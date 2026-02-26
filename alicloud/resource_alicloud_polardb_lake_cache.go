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

func resourceAliCloudPolardbLakeCache() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPolardbLakeCacheCreate,
		Read:   resourceAliCloudPolardbLakeCacheRead,
		Update: resourceAliCloudPolardbLakeCacheUpdate,
		Delete: resourceAliCloudPolardbLakeCacheDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"accelerated_storage_space": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"accelerating_enable": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"ON", "OFF", "ONLY"}, false),
			},
			"authorized_user_ids": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_renew": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"auto_use_coupon": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_bucket_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"custom_bucket_path": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"custom_oss_ak": {
				Type:      schema.TypeString,
				Optional:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"custom_oss_sk": {
				Type:      schema.TypeString,
				Optional:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"db_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"polardb_mysql", "polardb_pg"}, false),
			},
			"db_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"payment_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"PayAsYouGo", "Subscription"}, false),
			},
			"period": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"Year", "Month"}, false),
			},
			"promotion_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_space": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"storage_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tags": tagsSchema(),
			"used_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vswitch_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudPolardbLakeCacheCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreatePolarFs"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("custom_oss_sk"); ok {
		request["CustomOssSk"] = v
	}
	request["PayType"] = d.Get("payment_type")
	if v, ok := d.GetOk("storage_type"); ok {
		request["StorageType"] = v
	}
	if v, ok := d.GetOk("used_time"); ok {
		request["UsedTime"] = v
	}
	if v, ok := d.GetOk("vswitch_id"); ok {
		request["VSwitchId"] = v
	}
	request["DBClusterId"] = d.Get("db_cluster_id")
	if v, ok := d.GetOk("custom_oss_ak"); ok {
		request["CustomOssAk"] = v
	}
	if v, ok := d.GetOk("category"); ok {
		request["CreationCategory"] = v
	}
	if v, ok := d.GetOk("accelerated_storage_space"); ok {
		request["AccelerateStorageSize"] = v
	}
	if v, ok := d.GetOkExists("auto_renew"); ok {
		request["AutoRenew"] = v
	}
	if v, ok := d.GetOk("custom_bucket_path"); ok {
		request["CustomBucketPath"] = v
	}
	if v, ok := d.GetOk("promotion_code"); ok {
		request["PromotionCode"] = v
	}
	if v, ok := d.GetOkExists("storage_space"); ok {
		request["StorageSpace"] = v
	}
	if v, ok := d.GetOkExists("auto_use_coupon"); ok {
		request["AutoUseCoupon"] = v
	}
	if v, ok := d.GetOk("authorized_user_ids"); ok {
		request["AuthorizedUserIds"] = v
	}
	if v, ok := d.GetOk("accelerating_enable"); ok {
		request["AccelerateSwitch"] = v
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		request["VPCId"] = v
	}
	if v, ok := d.GetOk("period"); ok {
		request["Period"] = v
	}
	if v, ok := d.GetOkExists("custom_bucket_count"); ok {
		request["CustomBucketCount"] = v
	}
	if v, ok := d.GetOk("zone_id"); ok {
		request["ZoneId"] = v
	}
	if v, ok := d.GetOk("db_type"); ok {
		request["DBType"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_polardb_lake_cache", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["PolarFsInstanceId"]))

	return resourceAliCloudPolardbLakeCacheUpdate(d, meta)
}

func resourceAliCloudPolardbLakeCacheRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	polardbServiceV2 := PolardbServiceV2{client}

	objectRaw, err := polardbServiceV2.DescribePolardbLakeCache(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_polardb_lake_cache DescribePolardbLakeCache Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("accelerated_storage_space", objectRaw["AcceleratedStorageSpace"])
	d.Set("accelerating_enable", objectRaw["AcceleratingEnable"])
	d.Set("category", objectRaw["Category"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("custom_bucket_path", objectRaw["CustomBucketPath"])
	d.Set("db_type", objectRaw["DBType"])
	d.Set("db_cluster_id", objectRaw["RelativeDbClusterId"])
	d.Set("description", objectRaw["PolarFsInstanceDescription"])
	d.Set("payment_type", objectRaw["PayType"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("status", objectRaw["PolarFsStatus"])
	d.Set("storage_space", objectRaw["StorageSpace"])
	d.Set("storage_type", objectRaw["StorageType"])
	d.Set("vpc_id", objectRaw["VPCId"])
	d.Set("vswitch_id", objectRaw["VSwitchId"])
	d.Set("zone_id", objectRaw["ZoneId"])

	objectRaw, err = polardbServiceV2.DescribeLakeCacheListTagResources(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	tagsMaps := tagResourceChildRaw.(map[string]interface{})
	d.Set("tags", tagsToMap(tagsMaps))

	return nil
}

func resourceAliCloudPolardbLakeCacheUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "ModifyPolarFsInstanceDescription"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["PolarFsInstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("description") {
		update = true
	}
	request["PolarFsInstanceDescription"] = d.Get("description")
	if !d.IsNewResource() && d.HasChange("db_cluster_id") {
		update = true
	}
	request["DBClusterId"] = d.Get("db_cluster_id")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
	action = "TagResources"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ResourceId.1"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("tags") {
		update = true
	}
	if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	request["ResourceType"] = "lakecache"
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
	action = "UntagResources"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ResourceId.1"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("tags") {
		update = true
		if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
			tagsMap := ConvertTags(v.(map[string]interface{}))
			request = expandTagsToMap(request, tagsMap)
		}
	}

	request["ResourceType"] = "lakecache"
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
	return resourceAliCloudPolardbLakeCacheRead(d, meta)
}

func resourceAliCloudPolardbLakeCacheDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeletePolarFs"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["PolarFsInstanceId"] = d.Id()

	request["DBClusterId"] = d.Get("db_cluster_id")
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
