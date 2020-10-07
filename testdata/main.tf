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
      source  = "kvrhdn/honeycombio"
      version = "~> 0.0.9"
    }
  }
}

variable "github_run_number" {
  type = number
}

resource "honeycombio_marker" "dummy_resource_1" {
  message = "Integration - run ${var.github_run_number} - marker #1"
  dataset = "go-tfe-run"
}

resource "honeycombio_marker" "dummy_resource_2" {
  message = "Integration - run ${var.github_run_number} - marker #2"
  dataset = "go-tfe-run"
}

output "marker_message_1" {
  value = honeycombio_marker.dummy_resource_1.message
}

output "marker_message_2" {
  value = honeycombio_marker.dummy_resource_2.message
}
