---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_control_policy_order"
description: |-
  Provides a Alicloud Cloud Firewall Control Policy Order resource.
---

# alicloud_cloud_firewall_control_policy_order

Provides a Cloud Firewall Control Policy Order resource.

Access policy priority.

For information about Cloud Firewall Control Policy Order and how to use it, see [What is Control Policy Order](https://www.alibabacloud.com/help/doc-detail/138867.htm).

-> **NOTE:** Available since v1.130.0.

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

resource "alicloud_cloud_firewall_control_policy_order" "default" {
  acl_uuid  = alicloud_cloud_firewall_control_policy.default.acl_uuid
  direction = alicloud_cloud_firewall_control_policy.default.direction
  order     = 1
}
```

### Deleting `alicloud_cloud_firewall_control_policy_order` or removing it from your configuration

Terraform cannot destroy resource `alicloud_cloud_firewall_control_policy_order`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `acl_uuid` - (Required, ForceNew) The Security Access Control Strategy Is the Unique Identifier of the ID.
* `direction` - (Optional) Access Control Policy of the Direction of the Flow
* `lang` - (Optional, Available since v1.286.0) The language type of the received message. Value:
  - zh (default): Chinese.
  - en: English.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `new_order` - (Optional, Available since v1.286.0) Set a new priority for the Internet border firewall IPv4 access control policy.
Priority is represented by a number. Enter the number 1 to indicate the highest priority. The larger the number, the lower the priority.

-> **NOTE:**  The new priority number cannot exceed the scope of the Internet border firewall IPv4 policy priority you have created, otherwise it will cause an error when calling the interface.

-> **NOTE:**  We recommend that you call [DescribePolicyPriorUsed](~~ 138862 ~~) to query the IPv4 policy priority range of the specified traffic direction of the Internet boundary firewall before calling this interface.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `order` - (Required, ForceNew, Int) Security Access Control Policies Will Go into Effect of Priority. The Priority Value Starts from 1, the Smaller the Priority Number, the Higher the Priority.> **-1** Indicates the Lowest Priority.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The time when the policy was created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Control Policy Order.
* `update` - (Defaults to 5 mins) Used when update the Control Policy Order.

## Import

Cloud Firewall Control Policy Order can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_firewall_control_policy_order.example <acl_uuid>
```