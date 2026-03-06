package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test EventBridge ApiDestination. >>> Resource test cases, automatically generated.
// Case testApiDestination 3087
func TestAccAliCloudEventBridgeApiDestination_basic3087(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_event_bridge_api_destination.default"
	ra := resourceAttrInit(resourceId, AlicloudEventBridgeApiDestinationMap3087)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EventBridgeServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEventBridgeApiDestination")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceventbridge%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEventBridgeApiDestinationBasicDependence3087)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-huhehaote"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test-api-destination-connection",
					"http_api_parameters": []map[string]interface{}{
						{
							"endpoint": "http://127.0.0.1:8001",
							"method":   "POST",
						},
					},
					"api_destination_name": name,
					"connection_name":      "${alicloud_event_bridge_connection.defaultConnection.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "test-api-destination-connection",
						"api_destination_name": name,
						"connection_name":      CHECKSET,
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

var AlicloudEventBridgeApiDestinationMap3087 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEventBridgeApiDestinationBasicDependence3087(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_event_bridge_connection" "defaultConnection" {
  network_parameters {
    network_type = "PublicNetwork"
  }
  connection_name = "test-connection-test-1"
}


`, name)
}

// Case testApiDestination_test_zhaohai 4541
func TestAccAliCloudEventBridgeApiDestination_basic4541(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_event_bridge_api_destination.default"
	ra := resourceAttrInit(resourceId, AlicloudEventBridgeApiDestinationMap4541)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EventBridgeServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEventBridgeApiDestination")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceventbridge%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEventBridgeApiDestinationBasicDependence4541)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-chengdu"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test-api-destination-connection-2",
					"http_api_parameters": []map[string]interface{}{
						{
							"endpoint": "http://127.0.0.1:8002",
							"method":   "GET",
						},
					},
					"api_destination_name": name,
					"connection_name":      "${alicloud_event_bridge_connection.defaultConnection.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "test-api-destination-connection-2",
						"api_destination_name": name,
						"connection_name":      CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test-api-destination-connection",
					"http_api_parameters": []map[string]interface{}{
						{
							"endpoint": "http://127.0.0.1:8001",
							"method":   "POST",
						},
					},
					"connection_name": "${alicloud_event_bridge_connection.defaultfFbwZ0.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":     "test-api-destination-connection",
						"connection_name": CHECKSET,
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

var AlicloudEventBridgeApiDestinationMap4541 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEventBridgeApiDestinationBasicDependence4541(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_event_bridge_connection" "defaultConnection" {
  network_parameters {
    network_type = "PublicNetwork"
  }
  connection_name = "test-connection-test-1"
}

resource "alicloud_event_bridge_connection" "defaultfFbwZ0" {
  network_parameters {
    network_type = "PublicNetwork"
  }
  connection_name = "test-connection-test-1-update"
}


`, name)
}

// Case testApiDestination_副本1694657181745 4537
func TestAccAliCloudEventBridgeApiDestination_basic4537(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_event_bridge_api_destination.default"
	ra := resourceAttrInit(resourceId, AlicloudEventBridgeApiDestinationMap4537)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EventBridgeServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEventBridgeApiDestination")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceventbridge%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEventBridgeApiDestinationBasicDependence4537)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-huhehaote"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test-api-destination-connection-2",
					"http_api_parameters": []map[string]interface{}{
						{
							"endpoint": "http://127.0.0.1:8002",
							"method":   "GET",
						},
					},
					"api_destination_name": name,
					"connection_name":      "${alicloud_event_bridge_connection.defaultConnection.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "test-api-destination-connection-2",
						"api_destination_name": name,
						"connection_name":      CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test-api-destination-connection",
					"http_api_parameters": []map[string]interface{}{
						{
							"endpoint": "http://127.0.0.1:8001",
							"method":   "POST",
						},
					},
					"connection_name": "${alicloud_event_bridge_connection.defaultfFbwZ0.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":     "test-api-destination-connection",
						"connection_name": CHECKSET,
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

var AlicloudEventBridgeApiDestinationMap4537 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEventBridgeApiDestinationBasicDependence4537(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_event_bridge_connection" "defaultConnection" {
  network_parameters {
    network_type = "PublicNetwork"
  }
  connection_name = "test-connection-test-1"
}

resource "alicloud_event_bridge_connection" "defaultfFbwZ0" {
  network_parameters {
    network_type = "PublicNetwork"
  }
  connection_name = "test-connection-test-1-update"
}


`, name)
}

// Test EventBridge ApiDestination. <<< Resource test cases, automatically generated.

func TestAccAliCloudEventBridgeApiDestination_basic0(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.EventBridgeConnectionSupportRegions)
	resourceId := "alicloud_event_bridge_api_destination.default"
	ra := resourceAttrInit(resourceId, AliCloudEventBridgeApiDestinationMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EventBridgeServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEventBridgeApiDestination")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%seventbridgeapidestination%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudEventBridgeApiDestinationBasicDependence0)
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
					"connection_name":      "${alicloud_event_bridge_connection.default.connection_name}",
					"api_destination_name": name,
					"http_api_parameters": []map[string]interface{}{
						{
							"endpoint": "http://127.0.0.1:8001",
							"method":   "POST",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"connection_name":       CHECKSET,
						"api_destination_name":  name,
						"http_api_parameters.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test-api-destination-connection",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "test-api-destination-connection",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"http_api_parameters": []map[string]interface{}{
						{
							"endpoint": "http://127.0.0.1:8002",
							"method":   "GET",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"http_api_parameters.#": "1",
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

func TestAccAliCloudEventBridgeApiDestination_basic0_twin(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.EventBridgeConnectionSupportRegions)
	resourceId := "alicloud_event_bridge_api_destination.default"
	ra := resourceAttrInit(resourceId, AliCloudEventBridgeApiDestinationMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EventBridgeServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEventBridgeApiDestination")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%seventbridgeapidestination%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudEventBridgeApiDestinationBasicDependence0)
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
					"connection_name":      "${alicloud_event_bridge_connection.default.connection_name}",
					"api_destination_name": name,
					"description":          "test-api-destination-connection",
					"http_api_parameters": []map[string]interface{}{
						{
							"endpoint": "http://127.0.0.1:8001",
							"method":   "POST",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"connection_name":       CHECKSET,
						"api_destination_name":  name,
						"description":           "test-api-destination-connection",
						"http_api_parameters.#": "1",
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

var AliCloudEventBridgeApiDestinationMap0 = map[string]string{
	"create_time": CHECKSET,
}

func AliCloudEventBridgeApiDestinationBasicDependence0(name string) string {
	return fmt.Sprintf(`
	variable "name" {
    	default = "%s"
	}

	resource "alicloud_event_bridge_connection" "default" {
  		connection_name = var.name
  		network_parameters {
    		network_type = "PublicNetwork"
  		}
	}
`, name)
}
