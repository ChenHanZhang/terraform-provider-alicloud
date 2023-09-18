package alicloud

import (
	"fmt"
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

type ThreatDetectionServiceV2 struct {
	client *connectivity.AliyunClient
}

// DescribeThreatDetectionLogMeta <<< Encapsulated get interface for ThreatDetection LogMeta.

func (s *ThreatDetectionServiceV2) DescribeThreatDetectionLogMeta(id string) (object map[string]interface{}, err error) {

	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	action := "GetLogMeta"
	conn, err := client.NewThreatdetectionClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	query["LogStore"] = id

	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2018-12-03"), StringPointer("AK"), query, request, &runtime)

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
		if IsExpectedErrors(err, []string{}) {
			return object, WrapErrorf(Error(GetNotFoundMessage("LogMeta", id)), NotFoundMsg, response)
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.LogMeta", response)
	if err != nil {
		return object, WrapErrorf(Error(GetNotFoundMessage("LogMeta", id)), NotFoundMsg, response)
	}

	currentStatus := v.(map[string]interface{})["LogStore"]
	if currentStatus == nil {
		return object, WrapErrorf(Error(GetNotFoundMessage("LogMeta", id)), NotFoundMsg, response)
	}

	return v.(map[string]interface{}), nil
}

func (s *ThreatDetectionServiceV2) ThreatDetectionLogMetaStateRefreshFunc(id string, field string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeThreatDetectionLogMeta(id)
		if err != nil {
			if NotFoundError(err) {
				return object, "", nil
			}
			return nil, "", WrapError(err)
		}

		currentStatus := fmt.Sprint(object[field])
		for _, failState := range failStates {
			if currentStatus == failState {
				return object, currentStatus, WrapError(Error(FailedToReachTargetStatus, currentStatus))
			}
		}
		return object, currentStatus, nil
	}
}

// DescribeThreatDetectionLogMeta >>> Encapsulated.
