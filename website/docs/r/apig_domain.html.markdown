---
subcategory: "APIG"
layout: "alicloud"
page_title: "Alicloud: alicloud_apig_domain"
description: |-
  Provides a Alicloud APIG Domain resource.
---

# alicloud_apig_domain

Provides a APIG Domain resource.



For information about APIG Domain and how to use it, see [What is Domain](https://next.api.alibabacloud.com/document/APIG/2024-03-27/CreateDomain).

-> **NOTE:** Available since v1.275.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_ssl_certificates_service_certificate" "defaultwxeuip" {
  cert             = <<EOF
-----BEGIN CERTIFICATE----- MIIGBTCCBO2gAwIBAgIQBfXjfSqmbbYshNgm5P/P/jANBgkqhkiG9w0BAQsFADBu MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3 d3cuZGlnaWNlcnQuY29tMS0wKwYDVQQDEyRFbmNyeXB0aW9uIEV2ZXJ5d2hlcmUg RFYgVExTIENBIC0gRzIwHhcNMjUwMTEzMDAwMDAwWhcNMjUwNDEzMjM1OTU5WjAi MSAwHgYDVQQDExdkZW1vLnB0cy5jbGRuYXRpdmUudGVjaDCCASIwDQYJKoZIhvcN AQEBBQADggEPADCCAQoCggEBAIzzlFX1WK2iTAYr+ABki3zICfGxMZY9EyzhjgJb MYIvsxuavmMSPRCgBiP4N8qikMPMwzAU1CIrvWDQ0zp5c1vRReDmVCbTXMWs29Fu ljhdwOakNDe7eKwe7rOhLVl/fU579cPPOxAriK7kmynbLEzDQJVF0umt+/k1kFmR +tAfSBiUB8NjAAm34vs5jPNQtlff632NDohkDakl6o6g9//PT/gzuXKyb+L20o5G AfMwlv3Y8FijMw5vXtAu37xWw2+KHj81XoVjZaHZ9ikukF+OTshupDmdMxjVjs7L EPGsgGHQIYZIl5Cerp04BdKYrephvJlFuNFRy/rEwcRzIXMCAwEAAaOCAukwggLl MB8GA1UdIwQYMBaAFHjfkZBf7t6s9sV169VMVVPvJEq2MB0GA1UdDgQWBBRS+k7E 4r0iBMN7XsfqlpsdMldOTjAiBgNVHREEGzAZghdkZW1vLnB0cy5jbGRuYXRpdmUu dGVjaDA+BgNVHSAENzA1MDMGBmeBDAECATApMCcGCCsGAQUFBwIBFhtodHRwOi8v d3d3LmRpZ2ljZXJ0LmNvbS9DUFMwDgYDVR0PAQH/BAQDAgWgMB0GA1UdJQQWMBQG CCsGAQUFBwMBBggrBgEFBQcDAjCBgAYIKwYBBQUHAQEEdDByMCQGCCsGAQUFBzAB hhhodHRwOi8vb2NzcC5kaWdpY2VydC5jb20wSgYIKwYBBQUHMAKGPmh0dHA6Ly9j YWNlcnRzLmRpZ2ljZXJ0LmNvbS9FbmNyeXB0aW9uRXZlcnl3aGVyZURWVExTQ0Et RzIuY3J0MAwGA1UdEwEB/wQCMAAwggF9BgorBgEEAdZ5AgQCBIIBbQSCAWkBZwB1 AE51oydcmhDDOFts1N8/Uusd8OCOG41pwLH6ZLFimjnfAAABlF2cgF0AAAQDAEYw RAIgf01DcLaZanOBlyTFvhvZodKtkbcpJpWhk2Sl2p8NCPcCIFWHLUwJJyZOLAKI lexLFRKt2iIUzkNTk7ABwKfwaUCWAHcAcyAiDwgWivnzxKaLCrJqmkoA7vV3hYoI TQUA1KVCRFkAAAGUXZyAfwAABAMASDBGAiEA226s92d46MrlIK+J1hvMhT5pwzwK knUUSQXfEJXjTecCIQDFwrtn6CypuZ4tce6hEhjik9dx/3jGQrBqhIP21seMiAB1 AObSMWNAd4zBEEEG13G5zsHSQPaWhIb7uocyHf0eN45QAAABlF2cgJAAAAQDAEYw RAIgIpgYibS0+f4TeYINGZm5JxHy4jbL1cbKc37GaM5BZqICIHoAHyd04+2Atgd6 3I4Oc+jgSwY6OnGVAEP6crsySjEEMA0GCSqGSIb3DQEBCwUAA4IBAQAERKUT5CEb XzptbEAHesGAgVpeCFKw54+rmRxrlbmBj9S87T0AONsoefWdvIWvnhElhgeMEVVj UjZ55tnDZt8d8ah1iJN8/YUbPleshwHD8kpQiD/AgxsV4uJU0rVjQo4oG6n/zu7r O2PjOffNjD8V92LR/2Z86/71gK2ZBjUw1ohvkRj3hSjZysj/kw8dZTBYOd+P84e9 gKpnUKUhgmL9/25AzgXQQEFouwxxF76B/bmP63u24WN1ZjOjBwq4uwK+pXZeFuUi MmmOojiwzXTd9rZwyO07shbF0fSXWBXe0AZbieeFT+qioaHCgNiye6X689U8xmCp SZ8TnS2v0MjH -----END CERTIFICATE----- -----BEGIN CERTIFICATE----- MIIEqjCCA5KgAwIBAgIQDeD/te5iy2EQn2CMnO1e0zANBgkqhkiG9w0BAQsFADBh MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3 d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBH MjAeFw0xNzExMjcxMjQ2NDBaFw0yNzExMjcxMjQ2NDBaMG4xCzAJBgNVBAYTAlVT MRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j b20xLTArBgNVBAMTJEVuY3J5cHRpb24gRXZlcnl3aGVyZSBEViBUTFMgQ0EgLSBH MjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAO8Uf46i/nr7pkgTDqnE eSIfCFqvPnUq3aF1tMJ5hh9MnO6Lmt5UdHfBGwC9Si+XjK12cjZgxObsL6Rg1njv NhAMJ4JunN0JGGRJGSevbJsA3sc68nbPQzuKp5Jc8vpryp2mts38pSCXorPR+sch QisKA7OSQ1MjcFN0d7tbrceWFNbzgL2csJVQeogOBGSe/KZEIZw6gXLKeFe7mupn NYJROi2iC11+HuF79iAttMc32Cv6UOxixY/3ZV+LzpLnklFq98XORgwkIJL1HuvP ha8yvb+W6JislZJL+HLFtidoxmI7Qm3ZyIV66W533DsGFimFJkz3y0GeHWuSVMbI lfsCAwEAAaOCAU8wggFLMB0GA1UdDgQWBBR435GQX+7erPbFdevVTFVT7yRKtjAf BgNVHSMEGDAWgBROIlQgGJXm427mD/r6uRLtBhePOTAOBgNVHQ8BAf8EBAMCAYYw HQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMBIGA1UdEwEB/wQIMAYBAf8C AQAwNAYIKwYBBQUHAQEEKDAmMCQGCCsGAQUFBzABhhhodHRwOi8vb2NzcC5kaWdp Y2VydC5jb20wQgYDVR0fBDswOTA3oDWgM4YxaHR0cDovL2NybDMuZGlnaWNlcnQu Y29tL0RpZ2lDZXJ0R2xvYmFsUm9vdEcyLmNybDBMBgNVHSAERTBDMDcGCWCGSAGG /WwBAjAqMCgGCCsGAQUFBwIBFhxodHRwczovL3d3dy5kaWdpY2VydC5jb20vQ1BT MAgGBmeBDAECATANBgkqhkiG9w0BAQsFAAOCAQEAoBs1eCLKakLtVRPFRjBIJ9LJ L0s8ZWum8U8/1TMVkQMBn+CPb5xnCD0GSA6L/V0ZFrMNqBirrr5B241OesECvxIi 98bZ90h9+q/X5eMyOD35f8YTaEMpdnQCnawIwiHx06/0BfiTj+b/XQih+mqt3ZXe xNCJqKexdiB2IWGSKcgahPacWkk/BAQFisKIFYEqHzV974S3FAz/8LIfD58xnsEN GfzyIDkH3JrwYZ8caPTf6ZX9M1GrISN8HnWTtdNCH2xEajRa/h9ZBXjUyFKQrGk2 n2hcLrfZSbynEC/pSw/ET7H5nWwckjmAJ1l9fcnbqkU/pf6uMQmnfl0JQjJNSg== -----END CERTIFICATE-----
EOF
  certificate_name = "zhenyuanexample-cert"
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY----- MIIEogIBAAKCAQEAjPOUVfVYraJMBiv4AGSLfMgJ8bExlj0TLOGOAlsxgi+zG5q+ YxI9EKAGI/g3yqKQw8zDMBTUIiu9YNDTOnlzW9FF4OZUJtNcxazb0W6WOF3A5qQ0 N7t4rB7us6EtWX99Tnv1w887ECuIruSbKdssTMNAlUXS6a37+TWQWZH60B9IGJQH w2MACbfi+zmM81C2V9/rfY0OiGQNqSXqjqD3/89P+DO5crJv4vbSjkYB8zCW/djw WKMzDm9e0C7fvFbDb4oePzVehWNlodn2KS6QX45OyG6kOZ0zGNWOzssQ8ayAYdAh hkiXkJ6unTgF0pit6mG8mUW40VHL+sTBxHMhcwIDAQABAoIBAB2RJO8pIbp+nor2 4zWV4crBnZBlFp1I34JdKDBnV8nS3rPuMHzQwR2BK/iUo4NuuYZRKCqoW911Jmsu MvkDrAa5Lfrg5gsk8EOJIK1ckMdvJz1aeZetTEFy1ai5qS0qeKsjhqjWCjI4p5nk W4Yasz7BAIWvmXAqbn5W1ZpsbFwGvCNIZhqNPG57o42jpRnCIK8HwRYadDMxgg23 /kBLm3ivkAg83R1qiC2c54WurWxpNxYAyVDMpFioOmjuk8LznU6IDozwGlKQqSIp 76Kl08yERTSnBUd6izjfskZwxlq9mV3rhAEjgMwPWcvQ7SvdCdM8r1cegmWwZd2i 0hv3pskCgYEAvwWEjcp1DoZfpEuK0bEQ50ZA+sYU0j+Ve9R2e00Rh+OcrCvjjHBL l1hXeU1N4l6INfetyUKApqrKEGGSRBxk0nlZSCQuVshGFIyegIJlrJ2li+UVKca/ 2knIL3gjFV3bcUurkGJrNsPelHBmHisegpNnNSwoHZEYK2u4nUsEIGsCgYEAvOXf 4cBVCsKEXjUX0795nH1dzMU4kAz/99OC4g/fDKpfctYej7eT2jlbPN0UVUPYy43U PWmBJw3t5FnxKjNFnMobU7CEv7d49YYiLz3mQcMUwokkdse+urQzkmCu1tr61ZxV AMhnoDbhmumEd6PcWRZ7jnDe5eRI8UE0r7Y1pRkCgYB0nt7FXiRHANylRoVy7eI/ AFIth/8wOSd/GUBYmL2qM9jz+DuNAwXzWTjWUs7I9CO+pv7Pj61Rk5WXmKoesSMQ qItMkuJDkzPN2efu1y1RzFFEblCUV8XLdB4mdPceVVXi/mq41I1WOxDJmTGPG44d 8/YfrVy0aF/UKojkZucXDwKBgEm9s1h9rLU2mlXshBC1ftQIXMXDeAFEEei9OSrm 5zwahohd7jBMift2yJdJ9tYSVl+gTmpq9XA5j9TFB9Bhk5tRirlw/2GYTjAK3O10 TJ7/eOs4fnOgJmTcVAWgmKBne+RH0ycrFMkGo6vF3WCXJz3f+PvyNBGqvI7x1Y+C og6BAoGAEoeQFTdZ7wo1Bp2pqOe3POxjMIpYz2iZu71C72OM/XmjLYWRBWVyHpW+ D+QzZd0bhyiokzcDa5cMeDKUjGDJqD0R+xEATvDVx22bV/NbFCqljHy5GrwWOwJ5 Hp+yD0gABPP+UPQCeabvORMgk/AO70UfluxqJgU7C3fUhf9pWyU= -----END RSA PRIVATE KEY-----
EOF
}


resource "alicloud_apig_domain" "default" {
  domain_name = "www.baidu0.com"
  protocol    = "HTTP"
}
```

## Argument Reference

The following arguments are supported:
* `cert_identifier` - (Optional) tls cert identifier
* `domain_name` - (Required, ForceNew) domain name
* `force_https` - (Optional) Set the HTTPS protocol type and whether to enable forced HTTPS redirection.
* `http2_option` - (Optional) Whether to enable http2 settings.
* `protocol` - (Required) Protocol, HTTP/HTTPS
* `resource_group_id` - (Optional, Computed) The ID of the resource group
* `tls_cipher_suites_config` - (Optional, Set) TlsCipherSuitesConfig See [`tls_cipher_suites_config`](#tls_cipher_suites_config) below.
* `tls_max` - (Optional) The maximum version of the TLS protocol. The maximum version of the TLS protocol is 1.3.
* `tls_min` - (Optional) The minimum version of the TLS protocol. The minimum version of the TLS protocol is 1.0.

### `tls_cipher_suites_config`

The tls_cipher_suites_config supports the following:
* `config_type` - (Optional) config type, Default or Custom
* `tls_cipher_suite` - (Optional, List) tls Cipher Suite See [`tls_cipher_suite`](#tls_cipher_suites_config-tls_cipher_suite) below.

### `tls_cipher_suites_config-tls_cipher_suite`

The tls_cipher_suites_config-tls_cipher_suite supports the following:
* `name` - (Optional) cipher suite name

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