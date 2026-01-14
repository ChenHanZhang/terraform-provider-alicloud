---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_client_file_protect"
description: |-
  Provides a Alicloud Threat Detection Client File Protect resource.
---

# alicloud_threat_detection_client_file_protect

Provides a Threat Detection Client File Protect resource.

Client core file protection event monitoring, including file reading and writing, deletion, and permission change.

For information about Threat Detection Client File Protect and how to use it, see [What is Client File Protect](https://www.alibabacloud.com/help/zh/security-center/developer-reference/api-sas-2018-12-03-createfileprotectrule).

-> **NOTE:** Available since v1.212.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_threat_detection_client_file_protect" "default" {
  status      = "0"
  file_paths  = ["/usr/local"]
  file_ops    = ["CREATE"]
  rule_action = "pass"
  proc_paths  = ["/usr/local"]
  alert_level = "0"
  switch_id   = "FILE_PROTECT_RULE_SWITCH_TYPE_1693474122929"
  rule_name   = "rule_example"
}
```

## Argument Reference

The following arguments are supported:
* `alert_level` - (Optional, Int) The severity of alerts. Valid values:

  - 0: does not generate alerts
  - 1: sends notifications
  - 2: suspicious
  - 3: high-risk
* `file_ops` - (Required, List) The operations that you want to perform on the files.
* `file_paths` - (Required, List) The paths to the monitored files. Wildcard characters are supported.
* `platform` - (Optional, ForceNew, Available since v1.269.0) The type of the operating system. Valid values:

  - `windows`: Windows
  - `linux`: Linux
* `proc_paths` - (Required, List) The paths to the monitored processes.
* `rule_action` - (Required) The handling method of the rule. Valid values:

  - pass: allow
  - alert
* `rule_name` - (Required) The name of the rule.
* `status` - (Optional, Computed, Int) The status of the rule. Valid values:

  - `0`: disabled
  - `1`: enabled
* `switch_id` - (Optional, ForceNew) The switch ID of the rule.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Client File Protect.
* `delete` - (Defaults to 5 mins) Used when delete the Client File Protect.
* `update` - (Defaults to 5 mins) Used when update the Client File Protect.

## Import

Threat Detection Client File Protect can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_client_file_protect.example <rule_id>
```