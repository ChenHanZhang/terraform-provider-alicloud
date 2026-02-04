// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudEiamApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudEiamApplicationCreate,
		Read:   resourceAliCloudEiamApplicationRead,
		Update: resourceAliCloudEiamApplicationUpdate,
		Delete: resourceAliCloudEiamApplicationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"api_invoke_status": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"disabled", "enabled"}, false),
			},
			"application_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"application_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"application_profile_mapping_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expression_mapping_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"source_value_expression": {
							Type:     schema.TypeString,
							Required: true,
						},
						"target_field_description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"target_field": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"application_source_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"application_template_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"application_template_params": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_param_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"template_param_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"authorization_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"callback_provisioning_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"callback_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"encrypt_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"listen_event_scopes": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"encrypt_required": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"full_push_scopes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"grant_scopes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"group_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"groups": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"logo_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_zones": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"include_network_zones": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"oidc_sso_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"response_types": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"password_authentication_source_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"password_totp_mfa_required": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"subject_id_expression": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"custom_claims": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"claim_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"claim_value_expression": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"pkce_challenge_methods": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"post_logout_redirect_uris": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"refresh_token_effective": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"pkce_required": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"access_token_effective_time": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"grant_scopes": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"id_token_effective_time": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"redirect_uris": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"code_effective_time": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"grant_types": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"organizational_unit_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"organizational_units": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"organizational_unit_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"provision_password": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"provision_protocol_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"provisioning_actions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"resource_server_identifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_server_scope_ids": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"saml_sso_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name_id_value_expression": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"response_signed": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"sp_entity_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sp_sso_acs_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"default_relay_state": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"signature_algorithm": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"assertion_signed": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"name_id_format": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"attribute_statements": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attribute_value_expression": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"attribute_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"scim_provisioning_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authn_configuration": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authn_param": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"token_endpoint": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"authn_method": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"access_token": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"client_secret": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"client_id": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"grant_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"authn_mode": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"scim_base_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"provisioning_actions": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"full_push_scopes": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"sso_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"oidc", "saml2"}, false),
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"enabled", "disabled", "deleted"}, false),
			},
			"user_mapping_identity_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_primary_organizational_unit_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"users": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceAliCloudEiamApplicationCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateApplication"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	request["RegionId"] = client.RegionId

	request["ApplicationName"] = d.Get("application_name")
	if v, ok := d.GetOk("logo_url"); ok {
		request["LogoUrl"] = v
	}
	request["SsoType"] = d.Get("sso_type")
	request["ApplicationSourceType"] = d.Get("application_source_type")
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("application_template_id"); ok {
		request["ApplicationTemplateId"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_eiam_application", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["InstanceId"], response["ApplicationId"]))

	action = "AuthorizeApplicationToUsers"
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	if v, ok := d.GetOk("application_id"); ok {
		request["ApplicationId"] = v
	}
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("users"); ok {
		localData, err := jsonpath.Get("$[*].user_id", v)
		if err != nil {
			return WrapError(err)
		}
		localDataArray := convertToInterfaceArray(localData)

		request["UserIds"] = localDataArray
	}

	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_eiam_application", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["InstanceId"], request["ApplicationId"]))

	action = "AuthorizeApplicationToGroups"
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	if v, ok := d.GetOk("application_id"); ok {
		request["ApplicationId"] = v
	}
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("groups"); ok {
		localData, err := jsonpath.Get("$[*].group_id", v)
		if err != nil {
			return WrapError(err)
		}
		localDataArray := convertToInterfaceArray(localData)

		request["GroupIds"] = localDataArray
	}

	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_eiam_application", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["InstanceId"], request["ApplicationId"]))

	action = "AuthorizeApplicationToOrganizationalUnits"
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	if v, ok := d.GetOk("application_id"); ok {
		request["ApplicationId"] = v
	}
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("organizational_units"); ok {
		localData, err := jsonpath.Get("$[*].organizational_unit_id", v)
		if err != nil {
			return WrapError(err)
		}
		localDataArray := convertToInterfaceArray(localData)

		request["OrganizationalUnitIds"] = localDataArray
	}

	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_eiam_application", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["InstanceId"], request["ApplicationId"]))

	action = "AuthorizeResourceServerScopesToGroup"
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	if v, ok := d.GetOk("application_id"); ok {
		request["ApplicationId"] = v
	}
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_eiam_application", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["InstanceId"], request["ApplicationId"]))

	action = "AuthorizeResourceServerScopesToUser"
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	if v, ok := d.GetOk("application_id"); ok {
		request["ApplicationId"] = v
	}
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_eiam_application", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["InstanceId"], request["ApplicationId"]))

	action = "AuthorizeResourceServerScopesToOrganizationalUnit"
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	if v, ok := d.GetOk("application_id"); ok {
		request["ApplicationId"] = v
	}
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_eiam_application", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["InstanceId"], request["ApplicationId"]))

	action = "GenerateOauthToken"
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	if v, ok := d.GetOk("application_id"); ok {
		request["ApplicationId"] = v
	}
	request["RegionId"] = client.RegionId

	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_eiam_application", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["InstanceId"], request["ApplicationId"]))

	return resourceAliCloudEiamApplicationUpdate(d, meta)
}

func resourceAliCloudEiamApplicationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	eiamServiceV2 := EiamServiceV2{client}

	objectRaw, err := eiamServiceV2.DescribeEiamApplication(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_eiam_application DescribeEiamApplication Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("sso_status", objectRaw["SsoStatus"])

	oidcSsoConfigMaps := make([]map[string]interface{}, 0)
	oidcSsoConfigMap := make(map[string]interface{})
	oidcSsoConfigRaw := make(map[string]interface{})
	if objectRaw["OidcSsoConfig"] != nil {
		oidcSsoConfigRaw = objectRaw["OidcSsoConfig"].(map[string]interface{})
	}
	if len(oidcSsoConfigRaw) > 0 {
		oidcSsoConfigMap["access_token_effective_time"] = oidcSsoConfigRaw["AccessTokenEffectiveTime"]
		oidcSsoConfigMap["code_effective_time"] = oidcSsoConfigRaw["CodeEffectiveTime"]
		oidcSsoConfigMap["id_token_effective_time"] = oidcSsoConfigRaw["IdTokenEffectiveTime"]
		oidcSsoConfigMap["password_authentication_source_id"] = oidcSsoConfigRaw["PasswordAuthenticationSourceId"]
		oidcSsoConfigMap["password_totp_mfa_required"] = oidcSsoConfigRaw["PasswordTotpMfaRequired"]
		oidcSsoConfigMap["pkce_required"] = oidcSsoConfigRaw["PkceRequired"]
		oidcSsoConfigMap["refresh_token_effective"] = oidcSsoConfigRaw["RefreshTokenEffective"]
		oidcSsoConfigMap["subject_id_expression"] = oidcSsoConfigRaw["SubjectIdExpression"]

		customClaimsRaw := oidcSsoConfigRaw["CustomClaims"]
		customClaimsMaps := make([]map[string]interface{}, 0)
		if customClaimsRaw != nil {
			for _, customClaimsChildRaw := range convertToInterfaceArray(customClaimsRaw) {
				customClaimsMap := make(map[string]interface{})
				customClaimsChildRaw := customClaimsChildRaw.(map[string]interface{})
				customClaimsMap["claim_name"] = customClaimsChildRaw["ClaimName"]
				customClaimsMap["claim_value_expression"] = customClaimsChildRaw["ClaimValueExpression"]

				customClaimsMaps = append(customClaimsMaps, customClaimsMap)
			}
		}
		oidcSsoConfigMap["custom_claims"] = customClaimsMaps
		grantScopesRaw := make([]interface{}, 0)
		if oidcSsoConfigRaw["GrantScopes"] != nil {
			grantScopesRaw = convertToInterfaceArray(oidcSsoConfigRaw["GrantScopes"])
		}

		oidcSsoConfigMap["grant_scopes"] = grantScopesRaw
		grantTypesRaw := make([]interface{}, 0)
		if oidcSsoConfigRaw["GrantTypes"] != nil {
			grantTypesRaw = convertToInterfaceArray(oidcSsoConfigRaw["GrantTypes"])
		}

		oidcSsoConfigMap["grant_types"] = grantTypesRaw
		pkceChallengeMethodsRaw := make([]interface{}, 0)
		if oidcSsoConfigRaw["PkceChallengeMethods"] != nil {
			pkceChallengeMethodsRaw = convertToInterfaceArray(oidcSsoConfigRaw["PkceChallengeMethods"])
		}

		oidcSsoConfigMap["pkce_challenge_methods"] = pkceChallengeMethodsRaw
		postLogoutRedirectUrisRaw := make([]interface{}, 0)
		if oidcSsoConfigRaw["PostLogoutRedirectUris"] != nil {
			postLogoutRedirectUrisRaw = convertToInterfaceArray(oidcSsoConfigRaw["PostLogoutRedirectUris"])
		}

		oidcSsoConfigMap["post_logout_redirect_uris"] = postLogoutRedirectUrisRaw
		redirectUrisRaw := make([]interface{}, 0)
		if oidcSsoConfigRaw["RedirectUris"] != nil {
			redirectUrisRaw = convertToInterfaceArray(oidcSsoConfigRaw["RedirectUris"])
		}

		oidcSsoConfigMap["redirect_uris"] = redirectUrisRaw
		responseTypesRaw := make([]interface{}, 0)
		if oidcSsoConfigRaw["ResponseTypes"] != nil {
			responseTypesRaw = convertToInterfaceArray(oidcSsoConfigRaw["ResponseTypes"])
		}

		oidcSsoConfigMap["response_types"] = responseTypesRaw
		oidcSsoConfigMaps = append(oidcSsoConfigMaps, oidcSsoConfigMap)
	}
	if err := d.Set("oidc_sso_config", oidcSsoConfigMaps); err != nil {
		return err
	}
	samlSsoConfigMaps := make([]map[string]interface{}, 0)
	samlSsoConfigMap := make(map[string]interface{})
	samlSsoConfigRaw := make(map[string]interface{})
	if objectRaw["SamlSsoConfig"] != nil {
		samlSsoConfigRaw = objectRaw["SamlSsoConfig"].(map[string]interface{})
	}
	if len(samlSsoConfigRaw) > 0 {
		samlSsoConfigMap["assertion_signed"] = samlSsoConfigRaw["AssertionSigned"]
		samlSsoConfigMap["default_relay_state"] = samlSsoConfigRaw["DefaultRelayState"]
		samlSsoConfigMap["name_id_format"] = samlSsoConfigRaw["NameIdFormat"]
		samlSsoConfigMap["name_id_value_expression"] = samlSsoConfigRaw["NameIdValueExpression"]
		samlSsoConfigMap["response_signed"] = samlSsoConfigRaw["ResponseSigned"]
		samlSsoConfigMap["signature_algorithm"] = samlSsoConfigRaw["SignatureAlgorithm"]
		samlSsoConfigMap["sp_entity_id"] = samlSsoConfigRaw["SpEntityId"]
		samlSsoConfigMap["sp_sso_acs_url"] = samlSsoConfigRaw["SpSsoAcsUrl"]

		attributeStatementsRaw := samlSsoConfigRaw["AttributeStatements"]
		attributeStatementsMaps := make([]map[string]interface{}, 0)
		if attributeStatementsRaw != nil {
			for _, attributeStatementsChildRaw := range convertToInterfaceArray(attributeStatementsRaw) {
				attributeStatementsMap := make(map[string]interface{})
				attributeStatementsChildRaw := attributeStatementsChildRaw.(map[string]interface{})
				attributeStatementsMap["attribute_name"] = attributeStatementsChildRaw["AttributeName"]
				attributeStatementsMap["attribute_value_expression"] = attributeStatementsChildRaw["AttributeValueExpression"]

				attributeStatementsMaps = append(attributeStatementsMaps, attributeStatementsMap)
			}
		}
		samlSsoConfigMap["attribute_statements"] = attributeStatementsMaps
		samlSsoConfigMaps = append(samlSsoConfigMaps, samlSsoConfigMap)
	}
	if err := d.Set("saml_sso_config", samlSsoConfigMaps); err != nil {
		return err
	}

	objectRaw, err = eiamServiceV2.DescribeApplicationGetApplication(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("api_invoke_status", objectRaw["ApiInvokeStatus"])
	d.Set("application_name", objectRaw["ApplicationName"])
	d.Set("application_source_type", objectRaw["ApplicationSourceType"])
	d.Set("application_template_id", objectRaw["ApplicationTemplateId"])
	d.Set("authorization_type", objectRaw["AuthorizationType"])
	d.Set("description", objectRaw["Description"])
	d.Set("logo_url", objectRaw["LogoUrl"])
	d.Set("resource_server_identifier", objectRaw["ResourceServerIdentifier"])
	d.Set("sso_type", objectRaw["SsoType"])
	d.Set("status", objectRaw["Status"])
	d.Set("application_id", objectRaw["ApplicationId"])
	d.Set("instance_id", objectRaw["InstanceId"])

	objectRaw, err = eiamServiceV2.DescribeApplicationGetApplicationProvisioningConfig(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("provision_password", objectRaw["ProvisionPassword"])
	d.Set("provision_protocol_type", objectRaw["ProvisionProtocolType"])
	d.Set("status", objectRaw["Status"])
	d.Set("application_id", objectRaw["ApplicationId"])
	d.Set("instance_id", objectRaw["InstanceId"])

	callbackProvisioningConfigMaps := make([]map[string]interface{}, 0)
	callbackProvisioningConfigMap := make(map[string]interface{})
	callbackProvisioningConfigRaw := make(map[string]interface{})
	if objectRaw["CallbackProvisioningConfig"] != nil {
		callbackProvisioningConfigRaw = objectRaw["CallbackProvisioningConfig"].(map[string]interface{})
	}
	if len(callbackProvisioningConfigRaw) > 0 {
		callbackProvisioningConfigMap["callback_url"] = callbackProvisioningConfigRaw["CallbackUrl"]
		callbackProvisioningConfigMap["encrypt_key"] = callbackProvisioningConfigRaw["EncryptKey"]
		callbackProvisioningConfigMap["encrypt_required"] = callbackProvisioningConfigRaw["EncryptRequired"]

		listenEventScopesRaw := make([]interface{}, 0)
		if callbackProvisioningConfigRaw["ListenEventScopes"] != nil {
			listenEventScopesRaw = convertToInterfaceArray(callbackProvisioningConfigRaw["ListenEventScopes"])
		}

		callbackProvisioningConfigMap["listen_event_scopes"] = listenEventScopesRaw
		callbackProvisioningConfigMaps = append(callbackProvisioningConfigMaps, callbackProvisioningConfigMap)
	}
	if err := d.Set("callback_provisioning_config", callbackProvisioningConfigMaps); err != nil {
		return err
	}
	scimProvisioningConfigMaps := make([]map[string]interface{}, 0)
	scimProvisioningConfigMap := make(map[string]interface{})
	scimProvisioningConfigRaw := make(map[string]interface{})
	if objectRaw["ScimProvisioningConfig"] != nil {
		scimProvisioningConfigRaw = objectRaw["ScimProvisioningConfig"].(map[string]interface{})
	}
	if len(scimProvisioningConfigRaw) > 0 {
		scimProvisioningConfigMap["scim_base_url"] = scimProvisioningConfigRaw["ScimBaseUrl"]

		authnConfigurationMaps := make([]map[string]interface{}, 0)
		authnConfigurationMap := make(map[string]interface{})
		authnConfigurationRaw := make(map[string]interface{})
		if scimProvisioningConfigRaw["AuthnConfiguration"] != nil {
			authnConfigurationRaw = scimProvisioningConfigRaw["AuthnConfiguration"].(map[string]interface{})
		}
		if len(authnConfigurationRaw) > 0 {
			authnConfigurationMap["authn_mode"] = authnConfigurationRaw["AuthnMode"]
			authnConfigurationMap["grant_type"] = authnConfigurationRaw["GrantType"]

			authnParamMaps := make([]map[string]interface{}, 0)
			authnParamMap := make(map[string]interface{})
			authnParamRaw := make(map[string]interface{})
			if authnConfigurationRaw["AuthnParam"] != nil {
				authnParamRaw = authnConfigurationRaw["AuthnParam"].(map[string]interface{})
			}
			if len(authnParamRaw) > 0 {
				authnParamMap["access_token"] = authnParamRaw["AccessToken"]
				authnParamMap["authn_method"] = authnParamRaw["AuthnMethod"]
				authnParamMap["client_id"] = authnParamRaw["ClientId"]
				authnParamMap["client_secret"] = authnParamRaw["ClientSecret"]
				authnParamMap["token_endpoint"] = authnParamRaw["TokenEndpoint"]

				authnParamMaps = append(authnParamMaps, authnParamMap)
			}
			authnConfigurationMap["authn_param"] = authnParamMaps
			authnConfigurationMaps = append(authnConfigurationMaps, authnConfigurationMap)
		}
		scimProvisioningConfigMap["authn_configuration"] = authnConfigurationMaps
		fullPushScopesRaw := make([]interface{}, 0)
		if scimProvisioningConfigRaw["FullPushScopes"] != nil {
			fullPushScopesRaw = convertToInterfaceArray(scimProvisioningConfigRaw["FullPushScopes"])
		}

		scimProvisioningConfigMap["full_push_scopes"] = fullPushScopesRaw
		provisioningActionsRaw := make([]interface{}, 0)
		if scimProvisioningConfigRaw["ProvisioningActions"] != nil {
			provisioningActionsRaw = convertToInterfaceArray(scimProvisioningConfigRaw["ProvisioningActions"])
		}

		scimProvisioningConfigMap["provisioning_actions"] = provisioningActionsRaw
		scimProvisioningConfigMaps = append(scimProvisioningConfigMaps, scimProvisioningConfigMap)
	}
	if err := d.Set("scim_provisioning_config", scimProvisioningConfigMaps); err != nil {
		return err
	}

	objectRaw, err = eiamServiceV2.DescribeApplicationGetApplicationUserProfileMapping(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("user_mapping_identity_name", objectRaw["UserMappingIdentityName"])
	d.Set("application_id", objectRaw["ApplicationId"])
	d.Set("instance_id", objectRaw["InstanceId"])

	applicationProfileMappingAttributesRaw := objectRaw["ApplicationProfileMappingAttributes"]
	applicationProfileMappingAttributesMaps := make([]map[string]interface{}, 0)
	if applicationProfileMappingAttributesRaw != nil {
		for _, applicationProfileMappingAttributesChildRaw := range convertToInterfaceArray(applicationProfileMappingAttributesRaw) {
			applicationProfileMappingAttributesMap := make(map[string]interface{})
			applicationProfileMappingAttributesChildRaw := applicationProfileMappingAttributesChildRaw.(map[string]interface{})
			applicationProfileMappingAttributesMap["expression_mapping_type"] = applicationProfileMappingAttributesChildRaw["ExpressionMappingType"]
			applicationProfileMappingAttributesMap["source_value_expression"] = applicationProfileMappingAttributesChildRaw["SourceValueExpression"]
			applicationProfileMappingAttributesMap["target_field"] = applicationProfileMappingAttributesChildRaw["TargetField"]
			applicationProfileMappingAttributesMap["target_field_description"] = applicationProfileMappingAttributesChildRaw["TargetFieldDescription"]

			applicationProfileMappingAttributesMaps = append(applicationProfileMappingAttributesMaps, applicationProfileMappingAttributesMap)
		}
	}
	if err := d.Set("application_profile_mapping_attributes", applicationProfileMappingAttributesMaps); err != nil {
		return err
	}

	objectRaw, err = eiamServiceV2.DescribeApplicationListUsersForApplication(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	usersMaps := make([]map[string]interface{}, 0)
	if objectRaw != nil {
		for _, userIdRaw := range convertToInterfaceArray(objectRaw) {
			usersMap := make(map[string]interface{})

			usersMap["user_id"] = objectRaw["UserId"]

			usersMaps = append(usersMaps, usersMap)
		}
	}
	if err := d.Set("users", usersMaps); err != nil {
		return err
	}

	objectRaw, err = eiamServiceV2.DescribeApplicationGetApplicationProvisioningFormConfig(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	fullPushScopesRaw := make([]interface{}, 0)
	if objectRaw["FullPushScopes"] != nil {
		fullPushScopesRaw = convertToInterfaceArray(objectRaw["FullPushScopes"])
	}

	d.Set("full_push_scopes", fullPushScopesRaw)
	provisioningActionsRaw := make([]interface{}, 0)
	if objectRaw["ProvisioningActions"] != nil {
		provisioningActionsRaw = convertToInterfaceArray(objectRaw["ProvisioningActions"])
	}

	d.Set("provisioning_actions", provisioningActionsRaw)

	objectRaw, err = eiamServiceV2.DescribeApplicationGetApplicationProvisioningScope(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	groupIdsRaw := make([]interface{}, 0)
	if objectRaw["GroupIds"] != nil {
		groupIdsRaw = convertToInterfaceArray(objectRaw["GroupIds"])
	}

	d.Set("group_ids", groupIdsRaw)
	organizationalUnitIdsRaw := make([]interface{}, 0)
	if objectRaw["OrganizationalUnitIds"] != nil {
		organizationalUnitIdsRaw = convertToInterfaceArray(objectRaw["OrganizationalUnitIds"])
	}

	d.Set("organizational_unit_ids", organizationalUnitIdsRaw)

	objectRaw, err = eiamServiceV2.DescribeApplicationGetApplicationGrantScope(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("grant_scopes", objectRaw)

	objectRaw, err = eiamServiceV2.DescribeApplicationListGroupsForApplication(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	groupsMaps := make([]map[string]interface{}, 0)
	if objectRaw != nil {
		for _, groupIdRaw := range convertToInterfaceArray(objectRaw) {
			groupsMap := make(map[string]interface{})

			groupsMap["group_id"] = objectRaw["GroupId"]

			groupsMaps = append(groupsMaps, groupsMap)
		}
	}
	if err := d.Set("groups", groupsMaps); err != nil {
		return err
	}

	objectRaw, err = eiamServiceV2.DescribeApplicationListOrganizationalUnitsForApplication(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	organizationalUnitsMaps := make([]map[string]interface{}, 0)
	if objectRaw != nil {
		for _, organizationalUnitIdRaw := range convertToInterfaceArray(objectRaw) {
			organizationalUnitsMap := make(map[string]interface{})

			organizationalUnitsMap["organizational_unit_id"] = objectRaw["OrganizationalUnitId"]

			organizationalUnitsMaps = append(organizationalUnitsMaps, organizationalUnitsMap)
		}
	}
	if err := d.Set("organizational_units", organizationalUnitsMaps); err != nil {
		return err
	}

	objectRaw, err = eiamServiceV2.DescribeApplicationGetApplicationProvisioningUserPrimaryOrganizationalUnit(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("user_primary_organizational_unit_id", objectRaw["UserPrimaryOrganizationalUnitId"])

	objectRaw, err = eiamServiceV2.DescribeApplicationGetApplicationClientAccessPolicy(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	networkZonesMaps := make([]map[string]interface{}, 0)
	networkZonesMap := make(map[string]interface{})

	networkZonesMap["include_network_zones"] = objectRaw
	networkZonesMaps = append(networkZonesMaps, networkZonesMap)
	if err := d.Set("network_zones", networkZonesMaps); err != nil {
		return err
	}

	return nil
}

func resourceAliCloudEiamApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	eiamServiceV2 := EiamServiceV2{client}
	objectRaw, _ := eiamServiceV2.DescribeEiamApplication(d.Id())

	if d.HasChange("status") {
		var err error
		target := d.Get("status").(string)
		if target == "enabled" {
			parts := strings.Split(d.Id(), ":")
			action := "EnableApplication"
			request = make(map[string]interface{})
			query = make(map[string]interface{})
			request["InstanceId"] = parts[0]
			request["ApplicationId"] = parts[1]
			request["RegionId"] = client.RegionId
			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
				if err != nil {
					if NeedRetry(err) {
						wait()
						return resource.RetryableError(err)
					}
					return resource.NonRetryableError(err)
				}
				return nil
			})
			addDebug(action, response, request)
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}

		}
		if target == "disabled" {
			parts := strings.Split(d.Id(), ":")
			action := "DisableApplication"
			request = make(map[string]interface{})
			query = make(map[string]interface{})
			request["InstanceId"] = parts[0]
			request["ApplicationId"] = parts[1]
			request["RegionId"] = client.RegionId
			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
				if err != nil {
					if NeedRetry(err) {
						wait()
						return resource.RetryableError(err)
					}
					return resource.NonRetryableError(err)
				}
				return nil
			})
			addDebug(action, response, request)
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}

		}
	}

	if d.HasChange("sso_status") {
		var err error
		target := d.Get("sso_status").(string)

		currentStatus, err := jsonpath.Get("SsoStatus", objectRaw)
		if err != nil {
			return WrapErrorf(err, FailedGetAttributeMsg, d.Id(), "SsoStatus", objectRaw)
		}
		if fmt.Sprint(currentStatus) != target {
			if target == "enabled" {
				parts := strings.Split(d.Id(), ":")
				action := "EnableApplicationSso"
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				request["InstanceId"] = parts[0]
				request["ApplicationId"] = parts[1]
				request["RegionId"] = client.RegionId
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
					if err != nil {
						if NeedRetry(err) {
							wait()
							return resource.RetryableError(err)
						}
						return resource.NonRetryableError(err)
					}
					return nil
				})
				addDebug(action, response, request)
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}

			}
			if target == "disabled" {
				parts := strings.Split(d.Id(), ":")
				action := "DisableApplicationSso"
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				request["InstanceId"] = parts[0]
				request["ApplicationId"] = parts[1]
				request["RegionId"] = client.RegionId
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
					if err != nil {
						if NeedRetry(err) {
							wait()
							return resource.RetryableError(err)
						}
						return resource.NonRetryableError(err)
					}
					return nil
				})
				addDebug(action, response, request)
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}

			}
		}
	}

	if d.HasChange("api_invoke_status") {
		var err error
		target := d.Get("api_invoke_status").(string)
		if target == "enabled" {
			parts := strings.Split(d.Id(), ":")
			action := "EnableApplicationApiInvoke"
			request = make(map[string]interface{})
			query = make(map[string]interface{})
			request["InstanceId"] = parts[0]
			request["ApplicationId"] = parts[1]
			request["RegionId"] = client.RegionId
			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
				if err != nil {
					if NeedRetry(err) {
						wait()
						return resource.RetryableError(err)
					}
					return resource.NonRetryableError(err)
				}
				return nil
			})
			addDebug(action, response, request)
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}

		}
		if target == "disabled" {
			parts := strings.Split(d.Id(), ":")
			action := "DisableApplicationApiInvoke"
			request = make(map[string]interface{})
			query = make(map[string]interface{})
			request["InstanceId"] = parts[0]
			request["ApplicationId"] = parts[1]
			request["RegionId"] = client.RegionId
			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
				if err != nil {
					if NeedRetry(err) {
						wait()
						return resource.RetryableError(err)
					}
					return resource.NonRetryableError(err)
				}
				return nil
			})
			addDebug(action, response, request)
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}

		}
	}

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "UpdateApplicationDescription"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "UpdateApplicationInfo"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("application_name") {
		update = true
	}
	request["ApplicationName"] = d.Get("application_name")
	if !d.IsNewResource() && d.HasChange("logo_url") {
		update = true
		request["LogoUrl"] = d.Get("logo_url")
	}

	if d.HasChange("grant_scopes") {
		update = true
		if v, ok := d.GetOk("grant_scopes"); ok || d.HasChange("grant_scopes") {
			applicationVisibilityMapsArray := convertToInterfaceArray(v)

			request["ApplicationVisibility"] = applicationVisibilityMapsArray
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "SetApplicationSsoConfig"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("oidc_sso_config") {
		update = true
		oidcSsoConfig := make(map[string]interface{})

		if v := d.Get("oidc_sso_config"); v != nil {
			localData, err := jsonpath.Get("$[0].custom_claims", v)
			if err != nil {
				localData = make([]interface{}, 0)
			}
			localMaps := make([]interface{}, 0)
			for _, dataLoop := range convertToInterfaceArray(localData) {
				dataLoopTmp := make(map[string]interface{})
				if dataLoop != nil {
					dataLoopTmp = dataLoop.(map[string]interface{})
				}
				dataLoopMap := make(map[string]interface{})
				dataLoopMap["ClaimValueExpression"] = dataLoopTmp["claim_value_expression"]
				dataLoopMap["ClaimName"] = dataLoopTmp["claim_name"]
				localMaps = append(localMaps, dataLoopMap)
			}
			oidcSsoConfig["CustomClaims"] = localMaps

			idTokenEffectiveTime1, _ := jsonpath.Get("$[0].id_token_effective_time", v)
			if idTokenEffectiveTime1 != nil && idTokenEffectiveTime1 != "" {
				oidcSsoConfig["IdTokenEffectiveTime"] = idTokenEffectiveTime1
			}
			passwordAuthenticationSourceId1, _ := jsonpath.Get("$[0].password_authentication_source_id", v)
			if passwordAuthenticationSourceId1 != nil && passwordAuthenticationSourceId1 != "" {
				oidcSsoConfig["PasswordAuthenticationSourceId"] = passwordAuthenticationSourceId1
			}
			grantTypes1, _ := jsonpath.Get("$[0].response_types", v)
			if grantTypes1 != nil && grantTypes1 != "" {
				oidcSsoConfig["GrantTypes"] = grantTypes1
			}
			pkceChallengeMethods1, _ := jsonpath.Get("$[0].pkce_challenge_methods", v)
			if pkceChallengeMethods1 != nil && pkceChallengeMethods1 != "" {
				oidcSsoConfig["PkceChallengeMethods"] = pkceChallengeMethods1
			}
			passwordTotpMfaRequired1, _ := jsonpath.Get("$[0].password_totp_mfa_required", v)
			if passwordTotpMfaRequired1 != nil && passwordTotpMfaRequired1 != "" {
				oidcSsoConfig["PasswordTotpMfaRequired"] = passwordTotpMfaRequired1
			}
			codeEffectiveTime1, _ := jsonpath.Get("$[0].code_effective_time", v)
			if codeEffectiveTime1 != nil && codeEffectiveTime1 != "" {
				oidcSsoConfig["CodeEffectiveTime"] = codeEffectiveTime1
			}
			responseTypes1, _ := jsonpath.Get("$[0].grant_types", v)
			if responseTypes1 != nil && responseTypes1 != "" {
				oidcSsoConfig["ResponseTypes"] = responseTypes1
			}
			postLogoutRedirectUris1, _ := jsonpath.Get("$[0].post_logout_redirect_uris", v)
			if postLogoutRedirectUris1 != nil && postLogoutRedirectUris1 != "" {
				oidcSsoConfig["PostLogoutRedirectUris"] = postLogoutRedirectUris1
			}
			accessTokenEffectiveTime1, _ := jsonpath.Get("$[0].access_token_effective_time", v)
			if accessTokenEffectiveTime1 != nil && accessTokenEffectiveTime1 != "" {
				oidcSsoConfig["AccessTokenEffectiveTime"] = accessTokenEffectiveTime1
			}
			grantScopes1, _ := jsonpath.Get("$[0].grant_scopes", v)
			if grantScopes1 != nil && grantScopes1 != "" {
				oidcSsoConfig["GrantScopes"] = grantScopes1
			}
			redirectUris1, _ := jsonpath.Get("$[0].redirect_uris", v)
			if redirectUris1 != nil && redirectUris1 != "" {
				oidcSsoConfig["RedirectUris"] = redirectUris1
			}
			refreshTokenEffective1, _ := jsonpath.Get("$[0].refresh_token_effective", v)
			if refreshTokenEffective1 != nil && refreshTokenEffective1 != "" {
				oidcSsoConfig["RefreshTokenEffective"] = refreshTokenEffective1
			}
			pkceRequired1, _ := jsonpath.Get("$[0].pkce_required", v)
			if pkceRequired1 != nil && pkceRequired1 != "" {
				oidcSsoConfig["PkceRequired"] = pkceRequired1
			}
			subjectIdExpression1, _ := jsonpath.Get("$[0].subject_id_expression", v)
			if subjectIdExpression1 != nil && subjectIdExpression1 != "" {
				oidcSsoConfig["SubjectIdExpression"] = subjectIdExpression1
			}

			request["OidcSsoConfig"] = oidcSsoConfig
		}
	}

	if d.HasChange("oidc_sso_config.0.id_token_effective_time") {
		update = true
		oidcSsoConfigIdTokenEffectiveTimeJsonPath, err := jsonpath.Get("$[0].id_token_effective_time", d.Get("oidc_sso_config"))
		if err == nil {
			request["OidcSsoConfig.IdTokenEffectiveTime"] = oidcSsoConfigIdTokenEffectiveTimeJsonPath
		}
	}

	if d.HasChange("oidc_sso_config.0.password_authentication_source_id") {
		update = true
		oidcSsoConfigPasswordAuthenticationSourceIdJsonPath, err := jsonpath.Get("$[0].password_authentication_source_id", d.Get("oidc_sso_config"))
		if err == nil {
			request["OidcSsoConfig.PasswordAuthenticationSourceId"] = oidcSsoConfigPasswordAuthenticationSourceIdJsonPath
		}
	}

	if d.HasChange("saml_sso_config.0.sp_sso_acs_url") {
		update = true
		samlSsoConfigSpSsoAcsUrlJsonPath, err := jsonpath.Get("$[0].sp_sso_acs_url", d.Get("saml_sso_config"))
		if err == nil {
			request["SamlSsoConfig.SpSsoAcsUrl"] = samlSsoConfigSpSsoAcsUrlJsonPath
		}
	}

	if d.HasChange("saml_sso_config") {
		update = true
		samlSsoConfig := make(map[string]interface{})

		if v := d.Get("saml_sso_config"); v != nil {
			spSsoAcsUrl1, _ := jsonpath.Get("$[0].sp_sso_acs_url", v)
			if spSsoAcsUrl1 != nil && spSsoAcsUrl1 != "" {
				samlSsoConfig["SpSsoAcsUrl"] = spSsoAcsUrl1
			}
			localData1, err := jsonpath.Get("$[0].attribute_statements", v)
			if err != nil {
				localData1 = make([]interface{}, 0)
			}
			localMaps1 := make([]interface{}, 0)
			for _, dataLoop1 := range convertToInterfaceArray(localData1) {
				dataLoop1Tmp := make(map[string]interface{})
				if dataLoop1 != nil {
					dataLoop1Tmp = dataLoop1.(map[string]interface{})
				}
				dataLoop1Map := make(map[string]interface{})
				dataLoop1Map["AttributeValueExpression"] = dataLoop1Tmp["attribute_value_expression"]
				dataLoop1Map["AttributeName"] = dataLoop1Tmp["attribute_name"]
				localMaps1 = append(localMaps1, dataLoop1Map)
			}
			samlSsoConfig["AttributeStatements"] = localMaps1

			signatureAlgorithm1, _ := jsonpath.Get("$[0].signature_algorithm", v)
			if signatureAlgorithm1 != nil && signatureAlgorithm1 != "" {
				samlSsoConfig["SignatureAlgorithm"] = signatureAlgorithm1
			}
			defaultRelayState1, _ := jsonpath.Get("$[0].default_relay_state", v)
			if defaultRelayState1 != nil && defaultRelayState1 != "" {
				samlSsoConfig["DefaultRelayState"] = defaultRelayState1
			}
			nameIdFormat1, _ := jsonpath.Get("$[0].name_id_format", v)
			if nameIdFormat1 != nil && nameIdFormat1 != "" {
				samlSsoConfig["NameIdFormat"] = nameIdFormat1
			}
			spEntityId1, _ := jsonpath.Get("$[0].sp_entity_id", v)
			if spEntityId1 != nil && spEntityId1 != "" {
				samlSsoConfig["SpEntityId"] = spEntityId1
			}
			nameIdValueExpression1, _ := jsonpath.Get("$[0].name_id_value_expression", v)
			if nameIdValueExpression1 != nil && nameIdValueExpression1 != "" {
				samlSsoConfig["NameIdValueExpression"] = nameIdValueExpression1
			}
			assertionSigned1, _ := jsonpath.Get("$[0].assertion_signed", v)
			if assertionSigned1 != nil && assertionSigned1 != "" {
				samlSsoConfig["AssertionSigned"] = assertionSigned1
			}
			responseSigned1, _ := jsonpath.Get("$[0].response_signed", v)
			if responseSigned1 != nil && responseSigned1 != "" {
				samlSsoConfig["ResponseSigned"] = responseSigned1
			}

			request["SamlSsoConfig"] = samlSsoConfig
		}
	}

	if d.HasChange("saml_sso_config.0.signature_algorithm") {
		update = true
		samlSsoConfigSignatureAlgorithmJsonPath, err := jsonpath.Get("$[0].signature_algorithm", d.Get("saml_sso_config"))
		if err == nil {
			request["SamlSsoConfig.SignatureAlgorithm"] = samlSsoConfigSignatureAlgorithmJsonPath
		}
	}

	if d.HasChange("oidc_sso_config.0.password_totp_mfa_required") {
		update = true
		oidcSsoConfigPasswordTotpMfaRequiredJsonPath, err := jsonpath.Get("$[0].password_totp_mfa_required", d.Get("oidc_sso_config"))
		if err == nil {
			request["OidcSsoConfig.PasswordTotpMfaRequired"] = oidcSsoConfigPasswordTotpMfaRequiredJsonPath
		}
	}

	if d.HasChange("oidc_sso_config.0.code_effective_time") {
		update = true
		oidcSsoConfigCodeEffectiveTimeJsonPath, err := jsonpath.Get("$[0].code_effective_time", d.Get("oidc_sso_config"))
		if err == nil {
			request["OidcSsoConfig.CodeEffectiveTime"] = oidcSsoConfigCodeEffectiveTimeJsonPath
		}
	}

	if d.HasChange("saml_sso_config.0.default_relay_state") {
		update = true
		samlSsoConfigDefaultRelayStateJsonPath, err := jsonpath.Get("$[0].default_relay_state", d.Get("saml_sso_config"))
		if err == nil {
			request["SamlSsoConfig.DefaultRelayState"] = samlSsoConfigDefaultRelayStateJsonPath
		}
	}

	if d.HasChange("saml_sso_config.0.name_id_format") {
		update = true
		samlSsoConfigNameIdFormatJsonPath, err := jsonpath.Get("$[0].name_id_format", d.Get("saml_sso_config"))
		if err == nil {
			request["SamlSsoConfig.NameIdFormat"] = samlSsoConfigNameIdFormatJsonPath
		}
	}

	if d.HasChange("oidc_sso_config.0.access_token_effective_time") {
		update = true
		oidcSsoConfigAccessTokenEffectiveTimeJsonPath, err := jsonpath.Get("$[0].access_token_effective_time", d.Get("oidc_sso_config"))
		if err == nil {
			request["OidcSsoConfig.AccessTokenEffectiveTime"] = oidcSsoConfigAccessTokenEffectiveTimeJsonPath
		}
	}

	if d.HasChange("saml_sso_config.0.sp_entity_id") {
		update = true
		samlSsoConfigSpEntityIdJsonPath, err := jsonpath.Get("$[0].sp_entity_id", d.Get("saml_sso_config"))
		if err == nil {
			request["SamlSsoConfig.SpEntityId"] = samlSsoConfigSpEntityIdJsonPath
		}
	}

	if d.HasChange("oidc_sso_config.0.refresh_token_effective") {
		update = true
		oidcSsoConfigRefreshTokenEffectiveJsonPath, err := jsonpath.Get("$[0].refresh_token_effective", d.Get("oidc_sso_config"))
		if err == nil {
			request["OidcSsoConfig.RefreshTokenEffective"] = oidcSsoConfigRefreshTokenEffectiveJsonPath
		}
	}

	if d.HasChange("oidc_sso_config.0.pkce_required") {
		update = true
		oidcSsoConfigPkceRequiredJsonPath, err := jsonpath.Get("$[0].pkce_required", d.Get("oidc_sso_config"))
		if err == nil {
			request["OidcSsoConfig.PkceRequired"] = oidcSsoConfigPkceRequiredJsonPath
		}
	}

	if d.HasChange("oidc_sso_config.0.subject_id_expression") {
		update = true
		oidcSsoConfigSubjectIdExpressionJsonPath, err := jsonpath.Get("$[0].subject_id_expression", d.Get("oidc_sso_config"))
		if err == nil {
			request["OidcSsoConfig.SubjectIdExpression"] = oidcSsoConfigSubjectIdExpressionJsonPath
		}
	}

	if d.HasChange("saml_sso_config.0.name_id_value_expression") {
		update = true
		samlSsoConfigNameIdValueExpressionJsonPath, err := jsonpath.Get("$[0].name_id_value_expression", d.Get("saml_sso_config"))
		if err == nil {
			request["SamlSsoConfig.NameIdValueExpression"] = samlSsoConfigNameIdValueExpressionJsonPath
		}
	}

	if d.HasChange("saml_sso_config.0.assertion_signed") {
		update = true
		samlSsoConfigAssertionSignedJsonPath, err := jsonpath.Get("$[0].assertion_signed", d.Get("saml_sso_config"))
		if err == nil {
			request["SamlSsoConfig.AssertionSigned"] = samlSsoConfigAssertionSignedJsonPath
		}
	}

	if d.HasChange("saml_sso_config.0.response_signed") {
		update = true
		samlSsoConfigResponseSignedJsonPath, err := jsonpath.Get("$[0].response_signed", d.Get("saml_sso_config"))
		if err == nil {
			request["SamlSsoConfig.ResponseSigned"] = samlSsoConfigResponseSignedJsonPath
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "SetApplicationGrantScope"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if d.HasChange("grant_scopes") {
		update = true
		if v, ok := d.GetOk("grant_scopes"); ok || d.HasChange("grant_scopes") {
			grantScopesMapsArray := convertToInterfaceArray(v)

			request["GrantScopes"] = grantScopesMapsArray
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "SetApplicationProvisioningConfig"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ApplicationId"] = parts[1]
	request["InstanceId"] = parts[0]
	request["RegionId"] = client.RegionId
	callbackProvisioningConfig := make(map[string]interface{})

	if d.HasChange("scim_provisioning_config") {
		update = true
		listenEventScopes1, _ := jsonpath.Get("$[0].full_push_scopes", d.Get("scim_provisioning_config"))
		if listenEventScopes1 != nil && listenEventScopes1 != "" {
			callbackProvisioningConfig["ListenEventScopes"] = listenEventScopes1
		}
	}

	if d.HasChange("callback_provisioning_config") {
		update = true
		encryptKey1, _ := jsonpath.Get("$[0].encrypt_key", d.Get("callback_provisioning_config"))
		if encryptKey1 != nil && encryptKey1 != "" {
			callbackProvisioningConfig["EncryptKey"] = encryptKey1
		}
	}

	if d.HasChange("callback_provisioning_config") {
		update = true
		encryptRequired1, _ := jsonpath.Get("$[0].encrypt_required", d.Get("callback_provisioning_config"))
		if encryptRequired1 != nil && encryptRequired1 != "" {
			callbackProvisioningConfig["EncryptRequired"] = encryptRequired1
		}
	}

	if d.HasChange("callback_provisioning_config") {
		update = true
		callbackUrl1, _ := jsonpath.Get("$[0].callback_url", d.Get("callback_provisioning_config"))
		if callbackUrl1 != nil && callbackUrl1 != "" {
			callbackProvisioningConfig["CallbackUrl"] = callbackUrl1
		}
	}

	request["CallbackProvisioningConfig"] = callbackProvisioningConfig

	scimProvisioningConfig := make(map[string]interface{})

	if d.HasChange("callback_provisioning_config") {
		update = true
		fullPushScopes1, _ := jsonpath.Get("$[0].listen_event_scopes", d.Get("callback_provisioning_config"))
		if fullPushScopes1 != nil && fullPushScopes1 != "" {
			scimProvisioningConfig["FullPushScopes"] = fullPushScopes1
		}
	}

	if d.HasChange("scim_provisioning_config") {
		update = true
		provisioningActions1, _ := jsonpath.Get("$[0].provisioning_actions", d.Get("scim_provisioning_config"))
		if provisioningActions1 != nil && provisioningActions1 != "" {
			scimProvisioningConfig["ProvisioningActions"] = provisioningActions1
		}
	}

	if d.HasChange("scim_provisioning_config") {
		update = true
		if v := d.Get("scim_provisioning_config"); v != nil {
			authnConfiguration := make(map[string]interface{})
			authnParam := make(map[string]interface{})
			authnMethod1, _ := jsonpath.Get("$[0].authn_configuration[0].authn_param[0].authn_method", d.Get("scim_provisioning_config"))
			if authnMethod1 != nil && authnMethod1 != "" {
				authnParam["AuthnMethod"] = authnMethod1
			}
			tokenEndpoint1, _ := jsonpath.Get("$[0].authn_configuration[0].authn_param[0].token_endpoint", d.Get("scim_provisioning_config"))
			if tokenEndpoint1 != nil && tokenEndpoint1 != "" {
				authnParam["TokenEndpoint"] = tokenEndpoint1
			}
			accessToken1, _ := jsonpath.Get("$[0].authn_configuration[0].authn_param[0].access_token", d.Get("scim_provisioning_config"))
			if accessToken1 != nil && accessToken1 != "" {
				authnParam["AccessToken"] = accessToken1
			}
			clientId1, _ := jsonpath.Get("$[0].authn_configuration[0].authn_param[0].client_id", d.Get("scim_provisioning_config"))
			if clientId1 != nil && clientId1 != "" {
				authnParam["ClientId"] = clientId1
			}
			clientSecret1, _ := jsonpath.Get("$[0].authn_configuration[0].authn_param[0].client_secret", d.Get("scim_provisioning_config"))
			if clientSecret1 != nil && clientSecret1 != "" {
				authnParam["ClientSecret"] = clientSecret1
			}

			if len(authnParam) > 0 {
				authnConfiguration["AuthnParam"] = authnParam
			}
			grantType1, _ := jsonpath.Get("$[0].authn_configuration[0].grant_type", d.Get("scim_provisioning_config"))
			if grantType1 != nil && grantType1 != "" {
				authnConfiguration["GrantType"] = grantType1
			}
			authnMode1, _ := jsonpath.Get("$[0].authn_configuration[0].authn_mode", d.Get("scim_provisioning_config"))
			if authnMode1 != nil && authnMode1 != "" {
				authnConfiguration["AuthnMode"] = authnMode1
			}

			if len(authnConfiguration) > 0 {
				scimProvisioningConfig["AuthnConfiguration"] = authnConfiguration
			}
		}
	}

	if d.HasChange("scim_provisioning_config") {
		update = true
		scimBaseUrl1, _ := jsonpath.Get("$[0].scim_base_url", d.Get("scim_provisioning_config"))
		if scimBaseUrl1 != nil && scimBaseUrl1 != "" {
			scimProvisioningConfig["ScimBaseUrl"] = scimBaseUrl1
		}
	}

	request["ScimProvisioningConfig"] = scimProvisioningConfig

	if d.HasChange("callback_provisioning_config.0.encrypt_key") {
		update = true
		callbackProvisioningConfigEncryptKeyJsonPath, err := jsonpath.Get("$[0].encrypt_key", d.Get("callback_provisioning_config"))
		if err == nil {
			request["CallbackProvisioningConfig.EncryptKey"] = callbackProvisioningConfigEncryptKeyJsonPath
		}
	}

	if d.HasChange("callback_provisioning_config.0.encrypt_required") {
		update = true
		callbackProvisioningConfigEncryptRequiredJsonPath, err := jsonpath.Get("$[0].encrypt_required", d.Get("callback_provisioning_config"))
		if err == nil {
			request["CallbackProvisioningConfig.EncryptRequired"] = callbackProvisioningConfigEncryptRequiredJsonPath
		}
	}

	if d.HasChange("provision_password") {
		update = true
		request["ProvisionPassword"] = d.Get("provision_password")
	}

	if d.HasChange("provision_protocol_type") {
		update = true
	}
	request["ProvisionProtocolType"] = d.Get("provision_protocol_type")
	if d.HasChange("callback_provisioning_config.0.callback_url") {
		update = true
		callbackProvisioningConfigCallbackUrlJsonPath, err := jsonpath.Get("$[0].callback_url", d.Get("callback_provisioning_config"))
		if err == nil {
			request["CallbackProvisioningConfig.CallbackUrl"] = callbackProvisioningConfigCallbackUrlJsonPath
		}
	}

	if d.HasChange("scim_provisioning_config.0.scim_base_url") {
		update = true
		scimProvisioningConfigScimBaseUrlJsonPath, err := jsonpath.Get("$[0].scim_base_url", d.Get("scim_provisioning_config"))
		if err == nil {
			request["ScimProvisioningConfig.ScimBaseUrl"] = scimProvisioningConfigScimBaseUrlJsonPath
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "ChangeApplicationProvisioningConfigOperateMode"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "DisableApplicationProvisioning"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "EnableApplicationProvisioning"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "SetApplicationProvisioningUserPrimaryOrganizationalUnit"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if d.HasChange("user_primary_organizational_unit_id") {
		update = true
	}
	request["UserPrimaryOrganizationalUnitId"] = d.Get("user_primary_organizational_unit_id")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "SetApplicationUserProfileMapping"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if d.HasChange("application_profile_mapping_attributes") {
		update = true
		if v, ok := d.GetOk("application_profile_mapping_attributes"); ok || d.HasChange("application_profile_mapping_attributes") {
			applicationProfileMappingAttributesMapsArray := make([]interface{}, 0)
			for _, dataLoop := range convertToInterfaceArray(v) {
				dataLoopTmp := dataLoop.(map[string]interface{})
				dataLoopMap := make(map[string]interface{})
				dataLoopMap["ExpressionMappingType"] = dataLoopTmp["expression_mapping_type"]
				dataLoopMap["TargetFieldDescription"] = dataLoopTmp["target_field_description"]
				dataLoopMap["SourceValueExpression"] = dataLoopTmp["source_value_expression"]
				dataLoopMap["TargetField"] = dataLoopTmp["target_field"]
				applicationProfileMappingAttributesMapsArray = append(applicationProfileMappingAttributesMapsArray, dataLoopMap)
			}
			request["ApplicationProfileMappingAttributes"] = applicationProfileMappingAttributesMapsArray
		}
	}

	if d.HasChange("user_mapping_identity_name") {
		update = true
		request["UserMappingIdentityName"] = d.Get("user_mapping_identity_name")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "ChangeApplicationSourceType"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "SetApplicationProvisioningFormConfig"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	applicationTemplateParamsDataList := make(map[string]interface{})

	templateParamName1, _ := jsonpath.Get("$.template_param_name", d.Get("application_template_params"))
	if templateParamName1 != nil && templateParamName1 != "" {
		applicationTemplateParamsDataList["TemplateParamName"] = templateParamName1
	}

	templateParamValue1, _ := jsonpath.Get("$.template_param_value", d.Get("application_template_params"))
	if templateParamValue1 != nil && templateParamValue1 != "" {
		applicationTemplateParamsDataList["TemplateParamValue"] = templateParamValue1
	}

	ApplicationTemplateParamsMap := make([]interface{}, 0)
	ApplicationTemplateParamsMap = append(ApplicationTemplateParamsMap, applicationTemplateParamsDataList)
	request["ApplicationTemplateParams"] = ApplicationTemplateParamsMap

	if d.HasChange("provisioning_actions") {
		update = true
		if v, ok := d.GetOk("provisioning_actions"); ok || d.HasChange("provisioning_actions") {
			fullPushScopesMapsArray := convertToInterfaceArray(v)

			request["FullPushScopes"] = fullPushScopesMapsArray
		}
	}

	if d.HasChange("full_push_scopes") {
		update = true
		if v, ok := d.GetOk("full_push_scopes"); ok || d.HasChange("full_push_scopes") {
			provisioningActionsMapsArray := convertToInterfaceArray(v)

			request["ProvisioningActions"] = provisioningActionsMapsArray
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "SetApplicationProvisioningScope"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if d.HasChange("group_ids") {
		update = true
		if v, ok := d.GetOk("group_ids"); ok || d.HasChange("group_ids") {
			organizationalUnitIdsMapsArray := convertToInterfaceArray(v)

			request["OrganizationalUnitIds"] = organizationalUnitIdsMapsArray
		}
	}

	if d.HasChange("organizational_unit_ids") {
		update = true
		if v, ok := d.GetOk("organizational_unit_ids"); ok || d.HasChange("organizational_unit_ids") {
			groupIdsMapsArray := convertToInterfaceArray(v)

			request["GroupIds"] = groupIdsMapsArray
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "EnableApplicationM2MClient"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "DisableApplicationM2MClient"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "SetApplicationClientAccessPolicy"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if d.HasChange("network_zones") {
		update = true
		networkZones := make(map[string]interface{})

		if v := d.Get("network_zones"); v != nil {
			includeNetworkZones1, _ := jsonpath.Get("$[0].include_network_zones", v)
			if includeNetworkZones1 != nil && includeNetworkZones1 != "" {
				networkZones["IncludeNetworkZones"] = includeNetworkZones1
			}

			request["NetworkZones"] = networkZones
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "UpdateApplicationAuthorizationType"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if d.HasChange("authorization_type") {
		update = true
	}
	request["AuthorizationType"] = d.Get("authorization_type")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "UpdateApplicationSsoFormParams"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	applicationTemplateParamsDataList := make(map[string]interface{})

	templateParamName1, _ := jsonpath.Get("$.template_param_name", d.Get("application_template_params"))
	if templateParamName1 != nil && templateParamName1 != "" {
		applicationTemplateParamsDataList["TemplateParamName"] = templateParamName1
	}

	templateParamValue1, _ := jsonpath.Get("$.template_param_value", d.Get("application_template_params"))
	if templateParamValue1 != nil && templateParamValue1 != "" {
		applicationTemplateParamsDataList["TemplateParamValue"] = templateParamValue1
	}

	ApplicationTemplateParamsMap := make([]interface{}, 0)
	ApplicationTemplateParamsMap = append(ApplicationTemplateParamsMap, applicationTemplateParamsDataList)
	request["ApplicationTemplateParams"] = ApplicationTemplateParamsMap

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "EnableApplicationResourceServer"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "DisableApplicationResourceServer"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "SetApplicationResourceServerIdentifier"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("resource_server_identifier") {
		update = true
	}
	request["ResourceServerIdentifier"] = d.Get("resource_server_identifier")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "AuthorizeResourceServerToClient"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ClientApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if d.HasChange("resource_server_application_id") {
		update = true
	}
	request["ResourceServerApplicationId"] = d.Get("resource_server_application_id")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "RevokeResourceServerFromClient"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ClientApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if d.HasChange("resource_server_application_id") {
		update = true
	}
	request["ResourceServerApplicationId"] = d.Get("resource_server_application_id")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "AuthorizeResourceServerScopesToClient"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ClientApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if d.HasChange("resource_server_application_id") {
		update = true
	}
	request["ResourceServerApplicationId"] = d.Get("resource_server_application_id")
	if d.HasChange("resource_server_scope_ids") {
		update = true
	}
	if v, ok := d.GetOk("resource_server_scope_ids"); ok || d.HasChange("resource_server_scope_ids") {
		resourceServerScopeIdsMapsArray := convertToInterfaceArray(v)

		request["ResourceServerScopeIds"] = resourceServerScopeIdsMapsArray
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "RevokeResourceServerScopesFromClient"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ClientApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if d.HasChange("resource_server_application_id") {
		update = true
	}
	request["ResourceServerApplicationId"] = d.Get("resource_server_application_id")
	if d.HasChange("resource_server_scope_ids") {
		update = true
	}
	if v, ok := d.GetOk("resource_server_scope_ids"); ok || d.HasChange("resource_server_scope_ids") {
		resourceServerScopeIdsMapsArray := convertToInterfaceArray(v)

		request["ResourceServerScopeIds"] = resourceServerScopeIdsMapsArray
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "EnableResourceServerCustomSubject"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "DisableResourceServerCustomSubject"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "UpdateApplicationAdvancedConfig"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	if d.HasChange("application_advanced_config") {
		update = true
		scimServerAdvancedConfig := make(map[string]interface{})

		if v := d.Get("application_advanced_config"); v != nil {
			supportedUserCustomFieldIds1, _ := jsonpath.Get("$[0].scim_server_advanced_config[0].supported_user_custom_field_ids", v)
			if supportedUserCustomFieldIds1 != nil && supportedUserCustomFieldIds1 != "" {
				scimServerAdvancedConfig["SupportedUserCustomFieldIds"] = supportedUserCustomFieldIds1
			}
			userCustomFieldNamespace1, _ := jsonpath.Get("$[0].scim_server_advanced_config[0].user_custom_field_namespace", v)
			if userCustomFieldNamespace1 != nil && userCustomFieldNamespace1 != "" {
				scimServerAdvancedConfig["UserCustomFieldNamespace"] = userCustomFieldNamespace1
			}

			request["ScimServerAdvancedConfig"] = scimServerAdvancedConfig
		}
	}

	if d.HasChange("application_advanced_config.0.scim_server_advanced_config.0.user_custom_field_namespace") {
		update = true
		applicationAdvancedConfigScimServerAdvancedConfigUserCustomFieldNamespaceJsonPath, err := jsonpath.Get("$[0].scim_server_advanced_config[0].user_custom_field_namespace", d.Get("application_advanced_config"))
		if err == nil {
			request["ScimServerAdvancedConfig.UserCustomFieldNamespace"] = applicationAdvancedConfigScimServerAdvancedConfigUserCustomFieldNamespaceJsonPath
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "UpdateApplicationSsoType"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("sso_type") {
		update = true
	}
	request["SsoType"] = d.Get("sso_type")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}

	d.Partial(false)
	return resourceAliCloudEiamApplicationRead(d, meta)
}

func resourceAliCloudEiamApplicationDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteApplication"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	action = "RevokeApplicationFromUsers"
	request = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("users"); ok {
		localData, err := jsonpath.Get("$[*].user_id", v)
		if err != nil {
			return WrapError(err)
		}
		localDataArray := convertToInterfaceArray(localData)

		request["UserIds"] = localDataArray
	}

	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	action = "RevokeApplicationFromGroups"
	request = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("groups"); ok {
		localData, err := jsonpath.Get("$[*].group_id", v)
		if err != nil {
			return WrapError(err)
		}
		localDataArray := convertToInterfaceArray(localData)

		request["GroupIds"] = localDataArray
	}

	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	action = "RevokeApplicationFromOrganizationalUnits"
	request = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("organizational_units"); ok {
		localData, err := jsonpath.Get("$[*].organizational_unit_id", v)
		if err != nil {
			return WrapError(err)
		}
		localDataArray := convertToInterfaceArray(localData)

		request["OrganizationalUnitIds"] = localDataArray
	}

	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	action = "RevokeResourceServerScopesFromGroup"
	request = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("resource_server_scope_ids"); ok {
		resourceServerScopeIdsMapsArray := convertToInterfaceArray(v)

		request["ResourceServerScopeIds"] = resourceServerScopeIdsMapsArray
	}

	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	action = "RevokeResourceServerScopesFromOrganizationalUnit"
	request = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("resource_server_scope_ids"); ok {
		resourceServerScopeIdsMapsArray := convertToInterfaceArray(v)

		request["ResourceServerScopeIds"] = resourceServerScopeIdsMapsArray
	}

	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	action = "RevokeResourceServerScopesFromUser"
	request = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ApplicationId"] = parts[1]
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("resource_server_scope_ids"); ok {
		resourceServerScopeIdsMapsArray := convertToInterfaceArray(v)

		request["ResourceServerScopeIds"] = resourceServerScopeIdsMapsArray
	}

	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Eiam", "2021-12-01", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
