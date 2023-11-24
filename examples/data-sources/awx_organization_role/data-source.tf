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

data "awx_organization_role" "admin" {
  organization_id = data.awx_organization.default.id
  name            = "Project Admin"
}

output "awx_organization_role_admin_id" {
  value = data.awx_organization_role.admin.id
}
