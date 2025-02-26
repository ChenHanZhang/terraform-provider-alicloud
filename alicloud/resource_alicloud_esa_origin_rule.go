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

func resourceAliCloudEsaOriginRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudEsaOriginRuleCreate,
		Read:   resourceAliCloudEsaOriginRuleRead,
		Update: resourceAliCloudEsaOriginRuleUpdate,
		Delete: resourceAliCloudEsaOriginRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"config_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dns_record": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"origin_host": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"origin_http_port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"origin_https_port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"origin_scheme": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"http", "https", "follow"}, false),
			},
			"origin_sni": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"range": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"on", "off", "force"}, false),
			},
			"rule": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_enable": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"on", "off"}, false),
			},
			"rule_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"site_version": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudEsaOriginRuleCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateOriginRule"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("site_id"); ok {
		request["SiteId"] = v
	}
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOkExists("site_version"); ok {
		request["SiteVersion"] = v
	}
	if v, ok := d.GetOk("range"); ok {
		request["Range"] = v
	}
	if v, ok := d.GetOk("origin_https_port"); ok {
		request["OriginHttpsPort"] = v
	}
	if v, ok := d.GetOk("rule_enable"); ok {
		request["RuleEnable"] = v
	}
	if v, ok := d.GetOk("rule_name"); ok {
		request["RuleName"] = v
	}
	if v, ok := d.GetOk("origin_host"); ok {
		request["OriginHost"] = v
	}
	if v, ok := d.GetOk("origin_scheme"); ok {
		request["OriginScheme"] = v
	}
	if v, ok := d.GetOk("origin_sni"); ok {
		request["OriginSni"] = v
	}
	if v, ok := d.GetOk("dns_record"); ok {
		request["DnsRecord"] = v
	}
	if v, ok := d.GetOk("origin_http_port"); ok {
		request["OriginHttpPort"] = v
	}
	if v, ok := d.GetOk("rule"); ok {
		request["Rule"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("ESA", "2024-09-10", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_esa_origin_rule", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["SiteId"], response["ConfigId"]))

	return resourceAliCloudEsaOriginRuleRead(d, meta)
}

func resourceAliCloudEsaOriginRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	esaServiceV2 := EsaServiceV2{client}

	objectRaw, err := esaServiceV2.DescribeEsaOriginRule(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_esa_origin_rule DescribeEsaOriginRule Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("dns_record", objectRaw["DnsRecord"])
	d.Set("origin_host", objectRaw["OriginHost"])
	d.Set("origin_http_port", objectRaw["OriginHttpPort"])
	d.Set("origin_https_port", objectRaw["OriginHttpsPort"])
	d.Set("origin_scheme", objectRaw["OriginScheme"])
	d.Set("origin_sni", objectRaw["OriginSni"])
	d.Set("range", objectRaw["Range"])
	d.Set("rule", objectRaw["Rule"])
	d.Set("rule_enable", objectRaw["RuleEnable"])
	d.Set("rule_name", objectRaw["RuleName"])
	d.Set("site_version", objectRaw["SiteVersion"])
	d.Set("config_id", objectRaw["ConfigId"])

	parts := strings.Split(d.Id(), ":")
	d.Set("site_id", formatInt(parts[0]))

	return nil
}

func resourceAliCloudEsaOriginRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "UpdateOriginRule"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ConfigId"] = parts[1]
	request["SiteId"] = parts[0]
	request["RegionId"] = client.RegionId
	if d.HasChange("origin_sni") {
		update = true
		request["OriginSni"] = d.Get("origin_sni")
	}

	if d.HasChange("dns_record") {
		update = true
		request["DnsRecord"] = d.Get("dns_record")
	}

	if d.HasChange("range") {
		update = true
		request["Range"] = d.Get("range")
	}

	if d.HasChange("origin_https_port") {
		update = true
		request["OriginHttpsPort"] = d.Get("origin_https_port")
	}

	if d.HasChange("origin_http_port") {
		update = true
		request["OriginHttpPort"] = d.Get("origin_http_port")
	}

	if d.HasChange("rule_enable") {
		update = true
		request["RuleEnable"] = d.Get("rule_enable")
	}

	if d.HasChange("rule_name") {
		update = true
		request["RuleName"] = d.Get("rule_name")
	}

	if d.HasChange("origin_host") {
		update = true
		request["OriginHost"] = d.Get("origin_host")
	}

	if d.HasChange("origin_scheme") {
		update = true
		request["OriginScheme"] = d.Get("origin_scheme")
	}

	if d.HasChange("rule") {
		update = true
		request["Rule"] = d.Get("rule")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("ESA", "2024-09-10", action, query, request, true)
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

	return resourceAliCloudEsaOriginRuleRead(d, meta)
}

func resourceAliCloudEsaOriginRuleDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteOriginRule"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["ConfigId"] = parts[1]
	request["SiteId"] = parts[0]
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("ESA", "2024-09-10", action, query, request, true)

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
