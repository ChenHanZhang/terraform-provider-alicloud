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
* `address_ip_version` - (Optional, Computed, Available since v1.239.0) The IP version. Valid values:
  - `IPv4`: IPv4 (default).
  - `DualStack`: Dual stack.

-> **NOTE:**  To support dual stack, both the connected endpoint service and the VPC where the endpoint resides must support dual stack.

* `cross_region_bandwidth` - (Optional, Computed, Int, Available since v1.282.0) The cross-region bandwidth value. This parameter applies only when the endpoint and the endpoint service are in different regions. Unit: Mbps. Valid values:
  - `Minimum`: 100.
  - `Maximum`: Limited by your account quota (for details, see the [Quotas and limits](https://help.aliyun.com/zh/privatelink/quotas-and-limits?spm=a2c4g.11174283.help-menu-search-120462.d_0) section).

-> **NOTE:**  If you want to specify this parameter, ensure that you are working with a cross-region endpoint.

* `dry_run` - (Optional) Specifies whether to only precheck the request. Valid values:
  - `true`: Sends a dry-run request without creating the endpoint. The system checks whether required parameters are specified, whether the request format is valid, and whether the request complies with business limits. If the check fails, an error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
  - `false` (default): Sends a normal request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `endpoint_description` - (Optional) The description of the endpoint.
The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
* `endpoint_type` - (Optional, ForceNew, Computed, Available since v1.212.0) The endpoint type. Valid values:
  - `Interface`: indicates an interface endpoint. You can add service resource types such as Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB).
  - `Reverse`: indicates a reverse endpoint. You can add a VPC NAT gateway as a service resource.
  - `GatewayLoadBalancer`: indicates a Gateway Load Balancer endpoint. You can add Gateway Load Balancer (GWLB) as a service resource.

-> **NOTE:**  Services that support reverse endpoints can only be provided by Alibaba Cloud and its ecosystem partners. By default, you cannot create such services yourself. If you need to create one, contact your account manager to apply.

* `policy_document` - (Optional, Computed, Available since v1.223.2) The RAM policy. For more information about policy definitions, see [Basic elements of a policy](https://help.aliyun.com/document_detail/93738.html).
* `protected_enabled` - (Optional, Available since v1.212.0) Specifies whether managed protection is enabled. This setting takes effect only for STS-based API calls. Valid values:
  - `true`: Enabled. After managed protection is enabled, only the user who created the endpoint can modify or delete it through STS.
  - `false` (default): Disabled.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `resource_group_id` - (Optional, Computed, Available since v1.212.0) The ID of the resource group.
* `security_group_ids` - (Optional, List) The list of security groups associated with the endpoint ENI.
* `service_id` - (Optional, ForceNew, Computed) The ID of the endpoint service associated with the endpoint.
* `service_name` - (Optional, ForceNew, Computed) The name of the endpoint service associated with the endpoint.
* `service_region_id` - (Optional, ForceNew, Computed, Available since v1.282.0) The region ID of the endpoint service associated with the endpoint.
* `tags` - (Optional, Map, Available since v1.212.0) The list of tags.
* `vpc_endpoint_name` - (Optional) The name of the endpoint.
The name must be 2 to 128 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). It must start with a letter or a Chinese character.
* `vpc_id` - (Required, ForceNew) The virtual private cloud (VPC) to which the endpoint belongs.
* `zone_private_ip_address_count` - (Optional, ForceNew, Computed, Int) The number of private IP addresses for the ENI in each zone. Valid value: `1` only.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `bandwidth` - The connection bandwidth of the endpoint.
* `connection_status` - The endpoint connection status.
* `connections` - VPC endpoint connection information.
  * `bandwidth` - The bandwidth of the VPC endpoint connection, in Mbps.
  * `connection_status` - The status of the VPC endpoint connection.
  * `endpoint_owner_id` - The ID of the account to which the VPC endpoint belongs.
  * `modified_time` - The time when the connection was modified.
  * `zones` - List of zone information.
    * `eni_id` - The endpoint ENI ID.
    * `resource_id` - The service resource ID.
    * `vswitch_id` - The vSwitch to which the endpoint belongs.
    * `zone_domain` - The zone domain name.
    * `zone_id` - The zone ID.
* `create_time` - The creation time of the endpoint.
* `endpoint_business_status` - The business status of the endpoint.
* `endpoint_domain` - The endpoint domain name.
* `payer` - The payer.
* `resource_owner` - Indicates whether the endpoint and the endpoint service belong to the same account.
* `status` - The status of the endpoint.
* `zone_affinity_enabled` - Specifies whether the endpoint domain name of the connected service supports zone-affinity-based DNS resolution.

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