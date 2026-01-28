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

func resourceAliCloudRdsBackupFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudRdsBackupFileCreate,
		Read:   resourceAliCloudRdsBackupFileRead,
		Update: resourceAliCloudRdsBackupFileUpdate,
		Delete: resourceAliCloudRdsBackupFileDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"backup_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bucket_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"build_replication": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"engine_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"master_info": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restore_size": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"retention": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_info": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudRdsBackupFileCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "ImportUserBackupFile"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("comment"); ok {
		request["Comment"] = v
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	request["EngineVersion"] = d.Get("engine_version")
	if v, ok := d.GetOk("restore_size"); ok {
		request["RestoreSize"] = v
	}
	if v, ok := d.GetOk("bucket_region"); ok {
		request["BucketRegion"] = v
	}
	if v, ok := d.GetOkExists("build_replication"); ok {
		request["BuildReplication"] = v
	}
	if v, ok := d.GetOk("master_info"); ok {
		request["MasterInfo"] = v
	}
	if v, ok := d.GetOk("source_info"); ok {
		request["SourceInfo"] = v
	}
	if v, ok := d.GetOk("db_instance_id"); ok {
		request["DBInstanceId"] = v
	}
	if v, ok := d.GetOk("mode"); ok {
		request["Mode"] = v
	}
	if v, ok := d.GetOk("retention"); ok {
		request["Retention"] = v
	}
	if v, ok := d.GetOk("backup_file"); ok {
		request["BackupFile"] = v
	}
	if v, ok := d.GetOk("zone_id"); ok {
		request["ZoneId"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Rds", "2014-08-15", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_rds_backup_file", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["BackupId"]))

	return resourceAliCloudRdsBackupFileUpdate(d, meta)
}

func resourceAliCloudRdsBackupFileRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	rdsServiceV2 := RdsServiceV2{client}

	objectRaw, err := rdsServiceV2.DescribeRdsBackupFile(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_rds_backup_file DescribeRdsBackupFile Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("comment", objectRaw["Comment"])
	d.Set("create_time", objectRaw["CreationTime"])
	d.Set("engine_version", objectRaw["EngineVersion"])
	d.Set("restore_size", objectRaw["RestoreSize"])
	d.Set("retention", objectRaw["Retention"])
	d.Set("status", objectRaw["Status"])
	d.Set("zone_id", objectRaw["ZoneId"])

	return nil
}

func resourceAliCloudRdsBackupFileUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "UpdateUserBackupFile"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["BackupId"] = d.Id()
	request["RegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("comment") {
		update = true
		request["Comment"] = d.Get("comment")
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if !d.IsNewResource() && d.HasChange("retention") {
		update = true
		request["Retention"] = d.Get("retention")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Rds", "2014-08-15", action, query, request, true)
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

	return resourceAliCloudRdsBackupFileRead(d, meta)
}

func resourceAliCloudRdsBackupFileDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteUserBackupFile"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["BackupId"] = d.Id()
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Rds", "2014-08-15", action, query, request, true)
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
