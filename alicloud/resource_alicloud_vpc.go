// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudVpcVpc() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudVpcVpcCreate,
		Read:   resourceAlicloudVpcVpcRead,
		Update: resourceAlicloudVpcVpcUpdate,
		Delete: resourceAlicloudVpcVpcDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cidr_block": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"classic_link_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_options_set_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"enable_ipv6": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ipv6_cidr_block": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_cidr_blocks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_isp": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv6_cidr_block": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ipv6_isp": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_table_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"router_table_id"},
			},
			"router_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_cidr_blocks": {
				Type:       schema.TypeList,
				Optional:   true,
				Deprecated: "Field 'secondary_cidr_blocks' has been deprecated from provider version 1.185.0 and it will be removed in the future version. Please use the new resource 'alicloud_vpc_ipv4_cidr_block'. `secondary_cidr_blocks` attributes and `alicloud_vpc_ipv4_cidr_block` resource cannot be used at the same time.",
				Elem:       &schema.Schema{Type: schema.TypeString},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchema(),
			"user_cidrs": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vswitch_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpc_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"name"},
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"vpc_name"},
				Deprecated:    "Field 'name' has been deprecated from provider version 1.119.0. New field 'vpc_name' instead.",
			},
			"router_table_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"route_table_id"},
				Deprecated:    "Field 'router_table_id' has been deprecated from provider version 1.204.0. New field 'route_table_id' instead.",
			},
		},
	}
}

func resourceAlicloudVpcVpcCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "CreateVpc"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if v, ok := d.GetOk("cidr_block"); ok {
		request["CidrBlock"] = v
	}

	if v, ok := d.GetOk("vpc_name"); ok {
		request["VpcName"] = v
	}
	if v, ok := d.GetOk("name"); ok {
		request["VpcName"] = v
	}

	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}

	if v, ok := d.GetOk("ipv6_cidr_block"); ok {
		request["Ipv6CidrBlock"] = v
	}

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}

	if v, ok := d.GetOk("user_cidrs"); ok {
		jsonPathResult6, err := jsonpath.Get("$", v)
		if err != nil {
			return WrapError(err)
		}
		request["UserCidr"] = convertListToCommaSeparate(jsonPathResult6.([]interface{}))
	}

	if v, ok := d.GetOk("ipv6_isp"); ok {
		request["Ipv6Isp"] = v
	}

	if v, ok := d.GetOkExists("enable_ipv6"); ok {
		request["EnableIpv6"] = v
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"TaskConflict", "UnknownError"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_vpc", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["VpcId"]))

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, vpcServiceV2.VpcVpcStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAlicloudVpcVpcUpdate(d, meta)
}

func resourceAlicloudVpcVpcRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vpcServiceV2 := VpcServiceV2{client}

	object, err := vpcServiceV2.DescribeVpcVpc(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc .DescribeVpcVpc Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("cidr_block", object["cidr_block"])
	d.Set("classic_link_enabled", object["classic_link_enabled"])
	d.Set("create_time", object["create_time"])
	d.Set("description", object["description"])
	d.Set("dhcp_options_set_id", object["dhcp_options_set_id"])
	d.Set("ipv6_cidr_block", object["ipv6_cidr_block"])
	d.Set("ipv6_cidr_blocks", object["ipv6_cidr_blocks"])
	d.Set("is_default", object["is_default"])
	d.Set("resource_group_id", object["resource_group_id"])
	d.Set("route_table_id", object["route_table_id"])
	d.Set("router_id", object["router_id"])
	d.Set("secondary_cidr_blocks", object["secondary_cidr_blocks"])
	d.Set("status", object["status"])
	d.Set("tags", tagsToMap(object["tags"]))
	d.Set("user_cidrs", object["user_cidrs"])
	d.Set("vswitch_ids", object["vswitch_ids"])
	d.Set("vpc_id", object["vpc_id"])
	d.Set("vpc_name", object["vpc_name"])

	d.Set("name", d.Get("vpc_name"))
	d.Set("router_table_id", d.Get("route_table_id"))
	return nil
}

func resourceAlicloudVpcVpcUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	d.Partial(true)
	update = false
	action := "ModifyVpcAttribute"
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["VpcId"] = d.Id()
	request["RegionId"] = client.RegionId

	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		if v, ok := d.GetOk("description"); ok {
			request["Description"] = v
		}
	}
	if !d.IsNewResource() && (d.HasChange("vpc_name") || d.HasChange("name")) {
		update = true
		if d.HasChange("vpc_name") {
			if v, ok := d.GetOk("vpc_name"); ok {
				request["VpcName"] = v
			}
		}
		if d.HasChange("name") {
			if v, ok := d.GetOk("name"); ok {
				request["VpcName"] = v
			}
		}
	}
	if !d.IsNewResource() && d.HasChange("cidr_block") {
		update = true
		if v, ok := d.GetOk("cidr_block"); ok {
			request["CidrBlock"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("ipv6_cidr_block") {
		update = true
		if v, ok := d.GetOk("ipv6_cidr_block"); ok {
			request["Ipv6CidrBlock"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("enable_ipv6") {
		update = true
		if v, ok := d.GetOkExists("enable_ipv6"); ok {
			request["EnableIPv6"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("ipv6_isp") {
		update = true
		if v, ok := d.GetOk("ipv6_isp"); ok {
			request["Ipv6Isp"] = v
		}
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{"OperationFailed.LastTokenProcessing", "LastTokenProcessing", "OperationFailed.QueryCenIpv6Status", "IncorrectStatus.%s", "OperationConflict", "SystemBusy", "ServiceUnavailable", "IncorrectVpcStatus"}) || NeedRetry(err) {
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
			stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcVpcStateRefreshFunc(d.Id(), []string{}))
			if _, err := stateConf.WaitForState(); err != nil {
				return WrapErrorf(err, IdMsg, d.Id())
			}
		}
		d.SetPartial("description")
		d.SetPartial("vpc_name")
		d.SetPartial("cidr_block")
		d.SetPartial("ipv6_cidr_block")
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
	request["ResourceType"] = "VPC"
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

	if d.HasChange("classic_link_enabled") {
		client := meta.(*connectivity.AliyunClient)
		vpcServiceV2 := VpcServiceV2{client}
		object, err := vpcServiceV2.DescribeVpcVpc(d.Id())
		if err != nil {
			return WrapError(err)
		}

		target := d.Get("classic_link_enabled").(bool)
		if object["classic_link_enabled"].(bool) != target {
			if target == true {
				action = "EnableVpcClassicLink"
				conn, err = client.NewVpcClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["VpcId"] = d.Id()
				request["RegionId"] = client.RegionId
				request["ClientToken"] = buildClientToken(action)
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
					if err != nil {
						if IsExpectedErrors(err, []string{"IncorrectVpcStatus"}) || NeedRetry(err) {
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

			}
			if target == false {
				action = "DisableVpcClassicLink"
				conn, err = client.NewVpcClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["VpcId"] = d.Id()
				request["RegionId"] = client.RegionId
				request["ClientToken"] = buildClientToken(action)
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
					if err != nil {
						if IsExpectedErrors(err, []string{"InternalError", "IncorrectVpcStatus"}) || NeedRetry(err) {
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

			}
		}
	}

	update = false
	if d.HasChange("secondary_cidr_blocks") {
		update = true
		oldEntry, newEntry := d.GetChange("secondary_cidr_blocks")
		removed := oldEntry
		added := newEntry

		if len(removed.([]interface{})) > 0 {
			secondaryCidrBlocks := removed.([]interface{})

			for _, item := range secondaryCidrBlocks {
				action = "UnassociateVpcCidrBlock"
				conn, err = client.NewVpcClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["VpcId"] = d.Id()
				request["RegionId"] = client.RegionId

				if v, ok := item.(string); ok {
					jsonPathResult, err := jsonpath.Get("$", v)
					if err != nil {
						return WrapError(err)
					}
					request["SecondaryCidrBlock"] = jsonPathResult
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
					stateConf := BuildStateConf([]string{}, []string{"Created", "Available"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcVpcStateRefreshFunc(d.Id(), []string{}))
					if _, err := stateConf.WaitForState(); err != nil {
						return WrapErrorf(err, IdMsg, d.Id())
					}
				}

			}
			d.SetPartial("secondary_cidr_blocks")
		}

		if len(added.([]interface{})) > 0 {
			secondaryCidrBlocks := added.([]interface{})

			for _, item := range secondaryCidrBlocks {
				action = "AssociateVpcCidrBlock"
				conn, err = client.NewVpcClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["VpcId"] = d.Id()
				request["RegionId"] = client.RegionId

				if v, ok := item.(string); ok {
					jsonPathResult, err := jsonpath.Get("$", v)
					if err != nil {
						return WrapError(err)
					}
					request["SecondaryCidrBlock"] = jsonPathResult
				}
				request["IpVersion"] = "IPV4"
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
					stateConf := BuildStateConf([]string{}, []string{"Created", "Available"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcVpcStateRefreshFunc(d.Id(), []string{}))
					if _, err := stateConf.WaitForState(); err != nil {
						return WrapErrorf(err, IdMsg, d.Id())
					}
				}

			}
			d.SetPartial("secondary_cidr_blocks")
		}
	}
	update = false
	if d.HasChange("tags") {
		update = true
		vpcServiceV2 := VpcServiceV2{client}
		if err := vpcServiceV2.SetResourceTags(d, "VPC"); err != nil {
			return WrapError(err)
		}
	}
	update = false
	if d.HasChange("ipv6_cidr_blocks") {
		update = true
		oldEntry, newEntry := d.GetChange("ipv6_cidr_blocks")
		removed := oldEntry
		added := newEntry

		if len(removed.([]interface{})) > 0 {
			ipv6CidrBlocks := removed.([]interface{})

			for _, item := range ipv6CidrBlocks {
				action = "UnassociateVpcCidrBlock"
				conn, err = client.NewVpcClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["VpcId"] = d.Id()
				request["RegionId"] = client.RegionId

				if v, ok := item.(map[string]interface{}); ok {
					jsonPathResult, err := jsonpath.Get("$.ipv6_cidr_block", v)
					if err != nil {
						return WrapError(err)
					}
					request["IPv6CidrBlock"] = jsonPathResult
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
					stateConf := BuildStateConf([]string{}, []string{"Created", "Available"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcVpcStateRefreshFunc(d.Id(), []string{}))
					if _, err := stateConf.WaitForState(); err != nil {
						return WrapErrorf(err, IdMsg, d.Id())
					}
				}

			}
			d.SetPartial("ipv6_cidr_blocks")
		}

		if len(added.([]interface{})) > 0 {
			ipv6CidrBlocks := added.([]interface{})

			for _, item := range ipv6CidrBlocks {
				action = "AssociateVpcCidrBlock"
				conn, err = client.NewVpcClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["VpcId"] = d.Id()
				request["RegionId"] = client.RegionId

				request["IpVersion"] = "IPV6"
				if v, ok := item.(map[string]interface{}); ok {
					jsonPathResult, err := jsonpath.Get("$.ipv6_cidr_block", v)
					if err != nil {
						return WrapError(err)
					}
					request["IPv6CidrBlock"] = jsonPathResult
				}
				if v, ok := item.(map[string]interface{}); ok {
					jsonPathResult1, err := jsonpath.Get("$.ipv6_isp", v)
					if err != nil {
						return WrapError(err)
					}
					request["Ipv6Isp"] = jsonPathResult1
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
					stateConf := BuildStateConf([]string{}, []string{"Created", "Available"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcVpcStateRefreshFunc(d.Id(), []string{}))
					if _, err := stateConf.WaitForState(); err != nil {
						return WrapErrorf(err, IdMsg, d.Id())
					}
				}

			}
			d.SetPartial("ipv6_cidr_blocks")
		}
	}
	d.Partial(false)
	return resourceAlicloudVpcVpcRead(d, meta)
}

func resourceAlicloudVpcVpcDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "DeleteVpc"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["VpcId"] = d.Id()
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
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
		if IsExpectedErrors(err, []string{}) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcVpcStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}
