---
subcategory: "RAM"
layout: "alicloud"
page_title: "Alicloud: alicloud_ram_mfadevice"
description: |-
  Provides a Alicloud RAM M F A Device resource.
---

# alicloud_ram_mfadevice

Provides a RAM M F A Device resource.



For information about RAM M F A Device and how to use it, see [What is M F A Device](https://next.api.alibabacloud.com/document/Ims/2019-08-15/CreateVirtualMFADevice).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}


resource "alicloud_ram_mfadevice" "default" {
  virtual_mfa_device_name = "zaijiuTestMFA1"
}
```

## Argument Reference

The following arguments are supported:
* `virtual_mfa_device_name` - (Required) The name of the MFA device.
The name must be 1 to 64 characters in length and can contain letters, digits, periods (.), and hyphens (-).

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `activate_date` - Activation time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the M F A Device.
* `delete` - (Defaults to 5 mins) Used when delete the M F A Device.

## Import

RAM M F A Device can be imported using the id, e.g.

```shell
$ terraform import alicloud_ram_mfadevice.example <serial_number>
```