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

func resourceAliCloudCloudMonitorServiceDynamicTagGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCloudMonitorServiceDynamicTagGroupCreate,
		Read:   resourceAliCloudCloudMonitorServiceDynamicTagGroupRead,
		Update: resourceAliCloudCloudMonitorServiceDynamicTagGroupUpdate,
		Delete: resourceAliCloudCloudMonitorServiceDynamicTagGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"contact_group_list": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"enable_install_agent": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_subscribe_event": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"match_express": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tag_name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"tag_value": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"tag_value_match_function": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"match_express_filter_relation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tag_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tag_region_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"template_id_list": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceAliCloudCloudMonitorServiceDynamicTagGroupCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateDynamicTagGroup"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	if v, ok := d.GetOk("match_express"); ok {
		matchExpressMapsArray := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(v) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["TagValueMatchFunction"] = dataLoopTmp["tag_value_match_function"]
			dataLoopMap["TagValue"] = dataLoopTmp["tag_value"]
			dataLoopMap["TagName"] = dataLoopTmp["tag_name"]
			matchExpressMapsArray = append(matchExpressMapsArray, dataLoopMap)
		}
		request["MatchExpress"] = matchExpressMapsArray
	}

	if v, ok := d.GetOk("template_id_list"); ok {
		templateIdListMapsArray := convertToInterfaceArray(v)

		request["TemplateIdList"] = templateIdListMapsArray
	}

	request["TagKey"] = d.Get("tag_key")
	if v, ok := d.GetOk("match_express_filter_relation"); ok {
		request["MatchExpressFilterRelation"] = v
	}
	if v, ok := d.GetOk("contact_group_list"); ok {
		contactGroupListMapsArray := convertToInterfaceArray(v)

		request["ContactGroupList"] = contactGroupListMapsArray
	}

	if v, ok := d.GetOkExists("enable_subscribe_event"); ok {
		request["EnableSubscribeEvent"] = v
	}
	if v, ok := d.GetOkExists("enable_install_agent"); ok {
		request["EnableInstallAgent"] = v
	}
	if v, ok := d.GetOk("tag_region_id"); ok {
		request["TagRegionId"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Cms", "2019-01-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cms_dynamic_tag_group", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["Id"]))

	cloudMonitorServiceServiceV2 := CloudMonitorServiceServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"FINISH"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, cloudMonitorServiceServiceV2.CloudMonitorServiceDynamicTagGroupStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudCloudMonitorServiceDynamicTagGroupRead(d, meta)
}

func resourceAliCloudCloudMonitorServiceDynamicTagGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cloudMonitorServiceServiceV2 := CloudMonitorServiceServiceV2{client}

	objectRaw, err := cloudMonitorServiceServiceV2.DescribeCloudMonitorServiceDynamicTagGroup(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cms_dynamic_tag_group DescribeCloudMonitorServiceDynamicTagGroup Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("match_express_filter_relation", objectRaw["MatchExpressFilterRelation"])
	d.Set("status", objectRaw["Status"])
	d.Set("tag_key", objectRaw["TagKey"])
	d.Set("tag_region_id", objectRaw["RegionId"])

	contactGroupListRaw, _ := jsonpath.Get("$.ContactGroupList.ContactGroupList", objectRaw)
	d.Set("contact_group_list", contactGroupListRaw)
	matchExpressRaw, _ := jsonpath.Get("$.MatchExpress.MatchExpress", objectRaw)
	matchExpressMaps := make([]map[string]interface{}, 0)
	if matchExpressRaw != nil {
		for _, matchExpressChildRaw := range convertToInterfaceArray(matchExpressRaw) {
			matchExpressMap := make(map[string]interface{})
			matchExpressChildRaw := matchExpressChildRaw.(map[string]interface{})
			matchExpressMap["tag_name"] = matchExpressChildRaw["TagKey"]
			matchExpressMap["tag_value"] = matchExpressChildRaw["TagValue"]
			matchExpressMap["tag_value_match_function"] = matchExpressChildRaw["TagValueMatchFunction"]

			matchExpressMaps = append(matchExpressMaps, matchExpressMap)
		}
	}
	if err := d.Set("match_express", matchExpressMaps); err != nil {
		return err
	}
	templateIdListRaw, _ := jsonpath.Get("$.TemplateIdList.TemplateIdList", objectRaw)
	d.Set("template_id_list", templateIdListRaw)

	return nil
}

func resourceAliCloudCloudMonitorServiceDynamicTagGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Cannot update resource Alicloud Resource Dynamic Tag Group.")
	return nil
}

func resourceAliCloudCloudMonitorServiceDynamicTagGroupDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteDynamicTagGroup"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["DynamicTagRuleId"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Cms", "2019-01-01", action, query, request, true)
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
