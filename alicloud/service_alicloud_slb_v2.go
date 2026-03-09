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

type SlbServiceV2 struct {
	client *connectivity.AliyunClient
}

// DescribeSlbAccessControlListEntryAttachment <<< Encapsulated get interface for Slb AccessControlListEntryAttachment.

func (s *SlbServiceV2) DescribeSlbAccessControlListEntryAttachment(id string) (object map[string]interface{}, err error) {
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
	request["AclId"] = parts[0]
	request["RegionId"] = client.RegionId
	action := "DescribeAccessControlListAttribute"

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Slb", "2014-05-15", action, query, request, true)

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

	currentStatus, err := jsonpath.Get("$.AclEntrys.AclEntry[0].AclEntryIP", response)
	if currentStatus == nil {
		return object, WrapErrorf(NotFoundErr("AccessControlListEntryAttachment", id), NotFoundMsg, response)
	}

	return response, nil
}

func (s *SlbServiceV2) SlbAccessControlListEntryAttachmentStateRefreshFunc(id string, field string, failStates []string) resource.StateRefreshFunc {
	return s.SlbAccessControlListEntryAttachmentStateRefreshFuncWithApi(id, field, failStates, s.DescribeSlbAccessControlListEntryAttachment)
}

func (s *SlbServiceV2) SlbAccessControlListEntryAttachmentStateRefreshFuncWithApi(id string, field string, failStates []string, call func(id string) (map[string]interface{}, error)) resource.StateRefreshFunc {
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

// DescribeSlbAccessControlListEntryAttachment >>> Encapsulated.
