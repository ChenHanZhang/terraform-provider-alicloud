// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudWebsiteBuildAppInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudWebsiteBuildAppInstanceCreate,
		Read:   resourceAliCloudWebsiteBuildAppInstanceRead,
		Update: resourceAliCloudWebsiteBuildAppInstanceUpdate,
		Delete: resourceAliCloudWebsiteBuildAppInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"auto_renew": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"duration": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: StringMatch(regexp.MustCompile("^\\d+$"), "Number of subscription periods.  "),
			},
			"extend": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"payment_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pricing_cycle": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"年", "月", "天"}, false),
			},
			"profile": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_etag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"deploy_area": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"order_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_version": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"template_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"quantity": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: StringMatch(regexp.MustCompile("^\\d+$"), "Quantity  "),
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchema(),
		},
	}
}

func resourceAliCloudWebsiteBuildAppInstanceCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateAppInstance"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["ClientToken"] = buildClientToken(action)

	request["Duration"] = d.Get("duration")
	if v, ok := d.GetOkExists("auto_renew"); ok {
		request["AutoRenew"] = v
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if v, ok := d.GetOk("payment_type"); ok {
		request["PaymentType"] = v
	}
	if v, ok := d.GetOk("profile"); ok {
		profileSiteVersionJsonPath, err := jsonpath.Get("$[0].site_version", v)
		if err == nil && profileSiteVersionJsonPath != "" {
			request["SiteVersion"] = profileSiteVersionJsonPath
		}
	}
	tagsDataList := make(map[string]interface{})

	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		tagsDataList["TagValue"] = tagsMap
	}

	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		tagsDataList["TagKey"] = tagsMap
	}

	TagsMap := make([]interface{}, 0)
	TagsMap = append(TagsMap, tagsDataList)
	request["Tags"] = TagsMap

	if v, ok := d.GetOk("profile"); ok {
		profileApplicationTypeJsonPath, err := jsonpath.Get("$[0].application_type", v)
		if err == nil && profileApplicationTypeJsonPath != "" {
			request["ApplicationType"] = profileApplicationTypeJsonPath
		}
	}
	if v, ok := d.GetOk("pricing_cycle"); ok {
		request["PricingCycle"] = v
	}
	if v, ok := d.GetOk("profile"); ok {
		profileDeployAreaJsonPath, err := jsonpath.Get("$[0].deploy_area", v)
		if err == nil && profileDeployAreaJsonPath != "" {
			request["DeployArea"] = profileDeployAreaJsonPath
		}
	}
	if v, ok := d.GetOkExists("quantity"); ok {
		request["Quantity"] = v
	}
	if v, ok := d.GetOk("extend"); ok {
		request["Extend"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("WebsiteBuild", "2025-04-29", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_website_build_app_instance", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.Module.BizId", response)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudWebsiteBuildAppInstanceRead(d, meta)
}

func resourceAliCloudWebsiteBuildAppInstanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	websiteBuildServiceV2 := WebsiteBuildServiceV2{client}

	objectRaw, err := websiteBuildServiceV2.DescribeWebsiteBuildAppInstance(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_website_build_app_instance DescribeWebsiteBuildAppInstance Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("resource_group_id", objectRaw["ResourceGroupId"])
	d.Set("status", objectRaw["Status"])

	profileMaps := make([]map[string]interface{}, 0)
	profileMap := make(map[string]interface{})
	profileRaw := make(map[string]interface{})
	if objectRaw["Profile"] != nil {
		profileRaw = objectRaw["Profile"].(map[string]interface{})
	}
	if len(profileRaw) > 0 {
		profileMap["application_type"] = profileRaw["ApplicationType"]
		profileMap["deploy_area"] = profileRaw["DeployArea"]
		profileMap["instance_id"] = profileRaw["InstanceId"]
		profileMap["order_id"] = profileRaw["OrderId"]
		profileMap["site_version"] = profileRaw["SiteVersion"]
		profileMap["template_etag"] = profileRaw["TemplateEtag"]
		profileMap["template_id"] = profileRaw["TemplateId"]

		profileMaps = append(profileMaps, profileMap)
	}
	if err := d.Set("profile", profileMaps); err != nil {
		return err
	}
	tagsMaps := tagsChildRaw.(map[string]interface{})
	d.Set("tags", tagsToMap(tagsMaps))

	return nil
}

func resourceAliCloudWebsiteBuildAppInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "UpdateAppInstance"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["BizId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if _, ok := d.GetOk("resource_group_id"); ok && d.HasChange("resource_group_id") {
		update = true
		request["ResourceGroupId"] = d.Get("resource_group_id")
	}

	if v, ok := d.GetOk("payment_type"); ok {
		request["PaymentType"] = v
	}
	if d.HasChange("profile.0.site_version") {
		update = true
		profileSiteVersionJsonPath, err := jsonpath.Get("$[0].site_version", d.Get("profile"))
		if err == nil {
			request["SiteVersion"] = profileSiteVersionJsonPath
		}
	}

	tagsDataList := make(map[string]interface{})

	if d.HasChange("tags") {
		update = true
		tagsMap := ConvertTags(v.(map[string]interface{}))
		tagsDataList["TagValue"] = tagsMap
	}

	if d.HasChange("tags") {
		update = true
		tagsMap := ConvertTags(v.(map[string]interface{}))
		tagsDataList["TagKey"] = tagsMap
	}

	TagsMap := make([]interface{}, 0)
	TagsMap = append(TagsMap, tagsDataList)
	request["Tags"] = TagsMap

	if d.HasChange("profile.0.application_type") {
		update = true
		profileApplicationTypeJsonPath, err := jsonpath.Get("$[0].application_type", d.Get("profile"))
		if err == nil {
			request["ApplicationType"] = profileApplicationTypeJsonPath
		}
	}

	if v, ok := d.GetOk("extend"); ok {
		request["Extend"] = v
	}
	if d.HasChange("profile.0.deploy_area") {
		update = true
		profileDeployAreaJsonPath, err := jsonpath.Get("$[0].deploy_area", d.Get("profile"))
		if err == nil {
			request["DeployArea"] = profileDeployAreaJsonPath
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("WebsiteBuild", "2025-04-29", action, query, request, true)
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

	return resourceAliCloudWebsiteBuildAppInstanceRead(d, meta)
}

func resourceAliCloudWebsiteBuildAppInstanceDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteAppInstance"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["BizId"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("WebsiteBuild", "2025-04-29", action, query, request, true)
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
