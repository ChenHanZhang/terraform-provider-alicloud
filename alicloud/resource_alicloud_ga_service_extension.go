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

func resourceAliCloudGaServiceExtension() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudGaServiceExtensionCreate,
		Read:   resourceAliCloudGaServiceExtensionRead,
		Update: resourceAliCloudGaServiceExtensionUpdate,
		Delete: resourceAliCloudGaServiceExtensionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"components": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fail_policy": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service_component_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"timeout": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"config": {
							Type:     schema.TypeString,
							Required: true,
						},
						"create_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"component_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resources": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"create_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"update_time": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"associate_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"accelerator_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"tag": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceAliCloudGaServiceExtensionCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateServiceExtension"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("tag"); ok {
		tagMapsArray := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(v) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["Key"] = dataLoopTmp["key"]
			dataLoopMap["Value"] = dataLoopTmp["value"]
			tagMapsArray = append(tagMapsArray, dataLoopMap)
		}
		request["Tag"] = tagMapsArray
	}

	request["Name"] = d.Get("name")
	wait := incrementalWait(3*time.Second, 0*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Ga", "2019-11-20", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ga_service_extension", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["ServiceExtensionId"]))

	return resourceAliCloudGaServiceExtensionUpdate(d, meta)
}

func resourceAliCloudGaServiceExtensionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	gaServiceV2 := GaServiceV2{client}

	objectRaw, err := gaServiceV2.DescribeGaServiceExtension(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ga_service_extension DescribeGaServiceExtension Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("description", objectRaw["Description"])
	d.Set("name", objectRaw["Name"])

	return nil
}

func resourceAliCloudGaServiceExtensionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "UpdateServiceExtension"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ServiceExtensionId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if _, ok := d.GetOk("resource_group_id"); ok && !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
		request["ResourceGroupId"] = d.Get("resource_group_id")
	}

	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	if !d.IsNewResource() && d.HasChange("tag") {
		update = true
		if v, ok := d.GetOk("tag"); ok || d.HasChange("tag") {
			tagMapsArray := make([]interface{}, 0)
			for _, dataLoop := range convertToInterfaceArray(v) {
				dataLoopTmp := dataLoop.(map[string]interface{})
				dataLoopMap := make(map[string]interface{})
				dataLoopMap["Key"] = dataLoopTmp["key"]
				dataLoopMap["Value"] = dataLoopTmp["value"]
				tagMapsArray = append(tagMapsArray, dataLoopMap)
			}
			request["Tag"] = tagMapsArray
		}
	}

	if !d.IsNewResource() && d.HasChange("name") {
		update = true
	}
	request["Name"] = d.Get("name")
	if update {
		wait := incrementalWait(3*time.Second, 0*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ga", "2019-11-20", action, query, request, true)
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
	action = "AssociateResourcesWithServiceExtension"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ServiceExtensionId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("resources") {
		update = true
	}
	resourcesResourceTypeJsonPath, err := jsonpath.Get("$.resource_type", d.Get("resources"))
	if err == nil {
		request["ResourceType"] = resourcesResourceTypeJsonPath
	}

	resourcesDataList := make(map[string]interface{})

	if d.HasChange("resources") {
		update = true
	}
	resourceId1, _ := jsonpath.Get("$.resource_id", d.Get("resources"))
	if resourceId1 != nil && resourceId1 != "" {
		resourcesDataList["ResourceId"] = resourceId1
	}

	ResourcesMap := make([]interface{}, 0)
	ResourcesMap = append(ResourcesMap, resourcesDataList)
	request["Resources"] = ResourcesMap

	if d.HasChange("resources") {
		update = true
	}
	resourcesAcceleratorIdJsonPath, err := jsonpath.Get("$.accelerator_id", d.Get("resources"))
	if err == nil {
		request["AcceleratorId"] = resourcesAcceleratorIdJsonPath
	}

	if update {
		wait := incrementalWait(3*time.Second, 0*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ga", "2019-11-20", action, query, request, true)
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
	action = "DissociateResourcesFromServiceExtension"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ServiceExtensionId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	resourcesDataList := make(map[string]interface{})

	if d.HasChange("resources") {
		update = true
	}
	resourceId1, _ := jsonpath.Get("$.resource_id", d.Get("resources"))
	if resourceId1 != nil && resourceId1 != "" {
		resourcesDataList["ResourceId"] = resourceId1
	}

	ResourcesMap := make([]interface{}, 0)
	ResourcesMap = append(ResourcesMap, resourcesDataList)
	request["Resources"] = ResourcesMap

	if d.HasChange("resources") {
		update = true
	}
	resourcesAcceleratorIdJsonPath, err := jsonpath.Get("$.accelerator_id", d.Get("resources"))
	if err == nil {
		request["AcceleratorId"] = resourcesAcceleratorIdJsonPath
	}

	if update {
		wait := incrementalWait(3*time.Second, 0*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ga", "2019-11-20", action, query, request, true)
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
	action = "AddServiceComponentsToServiceExtension"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ServiceExtensionId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("components") {
		update = true
	}
	if v, ok := d.GetOk("components"); ok || d.HasChange("components") {
		componentsMapsArray := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(v) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["FailPolicy"] = dataLoopTmp["fail_policy"]
			dataLoopMap["Timeout"] = dataLoopTmp["timeout"]
			dataLoopMap["Priority"] = dataLoopTmp["priority"]
			dataLoopMap["Config"] = dataLoopTmp["config"]
			dataLoopMap["ServiceComponentId"] = dataLoopTmp["service_component_id"]
			componentsMapsArray = append(componentsMapsArray, dataLoopMap)
		}
		request["Components"] = componentsMapsArray
	}

	if update {
		wait := incrementalWait(3*time.Second, 0*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ga", "2019-11-20", action, query, request, true)
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
	action = "RemoveServiceComponentsFromServiceExtension"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ServiceExtensionId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	componentsDataList := make(map[string]interface{})

	if d.HasChange("components") {
		update = true
	}
	serviceComponentId1, _ := jsonpath.Get("$.service_component_id", d.Get("components"))
	if serviceComponentId1 != nil && serviceComponentId1 != "" {
		componentsDataList["ServiceComponentId"] = serviceComponentId1
	}

	ComponentsMap := make([]interface{}, 0)
	ComponentsMap = append(ComponentsMap, componentsDataList)
	request["Components"] = ComponentsMap

	if update {
		wait := incrementalWait(3*time.Second, 0*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ga", "2019-11-20", action, query, request, true)
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
	action = "UpdateServiceComponentsFromServiceExtension"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ServiceExtensionId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("components") {
		update = true
	}
	if v, ok := d.GetOk("components"); ok || d.HasChange("components") {
		componentsMapsArray := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(v) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["FailPolicy"] = dataLoopTmp["fail_policy"]
			dataLoopMap["Timeout"] = dataLoopTmp["timeout"]
			dataLoopMap["Priority"] = dataLoopTmp["priority"]
			dataLoopMap["Config"] = dataLoopTmp["config"]
			dataLoopMap["ServiceComponentId"] = dataLoopTmp["service_component_id"]
			componentsMapsArray = append(componentsMapsArray, dataLoopMap)
		}
		request["Components"] = componentsMapsArray
	}

	if update {
		wait := incrementalWait(3*time.Second, 0*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ga", "2019-11-20", action, query, request, true)
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
	return resourceAliCloudGaServiceExtensionRead(d, meta)
}

func resourceAliCloudGaServiceExtensionDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteServiceExtension"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["ServiceExtensionId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 0*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Ga", "2019-11-20", action, query, request, true)
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
