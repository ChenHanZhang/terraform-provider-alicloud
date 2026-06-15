package alicloud

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

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
	resource.AddTestSweepers("alicloud_fnf_flow", &resource.Sweeper{
		Name:         "alicloud_fnf_flow",
		F:            testSweepFnfFlow,
		Dependencies: []string{"alicloud_fnf_schedule"},
	})
}

func testSweepFnfFlow(region string) error {
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return WrapErrorf(err, "error getting Alicloud client.")
	}
	support := false
	for _, v := range connectivity.FnfSupportRegions {
		if v == connectivity.Region(region) {
			support = true
			break
		}
	}
	if !support {
		return nil
	}
	client := rawClient.(*connectivity.AliyunClient)

	prefixes := []string{
		"tf-testAcc",
		"tf_testAcc",
	}
	action := "ListFlows"
	request := make(map[string]interface{})
	var response map[string]interface{}
	response, err = client.RpcGet("fnf", "2019-03-15", action, request, nil)
	if err != nil {
		return WrapErrorf(err, DataDefaultErrorMsg, "alicloud_fnf_flows", action, AlibabaCloudSdkGoERROR)
	}
	addDebug(action, response, request)
	resp, err := jsonpath.Get("$.Flows", response)
	if err != nil {
		return WrapErrorf(err, FailedGetAttributeMsg, action, "$.Flows", response)
	}
	result, _ := resp.([]interface{})
	for _, v := range result {
		item := v.(map[string]interface{})
		name := item["Name"].(string)
		skip := true
		for _, prefix := range prefixes {
			if strings.HasPrefix(name, prefix) {
				skip = false
				break
			}
		}
		if skip {
			log.Printf("[INFO] Skipping Fnf Flow: %s ", name)
			continue
		}
		log.Printf("[Info] Delete Fnf Flow: %s", name)

		action := "DeleteFlow"
		request := map[string]interface{}{
			"Name": name,
		}
		_, err = client.RpcGet("fnf", "2019-03-15", action, request, nil)
		if err != nil {
			log.Printf("[ERROR] Failed to delete Fnf Flow (%s): %s", name, err)
		}
	}
	return nil
}

// Test FnF Flow. >>> Resource test cases, automatically generated.
// Case FnF下Flow测试用例 21
func TestAccAliCloudFnFFlow_basic21(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_fnf_flow.default"
	ra := resourceAttrInit(resourceId, AlicloudFnFFlowMap21)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &FnFServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeFnFFlow")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccfnf%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudFnFFlowBasicDependence21)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
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
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudFnFFlowMap21 = map[string]string{
	"create_time":             CHECKSET,
	"tracing_configuration.#": CHECKSET,
	"flow_id":                 CHECKSET,
	"last_modified_time":      CHECKSET,
}

func AlicloudFnFFlowBasicDependence21(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case FnF下Flow测试用例 361
func TestAccAliCloudFnFFlow_basic361(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_fnf_flow.default"
	ra := resourceAttrInit(resourceId, AlicloudFnFFlowMap361)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &FnFServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeFnFFlow")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccfnf%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudFnFFlowBasicDependence361)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
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
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudFnFFlowMap361 = map[string]string{
	"create_time":             CHECKSET,
	"tracing_configuration.#": CHECKSET,
	"flow_id":                 CHECKSET,
	"last_modified_time":      CHECKSET,
}

func AlicloudFnFFlowBasicDependence361(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 全生命周期 1393
func TestAccAliCloudFnFFlow_basic1393(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_fnf_flow.default"
	ra := resourceAttrInit(resourceId, AlicloudFnFFlowMap1393)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &FnFServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeFnFFlow")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 100)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudFnFFlowBasicDependence1393)
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
					"type":        "FDL",
					"description": "test",
					"flow_name":   name,
					"definition":  "version: v1beta1\\ntype: flow\\nsteps:\\n  - type: pass\\n    name: helloworld",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"type":        "FDL",
						"description": "test",
						"flow_name":   name,
						"definition":  "version: v1beta1\ntype: flow\nsteps:\n  - type: pass\n    name: helloworld",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "testupdate",
					"definition":  "version: v1beta1\\ntype: flow\\nsteps:\\n  - type: pass\\n    name: helloworld11",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "testupdate",
						"definition":  "version: v1beta1\ntype: flow\nsteps:\n  - type: pass\n    name: helloworld11",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudFnFFlowMap1393 = map[string]string{
	"create_time":             CHECKSET,
	"tracing_configuration.#": CHECKSET,
	"flow_id":                 CHECKSET,
	"last_modified_time":      CHECKSET,
}

func AlicloudFnFFlowBasicDependence1393(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test FnF Flow. <<< Resource test cases, automatically generated.

func TestAccAlicloudFnfFlow_basic(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_fnf_flow.default"
	ra := resourceAttrInit(resourceId, AlicloudFnfFlowMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &FnfService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeFnfFlow")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testAcc%sAlicloudFnfFlow%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudFnfFlowBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.FnfSupportRegions)
		},

		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"definition":  `version: v1beta1\ntype: flow\nsteps:\n  - type: pass\n    name: helloworld`,
					"description": "tf-testaccFnFFlow983041",
					"name":        "${var.name}",
					"type":        "FDL",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"definition":  strings.Replace(`version: v1beta1\ntype: flow\nsteps:\n  - type: pass\n    name: helloworld`, `\n`, "\n", -1),
						"description": "tf-testaccFnFFlow983041",
						"name":        name,
						"type":        "FDL",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"definition": `version: v1beta1\ntype: flow\nsteps:\n  - type: pass\n    name: helloworldchange`,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"definition": strings.Replace(`version: v1beta1\ntype: flow\nsteps:\n  - type: pass\n    name: helloworldchange`, `\n`, "\n", -1),
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "tf-testaccFnFFlow813242",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "tf-testaccFnFFlow813242",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"role_arn": `${alicloud_ram_role.default.arn}`,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"role_arn": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"type": "FDL",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"type": "FDL",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"definition":  `version: v1beta1\ntype: flow\nsteps:\n  - type: pass\n    name: helloworld`,
					"description": "tf-testaccFnFFlow983041",
					"role_arn":    `${alicloud_ram_role.default.arn}`,
					"type":        "FDL",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"definition":  strings.Replace(`version: v1beta1\ntype: flow\nsteps:\n  - type: pass\n    name: helloworld`, `\n`, "\n", -1),
						"description": "tf-testaccFnFFlow983041",
						"role_arn":    CHECKSET,
						"type":        "FDL",
					}),
				),
			},
		},
	})
}

var AlicloudFnfFlowMap0 = map[string]string{
	"flow_id":            CHECKSET,
	"last_modified_time": CHECKSET,
}

func AlicloudFnfFlowBasicDependence0(name string) string {
	return fmt.Sprintf(`
variable "name" {
	default = "%s"
}
resource "alicloud_ram_role" "default" {
  name = var.name
  document = <<EOF
  {
    "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Principal": {
          "Service": [
            "fnf.aliyuncs.com"
          ]
        }
      }
    ],
    "Version": "1"
  }
  EOF
}
`, name)
}

// lintignore: R001
func TestUnitAlicloudFnfFlow(t *testing.T) {
	p := Provider().(*schema.Provider).ResourcesMap
	checkoutSupportedRegions(t, true, connectivity.FnFSupportRegions)
	dInit, _ := schema.InternalMap(p["alicloud_fnf_flow"].Schema).Data(nil, nil)
	dExisted, _ := schema.InternalMap(p["alicloud_fnf_flow"].Schema).Data(nil, nil)
	dInit.MarkNewResource()
	attributes := map[string]interface{}{
		"definition":  "CreateFlowValue",
		"description": "CreateFlowValue",
		"name":        "CreateFlowValue",
		"type":        "CreateFlowValue",
		"role_arn":    "CreateFlowValue",
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
		// DescribeFlow
		"Name":             "CreateFlowValue",
		"Definition":       "CreateFlowValue",
		"Description":      "CreateFlowValue",
		"Id":               "CreateFlowValue",
		"LastModifiedTime": "CreateFlowValue",
		"RoleArn":          "CreateFlowValue",
		"Type":             "CreateFlowValue",
	}
	CreateMockResponse := map[string]interface{}{
		// CreateFlow
		"Name": "CreateFlowValue",
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
		return nil, GetNotFoundErrorFromString(GetNotFoundMessage("alicloud_fnf_flow", errorCode))
	}
	successResponseMock := func(operationMockResponse map[string]interface{}) (map[string]interface{}, error) {
		if len(operationMockResponse) > 0 {
			mapMerge(ReadMockResponse, operationMockResponse)
		}
		return ReadMockResponse, nil
	}

	// Create
	patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewFnfClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudFnFFlowCreate(dInit, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	ReadMockResponseDiff := map[string]interface{}{
		// DescribeFlow Response
		"Name": "CreateFlowValue",
	}
	errorCodes := []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1 // a counter used to cover retry scenario; the same below
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "CreateFlow" {
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
		err := resourceAliCloudFnFFlowCreate(dInit, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_fnf_flow"].Schema).Data(dInit.State(), nil)
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
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewFnfClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudFnFFlowUpdate(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	// UpdateFlow
	attributesDiff := map[string]interface{}{
		"definition":  "UpdateFlowValue",
		"description": "UpdateFlowValue",
		"role_arn":    "UpdateFlowValue",
		"type":        "UpdateFlowValue",
	}
	diff, err := newInstanceDiff("alicloud_fnf_flow", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_fnf_flow"].Schema).Data(dInit.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeFlow Response
		"Definition":  "UpdateFlowValue",
		"Description": "UpdateFlowValue",
		"RoleArn":     "UpdateFlowValue",
		"Type":        "UpdateFlowValue",
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "ConcurrentUpdateError", "InternalServerError", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "UpdateFlow" {
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
		err := resourceAliCloudFnFFlowUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_fnf_flow"].Schema).Data(dExisted.State(), nil)
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
			if *action == "DescribeFlow" {
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
		err := resourceAliCloudFnFFlowRead(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "{}":
			assert.Nil(t, err)
		}
	}

	// Delete
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewFnfClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudFnFFlowDelete(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	errorCodes = []string{"NonRetryableError", "Throttling", "InternalServerError", "nil", "FlowNotExists"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DeleteFlow" {
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
		err := resourceAliCloudFnFFlowDelete(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "FlowNotExists":
			assert.Nil(t, err)
		}
	}

}
