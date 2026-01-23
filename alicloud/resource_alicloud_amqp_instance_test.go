package alicloud

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"

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

func TestUnitAliCloudAmqpInstance(t *testing.T) {
	p := Provider().(*schema.Provider).ResourcesMap
	d, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, nil)
	dCreate, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, nil)
	dCreate.MarkNewResource()
	dCreateMock, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, nil)
	dCreateMock.MarkNewResource()
	dCreateRenewalStatus, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, nil)
	dCreateRenewalStatus.MarkNewResource()
	for key, value := range map[string]interface{}{
		"instance_name":    "${var.name}",
		"instance_type":    "professional",
		"max_tps":          "1000",
		"payment_type":     "Subscription",
		"period":           1,
		"queue_capacity":   "50",
		"support_eip":      false,
		"logistics":        "logistics",
		"max_eip_tps":      "128",
		"renewal_duration": 1,
		"renewal_status":   "ManualRenewal",
		"storage_size":     "800",
	} {
		err := dCreate.Set(key, value)
		assert.Nil(t, err)
		err = d.Set(key, value)
		assert.Nil(t, err)
	}
	for key, value := range map[string]interface{}{
		"instance_name":    "${var.name}",
		"instance_type":    "professional",
		"max_tps":          "1000",
		"payment_type":     "Subscription",
		"period":           1,
		"queue_capacity":   "50",
		"support_eip":      true,
		"logistics":        "logistics",
		"renewal_duration": 1,
		"renewal_status":   "ManualRenewal",
		"storage_size":     "800",
	} {
		err := dCreateMock.Set(key, value)
		assert.Nil(t, err)
		err = d.Set(key, value)
		assert.Nil(t, err)
	}
	for key, value := range map[string]interface{}{
		"instance_name":  "${var.name}",
		"instance_type":  "professional",
		"max_tps":        "1000",
		"payment_type":   "Subscription",
		"period":         1,
		"queue_capacity": "50",
		"support_eip":    false,
		"logistics":      "logistics",
		"renewal_status": "AutoRenewal",
		"storage_size":   "800",
	} {
		err := dCreateRenewalStatus.Set(key, value)
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
		"Data": map[string]interface{}{
			"Instances": []interface{}{
				map[string]interface{}{
					"InstanceId":   "MockInstanceId",
					"Status":       "SERVING",
					"InstanceName": "instance_name",
					"InstanceType": "PROFESSIONAL",
					"SupportEIP":   false,
				},
			},
			"InstanceList": []interface{}{
				map[string]interface{}{
					"SubscriptionType":    "payment_type",
					"RenewalDuration":     1,
					"RenewStatus":         "ManualRenewal",
					"RenewalDurationUnit": "M",
					"InstanceID":          "MockInstanceId",
				},
			},
			"InstanceId": "MockInstanceId",
		},
		"Code": "Success",
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
			return nil, GetNotFoundErrorFromString(GetNotFoundMessage("alicloud_amqp_instance", "MockInstanceId"))
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
			result["InstanceId"] = "MockInstanceId"
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
		"CreateResponseCode": func(errorCode string) (map[string]interface{}, error) {
			result := map[string]interface{}{
				"Code": "Failed",
			}
			return result, nil
		},
		"UpdateResponse": func(errorCode string) (map[string]interface{}, error) {
			result := map[string]interface{}{
				"Success": "false",
			}
			return result, nil
		},
	}
	// Create
	t.Run("CreateClientAbnormal", func(t *testing.T) {
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewBssopenapiClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
			return nil, &tea.SDKError{
				Code:       String("loadEndpoint error"),
				Data:       String("loadEndpoint error"),
				Message:    String("loadEndpoint error"),
				StatusCode: tea.Int(400),
			}
		})
		err := resourceAliCloudAmqpInstanceCreate(d, rawClient)
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
		err := resourceAliCloudAmqpInstanceCreate(d, rawClient)
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
		err := resourceAliCloudAmqpInstanceCreate(dCreate, rawClient)
		patches.Reset()
		assert.Nil(t, err)
	})

	t.Run("CreateIsExpectedErrors", func(t *testing.T) {
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("NotApplicable")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["CreateNormal"]("")
		})
		err := resourceAliCloudAmqpInstanceCreate(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("CreateAttributesSupportEip", func(t *testing.T) {
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("NotApplicable")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["CreateNormal"]("")
		})
		err := resourceAliCloudAmqpInstanceCreate(dCreateMock, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("CreateAttributesRenewalStatus", func(t *testing.T) {
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("NotApplicable")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["CreateNormal"]("")
		})
		err := resourceAliCloudAmqpInstanceCreate(dCreateRenewalStatus, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("CreateResponseCode", func(t *testing.T) {
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("NotApplicable")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["CreateResponseCode"]("")
		})
		err := resourceAliCloudAmqpInstanceCreate(dCreate, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	// Set ID for Update and Delete Method
	d.SetId("MockInstanceId")
	// Update
	t.Run("UpdateClientAbnormal", func(t *testing.T) {
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewBssopenapiClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
			return nil, &tea.SDKError{
				Code:       String("loadEndpoint error"),
				Data:       String("loadEndpoint error"),
				Message:    String("loadEndpoint error"),
				StatusCode: tea.Int(400),
			}
		})

		err := resourceAliCloudAmqpInstanceUpdate(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateInstanceNameAbnormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"instance_name"} {
			switch p["alicloud_amqp_instance"].Schema[key].Type {
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
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, diff)
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
		err := resourceAliCloudAmqpInstanceUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateInstanceNameNormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"instance_name"} {
			switch p["alicloud_amqp_instance"].Schema[key].Type {
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
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, diff)
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
		err := resourceAliCloudAmqpInstanceUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.Nil(t, err)
	})

	t.Run("UpdateInstanceNameResponseAbnormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"instance_name"} {
			switch p["alicloud_amqp_instance"].Schema[key].Type {
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
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, diff)
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
			return responseMock["UpdateResponse"]("")
		})
		err := resourceAliCloudAmqpInstanceUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateSetRenewalAbnormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"renewal_status", "renewal_duration", "payment_type", "renewal_duration_unit"} {
			switch p["alicloud_amqp_instance"].Schema[key].Type {
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
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, diff)
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
		err := resourceAliCloudAmqpInstanceUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateSetRenewalIsExpectedErrors", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"renewal_status", "renewal_duration", "payment_type", "renewal_duration_unit"} {
			switch p["alicloud_amqp_instance"].Schema[key].Type {
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
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, diff)
		resourceData1.SetId(d.Id())
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("NotApplicable")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["UpdateNormal"]("")
		})
		err := resourceAliCloudAmqpInstanceUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateSetRenewalNormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"renewal_status", "renewal_duration", "payment_type", "renewal_duration_unit"} {
			switch p["alicloud_amqp_instance"].Schema[key].Type {
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
			}

		}
		resourceData1, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, diff)
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
		err := resourceAliCloudAmqpInstanceUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.Nil(t, err)
	})

	t.Run("UpdateSetRenewalResponseAbnormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"renewal_status", "renewal_duration", "payment_type", "renewal_duration_unit"} {
			switch p["alicloud_amqp_instance"].Schema[key].Type {
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
			}

		}
		resourceData1, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, diff)
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
			return responseMock["UpdateResponse"]("")
		})
		err := resourceAliCloudAmqpInstanceUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateModifyInstanceAbnormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"max_tps", "payment_type", "queue_capacity", "support_eip", "max_eip_tps", "support_eip", "storage_size", "modify_type"} {
			switch p["alicloud_amqp_instance"].Schema[key].Type {
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
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, diff)
		resourceData1.SetId(d.Id())
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "SetRenewal" {
				return responseMock["UpdateNormal"]("")
			}
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["UpdateNormal"]("")
		})
		err := resourceAliCloudAmqpInstanceUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateModifyInstanceNormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"max_tps", "payment_type", "queue_capacity", "support_eip", "max_eip_tps", "support_eip", "storage_size", "modify_type"} {
			switch p["alicloud_amqp_instance"].Schema[key].Type {
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
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, diff)
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
		err := resourceAliCloudAmqpInstanceUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.Nil(t, err)
	})

	t.Run("UpdateResponseAbnormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"max_tps", "payment_type", "queue_capacity", "support_eip", "max_eip_tps", "support_eip", "storage_size", "modify_type"} {
			switch p["alicloud_amqp_instance"].Schema[key].Type {
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
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, diff)
		resourceData1.SetId(d.Id())
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "SetRenewal" {
				return responseMock["UpdateNormal"]("")
			}
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["UpdateResponse"]("")
		})
		err := resourceAliCloudAmqpInstanceUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateModifyInstanceIsExpectedErrors", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"max_tps", "payment_type", "queue_capacity", "support_eip", "max_eip_tps", "support_eip", "storage_size", "modify_type"} {
			switch p["alicloud_amqp_instance"].Schema[key].Type {
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
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, diff)
		resourceData1.SetId(d.Id())
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "SetRenewal" {
				return responseMock["UpdateNormal"]("")
			}
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("NotApplicable")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["UpdateNormal"]("")
		})
		err := resourceAliCloudAmqpInstanceUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateAttributesSupportEip", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		diff.SetAttribute("support_eip", &terraform.ResourceAttrDiff{Old: "false", New: "true"})
		resourceData1, _ := schema.InternalMap(p["alicloud_amqp_instance"].Schema).Data(nil, diff)
		resourceData1.SetId(d.Id())
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("NotApplicable")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["UpdateNormal"]("")
		})
		err := resourceAliCloudAmqpInstanceUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	// Delete
	t.Run("DeleteNormal", func(t *testing.T) {
		err := resourceAliCloudAmqpInstanceDelete(d, rawClient)
		assert.Nil(t, err)
	})

	//Read
	t.Run("ReadDescribeAmqpInstanceNotFound", func(t *testing.T) {
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
		err := resourceAliCloudAmqpInstanceRead(d, rawClient)
		patcheDorequest.Reset()
		assert.Nil(t, err)
	})

	t.Run("ReadDescribeAmqpInstanceAbnormal", func(t *testing.T) {
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
		err := resourceAliCloudAmqpInstanceRead(d, rawClient)
		patcheDorequest.Reset()
		assert.NotNil(t, err)
	})
}

// Case 创建serverless实例 6128
func TestAccAliCloudAmqpInstance_basic6128(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AliCloudAmqpInstanceMap6128)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%samqpinstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAmqpInstanceBasicDependence6128)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  nil,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"payment_type":           "PayAsYouGo",
					"serverless_charge_type": "onDemand",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"payment_type": "PayAsYouGo",
					}),
				),
			},
			// Currently, Modifying the Edition parameter involves migrating the instance cluster. Before making changes, submit a ticket to the cloud service provider.
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"edition": "dedicated",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"edition": "dedicated",
			//		}),
			//	),
			//},
			{
				Config: testAccConfig(map[string]interface{}{
					"serverless_charge_type": "provisioned",
					"provisioned_capacity":   "2000",
					"modify_type":            "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"provisioned_capacity": "2000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_name": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name": name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"provisioned_capacity": "20000",
					"modify_type":          "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"provisioned_capacity": "20000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"support_eip": "true",
					"modify_type": "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"support_eip": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"support_tracing": "true",
					"modify_type":     "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"support_tracing": "true",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type", "renewal_duration", "renewal_duration_unit", "renewal_status"},
			},
		},
	})
}

func TestAccAliCloudAmqpInstance_basic6128_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AliCloudAmqpInstanceMap6128)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%samqpinstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAmqpInstanceBasicDependence6128)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  nil,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"payment_type":           "PayAsYouGo",
					"serverless_charge_type": "provisioned",
					"provisioned_capacity":   "20000",
					"edition":                "dedicated",
					"instance_name":          name,
					"support_eip":            "true",
					"support_tracing":        "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"payment_type":         "PayAsYouGo",
						"provisioned_capacity": "20000",
						"edition":              "dedicated",
						"instance_name":        name,
						"support_eip":          "true",
						"support_tracing":      "true",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type", "renewal_duration", "renewal_duration_unit", "renewal_status"},
			},
		},
	})
}

var AliCloudAmqpInstanceMap6128 = map[string]string{
	"status":        CHECKSET,
	"create_time":   CHECKSET,
	"instance_name": CHECKSET,
}

func AliCloudAmqpInstanceBasicDependence6128(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test Amqp Instance. >>> Resource test cases, automatically generated.
// Case Serverless版实例生命周期用例_按量_v1.1 12435
func TestAccAliCloudAmqpInstance_basic12435(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap12435)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence12435)
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
					"max_tps":                "50000",
					"period_cycle":           "Month",
					"max_connections":        "10000",
					"support_eip":            "false",
					"auto_renew":             "false",
					"renewal_status":         "ManualRenewal",
					"period":                 "2",
					"queue_capacity":         "2000",
					"instance_name":          name,
					"serverless_charge_type": "onDemand",
					"support_tracing":        "false",
					"payment_type":           "PayAsYouGo",
					"renewal_duration_unit":  "Month",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_tps":                CHECKSET,
						"period_cycle":           "Month",
						"max_connections":        "10000",
						"support_eip":            "false",
						"auto_renew":             "false",
						"renewal_status":         "ManualRenewal",
						"period":                 "2",
						"queue_capacity":         CHECKSET,
						"instance_name":          name,
						"serverless_charge_type": "onDemand",
						"support_tracing":        "false",
						"payment_type":           "PayAsYouGo",
						"renewal_duration_unit":  "Month",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"support_eip":          "true",
					"support_tracing":      "true",
					"tracing_storage_time": "15",
					"modify_type":          "Upgrade",
					"max_eip_tps":          "128",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"support_eip":          "true",
						"support_tracing":      "true",
						"tracing_storage_time": "15",
						"modify_type":          "Upgrade",
						"max_eip_tps":          CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap12435 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence12435(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-hangzhou"
}


`, name)
}

// Case Serverless版实例生命周期用例_弹性_v1.1 12444
func TestAccAliCloudAmqpInstance_basic12444(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap12444)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence12444)
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
					"max_tps":                "2000",
					"max_eip_tps":            "128",
					"period_cycle":           "Month",
					"max_connections":        "10000",
					"support_eip":            "true",
					"auto_renew":             "false",
					"renewal_status":         "ManualRenewal",
					"period":                 "2",
					"queue_capacity":         "2000",
					"instance_name":          name,
					"serverless_charge_type": "provisioned",
					"support_tracing":        "false",
					"payment_type":           "PayAsYouGo",
					"renewal_duration_unit":  "Month",
					"provisioned_capacity":   "2000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_tps":                CHECKSET,
						"max_eip_tps":            CHECKSET,
						"period_cycle":           "Month",
						"max_connections":        "10000",
						"support_eip":            "true",
						"auto_renew":             "false",
						"renewal_status":         "ManualRenewal",
						"period":                 "2",
						"queue_capacity":         CHECKSET,
						"instance_name":          name,
						"serverless_charge_type": "provisioned",
						"support_tracing":        "false",
						"payment_type":           "PayAsYouGo",
						"renewal_duration_unit":  "Month",
						"provisioned_capacity":   "2000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"support_tracing":      "true",
					"provisioned_capacity": "4000",
					"tracing_storage_time": "15",
					"modify_type":          "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"support_tracing":      "true",
						"provisioned_capacity": "4000",
						"tracing_storage_time": "15",
						"modify_type":          "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap12444 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence12444(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-hangzhou"
}


`, name)
}

// Case 专业版实例生命周期用例_1_v1.1 12446
func TestAccAliCloudAmqpInstance_basic12446(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap12446)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence12446)
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
					"renewal_duration":      "1",
					"max_tps":               "3000",
					"period_cycle":          "Month",
					"max_connections":       "2000",
					"support_eip":           "false",
					"auto_renew":            "false",
					"renewal_status":        "AutoRenewal",
					"period":                "3",
					"instance_name":         name,
					"support_tracing":       "false",
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "enterprise",
					"queue_capacity":        "200",
					"storage_size":          "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "1",
						"max_tps":               CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "2000",
						"support_eip":           "false",
						"auto_renew":            "false",
						"renewal_status":        "AutoRenewal",
						"period":                "3",
						"instance_name":         name,
						"support_tracing":       "false",
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "enterprise",
						"queue_capacity":        CHECKSET,
						"storage_size":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"max_tps":              "5000",
					"max_connections":      "3000",
					"support_eip":          "true",
					"renewal_status":       "ManualRenewal",
					"instance_name":        name + "_update",
					"support_tracing":      "true",
					"queue_capacity":       "300",
					"storage_size":         "1",
					"max_eip_tps":          "256",
					"modify_type":          "Upgrade",
					"tracing_storage_time": "7",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_tps":              CHECKSET,
						"max_connections":      "3000",
						"support_eip":          "true",
						"renewal_status":       "ManualRenewal",
						"instance_name":        name + "_update",
						"support_tracing":      "true",
						"queue_capacity":       CHECKSET,
						"storage_size":         CHECKSET,
						"max_eip_tps":          CHECKSET,
						"modify_type":          "Upgrade",
						"tracing_storage_time": "7",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap12446 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence12446(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-qingdao"
}


`, name)
}

// Case 专业版实例生命周期用例_0_v1.1 12437
func TestAccAliCloudAmqpInstance_basic12437(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap12437)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence12437)
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
					"renewal_duration":      "1",
					"max_tps":               "3000",
					"period_cycle":          "Month",
					"max_connections":       "2000",
					"support_eip":           "true",
					"auto_renew":            "false",
					"renewal_status":        "AutoRenewal",
					"period":                "3",
					"instance_name":         name,
					"support_tracing":       "false",
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "enterprise",
					"queue_capacity":        "200",
					"max_eip_tps":           "128",
					"storage_size":          "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "1",
						"max_tps":               CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "2000",
						"support_eip":           "true",
						"auto_renew":            "false",
						"renewal_status":        "AutoRenewal",
						"period":                "3",
						"instance_name":         name,
						"support_tracing":       "false",
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "enterprise",
						"queue_capacity":        CHECKSET,
						"max_eip_tps":           CHECKSET,
						"storage_size":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"max_tps":              "5000",
					"max_connections":      "3000",
					"renewal_status":       "ManualRenewal",
					"instance_name":        name + "_update",
					"support_tracing":      "true",
					"queue_capacity":       "300",
					"max_eip_tps":          "256",
					"storage_size":         "1",
					"modify_type":          "Upgrade",
					"tracing_storage_time": "7",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_tps":              CHECKSET,
						"max_connections":      "3000",
						"renewal_status":       "ManualRenewal",
						"instance_name":        name + "_update",
						"support_tracing":      "true",
						"queue_capacity":       CHECKSET,
						"max_eip_tps":          CHECKSET,
						"storage_size":         CHECKSET,
						"modify_type":          "Upgrade",
						"tracing_storage_time": "7",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap12437 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence12437(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-qingdao"
}


`, name)
}

// Case 创建企业版实例5 11166
func TestAccAliCloudAmqpInstance_basic11166(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11166)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11166)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":      "1",
					"max_tps":               "3000",
					"max_eip_tps":           "128",
					"period_cycle":          "Month",
					"max_connections":       "2000",
					"auto_renew":            "false",
					"support_eip":           "true",
					"renewal_status":        "AutoRenewal",
					"storage_size":          "0",
					"period":                "3",
					"queue_capacity":        "200",
					"instance_name":         name,
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "enterprise",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "1",
						"max_tps":               CHECKSET,
						"max_eip_tps":           CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "2000",
						"auto_renew":            "false",
						"support_eip":           "true",
						"renewal_status":        "AutoRenewal",
						"storage_size":          CHECKSET,
						"period":                "3",
						"queue_capacity":        CHECKSET,
						"instance_name":         name,
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "enterprise",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11166 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11166(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建企业版实例6 11167
func TestAccAliCloudAmqpInstance_basic11167(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11167)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11167)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":      "1",
					"max_tps":               "3000",
					"max_eip_tps":           "128",
					"period_cycle":          "Month",
					"max_connections":       "2000",
					"auto_renew":            "false",
					"support_eip":           "true",
					"renewal_status":        "AutoRenewal",
					"storage_size":          "0",
					"period":                "12",
					"queue_capacity":        "200",
					"instance_name":         name,
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "enterprise",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "1",
						"max_tps":               CHECKSET,
						"max_eip_tps":           CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "2000",
						"auto_renew":            "false",
						"support_eip":           "true",
						"renewal_status":        "AutoRenewal",
						"storage_size":          CHECKSET,
						"period":                "12",
						"queue_capacity":        CHECKSET,
						"instance_name":         name,
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "enterprise",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11167 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11167(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建企业版实例7 11171
func TestAccAliCloudAmqpInstance_basic11171(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11171)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11171)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":      "1",
					"max_tps":               "3000",
					"max_eip_tps":           "128",
					"period_cycle":          "Month",
					"max_connections":       "2000",
					"auto_renew":            "false",
					"support_eip":           "true",
					"renewal_status":        "AutoRenewal",
					"storage_size":          "0",
					"period":                "24",
					"queue_capacity":        "200",
					"instance_name":         name,
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "enterprise",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "1",
						"max_tps":               CHECKSET,
						"max_eip_tps":           CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "2000",
						"auto_renew":            "false",
						"support_eip":           "true",
						"renewal_status":        "AutoRenewal",
						"storage_size":          CHECKSET,
						"period":                "24",
						"queue_capacity":        CHECKSET,
						"instance_name":         name,
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "enterprise",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11171 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11171(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建企业版实例_升级铂金版实例 11164
func TestAccAliCloudAmqpInstance_basic11164(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11164)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11164)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":      "1",
					"max_tps":               "3000",
					"max_eip_tps":           "128",
					"period_cycle":          "Month",
					"max_connections":       "2000",
					"auto_renew":            "false",
					"support_eip":           "true",
					"renewal_status":        "AutoRenewal",
					"storage_size":          "0",
					"period":                "1",
					"queue_capacity":        "200",
					"instance_name":         name,
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "enterprise",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "1",
						"max_tps":               CHECKSET,
						"max_eip_tps":           CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "2000",
						"auto_renew":            "false",
						"support_eip":           "true",
						"renewal_status":        "AutoRenewal",
						"storage_size":          CHECKSET,
						"period":                "1",
						"queue_capacity":        CHECKSET,
						"instance_name":         name,
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "enterprise",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"max_tps":         "8000",
					"max_connections": "50000",
					"renewal_status":  "ManualRenewal",
					"storage_size":    "700",
					"queue_capacity":  "10000",
					"instance_name":   name + "_update",
					"instance_type":   "vip",
					"modify_type":     "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_tps":         CHECKSET,
						"max_connections": "50000",
						"renewal_status":  "ManualRenewal",
						"storage_size":    CHECKSET,
						"queue_capacity":  CHECKSET,
						"instance_name":   name + "_update",
						"instance_type":   "vip",
						"modify_type":     "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11164 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11164(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建预留实例+升级独享 11145
func TestAccAliCloudAmqpInstance_basic11145(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11145)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11145)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"serverless_charge_type": "provisioned",
					"payment_type":           "PayAsYouGo",
					"period_cycle":           "Month",
					"provisioned_capacity":   "2000",
					"edition":                "shared",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"serverless_charge_type": "provisioned",
						"payment_type":           "PayAsYouGo",
						"period_cycle":           "Month",
						"provisioned_capacity":   "2000",
						"edition":                "shared",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"provisioned_capacity": "5000",
					"edition":              "dedicated",
					"modify_type":          "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"provisioned_capacity": "5000",
						"edition":              "dedicated",
						"modify_type":          "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11145 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11145(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建企业版实例3_副本1753705329927 11116
func TestAccAliCloudAmqpInstance_basic11116(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11116)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11116)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":      "2",
					"max_tps":               "3000",
					"max_eip_tps":           "128",
					"period_cycle":          "Year",
					"max_connections":       "2000",
					"auto_renew":            "false",
					"support_eip":           "true",
					"renewal_status":        "AutoRenewal",
					"storage_size":          "0",
					"period":                "2",
					"queue_capacity":        "200",
					"instance_name":         name,
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Year",
					"instance_type":         "enterprise",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "2",
						"max_tps":               CHECKSET,
						"max_eip_tps":           CHECKSET,
						"period_cycle":          "Year",
						"max_connections":       "2000",
						"auto_renew":            "false",
						"support_eip":           "true",
						"renewal_status":        "AutoRenewal",
						"storage_size":          CHECKSET,
						"period":                "2",
						"queue_capacity":        CHECKSET,
						"instance_name":         name,
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Year",
						"instance_type":         "enterprise",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"max_tps":         "5000",
					"max_eip_tps":     "256",
					"max_connections": "3000",
					"renewal_status":  "ManualRenewal",
					"queue_capacity":  "300",
					"instance_name":   name + "_update",
					"modify_type":     "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_tps":         CHECKSET,
						"max_eip_tps":     CHECKSET,
						"max_connections": "3000",
						"renewal_status":  "ManualRenewal",
						"queue_capacity":  CHECKSET,
						"instance_name":   name + "_update",
						"modify_type":     "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11116 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11116(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建企业版实例4_副本1753857148347 11122
func TestAccAliCloudAmqpInstance_basic11122(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11122)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11122)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":      "6",
					"max_tps":               "3000",
					"max_eip_tps":           "128",
					"period_cycle":          "Month",
					"max_connections":       "2000",
					"auto_renew":            "false",
					"support_eip":           "true",
					"renewal_status":        "AutoRenewal",
					"storage_size":          "0",
					"period":                "6",
					"queue_capacity":        "200",
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "enterprise",
					"tracing_storage_time":  "15",
					"support_tracing":       "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "6",
						"max_tps":               CHECKSET,
						"max_eip_tps":           CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "2000",
						"auto_renew":            "false",
						"support_eip":           "true",
						"renewal_status":        "AutoRenewal",
						"storage_size":          CHECKSET,
						"period":                "6",
						"queue_capacity":        CHECKSET,
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "enterprise",
						"tracing_storage_time":  "15",
						"support_tracing":       "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":     "1",
					"support_eip":          "false",
					"tracing_storage_time": "3",
					"modify_type":          "Downgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":     "1",
						"support_eip":          "false",
						"tracing_storage_time": "3",
						"modify_type":          "Downgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11122 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11122(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建企业版实例2_副本1753695877077 11114
func TestAccAliCloudAmqpInstance_basic11114(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11114)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11114)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":      "1",
					"max_tps":               "3000",
					"max_eip_tps":           "128",
					"period_cycle":          "Month",
					"max_connections":       "2000",
					"auto_renew":            "true",
					"support_eip":           "true",
					"renewal_status":        "AutoRenewal",
					"period":                "1",
					"queue_capacity":        "200",
					"tracing_storage_time":  "7",
					"support_tracing":       "true",
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "enterprise",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "1",
						"max_tps":               CHECKSET,
						"max_eip_tps":           CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "2000",
						"auto_renew":            "true",
						"support_eip":           "true",
						"renewal_status":        "AutoRenewal",
						"period":                "1",
						"queue_capacity":        CHECKSET,
						"tracing_storage_time":  "7",
						"support_tracing":       "true",
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "enterprise",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"support_tracing": "false",
					"modify_type":     "Downgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"support_tracing": "false",
						"modify_type":     "Downgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11114 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11114(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建企业版实例_副本1753684310337 11110
func TestAccAliCloudAmqpInstance_basic11110(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11110)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11110)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"max_tps":         "3000",
					"period_cycle":    "Month",
					"max_connections": "2000",
					"queue_capacity":  "200",
					"payment_type":    "Subscription",
					"instance_type":   "enterprise",
					"support_eip":     "false",
					"renewal_status":  "ManualRenewal",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_tps":         CHECKSET,
						"period_cycle":    "Month",
						"max_connections": "2000",
						"queue_capacity":  CHECKSET,
						"payment_type":    "Subscription",
						"instance_type":   "enterprise",
						"support_eip":     "false",
						"renewal_status":  "ManualRenewal",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"modify_type":          "Upgrade",
					"support_tracing":      "true",
					"tracing_storage_time": "3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"modify_type":          "Upgrade",
						"support_tracing":      "true",
						"tracing_storage_time": "3",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11110 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11110(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建serverless实例_副本1753681518593 11107
func TestAccAliCloudAmqpInstance_basic11107(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11107)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11107)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"serverless_charge_type": "onDemand",
					"payment_type":           "PayAsYouGo",
					"period_cycle":           "Month",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"serverless_charge_type": "onDemand",
						"payment_type":           "PayAsYouGo",
						"period_cycle":           "Month",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"modify_type":     "Upgrade",
					"support_tracing": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"modify_type":     "Upgrade",
						"support_tracing": "true",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11107 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11107(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建预留实例 11104
func TestAccAliCloudAmqpInstance_basic11104(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11104)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11104)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"provisioned_capacity":   "2000",
					"edition":                "shared",
					"serverless_charge_type": "provisioned",
					"payment_type":           "PayAsYouGo",
					"period_cycle":           "Month",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"provisioned_capacity":   "2000",
						"edition":                "shared",
						"serverless_charge_type": "provisioned",
						"payment_type":           "PayAsYouGo",
						"period_cycle":           "Month",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"provisioned_capacity": "4000",
					"modify_type":          "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"provisioned_capacity": "4000",
						"modify_type":          "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11104 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11104(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建企业版实例4 11100
func TestAccAliCloudAmqpInstance_basic11100(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11100)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11100)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":      "6",
					"max_tps":               "3000",
					"max_eip_tps":           "128",
					"period_cycle":          "Month",
					"max_connections":       "2000",
					"auto_renew":            "false",
					"support_eip":           "true",
					"renewal_status":        "AutoRenewal",
					"storage_size":          "0",
					"period":                "6",
					"queue_capacity":        "200",
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "enterprise",
					"tracing_storage_time":  "15",
					"support_tracing":       "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "6",
						"max_tps":               CHECKSET,
						"max_eip_tps":           CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "2000",
						"auto_renew":            "false",
						"support_eip":           "true",
						"renewal_status":        "AutoRenewal",
						"storage_size":          CHECKSET,
						"period":                "6",
						"queue_capacity":        CHECKSET,
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "enterprise",
						"tracing_storage_time":  "15",
						"support_tracing":       "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":     "1",
					"support_eip":          "false",
					"tracing_storage_time": "3",
					"modify_type":          "Downgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":     "1",
						"support_eip":          "false",
						"tracing_storage_time": "3",
						"modify_type":          "Downgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11100 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11100(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建企业版实例3 11096
func TestAccAliCloudAmqpInstance_basic11096(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11096)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11096)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":      "2",
					"max_tps":               "3000",
					"max_eip_tps":           "128",
					"period_cycle":          "Year",
					"max_connections":       "2000",
					"auto_renew":            "false",
					"support_eip":           "true",
					"renewal_status":        "AutoRenewal",
					"storage_size":          "0",
					"period":                "2",
					"queue_capacity":        "200",
					"edition":               "shared",
					"instance_name":         name,
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Year",
					"instance_type":         "enterprise",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "2",
						"max_tps":               CHECKSET,
						"max_eip_tps":           CHECKSET,
						"period_cycle":          "Year",
						"max_connections":       "2000",
						"auto_renew":            "false",
						"support_eip":           "true",
						"renewal_status":        "AutoRenewal",
						"storage_size":          CHECKSET,
						"period":                "2",
						"queue_capacity":        CHECKSET,
						"edition":               "shared",
						"instance_name":         name,
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Year",
						"instance_type":         "enterprise",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"max_tps":         "5000",
					"max_eip_tps":     "256",
					"max_connections": "3000",
					"renewal_status":  "ManualRenewal",
					"queue_capacity":  "300",
					"instance_name":   name + "_update",
					"modify_type":     "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_tps":         CHECKSET,
						"max_eip_tps":     CHECKSET,
						"max_connections": "3000",
						"renewal_status":  "ManualRenewal",
						"queue_capacity":  CHECKSET,
						"instance_name":   name + "_update",
						"modify_type":     "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11096 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11096(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建企业版实例2 11094
func TestAccAliCloudAmqpInstance_basic11094(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11094)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11094)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":      "1",
					"max_tps":               "3000",
					"max_eip_tps":           "128",
					"period_cycle":          "Month",
					"max_connections":       "2000",
					"auto_renew":            "true",
					"support_eip":           "true",
					"renewal_status":        "AutoRenewal",
					"period":                "1",
					"queue_capacity":        "200",
					"tracing_storage_time":  "7",
					"support_tracing":       "true",
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "enterprise",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "1",
						"max_tps":               CHECKSET,
						"max_eip_tps":           CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "2000",
						"auto_renew":            "true",
						"support_eip":           "true",
						"renewal_status":        "AutoRenewal",
						"period":                "1",
						"queue_capacity":        CHECKSET,
						"tracing_storage_time":  "7",
						"support_tracing":       "true",
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "enterprise",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"support_tracing": "false",
					"modify_type":     "Downgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"support_tracing": "false",
						"modify_type":     "Downgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11094 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11094(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建企业版实例 11067
func TestAccAliCloudAmqpInstance_basic11067(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11067)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11067)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"max_tps":         "3000",
					"period_cycle":    "Month",
					"max_connections": "2000",
					"queue_capacity":  "200",
					"payment_type":    "Subscription",
					"instance_type":   "enterprise",
					"support_eip":     "false",
					"renewal_status":  "ManualRenewal",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_tps":         CHECKSET,
						"period_cycle":    "Month",
						"max_connections": "2000",
						"queue_capacity":  CHECKSET,
						"payment_type":    "Subscription",
						"instance_type":   "enterprise",
						"support_eip":     "false",
						"renewal_status":  "ManualRenewal",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"modify_type":          "Upgrade",
					"support_tracing":      "true",
					"tracing_storage_time": "3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"modify_type":          "Upgrade",
						"support_tracing":      "true",
						"tracing_storage_time": "3",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11067 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11067(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建serverless实例 11069
func TestAccAliCloudAmqpInstance_basic11069(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap11069)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence11069)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"serverless_charge_type": "onDemand",
					"payment_type":           "PayAsYouGo",
					"period_cycle":           "Month",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"serverless_charge_type": "onDemand",
						"payment_type":           "PayAsYouGo",
						"period_cycle":           "Month",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"modify_type":     "Upgrade",
					"support_tracing": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"modify_type":     "Upgrade",
						"support_tracing": "true",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap11069 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence11069(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 专业版实例生命周期用例_0 6997
func TestAccAliCloudAmqpInstance_basic6997(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap6997)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence6997)
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
					"renewal_duration":      "1",
					"max_tps":               "1000",
					"period_cycle":          "Month",
					"max_connections":       "1000",
					"support_eip":           "true",
					"auto_renew":            "false",
					"renewal_status":        "AutoRenewal",
					"period":                "3",
					"instance_name":         name,
					"support_tracing":       "false",
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "professional",
					"queue_capacity":        "100",
					"max_eip_tps":           "128",
					"storage_size":          "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "1",
						"max_tps":               CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "1000",
						"support_eip":           "true",
						"auto_renew":            "false",
						"renewal_status":        "AutoRenewal",
						"period":                "3",
						"instance_name":         name,
						"support_tracing":       "false",
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "professional",
						"queue_capacity":        CHECKSET,
						"max_eip_tps":           CHECKSET,
						"storage_size":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"max_tps":              "5000",
					"max_connections":      "1500",
					"renewal_status":       "ManualRenewal",
					"instance_name":        name + "_update",
					"support_tracing":      "true",
					"instance_type":        "enterprise",
					"queue_capacity":       "300",
					"max_eip_tps":          "256",
					"storage_size":         "1",
					"modify_type":          "Upgrade",
					"tracing_storage_time": "7",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_tps":              CHECKSET,
						"max_connections":      "1500",
						"renewal_status":       "ManualRenewal",
						"instance_name":        name + "_update",
						"support_tracing":      "true",
						"instance_type":        "enterprise",
						"queue_capacity":       CHECKSET,
						"max_eip_tps":          CHECKSET,
						"storage_size":         CHECKSET,
						"modify_type":          "Upgrade",
						"tracing_storage_time": "7",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap6997 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence6997(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 专业版实例生命周期用例_1 7057
func TestAccAliCloudAmqpInstance_basic7057(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap7057)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence7057)
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
					"renewal_duration":      "2",
					"max_tps":               "1000",
					"period_cycle":          "Month",
					"max_connections":       "1000",
					"support_eip":           "false",
					"auto_renew":            "false",
					"renewal_status":        "AutoRenewal",
					"period":                "24",
					"instance_name":         name,
					"support_tracing":       "true",
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "professional",
					"tracing_storage_time":  "3",
					"queue_capacity":        "100",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "2",
						"max_tps":               CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "1000",
						"support_eip":           "false",
						"auto_renew":            "false",
						"renewal_status":        "AutoRenewal",
						"period":                "24",
						"instance_name":         name,
						"support_tracing":       "true",
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "professional",
						"tracing_storage_time":  "3",
						"queue_capacity":        CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":     "1",
					"max_tps":              "1500",
					"max_connections":      "1500",
					"support_eip":          "true",
					"instance_name":        name + "_update",
					"tracing_storage_time": "7",
					"queue_capacity":       "300",
					"max_eip_tps":          "256",
					"storage_size":         "1",
					"modify_type":          "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":     "1",
						"max_tps":              CHECKSET,
						"max_connections":      "1500",
						"support_eip":          "true",
						"instance_name":        name + "_update",
						"tracing_storage_time": "7",
						"queue_capacity":       CHECKSET,
						"max_eip_tps":          CHECKSET,
						"storage_size":         CHECKSET,
						"modify_type":          "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap7057 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence7057(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 专业版实例生命周期用例_4 6877
func TestAccAliCloudAmqpInstance_basic6877(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap6877)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence6877)
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
					"renewal_duration":      "2",
					"max_tps":               "1000",
					"period_cycle":          "Month",
					"max_connections":       "1600",
					"support_eip":           "false",
					"auto_renew":            "false",
					"renewal_status":        "AutoRenewal",
					"period":                "6",
					"instance_name":         name,
					"support_tracing":       "true",
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "professional",
					"tracing_storage_time":  "3",
					"queue_capacity":        "400",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "2",
						"max_tps":               CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "1600",
						"support_eip":           "false",
						"auto_renew":            "false",
						"renewal_status":        "AutoRenewal",
						"period":                "6",
						"instance_name":         name,
						"support_tracing":       "true",
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "professional",
						"tracing_storage_time":  "3",
						"queue_capacity":        CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":     "1",
					"max_tps":              "1500",
					"max_connections":      "2400",
					"support_eip":          "true",
					"instance_name":        name + "_update",
					"tracing_storage_time": "7",
					"queue_capacity":       "600",
					"max_eip_tps":          "256",
					"storage_size":         "1",
					"modify_type":          "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":     "1",
						"max_tps":              CHECKSET,
						"max_connections":      "2400",
						"support_eip":          "true",
						"instance_name":        name + "_update",
						"tracing_storage_time": "7",
						"queue_capacity":       CHECKSET,
						"max_eip_tps":          CHECKSET,
						"storage_size":         CHECKSET,
						"modify_type":          "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap6877 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence6877(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 专业版实例生命周期用例_5 7269
func TestAccAliCloudAmqpInstance_basic7269(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap7269)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence7269)
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
					"renewal_duration":      "1",
					"max_tps":               "1000",
					"period_cycle":          "Month",
					"max_connections":       "1000",
					"support_eip":           "true",
					"auto_renew":            "false",
					"renewal_status":        "AutoRenewal",
					"period":                "6",
					"instance_name":         name,
					"support_tracing":       "false",
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "professional",
					"queue_capacity":        "100",
					"max_eip_tps":           "128",
					"storage_size":          "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "1",
						"max_tps":               CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "1000",
						"support_eip":           "true",
						"auto_renew":            "false",
						"renewal_status":        "AutoRenewal",
						"period":                "6",
						"instance_name":         name,
						"support_tracing":       "false",
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "professional",
						"queue_capacity":        CHECKSET,
						"max_eip_tps":           CHECKSET,
						"storage_size":          CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap7269 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence7269(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 专业版实例生命周期用例_3 7055
func TestAccAliCloudAmqpInstance_basic7055(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap7055)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence7055)
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
					"renewal_duration":      "2",
					"max_tps":               "1000",
					"period_cycle":          "Year",
					"max_connections":       "1000",
					"support_eip":           "false",
					"auto_renew":            "false",
					"renewal_status":        "AutoRenewal",
					"period":                "1",
					"instance_name":         name,
					"support_tracing":       "true",
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Year",
					"instance_type":         "professional",
					"tracing_storage_time":  "3",
					"queue_capacity":        "100",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "2",
						"max_tps":               CHECKSET,
						"period_cycle":          "Year",
						"max_connections":       "1000",
						"support_eip":           "false",
						"auto_renew":            "false",
						"renewal_status":        "AutoRenewal",
						"period":                "1",
						"instance_name":         name,
						"support_tracing":       "true",
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Year",
						"instance_type":         "professional",
						"tracing_storage_time":  "3",
						"queue_capacity":        CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":     "1",
					"max_tps":              "1500",
					"max_connections":      "1500",
					"support_eip":          "true",
					"instance_name":        name + "_update",
					"tracing_storage_time": "7",
					"queue_capacity":       "300",
					"max_eip_tps":          "256",
					"storage_size":         "1",
					"modify_type":          "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":     "1",
						"max_tps":              CHECKSET,
						"max_connections":      "1500",
						"support_eip":          "true",
						"instance_name":        name + "_update",
						"tracing_storage_time": "7",
						"queue_capacity":       CHECKSET,
						"max_eip_tps":          CHECKSET,
						"storage_size":         CHECKSET,
						"modify_type":          "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap7055 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence7055(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case Serverless版实例生命周期用例 6991
func TestAccAliCloudAmqpInstance_basic6991(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap6991)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence6991)
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
					"max_tps":                "50000",
					"max_eip_tps":            "128",
					"period_cycle":           "Month",
					"max_connections":        "10000",
					"support_eip":            "true",
					"auto_renew":             "false",
					"renewal_status":         "ManualRenewal",
					"period":                 "2",
					"queue_capacity":         "2000",
					"instance_name":          name,
					"serverless_charge_type": "onDemand",
					"support_tracing":        "false",
					"payment_type":           "PayAsYouGo",
					"renewal_duration_unit":  "Month",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_tps":                CHECKSET,
						"max_eip_tps":            CHECKSET,
						"period_cycle":           "Month",
						"max_connections":        "10000",
						"support_eip":            "true",
						"auto_renew":             "false",
						"renewal_status":         "ManualRenewal",
						"period":                 "2",
						"queue_capacity":         CHECKSET,
						"instance_name":          name,
						"serverless_charge_type": "onDemand",
						"support_tracing":        "false",
						"payment_type":           "PayAsYouGo",
						"renewal_duration_unit":  "Month",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"support_tracing":      "true",
					"tracing_storage_time": "15",
					"modify_type":          "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"support_tracing":      "true",
						"tracing_storage_time": "15",
						"modify_type":          "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap6991 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence6991(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 专业版实例生命周期用例_6 7270
func TestAccAliCloudAmqpInstance_basic7270(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap7270)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence7270)
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
					"renewal_duration":      "1",
					"max_tps":               "1000",
					"period_cycle":          "Month",
					"max_connections":       "1000",
					"support_eip":           "true",
					"auto_renew":            "false",
					"renewal_status":        "AutoRenewal",
					"period":                "12",
					"instance_name":         name,
					"support_tracing":       "false",
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "professional",
					"queue_capacity":        "100",
					"max_eip_tps":           "128",
					"storage_size":          "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "1",
						"max_tps":               CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "1000",
						"support_eip":           "true",
						"auto_renew":            "false",
						"renewal_status":        "AutoRenewal",
						"period":                "12",
						"instance_name":         name,
						"support_tracing":       "false",
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "professional",
						"queue_capacity":        CHECKSET,
						"max_eip_tps":           CHECKSET,
						"storage_size":          CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap7270 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence7270(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 专业版实例生命周期用例_2 7056
func TestAccAliCloudAmqpInstance_basic7056(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap7056)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence7056)
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
					"renewal_duration":      "2",
					"max_tps":               "1000",
					"period_cycle":          "Month",
					"max_connections":       "1000",
					"support_eip":           "false",
					"auto_renew":            "false",
					"renewal_status":        "AutoRenewal",
					"period":                "12",
					"instance_name":         name,
					"support_tracing":       "true",
					"payment_type":          "Subscription",
					"renewal_duration_unit": "Month",
					"instance_type":         "professional",
					"tracing_storage_time":  "3",
					"queue_capacity":        "100",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "2",
						"max_tps":               CHECKSET,
						"period_cycle":          "Month",
						"max_connections":       "1000",
						"support_eip":           "false",
						"auto_renew":            "false",
						"renewal_status":        "AutoRenewal",
						"period":                "12",
						"instance_name":         name,
						"support_tracing":       "true",
						"payment_type":          "Subscription",
						"renewal_duration_unit": "Month",
						"instance_type":         "professional",
						"tracing_storage_time":  "3",
						"queue_capacity":        CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":     "1",
					"max_tps":              "1500",
					"max_connections":      "1500",
					"support_eip":          "true",
					"instance_name":        name + "_update",
					"tracing_storage_time": "7",
					"queue_capacity":       "300",
					"max_eip_tps":          "256",
					"storage_size":         "1",
					"modify_type":          "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":     "1",
						"max_tps":              CHECKSET,
						"max_connections":      "1500",
						"support_eip":          "true",
						"instance_name":        name + "_update",
						"tracing_storage_time": "7",
						"queue_capacity":       CHECKSET,
						"max_eip_tps":          CHECKSET,
						"storage_size":         CHECKSET,
						"modify_type":          "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap7056 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence7056(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 对接资源组测试用例_资源组完整用例 7538
func TestAccAliCloudAmqpInstance_basic7538(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_amqp_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudAmqpInstanceMap7538)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AmqpServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAmqpInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccamqp%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAmqpInstanceBasicDependence7538)
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
					"support_eip":            "true",
					"auto_renew":             "false",
					"period":                 "1",
					"support_tracing":        "true",
					"period_cycle":           "Month",
					"serverless_charge_type": "onDemand",
					"payment_type":           "PayAsYouGo",
					"tracing_storage_time":   "15",
					"max_eip_tps":            "128",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"support_eip":            "true",
						"auto_renew":             "false",
						"period":                 "1",
						"support_tracing":        "true",
						"period_cycle":           "Month",
						"serverless_charge_type": "onDemand",
						"payment_type":           "PayAsYouGo",
						"tracing_storage_time":   "15",
						"max_eip_tps":            CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"support_eip":     "false",
					"support_tracing": "false",
					"modify_type":     "Downgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"support_eip":     "false",
						"support_tracing": "false",
						"modify_type":     "Downgrade",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"support_eip":     "true",
					"support_tracing": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"support_eip":     "true",
						"support_tracing": "true",
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
				ImportStateVerifyIgnore: []string{"auto_renew", "modify_type", "period", "period_cycle", "serverless_charge_type"},
			},
		},
	})
}

var AlicloudAmqpInstanceMap7538 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudAmqpInstanceBasicDependence7538(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Test Amqp Instance. <<< Resource test cases, automatically generated.
