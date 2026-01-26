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

func resourceAliCloudComputeNestRestoreTask() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudComputeNestRestoreTaskCreate,
		Read:   resourceAliCloudComputeNestRestoreTaskRead,
		Delete: resourceAliCloudComputeNestRestoreTaskDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"backup_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudComputeNestRestoreTaskCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateRestoreTask"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["ServiceInstanceId"] = d.Get("service_instance_id")
	request["BackupId"] = d.Get("backup_id")
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_compute_nest_restore_task", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["RestoreTaskId"]))

	return resourceAliCloudComputeNestRestoreTaskRead(d, meta)
}

func resourceAliCloudComputeNestRestoreTaskRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	computeNestServiceV2 := ComputeNestServiceV2{client}

	objectRaw, err := computeNestServiceV2.DescribeComputeNestRestoreTask(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_compute_nest_restore_task DescribeComputeNestRestoreTask Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("backup_id", objectRaw["BackupId"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("service_instance_id", objectRaw["ServiceInstanceId"])
	d.Set("status", objectRaw["Status"])

	return nil
}

func resourceAliCloudComputeNestRestoreTaskDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Restore Task. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
