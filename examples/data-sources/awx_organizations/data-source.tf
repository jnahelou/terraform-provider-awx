terraform {
  required_providers {
    awx = {
      source = "denouche/awx"
    }
  }
}

data "awx_organizations" "default" {}

output "awx_organizations" {
  value = data.awx_organizations.default
}
