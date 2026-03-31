---
subcategory: "Cloud Enterprise Network (CEN)"
layout: "alicloud"
page_title: "Alicloud: alicloud_cen_transit_router_policy_table_association"
description: |-
  Provides a Alicloud Cloud Enterprise Network (CEN) Transit Router Policy Table Association resource.
---

# alicloud_cen_transit_router_policy_table_association

Provides a Cloud Enterprise Network (CEN) Transit Router Policy Table Association resource.

Transit Router Policy-Based Route Table Association.

For information about Cloud Enterprise Network (CEN) Transit Router Policy Table Association and how to use it, see [What is Transit Router Policy Table Association](https://next.api.alibabacloud.com/document/Cbn/2017-09-12/AssociateTransitRouterPolicyTable).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}


resource "alicloud_cen_transit_router_policy_table_association" "default" {
  policy_table_id = "ptbl-6ks0j63q2f5nrfees1"
  attachment_id   = "tr-attach-pljo1ollxcbin6mbgm"
}
```

## Argument Reference

The following arguments are supported:
* `attachment_id` - (Required, ForceNew) The ID of the transit router attachment instance.
* `dry_run` - (Optional) - `false` (default): Sends a normal request and performs the resource operation directly after passing the check.  
  - `true`: Sends a dry-run request to perform only validation without executing the resource operation. The validation includes checking whether required parameters are provided and whether the request format is correct. If the validation fails, an error is returned. If the validation passes, the error code `DryRunOperation` is returned.

-> **NOTE:** This parameter is only evaluated during resource creation and deletion. Modifying it in isolation will not trigger any action.

* `policy_table_id` - (Required, ForceNew) The ID of the policy-based route table instance.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `status` - The status of the policy-based route table association.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Transit Router Policy Table Association.
* `delete` - (Defaults to 5 mins) Used when delete the Transit Router Policy Table Association.

## Import

Cloud Enterprise Network (CEN) Transit Router Policy Table Association can be imported using the id, e.g.

```shell
$ terraform import alicloud_cen_transit_router_policy_table_association.example <association_id>
```