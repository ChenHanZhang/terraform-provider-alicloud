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

type OpenSearchServiceV2 struct {
	client *connectivity.AliyunClient
}

// DescribeOpenSearchAppGroupCredential <<< Encapsulated get interface for OpenSearch AppGroupCredential.

func (s *OpenSearchServiceV2) DescribeOpenSearchAppGroupCredential(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		err = WrapError(fmt.Errorf("invalid Resource Id %s. Expected parts' length %d, got %d", id, 2, len(parts)))
		return nil, err
	}
	token := parts[1]
	appGroupIdentity := parts[0]
	request = make(map[string]interface{})
	query = make(map[string]*string)

	action := fmt.Sprintf("/v4/openapi/app-groups/%s/credentials/%s", appGroupIdentity, token)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = client.RoaGet("OpenSearch", "2017-12-25", action, query, nil, nil)

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
		if IsExpectedErrors(err, []string{"Credentials.NotFound"}) {
			return object, WrapErrorf(NotFoundErr("AppGroupCredential", id), NotFoundMsg, response)
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.result", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.result", response)
	}

	return v.(map[string]interface{}), nil
}

func (s *OpenSearchServiceV2) OpenSearchAppGroupCredentialStateRefreshFunc(id string, field string, failStates []string) resource.StateRefreshFunc {
	return s.OpenSearchAppGroupCredentialStateRefreshFuncWithApi(id, field, failStates, s.DescribeOpenSearchAppGroupCredential)
}

func (s *OpenSearchServiceV2) OpenSearchAppGroupCredentialStateRefreshFuncWithApi(id string, field string, failStates []string, call func(id string) (map[string]interface{}, error)) resource.StateRefreshFunc {
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

// DescribeOpenSearchAppGroupCredential >>> Encapsulated.
