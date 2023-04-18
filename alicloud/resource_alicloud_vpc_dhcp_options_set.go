// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAliCloudVpcDhcpOptionsSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudVpcDhcpOptionsSetCreate,
		Read:   resourceAlicloudVpcDhcpOptionsSetRead,
		Update: resourceAlicloudVpcDhcpOptionsSetUpdate,
		Delete: resourceAlicloudVpcDhcpOptionsSetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"all": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"associate_vpcs": {
				Type:       schema.TypeSet,
				Optional:   true,
				Computed:   true,
				Deprecated: "Field 'associate_vpcs' has been deprecated from provider version 1.153.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_dhcp_options_set_attachment' to attach DhcpOptionsSet and Vpc.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"associate_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dhcp_options_set_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_options_set_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_options_set_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^[a-zA-Z\u4E00-\u9FA5][\u4E00-\u9FA5A-Za-z0-9_-]{2,128}$"), "The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character."),
			},
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_name_servers": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ipv6_lease_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lease_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchema(),
		},
	}
}

func resourceAlicloudVpcDhcpOptionsSetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "CreateDhcpOptionsSet"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if v, ok := d.GetOk("domain_name_servers"); ok {
		request["DomainNameServers"] = v
	}

	if v, ok := d.GetOk("dhcp_options_set_name"); ok {
		request["DhcpOptionsSetName"] = v
	}

	if v, ok := d.GetOk("dhcp_options_set_description"); ok {
		request["DhcpOptionsSetDescription"] = v
	}

	if v, ok := d.GetOk("domain_name"); ok {
		request["DomainName"] = v
	}

	if v, ok := d.GetOk("lease_time"); ok {
		request["LeaseTime"] = v
	}

	if v, ok := d.GetOk("ipv6_lease_time"); ok {
		request["Ipv6LeaseTime"] = v
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_vpc_dhcp_options_set", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["DhcpOptionsSetId"]))

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Available", "InUse"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, vpcServiceV2.VpcDhcpOptionsSetStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAlicloudVpcDhcpOptionsSetUpdate(d, meta)
}

func resourceAlicloudVpcDhcpOptionsSetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vpcServiceV2 := VpcServiceV2{client}

	object, err := vpcServiceV2.DescribeVpcDhcpOptionsSet(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_dhcp_options_set .DescribeVpcDhcpOptionsSet Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("associate_vpcs", object["associate_vpcs"])
	d.Set("dhcp_options_set_description", object["dhcp_options_set_description"])
	d.Set("dhcp_options_set_id", object["dhcp_options_set_id"])
	d.Set("dhcp_options_set_name", object["dhcp_options_set_name"])
	d.Set("domain_name", object["domain_name"])
	d.Set("domain_name_servers", object["domain_name_servers"])
	d.Set("ipv6_lease_time", object["ipv6_lease_time"])
	d.Set("lease_time", object["lease_time"])
	d.Set("resource_group_id", object["resource_group_id"])
	d.Set("status", object["status"])
	d.Set("tags", tagsToMap(object["tags"]))

	return nil
}

func resourceAlicloudVpcDhcpOptionsSetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	d.Partial(true)
	update = false
	action := "UpdateDhcpOptionsSetAttribute"
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["DhcpOptionsSetId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("domain_name_servers") {
		update = true
		if v, ok := d.GetOk("domain_name_servers"); ok {
			request["DomainNameServers"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("domain_name") {
		update = true
		if v, ok := d.GetOk("domain_name"); ok {
			request["DomainName"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("dhcp_options_set_name") {
		update = true
		if v, ok := d.GetOk("dhcp_options_set_name"); ok {
			request["DhcpOptionsSetName"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("dhcp_options_set_description") {
		update = true
		if v, ok := d.GetOk("dhcp_options_set_description"); ok {
			request["DhcpOptionsSetDescription"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("lease_time") {
		update = true
		if v, ok := d.GetOk("lease_time"); ok {
			request["LeaseTime"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("ipv6_lease_time") {
		update = true
		if v, ok := d.GetOk("ipv6_lease_time"); ok {
			request["Ipv6LeaseTime"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("dry_run") {
		update = true
		if v, ok := d.GetOkExists("dry_run"); ok {
			request["DryRun"] = v
		}
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		{
			vpcServiceV2 := VpcServiceV2{client}
			stateConf := BuildStateConf([]string{}, []string{"Available", "InUse"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcDhcpOptionsSetStateRefreshFunc(d.Id(), []string{}))
			if _, err := stateConf.WaitForState(); err != nil {
				return WrapErrorf(err, IdMsg, d.Id())
			}
		}
		d.SetPartial("domain_name_servers")
		d.SetPartial("domain_name")
		d.SetPartial("dhcp_options_set_name")
		d.SetPartial("dhcp_options_set_description")
		d.SetPartial("lease_time")
		d.SetPartial("ipv6_lease_time")
	}
	update = false
	action = "MoveResourceGroup"
	conn, err = client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId"] = d.Id()
	request["RegionId"] = client.RegionId

	if !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
		if v, ok := d.GetOk("resource_group_id"); ok {
			request["NewResourceGroupId"] = v
		}
	}
	request["ResourceType"] = "DhcpOptionsSet"
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("resource_group_id")
	}

	update = false
	if d.HasChange("associate_vpcs") {
		update = true
		oldEntry, newEntry := d.GetChange("associate_vpcs")
		oldEntrySet := oldEntry.(*schema.Set)
		newEntrySet := newEntry.(*schema.Set)
		removed := oldEntrySet.Difference(newEntrySet)
		added := newEntrySet.Difference(oldEntrySet)

		if removed.Len() > 0 {
			associateVpcs := removed.List()

			for _, item := range associateVpcs {
				action = "DetachDhcpOptionsSetFromVpc"
				conn, err = client.NewVpcClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["DhcpOptionsSetId"] = d.Id()
				request["RegionId"] = client.RegionId
				request["ClientToken"] = buildClientToken(action)
				if v, ok := item.(map[string]interface{}); ok {
					jsonPathResult, err := jsonpath.Get("$.vpc_id", v)
					if err != nil {
						return WrapError(err)
					}
					request["VpcId"] = jsonPathResult
				}
				if v, ok := item.(map[string]interface{}); ok {
					request["DryRun"] = v
				}
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
					if err != nil {
						if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
							wait()
							return resource.RetryableError(err)
						}
						return resource.NonRetryableError(err)
					}
					addDebug(action, response, request)
					return nil
				})
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}
				{
					vpcServiceV2 := VpcServiceV2{client}
					stateConf := BuildStateConf([]string{}, []string{"Available", "InUse"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcDhcpOptionsSetStateRefreshFunc(d.Id(), []string{}))
					if _, err := stateConf.WaitForState(); err != nil {
						return WrapErrorf(err, IdMsg, d.Id())
					}
				}

			}
			d.SetPartial("associate_vpcs")
		}

		if added.Len() > 0 {
			associateVpcs := added.List()

			for _, item := range associateVpcs {
				action = "AttachDhcpOptionsSetToVpc"
				conn, err = client.NewVpcClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["DhcpOptionsSetId"] = d.Id()
				request["RegionId"] = client.RegionId
				request["ClientToken"] = buildClientToken(action)
				if v, ok := item.(map[string]interface{}); ok {
					jsonPathResult, err := jsonpath.Get("$.vpc_id", v)
					if err != nil {
						return WrapError(err)
					}
					request["VpcId"] = jsonPathResult
				}
				if v, ok := item.(map[string]interface{}); ok {
					request["DryRun"] = v
				}
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
					if err != nil {
						if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
							wait()
							return resource.RetryableError(err)
						}
						return resource.NonRetryableError(err)
					}
					addDebug(action, response, request)
					return nil
				})
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}
				{
					vpcServiceV2 := VpcServiceV2{client}
					stateConf := BuildStateConf([]string{}, []string{"Available", "InUse"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcDhcpOptionsSetStateRefreshFunc(d.Id(), []string{}))
					if _, err := stateConf.WaitForState(); err != nil {
						return WrapErrorf(err, IdMsg, d.Id())
					}
				}

			}
			d.SetPartial("associate_vpcs")
		}

	}
	update = false
	if d.HasChange("tags") {
		update = true
		vpcServiceV2 := VpcServiceV2{client}
		if err := vpcServiceV2.SetResourceTags(d, "DhcpOptionsSet"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAlicloudVpcDhcpOptionsSetRead(d, meta)
}

func resourceAlicloudVpcDhcpOptionsSetDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "DeleteDhcpOptionsSet"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["DhcpOptionsSetId"] = d.Id()
	request["RegionId"] = client.RegionId

	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"IncorrectStatus.DhcpOptionsSet"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidDhcpOptionsSetId.NotFound"}) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcDhcpOptionsSetStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}
