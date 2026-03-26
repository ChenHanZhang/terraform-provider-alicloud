---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_integration_policy"
description: |-
  Provides a Alicloud Cms Integration Policy resource.
---

# alicloud_cms_integration_policy

Provides a Cms Integration Policy resource.

Policies used by the Integration Center.

For information about Cms Integration Policy and how to use it, see [What is Integration Policy](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateIntegrationPolicy).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}


resource "alicloud_cms_integration_policy" "default" {
  policy_type = "CS"
  entity_group {
    cluster_id          = "c3486629162d14953992eb8f015df3f0f"
    cluster_entity_type = "ManagedKubernetes/Default"
  }
  integration_policy_name = "example-create-name"
  workspace               = "prometheus"
}
```

## Argument Reference

The following arguments are supported:
* `entity_group` - (Optional, ForceNew, Set) Entity Group See [`entity_group`](#entity_group) below.
* `fee_package` - (Optional) Fee Package

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `force` - (Optional) Whether to forcibly delete.

-> **NOTE:** This parameter configures deletion behavior and is only evaluated when Terraform attempts to destroy the resource. Changes to this parameter during updates are stored but have no immediate effect.

* `integration_policy_name` - (Optional) The name of the resource
* `policy_type` - (Required, ForceNew) Access Center Policy Name
* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `tags` - (Optional, Map) The tag of the resource
* `workspace` - (Required, ForceNew) Workspace

### `entity_group`

The entity_group supports the following:
* `cluster_entity_type` - (Optional, ForceNew) Cluster entity type
* `cluster_id` - (Optional, ForceNew) Cluster ID
* `entity_group_id` - (Optional, ForceNew) Group ID
* `vpc_id` - (Optional, ForceNew) VPC ID

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `region_id` - The region ID of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Integration Policy.
* `delete` - (Defaults to 5 mins) Used when delete the Integration Policy.
* `update` - (Defaults to 5 mins) Used when update the Integration Policy.

## Import

Cms Integration Policy can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_integration_policy.example <integration_policy_id>
```