---
subcategory: "DMS Enterprise"
layout: "alicloud"
page_title: "Alicloud: alicloud_dms_enterprise_dify_instance"
description: |-
  Provides a Alicloud DMS Enterprise Dify Instance resource.
---

# alicloud_dms_enterprise_dify_instance

Provides a DMS Enterprise Dify Instance resource.

DMS Dify instance.

For information about DMS Enterprise Dify Instance and how to use it, see [What is Dify Instance](https://next.api.alibabacloud.com/document/dms-enterprise/2018-11-01/CreateDifyInstance).

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `dify_instance_name` - (Optional) A resource attribute field that represents the resource name.
* `payment_type` - (Optional, ForceNew, Computed) A resource attribute field that represents the payment type.
* `tags` - (Optional, Map) Tags
* `workspace_id` - (Optional, ForceNew) Workspace ID

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `status` - A resource attribute field that represents the resource status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Dify Instance.
* `delete` - (Defaults to 5 mins) Used when delete the Dify Instance.
* `update` - (Defaults to 5 mins) Used when update the Dify Instance.

## Import

DMS Enterprise Dify Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_dms_enterprise_dify_instance.example <dify_instance_id>
```