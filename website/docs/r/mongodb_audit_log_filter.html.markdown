---
subcategory: "MongoDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_mongodb_audit_log_filter"
description: |-
  Provides a Alicloud Mongodb Audit Log Filter resource.
---

# alicloud_mongodb_audit_log_filter

Provides a Mongodb Audit Log Filter resource.



For information about Mongodb Audit Log Filter and how to use it, see [What is Audit Log Filter](https://next.api.alibabacloud.com/document/Dds/2015-12-01/ModifyAuditLogFilter).

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

### Deleting `alicloud_mongodb_audit_log_filter` or removing it from your configuration

Terraform cannot destroy resource `alicloud_mongodb_audit_log_filter`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `db_instance_id` - (Required, ForceNew) Database Instance Id
* `filter` - (Required, ForceNew) Set the Audit Log Collection Type
* `role_type` - (Optional) The Role of the Node in the Instance

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Audit Log Filter.
* `update` - (Defaults to 5 mins) Used when update the Audit Log Filter.

## Import

Mongodb Audit Log Filter can be imported using the id, e.g.

```shell
$ terraform import alicloud_mongodb_audit_log_filter.example <db_instance_id>
```