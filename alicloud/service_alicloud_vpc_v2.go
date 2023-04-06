package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
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

	associatedEipAddresseRaw, _ := jsonpath.Get("$.AssociatedEipAddresses.associatedEipAddresse", objectRaw)
	instanceMap["associated_eip_addresses"] = associatedEipAddresseRaw

	associatedInstanceRaw, _ := jsonpath.Get("$.AssociatedInstances.associatedInstance", objectRaw)
	instanceMap["associated_instances"] = associatedInstanceRaw

	{
		tagRaw, _ := jsonpath.Get("$.Tags.Tag", objectRaw)
		tagsMaps := make([]map[string]interface{}, 0)
		if tagRaw != nil {
			for _, tagChildRaw := range tagRaw.([]interface{}) {
				tagsMap := make(map[string]interface{})
				tagChildRaw := tagChildRaw.(map[string]interface{})
				tagsMap["tag_key"] = tagChildRaw["Key"]
				tagsMap["tag_value"] = tagChildRaw["Value"]

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
