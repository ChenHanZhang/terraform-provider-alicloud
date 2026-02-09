---
subcategory: "Anti-DDoS Pro (DdosCoo)"
layout: "alicloud"
page_title: "Alicloud: alicloud_ddos_coo_web_cc_rule"
description: |-
  Provides a Alicloud Ddos Coo Web Cc Rule resource.
---

# alicloud_ddos_coo_web_cc_rule

Provides a Ddos Coo Web Cc Rule resource.

CC frequency control rules.

For information about Ddos Coo Web Cc Rule and how to use it, see [What is Web Cc Rule](https://next.api.alibabacloud.com/document/ddoscoo/2020-01-01/ConfigWebCCRuleV2).

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

resource "alicloud_ddoscoo_instance" "default8CmMdd" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "example"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "0"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ssl_certificates_service_certificate" "defaultqNOeDu" {
  cert             = <<EOF
-----BEGIN CERTIFICATE-----
MIID4TCCAsmgAwIBAgIRALw5sXZD1UHDhmh/t2VTQw4wDQYJKoZIhvcNAQELBQAw
XjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU
ZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w
HhcNMjUwOTI2MDI1NDU1WhcNMjYwOTI2MDI1NDU1WjAlMQswCQYDVQQGEwJDTjEW
MBQGA1UEAxMNdGVzdGxkLnFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBAMrfIvzgwhQegAeYFBRIR2LIwWT3cnKA7dLTmQUmusSqmx/AgA1ctaw8
/BUaRCCjamkYKnbDqBSYPUGMicLUVTbgiXuupoFwGBbkHN9AyetUiV86A8hebDi0
Hp3mK6AwIX432mb8nKiM3GCjVflJRt//xOybCpkqLyXFmOQxXUunZJEUUic+JHWa
bVlBxFzd4CDnBRrw0q0JPti0322TuL9HjiGkiJp2BvnMH++qtlTjwzOxMvTYeiz8
+E+yl0kzCW+bmMZK+t39nWrX57MvggNP8KsT6YCHcGgbSPQPcfx0kBL2IAU7eWbX
Xgpat3v/zRXxcIPjvg1EBmcw2rxu8dMCAwEAAaOB0jCBzzAOBgNVHQ8BAf8EBAMC
BaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMB8GA1UdIwQYMBaAFCiB
JgXRNBo/wXMPu5PPFRw/A79/MGMGCCsGAQUFBwEBBFcwVTAhBggrBgEFBQcwAYYV
aHR0cDovL29jc3AubXlzc2wuY29tMDAGCCsGAQUFBzAChiRodHRwOi8vY2EubXlz
c2wuY29tL215c3NsdGVzdHJzYS5jcnQwGAYDVR0RBBEwD4INdGVzdGxkLnFxLmNv
bTANBgkqhkiG9w0BAQsFAAOCAQEAKtDeQoQtloF6mvMOC0AYwJ2as7XyxfKKoqPs
dW7VHuASnB5AUeSmqPz3H8+qS7IX9VZDmTr2JxPRsJ+eYXMMI3UUlHUik0BcMt3Y
JfsV6nRgKm8JwktSHCsyVPDYU3zCO6KF1tUVKa18l61Twq81+gwX6jlmRy45/kPe
6yPUYA5FrNWc5ZWs4LcEM9F5L9xkhJVS8uICU09k8pwYsmU87z5mHaRxxSYjCoF2
gUrJjy6iWYfSJRWbDDA4p+BVZMuK3bGV4K7bS2lKjUPz7EZSUKQbWrzCMEOr7E8Y
9sFvHi49Blv8zllUS3clDdsP7nYPtU0hNysA9m9+eKkggFCo9Q==
-----END CERTIFICATE-----
EOF
  certificate_name = "1770603014"
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAyt8i/ODCFB6AB5gUFEhHYsjBZPdycoDt0tOZBSa6xKqbH8CA
DVy1rDz8FRpEIKNqaRgqdsOoFJg9QYyJwtRVNuCJe66mgXAYFuQc30DJ61SJXzoD
yF5sOLQeneYroDAhfjfaZvycqIzcYKNV+UlG3//E7JsKmSovJcWY5DFdS6dkkRRS
Jz4kdZptWUHEXN3gIOcFGvDSrQk+2LTfbZO4v0eOIaSImnYG+cwf76q2VOPDM7Ey
9Nh6LPz4T7KXSTMJb5uYxkr63f2datfnsy+CA0/wqxPpgIdwaBtI9A9x/HSQEvYg
BTt5ZtdeClq3e//NFfFwg+O+DUQGZzDavG7x0wIDAQABAoIBAF3lJb/t1NXUAgTB
xfVXOLcHXL866d0GQEyWQ9oHAWV54v8wrPPCR5F2zmOD+ykyCVOn7Ct4xif2CE/4
2F/0v5X9GZTFkmoRNA0QOi64QVxqjYQmCU8pKKSb9Rm2yNVEwZO2DR8iZu15+Ju/
rVCKMkQFkKLD5YVbaWPtjyR6lognFrzkehASnmk2xGbqPjpr6wXWVVQ/MkJd4Nwt
SFzF9veBZRjSmxFCl9yowE1WdsEmvkzQyX0bI8u/pY3z4hj9EuffQz9/zzL5xVS3
vNTN0CuRTyOnTHaB+6K/SIh2nqkJRwAb9UPTokAgUnO/EhN4W1CaGqimFCxxgSVT
yt3c/LECgYEAyvWUVtPAbDbrQhJ7L7dM5eR98xbhy7ZF4n0djJ6w6qPu/FSLTOjq
j2REtBzDn3xDo6Z/5U0vJG9P7RpleZ3s8g0vF4zN8Lu6EkDZfsRriq2sW9Nl+f0H
3zNocOVafQXyrrM3WytBbsdHwRi0oWs0z22h1pL2SZSn4TWnzgjrdUMCgYEA/+Ox
LXvAaCgoLYUZHuBhyoPmocl7a14MUzii8dnxjCIaawS7YjSXoc5mSoYdft6wgm+U
sjqGgkqTFasiSJUN+d7367Liy+aPnggCqK6rE9Hgi2T7sdQv+XtU+FMtd7cPzbkQ
WVsMHVpr7dA+bzO+IE6wgTxX8g8soSrCO1aUgDECgYAlLlQci/JGYNE8a3JRzXyy
6OcB74Ex9pRa55zQNAopEhsn8r0KO+ksl6vWayaTQwqJImlvsnIedJ3py9onK31K
4otr/wmDPoDZ8zNk+8rPvv1CXTnjUC1vAFXzyLCJEtvgkUhk1UnJZ4yHnWUJ5T/p
eCYbzxR7alZO9atmHVA1TQKBgD8Lx3SQX/iJpFSKzYSo/g8abnGAJdNvSZQbiTIT
Y4sGQAIMGWr50D5CjztfTdcbYNvSSA2dk9R4MUMOdhTx/I6K3ASLf4uDU/E4wgbz
eh0ZAbz2dXj78ZIDTA0e2T38sX0bUqbhYtu8koj2XNujIP3uxVgiGPz/thxDX7Wl
AFORAoGAEh5MIb1j3Z8n2isB9AxP+Ju6Q38AueX0iKvvjFUQiqzQKgaa16VajPw1
YDMn3aoIIA9UyghkSmKdBWXAgpRWqRLqakbN58PMdtmDRhp2qqz7xljpOBSKRs3t
G5w8hpXVQAflI/SUAKdnoQdHoezMX8WWQzQAlOlh4lfTKAPOh8E=
-----END RSA PRIVATE KEY-----
EOF
}

resource "alicloud_ddoscoo_domain_resource" "defaultuvS9OT" {
  rs_type      = "0"
  ocsp_enabled = false
  proxy_types {
    proxy_ports = ["80"]
    proxy_type  = "http"
  }
  proxy_types {
    proxy_ports = ["443"]
    proxy_type  = "https"
  }
  proxy_types {
    proxy_ports = ["80"]
    proxy_type  = "websocket"
  }
  real_servers     = ["1.1.1.1", "2.2.2.2", "3.3.3.3"]
  domain           = "exampleld.qq.com"
  instance_ids     = ["${alicloud_ddoscoo_instance.default8CmMdd.id}"]
  https_ext        = jsonencode({ "Https2http" : 0, "Http2" : 1, "Http2https" : 0 })
  cert_identifier  = "${alicloud_ssl_certificates_service_certificate.defaultqNOeDu.id}-cn-hangzhou"
  custom_headers   = jsonencode({ "3444" : "5555", "666" : "$ReqClientPort", "77777" : "$ReqClientIP" })
  white_list       = ["1.1.1.1"]
  ai_template      = "level30"
  bw_list_enable   = "1"
  ai_mode          = "defense"
  black_list       = ["2.2.2.2"]
  cc_global_switch = "open"
}


resource "alicloud_ddos_coo_web_cc_rule" "default" {
  rule_detail {
    action = "block"
    rate_limit {
      interval  = "11"
      threshold = "2"
      ttl       = "840"
      target    = "header"
      sub_key   = "33"
    }
    condition {
      match_method = "belong"
      field        = "ip"
      content      = "1.1.1.1"
    }
    condition {
      match_method = "contain"
      field        = "uri"
      content      = "/a"
    }
    condition {
      match_method = "contain"
      field        = "header"
      header_name  = "123"
      content      = "1234"
    }
    statistics {
      mode        = "distinct"
      field       = "header"
      header_name = "12"
    }
    status_code {
      enabled         = true
      code            = "100"
      use_ratio       = false
      count_threshold = "2"
      ratio_threshold = "5"
    }
  }
  name   = "ld_817"
  domain = alicloud_ddoscoo_domain_resource.defaultuvS9OT.id
}
```

## Argument Reference

The following arguments are supported:
* `domain` - (Required, ForceNew) The domain name of the website service.  

-> **NOTE:**  The domain name must already have website service forwarding rules configured. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.  

* `name` - (Required, ForceNew) Rule name.
* `rule_detail` - (Required, Set) Rule details.   See [`rule_detail`](#rule_detail) below.

### `rule_detail`

The rule_detail supports the following:
* `action` - (Required) The action to take when a match occurs. Valid values:
  - `accept`: Allow
  - `block`: Block
  - `challenge`: Challenge
  - `watch`: Monitor
* `condition` - (Required, List) List of matching conditions.   See [`condition`](#rule_detail-condition) below.
* `rate_limit` - (Required, Set) Rate limiting statistics. See [`rate_limit`](#rule_detail-rate_limit) below.
* `statistics` - (Optional, Set) Deduplicated statistics. This parameter is optional. If omitted, deduplication is not applied. See [`statistics`](#rule_detail-statistics) below.
* `status_code` - (Optional, Set) The HTTP status code. See [`status_code`](#rule_detail-status_code) below.

### `rule_detail-condition`

The rule_detail-condition supports the following:
* `content` - (Required) Matching content.
* `field` - (Required) Matching field.  
* `header_name` - (Optional) Custom HTTP header field name.

-> **NOTE:**  Valid only when `Field` is set to `header`.

* `match_method` - (Required) Matching method.  

### `rule_detail-rate_limit`

The rule_detail-rate_limit supports the following:
* `interval` - (Required, Int) Statistical interval. Unit: seconds.
* `sub_key` - (Optional) Header field name (required only when the statistic source is `header`).
* `target` - (Required) Statistic source. Valid values:
  - `ip`: Statistics are collected by IP address.
  - `header`: Statistics are collected by header.
* `threshold` - (Required, Int) The trigger threshold.
* `ttl` - (Required, Int) Block duration. Unit: seconds.

### `rule_detail-statistics`

The rule_detail-statistics supports the following:
* `field` - (Required) The statistic source. Valid values:
  - `ip`: Count by IP address
  - `header`: Count by HTTP header
  - `uri`: Count by URI
* `header_name` - (Optional) Set this parameter only when the statistic source is `header`.
* `mode` - (Required) The deduplication mode. Valid values:
  - `count`: No deduplication
  - `distinct`: Deduplicated count

### `rule_detail-status_code`

The rule_detail-status_code supports the following:
* `code` - (Required, Int) Status code. The value range is `100` to `599`:
  - `200`: Indicates success.
  - Others: Indicate error codes.
* `count_threshold` - (Optional, Int) When the ratio is not used, the enforcement action is triggered only when the corresponding status code reaches `CountThreshold`. The value range is `2` to `50000`.
* `enabled` - (Required) Whether the rule is enabled. Valid values:
  - `true`: Enabled.
  - `false`: Disabled.
* `ratio_threshold` - (Optional, Int) When the ratio is used, the enforcement action is triggered only when the corresponding status code reaches `RatioThreshold`. The value range is `1` to `100`.
* `use_ratio` - (Required) Whether to use a ratio:
  - `true`: Use ratio.
  - `false`: Do not use ratio.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<domain>:<name>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Web Cc Rule.
* `delete` - (Defaults to 5 mins) Used when delete the Web Cc Rule.
* `update` - (Defaults to 5 mins) Used when update the Web Cc Rule.

## Import

Ddos Coo Web Cc Rule can be imported using the id, e.g.

```shell
$ terraform import alicloud_ddos_coo_web_cc_rule.example <domain>:<name>
```