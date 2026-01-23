---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_global_config"
description: |-
  Provides a Alicloud Threat Detection Global Config resource.
---

# alicloud_threat_detection_global_config

Provides a Threat Detection Global Config resource.



For information about Threat Detection Global Config and how to use it, see [What is Global Config](https://next.api.alibabacloud.com/document/cloud-siem/2024-12-12/GetGlobalConfig).

-> **NOTE:** Available since v1.269.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

### Deleting `alicloud_threat_detection_global_config` or removing it from your configuration

Terraform cannot destroy resource `alicloud_threat_detection_global_config`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `global_config_name` - (Optional) The name of the resource
* `global_config_value` - (Optional) The ID of the resource group
* `lang` - (Optional) This property does not have a description in the spec, please add it before generating code.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `role_for` - (Optional, Int) This property does not have a description in the spec, please add it before generating code.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<Alibaba Cloud Account ID>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `update` - (Defaults to 5 mins) Used when update the Global Config.

## Import

Threat Detection Global Config can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_global_config.example <Alibaba Cloud Account ID>
```