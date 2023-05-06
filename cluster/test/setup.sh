#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"

echo "Creating cloud credential secret..."
${KUBECTL} -n upbound-system create secret generic provider-secret --from-literal=credentials="{\"token\":\"${UPTEST_CLOUD_CREDENTIALS}\"}" --dry-run=client -o yaml | ${KUBECTL} apply -f -
echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
${KUBECTL} -n upbound-system wait --for=condition=Available deployment --all --timeout=5m

echo "Creating vault namesapce"
${KUBECTL} create namespace vault --dry-run=client -o yaml | kubectl apply -f -

echo "Adding Hashicorp repo"
helm repo add hashicorp https://helm.releases.hashicorp.com --force-update

echo "Checking for HashiCorp Vault installation"
VAULT_DEPLOYMENT_STATUS=$(helm status vault -n vault|grep STATUS || true)
echo "$VAULT_DEPLOYMENT_STATUS"
if [[ "$VAULT_DEPLOYMENT_STATUS" == "STATUS: deployed" ]]; then
    echo "Uninstalling Hashicorp vault; need clean installation"
    helm uninstall vault -n vault --wait
fi

helm install vault hashicorp/vault -n vault

echo "Waiting for vault deployments"
${KUBECTL} -n vault wait --for=condition=Available deployment --all --timeout=5m

VAULT_NOT_RUNNING=true
while [[ "$VAULT_NOT_RUNNING" == "true" ]]; do
    echo "Waiting for vault-0 pod to run"
    sleep 5
    VAULT_STATUS=$(${KUBECTL} get pod -n vault vault-0 -o jsonpath='{.status}'|jq|grep phase|awk '{print $2}'|tr -d '"'|tr -d ',')
    echo "---$VAULT_STATUS---"
    if [[ "$VAULT_STATUS" == "Running" ]]; then
	VAULT_NOT_RUNNING=false
    fi
done

echo "Initializing vault and obtaining unseal keys"
${KUBECTL} exec -n vault --stdin vault-0 -- vault operator init -format=yaml > vault-auto-unseal-keys.yaml

echo "vault-auto-unseal-keys.yaml"
cat vault-auto-unseal-keys.yaml
VAULT_ROOT_TOKEN=$(grep root_token vault-auto-unseal-keys.yaml |awk -F: '{print $2}'|sed 's/ //g')

echo "Unsealing vault - part 1/3"
${KUBECTL} exec -n vault --stdin vault-0 -- vault operator unseal -tls-skip-verify $(grep -A 5 unseal_keys_b64 vault-auto-unseal-keys.yaml |head -2|tail -1|sed 's/- //g')

echo "Unsealing vault - part 2/3"
${KUBECTL} exec -n vault --stdin vault-0 -- vault operator unseal -tls-skip-verify $(grep -A 5 unseal_keys_b64 vault-auto-unseal-keys.yaml |head -3|tail -1|sed 's/- //g')

echo "Unsealing vault - part 3/3"
${KUBECTL} exec -n vault --stdin vault-0 -- vault operator unseal -tls-skip-verify $(grep -A 5 unseal_keys_b64 vault-auto-unseal-keys.yaml |head -4|tail -1|sed 's/- //g')

echo "Writing provider-vault secret"
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: vault-creds
  namespace: vault
type: Opaque
stringData:
  credentials: |
    {
      "token_name": "vault-creds-test-token",
      "token": "$VAULT_ROOT_TOKEN"
    }
EOF

echo "Applying providerconfig"
VAULT_0_POD_IP=$(${KUBECTL} get pod vault-0 -n vault -o yaml|grep podIP:|awk '{print $2}')
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: vault.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: vault-provider-config
spec:
  # address: http://127.0.0.1:8200
  address: http://$VAULT_0_POD_IP:8200
  add_address_to_env: false
  headers: {name: test, value: "e2e"}
  max_lease_ttl_seconds: 300
  max_retries: 10
  max_retries_ccc: 10
  namespace: vault
  skip_child_token: true
  skip_get_vault_version: true
  skip_tls_verify: true
  tls_server_name: ""
  vault_version_override: "1.12.0"
  credentials:
    source: Secret
    secretRef:
      name: vault-creds
      namespace: vault
      key: credentials
EOF

# More useful setup info
# https://itnext.io/vault-cluster-with-auto-unseal-on-kubernetes-8e469f9cdcfd

echo "Test setup complete"
echo "Running upjet vault tests"
