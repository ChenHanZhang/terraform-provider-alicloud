package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAliCloudKmsInstance_basic4048(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap4048)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%skmsinstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence4048)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.KmsInstanceSupportRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"product_version": "3",
					"vpc_id":          "${alicloud_vpc.default.id}",
					"log":             "0",
					"log_storage":     "0",
					"zone_ids": []string{
						"cn-hangzhou-k", "cn-hangzhou-j"},
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitch.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      "0",
						"spec":            "1000",
						"product_version": "3",
						"vpc_id":          CHECKSET,
						"zone_ids.#":      "2",
						"vswitch_ids.#":   "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num": "7",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num": "7",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bind_vpcs": []map[string]interface{}{
						{
							"vpc_id":       "${alicloud_vswitch.shareVswitch.vpc_id}",
							"region_id":    "cn-hangzhou",
							"vswitch_id":   "${alicloud_vswitch.shareVswitch.id}",
							"vpc_owner_id": "${data.alicloud_account.current.id}",
						},
						{
							"vpc_id":       "${alicloud_vswitch.share-vswitch2.vpc_id}",
							"region_id":    "cn-hangzhou",
							"vswitch_id":   "${alicloud_vswitch.share-vswitch2.id}",
							"vpc_owner_id": "${data.alicloud_account.current.id}",
						},
						{
							"vpc_id":       "${alicloud_vswitch.share-vsw3.vpc_id}",
							"region_id":    "cn-hangzhou",
							"vswitch_id":   "${alicloud_vswitch.share-vsw3.id}",
							"vpc_owner_id": "${data.alicloud_account.current.id}",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bind_vpcs.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"key_num": "1000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"key_num": "1000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"key_num": "2000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"key_num": "2000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"spec": "1000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"spec": "1000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"spec": "2000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"spec": "2000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"secret_num": "1000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"secret_num": "1000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"log":         "1",
					"log_storage": "1000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"log":         "1",
						"log_storage": "1000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bind_vpcs": []map[string]interface{}{
						{
							"vpc_id":       "${alicloud_vswitch.share-vswitch2.vpc_id}",
							"region_id":    "cn-hangzhou",
							"vswitch_id":   "${alicloud_vswitch.share-vswitch2.id}",
							"vpc_owner_id": "${data.alicloud_account.current.id}",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bind_vpcs.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bind_vpcs": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bind_vpcs.#": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "5",
					"key_num":         "2000",
					"secret_num":      "2000",
					"spec":            "2000",
					"renew_status":    "ManualRenewal",
					"product_version": "3",
					"vpc_id":          "${alicloud_vpc.default.id}",
					"zone_ids": []string{
						"cn-hangzhou-k", "cn-hangzhou-j"},
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitch.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "5",
						"key_num":         "2000",
						"secret_num":      "2000",
						"spec":            "2000",
						"renew_status":    "ManualRenewal",
						"product_version": "3",
						"vpc_id":          CHECKSET,
						"zone_ids.#":      "2",
						"vswitch_ids.#":   "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renew_status":        "AutoRenewal",
					"renew_period":        "1",
					"renewal_period_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renew_status":        "AutoRenewal",
						"renew_period":        "1",
						"renewal_period_unit": "M",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap4048 = map[string]string{
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
	"instance_name":            CHECKSET,
	"ca_certificate_chain_pem": CHECKSET,
}

func AlicloudKmsInstanceBasicDependence4048(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_account" "current" {}
data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "default" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = var.name
}

resource "alicloud_vswitch" "vswitch" {
  vpc_id     = alicloud_vpc.default.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vswitch" "vswitch-j" {
  vpc_id     = alicloud_vpc.default.id
  zone_id    = "cn-hangzhou-j"
  cidr_block = "172.16.2.0/24"
}

resource "alicloud_vpc" "shareVPC" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "${var.name}3"
}

resource "alicloud_vswitch" "shareVswitch" {
  vpc_id     = alicloud_vpc.shareVPC.id
  zone_id    = data.alicloud_zones.default.zones.1.id
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vpc" "share-VPC2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "${var.name}5"
}

resource "alicloud_vswitch" "share-vswitch2" {
  vpc_id     = alicloud_vpc.share-VPC2.id
  zone_id    = data.alicloud_zones.default.zones.1.id
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vpc" "share-VPC3" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "${var.name}7"
}

resource "alicloud_vswitch" "share-vsw3" {
  vpc_id     = alicloud_vpc.share-VPC3.id
  zone_id    = data.alicloud_zones.default.zones.1.id
  cidr_block = "172.16.1.0/24"
}


`, name)
}

func AlicloudKmsInstanceBasicDependence4048_intl(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_account" "current" {}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "default" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = var.name
}

resource "alicloud_vswitch" "vswitch" {
  vpc_id     = alicloud_vpc.default.id
  zone_id    = "ap-southeast-1a"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vswitch" "vswitch-j" {
  vpc_id     = alicloud_vpc.default.id
  zone_id    = "ap-southeast-1b"
  cidr_block = "172.16.2.0/24"
}

resource "alicloud_vpc" "shareVPC" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "${var.name}3"
}

resource "alicloud_vswitch" "shareVswitch" {
  vpc_id     = alicloud_vpc.shareVPC.id
  zone_id    = data.alicloud_zones.default.zones.1.id
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vpc" "share-VPC2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "${var.name}5"
}

resource "alicloud_vswitch" "share-vswitch2" {
  vpc_id     = alicloud_vpc.share-VPC2.id
  zone_id    = data.alicloud_zones.default.zones.1.id
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vpc" "share-VPC3" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "${var.name}7"
}

resource "alicloud_vswitch" "share-vsw3" {
  vpc_id     = alicloud_vpc.share-VPC3.id
  zone_id    = data.alicloud_zones.default.zones.1.id
  cidr_block = "172.16.1.0/24"
}


`, name)
}

// Case 4048  twin
func TestAccAliCloudKmsInstance_basic4048_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap4048)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%skmsinstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence4048)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithAccountSiteType(t, DomesticSite)
			testAccPreCheckWithRegions(t, true, connectivity.KmsInstanceSupportRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_name":   name,
					"vpc_num":         "7",
					"key_num":         "2000",
					"secret_num":      "1000",
					"spec":            "2000",
					"renew_status":    "ManualRenewal",
					"product_version": "3",
					"log":             "1",
					"log_storage":     "1000",
					"period":          "2",
					"vpc_id":          "${alicloud_vpc.default.id}",
					"zone_ids": []string{
						"cn-hangzhou-k", "cn-hangzhou-j"},
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitch.id}"},
					"bind_vpcs": []map[string]interface{}{
						{
							"vpc_id":       "${alicloud_vpc.shareVPC.id}",
							"region_id":    "cn-hangzhou",
							"vswitch_id":   "${alicloud_vswitch.shareVswitch.id}",
							"vpc_owner_id": "${data.alicloud_account.current.id}",
						},
						{
							"vpc_id":       "${alicloud_vswitch.share-vsw3.vpc_id}",
							"region_id":    "cn-hangzhou",
							"vswitch_id":   "${alicloud_vswitch.share-vsw3.id}",
							"vpc_owner_id": "${data.alicloud_account.current.id}",
						},
					},
					"tags": map[string]string{
						"Created": "TF",
						"For":     "Test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":   name,
						"vpc_num":         "7",
						"key_num":         "2000",
						"secret_num":      "1000",
						"spec":            "2000",
						"renew_status":    "ManualRenewal",
						"product_version": "3",
						"vpc_id":          CHECKSET,
						"zone_ids.#":      "2",
						"vswitch_ids.#":   "1",
						"bind_vpcs.#":     "2",
						"tags.%":          "2",
						"tags.Created":    "TF",
						"tags.For":        "Test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bind_vpcs": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bind_vpcs.#": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name": name + "_update",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

func TestAccAliCloudKmsInstance_basic4048_postpaid(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap4048)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%skmsinstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence4048)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithAccountSiteType(t, DomesticSite)
			testAccPreCheckWithRegions(t, true, connectivity.KmsInstanceSupportRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"payment_type":    "PayAsYouGo",
					"product_version": "3",
					"vpc_id":          "${alicloud_vpc.default.id}",
					"zone_ids": []string{
						"cn-hangzhou-k", "cn-hangzhou-j"},
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitch-j.id}"},
					"bind_vpcs": []map[string]interface{}{
						{
							"vpc_id":       "${alicloud_vpc.shareVPC.id}",
							"region_id":    "cn-hangzhou",
							"vswitch_id":   "${alicloud_vswitch.shareVswitch.id}",
							"vpc_owner_id": "${data.alicloud_account.current.id}",
						},
						{
							"vpc_id":       "${alicloud_vswitch.share-vsw3.vpc_id}",
							"region_id":    "cn-hangzhou",
							"vswitch_id":   "${alicloud_vswitch.share-vsw3.id}",
							"vpc_owner_id": "${data.alicloud_account.current.id}",
						},
					},
					"force_delete_without_backup": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_version": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bind_vpcs": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bind_vpcs.#": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bind_vpcs": []map[string]interface{}{
						{
							"vpc_id":       "${alicloud_vswitch.shareVswitch.vpc_id}",
							"region_id":    "cn-hangzhou",
							"vswitch_id":   "${alicloud_vswitch.shareVswitch.id}",
							"vpc_owner_id": "${data.alicloud_account.current.id}",
						},
						{
							"vpc_id":       "${alicloud_vswitch.share-vswitch2.vpc_id}",
							"region_id":    "cn-hangzhou",
							"vswitch_id":   "${alicloud_vswitch.share-vswitch2.id}",
							"vpc_owner_id": "${data.alicloud_account.current.id}",
						},
						{
							"vpc_id":       "${alicloud_vswitch.share-vsw3.vpc_id}",
							"region_id":    "cn-hangzhou",
							"vswitch_id":   "${alicloud_vswitch.share-vsw3.id}",
							"vpc_owner_id": "${data.alicloud_account.current.id}",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bind_vpcs.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bind_vpcs": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bind_vpcs.#": "0",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"period", "force_delete_without_backup", "renewal_period_unit"},
			},
		},
	})
}

func TestAccAliCloudKmsInstance_basic4048_postpaid_intl(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap4048)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%skmsinstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence5405)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.KmsInstanceIntlSupportRegions)
			testAccPreCheckWithAccountSiteType(t, IntlSite)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"payment_type":    "PayAsYouGo",
					"product_version": "3",
					"vpc_id":          "${alicloud_vswitch.vswitch.vpc_id}",
					"zone_ids": []string{
						"${alicloud_vswitch.vswitch.zone_id}", "${alicloud_vswitch.vswitch-j.zone_id}"},
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitch.id}"},
					"force_delete_without_backup": "true",
					"tags": map[string]string{
						"Created": "TF",
						"For":     "Test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_version": "3",
						"tags.%":          "2",
						"tags.Created":    "TF",
						"tags.For":        "Test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bind_vpcs": []map[string]interface{}{
						{
							"vpc_id":       "${alicloud_vswitch.shareVswitch.vpc_id}",
							"region_id":    "ap-southeast-1",
							"vswitch_id":   "${alicloud_vswitch.shareVswitch.id}",
							"vpc_owner_id": "${data.alicloud_account.current.id}",
						},
						{
							"vpc_id":       "${alicloud_vswitch.share-vswitch2.vpc_id}",
							"region_id":    "ap-southeast-1",
							"vswitch_id":   "${alicloud_vswitch.share-vswitch2.id}",
							"vpc_owner_id": "${data.alicloud_account.current.id}",
						},
						{
							"vpc_id":       "${alicloud_vswitch.share-vsw3.vpc_id}",
							"region_id":    "ap-southeast-1",
							"vswitch_id":   "${alicloud_vswitch.share-vsw3.id}",
							"vpc_owner_id": "${data.alicloud_account.current.id}",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bind_vpcs.#": "3",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"period", "force_delete_without_backup", "renewal_period_unit"},
			},
		},
	})
}

func TestAccAliCloudKmsInstance_basic4048_intl(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap4048)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%skmsinstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence4048_intl)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.KmsInstanceIntlSupportRegions)
			testAccPreCheckWithAccountSiteType(t, IntlSite)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "2",
					"key_num":         "1000",
					"secret_num":      "1000",
					"spec":            "1000",
					"renew_status":    "ManualRenewal",
					"product_version": "3",
					"vpc_id":          "${alicloud_vpc.default.id}",
					"zone_ids": []string{
						"ap-southeast-1a", "ap-southeast-1b"},
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitch.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "2",
						"key_num":         "1000",
						"secret_num":      "1000",
						"spec":            "1000",
						"renew_status":    "ManualRenewal",
						"product_version": "3",
						"vpc_id":          CHECKSET,
						"zone_ids.#":      "2",
						"vswitch_ids.#":   "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num": "7",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num": "7",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"key_num": "2000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"key_num": "2000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"spec": "2000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"spec": "2000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"secret_num": "2000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"secret_num": "2000",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

// Test Kms Instance. >>> Resource test cases, automatically generated.
// Case 国内站_河源线上测试_Period=36 11024
func TestAccAliCloudKmsInstance_basic11024(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap11024)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence11024)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "36",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "36",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap11024 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence11024(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站_河源线上测试_Period=24 11023
func TestAccAliCloudKmsInstance_basic11023(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap11023)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence11023)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "24",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "24",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap11023 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence11023(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站_河源线上测试_Period=12 11022
func TestAccAliCloudKmsInstance_basic11022(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap11022)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence11022)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "12",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "12",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap11022 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence11022(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 小规格实例升配标准规格_杭州线上 10993
func TestAccAliCloudKmsInstance_basic10993(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap10993)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence10993)
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
					"vpc_num":         "1",
					"key_num":         "100",
					"secret_num":      "100",
					"spec":            "200",
					"product_version": "5",
					"vpc_id":          "${alicloud_vpc.pre_amp_test.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitch.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitch.zone_id}", "${alicloud_vswitch.vswitch-j.zone_id}"},
					"log":          "0",
					"period":       "1",
					"log_storage":  "0",
					"payment_type": "Subscription",
					"renew_status": "AutoRenewal",
					"renew_period": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "100",
						"secret_num":      "100",
						"spec":            "200",
						"product_version": CHECKSET,
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"log":             CHECKSET,
						"period":          "1",
						"log_storage":     "0",
						"payment_type":    "Subscription",
						"renew_status":    "AutoRenewal",
						"renew_period":    "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"key_num":             "1000",
					"spec":                "1000",
					"product_version":     "3",
					"renewal_period_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"key_num":             "1000",
						"spec":                "1000",
						"product_version":     CHECKSET,
						"renewal_period_unit": "M",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap10993 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence10993(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "pre_amp_test" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vswitch" {
  vpc_id     = alicloud_vpc.pre_amp_test.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vswitch" "vswitch-j" {
  vpc_id     = alicloud_vpc.pre_amp_test.id
  zone_id    = "cn-hangzhou-j"
  cidr_block = "172.16.2.0/24"
}


`, name)
}

// Case 国内站_河源线上测试_Period=6 11021
func TestAccAliCloudKmsInstance_basic11021(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap11021)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence11021)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "6",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "6",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap11021 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence11021(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站_河源线上测试_Period=3 11020
func TestAccAliCloudKmsInstance_basic11020(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap11020)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence11020)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "3",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "3",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap11020 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence11020(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站_河源线上测试_Period=2 11019
func TestAccAliCloudKmsInstance_basic11019(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap11019)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence11019)
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
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":        "2",
					"log_storage":   "0",
					"log":           "0",
					"payment_type":  "Subscription",
					"instance_name": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "2",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
						"instance_name":   name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name": name + "_update",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap11019 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence11019(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test2.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.10.0/24"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block = "172.16.11.0/24"
  zone_id    = "cn-hangzhou-h"
}


`, name)
}

// Case 国内站_河源线上测试_Period=1——包含更新值测试 11013
func TestAccAliCloudKmsInstance_basic11013(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap11013)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence11013)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":        "1",
					"log_storage":   "0",
					"log":           "0",
					"payment_type":  "Subscription",
					"instance_name": name,
					"renew_period":  "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "1",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
						"instance_name":   name,
						"renew_period":    "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":             "5",
					"key_num":             "2000",
					"secret_num":          "2000",
					"spec":                "2000",
					"log_storage":         "1000",
					"log":                 "1",
					"instance_name":       name + "_update",
					"renew_period":        "2",
					"renewal_period_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":             "5",
						"key_num":             "2000",
						"secret_num":          "2000",
						"spec":                "2000",
						"log_storage":         "1000",
						"log":                 CHECKSET,
						"instance_name":       name + "_update",
						"renew_period":        "2",
						"renewal_period_unit": "M",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap11013 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence11013(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 手动续费升级自动续费_杭州线上 10991
func TestAccAliCloudKmsInstance_basic10991(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap10991)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence10991)
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
					"vpc_num":         "1",
					"key_num":         "100",
					"secret_num":      "100",
					"spec":            "200",
					"renew_status":    "ManualRenewal",
					"product_version": "5",
					"vpc_id":          "${alicloud_vpc.pre_amp_test.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitch.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitch.zone_id}", "${alicloud_vswitch.vswitch-j.zone_id}"},
					"log":          "0",
					"period":       "1",
					"log_storage":  "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "100",
						"secret_num":      "100",
						"spec":            "200",
						"renew_status":    "ManualRenewal",
						"product_version": CHECKSET,
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"log":             CHECKSET,
						"period":          "1",
						"log_storage":     "0",
						"payment_type":    "Subscription",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renew_status":        "AutoRenewal",
					"renew_period":        "1",
					"renewal_period_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renew_status":        "AutoRenewal",
						"renew_period":        "1",
						"renewal_period_unit": "M",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap10991 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence10991(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "pre_amp_test" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vswitch" {
  vpc_id     = alicloud_vpc.pre_amp_test.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vswitch" "vswitch-j" {
  vpc_id     = alicloud_vpc.pre_amp_test.id
  zone_id    = "cn-hangzhou-j"
  cidr_block = "172.16.2.0/24"
}


`, name)
}

// Case 国内站软件版多VPC 11014
func TestAccAliCloudKmsInstance_basic11014(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap11014)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence11014)
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
					"vpc_num":         "7",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "ManualRenewal",
					"product_version": "3",
					"vpc_id":          "${alicloud_vswitch.vswitch.vpc_id}",
					"zone_ids": []string{
						"cn-hangzhou-k", "cn-hangzhou-j"},
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitch.id}"},
					"bind_vpcs": []map[string]interface{}{
						{
							"vpc_id":       "${alicloud_vswitch.shareVswitch.vpc_id}",
							"region_id":    "${alicloud_vswitch.shareVswitch.region_id}",
							"vswitch_id":   "${alicloud_vswitch.shareVswitch.id}",
							"vpc_owner_id": "1511928242963727",
						},
						{
							"vpc_id":       "${alicloud_vswitch.share-vswitch2.vpc_id}",
							"region_id":    "${alicloud_vswitch.share-vswitch2.region_id}",
							"vswitch_id":   "${alicloud_vswitch.share-vswitch2.id}",
							"vpc_owner_id": "1511928242963727",
						},
						{
							"vpc_id":       "${alicloud_vswitch.share-vsw3.vpc_id}",
							"region_id":    "${alicloud_vswitch.share-vsw3.region_id}",
							"vswitch_id":   "${alicloud_vswitch.share-vsw3.id}",
							"vpc_owner_id": "1511928242963727",
						},
					},
					"log":          "0",
					"period":       "1",
					"log_storage":  "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "7",
						"key_num":         CHECKSET,
						"secret_num":      CHECKSET,
						"spec":            CHECKSET,
						"renew_status":    "ManualRenewal",
						"product_version": CHECKSET,
						"vpc_id":          CHECKSET,
						"zone_ids.#":      "2",
						"vswitch_ids.#":   "1",
						"bind_vpcs.#":     "3",
						"log":             CHECKSET,
						"period":          "1",
						"log_storage":     "0",
						"payment_type":    "Subscription",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"spec":         "1000",
					"renew_status": "AutoRenewal",
					"bind_vpcs": []map[string]interface{}{
						{
							"vpc_id":       "${alicloud_vpc.shareVPC.id}",
							"region_id":    "${alicloud_vpc.shareVPC.region_id}",
							"vswitch_id":   "${alicloud_vswitch.shareVswitch.id}",
							"vpc_owner_id": "1511928242963727",
						},
						{
							"vpc_id":       "vpc-bp14c07ucxg6h1xjmgcld",
							"region_id":    "cn-hangzhou",
							"vswitch_id":   "vsw-bp1wujtnspi1l3gvunvds",
							"vpc_owner_id": "1192853035118460",
						},
					},
					"renew_period":        "1",
					"renewal_period_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"spec":                "1000",
						"renew_status":        "AutoRenewal",
						"bind_vpcs.#":         "2",
						"renew_period":        "1",
						"renewal_period_unit": "M",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bind_vpcs": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bind_vpcs.#": "0",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap11014 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence11014(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test"
}

resource "alicloud_vswitch" "vswitch" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vswitch" "vswitch-j" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test.id
  zone_id    = "cn-hangzhou-j"
  cidr_block = "172.16.2.0/24"
}

resource "alicloud_vpc" "shareVPC" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-share-vpc-test1"
}

resource "alicloud_vswitch" "shareVswitch" {
  vpc_id     = alicloud_vpc.shareVPC.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vpc" "share-VPC2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-share-vpc-test2"
}

resource "alicloud_vswitch" "share-vswitch2" {
  vpc_id     = alicloud_vpc.share-VPC2.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vpc" "share-VPC3" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-share-vpc-test3"
}

resource "alicloud_vswitch" "share-vsw3" {
  vpc_id     = alicloud_vpc.share-VPC3.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}


`, name)
}

// Case 国内站杭州后付费 11015
func TestAccAliCloudKmsInstance_basic11015(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap11015)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence11015)
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
					"vpc_id": "${alicloud_vpc.vpc-amp-instance-test.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitch.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitch.zone_id}", "${alicloud_vswitch.vswitch-j.zone_id}"},
					"log_storage":     "0",
					"payment_type":    "PayAsYouGo",
					"product_version": "3",
					"log":             "0",
					"spec":            "1000",
					"period":          "1",
					"renew_status":    "ManualRenewal",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"log_storage":     "0",
						"payment_type":    "PayAsYouGo",
						"product_version": CHECKSET,
						"log":             CHECKSET,
						"spec":            "1000",
						"period":          "1",
						"renew_status":    "ManualRenewal",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap11015 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence11015(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test"
}

resource "alicloud_vswitch" "vswitch" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vswitch" "vswitch-j" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test.id
  zone_id    = "cn-hangzhou-j"
  cidr_block = "172.16.2.0/24"
}


`, name)
}

// Case 国内站_河源线上测试_Period=1——包含更新值测试_副本1744168044666 10662
func TestAccAliCloudKmsInstance_basic10662(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap10662)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence10662)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "ManualRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":        "1",
					"log_storage":   "0",
					"log":           "0",
					"payment_type":  "Subscription",
					"instance_name": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "ManualRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "1",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
						"instance_name":   name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":       "5",
					"key_num":       "2000",
					"secret_num":    "2000",
					"spec":          "2000",
					"log_storage":   "1000",
					"log":           "1",
					"instance_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":       "5",
						"key_num":       "2000",
						"secret_num":    "2000",
						"spec":          "2000",
						"log_storage":   "1000",
						"log":           CHECKSET,
						"instance_name": name + "_update",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap10662 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence10662(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站_河源线上测试_Period=12_副本1744194246900 10671
func TestAccAliCloudKmsInstance_basic10671(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap10671)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence10671)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "12",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "12",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap10671 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence10671(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站_河源线上测试_Period=36_副本1744194273259 10673
func TestAccAliCloudKmsInstance_basic10673(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap10673)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence10673)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "36",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "36",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap10673 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence10673(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站小规格_杭州_副本1744166796143 10660
func TestAccAliCloudKmsInstance_basic10660(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap10660)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence10660)
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
					"vpc_num":         "1",
					"key_num":         "100",
					"secret_num":      "100",
					"spec":            "200",
					"renew_status":    "AutoRenewal",
					"product_version": "5",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitch.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitch.zone_id}", "${alicloud_vswitch.vswitch-j.zone_id}"},
					"log":          "0",
					"period":       "1",
					"log_storage":  "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "100",
						"secret_num":      "100",
						"spec":            "200",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"log":             CHECKSET,
						"period":          "1",
						"log_storage":     "0",
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap10660 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence10660(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test"
}

resource "alicloud_vswitch" "vswitch" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vswitch" "vswitch-j" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test.id
  zone_id    = "cn-hangzhou-j"
  cidr_block = "172.16.2.0/24"
}


`, name)
}

// Case 国内站_河源线上测试_Period=24_副本1744194259946 10672
func TestAccAliCloudKmsInstance_basic10672(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap10672)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence10672)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "24",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "24",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap10672 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence10672(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站_河源线上测试_Period=6_副本1744189945179 10670
func TestAccAliCloudKmsInstance_basic10670(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap10670)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence10670)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "6",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "6",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap10670 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence10670(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站软件版多VPC 4048
func TestAccAliCloudKmsInstance_basic4048(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap4048)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence4048)
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
					"vpc_num":         "7",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "ManualRenewal",
					"product_version": "3",
					"renew_period":    "3",
					"vpc_id":          "${alicloud_vswitch.vswitch.vpc_id}",
					"zone_ids": []string{
						"cn-hangzhou-k", "cn-hangzhou-j"},
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitch.id}"},
					"bind_vpcs": []map[string]interface{}{
						{
							"vpc_id":       "${alicloud_vswitch.shareVswitch.vpc_id}",
							"region_id":    "${alicloud_vswitch.shareVswitch.region_id}",
							"vswitch_id":   "${alicloud_vswitch.shareVswitch.id}",
							"vpc_owner_id": "1511928242963727",
						},
						{
							"vpc_id":       "${alicloud_vswitch.share-vswitch2.vpc_id}",
							"region_id":    "${alicloud_vswitch.share-vswitch2.region_id}",
							"vswitch_id":   "${alicloud_vswitch.share-vswitch2.id}",
							"vpc_owner_id": "1511928242963727",
						},
						{
							"vpc_id":       "${alicloud_vswitch.share-vsw3.vpc_id}",
							"region_id":    "${alicloud_vswitch.share-vsw3.region_id}",
							"vswitch_id":   "${alicloud_vswitch.share-vsw3.id}",
							"vpc_owner_id": "1511928242963727",
						},
					},
					"log":          "0",
					"period":       "1",
					"log_storage":  "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "7",
						"key_num":         CHECKSET,
						"secret_num":      CHECKSET,
						"spec":            CHECKSET,
						"renew_status":    "ManualRenewal",
						"product_version": CHECKSET,
						"renew_period":    CHECKSET,
						"vpc_id":          CHECKSET,
						"zone_ids.#":      "2",
						"vswitch_ids.#":   "1",
						"bind_vpcs.#":     "3",
						"log":             CHECKSET,
						"period":          "1",
						"log_storage":     "0",
						"payment_type":    "Subscription",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"spec": "1000",
					"bind_vpcs": []map[string]interface{}{
						{
							"vpc_id":       "${alicloud_vpc.shareVPC.id}",
							"region_id":    "${alicloud_vpc.shareVPC.region_id}",
							"vswitch_id":   "${alicloud_vswitch.shareVswitch.id}",
							"vpc_owner_id": "1511928242963727",
						},
						{
							"vpc_id":       "vpc-bp14c07ucxg6h1xjmgcld",
							"region_id":    "cn-hangzhou",
							"vswitch_id":   "vsw-bp1wujtnspi1l3gvunvds",
							"vpc_owner_id": "1192853035118460",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"spec":        "1000",
						"bind_vpcs.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bind_vpcs": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"bind_vpcs.#": "0",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap4048 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence4048(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test"
}

resource "alicloud_vswitch" "vswitch" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vswitch" "vswitch-j" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test.id
  zone_id    = "cn-hangzhou-j"
  cidr_block = "172.16.2.0/24"
}

resource "alicloud_vpc" "shareVPC" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-share-vpc-test1"
}

resource "alicloud_vswitch" "shareVswitch" {
  vpc_id     = alicloud_vpc.shareVPC.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vpc" "share-VPC2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-share-vpc-test2"
}

resource "alicloud_vswitch" "share-vswitch2" {
  vpc_id     = alicloud_vpc.share-VPC2.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vpc" "share-VPC3" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-share-vpc-test3"
}

resource "alicloud_vswitch" "share-vsw3" {
  vpc_id     = alicloud_vpc.share-VPC3.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}


`, name)
}

// Case 国内站杭州后付费 10122
func TestAccAliCloudKmsInstance_basic10122(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap10122)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence10122)
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
					"renew_period": "1",
					"vpc_id":       "${alicloud_vpc.vpc-amp-instance-test.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitch.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitch.zone_id}", "${alicloud_vswitch.vswitch-j.zone_id}"},
					"log_storage":     "0",
					"payment_type":    "PayAsYouGo",
					"product_version": "3",
					"log":             "0",
					"spec":            "1000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"log_storage":     "0",
						"payment_type":    "PayAsYouGo",
						"product_version": CHECKSET,
						"log":             CHECKSET,
						"spec":            "1000",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap10122 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence10122(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test"
}

resource "alicloud_vswitch" "vswitch" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vswitch" "vswitch-j" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test.id
  zone_id    = "cn-hangzhou-j"
  cidr_block = "172.16.2.0/24"
}


`, name)
}

// Case 国内站_河源线上测试_Period=3_副本1744188004209 10669
func TestAccAliCloudKmsInstance_basic10669(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap10669)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence10669)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "3",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "3",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap10669 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence10669(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站_河源线上测试_Period=2_副本1741161153449 10474
func TestAccAliCloudKmsInstance_basic10474(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap10474)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence10474)
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
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":        "2",
					"log_storage":   "0",
					"log":           "0",
					"payment_type":  "Subscription",
					"instance_name": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "2",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
						"instance_name":   name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name": name + "_update",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap10474 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence10474(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test2.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.10.0/24"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block = "172.16.11.0/24"
  zone_id    = "cn-hangzhou-h"
}


`, name)
}

// Case 国内站_河源线上测试_Period=24_副本1744194259946_副本1744680811277 10694
func TestAccAliCloudKmsInstance_basic10694(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap10694)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence10694)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "24",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "24",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap10694 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence10694(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站小规格_杭州 5237
func TestAccAliCloudKmsInstance_basic5237(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap5237)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence5237)
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
					"vpc_num":         "1",
					"key_num":         "100",
					"secret_num":      "100",
					"spec":            "200",
					"renew_status":    "AutoRenewal",
					"product_version": "5",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitch.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitch.zone_id}", "${alicloud_vswitch.vswitch-j.zone_id}"},
					"log":          "0",
					"period":       "1",
					"log_storage":  "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "100",
						"secret_num":      "100",
						"spec":            "200",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"log":             CHECKSET,
						"period":          "1",
						"log_storage":     "0",
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap5237 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence5237(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test"
}

resource "alicloud_vswitch" "vswitch" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test.id
  zone_id    = "cn-hangzhou-k"
  cidr_block = "172.16.1.0/24"
}

resource "alicloud_vswitch" "vswitch-j" {
  vpc_id     = alicloud_vpc.vpc-amp-instance-test.id
  zone_id    = "cn-hangzhou-j"
  cidr_block = "172.16.2.0/24"
}


`, name)
}

// Case 国内站_河源线上测试_Period=1——包含更新值测试 9704
func TestAccAliCloudKmsInstance_basic9704(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap9704)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence9704)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "ManualRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":        "1",
					"log_storage":   "0",
					"log":           "0",
					"payment_type":  "Subscription",
					"instance_name": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "ManualRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "1",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
						"instance_name":   name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":       "5",
					"key_num":       "2000",
					"secret_num":    "2000",
					"spec":          "2000",
					"log_storage":   "1000",
					"log":           "1",
					"instance_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":       "5",
						"key_num":       "2000",
						"secret_num":    "2000",
						"spec":          "2000",
						"log_storage":   "1000",
						"log":           CHECKSET,
						"instance_name": name + "_update",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap9704 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence9704(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站_河源线上测试_Period=12 8472
func TestAccAliCloudKmsInstance_basic8472(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap8472)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence8472)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "12",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "12",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap8472 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence8472(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站_河源线上测试_Period=24 8473
func TestAccAliCloudKmsInstance_basic8473(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap8473)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence8473)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "24",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "24",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap8473 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence8473(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站_河源线上测试_Period=6 8471
func TestAccAliCloudKmsInstance_basic8471(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap8471)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence8471)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "6",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "6",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap8471 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence8471(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站_河源线上测试_Period=36 8474
func TestAccAliCloudKmsInstance_basic8474(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap8474)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence8474)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "36",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "36",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap8474 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence8474(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站_河源线上测试_Period=3 8470
func TestAccAliCloudKmsInstance_basic8470(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap8470)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence8470)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "3",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "3",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap8470 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence8470(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Case 国内站_河源线上测试_Period=2 8469
func TestAccAliCloudKmsInstance_basic8469(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_kms_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudKmsInstanceMap8469)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &KmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeKmsInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacckms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudKmsInstanceBasicDependence8469)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_num":         "1",
					"key_num":         "1000",
					"secret_num":      "0",
					"spec":            "1000",
					"renew_status":    "AutoRenewal",
					"product_version": "3",
					"renew_period":    "1",
					"vpc_id":          "${alicloud_vpc.vpc-amp-instance-test2.id}",
					"vswitch_ids": []string{
						"${alicloud_vswitch.vswitcha.id}"},
					"zone_ids": []string{
						"${alicloud_vswitch.vswitcha.zone_id}", "${alicloud_vswitch.vswitchb.zone_id}"},
					"period":       "2",
					"log_storage":  "0",
					"log":          "0",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_num":         "1",
						"key_num":         "1000",
						"secret_num":      CHECKSET,
						"spec":            "1000",
						"renew_status":    "AutoRenewal",
						"product_version": CHECKSET,
						"renew_period":    "1",
						"vpc_id":          CHECKSET,
						"vswitch_ids.#":   "1",
						"zone_ids.#":      "2",
						"period":          "2",
						"log_storage":     "0",
						"log":             CHECKSET,
						"payment_type":    "Subscription",
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
				ImportStateVerifyIgnore: []string{"period", "renewal_period_unit"},
			},
		},
	})
}

var AlicloudKmsInstanceMap8469 = map[string]string{
	"ca_certificate_chain_pem": CHECKSET,
	"status":                   CHECKSET,
	"create_time":              CHECKSET,
	"end_date":                 CHECKSET,
}

func AlicloudKmsInstanceBasicDependence8469(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc-amp-instance-test2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "vpc-amp-instance-test2"
}

resource "alicloud_vswitch" "vswitcha" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitcha"
  zone_id      = "cn-heyuan-a"
}

resource "alicloud_vswitch" "vswitchb" {
  vpc_id       = alicloud_vpc.vpc-amp-instance-test2.id
  cidr_block   = "172.16.2.0/24"
  vswitch_name = "vpc-amp-instance-test2-vswitchb"
  zone_id      = "cn-heyuan-b"
}


`, name)
}

// Test Kms Instance. <<< Resource test cases, automatically generated.
