---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_vpc_firewall_acl_group_list"
description: |-
  Provides a Alicloud Cloud Firewall Vpc Firewall Acl Group List resource.
---

# alicloud_cloud_firewall_vpc_firewall_acl_group_list

Provides a Cloud Firewall Vpc Firewall Acl Group List resource.

VPC Firewall Access Control Policy Group.

For information about Cloud Firewall Vpc Firewall Acl Group List and how to use it, see [What is Vpc Firewall Acl Group List](https://next.api.alibabacloud.com/document/Cloudfw/2017-12-07/DescribeVpcFirewallAclGroupList).

-> **NOTE:** Available since v1.286.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

### Deleting `alicloud_cloud_firewall_vpc_firewall_acl_group_list` or removing it from your configuration

Terraform cannot destroy resource `alicloud_cloud_firewall_vpc_firewall_acl_group_list`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 


## Import

Cloud Firewall Vpc Firewall Acl Group List can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_firewall_vpc_firewall_acl_group_list.example <vpc_firewall_acl_group_list_id>
```