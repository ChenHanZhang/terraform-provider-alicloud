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

type AiContentServiceV2 struct {
	client *connectivity.AliyunClient
}

// DescribeAiContentPersonalizedTextToImageImageAsset <<< Encapsulated get interface for AiContent PersonalizedTextToImageImageAsset.

func (s *AiContentServiceV2) DescribeAiContentPersonalizedTextToImageImageAsset(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var header map[string]*string
	request = make(map[string]interface{})
	query = make(map[string]*string)
	header = make(map[string]*string)
	query["inferenceJobId"] = StringPointer(id)

	action := fmt.Sprintf("/api/v1/personalizedtxt2img/queryPreModelInferenceJobInfo")

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = client.RoaGet("AiContent", "20240611", action, query, header, nil)

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

	v, err := jsonpath.Get("$.data", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.data", response)
	}

	return v.(map[string]interface{}), nil
}

func (s *AiContentServiceV2) AiContentPersonalizedTextToImageImageAssetStateRefreshFunc(id string, field string, failStates []string) resource.StateRefreshFunc {
	return s.AiContentPersonalizedTextToImageImageAssetStateRefreshFuncWithApi(id, field, failStates, s.DescribeAiContentPersonalizedTextToImageImageAsset)
}

func (s *AiContentServiceV2) AiContentPersonalizedTextToImageImageAssetStateRefreshFuncWithApi(id string, field string, failStates []string, call func(id string) (map[string]interface{}, error)) resource.StateRefreshFunc {
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

// DescribeAiContentPersonalizedTextToImageImageAsset >>> Encapsulated.
