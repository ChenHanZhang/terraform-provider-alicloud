---
subcategory: "Rocket M Q"
layout: "alicloud"
page_title: "Alicloud: alicloud_rocket_m_q_consumer_group"
description: |-
  Provides a Alicloud Rocket M Q Consumer Group resource.
---

# alicloud_rocket_m_q_consumer_group

Provides a Rocket M Q Consumer Group resource. 

For information about Rocket M Q Consumer Group and how to use it, see [What is Consumer Group](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available since v1.211.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "createVpc" {
  description = "111"
  cidr_block  = "172.16.0.0/12"
  vpc_name    = var.name

}

resource "alicloud_vswitch" "createVswitch" {
  description  = "1111"
  vpc_id       = alicloud_vpc.createVpc.id
  zone_id      = data.alicloud_zones.default.zones.0.id
  cidr_block   = "172.16.0.0/24"
  vswitch_name = var.name

}

resource "alicloud_rocket_m_q_instance" "createInstance" {
  auto_renew_period = "1"
  product_info {
    msg_process_spec       = "rmq.p2.4xlarge"
    send_receive_ratio     = 0.3
    message_retention_time = "70"
    support_auto_scaling   = true
  }
  network_info {
    vpc_info {
      vpc_id     = alicloud_vpc.createVpc.id
      vswitch_id = alicloud_vswitch.createVswitch.id
    }
    internet_info {
      internet_spec      = "enable"
      flow_out_type      = "payByBandwidth"
      flow_out_bandwidth = "30"
    }
  }
  period          = "1"
  sub_series_code = "cluster_ha"
  remark          = "自动化测试购买使用11"
  instance_name   = var.name

  service_code = "rmq"
  series_code  = "professional"
  payment_type = "PayAsYouGo"
  software {
    maintain_time    = "02:00-06:00"
    upgrade_method   = "auto"
    software_version = "5.0-rmq-20230619-1"
  }
  period_unit = "Month"
}


resource "alicloud_rocket_m_q_consumer_group" "default" {
  consumer_group_id = "pop-test-group"
  instance_id       = alicloud_rocket_m_q_instance.createInstance.id
  consume_retry_policy {
    max_retry_times = "10"
    retry_policy    = "DefaultRetryPolicy"
  }
  delivery_order_type = "Concurrently"
  remark              = "123321"
}
```

## Argument Reference

The following arguments are supported:
* `consume_retry_policy` - (Required) Consumption retry strategy. See [`consume_retry_policy`](#consume_retry_policy) below.
* `consumer_group_id` - (Required, ForceNew) The first ID of the resource.
* `delivery_order_type` - (Optional) Delivery order.
* `instance_id` - (Required, ForceNew) Instance ID.
* `remark` - (Optional) Custom remarks.

### `consume_retry_policy`

The consume_retry_policy supports the following:
* `max_retry_times` - (Optional) Maximum number of retries.
* `retry_policy` - (Optional) Consume retry policy.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<instance_id>:<consumer_group_id>`.
* `create_time` - The creation time of the resource.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Consumer Group.
* `delete` - (Defaults to 5 mins) Used when delete the Consumer Group.
* `update` - (Defaults to 5 mins) Used when update the Consumer Group.

## Import

Rocket M Q Consumer Group can be imported using the id, e.g.

```shell
$ terraform import alicloud_rocket_m_q_consumer_group.example <instance_id>:<consumer_group_id>
```