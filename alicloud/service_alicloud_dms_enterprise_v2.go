package alicloud

import (
	"fmt"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

type DMSEnterpriseServiceV2 struct {
	client *connectivity.AliyunClient
}

// DescribeDMSEnterpriseAuthorityTemplate <<< Encapsulated get interface for DMSEnterprise AuthorityTemplate.

func (s *DMSEnterpriseServiceV2) DescribeDMSEnterpriseAuthorityTemplate(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		err = WrapError(fmt.Errorf("invalid Resource Id %s. Expected parts' length %d, got %d", id, 2, len(parts)))
	}
	action := "GetAuthorityTemplate"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	query["TemplateId"] = parts[1]
	query["Tid"] = parts[0]

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("dms-enterprise", "2018-11-01", action, query, request, true)

		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	if _, ok := response["AuthorityTemplateView"]; !ok {
		return object, WrapErrorf(NotFoundErr("AuthorityTemplate", id), NotFoundMsg, response)
	}

	return response, nil
}

func (s *DMSEnterpriseServiceV2) DMSEnterpriseAuthorityTemplateStateRefreshFunc(id string, field string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeDMSEnterpriseAuthorityTemplate(id)
		if err != nil {
			if NotFoundError(err) {
				return object, "", nil
			}
			return nil, "", WrapError(err)
		}

		v, err := jsonpath.Get(field, object)
		currentStatus := fmt.Sprint(v)
		for _, failState := range failStates {
			if currentStatus == failState {
				return object, currentStatus, WrapError(Error(FailedToReachTargetStatus, currentStatus))
			}
		}
		return object, currentStatus, nil
	}
}

// DescribeDMSEnterpriseAuthorityTemplate >>> Encapsulated.

// DescribeDmsEnterpriseWorkspace <<< Encapsulated get interface for DmsEnterprise Workspace.

func (s *DMSEnterpriseServiceV2) DescribeDmsEnterpriseWorkspace(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["WorkspaceId"] = id
	request["RegionId"] = client.RegionId
	action := "GetWorkspace"

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("dms-enterprise", "2018-11-01", action, query, request, true)

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
		if IsExpectedErrors(err, []string{"WorkspaceNotExist"}) {
			return object, WrapErrorf(NotFoundErr("Workspace", id), NotFoundMsg, response)
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.Data", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Data", response)
	}

	return v.(map[string]interface{}), nil
}

func (s *DMSEnterpriseServiceV2) DmsEnterpriseWorkspaceStateRefreshFunc(id string, field string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeDmsEnterpriseWorkspace(id)
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

// DescribeDmsEnterpriseWorkspace >>> Encapsulated.
// DescribeDmsEnterpriseDifyInstance <<< Encapsulated get interface for DmsEnterprise DifyInstance.

func (s *DmsEnterpriseServiceV2) DescribeDmsEnterpriseDifyInstance(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["AppUuid"] = id
	request["DataRegion"] = client.RegionId
	action := "DescribeDifyAttribute"
	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("dms-enterprise", "2018-11-01", action, query, request, true)
		request["ClientToken"] = buildClientToken(action)

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
		if IsExpectedErrors(err, []string{"DMS.MLOPS.RESOURCE_DOES_NOT_EXIST"}) {
			return object, WrapErrorf(NotFoundErr("DifyInstance", id), NotFoundMsg, response)
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	return response, nil
}

func (s *DmsEnterpriseServiceV2) DmsEnterpriseDifyInstanceStateRefreshFunc(id string, field string, failStates []string) resource.StateRefreshFunc {
	return s.DmsEnterpriseDifyInstanceStateRefreshFuncWithApi(id, field, failStates, s.DescribeDmsEnterpriseDifyInstance)
}

func (s *DmsEnterpriseServiceV2) DmsEnterpriseDifyInstanceStateRefreshFuncWithApi(id string, field string, failStates []string, call func(id string) (map[string]interface{}, error)) resource.StateRefreshFunc {
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

// DescribeDmsEnterpriseDifyInstance >>> Encapsulated.
