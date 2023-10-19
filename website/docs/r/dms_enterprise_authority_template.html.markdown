---
subcategory: "DMS Enterprise"
layout: "alicloud"
page_title: "Alicloud: alicloud_dms_enterprise_authority_template"
description: |-
  Provides a Alicloud D M S Enterprise Authority Template resource.
---

# alicloud_dms_enterprise_authority_template

Provides a D M S Enterprise Authority Template resource. 

For information about D M S Enterprise Authority Template and how to use it, see [What is Authority Template](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available since v1.212.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

variable "v_tid" {
  default = "1"
}


resource "alicloud_dms_enterprise_authority_template" "default" {
  authority_template_name = var.name

  description = "资源用例测试权限模板"
  tid         = var.v_tid
}
```

### Deleting `alicloud_dms_enterprise_authority_template` or removing it from your configuration

Terraform cannot destroy resource `alicloud_dms_enterprise_authority_template`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `authority_template_name` - (Required) Permission Template name.
* `description` - (Optional) Permission template description information.
* `tid` - (Optional, ForceNew, Computed) Tenant ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<tid>:<authority_template_id>`.
* `authority_template_id` - Permission template ID.
* `create_time` - The creation time of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Authority Template.
* `update` - (Defaults to 5 mins) Used when update the Authority Template.

## Import

D M S Enterprise Authority Template can be imported using the id, e.g.

```shell
$ terraform import alicloud_dms_enterprise_authority_template.example <tid>:<authority_template_id>
```