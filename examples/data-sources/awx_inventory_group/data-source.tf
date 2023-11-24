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

data "awx_inventory" "example" {
  organization_id = data.awx_organization.default.id
  name            = "Demo Inventory"
}

data "awx_inventory_group" "default" {
  name         = "Demo"
  inventory_id = data.awx_inventory.example.id
}
