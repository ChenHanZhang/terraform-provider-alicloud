package alicloud

import (
	"fmt"
	"log"
	"os"
	"testing"

	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func init() {
	resource.AddTestSweepers("alicloud_instance", &resource.Sweeper{
		Name: "alicloud_instance",
		F:    testSweepInstances,
		// When implemented, these should be removed firstly
		// Now, the resource alicloud_havip_attachment has been published.
		//Dependencies: []string{
		//	"alicloud_havip_attachment",
		//},
	})
}

func testSweepInstances(region string) error {
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return fmt.Errorf("error getting Alicloud client: %s", err)
	}
	client := rawClient.(*connectivity.AliyunClient)

	prefixes := []string{
		"tf-testAcc",
		"tf_testAcc",
	}

	var insts []ecs.Instance
	req := ecs.CreateDescribeInstancesRequest()
	req.RegionId = client.RegionId
	req.PageSize = requests.NewInteger(PageSizeLarge)
	req.PageNumber = requests.NewInteger(1)
	for {
		raw, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.DescribeInstances(req)
		})
		if err != nil {
			return fmt.Errorf("Error retrieving Instances: %s", err)
		}
		resp, _ := raw.(*ecs.DescribeInstancesResponse)
		if resp == nil || len(resp.Instances.Instance) < 1 {
			break
		}
		insts = append(insts, resp.Instances.Instance...)

		if len(resp.Instances.Instance) < PageSizeLarge {
			break
		}

		page, err := getNextpageNumber(req.PageNumber)
		if err != nil {
			return err
		}
		req.PageNumber = page
	}

	sweeped := false
	vpcService := VpcService{client}
	for _, v := range insts {
		name := v.InstanceName
		id := v.InstanceId
		skip := true
		for _, prefix := range prefixes {
			if strings.HasPrefix(strings.ToLower(name), strings.ToLower(prefix)) {
				skip = false
				break
			}
		}
		// If a slb name is set by other service, it should be fetched by vswitch name and deleted.
		if skip {
			if need, err := vpcService.needSweepVpc(v.VpcAttributes.VpcId, v.VpcAttributes.VSwitchId); err == nil {
				skip = !need
			}

		}
		if skip {
			log.Printf("[INFO] Skipping Instance: %s (%s)", name, id)
			continue
		}
		log.Printf("[INFO] Deleting Instance: %s (%s)", name, id)
		if v.DeletionProtection {
			request := ecs.CreateModifyInstanceAttributeRequest()
			request.InstanceId = id
			request.DeletionProtection = requests.NewBoolean(false)
			_, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
				return ecsClient.ModifyInstanceAttribute(request)
			})
			if err != nil {
				log.Printf("[ERROR] %#v", WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR))
				continue
			}
		}
		if v.InstanceChargeType == string(PrePaid) {
			request := ecs.CreateModifyInstanceChargeTypeRequest()
			request.InstanceIds = convertListToJsonString(append(make([]interface{}, 0, 1), id))
			request.InstanceChargeType = string(PostPaid)
			request.IncludeDataDisks = requests.NewBoolean(true)
			_, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
				return ecsClient.ModifyInstanceChargeType(request)
			})
			if err != nil {
				log.Printf("[ERROR] %#v", WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR))
				continue
			}
			time.Sleep(3 * time.Second)
		}

		req := ecs.CreateDeleteInstanceRequest()
		req.InstanceId = id
		req.Force = requests.NewBoolean(true)
		_, err := client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.DeleteInstance(req)
		})
		if err != nil {
			log.Printf("[ERROR] Failed to delete Instance (%s (%s)): %s", name, id, err)
		} else {
			sweeped = true
		}
	}
	if sweeped {
		// Waiting 20 seconds to eusure these instances have been deleted.
		time.Sleep(20 * time.Second)
	}
	return nil
}

func TestAccAliCloudECSInstanceBasic(t *testing.T) {
	var v ecs.Instance
	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceBasicConfigDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, connectivity.EcsClassicSupportedRegions)
			testAccPreCheckWithAccountSiteType(t, DomesticSite)
			testAccClassicNetworkResources(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.0.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_efficiency",
					"instance_name":                 "${var.name}",
					"key_name":                      "${alicloud_key_pair.default.key_name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"security_enhancement_strategy": "Active",
					"resource_group_id":             "${var.resource_group_id}",
					// The specified parameter "UserData" only support the vpc and IoOptimized Instance.
					//"user_data" :                    "I_am_user_data",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":     name,
						"key_name":          name,
						"role_name":         NOSET,
						"vswitch_id":        REMOVEKEY,
						"user_data":         REMOVEKEY,
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run"},
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_release_time": time.Now().Add(10 * time.Hour).Format("2006-01-02T15:04:05Z"),
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_release_time": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_release_time": "",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_release_time": "",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"image_id": "${data.alicloud_images.default.images.1.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"image_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_groups": []string{"${alicloud_security_group.default.0.id}", "${alicloud_security_group.default.1.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_groups.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_name": fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_change", rand),
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name": fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_change", rand),
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_description", rand),
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_description", rand),
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"internet_max_bandwidth_out": "50",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"internet_max_bandwidth_out": "50",
						"public_ip":                  CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"internet_charge_type": "PayByBandwidth",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"internet_charge_type": "PayByBandwidth",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"host_name": "hostNameExample",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"host_name": "hostNameExample",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"password": "Password123",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"password": "Password123",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"system_disk_size": "50",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"system_disk_size": "50",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"system_disk_name":        fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_system_disk_name", rand),
					"system_disk_description": fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_system_disk_description", rand),
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"system_disk_name":        fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_system_disk_name", rand),
						"system_disk_description": fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_system_disk_description", rand),
					}),
				),
			},
			// private_ip cannot be set separately from vpc
			/*{
				Config: testAccConfig(map[string]interface{}{
					"private_ip": "172.16.0.10",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"private_ip": "172.16.0.10",
					}),
				),
			},*/
			// only burstable instances support this attribute.
			/*{
				Config: testAccConfig(map[string]interface{}{
					"credit_specification": "Unlimited",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"credit_specification": "Unlimited",
					}),
				),
			},*/
			{
				Config: testAccConfig(map[string]interface{}{
					"volume_tags": map[string]string{
						"tag1": "test",
						"Tag2": "Test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"volume_tags.%":    "2",
						"volume_tags.tag1": "test",
						"volume_tags.Tag2": "Test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"foo":  "foo",
						"Bar":  "Bar",
						"tag0": "value0",
						"tag1": "value1",
						"tag2": "value2",
						"tag3": "value3",
						"tag4": "value4",
						"tag5": "value5",
						"tag6": "value6",
						"tag7": "value7",
						"tag8": "value8",
						"tag9": "value9",
						"tagA": "valueA",
						"tagB": "valueB",
						"tagC": "valueC",
						"tagD": "valueD",
						"tagE": "valueE",
						"tagF": "valueF",
						"tagG": "valueG",
						"tagH": "valueH",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":    "20",
						"tags.foo":  "foo",
						"tags.Bar":  "Bar",
						"tags.tag9": "value9",
						"tags.tagH": "valueH",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"deletion_protection": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deletion_protection": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "Stopped",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "Stopped",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"system_disk_size": "60",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"system_disk_size": "60",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_type":              "${data.alicloud_instance_types.default.instance_types.0.id}",
					"security_groups":            []string{"${alicloud_security_group.default.0.id}"},
					"instance_name":              fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d", rand),
					"description":                fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d", rand),
					"internet_max_bandwidth_out": "0",
					"host_name":                  REMOVEKEY,
					"password":                   REMOVEKEY,
					// "credit_specification":       "Standard",

					"system_disk_size": "70",
					"volume_tags":      REMOVEKEY,
					"tags":             REMOVEKEY,

					"deletion_protection": "false",
					"status":              "Running",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{

						"tags.%":   "0",
						"tags.Bar": REMOVEKEY,
						"tags.foo": REMOVEKEY,

						"instance_name": fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d", rand),

						"volume_tags.%":    "0",
						"volume_tags.tag1": REMOVEKEY,
						"volume_tags.Tag2": REMOVEKEY,

						"image_id":          CHECKSET,
						"instance_type":     CHECKSET,
						"security_groups.#": "1",

						"availability_zone":             CHECKSET,
						"system_disk_category":          "cloud_efficiency",
						"spot_strategy":                 "NoSpot",
						"spot_price_limit":              "0",
						"security_enhancement_strategy": "Active",

						"description":      fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d", rand),
						"host_name":        CHECKSET,
						"password":         "",
						"is_outdated":      NOSET,
						"system_disk_size": "70",

						// "credit_specification": "Standard",

						"private_ip": CHECKSET,
						"public_ip":  CHECKSET,
						"status":     "Running",

						"internet_charge_type":       string(PayByBandwidth),
						"internet_max_bandwidth_out": "0",

						"deletion_protection": "false",
					}),
				),
			},
		},
	})
}

func TestAccAliCloudECSInstanceVpc(t *testing.T) {
	var v ecs.Instance
	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstanceConfigVpc%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceECSInstanceVpcDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.TestSalveRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.0.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_efficiency",
					"instance_name":                 "${var.name}",
					"key_name":                      "${alicloud_key_pair.default.key_name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"security_enhancement_strategy": "Active",
					"user_data":                     "I_am_user_data",
					"vswitch_id":                    "${alicloud_vswitch.default.id}",
					"role_name":                     "${alicloud_ram_role.default.name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name": name,
						"key_name":      name,
						"role_name":     name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_release_time": time.Now().Add(10 * time.Hour).Format("2006-01-02T15:04:05Z"),
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_release_time": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_release_time": "",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_release_time": "",
					}),
				),
			},
			// renew will be ignored for post paid
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_status": "AutoRenewal",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_status": NOSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_renew_period": "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_renew_period": NOSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "I_am_user_data_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "I_am_user_data_update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I_am_user_data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SV9hbV91c2VyX2RhdGE=",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_groups": []string{"${alicloud_security_group.default.0.id}", "${alicloud_security_group.default.1.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_groups.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_name": name + "_change",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name": name + "_change",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "_description",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "_description",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"internet_max_bandwidth_out": "50",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"internet_max_bandwidth_out": "50",
						"public_ip":                  CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"internet_charge_type": "PayByBandwidth",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"internet_charge_type": "PayByBandwidth",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"host_name": "hostNameExample",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"host_name": "hostNameExample",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"password": "Password123",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"password": "Password123",
					}),
				),
			},
			// only burstable instances support this attribute.
			/*{
				Config: testAccConfig(map[string]interface{}{
					"credit_specification": "Unlimited",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"credit_specification": "Unlimited",
					}),
				),
			},*/
			{
				Config: testAccConfig(map[string]interface{}{
					"system_disk_size": "50",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"system_disk_size": "50",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"private_ip": "${cidrhost(alicloud_vswitch.default.cidr_block, 100)}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"private_ip": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"volume_tags": map[string]string{
						"tag1": "test",
						"Tag2": "Test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"volume_tags.%":    "2",
						"volume_tags.tag1": "test",
						"volume_tags.Tag2": "Test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"foo":  "foo",
						"Bar":  "Bar",
						"tag0": "value0",
						"tag1": "value1",
						"tag2": "value2",
						"tag3": "value3",
						"tag4": "value4",
						"tag5": "value5",
						"tag6": "value6",
						"tag7": "value7",
						"tag8": "value8",
						"tag9": "value9",
						"tagA": "valueA",
						"tagB": "valueB",
						"tagC": "valueC",
						"tagD": "valueD",
						"tagE": "valueE",
						"tagF": "valueF",
						"tagG": "valueG",
						"tagH": "valueH",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":    "20",
						"tags.foo":  "foo",
						"tags.Bar":  "Bar",
						"tags.tag9": "value9",
						"tags.tagH": "valueH",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"deletion_protection": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deletion_protection": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_type": "${data.alicloud_instance_types.default.instance_types.1.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_type": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_type":              "${data.alicloud_instance_types.default.instance_types.0.id}",
					"security_groups":            []string{"${alicloud_security_group.default.0.id}"},
					"instance_name":              name,
					"description":                name,
					"internet_max_bandwidth_out": "0",
					"host_name":                  REMOVEKEY,
					"password":                   REMOVEKEY,
					// "credit_specification":       "Standard",

					"system_disk_size": "70",
					"private_ip":       REMOVEKEY,
					"volume_tags": map[string]string{
						"tag1": "test",
					},
					"tags": REMOVEKEY,

					"deletion_protection": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{

						"tags.%":    "0",
						"tags.Bar":  REMOVEKEY,
						"tags.foo":  REMOVEKEY,
						"tags.tag9": REMOVEKEY,
						"tags.tagH": REMOVEKEY,

						"instance_name": name,

						"volume_tags.%":    "1",
						"volume_tags.Tag2": REMOVEKEY,

						"image_id":          CHECKSET,
						"instance_type":     CHECKSET,
						"security_groups.#": "1",

						"availability_zone":             CHECKSET,
						"system_disk_category":          "cloud_efficiency",
						"spot_strategy":                 "NoSpot",
						"spot_price_limit":              "0",
						"security_enhancement_strategy": "Active",
						"vswitch_id":                    CHECKSET,

						"description":      name,
						"host_name":        CHECKSET,
						"password":         "",
						"is_outdated":      NOSET,
						"system_disk_size": "70",

						// "credit_specification": "Standard",

						"private_ip": CHECKSET,
						"public_ip":  "",
						"status":     "Running",

						"internet_charge_type":       string(PayByBandwidth),
						"internet_max_bandwidth_out": "0",

						"deletion_protection": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"system_disk_name": fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_system_disk_name", rand),
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"system_disk_name": fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_system_disk_name", rand),
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"system_disk_description": fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_system_disk_description", rand),
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"system_disk_description": fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_system_disk_description", rand),
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"system_disk_name":        fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_system_disk_name", rand),
					"system_disk_description": fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_system_disk_description", rand),
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"system_disk_name":        fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_system_disk_name", rand),
						"system_disk_description": fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d_system_disk_description", rand),
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_groups": []string{"${alicloud_security_group.default.0.id}", "${alicloud_security_group.default.1.id}", "${alicloud_security_group.default.2.id}", "${alicloud_security_group.default.3.id}", "${alicloud_security_group.default.4.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_groups.#": "5",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_groups": []string{"${alicloud_security_group.default.0.id}", "${alicloud_security_group.update.0.id}", "${alicloud_security_group.update.1.id}", "${alicloud_security_group.default.3.id}", "${alicloud_security_group.default.4.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_groups.#": "5",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_groups": []string{"${alicloud_security_group.default.0.id}", "${alicloud_security_group.update.0.id}", "${alicloud_security_group.update.1.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_groups.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_groups": []string{"${alicloud_security_group.update.0.id}", "${alicloud_security_group.default.1.id}", "${alicloud_security_group.update.1.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_groups.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_id":          "${alicloud_vpc.vpcUpdate.id}",
					"vswitch_id":      "${alicloud_vswitch.vswitchUpdate.id}",
					"security_groups": []string{"${alicloud_security_group.vpcUpdateSg.0.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_id":            CHECKSET,
						"vswitch_id":        CHECKSET,
						"security_groups.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_groups": []string{"${alicloud_security_group.vpcUpdateSg.0.id}", "${alicloud_security_group.vpcUpdateSg.1.id}", "${alicloud_security_group.vpcUpdateSg.2.id}", "${alicloud_security_group.vpcUpdateSg.3.id}", "${alicloud_security_group.vpcUpdateSg.4.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_groups.#": "5",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_groups": []string{"${alicloud_security_group.vpcUpdateSg.0.id}", "${alicloud_security_group.vpcUpdateSg2.0.id}", "${alicloud_security_group.vpcUpdateSg2.1.id}", "${alicloud_security_group.vpcUpdateSg.3.id}", "${alicloud_security_group.vpcUpdateSg.4.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_groups.#": "5",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_groups": []string{"${alicloud_security_group.vpcUpdateSg.0.id}", "${alicloud_security_group.vpcUpdateSg2.0.id}", "${alicloud_security_group.vpcUpdateSg2.1.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_groups.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_groups": []string{"${alicloud_security_group.vpcUpdateSg2.0.id}", "${alicloud_security_group.vpcUpdateSg.1.id}", "${alicloud_security_group.vpcUpdateSg2.1.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_groups.#": "3",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run"},
			},
		},
	})
}

func TestAccAliCloudECSInstancePrepaid(t *testing.T) {
	var v ecs.Instance
	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstanceConfigPrePaid%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstancePrePaidConfigDependence)
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
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.0.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_efficiency",
					"instance_name":                 "${var.name}",
					"key_name":                      "${alicloud_key_pair.default.key_name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"security_enhancement_strategy": "Active",
					"user_data":                     "I_am_user_data",
					"period":                        "1",
					"instance_charge_type":          "PrePaid",
					"vpc_id":                        "${alicloud_vpc.default.id}",
					"vswitch_id":                    "${alicloud_vswitch.default.id}",
					"role_name":                     "${alicloud_ram_role.default.name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_id":               CHECKSET,
						"instance_name":        name,
						"key_name":             name,
						"role_name":            name,
						"instance_charge_type": "PrePaid",
						"period":               "1",
						"period_unit":          "Month",
						"renewal_status":       "Normal",
						"auto_renew_period":    "0",
						"force_delete":         "false",
						"include_data_disks":   "true",
						"dry_run":              "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_charge_type": "PostPaid",
					"status":               "Stopped",
					"stopped_mode":         "StopCharging",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_charge_type": "PostPaid",
						"status":               "Stopped",
						"stopped_mode":         "StopCharging",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"period":               "1",
					"period_unit":          "Month",
					"instance_charge_type": "PrePaid",
					"status":               "Running",
					"stopped_mode":         REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"period":               "1",
						"period_unit":          "Month",
						"instance_charge_type": "PrePaid",
						"status":               "Running",
						"stopped_mode":         "Not-applicable",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"image_id": "${data.alicloud_images.default.images.1.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"image_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_groups": []string{"${alicloud_security_group.default.0.id}", "${alicloud_security_group.default.1.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_groups.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"force_delete": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"force_delete": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_name": name + "_change",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name": name + "_change",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "_description",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "_description",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"internet_max_bandwidth_out": "50",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"internet_max_bandwidth_out": "50",
						"public_ip":                  CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"internet_charge_type": "PayByBandwidth",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"internet_charge_type": "PayByBandwidth",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"host_name": "hostNameExample",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"host_name": "hostNameExample",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"password": "Password123",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"password": "Password123",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"system_disk_size": "50",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"system_disk_size": "50",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"private_ip": "${cidrhost(alicloud_vswitch.default.cidr_block, 100)}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"private_ip": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"volume_tags": map[string]string{
						"tag1": "test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"volume_tags.%":    "1",
						"volume_tags.tag1": "test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"foo": "foo",
						"bar": "bar",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":   "2",
						"tags.foo": "foo",
						"tags.bar": "bar",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			// Message: The operation is not permitted due to deletion protection only support postPaid instance
			/*{
				Config: testAccConfig(map[string]interface{}{
					"deletion_protection": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deletion_protection": "true",
					}),
				),
			},*/
			{
				Config: testAccConfig(map[string]interface{}{
					"period": "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"period": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"period_unit": "Week",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"period_unit": "Week",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_status": "AutoRenewal",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_status":    "AutoRenewal",
						"auto_renew_period": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_renew_period": "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_renew_period": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"include_data_disks": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"include_data_disks": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"dry_run": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"dry_run": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_groups":            []string{"${alicloud_security_group.default.0.id}"},
					"instance_name":              name,
					"description":                name,
					"internet_max_bandwidth_out": "0",
					"host_name":                  REMOVEKEY,
					"password":                   REMOVEKEY,

					"system_disk_size": "70",
					"private_ip":       REMOVEKEY,
					"volume_tags":      REMOVEKEY,
					"tags":             REMOVEKEY,

					"period":             "1",
					"period_unit":        REMOVEKEY,
					"renewal_status":     REMOVEKEY,
					"auto_renew_period":  REMOVEKEY,
					"include_data_disks": REMOVEKEY,
					"dry_run":            REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"period":             "1",
						"period_unit":        "Month",
						"renewal_status":     REMOVEKEY,
						"auto_renew_period":  REMOVEKEY,
						"include_data_disks": REMOVEKEY,
						"dry_run":            REMOVEKEY,

						"tags.%":   "0",
						"tags.bar": REMOVEKEY,
						"tags.foo": REMOVEKEY,

						"instance_name": name,

						"volume_tags.%":    REMOVEKEY,
						"volume_tags.tag1": REMOVEKEY,

						"image_id":          CHECKSET,
						"instance_type":     CHECKSET,
						"security_groups.#": "1",

						"availability_zone":             CHECKSET,
						"system_disk_category":          "cloud_efficiency",
						"spot_strategy":                 "NoSpot",
						"spot_price_limit":              "0",
						"security_enhancement_strategy": "Active",
						"vswitch_id":                    CHECKSET,

						"description":      name,
						"host_name":        CHECKSET,
						"password":         "",
						"is_outdated":      NOSET,
						"system_disk_size": "70",

						"private_ip": CHECKSET,
						"public_ip":  "",
						"status":     "Running",

						"internet_charge_type":       string(PayByBandwidth),
						"internet_max_bandwidth_out": "0",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"period", "security_enhancement_strategy", "data_disks", "dry_run", "force_delete",
					"include_data_disks"},
			},
		},
	})
}

func TestAccAliCloudECSInstancePrepaidAll(t *testing.T) {
	var v ecs.Instance

	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstanceConfigPrePaid%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstancePrePaidConfigDependence)

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
					"image_id":        "${data.alicloud_images.default.images.0.id}",
					"security_groups": []string{"${alicloud_security_group.default.0.id}"},
					"instance_type":   "${data.alicloud_instance_types.default.instance_types.0.id}",

					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_efficiency",
					"instance_name":                 "${var.name}",
					"key_name":                      "${alicloud_key_pair.default.key_name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"security_enhancement_strategy": "Active",
					"user_data":                     "I_am_user_data",
					"period":                        "1",
					"instance_charge_type":          "PrePaid",
					"vswitch_id":                    "${alicloud_vswitch.default.id}",
					"role_name":                     "${alicloud_ram_role.default.name}",
					"renewal_status":                "AutoRenewal",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name": name,
						"key_name":      name,
						"role_name":     name,

						"instance_charge_type": "PrePaid",
						"period":               "1",
						"period_unit":          "Month",
						"renewal_status":       "AutoRenewal",
						"auto_renew_period":    "1",
						"force_delete":         "false",
						"include_data_disks":   "true",
						"dry_run":              "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_renew_period": "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_renew_period": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_status": "NotRenewal",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_status": "NotRenewal",
						// "auto_renew_period": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"force_delete": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"force_delete": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"force_delete": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"force_delete": "true",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"period", "security_enhancement_strategy", "data_disks", "dry_run", "force_delete",
					"include_data_disks"},
			},
		},
	})
}

func TestAccAliCloudECSInstanceDataDisks(t *testing.T) {
	var v ecs.Instance
	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstanceDataDisks%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceVpcConfigDependence)
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
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.0.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_efficiency",
					"instance_name":                 "${var.name}",
					"key_name":                      "${alicloud_key_pair.default.key_name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"security_enhancement_strategy": "Active",
					"user_data":                     "${base64encode(\"I am the user data\")}",
					"instance_charge_type":          "PostPaid",
					"vswitch_id":                    "${alicloud_vswitch.default.id}",
					"role_name":                     "${alicloud_ram_role.default.name}",
					"network_interfaces": []map[string]string{
						{
							"network_interface_id": "${alicloud_ecs_network_interface.default.id}",
						},
					},
					"data_disks": []map[string]string{
						{
							"name":                 "${var.name}-1",
							"size":                 "20",
							"device":               "/dev/xvdb",
							"delete_with_instance": "true",
						},
						{
							"name":     "${var.name}-2",
							"size":     "20",
							"category": "cloud_ssd",
						},
						{
							"name":      "${var.name}-3",
							"size":      "20",
							"encrypted": "true",
						},
						{
							"name":       "${var.name}-4",
							"size":       "20",
							"encrypted":  "true",
							"kms_key_id": "${alicloud_kms_key.key.id}",
						},
						{
							"name":        "${var.name}-5",
							"size":        "20",
							"description": "${var.name} description",
						},
						{
							"name":              "${var.name}-6",
							"size":              "20",
							"category":          "cloud_ssd",
							"performance_level": "PL1",
						},
					},
					"force_delete": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":                        name,
						"key_name":                             name,
						"role_name":                            name,
						"user_data":                            "SSBhbSB0aGUgdXNlciBkYXRh",
						"network_interfaces.#":                 "1",
						"data_disks.#":                         "6",
						"data_disks.0.name":                    name + "-1",
						"data_disks.0.size":                    "20",
						"data_disks.0.category":                "cloud_efficiency",
						"data_disks.0.encrypted":               "false",
						"data_disks.0.kms_key_id":              "",
						"data_disks.0.snapshot_id":             "",
						"data_disks.0.auto_snapshot_policy_id": "",
						"data_disks.0.delete_with_instance":    "true",
						"data_disks.0.description":             "",
						"data_disks.0.performance_level":       "",
						"data_disks.0.device":                  "/dev/xvdb",
						"data_disks.1.name":                    name + "-2",
						"data_disks.1.size":                    "20",
						"data_disks.1.category":                "cloud_ssd",
						"data_disks.1.encrypted":               "false",
						"data_disks.1.kms_key_id":              "",
						"data_disks.1.snapshot_id":             "",
						"data_disks.1.auto_snapshot_policy_id": "",
						"data_disks.1.delete_with_instance":    "true",
						"data_disks.1.description":             "",
						"data_disks.1.performance_level":       "",
						"data_disks.2.name":                    name + "-3",
						"data_disks.2.category":                "cloud_efficiency",
						"data_disks.2.encrypted":               "true",
						"data_disks.2.kms_key_id":              "",
						"data_disks.2.snapshot_id":             "",
						"data_disks.2.auto_snapshot_policy_id": "",
						"data_disks.2.delete_with_instance":    "true",
						"data_disks.2.description":             "",
						"data_disks.2.performance_level":       "",
						"data_disks.3.name":                    name + "-4",
						"data_disks.3.category":                "cloud_efficiency",
						"data_disks.3.encrypted":               "true",
						"data_disks.3.kms_key_id":              CHECKSET,
						"data_disks.3.snapshot_id":             "",
						"data_disks.3.auto_snapshot_policy_id": "",
						"data_disks.3.delete_with_instance":    "true",
						"data_disks.3.description":             "",
						"data_disks.3.performance_level":       "",
						"data_disks.4.name":                    name + "-5",
						"data_disks.4.category":                "cloud_efficiency",
						"data_disks.4.encrypted":               "false",
						"data_disks.4.kms_key_id":              "",
						"data_disks.4.snapshot_id":             "",
						"data_disks.4.auto_snapshot_policy_id": "",
						"data_disks.4.delete_with_instance":    "true",
						"data_disks.4.description":             name + " description",
						"data_disks.4.performance_level":       "",
						"data_disks.5.name":                    name + "-6",
						"data_disks.5.category":                "cloud_ssd",
						"data_disks.5.encrypted":               "false",
						"data_disks.5.kms_key_id":              "",
						"data_disks.5.snapshot_id":             "",
						"data_disks.5.auto_snapshot_policy_id": "",
						"data_disks.5.delete_with_instance":    "true",
						"data_disks.5.description":             "",
						"data_disks.5.performance_level":       "PL1",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "system_disk_performance_level", "data_disks", "dry_run", "force_delete",
					"include_data_disks"},
			},
		},
	})
}

func TestAccAliCloudECSInstanceSpotInstanceLimit(t *testing.T) {
	var v ecs.Instance

	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAccEcsInstanceConfigSpot%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, testAccCheckSpotInstanceDependence)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, connectivity.TestSalveRegions)
			testAccPreCheckWithAccountSiteType(t, DomesticSite)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vswitch_id":                 "${alicloud_vswitch.default.id}",
					"image_id":                   "${data.alicloud_images.default.images.0.id}",
					"availability_zone":          "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"instance_type":              "${data.alicloud_instance_types.default.instance_types.0.id}",
					"system_disk_category":       "cloud_efficiency",
					"internet_charge_type":       "PayByTraffic",
					"internet_max_bandwidth_out": "5",
					"security_groups":            []string{"${alicloud_security_group.default.id}"},
					"instance_name":              "${var.name}",
					"spot_strategy":              "SpotWithPriceLimit",
					"spot_price_limit":           "1.002",
					"spot_duration":              "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_charge_type":       "PostPaid",
						"spot_strategy":              "SpotWithPriceLimit",
						"spot_price_limit":           "1.002",
						"internet_max_bandwidth_out": "5",
						"public_ip":                  CHECKSET,
						"user_data":                  REMOVEKEY,
						"spot_duration":              "1",
					}),
				),
			},
		},
	})
}

func TestAccAliCloudECSInstanceMulti(t *testing.T) {
	var v ecs.Instance
	resourceId := "alicloud_instance.default.9"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstanceConfigMulti%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceVpcConfigDependence)
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
					"count":                         "10",
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.0.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_efficiency",
					"instance_name":                 "${var.name}",
					"key_name":                      "${alicloud_key_pair.default.key_name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"security_enhancement_strategy": "Active",
					"user_data":                     "${base64encode(\"I am the user data\")}",
					"vswitch_id":                    "${alicloud_vswitch.default.id}",
					"role_name":                     "${alicloud_ram_role.default.name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name": name,
						"key_name":      name,
						"role_name":     name,
						"user_data":     "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
		},
	})
}

func TestAccAliCloudECSInstanceHpcCluster(t *testing.T) {
	var v ecs.Instance

	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstanceHpcCluster%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceVpcHpcClusterIDDependence)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.EcsSccSupportedRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.0.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_efficiency",
					"instance_name":                 "${var.name}",
					"key_name":                      "${alicloud_key_pair.default.key_name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"security_enhancement_strategy": "Active",
					"user_data":                     "${base64encode(\"I am the user data\")}",
					"vswitch_id":                    "${alicloud_vswitch.default.id}",
					"hpc_cluster_id":                "${alicloud_ecs_hpc_cluster.default.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":  name,
						"key_name":       name,
						"hpc_cluster_id": CHECKSET,
						"user_data":      "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run"},
			},
		},
	})
}

func resourceInstanceVpcHpcClusterIDDependence(name string) string {
	return fmt.Sprintf(`
data "alicloud_instance_types" "default" {
  instance_type_family = "ecs.sccg7"
  network_type         = "Vpc"
}
data "alicloud_images" "default" {
  name_regex  = "^aliyun_3_x64_20G_scc*"
  owners      = "system"
}

data "alicloud_instance_types" "essd" {
 	cpu_core_count    = 2
	memory_size       = 4
 	system_disk_category = "cloud_essd"
}
resource "alicloud_ecs_hpc_cluster" "default" {
  name          = "${var.name}"
  description   = "For Terraform Test"
}

resource "alicloud_vpc" "default" {
    vpc_name = var.name
}

resource "alicloud_vswitch" "default" {
  vpc_id  = alicloud_vpc.default.id
  zone_id = data.alicloud_instance_types.default.instance_types.0.availability_zones.0
  cidr_block = cidrsubnet(alicloud_vpc.default.cidr_block, 8, 8)
  vswitch_name   = var.name
}

resource "alicloud_security_group" "default" {
  count = "2"
  name   = "${var.name}"
  vpc_id = alicloud_vpc.default.id
}
resource "alicloud_security_group_rule" "default" {
	count = 2
  	type = "ingress"
  	ip_protocol = "tcp"
  	nic_type = "intranet"
  	policy = "accept"
  	port_range = "22/22"
  	priority = 1
  	security_group_id = "${element(alicloud_security_group.default.*.id,count.index)}"
  	cidr_ip = "172.16.0.0/24"
}

variable "name" {
	default = "%s"
}

resource "alicloud_key_pair" "default" {
	key_pair_name = "${var.name}"
}
resource "alicloud_kms_key" "key" {
        description             = var.name
        pending_window_in_days  = "7"
        status               = "Enabled"
}

`, name)
}

func TestAccAliCloudECSInstanceSecondaryIps(t *testing.T) {
	var v ecs.Instance

	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceVpcSecondaryIps)

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
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"system_disk_category":          "cloud_essd",
					"system_disk_performance_level": "PL1",
					"vswitch_id":                    "${alicloud_vswitch.default.id}",
					"host_name":                     "test",
					"instance_name":                 "${var.name}",
					"internet_charge_type":          "PayByTraffic",
					"instance_charge_type":          "PostPaid",
					"password":                      "Tftest123",
					"user_data":                     "${base64encode(\"I am the user data\")}",
					"security_enhancement_strategy": "Active",
					"internet_max_bandwidth_out":    "5",
					"secondary_private_ips": []string{
						"${cidrhost(alicloud_vswitch.default.cidr_block, 191)}",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"image_id":                      CHECKSET,
						"vswitch_id":                    CHECKSET,
						"secondary_private_ips.#":       "1",
						"instance_name":                 name,
						"password":                      "Tftest123",
						"internet_max_bandwidth_out":    "5",
						"public_ip":                     CHECKSET,
						"security_enhancement_strategy": "Active",
						"system_disk_category":          "cloud_essd",
						"system_disk_performance_level": "PL1",
						"user_data":                     "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"secondary_private_ips": []string{
						"${cidrhost(alicloud_vswitch.default.cidr_block, 195)}",
						"${cidrhost(alicloud_vswitch.default.cidr_block, 197)}",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"secondary_private_ips.#": "2",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password", "security_enhancement_strategy", "dry_run"},
			},
		},
	})
}

func TestAccAliCloudECSInstanceSecondaryIpCount(t *testing.T) {
	var v ecs.Instance

	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	rc := resourceCheckInit(resourceId, &v, func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	})
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAccEcsInstanceConfigBasic%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceVpcSecondaryIps)

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
					"image_id":                           "${data.alicloud_images.default.images.0.id}",
					"security_groups":                    []string{"${alicloud_security_group.default.id}"},
					"instance_type":                      "${data.alicloud_instance_types.default.instance_types.0.id}",
					"system_disk_category":               "cloud_essd",
					"vswitch_id":                         "${alicloud_vswitch.default.id}",
					"host_name":                          "test",
					"instance_name":                      "${var.name}",
					"internet_charge_type":               "PayByTraffic",
					"instance_charge_type":               "PostPaid",
					"password":                           "Tftest123",
					"user_data":                          "I_am_user_data",
					"security_enhancement_strategy":      "Active",
					"internet_max_bandwidth_out":         "5",
					"allocate_public_ip":                 "true",
					"secondary_private_ip_address_count": "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"image_id":                           CHECKSET,
						"vswitch_id":                         CHECKSET,
						"instance_name":                      name,
						"password":                           "Tftest123",
						"internet_max_bandwidth_out":         "5",
						"allocate_public_ip":                 "true",
						"public_ip":                          CHECKSET,
						"security_enhancement_strategy":      "Active",
						"system_disk_category":               "cloud_essd",
						"secondary_private_ip_address_count": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"secondary_private_ip_address_count": "3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"secondary_private_ip_address_count": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"secondary_private_ip_address_count": "0"},
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"secondary_private_ip_address_count": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"allocate_public_ip": "false"},
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"allocate_public_ip": "false",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password", "security_enhancement_strategy", "secondary_private_ip_address_count", "dry_run", "allocate_public_ip"},
			},
		},
	})
}

func TestAccAliCloudECSInstance_DeploymentSetID(t *testing.T) {
	var v ecs.Instance

	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstanceHpcCluster%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceVpcDeploymentSetIDDependence)

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
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_essd",
					"instance_name":                 "${var.name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"security_enhancement_strategy": "Active",
					"user_data":                     "I_am_user_data",
					"vswitch_id":                    "${alicloud_vswitch.default.id}",
					"deployment_set_id":             "${alicloud_ecs_deployment_set.default.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":        name,
						"system_disk_category": "cloud_essd",
						"deployment_set_id":    CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"deployment_set_id": "${alicloud_ecs_deployment_set.update.id}",
				},
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deployment_set_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"deployment_set_id": "",
				},
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deployment_set_id": "",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run"},
			},
		},
	})
}

func TestAccAliCloudECSInstance_StatusUpdated(t *testing.T) {
	var v ecs.Instance

	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstanceConfigVpc%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceVpcConfigDependence)

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
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.0.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_efficiency",
					"instance_name":                 "${var.name}",
					"key_name":                      "${alicloud_key_pair.default.key_name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"security_enhancement_strategy": "Active",
					"user_data":                     "I_am_user_data",
					"vswitch_id":                    "${alicloud_vswitch.default.id}",
					"role_name":                     "${alicloud_ram_role.default.name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name": name,
						"key_name":      name,
						"role_name":     name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "Stopped",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "Stopped",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "Running",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "Running",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":       "Stopped",
					"stopped_mode": "StopCharging",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":       "Stopped",
						"stopped_mode": "StopCharging",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run"},
			},
		},
	})
}

func TestAccAliCloudECSInstanceMetadataOptions(t *testing.T) {
	var v ecs.Instance

	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceMetadataOptionsDependence)

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
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_efficiency",
					"instance_name":                 "${var.name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"security_enhancement_strategy": "Active",
					"user_data":                     "I_am_user_data",
					"vswitch_id":                    "${alicloud_vswitch.default.id}",
					"http_tokens":                   "optional",
					"http_endpoint":                 "enabled",
					"http_put_response_hop_limit":   "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":               name,
						"http_tokens":                 "optional",
						"http_endpoint":               "enabled",
						"http_put_response_hop_limit": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"http_endpoint": "disabled",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"http_endpoint": "disabled",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"http_tokens": "required",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"http_tokens": "required",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"http_tokens":   "optional",
					"http_endpoint": "enabled",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name": name,
						"http_tokens":   "optional",
						"http_endpoint": "enabled",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run"},
			},
		},
	})
}

func TestAccAliCloudECSInstanceIpv6AddressesCount(t *testing.T) {
	var v ecs.Instance

	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceIpv6Dependence)

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
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_efficiency",
					"instance_name":                 "${var.name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"security_enhancement_strategy": "Active",
					"user_data":                     "I_am_user_data",
					"vswitch_id":                    "${alicloud_vswitch.vswitch.id}",
					"ipv6_address_count":            "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":        name,
						"ipv6_address_count":   "1",
						"system_disk_category": "cloud_efficiency",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run"},
			},
		},
	})
}

func TestAccAliCloudECSInstanceIpv6Addresses(t *testing.T) {
	var v ecs.Instance
	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceIpv6Dependence)
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
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_efficiency",
					"instance_name":                 "${var.name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"security_enhancement_strategy": "Active",
					"user_data":                     "I_am_user_data",
					"vswitch_id":                    "${alicloud_vswitch.vswitch.id}",
					"password_inherit":              "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":        name,
						"system_disk_category": "cloud_efficiency",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ipv6_addresses": []string{"${cidrhost(alicloud_vswitch.vswitch.ipv6_cidr_block, 8)}", "${cidrhost(alicloud_vswitch.vswitch.ipv6_cidr_block, 64)}", "${cidrhost(alicloud_vswitch.vswitch.ipv6_cidr_block, 100)}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ipv6_addresses.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ipv6_addresses": []string{"${cidrhost(alicloud_vswitch.vswitch.ipv6_cidr_block, 8)}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ipv6_addresses.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ipv6_addresses": []string{"${cidrhost(alicloud_vswitch.vswitch.ipv6_cidr_block, 8)}", "${cidrhost(alicloud_vswitch.vswitch.ipv6_cidr_block, 100)}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ipv6_addresses.#": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "password_inherit", "dry_run"},
			},
		},
	})
}

func resourceInstanceIpv6Dependence(name string) string {
	return fmt.Sprintf(`
data "alicloud_zones" default {
  available_disk_category     = "cloud_efficiency"
  available_resource_creation = "VSwitch"
}

data "alicloud_images" "default" {
  name_regex  = "^ubuntu_18.*64"
  owners     = "system"
}

data "alicloud_instance_types" "default" {
  availability_zone = data.alicloud_zones.default.zones.0.id
  image_id          = data.alicloud_images.default.images.0.id
  system_disk_category              = "cloud_efficiency"
  cpu_core_count                    = 4
  minimum_eni_ipv6_address_quantity = 2
}

resource "alicloud_vpc" "vpc" {
  vpc_name   = var.name
  cidr_block = "172.16.0.0/12"
  enable_ipv6 = "true"
}

resource "alicloud_vswitch" "vswitch" {
  vswitch_name = var.name
  cidr_block   = "172.16.0.0/21"
  zone_id      = data.alicloud_zones.default.zones.0.id
  vpc_id       = alicloud_vpc.vpc.id
  ipv6_cidr_block_mask = "22"
}

resource "alicloud_security_group" "default" {
  name   = "${var.name}"
  vpc_id = alicloud_vpc.vpc.id
}

variable "name" {
	default = "%s"
}

`, name)
}

func resourceInstanceMetadataOptionsDependence(name string) string {
	return fmt.Sprintf(`
data "alicloud_zones" default {
  available_resource_creation = "Instance"
}

data "alicloud_images" "default" {
  name_regex  = "^ubuntu_18.*64"
  owners     = "system"
}

data "alicloud_instance_types" "default" {
  availability_zone = data.alicloud_zones.default.zones.0.id
  image_id          = data.alicloud_images.default.images.0.id
  cpu_core_count       = 2
  memory_size          = 8
  instance_type_family = "ecs.g6"
}

resource "alicloud_vpc" "default" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = var.name
}

resource "alicloud_vswitch" "default" {
  vpc_id            = alicloud_vpc.default.id
  cidr_block        = cidrsubnet(alicloud_vpc.default.cidr_block, 8, 2)
  zone_id           = data.alicloud_zones.default.zones.0.id
  vswitch_name      = var.name
}

resource "alicloud_security_group" "default" {
  name   = "${var.name}"
  vpc_id = alicloud_vpc.default.id
}

variable "name" {
	default = "%s"
}

`, name)
}

func resourceInstanceVpcDeploymentSetIDDependence(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_zones" "default" {
		available_disk_category     = "cloud_essd"
  		available_resource_creation = "VSwitch"
	}

	data "alicloud_images" "default" {
  		name_regex  = "^ubuntu_[0-9]+_[0-9]+_x64*"
  		most_recent = true
  		owners      = "system"
	}

	data "alicloud_instance_types" "default" {
  		availability_zone    = data.alicloud_zones.default.zones.0.id
  		image_id             = data.alicloud_images.default.images.0.id
		system_disk_category = "cloud_essd"
	}

	resource "alicloud_vpc" "default" {
  		cidr_block = "192.168.0.0/16"
  		vpc_name   = var.name
	}

	resource "alicloud_vswitch" "default" {
  		vpc_id       = alicloud_vpc.default.id
  		cidr_block   = cidrsubnet(alicloud_vpc.default.cidr_block, 8, 2)
  		zone_id      = data.alicloud_zones.default.zones.0.id
  		vswitch_name = var.name
	}

	resource "alicloud_security_group" "default" {
  		name   = var.name
  		vpc_id = alicloud_vpc.default.id
	}

	resource "alicloud_ecs_deployment_set" "default" {
  		strategy            = "Availability"
  		domain              = "Default"
  		granularity         = "Host"
  		deployment_set_name = "example_value"
  		description         = "example_value"
	}

	resource "alicloud_ecs_deployment_set" "update" {
  		strategy            = "Availability"
  		domain              = "Default"
  		granularity         = "Host"
  		deployment_set_name = "example_value"
  		description         = "example_value"
	}
`, name)
}

func resourceInstanceVpcSecondaryIps(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_zones" "default" {
		available_disk_category     = "cloud_essd"
  		available_resource_creation = "VSwitch"
	}

	data "alicloud_images" "default" {
  		name_regex  = "^ubuntu_[0-9]+_[0-9]+_x64*"
  		most_recent = true
  		owners      = "system"
	}

	data "alicloud_instance_types" "default" {
  		availability_zone    = data.alicloud_zones.default.zones.0.id
  		image_id             = data.alicloud_images.default.images.0.id
		system_disk_category = "cloud_essd"
	}

	resource "alicloud_vpc" "default" {
  		vpc_name = var.name
	}

	resource "alicloud_vswitch" "default" {
  		vpc_id       = alicloud_vpc.default.id
  		zone_id      = data.alicloud_instance_types.default.instance_types.0.availability_zones.0
  		cidr_block   = cidrsubnet(alicloud_vpc.default.cidr_block, 8, 8)
  		vswitch_name = var.name
	}

	resource "alicloud_security_group" "default" {
  		name   = var.name
  		vpc_id = alicloud_vpc.default.id
	}
`, name)
}

func resourceInstanceVpcConfigDependence(name string) string {
	return fmt.Sprintf(`
data "alicloud_instance_types" "default" {
	instance_type_family = "ecs.sn1ne"
}

	data "alicloud_images" "default" {
  		name_regex  = "^ubuntu_[0-9]+_[0-9]+_x64*"
  		most_recent = true
  		owners      = "system"
	}

resource "alicloud_vpc" "default" {
    vpc_name = var.name
}

resource "alicloud_vswitch" "default" {
  vpc_id  = alicloud_vpc.default.id
  zone_id = data.alicloud_instance_types.default.instance_types.0.availability_zones.0
  cidr_block = cidrsubnet(alicloud_vpc.default.cidr_block, 8, 8)
  vswitch_name   = var.name
}

resource "alicloud_security_group" "default" {
  count = "2"
  name   = "${var.name}"
  vpc_id = alicloud_vpc.default.id
}
resource "alicloud_security_group_rule" "default" {
	count = 2
  	type = "ingress"
  	ip_protocol = "tcp"
  	nic_type = "intranet"
  	policy = "accept"
  	port_range = "22/22"
  	priority = 1
  	security_group_id = "${element(alicloud_security_group.default.*.id,count.index)}"
  	cidr_ip = "172.16.0.0/24"
}

variable "name" {
	default = "%s"
}

resource "alicloud_ram_role" "default" {
		  name = "${var.name}"
		  document = <<EOF
		{
		  "Statement": [
			{
			  "Action": "sts:AssumeRole",
			  "Effect": "Allow",
			  "Principal": {
				"Service": [
				  "ecs.aliyuncs.com"
				]
			  }
			}
		  ],
		  "Version": "1"
		}
	  EOF
		  force = "true"
}

resource "alicloud_key_pair" "default" {
	key_pair_name = "${var.name}"
}

resource "alicloud_kms_key" "key" {
  description            = var.name
  pending_window_in_days = "7"
  key_state              = "Enabled"
}

resource "alicloud_ecs_network_interface" "default" {
  network_interface_name = var.name
  vswitch_id             = alicloud_vswitch.default.id
  security_group_ids     = [alicloud_security_group.default.0.id]
}
`, name)
}

func resourceECSInstanceVpcDependence(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_zones" "default" {
  		available_resource_creation = "VSwitch"
	}

	data "alicloud_images" "default" {
  		name_regex = "^ubuntu_18.*64"
  		owners     = "system"
	}

	data "alicloud_instance_types" "default" {
  		image_id          = data.alicloud_images.default.images.0.id
  		availability_zone = data.alicloud_zones.default.zones.0.id
	}

	resource "alicloud_vpc" "default" {
  		cidr_block = "192.168.0.0/16"
  		vpc_name   = var.name
	}

	resource "alicloud_vswitch" "default" {
  		vpc_id       = alicloud_vpc.default.id
  		cidr_block   = cidrsubnet(alicloud_vpc.default.cidr_block, 8, 2)
  		zone_id      = data.alicloud_zones.default.zones.0.id
  		vswitch_name = var.name
	}

	resource "alicloud_security_group" "default" {
  		count  = 5
  		name   = var.name
  		vpc_id = alicloud_vpc.default.id
	}

	resource "alicloud_security_group" "update" {
  		count  = 2
  		name   = var.name
  		vpc_id = alicloud_vpc.default.id
	}

	resource "alicloud_vpc" "vpcUpdate" {
  		cidr_block = "10.0.0.0/8"
	}

	resource "alicloud_vswitch" "vswitchUpdate" {
  		cidr_block = "10.1.0.0/16"
  		vpc_id     = alicloud_vpc.vpcUpdate.id
  		zone_id    = data.alicloud_zones.default.zones.0.id
	}

	resource "alicloud_security_group" "vpcUpdateSg" {
  		count  = 5
  		vpc_id = alicloud_vpc.vpcUpdate.id
	}

	resource "alicloud_security_group" "vpcUpdateSg2" {
  		count  = 2
  		vpc_id = alicloud_vpc.vpcUpdate.id
	}

	resource "alicloud_key_pair" "default" {
  		key_pair_name = var.name
	}

	resource "alicloud_kms_key" "key" {
  		description            = var.name
  		pending_window_in_days = "7"
  		key_state              = "Enabled"
	}

	resource "alicloud_ram_role" "default" {
  		name     = var.name
  		force    = "true"
  		document = <<EOF
		{
		  "Statement": [
			{
			  "Action": "sts:AssumeRole",
			  "Effect": "Allow",
			  "Principal": {
				"Service": [
				  "ecs.aliyuncs.com"
				]
			  }
			}
		  ],
		  "Version": "1"
		}
	  EOF
	}
`, name)
}

func resourceInstancePrePaidConfigDependence(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_zones" "default" {
  		available_resource_creation = "VSwitch"
	}

	data "alicloud_images" "default" {
  		name_regex = "^ubuntu_[0-9]+_[0-9]+_x64*"
  		owners     = "system"
	}

	data "alicloud_instance_types" "default" {
  		availability_zone    = data.alicloud_zones.default.zones.0.id
  		image_id             = data.alicloud_images.default.images.0.id
  		instance_charge_type = "PrePaid"
	}

	resource "alicloud_vpc" "default" {
  		vpc_name = var.name
	}

	resource "alicloud_vswitch" "default" {
  		vpc_id       = alicloud_vpc.default.id
  		zone_id      = data.alicloud_instance_types.default.instance_types.0.availability_zones.0
  		cidr_block   = cidrsubnet(alicloud_vpc.default.cidr_block, 8, 8)
  		vswitch_name = var.name
	}

	resource "alicloud_security_group" "default" {
  		count  = 2
  		name   = var.name
  		vpc_id = alicloud_vpc.default.id
	}

	resource "alicloud_ram_role" "default" {
  		name     = var.name
  		document = <<EOF
		{
		  "Statement": [
			{
			  "Action": "sts:AssumeRole",
			  "Effect": "Allow",
			  "Principal": {
				"Service": [
				  "ecs.aliyuncs.com"
				]
			  }
			}
		  ],
		  "Version": "1"
		}
	  EOF
  		force    = "true"
	}

	resource "alicloud_key_pair" "default" {
  		key_pair_name = var.name
	}
`, name)
}

func resourceInstanceBasicConfigDependence(name string) string {
	return fmt.Sprintf(`

data "alicloud_instance_types" "default" {
  cpu_core_count    = 1
  memory_size       = 2
}

data "alicloud_images" "default" {
  name_regex  = "^ubuntu_[0-9]+_[0-9]+_x64*"
  owners      = "system"
}

variable "resource_group_id" {
		default = "%s"
}

resource "alicloud_security_group" "default" {
  count = "2"
  name   = "${var.name}"
}
resource "alicloud_security_group_rule" "default" {
	count = 2
  	type = "ingress"
  	ip_protocol = "tcp"
  	nic_type = "intranet"
  	policy = "accept"
  	port_range = "22/22"
  	priority = 1
  	security_group_id = "${element(alicloud_security_group.default.*.id,count.index)}"
  	cidr_ip = "172.16.0.0/24"
}

variable "name" {
	default = "%s"
}

resource "alicloud_key_pair" "default" {
	key_name = "${var.name}"
}

`, os.Getenv("ALICLOUD_RESOURCE_GROUP_ID"), name)
}

func testAccCheckSpotInstanceDependence(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_instance_types" "special" {
	  	spot_strategy = "SpotWithPriceLimit"
		cpu_core_count       = 2
  		memory_size          = 8
  		instance_type_family = "ecs.g6"
  		image_id             = data.alicloud_images.default.images.0.id
	}

	data "alicloud_images" "default" {
  		name_regex  = "^ubuntu_[0-9]+_[0-9]+_x64*"
  		most_recent = true
  		owners      = "system"
	}

	data "alicloud_instance_types" "default" {
	  	spot_strategy = "SpotWithPriceLimit"
		cpu_core_count       = 2
  		memory_size          = 8
  		instance_type_family = "ecs.g6"
  		image_id             = data.alicloud_images.default.images.0.id
	}

	resource "alicloud_vpc" "default" {
  		vpc_name = var.name
	}

	resource "alicloud_vswitch" "default" {
  		vpc_id       = alicloud_vpc.default.id
  		zone_id      = data.alicloud_instance_types.default.instance_types.0.availability_zones.0
  		cidr_block   = cidrsubnet(alicloud_vpc.default.cidr_block, 8, 8)
  		vswitch_name = var.name
	}

	resource "alicloud_security_group" "default" {
  		name   = var.name
  		vpc_id = alicloud_vpc.default.id
	}
	
	`, name)
}

var testAccInstanceCheckMap = map[string]string{
	"image_id":          CHECKSET,
	"instance_type":     CHECKSET,
	"security_groups.#": "1",

	"availability_zone":    CHECKSET,
	"system_disk_category": "cloud_efficiency",
	//"credit_specification":          "",
	"spot_strategy":    "NoSpot",
	"spot_price_limit": "0",
	// "security_enhancement_strategy": "Active",
	"vswitch_id": CHECKSET,
	"user_data":  "I_am_user_data",

	"description":      "",
	"host_name":        CHECKSET,
	"password":         "",
	"is_outdated":      NOSET,
	"system_disk_size": "40",
	"volume_tags.%":    "0",
	"tags.%":           NOSET,

	"private_ip":                 CHECKSET,
	"public_ip":                  "",
	"status":                     "Running",
	"internet_charge_type":       "PayByTraffic",
	"internet_max_bandwidth_out": "0",
	"instance_charge_type":       "PostPaid",
	// the attributes of below are suppressed  when the value of instance_charge_type is `PostPaid`
	"period":             NOSET,
	"period_unit":        NOSET,
	"renewal_status":     NOSET,
	"auto_renew_period":  NOSET,
	"force_delete":       NOSET,
	"include_data_disks": NOSET,
	"dry_run":            "false",
	"system_disk_id":     CHECKSET,
	"create_time":        CHECKSET,
	"start_time":         CHECKSET,
	"expired_time":       CHECKSET,
}

func TestAccAliCloudECSInstance_OperatorType(t *testing.T) {
	var v ecs.Instance

	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstanceHpcCluster%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceVpcOperatorTypeIDDependence)

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
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.2.id}",
					"availability_zone":             "${alicloud_vswitch.default.zone_id}",
					"system_disk_category":          "cloud_efficiency",
					"instance_name":                 "${var.name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"security_enhancement_strategy": "Active",
					"user_data":                     "I_am_user_data",
					"vswitch_id":                    "${alicloud_vswitch.default.id}",
					"period":                        "1",
					"instance_charge_type":          "PrePaid",
					"force_delete":                  "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":        name,
						"system_disk_category": "cloud_efficiency",
						"instance_charge_type": "PrePaid",
						"period":               "1",
						"period_unit":          "Month",
						"renewal_status":       "Normal",
						"auto_renew_period":    "0",
						"force_delete":         "true",
						"include_data_disks":   "true",
						"dry_run":              "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"operator_type": "downgrade",
					"instance_type": "${data.alicloud_instance_types.default.instance_types.0.id}",
				},
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_type": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"operator_type": "upgrade",
					"instance_type": "${data.alicloud_instance_types.default.instance_types.2.id}",
				},
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_type": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run", "operator_type", "force_delete", "include_data_disks", "period"},
			},
		},
	})
}

func resourceInstanceVpcOperatorTypeIDDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
	default = "%s"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

data "alicloud_images" "default" {
  name_regex  = "^aliyun_3_x64_20G_scc*"
  owners      = "system"
}

resource "alicloud_vpc" "default" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = var.name
}

resource "alicloud_vswitch" "default" {
  vpc_id            = alicloud_vpc.default.id
  cidr_block        = cidrsubnet(alicloud_vpc.default.cidr_block, 8, 2)
  zone_id           = data.alicloud_zones.default.zones.0.id
  vswitch_name      = var.name
}

data "alicloud_instance_types" "default" {
  availability_zone                 = alicloud_vswitch.default.zone_id
  system_disk_category              = "cloud_efficiency"
  cpu_core_count                    = 4
  minimum_eni_ipv6_address_quantity = 1
}

resource "alicloud_security_group" "default" {
  name   = var.name
  vpc_id = alicloud_vpc.default.id
}
`, name)
}

func TestAccAliCloudECSInstance_AutoSnapshotPolicyId(t *testing.T) {
	var v ecs.Instance

	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstanceConfigAutoSnapshotPolicyId%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceAutoSnapshotPolicyIdDependence)

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
					"image_id":                            "${data.alicloud_images.default.images.0.id}",
					"security_groups":                     []string{"${alicloud_security_group.default.id}"},
					"instance_type":                       "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":                   "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":                "cloud_auto",
					"instance_name":                       "${var.name}",
					"spot_strategy":                       "NoSpot",
					"spot_price_limit":                    "0",
					"security_enhancement_strategy":       "Active",
					"user_data":                           "I_am_user_data",
					"vswitch_id":                          "${alicloud_vswitch.default.id}",
					"system_disk_auto_snapshot_policy_id": "${alicloud_ecs_auto_snapshot_policy.default.id}",
					"system_disk_provisioned_iops":        "100",
					"system_disk_bursting_enabled":        "true",
					"data_disks": []map[string]string{
						{
							"name":             "${var.name}-1",
							"size":             "500",
							"category":         "cloud_auto",
							"provisioned_iops": "100",
							"bursting_enabled": "true",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":                       name,
						"system_disk_category":                "cloud_auto",
						"system_disk_auto_snapshot_policy_id": CHECKSET,
						"system_disk_provisioned_iops":        "100",
						"system_disk_bursting_enabled":        "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"system_disk_auto_snapshot_policy_id": "${alicloud_ecs_auto_snapshot_policy.update.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"system_disk_auto_snapshot_policy_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"system_disk_provisioned_iops": "200",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"system_disk_provisioned_iops": "200",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"system_disk_bursting_enabled": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"system_disk_bursting_enabled": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "data_disks", "dry_run"},
			},
		},
	})
}

func resourceInstanceAutoSnapshotPolicyIdDependence(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_images" "default" {
  		name_regex  = "^ubuntu_[0-9]+_[0-9]+_x64*"
  		most_recent = true
  		owners      = "system"
	}

	data "alicloud_instance_types" "default" {
  		image_id             = data.alicloud_images.default.images.0.id
		system_disk_category = "cloud_auto"
	}

	resource "alicloud_vpc" "default" {
  		vpc_name = var.name
	}

	resource "alicloud_vswitch" "default" {
  		vpc_id       = alicloud_vpc.default.id
  		zone_id      = data.alicloud_instance_types.default.instance_types.0.availability_zones.0
  		cidr_block   = cidrsubnet(alicloud_vpc.default.cidr_block, 8, 8)
  		vswitch_name = var.name
	}

	resource "alicloud_security_group" "default" {
  		name   = var.name
  		vpc_id = alicloud_vpc.default.id
	}

	resource "alicloud_ecs_auto_snapshot_policy" "default" {
  		name            = var.name
  		repeat_weekdays = ["1"]
  		retention_days  = -1
  		time_points     = ["1"]
  		tags = {
    		Created = "TF"
    		For     = "acceptance test"
  		}
	}

	resource "alicloud_ecs_auto_snapshot_policy" "update" {
  		name            = "${var.name}_update"
  		repeat_weekdays = ["1"]
  		retention_days  = -1
  		time_points     = ["1"]
  		tags = {
    		Created = "TF"
    		For     = "acceptance test"
  		}
	}
`, name)
}

func TestAccAliCloudECSInstanceSystemDisk(t *testing.T) {
	var v ecs.Instance

	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstanceSystemDisk%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceSystemDiskDependence)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.TestSalveRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.0.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_essd",
					"system_disk_size":              "500",
					"instance_name":                 "${var.name}",
					"key_name":                      "${alicloud_key_pair.default.key_name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"instance_charge_type":          "PostPaid",
					"vswitch_id":                    "${alicloud_vswitch.default.id}",
					"force_delete":                  "true",
					"security_enhancement_strategy": "Active",
					"user_data":                     "I_am_user_data",
					"system_disk_encrypted":         "true",
					"system_disk_kms_key_id":        "${alicloud_kms_key.key.id}",
					"system_disk_encrypt_algorithm": "aes-256",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":          name,
						"key_name":               name,
						"system_disk_encrypted":  "true",
						"system_disk_kms_key_id": CHECKSET,
						"system_disk_category":   "cloud_essd",
						"system_disk_size":       "500",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"system_disk_performance_level": "PL2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"system_disk_performance_level": "PL2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"image_id": "${data.alicloud_images.default.images.1.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"image_id": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_delete", "security_enhancement_strategy", "dry_run", "system_disk_encrypt_algorithm"},
			},
		},
	})
}

func TestAccAliCloudECSInstanceMaintenance(t *testing.T) {
	var v ecs.Instance

	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)

	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstanceMaintenance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceSystemDiskDependence)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.TestSalveRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.0.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_essd",
					"instance_name":                 "${var.name}",
					"key_name":                      "${alicloud_key_pair.default.key_name}",
					"spot_strategy":                 "NoSpot",
					"spot_price_limit":              "0",
					"instance_charge_type":          "PostPaid",
					"vswitch_id":                    "${alicloud_vswitch.default.id}",
					"force_delete":                  "true",
					"security_enhancement_strategy": "Active",
					"user_data":                     "I_am_user_data",
					"maintenance_time": []map[string]interface{}{
						{
							"start_time": "01:00:00",
							"end_time":   "02:00:00",
						},
					},
					"maintenance_action": "Stop",
					"maintenance_notify": "true",
					"enable_jumbo_frame": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":        name,
						"key_name":             name,
						"system_disk_category": "cloud_essd",
						"maintenance_time.#":   "1",
						"maintenance_action":   "Stop",
						"maintenance_notify":   "true",
						"enable_jumbo_frame":   "false",
						"cpu":                  CHECKSET,
						"memory":               CHECKSET,
						"os_name":              CHECKSET,
						"os_type":              CHECKSET,
						"primary_ip_address":   CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"maintenance_time": []map[string]interface{}{
						{
							"start_time": "02:00:00",
							"end_time":   "03:00:00",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"maintenance_time.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"maintenance_action": "AutoRecover",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"maintenance_action": "AutoRecover",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"maintenance_notify": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"maintenance_notify": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enable_jumbo_frame": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enable_jumbo_frame": "true",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_delete", "security_enhancement_strategy", "dry_run", "system_disk_encrypt_algorithm"},
			},
		},
	})
}

func TestAccAliCloudECSInstanceDedicatedHostId(t *testing.T) {
	var v ecs.Instance
	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstanceDedicatedHostId%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceDedicatedHostIdConfigDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.TestSalveRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"image_id":                      "${data.alicloud_images.default.images.0.id}",
					"security_groups":               []string{"${alicloud_security_group.default.id}"},
					"instance_type":                 "${data.alicloud_instance_types.default.instance_types.0.id}",
					"availability_zone":             "${data.alicloud_instance_types.default.instance_types.0.availability_zones.0}",
					"system_disk_category":          "cloud_efficiency",
					"instance_name":                 "${var.name}",
					"key_name":                      "${alicloud_key_pair.default.key_name}",
					"security_enhancement_strategy": "Active",
					"user_data":                     "${base64encode(\"I am the user data\")}",
					"vswitch_id":                    "${alicloud_vswitch.default.id}",
					"role_name":                     "${alicloud_ram_role.default.name}",
					"dedicated_host_id":             "${data.alicloud_ecs_dedicated_hosts.default.hosts.0.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":     name,
						"key_name":          name,
						"role_name":         name,
						"dedicated_host_id": CHECKSET,
						"user_data":         "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run"},
			},
		},
	})
}

func TestAccAliCloudECSInstance_LaunchTemplate(t *testing.T) {
	var v ecs.Instance
	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstanceLaunchTemplate%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceLaunchTemplateDependence)
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
					"launch_template_id":      "${alicloud_ecs_launch_template.default.id}",
					"launch_template_name":    "${alicloud_ecs_launch_template.default.launch_template_name}",
					"launch_template_version": "1",
					"system_disk_category":    "cloud_essd",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"launch_template_id":      CHECKSET,
						"launch_template_name":    CHECKSET,
						"launch_template_version": "1",
						"system_disk_category":    "cloud_essd",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_release_time": time.Now().Add(10 * time.Hour).Format("2006-01-02T15:04:05Z"),
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_release_time": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_release_time": "",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_release_time": "",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_name": name + "_change",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name": name + "_change",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run", "launch_template_id", "launch_template_name", "launch_template_version"},
			},
		},
	})
}

func TestAccAliCloudECSInstanceNetworkInterface0(t *testing.T) {
	var v ecs.Instance
	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceNetworkInterfaceDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.TestSalveRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"image_id":                       "${data.alicloud_images.default.images.0.id}",
					"instance_type":                  "ecs.c8i.48xlarge",
					"availability_zone":              "cn-hangzhou-k",
					"internet_charge_type":           "PayByTraffic",
					"vswitch_id":                     "${alicloud_vswitch.default.id}",
					"internet_max_bandwidth_out":     "10",
					"system_disk_category":           "cloud_essd",
					"instance_name":                  "${var.name}",
					"user_data":                      "I_am_user_data",
					"security_groups":                []string{"${alicloud_security_group.default.id}"},
					"network_interface_traffic_mode": "HighPerformance",
					"network_card_index":             "0",
					"queue_pair_number":              "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":                  name,
						"internet_max_bandwidth_out":     "10",
						"system_disk_category":           "cloud_essd",
						"public_ip":                      CHECKSET,
						"network_interface_traffic_mode": "HighPerformance",
						"network_card_index":             "0",
						"queue_pair_number":              "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"vswitch_id": "${alicloud_vswitch.update.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vswitch_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_groups": []string{"${alicloud_security_group.default.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_groups.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"private_ip": "${cidrhost(alicloud_vswitch.update.cidr_block, 100)}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"private_ip": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run"},
			},
		},
	})
}

func TestAccAliCloudECSInstanceNetworkInterface1(t *testing.T) {
	var v ecs.Instance
	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceNetworkInterfaceDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.TestSalveRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"image_id":                       "${data.alicloud_images.default.images.0.id}",
					"instance_type":                  "ecs.c8i.48xlarge",
					"availability_zone":              "cn-hangzhou-k",
					"internet_charge_type":           "PayByTraffic",
					"vswitch_id":                     "${alicloud_vswitch.default.id}",
					"internet_max_bandwidth_out":     "10",
					"system_disk_category":           "cloud_essd",
					"instance_name":                  "${var.name}",
					"user_data":                      "I_am_user_data",
					"security_groups":                []string{"${alicloud_security_group.default.id}"},
					"network_interface_traffic_mode": "HighPerformance",
					"network_card_index":             "0",
					"queue_pair_number":              "1",
					"network_interfaces": []map[string]interface{}{
						{
							"vswitch_id":                     "${alicloud_vswitch.networkInterface.id}",
							"network_interface_traffic_mode": "Standard",
							"network_card_index":             "1",
							"security_group_ids":             []string{"${alicloud_security_group.networkInterface.id}"},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":                  name,
						"internet_max_bandwidth_out":     "10",
						"system_disk_category":           "cloud_essd",
						"public_ip":                      CHECKSET,
						"network_interface_traffic_mode": "HighPerformance",
						"network_card_index":             "0",
						"queue_pair_number":              "1",
						"network_interfaces.#":           "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"vswitch_id": "${alicloud_vswitch.update.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vswitch_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_groups": []string{"${alicloud_security_group.default.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_groups.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"private_ip": "${cidrhost(alicloud_vswitch.update.cidr_block, 100)}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"private_ip": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run"},
			},
		},
	})
}

func TestAccAliCloudECSInstanceNetworkInterface2(t *testing.T) {
	var v ecs.Instance
	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceNetworkInterfaceDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.TestSalveRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"image_id":                            "${data.alicloud_images.default.images.0.id}",
					"instance_type":                       "ecs.c8i.48xlarge",
					"availability_zone":                   "cn-hangzhou-k",
					"internet_charge_type":                "PayByTraffic",
					"vswitch_id":                          "${alicloud_vswitch.default.id}",
					"internet_max_bandwidth_out":          "10",
					"system_disk_category":                "cloud_essd",
					"instance_name":                       "${var.name}",
					"private_pool_options_match_criteria": "Open",
					"user_data":                           "I_am_user_data",
					"security_groups":                     []string{"${alicloud_security_group.default.id}"},
					"ipv6_addresses":                      []string{"${cidrhost(alicloud_vswitch.default.ipv6_cidr_block, 64)}"},
					"network_interface_traffic_mode":      "Standard",
					"network_interfaces": []map[string]interface{}{
						{
							"vswitch_id":                     "${alicloud_vswitch.networkInterface.id}",
							"network_interface_traffic_mode": "HighPerformance",
							"network_card_index":             "1",
							"queue_pair_number":              "1",
							"security_group_ids":             []string{"${alicloud_security_group.networkInterface.id}"},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":                       name,
						"internet_max_bandwidth_out":          "10",
						"system_disk_category":                "cloud_essd",
						"public_ip":                           CHECKSET,
						"ipv6_addresses.#":                    "1",
						"network_interface_traffic_mode":      "Standard",
						"private_pool_options_match_criteria": "Open",
						"network_interfaces.#":                "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run"},
			},
		},
	})
}

func TestAccAliCloudECSInstanceNetworkInterface3(t *testing.T) {
	var v ecs.Instance
	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceInstanceNetworkInterfaceDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.TestSalveRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"image_id":                       "${data.alicloud_images.default.images.0.id}",
					"instance_type":                  "ecs.c8i.48xlarge",
					"availability_zone":              "cn-hangzhou-k",
					"internet_charge_type":           "PayByTraffic",
					"vswitch_id":                     "${alicloud_vswitch.default.id}",
					"internet_max_bandwidth_out":     "10",
					"system_disk_category":           "cloud_essd",
					"instance_name":                  "${var.name}",
					"user_data":                      "I_am_user_data",
					"security_groups":                []string{"${alicloud_security_group.default.id}"},
					"ipv6_address_count":             "1",
					"network_interface_traffic_mode": "Standard",
					"network_interfaces": []map[string]interface{}{
						{
							"vswitch_id":                     "${alicloud_vswitch.networkInterface.id}",
							"network_interface_traffic_mode": "HighPerformance",
							"network_card_index":             "1",
							"queue_pair_number":              "1",
							"security_group_ids":             []string{"${alicloud_security_group.networkInterface.id}"},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":                  name,
						"internet_max_bandwidth_out":     "10",
						"system_disk_category":           "cloud_essd",
						"public_ip":                      CHECKSET,
						"ipv6_address_count":             "1",
						"network_interface_traffic_mode": "Standard",
						"network_interfaces.#":           "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": tags0,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%": "50",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"tagH": "valueH",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": tags0,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%": "50",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run"},
			},
		},
	})
}

func TestAccAliCloudECSInstancePrivatePool(t *testing.T) {
	var v ecs.Instance
	resourceId := "alicloud_instance.default"
	ra := resourceAttrInit(resourceId, testAccInstanceCheckMap)
	serviceFunc := func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInit(resourceId, &v, serviceFunc)
	rac := resourceAttrCheckInit(rc, ra)
	rand := acctest.RandIntRange(1000, 9999)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAcc%sEcsInstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudEcsInstancePrivatePool)
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
					"image_id":                            "${data.alicloud_images.default.images.0.id}",
					"security_groups":                     []string{"${alicloud_security_group.default.id}"},
					"instance_type":                       "${data.alicloud_ecs_elasticity_assurances.default.assurances.0.allocated_resources.0.instance_type}",
					"availability_zone":                   "${data.alicloud_ecs_elasticity_assurances.default.assurances.0.allocated_resources.0.zone_id}",
					"system_disk_category":                "cloud_efficiency",
					"instance_name":                       "${var.name}",
					"spot_strategy":                       "NoSpot",
					"spot_price_limit":                    "0",
					"security_enhancement_strategy":       "Active",
					"vswitch_id":                          "${alicloud_vswitch.vswitch.id}",
					"user_data":                           "I_am_user_data",
					"private_pool_options_match_criteria": "Target",
					"private_pool_options_id":             "${data.alicloud_ecs_elasticity_assurances.default.assurances.0.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":                       name,
						"system_disk_category":                "cloud_efficiency",
						"private_pool_options_match_criteria": "Target",
						"private_pool_options_id":             CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"private_pool_options_match_criteria": "None",
					"private_pool_options_id":             REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"private_pool_options_match_criteria": "None",
						"private_pool_options_id":             REMOVEKEY,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"private_pool_options_match_criteria": "Open",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"private_pool_options_match_criteria": "Open",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"private_pool_options_match_criteria": "Target",
					"private_pool_options_id":             "${data.alicloud_ecs_elasticity_assurances.default.assurances.0.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"private_pool_options_match_criteria": "Target",
						"private_pool_options_id":             CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_data": "${base64encode(\"I am the user data\")}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_data": "SSBhbSB0aGUgdXNlciBkYXRh",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_enhancement_strategy", "dry_run"},
			},
		},
	})
}

func resourceInstanceDedicatedHostIdConfigDependence(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_ecs_dedicated_hosts" "default" {
  		name_regex = "default-NODELETING"
	}

	data "alicloud_images" "default" {
  		name_regex  = "^ubuntu_[0-9]+_[0-9]+_x64*"
  		most_recent = true
  		owners      = "system"
	}

	data "alicloud_instance_types" "default" {
  		instance_type_family = "ecs.c5"
  		availability_zone    = data.alicloud_ecs_dedicated_hosts.default.hosts.0.zone_id
  		image_id             = data.alicloud_images.default.images.0.id
  		system_disk_category = "cloud_efficiency"
	}

	resource "alicloud_vpc" "default" {
  		cidr_block = "192.168.0.0/16"
  		vpc_name   = var.name
	}

	resource "alicloud_vswitch" "default" {
  		vpc_id       = alicloud_vpc.default.id
  		cidr_block   = cidrsubnet(alicloud_vpc.default.cidr_block, 8, 2)
  		zone_id      = data.alicloud_ecs_dedicated_hosts.default.hosts.0.zone_id
  		vswitch_name = var.name
	}

	resource "alicloud_security_group" "default" {
  		name   = var.name
  		vpc_id = alicloud_vpc.default.id
	}

	resource "alicloud_ram_role" "default" {
  		name     = var.name
  		document = <<EOF
		{
		  "Statement": [
			{
			  "Action": "sts:AssumeRole",
			  "Effect": "Allow",
			  "Principal": {
				"Service": [
				  "ecs.aliyuncs.com"
				]
			  }
			}
		  ],
		  "Version": "1"
		}
	  EOF
  		force    = "true"
	}

	resource "alicloud_key_pair" "default" {
  		key_pair_name = var.name
	}
`, name)
}

func resourceInstanceSystemDiskDependence(name string) string {
	return fmt.Sprintf(`
data "alicloud_instance_types" "default" {
  instance_type_family = "ecs.g8i"
  system_disk_category = "cloud_essd"
  availability_zone    = alicloud_vswitch.default.zone_id
}

data "alicloud_images" "default" {
  name_regex = "^ubuntu_[0-9]+_[0-9]+_x64*"
  owners     = "system"
}

data "alicloud_zones" "default" {
  available_resource_creation = "Instance"
}

resource "alicloud_vpc" "default" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = var.name
}

resource "alicloud_vswitch" "default" {
  vpc_id            = alicloud_vpc.default.id
  cidr_block        = cidrsubnet(alicloud_vpc.default.cidr_block, 8, 2)
  zone_id           = data.alicloud_zones.default.zones.0.id
  vswitch_name      = var.name
}

resource "alicloud_security_group" "default" {
  	count  = "2"
  	name   = var.name
  	vpc_id = alicloud_vpc.default.id
}

resource "alicloud_security_group_rule" "default" {
  	count             = 2
  	type              = "ingress"
  	ip_protocol       = "tcp"
  	nic_type          = "intranet"
  	policy            = "accept"
  	port_range        = "22/22"
  	priority          = 1
  	security_group_id = element(alicloud_security_group.default.*.id, count.index)
  	cidr_ip           = "172.16.0.0/24"
}

variable "name" {
  	default = "%s"
}

resource "alicloud_key_pair" "default" {
  	key_pair_name = var.name
}

resource "alicloud_kms_key" "key" {
  	description            = var.name
  	pending_window_in_days = "7"
  	key_state              = "Enabled"
}
`, name)
}

func resourceInstanceLaunchTemplateDependence(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_zones" "default" {
  		available_disk_category     = "cloud_essd"
  		available_resource_creation = "VSwitch"
	}

	data "alicloud_images" "default" {
  		most_recent = true
  		owners      = "system"
	}

	data "alicloud_instance_types" "default" {
  		availability_zone    = data.alicloud_zones.default.zones.0.id
  		image_id             = data.alicloud_images.default.images.0.id
		system_disk_category = "cloud_essd"
	}

	resource "alicloud_vpc" "default" {
  		vpc_name   = var.name
  		cidr_block = "192.168.0.0/16"
	}

	resource "alicloud_vswitch" "default" {
  		vswitch_name = var.name
  		vpc_id       = alicloud_vpc.default.id
  		cidr_block   = "192.168.192.0/24"
  		zone_id      = data.alicloud_zones.default.zones.0.id
	}

	resource "alicloud_security_group" "default" {
  		name   = var.name
  		vpc_id = alicloud_vpc.default.id
	}

	resource "alicloud_ecs_launch_template" "default" {
  		launch_template_name          = var.name
  		image_id                      = data.alicloud_images.default.images.0.id
  		host_name                     = "hostNameExample"
  		instance_charge_type          = "PostPaid"
  		instance_name                 = var.name
  		instance_type                 = data.alicloud_instance_types.default.instance_types.0.id
  		internet_charge_type          = "PayByTraffic"
  		internet_max_bandwidth_in     = "5"
  		internet_max_bandwidth_out    = "0"
  		security_group_id             = alicloud_security_group.default.id
  		vswitch_id                    = alicloud_vswitch.default.id
  		vpc_id                        = alicloud_vpc.default.id
  		zone_id                       = data.alicloud_zones.default.zones.0.id
  		security_enhancement_strategy = "Active"
  		user_data                     = "I_am_user_data"
  		system_disk {
    		category             = "cloud_ssd"
    		description          = "Test For Terraform"
    		name                 = "terraform-example"
    		size                 = "40"
  		}
  		template_tags = {
    		Created = "tf"
    		For     = "example"
  		}
	}
`, name)
}

func resourceInstanceNetworkInterfaceDependence(name string) string {
	return fmt.Sprintf(`
resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "172.16.0.0/12"
  enable_ipv6 = "true"
}

resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  cidr_block   = "172.16.0.0/16"
  zone_id      = "cn-hangzhou-k"
  vpc_id       = alicloud_vpc.default.id
  ipv6_cidr_block_mask = "22"
}

resource "alicloud_vswitch" "update" {
  vswitch_name = var.name
  cidr_block   = "172.18.0.0/16"
  zone_id      = "cn-hangzhou-k"
  vpc_id       = alicloud_vpc.default.id
  ipv6_cidr_block_mask = "23"
}

resource "alicloud_vswitch" "networkInterface" {
  vswitch_name = var.name
  cidr_block   = "172.19.0.0/16"
  zone_id      = "cn-hangzhou-k"
  vpc_id       = alicloud_vpc.default.id
  ipv6_cidr_block_mask = "25"
}

	data "alicloud_images" "default" {
  		instance_type  = "ecs.c8i.48xlarge"
  		most_recent = true
  		owners      = "system"
	}

resource "alicloud_security_group" "default" {
  name   = "${var.name}"
  vpc_id = alicloud_vpc.default.id
}

resource "alicloud_security_group" "update" {
  name   = "${var.name}"
  vpc_id = alicloud_vpc.default.id
}

resource "alicloud_security_group" "networkInterface" {
  name   = "${var.name}"
  vpc_id = alicloud_vpc.default.id
}

variable "name" {
	default = "%s"
}

`, name)
}

func AliCloudEcsInstancePrivatePool(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_ecs_elasticity_assurances" "default" {
  		status = "Active"
	}

	data "alicloud_images" "default" {
  		instance_type = data.alicloud_ecs_elasticity_assurances.default.assurances.0.allocated_resources.0.instance_type
  		owners        = "system"
	}

	resource "alicloud_vpc" "vpc" {
  		vpc_name   = var.name
  		cidr_block = "172.16.0.0/12"
	}

	resource "alicloud_vswitch" "vswitch" {
  		vswitch_name = var.name
  		cidr_block   = "172.16.0.0/21"
  		zone_id      = data.alicloud_ecs_elasticity_assurances.default.assurances.0.allocated_resources.0.zone_id
  		vpc_id       = alicloud_vpc.vpc.id
	}

	resource "alicloud_security_group" "default" {
  		name   = var.name
  		vpc_id = alicloud_vpc.vpc.id
	}
`, name)
}

var tags0 = map[string]string{
	"foo":    "foo",
	"bar":    "bar",
	"tf":     "tf",
	"create": "create",
	"tag0":   "value0",
	"tag1":   "value1",
	"tag2":   "value2",
	"tag3":   "value3",
	"tag4":   "value4",
	"tag5":   "value5",
	"tag6":   "value6",
	"tag7":   "value7",
	"tag8":   "value8",
	"tag9":   "value9",
	"tag10":  "value10",
	"tag11":  "value11",
	"tag12":  "value12",
	"tag13":  "value13",
	"tag14":  "value14",
	"tag15":  "value15",
	"tag16":  "value16",
	"tag17":  "value17",
	"tag18":  "value18",
	"tag19":  "value19",
	"tagA":   "valueA",
	"tagB":   "valueB",
	"tagC":   "valueC",
	"tagD":   "valueD",
	"tagE":   "valueE",
	"tagF":   "valueF",
	"tagG":   "valueG",
	"tagH":   "valueH",
	"tagI":   "valueI",
	"tagJ":   "valueJ",
	"tagK":   "valueK",
	"tagL":   "valueL",
	"tagM":   "valueM",
	"tagN":   "valueN",
	"tagO":   "valueO",
	"tagP":   "valueP",
	"tagQ":   "valueQ",
	"tagR":   "valueR",
	"tagS":   "valueS",
	"tagT":   "valueT",
	"tagU":   "valueU",
	"tagV":   "valueV",
	"tagW":   "valueW",
	"tagX":   "valueX",
	"tagY":   "valueY",
	"tagZ":   "valueZ",
}
