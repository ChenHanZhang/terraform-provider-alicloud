---
subcategory: "Alidns"
layout: "alicloud"
page_title: "Alicloud: alicloud_alidns_cloud_gtm_monitor_template"
description: |-
  Provides a Alicloud Alidns Cloud Gtm Monitor Template resource.
---

# alicloud_alidns_cloud_gtm_monitor_template

Provides a Alidns Cloud Gtm Monitor Template resource.

CloudGtm Monitor Template.

For information about Alidns Cloud Gtm Monitor Template and how to use it, see [What is Cloud Gtm Monitor Template](https://next.api.alibabacloud.com/document/Alidns/2015-01-09/CreateCloudGtmMonitorTemplate).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `evaluation_count` - (Required) This property does not have a description in the spec, please add it before generating code.
* `extend_info` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `failure_rate` - (Required, Int) This property does not have a description in the spec, please add it before generating code.
* `interval` - (Required) This property does not have a description in the spec, please add it before generating code.
* `ip_version` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `name` - (Required) The name of the resource
* `protocol` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `remark` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `timeout` - (Required) This property does not have a description in the spec, please add it before generating code.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Cloud Gtm Monitor Template.
* `delete` - (Defaults to 5 mins) Used when delete the Cloud Gtm Monitor Template.
* `update` - (Defaults to 5 mins) Used when update the Cloud Gtm Monitor Template.

## Import

Alidns Cloud Gtm Monitor Template can be imported using the id, e.g.

```shell
$ terraform import alicloud_alidns_cloud_gtm_monitor_template.example <template_id>
```