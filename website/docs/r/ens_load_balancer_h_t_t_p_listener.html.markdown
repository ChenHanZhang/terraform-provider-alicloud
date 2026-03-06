---
subcategory: "ENS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ens_load_balancer_h_t_t_p_listener"
description: |-
  Provides a Alicloud ENS Load Balancer H T T P Listener resource.
---

# alicloud_ens_load_balancer_h_t_t_p_listener

Provides a ENS Load Balancer H T T P Listener resource.

Http Listener for load balancing.

For information about ENS Load Balancer H T T P Listener and how to use it, see [What is Load Balancer H T T P Listener](https://next.api.alibabacloud.com/document/Ens/2017-11-10/CreateLoadBalancerHTTPListener).

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
  default = "cn-hangzhou-47"
}

resource "alicloud_ens_network" "defaultgkJXuV" {
  network_name  = "example用例-http监听示例"
  cidr_block    = "10.0.0.0/8"
  ens_region_id = var.ens_region_id
}

resource "alicloud_ens_vswitch" "defaultQy3IhQ" {
  cidr_block    = "10.0.8.0/24"
  vswitch_name  = "example用例-负载均衡-http监听"
  ens_region_id = alicloud_ens_network.defaultgkJXuV.ens_region_id
  network_id    = alicloud_ens_network.defaultgkJXuV.id
}

resource "alicloud_ens_load_balancer" "defaultmaG48i" {
  load_balancer_name = "example用例-http监听"
  vswitch_id         = alicloud_ens_vswitch.defaultQy3IhQ.id
  payment_type       = "PayAsYouGo"
  ens_region_id      = alicloud_ens_vswitch.defaultQy3IhQ.ens_region_id
  network_id         = alicloud_ens_vswitch.defaultQy3IhQ.network_id
  load_balancer_spec = "elb.s1.small"
}


resource "alicloud_ens_load_balancer_h_t_t_p_listener" "default" {
  listener_port             = "8082"
  request_timeout           = "60"
  health_check_interval     = "5"
  description               = "exampleD5esc1"
  unhealthy_threshold       = "5"
  scheduler                 = "wrr"
  health_check_uri          = "/example1.html"
  health_check              = "on"
  idle_timeout              = "5"
  load_balancer_id          = alicloud_ens_load_balancer.defaultmaG48i.id
  health_check_timeout      = "5"
  health_check_connect_port = "8082"
  health_check_method       = "head"
  healthy_threshold         = "5"
  x_forwarded_for           = "off"
  health_check_domain       = "example1.com"
  health_check_http_code    = "http_2xx"
  status                    = "Stopped"
  backend_server_port       = "8082"
}
```

## Argument Reference

The following arguments are supported:
* `backend_server_port` - (Optional, ForceNew, Int) The port used by the backend of the SLB instance. Valid values: `1` to **65535 * *.
* `description` - (Optional) Sets the description of the listener. The length is limited to `1` to `80` characters.

-> **NOTE:**  cannot start with 'http:// 'and 'https.

* `health_check` - (Required) Whether to enable health check. Value:
  - `on`: on.
  - `off`: closed.
* `health_check_connect_port` - (Optional, Int) The port used for health check. Valid values: `1` to **65535 * *. If this parameter is not set, the backend service port (BackendServerPort) is used.

-> **NOTE:**  is valid only when the HealthCheck value is on.

* `health_check_domain` - (Optional) The domain name used for health checks.

-> **NOTE:**  is valid only when the HealthCheck value is on.

* `health_check_http_code` - (Optional) The HTTP status code when the health check is normal. Value:
  - `http_2xx` (default).
  - **http_3xx * *.
  - **http_4xx * *.
  - **http_5xx * *.

-> **NOTE:**  is valid only when the HealthCheck value is on.

* `health_check_interval` - (Optional, Int) The interval between health checks. Value: `1` to `50`, in seconds.

-> **NOTE:**  is valid only when the HealthCheck value is on.

* `health_check_method` - (Optional) The health check method for listening to the health check of the HTTP type. Value:
  - `head`: Only the first part of the page is requested.
  - `get`: requests the specified page information and returns the entity body.

-> **NOTE:**  is valid only when the HealthCheck value is on.

* `health_check_timeout` - (Optional, Int) The amount of time to wait to receive a response from the health check. If the backend ECS instances do not respond within the specified time, the health check fails.
  - Default value: 5 seconds.
  - Value: `1` ~ **300 * *.
  - Unit: seconds.

-> **NOTE:** - is only valid when the HealthCheck value is on.
  - If the value of HealthCHeckTimeout is less than the value of HealthCheckInterval, HealthCHeckTimeout is invalid and the timeout period is the value of HealthCheckInterval.
* `health_check_uri` - (Optional) The URI used for health check. Length limit is `1` ~ **80 * *.

-> **NOTE:**  - URL cannot have only '/', but must start.
The-HealthCheck value is only valid when it is on.
* `healthy_threshold` - (Optional, Int) After the number of consecutive successful health checks, the health check status of the backend server is determined from fail (the backend server is unreachable) to success (the backend server is reachable). Value: `2` ~ **10 * *.

-> **NOTE:**  is valid only when the HealthCheck value is on.

* `idle_timeout` - (Optional, Int) The connection idle timeout period. Default value: 15 seconds. Valid values: `1` to `60`. Unit: seconds.

-> **NOTE:**  If there is no access request within the timeout period, load balancing will temporarily interrupt the current connection until the next request comes and re-establishes a new connection.

* `listener_port` - (Required, ForceNew, Int) The port used by the front end of the Server Load Balancer instance. Value: `1` ~ **65535 * *.
* `load_balancer_id` - (Required, ForceNew) The ID of the load balancing instance.
* `request_timeout` - (Optional, Int) The request timeout period. Default value: 60 seconds. Value: `1` to `180`. Unit: seconds.

-> **NOTE:**  If the backend server fails to respond within the timeout period, the server load balancer will stop waiting and return the HTTP 504 error code to the client.

* `scheduler` - (Optional) Scheduling algorithm. Value:
  - `wrr`: The higher the weight value, the higher the number of times (probability) that the backend server is polled.
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
* `unhealthy_threshold` - (Optional, Int) The number of consecutive health check failures that determine the health check status of the backend server from success (the backend server is reachable) to fail (the backend server is unreachable). Value: `2` ~ **10 * *.

-> **NOTE:**  is valid only when the HealthCheck value is on.

* `x_forwarded_for` - (Optional) Whether to enable the X-Forwarded-For field to obtain the real IP address of the visitor. Value:
  - `on`: on.
  - `off` (default): off.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<load_balancer_id>:<listener_port>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Load Balancer H T T P Listener.
* `delete` - (Defaults to 5 mins) Used when delete the Load Balancer H T T P Listener.
* `update` - (Defaults to 5 mins) Used when update the Load Balancer H T T P Listener.

## Import

ENS Load Balancer H T T P Listener can be imported using the id, e.g.

```shell
$ terraform import alicloud_ens_load_balancer_h_t_t_p_listener.example <load_balancer_id>:<listener_port>
```