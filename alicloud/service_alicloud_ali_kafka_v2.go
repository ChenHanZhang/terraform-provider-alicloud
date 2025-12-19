// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/blues/jsonata-go"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

type AliKafkaServiceV2 struct {
	client *connectivity.AliyunClient
}

// DescribeAliKafkaAllowedIpAttachment <<< Encapsulated get interface for AliKafka AllowedIpAttachment.

func (s *AliKafkaServiceV2) DescribeAliKafkaAllowedIpAttachment(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	parts := strings.Split(id, ":")
	if len(parts) != 4 {
		err = WrapError(fmt.Errorf("invalid Resource Id %s. Expected parts' length %d, got %d", id, 4, len(parts)))
		return nil, err
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["RegionId"] = client.RegionId
	action := "GetAllowedIpList"

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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.AllowedList", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.AllowedList", response)
	}

	return v.(map[string]interface{}), nil
}

func (s *AliKafkaServiceV2) AliKafkaAllowedIpAttachmentStateRefreshFunc(id string, field string, failStates []string) resource.StateRefreshFunc {
	return s.AliKafkaAllowedIpAttachmentStateRefreshFuncWithApi(id, field, failStates, s.DescribeAliKafkaAllowedIpAttachment)
}

func (s *AliKafkaServiceV2) AliKafkaAllowedIpAttachmentStateRefreshFuncWithApi(id string, field string, failStates []string, call func(id string) (map[string]interface{}, error)) resource.StateRefreshFunc {
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
		if field == "$.InternetList[*].AllowedIpList[*]" {
			e := jsonata.MustCompile("$reduce($split($reduce($.AllowedList.VpcList.AllowedIpList, function($i, $j){$i & ',' & $j}), ',', 1), function($i, $j){$i & '' & $j})")
			v, _ = e.Eval(object)
			currentStatus = fmt.Sprint(v)
		}

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

// DescribeAliKafkaAllowedIpAttachment >>> Encapsulated.
