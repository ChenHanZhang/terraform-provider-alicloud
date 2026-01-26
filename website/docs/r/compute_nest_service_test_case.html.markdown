---
subcategory: "Compute Nest"
layout: "alicloud"
page_title: "Alicloud: alicloud_compute_nest_service_test_case"
description: |-
  Provides a Alicloud Compute Nest Service Test Case resource.
---

# alicloud_compute_nest_service_test_case

Provides a Compute Nest Service Test Case resource.

Service Test Case.

For information about Compute Nest Service Test Case and how to use it, see [What is Service Test Case](https://next.api.alibabacloud.com/document/ComputeNestSupplier/2021-05-21/CreateServiceTestCase).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_compute_nest_service_example_case" "default" {
        example_config = "{ parameters: { SystemDiskSize: 100, PayType: \'PostPaid\',  DataDiskSize: 40, InstanceType: \'$[iact3-auto]\',  AllocatePublicIp: \'true\',  DataDiskCategory: \'cloud_efficiency\',InstancePassword: \'$[iact3-auto]\',SystemDiskCategory: \'cloud_essd\'} }"
          template_name = "模版2"
          service_version = "1"
          example_case_name = "caseexample"
          service_id = "service-b7f967cacfa04d699a8f"
  }
```

## Argument Reference

The following arguments are supported:
* `service_id` - (Required) Service ID

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `service_version` - (Required) Service Version

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `template_name` - (Required, ForceNew) Template Name
* `test_case_name` - (Required) Test Case Name
* `test_config` - (Required) Test Configuration

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Service Test Case.
* `delete` - (Defaults to 5 mins) Used when delete the Service Test Case.
* `update` - (Defaults to 5 mins) Used when update the Service Test Case.

## Import

Compute Nest Service Test Case can be imported using the id, e.g.

```shell
$ terraform import alicloud_compute_nest_service_test_case.example <test_case_id>
```