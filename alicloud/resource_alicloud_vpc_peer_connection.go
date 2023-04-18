// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"regexp"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAliCloudVpcPeerConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudVpcPeerConnectionCreate,
		Read:   resourceAlicloudVpcPeerConnectionRead,
		Update: resourceAlicloudVpcPeerConnectionUpdate,
		Delete: resourceAlicloudVpcPeerConnectionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"accepting_ali_uid": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"accepting_region_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"accepting_vpc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bandwidth": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"biz_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_all": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"expire_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modify_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peer_connection_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("(.*)"), "The name of the resource"),
			},
			"peering_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": tagsSchema(),
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAlicloudVpcPeerConnectionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "CreateVpcPeerConnection"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcpeerClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if v, ok := d.GetOk("vpc_id"); ok {
		request["VpcId"] = v
	}

	if v, ok := d.GetOk("accepting_ali_uid"); ok {
		request["AcceptingAliUid"] = v
	}

	if v, ok := d.GetOk("accepting_region_id"); ok {
		request["AcceptingRegionId"] = v
	}

	if v, ok := d.GetOk("accepting_vpc_id"); ok {
		request["AcceptingVpcId"] = v
	}

	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}

	if v, ok := d.GetOk("peer_connection_name"); ok {
		request["Name"] = v
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2022-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_vpc_peer_connection", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["InstanceId"]))

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Activated", "Accepting"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, vpcServiceV2.VpcPeerConnectionStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAlicloudVpcPeerConnectionUpdate(d, meta)
}

func resourceAlicloudVpcPeerConnectionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vpcServiceV2 := VpcServiceV2{client}

	object, err := vpcServiceV2.DescribeVpcPeerConnection(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_peer_connection .DescribeVpcPeerConnection Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("accepting_ali_uid", object["accepting_ali_uid"])
	d.Set("accepting_region_id", object["accepting_region_id"])
	d.Set("accepting_vpc_id", object["accepting_vpc_id"])
	d.Set("bandwidth", object["bandwidth"])
	d.Set("biz_status", object["biz_status"])
	d.Set("create_time", object["create_time"])
	d.Set("description", object["description"])
	d.Set("expire_time", object["expire_time"])
	d.Set("modify_time", object["modify_time"])
	d.Set("peer_connection_name", object["peer_connection_name"])
	d.Set("peering_id", object["peering_id"])
	d.Set("status", object["status"])
	d.Set("tags", tagsToMap(object["tags"]))
	d.Set("vpc_id", object["vpc_id"])

	return nil
}

func resourceAlicloudVpcPeerConnectionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	d.Partial(true)
	update = false
	action := "ModifyVpcPeerConnection"
	conn, err := client.NewVpcpeerClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["InstanceId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		if v, ok := d.GetOk("description"); ok {
			request["Description"] = v
		}
	}
	if d.HasChange("bandwidth") {
		update = true
		if v, ok := d.GetOk("bandwidth"); ok {
			request["Bandwidth"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("peer_connection_name") {
		update = true
		if v, ok := d.GetOk("peer_connection_name"); ok {
			request["Name"] = v
		}
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2022-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
			stateConf := BuildStateConf([]string{}, []string{"Activated"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcPeerConnectionStateRefreshFunc(d.Id(), []string{}))
			if _, err := stateConf.WaitForState(); err != nil {
				return WrapErrorf(err, IdMsg, d.Id())
			}
		}
		d.SetPartial("description")
		d.SetPartial("bandwidth")
		d.SetPartial("peer_connection_name")
	}

	if d.HasChange("status") {
		client := meta.(*connectivity.AliyunClient)
		vpcServiceV2 := VpcServiceV2{client}
		object, err := vpcServiceV2.DescribeVpcPeerConnection(d.Id())
		if err != nil {
			return WrapError(err)
		}

		target := d.Get("status").(string)
		if object["status"].(string) != target {
			if target == "Activated" {
				action = "AcceptVpcPeerConnection"
				conn, err = client.NewVpcpeerClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["InstanceId"] = d.Id()

				request["ClientToken"] = buildClientToken(action)
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2022-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
					stateConf := BuildStateConf([]string{}, []string{"Activated"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcPeerConnectionStateRefreshFunc(d.Id(), []string{}))
					if _, err := stateConf.WaitForState(); err != nil {
						return WrapErrorf(err, IdMsg, d.Id())
					}
				}

			}
			if target == "Rejected" {
				action = "RejectVpcPeerConnection"
				conn, err = client.NewVpcpeerClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["InstanceId"] = d.Id()

				request["ClientToken"] = buildClientToken(action)
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2022-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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

			}
		}
	}

	update = false
	if d.HasChange("tags") {
		update = true
		vpcServiceV2 := VpcServiceV2{client}
		if err := vpcServiceV2.SetResourceTags(d, "PeerConnection"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAlicloudVpcPeerConnectionRead(d, meta)
}

func resourceAlicloudVpcPeerConnectionDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "DeleteVpcPeerConnection"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcpeerClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["InstanceId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2022-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcPeerConnectionStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}
