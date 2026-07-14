---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_vpc_firewall_control_policy"
description: |-
  Provides a Alicloud Cloud Firewall Vpc Firewall Control Policy resource.
---

# alicloud_cloud_firewall_vpc_firewall_control_policy

Provides a Cloud Firewall Vpc Firewall Control Policy resource.

VPC Firewall Control Policy.

For information about Cloud Firewall Vpc Firewall Control Policy and how to use it, see [What is Vpc Firewall Control Policy](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallcontrolpolicy).

-> **NOTE:** Available since v1.194.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

data "alicloud_account" "default" {
}

resource "alicloud_cen_instance" "default" {
  cen_instance_name = var.name
  description       = "example_value"
  tags = {
    Created = "TF"
    For     = "acceptance test"
  }
}

resource "alicloud_cloud_firewall_vpc_firewall_control_policy" "default" {
  order            = "1"
  destination      = "127.0.0.2/32"
  application_name = "ANY"
  description      = "example_value"
  source_type      = "net"
  dest_port        = "80/88"
  acl_action       = "accept"
  lang             = "zh"
  destination_type = "net"
  source           = "127.0.0.1/32"
  dest_port_type   = "port"
  proto            = "TCP"
  release          = true
  member_uid       = data.alicloud_account.default.id
  vpc_firewall_id  = alicloud_cen_instance.default.id
}
```

## Argument Reference

The following arguments are supported:
* `acl_action` - (Required) The action that Cloud Firewall performs on traffic based on the access control policy configured for the VPC firewall. Valid values:
  - `accept`: allows the traffic.
  - `drop`: blocks the traffic.
  - `log`: monitors the traffic.

-> **NOTE:**  If this parameter is not specified, all actions are queried.

* `application_name` - (Optional) The application type supported by the access control policy of the VPC firewall. Valid values:
  - `FTP`
  - `HTTP`
  - `HTTPS`
  - `MySQL`
  - `SMTP`
  - `SMTPS`
  - `RDP`
  - `VNC`
  - `SSH`
  - `Redis`
  - `MQTT`
  - `MongoDB`
  - `Memcache`
  - `SSL`
  - `ANY` (indicates all application types).
* `application_name_list` - (Optional, List, Available since v1.267.0) The list of application types supported by the access control policy.
* `description` - (Required) The description of the VPC firewall access control policy. Fuzzy search is supported.
* `dest_port` - (Optional, Computed) The destination port of the traffic in the VPC firewall access control policy.
* `dest_port_group` - (Optional) The name of the destination port address book for traffic access in the VPC firewall access control policy.

-> **NOTE:**  Set this parameter when `DestPortType` is set to `group`.

* `dest_port_type` - (Optional, Computed) The type of the destination port for traffic in a VPC firewall access control policy. Valid values:
  - `port`: single port
  - `group`: port address book
* `destination` - (Required) The destination address in the access control policy.
  - When `DestinationType` is set to `net`, Destination specifies the destination CIDR block.
  
  Example: 10.2.3.0/24
  - When `DestinationType` is set to `group`, Destination specifies the name of the destination address book.

  Example: db_group
  - When `DestinationType` is set to `domain`, Destination specifies the destination domain name.

  Example: *.aliyuncs.com.
* `destination_type` - (Required) The type of the destination address in a VPC firewall access control policy. Valid values:
  - `net`: destination CIDR block
  - `group`: destination address book
  - `domain`: destination domain name
* `domain_resolve_type` - (Optional, Computed, Available since v1.267.0) The domain name resolution method of the access control policy. Valid values:

  - `FQDN`: FQDN-based.
  - `DNS`: DNS-based dynamic resolution.
  - `FQDN_AND_DNS`: FQDN and DNS-based dynamic resolution.
* `end_time` - (Optional, Int, Available since v1.267.0) The end time of the validity period for the access control policy. The value is a UNIX timestamp in seconds. The time must be on the hour or half-hour and must be at least 30 minutes later than the start time.

-> **NOTE:**  If RepeatType is set to Permanent, EndTime is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, EndTime is required and you must specify an end time.

* `lang` - (Optional) The language of the request and response.

Valid values:
  - `zh` (default): Chinese
  - `en`: English

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `member_uid` - (Optional, ForceNew, Computed) The UID of the member account that belongs to the current Alibaba Cloud account.
* `old_order` - (Optional, Available since v1.286.0) The original priority of the access control policy before the priority is modified.

-> **NOTE:**  This parameter is not recommended. We recommend that you use the AclUuid parameter to specify the policy to be modified.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `order` - (Required, Int, Deprecated since v1.286.0) The new priority of the access control policy after the priority is modified.

-> **NOTE:**  For the valid values of the new priority, see [Query the effective range of priorities](https://help.aliyun.com/document_detail/474145.html).

* `proto` - (Required) The protocol type of traffic in a VPC firewall access control policy. Valid values:
  - `TCP`
  - `UDP`
  - `ICMP`
  - `ANY` (all protocol types)

-> **NOTE:**  If you do not specify this parameter, all protocol types are queried.

* `release` - (Optional, Computed) The enabled status of the access control policy. By default, the policy is enabled after it is created. Valid values:
  - `true`: The access control policy is enabled.
  - `false`: The access control policy is not enabled.
* `repeat_days` - (Optional, List, Available since v1.267.0) The set of recurring days for the validity period of the access control policy.
  - When RepeatType is set to `Permanent`, `None`, or `Daily`, RepeatDays is an empty set.
  Example: []
  - When RepeatType is set to Weekly, RepeatDays cannot be empty.
  Example: [0, 6]

-> **NOTE:**  When RepeatType is set to Weekly, duplicate values are not allowed in RepeatDays.
  - When RepeatType is set to `Monthly`, RepeatDays cannot be empty.
  Example: [1, 31]

-> **NOTE:**  When RepeatType is set to Monthly, duplicate values are not allowed in RepeatDays.

* `repeat_end_time` - (Optional, Available since v1.267.0) The repeat end time of the validity period for the access control policy. Example: 23:30. The value must be on the hour or half-hour and must be at least 30 minutes later than the repeat start time.

-> **NOTE:**  If RepeatType is set to Permanent or None, RepeatEndTime is empty. If RepeatType is set to Daily, Weekly, or Monthly, RepeatEndTime is required and you must specify a repeat end time. The format is HH:MM (24-hour clock). Example: 08:00.

* `repeat_start_time` - (Optional, Available since v1.267.0) The repeat start time of the validity period for an access control policy. Example: 08:00. The value must be on the hour or half-hour and must be at least 30 minutes earlier than the repeat end time.

-> **NOTE:**  When RepeatType is set to Permanent or None, RepeatStartTime is empty. When RepeatType is set to Daily, Weekly, or Monthly, RepeatStartTime is required and you must specify a repeat start time.
The format is HH:MM (24-hour clock). Example: 08:00.
* `repeat_type` - (Optional, Computed, Available since v1.267.0) The repeat type of the validity period for the access control policy. Valid values:
  - `Permanent` (default): Always.
  - `None`: One-time only.
  - `Daily`: Every day.
  - `Weekly`: Every week.
  - `Monthly`: Every month.
* `source` - (Required) The source address in a VPC firewall access control policy. Valid values:
  - If `SourceType` is set to `net`, Source specifies the source CIDR block.
  - If `SourceType` is set to `group`, Source specifies the name of the address book.
* `source_type` - (Required) The type of the source address in the VPC firewall access control policy. Valid values:
  - `net`: source CIDR block
  - `group`: source address book
* `start_time` - (Optional, Int, Available since v1.267.0) The start time of the validity period for the access control policy. The value is a timestamp in seconds. The time must be on the hour or half-hour and must be at least 30 minutes earlier than the end time.

-> **NOTE:**  If RepeatType is set to Permanent, StartTime is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, StartTime is required and you must specify a start time.

* `vpc_firewall_id` - (Required, ForceNew) The instance ID of the VPC firewall. You can call the [DescribeVpcFirewallAclGroupList](https://help.aliyun.com/document_detail/159760.html) operation to obtain the ID.
  - If the VPC firewall protects a Cloud Enterprise Network (CEN), set this parameter to the CEN instance ID.

  Example: cen-ervw0g12b5jbw*\*\*\*
  - If the VPC firewall protects an Express Connect circuit, set this parameter to the VPC firewall instance ID.

  Example: vfw-a42bbb7b887148c9*\*\*\*

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<vpc_firewall_id>:<acl_uuid>`.
* `acl_uuid` - The unique ID of the access control policy for the VPC firewall.
* `application_id` - The ID of the application that handles traffic in the access control policy for the VPC firewall.
* `create_time` - The time when the policy was created.
* `dest_port_group_ports` - The details of the destination port address book in a VPC firewall access control policy.
* `destination_group_cidrs` - The CIDR block information in the destination address book of the VPC firewall access control policy.
* `destination_group_type` - The type of the destination address book in the access control policy.
* `hit_times` - The number of times that the VPC firewall access control policy is matched.
* `source_group_cidrs` - The details of the source address book in the VPC firewall access control policy.
* `source_group_type` - The type of the source address book in the access control policy.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Vpc Firewall Control Policy.
* `delete` - (Defaults to 5 mins) Used when delete the Vpc Firewall Control Policy.
* `update` - (Defaults to 5 mins) Used when update the Vpc Firewall Control Policy.

## Import

Cloud Firewall Vpc Firewall Control Policy can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_firewall_vpc_firewall_control_policy.example <vpc_firewall_id>:<acl_uuid>
```