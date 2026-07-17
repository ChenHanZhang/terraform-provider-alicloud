---
subcategory: "Cloud Native API Gateway (APIG)"
layout: "alicloud"
page_title: "Alicloud: alicloud_apig_service"
description: |-
  Provides a Alicloud APIG Service resource.
---

# alicloud_apig_service

Provides a APIG Service resource.



For information about APIG Service and how to use it, see [What is Service](https://next.api.alibabacloud.com/document/APIG/2024-03-27/CreateService).

-> **NOTE:** Available since v1.286.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

variable "address" {
  default = "127.0.0.1:8080"
}

variable "address_1" {
  default = "127.0.0.1:7891"
}

variable "address_2" {
  default = "127.0.0.1:7890"
}

resource "alicloud_vpc" "defaultvpc" {
  cidr_block = "172.32.0.0/12"
  vpc_name   = "zhenyuan-example"
}

resource "alicloud_vswitch" "defaultvswitch" {
  vpc_id       = alicloud_vpc.defaultvpc.id
  zone_id      = "cn-hangzhou-g"
  cidr_block   = "172.32.100.0/24"
  vswitch_name = "zhenyuan-example"
}

resource "alicloud_apig_gateway" "defaultFsRKYn" {
  network_access_config {
    type = "Intranet"
  }
  vswitch {
    vswitch_id = alicloud_vswitch.defaultvswitch.id
    name       = alicloud_vswitch.defaultvswitch.vswitch_name
  }
  zone_config {
    select_option = "Auto"
  }
  vpc {
    vpc_id = alicloud_vpc.defaultvpc.id
  }
  payment_type = "PayAsYouGo"
  gateway_name = "zhenyuanexample"
  spec         = "apigw.small.x1"
  log_config {
    sls {
      enable = false
    }
  }
}


resource "alicloud_apig_service" "default" {
  addresses    = ["${var.address}"]
  service_name = "1784259168"
  source_type  = "VIP"
  gateway_id   = alicloud_apig_gateway.defaultFsRKYn.id
  namespace    = "default"
}
```

## Argument Reference

The following arguments are supported:
* `addresses` - (Optional, List) A list of domain names or fixed addresses.
* `ai_service_config` - (Optional, ForceNew, Set) ai service configuration when sourceType equals AI. See [`ai_service_config`](#ai_service_config) below.
* `gateway_id` - (Optional, ForceNew) The ID of the Cloud Native API Gateway.
* `group_name` - (Optional, ForceNew) The service group name.
Required when sourceType is MSE_NACOS.
* `namespace` - (Optional, ForceNew) The namespace of the service:
  - sourceType is K8S, indicating the namespace of the K8S service.
When-sourceType is set to MSE_NACOS, it indicates the namespace in Nacos.

When the sourceType is K8S and MSE_NACOS, it needs to be specified.
* `qualifier` - (Optional, ForceNew) The function version or alias.
* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `service_name` - (Optional, ForceNew) Service Name, need to fill in manually when sourceType is VIP/DNS/AI.
* `source_type` - (Optional, ForceNew) service source type, optional value is K8S/MSE_NACOS/FC3/SAE_K8S_SERVICE/VIP/DNS/AI

### `ai_service_config`

The ai_service_config supports the following:
* `address` - (Optional) ai provider address
* `api_keys` - (Optional, List) api key list
* `enable_health_check` - (Optional) whether enable health check
* `protocols` - (Optional, List) model protocol list
* `provider` - (Optional) ai model provider

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Service.
* `delete` - (Defaults to 5 mins) Used when delete the Service.
* `update` - (Defaults to 5 mins) Used when update the Service.

## Import

APIG Service can be imported using the id, e.g.

```shell
$ terraform import alicloud_apig_service.example <service_id>
```