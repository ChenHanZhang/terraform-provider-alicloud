---
subcategory: "ESA"
layout: "alicloud"
page_title: "Alicloud: alicloud_esa_origin_pool"
description: |-
  Provides a Alicloud ESA Origin Pool resource.
---

# alicloud_esa_origin_pool

Provides a ESA Origin Pool resource.



For information about ESA Origin Pool and how to use it, see [What is Origin Pool](https://next.api.alibabacloud.com/document/ESA/2024-09-10/CreateOriginPool).

-> **NOTE:** Available since v1.244.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "random_integer" "default" {
  min = 10000
  max = 99999
}

resource "alicloud_esa_site" "default" {
  site_name   = "gositecdn-${random_integer.default.result}.cn"
  instance_id = data.alicloud_esa_sites.default.sites.0.instance_id
  coverage    = "overseas"
  access_type = "NS"
}


resource "alicloud_esa_origin_pool" "default" {
  origins {
    type    = "OSS"
    address = "example.oss-cn-beijing.aliyuncs.com"
    header  = "{\"Host\":[\"example.oss-cn-beijing.aliyuncs.com\"]}"
    enabled = "true"
    auth_conf {
      secret_key = "<SecretKeyId>"
      auth_type  = "private_cross_account"
      access_key = "<AccessKeyId>"
    }

    weight = "50"
    name   = "origin1"
  }
  origins {
    address = "example.s3.com"
    header  = "{\"Host\": [\"example1.com\"]}"
    enabled = "true"
    auth_conf {
      version    = "v2"
      region     = "us-east-1"
      auth_type  = "private"
      access_key = "<AccessKeyId>"
      secret_key = "<SecretKeyId>"
    }

    weight = "50"
    name   = "origin2"
    type   = "S3"
  }
  origins {
    type    = "S3"
    address = "example1111.s3.com"
    header  = "{\"Host\":[\"example1111.com\"]}"
    enabled = "true"
    auth_conf {
      secret_key = "<SecretKeyId>"
      version    = "v2"
      region     = "us-east-1"
      auth_type  = "private"
      access_key = "<AccessKeyId>"
    }

    weight = "30"
    name   = "origin3"
  }

  site_id          = alicloud_esa_site.default.id
  origin_pool_name = "exampleoriginpool"
  enabled          = "true"
}
```

## Argument Reference

The following arguments are supported:
* `enabled` - (Optional, Computed) Whether the source address pool is enabled:
  - `true`: Enabled;
  - `false`: Not enabled.
* `origin_pool_name` - (Required, ForceNew) The source address pool name.
* `origins` - (Optional, List) The Source station information added to the source address pool. Multiple Source stations use arrays to transfer values. See [`origins`](#origins) below.
* `site_id` - (Required, ForceNew) The site ID.

### `origins`

The origins supports the following:
* `address` - (Optional) Origin Address.
* `auth_conf` - (Optional, Set) The authentication information. When the source Station is an OSS or S3 and other source stations need to be authenticated, the authentication-related configuration information needs to be transmitted. See [`auth_conf`](#origins-auth_conf) below.
* `enabled` - (Optional) Whether the source station is enabled:
  - `true`: Enabled;
  - `false`: Not enabled.
* `header` - (Optional) The request header that is sent when returning to the source. Only Host is supported.
* `ip_version_policy` - (Optional, Computed, Available since v1.283.0) IP protocol version used for origin requests:
  - `round_robin`: Default policy; randomly alternates between IPv4 and IPv6 origin servers.
  - `ipv4_first`: Prefers IPv4 origin servers.
  - `ipv6_first`: Prefers IPv6 origin servers.
  - `follow`: Follows the IP version used by the client.
* `name` - (Optional) Origin Name.
* `type` - (Optional) Source station type:
ip_domain: ip or domain name type origin station;
  - `OSS`:OSS address source station;
  - `S3`:AWS S3 Source station.
* `weight` - (Optional, Int) Weight, 0-100.

### `origins-auth_conf`

The origins-auth_conf supports the following:
* `access_key` - (Optional) The AccessKey to be passed when AuthType is set to private_cross_account or private.
* `auth_type` - (Optional) Authentication type.
  - `public`: public read/write, which is used when the source station is OSS or S3 and is public read/write;
  - `private_same_account`: Used when the same account is private, the source station is OSS, and the authentication type is private authentication of the same account;
  - `private_cross_account`: private cross-account, used when the origin station is OSS and the authentication type is cross-account private authentication;
  - `private`: Used when the source station is S3 and the authentication type is private.
* `region` - (Optional) The Region of the source station to be transmitted when the source station is AWS S3.
* `secret_key` - (Optional) The SecretKey to be passed when AuthType is set to private_cross_account or private.
* `version` - (Optional) The signature version to be transmitted when the source station is AWS S3.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<site_id>:<origin_pool_id>`.
* `origin_pool_id` - OriginPool Id.
* `origins` - The Source station information added to the source address pool.
  * `origin_id` - Origin ID.
* `record_name` - The domain name assigned to the source address pool can be used as the source address recorded under the site.
* `reference_lb_count` - How many load balancers are referenced.
* `references` - The source address pool is referred to when the source address pool is configured by the load balancer or recorded as the source station.
  * `dns_records` - Use this source address pool for the seven-level record list of the source station.
    * `dns_record_id` - Record ID.
    * `name` - Record Name.
  * `ip_a_records` - Use this source address pool as a list of Layer 4 records for the source station.
    * `ipa_record_id` - Record ID.
    * `name` - Record Name.
  * `load_balancers` - A list of load balancers that use this source address pool.
    * `load_balancer_id` - Load Balancer ID.
    * `name` - Load Balancer Name.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Origin Pool.
* `delete` - (Defaults to 5 mins) Used when delete the Origin Pool.
* `update` - (Defaults to 5 mins) Used when update the Origin Pool.

## Import

ESA Origin Pool can be imported using the id, e.g.

```shell
$ terraform import alicloud_esa_origin_pool.example <site_id>:<origin_pool_id>
```