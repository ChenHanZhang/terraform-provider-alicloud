package alicloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudCdnDomainConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCdnDomainConfigCreate,
		Read:   resourceAliCloudCdnDomainConfigRead,
		Update: resourceAliCloudCdnDomainConfigUpdate,
		Delete: resourceAliCloudCdnDomainConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(15 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"domain_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(5, 67),
			},
			"function_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"function_args": {
				Type:     schema.TypeSet,
				Set:      expirationCdnDomainConfigHash,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"arg_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"arg_value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"config_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudCdnDomainConfigCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cdnService := &CdnService{client: client}
	var response map[string]interface{}
	var err error
	action := "BatchSetCdnDomainConfig"

	config := make([]map[string]interface{}, 1)
	functionArgs := d.Get("function_args").(*schema.Set).List()
	args := make([]map[string]interface{}, len(functionArgs))
	for key, value := range functionArgs {
		arg := value.(map[string]interface{})
		args[key] = map[string]interface{}{
			"argName":  arg["arg_name"],
			"argValue": arg["arg_value"],
		}
	}
	config[0] = map[string]interface{}{
		"functionArgs": args,
		"functionName": d.Get("function_name").(string),
	}
	if v, ok := d.GetOk("parent_id"); ok {
		config[0]["parentId"] = v.(string)
	}
	bytconfig, _ := json.Marshal(config)

	request := map[string]interface{}{
		"RegionId":    client.RegionId,
		"DomainNames": d.Get("domain_name").(string),
		"Functions":   string(bytconfig),
	}

	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(client.GetRetryTimeout(d.Timeout(schema.TimeoutCreate)), func() *resource.RetryError {
		response, err = client.RpcPost("Cdn", "2018-05-10", action, nil, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"ServiceBusy"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cdn_domain_config", action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.DomainConfigList.DomainConfigModel", response)
	if err != nil {
		return WrapErrorf(err, FailedGetAttributeMsg, "request.DomainNames", "$.DomainConfigList.DomainConfigModel", err)
	}

	if len(v.([]interface{})) != 1 {
		return WrapErrorf(NotFoundErr("CDN", d.Get("function_name").(string)), DefaultErrorMsg, err)
	}

	val := v.([]interface{})[0].(map[string]interface{})
	d.SetId(fmt.Sprintf("%s:%s:%s", val["DomainName"], val["FunctionName"], val["ConfigId"]))

	stateConf := BuildStateConf([]string{}, []string{"success"}, d.Timeout(schema.TimeoutCreate), 30*time.Second, cdnService.CdnDomainConfigRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	err = cdnService.WaitForCdnDomain(d.Get("domain_name").(string), Online, DefaultTimeoutMedium)
	if err != nil {
		return WrapError(err)
	}

	return resourceAliCloudCdnDomainConfigRead(d, meta)
}

func resourceAliCloudCdnDomainConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cdnService := &CdnService{client: client}
	var response map[string]interface{}
	var err error

	parts, err := ParseResourceId(d.Id(), 3)
	if err != nil {
		return WrapError(err)
	}

	request := map[string]interface{}{
		"RegionId":    client.RegionId,
		"DomainNames": parts[0],
	}

	update := false
	if d.HasChange("function_args") || d.HasChange("parent_id") {
		update = true
		config := make([]map[string]interface{}, 1)
		functionArgs := d.Get("function_args").(*schema.Set).List()
		args := make([]map[string]interface{}, len(functionArgs))
		for key, value := range functionArgs {
			arg := value.(map[string]interface{})
			args[key] = map[string]interface{}{
				"argName":  arg["arg_name"],
				"argValue": arg["arg_value"],
			}
		}

		config[0] = map[string]interface{}{
			"functionArgs": args,
			"functionName": parts[1],
			"configId":     parts[2],
		}

		if v, ok := d.GetOk("parent_id"); ok {
			config[0]["parentId"] = v
		}

		bytconfig, _ := json.Marshal(config)
		request["Functions"] = string(bytconfig)
	}

	if update {
		action := "BatchSetCdnDomainConfig"

		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(client.GetRetryTimeout(d.Timeout(schema.TimeoutUpdate)), func() *resource.RetryError {
			response, err = client.RpcPost("Cdn", "2018-05-10", action, nil, request, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"ServiceBusy"}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)

		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, "alicloud_cdn_domain_config", action, AlibabaCloudSdkGoERROR)
		}

		v, err := jsonpath.Get("$.DomainConfigList.DomainConfigModel", response)
		if err != nil {
			return WrapErrorf(err, FailedGetAttributeMsg, "request.DomainNames", "$.DomainConfigList.DomainConfigModel", err)
		}

		if len(v.([]interface{})) > 1 {
			return WrapErrorf(NotFoundErr("CDN", d.Get("function_name").(string)), DefaultErrorMsg, err)
		}

		stateConf := BuildStateConf([]string{}, []string{"success"}, d.Timeout(schema.TimeoutCreate), 30*time.Second, cdnService.CdnDomainConfigRefreshFunc(d.Id(), []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	err = cdnService.WaitForCdnDomain(d.Get("domain_name").(string), Online, DefaultTimeoutMedium)
	if err != nil {
		return WrapError(err)
	}

	return resourceAliCloudCdnDomainConfigRead(d, meta)
}

func resourceAliCloudCdnDomainConfigRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cdnService := &CdnService{client: client}

	v, err := cdnService.DescribeCdnDomainConfig(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	val := v.(map[string]interface{})
	args := make([]interface{}, 0)
	if v, ok := val["FunctionArgs"].(map[string]interface{}); ok {
		if tmp, exist := v["FunctionArg"].([]interface{}); exist {
			args = tmp
		}
	}

	var funArgs []map[string]string
	for _, k := range args {

		arg := k.(map[string]interface{})
		// This two function args is extra, filter them to pass test check.
		if arg["ArgName"] == "aliyun_id" || arg["ArgName"] == "scheme_origin_port" || arg["ArgName"] == "dsl" || arg["ArgName"] == "session_timeout" || arg["ArgName"] == "oss_pri_buckets" {
			continue
		}
		// private_oss_tbl always is changed and used to enable Alibaba Cloud OSS Private Bucket Back to Source Authorization
		if arg["ArgName"] == "private_oss_tbl" {
			continue
		}
		funArgs = append(funArgs, map[string]string{
			"arg_name":  arg["ArgName"].(string),
			"arg_value": arg["ArgValue"].(string),
		})
	}

	parts := strings.Split(d.Id(), ":")
	d.Set("domain_name", parts[0])
	d.Set("function_name", parts[1])
	d.Set("config_id", val["ConfigId"])
	if len(parts) == 2 {
		d.SetId(fmt.Sprintf("%s:%s:%s", parts[0], parts[1], val["ConfigId"]))
	}

	d.Set("status", val["Status"])
	d.Set("function_args", funArgs)
	d.Set("parent_id", val["ParentId"])

	return nil
}

func resourceAliCloudCdnDomainConfigDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cdnService := &CdnService{client: client}

	_, err := cdnService.DescribeCdnDomainConfig(d.Id())
	if err != nil {
		if NotFoundError(err) {
			return nil
		}
		return WrapError(err)
	}

	action := "DeleteSpecificConfig"
	var response map[string]interface{}

	parts, err := ParseResourceId(d.Id(), 3)
	if err != nil {
		return WrapError(err)
	}

	request := map[string]interface{}{
		"RegionId":   client.RegionId,
		"DomainName": parts[0],
		"ConfigId":   parts[2],
	}

	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(client.GetRetryTimeout(d.Timeout(schema.TimeoutDelete)), func() *resource.RetryError {
		response, err = client.RpcPost("Cdn", "2018-05-10", action, nil, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"ServiceBusy"}) || NeedRetry(err) {
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

	return nil
}

func expirationCdnDomainConfigHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if v, ok := m["arg_name"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	if v, ok := m["arg_value"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}
	return hashcode.String(buf.String())
}
