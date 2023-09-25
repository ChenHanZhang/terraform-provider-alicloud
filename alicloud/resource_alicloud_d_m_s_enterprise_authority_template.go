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

func resourceAliCloudDMSEnterpriseAuthorityTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudDMSEnterpriseAuthorityTemplateCreate,
		Read:   resourceAliCloudDMSEnterpriseAuthorityTemplateRead,
		Update: resourceAliCloudDMSEnterpriseAuthorityTemplateUpdate,
		Delete: resourceAliCloudDMSEnterpriseAuthorityTemplateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"authority_template_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tid": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudDMSEnterpriseAuthorityTemplateCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateAuthorityTemplate"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewDmsenterpriseClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["Name"] = d.Get("authority_template_name")
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("tid"); ok {
		request["Tid"] = v
	}
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2018-11-01"), StringPointer("AK"), nil, request, &runtime)

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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_d_m_s_enterprise_authority_template", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.AuthorityTemplateView.TemplateId", response)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudDMSEnterpriseAuthorityTemplateRead(d, meta)
}

func resourceAliCloudDMSEnterpriseAuthorityTemplateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	dMSEnterpriseServiceV2 := DMSEnterpriseServiceV2{client}

	objectRaw, err := dMSEnterpriseServiceV2.DescribeDMSEnterpriseAuthorityTemplate(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_d_m_s_enterprise_authority_template DescribeDMSEnterpriseAuthorityTemplate Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("authority_template_name", objectRaw["Name"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("description", objectRaw["Description"])

	return nil
}

func resourceAliCloudDMSEnterpriseAuthorityTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	action := "UpdateAuthorityTemplate"
	conn, err := client.NewDmsenterpriseClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["TemplateId"] = d.Id()
	if d.HasChange("authority_template_name") {
		update = true
	}
	request["Name"] = d.Get("authority_template_name")
	if d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	if v, ok := d.GetOk("tid"); ok {
		request["Tid"] = v
	}
	if update {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2018-11-01"), StringPointer("AK"), nil, request, &runtime)

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

	return resourceAliCloudDMSEnterpriseAuthorityTemplateRead(d, meta)
}

func resourceAliCloudDMSEnterpriseAuthorityTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Authority Template. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
