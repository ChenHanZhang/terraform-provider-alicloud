// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAliCloudDataWorksDataSource() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudDataWorksDataSourceCreate,
		Read:   resourceAliCloudDataWorksDataSourceRead,
		Update: resourceAliCloudDataWorksDataSourceUpdate,
		Delete: resourceAliCloudDataWorksDataSourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"connection_properties": {
				Type:         schema.TypeString,
				Required:     true,
				Sensitive:    true,
				ValidateFunc: validation.ValidateJsonString,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					equal, _ := compareJsonTemplateAreEquivalent(old, new)
					return equal
				},
			},
			"connection_properties_mode": {
				Type:     schema.TypeString,
				Required: true,
			},
			"create_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"create_user": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_source_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"data_source_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"modify_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"modify_user": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"qualified_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudDataWorksDataSourceCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateDataSource"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("project_id"); ok {
		request["ProjectId"] = v
	}
	request["RegionId"] = client.RegionId

	request["Type"] = d.Get("type")
	request["ConnectionPropertiesMode"] = d.Get("connection_properties_mode")
	request["ConnectionProperties"] = d.Get("connection_properties")
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	request["Name"] = d.Get("data_source_name")
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_data_works_data_source", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["ProjectId"], response["Id"]))

	return resourceAliCloudDataWorksDataSourceRead(d, meta)
}

func resourceAliCloudDataWorksDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	dataWorksServiceV2 := DataWorksServiceV2{client}

	objectRaw, err := dataWorksServiceV2.DescribeDataWorksDataSource(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_data_works_data_source DescribeDataWorksDataSource Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("connection_properties", objectRaw["ConnectionProperties"])
	d.Set("connection_properties_mode", objectRaw["ConnectionPropertiesMode"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("create_user", objectRaw["CreateUser"])
	d.Set("data_source_name", objectRaw["Name"])
	d.Set("description", objectRaw["Description"])
	d.Set("modify_time", objectRaw["ModifyTime"])
	d.Set("modify_user", objectRaw["ModifyUser"])
	d.Set("qualified_name", objectRaw["QualifiedName"])
	d.Set("type", objectRaw["Type"])
	d.Set("data_source_id", objectRaw["Id"])
	d.Set("project_id", objectRaw["ProjectId"])

	return nil
}

func resourceAliCloudDataWorksDataSourceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "UpdateDataSource"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ProjectId"] = parts[0]
	request["Id"] = parts[1]
	request["RegionId"] = client.RegionId
	if d.HasChange("connection_properties_mode") {
		update = true
	}
	request["ConnectionPropertiesMode"] = d.Get("connection_properties_mode")
	if d.HasChange("connection_properties") {
		update = true
	}
	request["ConnectionProperties"] = d.Get("connection_properties")
	if d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

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
	}

	return resourceAliCloudDataWorksDataSourceRead(d, meta)
}

func resourceAliCloudDataWorksDataSourceDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteDataSource"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	query["Id"] = parts[1]
	query["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcGet("dataworks-public", "2024-05-18", action, query, request)
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
