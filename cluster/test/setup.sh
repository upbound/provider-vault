#!/usr/bin/env bash
set -aeuo pipefail

# setting up colors
BLU='\033[0;104m'
YLW='\033[0;33m'
GRN='\033[0;32m'
RED='\033[0;31m'
NOC='\033[0m' # No Color

echo_info(){
    printf "\n${BLU}%s${NOC}\n" "$1"
}
echo_step(){
    printf "\n${BLU}>>>>>>> %s${NOC}\n" "$1"
}
echo_step_completed(){
    printf "${GRN} [✔] %s${NOC}\n" "$1"
}

echo_info "Running setup.sh"

SCRIPT_DIR=$( cd -- $( dirname -- "${BASH_SOURCE[0]}" ) &> /dev/null && pwd )
KUBECONFIG_PATH="${SCRIPT_DIR}/../../kubeconfig"
if [ -f "${KUBECONFIG_PATH}" ]; then
    chmod 0600 ${KUBECONFIG_PATH}
fi

echo_info "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo_info "Waiting for all pods to come online..."
${KUBECTL} -n upbound-system wait --for=condition=Available deployment --all --timeout=5m

echo_info "Creating vault namespace"
${KUBECTL} create namespace vault --dry-run=client -o yaml | kubectl apply -f -

echo_info "Adding Hashicorp repo"
helm repo add hashicorp https://helm.releases.hashicorp.com --force-update

echo_info "Checking for HashiCorp Vault installation"
VAULT_DEPLOYMENT_STATUS=$(helm status vault -n vault|grep STATUS || true)
echo_info "$VAULT_DEPLOYMENT_STATUS"
if [[ "$VAULT_DEPLOYMENT_STATUS" == "STATUS: deployed" ]]; then
    echo "Uninstalling Hashicorp vault; need clean installation"
    helm uninstall vault -n vault --wait
fi

helm install vault hashicorp/vault -n vault

echo_info "Waiting for vault deployments"
${KUBECTL} -n vault wait --for=condition=Available deployment --all --timeout=5m

VAULT_NOT_RUNNING=true
while [[ "$VAULT_NOT_RUNNING" == "true" ]]; do
    echo_step "Waiting for vault-0 pod to run"
    sleep 5
    VAULT_STATUS=$(${KUBECTL} get pod -n vault vault-0 -o jsonpath='{.status}'|jq|grep phase|awk '{print $2}'|tr -d '"'|tr -d ',')
    echo_info "---$VAULT_STATUS---"
    if [[ "$VAULT_STATUS" == "Running" ]]; then
	VAULT_NOT_RUNNING=false
    fi
done

SECRET_EXISTS=$(${KUBECTL} get secret vault-auto-unseal-keys --ignore-not-found)
if [[ "$SECRET_EXISTS" = "" ]]; then
    echo "Initializing vault and obtaining unseal keys"
    ${KUBECTL} exec -n vault --stdin vault-0 -- vault operator init -format=yaml > vault-auto-unseal-keys.yaml
    echo "Creating secret with unseal keys"
    ${KUBECTL} create secret generic vault-auto-unseal-keys --from-file=keys=./vault-auto-unseal-keys.yaml
else
    echo "Secret vault-auto-unseal-keys exists. Fetching and decoding the keys."
    ${KUBECTL} get secret vault-auto-unseal-keys -o jsonpath='{.data.keys}' | base64 --decode > vault-auto-unseal-keys.yaml
fi

echo_info "vault-auto-unseal-keys.yaml"
cat vault-auto-unseal-keys.yaml
VAULT_ROOT_TOKEN=$(grep root_token vault-auto-unseal-keys.yaml |awk -F: '{print $2}'|sed 's/ //g')

echo_info "Unsealing vault - part 1/3"
${KUBECTL} exec -n vault --stdin vault-0 -- vault operator unseal -tls-skip-verify $(grep -A 5 unseal_keys_b64 vault-auto-unseal-keys.yaml |head -2|tail -1|sed 's/- //g')

echo_info "Unsealing vault - part 2/3"
${KUBECTL} exec -n vault --stdin vault-0 -- vault operator unseal -tls-skip-verify $(grep -A 5 unseal_keys_b64 vault-auto-unseal-keys.yaml |head -3|tail -1|sed 's/- //g')

echo_info "Unsealing vault - part 3/3"
${KUBECTL} exec -n vault --stdin vault-0 -- vault operator unseal -tls-skip-verify $(grep -A 5 unseal_keys_b64 vault-auto-unseal-keys.yaml |head -4|tail -1|sed 's/- //g')

echo_info "Writing provider-vault secret"
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: vault-creds
  namespace: vault
type: Opaque
stringData:
  credentials: '{"token": "$VAULT_ROOT_TOKEN"}'
EOF

echo_info "Applying providerconfig"
VAULT_0_POD_IP=$(${KUBECTL} get pod vault-0 -n vault -o yaml|grep podIP:|awk '{print $2}')
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: vault.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: vault-provider-config
spec:
  address: http://$VAULT_0_POD_IP:8200
  skip_child_token: true
  skip_tls_verify: true
  credentials:
    source: Secret
    secretRef:
      namespace: vault
      name: vault-creds
      key: credentials
EOF
