terraform {
  required_providers {
    awx = {
      source = "denouche/awx"
    }
  }
}

data "awx_execution_environment" "by_id" {
  id = 1
}

data "awx_execution_environment" "by_name" {
  name = "awx"
}

output "name_by_id" {
  value = data.awx_execution_environment.by_id.name
}

output "id_by_name" {
  value = data.awx_execution_environment.by_name.id
}
