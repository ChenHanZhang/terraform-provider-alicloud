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
			"enable_auth": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"http_api_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ingress_info": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ingress_class": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"cluster_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"override_ingress_ip": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"source_id": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "Field 'source_id' has been deprecated from provider version 1.283.0. Source ID",
						},
						"watch_namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"model_category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocols": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_backend": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"services": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
					},
				},
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

	if v, ok := d.GetOkExists("enable_auth"); ok {
		request["enableAuth"] = v
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["resourceGroupId"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["description"] = v
	}
	request["name"] = d.Get("http_api_name")
	if v, ok := d.GetOk("model_category"); ok {
		request["modelCategory"] = v
	}
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

	d.Set("description", objectRaw["description"])
	d.Set("http_api_name", objectRaw["name"])
	d.Set("type", objectRaw["type"])

	protocolsRaw := make([]interface{}, 0)
	if objectRaw["protocols"] != nil {
		protocolsRaw = convertToInterfaceArray(objectRaw["protocols"])
	}

	d.Set("protocols", protocolsRaw)
	routeBackendMaps := make([]map[string]interface{}, 0)
	routeBackendMap := make(map[string]interface{})
	routeBackendRawObj, _ := jsonpath.Get("$.deployConfigs[*].routeBackend", objectRaw)
	routeBackendRaw := make(map[string]interface{})
	if routeBackendRawObj != nil {
		routeBackendRaw = routeBackendRawObj.(map[string]interface{})
	}
	if len(routeBackendRaw) > 0 {

		servicesRaw, _ := jsonpath.Get("$.deployConfigs[*].routeBackend.services", objectRaw)
		servicesMaps := make([]map[string]interface{}, 0)
		if servicesRaw != nil {
			for _, servicesChildRaw := range convertToInterfaceArray(servicesRaw) {
				servicesMap := make(map[string]interface{})
				servicesChildRawObj, _ := jsonpath.Get("$.deployConfigs[*].routeBackend.services[*]", objectRaw)
				servicesChildRaw := make([]interface{}, 0)
				if servicesChildRawObj != nil {
					servicesChildRaw = convertToInterfaceArray(servicesChildRawObj)
				}

				servicesMap["name"] = servicesChildRaw["name"]

				servicesMaps = append(servicesMaps, servicesMap)
			}
		}
		routeBackendMap["services"] = servicesMaps
		routeBackendMaps = append(routeBackendMaps, routeBackendMap)
	}
	if err := d.Set("route_backend", routeBackendMaps); err != nil {
		return err
	}

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
	if d.HasChange("route_backend") {
		update = true
	}
	if v, ok := d.GetOk("route_backend"); ok || d.HasChange("route_backend") {
		localData1, _ := jsonpath.Get("$[0].s", v)
		deployConfigsMapsArray := make([]interface{}, 0)
		for _, dataLoop1 := range convertToInterfaceArray(localData1) {
			dataLoop1Tmp := dataLoop1.(interface{})
			dataLoop1Map := make(interface{})
			if !IsNil(dataLoop1Tmp["routeBackend"]) {
				localData2 := make(map[string]interface{})
				if len(localData2) > 0 {
					dataLoop1Map["routeBackend"] = localData2
				}
			}
			deployConfigsMapsArray = append(deployConfigsMapsArray, dataLoop1Map)
		}
		request["deployConfigs"] = deployConfigsMapsArray
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
