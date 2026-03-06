---
subcategory: "ENS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ens_bucket_lifecycle"
description: |-
  Provides a Alicloud ENS Bucket Lifecycle resource.
---

# alicloud_ens_bucket_lifecycle

Provides a ENS Bucket Lifecycle resource.

Bucket Lifecycle Management Resources.

For information about ENS Bucket Lifecycle and how to use it, see [What is Bucket Lifecycle](https://next.api.alibabacloud.com/document/Ens/2017-11-10/PutBucketLifecycle).

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
  default = "pop-auto-lifecycle"
}

variable "rule_id" {
  default = "pop-auto-rule-id"
}

variable "expiration_days_modify" {
  default = <<EOF
1
EOF
}

variable "prefix" {
  default = "/popauto"
}

variable "ens_region_id" {
  default = "cn-jiaozuo-4"
}

variable "expiration_days" {
  default = <<EOF
3
EOF
}

variable "expir_date" {
  default = "2024-08-15T00:00:00.000Z"
}

resource "alicloud_ens_bucket" "BUCKETNAME" {
  logical_bucket_type = "sink"
  bucket_acl          = "private"
  bucket_name         = var.bucket_name
  comment             = "bucketlifecycle资源example"
  ens_region_id       = var.ens_region_id
}


resource "alicloud_ens_bucket_lifecycle" "default" {
  bucket_name               = alicloud_ens_bucket.BUCKETNAME.id
  prefix                    = "/example"
  expiration_days           = var.expiration_days
  status                    = "Enabled"
  allow_same_action_overlap = false
}
```

## Argument Reference

The following arguments are supported:
* `allow_same_action_overlap` - (Optional) Whether prefix overlap is allowed. Value:
  - `true`: allows duplicates.
  - `false` (default): Duplicate is not allowed.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `bucket_name` - (Required, ForceNew) The Bucket name.
* `created_before_date` - (Optional) Specify the expiration period, the storage will execute the life cycle rule for data that is last updated before the expiration period.

The second time settings must follow the ISO8601 standard and use UTC time. The format is yyyy-MM-ddTHH mm:ssZ.

-> **NOTE:**  ExpirationDays and CreateBeforeDate are mutually exclusive and must be set.

* `expiration_days` - (Optional, Int) Specifies the number of days after the last update of the Object. The value must be greater than 0 and is a positive integer.

-> **NOTE:**  ExpirationDays and CreateBeforeDate are mutually exclusive and must be set.

* `prefix` - (Optional) Specifies the Prefix that the rule applies to, and does not hold the same Prefix.
  - If Prefix is specified, the rule applies only to objects in the Bucket that match the Prefix.
  - If Prefix is set to null, the rule applies to all objects in the Bucket.
* `rule_id` - (Optional, ForceNew, Computed) Identifies the unique ID of the rule. Consists of up to 255 bytes.
  - You do not need to configure this parameter when creating a new rule. The system automatically generates a unique ID.
  - To update an existing rule, you must specify a RuleId, and the rule corresponding to the RuleId must exist. Otherwise, an error is reported.

* `status` - (Required) Rule status. Value range:
  - `Enabled`: Periodically execute the rule.
  - `Disabled`: The rule is ignored.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<bucket_name>:<rule_id>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Bucket Lifecycle.
* `delete` - (Defaults to 5 mins) Used when delete the Bucket Lifecycle.
* `update` - (Defaults to 5 mins) Used when update the Bucket Lifecycle.

## Import

ENS Bucket Lifecycle can be imported using the id, e.g.

```shell
$ terraform import alicloud_ens_bucket_lifecycle.example <bucket_name>:<rule_id>
```