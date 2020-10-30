terraform {
  required_providers {
    meraki = {
      versions = ["0.2"]
      source = "hashicorp.com/edu/meraki"
    }
  }
}

data "meraki_organizations" "all" {}
data "meraki_devices" "devices" {}


output "org" {
  value = data.meraki_organizations.all
}

output "devices" {
  value = data.meraki_devices.devices
}