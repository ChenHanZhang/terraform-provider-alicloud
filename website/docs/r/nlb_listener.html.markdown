---
subcategory: "Network Load Balancer (NLB)"
layout: "alicloud"
page_title: "Alicloud: alicloud_nlb_listener"
description: |-
  Provides a Alicloud Network Load Balancer (NLB) Listener resource.
---

# alicloud_nlb_listener

Provides a Network Load Balancer (NLB) Listener resource.



For information about Network Load Balancer (NLB) Listener and how to use it, see [What is Listener](https://www.alibabacloud.com/help/en/server-load-balancer/latest/api-nlb-2022-04-30-createlistener).

-> **NOTE:** Available since v1.191.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf-example"
}
data "alicloud_resource_manager_resource_groups" "default" {}
data "alicloud_nlb_zones" "default" {}
resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "10.4.0.0/16"
}
resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  cidr_block   = "10.4.0.0/24"
  vpc_id       = alicloud_vpc.default.id
  zone_id      = data.alicloud_nlb_zones.default.zones.0.id
}
resource "alicloud_vswitch" "default1" {
  vswitch_name = var.name
  cidr_block   = "10.4.1.0/24"
  vpc_id       = alicloud_vpc.default.id
  zone_id      = data.alicloud_nlb_zones.default.zones.1.id
}

resource "alicloud_security_group" "default" {
  name   = var.name
  vpc_id = alicloud_vpc.default.id
}

resource "alicloud_nlb_load_balancer" "default" {
  load_balancer_name = var.name
  resource_group_id  = data.alicloud_resource_manager_resource_groups.default.ids.0
  load_balancer_type = "Network"
  address_type       = "Internet"
  address_ip_version = "Ipv4"
  vpc_id             = alicloud_vpc.default.id
  tags = {
    Created = "TF",
    For     = "example",
  }
  zone_mappings {
    vswitch_id = alicloud_vswitch.default.id
    zone_id    = data.alicloud_nlb_zones.default.zones.0.id
  }
  zone_mappings {
    vswitch_id = alicloud_vswitch.default1.id
    zone_id    = data.alicloud_nlb_zones.default.zones.1.id
  }
}

resource "alicloud_nlb_server_group" "default" {
  resource_group_id        = data.alicloud_resource_manager_resource_groups.default.ids.0
  server_group_name        = var.name
  server_group_type        = "Instance"
  vpc_id                   = alicloud_vpc.default.id
  scheduler                = "Wrr"
  protocol                 = "TCP"
  connection_drain_enabled = true
  connection_drain_timeout = 60
  address_ip_version       = "Ipv4"
  health_check {
    health_check_enabled         = true
    health_check_type            = "TCP"
    health_check_connect_port    = 0
    healthy_threshold            = 2
    unhealthy_threshold          = 2
    health_check_connect_timeout = 5
    health_check_interval        = 10
    http_check_method            = "GET"
    health_check_http_code       = ["http_2xx", "http_3xx", "http_4xx"]
  }
  tags = {
    Created = "TF",
    For     = "example",
  }
}

resource "alicloud_nlb_listener" "default" {
  listener_protocol      = "TCP"
  listener_port          = "80"
  listener_description   = var.name
  load_balancer_id       = alicloud_nlb_load_balancer.default.id
  server_group_id        = alicloud_nlb_server_group.default.id
  idle_timeout           = "900"
  proxy_protocol_enabled = "true"
  cps                    = "10000"
  mss                    = "0"
}
```

## Argument Reference

The following arguments are supported:
* `alpn_enabled` - (Optional, Computed) Specifies whether to enable ALPN. Valid values:
  - `true`: Enable.
  - `false` (default): Disable.
* `alpn_policy` - (Optional) The ALPN policy. Valid values:
  - `HTTP1Only`: Only HTTP/1.x is negotiated. Priority: HTTP/1.1 > HTTP/1.0.
  - `HTTP2Only`: Only HTTP/2.0 is negotiated.
  - `HTTP2Optional`: HTTP/1.x is preferred, but HTTP/2.0 is also accepted. Priority: HTTP/1.1 > HTTP/1.0 > HTTP/2.0.
  - `HTTP2Preferred`: HTTP/2.0 is preferred, but HTTP/1.x is also accepted. Priority: HTTP/2.0 > HTTP/1.1 > HTTP/1.0.

-> **NOTE:**  You must configure this parameter when `AlpnEnabled` is set to true.

* `ca_certificate_ids` - (Optional, List) The information about the Certificate Authority (CA) certificate list. You can add only one CA certificate.

-> **NOTE:**  This parameter takes effect only for TCPSSL listeners.

* `ca_enabled` - (Optional, Computed) Specifies whether to enable mutual authentication. Valid values:
  - `true`: enables mutual authentication.
  - `false`: disables mutual authentication.
* `certificate_ids` - (Optional, List) The information about the server certificate list. You can add only one server certificate.

-> **NOTE:**  This parameter takes effect only for TCPSSL listeners.

* `cps` - (Optional, Int) The limit on new connections per second processed by the listener in each zone (VIP). Valid values: `0` to `1000000`. `0` indicates that no limit is applied.
* `end_port` - (Optional, ForceNew, Int) The end port for full-port listening. Valid values: `1` to `65535`.
The value of the end port must be greater than that of the start port.

-> **NOTE:**  This parameter is required when `ListenerPort` is set to `0`.

* `idle_timeout` - (Optional, Computed, Int) The idle connection timeout period. Unit: seconds.
  - When the listener protocol is `TCP` or `TCPSSL`, the valid values of the idle connection timeout period range from `10` to `900`. Default value: `900`.
  - When the listener protocol is `UDP`, the valid values of the idle connection timeout period range from `10` to `20`. Default value: `20`.
* `listener_description` - (Optional) The custom listener name.
The name must be 2 to 256 characters in length and can contain Chinese characters, English letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
* `listener_port` - (Required, ForceNew, Int) The listening port. Valid values: `0` to `65535`.

  - *0**: indicates that the all-port listening feature is used. If this parameter is set to `0`, you must configure `StartPort` and `EndPort`.
* `listener_protocol` - (Required, ForceNew) The listener protocol. Valid values: `TCP`, `UDP`, and `TCPSSL`.
* `load_balancer_id` - (Required, ForceNew) The ID of the Network Load Balancer instance.
* `mss` - (Optional, Int) The maximum segment size (MSS) of TCP packets. Unit: bytes. Valid values: `0` to `1500`. A value of `0` indicates that the MSS value of user TCP packets is not modified.

-> **NOTE:**  This field is supported only by TCP and TCPSSL listeners.

* `proxy_protocol_config` - (Optional, Computed, Set, Available since v1.243.0) The configuration for carrying VpcId, PrivateLinkEpId, and PrivateLinkEpsId information to backend servers through Proxy Protocol. See [`proxy_protocol_config`](#proxy_protocol_config) below.
* `proxy_protocol_enabled` - (Optional, Computed) Specifies whether to enable Proxy Protocol to pass client source IP addresses to backend servers. Valid values:
  - `true`: Enable.
  - `false`: Disable.
* `sec_sensor_enabled` - (Optional, Computed) Specifies whether to enable second-level monitoring. Valid values:
  - `true`: Enable second-level monitoring.
  - `false` (default): Disable second-level monitoring.
* `security_policy_id` - (Optional, Computed) The ID of the security policy. System security policies and custom security policies are supported.
  - System policy valid values: `tls_cipher_policy_1_0` (default), `tls_cipher_policy_1_1`, `tls_cipher_policy_1_2`, `tls_cipher_policy_1_2_strict`, or `tls_cipher_policy_1_2_strict_with_1_3`.
  - Custom security policy: Enter a custom security policy ID.
    - To create a custom security policy, see [CreateSecurityPolicy](https://help.aliyun.com/document_detail/445901.html).
    - To query security policies, see [ListSecurityPolicy](https://help.aliyun.com/document_detail/445900.html).

-> **NOTE:**  This parameter takes effect only for TCPSSL listeners.

* `server_group_id` - (Optional) The ID of the server group.

-> **NOTE:**  - When `ListenerProtocol` is set to `TCP`, the listener supports server groups whose backend protocol is `TCP` or `TCP_UDP`, but does not support server groups whose backend protocol is `UDP`.

-> **NOTE:**  - When `ListenerProtocol` is set to `UDP`, the listener supports server groups whose backend protocol is `UDP` or `TCP_UDP`, but does not support server groups whose backend protocol is `TCP`.

-> **NOTE:**  - When `ListenerProtocol` is set to `TCPSSL`, the listener supports server groups whose backend protocol is `TCP` and for which **client address preservation is disabled**. It does not support server groups whose backend protocol is `TCP` and for which **client address preservation is enabled**, or server groups whose backend protocol is `UDP` or `TCP_UDP`.

* `server_group_tuples` - (Optional, List, Available since v1.287.0) The list of multiple server groups. See [`server_group_tuples`](#server_group_tuples) below.
* `start_port` - (Optional, ForceNew, Int) The start port for all-port listening. Valid values: `1` to `65535`.

-> **NOTE:**  This parameter is required when `ListenerPort` is set to `0`.

* `status` - (Optional, Computed) The current status of the listener. Valid values:
  - `Provisioning`: The listener is being created.
  - `Running`: The listener is running.
  - `Configuring`: The listener is being configured.
  - `Stopping`: The listener is being stopped.
  - `Stopped`: The listener is stopped.
  - `Starting`: The listener is being started.
  - `Deleting`: The listener is being deleted.
  - `Deleted`: The listener is deleted.
* `tags` - (Optional, Map) A collection of resources and their tags, which includes information such as resource IDs, resource types, and tag key-value pairs.

### `proxy_protocol_config`

The proxy_protocol_config supports the following:
* `proxy_protocol_config_private_link_ep_id_enabled` - (Optional, Computed, Available since v1.243.0) Specifies whether to enable the Proxy Protocol to pass Ppv2PrivateLinkEpId to backend servers. Valid values:
  - `true`: Enabled.
  - `false` (default): Disabled.
* `proxy_protocol_config_private_link_eps_id_enabled` - (Optional, Available since v1.243.0) Specifies whether to enable carrying PrivateLinkEpsId to backend servers through Proxy Protocol. Valid values:
  - `true`: Enabled.
  - `false` (default): Disabled.
* `proxy_protocol_config_vpc_id_enabled` - (Optional) Specifies whether to enable the Proxy Protocol to pass VpcId to backend servers. Valid values:
  - `true`: Enabled.
  - `false` (default): Disabled.

### `server_group_tuples`

The server_group_tuples supports the following:
* `server_group_id` - (Optional, Available since v1.287.0) The server group ID.
* `weight` - (Optional, Int, Available since v1.287.0) The weight.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `region_id` - The region ID of the Network Load Balancer (NLB) instance.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 30 mins) Used when create the Listener.
* `delete` - (Defaults to 30 mins) Used when delete the Listener.
* `update` - (Defaults to 30 mins) Used when update the Listener.

## Import

Network Load Balancer (NLB) Listener can be imported using the id, e.g.

```shell
$ terraform import alicloud_nlb_listener.example <listener_id>
```