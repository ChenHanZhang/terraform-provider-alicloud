---
subcategory: "Cloud Native API Gateway (APIG)"
layout: "alicloud"
page_title: "Alicloud: alicloud_apig_domain"
description: |-
  Provides a Alicloud APIG Domain resource.
---

# alicloud_apig_domain

Provides a APIG Domain resource.



For information about APIG Domain and how to use it, see [What is Domain](https://next.api.alibabacloud.com/document/APIG/2024-03-27/CreateDomain).

-> **NOTE:** Available since v1.287.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_apig_domain" "default" {
  domain_name  = "example-domain-cspec-v6.example.com"
  gateway_type = "API"
  protocol     = "HTTP"
}
```

## Argument Reference

The following arguments are supported:
* `ca_cert_identifier` - (Optional) The CA certificate identifier. This parameter is optional for dedicated gateways over HTTPS. This parameter is not supported for serverless gateways. Verification is not performed for dedicated gateways over HTTP.
* `cert_identifier` - (Optional) The certificate identifier. This parameter is required for Dedicated+HTTPS and must pass validation before submission. It is not allowed for Serverless. It is not validated for Dedicated+HTTP.
* `client_ca_cert` - (Optional) The client CA certificate. This parameter is required for Dedicated+HTTPS when mTLS is enabled (MTLSEnabled=true). It is not allowed for Serverless. It is not validated for Dedicated+HTTP.
* `domain_name` - (Required) The domain name. The name must be 1 to 128 characters in length, such as abc.com.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `domain_scope` - (Optional, Computed) The domain name scope. Valid values: Dedicated (dedicated gateway domain name) and Serverless (serverless gateway domain name). Default value: Dedicated.
* `force_https` - (Optional) Sets the HTTPS protocol type and specifies whether to enable forced HTTPS redirection. When the protocol is HTTPS, forceHttps is required.
* `gateway_type` - (Optional) The gateway type. This parameter is optional. If you do not specify this parameter, the default value API is used.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `http2_option` - (Optional) The HTTP/2 setting. Valid values: GlobalConfig (follow global configuration), Open (enable), and Close (disable). Default value: GlobalConfig. This setting is supported only for HTTPS domains in the Dedicated scope.
* `m_tls_enabled` - (Optional) Specifies whether to enable mTLS mutual authentication. This parameter is optional for dedicated gateways over HTTPS. If this parameter is set to true, ClientCACert is required. This parameter is not supported for serverless gateways.
* `protocol` - (Optional, Computed) The protocol type used by the domain name. Valid values: HTTP and HTTPS. This parameter is required for the Dedicated scope and is not allowed for the Serverless scope.
* `resource_group_id` - (Optional, Computed) [Resource Group ID](https://help.aliyun.com/document_detail/151181.html).
* `tls_cipher_suites_config` - (Optional, Computed, Set) The TLS cipher suite configuration, including the configuration type, cipher suite names, and supported TLS versions. This configuration is supported only for HTTPS domain names with a Dedicated scope. See [`tls_cipher_suites_config`](#tls_cipher_suites_config) below.
* `tls_max` - (Optional) The maximum TLS protocol version. The maximum supported version is TLS 1.3.
* `tls_min` - (Optional) The minimum TLS protocol version. The minimum supported version is TLS 1.0.

### `tls_cipher_suites_config`

The tls_cipher_suites_config supports the following:
* `config_type` - (Optional) The configuration type: Default or Custom.
* `tls_cipher_suite` - (Optional, List) The TLS cipher suite. See [`tls_cipher_suite`](#tls_cipher_suites_config-tls_cipher_suite) below.

### `tls_cipher_suites_config-tls_cipher_suite`

The tls_cipher_suites_config-tls_cipher_suite supports the following:
* `name` - (Optional) The name of the TLS cipher suite, such as ECDHE-ECDSA-AES256-GCM-SHA384.
* `support_versions` - (Optional, List) The supported TLS versions.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Domain.
* `delete` - (Defaults to 5 mins) Used when delete the Domain.
* `update` - (Defaults to 6 mins) Used when update the Domain.

## Import

APIG Domain can be imported using the id, e.g.

```shell
$ terraform import alicloud_apig_domain.example <domain_id>
```