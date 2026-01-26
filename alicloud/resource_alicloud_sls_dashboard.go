// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudSlsDashboard() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudSlsDashboardCreate,
		Read:   resourceAliCloudSlsDashboardRead,
		Update: resourceAliCloudSlsDashboardUpdate,
		Delete: resourceAliCloudSlsDashboardDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"attribute": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"charts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"action": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"search": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"title": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"display": {
							Type:     schema.TypeMap,
							Optional: true,
						},
					},
				},
			},
			"dashboard_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudSlsDashboardCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/dashboards")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	hostMap := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	hostMap["project"] = StringPointer(d.Get("project_name").(string))
	if v, ok := d.GetOk("dashboard_name"); ok {
		request["dashboardName"] = v
	}

	request["displayName"] = d.Get("display_name")
	if v, ok := d.GetOk("description"); ok {
		request["description"] = v
	}
	if v, ok := d.GetOk("attribute"); ok {
		attributeMapsArray := convertToInterfaceArray(v)

		request["attribute"] = attributeMapsArray
	}

	if v, ok := d.GetOk("charts"); ok {
		chartsMapsArray := make([]interface{}, 0)
		for _, dataLoop1 := range convertToInterfaceArray(v) {
			dataLoop1Tmp := dataLoop1.(map[string]interface{})
			dataLoop1Map := make(map[string]interface{})
			dataLoop1Map["title"] = dataLoop1Tmp["title"]
			dataLoop1Map["type"] = dataLoop1Tmp["type"]
			dataLoop1Map["search"] = dataLoop1Tmp["search"]
			dataLoop1Map["display"] = dataLoop1Tmp["display"]
			dataLoop1Map["action"] = dataLoop1Tmp["action"]
			chartsMapsArray = append(chartsMapsArray, dataLoop1Map)
		}
		request["charts"] = chartsMapsArray
	}

	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.Do("Sls", roaParam("POST", "2020-12-30", "CreateDashboard", action), query, body, nil, hostMap, false)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_sls_dashboard", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", *hostMap["project"], request["dashboardName"]))

	return resourceAliCloudSlsDashboardUpdate(d, meta)
}

func resourceAliCloudSlsDashboardRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	slsServiceV2 := SlsServiceV2{client}

	objectRaw, err := slsServiceV2.DescribeSlsDashboard(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_sls_dashboard DescribeSlsDashboard Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("description", objectRaw["description"])
	d.Set("display_name", objectRaw["displayName"])
	d.Set("dashboard_name", objectRaw["dashboardName"])

	chartsRaw := objectRaw["charts"]
	chartsMaps := make([]map[string]interface{}, 0)
	if chartsRaw != nil {
		for _, chartsChildRaw := range convertToInterfaceArray(chartsRaw) {
			chartsMap := make(map[string]interface{})
			chartsChildRaw := chartsChildRaw.(map[string]interface{})
			chartsMap["action"] = chartsChildRaw["action"]
			chartsMap["display"] = chartsChildRaw["display"]
			chartsMap["search"] = chartsChildRaw["search"]
			chartsMap["title"] = chartsChildRaw["title"]
			chartsMap["type"] = chartsChildRaw["type"]

			chartsMaps = append(chartsMaps, chartsMap)
		}
	}
	if err := d.Set("charts", chartsMaps); err != nil {
		return err
	}

	parts := strings.Split(d.Id(), ":")
	d.Set("project_name", parts[0])

	return nil
}

func resourceAliCloudSlsDashboardUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var header map[string]*string
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	dashboardName := parts[1]
	action := fmt.Sprintf("/dashboards/%s", dashboardName)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	hostMap := make(map[string]*string)
	hostMap["project"] = StringPointer(parts[0])

	if !d.IsNewResource() && d.HasChange("display_name") {
		update = true
	}
	request["displayName"] = d.Get("display_name")
	if !d.IsNewResource() && d.HasChange("description") {
		update = true
	}
	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		request["description"] = v
	}
	if !d.IsNewResource() && d.HasChange("attribute") {
		update = true
	}
	if v, ok := d.GetOk("attribute"); ok || d.HasChange("attribute") {
		attributeMapsArray := convertToInterfaceArray(v)

		request["attribute"] = attributeMapsArray
	}

	if !d.IsNewResource() && d.HasChange("charts") {
		update = true
	}
	if v, ok := d.GetOk("charts"); ok || d.HasChange("charts") {
		chartsMapsArray := make([]interface{}, 0)
		for _, dataLoop1 := range convertToInterfaceArray(v) {
			dataLoop1Tmp := dataLoop1.(map[string]interface{})
			dataLoop1Map := make(map[string]interface{})
			dataLoop1Map["title"] = dataLoop1Tmp["title"]
			dataLoop1Map["type"] = dataLoop1Tmp["type"]
			dataLoop1Map["search"] = dataLoop1Tmp["search"]
			dataLoop1Map["display"] = dataLoop1Tmp["display"]
			dataLoop1Map["action"] = dataLoop1Tmp["action"]
			chartsMapsArray = append(chartsMapsArray, dataLoop1Map)
		}
		request["charts"] = chartsMapsArray
	}

	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.Do("Sls", roaParam("PUT", "2020-12-30", "UpdateDashboard", action), query, body, nil, hostMap, false)
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

	return resourceAliCloudSlsDashboardRead(d, meta)
}

func resourceAliCloudSlsDashboardDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	dashboardName := parts[1]
	action := fmt.Sprintf("/dashboards/%s", dashboardName)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	hostMap := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	hostMap["project"] = StringPointer(parts[0])

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.Do("Sls", roaParam("DELETE", "2020-12-30", "DeleteDashboard", action), query, nil, nil, hostMap, false)
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
