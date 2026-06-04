package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAliCloudEipanycastAnycastEipAddressAttachment_basic(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eipanycast_anycast_eip_address_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudEipanycastAnycastEipAddressAttachmentMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EipanycastServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEipanycastAnycastEipAddressAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testAcc%sAlicloudEipanycastAnycastEipAddressAttachment%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEipanycastAnycastEipAddressAttachmentBasicDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
		},

		IDRefreshName: resourceId,
		ProviderFactories: testAccProviderFactory,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					// "bind_instance_region_id" must be consistent with the region of slb instance.
					"anycast_id":              "${alicloud_eipanycast_anycast_eip_address.default.id}",
					"bind_instance_id":        "${alicloud_slb_load_balancer.default.id}",
					"bind_instance_region_id": "${data.alicloud_regions.current_regions.regions.0.id}",
					"bind_instance_type":      "SlbInstance",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"anycast_id":              CHECKSET,
						"bind_instance_id":        CHECKSET,
						"bind_instance_region_id": CHECKSET,
						"bind_instance_type":      "SlbInstance",
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

var AlicloudEipanycastAnycastEipAddressAttachmentMap = map[string]string{
	"bind_time": CHECKSET,
}

func AlicloudEipanycastAnycastEipAddressAttachmentBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
	default = "%s"
}

data "alicloud_vpcs" "default"{
	name_regex = "default-NODELETING"
}
data "alicloud_slb_zones" "default" {
	available_slb_address_type = "vpc"
}

data "alicloud_vswitches" "default" {
	vpc_id  = data.alicloud_vpcs.default.ids.0
	zone_id = data.alicloud_slb_zones.default.zones.0.id
}

resource "alicloud_slb_load_balancer" "default" {
	address_type = "intranet"
	vswitch_id = data.alicloud_vswitches.default.ids[0]
	load_balancer_name = var.name
	load_balancer_spec = "slb.s1.small"
    master_zone_id = "${data.alicloud_slb_zones.default.zones.0.id}"
}

resource "alicloud_eipanycast_anycast_eip_address" "default" {
  anycast_eip_address_name = "${var.name}"
  service_location = "ChineseMainland"
}

data "alicloud_regions" "current_regions" {
  current = true
}

`, name)
}

// Test Eipanycast AnycastEipAddressAttachment. >>> Resource test cases, automatically generated.
// Case 3732
func TestAccAliCloudEipanycastAnycastEipAddressAttachment_basic3732(t *testing.T) {
	var v map[string]interface{}
	testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
	resourceId := "alicloud_eipanycast_anycast_eip_address_attachment.default"

	ra := resourceAttrInit(resourceId, AlicloudEipanycastAnycastEipAddressAttachmentMap3732)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EipanycastServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEipanycastAnycastEipAddressAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%seipanycastanycasteipaddressattachment%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEipanycastAnycastEipAddressAttachmentBasicDependence3732)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName:     resourceId,
		ProviderFactories: testAccProviderFactoriesMultiRegion(4),
		CheckDestroy:      rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"bind_instance_id":        "${alicloud_instance.defaultEcs2.network_interface_id}",
					"bind_instance_type":      "NetworkInterface",
					"anycast_id":              "${alicloud_eipanycast_anycast_eip_address_attachment.defaultEfYBJY.anycast_id}",
					"bind_instance_region_id": "${data.alicloud_regions.current_regions.regions.0.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
				testAccCheck(map[string]string{
						"bind_instance_id":        CHECKSET,
						"bind_instance_type":      "NetworkInterface",
						"anycast_id":              CHECKSET,
						"bind_instance_region_id": CHECKSET,
					}),
				),
			},
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"association_mode": "Default",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheckEipanycastEipAddressAttachmentExistsWithProviders(resourceId, v, &providers),
			//		testAccCheck(map[string]string{
			//			"association_mode": "Default",
			//		}),
			//	),
			//},
			{
				Config: testAccConfig(map[string]interface{}{
					"pop_locations": []map[string]interface{}{
						{
							"pop_location": "cn-guangzhou-pop",
						},
						{
							"pop_location": "cn-shanghai-pop",
						},
						{
							"pop_location": "cn-beijing-pop",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
				testAccCheck(map[string]string{
						"pop_locations.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"pop_locations": []map[string]interface{}{
						{
							"pop_location": "cn-chengdu-pop",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
				testAccCheck(map[string]string{
						"pop_locations.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"pop_locations": []map[string]interface{}{
						{
							"pop_location": "cn-guangzhou-pop",
						},
						{
							"pop_location": "cn-shanghai-pop",
						},
						{
							"pop_location": "cn-beijing-pop",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
				testAccCheck(map[string]string{
						"pop_locations.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"pop_locations": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
				testAccCheck(map[string]string{
						"pop_locations.#": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"pop_locations": []map[string]interface{}{
						{
							"pop_location": "cn-chengdu-pop",
						},
						{
							"pop_location": "cn-beijing-pop",
						},
					},
					"association_mode": "Normal",
				}),
				Check: resource.ComposeTestCheckFunc(
				testAccCheck(map[string]string{
						"pop_locations.#":  "2",
						"association_mode": "Normal",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"bind_instance_id":        "${alicloud_instance.defaultEcs2.network_interface_id}",
					"bind_instance_type":      "NetworkInterface",
					"association_mode":        "Normal",
					"anycast_id":              "${alicloud_eipanycast_anycast_eip_address.defaultXkpFRs.id}",
					"bind_instance_region_id": "${data.alicloud_regions.current_regions.regions.0.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
				testAccCheck(map[string]string{
						"bind_instance_id":        CHECKSET,
						"bind_instance_type":      "NetworkInterface",
						"association_mode":        "Normal",
						"anycast_id":              CHECKSET,
						"bind_instance_region_id": CHECKSET,
					}),
				),
			},
		},
	})
}

var AlicloudEipanycastAnycastEipAddressAttachmentMap3732 = map[string]string{
	"status": CHECKSET,
}

func AlicloudEipanycastAnycastEipAddressAttachmentBasicDependence3732(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

%s

data "alicloud_zones" "default" {
  provider                    = alicloudalt
  available_disk_category     = "cloud_efficiency"
  available_resource_creation = "VSwitch"
}

data "alicloud_images" "default" {
  provider    = alicloudalt
  name_regex  = "^ubuntu_18.*64"
  most_recent = true
  owners      = "system"
}

data "alicloud_instance_types" "default" {
  provider          = alicloudalt
  availability_zone = data.alicloud_zones.default.zones[0].id
  cpu_core_count    = 1
  memory_size       = 2
}

resource "alicloud_vpc" "defaultVpc" {
  provider   = alicloudalt
  vpc_name   = var.name
  cidr_block = "192.168.0.0/16"
}

resource "alicloud_vswitch" "defaultVsw" {
  provider   = alicloudalt
  vpc_id     = alicloud_vpc.defaultVpc.id
  cidr_block = "192.168.0.0/24"
  zone_id    = data.alicloud_zones.default.zones.0.id
}

resource "alicloud_security_group" "defaultuBsECI" {
  provider = alicloudalt
  vpc_id   = alicloud_vpc.defaultVpc.id
}

resource "alicloud_instance" "default9KDlN7" {
  provider                   = alicloudalt
  image_id                   = data.alicloud_images.default.images[0].id
  instance_type              = data.alicloud_instance_types.default.instance_types[0].id
  instance_name              = var.name
  security_groups            = ["${alicloud_security_group.defaultuBsECI.id}"]
  availability_zone          = alicloud_vswitch.defaultVsw.zone_id
  instance_charge_type       = "PostPaid"
  system_disk_category       = "cloud_efficiency"
  vswitch_id                 = alicloud_vswitch.defaultVsw.id
}

resource "alicloud_eipanycast_anycast_eip_address" "defaultXkpFRs" {
  service_location = "ChineseMainland"
}

resource "alicloud_eipanycast_anycast_eip_address_attachment" "defaultEfYBJY" {
  provider                = alicloudalt
  bind_instance_id        = alicloud_instance.default9KDlN7.network_interface_id
  bind_instance_type      = "NetworkInterface"
  bind_instance_region_id = "cn-beijing"
  anycast_id              = alicloud_eipanycast_anycast_eip_address.defaultXkpFRs.id
}

resource "alicloud_vpc" "defaultVpc2" {
  vpc_name   = "${var.name}6"
  cidr_block = "192.168.0.0/16"
}

data "alicloud_regions" "current_regions" {
  current = true
}

data "alicloud_zones" "default2" {
  available_disk_category     = "cloud_efficiency"
  available_resource_creation = "VSwitch"
}

data "alicloud_images" "default2" {
  name_regex  = "^ubuntu_18.*64"
  most_recent = true
  owners      = "system"
}

data "alicloud_instance_types" "default2" {
  availability_zone = data.alicloud_zones.default2.zones[0].id
  cpu_core_count    = 1
  memory_size       = 2
}

resource "alicloud_vswitch" "defaultdsVsw2" {
  vpc_id     = alicloud_vpc.defaultVpc2.id
  cidr_block = "192.168.0.0/24"
  zone_id    = data.alicloud_zones.default2.zones.1.id
}

resource "alicloud_security_group" "defaultuBsECI2" {
  vpc_id   = alicloud_vpc.defaultVpc2.id
}

resource "alicloud_instance" "defaultEcs2" {
  image_id                   = data.alicloud_images.default2.images[0].id
  instance_type              = data.alicloud_instance_types.default2.instance_types[0].id
  instance_name              = var.name
  security_groups            = ["${alicloud_security_group.defaultuBsECI2.id}"]
  availability_zone          = alicloud_vswitch.defaultdsVsw2.zone_id
  instance_charge_type       = "PostPaid"
  system_disk_category       = "cloud_efficiency"
  vswitch_id                 = alicloud_vswitch.defaultdsVsw2.id
}


`, name, configAlternateRegionProvider("cn-beijing"))
}

func AlicloudEipanycastAnycastEipAddressAttachmentBasicDependence3732_region(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

%s
%s

data "alicloud_zones" "default" {
  provider                    = alicloudalt
  available_disk_category     = "cloud_efficiency"
  available_resource_creation = "VSwitch"
}

data "alicloud_images" "default" {
  provider    = alicloudalt
  name_regex  = "^ubuntu_18.*64"
  most_recent = true
  owners      = "system"
}

data "alicloud_instance_types" "default" {
  provider          = alicloudalt
  availability_zone = data.alicloud_zones.default.zones[0].id
  cpu_core_count    = 1
  memory_size       = 2
}

resource "alicloud_vpc" "defaultVpc" {
  provider   = alicloudalt
  vpc_name   = var.name
  cidr_block = "192.168.0.0/16"
}

resource "alicloud_vswitch" "defaultVsw" {
  provider   = alicloudalt
  vpc_id     = alicloud_vpc.defaultVpc.id
  cidr_block = "192.168.0.0/24"
  zone_id    = data.alicloud_zones.default.zones.0.id
}

resource "alicloud_security_group" "defaultuBsECI" {
  provider = alicloudalt
  vpc_id   = alicloud_vpc.defaultVpc.id
}

resource "alicloud_instance" "default9KDlN7" {
  provider                   = alicloudalt
  image_id                   = data.alicloud_images.default.images[0].id
  instance_type              = data.alicloud_instance_types.default.instance_types[0].id
  instance_name              = var.name
  security_groups            = ["${alicloud_security_group.defaultuBsECI.id}"]
  availability_zone          = alicloud_vswitch.defaultVsw.zone_id
  instance_charge_type       = "PostPaid"
  system_disk_category       = "cloud_efficiency"
  vswitch_id                 = alicloud_vswitch.defaultVsw.id
}

data "alicloud_regions" "current_regions" {
  provider = alicloudalt2
  current = true
}

resource "alicloud_eipanycast_anycast_eip_address" "defaultXkpFRs" {
  provider         = alicloudalt2
  service_location = "international"
}

resource "alicloud_eipanycast_anycast_eip_address_attachment" "defaultEfYBJY" {
  provider                = alicloudalt
  bind_instance_id        = alicloud_instance.default9KDlN7.network_interface_id
  bind_instance_type      = "NetworkInterface"
  bind_instance_region_id = "cn-hongkong"
  anycast_id              = alicloud_eipanycast_anycast_eip_address.defaultXkpFRs.id
}

resource "alicloud_vpc" "defaultVpc2" {
  provider   = alicloudalt2
  vpc_name   = "${var.name}6"
  cidr_block = "192.168.0.0/16"
}

data "alicloud_zones" "default2" {
  provider                    = alicloudalt2
  available_disk_category     = "cloud_efficiency"
  available_resource_creation = "VSwitch"
}

data "alicloud_images" "default2" {
  provider    = alicloudalt2
  name_regex  = "^ubuntu_18.*64"
  most_recent = true
  owners      = "system"
}

data "alicloud_instance_types" "default2" {
  provider          = alicloudalt2
  availability_zone = data.alicloud_zones.default2.zones[0].id
  cpu_core_count    = 1
  memory_size       = 2
}

resource "alicloud_vswitch" "defaultdsVsw2" {
  provider   = alicloudalt2
  vpc_id     = alicloud_vpc.defaultVpc2.id
  cidr_block = "192.168.0.0/24"
  zone_id    = data.alicloud_zones.default2.zones.1.id
}

resource "alicloud_security_group" "defaultuBsECI2" {
  provider = alicloudalt2
  vpc_id   = alicloud_vpc.defaultVpc2.id
}

resource "alicloud_instance" "defaultEcs2" {
  provider                   = alicloudalt2
  image_id                   = data.alicloud_images.default2.images[0].id
  instance_type              = data.alicloud_instance_types.default2.instance_types[0].id
  instance_name              = var.name
  security_groups            = ["${alicloud_security_group.defaultuBsECI2.id}"]
  availability_zone          = alicloud_vswitch.defaultdsVsw2.zone_id
  instance_charge_type       = "PostPaid"
  system_disk_category       = "cloud_efficiency"
  vswitch_id                 = alicloud_vswitch.defaultdsVsw2.id
}


`, name, configNamedRegionalProvider(ProviderNameAlicloudAlt, "cn-hongkong"), configNamedRegionalProvider(ProviderNameAlicloudAlt2, "eu-central-1"))
}

// Case 3732  twin
func TestAccAliCloudEipanycastAnycastEipAddressAttachment_basic3732_twin(t *testing.T) {
	var v map[string]interface{}
	testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
	resourceId := "alicloud_eipanycast_anycast_eip_address_attachment.default"

	ra := resourceAttrInit(resourceId, AlicloudEipanycastAnycastEipAddressAttachmentMap3732)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EipanycastServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEipanycastAnycastEipAddressAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%seipanycastanycasteipaddressattachment%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEipanycastAnycastEipAddressAttachmentBasicDependence3732)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName:     resourceId,
		ProviderFactories: testAccProviderFactoriesMultiRegion(4),
		CheckDestroy:      rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"bind_instance_id":        "${alicloud_instance.defaultEcs2.network_interface_id}",
					"bind_instance_type":      "NetworkInterface",
					"anycast_id":              "${alicloud_eipanycast_anycast_eip_address_attachment.defaultEfYBJY.anycast_id}",
					"bind_instance_region_id": "${data.alicloud_regions.current_regions.regions.0.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
				testAccCheck(map[string]string{
						"bind_instance_id":        CHECKSET,
						"bind_instance_type":      "NetworkInterface",
						"anycast_id":              CHECKSET,
						"bind_instance_region_id": CHECKSET,
					}),
				),
			},
		},
	})
}

func TestAccAliCloudEipanycastAnycastEipAddressAttachment_basic3732_region(t *testing.T) {
	var v map[string]interface{}
	testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
	resourceId := "alicloud_eipanycast_anycast_eip_address_attachment.default"

	ra := resourceAttrInit(resourceId, AlicloudEipanycastAnycastEipAddressAttachmentMap3732)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EipanycastServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEipanycastAnycastEipAddressAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%seipanycastanycasteipaddressattachment%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEipanycastAnycastEipAddressAttachmentBasicDependence3732_region)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName:     resourceId,
		ProviderFactories: testAccProviderFactoriesMultiRegion(4),
		CheckDestroy:      rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"bind_instance_id":        "${alicloud_instance.defaultEcs2.network_interface_id}",
					"bind_instance_type":      "NetworkInterface",
					"association_mode":        "Normal",
					"anycast_id":              "${alicloud_eipanycast_anycast_eip_address.defaultXkpFRs.id}",
					"bind_instance_region_id": "eu-central-1",
				}),
				Check: resource.ComposeTestCheckFunc(
				testAccCheck(map[string]string{
						"bind_instance_id":        CHECKSET,
						"bind_instance_type":      "NetworkInterface",
						"association_mode":        "Normal",
						"anycast_id":              CHECKSET,
						"bind_instance_region_id": CHECKSET,
					}),
				),
			},
		},
	})
}

func testAccCheckEipanycastEipAddressAttachmentDestroyWithProviders(providers *[]*schema.Provider) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, provider := range *providers {
			if provider.Meta() == nil {
				continue
			}
			if err := testAccCheckEipanycastEipAddressAttachmentDestroyWithProvider(s, provider); err != nil {
				return err
			}
		}
		return nil
	}
}

func testAccCheckEipanycastEipAddressAttachmentDestroyWithProvider(s *terraform.State, provider *schema.Provider) error {
	client := provider.Meta().(*connectivity.AliyunClient)
	eipanycastServiceV2 := EipanycastServiceV2{client}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "alicloud_eipanycast_anycast_eip_address_attachment" {
			continue
		}
		_, err := eipanycastServiceV2.DescribeEipanycastAnycastEipAddressAttachment(rs.Primary.ID)
		if err != nil {
			if NotFoundError(err) {
				continue
			}
			return err
		} else {
			return fmt.Errorf("alicloud_eipanycast_anycast_eip_address_attachment still exist, ID %s ", fmt.Sprint(rs.Primary.ID))
		}
	}

	return nil
}

func testAccCheckEipanycastEipAddressAttachmentExistsWithProviders(n string, res map[string]interface{}, providers *[]*schema.Provider) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s ", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No alicloud_eipanycast_anycast_eip_address_attachment ID is set. ")
		}
		for _, provider := range *providers {
			if provider.Meta() == nil {
				continue
			}

			client := provider.Meta().(*connectivity.AliyunClient)
			eipanycastServiceV2 := EipanycastServiceV2{client}

			resp, err := eipanycastServiceV2.DescribeEipanycastAnycastEipAddressAttachment(rs.Primary.ID)
			if err != nil {
				if NotFoundError(err) {
					continue
				}
				return err
			}
			res = resp
			return nil
		}
		return fmt.Errorf("alicloud_cen_inter_region_traffic_qos_queue not found")
	}
}

// Test Eipanycast AnycastEipAddressAttachment. <<< Resource test cases, automatically generated.
