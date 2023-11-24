terraform {
  required_providers {
    awx = {
      source = "denouche/awx"
    }
  }
}

data "awx_schedule" "default" {
  name = "default"
}

output "awx_schedule_id" {
  value = data.awx_schedule.default.id
}
