---
subcategory: "SSL Certificates"
layout: "alicloud"
page_title: "Alicloud: alicloud_ssl_certificates_service_pca_certificate"
description: |-
  Provides a Alicloud SSL Certificates Pca Certificate resource.
---

# alicloud_ssl_certificates_service_pca_certificate

Provides a SSL Certificates Pca Certificate resource.



For information about SSL Certificates Pca Certificate and how to use it, see [What is Pca Certificate](https://next.api.alibabacloud.com/document/cas/2020-06-30/CreateRootCACertificate).

-> **NOTE:** Available since v1.257.0.

## Example Usage

Basic Usage

```terraform
provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_ssl_certificates_service_pca_certificate" "default" {
  organization      = "a"
  years             = "1"
  locality          = "a"
  organization_unit = "a"
  state             = "a"
  country_code      = "cn"
  common_name       = "cbc.certqa.cn"
  algorithm         = "RSA_2048"
}
```

## Argument Reference

The following arguments are supported:
* `algorithm` - (Optional, ForceNew, Computed) The key algorithm type of the root CA certificate. The key algorithm is expressed using the '_< key length>' format. Value:
  - `RSA_1024`: The corresponding signature algorithm is Sha256WithRSA.
  - `RSA_2048`: The corresponding signature algorithm is Sha256WithRSA.
  - `RSA_4096`: The corresponding signature algorithm is Sha256WithRSA.
  - `ECC_256`: The signature algorithm is Sha256WithECDSA.
  - `ECC_384`: The corresponding signature algorithm is Sha256WithECDSA.
  - `ECC_512`: The signature algorithm is Sha256WithECDSA.
  - `SM2_256`: The corresponding signature algorithm is SM3WithSM2.

The encryption algorithm of the root CA certificate must be the same as the **certificate algorithm** of the private Root CA you purchased. Example: If the **certificate algorithm** selected when you purchase a private Root CA is `RSA`, the key algorithm of the root CA certificate must be **RSA\_1024**, **RSA\_2048**, or **RSA\_4096**.
* `alias_name` - (Optional, Available since v1.266.0) Alias Name.

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `common_name` - (Required, ForceNew) The common name or abbreviation of the organization. Support the use of Chinese, English characters.
* `country_code` - (Optional, ForceNew) The code of the country or region in which the organization is located, using a two-digit capital abbreviation. For example, `CN` represents China and `US` represents the United States.
For more information about the codes of different countries, see **International Code** in [Management Company Information](~~ 198289 ~~).
* `crl_day` - (Optional, ForceNew, Computed, Int, Available since v1.269.0) CRL cycle.
* `enable_crl` - (Optional, Available since v1.269.0) Whether to enable CRL.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `extended_key_usages` - (Optional, List, Available since v1.269.0) Extend the usage of the key, allowing one of the following:
any
serverAuth
clientAuth
codeSigning
emailProtection
timeStamping
OCSPSigning
Other Extended Key Usage OID
Enumerated values:
codeSigning: code signing
emailProtection: email protection
serverAuth: server authentication
timeStamping: Issuing timestamp
any: any
clientAuth: client authentication
OCSPSigning: OCSP Signature

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `locality` - (Required, ForceNew) Name of the city where the organization is located. Support the use of Chinese, English characters.
* `organization` - (Required, ForceNew) The name of the organization (corresponding to your enterprise or company) associated with the root CA certificate. Support the use of Chinese, English characters.
* `organization_unit` - (Required, ForceNew) The name of the department or branch under the organization. Support the use of Chinese, English characters.
* `parent_identifier` - (Optional, ForceNew, Available since v1.269.0) Parent node identifier.
* `path_len_constraint` - (Optional, Int, Available since v1.269.0) Child node level.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `resource_group_id` - (Optional, Computed, Available since v1.266.0) The ID of the resource group
* `state` - (Required, ForceNew)  The name of the province, municipality, or autonomous region in which the organization is located. Support the use of Chinese, English characters. 
 Name of the province or state where the organization is located. Support the use of Chinese, English characters. 
* `tags` - (Optional, Map, Available since v1.266.0) Tags
* `years` - (Required, ForceNew, Int) The validity period of the root CA certificate, in years.

-> **NOTE:**  It is recommended to set to 5~10 years.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `certificate_type` - The type of CA.
* `status` - The current CA status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Pca Certificate.
* `delete` - (Defaults to 5 mins) Used when delete the Pca Certificate.
* `update` - (Defaults to 5 mins) Used when update the Pca Certificate.

## Import

SSL Certificates Pca Certificate can be imported using the id, e.g.

```shell
$ terraform import alicloud_ssl_certificates_service_pca_certificate.example <identifier>
```