---
subcategory: "Container Registry (CR)"
layout: "alicloud"
page_title: "Alicloud: alicloud_cr_internet_endpoint"
description: |-
  Provides a Alicloud CR Internet Endpoint resource.
---

# alicloud_cr_internet_endpoint

Provides a CR Internet Endpoint resource.

The public endpoint of the instance.

For information about CR Internet Endpoint and how to use it, see [What is Internet Endpoint](https://next.api.alibabacloud.com/document/cr/2018-12-01/UpdateInstanceEndpointStatus).

-> **NOTE:** Available since v1.287.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_cr_ee_instance" "internet_endpoint_pre_Instance" {
  instance_name      = "cspec-example-instance"
  default_oss_bucket = "true"
  payment_type       = "Subscription"
  period             = "1"
  instance_type      = "Basic"
}


resource "alicloud_cr_internet_endpoint" "default" {
  instance_id = alicloud_cr_ee_instance.internet_endpoint_pre_Instance.id
  entries {
    comment = "jj"
    entry   = "127.0.0.9/32"
  }
  entries {
    entry = "127.0.0.10/32"
  }
}
```

## Argument Reference

The following arguments are supported:
* `entries` - (Optional, List) The public network access control policy group. See [`entries`](#entries) below.
* `instance_id` - (Required, ForceNew) The instance ID.

### `entries`

The entries supports the following:
* `comment` - (Optional) The description of the access control policy group.
* `entry` - (Optional) The IP entry in the access control policy group. You can specify a CIDR block.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `status` - The running status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Internet Endpoint.
* `delete` - (Defaults to 5 mins) Used when delete the Internet Endpoint.
* `update` - (Defaults to 5 mins) Used when update the Internet Endpoint.

## Import

CR Internet Endpoint can be imported using the id, e.g.

```shell
$ terraform import alicloud_cr_internet_endpoint.example <instance_id>
```