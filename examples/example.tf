terraform {
  required_providers {
    terraform = {
      source  = ""
      version = "0.0.1"
    }
  }
}

variable "fivetran_api_key" {
  type = string
}

provider "terraform" {
  authorization = var.fivetran_api_key
}


resource "fivetran_connector" "my_connector" {
  group_id          = "test"
  pause_after_trial = "true"
  paused            = "true"
  service           = "google_analytics"
  sync_frequency    = "on_demand"
  config {}
  destination_schema {}
}
