package alicloud

import (
	"fmt"
	"log"
	"strconv"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAlicloudHbrNasBackupPlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudHbrNasBackupPlanCreate,
		Read:   resourceAlicloudHbrNasBackupPlanRead,
		Update: resourceAlicloudHbrNasBackupPlanUpdate,
		Delete: resourceAlicloudHbrNasBackupPlanDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"backup_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"COMPLETE"}, false),
			},
			"bucket": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"detail": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disk_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"exclude": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"file_system_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"include": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nas_backup_plan_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"options": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"path": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"source_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"retention": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rule": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_region_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"destination_retention": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"disabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"do_copy": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"retention": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"rule_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"schedule": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"schedule": {
				Type:     schema.TypeString,
				Required: true,
			},
			"speed_limit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"udm_region_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"update_paths": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vault_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAlicloudHbrNasBackupPlanCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var response map[string]interface{}
	action := "CreateBackupPlan"
	request := make(map[string]interface{})
	conn, err := client.NewHbrClient()
	if err != nil {
		return WrapError(err)
	}
	request["BackupType"] = d.Get("backup_type")
	if v, ok := d.GetOk("bucket"); ok {
		request["Bucket"] = v
	}
	if v, ok := d.GetOk("create_time"); ok {
		request["CreateTime"] = v
	}
	if v, ok := d.GetOk("detail"); ok {
		request["Detail"] = v
	}
	if v, ok := d.GetOk("disk_id"); ok {
		request["DiskId"] = v
	}
	if v, ok := d.GetOk("exclude"); ok {
		request["Exclude"] = v
	}
	if v, ok := d.GetOk("file_system_id"); ok {
		request["FileSystemId"] = v
	}
	if v, ok := d.GetOk("include"); ok {
		request["Include"] = v
	}
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	request["PlanName"] = d.Get("nas_backup_plan_name")
	if v, ok := d.GetOk("options"); ok {
		request["Options"] = v
	}
	if v, ok := d.GetOk("prefix"); ok {
		request["Prefix"] = v
	}
	if v, ok := d.GetOk("resource"); ok {
		for resourcePtr, resource := range v.(*schema.Set).List() {
			resourceArg := resource.(map[string]interface{})
			request["Resource."+fmt.Sprint(resourcePtr+1)+".ResourceId"] = resourceArg["resource_id"]
			request["Resource."+fmt.Sprint(resourcePtr+1)+".SourceType"] = resourceArg["source_type"]
		}
	}
	if v, ok := d.GetOk("retention"); ok {
		request["Retention"] = v
	}
	if v, ok := d.GetOk("rule"); ok {
		for rulePtr, rule := range v.(*schema.Set).List() {
			ruleArg := rule.(map[string]interface{})
			request["Rule."+fmt.Sprint(rulePtr+1)+".DestinationRegionId"] = ruleArg["destination_region_id"]
			request["Rule."+fmt.Sprint(rulePtr+1)+".DestinationRetention"] = ruleArg["destination_retention"]
			request["Rule."+fmt.Sprint(rulePtr+1)+".Disabled"] = ruleArg["disabled"]
			request["Rule."+fmt.Sprint(rulePtr+1)+".DoCopy"] = ruleArg["do_copy"]
			request["Rule."+fmt.Sprint(rulePtr+1)+".Retention"] = ruleArg["retention"]
			request["Rule."+fmt.Sprint(rulePtr+1)+".RuleName"] = ruleArg["rule_name"]
			request["Rule."+fmt.Sprint(rulePtr+1)+".Schedule"] = ruleArg["schedule"]
		}
	}
	request["Schedule"] = d.Get("schedule")
	request["SourceType"] = "NAS"
	if v, ok := d.GetOk("speed_limit"); ok {
		request["SpeedLimit"] = v
	}
	if v, ok := d.GetOk("udm_region_id"); ok {
		request["UdmRegionId"] = v
	}
	if v, ok := d.GetOk("vault_id"); ok {
		request["VaultId"] = v
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2017-09-08"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_hbr_nas_backup_plan", action, AlibabaCloudSdkGoERROR)
	}
	if fmt.Sprint(response["Success"]) == "false" {
		return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
	}

	d.SetId(fmt.Sprint(response["PlanId"]))

	return resourceAlicloudHbrNasBackupPlanUpdate(d, meta)
}
func resourceAlicloudHbrNasBackupPlanRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	hbrService := HbrService{client}
	object, err := hbrService.DescribeHbrNasBackupPlan(d.Id())
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_hbr_nas_backup_plan hbrService.DescribeHbrNasBackupPlan Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}
	d.Set("backup_type", object["BackupType"])
	d.Set("create_time", fmt.Sprint(formatInt(object["CreateTime"])))
	d.Set("file_system_id", object["FileSystemId"])
	d.Set("nas_backup_plan_name", object["PlanName"])
	d.Set("options", object["Options"])
	if object["Paths"] != nil {
		d.Set("path", object["Paths"].(map[string]interface{})["Path"])
	}
	d.Set("retention", fmt.Sprint(formatInt(object["Retention"])))
	d.Set("schedule", object["Schedule"])
	d.Set("vault_id", object["VaultId"])
	return nil
}
func resourceAlicloudHbrNasBackupPlanUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	hbrService := HbrService{client}
	var response map[string]interface{}
	d.Partial(true)

	update := false
	request := map[string]interface{}{
		"PlanId": d.Id(),
	}
	if !d.IsNewResource() && d.HasChange("retention") {
		update = true
	}
	if v, ok := d.GetOk("retention"); ok {
		request["Retention"] = v
	}

	if v, ok := d.GetOk("vault_id"); ok {
		request["VaultId"] = v
	}
	if !d.IsNewResource() && d.HasChange("nas_backup_plan_name") {
		update = true
		request["PlanName"] = d.Get("nas_backup_plan_name")
	}
	if !d.IsNewResource() && d.HasChange("options") {
		update = true
		if v, ok := d.GetOk("options"); ok {
			request["Options"] = v
		}
	}
	if d.HasChange("path") {
		update = true
		if v, ok := d.GetOk("path"); ok {
			request["Path"] = v
			request["UpdatePaths"] = true
		}
	}
	if !d.IsNewResource() && d.HasChange("schedule") {
		update = true
		request["Schedule"] = d.Get("schedule")
	}
	request["SourceType"] = "NAS"
	if update {
		if v, ok := d.GetOk("detail"); ok {
			request["Detail"] = v
		}
		if v, ok := d.GetOk("exclude"); ok {
			request["Exclude"] = v
		}
		if v, ok := d.GetOk("include"); ok {
			request["Include"] = v
		}
		if v, ok := d.GetOk("prefix"); ok {
			request["Prefix"] = v
		}
		if v, ok := d.GetOk("resource"); ok {
			for resourcePtr, resource := range v.(*schema.Set).List() {
				resourceArg := resource.(map[string]interface{})
				request["Resource."+fmt.Sprint(resourcePtr+1)+".ResourceId"] = resourceArg["resource_id"]
				request["Resource."+fmt.Sprint(resourcePtr+1)+".SourceType"] = resourceArg["source_type"]
			}
		}
		if v, ok := d.GetOk("rule"); ok {
			for rulePtr, rule := range v.(*schema.Set).List() {
				ruleArg := rule.(map[string]interface{})
				request["Rule."+fmt.Sprint(rulePtr+1)+".DestinationRegionId"] = ruleArg["destination_region_id"]
				request["Rule."+fmt.Sprint(rulePtr+1)+".DestinationRetention"] = ruleArg["destination_retention"]
				request["Rule."+fmt.Sprint(rulePtr+1)+".Disabled"] = ruleArg["disabled"]
				request["Rule."+fmt.Sprint(rulePtr+1)+".DoCopy"] = ruleArg["do_copy"]
				request["Rule."+fmt.Sprint(rulePtr+1)+".Retention"] = ruleArg["retention"]
				request["Rule."+fmt.Sprint(rulePtr+1)+".RuleName"] = ruleArg["rule_name"]
				request["Rule."+fmt.Sprint(rulePtr+1)+".Schedule"] = ruleArg["schedule"]
			}
		}
		if v, ok := d.GetOk("speed_limit"); ok {
			request["SpeedLimit"] = v
		}
		if v, ok := d.GetOkExists("update_paths"); ok {
			request["UpdatePaths"] = v
		}
		action := "UpdateBackupPlan"
		conn, err := client.NewHbrClient()
		if err != nil {
			return WrapError(err)
		}
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2017-09-08"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
		if fmt.Sprint(response["Success"]) == "false" {
			return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
		}
		d.SetPartial("retention")
		d.SetPartial("vault_id")
		d.SetPartial("nas_backup_plan_name")
		d.SetPartial("options")
		d.SetPartial("path")
		d.SetPartial("schedule")
	}
	if d.HasChange("disabled") {
		object, err := hbrService.DescribeHbrNasBackupPlan(d.Id())
		if err != nil {
			return WrapError(err)
		}
		target := strconv.FormatBool(d.Get("disabled").(bool))
		if strconv.FormatBool(object["Disabled"].(bool)) != target {
			if target == "false" {
				request := map[string]interface{}{
					"PlanId": d.Id(),
				}
				request["VaultId"] = d.Get("vault_id")
				request["SourceType"] = "NAS"
			}
			action := "EnableBackupPlan"
			conn, err := client.NewHbrClient()
			if err != nil {
				return WrapError(err)
			}
			wait := incrementalWait(3*time.Second, 3*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2017-09-08"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
			if fmt.Sprint(response["Success"]) == "false" {
				return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
			}
		}
		if target == "true" {
			request := map[string]interface{}{
				"PlanId": d.Id(),
			}
			request["VaultId"] = d.Get("vault_id")
			request["SourceType"] = "NAS"
		}
		action := "DisableBackupPlan"
		conn, err := client.NewHbrClient()
		if err != nil {
			return WrapError(err)
		}
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2017-09-08"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
		if fmt.Sprint(response["Success"]) == "false" {
			return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
		}
	}
	d.SetPartial("disabled")
	d.Partial(false)
	return resourceAlicloudHbrNasBackupPlanRead(d, meta)
}
func resourceAlicloudHbrNasBackupPlanDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	action := "DeleteBackupPlan"
	var response map[string]interface{}
	conn, err := client.NewHbrClient()
	if err != nil {
		return WrapError(err)
	}
	request := map[string]interface{}{
		"PlanId": d.Id(),
	}

	request["SourceType"] = "NAS"
	if v, ok := d.GetOk("vault_id"); ok {
		request["VaultId"] = v
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2017-09-08"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
	if fmt.Sprint(response["Success"]) == "false" {
		return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
	}
	return nil
}
