---
subcategory: "Certificate Management Service (Original SSL Certificate)"
layout: "alicloud"
page_title: "Alicloud: alicloud_ssl_certificates_service_instance"
description: |-
  Provides a Alicloud Certificate Management Service (Original SSL Certificate) Instance resource.
---

# alicloud_ssl_certificates_service_instance

Provides a Certificate Management Service (Original SSL Certificate) Instance resource.



For information about Certificate Management Service (Original SSL Certificate) Instance and how to use it, see [What is Instance](https://next.api.alibabacloud.com/document/BssOpenApi/2017-12-14/CreateInstance).

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


resource "alicloud_ssl_certificates_service_instance" "default" {
  contact_id_list = []
  certificate_id  = "25747924"
  product_type    = "cas_dv_public_cn"
  period          = "3"
  parameter {
    value = "ss.dv.t"
    code  = "fullSpec"
  }
  parameter {
    value = "1"
    code  = "fullDomainCount"
  }
  parameter {
    value = "exampleCert_product"
    code  = "product"
  }
}
```

## Argument Reference

The following arguments are supported:
* `auto_reissue` - (Optional) Specifies whether to enable managed renewal.
enable: Enabled
disable: Disabled

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `certificate_id` - (Optional, Int) The certificate ID. This field exists only after the certificate is issued.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `certificate_name` - (Optional) The name of the instance. When a certificate is issued, this value is used as the default name of the certificate.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `city` - (Optional) The city where the company or organization to which the certificate purchaser belongs is located. This field is required when generating a CSR for a DV certificate. Default value: Beijing.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `company_id` - (Optional) The company information ID. This parameter is required for OV and EV certificates. Otherwise, you cannot call the ApplyCertificate operation to apply for a certificate.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `contact_id_list` - (Optional, List) The list of contact IDs. At least one contact ID is required. Otherwise, you cannot call the ApplyCertificate operation to apply for a certificate.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `country_code` - (Optional) The code of the country or region where the certificate organization is located. For example, CN indicates China and US indicates the United States. This field is required when a Certificate Signing Request (CSR) is generated for a DV certificate. The default value is CN.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `csr` - (Optional) The CSR content. You can generate a CSR by using OpenSSL or Keytool. For more information, see [How to create a CSR file](https://help.aliyun.com/document_detail/42218.html).

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `domain` - (Optional) The domain names to be bound to the certificate. Specific requirements are as follows:
  - You can specify a single domain name or a wildcard domain name (such as `*.aliyundoc.com`).
  - You can specify multiple domain names. Separate multiple domain names with commas (,). Whether free domain names are provided is determined based on the first domain name.

-> **NOTE:** 

When you bind multiple domain names to a certificate, this parameter is required. This parameter and the `Csr` parameter cannot both be empty. If you specify both this parameter and the `Csr` parameter, the value of the `CN` field in the `Csr` parameter is used as the domain name bound to the certificate.

-> **NOTE:** 


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `generate_csr_method` - (Optional) The method used to generate the Certificate Signing Request (CSR). Default value: online.
online: The system generates the CSR. In this case, the Csr parameter is ignored.
upload: You upload the CSR. In this case, the Csr parameter is required.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `key_algorithm` - (Optional) The certificate algorithm. Default value: RSA_2048.
  - `RSA_2048`
  - `RSA_3072`
  - `RSA_4096`
  - `ECC_256`
  - `SM2`

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `parameter` - (Optional, List) The list of modules. See [`parameter`](#parameter) below.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `period` - (Optional, Int) The subscription period. Unit: months. For products billed on a yearly basis, enter an integer multiple of 12.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `pricing_cycle` - (Optional, Int) The unit of the subscription period.
  - PricingCycle=1: The unit of the subscription period is year.
  - PricingCycle=2: The unit of the subscription period is month.
  - PricingCycle=3: The unit of the subscription period is day.

Default value: PricingCycle=2

This parameter applies only to specific product types (ProductType is ddos_originpre_public_cn, ddosDip, ddoscoo, ddos_originpre_public_intl, ddosDip_intl, or ddoscoo_intl).

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `product_type` - (Optional) The product type.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `province` - (Optional) The province or region where the company is located. This field is required when generating a CSR for a DV certificate. Default value: Beijing.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `resource_group_id` - (Optional, Computed) The ID of the resource group.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `tags` - (Optional, Map) The list of tags.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `validation_method` - (Optional) The verification method for the certificate application.
DNS: DNS verification, using TXT or CNAME records.
HTTP: File verification.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.


### `parameter`

The parameter supports the following:
* `code` - (Required) The code configuration of the nth module property.
* `value` - (Required) The value configuration of the nth module property. Valid values of n: 1 to 100.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `brand` - The certificate brand.
* `certificate_status` - The status of the certificate.
* `certificate_type` - The type of the certificate.
* `instance_type` - The instance type.
* `status` - The instance status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Instance.
* `delete` - (Defaults to 5 mins) Used when delete the Instance.
* `update` - (Defaults to 5 mins) Used when update the Instance.

## Import

Certificate Management Service (Original SSL Certificate) Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_ssl_certificates_service_instance.example <instance_id>
```