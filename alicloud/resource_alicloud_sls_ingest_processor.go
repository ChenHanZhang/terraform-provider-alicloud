// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudSlsIngestProcessor() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudSlsIngestProcessorCreate,
		Read:   resourceAliCloudSlsIngestProcessorRead,
		Update: resourceAliCloudSlsIngestProcessorUpdate,
		Delete: resourceAliCloudSlsIngestProcessorDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"configuration": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"parse_fail": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"spl": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"processor_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"project_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudSlsIngestProcessorCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	processorName := d.Get("processor_name")
	action := fmt.Sprintf("/ingestprocessors/%s", processorName)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	hostMap := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	hostMap["project"] = StringPointer(d.Get("project_name").(string))

	request["displayName"] = d.Get("display_name")
	if v, ok := d.GetOk("description"); ok {
		request["description"] = v
	}
	configuration := make(map[string]interface{})

	if v := d.Get("configuration"); v != nil {
		spl1, _ := jsonpath.Get("$[0].spl", v)
		if spl1 != nil && spl1 != "" {
			configuration["spl"] = spl1
		}
		parseFail1, _ := jsonpath.Get("$[0].parse_fail", v)
		if parseFail1 != nil && parseFail1 != "" {
			configuration["parseFail"] = parseFail1
		}

		request["configuration"] = configuration
	}

	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.Do("Sls", roaParam("PUT", "2020-12-30", "PutIngestProcessor", action), query, body, nil, hostMap, false)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_sls_ingest_processor", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", *hostMap["project"], processorName))

	return resourceAliCloudSlsIngestProcessorRead(d, meta)
}

func resourceAliCloudSlsIngestProcessorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	slsServiceV2 := SlsServiceV2{client}

	objectRaw, err := slsServiceV2.DescribeSlsIngestProcessor(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_sls_ingest_processor DescribeSlsIngestProcessor Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("description", objectRaw["description"])
	d.Set("display_name", objectRaw["displayName"])
	d.Set("processor_name", objectRaw["processorName"])

	configurationMaps := make([]map[string]interface{}, 0)
	configurationMap := make(map[string]interface{})
	configurationRaw := make(map[string]interface{})
	if objectRaw["configuration"] != nil {
		configurationRaw = objectRaw["configuration"].(map[string]interface{})
	}
	if len(configurationRaw) > 0 {
		configurationMap["parse_fail"] = configurationRaw["parseFail"]
		configurationMap["spl"] = configurationRaw["spl"]

		configurationMaps = append(configurationMaps, configurationMap)
	}
	if err := d.Set("configuration", configurationMaps); err != nil {
		return err
	}

	parts := strings.Split(d.Id(), ":")
	d.Set("project_name", parts[0])

	return nil
}

func resourceAliCloudSlsIngestProcessorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var header map[string]*string
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	processorName := parts[1]
	action := fmt.Sprintf("/ingestprocessors/%s", processorName)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	hostMap := make(map[string]*string)
	hostMap["project"] = StringPointer(parts[0])

	if d.HasChange("display_name") {
		update = true
	}
	request["displayName"] = d.Get("display_name")
	if d.HasChange("description") {
		update = true
	}
	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		request["description"] = v
	}
	if d.HasChange("configuration") {
		update = true
	}
	configuration := make(map[string]interface{})

	if v := d.Get("configuration"); v != nil {
		spl1, _ := jsonpath.Get("$[0].spl", v)
		if spl1 != nil && spl1 != "" {
			configuration["spl"] = spl1
		}
		parseFail1, _ := jsonpath.Get("$[0].parse_fail", v)
		if parseFail1 != nil && parseFail1 != "" {
			configuration["parseFail"] = parseFail1
		}

		request["configuration"] = configuration
	}

	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.Do("Sls", roaParam("PUT", "2020-12-30", "PutIngestProcessor", action), query, body, nil, hostMap, false)
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

	return resourceAliCloudSlsIngestProcessorRead(d, meta)
}

func resourceAliCloudSlsIngestProcessorDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	processorName := parts[1]
	action := fmt.Sprintf("/ingestprocessors/%s", processorName)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	hostMap := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	hostMap["project"] = StringPointer(parts[0])

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.Do("Sls", roaParam("DELETE", "2020-12-30", "DeleteIngestProcessor", action), query, nil, nil, hostMap, false)
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
