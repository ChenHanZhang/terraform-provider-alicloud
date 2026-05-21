---
subcategory: "Operation Orchestration Service (OOS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_oos_inventory_search"
sidebar_current: "docs-alicloud-datasource-oos-inventory-search"
description: |-
  Provides OOS inventory search results.
---

# alicloud_oos_inventory_search

This data source provides OOS inventory search results, including detailed inventory entities or aggregation results.

-> **NOTE:** Available since v1.279.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_oos_inventory_search" "default" {
  filter {
    name     = "ACS:InstanceInformation.InstanceId"
    operator = "Equal"
    values   = ["i-bp1cpoxxxwxxxxxxxxxx"]
  }
}

output "first_entity" {
  value = jsondecode(data.alicloud_oos_inventory_search.default.entities[0])
}
```

Aggregator Usage

```terraform
data "alicloud_oos_inventory_search" "default" {
  aggregators = ["ACS:Application.Name"]
}
```

## Argument Reference

The following arguments are supported:

* `filter` - (Optional) The inventory filters. You can specify up to 5 filters. See [`filter`](#filter) below.
* `aggregators` - (Optional) A list of aggregators used to query inventory aggregation information. Valid values include `ACS:Application.Name` and `ACS:Application.Version`.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).

### `filter`

The `filter` block supports the following:

* `name` - (Required) The component attribute name. The value should include the component prefix, such as `ACS:InstanceInformation.InstanceId`.
* `operator` - (Optional) The comparison operator. Valid values: `Equal`, `NotEqual`, `BeginWith`, `LessThan`, `GreaterThan`. Default to `Equal`.
* `values` - (Required) The attribute values to match. You can specify up to 20 values.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:

* `ids` - A list of entity IDs. If an entity does not include an `Id` field, its zero-based result index is used.
* `entities` - A list of inventory entity JSON strings. Use `jsondecode` to access nested attributes.
* `entities_json` - A JSON string containing all inventory entities.
