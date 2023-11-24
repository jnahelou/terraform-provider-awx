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

data "awx_inventory_role" "example" {
  name         = "Admin"
  inventory_id = data.awx_inventory.example.id
}
