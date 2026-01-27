---
subcategory: "Container Registry (CR)"
layout: "alicloud"
page_title: "Alicloud: alicloud_cr_artifact_lifecycle_rule"
description: |-
  Provides a Alicloud CR Artifact Lifecycle Rule resource.
---

# alicloud_cr_artifact_lifecycle_rule

Provides a CR Artifact Lifecycle Rule resource.

Retention policies for versions in the warehouse.

For information about CR Artifact Lifecycle Rule and how to use it, see [What is Artifact Lifecycle Rule](https://next.api.alibabacloud.com/document/cr/2018-12-01/CreateArtifactLifecycleRule).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-shenzhen"
}

resource "alicloud_cr_ee_instance" "defaultnKIyBE" {
  instance_name  = "example-artifact-life-rule-706"
  renewal_status = "ManualRenewal"
  image_scanner  = "ACR"
  period         = "1"
  instance_type  = "Basic"
}


resource "alicloud_cr_artifact_lifecycle_rule" "default" {
  auto                = false
  retention_tag_count = "30"
  scope               = "INSTANCE"
  instance_id         = alicloud_cr_ee_instance.defaultnKIyBE.id
  tag_regexp          = " "
}
```

## Argument Reference

The following arguments are supported:
* `auto` - (Required) Whether to execute automatically
* `instance_id` - (Required, ForceNew) Instance ID
* `namespace_name` - (Optional) Namespace name
* `repo_name` - (Optional) Repository Name
* `retention_tag_count` - (Optional, Int) Number of Retention Tags
* `schedule_time` - (Optional) Execution cycle
* `scope` - (Optional) Scope of cleaning
* `tag_regexp` - (Optional) Retain regular expressions for mirrored versions

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<instance_id>:<artifact_lifecycle_rule_id>`.
* `artifact_lifecycle_rule_id` - The first ID of the resource.
* `create_time` - Creation time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Artifact Lifecycle Rule.
* `delete` - (Defaults to 5 mins) Used when delete the Artifact Lifecycle Rule.
* `update` - (Defaults to 5 mins) Used when update the Artifact Lifecycle Rule.

## Import

CR Artifact Lifecycle Rule can be imported using the id, e.g.

```shell
$ terraform import alicloud_cr_artifact_lifecycle_rule.example <instance_id>:<artifact_lifecycle_rule_id>
```