---
subcategory: "PAI Workspace"
layout: "alicloud"
page_title: "Alicloud: alicloud_pai_workspace_connection"
description: |-
  Provides a Alicloud PAI Workspace Connection resource.
---

# alicloud_pai_workspace_connection

Provides a PAI Workspace Connection resource.

Connection configuration for model services and databases.

For information about PAI Workspace Connection and how to use it, see [What is Connection](https://next.api.alibabacloud.com/document/AIWorkSpace/2021-02-04/CreateConnection).

-> **NOTE:** Available since v1.269.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_pai_workspace_workspace" "defaultMr2RQE" {
  description    = "example_connection_1769080362"
  display_name   = "连接管理镇元资源example用例"
  workspace_name = "connection_1769080362"
  env_types      = ["prod"]
}


resource "alicloud_pai_workspace_connection" "default" {
  connection_name = "example_connection_1769080363"
  description     = "镇元资源example用例"
  accessibility   = "PRIVATE"
  connection_type = "MilvusConnection"
  secrets {
  }
  configs {
  }
  workspace_id = alicloud_pai_workspace_workspace.defaultMr2RQE.id
  resource_meta {
    instance_name = "example_inst"
    instance_id   = "c-b1c5222fba7cxxxx"
  }
  models {
    model        = "qwen-vl-max"
    display_name = "通义千问VL-Max"
    model_type   = "LLM"
  }
  models {
    model        = "qwen-qwq"
    display_name = "通义千问qwq"
    model_type   = "LLM"
  }
  models {
    model        = "qwen-plus"
    display_name = "通义千问plus"
    model_type   = "LLM"
    tool_call    = true
  }
}
```

## Argument Reference

The following arguments are supported:
* `accessibility` - (Optional, ForceNew) Resource visibility. Valid values:  
  - PUBLIC: Accessible to all members of the current workspace.  
  - PRIVATE: Accessible only to the creator.  
* `configs` - (Required, Map) Configuration information for the connection. Configurations are specified as key-value pairs. The keys vary depending on the connection type. For details, see the supplementary description of request parameters in the CreateConnection API.  
* `connection_id` - (Optional, ForceNew, Computed) The connection ID. For information about how to obtain a connection ID, see [ListConnections](url).  
* `connection_name` - (Required, ForceNew) The connection name.  
* `connection_type` - (Required, ForceNew) Type of the connection. Valid values:  
  - DashScopeConnection: DashScope service connection.  
  - OpenLLMConnection: Open-source LLM connection.  
  - MilvusConnection: Milvus connection.  
  - OpenSearchConnection: OpenSearch connection.  
  - LindormConnection: Lindorm connection.  
  - ElasticsearchConnection: Elasticsearch connection.  
  - HologresConnection: Hologres connection.  
  - RDSConnection: ApsaraDB RDS connection.  
  - CustomConnection: Custom connection.  
* `description` - (Optional) Description of the connection.  
* `models` - (Optional, List) List of models. Applicable to connections of the model service type.   See [`models`](#models) below.
* `resource_meta` - (Optional, ForceNew, Set) Resource metadata. See [`resource_meta`](#resource_meta) below.
* `secrets` - (Optional, Map) Key-value pairs that require encryption, such as database login passwords or model access keys.  
* `validate_type` - (Optional) The validation type used when verifying the connection. Valid values:  
  - Connectivity: connectivity test.  

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `workspace_id` - (Required, ForceNew) Workspace ID. For information about how to obtain a workspace ID, see [ListWorkspaces](~~449124~~).  

### `models`

The models supports the following:
* `display_name` - (Optional) The display name of the model.  
* `model` - (Optional) The model identifier.  
* `model_type` - (Optional) The model type.  
* `tool_call` - (Optional) Indicates whether tool calling is supported.  

### `resource_meta`

The resource_meta supports the following:
* `instance_id` - (Optional, ForceNew) The instance ID.  
* `instance_name` - (Optional, ForceNew) The instance name.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Connection.
* `delete` - (Defaults to 5 mins) Used when delete the Connection.
* `update` - (Defaults to 5 mins) Used when update the Connection.

## Import

PAI Workspace Connection can be imported using the id, e.g.

```shell
$ terraform import alicloud_pai_workspace_connection.example <connection_id>
```