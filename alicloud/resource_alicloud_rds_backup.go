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

func resourceAliCloudRdsBackup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudRdsBackupCreate,
		Read:   resourceAliCloudRdsBackupRead,
		Update: resourceAliCloudRdsBackupUpdate,
		Delete: resourceAliCloudRdsBackupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"backup_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"backup_retention_period": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"backup_strategy": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"db_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"store_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudRdsBackupCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateBackup"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("backup_method"); ok {
		request["BackupMethod"] = v
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if v, ok := d.GetOkExists("backup_retention_period"); ok {
		request["BackupRetentionPeriod"] = v
	}
	request["DBInstanceId"] = d.Get("db_instance_id")
	if v, ok := d.GetOk("backup_strategy"); ok {
		request["BackupStrategy"] = v
	}
	if v, ok := d.GetOk("backup_type"); ok {
		request["BackupType"] = v
	}
	if v, ok := d.GetOk("db_name"); ok {
		request["DBName"] = v
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_rds_backup", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["BackupJobId"]))

	return resourceAliCloudRdsBackupRead(d, meta)
}

func resourceAliCloudRdsBackupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	rdsServiceV2 := RdsServiceV2{client}

	objectRaw, err := rdsServiceV2.DescribeRdsBackup(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_rds_backup DescribeRdsBackup Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("backup_method", objectRaw["BackupMethod"])
	d.Set("backup_type", objectRaw["BackupType"])
	d.Set("db_instance_id", objectRaw["DBInstanceId"])
	d.Set("resource_group_id", objectRaw["ResourceGroupId"])
	d.Set("status", objectRaw["MetaStatus"])
	d.Set("store_status", objectRaw["StoreStatus"])
	d.Set("backup_id", objectRaw["BackupId"])

	return nil
}

func resourceAliCloudRdsBackupUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Cannot update resource Alicloud Resource Backup.")
	return nil
}

func resourceAliCloudRdsBackupDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteBackup"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["BackupId"] = d.Id()
	request["RegionId"] = client.RegionId

	request["DBInstanceId"] = d.Get("db_instance_id")
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
