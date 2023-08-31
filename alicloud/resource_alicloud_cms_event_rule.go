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

func resourceAliCloudCloudMonitorServiceEventRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCloudMonitorServiceEventRuleCreate,
		Read:   resourceAliCloudCloudMonitorServiceEventRuleRead,
		Update: resourceAliCloudCloudMonitorServiceEventRuleUpdate,
		Delete: resourceAliCloudCloudMonitorServiceEventRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"contact_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"contact_group_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"contact_parameters_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"level": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_pattern": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"level_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"sql_filter": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"status_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"event_type_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"product": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"event_rule_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fc_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"function_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"service_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"fc_parameters_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"arn": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mns_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mns_parameters_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"arn": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"topic": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"queue": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"open_api_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"action": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"product": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"open_api_parameters_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"arn": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"silence_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sls_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_store": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"sls_parameters_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"arn": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"ENABLED", "DISABLED"}, false),
			},
			"webhook_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"webhook_parameters_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"method": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"url": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceAliCloudCloudMonitorServiceEventRuleCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "PutEventRule"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewCloudmonitorserviceClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RuleName"] = d.Get("event_rule_name")

	eventPatternMaps := make([]map[string]interface{}, 0)
	for _, eventPattern := range d.Get("event_pattern").(*schema.Set).List() {
		eventPatternMap := make(map[string]interface{})
		eventPatternArg := eventPattern.(map[string]interface{})

		if v, ok := eventPatternArg["product"].(string); ok {
			eventPatternMap["Product"] = v
		}
		if v, ok := eventPatternArg["event_type_list"]; ok {
			eventPatternMap["EventTypeList"] = v
		}
		if v, ok := eventPatternArg["level_list"]; ok {
			eventPatternMap["LevelList"] = v
		}
		if v, ok := eventPatternArg["name_list"]; ok {
			eventPatternMap["NameList"] = v
		}
		if v, ok := eventPatternArg["sql_filter"].(string); ok && v != "" {
			eventPatternMap["SQLFilter"] = v
		}

		eventPatternMaps = append(eventPatternMaps, eventPatternMap)
	}
	request["EventPattern"] = eventPatternMaps

	if v, ok := d.GetOk("group_id"); ok {
		request["GroupId"] = convertCloudMonitorServiceGroupIdRequest(v.(string))
	}
	request["EventType"] = "SYSTEM"
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("status"); ok {
		request["State"] = v
	}
	if v, ok := d.GetOk("silence_time"); ok {
		request["SilenceTime"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cms_event_rule", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["RuleName"]))

	return resourceAliCloudCloudMonitorServiceEventRuleUpdate(d, meta)
}

func resourceAliCloudCloudMonitorServiceEventRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cloudMonitorServiceServiceV2 := CloudMonitorServiceServiceV2{client}

	objectRaw, err := cloudMonitorServiceServiceV2.DescribeCloudMonitorServiceEventRule(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cms_event_rule DescribeCloudMonitorServiceEventRule Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("description", objectRaw["Description"])
	d.Set("group_id", objectRaw["GroupId"])
	d.Set("status", objectRaw["State"])
	d.Set("event_rule_name", objectRaw["Name"])
	eventPatternMaps := make([]map[string]interface{}, 0)
	eventPatternMap := make(map[string]interface{})
	eventPattern1Raw := make(map[string]interface{})
	if objectRaw["EventPattern"] != nil {
		eventPattern1Raw = objectRaw["EventPattern"].(map[string]interface{})
	}
	if len(eventPattern1Raw) > 0 {
		eventPatternMap["product"] = eventPattern1Raw["Product"]
		eventPatternMap["sql_filter"] = eventPattern1Raw["SQLFilter"]
		eventTypeList3Raw, _ := jsonpath.Get("$.EventPattern.EventTypeList.EventTypeList", objectRaw)
		eventPatternMap["event_type_list"] = eventTypeList3Raw
		levelList3Raw, _ := jsonpath.Get("$.EventPattern.LevelList.LevelList", objectRaw)
		eventPatternMap["level_list"] = levelList3Raw
		nameList3Raw, _ := jsonpath.Get("$.EventPattern.NameList.NameList", objectRaw)
		eventPatternMap["name_list"] = nameList3Raw
		statusList3Raw, _ := jsonpath.Get("$.EventPattern.StatusList.StatusList", objectRaw)
		eventPatternMap["status_list"] = statusList3Raw
		eventPatternMaps = append(eventPatternMaps, eventPatternMap)
	}
	d.Set("event_pattern", eventPatternMaps)

	objectRaw, err = cloudMonitorServiceServiceV2.DescribeDescribeEventRuleTargetList(d.Id())
	if err != nil {
		return WrapError(err)
	}

	contactParameter1Raw, _ := jsonpath.Get("$.ContactParameters.ContactParameter", objectRaw)
	contactParametersMaps := make([]map[string]interface{}, 0)
	if contactParameter1Raw != nil {
		for _, contactParameterChild1Raw := range contactParameter1Raw.([]interface{}) {
			contactParametersMap := make(map[string]interface{})
			contactParameterChild1Raw := contactParameterChild1Raw.(map[string]interface{})
			contactParametersMap["contact_group_name"] = contactParameterChild1Raw["ContactGroupName"]
			contactParametersMap["contact_parameters_id"] = contactParameterChild1Raw["Id"]
			contactParametersMap["level"] = contactParameterChild1Raw["Level"]
			contactParametersMaps = append(contactParametersMaps, contactParametersMap)
		}
	}
	d.Set("contact_parameters", contactParametersMaps)
	fCParameter1Raw, _ := jsonpath.Get("$.FcParameters.FCParameter", objectRaw)
	fcParametersMaps := make([]map[string]interface{}, 0)
	if fCParameter1Raw != nil {
		for _, fCParameterChild1Raw := range fCParameter1Raw.([]interface{}) {
			fcParametersMap := make(map[string]interface{})
			fCParameterChild1Raw := fCParameterChild1Raw.(map[string]interface{})
			fcParametersMap["arn"] = fCParameterChild1Raw["Arn"]
			fcParametersMap["fc_parameters_id"] = fCParameterChild1Raw["Id"]
			fcParametersMap["function_name"] = fCParameterChild1Raw["FunctionName"]
			fcParametersMap["region"] = fCParameterChild1Raw["Region"]
			fcParametersMap["service_name"] = fCParameterChild1Raw["ServiceName"]
			fcParametersMaps = append(fcParametersMaps, fcParametersMap)
		}
	}
	d.Set("fc_parameters", fcParametersMaps)
	mnsParameter1Raw, _ := jsonpath.Get("$.MnsParameters.MnsParameter", objectRaw)
	mnsParametersMaps := make([]map[string]interface{}, 0)
	if mnsParameter1Raw != nil {
		for _, mnsParameterChild1Raw := range mnsParameter1Raw.([]interface{}) {
			mnsParametersMap := make(map[string]interface{})
			mnsParameterChild1Raw := mnsParameterChild1Raw.(map[string]interface{})
			mnsParametersMap["arn"] = mnsParameterChild1Raw["Arn"]
			mnsParametersMap["mns_parameters_id"] = mnsParameterChild1Raw["Id"]
			mnsParametersMap["queue"] = mnsParameterChild1Raw["Queue"]
			mnsParametersMap["region"] = mnsParameterChild1Raw["Region"]
			mnsParametersMap["topic"] = mnsParameterChild1Raw["Topic"]
			mnsParametersMaps = append(mnsParametersMaps, mnsParametersMap)
		}
	}
	d.Set("mns_parameters", mnsParametersMaps)
	openApiParameters3Raw, _ := jsonpath.Get("$.OpenApiParameters.OpenApiParameters", objectRaw)
	openApiParametersMaps := make([]map[string]interface{}, 0)
	if openApiParameters3Raw != nil {
		for _, openApiParametersChild1Raw := range openApiParameters3Raw.([]interface{}) {
			openApiParametersMap := make(map[string]interface{})
			openApiParametersChild1Raw := openApiParametersChild1Raw.(map[string]interface{})
			openApiParametersMap["action"] = openApiParametersChild1Raw["Action"]
			openApiParametersMap["arn"] = openApiParametersChild1Raw["Arn"]
			openApiParametersMap["open_api_parameters_id"] = openApiParametersChild1Raw["Id"]
			openApiParametersMap["product"] = openApiParametersChild1Raw["Product"]
			openApiParametersMap["region"] = openApiParametersChild1Raw["Region"]
			openApiParametersMap["role"] = openApiParametersChild1Raw["Role"]
			openApiParametersMap["version"] = openApiParametersChild1Raw["Version"]
			openApiParametersMaps = append(openApiParametersMaps, openApiParametersMap)
		}
	}
	d.Set("open_api_parameters", openApiParametersMaps)
	slsParameter1Raw, _ := jsonpath.Get("$.SlsParameters.SlsParameter", objectRaw)
	slsParametersMaps := make([]map[string]interface{}, 0)
	if slsParameter1Raw != nil {
		for _, slsParameterChild1Raw := range slsParameter1Raw.([]interface{}) {
			slsParametersMap := make(map[string]interface{})
			slsParameterChild1Raw := slsParameterChild1Raw.(map[string]interface{})
			slsParametersMap["arn"] = slsParameterChild1Raw["Arn"]
			slsParametersMap["log_store"] = slsParameterChild1Raw["LogStore"]
			slsParametersMap["project"] = slsParameterChild1Raw["Project"]
			slsParametersMap["region"] = slsParameterChild1Raw["Region"]
			slsParametersMap["sls_parameters_id"] = slsParameterChild1Raw["Id"]
			slsParametersMaps = append(slsParametersMaps, slsParametersMap)
		}
	}
	d.Set("sls_parameters", slsParametersMaps)
	webhookParameter1Raw, _ := jsonpath.Get("$.WebhookParameters.WebhookParameter", objectRaw)
	webhookParametersMaps := make([]map[string]interface{}, 0)
	if webhookParameter1Raw != nil {
		for _, webhookParameterChild1Raw := range webhookParameter1Raw.([]interface{}) {
			webhookParametersMap := make(map[string]interface{})
			webhookParameterChild1Raw := webhookParameterChild1Raw.(map[string]interface{})
			webhookParametersMap["method"] = webhookParameterChild1Raw["Method"]
			webhookParametersMap["protocol"] = webhookParameterChild1Raw["Protocol"]
			webhookParametersMap["url"] = webhookParameterChild1Raw["Url"]
			webhookParametersMap["webhook_parameters_id"] = webhookParameterChild1Raw["Id"]
			webhookParametersMaps = append(webhookParametersMaps, webhookParametersMap)
		}
	}
	d.Set("webhook_parameters", webhookParametersMaps)

	return nil
}

func resourceAliCloudCloudMonitorServiceEventRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	d.Partial(true)
	action := "PutEventRule"
	conn, err := client.NewCloudmonitorserviceClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RuleName"] = d.Id()
	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	request["EventType"] = "SYSTEM"
	if !d.IsNewResource() && d.HasChange("group_id") {
		update = true
		request["GroupId"] = d.Get("group_id")
	}

	if !d.IsNewResource() && d.HasChange("status") {
		update = true
		request["State"] = d.Get("status")
	}

	if v, ok := d.GetOk("silence_time"); ok {
		request["SilenceTime"] = v
	}

	if !d.IsNewResource() && d.HasChange("event_pattern") {
		update = true
	}
	eventPatternMaps := make([]map[string]interface{}, 0)
	for _, eventPattern := range d.Get("event_pattern").(*schema.Set).List() {
		eventPatternMap := make(map[string]interface{}, 0)
		eventPatternArg := eventPattern.(map[string]interface{})

		if v, ok := eventPatternArg["product"].(string); ok {
			eventPatternMap["Product"] = v
		}
		if v, ok := eventPatternArg["event_type_list"]; ok {
			eventPatternMap["EventTypeList"] = v
		}
		if v, ok := eventPatternArg["level_list"]; ok {
			eventPatternMap["LevelList"] = v
		}
		if v, ok := eventPatternArg["name_list"]; ok {
			eventPatternMap["NameList"] = v
		}
		if v, ok := eventPatternArg["sql_filter"].(string); ok && v != "" {
			eventPatternMap["SQLFilter"] = v
		}

		eventPatternMaps = append(eventPatternMaps, eventPatternMap)
	}
	request["EventPattern"] = eventPatternMaps

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

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
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("description")
		d.SetPartial("group_id")
		d.SetPartial("status")
	}
	update = false
	action = "PutEventRuleTargets"
	conn, err = client.NewCloudmonitorserviceClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RuleName"] = d.Id()
	if d.HasChange("fc_parameters") {
		update = true
		if v, ok := d.GetOk("fc_parameters"); ok {
			fcParametersMaps := make([]map[string]interface{}, 0)
			for _, dataLoop := range v.([]interface{}) {
				dataLoopTmp := dataLoop.(map[string]interface{})
				dataLoopMap := make(map[string]interface{})
				dataLoopMap["FunctionName"] = dataLoopTmp["function_name"]
				dataLoopMap["Region"] = dataLoopTmp["region"]
				dataLoopMap["ServiceName"] = dataLoopTmp["service_name"]
				dataLoopMap["Id"] = dataLoopTmp["fc_parameters_id"]
				fcParametersMaps = append(fcParametersMaps, dataLoopMap)
			}
			request["FcParameters"] = fcParametersMaps
		}
	}

	if d.HasChange("contact_parameters") {
		update = true
		if v, ok := d.GetOk("contact_parameters"); ok {
			contactParametersMaps := make([]map[string]interface{}, 0)
			for _, dataLoop1 := range v.([]interface{}) {
				dataLoop1Tmp := dataLoop1.(map[string]interface{})
				dataLoop1Map := make(map[string]interface{})
				dataLoop1Map["ContactGroupName"] = dataLoop1Tmp["contact_group_name"]
				dataLoop1Map["Level"] = dataLoop1Tmp["level"]
				dataLoop1Map["Id"] = dataLoop1Tmp["contact_parameters_id"]
				contactParametersMaps = append(contactParametersMaps, dataLoop1Map)
			}
			request["ContactParameters"] = contactParametersMaps
		}
	}

	if d.HasChange("mns_parameters") {
		update = true
		if v, ok := d.GetOk("mns_parameters"); ok {
			mnsParametersMaps := make([]map[string]interface{}, 0)
			for _, dataLoop2 := range v.([]interface{}) {
				dataLoop2Tmp := dataLoop2.(map[string]interface{})
				dataLoop2Map := make(map[string]interface{})
				dataLoop2Map["Region"] = dataLoop2Tmp["region"]
				dataLoop2Map["Queue"] = dataLoop2Tmp["queue"]
				dataLoop2Map["Topic"] = dataLoop2Tmp["topic"]
				dataLoop2Map["Id"] = dataLoop2Tmp["mns_parameters_id"]
				mnsParametersMaps = append(mnsParametersMaps, dataLoop2Map)
			}
			request["MnsParameters"] = mnsParametersMaps
		}
	}

	if d.HasChange("webhook_parameters") {
		update = true
		if v, ok := d.GetOk("webhook_parameters"); ok {
			webhookParametersMaps := make([]map[string]interface{}, 0)
			for _, dataLoop3 := range v.([]interface{}) {
				dataLoop3Tmp := dataLoop3.(map[string]interface{})
				dataLoop3Map := make(map[string]interface{})
				dataLoop3Map["Protocol"] = dataLoop3Tmp["protocol"]
				dataLoop3Map["Url"] = dataLoop3Tmp["url"]
				dataLoop3Map["Method"] = dataLoop3Tmp["method"]
				dataLoop3Map["Id"] = dataLoop3Tmp["webhook_parameters_id"]
				webhookParametersMaps = append(webhookParametersMaps, dataLoop3Map)
			}
			request["WebhookParameters"] = webhookParametersMaps
		}
	}

	if d.HasChange("sls_parameters") {
		update = true
		if v, ok := d.GetOk("sls_parameters"); ok {
			slsParametersMaps := make([]map[string]interface{}, 0)
			for _, dataLoop4 := range v.([]interface{}) {
				dataLoop4Tmp := dataLoop4.(map[string]interface{})
				dataLoop4Map := make(map[string]interface{})
				dataLoop4Map["LogStore"] = dataLoop4Tmp["log_store"]
				dataLoop4Map["Region"] = dataLoop4Tmp["region"]
				dataLoop4Map["Project"] = dataLoop4Tmp["project"]
				dataLoop4Map["Id"] = dataLoop4Tmp["sls_parameters_id"]
				slsParametersMaps = append(slsParametersMaps, dataLoop4Map)
			}
			request["SlsParameters"] = slsParametersMaps
		}
	}

	if d.HasChange("open_api_parameters") {
		update = true
		if v, ok := d.GetOk("open_api_parameters"); ok {
			openApiParametersMaps := make([]map[string]interface{}, 0)
			for _, dataLoop5 := range v.([]interface{}) {
				dataLoop5Tmp := dataLoop5.(map[string]interface{})
				dataLoop5Map := make(map[string]interface{})
				dataLoop5Map["Action"] = dataLoop5Tmp["action"]
				dataLoop5Map["Region"] = dataLoop5Tmp["region"]
				dataLoop5Map["Role"] = dataLoop5Tmp["role"]
				dataLoop5Map["Version"] = dataLoop5Tmp["version"]
				dataLoop5Map["Product"] = dataLoop5Tmp["product"]
				dataLoop5Map["Id"] = dataLoop5Tmp["open_api_parameters_id"]
				openApiParametersMaps = append(openApiParametersMaps, dataLoop5Map)
			}
			request["OpenApiParameters"] = openApiParametersMaps
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

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
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}

	d.Partial(false)
	return resourceAliCloudCloudMonitorServiceEventRuleRead(d, meta)
}

func resourceAliCloudCloudMonitorServiceEventRuleDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteEventRules"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewCloudmonitorserviceClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RuleNames.1"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

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
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}

func convertCloudMonitorServiceGroupIdRequest(source interface{}) interface{} {
	switch source {
	}
	return source
}
