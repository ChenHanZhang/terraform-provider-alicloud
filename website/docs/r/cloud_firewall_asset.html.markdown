---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_asset"
description: |-
  Provides a Alicloud Cloud Firewall Asset resource.
---

# alicloud_cloud_firewall_asset

Provides a Cloud Firewall Asset resource.

Public IP Address Asset.

For information about Cloud Firewall Asset and how to use it, see [What is Asset](https://next.api.alibabacloud.com/document/Cloudfw/2017-12-07/PutEnableFwSwitch).

-> **NOTE:** Available since v1.286.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_cloud_firewall_asset" "default" {
  region           = "cn-hangzhou"
  resource_type    = "EcsPublicIp"
  internet_address = "192.0.0.1"
  lang             = "zh"
}
```

## Argument Reference

The following arguments are supported:
* `internet_address` - (Optional, ForceNew, Computed) The IP address or instance ID of the asset.
* `lang` - (Optional) The language of the notification messages.
  - `zh`: Chinese.
  - `en`: English.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `region` - (Optional, ForceNew) The ID of the region where the asset resides.
* `resource_type` - (Optional, ForceNew) The list of asset types.

Valid values:
  - `BastionHostEgressIP`: Bastionhost egress IP address.
  - `BastionHostIngressIP`: Bastionhost ingress IP address.
  - `EcsEIP`: ECS EIP.
  - `EcsPublicIP`: ECS public IP address.
  - `EIP`: Elastic IP Address (EIP).
  - `EniEIP`: ENI EIP.
  - `NatEIP`: NAT EIP.
  - `SlbEIP`: SLB EIP (CLB EIP).
  - `SlbPublicIP`: SLB public IP address (CLB public IP address).
  - `NatPublicIP`: NAT public IP address.
  - `HAVIP`: High-Availability Virtual IP Address (HAVIP).
  - `NlbEIP`: NLB EIP.
  - `ApiGatewayEIP`: API Gateway public IP address.
  - `AlbEIP`: ALB EIP.
  - `AiGatewayEIP`: AI Gateway public IP address.
  - `GaEIP`: GA EIP.
  - `SwasEIP`: Simple Application Server public IP address.
  - `EcdEIP`: Wuying public IP address.
  - `BastionHostIP`: Bastionhost IP address.

-> **NOTE:**  The IpaddrList, RegionList, and ResourceTypeList parameters cannot all be empty at the same time. You must specify a value for at least one of these parameters.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `ali_uid` - The UID of the Alibaba Cloud account.
* `asset_name` - The instance name of the asset protected by Cloud Firewall.
* `bind_instance_id` - The ID of the bound asset instance.
* `bind_instance_name` - The name of the bound asset instance.
* `intranet_address` - The internal IP address of the server.
* `note` - The remarks on the asset.
* `protect_status` - The firewall status.
* `region_status` - Indicates whether the region where the asset resides supports enabling Cloud Firewall protection.
* `resource_instance_id` - The ID of the asset instance.
* `sg_status` - The status of the security group policy.
* `sg_status_time` - The last time the security group status was checked, in timestamp format.
* `sync_status` - The traffic redirection support status of the asset.
* `type` - This parameter is deprecated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Asset.
* `delete` - (Defaults to 5 mins) Used when delete the Asset.

## Import

Cloud Firewall Asset can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_firewall_asset.example <internet_address>
```