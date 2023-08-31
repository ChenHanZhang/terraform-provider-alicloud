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
	"github.com/alibabacloud-go/tea/tea"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"

	util "github.com/alibabacloud-go/tea-utils/service"

	"github.com/alibabacloud-go/tea-rpc/client"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func init() {
	resource.AddTestSweepers("alicloud_cms_event_rule", &resource.Sweeper{
		Name: "alicloud_cms_event_rule",
		F:    testSweepCmsEventRules,
	})
}

func testSweepCmsEventRules(region string) error {
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return fmt.Errorf("error getting Alicloud client: %s", err)
	}
	client := rawClient.(*connectivity.AliyunClient)

	conn, err := client.NewCmsClient()
	if err != nil {
		return WrapError(err)
	}

	prefixes := []string{
		"tf-testAcc",
		"tf_testAcc",
	}
	action := "DescribeEventRuleList"
	request := make(map[string]interface{})
	request["PageSize"] = PageSizeLarge
	request["PageNumber"] = 1
	var response map[string]interface{}
	cmsEventRuleIds := make([]string, 0)
	for {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-01-01"), StringPointer("AK"), nil, request, &runtime)
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
			return WrapErrorf(err, DataDefaultErrorMsg, "alicloud_cms_event_rule", action, AlibabaCloudSdkGoERROR)
		}
		if fmt.Sprint(response["Success"]) == "false" {
			return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
		}
		resp, err := jsonpath.Get("$.EventRules.EventRule", response)
		if err != nil {
			return WrapErrorf(err, FailedGetAttributeMsg, action, "$.EventRules.EventRule", response)
		}
		result, _ := resp.([]interface{})
		for _, v := range result {
			skip := true
			item := v.(map[string]interface{})
			if !sweepAll() {
				for _, prefix := range prefixes {
					if strings.HasPrefix(strings.ToLower(fmt.Sprint(item["Name"])), strings.ToLower(prefix)) {
						skip = false
						break
					}
				}
				if skip {
					log.Printf("[INFO] Skipping CmsEventRule Instance: %v", item["Name"])
					continue
				}
			}
			cmsEventRuleIds = append(cmsEventRuleIds, fmt.Sprint(item["Name"]))
		}
		if len(result) < PageSizeLarge {
			break
		}
		request["PageNumber"] = request["PageNumber"].(int) + 1
	}

	log.Printf("[INFO] Deleting CmsEventRule Instances: %s", cmsEventRuleIds)
	deleteAction := "DeleteEventRules"
	if err != nil {
		return WrapError(err)
	}
	request = map[string]interface{}{
		"RuleNames": cmsEventRuleIds,
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		_, err = conn.DoRequest(StringPointer(deleteAction), nil, StringPointer("POST"), StringPointer("2019-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
		log.Printf("[ERROR] Failed to delete CmsEventRule Instance (%s): %s", cmsEventRuleIds, err)
	}

	return nil
}

func TestAccAlicloudCmsEventRule_basic00(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cms_event_rule.default"
	ra := resourceAttrInit(resourceId, resourceAlicloudCmsEventRuleMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CmsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCmsEventRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testAcc%sCmsEventRule-name%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceAlicloudCmsEventRuleBasicDependence)
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
					"rule_name":   name,
					"group_id":    "${alicloud_cms_monitor_group.default.id}",
					"description": "tf-testAcc",
					"status":      "ENABLED",
					"event_pattern": []map[string]interface{}{
						{
							"product":         "ecs",
							"event_type_list": []string{"StatusNotification"},
							"level_list":      []string{"CRITICAL"},
							"name_list":       []string{"test"},
							"sql_filter":      "test",
						},
					},
					"silence_time": "100",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"rule_name":       name,
						"group_id":        CHECKSET,
						"description":     "tf-testAcc",
						"status":          "ENABLED",
						"event_pattern.#": "1",
						"silence_time":    "100",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_id": "${alicloud_cms_monitor_group.new.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "update-tf-testAcc",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "update-tf-testAcc",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "DISABLED",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "DISABLED",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"event_pattern": []map[string]interface{}{
						{
							"product":         "ads",
							"event_type_list": []string{"Exception"},
							"level_list":      []string{"WARN"},
							"name_list":       []string{"update_test"},
							"sql_filter":      "update_test",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"event_pattern.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"silence_time": "1000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"silence_time": "1000",
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

var resourceAlicloudCmsEventRuleMap = map[string]string{}

func resourceAlicloudCmsEventRuleBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
	default = "%s"
}

resource "alicloud_cms_monitor_group" "default" {
	monitor_group_name = var.name
}

resource "alicloud_cms_monitor_group" "new" {
	monitor_group_name = "${var.name}_new"
}

`, name)
}

func TestUnitAlicloudCmsEventRule(t *testing.T) {
	p := Provider().(*schema.Provider).ResourcesMap
	dInit, _ := schema.InternalMap(p["alicloud_cms_event_rule"].Schema).Data(nil, nil)
	dExisted, _ := schema.InternalMap(p["alicloud_cms_event_rule"].Schema).Data(nil, nil)
	dInit.MarkNewResource()
	attributes := map[string]interface{}{
		"rule_name":   "CreateEventRuleValue",
		"group_id":    "CreateEventRuleValue",
		"description": "CreateEventRuleValue",
		"status":      "CreateEventRuleValue",
		"event_pattern": []map[string]interface{}{
			{
				"product":         "CreateEventRuleValue",
				"event_type_list": []string{"CreateEventRuleValue"},
				"level_list":      []string{"CreateEventRuleValue"},
				"name_list":       []string{"CreateEventRuleValue"},
				"sql_filter":      "CreateEventRuleValue",
			},
		},
		"silence_time": 10,
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
		// DescribeEventRuleList
		"EventRules": map[string]interface{}{
			"EventRule": []interface{}{
				map[string]interface{}{
					"Description": "CreateEventRuleValue",
					"GroupId":     "CreateEventRuleValue",
					"Name":        "CreateEventRuleValue",
					"State":       "CreateEventRuleValue",
					"EventPattern": map[string]interface{}{
						"Product": "CreateEventRuleValue",
						"EventTypeList": map[string]interface{}{
							"EventTypeList": []interface{}{"CreateEventRuleValue"},
						},
						"LevelList": map[string]interface{}{
							"LevelList": []interface{}{"CreateEventRuleValue"},
						},
						"NameList": map[string]interface{}{
							"NameList": []interface{}{"CreateEventRuleValue"},
						},
						"SQLFilter": "CreateEventRuleValue",
					},
					"SilenceTime": 10,
				},
			},
		},
	}
	CreateMockResponse := map[string]interface{}{
		"Code":    200,
		"Success": "true",
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
		return nil, GetNotFoundErrorFromString(GetNotFoundMessage("alicloud_cms_event_rule", errorCode))
	}
	successResponseMock := func(operationMockResponse map[string]interface{}) (map[string]interface{}, error) {
		if len(operationMockResponse) > 0 {
			mapMerge(ReadMockResponse, operationMockResponse)
		}
		return ReadMockResponse, nil
	}
	// Create
	patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewCmsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudCloudMonitorServiceEventRuleCreate(dInit, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	ReadMockResponseDiff := map[string]interface{}{}
	errorCodes := []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1 // a counter used to cover retry scenario; the same below
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "PutEventRule" {
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
		err := resourceAliCloudCloudMonitorServiceEventRuleCreate(dInit, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_cms_event_rule"].Schema).Data(dInit.State(), nil)
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
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewCmsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudCloudMonitorServiceEventRuleUpdate(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	attributesDiff := map[string]interface{}{
		"rule_name":   "CreateEventRuleValue",
		"group_id":    "PutEventRuleValue",
		"description": "PutEventRuleValue",
		"status":      "PutEventRuleValue",
		"event_pattern": []map[string]interface{}{
			{
				"product":         "PutEventRuleValue",
				"event_type_list": []string{"PutEventRuleValue"},
				"level_list":      []string{"PutEventRuleValue"},
				"name_list":       []string{"PutEventRuleValue"},
				"sql_filter":      "PutEventRuleValue",
			},
		},
		"silence_time": 20,
	}
	diff, err := newInstanceDiff("alicloud_cms_event_rule", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_cms_event_rule"].Schema).Data(dInit.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeEventRuleList Response
		"EventRules": map[string]interface{}{
			"EventRule": []interface{}{
				map[string]interface{}{
					"Description": "PutEventRuleValue",
					"GroupId":     "PutEventRuleValue",
					"Name":        "CreateEventRuleValue",
					"State":       "PutEventRuleValue",
					"EventPattern": map[string]interface{}{
						"Product": "PutEventRuleValue",
						"EventTypeList": map[string]interface{}{
							"EventTypeList": []interface{}{"PutEventRuleValue"},
						},
						"LevelList": map[string]interface{}{
							"LevelList": []interface{}{"PutEventRuleValue"},
						},
						"NameList": map[string]interface{}{
							"NameList": []interface{}{"PutEventRuleValue"},
						},
						"SQLFilter": "PutEventRuleValue",
					},
					"SilenceTime": 20,
				},
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "PutEventRule" {
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
		err := resourceAliCloudCloudMonitorServiceEventRuleUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_cms_event_rule"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// Read
	diff, err = newInstanceDiff("alicloud_cms_event_rule", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_cms_event_rule"].Schema).Data(dInit.State(), diff)
	errorCodes = []string{"NonRetryableError", "Throttling", "nil", "{}"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DescribeEventRuleList" {
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
		err := resourceAliCloudCloudMonitorServiceEventRuleRead(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "{}":
			assert.Nil(t, err)
		}
	}

	// Delete
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewCmsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudCloudMonitorServiceEventRuleDelete(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	attributesDiff = map[string]interface{}{}
	diff, err = newInstanceDiff("alicloud_cms_event_rule", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_cms_event_rule"].Schema).Data(dInit.State(), diff)
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DeleteEventRules" {
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
		err := resourceAliCloudCloudMonitorServiceEventRuleDelete(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "nil":
			assert.Nil(t, err)
		}
	}
}

// Test CloudMonitorService EventRule. >>> Resource test cases, automatically generated.
// Case 4285
func TestAccAliCloudCloudMonitorServiceEventRule_basic4285(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cms_event_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudMonitorServiceEventRuleMap4285)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudMonitorServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudMonitorServiceEventRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudmonitorserviceeventrule%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudMonitorServiceEventRuleBasicDependence4285)
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
					"event_rule_name": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"event_rule_name": name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "ENABLED",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "ENABLED",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "Description1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "Description1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_id": "${alicloud_cms_monitor_group.defaultMs3hsi.group_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"contact_parameters": []map[string]interface{}{
						{
							"contact_group_name":    "ContactGroupName111",
							"level":                 "3",
							"contact_parameters_id": "111",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"contact_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"open_api_parameters": []map[string]interface{}{
						{
							"action":                 "PutLogs1",
							"arn":                    "acs:openapi:cn-hangzhou:null:log1/PutLogs1/2019-01-01/MyRole1",
							"product":                "log1",
							"region":                 "cn-hangzhou",
							"role":                   "MyRole1",
							"version":                "2019-01-01",
							"open_api_parameters_id": "22",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"open_api_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sls_parameters": []map[string]interface{}{
						{
							"project":           "testproject3",
							"log_store":         "testlogstore3",
							"arn":               "acs:log:cn-hangzhou::project/testproject3/logstore/testlogstore3",
							"region":            "cn-hangzhou",
							"sls_parameters_id": "33",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"sls_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"webhook_parameters": []map[string]interface{}{
						{
							"url":                   "http://www.aliyun.com1",
							"method":                "POST",
							"protocol":              "http",
							"webhook_parameters_id": "44",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"webhook_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"fc_parameters": []map[string]interface{}{
						{
							"service_name":     "ServiceName15",
							"function_name":    "FunctionNamee5",
							"arn":              "acs:log:cn-hangzhou::services/ServiceName15/functions/FunctionNamee5",
							"region":           "cn-hangzhou",
							"fc_parameters_id": "55",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"fc_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"mns_parameters": []map[string]interface{}{
						{
							"queue":             "queue16",
							"arn":               "acs:mns:cn-hangzhou::/queues/queue16/messages",
							"region":            "cn-hangzhou",
							"mns_parameters_id": "66",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"mns_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "DISABLED",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "DISABLED",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "Description2312",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "Description2312",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_id": "${alicloud_cms_monitor_group.defaultGp1.group_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"contact_parameters": []map[string]interface{}{
						{
							"contact_group_name":    "云账号报警联系人1",
							"level":                 "2",
							"contact_parameters_id": "1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"contact_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"open_api_parameters": []map[string]interface{}{
						{
							"action":                 "PutLogs",
							"arn":                    "acs:openapi:cn-beijing:null:log/PutLogs/2018-03-08/MyRole",
							"product":                "log",
							"region":                 "cn-beijing",
							"role":                   "MyRole",
							"version":                "2018-03-08",
							"open_api_parameters_id": "2",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"open_api_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sls_parameters": []map[string]interface{}{
						{
							"project":           "testproject",
							"log_store":         "testlogstore",
							"arn":               "acs:log:cn-beijing::project/testproject/logstore/testlogstore",
							"region":            "cn-beijing",
							"sls_parameters_id": "3",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"sls_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"webhook_parameters": []map[string]interface{}{
						{
							"url":                   "http://www.aliyun.com",
							"method":                "get",
							"protocol":              "telnet",
							"webhook_parameters_id": "4",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"webhook_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"fc_parameters": []map[string]interface{}{
						{
							"service_name":     "ServiceName1",
							"function_name":    "FunctionNamee",
							"arn":              "acs:log:cn-beijing::services/ServiceName1/functions/FunctionNamee",
							"region":           "cn-beijing",
							"fc_parameters_id": "5",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"fc_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"mns_parameters": []map[string]interface{}{
						{
							"queue":             "queue1",
							"arn":               "acs:mns:cn-beijing::/queues/queue1/messages",
							"region":            "cn-beijing",
							"mns_parameters_id": "6",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"mns_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "ENABLED",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "ENABLED",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "Description1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "Description1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_id": "${alicloud_cms_monitor_group.defaultMs3hsi.group_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"contact_parameters": []map[string]interface{}{
						{
							"contact_group_name":    "ContactGroupName111",
							"level":                 "3",
							"contact_parameters_id": "111",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"contact_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"open_api_parameters": []map[string]interface{}{
						{
							"action":                 "PutLogs1",
							"arn":                    "acs:openapi:cn-hangzhou:null:log1/PutLogs1/2019-01-01/MyRole1",
							"product":                "log1",
							"region":                 "cn-hangzhou",
							"role":                   "MyRole1",
							"version":                "2019-01-01",
							"open_api_parameters_id": "22",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"open_api_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sls_parameters": []map[string]interface{}{
						{
							"project":           "testproject3",
							"log_store":         "testlogstore3",
							"arn":               "acs:log:cn-hangzhou::project/testproject3/logstore/testlogstore3",
							"region":            "cn-hangzhou",
							"sls_parameters_id": "33",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"sls_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"webhook_parameters": []map[string]interface{}{
						{
							"url":                   "http://www.aliyun.com1",
							"method":                "POST",
							"protocol":              "http",
							"webhook_parameters_id": "44",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"webhook_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"fc_parameters": []map[string]interface{}{
						{
							"service_name":     "ServiceName15",
							"function_name":    "FunctionNamee5",
							"arn":              "acs:log:cn-hangzhou::services/ServiceName15/functions/FunctionNamee5",
							"region":           "cn-hangzhou",
							"fc_parameters_id": "55",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"fc_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"mns_parameters": []map[string]interface{}{
						{
							"queue":             "queue16",
							"arn":               "acs:mns:cn-hangzhou::/queues/queue16/messages",
							"region":            "cn-hangzhou",
							"mns_parameters_id": "66",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"mns_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "DISABLED",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "DISABLED",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "Description2312",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "Description2312",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_id": "${alicloud_cms_monitor_group.defaultGp1.group_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"contact_parameters": []map[string]interface{}{
						{
							"contact_group_name":    "云账号报警联系人1",
							"level":                 "2",
							"contact_parameters_id": "1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"contact_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"open_api_parameters": []map[string]interface{}{
						{
							"action":                 "PutLogs",
							"arn":                    "acs:openapi:cn-beijing:null:log/PutLogs/2018-03-08/MyRole",
							"product":                "log",
							"region":                 "cn-beijing",
							"role":                   "MyRole",
							"version":                "2018-03-08",
							"open_api_parameters_id": "2",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"open_api_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sls_parameters": []map[string]interface{}{
						{
							"project":           "testproject",
							"log_store":         "testlogstore",
							"arn":               "acs:log:cn-beijing::project/testproject/logstore/testlogstore",
							"region":            "cn-beijing",
							"sls_parameters_id": "3",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"sls_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"webhook_parameters": []map[string]interface{}{
						{
							"url":                   "http://www.aliyun.com",
							"method":                "get",
							"protocol":              "telnet",
							"webhook_parameters_id": "4",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"webhook_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"fc_parameters": []map[string]interface{}{
						{
							"service_name":     "ServiceName1",
							"function_name":    "FunctionNamee",
							"arn":              "acs:log:cn-beijing::services/ServiceName1/functions/FunctionNamee",
							"region":           "cn-beijing",
							"fc_parameters_id": "5",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"fc_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"mns_parameters": []map[string]interface{}{
						{
							"queue":             "queue1",
							"arn":               "acs:mns:cn-beijing::/queues/queue1/messages",
							"region":            "cn-beijing",
							"mns_parameters_id": "6",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"mns_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"contact_parameters": []map[string]interface{}{
						{},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"contact_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"open_api_parameters": []map[string]interface{}{
						{},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"open_api_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sls_parameters": []map[string]interface{}{
						{},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"sls_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"webhook_parameters": []map[string]interface{}{
						{},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"webhook_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"fc_parameters": []map[string]interface{}{
						{},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"fc_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"mns_parameters": []map[string]interface{}{
						{},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"mns_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":       "ENABLED",
					"silence_time": "86400",
					"event_pattern": []map[string]interface{}{
						{
							"sql_filter": "11223ANDxx",
							"product":    "ECS",
							"level_list": []string{
								"CRITICAL"},
							"event_type_list": []string{
								"Exception"},
							"name_list": []string{
								"Disk:ErrorDetected:Executing"},
							"status_list": []string{
								"Normal"},
						},
					},
					"description":     "Description1",
					"event_rule_name": name + "_update",
					"group_id":        "${alicloud_cms_monitor_group.defaultMs3hsi.group_id}",
					"contact_parameters": []map[string]interface{}{
						{
							"contact_group_name":    "ContactGroupName111",
							"level":                 "3",
							"contact_parameters_id": "111",
						},
					},
					"open_api_parameters": []map[string]interface{}{
						{
							"action":                 "PutLogs1",
							"arn":                    "acs:openapi:cn-hangzhou:null:log1/PutLogs1/2019-01-01/MyRole1",
							"product":                "log1",
							"region":                 "cn-hangzhou",
							"role":                   "MyRole1",
							"version":                "2019-01-01",
							"open_api_parameters_id": "22",
						},
					},
					"sls_parameters": []map[string]interface{}{
						{
							"project":           "testproject3",
							"log_store":         "testlogstore3",
							"arn":               "acs:log:cn-hangzhou::project/testproject3/logstore/testlogstore3",
							"region":            "cn-hangzhou",
							"sls_parameters_id": "33",
						},
					},
					"webhook_parameters": []map[string]interface{}{
						{
							"url":                   "http://www.aliyun.com1",
							"method":                "POST",
							"protocol":              "http",
							"webhook_parameters_id": "44",
						},
					},
					"fc_parameters": []map[string]interface{}{
						{
							"service_name":     "ServiceName15",
							"function_name":    "FunctionNamee5",
							"arn":              "acs:log:cn-hangzhou::services/ServiceName15/functions/FunctionNamee5",
							"region":           "cn-hangzhou",
							"fc_parameters_id": "55",
						},
					},
					"mns_parameters": []map[string]interface{}{
						{
							"queue":             "queue16",
							"arn":               "acs:mns:cn-hangzhou::/queues/queue16/messages",
							"region":            "cn-hangzhou",
							"mns_parameters_id": "66",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":                "ENABLED",
						"silence_time":          "86400",
						"description":           "Description1",
						"event_rule_name":       name + "_update",
						"group_id":              CHECKSET,
						"contact_parameters.#":  "1",
						"open_api_parameters.#": "1",
						"sls_parameters.#":      "1",
						"webhook_parameters.#":  "1",
						"fc_parameters.#":       "1",
						"mns_parameters.#":      "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"silence_time"},
			},
		},
	})
}

var AlicloudCloudMonitorServiceEventRuleMap4285 = map[string]string{
	"status": CHECKSET,
}

func AlicloudCloudMonitorServiceEventRuleBasicDependence4285(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cms_monitor_group" "defaultMs3hsi" {
  monitor_group_name = var.name
}

resource "alicloud_cms_monitor_group" "defaultGp1" {
  monitor_group_name = "${var.name}1"
}


`, name)
}

// Case 4285  twin
func TestAccAliCloudCloudMonitorServiceEventRule_basic4285_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cms_event_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudMonitorServiceEventRuleMap4285)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudMonitorServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudMonitorServiceEventRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudmonitorserviceeventrule%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudMonitorServiceEventRuleBasicDependence4285)
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
					"status":       "DISABLED",
					"silence_time": "900",
					"event_pattern": []map[string]interface{}{
						{
							"sql_filter": "Warn and i-123xxx222",
							"product":    "CloudMonitor",
							"level_list": []string{
								"INFO", "CRITICAL"},
							"event_type_list": []string{
								"Exception", "Alert"},
							"name_list": []string{
								"Agent_Status_Running"},
							"status_list": []string{
								"running"},
						},
					},
					"description":     "Description2312",
					"event_rule_name": name,
					"group_id":        "${alicloud_cms_monitor_group.defaultGp1.id}",
					"contact_parameters": []map[string]interface{}{
						{
							"contact_group_name":    "云账号报警联系人1",
							"level":                 "2",
							"contact_parameters_id": "1",
						},
					},
					"open_api_parameters": []map[string]interface{}{
						{
							"action":                 "PutLogs",
							"arn":                    "acs:openapi:cn-beijing:null:log/PutLogs/2018-03-08/MyRole",
							"product":                "log",
							"region":                 "cn-beijing",
							"role":                   "MyRole",
							"version":                "2018-03-08",
							"open_api_parameters_id": "2",
						},
					},
					"sls_parameters": []map[string]interface{}{
						{
							"project":           "testproject",
							"log_store":         "testlogstore",
							"arn":               "acs:log:cn-beijing::project/testproject/logstore/testlogstore",
							"region":            "cn-beijing",
							"sls_parameters_id": "3",
						},
					},
					"webhook_parameters": []map[string]interface{}{
						{
							"url":                   "http://www.aliyun.com",
							"method":                "get",
							"protocol":              "telnet",
							"webhook_parameters_id": "4",
						},
					},
					"fc_parameters": []map[string]interface{}{
						{
							"service_name":     "ServiceName1",
							"function_name":    "FunctionNamee",
							"arn":              "acs:log:cn-beijing::services/ServiceName1/functions/FunctionNamee",
							"region":           "cn-beijing",
							"fc_parameters_id": "5",
						},
					},
					"mns_parameters": []map[string]interface{}{
						{
							"queue":             "queue1",
							"arn":               "acs:mns:cn-beijing::/queues/queue1/messages",
							"region":            "cn-beijing",
							"mns_parameters_id": "6",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":                "DISABLED",
						"silence_time":          "900",
						"description":           "Description2312",
						"event_rule_name":       name,
						"group_id":              CHECKSET,
						"contact_parameters.#":  "1",
						"open_api_parameters.#": "1",
						"sls_parameters.#":      "1",
						"webhook_parameters.#":  "1",
						"fc_parameters.#":       "1",
						"mns_parameters.#":      "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"silence_time"},
			},
		},
	})
}

// Test CloudMonitorService EventRule. <<< Resource test cases, automatically generated.
