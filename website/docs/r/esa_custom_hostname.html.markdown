---
subcategory: "ESA"
layout: "alicloud"
page_title: "Alicloud: alicloud_esa_custom_hostname"
description: |-
  Provides a Alicloud ESA Custom Hostname resource.
---

# alicloud_esa_custom_hostname

Provides a ESA Custom Hostname resource.



For information about ESA Custom Hostname and how to use it, see [What is Custom Hostname](https://next.api.alibabacloud.com/document/ESA/2024-09-10/CreateCustomHostname).

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_example_CustomHostname" {
  type         = "NS"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "overseas"
  auto_pay     = true
  plan_name    = "high"
}

resource "alicloud_esa_site" "resource_Site_example_CustomHostname" {
  site_name   = "gositecdn.cn"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_example_CustomHostname.id
  coverage    = "overseas"
  access_type = "NS"
}

resource "alicloud_esa_record" "resource_Record_example_CustomHostname" {
  record_name = "www.gositecdn.cn"
  comment     = "This is a remark"
  proxied     = true
  site_id     = alicloud_esa_site.resource_Site_example_CustomHostname.id
  record_type = "CNAME"
  source_type = "S3"
  data {
    value = "www.idlexampler.com"
  }
  biz_name    = "api"
  host_policy = "follow_hostname"
  ttl         = "100"
  auth_conf {
    secret_key = "hijklmnhijklmnhijklmnhijklmn"
    version    = "v4"
    region     = "us-east-1"
    auth_type  = "private"
    access_key = "abcdefgabcdefgabcdefgabcdefg"
  }
}


resource "alicloud_esa_custom_hostname" "default" {
  site_id   = alicloud_esa_site.resource_Site_example_CustomHostname.id
  hostname  = "unitexample.ialicdn.com"
  cert_type = "free"
  record_id = alicloud_esa_record.resource_Record_example_CustomHostname.id
  ssl_flag  = "on"
  cas_id    = "0"
}
```

## Argument Reference

The following arguments are supported:
* `cas_id` - (Optional, Int) The ID of the Alibaba Cloud security certificate.
* `cas_region` - (Optional) The region where the Apsara Stack Security certificate is located. It is required when using an Apsara Stack Security certificate.

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `cert_type` - (Required) The certificate type.
* `certificate` - (Optional) The public key of the uploaded certificate.
* `hostname` - (Required, ForceNew) User-defined host name.
* `private_key` - (Optional) Upload the private key of the certificate.
* `record_id` - (Required, Int) The ID of the bound Origin record.
* `site_id` - (Required, ForceNew, Int) The website ID.
* `ssl_flag` - (Required) The state of the SSL switch.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.
* `status` - Custom hostname status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Custom Hostname.
* `delete` - (Defaults to 5 mins) Used when delete the Custom Hostname.
* `update` - (Defaults to 5 mins) Used when update the Custom Hostname.

## Import

ESA Custom Hostname can be imported using the id, e.g.

```shell
$ terraform import alicloud_esa_custom_hostname.example <hostname_id>
```