package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type CbwpServiceV2 struct {
	client *connectivity.AliyunClient
}

// SetResourceTags <<< Encapsulated tag function for Cbwp.
func (s *CbwpServiceV2) SetResourceTags(d *schema.ResourceData, resourceType string) error {
	if d.HasChange("tags") {
		var err error
		var action string
		var conn *rpc.Client
		client := s.client
		var request map[string]interface{}
		var response map[string]interface{}

		added, removed := parsingTags(d)
		removedTagKeys := make([]string, 0)
		for _, v := range removed {
			if !ignoredTags(v, "") {
				removedTagKeys = append(removedTagKeys, v)
			}
		}
		if len(removedTagKeys) > 0 {
			action = "UnTagResources"
			conn, err = client.NewCbwpClient()
			if err != nil {
				return WrapError(err)
			}
			request = make(map[string]interface{})

			request["ResourceId.1"] = d.Id()
			request["RegionId"] = client.RegionId

			for i, key := range removedTagKeys {
				request[fmt.Sprintf("TagKey.%d", i+1)] = key
			}

			request["ResourceType"] = resourceType
			if v, ok := d.GetOkExists("un_tag_all_tags"); ok {
				request["All"] = v
			}

			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
				if err != nil {
					if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
						wait()
						return resource.RetryableError(err)
					}
					return resource.NonRetryableError(err)
				}
				addDebug(action, response, request)
				return nil
			})
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}

		}

		if len(added) > 0 {
			action = "TagResources"
			conn, err = client.NewCbwpClient()
			if err != nil {
				return WrapError(err)
			}
			request = make(map[string]interface{})

			request["ResourceId.1"] = d.Id()
			request["RegionId"] = client.RegionId

			request["ResourceType"] = resourceType
			count := 1
			for key, value := range added {
				request[fmt.Sprintf("Tag.%d.Key", count)] = key
				request[fmt.Sprintf("Tag.%d.Value", count)] = value
				count++
			}

			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
				if err != nil {
					if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
						wait()
						return resource.RetryableError(err)
					}
					return resource.NonRetryableError(err)
				}
				addDebug(action, response, request)
				return nil
			})
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}

		}
	}

	return nil
}

// SetResourceTags >>> tag function encapsulated.// DescribeCbwpCommonBandwidthPackage <<< Encapsulated get interface for Cbwp CommonBandwidthPackage.
func (s *CbwpServiceV2) DescribeCbwpCommonBandwidthPackage(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeCbwpCommonBandwidthPackageDescribeCommonBandwidthPackagesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_common_bandwidth_package CbwpServiceV2.DescribeCbwpCommonBandwidthPackage Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)
	object1, err := s.describeCbwpCommonBandwidthPackageListTagResourcesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_common_bandwidth_package CbwpServiceV2.DescribeCbwpCommonBandwidthPackage Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object1)

	return objectSearch, nil
}

func (s *CbwpServiceV2) DescribeCbwpCommonBandwidthPackagePrimary(id string) (object map[string]interface{}, err error) {
	object, err = s.describeCbwpCommonBandwidthPackageDescribeCommonBandwidthPackagesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_common_bandwidth_package CbwpServiceV2.DescribeCbwpCommonBandwidthPackagePrimary Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	return object, nil
}

func (s *CbwpServiceV2) describeCbwpCommonBandwidthPackageDescribeCommonBandwidthPackagesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "DescribeCommonBandwidthPackages"
	conn, err := client.NewCbwpClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["BandwidthPackageId"] = id
	request["RegionId"] = client.RegionId
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})
	if err != nil {
		if IsExpectedErrors(err, []string{}) {
			return object, WrapErrorf(Error(GetNotFoundMessage("CommonBandwidthPackage", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.CommonBandwidthPackages.CommonBandwidthPackage[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.CommonBandwidthPackages.CommonBandwidthPackage[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("CommonBandwidthPackage", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["bandwidth"] = objectRaw["Bandwidth"]
	instanceMap["business_status"] = objectRaw["BusinessStatus"]
	instanceMap["common_bandwidth_package_name"] = objectRaw["Name"]
	instanceMap["create_time"] = objectRaw["CreationTime"]
	instanceMap["deletion_protection"] = objectRaw["DeletionProtection"]
	instanceMap["description"] = objectRaw["Description"]
	instanceMap["expired_time"] = objectRaw["ExpiredTime"]
	instanceMap["has_reservation_data"] = objectRaw["HasReservationData"]
	instanceMap["internet_charge_type"] = objectRaw["InternetChargeType"]
	instanceMap["isp"] = objectRaw["ISP"]
	instanceMap["payment_type"] = objectRaw["InstanceChargeType"]
	instanceMap["ratio"] = objectRaw["Ratio"]
	instanceMap["reservation_active_time"] = objectRaw["ReservationActiveTime"]
	instanceMap["reservation_bandwidth"] = objectRaw["ReservationBandwidth"]
	instanceMap["reservation_internet_charge_type"] = objectRaw["ReservationInternetChargeType"]
	instanceMap["reservation_order_type"] = objectRaw["ReservationOrderType"]
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["service_managed"] = objectRaw["ServiceManaged"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["common_bandwidth_package_id"] = objectRaw["BandwidthPackageId"]
	instanceMap["region_id"] = objectRaw["RegionId"]

	{
		publicIpAddresse1Raw, _ := jsonpath.Get("$.PublicIpAddresses.PublicIpAddresse", objectRaw)
		publicIpAddressesMaps := make([]map[string]interface{}, 0)
		if publicIpAddresse1Raw != nil {
			for _, publicIpAddresseChild1Raw := range publicIpAddresse1Raw.([]interface{}) {
				publicIpAddressesMap := make(map[string]interface{})
				publicIpAddresseChild1Raw := publicIpAddresseChild1Raw.(map[string]interface{})
				publicIpAddressesMap["allocation_id"] = publicIpAddresseChild1Raw["AllocationId"]
				publicIpAddressesMap["bandwidth_package_ip_relation_status"] = publicIpAddresseChild1Raw["BandwidthPackageIpRelationStatus"]
				publicIpAddressesMap["ip_address"] = publicIpAddresseChild1Raw["IpAddress"]
				publicIpAddressesMaps = append(publicIpAddressesMaps, publicIpAddressesMap)
			}
		}
		instanceMap["public_ip_addresses"] = publicIpAddressesMaps
	}

	securityProtectionType1Raw, _ := jsonpath.Get("$.SecurityProtectionTypes.SecurityProtectionType", objectRaw)
	instanceMap["security_protection_types"] = securityProtectionType1Raw

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *CbwpServiceV2) describeCbwpCommonBandwidthPackageListTagResourcesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTagResources"
	conn, err := client.NewCbwpClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId.1"] = id
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "bandwidthpackage"
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})
	if err != nil {
		if IsExpectedErrors(err, []string{}) {
			return object, WrapErrorf(Error(GetNotFoundMessage("CommonBandwidthPackage", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$", response)
	}

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["next_token"] = objectRaw["NextToken"]

	{
		tagResource1Raw, _ := jsonpath.Get("$.TagResources.TagResource", objectRaw)
		tagsMaps := make([]map[string]interface{}, 0)
		if tagResource1Raw != nil {
			for _, tagResourceChild1Raw := range tagResource1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagResourceChild1Raw := tagResourceChild1Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagResourceChild1Raw["TagKey"]
				tagsMap["tag_value"] = tagResourceChild1Raw["TagValue"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *CbwpServiceV2) CbwpCommonBandwidthPackageStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeCbwpCommonBandwidthPackagePrimary(id)
		if err != nil {
			if NotFoundError(err) {
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		currentStatus := object["status"]
		if _, ok := currentStatus.(string); !ok {
			return nil, "", nil
		}
		for _, failState := range failStates {
			if currentStatus.(string) == failState {
				return object, currentStatus.(string), WrapError(Error(FailedToReachTargetStatus, currentStatus.(string)))
			}
		}
		return object, currentStatus.(string), nil
	}
}

// DescribeCbwpCommonBandwidthPackage >>> Encapsulated.
