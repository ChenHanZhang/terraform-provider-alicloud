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
* `api_invoke_status` - (Optional) The Developer API feature status of the application. Valid values:  
  - enabled: Enabled.  
  - disabled: Disabled.
* `application_id` - (Optional, ForceNew, Computed) The ID of the application resource in IDaaS.
* `application_name` - (Required) Application name.
* `application_profile_mapping_attributes` - (Optional, List) A list of SCIM synchronization organization field mapping configurations. See [`application_profile_mapping_attributes`](#application_profile_mapping_attributes) below.
* `application_source_type` - (Required, ForceNew) Source of application creation. Valid values:
  - urn:alibaba:idaas:app:source:template: Application template.
  - urn:alibaba:idaas:app:source:standard: Standard protocol.
* `application_template_id` - (Optional, ForceNew) The ID of the application template associated during creation. This value is returned only when the application was created from an application template.
* `application_template_params` - (Optional, List) Application template creation parameters. These can be specified only when the application is created from a template. See [`application_template_params`](#application_template_params) below.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `authorization_type` - (Optional) Application access authorization type. Valid values:
  - authorize_required: Explicit authorization is required for access.
  - default_all: All members have access by default.
* `callback_provisioning_config` - (Optional, Set) Application event callback synchronization configuration. This field is required when ProvisionProtocolType is set to idaas_callback. See [`callback_provisioning_config`](#callback_provisioning_config) below.
* `description` - (Optional) Application description.
* `full_push_scopes` - (Optional, List) Full push scope.
* `grant_scopes` - (Optional, List) The set of application visibility scopes.
* `group_ids` - (Optional, List) A list of organizational units authorized for account synchronization.
* `groups` - (Optional, ForceNew, List) Group information. See [`groups`](#groups) below.
* `instance_id` - (Required, ForceNew) The instance ID of IDaaS EIAM.  
* `logo_url` - (Optional) URL of the application icon.
* `network_zones` - (Optional, Set) Network zone configuration for client access policies. See [`network_zones`](#network_zones) below.
* `oidc_sso_config` - (Optional, Set) Configuration parameters for the application's SSO attributes using the OpenID Connect (OIDC) protocol. See [`oidc_sso_config`](#oidc_sso_config) below.
* `organizational_unit_ids` - (Optional, List) List of organizational units included in account synchronization authorization.  
* `organizational_units` - (Optional, ForceNew, List) Organizational unit information. See [`organizational_units`](#organizational_units) below.
* `provision_password` - (Optional) Specifies whether the IDaaS user event callback synchronizes passwords. Valid values:  
  - true: Synchronize passwords.  
  - false: Do not synchronize passwords.  
* `provision_protocol_type` - (Optional) Application provisioning protocol type. Valid values:
  - idaas_callback: IDaaS custom event callback protocol.
  - scim2: SCIM protocol.
* `provisioning_actions` - (Optional, List) Target resource provisioning actions.
* `resource_server_identifier` - (Optional) Unique identifier of the Resource Server, corresponding to the Resource Server audience.
* `resource_server_scope_ids` - (Optional, ForceNew, List) A list of scope permission IDs under the Resource Server.  
* `saml_sso_config` - (Optional, Set) Configuration parameters for SAML-based SSO attributes of the application. See [`saml_sso_config`](#saml_sso_config) below.
* `scim_provisioning_config` - (Optional, Set) IDaaS SCIM protocol provisioning configuration parameters. This parameter is required when ProvisionProtocolType is set to scim2. See [`scim_provisioning_config`](#scim_provisioning_config) below.
* `sso_status` - (Optional) The SSO feature status of the application. Valid values:  
  - enabled: Enabled.  
  - disabled: Disabled.  
* `sso_type` - (Required, ForceNew) Single sign-on (SSO) protocol. Valid values:
  - saml2: SAML 2.0 protocol.
  - oidc: OpenID Connect protocol.
* `status` - (Optional, Computed) The application status. Valid values:  
  - enabled: Enabled.  
  - disabled: Disabled.  
* `user_mapping_identity_name` - (Optional) SCIM user mapping identifier.
* `user_primary_organizational_unit_id` - (Optional) The organizational unit ID.
* `users` - (Optional, ForceNew, List) User information. See [`users`](#users) below.

### `application_profile_mapping_attributes`

The application_profile_mapping_attributes supports the following:
* `expression_mapping_type` - (Required) The type of expression.
* `source_value_expression` - (Required) Expression for the value of the mapped attribute.
* `target_field` - (Required) Name of the target attribute for mapping.
* `target_field_description` - (Optional) Name of the target attribute for mapping.

### `application_template_params`

The application_template_params supports the following:
* `template_param_name` - (Optional) The specific name of an application template creation parameter.
* `template_param_value` - (Optional) Actual value of the application template creation parameter.

### `callback_provisioning_config`

The callback_provisioning_config supports the following:
* `callback_url` - (Optional) The endpoint URL where the application receives IDaaS event callbacks.
* `encrypt_key` - (Optional) Symmetric encryption and decryption key for IDaaS event callbacks, using the AES-256 algorithm and formatted in hexadecimal encoding.
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
* `access_token_effective_time` - (Optional, Int) The validity period of the issued access token, in seconds. The default value is 1 200 seconds (20 minutes).
* `code_effective_time` - (Optional, Int) The validity period of the issued authorization code, in seconds. The default value is 60 seconds (1 minute).  
* `custom_claims` - (Optional, List) Custom user information included in the ID token response. See [`custom_claims`](#oidc_sso_config-custom_claims) below.
* `grant_scopes` - (Optional, List) The set of API authorization scopes visible to the application.
* `grant_types` - (Optional, List) A list of OIDC grant types supported by the application.  
* `id_token_effective_time` - (Optional, Int) The validity period of the issued ID token, in seconds. The default value is 300 seconds (5 minutes).
* `password_authentication_source_id` - (Optional) The authentication source ID used for the password grant type. This setting takes effect only when the OIDC application's GrantTypes includes the password grant type.
* `password_totp_mfa_required` - (Optional) Specifies whether TOTP multi-factor authentication (MFA) is required for the password grant type. This setting takes effect only when the OIDC application's GrantTypes includes the password grant type.
* `pkce_challenge_methods` - (Optional, List) The algorithm used to compute the Code Challenge in PKCE.
* `pkce_required` - (Optional) Whether PKCE (RFC 7636) is required for the application's SSO.
* `post_logout_redirect_uris` - (Optional, List) A list of logout callback URIs supported by the application.
* `redirect_uris` - (Optional, List) A list of Redirect URIs supported by the application.  
* `refresh_token_effective` - (Optional, Int) The validity period of the issued refresh token, in seconds. The default value is 86 400 seconds (1 day).
* `response_types` - (Optional, List) The response types supported by the application when OidcSsoConfig.GrantTypes includes the implicit grant type.
* `subject_id_expression` - (Optional) Custom expression for the "sub" claim value returned in the ID token.

### `oidc_sso_config-custom_claims`

The oidc_sso_config-custom_claims supports the following:
* `claim_name` - (Optional) The name of the claim to be returned.
* `claim_value_expression` - (Optional) The expression used to determine the value of the claim to be returned.

### `organizational_units`

The organizational_units supports the following:
* `organizational_unit_id` - (Optional, ForceNew) The organizational unit ID.

### `saml_sso_config`

The saml_sso_config supports the following:
* `assertion_signed` - (Optional) Specifies whether the assertion requires signing. ResponseSigned and AssertionSigned cannot both be false.  
  - true: Signing is required.  
  - false: Signing is not required.  
* `attribute_statements` - (Optional, List) Additional user attribute configurations included in the SAML assertion.   See [`attribute_statements`](#saml_sso_config-attribute_statements) below.
* `default_relay_state` - (Optional) The default RelayState value. When the user's SSO request is initiated by EIAM, the SAML Response provided by EIAM sets RelayState to this value.
* `name_id_format` - (Optional) The NameID format defined in the SAML protocol standard. Valid values include:  
  - urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified: Unspecified; the application determines how to parse the NameID.  
  - urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress: Email address format.  
  - urn:oasis:names:tc:SAML:2.0:nameid-format:persistent: Persistent NameID.  
  - urn:oasis:names:tc:SAML:2.0:nameid-format:transient: Transient NameID.  
* `name_id_value_expression` - (Optional) The expression used to generate the actual value of the NameID in the SAML protocol.
* `response_signed` - (Optional) Indicates whether the Response must be signed. ResponseSigned and AssertionSigned cannot both be false.  
  - true: Signing is required.  
  - false: Signing is not required.
* `signature_algorithm` - (Optional) The signature algorithm used for SAML assertions.
* `sp_entity_id` - (Optional) The SAML EntityId of the service provider (SP).
* `sp_sso_acs_url` - (Optional) The SAML assertion consumer service (ACS) URL of the service provider (SP).

### `saml_sso_config-attribute_statements`

The saml_sso_config-attribute_statements supports the following:
* `attribute_name` - (Optional) The name of the attribute in the SAML assertion.
* `attribute_value_expression` - (Optional) The attribute value expression used in the SAML assertion.  

### `scim_provisioning_config`

The scim_provisioning_config supports the following:
* `authn_configuration` - (Optional, Set) Configuration parameters related to SCIM protocol provisioning.   See [`authn_configuration`](#scim_provisioning_config-authn_configuration) below.
* `full_push_scopes` - (Optional, List) Full push scope.
* `provisioning_actions` - (Optional, List) Target resource provisioning actions.
* `scim_base_url` - (Optional) The base URL of the application that accepts SCIM provisioning requests from IDaaS.

### `scim_provisioning_config-authn_configuration`

The scim_provisioning_config-authn_configuration supports the following:
* `authn_mode` - (Optional) The authentication mode for SCIM protocol endpoints. Valid values:
  - oauth2: OAuth2 mode.
* `authn_param` - (Optional, Set) Authorization-related configuration parameters. Usage is as follows:  
  - When GrantType is client_credentials, you can update ClientId, ClientSecret, and AuthnMethod.  
  - When GrantType is bearer_token, you can update AccessToken.   See [`authn_param`](#scim_provisioning_config-authn_configuration-authn_param) below.
* `grant_type` - (Optional) The SCIM protocol grant type. Valid values:
  - client_credentials: Client credentials grant.
  - bearer_token: Bearer token grant.

### `scim_provisioning_config-authn_configuration-authn_param`

The scim_provisioning_config-authn_configuration-authn_param supports the following:
* `access_token` - (Optional) Access token. This field is required when using the bearer_token authentication method.  
* `authn_method` - (Optional) SCIM protocol authentication method. Valid values include:  
  - client_secret_basic: Pass the secret through the request header.  
  - client_secret_post: Pass the secret through the request body.  
* `client_id` - (Optional) The client ID of the application.
* `client_secret` - (Optional) The client secret of the application.
* `token_endpoint` - (Optional) Token endpoint.  

### `users`

The users supports the following:
* `user_id` - (Optional, ForceNew) User ID.

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