package alicloud

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tidwall/sjson"
)

type HologramServiceV2 struct {
	client *connectivity.AliyunClient
}

// DescribeHologramInstance <<< Encapsulated get interface for Hologram Instance.

func (s *HologramServiceV2) DescribeHologramInstance(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var header map[string]*string
	instanceId := id
	request = make(map[string]interface{})
	query = make(map[string]*string)
	header = make(map[string]*string)

	action := fmt.Sprintf("/api/v1/instances/%s", instanceId)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = client.RoaGet("Hologram", "2022-06-01", action, query, header, nil)

		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)
	if err != nil {
		if IsExpectedErrors(err, []string{"resource not exists failed"}) {
			return object, WrapErrorf(NotFoundErr("Instance", id), NotFoundMsg, response)
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.Instance", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Instance", response)
	}

	return v.(map[string]interface{}), nil
}

func (s *HologramServiceV2) HologramInstanceStateRefreshFunc(id string, field string, failStates []string) resource.StateRefreshFunc {
	return s.HologramInstanceStateRefreshFuncWithApi(id, field, failStates, s.DescribeHologramInstance)
}

func (s *HologramServiceV2) HologramInstanceStateRefreshFuncWithApi(id string, field string, failStates []string, call func(id string) (map[string]interface{}, error)) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := call(id)
		if err != nil {
			if NotFoundError(err) {
				return object, "", nil
			}
			return nil, "", WrapError(err)
		}
		object["InstanceChargeType"] = convertHologramInstanceInstanceInstanceChargeTypeResponse(object["InstanceChargeType"])
		v, err := jsonpath.Get(field, object)
		currentStatus := fmt.Sprint(v)

		if strings.HasPrefix(field, "#") {
			v, _ := jsonpath.Get(strings.TrimPrefix(field, "#"), object)
			if v != nil {
				currentStatus = "#CHECKSET"
			}
		}

		for _, failState := range failStates {
			if currentStatus == failState {
				return object, currentStatus, WrapError(Error(FailedToReachTargetStatus, currentStatus))
			}
		}
		return object, currentStatus, nil
	}
}

// DescribeHologramInstance >>> Encapsulated.

// SetResourceTags <<< Encapsulated tag function for Hologram.
func (s *HologramServiceV2) SetResourceTags(d *schema.ResourceData, resourceType string) error {
	if d.HasChange("tags") {
		var action string
		var err error
		client := s.client
		var request map[string]interface{}
		var response map[string]interface{}
		header := make(map[string]*string)
		query := make(map[string]*string)
		body := make(map[string]interface{})

		added, removed := parsingTags(d)
		removedTagKeys := make([]string, 0)
		for _, v := range removed {
			if !ignoredTags(v, "") {
				removedTagKeys = append(removedTagKeys, v)
			}
		}
		if len(removedTagKeys) > 0 {
			action = fmt.Sprintf("/api/v1/tag/unTag")
			request = make(map[string]interface{})
			query = make(map[string]*string)
			body = make(map[string]interface{})
			query["RegionId"] = StringPointer(client.RegionId)
			request["tagKeys"] = convertListStringToListInterface(removedTagKeys)
			jsonString := convertObjectToJsonString(request)
			jsonString, _ = sjson.Set(jsonString, "resourceIds.0", d.Id())
			_ = json.Unmarshal([]byte(jsonString), &request)

			body = request
			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = client.RoaPost("Hologram", "2022-06-01", action, query, header, body, true)
				if err != nil {
					if NeedRetry(err) {
						wait()
						return resource.RetryableError(err)
					}
					return resource.NonRetryableError(err)
				}
				return nil
			})
			addDebug(action, response, request)
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}

		}

		if len(added) > 0 {
			action = fmt.Sprintf("/api/v1/tag")
			request = make(map[string]interface{})
			query = make(map[string]*string)
			body = make(map[string]interface{})
			query["RegionId"] = StringPointer(client.RegionId)
			count := 1
			tagsMaps := make([]map[string]interface{}, 0)
			for key, value := range added {
				tagsMap := make(map[string]interface{})
				tagsMap["value"] = value
				tagsMap["key"] = key
				tagsMaps = append(tagsMaps, tagsMap)
				count++
			}
			request["tags"] = tagsMaps

			jsonString := convertObjectToJsonString(request)
			jsonString, _ = sjson.Set(jsonString, "resourceIds.0", d.Id())
			_ = json.Unmarshal([]byte(jsonString), &request)

			body = request
			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = client.RoaPost("Hologram", "2022-06-01", action, query, header, body, true)
				if err != nil {
					if NeedRetry(err) {
						wait()
						return resource.RetryableError(err)
					}
					return resource.NonRetryableError(err)
				}
				return nil
			})
			addDebug(action, response, request)
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}
			hologramServiceV2 := HologramServiceV2{client}
			stateConf := BuildStateConf([]string{}, []string{"#CHECKSET"}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, hologramServiceV2.HologramInstanceStateRefreshFunc(d.Id(), "#TagsChild", []string{}))
			if _, err := stateConf.WaitForState(); err != nil {
				return WrapErrorf(err, IdMsg, d.Id())
			}

		}
	}

	return nil
}

// SetResourceTags >>> tag function encapsulated.
