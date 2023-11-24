terraform {
  required_providers {
    awx = {
      source = "denouche/awx"
    }
  }
}

data "awx_notification_template" "default" {
  name = "demo"
}

output "awx_job_template_id" {
  value = data.awx_job_template.default.id
}
