---
subcategory: "Operation Orchestration Service (OOS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_oos_patch_baseline"
description: |-
  Provides a Alicloud Operation Orchestration Service (OOS) Patch Baseline resource.
---

# alicloud_oos_patch_baseline

Provides a Operation Orchestration Service (OOS) Patch Baseline resource.



For information about Operation Orchestration Service (OOS) Patch Baseline and how to use it, see [What is Patch Baseline](https://www.alibabacloud.com/help/en/operation-orchestration-service/latest/patch-manager-overview).

-> **NOTE:** Available since v1.146.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}


resource "alicloud_oos_patch_baseline" "default" {
  patch_baseline_name = var.name
  operation_system    = "Windows"
  approval_rules      = "{\"PatchRules\":[{\"EnableNonSecurity\":true,\"PatchFilterGroup\":[{\"Values\":[\"*\"],\"Key\":\"Product\"},{\"Values\":[\"Security\",\"Bugfix\"],\"Key\":\"Classification\"},{\"Values\":[\"Critical\",\"Important\"],\"Key\":\"Severity\"}],\"ApproveAfterDays\":7,\"ComplianceLevel\":\"Unspecified\"}]}"
}
```

## Argument Reference

The following arguments are supported:
* `approval_rules` - (Required) Accept the rules.
* `approved_patches` - (Optional, List, Available since v1.219.0) Approved Patch
* `approved_patches_enable_non_security` - (Optional, Available since v1.219.0) ApprovedPatchesEnableNonSecurity
* `description` - (Optional) Patches baseline description information.
* `operation_system` - (Required, ForceNew) Operating system type.
* `patch_baseline_name` - (Required, ForceNew) The name of the patch baseline.
* `rejected_patches` - (Optional, List, Available since v1.210.0) Reject patches
* `rejected_patches_action` - (Optional, Computed, Available since v1.210.0) Rejected patches action
* `resource_group_id` - (Optional, Computed, Available since v1.219.0) The ID of the resource group
* `sources` - (Optional, List, Available since v1.219.0) Source
* `tags` - (Optional, Map) Label

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Patch Baseline.
* `delete` - (Defaults to 5 mins) Used when delete the Patch Baseline.
* `update` - (Defaults to 5 mins) Used when update the Patch Baseline.

## Import

Operation Orchestration Service (OOS) Patch Baseline can be imported using the id, e.g.

```shell
$ terraform import alicloud_oos_patch_baseline.example <patch_baseline_name>
```