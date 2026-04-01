// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudCloudFirewallUserAlarmConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCloudFirewallUserAlarmConfigCreate,
		Read:   resourceAliCloudCloudFirewallUserAlarmConfigRead,
		Update: resourceAliCloudCloudFirewallUserAlarmConfigUpdate,
		Delete: resourceAliCloudCloudFirewallUserAlarmConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"alarm_config": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alarm_period": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"alarm_hour": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"alarm_value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"alarm_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"alarm_week_day": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"alarm_notify": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"alarm_lang": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"contact_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"email": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"mobile_phone": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"lang": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_default_contact": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudCloudFirewallUserAlarmConfigCreate(d *schema.ResourceData, meta interface{}) error {

	return resourceAliCloudCloudFirewallUserAlarmConfigUpdate(d, meta)
}

func resourceAliCloudCloudFirewallUserAlarmConfigRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cloudFirewallServiceV2 := CloudFirewallServiceV2{client}

	objectRaw, err := cloudFirewallServiceV2.DescribeCloudFirewallUserAlarmConfig(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cloud_firewall_user_alarm_config DescribeCloudFirewallUserAlarmConfig Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("alarm_lang", objectRaw["AlarmLang"])

	alarmConfigRaw := objectRaw["AlarmConfig"]
	alarmConfigMaps := make([]map[string]interface{}, 0)
	if alarmConfigRaw != nil {
		for _, alarmConfigChildRaw := range convertToInterfaceArray(alarmConfigRaw) {
			alarmConfigMap := make(map[string]interface{})
			alarmConfigChildRaw := alarmConfigChildRaw.(map[string]interface{})
			alarmConfigMap["alarm_hour"] = alarmConfigChildRaw["AlarmHour"]
			alarmConfigMap["alarm_notify"] = alarmConfigChildRaw["AlarmNotify"]
			alarmConfigMap["alarm_period"] = alarmConfigChildRaw["AlarmPeriod"]
			alarmConfigMap["alarm_type"] = alarmConfigChildRaw["AlarmType"]
			alarmConfigMap["alarm_value"] = alarmConfigChildRaw["AlarmValue"]
			alarmConfigMap["alarm_week_day"] = alarmConfigChildRaw["AlarmWeekDay"]

			alarmConfigMaps = append(alarmConfigMaps, alarmConfigMap)
		}
	}
	if err := d.Set("alarm_config", alarmConfigMaps); err != nil {
		return err
	}
	contactConfigRaw := objectRaw["ContactConfig"]
	contactConfigMaps := make([]map[string]interface{}, 0)
	if contactConfigRaw != nil {
		for _, contactConfigChildRaw := range convertToInterfaceArray(contactConfigRaw) {
			contactConfigMap := make(map[string]interface{})
			contactConfigChildRaw := contactConfigChildRaw.(map[string]interface{})
			contactConfigMap["email"] = contactConfigChildRaw["Email"]
			contactConfigMap["mobile_phone"] = contactConfigChildRaw["MobilePhone"]
			contactConfigMap["name"] = contactConfigChildRaw["Name"]
			contactConfigMap["status"] = contactConfigChildRaw["Status"]

			contactConfigMaps = append(contactConfigMaps, contactConfigMap)
		}
	}
	if err := d.Set("contact_config", contactConfigMaps); err != nil {
		return err
	}

	return nil
}

func resourceAliCloudCloudFirewallUserAlarmConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "ModifyUserAlarmConfig"
	request = make(map[string]interface{})
	query = make(map[string]interface{})

	if v, ok := d.GetOk("use_default_contact"); ok {
		request["UseDefaultContact"] = v
	}
	if v, ok := d.GetOk("lang"); ok {
		request["Lang"] = v
	}
	if d.HasChange("alarm_config") {
		update = true
	}
	if v, ok := d.GetOk("alarm_config"); ok || d.HasChange("alarm_config") {
		alarmConfigMapsArray := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(v) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["AlarmValue"] = dataLoopTmp["alarm_value"]
			dataLoopMap["AlarmPeriod"] = dataLoopTmp["alarm_period"]
			dataLoopMap["AlarmWeekDay"] = dataLoopTmp["alarm_week_day"]
			dataLoopMap["AlarmNotify"] = dataLoopTmp["alarm_notify"]
			dataLoopMap["AlarmType"] = dataLoopTmp["alarm_type"]
			dataLoopMap["AlarmHour"] = dataLoopTmp["alarm_hour"]
			alarmConfigMapsArray = append(alarmConfigMapsArray, dataLoopMap)
		}
		request["AlarmConfig"] = alarmConfigMapsArray
	}

	if d.HasChange("contact_config") {
		update = true
	}
	if v, ok := d.GetOk("contact_config"); ok || d.HasChange("contact_config") {
		contactConfigMapsArray := make([]interface{}, 0)
		for _, dataLoop1 := range convertToInterfaceArray(v) {
			dataLoop1Tmp := dataLoop1.(map[string]interface{})
			dataLoop1Map := make(map[string]interface{})
			dataLoop1Map["Name"] = dataLoop1Tmp["name"]
			dataLoop1Map["MobilePhone"] = dataLoop1Tmp["mobile_phone"]
			dataLoop1Map["Email"] = dataLoop1Tmp["email"]
			dataLoop1Map["Status"] = dataLoop1Tmp["status"]
			contactConfigMapsArray = append(contactConfigMapsArray, dataLoop1Map)
		}
		contactConfigMapsJson, err := json.Marshal(contactConfigMapsArray)
		if err != nil {
			return WrapError(err)
		}
		request["ContactConfig"] = string(contactConfigMapsJson)
	}

	if d.HasChange("alarm_lang") {
		update = true
	}
	if v, ok := d.GetOk("alarm_lang"); ok || d.HasChange("alarm_lang") {
		request["AlarmLang"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 0*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Cloudfw", "2017-12-07", action, query, request, true)
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

	return resourceAliCloudCloudFirewallUserAlarmConfigRead(d, meta)
}

func resourceAliCloudCloudFirewallUserAlarmConfigDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource User Alarm Config. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
