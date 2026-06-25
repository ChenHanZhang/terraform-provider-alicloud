---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_service_linked_role"
description: |-
  Provides a Alicloud Threat Detection Service Linked Role resource.
---

# alicloud_threat_detection_service_linked_role

Provides a Threat Detection Service Linked Role resource.

Service Linked Role.

For information about Threat Detection Service Linked Role and how to use it, see [What is Service Linked Role](https://www.alibabacloud.com/help/en/doc-detail/42302.htm).

-> **NOTE:** Available since v1.142.0.

## Example Usage

Basic Usage

```terraform
resource "alicloud_security_center_service_linked_role" "service_linked_role" {
}

```

## Argument Reference

The following arguments are supported:
* `service_linked_role` - (Optional, ForceNew, Computed, Available since v1.283.0) This property does not have a description in the spec, please add it before generating code.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `role_status` - The creation time of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Service Linked Role.
* `delete` - (Defaults to 5 mins) Used when delete the Service Linked Role.

## Import

Threat Detection Service Linked Role can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_service_linked_role.example <service_linked_role>
```