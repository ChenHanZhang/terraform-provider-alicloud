---
subcategory: "Anti-DDoS Pro (DdosCoo)"
layout: "alicloud"
page_title: "Alicloud: alicloud_ddoscoo_domain_resource"
description: |-
  Provides a Alicloud Ddos Coo Domain Resource resource.
---

# alicloud_ddoscoo_domain_resource

Provides a Ddos Coo Domain Resource resource.

Domain name resource.

For information about Ddos Coo Domain Resource and how to use it, see [What is Domain Resource](https://www.alibabacloud.com/help/en/anti-ddos/anti-ddos-pro-and-premium/developer-reference/api-ddoscoo-2020-01-01-createdomainresource).

-> **NOTE:** Available since v1.123.0.

## Example Usage

Basic Usage

```terraform
provider "alicloud" {
  region = "cn-hangzhou"
}

variable "name" {
  default = "tf-example"
}
variable "domain" {
  default = "tf-example.alibaba.com"
}

resource "alicloud_ddoscoo_instance" "default" {
  name              = var.name
  bandwidth         = "30"
  base_bandwidth    = "30"
  service_bandwidth = "100"
  port_count        = "50"
  domain_count      = "50"
  period            = "1"
  product_type      = "ddoscoo"
}

resource "alicloud_ddoscoo_domain_resource" "default" {
  domain       = var.domain
  rs_type      = 0
  instance_ids = [alicloud_ddoscoo_instance.default.id]
  real_servers = ["177.167.32.11"]
  https_ext    = <<EOF
    {
    "Http2": 1,
    "Http2https": 0,
    "Https2http": 0
  }
  EOF
  proxy_types {
    proxy_ports = [443]
    proxy_type  = "https"
  }
}
```

## Argument Reference

The following arguments are supported:
* `ai_mode` - (Optional, Computed, Available since v1.270.0) The mode of AI-powered protection. Valid values:
  - `watch`: Alert mode
  - `defense`: Protection mode
* `ai_template` - (Optional, Computed, Available since v1.270.0) The AI-based intelligent protection level. Valid values:  
  - `level30`: loose  
  - `level60`: normal  
  - `level90`: strict
* `black_list` - (Optional, List, Available since v1.270.0) The IP blacklist for the current domain.  
* `bw_list_enable` - (Optional, Computed, Int, Available since v1.270.0) The status of the allowlist and blocklist switch. Valid values:
  - `0`: Disabled
  - `1`: Enabled
* `cc_global_switch` - (Optional, Computed, Available since v1.270.0) The status of the custom and precise CC protection switch. Valid values:  
  - `open`: enabled  
  - `close`: disabled  
* `cert` - (Optional, Available since v1.231.0) Public key of the certificate to be associated. This parameter must be used together with CertName and Key.  

-> **NOTE:**  After you specify CertName, Cert, and Key, you do not need to specify CertId.  

* `cert_identifier` - (Optional, Available since v1.231.0) Global certificate ID, which is the certificate ID appended with "-cn-hangzhou". For example, if the certificate ID is 123, then CertIdentifier is "123-cn-hangzhou".  

-> **NOTE:**  You cannot specify both CertIdentifier and CertId at the same time.  

* `cert_name` - (Optional, Computed, Available since v1.231.0) Certificate name.

-> **NOTE:**  The name of the certificate to be associated. This parameter must be used together with Cert and Key.

* `cert_region` - (Optional, Available since v1.231.0) The region where the certificate is located. Supported regions are **cn-hangzhou** and **ap-southeast-1**. The default value is **cn-hangzhou**.
* `custom_ciphers` - (Optional, List, Available since v1.281.0) List of custom cipher suites.
* `custom_headers` - (Optional, Computed, JsonString, Available since v1.261.0) Key-value pairs for custom headers. The key is the header name and the value is the corresponding value. You can configure up to five key-value pairs, with a total length of no more than 200 characters.

Notes:
  - Do not use the following default fields as custom headers:
  - X-Forwarded-ClientSrcPort: Used by default to obtain the client port accessing the Anti-DDoS Layer 7 engine.
  - X-Forwarded-ProxyPort: Used by default to obtain the listener port accessing the Anti-DDoS Layer 7 engine.
  - X-Forwarded-For: Used by default to obtain the client IP address accessing the Anti-DDoS Layer 7 engine.
  - Do not use standard HTTP header fields (such as host, user-agent, connection, upgrade) or commonly used custom HTTP header fields (such as x-real-ip, x-true-ip, x-client-ip, web-server-type, wl-proxy-client-ip, eagleeye-rpcid, eagleeye-traceid, x-forwarded-cluster, x-forwarded-proto), as doing so will overwrite the original header field values in the request.

-> **NOTE:**  If you set the key to X-Forwarded-ClientSrcPort to obtain the client port accessing the Anti-DDoS Layer 7 engine, set the value to "".

* `domain` - (Required, ForceNew) The domain name of the website to query.

-> **NOTE:**  The domain name must already have website traffic forwarding rules configured. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names that have website traffic forwarding rules configured.

* `https_ext` - (Optional, Computed, JsonString) Advanced HTTPS settings. These settings take effect only when the website protocol supports HTTPS (`ProxyType` includes `https`). The settings are provided as a JSON-formatted string with the following fields:  
  - `Http2https`: Integer | Optional | Specifies whether to enable forced redirection to HTTPS. Valid values: `0` (disabled) | `1` (enabled). Disabled by default.  

This setting applies when your website supports both HTTP and HTTPS protocols. After this setting is enabled, all HTTP requests are forcibly redirected to HTTPS requests, and the default destination port is 443.  
  - `Https2http`: Integer | Optional | Specifies whether to enable HTTP back-to-origin for HTTPS requests. Valid values: `0` (disabled) | `1` (enabled). Disabled by default.  

This setting applies when your origin server does not support HTTPS back-to-origin. After this setting is enabled, all HTTPS requests use HTTP for back-to-origin (WebSocket requests use WebSocket for back-to-origin), and the default back-to-origin port is 80.  
  - `Http2`: Integer | Optional | Specifies whether to enable HTTP/2 support. Valid values: `0` (disabled) | `1` (enabled). Disabled by default.  

After this setting is enabled, the protocol version is HTTP/2.
* `instance_ids` - (Required, List) List of IDs of Anti-DDoS Pro/Premium instances associated with the domain name.  
* `key` - (Optional) The private key of the certificate to be associated. This parameter must be used together with CertName and Cert.

-> **NOTE:**  If you specify CertName, Cert, and Key, you do not need to specify CertId.

* `ocsp_enabled` - (Optional, Available since v1.208.0) Specifies whether the Online Certificate Status Protocol (OCSP) feature is enabled. Valid values:  
  - `true`: enabled  
  - `false`: disabled
* `proxy_types` - (Required, List) The protocol and port configuration of the website.   See [`proxy_types`](#proxy_types) below.
* `real_servers` - (Required, List, Available since v1.231.0) List of origin server addresses.  
* `rs_type` - (Required, Int) The address type of the origin server to be configured. Valid values:  
  - `0`: indicates that the IP address of the origin server is configured.  
  - `1`: indicates that the domain name of the origin server is configured.  

This applies to scenarios where another proxy service (such as a WAF) is deployed between the origin server and Anti-DDoS Pro/Premium. Specifically, it refers to the forwarding address of the proxy service (for example, the CNAME address of the WAF).  
* `ssl13_enabled` - (Optional, Available since v1.281.0) Indicates whether support for the TLS 1.3 protocol is enabled. Valid values:
  - `true`: Enabled
  - `false`: Disabled
* `ssl_ciphers` - (Optional, Computed, Available since v1.281.0) The type of cipher suite. Valid values:  
  - `default`: custom cipher suite.  
  - `all`: all cipher suites.  
  - `strong`: strong cipher suite.
* `ssl_protocols` - (Optional, Computed, Available since v1.281.0) The TLS protocol version. Valid values:  
  - **tls1.0**: TLS 1.0 or later  
  - **tls1.1**: TLS 1.1 or later  
  - **tls1.2**: TLS 1.2 or later  
* `tls13_custom_ciphers` - (Optional, List, Available since v1.281.0) The list of cipher suites used for TLS 1.3.  
* `white_list` - (Optional, List) The IP whitelist for the current domain.  

### `proxy_types`

The proxy_types supports the following:
* `proxy_ports` - (Required, List) The list of ports used by the website for external services.  
* `proxy_type` - (Optional) The protocol type used by the website for external services. Valid values:  
  - `http`: HTTP protocol  
  - `https`: HTTPS protocol  
  - `websocket`: WebSocket protocol  
  - `websockets`: WebSockets protocol

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `cc_enabled` - Indicates whether rate-based control protection (CC protection) is enabled.
* `cc_rule_enabled` - Indicates whether custom rate-based protection (CC protection) is enabled.
* `cc_template` - The mode of HTTP flood protection (CC protection).
* `cname` - The DDoS protection CNAME address corresponding to the domain name.
* `http2_enable` - Indicates whether HTTP/2 is enabled.
* `policy_mode` - The type of origin load balancing algorithm.
* `proxy_enabled` - Specifies whether Anti-DDoS Pro forwards traffic for the website service.
* `punish_reason` - Reason for the penalty imposed on the domain name.
* `punish_status` - Indicates whether the domain has been penalized for policy violations.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Domain Resource.
* `delete` - (Defaults to 5 mins) Used when delete the Domain Resource.
* `update` - (Defaults to 5 mins) Used when update the Domain Resource.

## Import

Ddos Coo Domain Resource can be imported using the id, e.g.

```shell
$ terraform import alicloud_ddoscoo_domain_resource.example <domain>
```