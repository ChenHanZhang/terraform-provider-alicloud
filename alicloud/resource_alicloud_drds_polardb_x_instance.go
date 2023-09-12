// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudDrdsPolardbXInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudDrdsPolardbXInstanceCreate,
		Read:   resourceAliCloudDrdsPolardbXInstanceRead,
		Update: resourceAliCloudDrdsPolardbXInstanceUpdate,
		Delete: resourceAliCloudDrdsPolardbXInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cn_class": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"cn_node_count": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dn_class": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"dn_node_count": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_read_db_instance": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"primary_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tertiary_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"topology_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vswitch_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudDrdsPolardbXInstanceCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateDBInstance"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewDrdsClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	request["NetworkType"] = "VPC"
	if v, ok := d.GetOk("vpc_id"); ok {
		request["VPCId"] = v
	}
	if v, ok := d.GetOk("vswitch_id"); ok {
		request["VSwitchId"] = v
	}
	request["EngineVersion"] = "5.7"
	request["PrimaryDBInstanceName"] = "null"
	request["TopologyType"] = d.Get("topology_type")
	if v, ok := d.GetOk("primary_zone"); ok {
		request["PrimaryZone"] = v
	}
	if v, ok := d.GetOk("secondary_zone"); ok {
		request["SecondaryZone"] = v
	}
	if v, ok := d.GetOk("tertiary_zone"); ok {
		request["TertiaryZone"] = v
	}
	if v, ok := d.GetOk("cn_class"); ok {
		request["CnClass"] = v
	}
	if v, ok := d.GetOk("dn_class"); ok {
		request["DnClass"] = v
	}
	if v, ok := d.GetOk("cn_node_count"); ok {
		request["CNNodeCount"] = v
	}
	if v, ok := d.GetOk("dn_node_count"); ok {
		request["DNNodeCount"] = v
	}
	request["PayType"] = "Postpaid"
	request["UsedTime"] = "1"
	request["Period"] = "Month"
	request["AutoRenew"] = "false"
	request["ZoneId"] = "null"
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if v, ok := d.GetOkExists("is_read_db_instance"); ok {
		request["IsReadDBInstance"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-02-02"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		request["ClientToken"] = buildClientToken(action)

		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_drds_polardb_x_instance", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["DBInstanceName"]))

	return resourceAliCloudDrdsPolardbXInstanceUpdate(d, meta)
}

func resourceAliCloudDrdsPolardbXInstanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	drdsServiceV2 := DrdsServiceV2{client}

	objectRaw, err := drdsServiceV2.DescribeDrdsPolardbXInstance(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_drds_polardb_x_instance DescribeDrdsPolardbXInstance Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("cn_class", objectRaw["CnNodeClassCode"])
	d.Set("cn_node_count", objectRaw["CnNodeCount"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("dn_class", objectRaw["DnNodeClassCode"])
	d.Set("dn_node_count", objectRaw["DnNodeCount"])
	d.Set("resource_group_id", objectRaw["ResourceGroupId"])
	d.Set("status", objectRaw["Status"])
	d.Set("vswitch_id", objectRaw["VSwitchId"])
	d.Set("vpc_id", objectRaw["VPCId"])
	d.Set("zone", objectRaw["ZoneId"])

	return nil
}

func resourceAliCloudDrdsPolardbXInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	d.Partial(true)
	action := "UpdatePolarDBXInstanceNode"
	conn, err := client.NewDrdsClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["DBInstanceName"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("cn_node_count") {
		update = true
		request["CNNodeCount"] = d.Get("cn_node_count")
	}

	if !d.IsNewResource() && d.HasChange("dn_node_count") {
		update = true
		request["DNNodeCount"] = d.Get("dn_node_count")
	}

	request["DbInstanceNodeCount"] = "0"
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-02-02"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			request["ClientToken"] = buildClientToken(action)

			if err != nil {
				if NeedRetry(err) {
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
		d.SetPartial("cn_node_count")
		d.SetPartial("dn_node_count")
	}
	update = false
	action = "ChangeResourceGroup"
	conn, err = client.NewDrdsClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["ResourceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
		request["NewResourceGroupId"] = d.Get("resource_group_id")
	}

	request["ResourceType"] = "PolarDBXInstance"
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-02-02"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

			if err != nil {
				if NeedRetry(err) {
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

	d.Partial(false)
	return resourceAliCloudDrdsPolardbXInstanceRead(d, meta)
}

func resourceAliCloudDrdsPolardbXInstanceDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteDBInstance"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewDrdsClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["DBInstanceName"] = d.Id()
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-02-02"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		if IsExpectedErrors(err, []string{"DBInstance.NotFound"}) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}

func convertDrdsDBInstancePayTypeResponse(source interface{}) interface{} {
	switch source {
	}
	return source
}
