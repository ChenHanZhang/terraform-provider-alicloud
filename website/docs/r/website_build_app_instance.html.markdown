---
subcategory: "Website Build"
layout: "alicloud"
page_title: "Alicloud: alicloud_website_build_app_instance"
description: |-
  Provides a Alicloud Website Build App Instance resource.
---

# alicloud_website_build_app_instance

Provides a Website Build App Instance resource.



For information about Website Build App Instance and how to use it, see [What is App Instance](https://next.api.alibabacloud.com/document/WebsiteBuild/2025-04-29/CreateAppInstance).

-> **NOTE:** Available since v1.269.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `auto_renew` - (Optional) Whether auto-renewal is enabled.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `duration` - (Required, Int) Number of subscription periods.  

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `extend` - (Optional) Extended parameters.  

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `payment_type` - (Optional, Computed) A resource property field that represents the payment type.

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `pricing_cycle` - (Optional) Billing cycle.

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `profile` - (Optional, Set) Website configuration information. See [`profile`](#profile) below.
* `quantity` - (Optional, Int) Quantity  

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `resource_group_id` - (Optional, ForceNew, Computed) The ID of the resource group
* `tags` - (Optional, Map) The tag of the resource

### `profile`

The profile supports the following:
* `application_type` - (Optional) Application type  
* `deploy_area` - (Optional) Deployment region.  
* `site_version` - (Optional) Version.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Resource property field representing the creation time.  
* `profile` - Website configuration information.
  * `instance_id` - Alibaba Cloud instance ID  
  * `order_id` - Order ID.  
  * `template_etag` - Template identifier.  
  * `template_id` - Model template ID.  
* `status` - Instance running status.  

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the App Instance.
* `delete` - (Defaults to 5 mins) Used when delete the App Instance.
* `update` - (Defaults to 5 mins) Used when update the App Instance.

## Import

Website Build App Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_website_build_app_instance.example <biz_id>
```