terraform {
  required_providers {
    awx = {
      source = "denouche/awx"
    }
  }
}

data "awx_organization" "default" {
  name = "Default"
}
