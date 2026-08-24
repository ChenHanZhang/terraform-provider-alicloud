---
subcategory: "Cloud Native API Gateway (APIG)"
layout: "alicloud"
page_title: "Alicloud: alicloud_apig_route"
description: |-
  Provides a Alicloud APIG Route resource.
---

# alicloud_apig_route

Provides a APIG Route resource.



For information about APIG Route and how to use it, see [What is Route](https://next.api.alibabacloud.com/document/APIG/2024-03-27/CreateHttpApiRoute).

-> **NOTE:** Available since v1.287.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_apig_http_api" "route_httpapi_arr" {
  http_api_name = "cspec-route-arr-httpapi"
  protocols     = ["HTTP"]
  type          = "Rest"
  description   = "array example httpapi"
  base_path     = "/cspec-route-arr"
}


resource "alicloud_apig_route" "default" {
  backend {
    services {
      version    = "v1.0"
      port       = "8080"
      protocol   = "HTTP"
      weight     = "70"
      service_id = "svc-a"
    }
    services {
      version    = "v1.0"
      port       = "8081"
      protocol   = "HTTP"
      weight     = "30"
      service_id = "svc-b"
    }
    scene = "SingleService"
  }
  description = "array route description"
  route_name  = "cspec-arr-route"
  domain_infos {
  }
  domain_infos {
  }
  http_api_id = alicloud_apig_http_api.route_httpapi_arr.id
  match {
    path {
      type  = "Prefix"
      value = "/arr-path"
    }
    headers {
      type  = "Exact"
      value = "v1"
      name  = "h1"
    }
    headers {
      type  = "Prefix"
      value = "v2"
      name  = "h2"
    }
    query_params {
      type  = "Exact"
      value = "v1"
      name  = "q1"
    }
    query_params {
      type  = "Prefix"
      value = "v2"
      name  = "q2"
    }
    methods         = ["GET", "POST", "PUT"]
    ignore_uri_case = false
  }
}
```

## Argument Reference

The following arguments are supported:
* `backend` - (Optional, Set) Backend service. See [`backend`](#backend) below.
* `description` - (Optional) The description of the route.
* `domain_ids` - (Optional, List) The domain name IDs.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `environment_info` - (Optional, ForceNew, Set) The environment information of the route. See [`environment_info`](#environment_info) below.
* `http_api_id` - (Optional, ForceNew, Computed) The ID of the HTTP API to which the route belongs.
* `match` - (Optional, Set) The route match rule. See [`match`](#match) below.
* `route_name` - (Optional) The name of the route.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.


### `backend`

The backend supports the following:
* `scene` - (Optional) The backend service scenario.
  - SingleService: Single service.
  - MultiServiceByRatio: Canary release across multiple services by ratio.
  - Mock: Mock service.
  - Redirect: Redirect service.
* `services` - (Optional, List) Backend service. See [`services`](#backend-services) below.

### `backend-services`

The backend-services supports the following:
* `group_name` - (Optional, ForceNew) service group
* `http_dubbo_transcoder` - (Optional, ForceNew, Set) HTTP to Dubbo transcoder configuration See [`http_dubbo_transcoder`](#backend-services-http_dubbo_transcoder) below.
* `model_name` - (Optional) Target model name shared by multiple existing model backend scenes. Its routing or model-rewrite semantics depend on Backend.Scene. It is required for SemanticRouter; AiAutoRouter uses the AI service default model when it is omitted.
* `name` - (Optional, ForceNew) The name of the service.
* `namespace` - (Optional, ForceNew) service namespace
* `port` - (Optional, ForceNew, Int) Service port. Do not specify this parameter for dynamic ports.
* `protocol` - (Optional, ForceNew) Service protocol. Valid values: HTTP, TCP, and DUBBO.
* `service_id` - (Optional) Service ID.
* `source_type` - (Optional, ForceNew) service source type
* `version` - (Optional, ForceNew) Service version.
* `weight` - (Optional, Int) The percentage value of the traffic ratio. You can specify the weight of the service when the scenario is proportional (canary) routing. This parameter is not required in other scenarios.

### `backend-services-http_dubbo_transcoder`

The backend-services-http_dubbo_transcoder supports the following:
* `dubbo_service_group` - (Optional, ForceNew) Dubbo service group
* `dubbo_service_name` - (Optional, ForceNew) Dubbo service name
* `dubbo_service_version` - (Optional, ForceNew) Dubbo service version
* `method_map_list` - (Optional, ForceNew, List) method mapping list See [`method_map_list`](#backend-services-http_dubbo_transcoder-method_map_list) below.

### `backend-services-http_dubbo_transcoder-method_map_list`

The backend-services-http_dubbo_transcoder-method_map_list supports the following:
* `dubbo_method_name` - (Optional, ForceNew) Dubbo method name
* `http_method` - (Optional, ForceNew) HTTP method
* `method_path` - (Optional, ForceNew) method path
* `param_maps_list` - (Optional, ForceNew, List) parameter mapping list See [`param_maps_list`](#backend-services-http_dubbo_transcoder-method_map_list-param_maps_list) below.
* `pass_through_all_headers` - (Optional, ForceNew) header passthrough mode
* `pass_through_list` - (Optional, ForceNew, List) passthrough header list

### `backend-services-http_dubbo_transcoder-method_map_list-param_maps_list`

The backend-services-http_dubbo_transcoder-method_map_list-param_maps_list supports the following:
* `extract_key` - (Optional, ForceNew) parameter extraction key
* `extract_key_spec` - (Optional, ForceNew) parameter extraction location
* `mapping_type` - (Optional, ForceNew) backend parameter type

### `environment_info`

The environment_info supports the following:
* `environment_id` - (Optional, ForceNew) The environment ID.

### `match`

The match supports the following:
* `headers` - (Optional, List) The list of HTTP request header matching rules. See [`headers`](#match-headers) below.
* `ignore_uri_case` - (Optional) Specifies whether the path is case-sensitive.
* `methods` - (Optional, List) The request method. Valid values: GET, HEAD, POST, PUT, DELETE, CONNECT, OPTION, TRACE, and PATCH.
* `path` - (Optional, Set) The path rule. See [`path`](#match-path) below.
* `query_params` - (Optional, List) The matching rules for query parameters. See [`query_params`](#match-query_params) below.

### `match-headers`

The match-headers supports the following:
* `name` - (Optional) The header name.
* `type` - (Optional) The header matching rule type. Valid values: Exact (exact match), Prefix (prefix match), and Regex (regular expression match).
* `value` - (Optional) The header value.

### `match-path`

The match-path supports the following:
* `type` - (Optional) The path matching type. Valid values: Exact (exact match), Prefix (prefix match), and Regex (regular expression match).
* `value` - (Optional) The path value.

### `match-query_params`

The match-query_params supports the following:
* `name` - (Optional) The parameter name.
* `type` - (Optional) The matching rule for the query parameter. Valid values: Exact (exact match), Prefix (prefix match), and Regex (regular expression match).
* `value` - (Optional) The parameter value.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<http_api_id>:<route_id>`.
* `backend` - Backend service.
  * `enable_system_models` - Read-only projection of whether the Model API that generated this route enabled platform system models.
* `builtin` - Indicates whether the route is a built-in route.
* `create_time` - The creation time in UTC format: yyyy-MM-ddTHH:mm:ssZ.
* `environment_info` - The environment information of the route.
  * `alias` - The alias of the environment name.
  * `gateway_info` - The gateway instance information corresponding to the environment.
    * `gateway_edition` - The edition of the gateway instance.
    * `gateway_id` - The ID of the Cloud-native API Gateway.
    * `name` - The name of the gateway.
  * `name` - The environment name.
  * `sub_domains` - The default second-level domain names of the environment.
    * `domain_id` - The ID of the second-level domain name.
    * `name` - The name of the second-level domain name.
    * `network_type` - The domain access type, such as Intranet or Internet.
    * `protocol` - The domain protocol, such as HTTP or HTTPS.
* `route_id` - The route ID.
* `status` - The deployment status of the route.
* `update_time` - The update time in Greenwich Mean Time (GMT).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Route.
* `delete` - (Defaults to 5 mins) Used when delete the Route.
* `update` - (Defaults to 5 mins) Used when update the Route.

## Import

APIG Route can be imported using the id, e.g.

```shell
$ terraform import alicloud_apig_route.example <http_api_id>:<route_id>
```