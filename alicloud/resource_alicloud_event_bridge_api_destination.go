// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudEventBridgeApiDestination() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudEventBridgeApiDestinationCreate,
		Read:   resourceAliCloudEventBridgeApiDestinationRead,
		Update: resourceAliCloudEventBridgeApiDestinationUpdate,
		Delete: resourceAliCloudEventBridgeApiDestinationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"api_destination_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"connection_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"create_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_api_parameters": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint": {
							Type:     schema.TypeString,
							Required: true,
						},
						"method": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceAliCloudEventBridgeApiDestinationCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateApiDestination"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("api_destination_name"); ok {
		request["ApiDestinationName"] = v
	}

	request["ConnectionName"] = d.Get("connection_name")
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	httpApiParameters := make(map[string]interface{})

	if v := d.Get("http_api_parameters"); v != nil {
		endpoint1, _ := jsonpath.Get("$[0].endpoint", v)
		if endpoint1 != nil && endpoint1 != "" {
			httpApiParameters["Endpoint"] = endpoint1
		}
		method1, _ := jsonpath.Get("$[0].method", v)
		if method1 != nil && method1 != "" {
			httpApiParameters["Method"] = method1
		}

		httpApiParametersJson, err := json.Marshal(httpApiParameters)
		if err != nil {
			return WrapError(err)
		}
		request["HttpApiParameters"] = string(httpApiParametersJson)
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("eventbridge", "2020-04-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_event_bridge_api_destination", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.Date.ApiDestinationName", response)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudEventBridgeApiDestinationRead(d, meta)
}

func resourceAliCloudEventBridgeApiDestinationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	eventBridgeServiceV2 := EventBridgeServiceV2{client}

	objectRaw, err := eventBridgeServiceV2.DescribeEventBridgeApiDestination(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_event_bridge_api_destination DescribeEventBridgeApiDestination Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("connection_name", objectRaw["ConnectionName"])
	d.Set("create_time", objectRaw["GmtCreate"])
	d.Set("description", objectRaw["Description"])
	d.Set("api_destination_name", objectRaw["ApiDestinationName"])

	httpApiParametersMaps := make([]map[string]interface{}, 0)
	httpApiParametersMap := make(map[string]interface{})
	httpApiParametersRaw := make(map[string]interface{})
	if objectRaw["HttpApiParameters"] != nil {
		httpApiParametersRaw = objectRaw["HttpApiParameters"].(map[string]interface{})
	}
	if len(httpApiParametersRaw) > 0 {
		httpApiParametersMap["endpoint"] = httpApiParametersRaw["Endpoint"]
		httpApiParametersMap["method"] = httpApiParametersRaw["Method"]

		httpApiParametersMaps = append(httpApiParametersMaps, httpApiParametersMap)
	}
	if err := d.Set("http_api_parameters", httpApiParametersMaps); err != nil {
		return err
	}

	d.Set("api_destination_name", d.Id())

	return nil
}

func resourceAliCloudEventBridgeApiDestinationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "UpdateApiDestination"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ApiDestinationName"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	if d.HasChange("connection_name") {
		update = true
	}
	request["ConnectionName"] = d.Get("connection_name")
	if d.HasChange("http_api_parameters") {
		update = true
	}
	httpApiParameters := make(map[string]interface{})

	if v := d.Get("http_api_parameters"); v != nil {
		endpoint1, _ := jsonpath.Get("$[0].endpoint", v)
		if endpoint1 != nil && endpoint1 != "" {
			httpApiParameters["Endpoint"] = endpoint1
		}
		method1, _ := jsonpath.Get("$[0].method", v)
		if method1 != nil && method1 != "" {
			httpApiParameters["Method"] = method1
		}

		httpApiParametersJson, err := json.Marshal(httpApiParameters)
		if err != nil {
			return WrapError(err)
		}
		request["HttpApiParameters"] = string(httpApiParametersJson)
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("eventbridge", "2020-04-01", action, query, request, true)
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

	return resourceAliCloudEventBridgeApiDestinationRead(d, meta)
}

func resourceAliCloudEventBridgeApiDestinationDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteApiDestination"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["ApiDestinationName"] = d.Id()

	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("eventbridge", "2020-04-01", action, query, request, true)
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
