// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudPolardbBackup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPolardbBackupCreate,
		Read:   resourceAliCloudPolardbBackupRead,
		Update: resourceAliCloudPolardbBackupUpdate,
		Delete: resourceAliCloudPolardbBackupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"backup_frequency": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_retention_policy_on_cluster_deletion": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data_level1_backup_frequency": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_level1_backup_period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_level1_backup_retention_period": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data_level1_backup_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_level2_backup_another_region_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_level2_backup_another_region_retention_period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_level2_backup_period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_level2_backup_retention_period": {
				Type:     schema.TypeString,
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
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudPolardbBackupCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateBackup"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("db_cluster_id"); ok {
		request["DBClusterId"] = v
	}

	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("data_level2_backup_retention_period"); ok {
		request["DataLevel2BackupRetentionPeriod"] = v
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_polardb_backup", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["DBClusterId"]))

	return resourceAliCloudPolardbBackupUpdate(d, meta)
}

func resourceAliCloudPolardbBackupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	polardbServiceV2 := PolardbServiceV2{client}

	objectRaw, err := polardbServiceV2.DescribePolardbBackup(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_polardb_backup DescribePolardbBackup Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	backupRawArrayObj, _ := jsonpath.Get("$.Items.Backup[*]", objectRaw)
	backupRawArray := make([]interface{}, 0)
	if backupRawArrayObj != nil {
		backupRawArray = convertToInterfaceArray(backupRawArrayObj)
	}
	backupRaw := make(map[string]interface{})
	if len(backupRawArray) > 0 {
		backupRaw = backupRawArray[0].(map[string]interface{})
	}

	d.Set("status", backupRaw["BackupStatus"])
	d.Set("db_cluster_id", backupRaw["DBClusterId"])

	objectRaw, err = polardbServiceV2.DescribeBackupDescribeBackupPolicy(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("backup_frequency", objectRaw["BackupFrequency"])
	d.Set("backup_retention_policy_on_cluster_deletion", objectRaw["BackupRetentionPolicyOnClusterDeletion"])
	d.Set("data_level1_backup_frequency", objectRaw["DataLevel1BackupFrequency"])
	d.Set("data_level1_backup_period", objectRaw["DataLevel1BackupPeriod"])
	d.Set("data_level1_backup_retention_period", objectRaw["DataLevel1BackupRetentionPeriod"])
	d.Set("data_level1_backup_time", objectRaw["DataLevel1BackupTime"])
	d.Set("data_level2_backup_another_region_region", objectRaw["DataLevel2BackupAnotherRegionRegion"])
	d.Set("data_level2_backup_another_region_retention_period", objectRaw["DataLevel2BackupAnotherRegionRetentionPeriod"])
	d.Set("data_level2_backup_period", objectRaw["DataLevel2BackupPeriod"])
	d.Set("data_level2_backup_retention_period", objectRaw["DataLevel2BackupRetentionPeriod"])
	d.Set("preferred_backup_period", objectRaw["PreferredBackupPeriod"])
	d.Set("preferred_backup_time", objectRaw["PreferredBackupTime"])

	objectRaw, err = polardbServiceV2.DescribeBackupDescribeBackupTasks(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["StartTime"])

	d.Set("db_cluster_id", d.Id())

	return nil
}

func resourceAliCloudPolardbBackupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "ModifyBackupPolicy"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if d.HasChange("data_level2_backup_another_region_region") {
		update = true
		request["DataLevel2BackupAnotherRegionRegion"] = d.Get("data_level2_backup_another_region_region")
	}

	if d.HasChange("data_level1_backup_period") {
		update = true
		request["DataLevel1BackupPeriod"] = d.Get("data_level1_backup_period")
	}

	if d.HasChange("data_level1_backup_time") {
		update = true
		request["DataLevel1BackupTime"] = d.Get("data_level1_backup_time")
	}

	if d.HasChange("preferred_backup_time") {
		update = true
		request["PreferredBackupTime"] = d.Get("preferred_backup_time")
	}

	if d.HasChange("preferred_backup_period") {
		update = true
		request["PreferredBackupPeriod"] = d.Get("preferred_backup_period")
	}

	if d.HasChange("data_level2_backup_period") {
		update = true
		request["DataLevel2BackupPeriod"] = d.Get("data_level2_backup_period")
	}

	if !d.IsNewResource() && d.HasChange("data_level2_backup_retention_period") {
		update = true
		request["DataLevel2BackupRetentionPeriod"] = d.Get("data_level2_backup_retention_period")
	}

	if d.HasChange("data_level1_backup_retention_period") {
		update = true
		request["DataLevel1BackupRetentionPeriod"] = d.Get("data_level1_backup_retention_period")
	}

	if d.HasChange("backup_retention_policy_on_cluster_deletion") {
		update = true
		request["BackupRetentionPolicyOnClusterDeletion"] = d.Get("backup_retention_policy_on_cluster_deletion")
	}

	if d.HasChange("backup_frequency") {
		update = true
		request["BackupFrequency"] = d.Get("backup_frequency")
	}

	if d.HasChange("data_level2_backup_another_region_retention_period") {
		update = true
		request["DataLevel2BackupAnotherRegionRetentionPeriod"] = d.Get("data_level2_backup_another_region_retention_period")
	}

	if d.HasChange("data_level1_backup_frequency") {
		update = true
		request["DataLevel1BackupFrequency"] = d.Get("data_level1_backup_frequency")
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

	return resourceAliCloudPolardbBackupRead(d, meta)
}

func resourceAliCloudPolardbBackupDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteBackup"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	request["BackupId"] = d.Get("backup_id")
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
