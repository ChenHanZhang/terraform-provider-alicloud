// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
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
						},
						"alarm_hour": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"alarm_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"alarm_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"alarm_week_day": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"alarm_notify": {
							Type:     schema.TypeString,
							Optional: true,
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"email": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"mobile_phone": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"lang": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"notify_config": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notify_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"notify_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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
	notifyConfigRaw := objectRaw["NotifyConfig"]
	notifyConfigMaps := make([]map[string]interface{}, 0)
	if notifyConfigRaw != nil {
		for _, notifyConfigChildRaw := range convertToInterfaceArray(notifyConfigRaw) {
			notifyConfigMap := make(map[string]interface{})
			notifyConfigChildRaw := notifyConfigChildRaw.(map[string]interface{})
			notifyConfigMap["notify_type"] = notifyConfigChildRaw["NotifyType"]
			notifyConfigMap["notify_value"] = notifyConfigChildRaw["NotifyValue"]

			notifyConfigMaps = append(notifyConfigMaps, notifyConfigMap)
		}
	}
	if err := d.Set("notify_config", notifyConfigMaps); err != nil {
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
	alarmConfigDataList := make(map[string]interface{})

	if d.HasChange("alarm_config") {
		update = true
	}
	alarmValue1, _ := jsonpath.Get("$.alarm_value", d.Get("alarm_config"))
	if alarmValue1 != nil && alarmValue1 != "" {
		alarmConfigDataList["AlarmValue"] = alarmValue1
	}

	if d.HasChange("alarm_config") {
		update = true
	}
	alarmPeriod1, _ := jsonpath.Get("$.alarm_period", d.Get("alarm_config"))
	if alarmPeriod1 != nil && alarmPeriod1 != "" {
		alarmConfigDataList["AlarmPeriod"] = alarmPeriod1
	}

	if d.HasChange("alarm_config") {
		update = true
	}
	alarmWeekDay1, _ := jsonpath.Get("$.alarm_week_day", d.Get("alarm_config"))
	if alarmWeekDay1 != nil && alarmWeekDay1 != "" {
		alarmConfigDataList["AlarmWeekDay"] = alarmWeekDay1
	}

	if d.HasChange("alarm_config") {
		update = true
	}
	alarmNotify1, _ := jsonpath.Get("$.alarm_notify", d.Get("alarm_config"))
	if alarmNotify1 != nil && alarmNotify1 != "" {
		alarmConfigDataList["AlarmNotify"] = alarmNotify1
	}

	if d.HasChange("alarm_config") {
		update = true
	}
	alarmType1, _ := jsonpath.Get("$.alarm_type", d.Get("alarm_config"))
	if alarmType1 != nil && alarmType1 != "" {
		alarmConfigDataList["AlarmType"] = alarmType1
	}

	if d.HasChange("alarm_config") {
		update = true
	}
	alarmHour1, _ := jsonpath.Get("$.alarm_hour", d.Get("alarm_config"))
	if alarmHour1 != nil && alarmHour1 != "" {
		alarmConfigDataList["AlarmHour"] = alarmHour1
	}

	AlarmConfigMap := make([]interface{}, 0)
	AlarmConfigMap = append(AlarmConfigMap, alarmConfigDataList)
	request["AlarmConfig"] = AlarmConfigMap

	contactConfigDataList := make(map[string]interface{})

	if d.HasChange("contact_config") {
		update = true
		name1, _ := jsonpath.Get("$.name", d.Get("contact_config"))
		if name1 != nil && name1 != "" {
			contactConfigDataList["Name"] = name1
		}
	}

	if d.HasChange("contact_config") {
		update = true
		mobilePhone1, _ := jsonpath.Get("$.mobile_phone", d.Get("contact_config"))
		if mobilePhone1 != nil && mobilePhone1 != "" {
			contactConfigDataList["MobilePhone"] = mobilePhone1
		}
	}

	if d.HasChange("contact_config") {
		update = true
		email1, _ := jsonpath.Get("$.email", d.Get("contact_config"))
		if email1 != nil && email1 != "" {
			contactConfigDataList["Email"] = email1
		}
	}

	if d.HasChange("contact_config") {
		update = true
		status1, _ := jsonpath.Get("$.status", d.Get("contact_config"))
		if status1 != nil && status1 != "" {
			contactConfigDataList["Status"] = status1
		}
	}

	ContactConfigMap := make([]interface{}, 0)
	ContactConfigMap = append(ContactConfigMap, contactConfigDataList)
	request["ContactConfig"] = ContactConfigMap

	if d.HasChange("alarm_lang") {
		update = true
		request["AlarmLang"] = d.Get("alarm_lang")
	}

	notifyConfigDataList := make(map[string]interface{})

	if d.HasChange("notify_config") {
		update = true
		notifyType1, _ := jsonpath.Get("$.notify_type", d.Get("notify_config"))
		if notifyType1 != nil && notifyType1 != "" {
			notifyConfigDataList["NotifyType"] = notifyType1
		}
	}

	if d.HasChange("notify_config") {
		update = true
		notifyValue1, _ := jsonpath.Get("$.notify_value", d.Get("notify_config"))
		if notifyValue1 != nil && notifyValue1 != "" {
			notifyConfigDataList["NotifyValue"] = notifyValue1
		}
	}

	NotifyConfigMap := make([]interface{}, 0)
	NotifyConfigMap = append(NotifyConfigMap, notifyConfigDataList)
	request["NotifyConfig"] = NotifyConfigMap

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
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
