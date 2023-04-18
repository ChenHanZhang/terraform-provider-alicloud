---
subcategory: "EIP"
layout: "alicloud"
page_title: "Alicloud: alicloud_eip_addresss"
sidebar_current: "docs-alicloud-datasource-eip-addresss"
description: |-
  Provides a list of Eip Address owned by an Alibaba Cloud account.
---

# alicloud_eip_addresss

This data source provides Eip Address available to the user.[What is Address](https://www.alibabacloud.com/help/en/)

-> **NOTE:** Available in 1.204.0+.

## Example Usage

```
data "alicloud_eip_addresss" "default" {
  ids               = ["${alicloud_eip_address.default.id}"]
  name_regex        = alicloud_eip_address.default.name
  address_name      = "Test12"
  isp               = "BGP"
  resource_group_id = "rg-aek2xl5qajpkquq"
}

output "alicloud_eip_address_example_id" {
  value = data.alicloud_eip_addresss.default.addresss.0.id
}
```

## Argument Reference

The following arguments are supported:
* `address_name` - (ForceNew,Optional) The name of the EIP.
* `allocation_id` - (ForceNew,Optional) The ID of the EIP instance.
* `instance_id` - (ForceNew,Optional) The ID of the current bound instance.
* `instance_type` - (ForceNew,Optional) The type of the current bound instance.-**EcsInstance**: an ECS instance of the VPC type.-**SlbInstance**: an SLB instance of the VPC type.-**Nat**:NAT gateway.-**HaVip**: a highly available virtual IP address.-**NetworkInterface**: Secondary ENI.
* `ip_address` - (ForceNew,Optional) The IP address of the EIP.
* `isp` - (ForceNew,Optional) Service providers.
* `payment_type` - (ForceNew,Optional) The billing method of the EIP. Valid values:  PrePaid: subscription PostPaid: pay-as-you-go
* `resource_group_id` - (ForceNew,Optional) The ID of the resource group.
* `segment_instance_id` - (ForceNew,Optional) The ID of the consecutive EIPs.This parameter value is returned only if the EIP is a continuous EIP.
* `status` - (ForceNew,Optional) The status of the EIP.-**Associating**: Binding.-**Unassociating**: Unbinding.-**InUse**: Assigned.-**Available**: Available.
* `ids` - (Optional, ForceNew, Computed) A list of Address IDs.
* `address_names` - (Optional, ForceNew) The name of the Address. You can specify at most 10 names.
* `name_regex` - (Optional, ForceNew) A regex string to filter results by Group Metric Rule name.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).


## Attributes Reference

The following attributes are exported in addition to the arguments listed above:
* `ids` - A list of Address IDs.
* `names` - A list of name of Addresss.
* `addresss` - A list of Address Entries. Each element contains the following attributes:
  * `address_name` - The name of the EIP.
  * `allocation_id` - The ID of the EIP instance.
  * `bandwidth` - The peak bandwidth of the EIP. Unit: Mbps.
  * `bandwidth_package_bandwidth` - The bandwidth value of the Internet shared bandwidth added to the EIP. The unit is Mbps.
  * `bandwidth_package_id` - The ID of the added shared bandwidth.
  * `bandwidth_package_type` - The type of bandwidth. Only **CommonBandwidthPackage** (shared bandwidth) is supported.
  * `business_status` - The business status of the EIP instance. Value:-**Normal**: Normal.-**Financialized**: locked.
  * `create_time` - The time when the EIP was created.
  * `deletion_protection` - Whether the delete protection function is turned on.-**true**: enabled.-**false**: not enabled.
  * `description` - 描述
  * `eip_bandwidth` - AN EIP is added to an Internet shared bandwidth before or exit purchase an Internet shared bandwidth instance after the bandwidth Mbps Set for the bandwidth.
  * `expired_time` - Expiration time, format' YYYY-MM-DDThh:mm:ssZ '.
  * `has_reservation_data` - Whether there is renewal data.The value is **true** only when the input parameter **IncludeReservationData** is **true** and there is no valid subscription data * *.
  * `hd_monitor_status` - Whether the second-level monitoring is enabled for the EIP.-**false**: not enabled.-**true**: enabled.
  * `instance_id` - The ID of the current bound instance.
  * `instance_region_id` - The region ID of the currently bound resource.
  * `instance_type` - The type of the current bound instance.-**EcsInstance**: an ECS instance of the VPC type.-**SlbInstance**: an SLB instance of the VPC type.-**Nat**:NAT gateway.-**HaVip**: a highly available virtual IP address.-**NetworkInterface**: Secondary ENI.
  * `internet_charge_type` - The billing method of EIP.-**PayByBandwidth**: Billed by fixed bandwidth.-**PayByTraffic**: Billing by using traffic.
  * `ip_address` - The IP address of the EIP.
  * `isp` - Service providers.
  * `netmode` - Network type.
  * `operation_locks` - Lock details.
    * `lock_reason` - Lock type.-**financial**: locked due to arrears.-**security**: locked for security reasons.
  * `payment_type` - The billing method of the EIP. Valid values:  PrePaid: subscription PostPaid: pay-as-you-go
  * `reservation_active_time` - The effective time of the renewal fee, in the format of 'YYYY-MM-DDThh:mm:ssZ'.
  * `reservation_bandwidth` - Renewal bandwidth, in Mbps.
  * `reservation_internet_charge_type` - Renewal Payment type.-**PayByBandwidth**: billed by fixed bandwidth.-**PayByTraffic**: Billing by traffic.
  * `reservation_order_type` - Renewal order type.-**RENEWCHANGE**: Renewal change.-**TEMP_UPGRADE**: Short-term upgrade.-**UPGRADE**: UPGRADE.
  * `resource_group_id` - The ID of the resource group.
  * `second_limited` - Whether a secondary speed limit is configured.-**true**: configured.-**false**: Not configured.
  * `segment_instance_id` - The ID of the consecutive EIPs.This parameter value is returned only if the EIP is a continuous EIP.
  * `service_managed` - Indicates the resource created for the service account. Value:-**0**: Create a non-service account.-**1**: Create a service account.
  * `status` - The status of the EIP.-**Associating**: Binding.-**Unassociating**: Unbinding.-**InUse**: Assigned.-**Available**: Available.
  * `tags` - The tag of the resource
  * `vpc_id` - Private network Id
