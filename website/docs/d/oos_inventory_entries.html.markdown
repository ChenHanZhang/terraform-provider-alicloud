---
subcategory: "Operation Orchestration Service (OOS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_oos_inventory_entries"
sidebar_current: "docs-alicloud-datasource-oos-inventory-entries"
description: |-
  Provides OOS inventory entries of a specified instance and component.
---

# alicloud_oos_inventory_entries

This data source provides OOS inventory entries of a specified ECS instance and inventory component.

-> **NOTE:** Available since v1.279.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_oos_inventory_entries" "default" {
  instance_id = "i-bp1cpoxxxwxxxxxxxxxx"
  type_name   = "ACS:InstanceInformation"
}

output "first_entry" {
  value = jsondecode(data.alicloud_oos_inventory_entries.default.entries[0])
}
```

Filter Usage

```terraform
data "alicloud_oos_inventory_entries" "default" {
  instance_id = "i-bp1cpoxxxwxxxxxxxxxx"
  type_name   = "ACS:InstanceInformation"

  filter {
    name     = "PlatformName"
    operator = "Equal"
    values   = ["ubuntu"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required) The ID of the ECS instance.
* `type_name` - (Required) The inventory component name. Valid values: `ACS:InstanceInformation`, `ACS:Application`, `ACS:File`, `ACS:Network`, `ACS:WindowsRole`, `ACS:Service`, `ACS:WindowsRegistry`, `ACS:WindowsUpdate`.
* `filter` - (Optional) The inventory component filters. You can specify up to 5 filters. See [`filter`](#filter) below.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).

### `filter`

The `filter` block supports the following:

* `name` - (Required) The attribute name in the component result, such as `PlatformName`.
* `operator` - (Optional) The comparison operator. Valid values: `Equal`, `NotEqual`, `BeginWith`, `LessThan`, `GreaterThan`. Default to `Equal`.
* `values` - (Required) The attribute values to match. You can specify up to 20 values.

## Attributes Reference

The following attributes are exported in addition to the arguments listed above:

* `schema_version` - The inventory component schema version.
* `capture_time` - The time when the inventory information was synchronized.
* `entries` - A list of inventory entry JSON strings. Use `jsondecode` to access nested attributes.
* `entries_json` - A JSON string containing all inventory entries.
