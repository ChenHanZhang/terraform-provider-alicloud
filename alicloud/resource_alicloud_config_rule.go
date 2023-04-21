// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudConfigRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudConfigRuleCreate,
		Read:   resourceAlicloudConfigRuleRead,
		Update: resourceAlicloudConfigRuleUpdate,
		Delete: resourceAlicloudConfigRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"compliance": {
				Type:     schema.TypeSet,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compliance_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"compliance_pack_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_rule_arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_rule_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_rule_trigger_types": {
				Type:     schema.TypeString,
				Required: true,
			},
			"create_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_source": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_resource_ids_scope": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"input_parameters": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"maximum_execution_frequency": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modified_timestamp": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"region_ids_scope": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group_ids_scope": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_types_scope": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"risk_level": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"rule_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_identifier": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_owner": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag_key_scope": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tag_value_scope": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAlicloudConfigRuleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "CreateConfigRule"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewConfigClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}

	if v, ok := d.GetOk("source_owner"); ok {
		request["SourceOwner"] = v
	}

	if v, ok := d.GetOk("source_identifier"); ok {
		request["SourceIdentifier"] = v
	}

	if v, ok := d.GetOk("config_rule_trigger_types"); ok {
		request["ConfigRuleTriggerTypes"] = v
	}

	if v, ok := d.GetOk("risk_level"); ok {
		request["RiskLevel"] = v
	}

	if v, ok := d.GetOk("rule_name"); ok {
		request["ConfigRuleName"] = v
	}

	if v, ok := d.GetOk("tag_key_scope"); ok {
		request["TagKeyScope"] = v
	}

	if v, ok := d.GetOk("tag_value_scope"); ok {
		request["TagValueScope"] = v
	}

	if v, ok := d.GetOk("region_ids_scope"); ok {
		request["RegionIdsScope"] = v
	}

	if v, ok := d.GetOk("resource_group_ids_scope"); ok {
		request["ResourceGroupIdsScope"] = v
	}

	if v, ok := d.GetOk("exclude_resource_ids_scope"); ok {
		request["ExcludeResourceIdsScope"] = v
	}

	if v, ok := d.GetOk("input_parameters"); ok {
		request["InputParameters"] = convertMapToJsonStringIgnoreError(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("resource_types_scope"); ok {
		jsonPathResult12, err := jsonpath.Get("$", v)
		if err != nil {
			return WrapError(err)
		}
		request["ResourceTypesScope"] = convertListToCommaSeparate(jsonPathResult12.([]interface{}))
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-09-07"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		request["ClientToken"] = buildClientToken(action)

		if err != nil {
			if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_config_rule", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["ConfigRuleId"]))

	configServiceV2 := ConfigServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"ACTIVE", "EVALUATING"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, configServiceV2.ConfigRuleStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAlicloudConfigRuleUpdate(d, meta)
}

func resourceAlicloudConfigRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	configServiceV2 := ConfigServiceV2{client}

	object, err := configServiceV2.DescribeConfigRule(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_config_rule .DescribeConfigRule Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("account_id", object["account_id"])
	d.Set("compliance", object["compliance"])
	d.Set("compliance_pack_id", object["compliance_pack_id"])
	d.Set("config_rule_arn", object["config_rule_arn"])
	d.Set("config_rule_id", object["config_rule_id"])
	d.Set("config_rule_trigger_types", object["config_rule_trigger_types"])
	d.Set("create_time", object["create_time"])
	d.Set("description", object["description"])
	d.Set("event_source", object["event_source"])
	d.Set("exclude_resource_ids_scope", object["exclude_resource_ids_scope"])
	d.Set("input_parameters", object["input_parameters"])
	d.Set("maximum_execution_frequency", object["maximum_execution_frequency"])
	d.Set("modified_timestamp", object["modified_timestamp"])
	d.Set("region_ids_scope", object["region_ids_scope"])
	d.Set("resource_group_ids_scope", object["resource_group_ids_scope"])
	d.Set("resource_types_scope", object["resource_types_scope"])
	d.Set("risk_level", object["risk_level"])
	d.Set("rule_name", object["rule_name"])
	d.Set("source_identifier", object["source_identifier"])
	d.Set("source_owner", object["source_owner"])
	d.Set("status", object["status"])
	d.Set("tag_key_scope", object["tag_key_scope"])
	d.Set("tag_value_scope", object["tag_value_scope"])

	return nil
}

func resourceAlicloudConfigRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	d.Partial(true)
	update = false
	action := "UpdateConfigRule"
	conn, err := client.NewConfigClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["ConfigRuleId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)

	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		if v, ok := d.GetOk("description"); ok {
			request["Description"] = v
		}
	}
	if d.HasChange("maximum_execution_frequency") {
		update = true
		if v, ok := d.GetOk("maximum_execution_frequency"); ok {
			request["MaximumExecutionFrequency"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("tag_key_scope") {
		update = true
		if v, ok := d.GetOk("tag_key_scope"); ok {
			request["TagKeyScope"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("tag_value_scope") {
		update = true
		if v, ok := d.GetOk("tag_value_scope"); ok {
			request["TagValueScope"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("region_ids_scope") {
		update = true
		if v, ok := d.GetOk("region_ids_scope"); ok {
			request["RegionIdsScope"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("resource_group_ids_scope") {
		update = true
		if v, ok := d.GetOk("resource_group_ids_scope"); ok {
			request["ResourceGroupIdsScope"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("exclude_resource_ids_scope") {
		update = true
		if v, ok := d.GetOk("exclude_resource_ids_scope"); ok {
			request["ExcludeResourceIdsScope"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("risk_level") {
		update = true
		if v, ok := d.GetOk("risk_level"); ok {
			request["RiskLevel"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("config_rule_trigger_types") {
		update = true
		if v, ok := d.GetOk("config_rule_trigger_types"); ok {
			request["ConfigRuleTriggerTypes"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("input_parameters") {
		update = true
		if v, ok := d.GetOk("input_parameters"); ok {
			request["InputParameters"] = convertMapToJsonStringIgnoreError(v.(map[string]interface{}))
		}
	}
	if !d.IsNewResource() && d.HasChange("resource_types_scope") {
		update = true
		if v, ok := d.GetOk("resource_types_scope"); ok {
			jsonPathResult10, err := jsonpath.Get("$", v)
			if err != nil {
				return WrapError(err)
			}
			request["ResourceTypesScope"] = convertListToCommaSeparate(jsonPathResult10.([]interface{}))
		}
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-09-07"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			request["ClientToken"] = buildClientToken(action)

			if err != nil {
				if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		{
			configServiceV2 := ConfigServiceV2{client}
			stateConf := BuildStateConf([]string{}, []string{"ACTIVE", "EVALUATING"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, configServiceV2.ConfigRuleStateRefreshFunc(d.Id(), []string{}))
			if _, err := stateConf.WaitForState(); err != nil {
				return WrapErrorf(err, IdMsg, d.Id())
			}
		}
		d.SetPartial("description")
		d.SetPartial("maximum_execution_frequency")
		d.SetPartial("tag_key_scope")
		d.SetPartial("tag_value_scope")
		d.SetPartial("region_ids_scope")
		d.SetPartial("resource_group_ids_scope")
		d.SetPartial("exclude_resource_ids_scope")
		d.SetPartial("risk_level")
		d.SetPartial("config_rule_trigger_types")
		d.SetPartial("input_parameters")
		d.SetPartial("resource_types_scope")
	}

	if d.HasChange("status") {
		client := meta.(*connectivity.AliyunClient)
		configServiceV2 := ConfigServiceV2{client}
		object, err := configServiceV2.DescribeConfigRule(d.Id())
		if err != nil {
			return WrapError(err)
		}

		target := d.Get("status").(string)
		if object["status"].(string) != target {
			if target == "INACTIVE" {
				action = "StopConfigRules"
				conn, err = client.NewConfigClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["ConfigRuleIds"] = d.Id()

				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-01-08"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

					if err != nil {
						if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
							wait()
							return resource.RetryableError(err)
						}
						return resource.NonRetryableError(err)
					}
					addDebug(action, response, request)
					return nil
				})
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}

			}
			if target == "ACTIVE" {
				action = "ActiveConfigRules"
				conn, err = client.NewConfigClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["ConfigRuleIds"] = d.Id()

				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-01-08"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

					if err != nil {
						if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
							wait()
							return resource.RetryableError(err)
						}
						return resource.NonRetryableError(err)
					}
					addDebug(action, response, request)
					return nil
				})
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}

			}
		}
	}

	d.Partial(false)
	return resourceAlicloudConfigRuleRead(d, meta)
}

func resourceAlicloudConfigRuleDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "DeleteConfigRules"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewConfigClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["ConfigRuleIds"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-01-08"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

		if err != nil {
			if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		if IsExpectedErrors(err, []string{"ConfigRuleNotExists"}) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
