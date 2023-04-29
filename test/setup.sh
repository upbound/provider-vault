#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"
echo "Waiting until configuration package is healthy/installed..."
${KUBECTL} wait configuration.pkg platform-ref-aws --for=condition=Healthy --timeout 5m
${KUBECTL} wait configuration.pkg platform-ref-aws --for=condition=Installed --timeout 5m

echo "Creating cloud credential secret..."
${KUBECTL} -n upbound-system create secret generic aws-creds --from-file=credentials="UPBOUND_AWS_CREDENTIALS_PATH" \
    --dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Waiting until provider-aws is healthy..."
${KUBECTL} wait provider.pkg upbound-provider-aws --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
"${KUBECTL}" -n upbound-system wait --for=condition=Available deployment --all --timeout=5m

echo "Waiting for all XRDs to be established..."
kubectl wait xrd --all --for condition=Established

echo "Creating a default provider config..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    secretRef:
      key: credentials
      name: aws-creds
      namespace: upbound-system
    source: Secret
EOF

echo "Installing HashiCorp Vault"
${KUBECTL} create namespace Vault
i${HELM} repo add hashicorp https://helm.releases.hashicorp.com
${HELM} install vault hashicorp/vault -n vault
# Initialize vault
${KUBECTL} exec --namespace=vault --stdin --tty vault-0 -- vault operator init -format=yaml > vault-auto-unseal-keys.txt   
# Unseal it
${KUBECTL} exec --namespace=vault --stdin --tty vault-0 -- vault operator unseal -tls-skip-verify $(grep -A 5 unseal_keys_b64 vault-auto-unseal-keys.txt |head -2|tail -1|sed 's/- //g')
${KUBECTL} exec --namespace=vault --stdin --tty vault-0 -- vault operator unseal -tls-skip-verify $(grep -A 5 unseal_keys_b64 vault-auto-unseal-keys.txt |head -3|tail -1|sed 's/- //g')
${KUBECTL} exec --namespace=vault --stdin --tty vault-0 -- vault operator unseal -tls-skip-verify $(grep -A 5 unseal_keys_b64 vault-auto-unseal-keys.txt |head -4|tail -1|sed 's/- //g')
VAULT_AUTO_UNSEAL_ROOT_TOKEN=`${KUBECTL} exec --namespace=vault --stdin --tty vault-0 -- $(grep root_token vault-auto-unseal-keys.txt |awk -F: '{print $2}'|sed 's/ //g')`

# More useful setup info
# https://itnext.io/vault-cluster-with-auto-unseal-on-kubernetes-8e469f9cdcfd
