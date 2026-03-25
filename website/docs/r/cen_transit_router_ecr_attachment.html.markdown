---
subcategory: "Cloud Enterprise Network (CEN)"
layout: "alicloud"
page_title: "Alicloud: alicloud_cen_transit_router_ecr_attachment"
description: |-
  Provides a Alicloud Cloud Enterprise Network (CEN) Transit Router Ecr Attachment resource.
---

# alicloud_cen_transit_router_ecr_attachment

Provides a Cloud Enterprise Network (CEN) Transit Router Ecr Attachment resource.



For information about Cloud Enterprise Network (CEN) Transit Router Ecr Attachment and how to use it, see [What is Transit Router Ecr Attachment](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createtransitrouterecrattachment).

-> **NOTE:** Available since v1.235.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

variable "asn" {
  default = "4200000667"
}

resource "alicloud_express_connect_router_express_connect_router" "defaultO8Hcfx" {
  alibaba_side_asn = var.asn
  ecr_name         = var.name
}

resource "alicloud_cen_instance" "defaultQKBiay" {
  cen_instance_name = var.name
}

resource "alicloud_cen_transit_router" "defaultQa94Y1" {
  cen_id              = alicloud_cen_instance.defaultQKBiay.id
  transit_router_name = var.name
}

data "alicloud_account" "current" {
}

resource "alicloud_express_connect_router_tr_association" "defaultedPu6c" {
  association_region_id   = "cn-hangzhou"
  ecr_id                  = alicloud_express_connect_router_express_connect_router.defaultO8Hcfx.id
  cen_id                  = alicloud_cen_instance.defaultQKBiay.id
  transit_router_id       = alicloud_cen_transit_router.defaultQa94Y1.transit_router_id
  transit_router_owner_id = data.alicloud_account.current.id
}

resource "alicloud_cen_transit_router_ecr_attachment" "default" {
  ecr_id                                = alicloud_express_connect_router_express_connect_router.defaultO8Hcfx.id
  cen_id                                = alicloud_express_connect_router_tr_association.defaultedPu6c.cen_id
  transit_router_ecr_attachment_name    = var.name
  transit_router_attachment_description = var.name
  transit_router_id                     = alicloud_cen_transit_router.defaultQa94Y1.transit_router_id
  ecr_owner_id                          = data.alicloud_account.current.id
}
```

## Argument Reference

The following arguments are supported:
* `cen_id` - (Optional, ForceNew) The ID of the Cloud Enterprise Network (CEN) instance.
* `ecr_id` - (Required, ForceNew) The ID of the associated Express Connect router (ECR) instance.
* `ecr_owner_id` - (Optional, ForceNew, Int) The ID of the Alibaba Cloud account (main account) to which the ECR instance belongs. The default value is the ID of the currently logged-in Alibaba Cloud account.

-> **NOTE:**  If you want to attach a network instance from another Alibaba Cloud account, you must specify this parameter.

* `order_type` - (Optional, Computed, Available since v1.274.0) The payer of the network instance. Valid values:
  - `PayByCenOwner`: The owner of the transit router instance pays for the connection fee and traffic processing fee of the ECR instance.
  - `PayByResourceOwner`: The owner of the ECR instance pays for the connection fee and traffic processing fee of the ECR instance.
* `tags` - (Optional, Map) The list of tags.
* `transit_router_attachment_description` - (Optional) The new description of the ECR attachment.
The description can be empty or contain 1 to 256 characters, and must not start with http:// or https://.
* `transit_router_ecr_attachment_name` - (Optional) The new name of the ECR attachment.
The name can be empty or contain 1 to 128 characters, and must not start with http:// or https://.
* `transit_router_id` - (Optional, ForceNew) The ID of the Enterprise Edition transit router instance.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the ECR instance attachment.
* `status` - The status of the ECR attachment.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Transit Router Ecr Attachment.
* `delete` - (Defaults to 5 mins) Used when delete the Transit Router Ecr Attachment.
* `update` - (Defaults to 5 mins) Used when update the Transit Router Ecr Attachment.

## Import

Cloud Enterprise Network (CEN) Transit Router Ecr Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_cen_transit_router_ecr_attachment.example <transit_router_attachment_id>
```