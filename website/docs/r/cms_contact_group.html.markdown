---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_contact_group"
description: |-
  Provides a Alicloud Cms Contact Group resource.
---

# alicloud_cms_contact_group

Provides a Cms Contact Group resource.

Alarm Contact Group.

For information about Cms Contact Group and how to use it, see [What is Contact Group](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateContactGroup).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}


resource "alicloud_cms_contact_group" "default" {
  contact_group_name = "AlarmContactGroupNameTest"
  contact_group_id   = "yuangenexample"
}
```

## Argument Reference

The following arguments are supported:
* `contact_group_id` - (Optional, ForceNew, Computed) The first ID of the resource
* `contact_group_name` - (Optional) The name of the resource
* `contact_ids` - (Optional, List) Contact ID

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Contact Group.
* `delete` - (Defaults to 5 mins) Used when delete the Contact Group.
* `update` - (Defaults to 5 mins) Used when update the Contact Group.

## Import

Cms Contact Group can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_contact_group.example <contact_group_id>
```