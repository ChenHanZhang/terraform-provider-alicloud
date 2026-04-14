---
subcategory: "Alidns"
layout: "alicloud"
page_title: "Alicloud: alicloud_alidns_cloud_gtm_address"
description: |-
  Provides a Alicloud Alidns Cloud Gtm Address resource.
---

# alicloud_alidns_cloud_gtm_address

Provides a Alidns Cloud Gtm Address resource.



For information about Alidns Cloud Gtm Address and how to use it, see [What is Cloud Gtm Address](https://next.api.alibabacloud.com/document/Alidns/2015-01-09/CreateCloudGtmAddress).

-> **NOTE:** Available since v1.276.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

resource "alicloud_alidns_cloud_gtm_monitor_template" "create_tcp_template" {
  ip_version = "IPv4"
  timeout    = "3000"
  isp_city_nodes {
    city_code = "357"
    isp_code  = "465"
  }
  isp_city_nodes {
    city_code = "738"
    isp_code  = "465"
  }
  evaluation_count = "1"
  protocol         = "tcp"
  failure_rate     = "50"
  extend_info      = "{}"
  name             = "example-case-2"
  interval         = "60"
}

resource "alicloud_alidns_cloud_gtm_monitor_template" "create_https_template" {
  ip_version = "IPv4"
  timeout    = "2000"
  isp_city_nodes {
    city_code = "357"
    isp_code  = "465"
  }
  isp_city_nodes {
    city_code = "738"
    isp_code  = "465"
  }
  evaluation_count = "1"
  protocol         = "https"
  failure_rate     = "50"
  extend_info      = "{\"code\":400,\"followRedirect\":true,\"path\":\"/\",\"sni\":false}"
  name             = "example-case-3"
  interval         = "60"
}

resource "alicloud_alidns_cloud_gtm_monitor_template" "create_ping_template" {
  ip_version = "IPv4"
  timeout    = "3000"
  isp_city_nodes {
    city_code = "357"
    isp_code  = "465"
  }
  isp_city_nodes {
    city_code = "738"
    isp_code  = "465"
  }
  evaluation_count = "1"
  protocol         = "ping"
  failure_rate     = "50"
  extend_info      = "{\"packetLossRate\":10,\"packetNum\":20}"
  name             = "example-case-1"
  interval         = "60"
}


resource "alicloud_alidns_cloud_gtm_address" "default" {
  type             = "IPv4"
  health_judgement = "all_ok"
  health_tasks {
    template_id = alicloud_alidns_cloud_gtm_monitor_template.create_ping_template.id
  }
  health_tasks {
    port        = "53"
    template_id = alicloud_alidns_cloud_gtm_monitor_template.create_tcp_template.id
  }
  health_tasks {
    port        = "443"
    template_id = alicloud_alidns_cloud_gtm_monitor_template.create_https_template.id
  }
  address                 = "1.1.1.1"
  enable_status           = "enable"
  available_mode          = "manual"
  manual_available_status = "available"
  name                    = "addr-example-1"
  remark                  = "remark"
}
```

## Argument Reference

The following arguments are supported:
* `address` - (Required) This property does not have a description in the spec, please add it before generating code.
* `available_mode` - (Required) This property does not have a description in the spec, please add it before generating code.
* `enable_status` - (Required) This property does not have a description in the spec, please add it before generating code.
* `health_judgement` - (Required) This property does not have a description in the spec, please add it before generating code.
* `health_tasks` - (Optional, List) This property does not have a description in the spec, please add it before generating code. See [`health_tasks`](#health_tasks) below.
* `manual_available_status` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `name` - (Required) The name of the resource
* `remark` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `type` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.

### `health_tasks`

The health_tasks supports the following:
* `port` - (Optional, Int) This property does not have a description in the spec, please add it before generating code.
* `template_id` - (Optional) This property does not have a description in the spec, please add it before generating code.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Cloud Gtm Address.
* `delete` - (Defaults to 5 mins) Used when delete the Cloud Gtm Address.
* `update` - (Defaults to 5 mins) Used when update the Cloud Gtm Address.

## Import

Alidns Cloud Gtm Address can be imported using the id, e.g.

```shell
$ terraform import alicloud_alidns_cloud_gtm_address.example <address_id>
```