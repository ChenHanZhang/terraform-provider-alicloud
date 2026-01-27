package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test DataWorks DataSourceSharedRule. >>> Resource test cases, automatically generated.
// Case DataSourceSharedRule-匠承测试_正式副本 8165
func TestAccAliCloudDataWorksDataSourceSharedRule_basic8165(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_data_works_data_source_shared_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudDataWorksDataSourceSharedRuleMap8165)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DataWorksServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDataWorksDataSourceSharedRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccdataworks%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDataWorksDataSourceSharedRuleBasicDependence8165)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shanghai"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"target_project_id": "${alicloud_data_works_project.defaultasjsH5.id}",
					"data_source_id":    "${alicloud_data_works_data_source.defaultvzu0wG.data_source_id}",
					"env_type":          "Prod",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"target_project_id": CHECKSET,
						"data_source_id":    CHECKSET,
						"env_type":          "Prod",
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

var AlicloudDataWorksDataSourceSharedRuleMap8165 = map[string]string{
	"create_time":                CHECKSET,
	"data_source_shared_rule_id": CHECKSET,
}

func AlicloudDataWorksDataSourceSharedRuleBasicDependence8165(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_data_works_project" "defaultQeRfvU" {
  description      = "源项目"
  project_name     = "shared_source2"
  display_name     = "shared_source2"
  pai_task_enabled = false
}

resource "alicloud_data_works_project" "defaultasjsH5" {
  description      = "目标空间"
  project_name     = "shared_target2"
  pai_task_enabled = false
  display_name     = "shared_target2"
}

resource "alicloud_data_works_data_source" "defaultvzu0wG" {
  type                       = "hive"
  data_source_name           = "terraform_test"
  connection_properties      = jsonencode({ "address" : [{ "host" : "127.0.0.1", "port" : "1234" }], "database" : "hive_database", "metaType" : "HiveMetastore", "metastoreUris" : "thrift://123:123", "version" : "2.3.9", "loginMode" : "Anonymous", "securityProtocol" : "authTypeNone", "envType" : "Prod", "properties" : { "key1" : "value1" } })
  project_id                 = alicloud_data_works_project.defaultQeRfvU.id
  connection_properties_mode = "UrlMode"
}


`, name)
}

// Case DataSourceSharedRule-SharedUser_正式副本 8166
func TestAccAliCloudDataWorksDataSourceSharedRule_basic8166(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_data_works_data_source_shared_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudDataWorksDataSourceSharedRuleMap8166)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DataWorksServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDataWorksDataSourceSharedRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccdataworks%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDataWorksDataSourceSharedRuleBasicDependence8166)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shanghai"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"shared_user":       "${alicloud_data_works_data_source.defaultsQybqp.create_user}",
					"target_project_id": "${alicloud_data_works_project.defaultGTFU1x.id}",
					"data_source_id":    "${alicloud_data_works_data_source.defaultsQybqp.data_source_id}",
					"env_type":          "Prod",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"shared_user":       CHECKSET,
						"target_project_id": CHECKSET,
						"data_source_id":    CHECKSET,
						"env_type":          "Prod",
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

var AlicloudDataWorksDataSourceSharedRuleMap8166 = map[string]string{
	"create_time":                CHECKSET,
	"data_source_shared_rule_id": CHECKSET,
}

func AlicloudDataWorksDataSourceSharedRuleBasicDependence8166(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_data_works_project" "defaultFW8Oua" {
  description      = "源项目"
  project_name     = "shared_source1"
  pai_task_enabled = false
  display_name     = "shared_source1"
}

resource "alicloud_data_works_project" "defaultGTFU1x" {
  description      = "目标项目"
  project_name     = "target_source1"
  pai_task_enabled = false
  display_name     = "target_source1"
}

resource "alicloud_data_works_data_source" "defaultsQybqp" {
  type                       = "hive"
  data_source_name           = "terraform_test"
  connection_properties      = jsonencode({ "address" : [{ "host" : "127.0.0.1", "port" : "1234" }], "database" : "hive_database", "metaType" : "HiveMetastore", "metastoreUris" : "thrift://123:123", "version" : "2.3.9", "loginMode" : "Anonymous", "securityProtocol" : "authTypeNone", "envType" : "Prod", "properties" : { "key1" : "value1" } })
  project_id                 = alicloud_data_works_project.defaultFW8Oua.id
  connection_properties_mode = "UrlMode"
}


`, name)
}

// Test DataWorks DataSourceSharedRule. <<< Resource test cases, automatically generated.
