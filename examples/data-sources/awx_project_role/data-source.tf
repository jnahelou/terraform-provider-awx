terraform {
  required_providers {
    awx = {
      source = "denouche/awx"
    }
  }
}

data "awx_project" "default" {
  name = "Default"
}

data "awx_project_role" "default" {
  name       = "Admin"
  project_id = data.awx_project.default.id
}

output "awx_project_role_id" {
  value = data.awx_project_role.default.id
}
