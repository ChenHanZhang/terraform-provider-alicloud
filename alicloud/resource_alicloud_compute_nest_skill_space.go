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

func resourceAliCloudComputeNestSkillSpace() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudComputeNestSkillSpaceCreate,
		Read:   resourceAliCloudComputeNestSkillSpaceRead,
		Delete: resourceAliCloudComputeNestSkillSpaceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"skill_space_description": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"skill_space_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudComputeNestSkillSpaceCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateSkillSpace"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["ClientToken"] = buildClientToken(action)

	request["SkillSpaceDescription"] = d.Get("skill_space_description")
	request["SkillSpaceName"] = d.Get("skill_space_name")
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_compute_nest_skill_space", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["SkillSpaceId"]))

	return resourceAliCloudComputeNestSkillSpaceRead(d, meta)
}

func resourceAliCloudComputeNestSkillSpaceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	computeNestServiceV2 := ComputeNestServiceV2{client}

	objectRaw, err := computeNestServiceV2.DescribeComputeNestSkillSpace(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_compute_nest_skill_space DescribeComputeNestSkillSpace Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("skill_space_description", objectRaw["SkillSpaceDescription"])
	d.Set("skill_space_name", objectRaw["SkillSpaceName"])
	d.Set("update_time", objectRaw["UpdateTime"])

	return nil
}

func resourceAliCloudComputeNestSkillSpaceDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteSkillSpace"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["SkillSpaceId"] = d.Id()

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
		if IsExpectedErrors(err, []string{"InvalidSkillSpace.NotFound"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
