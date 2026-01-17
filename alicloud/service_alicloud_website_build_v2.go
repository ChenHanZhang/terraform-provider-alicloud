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

type WebsiteBuildServiceV2 struct {
	client *connectivity.AliyunClient
}

// DescribeWebsiteBuildAppInstance <<< Encapsulated get interface for WebsiteBuild AppInstance.

func (s *WebsiteBuildServiceV2) DescribeWebsiteBuildAppInstance(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["BizId"] = id

	action := "GetAppInstance"

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("WebsiteBuild", "2025-04-29", action, query, request, true)

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
		if IsExpectedErrors(err, []string{"AppInstance.NotFound.Exception"}) {
			return object, WrapErrorf(NotFoundErr("AppInstance", id), NotFoundMsg, response)
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.Module", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Module", response)
	}

	return v.(map[string]interface{}), nil
}

func (s *WebsiteBuildServiceV2) WebsiteBuildAppInstanceStateRefreshFunc(id string, field string, failStates []string) resource.StateRefreshFunc {
	return s.WebsiteBuildAppInstanceStateRefreshFuncWithApi(id, field, failStates, s.DescribeWebsiteBuildAppInstance)
}

func (s *WebsiteBuildServiceV2) WebsiteBuildAppInstanceStateRefreshFuncWithApi(id string, field string, failStates []string, call func(id string) (map[string]interface{}, error)) resource.StateRefreshFunc {
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

// DescribeWebsiteBuildAppInstance >>> Encapsulated.
