#!/bin/bash

set -euo pipefail

VAULT_TOKEN=$(make vault.token)
export VAULT_TOKEN

# requires Vault to be port-forwarded
VAULT_ADDR="http://127.0.0.1:8200"
export VAULT_ADDR

if vault auth list | grep -q "approle"; then
    echo "Approle auth method already enabled"
else
    echo "Enabling approle auth method"
    vault auth enable approle
fi

echo "Creating development admin policy"
curl \
  --request POST \
  --header "X-Vault-Token: $VAULT_TOKEN" \
  --data '{"policy": "path \"*\" { capabilities = [\"create\", \"read\", \"update\", \"delete\", \"list\", \"sudo\"] }"}' \
  "$VAULT_ADDR/v1/sys/policy/dev-admin"

echo "Creating AppRole role my-role"
vault write auth/approle/role/my-role \
    token_type=batch \
    token_max_ttl=10m \
    bind_secret_id=false \
    secret_id_bound_cidrs="0.0.0.0/0" \
    token_bound_cidrs="0.0.0.0/0" \
    token_policies="dev-admin"

vault write auth/approle/role/my-role/role-id \
    role_id=my-role

echo "Authentication set up!"
