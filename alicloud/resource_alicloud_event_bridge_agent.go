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

func resourceAliCloudEventBridgeAgent() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudEventBridgeAgentCreate,
		Read:   resourceAliCloudEventBridgeAgentRead,
		Update: resourceAliCloudEventBridgeAgentUpdate,
		Delete: resourceAliCloudEventBridgeAgentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metadata": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attachments": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"arn": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"mime_type": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"prompt": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudEventBridgeAgentCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateAgent"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("name"); ok {
		request["Name"] = v
	}

	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("prompt"); ok {
		request["Prompt"] = v
	}
	metadata := make(map[string]interface{})

	if v := d.Get("metadata"); !IsNil(v) {
		localData, err := jsonpath.Get("$[0].attachments", v)
		if err != nil {
			localData = make([]interface{}, 0)
		}
		localMaps := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(localData) {
			dataLoopTmp := make(map[string]interface{})
			if dataLoop != nil {
				dataLoopTmp = dataLoop.(map[string]interface{})
			}
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["Arn"] = dataLoopTmp["arn"]
			dataLoopMap["MimeType"] = dataLoopTmp["mime_type"]
			localMaps = append(localMaps, dataLoopMap)
		}
		metadata["Attachments"] = localMaps

		metadataJson, err := json.Marshal(metadata)
		if err != nil {
			return WrapError(err)
		}
		request["Metadata"] = string(metadataJson)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_event_bridge_agent", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["Name"]))

	return resourceAliCloudEventBridgeAgentRead(d, meta)
}

func resourceAliCloudEventBridgeAgentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	eventBridgeServiceV2 := EventBridgeServiceV2{client}

	objectRaw, err := eventBridgeServiceV2.DescribeEventBridgeAgent(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_event_bridge_agent DescribeEventBridgeAgent Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("description", objectRaw["Description"])
	d.Set("prompt", objectRaw["Prompt"])
	d.Set("name", objectRaw["Name"])

	metadataMaps := make([]map[string]interface{}, 0)
	metadataMap := make(map[string]interface{})
	attachmentsRaw, _ := jsonpath.Get("$.Metadata.Attachments", objectRaw)

	attachmentsMaps := make([]map[string]interface{}, 0)
	if attachmentsRaw != nil {
		for _, attachmentsChildRaw := range convertToInterfaceArray(attachmentsRaw) {
			attachmentsMap := make(map[string]interface{})
			attachmentsChildRaw := attachmentsChildRaw.(map[string]interface{})
			attachmentsMap["arn"] = attachmentsChildRaw["Arn"]
			attachmentsMap["mime_type"] = attachmentsChildRaw["MimeType"]

			attachmentsMaps = append(attachmentsMaps, attachmentsMap)
		}
	}
	metadataMap["attachments"] = attachmentsMaps
	metadataMaps = append(metadataMaps, metadataMap)
	if err := d.Set("metadata", metadataMaps); err != nil {
		return err
	}

	d.Set("name", d.Id())

	return nil
}

func resourceAliCloudEventBridgeAgentUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "UpdateAgent"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["Name"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	if d.HasChange("prompt") {
		update = true
		request["Prompt"] = d.Get("prompt")
	}

	if d.HasChange("metadata") {
		update = true
		metadata := make(map[string]interface{})

		if v := d.Get("metadata"); v != nil {
			localData, err := jsonpath.Get("$[0].attachments", v)
			if err != nil {
				localData = make([]interface{}, 0)
			}
			localMaps := make([]interface{}, 0)
			for _, dataLoop := range convertToInterfaceArray(localData) {
				dataLoopTmp := make(map[string]interface{})
				if dataLoop != nil {
					dataLoopTmp = dataLoop.(map[string]interface{})
				}
				dataLoopMap := make(map[string]interface{})
				dataLoopMap["Arn"] = dataLoopTmp["arn"]
				dataLoopMap["MimeType"] = dataLoopTmp["mime_type"]
				localMaps = append(localMaps, dataLoopMap)
			}
			metadata["Attachments"] = localMaps

			metadataJson, err := json.Marshal(metadata)
			if err != nil {
				return WrapError(err)
			}
			request["Metadata"] = string(metadataJson)
		}
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

	return resourceAliCloudEventBridgeAgentRead(d, meta)
}

func resourceAliCloudEventBridgeAgentDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteAgent"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["Name"] = d.Id()

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
