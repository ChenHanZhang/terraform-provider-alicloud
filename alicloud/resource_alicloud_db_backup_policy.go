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

func resourceAliCloudRdsBackupPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudRdsBackupPolicyCreate,
		Read:   resourceAliCloudRdsBackupPolicyRead,
		Update: resourceAliCloudRdsBackupPolicyUpdate,
		Delete: resourceAliCloudRdsBackupPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"archive_backup_keep_count": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"archive_backup_keep_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"archive_backup_retention_period": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backup_interval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backup_log": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backup_policy_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"backup_retention_period": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compress_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"db_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_backup_log": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enable_increment_data_backup": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"high_space_usage_protection": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_log_retention_hours": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"local_log_retention_space": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_backup_frequency": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_backup_local_retention_number": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"log_backup_retention_period": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"preferred_backup_period": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preferred_backup_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"released_keep_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudRdsBackupPolicyCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "ModifyBackupPolicy"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("db_instance_id"); ok {
		request["DBInstanceId"] = v
	}
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOkExists("backup_priority"); ok {
		request["BackupPriority"] = v
	}
	if v, ok := d.GetOkExists("enable_increment_data_backup"); ok {
		request["EnableIncrementDataBackup"] = v
	}
	if v, ok := d.GetOk("archive_backup_retention_period"); ok {
		request["ArchiveBackupRetentionPeriod"] = v
	}
	if v, ok := d.GetOk("preferred_backup_time"); ok {
		request["PreferredBackupTime"] = v
	}
	if v, ok := d.GetOk("compress_type"); ok {
		request["CompressType"] = v
	}
	if v, ok := d.GetOk("archive_backup_keep_count"); ok {
		request["ArchiveBackupKeepCount"] = v
	}
	if v, ok := d.GetOkExists("log_backup_retention_period"); ok {
		request["LogBackupRetentionPeriod"] = v
	}
	if v, ok := d.GetOkExists("log_backup_local_retention_number"); ok {
		request["LogBackupLocalRetentionNumber"] = v
	}
	if v, ok := d.GetOk("preferred_backup_period"); ok {
		request["PreferredBackupPeriod"] = v
	}
	if v, ok := d.GetOk("backup_log"); ok {
		request["BackupLog"] = v
	}
	if v, ok := d.GetOk("released_keep_policy"); ok {
		request["ReleasedKeepPolicy"] = v
	}
	if v, ok := d.GetOk("backup_method"); ok {
		request["BackupMethod"] = v
	}
	if v, ok := d.GetOkExists("backup_retention_period"); ok {
		request["BackupRetentionPeriod"] = v
	}
	if v, ok := d.GetOk("category"); ok {
		request["Category"] = v
	}
	if v, ok := d.GetOk("enable_backup_log"); ok {
		request["EnableBackupLog"] = v
	}
	if v, ok := d.GetOkExists("local_log_retention_hours"); ok {
		request["LocalLogRetentionHours"] = v
	}
	if v, ok := d.GetOk("local_log_retention_space"); ok {
		request["LocalLogRetentionSpace"] = v
	}
	if v, ok := d.GetOk("backup_policy_mode"); ok {
		request["BackupPolicyMode"] = v
	}
	if v, ok := d.GetOk("high_space_usage_protection"); ok {
		request["HighSpaceUsageProtection"] = v
	}
	if v, ok := d.GetOk("log_backup_frequency"); ok {
		request["LogBackupFrequency"] = v
	}
	if v, ok := d.GetOk("backup_interval"); ok {
		request["BackupInterval"] = v
	}
	if v, ok := d.GetOk("archive_backup_keep_policy"); ok {
		request["ArchiveBackupKeepPolicy"] = v
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_db_backup_policy", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["DBInstanceID"]))

	return resourceAliCloudRdsBackupPolicyRead(d, meta)
}

func resourceAliCloudRdsBackupPolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	rdsServiceV2 := RdsServiceV2{client}

	objectRaw, err := rdsServiceV2.DescribeRdsBackupPolicy(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_db_backup_policy DescribeRdsBackupPolicy Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("archive_backup_keep_count", objectRaw["ArchiveBackupKeepCount"])
	d.Set("archive_backup_keep_policy", objectRaw["ArchiveBackupKeepPolicy"])
	d.Set("archive_backup_retention_period", objectRaw["ArchiveBackupRetentionPeriod"])
	d.Set("backup_interval", objectRaw["BackupInterval"])
	d.Set("backup_log", objectRaw["BackupLog"])
	d.Set("backup_method", objectRaw["BackupMethod"])
	d.Set("backup_priority", objectRaw["BackupPriority"])
	d.Set("backup_retention_period", objectRaw["BackupRetentionPeriod"])
	d.Set("category", objectRaw["Category"])
	d.Set("compress_type", objectRaw["CompressType"])
	d.Set("enable_backup_log", objectRaw["EnableBackupLog"])
	d.Set("enable_increment_data_backup", objectRaw["EnableIncrementDataBackup"])
	d.Set("high_space_usage_protection", objectRaw["HighSpaceUsageProtection"])
	d.Set("local_log_retention_hours", objectRaw["LocalLogRetentionHours"])
	d.Set("local_log_retention_space", objectRaw["LocalLogRetentionSpace"])
	d.Set("log_backup_frequency", objectRaw["LogBackupFrequency"])
	d.Set("log_backup_local_retention_number", objectRaw["LogBackupLocalRetentionNumber"])
	d.Set("log_backup_retention_period", objectRaw["LogBackupRetentionPeriod"])
	d.Set("preferred_backup_period", objectRaw["PreferredBackupPeriod"])
	d.Set("preferred_backup_time", objectRaw["PreferredBackupTime"])
	d.Set("released_keep_policy", objectRaw["ReleasedKeepPolicy"])

	d.Set("db_instance_id", d.Id())

	return nil
}

func resourceAliCloudRdsBackupPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "ModifyBackupPolicy"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBInstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("backup_priority") {
		update = true
		request["BackupPriority"] = d.Get("backup_priority")
	}

	if d.HasChange("enable_increment_data_backup") {
		update = true
		request["EnableIncrementDataBackup"] = d.Get("enable_increment_data_backup")
	}

	if d.HasChange("archive_backup_retention_period") {
		update = true
		request["ArchiveBackupRetentionPeriod"] = d.Get("archive_backup_retention_period")
	}

	if d.HasChange("preferred_backup_time") {
		update = true
		request["PreferredBackupTime"] = d.Get("preferred_backup_time")
	}

	if d.HasChange("compress_type") {
		update = true
		request["CompressType"] = d.Get("compress_type")
	}

	if d.HasChange("archive_backup_keep_count") {
		update = true
		request["ArchiveBackupKeepCount"] = d.Get("archive_backup_keep_count")
	}

	if d.HasChange("log_backup_retention_period") {
		update = true
		request["LogBackupRetentionPeriod"] = d.Get("log_backup_retention_period")
	}

	if d.HasChange("log_backup_local_retention_number") {
		update = true
		request["LogBackupLocalRetentionNumber"] = d.Get("log_backup_local_retention_number")
	}

	if d.HasChange("preferred_backup_period") {
		update = true
		request["PreferredBackupPeriod"] = d.Get("preferred_backup_period")
	}

	if d.HasChange("backup_log") {
		update = true
		request["BackupLog"] = d.Get("backup_log")
	}

	if d.HasChange("released_keep_policy") {
		update = true
		request["ReleasedKeepPolicy"] = d.Get("released_keep_policy")
	}

	if d.HasChange("backup_method") {
		update = true
		request["BackupMethod"] = d.Get("backup_method")
	}

	if d.HasChange("backup_retention_period") {
		update = true
		request["BackupRetentionPeriod"] = d.Get("backup_retention_period")
	}

	if d.HasChange("category") {
		update = true
		request["Category"] = d.Get("category")
	}

	if d.HasChange("enable_backup_log") {
		update = true
		request["EnableBackupLog"] = d.Get("enable_backup_log")
	}

	if d.HasChange("local_log_retention_hours") {
		update = true
		request["LocalLogRetentionHours"] = d.Get("local_log_retention_hours")
	}

	if d.HasChange("local_log_retention_space") {
		update = true
		request["LocalLogRetentionSpace"] = d.Get("local_log_retention_space")
	}

	if v, ok := d.GetOk("backup_policy_mode"); ok {
		request["BackupPolicyMode"] = v
	}
	if d.HasChange("high_space_usage_protection") {
		update = true
		request["HighSpaceUsageProtection"] = d.Get("high_space_usage_protection")
	}

	if d.HasChange("log_backup_frequency") {
		update = true
		request["LogBackupFrequency"] = d.Get("log_backup_frequency")
	}

	if d.HasChange("backup_interval") {
		update = true
		request["BackupInterval"] = d.Get("backup_interval")
	}

	if d.HasChange("archive_backup_keep_policy") {
		update = true
		request["ArchiveBackupKeepPolicy"] = d.Get("archive_backup_keep_policy")
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

	return resourceAliCloudRdsBackupPolicyRead(d, meta)
}

func resourceAliCloudRdsBackupPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Backup Policy. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
