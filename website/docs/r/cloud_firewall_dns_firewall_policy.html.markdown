---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_dns_firewall_policy"
description: |-
  Provides a Alicloud Cloud Firewall Dns Firewall Policy resource.
---

# alicloud_cloud_firewall_dns_firewall_policy

Provides a Cloud Firewall Dns Firewall Policy resource.

DNS Firewall Policy.

For information about Cloud Firewall Dns Firewall Policy and how to use it, see [What is Dns Firewall Policy](https://next.api.alibabacloud.com/document/Cloudfw/2017-12-07/AddDnsFirewallPolicy).

-> **NOTE:** Available since v1.286.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_cloud_firewall_dns_firewall_policy" "default" {
  description      = "example"
  ip_version       = "4"
  priority         = "1"
  source_type      = "net"
  acl_action       = "accept"
  destination_type = "group"
  release          = "true"
  source           = "8.8.8.8/32"
}
```

## Argument Reference

The following arguments are supported:
* `acl_action` - (Required) The action that Cloud Firewall performs on traffic based on the access control policy. Valid values:
  - `accept`: Allow
  - `drop`: Deny
  - `log`: Monitor
* `description` - (Required) The description of the DNS firewall access control policy.
* `destination_type` - (Required) The type of the destination address in the access control policy. Valid values:
  - `net`: Destination CIDR block
  - `group`: Destination address book
  - `domain`: Destination domain name
  - `location`: Destination region
* `direction` - (Optional, ForceNew) The direction of the traffic to which the access control policy applies. Valid values:
  - `in`: Inbound traffic
  - `out`: Outbound traffic
* `ip_version` - (Required, ForceNew, Int) The supported IP address version.

Valid values:
  - `4`: IPv4 address
  - `6`: IPv6 address
* `priority` - (Required, Int) The priority of the access control policy before the modification.
* `release` - (Required) The enabled status of the access control policy. By default, the policy is enabled after it is created. Valid values:
  - `true`: Enable the access control policy.
  - `false`: Do not enable the access control policy.
* `source` - (Required) The source address in the access control policy. Valid values:
  - When `SourceType` is set to `net`, Source specifies the source CIDR block. Example: 10.2.XX.XX/24.
  - When `SourceType` is set to `group`, Source specifies the name of the source address book. Example: db_group.
* `source_type` - (Required) The type of the source address in the DNS firewall access control policy. Valid values:
  - `net`: Source CIDR block
  - `group`: Source address book

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `destination_addrs` - The list of addresses in the destination address book.
* `destination_group_type` - The type of the destination address book in the access control policy.
* `hit_last_time` - The timestamp of the most recent hit.
* `hit_times` - The number of times the access control policy is matched.
* `source_addrs` - The source addresses.
* `source_group_type` - The type of the source address book in the access control policy.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Dns Firewall Policy.
* `delete` - (Defaults to 5 mins) Used when delete the Dns Firewall Policy.
* `update` - (Defaults to 5 mins) Used when update the Dns Firewall Policy.

## Import

Cloud Firewall Dns Firewall Policy can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_firewall_dns_firewall_policy.example <acl_uuid>
```