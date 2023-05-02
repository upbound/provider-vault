#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"

echo "Creating cloud credential secret..."
${KUBECTL} -n upbound-system create secret generic provider-secret --from-literal=credentials="{\"token\":\"${UPTEST_CLOUD_CREDENTIALS}\"}" --dry-run=client -o yaml | ${KUBECTL} apply -f -
echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
${KUBECTL} -n upbound-system wait --for=condition=Available deployment --all --timeout=5m

echo "Creating a default provider config..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: vault.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  add_address_to_env: false
  address: http://127.0.0.1:8200
  headers: {name: test, value: "e2e"}
  max_lease_ttl_seconds: 300
  max_retries: 10
  max_retries_ccc: 10
  namespace: vault
  skip_child_token: true
  skip_get_vault_version: true
  skip_tls_verify: true
  tls_server_name: "1.15.0"
  vault_version_override: "false"
  credentials:
    secretRef:
      name: provider-secret
      namespace: upbound-system
      key: credentials
      name: provider-creds
      namespace: upbound-system
    source: Secret
EOF

echo "Creating vault namesapce"
${KUBECTL} create namespace vault --dry-run=client -o yaml | kubectl apply -f -

echo "Adding Hashicorp repo"
helm repo add hashicorp https://helm.releases.hashicorp.com --force-update

echo "Installing Hashicorp vault"
# TODO: Conditional uninstall
# helm uninstall vault -n vault --wait --cascade
helm install vault hashicorp/vault -n vault

echo "Waiting for vault deployments"
${KUBECTL} -n vault wait --for=condition=Available deployment --all --timeout=5m

echo "Initializing vault and obtaining unseal keys"
${KUBECTL} exec --namespace vault --stdin --tty vault-0 -- vault operator init -format=yaml > vault-auto-unseal-keys.txt

echo "Unsealing vault - part 1/3"
${KUBECTL} exec --namespace vault --stdin vault-0 -- vault operator unseal -tls-skip-verify $(grep -A 5 unseal_keys_b64 vault-auto-unseal-keys.txt |head -2|tail -1|sed 's/- //g')

echo "Unsealing vault - part 2/3"
${KUBECTL} exec --namespace vault --stdin vault-0 -- vault operator unseal -tls-skip-verify $(grep -A 5 unseal_keys_b64 vault-auto-unseal-keys.txt |head -3|tail -1|sed 's/- //g')

echo "Unsealing vault - part 3/3"
${KUBECTL} exec --namespace vault --stdin vault-0 -- vault operator unseal -tls-skip-verify $(grep -A 5 unseal_keys_b64 vault-auto-unseal-keys.txt |head -4|tail -1|sed 's/- //g')

echo "Enable secrets, policy and token"

export VAULT_AUTO_UNSEAL_ROOT_TOKEN=`${KUBECTL} exec --namespace vault --stdin --tty vault-0 -- $(grep root_token vault-auto-unseal-keys.txt |awk -F: '{print $2}'|sed 's/ //g')`

echo "Logging into vault with ${VAULT_AUTO_UNSEAL_ROOT_TOKEN}"

${KUBECTL} exec --namespace vault --stdin vault-0 -- vault login -tls-skip-verify ${VAULT_AUTO_UNSEAL_ROOT_TOKEN}
${KUBECTL} exec --namespace vault --stdin vault-0 -- vault audit enable file file_path=/vault/logs/audit.log
${KUBECTL} exec --namespace vault --stdin vault-0 -- vault secrets enable transit
${KUBECTL} exec --namespace vault --stdin vault-0 -- vault write -f transit/keys/autounseal
${KUBECTL} exec --namespace vault --stdin vault-0 -- vault policy write autounseal /vault/unseal/autounseal.hcl
${KUBECTL} exec --namespace vault --stdin vault-0 -- vault token create -policy="autounseal" -wrap-ttl=12000 -format=yaml > .vault-auto-unseal-token.txt

# More useful setup info
# https://itnext.io/vault-cluster-with-auto-unseal-on-kubernetes-8e469f9cdcfd
