terraform {
  required_providers {
    awx = {
      source = "denouche/awx"
    }
  }
}

data "awx_workflow_job_template" "default" {
  name = "default"
}

output "awx_workflow_job_template_id" {
  value = data.awx_workflow_job_template.default.id
}
