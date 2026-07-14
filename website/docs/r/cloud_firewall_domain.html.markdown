---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_domain"
description: |-
  Provides a Alicloud Cloud Firewall Domain resource.
---

# alicloud_cloud_firewall_domain

Provides a Cloud Firewall Domain resource.



For information about Cloud Firewall Domain and how to use it, see [What is Domain](https://next.api.alibabacloud.com/document/Domain/2018-01-29/SaveSingleTaskForCreatingOrderActivate).

-> **NOTE:** Available since v1.286.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

### Deleting `alicloud_cloud_firewall_domain` or removing it from your configuration

Terraform cannot destroy resource `alicloud_cloud_firewall_domain`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `domain` - (Required, ForceNew) The domain name to be registered.

-> **NOTE:**  You must specify the domain name registrant information when you register a domain name. If you do not specify the domain name registrant information, the domain name registration fails. You can use the RegistrantProfileId parameter to specify an information template that defines the domain name registrant information.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `ip_addrs` - The domain name resolution results.
* `update_time` - The timestamp of the resolution.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Domain.

## Import

Cloud Firewall Domain can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_firewall_domain.example <domain>
```