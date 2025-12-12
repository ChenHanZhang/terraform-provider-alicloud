---
subcategory: "Hologres (Hologram)"
layout: "alicloud"
page_title: "Alicloud: alicloud_hologram_instance"
description: |-
  Provides a Alicloud Hologres (Hologram) Instance resource.
---

# alicloud_hologram_instance

Provides a Hologres (Hologram) Instance resource.



For information about Hologres (Hologram) Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/zh/hologres/developer-reference/api-hologram-2022-06-01-createinstance).

-> **NOTE:** Available since v1.213.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_vpc" "defaultVpc" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = var.name
}

resource "alicloud_vswitch" "defaultVSwitch" {
  vpc_id       = alicloud_vpc.defaultVpc.id
  zone_id      = "cn-hangzhou-j"
  cidr_block   = "172.16.53.0/24"
  vswitch_name = var.name
}

resource "alicloud_hologram_instance" "default" {
  instance_type = "Standard"
  pricing_cycle = "Hour"
  cpu           = "32"
  endpoints {
    type = "Intranet"
  }
  endpoints {
    type       = "VPCSingleTunnel"
    vswitch_id = alicloud_vswitch.defaultVSwitch.id
    vpc_id     = alicloud_vswitch.defaultVSwitch.vpc_id
  }

  zone_id       = alicloud_vswitch.defaultVSwitch.zone_id
  instance_name = var.name
  payment_type  = "PayAsYouGo"
}
```

### Deleting `alicloud_hologram_instance` or removing it from your configuration

The `alicloud_hologram_instance` resource allows you to manage  `payment_type = "Subscription"`  instance, but Terraform cannot destroy it.
Deleting the subscription resource or removing it from your configuration will remove it from your state file and management, but will not destroy the Instance.
You can resume managing the subscription instance via the AlibabaCloud Console.

## Argument Reference

The following arguments are supported:
* `auto_pay` - (Optional) Whether to pay automatically. The default value is true. Value:
  - true: automatic payment
  - false: only generate orders, not pay

-> **NOTE:**  The default value is true. If the balance of your payment method is insufficient, you can set the parameter AutoPay to false, and an unpaid order will be generated. You can log in to the user Center to pay by yourself.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `cold_storage_size` - (Optional, Int) Instance low-frequency storage space. Unit: GB.

-> **NOTE:**  Pay-As-You-Go (PostPaid) instances ignore this parameter.

* `cpu` - (Optional, Computed, Int) Instance specifications. Value:
  - 8 cores 32 GB (number of compute nodes: 1)
  - 16 cores 64 GB (number of compute nodes: 1)
  - 32 core 128 GB (number of compute nodes: 2)
  - 64 core 256 GB (number of compute nodes: 4)
  - 96 core 384 GB (number of computing nodes: 6)
  - 128 core 512 GB (number of compute nodes: 8)
  - Wait

-> **NOTE:** >

-> **NOTE:** - just fill in the audit number.

-> **NOTE:** - Please submit a work order application for purchasing 1024 or above specifications.

-> **NOTE:** - Shared instance types do not need to specify specifications.

-> **NOTE:**  The specification of - 8 core 32GB (number of computing nodes: 1) is only for experience use and cannot be used for production.

* `duration` - (Optional, Int) The buying cycle. Buy for 2 months.

-> **NOTE:**  If the Payment type is PostPaid, you do not need to specify it.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `enable_ssl` - (Optional, Available since v1.266.0) Secure Sockets Layer
* `endpoints` - (Required, List) List of domain names. See [`endpoints`](#endpoints) below.
* `gateway_count` - (Optional, Int) Number of gateway nodes.
* `initial_databases` - (Optional) Initialize the database and split multiple database names ","

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `instance_name` - (Required) The name of the resource
* `instance_type` - (Required, ForceNew, Available since v1.259.0) The instance type. Value:
  - Standard: Universal.
  - Follower: Read-only slave instance.
  - Warehouse: calculation group type.
  - Shared: Shared.
* `leader_instance_id` - (Optional, ForceNew) The ID of the primary instance.
* `payment_type` - (Required, ForceNew) The payment type of the resource
* `pricing_cycle` - (Optional) Billing cycle. Value:
  - Month: monthly billing
  - Hour: hourly billing

-> **NOTE:** >

-> **NOTE:**  - PrePaid only supports Month

-> **NOTE:**  - PostPaid only supports Hour

-> **NOTE:** - The Shared type is automatically set to Hour without specifying it.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `scale_type` - (Optional) Change matching type. Value:
  - UPGRADE: UPGRADE
  - DOWNGRADE: Downgrading

-> **NOTE:** >

-> **NOTE:** - The upgrade specification cannot be less than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is larger than the original specification.

-> **NOTE:** - The downgrading specification cannot be greater than the original specification. A blank field indicates that the original specification remains unchanged. On this basis, at least one specification is smaller than the original specification.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `status` - (Optional, Computed) The status of the resource
* `storage_size` - (Optional, Int) The standard storage space of the instance. Unit: GB.

-> **NOTE:**  Pay-As-You-Go instances (PostPaid) ignore this parameter.

* `tags` - (Optional, Map) Instance tag
* `zone_id` - (Required, ForceNew) The zone Id. Refer to "Instructions for Use"

### `endpoints`

The endpoints supports the following:
* `type` - (Optional, Computed) The network type.
* `vswitch_id` - (Required) The ID of the virtual switch.
* `vpc_id` - (Required) VPC primary key

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.
* `create_time` - The creation time of the resource
* `endpoints` - List of domain names.
  * `alternative_endpoints` - Some old instances have both AnyTunnel and SingleTunnel enabled. When switching from AnyTunnel to SingleTunnel, the endpoints of both are retained. Therefore, one more field is required to store the Endpoint.
  * `enabled` - Whether to turn on the network.
  * `endpoint` - Domain name.
  * `vpc_instance_id` - The vpc instance ID.
* `region_id` - Region Id. You can go to [API Portal](https://api.aliyun.com/product/Hologram) or "instructions for use" to view.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 10 mins) Used when create the Instance.
* `delete` - (Defaults to 5 mins) Used when delete the Instance.
* `update` - (Defaults to 10 mins) Used when update the Instance.

## Import

Hologres (Hologram) Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_hologram_instance.example <id>
```