terraform {
  required_providers {
    awx = {
      source = "denouche/awx"
    }
  }
}

data "awx_credentials" "example" {}
