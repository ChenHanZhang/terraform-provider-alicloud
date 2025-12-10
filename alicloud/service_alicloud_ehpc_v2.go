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

type EhpcServiceV2 struct {
	client *connectivity.AliyunClient
}

// DescribeEhpcClusterV2 <<< Encapsulated get interface for Ehpc ClusterV2.

func (s *EhpcServiceV2) DescribeEhpcClusterV2(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ClusterId"] = id

	action := "ListSharedStorages"

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("EHPC", "2024-07-30", action, query, request, true)

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
		if IsExpectedErrors(err, []string{"ClusterNotFound"}) {
			return object, WrapErrorf(NotFoundErr("ClusterV2", id), NotFoundMsg, response)
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.SharedStorages[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.SharedStorages[*]", response)
	}

	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(NotFoundErr("ClusterV2", id), NotFoundMsg, response)
	}

	return v.([]interface{})[0].(map[string]interface{}), nil
}
func (s *EhpcServiceV2) DescribeClusterV2GetCluster(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ClusterId"] = id

	action := "GetCluster"

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("EHPC", "2024-07-30", action, query, request, true)

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
		if IsExpectedErrors(err, []string{"ClusterNotFound"}) {
			return object, WrapErrorf(NotFoundErr("ClusterV2", id), NotFoundMsg, response)
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	return response, nil
}

func (s *EhpcServiceV2) EhpcClusterV2StateRefreshFunc(id string, field string, failStates []string) resource.StateRefreshFunc {
	return s.EhpcClusterV2StateRefreshFuncWithApi(id, field, failStates, s.DescribeEhpcClusterV2)
}

func (s *EhpcServiceV2) EhpcClusterV2StateRefreshFuncWithApi(id string, field string, failStates []string, call func(id string) (map[string]interface{}, error)) resource.StateRefreshFunc {
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
		if field == "$.FileSystemId" {
			e := jsonata.MustCompile("$.$.MountInfo.($$.$.FileSystemId)")
			v, _ = e.Eval(object)
			currentStatus = fmt.Sprint(v)
		}
		if field == "$.MountInfo[*].StorageDirectory" {
			e := jsonata.MustCompile("$.$.MountInfo.(MountDirectory = '/home' or MountDirectory = '/opt' ? $split(StorageDirectory, '/ehpc')[0]: StorageDirectory)")
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

// DescribeEhpcClusterV2 >>> Encapsulated.
