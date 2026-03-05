---
subcategory: "Eflo"
layout: "alicloud"
page_title: "Alicloud: alicloud_eflo_er_attachment"
description: |-
  Provides a Alicloud Eflo Er Attachment resource.
---

# alicloud_eflo_er_attachment

Provides a Eflo Er Attachment resource.

Network instance connection.

For information about Eflo Er Attachment and how to use it, see [What is Er Attachment](https://next.api.alibabacloud.com/document/eflo/2022-05-30/CreateErAttachment).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-wulanchabu"
}

variable "zone_id" {
  default = "cn-wulanchabu-b"
}

variable "region" {
  default = "cn-wulanchabu"
}

resource "alicloud_eflo_er" "ER" {
  er_name        = "er_at_example"
  master_zone_id = var.zone_id
}

resource "alicloud_vpc" "VPC" {
  is_default  = false
  dry_run     = false
  cidr_block  = "192.168.0.0/16"
  enable_ipv6 = false
}

resource "alicloud_vswitch" "VSW" {
  vpc_id     = alicloud_vpc.VPC.id
  zone_id    = var.zone_id
  cidr_block = "192.168.0.0/24"
}

resource "alicloud_eflo_vcc" "VCC" {
  connection_type = "VPC"
  zone_id         = var.zone_id
  vswitch_id      = alicloud_vswitch.VSW.id
  vpc_id          = alicloud_vpc.VPC.id
  bandwidth       = "1000"
  vcc_name        = "ER_AT_TEST"
}


resource "alicloud_eflo_er_attachment" "default" {
  instance_id            = alicloud_eflo_vcc.VCC.id
  resource_tenant_id     = "1511928242963727"
  er_id                  = alicloud_eflo_er.ER.id
  er_attachment_name     = "er_at_tf_example_vcc"
  instance_type          = "VCC"
  auto_receive_all_route = true
}
```

## Argument Reference

The following arguments are supported:
* `auto_receive_all_route` - (Required, ForceNew) Automatically receive all routes
* `er_attachment_name` - (Optional) Network instance connection name
* `er_id` - (Required, ForceNew) Lingjun HUB
* `instance_id` - (Required, ForceNew) Instance ID
* `instance_type` - (Required, ForceNew) Instance type
* `resource_tenant_id` - (Required, ForceNew) Tenant ID to which the resource belongs

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<er_id>:<er_attachment_id>`.
* `create_time` - Creation time.
* `er_attachment_id` - The connection ID of the network instance.
* `region_id` - Geographical information.
* `status` - Status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Er Attachment.
* `delete` - (Defaults to 5 mins) Used when delete the Er Attachment.
* `update` - (Defaults to 5 mins) Used when update the Er Attachment.

## Import

Eflo Er Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_eflo_er_attachment.example <er_id>:<er_attachment_id>
```