// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudDmsEnterpriseDifyInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudDmsEnterpriseDifyInstanceCreate,
		Read:   resourceAliCloudDmsEnterpriseDifyInstanceRead,
		Update: resourceAliCloudDmsEnterpriseDifyInstanceUpdate,
		Delete: resourceAliCloudDmsEnterpriseDifyInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"dify_instance_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"payment_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"PREPAY", "POSTPAY"}, false),
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchema(),
			"workspace_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudDmsEnterpriseDifyInstanceCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateDifyInstance"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["DataRegion"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	if v, ok := d.GetOk("payment_type"); ok {
		request["PayType"] = v
	}
	if v, ok := d.GetOk("dify_instance_name"); ok {
		request["DifyInstanceName"] = v
	}
	if v, ok := d.GetOk("workspace_id"); ok {
		request["WorkspaceId"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("dms-enterprise", "2018-11-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_dms_enterprise_dify_instance", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.Data.DifyInstanceId", response)
	d.SetId(fmt.Sprint(id))

	dmsEnterpriseServiceV2 := DmsEnterpriseServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"RUNNING"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, dmsEnterpriseServiceV2.DmsEnterpriseDifyInstanceStateRefreshFunc(d.Id(), "$.Root.Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudDmsEnterpriseDifyInstanceRead(d, meta)
}

func resourceAliCloudDmsEnterpriseDifyInstanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	dmsEnterpriseServiceV2 := DmsEnterpriseServiceV2{client}

	objectRaw, err := dmsEnterpriseServiceV2.DescribeDmsEnterpriseDifyInstance(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_dms_enterprise_dify_instance DescribeDmsEnterpriseDifyInstance Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	rootRawObj, _ := jsonpath.Get("$.Root", objectRaw)
	rootRaw := make(map[string]interface{})
	if rootRawObj != nil {
		rootRaw = rootRawObj.(map[string]interface{})
	}
	d.Set("dify_instance_name", rootRaw["DifyInstanceName"])
	d.Set("payment_type", rootRaw["ChargeType"])
	d.Set("status", rootRaw["Status"])
	d.Set("workspace_id", rootRaw["WorkspaceId"])

	tagsMaps := objectRaw["Tags"]
	d.Set("tags", tagsToMap(tagsMaps))

	return nil
}

func resourceAliCloudDmsEnterpriseDifyInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "UntagResources"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "DifyInstance"
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("dms-enterprise", "2018-11-01", action, query, request, true)
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
	update = false
	action = "UpdateDifyMeta"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["AppUuid"] = d.Id()
	request["DataRegion"] = client.RegionId
	if d.HasChange("dify_instance_name") {
		update = true
		request["Description"] = d.Get("dify_instance_name")
	}

	if d.HasChange("workspace_id") {
		update = true
		request["WorkspaceId"] = d.Get("workspace_id")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("dms-enterprise", "2018-11-01", action, query, request, true)
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

	d.Partial(false)
	return resourceAliCloudDmsEnterpriseDifyInstanceRead(d, meta)
}

func resourceAliCloudDmsEnterpriseDifyInstanceDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "StopDifyInstance"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["DifyInstanceId"] = d.Id()
	request["DataRegion"] = client.RegionId

	if v, ok := d.GetOk("workspace_id"); ok {
		request["WorkspaceId"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("dms-enterprise", "2018-11-01", action, query, request, true)
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
		if IsExpectedErrors(err, []string{"DMS.MLOPS.RESOURCE_DOES_NOT_EXIST"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
