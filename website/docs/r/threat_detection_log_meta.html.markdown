---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_log_meta"
description: |-
  Provides a Alicloud Threat Detection Log Meta resource.
---

# alicloud_threat_detection_log_meta

Provides a Threat Detection Log Meta resource. Log analysis shipping status.

For information about Threat Detection Log Meta and how to use it, see [What is Log Meta](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available since v1.211.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_threat_detection_log_meta" "default" {
  status      = "disabled"
  log_meta_id = "aegis-log-network"
}
```

### Deleting `alicloud_threat_detection_log_meta` or removing it from your configuration

Terraform cannot destroy resource `alicloud_threat_detection_log_meta`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `log_meta_id` - (Required, ForceNew) The name of the exclusive Logstore where logs are stored. Value:
  - aegis-log-crack: Brute Force log
  - aegis-log-dns-query:DNS request log
  - aegis-log-login: log in to the flow log
  - aegis-log-network: network connection log
  - aegis-log-process: process startup log
  - aegis-snapshot-host: account snapshot log
  - aegis-snapshot-port: port snapshot log
  - aegis-snapshot-process: process snapshot log
  - local-dns: local DNS log
  - sas-cspm-log: Cloud platform configuration check log
  - sas-hc-log: baseline log
  - sas-log-dns:DNS resolution log
  - sas-log-http:WEB access log
  - sas-log-session: Web session log
  - sas-security-log: alarm log
  - sas-vul-log: Vulnerability log.
* `status` - (Required) The status of the resource.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Log Meta.
* `update` - (Defaults to 5 mins) Used when update the Log Meta.

## Import

Threat Detection Log Meta can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_log_meta.example <id>
```