---
subcategory: "OSS"
layout: "alicloud"
page_title: "Alicloud: alicloud_oss_bucket"
description: |-
  Provides a Alicloud OSS Bucket resource.
---

# alicloud_oss_bucket

Provides a OSS Bucket resource.



For information about OSS Bucket and how to use it, see [What is Bucket](https://next.api.alibabacloud.com/document/Oss/2019-05-17/PutBucket).

-> **NOTE:** Available since v1.2.0.

## Example Usage

Basic Usage

```terraform
resource "random_integer" "default" {
  max = 99999
  min = 10000
}

resource "alicloud_oss_bucket" "bucket-acl" {
  bucket = "example-value-${random_integer.default.result}"
}

resource "alicloud_oss_bucket_acl" "bucket-acl" {
  bucket = alicloud_oss_bucket.bucket-acl.bucket
  acl    = "private"
}
```

## Argument Reference

The following arguments are supported:
* `bucket_name` - (Required, ForceNew, Available since v1.283.0) Bucket name. The naming conventions for buckets are as follows:
Only lowercase letters, numbers, and dashes (-) can be included.
Must start and end with lowercase letters or numbers.
The length must be between 3 and 63 characters.
* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `storage_class` - (Optional, ForceNew, Computed) The storage type of the Bucket. The range of values is as follows:
Standard (default): Standard storage
IA: low frequency access
Archive: Archive storage
ColdArchive: cold archive storage

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the Bucket.
* `extranet_endpoint` - The Internet domain name of the Bucket.
* `intranet_endpoint` - The intranet domain name of the ECS instance that accesses the Bucket in the same region.
* `location` - The region where the Bucket is located.
* `owner` - Container for storing Bucket owner information.
  * `display_name` - The name of the Bucket owner (currently the same as the user ID).
  * `id` - The user ID of the Bucket owner.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Bucket.
* `delete` - (Defaults to 5 mins) Used when delete the Bucket.
* `update` - (Defaults to 5 mins) Used when update the Bucket.

## Import

OSS Bucket can be imported using the id, e.g.

```shell
$ terraform import alicloud_oss_bucket.example <bucket_name>
```