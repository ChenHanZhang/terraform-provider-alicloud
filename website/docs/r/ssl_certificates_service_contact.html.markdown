---
subcategory: "Certificate Management Service (Original SSL Certificate)"
layout: "alicloud"
page_title: "Alicloud: alicloud_ssl_certificates_service_contact"
description: |-
  Provides a Alicloud Ssl Certificates Service Contact resource.
---

# alicloud_ssl_certificates_service_contact

Provides a Ssl Certificates Service Contact resource.

Certificate Contact Person.

For information about Ssl Certificates Service Contact and how to use it, see [What is Contact](https://next.api.alibabacloud.com/document/cas/2020-04-07/CreateContact).

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


resource "alicloud_ssl_certificates_service_contact" "default" {
  email  = "examplecontact@example.com"
  idcard = "110101199001011234"
  mobile = "13800138001"
  name   = "TestContact-SDK"
}
```

## Argument Reference

The following arguments are supported:
* `email` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `idcard` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `mobile` - (Required) This property does not have a description in the spec, please add it before generating code.
* `name` - (Required) The name of the resource
* `webhooks` - (Optional) This property does not have a description in the spec, please add it before generating code.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Contact.
* `delete` - (Defaults to 5 mins) Used when delete the Contact.
* `update` - (Defaults to 5 mins) Used when update the Contact.

## Import

Ssl Certificates Service Contact can be imported using the id, e.g.

```shell
$ terraform import alicloud_ssl_certificates_service_contact.example <contact_id>
```