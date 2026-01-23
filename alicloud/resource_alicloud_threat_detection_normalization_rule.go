// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudThreatDetectionNormalizationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudThreatDetectionNormalizationRuleCreate,
		Read:   resourceAliCloudThreatDetectionNormalizationRuleRead,
		Update: resourceAliCloudThreatDetectionNormalizationRuleUpdate,
		Delete: resourceAliCloudThreatDetectionNormalizationRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"extend_content_packed": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"extend_field_store_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lang": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"normalization_category_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"normalization_rule_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"normalization_rule_expression": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"normalization_rule_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"normalization_rule_ids": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"normalization_rule_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"normalization_rule_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"normalization_rule_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"normalization_rule_version": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"normalization_schema_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"order_field": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role_for": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vendor_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudThreatDetectionNormalizationRuleCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateNormalizationRule"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("normalization_rule_type"); ok {
		request["NormalizationRuleType"] = v
	}
	if v, ok := d.GetOk("lang"); ok {
		request["Lang"] = v
	}
	if v, ok := d.GetOk("extend_content_packed"); ok {
		request["ExtendContentPacked"] = v
	}
	if v, ok := d.GetOk("order_field"); ok {
		request["OrderField"] = v
	}
	if v, ok := d.GetOk("vendor_id"); ok {
		request["VendorId"] = v
	}
	if v, ok := d.GetOk("normalization_rule_format"); ok {
		request["NormalizationRuleFormat"] = v
	}
	if v, ok := d.GetOk("normalization_rule_description"); ok {
		request["NormalizationRuleDescription"] = v
	}
	if v, ok := d.GetOk("normalization_category_id"); ok {
		request["NormalizationCategoryId"] = v
	}
	if v, ok := d.GetOkExists("normalization_rule_version"); ok {
		request["NormalizationRuleVersion"] = v
	}
	if v, ok := d.GetOk("normalization_rule_mode"); ok {
		request["NormalizationRuleMode"] = v
	}
	if v, ok := d.GetOk("normalization_rule_ids"); ok {
		normalizationRuleIdsMapsArray := convertToInterfaceArray(v)

		normalizationRuleIdsMapsJson, err := json.Marshal(normalizationRuleIdsMapsArray)
		if err != nil {
			return WrapError(err)
		}
		request["NormalizationRuleIds"] = string(normalizationRuleIdsMapsJson)
	}

	if v, ok := d.GetOkExists("role_for"); ok {
		request["RoleFor"] = v
	}
	if v, ok := d.GetOk("product_id"); ok {
		request["ProductId"] = v
	}
	if v, ok := d.GetOk("extend_field_store_mode"); ok {
		request["ExtendFieldStoreMode"] = v
	}
	if v, ok := d.GetOk("normalization_rule_name"); ok {
		request["NormalizationRuleName"] = v
	}
	if v, ok := d.GetOk("normalization_rule_expression"); ok {
		request["NormalizationRuleExpression"] = v
	}
	if v, ok := d.GetOk("normalization_schema_id"); ok {
		request["NormalizationSchemaId"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("cloud-siem", "2024-12-12", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_threat_detection_normalization_rule", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["NormalizationRuleId"]))

	return resourceAliCloudThreatDetectionNormalizationRuleRead(d, meta)
}

func resourceAliCloudThreatDetectionNormalizationRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	threatDetectionServiceV2 := ThreatDetectionServiceV2{client}

	objectRaw, err := threatDetectionServiceV2.DescribeThreatDetectionNormalizationRule(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_threat_detection_normalization_rule DescribeThreatDetectionNormalizationRule Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("extend_content_packed", objectRaw["ExtendContentPacked"])
	d.Set("extend_field_store_mode", objectRaw["ExtendFieldStoreMode"])
	d.Set("normalization_category_id", objectRaw["NormalizationCategoryId"])
	d.Set("normalization_rule_description", objectRaw["NormalizationRuleDescription"])
	d.Set("normalization_rule_expression", objectRaw["NormalizationRuleExpression"])
	d.Set("normalization_rule_format", objectRaw["NormalizationRuleFormat"])
	d.Set("normalization_rule_mode", objectRaw["NormalizationRuleMode"])
	d.Set("normalization_rule_name", objectRaw["NormalizationRuleName"])
	d.Set("normalization_rule_type", objectRaw["NormalizationRuleType"])
	d.Set("normalization_rule_version", objectRaw["NormalizationRuleVersion"])
	d.Set("normalization_schema_id", objectRaw["NormalizationSchemaId"])
	d.Set("order_field", objectRaw["OrderField"])
	d.Set("product_id", objectRaw["ProductId"])
	d.Set("vendor_id", objectRaw["VendorId"])

	normalizationRuleIdsRaw := make([]interface{}, 0)
	if objectRaw["NormalizationRuleIds"] != nil {
		normalizationRuleIdsRaw = convertToInterfaceArray(objectRaw["NormalizationRuleIds"])
	}

	d.Set("normalization_rule_ids", normalizationRuleIdsRaw)

	return nil
}

func resourceAliCloudThreatDetectionNormalizationRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "UpdateNormalizationRule"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["NormalizationRuleId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("normalization_rule_type") {
		update = true
		request["NormalizationRuleType"] = d.Get("normalization_rule_type")
	}

	if v, ok := d.GetOk("lang"); ok {
		request["Lang"] = v
	}
	if d.HasChange("order_field") {
		update = true
		request["OrderField"] = d.Get("order_field")
	}

	if d.HasChange("vendor_id") {
		update = true
		request["VendorId"] = d.Get("vendor_id")
	}

	if d.HasChange("normalization_rule_format") {
		update = true
		request["NormalizationRuleFormat"] = d.Get("normalization_rule_format")
	}

	if d.HasChange("normalization_rule_description") {
		update = true
		request["NormalizationRuleDescription"] = d.Get("normalization_rule_description")
	}

	if d.HasChange("normalization_category_id") {
		update = true
		request["NormalizationCategoryId"] = d.Get("normalization_category_id")
	}

	if d.HasChange("normalization_rule_mode") {
		update = true
		request["NormalizationRuleMode"] = d.Get("normalization_rule_mode")
	}

	if d.HasChange("normalization_rule_ids") {
		update = true
		if v, ok := d.GetOk("normalization_rule_ids"); ok || d.HasChange("normalization_rule_ids") {
			normalizationRuleIdsMapsArray := convertToInterfaceArray(v)

			normalizationRuleIdsMapsJson, err := json.Marshal(normalizationRuleIdsMapsArray)
			if err != nil {
				return WrapError(err)
			}
			request["NormalizationRuleIds"] = string(normalizationRuleIdsMapsJson)
		}
	}

	if v, ok := d.GetOkExists("role_for"); ok {
		request["RoleFor"] = v
	}
	if d.HasChange("product_id") {
		update = true
		request["ProductId"] = d.Get("product_id")
	}

	if d.HasChange("extend_field_store_mode") {
		update = true
		request["ExtendFieldStoreMode"] = d.Get("extend_field_store_mode")
	}

	if d.HasChange("normalization_rule_name") {
		update = true
		request["NormalizationRuleName"] = d.Get("normalization_rule_name")
	}

	if d.HasChange("normalization_rule_expression") {
		update = true
		request["NormalizationRuleExpression"] = d.Get("normalization_rule_expression")
	}

	if d.HasChange("normalization_schema_id") {
		update = true
		request["NormalizationSchemaId"] = d.Get("normalization_schema_id")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("cloud-siem", "2024-12-12", action, query, request, true)
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
	action = "SetDefaultNormalizationRuleVersion"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["NormalizationRuleId"] = d.Id()
	request["RegionId"] = client.RegionId
	if v, ok := d.GetOk("lang"); ok {
		request["Lang"] = v
	}
	if d.HasChange("normalization_rule_version") {
		update = true
		request["NormalizationRuleVersion"] = d.Get("normalization_rule_version")
	}

	if v, ok := d.GetOkExists("role_for"); ok {
		request["RoleFor"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("cloud-siem", "2024-12-12", action, query, request, true)
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
	return resourceAliCloudThreatDetectionNormalizationRuleRead(d, meta)
}

func resourceAliCloudThreatDetectionNormalizationRuleDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteNormalizationRule"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["NormalizationRuleId"] = d.Id()
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("lang"); ok {
		request["Lang"] = v
	}
	if v, ok := d.GetOkExists("role_for"); ok {
		request["RoleFor"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("cloud-siem", "2024-12-12", action, query, request, true)
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
