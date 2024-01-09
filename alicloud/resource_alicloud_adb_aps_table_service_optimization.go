// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"regexp"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudAdbApsTableServiceOptimization() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAdbApsTableServiceOptimizationCreate,
		Read:   resourceAliCloudAdbApsTableServiceOptimizationRead,
		Update: resourceAliCloudAdbApsTableServiceOptimizationUpdate,
		Delete: resourceAliCloudAdbApsTableServiceOptimizationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"db_cluster_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringMatch(regexp.MustCompile("amv-[A-Za-z0-9]{16}"), "The ID of the ADB instance to which the resource belongs."),
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"on", "off"}, true),
			},
		},
	}
}

func resourceAliCloudAdbApsTableServiceOptimizationCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "SwitchApsOptimizationStrategy"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	conn, err := client.NewAdbClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["DBClusterId"] = d.Get("db_cluster_id")
	request["RegionId"] = client.RegionId

	request["StrategyType"] = "CLEAN"
	if v, ok := d.GetOk("status"); ok {
		request["Status"] = v
	}
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2021-12-01"), StringPointer("AK"), query, request, &runtime)

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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_adb_aps_table_service_optimization", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["DBClusterId"]))

	return resourceAliCloudAdbApsTableServiceOptimizationRead(d, meta)
}

func resourceAliCloudAdbApsTableServiceOptimizationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	adbServiceV2 := AdbServiceV2{client}

	objectRaw, err := adbServiceV2.DescribeAdbApsTableServiceOptimization(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_adb_aps_table_service_optimization DescribeAdbApsTableServiceOptimization Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("status", objectRaw["Status"])
	d.Set("db_cluster_id", objectRaw["DBClusterId"])

	d.Set("db_cluster_id", d.Id())

	return nil
}

func resourceAliCloudAdbApsTableServiceOptimizationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	action := "SwitchApsOptimizationStrategy"
	conn, err := client.NewAdbClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("status") {
		update = true
		request["Status"] = d.Get("status")
	}

	request["StrategyType"] = "CLEAN"
	if update {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2021-12-01"), StringPointer("AK"), query, request, &runtime)

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

	return resourceAliCloudAdbApsTableServiceOptimizationRead(d, meta)
}

func resourceAliCloudAdbApsTableServiceOptimizationDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Aps Table Service Optimization. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
