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

func resourceAliCloudSlsIndex() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudSlsIndexCreate,
		Read:   resourceAliCloudSlsIndexRead,
		Update: resourceAliCloudSlsIndexUpdate,
		Delete: resourceAliCloudSlsIndexDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"keys": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"chn": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"case_sensitive": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"token": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"alias": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"doc_value": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"line": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"include_keys": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"chn": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"case_sensitive": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"exclude_keys": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"token": {
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"log_reduce": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"log_reduce_black_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"log_reduce_white_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"logstore_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"max_text_len": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"project_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudSlsIndexCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	logstore := d.Get("logstore_name")
	action := fmt.Sprintf("/logstores/%s/index", logstore)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	hostMap := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	hostMap["project"] = StringPointer(d.Get("project_name").(string))

	line := make(map[string]interface{})

	if v := d.Get("line"); !IsNil(v) {
		chn1, _ := jsonpath.Get("$[0].chn", v)
		if chn1 != nil && chn1 != "" {
			line["chn"] = chn1
		}
		caseSensitive1, _ := jsonpath.Get("$[0].case_sensitive", v)
		if caseSensitive1 != nil && caseSensitive1 != "" {
			line["caseSensitive"] = caseSensitive1
		}
		token1, _ := jsonpath.Get("$[0].token", v)
		if token1 != nil && token1 != "" {
			line["token"] = token1
		}
		excludeKeys, _ := jsonpath.Get("$[0].exclude_keys", v)
		if excludeKeys != nil && excludeKeys != "" {
			line["exclude_keys"] = excludeKeys
		}
		includeKeys, _ := jsonpath.Get("$[0].include_keys", v)
		if includeKeys != nil && includeKeys != "" {
			line["include_keys"] = includeKeys
		}

		request["line"] = line
	}

	if v, ok := d.GetOkExists("log_reduce"); ok {
		request["log_reduce"] = v
	}
	if v, ok := d.GetOkExists("max_text_len"); ok {
		request["max_text_len"] = v
	}
	if v, ok := d.GetOk("log_reduce_white_list"); ok {
		log_reduce_white_listMapsArray := convertToInterfaceArray(v)

		request["log_reduce_white_list"] = log_reduce_white_listMapsArray
	}

	if v, ok := d.GetOk("log_reduce_black_list"); ok {
		log_reduce_black_listMapsArray := convertToInterfaceArray(v)

		request["log_reduce_black_list"] = log_reduce_black_listMapsArray
	}

	if v, ok := d.GetOk("keys"); ok {
		request["keys"] = v
	}
	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.Do("Sls", roaParam("POST", "2020-12-30", "CreateIndex", action), query, body, nil, hostMap, false)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_sls_index", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", *hostMap["project"], logstore))

	return resourceAliCloudSlsIndexUpdate(d, meta)
}

func resourceAliCloudSlsIndexRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	slsServiceV2 := SlsServiceV2{client}

	objectRaw, err := slsServiceV2.DescribeSlsIndex(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_sls_index DescribeSlsIndex Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("keys", objectRaw["keys"])
	d.Set("log_reduce", objectRaw["log_reduce"])
	d.Set("max_text_len", objectRaw["max_text_len"])

	lineMaps := make([]map[string]interface{}, 0)
	lineMap := make(map[string]interface{})
	lineRaw := make(map[string]interface{})
	if objectRaw["line"] != nil {
		lineRaw = objectRaw["line"].(map[string]interface{})
	}
	if len(lineRaw) > 0 {
		lineMap["case_sensitive"] = lineRaw["caseSensitive"]
		lineMap["chn"] = lineRaw["chn"]

		exclude_keysRaw := make([]interface{}, 0)
		if lineRaw["exclude_keys"] != nil {
			exclude_keysRaw = convertToInterfaceArray(lineRaw["exclude_keys"])
		}

		lineMap["exclude_keys"] = exclude_keysRaw
		include_keysRaw := make([]interface{}, 0)
		if lineRaw["include_keys"] != nil {
			include_keysRaw = convertToInterfaceArray(lineRaw["include_keys"])
		}

		lineMap["include_keys"] = include_keysRaw
		tokenRaw := make([]interface{}, 0)
		if lineRaw["token"] != nil {
			tokenRaw = convertToInterfaceArray(lineRaw["token"])
		}

		lineMap["token"] = tokenRaw
		lineMaps = append(lineMaps, lineMap)
	}
	if err := d.Set("line", lineMaps); err != nil {
		return err
	}
	log_reduce_black_listRaw := make([]interface{}, 0)
	if objectRaw["log_reduce_black_list"] != nil {
		log_reduce_black_listRaw = convertToInterfaceArray(objectRaw["log_reduce_black_list"])
	}

	d.Set("log_reduce_black_list", log_reduce_black_listRaw)
	log_reduce_white_listRaw := make([]interface{}, 0)
	if objectRaw["log_reduce_white_list"] != nil {
		log_reduce_white_listRaw = convertToInterfaceArray(objectRaw["log_reduce_white_list"])
	}

	d.Set("log_reduce_white_list", log_reduce_white_listRaw)

	parts := strings.Split(d.Id(), ":")
	d.Set("project_name", parts[0])
	d.Set("logstore_name", parts[1])

	return nil
}

func resourceAliCloudSlsIndexUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var header map[string]*string
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	logstore := parts[1]
	action := fmt.Sprintf("/logstores/%s/index", logstore)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	hostMap := make(map[string]*string)
	hostMap["project"] = StringPointer(parts[0])

	if !d.IsNewResource() && d.HasChange("line") {
		update = true
	}
	line := make(map[string]interface{})

	if v := d.Get("line"); !IsNil(v) || d.HasChange("line") {
		chn1, _ := jsonpath.Get("$[0].chn", v)
		if chn1 != nil && chn1 != "" {
			line["chn"] = chn1
		}
		caseSensitive1, _ := jsonpath.Get("$[0].case_sensitive", v)
		if caseSensitive1 != nil && caseSensitive1 != "" {
			line["caseSensitive"] = caseSensitive1
		}
		token1, _ := jsonpath.Get("$[0].token", v)
		if token1 != nil && token1 != "" {
			line["token"] = token1
		}
		excludeKeys, _ := jsonpath.Get("$[0].exclude_keys", v)
		if excludeKeys != nil && excludeKeys != "" {
			line["exclude_keys"] = excludeKeys
		}
		includeKeys, _ := jsonpath.Get("$[0].include_keys", v)
		if includeKeys != nil && includeKeys != "" {
			line["include_keys"] = includeKeys
		}

		request["line"] = line
	}

	if !d.IsNewResource() && d.HasChange("max_text_len") {
		update = true
	}
	if v, ok := d.GetOkExists("max_text_len"); ok || d.HasChange("max_text_len") {
		request["max_text_len"] = v
	}
	if !d.IsNewResource() && d.HasChange("log_reduce") {
		update = true
	}
	if v, ok := d.GetOkExists("log_reduce"); ok || d.HasChange("log_reduce") {
		request["log_reduce"] = v
	}
	if !d.IsNewResource() && d.HasChange("log_reduce_white_list") {
		update = true
	}
	if v, ok := d.GetOk("log_reduce_white_list"); ok || d.HasChange("log_reduce_white_list") {
		log_reduce_white_listMapsArray := convertToInterfaceArray(v)

		request["log_reduce_white_list"] = log_reduce_white_listMapsArray
	}

	if !d.IsNewResource() && d.HasChange("log_reduce_black_list") {
		update = true
	}
	if v, ok := d.GetOk("log_reduce_black_list"); ok || d.HasChange("log_reduce_black_list") {
		log_reduce_black_listMapsArray := convertToInterfaceArray(v)

		request["log_reduce_black_list"] = log_reduce_black_listMapsArray
	}

	if !d.IsNewResource() && d.HasChange("keys") {
		update = true
	}
	if v, ok := d.GetOk("keys"); ok || d.HasChange("keys") {
		request["keys"] = v
	}
	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.Do("Sls", roaParam("PUT", "2020-12-30", "UpdateIndex", action), query, body, nil, hostMap, false)
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

	return resourceAliCloudSlsIndexRead(d, meta)
}

func resourceAliCloudSlsIndexDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	logstore := parts[1]
	action := fmt.Sprintf("/logstores/%s/index", logstore)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	hostMap := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	hostMap["project"] = StringPointer(parts[0])

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.Do("Sls", roaParam("DELETE", "2020-12-30", "DeleteIndex", action), query, nil, nil, hostMap, false)
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
