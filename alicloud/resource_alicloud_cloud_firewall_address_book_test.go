package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test CloudFirewall AddressBook. >>> Resource test cases, automatically generated.
// Case resourceCase_20260622_oZG09z 12951
func TestAccAliCloudCloudFirewallAddressBook_basic12951(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallAddressBookMap12951)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallAddressBookBasicDependence12951)
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
					"group_name": "IPv4AddressBook75",
					"asset_member_uids": []string{
						"1580394279029207"},
					"description":  "AddressBook create by Terraform",
					"tag_relation": "and",
					"group_type":   "ip",
					"lang":         "zh",
					"address_list": []string{
						"1.1.1.1/32", "2.2.2.2/32", "3.3.3.3/32", "4.4.4.4/32"},
					"auto_add_tag_ecs": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":          "IPv4AddressBook75",
						"asset_member_uids.#": "1",
						"description":         "AddressBook create by Terraform",
						"tag_relation":        "and",
						"group_type":          "ip",
						"lang":                "zh",
						"address_list.#":      "4",
						"auto_add_tag_ecs":    "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tag_relation": "or",
					"address_list": []string{
						"5.5.5.5/32"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tag_relation":   "or",
						"address_list.#": "1",
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
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallAddressBookMap12951 = map[string]string{
	"address_list_count": CHECKSET,
	"reference_count":    CHECKSET,
}

func AlicloudCloudFirewallAddressBookBasicDependence12951(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case resourceCase_20260707_6lVaF3 12952
func TestAccAliCloudCloudFirewallAddressBook_basic12952(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallAddressBookMap12952)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallAddressBookBasicDependence12952)
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
					"group_name":  "DomainAddressBook8",
					"description": "AddressBook create by Terraform",
					"group_type":  "domain",
					"lang":        "zh",
					"address_list": []string{
						"1.com", "2.com", "3.com", "4.com"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     "DomainAddressBook8",
						"description":    "AddressBook create by Terraform",
						"group_type":     "domain",
						"lang":           "zh",
						"address_list.#": "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":  "AddressBook164",
					"description": "AddressBook update by Terraform",
					"address_list": []string{
						"1.com", "2.com"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     "AddressBook164",
						"description":    "AddressBook update by Terraform",
						"address_list.#": "2",
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
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallAddressBookMap12952 = map[string]string{
	"address_list_count": CHECKSET,
	"reference_count":    CHECKSET,
}

func AlicloudCloudFirewallAddressBookBasicDependence12952(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case resourceCase_20260624_R4KXek_clone_1 12953
func TestAccAliCloudCloudFirewallAddressBook_basic12953(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallAddressBookMap12953)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallAddressBookBasicDependence12953)
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
					"group_name":  "IPv4AddressBook81",
					"description": "AddressBook create by Terraform",
					"group_type":  "ip",
					"lang":        "zh",
					"address_list": []string{
						"1.1.1.1/32", "2.2.2.2/32", "3.3.3.3/32", "4.4.4.4/32"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     "IPv4AddressBook81",
						"description":    "AddressBook create by Terraform",
						"group_type":     "ip",
						"lang":           "zh",
						"address_list.#": "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":  "IPv4AddressBookUpdate183",
					"description": "AddressBook update by Terraform",
					"address_list": []string{
						"1.1.1.1/32", "2.2.2.2/32"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     "IPv4AddressBookUpdate183",
						"description":    "AddressBook update by Terraform",
						"address_list.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": "IPv4AddressBookUpdate288",
					"address_list": []string{
						"1.1.1.1/32", "2.2.2.2/32"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     "IPv4AddressBookUpdate288",
						"address_list.#": "2",
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
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallAddressBookMap12953 = map[string]string{
	"address_list_count": CHECKSET,
	"reference_count":    CHECKSET,
}

func AlicloudCloudFirewallAddressBookBasicDependence12953(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case resourceCase_20260706_C2KUqs 12954
func TestAccAliCloudCloudFirewallAddressBook_basic12954(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallAddressBookMap12954)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallAddressBookBasicDependence12954)
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
					"group_name":  "PortAddressBook43",
					"description": "AddressBook create by Terraform",
					"group_type":  "port",
					"lang":        "zh",
					"address_list": []string{
						"8080/8080"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     "PortAddressBook43",
						"description":    "AddressBook create by Terraform",
						"group_type":     "port",
						"lang":           "zh",
						"address_list.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":  "PortAddressBookUpdate178",
					"description": "AddressBook update by Terraform",
					"address_list": []string{
						"8082/8082"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     "PortAddressBookUpdate178",
						"description":    "AddressBook update by Terraform",
						"address_list.#": "1",
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
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallAddressBookMap12954 = map[string]string{
	"address_list_count": CHECKSET,
	"reference_count":    CHECKSET,
}

func AlicloudCloudFirewallAddressBookBasicDependence12954(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case resourceCase_20260624_R4KXek_clone_0_clone_0 12955
func TestAccAliCloudCloudFirewallAddressBook_basic12955(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallAddressBookMap12955)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallAddressBookBasicDependence12955)
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
					"group_name": "yqc-test-01",
					"asset_member_uids": []string{
						"1580394279029207"},
					"asset_region_resource_types": []map[string]interface{}{
						{
							"asset_region_id": "all",
							"resource_type": []map[string]interface{}{
								{
									"ipv6": []map[string]interface{}{
										{
											"api_gateway_e_ipv6": "true",
											"ai_gateway_e_ipv6":  "true",
											"nlb_ipv6":           "true",
											"eni_e_ipv6":         "true",
											"slb_ipv6":           "true",
											"ecs_ipv6":           "false",
											"alb_ipv6":           "true",
											"ga_e_ipv6":          "true",
										},
									},
								},
							},
						},
					},
					"description":      "AddressBook create by Terraform",
					"tag_relation":     "and",
					"group_type":       "assetIpv6",
					"lang":             "zh",
					"auto_add_tag_ecs": "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":                    "yqc-test-01",
						"asset_member_uids.#":           "1",
						"asset_region_resource_types.#": "1",
						"description":                   "AddressBook create by Terraform",
						"tag_relation":                  "and",
						"group_type":                    "assetIpv6",
						"lang":                          "zh",
						"auto_add_tag_ecs":              "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": "test11",
					"asset_region_resource_types": []map[string]interface{}{
						{
							"asset_region_id": "all",
							"resource_type": []map[string]interface{}{
								{
									"ipv6": []map[string]interface{}{
										{
											"ai_gateway_e_ipv6":  "false",
											"api_gateway_e_ipv6": "false",
											"nlb_ipv6":           "false",
											"eni_e_ipv6":         "false",
											"slb_ipv6":           "false",
											"ecs_ipv6":           "true",
											"alb_ipv6":           "false",
											"ga_e_ipv6":          "true",
										},
									},
								},
							},
						},
					},
					"description": "222",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":                    "test11",
						"asset_region_resource_types.#": "1",
						"description":                   CHECKSET,
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
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallAddressBookMap12955 = map[string]string{
	"address_list_count": CHECKSET,
	"reference_count":    CHECKSET,
}

func AlicloudCloudFirewallAddressBookBasicDependence12955(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case resourceCase_20260624_R4KXek 12956
func TestAccAliCloudCloudFirewallAddressBook_basic12956(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallAddressBookMap12956)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallAddressBookBasicDependence12956)
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
					"group_name":  "IPv4AddressBook81",
					"description": "AddressBook create by Terraform",
					"group_type":  "ip",
					"lang":        "zh",
					"address_list": []string{
						"1.1.1.1/32", "2.2.2.2/32", "3.3.3.3/32", "4.4.4.4/32"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     "IPv4AddressBook81",
						"description":    "AddressBook create by Terraform",
						"group_type":     "ip",
						"lang":           "zh",
						"address_list.#": "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":  "IPv4AddressBookUpdate183",
					"description": "AddressBook update by Terraform",
					"address_list": []string{
						"1.1.1.1/32", "2.2.2.2/32"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     "IPv4AddressBookUpdate183",
						"description":    "AddressBook update by Terraform",
						"address_list.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": "IPv4AddressBookUpdate288",
					"address_list": []string{
						"1.1.1.1/32", "2.2.2.2/32"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     "IPv4AddressBookUpdate288",
						"address_list.#": "2",
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
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallAddressBookMap12956 = map[string]string{
	"address_list_count": CHECKSET,
	"reference_count":    CHECKSET,
}

func AlicloudCloudFirewallAddressBookBasicDependence12956(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case resourceCase_20260706_I4hrnB 12957
func TestAccAliCloudCloudFirewallAddressBook_basic12957(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallAddressBookMap12957)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallAddressBookBasicDependence12957)
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
					"group_name":  "IPv6AddressBook{{function.randomIntString(0,100)}}",
					"description": "AddressBook create by Terraform",
					"group_type":  "ipv6",
					"lang":        "zh",
					"address_list": []string{
						"::1/128", "::2/128", "::3/128", "::4/128"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     "IPv6AddressBook{{function.randomIntString(0,100)}}",
						"description":    "AddressBook create by Terraform",
						"group_type":     "ipv6",
						"lang":           "zh",
						"address_list.#": "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":  "IPv6AddressBookUpdate1{{function.randomIntString(0,100)}}",
					"description": "AddressBook update by Terraform",
					"address_list": []string{
						"1::/128", "2::/128"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     "IPv6AddressBookUpdate1{{function.randomIntString(0,100)}}",
						"description":    "AddressBook update by Terraform",
						"address_list.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":   "IPv6AddressBookUpdate2{{function.randomIntString(0,100)}}",
					"address_list": []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     "IPv6AddressBookUpdate2{{function.randomIntString(0,100)}}",
						"address_list.#": "0",
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
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallAddressBookMap12957 = map[string]string{
	"address_list_count": CHECKSET,
	"reference_count":    CHECKSET,
}

func AlicloudCloudFirewallAddressBookBasicDependence12957(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case resourceCase_20260706_hEWwyB 12958
func TestAccAliCloudCloudFirewallAddressBook_basic12958(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallAddressBookMap12958)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallAddressBookBasicDependence12958)
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
					"group_name":       "ECSTagAddressBook29",
					"description":      "AddressBook create by Terraform",
					"tag_relation":     "and",
					"group_type":       "tag",
					"lang":             "zh",
					"address_list":     []string{},
					"auto_add_tag_ecs": "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":       "ECSTagAddressBook29",
						"description":      "AddressBook create by Terraform",
						"tag_relation":     "and",
						"group_type":       "tag",
						"lang":             "zh",
						"auto_add_tag_ecs": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":       "ECSTagAddressBookUpdate115",
					"description":      "AddressBook update by Terraform",
					"tag_relation":     "or",
					"auto_add_tag_ecs": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":       "ECSTagAddressBookUpdate115",
						"description":      "AddressBook update by Terraform",
						"tag_relation":     "or",
						"auto_add_tag_ecs": "1",
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
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallAddressBookMap12958 = map[string]string{
	"address_list_count": CHECKSET,
	"reference_count":    CHECKSET,
}

func AlicloudCloudFirewallAddressBookBasicDependence12958(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case resourceCase_20260624_R4KXek_clone_0 12959
func TestAccAliCloudCloudFirewallAddressBook_basic12959(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallAddressBookMap12959)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallAddressBookBasicDependence12959)
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
					"group_name": "yqc-test-00",
					"asset_region_resource_types": []map[string]interface{}{
						{
							"asset_region_id": "all",
							"resource_type": []map[string]interface{}{
								{
									"ipv4": []map[string]interface{}{
										{
											"slb_e_ip":                "true",
											"ai_gateway_e_ip":         "true",
											"ga_e_ip":                 "false",
											"nat_public_ip":           "true",
											"bastion_host_ingress_ip": "true",
											"eip":                     "true",
											"nat_e_ip":                "true",
											"slb_public_ip":           "true",
											"nlb_e_ip":                "true",
											"hav_ip":                  "true",
											"ecs_e_ip":                "false",
											"eni_e_ip":                "true",
											"api_gateway_e_ip":        "true",
											"ecs_public_ip":           "true",
											"bastion_host_ip":         "true",
											"bastion_host_egress_ip":  "true",
											"alb_e_ip":                "true",
										},
									},
								},
							},
						},
					},
					"description": "AddressBook create by Terraform",
					"group_type":  "asset",
					"lang":        "zh",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":                    "yqc-test-00",
						"asset_region_resource_types.#": "1",
						"description":                   "AddressBook create by Terraform",
						"group_type":                    "asset",
						"lang":                          "zh",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"asset_region_resource_types": []map[string]interface{}{
						{
							"asset_region_id": "all",
							"resource_type": []map[string]interface{}{
								{
									"ipv4": []map[string]interface{}{
										{
											"ai_gateway_e_ip":         "false",
											"slb_e_ip":                "false",
											"ga_e_ip":                 "false",
											"nat_public_ip":           "false",
											"bastion_host_ingress_ip": "false",
											"eip":                     "false",
											"nat_e_ip":                "false",
											"slb_public_ip":           "false",
											"ecs_e_ip":                "false",
											"hav_ip":                  "false",
											"nlb_e_ip":                "true",
											"eni_e_ip":                "false",
											"api_gateway_e_ip":        "false",
											"ecs_public_ip":           "false",
											"bastion_host_ip":         "false",
											"bastion_host_egress_ip":  "false",
											"alb_e_ip":                "false",
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"asset_region_resource_types.#": "1",
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
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallAddressBookMap12959 = map[string]string{
	"address_list_count": CHECKSET,
	"reference_count":    CHECKSET,
}

func AlicloudCloudFirewallAddressBookBasicDependence12959(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case IPv4地址簿 5031
func TestAccAliCloudCloudFirewallAddressBook_basic5031(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallAddressBookMap5031)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallAddressBookBasicDependence5031)
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
					"group_name":  "IPv4AddressBook352",
					"description": "AddressBook create by Terraform",
					"lang":        "zh",
					"address_list": []string{
						"1.1.1.1/32", "2.2.2.2/32", "3.3.3.3/32", "4.4.4.4/32"},
					"group_type": "ip",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     CHECKSET,
						"description":    "AddressBook create by Terraform",
						"lang":           "zh",
						"address_list.#": "4",
						"group_type":     "ip",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":  "IPv4AddressBookUpdate1342",
					"description": "AddressBook update by Terraform",
					"address_list": []string{
						"1.1.1.1/32", "2.2.2.2/32"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     CHECKSET,
						"description":    "AddressBook update by Terraform",
						"address_list.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":   "IPv4AddressBookUpdate2675",
					"address_list": []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     CHECKSET,
						"address_list.#": "0",
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
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallAddressBookMap5031 = map[string]string{
	"address_list_count": CHECKSET,
	"reference_count":    CHECKSET,
}

func AlicloudCloudFirewallAddressBookBasicDependence5031(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case ECS Tag地址簿 5139
func TestAccAliCloudCloudFirewallAddressBook_basic5139(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallAddressBookMap5139)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallAddressBookBasicDependence5139)
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
					"group_name":       "ECSTagAddressBook430",
					"description":      "AddressBook create by Terraform",
					"lang":             "zh",
					"address_list":     []string{},
					"group_type":       "tag",
					"auto_add_tag_ecs": "0",
					"tag_relation":     "and",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":       CHECKSET,
						"description":      "AddressBook create by Terraform",
						"lang":             "zh",
						"group_type":       "tag",
						"auto_add_tag_ecs": "0",
						"tag_relation":     "and",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":       "ECSTagAddressBookUpdate1815",
					"description":      "AddressBook update by Terraform",
					"auto_add_tag_ecs": "1",
					"tag_relation":     "or",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":       CHECKSET,
						"description":      "AddressBook update by Terraform",
						"auto_add_tag_ecs": "1",
						"tag_relation":     "or",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": "ECSTagAddressBookUpdate219",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallAddressBookMap5139 = map[string]string{
	"address_list_count": CHECKSET,
	"reference_count":    CHECKSET,
}

func AlicloudCloudFirewallAddressBookBasicDependence5139(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case IPv6地址簿 5038
func TestAccAliCloudCloudFirewallAddressBook_basic5038(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallAddressBookMap5038)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallAddressBookBasicDependence5038)
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
					"group_name":  "IPv6AddressBook436",
					"description": "AddressBook create by Terraform",
					"lang":        "zh",
					"address_list": []string{
						"::1/128", "::2/128", "::3/128", "::4/128"},
					"group_type": "ipv6",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     CHECKSET,
						"description":    "AddressBook create by Terraform",
						"lang":           "zh",
						"address_list.#": "4",
						"group_type":     "ipv6",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":  "IPv6AddressBookUpdate1452",
					"description": "AddressBook update by Terraform",
					"address_list": []string{
						"1::/128", "2::/128"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     CHECKSET,
						"description":    "AddressBook update by Terraform",
						"address_list.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":   "IPv6AddressBookUpdate2310",
					"address_list": []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     CHECKSET,
						"address_list.#": "0",
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
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallAddressBookMap5038 = map[string]string{
	"address_list_count": CHECKSET,
	"reference_count":    CHECKSET,
}

func AlicloudCloudFirewallAddressBookBasicDependence5038(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 域名地址簿 5155
func TestAccAliCloudCloudFirewallAddressBook_basic5155(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallAddressBookMap5155)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallAddressBookBasicDependence5155)
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
					"group_name":  "DomainAddressBook381",
					"description": "AddressBook create by Terraform",
					"lang":        "zh",
					"address_list": []string{
						"1.com", "2.com", "3.com", "4.com"},
					"group_type": "domain",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     CHECKSET,
						"description":    "AddressBook create by Terraform",
						"lang":           "zh",
						"address_list.#": "4",
						"group_type":     "domain",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":  "AddressBook1692",
					"description": "AddressBook update by Terraform",
					"address_list": []string{
						"1.com", "2.com"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     CHECKSET,
						"description":    "AddressBook update by Terraform",
						"address_list.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name":   "AddressBook211",
					"address_list": []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     CHECKSET,
						"address_list.#": "0",
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
				ImportStateVerifyIgnore: []string{"lang"},
			},
		},
	})
}

var AlicloudCloudFirewallAddressBookMap5155 = map[string]string{
	"address_list_count": CHECKSET,
	"reference_count":    CHECKSET,
}

func AlicloudCloudFirewallAddressBookBasicDependence5155(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test CloudFirewall AddressBook. <<< Resource test cases, automatically generated.

func TestAccAliCloudCloudFirewallAddressBook_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudfwService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
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
					"group_name":   name,
					"group_type":   "ip",
					"description":  name,
					"address_list": []string{"10.21.0.0/16", "10.22.0.0/16", "10.168.0.0/16"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     name,
						"group_type":     "ip",
						"description":    name,
						"address_list.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"lang": "en",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"lang": "en",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"address_list": []string{"10.21.0.0/16"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"address_list.#": "1",
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

func TestAccAliCloudCloudFirewallAddressBook_basic1(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudfwService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
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
					"group_name":   name,
					"group_type":   "ipv6",
					"description":  name,
					"address_list": []string{"::1/128", "::2/128", "::3/128"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     name,
						"group_type":     "ipv6",
						"description":    name,
						"address_list.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"lang": "en",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"lang": "en",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"address_list": []string{"::1/128"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"address_list.#": "1",
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

func TestAccAliCloudCloudFirewallAddressBook_basic2(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudfwService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
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
					"group_name":   name,
					"group_type":   "domain",
					"description":  name,
					"address_list": []string{"alibaba.com", "aliyun.com", "alicloud.com"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     name,
						"group_type":     "domain",
						"description":    name,
						"address_list.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"lang": "en",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"lang": "en",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"address_list": []string{"alibaba.com"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"address_list.#": "1",
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

func TestAccAliCloudCloudFirewallAddressBook_basic3(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudfwService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
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
					"group_name":   name,
					"group_type":   "port",
					"description":  name,
					"address_list": []string{"1/1", "22/22", "88/88"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":     name,
						"group_type":     "port",
						"description":    name,
						"address_list.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"lang": "en",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"lang": "en",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"address_list": []string{"1/1"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"address_list.#": "1",
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

func TestAccAliCloudCloudFirewallAddressBook_basic4(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudfwService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
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
					"group_name":  name,
					"group_type":  "tag",
					"description": name,
					"ecs_tags": []map[string]interface{}{
						{
							"tag_key":   "created",
							"tag_value": "tfTestAcc0",
						},
						{
							"tag_key":   "for",
							"tag_value": "tfTestAcc1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":  name,
						"group_type":  "tag",
						"description": name,
						"ecs_tags.#":  "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_add_tag_ecs": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_add_tag_ecs": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tag_relation": "or",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tag_relation": "or",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"lang": "en",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"lang": "en",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ecs_tags": []map[string]interface{}{
						{
							"tag_key":   "created",
							"tag_value": "tfTestAcc0",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ecs_tags.#": "1",
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

func TestAccAliCloudCloudFirewallAddressBook_basic5(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_address_book.default"
	checkoutSupportedRegions(t, true, connectivity.CloudFirewallSupportRegions)
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallAddressBookMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudfwService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallAddressBook")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudfirewalladdressbook%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallAddressBookBasicDependence0)
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
					"group_name":       name,
					"group_type":       "tag",
					"description":      name,
					"auto_add_tag_ecs": "1",
					"tag_relation":     "or",
					"lang":             "en",
					"ecs_tags": []map[string]interface{}{
						{
							"tag_key":   "created",
							"tag_value": "tfTestAcc0",
						},
						{
							"tag_key":   "for",
							"tag_value": "tfTestAcc1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":       name,
						"group_type":       "tag",
						"description":      name,
						"auto_add_tag_ecs": "1",
						"tag_relation":     "or",
						"lang":             "en",
						"ecs_tags.#":       "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_name": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_add_tag_ecs": "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_add_tag_ecs": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tag_relation": "and",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tag_relation": "and",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"lang": "zh",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"lang": "zh",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ecs_tags": []map[string]interface{}{
						{
							"tag_key":   "created",
							"tag_value": "tfTestAcc0",
						},
						{
							"tag_key":   "for",
							"tag_value": "tfTestAcc1",
						},
						{
							"tag_key":   "by",
							"tag_value": "tfTestAcc2",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ecs_tags.#": "3",
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

var AliCloudCloudFirewallAddressBookMap0 = map[string]string{}

func AliCloudCloudFirewallAddressBookBasicDependence0(name string) string {
	return fmt.Sprintf(` 
	variable "name" {
  		default = "%s"
	}
`, name)
}
