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

func resourceAliCloudCmsDeliveryTask() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCmsDeliveryTaskCreate,
		Read:   resourceAliCloudCmsDeliveryTaskRead,
		Update: resourceAliCloudCmsDeliveryTaskUpdate,
		Delete: resourceAliCloudCmsDeliveryTaskDelete,
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
			"data_source_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"label_filters": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"label_filters_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sink_list": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sink_configs": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"sink_type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"Running", "Disable"}, false),
			},
			"tags": tagsSchema(),
			"task_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"task_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAliCloudCmsDeliveryTaskCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/delivery-tasks")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request["Tags"] = tagsMap
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["resourceGroupId"] = v
	}
	if v, ok := d.GetOk("label_filters_type"); ok {
		request["labelFiltersType"] = v
	}
	if v, ok := d.GetOk("sink_list"); ok {
		sinkListMapsArray := make([]interface{}, 0)
		for _, dataLoop1 := range convertToInterfaceArray(v) {
			dataLoop1Tmp := dataLoop1.(map[string]interface{})
			dataLoop1Map := make(map[string]interface{})
			dataLoop1Map["sinkConfigs"] = dataLoop1Tmp["sink_configs"]
			dataLoop1Map["sinkType"] = dataLoop1Tmp["sink_type"]
			sinkListMapsArray = append(sinkListMapsArray, dataLoop1Map)
		}
		request["sinkList"] = sinkListMapsArray
	}

	if v, ok := d.GetOk("task_description"); ok {
		request["taskDescription"] = v
	}
	request["dataSourceId"] = d.Get("data_source_id")
	if v, ok := d.GetOk("external_labels"); ok {
		request["externalLabels"] = v
	}
	if v, ok := d.GetOk("label_filters"); ok {
		request["labelFilters"] = v
	}
	request["taskName"] = d.Get("task_name")
	body = request
	wait := incrementalWait(10*time.Second, 10*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RoaPost("Cms", "2024-03-30", action, query, nil, body, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"500"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cms_delivery_task", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["taskId"]))

	return resourceAliCloudCmsDeliveryTaskUpdate(d, meta)
}

func resourceAliCloudCmsDeliveryTaskRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cmsServiceV2 := CmsServiceV2{client}

	objectRaw, err := cmsServiceV2.DescribeCmsDeliveryTask(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cms_delivery_task DescribeCmsDeliveryTask Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["createTime"])
	d.Set("data_source_id", objectRaw["dataSourceId"])
	d.Set("external_labels", objectRaw["externalLabels"])
	d.Set("label_filters", objectRaw["labelFilters"])
	d.Set("label_filters_type", objectRaw["labelFiltersType"])
	d.Set("region_id", objectRaw["regionId"])
	d.Set("status", objectRaw["status"])
	d.Set("task_description", objectRaw["taskDescription"])
	d.Set("task_name", objectRaw["taskName"])

	sinkListRaw := objectRaw["sinkList"]
	sinkListMaps := make([]map[string]interface{}, 0)
	if sinkListRaw != nil {
		for _, sinkListChildRaw := range convertToInterfaceArray(sinkListRaw) {
			sinkListMap := make(map[string]interface{})
			sinkListChildRaw := sinkListChildRaw.(map[string]interface{})
			sinkListMap["sink_configs"] = sinkListChildRaw["sinkConfigs"]
			sinkListMap["sink_type"] = sinkListChildRaw["sinkType"]

			sinkListMaps = append(sinkListMaps, sinkListMap)
		}
	}
	if err := d.Set("sink_list", sinkListMaps); err != nil {
		return err
	}
	tagsMaps := objectRaw["tags"]
	d.Set("tags", tagsToMap(tagsMaps))

	return nil
}

func resourceAliCloudCmsDeliveryTaskUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	taskId := d.Id()
	action := fmt.Sprintf("/delivery-task/%s", taskId)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})

	if !d.IsNewResource() && d.HasChange("task_description") {
		update = true
	}
	if v, ok := d.GetOk("task_description"); ok || d.HasChange("task_description") {
		request["taskDescription"] = v
	}
	if !d.IsNewResource() && d.HasChange("data_source_id") {
		update = true
	}
	request["dataSourceId"] = d.Get("data_source_id")
	if !d.IsNewResource() && d.HasChange("label_filters_type") {
		update = true
	}
	if v, ok := d.GetOk("label_filters_type"); ok || d.HasChange("label_filters_type") {
		request["labelFiltersType"] = v
	}
	if !d.IsNewResource() && d.HasChange("external_labels") {
		update = true
	}
	if v, ok := d.GetOk("external_labels"); ok || d.HasChange("external_labels") {
		request["externalLabels"] = v
	}
	if !d.IsNewResource() && d.HasChange("label_filters") {
		update = true
	}
	if v, ok := d.GetOk("label_filters"); ok || d.HasChange("label_filters") {
		request["labelFilters"] = v
	}
	if d.HasChange("status") {
		update = true
	}
	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		request["status"] = v
	}
	if !d.IsNewResource() && d.HasChange("sink_list") {
		update = true
	}
	if v, ok := d.GetOk("sink_list"); ok || d.HasChange("sink_list") {
		sinkListMapsArray := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(v) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["sinkType"] = dataLoopTmp["sink_type"]
			dataLoopMap["sinkConfigs"] = dataLoopTmp["sink_configs"]
			sinkListMapsArray = append(sinkListMapsArray, dataLoopMap)
		}
		request["sinkList"] = sinkListMapsArray
	}

	if !d.IsNewResource() && d.HasChange("task_name") {
		update = true
	}
	request["taskName"] = d.Get("task_name")
	body = request
	if update {
		wait := incrementalWait(10*time.Second, 10*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPut("Cms", "2024-03-30", action, query, nil, body, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"500"}) || NeedRetry(err) {
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

	return resourceAliCloudCmsDeliveryTaskRead(d, meta)
}

func resourceAliCloudCmsDeliveryTaskDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	taskId := d.Id()
	action := fmt.Sprintf("/delivery-task/%s", taskId)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	var err error
	request = make(map[string]interface{})

	wait := incrementalWait(10*time.Second, 10*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RoaDelete("Cms", "2024-03-30", action, query, nil, nil, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"500"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		if IsExpectedErrors(err, []string{"404"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
