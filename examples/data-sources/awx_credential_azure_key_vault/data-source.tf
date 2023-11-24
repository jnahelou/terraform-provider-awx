terraform {
  required_providers {
    awx = {
      source = "denouche/awx"
    }
  }
}

data "awx_credential_azure_key_vault" "example" {
  credential_id = 2
}

output "example" {
  value = data.awx_credential_azure_key_vault.example.name
}
