// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudPAIWorkspaceModel() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPAIWorkspaceModelCreate,
		Read:   resourceAliCloudPAIWorkspaceModelRead,
		Update: resourceAliCloudPAIWorkspaceModelUpdate,
		Delete: resourceAliCloudPAIWorkspaceModelDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"accessibility": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"PRIVATE", "PUBLIC"}, false),
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"extra_info": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"labels": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"model_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model_doc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"order_number": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"origin": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"task": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"workspace_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudPAIWorkspaceModelCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/api/v1/models")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	conn, err := client.NewPaiworkspaceClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["ModelName"] = d.Get("model_name")
	if v, ok := d.GetOk("labels"); ok {
		labelsMaps := make([]map[string]interface{}, 0)
		for _, dataLoop := range v.([]interface{}) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["Key"] = dataLoopTmp["key"]
			dataLoopMap["Value"] = dataLoopTmp["value"]
			labelsMaps = append(labelsMaps, dataLoopMap)
		}
		request["Labels"] = labelsMaps
	}

	if v, ok := d.GetOk("model_description"); ok {
		request["ModelDescription"] = v
	}
	if v, ok := d.GetOk("workspace_id"); ok {
		request["WorkspaceId"] = v
	}
	if v, ok := d.GetOk("accessibility"); ok {
		request["Accessibility"] = v
	}
	if v, ok := d.GetOk("origin"); ok {
		request["Origin"] = v
	}
	if v, ok := d.GetOk("domain"); ok {
		request["Domain"] = v
	}
	if v, ok := d.GetOk("task"); ok {
		request["Task"] = v
	}
	if v, ok := d.GetOk("model_doc"); ok {
		request["ModelDoc"] = v
	}
	if v, ok := d.GetOk("order_number"); ok {
		request["OrderNumber"] = v
	}
	if v, ok := d.GetOk("model_type"); ok {
		request["ModelType"] = v
	}
	if v, ok := d.GetOk("extra_info"); ok {
		request["ExtraInfo"] = convertMapToJsonStringIgnoreError(v.(map[string]interface{}))
	}
	body = request
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer("2021-02-04"), nil, StringPointer("POST"), StringPointer("AK"), StringPointer(action), query, nil, body, &runtime)

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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_p_a_i_workspace_model", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["ModelId"]))

	return resourceAliCloudPAIWorkspaceModelRead(d, meta)
}

func resourceAliCloudPAIWorkspaceModelRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	pAIWorkspaceServiceV2 := PAIWorkspaceServiceV2{client}

	objectRaw, err := pAIWorkspaceServiceV2.DescribePAIWorkspaceModel(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_p_a_i_workspace_model DescribePAIWorkspaceModel Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("accessibility", objectRaw["Accessibility"])
	d.Set("domain", objectRaw["Domain"])
	d.Set("extra_info", objectRaw["ExtraInfo"])
	d.Set("model_description", objectRaw["ModelDescription"])
	d.Set("model_doc", objectRaw["ModelDoc"])
	d.Set("model_name", objectRaw["ModelName"])
	d.Set("model_type", objectRaw["ModelType"])
	d.Set("order_number", objectRaw["OrderNumber"])
	d.Set("origin", objectRaw["Origin"])
	d.Set("task", objectRaw["Task"])
	d.Set("workspace_id", objectRaw["WorkspaceId"])
	labels1Raw := objectRaw["Labels"]
	labelsMaps := make([]map[string]interface{}, 0)
	if labels1Raw != nil {
		for _, labelsChild1Raw := range labels1Raw.([]interface{}) {
			labelsMap := make(map[string]interface{})
			labelsChild1Raw := labelsChild1Raw.(map[string]interface{})
			labelsMap["key"] = labelsChild1Raw["Key"]
			labelsMap["value"] = labelsChild1Raw["Value"]
			labelsMaps = append(labelsMaps, labelsMap)
		}
	}
	d.Set("labels", labelsMaps)

	return nil
}

func resourceAliCloudPAIWorkspaceModelUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	update := false
	ModelId := d.Id()
	action := fmt.Sprintf("/api/v1/models/%s", ModelId)
	conn, err := client.NewPaiworkspaceClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	request["ModelId"] = d.Id()
	if d.HasChange("model_name") {
		update = true
	}
	request["ModelName"] = d.Get("model_name")
	if d.HasChange("model_description") {
		update = true
		request["ModelDescription"] = d.Get("model_description")
	}

	if d.HasChange("accessibility") {
		update = true
		request["Accessibility"] = d.Get("accessibility")
	}

	if d.HasChange("origin") {
		update = true
		request["Origin"] = d.Get("origin")
	}

	if d.HasChange("domain") {
		update = true
		request["Domain"] = d.Get("domain")
	}

	if d.HasChange("task") {
		update = true
		request["Task"] = d.Get("task")
	}

	if d.HasChange("model_doc") {
		update = true
		request["ModelDoc"] = d.Get("model_doc")
	}

	if d.HasChange("order_number") {
		update = true
		request["OrderNumber"] = d.Get("order_number")
	}

	if d.HasChange("model_type") {
		update = true
		request["ModelType"] = d.Get("model_type")
	}

	if d.HasChange("extra_info") {
		update = true
		request["ExtraInfo"] = convertMapToJsonStringIgnoreError(d.Get("extra_info").(map[string]interface{}))
	}

	body = request
	if update {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer("2021-02-04"), nil, StringPointer("PUT"), StringPointer("AK"), StringPointer(action), query, nil, body, &runtime)

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

	return resourceAliCloudPAIWorkspaceModelRead(d, meta)
}

func resourceAliCloudPAIWorkspaceModelDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	ModelId := d.Id()
	action := fmt.Sprintf("/api/v1/models/%s", ModelId)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	conn, err := client.NewPaiworkspaceClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["ModelId"] = d.Id()

	body = request
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer("2021-02-04"), nil, StringPointer("DELETE"), StringPointer("AK"), StringPointer(action), query, nil, body, &runtime)

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

	return nil
}
