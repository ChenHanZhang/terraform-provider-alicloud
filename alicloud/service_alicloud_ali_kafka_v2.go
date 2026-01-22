// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

type AliKafkaServiceV2 struct {
	client *connectivity.AliyunClient
}

// DescribeAliKafkaScheduledScalingRule <<< Encapsulated get interface for AliKafka ScheduledScalingRule.

func (s *AliKafkaServiceV2) DescribeAliKafkaScheduledScalingRule(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		err = WrapError(fmt.Errorf("invalid Resource Id %s. Expected parts' length %d, got %d", id, 2, len(parts)))
		return nil, err
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["RegionId"] = client.RegionId
	action := "GetAutoScalingConfiguration"

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)

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
		if IsExpectedErrors(err, []string{"AUTH_RESOURCE_OWNER_ERROR"}) {
			return object, WrapErrorf(NotFoundErr("ScheduledScalingRule", id), NotFoundMsg, response)
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.Data.ScheduledScalingRules.ScheduledScalingRules[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Data.ScheduledScalingRules.ScheduledScalingRules[*]", response)
	}

	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(NotFoundErr("ScheduledScalingRule", id), NotFoundMsg, response)
	}

	result, _ := v.([]interface{})
	for _, v := range result {
		item := v.(map[string]interface{})
		if fmt.Sprint(item["RuleName"]) != parts[1] {
			continue
		}
		return item, nil
	}
	return object, WrapErrorf(NotFoundErr("ScheduledScalingRule", id), NotFoundMsg, response)
}

func (s *AliKafkaServiceV2) AliKafkaScheduledScalingRuleStateRefreshFunc(id string, field string, failStates []string) resource.StateRefreshFunc {
	return s.AliKafkaScheduledScalingRuleStateRefreshFuncWithApi(id, field, failStates, s.DescribeAliKafkaScheduledScalingRule)
}

func (s *AliKafkaServiceV2) AliKafkaScheduledScalingRuleStateRefreshFuncWithApi(id string, field string, failStates []string, call func(id string) (map[string]interface{}, error)) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := call(id)
		if err != nil {
			if NotFoundError(err) {
				return object, "", nil
			}
			return nil, "", WrapError(err)
		}
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

// DescribeAliKafkaScheduledScalingRule >>> Encapsulated.
