---
subcategory: "Live"
layout: "alicloud"
page_title: "Alicloud: alicloud_live_domain"
description: |-
  Provides a Alicloud Live Domain resource.
---

# alicloud_live_domain

Provides a Live Domain resource.

Live domain name.

For information about Live Domain and how to use it, see [What is Domain](https://next.api.alibabacloud.com/document/live/2016-11-01/AddLiveDomain).

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

variable "domain_name" {
  default = "antang-terraform-domain20260210020202025151"
}

data "alicloud_resource_manager_resource_groups" "default" {}


resource "alicloud_live_domain" "default" {
        domain_type = "liveVideo"
          scope = "overseas"
          domain_name = "${var.domain_name}"
          region = "ap-southeast-1"
          check_url = "http://${{ref(variable, domainName)}}/example.html"
          top_level_domain = " "
          status = "online"
  }
```

## Argument Reference

The following arguments are supported:
* `check_url` - (Optional) The health check URL.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `domain_name` - (Required, ForceNew) Broadcast the basin name or push the basin name.
* `domain_type` - (Required, ForceNew) Domain name business type. Value:
  - `liveVideo`: the name of the broadcast basin.
  - `liveEdge`: the name of the edge push basin.
* `owner_id` - (Optional, Int) The Id of the resource owner.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `region` - (Required, ForceNew) The region to which the domain name belongs.
* `resource_group_id` - (Optional, Computed) The Id of the resource group to which the domain name belongs.
* `scope` - (Optional) Acceleration zone. International users, China station L3 and above user settings are valid. Value:
  - `domestic` (default): Mainland China.
  - `overseas`: overseas and Hong Kong, Macao and Taiwan accelerated.
  - `global`: global acceleration.
* `status` - (Optional, Computed) The status of the resource
* `tags` - (Optional, Map) Resource Tags
* `top_level_domain` - (Optional) Top-level access domain.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - When the live domain name was created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 14 mins) Used when create the Domain.
* `delete` - (Defaults to 5 mins) Used when delete the Domain.
* `update` - (Defaults to 5 mins) Used when update the Domain.

## Import

Live Domain can be imported using the id, e.g.

```shell
$ terraform import alicloud_live_domain.example <domain_name>
```