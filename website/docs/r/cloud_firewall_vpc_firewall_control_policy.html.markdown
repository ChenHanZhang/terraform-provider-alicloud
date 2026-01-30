---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_vpc_firewall_control_policy"
description: |-
  Provides a Alicloud Cloud Firewall Vpc Firewall Control Policy resource.
---

# alicloud_cloud_firewall_vpc_firewall_control_policy

Provides a Cloud Firewall Vpc Firewall Control Policy resource.

VPC Control Policy.

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
* `acl_action` - (Required) The action that the VPC firewall applies to matching traffic in the access control policy. Valid values:
  - `accept`: Allow
  - `drop`: Deny
  - `log`: Monitor

-> **NOTE:**  If this parameter is not specified, all actions are queried.

* `application_name` - (Optional) The application type supported by the VPC firewall access control policy. Valid values:
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
  - `ANY` (indicates all application types are supported).
* `application_name_list` - (Optional, List, Available since v1.267.0) The list of application types supported by the access control policy.
* `description` - (Required) The description of the VPC firewall access control policy. Fuzzy search is supported.
* `dest_port` - (Optional, Computed) The destination port of the traffic in the VPC firewall access control policy.
* `dest_port_group` - (Optional) Name of the destination port address book for traffic in the VPC boundary firewall access control policy.  

-> **NOTE:**  Set this parameter when `DestPortType` is `group`.  

* `dest_port_type` - (Optional, Computed) The type of destination port for traffic in the VPC border firewall access control policy. Valid values:
  - `port`: a single port
  - `group`: a port address book.
* `destination` - (Required) The destination address in the access control policy.
  - When `DestinationType` is `net`, Destination is a destination CIDR address.  
  Example: 10.2.3.0/24
  - When `DestinationType` is `group`, Destination is the name of a destination address book.  
  Example: db_group
  - When `DestinationType` is `domain`, Destination is a destination domain name.  
  Example: *.aliyuncs.com.
* `destination_type` - (Required) The type of destination address in the VPC border firewall access control policy. Valid values:
  - `net`: destination CIDR block
  - `group`: destination address book
  - `domain`: destination domain name.
* `domain_resolve_type` - (Optional, Computed, Available since v1.267.0) The domain name resolution method of the access control policy. Valid values:  

  - `FQDN`: Based on FQDN  
  - `DNS`: Based on dynamic DNS resolution  
  - `FQDN_AND_DNS`: Based on both FQDN and dynamic DNS resolution.
* `end_time` - (Optional, Int, Available since v1.267.0) The end time of the policy validity period for the access control policy, expressed as a Unix timestamp in seconds. The time must be on the hour or half-hour and at least 30 minutes later than the start time.  

-> **NOTE:**  When RepeatType is Permanent, EndTime is empty. When RepeatType is None, Daily, Weekly, or Monthly, EndTime must be specified.

* `lang` - (Optional) The language used for requests and responses.

Valid values:
  - `zh` (default): Chinese
  - `en`: English

-> **NOTE:** This parameter only applies during resource creation, update or deletion. If modified in isolation without other property changes, Terraform will not trigger any action.

* `member_uid` - (Optional, ForceNew, Computed) The UID of the member account under the current Alibaba Cloud account.
* `old_order` - (Optional, Available since v1.270.0) The original priority of the access control policy before its priority is modified.

-> **NOTE:**  We do not recommend that you use this parameter. We recommend that you use the AclUuid parameter to specify the policy to be modified.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `order` - (Required, Int) Priority at which the VPC boundary firewall access control policy takes effect.  
Priority numbers start from 1 and increment sequentially. A smaller number indicates a higher priority.  
* `proto` - (Required) The protocol type of traffic in the VPC border firewall access control policy. Valid values:
  - `TCP`
  - `UDP`
  - `ICMP`
  - `ANY` (queries all protocol types)

-> **NOTE:**  If this parameter is not specified, all protocol types are queried.

* `release` - (Optional, Computed) The enabled status of the access control policy. By default, the policy is enabled after creation. Valid values:  
  - `true`: Enable the access control policy  
  - `false`: Do not enable the access control policy.  
* `repeat_days` - (Optional, List, Available since v1.267.0) The set of repeat days for the validity period of the access control policy.
  - When RepeatType is `Permanent`, `None`, or `Daily`, RepeatDays is an empty set.
  Example: []
  - When RepeatType is `Weekly`, RepeatDays must not be empty.
  Example: [0, 6]

-> **NOTE:**  When RepeatType is set to `Weekly`, RepeatDays must not contain duplicate values.
  - When RepeatType is `Monthly`, RepeatDays must not be empty.
  Example: [1, 31]

-> **NOTE:**  When RepeatType is set to `Monthly`, RepeatDays must not contain duplicate values.

* `repeat_end_time` - (Optional, Available since v1.267.0) The repeat end time of the access control policy validity period. For example, 23:30. The time must be on the hour or half-hour and at least 30 minutes later than the repeat start time.  

-> **NOTE:**  When RepeatType is Permanent or None, RepeatEndTime is empty. When RepeatType is Daily, Weekly, or Monthly, RepeatEndTime must be specified, and you must set the repeat end time.

* `repeat_start_time` - (Optional, Available since v1.267.0) The repeat start time of the access control policy validity period. For example, 08:00. The time must be on the hour or half-hour and at least 30 minutes earlier than the repeat end time.  

-> **NOTE:**  When RepeatType is Permanent or None, RepeatStartTime is empty. When RepeatType is Daily, Weekly, or Monthly, RepeatStartTime must be specified, and you must set the repeat start time.

* `repeat_type` - (Optional, Computed, Available since v1.267.0) The repeat type of the access control policy validity period. Valid values:  
  - `Permanent` (default): Always  
  - `None`: Single specified time  
  - `Daily`: Daily  
  - `Weekly`: Weekly  
  - `Monthly`: Monthly.
* `source` - (Required) The source address in the VPC border firewall access control policy. Valid values:
  - When `SourceType` is `net`, Source specifies the source CIDR block.
  - When `SourceType` is `group`, Source specifies the name of the address book.
* `source_type` - (Required) The source address type in the VPC firewall access control policy. Valid values:  
  - `net`: Source CIDR block  
  - `group`: Source address book
* `start_time` - (Optional, Int, Available since v1.267.0) The start time of the validity period for the access control policy, expressed as a Unix timestamp in seconds. The time must be on the hour or half-hour and at least 30 minutes earlier than the end time.

-> **NOTE:**  When RepeatType is set to Permanent, StartTime must be empty. When RepeatType is set to None, Daily, Weekly, or Monthly, StartTime must be specified.

* `vpc_firewall_id` - (Required, ForceNew) The instance ID of the VPC firewall. You can call the [DescribeVpcFirewallAclGroupList](https://help.aliyun.com/document_detail/159760.html) operation to obtain this ID.
  - When the VPC firewall protects traffic between two VPCs connected through Cloud Enterprise Network (CEN), use the CEN instance ID.

  Example: cen-ervw0g12b5jbw***
  - When the VPC firewall protects traffic between two VPCs connected through Express Connect, use the VPC firewall instance ID.

  Example: vfw-a42bbb7b887148c9***

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<vpc_firewall_id>:<acl_uuid>`.
* `acl_uuid` - The unique identifier ID of the VPC border firewall access control policy.
* `application_id` - The ID of the application for which access traffic is configured in the VPC border firewall access control policy.
* `create_time` - The time when the policy was created, expressed as a Unix timestamp in seconds.
* `dest_port_group_ports` - The detailed information about the destination port address book in the VPC border firewall access control policy.
* `destination_group_cidrs` - The CIDR blocks in the destination address book of the VPC firewall access control policy.
* `destination_group_type` - The type of destination address book in the access control policy.
* `hit_times` - The number of times the VPC border firewall access control policy has been matched.
* `source_group_cidrs` - Details of the source address book in the VPC boundary firewall access control policy.
* `source_group_type` - The type of source address book in the access control policy.

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