---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_ipv6_gateway"
description: |-
  Provides a Alicloud VPC Ipv6 Gateway resource.
---

# alicloud_vpc_ipv6_gateway

Provides a VPC Ipv6 Gateway resource. Gateway Based on Internet Protocol Version 6.

For information about VPC Ipv6 Gateway and how to use it, see [What is Ipv6 Gateway](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createipv6gateway).

-> **NOTE:** Available since v1.142.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

resource "alicloud_vpc" "defaultVpc" {
  description = "tf-testacc"
  enable_ipv6 = true
}

resource "alicloud_resource_manager_resource_group" "defaultRg" {
  display_name        = "tf-testacc-ipv6gateway510"
  resource_group_name = "${var.name}1"
}

resource "alicloud_resource_manager_resource_group" "changeRg" {
  display_name        = "tf-testacc-ipv6gateway151"
  resource_group_name = "${var.name}2"
}


resource "alicloud_vpc_ipv6_gateway" "default" {
  description       = "test"
  ipv6_gateway_name = var.name
  vpc_id            = alicloud_vpc.defaultVpc.id
  resource_group_id = alicloud_resource_manager_resource_group.defaultRg.id
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) The description of the IPv6 gateway. The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
* `ipv6_gateway_name` - (Optional) The name of the IPv6 gateway. The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
* `resource_group_id` - (Optional, Computed) The ID of the resource group to which the instance belongs.
* `spec` - (Optional, Deprecated since v1.208.0) IPv6 gateways do not distinguish between specifications. This parameter is no longer used.
* `tags` - (Optional, Map) The tags for the resource.
* `vpc_id` - (Required, ForceNew) The ID of the virtual private cloud (VPC) for which you want to create the IPv6 gateway.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `create_time` - The creation time of the resource.
* `status` - The status of the resource.  Available, Pending and Deleting.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Ipv6 Gateway.
* `delete` - (Defaults to 5 mins) Used when delete the Ipv6 Gateway.
* `update` - (Defaults to 5 mins) Used when update the Ipv6 Gateway.

## Import

VPC Ipv6 Gateway can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_ipv6_gateway.example <id>
```