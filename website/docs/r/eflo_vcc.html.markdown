---
subcategory: "Eflo"
layout: "alicloud"
page_title: "Alicloud: alicloud_eflo_vcc"
description: |-
  Provides a Alicloud Eflo Vcc resource.
---

# alicloud_eflo_vcc

Provides a Eflo Vcc resource.



For information about Eflo Vcc and how to use it, see [What is Vcc](https://next.api.alibabacloud.com/document/eflo/2022-05-30/CreateVcc).

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

variable "cen_owner_id" {
  default = <<EOF
1013666993027780
EOF
}

variable "vsw_id" {
  default = "vsw-0jlpbevcjvsbe1b29cwkb"
}

variable "region_id" {
  default = "cn-wulanchabu"
}

variable "cen_id" {
  default = "cen-czo1beum1uku1vq4gp"
}

variable "vpc_id" {
  default = "vpc-0jll0blr5o98680qxqujn"
}

variable "zone_id" {
  default = "cn-wulanchabu-b"
}


resource "alicloud_eflo_vcc" "default" {
  description          = "example-tf-cross-account"
  access_could_service = false
  connection_type      = "CENTR"
  zone_id              = var.zone_id
  cen_owner_id         = var.cen_owner_id
  bandwidth            = "1000"
  vcc_name             = "vcc-example-cross-account"
  vswitch_id           = var.vsw_id
  vpc_id               = var.vpc_id
  cen_id               = var.cen_id
}
```

## Argument Reference

The following arguments are supported:
* `access_could_service` - (Optional) Start access to cloud service, optional value:
  - `true`: enable access to cloud services
  - `false`: does not enable access to cloud services

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `bandwidth` - (Optional, Int) Bandwidth, in Mbps. The minimum value is 1000, representing 1Gbps bandwidth; The maximum value is 400000, representing 400Gbps bandwidth.

-> **NOTE:**  1Gbps = 1000Mbps

* `bgp_asn` - (Optional, Int) Bgp as number

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `bgp_cidr` - (Optional, ForceNew) BGP network segment.
* `cen_id` - (Optional) CEN
* `cen_owner_id` - (Optional, ForceNew) Account to which cen belongs
* `connection_type` - (Optional) Connection method, optional value:
  - `VPC`
  - `CENTR`
* `description` - (Optional) Description

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `resource_group_id` - (Optional, Computed) Resource group id
* `tags` - (Optional, Map) The tag of the resource
* `vswitch_id` - (Optional) vSwitch
* `vcc_name` - (Optional) Lingjun Connection Name
* `vpc_id` - (Required) VPC
* `zone_id` - (Optional) Availability Zone

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.
* `region_id` - Region.
* `status` - Status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 25 mins) Used when create the Vcc.
* `delete` - (Defaults to 143 mins) Used when delete the Vcc.
* `update` - (Defaults to 70 mins) Used when update the Vcc.

## Import

Eflo Vcc can be imported using the id, e.g.

```shell
$ terraform import alicloud_eflo_vcc.example <vcc_id>
```