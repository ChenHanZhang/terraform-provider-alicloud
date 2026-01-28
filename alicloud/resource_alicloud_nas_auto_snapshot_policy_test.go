package alicloud

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
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
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

func init() {
	resource.AddTestSweepers(
		"alicloud_nas_auto_snapshot_policy",
		&resource.Sweeper{
			Name: "alicloud_nas_auto_snapshot_policy",
			F:    testSweepNasAutoSnapshotPolicy,
		})
}

func testSweepNasAutoSnapshotPolicy(region string) error {
	if testSweepPreCheckWithRegions(region, true, connectivity.NASSupportRegions) {
		log.Printf("[INFO] Skipping Nas Auto Snapshot Policy unsupported region: %s", region)
		return nil
	}

	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return fmt.Errorf("error getting Alicloud client: %s", err)
	}
	client := rawClient.(*connectivity.AliyunClient)
	prefixes := []string{
		"tf-testAcc",
		"tf_testAcc",
	}

	action := "DescribeAutoSnapshotPolicies"
	request := map[string]interface{}{}
	request["FileSystemType"] = "extreme"
	request["PageSize"] = PageSizeLarge
	request["PageNumber"] = 1

	var response map[string]interface{}
	for {
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(1*time.Minute, func() *resource.RetryError {
			response, err = client.RpcPost("NAS", "2017-06-26", action, nil, request, true)
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
			log.Printf("[ERROR] %s get an error: %#v", action, err)
			return nil
		}

		resp, err := jsonpath.Get("$.AutoSnapshotPolicies.AutoSnapshotPolicy", response)
		if err != nil {
			log.Printf("[ERROR] Getting resource %s attribute by path %s failed!!! Body: %v.", "$.AutoSnapshotPolicies.AutoSnapshotPolicy", action, err)
			return nil
		}
		result, _ := resp.([]interface{})
		for _, v := range result {
			item := v.(map[string]interface{})

			skip := true
			for _, prefix := range prefixes {
				if strings.HasPrefix(strings.ToLower(item["AutoSnapshotPolicyName"].(string)), strings.ToLower(prefix)) {
					skip = false
				}
			}
			if skip {
				log.Printf("[INFO] Skipping Nas Auto Snapshot Policy: %s", item["AutoSnapshotPolicyName"].(string))
				continue
			}
			action := "DeleteAutoSnapshotPolicy"
			request := map[string]interface{}{
				"AutoSnapshotPolicyId": item["AutoSnapshotPolicyId"],
			}
			_, err = client.RpcPost("NAS", "2017-06-26", action, nil, request, false)
			if err != nil {
				log.Printf("[ERROR] Failed to delete Nas Auto Snapshot Policy (%s): %s", item["AutoSnapshotPolicyName"].(string), err)
			}
			log.Printf("[INFO] Delete Nas Auto Snapshot Policy success: %s ", item["AutoSnapshotPolicyName"].(string))
		}
		if len(result) < PageSizeLarge {
			break
		}
		request["PageNumber"] = request["PageNumber"].(int) + 1
	}
	return nil
}

func TestAccAliCloudNASAutoSnapshotPolicy_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_auto_snapshot_policy.default"
	checkoutSupportedRegions(t, true, connectivity.NASSupportRegions)
	ra := resourceAttrInit(resourceId, AlicloudNASAutoSnapshotPolicyMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasAutoSnapshotPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc-nasautosnapshotpolicy%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNASAutoSnapshotPolicyBasicDependence0)
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
					"repeat_weekdays": []string{"2", "3", "4"},
					"time_points":     []string{"0", "1", "2"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repeat_weekdays.#": "3",
						"time_points.#":     "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_snapshot_policy_name": "${var.name}_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_snapshot_policy_name": name + "_update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"repeat_weekdays": []string{"3", "4", "5"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"repeat_weekdays.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"time_points": []string{"1", "2", "3"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"time_points.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"retention_days": "30",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"retention_days": "30",
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

func TestAccAliCloudNASAutoSnapshotPolicy_basic1(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_auto_snapshot_policy.default"
	checkoutSupportedRegions(t, true, connectivity.NASSupportRegions)
	ra := resourceAttrInit(resourceId, AlicloudNASAutoSnapshotPolicyMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasAutoSnapshotPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc-nasautosnapshotpolicy%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNASAutoSnapshotPolicyBasicDependence0)
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
					"auto_snapshot_policy_name": "${var.name}",
					"repeat_weekdays":           []string{"3", "4"},
					"time_points":               []string{"1", "2"},
					"retention_days":            "30",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_snapshot_policy_name": name,
						"repeat_weekdays.#":         "2",
						"time_points.#":             "2",
						"retention_days":            "30",
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

var AlicloudNASAutoSnapshotPolicyMap0 = map[string]string{
	"status": CHECKSET,
}

func AlicloudNASAutoSnapshotPolicyBasicDependence0(name string) string {
	return fmt.Sprintf(` 
variable "name" {
  default = "%s"
}
`, name)
}

func TestUnitAlicloudNASAutoSnapshotPolicy(t *testing.T) {
	p := Provider().(*schema.Provider).ResourcesMap
	d, _ := schema.InternalMap(p["alicloud_nas_auto_snapshot_policy"].Schema).Data(nil, nil)
	dCreate, _ := schema.InternalMap(p["alicloud_nas_auto_snapshot_policy"].Schema).Data(nil, nil)
	dCreate.MarkNewResource()
	for key, value := range map[string]interface{}{
		"auto_snapshot_policy_name": "auto_snapshot_policy_name",
		"repeat_weekdays":           []string{"3", "4"},
		"time_points":               []string{"1", "2"},
		"retention_days":            30,
	} {
		err := dCreate.Set(key, value)
		assert.Nil(t, err)
		err = d.Set(key, value)
		assert.Nil(t, err)
	}
	region := os.Getenv("ALICLOUD_REGION")
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		t.Skipf("Skipping the test case with err: %s", err)
		t.Skipped()
	}
	rawClient = rawClient.(*connectivity.AliyunClient)
	ReadMockResponse := map[string]interface{}{
		"AutoSnapshotPolicies": map[string]interface{}{
			"AutoSnapshotPolicy": []interface{}{
				map[string]interface{}{
					"AutoSnapshotPolicyName": "auto_snapshot_policy_name",
					"RepeatWeekdays":         "3,4",
					"RetentionDays":          30,
					"Status":                 "Available",
					"AutoSnapshotPolicyId":   "MockAutoSnapshotPolicyId",
					"TimePoints":             "1,2",
				},
			},
		},
	}

	responseMock := map[string]func(errorCode string) (map[string]interface{}, error){
		"RetryError": func(errorCode string) (map[string]interface{}, error) {
			return nil, &tea.SDKError{
				Code:       String(errorCode),
				Data:       String(errorCode),
				Message:    String(errorCode),
				StatusCode: tea.Int(400),
			}
		},
		"NotFoundError": func(errorCode string) (map[string]interface{}, error) {
			return nil, GetNotFoundErrorFromString(GetNotFoundMessage("alicloud_nas_auto_snapshot_policy", "MockAutoSnapshotPolicyId"))
		},
		"NoRetryError": func(errorCode string) (map[string]interface{}, error) {
			return nil, &tea.SDKError{
				Code:       String(errorCode),
				Data:       String(errorCode),
				Message:    String(errorCode),
				StatusCode: tea.Int(400),
			}
		},
		"CreateNormal": func(errorCode string) (map[string]interface{}, error) {
			result := ReadMockResponse
			result["AutoSnapshotPolicyId"] = "MockAutoSnapshotPolicyId"
			return result, nil
		},
		"UpdateNormal": func(errorCode string) (map[string]interface{}, error) {
			result := ReadMockResponse
			return result, nil
		},
		"DeleteNormal": func(errorCode string) (map[string]interface{}, error) {
			result := ReadMockResponse
			return result, nil
		},
		"ReadNormal": func(errorCode string) (map[string]interface{}, error) {
			result := ReadMockResponse
			return result, nil
		},
	}
	// Create
	t.Run("CreateClientAbnormal", func(t *testing.T) {
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewNasClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
			return nil, &tea.SDKError{
				Code:       String("loadEndpoint error"),
				Data:       String("loadEndpoint error"),
				Message:    String("loadEndpoint error"),
				StatusCode: tea.Int(400),
			}
		})
		err := resourceAliCloudNasAutoSnapshotPolicyCreate(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})
	t.Run("CreateAbnormal", func(t *testing.T) {
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["CreateNormal"]("")
		})
		err := resourceAliCloudNasAutoSnapshotPolicyCreate(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})
	t.Run("CreateNormal", func(t *testing.T) {
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["CreateNormal"]("")
		})
		err := resourceAliCloudNasAutoSnapshotPolicyCreate(dCreate, rawClient)
		patches.Reset()
		assert.Nil(t, err)
	})

	// Set ID for Update and Delete Method
	d.SetId("MockAutoSnapshotPolicyId")
	// Update
	t.Run("UpdateClientAbnormal", func(t *testing.T) {
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewNasClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
			return nil, &tea.SDKError{
				Code:       String("loadEndpoint error"),
				Data:       String("loadEndpoint error"),
				Message:    String("loadEndpoint error"),
				StatusCode: tea.Int(400),
			}
		})

		err := resourceAliCloudNasAutoSnapshotPolicyUpdate(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateModifyAutoSnapshotPolicyAbnormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"auto_snapshot_policy_name", "repeat_weekdays", "retention_days", "time_points"} {
			switch p["alicloud_nas_auto_snapshot_policy"].Schema[key].Type {
			case schema.TypeString:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: d.Get(key).(string), New: d.Get(key).(string) + "_update"})
			case schema.TypeBool:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: strconv.FormatBool(d.Get(key).(bool)), New: strconv.FormatBool(true)})
			case schema.TypeInt:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: strconv.Itoa(d.Get(key).(int)), New: strconv.Itoa(3)})
			case schema.TypeMap:
				diff.SetAttribute("tags.%", &terraform.ResourceAttrDiff{Old: "0", New: "2"})
				diff.SetAttribute("tags.For", &terraform.ResourceAttrDiff{Old: "", New: "Test"})
				diff.SetAttribute("tags.Created", &terraform.ResourceAttrDiff{Old: "", New: "TF"})
			case schema.TypeSet:
				diff.SetAttribute("repeat_weekdays.0", &terraform.ResourceAttrDiff{Old: "", New: "4"})
				diff.SetAttribute("time_points.0", &terraform.ResourceAttrDiff{Old: "", New: "3"})
			}

		}
		resourceData1, _ := schema.InternalMap(p["alicloud_nas_auto_snapshot_policy"].Schema).Data(nil, diff)
		resourceData1.SetId(d.Id())
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["UpdateNormal"]("")
		})
		err := resourceAliCloudNasAutoSnapshotPolicyUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateModifyAutoSnapshotPolicyNormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"auto_snapshot_policy_name", "repeat_weekdays", "retention_days", "time_points"} {
			switch p["alicloud_nas_auto_snapshot_policy"].Schema[key].Type {
			case schema.TypeString:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: d.Get(key).(string), New: d.Get(key).(string) + "_update"})
			case schema.TypeBool:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: strconv.FormatBool(d.Get(key).(bool)), New: strconv.FormatBool(true)})
			case schema.TypeInt:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: strconv.Itoa(d.Get(key).(int)), New: strconv.Itoa(3)})
			case schema.TypeMap:
				diff.SetAttribute("tags.%", &terraform.ResourceAttrDiff{Old: "0", New: "2"})
				diff.SetAttribute("tags.For", &terraform.ResourceAttrDiff{Old: "", New: "Test"})
				diff.SetAttribute("tags.Created", &terraform.ResourceAttrDiff{Old: "", New: "TF"})
			case schema.TypeSet:
				diff.SetAttribute("repeat_weekdays.0", &terraform.ResourceAttrDiff{Old: "", New: "4"})
				diff.SetAttribute("time_points.0", &terraform.ResourceAttrDiff{Old: "", New: "3"})
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_nas_auto_snapshot_policy"].Schema).Data(nil, diff)
		resourceData1.SetId(d.Id())
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["UpdateNormal"]("")
		})
		err := resourceAliCloudNasAutoSnapshotPolicyUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.Nil(t, err)
	})

	// Delete
	t.Run("DeleteClientAbnormal", func(t *testing.T) {
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewNasClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
			return nil, &tea.SDKError{
				Code:       String("loadEndpoint error"),
				Data:       String("loadEndpoint error"),
				Message:    String("loadEndpoint error"),
				StatusCode: tea.Int(400),
			}
		})
		err := resourceAliCloudNasAutoSnapshotPolicyDelete(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})
	t.Run("DeleteMockAbnormal", func(t *testing.T) {
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["DeleteNormal"]("")
		})
		err := resourceAliCloudNasAutoSnapshotPolicyDelete(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})
	t.Run("DeleteMockNormal", func(t *testing.T) {
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["DeleteNormal"]("")
		})
		patcheDescribeNasAutoSnapshotPolicy := gomonkey.ApplyMethod(reflect.TypeOf(&NasService{}), "DescribeNasAutoSnapshotPolicy", func(*NasService, string) (map[string]interface{}, error) {
			return responseMock["NotFoundError"]("ResourceNotfound")
		})
		err := resourceAliCloudNasAutoSnapshotPolicyDelete(d, rawClient)
		patches.Reset()
		patcheDescribeNasAutoSnapshotPolicy.Reset()
		assert.Nil(t, err)
	})

	t.Run("DeleteIsExpectedErrors", func(t *testing.T) {
		retryFlag := true
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				return responseMock["RetryError"]("InvalidLifecyclePolicy.NotFound")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["DeleteNormal"]("")
		})
		err := resourceAliCloudNasAutoSnapshotPolicyDelete(d, rawClient)
		patches.Reset()
		assert.Nil(t, err)
	})

	//Read
	t.Run("ReadDescribeNasAutoSnapshotPolicyNotFound", func(t *testing.T) {
		patcheDorequest := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			NotFoundFlag := true
			noRetryFlag := false
			if NotFoundFlag {
				return responseMock["NotFoundError"]("ResourceNotfound")
			} else if noRetryFlag {
				return responseMock["NoRetryError"]("NoRetryError")
			}
			return responseMock["ReadNormal"]("")
		})
		err := resourceAliCloudNasAutoSnapshotPolicyRead(d, rawClient)
		patcheDorequest.Reset()
		assert.Nil(t, err)
	})

	t.Run("ReadDescribeNasAutoSnapshotPolicyAbnormal", func(t *testing.T) {
		patcheDorequest := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			retryFlag := false
			noRetryFlag := true
			if retryFlag {
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["ReadNormal"]("")
		})
		err := resourceAliCloudNasAutoSnapshotPolicyRead(d, rawClient)
		patcheDorequest.Reset()
		assert.NotNil(t, err)
	})
}

var AlicloudNasAutoSnapshotPolicyMap6532 = map[string]string{
	"status":         CHECKSET,
	"create_time":    CHECKSET,
	"retention_days": CHECKSET,
}

func AlicloudNasAutoSnapshotPolicyBasicDependence6532(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 自动快照策略 6532  raw
func TestAccAliCloudNasAutoSnapshotPolicy_basic6532_raw(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_auto_snapshot_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudNasAutoSnapshotPolicyMap6532)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasAutoSnapshotPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%snasautosnapshotpolicy%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasAutoSnapshotPolicyBasicDependence6532)
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
					"time_points":               []string{"0", "1", "2"},
					"retention_days":            "1",
					"repeat_weekdays":           []string{"2", "3", "4"},
					"auto_snapshot_policy_name": name,
					"file_system_type":          "extreme",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"time_points.#":             "3",
						"retention_days":            "1",
						"repeat_weekdays.#":         "3",
						"auto_snapshot_policy_name": name,
						"file_system_type":          "extreme",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"time_points":               []string{"0", "5"},
					"retention_days":            "2",
					"repeat_weekdays":           []string{"4", "5", "6", "7"},
					"auto_snapshot_policy_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"time_points.#":             "2",
						"retention_days":            "2",
						"repeat_weekdays.#":         "4",
						"auto_snapshot_policy_name": name + "_update",
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

// Test Nas AutoSnapshotPolicy. <<< Resource test cases, automatically generated.
// Test Nas AutoSnapshotPolicy. >>> Resource test cases, automatically generated.
// Case resource_AutoSnapshotPolicy_test 12376
func TestAccAliCloudNasAutoSnapshotPolicy_basic12376(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_auto_snapshot_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudNasAutoSnapshotPolicyMap12376)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasAutoSnapshotPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnas%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasAutoSnapshotPolicyBasicDependence12376)
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
					"time_points": []string{
						"0", "8", "16"},
					"file_system_type": "extreme",
					"repeat_weekdays": []string{
						"1", "2", "3"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"time_points.#":     "3",
						"file_system_type":  "extreme",
						"repeat_weekdays.#": "3",
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
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudNasAutoSnapshotPolicyMap12376 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudNasAutoSnapshotPolicyBasicDependence12376(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case resource_AutoSnapshotPolicy_test_1 12385
func TestAccAliCloudNasAutoSnapshotPolicy_basic12385(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_auto_snapshot_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudNasAutoSnapshotPolicyMap12385)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasAutoSnapshotPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnas%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasAutoSnapshotPolicyBasicDependence12385)
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
					"time_points": []string{
						"0", "6", "12"},
					"file_system_type": "extreme",
					"retention_days":   "30",
					"repeat_weekdays": []string{
						"1", "3", "5"},
					"auto_snapshot_policy_name": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"time_points.#":             "3",
						"file_system_type":          "extreme",
						"retention_days":            "30",
						"repeat_weekdays.#":         "3",
						"auto_snapshot_policy_name": name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"retention_days":            "7",
					"auto_snapshot_policy_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"retention_days":            "7",
						"auto_snapshot_policy_name": name + "_update",
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

var AlicloudNasAutoSnapshotPolicyMap12385 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudNasAutoSnapshotPolicyBasicDependence12385(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 自动快照策略_副本 6770
func TestAccAliCloudNasAutoSnapshotPolicy_basic6770(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_auto_snapshot_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudNasAutoSnapshotPolicyMap6770)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasAutoSnapshotPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnas%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasAutoSnapshotPolicyBasicDependence6770)
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
					"time_points":               []string{},
					"retention_days":            "1",
					"repeat_weekdays":           []string{},
					"auto_snapshot_policy_name": name,
					"file_system_type":          "extreme",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"time_points.#":             "1",
						"retention_days":            CHECKSET,
						"repeat_weekdays.#":         "1",
						"auto_snapshot_policy_name": name,
						"file_system_type":          "extreme",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"time_points":               []string{},
					"retention_days":            "2",
					"repeat_weekdays":           []string{},
					"auto_snapshot_policy_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"time_points.#":             "1",
						"retention_days":            CHECKSET,
						"repeat_weekdays.#":         "1",
						"auto_snapshot_policy_name": name + "_update",
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

var AlicloudNasAutoSnapshotPolicyMap6770 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudNasAutoSnapshotPolicyBasicDependence6770(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 自动快照策略 6532
func TestAccAliCloudNasAutoSnapshotPolicy_basic6532(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_auto_snapshot_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudNasAutoSnapshotPolicyMap6532)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasAutoSnapshotPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnas%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasAutoSnapshotPolicyBasicDependence6532)
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
					"retention_days":            "1",
					"auto_snapshot_policy_name": name,
					"file_system_type":          "extreme",
					"time_points": []string{
						"${var.time_point1}"},
					"repeat_weekdays": []string{
						"${var.repeat_weekday1}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"retention_days":            CHECKSET,
						"auto_snapshot_policy_name": name,
						"file_system_type":          "extreme",
						"time_points.#":             "1",
						"repeat_weekdays.#":         "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"retention_days":            "2",
					"auto_snapshot_policy_name": name + "_update",
					"time_points": []string{
						"${var.time_point1}", "${var.time_point2}", "${var.time_point3}"},
					"repeat_weekdays": []string{
						"${var.repeat_weekday1}", "${var.repeat_weekday2}", "${var.repeat_weekday3}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"retention_days":            CHECKSET,
						"auto_snapshot_policy_name": name + "_update",
						"time_points.#":             "3",
						"repeat_weekdays.#":         "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"time_points": []string{
						"${var.time_point3}"},
					"repeat_weekdays": []string{
						"${var.repeat_weekday2}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"time_points.#":     "1",
						"repeat_weekdays.#": "1",
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

var AlicloudNasAutoSnapshotPolicyMap6532 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudNasAutoSnapshotPolicyBasicDependence6532(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "time_point3" {
  default = <<EOF
3
EOF
}

variable "repeat_weekday2" {
  default = <<EOF
2
EOF
}

variable "repeat_weekday1" {
  default = <<EOF
1
EOF
}

variable "repeat_weekday3" {
  default = <<EOF
3
EOF
}

variable "time_point2" {
  default = <<EOF
2
EOF
}

variable "time_point1" {
  default = <<EOF
1
EOF
}


`, name)
}

// Test Nas AutoSnapshotPolicy. <<< Resource test cases, automatically generated.
