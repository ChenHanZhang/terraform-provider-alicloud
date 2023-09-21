---
subcategory: "Rocket M Q"
layout: "alicloud"
page_title: "Alicloud: alicloud_rocket_m_q_topic"
description: |-
  Provides a Alicloud Rocket M Q Topic resource.
---

# alicloud_rocket_m_q_topic

Provides a Rocket M Q Topic resource. 

For information about Rocket M Q Topic and how to use it, see [What is Topic](https://www.alibabacloud.com/help/en/).

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
  description  = "111"
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


resource "alicloud_rocket_m_q_topic" "default" {
  instance_id  = alicloud_rocket_m_q_instance.createInstance.id
  message_type = "NORMAL"
  topic_name   = var.name

  remark = "1111"
}
```

## Argument Reference

The following arguments are supported:
* `instance_id` - (Required, ForceNew) Instance ID.
* `message_type` - (Optional, ForceNew) Message type.
* `remark` - (Optional) Custom remarks.
* `topic_name` - (Required, ForceNew) Topic name and identification.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<instance_id>:<topic_name>`.
* `create_time` - The creation time of the resource.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Topic.
* `delete` - (Defaults to 5 mins) Used when delete the Topic.
* `update` - (Defaults to 5 mins) Used when update the Topic.

## Import

Rocket M Q Topic can be imported using the id, e.g.

```shell
$ terraform import alicloud_rocket_m_q_topic.example <instance_id>:<topic_name>
```