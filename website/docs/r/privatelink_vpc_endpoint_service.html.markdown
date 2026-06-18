---
subcategory: "Private Link"
layout: "alicloud"
page_title: "Alicloud: alicloud_privatelink_vpc_endpoint_service"
description: |-
  Provides a Alicloud Private Link Vpc Endpoint Service resource.
---

# alicloud_privatelink_vpc_endpoint_service

Provides a Private Link Vpc Endpoint Service resource.

Endpoint service resource.

For information about Private Link Vpc Endpoint Service and how to use it, see [What is Vpc Endpoint Service](https://www.alibabacloud.com/help/en/privatelink/latest/api-privatelink-2020-04-15-createvpcendpointservice).

-> **NOTE:** Available since v1.109.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf_example"
}

resource "alicloud_privatelink_vpc_endpoint_service" "example" {
  service_description    = var.name
  connect_bandwidth      = 103
  auto_accept_connection = false
}
```

## Argument Reference

The following arguments are supported:
* `address_ip_version` - (Optional, Computed, Available since v1.239.0) The IP version. Valid values:
  - `IPv4`: IPv4 type (default).
  - `DualStack`: Dual-stack type.

-> **NOTE:**  Currently, only endpoint services whose backend resource types are NLB or GWLB support specifying the IP address protocol as DualStack. For an endpoint service to support dual stack, its backend resources must also support dual stack.

* `auto_accept_connection` - (Optional) Specifies whether to automatically accept endpoint connections. Valid values:
  - `true`: Automatically accepts endpoint connections.
  - `false`: Does not automatically accept endpoint connections.
* `connect_bandwidth` - (Optional, Computed, Int) The default peak bandwidth for connections. Default value: `3072`. Unit: Mbps.

Valid range: `100` to `10240`.

-> **NOTE:**  When the service resource is a Classic Load Balancer (CLB) or an Application Load Balancer (ALB), you can set the default peak bandwidth. When the service resource is a Network Load Balancer (NLB), setting the connection bandwidth is not supported.

* `dry_run` - (Optional) Specifies whether to only precheck the request. Valid values:
  - `true`: Sends a dry-run request without modifying any attributes of the endpoint service resource. The check includes validation of required parameters, request format, and business constraints. If the check fails, an error is returned. If the check passes, the error code `DryRunOperation` is returned.
  - `false` (default): Sends a normal request. After passing the checks, an HTTP 2xx status code is returned and the operation is performed.

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `payer` - (Optional, ForceNew, Computed) The payer. Valid values:
  - `Endpoint`: The service consumer.
  - `EndpointService`: The service provider.
* `resource_group_id` - (Optional, Computed) The resource group ID.
* `service_description` - (Optional) The description of the endpoint service.
* `service_resource_type` - (Optional, ForceNew, Computed) The type of service resource. Valid values:
  - `slb`: indicates that the service resource type is Classic Load Balancer (CLB).
  - `alb`: indicates that the service resource type is Application Load Balancer (ALB).
  - `nlb`: indicates that the service resource type is Network Load Balancer (NLB).
  - `gwlb`: indicates that the service resource type is Gateway Load Balancer (GWLB).

-> **NOTE:**  TCPSSL listeners of NLB are not supported.

* `service_support_ipv6` - (Optional, Computed) Specifies whether the endpoint service supports IPv6. Valid values:
  - `true`: Yes.
  - `false` (default): No.
* `tags` - (Optional, Map) A list of tags.
* `zone_affinity_enabled` - (Optional, Computed) Specifies whether the endpoint domain name of the connected service supports proximity-based DNS resolution. Valid values:
  - `true` (default): Yes.
  - `false`: No.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the endpoint service.
* `max_bandwidth` - The maximum bandwidth of the endpoint connection.
* `min_bandwidth` - The minimum bandwidth of the endpoint connection.
* `service_business_status` - The business status of the endpoint service.
* `service_domain` - The service domain name of the endpoint service.
* `status` - The status of the endpoint service.
* `vpc_endpoint_service_name` - The name of the endpoint service that can be associated when creating an endpoint.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Vpc Endpoint Service.
* `delete` - (Defaults to 5 mins) Used when delete the Vpc Endpoint Service.
* `update` - (Defaults to 5 mins) Used when update the Vpc Endpoint Service.

## Import

Private Link Vpc Endpoint Service can be imported using the id, e.g.

```shell
$ terraform import alicloud_privatelink_vpc_endpoint_service.example <service_id>
```