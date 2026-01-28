// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudDataWorksProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudDataWorksProjectCreate,
		Read:   resourceAliCloudDataWorksProjectRead,
		Update: resourceAliCloudDataWorksProjectUpdate,
		Delete: resourceAliCloudDataWorksProjectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dev_environment_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"dev_role_disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringMatch(regexp.MustCompile("^[\\w.,;/@-]+$"), "Workspace Display Name"),
			},
			"pai_task_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"project_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringMatch(regexp.MustCompile("^[\\w.,;/@-]+$"), "Workspace Name"),
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"Available", "Initializing", "Forbidden", "InitFailed", "Deleting", "DeleteFailed", "Frozen", "Updating", "UpdateFailed"}, false),
			},
			"tags": tagsSchema(),
		},
	}
}

func resourceAliCloudDataWorksProjectCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateProject"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId

	request["Name"] = d.Get("project_name")
	request["DisplayName"] = d.Get("display_name")
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["AliyunResourceGroupId"] = v
	}
	request["PaiTaskEnabled"] = d.Get("pai_task_enabled")
	if v, ok := d.GetOkExists("dev_environment_enabled"); ok {
		request["DevEnvironmentEnabled"] = v
	}
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request["Tags"] = tagsMap
	}

	if v, ok := d.GetOkExists("dev_role_disabled"); ok {
		request["DevRoleDisabled"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("dataworks-public", "2024-05-18", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_data_works_project", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["Id"]))

	dataWorksServiceV2 := DataWorksServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutCreate), 10*time.Second, dataWorksServiceV2.DataWorksProjectStateRefreshFunc(d.Id(), "$.Project.Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudDataWorksProjectUpdate(d, meta)
}

func resourceAliCloudDataWorksProjectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	dataWorksServiceV2 := DataWorksServiceV2{client}

	objectRaw, err := dataWorksServiceV2.DescribeDataWorksProject(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_data_works_project DescribeDataWorksProject Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	projectRawObj, _ := jsonpath.Get("$.Project", objectRaw)
	projectRaw := make(map[string]interface{})
	if projectRawObj != nil {
		projectRaw = projectRawObj.(map[string]interface{})
	}
	d.Set("description", projectRaw["Description"])
	d.Set("dev_environment_enabled", projectRaw["DevEnvironmentEnabled"])
	d.Set("dev_role_disabled", projectRaw["DevRoleDisabled"])
	d.Set("display_name", projectRaw["DisplayName"])
	d.Set("pai_task_enabled", projectRaw["PaiTaskEnabled"])
	d.Set("project_name", projectRaw["Name"])
	d.Set("resource_group_id", projectRaw["AliyunResourceGroupId"])
	d.Set("status", projectRaw["Status"])

	tagsMaps, _ := jsonpath.Get("$.Project.AliyunResourceTags", objectRaw)
	d.Set("tags", tagsToMap(tagsMaps))

	return nil
}

func resourceAliCloudDataWorksProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "ChangeResourceManagerResourceGroup"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ResourceId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ResourceType"] = "project"
	if _, ok := d.GetOk("resource_group_id"); ok && !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
	}
	request["ResourceManagerResourceGroupId"] = d.Get("resource_group_id")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("dataworks-public", "2020-05-18", action, query, request, true)
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
	action = "UpdateProject"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["Id"] = d.Id()
	request["RegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	if d.HasChange("status") {
		update = true
		request["Status"] = d.Get("status")
	}

	if !d.IsNewResource() && d.HasChange("dev_environment_enabled") {
		update = true
		request["DevEnvironmentEnabled"] = d.Get("dev_environment_enabled")
	}

	if !d.IsNewResource() && d.HasChange("dev_role_disabled") {
		update = true
		request["DevRoleDisabled"] = d.Get("dev_role_disabled")
	}

	if !d.IsNewResource() && d.HasChange("display_name") {
		update = true
	}
	request["DisplayName"] = d.Get("display_name")
	if !d.IsNewResource() && d.HasChange("pai_task_enabled") {
		update = true
	}
	request["PaiTaskEnabled"] = d.Get("pai_task_enabled")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("dataworks-public", "2024-05-18", action, query, request, true)
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
		dataWorksServiceV2 := DataWorksServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"Available", "Forbidden"}, d.Timeout(schema.TimeoutUpdate), 10*time.Second, dataWorksServiceV2.DataWorksProjectStateRefreshFunc(d.Id(), "$.Project.Status", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	if d.HasChange("tags") {
		dataWorksServiceV2 := DataWorksServiceV2{client}
		if err := dataWorksServiceV2.SetResourceTags(d, "project"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAliCloudDataWorksProjectRead(d, meta)
}

func resourceAliCloudDataWorksProjectDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteProject"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["Id"] = d.Id()
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("dataworks-public", "2024-05-18", action, query, request, true)
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
		if IsExpectedErrors(err, []string{"1101080008"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	dataWorksServiceV2 := DataWorksServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 10*time.Second, dataWorksServiceV2.DataWorksProjectStateRefreshFunc(d.Id(), "$.Project.Id", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
