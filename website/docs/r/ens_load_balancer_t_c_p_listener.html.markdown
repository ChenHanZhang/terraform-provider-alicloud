---
subcategory: "ENS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ens_load_balancer_t_c_p_listener"
description: |-
  Provides a Alicloud ENS Load Balancer T C P Listener resource.
---

# alicloud_ens_load_balancer_t_c_p_listener

Provides a ENS Load Balancer T C P Listener resource.

TCP listening of the load balancing instance.

For information about ENS Load Balancer T C P Listener and how to use it, see [What is Load Balancer T C P Listener](https://next.api.alibabacloud.com/document/Ens/2017-11-10/CreateLoadBalancerTCPListener).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

variable "ens_region_id" {
  default = "cn-hangzhou-44"
}

resource "alicloud_ens_network" "defaultzsNbyZ" {
  network_name  = "example用例-tcp监听示例"
  cidr_block    = "10.0.0.0/8"
  ens_region_id = var.ens_region_id
}

resource "alicloud_ens_vswitch" "defaultYueZqF" {
  cidr_block    = "10.0.8.0/24"
  ens_region_id = alicloud_ens_network.defaultzsNbyZ.ens_region_id
  network_id    = alicloud_ens_network.defaultzsNbyZ.id
}

resource "alicloud_ens_load_balancer" "defaultrlO87U" {
  load_balancer_name = "example用例-TCP监听"
  vswitch_id         = alicloud_ens_vswitch.defaultYueZqF.id
  payment_type       = "PayAsYouGo"
  ens_region_id      = alicloud_ens_vswitch.defaultYueZqF.ens_region_id
  network_id         = alicloud_ens_vswitch.defaultYueZqF.network_id
  load_balancer_spec = "elb.s1.small"
}


resource "alicloud_ens_load_balancer_t_c_p_listener" "default" {
  listener_port                = "5000"
  description                  = "example"
  health_check_interval        = "10"
  unhealthy_threshold          = "10"
  scheduler                    = "wrr"
  health_check_uri             = "/example.html"
  health_check_connect_timeout = "20"
  load_balancer_id             = alicloud_ens_load_balancer.defaultrlO87U.id
  backend_server_port          = "5000"
  persistence_timeout          = "0"
  health_check_connect_port    = "5000"
  healthy_threshold            = "10"
  health_check_domain          = "example.com"
  eip_transmit                 = "off"
  health_check_http_code       = "http_2xx"
  health_check_type            = "http"
  established_timeout          = "50"
  status                       = "Stopped"
}
```

## Argument Reference

The following arguments are supported:
* `backend_server_port` - (Optional, ForceNew, Int) The port used by the backend of the SLB instance. Valid values: `1` to **65535 * *.
* `description` - (Optional) Sets the description of the listener. The length is limited to `1` to `80` characters.

-> **NOTE:**  cannot start with 'http:// 'and 'https.

* `eip_transmit` - (Optional) Whether to enable EIP transparent transmission. Value:
  - `on`: on.
  - `off` (default): off.
* `established_timeout` - (Optional, Int) The connection timeout period. Valid values: `10` to `900`. The default value is `900`. Unit: seconds.
* `health_check_connect_port` - (Optional, Int) The port used for health check. Valid values: `1` to **65535 * *. If this parameter is not set, the backend service port (BackendServerPort) is used.
* `health_check_connect_timeout` - (Optional, Int) The amount of time to wait to receive a response from the health check. If the backend ECS instances do not respond within the specified time, the health check fails.
  - Default value: 5 seconds.
  - Value: `1` ~ **300 * *.
  - Unit: seconds.

-> **NOTE:**  If the value of HealthCheckConnectTimeout is less than the value of HealthCheckInterval, HealthCheckConnectTimeout is invalid and the timeout period is the value of HealthCheckInterval.

* `health_check_domain` - (Optional) The domain name used for health checks.
* `health_check_http_code` - (Optional) The HTTP status code when the health check is normal. Value:
  - `http_2xx` (default).
  - **http_3xx * *.
  - **http_4xx * *.
  - **http_5xx * *.

* `health_check_interval` - (Optional, Computed, Int) The interval between health checks. Value: `1` to `50`. Default value: `2`. Unit: seconds.
* `health_check_type` - (Optional) Type of health check. Value:
  - `tcp` (default).
  - **http * *.
* `health_check_uri` - (Optional) The URI used for health check. Length limit is `1` ~ **80 * *.

-> **NOTE:**  URL cannot have only '/', but must start.

* `healthy_threshold` - (Optional, Int) After the number of consecutive successful health checks, the health check status of the backend server is determined from fail (the backend server is unreachable) to success (the backend server is reachable). Value: `2` to `10`. The default value is **3 * *.
* `listener_port` - (Required, ForceNew, Int) The port used by the front end of the Server Load Balancer instance. Value: `1` ~ **65535 * *.
* `load_balancer_id` - (Required, ForceNew) The ID of the load balancing instance.
* `persistence_timeout` - (Optional, Int) The timeout period for session retention.
  - Default value: 0, which means that session persistence is turned off.
  - Value: `0` ~ **3600 * *.
  - Unit: seconds.
* `scheduler` - (Optional) Scheduling algorithm. Value:
  - `wrr` (default): The higher the weight value, the higher the number of times (probability) the backend server is polled.
  - `wlc`: In addition to polling based on the weight value set for each backend server, the actual load of the backend server (that is, the number of connections) is also considered. When the weight value is the same, the number of times (probability) that a backend server with a smaller number of current connections is polled is higher.
  - `rr`: distributes external requests to backend servers in the order they are accessed.
  - `sch`: a consistent Hash based on the source IP address. The same source IP address is dispatched to the same backend server.
  - `qch`: Consistent Hash based on QUIC Connection ID. The same QUIC Connection ID is dispatched to the same backend server.
  - `iqch`: performs consistent Hash on specific three bytes of iQUIC CID, and dispatches the same second to fourth bytes to the same backend server.
* `status` - (Optional, Computed) The current status of the listener. Value:
  - `Running`: Normal operation.
  - `Stopped`: The run is Stopped.
  - `Starting`: Starting.
  - `Configuring`: Configuring.
  - `Stopping`: The running is being stopped.
* `unhealthy_threshold` - (Optional, Int) The number of consecutive health check failures that determine the health check status of the backend server from success (the backend server is reachable) to fail (the backend server is unreachable). Value: `2` to `10`. The default value is **3 * *.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<load_balancer_id>:<listener_port>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Load Balancer T C P Listener.
* `delete` - (Defaults to 5 mins) Used when delete the Load Balancer T C P Listener.
* `update` - (Defaults to 5 mins) Used when update the Load Balancer T C P Listener.

## Import

ENS Load Balancer T C P Listener can be imported using the id, e.g.

```shell
$ terraform import alicloud_ens_load_balancer_t_c_p_listener.example <load_balancer_id>:<listener_port>
```