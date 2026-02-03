package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test EventBridge Connection. >>> Resource test cases, automatically generated.
// Case testConnection 3084
func TestAccAliCloudEventBridgeConnection_basic3084(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_event_bridge_connection.default"
	ra := resourceAttrInit(resourceId, AlicloudEventBridgeConnectionMap3084)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EventBridgeServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEventBridgeConnection")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceventbridge%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEventBridgeConnectionBasicDependence3084)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-qingdao"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"connection_name": name,
					"description":     "test-connection-basic-pre",
					"network_parameters": []map[string]interface{}{
						{
							"network_type":      "PublicNetwork",
							"vpc_id":            "eb-cn-huhehaote/vpc-hp3bdy0vbee0vb87fq2i6",
							"vswitche_id":       "vsw-hp3uinuttt9qbl27482v9",
							"security_group_id": "/sg-hp37abv61w1mpsuc0zco",
						},
					},
					"auth_parameters": []map[string]interface{}{
						{
							"api_key_auth_parameters": []map[string]interface{}{
								{
									"api_key_name":  "Token",
									"api_key_value": "Token-value",
								},
							},
							"basic_auth_parameters": []map[string]interface{}{
								{
									"password": "admin",
									"username": "admin",
								},
							},
							"oauth_parameters": []map[string]interface{}{
								{
									"authorization_endpoint": "http://127.0.0.1:8080",
									"client_parameters": []map[string]interface{}{
										{
											"client_secret": "ClientSecret",
											"client_id":     "ClientId",
										},
									},
									"http_method": "POST",
									"oauth_http_parameters": []map[string]interface{}{
										{
											"body_parameters": []map[string]interface{}{
												{
													"is_value_secret": "true",
													"key":             "name",
													"value":           "name",
												},
											},
											"header_parameters": []map[string]interface{}{
												{
													"is_value_secret": "true",
													"key":             "name",
													"value":           "name",
												},
											},
											"query_string_parameters": []map[string]interface{}{
												{
													"is_value_secret": "true",
													"key":             "name",
													"value":           "name",
												},
											},
										},
									},
								},
							},
							"authorization_type": "BASIC_AUTH",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"connection_name": name,
						"description":     "test-connection-basic-pre",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test-connection-basic-pre-update",
					"auth_parameters": []map[string]interface{}{
						{
							"authorization_type": "BASIC_AUTH",
							"basic_auth_parameters": []map[string]interface{}{
								{
									"password": "admin",
									"username": "admin",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "test-connection-basic-pre-update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test-connection-basic-pre-update-api-key",
					"network_parameters": []map[string]interface{}{
						{
							"network_type": "PublicNetwork",
						},
					},
					"auth_parameters": []map[string]interface{}{
						{
							"authorization_type": "API_KEY_AUTH",
							"api_key_auth_parameters": []map[string]interface{}{
								{
									"api_key_name":  "Token",
									"api_key_value": "Token-value",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "test-connection-basic-pre-update-api-key",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test-connection-basic-pre-update-oauth",
					"network_parameters": []map[string]interface{}{
						{
							"network_type": "PublicNetwork",
						},
					},
					"auth_parameters": []map[string]interface{}{
						{
							"authorization_type": "OAUTH_AUTH",
							"oauth_parameters": []map[string]interface{}{
								{
									"authorization_endpoint": "http://127.0.0.1:8080",
									"client_parameters": []map[string]interface{}{
										{
											"client_secret": "clientSecret",
											"client_id":     "clientId",
										},
									},
									"http_method": "POST",
									"oauth_http_parameters": []map[string]interface{}{
										{
											"body_parameters": []map[string]interface{}{
												{
													"is_value_secret": "true",
													"key":             "name",
													"value":           "name",
												},
											},
											"header_parameters": []map[string]interface{}{
												{
													"is_value_secret": "true",
													"key":             "age",
													"value":           "18",
												},
											},
											"query_string_parameters": []map[string]interface{}{
												{
													"is_value_secret": "true",
													"key":             "name",
													"value":           "name",
												},
											},
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "test-connection-basic-pre-update-oauth",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auth_parameters"},
			},
		},
	})
}

var AlicloudEventBridgeConnectionMap3084 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEventBridgeConnectionBasicDependence3084(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case testConnection_zhaohai_basic 3561
func TestAccAliCloudEventBridgeConnection_basic3561(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_event_bridge_connection.default"
	ra := resourceAttrInit(resourceId, AlicloudEventBridgeConnectionMap3561)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EventBridgeServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEventBridgeConnection")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceventbridge%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEventBridgeConnectionBasicDependence3561)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-qingdao"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"connection_name": name,
					"description":     "test-connection-basic-pre",
					"network_parameters": []map[string]interface{}{
						{
							"network_type":      "PublicNetwork",
							"vpc_id":            "eb-cn-huhehaote/vpc-hp3bdy0vbee0vb87fq2i6",
							"vswitche_id":       "vsw-hp3uinuttt9qbl27482v9",
							"security_group_id": "/sg-hp37abv61w1mpsuc0zco",
						},
					},
					"auth_parameters": []map[string]interface{}{
						{
							"api_key_auth_parameters": []map[string]interface{}{
								{
									"api_key_name":  "Token",
									"api_key_value": "Token-value",
								},
							},
							"basic_auth_parameters": []map[string]interface{}{
								{
									"password": "admin",
									"username": "admin",
								},
							},
							"oauth_parameters": []map[string]interface{}{
								{
									"authorization_endpoint": "http://127.0.0.1:8080",
									"client_parameters": []map[string]interface{}{
										{
											"client_secret": "ClientSecret",
											"client_id":     "ClientId",
										},
									},
									"http_method": "POST",
									"oauth_http_parameters": []map[string]interface{}{
										{
											"body_parameters": []map[string]interface{}{
												{
													"is_value_secret": "true",
													"key":             "name",
													"value":           "name",
												},
												{
													"is_value_secret": "true",
													"key":             "name",
													"value":           "name",
												},
												{
													"is_value_secret": "true",
													"key":             "name",
													"value":           "name",
												},
											},
											"header_parameters": []map[string]interface{}{
												{
													"is_value_secret": "true",
													"key":             "name",
													"value":           "name",
												},
												{
													"is_value_secret": "true",
													"key":             "name",
													"value":           "name",
												},
												{
													"is_value_secret": "true",
													"key":             "name",
													"value":           "name",
												},
											},
											"query_string_parameters": []map[string]interface{}{
												{
													"is_value_secret": "true",
													"key":             "name",
													"value":           "name",
												},
												{
													"is_value_secret": "true",
													"key":             "name",
													"value":           "name",
												},
												{
													"is_value_secret": "true",
													"key":             "name",
													"value":           "name",
												},
											},
										},
									},
								},
							},
							"authorization_type": "API_KEY_AUTH",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"connection_name": name,
						"description":     "test-connection-basic-pre",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test-connection-basic-pre-update",
					"network_parameters": []map[string]interface{}{
						{
							"network_type":      "PrivateNetwork",
							"vpc_id":            "eb-test/vpc-m5ecylo6dsk459zp1vp2k",
							"vswitche_id":       "vsw-m5eozz30ol7g427ucp7e7",
							"security_group_id": "/sg-m5eaoau86plrhghr3eda",
						},
					},
					"auth_parameters": []map[string]interface{}{
						{
							"authorization_type": "BASIC_AUTH",
							"basic_auth_parameters": []map[string]interface{}{
								{
									"password": "admin1",
									"username": "admin1",
								},
							},
							"api_key_auth_parameters": []map[string]interface{}{
								{
									"api_key_name":  "Token1",
									"api_key_value": "Token-value1",
								},
							},
							"oauth_parameters": []map[string]interface{}{
								{
									"authorization_endpoint": "http://127.0.0.1:8081",
									"client_parameters": []map[string]interface{}{
										{
											"client_secret": "ClientSecret1",
											"client_id":     "ClientId1",
										},
									},
									"http_method": "GET",
									"oauth_http_parameters": []map[string]interface{}{
										{
											"body_parameters": []map[string]interface{}{
												{
													"is_value_secret": "false",
													"key":             "name1",
													"value":           "name1",
												},
												{
													"is_value_secret": "false",
													"key":             "name1",
													"value":           "name1",
												},
											},
											"header_parameters": []map[string]interface{}{
												{
													"is_value_secret": "false",
													"key":             "name1",
													"value":           "name1",
												},
												{
													"is_value_secret": "false",
													"key":             "name1",
													"value":           "name1",
												},
											},
											"query_string_parameters": []map[string]interface{}{
												{
													"is_value_secret": "false",
													"key":             "name1",
													"value":           "name1",
												},
												{
													"is_value_secret": "false",
													"key":             "name1",
													"value":           "name1",
												},
											},
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "test-connection-basic-pre-update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"network_parameters": []map[string]interface{}{
						{
							"network_type": "PublicNetwork",
						},
					},
					"auth_parameters": []map[string]interface{}{
						{
							"authorization_type": "BASIC_AUTH",
							"api_key_auth_parameters": []map[string]interface{}{
								{
									"api_key_name":  "Token2",
									"api_key_value": "Token-value2",
								},
							},
							"basic_auth_parameters": []map[string]interface{}{
								{
									"password": "admin2",
									"username": "admin2",
								},
							},
							"oauth_parameters": []map[string]interface{}{
								{
									"authorization_endpoint": "http://127.0.0.1:8083",
									"client_parameters": []map[string]interface{}{
										{
											"client_secret": "ClientSecret2",
											"client_id":     "ClientId2",
										},
									},
									"http_method": "GET",
									"oauth_http_parameters": []map[string]interface{}{
										{
											"body_parameters": []map[string]interface{}{
												{
													"is_value_secret": "false",
													"key":             "name2",
													"value":           "name2",
												},
												{
													"is_value_secret": "false",
													"key":             "nam4",
													"value":           "name4",
												},
												{
													"is_value_secret": "true",
													"key":             "name5",
													"value":           "name5",
												},
											},
											"header_parameters": []map[string]interface{}{
												{
													"is_value_secret": "true",
													"key":             "name6",
													"value":           "name6",
												},
												{
													"is_value_secret": "false",
													"key":             "name2",
													"value":           "name2",
												},
												{
													"is_value_secret": "true",
													"key":             "name8",
													"value":           "name8",
												},
											},
											"query_string_parameters": []map[string]interface{}{
												{
													"is_value_secret": "false",
													"key":             "name9",
													"value":           "name9",
												},
												{
													"is_value_secret": "true",
													"key":             "name10",
													"value":           "name10",
												},
												{
													"is_value_secret": "true",
													"key":             "name11",
													"value":           "name11",
												},
											},
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"network_parameters": []map[string]interface{}{
						{
							"network_type": "PublicNetwork",
						},
					},
					"auth_parameters": []map[string]interface{}{
						{
							"oauth_parameters": []map[string]interface{}{
								{
									"oauth_http_parameters": []map[string]interface{}{
										{},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auth_parameters"},
			},
		},
	})
}

var AlicloudEventBridgeConnectionMap3561 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEventBridgeConnectionBasicDependence3561(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test EventBridge Connection. <<< Resource test cases, automatically generated.

func TestAccAliCloudEventBridgeConnection_basic0(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.EventBridgeConnectionSupportRegions)
	resourceId := "alicloud_event_bridge_connection.default"
	ra := resourceAttrInit(resourceId, AliCloudEventBridgeConnectionMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EventBridgeServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEventBridgeConnection")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%seventbridgeconnection%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudEventBridgeConnectionBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"connection_name": name,
					"network_parameters": []map[string]interface{}{
						{
							"network_type": "PublicNetwork",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"connection_name":      name,
						"network_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test-connection-basic-pre",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "test-connection-basic-pre",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"network_parameters": []map[string]interface{}{
						{
							"network_type":      "PrivateNetwork",
							"vpc_id":            "${alicloud_vpc.default.id}",
							"vswitche_id":       "${alicloud_vswitch.default.id}",
							"security_group_id": "${alicloud_security_group.default.id}",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"network_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auth_parameters": []map[string]interface{}{
						{
							"authorization_type": "API_KEY_AUTH",
							"api_key_auth_parameters": []map[string]interface{}{
								{
									"api_key_name":  "Token",
									"api_key_value": "Token-value",
								},
							},
							"basic_auth_parameters": []map[string]interface{}{
								{
									"username": "admin",
									"password": "admin",
								},
							},
							"oauth_parameters": []map[string]interface{}{
								{
									"authorization_endpoint": "http://127.0.0.1:8080",
									"http_method":            "POST",
									"client_parameters": []map[string]interface{}{
										{
											"client_id":     "ClientId",
											"client_secret": "ClientSecret",
										},
									},
									"oauth_http_parameters": []map[string]interface{}{
										{
											"header_parameters": []map[string]interface{}{
												{
													"key":             "name",
													"value":           "name",
													"is_value_secret": "true",
												},
											},
											"body_parameters": []map[string]interface{}{
												{
													"key":             "name",
													"value":           "name",
													"is_value_secret": "true",
												},
											},
											"query_string_parameters": []map[string]interface{}{
												{
													"key":             "name",
													"value":           "name",
													"is_value_secret": "true",
												},
											},
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auth_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test-connection-basic-pre-update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "test-connection-basic-pre-update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"network_parameters": []map[string]interface{}{
						{
							"network_type":      "PublicNetwork",
							"vpc_id":            "${alicloud_vpc.default.id}",
							"vswitche_id":       "${alicloud_vswitch.default.id}",
							"security_group_id": "${alicloud_security_group.default.id}",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"network_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auth_parameters": []map[string]interface{}{
						{
							"authorization_type": "BASIC_AUTH",
							"api_key_auth_parameters": []map[string]interface{}{
								{
									"api_key_name":  "Token-update",
									"api_key_value": "Token-value-update",
								},
							},
							"basic_auth_parameters": []map[string]interface{}{
								{
									"username": "admin-update",
									"password": "admin-update",
								},
							},
							"oauth_parameters": []map[string]interface{}{
								{
									"authorization_endpoint": "http://127.0.0.1:8080",
									"http_method":            "POST",
									"client_parameters": []map[string]interface{}{
										{
											"client_id":     "clientId",
											"client_secret": "clientSecret",
										},
									},
									"oauth_http_parameters": []map[string]interface{}{
										{
											"header_parameters": []map[string]interface{}{
												{
													"key":             "name-update",
													"value":           "name-update",
													"is_value_secret": "false",
												},
											},
											"body_parameters": []map[string]interface{}{
												{
													"key":             "name-update",
													"value":           "name-update",
													"is_value_secret": "false",
												},
											},
											"query_string_parameters": []map[string]interface{}{
												{
													"key":             "name-update",
													"value":           "name-update",
													"is_value_secret": "false",
												},
											},
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auth_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auth_parameters": []map[string]interface{}{
						{
							"authorization_type": "OAUTH_AUTH",
							"api_key_auth_parameters": []map[string]interface{}{
								{
									"api_key_name":  "Token",
									"api_key_value": "Token-value",
								},
							},
							"basic_auth_parameters": []map[string]interface{}{
								{
									"username": "admin",
									"password": "admin",
								},
							},
							"oauth_parameters": []map[string]interface{}{
								{
									"authorization_endpoint": "http://127.0.0.1:8080",
									"http_method":            "POST",
									"client_parameters": []map[string]interface{}{
										{
											"client_id":     "ClientId",
											"client_secret": "ClientSecret",
										},
									},
									"oauth_http_parameters": []map[string]interface{}{
										{
											"header_parameters": []map[string]interface{}{
												{
													"key":             "name",
													"value":           "name",
													"is_value_secret": "true",
												},
											},
											"body_parameters": []map[string]interface{}{
												{
													"key":             "name",
													"value":           "name",
													"is_value_secret": "true",
												},
											},
											"query_string_parameters": []map[string]interface{}{
												{
													"key":             "name",
													"value":           "name",
													"is_value_secret": "true",
												},
											},
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auth_parameters.#": "1",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAliCloudEventBridgeConnection_basic0_twin(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.EventBridgeConnectionSupportRegions)
	resourceId := "alicloud_event_bridge_connection.default"
	ra := resourceAttrInit(resourceId, AliCloudEventBridgeConnectionMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EventBridgeServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEventBridgeConnection")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%seventbridgeconnection%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudEventBridgeConnectionBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"connection_name": name,
					"description":     "test-connection-basic-pre-update-oauth",
					"network_parameters": []map[string]interface{}{
						{
							"network_type":      "PrivateNetwork",
							"vpc_id":            "${alicloud_vpc.default.id}",
							"vswitche_id":       "${alicloud_vswitch.default.id}",
							"security_group_id": "${alicloud_security_group.default.id}",
						},
					},
					"auth_parameters": []map[string]interface{}{
						{
							"authorization_type": "API_KEY_AUTH",
							"api_key_auth_parameters": []map[string]interface{}{
								{
									"api_key_name":  "Token",
									"api_key_value": "Token-value",
								},
							},
							"basic_auth_parameters": []map[string]interface{}{
								{
									"username": "admin",
									"password": "admin",
								},
							},
							"oauth_parameters": []map[string]interface{}{
								{
									"authorization_endpoint": "http://127.0.0.1:8080",
									"http_method":            "POST",
									"client_parameters": []map[string]interface{}{
										{
											"client_id":     "ClientId",
											"client_secret": "ClientSecret",
										},
									},
									"oauth_http_parameters": []map[string]interface{}{
										{
											"header_parameters": []map[string]interface{}{
												{
													"key":             "name",
													"value":           "name",
													"is_value_secret": "true",
												},
											},
											"body_parameters": []map[string]interface{}{
												{
													"key":             "name",
													"value":           "name",
													"is_value_secret": "true",
												},
											},
											"query_string_parameters": []map[string]interface{}{
												{
													"key":             "name",
													"value":           "name",
													"is_value_secret": "true",
												},
											},
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"connection_name":      name,
						"description":          "test-connection-basic-pre-update-oauth",
						"network_parameters.#": "1",
						"auth_parameters.#":    "1",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

var AliCloudEventBridgeConnectionMap0 = map[string]string{
	"create_time": CHECKSET,
}

func AliCloudEventBridgeConnectionBasicDependence0(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_zones" "default" {
	}

	resource "alicloud_vpc" "default" {
  		vpc_name   = var.name
  		cidr_block = "172.16.0.0/16"
	}

	resource "alicloud_vswitch" "default" {
  		vpc_id       = alicloud_vpc.default.id
  		cidr_block   = "172.16.0.0/24"
  		zone_id      = data.alicloud_zones.default.zones[0].id
  		vswitch_name = var.name
	}

	resource "alicloud_security_group" "default" {
  		name   = var.name
  		vpc_id = alicloud_vswitch.default.vpc_id
	}
`, name)
}
