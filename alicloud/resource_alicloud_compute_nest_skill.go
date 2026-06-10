// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudComputeNestSkill() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudComputeNestSkillCreate,
		Read:   resourceAliCloudComputeNestSkillRead,
		Update: resourceAliCloudComputeNestSkillUpdate,
		Delete: resourceAliCloudComputeNestSkillDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oss_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"skill_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"skill_labels": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"skill_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"skill_space_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_skill_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudComputeNestSkillCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateSkill"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("source_skill_id"); ok {
		request["SourceSkillId"] = v
	}
	request["SourceType"] = "UPLOAD"
	if v, ok := d.GetOk("skill_name"); ok {
		request["SkillName"] = v
	}
	request["SkillSpaceId"] = d.Get("skill_space_id")
	if v, ok := d.GetOk("oss_url"); ok {
		request["OssUrl"] = v
	}
	if v, ok := d.GetOk("skill_description"); ok {
		request["SkillDescription"] = v
	}
	if v, ok := d.GetOk("skill_labels"); ok {
		skillLabelsMapsArray := convertToInterfaceArray(v)

		request["SkillLabels"] = skillLabelsMapsArray
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("ComputeNest", "2021-06-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_compute_nest_skill", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["SkillId"]))

	return resourceAliCloudComputeNestSkillRead(d, meta)
}

func resourceAliCloudComputeNestSkillRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	computeNestServiceV2 := ComputeNestServiceV2{client}

	objectRaw, err := computeNestServiceV2.DescribeComputeNestSkill(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_compute_nest_skill DescribeComputeNestSkill Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("skill_description", objectRaw["SkillDescription"])
	d.Set("skill_name", objectRaw["SkillName"])
	d.Set("skill_space_id", objectRaw["SkillSpaceId"])
	d.Set("update_time", objectRaw["UpdateTime"])

	skillLabelsRaw := make([]interface{}, 0)
	if objectRaw["SkillLabels"] != nil {
		skillLabelsRaw = convertToInterfaceArray(objectRaw["SkillLabels"])
	}

	d.Set("skill_labels", skillLabelsRaw)

	return nil
}

func resourceAliCloudComputeNestSkillUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "UpdateSkill"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["SkillId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if v, ok := d.GetOk("source_skill_id"); ok {
		request["SourceSkillId"] = v
	}
	request["SourceType"] = "UPLOAD"
	if v, ok := d.GetOk("oss_url"); ok {
		request["OssUrl"] = v
	}
	if d.HasChange("skill_description") {
		update = true
		request["SkillDescription"] = d.Get("skill_description")
	}

	if d.HasChange("skill_labels") {
		update = true
		if v, ok := d.GetOk("skill_labels"); ok || d.HasChange("skill_labels") {
			skillLabelsMapsArray := convertToInterfaceArray(v)

			request["SkillLabels"] = skillLabelsMapsArray
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("ComputeNest", "2021-06-01", action, query, request, true)
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

	return resourceAliCloudComputeNestSkillRead(d, meta)
}

func resourceAliCloudComputeNestSkillDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteSkill"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["SkillId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("ComputeNest", "2021-06-01", action, query, request, true)
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
		if IsExpectedErrors(err, []string{"InvalidSkill.NotFound"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
