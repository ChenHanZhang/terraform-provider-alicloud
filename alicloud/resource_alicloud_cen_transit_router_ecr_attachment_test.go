package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Cen TransitRouterEcrAttachment. >>> Resource test cases, automatically generated.
// Case ECR Attachment_跨账号 12499
func TestAccAliCloudCenTransitRouterEcrAttachment_basic12499(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cen_transit_router_ecr_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudCenTransitRouterEcrAttachmentMap12499)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CenServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCenTransitRouterEcrAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccen%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCenTransitRouterEcrAttachmentBasicDependence12499)
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
					"ecr_id":                                "${alicloud_express_connect_router_express_connect_router.defaultO8Hcfx.id}",
					"cen_id":                                "${alicloud_cen_instance.defaultQKBiay.id}",
					"transit_router_ecr_attachment_name":    name,
					"transit_router_attachment_description": "ecr attachment",
					"transit_router_id":                     "${alicloud_cen_transit_router.defaultQa94Y1.transit_router_id}",
					"ecr_owner_id":                          "${{ref(system, defaultO8Hcfx.AccountId)}}",
					"order_type":                            "PayByCenOwner",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ecr_id":                                CHECKSET,
						"cen_id":                                CHECKSET,
						"transit_router_ecr_attachment_name":    name,
						"transit_router_attachment_description": "ecr attachment",
						"transit_router_id":                     CHECKSET,
						"ecr_owner_id":                          CHECKSET,
						"order_type":                            "PayByCenOwner",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"transit_router_ecr_attachment_name":    name + "_update",
					"transit_router_attachment_description": "test2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"transit_router_ecr_attachment_name":    name + "_update",
						"transit_router_attachment_description": "test2",
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudCenTransitRouterEcrAttachmentMap12499 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudCenTransitRouterEcrAttachmentBasicDependence12499(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "asn" {
    default = <<EOF
4200000666
EOF
}

resource "alicloud_express_connect_router_express_connect_router" "defaultO8Hcfx" {
        alibaba_side_asn = "${var.asn}"
        ecr_name = "resource-test-123456"
}

resource "alicloud_cen_instance" "defaultQKBiay" {
        cen_instance_name = "test-ecr-attachment"
}

resource "alicloud_cen_transit_router" "defaultQa94Y1" {
        cen_id = "${alicloud_cen_instance.defaultQKBiay.id}"
}

resource "alicloud_cen_transit_router_grant_attachment" "defaultnYqPDC" {
        order_type = "PayByCenOwner"
        instance_id = "${alicloud_express_connect_router_express_connect_router.defaultO8Hcfx.id}"
        cen_owner_id = "${{ref(system, defaultQa94Y1.AccountId)}}"
        cen_id = "${alicloud_cen_transit_router.defaultQa94Y1.cen_id}"
        instance_type = "ECR"
}

resource "alicloud_express_connect_router_tr_association" "defaultedPu6c" {
        association_region_id = "${alicloud_cen_transit_router.defaultQa94Y1.region_id}"
        ecr_id = "${alicloud_express_connect_router_express_connect_router.defaultO8Hcfx.id}"
        cen_id = "${alicloud_cen_instance.defaultQKBiay.id}"
        transit_router_id = "${alicloud_cen_transit_router.defaultQa94Y1.transit_router_id}"
        transit_router_owner_id = "${{ref(system, defaultQa94Y1.AccountId)}}"
}


`, name)
}

// Case ECR Attachment 5366
func TestAccAliCloudCenTransitRouterEcrAttachment_basic5366(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cen_transit_router_ecr_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudCenTransitRouterEcrAttachmentMap5366)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CenServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCenTransitRouterEcrAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccen%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCenTransitRouterEcrAttachmentBasicDependence5366)
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
					"ecr_id":                                "${alicloud_express_connect_router_express_connect_router.defaultO8Hcfx.id}",
					"cen_id":                                "${alicloud_cen_instance.defaultQKBiay.id}",
					"transit_router_ecr_attachment_name":    name,
					"transit_router_attachment_description": "ecr attachment",
					"transit_router_id":                     "${alicloud_cen_transit_router.defaultQa94Y1.transit_router_id}",
					"ecr_owner_id":                          "${alicloud_express_connect_router_express_connect_router.defaultO8Hcfx.owner_id}",
					"order_type":                            "PayByCenOwner",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ecr_id":                                CHECKSET,
						"cen_id":                                CHECKSET,
						"transit_router_ecr_attachment_name":    name,
						"transit_router_attachment_description": "ecr attachment",
						"transit_router_id":                     CHECKSET,
						"ecr_owner_id":                          CHECKSET,
						"order_type":                            "PayByCenOwner",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"transit_router_ecr_attachment_name":    name + "_update",
					"transit_router_attachment_description": "test2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"transit_router_ecr_attachment_name":    name + "_update",
						"transit_router_attachment_description": "test2",
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudCenTransitRouterEcrAttachmentMap5366 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudCenTransitRouterEcrAttachmentBasicDependence5366(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "asn" {
  default = <<EOF
4200000666
EOF
}

resource "alicloud_express_connect_router_express_connect_router" "defaultO8Hcfx" {
  alibaba_side_asn = var.asn
  ecr_name         = "resource-test-123456"
}

resource "alicloud_cen_instance" "defaultQKBiay" {
  cen_instance_name = "ecr-attachment-test"
}

resource "alicloud_cen_transit_router" "defaultQa94Y1" {
  cen_id = alicloud_cen_instance.defaultQKBiay.id
}

resource "alicloud_express_connect_router_tr_association" "defaultedPu6c" {
  association_region_id = alicloud_cen_transit_router.defaultQa94Y1.region_id
  ecr_id                = alicloud_express_connect_router_express_connect_router.defaultO8Hcfx.id
  cen_id                = alicloud_cen_instance.defaultQKBiay.id
  transit_router_id     = alicloud_cen_transit_router.defaultQa94Y1.transit_router_id
}


`, name)
}

// Case ECR Attachment_修改OrderType 12490
func TestAccAliCloudCenTransitRouterEcrAttachment_basic12490(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cen_transit_router_ecr_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudCenTransitRouterEcrAttachmentMap12490)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CenServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCenTransitRouterEcrAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccen%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCenTransitRouterEcrAttachmentBasicDependence12490)
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
					"ecr_id":                                "${alicloud_express_connect_router_express_connect_router.defaultO8Hcfx.id}",
					"cen_id":                                "${alicloud_cen_instance.defaultQKBiay.id}",
					"transit_router_ecr_attachment_name":    name,
					"transit_router_attachment_description": "ecr attachment",
					"transit_router_id":                     "${alicloud_cen_transit_router.defaultQa94Y1.transit_router_id}",
					"ecr_owner_id":                          "${alicloud_express_connect_router_express_connect_router.defaultO8Hcfx.owner_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ecr_id":                                CHECKSET,
						"cen_id":                                CHECKSET,
						"transit_router_ecr_attachment_name":    name,
						"transit_router_attachment_description": "ecr attachment",
						"transit_router_id":                     CHECKSET,
						"ecr_owner_id":                          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"transit_router_ecr_attachment_name":    name + "_update",
					"transit_router_attachment_description": "test2",
					"order_type":                            "PayByCenOwner",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"transit_router_ecr_attachment_name":    name + "_update",
						"transit_router_attachment_description": "test2",
						"order_type":                            "PayByCenOwner",
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudCenTransitRouterEcrAttachmentMap12490 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudCenTransitRouterEcrAttachmentBasicDependence12490(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "asn" {
  default = <<EOF
4200000666
EOF
}

resource "alicloud_express_connect_router_express_connect_router" "defaultO8Hcfx" {
  alibaba_side_asn = var.asn
  ecr_name         = "resource-test-123456"
}

resource "alicloud_cen_instance" "defaultQKBiay" {
  cen_instance_name = "test-ecr-attachment"
}

resource "alicloud_cen_transit_router" "defaultQa94Y1" {
  cen_id = alicloud_cen_instance.defaultQKBiay.id
}

resource "alicloud_express_connect_router_tr_association" "defaultedPu6c" {
  association_region_id = alicloud_cen_transit_router.defaultQa94Y1.region_id
  ecr_id                = alicloud_express_connect_router_express_connect_router.defaultO8Hcfx.id
  cen_id                = alicloud_cen_instance.defaultQKBiay.id
  transit_router_id     = alicloud_cen_transit_router.defaultQa94Y1.transit_router_id
}


`, name)
}

// Test Cen TransitRouterEcrAttachment. <<< Resource test cases, automatically generated.
