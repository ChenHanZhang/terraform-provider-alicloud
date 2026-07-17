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

func resourceAliCloudApigService() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudApigServiceCreate,
		Read:   resourceAliCloudApigServiceRead,
		Update: resourceAliCloudApigServiceUpdate,
		Delete: resourceAliCloudApigServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"addresses": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ai_service_config": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"api_keys": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"protocols": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"address": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable_health_check": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"provider": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: StringInSlice([]string{"qwen", "moonshot", "baichuan", "yi", "zhipuai", "360ai", "hunyuan", "stepfun", "spark", "openai", "claude", "doubao", "minimax", "gemini", "azure", "deepseek", "custom", "pai-eas", "bedrock", "dify", "vertex", "vllm"}, false),
						},
					},
				},
			},
			"gateway_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"group_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"qualifier": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudApigServiceCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/v1/services")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	serviceConfigsDataList := make(map[string]interface{})
	if v, ok := d.GetOk("namespace"); ok {
		serviceConfigsDataList["namespace"] = v
	}
	if v, ok := d.GetOk("group_name"); ok {
		serviceConfigsDataList["groupName"] = v
	}
	if v := d.Get("ai_service_config"); !IsNil(v) {
		aiServiceConfig := make(map[string]interface{})
		apiKeys1, _ := jsonpath.Get("$[0].api_keys", d.Get("ai_service_config"))
		if apiKeys1 != nil && apiKeys1 != "" {
			aiServiceConfig["apiKeys"] = apiKeys1
		}
		protocols1, _ := jsonpath.Get("$[0].protocols", d.Get("ai_service_config"))
		if protocols1 != nil && protocols1 != "" {
			aiServiceConfig["protocols"] = protocols1
		}
		provider1, _ := jsonpath.Get("$[0].provider", d.Get("ai_service_config"))
		if provider1 != nil && provider1 != "" {
			aiServiceConfig["provider"] = provider1
		}
		enableHealthCheck1, _ := jsonpath.Get("$[0].enable_health_check", d.Get("ai_service_config"))
		if enableHealthCheck1 != nil && enableHealthCheck1 != "" {
			aiServiceConfig["enableHealthCheck"] = enableHealthCheck1
		}
		address1, _ := jsonpath.Get("$[0].address", d.Get("ai_service_config"))
		if address1 != nil && address1 != "" {
			aiServiceConfig["address"] = address1
		}

		if len(aiServiceConfig) > 0 {
			serviceConfigsDataList["aiServiceConfig"] = aiServiceConfig
		}
	}
	if v, ok := d.GetOk("qualifier"); ok {
		serviceConfigsDataList["qualifier"] = v
	}
	if v, ok := d.GetOk("addresses"); ok {
		addresses1, _ := jsonpath.Get("$", v)
		if addresses1 != nil && addresses1 != "" {
			serviceConfigsDataList["addresses"] = addresses1
		}
	}
	if v, ok := d.GetOk("service_name"); ok {
		serviceConfigsDataList["name"] = v
	}

	serviceConfigsMap := make([]interface{}, 0)
	serviceConfigsMap = append(serviceConfigsMap, serviceConfigsDataList)
	request["serviceConfigs"] = serviceConfigsMap

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["resourceGroupId"] = v
	}
	if v, ok := d.GetOk("source_type"); ok {
		request["sourceType"] = v
	}
	if v, ok := d.GetOk("gateway_id"); ok {
		request["gatewayId"] = v
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_apig_service", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.data.serviceIds[0]", response)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudApigServiceRead(d, meta)
}

func resourceAliCloudApigServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	apigServiceV2 := ApigServiceV2{client}

	objectRaw, err := apigServiceV2.DescribeApigService(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_apig_service DescribeApigService Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("gateway_id", objectRaw["gatewayId"])
	d.Set("group_name", objectRaw["groupName"])
	d.Set("namespace", objectRaw["namespace"])
	d.Set("qualifier", objectRaw["qualifier"])
	d.Set("resource_group_id", objectRaw["resourceGroupId"])
	d.Set("service_name", objectRaw["name"])
	d.Set("source_type", objectRaw["sourceType"])

	addressesRaw := make([]interface{}, 0)
	if objectRaw["addresses"] != nil {
		addressesRaw = convertToInterfaceArray(objectRaw["addresses"])
	}

	d.Set("addresses", addressesRaw)
	aiServiceConfigMaps := make([]map[string]interface{}, 0)
	aiServiceConfigMap := make(map[string]interface{})
	aiServiceConfigRaw := make(map[string]interface{})
	if objectRaw["aiServiceConfig"] != nil {
		aiServiceConfigRaw = objectRaw["aiServiceConfig"].(map[string]interface{})
	}
	if len(aiServiceConfigRaw) > 0 {
		aiServiceConfigMap["address"] = aiServiceConfigRaw["address"]
		aiServiceConfigMap["enable_health_check"] = aiServiceConfigRaw["enableHealthCheck"]
		aiServiceConfigMap["provider"] = aiServiceConfigRaw["provider"]

		apiKeysRaw := make([]interface{}, 0)
		if aiServiceConfigRaw["apiKeys"] != nil {
			apiKeysRaw = convertToInterfaceArray(aiServiceConfigRaw["apiKeys"])
		}

		aiServiceConfigMap["api_keys"] = apiKeysRaw
		protocolsRaw := make([]interface{}, 0)
		if aiServiceConfigRaw["protocols"] != nil {
			protocolsRaw = convertToInterfaceArray(aiServiceConfigRaw["protocols"])
		}

		aiServiceConfigMap["protocols"] = protocolsRaw
		aiServiceConfigMaps = append(aiServiceConfigMaps, aiServiceConfigMap)
	}
	if err := d.Set("ai_service_config", aiServiceConfigMaps); err != nil {
		return err
	}

	return nil
}

func resourceAliCloudApigServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	serviceId := d.Id()
	action := fmt.Sprintf("/v1/services/%s", serviceId)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})

	if d.HasChange("ai_service_config") {
		update = true
	}
	aiServiceConfig := make(map[string]interface{})

	if v := d.Get("ai_service_config"); !IsNil(v) || d.HasChange("ai_service_config") {
		protocols1, _ := jsonpath.Get("$[0].protocols", v)
		if protocols1 != nil && protocols1 != "" {
			aiServiceConfig["protocols"] = protocols1
		}
		provider1, _ := jsonpath.Get("$[0].provider", v)
		if provider1 != nil && provider1 != "" {
			aiServiceConfig["provider"] = provider1
		}
		apiKeys1, _ := jsonpath.Get("$[0].api_keys", v)
		if apiKeys1 != nil && apiKeys1 != "" {
			aiServiceConfig["apiKeys"] = apiKeys1
		}
		enableHealthCheck1, _ := jsonpath.Get("$[0].enable_health_check", v)
		if enableHealthCheck1 != nil && enableHealthCheck1 != "" {
			aiServiceConfig["enableHealthCheck"] = enableHealthCheck1
		}
		address1, _ := jsonpath.Get("$[0].address", v)
		if address1 != nil && address1 != "" {
			aiServiceConfig["address"] = address1
		}

		request["aiServiceConfig"] = aiServiceConfig
	}

	if d.HasChange("addresses") {
		update = true
	}
	if v, ok := d.GetOk("addresses"); ok || d.HasChange("addresses") {
		addressesMapsArray := convertToInterfaceArray(v)

		request["addresses"] = addressesMapsArray
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

	query["Service"] = StringPointer("APIG")
	query["ResourceType"] = StringPointer("Service")
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

	return resourceAliCloudApigServiceRead(d, meta)
}

func resourceAliCloudApigServiceDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	serviceId := d.Id()
	action := fmt.Sprintf("/v1/services/%s", serviceId)
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
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
