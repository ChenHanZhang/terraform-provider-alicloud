// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudCmsIntegrationPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCmsIntegrationPolicyCreate,
		Read:   resourceAliCloudCmsIntegrationPolicyRead,
		Update: resourceAliCloudCmsIntegrationPolicyUpdate,
		Delete: resourceAliCloudCmsIntegrationPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"entity_group": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"cluster_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"cluster_entity_type": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"entity_group_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"fee_package": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"force": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"integration_policy_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_type": {
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
			"tags": tagsSchema(),
			"workspace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudCmsIntegrationPolicyCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/integration-policies")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	entityGroup := make(map[string]interface{})

	if v := d.Get("entity_group"); !IsNil(v) {
		clusterId1, _ := jsonpath.Get("$[0].cluster_id", v)
		if clusterId1 != nil && clusterId1 != "" {
			entityGroup["clusterId"] = clusterId1
		}
		vpcId1, _ := jsonpath.Get("$[0].vpc_id", v)
		if vpcId1 != nil && vpcId1 != "" {
			entityGroup["vpcId"] = vpcId1
		}
		entityGroupId1, _ := jsonpath.Get("$[0].entity_group_id", v)
		if entityGroupId1 != nil && entityGroupId1 != "" {
			entityGroup["entityGroupId"] = entityGroupId1
		}
		clusterEntityType1, _ := jsonpath.Get("$[0].cluster_entity_type", v)
		if clusterEntityType1 != nil && clusterEntityType1 != "" {
			entityGroup["clusterEntityType"] = clusterEntityType1
		}

		request["entityGroup"] = entityGroup
	}

	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request["Tags"] = tagsMap
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["resourceGroupId"] = v
	}
	request["policyType"] = d.Get("policy_type")
	if v, ok := d.GetOk("integration_policy_name"); ok {
		request["policyName"] = v
	}
	request["workspace"] = d.Get("workspace")
	body = request
	wait := incrementalWait(3*time.Second, 0*time.Second)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cms_integration_policy", action, AlibabaCloudSdkGoERROR)
	}
	if fmt.Sprint(response["code"]) != "200" {
		return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
	}

	id, _ := jsonpath.Get("$.policy.policyId", response)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudCmsIntegrationPolicyRead(d, meta)
}

func resourceAliCloudCmsIntegrationPolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cmsServiceV2 := CmsServiceV2{client}

	objectRaw, err := cmsServiceV2.DescribeCmsIntegrationPolicy(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cms_integration_policy DescribeCmsIntegrationPolicy Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	policyRawObj, _ := jsonpath.Get("$.policy", objectRaw)
	policyRaw := make(map[string]interface{})
	if policyRawObj != nil {
		policyRaw = policyRawObj.(map[string]interface{})
	}
	d.Set("integration_policy_name", policyRaw["policyName"])
	d.Set("policy_type", policyRaw["policyType"])
	d.Set("region_id", policyRaw["regionId"])
	d.Set("resource_group_id", policyRaw["resourceGroupId"])
	d.Set("workspace", policyRaw["workspace"])

	entityGroupMaps := make([]map[string]interface{}, 0)
	entityGroupMap := make(map[string]interface{})
	bindResourceRawObj, _ := jsonpath.Get("$.policy.bindResource", objectRaw)
	bindResourceRaw := make(map[string]interface{})
	if bindResourceRawObj != nil {
		bindResourceRaw = bindResourceRawObj.(map[string]interface{})
	}
	if len(bindResourceRaw) > 0 {
		entityGroupMap["cluster_entity_type"] = bindResourceRaw["clusterType"]
		entityGroupMap["cluster_id"] = bindResourceRaw["clusterId"]
		entityGroupMap["vpc_id"] = bindResourceRaw["vpcId"]

		entityGroupRawObj, _ := jsonpath.Get("$.policy.entityGroup", objectRaw)
		entityGroupRaw := make(map[string]interface{})
		if entityGroupRawObj != nil {
			entityGroupRaw = entityGroupRawObj.(map[string]interface{})
		}
		if len(entityGroupRaw) > 0 {
			entityGroupMap["entity_group_id"] = entityGroupRaw["entityGroupId"]
		}
		entityGroupMaps = append(entityGroupMaps, entityGroupMap)
	}
	if err := d.Set("entity_group", entityGroupMaps); err != nil {
		return err
	}
	tagsMaps, _ := jsonpath.Get("$.policy.tags", objectRaw)
	d.Set("tags", tagsToMap(tagsMaps))

	return nil
}

func resourceAliCloudCmsIntegrationPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	integrationPolicyId := d.Id()
	action := fmt.Sprintf("/integration-policies/%s", integrationPolicyId)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})

	if d.HasChange("tags") {
		update = true
	}
	if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request["Tags"] = tagsMap
	}

	if _, ok := d.GetOk("resource_group_id"); ok && d.HasChange("resource_group_id") {
		update = true
	}
	if v, ok := d.GetOk("resource_group_id"); ok || d.HasChange("resource_group_id") {
		request["resourceGroupId"] = v
	}
	if v, ok := d.GetOk("fee_package"); ok {
		request["feePackage"] = v
	}
	if d.HasChange("integration_policy_name") {
		update = true
	}
	if v, ok := d.GetOk("integration_policy_name"); ok || d.HasChange("integration_policy_name") {
		request["policyName"] = v
	}
	body = request
	if update {
		wait := incrementalWait(3*time.Second, 0*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPut("Cms", "2024-03-30", action, query, nil, body, true)
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

	return resourceAliCloudCmsIntegrationPolicyRead(d, meta)
}

func resourceAliCloudCmsIntegrationPolicyDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	policyId := d.Id()
	action := fmt.Sprintf("/integration-policies/%s", policyId)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	var err error
	request = make(map[string]interface{})

	if v, ok := d.GetOkExists("force"); ok {
		query["force"] = StringPointer(strconv.FormatBool(v.(bool)))
	}

	wait := incrementalWait(3*time.Second, 0*time.Second)
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
		if IsExpectedErrors(err, []string{"404", "15007"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
