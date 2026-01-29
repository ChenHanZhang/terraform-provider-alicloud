---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_rds_gad_instance"
description: |-
  Provides a Alicloud RDS Gad Instance resource.
---

# alicloud_rds_gad_instance

Provides a RDS Gad Instance resource.

Global Multi-Live Database.

For information about RDS Gad Instance and how to use it, see [What is Gad Instance](https://next.api.alibabacloud.com/document/Rds/2014-08-15/CreateGADInstance).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_rds_gad_instance" "default" {
}
```

### Deleting `alicloud_rds_gad_instance` or removing it from your configuration

Terraform cannot destroy resource `alicloud_rds_gad_instance`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `resource_group_id` - (Optional, ForceNew, Computed) The ID of the resource group

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Gad Instance.
* `update` - (Defaults to 5 mins) Used when update the Gad Instance.

## Import

RDS Gad Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_rds_gad_instance.example <gad_instance_name>
```