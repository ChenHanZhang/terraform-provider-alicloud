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

func resourceAliCloudSlsCollectionPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudSlsCollectionPolicyCreate,
		Read:   resourceAliCloudSlsCollectionPolicyRead,
		Update: resourceAliCloudSlsCollectionPolicyUpdate,
		Delete: resourceAliCloudSlsCollectionPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"centralize_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dest_ttl": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"dest_logstore": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"dest_region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"dest_project": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"centralize_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"data_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_region": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"data_project": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"policy_config": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_tags": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"regions": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"instance_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"resource_mode": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: StringInSlice([]string{"all", "instanceMode", "attributeMode"}, false),
						},
					},
				},
			},
			"policy_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"product_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_directory": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_group_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"members": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func resourceAliCloudSlsCollectionPolicyCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/collectionpolicy")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	hostMap := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("policy_name"); ok {
		request["policyName"] = v
	}

	request["productCode"] = d.Get("product_code")
	resourceDirectory := make(map[string]interface{})

	if v := d.Get("resource_directory"); !IsNil(v) {
		accountGroupType1, _ := jsonpath.Get("$[0].account_group_type", v)
		if accountGroupType1 != nil && accountGroupType1 != "" {
			resourceDirectory["accountGroupType"] = accountGroupType1
		}
		members1, _ := jsonpath.Get("$[0].members", v)
		if members1 != nil && members1 != "" {
			resourceDirectory["members"] = members1
		}

		request["resourceDirectory"] = resourceDirectory
	}

	centralizeConfig := make(map[string]interface{})

	if v := d.Get("centralize_config"); !IsNil(v) {
		destTtl, _ := jsonpath.Get("$[0].dest_ttl", v)
		if destTtl != nil && destTtl != "" {
			centralizeConfig["destTTL"] = destTtl
		}
		destRegion1, _ := jsonpath.Get("$[0].dest_region", v)
		if destRegion1 != nil && destRegion1 != "" {
			centralizeConfig["destRegion"] = destRegion1
		}
		destLogstore1, _ := jsonpath.Get("$[0].dest_logstore", v)
		if destLogstore1 != nil && destLogstore1 != "" {
			centralizeConfig["destLogstore"] = destLogstore1
		}
		destProject1, _ := jsonpath.Get("$[0].dest_project", v)
		if destProject1 != nil && destProject1 != "" {
			centralizeConfig["destProject"] = destProject1
		}

		request["centralizeConfig"] = centralizeConfig
	}

	policyConfig := make(map[string]interface{})

	if v := d.Get("policy_config"); v != nil {
		instanceIds1, _ := jsonpath.Get("$[0].instance_ids", v)
		if instanceIds1 != nil && instanceIds1 != "" {
			policyConfig["instanceIds"] = instanceIds1
		}
		regions1, _ := jsonpath.Get("$[0].regions", v)
		if regions1 != nil && regions1 != "" {
			policyConfig["regions"] = regions1
		}
		resourceTags1, _ := jsonpath.Get("$[0].resource_tags", v)
		if resourceTags1 != nil && resourceTags1 != "" {
			policyConfig["resourceTags"] = resourceTags1
		}
		resourceMode1, _ := jsonpath.Get("$[0].resource_mode", v)
		if resourceMode1 != nil && resourceMode1 != "" {
			policyConfig["resourceMode"] = resourceMode1
		}

		request["policyConfig"] = policyConfig
	}

	request["enabled"] = d.Get("enabled")
	if v, ok := d.GetOkExists("centralize_enabled"); ok {
		request["centralizeEnabled"] = v
	}
	dataConfig := make(map[string]interface{})

	if v := d.Get("data_config"); !IsNil(v) {
		dataRegion1, _ := jsonpath.Get("$[0].data_region", v)
		if dataRegion1 != nil && dataRegion1 != "" {
			dataConfig["dataRegion"] = dataRegion1
		}

		request["dataConfig"] = dataConfig
	}

	request["dataCode"] = d.Get("data_code")
	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.Do("Sls", roaParam("POST", "2020-12-30", "UpsertCollectionPolicy", action), query, body, nil, hostMap, false)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_sls_collection_policy", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["policyName"]))

	slsServiceV2 := SlsServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"#CHECKSET"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, slsServiceV2.SlsCollectionPolicyStateRefreshFunc(d.Id(), "#policyName", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudSlsCollectionPolicyRead(d, meta)
}

func resourceAliCloudSlsCollectionPolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	slsServiceV2 := SlsServiceV2{client}

	objectRaw, err := slsServiceV2.DescribeSlsCollectionPolicy(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_sls_collection_policy DescribeSlsCollectionPolicy Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("centralize_enabled", objectRaw["centralizeEnabled"])
	d.Set("data_code", objectRaw["dataCode"])
	d.Set("enabled", objectRaw["enabled"])
	d.Set("product_code", objectRaw["productCode"])
	d.Set("policy_name", objectRaw["policyName"])

	centralizeConfigMaps := make([]map[string]interface{}, 0)
	centralizeConfigMap := make(map[string]interface{})
	centralizeConfigRaw := make(map[string]interface{})
	if objectRaw["centralizeConfig"] != nil {
		centralizeConfigRaw = objectRaw["centralizeConfig"].(map[string]interface{})
	}
	if len(centralizeConfigRaw) > 0 {
		centralizeConfigMap["dest_logstore"] = centralizeConfigRaw["destLogstore"]
		centralizeConfigMap["dest_project"] = centralizeConfigRaw["destProject"]
		centralizeConfigMap["dest_region"] = centralizeConfigRaw["destRegion"]
		centralizeConfigMap["dest_ttl"] = centralizeConfigRaw["destTTL"]

		centralizeConfigMaps = append(centralizeConfigMaps, centralizeConfigMap)
	}
	if err := d.Set("centralize_config", centralizeConfigMaps); err != nil {
		return err
	}
	dataConfigMaps := make([]map[string]interface{}, 0)
	dataConfigMap := make(map[string]interface{})
	dataConfigRaw := make(map[string]interface{})
	if objectRaw["dataConfig"] != nil {
		dataConfigRaw = objectRaw["dataConfig"].(map[string]interface{})
	}
	if len(dataConfigRaw) > 0 {
		dataConfigMap["data_project"] = dataConfigRaw["dataProject"]
		dataConfigMap["data_region"] = dataConfigRaw["dataRegion"]

		dataConfigMaps = append(dataConfigMaps, dataConfigMap)
	}
	if err := d.Set("data_config", dataConfigMaps); err != nil {
		return err
	}
	policyConfigMaps := make([]map[string]interface{}, 0)
	policyConfigMap := make(map[string]interface{})
	policyConfigRaw := make(map[string]interface{})
	if objectRaw["policyConfig"] != nil {
		policyConfigRaw = objectRaw["policyConfig"].(map[string]interface{})
	}
	if len(policyConfigRaw) > 0 {
		policyConfigMap["resource_mode"] = policyConfigRaw["resourceMode"]
		policyConfigMap["resource_tags"] = policyConfigRaw["resourceTags"]

		instanceIdsRaw := make([]interface{}, 0)
		if policyConfigRaw["instanceIds"] != nil {
			instanceIdsRaw = convertToInterfaceArray(policyConfigRaw["instanceIds"])
		}

		policyConfigMap["instance_ids"] = instanceIdsRaw
		regionsRaw := make([]interface{}, 0)
		if policyConfigRaw["regions"] != nil {
			regionsRaw = convertToInterfaceArray(policyConfigRaw["regions"])
		}

		policyConfigMap["regions"] = regionsRaw
		policyConfigMaps = append(policyConfigMaps, policyConfigMap)
	}
	if err := d.Set("policy_config", policyConfigMaps); err != nil {
		return err
	}
	resourceDirectoryMaps := make([]map[string]interface{}, 0)
	resourceDirectoryMap := make(map[string]interface{})
	resourceDirectoryRaw := make(map[string]interface{})
	if objectRaw["resourceDirectory"] != nil && objectRaw["resourceDirectory"].(map[string]interface{})["accountGroupType"] != "" {
		resourceDirectoryRaw = objectRaw["resourceDirectory"].(map[string]interface{})
	}
	if len(resourceDirectoryRaw) > 0 {
		resourceDirectoryMap["account_group_type"] = resourceDirectoryRaw["accountGroupType"]

		membersRaw := make([]interface{}, 0)
		if resourceDirectoryRaw["members"] != nil {
			membersRaw = convertToInterfaceArray(resourceDirectoryRaw["members"])
		}

		resourceDirectoryMap["members"] = membersRaw
		resourceDirectoryMaps = append(resourceDirectoryMaps, resourceDirectoryMap)
	}
	if err := d.Set("resource_directory", resourceDirectoryMaps); err != nil {
		return err
	}

	d.Set("policy_name", d.Id())

	return nil
}

func resourceAliCloudSlsCollectionPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	action := fmt.Sprintf("/collectionpolicy")
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	hostMap := make(map[string]*string)
	request["policyName"] = d.Id()

	if d.HasChange("product_code") {
		update = true
	}
	request["productCode"] = d.Get("product_code")
	if d.HasChange("resource_directory") {
		update = true
	}
	resourceDirectory := make(map[string]interface{})

	if v := d.Get("resource_directory"); !IsNil(v) || d.HasChange("resource_directory") {
		accountGroupType1, _ := jsonpath.Get("$[0].account_group_type", v)
		if accountGroupType1 != nil && accountGroupType1 != "" {
			resourceDirectory["accountGroupType"] = accountGroupType1
		}
		members1, _ := jsonpath.Get("$[0].members", v)
		if members1 != nil && members1 != "" {
			resourceDirectory["members"] = members1
		}

		request["resourceDirectory"] = resourceDirectory
	}

	if d.HasChange("centralize_config") {
		update = true
	}
	centralizeConfig := make(map[string]interface{})

	if v := d.Get("centralize_config"); !IsNil(v) || d.HasChange("centralize_config") {
		destTtl, _ := jsonpath.Get("$[0].dest_ttl", v)
		if destTtl != nil && destTtl != "" {
			centralizeConfig["destTTL"] = destTtl
		}
		destRegion1, _ := jsonpath.Get("$[0].dest_region", v)
		if destRegion1 != nil && destRegion1 != "" {
			centralizeConfig["destRegion"] = destRegion1
		}
		destLogstore1, _ := jsonpath.Get("$[0].dest_logstore", v)
		if destLogstore1 != nil && destLogstore1 != "" {
			centralizeConfig["destLogstore"] = destLogstore1
		}
		destProject1, _ := jsonpath.Get("$[0].dest_project", v)
		if destProject1 != nil && destProject1 != "" {
			centralizeConfig["destProject"] = destProject1
		}

		request["centralizeConfig"] = centralizeConfig
	}

	if d.HasChange("policy_config") {
		update = true
	}
	policyConfig := make(map[string]interface{})

	if v := d.Get("policy_config"); v != nil {
		instanceIds1, _ := jsonpath.Get("$[0].instance_ids", v)
		if instanceIds1 != nil && instanceIds1 != "" {
			policyConfig["instanceIds"] = instanceIds1
		}
		regions1, _ := jsonpath.Get("$[0].regions", v)
		if regions1 != nil && regions1 != "" {
			policyConfig["regions"] = regions1
		}
		resourceTags1, _ := jsonpath.Get("$[0].resource_tags", v)
		if resourceTags1 != nil && resourceTags1 != "" {
			policyConfig["resourceTags"] = resourceTags1
		}
		resourceMode1, _ := jsonpath.Get("$[0].resource_mode", v)
		if resourceMode1 != nil && resourceMode1 != "" {
			policyConfig["resourceMode"] = resourceMode1
		}

		request["policyConfig"] = policyConfig
	}

	if d.HasChange("enabled") {
		update = true
	}
	request["enabled"] = d.Get("enabled")
	if d.HasChange("centralize_enabled") {
		update = true
	}
	if v, ok := d.GetOkExists("centralize_enabled"); ok || d.HasChange("centralize_enabled") {
		request["centralizeEnabled"] = v
	}
	if d.HasChange("data_config") {
		update = true
	}
	dataConfig := make(map[string]interface{})

	if v := d.Get("data_config"); !IsNil(v) || d.HasChange("data_config") {
		dataRegion1, _ := jsonpath.Get("$[0].data_region", v)
		if dataRegion1 != nil && dataRegion1 != "" {
			dataConfig["dataRegion"] = dataRegion1
		}

		request["dataConfig"] = dataConfig
	}

	if d.HasChange("data_code") {
		update = true
	}
	request["dataCode"] = d.Get("data_code")
	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.Do("Sls", roaParam("POST", "2020-12-30", "UpsertCollectionPolicy", action), query, body, nil, hostMap, false)
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
		slsServiceV2 := SlsServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"#CHECKSET"}, d.Timeout(schema.TimeoutUpdate), 10*time.Second, slsServiceV2.SlsCollectionPolicyStateRefreshFunc(d.Id(), "#policyName", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	return resourceAliCloudSlsCollectionPolicyRead(d, meta)
}

func resourceAliCloudSlsCollectionPolicyDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	policyName := d.Id()
	action := fmt.Sprintf("/collectionpolicy/%s", policyName)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	hostMap := make(map[string]*string)
	var err error
	request = make(map[string]interface{})

	if v, ok := d.GetOk("product_code"); ok {
		query["productCode"] = StringPointer(v.(string))
	}

	if v, ok := d.GetOk("data_code"); ok {
		query["dataCode"] = StringPointer(v.(string))
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.Do("Sls", roaParam("DELETE", "2020-12-30", "DeleteCollectionPolicy", action), query, nil, nil, hostMap, false)
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
		if IsExpectedErrors(err, []string{"PolicyNotExist"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	slsServiceV2 := SlsServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 5*time.Second, slsServiceV2.SlsCollectionPolicyStateRefreshFunc(d.Id(), "policyName", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
