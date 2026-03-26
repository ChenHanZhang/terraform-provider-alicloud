---
subcategory: "CMN"
layout: "alicloud"
page_title: "Alicloud: alicloud_cmn_ipa_app"
description: |-
  Provides a Alicloud CMN Ipa App resource.
---

# alicloud_cmn_ipa_app

Provides a CMN Ipa App resource.

Has an independent secret key, with a one-to-one correspondence to probe configurations.

For information about CMN Ipa App and how to use it, see [What is Ipa App](https://next.api.alibabacloud.com/document/cmn-ipa/2022-05-26/AddApp).

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


resource "alicloud_cmn_ipa_app" "default" {
  status  = "1"
  channel = "cms"
  tag {
    key   = "ke111"
    value = "exampleTag"
  }
  app_name = "tagTest"
}
```

## Argument Reference

The following arguments are supported:
* `app_name` - (Required) A resource property field representing the resource name.
* `channel` - (Required) Channel.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `status` - (Optional, Computed, Int) Application status. A value of 1 indicates enabled, and 0 indicates disabled.
* `tag` - (Optional, List) Tag information See [`tag`](#tag) below.

### `tag`

The tag supports the following:
* `key` - (Required) Tag key
* `value` - (Required) Tag value

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Ipa App.
* `delete` - (Defaults to 5 mins) Used when delete the Ipa App.
* `update` - (Defaults to 5 mins) Used when update the Ipa App.

## Import

CMN Ipa App can be imported using the id, e.g.

```shell
$ terraform import alicloud_cmn_ipa_app.example <ipa_app_id>
```