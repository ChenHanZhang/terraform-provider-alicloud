---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_incident_investigation"
description: |-
  Provides a Alicloud Threat Detection Incident Investigation resource.
---

# alicloud_threat_detection_incident_investigation

Provides a Threat Detection Incident Investigation resource.

Incident Investigation Record.

For information about Threat Detection Incident Investigation and how to use it, see [What is Incident Investigation](https://next.api.alibabacloud.com/document/cloud-siem/2024-12-12/GetIncidentInvestigation).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

### Deleting `alicloud_threat_detection_incident_investigation` or removing it from your configuration

Terraform cannot destroy resource `alicloud_threat_detection_incident_investigation`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `incident_investigation_display_id` - (Optional, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `incident_investigation_id` - (Optional, ForceNew, Computed) The first ID of the resource
* `incident_uuid` - (Optional, ForceNew) This property does not have a description in the spec, please add it before generating code.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 


## Import

Threat Detection Incident Investigation can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_incident_investigation.example <incident_investigation_id>
```