terraform {
  required_providers {
    awx = {
      source = "denouche/awx"
    }
  }
}

data "awx_credential" "example" {
  id = 1
}
