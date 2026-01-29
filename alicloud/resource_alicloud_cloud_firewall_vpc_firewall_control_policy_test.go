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

func TestUnitAliCloudCloudFirewallVpcFirewallControlPolicy(t *testing.T) {
	p := Provider().(*schema.Provider).ResourcesMap
	dInit, _ := schema.InternalMap(p["alicloud_cloud_firewall_vpc_firewall_control_policy"].Schema).Data(nil, nil)
	dExisted, _ := schema.InternalMap(p["alicloud_cloud_firewall_vpc_firewall_control_policy"].Schema).Data(nil, nil)
	dInit.MarkNewResource()
	attributes := map[string]interface{}{
		"acl_action":       "CreateVpcFirewallControlPolicyValue",
		"application_name": "CreateVpcFirewallControlPolicyValue",
		"description":      "CreateVpcFirewallControlPolicyValue",
		"dest_port":        "CreateVpcFirewallControlPolicyValue",
		"dest_port_group":  "CreateVpcFirewallControlPolicyValue",
		"dest_port_type":   "CreateVpcFirewallControlPolicyValue",
		"destination":      "CreateVpcFirewallControlPolicyValue",
		"destination_type": "CreateVpcFirewallControlPolicyValue",
		"lang":             "CreateVpcFirewallControlPolicyValue",
		"member_uid":       "CreateVpcFirewallControlPolicyValue",
		"order":            1,
		"proto":            "CreateVpcFirewallControlPolicyValue",
		"release":          true,
		"source":           "CreateVpcFirewallControlPolicyValue",
		"source_type":      "CreateVpcFirewallControlPolicyValue",
		"vpc_firewall_id":  "CreateVpcFirewallControlPolicyValue",
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
		// DescribeVpcFirewallControlPolicy
		"Policys": []interface{}{
			map[string]interface{}{
				"AclAction":             "CreateVpcFirewallControlPolicyValue",
				"AclUuid":               "CreateVpcFirewallControlPolicyValue",
				"ApplicationId":         "DefaultValue",
				"ApplicationName":       "CreateVpcFirewallControlPolicyValue",
				"Description":           "CreateVpcFirewallControlPolicyValue",
				"DestPort":              "CreateVpcFirewallControlPolicyValue",
				"DestPortGroup":         "CreateVpcFirewallControlPolicyValue",
				"DestPortGroupPorts":    []interface{}{},
				"DestPortType":          "CreateVpcFirewallControlPolicyValue",
				"Destination":           "CreateVpcFirewallControlPolicyValue",
				"DestinationGroupCidrs": []interface{}{},
				"DestinationGroupType":  "DefaultValue",
				"DestinationType":       "CreateVpcFirewallControlPolicyValue",
				"HitTimes":              0,
				"MemberUid":             "CreateVpcFirewallControlPolicyValue",
				"Order":                 1,
				"Proto":                 "CreateVpcFirewallControlPolicyValue",
				"Release":               "true",
				"Source":                "CreateVpcFirewallControlPolicyValue",
				"SourceGroupCidrs":      []interface{}{},
				"SourceGroupType":       "DefaultValue",
				"SourceType":            "CreateVpcFirewallControlPolicyValue",
			},
		},
	}
	CreateMockResponse := map[string]interface{}{
		// CreateVpcFirewallControlPolicy
		"AclUuid": "CreateVpcFirewallControlPolicyValue",
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
		return nil, GetNotFoundErrorFromString(GetNotFoundMessage("alicloud_cloud_firewall_vpc_firewall_control_policy", errorCode))
	}
	successResponseMock := func(operationMockResponse map[string]interface{}) (map[string]interface{}, error) {
		if len(operationMockResponse) > 0 {
			mapMerge(ReadMockResponse, operationMockResponse)
		}
		return ReadMockResponse, nil
	}

	// Create
	patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewCloudfirewallClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:    String("loadEndpoint error"),
			Data:    String("loadEndpoint error"),
			Message: String("loadEndpoint error"),
		}
	})
	err = resourceAliCloudCloudFirewallVpcFirewallControlPolicyCreate(dInit, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	ReadMockResponseDiff := map[string]interface{}{
		// DescribeVpcFirewallControlPolicy Response
		"Policys": []interface{}{
			map[string]interface{}{
				"AclUuid": "CreateVpcFirewallControlPolicyValue",
			},
		},
	}
	errorCodes := []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1 // a counter used to cover retry scenario; the same below
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "CreateVpcFirewallControlPolicy" {
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
		err := resourceAliCloudCloudFirewallVpcFirewallControlPolicyCreate(dInit, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_cloud_firewall_vpc_firewall_control_policy"].Schema).Data(dInit.State(), nil)
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
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewCloudfirewallClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:    String("loadEndpoint error"),
			Data:    String("loadEndpoint error"),
			Message: String("loadEndpoint error"),
		}
	})
	err = resourceAliCloudCloudFirewallVpcFirewallControlPolicyUpdate(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	// ModifyVpcFirewallControlPolicy
	attributesDiff := map[string]interface{}{
		"acl_action":       "ModifyVpcFirewallControlPolicyValue",
		"application_name": "ModifyVpcFirewallControlPolicyValue",
		"description":      "ModifyVpcFirewallControlPolicyValue",
		"dest_port":        "ModifyVpcFirewallControlPolicyValue",
		"dest_port_group":  "ModifyVpcFirewallControlPolicyValue",
		"dest_port_type":   "ModifyVpcFirewallControlPolicyValue",
		"destination":      "ModifyVpcFirewallControlPolicyValue",
		"destination_type": "ModifyVpcFirewallControlPolicyValue",
		"proto":            "ModifyVpcFirewallControlPolicyValue",
		"release":          false,
		"source":           "ModifyVpcFirewallControlPolicyValue",
		"source_type":      "ModifyVpcFirewallControlPolicyValue",
	}
	diff, err := newInstanceDiff("alicloud_cloud_firewall_vpc_firewall_control_policy", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_cloud_firewall_vpc_firewall_control_policy"].Schema).Data(dInit.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeVpcFirewallControlPolicy Response
		"Policys": []interface{}{
			map[string]interface{}{
				"AclAction":       "ModifyVpcFirewallControlPolicyValue",
				"ApplicationName": "ModifyVpcFirewallControlPolicyValue",
				"Description":     "ModifyVpcFirewallControlPolicyValue",
				"DestPort":        "ModifyVpcFirewallControlPolicyValue",
				"DestPortGroup":   "ModifyVpcFirewallControlPolicyValue",
				"DestPortType":    "ModifyVpcFirewallControlPolicyValue",
				"Destination":     "ModifyVpcFirewallControlPolicyValue",
				"DestinationType": "ModifyVpcFirewallControlPolicyValue",
				"Proto":           "ModifyVpcFirewallControlPolicyValue",
				"Release":         "ModifyVpcFirewallControlPolicyValue",
				"Source":          "ModifyVpcFirewallControlPolicyValue",
				"SourceType":      "ModifyVpcFirewallControlPolicyValue",
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "ModifyVpcFirewallControlPolicy" {
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
		err := resourceAliCloudCloudFirewallVpcFirewallControlPolicyUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_cloud_firewall_vpc_firewall_control_policy"].Schema).Data(dExisted.State(), nil)
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
	errorCodes = []string{"NonRetryableError", "Throttling", "nil", "{}"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DescribeVpcFirewallControlPolicy" {
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
		err := resourceAliCloudCloudFirewallVpcFirewallControlPolicyRead(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "{}":
			assert.Nil(t, err)
		}
	}

	// Delete
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewCloudfirewallClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:    String("loadEndpoint error"),
			Data:    String("loadEndpoint error"),
			Message: String("loadEndpoint error"),
		}
	})
	err = resourceAliCloudCloudFirewallVpcFirewallControlPolicyDelete(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	attributesDiff = map[string]interface{}{}
	diff, err = newInstanceDiff("alicloud_cloud_firewall_vpc_firewall_control_policy", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_cloud_firewall_vpc_firewall_control_policy"].Schema).Data(dInit.State(), diff)
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DeleteVpcFirewallControlPolicy" {
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
		err := resourceAliCloudCloudFirewallVpcFirewallControlPolicyDelete(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "nil":
			assert.Nil(t, err)
		}
	}
}

// Test CloudFirewall VpcFirewallControlPolicy. >>> Resource test cases, automatically generated.
// Case VPC边界安全策略_域名_0 12208
func TestAccAliCloudCloudFirewallVpcFirewallControlPolicy_basic12208(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_vpc_firewall_control_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallVpcFirewallControlPolicyMap12208)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallVpcFirewallControlPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12208)
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
					"order":                 "1",
					"destination":           "baidu.com",
					"description":           "miaoshu",
					"source_type":           "net",
					"dest_port":             "8082/8082",
					"application_name_list": []string{},
					"acl_action":            "log",
					"lang":                  "zh",
					"destination_type":      "domain",
					"vpc_firewall_id":       "${alicloud_cen_instance.cen.id}",
					"source":                "8.8.8.8/32",
					"dest_port_type":        "port",
					"proto":                 "TCP",
					"release":               "true",
					"member_uid":            "1511928242963727",
					"repeat_days":           []string{},
					"application_name":      "HTTP",
					"repeat_type":           "Permanent",
					"domain_resolve_type":   "FQDN",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"order":               "1",
						"destination":         "baidu.com",
						"description":         "miaoshu",
						"source_type":         "net",
						"dest_port":           "8082/8082",
						"acl_action":          "log",
						"lang":                "zh",
						"destination_type":    "domain",
						"vpc_firewall_id":     CHECKSET,
						"source":              "8.8.8.8/32",
						"dest_port_type":      "port",
						"proto":               "TCP",
						"release":             CHECKSET,
						"member_uid":          CHECKSET,
						"application_name":    "HTTP",
						"repeat_type":         "Permanent",
						"domain_resolve_type": "FQDN",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"destination": "8.8.8.8/32",
					"description": "test",
					"dest_port":   "8080/8080",
					"application_name_list": []string{
						"HTTP"},
					"acl_action":       "drop",
					"lang":             "en",
					"destination_type": "net",
					"repeat_type":      "None",
					"end_time":         "1766676600",
					"start_time":       "1766592000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"destination":             "8.8.8.8/32",
						"description":             "test",
						"dest_port":               "8080/8080",
						"application_name_list.#": "1",
						"acl_action":              "drop",
						"lang":                    "en",
						"destination_type":        "net",
						"repeat_type":             "None",
						"end_time":                "1766676600",
						"start_time":              "1766592000",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallVpcFirewallControlPolicyMap12208 = map[string]string{
	"hit_times":                 CHECKSET,
	"dest_port_group_ports.#":   CHECKSET,
	"destination_group_type":    CHECKSET,
	"destination_group_cidrs.#": CHECKSET,
	"source_group_type":         CHECKSET,
	"source_group_cidrs.#":      CHECKSET,
	"create_time":               CHECKSET,
	"acl_uuid":                  CHECKSET,
	"application_id":            CHECKSET,
}

func AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12208(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_instance" "cen" {
  cen_instance_name = "yqc-test-CenInstanceName1"
}


`, name)
}

// Case VPC边界安全策略_地址簿_策略有效期(总是)_0 12203
func TestAccAliCloudCloudFirewallVpcFirewallControlPolicy_basic12203(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_vpc_firewall_control_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallVpcFirewallControlPolicyMap12203)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallVpcFirewallControlPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12203)
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
					"order":       "1",
					"destination": "${alicloud_cloud_firewall_address_book.addressBook.group_name}",
					"description": "miaoshu",
					"source_type": "group",
					"application_name_list": []string{
						"HTTPS", "HTTP", "FTP"},
					"acl_action":       "log",
					"lang":             "en",
					"destination_type": "group",
					"vpc_firewall_id":  "${alicloud_cen_instance.cen.id}",
					"source":           "${alicloud_cloud_firewall_address_book.addressBook.group_name}",
					"dest_port_type":   "group",
					"proto":            "TCP",
					"repeat_days":      []string{},
					"member_uid":       "1511928242963727",
					"release":          "true",
					"repeat_type":      "Permanent",
					"dest_port_group":  "${alicloud_cloud_firewall_address_book.portAddressBook.group_name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"order":                   "1",
						"destination":             CHECKSET,
						"description":             "miaoshu",
						"source_type":             "group",
						"application_name_list.#": "3",
						"acl_action":              "log",
						"lang":                    "en",
						"destination_type":        "group",
						"vpc_firewall_id":         CHECKSET,
						"source":                  CHECKSET,
						"dest_port_type":          "group",
						"proto":                   "TCP",
						"member_uid":              CHECKSET,
						"release":                 CHECKSET,
						"repeat_type":             "Permanent",
						"dest_port_group":         CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"destination": "${alicloud_cloud_firewall_address_book.addressBook2.group_name}",
					"description": "test",
					"application_name_list": []string{
						"SSH", "SSL"},
					"lang":            "zh",
					"source":          "${alicloud_cloud_firewall_address_book.addressBook2.group_name}",
					"release":         "false",
					"dest_port_group": "${alicloud_cloud_firewall_address_book.portAddressBook2.group_name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"destination":             CHECKSET,
						"description":             "test",
						"application_name_list.#": "2",
						"lang":                    "zh",
						"source":                  CHECKSET,
						"release":                 CHECKSET,
						"dest_port_group":         CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallVpcFirewallControlPolicyMap12203 = map[string]string{
	"hit_times":                 CHECKSET,
	"dest_port_group_ports.#":   CHECKSET,
	"destination_group_type":    CHECKSET,
	"destination_group_cidrs.#": CHECKSET,
	"source_group_type":         CHECKSET,
	"source_group_cidrs.#":      CHECKSET,
	"create_time":               CHECKSET,
	"acl_uuid":                  CHECKSET,
	"application_id":            CHECKSET,
}

func AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12203(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_instance" "cen" {
  cen_instance_name = "yqc-test-CenInstanceName1"
}

resource "alicloud_cloud_firewall_address_book" "addressBook" {
  group_name   = "yqc-test-addressBook"
  description  = "tqc-description"
  group_type   = "ip"
  lang         = "zh"
  address_list = ["172.16.1.2/32"]
}

resource "alicloud_cloud_firewall_address_book" "portAddressBook" {
  group_name   = "yqc-address-port"
  group_type   = "port"
  address_list = ["8080/8082"]
  description  = "test"
}

resource "alicloud_cloud_firewall_address_book" "addressBook2" {
  group_name   = "yqc-test-addressBook2"
  description  = "tqc-description2"
  group_type   = "ip"
  lang         = "zh"
  address_list = ["172.16.1.3/32"]
}

resource "alicloud_cloud_firewall_address_book" "portAddressBook2" {
  group_name   = "yqc-address-port2"
  group_type   = "port"
  address_list = ["8081/8081"]
  description  = "test2"
}


`, name)
}

// Case VPC边界安全策略_IP_策略有效期(重复周期)_0 12202
func TestAccAliCloudCloudFirewallVpcFirewallControlPolicy_basic12202(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_vpc_firewall_control_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallVpcFirewallControlPolicyMap12202)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallVpcFirewallControlPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12202)
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
					"order":                 "1",
					"destination":           "8.8.8.8/32",
					"description":           "miaoshu",
					"source_type":           "net",
					"dest_port":             "8082/8082",
					"application_name_list": []string{},
					"acl_action":            "log",
					"lang":                  "zh",
					"destination_type":      "net",
					"vpc_firewall_id":       "${alicloud_cen_instance.cen.id}",
					"source":                "8.8.8.8/32",
					"dest_port_type":        "port",
					"proto":                 "TCP",
					"repeat_days": []string{
						"1", "2", "3"},
					"application_name":  "HTTP",
					"member_uid":        "1511928242963727",
					"end_time":          "1766676600",
					"start_time":        "1765209600",
					"repeat_end_time":   "19:00",
					"repeat_type":       "Weekly",
					"repeat_start_time": "18:00",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"order":             "1",
						"destination":       "8.8.8.8/32",
						"description":       "miaoshu",
						"source_type":       "net",
						"dest_port":         "8082/8082",
						"acl_action":        "log",
						"lang":              "zh",
						"destination_type":  "net",
						"vpc_firewall_id":   CHECKSET,
						"source":            "8.8.8.8/32",
						"dest_port_type":    "port",
						"proto":             "TCP",
						"repeat_days.#":     "3",
						"application_name":  "HTTP",
						"member_uid":        CHECKSET,
						"end_time":          "1766676600",
						"start_time":        "1765209600",
						"repeat_end_time":   "19:00",
						"repeat_type":       "Weekly",
						"repeat_start_time": "18:00",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description":       "test",
					"dest_port":         "8080/8080",
					"lang":              "en",
					"proto":             "UDP",
					"application_name":  "ANY",
					"end_time":          "1766678400",
					"start_time":        "1765211400",
					"repeat_end_time":   "19:30",
					"repeat_start_time": "18:30",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":       "test",
						"dest_port":         "8080/8080",
						"lang":              "en",
						"proto":             "UDP",
						"application_name":  "ANY",
						"end_time":          "1766678400",
						"start_time":        "1765211400",
						"repeat_end_time":   "19:30",
						"repeat_start_time": "18:30",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallVpcFirewallControlPolicyMap12202 = map[string]string{
	"hit_times":                 CHECKSET,
	"dest_port_group_ports.#":   CHECKSET,
	"destination_group_type":    CHECKSET,
	"destination_group_cidrs.#": CHECKSET,
	"source_group_type":         CHECKSET,
	"source_group_cidrs.#":      CHECKSET,
	"create_time":               CHECKSET,
	"acl_uuid":                  CHECKSET,
	"application_id":            CHECKSET,
}

func AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12202(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_instance" "cen" {
  cen_instance_name = "yqc-test-CenInstanceName1"
}


`, name)
}

// Case VPC边界安全策略_ip修改为域名_0 12201
func TestAccAliCloudCloudFirewallVpcFirewallControlPolicy_basic12201(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_vpc_firewall_control_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallVpcFirewallControlPolicyMap12201)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallVpcFirewallControlPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12201)
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
					"order":                 "1",
					"destination":           "8.8.8.8/32",
					"description":           "miaoshu",
					"source_type":           "net",
					"dest_port":             "8082/8082",
					"application_name_list": []string{},
					"acl_action":            "log",
					"lang":                  "zh",
					"destination_type":      "net",
					"vpc_firewall_id":       "${alicloud_cen_instance.cen.id}",
					"source":                "8.8.8.8/32",
					"dest_port_type":        "port",
					"proto":                 "TCP",
					"release":               "true",
					"end_time":              "4102414200",
					"start_time":            "1765209600",
					"repeat_end_time":       "18:00",
					"repeat_start_time":     "19:00",
					"member_uid":            "1511928242963727",
					"repeat_days":           []string{},
					"application_name":      "HTTP",
					"repeat_type":           "Daily",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"order":             "1",
						"destination":       "8.8.8.8/32",
						"description":       "miaoshu",
						"source_type":       "net",
						"dest_port":         "8082/8082",
						"acl_action":        "log",
						"lang":              "zh",
						"destination_type":  "net",
						"vpc_firewall_id":   CHECKSET,
						"source":            "8.8.8.8/32",
						"dest_port_type":    "port",
						"proto":             "TCP",
						"release":           CHECKSET,
						"end_time":          "4102414200",
						"start_time":        "1765209600",
						"repeat_end_time":   "18:00",
						"repeat_start_time": "19:00",
						"member_uid":        CHECKSET,
						"application_name":  "HTTP",
						"repeat_type":       "Daily",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"destination": "baidu.com",
					"description": "test",
					"dest_port":   "8080/8080",
					"application_name_list": []string{
						"HTTP"},
					"lang":                "en",
					"destination_type":    "domain",
					"domain_resolve_type": "FQDN",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"destination":             "baidu.com",
						"description":             "test",
						"dest_port":               "8080/8080",
						"application_name_list.#": "1",
						"lang":                    "en",
						"destination_type":        "domain",
						"domain_resolve_type":     "FQDN",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallVpcFirewallControlPolicyMap12201 = map[string]string{
	"hit_times":                 CHECKSET,
	"dest_port_group_ports.#":   CHECKSET,
	"destination_group_type":    CHECKSET,
	"destination_group_cidrs.#": CHECKSET,
	"source_group_type":         CHECKSET,
	"source_group_cidrs.#":      CHECKSET,
	"create_time":               CHECKSET,
	"acl_uuid":                  CHECKSET,
	"application_id":            CHECKSET,
}

func AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12201(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_instance" "cen" {
  cen_instance_name = "yqc-test-CenInstanceName1"
}


`, name)
}

// Case VPC边界安全策略_IP_策略有效期(总是)_0 12206
func TestAccAliCloudCloudFirewallVpcFirewallControlPolicy_basic12206(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_vpc_firewall_control_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallVpcFirewallControlPolicyMap12206)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallVpcFirewallControlPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12206)
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
					"order":       "1",
					"destination": "8.8.8.8/32",
					"description": "miaoshu",
					"source_type": "net",
					"dest_port":   "8082/8082",
					"application_name_list": []string{
						"HTTPS", "HTTP", "FTP"},
					"acl_action":       "log",
					"lang":             "zh",
					"destination_type": "net",
					"vpc_firewall_id":  "${alicloud_cen_instance.cen.id}",
					"source":           "8.8.8.8/32",
					"dest_port_type":   "port",
					"proto":            "TCP",
					"repeat_days":      []string{},
					"member_uid":       "1511928242963727",
					"release":          "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"order":                   "1",
						"destination":             "8.8.8.8/32",
						"description":             "miaoshu",
						"source_type":             "net",
						"dest_port":               "8082/8082",
						"application_name_list.#": "3",
						"acl_action":              "log",
						"lang":                    "zh",
						"destination_type":        "net",
						"vpc_firewall_id":         CHECKSET,
						"source":                  "8.8.8.8/32",
						"dest_port_type":          "port",
						"proto":                   "TCP",
						"member_uid":              CHECKSET,
						"release":                 CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test",
					"dest_port":   "8080/8080",
					"application_name_list": []string{
						"SSH", "SSL"},
					"acl_action": "drop",
					"lang":       "en",
					"release":    "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":             "test",
						"dest_port":               "8080/8080",
						"application_name_list.#": "2",
						"acl_action":              "drop",
						"lang":                    "en",
						"release":                 CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallVpcFirewallControlPolicyMap12206 = map[string]string{
	"hit_times":                 CHECKSET,
	"dest_port_group_ports.#":   CHECKSET,
	"destination_group_type":    CHECKSET,
	"destination_group_cidrs.#": CHECKSET,
	"source_group_type":         CHECKSET,
	"source_group_cidrs.#":      CHECKSET,
	"create_time":               CHECKSET,
	"acl_uuid":                  CHECKSET,
	"application_id":            CHECKSET,
}

func AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12206(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_instance" "cen" {
  cen_instance_name = "yqc-test-CenInstanceName1"
}


`, name)
}

// Case VPC边界安全策略_域名_更新_0 12204
func TestAccAliCloudCloudFirewallVpcFirewallControlPolicy_basic12204(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_vpc_firewall_control_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallVpcFirewallControlPolicyMap12204)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallVpcFirewallControlPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12204)
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
					"order":                 "1",
					"destination":           "baidu.com",
					"description":           "miaoshu",
					"source_type":           "net",
					"dest_port":             "8082/8082",
					"application_name_list": []string{},
					"acl_action":            "log",
					"lang":                  "zh",
					"destination_type":      "domain",
					"vpc_firewall_id":       "${alicloud_cen_instance.cen.id}",
					"source":                "8.8.8.8/32",
					"dest_port_type":        "port",
					"proto":                 "TCP",
					"release":               "true",
					"member_uid":            "1511928242963727",
					"repeat_days":           []string{},
					"application_name":      "HTTP",
					"repeat_type":           "Permanent",
					"domain_resolve_type":   "FQDN",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"order":               "1",
						"destination":         "baidu.com",
						"description":         "miaoshu",
						"source_type":         "net",
						"dest_port":           "8082/8082",
						"acl_action":          "log",
						"lang":                "zh",
						"destination_type":    "domain",
						"vpc_firewall_id":     CHECKSET,
						"source":              "8.8.8.8/32",
						"dest_port_type":      "port",
						"proto":               "TCP",
						"release":             CHECKSET,
						"member_uid":          CHECKSET,
						"application_name":    "HTTP",
						"repeat_type":         "Permanent",
						"domain_resolve_type": "FQDN",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test",
					"dest_port":   "8080/8080",
					"application_name_list": []string{
						"HTTP", "HTTPS", "SSH"},
					"acl_action":          "drop",
					"domain_resolve_type": "DNS",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":             "test",
						"dest_port":               "8080/8080",
						"application_name_list.#": "3",
						"acl_action":              "drop",
						"domain_resolve_type":     "DNS",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallVpcFirewallControlPolicyMap12204 = map[string]string{
	"hit_times":                 CHECKSET,
	"dest_port_group_ports.#":   CHECKSET,
	"destination_group_type":    CHECKSET,
	"destination_group_cidrs.#": CHECKSET,
	"source_group_type":         CHECKSET,
	"source_group_cidrs.#":      CHECKSET,
	"create_time":               CHECKSET,
	"acl_uuid":                  CHECKSET,
	"application_id":            CHECKSET,
}

func AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12204(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_instance" "cen" {
  cen_instance_name = "yqc-test-CenInstanceName1"
}


`, name)
}

// Case VPC边界安全策略_ip修改为域名 11929
func TestAccAliCloudCloudFirewallVpcFirewallControlPolicy_basic11929(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_vpc_firewall_control_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallVpcFirewallControlPolicyMap11929)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallVpcFirewallControlPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence11929)
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
					"order":                 "1",
					"destination":           "8.8.8.8/32",
					"description":           "miaoshu",
					"source_type":           "net",
					"dest_port":             "8082/8082",
					"application_name_list": []string{},
					"acl_action":            "log",
					"lang":                  "zh",
					"destination_type":      "net",
					"vpc_firewall_id":       "${alicloud_cen_instance.cen.id}",
					"source":                "8.8.8.8/32",
					"dest_port_type":        "port",
					"proto":                 "TCP",
					"release":               "true",
					"end_time":              "4102414200",
					"start_time":            "1765209600",
					"repeat_end_time":       "18:00",
					"repeat_start_time":     "19:00",
					"member_uid":            "1511928242963727",
					"repeat_days":           []string{},
					"application_name":      "HTTP",
					"repeat_type":           "Daily",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"order":             "1",
						"destination":       "8.8.8.8/32",
						"description":       "miaoshu",
						"source_type":       "net",
						"dest_port":         "8082/8082",
						"acl_action":        "log",
						"lang":              "zh",
						"destination_type":  "net",
						"vpc_firewall_id":   CHECKSET,
						"source":            "8.8.8.8/32",
						"dest_port_type":    "port",
						"proto":             "TCP",
						"release":           CHECKSET,
						"end_time":          "4102414200",
						"start_time":        "1765209600",
						"repeat_end_time":   "18:00",
						"repeat_start_time": "19:00",
						"member_uid":        CHECKSET,
						"application_name":  "HTTP",
						"repeat_type":       "Daily",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"destination": "baidu.com",
					"description": "test",
					"dest_port":   "8080/8080",
					"application_name_list": []string{
						"HTTP"},
					"lang":                "en",
					"destination_type":    "domain",
					"domain_resolve_type": "FQDN",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"destination":             "baidu.com",
						"description":             "test",
						"dest_port":               "8080/8080",
						"application_name_list.#": "1",
						"lang":                    "en",
						"destination_type":        "domain",
						"domain_resolve_type":     "FQDN",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallVpcFirewallControlPolicyMap11929 = map[string]string{
	"hit_times":                 CHECKSET,
	"dest_port_group_ports.#":   CHECKSET,
	"destination_group_type":    CHECKSET,
	"destination_group_cidrs.#": CHECKSET,
	"source_group_type":         CHECKSET,
	"source_group_cidrs.#":      CHECKSET,
	"create_time":               CHECKSET,
	"acl_uuid":                  CHECKSET,
	"application_id":            CHECKSET,
}

func AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence11929(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_instance" "cen" {
  cen_instance_name = "yqc-test-CenInstanceName1"
}


`, name)
}

// Case VPC边界安全策略_IP_策略有效期(重复周期) 12022
func TestAccAliCloudCloudFirewallVpcFirewallControlPolicy_basic12022(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_vpc_firewall_control_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallVpcFirewallControlPolicyMap12022)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallVpcFirewallControlPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12022)
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
					"order":                 "1",
					"destination":           "8.8.8.8/32",
					"description":           "miaoshu",
					"source_type":           "net",
					"dest_port":             "8082/8082",
					"application_name_list": []string{},
					"acl_action":            "log",
					"lang":                  "zh",
					"destination_type":      "net",
					"vpc_firewall_id":       "${alicloud_cen_instance.cen.id}",
					"source":                "8.8.8.8/32",
					"dest_port_type":        "port",
					"proto":                 "TCP",
					"repeat_days": []string{
						"1", "2", "3"},
					"application_name":  "HTTP",
					"member_uid":        "1511928242963727",
					"end_time":          "1766676600",
					"start_time":        "1765209600",
					"repeat_end_time":   "19:00",
					"repeat_type":       "Weekly",
					"repeat_start_time": "18:00",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"order":             "1",
						"destination":       "8.8.8.8/32",
						"description":       "miaoshu",
						"source_type":       "net",
						"dest_port":         "8082/8082",
						"acl_action":        "log",
						"lang":              "zh",
						"destination_type":  "net",
						"vpc_firewall_id":   CHECKSET,
						"source":            "8.8.8.8/32",
						"dest_port_type":    "port",
						"proto":             "TCP",
						"repeat_days.#":     "3",
						"application_name":  "HTTP",
						"member_uid":        CHECKSET,
						"end_time":          "1766676600",
						"start_time":        "1765209600",
						"repeat_end_time":   "19:00",
						"repeat_type":       "Weekly",
						"repeat_start_time": "18:00",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description":       "test",
					"dest_port":         "8080/8080",
					"lang":              "en",
					"proto":             "UDP",
					"application_name":  "ANY",
					"end_time":          "1766678400",
					"start_time":        "1765211400",
					"repeat_end_time":   "19:30",
					"repeat_start_time": "18:30",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":       "test",
						"dest_port":         "8080/8080",
						"lang":              "en",
						"proto":             "UDP",
						"application_name":  "ANY",
						"end_time":          "1766678400",
						"start_time":        "1765211400",
						"repeat_end_time":   "19:30",
						"repeat_start_time": "18:30",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallVpcFirewallControlPolicyMap12022 = map[string]string{
	"hit_times":                 CHECKSET,
	"dest_port_group_ports.#":   CHECKSET,
	"destination_group_type":    CHECKSET,
	"destination_group_cidrs.#": CHECKSET,
	"source_group_type":         CHECKSET,
	"source_group_cidrs.#":      CHECKSET,
	"create_time":               CHECKSET,
	"acl_uuid":                  CHECKSET,
	"application_id":            CHECKSET,
}

func AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12022(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_instance" "cen" {
  cen_instance_name = "yqc-test-CenInstanceName1"
}


`, name)
}

// Case VPC边界安全策略_地址簿_策略有效期(总是) 12034
func TestAccAliCloudCloudFirewallVpcFirewallControlPolicy_basic12034(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_vpc_firewall_control_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallVpcFirewallControlPolicyMap12034)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallVpcFirewallControlPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12034)
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
					"order":       "1",
					"destination": "${alicloud_cloud_firewall_address_book.addressBook.group_name}",
					"description": "miaoshu",
					"source_type": "group",
					"application_name_list": []string{
						"HTTPS", "HTTP", "FTP"},
					"acl_action":       "log",
					"lang":             "zh",
					"destination_type": "group",
					"vpc_firewall_id":  "${alicloud_cen_instance.cen.id}",
					"source":           "${alicloud_cloud_firewall_address_book.addressBook.group_name}",
					"dest_port_type":   "group",
					"proto":            "TCP",
					"repeat_days":      []string{},
					"member_uid":       "1511928242963727",
					"release":          "true",
					"repeat_type":      "Permanent",
					"dest_port_group":  "${alicloud_cloud_firewall_address_book.portAddressBook.group_name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"order":                   "1",
						"destination":             CHECKSET,
						"description":             "miaoshu",
						"source_type":             "group",
						"application_name_list.#": "3",
						"acl_action":              "log",
						"lang":                    "zh",
						"destination_type":        "group",
						"vpc_firewall_id":         CHECKSET,
						"source":                  CHECKSET,
						"dest_port_type":          "group",
						"proto":                   "TCP",
						"member_uid":              CHECKSET,
						"release":                 CHECKSET,
						"repeat_type":             "Permanent",
						"dest_port_group":         CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"destination": "${alicloud_cloud_firewall_address_book.addressBook2.group_name}",
					"description": "test",
					"application_name_list": []string{
						"SSH", "SSL"},
					"source":          "${alicloud_cloud_firewall_address_book.addressBook2.group_name}",
					"release":         "false",
					"dest_port_group": "${alicloud_cloud_firewall_address_book.portAddressBook2.group_name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"destination":             CHECKSET,
						"description":             "test",
						"application_name_list.#": "2",
						"source":                  CHECKSET,
						"release":                 CHECKSET,
						"dest_port_group":         CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallVpcFirewallControlPolicyMap12034 = map[string]string{
	"hit_times":                 CHECKSET,
	"dest_port_group_ports.#":   CHECKSET,
	"destination_group_type":    CHECKSET,
	"destination_group_cidrs.#": CHECKSET,
	"source_group_type":         CHECKSET,
	"source_group_cidrs.#":      CHECKSET,
	"create_time":               CHECKSET,
	"acl_uuid":                  CHECKSET,
	"application_id":            CHECKSET,
}

func AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12034(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_instance" "cen" {
  cen_instance_name = "yqc-test-CenInstanceName1"
}

resource "alicloud_cloud_firewall_address_book" "addressBook" {
  group_name   = "yqc-test-addressBook"
  description  = "tqc-description"
  group_type   = "ip"
  lang         = "zh"
  address_list = ["172.16.1.2/32"]
}

resource "alicloud_cloud_firewall_address_book" "portAddressBook" {
  group_name   = "yqc-address-port"
  group_type   = "port"
  address_list = ["8080/8082"]
  description  = "test"
}

resource "alicloud_cloud_firewall_address_book" "addressBook2" {
  group_name   = "yqc-test-addressBook2"
  description  = "tqc-description2"
  group_type   = "ip"
  lang         = "zh"
  address_list = ["172.16.1.3/32"]
}

resource "alicloud_cloud_firewall_address_book" "portAddressBook2" {
  group_name   = "yqc-address-port2"
  group_type   = "port"
  address_list = ["8081/8081"]
  description  = "test2"
}


`, name)
}

// Case VPC边界安全策略_域名_更新 12041
func TestAccAliCloudCloudFirewallVpcFirewallControlPolicy_basic12041(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_vpc_firewall_control_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallVpcFirewallControlPolicyMap12041)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallVpcFirewallControlPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12041)
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
					"order":                 "1",
					"destination":           "baidu.com",
					"description":           "miaoshu",
					"source_type":           "net",
					"dest_port":             "8082/8082",
					"application_name_list": []string{},
					"acl_action":            "log",
					"lang":                  "zh",
					"destination_type":      "domain",
					"vpc_firewall_id":       "${alicloud_cen_instance.cen.id}",
					"source":                "8.8.8.8/32",
					"dest_port_type":        "port",
					"proto":                 "TCP",
					"release":               "true",
					"member_uid":            "1511928242963727",
					"repeat_days":           []string{},
					"application_name":      "HTTP",
					"repeat_type":           "Permanent",
					"domain_resolve_type":   "FQDN",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"order":               "1",
						"destination":         "baidu.com",
						"description":         "miaoshu",
						"source_type":         "net",
						"dest_port":           "8082/8082",
						"acl_action":          "log",
						"lang":                "zh",
						"destination_type":    "domain",
						"vpc_firewall_id":     CHECKSET,
						"source":              "8.8.8.8/32",
						"dest_port_type":      "port",
						"proto":               "TCP",
						"release":             CHECKSET,
						"member_uid":          CHECKSET,
						"application_name":    "HTTP",
						"repeat_type":         "Permanent",
						"domain_resolve_type": "FQDN",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test",
					"dest_port":   "8080/8080",
					"application_name_list": []string{
						"HTTP", "HTTPS", "SSH"},
					"acl_action":          "drop",
					"domain_resolve_type": "DNS",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":             "test",
						"dest_port":               "8080/8080",
						"application_name_list.#": "3",
						"acl_action":              "drop",
						"domain_resolve_type":     "DNS",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallVpcFirewallControlPolicyMap12041 = map[string]string{
	"hit_times":                 CHECKSET,
	"dest_port_group_ports.#":   CHECKSET,
	"destination_group_type":    CHECKSET,
	"destination_group_cidrs.#": CHECKSET,
	"source_group_type":         CHECKSET,
	"source_group_cidrs.#":      CHECKSET,
	"create_time":               CHECKSET,
	"acl_uuid":                  CHECKSET,
	"application_id":            CHECKSET,
}

func AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12041(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_instance" "cen" {
  cen_instance_name = "yqc-test-CenInstanceName1"
}


`, name)
}

// Case VPC边界安全策略_域名 12040
func TestAccAliCloudCloudFirewallVpcFirewallControlPolicy_basic12040(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_vpc_firewall_control_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallVpcFirewallControlPolicyMap12040)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallVpcFirewallControlPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12040)
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
					"order":                 "1",
					"destination":           "baidu.com",
					"description":           "miaoshu",
					"source_type":           "net",
					"dest_port":             "8082/8082",
					"application_name_list": []string{},
					"acl_action":            "log",
					"lang":                  "zh",
					"destination_type":      "domain",
					"vpc_firewall_id":       "${alicloud_cen_instance.cen.id}",
					"source":                "8.8.8.8/32",
					"dest_port_type":        "port",
					"proto":                 "TCP",
					"release":               "true",
					"member_uid":            "1511928242963727",
					"repeat_days":           []string{},
					"application_name":      "HTTP",
					"repeat_type":           "Permanent",
					"domain_resolve_type":   "FQDN",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"order":               "1",
						"destination":         "baidu.com",
						"description":         "miaoshu",
						"source_type":         "net",
						"dest_port":           "8082/8082",
						"acl_action":          "log",
						"lang":                "zh",
						"destination_type":    "domain",
						"vpc_firewall_id":     CHECKSET,
						"source":              "8.8.8.8/32",
						"dest_port_type":      "port",
						"proto":               "TCP",
						"release":             CHECKSET,
						"member_uid":          CHECKSET,
						"application_name":    "HTTP",
						"repeat_type":         "Permanent",
						"domain_resolve_type": "FQDN",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"destination": "8.8.8.8/32",
					"description": "test",
					"dest_port":   "8080/8080",
					"application_name_list": []string{
						"HTTP"},
					"acl_action":       "drop",
					"lang":             "en",
					"destination_type": "net",
					"repeat_type":      "None",
					"end_time":         "1765380600",
					"start_time":       "1765296000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"destination":             "8.8.8.8/32",
						"description":             "test",
						"dest_port":               "8080/8080",
						"application_name_list.#": "1",
						"acl_action":              "drop",
						"lang":                    "en",
						"destination_type":        "net",
						"repeat_type":             "None",
						"end_time":                "1765380600",
						"start_time":              "1765296000",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallVpcFirewallControlPolicyMap12040 = map[string]string{
	"hit_times":                 CHECKSET,
	"dest_port_group_ports.#":   CHECKSET,
	"destination_group_type":    CHECKSET,
	"destination_group_cidrs.#": CHECKSET,
	"source_group_type":         CHECKSET,
	"source_group_cidrs.#":      CHECKSET,
	"create_time":               CHECKSET,
	"acl_uuid":                  CHECKSET,
	"application_id":            CHECKSET,
}

func AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12040(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_instance" "cen" {
  cen_instance_name = "yqc-test-CenInstanceName1"
}


`, name)
}

// Case VPC边界安全策略_IP_策略有效期(总是) 12032
func TestAccAliCloudCloudFirewallVpcFirewallControlPolicy_basic12032(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_vpc_firewall_control_policy.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallVpcFirewallControlPolicyMap12032)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallVpcFirewallControlPolicy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12032)
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
					"order":       "1",
					"destination": "8.8.8.8/32",
					"description": "miaoshu",
					"source_type": "net",
					"dest_port":   "8082/8082",
					"application_name_list": []string{
						"HTTPS", "HTTP", "FTP"},
					"acl_action":       "log",
					"lang":             "zh",
					"destination_type": "net",
					"vpc_firewall_id":  "${alicloud_cen_instance.cen.id}",
					"source":           "8.8.8.8/32",
					"dest_port_type":   "port",
					"proto":            "TCP",
					"repeat_days":      []string{},
					"member_uid":       "1511928242963727",
					"release":          "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"order":                   "1",
						"destination":             "8.8.8.8/32",
						"description":             "miaoshu",
						"source_type":             "net",
						"dest_port":               "8082/8082",
						"application_name_list.#": "3",
						"acl_action":              "log",
						"lang":                    "zh",
						"destination_type":        "net",
						"vpc_firewall_id":         CHECKSET,
						"source":                  "8.8.8.8/32",
						"dest_port_type":          "port",
						"proto":                   "TCP",
						"member_uid":              CHECKSET,
						"release":                 CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test",
					"dest_port":   "8080/8080",
					"application_name_list": []string{
						"SSH", "SSL"},
					"acl_action": "drop",
					"lang":       "en",
					"release":    "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":             "test",
						"dest_port":               "8080/8080",
						"application_name_list.#": "2",
						"acl_action":              "drop",
						"lang":                    "en",
						"release":                 CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallVpcFirewallControlPolicyMap12032 = map[string]string{
	"hit_times":                 CHECKSET,
	"dest_port_group_ports.#":   CHECKSET,
	"destination_group_type":    CHECKSET,
	"destination_group_cidrs.#": CHECKSET,
	"source_group_type":         CHECKSET,
	"source_group_cidrs.#":      CHECKSET,
	"create_time":               CHECKSET,
	"acl_uuid":                  CHECKSET,
	"application_id":            CHECKSET,
}

func AlicloudCloudFirewallVpcFirewallControlPolicyBasicDependence12032(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cen_instance" "cen" {
  cen_instance_name = "yqc-test-CenInstanceName1"
}


`, name)
}

// Test CloudFirewall VpcFirewallControlPolicy. <<< Resource test cases, automatically generated.
