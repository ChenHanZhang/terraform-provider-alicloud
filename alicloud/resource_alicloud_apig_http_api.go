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

func resourceAliCloudApigHttpApi() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudApigHttpApiCreate,
		Read:   resourceAliCloudApigHttpApiRead,
		Update: resourceAliCloudApigHttpApiUpdate,
		Delete: resourceAliCloudApigHttpApiDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"ai_protocols": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"base_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deploy_configs": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"environments": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"environment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"gateway_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"gateway_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"API", "AI"}, false),
			},
			"http_api_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"model_category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocols": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudApigHttpApiCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/v1/http-apis")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	query["RegionId"] = StringPointer(client.RegionId)

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["resourceGroupId"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["description"] = v
	}
	request["name"] = d.Get("http_api_name")
	if v, ok := d.GetOk("environments"); ok {
		localData, _ := jsonpath.Get("$.s", v)
		deployConfigsMapsArray := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(localData) {
			dataLoopTmp := dataLoop.(interface{})
			dataLoopMap := make(interface{})
			dataLoopMap["environmentId"] = dataLoopTmp["environment_id"]
			dataLoopMap["gatewayId"] = dataLoopTmp[""]
			dataLoopMap["gatewayType"] = dataLoopTmp[""]
			localData1 := make(map[string]interface{})
			if len(localData1) > 0 {
				dataLoopMap["routeBackend"] = localData1
			}
			deployConfigsMapsArray = append(deployConfigsMapsArray, dataLoopMap)
		}
		request["deployConfigs"] = deployConfigsMapsArray
	}

	if v, ok := d.GetOk("model_category"); ok {
		request["modelCategory"] = v
	}
	if v, ok := d.GetOk("protocols"); ok {
		protocolsMapsArray := convertToInterfaceArray(v)

		request["protocols"] = protocolsMapsArray
	}

	if v, ok := d.GetOk("ai_protocols"); ok {
		aiProtocolsMapsArray := convertToInterfaceArray(v)

		request["aiProtocols"] = aiProtocolsMapsArray
	}

	if v, ok := d.GetOk("type"); ok {
		request["type"] = v
	}
	if v, ok := d.GetOk("base_path"); ok {
		request["basePath"] = v
	}
	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RoaPost("APIG", "2024-03-27", action, query, nil, body, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_apig_http_api", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.data.httpApiId", response)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudApigHttpApiRead(d, meta)
}

func resourceAliCloudApigHttpApiRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	apigServiceV2 := ApigServiceV2{client}

	objectRaw, err := apigServiceV2.DescribeApigHttpApi(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_apig_http_api DescribeApigHttpApi Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("base_path", objectRaw["basePath"])
	d.Set("description", objectRaw["description"])
	d.Set("gateway_id", objectRaw["gatewayId"])
	d.Set("http_api_name", objectRaw["name"])
	d.Set("resource_group_id", objectRaw["resourceGroupId"])
	d.Set("type", objectRaw["type"])

	deployConfigsChildRawObj, _ := jsonpath.Get("$.deployConfigs[*]", objectRaw)
	deployConfigsChildRaw := make([]interface{}, 0)
	if deployConfigsChildRawObj != nil {
		deployConfigsChildRaw = convertToInterfaceArray(deployConfigsChildRawObj)
	}

	d.Set("gateway_type", deployConfigsRaw["gatewayType"])

	environmentsRaw := objectRaw["environments"]
	environmentsMaps := make([]map[string]interface{}, 0)
	if environmentsRaw != nil {
		for _, environmentsChildRaw := range convertToInterfaceArray(environmentsRaw) {
			environmentsMap := make(map[string]interface{})
			environmentsChildRaw := environmentsChildRaw.(map[string]interface{})
			environmentsMap["environment_id"] = environmentsChildRaw["environmentId"]
			environmentsMap["name"] = environmentsChildRaw["name"]

			environmentsMaps = append(environmentsMaps, environmentsMap)
		}
	}
	if err := d.Set("environments", environmentsMaps); err != nil {
		return err
	}
	protocolsRaw := make([]interface{}, 0)
	if objectRaw["protocols"] != nil {
		protocolsRaw = convertToInterfaceArray(objectRaw["protocols"])
	}

	d.Set("protocols", protocolsRaw)

	return nil
}

func resourceAliCloudApigHttpApiUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	httpApiId := d.Id()
	action := fmt.Sprintf("/v1/http-apis/%s", httpApiId)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})

	if d.HasChange("protocols") {
		update = true
	}
	if v, ok := d.GetOk("protocols"); ok || d.HasChange("protocols") {
		protocolsMapsArray := convertToInterfaceArray(v)

		request["protocols"] = protocolsMapsArray
	}

	if d.HasChange("description") {
		update = true
	}
	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		request["description"] = v
	}
	if d.HasChange("ai_protocols") {
		update = true
	}
	if v, ok := d.GetOk("ai_protocols"); ok || d.HasChange("ai_protocols") {
		aiProtocolsJsonPath, err := jsonpath.Get("$", v)
		if err == nil && aiProtocolsJsonPath != "" {
			request["aiProtocols"] = aiProtocolsJsonPath
		}
	}
	if d.HasChange("deploy_configs") {
		update = true
	}
	if v, ok := d.GetOk("deploy_configs"); ok || d.HasChange("deploy_configs") {
		deployConfigsMapsArray := make([]interface{}, 0)
		for _, dataLoop1 := range convertToInterfaceArray(v) {
			dataLoop1Tmp := dataLoop1.(interface{})
			dataLoop1Map := make(interface{})
			dataLoop1Map["gatewayId"] = dataLoop1Tmp[""]
			localMaps := make([]interface{}, 0)
			localData2 := dataLoop1Tmp[""]
			for _, dataLoop2 := range convertToInterfaceArray(localData2) {
				dataLoop2Tmp := dataLoop2.(map[string]interface{})
				dataLoop2Map := make(map[string]interface{})
				dataLoop2Map["type"] = dataLoop2Tmp["$_type"]
				localMaps = append(localMaps, dataLoop2Map)
			}
			dataLoop1Map["policyConfigs"] = localMaps
			dataLoop1Map["gatewayType"] = dataLoop1Tmp[""]
			if !IsNil(dataLoop1Tmp["routeBackend"]) {
				localData3 := make(map[string]interface{})
				if len(localData3) > 0 {
					dataLoop1Map["routeBackend"] = localData3
				}
			}
			deployConfigsMapsArray = append(deployConfigsMapsArray, dataLoop1Map)
		}
		request["deployConfigs"] = deployConfigsMapsArray
	}

	if d.HasChange("base_path") {
		update = true
	}
	request["basePath"] = d.Get("base_path")
	if d.HasChange("environments") {
		update = true
	}
	ingressConfig := make(map[string]interface{})

	if v := d.Get("environments"); !IsNil(v) || d.HasChange("environments") {
		environmentId1, _ := jsonpath.Get("$.environment_id", v)
		if environmentId1 != nil && environmentId1 != "" {
			ingressConfig["environmentId"] = environmentId1
		}

		request["ingressConfig"] = ingressConfig
	}

	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPut("APIG", "2024-03-27", action, query, nil, body, true)
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
	action = fmt.Sprintf("/move-resource-group")
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	query["ResourceId"] = StringPointer(d.Id())
	query["RegionId"] = StringPointer(client.RegionId)
	if d.HasChange("resource_group_id") {
		update = true
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		query["ResourceGroupId"] = StringPointer(v.(string))
	}

	query["ResourceType"] = StringPointer("HttpApi")
	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPost("APIG", "2024-03-27", action, query, nil, body, true)
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

	return resourceAliCloudApigHttpApiRead(d, meta)
}

func resourceAliCloudApigHttpApiDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	httpApiId := d.Id()
	action := fmt.Sprintf("/v1/http-apis/%s", httpApiId)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	var err error
	request = make(map[string]interface{})

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RoaDelete("APIG", "2024-03-27", action, query, nil, nil, true)
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
		if IsExpectedErrors(err, []string{"Error.DatabaseError.RecordNotFound"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
