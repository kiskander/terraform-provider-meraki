terraform {
  required_providers {
    meraki = {
      versions = ["0.2"]
      source = "local/devnet/meraki"
    }
  }
}

# data "meraki_organizations" "all" {}

# data "meraki_devices" "devices" {
#   org_id = "681155"
# }

data "meraki_org_create" "organizations" {
  org_name = "Kareem Test"
}

# output "org" {
#   value = data.meraki_organizations.all
# }

# output "devices" {
#   value = data.meraki_devices.devices
# }

output "neworg"{
  value = data.meraki_org_create.organizations
}

