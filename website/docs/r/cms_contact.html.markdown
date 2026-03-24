---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_contact"
description: |-
  Provides a Alicloud Cms Contact resource.
---

# alicloud_cms_contact

Provides a Cms Contact resource.



For information about Cms Contact and how to use it, see [What is Contact](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateContact).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_cms_contact" "default" {
  lang = "en"
}
```

## Argument Reference

The following arguments are supported:
* `contact_id` - (Optional, ForceNew, Computed) The first ID of the resource
* `contact_name` - (Optional) The name of the resource
* `email` - (Optional) Mailbox
* `im_user_ids` - (Optional, Map) DingTalk and other communication tools user id
* `lang` - (Optional) Language
* `phone` - (Optional) Telephone
* `workspace` - (Optional, ForceNew) This property does not have a description in the spec, please add it before generating code.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Contact.
* `delete` - (Defaults to 5 mins) Used when delete the Contact.
* `update` - (Defaults to 5 mins) Used when update the Contact.

## Import

Cms Contact can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_contact.example <contact_id>
```