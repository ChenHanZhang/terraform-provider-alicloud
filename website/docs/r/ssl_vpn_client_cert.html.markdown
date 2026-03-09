---
subcategory: "VPN Gateway"
layout: "alicloud"
page_title: "Alicloud: alicloud_ssl_vpn_client_cert"
description: |-
  Provides a Alicloud Vpn Gateway Ssl Vpn Client Cert resource.
---

# alicloud_ssl_vpn_client_cert

Provides a Vpn Gateway Ssl Vpn Client Cert resource.

SSL-VPN client certificate.

For information about Vpn Gateway Ssl Vpn Client Cert and how to use it, see [What is Ssl Vpn Client Cert](https://next.api.alibabacloud.com/document/Vpc/2016-04-28/CreateSslVpnClientCert).

-> **NOTE:** Available since v1.15.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

data "alicloud_vpcs" "default" {
  name_regex = "^default-NODELETING$"
  cidr_block = "172.16.0.0/16"
}

data "alicloud_vswitches" "default0" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_zones.default.ids.0
}

data "alicloud_vswitches" "default1" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_zones.default.ids.1
}

resource "alicloud_vpn_gateway" "default" {
  vpn_gateway_name             = var.name
  vpc_id                       = data.alicloud_vpcs.default.ids.0
  bandwidth                    = "10"
  enable_ssl                   = true
  description                  = var.name
  payment_type                 = "Subscription"
  vswitch_id                   = data.alicloud_vswitches.default0.ids.0
  disaster_recovery_vswitch_id = data.alicloud_vswitches.default1.ids.0
}


resource "alicloud_ssl_vpn_server" "default" {
  name           = var.name
  vpn_gateway_id = alicloud_vpn_gateway.default.id
  client_ip_pool = "192.168.0.0/16"
  local_subnet   = cidrsubnet(data.alicloud_vpcs.default.vpcs.0.cidr_block, 8, 8)
  protocol       = "UDP"
  cipher         = "AES-128-CBC"
  port           = "1194"
  compress       = "false"
}

resource "alicloud_ssl_vpn_client_cert" "default" {
  ssl_vpn_server_id = alicloud_ssl_vpn_server.default.id
  name              = var.name
}
```

## Argument Reference

The following arguments are supported:
* `ssl_vpn_client_cert_name` - (Optional, Available since v1.273.0) The name of the client certificate.
* `ssl_vpn_server_id` - (Required, ForceNew) The ID of the SSL server.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `ca_cert` - The CA certificate.
* `client_cert` - The client certificate.
* `client_config` - The client configuration.
* `client_key` - The client key.
* `create_time` - The time when the SSL client certificate was created.
* `region_id` - The region ID of the resource.
* `status` - The status of the client certificate.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Ssl Vpn Client Cert.
* `delete` - (Defaults to 5 mins) Used when delete the Ssl Vpn Client Cert.
* `update` - (Defaults to 5 mins) Used when update the Ssl Vpn Client Cert.

## Import

Vpn Gateway Ssl Vpn Client Cert can be imported using the id, e.g.

```shell
$ terraform import alicloud_ssl_vpn_client_cert.example <ssl_vpn_client_cert_id>
```