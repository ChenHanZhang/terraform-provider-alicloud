---
subcategory: "OSS"
layout: "alicloud"
page_title: "Alicloud: alicloud_oss_bucket_worm"
description: |-
  Provides a Alicloud OSS Bucket Worm resource.
---

# alicloud_oss_bucket_worm

Provides a OSS Bucket Worm resource.

Bucket Retention Policy.

For information about OSS Bucket Worm and how to use it, see [What is Bucket Worm](https://www.alibabacloud.com/help/en/oss/developer-reference/initiatebucketworm).

-> **NOTE:** Available since v1.240.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_oss_bucket" "defaulthNMfIF" {
  storage_class = "Standard"
}


resource "alicloud_oss_bucket_worm" "default" {
  bucket                   = alicloud_oss_bucket.defaulthNMfIF.bucket
  retention_period_in_days = "1"
  status                   = "InProgress"
}
```

### Deleting `alicloud_oss_bucket_worm` or removing it from your configuration

The `alicloud_oss_bucket_worm` resource allows you to manage  `status = "Locked"`  instance, but Terraform cannot destroy it.
Deleting the subscription resource or removing it from your configuration will remove it from your state file and management, but will not destroy the Instance.
You can resume managing the subscription instance via the AlibabaCloud Console.

## Argument Reference

The following arguments are supported:
* `bucket` - (Required, ForceNew) The name of the bucket
* `retention_period_in_days` - (Optional, Int) The specified number of days to retain the Object.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<bucket>:<worm_id>`.
* `create_time` - The creation time of the resource
* `status` - The status of the compliance retention policy.
* `worm_id` - The ID of the retention policy.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Bucket Worm.
* `delete` - (Defaults to 5 mins) Used when delete the Bucket Worm.
* `update` - (Defaults to 5 mins) Used when update the Bucket Worm.

## Import

OSS Bucket Worm can be imported using the id, e.g.

```shell
$ terraform import alicloud_oss_bucket_worm.example <bucket>:<worm_id>
```