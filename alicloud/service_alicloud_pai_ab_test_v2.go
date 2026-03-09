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

type PaiAbTestServiceV2 struct {
	client *connectivity.AliyunClient
}

// DescribePaiAbTestCrowd <<< Encapsulated get interface for PaiAbTest Crowd.

func (s *PaiAbTestServiceV2) DescribePaiAbTestCrowd(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	CrowdId := id
	request = make(map[string]interface{})
	query = make(map[string]*string)
	query["Regionid"] = StringPointer(client.RegionId)
	action := fmt.Sprintf("/api/v1/crowds/%s", CrowdId)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = client.RoaGet("PAIABTest", "2024-01-19", action, query, nil, nil)

		if err != nil {
			if IsExpectedErrors(err, []string{"503"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)
	if err != nil {
		if IsExpectedErrors(err, []string{"DoConfigCenterRequestErrorProblem"}) {
			return object, WrapErrorf(NotFoundErr("Crowd", id), NotFoundMsg, response)
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	return response, nil
}

func (s *PaiAbTestServiceV2) PaiAbTestCrowdStateRefreshFunc(id string, field string, failStates []string) resource.StateRefreshFunc {
	return s.PaiAbTestCrowdStateRefreshFuncWithApi(id, field, failStates, s.DescribePaiAbTestCrowd)
}

func (s *PaiAbTestServiceV2) PaiAbTestCrowdStateRefreshFuncWithApi(id string, field string, failStates []string, call func(id string) (map[string]interface{}, error)) resource.StateRefreshFunc {
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

// DescribePaiAbTestCrowd >>> Encapsulated.
