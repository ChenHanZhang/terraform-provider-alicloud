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

func resourceAliCloudPolardbBackupLog() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPolardbBackupLogCreate,
		Read:   resourceAliCloudPolardbBackupLogRead,
		Update: resourceAliCloudPolardbBackupLogUpdate,
		Delete: resourceAliCloudPolardbBackupLogDelete,
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"log_backup_another_region_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_backup_another_region_retention_period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_backup_retention_period": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudPolardbBackupLogCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "ModifyLogBackupPolicy"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("db_cluster_id"); ok {
		request["DBClusterId"] = v
	}

	if v, ok := d.GetOkExists("log_backup_retention_period"); ok {
		request["LogBackupRetentionPeriod"] = v
	}
	if v, ok := d.GetOk("log_backup_another_region_region"); ok {
		request["LogBackupAnotherRegionRegion"] = v
	}
	if v, ok := d.GetOk("log_backup_another_region_retention_period"); ok {
		request["LogBackupAnotherRegionRetentionPeriod"] = v
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_polardb_backup_log", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["DBClusterId"]))

	return resourceAliCloudPolardbBackupLogUpdate(d, meta)
}

func resourceAliCloudPolardbBackupLogRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	polardbServiceV2 := PolardbServiceV2{client}

	objectRaw, err := polardbServiceV2.DescribePolardbBackupLog(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_polardb_backup_log DescribePolardbBackupLog Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("log_backup_another_region_region", objectRaw["LogBackupAnotherRegionRegion"])
	d.Set("log_backup_another_region_retention_period", objectRaw["LogBackupAnotherRegionRetentionPeriod"])
	d.Set("log_backup_retention_period", objectRaw["LogBackupRetentionPeriod"])

	d.Set("db_cluster_id", d.Id())

	return nil
}

func resourceAliCloudPolardbBackupLogUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "ModifyLogBackupPolicy"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if !d.IsNewResource() && d.HasChange("log_backup_retention_period") {
		update = true
		request["LogBackupRetentionPeriod"] = d.Get("log_backup_retention_period")
	}

	if !d.IsNewResource() && d.HasChange("log_backup_another_region_region") {
		update = true
		request["LogBackupAnotherRegionRegion"] = d.Get("log_backup_another_region_region")
	}

	if !d.IsNewResource() && d.HasChange("log_backup_another_region_retention_period") {
		update = true
		request["LogBackupAnotherRegionRetentionPeriod"] = d.Get("log_backup_another_region_retention_period")
	}

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

	return resourceAliCloudPolardbBackupLogRead(d, meta)
}

func resourceAliCloudPolardbBackupLogDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Backup Log. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
