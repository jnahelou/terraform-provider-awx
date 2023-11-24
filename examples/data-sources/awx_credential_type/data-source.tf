terraform {
  required_providers {
    awx = {
      source = "denouche/awx"
    }
  }
}

data "awx_credential_type" "example" {
  id = 1
}
