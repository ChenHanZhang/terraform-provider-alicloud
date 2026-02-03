// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudEventBridgeConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudEventBridgeConnectionCreate,
		Read:   resourceAliCloudEventBridgeConnectionRead,
		Update: resourceAliCloudEventBridgeConnectionUpdate,
		Delete: resourceAliCloudEventBridgeConnectionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"auth_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"oauth_parameters": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_parameters": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"client_secret": {
													Type:      schema.TypeString,
													Optional:  true,
													Sensitive: true,
												},
												"client_id": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"oauth_http_parameters": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"header_parameters": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"is_value_secret": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"value": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"key": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"query_string_parameters": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"is_value_secret": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"value": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"key": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"body_parameters": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"is_value_secret": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"value": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"key": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"authorization_endpoint": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"http_method": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"basic_auth_parameters": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"username": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"password": {
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
								},
							},
						},
						"api_key_auth_parameters": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"api_key_value": {
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
									"api_key_name": {
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
					},
				},
			},
			"connection_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_parameters": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"network_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"security_group_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"vswitche_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceAliCloudEventBridgeConnectionCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateConnection"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("connection_name"); ok {
		request["ConnectionName"] = v
	}

	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	networkParameters := make(map[string]interface{})

	if v := d.Get("network_parameters"); v != nil {
		networkType1, _ := jsonpath.Get("$[0].network_type", v)
		if networkType1 != nil && networkType1 != "" {
			networkParameters["NetworkType"] = networkType1
		}
		vpcId1, _ := jsonpath.Get("$[0].vpc_id", v)
		if vpcId1 != nil && vpcId1 != "" {
			networkParameters["VpcId"] = vpcId1
		}
		vswitcheId1, _ := jsonpath.Get("$[0].vswitche_id", v)
		if vswitcheId1 != nil && vswitcheId1 != "" {
			networkParameters["VswitcheId"] = vswitcheId1
		}
		securityGroupId1, _ := jsonpath.Get("$[0].security_group_id", v)
		if securityGroupId1 != nil && securityGroupId1 != "" {
			networkParameters["SecurityGroupId"] = securityGroupId1
		}

		networkParametersJson, err := json.Marshal(networkParameters)
		if err != nil {
			return WrapError(err)
		}
		request["NetworkParameters"] = string(networkParametersJson)
	}

	authParameters := make(map[string]interface{})

	if v := d.Get("auth_parameters"); !IsNil(v) {
		authorizationType1, _ := jsonpath.Get("$[0].authorization_type", v)
		if authorizationType1 != nil && authorizationType1 != "" {
			authParameters["AuthorizationType"] = authorizationType1
		}
		apiKeyAuthParameters := make(map[string]interface{})
		apiKeyName1, _ := jsonpath.Get("$[0].api_key_auth_parameters[0].api_key_name", d.Get("auth_parameters"))
		if apiKeyName1 != nil && apiKeyName1 != "" {
			apiKeyAuthParameters["ApiKeyName"] = apiKeyName1
		}
		apiKeyValue1, _ := jsonpath.Get("$[0].api_key_auth_parameters[0].api_key_value", d.Get("auth_parameters"))
		if apiKeyValue1 != nil && apiKeyValue1 != "" {
			apiKeyAuthParameters["ApiKeyValue"] = apiKeyValue1
		}

		if len(apiKeyAuthParameters) > 0 {
			authParameters["ApiKeyAuthParameters"] = apiKeyAuthParameters
		}
		basicAuthParameters := make(map[string]interface{})
		password1, _ := jsonpath.Get("$[0].basic_auth_parameters[0].password", d.Get("auth_parameters"))
		if password1 != nil && password1 != "" {
			basicAuthParameters["Password"] = password1
		}
		username1, _ := jsonpath.Get("$[0].basic_auth_parameters[0].username", d.Get("auth_parameters"))
		if username1 != nil && username1 != "" {
			basicAuthParameters["Username"] = username1
		}

		if len(basicAuthParameters) > 0 {
			authParameters["BasicAuthParameters"] = basicAuthParameters
		}
		oAuthParameters := make(map[string]interface{})
		clientParameters := make(map[string]interface{})
		clientId, _ := jsonpath.Get("$[0].oauth_parameters[0].client_parameters[0].client_id", d.Get("auth_parameters"))
		if clientId != nil && clientId != "" {
			clientParameters["ClientID"] = clientId
		}
		clientSecret1, _ := jsonpath.Get("$[0].oauth_parameters[0].client_parameters[0].client_secret", d.Get("auth_parameters"))
		if clientSecret1 != nil && clientSecret1 != "" {
			clientParameters["ClientSecret"] = clientSecret1
		}

		if len(clientParameters) > 0 {
			oAuthParameters["ClientParameters"] = clientParameters
		}
		httpMethod1, _ := jsonpath.Get("$[0].oauth_parameters[0].http_method", d.Get("auth_parameters"))
		if httpMethod1 != nil && httpMethod1 != "" {
			oAuthParameters["HttpMethod"] = httpMethod1
		}
		authorizationEndpoint1, _ := jsonpath.Get("$[0].oauth_parameters[0].authorization_endpoint", d.Get("auth_parameters"))
		if authorizationEndpoint1 != nil && authorizationEndpoint1 != "" {
			oAuthParameters["AuthorizationEndpoint"] = authorizationEndpoint1
		}
		oAuthHttpParameters := make(map[string]interface{})
		localData, err := jsonpath.Get("$[0].oauth_parameters[0].oauth_http_parameters[0].body_parameters", v)
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
			dataLoopMap["IsValueSecret"] = dataLoopTmp["is_value_secret"]
			dataLoopMap["Key"] = dataLoopTmp["key"]
			dataLoopMap["Value"] = dataLoopTmp["value"]
			localMaps = append(localMaps, dataLoopMap)
		}
		oAuthHttpParameters["BodyParameters"] = localMaps

		localData1, err := jsonpath.Get("$[0].oauth_parameters[0].oauth_http_parameters[0].header_parameters", v)
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
			dataLoop1Map["IsValueSecret"] = dataLoop1Tmp["is_value_secret"]
			dataLoop1Map["Key"] = dataLoop1Tmp["key"]
			dataLoop1Map["Value"] = dataLoop1Tmp["value"]
			localMaps1 = append(localMaps1, dataLoop1Map)
		}
		oAuthHttpParameters["HeaderParameters"] = localMaps1

		localData2, err := jsonpath.Get("$[0].oauth_parameters[0].oauth_http_parameters[0].query_string_parameters", v)
		if err != nil {
			localData2 = make([]interface{}, 0)
		}
		localMaps2 := make([]interface{}, 0)
		for _, dataLoop2 := range convertToInterfaceArray(localData2) {
			dataLoop2Tmp := make(map[string]interface{})
			if dataLoop2 != nil {
				dataLoop2Tmp = dataLoop2.(map[string]interface{})
			}
			dataLoop2Map := make(map[string]interface{})
			dataLoop2Map["IsValueSecret"] = dataLoop2Tmp["is_value_secret"]
			dataLoop2Map["Key"] = dataLoop2Tmp["key"]
			dataLoop2Map["Value"] = dataLoop2Tmp["value"]
			localMaps2 = append(localMaps2, dataLoop2Map)
		}
		oAuthHttpParameters["QueryStringParameters"] = localMaps2

		if len(oAuthHttpParameters) > 0 {
			oAuthParameters["OAuthHttpParameters"] = oAuthHttpParameters
		}

		if len(oAuthParameters) > 0 {
			authParameters["OAuthParameters"] = oAuthParameters
		}

		authParametersJson, err := json.Marshal(authParameters)
		if err != nil {
			return WrapError(err)
		}
		request["AuthParameters"] = string(authParametersJson)
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("eventbridge", "2020-04-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_event_bridge_connection", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.Data.ConnectionName", response)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudEventBridgeConnectionRead(d, meta)
}

func resourceAliCloudEventBridgeConnectionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	eventBridgeServiceV2 := EventBridgeServiceV2{client}

	objectRaw, err := eventBridgeServiceV2.DescribeEventBridgeConnection(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_event_bridge_connection DescribeEventBridgeConnection Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["GmtCreate"])
	d.Set("description", objectRaw["Description"])
	d.Set("connection_name", objectRaw["ConnectionName"])

	authParametersMaps := make([]map[string]interface{}, 0)
	authParametersMap := make(map[string]interface{})
	authParametersRaw := make(map[string]interface{})
	if objectRaw["AuthParameters"] != nil {
		authParametersRaw = objectRaw["AuthParameters"].(map[string]interface{})
	}
	if len(authParametersRaw) > 0 {
		authParametersMap["authorization_type"] = authParametersRaw["AuthorizationType"]

		apiKeyAuthParametersMaps := make([]map[string]interface{}, 0)
		apiKeyAuthParametersMap := make(map[string]interface{})
		apiKeyAuthParametersRaw := make(map[string]interface{})
		if authParametersRaw["ApiKeyAuthParameters"] != nil {
			apiKeyAuthParametersRaw = authParametersRaw["ApiKeyAuthParameters"].(map[string]interface{})
		}
		if len(apiKeyAuthParametersRaw) > 0 {
			apiKeyAuthParametersMap["api_key_name"] = apiKeyAuthParametersRaw["ApiKeyName"]
			apiKeyAuthParametersMap["api_key_value"] = apiKeyAuthParametersRaw["ApiKeyValue"]

			apiKeyAuthParametersMaps = append(apiKeyAuthParametersMaps, apiKeyAuthParametersMap)
		}
		authParametersMap["api_key_auth_parameters"] = apiKeyAuthParametersMaps
		basicAuthParametersMaps := make([]map[string]interface{}, 0)
		basicAuthParametersMap := make(map[string]interface{})
		basicAuthParametersRaw := make(map[string]interface{})
		if authParametersRaw["BasicAuthParameters"] != nil {
			basicAuthParametersRaw = authParametersRaw["BasicAuthParameters"].(map[string]interface{})
		}
		if len(basicAuthParametersRaw) > 0 {
			basicAuthParametersMap["password"] = basicAuthParametersRaw["Password"]
			basicAuthParametersMap["username"] = basicAuthParametersRaw["Username"]

			basicAuthParametersMaps = append(basicAuthParametersMaps, basicAuthParametersMap)
		}
		authParametersMap["basic_auth_parameters"] = basicAuthParametersMaps
		oauthParametersMaps := make([]map[string]interface{}, 0)
		oauthParametersMap := make(map[string]interface{})
		oAuthParametersRaw := make(map[string]interface{})
		if authParametersRaw["OAuthParameters"] != nil {
			oAuthParametersRaw = authParametersRaw["OAuthParameters"].(map[string]interface{})
		}
		if len(oAuthParametersRaw) > 0 {
			oauthParametersMap["authorization_endpoint"] = oAuthParametersRaw["AuthorizationEndpoint"]
			oauthParametersMap["http_method"] = oAuthParametersRaw["HttpMethod"]

			clientParametersMaps := make([]map[string]interface{}, 0)
			clientParametersMap := make(map[string]interface{})
			clientParametersRaw := make(map[string]interface{})
			if oAuthParametersRaw["ClientParameters"] != nil {
				clientParametersRaw = oAuthParametersRaw["ClientParameters"].(map[string]interface{})
			}
			if len(clientParametersRaw) > 0 {
				clientParametersMap["client_id"] = clientParametersRaw["ClientID"]
				clientParametersMap["client_secret"] = clientParametersRaw["ClientSecret"]

				clientParametersMaps = append(clientParametersMaps, clientParametersMap)
			}
			oauthParametersMap["client_parameters"] = clientParametersMaps
			oauthHttpParametersMaps := make([]map[string]interface{}, 0)
			oauthHttpParametersMap := make(map[string]interface{})
			oAuthHttpParametersRaw := make(map[string]interface{})
			if oAuthParametersRaw["OAuthHttpParameters"] != nil {
				oAuthHttpParametersRaw = oAuthParametersRaw["OAuthHttpParameters"].(map[string]interface{})
			}
			if len(oAuthHttpParametersRaw) > 0 {

				bodyParametersRaw := oAuthHttpParametersRaw["BodyParameters"]
				bodyParametersMaps := make([]map[string]interface{}, 0)
				if bodyParametersRaw != nil {
					for _, bodyParametersChildRaw := range convertToInterfaceArray(bodyParametersRaw) {
						bodyParametersMap := make(map[string]interface{})
						bodyParametersChildRaw := bodyParametersChildRaw.(map[string]interface{})
						bodyParametersMap["is_value_secret"] = bodyParametersChildRaw["IsValueSecret"]
						bodyParametersMap["key"] = bodyParametersChildRaw["Key"]
						bodyParametersMap["value"] = bodyParametersChildRaw["Value"]

						bodyParametersMaps = append(bodyParametersMaps, bodyParametersMap)
					}
				}
				oauthHttpParametersMap["body_parameters"] = bodyParametersMaps
				headerParametersRaw := oAuthHttpParametersRaw["HeaderParameters"]
				headerParametersMaps := make([]map[string]interface{}, 0)
				if headerParametersRaw != nil {
					for _, headerParametersChildRaw := range convertToInterfaceArray(headerParametersRaw) {
						headerParametersMap := make(map[string]interface{})
						headerParametersChildRaw := headerParametersChildRaw.(map[string]interface{})
						headerParametersMap["is_value_secret"] = headerParametersChildRaw["IsValueSecret"]
						headerParametersMap["key"] = headerParametersChildRaw["Key"]
						headerParametersMap["value"] = headerParametersChildRaw["Value"]

						headerParametersMaps = append(headerParametersMaps, headerParametersMap)
					}
				}
				oauthHttpParametersMap["header_parameters"] = headerParametersMaps
				queryStringParametersRaw := oAuthHttpParametersRaw["QueryStringParameters"]
				queryStringParametersMaps := make([]map[string]interface{}, 0)
				if queryStringParametersRaw != nil {
					for _, queryStringParametersChildRaw := range convertToInterfaceArray(queryStringParametersRaw) {
						queryStringParametersMap := make(map[string]interface{})
						queryStringParametersChildRaw := queryStringParametersChildRaw.(map[string]interface{})
						queryStringParametersMap["is_value_secret"] = queryStringParametersChildRaw["IsValueSecret"]
						queryStringParametersMap["key"] = queryStringParametersChildRaw["Key"]
						queryStringParametersMap["value"] = queryStringParametersChildRaw["Value"]

						queryStringParametersMaps = append(queryStringParametersMaps, queryStringParametersMap)
					}
				}
				oauthHttpParametersMap["query_string_parameters"] = queryStringParametersMaps
				oauthHttpParametersMaps = append(oauthHttpParametersMaps, oauthHttpParametersMap)
			}
			oauthParametersMap["oauth_http_parameters"] = oauthHttpParametersMaps
			oauthParametersMaps = append(oauthParametersMaps, oauthParametersMap)
		}
		authParametersMap["oauth_parameters"] = oauthParametersMaps
		authParametersMaps = append(authParametersMaps, authParametersMap)
	}
	if err := d.Set("auth_parameters", authParametersMaps); err != nil {
		return err
	}
	networkParametersMaps := make([]map[string]interface{}, 0)
	networkParametersMap := make(map[string]interface{})
	networkParametersRaw := make(map[string]interface{})
	if objectRaw["NetworkParameters"] != nil {
		networkParametersRaw = objectRaw["NetworkParameters"].(map[string]interface{})
	}
	if len(networkParametersRaw) > 0 {
		networkParametersMap["network_type"] = networkParametersRaw["NetworkType"]
		networkParametersMap["security_group_id"] = networkParametersRaw["SecurityGroupId"]
		networkParametersMap["vpc_id"] = networkParametersRaw["VpcId"]
		networkParametersMap["vswitche_id"] = networkParametersRaw["VswitcheId"]

		networkParametersMaps = append(networkParametersMaps, networkParametersMap)
	}
	if err := d.Set("network_parameters", networkParametersMaps); err != nil {
		return err
	}

	d.Set("connection_name", d.Id())

	return nil
}

func resourceAliCloudEventBridgeConnectionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "UpdateConnection"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ConnectionName"] = d.Id()

	if d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	if d.HasChange("network_parameters") {
		update = true
	}
	networkParameters := make(map[string]interface{})

	if v := d.Get("network_parameters"); v != nil {
		networkType1, _ := jsonpath.Get("$[0].network_type", v)
		if networkType1 != nil && networkType1 != "" {
			networkParameters["NetworkType"] = networkType1
		}
		vpcId1, _ := jsonpath.Get("$[0].vpc_id", v)
		if vpcId1 != nil && vpcId1 != "" {
			networkParameters["VpcId"] = vpcId1
		}
		vswitcheId1, _ := jsonpath.Get("$[0].vswitche_id", v)
		if vswitcheId1 != nil && vswitcheId1 != "" {
			networkParameters["VswitcheId"] = vswitcheId1
		}
		securityGroupId1, _ := jsonpath.Get("$[0].security_group_id", v)
		if securityGroupId1 != nil && securityGroupId1 != "" {
			networkParameters["SecurityGroupId"] = securityGroupId1
		}

		networkParametersJson, err := json.Marshal(networkParameters)
		if err != nil {
			return WrapError(err)
		}
		request["NetworkParameters"] = string(networkParametersJson)
	}

	if d.HasChange("auth_parameters") {
		update = true
		authParameters := make(map[string]interface{})

		if v := d.Get("auth_parameters"); v != nil {
			authorizationType1, _ := jsonpath.Get("$[0].authorization_type", v)
			if authorizationType1 != nil && authorizationType1 != "" {
				authParameters["AuthorizationType"] = authorizationType1
			}
			apiKeyAuthParameters := make(map[string]interface{})
			apiKeyName1, _ := jsonpath.Get("$[0].api_key_auth_parameters[0].api_key_name", d.Get("auth_parameters"))
			if apiKeyName1 != nil && apiKeyName1 != "" {
				apiKeyAuthParameters["ApiKeyName"] = apiKeyName1
			}
			apiKeyValue1, _ := jsonpath.Get("$[0].api_key_auth_parameters[0].api_key_value", d.Get("auth_parameters"))
			if apiKeyValue1 != nil && apiKeyValue1 != "" {
				apiKeyAuthParameters["ApiKeyValue"] = apiKeyValue1
			}

			if len(apiKeyAuthParameters) > 0 {
				authParameters["ApiKeyAuthParameters"] = apiKeyAuthParameters
			}
			basicAuthParameters := make(map[string]interface{})
			password1, _ := jsonpath.Get("$[0].basic_auth_parameters[0].password", d.Get("auth_parameters"))
			if password1 != nil && password1 != "" {
				basicAuthParameters["Password"] = password1
			}
			username1, _ := jsonpath.Get("$[0].basic_auth_parameters[0].username", d.Get("auth_parameters"))
			if username1 != nil && username1 != "" {
				basicAuthParameters["Username"] = username1
			}

			if len(basicAuthParameters) > 0 {
				authParameters["BasicAuthParameters"] = basicAuthParameters
			}
			oAuthParameters := make(map[string]interface{})
			authorizationEndpoint1, _ := jsonpath.Get("$[0].oauth_parameters[0].authorization_endpoint", d.Get("auth_parameters"))
			if authorizationEndpoint1 != nil && authorizationEndpoint1 != "" {
				oAuthParameters["AuthorizationEndpoint"] = authorizationEndpoint1
			}
			clientParameters := make(map[string]interface{})
			clientId, _ := jsonpath.Get("$[0].oauth_parameters[0].client_parameters[0].client_id", d.Get("auth_parameters"))
			if clientId != nil && clientId != "" {
				clientParameters["ClientID"] = clientId
			}
			clientSecret1, _ := jsonpath.Get("$[0].oauth_parameters[0].client_parameters[0].client_secret", d.Get("auth_parameters"))
			if clientSecret1 != nil && clientSecret1 != "" {
				clientParameters["ClientSecret"] = clientSecret1
			}

			if len(clientParameters) > 0 {
				oAuthParameters["ClientParameters"] = clientParameters
			}
			httpMethod1, _ := jsonpath.Get("$[0].oauth_parameters[0].http_method", d.Get("auth_parameters"))
			if httpMethod1 != nil && httpMethod1 != "" {
				oAuthParameters["HttpMethod"] = httpMethod1
			}
			oAuthHttpParameters := make(map[string]interface{})
			localData, err := jsonpath.Get("$[0].oauth_parameters[0].oauth_http_parameters[0].body_parameters", v)
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
				dataLoopMap["IsValueSecret"] = dataLoopTmp["is_value_secret"]
				dataLoopMap["Key"] = dataLoopTmp["key"]
				dataLoopMap["Value"] = dataLoopTmp["value"]
				localMaps = append(localMaps, dataLoopMap)
			}
			oAuthHttpParameters["BodyParameters"] = localMaps

			localData1, err := jsonpath.Get("$[0].oauth_parameters[0].oauth_http_parameters[0].header_parameters", v)
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
				dataLoop1Map["IsValueSecret"] = dataLoop1Tmp["is_value_secret"]
				dataLoop1Map["Key"] = dataLoop1Tmp["key"]
				dataLoop1Map["Value"] = dataLoop1Tmp["value"]
				localMaps1 = append(localMaps1, dataLoop1Map)
			}
			oAuthHttpParameters["HeaderParameters"] = localMaps1

			localData2, err := jsonpath.Get("$[0].oauth_parameters[0].oauth_http_parameters[0].query_string_parameters", v)
			if err != nil {
				localData2 = make([]interface{}, 0)
			}
			localMaps2 := make([]interface{}, 0)
			for _, dataLoop2 := range convertToInterfaceArray(localData2) {
				dataLoop2Tmp := make(map[string]interface{})
				if dataLoop2 != nil {
					dataLoop2Tmp = dataLoop2.(map[string]interface{})
				}
				dataLoop2Map := make(map[string]interface{})
				dataLoop2Map["IsValueSecret"] = dataLoop2Tmp["is_value_secret"]
				dataLoop2Map["Key"] = dataLoop2Tmp["key"]
				dataLoop2Map["Value"] = dataLoop2Tmp["value"]
				localMaps2 = append(localMaps2, dataLoop2Map)
			}
			oAuthHttpParameters["QueryStringParameters"] = localMaps2

			if len(oAuthHttpParameters) > 0 {
				oAuthParameters["OAuthHttpParameters"] = oAuthHttpParameters
			}

			if len(oAuthParameters) > 0 {
				authParameters["OAuthParameters"] = oAuthParameters
			}

			authParametersJson, err := json.Marshal(authParameters)
			if err != nil {
				return WrapError(err)
			}
			request["AuthParameters"] = string(authParametersJson)
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("eventbridge", "2020-04-01", action, query, request, true)
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

	return resourceAliCloudEventBridgeConnectionRead(d, meta)
}

func resourceAliCloudEventBridgeConnectionDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteConnection"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["ConnectionName"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("eventbridge", "2020-04-01", action, query, request, true)
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
