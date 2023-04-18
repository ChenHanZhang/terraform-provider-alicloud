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

type EipServiceV2 struct {
	client *connectivity.AliyunClient
}

// SetResourceTags <<< Encapsulated tag function for Eip.
func (s *EipServiceV2) SetResourceTags(d *schema.ResourceData, resourceType string) error {
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
			conn, err = client.NewEipClient()
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
			conn, err = client.NewEipClient()
			if err != nil {
				return WrapError(err)
			}
			request = make(map[string]interface{})

			request["ResourceId.1"] = d.Id()
			request["RegionId"] = client.RegionId

			count := 1
			for key, value := range added {
				request[fmt.Sprintf("Tag.%d.Key", count)] = key
				request[fmt.Sprintf("Tag.%d.Value", count)] = value
				count++
			}

			request["ResourceType"] = resourceType
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

// SetResourceTags >>> tag function encapsulated.// DescribeEipAddress <<< Encapsulated get interface for Eip Address.
func (s *EipServiceV2) DescribeEipAddress(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeEipAddressDescribeEipAddressesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_eip_address EipServiceV2.DescribeEipAddress Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)
	object1, err := s.describeEipAddressListTagResourcesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_eip_address EipServiceV2.DescribeEipAddress Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object1)
	checkValue00, _ := jsonpath.Get("$.hd_monitor_status", objectSearch)
	if checkValue00 == "ON" {

		object2, err := s.describeEipAddressDescribeHighDefinitionMonitorLogAttributeApi(id)
		if err != nil {
			if NotFoundError(err) {
				log.Printf("[DEBUG] Resource alicloud_eip_address EipServiceV2.DescribeEipAddress Failed!!! %s", err)
				return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
			}
			return object, WrapError(err)
		}
		objectSearch = MergeMaps(objectSearch, object2)
	}

	return objectSearch, nil
}

func (s *EipServiceV2) DescribeEipAddressPrimary(id string) (object map[string]interface{}, err error) {
	object, err = s.describeEipAddressDescribeEipAddressesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_eip_address EipServiceV2.DescribeEipAddressPrimary Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	return object, nil
}

func (s *EipServiceV2) describeEipAddressDescribeEipAddressesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "DescribeEipAddresses"
	conn, err := client.NewEipClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["AllocationId"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("Address", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.EipAddresses.EipAddress[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.EipAddresses.EipAddress[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("Address", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["address_name"] = objectRaw["Name"]
	instanceMap["bandwidth"] = objectRaw["Bandwidth"]
	instanceMap["bandwidth_package_bandwidth"] = objectRaw["BandwidthPackageBandwidth"]
	instanceMap["bandwidth_package_id"] = objectRaw["BandwidthPackageId"]
	instanceMap["bandwidth_package_type"] = objectRaw["BandwidthPackageType"]
	instanceMap["business_status"] = objectRaw["BusinessStatus"]
	instanceMap["create_time"] = objectRaw["AllocationTime"]
	instanceMap["deletion_protection"] = objectRaw["DeletionProtection"]
	instanceMap["description"] = objectRaw["Description"]
	instanceMap["eip_bandwidth"] = objectRaw["EipBandwidth"]
	instanceMap["expired_time"] = objectRaw["ExpiredTime"]
	instanceMap["has_reservation_data"] = objectRaw["HasReservationData"]
	instanceMap["hd_monitor_status"] = objectRaw["HDMonitorStatus"]
	instanceMap["instance_id"] = objectRaw["InstanceId"]
	instanceMap["instance_region_id"] = objectRaw["InstanceRegionId"]
	instanceMap["internet_charge_type"] = objectRaw["InternetChargeType"]
	instanceMap["ip_address"] = objectRaw["IpAddress"]
	instanceMap["isp"] = objectRaw["ISP"]
	instanceMap["netmode"] = objectRaw["Netmode"]
	instanceMap["payment_type"] = objectRaw["ChargeType"]
	instanceMap["public_ip_address_pool_id"] = objectRaw["PublicIpAddressPoolId"]
	instanceMap["reservation_active_time"] = objectRaw["ReservationActiveTime"]
	instanceMap["reservation_bandwidth"] = objectRaw["ReservationBandwidth"]
	instanceMap["reservation_internet_charge_type"] = objectRaw["ReservationInternetChargeType"]
	instanceMap["reservation_order_type"] = objectRaw["ReservationOrderType"]
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["second_limited"] = objectRaw["SecondLimited"]
	instanceMap["segment_instance_id"] = objectRaw["SegmentInstanceId"]
	instanceMap["service_managed"] = objectRaw["ServiceManaged"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["vpc_id"] = objectRaw["VpcId"]
	instanceMap["zone"] = objectRaw["Zone"]
	instanceMap["allocation_id"] = objectRaw["AllocationId"]
	instanceMap["instance_type"] = objectRaw["InstanceType"]
	instanceMap["region_id"] = objectRaw["RegionId"]

	{
		lockReason3Raw, _ := jsonpath.Get("$.OperationLocks.LockReason", objectRaw)
		operationLocksMaps := make([]map[string]interface{}, 0)
		if lockReason3Raw != nil {
			for _, lockReasonChild1Raw := range lockReason3Raw.([]interface{}) {
				operationLocksMap := make(map[string]interface{})
				lockReasonChild1Raw := lockReasonChild1Raw.(map[string]interface{})
				operationLocksMap["lock_reason"] = lockReasonChild1Raw["LockReason"]
				operationLocksMaps = append(operationLocksMaps, operationLocksMap)
			}
		}
		instanceMap["operation_locks"] = operationLocksMaps
	}

	securityProtectionType1Raw, _ := jsonpath.Get("$.SecurityProtectionTypes.SecurityProtectionType", objectRaw)
	instanceMap["security_protection_types"] = securityProtectionType1Raw

	{
		tag1Raw, _ := jsonpath.Get("$.Tags.Tag", objectRaw)
		tagsMaps := make([]map[string]interface{}, 0)
		if tag1Raw != nil {
			for _, tagChild1Raw := range tag1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagChild1Raw := tagChild1Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagChild1Raw["Key"]
				tagsMap["tag_value"] = tagChild1Raw["Value"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *EipServiceV2) describeEipAddressListTagResourcesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTagResources"
	conn, err := client.NewEipClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId.1"] = id
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "EIP"
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
			return object, WrapErrorf(Error(GetNotFoundMessage("Address", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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

func (s *EipServiceV2) describeEipAddressDescribeHighDefinitionMonitorLogAttributeApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "DescribeHighDefinitionMonitorLogAttribute"
	conn, err := client.NewEipClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["InstanceId"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("Address", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	instanceMap["hd_monitor_log_project"] = objectRaw["LogProject"]
	instanceMap["hd_monitor_log_store"] = objectRaw["LogStore"]
	instanceMap["allocation_id"] = objectRaw["InstanceId"]

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *EipServiceV2) EipAddressStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEipAddressPrimary(id)
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

// DescribeEipAddress >>> Encapsulated.
// DescribeEipSegmentAddress <<< Encapsulated get interface for Eip SegmentAddress.
func (s *EipServiceV2) DescribeEipSegmentAddress(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeEipSegmentAddressDescribeEipSegmentApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_eip_segment_address EipServiceV2.DescribeEipSegmentAddress Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)

	return objectSearch, nil
}

func (s *EipServiceV2) describeEipSegmentAddressDescribeEipSegmentApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "DescribeEipSegment"
	conn, err := client.NewEipClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["SegmentInstanceId"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("SegmentAddress", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.EipSegments.EipSegment[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.EipSegments.EipSegment[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("SegmentAddress", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["create_time"] = objectRaw["CreationTime"]
	instanceMap["descritpion"] = objectRaw["Descritpion"]
	instanceMap["ip_count"] = objectRaw["IpCount"]
	instanceMap["segment"] = objectRaw["Segment"]
	instanceMap["segment_address_name"] = objectRaw["Name"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["zone"] = objectRaw["Zone"]
	instanceMap["region_id"] = objectRaw["RegionId"]
	instanceMap["segment_instance_id"] = objectRaw["InstanceId"]

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *EipServiceV2) EipSegmentAddressStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEipSegmentAddress(id)
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

// DescribeEipSegmentAddress >>> Encapsulated.
