---
subcategory: "SSL Certificates"
layout: "alicloud"
page_title: "Alicloud: alicloud_ssl_certificates_service_pca_cert"
description: |-
  Provides a Alicloud SSL Certificates Pca Cert resource.
---

# alicloud_ssl_certificates_service_pca_cert

Provides a SSL Certificates Pca Cert resource.



For information about SSL Certificates Pca Cert and how to use it, see [What is Pca Cert](https://next.api.alibabacloud.com/document/cas/2020-06-30/CreateClientCertificate).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_ssl_certificates_service_pca_certificate" "创建根CA" {
  organization      = "a"
  years             = "1"
  locality          = "a"
  organization_unit = "a"
  state             = "a"
  country_code      = "cn"
  common_name       = "cbc.certqa.cn"
  algorithm         = "RSA_2048"
  certificate_type  = "ROOT"
}

resource "alicloud_ssl_certificates_service_pca_certificate" "创建子CA" {
  organization        = "a"
  years               = "1"
  locality            = "a"
  organization_unit   = "a"
  state               = "a"
  country_code        = "cn"
  common_name         = "cbc.certqa.cn"
  algorithm           = "RSA_2048"
  certificate_type    = "SUB_ROOT"
  enable_crl          = true
  crl_day             = "1"
  path_len_constraint = "0"
  parent_identifier   = alicloud_ssl_certificates_service_pca_certificate.创建根CA.id
}


resource "alicloud_ssl_certificates_service_pca_cert" "default" {
  immediately       = "0"
  organization      = "terraform"
  years             = "1"
  upload_flag       = "0"
  locality          = "terraform"
  months            = "1"
  custom_identifier = "498"
  algorithm         = "RSA_2048"
  parent_identifier = alicloud_ssl_certificates_service_pca_certificate.创建子CA.id
  san_value         = "somebody@example.com"
  enable_crl        = "1"
  organization_unit = "aliyun"
  state             = "Beijing"
  before_time       = "1767948807"
  days              = "1"
  san_type          = "1"
  after_time        = "1768035207"
  country_code      = "cn"
  common_name       = "exampleTerraform"
  alias_name        = "AliasName"
  status            = "ISSUE"
}
```

## Argument Reference

The following arguments are supported:
* `after_time` - (Optional, Int) The expiration time of the certificate is represented in the format of a timestamp. Unit: seconds.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `algorithm` - (Optional, ForceNew, Computed) The key algorithm type of the root CA certificate. The key algorithm is expressed using the '_< key length>' format. Value:
  - `RSA_1024`: The corresponding signature algorithm is Sha256WithRSA.
  - `RSA_2048`: The corresponding signature algorithm is Sha256WithRSA.
  - `RSA_4096`: The corresponding signature algorithm is Sha256WithRSA.
  - `ECC_256`: The signature algorithm is Sha256WithECDSA.
  - `ECC_384`: The corresponding signature algorithm is Sha256WithECDSA.
  - `ECC_512`: The signature algorithm is Sha256WithECDSA.
  - `SM2_256`: The corresponding signature algorithm is SM3WithSM2.

The encryption algorithm of the root CA certificate must be the same as the **certificate algorithm** of the private Root CA you purchased. Example: If the **certificate algorithm** selected when you purchase a private Root CA is `RSA`, the key algorithm of the root CA certificate must be **RSA\_1024**, **RSA\_2048**, or **RSA\_4096**.
* `alias_name` - (Optional) Alias Name For cert.
* `before_time` - (Optional, Int) The issuance time of the certificate is in timestamp format, and defaults to the time when you called this interface. Unit: seconds.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `common_name` - (Optional, ForceNew) The common name or abbreviation of the organization. Support the use of Chinese, English characters.
* `country_code` - (Optional, ForceNew) The code of the country or region in which the organization is located, using a two-digit capital abbreviation. For example, `CN` represents China and `US` represents the United States.
For more information about the codes of different countries, see **International Code** in [Management Company Information](~~ 198289 ~~).
* `custom_identifier` - (Optional, ForceNew) 用户创建私有证书时可以指定自定义标识，用来唯一标识所创建的私有证书。
* `days` - (Optional, ForceNew, Computed, Int) The validity period of the client certificate. Unit: days. The Days, BeforeTime, and AfterTime parameters cannot be empty simultaneously, and the BeforeTime and AfterTime parameters must either both be empty or both be set. The specific settings for this parameter are as follows:
When setting the Days parameter, you have the option to set both the BeforeTime and AfterTime parameters simultaneously, or choose not to set the BeforeTime and AfterTime parameters.
When the Days parameter is not set, you must set the BeforeTime and AfterTime parameters.
When you set the Days, BeforeTime, and AfterTime parameters simultaneously, the validity period of the client certificate is determined by the value of the Days parameter.
The validity period of the client certificate cannot exceed that of the subordinate CA certificate. You can call DescribeCACertificate to view the validity period of the subordinate CA certificate.
* `enable_crl` - (Optional, Int) Whether to enable CRL.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `immediately` - (Optional, Int) Return the digital certificate immediately.
0, no return. Default value.
1. Return the certificate.
2. Return the certificate and its certificate chain.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `locality` - (Optional, ForceNew) Name of the city where the organization is located. Support the use of Chinese, English characters.
* `months` - (Optional, Int) Duration of certificate purchase. Unit: month.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `organization` - (Optional, ForceNew) The name of the organization (corresponding to your enterprise or company) associated with the root CA certificate. Support the use of Chinese, English characters.
* `organization_unit` - (Optional, ForceNew) The name of the department or branch under the organization. Support the use of Chinese, English characters.
* `parent_identifier` - (Required, ForceNew) Parent node identifier.
* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `san_type` - (Optional) The extension information supported by the client certificate is of type SAN (Subject Alternative Name). Values:
1: Denotes an email address.
6: Denotes the Uniform Resource Identifier (URI).

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `san_value` - (Optional) Specific extension information of the client certificate. Multiple extension information is supported. If you need to enter multiple extension information, please separate them with a half-width comma (,).

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `state` - (Optional, ForceNew)  The name of the province, municipality, or autonomous region in which the organization is located. Support the use of Chinese, English characters. 
 Name of the province or state where the organization is located. Support the use of Chinese, English characters. 
* `status` - (Optional, Computed) The current CA status. Value:
  - issue: Enabled.
  - forbidden: Disabled.
  - revoke: Revoked.
* `tags` - (Optional, Map) Tags
* `upload_flag` - (Optional, Int) Indicate whether this certificate has been uploaded to the SSL certificate management platform.
* `years` - (Optional, Int) The validity period of the root CA certificate, in years.

-> **NOTE:**  It is recommended to set to 5~10 years.


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Pca Cert.
* `delete` - (Defaults to 5 mins) Used when delete the Pca Cert.
* `update` - (Defaults to 5 mins) Used when update the Pca Cert.

## Import

SSL Certificates Pca Cert can be imported using the id, e.g.

```shell
$ terraform import alicloud_ssl_certificates_service_pca_cert.example <identifier>
```