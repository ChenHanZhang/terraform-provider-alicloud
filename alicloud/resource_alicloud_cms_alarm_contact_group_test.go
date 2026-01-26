package alicloud

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func init() {
	resource.AddTestSweepers("alicloud_cms_alarm_contact_group", &resource.Sweeper{
		Name: "alicloud_cms_alarm_contact_group",
		F:    testSweepCmsAlarmContactGroup,
	})
}

func testSweepCmsAlarmContactGroup(region string) error {
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return WrapErrorf(err, "error getting Alicloud client.")
	}
	client := rawClient.(*connectivity.AliyunClient)
	cmsService := CmsService{client}

	prefixes := []string{
		"tf-testAcc",
		"tf_testacc",
	}

	request := cms.CreateDescribeContactGroupListRequest()

	raw, err := cmsService.client.WithCmsClient(func(cmsClient *cms.Client) (interface{}, error) {
		return cmsClient.DescribeContactGroupList(request)
	})
	if err != nil {
		log.Printf("[ERROR] Failed to retrieve Cms Alarm Contact Group in service list: %s", err)
	}

	var response *cms.DescribeContactGroupListResponse
	response, _ = raw.(*cms.DescribeContactGroupListResponse)

	for _, v := range response.ContactGroupList.ContactGroup {
		name := v.Name
		skip := true
		if !sweepAll() {
			for _, prefix := range prefixes {
				if strings.HasPrefix(strings.ToLower(name), strings.ToLower(prefix)) {
					skip = false
					break
				}
			}
			if skip {
				log.Printf("[INFO] Skipping alarm contact group: %s ", name)
				continue
			}
		}
		log.Printf("[INFO] delete alarm contact group: %s ", name)

		request := cms.CreateDeleteContactGroupRequest()
		request.ContactGroupName = v.Name
		_, err := client.WithCmsClient(func(cmsClient *cms.Client) (interface{}, error) {
			return cmsClient.DeleteContactGroup(request)
		})

		if err != nil {
			log.Printf("[ERROR] Failed to delete alarm contact group (%s): %s", name, err)
		}
	}

	return nil
}

func TestAccAlicloudCmsAlarmContactGroup_basic(t *testing.T) {
	var v cms.ContactGroup
	resourceId := "alicloud_cms_alarm_contact_group.default"
	ra := resourceAttrInit(resourceId, CmsAlarmContactGroupMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CmsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCmsAlarmContactGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1000000, 9999999)
	name := fmt.Sprintf("tf-testAccCmsAlarmContactGroup%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, CmsAlarmContactGroupBasicdependence)
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
					"alarm_contact_group_name": "${var.name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"alarm_contact_group_name": name,
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
					"describe": "tf-test-describe",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"describe": "tf-test-describe",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enable_subscribed": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enable_subscribed": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"contacts": []string{"${alicloud_cms_alarm_contact.default.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"contacts.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"describe":          "tf-test-describe-update",
					"contacts":          []string{"${alicloud_cms_alarm_contact.default.id}", "${alicloud_cms_alarm_contact.default0.id}"},
					"enable_subscribed": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"describe":          "tf-test-describe-update",
						"contacts.#":        "2",
						"enable_subscribed": "false",
					}),
				),
			},
		},
	})
}

var CmsAlarmContactGroupMap = map[string]string{}

func CmsAlarmContactGroupBasicdependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
	default = "%s"
}
resource "alicloud_cms_alarm_contact" "default" {
  alarm_contact_name = "${var.name}-1"
  describe           = "For Test 1234567"
  channels_mail      = "hello.uuuu@aaa.com"
  lifecycle {
    ignore_changes = [channels_mail]
  }
}
resource "alicloud_cms_alarm_contact" "default0" {
  alarm_contact_name = "${var.name}-0"
  describe           = "For Test 1234567"
  channels_mail      = "hello.uuuu@aaa.com"
  lifecycle {
    ignore_changes = [channels_mail]
  }
}
`, name)
}

// Test CloudMonitorService AlarmContactGroup. >>> Resource test cases, automatically generated.
// Case AlarmContactGroup资源用例_副本1694769992264 4570
func TestAccAliCloudCloudMonitorServiceAlarmContactGroup_basic4570(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cms_alarm_contact_group.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudMonitorServiceAlarmContactGroupMap4570)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudMonitorServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudMonitorServiceAlarmContactGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudmonitorservice%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudMonitorServiceAlarmContactGroupBasicDependence4570)
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
					"alarm_contact_group_name": name,
					"enable_subscribed":        "true",
					"contact_names": []string{
						"${alicloud_cms_alarm_contact.defaultContact.id}"},
					"describe": "Describe",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"alarm_contact_group_name": name,
						"enable_subscribed":        "true",
						"contact_names.#":          "1",
						"describe":                 "Describe",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enable_subscribed": "false",
					"contact_names": []string{
						"${alicloud_cms_alarm_contact.defaultWESknD.id}", "${alicloud_cms_alarm_contact.defaultwNY3m3.id}", "${alicloud_cms_alarm_contact.defaultdDTgGF.id}"},
					"describe": "Describe33",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enable_subscribed": "false",
						"contact_names.#":   "3",
						"describe":          "Describe33",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"contact_names": []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"contact_names.#": "0",
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

var AlicloudCloudMonitorServiceAlarmContactGroupMap4570 = map[string]string{}

func AlicloudCloudMonitorServiceAlarmContactGroupBasicDependence4570(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cms_alarm_contact" "defaultContact" {
  describe           = "Describe486"
  lang               = "zh-cn"
  alarm_contact_name = "rdktestname24"
}

resource "alicloud_cms_alarm_contact" "defaultWESknD" {
  describe           = "Describe200"
  lang               = "zh-cn"
  alarm_contact_name = "rdktestname779"
}

resource "alicloud_cms_alarm_contact" "defaultwNY3m3" {
  describe           = "Describe432"
  lang               = "zh-cn"
  alarm_contact_name = "rdktestname9"
}

resource "alicloud_cms_alarm_contact" "defaultdDTgGF" {
  describe           = "Describe288"
  lang               = "zh-cn"
  alarm_contact_name = "rdktestname261"
}


`, name)
}

// Case AlarmContactGroup资源用例 3694
func TestAccAliCloudCloudMonitorServiceAlarmContactGroup_basic3694(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cms_alarm_contact_group.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudMonitorServiceAlarmContactGroupMap3694)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudMonitorServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudMonitorServiceAlarmContactGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudmonitorservice%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudMonitorServiceAlarmContactGroupBasicDependence3694)
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
					"alarm_contact_group_name": name,
					"enable_subscribed":        "true",
					"contact_names": []string{
						"${alicloud_cms_alarm_contact.defaultContact.id}"},
					"describe": "Describe",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"alarm_contact_group_name": name,
						"enable_subscribed":        "true",
						"contact_names.#":          "1",
						"describe":                 "Describe",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enable_subscribed": "false",
					"contact_names": []string{
						"${alicloud_cms_alarm_contact.defaultWESknD.id}", "${alicloud_cms_alarm_contact.defaultwNY3m3.id}", "${alicloud_cms_alarm_contact.defaultdDTgGF.id}"},
					"describe": "Describe33",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enable_subscribed": "false",
						"contact_names.#":   "3",
						"describe":          "Describe33",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"contact_names": []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"contact_names.#": "0",
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

var AlicloudCloudMonitorServiceAlarmContactGroupMap3694 = map[string]string{}

func AlicloudCloudMonitorServiceAlarmContactGroupBasicDependence3694(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_cms_alarm_contact" "defaultContact" {
  describe           = "Describe50"
  lang               = "zh-cn"
  alarm_contact_name = "rdktestname258"
}

resource "alicloud_cms_alarm_contact" "defaultWESknD" {
  describe           = "Describe312"
  lang               = "zh-cn"
  alarm_contact_name = "rdktestname228"
}

resource "alicloud_cms_alarm_contact" "defaultwNY3m3" {
  describe           = "Describe611"
  lang               = "zh-cn"
  alarm_contact_name = "rdktestname732"
}

resource "alicloud_cms_alarm_contact" "defaultdDTgGF" {
  describe           = "Describe942"
  lang               = "zh-cn"
  alarm_contact_name = "rdktestname906"
}


`, name)
}

// Test CloudMonitorService AlarmContactGroup. <<< Resource test cases, automatically generated.
