---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_public_ip_address_pool"
description: |-
  Provides a Alicloud VPC Public Ip Address Pool resource.
---

# alicloud_vpc_public_ip_address_pool

Provides a VPC Public Ip Address Pool resource. 

For information about VPC Public Ip Address Pool and how to use it, see [What is Public Ip Address Pool](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createpublicipaddresspool).

-> **NOTE:** Available since v1.186.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

resource "alicloud_resource_manager_resource_group" "defaultRg" {
  display_name        = "tf-test-acc-publicaddresspool-880"
  resource_group_name = var.name
}

resource "alicloud_resource_manager_resource_group" "changeRg" {
  display_name        = "tf-testacc-publicaddresspool-change-933"
  resource_group_name = "${var.name}1"
}


resource "alicloud_vpc_public_ip_address_pool" "default" {
  description                 = "rdk-test"
  public_ip_address_pool_name = var.name
  isp                         = "BGP"
  resource_group_id           = alicloud_resource_manager_resource_group.defaultRg.id
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) Description.
* `isp` - (Optional, ForceNew, Computed) The Internet service provider. Valid values: `BGP`, `BGP_PRO`, `ChinaTelecom`, `ChinaUnicom`, `ChinaMobile`, `ChinaTelecom_L2`, `ChinaUnicom_L2`, `ChinaMobile_L2`, `BGP_FinanceCloud`. Default Value: `BGP`.
* `public_ip_address_pool_name` - (Optional) The name of the VPC Public IP address pool.
* `resource_group_id` - (Optional, Computed) The resource group ID of the VPC Public IP address pool.
* `tags` - (Optional, Map) The tags of PrefixList.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `create_time` - The creation time of the resource.
* `status` - The status of the VPC Public IP address pool.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Public Ip Address Pool.
* `delete` - (Defaults to 5 mins) Used when delete the Public Ip Address Pool.
* `update` - (Defaults to 5 mins) Used when update the Public Ip Address Pool.

## Import

VPC Public Ip Address Pool can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_public_ip_address_pool.example <id>
```