terraform {
  required_providers {
    meraki = {
      versions = ["0.2"]
      source = "hashicorp.com/edu/meraki"
    }
  }
}

data "meraki_organizations" "all" {}
output "all" {
  value = data.meraki_organizations.all
}