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

func resourceAliCloudPaiWorkspaceConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPaiWorkspaceConnectionCreate,
		Read:   resourceAliCloudPaiWorkspaceConnectionRead,
		Update: resourceAliCloudPaiWorkspaceConnectionUpdate,
		Delete: resourceAliCloudPaiWorkspaceConnectionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"accessibility": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"configs": {
				Type:     schema.TypeMap,
				Required: true,
			},
			"connection_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: StringMatch(regexp.MustCompile("^[0-9a-zA-Z_-]+$"), "The connection ID. For information about how to obtain a connection ID, see [ListConnections](url).  "),
			},
			"connection_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"connection_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"models": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tool_call": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"model": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"model_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"resource_meta": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"secrets": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"validate_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringMatch(regexp.MustCompile("^[0-9a-zA-Z_-]+$"), "The validation type used when verifying the connection. Valid values:  - Connectivity: connectivity test.  "),
			},
			"workspace_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudPaiWorkspaceConnectionCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/api/v1/connections")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	if v, ok := d.GetOk("models"); ok {
		modelsMapsArray := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(v) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["DisplayName"] = dataLoopTmp["display_name"]
			dataLoopMap["ModelType"] = dataLoopTmp["model_type"]
			dataLoopMap["Model"] = dataLoopTmp["model"]
			dataLoopMap["ToolCall"] = dataLoopTmp["tool_call"]
			modelsMapsArray = append(modelsMapsArray, dataLoopMap)
		}
		request["Models"] = modelsMapsArray
	}

	if v, ok := d.GetOk("accessibility"); ok {
		request["Accessibility"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("secrets"); ok {
		request["Secrets"] = v
	}
	resourceMeta := make(map[string]interface{})

	if v := d.Get("resource_meta"); !IsNil(v) {
		instanceId1, _ := jsonpath.Get("$[0].instance_id", v)
		if instanceId1 != nil && instanceId1 != "" {
			resourceMeta["InstanceId"] = instanceId1
		}
		instanceName1, _ := jsonpath.Get("$[0].instance_name", v)
		if instanceName1 != nil && instanceName1 != "" {
			resourceMeta["InstanceName"] = instanceName1
		}

		request["ResourceMeta"] = resourceMeta
	}

	request["ConnectionType"] = d.Get("connection_type")
	request["WorkspaceId"] = d.Get("workspace_id")
	request["Configs"] = d.Get("configs")
	request["ConnectionName"] = d.Get("connection_name")
	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RoaPost("AIWorkSpace", "2021-02-04", action, query, nil, body, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_pai_workspace_connection", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["ConnectionId"]))

	action = fmt.Sprintf("/api/v1/connections/validate")
	request = make(map[string]interface{})
	if v, ok := d.GetOk("connection_id"); ok {
		request["ConnectionId"] = v
	}

	if v, ok := d.GetOk("validate_type"); ok {
		request["ValidateType"] = v
	}
	if v, ok := d.GetOk("secrets"); ok {
		request["Secrets"] = v
	}
	request["WorkspaceId"] = d.Get("workspace_id")
	request["Configs"] = d.Get("configs")
	request["ConnectionType"] = d.Get("connection_type")
	body = request
	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RoaPost("AIWorkSpace", "2021-02-04", action, query, nil, body, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_pai_workspace_connection", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["ConnectionId"]))

	return resourceAliCloudPaiWorkspaceConnectionRead(d, meta)
}

func resourceAliCloudPaiWorkspaceConnectionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	paiWorkspaceServiceV2 := PaiWorkspaceServiceV2{client}

	objectRaw, err := paiWorkspaceServiceV2.DescribePaiWorkspaceConnection(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_pai_workspace_connection DescribePaiWorkspaceConnection Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("accessibility", objectRaw["Accessibility"])
	d.Set("configs", objectRaw["Configs"])
	d.Set("connection_name", objectRaw["ConnectionName"])
	d.Set("connection_type", objectRaw["ConnectionType"])
	d.Set("description", objectRaw["Description"])
	d.Set("secrets", objectRaw["Secrets"])
	d.Set("workspace_id", objectRaw["WorkspaceId"])

	modelsRaw := objectRaw["Models"]
	modelsMaps := make([]map[string]interface{}, 0)
	if modelsRaw != nil {
		for _, modelsChildRaw := range convertToInterfaceArray(modelsRaw) {
			modelsMap := make(map[string]interface{})
			modelsChildRaw := modelsChildRaw.(map[string]interface{})
			modelsMap["display_name"] = modelsChildRaw["DisplayName"]
			modelsMap["model"] = modelsChildRaw["Model"]
			modelsMap["model_type"] = modelsChildRaw["ModelType"]
			modelsMap["tool_call"] = modelsChildRaw["ToolCall"]

			modelsMaps = append(modelsMaps, modelsMap)
		}
	}
	if err := d.Set("models", modelsMaps); err != nil {
		return err
	}
	resourceMetaMaps := make([]map[string]interface{}, 0)
	resourceMetaMap := make(map[string]interface{})
	resourceMetaRaw := make(map[string]interface{})
	if objectRaw["ResourceMeta"] != nil {
		resourceMetaRaw = objectRaw["ResourceMeta"].(map[string]interface{})
	}
	if len(resourceMetaRaw) > 0 {
		resourceMetaMap["instance_id"] = resourceMetaRaw["InstanceId"]
		resourceMetaMap["instance_name"] = resourceMetaRaw["InstanceName"]

		resourceMetaMaps = append(resourceMetaMaps, resourceMetaMap)
	}
	if err := d.Set("resource_meta", resourceMetaMaps); err != nil {
		return err
	}

	d.Set("connection_id", d.Id())

	return nil
}

func resourceAliCloudPaiWorkspaceConnectionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var header map[string]*string
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	ConnectionId := d.Id()
	action := fmt.Sprintf("/api/v1/connections/%s", ConnectionId)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})

	if d.HasChange("models") {
		update = true
	}
	if v, ok := d.GetOk("models"); ok || d.HasChange("models") {
		modelsMapsArray := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(v) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["DisplayName"] = dataLoopTmp["display_name"]
			dataLoopMap["ModelType"] = dataLoopTmp["model_type"]
			dataLoopMap["Model"] = dataLoopTmp["model"]
			dataLoopMap["ToolCall"] = dataLoopTmp["tool_call"]
			modelsMapsArray = append(modelsMapsArray, dataLoopMap)
		}
		request["Models"] = modelsMapsArray
	}

	if d.HasChange("description") {
		update = true
	}
	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		request["Description"] = v
	}
	if d.HasChange("secrets") {
		update = true
	}
	if v, ok := d.GetOk("secrets"); ok || d.HasChange("secrets") {
		request["Secrets"] = v
	}
	if d.HasChange("configs") {
		update = true
	}
	request["Configs"] = d.Get("configs")
	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPut("AIWorkSpace", "2021-02-04", action, query, header, body, true)
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

	return resourceAliCloudPaiWorkspaceConnectionRead(d, meta)
}

func resourceAliCloudPaiWorkspaceConnectionDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	ConnectionId := d.Id()
	action := fmt.Sprintf("/api/v1/connections/%s", ConnectionId)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	var err error
	request = make(map[string]interface{})

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RoaDelete("AIWorkSpace", "2021-02-04", action, query, nil, nil, true)
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
		if IsExpectedErrors(err, []string{"RequestFailedErrorProblem"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
