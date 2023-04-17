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

	return objectSearch, nil
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

func (s *VpcServiceV2) VpcHavipStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeVpcHavip(id)
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
			for i, key := range removedTagKeys {
				request[fmt.Sprintf("TagKey.%d", i+1)] = key
			}

			if v, ok := d.GetOkExists("all"); ok {
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
