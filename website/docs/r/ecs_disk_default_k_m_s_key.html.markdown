---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_disk_default_k_m_s_key"
description: |-
  Provides a Alicloud ECS Disk Default K M S Key resource.
---

# alicloud_ecs_disk_default_k_m_s_key

Provides a ECS Disk Default K M S Key resource.

The encryption key used by default for cloud storage encryption.

For information about ECS Disk Default K M S Key and how to use it, see [What is Disk Default K M S Key](https://next.api.alibabacloud.com/document/Ecs/2014-05-26/ModifyDiskDefaultKMSKeyId).

-> **NOTE:** Available since v1.279.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

variable "cidr_block_vpc" {
  default = "172.16.0.0/12"
}

variable "cidr_block_vsw_k" {
  default = "172.17.0.0/16"
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "zone_id_k" {
  default = "cn-hangzhou-k"
}

variable "zone_id_j" {
  default = "cn-hangzhou-j"
}

resource "alicloud_vpc" "default5RztkC" {
  is_default = false
  cidr_block = var.cidr_block_vpc
}

resource "alicloud_vswitch" "defaultTHeK3U" {
  vpc_id     = alicloud_vpc.default5RztkC.id
  zone_id    = var.zone_id_k
  cidr_block = var.cidr_block_vsw_k
}

resource "alicloud_kms_instance" "defaultLbatIU" {
  vpc_num         = "7"
  key_num         = "2000"
  renew_period    = "1"
  secret_num      = "0"
  product_version = "3"
  renew_status    = "AutoRenewal"
  vpc_id          = alicloud_vpc.default5RztkC.id
  vswitch_ids     = ["${alicloud_vswitch.defaultTHeK3U.id}"]
  zone_ids        = ["cn-hangzhou-k", "${var.zone_id_j}"]
  spec            = "2000"
}

resource "alicloud_kms_key" "defaultC3EYIX" {
  origin           = "Aliyun_KMS"
  status           = "Enabled"
  protection_level = "SOFTWARE"
  key_spec         = "Aliyun_AES_256"
  key_usage        = "ENCRYPT/DECRYPT"
  dkms_instance_id = alicloud_kms_instance.defaultLbatIU.id
}

resource "alicloud_kms_key" "defaultDDxPUN" {
  origin           = "Aliyun_KMS"
  status           = "Enabled"
  protection_level = "SOFTWARE"
  key_spec         = "Aliyun_AES_256"
  key_usage        = "ENCRYPT/DECRYPT"
  dkms_instance_id = alicloud_kms_instance.defaultLbatIU.id
  description      = "KmsKey02"
}


resource "alicloud_ecs_disk_default_k_m_s_key" "default" {
  kms_key_id = alicloud_kms_key.defaultC3EYIX.id
}
```

## Argument Reference

The following arguments are supported:
* `kms_key_id` - (Required) The KMS Key ID of the resource

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Disk Default K M S Key.
* `delete` - (Defaults to 5 mins) Used when delete the Disk Default K M S Key.
* `update` - (Defaults to 5 mins) Used when update the Disk Default K M S Key.

## Import

ECS Disk Default K M S Key can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_disk_default_k_m_s_key.example <region_id>
```