---
subcategory: "Cloud Enterprise Network (CEN)"
layout: "alicloud"
page_title: "Alicloud: alicloud_cen_transit_router_vpc_attachment"
description: |-
  Provides a Alicloud Cloud Enterprise Network (CEN) Transit Router Vpc Attachment resource.
---

# alicloud_cen_transit_router_vpc_attachment

Provides a Cloud Enterprise Network (CEN) Transit Router Vpc Attachment resource.

CenTransitRouterVpcAttachment  .

For information about Cloud Enterprise Network (CEN) Transit Router Vpc Attachment and how to use it, see [What is Transit Router Vpc Attachment](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createtransitroutervpcattachment).

-> **NOTE:** Available since v1.126.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

data "alicloud_cen_transit_router_available_resources" "default" {
}

locals {
  master_zone = data.alicloud_cen_transit_router_available_resources.default.resources[0].master_zones[0]
  slave_zone  = data.alicloud_cen_transit_router_available_resources.default.resources[0].slave_zones[1]
}

resource "alicloud_vpc" "example" {
  vpc_name   = var.name
  cidr_block = "192.168.0.0/16"
}

resource "alicloud_vswitch" "example_master" {
  vswitch_name = var.name
  cidr_block   = "192.168.1.0/24"
  vpc_id       = alicloud_vpc.example.id
  zone_id      = local.master_zone
}

resource "alicloud_vswitch" "example_slave" {
  vswitch_name = var.name
  cidr_block   = "192.168.2.0/24"
  vpc_id       = alicloud_vpc.example.id
  zone_id      = local.slave_zone
}

resource "alicloud_cen_instance" "example" {
  cen_instance_name = var.name
  protection_level  = "REDUCED"
}

resource "alicloud_cen_transit_router" "example" {
  transit_router_name = var.name
  cen_id              = alicloud_cen_instance.example.id
}

resource "alicloud_cen_transit_router_vpc_attachment" "example" {
  cen_id            = alicloud_cen_instance.example.id
  transit_router_id = alicloud_cen_transit_router.example.transit_router_id
  vpc_id            = alicloud_vpc.example.id
  zone_mappings {
    zone_id    = local.master_zone
    vswitch_id = alicloud_vswitch.example_master.id
  }
  zone_mappings {
    zone_id    = local.slave_zone
    vswitch_id = alicloud_vswitch.example_slave.id
  }
  transit_router_vpc_attachment_name    = var.name
  transit_router_attachment_description = var.name
}
```

## Argument Reference

The following arguments are supported:
* `auto_publish_route_enabled` - (Optional) Specifies whether the Enterprise Edition transit router automatically advertises routes to the VPC instance.
  - `false` (default): No.
  - `true`: Yes.
* `cen_id` - (Optional) The ID of the Cloud Enterprise Network (CEN) instance.

-> **NOTE:** This parameter is only evaluated during resource creation and deletion. Modifying it in isolation will not trigger any action.

* `dry_run` - (Optional) Specifies whether to perform a dry run of the request, including checks for permissions and instance status. Valid values:
  - `false` (default): Sends a normal request. If the checks pass, the VPC attachment is created immediately.
  - `true`: Sends a dry run request for validation only. The VPC attachment is not created. Validation includes checking whether required parameters are specified and whether the request format is valid. If validation fails, an error is returned. If validation passes, the error code `DryRunOperation` is returned.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `force_delete` - (Optional, Available since v1.230.1) Specifies whether to forcibly delete the VPC attachment. Valid values:
  - `false` (default): Before deleting the VPC attachment, the system checks for dependencies such as associated forwarding entries or route learning configurations. If dependencies exist, the deletion is blocked and an error is returned.
  - `true`: When deleting the VPC attachment, all related dependencies are automatically removed.

-> **NOTE:** This parameter configures deletion behavior and is only evaluated when Terraform attempts to destroy the resource. Changes to this parameter during updates are stored but have no immediate effect.

* `order_type` - (Optional, Computed, Available since v1.276.0) The billing party for the network instance. Valid values:
  - `PayByCenOwner`: The account that owns the Cloud Enterprise Network (CEN) instance pays for the network instance.
  - `PayByResourceOwner`: The account that owns the network instance pays for the network instance.
* `payment_type` - (Optional, ForceNew, Computed) Billing method of the VPC attachment.
Valid value: `POSTPAY`, which indicates pay-as-you-go billing.
* `tags` - (Optional, Map) List of tags.
You can specify up to 20 tags at a time.
* `transit_router_attachment_description` - (Optional) The description of the VPC attachment.
The description can be empty or contain 1 to 256 characters, and must not start with http:// or https://.
* `transit_router_id` - (Optional, ForceNew, Computed) The ID of the Enterprise Edition transit router instance.  
* `transit_router_vpc_attachment_name` - (Optional, Computed, Available since v1.230.1) The name of the VPC attachment.
The name can be empty or contain 1 to 128 characters, and must not start with http:// or https://.
* `transit_router_vpc_attachment_options` - (Optional, Computed, Map, Available since v1.230.1) A list of feature attributes for the VPC attachment.
* `vpc_id` - (Required, ForceNew) The ID of the VPC instance.
* `vpc_owner_id` - (Optional, ForceNew, Computed) The ID of the Alibaba Cloud account (main account) to which the VPC instance belongs. The default value is the ID of the Alibaba Cloud account you are currently logged in to.

-> **NOTE:**  If you want to attach a cross-account network instance, you must specify this parameter.

* `zone_mappings` - (Required, List) Select a vSwitch in an availability zone supported by the Enterprise Edition transit router.
You can specify up to 10 entries at a time. See [`zone_mappings`](#zone_mappings) below.

The following arguments will be discarded. Please use new fields as soon as possible:
* `transit_router_attachment_name` - (Deprecated since v1.279.0). Field 'transit_router_attachment_name' has been deprecated from provider version 1.279.0. New field 'transit_router_vpc_attachment_name' instead.

### `zone_mappings`

The zone_mappings supports the following:
* `vswitch_id` - (Required) The ID of the vSwitch to be added to the VPC attachment.

You can add up to 10 vSwitches at a time.
  - If the Alibaba Cloud account you are currently logged in to is the same as the account to which the VPC instance belongs, you can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation to query the IDs of vSwitches in the VPC instance and the IDs of their associated zones.
  - If the Alibaba Cloud account you are currently logged in to is different from the account to which the VPC instance belongs, you can call the [ListGrantVSwitchesToCen](https://help.aliyun.com/document_detail/427599.html) operation to query the IDs of vSwitches in the VPC instance and the IDs of their associated zones.

If you specify `VSwitchId`, you must also specify `ZoneId`.
* `zone_id` - (Required) The ID of the availability zone supported by the Enterprise Edition transit router.
You can query the availability zone IDs by using the [DescribeZones](https://help.aliyun.com/document_detail/36064.html) operation.
You can select up to 10 availability zones at a time.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the VPC attachment.
* `region_id` - The region ID of the VPC instance.
* `status` - The status of the VPC attachment.
* `transit_router_attachment_id` - The resource ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Transit Router Vpc Attachment.
* `delete` - (Defaults to 20 mins) Used when delete the Transit Router Vpc Attachment.
* `update` - (Defaults to 8 mins) Used when update the Transit Router Vpc Attachment.

## Import

Cloud Enterprise Network (CEN) Transit Router Vpc Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_cen_transit_router_vpc_attachment.example <transit_router_attachment_id>
```