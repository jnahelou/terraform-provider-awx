terraform {
  required_providers {
    awx = {
      source = "denouche/awx"
    }
  }
}

data "awx_team" "default" {
  name = "default"
}

output "awx_team_id" {
  value = data.awx_team.default.id
}
