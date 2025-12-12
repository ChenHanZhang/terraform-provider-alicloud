---
subcategory: "OSS"
layout: "alicloud"
page_title: "Alicloud: alicloud_oss_bucket_cname"
description: |-
  Provides a Alicloud OSS Bucket Cname resource.
---

# alicloud_oss_bucket_cname

Provides a OSS Bucket Cname resource.

Manage user-defined domain names.

For information about OSS Bucket Cname and how to use it, see [What is Bucket Cname](https://www.alibabacloud.com/help/en/oss/developer-reference/putcname).

-> **NOTE:** Available since v1.233.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_oss_bucket" "CreateBucket" {
  bucket        = var.name
  storage_class = "Standard"
}

resource "alicloud_oss_bucket_cname_token" "defaultZaWJfG" {
  bucket = alicloud_oss_bucket.CreateBucket.bucket
  domain = "tftestacc.com"
}

resource "alicloud_alidns_record" "defaultnHqm5p" {
  status      = "ENABLE"
  line        = "default"
  rr          = "_dnsauth"
  type        = "TXT"
  domain_name = "tftestacc.com"
  priority    = "1"
  value       = alicloud_oss_bucket_cname_token.defaultZaWJfG.token
  ttl         = "600"
  lifecycle {
    ignore_changes = [
      value,
    ]
  }
}

resource "alicloud_oss_bucket_cname" "default" {
  bucket = alicloud_oss_bucket.CreateBucket.bucket
  domain = alicloud_alidns_record.defaultnHqm5p.domain_name
}
```

## Argument Reference

The following arguments are supported:
* `bucket` - (Required, ForceNew) The bucket to which the custom domain name belongs
* `certificate` - (Optional, List) The container for the certificate configuration. See [`certificate`](#certificate) below.
* `delete_certificate` - (Optional) Whether to delete the certificate.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `domain` - (Required, ForceNew) User-defined domain name
* `force` - (Optional) Whether to force overwrite certificate.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `previous_cert_id` - (Optional) The current certificate ID. If the Force value is not true, the OSS Server checks whether the value matches the current certificate ID. If the value does not match, an error is reported.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.


### `certificate`

The certificate supports the following:
* `cert_id` - (Optional, Computed) Certificate Identifier
* `certificate` - (Optional) The certificate public key.
* `private_key` - (Optional) The certificate private key.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<bucket>:<domain>`.
* `certificate` - The container for the certificate configuration.
  * `creation_date` - Certificate creation time
  * `fingerprint` - Certificate Fingerprint
  * `status` - Certificate Status
  * `type` - Certificate Type
  * `valid_end_date` - Certificate validity period end time
  * `valid_start_date` - Certificate validity period start time
* `status` - Cname status

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Bucket Cname.
* `delete` - (Defaults to 5 mins) Used when delete the Bucket Cname.
* `update` - (Defaults to 5 mins) Used when update the Bucket Cname.

## Import

OSS Bucket Cname can be imported using the id, e.g.

```shell
$ terraform import alicloud_oss_bucket_cname.example <bucket>:<domain>
```