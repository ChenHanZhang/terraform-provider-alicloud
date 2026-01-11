// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudSlsAlert() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudSlsAlertCreate,
		Read:   resourceAliCloudSlsAlertRead,
		Update: resourceAliCloudSlsAlertUpdate,
		Delete: resourceAliCloudSlsAlertDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"alert_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"configuration": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"severity_configurations": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"severity": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"eval_condition": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"condition": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"count_condition": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"auto_annotation": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"no_data_fire": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"sink_cms": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"dashboard": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"mute_until": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"template_configuration": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"annotations": {
										Type:     schema.TypeMap,
										Optional: true,
									},
									"version": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"lang": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"template_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"tokens": {
										Type:     schema.TypeMap,
										Optional: true,
									},
								},
							},
						},
						"labels": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"key": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"group_configuration": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: StringInSlice([]string{"no_group", "custom", "labels_auto"}, false),
									},
									"fields": {
										Type:     schema.TypeList,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"no_data_severity": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"annotations": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"key": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"condition_configuration": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"count_condition": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"join_configurations": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"condition": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: StringInSlice([]string{"cross_join", "inner_join", "left_join", "right_join", "full_join", "left_exclude", "right_exclude", "concat", "no_join"}, false),
									},
								},
							},
						},
						"policy_configuration": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alert_policy_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"action_policy_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"repeat_interval": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"sink_event_store": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"project": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"event_store": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"endpoint": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"role_arn": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"query_list": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"query": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"time_span_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"start": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"store": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"power_sql_mode": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"dashboard_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"role_arn": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"store_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"project": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"ui": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"region": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"end": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"chart_title": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"sink_alerthub": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"tags": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"send_resolved": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"threshold": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"create_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schedule": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: StringInSlice([]string{"FixedRate", "Cron"}, false),
						},
						"time_zone": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"run_immdiately": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"cron_expression": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"delay": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"interval": {
							Type:     schema.TypeString,
							Optional: true,
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
		},
	}
}

func resourceAliCloudSlsAlertCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/alerts")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	hostMap := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("alert_name"); ok {
		request["name"] = v
	}
	hostMap["project"] = StringPointer(d.Get("project_name").(string))

	schedule := make(map[string]interface{})

	if v := d.Get("schedule"); v != nil {
		cronExpression1, _ := jsonpath.Get("$[0].cron_expression", v)
		if cronExpression1 != nil && cronExpression1 != "" {
			schedule["cronExpression"] = cronExpression1
		}
		runImmdiately, _ := jsonpath.Get("$[0].run_immdiately", v)
		if runImmdiately != nil && runImmdiately != "" {
			schedule["runImmediately"] = runImmdiately
		}
		timeZone1, _ := jsonpath.Get("$[0].time_zone", v)
		if timeZone1 != nil && timeZone1 != "" {
			schedule["timeZone"] = timeZone1
		}
		interval1, _ := jsonpath.Get("$[0].interval", v)
		if interval1 != nil && interval1 != "" {
			schedule["interval"] = interval1
		}
		delay1, _ := jsonpath.Get("$[0].delay", v)
		if delay1 != nil && delay1 != "" {
			schedule["delay"] = delay1
		}
		type1, _ := jsonpath.Get("$[0].type", v)
		if type1 != nil && type1 != "" {
			schedule["type"] = type1
		}

		request["schedule"] = schedule
	}

	configuration := make(map[string]interface{})

	if v := d.Get("configuration"); v != nil {
		templateConfiguration := make(map[string]interface{})
		annotations, _ := jsonpath.Get("$[0].template_configuration[0].annotations", d.Get("configuration"))
		if annotations != nil && annotations != "" {
			templateConfiguration["aonotations"] = annotations
		}
		type3, _ := jsonpath.Get("$[0].template_configuration[0].type", d.Get("configuration"))
		if type3 != nil && type3 != "" {
			templateConfiguration["type"] = type3
		}
		tokens1, _ := jsonpath.Get("$[0].template_configuration[0].tokens", d.Get("configuration"))
		if tokens1 != nil && tokens1 != "" {
			templateConfiguration["tokens"] = tokens1
		}
		templateId, _ := jsonpath.Get("$[0].template_configuration[0].template_id", d.Get("configuration"))
		if templateId != nil && templateId != "" {
			templateConfiguration["id"] = templateId
		}
		lang1, _ := jsonpath.Get("$[0].template_configuration[0].lang", d.Get("configuration"))
		if lang1 != nil && lang1 != "" {
			templateConfiguration["lang"] = lang1
		}
		version1, _ := jsonpath.Get("$[0].template_configuration[0].version", d.Get("configuration"))
		if version1 != nil && version1 != "" {
			templateConfiguration["version"] = version1
		}

		if len(templateConfiguration) > 0 {
			configuration["templateConfiguration"] = templateConfiguration
		}
		type5, _ := jsonpath.Get("$[0].type", v)
		if type5 != nil && type5 != "" {
			configuration["type"] = type5
		}
		if v, ok := d.GetOk("configuration"); ok {
			localData, err := jsonpath.Get("$[0].query_list", v)
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
				dataLoopMap["chartTitle"] = dataLoopTmp["chart_title"]
				dataLoopMap["project"] = dataLoopTmp["project"]
				dataLoopMap["roleArn"] = dataLoopTmp["role_arn"]
				dataLoopMap["storeType"] = dataLoopTmp["store_type"]
				dataLoopMap["dashboardId"] = dataLoopTmp["dashboard_id"]
				dataLoopMap["region"] = dataLoopTmp["region"]
				dataLoopMap["ui"] = dataLoopTmp["ui"]
				dataLoopMap["query"] = dataLoopTmp["query"]
				dataLoopMap["start"] = dataLoopTmp["start"]
				dataLoopMap["store"] = dataLoopTmp["store"]
				dataLoopMap["powerSqlMode"] = dataLoopTmp["power_sql_mode"]
				dataLoopMap["end"] = dataLoopTmp["end"]
				dataLoopMap["timeSpanType"] = dataLoopTmp["time_span_type"]
				localMaps = append(localMaps, dataLoopMap)
			}
			configuration["queryList"] = localMaps
		}

		autoAnnotation1, _ := jsonpath.Get("$[0].auto_annotation", v)
		if autoAnnotation1 != nil && autoAnnotation1 != "" {
			configuration["autoAnnotation"] = autoAnnotation1
		}
		conditionConfiguration := make(map[string]interface{})
		countCondition1, _ := jsonpath.Get("$[0].condition_configuration[0].count_condition", d.Get("configuration"))
		if countCondition1 != nil && countCondition1 != "" {
			conditionConfiguration["countCondition"] = countCondition1
		}
		condition1, _ := jsonpath.Get("$[0].condition_configuration[0].condition", d.Get("configuration"))
		if condition1 != nil && condition1 != "" {
			conditionConfiguration["condition"] = condition1
		}

		if len(conditionConfiguration) > 0 {
			configuration["conditionConfiguration"] = conditionConfiguration
		}
		noDataSeverity1, _ := jsonpath.Get("$[0].no_data_severity", v)
		if noDataSeverity1 != nil && noDataSeverity1 != "" {
			configuration["noDataSeverity"] = noDataSeverity1
		}
		if v, ok := d.GetOk("configuration"); ok {
			localData1, err := jsonpath.Get("$[0].labels", v)
			if err != nil {
				localData1 = make([]interface{}, 0)
			}
			localMaps1 := make([]interface{}, 0)
			for _, dataLoop1 := range convertToInterfaceArray(localData1) {
				dataLoop1Tmp := make(map[string]interface{})
				if dataLoop1 != nil {
					dataLoop1Tmp = dataLoop1.(map[string]interface{})
				}
				dataLoop1Map := make(map[string]interface{})
				dataLoop1Map["key"] = dataLoop1Tmp["key"]
				dataLoop1Map["value"] = dataLoop1Tmp["value"]
				localMaps1 = append(localMaps1, dataLoop1Map)
			}
			configuration["labels"] = localMaps1
		}

		muteUntil1, _ := jsonpath.Get("$[0].mute_until", v)
		if muteUntil1 != nil && muteUntil1 != "" {
			configuration["muteUntil"] = muteUntil1
		}
		policyConfiguration := make(map[string]interface{})
		alertPolicyId1, _ := jsonpath.Get("$[0].policy_configuration[0].alert_policy_id", d.Get("configuration"))
		if alertPolicyId1 != nil && alertPolicyId1 != "" {
			policyConfiguration["alertPolicyId"] = alertPolicyId1
		}
		repeatInterval1, _ := jsonpath.Get("$[0].policy_configuration[0].repeat_interval", d.Get("configuration"))
		if repeatInterval1 != nil && repeatInterval1 != "" {
			policyConfiguration["repeatInterval"] = repeatInterval1
		}
		actionPolicyId1, _ := jsonpath.Get("$[0].policy_configuration[0].action_policy_id", d.Get("configuration"))
		if actionPolicyId1 != nil && actionPolicyId1 != "" {
			policyConfiguration["actionPolicyId"] = actionPolicyId1
		}

		if len(policyConfiguration) > 0 {
			configuration["policyConfiguration"] = policyConfiguration
		}
		noDataFire1, _ := jsonpath.Get("$[0].no_data_fire", v)
		if noDataFire1 != nil && noDataFire1 != "" {
			configuration["noDataFire"] = noDataFire1
		}
		sinkEventStore := make(map[string]interface{})
		enabled1, _ := jsonpath.Get("$[0].sink_event_store[0].enabled", d.Get("configuration"))
		if enabled1 != nil && enabled1 != "" {
			sinkEventStore["enabled"] = enabled1
		}
		eventStore1, _ := jsonpath.Get("$[0].sink_event_store[0].event_store", d.Get("configuration"))
		if eventStore1 != nil && eventStore1 != "" {
			sinkEventStore["eventStore"] = eventStore1
		}
		endpoint1, _ := jsonpath.Get("$[0].sink_event_store[0].endpoint", d.Get("configuration"))
		if endpoint1 != nil && endpoint1 != "" {
			sinkEventStore["endpoint"] = endpoint1
		}
		roleArn3, _ := jsonpath.Get("$[0].sink_event_store[0].role_arn", d.Get("configuration"))
		if roleArn3 != nil && roleArn3 != "" {
			sinkEventStore["roleArn"] = roleArn3
		}
		project3, _ := jsonpath.Get("$[0].sink_event_store[0].project", d.Get("configuration"))
		if project3 != nil && project3 != "" {
			sinkEventStore["project"] = project3
		}

		if len(sinkEventStore) > 0 {
			configuration["sinkEventStore"] = sinkEventStore
		}
		if v, ok := d.GetOk("configuration"); ok {
			localData2, err := jsonpath.Get("$[0].severity_configurations", v)
			if err != nil {
				localData2 = make([]interface{}, 0)
			}
			localMaps2 := make([]interface{}, 0)
			for _, dataLoop2 := range convertToInterfaceArray(localData2) {
				dataLoop2Tmp := make(map[string]interface{})
				if dataLoop2 != nil {
					dataLoop2Tmp = dataLoop2.(map[string]interface{})
				}
				dataLoop2Map := make(map[string]interface{})
				localData3 := make(map[string]interface{})
				countCondition3, _ := jsonpath.Get("$[0].count_condition", dataLoop2Tmp["eval_condition"])
				if countCondition3 != nil && countCondition3 != "" {
					localData3["countCondition"] = countCondition3
				}
				condition3, _ := jsonpath.Get("$[0].condition", dataLoop2Tmp["eval_condition"])
				if condition3 != nil && condition3 != "" {
					localData3["condition"] = condition3
				}
				if len(localData3) > 0 {
					dataLoop2Map["evalCondition"] = localData3
				}
				dataLoop2Map["severity"] = dataLoop2Tmp["severity"]
				localMaps2 = append(localMaps2, dataLoop2Map)
			}
			configuration["severityConfigurations"] = localMaps2
		}

		version3, _ := jsonpath.Get("$[0].version", v)
		if version3 != nil && version3 != "" {
			configuration["version"] = version3
		}
		sinkCms := make(map[string]interface{})
		enabled3, _ := jsonpath.Get("$[0].sink_cms[0].enabled", d.Get("configuration"))
		if enabled3 != nil && enabled3 != "" {
			sinkCms["enabled"] = enabled3
		}

		if len(sinkCms) > 0 {
			configuration["sinkCms"] = sinkCms
		}
		groupConfiguration := make(map[string]interface{})
		fields1, _ := jsonpath.Get("$[0].group_configuration[0].fields", d.Get("configuration"))
		if fields1 != nil && fields1 != "" {
			groupConfiguration["fields"] = fields1
		}
		type7, _ := jsonpath.Get("$[0].group_configuration[0].type", d.Get("configuration"))
		if type7 != nil && type7 != "" {
			groupConfiguration["type"] = type7
		}

		if len(groupConfiguration) > 0 {
			configuration["groupConfiguration"] = groupConfiguration
		}
		if v, ok := d.GetOk("configuration"); ok {
			localData4, err := jsonpath.Get("$[0].annotations", v)
			if err != nil {
				localData4 = make([]interface{}, 0)
			}
			localMaps4 := make([]interface{}, 0)
			for _, dataLoop4 := range convertToInterfaceArray(localData4) {
				dataLoop4Tmp := make(map[string]interface{})
				if dataLoop4 != nil {
					dataLoop4Tmp = dataLoop4.(map[string]interface{})
				}
				dataLoop4Map := make(map[string]interface{})
				dataLoop4Map["key"] = dataLoop4Tmp["key"]
				dataLoop4Map["value"] = dataLoop4Tmp["value"]
				localMaps4 = append(localMaps4, dataLoop4Map)
			}
			configuration["annotations"] = localMaps4
		}

		sinkAlerthub := make(map[string]interface{})
		enabled5, _ := jsonpath.Get("$[0].sink_alerthub[0].enabled", d.Get("configuration"))
		if enabled5 != nil && enabled5 != "" {
			sinkAlerthub["enabled"] = enabled5
		}

		if len(sinkAlerthub) > 0 {
			configuration["sinkAlerthub"] = sinkAlerthub
		}
		if v, ok := d.GetOk("configuration"); ok {
			localData5, err := jsonpath.Get("$[0].join_configurations", v)
			if err != nil {
				localData5 = make([]interface{}, 0)
			}
			localMaps5 := make([]interface{}, 0)
			for _, dataLoop5 := range convertToInterfaceArray(localData5) {
				dataLoop5Tmp := make(map[string]interface{})
				if dataLoop5 != nil {
					dataLoop5Tmp = dataLoop5.(map[string]interface{})
				}
				dataLoop5Map := make(map[string]interface{})
				dataLoop5Map["condition"] = dataLoop5Tmp["condition"]
				dataLoop5Map["type"] = dataLoop5Tmp["type"]
				localMaps5 = append(localMaps5, dataLoop5Map)
			}
			configuration["joinConfigurations"] = localMaps5
		}

		dashboard1, _ := jsonpath.Get("$[0].dashboard", v)
		if dashboard1 != nil && dashboard1 != "" {
			configuration["dashboard"] = dashboard1
		}
		tags1, _ := jsonpath.Get("$[0].tags", v)
		if tags1 != nil && tags1 != "" {
			configuration["tags"] = tags1
		}
		threshold1, _ := jsonpath.Get("$[0].threshold", v)
		if threshold1 != nil && threshold1 != "" {
			configuration["threshold"] = threshold1
		}
		sendResolved1, _ := jsonpath.Get("$[0].send_resolved", v)
		if sendResolved1 != nil && sendResolved1 != "" {
			configuration["sendResolved"] = sendResolved1
		}

		request["configuration"] = configuration
	}

	if v, ok := d.GetOk("description"); ok {
		request["description"] = v
	}
	request["displayName"] = d.Get("display_name")
	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.Do("Sls", roaParam("POST", "2020-12-30", "CreateAlert", action), query, body, nil, hostMap, false)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_sls_alert", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", *hostMap["project"], request["name"]))

	return resourceAliCloudSlsAlertUpdate(d, meta)
}

func resourceAliCloudSlsAlertRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	slsServiceV2 := SlsServiceV2{client}

	objectRaw, err := slsServiceV2.DescribeSlsAlert(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_sls_alert DescribeSlsAlert Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["createTime"])
	d.Set("description", objectRaw["description"])
	d.Set("display_name", objectRaw["displayName"])
	d.Set("status", objectRaw["status"])
	d.Set("alert_name", objectRaw["name"])

	configurationMaps := make([]map[string]interface{}, 0)
	configurationMap := make(map[string]interface{})
	configurationRaw := make(map[string]interface{})
	if objectRaw["configuration"] != nil {
		configurationRaw = objectRaw["configuration"].(map[string]interface{})
	}
	if len(configurationRaw) > 0 {
		configurationMap["auto_annotation"] = configurationRaw["autoAnnotation"]
		configurationMap["dashboard"] = configurationRaw["dashboard"]
		configurationMap["mute_until"] = configurationRaw["muteUntil"]
		configurationMap["no_data_fire"] = configurationRaw["noDataFire"]
		configurationMap["no_data_severity"] = configurationRaw["noDataSeverity"]
		configurationMap["send_resolved"] = configurationRaw["sendResolved"]
		configurationMap["threshold"] = configurationRaw["threshold"]
		configurationMap["type"] = configurationRaw["type"]
		configurationMap["version"] = configurationRaw["version"]

		annotationsRaw := configurationRaw["annotations"]
		annotationsMaps := make([]map[string]interface{}, 0)
		if annotationsRaw != nil {
			for _, annotationsChildRaw := range convertToInterfaceArray(annotationsRaw) {
				annotationsMap := make(map[string]interface{})
				annotationsChildRaw := annotationsChildRaw.(map[string]interface{})
				annotationsMap["key"] = annotationsChildRaw["key"]
				annotationsMap["value"] = annotationsChildRaw["value"]

				annotationsMaps = append(annotationsMaps, annotationsMap)
			}
		}
		configurationMap["annotations"] = annotationsMaps
		conditionConfigurationMaps := make([]map[string]interface{}, 0)
		conditionConfigurationMap := make(map[string]interface{})
		conditionConfigurationRaw := make(map[string]interface{})
		if configurationRaw["conditionConfiguration"] != nil {
			conditionConfigurationRaw = configurationRaw["conditionConfiguration"].(map[string]interface{})
		}
		if len(conditionConfigurationRaw) > 0 {
			conditionConfigurationMap["condition"] = conditionConfigurationRaw["condition"]
			conditionConfigurationMap["count_condition"] = conditionConfigurationRaw["countCondition"]

			conditionConfigurationMaps = append(conditionConfigurationMaps, conditionConfigurationMap)
		}
		configurationMap["condition_configuration"] = conditionConfigurationMaps
		groupConfigurationMaps := make([]map[string]interface{}, 0)
		groupConfigurationMap := make(map[string]interface{})
		groupConfigurationRaw := make(map[string]interface{})
		if configurationRaw["groupConfiguration"] != nil {
			groupConfigurationRaw = configurationRaw["groupConfiguration"].(map[string]interface{})
		}
		if len(groupConfigurationRaw) > 0 {
			groupConfigurationMap["type"] = groupConfigurationRaw["type"]

			fieldsRaw := make([]interface{}, 0)
			if groupConfigurationRaw["fields"] != nil {
				fieldsRaw = convertToInterfaceArray(groupConfigurationRaw["fields"])
			}

			groupConfigurationMap["fields"] = fieldsRaw
			groupConfigurationMaps = append(groupConfigurationMaps, groupConfigurationMap)
		}
		configurationMap["group_configuration"] = groupConfigurationMaps
		joinConfigurationsRaw := configurationRaw["joinConfigurations"]
		joinConfigurationsMaps := make([]map[string]interface{}, 0)
		if joinConfigurationsRaw != nil {
			for _, joinConfigurationsChildRaw := range convertToInterfaceArray(joinConfigurationsRaw) {
				joinConfigurationsMap := make(map[string]interface{})
				joinConfigurationsChildRaw := joinConfigurationsChildRaw.(map[string]interface{})
				joinConfigurationsMap["condition"] = joinConfigurationsChildRaw["condition"]
				joinConfigurationsMap["type"] = joinConfigurationsChildRaw["type"]

				joinConfigurationsMaps = append(joinConfigurationsMaps, joinConfigurationsMap)
			}
		}
		configurationMap["join_configurations"] = joinConfigurationsMaps
		labelsRaw := configurationRaw["labels"]
		labelsMaps := make([]map[string]interface{}, 0)
		if labelsRaw != nil {
			for _, labelsChildRaw := range convertToInterfaceArray(labelsRaw) {
				labelsMap := make(map[string]interface{})
				labelsChildRaw := labelsChildRaw.(map[string]interface{})
				labelsMap["key"] = labelsChildRaw["key"]
				labelsMap["value"] = labelsChildRaw["value"]

				labelsMaps = append(labelsMaps, labelsMap)
			}
		}
		configurationMap["labels"] = labelsMaps
		policyConfigurationMaps := make([]map[string]interface{}, 0)
		policyConfigurationMap := make(map[string]interface{})
		policyConfigurationRaw := make(map[string]interface{})
		if configurationRaw["policyConfiguration"] != nil {
			policyConfigurationRaw = configurationRaw["policyConfiguration"].(map[string]interface{})
		}
		if len(policyConfigurationRaw) > 0 {
			policyConfigurationMap["action_policy_id"] = policyConfigurationRaw["actionPolicyId"]
			policyConfigurationMap["alert_policy_id"] = policyConfigurationRaw["alertPolicyId"]
			policyConfigurationMap["repeat_interval"] = policyConfigurationRaw["repeatInterval"]

			policyConfigurationMaps = append(policyConfigurationMaps, policyConfigurationMap)
		}
		configurationMap["policy_configuration"] = policyConfigurationMaps
		queryListRaw := configurationRaw["queryList"]
		queryListMaps := make([]map[string]interface{}, 0)
		if queryListRaw != nil {
			for _, queryListChildRaw := range convertToInterfaceArray(queryListRaw) {
				queryListMap := make(map[string]interface{})
				queryListChildRaw := queryListChildRaw.(map[string]interface{})
				queryListMap["chart_title"] = queryListChildRaw["chartTitle"]
				queryListMap["dashboard_id"] = queryListChildRaw["dashboardId"]
				queryListMap["end"] = queryListChildRaw["end"]
				queryListMap["power_sql_mode"] = queryListChildRaw["powerSqlMode"]
				queryListMap["project"] = queryListChildRaw["project"]
				queryListMap["query"] = queryListChildRaw["query"]
				queryListMap["region"] = queryListChildRaw["region"]
				queryListMap["role_arn"] = queryListChildRaw["roleArn"]
				queryListMap["start"] = queryListChildRaw["start"]
				queryListMap["store"] = queryListChildRaw["store"]
				queryListMap["store_type"] = queryListChildRaw["storeType"]
				queryListMap["time_span_type"] = queryListChildRaw["timeSpanType"]
				queryListMap["ui"] = queryListChildRaw["ui"]

				queryListMaps = append(queryListMaps, queryListMap)
			}
		}
		configurationMap["query_list"] = queryListMaps
		severityConfigurationsRaw := configurationRaw["severityConfigurations"]
		severityConfigurationsMaps := make([]map[string]interface{}, 0)
		if severityConfigurationsRaw != nil {
			for _, severityConfigurationsChildRaw := range convertToInterfaceArray(severityConfigurationsRaw) {
				severityConfigurationsMap := make(map[string]interface{})
				severityConfigurationsChildRaw := severityConfigurationsChildRaw.(map[string]interface{})
				severityConfigurationsMap["severity"] = severityConfigurationsChildRaw["severity"]

				evalConditionMaps := make([]map[string]interface{}, 0)
				evalConditionMap := make(map[string]interface{})
				evalConditionRaw := make(map[string]interface{})
				if severityConfigurationsChildRaw["evalCondition"] != nil {
					evalConditionRaw = severityConfigurationsChildRaw["evalCondition"].(map[string]interface{})
				}
				if len(evalConditionRaw) > 0 {
					evalConditionMap["condition"] = evalConditionRaw["condition"]
					evalConditionMap["count_condition"] = evalConditionRaw["countCondition"]

					evalConditionMaps = append(evalConditionMaps, evalConditionMap)
				}
				severityConfigurationsMap["eval_condition"] = evalConditionMaps
				severityConfigurationsMaps = append(severityConfigurationsMaps, severityConfigurationsMap)
			}
		}
		configurationMap["severity_configurations"] = severityConfigurationsMaps
		sinkAlerthubMaps := make([]map[string]interface{}, 0)
		sinkAlerthubMap := make(map[string]interface{})
		sinkAlerthubRaw := make(map[string]interface{})
		if configurationRaw["sinkAlerthub"] != nil {
			sinkAlerthubRaw = configurationRaw["sinkAlerthub"].(map[string]interface{})
		}
		if len(sinkAlerthubRaw) > 0 {
			sinkAlerthubMap["enabled"] = sinkAlerthubRaw["enabled"]

			sinkAlerthubMaps = append(sinkAlerthubMaps, sinkAlerthubMap)
		}
		configurationMap["sink_alerthub"] = sinkAlerthubMaps
		sinkCmsMaps := make([]map[string]interface{}, 0)
		sinkCmsMap := make(map[string]interface{})
		sinkCmsRaw := make(map[string]interface{})
		if configurationRaw["sinkCms"] != nil {
			sinkCmsRaw = configurationRaw["sinkCms"].(map[string]interface{})
		}
		if len(sinkCmsRaw) > 0 {
			sinkCmsMap["enabled"] = sinkCmsRaw["enabled"]

			sinkCmsMaps = append(sinkCmsMaps, sinkCmsMap)
		}
		configurationMap["sink_cms"] = sinkCmsMaps
		sinkEventStoreMaps := make([]map[string]interface{}, 0)
		sinkEventStoreMap := make(map[string]interface{})
		sinkEventStoreRaw := make(map[string]interface{})
		if configurationRaw["sinkEventStore"] != nil {
			sinkEventStoreRaw = configurationRaw["sinkEventStore"].(map[string]interface{})
		}
		if len(sinkEventStoreRaw) > 0 {
			sinkEventStoreMap["enabled"] = sinkEventStoreRaw["enabled"]
			sinkEventStoreMap["endpoint"] = sinkEventStoreRaw["endpoint"]
			sinkEventStoreMap["event_store"] = sinkEventStoreRaw["eventStore"]
			sinkEventStoreMap["project"] = sinkEventStoreRaw["project"]
			sinkEventStoreMap["role_arn"] = sinkEventStoreRaw["roleArn"]

			sinkEventStoreMaps = append(sinkEventStoreMaps, sinkEventStoreMap)
		}
		configurationMap["sink_event_store"] = sinkEventStoreMaps
		tagsRaw := make([]interface{}, 0)
		if configurationRaw["tags"] != nil {
			tagsRaw = convertToInterfaceArray(configurationRaw["tags"])
		}

		configurationMap["tags"] = tagsRaw
		templateConfigurationMaps := make([]map[string]interface{}, 0)
		templateConfigurationMap := make(map[string]interface{})
		templateConfigurationRaw := make(map[string]interface{})
		if configurationRaw["templateConfiguration"] != nil {
			templateConfigurationRaw = configurationRaw["templateConfiguration"].(map[string]interface{})
		}
		if len(templateConfigurationRaw) > 0 {
			templateConfigurationMap["annotations"] = templateConfigurationRaw["aonotations"]
			templateConfigurationMap["lang"] = templateConfigurationRaw["lang"]
			templateConfigurationMap["template_id"] = templateConfigurationRaw["id"]
			templateConfigurationMap["tokens"] = templateConfigurationRaw["tokens"]
			templateConfigurationMap["type"] = templateConfigurationRaw["type"]
			templateConfigurationMap["version"] = templateConfigurationRaw["version"]

			templateConfigurationMaps = append(templateConfigurationMaps, templateConfigurationMap)
		}
		configurationMap["template_configuration"] = templateConfigurationMaps
		configurationMaps = append(configurationMaps, configurationMap)
	}
	if err := d.Set("configuration", configurationMaps); err != nil {
		return err
	}
	scheduleMaps := make([]map[string]interface{}, 0)
	scheduleMap := make(map[string]interface{})
	scheduleRaw := make(map[string]interface{})
	if objectRaw["schedule"] != nil {
		scheduleRaw = objectRaw["schedule"].(map[string]interface{})
	}
	if len(scheduleRaw) > 0 {
		scheduleMap["cron_expression"] = scheduleRaw["cronExpression"]
		scheduleMap["delay"] = scheduleRaw["delay"]
		scheduleMap["interval"] = scheduleRaw["interval"]
		scheduleMap["run_immdiately"] = scheduleRaw["runImmediately"]
		scheduleMap["time_zone"] = scheduleRaw["timeZone"]
		scheduleMap["type"] = scheduleRaw["type"]

		scheduleMaps = append(scheduleMaps, scheduleMap)
	}
	if err := d.Set("schedule", scheduleMaps); err != nil {
		return err
	}

	parts := strings.Split(d.Id(), ":")
	d.Set("project_name", parts[0])

	return nil
}

func resourceAliCloudSlsAlertUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	update := false

	slsServiceV2 := SlsServiceV2{client}
	objectRaw, _ := slsServiceV2.DescribeSlsAlert(d.Id())

	if d.HasChange("status") {
		var err error
		target := d.Get("status").(string)

		currentStatus, err := jsonpath.Get("status", objectRaw)
		if err != nil {
			return WrapErrorf(err, FailedGetAttributeMsg, d.Id(), "status", objectRaw)
		}
		if fmt.Sprint(currentStatus) != target {
			if target == "ENABLED" {
				parts := strings.Split(d.Id(), ":")
				alertName := parts[1]
				action := fmt.Sprintf("/alerts/%s?action=enable", alertName)
				request = make(map[string]interface{})
				query = make(map[string]*string)
				body = make(map[string]interface{})
				hostMap := make(map[string]*string)
				hostMap["project"] = StringPointer(parts[0])

				body = request
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.Do("Sls", roaParam("PUT", "2020-12-30", "EnableAlert", action), query, body, nil, hostMap, false)
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
			if target == "DISABLED" {
				parts := strings.Split(d.Id(), ":")
				alertName := parts[1]
				action := fmt.Sprintf("/alerts/%s?action=disable", alertName)
				request = make(map[string]interface{})
				query = make(map[string]*string)
				body = make(map[string]interface{})
				hostMap := make(map[string]*string)
				hostMap["project"] = StringPointer(parts[0])

				body = request
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.Do("Sls", roaParam("PUT", "2020-12-30", "DisableAlert", action), query, body, nil, hostMap, false)
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
		}
	}

	var err error
	parts := strings.Split(d.Id(), ":")
	alertName := parts[1]
	action := fmt.Sprintf("/alerts/%s", alertName)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	hostMap := make(map[string]*string)
	hostMap["project"] = StringPointer(parts[0])

	if !d.IsNewResource() && d.HasChange("schedule") {
		update = true
	}
	schedule := make(map[string]interface{})

	if v := d.Get("schedule"); v != nil {
		cronExpression1, _ := jsonpath.Get("$[0].cron_expression", v)
		if cronExpression1 != nil && cronExpression1 != "" {
			schedule["cronExpression"] = cronExpression1
		}
		runImmdiately, _ := jsonpath.Get("$[0].run_immdiately", v)
		if runImmdiately != nil && runImmdiately != "" {
			schedule["runImmediately"] = runImmdiately
		}
		timeZone1, _ := jsonpath.Get("$[0].time_zone", v)
		if timeZone1 != nil && timeZone1 != "" {
			schedule["timeZone"] = timeZone1
		}
		interval1, _ := jsonpath.Get("$[0].interval", v)
		if interval1 != nil && interval1 != "" {
			schedule["interval"] = interval1
		}
		delay1, _ := jsonpath.Get("$[0].delay", v)
		if delay1 != nil && delay1 != "" {
			schedule["delay"] = delay1
		}
		type1, _ := jsonpath.Get("$[0].type", v)
		if type1 != nil && type1 != "" {
			schedule["type"] = type1
		}

		request["schedule"] = schedule
	}

	if d.HasChange("configuration") {
		update = true
	}
	configuration := make(map[string]interface{})

	if v := d.Get("configuration"); v != nil {
		templateConfiguration := make(map[string]interface{})
		annotations, _ := jsonpath.Get("$[0].template_configuration[0].annotations", d.Get("configuration"))
		if annotations != nil && annotations != "" {
			templateConfiguration["aonotations"] = annotations
		}
		type3, _ := jsonpath.Get("$[0].template_configuration[0].type", d.Get("configuration"))
		if type3 != nil && type3 != "" {
			templateConfiguration["type"] = type3
		}
		tokens1, _ := jsonpath.Get("$[0].template_configuration[0].tokens", d.Get("configuration"))
		if tokens1 != nil && tokens1 != "" {
			templateConfiguration["tokens"] = tokens1
		}
		templateId, _ := jsonpath.Get("$[0].template_configuration[0].template_id", d.Get("configuration"))
		if templateId != nil && templateId != "" {
			templateConfiguration["id"] = templateId
		}
		lang1, _ := jsonpath.Get("$[0].template_configuration[0].lang", d.Get("configuration"))
		if lang1 != nil && lang1 != "" {
			templateConfiguration["lang"] = lang1
		}
		version1, _ := jsonpath.Get("$[0].template_configuration[0].version", d.Get("configuration"))
		if version1 != nil && version1 != "" {
			templateConfiguration["version"] = version1
		}

		if len(templateConfiguration) > 0 {
			configuration["templateConfiguration"] = templateConfiguration
		}
		type5, _ := jsonpath.Get("$[0].type", v)
		if type5 != nil && type5 != "" {
			configuration["type"] = type5
		}
		if v, ok := d.GetOk("configuration"); ok {
			localData, err := jsonpath.Get("$[0].query_list", v)
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
				dataLoopMap["chartTitle"] = dataLoopTmp["chart_title"]
				dataLoopMap["project"] = dataLoopTmp["project"]
				dataLoopMap["roleArn"] = dataLoopTmp["role_arn"]
				dataLoopMap["storeType"] = dataLoopTmp["store_type"]
				dataLoopMap["dashboardId"] = dataLoopTmp["dashboard_id"]
				dataLoopMap["region"] = dataLoopTmp["region"]
				dataLoopMap["ui"] = dataLoopTmp["ui"]
				dataLoopMap["query"] = dataLoopTmp["query"]
				dataLoopMap["start"] = dataLoopTmp["start"]
				dataLoopMap["store"] = dataLoopTmp["store"]
				dataLoopMap["powerSqlMode"] = dataLoopTmp["power_sql_mode"]
				dataLoopMap["end"] = dataLoopTmp["end"]
				dataLoopMap["timeSpanType"] = dataLoopTmp["time_span_type"]
				localMaps = append(localMaps, dataLoopMap)
			}
			configuration["queryList"] = localMaps
		}

		autoAnnotation1, _ := jsonpath.Get("$[0].auto_annotation", v)
		if autoAnnotation1 != nil && autoAnnotation1 != "" {
			configuration["autoAnnotation"] = autoAnnotation1
		}
		conditionConfiguration := make(map[string]interface{})
		countCondition1, _ := jsonpath.Get("$[0].condition_configuration[0].count_condition", d.Get("configuration"))
		if countCondition1 != nil && countCondition1 != "" {
			conditionConfiguration["countCondition"] = countCondition1
		}
		condition1, _ := jsonpath.Get("$[0].condition_configuration[0].condition", d.Get("configuration"))
		if condition1 != nil && condition1 != "" {
			conditionConfiguration["condition"] = condition1
		}

		if len(conditionConfiguration) > 0 {
			configuration["conditionConfiguration"] = conditionConfiguration
		}
		noDataSeverity1, _ := jsonpath.Get("$[0].no_data_severity", v)
		if noDataSeverity1 != nil && noDataSeverity1 != "" {
			configuration["noDataSeverity"] = noDataSeverity1
		}
		if v, ok := d.GetOk("configuration"); ok {
			localData1, err := jsonpath.Get("$[0].labels", v)
			if err != nil {
				localData1 = make([]interface{}, 0)
			}
			localMaps1 := make([]interface{}, 0)
			for _, dataLoop1 := range convertToInterfaceArray(localData1) {
				dataLoop1Tmp := make(map[string]interface{})
				if dataLoop1 != nil {
					dataLoop1Tmp = dataLoop1.(map[string]interface{})
				}
				dataLoop1Map := make(map[string]interface{})
				dataLoop1Map["key"] = dataLoop1Tmp["key"]
				dataLoop1Map["value"] = dataLoop1Tmp["value"]
				localMaps1 = append(localMaps1, dataLoop1Map)
			}
			configuration["labels"] = localMaps1
		}

		muteUntil1, _ := jsonpath.Get("$[0].mute_until", v)
		if muteUntil1 != nil && muteUntil1 != "" {
			configuration["muteUntil"] = muteUntil1
		}
		policyConfiguration := make(map[string]interface{})
		alertPolicyId1, _ := jsonpath.Get("$[0].policy_configuration[0].alert_policy_id", d.Get("configuration"))
		if alertPolicyId1 != nil && alertPolicyId1 != "" {
			policyConfiguration["alertPolicyId"] = alertPolicyId1
		}
		repeatInterval1, _ := jsonpath.Get("$[0].policy_configuration[0].repeat_interval", d.Get("configuration"))
		if repeatInterval1 != nil && repeatInterval1 != "" {
			policyConfiguration["repeatInterval"] = repeatInterval1
		}
		actionPolicyId1, _ := jsonpath.Get("$[0].policy_configuration[0].action_policy_id", d.Get("configuration"))
		if actionPolicyId1 != nil && actionPolicyId1 != "" {
			policyConfiguration["actionPolicyId"] = actionPolicyId1
		}

		if len(policyConfiguration) > 0 {
			configuration["policyConfiguration"] = policyConfiguration
		}
		noDataFire1, _ := jsonpath.Get("$[0].no_data_fire", v)
		if noDataFire1 != nil && noDataFire1 != "" {
			configuration["noDataFire"] = noDataFire1
		}
		sinkEventStore := make(map[string]interface{})
		enabled1, _ := jsonpath.Get("$[0].sink_event_store[0].enabled", d.Get("configuration"))
		if enabled1 != nil && enabled1 != "" {
			sinkEventStore["enabled"] = enabled1
		}
		eventStore1, _ := jsonpath.Get("$[0].sink_event_store[0].event_store", d.Get("configuration"))
		if eventStore1 != nil && eventStore1 != "" {
			sinkEventStore["eventStore"] = eventStore1
		}
		endpoint1, _ := jsonpath.Get("$[0].sink_event_store[0].endpoint", d.Get("configuration"))
		if endpoint1 != nil && endpoint1 != "" {
			sinkEventStore["endpoint"] = endpoint1
		}
		roleArn3, _ := jsonpath.Get("$[0].sink_event_store[0].role_arn", d.Get("configuration"))
		if roleArn3 != nil && roleArn3 != "" {
			sinkEventStore["roleArn"] = roleArn3
		}
		project3, _ := jsonpath.Get("$[0].sink_event_store[0].project", d.Get("configuration"))
		if project3 != nil && project3 != "" {
			sinkEventStore["project"] = project3
		}

		if len(sinkEventStore) > 0 {
			configuration["sinkEventStore"] = sinkEventStore
		}
		if v, ok := d.GetOk("configuration"); ok {
			localData2, err := jsonpath.Get("$[0].severity_configurations", v)
			if err != nil {
				localData2 = make([]interface{}, 0)
			}
			localMaps2 := make([]interface{}, 0)
			for _, dataLoop2 := range convertToInterfaceArray(localData2) {
				dataLoop2Tmp := make(map[string]interface{})
				if dataLoop2 != nil {
					dataLoop2Tmp = dataLoop2.(map[string]interface{})
				}
				dataLoop2Map := make(map[string]interface{})
				if !IsNil(dataLoop2Tmp["eval_condition"]) {
					localData3 := make(map[string]interface{})
					countCondition3, _ := jsonpath.Get("$[0].count_condition", dataLoop2Tmp["eval_condition"])
					if countCondition3 != nil && countCondition3 != "" {
						localData3["countCondition"] = countCondition3
					}
					condition3, _ := jsonpath.Get("$[0].condition", dataLoop2Tmp["eval_condition"])
					if condition3 != nil && condition3 != "" {
						localData3["condition"] = condition3
					}
					if len(localData3) > 0 {
						dataLoop2Map["evalCondition"] = localData3
					}
				}
				dataLoop2Map["severity"] = dataLoop2Tmp["severity"]
				localMaps2 = append(localMaps2, dataLoop2Map)
			}
			configuration["severityConfigurations"] = localMaps2
		}

		version3, _ := jsonpath.Get("$[0].version", v)
		if version3 != nil && version3 != "" {
			configuration["version"] = version3
		}
		sinkCms := make(map[string]interface{})
		enabled3, _ := jsonpath.Get("$[0].sink_cms[0].enabled", d.Get("configuration"))
		if enabled3 != nil && enabled3 != "" {
			sinkCms["enabled"] = enabled3
		}

		if len(sinkCms) > 0 {
			configuration["sinkCms"] = sinkCms
		}
		groupConfiguration := make(map[string]interface{})
		fields1, _ := jsonpath.Get("$[0].group_configuration[0].fields", d.Get("configuration"))
		if fields1 != nil && fields1 != "" {
			groupConfiguration["fields"] = fields1
		}
		type7, _ := jsonpath.Get("$[0].group_configuration[0].type", d.Get("configuration"))
		if type7 != nil && type7 != "" {
			groupConfiguration["type"] = type7
		}

		if len(groupConfiguration) > 0 {
			configuration["groupConfiguration"] = groupConfiguration
		}
		if v, ok := d.GetOk("configuration"); ok {
			localData4, err := jsonpath.Get("$[0].annotations", v)
			if err != nil {
				localData4 = make([]interface{}, 0)
			}
			localMaps4 := make([]interface{}, 0)
			for _, dataLoop4 := range convertToInterfaceArray(localData4) {
				dataLoop4Tmp := make(map[string]interface{})
				if dataLoop4 != nil {
					dataLoop4Tmp = dataLoop4.(map[string]interface{})
				}
				dataLoop4Map := make(map[string]interface{})
				dataLoop4Map["key"] = dataLoop4Tmp["key"]
				dataLoop4Map["value"] = dataLoop4Tmp["value"]
				localMaps4 = append(localMaps4, dataLoop4Map)
			}
			configuration["annotations"] = localMaps4
		}

		sinkAlerthub := make(map[string]interface{})
		enabled5, _ := jsonpath.Get("$[0].sink_alerthub[0].enabled", d.Get("configuration"))
		if enabled5 != nil && enabled5 != "" {
			sinkAlerthub["enabled"] = enabled5
		}

		if len(sinkAlerthub) > 0 {
			configuration["sinkAlerthub"] = sinkAlerthub
		}
		if v, ok := d.GetOk("configuration"); ok {
			localData5, err := jsonpath.Get("$[0].join_configurations", v)
			if err != nil {
				localData5 = make([]interface{}, 0)
			}
			localMaps5 := make([]interface{}, 0)
			for _, dataLoop5 := range convertToInterfaceArray(localData5) {
				dataLoop5Tmp := make(map[string]interface{})
				if dataLoop5 != nil {
					dataLoop5Tmp = dataLoop5.(map[string]interface{})
				}
				dataLoop5Map := make(map[string]interface{})
				dataLoop5Map["condition"] = dataLoop5Tmp["condition"]
				dataLoop5Map["type"] = dataLoop5Tmp["type"]
				localMaps5 = append(localMaps5, dataLoop5Map)
			}
			configuration["joinConfigurations"] = localMaps5
		}

		dashboard1, _ := jsonpath.Get("$[0].dashboard", v)
		if dashboard1 != nil && dashboard1 != "" {
			configuration["dashboard"] = dashboard1
		}
		tags1, _ := jsonpath.Get("$[0].tags", v)
		if tags1 != nil && tags1 != "" {
			configuration["tags"] = tags1
		}
		threshold1, _ := jsonpath.Get("$[0].threshold", v)
		if threshold1 != nil && threshold1 != "" {
			configuration["threshold"] = threshold1
		}
		sendResolved1, _ := jsonpath.Get("$[0].send_resolved", v)
		if sendResolved1 != nil && sendResolved1 != "" {
			configuration["sendResolved"] = sendResolved1
		}

		request["configuration"] = configuration
	}

	if !d.IsNewResource() && d.HasChange("description") {
		update = true
	}
	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		request["description"] = v
	}
	if !d.IsNewResource() && d.HasChange("display_name") {
		update = true
	}
	request["displayName"] = d.Get("display_name")
	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.Do("Sls", roaParam("PUT", "2020-12-30", "UpdateAlert", action), query, body, nil, hostMap, false)
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

	return resourceAliCloudSlsAlertRead(d, meta)
}

func resourceAliCloudSlsAlertDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	alertName := parts[1]
	action := fmt.Sprintf("/alerts/%s", alertName)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	hostMap := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	hostMap["project"] = StringPointer(parts[0])

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.Do("Sls", roaParam("DELETE", "2020-12-30", "DeleteAlert", action), query, nil, nil, hostMap, false)
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
