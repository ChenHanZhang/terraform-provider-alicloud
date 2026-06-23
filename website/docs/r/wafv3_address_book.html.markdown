---
subcategory: "Web Application Firewall(WAF)"
layout: "alicloud"
page_title: "Alicloud: alicloud_wafv3_address_book"
description: |-
  Provides a Alicloud WAFV3 Address Book resource.
---

# alicloud_wafv3_address_book

Provides a WAFV3 Address Book resource.



For information about WAFV3 Address Book and how to use it, see [What is Address Book](https://next.api.alibabacloud.com/document/waf-openapi/2021-10-01/CreateDefenseRule).

-> **NOTE:** Available since v1.283.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_wafv3_instance" "defaultd84Zq8" {
}


resource "alicloud_wafv3_address_book" "default" {
  description       = "example"
  instance_id       = alicloud_wafv3_instance.defaultd84Zq8.id
  address_book_name = "example_from_tf_220"
  address_list      = ["100.100.100.100/32", "101.101.101.101/32", "102.102.102.102/32"]
  address_book_type = "ip"
}
```

## Argument Reference

The following arguments are supported:
* `address_book_name` - (Optional) The name of the resource
* `address_book_type` - (Required) This property does not have a description in the spec, please add it before generating code.
* `address_list` - (Optional, List) The address list to add.
* `description` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `instance_id` - (Required, ForceNew) The ID of the WAF instance.

-> **NOTE:**  You can call [DescribeInstance](~~ 433756 ~~) to view the ID of the current WAF instance.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<instance_id>:<address_book_id>`.
* `address_book_id` - The first ID of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Address Book.
* `delete` - (Defaults to 5 mins) Used when delete the Address Book.
* `update` - (Defaults to 5 mins) Used when update the Address Book.

## Import

WAFV3 Address Book can be imported using the id, e.g.

```shell
$ terraform import alicloud_wafv3_address_book.example <instance_id>:<address_book_id>
```