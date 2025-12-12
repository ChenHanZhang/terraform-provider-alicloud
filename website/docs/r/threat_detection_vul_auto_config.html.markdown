---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_vul_auto_config"
description: |-
  Provides a Alicloud Threat Detection Vul Auto Config resource.
---

# alicloud_threat_detection_vul_auto_config

Provides a Threat Detection Vul Auto Config resource.

Automatic vulnerability repair configuration.

For information about Threat Detection Vul Auto Config and how to use it, see [What is Vul Auto Config](https://next.api.alibabacloud.com/document/Sas/2018-12-03/AddOrUpdateAutoFixConfig).

-> **NOTE:** Available since v1.266.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_threat_detection_vul_auto_config" "default" {
  all_uuid          = "1"
  start_time        = "1796903984000"
  necessity         = "asap"
  target_end_time   = "6"
  type              = "cron"
  enable            = "0"
  target_start_time = "0"
  period_unit       = "day"
  need_snapshot     = "0"
  rules             = "[{\"Type\":\"cve\",\"Name\":\"alilinux3:3:ALINUX3-SA-2025:0163\"}]"
  snapshot_name     = "Test 1"
  snapshot_time     = "1"
}
```

### Deleting `alicloud_threat_detection_vul_auto_config` or removing it from your configuration

Terraform cannot destroy resource `alicloud_threat_detection_vul_auto_config`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `all_uuid` - (Required, Int) This property does not have a description in the spec, please add it before generating code.
* `config_id` - (Optional, ForceNew, Computed) This property does not have a description in the spec, please add it before generating code.
* `enable` - (Required, Int) This property does not have a description in the spec, please add it before generating code.
* `necessity` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `need_snapshot` - (Required, Int) This property does not have a description in the spec, please add it before generating code.
* `period_unit` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `rules` - (Optional) This property does not have a description in the spec, please add it before generating code.

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `snapshot_name` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `snapshot_time` - (Optional, Int) This property does not have a description in the spec, please add it before generating code.
* `start_time` - (Required, Int) This property does not have a description in the spec, please add it before generating code.
* `target_end_time` - (Optional, Int) This property does not have a description in the spec, please add it before generating code.
* `target_start_time` - (Optional, Int) This property does not have a description in the spec, please add it before generating code.
* `type` - (Required) This property does not have a description in the spec, please add it before generating code.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Vul Auto Config.
* `update` - (Defaults to 5 mins) Used when update the Vul Auto Config.

## Import

Threat Detection Vul Auto Config can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_vul_auto_config.example <id>
```