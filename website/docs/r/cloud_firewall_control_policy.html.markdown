---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_control_policy"
description: |-
  Provides a Alicloud Cloud Firewall Control Policy resource.
---

# alicloud_cloud_firewall_control_policy

Provides a Cloud Firewall Control Policy resource.

Access Control Policy.

For information about Cloud Firewall Control Policy and how to use it, see [What is Control Policy](https://www.alibabacloud.com/help/doc-detail/138867.htm).

-> **NOTE:** Available since v1.129.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

resource "alicloud_cloud_firewall_control_policy" "default" {
  direction        = "in"
  application_name = "ANY"
  description      = var.name
  acl_action       = "accept"
  source           = "127.0.0.1/32"
  source_type      = "net"
  destination      = "127.0.0.2/32"
  destination_type = "net"
  proto            = "ANY"
}
```

## Argument Reference

The following arguments are supported:
* `acl_action` - (Required) The action that Cloud Firewall performs on traffic based on the access control policy. Valid values:
  - `accept`: allows the traffic
  - `drop`: blocks the traffic
  - `log`: monitors the traffic
* `application_name` - (Optional) The application type supported by the access control policy. The following application types are supported:
  - `ANY`
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

-> **NOTE:**  `ANY` indicates that the policy applies to all application types.

-> **NOTE:**  You must specify either ApplicationNameList or ApplicationName. They cannot both be left empty. If both parameters are specified, ApplicationNameList takes precedence.

* `application_name_list` - (Optional, List, Available since v1.232.0) The list of application names.

-> **NOTE:**  You must specify either ApplicationNameList or ApplicationName. They cannot both be empty. If both ApplicationNameList and ApplicationName are specified, ApplicationNameList takes precedence.

* `description` - (Required) The description of the access control policy.
* `dest_port` - (Optional, Computed) The destination port for traffic in the access control policy.
* `dest_port_group` - (Optional) The name of the destination port address book for traffic in the access control policy.
* `dest_port_type` - (Optional, Computed) The destination port type for traffic in the access control policy. Valid values:
  - `port`: port
  - `group`: port address book.
* `destination` - (Required) The destination address in the access control policy.
  - When `DestinationType` is set to net, `Destination` specifies the destination CIDR block. Example: 1.2.XX.XX/24.
  - When `DestinationType` is set to group, `Destination` specifies the name of the destination address book. Example: db_group.
  - When `DestinationType` is set to domain, `Destination` specifies the destination domain name. Example: *.aliyuncs.com.
  - When `DestinationType` is set to location, `Destination` specifies the destination region (see the following section for specific region codes). Example: \["BJ11", "ZB"\].
* `destination_type` - (Required) The destination address type in the access control policy. Valid values:
  - `net`: destination CIDR block
  - `group`: destination address book
  - `domain`: destination domain name
  - `location`: destination region.
* `direction` - (Required, ForceNew) The traffic direction of the access control policy. Valid values:
  - `in`: inbound traffic (from the Internet to an internal network)
  - `out`: outbound traffic (from an internal network to the Internet).
* `domain_resolve_type` - (Optional, Available since v1.232.0) The domain name resolution method of the access control policy. Valid values:

  - `FQDN`: FQDN-based
  - `DNS`: DNS-based dynamic resolution
  - `FQDN_AND_DNS`: FQDN and DNS-based dynamic resolution
* `end_time` - (Optional, Int, Available since v1.232.0) The end time of the validity period for the access control policy. This value is a UNIX timestamp in seconds. The time must be on the hour or half-hour and must be at least 30 minutes later than the start time.

-> **NOTE:**  If RepeatType is set to Permanent, EndTime is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, EndTime is required and you must specify an end time.

* `ip_version` - (Optional, ForceNew, Computed, Int) The supported IP address version.

Valid values:
  - `4`: IPv4 address
  - `6`: IPv6 address
* `lang` - (Optional) The language of the request and response messages.

Valid values:
  - `zh` (default): Chinese
  - `en`: English

-> **NOTE:** This parameter is only evaluated during resource creation and deletion. Modifying it in isolation will not trigger any action.

* `proto` - (Required) The protocol type for traffic in the access control policy. Supported application types include:
  - `ANY`
  - `TCP`
  - `UDP`
  - `ICMP`

-> **NOTE:**  `ANY` indicates that the policy applies to all application types.

-> **NOTE:**  For outbound traffic, if the destination address belongs to a threat intelligence address book or a cloud service address book of the domain name type, you can select TCP or ANY as the protocol. If you select TCP, you can choose from five applications: HTTP, HTTPS, SMTP, SMTPS, and SSL. If you select ANY, the application must be set to ANY.

* `release` - (Optional, Computed) Specifies whether the access control policy is enabled. By default, a newly created policy is enabled. Valid values:
  - `true`: Enable the access control policy.
  - `false`: Do not enable the access control policy.
* `repeat_days` - (Optional, List, Available since v1.232.0) The set of recurring days for the validity period of the access control policy.
  - If RepeatType is set to `Permanent`, `None`, or `Daily`, RepeatDays is an empty set.
  Example: []
  - If RepeatType is set to Weekly, RepeatDays cannot be empty.
  Example: [0, 6]

-> **NOTE:**  If RepeatType is set to Weekly, duplicate values are not allowed in RepeatDays.
  - If RepeatType is set to `Monthly`, RepeatDays cannot be empty.
  Example: [1, 31]

-> **NOTE:**  If RepeatType is set to Monthly, duplicate values are not allowed in RepeatDays.

* `repeat_end_time` - (Optional, Available since v1.232.0) The repeat end time of the validity period for the access control policy. Example: 23:30. The value must be on the hour or half-hour and must be at least 30 minutes later than the repeat start time.

-> **NOTE:**  When RepeatType is set to Permanent or None, RepeatEndTime is left empty. When RepeatType is set to Daily, Weekly, or Monthly, RepeatEndTime is required and you must specify a repeat end time.

-> **NOTE:**  The time format is HH:mm (24-hour clock), such as 08:00 or 23:30.

* `repeat_start_time` - (Optional, Available since v1.232.0) The repeat start time of the validity period for the access control policy. Example: 08:00. The value must be on the hour or half-hour and must be at least 30 minutes earlier than the repeat end time.

-> **NOTE:**  When RepeatType is set to Permanent or None, RepeatStartTime is left empty. When RepeatType is set to Daily, Weekly, or Monthly, RepeatStartTime is required and you must specify a repeat start time.

-> **NOTE:**  The time format is HH:mm (24-hour clock), such as 08:00 or 23:30.

* `repeat_type` - (Optional, Computed, Available since v1.232.0) The repeat type of the validity period for the access control policy. Valid values:
  - `Permanent` (default): Always
  - `None`: One-time only
  - `Daily`: Every day
  - `Weekly`: Every week
  - `Monthly`: Every month
* `source` - (Required) The source address in the access control policy.
  - If `SourceType` is set to net, `Source` specifies the source CIDR block. Example: 1.2.XX.XX/24.
  - If `SourceType` is set to group, `Source` specifies the name of the source address book. Example: db_group.
  - If `SourceType` is set to location, `Source` specifies the source region code (see the following section for specific region codes). Example: \["BJ11", "ZB"\].
* `source_type` - (Required) The type of the source address in the access control policy. Valid values:
  - `net`: source CIDR block
  - `group`: source address book
  - `location`: source region
* `start_time` - (Optional, Int, Available since v1.232.0) The start time of the validity period for the access control policy. The value is a timestamp in seconds. The time must be on the hour or half-hour and must be at least 30 minutes earlier than the end time.

-> **NOTE:**  If RepeatType is set to Permanent, StartTime is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, StartTime is required and you must specify a start time.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<direction>:<ip_version>:<acl_uuid>`.
* `acl_uuid` - The unique identifier of the access control policy.
* `application_id` - The application ID configured for traffic access in the access control policy.
* `create_time` - The time when the policy was created.
* `dest_port_group_ports` - The list of ports contained in the destination port address book.
* `destination_group_cidrs` - The list of CIDR blocks in the destination address book of the access control policy.
* `destination_group_type` - The type of the destination address book in the access control policy.
* `dns_result` - The DNS resolution result.
* `dns_result_time` - The timestamp of DNS resolution.
* `hit_times` - The number of times the access control policy is matched.
* `source_group_cidrs` - The list of CIDR blocks in the source address book of the access control policy.
* `source_group_type` - The type of the source address book in the access control policy.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Control Policy.
* `delete` - (Defaults to 5 mins) Used when delete the Control Policy.
* `update` - (Defaults to 5 mins) Used when update the Control Policy.

## Import

Cloud Firewall Control Policy can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_firewall_control_policy.example <direction>:<ip_version>:<acl_uuid>
```