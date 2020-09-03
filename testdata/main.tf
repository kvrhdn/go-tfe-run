terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "kvrhdn"

    workspaces {
      name = "go-tfe-run"
    }
  }
  required_providers {
    honeycombio = {
      source = "kvrhdn/honeycombio"
      version = "~> 0.0.9"
    }
  }
}

variable "github_run_number" {
  type = number
}

resource "honeycombio_marker" "dummy_resource" {
  message = "Integration - run ${var.github_run_number}"
  dataset = "go-tfe-run"
}

output "marker_message" {
  value = honeycombio_marker.dummy_resource.message
}
