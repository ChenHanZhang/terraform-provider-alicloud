package alicloud

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestAccAliCloudDdosCooDomainResource_https_ext(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AliCloudDdoscooDomainResourceMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandString(10)
	name := fmt.Sprintf("tf-testacc%s.alibaba.com", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudDdoscooDomainResourceBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.DdoscooSupportedRegions)
			testAccPreCheckWithAccountSiteType(t, DomesticSite)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"domain":       name,
					"instance_ids": []string{"${data.alicloud_ddoscoo_instances.default.ids.0}"},
					"real_servers": []string{"177.167.32.11", "177.167.32.12", "177.167.32.13"},
					"rs_type":      `0`,
					"https_ext":    `{\"Http2\":1,\"Http2https\":0,\"Https2http\":0}`,
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{"80", "8080"},
							"proxy_type":  "http",
						},
						{
							"proxy_ports": []string{"443", "8443"},
							"proxy_type":  "https",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"domain":         name,
						"instance_ids.#": "1",
						"real_servers.#": "3",
						"rs_type":        "0",
						"https_ext":      CHECKSET,
						"proxy_types.#":  "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"https_ext": `{\"Http2\":1,\"Http2https\":0,\"Https2http\":0}`,
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{"443"},
							"proxy_type":  "https",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"https_ext":     CHECKSET,
						"proxy_types.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_ids": []string{"${data.alicloud_ddoscoo_instances.default.ids.0}", "${data.alicloud_ddoscoo_instances.default.ids.1}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_ids.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"real_servers": []string{"aliyun.com", "taobao.com", "alibaba.com"},
					"rs_type":      `1`,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"real_servers.#": "3",
						"rs_type":        "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_ids": []string{"${data.alicloud_ddoscoo_instances.default.ids.0}"},
					"real_servers": []string{"177.167.32.11", "177.167.32.12", "177.167.32.13", "177.167.32.14", "177.167.32.15"},
					"rs_type":      `0`,
					"https_ext":    `{\"Http2\":0,\"Http2https\":0,\"Https2http\":0}`,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_ids.#": "1",
						"real_servers.#": "5",
						"rs_type":        "0",
						"https_ext":      CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ocsp_enabled": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ocsp_enabled": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ocsp_enabled": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ocsp_enabled": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ai_mode": "watch",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ai_mode": "watch",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ai_template": "level30",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ai_template": "level30",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bw_list_enable": "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bw_list_enable": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bw_list_enable": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bw_list_enable": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"black_list": []string{"3.3.3.3", "6.6.6.6", "5.5.5.5"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"black_list.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"black_list": []string{"3.3.3.3"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"black_list.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"black_list": []string{"3.3.3.3", "6.6.6.6", "5.5.5.5"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"black_list.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cc_global_switch": "close",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cc_global_switch": "close",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"white_list": []string{"3.3.3.3", "1.1.1.1", "2.2.2.2"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"white_list.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"white_list": []string{"1.1.1.1"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"white_list.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"white_list": []string{"3.3.3.3", "1.1.1.1", "2.2.2.2"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"white_list.#": "3",
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

func TestAccAliCloudDdosCooDomainResource_none_https_ext(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AliCloudDdoscooDomainResourceMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandInt()
	name := fmt.Sprintf("tf-testacc%d.alibaba.com", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudDdoscooDomainResourceBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.DdoscooSupportedRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"domain":           name,
					"instance_ids":     []string{"${data.alicloud_ddoscoo_instances.default.ids.0}"},
					"real_servers":     []string{"177.167.32.11", "177.167.32.12", "177.167.32.13"},
					"rs_type":          `0`,
					"ocsp_enabled":     "true",
					"custom_headers":   "{\\\"22\\\":\\\"$ReqClientIP\\\",\\\"77\\\":\\\"88\\\",\\\"99\\\":\\\"$ReqClientPort\\\"}",
					"ai_mode":          "watch",
					"ai_template":      "level30",
					"bw_list_enable":   "0",
					"cc_global_switch": "close",
					"black_list":       []string{"3.3.3.3", "6.6.6.6", "5.5.5.5"},
					"white_list":       []string{"3.3.3.3", "1.1.1.1", "2.2.2.2"},
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{"80", "8080"},
							"proxy_type":  "http",
						},
						{
							"proxy_ports": []string{"443", "8443"},
							"proxy_type":  "https",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"domain":           name,
						"instance_ids.#":   "1",
						"real_servers.#":   "3",
						"rs_type":          "0",
						"ocsp_enabled":     "true",
						"custom_headers":   CHECKSET,
						"ai_mode":          "watch",
						"ai_template":      "level30",
						"bw_list_enable":   "0",
						"cc_global_switch": "close",
						"black_list.#":     "3",
						"white_list.#":     "3",
						"proxy_types.#":    "2",
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

var AliCloudDdoscooDomainResourceMap0 = map[string]string{
	"https_ext":      CHECKSET,
	"cname":          CHECKSET,
	"instance_ids.#": "1",
	"proxy_types.#":  "1",
	"real_servers.#": "1",
	"rs_type":        "0",
}

func AliCloudDdoscooDomainResourceBasicDependence0(name string) string {
	return fmt.Sprintf(`
	data "alicloud_ddoscoo_instances" "default" {
	}
`)
}

// lintignore: R001
func TestUnitAliCloudDdoscooDomainResource(t *testing.T) {
	p := Provider().(*schema.Provider).ResourcesMap
	checkoutSupportedRegions(t, true, connectivity.DdoscooSupportedRegions)
	dInit, _ := schema.InternalMap(p["alicloud_ddoscoo_domain_resource"].Schema).Data(nil, nil)
	dExisted, _ := schema.InternalMap(p["alicloud_ddoscoo_domain_resource"].Schema).Data(nil, nil)
	dInit.MarkNewResource()
	attributes := map[string]interface{}{
		"domain":    "CreateDomainResourceValue",
		"https_ext": "CreateDomainResourceValue",
		"proxy_types": []map[string]interface{}{
			{
				"proxy_ports": []int{443},
				"proxy_type":  "https",
			},
		},
		"instance_ids": []string{"CreateDomainResourceValue"},
		"real_servers": []string{"CreateDomainResourceValue"},
		"rs_type":      1,
	}
	for key, value := range attributes {
		err := dInit.Set(key, value)
		assert.Nil(t, err)
		err = dExisted.Set(key, value)
		assert.Nil(t, err)
		if err != nil {
			log.Printf("[ERROR] the field %s setting error", key)
		}
	}
	region := os.Getenv("ALICLOUD_REGION")
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		t.Skipf("Skipping the test case with err: %s", err)
		t.Skipped()
	}
	rawClient = rawClient.(*connectivity.AliyunClient)
	ReadMockResponse := map[string]interface{}{
		// DescribeDomainResource
		"WebRules": []interface{}{
			map[string]interface{}{
				"Domain":   "CreateDomainResourceValue",
				"HttpsExt": "CreateDomainResourceValue",
				"InstanceIds": []interface{}{
					"CreateDomainResourceValue",
				},
				"ProxyTypes": []interface{}{
					map[string]interface{}{
						"ProxyPorts": []int{443},
						"ProxyType":  "https",
					},
				},
				"RealServers": "CreateDomainResourceValue",
				"RsType":      1,
			},
		},
	}
	CreateMockResponse := map[string]interface{}{}
	ReadMockResponseDiff := map[string]interface{}{}
	failedResponseMock := func(errorCode string) (map[string]interface{}, error) {
		return nil, &tea.SDKError{
			Code:       String(errorCode),
			Data:       String(errorCode),
			Message:    String(errorCode),
			StatusCode: tea.Int(400),
		}
	}
	notFoundResponseMock := func(errorCode string) (map[string]interface{}, error) {
		return nil, GetNotFoundErrorFromString(GetNotFoundMessage("alicloud_ddoscoo_domain_resource", errorCode))
	}
	successResponseMock := func(operationMockResponse map[string]interface{}) (map[string]interface{}, error) {
		if len(operationMockResponse) > 0 {
			mapMerge(ReadMockResponse, operationMockResponse)
		}
		return ReadMockResponse, nil
	}

	patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewDdoscooClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudDdosCooDomainResourceCreate(dInit, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	errorCodes := []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1 // a counter used to cover retry scenario; the same below
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "CreateDomainResource" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if retryIndex >= len(errorCodes)-1 {
						successResponseMock(ReadMockResponseDiff)
						return CreateMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudDdosCooDomainResourceCreate(dInit, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_ddoscoo_domain_resource"].Schema).Data(dInit.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dInit.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// Update
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewDdoscooClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudDdosCooDomainResourceUpdate(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	//ModifyDomainResource
	attributesDiff := map[string]interface{}{
		"proxy_types": []map[string]interface{}{
			{
				"proxy_ports": []int{80},
				"proxy_type":  "http",
			},
		},
		"real_servers": []string{"ModifyDomainResourceValue"},
		"rs_type":      2,
		"https_ext":    "ModifyDomainResourceValue",
		"instance_ids": []string{"ModifyDomainResourceValue"},
	}
	diff, err := newInstanceDiff("alicloud_ddoscoo_domain_resource", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_ddoscoo_domain_resource"].Schema).Data(dInit.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeDomainResource Response
		"WebRules": []interface{}{
			map[string]interface{}{
				"HttpsExt": "ModifyDomainResourceValue",
				"ProxyTypes": []interface{}{
					map[string]interface{}{
						"ProxyPorts": []int{80},
						"ProxyType":  "http",
					},
				},
				"InstanceIds": []interface{}{
					"ModifyDomainResourceValue",
				},
				"RealServers": "ModifyDomainResourceValue",
				"RsType":      2,
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "ModifyDomainResource" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if retryIndex >= len(errorCodes)-1 {
						return successResponseMock(ReadMockResponseDiff)
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudDdosCooDomainResourceUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_ddoscoo_domain_resource"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	//Read
	errorCodes = []string{"NonRetryableError", "Throttling", "nil", "{}"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DescribeDomainResource" {
				switch errorCode {
				case "{}":
					return notFoundResponseMock(errorCode)
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if errorCodes[retryIndex] == "nil" {
						return ReadMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudDdosCooDomainResourceRead(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "{}":
			assert.Nil(t, err)
		}
	}

	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewDdoscooClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudDdosCooDomainResourceDelete(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DeleteDomainResource" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if errorCodes[retryIndex] == "nil" {
						ReadMockResponse = map[string]interface{}{
							"Success": true,
						}
						return ReadMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudDdosCooDomainResourceDelete(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "nil":
			assert.Nil(t, err)
		}
	}
}

// Test DdosCoo DomainResource. >>> Resource test cases, automatically generated.
// Case 国内高防测试支持策略 12258
func TestAccAliCloudDdosCooDomainResource_basic12258(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap12258)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence12258)
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
					"rs_type":      "0",
					"ocsp_enabled": "false",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "websocket",
						},
					},
					"real_servers": []string{
						"1.1.1.1", "2.2.2.2", "3.3.3.3"},
					"domain": "testld.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.default9nJ7Ie.id}"},
					"https_ext":       "{\\\"Https2http\\\":0,\\\"Http2\\\":1,\\\"Http2https\\\":0}",
					"cert_identifier": "${alicloud_ssl_certificates_service_certificate.default3MYZEt.id}-cn-hangzhou",
					"custom_headers":  "{\\\"3444\\\":\\\"5555\\\",\\\"666\\\":\\\"$ReqClientPort\\\",\\\"77777\\\":\\\"$ReqClientIP\\\"}",
					"white_list": []string{
						"1.1.1.1"},
					"ai_template":    "level30",
					"bw_list_enable": "1",
					"ai_mode":        "defense",
					"black_list": []string{
						"2.2.2.2"},
					"cc_global_switch": "open",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":          "0",
						"ocsp_enabled":     "false",
						"proxy_types.#":    "3",
						"real_servers.#":   "3",
						"domain":           "testld.qq.com",
						"instance_ids.#":   "1",
						"https_ext":        CHECKSET,
						"cert_identifier":  CHECKSET,
						"custom_headers":   CHECKSET,
						"white_list.#":     "1",
						"ai_template":      "level30",
						"bw_list_enable":   "1",
						"ai_mode":          "defense",
						"black_list.#":     "1",
						"cc_global_switch": "open",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ocsp_enabled": "true",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80", "8011", "8022"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "websockets",
						},
					},
					"https_ext":      "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
					"custom_headers": "{\\\"22\\\":\\\"$ReqClientIP\\\",\\\"77\\\":\\\"88\\\",\\\"99\\\":\\\"$ReqClientPort\\\"}",
					"white_list": []string{
						"3.3.3.3", "1.1.1.1", "2.2.2.2"},
					"ai_template": "level60",
					"black_list": []string{
						"3.3.3.3", "4.4.4.4", "5.5.5.5"},
					"cc_global_switch": "close",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ocsp_enabled":     "true",
						"proxy_types.#":    "3",
						"https_ext":        CHECKSET,
						"custom_headers":   CHECKSET,
						"white_list.#":     "3",
						"ai_template":      "level60",
						"black_list.#":     "3",
						"cc_global_switch": "close",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type": "1",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
					},
					"real_servers": []string{
						"1.qq.com"},
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.default9nJ7Ie.id}", "${alicloud_ddoscoo_instance.defaultYi1ySx.id}", "${alicloud_ddoscoo_instance.default9KfXQ3.id}"},
					"cert_identifier": "${alicloud_ssl_certificates_service_certificate.defaultAeCL0Q.id}-cn-hangzhou",
					"custom_headers":  "{\\\"44\\\":\\\"55\\\",\\\"66\\\":\\\"$ReqClientPort\\\",\\\"77\\\":\\\"$ReqClientIP\\\"}",
					"white_list":      []string{},
					"ai_template":     "level90",
					"bw_list_enable":  "0",
					"ai_mode":         "watch",
					"black_list":      []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":         "1",
						"proxy_types.#":   "2",
						"real_servers.#":  "1",
						"instance_ids.#":  "3",
						"cert_identifier": CHECKSET,
						"custom_headers":  CHECKSET,
						"white_list.#":    "0",
						"ai_template":     "level90",
						"bw_list_enable":  "0",
						"ai_mode":         "watch",
						"black_list.#":    "0",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap12258 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence12258(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "cert_region" {
  default = "default3MYZEt.CertId-${cn-hangzhou}"
}

resource "alicloud_ddoscoo_instance" "default9nJ7Ie" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "0"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ssl_certificates_service_certificate" "default3MYZEt" {
  cert             = <<EOF
-----BEGIN CERTIFICATE-----
MIID4TCCAsmgAwIBAgIRALw5sXZD1UHDhmh/t2VTQw4wDQYJKoZIhvcNAQELBQAw
XjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU
ZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w
HhcNMjUwOTI2MDI1NDU1WhcNMjYwOTI2MDI1NDU1WjAlMQswCQYDVQQGEwJDTjEW
MBQGA1UEAxMNdGVzdGxkLnFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBAMrfIvzgwhQegAeYFBRIR2LIwWT3cnKA7dLTmQUmusSqmx/AgA1ctaw8
/BUaRCCjamkYKnbDqBSYPUGMicLUVTbgiXuupoFwGBbkHN9AyetUiV86A8hebDi0
Hp3mK6AwIX432mb8nKiM3GCjVflJRt//xOybCpkqLyXFmOQxXUunZJEUUic+JHWa
bVlBxFzd4CDnBRrw0q0JPti0322TuL9HjiGkiJp2BvnMH++qtlTjwzOxMvTYeiz8
+E+yl0kzCW+bmMZK+t39nWrX57MvggNP8KsT6YCHcGgbSPQPcfx0kBL2IAU7eWbX
Xgpat3v/zRXxcIPjvg1EBmcw2rxu8dMCAwEAAaOB0jCBzzAOBgNVHQ8BAf8EBAMC
BaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMB8GA1UdIwQYMBaAFCiB
JgXRNBo/wXMPu5PPFRw/A79/MGMGCCsGAQUFBwEBBFcwVTAhBggrBgEFBQcwAYYV
aHR0cDovL29jc3AubXlzc2wuY29tMDAGCCsGAQUFBzAChiRodHRwOi8vY2EubXlz
c2wuY29tL215c3NsdGVzdHJzYS5jcnQwGAYDVR0RBBEwD4INdGVzdGxkLnFxLmNv
bTANBgkqhkiG9w0BAQsFAAOCAQEAKtDeQoQtloF6mvMOC0AYwJ2as7XyxfKKoqPs
dW7VHuASnB5AUeSmqPz3H8+qS7IX9VZDmTr2JxPRsJ+eYXMMI3UUlHUik0BcMt3Y
JfsV6nRgKm8JwktSHCsyVPDYU3zCO6KF1tUVKa18l61Twq81+gwX6jlmRy45/kPe
6yPUYA5FrNWc5ZWs4LcEM9F5L9xkhJVS8uICU09k8pwYsmU87z5mHaRxxSYjCoF2
gUrJjy6iWYfSJRWbDDA4p+BVZMuK3bGV4K7bS2lKjUPz7EZSUKQbWrzCMEOr7E8Y
9sFvHi49Blv8zllUS3clDdsP7nYPtU0hNysA9m9+eKkggFCo9Q==
-----END CERTIFICATE-----

EOF
  certificate_name = "4"
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAyt8i/ODCFB6AB5gUFEhHYsjBZPdycoDt0tOZBSa6xKqbH8CA
DVy1rDz8FRpEIKNqaRgqdsOoFJg9QYyJwtRVNuCJe66mgXAYFuQc30DJ61SJXzoD
yF5sOLQeneYroDAhfjfaZvycqIzcYKNV+UlG3//E7JsKmSovJcWY5DFdS6dkkRRS
Jz4kdZptWUHEXN3gIOcFGvDSrQk+2LTfbZO4v0eOIaSImnYG+cwf76q2VOPDM7Ey
9Nh6LPz4T7KXSTMJb5uYxkr63f2datfnsy+CA0/wqxPpgIdwaBtI9A9x/HSQEvYg
BTt5ZtdeClq3e//NFfFwg+O+DUQGZzDavG7x0wIDAQABAoIBAF3lJb/t1NXUAgTB
xfVXOLcHXL866d0GQEyWQ9oHAWV54v8wrPPCR5F2zmOD+ykyCVOn7Ct4xif2CE/4
2F/0v5X9GZTFkmoRNA0QOi64QVxqjYQmCU8pKKSb9Rm2yNVEwZO2DR8iZu15+Ju/
rVCKMkQFkKLD5YVbaWPtjyR6lognFrzkehASnmk2xGbqPjpr6wXWVVQ/MkJd4Nwt
SFzF9veBZRjSmxFCl9yowE1WdsEmvkzQyX0bI8u/pY3z4hj9EuffQz9/zzL5xVS3
vNTN0CuRTyOnTHaB+6K/SIh2nqkJRwAb9UPTokAgUnO/EhN4W1CaGqimFCxxgSVT
yt3c/LECgYEAyvWUVtPAbDbrQhJ7L7dM5eR98xbhy7ZF4n0djJ6w6qPu/FSLTOjq
j2REtBzDn3xDo6Z/5U0vJG9P7RpleZ3s8g0vF4zN8Lu6EkDZfsRriq2sW9Nl+f0H
3zNocOVafQXyrrM3WytBbsdHwRi0oWs0z22h1pL2SZSn4TWnzgjrdUMCgYEA/+Ox
LXvAaCgoLYUZHuBhyoPmocl7a14MUzii8dnxjCIaawS7YjSXoc5mSoYdft6wgm+U
sjqGgkqTFasiSJUN+d7367Liy+aPnggCqK6rE9Hgi2T7sdQv+XtU+FMtd7cPzbkQ
WVsMHVpr7dA+bzO+IE6wgTxX8g8soSrCO1aUgDECgYAlLlQci/JGYNE8a3JRzXyy
6OcB74Ex9pRa55zQNAopEhsn8r0KO+ksl6vWayaTQwqJImlvsnIedJ3py9onK31K
4otr/wmDPoDZ8zNk+8rPvv1CXTnjUC1vAFXzyLCJEtvgkUhk1UnJZ4yHnWUJ5T/p
eCYbzxR7alZO9atmHVA1TQKBgD8Lx3SQX/iJpFSKzYSo/g8abnGAJdNvSZQbiTIT
Y4sGQAIMGWr50D5CjztfTdcbYNvSSA2dk9R4MUMOdhTx/I6K3ASLf4uDU/E4wgbz
eh0ZAbz2dXj78ZIDTA0e2T38sX0bUqbhYtu8koj2XNujIP3uxVgiGPz/thxDX7Wl
AFORAoGAEh5MIb1j3Z8n2isB9AxP+Ju6Q38AueX0iKvvjFUQiqzQKgaa16VajPw1
YDMn3aoIIA9UyghkSmKdBWXAgpRWqRLqakbN58PMdtmDRhp2qqz7xljpOBSKRs3t
G5w8hpXVQAflI/SUAKdnoQdHoezMX8WWQzQAlOlh4lfTKAPOh8E=
-----END RSA PRIVATE KEY-----
EOF
}

resource "alicloud_ssl_certificates_service_certificate" "defaultAeCL0Q" {
  certificate_name = "ld测试headers622"
  cert             = <<EOF
-----BEGIN CERTIFICATE----- MIID4TCCAsmgAwIBAgIRAJ3bnm3BlEw5iDBJ9msY+YgwDQYJKoZIhvcNAQELBQAw XjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU ZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w HhcNMjUwOTI2MDI1NjMzWhcNMjYwOTI2MDI1NjMzWjAlMQswCQYDVQQGEwJDTjEW MBQGA1UEAxMNdGVzdGxkLnFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC AQoCggEBAK0Owiyqlde8Y3QNQdtQYLZkqm7zf2HAlJksd6EZeG/KBDTDnHOEoGs9 v5xdj2Wld0OkjC1M+Y557cjW6nKp+espiGl+LJ8s4OAKssupUaXktKG4Hs2EoA+e UIvFI7OfMzgv5prG9/reAEPPMkebbpVuuS2WVqtzRc3G7CShkcv75viVandQN3Xf GezpU9G5Dhfa3Ib6etqO6LHC9wCT06LK0NKscp7NFSbeGY/AHu2qHm5EuaityWpk xRDznaMveEEvs0fav0wW/zv4WvyDP+IQ6hjFgAUSw17iSW8VJEaGJ8KG9arddjG+ oZhLkYViTseER3HS0H69qrm7Qcvks9sCAwEAAaOB0jCBzzAOBgNVHQ8BAf8EBAMC BaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMB8GA1UdIwQYMBaAFCiB JgXRNBo/wXMPu5PPFRw/A79/MGMGCCsGAQUFBwEBBFcwVTAhBggrBgEFBQcwAYYV aHR0cDovL29jc3AubXlzc2wuY29tMDAGCCsGAQUFBzAChiRodHRwOi8vY2EubXlz c2wuY29tL215c3NsdGVzdHJzYS5jcnQwGAYDVR0RBBEwD4INdGVzdGxkLnFxLmNv bTANBgkqhkiG9w0BAQsFAAOCAQEAItakLjDpt86k9nl7ZKa0nKMMov3IGfpv94yW /AUUM6cB73ygVItJGaOG9QLO2r44jdum0pyJNusJAsHPgURUb/0RlEsGYOgTuhPG +gLlr2Yo7ekI5PdQAcVNkG2vaLyVGVo7WY6OuF3kHqt09HL8HvEpj0gwxumNsCEu sEq6nFwqgqhLa0yfVT6wAdp4he3vuo+Kgla9IOUTMInVw5WerxL2htCYQRhHme7q 98QFJVxkrnSkGYaNYnIm1opwwg9U52I5vGpd5KZkR4PiM8gxeBCLbIr9+RuJdoIa jFBpKbR/44TV3e6uwRWxvip83pE3DWxQEn8Z4rlrrnTtUxPsCw== -----END CERTIFICATE-----
EOF
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY----- MIIEogIBAAKCAQEArQ7CLKqV17xjdA1B21BgtmSqbvN/YcCUmSx3oRl4b8oENMOc c4Sgaz2/nF2PZaV3Q6SMLUz5jnntyNbqcqn56ymIaX4snyzg4Aqyy6lRpeS0obge zYSgD55Qi8Ujs58zOC/mmsb3+t4AQ88yR5tulW65LZZWq3NFzcbsJKGRy/vm+JVq d1A3dd8Z7OlT0bkOF9rchvp62o7oscL3AJPTosrQ0qxyns0VJt4Zj8Ae7aoebkS5 qK3JamTFEPOdoy94QS+zR9q/TBb/O/ha/IM/4hDqGMWABRLDXuJJbxUkRoYnwob1 qt12Mb6hmEuRhWJOx4RHcdLQfr2qubtBy+Sz2wIDAQABAoIBAAXcX5gaqNt3Dlky T74pMTVMIHeEeJZrarzrRBvpHGqQyWauD0DcR4CKRVB63K3hFjJswrCQEE2SdIqe OK9scUHVFMEZ3FIBt5Xu1tJN6C15muJ1NVnZeYA96NVq9kQRiq8G9ETeoyxUU2b4 f+fr7ClUaCISmtnQnBcVew7ch+8ECSeyrnSFXE8MBDghEYGVwJdTYye7O5pbBlUI VkAjme/vZ5tgBEizTbOjJ1wToprXS90wbUN2SNrZ3O80ctnfQOYeN4FHZoKK4iZ5 YWeIY/3s/EhgHL4BpVE9xBr4l4Ji1OhVj55vnEca7V80gBLI8KsKmswgPUDErQZQ 59HeFU0CgYEA4+XTWyr/OHaxIZvrgTbroB8IQcRS8N1yp9Mhi0OxNNGNN/0mzS3J R1H0nmBqH3uiJtUHekHkaILZeB6AJsmBdF4wvqFwhinC879He4UxVJhwrNmrhyN7 e4U8JpcScfqkRSCw7IyWXA8utgJwkrwSBRWpPJH0uDzEBdG6SQUbRv0CgYEAwmXD 1HNhc6zYljOpYJJ303aa1tbo5iLesnlmKRLGH02niHsjTKlKk8OgpL0mDXZPZcBO cGBv89VpEadgmXkvSBF3EcYzznMx4bE5roODG/uZpAgYhFCtck8OiaSpMaLAsNfg iX5UoUpGC/sjzRlb4pTuYXgdJIpq7uYNESPDWbcCgYB4gkw8NkVVNzY9PnTPFBqO xjoYhNcS9Rau9e8T9EydUH3rcFr3PlHj+8ttgDA9y6OYAyf8FyHPvfl/EinT2NQy k6gU9ctJtoWXXLtQ7sKCKEXUsRmJ0VJ3o6GGCna/fLkmsL29qC3OKl0Z87pRbU/e DnE6TTfpwIArT5bBDAk7hQKBgEOyuuL2l354OCj+rsKx30bmLAAbcZoLlLBktJ4j KpnmKizObTmaTx1saDpWoLTZecb+lJJwEyzAKkn6lxp7aGrZojSvaIVB+qIVXPgY VoVdNmE4bIlPq0d+nppynHVrr10moMxhEUnmGsS7XflkFmqu4OFSuhtqlNjDhtNs N/C7AoGAVYQOLWtIyST4GnAIqnGzNHe2iBm28lUkPqaqR69x7MIxq7gr9XsLpaRo ocoes5G2ZLILhifitPWm/LDS+TX762zGb0lotm2y6pDSR1l0A5Zjbw67H6mJHgDN CaGUag785i8Fln+pJ3DkuQ0tH70wAcZooqdWk7aLs81JTLaImTY= -----END RSA PRIVATE KEY-----
EOF
}

resource "alicloud_ddoscoo_instance" "defaultYi1ySx" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "0"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ddoscoo_instance" "default9KfXQ3" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "0"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}


`, name)
}

// Case 国际高防策略 12079
func TestAccAliCloudDdosCooDomainResource_basic12079(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap12079)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence12079)
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
					"rs_type": "0",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
					},
					"real_servers": []string{
						"1.1.1.1"},
					"domain": "testcert.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultdPWcCl.id}"},
					"https_ext":      "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
					"cert_name":      "532",
					"cert":           "-----BEGIN CERTIFICATE-----\\nMIID5TCCAs2gAwIBAgIRAJ8XVm8/e0sNiF7SQKrNVpUwDQYJKoZIhvcNAQELBQAw\\nXjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU\\nZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w\\nHhcNMjUwOTI5MDY1NDI3WhcNMjYwOTI5MDY1NDI3WjAnMQswCQYDVQQGEwJDTjEY\\nMBYGA1UEAxMPdGVzdGNlcnQucXEuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A\\nMIIBCgKCAQEAqjAfE7/ToSl6okR9YhQSz/vw+NvQUU4i/p+7HKqCoLC+l+rY4HzU\\n7m7fvgfkPZtTItgaiIKURo/Y+UmIaN+JNvnaTPZTXRgkfdccdTvYOeYzHn77AzFd\\nsjCcystdpXdJ/wrtfABa5XGNrFMeXYpf9Gxp+1A56modo8wiWAS48rysfzS8vHUw\\ncYuog5Nxttg5Kaqs2w1eUxiPLsXp+mN5CIhP46m1U5cisq9J1utAbKaHeRO+pWFn\\nAu3e4AK1utaxhGdYCDMX3xwlNWFIgpNepFpdrs+H8btwsoPUkbwdYM+lFkGRI8XZ\\n4dMwQEI6p6iohthk6o9GQj+otOEqkGHyHQIDAQABo4HUMIHRMA4GA1UdDwEB/wQE\\nAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwHwYDVR0jBBgwFoAU\\nKIEmBdE0Gj/Bcw+7k88VHD8Dv38wYwYIKwYBBQUHAQEEVzBVMCEGCCsGAQUFBzAB\\nhhVodHRwOi8vb2NzcC5teXNzbC5jb20wMAYIKwYBBQUHMAKGJGh0dHA6Ly9jYS5t\\neXNzbC5jb20vbXlzc2x0ZXN0cnNhLmNydDAaBgNVHREEEzARgg90ZXN0Y2VydC5x\\ncS5jb20wDQYJKoZIhvcNAQELBQADggEBAKN/GA4FZGLzaEGs+8zbNkCbNaq/zs+V\\nDBmdWF124wn+xke7PMN6PByWHBWFg7vJIfN7J7O4cFuShz23qcKj3qnDbUNKuK0X\\nTjGOJcFqSaHydlrPnmVSKEIM2J5O+VyR2gjuP+J6GvJavr0hFLRXB5gVmDfpxX6P\\nQX3SOilfn1rM0L7x1UG9t3t1USTumAgOuV/rflhOZs88nDjllNiR7RzsczBvrjes\\n5dP+mBcCVy9PUcVHJqNNNvbR4Vly8Ki+3c10jkxR9fVMnNpi2uI/eya1SLKGM6yr\\nHxEBkAxaANt8vEcCyiEBraqUvf9gBydvRKG0V6gVeeZ9+5WJd84WBzY=\\n-----END CERTIFICATE-----\\n",
					"key":            "-----BEGIN RSA PRIVATE KEY-----\\nMIIEogIBAAKCAQEAqjAfE7/ToSl6okR9YhQSz/vw+NvQUU4i/p+7HKqCoLC+l+rY\\n4HzU7m7fvgfkPZtTItgaiIKURo/Y+UmIaN+JNvnaTPZTXRgkfdccdTvYOeYzHn77\\nAzFdsjCcystdpXdJ/wrtfABa5XGNrFMeXYpf9Gxp+1A56modo8wiWAS48rysfzS8\\nvHUwcYuog5Nxttg5Kaqs2w1eUxiPLsXp+mN5CIhP46m1U5cisq9J1utAbKaHeRO+\\npWFnAu3e4AK1utaxhGdYCDMX3xwlNWFIgpNepFpdrs+H8btwsoPUkbwdYM+lFkGR\\nI8XZ4dMwQEI6p6iohthk6o9GQj+otOEqkGHyHQIDAQABAoIBAEA02XSk+V2i/X48\\noqUe95356gapP2V9Ohyf/IKrHY8sPyunUV0YG2k88TKLXaOUdv/9Ub1QrkoUuQIL\\nqOgP9X+FMcO5ZugHVLUZM8ZS5pepbn3B4EdrF3NDfdPQd6sWXxdWcxRGOgS4G3/4\\n98rIirz3LeC/eqoikL4cJJTaa0BAKpmxy2A5zM+2mtWjte8kLFmll1dEudVgwjI7\\nGISpw0vQQ6m6Odeu8rShW3EfXK9+VVQcOyKOEoNH9UlhugtkJvBDX/HZQylebHDX\\nKlT5Zs3uzpTsBKUC/fdv6S2vMl5YTvTbXPx/v5IOXhFsP/dMZ0Pw9/+feh5ILyo8\\ni0sh1o0CgYEAx7tuqP32bFajabqPzY4xz1LYo8rb15orilbNs3saxpAmo3h+CkpC\\nHCyyeBcnW9jRCSDGkFY+hN9zdyLMVp+2oURY3y/8iH5AQPfHHxJ8pcJNbpcL0xjk\\n6YfcEFwGbjz/c8sNYLds4Vj+/Ztn1l36LGq3hxPc9mz56J/T8+9JtUcCgYEA2iH6\\nzwu7wAU2JGvyolvWrZ4aOQ4UoKMX/fKzR32NTW5VG/te0h+TD93xeIEAfwcqcE+R\\nPq5L94ashcWxLLO1Ykeh4NY2cjHOB5iwgPfcs4bOQar9mXs7cpbIRXOHZtVC+QIf\\nbmSIeUmD9teQK8fsOsXzVKys7jVvsHbWN+f533sCgYBp2y7lJeRquuhU6um4Ofqw\\nNOpYtPUbKbyVzzeqPj6Mqm2wCfZTIdQz9oSIHU1g4mK3gcV7ThpIdd1OcQT3jCJR\\nClZHw7kF5lPLmwbPsw6ZE4fSav97XCBGnRjHGt7VokKJbj6i2dQ9AtvyMWExPhGP\\nZOfsEVz0xgEVh0/gYxwflQKBgEaKupt+hechSFMa/cp3vMNE3+IXpaAVXkrn1AaI\\nmcuypJ95+T4mq95zvgVkYo6x+I19jdEheLvBt8VDG3sDWuM/myAk5lKjvKdEP8BV\\nZ/A+5jDiZUTfT2hNTtL2+5DL4u64OwXuSRxAJAcNyzf9XW3cWkbF2N7oNQhyRfCq\\nempFAoGAAMjporbd5G7n1FjgAzZvlIJ2hIXUtCwIpIgRobEZYY9VQ9Y3OB9OAd0O\\nVDJbJqlhx0UUsDdyelW+L6NNADhF1Sq//VzL1uCPdKsM6Pn/UoxCbyswzycC+vo2\\n7co3GwupojSToXvidAYTYKupGoU5Qj227c7YGsU69qhfFq5T+Uc=\\n-----END RSA PRIVATE KEY-----\\n",
					"cert_region":    "ap-southeast-1",
					"custom_headers": "{\\\"3444\\\":\\\"5555\\\",\\\"666\\\":\\\"$ReqClientPort\\\",\\\"77777\\\":\\\"$ReqClientIP\\\"}",
					"ai_template":    "level30",
					"ai_mode":        "defense",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "0",
						"proxy_types.#":  "2",
						"real_servers.#": "1",
						"domain":         "testcert.qq.com",
						"instance_ids.#": "1",
						"https_ext":      CHECKSET,
						"cert_name":      CHECKSET,
						"cert":           "-----BEGIN CERTIFICATE-----\nMIID5TCCAs2gAwIBAgIRAJ8XVm8/e0sNiF7SQKrNVpUwDQYJKoZIhvcNAQELBQAw\nXjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU\nZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w\nHhcNMjUwOTI5MDY1NDI3WhcNMjYwOTI5MDY1NDI3WjAnMQswCQYDVQQGEwJDTjEY\nMBYGA1UEAxMPdGVzdGNlcnQucXEuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A\nMIIBCgKCAQEAqjAfE7/ToSl6okR9YhQSz/vw+NvQUU4i/p+7HKqCoLC+l+rY4HzU\n7m7fvgfkPZtTItgaiIKURo/Y+UmIaN+JNvnaTPZTXRgkfdccdTvYOeYzHn77AzFd\nsjCcystdpXdJ/wrtfABa5XGNrFMeXYpf9Gxp+1A56modo8wiWAS48rysfzS8vHUw\ncYuog5Nxttg5Kaqs2w1eUxiPLsXp+mN5CIhP46m1U5cisq9J1utAbKaHeRO+pWFn\nAu3e4AK1utaxhGdYCDMX3xwlNWFIgpNepFpdrs+H8btwsoPUkbwdYM+lFkGRI8XZ\n4dMwQEI6p6iohthk6o9GQj+otOEqkGHyHQIDAQABo4HUMIHRMA4GA1UdDwEB/wQE\nAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwHwYDVR0jBBgwFoAU\nKIEmBdE0Gj/Bcw+7k88VHD8Dv38wYwYIKwYBBQUHAQEEVzBVMCEGCCsGAQUFBzAB\nhhVodHRwOi8vb2NzcC5teXNzbC5jb20wMAYIKwYBBQUHMAKGJGh0dHA6Ly9jYS5t\neXNzbC5jb20vbXlzc2x0ZXN0cnNhLmNydDAaBgNVHREEEzARgg90ZXN0Y2VydC5x\ncS5jb20wDQYJKoZIhvcNAQELBQADggEBAKN/GA4FZGLzaEGs+8zbNkCbNaq/zs+V\nDBmdWF124wn+xke7PMN6PByWHBWFg7vJIfN7J7O4cFuShz23qcKj3qnDbUNKuK0X\nTjGOJcFqSaHydlrPnmVSKEIM2J5O+VyR2gjuP+J6GvJavr0hFLRXB5gVmDfpxX6P\nQX3SOilfn1rM0L7x1UG9t3t1USTumAgOuV/rflhOZs88nDjllNiR7RzsczBvrjes\n5dP+mBcCVy9PUcVHJqNNNvbR4Vly8Ki+3c10jkxR9fVMnNpi2uI/eya1SLKGM6yr\nHxEBkAxaANt8vEcCyiEBraqUvf9gBydvRKG0V6gVeeZ9+5WJd84WBzY=\n-----END CERTIFICATE-----\n",
						"key":            "-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEAqjAfE7/ToSl6okR9YhQSz/vw+NvQUU4i/p+7HKqCoLC+l+rY\n4HzU7m7fvgfkPZtTItgaiIKURo/Y+UmIaN+JNvnaTPZTXRgkfdccdTvYOeYzHn77\nAzFdsjCcystdpXdJ/wrtfABa5XGNrFMeXYpf9Gxp+1A56modo8wiWAS48rysfzS8\nvHUwcYuog5Nxttg5Kaqs2w1eUxiPLsXp+mN5CIhP46m1U5cisq9J1utAbKaHeRO+\npWFnAu3e4AK1utaxhGdYCDMX3xwlNWFIgpNepFpdrs+H8btwsoPUkbwdYM+lFkGR\nI8XZ4dMwQEI6p6iohthk6o9GQj+otOEqkGHyHQIDAQABAoIBAEA02XSk+V2i/X48\noqUe95356gapP2V9Ohyf/IKrHY8sPyunUV0YG2k88TKLXaOUdv/9Ub1QrkoUuQIL\nqOgP9X+FMcO5ZugHVLUZM8ZS5pepbn3B4EdrF3NDfdPQd6sWXxdWcxRGOgS4G3/4\n98rIirz3LeC/eqoikL4cJJTaa0BAKpmxy2A5zM+2mtWjte8kLFmll1dEudVgwjI7\nGISpw0vQQ6m6Odeu8rShW3EfXK9+VVQcOyKOEoNH9UlhugtkJvBDX/HZQylebHDX\nKlT5Zs3uzpTsBKUC/fdv6S2vMl5YTvTbXPx/v5IOXhFsP/dMZ0Pw9/+feh5ILyo8\ni0sh1o0CgYEAx7tuqP32bFajabqPzY4xz1LYo8rb15orilbNs3saxpAmo3h+CkpC\nHCyyeBcnW9jRCSDGkFY+hN9zdyLMVp+2oURY3y/8iH5AQPfHHxJ8pcJNbpcL0xjk\n6YfcEFwGbjz/c8sNYLds4Vj+/Ztn1l36LGq3hxPc9mz56J/T8+9JtUcCgYEA2iH6\nzwu7wAU2JGvyolvWrZ4aOQ4UoKMX/fKzR32NTW5VG/te0h+TD93xeIEAfwcqcE+R\nPq5L94ashcWxLLO1Ykeh4NY2cjHOB5iwgPfcs4bOQar9mXs7cpbIRXOHZtVC+QIf\nbmSIeUmD9teQK8fsOsXzVKys7jVvsHbWN+f533sCgYBp2y7lJeRquuhU6um4Ofqw\nNOpYtPUbKbyVzzeqPj6Mqm2wCfZTIdQz9oSIHU1g4mK3gcV7ThpIdd1OcQT3jCJR\nClZHw7kF5lPLmwbPsw6ZE4fSav97XCBGnRjHGt7VokKJbj6i2dQ9AtvyMWExPhGP\nZOfsEVz0xgEVh0/gYxwflQKBgEaKupt+hechSFMa/cp3vMNE3+IXpaAVXkrn1AaI\nmcuypJ95+T4mq95zvgVkYo6x+I19jdEheLvBt8VDG3sDWuM/myAk5lKjvKdEP8BV\nZ/A+5jDiZUTfT2hNTtL2+5DL4u64OwXuSRxAJAcNyzf9XW3cWkbF2N7oNQhyRfCq\nempFAoGAAMjporbd5G7n1FjgAzZvlIJ2hIXUtCwIpIgRobEZYY9VQ9Y3OB9OAd0O\nVDJbJqlhx0UUsDdyelW+L6NNADhF1Sq//VzL1uCPdKsM6Pn/UoxCbyswzycC+vo2\n7co3GwupojSToXvidAYTYKupGoU5Qj227c7YGsU69qhfFq5T+Uc=\n-----END RSA PRIVATE KEY-----\n",
						"cert_region":    "ap-southeast-1",
						"custom_headers": CHECKSET,
						"ai_template":    "level30",
						"ai_mode":        "defense",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"real_servers": []string{
						"1.1.1.1", "2.2.2.2", "3.3.3.3"},
					"cert_name":      "97",
					"cert":           "-----BEGIN CERTIFICATE-----\\nMIID5DCCAsygAwIBAgIQXWOU7ILVRC20+c/4dgiiJDANBgkqhkiG9w0BAQsFADBe\\nMQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl\\nc3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe\\nFw0yNTA5MjkwNjU1MThaFw0yNjA5MjkwNjU1MThaMCcxCzAJBgNVBAYTAkNOMRgw\\nFgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw\\nggEKAoIBAQDe24eCDI+IDalD4pYU4NuvS5/S40Hwu+B3yKB9KN1HL2pTEH7kOYG/\\nRQHvkDUpsCyx0H4inpzVWu36qXAEXj9QLCNFN3+SdlBZBKaaWb76xExXPtlmyURV\\nXXlmMSSx7MeqS2euOmeIBd7WqsPX4kIok71E8tnMaLqlNEnx9SmksZukLQYusZC5\\nuECr+kgEsCDFf4JM1ZYADNL1csc1HzFPmUzZyOl99ZxAdrVNuDhc/SKdpHb80FhV\\nRBv72WW1JDYYbkP9dZALUezR66uF0/+bPzmxi2QlSu8q4FfgCSCRV2GLRLNpxMTw\\nzOrwXen4gQCzm1jTpWsb1K/n1CPHfe3TAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD\\nAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo\\ngSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG\\nFWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15\\nc3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx\\nLmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAqlpc+4pcaKbUdSdWCIC2vQx40D/51AhX\\nMQk1bKOANo/2PGbWLcYPJ0QLiE/5umKz0cRDEcq4oVjdmJ1VUvUC986/cdnFlp+l\\nnbujrCby3dJFSsG/ZWcVGfcr1Ioy3LzcEnNto2SEUrj6lMCuyCI6zHeSHAGokxAp\\n8RpbhYooisAiDj74g7Xx2/vJgOMFw6i6KhvtHYwct0JpNNsdE5EHgAJ3FWLOqKQs\\n2D/YC7I72ae3X9IGHPJjzMiQiBAfh1BKZBYNeo3UVZLmidsmhoN9QLnKCTyiAciF\\nRnodqdg8gSb4j0WmIoAua5nWjhz9V6N5rS2FDBmyaeSep0XY7SWMJQ==\\n-----END CERTIFICATE-----\\n",
					"key":            "-----BEGIN RSA PRIVATE KEY-----\\nMIIEpAIBAAKCAQEA3tuHggyPiA2pQ+KWFODbr0uf0uNB8Lvgd8igfSjdRy9qUxB+\\n5DmBv0UB75A1KbAssdB+Ip6c1Vrt+qlwBF4/UCwjRTd/knZQWQSmmlm++sRMVz7Z\\nZslEVV15ZjEksezHqktnrjpniAXe1qrD1+JCKJO9RPLZzGi6pTRJ8fUppLGbpC0G\\nLrGQubhAq/pIBLAgxX+CTNWWAAzS9XLHNR8xT5lM2cjpffWcQHa1Tbg4XP0inaR2\\n/NBYVUQb+9lltSQ2GG5D/XWQC1Hs0eurhdP/mz85sYtkJUrvKuBX4AkgkVdhi0Sz\\nacTE8Mzq8F3p+IEAs5tY06VrG9Sv59Qjx33t0wIDAQABAoIBAAoLR3HrVwFKMzB8\\nov7/CkMZbgV3MJ7E9F861CoGxy2+uVSdD/+Th+8y7G9pWQUvZ4H+ZCfOtkGuSAgU\\nqbztO6L8PthqoaJuhhiGIkeM5RPCitUhhAUjEH+mJ2xB8sFv2M2YC9qAGIJxfLfl\\n0BgaML223S4O6sxQzlQgXFKHX0mtbXGVVA0DSOwhTy5x+qLhp/Bx4TD6Qfvl9Pt8\\n3Ip8hRRtepV03h2W4TCZuU0HdLTgFcO7WXgl9HLNuubrf9EhzhRfBMGVGTfYzwMW\\npQHYXBRqFyq1+KFFhsiqT6K5UzyxWw3Rpwy03tsaIgukg/HRDrB2icjODDWpCSyD\\nM6KAYaECgYEA6E5+HqUJcqvOYXkiB6FwO6uKvCEh2GGdmcdjh8xi0jF4S78XG4RF\\nXx709z7kEY5eBRfb2Nnby1JzDkTG/NgDStcVYB/qQHK73EDb/jAsniRxA8nLAZfD\\ni0ytPpDX5dZok1gIbNx0koRQhlBGudHr5PoY7N1QQcvcmnVQWdW/EzcCgYEA9ZZS\\nckFbVTo+ePLbB8EgJuNR48mKLefWT0q5UUuFyC3ncRNpiJ/DEMBIRqjjWq6CdWXK\\n12Hz56CdpXS/HgZXnU+QqGHQ5l7Pz9hVjM2WIJ3w6kJVw1fBO1B11XfRxXnN5IJE\\na/ATOAXIRDrik+Rg3TCpZusBeTkg8do9QENQQEUCgYEAn/llup62OeR8U/2B5LVU\\nv5KrEFDUqNjYGg0Hyn2CU/NDPw5R0F4vE4kS8qy5jCl5L5K1j8i/Jm4Z02qjiX0M\\nD168VpzDySv5mHyFwq7UGvdHaG9vQCKNw4DDEQHX22viSg1mh+js0fUSKtxfSBl6\\nlA1yWrMxUI4d1bQR6DtcwNMCgYA8MWbtyCUZo5fyTxvuL6Cwx2Cn4xryG3PEpXz8\\ndvVIVi/24BoquXW3IlnUr8phzIn/Oj7YQZLlf9GD3zSEqGtLDFhZXPg1rqFiwRRe\\n2Xjlb7C/yhh5M4YSAquO8bpBm5QiYOdiSUp8nbYzOveT5hLzw9yRdCI9UwpHHQWM\\nPbNqvQKBgQDaxQz3ma5o348nRWTJrgLaipYveqFQLtosIMxDL1p8SHfiRZmSvUAC\\n0JJa7gD3LWoKPfTtMhR7Tfdkcw6x0WEzVNYy5xiWce9yhpD0cN+RoNl0xjxaAcXN\\n+hRL+9t7LIJr5zZoFL4Fj3h9FtSnhPLYk13AUa9WnY5Ih7SIaFD9hQ==\\n-----END RSA PRIVATE KEY-----\\n",
					"cert_region":    "cn-hangzhou",
					"custom_headers": "{\\\"22\\\":\\\"$ReqClientIP\\\",\\\"77\\\":\\\"88\\\",\\\"99\\\":\\\"$ReqClientPort\\\"}",
					"ocsp_enabled":   "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"real_servers.#": "3",
						"cert_name":      CHECKSET,
						"cert":           "-----BEGIN CERTIFICATE-----\nMIID5DCCAsygAwIBAgIQXWOU7ILVRC20+c/4dgiiJDANBgkqhkiG9w0BAQsFADBe\nMQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl\nc3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe\nFw0yNTA5MjkwNjU1MThaFw0yNjA5MjkwNjU1MThaMCcxCzAJBgNVBAYTAkNOMRgw\nFgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw\nggEKAoIBAQDe24eCDI+IDalD4pYU4NuvS5/S40Hwu+B3yKB9KN1HL2pTEH7kOYG/\nRQHvkDUpsCyx0H4inpzVWu36qXAEXj9QLCNFN3+SdlBZBKaaWb76xExXPtlmyURV\nXXlmMSSx7MeqS2euOmeIBd7WqsPX4kIok71E8tnMaLqlNEnx9SmksZukLQYusZC5\nuECr+kgEsCDFf4JM1ZYADNL1csc1HzFPmUzZyOl99ZxAdrVNuDhc/SKdpHb80FhV\nRBv72WW1JDYYbkP9dZALUezR66uF0/+bPzmxi2QlSu8q4FfgCSCRV2GLRLNpxMTw\nzOrwXen4gQCzm1jTpWsb1K/n1CPHfe3TAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD\nAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo\ngSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG\nFWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15\nc3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx\nLmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAqlpc+4pcaKbUdSdWCIC2vQx40D/51AhX\nMQk1bKOANo/2PGbWLcYPJ0QLiE/5umKz0cRDEcq4oVjdmJ1VUvUC986/cdnFlp+l\nnbujrCby3dJFSsG/ZWcVGfcr1Ioy3LzcEnNto2SEUrj6lMCuyCI6zHeSHAGokxAp\n8RpbhYooisAiDj74g7Xx2/vJgOMFw6i6KhvtHYwct0JpNNsdE5EHgAJ3FWLOqKQs\n2D/YC7I72ae3X9IGHPJjzMiQiBAfh1BKZBYNeo3UVZLmidsmhoN9QLnKCTyiAciF\nRnodqdg8gSb4j0WmIoAua5nWjhz9V6N5rS2FDBmyaeSep0XY7SWMJQ==\n-----END CERTIFICATE-----\n",
						"key":            "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA3tuHggyPiA2pQ+KWFODbr0uf0uNB8Lvgd8igfSjdRy9qUxB+\n5DmBv0UB75A1KbAssdB+Ip6c1Vrt+qlwBF4/UCwjRTd/knZQWQSmmlm++sRMVz7Z\nZslEVV15ZjEksezHqktnrjpniAXe1qrD1+JCKJO9RPLZzGi6pTRJ8fUppLGbpC0G\nLrGQubhAq/pIBLAgxX+CTNWWAAzS9XLHNR8xT5lM2cjpffWcQHa1Tbg4XP0inaR2\n/NBYVUQb+9lltSQ2GG5D/XWQC1Hs0eurhdP/mz85sYtkJUrvKuBX4AkgkVdhi0Sz\nacTE8Mzq8F3p+IEAs5tY06VrG9Sv59Qjx33t0wIDAQABAoIBAAoLR3HrVwFKMzB8\nov7/CkMZbgV3MJ7E9F861CoGxy2+uVSdD/+Th+8y7G9pWQUvZ4H+ZCfOtkGuSAgU\nqbztO6L8PthqoaJuhhiGIkeM5RPCitUhhAUjEH+mJ2xB8sFv2M2YC9qAGIJxfLfl\n0BgaML223S4O6sxQzlQgXFKHX0mtbXGVVA0DSOwhTy5x+qLhp/Bx4TD6Qfvl9Pt8\n3Ip8hRRtepV03h2W4TCZuU0HdLTgFcO7WXgl9HLNuubrf9EhzhRfBMGVGTfYzwMW\npQHYXBRqFyq1+KFFhsiqT6K5UzyxWw3Rpwy03tsaIgukg/HRDrB2icjODDWpCSyD\nM6KAYaECgYEA6E5+HqUJcqvOYXkiB6FwO6uKvCEh2GGdmcdjh8xi0jF4S78XG4RF\nXx709z7kEY5eBRfb2Nnby1JzDkTG/NgDStcVYB/qQHK73EDb/jAsniRxA8nLAZfD\ni0ytPpDX5dZok1gIbNx0koRQhlBGudHr5PoY7N1QQcvcmnVQWdW/EzcCgYEA9ZZS\nckFbVTo+ePLbB8EgJuNR48mKLefWT0q5UUuFyC3ncRNpiJ/DEMBIRqjjWq6CdWXK\n12Hz56CdpXS/HgZXnU+QqGHQ5l7Pz9hVjM2WIJ3w6kJVw1fBO1B11XfRxXnN5IJE\na/ATOAXIRDrik+Rg3TCpZusBeTkg8do9QENQQEUCgYEAn/llup62OeR8U/2B5LVU\nv5KrEFDUqNjYGg0Hyn2CU/NDPw5R0F4vE4kS8qy5jCl5L5K1j8i/Jm4Z02qjiX0M\nD168VpzDySv5mHyFwq7UGvdHaG9vQCKNw4DDEQHX22viSg1mh+js0fUSKtxfSBl6\nlA1yWrMxUI4d1bQR6DtcwNMCgYA8MWbtyCUZo5fyTxvuL6Cwx2Cn4xryG3PEpXz8\ndvVIVi/24BoquXW3IlnUr8phzIn/Oj7YQZLlf9GD3zSEqGtLDFhZXPg1rqFiwRRe\n2Xjlb7C/yhh5M4YSAquO8bpBm5QiYOdiSUp8nbYzOveT5hLzw9yRdCI9UwpHHQWM\nPbNqvQKBgQDaxQz3ma5o348nRWTJrgLaipYveqFQLtosIMxDL1p8SHfiRZmSvUAC\n0JJa7gD3LWoKPfTtMhR7Tfdkcw6x0WEzVNYy5xiWce9yhpD0cN+RoNl0xjxaAcXN\n+hRL+9t7LIJr5zZoFL4Fj3h9FtSnhPLYk13AUa9WnY5Ih7SIaFD9hQ==\n-----END RSA PRIVATE KEY-----\n",
						"cert_region":    "cn-hangzhou",
						"custom_headers": CHECKSET,
						"ocsp_enabled":   "true",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap12079 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence12079(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_ddoscoo_instance" "defaultdPWcCl" {
  normal_bandwidth = "100"
  product_plan     = "0"
  product_type     = "ddosDip"
  period           = "1"
  normal_qps       = "500"
  port_count       = "5"
  function_version = "0"
  domain_count     = "10"
  bandwidth_mode   = "2"
  name             = "测试"
}


`, name)
}

// Case 域名资源测试 3256
func TestAccAliCloudDdosCooDomainResource_basic3256(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap3256)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence3256)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type": "0",
					"real_servers": []string{
						"1.1.1.1"},
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
					},
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaulttV5f6b.id}"},
					"domain":    "ldtest123.qq.com",
					"https_ext": "{\\\"Https2http\\\":1,\\\"Http2\\\":1,\\\"Http2https\\\":0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "0",
						"real_servers.#": "1",
						"proxy_types.#":  "1",
						"instance_ids.#": "1",
						"domain":         "ldtest123.qq.com",
						"https_ext":      CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type": "1",
					"real_servers": []string{
						"qq.com"},
					"https_ext": "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "1",
						"real_servers.#": "1",
						"https_ext":      CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap3256 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence3256(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_ddoscoo_instance" "defaulttV5f6b" {
  product_type = "ddoscoo"
  period       = "1"
}


`, name)
}

// Case 测试选择证书-李铎_测试 7808
func TestAccAliCloudDdosCooDomainResource_basic7808(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap7808)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence7808)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type":      "0",
					"ocsp_enabled": "0",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "websocket",
						},
					},
					"real_servers": []string{
						"1.1.1.1", "2.2.2.2", "3.3.3.3"},
					"domain": "testld.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultSJe7n8.id}", "${alicloud_ddoscoo_instance.default6lyurZ.id}", "${alicloud_ddoscoo_instance.defaultTTvY0D.id}"},
					"https_ext": "{\\\"Https2http\\\":1,\\\"Http2\\\":1,\\\"Http2https\\\":0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "0",
						"ocsp_enabled":   "0",
						"proxy_types.#":  "3",
						"real_servers.#": "3",
						"domain":         "testld.qq.com",
						"instance_ids.#": "3",
						"https_ext":      CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type":      "1",
					"ocsp_enabled": "1",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80", "8080"},
							"proxy_type": "http",
						},
					},
					"real_servers": []string{
						"1.qq.com"},
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultTTvY0D.id}"},
					"https_ext":       "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
					"cert_identifier": "${alicloud_ssl_certificates_service_certificate.default3MYZEt.id}",
					"cert_region":     "cn-hangzhou",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":         "1",
						"ocsp_enabled":    "1",
						"proxy_types.#":   "1",
						"real_servers.#":  "1",
						"instance_ids.#":  "1",
						"https_ext":       CHECKSET,
						"cert_identifier": CHECKSET,
						"cert_region":     "cn-hangzhou",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap7808 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence7808(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_ddoscoo_instance" "defaultTTvY0D" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "0"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ddoscoo_instance" "defaultSJe7n8" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test2"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "1"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ddoscoo_instance" "default6lyurZ" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test2"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "1"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ssl_certificates_service_certificate" "default3MYZEt" {
  cert             = <<EOF
-----BEGIN CERTIFICATE----- MIID4TCCAsmgAwIBAgIRANZGvLwT8kuWpPlZ/Aj+uPgwDQYJKoZIhvcNAQELBQAw XjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU ZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w HhcNMjQwODIzMDk0NzA0WhcNMjUwODIzMDk0NzA0WjAlMQswCQYDVQQGEwJDTjEW MBQGA1UEAxMNdGVzdGxkLnFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC AQoCggEBALmZY2geTFi+50gAVyDQH9Y5sTv8LLX6+MET1l3larzjX1M0Az9ZEIc0 TNrAp8mtJRlpQCzyDPZg88AwSdEwqSOSsnGzfS2DUcPJmdn2a2n5PLvWE28qPuSf 6fl3IhNiPzLYR51+7ccJKEQRhfOK2usmJo6oTG/0Lhh4BRH5owcclKv6n3YHaBVj JNigiq1/tlqU46toZvotPOORjpy21kJPZioHqOVCDO4zreMy2xuIiYtpSSmXwkEO zcQQ3K8sbRx9ED8SCdb229h7ioTug02YBXs0YOQZ024HFaIF8Nz1M+mdHy1jCbLd yJoT/jzE4RdldZKZJFaSKV1c7EYlzhkCAwEAAaOB0jCBzzAOBgNVHQ8BAf8EBAMC BaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMB8GA1UdIwQYMBaAFCiB JgXRNBo/wXMPu5PPFRw/A79/MGMGCCsGAQUFBwEBBFcwVTAhBggrBgEFBQcwAYYV aHR0cDovL29jc3AubXlzc2wuY29tMDAGCCsGAQUFBzAChiRodHRwOi8vY2EubXlz c2wuY29tL215c3NsdGVzdHJzYS5jcnQwGAYDVR0RBBEwD4INdGVzdGxkLnFxLmNv bTANBgkqhkiG9w0BAQsFAAOCAQEAnPJl1GrePDIulWfsETPbGnrZv3j3ZRXuou0o K32X/nydS/i/j+AUzKSyezmnR1edkgY1hbGaza702SLQJuGh2IqJvAFyifwV/CZ5 cpJIi5G7kWTBjZo9NgVnDMhR8y5DCKE8BhiUBwcSvKKC8se2yWHm1fk9pRxG0Mc6 0fstl40jtR5XZYsW1GhX4fzwrWuBodPKticgXPn2e24ec+4rVrziu5R7D77AzJjG Y/wzNYvAUWEzEya7Ve53nhu+WpIuIQn0ux8nPDioFdOjckn4jK3ePYdS2mWT6EBU BC74GYiBNDz0QgHADq1VTExeLzC0tw9PPdWl0WfoTgCCKLz0yA== -----END CERTIFICATE-----
EOF
  certificate_name = "501"
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuZljaB5MWL7nSABXINAf1jmxO/wstfr4wRPWXeVqvONfUzQD P1kQhzRM2sCnya0lGWlALPIM9mDzwDBJ0TCpI5KycbN9LYNRw8mZ2fZrafk8u9YT byo+5J/p+XciE2I/MthHnX7txwkoRBGF84ra6yYmjqhMb/QuGHgFEfmjBxyUq/qf dgdoFWMk2KCKrX+2WpTjq2hm+i0845GOnLbWQk9mKgeo5UIM7jOt4zLbG4iJi2lJ KZfCQQ7NxBDcryxtHH0QPxIJ1vbb2HuKhO6DTZgFezRg5BnTbgcVogXw3PUz6Z0f LWMJst3ImhP+PMThF2V1kpkkVpIpXVzsRiXOGQIDAQABAoIBAQCBPiw4A+k8X2vk +r+xjNyurCwcTmXAL81rfmnnputmL5tg8DZWtanJzQS7zC7LRPQxttZGtiOKqkbz DW1J6+3MZMo4XToNKIYWpduqKWvxNusxDkkoPy3evPEMlAY5o0/JE00DgrEHyfut MtqplocN+tocu1vHFi3HQkSdmM4LE46ZfFu5w1FRbNI1Gqjj/cwlF/T93V3qMap1 WfsJjhMIX9LjBq3y9GAfAtAw7JYwkztr2AhYzCsK25wAj72zFY6FJTZ8LklfS41q DrVtdjMx42IonDQtkzrqzfYlXdzzhZzuQHxn+qJODseoU8oDG9j3eKhp1dKgqLfx tv1o3km1AoGBAOqXGEw2w94uVchCjuTum3XFYieEla0IUbHJCaWKU/hoSbht7j23 K7tA9//epBuRLGtYE0sPBK6i31mQT216muspO1g3pwGJhPy8VSpsJ1GQhB2G2UNz kZlRK+2/gx35TdTi9x0C6UWk5XkhgWO3R35BlEnuV7EOyJunobiUcoObAoGBAMqJ reuSbJajNGfeBzPel7F62ZDufC85hWaGKzeXIk3DkXcsEpeR5ogMGHCZ3dGf4Yz/ pcfjnCMIWjc+MkA4ppFd4432FJkxNQQP0z7njpXW3e5tionsMN+UwPzi4wZPKufK osjw43JpBzpHGxG3ynLgZg7bSfrQZhPDTQw84nJbAoGBAKpy8mKeAB71R7rkMXNB s48Uxca03RQGUWV+DxZKtcxt6fKpXUtWRd4ezJMLL+4fw0iTjCEjXmGNUf9/jVac mOd44/erKBtD0m7YYIEcaE0pVfUmP8J0vDvL8MEkP56Nv/GIn8hijx/dOiaTI7JS Pw4LlDVLikfJ2BTQ7f5xTes1AoGBAJ7h4HiDFgIZp1uvtfC/tjn5CEGEhBC7y+VA bRify747I5rcDP2v66tf6bAzU+pExLhKN++Vov9sZvEdLmhoyGoSwBa2KzR9gHxe ObYICjeLJfALKHnHuhM6ayY2iieB5UOOF6MQLSysLYpPC3IbvonddNJEvkUuRFVO iNuHy5AvAoGAPqbzSNA05gf85zRO/JZAmZuWXG3o3pVougbkc211p+ynMpS+/bMb c/nR36kOE551lFjIoAjoeIs16Wbq+00u9GlcQmyAdFpfaFCNHa3dayJwKJMW9Nia fKbiiiOAQE8s2v8Paa+b00GspeWLow4u0G5lBVau4JjEVnl6ivLXlzY= -----END RSA PRIVATE KEY-----
EOF
}


`, name)
}

// Case 手动上传证书_李铎 7629
func TestAccAliCloudDdosCooDomainResource_basic7629(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap7629)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence7629)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"ap-southeast-1"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type":      "0",
					"ocsp_enabled": "0",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
					},
					"real_servers": []string{
						"1.1.1.1"},
					"domain": "testcert.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultaRfzZ9.id}"},
					"https_ext":   "{\\\"Https2http\\\":1,\\\"Http2\\\":1,\\\"Http2https\\\":0}",
					"cert_name":   "201",
					"cert":        "-----BEGIN CERTIFICATE----- MIID5DCCAsygAwIBAgIQWen3GebvT0GcE/a1MJVFgjANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MjMwOTE1MjNaFw0yNTA4MjMwOTE1MjNaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDF/H2/Oas8trhkHs7B4B8mN00eup/5Tqar2QO4Hm499GJECKc3eMiC v+aAlW74Iymb2Varnv+WMdFRVMQgpXesi3akvVp0QxecvcDliilkh4ddTK731Rd7 PaSK1JdQX1jdGGhVnhQz+cPNFBGZ3tMYGhUkgNfqa3UFucJcBuRub/Ircr+5Ob4D FxSglfTHi+/EFcp7vMAOztLD4zXmEz3NysDNP6NzN7SD72DwPp0nxyRjrBlHSOVg szB/bFasQdAhZGeo64MvSb+SivdWEMhHkwKA5MhhYOkDeNPPSmlxbw0Z3nOyeMmI YkaxzhpO5DZN382duTQmiQ+Yg60OfL3NAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAKUCvcbUCJRFbVowd1YorILivqRmS6ztR 9vLdj4YZBWxmmQrgkDlkl78r/rXlJbqHunSh2Wbag7y+GaQQwg8xcL4Z3KrKj4zg aHEP1DyFiaxMTEuC/L2RgSX3xlXcf6fQ46D3y3Ja3iHFQnjx6npNaSZ2bSEULvJg IjPiJ/nbid1TYR5vtg8vtwDQfY7/+q8/3DWKcQq+SGcd9dDS6u9vulNdlW8e14bS CKDEzS/axjoICl9JagASLcElWIit/eD5zGIKzTPC9mEXiX/J/gUr70y9GiE7Ue++ 4nFSGQOwjMh/wO2HlmRfToeZ3g6rRCijibBHKHBmVym7NCai2voZJQ== -----END CERTIFICATE-----",
					"key":         "-----BEGIN RSA PRIVATE KEY----- MIIEowIBAAKCAQEAxfx9vzmrPLa4ZB7OweAfJjdNHrqf+U6mq9kDuB5uPfRiRAin N3jIgr/mgJVu+CMpm9lWq57/ljHRUVTEIKV3rIt2pL1adEMXnL3A5YopZIeHXUyu 99UXez2kitSXUF9Y3RhoVZ4UM/nDzRQRmd7TGBoVJIDX6mt1BbnCXAbkbm/yK3K/ uTm+AxcUoJX0x4vvxBXKe7zADs7Sw+M15hM9zcrAzT+jcze0g+9g8D6dJ8ckY6wZ R0jlYLMwf2xWrEHQIWRnqOuDL0m/kor3VhDIR5MCgOTIYWDpA3jTz0ppcW8NGd5z snjJiGJGsc4aTuQ2Td/Nnbk0JokPmIOtDny9zQIDAQABAoIBAQDD6fU4y8UhwCG4 mS+5c6D/PQvoU35Hwkd1l7pxcFNgpTqz3egyISgxEdny9WwoyQq8eJWmICEEK+nY VEv7jiFdMWhG3kTq9RUhejeuLEiHfQE7Fs2w2kFxJ29yHapZ0u/pYOSljFarlATo I2rDW1aB7BVt2L1P7+ONteKZFAzpJckft5ceRUzs5Jm1Cqt8OWO3Km+FBbCROv8M TevW44aoMwBGXuqs06FV1Z4dafglskjt2O38V4acZpH8Nc8j+nCONKL3OxwKY6HQ WfnbXnTLCF3IuMiy8ntrY8HYU6EABiCdr+Pl5HmhI2nmtSFTFbD4Gq70vgPL0P1m iULJGJ7hAoGBANQPrOGe9qHcBydvcBHE7qA9v1+IaTj03qzDTopTi/jxcd9pEkei skLyHNQ5yJT0QjTxB9iYRLfZccOGFyqz/Sdz6CwwTWBZeXOQ2AX7FPEcCnNr1TpF yMrgOY3H93KJISEVS6kYskByjK7XzXCp0KQNS2EeIhAXcqXxNmSwylvpAoGBAO8C PdZHd6aLLEZyVO1aZVHDxqmbhmGVoY9wZ+uwR4K2Hu/fjk0qlR9cYpw8+N675Wr9 E9Ff5/wjK2+/+uocQV9Zoap2vgrwX7GASuO5KYdCOBn6oUOSa+Ru+LgBNyUkXYES mM8eFC1QqfcSrETLAQqd2lmLcuaMq6jJtbBpvzhFAoGAYd8mNC9wtr1dE+dLuvfA BnbZJ1dG8QGa7/NoAVGT7X5JxwmwZR2C1oD1q0FMAOtGzzZbH60PMicKaWousQfH E/lbs2FLpOdGtX6pJQF/5dPCQwkGrVFd3bxk87nRy6vcfW9drxp10mbL5To2WAQY Bk8Ydic5I2IfCNVt/ETX8FkCgYA+OkAtVQgi9WM+qC/SaFGu2yETManoKFQbC3IT HB9SOeaOH49mKesPcjc+ZGWLYDJYC7IoNicpL2L0wnAqmdavY5/CyQ2rvW+8wCE/ bwsP6z6+DNIFzM6IeBgLmE1qPzCVFWlxq2wnbDQEXvk5I/2ObRDXdYYh3ogm9vV2 C+I8XQKBgHIquvifRVvWf1q9WFZLQXZMv1flPhNaLmR+2k6gNpJ8SeiOmBtE7gT6 Je+YOXEKvfr6jaaJwYHPi6IhWHs4fQbgdK4jei30sRL7c8QdKEuwRdHPimmGNAPb UapzHY7xq0Wk9enAnM/SXkjTAJEkrpiQiDuPZVi4sIYCOqb+Ovu5 -----END RSA PRIVATE KEY-----",
					"cert_region": "ap-southeast-1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "0",
						"ocsp_enabled":   "0",
						"proxy_types.#":  "2",
						"real_servers.#": "1",
						"domain":         "testcert.qq.com",
						"instance_ids.#": "1",
						"https_ext":      CHECKSET,
						"cert_name":      CHECKSET,
						"cert":           "-----BEGIN CERTIFICATE----- MIID5DCCAsygAwIBAgIQWen3GebvT0GcE/a1MJVFgjANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MjMwOTE1MjNaFw0yNTA4MjMwOTE1MjNaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDF/H2/Oas8trhkHs7B4B8mN00eup/5Tqar2QO4Hm499GJECKc3eMiC v+aAlW74Iymb2Varnv+WMdFRVMQgpXesi3akvVp0QxecvcDliilkh4ddTK731Rd7 PaSK1JdQX1jdGGhVnhQz+cPNFBGZ3tMYGhUkgNfqa3UFucJcBuRub/Ircr+5Ob4D FxSglfTHi+/EFcp7vMAOztLD4zXmEz3NysDNP6NzN7SD72DwPp0nxyRjrBlHSOVg szB/bFasQdAhZGeo64MvSb+SivdWEMhHkwKA5MhhYOkDeNPPSmlxbw0Z3nOyeMmI YkaxzhpO5DZN382duTQmiQ+Yg60OfL3NAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAKUCvcbUCJRFbVowd1YorILivqRmS6ztR 9vLdj4YZBWxmmQrgkDlkl78r/rXlJbqHunSh2Wbag7y+GaQQwg8xcL4Z3KrKj4zg aHEP1DyFiaxMTEuC/L2RgSX3xlXcf6fQ46D3y3Ja3iHFQnjx6npNaSZ2bSEULvJg IjPiJ/nbid1TYR5vtg8vtwDQfY7/+q8/3DWKcQq+SGcd9dDS6u9vulNdlW8e14bS CKDEzS/axjoICl9JagASLcElWIit/eD5zGIKzTPC9mEXiX/J/gUr70y9GiE7Ue++ 4nFSGQOwjMh/wO2HlmRfToeZ3g6rRCijibBHKHBmVym7NCai2voZJQ== -----END CERTIFICATE-----",
						"key":            "-----BEGIN RSA PRIVATE KEY----- MIIEowIBAAKCAQEAxfx9vzmrPLa4ZB7OweAfJjdNHrqf+U6mq9kDuB5uPfRiRAin N3jIgr/mgJVu+CMpm9lWq57/ljHRUVTEIKV3rIt2pL1adEMXnL3A5YopZIeHXUyu 99UXez2kitSXUF9Y3RhoVZ4UM/nDzRQRmd7TGBoVJIDX6mt1BbnCXAbkbm/yK3K/ uTm+AxcUoJX0x4vvxBXKe7zADs7Sw+M15hM9zcrAzT+jcze0g+9g8D6dJ8ckY6wZ R0jlYLMwf2xWrEHQIWRnqOuDL0m/kor3VhDIR5MCgOTIYWDpA3jTz0ppcW8NGd5z snjJiGJGsc4aTuQ2Td/Nnbk0JokPmIOtDny9zQIDAQABAoIBAQDD6fU4y8UhwCG4 mS+5c6D/PQvoU35Hwkd1l7pxcFNgpTqz3egyISgxEdny9WwoyQq8eJWmICEEK+nY VEv7jiFdMWhG3kTq9RUhejeuLEiHfQE7Fs2w2kFxJ29yHapZ0u/pYOSljFarlATo I2rDW1aB7BVt2L1P7+ONteKZFAzpJckft5ceRUzs5Jm1Cqt8OWO3Km+FBbCROv8M TevW44aoMwBGXuqs06FV1Z4dafglskjt2O38V4acZpH8Nc8j+nCONKL3OxwKY6HQ WfnbXnTLCF3IuMiy8ntrY8HYU6EABiCdr+Pl5HmhI2nmtSFTFbD4Gq70vgPL0P1m iULJGJ7hAoGBANQPrOGe9qHcBydvcBHE7qA9v1+IaTj03qzDTopTi/jxcd9pEkei skLyHNQ5yJT0QjTxB9iYRLfZccOGFyqz/Sdz6CwwTWBZeXOQ2AX7FPEcCnNr1TpF yMrgOY3H93KJISEVS6kYskByjK7XzXCp0KQNS2EeIhAXcqXxNmSwylvpAoGBAO8C PdZHd6aLLEZyVO1aZVHDxqmbhmGVoY9wZ+uwR4K2Hu/fjk0qlR9cYpw8+N675Wr9 E9Ff5/wjK2+/+uocQV9Zoap2vgrwX7GASuO5KYdCOBn6oUOSa+Ru+LgBNyUkXYES mM8eFC1QqfcSrETLAQqd2lmLcuaMq6jJtbBpvzhFAoGAYd8mNC9wtr1dE+dLuvfA BnbZJ1dG8QGa7/NoAVGT7X5JxwmwZR2C1oD1q0FMAOtGzzZbH60PMicKaWousQfH E/lbs2FLpOdGtX6pJQF/5dPCQwkGrVFd3bxk87nRy6vcfW9drxp10mbL5To2WAQY Bk8Ydic5I2IfCNVt/ETX8FkCgYA+OkAtVQgi9WM+qC/SaFGu2yETManoKFQbC3IT HB9SOeaOH49mKesPcjc+ZGWLYDJYC7IoNicpL2L0wnAqmdavY5/CyQ2rvW+8wCE/ bwsP6z6+DNIFzM6IeBgLmE1qPzCVFWlxq2wnbDQEXvk5I/2ObRDXdYYh3ogm9vV2 C+I8XQKBgHIquvifRVvWf1q9WFZLQXZMv1flPhNaLmR+2k6gNpJ8SeiOmBtE7gT6 Je+YOXEKvfr6jaaJwYHPi6IhWHs4fQbgdK4jei30sRL7c8QdKEuwRdHPimmGNAPb UapzHY7xq0Wk9enAnM/SXkjTAJEkrpiQiDuPZVi4sIYCOqb+Ovu5 -----END RSA PRIVATE KEY-----",
						"cert_region":    "ap-southeast-1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"https_ext":   "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
					"cert_name":   "692",
					"cert":        "-----BEGIN CERTIFICATE----- MIID5DCCAsygAwIBAgIQFO5STIOlR/KkRB3gDHsi5zANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MTMwNTEyMDdaFw0yNTA4MTMwNTEyMDdaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDwg/HpVcY9yaLrxApP4kvrGwwPbpqOXw9NDWqSD2ms7Qvhi84dTnZT 9xyLXjW0MYI+W13kes4Prl37L8VHOZFvxdzZdzuHvZMkJDaRcekIiUYp+5SFqzzT EtlQodDZBLy0uSlwHnDHymUNJg2nXma+cOOwBBhvb9h+j5v4uuUkQoHRzIUHCkd6 8LR4zbOe+zrhxUU5AYm8C77ZJphDXi9GYm/moajpk+0biCDZ/vZSjTngZEujQKam dgRfVCiWgoUrueiijh1cLGf+W15A3kNo5UXfQrhbHVRQY0vy9dqU1ZBms0pWdc9G dJGd2kXkGNWtEbk4NCN6c64AW2L3G2JNAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAFl4tmH+m/lGr7aKxxz5uYdqepcjCHTzW KpxxEHr+sL9fOkap6I/NA9MYKsuVfnrTMnEJXeAj24g8q6t38bGTEMcUvh3YGS+P 7DLIF2T2/f9oiO6LVD72zC/HkqaenzWla536G4FLAW4+MzAdn+R7hd5QN4CuFaKI nWWr5KMjZfgyxPkbe+WiD8QryR/EOd3za5vdycZRnBVIcbZQzEHwxizm3JyPNYy8 KeIjt5kGjFt9ccb1oPXjVibB2SUO322kbhMeCXKq2tswnZBICy+1EcbmaxNlkxwc hcdb2wz/3pzH0jbexgTFrNWeoJyYykM8s/3cMY/t0dKppve4FXLG9A== -----END CERTIFICATE-----",
					"key":         "-----BEGIN RSA PRIVATE KEY----- MIIEowIBAAKCAQEA8IPx6VXGPcmi68QKT+JL6xsMD26ajl8PTQ1qkg9prO0L4YvO HU52U/cci141tDGCPltd5HrOD65d+y/FRzmRb8Xc2Xc7h72TJCQ2kXHpCIlGKfuU has80xLZUKHQ2QS8tLkpcB5wx8plDSYNp15mvnDjsAQYb2/Yfo+b+LrlJEKB0cyF BwpHevC0eM2znvs64cVFOQGJvAu+2SaYQ14vRmJv5qGo6ZPtG4gg2f72Uo054GRL o0CmpnYEX1QoloKFK7nooo4dXCxn/lteQN5DaOVF30K4Wx1UUGNL8vXalNWQZrNK VnXPRnSRndpF5BjVrRG5ODQjenOuAFti9xtiTQIDAQABAoIBAQCkuln3bA3ox69U NuKxL9a7Ybzy3NfyZtz98xBolTHVhE083xn+LH0SqQ7dzVqO3dHMj5tRH2L+jnhD z8YYMC+SFDxcnTMilw6uFDdjilcGx65Mlsh0fIGeNyyr8wgtevcb+C2PYunvjImF Zei4FwnbqUnohgWOXVYz6Hv08Vx7ZdW+QiH62I/LS73G7d2EPb26Zo3zMKg/H5AD xuNk82MDW0lCgrw869Yqhcd3GkkmgWi+S71AE0ftY03QeBrsSZbzz4Zsgk2GsEBt fGclOu2c5sNRhLy4o7GiZghPS32zkiec9H76Ip5n/nwXgcYCoxfvOQL+b8U5vpap gbgjWGLBAoGBAPQsJGrYAzt1naGbzwEgYbbyItKD3Bf30YVcZxnZW1fZn2H9NpWl oIBAiO8ls1WP32Mf1us91Bp0V4ESmmklFb5ZllRZXnGt6U2pvxl2FZbK18Nvcvw/ kc66t643mKdfsTAF3sRSmeeWxNSB5C+/yZfHARbBoZ64hYqSut2sivjlAoGBAPwq dDBrz4P/PGn191ZfpWyTC4NRMpC2zYkHoWtMJNs96bsZub4phWyo1KLL4sbAZ+uV GfpRpE2u8mKQWEUU2Gz/KI0cY1e40Icg/ZBlzNglb7ssBXKMrI3R2pigIzEAArhU KwsjhreGF7ix5DEyNT4i7PQL/tOu3uh9SUXl9zVJAoGANta/KxvuxejpiUVUHZ2n NI53Ua55vQxUi04wfba6dCWVTU2wd7WmMYfM+WEPQPU6J6ob++N8AqEEkiGaemjw 1DqMr88OjhuQHXg1SkOiH6bZBLTAL3Ubi0GWRVOJPnYYdn+rA47FsCTFejDeDfdW EHeKgBDm+p3YqEHCJE0/PR0CgYA+mw+zweCIhgLqz81znVWFylAubyddtHT9E27p I8N2xz1TXYS3CLn+i0AXlwUbkUN7ws3rTv+65bd57xprNEyzavoXZrfnXJQxKGir xAqCk3DVCI3lrbVdlH9wKznxfW4vc34oSs60m88h5NChwjRj0+n+gUfoKF9hW1Go z/p7OQKBgCqfExvjjmBibsr0ZtTprMGV2qI55LoVnUXcM1xSBbgU0rLwvw2YBrWg MRl3ixYN851wGD1LhgpwFjr7SnwEhpKdDloKSE1ANM5LE7zvJHsPY2uvmY/Rbn7i RqIzeCbUVYt1Ow0WEqy0DUy/fGLQEz9viwLnvTcDNWeuSsljeAO7 -----END RSA PRIVATE KEY-----",
					"cert_region": "cn-hangzhou",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"https_ext":   CHECKSET,
						"cert_name":   CHECKSET,
						"cert":        "-----BEGIN CERTIFICATE----- MIID5DCCAsygAwIBAgIQFO5STIOlR/KkRB3gDHsi5zANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MTMwNTEyMDdaFw0yNTA4MTMwNTEyMDdaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDwg/HpVcY9yaLrxApP4kvrGwwPbpqOXw9NDWqSD2ms7Qvhi84dTnZT 9xyLXjW0MYI+W13kes4Prl37L8VHOZFvxdzZdzuHvZMkJDaRcekIiUYp+5SFqzzT EtlQodDZBLy0uSlwHnDHymUNJg2nXma+cOOwBBhvb9h+j5v4uuUkQoHRzIUHCkd6 8LR4zbOe+zrhxUU5AYm8C77ZJphDXi9GYm/moajpk+0biCDZ/vZSjTngZEujQKam dgRfVCiWgoUrueiijh1cLGf+W15A3kNo5UXfQrhbHVRQY0vy9dqU1ZBms0pWdc9G dJGd2kXkGNWtEbk4NCN6c64AW2L3G2JNAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAFl4tmH+m/lGr7aKxxz5uYdqepcjCHTzW KpxxEHr+sL9fOkap6I/NA9MYKsuVfnrTMnEJXeAj24g8q6t38bGTEMcUvh3YGS+P 7DLIF2T2/f9oiO6LVD72zC/HkqaenzWla536G4FLAW4+MzAdn+R7hd5QN4CuFaKI nWWr5KMjZfgyxPkbe+WiD8QryR/EOd3za5vdycZRnBVIcbZQzEHwxizm3JyPNYy8 KeIjt5kGjFt9ccb1oPXjVibB2SUO322kbhMeCXKq2tswnZBICy+1EcbmaxNlkxwc hcdb2wz/3pzH0jbexgTFrNWeoJyYykM8s/3cMY/t0dKppve4FXLG9A== -----END CERTIFICATE-----",
						"key":         "-----BEGIN RSA PRIVATE KEY----- MIIEowIBAAKCAQEA8IPx6VXGPcmi68QKT+JL6xsMD26ajl8PTQ1qkg9prO0L4YvO HU52U/cci141tDGCPltd5HrOD65d+y/FRzmRb8Xc2Xc7h72TJCQ2kXHpCIlGKfuU has80xLZUKHQ2QS8tLkpcB5wx8plDSYNp15mvnDjsAQYb2/Yfo+b+LrlJEKB0cyF BwpHevC0eM2znvs64cVFOQGJvAu+2SaYQ14vRmJv5qGo6ZPtG4gg2f72Uo054GRL o0CmpnYEX1QoloKFK7nooo4dXCxn/lteQN5DaOVF30K4Wx1UUGNL8vXalNWQZrNK VnXPRnSRndpF5BjVrRG5ODQjenOuAFti9xtiTQIDAQABAoIBAQCkuln3bA3ox69U NuKxL9a7Ybzy3NfyZtz98xBolTHVhE083xn+LH0SqQ7dzVqO3dHMj5tRH2L+jnhD z8YYMC+SFDxcnTMilw6uFDdjilcGx65Mlsh0fIGeNyyr8wgtevcb+C2PYunvjImF Zei4FwnbqUnohgWOXVYz6Hv08Vx7ZdW+QiH62I/LS73G7d2EPb26Zo3zMKg/H5AD xuNk82MDW0lCgrw869Yqhcd3GkkmgWi+S71AE0ftY03QeBrsSZbzz4Zsgk2GsEBt fGclOu2c5sNRhLy4o7GiZghPS32zkiec9H76Ip5n/nwXgcYCoxfvOQL+b8U5vpap gbgjWGLBAoGBAPQsJGrYAzt1naGbzwEgYbbyItKD3Bf30YVcZxnZW1fZn2H9NpWl oIBAiO8ls1WP32Mf1us91Bp0V4ESmmklFb5ZllRZXnGt6U2pvxl2FZbK18Nvcvw/ kc66t643mKdfsTAF3sRSmeeWxNSB5C+/yZfHARbBoZ64hYqSut2sivjlAoGBAPwq dDBrz4P/PGn191ZfpWyTC4NRMpC2zYkHoWtMJNs96bsZub4phWyo1KLL4sbAZ+uV GfpRpE2u8mKQWEUU2Gz/KI0cY1e40Icg/ZBlzNglb7ssBXKMrI3R2pigIzEAArhU KwsjhreGF7ix5DEyNT4i7PQL/tOu3uh9SUXl9zVJAoGANta/KxvuxejpiUVUHZ2n NI53Ua55vQxUi04wfba6dCWVTU2wd7WmMYfM+WEPQPU6J6ob++N8AqEEkiGaemjw 1DqMr88OjhuQHXg1SkOiH6bZBLTAL3Ubi0GWRVOJPnYYdn+rA47FsCTFejDeDfdW EHeKgBDm+p3YqEHCJE0/PR0CgYA+mw+zweCIhgLqz81znVWFylAubyddtHT9E27p I8N2xz1TXYS3CLn+i0AXlwUbkUN7ws3rTv+65bd57xprNEyzavoXZrfnXJQxKGir xAqCk3DVCI3lrbVdlH9wKznxfW4vc34oSs60m88h5NChwjRj0+n+gUfoKF9hW1Go z/p7OQKBgCqfExvjjmBibsr0ZtTprMGV2qI55LoVnUXcM1xSBbgU0rLwvw2YBrWg MRl3ixYN851wGD1LhgpwFjr7SnwEhpKdDloKSE1ANM5LE7zvJHsPY2uvmY/Rbn7i RqIzeCbUVYt1Ow0WEqy0DUy/fGLQEz9viwLnvTcDNWeuSsljeAO7 -----END RSA PRIVATE KEY-----",
						"cert_region": "cn-hangzhou",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap7629 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence7629(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_ddoscoo_instance" "defaultaRfzZ9" {
  normal_bandwidth = "100"
  normal_qps       = "500"
  bandwidth_mode   = "2"
  product_plan     = "3"
  product_type     = "ddosDip"
  period           = "1"
  port_count       = "5"
  name             = "测试手动上传证书"
  function_version = "0"
  domain_count     = "10"
}


`, name)
}

// Case 域名资源测试_副本1686721083373 3448
func TestAccAliCloudDdosCooDomainResource_basic3448(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap3448)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence3448)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type": "0",
					"real_servers": []string{
						"1.1.1.1"},
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
					},
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaulttV5f6b.id}"},
					"domain":       "ldtest123.qq.com",
					"https_ext":    "{\\\"Https2http\\\":1,\\\"Http2\\\":1,\\\"Http2https\\\":0}",
					"ocsp_enabled": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "0",
						"real_servers.#": "1",
						"proxy_types.#":  "1",
						"instance_ids.#": "1",
						"domain":         "ldtest123.qq.com",
						"https_ext":      CHECKSET,
						"ocsp_enabled":   "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type": "1",
					"real_servers": []string{
						"qq.com"},
					"https_ext": "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "1",
						"real_servers.#": "1",
						"https_ext":      CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap3448 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence3448(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_ddoscoo_instance" "defaulttV5f6b" {
  product_type = "ddoscoo"
  period       = "1"
}


`, name)
}

// Case 测试选择证书-李铎 7632
func TestAccAliCloudDdosCooDomainResource_basic7632(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap7632)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence7632)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type":      "0",
					"ocsp_enabled": "0",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "websocket",
						},
					},
					"real_servers": []string{
						"1.1.1.1", "2.2.2.2", "3.3.3.3"},
					"domain": "testld.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultSJe7n8.id}", "${alicloud_ddoscoo_instance.default6lyurZ.id}", "${alicloud_ddoscoo_instance.defaultTTvY0D.id}"},
					"https_ext": "{\\\"Https2http\\\":1,\\\"Http2\\\":1,\\\"Http2https\\\":0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "0",
						"ocsp_enabled":   "0",
						"proxy_types.#":  "3",
						"real_servers.#": "3",
						"domain":         "testld.qq.com",
						"instance_ids.#": "3",
						"https_ext":      CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type":      "1",
					"ocsp_enabled": "1",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80", "8080"},
							"proxy_type": "http",
						},
					},
					"real_servers": []string{
						"1.qq.com"},
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultTTvY0D.id}"},
					"https_ext":       "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
					"cert_identifier": "${alicloud_ssl_certificates_service_certificate.default3MYZEt.id}",
					"cert_region":     "cn-hangzhou",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":         "1",
						"ocsp_enabled":    "1",
						"proxy_types.#":   "1",
						"real_servers.#":  "1",
						"instance_ids.#":  "1",
						"https_ext":       CHECKSET,
						"cert_identifier": CHECKSET,
						"cert_region":     "cn-hangzhou",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap7632 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence7632(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_ddoscoo_instance" "defaultTTvY0D" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "0"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ddoscoo_instance" "defaultSJe7n8" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test2"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "1"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ddoscoo_instance" "default6lyurZ" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test2"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "1"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ssl_certificates_service_certificate" "default3MYZEt" {
  cert             = <<EOF
-----BEGIN CERTIFICATE----- MIID4TCCAsmgAwIBAgIRANZGvLwT8kuWpPlZ/Aj+uPgwDQYJKoZIhvcNAQELBQAw XjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU ZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w HhcNMjQwODIzMDk0NzA0WhcNMjUwODIzMDk0NzA0WjAlMQswCQYDVQQGEwJDTjEW MBQGA1UEAxMNdGVzdGxkLnFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC AQoCggEBALmZY2geTFi+50gAVyDQH9Y5sTv8LLX6+MET1l3larzjX1M0Az9ZEIc0 TNrAp8mtJRlpQCzyDPZg88AwSdEwqSOSsnGzfS2DUcPJmdn2a2n5PLvWE28qPuSf 6fl3IhNiPzLYR51+7ccJKEQRhfOK2usmJo6oTG/0Lhh4BRH5owcclKv6n3YHaBVj JNigiq1/tlqU46toZvotPOORjpy21kJPZioHqOVCDO4zreMy2xuIiYtpSSmXwkEO zcQQ3K8sbRx9ED8SCdb229h7ioTug02YBXs0YOQZ024HFaIF8Nz1M+mdHy1jCbLd yJoT/jzE4RdldZKZJFaSKV1c7EYlzhkCAwEAAaOB0jCBzzAOBgNVHQ8BAf8EBAMC BaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMB8GA1UdIwQYMBaAFCiB JgXRNBo/wXMPu5PPFRw/A79/MGMGCCsGAQUFBwEBBFcwVTAhBggrBgEFBQcwAYYV aHR0cDovL29jc3AubXlzc2wuY29tMDAGCCsGAQUFBzAChiRodHRwOi8vY2EubXlz c2wuY29tL215c3NsdGVzdHJzYS5jcnQwGAYDVR0RBBEwD4INdGVzdGxkLnFxLmNv bTANBgkqhkiG9w0BAQsFAAOCAQEAnPJl1GrePDIulWfsETPbGnrZv3j3ZRXuou0o K32X/nydS/i/j+AUzKSyezmnR1edkgY1hbGaza702SLQJuGh2IqJvAFyifwV/CZ5 cpJIi5G7kWTBjZo9NgVnDMhR8y5DCKE8BhiUBwcSvKKC8se2yWHm1fk9pRxG0Mc6 0fstl40jtR5XZYsW1GhX4fzwrWuBodPKticgXPn2e24ec+4rVrziu5R7D77AzJjG Y/wzNYvAUWEzEya7Ve53nhu+WpIuIQn0ux8nPDioFdOjckn4jK3ePYdS2mWT6EBU BC74GYiBNDz0QgHADq1VTExeLzC0tw9PPdWl0WfoTgCCKLz0yA== -----END CERTIFICATE-----
EOF
  certificate_name = "755"
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuZljaB5MWL7nSABXINAf1jmxO/wstfr4wRPWXeVqvONfUzQD P1kQhzRM2sCnya0lGWlALPIM9mDzwDBJ0TCpI5KycbN9LYNRw8mZ2fZrafk8u9YT byo+5J/p+XciE2I/MthHnX7txwkoRBGF84ra6yYmjqhMb/QuGHgFEfmjBxyUq/qf dgdoFWMk2KCKrX+2WpTjq2hm+i0845GOnLbWQk9mKgeo5UIM7jOt4zLbG4iJi2lJ KZfCQQ7NxBDcryxtHH0QPxIJ1vbb2HuKhO6DTZgFezRg5BnTbgcVogXw3PUz6Z0f LWMJst3ImhP+PMThF2V1kpkkVpIpXVzsRiXOGQIDAQABAoIBAQCBPiw4A+k8X2vk +r+xjNyurCwcTmXAL81rfmnnputmL5tg8DZWtanJzQS7zC7LRPQxttZGtiOKqkbz DW1J6+3MZMo4XToNKIYWpduqKWvxNusxDkkoPy3evPEMlAY5o0/JE00DgrEHyfut MtqplocN+tocu1vHFi3HQkSdmM4LE46ZfFu5w1FRbNI1Gqjj/cwlF/T93V3qMap1 WfsJjhMIX9LjBq3y9GAfAtAw7JYwkztr2AhYzCsK25wAj72zFY6FJTZ8LklfS41q DrVtdjMx42IonDQtkzrqzfYlXdzzhZzuQHxn+qJODseoU8oDG9j3eKhp1dKgqLfx tv1o3km1AoGBAOqXGEw2w94uVchCjuTum3XFYieEla0IUbHJCaWKU/hoSbht7j23 K7tA9//epBuRLGtYE0sPBK6i31mQT216muspO1g3pwGJhPy8VSpsJ1GQhB2G2UNz kZlRK+2/gx35TdTi9x0C6UWk5XkhgWO3R35BlEnuV7EOyJunobiUcoObAoGBAMqJ reuSbJajNGfeBzPel7F62ZDufC85hWaGKzeXIk3DkXcsEpeR5ogMGHCZ3dGf4Yz/ pcfjnCMIWjc+MkA4ppFd4432FJkxNQQP0z7njpXW3e5tionsMN+UwPzi4wZPKufK osjw43JpBzpHGxG3ynLgZg7bSfrQZhPDTQw84nJbAoGBAKpy8mKeAB71R7rkMXNB s48Uxca03RQGUWV+DxZKtcxt6fKpXUtWRd4ezJMLL+4fw0iTjCEjXmGNUf9/jVac mOd44/erKBtD0m7YYIEcaE0pVfUmP8J0vDvL8MEkP56Nv/GIn8hijx/dOiaTI7JS Pw4LlDVLikfJ2BTQ7f5xTes1AoGBAJ7h4HiDFgIZp1uvtfC/tjn5CEGEhBC7y+VA bRify747I5rcDP2v66tf6bAzU+pExLhKN++Vov9sZvEdLmhoyGoSwBa2KzR9gHxe ObYICjeLJfALKHnHuhM6ayY2iieB5UOOF6MQLSysLYpPC3IbvonddNJEvkUuRFVO iNuHy5AvAoGAPqbzSNA05gf85zRO/JZAmZuWXG3o3pVougbkc211p+ynMpS+/bMb c/nR36kOE551lFjIoAjoeIs16Wbq+00u9GlcQmyAdFpfaFCNHa3dayJwKJMW9Nia fKbiiiOAQE8s2v8Paa+b00GspeWLow4u0G5lBVau4JjEVnl6ivLXlzY= -----END RSA PRIVATE KEY-----
EOF
}


`, name)
}

// Case 测试选择证书-TF接入_ld 7936
func TestAccAliCloudDdosCooDomainResource_basic7936(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap7936)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence7936)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type":      "0",
					"ocsp_enabled": "false",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "websocket",
						},
					},
					"real_servers": []string{
						"1.1.1.1", "2.2.2.2", "3.3.3.3"},
					"domain": "testld.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultSJe7n8.id}", "${alicloud_ddoscoo_instance.default6lyurZ.id}", "${alicloud_ddoscoo_instance.defaultTTvY0D.id}"},
					"https_ext": "{\\\"Https2http\\\":1,\\\"Http2\\\":1,\\\"Http2https\\\":0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "0",
						"ocsp_enabled":   "false",
						"proxy_types.#":  "3",
						"real_servers.#": "3",
						"domain":         "testld.qq.com",
						"instance_ids.#": "3",
						"https_ext":      CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type":      "1",
					"ocsp_enabled": "true",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80", "8080"},
							"proxy_type": "http",
						},
					},
					"real_servers": []string{
						"1.qq.com"},
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultTTvY0D.id}"},
					"https_ext":       "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
					"cert_identifier": "${alicloud_ssl_certificates_service_certificate.default3MYZEt.id}",
					"cert_region":     "cn-hangzhou",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":         "1",
						"ocsp_enabled":    "true",
						"proxy_types.#":   "1",
						"real_servers.#":  "1",
						"instance_ids.#":  "1",
						"https_ext":       CHECKSET,
						"cert_identifier": CHECKSET,
						"cert_region":     "cn-hangzhou",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap7936 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence7936(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_ddoscoo_instance" "defaultTTvY0D" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "0"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ddoscoo_instance" "defaultSJe7n8" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test2"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "1"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ddoscoo_instance" "default6lyurZ" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test2"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "1"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ssl_certificates_service_certificate" "default3MYZEt" {
  cert             = <<EOF
-----BEGIN CERTIFICATE-----
MIID4TCCAsmgAwIBAgIRANZGvLwT8kuWpPlZ/Aj+uPgwDQYJKoZIhvcNAQELBQAw
XjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU
ZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w
HhcNMjQwODIzMDk0NzA0WhcNMjUwODIzMDk0NzA0WjAlMQswCQYDVQQGEwJDTjEW
MBQGA1UEAxMNdGVzdGxkLnFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBALmZY2geTFi+50gAVyDQH9Y5sTv8LLX6+MET1l3larzjX1M0Az9ZEIc0
TNrAp8mtJRlpQCzyDPZg88AwSdEwqSOSsnGzfS2DUcPJmdn2a2n5PLvWE28qPuSf
6fl3IhNiPzLYR51+7ccJKEQRhfOK2usmJo6oTG/0Lhh4BRH5owcclKv6n3YHaBVj
JNigiq1/tlqU46toZvotPOORjpy21kJPZioHqOVCDO4zreMy2xuIiYtpSSmXwkEO
zcQQ3K8sbRx9ED8SCdb229h7ioTug02YBXs0YOQZ024HFaIF8Nz1M+mdHy1jCbLd
yJoT/jzE4RdldZKZJFaSKV1c7EYlzhkCAwEAAaOB0jCBzzAOBgNVHQ8BAf8EBAMC
BaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMB8GA1UdIwQYMBaAFCiB
JgXRNBo/wXMPu5PPFRw/A79/MGMGCCsGAQUFBwEBBFcwVTAhBggrBgEFBQcwAYYV
aHR0cDovL29jc3AubXlzc2wuY29tMDAGCCsGAQUFBzAChiRodHRwOi8vY2EubXlz
c2wuY29tL215c3NsdGVzdHJzYS5jcnQwGAYDVR0RBBEwD4INdGVzdGxkLnFxLmNv
bTANBgkqhkiG9w0BAQsFAAOCAQEAnPJl1GrePDIulWfsETPbGnrZv3j3ZRXuou0o
K32X/nydS/i/j+AUzKSyezmnR1edkgY1hbGaza702SLQJuGh2IqJvAFyifwV/CZ5
cpJIi5G7kWTBjZo9NgVnDMhR8y5DCKE8BhiUBwcSvKKC8se2yWHm1fk9pRxG0Mc6
0fstl40jtR5XZYsW1GhX4fzwrWuBodPKticgXPn2e24ec+4rVrziu5R7D77AzJjG
Y/wzNYvAUWEzEya7Ve53nhu+WpIuIQn0ux8nPDioFdOjckn4jK3ePYdS2mWT6EBU
BC74GYiBNDz0QgHADq1VTExeLzC0tw9PPdWl0WfoTgCCKLz0yA==
-----END CERTIFICATE-----
EOF
  certificate_name = "530"
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAuZljaB5MWL7nSABXINAf1jmxO/wstfr4wRPWXeVqvONfUzQD
P1kQhzRM2sCnya0lGWlALPIM9mDzwDBJ0TCpI5KycbN9LYNRw8mZ2fZrafk8u9YT
byo+5J/p+XciE2I/MthHnX7txwkoRBGF84ra6yYmjqhMb/QuGHgFEfmjBxyUq/qf
dgdoFWMk2KCKrX+2WpTjq2hm+i0845GOnLbWQk9mKgeo5UIM7jOt4zLbG4iJi2lJ
KZfCQQ7NxBDcryxtHH0QPxIJ1vbb2HuKhO6DTZgFezRg5BnTbgcVogXw3PUz6Z0f
LWMJst3ImhP+PMThF2V1kpkkVpIpXVzsRiXOGQIDAQABAoIBAQCBPiw4A+k8X2vk
+r+xjNyurCwcTmXAL81rfmnnputmL5tg8DZWtanJzQS7zC7LRPQxttZGtiOKqkbz
DW1J6+3MZMo4XToNKIYWpduqKWvxNusxDkkoPy3evPEMlAY5o0/JE00DgrEHyfut
MtqplocN+tocu1vHFi3HQkSdmM4LE46ZfFu5w1FRbNI1Gqjj/cwlF/T93V3qMap1
WfsJjhMIX9LjBq3y9GAfAtAw7JYwkztr2AhYzCsK25wAj72zFY6FJTZ8LklfS41q
DrVtdjMx42IonDQtkzrqzfYlXdzzhZzuQHxn+qJODseoU8oDG9j3eKhp1dKgqLfx
tv1o3km1AoGBAOqXGEw2w94uVchCjuTum3XFYieEla0IUbHJCaWKU/hoSbht7j23
K7tA9//epBuRLGtYE0sPBK6i31mQT216muspO1g3pwGJhPy8VSpsJ1GQhB2G2UNz
kZlRK+2/gx35TdTi9x0C6UWk5XkhgWO3R35BlEnuV7EOyJunobiUcoObAoGBAMqJ
reuSbJajNGfeBzPel7F62ZDufC85hWaGKzeXIk3DkXcsEpeR5ogMGHCZ3dGf4Yz/
pcfjnCMIWjc+MkA4ppFd4432FJkxNQQP0z7njpXW3e5tionsMN+UwPzi4wZPKufK
osjw43JpBzpHGxG3ynLgZg7bSfrQZhPDTQw84nJbAoGBAKpy8mKeAB71R7rkMXNB
s48Uxca03RQGUWV+DxZKtcxt6fKpXUtWRd4ezJMLL+4fw0iTjCEjXmGNUf9/jVac
mOd44/erKBtD0m7YYIEcaE0pVfUmP8J0vDvL8MEkP56Nv/GIn8hijx/dOiaTI7JS
Pw4LlDVLikfJ2BTQ7f5xTes1AoGBAJ7h4HiDFgIZp1uvtfC/tjn5CEGEhBC7y+VA
bRify747I5rcDP2v66tf6bAzU+pExLhKN++Vov9sZvEdLmhoyGoSwBa2KzR9gHxe
ObYICjeLJfALKHnHuhM6ayY2iieB5UOOF6MQLSysLYpPC3IbvonddNJEvkUuRFVO
iNuHy5AvAoGAPqbzSNA05gf85zRO/JZAmZuWXG3o3pVougbkc211p+ynMpS+/bMb
c/nR36kOE551lFjIoAjoeIs16Wbq+00u9GlcQmyAdFpfaFCNHa3dayJwKJMW9Nia
fKbiiiOAQE8s2v8Paa+b00GspeWLow4u0G5lBVau4JjEVnl6ivLXlzY=
-----END RSA PRIVATE KEY-----
EOF
}


`, name)
}

// Case 手动上传证书_TF接入 7935
func TestAccAliCloudDdosCooDomainResource_basic7935(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap7935)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence7935)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"ap-southeast-1"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type": "0",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
					},
					"real_servers": []string{
						"1.1.1.1"},
					"domain": "testcert.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultaRfzZ9.id}"},
					"https_ext":   "{\\\"Https2http\\\":1,\\\"Http2\\\":1,\\\"Http2https\\\":0}",
					"cert_name":   "808",
					"cert":        "-----BEGIN CERTIFICATE-----\\nMIID5DCCAsygAwIBAgIQWen3GebvT0GcE/a1MJVFgjANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MjMwOTE1MjNaFw0yNTA4MjMwOTE1MjNaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDF/H2/Oas8trhkHs7B4B8mN00eup/5Tqar2QO4Hm499GJECKc3eMiC v+aAlW74Iymb2Varnv+WMdFRVMQgpXesi3akvVp0QxecvcDliilkh4ddTK731Rd7 PaSK1JdQX1jdGGhVnhQz+cPNFBGZ3tMYGhUkgNfqa3UFucJcBuRub/Ircr+5Ob4D FxSglfTHi+/EFcp7vMAOztLD4zXmEz3NysDNP6NzN7SD72DwPp0nxyRjrBlHSOVg szB/bFasQdAhZGeo64MvSb+SivdWEMhHkwKA5MhhYOkDeNPPSmlxbw0Z3nOyeMmI YkaxzhpO5DZN382duTQmiQ+Yg60OfL3NAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAKUCvcbUCJRFbVowd1YorILivqRmS6ztR 9vLdj4YZBWxmmQrgkDlkl78r/rXlJbqHunSh2Wbag7y+GaQQwg8xcL4Z3KrKj4zg aHEP1DyFiaxMTEuC/L2RgSX3xlXcf6fQ46D3y3Ja3iHFQnjx6npNaSZ2bSEULvJg IjPiJ/nbid1TYR5vtg8vtwDQfY7/+q8/3DWKcQq+SGcd9dDS6u9vulNdlW8e14bS CKDEzS/axjoICl9JagASLcElWIit/eD5zGIKzTPC9mEXiX/J/gUr70y9GiE7Ue++ 4nFSGQOwjMh/wO2HlmRfToeZ3g6rRCijibBHKHBmVym7NCai2voZJQ==\\n-----END CERTIFICATE-----",
					"key":         "-----BEGIN RSA PRIVATE KEY-----\\nMIIEowIBAAKCAQEAxfx9vzmrPLa4ZB7OweAfJjdNHrqf+U6mq9kDuB5uPfRiRAin N3jIgr/mgJVu+CMpm9lWq57/ljHRUVTEIKV3rIt2pL1adEMXnL3A5YopZIeHXUyu 99UXez2kitSXUF9Y3RhoVZ4UM/nDzRQRmd7TGBoVJIDX6mt1BbnCXAbkbm/yK3K/ uTm+AxcUoJX0x4vvxBXKe7zADs7Sw+M15hM9zcrAzT+jcze0g+9g8D6dJ8ckY6wZ R0jlYLMwf2xWrEHQIWRnqOuDL0m/kor3VhDIR5MCgOTIYWDpA3jTz0ppcW8NGd5z snjJiGJGsc4aTuQ2Td/Nnbk0JokPmIOtDny9zQIDAQABAoIBAQDD6fU4y8UhwCG4 mS+5c6D/PQvoU35Hwkd1l7pxcFNgpTqz3egyISgxEdny9WwoyQq8eJWmICEEK+nY VEv7jiFdMWhG3kTq9RUhejeuLEiHfQE7Fs2w2kFxJ29yHapZ0u/pYOSljFarlATo I2rDW1aB7BVt2L1P7+ONteKZFAzpJckft5ceRUzs5Jm1Cqt8OWO3Km+FBbCROv8M TevW44aoMwBGXuqs06FV1Z4dafglskjt2O38V4acZpH8Nc8j+nCONKL3OxwKY6HQ WfnbXnTLCF3IuMiy8ntrY8HYU6EABiCdr+Pl5HmhI2nmtSFTFbD4Gq70vgPL0P1m iULJGJ7hAoGBANQPrOGe9qHcBydvcBHE7qA9v1+IaTj03qzDTopTi/jxcd9pEkei skLyHNQ5yJT0QjTxB9iYRLfZccOGFyqz/Sdz6CwwTWBZeXOQ2AX7FPEcCnNr1TpF yMrgOY3H93KJISEVS6kYskByjK7XzXCp0KQNS2EeIhAXcqXxNmSwylvpAoGBAO8C PdZHd6aLLEZyVO1aZVHDxqmbhmGVoY9wZ+uwR4K2Hu/fjk0qlR9cYpw8+N675Wr9 E9Ff5/wjK2+/+uocQV9Zoap2vgrwX7GASuO5KYdCOBn6oUOSa+Ru+LgBNyUkXYES mM8eFC1QqfcSrETLAQqd2lmLcuaMq6jJtbBpvzhFAoGAYd8mNC9wtr1dE+dLuvfA BnbZJ1dG8QGa7/NoAVGT7X5JxwmwZR2C1oD1q0FMAOtGzzZbH60PMicKaWousQfH E/lbs2FLpOdGtX6pJQF/5dPCQwkGrVFd3bxk87nRy6vcfW9drxp10mbL5To2WAQY Bk8Ydic5I2IfCNVt/ETX8FkCgYA+OkAtVQgi9WM+qC/SaFGu2yETManoKFQbC3IT HB9SOeaOH49mKesPcjc+ZGWLYDJYC7IoNicpL2L0wnAqmdavY5/CyQ2rvW+8wCE/ bwsP6z6+DNIFzM6IeBgLmE1qPzCVFWlxq2wnbDQEXvk5I/2ObRDXdYYh3ogm9vV2 C+I8XQKBgHIquvifRVvWf1q9WFZLQXZMv1flPhNaLmR+2k6gNpJ8SeiOmBtE7gT6 Je+YOXEKvfr6jaaJwYHPi6IhWHs4fQbgdK4jei30sRL7c8QdKEuwRdHPimmGNAPb UapzHY7xq0Wk9enAnM/SXkjTAJEkrpiQiDuPZVi4sIYCOqb+Ovu5\\n-----END RSA PRIVATE KEY-----",
					"cert_region": "ap-southeast-1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "0",
						"proxy_types.#":  "2",
						"real_servers.#": "1",
						"domain":         "testcert.qq.com",
						"instance_ids.#": "1",
						"https_ext":      CHECKSET,
						"cert_name":      CHECKSET,
						"cert":           "-----BEGIN CERTIFICATE-----\nMIID5DCCAsygAwIBAgIQWen3GebvT0GcE/a1MJVFgjANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MjMwOTE1MjNaFw0yNTA4MjMwOTE1MjNaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDF/H2/Oas8trhkHs7B4B8mN00eup/5Tqar2QO4Hm499GJECKc3eMiC v+aAlW74Iymb2Varnv+WMdFRVMQgpXesi3akvVp0QxecvcDliilkh4ddTK731Rd7 PaSK1JdQX1jdGGhVnhQz+cPNFBGZ3tMYGhUkgNfqa3UFucJcBuRub/Ircr+5Ob4D FxSglfTHi+/EFcp7vMAOztLD4zXmEz3NysDNP6NzN7SD72DwPp0nxyRjrBlHSOVg szB/bFasQdAhZGeo64MvSb+SivdWEMhHkwKA5MhhYOkDeNPPSmlxbw0Z3nOyeMmI YkaxzhpO5DZN382duTQmiQ+Yg60OfL3NAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAKUCvcbUCJRFbVowd1YorILivqRmS6ztR 9vLdj4YZBWxmmQrgkDlkl78r/rXlJbqHunSh2Wbag7y+GaQQwg8xcL4Z3KrKj4zg aHEP1DyFiaxMTEuC/L2RgSX3xlXcf6fQ46D3y3Ja3iHFQnjx6npNaSZ2bSEULvJg IjPiJ/nbid1TYR5vtg8vtwDQfY7/+q8/3DWKcQq+SGcd9dDS6u9vulNdlW8e14bS CKDEzS/axjoICl9JagASLcElWIit/eD5zGIKzTPC9mEXiX/J/gUr70y9GiE7Ue++ 4nFSGQOwjMh/wO2HlmRfToeZ3g6rRCijibBHKHBmVym7NCai2voZJQ==\n-----END CERTIFICATE-----",
						"key":            "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAxfx9vzmrPLa4ZB7OweAfJjdNHrqf+U6mq9kDuB5uPfRiRAin N3jIgr/mgJVu+CMpm9lWq57/ljHRUVTEIKV3rIt2pL1adEMXnL3A5YopZIeHXUyu 99UXez2kitSXUF9Y3RhoVZ4UM/nDzRQRmd7TGBoVJIDX6mt1BbnCXAbkbm/yK3K/ uTm+AxcUoJX0x4vvxBXKe7zADs7Sw+M15hM9zcrAzT+jcze0g+9g8D6dJ8ckY6wZ R0jlYLMwf2xWrEHQIWRnqOuDL0m/kor3VhDIR5MCgOTIYWDpA3jTz0ppcW8NGd5z snjJiGJGsc4aTuQ2Td/Nnbk0JokPmIOtDny9zQIDAQABAoIBAQDD6fU4y8UhwCG4 mS+5c6D/PQvoU35Hwkd1l7pxcFNgpTqz3egyISgxEdny9WwoyQq8eJWmICEEK+nY VEv7jiFdMWhG3kTq9RUhejeuLEiHfQE7Fs2w2kFxJ29yHapZ0u/pYOSljFarlATo I2rDW1aB7BVt2L1P7+ONteKZFAzpJckft5ceRUzs5Jm1Cqt8OWO3Km+FBbCROv8M TevW44aoMwBGXuqs06FV1Z4dafglskjt2O38V4acZpH8Nc8j+nCONKL3OxwKY6HQ WfnbXnTLCF3IuMiy8ntrY8HYU6EABiCdr+Pl5HmhI2nmtSFTFbD4Gq70vgPL0P1m iULJGJ7hAoGBANQPrOGe9qHcBydvcBHE7qA9v1+IaTj03qzDTopTi/jxcd9pEkei skLyHNQ5yJT0QjTxB9iYRLfZccOGFyqz/Sdz6CwwTWBZeXOQ2AX7FPEcCnNr1TpF yMrgOY3H93KJISEVS6kYskByjK7XzXCp0KQNS2EeIhAXcqXxNmSwylvpAoGBAO8C PdZHd6aLLEZyVO1aZVHDxqmbhmGVoY9wZ+uwR4K2Hu/fjk0qlR9cYpw8+N675Wr9 E9Ff5/wjK2+/+uocQV9Zoap2vgrwX7GASuO5KYdCOBn6oUOSa+Ru+LgBNyUkXYES mM8eFC1QqfcSrETLAQqd2lmLcuaMq6jJtbBpvzhFAoGAYd8mNC9wtr1dE+dLuvfA BnbZJ1dG8QGa7/NoAVGT7X5JxwmwZR2C1oD1q0FMAOtGzzZbH60PMicKaWousQfH E/lbs2FLpOdGtX6pJQF/5dPCQwkGrVFd3bxk87nRy6vcfW9drxp10mbL5To2WAQY Bk8Ydic5I2IfCNVt/ETX8FkCgYA+OkAtVQgi9WM+qC/SaFGu2yETManoKFQbC3IT HB9SOeaOH49mKesPcjc+ZGWLYDJYC7IoNicpL2L0wnAqmdavY5/CyQ2rvW+8wCE/ bwsP6z6+DNIFzM6IeBgLmE1qPzCVFWlxq2wnbDQEXvk5I/2ObRDXdYYh3ogm9vV2 C+I8XQKBgHIquvifRVvWf1q9WFZLQXZMv1flPhNaLmR+2k6gNpJ8SeiOmBtE7gT6 Je+YOXEKvfr6jaaJwYHPi6IhWHs4fQbgdK4jei30sRL7c8QdKEuwRdHPimmGNAPb UapzHY7xq0Wk9enAnM/SXkjTAJEkrpiQiDuPZVi4sIYCOqb+Ovu5\n-----END RSA PRIVATE KEY-----",
						"cert_region":    "ap-southeast-1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"https_ext":   "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
					"cert_name":   "519",
					"cert":        "-----BEGIN CERTIFICATE-----\\nMIID5DCCAsygAwIBAgIQFO5STIOlR/KkRB3gDHsi5zANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MTMwNTEyMDdaFw0yNTA4MTMwNTEyMDdaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDwg/HpVcY9yaLrxApP4kvrGwwPbpqOXw9NDWqSD2ms7Qvhi84dTnZT 9xyLXjW0MYI+W13kes4Prl37L8VHOZFvxdzZdzuHvZMkJDaRcekIiUYp+5SFqzzT EtlQodDZBLy0uSlwHnDHymUNJg2nXma+cOOwBBhvb9h+j5v4uuUkQoHRzIUHCkd6 8LR4zbOe+zrhxUU5AYm8C77ZJphDXi9GYm/moajpk+0biCDZ/vZSjTngZEujQKam dgRfVCiWgoUrueiijh1cLGf+W15A3kNo5UXfQrhbHVRQY0vy9dqU1ZBms0pWdc9G dJGd2kXkGNWtEbk4NCN6c64AW2L3G2JNAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAFl4tmH+m/lGr7aKxxz5uYdqepcjCHTzW KpxxEHr+sL9fOkap6I/NA9MYKsuVfnrTMnEJXeAj24g8q6t38bGTEMcUvh3YGS+P 7DLIF2T2/f9oiO6LVD72zC/HkqaenzWla536G4FLAW4+MzAdn+R7hd5QN4CuFaKI nWWr5KMjZfgyxPkbe+WiD8QryR/EOd3za5vdycZRnBVIcbZQzEHwxizm3JyPNYy8 KeIjt5kGjFt9ccb1oPXjVibB2SUO322kbhMeCXKq2tswnZBICy+1EcbmaxNlkxwc hcdb2wz/3pzH0jbexgTFrNWeoJyYykM8s/3cMY/t0dKppve4FXLG9A==\\n-----END CERTIFICATE-----",
					"key":         "-----BEGIN RSA PRIVATE KEY-----\\nMIIEowIBAAKCAQEA8IPx6VXGPcmi68QKT+JL6xsMD26ajl8PTQ1qkg9prO0L4YvO HU52U/cci141tDGCPltd5HrOD65d+y/FRzmRb8Xc2Xc7h72TJCQ2kXHpCIlGKfuU has80xLZUKHQ2QS8tLkpcB5wx8plDSYNp15mvnDjsAQYb2/Yfo+b+LrlJEKB0cyF BwpHevC0eM2znvs64cVFOQGJvAu+2SaYQ14vRmJv5qGo6ZPtG4gg2f72Uo054GRL o0CmpnYEX1QoloKFK7nooo4dXCxn/lteQN5DaOVF30K4Wx1UUGNL8vXalNWQZrNK VnXPRnSRndpF5BjVrRG5ODQjenOuAFti9xtiTQIDAQABAoIBAQCkuln3bA3ox69U NuKxL9a7Ybzy3NfyZtz98xBolTHVhE083xn+LH0SqQ7dzVqO3dHMj5tRH2L+jnhD z8YYMC+SFDxcnTMilw6uFDdjilcGx65Mlsh0fIGeNyyr8wgtevcb+C2PYunvjImF Zei4FwnbqUnohgWOXVYz6Hv08Vx7ZdW+QiH62I/LS73G7d2EPb26Zo3zMKg/H5AD xuNk82MDW0lCgrw869Yqhcd3GkkmgWi+S71AE0ftY03QeBrsSZbzz4Zsgk2GsEBt fGclOu2c5sNRhLy4o7GiZghPS32zkiec9H76Ip5n/nwXgcYCoxfvOQL+b8U5vpap gbgjWGLBAoGBAPQsJGrYAzt1naGbzwEgYbbyItKD3Bf30YVcZxnZW1fZn2H9NpWl oIBAiO8ls1WP32Mf1us91Bp0V4ESmmklFb5ZllRZXnGt6U2pvxl2FZbK18Nvcvw/ kc66t643mKdfsTAF3sRSmeeWxNSB5C+/yZfHARbBoZ64hYqSut2sivjlAoGBAPwq dDBrz4P/PGn191ZfpWyTC4NRMpC2zYkHoWtMJNs96bsZub4phWyo1KLL4sbAZ+uV GfpRpE2u8mKQWEUU2Gz/KI0cY1e40Icg/ZBlzNglb7ssBXKMrI3R2pigIzEAArhU KwsjhreGF7ix5DEyNT4i7PQL/tOu3uh9SUXl9zVJAoGANta/KxvuxejpiUVUHZ2n NI53Ua55vQxUi04wfba6dCWVTU2wd7WmMYfM+WEPQPU6J6ob++N8AqEEkiGaemjw 1DqMr88OjhuQHXg1SkOiH6bZBLTAL3Ubi0GWRVOJPnYYdn+rA47FsCTFejDeDfdW EHeKgBDm+p3YqEHCJE0/PR0CgYA+mw+zweCIhgLqz81znVWFylAubyddtHT9E27p I8N2xz1TXYS3CLn+i0AXlwUbkUN7ws3rTv+65bd57xprNEyzavoXZrfnXJQxKGir xAqCk3DVCI3lrbVdlH9wKznxfW4vc34oSs60m88h5NChwjRj0+n+gUfoKF9hW1Go z/p7OQKBgCqfExvjjmBibsr0ZtTprMGV2qI55LoVnUXcM1xSBbgU0rLwvw2YBrWg MRl3ixYN851wGD1LhgpwFjr7SnwEhpKdDloKSE1ANM5LE7zvJHsPY2uvmY/Rbn7i RqIzeCbUVYt1Ow0WEqy0DUy/fGLQEz9viwLnvTcDNWeuSsljeAO7\\n-----END RSA PRIVATE KEY-----",
					"cert_region": "cn-hangzhou",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"https_ext":   CHECKSET,
						"cert_name":   CHECKSET,
						"cert":        "-----BEGIN CERTIFICATE-----\nMIID5DCCAsygAwIBAgIQFO5STIOlR/KkRB3gDHsi5zANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MTMwNTEyMDdaFw0yNTA4MTMwNTEyMDdaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDwg/HpVcY9yaLrxApP4kvrGwwPbpqOXw9NDWqSD2ms7Qvhi84dTnZT 9xyLXjW0MYI+W13kes4Prl37L8VHOZFvxdzZdzuHvZMkJDaRcekIiUYp+5SFqzzT EtlQodDZBLy0uSlwHnDHymUNJg2nXma+cOOwBBhvb9h+j5v4uuUkQoHRzIUHCkd6 8LR4zbOe+zrhxUU5AYm8C77ZJphDXi9GYm/moajpk+0biCDZ/vZSjTngZEujQKam dgRfVCiWgoUrueiijh1cLGf+W15A3kNo5UXfQrhbHVRQY0vy9dqU1ZBms0pWdc9G dJGd2kXkGNWtEbk4NCN6c64AW2L3G2JNAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAFl4tmH+m/lGr7aKxxz5uYdqepcjCHTzW KpxxEHr+sL9fOkap6I/NA9MYKsuVfnrTMnEJXeAj24g8q6t38bGTEMcUvh3YGS+P 7DLIF2T2/f9oiO6LVD72zC/HkqaenzWla536G4FLAW4+MzAdn+R7hd5QN4CuFaKI nWWr5KMjZfgyxPkbe+WiD8QryR/EOd3za5vdycZRnBVIcbZQzEHwxizm3JyPNYy8 KeIjt5kGjFt9ccb1oPXjVibB2SUO322kbhMeCXKq2tswnZBICy+1EcbmaxNlkxwc hcdb2wz/3pzH0jbexgTFrNWeoJyYykM8s/3cMY/t0dKppve4FXLG9A==\n-----END CERTIFICATE-----",
						"key":         "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA8IPx6VXGPcmi68QKT+JL6xsMD26ajl8PTQ1qkg9prO0L4YvO HU52U/cci141tDGCPltd5HrOD65d+y/FRzmRb8Xc2Xc7h72TJCQ2kXHpCIlGKfuU has80xLZUKHQ2QS8tLkpcB5wx8plDSYNp15mvnDjsAQYb2/Yfo+b+LrlJEKB0cyF BwpHevC0eM2znvs64cVFOQGJvAu+2SaYQ14vRmJv5qGo6ZPtG4gg2f72Uo054GRL o0CmpnYEX1QoloKFK7nooo4dXCxn/lteQN5DaOVF30K4Wx1UUGNL8vXalNWQZrNK VnXPRnSRndpF5BjVrRG5ODQjenOuAFti9xtiTQIDAQABAoIBAQCkuln3bA3ox69U NuKxL9a7Ybzy3NfyZtz98xBolTHVhE083xn+LH0SqQ7dzVqO3dHMj5tRH2L+jnhD z8YYMC+SFDxcnTMilw6uFDdjilcGx65Mlsh0fIGeNyyr8wgtevcb+C2PYunvjImF Zei4FwnbqUnohgWOXVYz6Hv08Vx7ZdW+QiH62I/LS73G7d2EPb26Zo3zMKg/H5AD xuNk82MDW0lCgrw869Yqhcd3GkkmgWi+S71AE0ftY03QeBrsSZbzz4Zsgk2GsEBt fGclOu2c5sNRhLy4o7GiZghPS32zkiec9H76Ip5n/nwXgcYCoxfvOQL+b8U5vpap gbgjWGLBAoGBAPQsJGrYAzt1naGbzwEgYbbyItKD3Bf30YVcZxnZW1fZn2H9NpWl oIBAiO8ls1WP32Mf1us91Bp0V4ESmmklFb5ZllRZXnGt6U2pvxl2FZbK18Nvcvw/ kc66t643mKdfsTAF3sRSmeeWxNSB5C+/yZfHARbBoZ64hYqSut2sivjlAoGBAPwq dDBrz4P/PGn191ZfpWyTC4NRMpC2zYkHoWtMJNs96bsZub4phWyo1KLL4sbAZ+uV GfpRpE2u8mKQWEUU2Gz/KI0cY1e40Icg/ZBlzNglb7ssBXKMrI3R2pigIzEAArhU KwsjhreGF7ix5DEyNT4i7PQL/tOu3uh9SUXl9zVJAoGANta/KxvuxejpiUVUHZ2n NI53Ua55vQxUi04wfba6dCWVTU2wd7WmMYfM+WEPQPU6J6ob++N8AqEEkiGaemjw 1DqMr88OjhuQHXg1SkOiH6bZBLTAL3Ubi0GWRVOJPnYYdn+rA47FsCTFejDeDfdW EHeKgBDm+p3YqEHCJE0/PR0CgYA+mw+zweCIhgLqz81znVWFylAubyddtHT9E27p I8N2xz1TXYS3CLn+i0AXlwUbkUN7ws3rTv+65bd57xprNEyzavoXZrfnXJQxKGir xAqCk3DVCI3lrbVdlH9wKznxfW4vc34oSs60m88h5NChwjRj0+n+gUfoKF9hW1Go z/p7OQKBgCqfExvjjmBibsr0ZtTprMGV2qI55LoVnUXcM1xSBbgU0rLwvw2YBrWg MRl3ixYN851wGD1LhgpwFjr7SnwEhpKdDloKSE1ANM5LE7zvJHsPY2uvmY/Rbn7i RqIzeCbUVYt1Ow0WEqy0DUy/fGLQEz9viwLnvTcDNWeuSsljeAO7\n-----END RSA PRIVATE KEY-----",
						"cert_region": "cn-hangzhou",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap7935 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence7935(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_ddoscoo_instance" "defaultaRfzZ9" {
  normal_bandwidth = "100"
  normal_qps       = "500"
  bandwidth_mode   = "2"
  product_plan     = "3"
  product_type     = "ddosDip"
  period           = "1"
  port_count       = "5"
  name             = "测试手动上传证书"
  function_version = "0"
  domain_count     = "10"
}


`, name)
}

// Case 测试选择证书-Terraform接入 7937
func TestAccAliCloudDdosCooDomainResource_basic7937(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap7937)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence7937)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type":      "0",
					"ocsp_enabled": "false",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "websocket",
						},
					},
					"real_servers": []string{
						"1.1.1.1", "2.2.2.2", "3.3.3.3"},
					"domain": "testld.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultSJe7n8.id}", "${alicloud_ddoscoo_instance.default6lyurZ.id}", "${alicloud_ddoscoo_instance.defaultTTvY0D.id}"},
					"https_ext": "{\\\"Https2http\\\":1,\\\"Http2\\\":1,\\\"Http2https\\\":0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "0",
						"ocsp_enabled":   "false",
						"proxy_types.#":  "3",
						"real_servers.#": "3",
						"domain":         "testld.qq.com",
						"instance_ids.#": "3",
						"https_ext":      CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type":      "1",
					"ocsp_enabled": "true",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80", "8080"},
							"proxy_type": "http",
						},
					},
					"real_servers": []string{
						"1.qq.com"},
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultTTvY0D.id}"},
					"https_ext":       "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
					"cert_identifier": "${alicloud_ssl_certificates_service_certificate.default3MYZEt.id}",
					"cert_region":     "cn-hangzhou",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":         "1",
						"ocsp_enabled":    "true",
						"proxy_types.#":   "1",
						"real_servers.#":  "1",
						"instance_ids.#":  "1",
						"https_ext":       CHECKSET,
						"cert_identifier": CHECKSET,
						"cert_region":     "cn-hangzhou",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap7937 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence7937(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_ddoscoo_instance" "defaultTTvY0D" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "0"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ddoscoo_instance" "defaultSJe7n8" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test2"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "1"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ddoscoo_instance" "default6lyurZ" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test2"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "1"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ssl_certificates_service_certificate" "default3MYZEt" {
  cert             = <<EOF
-----BEGIN CERTIFICATE-----
MIID4TCCAsmgAwIBAgIRANZGvLwT8kuWpPlZ/Aj+uPgwDQYJKoZIhvcNAQELBQAw
XjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU
ZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w
HhcNMjQwODIzMDk0NzA0WhcNMjUwODIzMDk0NzA0WjAlMQswCQYDVQQGEwJDTjEW
MBQGA1UEAxMNdGVzdGxkLnFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBALmZY2geTFi+50gAVyDQH9Y5sTv8LLX6+MET1l3larzjX1M0Az9ZEIc0
TNrAp8mtJRlpQCzyDPZg88AwSdEwqSOSsnGzfS2DUcPJmdn2a2n5PLvWE28qPuSf
6fl3IhNiPzLYR51+7ccJKEQRhfOK2usmJo6oTG/0Lhh4BRH5owcclKv6n3YHaBVj
JNigiq1/tlqU46toZvotPOORjpy21kJPZioHqOVCDO4zreMy2xuIiYtpSSmXwkEO
zcQQ3K8sbRx9ED8SCdb229h7ioTug02YBXs0YOQZ024HFaIF8Nz1M+mdHy1jCbLd
yJoT/jzE4RdldZKZJFaSKV1c7EYlzhkCAwEAAaOB0jCBzzAOBgNVHQ8BAf8EBAMC
BaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMB8GA1UdIwQYMBaAFCiB
JgXRNBo/wXMPu5PPFRw/A79/MGMGCCsGAQUFBwEBBFcwVTAhBggrBgEFBQcwAYYV
aHR0cDovL29jc3AubXlzc2wuY29tMDAGCCsGAQUFBzAChiRodHRwOi8vY2EubXlz
c2wuY29tL215c3NsdGVzdHJzYS5jcnQwGAYDVR0RBBEwD4INdGVzdGxkLnFxLmNv
bTANBgkqhkiG9w0BAQsFAAOCAQEAnPJl1GrePDIulWfsETPbGnrZv3j3ZRXuou0o
K32X/nydS/i/j+AUzKSyezmnR1edkgY1hbGaza702SLQJuGh2IqJvAFyifwV/CZ5
cpJIi5G7kWTBjZo9NgVnDMhR8y5DCKE8BhiUBwcSvKKC8se2yWHm1fk9pRxG0Mc6
0fstl40jtR5XZYsW1GhX4fzwrWuBodPKticgXPn2e24ec+4rVrziu5R7D77AzJjG
Y/wzNYvAUWEzEya7Ve53nhu+WpIuIQn0ux8nPDioFdOjckn4jK3ePYdS2mWT6EBU
BC74GYiBNDz0QgHADq1VTExeLzC0tw9PPdWl0WfoTgCCKLz0yA==
-----END CERTIFICATE-----
EOF
  certificate_name = "173"
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAuZljaB5MWL7nSABXINAf1jmxO/wstfr4wRPWXeVqvONfUzQD
P1kQhzRM2sCnya0lGWlALPIM9mDzwDBJ0TCpI5KycbN9LYNRw8mZ2fZrafk8u9YT
byo+5J/p+XciE2I/MthHnX7txwkoRBGF84ra6yYmjqhMb/QuGHgFEfmjBxyUq/qf
dgdoFWMk2KCKrX+2WpTjq2hm+i0845GOnLbWQk9mKgeo5UIM7jOt4zLbG4iJi2lJ
KZfCQQ7NxBDcryxtHH0QPxIJ1vbb2HuKhO6DTZgFezRg5BnTbgcVogXw3PUz6Z0f
LWMJst3ImhP+PMThF2V1kpkkVpIpXVzsRiXOGQIDAQABAoIBAQCBPiw4A+k8X2vk
+r+xjNyurCwcTmXAL81rfmnnputmL5tg8DZWtanJzQS7zC7LRPQxttZGtiOKqkbz
DW1J6+3MZMo4XToNKIYWpduqKWvxNusxDkkoPy3evPEMlAY5o0/JE00DgrEHyfut
MtqplocN+tocu1vHFi3HQkSdmM4LE46ZfFu5w1FRbNI1Gqjj/cwlF/T93V3qMap1
WfsJjhMIX9LjBq3y9GAfAtAw7JYwkztr2AhYzCsK25wAj72zFY6FJTZ8LklfS41q
DrVtdjMx42IonDQtkzrqzfYlXdzzhZzuQHxn+qJODseoU8oDG9j3eKhp1dKgqLfx
tv1o3km1AoGBAOqXGEw2w94uVchCjuTum3XFYieEla0IUbHJCaWKU/hoSbht7j23
K7tA9//epBuRLGtYE0sPBK6i31mQT216muspO1g3pwGJhPy8VSpsJ1GQhB2G2UNz
kZlRK+2/gx35TdTi9x0C6UWk5XkhgWO3R35BlEnuV7EOyJunobiUcoObAoGBAMqJ
reuSbJajNGfeBzPel7F62ZDufC85hWaGKzeXIk3DkXcsEpeR5ogMGHCZ3dGf4Yz/
pcfjnCMIWjc+MkA4ppFd4432FJkxNQQP0z7njpXW3e5tionsMN+UwPzi4wZPKufK
osjw43JpBzpHGxG3ynLgZg7bSfrQZhPDTQw84nJbAoGBAKpy8mKeAB71R7rkMXNB
s48Uxca03RQGUWV+DxZKtcxt6fKpXUtWRd4ezJMLL+4fw0iTjCEjXmGNUf9/jVac
mOd44/erKBtD0m7YYIEcaE0pVfUmP8J0vDvL8MEkP56Nv/GIn8hijx/dOiaTI7JS
Pw4LlDVLikfJ2BTQ7f5xTes1AoGBAJ7h4HiDFgIZp1uvtfC/tjn5CEGEhBC7y+VA
bRify747I5rcDP2v66tf6bAzU+pExLhKN++Vov9sZvEdLmhoyGoSwBa2KzR9gHxe
ObYICjeLJfALKHnHuhM6ayY2iieB5UOOF6MQLSysLYpPC3IbvonddNJEvkUuRFVO
iNuHy5AvAoGAPqbzSNA05gf85zRO/JZAmZuWXG3o3pVougbkc211p+ynMpS+/bMb
c/nR36kOE551lFjIoAjoeIs16Wbq+00u9GlcQmyAdFpfaFCNHa3dayJwKJMW9Nia
fKbiiiOAQE8s2v8Paa+b00GspeWLow4u0G5lBVau4JjEVnl6ivLXlzY=
-----END RSA PRIVATE KEY-----
EOF
}


`, name)
}

// Case 测试选择证书-Terraform接入_副本1744168509262_副本1744960182068 10709
func TestAccAliCloudDdosCooDomainResource_basic10709(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap10709)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence10709)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type":      "0",
					"ocsp_enabled": "false",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "websocket",
						},
					},
					"real_servers": []string{
						"1.1.1.1", "2.2.2.2", "3.3.3.3"},
					"domain": "testld.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.default9nJ7Ie.id}"},
					"https_ext": "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "0",
						"ocsp_enabled":   "false",
						"proxy_types.#":  "3",
						"real_servers.#": "3",
						"domain":         "testld.qq.com",
						"instance_ids.#": "1",
						"https_ext":      CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80", "8080"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
					},
					"cert_identifier": "${alicloud_ssl_certificates_service_certificate.default3MYZEt.id}-cn-hangzhou",
					"cert_region":     "cn-hangzhou",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"proxy_types.#":   "2",
						"cert_identifier": CHECKSET,
						"cert_region":     "cn-hangzhou",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type": "1",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
					},
					"real_servers": []string{
						"1.qq.com"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "1",
						"proxy_types.#":  "2",
						"real_servers.#": "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap10709 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence10709(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "cert_region" {
  default = "default3MYZEt.CertId-${cn-hangzhou}"
}

resource "alicloud_ddoscoo_instance" "default9nJ7Ie" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "0"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ssl_certificates_service_certificate" "default3MYZEt" {
  cert             = <<EOF
-----BEGIN CERTIFICATE-----
MIID4TCCAsmgAwIBAgIRANZGvLwT8kuWpPlZ/Aj+uPgwDQYJKoZIhvcNAQELBQAw
XjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU
ZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w
HhcNMjQwODIzMDk0NzA0WhcNMjUwODIzMDk0NzA0WjAlMQswCQYDVQQGEwJDTjEW
MBQGA1UEAxMNdGVzdGxkLnFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBALmZY2geTFi+50gAVyDQH9Y5sTv8LLX6+MET1l3larzjX1M0Az9ZEIc0
TNrAp8mtJRlpQCzyDPZg88AwSdEwqSOSsnGzfS2DUcPJmdn2a2n5PLvWE28qPuSf
6fl3IhNiPzLYR51+7ccJKEQRhfOK2usmJo6oTG/0Lhh4BRH5owcclKv6n3YHaBVj
JNigiq1/tlqU46toZvotPOORjpy21kJPZioHqOVCDO4zreMy2xuIiYtpSSmXwkEO
zcQQ3K8sbRx9ED8SCdb229h7ioTug02YBXs0YOQZ024HFaIF8Nz1M+mdHy1jCbLd
yJoT/jzE4RdldZKZJFaSKV1c7EYlzhkCAwEAAaOB0jCBzzAOBgNVHQ8BAf8EBAMC
BaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMB8GA1UdIwQYMBaAFCiB
JgXRNBo/wXMPu5PPFRw/A79/MGMGCCsGAQUFBwEBBFcwVTAhBggrBgEFBQcwAYYV
aHR0cDovL29jc3AubXlzc2wuY29tMDAGCCsGAQUFBzAChiRodHRwOi8vY2EubXlz
c2wuY29tL215c3NsdGVzdHJzYS5jcnQwGAYDVR0RBBEwD4INdGVzdGxkLnFxLmNv
bTANBgkqhkiG9w0BAQsFAAOCAQEAnPJl1GrePDIulWfsETPbGnrZv3j3ZRXuou0o
K32X/nydS/i/j+AUzKSyezmnR1edkgY1hbGaza702SLQJuGh2IqJvAFyifwV/CZ5
cpJIi5G7kWTBjZo9NgVnDMhR8y5DCKE8BhiUBwcSvKKC8se2yWHm1fk9pRxG0Mc6
0fstl40jtR5XZYsW1GhX4fzwrWuBodPKticgXPn2e24ec+4rVrziu5R7D77AzJjG
Y/wzNYvAUWEzEya7Ve53nhu+WpIuIQn0ux8nPDioFdOjckn4jK3ePYdS2mWT6EBU
BC74GYiBNDz0QgHADq1VTExeLzC0tw9PPdWl0WfoTgCCKLz0yA==
-----END CERTIFICATE-----
EOF
  certificate_name = "622"
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAuZljaB5MWL7nSABXINAf1jmxO/wstfr4wRPWXeVqvONfUzQD
P1kQhzRM2sCnya0lGWlALPIM9mDzwDBJ0TCpI5KycbN9LYNRw8mZ2fZrafk8u9YT
byo+5J/p+XciE2I/MthHnX7txwkoRBGF84ra6yYmjqhMb/QuGHgFEfmjBxyUq/qf
dgdoFWMk2KCKrX+2WpTjq2hm+i0845GOnLbWQk9mKgeo5UIM7jOt4zLbG4iJi2lJ
KZfCQQ7NxBDcryxtHH0QPxIJ1vbb2HuKhO6DTZgFezRg5BnTbgcVogXw3PUz6Z0f
LWMJst3ImhP+PMThF2V1kpkkVpIpXVzsRiXOGQIDAQABAoIBAQCBPiw4A+k8X2vk
+r+xjNyurCwcTmXAL81rfmnnputmL5tg8DZWtanJzQS7zC7LRPQxttZGtiOKqkbz
DW1J6+3MZMo4XToNKIYWpduqKWvxNusxDkkoPy3evPEMlAY5o0/JE00DgrEHyfut
MtqplocN+tocu1vHFi3HQkSdmM4LE46ZfFu5w1FRbNI1Gqjj/cwlF/T93V3qMap1
WfsJjhMIX9LjBq3y9GAfAtAw7JYwkztr2AhYzCsK25wAj72zFY6FJTZ8LklfS41q
DrVtdjMx42IonDQtkzrqzfYlXdzzhZzuQHxn+qJODseoU8oDG9j3eKhp1dKgqLfx
tv1o3km1AoGBAOqXGEw2w94uVchCjuTum3XFYieEla0IUbHJCaWKU/hoSbht7j23
K7tA9//epBuRLGtYE0sPBK6i31mQT216muspO1g3pwGJhPy8VSpsJ1GQhB2G2UNz
kZlRK+2/gx35TdTi9x0C6UWk5XkhgWO3R35BlEnuV7EOyJunobiUcoObAoGBAMqJ
reuSbJajNGfeBzPel7F62ZDufC85hWaGKzeXIk3DkXcsEpeR5ogMGHCZ3dGf4Yz/
pcfjnCMIWjc+MkA4ppFd4432FJkxNQQP0z7njpXW3e5tionsMN+UwPzi4wZPKufK
osjw43JpBzpHGxG3ynLgZg7bSfrQZhPDTQw84nJbAoGBAKpy8mKeAB71R7rkMXNB
s48Uxca03RQGUWV+DxZKtcxt6fKpXUtWRd4ezJMLL+4fw0iTjCEjXmGNUf9/jVac
mOd44/erKBtD0m7YYIEcaE0pVfUmP8J0vDvL8MEkP56Nv/GIn8hijx/dOiaTI7JS
Pw4LlDVLikfJ2BTQ7f5xTes1AoGBAJ7h4HiDFgIZp1uvtfC/tjn5CEGEhBC7y+VA
bRify747I5rcDP2v66tf6bAzU+pExLhKN++Vov9sZvEdLmhoyGoSwBa2KzR9gHxe
ObYICjeLJfALKHnHuhM6ayY2iieB5UOOF6MQLSysLYpPC3IbvonddNJEvkUuRFVO
iNuHy5AvAoGAPqbzSNA05gf85zRO/JZAmZuWXG3o3pVougbkc211p+ynMpS+/bMb
c/nR36kOE551lFjIoAjoeIs16Wbq+00u9GlcQmyAdFpfaFCNHa3dayJwKJMW9Nia
fKbiiiOAQE8s2v8Paa+b00GspeWLow4u0G5lBVau4JjEVnl6ivLXlzY=
-----END RSA PRIVATE KEY-----
EOF
}


`, name)
}

// Case 手动上传证书_TF接入_副本1744960888974 10710
func TestAccAliCloudDdosCooDomainResource_basic10710(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap10710)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence10710)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"ap-southeast-1"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type": "0",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
					},
					"real_servers": []string{
						"1.1.1.1"},
					"domain": "testcert.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultaRfzZ9.id}"},
					"https_ext":   "{\\\"Https2http\\\":1,\\\"Http2\\\":1,\\\"Http2https\\\":0}",
					"cert_name":   "694",
					"cert":        "-----BEGIN CERTIFICATE-----\\nMIID5DCCAsygAwIBAgIQWen3GebvT0GcE/a1MJVFgjANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MjMwOTE1MjNaFw0yNTA4MjMwOTE1MjNaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDF/H2/Oas8trhkHs7B4B8mN00eup/5Tqar2QO4Hm499GJECKc3eMiC v+aAlW74Iymb2Varnv+WMdFRVMQgpXesi3akvVp0QxecvcDliilkh4ddTK731Rd7 PaSK1JdQX1jdGGhVnhQz+cPNFBGZ3tMYGhUkgNfqa3UFucJcBuRub/Ircr+5Ob4D FxSglfTHi+/EFcp7vMAOztLD4zXmEz3NysDNP6NzN7SD72DwPp0nxyRjrBlHSOVg szB/bFasQdAhZGeo64MvSb+SivdWEMhHkwKA5MhhYOkDeNPPSmlxbw0Z3nOyeMmI YkaxzhpO5DZN382duTQmiQ+Yg60OfL3NAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAKUCvcbUCJRFbVowd1YorILivqRmS6ztR 9vLdj4YZBWxmmQrgkDlkl78r/rXlJbqHunSh2Wbag7y+GaQQwg8xcL4Z3KrKj4zg aHEP1DyFiaxMTEuC/L2RgSX3xlXcf6fQ46D3y3Ja3iHFQnjx6npNaSZ2bSEULvJg IjPiJ/nbid1TYR5vtg8vtwDQfY7/+q8/3DWKcQq+SGcd9dDS6u9vulNdlW8e14bS CKDEzS/axjoICl9JagASLcElWIit/eD5zGIKzTPC9mEXiX/J/gUr70y9GiE7Ue++ 4nFSGQOwjMh/wO2HlmRfToeZ3g6rRCijibBHKHBmVym7NCai2voZJQ==\\n-----END CERTIFICATE-----",
					"key":         "-----BEGIN RSA PRIVATE KEY-----\\nMIIEowIBAAKCAQEAxfx9vzmrPLa4ZB7OweAfJjdNHrqf+U6mq9kDuB5uPfRiRAin N3jIgr/mgJVu+CMpm9lWq57/ljHRUVTEIKV3rIt2pL1adEMXnL3A5YopZIeHXUyu 99UXez2kitSXUF9Y3RhoVZ4UM/nDzRQRmd7TGBoVJIDX6mt1BbnCXAbkbm/yK3K/ uTm+AxcUoJX0x4vvxBXKe7zADs7Sw+M15hM9zcrAzT+jcze0g+9g8D6dJ8ckY6wZ R0jlYLMwf2xWrEHQIWRnqOuDL0m/kor3VhDIR5MCgOTIYWDpA3jTz0ppcW8NGd5z snjJiGJGsc4aTuQ2Td/Nnbk0JokPmIOtDny9zQIDAQABAoIBAQDD6fU4y8UhwCG4 mS+5c6D/PQvoU35Hwkd1l7pxcFNgpTqz3egyISgxEdny9WwoyQq8eJWmICEEK+nY VEv7jiFdMWhG3kTq9RUhejeuLEiHfQE7Fs2w2kFxJ29yHapZ0u/pYOSljFarlATo I2rDW1aB7BVt2L1P7+ONteKZFAzpJckft5ceRUzs5Jm1Cqt8OWO3Km+FBbCROv8M TevW44aoMwBGXuqs06FV1Z4dafglskjt2O38V4acZpH8Nc8j+nCONKL3OxwKY6HQ WfnbXnTLCF3IuMiy8ntrY8HYU6EABiCdr+Pl5HmhI2nmtSFTFbD4Gq70vgPL0P1m iULJGJ7hAoGBANQPrOGe9qHcBydvcBHE7qA9v1+IaTj03qzDTopTi/jxcd9pEkei skLyHNQ5yJT0QjTxB9iYRLfZccOGFyqz/Sdz6CwwTWBZeXOQ2AX7FPEcCnNr1TpF yMrgOY3H93KJISEVS6kYskByjK7XzXCp0KQNS2EeIhAXcqXxNmSwylvpAoGBAO8C PdZHd6aLLEZyVO1aZVHDxqmbhmGVoY9wZ+uwR4K2Hu/fjk0qlR9cYpw8+N675Wr9 E9Ff5/wjK2+/+uocQV9Zoap2vgrwX7GASuO5KYdCOBn6oUOSa+Ru+LgBNyUkXYES mM8eFC1QqfcSrETLAQqd2lmLcuaMq6jJtbBpvzhFAoGAYd8mNC9wtr1dE+dLuvfA BnbZJ1dG8QGa7/NoAVGT7X5JxwmwZR2C1oD1q0FMAOtGzzZbH60PMicKaWousQfH E/lbs2FLpOdGtX6pJQF/5dPCQwkGrVFd3bxk87nRy6vcfW9drxp10mbL5To2WAQY Bk8Ydic5I2IfCNVt/ETX8FkCgYA+OkAtVQgi9WM+qC/SaFGu2yETManoKFQbC3IT HB9SOeaOH49mKesPcjc+ZGWLYDJYC7IoNicpL2L0wnAqmdavY5/CyQ2rvW+8wCE/ bwsP6z6+DNIFzM6IeBgLmE1qPzCVFWlxq2wnbDQEXvk5I/2ObRDXdYYh3ogm9vV2 C+I8XQKBgHIquvifRVvWf1q9WFZLQXZMv1flPhNaLmR+2k6gNpJ8SeiOmBtE7gT6 Je+YOXEKvfr6jaaJwYHPi6IhWHs4fQbgdK4jei30sRL7c8QdKEuwRdHPimmGNAPb UapzHY7xq0Wk9enAnM/SXkjTAJEkrpiQiDuPZVi4sIYCOqb+Ovu5\\n-----END RSA PRIVATE KEY-----",
					"cert_region": "ap-southeast-1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "0",
						"proxy_types.#":  "2",
						"real_servers.#": "1",
						"domain":         "testcert.qq.com",
						"instance_ids.#": "1",
						"https_ext":      CHECKSET,
						"cert_name":      CHECKSET,
						"cert":           "-----BEGIN CERTIFICATE-----\nMIID5DCCAsygAwIBAgIQWen3GebvT0GcE/a1MJVFgjANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MjMwOTE1MjNaFw0yNTA4MjMwOTE1MjNaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDF/H2/Oas8trhkHs7B4B8mN00eup/5Tqar2QO4Hm499GJECKc3eMiC v+aAlW74Iymb2Varnv+WMdFRVMQgpXesi3akvVp0QxecvcDliilkh4ddTK731Rd7 PaSK1JdQX1jdGGhVnhQz+cPNFBGZ3tMYGhUkgNfqa3UFucJcBuRub/Ircr+5Ob4D FxSglfTHi+/EFcp7vMAOztLD4zXmEz3NysDNP6NzN7SD72DwPp0nxyRjrBlHSOVg szB/bFasQdAhZGeo64MvSb+SivdWEMhHkwKA5MhhYOkDeNPPSmlxbw0Z3nOyeMmI YkaxzhpO5DZN382duTQmiQ+Yg60OfL3NAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAKUCvcbUCJRFbVowd1YorILivqRmS6ztR 9vLdj4YZBWxmmQrgkDlkl78r/rXlJbqHunSh2Wbag7y+GaQQwg8xcL4Z3KrKj4zg aHEP1DyFiaxMTEuC/L2RgSX3xlXcf6fQ46D3y3Ja3iHFQnjx6npNaSZ2bSEULvJg IjPiJ/nbid1TYR5vtg8vtwDQfY7/+q8/3DWKcQq+SGcd9dDS6u9vulNdlW8e14bS CKDEzS/axjoICl9JagASLcElWIit/eD5zGIKzTPC9mEXiX/J/gUr70y9GiE7Ue++ 4nFSGQOwjMh/wO2HlmRfToeZ3g6rRCijibBHKHBmVym7NCai2voZJQ==\n-----END CERTIFICATE-----",
						"key":            "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAxfx9vzmrPLa4ZB7OweAfJjdNHrqf+U6mq9kDuB5uPfRiRAin N3jIgr/mgJVu+CMpm9lWq57/ljHRUVTEIKV3rIt2pL1adEMXnL3A5YopZIeHXUyu 99UXez2kitSXUF9Y3RhoVZ4UM/nDzRQRmd7TGBoVJIDX6mt1BbnCXAbkbm/yK3K/ uTm+AxcUoJX0x4vvxBXKe7zADs7Sw+M15hM9zcrAzT+jcze0g+9g8D6dJ8ckY6wZ R0jlYLMwf2xWrEHQIWRnqOuDL0m/kor3VhDIR5MCgOTIYWDpA3jTz0ppcW8NGd5z snjJiGJGsc4aTuQ2Td/Nnbk0JokPmIOtDny9zQIDAQABAoIBAQDD6fU4y8UhwCG4 mS+5c6D/PQvoU35Hwkd1l7pxcFNgpTqz3egyISgxEdny9WwoyQq8eJWmICEEK+nY VEv7jiFdMWhG3kTq9RUhejeuLEiHfQE7Fs2w2kFxJ29yHapZ0u/pYOSljFarlATo I2rDW1aB7BVt2L1P7+ONteKZFAzpJckft5ceRUzs5Jm1Cqt8OWO3Km+FBbCROv8M TevW44aoMwBGXuqs06FV1Z4dafglskjt2O38V4acZpH8Nc8j+nCONKL3OxwKY6HQ WfnbXnTLCF3IuMiy8ntrY8HYU6EABiCdr+Pl5HmhI2nmtSFTFbD4Gq70vgPL0P1m iULJGJ7hAoGBANQPrOGe9qHcBydvcBHE7qA9v1+IaTj03qzDTopTi/jxcd9pEkei skLyHNQ5yJT0QjTxB9iYRLfZccOGFyqz/Sdz6CwwTWBZeXOQ2AX7FPEcCnNr1TpF yMrgOY3H93KJISEVS6kYskByjK7XzXCp0KQNS2EeIhAXcqXxNmSwylvpAoGBAO8C PdZHd6aLLEZyVO1aZVHDxqmbhmGVoY9wZ+uwR4K2Hu/fjk0qlR9cYpw8+N675Wr9 E9Ff5/wjK2+/+uocQV9Zoap2vgrwX7GASuO5KYdCOBn6oUOSa+Ru+LgBNyUkXYES mM8eFC1QqfcSrETLAQqd2lmLcuaMq6jJtbBpvzhFAoGAYd8mNC9wtr1dE+dLuvfA BnbZJ1dG8QGa7/NoAVGT7X5JxwmwZR2C1oD1q0FMAOtGzzZbH60PMicKaWousQfH E/lbs2FLpOdGtX6pJQF/5dPCQwkGrVFd3bxk87nRy6vcfW9drxp10mbL5To2WAQY Bk8Ydic5I2IfCNVt/ETX8FkCgYA+OkAtVQgi9WM+qC/SaFGu2yETManoKFQbC3IT HB9SOeaOH49mKesPcjc+ZGWLYDJYC7IoNicpL2L0wnAqmdavY5/CyQ2rvW+8wCE/ bwsP6z6+DNIFzM6IeBgLmE1qPzCVFWlxq2wnbDQEXvk5I/2ObRDXdYYh3ogm9vV2 C+I8XQKBgHIquvifRVvWf1q9WFZLQXZMv1flPhNaLmR+2k6gNpJ8SeiOmBtE7gT6 Je+YOXEKvfr6jaaJwYHPi6IhWHs4fQbgdK4jei30sRL7c8QdKEuwRdHPimmGNAPb UapzHY7xq0Wk9enAnM/SXkjTAJEkrpiQiDuPZVi4sIYCOqb+Ovu5\n-----END RSA PRIVATE KEY-----",
						"cert_region":    "ap-southeast-1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"real_servers": []string{
						"1.1.1.1", "2.2.2.2", "3.3.3.3"},
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultaRfzZ9.id}", "${alicloud_ddoscoo_instance.defaultQPviUE.id}", "${alicloud_ddoscoo_instance.defaultgyrXgo.id}"},
					"https_ext":    "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
					"cert_name":    "869",
					"cert":         "-----BEGIN CERTIFICATE-----\\nMIID5DCCAsygAwIBAgIQFO5STIOlR/KkRB3gDHsi5zANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MTMwNTEyMDdaFw0yNTA4MTMwNTEyMDdaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDwg/HpVcY9yaLrxApP4kvrGwwPbpqOXw9NDWqSD2ms7Qvhi84dTnZT 9xyLXjW0MYI+W13kes4Prl37L8VHOZFvxdzZdzuHvZMkJDaRcekIiUYp+5SFqzzT EtlQodDZBLy0uSlwHnDHymUNJg2nXma+cOOwBBhvb9h+j5v4uuUkQoHRzIUHCkd6 8LR4zbOe+zrhxUU5AYm8C77ZJphDXi9GYm/moajpk+0biCDZ/vZSjTngZEujQKam dgRfVCiWgoUrueiijh1cLGf+W15A3kNo5UXfQrhbHVRQY0vy9dqU1ZBms0pWdc9G dJGd2kXkGNWtEbk4NCN6c64AW2L3G2JNAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAFl4tmH+m/lGr7aKxxz5uYdqepcjCHTzW KpxxEHr+sL9fOkap6I/NA9MYKsuVfnrTMnEJXeAj24g8q6t38bGTEMcUvh3YGS+P 7DLIF2T2/f9oiO6LVD72zC/HkqaenzWla536G4FLAW4+MzAdn+R7hd5QN4CuFaKI nWWr5KMjZfgyxPkbe+WiD8QryR/EOd3za5vdycZRnBVIcbZQzEHwxizm3JyPNYy8 KeIjt5kGjFt9ccb1oPXjVibB2SUO322kbhMeCXKq2tswnZBICy+1EcbmaxNlkxwc hcdb2wz/3pzH0jbexgTFrNWeoJyYykM8s/3cMY/t0dKppve4FXLG9A==\\n-----END CERTIFICATE-----",
					"key":          "-----BEGIN RSA PRIVATE KEY-----\\nMIIEowIBAAKCAQEA8IPx6VXGPcmi68QKT+JL6xsMD26ajl8PTQ1qkg9prO0L4YvO HU52U/cci141tDGCPltd5HrOD65d+y/FRzmRb8Xc2Xc7h72TJCQ2kXHpCIlGKfuU has80xLZUKHQ2QS8tLkpcB5wx8plDSYNp15mvnDjsAQYb2/Yfo+b+LrlJEKB0cyF BwpHevC0eM2znvs64cVFOQGJvAu+2SaYQ14vRmJv5qGo6ZPtG4gg2f72Uo054GRL o0CmpnYEX1QoloKFK7nooo4dXCxn/lteQN5DaOVF30K4Wx1UUGNL8vXalNWQZrNK VnXPRnSRndpF5BjVrRG5ODQjenOuAFti9xtiTQIDAQABAoIBAQCkuln3bA3ox69U NuKxL9a7Ybzy3NfyZtz98xBolTHVhE083xn+LH0SqQ7dzVqO3dHMj5tRH2L+jnhD z8YYMC+SFDxcnTMilw6uFDdjilcGx65Mlsh0fIGeNyyr8wgtevcb+C2PYunvjImF Zei4FwnbqUnohgWOXVYz6Hv08Vx7ZdW+QiH62I/LS73G7d2EPb26Zo3zMKg/H5AD xuNk82MDW0lCgrw869Yqhcd3GkkmgWi+S71AE0ftY03QeBrsSZbzz4Zsgk2GsEBt fGclOu2c5sNRhLy4o7GiZghPS32zkiec9H76Ip5n/nwXgcYCoxfvOQL+b8U5vpap gbgjWGLBAoGBAPQsJGrYAzt1naGbzwEgYbbyItKD3Bf30YVcZxnZW1fZn2H9NpWl oIBAiO8ls1WP32Mf1us91Bp0V4ESmmklFb5ZllRZXnGt6U2pvxl2FZbK18Nvcvw/ kc66t643mKdfsTAF3sRSmeeWxNSB5C+/yZfHARbBoZ64hYqSut2sivjlAoGBAPwq dDBrz4P/PGn191ZfpWyTC4NRMpC2zYkHoWtMJNs96bsZub4phWyo1KLL4sbAZ+uV GfpRpE2u8mKQWEUU2Gz/KI0cY1e40Icg/ZBlzNglb7ssBXKMrI3R2pigIzEAArhU KwsjhreGF7ix5DEyNT4i7PQL/tOu3uh9SUXl9zVJAoGANta/KxvuxejpiUVUHZ2n NI53Ua55vQxUi04wfba6dCWVTU2wd7WmMYfM+WEPQPU6J6ob++N8AqEEkiGaemjw 1DqMr88OjhuQHXg1SkOiH6bZBLTAL3Ubi0GWRVOJPnYYdn+rA47FsCTFejDeDfdW EHeKgBDm+p3YqEHCJE0/PR0CgYA+mw+zweCIhgLqz81znVWFylAubyddtHT9E27p I8N2xz1TXYS3CLn+i0AXlwUbkUN7ws3rTv+65bd57xprNEyzavoXZrfnXJQxKGir xAqCk3DVCI3lrbVdlH9wKznxfW4vc34oSs60m88h5NChwjRj0+n+gUfoKF9hW1Go z/p7OQKBgCqfExvjjmBibsr0ZtTprMGV2qI55LoVnUXcM1xSBbgU0rLwvw2YBrWg MRl3ixYN851wGD1LhgpwFjr7SnwEhpKdDloKSE1ANM5LE7zvJHsPY2uvmY/Rbn7i RqIzeCbUVYt1Ow0WEqy0DUy/fGLQEz9viwLnvTcDNWeuSsljeAO7\\n-----END RSA PRIVATE KEY-----",
					"cert_region":  "cn-hangzhou",
					"ocsp_enabled": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"real_servers.#": "3",
						"instance_ids.#": "3",
						"https_ext":      CHECKSET,
						"cert_name":      CHECKSET,
						"cert":           "-----BEGIN CERTIFICATE-----\nMIID5DCCAsygAwIBAgIQFO5STIOlR/KkRB3gDHsi5zANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MTMwNTEyMDdaFw0yNTA4MTMwNTEyMDdaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDwg/HpVcY9yaLrxApP4kvrGwwPbpqOXw9NDWqSD2ms7Qvhi84dTnZT 9xyLXjW0MYI+W13kes4Prl37L8VHOZFvxdzZdzuHvZMkJDaRcekIiUYp+5SFqzzT EtlQodDZBLy0uSlwHnDHymUNJg2nXma+cOOwBBhvb9h+j5v4uuUkQoHRzIUHCkd6 8LR4zbOe+zrhxUU5AYm8C77ZJphDXi9GYm/moajpk+0biCDZ/vZSjTngZEujQKam dgRfVCiWgoUrueiijh1cLGf+W15A3kNo5UXfQrhbHVRQY0vy9dqU1ZBms0pWdc9G dJGd2kXkGNWtEbk4NCN6c64AW2L3G2JNAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAFl4tmH+m/lGr7aKxxz5uYdqepcjCHTzW KpxxEHr+sL9fOkap6I/NA9MYKsuVfnrTMnEJXeAj24g8q6t38bGTEMcUvh3YGS+P 7DLIF2T2/f9oiO6LVD72zC/HkqaenzWla536G4FLAW4+MzAdn+R7hd5QN4CuFaKI nWWr5KMjZfgyxPkbe+WiD8QryR/EOd3za5vdycZRnBVIcbZQzEHwxizm3JyPNYy8 KeIjt5kGjFt9ccb1oPXjVibB2SUO322kbhMeCXKq2tswnZBICy+1EcbmaxNlkxwc hcdb2wz/3pzH0jbexgTFrNWeoJyYykM8s/3cMY/t0dKppve4FXLG9A==\n-----END CERTIFICATE-----",
						"key":            "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA8IPx6VXGPcmi68QKT+JL6xsMD26ajl8PTQ1qkg9prO0L4YvO HU52U/cci141tDGCPltd5HrOD65d+y/FRzmRb8Xc2Xc7h72TJCQ2kXHpCIlGKfuU has80xLZUKHQ2QS8tLkpcB5wx8plDSYNp15mvnDjsAQYb2/Yfo+b+LrlJEKB0cyF BwpHevC0eM2znvs64cVFOQGJvAu+2SaYQ14vRmJv5qGo6ZPtG4gg2f72Uo054GRL o0CmpnYEX1QoloKFK7nooo4dXCxn/lteQN5DaOVF30K4Wx1UUGNL8vXalNWQZrNK VnXPRnSRndpF5BjVrRG5ODQjenOuAFti9xtiTQIDAQABAoIBAQCkuln3bA3ox69U NuKxL9a7Ybzy3NfyZtz98xBolTHVhE083xn+LH0SqQ7dzVqO3dHMj5tRH2L+jnhD z8YYMC+SFDxcnTMilw6uFDdjilcGx65Mlsh0fIGeNyyr8wgtevcb+C2PYunvjImF Zei4FwnbqUnohgWOXVYz6Hv08Vx7ZdW+QiH62I/LS73G7d2EPb26Zo3zMKg/H5AD xuNk82MDW0lCgrw869Yqhcd3GkkmgWi+S71AE0ftY03QeBrsSZbzz4Zsgk2GsEBt fGclOu2c5sNRhLy4o7GiZghPS32zkiec9H76Ip5n/nwXgcYCoxfvOQL+b8U5vpap gbgjWGLBAoGBAPQsJGrYAzt1naGbzwEgYbbyItKD3Bf30YVcZxnZW1fZn2H9NpWl oIBAiO8ls1WP32Mf1us91Bp0V4ESmmklFb5ZllRZXnGt6U2pvxl2FZbK18Nvcvw/ kc66t643mKdfsTAF3sRSmeeWxNSB5C+/yZfHARbBoZ64hYqSut2sivjlAoGBAPwq dDBrz4P/PGn191ZfpWyTC4NRMpC2zYkHoWtMJNs96bsZub4phWyo1KLL4sbAZ+uV GfpRpE2u8mKQWEUU2Gz/KI0cY1e40Icg/ZBlzNglb7ssBXKMrI3R2pigIzEAArhU KwsjhreGF7ix5DEyNT4i7PQL/tOu3uh9SUXl9zVJAoGANta/KxvuxejpiUVUHZ2n NI53Ua55vQxUi04wfba6dCWVTU2wd7WmMYfM+WEPQPU6J6ob++N8AqEEkiGaemjw 1DqMr88OjhuQHXg1SkOiH6bZBLTAL3Ubi0GWRVOJPnYYdn+rA47FsCTFejDeDfdW EHeKgBDm+p3YqEHCJE0/PR0CgYA+mw+zweCIhgLqz81znVWFylAubyddtHT9E27p I8N2xz1TXYS3CLn+i0AXlwUbkUN7ws3rTv+65bd57xprNEyzavoXZrfnXJQxKGir xAqCk3DVCI3lrbVdlH9wKznxfW4vc34oSs60m88h5NChwjRj0+n+gUfoKF9hW1Go z/p7OQKBgCqfExvjjmBibsr0ZtTprMGV2qI55LoVnUXcM1xSBbgU0rLwvw2YBrWg MRl3ixYN851wGD1LhgpwFjr7SnwEhpKdDloKSE1ANM5LE7zvJHsPY2uvmY/Rbn7i RqIzeCbUVYt1Ow0WEqy0DUy/fGLQEz9viwLnvTcDNWeuSsljeAO7\n-----END RSA PRIVATE KEY-----",
						"cert_region":    "cn-hangzhou",
						"ocsp_enabled":   "true",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap10710 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence10710(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_ddoscoo_instance" "defaultaRfzZ9" {
  normal_bandwidth = "100"
  normal_qps       = "500"
  bandwidth_mode   = "2"
  product_plan     = "3"
  product_type     = "ddosDip"
  period           = "1"
  port_count       = "5"
  name             = "测试手动上传证书"
  function_version = "0"
  domain_count     = "10"
}

resource "alicloud_ddoscoo_instance" "defaultgyrXgo" {
  normal_bandwidth = "100"
  normal_qps       = "500"
  bandwidth_mode   = "2"
  product_plan     = "3"
  product_type     = "ddosDip"
  period           = "1"
  port_count       = "5"
  name             = "测试手动上传证书"
  function_version = "0"
  domain_count     = "10"
}

resource "alicloud_ddoscoo_instance" "defaultQPviUE" {
  normal_bandwidth = "100"
  normal_qps       = "500"
  bandwidth_mode   = "2"
  product_plan     = "3"
  product_type     = "ddosDip"
  period           = "1"
  port_count       = "5"
  name             = "测试手动上传证书"
  function_version = "0"
  domain_count     = "10"
}


`, name)
}

// Case 测试选择证书-Terraform接入证书ID问题修复 10784
func TestAccAliCloudDdosCooDomainResource_basic10784(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap10784)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence10784)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type":      "0",
					"ocsp_enabled": "false",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "websocket",
						},
					},
					"real_servers": []string{
						"1.1.1.1", "2.2.2.2", "3.3.3.3"},
					"domain": "testld.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.default9nJ7Ie.id}"},
					"https_ext": "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "0",
						"ocsp_enabled":   "false",
						"proxy_types.#":  "3",
						"real_servers.#": "3",
						"domain":         "testld.qq.com",
						"instance_ids.#": "1",
						"https_ext":      CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80", "8080"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "websockets",
						},
					},
					"cert_identifier": "${alicloud_ssl_certificates_service_certificate.default3MYZEt.id}-cn-hangzhou",
					"cert_region":     "cn-hangzhou",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"proxy_types.#":   "3",
						"cert_identifier": CHECKSET,
						"cert_region":     "cn-hangzhou",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type": "1",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
					},
					"real_servers": []string{
						"1.qq.com"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "1",
						"proxy_types.#":  "2",
						"real_servers.#": "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap10784 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence10784(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "cert_region" {
  default = "default3MYZEt.CertId-${cn-hangzhou}"
}

resource "alicloud_ddoscoo_instance" "default9nJ7Ie" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "0"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ssl_certificates_service_certificate" "default3MYZEt" {
  cert             = <<EOF
-----BEGIN CERTIFICATE-----
MIID4TCCAsmgAwIBAgIRANZGvLwT8kuWpPlZ/Aj+uPgwDQYJKoZIhvcNAQELBQAw
XjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU
ZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w
HhcNMjQwODIzMDk0NzA0WhcNMjUwODIzMDk0NzA0WjAlMQswCQYDVQQGEwJDTjEW
MBQGA1UEAxMNdGVzdGxkLnFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBALmZY2geTFi+50gAVyDQH9Y5sTv8LLX6+MET1l3larzjX1M0Az9ZEIc0
TNrAp8mtJRlpQCzyDPZg88AwSdEwqSOSsnGzfS2DUcPJmdn2a2n5PLvWE28qPuSf
6fl3IhNiPzLYR51+7ccJKEQRhfOK2usmJo6oTG/0Lhh4BRH5owcclKv6n3YHaBVj
JNigiq1/tlqU46toZvotPOORjpy21kJPZioHqOVCDO4zreMy2xuIiYtpSSmXwkEO
zcQQ3K8sbRx9ED8SCdb229h7ioTug02YBXs0YOQZ024HFaIF8Nz1M+mdHy1jCbLd
yJoT/jzE4RdldZKZJFaSKV1c7EYlzhkCAwEAAaOB0jCBzzAOBgNVHQ8BAf8EBAMC
BaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMB8GA1UdIwQYMBaAFCiB
JgXRNBo/wXMPu5PPFRw/A79/MGMGCCsGAQUFBwEBBFcwVTAhBggrBgEFBQcwAYYV
aHR0cDovL29jc3AubXlzc2wuY29tMDAGCCsGAQUFBzAChiRodHRwOi8vY2EubXlz
c2wuY29tL215c3NsdGVzdHJzYS5jcnQwGAYDVR0RBBEwD4INdGVzdGxkLnFxLmNv
bTANBgkqhkiG9w0BAQsFAAOCAQEAnPJl1GrePDIulWfsETPbGnrZv3j3ZRXuou0o
K32X/nydS/i/j+AUzKSyezmnR1edkgY1hbGaza702SLQJuGh2IqJvAFyifwV/CZ5
cpJIi5G7kWTBjZo9NgVnDMhR8y5DCKE8BhiUBwcSvKKC8se2yWHm1fk9pRxG0Mc6
0fstl40jtR5XZYsW1GhX4fzwrWuBodPKticgXPn2e24ec+4rVrziu5R7D77AzJjG
Y/wzNYvAUWEzEya7Ve53nhu+WpIuIQn0ux8nPDioFdOjckn4jK3ePYdS2mWT6EBU
BC74GYiBNDz0QgHADq1VTExeLzC0tw9PPdWl0WfoTgCCKLz0yA==
-----END CERTIFICATE-----
EOF
  certificate_name = "868"
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAuZljaB5MWL7nSABXINAf1jmxO/wstfr4wRPWXeVqvONfUzQD
P1kQhzRM2sCnya0lGWlALPIM9mDzwDBJ0TCpI5KycbN9LYNRw8mZ2fZrafk8u9YT
byo+5J/p+XciE2I/MthHnX7txwkoRBGF84ra6yYmjqhMb/QuGHgFEfmjBxyUq/qf
dgdoFWMk2KCKrX+2WpTjq2hm+i0845GOnLbWQk9mKgeo5UIM7jOt4zLbG4iJi2lJ
KZfCQQ7NxBDcryxtHH0QPxIJ1vbb2HuKhO6DTZgFezRg5BnTbgcVogXw3PUz6Z0f
LWMJst3ImhP+PMThF2V1kpkkVpIpXVzsRiXOGQIDAQABAoIBAQCBPiw4A+k8X2vk
+r+xjNyurCwcTmXAL81rfmnnputmL5tg8DZWtanJzQS7zC7LRPQxttZGtiOKqkbz
DW1J6+3MZMo4XToNKIYWpduqKWvxNusxDkkoPy3evPEMlAY5o0/JE00DgrEHyfut
MtqplocN+tocu1vHFi3HQkSdmM4LE46ZfFu5w1FRbNI1Gqjj/cwlF/T93V3qMap1
WfsJjhMIX9LjBq3y9GAfAtAw7JYwkztr2AhYzCsK25wAj72zFY6FJTZ8LklfS41q
DrVtdjMx42IonDQtkzrqzfYlXdzzhZzuQHxn+qJODseoU8oDG9j3eKhp1dKgqLfx
tv1o3km1AoGBAOqXGEw2w94uVchCjuTum3XFYieEla0IUbHJCaWKU/hoSbht7j23
K7tA9//epBuRLGtYE0sPBK6i31mQT216muspO1g3pwGJhPy8VSpsJ1GQhB2G2UNz
kZlRK+2/gx35TdTi9x0C6UWk5XkhgWO3R35BlEnuV7EOyJunobiUcoObAoGBAMqJ
reuSbJajNGfeBzPel7F62ZDufC85hWaGKzeXIk3DkXcsEpeR5ogMGHCZ3dGf4Yz/
pcfjnCMIWjc+MkA4ppFd4432FJkxNQQP0z7njpXW3e5tionsMN+UwPzi4wZPKufK
osjw43JpBzpHGxG3ynLgZg7bSfrQZhPDTQw84nJbAoGBAKpy8mKeAB71R7rkMXNB
s48Uxca03RQGUWV+DxZKtcxt6fKpXUtWRd4ezJMLL+4fw0iTjCEjXmGNUf9/jVac
mOd44/erKBtD0m7YYIEcaE0pVfUmP8J0vDvL8MEkP56Nv/GIn8hijx/dOiaTI7JS
Pw4LlDVLikfJ2BTQ7f5xTes1AoGBAJ7h4HiDFgIZp1uvtfC/tjn5CEGEhBC7y+VA
bRify747I5rcDP2v66tf6bAzU+pExLhKN++Vov9sZvEdLmhoyGoSwBa2KzR9gHxe
ObYICjeLJfALKHnHuhM6ayY2iieB5UOOF6MQLSysLYpPC3IbvonddNJEvkUuRFVO
iNuHy5AvAoGAPqbzSNA05gf85zRO/JZAmZuWXG3o3pVougbkc211p+ynMpS+/bMb
c/nR36kOE551lFjIoAjoeIs16Wbq+00u9GlcQmyAdFpfaFCNHa3dayJwKJMW9Nia
fKbiiiOAQE8s2v8Paa+b00GspeWLow4u0G5lBVau4JjEVnl6ivLXlzY=
-----END RSA PRIVATE KEY-----
EOF
}


`, name)
}

// Case 测试选择证书-Terraform接入证书ID问题修复_副本1747879542394 10823
func TestAccAliCloudDdosCooDomainResource_basic10823(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap10823)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence10823)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type":      "0",
					"ocsp_enabled": "false",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "websocket",
						},
					},
					"real_servers": []string{
						"1.1.1.1", "2.2.2.2", "3.3.3.3"},
					"domain": "testld.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.default9nJ7Ie.id}"},
					"https_ext":       "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
					"cert_identifier": "${alicloud_ssl_certificates_service_certificate.default3MYZEt.id}-cn-hangzhou",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":         "0",
						"ocsp_enabled":    "false",
						"proxy_types.#":   "3",
						"real_servers.#":  "3",
						"domain":          "testld.qq.com",
						"instance_ids.#":  "1",
						"https_ext":       CHECKSET,
						"cert_identifier": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80", "8080"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "websockets",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"proxy_types.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type": "1",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
					},
					"real_servers": []string{
						"1.qq.com"},
					"cert_identifier": "${alicloud_ssl_certificates_service_certificate.defaultAeCL0Q.id}-cn-hangzhou",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":         "1",
						"proxy_types.#":   "2",
						"real_servers.#":  "1",
						"cert_identifier": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap10823 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence10823(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "cert_region" {
  default = "default3MYZEt.CertId-${cn-hangzhou}"
}

resource "alicloud_ddoscoo_instance" "default9nJ7Ie" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "0"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ssl_certificates_service_certificate" "default3MYZEt" {
  cert             = <<EOF
-----BEGIN CERTIFICATE-----
MIID4TCCAsmgAwIBAgIRANZGvLwT8kuWpPlZ/Aj+uPgwDQYJKoZIhvcNAQELBQAw
XjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU
ZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w
HhcNMjQwODIzMDk0NzA0WhcNMjUwODIzMDk0NzA0WjAlMQswCQYDVQQGEwJDTjEW
MBQGA1UEAxMNdGVzdGxkLnFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBALmZY2geTFi+50gAVyDQH9Y5sTv8LLX6+MET1l3larzjX1M0Az9ZEIc0
TNrAp8mtJRlpQCzyDPZg88AwSdEwqSOSsnGzfS2DUcPJmdn2a2n5PLvWE28qPuSf
6fl3IhNiPzLYR51+7ccJKEQRhfOK2usmJo6oTG/0Lhh4BRH5owcclKv6n3YHaBVj
JNigiq1/tlqU46toZvotPOORjpy21kJPZioHqOVCDO4zreMy2xuIiYtpSSmXwkEO
zcQQ3K8sbRx9ED8SCdb229h7ioTug02YBXs0YOQZ024HFaIF8Nz1M+mdHy1jCbLd
yJoT/jzE4RdldZKZJFaSKV1c7EYlzhkCAwEAAaOB0jCBzzAOBgNVHQ8BAf8EBAMC
BaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMB8GA1UdIwQYMBaAFCiB
JgXRNBo/wXMPu5PPFRw/A79/MGMGCCsGAQUFBwEBBFcwVTAhBggrBgEFBQcwAYYV
aHR0cDovL29jc3AubXlzc2wuY29tMDAGCCsGAQUFBzAChiRodHRwOi8vY2EubXlz
c2wuY29tL215c3NsdGVzdHJzYS5jcnQwGAYDVR0RBBEwD4INdGVzdGxkLnFxLmNv
bTANBgkqhkiG9w0BAQsFAAOCAQEAnPJl1GrePDIulWfsETPbGnrZv3j3ZRXuou0o
K32X/nydS/i/j+AUzKSyezmnR1edkgY1hbGaza702SLQJuGh2IqJvAFyifwV/CZ5
cpJIi5G7kWTBjZo9NgVnDMhR8y5DCKE8BhiUBwcSvKKC8se2yWHm1fk9pRxG0Mc6
0fstl40jtR5XZYsW1GhX4fzwrWuBodPKticgXPn2e24ec+4rVrziu5R7D77AzJjG
Y/wzNYvAUWEzEya7Ve53nhu+WpIuIQn0ux8nPDioFdOjckn4jK3ePYdS2mWT6EBU
BC74GYiBNDz0QgHADq1VTExeLzC0tw9PPdWl0WfoTgCCKLz0yA==
-----END CERTIFICATE-----
EOF
  certificate_name = "587"
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAuZljaB5MWL7nSABXINAf1jmxO/wstfr4wRPWXeVqvONfUzQD
P1kQhzRM2sCnya0lGWlALPIM9mDzwDBJ0TCpI5KycbN9LYNRw8mZ2fZrafk8u9YT
byo+5J/p+XciE2I/MthHnX7txwkoRBGF84ra6yYmjqhMb/QuGHgFEfmjBxyUq/qf
dgdoFWMk2KCKrX+2WpTjq2hm+i0845GOnLbWQk9mKgeo5UIM7jOt4zLbG4iJi2lJ
KZfCQQ7NxBDcryxtHH0QPxIJ1vbb2HuKhO6DTZgFezRg5BnTbgcVogXw3PUz6Z0f
LWMJst3ImhP+PMThF2V1kpkkVpIpXVzsRiXOGQIDAQABAoIBAQCBPiw4A+k8X2vk
+r+xjNyurCwcTmXAL81rfmnnputmL5tg8DZWtanJzQS7zC7LRPQxttZGtiOKqkbz
DW1J6+3MZMo4XToNKIYWpduqKWvxNusxDkkoPy3evPEMlAY5o0/JE00DgrEHyfut
MtqplocN+tocu1vHFi3HQkSdmM4LE46ZfFu5w1FRbNI1Gqjj/cwlF/T93V3qMap1
WfsJjhMIX9LjBq3y9GAfAtAw7JYwkztr2AhYzCsK25wAj72zFY6FJTZ8LklfS41q
DrVtdjMx42IonDQtkzrqzfYlXdzzhZzuQHxn+qJODseoU8oDG9j3eKhp1dKgqLfx
tv1o3km1AoGBAOqXGEw2w94uVchCjuTum3XFYieEla0IUbHJCaWKU/hoSbht7j23
K7tA9//epBuRLGtYE0sPBK6i31mQT216muspO1g3pwGJhPy8VSpsJ1GQhB2G2UNz
kZlRK+2/gx35TdTi9x0C6UWk5XkhgWO3R35BlEnuV7EOyJunobiUcoObAoGBAMqJ
reuSbJajNGfeBzPel7F62ZDufC85hWaGKzeXIk3DkXcsEpeR5ogMGHCZ3dGf4Yz/
pcfjnCMIWjc+MkA4ppFd4432FJkxNQQP0z7njpXW3e5tionsMN+UwPzi4wZPKufK
osjw43JpBzpHGxG3ynLgZg7bSfrQZhPDTQw84nJbAoGBAKpy8mKeAB71R7rkMXNB
s48Uxca03RQGUWV+DxZKtcxt6fKpXUtWRd4ezJMLL+4fw0iTjCEjXmGNUf9/jVac
mOd44/erKBtD0m7YYIEcaE0pVfUmP8J0vDvL8MEkP56Nv/GIn8hijx/dOiaTI7JS
Pw4LlDVLikfJ2BTQ7f5xTes1AoGBAJ7h4HiDFgIZp1uvtfC/tjn5CEGEhBC7y+VA
bRify747I5rcDP2v66tf6bAzU+pExLhKN++Vov9sZvEdLmhoyGoSwBa2KzR9gHxe
ObYICjeLJfALKHnHuhM6ayY2iieB5UOOF6MQLSysLYpPC3IbvonddNJEvkUuRFVO
iNuHy5AvAoGAPqbzSNA05gf85zRO/JZAmZuWXG3o3pVougbkc211p+ynMpS+/bMb
c/nR36kOE551lFjIoAjoeIs16Wbq+00u9GlcQmyAdFpfaFCNHa3dayJwKJMW9Nia
fKbiiiOAQE8s2v8Paa+b00GspeWLow4u0G5lBVau4JjEVnl6ivLXlzY=
-----END RSA PRIVATE KEY-----
EOF
}

resource "alicloud_ssl_certificates_service_certificate" "defaultAeCL0Q" {
  certificate_name = "ld测试"
  cert             = <<EOF
-----BEGIN CERTIFICATE----- MIID4TCCAsmgAwIBAgIRANZGvLwT8kuWpPlZ/Aj+uPgwDQYJKoZIhvcNAQELBQAw XjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU ZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w HhcNMjQwODIzMDk0NzA0WhcNMjUwODIzMDk0NzA0WjAlMQswCQYDVQQGEwJDTjEW MBQGA1UEAxMNdGVzdGxkLnFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC AQoCggEBALmZY2geTFi+50gAVyDQH9Y5sTv8LLX6+MET1l3larzjX1M0Az9ZEIc0 TNrAp8mtJRlpQCzyDPZg88AwSdEwqSOSsnGzfS2DUcPJmdn2a2n5PLvWE28qPuSf 6fl3IhNiPzLYR51+7ccJKEQRhfOK2usmJo6oTG/0Lhh4BRH5owcclKv6n3YHaBVj JNigiq1/tlqU46toZvotPOORjpy21kJPZioHqOVCDO4zreMy2xuIiYtpSSmXwkEO zcQQ3K8sbRx9ED8SCdb229h7ioTug02YBXs0YOQZ024HFaIF8Nz1M+mdHy1jCbLd yJoT/jzE4RdldZKZJFaSKV1c7EYlzhkCAwEAAaOB0jCBzzAOBgNVHQ8BAf8EBAMC BaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMB8GA1UdIwQYMBaAFCiB JgXRNBo/wXMPu5PPFRw/A79/MGMGCCsGAQUFBwEBBFcwVTAhBggrBgEFBQcwAYYV aHR0cDovL29jc3AubXlzc2wuY29tMDAGCCsGAQUFBzAChiRodHRwOi8vY2EubXlz c2wuY29tL215c3NsdGVzdHJzYS5jcnQwGAYDVR0RBBEwD4INdGVzdGxkLnFxLmNv bTANBgkqhkiG9w0BAQsFAAOCAQEAnPJl1GrePDIulWfsETPbGnrZv3j3ZRXuou0o K32X/nydS/i/j+AUzKSyezmnR1edkgY1hbGaza702SLQJuGh2IqJvAFyifwV/CZ5 cpJIi5G7kWTBjZo9NgVnDMhR8y5DCKE8BhiUBwcSvKKC8se2yWHm1fk9pRxG0Mc6 0fstl40jtR5XZYsW1GhX4fzwrWuBodPKticgXPn2e24ec+4rVrziu5R7D77AzJjG Y/wzNYvAUWEzEya7Ve53nhu+WpIuIQn0ux8nPDioFdOjckn4jK3ePYdS2mWT6EBU BC74GYiBNDz0QgHADq1VTExeLzC0tw9PPdWl0WfoTgCCKLz0yA== -----END CERTIFICATE-----
EOF
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY----- MIIEpQIBAAKCAQEAuZljaB5MWL7nSABXINAf1jmxO/wstfr4wRPWXeVqvONfUzQD P1kQhzRM2sCnya0lGWlALPIM9mDzwDBJ0TCpI5KycbN9LYNRw8mZ2fZrafk8u9YT byo+5J/p+XciE2I/MthHnX7txwkoRBGF84ra6yYmjqhMb/QuGHgFEfmjBxyUq/qf dgdoFWMk2KCKrX+2WpTjq2hm+i0845GOnLbWQk9mKgeo5UIM7jOt4zLbG4iJi2lJ KZfCQQ7NxBDcryxtHH0QPxIJ1vbb2HuKhO6DTZgFezRg5BnTbgcVogXw3PUz6Z0f LWMJst3ImhP+PMThF2V1kpkkVpIpXVzsRiXOGQIDAQABAoIBAQCBPiw4A+k8X2vk +r+xjNyurCwcTmXAL81rfmnnputmL5tg8DZWtanJzQS7zC7LRPQxttZGtiOKqkbz DW1J6+3MZMo4XToNKIYWpduqKWvxNusxDkkoPy3evPEMlAY5o0/JE00DgrEHyfut MtqplocN+tocu1vHFi3HQkSdmM4LE46ZfFu5w1FRbNI1Gqjj/cwlF/T93V3qMap1 WfsJjhMIX9LjBq3y9GAfAtAw7JYwkztr2AhYzCsK25wAj72zFY6FJTZ8LklfS41q DrVtdjMx42IonDQtkzrqzfYlXdzzhZzuQHxn+qJODseoU8oDG9j3eKhp1dKgqLfx tv1o3km1AoGBAOqXGEw2w94uVchCjuTum3XFYieEla0IUbHJCaWKU/hoSbht7j23 K7tA9//epBuRLGtYE0sPBK6i31mQT216muspO1g3pwGJhPy8VSpsJ1GQhB2G2UNz kZlRK+2/gx35TdTi9x0C6UWk5XkhgWO3R35BlEnuV7EOyJunobiUcoObAoGBAMqJ reuSbJajNGfeBzPel7F62ZDufC85hWaGKzeXIk3DkXcsEpeR5ogMGHCZ3dGf4Yz/ pcfjnCMIWjc+MkA4ppFd4432FJkxNQQP0z7njpXW3e5tionsMN+UwPzi4wZPKufK osjw43JpBzpHGxG3ynLgZg7bSfrQZhPDTQw84nJbAoGBAKpy8mKeAB71R7rkMXNB s48Uxca03RQGUWV+DxZKtcxt6fKpXUtWRd4ezJMLL+4fw0iTjCEjXmGNUf9/jVac mOd44/erKBtD0m7YYIEcaE0pVfUmP8J0vDvL8MEkP56Nv/GIn8hijx/dOiaTI7JS Pw4LlDVLikfJ2BTQ7f5xTes1AoGBAJ7h4HiDFgIZp1uvtfC/tjn5CEGEhBC7y+VA bRify747I5rcDP2v66tf6bAzU+pExLhKN++Vov9sZvEdLmhoyGoSwBa2KzR9gHxe ObYICjeLJfALKHnHuhM6ayY2iieB5UOOF6MQLSysLYpPC3IbvonddNJEvkUuRFVO iNuHy5AvAoGAPqbzSNA05gf85zRO/JZAmZuWXG3o3pVougbkc211p+ynMpS+/bMb c/nR36kOE551lFjIoAjoeIs16Wbq+00u9GlcQmyAdFpfaFCNHa3dayJwKJMW9Nia fKbiiiOAQE8s2v8Paa+b00GspeWLow4u0G5lBVau4JjEVnl6ivLXlzY= -----END RSA PRIVATE KEY-----
EOF
}


`, name)
}

// Case 手动上传证书_TF接入_副本1744960888974_副本1748569871102 10867
func TestAccAliCloudDdosCooDomainResource_basic10867(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap10867)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence10867)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"ap-southeast-1"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type": "0",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
					},
					"real_servers": []string{
						"1.1.1.1"},
					"domain": "testcert.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultaRfzZ9.id}"},
					"https_ext":   "{\\\"Https2http\\\":1,\\\"Http2\\\":1,\\\"Http2https\\\":0}",
					"cert_name":   "523",
					"cert":        "-----BEGIN CERTIFICATE-----\\nMIID5DCCAsygAwIBAgIQWen3GebvT0GcE/a1MJVFgjANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MjMwOTE1MjNaFw0yNTA4MjMwOTE1MjNaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDF/H2/Oas8trhkHs7B4B8mN00eup/5Tqar2QO4Hm499GJECKc3eMiC v+aAlW74Iymb2Varnv+WMdFRVMQgpXesi3akvVp0QxecvcDliilkh4ddTK731Rd7 PaSK1JdQX1jdGGhVnhQz+cPNFBGZ3tMYGhUkgNfqa3UFucJcBuRub/Ircr+5Ob4D FxSglfTHi+/EFcp7vMAOztLD4zXmEz3NysDNP6NzN7SD72DwPp0nxyRjrBlHSOVg szB/bFasQdAhZGeo64MvSb+SivdWEMhHkwKA5MhhYOkDeNPPSmlxbw0Z3nOyeMmI YkaxzhpO5DZN382duTQmiQ+Yg60OfL3NAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAKUCvcbUCJRFbVowd1YorILivqRmS6ztR 9vLdj4YZBWxmmQrgkDlkl78r/rXlJbqHunSh2Wbag7y+GaQQwg8xcL4Z3KrKj4zg aHEP1DyFiaxMTEuC/L2RgSX3xlXcf6fQ46D3y3Ja3iHFQnjx6npNaSZ2bSEULvJg IjPiJ/nbid1TYR5vtg8vtwDQfY7/+q8/3DWKcQq+SGcd9dDS6u9vulNdlW8e14bS CKDEzS/axjoICl9JagASLcElWIit/eD5zGIKzTPC9mEXiX/J/gUr70y9GiE7Ue++ 4nFSGQOwjMh/wO2HlmRfToeZ3g6rRCijibBHKHBmVym7NCai2voZJQ==\\n-----END CERTIFICATE-----",
					"key":         "-----BEGIN RSA PRIVATE KEY-----\\nMIIEowIBAAKCAQEAxfx9vzmrPLa4ZB7OweAfJjdNHrqf+U6mq9kDuB5uPfRiRAin N3jIgr/mgJVu+CMpm9lWq57/ljHRUVTEIKV3rIt2pL1adEMXnL3A5YopZIeHXUyu 99UXez2kitSXUF9Y3RhoVZ4UM/nDzRQRmd7TGBoVJIDX6mt1BbnCXAbkbm/yK3K/ uTm+AxcUoJX0x4vvxBXKe7zADs7Sw+M15hM9zcrAzT+jcze0g+9g8D6dJ8ckY6wZ R0jlYLMwf2xWrEHQIWRnqOuDL0m/kor3VhDIR5MCgOTIYWDpA3jTz0ppcW8NGd5z snjJiGJGsc4aTuQ2Td/Nnbk0JokPmIOtDny9zQIDAQABAoIBAQDD6fU4y8UhwCG4 mS+5c6D/PQvoU35Hwkd1l7pxcFNgpTqz3egyISgxEdny9WwoyQq8eJWmICEEK+nY VEv7jiFdMWhG3kTq9RUhejeuLEiHfQE7Fs2w2kFxJ29yHapZ0u/pYOSljFarlATo I2rDW1aB7BVt2L1P7+ONteKZFAzpJckft5ceRUzs5Jm1Cqt8OWO3Km+FBbCROv8M TevW44aoMwBGXuqs06FV1Z4dafglskjt2O38V4acZpH8Nc8j+nCONKL3OxwKY6HQ WfnbXnTLCF3IuMiy8ntrY8HYU6EABiCdr+Pl5HmhI2nmtSFTFbD4Gq70vgPL0P1m iULJGJ7hAoGBANQPrOGe9qHcBydvcBHE7qA9v1+IaTj03qzDTopTi/jxcd9pEkei skLyHNQ5yJT0QjTxB9iYRLfZccOGFyqz/Sdz6CwwTWBZeXOQ2AX7FPEcCnNr1TpF yMrgOY3H93KJISEVS6kYskByjK7XzXCp0KQNS2EeIhAXcqXxNmSwylvpAoGBAO8C PdZHd6aLLEZyVO1aZVHDxqmbhmGVoY9wZ+uwR4K2Hu/fjk0qlR9cYpw8+N675Wr9 E9Ff5/wjK2+/+uocQV9Zoap2vgrwX7GASuO5KYdCOBn6oUOSa+Ru+LgBNyUkXYES mM8eFC1QqfcSrETLAQqd2lmLcuaMq6jJtbBpvzhFAoGAYd8mNC9wtr1dE+dLuvfA BnbZJ1dG8QGa7/NoAVGT7X5JxwmwZR2C1oD1q0FMAOtGzzZbH60PMicKaWousQfH E/lbs2FLpOdGtX6pJQF/5dPCQwkGrVFd3bxk87nRy6vcfW9drxp10mbL5To2WAQY Bk8Ydic5I2IfCNVt/ETX8FkCgYA+OkAtVQgi9WM+qC/SaFGu2yETManoKFQbC3IT HB9SOeaOH49mKesPcjc+ZGWLYDJYC7IoNicpL2L0wnAqmdavY5/CyQ2rvW+8wCE/ bwsP6z6+DNIFzM6IeBgLmE1qPzCVFWlxq2wnbDQEXvk5I/2ObRDXdYYh3ogm9vV2 C+I8XQKBgHIquvifRVvWf1q9WFZLQXZMv1flPhNaLmR+2k6gNpJ8SeiOmBtE7gT6 Je+YOXEKvfr6jaaJwYHPi6IhWHs4fQbgdK4jei30sRL7c8QdKEuwRdHPimmGNAPb UapzHY7xq0Wk9enAnM/SXkjTAJEkrpiQiDuPZVi4sIYCOqb+Ovu5\\n-----END RSA PRIVATE KEY-----",
					"cert_region": "ap-southeast-1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "0",
						"proxy_types.#":  "2",
						"real_servers.#": "1",
						"domain":         "testcert.qq.com",
						"instance_ids.#": "1",
						"https_ext":      CHECKSET,
						"cert_name":      CHECKSET,
						"cert":           "-----BEGIN CERTIFICATE-----\nMIID5DCCAsygAwIBAgIQWen3GebvT0GcE/a1MJVFgjANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MjMwOTE1MjNaFw0yNTA4MjMwOTE1MjNaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDF/H2/Oas8trhkHs7B4B8mN00eup/5Tqar2QO4Hm499GJECKc3eMiC v+aAlW74Iymb2Varnv+WMdFRVMQgpXesi3akvVp0QxecvcDliilkh4ddTK731Rd7 PaSK1JdQX1jdGGhVnhQz+cPNFBGZ3tMYGhUkgNfqa3UFucJcBuRub/Ircr+5Ob4D FxSglfTHi+/EFcp7vMAOztLD4zXmEz3NysDNP6NzN7SD72DwPp0nxyRjrBlHSOVg szB/bFasQdAhZGeo64MvSb+SivdWEMhHkwKA5MhhYOkDeNPPSmlxbw0Z3nOyeMmI YkaxzhpO5DZN382duTQmiQ+Yg60OfL3NAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAKUCvcbUCJRFbVowd1YorILivqRmS6ztR 9vLdj4YZBWxmmQrgkDlkl78r/rXlJbqHunSh2Wbag7y+GaQQwg8xcL4Z3KrKj4zg aHEP1DyFiaxMTEuC/L2RgSX3xlXcf6fQ46D3y3Ja3iHFQnjx6npNaSZ2bSEULvJg IjPiJ/nbid1TYR5vtg8vtwDQfY7/+q8/3DWKcQq+SGcd9dDS6u9vulNdlW8e14bS CKDEzS/axjoICl9JagASLcElWIit/eD5zGIKzTPC9mEXiX/J/gUr70y9GiE7Ue++ 4nFSGQOwjMh/wO2HlmRfToeZ3g6rRCijibBHKHBmVym7NCai2voZJQ==\n-----END CERTIFICATE-----",
						"key":            "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAxfx9vzmrPLa4ZB7OweAfJjdNHrqf+U6mq9kDuB5uPfRiRAin N3jIgr/mgJVu+CMpm9lWq57/ljHRUVTEIKV3rIt2pL1adEMXnL3A5YopZIeHXUyu 99UXez2kitSXUF9Y3RhoVZ4UM/nDzRQRmd7TGBoVJIDX6mt1BbnCXAbkbm/yK3K/ uTm+AxcUoJX0x4vvxBXKe7zADs7Sw+M15hM9zcrAzT+jcze0g+9g8D6dJ8ckY6wZ R0jlYLMwf2xWrEHQIWRnqOuDL0m/kor3VhDIR5MCgOTIYWDpA3jTz0ppcW8NGd5z snjJiGJGsc4aTuQ2Td/Nnbk0JokPmIOtDny9zQIDAQABAoIBAQDD6fU4y8UhwCG4 mS+5c6D/PQvoU35Hwkd1l7pxcFNgpTqz3egyISgxEdny9WwoyQq8eJWmICEEK+nY VEv7jiFdMWhG3kTq9RUhejeuLEiHfQE7Fs2w2kFxJ29yHapZ0u/pYOSljFarlATo I2rDW1aB7BVt2L1P7+ONteKZFAzpJckft5ceRUzs5Jm1Cqt8OWO3Km+FBbCROv8M TevW44aoMwBGXuqs06FV1Z4dafglskjt2O38V4acZpH8Nc8j+nCONKL3OxwKY6HQ WfnbXnTLCF3IuMiy8ntrY8HYU6EABiCdr+Pl5HmhI2nmtSFTFbD4Gq70vgPL0P1m iULJGJ7hAoGBANQPrOGe9qHcBydvcBHE7qA9v1+IaTj03qzDTopTi/jxcd9pEkei skLyHNQ5yJT0QjTxB9iYRLfZccOGFyqz/Sdz6CwwTWBZeXOQ2AX7FPEcCnNr1TpF yMrgOY3H93KJISEVS6kYskByjK7XzXCp0KQNS2EeIhAXcqXxNmSwylvpAoGBAO8C PdZHd6aLLEZyVO1aZVHDxqmbhmGVoY9wZ+uwR4K2Hu/fjk0qlR9cYpw8+N675Wr9 E9Ff5/wjK2+/+uocQV9Zoap2vgrwX7GASuO5KYdCOBn6oUOSa+Ru+LgBNyUkXYES mM8eFC1QqfcSrETLAQqd2lmLcuaMq6jJtbBpvzhFAoGAYd8mNC9wtr1dE+dLuvfA BnbZJ1dG8QGa7/NoAVGT7X5JxwmwZR2C1oD1q0FMAOtGzzZbH60PMicKaWousQfH E/lbs2FLpOdGtX6pJQF/5dPCQwkGrVFd3bxk87nRy6vcfW9drxp10mbL5To2WAQY Bk8Ydic5I2IfCNVt/ETX8FkCgYA+OkAtVQgi9WM+qC/SaFGu2yETManoKFQbC3IT HB9SOeaOH49mKesPcjc+ZGWLYDJYC7IoNicpL2L0wnAqmdavY5/CyQ2rvW+8wCE/ bwsP6z6+DNIFzM6IeBgLmE1qPzCVFWlxq2wnbDQEXvk5I/2ObRDXdYYh3ogm9vV2 C+I8XQKBgHIquvifRVvWf1q9WFZLQXZMv1flPhNaLmR+2k6gNpJ8SeiOmBtE7gT6 Je+YOXEKvfr6jaaJwYHPi6IhWHs4fQbgdK4jei30sRL7c8QdKEuwRdHPimmGNAPb UapzHY7xq0Wk9enAnM/SXkjTAJEkrpiQiDuPZVi4sIYCOqb+Ovu5\n-----END RSA PRIVATE KEY-----",
						"cert_region":    "ap-southeast-1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"real_servers": []string{
						"1.1.1.1", "2.2.2.2", "3.3.3.3"},
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultaRfzZ9.id}", "${alicloud_ddoscoo_instance.defaultQPviUE.id}", "${alicloud_ddoscoo_instance.defaultgyrXgo.id}"},
					"https_ext":    "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
					"cert_name":    "729",
					"cert":         "-----BEGIN CERTIFICATE-----\\nMIID5DCCAsygAwIBAgIQFO5STIOlR/KkRB3gDHsi5zANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MTMwNTEyMDdaFw0yNTA4MTMwNTEyMDdaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDwg/HpVcY9yaLrxApP4kvrGwwPbpqOXw9NDWqSD2ms7Qvhi84dTnZT 9xyLXjW0MYI+W13kes4Prl37L8VHOZFvxdzZdzuHvZMkJDaRcekIiUYp+5SFqzzT EtlQodDZBLy0uSlwHnDHymUNJg2nXma+cOOwBBhvb9h+j5v4uuUkQoHRzIUHCkd6 8LR4zbOe+zrhxUU5AYm8C77ZJphDXi9GYm/moajpk+0biCDZ/vZSjTngZEujQKam dgRfVCiWgoUrueiijh1cLGf+W15A3kNo5UXfQrhbHVRQY0vy9dqU1ZBms0pWdc9G dJGd2kXkGNWtEbk4NCN6c64AW2L3G2JNAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAFl4tmH+m/lGr7aKxxz5uYdqepcjCHTzW KpxxEHr+sL9fOkap6I/NA9MYKsuVfnrTMnEJXeAj24g8q6t38bGTEMcUvh3YGS+P 7DLIF2T2/f9oiO6LVD72zC/HkqaenzWla536G4FLAW4+MzAdn+R7hd5QN4CuFaKI nWWr5KMjZfgyxPkbe+WiD8QryR/EOd3za5vdycZRnBVIcbZQzEHwxizm3JyPNYy8 KeIjt5kGjFt9ccb1oPXjVibB2SUO322kbhMeCXKq2tswnZBICy+1EcbmaxNlkxwc hcdb2wz/3pzH0jbexgTFrNWeoJyYykM8s/3cMY/t0dKppve4FXLG9A==\\n-----END CERTIFICATE-----",
					"key":          "-----BEGIN RSA PRIVATE KEY-----\\nMIIEowIBAAKCAQEA8IPx6VXGPcmi68QKT+JL6xsMD26ajl8PTQ1qkg9prO0L4YvO HU52U/cci141tDGCPltd5HrOD65d+y/FRzmRb8Xc2Xc7h72TJCQ2kXHpCIlGKfuU has80xLZUKHQ2QS8tLkpcB5wx8plDSYNp15mvnDjsAQYb2/Yfo+b+LrlJEKB0cyF BwpHevC0eM2znvs64cVFOQGJvAu+2SaYQ14vRmJv5qGo6ZPtG4gg2f72Uo054GRL o0CmpnYEX1QoloKFK7nooo4dXCxn/lteQN5DaOVF30K4Wx1UUGNL8vXalNWQZrNK VnXPRnSRndpF5BjVrRG5ODQjenOuAFti9xtiTQIDAQABAoIBAQCkuln3bA3ox69U NuKxL9a7Ybzy3NfyZtz98xBolTHVhE083xn+LH0SqQ7dzVqO3dHMj5tRH2L+jnhD z8YYMC+SFDxcnTMilw6uFDdjilcGx65Mlsh0fIGeNyyr8wgtevcb+C2PYunvjImF Zei4FwnbqUnohgWOXVYz6Hv08Vx7ZdW+QiH62I/LS73G7d2EPb26Zo3zMKg/H5AD xuNk82MDW0lCgrw869Yqhcd3GkkmgWi+S71AE0ftY03QeBrsSZbzz4Zsgk2GsEBt fGclOu2c5sNRhLy4o7GiZghPS32zkiec9H76Ip5n/nwXgcYCoxfvOQL+b8U5vpap gbgjWGLBAoGBAPQsJGrYAzt1naGbzwEgYbbyItKD3Bf30YVcZxnZW1fZn2H9NpWl oIBAiO8ls1WP32Mf1us91Bp0V4ESmmklFb5ZllRZXnGt6U2pvxl2FZbK18Nvcvw/ kc66t643mKdfsTAF3sRSmeeWxNSB5C+/yZfHARbBoZ64hYqSut2sivjlAoGBAPwq dDBrz4P/PGn191ZfpWyTC4NRMpC2zYkHoWtMJNs96bsZub4phWyo1KLL4sbAZ+uV GfpRpE2u8mKQWEUU2Gz/KI0cY1e40Icg/ZBlzNglb7ssBXKMrI3R2pigIzEAArhU KwsjhreGF7ix5DEyNT4i7PQL/tOu3uh9SUXl9zVJAoGANta/KxvuxejpiUVUHZ2n NI53Ua55vQxUi04wfba6dCWVTU2wd7WmMYfM+WEPQPU6J6ob++N8AqEEkiGaemjw 1DqMr88OjhuQHXg1SkOiH6bZBLTAL3Ubi0GWRVOJPnYYdn+rA47FsCTFejDeDfdW EHeKgBDm+p3YqEHCJE0/PR0CgYA+mw+zweCIhgLqz81znVWFylAubyddtHT9E27p I8N2xz1TXYS3CLn+i0AXlwUbkUN7ws3rTv+65bd57xprNEyzavoXZrfnXJQxKGir xAqCk3DVCI3lrbVdlH9wKznxfW4vc34oSs60m88h5NChwjRj0+n+gUfoKF9hW1Go z/p7OQKBgCqfExvjjmBibsr0ZtTprMGV2qI55LoVnUXcM1xSBbgU0rLwvw2YBrWg MRl3ixYN851wGD1LhgpwFjr7SnwEhpKdDloKSE1ANM5LE7zvJHsPY2uvmY/Rbn7i RqIzeCbUVYt1Ow0WEqy0DUy/fGLQEz9viwLnvTcDNWeuSsljeAO7\\n-----END RSA PRIVATE KEY-----",
					"cert_region":  "cn-hangzhou",
					"ocsp_enabled": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"real_servers.#": "3",
						"instance_ids.#": "3",
						"https_ext":      CHECKSET,
						"cert_name":      CHECKSET,
						"cert":           "-----BEGIN CERTIFICATE-----\nMIID5DCCAsygAwIBAgIQFO5STIOlR/KkRB3gDHsi5zANBgkqhkiG9w0BAQsFADBe MQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl c3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe Fw0yNDA4MTMwNTEyMDdaFw0yNTA4MTMwNTEyMDdaMCcxCzAJBgNVBAYTAkNOMRgw FgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw ggEKAoIBAQDwg/HpVcY9yaLrxApP4kvrGwwPbpqOXw9NDWqSD2ms7Qvhi84dTnZT 9xyLXjW0MYI+W13kes4Prl37L8VHOZFvxdzZdzuHvZMkJDaRcekIiUYp+5SFqzzT EtlQodDZBLy0uSlwHnDHymUNJg2nXma+cOOwBBhvb9h+j5v4uuUkQoHRzIUHCkd6 8LR4zbOe+zrhxUU5AYm8C77ZJphDXi9GYm/moajpk+0biCDZ/vZSjTngZEujQKam dgRfVCiWgoUrueiijh1cLGf+W15A3kNo5UXfQrhbHVRQY0vy9dqU1ZBms0pWdc9G dJGd2kXkGNWtEbk4NCN6c64AW2L3G2JNAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo gSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG FWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15 c3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx LmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAFl4tmH+m/lGr7aKxxz5uYdqepcjCHTzW KpxxEHr+sL9fOkap6I/NA9MYKsuVfnrTMnEJXeAj24g8q6t38bGTEMcUvh3YGS+P 7DLIF2T2/f9oiO6LVD72zC/HkqaenzWla536G4FLAW4+MzAdn+R7hd5QN4CuFaKI nWWr5KMjZfgyxPkbe+WiD8QryR/EOd3za5vdycZRnBVIcbZQzEHwxizm3JyPNYy8 KeIjt5kGjFt9ccb1oPXjVibB2SUO322kbhMeCXKq2tswnZBICy+1EcbmaxNlkxwc hcdb2wz/3pzH0jbexgTFrNWeoJyYykM8s/3cMY/t0dKppve4FXLG9A==\n-----END CERTIFICATE-----",
						"key":            "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA8IPx6VXGPcmi68QKT+JL6xsMD26ajl8PTQ1qkg9prO0L4YvO HU52U/cci141tDGCPltd5HrOD65d+y/FRzmRb8Xc2Xc7h72TJCQ2kXHpCIlGKfuU has80xLZUKHQ2QS8tLkpcB5wx8plDSYNp15mvnDjsAQYb2/Yfo+b+LrlJEKB0cyF BwpHevC0eM2znvs64cVFOQGJvAu+2SaYQ14vRmJv5qGo6ZPtG4gg2f72Uo054GRL o0CmpnYEX1QoloKFK7nooo4dXCxn/lteQN5DaOVF30K4Wx1UUGNL8vXalNWQZrNK VnXPRnSRndpF5BjVrRG5ODQjenOuAFti9xtiTQIDAQABAoIBAQCkuln3bA3ox69U NuKxL9a7Ybzy3NfyZtz98xBolTHVhE083xn+LH0SqQ7dzVqO3dHMj5tRH2L+jnhD z8YYMC+SFDxcnTMilw6uFDdjilcGx65Mlsh0fIGeNyyr8wgtevcb+C2PYunvjImF Zei4FwnbqUnohgWOXVYz6Hv08Vx7ZdW+QiH62I/LS73G7d2EPb26Zo3zMKg/H5AD xuNk82MDW0lCgrw869Yqhcd3GkkmgWi+S71AE0ftY03QeBrsSZbzz4Zsgk2GsEBt fGclOu2c5sNRhLy4o7GiZghPS32zkiec9H76Ip5n/nwXgcYCoxfvOQL+b8U5vpap gbgjWGLBAoGBAPQsJGrYAzt1naGbzwEgYbbyItKD3Bf30YVcZxnZW1fZn2H9NpWl oIBAiO8ls1WP32Mf1us91Bp0V4ESmmklFb5ZllRZXnGt6U2pvxl2FZbK18Nvcvw/ kc66t643mKdfsTAF3sRSmeeWxNSB5C+/yZfHARbBoZ64hYqSut2sivjlAoGBAPwq dDBrz4P/PGn191ZfpWyTC4NRMpC2zYkHoWtMJNs96bsZub4phWyo1KLL4sbAZ+uV GfpRpE2u8mKQWEUU2Gz/KI0cY1e40Icg/ZBlzNglb7ssBXKMrI3R2pigIzEAArhU KwsjhreGF7ix5DEyNT4i7PQL/tOu3uh9SUXl9zVJAoGANta/KxvuxejpiUVUHZ2n NI53Ua55vQxUi04wfba6dCWVTU2wd7WmMYfM+WEPQPU6J6ob++N8AqEEkiGaemjw 1DqMr88OjhuQHXg1SkOiH6bZBLTAL3Ubi0GWRVOJPnYYdn+rA47FsCTFejDeDfdW EHeKgBDm+p3YqEHCJE0/PR0CgYA+mw+zweCIhgLqz81znVWFylAubyddtHT9E27p I8N2xz1TXYS3CLn+i0AXlwUbkUN7ws3rTv+65bd57xprNEyzavoXZrfnXJQxKGir xAqCk3DVCI3lrbVdlH9wKznxfW4vc34oSs60m88h5NChwjRj0+n+gUfoKF9hW1Go z/p7OQKBgCqfExvjjmBibsr0ZtTprMGV2qI55LoVnUXcM1xSBbgU0rLwvw2YBrWg MRl3ixYN851wGD1LhgpwFjr7SnwEhpKdDloKSE1ANM5LE7zvJHsPY2uvmY/Rbn7i RqIzeCbUVYt1Ow0WEqy0DUy/fGLQEz9viwLnvTcDNWeuSsljeAO7\n-----END RSA PRIVATE KEY-----",
						"cert_region":    "cn-hangzhou",
						"ocsp_enabled":   "true",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap10867 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence10867(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_ddoscoo_instance" "defaultaRfzZ9" {
  normal_bandwidth = "100"
  normal_qps       = "500"
  bandwidth_mode   = "2"
  product_plan     = "3"
  product_type     = "ddosDip"
  period           = "1"
  port_count       = "5"
  name             = "测试手动上传证书"
  function_version = "0"
  domain_count     = "10"
}

resource "alicloud_ddoscoo_instance" "defaultgyrXgo" {
  normal_bandwidth = "100"
  normal_qps       = "500"
  bandwidth_mode   = "2"
  product_plan     = "3"
  product_type     = "ddosDip"
  period           = "1"
  port_count       = "5"
  name             = "测试手动上传证书"
  function_version = "0"
  domain_count     = "10"
}

resource "alicloud_ddoscoo_instance" "defaultQPviUE" {
  normal_bandwidth = "100"
  normal_qps       = "500"
  bandwidth_mode   = "2"
  product_plan     = "3"
  product_type     = "ddosDip"
  period           = "1"
  port_count       = "5"
  name             = "测试手动上传证书"
  function_version = "0"
  domain_count     = "10"
}


`, name)
}

// Case 国内高防测试自定义Headers 11574
func TestAccAliCloudDdosCooDomainResource_basic11574(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap11574)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence11574)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type":      "0",
					"ocsp_enabled": "false",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "websocket",
						},
					},
					"real_servers": []string{
						"1.1.1.1", "2.2.2.2", "3.3.3.3"},
					"domain": "testld.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.default9nJ7Ie.id}"},
					"https_ext":       "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
					"cert_identifier": "${alicloud_ssl_certificates_service_certificate.default3MYZEt.id}-cn-hangzhou",
					"custom_headers":  "{\\\"3444\\\":\\\"5555\\\",\\\"666\\\":\\\"$ReqClientPort\\\",\\\"77777\\\":\\\"$ReqClientIP\\\"}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":         "0",
						"ocsp_enabled":    "false",
						"proxy_types.#":   "3",
						"real_servers.#":  "3",
						"domain":          "testld.qq.com",
						"instance_ids.#":  "1",
						"https_ext":       CHECKSET,
						"cert_identifier": CHECKSET,
						"custom_headers":  CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80", "8080"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "websockets",
						},
					},
					"custom_headers": "{\\\"22\\\":\\\"$ReqClientIP\\\",\\\"77\\\":\\\"88\\\",\\\"99\\\":\\\"$ReqClientPort\\\"}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"proxy_types.#":  "3",
						"custom_headers": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type": "1",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
					},
					"real_servers": []string{
						"1.qq.com"},
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.default9nJ7Ie.id}", "${alicloud_ddoscoo_instance.defaultYi1ySx.id}", "${alicloud_ddoscoo_instance.default9KfXQ3.id}"},
					"cert_identifier": "${alicloud_ssl_certificates_service_certificate.defaultAeCL0Q.id}-cn-hangzhou",
					"custom_headers":  "{\\\"44\\\":\\\"55\\\",\\\"66\\\":\\\"$ReqClientPort\\\",\\\"77\\\":\\\"$ReqClientIP\\\"}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":         "1",
						"proxy_types.#":   "2",
						"real_servers.#":  "1",
						"instance_ids.#":  "3",
						"cert_identifier": CHECKSET,
						"custom_headers":  CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap11574 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence11574(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "cert_region" {
  default = "default3MYZEt.CertId-${cn-hangzhou}"
}

resource "alicloud_ddoscoo_instance" "default9nJ7Ie" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "0"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ssl_certificates_service_certificate" "default3MYZEt" {
  cert             = <<EOF
-----BEGIN CERTIFICATE-----
MIID4TCCAsmgAwIBAgIRALw5sXZD1UHDhmh/t2VTQw4wDQYJKoZIhvcNAQELBQAw
XjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU
ZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w
HhcNMjUwOTI2MDI1NDU1WhcNMjYwOTI2MDI1NDU1WjAlMQswCQYDVQQGEwJDTjEW
MBQGA1UEAxMNdGVzdGxkLnFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBAMrfIvzgwhQegAeYFBRIR2LIwWT3cnKA7dLTmQUmusSqmx/AgA1ctaw8
/BUaRCCjamkYKnbDqBSYPUGMicLUVTbgiXuupoFwGBbkHN9AyetUiV86A8hebDi0
Hp3mK6AwIX432mb8nKiM3GCjVflJRt//xOybCpkqLyXFmOQxXUunZJEUUic+JHWa
bVlBxFzd4CDnBRrw0q0JPti0322TuL9HjiGkiJp2BvnMH++qtlTjwzOxMvTYeiz8
+E+yl0kzCW+bmMZK+t39nWrX57MvggNP8KsT6YCHcGgbSPQPcfx0kBL2IAU7eWbX
Xgpat3v/zRXxcIPjvg1EBmcw2rxu8dMCAwEAAaOB0jCBzzAOBgNVHQ8BAf8EBAMC
BaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMB8GA1UdIwQYMBaAFCiB
JgXRNBo/wXMPu5PPFRw/A79/MGMGCCsGAQUFBwEBBFcwVTAhBggrBgEFBQcwAYYV
aHR0cDovL29jc3AubXlzc2wuY29tMDAGCCsGAQUFBzAChiRodHRwOi8vY2EubXlz
c2wuY29tL215c3NsdGVzdHJzYS5jcnQwGAYDVR0RBBEwD4INdGVzdGxkLnFxLmNv
bTANBgkqhkiG9w0BAQsFAAOCAQEAKtDeQoQtloF6mvMOC0AYwJ2as7XyxfKKoqPs
dW7VHuASnB5AUeSmqPz3H8+qS7IX9VZDmTr2JxPRsJ+eYXMMI3UUlHUik0BcMt3Y
JfsV6nRgKm8JwktSHCsyVPDYU3zCO6KF1tUVKa18l61Twq81+gwX6jlmRy45/kPe
6yPUYA5FrNWc5ZWs4LcEM9F5L9xkhJVS8uICU09k8pwYsmU87z5mHaRxxSYjCoF2
gUrJjy6iWYfSJRWbDDA4p+BVZMuK3bGV4K7bS2lKjUPz7EZSUKQbWrzCMEOr7E8Y
9sFvHi49Blv8zllUS3clDdsP7nYPtU0hNysA9m9+eKkggFCo9Q==
-----END CERTIFICATE-----

EOF
  certificate_name = "146"
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAyt8i/ODCFB6AB5gUFEhHYsjBZPdycoDt0tOZBSa6xKqbH8CA
DVy1rDz8FRpEIKNqaRgqdsOoFJg9QYyJwtRVNuCJe66mgXAYFuQc30DJ61SJXzoD
yF5sOLQeneYroDAhfjfaZvycqIzcYKNV+UlG3//E7JsKmSovJcWY5DFdS6dkkRRS
Jz4kdZptWUHEXN3gIOcFGvDSrQk+2LTfbZO4v0eOIaSImnYG+cwf76q2VOPDM7Ey
9Nh6LPz4T7KXSTMJb5uYxkr63f2datfnsy+CA0/wqxPpgIdwaBtI9A9x/HSQEvYg
BTt5ZtdeClq3e//NFfFwg+O+DUQGZzDavG7x0wIDAQABAoIBAF3lJb/t1NXUAgTB
xfVXOLcHXL866d0GQEyWQ9oHAWV54v8wrPPCR5F2zmOD+ykyCVOn7Ct4xif2CE/4
2F/0v5X9GZTFkmoRNA0QOi64QVxqjYQmCU8pKKSb9Rm2yNVEwZO2DR8iZu15+Ju/
rVCKMkQFkKLD5YVbaWPtjyR6lognFrzkehASnmk2xGbqPjpr6wXWVVQ/MkJd4Nwt
SFzF9veBZRjSmxFCl9yowE1WdsEmvkzQyX0bI8u/pY3z4hj9EuffQz9/zzL5xVS3
vNTN0CuRTyOnTHaB+6K/SIh2nqkJRwAb9UPTokAgUnO/EhN4W1CaGqimFCxxgSVT
yt3c/LECgYEAyvWUVtPAbDbrQhJ7L7dM5eR98xbhy7ZF4n0djJ6w6qPu/FSLTOjq
j2REtBzDn3xDo6Z/5U0vJG9P7RpleZ3s8g0vF4zN8Lu6EkDZfsRriq2sW9Nl+f0H
3zNocOVafQXyrrM3WytBbsdHwRi0oWs0z22h1pL2SZSn4TWnzgjrdUMCgYEA/+Ox
LXvAaCgoLYUZHuBhyoPmocl7a14MUzii8dnxjCIaawS7YjSXoc5mSoYdft6wgm+U
sjqGgkqTFasiSJUN+d7367Liy+aPnggCqK6rE9Hgi2T7sdQv+XtU+FMtd7cPzbkQ
WVsMHVpr7dA+bzO+IE6wgTxX8g8soSrCO1aUgDECgYAlLlQci/JGYNE8a3JRzXyy
6OcB74Ex9pRa55zQNAopEhsn8r0KO+ksl6vWayaTQwqJImlvsnIedJ3py9onK31K
4otr/wmDPoDZ8zNk+8rPvv1CXTnjUC1vAFXzyLCJEtvgkUhk1UnJZ4yHnWUJ5T/p
eCYbzxR7alZO9atmHVA1TQKBgD8Lx3SQX/iJpFSKzYSo/g8abnGAJdNvSZQbiTIT
Y4sGQAIMGWr50D5CjztfTdcbYNvSSA2dk9R4MUMOdhTx/I6K3ASLf4uDU/E4wgbz
eh0ZAbz2dXj78ZIDTA0e2T38sX0bUqbhYtu8koj2XNujIP3uxVgiGPz/thxDX7Wl
AFORAoGAEh5MIb1j3Z8n2isB9AxP+Ju6Q38AueX0iKvvjFUQiqzQKgaa16VajPw1
YDMn3aoIIA9UyghkSmKdBWXAgpRWqRLqakbN58PMdtmDRhp2qqz7xljpOBSKRs3t
G5w8hpXVQAflI/SUAKdnoQdHoezMX8WWQzQAlOlh4lfTKAPOh8E=
-----END RSA PRIVATE KEY-----
EOF
}

resource "alicloud_ssl_certificates_service_certificate" "defaultAeCL0Q" {
  certificate_name = "ld测试headers567"
  cert             = <<EOF
-----BEGIN CERTIFICATE----- MIID4TCCAsmgAwIBAgIRAJ3bnm3BlEw5iDBJ9msY+YgwDQYJKoZIhvcNAQELBQAw XjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU ZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w HhcNMjUwOTI2MDI1NjMzWhcNMjYwOTI2MDI1NjMzWjAlMQswCQYDVQQGEwJDTjEW MBQGA1UEAxMNdGVzdGxkLnFxLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC AQoCggEBAK0Owiyqlde8Y3QNQdtQYLZkqm7zf2HAlJksd6EZeG/KBDTDnHOEoGs9 v5xdj2Wld0OkjC1M+Y557cjW6nKp+espiGl+LJ8s4OAKssupUaXktKG4Hs2EoA+e UIvFI7OfMzgv5prG9/reAEPPMkebbpVuuS2WVqtzRc3G7CShkcv75viVandQN3Xf GezpU9G5Dhfa3Ib6etqO6LHC9wCT06LK0NKscp7NFSbeGY/AHu2qHm5EuaityWpk xRDznaMveEEvs0fav0wW/zv4WvyDP+IQ6hjFgAUSw17iSW8VJEaGJ8KG9arddjG+ oZhLkYViTseER3HS0H69qrm7Qcvks9sCAwEAAaOB0jCBzzAOBgNVHQ8BAf8EBAMC BaAwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMB8GA1UdIwQYMBaAFCiB JgXRNBo/wXMPu5PPFRw/A79/MGMGCCsGAQUFBwEBBFcwVTAhBggrBgEFBQcwAYYV aHR0cDovL29jc3AubXlzc2wuY29tMDAGCCsGAQUFBzAChiRodHRwOi8vY2EubXlz c2wuY29tL215c3NsdGVzdHJzYS5jcnQwGAYDVR0RBBEwD4INdGVzdGxkLnFxLmNv bTANBgkqhkiG9w0BAQsFAAOCAQEAItakLjDpt86k9nl7ZKa0nKMMov3IGfpv94yW /AUUM6cB73ygVItJGaOG9QLO2r44jdum0pyJNusJAsHPgURUb/0RlEsGYOgTuhPG +gLlr2Yo7ekI5PdQAcVNkG2vaLyVGVo7WY6OuF3kHqt09HL8HvEpj0gwxumNsCEu sEq6nFwqgqhLa0yfVT6wAdp4he3vuo+Kgla9IOUTMInVw5WerxL2htCYQRhHme7q 98QFJVxkrnSkGYaNYnIm1opwwg9U52I5vGpd5KZkR4PiM8gxeBCLbIr9+RuJdoIa jFBpKbR/44TV3e6uwRWxvip83pE3DWxQEn8Z4rlrrnTtUxPsCw== -----END CERTIFICATE-----
EOF
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY----- MIIEogIBAAKCAQEArQ7CLKqV17xjdA1B21BgtmSqbvN/YcCUmSx3oRl4b8oENMOc c4Sgaz2/nF2PZaV3Q6SMLUz5jnntyNbqcqn56ymIaX4snyzg4Aqyy6lRpeS0obge zYSgD55Qi8Ujs58zOC/mmsb3+t4AQ88yR5tulW65LZZWq3NFzcbsJKGRy/vm+JVq d1A3dd8Z7OlT0bkOF9rchvp62o7oscL3AJPTosrQ0qxyns0VJt4Zj8Ae7aoebkS5 qK3JamTFEPOdoy94QS+zR9q/TBb/O/ha/IM/4hDqGMWABRLDXuJJbxUkRoYnwob1 qt12Mb6hmEuRhWJOx4RHcdLQfr2qubtBy+Sz2wIDAQABAoIBAAXcX5gaqNt3Dlky T74pMTVMIHeEeJZrarzrRBvpHGqQyWauD0DcR4CKRVB63K3hFjJswrCQEE2SdIqe OK9scUHVFMEZ3FIBt5Xu1tJN6C15muJ1NVnZeYA96NVq9kQRiq8G9ETeoyxUU2b4 f+fr7ClUaCISmtnQnBcVew7ch+8ECSeyrnSFXE8MBDghEYGVwJdTYye7O5pbBlUI VkAjme/vZ5tgBEizTbOjJ1wToprXS90wbUN2SNrZ3O80ctnfQOYeN4FHZoKK4iZ5 YWeIY/3s/EhgHL4BpVE9xBr4l4Ji1OhVj55vnEca7V80gBLI8KsKmswgPUDErQZQ 59HeFU0CgYEA4+XTWyr/OHaxIZvrgTbroB8IQcRS8N1yp9Mhi0OxNNGNN/0mzS3J R1H0nmBqH3uiJtUHekHkaILZeB6AJsmBdF4wvqFwhinC879He4UxVJhwrNmrhyN7 e4U8JpcScfqkRSCw7IyWXA8utgJwkrwSBRWpPJH0uDzEBdG6SQUbRv0CgYEAwmXD 1HNhc6zYljOpYJJ303aa1tbo5iLesnlmKRLGH02niHsjTKlKk8OgpL0mDXZPZcBO cGBv89VpEadgmXkvSBF3EcYzznMx4bE5roODG/uZpAgYhFCtck8OiaSpMaLAsNfg iX5UoUpGC/sjzRlb4pTuYXgdJIpq7uYNESPDWbcCgYB4gkw8NkVVNzY9PnTPFBqO xjoYhNcS9Rau9e8T9EydUH3rcFr3PlHj+8ttgDA9y6OYAyf8FyHPvfl/EinT2NQy k6gU9ctJtoWXXLtQ7sKCKEXUsRmJ0VJ3o6GGCna/fLkmsL29qC3OKl0Z87pRbU/e DnE6TTfpwIArT5bBDAk7hQKBgEOyuuL2l354OCj+rsKx30bmLAAbcZoLlLBktJ4j KpnmKizObTmaTx1saDpWoLTZecb+lJJwEyzAKkn6lxp7aGrZojSvaIVB+qIVXPgY VoVdNmE4bIlPq0d+nppynHVrr10moMxhEUnmGsS7XflkFmqu4OFSuhtqlNjDhtNs N/C7AoGAVYQOLWtIyST4GnAIqnGzNHe2iBm28lUkPqaqR69x7MIxq7gr9XsLpaRo ocoes5G2ZLILhifitPWm/LDS+TX762zGb0lotm2y6pDSR1l0A5Zjbw67H6mJHgDN CaGUag785i8Fln+pJ3DkuQ0tH70wAcZooqdWk7aLs81JTLaImTY= -----END RSA PRIVATE KEY-----
EOF
}

resource "alicloud_ddoscoo_instance" "defaultYi1ySx" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "0"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}

resource "alicloud_ddoscoo_instance" "default9KfXQ3" {
  normal_qps        = "3000"
  bandwidth_mode    = "2"
  product_type      = "ddoscoo"
  period            = "1"
  port_count        = "50"
  name              = "test"
  service_bandwidth = "200"
  base_bandwidth    = "30"
  bandwidth         = "50"
  function_version  = "0"
  address_type      = "Ipv4"
  edition_sale      = "coop"
  domain_count      = "50"
  product_plan      = "9"
}


`, name)
}

// Case 国际高防自定义header 11581
func TestAccAliCloudDdosCooDomainResource_basic11581(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ddoscoo_domain_resource.default"
	ra := resourceAttrInit(resourceId, AlicloudDdosCooDomainResourceMap11581)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DdosCooServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDdosCooDomainResource")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccddoscoo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDdosCooDomainResourceBasicDependence11581)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"ap-southeast-1"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"rs_type": "0",
					"proxy_types": []map[string]interface{}{
						{
							"proxy_ports": []string{
								"80"},
							"proxy_type": "http",
						},
						{
							"proxy_ports": []string{
								"443"},
							"proxy_type": "https",
						},
					},
					"real_servers": []string{
						"1.1.1.1"},
					"domain": "testcert.qq.com",
					"instance_ids": []string{
						"${alicloud_ddoscoo_instance.defaultdPWcCl.id}"},
					"https_ext":      "{\\\"Https2http\\\":1,\\\"Http2\\\":1,\\\"Http2https\\\":0}",
					"cert_name":      "202",
					"cert":           "-----BEGIN CERTIFICATE-----\\nMIID5TCCAs2gAwIBAgIRAJ8XVm8/e0sNiF7SQKrNVpUwDQYJKoZIhvcNAQELBQAw\\nXjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU\\nZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w\\nHhcNMjUwOTI5MDY1NDI3WhcNMjYwOTI5MDY1NDI3WjAnMQswCQYDVQQGEwJDTjEY\\nMBYGA1UEAxMPdGVzdGNlcnQucXEuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A\\nMIIBCgKCAQEAqjAfE7/ToSl6okR9YhQSz/vw+NvQUU4i/p+7HKqCoLC+l+rY4HzU\\n7m7fvgfkPZtTItgaiIKURo/Y+UmIaN+JNvnaTPZTXRgkfdccdTvYOeYzHn77AzFd\\nsjCcystdpXdJ/wrtfABa5XGNrFMeXYpf9Gxp+1A56modo8wiWAS48rysfzS8vHUw\\ncYuog5Nxttg5Kaqs2w1eUxiPLsXp+mN5CIhP46m1U5cisq9J1utAbKaHeRO+pWFn\\nAu3e4AK1utaxhGdYCDMX3xwlNWFIgpNepFpdrs+H8btwsoPUkbwdYM+lFkGRI8XZ\\n4dMwQEI6p6iohthk6o9GQj+otOEqkGHyHQIDAQABo4HUMIHRMA4GA1UdDwEB/wQE\\nAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwHwYDVR0jBBgwFoAU\\nKIEmBdE0Gj/Bcw+7k88VHD8Dv38wYwYIKwYBBQUHAQEEVzBVMCEGCCsGAQUFBzAB\\nhhVodHRwOi8vb2NzcC5teXNzbC5jb20wMAYIKwYBBQUHMAKGJGh0dHA6Ly9jYS5t\\neXNzbC5jb20vbXlzc2x0ZXN0cnNhLmNydDAaBgNVHREEEzARgg90ZXN0Y2VydC5x\\ncS5jb20wDQYJKoZIhvcNAQELBQADggEBAKN/GA4FZGLzaEGs+8zbNkCbNaq/zs+V\\nDBmdWF124wn+xke7PMN6PByWHBWFg7vJIfN7J7O4cFuShz23qcKj3qnDbUNKuK0X\\nTjGOJcFqSaHydlrPnmVSKEIM2J5O+VyR2gjuP+J6GvJavr0hFLRXB5gVmDfpxX6P\\nQX3SOilfn1rM0L7x1UG9t3t1USTumAgOuV/rflhOZs88nDjllNiR7RzsczBvrjes\\n5dP+mBcCVy9PUcVHJqNNNvbR4Vly8Ki+3c10jkxR9fVMnNpi2uI/eya1SLKGM6yr\\nHxEBkAxaANt8vEcCyiEBraqUvf9gBydvRKG0V6gVeeZ9+5WJd84WBzY=\\n-----END CERTIFICATE-----\\n",
					"key":            "-----BEGIN RSA PRIVATE KEY-----\\nMIIEogIBAAKCAQEAqjAfE7/ToSl6okR9YhQSz/vw+NvQUU4i/p+7HKqCoLC+l+rY\\n4HzU7m7fvgfkPZtTItgaiIKURo/Y+UmIaN+JNvnaTPZTXRgkfdccdTvYOeYzHn77\\nAzFdsjCcystdpXdJ/wrtfABa5XGNrFMeXYpf9Gxp+1A56modo8wiWAS48rysfzS8\\nvHUwcYuog5Nxttg5Kaqs2w1eUxiPLsXp+mN5CIhP46m1U5cisq9J1utAbKaHeRO+\\npWFnAu3e4AK1utaxhGdYCDMX3xwlNWFIgpNepFpdrs+H8btwsoPUkbwdYM+lFkGR\\nI8XZ4dMwQEI6p6iohthk6o9GQj+otOEqkGHyHQIDAQABAoIBAEA02XSk+V2i/X48\\noqUe95356gapP2V9Ohyf/IKrHY8sPyunUV0YG2k88TKLXaOUdv/9Ub1QrkoUuQIL\\nqOgP9X+FMcO5ZugHVLUZM8ZS5pepbn3B4EdrF3NDfdPQd6sWXxdWcxRGOgS4G3/4\\n98rIirz3LeC/eqoikL4cJJTaa0BAKpmxy2A5zM+2mtWjte8kLFmll1dEudVgwjI7\\nGISpw0vQQ6m6Odeu8rShW3EfXK9+VVQcOyKOEoNH9UlhugtkJvBDX/HZQylebHDX\\nKlT5Zs3uzpTsBKUC/fdv6S2vMl5YTvTbXPx/v5IOXhFsP/dMZ0Pw9/+feh5ILyo8\\ni0sh1o0CgYEAx7tuqP32bFajabqPzY4xz1LYo8rb15orilbNs3saxpAmo3h+CkpC\\nHCyyeBcnW9jRCSDGkFY+hN9zdyLMVp+2oURY3y/8iH5AQPfHHxJ8pcJNbpcL0xjk\\n6YfcEFwGbjz/c8sNYLds4Vj+/Ztn1l36LGq3hxPc9mz56J/T8+9JtUcCgYEA2iH6\\nzwu7wAU2JGvyolvWrZ4aOQ4UoKMX/fKzR32NTW5VG/te0h+TD93xeIEAfwcqcE+R\\nPq5L94ashcWxLLO1Ykeh4NY2cjHOB5iwgPfcs4bOQar9mXs7cpbIRXOHZtVC+QIf\\nbmSIeUmD9teQK8fsOsXzVKys7jVvsHbWN+f533sCgYBp2y7lJeRquuhU6um4Ofqw\\nNOpYtPUbKbyVzzeqPj6Mqm2wCfZTIdQz9oSIHU1g4mK3gcV7ThpIdd1OcQT3jCJR\\nClZHw7kF5lPLmwbPsw6ZE4fSav97XCBGnRjHGt7VokKJbj6i2dQ9AtvyMWExPhGP\\nZOfsEVz0xgEVh0/gYxwflQKBgEaKupt+hechSFMa/cp3vMNE3+IXpaAVXkrn1AaI\\nmcuypJ95+T4mq95zvgVkYo6x+I19jdEheLvBt8VDG3sDWuM/myAk5lKjvKdEP8BV\\nZ/A+5jDiZUTfT2hNTtL2+5DL4u64OwXuSRxAJAcNyzf9XW3cWkbF2N7oNQhyRfCq\\nempFAoGAAMjporbd5G7n1FjgAzZvlIJ2hIXUtCwIpIgRobEZYY9VQ9Y3OB9OAd0O\\nVDJbJqlhx0UUsDdyelW+L6NNADhF1Sq//VzL1uCPdKsM6Pn/UoxCbyswzycC+vo2\\n7co3GwupojSToXvidAYTYKupGoU5Qj227c7YGsU69qhfFq5T+Uc=\\n-----END RSA PRIVATE KEY-----\\n",
					"cert_region":    "ap-southeast-1",
					"custom_headers": "{\\\"3444\\\":\\\"5555\\\",\\\"666\\\":\\\"$ReqClientPort\\\",\\\"77777\\\":\\\"$ReqClientIP\\\"}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rs_type":        "0",
						"proxy_types.#":  "2",
						"real_servers.#": "1",
						"domain":         "testcert.qq.com",
						"instance_ids.#": "1",
						"https_ext":      CHECKSET,
						"cert_name":      CHECKSET,
						"cert":           "-----BEGIN CERTIFICATE-----\nMIID5TCCAs2gAwIBAgIRAJ8XVm8/e0sNiF7SQKrNVpUwDQYJKoZIhvcNAQELBQAw\nXjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU\nZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w\nHhcNMjUwOTI5MDY1NDI3WhcNMjYwOTI5MDY1NDI3WjAnMQswCQYDVQQGEwJDTjEY\nMBYGA1UEAxMPdGVzdGNlcnQucXEuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A\nMIIBCgKCAQEAqjAfE7/ToSl6okR9YhQSz/vw+NvQUU4i/p+7HKqCoLC+l+rY4HzU\n7m7fvgfkPZtTItgaiIKURo/Y+UmIaN+JNvnaTPZTXRgkfdccdTvYOeYzHn77AzFd\nsjCcystdpXdJ/wrtfABa5XGNrFMeXYpf9Gxp+1A56modo8wiWAS48rysfzS8vHUw\ncYuog5Nxttg5Kaqs2w1eUxiPLsXp+mN5CIhP46m1U5cisq9J1utAbKaHeRO+pWFn\nAu3e4AK1utaxhGdYCDMX3xwlNWFIgpNepFpdrs+H8btwsoPUkbwdYM+lFkGRI8XZ\n4dMwQEI6p6iohthk6o9GQj+otOEqkGHyHQIDAQABo4HUMIHRMA4GA1UdDwEB/wQE\nAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwHwYDVR0jBBgwFoAU\nKIEmBdE0Gj/Bcw+7k88VHD8Dv38wYwYIKwYBBQUHAQEEVzBVMCEGCCsGAQUFBzAB\nhhVodHRwOi8vb2NzcC5teXNzbC5jb20wMAYIKwYBBQUHMAKGJGh0dHA6Ly9jYS5t\neXNzbC5jb20vbXlzc2x0ZXN0cnNhLmNydDAaBgNVHREEEzARgg90ZXN0Y2VydC5x\ncS5jb20wDQYJKoZIhvcNAQELBQADggEBAKN/GA4FZGLzaEGs+8zbNkCbNaq/zs+V\nDBmdWF124wn+xke7PMN6PByWHBWFg7vJIfN7J7O4cFuShz23qcKj3qnDbUNKuK0X\nTjGOJcFqSaHydlrPnmVSKEIM2J5O+VyR2gjuP+J6GvJavr0hFLRXB5gVmDfpxX6P\nQX3SOilfn1rM0L7x1UG9t3t1USTumAgOuV/rflhOZs88nDjllNiR7RzsczBvrjes\n5dP+mBcCVy9PUcVHJqNNNvbR4Vly8Ki+3c10jkxR9fVMnNpi2uI/eya1SLKGM6yr\nHxEBkAxaANt8vEcCyiEBraqUvf9gBydvRKG0V6gVeeZ9+5WJd84WBzY=\n-----END CERTIFICATE-----\n",
						"key":            "-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEAqjAfE7/ToSl6okR9YhQSz/vw+NvQUU4i/p+7HKqCoLC+l+rY\n4HzU7m7fvgfkPZtTItgaiIKURo/Y+UmIaN+JNvnaTPZTXRgkfdccdTvYOeYzHn77\nAzFdsjCcystdpXdJ/wrtfABa5XGNrFMeXYpf9Gxp+1A56modo8wiWAS48rysfzS8\nvHUwcYuog5Nxttg5Kaqs2w1eUxiPLsXp+mN5CIhP46m1U5cisq9J1utAbKaHeRO+\npWFnAu3e4AK1utaxhGdYCDMX3xwlNWFIgpNepFpdrs+H8btwsoPUkbwdYM+lFkGR\nI8XZ4dMwQEI6p6iohthk6o9GQj+otOEqkGHyHQIDAQABAoIBAEA02XSk+V2i/X48\noqUe95356gapP2V9Ohyf/IKrHY8sPyunUV0YG2k88TKLXaOUdv/9Ub1QrkoUuQIL\nqOgP9X+FMcO5ZugHVLUZM8ZS5pepbn3B4EdrF3NDfdPQd6sWXxdWcxRGOgS4G3/4\n98rIirz3LeC/eqoikL4cJJTaa0BAKpmxy2A5zM+2mtWjte8kLFmll1dEudVgwjI7\nGISpw0vQQ6m6Odeu8rShW3EfXK9+VVQcOyKOEoNH9UlhugtkJvBDX/HZQylebHDX\nKlT5Zs3uzpTsBKUC/fdv6S2vMl5YTvTbXPx/v5IOXhFsP/dMZ0Pw9/+feh5ILyo8\ni0sh1o0CgYEAx7tuqP32bFajabqPzY4xz1LYo8rb15orilbNs3saxpAmo3h+CkpC\nHCyyeBcnW9jRCSDGkFY+hN9zdyLMVp+2oURY3y/8iH5AQPfHHxJ8pcJNbpcL0xjk\n6YfcEFwGbjz/c8sNYLds4Vj+/Ztn1l36LGq3hxPc9mz56J/T8+9JtUcCgYEA2iH6\nzwu7wAU2JGvyolvWrZ4aOQ4UoKMX/fKzR32NTW5VG/te0h+TD93xeIEAfwcqcE+R\nPq5L94ashcWxLLO1Ykeh4NY2cjHOB5iwgPfcs4bOQar9mXs7cpbIRXOHZtVC+QIf\nbmSIeUmD9teQK8fsOsXzVKys7jVvsHbWN+f533sCgYBp2y7lJeRquuhU6um4Ofqw\nNOpYtPUbKbyVzzeqPj6Mqm2wCfZTIdQz9oSIHU1g4mK3gcV7ThpIdd1OcQT3jCJR\nClZHw7kF5lPLmwbPsw6ZE4fSav97XCBGnRjHGt7VokKJbj6i2dQ9AtvyMWExPhGP\nZOfsEVz0xgEVh0/gYxwflQKBgEaKupt+hechSFMa/cp3vMNE3+IXpaAVXkrn1AaI\nmcuypJ95+T4mq95zvgVkYo6x+I19jdEheLvBt8VDG3sDWuM/myAk5lKjvKdEP8BV\nZ/A+5jDiZUTfT2hNTtL2+5DL4u64OwXuSRxAJAcNyzf9XW3cWkbF2N7oNQhyRfCq\nempFAoGAAMjporbd5G7n1FjgAzZvlIJ2hIXUtCwIpIgRobEZYY9VQ9Y3OB9OAd0O\nVDJbJqlhx0UUsDdyelW+L6NNADhF1Sq//VzL1uCPdKsM6Pn/UoxCbyswzycC+vo2\n7co3GwupojSToXvidAYTYKupGoU5Qj227c7YGsU69qhfFq5T+Uc=\n-----END RSA PRIVATE KEY-----\n",
						"cert_region":    "ap-southeast-1",
						"custom_headers": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"real_servers": []string{
						"1.1.1.1", "2.2.2.2", "3.3.3.3"},
					"https_ext":      "{\\\"Https2http\\\":0,\\\"Http2\\\":0,\\\"Http2https\\\":0}",
					"cert_name":      "146",
					"cert":           "-----BEGIN CERTIFICATE-----\\nMIID5DCCAsygAwIBAgIQXWOU7ILVRC20+c/4dgiiJDANBgkqhkiG9w0BAQsFADBe\\nMQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl\\nc3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe\\nFw0yNTA5MjkwNjU1MThaFw0yNjA5MjkwNjU1MThaMCcxCzAJBgNVBAYTAkNOMRgw\\nFgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw\\nggEKAoIBAQDe24eCDI+IDalD4pYU4NuvS5/S40Hwu+B3yKB9KN1HL2pTEH7kOYG/\\nRQHvkDUpsCyx0H4inpzVWu36qXAEXj9QLCNFN3+SdlBZBKaaWb76xExXPtlmyURV\\nXXlmMSSx7MeqS2euOmeIBd7WqsPX4kIok71E8tnMaLqlNEnx9SmksZukLQYusZC5\\nuECr+kgEsCDFf4JM1ZYADNL1csc1HzFPmUzZyOl99ZxAdrVNuDhc/SKdpHb80FhV\\nRBv72WW1JDYYbkP9dZALUezR66uF0/+bPzmxi2QlSu8q4FfgCSCRV2GLRLNpxMTw\\nzOrwXen4gQCzm1jTpWsb1K/n1CPHfe3TAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD\\nAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo\\ngSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG\\nFWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15\\nc3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx\\nLmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAqlpc+4pcaKbUdSdWCIC2vQx40D/51AhX\\nMQk1bKOANo/2PGbWLcYPJ0QLiE/5umKz0cRDEcq4oVjdmJ1VUvUC986/cdnFlp+l\\nnbujrCby3dJFSsG/ZWcVGfcr1Ioy3LzcEnNto2SEUrj6lMCuyCI6zHeSHAGokxAp\\n8RpbhYooisAiDj74g7Xx2/vJgOMFw6i6KhvtHYwct0JpNNsdE5EHgAJ3FWLOqKQs\\n2D/YC7I72ae3X9IGHPJjzMiQiBAfh1BKZBYNeo3UVZLmidsmhoN9QLnKCTyiAciF\\nRnodqdg8gSb4j0WmIoAua5nWjhz9V6N5rS2FDBmyaeSep0XY7SWMJQ==\\n-----END CERTIFICATE-----\\n",
					"key":            "-----BEGIN RSA PRIVATE KEY-----\\nMIIEpAIBAAKCAQEA3tuHggyPiA2pQ+KWFODbr0uf0uNB8Lvgd8igfSjdRy9qUxB+\\n5DmBv0UB75A1KbAssdB+Ip6c1Vrt+qlwBF4/UCwjRTd/knZQWQSmmlm++sRMVz7Z\\nZslEVV15ZjEksezHqktnrjpniAXe1qrD1+JCKJO9RPLZzGi6pTRJ8fUppLGbpC0G\\nLrGQubhAq/pIBLAgxX+CTNWWAAzS9XLHNR8xT5lM2cjpffWcQHa1Tbg4XP0inaR2\\n/NBYVUQb+9lltSQ2GG5D/XWQC1Hs0eurhdP/mz85sYtkJUrvKuBX4AkgkVdhi0Sz\\nacTE8Mzq8F3p+IEAs5tY06VrG9Sv59Qjx33t0wIDAQABAoIBAAoLR3HrVwFKMzB8\\nov7/CkMZbgV3MJ7E9F861CoGxy2+uVSdD/+Th+8y7G9pWQUvZ4H+ZCfOtkGuSAgU\\nqbztO6L8PthqoaJuhhiGIkeM5RPCitUhhAUjEH+mJ2xB8sFv2M2YC9qAGIJxfLfl\\n0BgaML223S4O6sxQzlQgXFKHX0mtbXGVVA0DSOwhTy5x+qLhp/Bx4TD6Qfvl9Pt8\\n3Ip8hRRtepV03h2W4TCZuU0HdLTgFcO7WXgl9HLNuubrf9EhzhRfBMGVGTfYzwMW\\npQHYXBRqFyq1+KFFhsiqT6K5UzyxWw3Rpwy03tsaIgukg/HRDrB2icjODDWpCSyD\\nM6KAYaECgYEA6E5+HqUJcqvOYXkiB6FwO6uKvCEh2GGdmcdjh8xi0jF4S78XG4RF\\nXx709z7kEY5eBRfb2Nnby1JzDkTG/NgDStcVYB/qQHK73EDb/jAsniRxA8nLAZfD\\ni0ytPpDX5dZok1gIbNx0koRQhlBGudHr5PoY7N1QQcvcmnVQWdW/EzcCgYEA9ZZS\\nckFbVTo+ePLbB8EgJuNR48mKLefWT0q5UUuFyC3ncRNpiJ/DEMBIRqjjWq6CdWXK\\n12Hz56CdpXS/HgZXnU+QqGHQ5l7Pz9hVjM2WIJ3w6kJVw1fBO1B11XfRxXnN5IJE\\na/ATOAXIRDrik+Rg3TCpZusBeTkg8do9QENQQEUCgYEAn/llup62OeR8U/2B5LVU\\nv5KrEFDUqNjYGg0Hyn2CU/NDPw5R0F4vE4kS8qy5jCl5L5K1j8i/Jm4Z02qjiX0M\\nD168VpzDySv5mHyFwq7UGvdHaG9vQCKNw4DDEQHX22viSg1mh+js0fUSKtxfSBl6\\nlA1yWrMxUI4d1bQR6DtcwNMCgYA8MWbtyCUZo5fyTxvuL6Cwx2Cn4xryG3PEpXz8\\ndvVIVi/24BoquXW3IlnUr8phzIn/Oj7YQZLlf9GD3zSEqGtLDFhZXPg1rqFiwRRe\\n2Xjlb7C/yhh5M4YSAquO8bpBm5QiYOdiSUp8nbYzOveT5hLzw9yRdCI9UwpHHQWM\\nPbNqvQKBgQDaxQz3ma5o348nRWTJrgLaipYveqFQLtosIMxDL1p8SHfiRZmSvUAC\\n0JJa7gD3LWoKPfTtMhR7Tfdkcw6x0WEzVNYy5xiWce9yhpD0cN+RoNl0xjxaAcXN\\n+hRL+9t7LIJr5zZoFL4Fj3h9FtSnhPLYk13AUa9WnY5Ih7SIaFD9hQ==\\n-----END RSA PRIVATE KEY-----\\n",
					"cert_region":    "cn-hangzhou",
					"custom_headers": "{\\\"22\\\":\\\"$ReqClientIP\\\",\\\"77\\\":\\\"88\\\",\\\"99\\\":\\\"$ReqClientPort\\\"}",
					"ocsp_enabled":   "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"real_servers.#": "3",
						"https_ext":      CHECKSET,
						"cert_name":      CHECKSET,
						"cert":           "-----BEGIN CERTIFICATE-----\nMIID5DCCAsygAwIBAgIQXWOU7ILVRC20+c/4dgiiJDANBgkqhkiG9w0BAQsFADBe\nMQswCQYDVQQGEwJDTjEOMAwGA1UEChMFTXlTU0wxKzApBgNVBAsTIk15U1NMIFRl\nc3QgUlNBIC0gRm9yIHRlc3QgdXNlIG9ubHkxEjAQBgNVBAMTCU15U1NMLmNvbTAe\nFw0yNTA5MjkwNjU1MThaFw0yNjA5MjkwNjU1MThaMCcxCzAJBgNVBAYTAkNOMRgw\nFgYDVQQDEw90ZXN0Y2VydC5xcS5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw\nggEKAoIBAQDe24eCDI+IDalD4pYU4NuvS5/S40Hwu+B3yKB9KN1HL2pTEH7kOYG/\nRQHvkDUpsCyx0H4inpzVWu36qXAEXj9QLCNFN3+SdlBZBKaaWb76xExXPtlmyURV\nXXlmMSSx7MeqS2euOmeIBd7WqsPX4kIok71E8tnMaLqlNEnx9SmksZukLQYusZC5\nuECr+kgEsCDFf4JM1ZYADNL1csc1HzFPmUzZyOl99ZxAdrVNuDhc/SKdpHb80FhV\nRBv72WW1JDYYbkP9dZALUezR66uF0/+bPzmxi2QlSu8q4FfgCSCRV2GLRLNpxMTw\nzOrwXen4gQCzm1jTpWsb1K/n1CPHfe3TAgMBAAGjgdQwgdEwDgYDVR0PAQH/BAQD\nAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQo\ngSYF0TQaP8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGG\nFWh0dHA6Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15\nc3NsLmNvbS9teXNzbHRlc3Ryc2EuY3J0MBoGA1UdEQQTMBGCD3Rlc3RjZXJ0LnFx\nLmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAqlpc+4pcaKbUdSdWCIC2vQx40D/51AhX\nMQk1bKOANo/2PGbWLcYPJ0QLiE/5umKz0cRDEcq4oVjdmJ1VUvUC986/cdnFlp+l\nnbujrCby3dJFSsG/ZWcVGfcr1Ioy3LzcEnNto2SEUrj6lMCuyCI6zHeSHAGokxAp\n8RpbhYooisAiDj74g7Xx2/vJgOMFw6i6KhvtHYwct0JpNNsdE5EHgAJ3FWLOqKQs\n2D/YC7I72ae3X9IGHPJjzMiQiBAfh1BKZBYNeo3UVZLmidsmhoN9QLnKCTyiAciF\nRnodqdg8gSb4j0WmIoAua5nWjhz9V6N5rS2FDBmyaeSep0XY7SWMJQ==\n-----END CERTIFICATE-----\n",
						"key":            "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA3tuHggyPiA2pQ+KWFODbr0uf0uNB8Lvgd8igfSjdRy9qUxB+\n5DmBv0UB75A1KbAssdB+Ip6c1Vrt+qlwBF4/UCwjRTd/knZQWQSmmlm++sRMVz7Z\nZslEVV15ZjEksezHqktnrjpniAXe1qrD1+JCKJO9RPLZzGi6pTRJ8fUppLGbpC0G\nLrGQubhAq/pIBLAgxX+CTNWWAAzS9XLHNR8xT5lM2cjpffWcQHa1Tbg4XP0inaR2\n/NBYVUQb+9lltSQ2GG5D/XWQC1Hs0eurhdP/mz85sYtkJUrvKuBX4AkgkVdhi0Sz\nacTE8Mzq8F3p+IEAs5tY06VrG9Sv59Qjx33t0wIDAQABAoIBAAoLR3HrVwFKMzB8\nov7/CkMZbgV3MJ7E9F861CoGxy2+uVSdD/+Th+8y7G9pWQUvZ4H+ZCfOtkGuSAgU\nqbztO6L8PthqoaJuhhiGIkeM5RPCitUhhAUjEH+mJ2xB8sFv2M2YC9qAGIJxfLfl\n0BgaML223S4O6sxQzlQgXFKHX0mtbXGVVA0DSOwhTy5x+qLhp/Bx4TD6Qfvl9Pt8\n3Ip8hRRtepV03h2W4TCZuU0HdLTgFcO7WXgl9HLNuubrf9EhzhRfBMGVGTfYzwMW\npQHYXBRqFyq1+KFFhsiqT6K5UzyxWw3Rpwy03tsaIgukg/HRDrB2icjODDWpCSyD\nM6KAYaECgYEA6E5+HqUJcqvOYXkiB6FwO6uKvCEh2GGdmcdjh8xi0jF4S78XG4RF\nXx709z7kEY5eBRfb2Nnby1JzDkTG/NgDStcVYB/qQHK73EDb/jAsniRxA8nLAZfD\ni0ytPpDX5dZok1gIbNx0koRQhlBGudHr5PoY7N1QQcvcmnVQWdW/EzcCgYEA9ZZS\nckFbVTo+ePLbB8EgJuNR48mKLefWT0q5UUuFyC3ncRNpiJ/DEMBIRqjjWq6CdWXK\n12Hz56CdpXS/HgZXnU+QqGHQ5l7Pz9hVjM2WIJ3w6kJVw1fBO1B11XfRxXnN5IJE\na/ATOAXIRDrik+Rg3TCpZusBeTkg8do9QENQQEUCgYEAn/llup62OeR8U/2B5LVU\nv5KrEFDUqNjYGg0Hyn2CU/NDPw5R0F4vE4kS8qy5jCl5L5K1j8i/Jm4Z02qjiX0M\nD168VpzDySv5mHyFwq7UGvdHaG9vQCKNw4DDEQHX22viSg1mh+js0fUSKtxfSBl6\nlA1yWrMxUI4d1bQR6DtcwNMCgYA8MWbtyCUZo5fyTxvuL6Cwx2Cn4xryG3PEpXz8\ndvVIVi/24BoquXW3IlnUr8phzIn/Oj7YQZLlf9GD3zSEqGtLDFhZXPg1rqFiwRRe\n2Xjlb7C/yhh5M4YSAquO8bpBm5QiYOdiSUp8nbYzOveT5hLzw9yRdCI9UwpHHQWM\nPbNqvQKBgQDaxQz3ma5o348nRWTJrgLaipYveqFQLtosIMxDL1p8SHfiRZmSvUAC\n0JJa7gD3LWoKPfTtMhR7Tfdkcw6x0WEzVNYy5xiWce9yhpD0cN+RoNl0xjxaAcXN\n+hRL+9t7LIJr5zZoFL4Fj3h9FtSnhPLYk13AUa9WnY5Ih7SIaFD9hQ==\n-----END RSA PRIVATE KEY-----\n",
						"cert_region":    "cn-hangzhou",
						"custom_headers": CHECKSET,
						"ocsp_enabled":   "true",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cert", "cert_identifier", "cert_region", "key"},
			},
		},
	})
}

var AlicloudDdosCooDomainResourceMap11581 = map[string]string{
	"cname": CHECKSET,
}

func AlicloudDdosCooDomainResourceBasicDependence11581(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_ddoscoo_instance" "defaultdPWcCl" {
  normal_bandwidth = "100"
  product_plan     = "0"
  product_type     = "ddosDip"
  period           = "1"
  normal_qps       = "500"
  port_count       = "5"
  function_version = "0"
  domain_count     = "10"
  bandwidth_mode   = "2"
  name             = "测试"
}


`, name)
}

// Test DdosCoo DomainResource. <<< Resource test cases, automatically generated.
