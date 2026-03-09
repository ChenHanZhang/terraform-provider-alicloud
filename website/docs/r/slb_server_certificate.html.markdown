---
subcategory: "Classic Load Balancer (SLB)"
layout: "alicloud"
page_title: "Alicloud: alicloud_slb_server_certificate"
description: |-
  Provides a Alicloud SLB Server Certificate resource.
---

# alicloud_slb_server_certificate

Provides a SLB Server Certificate resource.

To configure HTTPS one-way authentication, you must provide a server certificate.

For information about SLB Server Certificate and how to use it, see [What is Server Certificate](https://www.alibabacloud.com/help/doc-detail/85968.htm).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

```terraform
# create a server certificate
resource "alicloud_slb_server_certificate" "foo" {
  name               = "slbservercertificate"
  server_certificate = "-----BEGIN CERTIFICATE-----\nMIICWDCCAcGgAwIBAgIJAP7vOtjPtQIjMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkNOMRMwEQYDVQQIDApjbi1iZWlqaW5nMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMjAxMDIwMDYxOTUxWhcNMjAxMTE5MDYxOTUxWjBF\nMQswCQYDVQQGEwJDTjETMBEGA1UECAwKY24tYmVpamluZzEhMB8GA1UECgwYSW50\nZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKB\ngQDEdoyaJ0kdtjtbLRx5X9qwI7FblhJPRcScvhQSE8P5y/b/T8J9BVuFIBoU8nrP\nY9ABz4JFklZ6SznxLbFBqtXoJTmzV6ixyjjH+AGEw6hCiA8Pqy2CNIzxr9DjCzN5\ntWruiHqO60O3Bve6cHipH0VyLAhrB85mflvOZSH4xGsJkwIDAQABo1AwTjAdBgNV\nHQ4EFgQUYDwuuqC2a2UPrfm1v31vE7+GRM4wHwYDVR0jBBgwFoAUYDwuuqC2a2UP\nrfm1v31vE7+GRM4wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOBgQAovSB0\n5JRKrg7lYR/KlTuKHmozfyL9UER0/dpTSoqsCyt8yc1BbtAKUJWh09BujBE1H22f\nlKvCAjhPmnNdfd/l9GrmAWNDWEDPLdUTkGSkKAScMpdS+mLmOBuYWgdnOtq3eQGf\nt07tlBL+dtzrrohHpfLeuNyYb40g8VQdp3RRRQ==\n-----END CERTIFICATE-----"
  private_key        = "-----BEGIN RSA PRIVATE KEY-----\nMIICXAIBAAKBgQDEdoyaJ0kdtjtbLRx5X9qwI7FblhJPRcScvhQSE8P5y/b/T8J9\nBVuFIBoU8nrPY9ABz4JFklZ6SznxLbFBqtXoJTmzV6ixyjjH+AGEw6hCiA8Pqy2C\nNIzxr9DjCzN5tWruiHqO60O3Bve6cHipH0VyLAhrB85mflvOZSH4xGsJkwIDAQAB\nAoGARe2oaCo5lTDK+c4Zx3392hoqQ94r0DmWHPBvNmwAooYd+YxLPrLMe5sMjY4t\ndmohnLNevCK1Uzw5eIX6BNSo5CORBcIDRmiAgwiYiS3WOv2+qi9g5uIdMiDr+EED\nK8wZJjB5E2WyfxL507vtW4T5L36yfr8SkmqH3GvzpI2jCqECQQDsy0AmBzyfK0tG\nNw1+iF9SReJWgb1f5iHvz+6Dt5ueVQngrl/5++Gp5bNoaQMkLEDsy0iHIj9j43ji\n0DON05uDAkEA1GXgGn8MXXKyuzYuoyYXCBH7aF579d7KEGET/jjnXx9DHcfRJZBY\nB9ghMnnonSOGboF04Zsdd3xwYF/3OHYssQJAekd/SeQEzyE5TvoQ8t2Tc9X4yrlW\nxNX/gmp6/fPr3biGUEtb7qi+4NBodCt+XsingmB7hKUP3RJTk7T2WnAC5wJAMqHi\njY5x3SkFkHl3Hq9q2CKpQxUbCd7FXqg1wum/xj5GmqfSpNjHE3+jUkwbdrJMTrWP\nrmRy3tQMWf0mixAo0QJBAN4IcZChanq8cZyNqqoNbxGm4hkxUmE0W4hxHmLC2CYZ\nV4JpNm8dpi4CiMWLasF6TYlVMgX+aPxYRUWc/qqf1/Q=\n-----END RSA PRIVATE KEY-----"
}
```

## Argument Reference

The following arguments are supported:
* `ali_cloud_certificate_id` - (Optional, ForceNew) The ID of the server certificate from Alibaba Cloud SSL Certificate Service.
* `ali_cloud_certificate_name` - (Optional, ForceNew) The name of the server certificate from Alibaba Cloud SSL Certificate Service.
* `private_key` - (Optional) The private key to be uploaded.

-> **NOTE:**  If you do not use the Alibaba Cloud certificate, this parameter is required.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `resource_group_id` - (Optional, Computed) The ID of the enterprise resource group.
* `server_certificate` - (Optional) The public key certificate to be uploaded.

-> **NOTE:**  If you do not use the Alibaba Cloud certificate, this parameter is required.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `server_certificate_name` - (Optional) The name of the server certificate.
* `tags` - (Optional, Map) The tag of the resource

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.
* `region_id` - The ID of the region to which the server certificate belongs.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Server Certificate.
* `delete` - (Defaults to 5 mins) Used when delete the Server Certificate.
* `update` - (Defaults to 5 mins) Used when update the Server Certificate.

## Import

SLB Server Certificate can be imported using the id, e.g.

```shell
$ terraform import alicloud_slb_server_certificate.example <server_certificate_id>
```