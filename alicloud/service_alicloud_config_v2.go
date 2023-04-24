package alicloud

import (
	"fmt"
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

type ConfigServiceV2 struct {
	client *connectivity.AliyunClient
}

// DescribeConfigRule <<< Encapsulated get interface for Config Rule.
func (s *ConfigServiceV2) DescribeConfigRule(id string) (object map[string]interface{}, err error) {

	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	action := "GetConfigRule"
	conn, err := client.NewConfigClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})

	query["ConfigRuleId"] = id

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-09-07"), StringPointer("AK"), query, request, &util.RuntimeOptions{})

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
		if IsExpectedErrors(err, []string{"ConfigRuleNotExists", "Invalid.ConfigRuleId.Value"}) {
			return object, WrapErrorf(Error(GetNotFoundMessage("Rule", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.ConfigRule", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.ConfigRule", response)
	}

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["account_id"] = objectRaw["AccountId"]
	instanceMap["config_rule_arn"] = objectRaw["ConfigRuleArn"]
	instanceMap["create_time"] = objectRaw["CreateTimestamp"]
	instanceMap["description"] = objectRaw["Description"]
	instanceMap["exclude_resource_ids_scope"] = objectRaw["ExcludeResourceIdsScope"]
	instanceMap["input_parameters"] = objectRaw["InputParameters"]
	instanceMap["maximum_execution_frequency"] = objectRaw["MaximumExecutionFrequency"]
	instanceMap["modified_timestamp"] = objectRaw["ModifiedTimestamp"]
	instanceMap["region_ids_scope"] = objectRaw["RegionIdsScope"]
	instanceMap["resource_group_ids_scope"] = objectRaw["ResourceGroupIdsScope"]
	instanceMap["risk_level"] = objectRaw["RiskLevel"]
	instanceMap["rule_name"] = objectRaw["ConfigRuleName"]
	instanceMap["status"] = objectRaw["ConfigRuleState"]
	instanceMap["tag_key_scope"] = objectRaw["TagKeyScope"]
	instanceMap["tag_value_scope"] = objectRaw["TagValueScope"]
	instanceMap["config_rule_id"] = objectRaw["ConfigRuleId"]
	createBy1RawObj, _ := jsonpath.Get("$.CreateBy", objectRaw)
	createBy1Raw := make(map[string]interface{})
	if createBy1RawObj != nil {
		createBy1Raw = createBy1RawObj.(map[string]interface{})
	}
	instanceMap["compliance_pack_id"] = createBy1Raw["CompliancePackId"]
	source1RawObj, _ := jsonpath.Get("$.Source", objectRaw)
	source1Raw := make(map[string]interface{})
	if source1RawObj != nil {
		source1Raw = source1RawObj.(map[string]interface{})
	}
	instanceMap["source_identifier"] = source1Raw["Identifier"]
	instanceMap["source_owner"] = source1Raw["Owner"]
	sourceDetails1RawObj, _ := jsonpath.Get("$.Source.SourceDetails[*]", objectRaw)
	sourceDetails1Raw := sourceDetails1RawObj.([]interface{})

	sourceDetailsChild1Raw := make(map[string]interface{})
	if len(sourceDetails1Raw) > 0 {
		sourceDetailsChild1Raw = sourceDetails1Raw[0].(map[string]interface{})
	}
	instanceMap["config_rule_trigger_types"] = sourceDetailsChild1Raw["MessageType"]
	instanceMap["event_source"] = sourceDetailsChild1Raw["EventSource"]
	{
		complianceMaps := make([]map[string]interface{}, 0)
		complianceMap := make(map[string]interface{})
		compliance1Raw := make(map[string]interface{})
		if objectRaw["Compliance"] != nil {
			compliance1Raw = objectRaw["Compliance"].(map[string]interface{})
		}

		complianceMap["compliance_type"] = compliance1Raw["ComplianceType"]
		complianceMap["count"] = compliance1Raw["Count"]
		complianceMaps = append(complianceMaps, complianceMap)
		instanceMap["compliance"] = complianceMaps
	}

	complianceResourceTypes1Raw, _ := jsonpath.Get("$.Scope.ComplianceResourceTypes", objectRaw)
	instanceMap["resource_types_scope"] = complianceResourceTypes1Raw

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *ConfigServiceV2) ConfigRuleStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeConfigRule(id)
		if err != nil {
			if NotFoundError(err) {
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		currentStatus := object["status"]
		if _, ok := currentStatus.(string); !ok {
			return nil, "", nil
		}
		for _, failState := range failStates {
			if currentStatus.(string) == failState {
				return object, currentStatus.(string), WrapError(Error(FailedToReachTargetStatus, currentStatus.(string)))
			}
		}
		return object, currentStatus.(string), nil
	}
}

// DescribeConfigRule >>> Encapsulated.
// DescribeConfigRemediation <<< Encapsulated get interface for Config Remediation.
func (s *ConfigServiceV2) DescribeConfigRemediation(id string) (object map[string]interface{}, err error) {

	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	action := "DescribeRemediation"
	conn, err := client.NewConfigClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})

	query["RemediationId"] = id

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("GET"), StringPointer("2020-09-07"), StringPointer("AK"), query, request, &util.RuntimeOptions{})

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
		if IsExpectedErrors(err, []string{"RemediationConfigNotExist", "RemediationNotExist"}) {
			return object, WrapErrorf(Error(GetNotFoundMessage("Remediation", id)), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.Remediation", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Remediation", response)
	}

	instanceMaps := make([]map[string]interface{}, 0)
	instanceMap := make(map[string]interface{})
	objectRaw := v.(map[string]interface{})
	instanceMap["config_rule_id"] = objectRaw["ConfigRuleId"]
	instanceMap["invoke_type"] = objectRaw["InvokeType"]
	instanceMap["params"] = objectRaw["RemediationOriginParams"]
	instanceMap["remediation_source_type"] = objectRaw["RemediationSourceType"]
	instanceMap["remediation_template_id"] = objectRaw["RemediationTemplateId"]
	instanceMap["remediation_type"] = objectRaw["RemediationType"]
	instanceMap["remediation_id"] = objectRaw["RemediationId"]

	instanceMaps = append(instanceMaps, instanceMap)
	return instanceMaps[0], nil
}

func (s *ConfigServiceV2) ConfigRemediationStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeConfigRemediation(id)
		if err != nil {
			if NotFoundError(err) {
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		currentStatus := object["status"]
		if _, ok := currentStatus.(string); !ok {
			return nil, "", nil
		}
		for _, failState := range failStates {
			if currentStatus.(string) == failState {
				return object, currentStatus.(string), WrapError(Error(FailedToReachTargetStatus, currentStatus.(string)))
			}
		}
		return object, currentStatus.(string), nil
	}
}

// DescribeConfigRemediation >>> Encapsulated.
