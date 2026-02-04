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
	"github.com/tidwall/sjson"
)

func resourceAliCloudCmsDigitalEmployee() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCmsDigitalEmployeeCreate,
		Read:   resourceAliCloudCmsDigitalEmployeeRead,
		Update: resourceAliCloudCmsDigitalEmployeeUpdate,
		Delete: resourceAliCloudCmsDigitalEmployeeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_rule": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"knowledges": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bailian": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"index_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"workspace_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"attributes": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"region": {
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
				Required: true,
				ForceNew: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role_arn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": tagsSchema(),
		},
	}
}

func resourceAliCloudCmsDigitalEmployeeCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/digital-employee")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("name"); ok {
		request["name"] = v
	}
	query["RegionId"] = StringPointer(client.RegionId)

	request["roleArn"] = d.Get("role_arn")
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["resourceGroupId"] = v
	}
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request["Tags"] = tagsMap
	}

	if v, ok := d.GetOk("default_rule"); ok {
		request["defaultRule"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["description"] = v
	}
	if v, ok := d.GetOk("display_name"); ok {
		request["displayName"] = v
	}
	knowledges := make(map[string]interface{})

	if v := d.Get("knowledges"); !IsNil(v) {
		if v, ok := d.GetOk("knowledges"); ok {
			localData1, err := jsonpath.Get("$[0].bailian", v)
			if err != nil {
				localData1 = make([]interface{}, 0)
			}
			localMaps := make([]interface{}, 0)
			for _, dataLoop1 := range convertToInterfaceArray(localData1) {
				dataLoop1Tmp := make(map[string]interface{})
				if dataLoop1 != nil {
					dataLoop1Tmp = dataLoop1.(map[string]interface{})
				}
				dataLoop1Map := make(map[string]interface{})
				dataLoop1Map["indexId"] = dataLoop1Tmp["index_id"]
				dataLoop1Map["attributes"] = dataLoop1Tmp["attributes"]
				dataLoop1Map["region"] = dataLoop1Tmp["region"]
				dataLoop1Map["workspaceId"] = dataLoop1Tmp["workspace_id"]
				localMaps = append(localMaps, dataLoop1Map)
			}
			knowledges["bailian"] = localMaps
		}

		request["knowledges"] = knowledges
	}

	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RoaPost("Cms", "2024-03-30", action, query, nil, body, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cms_digital_employee", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["name"]))

	return resourceAliCloudCmsDigitalEmployeeRead(d, meta)
}

func resourceAliCloudCmsDigitalEmployeeRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cmsServiceV2 := CmsServiceV2{client}

	objectRaw, err := cmsServiceV2.DescribeCmsDigitalEmployee(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cms_digital_employee DescribeCmsDigitalEmployee Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["createTime"])
	d.Set("default_rule", objectRaw["defaultRule"])
	d.Set("description", objectRaw["description"])
	d.Set("display_name", objectRaw["displayName"])
	d.Set("region_id", objectRaw["regionId"])
	d.Set("resource_group_id", objectRaw["resourceGroupId"])
	d.Set("role_arn", objectRaw["roleArn"])
	d.Set("name", objectRaw["name"])

	knowledgesMaps := make([]map[string]interface{}, 0)
	knowledgesMap := make(map[string]interface{})
	bailianRaw, _ := jsonpath.Get("$.knowledges.bailian", objectRaw)

	bailianMaps := make([]map[string]interface{}, 0)
	if bailianRaw != nil {
		for _, bailianChildRaw := range convertToInterfaceArray(bailianRaw) {
			bailianMap := make(map[string]interface{})
			bailianChildRaw := bailianChildRaw.(map[string]interface{})
			bailianMap["attributes"] = bailianChildRaw["attributes"]
			bailianMap["index_id"] = bailianChildRaw["indexId"]
			bailianMap["region"] = bailianChildRaw["region"]
			bailianMap["workspace_id"] = bailianChildRaw["workspaceId"]

			bailianMaps = append(bailianMaps, bailianMap)
		}
	}
	knowledgesMap["bailian"] = bailianMaps
	knowledgesMaps = append(knowledgesMaps, knowledgesMap)
	if err := d.Set("knowledges", knowledgesMaps); err != nil {
		return err
	}
	tagsMaps := objectRaw["tags"]
	d.Set("tags", tagsToMap(tagsMaps))

	d.Set("name", d.Id())

	return nil
}

func resourceAliCloudCmsDigitalEmployeeUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var header map[string]*string
	var query map[string]*string
	var body map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	name := d.Id()
	action := fmt.Sprintf("/digital-employee/%s", name)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	query["RegionId"] = StringPointer(client.RegionId)
	if d.HasChange("role_arn") {
		update = true
	}
	request["roleArn"] = d.Get("role_arn")
	if d.HasChange("knowledges") {
		update = true
	}
	knowledges := make(map[string]interface{})

	if v := d.Get("knowledges"); !IsNil(v) || d.HasChange("knowledges") {
		if v, ok := d.GetOk("knowledges"); ok {
			localData, err := jsonpath.Get("$[0].bailian", v)
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
				dataLoopMap["attributes"] = dataLoopTmp["attributes"]
				dataLoopMap["region"] = dataLoopTmp["region"]
				dataLoopMap["indexId"] = dataLoopTmp["index_id"]
				dataLoopMap["workspaceId"] = dataLoopTmp["workspace_id"]
				localMaps = append(localMaps, dataLoopMap)
			}
			knowledges["bailian"] = localMaps
		}

		request["knowledges"] = knowledges
	}

	if d.HasChange("default_rule") {
		update = true
	}
	if v, ok := d.GetOk("default_rule"); ok || d.HasChange("default_rule") {
		request["defaultRule"] = v
	}
	if d.HasChange("description") {
		update = true
	}
	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		request["description"] = v
	}
	if d.HasChange("display_name") {
		update = true
	}
	if v, ok := d.GetOk("display_name"); ok || d.HasChange("display_name") {
		request["displayName"] = v
	}
	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPatch("Cms", "2024-03-30", action, query, header, body, true)
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
	action = fmt.Sprintf("/resourcegroup")
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	request["resourceId"] = d.Id()
	query["RegionId"] = StringPointer(client.RegionId)
	if _, ok := d.GetOk("resource_group_id"); ok && d.HasChange("resource_group_id") {
		update = true
	}
	if v, ok := d.GetOk("resource_group_id"); ok || d.HasChange("resource_group_id") {
		request["resourceGroupId"] = v
	}
	request["resourceType"] = "ALIYUN::CMS::DIGITALEMPLOYEE"
	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPut("Cms", "2024-03-30", action, query, header, body, true)
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
	action = fmt.Sprintf("/tags")
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	query["resourceId"] = StringPointer(convertListToJsonString(expandSingletonToList(d.Id())))
	query["RegionId"] = StringPointer(client.RegionId)
	if d.HasChange("tags") {
		update = true
	}
	if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		query["Tags"] = tagsMap
	}

	query["resourceType"] = StringPointer("ALIYUN::CMS::DIGITALEMPLOYEE")
	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaDelete("Cms", "2024-03-30", action, query, header, nil, true)
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
	action = fmt.Sprintf("/tags")
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	query["RegionId"] = StringPointer(client.RegionId)
	if d.HasChange("tags") {
		update = true
	}
	if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request["Tags"] = tagsMap
	}

	request["resourceType"] = "ALIYUN::CMS::DIGITALEMPLOYEE"
	jsonString := convertObjectToJsonString(request)
	jsonString, _ = sjson.Set(jsonString, "resourceId.0", d.Id())
	_ = json.Unmarshal([]byte(jsonString), &request)

	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPost("Cms", "2024-03-30", action, query, header, body, true)
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

	d.Partial(false)
	return resourceAliCloudCmsDigitalEmployeeRead(d, meta)
}

func resourceAliCloudCmsDigitalEmployeeDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	name := d.Id()
	action := fmt.Sprintf("/digital-employee/%s", name)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	query["RegionId"] = StringPointer(client.RegionId)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RoaDelete("Cms", "2024-03-30", action, query, nil, nil, true)
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
		if IsExpectedErrors(err, []string{"DigitalEmployeeNotExist"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
