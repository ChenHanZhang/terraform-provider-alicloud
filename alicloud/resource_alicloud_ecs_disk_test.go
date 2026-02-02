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
	resource.AddTestSweepers("alicloud_ecs_disk", &resource.Sweeper{
		Name: "alicloud_ecs_disk",
		F:    testAliCloudEcsDisk,
	})
}

func testAliCloudEcsDisk(region string) error {
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return WrapErrorf(err, "Error getting AliCloud client.")
	}
	client := rawClient.(*connectivity.AliyunClient)
	prefixes := []string{
		"tf-testAcc",
		"tf-testacc",
	}
	request := map[string]interface{}{
		"PageSize":   PageSizeLarge,
		"PageNumber": 1,
		"RegionId":   client.RegionId,
	}

	var response map[string]interface{}
	for {
		action := "DescribeDisks"
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
		if err != nil {
			log.Printf("[ERROR] %s got an error: %s", action, err)
			return nil
		}
		addDebug(action, response, request)

		resp, err := jsonpath.Get("$.Disks.Disk", response)
		if err != nil {
			return WrapErrorf(err, FailedGetAttributeMsg, action, "$.Disks.Disk", response)
		}

		result, _ := resp.([]interface{})
		for _, v := range result {
			item := v.(map[string]interface{})
			skip := true
			for _, prefix := range prefixes {
				if strings.HasPrefix(strings.ToLower(item["DiskName"].(string)), strings.ToLower(prefix)) {
					skip = false
				}
			}
			if skip {
				log.Printf("[INFO] Skipping Disk: %s", item["DiskName"].(string))
				continue
			}
			action = "DeleteDisk"
			request := map[string]interface{}{
				"DiskId":   item["DiskId"],
				"RegionId": client.RegionId,
			}
			_, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, false)
			if err != nil {
				log.Printf("[ERROR] Failed to delete Disk (%s): %s", item["DiskName"].(string), err)
			}
			log.Printf("[INFO] Delete Disk success: %s ", item["DiskName"].(string))
		}
		if len(result) < PageSizeLarge {
			break
		}
		request["PageNumber"] = request["PageNumber"].(int) + 1
	}
	return nil
}

func TestUnitECSDisk(t *testing.T) {
	p := Provider().(*schema.Provider).ResourcesMap
	dInit, _ := schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(nil, nil)
	dExisted, _ := schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(nil, nil)
	dInit.MarkNewResource()
	attributes := map[string]interface{}{
		"advanced_features": "CreateDiskValue",
		"disk_name":         "CreateDiskValue",
		"description":       "CreateDiskValue",
		"zone_id":           "CreateDiskValue",
		"size":              500,
		"payment_type":      "PayAsYouGo",
		"category":          "CreateDiskValue",
		"resource_group_id": "CreateDiskValue",
		"snapshot_id":       "CreateDiskValue",
		"tags": map[string]string{
			"TagResourcesValue_1": "CreateDiskValue",
			"TagResourcesValue_2": "CreateDiskValue",
		},
		"encrypt_algorithm":                  "CreateDiskValue",
		"instance_id":                        "CreateDiskValue",
		"kms_key_id":                         "CreateDiskValue",
		"performance_level":                  "CreateDiskValue",
		"storage_set_id":                     "CreateDiskValue",
		"storage_set_partition_number":       1,
		"dedicated_block_storage_cluster_id": "CreateDiskValue",
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
		// DescribeDisks
		"Disks": map[string]interface{}{
			"Disk": []interface{}{
				map[string]interface{}{
					"DiskId":             "CreateDiskValue",
					"Category":           "CreateDiskValue",
					"DeleteAutoSnapshot": true,
					"DeleteWithInstance": true,
					"Description":        "CreateDiskValue",
					"DiskName":           "CreateDiskValue",
					"EnableAutoSnapshot": false,
					"Encrypted":          "CreateDiskValue",
					"InstanceId":         "CreateDiskValue",
					"KMSKeyId":           "CreateDiskValue",
					"DiskChargeType":     "PostPaid",
					"PerformanceLevel":   "CreateDiskValue",
					"ResourceGroupId":    "CreateDiskValue",
					"Size":               500,
					"SourceSnapshotId":   "CreateDiskValue",
					"Status":             "Available",
					"ZoneId":             "CreateDiskValue",
					"Tags": map[string]interface{}{
						"Tag": []interface{}{
							map[string]interface{}{
								"Key":   "TagResourcesValue_1",
								"Value": "CreateDiskValue",
							},
							map[string]interface{}{
								"Key":   "TagResourcesValue_2",
								"Value": "CreateDiskValue",
							},
						},
					},
				},
			},
		},
		"DiskId": "CreateDiskValue",
	}
	CreateMockResponse := map[string]interface{}{
		// CreateDisk
		"DiskId": "CreateDiskValue",
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
		return nil, GetNotFoundErrorFromString(GetNotFoundMessage("alicloud_ecs_disk", errorCode))
	}
	successResponseMock := func(operationMockResponse map[string]interface{}) (map[string]interface{}, error) {
		if len(operationMockResponse) > 0 {
			mapMerge(ReadMockResponse, operationMockResponse)
		}
		return ReadMockResponse, nil
	}

	// Create
	patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewEcsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudEcsDiskCreate(dInit, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	ReadMockResponseDiff := map[string]interface{}{
		// DescribeDisks Response
		"DiskId": "CreateDiskValue",
	}
	errorCodes := []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1 // a counter used to cover retry scenario; the same below
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "CreateDisk" {
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
		err := resourceAliCloudEcsDiskCreate(dInit, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dInit.State(), nil)
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
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewEcsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudEcsDiskUpdate(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	// ResizeDisk
	attributesDiff := map[string]interface{}{
		"size": 600,
		"type": "ResizeDiskValue",
	}
	diff, err := newInstanceDiff("alicloud_ecs_disk", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dInit.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeDisks Response
		"Disks": map[string]interface{}{
			"Disk": []interface{}{
				map[string]interface{}{
					"Size": 600,
				},
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "ResizeDisk" {
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
		err := resourceAliCloudEcsDiskUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}
	// JoinResourceGroup
	attributesDiff = map[string]interface{}{
		"resource_group_id": "JoinResourceGroupValue",
	}
	diff, err = newInstanceDiff("alicloud_ecs_disk", attributes, attributesDiff, dExisted.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dExisted.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeDisks Response
		"Disks": map[string]interface{}{
			"Disk": []interface{}{
				map[string]interface{}{
					"ResourceGroupId": "JoinResourceGroupValue",
				},
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "JoinResourceGroup" {
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
		err := resourceAliCloudEcsDiskUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}
	// ModifyDiskSpec
	attributesDiff = map[string]interface{}{
		"category":          "ModifyDiskSpecValue",
		"performance_level": "ModifyDiskSpecValue",
		"dry_run":           false,
	}
	diff, err = newInstanceDiff("alicloud_ecs_disk", attributes, attributesDiff, dExisted.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dExisted.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeDisks Response
		"Disks": map[string]interface{}{
			"Disk": []interface{}{
				map[string]interface{}{
					"Category":         "ModifyDiskSpecValue",
					"PerformanceLevel": "ModifyDiskSpecValue",
				},
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "ModifyDiskSpec" {
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
		err := resourceAliCloudEcsDiskUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}
	// ModifyDiskChargeType
	attributesDiff = map[string]interface{}{
		"instance_id":  "ModifyDiskChargeTypeValue",
		"payment_type": "Subscription",
	}
	diff, err = newInstanceDiff("alicloud_ecs_disk", attributes, attributesDiff, dExisted.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dExisted.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeDisks Response
		"Disks": map[string]interface{}{
			"Disk": []interface{}{
				map[string]interface{}{
					"InstanceId":     "ModifyDiskChargeTypeValue",
					"DiskChargeType": "PrePaid",
				},
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "ModifyDiskChargeType" {
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
		err := resourceAliCloudEcsDiskUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}
	// ModifyDiskAttribute
	attributesDiff = map[string]interface{}{
		"delete_auto_snapshot": false,
		"delete_with_instance": false,
		"disk_name":            "ModifyDiskAttributeValue",
		"name":                 "ModifyDiskAttributeValue",
		"description":          "ModifyDiskAttributeValue",
		"enable_auto_snapshot": true,
	}
	diff, err = newInstanceDiff("alicloud_ecs_disk", attributes, attributesDiff, dExisted.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dExisted.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeDisks Response
		"Disks": map[string]interface{}{
			"Disk": []interface{}{
				map[string]interface{}{
					"DeleteAutoSnapshot": false,
					"DeleteWithInstance": false,
					"Description":        "ModifyDiskAttributeValue",
					"DiskName":           "ModifyDiskAttributeValue",
					"EnableAutoSnapshot": true,
				},
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "ModifyDiskAttribute" {
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
		err := resourceAliCloudEcsDiskUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}
	// TagResources
	attributesDiff = map[string]interface{}{
		"tags": map[string]interface{}{
			"TagResourcesValue_1": "TagResourcesValue_1",
			"TagResourcesValue_2": "TagResourcesValue_2",
		},
	}
	diff, err = newInstanceDiff("alicloud_ecs_disk", attributes, attributesDiff, dExisted.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dExisted.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeDisks Response
		"Disks": map[string]interface{}{
			"Disk": []interface{}{
				map[string]interface{}{
					"Tags": map[string]interface{}{
						"Tag": []interface{}{
							map[string]interface{}{
								"Key":   "TagResourcesValue_1",
								"Value": "TagResourcesValue_1",
							},
							map[string]interface{}{
								"Key":   "TagResourcesValue_2",
								"Value": "TagResourcesValue_2",
							},
						},
					},
				},
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "TagResources" {
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
		err := resourceAliCloudEcsDiskUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// UntagResources
	attributesDiff = map[string]interface{}{
		"tags": map[string]interface{}{
			"UntagResourcesValue3_1": "UnTagResourcesValue3_1",
			"UntagResourcesValue3_2": "UnTagResourcesValue3_2",
		},
	}
	diff, err = newInstanceDiff("alicloud_ecs_disk", attributes, attributesDiff, dExisted.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dExisted.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeDisks Response
		"Disks": map[string]interface{}{
			"Disk": []interface{}{
				map[string]interface{}{
					"Tags": map[string]interface{}{
						"Tag": []interface{}{
							map[string]interface{}{
								"Key":   "UntagResourcesValue3_1",
								"Value": "UnTagResourcesValue3_1",
							},
							map[string]interface{}{
								"Key":   "UntagResourcesValue3_2",
								"Value": "UnTagResourcesValue3_2",
							},
						},
					},
				},
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "UntagResources" {
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
		err := resourceAliCloudEcsDiskUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dExisted.State(), nil)
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
	attributesDiff = map[string]interface{}{}
	diff, err = newInstanceDiff("alicloud_ecs_disk", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dInit.State(), diff)
	errorCodes = []string{"NonRetryableError", "Throttling", "nil", "{}"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DescribeDisks" {
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
		err := resourceAliCloudEcsDiskRead(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "{}":
			assert.Nil(t, err)
		}
	}

	// Delete
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewEcsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudEcsDiskDelete(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	attributesDiff = map[string]interface{}{}
	diff, err = newInstanceDiff("alicloud_ecs_disk", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_ecs_disk"].Schema).Data(dInit.State(), diff)
	errorCodes = []string{"NonRetryableError", "Throttling", "IncorrectDiskStatus.Initializing", "nil", "InvalidDiskId.NotFound"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DeleteDisk" {
				switch errorCode {
				case "NonRetryableError", "InvalidDiskId.NotFound":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if errorCodes[retryIndex] == "nil" {
						ReadMockResponse = map[string]interface{}{
							"Disks": map[string]interface{}{
								"Disk": []interface{}{
									map[string]interface{}{
										"DiskId": "CreateDiskValue",
									},
								},
							},
						}
						return ReadMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudEcsDiskDelete(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "InvalidDiskId.NotFound":
			assert.Nil(t, err)
		}
	}
}

// Test Ecs Disk. >>> Resource test cases, automatically generated.
// Case 简单生命周期 8017
func TestAccAliCloudEcsDisk_basic8017(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsDiskMap8017)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 100)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsDiskBasicDependence8017)
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
					"category":                     "cloud_essd",
					"description":                  "加密测试",
					"zone_id":                      "${var.zone_id}",
					"size":                         "461",
					"delete_auto_snapshot":         "true",
					"delete_with_instance":         "true",
					"enable_auto_snapshot":         "false",
					"disk_name":                    name,
					"multi_attach":                 "Disabled",
					"performance_level":            "PL2",
					"kms_key_id":                   "${alicloud_kms_key.key.id}",
					"encrypted":                    "true",
					"storage_set_partition_number": "2",
					"storage_set_id":               "${alicloud_ecs_storage_set.StorageSet.id}",
					"advanced_features":            "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"category":                     "cloud_essd",
						"description":                  "加密测试",
						"zone_id":                      CHECKSET,
						"size":                         "461",
						"delete_auto_snapshot":         "true",
						"delete_with_instance":         "true",
						"enable_auto_snapshot":         "false",
						"disk_name":                    name,
						"multi_attach":                 "Disabled",
						"performance_level":            "PL2",
						"kms_key_id":                   CHECKSET,
						"encrypted":                    "true",
						"storage_set_partition_number": "2",
						"storage_set_id":               CHECKSET,
						"advanced_features":            CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description":       "zoneIdUpdate",
					"disk_name":         name + "_update",
					"performance_level": "PL1",
					"payment_type":      "PayAsYouGo",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":       "zoneIdUpdate",
						"disk_name":         name + "_update",
						"performance_level": "PL1",
						"payment_type":      "PayAsYouGo",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"category":  "cloud_auto",
					"disk_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"category":  "cloud_auto",
						"disk_name": name + "_update",
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
				ImportStateVerifyIgnore: []string{"advanced_features", "dry_run", "encrypt_algorithm", "image_id", "type"},
			},
		},
	})
}

var AlicloudEcsDiskMap8017 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudEcsDiskBasicDependence8017(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "zone_id" {
  default = "cn-hangzhou-i"
}

resource "alicloud_ecs_storage_set" "StorageSet" {
  zone_id              = var.zone_id
  max_partition_number = "2"
}

resource "alicloud_kms_key" "key" {
}


`, name)
}

// Case 用实例创建盘 7848
func TestAccAliCloudEcsDisk_basic7848(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_disk.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsDiskMap7848)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsDisk")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 100)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsDiskBasicDependence7848)
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
					"category":             "cloud_auto",
					"description":          "挂盘测试",
					"size":                 "20",
					"instance_id":          "${alicloud_ecs_instance.instance.id}",
					"disk_name":            name,
					"resource_group_id":    "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"delete_auto_snapshot": "false",
					"delete_with_instance": "false",
					"enable_auto_snapshot": "true",
					"bursting_enabled":     "true",
					"provisioned_iops":     "10",
					"snapshot_id":          "${alicloud_ecs_snapshot.createSnapshot.id}",
					"payment_type":         "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"category":             "cloud_auto",
						"description":          "挂盘测试",
						"size":                 "20",
						"instance_id":          CHECKSET,
						"disk_name":            name,
						"resource_group_id":    CHECKSET,
						"delete_auto_snapshot": "false",
						"delete_with_instance": "false",
						"enable_auto_snapshot": "true",
						"bursting_enabled":     "true",
						"provisioned_iops":     "10",
						"snapshot_id":          CHECKSET,
						"payment_type":         "Subscription",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description":          "挂盘",
					"size":                 "30",
					"disk_name":            name + "_update",
					"resource_group_id":    "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"delete_auto_snapshot": "true",
					"delete_with_instance": "true",
					"enable_auto_snapshot": "false",
					"bursting_enabled":     "false",
					"provisioned_iops":     "20",
					"payment_type":         "PayAsYouGo",
					"image_id":             "${alicloud_image.createImage.id}",
					"type":                 "offline",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "挂盘",
						"size":                 "30",
						"disk_name":            name + "_update",
						"resource_group_id":    CHECKSET,
						"delete_auto_snapshot": "true",
						"delete_with_instance": "true",
						"enable_auto_snapshot": "false",
						"bursting_enabled":     "false",
						"provisioned_iops":     "20",
						"payment_type":         "PayAsYouGo",
						"image_id":             CHECKSET,
						"type":                 "offline",
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
				ImportStateVerifyIgnore: []string{"advanced_features", "dry_run", "encrypt_algorithm", "image_id", "type"},
			},
		},
	})
}

var AlicloudEcsDiskMap7848 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudEcsDiskBasicDependence7848(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpc" {
  vpc_name = "w-TestVpc"
}

resource "alicloud_vswitch" "vsw" {
  vpc_id       = alicloud_vpc.vpc.id
  cidr_block   = "172.16.0.0/24"
  description  = "w-testVsw"
  zone_id      = "cn-hangzhou-i"
  vswitch_name = "w-test"
}

resource "alicloud_security_group" "sg" {
  description         = "w-test"
  security_group_name = "w-test"
  vpc_id              = alicloud_vpc.vpc.id
  security_group_type = "normal"
}

resource "alicloud_ecs_instance" "instance" {
  image_family = "acs:alibaba_cloud_linux_3_2104_lts_x64"
  vpc_attributes {
    vpc_id     = alicloud_vpc.vpc.id
    vswitch_id = alicloud_vswitch.vsw.id
  }
  security_group_id = alicloud_security_group.sg.id
  instance_type     = "ecs.g6.large"
  system_disk {
    size     = "20"
    category = "cloud_essd"
  }
  payment_type = "Subscription"
  period       = "1"
  period_unit  = "Week"
  status       = "Running"
}

resource "alicloud_image" "createImage" {
  instance_id = alicloud_ecs_instance.instance.id
}

resource "alicloud_ecs_disk" "createDisk" {
  category          = "cloud_essd"
  description       = "快照依赖盘"
  instance_id       = alicloud_ecs_instance.instance.id
  performance_level = "PL0"
  size              = "20"
  disk_name         = "快照依赖盘"
}

resource "alicloud_ecs_snapshot" "createSnapshot" {
  description = "挂盘测试"
  disk_id     = alicloud_ecs_disk.createDisk.id
}


`, name)
}

// Test Ecs Disk. <<< Resource test cases, automatically generated.
