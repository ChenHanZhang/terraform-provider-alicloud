---
subcategory: "Open Api Explorer"
layout: "alicloud"
page_title: "Alicloud: alicloud_open_api_explorer_api_mcp_server_core"
description: |-
  Provides a Alicloud Open Api Explorer Api Mcp Server Core resource.
---

# alicloud_open_api_explorer_api_mcp_server_core

Provides a Open Api Explorer Api Mcp Server Core resource.

Provides a streamlined version of ApiMcpServer that does not require users to manually select APIs. It offers atomic interfaces to provide the model with sufficient context, enabling the model to select and orchestrate APIs.

For information about Open Api Explorer Api Mcp Server Core and how to use it, see [What is Api Mcp Server Core](https://next.api.alibabacloud.com/document/OpenAPIExplorer/2024-11-30/CreateApiMcpServerCore).

-> **NOTE:** Available since v1.277.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}


resource "alicloud_open_api_explorer_api_mcp_server_core" "default" {
  oauth_client_id             = "create-oauth-client-id"
  enable_assume_role          = true
  assume_role_name            = "create-assume-role"
  assume_role_override_policy = "{\"Version\":\"1\",\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"ecs:Describe*\",\"vpc:Describe*\"],\"Resource\":\"*\"}]}"
  vpc_whitelists              = ["vpc-create-a", "vpc-create-b", "vpc-create-c"]
  public_access_type          = "public"
  enable_custom_vpc_whitelist = true
}
```

## Argument Reference

The following arguments are supported:
* `assume_role_name` - (Optional) The name of the RAM role in the target account to assume when performing cross-account operations with multi-account access enabled.
* `assume_role_override_policy` - (Optional) An additional policy applied during role assumption when multi-account access is enabled. If this policy is specified, the permissions for the assumed role are determined by this policy, overriding the role's original permission definition.
* `enable_assume_role` - (Optional) Specifies whether to enable multi-account access.
* `enable_custom_vpc_whitelist` - (Optional) Whether to enable a custom VPC allowlist. If disabled, the resource inherits the account-level configuration.
* `oauth_client_id` - (Optional) The custom OAuth Client ID when selecting a custom OAuth configuration.  
`Supported only for Web/Native applications, and the OAuth scope must include /acs/mcp-server`.
* `public_access_type` - (Optional) Public access type: allow public access, disallow public access, or inherit account-level configuration.
* `vpc_whitelists` - (Optional, List) After public network access is disabled, this field specifies the VPC allowlist that restricts traffic sources. If not set or left empty, no source restrictions apply.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - A resource property field that represents the creation time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Api Mcp Server Core.
* `delete` - (Defaults to 5 mins) Used when delete the Api Mcp Server Core.
* `update` - (Defaults to 5 mins) Used when update the Api Mcp Server Core.

## Import

Open Api Explorer Api Mcp Server Core can be imported using the id, e.g.

```shell
$ terraform import alicloud_open_api_explorer_api_mcp_server_core.example <id>
```