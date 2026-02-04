---
subcategory: "EIAM"
layout: "alicloud"
page_title: "Alicloud: alicloud_eiam_application"
description: |-
  Provides a Alicloud EIAM Application resource.
---

# alicloud_eiam_application

Provides a EIAM Application resource.

The application in IDaaS EIAM.

For information about EIAM Application and how to use it, see [What is Application](https://next.api.alibabacloud.com/document/Eiam/2021-12-01/CreateApplication).

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_eiam_instance" "defaultiNPi8X" {
  description = "lcw-example"
}


resource "alicloud_eiam_application" "default" {
  instance_id             = alicloud_eiam_instance.defaultiNPi8X.id
  application_source_type = "urn:alibaba:idaas:app:source:standard"
  sso_type                = "saml2"
  application_name        = "example-application"
}
```

## Argument Reference

The following arguments are supported:
* `api_invoke_status` - (Optional) The status of the application's Developer API feature. Valid values include:
  - enabled: Enabled.
  - disabled: Disabled.
* `application_id` - (Optional, ForceNew, Computed) The application resource ID in IDaaS.  
* `application_name` - (Required) Application name.  
* `application_profile_mapping_attributes` - (Optional, List) List of SCIM synchronization organization field mapping configurations. See [`application_profile_mapping_attributes`](#application_profile_mapping_attributes) below.
* `application_source_type` - (Required, ForceNew) Source of the application creation. Valid values:
  - urn:alibaba:idaas:app:source:template: Application template.
  - urn:alibaba:idaas:app:source:standard: Standard protocol.
* `application_template_id` - (Optional, ForceNew) The ID of the application template associated during creation. This value is returned only when the application was created from an application template.
* `application_template_params` - (Optional, List) Application template creation parameters. These can only be specified when the application is created from a template.   See [`application_template_params`](#application_template_params) below.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `authorization_type` - (Optional) The application access authorization type. Valid values:  
  - authorize_required: Explicit authorization is required for access.  
  - default_all: All members have access permission by default.  
* `callback_provisioning_config` - (Optional, Set) Application event callback synchronization configuration. This field is required when ProvisionProtocolType is set to idaas_callback. See [`callback_provisioning_config`](#callback_provisioning_config) below.
* `description` - (Optional) Application description.
* `full_push_scopes` - (Optional, List) Full push scope.
* `grant_scopes` - (Optional, List) The set of application visibility scopes.
* `group_ids` - (Optional, List) A list of organizational units authorized for account synchronization.
* `groups` - (Optional, ForceNew, List) Group information. See [`groups`](#groups) below.
* `instance_id` - (Required, ForceNew) The instance ID of IDaaS EIAM.
* `logo_url` - (Optional) URL of the application icon.
* `network_zones` - (Optional, Set) Network zone configuration for client access policies. See [`network_zones`](#network_zones) below.
* `oidc_sso_config` - (Optional, Set) OIDC-based application SSO configuration parameters. See [`oidc_sso_config`](#oidc_sso_config) below.
* `organizational_unit_ids` - (Optional, List) A list of organizational units authorized for account synchronization.  
* `organizational_units` - (Optional, ForceNew, List) Organizational information. See [`organizational_units`](#organizational_units) below.
* `provision_password` - (Optional) Specifies whether passwords are synchronized in IDaaS user event callbacks. Valid values:
  - true: Synchronize passwords.
  - false: Do not synchronize passwords.
* `provision_protocol_type` - (Optional) The application provisioning protocol type. Valid values:  
  - idaas_callback: IDaaS custom event callback protocol.  
  - scim2: SCIM protocol.  
* `provisioning_actions` - (Optional, List) Target resource provisioning actions.
* `resource_server_identifier` - (Optional) The unique identifier of the ResourceServer, corresponding to the ResourceServer audience.
* `resource_server_scope_ids` - (Optional, ForceNew, List) A list of scope permission IDs under the ResourceServer.
* `saml_sso_config` - (Optional, Set) SAML-based application SSO attribute configuration parameters.   See [`saml_sso_config`](#saml_sso_config) below.
* `scim_provisioning_config` - (Optional, Set) IDaaS SCIM protocol provisioning configuration parameters. This parameter is required when ProvisionProtocolType is set to scim2.   See [`scim_provisioning_config`](#scim_provisioning_config) below.
* `sso_status` - (Optional) The SSO feature status of the application. Valid values include:  
  - enabled: Enabled.  
  - disabled: Disabled.  
* `sso_type` - (Required, ForceNew) Single sign-on (SSO) protocol. Valid values:
  - saml2: SAML 2.0 protocol.
  - oidc: OpenID Connect protocol.
* `status` - (Optional, Computed) Application status. Valid values include:  
  - enabled: Enabled.  
  - disabled: Disabled.  
* `user_mapping_identity_name` - (Optional) SCIM user mapping identifier.  
* `user_primary_organizational_unit_id` - (Optional) Organization ID.
* `users` - (Optional, ForceNew, List) User information. See [`users`](#users) below.

### `application_profile_mapping_attributes`

The application_profile_mapping_attributes supports the following:
* `expression_mapping_type` - (Required) The type of the expression.
* `source_value_expression` - (Required) Mapping attribute value expression.
* `target_field` - (Required) Target attribute name for mapping.
* `target_field_description` - (Optional) Target attribute name for mapping.

### `application_template_params`

The application_template_params supports the following:
* `template_param_name` - (Optional) The specific name of the application template creation parameter.  
* `template_param_value` - (Optional) The actual value of the application template creation parameter.  

### `callback_provisioning_config`

The callback_provisioning_config supports the following:
* `callback_url` - (Optional) The endpoint URL where the application receives IDaaS event callbacks.
* `encrypt_key` - (Optional) Symmetric encryption and decryption key for IDaaS event callbacks, using the AES-256 algorithm and encoded in hexadecimal format.
* `encrypt_required` - (Optional) Specifies whether IDaaS event callback messages are encrypted. Valid values:
  - true: Encrypted.
  - false: Not encrypted; messages are transmitted in plaintext.
* `listen_event_scopes` - (Optional, List) List of message types for which IDaaS event callbacks are monitored.

### `groups`

The groups supports the following:
* `group_id` - (Optional, ForceNew) Group ID.

### `network_zones`

The network_zones supports the following:
* `include_network_zones` - (Optional, List) Selected network zones.

### `oidc_sso_config`

The oidc_sso_config supports the following:
* `access_token_effective_time` - (Optional, Int) The validity period of the issued access token, in seconds. The default value is 1200 seconds (20 minutes).
* `code_effective_time` - (Optional, Int) The validity period of the issued authorization code, in seconds. The default value is 60 seconds (1 minute).
* `custom_claims` - (Optional, List) Custom user information included in the returned ID token. See [`custom_claims`](#oidc_sso_config-custom_claims) below.
* `grant_scopes` - (Optional, List) The set of application visibility scopes.
* `grant_types` - (Optional, List) A list of supported OIDC protocol grant types.
* `id_token_effective_time` - (Optional, Int) The validity period of the issued id token, in seconds. The default value is 300 seconds (5 minutes).
* `password_authentication_source_id` - (Optional) The authentication source ID used for the password grant type. This parameter takes effect only when the specified GrantTypes of the OIDC application include the password grant type.
* `password_totp_mfa_required` - (Optional) Specifies whether TOTP multi-factor authentication (MFA) is required for the password grant type. This parameter takes effect only when the specified GrantTypes of the OIDC application include the password grant type.
* `pkce_challenge_methods` - (Optional, List) Algorithm used to calculate the Code Challenge in PKCE.  
* `pkce_required` - (Optional) Whether PKCE (RFC 7636) is required for application SSO.
* `post_logout_redirect_uris` - (Optional, List) List of logout callback URIs supported by the application.  
* `redirect_uris` - (Optional, List) A list of RedirectURIs supported by the application.
* `refresh_token_effective` - (Optional, Int) Validity period of the issued refresh token, in seconds. The default value is 86 400 seconds (1 day).  
* `response_types` - (Optional, List) Response types supported by the application when OidcSsoConfig.GrantTypes includes the implicit grant type.
* `subject_id_expression` - (Optional) An expression that defines the value of the "sub" claim returned in the ID token.

### `oidc_sso_config-custom_claims`

The oidc_sso_config-custom_claims supports the following:
* `claim_name` - (Optional) The name of the returned claim.
* `claim_value_expression` - (Optional) Expression for the claim value to be returned.  

### `organizational_units`

The organizational_units supports the following:
* `organizational_unit_id` - (Optional, ForceNew) Organization ID.

### `saml_sso_config`

The saml_sso_config supports the following:
* `assertion_signed` - (Optional) Whether the Assertion requires signing. ResponseSigned and AssertionSigned cannot both be false.  
  - true: Signing is required.  
  - false: Signing is not required.  
* `attribute_statements` - (Optional, List) Additional user attribute configurations included in the SAML assertion. See [`attribute_statements`](#saml_sso_config-attribute_statements) below.
* `default_relay_state` - (Optional) The default RelayState value. When the user's single sign-on (SSO) request is initiated by EIAM, the SAML Response provided by EIAM sets RelayState to this value.  
* `name_id_format` - (Optional) The NameID format defined in the SAML standard. Valid values include:  
  - urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified: Unspecified; the application determines how to parse the NameID.  
  - urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress: Email address format.  
  - urn:oasis:names:tc:SAML:2.0:nameid-format:persistent: Persistent NameID.  
  - urn:oasis:names:tc:SAML:2.0:nameid-format:transient: Transient NameID.  
* `name_id_value_expression` - (Optional) Expression used to generate the actual value of the NameID in the SAML protocol.  
* `response_signed` - (Optional) Indicates whether the Response must be signed. ResponseSigned and AssertionSigned cannot both be false.  
  - true: Signing is required.  
  - false: Signing is not required.
* `signature_algorithm` - (Optional) The SAML assertion signature algorithm.  
* `sp_entity_id` - (Optional) The SAML EntityId of the application (SP).  
* `sp_sso_acs_url` - (Optional) The SAML assertion consumer service (ACS) URL of the application (SP).  

### `saml_sso_config-attribute_statements`

The saml_sso_config-attribute_statements supports the following:
* `attribute_name` - (Optional) The Name of the attribute in the SAML assertion.
* `attribute_value_expression` - (Optional) The attribute value expression in the SAML assertion.

### `scim_provisioning_config`

The scim_provisioning_config supports the following:
* `authn_configuration` - (Optional, Set) SCIM protocol provisioning-related configuration parameters.   See [`authn_configuration`](#scim_provisioning_config-authn_configuration) below.
* `full_push_scopes` - (Optional, List) Full push scope.
* `provisioning_actions` - (Optional, List) Target resource operation actions.
* `scim_base_url` - (Optional) The base URL at which the application accepts SCIM provisioning requests from IDaaS.

### `scim_provisioning_config-authn_configuration`

The scim_provisioning_config-authn_configuration supports the following:
* `authn_mode` - (Optional) The authorization mode for the SCIM protocol API. Valid values include:
  - oauth2: OAuth 2.0 mode.
* `authn_param` - (Optional, Set) Authorization-related configuration parameters. Usage:  
  - When GrantType is client_credentials, you can update ClientId, ClientSecret, and AuthnMethod.  
  - When GrantType is bearer_token, you can update AccessToken.   See [`authn_param`](#scim_provisioning_config-authn_configuration-authn_param) below.
* `grant_type` - (Optional) The SCIM protocol grant type. Valid values include:
  - client\_credentials: Client credentials grant type.
  - bearer\_token: Bearer token grant type.

### `scim_provisioning_config-authn_configuration-authn_param`

The scim_provisioning_config-authn_configuration-authn_param supports the following:
* `access_token` - (Optional) The access token. This field is required when using the bearer\_token authentication method.
* `authn_method` - (Optional) The SCIM protocol authentication method. Valid values include:
  - client\_secret\_basic: The secret is passed in the request header.
  - client\_secret\_post: The secret is passed in the request body.
* `client_id` - (Optional) The client ID of the application.
* `client_secret` - (Optional) The client secret of the application.
* `token_endpoint` - (Optional) The token endpoint.  

### `users`

The users supports the following:
* `user_id` - (Optional, ForceNew) The user ID.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<instance_id>:<application_id>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Application.
* `delete` - (Defaults to 5 mins) Used when delete the Application.
* `update` - (Defaults to 5 mins) Used when update the Application.

## Import

EIAM Application can be imported using the id, e.g.

```shell
$ terraform import alicloud_eiam_application.example <instance_id>:<application_id>
```