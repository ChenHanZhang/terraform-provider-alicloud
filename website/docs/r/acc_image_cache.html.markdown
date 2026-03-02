---
subcategory: "ACC"
layout: "alicloud"
page_title: "Alicloud: alicloud_acc_image_cache"
description: |-
  Provides a Alicloud ACC Image Cache resource.
---

# alicloud_acc_image_cache

Provides a ACC Image Cache resource.



For information about ACC Image Cache and how to use it, see [What is Image Cache](https://next.api.alibabacloud.com/document/acc/2024-04-02/CreateImageCache).

-> **NOTE:** Available since v1.272.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `acr_registry_infos` - (Optional, List) ACR Registry Information See [`acr_registry_infos`](#acr_registry_infos) below.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `image_cache_name` - (Required, ForceNew) The name of the resource
* `image_registry_credentials` - (Optional, List) The image registry credentials for creating the resource. See [`image_registry_credentials`](#image_registry_credentials) below.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `images` - (Required, ForceNew, List) Images used for creating image cache.
* `network_config` - (Required, ForceNew, Set) The nework config for creating image cache. See [`network_config`](#network_config) below.
* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `tags` - (Optional, Map) Tags

### `acr_registry_infos`

The acr_registry_infos supports the following:
* `instance_id` - (Optional) ACR Repository Instance ID
* `region_id` - (Optional) ACR Repository Region

### `image_registry_credentials`

The image_registry_credentials supports the following:
* `password` - (Optional) The password of the image registry credential.
* `server` - (Optional) The server of the image registry credential.
* `skip_cert_verification` - (Optional) The operation property field that represents whether to skip certificate authentication
* `use_plain_http` - (Optional) The operation property field that represents whether to use the HTTP protocol
* `username` - (Optional) The username of the image registry credential.

### `network_config`

The network_config supports the following:
* `eip_instance` - (Optional, ForceNew, Set) The elastic IP address instance for creating image cache. See [`eip_instance`](#network_config-eip_instance) below.
* `security_group_id` - (Required, ForceNew) The security group id for creating image cache.
* `vswitch_ids` - (Required, ForceNew, List) The vSwitch ids for creating image cache.

### `network_config-eip_instance`

The network_config-eip_instance supports the following:
* `auto_create` - (Optional, ForceNew) Automatically create an elastic IP instance.
* `bandwidth` - (Optional, ForceNew, Int) The bandwidth of elastic IP address instance.
* `instance_id` - (Optional, ForceNew) The elastic IP address instance id for creating image cache.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.
* `region_id` - The region ID of the resource.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Image Cache.
* `delete` - (Defaults to 5 mins) Used when delete the Image Cache.
* `update` - (Defaults to 5 mins) Used when update the Image Cache.

## Import

ACC Image Cache can be imported using the id, e.g.

```shell
$ terraform import alicloud_acc_image_cache.example <image_cache_id>
```