// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tidwall/sjson"
)

func resourceAliCloudCmsPrometheusCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCmsPrometheusClusterCreate,
		Read:   resourceAliCloudCmsPrometheusClusterRead,
		Update: resourceAliCloudCmsPrometheusClusterUpdate,
		Delete: resourceAliCloudCmsPrometheusClusterDelete,
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
			"hash_label_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"payment_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"prometheus_cluster_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replicas": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: IntBetween(2, 48),
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_duration": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tags": tagsSchema(),
		},
	}
}

func resourceAliCloudCmsPrometheusClusterCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/prometheus-clusters")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("prometheus_cluster_name"); ok {
		request["prometheusClusterName"] = v
	}

	request["replicas"] = d.Get("replicas")
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request["Tags"] = tagsMap
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["resourceGroupId"] = v
	}
	if v, ok := d.GetOk("payment_type"); ok {
		request["paymentType"] = v
	}
	request["hashLabelKey"] = d.Get("hash_label_key")
	if v, ok := d.GetOkExists("storage_duration"); ok {
		request["storageDuration"] = v
	}
	body = request
	wait := incrementalWait(5*time.Second, 5*time.Second)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cms_prometheus_cluster", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["prometheusClusterName"]))

	return resourceAliCloudCmsPrometheusClusterRead(d, meta)
}

func resourceAliCloudCmsPrometheusClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cmsServiceV2 := CmsServiceV2{client}

	objectRaw, err := cmsServiceV2.DescribeCmsPrometheusCluster(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cms_prometheus_cluster DescribeCmsPrometheusCluster Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["createTime"])
	d.Set("hash_label_key", objectRaw["hashLabelKey"])
	d.Set("payment_type", objectRaw["paymentType"])
	d.Set("region_id", objectRaw["regionId"])
	d.Set("replicas", objectRaw["replicas"])
	d.Set("resource_group_id", objectRaw["resourceGroupId"])
	d.Set("storage_duration", objectRaw["storageDuration"])
	d.Set("prometheus_cluster_name", objectRaw["prometheusClusterName"])

	tagsMaps := objectRaw["tags"]
	d.Set("tags", tagsToMap(tagsMaps))

	d.Set("prometheus_cluster_name", d.Id())

	return nil
}

func resourceAliCloudCmsPrometheusClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	prometheusClusterName := d.Id()
	action := fmt.Sprintf("/prometheus-clusters/%s", prometheusClusterName)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})

	if d.HasChange("storage_duration") {
		update = true
	}
	if v, ok := d.GetOkExists("storage_duration"); ok || d.HasChange("storage_duration") {
		request["storageDuration"] = v
	}
	body = request
	if update {
		wait := incrementalWait(5*time.Second, 5*time.Second)
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
	update = false
	action = fmt.Sprintf("/resourcegroup")
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	request["resourceId"] = d.Id()
	query["RegionId"] = StringPointer(client.RegionId)
	if _, ok := d.GetOk("resource_group_id"); ok && d.HasChange("resource_group_id") {
		update = true
	}
	if v, ok := d.GetOk("resource_group_id"); ok || d.HasChange("resource_group_id") {
		request["resourceGroupId"] = v
	}
	if v, ok := d.GetOk("resource_type"); ok {
		request["resourceType"] = v
	}
	body = request
	if update {
		wait := incrementalWait(5*time.Second, 5*time.Second)
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
	update = false
	action = fmt.Sprintf("/tags")
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	query["RegionId"] = StringPointer(client.RegionId)
	tagDataList := make(map[string]interface{})

	if d.HasChange("tags") {
		update = true
	}
	tagsMap := ConvertTags(v.(map[string]interface{}))
	tagDataList["key"] = tagsMap

	if d.HasChange("tags") {
		update = true
	}
	tagsMap := ConvertTags(v.(map[string]interface{}))
	tagDataList["value"] = tagsMap

	tagMap := make([]interface{}, 0)
	tagMap = append(tagMap, tagDataList)
	request["tag"] = tagMap

	request["resourceType"] = d.Get("resource_type")
	jsonString := convertObjectToJsonString(request)
	jsonString, _ = sjson.Set(jsonString, "resourceId.0", d.Id())
	_ = json.Unmarshal([]byte(jsonString), &request)

	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPost("Cms", "2024-03-30", action, query, nil, body, true)
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
	action = fmt.Sprintf("/tags")
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	query["resourceId"] = StringPointer(convertListToJsonString(expandSingletonToList(d.Id())))
	query["RegionId"] = StringPointer(client.RegionId)
	if v, ok := d.GetOk("resource_type"); ok {
		query["resourceType"] = StringPointer(v.(string))
	}

	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaDelete("Cms", "2024-03-30", action, query, nil, nil, true)
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

	return resourceAliCloudCmsPrometheusClusterRead(d, meta)
}

func resourceAliCloudCmsPrometheusClusterDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	prometheusClusterName := d.Id()
	action := fmt.Sprintf("/prometheus-clusters/%s", prometheusClusterName)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	var err error
	request = make(map[string]interface{})

	wait := incrementalWait(5*time.Second, 5*time.Second)
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
