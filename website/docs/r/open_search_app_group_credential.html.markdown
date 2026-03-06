---
subcategory: "Open Search"
layout: "alicloud"
page_title: "Alicloud: alicloud_open_search_app_group_credential"
description: |-
  Provides a Alicloud Open Search App Group Credential resource.
---

# alicloud_open_search_app_group_credential

Provides a Open Search App Group Credential resource.

Application-level authentication credentials.

For information about Open Search App Group Credential and how to use it, see [What is App Group Credential](https://next.api.alibabacloud.com/document/OpenSearch/2017-12-25/CreateAppGroupCredentials).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-shanghai"
}

resource "alicloud_open_search_app_group" "defaultWQxaIV" {
  quota {
    spec             = "opensearch.share.common"
    doc_size         = "1"
    compute_resource = "20"
  }
  app_group_name = "credential_example"
  type           = "standard"
}


resource "alicloud_open_search_app_group_credential" "default" {
  enabled      = true
  type         = "api-token"
  app_group_id = alicloud_open_search_app_group.defaultWQxaIV.group_id
  dry_run      = false
}
```

## Argument Reference

The following arguments are supported:
* `app_group_id` - (Optional, ForceNew, Computed) Application id
* `dry_run` - (Optional) Whether running empty

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `enabled` - (Optional) Whether credentials are disabled.
* `type` - (Required, ForceNew) voucher type, valid values: `api-token`.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<app_group_id>:<app_group_credential_id>`.
* `app_group_credential_id` - The first ID of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the App Group Credential.
* `delete` - (Defaults to 5 mins) Used when delete the App Group Credential.
* `update` - (Defaults to 5 mins) Used when update the App Group Credential.

## Import

Open Search App Group Credential can be imported using the id, e.g.

```shell
$ terraform import alicloud_open_search_app_group_credential.example <app_group_id>:<app_group_credential_id>
```