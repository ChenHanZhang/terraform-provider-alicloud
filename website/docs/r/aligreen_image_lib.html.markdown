---
subcategory: "Aligreen"
layout: "alicloud"
page_title: "Alicloud: alicloud_aligreen_image_lib"
description: |-
  Provides a Alicloud Aligreen Image Lib resource.
---

# alicloud_aligreen_image_lib

Provides a Aligreen Image Lib resource.

Image libraries used for image moderation.

For information about Aligreen Image Lib and how to use it, see [What is Image Lib](https://next.api.alibabacloud.com/document/Green/2017-08-23/CreateImageLib).

-> **NOTE:** Available since v1.228.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_aligreen_biz_type" "defaultUalunB" {
  biz_type_name = var.name
}


resource "alicloud_aligreen_image_lib" "default" {
  category       = "BLACK"
  enable         = true
  scene          = "PORN"
  image_lib_name = var.name
  biz_types      = [alicloud_aligreen_biz_type.defaultUalunB.biz_type_name]
}
```

## Argument Reference

The following arguments are supported:
* `biz_types` - (Optional, List) Business scenarios.
* `category` - (Required, ForceNew) The purpose of the image library. Valid values: - `black`: blacklist - `white`: whitelist.
* `enable` - (Optional, Computed) Specifies whether to enable the image library. Valid values:
  - `true`: enabled
  - `false`: disabled.
* `image_lib_name` - (Required) The name of the image library.
* `scene` - (Required, ForceNew) The usage scenario of the image library. Valid values:
  - `PORN`: pornography detection
  - `AD`: ad recognition
  - `ILLEGAL`: violence and terrorism detection.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Image Lib.
* `delete` - (Defaults to 5 mins) Used when delete the Image Lib.
* `update` - (Defaults to 5 mins) Used when update the Image Lib.

## Import

Aligreen Image Lib can be imported using the id, e.g.

```shell
$ terraform import alicloud_aligreen_image_lib.example <image_lib_id>
```