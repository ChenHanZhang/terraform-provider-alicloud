// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/blues/jsonata-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tidwall/sjson"
)

func resourceAliCloudDdosCooDomainPreciseAccessRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudDdosCooDomainPreciseAccessRuleCreate,
		Read:   resourceAliCloudDdosCooDomainPreciseAccessRuleRead,
		Update: resourceAliCloudDdosCooDomainPreciseAccessRuleUpdate,
		Delete: resourceAliCloudDdosCooDomainPreciseAccessRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"accept", "block", "challenge", "watch"}, false),
			},
			"condition": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"match_method": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"field": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"header_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"content": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"expires": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudDdosCooDomainPreciseAccessRuleCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "ModifyWebPreciseAccessRule"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("domain"); ok {
		request["Domain"] = v
	}
	request["RegionId"] = client.RegionId

	rulesDataList := make(map[string]interface{})

	if v := d.Get("condition"); !IsNil(v) {
		localData, err := jsonpath.Get("$", v)
		if err != nil {
			localData = make([]interface{}, 0)
		}
		localMaps := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(localData) {
			dataLoopTmp := make(map[string]interface{})
			if dataLoop != nil {
				dataLoopTmp = dataLoop.(map[string]interface{})
			}
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["header_name"] = dataLoopTmp["header_name"]
			dataLoopMap["content"] = dataLoopTmp["content"]
			dataLoopMap["match_method"] = dataLoopTmp["match_method"]
			dataLoopMap["field"] = dataLoopTmp["field"]
			localMaps = append(localMaps, dataLoopMap)
		}
		rulesDataList["condition"] = localMaps

	}

	if v, ok := d.GetOk("action"); ok {
		rulesDataList["action"] = v
	}

	if v, ok := d.GetOk("name"); ok {
		rulesDataList["name"] = v
	}

	RulesMap := make([]interface{}, 0)
	RulesMap = append(RulesMap, rulesDataList)
	request["Rules"] = RulesMap

	if v, ok := d.GetOkExists("expires"); ok {
		request["Expires"] = v
	}
	jsonString := convertObjectToJsonString(request)
	jsonString, _ = sjson.Set(jsonString, "Rules.0.name", d.Get("name"))
	_ = json.Unmarshal([]byte(jsonString), &request)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("ddoscoo", "2020-01-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ddos_coo_domain_precise_access_rule", action, AlibabaCloudSdkGoERROR)
	}

	RulesnameVar, _ := jsonpath.Get("Rules[0].name", request)
	d.SetId(fmt.Sprintf("%v:%v", request["Domain"], RulesnameVar))

	return resourceAliCloudDdosCooDomainPreciseAccessRuleRead(d, meta)
}

func resourceAliCloudDdosCooDomainPreciseAccessRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	ddosCooServiceV2 := DdosCooServiceV2{client}

	objectRaw, err := ddosCooServiceV2.DescribeDdosCooDomainPreciseAccessRule(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ddos_coo_domain_precise_access_rule DescribeDdosCooDomainPreciseAccessRule Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("domain", objectRaw["Domain"])

	ruleListRawObj, _ := jsonpath.Get("$.RuleList[*]", objectRaw)
	ruleListRaw := make([]interface{}, 0)
	if ruleListRawObj != nil {
		ruleListRaw = convertToInterfaceArray(ruleListRawObj)
	}

	d.Set("action", ruleListRaw["Action"])
	d.Set("expires", ruleListRaw["ExpirePeriod"])
	d.Set("name", ruleListRaw["Name"])

	conditionListChildRawObj, _ := jsonpath.Get("$.RuleList[*].ConditionList[*]", objectRaw)
	conditionListChildRaw := make([]interface{}, 0)
	if conditionListChildRawObj != nil {
		conditionListChildRaw = convertToInterfaceArray(conditionListChildRawObj)
	}

	conditionMaps := make([]map[string]interface{}, 0)
	if conditionListChildRaw != nil {
		for _, contentRaw := range convertToInterfaceArray(conditionListChildRaw) {
			conditionMap := make(map[string]interface{})
			conditionListChildRawObj, _ := jsonpath.Get("$.RuleList[*].ConditionList[*]", objectRaw)
			conditionListChildRaw := make([]interface{}, 0)
			if conditionListChildRawObj != nil {
				conditionListChildRaw = convertToInterfaceArray(conditionListChildRawObj)
			}

			conditionMap["content"] = conditionListChildRaw["Content"]
			conditionMap["field"] = conditionListChildRaw["Field"]
			conditionMap["header_name"] = conditionListChildRaw["HeaderName"]
			conditionMap["match_method"] = conditionListChildRaw["MatchMethod"]

			conditionMaps = append(conditionMaps, conditionMap)
		}
	}
	if err := d.Set("condition", conditionMaps); err != nil {
		return err
	}

	e := jsonata.MustCompile("$.RuleList.ConditionList")
	evaluation, _ := e.Eval(objectRaw)
	d.Set("condition", evaluation)
	e = jsonata.MustCompile("$.RuleList.ConditionList.($exists(Content) ? Content : null)")
	evaluation, _ = e.Eval(objectRaw)
	d.Set("condition", evaluation)
	e = jsonata.MustCompile("$.RuleList.ConditionList.($exists(Field) ? Field : null)")
	evaluation, _ = e.Eval(objectRaw)
	d.Set("condition", evaluation)
	e = jsonata.MustCompile("$.RuleList.ConditionList.($exists(HeaderName) ? HeaderName : null)")
	evaluation, _ = e.Eval(objectRaw)
	d.Set("condition", evaluation)
	e = jsonata.MustCompile("$.RuleList.ConditionList.($exists(MatchMethod) ? MatchMethod : null)")
	evaluation, _ = e.Eval(objectRaw)
	d.Set("condition", evaluation)

	return nil
}

func resourceAliCloudDdosCooDomainPreciseAccessRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "ModifyWebPreciseAccessRule"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["Domain"] = parts[0]
	request["RegionId"] = client.RegionId
	rulesDataList := make(map[string]interface{})

	if d.HasChange("condition") {
		update = true
	}
	if v := d.Get("condition"); v != nil {
		localData, err := jsonpath.Get("$", v)
		if err != nil {
			localData = make([]interface{}, 0)
		}
		localMaps := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(localData) {
			dataLoopTmp := make(map[string]interface{})
			if dataLoop != nil {
				dataLoopTmp = dataLoop.(map[string]interface{})
			}
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["header_name"] = dataLoopTmp["header_name"]
			dataLoopMap["content"] = dataLoopTmp["content"]
			dataLoopMap["match_method"] = dataLoopTmp["match_method"]
			dataLoopMap["field"] = dataLoopTmp["field"]
			localMaps = append(localMaps, dataLoopMap)
		}
		rulesDataList["condition"] = localMaps

	}

	if d.HasChange("action") {
		update = true
	}
	if v, ok := d.GetOk("action"); ok {
		rulesDataList["action"] = v
	}

	if d.HasChange("name") {
		update = true
	}
	if v, ok := d.GetOk("name"); ok {
		rulesDataList["name"] = v
	}

	RulesMap := make([]interface{}, 0)
	RulesMap = append(RulesMap, rulesDataList)
	request["Rules"] = RulesMap

	if d.HasChange("expires") {
		update = true
		request["Expires"] = d.Get("expires")
	}

	jsonString := convertObjectToJsonString(request)
	jsonString, _ = sjson.Set(jsonString, "Rules.0.name", parts[1])
	_ = json.Unmarshal([]byte(jsonString), &request)

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("ddoscoo", "2020-01-01", action, query, request, true)
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
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}

	return resourceAliCloudDdosCooDomainPreciseAccessRuleRead(d, meta)
}

func resourceAliCloudDdosCooDomainPreciseAccessRuleDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteWebPreciseAccessRule"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RuleNames.1"] = parts[1]
	request["Domain"] = parts[0]
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("ddoscoo", "2020-01-01", action, query, request, true)
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
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
