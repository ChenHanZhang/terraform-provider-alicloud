---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_vpc_route_target_group"
description: |-
  Provides a Alicloud VPC Route Target Group resource.
---

# alicloud_vpc_route_target_group

Provides a VPC Route Target Group resource.



For information about VPC Route Target Group and how to use it, see [What is Route Target Group](https://next.api.alibabacloud.com/document/Vpc/2016-04-28/CreateRouteTargetGroup).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "eu-central-1"
}

variable "vpc_name" {
  default = "预发-GWLB-业务VPC-勿删"
}

variable "gwlbe_name_a1" {
  default = "预发-GWLBE-可用区A-勿删"
}

variable "gwlbe_name_b2" {
  default = "预发-GWLBE-可用区B-2-勿删"
}

variable "eps_name_a" {
  default = "预发-GWLB-GES-可用区A-勿删"
}

variable "eps_name_b" {
  default = "预发-GWLB-GES-可用区B-勿删"
}

variable "gwlbe_name_b1" {
  default = "预发-GWLBE-可用区B-勿删"
}

variable "region" {
  default = "eu-central-1"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "defaultVpc" {
  vpc_name    = var.vpc_name
  cidr_block  = "192.168.0.0/16"
  enable_ipv6 = false
}

resource "alicloud_privatelink_vpc_endpoint_service" "getVpcEndpointServiceA" {
  auto_accept_connection = true
  service_description    = var.eps_name_a
}

resource "alicloud_privatelink_vpc_endpoint" "getVpcEndpointA" {
  service_id        = alicloud_privatelink_vpc_endpoint_service.getVpcEndpointServiceA.id
  vpc_endpoint_name = var.gwlbe_name_a1
  vpc_id            = alicloud_vpc.defaultVpc.id
  service_name      = alicloud_privatelink_vpc_endpoint_service.getVpcEndpointServiceA.vpc_endpoint_service_name
  endpoint_type     = "GatewayLoadBalancer"
}

resource "alicloud_privatelink_vpc_endpoint_service" "getVpcEndpointServiceB" {
  auto_accept_connection = true
  service_description    = var.eps_name_b
}

resource "alicloud_privatelink_vpc_endpoint" "getVpcEndpointB" {
  service_id        = alicloud_privatelink_vpc_endpoint_service.getVpcEndpointServiceB.id
  vpc_endpoint_name = var.gwlbe_name_b1
  vpc_id            = alicloud_vpc.defaultVpc.id
  service_name      = alicloud_privatelink_vpc_endpoint_service.getVpcEndpointServiceB.vpc_endpoint_service_name
  endpoint_type     = "GatewayLoadBalancer"
}

resource "alicloud_privatelink_vpc_endpoint" "getVpcEndpointB2" {
  service_id        = alicloud_privatelink_vpc_endpoint_service.getVpcEndpointServiceB.id
  vpc_endpoint_name = var.gwlbe_name_b2
  vpc_id            = alicloud_vpc.defaultVpc.id
  service_name      = alicloud_privatelink_vpc_endpoint_service.getVpcEndpointServiceB.vpc_endpoint_service_name
  endpoint_type     = "GatewayLoadBalancer"
}


resource "alicloud_vpc_route_target_group" "default" {
  route_target_group_description = "预发-路由目标组-资源example"
  vpc_id                         = alicloud_vpc.defaultVpc.id
  route_target_member_list {
    member_id           = alicloud_privatelink_vpc_endpoint.getVpcEndpointA.id
    member_type         = "GatewayLoadBalancerEndpoint"
    weight              = "100"
    enable_status       = "Enable"
    health_check_status = "Normal"
  }
  route_target_member_list {
    member_id           = alicloud_privatelink_vpc_endpoint.getVpcEndpointB.id
    member_type         = "GatewayLoadBalancerEndpoint"
    weight              = "0"
    enable_status       = "Disable"
    health_check_status = "Normal"
  }
  config_mode             = "Active-Standby"
  route_target_group_name = "预发-路由目标组-资源example"
}
```

## Argument Reference

The following arguments are supported:
* `config_mode` - (Required, ForceNew) The working mode of the routing target group:
  - Active-Standby
* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `route_target_group_description` - (Optional) The description.
* `route_target_group_name` - (Optional) The name of the resource
* `route_target_member_list` - (Required, List) The list of RouteTargetMember. See [`route_target_member_list`](#route_target_member_list) below.
* `tags` - (Optional, Map) The tag of the resource
* `vpc_id` - (Required, ForceNew) The VPC ID of RoutetargetGroup.

### `route_target_member_list`

The route_target_member_list supports the following:
* `member_id` - (Required) The instance id of RouteTargetMember.
* `member_type` - (Required) The instance type of the RouteTargetConfig, now support:
  - GatewayLoadBalancerEndpoint
* `weight` - (Required, Int) Sets the weight property of the current route target configuration.

In Active-Standby mode can be set to 0 or 100:
  - Only one route target configuration can be set to 100, as the active instance.
  - Only one route target configuration can be set to 0, as the standby instance.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.
* `region_id` - The region ID of the resource.
* `route_target_member_list` - The list of RouteTargetMember.
  * `enable_status` - Identify the enabled state of the current route target configuration:.
  * `health_check_status` - The health check status of the current RouteTargetConfig.
* `status` - The Status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Route Target Group.
* `delete` - (Defaults to 5 mins) Used when delete the Route Target Group.
* `update` - (Defaults to 5 mins) Used when update the Route Target Group.

## Import

VPC Route Target Group can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_route_target_group.example <route_target_group_id>
```