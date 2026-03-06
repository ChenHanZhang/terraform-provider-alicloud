---
subcategory: "ENS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ens_bucket"
description: |-
  Provides a Alicloud ENS Bucket resource.
---

# alicloud_ens_bucket

Provides a ENS Bucket resource.

Logical bucket name, which is used to aggregate and manage files.

For information about ENS Bucket and how to use it, see [What is Bucket](https://next.api.alibabacloud.com/document/Ens/2017-11-10/PutBucket).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

variable "bucket_name" {
  default = "pop-autoexample-sink-jiaozuo4-private"
}

variable "ens_region_id" {
  default = "cn-jiaozuo-4"
}


resource "alicloud_ens_bucket" "default" {
  comment             = "bucket资源自动化example"
  bucket_acl          = "private"
  bucket_name         = var.bucket_name
  logical_bucket_type = "sink"
  ens_region_id       = var.ens_region_id
}
```

## Argument Reference

The following arguments are supported:
* `bucket_acl` - (Optional, ForceNew) Bucket read/write permission type:
  - **public-read-write**: public read-write
  - **public-read**: public read
  - `private`: private (default)
* `bucket_name` - (Required, ForceNew) Logical bucket name, which is used to aggregate and manage files.
* `comment` - (Optional) Bucket Description
* `dispatch_scope` - (Optional) Scheduling range. This parameter is valid only for global scheduling buckets. Valid values:
  - `domestic`: Mainland China
  - `oversea`: Outside Mainland China

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `ens_region_id` - (Optional) ENS Region Id

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `logical_bucket_type` - (Optional, ForceNew) Single-node storage, value: sink

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The Bucket creation time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Bucket.
* `delete` - (Defaults to 5 mins) Used when delete the Bucket.
* `update` - (Defaults to 5 mins) Used when update the Bucket.

## Import

ENS Bucket can be imported using the id, e.g.

```shell
$ terraform import alicloud_ens_bucket.example <bucket_name>
```