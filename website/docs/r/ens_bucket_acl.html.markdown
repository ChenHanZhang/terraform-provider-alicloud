---
subcategory: "ENS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ens_bucket_acl"
description: |-
  Provides a Alicloud ENS Bucket Acl resource.
---

# alicloud_ens_bucket_acl

Provides a ENS Bucket Acl resource.

Bucket Security Configuration.

For information about ENS Bucket Acl and how to use it, see [What is Bucket Acl](https://next.api.alibabacloud.com/document/Ens/2017-11-10/PutBucketAcl).

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
  default = "pop-autocase-acl-example"
}

variable "ens_region_id" {
  default = "cn-jiaozuo-4"
}

resource "alicloud_ens_bucket" "defaultmRFqwe" {
  comment             = "bucket资源自动化example"
  bucket_acl          = "private"
  bucket_name         = var.bucket_name
  logical_bucket_type = "sink"
  ens_region_id       = var.ens_region_id
}


resource "alicloud_ens_bucket_acl" "default" {
  bucket_acl  = "public-read-write"
  bucket_name = alicloud_ens_bucket.defaultmRFqwe.id
}
```

### Deleting `alicloud_ens_bucket_acl` or removing it from your configuration

Terraform cannot destroy resource `alicloud_ens_bucket_acl`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `bucket_acl` - (Required) Bucket read and write permissions.
* `bucket_name` - (Required, ForceNew) The Bucket name.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Bucket Acl.
* `update` - (Defaults to 5 mins) Used when update the Bucket Acl.

## Import

ENS Bucket Acl can be imported using the id, e.g.

```shell
$ terraform import alicloud_ens_bucket_acl.example <bucket_name>
```