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

func resourceAliCloudAdbConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAdbConnectionCreate,
		Read:   resourceAliCloudAdbConnectionRead,
		Update: resourceAliCloudAdbConnectionUpdate,
		Delete: resourceAliCloudAdbConnectionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"connection_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_string_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"db_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudAdbConnectionCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "AllocateClusterPublicConnection"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("db_cluster_id"); ok {
		request["DBClusterId"] = v
	}

	if v, ok := d.GetOk("connection_string_prefix"); ok {
		request["ConnectionStringPrefix"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_adb_connection", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["DBClusterId"]))

	return resourceAliCloudAdbConnectionUpdate(d, meta)
}

func resourceAliCloudAdbConnectionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	adbServiceV2 := AdbServiceV2{client}

	objectRaw, err := adbServiceV2.DescribeAdbConnection(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_adb_connection DescribeAdbConnection Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("connection_string", objectRaw["ConnectionString"])
	d.Set("connection_string_prefix", objectRaw["ConnectionStringPrefix"])
	d.Set("ip_address", objectRaw["IPAddress"])
	d.Set("port", formatInt(objectRaw["Port"]))

	d.Set("db_cluster_id", d.Id())

	return nil
}

func resourceAliCloudAdbConnectionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "ModifyClusterConnectionString"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if !d.IsNewResource() && d.HasChange("connection_string_prefix") {
		update = true
	}
	request["ConnectionStringPrefix"] = d.Get("connection_string_prefix")
	if d.HasChange("port") {
		update = true
		request["Port"] = d.Get("port")
	}

	if d.HasChange("connection_string") {
		update = true
	}
	request["CurrentConnectionString"] = d.Get("connection_string")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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

	return resourceAliCloudAdbConnectionRead(d, meta)
}

func resourceAliCloudAdbConnectionDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "ReleaseClusterPublicConnection"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
		if IsExpectedErrors(err, []string{"InvalidDBCluster.NotFound"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
