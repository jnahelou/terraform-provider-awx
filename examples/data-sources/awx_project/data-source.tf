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

output "awx_project_id" {
  value = data.awx_project.default.id
}
