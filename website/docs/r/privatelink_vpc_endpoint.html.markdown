---
subcategory: "Private Link"
layout: "alicloud"
page_title: "Alicloud: alicloud_privatelink_vpc_endpoint"
description: |-
  Provides a Alicloud Private Link Vpc Endpoint resource.
---

# alicloud_privatelink_vpc_endpoint

Provides a Private Link Vpc Endpoint resource.



For information about Private Link Vpc Endpoint and how to use it, see [What is Vpc Endpoint](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-createvpcendpoint).

-> **NOTE:** Available since v1.109.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "ap-southeast-5"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "defaultbFzA4a" {
  description = "example-terraform"
  cidr_block  = "172.16.0.0/12"
  vpc_name    = var.name
}

resource "alicloud_security_group" "default1FTFrP" {
  name   = var.name
  vpc_id = alicloud_vpc.defaultbFzA4a.id
}

resource "alicloud_security_group" "defaultjljY5S" {
  name   = var.name
  vpc_id = alicloud_vpc.defaultbFzA4a.id
}

resource "alicloud_privatelink_vpc_endpoint" "default" {
  endpoint_description          = var.name
  vpc_endpoint_name             = var.name
  resource_group_id             = data.alicloud_resource_manager_resource_groups.default.ids.0
  endpoint_type                 = "Interface"
  vpc_id                        = alicloud_vpc.defaultbFzA4a.id
  service_name                  = "com.aliyuncs.privatelink.ap-southeast-5.oss"
  dry_run                       = "false"
  zone_private_ip_address_count = "1"
  policy_document               = jsonencode({ "Version" : "1", "Statement" : [{ "Effect" : "Allow", "Action" : ["*"], "Resource" : ["*"], "Principal" : "*" }] })
  security_group_ids = [
    "${alicloud_security_group.default1FTFrP.id}"
  ]
  service_id        = "epsrv-k1apjysze8u1l9t6uyg9"
  protected_enabled = "false"
}
```

## Argument Reference

The following arguments are supported:
* `address_ip_version` - (Optional, Computed, Available since v1.239.0) The IP address version.
* `dry_run` - (Optional) Specifies whether to perform only a dry run, without performing the actual request. Valid values:
  - `true`: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
  - **false (default)**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `endpoint_description` - (Optional) The description of the endpoint.
* `endpoint_type` - (Optional, ForceNew, Computed, Available since v1.212.0) The endpoint type. Value:
  - `Interface`: Indicates an Interface endpoint. You can add service resource types for Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB).
  - `Reverse`: indicates a Reverse terminal node. You can add a VPC NAT Gateway as a service resource.
  - `GatewayLoadBalancer`: indicates the gateway-type load balancing endpoint. You can add a Gateway Load Balancer (GWLB) as a service resource.
* `policy_document` - (Optional, Computed, Available since v1.223.2) RAM access policies. For more information about policy definitions, see Alibaba Cloud-access control (RAM) official guidance.
* `protected_enabled` - (Optional, Available since v1.212.0) Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
  - `true`: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
  - **false (default)**: disables user authentication.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `resource_group_id` - (Optional, Computed, Available since v1.212.0) The resource group ID.
* `security_group_ids` - (Optional, List) The ID of the security group that is associated with the endpoint ENI. The security group can be used to control data transfer between the VPC and the endpoint ENI.
The endpoint can be associated with up to 10 security groups.
* `service_id` - (Optional, ForceNew, Computed) The ID of the endpoint service with which the endpoint is associated.
* `service_name` - (Optional, ForceNew, Computed) The name of the endpoint service with which the endpoint is associated.
* `tags` - (Optional, Map, Available since v1.212.0) The list of tags.
* `vpc_endpoint_name` - (Optional) The name of the endpoint.
* `vpc_id` - (Required, ForceNew) The ID of the VPC to which the endpoint belongs.
* `zone_private_ip_address_count` - (Optional, ForceNew, Computed, Int) The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only 1 is returned.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `bandwidth` - The bandwidth of the endpoint connection.
* `connection_status` - The state of the endpoint connection.
* `connections` - The endpoint connection list.
  * `bandwidth` - The bandwidth of the endpoint connection.
  * `connection_status` - The state of the endpoint connection.
  * `endpoint_owner_id` - The ID of the Alibaba Cloud account to which the endpoint belongs.
  * `modified_time` - The time when the endpoint connection was modified.
  * `zones` - The zone list.
    * `eni_id` - The endpoint ENI ID.
    * `resource_id` - The service resource ID.
    * `vswitch_id` - The ID of the vSwitch to which the endpoint belongs.
    * `zone_domain` - The domain name of the zone.
    * `zone_id` - The zone ID.
* `create_time` - The time when the endpoint was created.
* `endpoint_business_status` - The service state of the endpoint.
* `endpoint_domain` - The domain name of the endpoint.
* `payer` - The payer.
* `resource_owner` - Indicates whether the endpoint and the endpoint service belong to the same Alibaba Cloud account.
* `status` - The state of the endpoint.
* `zone_affinity_enabled` - Indicates whether zone affinity is enabled.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 6 mins) Used when create the Vpc Endpoint.
* `delete` - (Defaults to 5 mins) Used when delete the Vpc Endpoint.
* `update` - (Defaults to 7 mins) Used when update the Vpc Endpoint.

## Import

Private Link Vpc Endpoint can be imported using the id, e.g.

```shell
$ terraform import alicloud_privatelink_vpc_endpoint.example <endpoint_id>
```