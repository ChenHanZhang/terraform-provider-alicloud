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

func resourceAliCloudRdsGadInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudRdsGadInstanceCreate,
		Read:   resourceAliCloudRdsGadInstanceRead,
		Delete: resourceAliCloudRdsGadInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
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
		},
	}
}

func resourceAliCloudRdsGadInstanceCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateGADInstance"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId

	unitNodeDataList := make(map[string]interface{})

	if v, ok := d.GetOk("gad_instance_members"); ok {
		engineVersion1, _ := jsonpath.Get("$.engine_version", v)
		if engineVersion1 != nil && engineVersion1 != "" {
			unitNodeDataList["EngineVersion"] = engineVersion1
		}
	}

	if v, ok := d.GetOk("gad_instance_members"); ok {
		regionId, _ := jsonpath.Get("$.region_id", v)
		if regionId != nil && regionId != "" {
			unitNodeDataList["RegionID"] = regionId
		}
	}

	if v, ok := d.GetOk("gad_instance_members"); ok {
		engine1, _ := jsonpath.Get("$.engine", v)
		if engine1 != nil && engine1 != "" {
			unitNodeDataList["Engine"] = engine1
		}
	}

	UnitNodeMap := make([]interface{}, 0)
	UnitNodeMap = append(UnitNodeMap, unitNodeDataList)
	request["UnitNode"] = UnitNodeMap

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Rds", "2014-08-15", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_rds_gad_instance", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.Result.GadInstanceName", response)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudRdsGadInstanceRead(d, meta)
}

func resourceAliCloudRdsGadInstanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	rdsServiceV2 := RdsServiceV2{client}

	objectRaw, err := rdsServiceV2.DescribeRdsGadInstance(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_rds_gad_instance DescribeRdsGadInstance Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreationTime"])
	d.Set("status", objectRaw["Status"])

	gadInstanceMembersRawObj, _ := jsonpath.Get("$.GadInstanceMembers[*]", objectRaw)
	gadInstanceMembersRaw := make([]interface{}, 0)
	if gadInstanceMembersRawObj != nil {
		gadInstanceMembersRaw = convertToInterfaceArray(gadInstanceMembersRawObj)
	}

	d.Set("resource_group_id", gadInstanceMembersRaw["ResourceGroupId"])

	return nil
}

func resourceAliCloudRdsGadInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Gad Instance. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
