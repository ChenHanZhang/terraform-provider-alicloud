---
subcategory: "EIP"
layout: "alicloud"
page_title: "Alicloud: alicloud_eip_address"
sidebar_current: "docs-alicloud-resource-eip-address"
description: |-
  Provides a Alicloud Eip Address resource.
---

# alicloud_eip_address

Provides a Eip Address resource.

For information about Eip Address and how to use it, see [What is Address](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_eip_address" "default" {
  address_name      = "Test12"
  description       = "ab"
  resource_group_id = "rg-aek2xl5qajpkquq"
  isp               = "BGP"
  bandwidth         = "5"
  region_id         = "cn-hangzhou"
}
```

## Argument Reference

The following arguments are supported:
* `activity_id` - (ForceNew,Optional) Special activity ID. This parameter is not required.
* `address_name` - (Optional) The name of the EIP.
* `auto_pay` - (ForceNew,Optional) False (default): automatic payment is not enabled. After generating an order, you need to complete the payment at the order center.true: Turn on automatic payment and automatically pay for orders.This parameter is required when the value of the InstanceChargeType parameter is PrePaid. This parameter is optional when the value of the InstanceChargeType parameter is PostPaid.
* `bandwidth` - (Computed,Optional) The peak bandwidth of the EIP. Unit: Mbps.
* `bandwidth_package_id` - (ForceNew,Computed,Optional) The ID of the added shared bandwidth.
* `deletion_protection` - (Computed,Optional) Whether the delete protection function is turned on.-**true**: enabled.-**false**: not enabled.
* `description` - (Optional) 描述
* `hd_monitor_log_project` - (Optional) The name of the LogProject of the high-precision second-level monitoring log service.
* `hd_monitor_log_store` - (Optional) The name of the LogStore of the high-precision second-level monitoring log service.
* `hd_monitor_status` - (ForceNew,Computed,Optional) Whether the second-level monitoring is enabled for the EIP.-**false**: not enabled.-**true**: enabled.
* `instance_id` - (Optional) The ID of the current bound instance.
* `instance_region_id` - (Optional) The region ID of the currently bound resource.
* `instance_type` - (Optional) The type of the current bound instance.-**EcsInstance**: an ECS instance of the VPC type.-**SlbInstance**: an SLB instance of the VPC type.-**Nat**:NAT gateway.-**HaVip**: a highly available virtual IP address.-**NetworkInterface**: Secondary ENI.
* `internet_charge_type` - (ForceNew,Computed,Optional) The billing method of EIP.-**PayByBandwidth**: Billed by fixed bandwidth.-**PayByTraffic**: Billing by using traffic.
* `isp` - (ForceNew,Computed,Optional) Service providers.
* `netmode` - (ForceNew,Computed,Optional) Network type.
* `payment_type` - (ForceNew,Optional) The billing method of the EIP. Valid values:  PrePaid: subscription PostPaid: pay-as-you-go
* `period` - (ForceNew,Optional) When the PricingCycle is set to Month, the Period value ranges from 1 to 9.When the PricingCycle is set to Year, the Period range is 1 to 5.If the value of the InstanceChargeType parameter is PrePaid, this parameter is required. If the value of the InstanceChargeType parameter is PostPaid, this parameter is not filled in.
* `pricing_cycle` - (ForceNew,Optional) Value:Month (default): Pay monthly.Year: Pay per Year.This parameter is required when the value of the InstanceChargeType parameter is PrePaid. This parameter is optional when the value of the InstanceChargeType parameter is PostPaid.
* `public_ip_address_pool_id` - (ForceNew,Optional) The ID of the IP address pool to which the EIP belongs.
* `resource_group_id` - (Computed,Optional) The ID of the resource group.
* `second_limited` - (Computed,Optional) Whether a secondary speed limit is configured.-**true**: configured.-**false**: Not configured.
* `security_protection_types` - (ForceNew,Optional) Security protection level.-When the return is empty, the basic DDoS protection is specified.-When **antidos_enhanced** is returned, it indicates DDoS protection (enhanced version).
* `tags` - (Optional) The tag of the resourceSee the following `Block Tags`.
* `vpc_id` - (Optional) Private network Id

The following arguments will be discarded. Please use new fields as soon as possible:
* `name` - Field 'name' has been deprecated from provider version 1.126.0. New field 'address_name' instead.
* `instance_charge_type` - Field 'instance_charge_type' has been deprecated from provider version 1.126.0. New field 'payment_type' instead.
* `high_definition_monitor_log_status` - Field 'high_definition_monitor_log_status' has been deprecated from provider version 1.204.0. New field 'hd_monitor_status' instead.
* `log_project` - Field 'log_project' has been deprecated from provider version 1.204.0. New field 'hd_monitor_log_project' instead.
* `log_store` - Field 'log_store' has been deprecated from provider version 1.204.0. New field 'hd_monitor_log_store' instead.

#### Block Tags

The Tags supports the following:
* `tag_key` - (Optional) The key of the tags.
* `tag_value` - (Optional) The value of the tags.



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `allocation_id` - The ID of the EIP instance.
* `bandwidth` - The peak bandwidth of the EIP. Unit: Mbps.
* `bandwidth_package_bandwidth` - The bandwidth value of the Internet shared bandwidth added to the EIP. The unit is Mbps.
* `bandwidth_package_id` - The ID of the added shared bandwidth.
* `bandwidth_package_type` - The type of bandwidth. Only **CommonBandwidthPackage** (shared bandwidth) is supported.
* `business_status` - The business status of the EIP instance. Value:-**Normal**: Normal.-**Financialized**: locked.
* `create_time` - The time when the EIP was created.
* `deletion_protection` - Whether the delete protection function is turned on.-**true**: enabled.-**false**: not enabled.
* `eip_bandwidth` - AN EIP is added to an Internet shared bandwidth before or exit purchase an Internet shared bandwidth instance after the bandwidth Mbps Set for the bandwidth.
* `expired_time` - Expiration time, format' YYYY-MM-DDThh:mm:ssZ '.
* `has_reservation_data` - Whether there is renewal data.The value is **true** only when the input parameter **IncludeReservationData** is **true** and there is no valid subscription data * *.
* `hd_monitor_status` - Whether the second-level monitoring is enabled for the EIP.-**false**: not enabled.-**true**: enabled.
* `internet_charge_type` - The billing method of EIP.-**PayByBandwidth**: Billed by fixed bandwidth.-**PayByTraffic**: Billing by using traffic.
* `ip_address` - The IP address of the EIP.
* `isp` - Service providers.
* `netmode` - Network type.
* `operation_locks` - Lock details.
  * `lock_reason` - Lock type.-**financial**: locked due to arrears.-**security**: locked for security reasons.
* `reservation_active_time` - The effective time of the renewal fee, in the format of 'YYYY-MM-DDThh:mm:ssZ'.
* `reservation_bandwidth` - Renewal bandwidth, in Mbps.
* `reservation_internet_charge_type` - Renewal Payment type.-**PayByBandwidth**: billed by fixed bandwidth.-**PayByTraffic**: Billing by traffic.
* `reservation_order_type` - Renewal order type.-**RENEWCHANGE**: Renewal change.-**TEMP_UPGRADE**: Short-term upgrade.-**UPGRADE**: UPGRADE.
* `resource_group_id` - The ID of the resource group.
* `second_limited` - Whether a secondary speed limit is configured.-**true**: configured.-**false**: Not configured.
* `segment_instance_id` - The ID of the consecutive EIPs.This parameter value is returned only if the EIP is a continuous EIP.
* `service_managed` - Indicates the resource created for the service account. Value:-**0**: Create a non-service account.-**1**: Create a service account.
* `status` - The status of the EIP.-**Associating**: Binding.-**Unassociating**: Unbinding.-**InUse**: Assigned.-**Available**: Available.
* `zone` - The zone of the EIP.This parameter is returned only for whitelist users that are visible to the zone.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Address.
* `delete` - (Defaults to 5 mins) Used when delete the Address.
* `update` - (Defaults to 5 mins) Used when update the Address.

## Import

Eip Address can be imported using the id, e.g.

```shell
$ terraform import alicloud_eip_address.example 
```