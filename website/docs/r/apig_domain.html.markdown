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

-> **NOTE:** Available since v1.286.0.

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
* `ca_cert_identifier` - (Optional) CA certificate identifier
* `cert_identifier` - (Optional) The certificate identifier.
* `client_ca_cert` - (Optional) client CA certificate
* `domain_name` - (Required, ForceNew) Domain name.
* `domain_scope` - (Optional) domain scope
* `force_https` - (Optional) Specifies whether to enable forced HTTPS redirection when the protocol is set to HTTPS. This parameter is required if the protocol is HTTPS.
* `gateway_type` - (Optional) Gateway type.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `http2_option` - (Optional) HTTP/2 settings.
* `m_tls_enabled` - (Optional) Whether to enable mTLS mutual authentication
* `protocol` - (Required) The protocol types supported by the domain.
  - HTTP: Supports HTTP only.
  - HTTPS: Supports HTTPS only.
* `resource_group_id` - (Optional, Computed) Resource group ID (https://help.aliyun.com/document_detail/151181.html).
* `tls_cipher_suites_config` - (Optional, Set) TLS cipher suites configuration. See [`tls_cipher_suites_config`](#tls_cipher_suites_config) below.
* `tls_max` - (Optional) The maximum version of the TLS protocol supported, up to TLS 1.3.
* `tls_min` - (Optional) Minimum TLS protocol version. TLS 1.0 is the minimum supported version.

### `tls_cipher_suites_config`

The tls_cipher_suites_config supports the following:
* `config_type` - (Optional) The configuration type, which can be Default or Custom.
* `tls_cipher_suite` - (Optional, List) TLS cipher suite. See [`tls_cipher_suite`](#tls_cipher_suites_config-tls_cipher_suite) below.

### `tls_cipher_suites_config-tls_cipher_suite`

The tls_cipher_suites_config-tls_cipher_suite supports the following:
* `name` - (Optional) The name of the cipher suite.
* `support_versions` - (Optional, List) support versions

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Domain.
* `delete` - (Defaults to 5 mins) Used when delete the Domain.
* `update` - (Defaults to 5 mins) Used when update the Domain.

## Import

APIG Domain can be imported using the id, e.g.

```shell
$ terraform import alicloud_apig_domain.example <domain_id>
```