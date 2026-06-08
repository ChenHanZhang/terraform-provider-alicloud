package alicloud

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/PaesslerAG/jsonpath"
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

func init() {
	resource.AddTestSweepers("alicloud_express_connect_virtual_border_router", &resource.Sweeper{
		Name: "alicloud_express_connect_virtual_border_router",
		F:    testSweepExpressConnectVirtualBorderRouters,
		Dependencies: []string{
			"alicloud_cen_instance",
		},
	})
}

func testSweepExpressConnectVirtualBorderRouters(region string) error {
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return fmt.Errorf("error getting AliCloud client: %s", err)
	}
	client := rawClient.(*connectivity.AliyunClient)

	prefixes := []string{
		"tf-testAcc",
		"tf_testAcc",
	}

	request := make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["PageSize"] = PageSizeLarge
	request["PageNumber"] = 1
	var response interface{}
	for {
		action := "DescribeVirtualBorderRouters"
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(1*time.Minute, func() *resource.RetryError {
			response, err = client.RpcPost("Vpc", "2016-04-28", action, nil, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		if err != nil {
			log.Printf("[ERROR] %s got an error: %v", action, err)
			break
		}
		resp, err := jsonpath.Get("$.VirtualBorderRouterSet.VirtualBorderRouterType", response)
		if err != nil {
			log.Printf("[ERROR] parsing %s response got an error: %s", action, err)
			break
		}
		result, _ := resp.([]interface{})
		for _, v := range result {
			item := v.(map[string]interface{})
			vbrName := fmt.Sprint(item["Name"])
			vbrId := fmt.Sprint(item["VbrId"])
			skip := true
			if !sweepAll() {
				for _, prefix := range prefixes {
					if strings.HasPrefix(strings.ToLower(vbrName), strings.ToLower(prefix)) {
						skip = false
						break
					}
				}
				if skip {
					log.Printf("[INFO] Skipping VirtualBorderRouter: %s (%s)", vbrName, vbrId)
					continue
				}
			}
			action = "DeleteVirtualBorderRouter"
			request := map[string]interface{}{
				"VbrId":       vbrId,
				"RegionId":    client.RegionId,
				"ClientToken": buildClientToken("DeleteVirtualBorderRouter"),
			}
			runtime := util.RuntimeOptions{}
			runtime.SetAutoretry(true)
			wait := incrementalWait(3*time.Second, 3*time.Second)
			err = resource.Retry(1*time.Minute, func() *resource.RetryError {
				_, err = client.RpcPost("Vpc", "2016-04-28", action, nil, request, true)
				if err != nil {
					if NeedRetry(err) || IsExpectedErrors(err, []string{"DependencyViolation.BgpGroup"}) {
						wait()
						return resource.RetryableError(err)
					}
					return resource.NonRetryableError(err)
				}
				return nil
			})
			if err != nil {
				log.Printf("[ERROR] %s got an error: %v", action, err)
			}
		}
		if len(result) < PageSizeLarge {
			break
		}
		request["PageNumber"] = request["PageNumber"].(int) + 1
	}
	return nil
}

// lintignore: R001
func TestUnitAliCloudExpressConnectVirtualBorderRouter(t *testing.T) {
	p := Provider().(*schema.Provider).ResourcesMap
	dInit, _ := schema.InternalMap(p["alicloud_express_connect_virtual_border_router"].Schema).Data(nil, nil)
	dExisted, _ := schema.InternalMap(p["alicloud_express_connect_virtual_border_router"].Schema).Data(nil, nil)
	dInit.MarkNewResource()
	attributes := map[string]interface{}{
		"bandwidth":                  10,
		"circuit_code":               "CreateVirtualBorderRouterValue",
		"description":                "CreateVirtualBorderRouterValue",
		"enable_ipv6":                false,
		"local_gateway_ip":           "CreateVirtualBorderRouterValue",
		"local_ipv6_gateway_ip":      "CreateVirtualBorderRouterValue",
		"peer_gateway_ip":            "CreateVirtualBorderRouterValue",
		"peering_ipv6_subnet_mask":   "CreateVirtualBorderRouterValue",
		"peering_subnet_mask":        "CreateVirtualBorderRouterValue",
		"physical_connection_id":     "CreateVirtualBorderRouterValue",
		"vbr_owner_id":               "CreateVirtualBorderRouterValue",
		"virtual_border_router_name": "CreateVirtualBorderRouterValue",
		"vlan_id":                    1,
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
		// DescribeVirtualBorderRouters
		"VirtualBorderRouterSet": map[string]interface{}{
			"VirtualBorderRouterType": []interface{}{
				map[string]interface{}{
					"CircuitCode":           "CreateVirtualBorderRouterValue",
					"Description":           "CreateVirtualBorderRouterValue",
					"DetectMultiplier":      3,
					"EnableIpv6":            false,
					"LocalGatewayIp":        "CreateVirtualBorderRouterValue",
					"LocalIpv6GatewayIp":    "CreateVirtualBorderRouterValue",
					"MinRxInterval":         200,
					"MinTxInterval":         200,
					"PeerGatewayIp":         "CreateVirtualBorderRouterValue",
					"PeerIpv6GatewayIp":     "CreateVirtualBorderRouterValue",
					"PeeringIpv6SubnetMask": "CreateVirtualBorderRouterValue",
					"PeeringSubnetMask":     "CreateVirtualBorderRouterValue",
					"PhysicalConnectionId":  "CreateVirtualBorderRouterValue",
					"Status":                "CreateVirtualBorderRouterValue",
					"Name":                  "CreateVirtualBorderRouterValue",
					"VlanId":                1,
					"VbrId":                 "CreateVirtualBorderRouterValue",
				},
			},
		},
		"VbrId": "CreateVirtualBorderRouterValue",
	}
	CreateMockResponse := map[string]interface{}{
		// CreateVirtualBorderRouter
		"VbrId": "CreateVirtualBorderRouterValue",
	}
	failedResponseMock := func(errorCode string) (map[string]interface{}, error) {
		return nil, &tea.SDKError{
			Code:       String(errorCode),
			Data:       String(errorCode),
			Message:    String(errorCode),
			StatusCode: tea.Int(400),
		}
	}
	notFoundResponseMock := func(errorCode string) (map[string]interface{}, error) {
		return nil, GetNotFoundErrorFromString(GetNotFoundMessage("alicloud_express_connect_virtual_border_router", errorCode))
	}
	successResponseMock := func(operationMockResponse map[string]interface{}) (map[string]interface{}, error) {
		if len(operationMockResponse) > 0 {
			mapMerge(ReadMockResponse, operationMockResponse)
		}
		return ReadMockResponse, nil
	}

	// Create
	patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewVpcClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudExpressConnectVirtualBorderRouterCreate(dInit, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	ReadMockResponseDiff := map[string]interface{}{
		// DescribeVirtualBorderRouters Response
		"VirtualBorderRouterSet": map[string]interface{}{
			"VirtualBorderRouterType": []interface{}{
				map[string]interface{}{
					"VbrId": "CreateVirtualBorderRouterValue",
				},
			},
		},
	}
	errorCodes := []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1 // a counter used to cover retry scenario; the same below
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "CreateVirtualBorderRouter" {
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
		err := resourceAliCloudExpressConnectVirtualBorderRouterCreate(dInit, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_express_connect_virtual_border_router"].Schema).Data(dInit.State(), nil)
			for key, value := range attributes {
				dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dInit.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// Update
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewVpcClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudExpressConnectVirtualBorderRouterUpdate(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	// ModifyVirtualBorderRouterAttribute
	attributesDiff := map[string]interface{}{
		"circuit_code":                    "ModifyVirtualBorderRouterAttributeValue",
		"description":                     "ModifyVirtualBorderRouterAttributeValue",
		"detect_multiplier":               5,
		"min_rx_interval":                 300,
		"min_tx_interval":                 300,
		"enable_ipv6":                     true,
		"local_gateway_ip":                "ModifyVirtualBorderRouterAttributeValue",
		"local_ipv6_gateway_ip":           "ModifyVirtualBorderRouterAttributeValue",
		"peer_gateway_ip":                 "ModifyVirtualBorderRouterAttributeValue",
		"peer_ipv6_gateway_ip":            "ModifyVirtualBorderRouterAttributeValue",
		"peering_ipv6_subnet_mask":        "ModifyVirtualBorderRouterAttributeValue",
		"peering_subnet_mask":             "ModifyVirtualBorderRouterAttributeValue",
		"virtual_border_router_name":      "ModifyVirtualBorderRouterAttributeValue",
		"vlan_id":                         2,
		"associated_physical_connections": "ModifyVirtualBorderRouterAttributeValue",
		"bandwidth":                       20,
	}
	diff, err := newInstanceDiff("alicloud_express_connect_virtual_border_router", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_express_connect_virtual_border_router"].Schema).Data(dInit.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeVirtualBorderRouters Response
		"VirtualBorderRouterSet": map[string]interface{}{
			"VirtualBorderRouterType": []interface{}{
				map[string]interface{}{
					"CircuitCode":           "ModifyVirtualBorderRouterAttributeValue",
					"Description":           "ModifyVirtualBorderRouterAttributeValue",
					"DetectMultiplier":      5,
					"EnableIpv6":            true,
					"LocalGatewayIp":        "ModifyVirtualBorderRouterAttributeValue",
					"LocalIpv6GatewayIp":    "ModifyVirtualBorderRouterAttributeValue",
					"MinRxInterval":         300,
					"MinTxInterval":         300,
					"PeerGatewayIp":         "ModifyVirtualBorderRouterAttributeValue",
					"PeerIpv6GatewayIp":     "ModifyVirtualBorderRouterAttributeValue",
					"PeeringIpv6SubnetMask": "ModifyVirtualBorderRouterAttributeValue",
					"PeeringSubnetMask":     "ModifyVirtualBorderRouterAttributeValue",
					"Status":                "ModifyVirtualBorderRouterAttributeValue",
					"Name":                  "ModifyVirtualBorderRouterAttributeValue",
					"VlanId":                2,
				},
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "ModifyVirtualBorderRouterAttribute" {
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
		err := resourceAliCloudExpressConnectVirtualBorderRouterUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_express_connect_virtual_border_router"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// RecoverVirtualBorderRouter
	attributesDiff = map[string]interface{}{
		"status": "active",
	}
	diff, err = newInstanceDiff("alicloud_express_connect_virtual_border_router", attributes, attributesDiff, dExisted.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_express_connect_virtual_border_router"].Schema).Data(dExisted.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeVirtualBorderRouters Response
		"VirtualBorderRouterSet": map[string]interface{}{
			"VirtualBorderRouterType": []interface{}{
				map[string]interface{}{
					"Status": "active",
				},
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "RecoverVirtualBorderRouter" {
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
		err := resourceAliCloudExpressConnectVirtualBorderRouterUpdate(dExisted, rawClient)
		patches.Reset()

		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_express_connect_virtual_border_router"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// TerminateVirtualBorderRouter
	attributesDiff = map[string]interface{}{
		"status": "terminated",
	}
	diff, err = newInstanceDiff("alicloud_express_connect_virtual_border_router", attributes, attributesDiff, dExisted.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_express_connect_virtual_border_router"].Schema).Data(dExisted.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeVirtualBorderRouters Response
		"VirtualBorderRouterSet": map[string]interface{}{
			"VirtualBorderRouterType": []interface{}{
				map[string]interface{}{
					"Status": "terminated",
				},
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "TerminateVirtualBorderRouter" {
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
		err := resourceAliCloudExpressConnectVirtualBorderRouterUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_express_connect_virtual_border_router"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// Read
	errorCodes = []string{"NonRetryableError", "Throttling", "nil", "{}"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DescribeVirtualBorderRouters" {
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
		err := resourceAliCloudExpressConnectVirtualBorderRouterRead(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "{}":
			assert.Nil(t, err)
		}
	}

	// Delete
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewVpcClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudExpressConnectVirtualBorderRouterDelete(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	errorCodes = []string{"NonRetryableError", "Throttling", "DependencyViolation.BgpGroup", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DeleteVirtualBorderRouter" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if errorCodes[retryIndex] == "nil" {
						ReadMockResponse = map[string]interface{}{}
						return ReadMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudExpressConnectVirtualBorderRouterDelete(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "nil":
			assert.Nil(t, err)
		}
	}

}

// Test ExpressConnect VirtualBorderRouter. >>> Resource test cases, automatically generated.
// Case VirtualBorderRouter_v4 11964
func TestAccAliCloudExpressConnectVirtualBorderRouter_basic11964(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_express_connect_virtual_border_router.default"
	ra := resourceAttrInit(resourceId, AlicloudExpressConnectVirtualBorderRouterMap11964)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ExpressConnectServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeExpressConnectVirtualBorderRouter")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccexpressconnect%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudExpressConnectVirtualBorderRouterBasicDependence11964)
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
					"physical_connection_id":     "${var.pconn1}",
					"peer_gateway_ip":            "10.0.0.1",
					"description":                "kongyan-test",
					"virtual_border_router_name": name,
					"peering_subnet_mask":        "255.255.255.0",
					"local_gateway_ip":           "10.0.0.2",
					"vlan_id":                    "2088",
					"circuit_code":               "1024",
					"min_rx_interval":            "1000",
					"min_tx_interval":            "1000",
					"detect_multiplier":          "3",
					"enable_ipv6":                "true",
					"vbr_owner_id":               "1903286717072367",
					"resource_group_id":          "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"status":                     "active",
					"peer_ipv6_gateway_ip":       "2408:4004:cc:400::8",
					"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/55",
					"local_ipv6_gateway_ip":      "2408:4004:cc:400::7",
					"mtu":                        "1500",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"physical_connection_id":     CHECKSET,
						"peer_gateway_ip":            "10.0.0.1",
						"description":                "kongyan-test",
						"virtual_border_router_name": name,
						"peering_subnet_mask":        "255.255.255.0",
						"local_gateway_ip":           "10.0.0.2",
						"vlan_id":                    "2088",
						"circuit_code":               CHECKSET,
						"min_rx_interval":            "1000",
						"min_tx_interval":            "1000",
						"detect_multiplier":          "3",
						"enable_ipv6":                "true",
						"vbr_owner_id":               "1903286717072367",
						"resource_group_id":          CHECKSET,
						"status":                     "active",
						"peer_ipv6_gateway_ip":       "2408:4004:cc:400::8",
						"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/55",
						"local_ipv6_gateway_ip":      "2408:4004:cc:400::7",
						"mtu":                        "1500",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"virtual_border_router_name": name + "_update",
					"vlan_id":                    "2089",
					"min_rx_interval":            "600",
					"min_tx_interval":            "600",
					"detect_multiplier":          "6",
					"peer_ipv6_gateway_ip":       "2408:4004:cc:400::2",
					"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/56",
					"local_ipv6_gateway_ip":      "2408:4004:cc:400::1",
					"sitelink_enable":            "true",
					"bandwidth":                  "50",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"virtual_border_router_name": name + "_update",
						"vlan_id":                    "2089",
						"min_rx_interval":            "600",
						"min_tx_interval":            "600",
						"detect_multiplier":          "6",
						"peer_ipv6_gateway_ip":       "2408:4004:cc:400::2",
						"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/56",
						"local_ipv6_gateway_ip":      "2408:4004:cc:400::1",
						"sitelink_enable":            "true",
						"bandwidth":                  "50",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"peer_ipv6_gateway_ip":     "2408:4004:cc:400::4",
					"peering_ipv6_subnet_mask": "2408:4004:cc:400::/57",
					"local_ipv6_gateway_ip":    "2408:4004:cc:400::3",
					"sitelink_enable":          "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"peer_ipv6_gateway_ip":     "2408:4004:cc:400::4",
						"peering_ipv6_subnet_mask": "2408:4004:cc:400::/57",
						"local_ipv6_gateway_ip":    "2408:4004:cc:400::3",
						"sitelink_enable":          "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "kongyan-tes",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "kongyan-tes",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bandwidth": "100",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bandwidth": "100",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "terminated",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "terminated",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "active",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "active",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"peer_gateway_ip":     "192.168.0.2",
					"peering_subnet_mask": "255.255.0.0",
					"local_gateway_ip":    "192.168.0.1",
					"circuit_code":        "4096",
					"resource_group_id":   "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"peer_gateway_ip":     "192.168.0.2",
						"peering_subnet_mask": "255.255.0.0",
						"local_gateway_ip":    "192.168.0.1",
						"circuit_code":        CHECKSET,
						"resource_group_id":   CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF",
						"For":     "Test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "Test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF-update",
						"For":     "Test-update",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF-update",
						"tags.For":     "Test-update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "0",
						"tags.Created": REMOVEKEY,
						"tags.For":     REMOVEKEY,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"vbr_owner_id"},
			},
		},
	})
}

var AlicloudExpressConnectVirtualBorderRouterMap11964 = map[string]string{
	"route_table_id":                      CHECKSET,
	"activation_time":                     CHECKSET,
	"associated_physical_connections.#":   CHECKSET,
	"vlan_interface_id":                   CHECKSET,
	"physical_connection_status":          CHECKSET,
	"physical_connection_owner_uid":       CHECKSET,
	"termination_time":                    CHECKSET,
	"recovery_time":                       CHECKSET,
	"associated_cens.#":                   CHECKSET,
	"pconn_vbr_expire_time":               CHECKSET,
	"create_time":                         CHECKSET,
	"type":                                CHECKSET,
	"access_point_id":                     CHECKSET,
	"physical_connection_business_status": CHECKSET,
	"cloud_box_instance_id":               CHECKSET,
}

func AlicloudExpressConnectVirtualBorderRouterBasicDependence11964(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "pconn1" {
  default = "pc-bp1j6up8jcs8jogazdb6j"
}

variable "pconn2" {
  default = "pc-bp1oasymmf8oe0j92m125"
}

variable "region_id" {
  default = "cn-hangzhou"
}


`, name)
}

// Case VirtualBorderRouter_v3 11835
func TestAccAliCloudExpressConnectVirtualBorderRouter_basic11835(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_express_connect_virtual_border_router.default"
	ra := resourceAttrInit(resourceId, AlicloudExpressConnectVirtualBorderRouterMap11835)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ExpressConnectServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeExpressConnectVirtualBorderRouter")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccexpressconnect%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudExpressConnectVirtualBorderRouterBasicDependence11835)
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
					"physical_connection_id":     "${var.pconn1}",
					"peer_gateway_ip":            "10.0.0.1",
					"description":                "kongyan-test",
					"virtual_border_router_name": name,
					"peering_subnet_mask":        "255.255.255.0",
					"local_gateway_ip":           "10.0.0.2",
					"vlan_id":                    "2088",
					"circuit_code":               "1024",
					"bandwidth":                  "50",
					"min_rx_interval":            "1000",
					"min_tx_interval":            "1000",
					"detect_multiplier":          "3",
					"enable_ipv6":                "true",
					"vbr_owner_id":               "1903286717072367",
					"resource_group_id":          "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"status":                     "active",
					"peer_ipv6_gateway_ip":       "2408:4004:cc:400::8",
					"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/55",
					"local_ipv6_gateway_ip":      "2408:4004:cc:400::7",
					"mtu":                        "1500",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"physical_connection_id":     CHECKSET,
						"peer_gateway_ip":            "10.0.0.1",
						"description":                "kongyan-test",
						"virtual_border_router_name": name,
						"peering_subnet_mask":        "255.255.255.0",
						"local_gateway_ip":           "10.0.0.2",
						"vlan_id":                    "2088",
						"circuit_code":               CHECKSET,
						"bandwidth":                  "50",
						"min_rx_interval":            "1000",
						"min_tx_interval":            "1000",
						"detect_multiplier":          "3",
						"enable_ipv6":                "true",
						"vbr_owner_id":               "1903286717072367",
						"resource_group_id":          CHECKSET,
						"status":                     "active",
						"peer_ipv6_gateway_ip":       "2408:4004:cc:400::8",
						"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/55",
						"local_ipv6_gateway_ip":      "2408:4004:cc:400::7",
						"mtu":                        "1500",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"virtual_border_router_name": name + "_update",
					"vlan_id":                    "2089",
					"min_rx_interval":            "600",
					"min_tx_interval":            "600",
					"detect_multiplier":          "6",
					"peer_ipv6_gateway_ip":       "2408:4004:cc:400::2",
					"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/56",
					"local_ipv6_gateway_ip":      "2408:4004:cc:400::1",
					"sitelink_enable":            "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"virtual_border_router_name": name + "_update",
						"vlan_id":                    "2089",
						"min_rx_interval":            "600",
						"min_tx_interval":            "600",
						"detect_multiplier":          "6",
						"peer_ipv6_gateway_ip":       "2408:4004:cc:400::2",
						"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/56",
						"local_ipv6_gateway_ip":      "2408:4004:cc:400::1",
						"sitelink_enable":            "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"peer_ipv6_gateway_ip":     "2408:4004:cc:400::4",
					"peering_ipv6_subnet_mask": "2408:4004:cc:400::/57",
					"local_ipv6_gateway_ip":    "2408:4004:cc:400::3",
					"sitelink_enable":          "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"peer_ipv6_gateway_ip":     "2408:4004:cc:400::4",
						"peering_ipv6_subnet_mask": "2408:4004:cc:400::/57",
						"local_ipv6_gateway_ip":    "2408:4004:cc:400::3",
						"sitelink_enable":          "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "kongyan-tes",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "kongyan-tes",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bandwidth": "100",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bandwidth": "100",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "terminated",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "terminated",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "active",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "active",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"peer_gateway_ip":     "192.168.0.2",
					"peering_subnet_mask": "255.255.0.0",
					"local_gateway_ip":    "192.168.0.1",
					"circuit_code":        "4096",
					"resource_group_id":   "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"peer_gateway_ip":     "192.168.0.2",
						"peering_subnet_mask": "255.255.0.0",
						"local_gateway_ip":    "192.168.0.1",
						"circuit_code":        CHECKSET,
						"resource_group_id":   CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF",
						"For":     "Test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "Test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF-update",
						"For":     "Test-update",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF-update",
						"tags.For":     "Test-update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "0",
						"tags.Created": REMOVEKEY,
						"tags.For":     REMOVEKEY,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"vbr_owner_id"},
			},
		},
	})
}

var AlicloudExpressConnectVirtualBorderRouterMap11835 = map[string]string{
	"route_table_id":                      CHECKSET,
	"activation_time":                     CHECKSET,
	"associated_physical_connections.#":   CHECKSET,
	"vlan_interface_id":                   CHECKSET,
	"physical_connection_status":          CHECKSET,
	"physical_connection_owner_uid":       CHECKSET,
	"termination_time":                    CHECKSET,
	"recovery_time":                       CHECKSET,
	"associated_cens.#":                   CHECKSET,
	"pconn_vbr_expire_time":               CHECKSET,
	"create_time":                         CHECKSET,
	"type":                                CHECKSET,
	"access_point_id":                     CHECKSET,
	"physical_connection_business_status": CHECKSET,
	"cloud_box_instance_id":               CHECKSET,
}

func AlicloudExpressConnectVirtualBorderRouterBasicDependence11835(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "pconn1" {
  default = "pc-bp1j6up8jcs8jogazdb6j"
}

variable "pconn2" {
  default = "pc-bp1oasymmf8oe0j92m125"
}

variable "region_id" {
  default = "cn-hangzhou"
}


`, name)
}

// Case VirtualBorderRouter_v2 11568
func TestAccAliCloudExpressConnectVirtualBorderRouter_basic11568(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_express_connect_virtual_border_router.default"
	ra := resourceAttrInit(resourceId, AlicloudExpressConnectVirtualBorderRouterMap11568)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ExpressConnectServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeExpressConnectVirtualBorderRouter")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccexpressconnect%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudExpressConnectVirtualBorderRouterBasicDependence11568)
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
					"physical_connection_id":     "${var.pconn1}",
					"peer_gateway_ip":            "10.0.0.1",
					"description":                "kongyan-test",
					"virtual_border_router_name": name,
					"peering_subnet_mask":        "255.255.255.0",
					"local_gateway_ip":           "10.0.0.2",
					"vlan_id":                    "2088",
					"circuit_code":               "1024",
					"bandwidth":                  "50",
					"min_rx_interval":            "1000",
					"min_tx_interval":            "1000",
					"detect_multiplier":          "3",
					"enable_ipv6":                "true",
					"vbr_owner_id":               "1903286717072367",
					"resource_group_id":          "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"status":                     "active",
					"peer_ipv6_gateway_ip":       "2408:4004:cc:400::8",
					"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/55",
					"local_ipv6_gateway_ip":      "2408:4004:cc:400::7",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"physical_connection_id":     CHECKSET,
						"peer_gateway_ip":            "10.0.0.1",
						"description":                "kongyan-test",
						"virtual_border_router_name": name,
						"peering_subnet_mask":        "255.255.255.0",
						"local_gateway_ip":           "10.0.0.2",
						"vlan_id":                    "2088",
						"circuit_code":               CHECKSET,
						"bandwidth":                  "50",
						"min_rx_interval":            "1000",
						"min_tx_interval":            "1000",
						"detect_multiplier":          "3",
						"enable_ipv6":                "true",
						"vbr_owner_id":               "1903286717072367",
						"resource_group_id":          CHECKSET,
						"status":                     "active",
						"peer_ipv6_gateway_ip":       "2408:4004:cc:400::8",
						"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/55",
						"local_ipv6_gateway_ip":      "2408:4004:cc:400::7",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"virtual_border_router_name": name + "_update",
					"vlan_id":                    "2089",
					"min_rx_interval":            "600",
					"min_tx_interval":            "600",
					"detect_multiplier":          "6",
					"peer_ipv6_gateway_ip":       "2408:4004:cc:400::2",
					"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/56",
					"local_ipv6_gateway_ip":      "2408:4004:cc:400::1",
					"sitelink_enable":            "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"virtual_border_router_name": name + "_update",
						"vlan_id":                    "2089",
						"min_rx_interval":            "600",
						"min_tx_interval":            "600",
						"detect_multiplier":          "6",
						"peer_ipv6_gateway_ip":       "2408:4004:cc:400::2",
						"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/56",
						"local_ipv6_gateway_ip":      "2408:4004:cc:400::1",
						"sitelink_enable":            "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"peer_ipv6_gateway_ip":     "2408:4004:cc:400::4",
					"peering_ipv6_subnet_mask": "2408:4004:cc:400::/57",
					"local_ipv6_gateway_ip":    "2408:4004:cc:400::3",
					"sitelink_enable":          "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"peer_ipv6_gateway_ip":     "2408:4004:cc:400::4",
						"peering_ipv6_subnet_mask": "2408:4004:cc:400::/57",
						"local_ipv6_gateway_ip":    "2408:4004:cc:400::3",
						"sitelink_enable":          "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "kongyan-tes",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "kongyan-tes",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bandwidth": "100",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bandwidth": "100",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "terminated",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "terminated",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "active",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "active",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"peer_gateway_ip":     "192.168.0.2",
					"peering_subnet_mask": "255.255.0.0",
					"local_gateway_ip":    "192.168.0.1",
					"circuit_code":        "4096",
					"resource_group_id":   "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"peer_gateway_ip":     "192.168.0.2",
						"peering_subnet_mask": "255.255.0.0",
						"local_gateway_ip":    "192.168.0.1",
						"circuit_code":        CHECKSET,
						"resource_group_id":   CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF",
						"For":     "Test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "Test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF-update",
						"For":     "Test-update",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF-update",
						"tags.For":     "Test-update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "0",
						"tags.Created": REMOVEKEY,
						"tags.For":     REMOVEKEY,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"vbr_owner_id"},
			},
		},
	})
}

var AlicloudExpressConnectVirtualBorderRouterMap11568 = map[string]string{
	"route_table_id":                      CHECKSET,
	"activation_time":                     CHECKSET,
	"associated_physical_connections.#":   CHECKSET,
	"vlan_interface_id":                   CHECKSET,
	"physical_connection_status":          CHECKSET,
	"physical_connection_owner_uid":       CHECKSET,
	"termination_time":                    CHECKSET,
	"recovery_time":                       CHECKSET,
	"associated_cens.#":                   CHECKSET,
	"pconn_vbr_expire_time":               CHECKSET,
	"create_time":                         CHECKSET,
	"type":                                CHECKSET,
	"access_point_id":                     CHECKSET,
	"physical_connection_business_status": CHECKSET,
	"cloud_box_instance_id":               CHECKSET,
}

func AlicloudExpressConnectVirtualBorderRouterBasicDependence11568(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "pconn1" {
  default = "pc-bp1j6up8jcs8jogazdb6j"
}

variable "pconn2" {
  default = "pc-bp1oasymmf8oe0j92m125"
}

variable "region_id" {
  default = "cn-hangzhou"
}


`, name)
}

// Case VirtualBorderRouter_202509_小浒_副本 11493
func TestAccAliCloudExpressConnectVirtualBorderRouter_basic11493(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_express_connect_virtual_border_router.default"
	ra := resourceAttrInit(resourceId, AlicloudExpressConnectVirtualBorderRouterMap11493)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ExpressConnectServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeExpressConnectVirtualBorderRouter")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccexpressconnect%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudExpressConnectVirtualBorderRouterBasicDependence11493)
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
					"physical_connection_id":     "${var.pconn1}",
					"peer_gateway_ip":            "10.0.0.1",
					"description":                "kongyan-test",
					"virtual_border_router_name": name,
					"peering_subnet_mask":        "255.255.255.0",
					"local_gateway_ip":           "10.0.0.2",
					"vlan_id":                    "2088",
					"circuit_code":               "1024",
					"bandwidth":                  "50",
					"min_rx_interval":            "1000",
					"min_tx_interval":            "1000",
					"detect_multiplier":          "3",
					"enable_ipv6":                "true",
					"vbr_owner_id":               "1903286717072367",
					"resource_group_id":          "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"status":                     "active",
					"peer_ipv6_gateway_ip":       "2408:4004:cc:400::8",
					"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/55",
					"local_ipv6_gateway_ip":      "2408:4004:cc:400::7",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"physical_connection_id":     CHECKSET,
						"peer_gateway_ip":            "10.0.0.1",
						"description":                "kongyan-test",
						"virtual_border_router_name": name,
						"peering_subnet_mask":        "255.255.255.0",
						"local_gateway_ip":           "10.0.0.2",
						"vlan_id":                    "2088",
						"circuit_code":               CHECKSET,
						"bandwidth":                  "50",
						"min_rx_interval":            "1000",
						"min_tx_interval":            "1000",
						"detect_multiplier":          "3",
						"enable_ipv6":                "true",
						"vbr_owner_id":               "1903286717072367",
						"resource_group_id":          CHECKSET,
						"status":                     "active",
						"peer_ipv6_gateway_ip":       "2408:4004:cc:400::8",
						"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/55",
						"local_ipv6_gateway_ip":      "2408:4004:cc:400::7",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"virtual_border_router_name": name + "_update",
					"vlan_id":                    "2089",
					"min_rx_interval":            "600",
					"min_tx_interval":            "600",
					"detect_multiplier":          "6",
					"peer_ipv6_gateway_ip":       "2408:4004:cc:400::2",
					"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/56",
					"local_ipv6_gateway_ip":      "2408:4004:cc:400::1",
					"sitelink_enable":            "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"virtual_border_router_name": name + "_update",
						"vlan_id":                    "2089",
						"min_rx_interval":            "600",
						"min_tx_interval":            "600",
						"detect_multiplier":          "6",
						"peer_ipv6_gateway_ip":       "2408:4004:cc:400::2",
						"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/56",
						"local_ipv6_gateway_ip":      "2408:4004:cc:400::1",
						"sitelink_enable":            "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"peer_ipv6_gateway_ip":     "2408:4004:cc:400::4",
					"peering_ipv6_subnet_mask": "2408:4004:cc:400::/57",
					"local_ipv6_gateway_ip":    "2408:4004:cc:400::3",
					"sitelink_enable":          "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"peer_ipv6_gateway_ip":     "2408:4004:cc:400::4",
						"peering_ipv6_subnet_mask": "2408:4004:cc:400::/57",
						"local_ipv6_gateway_ip":    "2408:4004:cc:400::3",
						"sitelink_enable":          "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "kongyan-tes",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "kongyan-tes",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bandwidth": "100",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bandwidth": "100",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "terminated",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "terminated",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "active",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "active",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"peer_gateway_ip":     "192.168.0.2",
					"peering_subnet_mask": "255.255.0.0",
					"local_gateway_ip":    "192.168.0.1",
					"circuit_code":        "4096",
					"resource_group_id":   "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"peer_gateway_ip":     "192.168.0.2",
						"peering_subnet_mask": "255.255.0.0",
						"local_gateway_ip":    "192.168.0.1",
						"circuit_code":        CHECKSET,
						"resource_group_id":   CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF",
						"For":     "Test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "Test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF-update",
						"For":     "Test-update",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF-update",
						"tags.For":     "Test-update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "0",
						"tags.Created": REMOVEKEY,
						"tags.For":     REMOVEKEY,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"vbr_owner_id"},
			},
		},
	})
}

var AlicloudExpressConnectVirtualBorderRouterMap11493 = map[string]string{
	"route_table_id":                      CHECKSET,
	"activation_time":                     CHECKSET,
	"associated_physical_connections.#":   CHECKSET,
	"vlan_interface_id":                   CHECKSET,
	"physical_connection_status":          CHECKSET,
	"physical_connection_owner_uid":       CHECKSET,
	"termination_time":                    CHECKSET,
	"recovery_time":                       CHECKSET,
	"associated_cens.#":                   CHECKSET,
	"pconn_vbr_expire_time":               CHECKSET,
	"create_time":                         CHECKSET,
	"type":                                CHECKSET,
	"access_point_id":                     CHECKSET,
	"physical_connection_business_status": CHECKSET,
	"cloud_box_instance_id":               CHECKSET,
}

func AlicloudExpressConnectVirtualBorderRouterBasicDependence11493(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "pconn1" {
  default = "pc-bp1j6up8jcs8jogazdb6j"
}

variable "pconn2" {
  default = "pc-bp1oasymmf8oe0j92m125"
}

variable "region_id" {
  default = "cn-hangzhou"
}


`, name)
}

// Case VirtualBorderRouter_202509_小浒 11440
func TestAccAliCloudExpressConnectVirtualBorderRouter_basic11440(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_express_connect_virtual_border_router.default"
	ra := resourceAttrInit(resourceId, AlicloudExpressConnectVirtualBorderRouterMap11440)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ExpressConnectServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeExpressConnectVirtualBorderRouter")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccexpressconnect%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudExpressConnectVirtualBorderRouterBasicDependence11440)
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
					"physical_connection_id":     "${var.pconn1}",
					"peer_gateway_ip":            "10.0.0.1",
					"description":                "kongyan-test",
					"virtual_border_router_name": name,
					"peering_subnet_mask":        "255.255.255.0",
					"local_gateway_ip":           "10.0.0.2",
					"vlan_id":                    "2088",
					"circuit_code":               "1024",
					"bandwidth":                  "50",
					"min_rx_interval":            "1000",
					"min_tx_interval":            "1000",
					"detect_multiplier":          "3",
					"enable_ipv6":                "true",
					"vbr_owner_id":               "1903286717072367",
					"resource_group_id":          "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"status":                     "active",
					"peer_ipv6_gateway_ip":       "2408:4004:cc:400::8",
					"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/55",
					"local_ipv6_gateway_ip":      "2408:4004:cc:400::7",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"physical_connection_id":     CHECKSET,
						"peer_gateway_ip":            "10.0.0.1",
						"description":                "kongyan-test",
						"virtual_border_router_name": name,
						"peering_subnet_mask":        "255.255.255.0",
						"local_gateway_ip":           "10.0.0.2",
						"vlan_id":                    "2088",
						"circuit_code":               CHECKSET,
						"bandwidth":                  "50",
						"min_rx_interval":            "1000",
						"min_tx_interval":            "1000",
						"detect_multiplier":          "3",
						"enable_ipv6":                "true",
						"vbr_owner_id":               "1903286717072367",
						"resource_group_id":          CHECKSET,
						"status":                     "active",
						"peer_ipv6_gateway_ip":       "2408:4004:cc:400::8",
						"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/55",
						"local_ipv6_gateway_ip":      "2408:4004:cc:400::7",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"virtual_border_router_name": name + "_update",
					"vlan_id":                    "2089",
					"min_rx_interval":            "600",
					"min_tx_interval":            "600",
					"detect_multiplier":          "6",
					"peer_ipv6_gateway_ip":       "2408:4004:cc:400::2",
					"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/56",
					"local_ipv6_gateway_ip":      "2408:4004:cc:400::1",
					"sitelink_enable":            "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"virtual_border_router_name": name + "_update",
						"vlan_id":                    "2089",
						"min_rx_interval":            "600",
						"min_tx_interval":            "600",
						"detect_multiplier":          "6",
						"peer_ipv6_gateway_ip":       "2408:4004:cc:400::2",
						"peering_ipv6_subnet_mask":   "2408:4004:cc:400::/56",
						"local_ipv6_gateway_ip":      "2408:4004:cc:400::1",
						"sitelink_enable":            "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"peer_ipv6_gateway_ip":     "2408:4004:cc:400::4",
					"peering_ipv6_subnet_mask": "2408:4004:cc:400::/57",
					"local_ipv6_gateway_ip":    "2408:4004:cc:400::3",
					"sitelink_enable":          "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"peer_ipv6_gateway_ip":     "2408:4004:cc:400::4",
						"peering_ipv6_subnet_mask": "2408:4004:cc:400::/57",
						"local_ipv6_gateway_ip":    "2408:4004:cc:400::3",
						"sitelink_enable":          "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "kongyan-tes",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "kongyan-tes",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bandwidth": "100",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bandwidth": "100",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "terminated",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "terminated",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "active",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "active",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"peer_gateway_ip":     "192.168.0.2",
					"peering_subnet_mask": "255.255.0.0",
					"local_gateway_ip":    "192.168.0.1",
					"circuit_code":        "4096",
					"resource_group_id":   "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"peer_gateway_ip":     "192.168.0.2",
						"peering_subnet_mask": "255.255.0.0",
						"local_gateway_ip":    "192.168.0.1",
						"circuit_code":        CHECKSET,
						"resource_group_id":   CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF",
						"For":     "Test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "Test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF-update",
						"For":     "Test-update",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF-update",
						"tags.For":     "Test-update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "0",
						"tags.Created": REMOVEKEY,
						"tags.For":     REMOVEKEY,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"vbr_owner_id"},
			},
		},
	})
}

var AlicloudExpressConnectVirtualBorderRouterMap11440 = map[string]string{
	"route_table_id":                      CHECKSET,
	"activation_time":                     CHECKSET,
	"associated_physical_connections.#":   CHECKSET,
	"vlan_interface_id":                   CHECKSET,
	"physical_connection_status":          CHECKSET,
	"physical_connection_owner_uid":       CHECKSET,
	"termination_time":                    CHECKSET,
	"recovery_time":                       CHECKSET,
	"associated_cens.#":                   CHECKSET,
	"pconn_vbr_expire_time":               CHECKSET,
	"create_time":                         CHECKSET,
	"type":                                CHECKSET,
	"access_point_id":                     CHECKSET,
	"physical_connection_business_status": CHECKSET,
	"cloud_box_instance_id":               CHECKSET,
}

func AlicloudExpressConnectVirtualBorderRouterBasicDependence11440(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "pconn1" {
  default = "pc-bp1j6up8jcs8jogazdb6j"
}

variable "pconn2" {
  default = "pc-bp1oasymmf8oe0j92m125"
}

variable "region_id" {
  default = "cn-hangzhou"
}


`, name)
}

// Test ExpressConnect VirtualBorderRouter. <<< Resource test cases, automatically generated.
