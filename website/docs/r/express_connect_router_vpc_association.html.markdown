---
subcategory: "Express Connect Router"
layout: "alicloud"
page_title: "Alicloud: alicloud_express_connect_router_vpc_association"
description: |-
  Provides a Alicloud Express Connect Router Express Connect Router Vpc Association resource.
---

# alicloud_express_connect_router_vpc_association

Provides a Express Connect Router Express Connect Router Vpc Association resource.

Express Connect router and VPC association object  .

For information about Express Connect Router Express Connect Router Vpc Association and how to use it, see [What is Express Connect Router Vpc Association](https://www.alibabacloud.com/help/en/express-connect/developer-reference/api-expressconnectrouter-2023-09-01-createexpressconnectrouterassociation).

-> **NOTE:** Available since v1.224.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_vpc" "default8qAtD6" {
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_express_connect_router_express_connect_router" "defaultM9YxGW" {
  alibaba_side_asn = "65533"
}

data "alicloud_account" "current" {
}

resource "alicloud_express_connect_router_vpc_association" "default" {
  ecr_id = alicloud_express_connect_router_express_connect_router.defaultM9YxGW.id
  allowed_prefixes = [
    "172.16.4.0/24",
    "172.16.3.0/24",
    "172.16.2.0/24",
    "172.16.1.0/24"
  ]
  vpc_owner_id          = data.alicloud_account.current.id
  association_region_id = "cn-hangzhou"
  vpc_id                = alicloud_vpc.default8qAtD6.id
}
```

## Argument Reference

The following arguments are supported:
* `allowed_prefixes` - (Optional, List) The list of allowed route prefixes.
* `association_region_id` - (Required, ForceNew) The region ID of the VPC or Transit Router (TR).
* `description` - (Optional, Available since v1.277.0) The description of the associated resource. The length must be 0 to 128 characters.
* `ecr_id` - (Required, ForceNew) The ID of the Express Connect router instance.
* `vpc_id` - (Required, ForceNew) The ID of the VPC instance.
* `vpc_owner_id` - (Optional, ForceNew, Computed, Int) The Alibaba Cloud account ID to which the VPC belongs.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<ecr_id>:<association_id>:<vpc_id>`.
* `association_id` - The association ID between the Express Connect router and a VPC or TR.
* `create_time` - The time when the association was created.
* `status` - The deployment status of the associated resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Express Connect Router Vpc Association.
* `delete` - (Defaults to 5 mins) Used when delete the Express Connect Router Vpc Association.
* `update` - (Defaults to 5 mins) Used when update the Express Connect Router Vpc Association.

## Import

Express Connect Router Express Connect Router Vpc Association can be imported using the id, e.g.

```shell
$ terraform import alicloud_express_connect_router_vpc_association.example <ecr_id>:<association_id>:<vpc_id>
```