---
subcategory: "Alidns"
layout: "alicloud"
page_title: "Alicloud: alicloud_alidns_cloud_gtm_monitor_template"
description: |-
  Provides a Alicloud Alidns Cloud Gtm Monitor Template resource.
---

# alicloud_alidns_cloud_gtm_monitor_template

Provides a Alidns Cloud Gtm Monitor Template resource.

CloudGtm Monitor Template
.

For information about Alidns Cloud Gtm Monitor Template and how to use it, see [What is Cloud Gtm Monitor Template](https://next.api.alibabacloud.com/document/Alidns/2015-01-09/CreateCloudGtmMonitorTemplate).

-> **NOTE:** Available since v1.276.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}


resource "alicloud_alidns_cloud_gtm_monitor_template" "default" {
  ip_version = "IPv4"
  timeout    = "2000"
  isp_city_nodes {
    city_code = "357"
    isp_code  = "465"
  }
  isp_city_nodes {
    city_code = "738"
    isp_code  = "465"
  }
  evaluation_count = "1"
  protocol         = "ping"
  failure_rate     = "20"
  extend_info      = "{\"packetLossRate\":10,\"packetNum\":20}"
  name             = "template-example-1"
  interval         = "60"
  remark           = "remark"
}
```

## Argument Reference

The following arguments are supported:
* `evaluation_count` - (Required, Int) Number of retries. The system determines that the application service is abnormal only after consecutive monitoring failures occur multiple times. This prevents transient network fluctuations or other temporary issues from affecting monitoring accuracy. Valid retry counts are:
  - 1
  - 2
  - 3.
* `extend_info` - (Optional) JSON string containing extended information with required parameters for each protocol check:
  - HTTP(S):

  `host`: When performing HTTP(S) monitoring, specifies the Host field carried in the HTTP(S) request header to identify the specific website to access. The default value is the primary domain name. If the target website has special requirements for the host, modify it accordingly.

  `path`: The HTTP path used for configuring the URL path in HTTP(S) health checks. The system defaults to "/".

  `code`: When performing HTTP(S) monitoring, the system determines whether the web server is functioning normally based on the response code returned by the server. If the response code exceeds the configured alert threshold, the system considers the application service abnormal:

    - 400: Bad Request. If the HTTP(S) request carries invalid parameters, the web server returns a response code greater than 400. If the validation condition is set to "abnormal code greater than 400", ensure that an exact URL path with parameters is specified.

    - 500: Server Error. If the web server encounters certain errors, it returns an error code greater than 500. By default, the system uses error codes greater than 500 as the alert threshold.

  `sni`: Specifies whether to enable SNI (Server Name Indication), which applies only to HTTPS. SNI is an extension of the TLS protocol that allows a client to indicate the hostname it is attempting to connect to during the TLS handshake. Because the TLS handshake occurs before any HTTP data is sent, SNI enables the server to select the correct certificate to present to the client before sending it.

    - true: Enabled

    - false: Disabled

  `followRedirect`: Specifies whether to follow 3XX redirects.

    - true: If the monitoring node receives a 3XX status code (301, 302, 303, 307, or 308), it follows the redirect.

    - false: Does not follow redirects.
  - Ping:

  `packetNum`: The number of ICMP packets sent simultaneously during each ping monitoring session. Valid values are 20, 50, or 100.

  `packetLossRate`: The packet loss rate (%) that triggers an alert. During each ping monitoring session, the system calculates the packet loss rate based on the sent ICMP packets. Packet loss rate = (number of lost packets / total number of ICMP packets sent) × 100%. An alert is triggered when this rate reaches the configured threshold. Valid values are: 10, 30, 40, 80, 90, or 100.
* `failure_rate` - (Required, Int) Failure rate (%) of selected nodes, which represents the percentage of probe points with health check failures out of the total number of probe points. If the failure rate exceeds the configured threshold, the service address is marked as abnormal. Valid failure rate thresholds are:
  - 20
  - 50
  - 80
  - 100.
* `interval` - (Required) The interval (in seconds) between consecutive health probes. By default, a probe is performed every 1 minute. The minimum supported health check interval is 15 seconds (available for Flagship Edition instances).
* `ip_version` - (Required, ForceNew) Perform an exact search based on the IP address type of the probing node:
  - IPv4: Applicable when the target address type for probing is an IPv4 address.
  - IPv6: Applicable when the target address type for probing is an IPv6 address.
* `isp_city_nodes` - (Required, List) The list of monitoring nodes. For detailed information, call [ListCloudGtmMonitorNodes](https://help.aliyun.com/document_detail/2797349.html). See [`isp_city_nodes`](#isp_city_nodes) below.
* `name` - (Required) Name of the health check probe template. To help configuration personnel distinguish and remember templates easily, it is recommended that the name reflect the health check protocol used.
* `protocol` - (Required, ForceNew) Perform an exact search by monitoring protocol type:
  - ping
  - tcp
  - http
  - https.
* `remark` - (Optional) Enter the updated remark. Passing an empty value deletes the existing remark.
* `timeout` - (Required) Probe timeout duration in milliseconds. Packets that do not return within this duration are considered as health check timeouts. Valid values are:
  - 2000
  - 3000
  - 5000
  - 10000.

### `isp_city_nodes`

The isp_city_nodes supports the following:
* `city_code` - (Optional) The city code.
* `isp_code` - (Optional) The ISP code.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Cloud Gtm Monitor Template.
* `delete` - (Defaults to 5 mins) Used when delete the Cloud Gtm Monitor Template.
* `update` - (Defaults to 5 mins) Used when update the Cloud Gtm Monitor Template.

## Import

Alidns Cloud Gtm Monitor Template can be imported using the id, e.g.

```shell
$ terraform import alicloud_alidns_cloud_gtm_monitor_template.example <template_id>
```