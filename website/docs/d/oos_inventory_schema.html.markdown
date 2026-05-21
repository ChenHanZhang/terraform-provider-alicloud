---
subcategory: "Operation Orchestration Service (OOS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_oos_inventory_schema"
sidebar_current: "docs-alicloud-datasource-oos-inventory-schema"
description: |-
  Provides OOS inventory schemas.
---

# alicloud_oos_inventory_schema

This data source provides OOS inventory schemas.

-> **NOTE:** Available since v1.279.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_oos_inventory_schema" "default" {
  type_name = "ACS:InstanceInformation"
}

output "first_attribute_name" {
  value = data.alicloud_oos_inventory_schema.default.schemas.0.attributes.0.name
}
```

Aggregator Schema Usage

```terraform
data "alicloud_oos_inventory_schema" "default" {
  aggregator = true
}
```

## Argument Reference

The following arguments are supported:

* `type_name` - (Optional) The inventory component name. Valid values: `ACS:InstanceInformation`, `ACS:Application`, `ACS:File`, `ACS:Network`, `ACS:WindowsRole`, `ACS:Service`, `ACS:WindowsRegistry`, `ACS:WindowsUpdate`.
* `aggregator` - (Optional) Whether to return only schemas that support aggregation.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:

* `ids` - A list of inventory component names.
* `schemas` - A list of OOS inventory schemas. Each element contains the following attributes:
    * `id` - The inventory component name.
    * `type_name` - The inventory component name.
    * `version` - The inventory component schema version.
    * `attributes` - The schema attributes. Each element contains the following attributes:
        * `name` - The attribute name.
        * `data_type` - The attribute data type.
    * `attributes_json` - A JSON string containing all raw attributes of this schema.
    * `schema_json` - A JSON string containing the raw schema.
* `schemas_json` - A JSON string containing all raw schemas.
