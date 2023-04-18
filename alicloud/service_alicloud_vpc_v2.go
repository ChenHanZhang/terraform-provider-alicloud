package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tidwall/sjson"
)

type VpcServiceV2 struct {
	client *connectivity.AliyunClient
}

// DescribeVpcHavip <<< Encapsulated get interface for Vpc Havip.
func (s *VpcServiceV2) DescribeVpcHavip(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcHavipDescribeHaVipsApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_havip VpcServiceV2.DescribeVpcHavip Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)
	object1, err := s.describeVpcHavipListTagResourcesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_havip VpcServiceV2.DescribeVpcHavip Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object1)

	return objectSearch, nil
}

func (s *VpcServiceV2) DescribeVpcHavipPrimary(id string) (object map[string]interface{}, err error) {
	object, err = s.describeVpcHavipDescribeHaVipsApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_havip VpcServiceV2.DescribeVpcHavipPrimary Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	return object, nil
}

func (s *VpcServiceV2) describeVpcHavipDescribeHaVipsApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "DescribeHaVips"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	jsonString := "{}"
	jsonString, _ = sjson.Set(jsonString, "Filter[0].Value[0]", id)
	jsonString, _ = sjson.Set(jsonString, "Filter[0].Key", "HaVipId")
	err = json.Unmarshal([]byte(jsonString), &request)
	if err != nil {
		return object, WrapError(err)
	}
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
		if IsExpectedErrors(err, []string{"InvalidHaVipId.NotFound"}) {
			return object, WrapErrorf(Error(GetNotFoundMessage("Havip", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.HaVips.HaVip[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.HaVips.HaVip[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("Havip", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["associated_instance_type"] = objectRaw["AssociatedInstanceType"]
	instanceMap["create_time"] = objectRaw["CreateTime"]
	instanceMap["description"] = objectRaw["Description"]
	instanceMap["ha_vip_name"] = objectRaw["Name"]
	instanceMap["ip_address"] = objectRaw["IpAddress"]
	instanceMap["master_instance_id"] = objectRaw["MasterInstanceId"]
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["vswitch_id"] = objectRaw["VSwitchId"]
	instanceMap["vpc_id"] = objectRaw["VpcId"]
	instanceMap["ha_vip_id"] = objectRaw["HaVipId"]
	instanceMap["region_id"] = objectRaw["RegionId"]

	associatedEipAddresse1Raw, _ := jsonpath.Get("$.AssociatedEipAddresses.associatedEipAddresse", objectRaw)
	instanceMap["associated_eip_addresses"] = associatedEipAddresse1Raw

	associatedInstance1Raw, _ := jsonpath.Get("$.AssociatedInstances.associatedInstance", objectRaw)
	instanceMap["associated_instances"] = associatedInstance1Raw

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

func (s *VpcServiceV2) describeVpcHavipListTagResourcesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTagResources"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId.1"] = id
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "HAVIP"
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
			return object, WrapErrorf(Error(GetNotFoundMessage("Havip", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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

func (s *VpcServiceV2) VpcHavipStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcHavipPrimary(id)
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

// DescribeVpcHavip >>> Encapsulated.

// SetResourceTags <<< Encapsulated tag function for Vpc.
func (s *VpcServiceV2) SetResourceTags(d *schema.ResourceData, resourceType string) error {
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
			conn, err = client.NewVpcClient()
			if err != nil {
				return WrapError(err)
			}
			request = make(map[string]interface{})

			request["ResourceId.1"] = d.Id()
			request["RegionId"] = client.RegionId

			request["ResourceType"] = resourceType
			if v, ok := d.GetOk("tags"); ok {
				jsonPathResult, err := jsonpath.Get("$.tag_key", v)
				if err != nil {
					return WrapError(err)
				}
				request["TagKey"] = jsonPathResult
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
			conn, err = client.NewVpcClient()
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

// SetResourceTags >>> tag function encapsulated.

// DescribeVpcIpv6EgressRule <<< Encapsulated get interface for Vpc Ipv6EgressRule.
func (s *VpcServiceV2) DescribeVpcIpv6EgressRule(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcIpv6EgressRuleDescribeIpv6EgressOnlyRulesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_ipv6_egress_rule VpcServiceV2.DescribeVpcIpv6EgressRule Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)

	return objectSearch, nil
}

func (s *VpcServiceV2) describeVpcIpv6EgressRuleDescribeIpv6EgressOnlyRulesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "DescribeIpv6EgressOnlyRules"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		err = WrapError(fmt.Errorf("invalid Resource Id %s. Expected parts' length %d, got %d", id, 2, len(parts)))
	}

	request["Ipv6GatewayId"] = parts[0]
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
			return object, WrapErrorf(Error(GetNotFoundMessage("Ipv6EgressRule", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.Ipv6EgressOnlyRules.Ipv6EgressOnlyRule[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Ipv6EgressOnlyRules.Ipv6EgressOnlyRule[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("Ipv6EgressRule", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["description"] = objectRaw["Description"]
	instanceMap["instance_id"] = objectRaw["InstanceId"]
	instanceMap["instance_type"] = objectRaw["InstanceType"]
	instanceMap["ipv6_egress_rule_name"] = objectRaw["Name"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["ipv6_egress_rule_id"] = objectRaw["Ipv6EgressOnlyRuleId"]

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcIpv6EgressRuleStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcIpv6EgressRule(id)
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

// DescribeVpcIpv6EgressRule >>> Encapsulated.
// DescribeVpcVswitch <<< Encapsulated get interface for Vpc Vswitch.
func (s *VpcServiceV2) DescribeVpcVswitch(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcVswitchDescribeVSwitchAttributesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vswitch VpcServiceV2.DescribeVpcVswitch Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	if object0["status"] == "" {
		return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
	}
	objectSearch = MergeMaps(objectSearch, object0)
	object1, err := s.describeVpcVswitchListTagResourcesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vswitch VpcServiceV2.DescribeVpcVswitch Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object1)

	return objectSearch, nil
}

func (s *VpcServiceV2) DescribeVpcVswitchPrimary(id string) (object map[string]interface{}, err error) {
	object, err = s.describeVpcVswitchDescribeVSwitchAttributesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vswitch VpcServiceV2.DescribeVpcVswitchPrimary Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	return object, nil
}

func (s *VpcServiceV2) describeVpcVswitchDescribeVSwitchAttributesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "DescribeVSwitchAttributes"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["VSwitchId"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("Vswitch", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	instanceMap["available_ip_address_count"] = objectRaw["AvailableIpAddressCount"]
	instanceMap["cidr_block"] = objectRaw["CidrBlock"]
	instanceMap["create_time"] = objectRaw["CreationTime"]
	instanceMap["description"] = objectRaw["Description"]
	instanceMap["ipv6_cidr_block"] = objectRaw["Ipv6CidrBlock"]
	instanceMap["ipv6_cidr_block_mask"] = objectRaw["Ipv6CidrBlock"]
	instanceMap["is_default"] = objectRaw["IsDefault"]
	instanceMap["network_acl_id"] = objectRaw["NetworkAclId"]
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["vswitch_name"] = objectRaw["VSwitchName"]
	instanceMap["vpc_id"] = objectRaw["VpcId"]
	instanceMap["zone_id"] = objectRaw["ZoneId"]
	instanceMap["vswitch_id"] = objectRaw["VSwitchId"]
	routeTable1RawObj, _ := jsonpath.Get("$.RouteTable", objectRaw)
	routeTable1Raw := make(map[string]interface{})
	if routeTable1RawObj != nil {
		routeTable1Raw = routeTable1RawObj.(map[string]interface{})
	}
	instanceMap["route_table_id"] = routeTable1Raw["RouteTableId"]

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

func (s *VpcServiceV2) describeVpcVswitchListTagResourcesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTagResources"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId.1"] = id
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "VSWITCH"
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
			return object, WrapErrorf(Error(GetNotFoundMessage("Vswitch", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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

func (s *VpcServiceV2) VpcVswitchStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcVswitchPrimary(id)
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

// DescribeVpcVswitch >>> Encapsulated.
// DescribeVpcPrefixList <<< Encapsulated get interface for Vpc PrefixList.
func (s *VpcServiceV2) DescribeVpcPrefixList(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcPrefixListListPrefixListsApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_prefix_list VpcServiceV2.DescribeVpcPrefixList Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)
	object1, err := s.describeVpcPrefixListGetVpcPrefixListEntriesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_prefix_list VpcServiceV2.DescribeVpcPrefixList Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object1)
	object2, err := s.describeVpcPrefixListGetVpcPrefixListAssociationsApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_prefix_list VpcServiceV2.DescribeVpcPrefixList Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object2)
	object3, err := s.describeVpcPrefixListListTagResourcesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_prefix_list VpcServiceV2.DescribeVpcPrefixList Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object3)

	return objectSearch, nil
}

func (s *VpcServiceV2) DescribeVpcPrefixListPrimary(id string) (object map[string]interface{}, err error) {
	object, err = s.describeVpcPrefixListListPrefixListsApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_prefix_list VpcServiceV2.DescribeVpcPrefixListPrimary Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	return object, nil
}

func (s *VpcServiceV2) describeVpcPrefixListListPrefixListsApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListPrefixLists"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["PrefixListIds.1"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("PrefixList", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.PrefixLists[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.PrefixLists[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("PrefixList", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["create_time"] = objectRaw["CreationTime"]
	instanceMap["ip_version"] = objectRaw["IpVersion"]
	instanceMap["max_entries"] = objectRaw["MaxEntries"]
	instanceMap["owner_id"] = objectRaw["OwnerId"]
	instanceMap["prefix_list_description"] = objectRaw["PrefixListDescription"]
	instanceMap["prefix_list_name"] = objectRaw["PrefixListName"]
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["share_type"] = objectRaw["ShareType"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["prefix_list_id"] = objectRaw["PrefixListId"]
	instanceMap["region_id"] = objectRaw["RegionId"]

	{
		tags1Raw := objectRaw["Tags"]
		tagsMaps := make([]map[string]interface{}, 0)
		if tags1Raw != nil {
			for _, tagsChild1Raw := range tags1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagsChild1Raw := tagsChild1Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagsChild1Raw["Key"]
				tagsMap["tag_value"] = tagsChild1Raw["Value"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) describeVpcPrefixListGetVpcPrefixListEntriesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "GetVpcPrefixListEntries"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["PrefixListId"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("PrefixList", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	prefixListEntry1RawObj, _ := jsonpath.Get("$.PrefixListEntry[*]", objectRaw)
	prefixListEntry1Raw := prefixListEntry1RawObj.([]interface{})

	prefixListEntryChild1Raw := make(map[string]interface{})
	if len(prefixListEntry1Raw) > 0 {
		prefixListEntryChild1Raw = prefixListEntry1Raw[0].(map[string]interface{})
	}
	instanceMap["prefix_list_id"] = prefixListEntryChild1Raw["PrefixListId"]

	{
		prefixListEntry1Raw := objectRaw["PrefixListEntry"]
		entriesMaps := make([]map[string]interface{}, 0)
		if prefixListEntry1Raw != nil {
			for _, prefixListEntryChild1Raw := range prefixListEntry1Raw.([]interface{}) {
				entriesMap := make(map[string]interface{})
				prefixListEntryChild1Raw := prefixListEntryChild1Raw.(map[string]interface{})
				entriesMap["cidr"] = prefixListEntryChild1Raw["Cidr"]
				entriesMap["description"] = prefixListEntryChild1Raw["Description"]
				entriesMaps = append(entriesMaps, entriesMap)
			}
		}
		instanceMap["entries"] = entriesMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) describeVpcPrefixListGetVpcPrefixListAssociationsApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "GetVpcPrefixListAssociations"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["PrefixListId"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("PrefixList", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
		prefixListAssociation1Raw := objectRaw["PrefixListAssociation"]
		prefixListAssociationMaps := make([]map[string]interface{}, 0)
		if prefixListAssociation1Raw != nil {
			for _, prefixListAssociationChild1Raw := range prefixListAssociation1Raw.([]interface{}) {
				prefixListAssociationMap := make(map[string]interface{})
				prefixListAssociationChild1Raw := prefixListAssociationChild1Raw.(map[string]interface{})
				prefixListAssociationMap["owner_id"] = prefixListAssociationChild1Raw["OwnerId"]
				prefixListAssociationMap["prefix_list_id"] = prefixListAssociationChild1Raw["PrefixListId"]
				prefixListAssociationMap["reason"] = prefixListAssociationChild1Raw["Reason"]
				prefixListAssociationMap["region_id"] = prefixListAssociationChild1Raw["RegionId"]
				prefixListAssociationMap["resource_id"] = prefixListAssociationChild1Raw["ResourceId"]
				prefixListAssociationMap["resource_type"] = prefixListAssociationChild1Raw["ResourceType"]
				prefixListAssociationMap["resource_uid"] = prefixListAssociationChild1Raw["ResourceUid"]
				prefixListAssociationMap["status"] = prefixListAssociationChild1Raw["Status"]
				prefixListAssociationMaps = append(prefixListAssociationMaps, prefixListAssociationMap)
			}
		}
		instanceMap["prefix_list_association"] = prefixListAssociationMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) describeVpcPrefixListListTagResourcesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTagResources"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId.1"] = id
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "PrefixList"
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
			return object, WrapErrorf(Error(GetNotFoundMessage("PrefixList", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	tagResource1RawObj, _ := jsonpath.Get("$.TagResources.TagResource[*]", objectRaw)
	tagResource1Raw := tagResource1RawObj.([]interface{})

	tagResourceChild2Raw := make(map[string]interface{})
	if len(tagResource1Raw) > 0 {
		tagResourceChild2Raw = tagResource1Raw[0].(map[string]interface{})
	}
	instanceMap["prefix_list_id"] = tagResourceChild2Raw["ResourceId"]

	{
		tagResource1Raw, _ := jsonpath.Get("$.TagResources.TagResource", objectRaw)
		tagsMaps := make([]map[string]interface{}, 0)
		if tagResource1Raw != nil {
			for _, tagResourceChild3Raw := range tagResource1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagResourceChild3Raw := tagResourceChild3Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagResourceChild3Raw["TagKey"]
				tagsMap["tag_value"] = tagResourceChild3Raw["TagValue"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcPrefixListStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcPrefixListPrimary(id)
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

// DescribeVpcPrefixList >>> Encapsulated.
// DescribeVpcPublicIpAddressPool <<< Encapsulated get interface for Vpc PublicIpAddressPool.
func (s *VpcServiceV2) DescribeVpcPublicIpAddressPool(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcPublicIpAddressPoolListPublicIpAddressPoolsApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_public_ip_address_pool VpcServiceV2.DescribeVpcPublicIpAddressPool Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)
	object1, err := s.describeVpcPublicIpAddressPoolListTagResourcesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_public_ip_address_pool VpcServiceV2.DescribeVpcPublicIpAddressPool Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object1)

	return objectSearch, nil
}

func (s *VpcServiceV2) DescribeVpcPublicIpAddressPoolPrimary(id string) (object map[string]interface{}, err error) {
	object, err = s.describeVpcPublicIpAddressPoolListPublicIpAddressPoolsApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_public_ip_address_pool VpcServiceV2.DescribeVpcPublicIpAddressPoolPrimary Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	return object, nil
}

func (s *VpcServiceV2) describeVpcPublicIpAddressPoolListPublicIpAddressPoolsApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListPublicIpAddressPools"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["PublicIpAddressPoolIds.1"] = id
	request["RegionId"] = client.RegionId
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationConflict", "IncorrectStatus.%s", "SystemBusy", "ServiceUnavailable", "OperationFailed.LastTokenProcessing", "LastTokenProcessing"}) || NeedRetry(err) {
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
			return object, WrapErrorf(Error(GetNotFoundMessage("PublicIpAddressPool", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.PublicIpAddressPoolList[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.PublicIpAddressPoolList[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("PublicIpAddressPool", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["create_time"] = objectRaw["CreationTime"]
	instanceMap["description"] = objectRaw["Description"]
	instanceMap["ip_address_remaining"] = objectRaw["IpAddressRemaining"]
	instanceMap["isp"] = objectRaw["Isp"]
	instanceMap["public_ip_address_pool_name"] = objectRaw["Name"]
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["total_ip_num"] = objectRaw["TotalIpNum"]
	instanceMap["used_ip_num"] = objectRaw["UsedIpNum"]
	instanceMap["public_ip_address_pool_id"] = objectRaw["PublicIpAddressPoolId"]
	instanceMap["region_id"] = objectRaw["RegionId"]

	{
		tags1Raw := objectRaw["Tags"]
		tagsMaps := make([]map[string]interface{}, 0)
		if tags1Raw != nil {
			for _, tagsChild1Raw := range tags1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagsChild1Raw := tagsChild1Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagsChild1Raw["Key"]
				tagsMap["tag_value"] = tagsChild1Raw["Value"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) describeVpcPublicIpAddressPoolListTagResourcesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTagResources"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId.1"] = id
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "PUBLICIPADDRESSPOOL"
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
			return object, WrapErrorf(Error(GetNotFoundMessage("PublicIpAddressPool", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	tagResource1RawObj, _ := jsonpath.Get("$.TagResources.TagResource[*]", objectRaw)
	tagResource1Raw := tagResource1RawObj.([]interface{})

	tagResourceChild2Raw := make(map[string]interface{})
	if len(tagResource1Raw) > 0 {
		tagResourceChild2Raw = tagResource1Raw[0].(map[string]interface{})
	}
	instanceMap["public_ip_address_pool_id"] = tagResourceChild2Raw["ResourceId"]

	{
		tagResource1Raw, _ := jsonpath.Get("$.TagResources.TagResource", objectRaw)
		tagsMaps := make([]map[string]interface{}, 0)
		if tagResource1Raw != nil {
			for _, tagResourceChild3Raw := range tagResource1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagResourceChild3Raw := tagResourceChild3Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagResourceChild3Raw["TagKey"]
				tagsMap["tag_value"] = tagResourceChild3Raw["TagValue"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcPublicIpAddressPoolStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcPublicIpAddressPoolPrimary(id)
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

// DescribeVpcPublicIpAddressPool >>> Encapsulated.
// DescribeVpcPublicIpAddressPoolCidrBlock <<< Encapsulated get interface for Vpc PublicIpAddressPoolCidrBlock.
func (s *VpcServiceV2) DescribeVpcPublicIpAddressPoolCidrBlock(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcPublicIpAddressPoolCidrBlockListPublicIpAddressPoolCidrBlocksApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_public_ip_address_pool_cidr_block VpcServiceV2.DescribeVpcPublicIpAddressPoolCidrBlock Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)

	return objectSearch, nil
}

func (s *VpcServiceV2) describeVpcPublicIpAddressPoolCidrBlockListPublicIpAddressPoolCidrBlocksApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListPublicIpAddressPoolCidrBlocks"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		err = WrapError(fmt.Errorf("invalid Resource Id %s. Expected parts' length %d, got %d", id, 2, len(parts)))
	}

	request["CidrBlock"] = parts[1]
	request["PublicIpAddressPoolId"] = parts[0]
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
			return object, WrapErrorf(Error(GetNotFoundMessage("PublicIpAddressPoolCidrBlock", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.PublicIpPoolCidrBlockList[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.PublicIpPoolCidrBlockList[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("PublicIpAddressPoolCidrBlock", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["create_time"] = objectRaw["CreationTime"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["total_ip_num"] = objectRaw["TotalIpNum"]
	instanceMap["used_ip_num"] = objectRaw["UsedIpNum"]
	instanceMap["cidr_block"] = objectRaw["CidrBlock"]
	instanceMap["public_ip_address_pool_id"] = objectRaw["PublicIpAddressPoolId"]

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcPublicIpAddressPoolCidrBlockStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcPublicIpAddressPoolCidrBlock(id)
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

// DescribeVpcPublicIpAddressPoolCidrBlock >>> Encapsulated.
// DescribeVpcFlowLog <<< Encapsulated get interface for Vpc FlowLog.
func (s *VpcServiceV2) DescribeVpcFlowLog(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcFlowLogDescribeFlowLogsApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_flow_log VpcServiceV2.DescribeVpcFlowLog Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)

	return objectSearch, nil
}

func (s *VpcServiceV2) describeVpcFlowLogDescribeFlowLogsApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "DescribeFlowLogs"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["FlowLogId"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("FlowLog", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.FlowLogs.FlowLog[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.FlowLogs.FlowLog[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("FlowLog", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["aggregation_interval"] = objectRaw["AggregationInterval"]
	instanceMap["business_status"] = objectRaw["BusinessStatus"]
	instanceMap["create_time"] = objectRaw["CreationTime"]
	instanceMap["description"] = objectRaw["Description"]
	instanceMap["flow_log_name"] = objectRaw["FlowLogName"]
	instanceMap["log_store_name"] = objectRaw["LogStoreName"]
	instanceMap["project_name"] = objectRaw["ProjectName"]
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["resource_id"] = objectRaw["ResourceId"]
	instanceMap["resource_type"] = objectRaw["ResourceType"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["traffic_type"] = objectRaw["TrafficType"]
	instanceMap["flow_log_id"] = objectRaw["FlowLogId"]
	instanceMap["region_id"] = objectRaw["RegionId"]

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

	trafficPathList1Raw, _ := jsonpath.Get("$.TrafficPath.TrafficPathList", objectRaw)
	instanceMap["traffic_path"] = trafficPathList1Raw

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcFlowLogStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcFlowLog(id)
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

// DescribeVpcFlowLog >>> Encapsulated.
// DescribeVpcIpv4Gateway <<< Encapsulated get interface for Vpc Ipv4Gateway.
func (s *VpcServiceV2) DescribeVpcIpv4Gateway(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcIpv4GatewayGetIpv4GatewayAttributeApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_ipv4_gateway VpcServiceV2.DescribeVpcIpv4Gateway Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)
	object1, err := s.describeVpcIpv4GatewayListTagResourcesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_ipv4_gateway VpcServiceV2.DescribeVpcIpv4Gateway Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object1)

	return objectSearch, nil
}

func (s *VpcServiceV2) DescribeVpcIpv4GatewayPrimary(id string) (object map[string]interface{}, err error) {
	object, err = s.describeVpcIpv4GatewayGetIpv4GatewayAttributeApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_ipv4_gateway VpcServiceV2.DescribeVpcIpv4GatewayPrimary Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	return object, nil
}

func (s *VpcServiceV2) describeVpcIpv4GatewayGetIpv4GatewayAttributeApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "GetIpv4GatewayAttribute"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["Ipv4GatewayId"] = id
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
		if IsExpectedErrors(err, []string{"ResourceNotFound.Ipv4Gateway"}) {
			return object, WrapErrorf(Error(GetNotFoundMessage("Ipv4Gateway", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	instanceMap["create_time"] = objectRaw["CreateTime"]
	instanceMap["enabled"] = objectRaw["Enabled"]
	instanceMap["ipv4_gateway_description"] = objectRaw["Ipv4GatewayDescription"]
	instanceMap["ipv4_gateway_name"] = objectRaw["Ipv4GatewayName"]
	instanceMap["ipv4_gateway_route_table_id"] = objectRaw["Ipv4GatewayRouteTableId"]
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["vpc_id"] = objectRaw["VpcId"]
	instanceMap["ipv4_gateway_id"] = objectRaw["Ipv4GatewayId"]

	{
		tags1Raw := objectRaw["Tags"]
		tagsMaps := make([]map[string]interface{}, 0)
		if tags1Raw != nil {
			for _, tagsChild1Raw := range tags1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagsChild1Raw := tagsChild1Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagsChild1Raw["Key"]
				tagsMap["tag_value"] = tagsChild1Raw["Value"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) describeVpcIpv4GatewayListTagResourcesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTagResources"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId.1"] = id
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "IPV4GATEWAY"
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
			return object, WrapErrorf(Error(GetNotFoundMessage("Ipv4Gateway", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	tagResource1RawObj, _ := jsonpath.Get("$.TagResources.TagResource[*]", objectRaw)
	tagResource1Raw := tagResource1RawObj.([]interface{})

	tagResourceChild2Raw := make(map[string]interface{})
	if len(tagResource1Raw) > 0 {
		tagResourceChild2Raw = tagResource1Raw[0].(map[string]interface{})
	}
	instanceMap["ipv4_gateway_id"] = tagResourceChild2Raw["ResourceId"]

	{
		tagResource1Raw, _ := jsonpath.Get("$.TagResources.TagResource", objectRaw)
		tagsMaps := make([]map[string]interface{}, 0)
		if tagResource1Raw != nil {
			for _, tagResourceChild3Raw := range tagResource1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagResourceChild3Raw := tagResourceChild3Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagResourceChild3Raw["TagKey"]
				tagsMap["tag_value"] = tagResourceChild3Raw["TagValue"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcIpv4GatewayStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcIpv4GatewayPrimary(id)
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

// DescribeVpcIpv4Gateway >>> Encapsulated.
// DescribeVpcIpv6Gateway <<< Encapsulated get interface for Vpc Ipv6Gateway.
func (s *VpcServiceV2) DescribeVpcIpv6Gateway(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcIpv6GatewayDescribeIpv6GatewayAttributeApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_ipv6_gateway VpcServiceV2.DescribeVpcIpv6Gateway Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)
	object1, err := s.describeVpcIpv6GatewayListTagResourcesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_ipv6_gateway VpcServiceV2.DescribeVpcIpv6Gateway Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object1)

	return objectSearch, nil
}

func (s *VpcServiceV2) DescribeVpcIpv6GatewayPrimary(id string) (object map[string]interface{}, err error) {
	object, err = s.describeVpcIpv6GatewayDescribeIpv6GatewayAttributeApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_ipv6_gateway VpcServiceV2.DescribeVpcIpv6GatewayPrimary Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	return object, nil
}

func (s *VpcServiceV2) describeVpcIpv6GatewayDescribeIpv6GatewayAttributeApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "DescribeIpv6GatewayAttribute"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["Ipv6GatewayId"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("Ipv6Gateway", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	instanceMap["business_status"] = objectRaw["BusinessStatus"]
	instanceMap["create_time"] = objectRaw["CreationTime"]
	instanceMap["description"] = objectRaw["Description"]
	instanceMap["expired_time"] = objectRaw["ExpiredTime"]
	instanceMap["instance_charge_type"] = objectRaw["InstanceChargeType"]
	instanceMap["ipv6_gateway_name"] = objectRaw["Name"]
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["vpc_id"] = objectRaw["VpcId"]
	instanceMap["ipv6_gateway_id"] = objectRaw["Ipv6GatewayId"]
	instanceMap["region_id"] = objectRaw["RegionId"]

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

func (s *VpcServiceV2) describeVpcIpv6GatewayListTagResourcesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTagResources"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId.1"] = id
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "IPV6GATEWAY"
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
			return object, WrapErrorf(Error(GetNotFoundMessage("Ipv6Gateway", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	tagResource1RawObj, _ := jsonpath.Get("$.TagResources.TagResource[*]", objectRaw)
	tagResource1Raw := tagResource1RawObj.([]interface{})

	tagResourceChild2Raw := make(map[string]interface{})
	if len(tagResource1Raw) > 0 {
		tagResourceChild2Raw = tagResource1Raw[0].(map[string]interface{})
	}
	instanceMap["ipv6_gateway_id"] = tagResourceChild2Raw["ResourceId"]

	{
		tagResource1Raw, _ := jsonpath.Get("$.TagResources.TagResource", objectRaw)
		tagsMaps := make([]map[string]interface{}, 0)
		if tagResource1Raw != nil {
			for _, tagResourceChild3Raw := range tagResource1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagResourceChild3Raw := tagResourceChild3Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagResourceChild3Raw["TagKey"]
				tagsMap["tag_value"] = tagResourceChild3Raw["TagValue"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcIpv6GatewayStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcIpv6GatewayPrimary(id)
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

// DescribeVpcIpv6Gateway >>> Encapsulated.
// DescribeVpcRouteTable <<< Encapsulated get interface for Vpc RouteTable.
func (s *VpcServiceV2) DescribeVpcRouteTable(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcRouteTableDescribeRouteTableListApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_route_table VpcServiceV2.DescribeVpcRouteTable Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)
	object1, err := s.describeVpcRouteTableListTagResourcesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_route_table VpcServiceV2.DescribeVpcRouteTable Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object1)

	return objectSearch, nil
}

func (s *VpcServiceV2) DescribeVpcRouteTablePrimary(id string) (object map[string]interface{}, err error) {
	object, err = s.describeVpcRouteTableDescribeRouteTableListApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_route_table VpcServiceV2.DescribeVpcRouteTablePrimary Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	return object, nil
}

func (s *VpcServiceV2) describeVpcRouteTableDescribeRouteTableListApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "DescribeRouteTableList"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["RouteTableId"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("RouteTable", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.RouterTableList.RouterTableListType[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.RouterTableList.RouterTableListType[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("RouteTable", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["associate_type"] = objectRaw["AssociateType"]
	instanceMap["create_time"] = objectRaw["CreationTime"]
	instanceMap["description"] = objectRaw["Description"]
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["route_table_name"] = objectRaw["RouteTableName"]
	instanceMap["route_table_type"] = objectRaw["RouteTableType"]
	instanceMap["router_id"] = objectRaw["RouterId"]
	instanceMap["router_type"] = objectRaw["RouterType"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["vpc_id"] = objectRaw["VpcId"]
	instanceMap["route_table_id"] = objectRaw["RouteTableId"]

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

	vSwitchId1Raw, _ := jsonpath.Get("$.VSwitchIds.VSwitchId", objectRaw)
	instanceMap["vswitch_ids"] = vSwitchId1Raw

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) describeVpcRouteTableListTagResourcesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTagResources"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId.1"] = id
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "ROUTETABLE"
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
			return object, WrapErrorf(Error(GetNotFoundMessage("RouteTable", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	tagResource1RawObj, _ := jsonpath.Get("$.TagResources.TagResource[*]", objectRaw)
	tagResource1Raw := tagResource1RawObj.([]interface{})

	tagResourceChild2Raw := make(map[string]interface{})
	if len(tagResource1Raw) > 0 {
		tagResourceChild2Raw = tagResource1Raw[0].(map[string]interface{})
	}
	instanceMap["route_table_id"] = tagResourceChild2Raw["ResourceId"]

	{
		tagResource1Raw, _ := jsonpath.Get("$.TagResources.TagResource", objectRaw)
		tagsMaps := make([]map[string]interface{}, 0)
		if tagResource1Raw != nil {
			for _, tagResourceChild3Raw := range tagResource1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagResourceChild3Raw := tagResourceChild3Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagResourceChild3Raw["TagKey"]
				tagsMap["tag_value"] = tagResourceChild3Raw["TagValue"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcRouteTableStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcRouteTablePrimary(id)
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

// DescribeVpcRouteTable >>> Encapsulated.
// DescribeVpcGatewayRouteTableAttachment <<< Encapsulated get interface for Vpc GatewayRouteTableAttachment.
func (s *VpcServiceV2) DescribeVpcGatewayRouteTableAttachment(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcGatewayRouteTableAttachmentGetIpv4GatewayAttributeApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_gateway_route_table_attachment VpcServiceV2.DescribeVpcGatewayRouteTableAttachment Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)

	return objectSearch, nil
}

func (s *VpcServiceV2) describeVpcGatewayRouteTableAttachmentGetIpv4GatewayAttributeApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "GetIpv4GatewayAttribute"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		err = WrapError(fmt.Errorf("invalid Resource Id %s. Expected parts' length %d, got %d", id, 2, len(parts)))
	}

	request["Ipv4GatewayId"] = parts[1]
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
			return object, WrapErrorf(Error(GetNotFoundMessage("GatewayRouteTableAttachment", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	instanceMap["create_time"] = objectRaw["CreateTime"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["ipv4_gateway_id"] = objectRaw["Ipv4GatewayId"]
	instanceMap["route_table_id"] = objectRaw["Ipv4GatewayRouteTableId"]

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcGatewayRouteTableAttachmentStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcGatewayRouteTableAttachment(id)
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

// DescribeVpcGatewayRouteTableAttachment >>> Encapsulated.
// DescribeVpcNetworkAcl <<< Encapsulated get interface for Vpc NetworkAcl.
func (s *VpcServiceV2) DescribeVpcNetworkAcl(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcNetworkAclDescribeNetworkAclAttributesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_network_acl VpcServiceV2.DescribeVpcNetworkAcl Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)
	object1, err := s.describeVpcNetworkAclListTagResourcesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_network_acl VpcServiceV2.DescribeVpcNetworkAcl Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object1)

	return objectSearch, nil
}

func (s *VpcServiceV2) DescribeVpcNetworkAclPrimary(id string) (object map[string]interface{}, err error) {
	object, err = s.describeVpcNetworkAclDescribeNetworkAclAttributesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_network_acl VpcServiceV2.DescribeVpcNetworkAclPrimary Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	return object, nil
}

func (s *VpcServiceV2) describeVpcNetworkAclDescribeNetworkAclAttributesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "DescribeNetworkAclAttributes"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["NetworkAclId"] = id
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
		if IsExpectedErrors(err, []string{"InvalidNetworkAcl.NotFound"}) {
			return object, WrapErrorf(Error(GetNotFoundMessage("NetworkAcl", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.NetworkAclAttribute", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.NetworkAclAttribute", response)
	}

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["create_time"] = objectRaw["CreationTime"]
	instanceMap["description"] = objectRaw["Description"]
	instanceMap["network_acl_name"] = objectRaw["NetworkAclName"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["vpc_id"] = objectRaw["VpcId"]
	instanceMap["network_acl_id"] = objectRaw["NetworkAclId"]
	instanceMap["region_id"] = objectRaw["RegionId"]

	{
		egressAclEntry1Raw, _ := jsonpath.Get("$.EgressAclEntries.EgressAclEntry", objectRaw)
		egressAclEntriesMaps := make([]map[string]interface{}, 0)
		if egressAclEntry1Raw != nil {
			for _, egressAclEntryChild1Raw := range egressAclEntry1Raw.([]interface{}) {
				egressAclEntriesMap := make(map[string]interface{})
				egressAclEntryChild1Raw := egressAclEntryChild1Raw.(map[string]interface{})
				egressAclEntriesMap["description"] = egressAclEntryChild1Raw["Description"]
				egressAclEntriesMap["destination_cidr_ip"] = egressAclEntryChild1Raw["DestinationCidrIp"]
				egressAclEntriesMap["network_acl_entry_name"] = egressAclEntryChild1Raw["NetworkAclEntryName"]
				egressAclEntriesMap["policy"] = egressAclEntryChild1Raw["Policy"]
				egressAclEntriesMap["port"] = egressAclEntryChild1Raw["Port"]
				egressAclEntriesMap["protocol"] = egressAclEntryChild1Raw["Protocol"]
				egressAclEntriesMaps = append(egressAclEntriesMaps, egressAclEntriesMap)
			}
		}
		instanceMap["egress_acl_entries"] = egressAclEntriesMaps
	}

	{
		ingressAclEntry1Raw, _ := jsonpath.Get("$.IngressAclEntries.IngressAclEntry", objectRaw)
		ingressAclEntriesMaps := make([]map[string]interface{}, 0)
		if ingressAclEntry1Raw != nil {
			for _, ingressAclEntryChild1Raw := range ingressAclEntry1Raw.([]interface{}) {
				ingressAclEntriesMap := make(map[string]interface{})
				ingressAclEntryChild1Raw := ingressAclEntryChild1Raw.(map[string]interface{})
				ingressAclEntriesMap["description"] = ingressAclEntryChild1Raw["Description"]
				ingressAclEntriesMap["network_acl_entry_name"] = ingressAclEntryChild1Raw["NetworkAclEntryName"]
				ingressAclEntriesMap["policy"] = ingressAclEntryChild1Raw["Policy"]
				ingressAclEntriesMap["port"] = ingressAclEntryChild1Raw["Port"]
				ingressAclEntriesMap["protocol"] = ingressAclEntryChild1Raw["Protocol"]
				ingressAclEntriesMap["source_cidr_ip"] = ingressAclEntryChild1Raw["SourceCidrIp"]
				ingressAclEntriesMaps = append(ingressAclEntriesMaps, ingressAclEntriesMap)
			}
		}
		instanceMap["ingress_acl_entries"] = ingressAclEntriesMaps
	}

	{
		resource1Raw, _ := jsonpath.Get("$.Resources.Resource", objectRaw)
		resourcesMaps := make([]map[string]interface{}, 0)
		if resource1Raw != nil {
			for _, resourceChild1Raw := range resource1Raw.([]interface{}) {
				resourcesMap := make(map[string]interface{})
				resourceChild1Raw := resourceChild1Raw.(map[string]interface{})
				resourcesMap["resource_id"] = resourceChild1Raw["ResourceId"]
				resourcesMap["resource_type"] = resourceChild1Raw["ResourceType"]
				resourcesMap["status"] = resourceChild1Raw["Status"]
				resourcesMaps = append(resourcesMaps, resourcesMap)
			}
		}
		instanceMap["resources"] = resourcesMaps
	}

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

func (s *VpcServiceV2) describeVpcNetworkAclListTagResourcesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTagResources"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId.1"] = id
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "NETWORKACL"
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
			return object, WrapErrorf(Error(GetNotFoundMessage("NetworkAcl", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	tagResource1RawObj, _ := jsonpath.Get("$.TagResources.TagResource[*]", objectRaw)
	tagResource1Raw := tagResource1RawObj.([]interface{})

	tagResourceChild2Raw := make(map[string]interface{})
	if len(tagResource1Raw) > 0 {
		tagResourceChild2Raw = tagResource1Raw[0].(map[string]interface{})
	}
	instanceMap["network_acl_id"] = tagResourceChild2Raw["ResourceId"]

	{
		tagResource1Raw, _ := jsonpath.Get("$.TagResources.TagResource", objectRaw)
		tagsMaps := make([]map[string]interface{}, 0)
		if tagResource1Raw != nil {
			for _, tagResourceChild3Raw := range tagResource1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagResourceChild3Raw := tagResourceChild3Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagResourceChild3Raw["TagKey"]
				tagsMap["tag_value"] = tagResourceChild3Raw["TagValue"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcNetworkAclStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcNetworkAclPrimary(id)
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

// DescribeVpcNetworkAcl >>> Encapsulated.
// DescribeVpcTrafficMirrorFilter <<< Encapsulated get interface for Vpc TrafficMirrorFilter.
func (s *VpcServiceV2) DescribeVpcTrafficMirrorFilter(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcTrafficMirrorFilterListTrafficMirrorFiltersApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_traffic_mirror_filter VpcServiceV2.DescribeVpcTrafficMirrorFilter Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)
	object1, err := s.describeVpcTrafficMirrorFilterListTagResourcesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_traffic_mirror_filter VpcServiceV2.DescribeVpcTrafficMirrorFilter Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object1)

	return objectSearch, nil
}

func (s *VpcServiceV2) DescribeVpcTrafficMirrorFilterPrimary(id string) (object map[string]interface{}, err error) {
	object, err = s.describeVpcTrafficMirrorFilterListTrafficMirrorFiltersApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_traffic_mirror_filter VpcServiceV2.DescribeVpcTrafficMirrorFilterPrimary Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	return object, nil
}

func (s *VpcServiceV2) describeVpcTrafficMirrorFilterListTrafficMirrorFiltersApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTrafficMirrorFilters"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["TrafficMirrorFilterIds.1"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("TrafficMirrorFilter", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.TrafficMirrorFilters[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.TrafficMirrorFilters[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("TrafficMirrorFilter", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["status"] = objectRaw["TrafficMirrorFilterStatus"]
	instanceMap["traffic_mirror_filter_description"] = objectRaw["TrafficMirrorFilterDescription"]
	instanceMap["traffic_mirror_filter_name"] = objectRaw["TrafficMirrorFilterName"]
	instanceMap["traffic_mirror_filter_id"] = objectRaw["TrafficMirrorFilterId"]

	{
		egressRules1Raw := objectRaw["EgressRules"]
		egressRulesMaps := make([]map[string]interface{}, 0)
		if egressRules1Raw != nil {
			for _, egressRulesChild1Raw := range egressRules1Raw.([]interface{}) {
				egressRulesMap := make(map[string]interface{})
				egressRulesChild1Raw := egressRulesChild1Raw.(map[string]interface{})
				egressRulesMap["action"] = egressRulesChild1Raw["Action"]
				egressRulesMap["destination_cidr_block"] = egressRulesChild1Raw["DestinationCidrBlock"]
				egressRulesMap["destination_port_range"] = egressRulesChild1Raw["DestinationPortRange"]
				egressRulesMap["priority"] = egressRulesChild1Raw["Priority"]
				egressRulesMap["protocol"] = egressRulesChild1Raw["Protocol"]
				egressRulesMap["source_cidr_block"] = egressRulesChild1Raw["SourceCidrBlock"]
				egressRulesMap["source_port_range"] = egressRulesChild1Raw["SourcePortRange"]
				egressRulesMap["traffic_mirror_filter_rule_status"] = egressRulesChild1Raw["TrafficMirrorFilterRuleStatus"]
				egressRulesMaps = append(egressRulesMaps, egressRulesMap)
			}
		}
		instanceMap["egress_rules"] = egressRulesMaps
	}

	{
		ingressRules1Raw := objectRaw["IngressRules"]
		ingressRulesMaps := make([]map[string]interface{}, 0)
		if ingressRules1Raw != nil {
			for _, ingressRulesChild1Raw := range ingressRules1Raw.([]interface{}) {
				ingressRulesMap := make(map[string]interface{})
				ingressRulesChild1Raw := ingressRulesChild1Raw.(map[string]interface{})
				ingressRulesMap["action"] = ingressRulesChild1Raw["Action"]
				ingressRulesMap["destination_cidr_block"] = ingressRulesChild1Raw["DestinationCidrBlock"]
				ingressRulesMap["destination_port_range"] = ingressRulesChild1Raw["DestinationPortRange"]
				ingressRulesMap["priority"] = ingressRulesChild1Raw["Priority"]
				ingressRulesMap["protocol"] = ingressRulesChild1Raw["Protocol"]
				ingressRulesMap["source_cidr_block"] = ingressRulesChild1Raw["SourceCidrBlock"]
				ingressRulesMap["source_port_range"] = ingressRulesChild1Raw["SourcePortRange"]
				ingressRulesMap["traffic_mirror_filter_rule_status"] = ingressRulesChild1Raw["TrafficMirrorFilterRuleStatus"]
				ingressRulesMaps = append(ingressRulesMaps, ingressRulesMap)
			}
		}
		instanceMap["ingress_rules"] = ingressRulesMaps
	}

	{
		tags1Raw := objectRaw["Tags"]
		tagsMaps := make([]map[string]interface{}, 0)
		if tags1Raw != nil {
			for _, tagsChild1Raw := range tags1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagsChild1Raw := tagsChild1Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagsChild1Raw["Key"]
				tagsMap["tag_value"] = tagsChild1Raw["Value"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) describeVpcTrafficMirrorFilterListTagResourcesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTagResources"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId.1"] = id
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "TRAFFICMIRRORFILTER"
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
			return object, WrapErrorf(Error(GetNotFoundMessage("TrafficMirrorFilter", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	tagResource1RawObj, _ := jsonpath.Get("$.TagResources.TagResource[*]", objectRaw)
	tagResource1Raw := tagResource1RawObj.([]interface{})

	tagResourceChild1Raw := make(map[string]interface{})
	if len(tagResource1Raw) > 0 {
		tagResourceChild1Raw = tagResource1Raw[0].(map[string]interface{})
	}
	instanceMap["traffic_mirror_filter_id"] = tagResourceChild1Raw["ResourceId"]

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

func (s *VpcServiceV2) VpcTrafficMirrorFilterStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcTrafficMirrorFilterPrimary(id)
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

// DescribeVpcTrafficMirrorFilter >>> Encapsulated.

// DescribeVpcTrafficMirrorSession <<< Encapsulated get interface for Vpc TrafficMirrorSession.
func (s *VpcServiceV2) DescribeVpcTrafficMirrorSession(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcTrafficMirrorSessionListTrafficMirrorSessionsApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_traffic_mirror_session VpcServiceV2.DescribeVpcTrafficMirrorSession Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)

	return objectSearch, nil
}

func (s *VpcServiceV2) describeVpcTrafficMirrorSessionListTrafficMirrorSessionsApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTrafficMirrorSessions"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["TrafficMirrorSessionIds.1"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("TrafficMirrorSession", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.TrafficMirrorSessions[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.TrafficMirrorSessions[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("TrafficMirrorSession", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["enabled"] = objectRaw["Enabled"]
	instanceMap["packet_length"] = objectRaw["PacketLength"]
	instanceMap["priority"] = objectRaw["Priority"]
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["status"] = objectRaw["TrafficMirrorSessionStatus"]
	instanceMap["traffic_mirror_filter_id"] = objectRaw["TrafficMirrorFilterId"]
	instanceMap["traffic_mirror_session_business_status"] = objectRaw["TrafficMirrorSessionBusinessStatus"]
	instanceMap["traffic_mirror_session_description"] = objectRaw["TrafficMirrorSessionDescription"]
	instanceMap["traffic_mirror_session_name"] = objectRaw["TrafficMirrorSessionName"]
	instanceMap["traffic_mirror_target_id"] = objectRaw["TrafficMirrorTargetId"]
	instanceMap["traffic_mirror_target_type"] = objectRaw["TrafficMirrorTargetType"]
	instanceMap["virtual_network_id"] = objectRaw["VirtualNetworkId"]
	instanceMap["traffic_mirror_session_id"] = objectRaw["TrafficMirrorSessionId"]

	{
		tags1Raw := objectRaw["Tags"]
		tagsMaps := make([]map[string]interface{}, 0)
		if tags1Raw != nil {
			for _, tagsChild1Raw := range tags1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagsChild1Raw := tagsChild1Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagsChild1Raw["Key"]
				tagsMap["tag_value"] = tagsChild1Raw["Value"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	trafficMirrorSourceIds1Raw := objectRaw["TrafficMirrorSourceIds"].([]interface{})
	instanceMap["traffic_mirror_source_ids"] = trafficMirrorSourceIds1Raw

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcTrafficMirrorSessionStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcTrafficMirrorSession(id)
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

// DescribeVpcTrafficMirrorSession >>> Encapsulated.
// DescribeVpcIpv6InternetBandwidth <<< Encapsulated get interface for Vpc Ipv6InternetBandwidth.
func (s *VpcServiceV2) DescribeVpcIpv6InternetBandwidth(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcIpv6InternetBandwidthDescribeIpv6AddressesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_ipv6_internet_bandwidth VpcServiceV2.DescribeVpcIpv6InternetBandwidth Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)

	return objectSearch, nil
}

func (s *VpcServiceV2) describeVpcIpv6InternetBandwidthDescribeIpv6AddressesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "DescribeIpv6Addresses"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["Ipv6InternetBandwidthId"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("Ipv6InternetBandwidth", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.Ipv6Addresses.Ipv6Address[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Ipv6Addresses.Ipv6Address[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("Ipv6InternetBandwidth", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["ipv6_address_id"] = objectRaw["Ipv6AddressId"]
	instanceMap["ipv6_gateway_id"] = objectRaw["Ipv6GatewayId"]
	ipv6InternetBandwidth1RawObj, _ := jsonpath.Get("$.Ipv6InternetBandwidth", objectRaw)
	ipv6InternetBandwidth1Raw := make(map[string]interface{})
	if ipv6InternetBandwidth1RawObj != nil {
		ipv6InternetBandwidth1Raw = ipv6InternetBandwidth1RawObj.(map[string]interface{})
	}
	instanceMap["bandwidth"] = ipv6InternetBandwidth1Raw["Bandwidth"]
	instanceMap["internet_charge_type"] = ipv6InternetBandwidth1Raw["InternetChargeType"]
	instanceMap["payment_type"] = ipv6InternetBandwidth1Raw["InstanceChargeType"]
	instanceMap["status"] = ipv6InternetBandwidth1Raw["BusinessStatus"]
	instanceMap["ipv6_internet_bandwidth_id"] = ipv6InternetBandwidth1Raw["Ipv6InternetBandwidthId"]

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcIpv6InternetBandwidthStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcIpv6InternetBandwidth(id)
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

// DescribeVpcIpv6InternetBandwidth >>> Encapsulated.
// DescribeVpcDhcpOptionsSet <<< Encapsulated get interface for Vpc DhcpOptionsSet.
func (s *VpcServiceV2) DescribeVpcDhcpOptionsSet(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcDhcpOptionsSetGetDhcpOptionsSetApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_dhcp_options_set VpcServiceV2.DescribeVpcDhcpOptionsSet Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)

	return objectSearch, nil
}

func (s *VpcServiceV2) describeVpcDhcpOptionsSetGetDhcpOptionsSetApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "GetDhcpOptionsSet"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["DhcpOptionsSetId"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("DhcpOptionsSet", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	instanceMap["dhcp_options_set_description"] = objectRaw["DhcpOptionsSetDescription"]
	instanceMap["dhcp_options_set_name"] = objectRaw["DhcpOptionsSetName"]
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["dhcp_options_set_id"] = objectRaw["DhcpOptionsSetId"]
	dhcpOptions1RawObj, _ := jsonpath.Get("$.DhcpOptions", objectRaw)
	dhcpOptions1Raw := make(map[string]interface{})
	if dhcpOptions1RawObj != nil {
		dhcpOptions1Raw = dhcpOptions1RawObj.(map[string]interface{})
	}
	instanceMap["domain_name"] = dhcpOptions1Raw["DomainName"]
	instanceMap["domain_name_servers"] = dhcpOptions1Raw["DomainNameServers"]
	instanceMap["ipv6_lease_time"] = dhcpOptions1Raw["Ipv6LeaseTime"]
	instanceMap["lease_time"] = dhcpOptions1Raw["LeaseTime"]

	{
		associateVpcs1Raw := objectRaw["AssociateVpcs"]
		associateVpcsMaps := make([]map[string]interface{}, 0)
		if associateVpcs1Raw != nil {
			for _, associateVpcsChild1Raw := range associateVpcs1Raw.([]interface{}) {
				associateVpcsMap := make(map[string]interface{})
				associateVpcsChild1Raw := associateVpcsChild1Raw.(map[string]interface{})
				associateVpcsMap["associate_status"] = associateVpcsChild1Raw["AssociateStatus"]
				associateVpcsMap["vpc_id"] = associateVpcsChild1Raw["VpcId"]
				associateVpcsMaps = append(associateVpcsMaps, associateVpcsMap)
			}
		}
		instanceMap["associate_vpcs"] = associateVpcsMaps
	}

	{
		tags1Raw := objectRaw["Tags"]
		tagsMaps := make([]map[string]interface{}, 0)
		if tags1Raw != nil {
			for _, tagsChild1Raw := range tags1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagsChild1Raw := tagsChild1Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagsChild1Raw["Key"]
				tagsMap["tag_value"] = tagsChild1Raw["Value"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcDhcpOptionsSetStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcDhcpOptionsSet(id)
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

// DescribeVpcDhcpOptionsSet >>> Encapsulated.
// DescribeVpcVpc <<< Encapsulated get interface for Vpc Vpc.
func (s *VpcServiceV2) DescribeVpcVpc(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcVpcDescribeVpcAttributeApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc VpcServiceV2.DescribeVpcVpc Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)
	object1, err := s.describeVpcVpcDescribeRouteTableListApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc VpcServiceV2.DescribeVpcVpc Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object1)
	object2, err := s.describeVpcVpcListTagResourcesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc VpcServiceV2.DescribeVpcVpc Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object2)

	return objectSearch, nil
}

func (s *VpcServiceV2) DescribeVpcVpcPrimary(id string) (object map[string]interface{}, err error) {
	object, err = s.describeVpcVpcDescribeVpcAttributeApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc VpcServiceV2.DescribeVpcVpcPrimary Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	return object, nil
}

func (s *VpcServiceV2) describeVpcVpcDescribeVpcAttributeApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "DescribeVpcAttribute"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["VpcId"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("Vpc", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	instanceMap["cidr_block"] = objectRaw["CidrBlock"]
	instanceMap["classic_link_enabled"] = objectRaw["ClassicLinkEnabled"]
	instanceMap["create_time"] = objectRaw["CreationTime"]
	instanceMap["description"] = objectRaw["Description"]
	instanceMap["dhcp_options_set_id"] = objectRaw["DhcpOptionsSetId"]
	instanceMap["ipv6_cidr_block"] = objectRaw["Ipv6CidrBlock"]
	instanceMap["is_default"] = objectRaw["IsDefault"]
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["router_id"] = objectRaw["VRouterId"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["vpc_name"] = objectRaw["VpcName"]
	instanceMap["region_id"] = objectRaw["RegionId"]
	instanceMap["vpc_id"] = objectRaw["VpcId"]

	{
		ipv6CidrBlock7Raw, _ := jsonpath.Get("$.Ipv6CidrBlocks.Ipv6CidrBlock", objectRaw)
		ipv6CidrBlocksMaps := make([]map[string]interface{}, 0)
		if ipv6CidrBlock7Raw != nil {
			for _, ipv6CidrBlockChild1Raw := range ipv6CidrBlock7Raw.([]interface{}) {
				ipv6CidrBlocksMap := make(map[string]interface{})
				ipv6CidrBlockChild1Raw := ipv6CidrBlockChild1Raw.(map[string]interface{})
				ipv6CidrBlocksMap["ipv6_cidr_block"] = ipv6CidrBlockChild1Raw["Ipv6CidrBlock"]
				ipv6CidrBlocksMap["ipv6_isp"] = ipv6CidrBlockChild1Raw["Ipv6Isp"]
				ipv6CidrBlocksMaps = append(ipv6CidrBlocksMaps, ipv6CidrBlocksMap)
			}
		}
		instanceMap["ipv6_cidr_blocks"] = ipv6CidrBlocksMaps
	}

	secondaryCidrBlock1Raw, _ := jsonpath.Get("$.SecondaryCidrBlocks.SecondaryCidrBlock", objectRaw)
	instanceMap["secondary_cidr_blocks"] = secondaryCidrBlock1Raw

	userCidr1Raw, _ := jsonpath.Get("$.UserCidrs.UserCidr", objectRaw)
	instanceMap["user_cidrs"] = userCidr1Raw

	vSwitchId1Raw, _ := jsonpath.Get("$.VSwitchIds.VSwitchId", objectRaw)
	instanceMap["vswitch_ids"] = vSwitchId1Raw

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) describeVpcVpcDescribeRouteTableListApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "DescribeRouteTableList"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["VpcId"] = id
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
			return object, WrapErrorf(Error(GetNotFoundMessage("Vpc", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.RouterTableList.RouterTableListType[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.RouterTableList.RouterTableListType[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("Vpc", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["route_table_id"] = objectRaw["RouteTableId"]
	instanceMap["router_id"] = objectRaw["RouterId"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["route_table_type"] = objectRaw["RouteTableType"]
	instanceMap["vpc_id"] = objectRaw["VpcId"]

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) describeVpcVpcListTagResourcesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTagResources"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId.1"] = id
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "VPC"
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
			return object, WrapErrorf(Error(GetNotFoundMessage("Vpc", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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

func (s *VpcServiceV2) VpcVpcStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcVpcPrimary(id)
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

// DescribeVpcVpc >>> Encapsulated.
// DescribeVpcPeerConnection <<< Encapsulated get interface for Vpc PeerConnection.
func (s *VpcServiceV2) DescribeVpcPeerConnection(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcPeerConnectionGetVpcPeerConnectionAttributeApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_peer_connection VpcServiceV2.DescribeVpcPeerConnection Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)
	object1, err := s.describeVpcPeerConnectionListTagResourcesApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_peer_connection VpcServiceV2.DescribeVpcPeerConnection Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object1)

	return objectSearch, nil
}

func (s *VpcServiceV2) DescribeVpcPeerConnectionPrimary(id string) (object map[string]interface{}, err error) {
	object, err = s.describeVpcPeerConnectionGetVpcPeerConnectionAttributeApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_peer_connection VpcServiceV2.DescribeVpcPeerConnectionPrimary Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	return object, nil
}

func (s *VpcServiceV2) describeVpcPeerConnectionGetVpcPeerConnectionAttributeApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "GetVpcPeerConnectionAttribute"
	conn, err := client.NewVpcpeerClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["InstanceId"] = id

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2022-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
		if IsExpectedErrors(err, []string{"ResourceNotFound.InstanceId"}) {
			return object, WrapErrorf(Error(GetNotFoundMessage("PeerConnection", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	instanceMap["accepting_ali_uid"] = objectRaw["AcceptingOwnerUid"]
	instanceMap["accepting_region_id"] = objectRaw["AcceptingRegionId"]
	instanceMap["bandwidth"] = objectRaw["Bandwidth"]
	instanceMap["biz_status"] = objectRaw["BizStatus"]
	instanceMap["create_time"] = objectRaw["GmtCreate"]
	instanceMap["description"] = objectRaw["Description"]
	instanceMap["expire_time"] = objectRaw["GmtExpired"]
	instanceMap["modify_time"] = objectRaw["GmtModified"]
	instanceMap["peer_connection_name"] = objectRaw["Name"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["peering_id"] = objectRaw["InstanceId"]
	instanceMap["region_id"] = objectRaw["RegionId"]
	acceptingVpc1RawObj, _ := jsonpath.Get("$.AcceptingVpc", objectRaw)
	acceptingVpc1Raw := make(map[string]interface{})
	if acceptingVpc1RawObj != nil {
		acceptingVpc1Raw = acceptingVpc1RawObj.(map[string]interface{})
	}
	instanceMap["accepting_vpc_id"] = acceptingVpc1Raw["VpcId"]
	vpc1RawObj, _ := jsonpath.Get("$.Vpc", objectRaw)
	vpc1Raw := make(map[string]interface{})
	if vpc1RawObj != nil {
		vpc1Raw = vpc1RawObj.(map[string]interface{})
	}
	instanceMap["vpc_id"] = vpc1Raw["VpcId"]

	{
		tags1Raw := objectRaw["Tags"]
		tagsMaps := make([]map[string]interface{}, 0)
		if tags1Raw != nil {
			for _, tagsChild1Raw := range tags1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagsChild1Raw := tagsChild1Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagsChild1Raw["Key"]
				tagsMap["tag_value"] = tagsChild1Raw["Value"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) describeVpcPeerConnectionListTagResourcesApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListTagResources"
	conn, err := client.NewVpcpeerClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId.1"] = id
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "PeerConnection"
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2022-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
			return object, WrapErrorf(Error(GetNotFoundMessage("PeerConnection", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
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
	tagResources1RawObj, _ := jsonpath.Get("$.TagResources[*]", objectRaw)
	tagResources1Raw := tagResources1RawObj.([]interface{})

	tagResourcesChild2Raw := make(map[string]interface{})
	if len(tagResources1Raw) > 0 {
		tagResourcesChild2Raw = tagResources1Raw[0].(map[string]interface{})
	}
	instanceMap["peering_id"] = tagResourcesChild2Raw["ResourceId"]

	{
		tagResources1Raw := objectRaw["TagResources"]
		tagsMaps := make([]map[string]interface{}, 0)
		if tagResources1Raw != nil {
			for _, tagResourcesChild3Raw := range tagResources1Raw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagResourcesChild3Raw := tagResourcesChild3Raw.(map[string]interface{})
				tagsMap["tag_key"] = tagResourcesChild3Raw["TagKey"]
				tagsMap["tag_value"] = tagResourcesChild3Raw["TagValue"]
				tagsMaps = append(tagsMaps, tagsMap)
			}
		}
		instanceMap["tags"] = tagsMaps
	}

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcPeerConnectionStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcPeerConnectionPrimary(id)
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

// DescribeVpcPeerConnection >>> Encapsulated.
// DescribeVpcVswitchCidrReservation <<< Encapsulated get interface for Vpc VswitchCidrReservation.
func (s *VpcServiceV2) DescribeVpcVswitchCidrReservation(id string) (object map[string]interface{}, err error) {
	objectSearch := make(map[string]interface{}, 0)
	object0, err := s.describeVpcVswitchCidrReservationListVSwitchCidrReservationsApi(id)
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_vswitch_cidr_reservation VpcServiceV2.DescribeVpcVswitchCidrReservation Failed!!! %s", err)
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapError(err)
	}
	objectSearch = MergeMaps(objectSearch, object0)

	return objectSearch, nil
}

func (s *VpcServiceV2) describeVpcVswitchCidrReservationListVSwitchCidrReservationsApi(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	action := "ListVSwitchCidrReservations"
	conn, err := client.NewVpcClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})

	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		err = WrapError(fmt.Errorf("invalid Resource Id %s. Expected parts' length %d, got %d", id, 2, len(parts)))
	}

	request["VSwitchCidrReservationIds.1"] = parts[1]
	request["VSwitchId"] = parts[0]
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
			return object, WrapErrorf(Error(GetNotFoundMessage("VswitchCidrReservation", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.VSwitchCidrReservations[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.VSwitchCidrReservations[*]", response)
	}
	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("VswitchCidrReservation", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
	}
	v = v.([]interface{})[0]

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["cdir_reservation_type"] = objectRaw["Type"]
	instanceMap["cidr_reservation_cidr"] = objectRaw["VSwitchCidrReservationCidr"]
	instanceMap["cidr_reservation_description"] = objectRaw["VSwitchCidrReservationDescription"]
	instanceMap["create_time"] = objectRaw["CreationTime"]
	instanceMap["ip_version"] = objectRaw["IpVersion"]
	instanceMap["resource_group_id"] = objectRaw["ResourceGroupId"]
	instanceMap["status"] = objectRaw["Status"]
	instanceMap["vswitch_cidr_reservation_name"] = objectRaw["VSwitchCidrReservationName"]
	instanceMap["vpc_instance_id"] = objectRaw["VpcId"]
	instanceMap["vswitch_cidr_reservation_id"] = objectRaw["VSwitchCidrReservationId"]
	instanceMap["vswitch_id"] = objectRaw["VSwitchId"]

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *VpcServiceV2) VpcVswitchCidrReservationStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcVswitchCidrReservation(id)
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

// DescribeVpcVswitchCidrReservation >>> Encapsulated.
