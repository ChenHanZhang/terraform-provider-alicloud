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

func resourceAliCloudPolardbAICluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPolardbAIClusterCreate,
		Read:   resourceAliCloudPolardbAIClusterRead,
		Update: resourceAliCloudPolardbAIClusterUpdate,
		Delete: resourceAliCloudPolardbAIClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"ack_admin": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ai_node_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"auto_renew": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_use_coupon": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"db_cluster_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"db_node_class": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"extension": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inference_engine": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"kube_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"kube_config": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"kube_management": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"kubernetes_config": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"management_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"pay_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"payment_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"promotion_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_space": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"storage_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": tagsSchema(),
			"time_slices": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_time": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"begin_time": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"used_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vswitch_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudPolardbAIClusterCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateAIDBCluster"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("db_cluster_id"); ok {
		request["DBClusterId"] = v
	}
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	request["DBNodeClass"] = d.Get("db_node_class")
	request["StorageType"] = d.Get("storage_type")
	request["VSwitchId"] = d.Get("vswitch_id")
	if v, ok := d.GetOk("password"); ok {
		request["Password"] = v
	}
	if v, ok := d.GetOk("used_time"); ok {
		request["UsedTime"] = v
	}
	if v, ok := d.GetOk("security_group_id"); ok {
		request["SecurityGroupId"] = v
	}
	request["PayType"] = d.Get("pay_type")
	if v, ok := d.GetOk("ai_node_type"); ok {
		request["KubeType"] = v
	}
	if v, ok := d.GetOk("ack_admin"); ok {
		request["AckAdmin"] = v
	}
	if v, ok := d.GetOk("extension"); ok {
		request["Extension"] = v
	}
	if v, ok := d.GetOk("kube_cluster_id"); ok {
		request["KubeClusterId"] = v
	}
	if v, ok := d.GetOk("auto_renew"); ok {
		request["AutoRenew"] = v
	}
	if v, ok := d.GetOk("time_slices"); ok {
		timeSlicesMapsArray := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(v) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["BeginTime"] = dataLoopTmp["begin_time"]
			dataLoopMap["EndTime"] = dataLoopTmp["end_time"]
			timeSlicesMapsArray = append(timeSlicesMapsArray, dataLoopMap)
		}
		request["TimeSlices"] = timeSlicesMapsArray
	}

	if v, ok := d.GetOk("promotion_code"); ok {
		request["PromotionCode"] = v
	}
	if v, ok := d.GetOk("kube_management"); ok {
		request["KubeManagement"] = v
	}
	if v, ok := d.GetOkExists("storage_space"); ok {
		request["StorageSpace"] = v
	}
	if v, ok := d.GetOk("inference_engine"); ok {
		request["InferenceEngine"] = v
	}
	if v, ok := d.GetOk("kube_config"); ok {
		request["KubeConfig"] = v
	}
	if v, ok := d.GetOkExists("auto_use_coupon"); ok {
		request["AutoUseCoupon"] = v
	}
	if v, ok := d.GetOk("management_mode"); ok {
		request["ManagementMode"] = v
	}
	if v, ok := d.GetOk("kubernetes_config"); ok {
		request["KubernetesConfig"] = v
	}
	if v, ok := d.GetOk("db_cluster_description"); ok {
		request["DBClusterDescription"] = v
	}
	request["VPCId"] = d.Get("vpc_id")
	if v, ok := d.GetOk("period"); ok {
		request["Period"] = v
	}
	if v, ok := d.GetOk("zone_id"); ok {
		request["ZoneId"] = v
	}
	if v, ok := d.GetOk("mode_name"); ok {
		request["ModeName"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_polardb_a_i_cluster", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["DBClusterId"]))

	action = "TagResources"
	request = make(map[string]interface{})
	if v, ok := d.GetOk("db_cluster_id"); ok {
		request["ResourceId.1"] = v
	}
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	request["ResourceType"] = "aicluster"
	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_polardb_a_i_cluster", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("ResourceId[0]", request)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudPolardbAIClusterRead(d, meta)
}

func resourceAliCloudPolardbAIClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	polardbServiceV2 := PolardbServiceV2{client}

	objectRaw, err := polardbServiceV2.DescribePolardbAICluster(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_polardb_a_i_cluster DescribePolardbAICluster Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("ai_node_type", objectRaw["AiNodeType"])
	d.Set("db_cluster_description", objectRaw["DBClusterDescription"])
	d.Set("kube_cluster_id", objectRaw["KubeClusterId"])
	d.Set("pay_type", objectRaw["PayType"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("storage_type", objectRaw["StorageType"])
	d.Set("vpc_id", objectRaw["VPCId"])
	d.Set("vswitch_id", objectRaw["VSwitchId"])
	d.Set("zone_id", objectRaw["ZoneId"])
	d.Set("db_cluster_id", objectRaw["DBClusterId"])

	dBNodesRawArrayObj, _ := jsonpath.Get("$.DBNodes[*]", objectRaw)
	dBNodesRawArray := make([]interface{}, 0)
	if dBNodesRawArrayObj != nil {
		dBNodesRawArray = convertToInterfaceArray(dBNodesRawArrayObj)
	}
	dBNodesRaw := make(map[string]interface{})
	if len(dBNodesRawArray) > 0 {
		dBNodesRaw = dBNodesRawArray[0].(map[string]interface{})
	}

	d.Set("db_node_class", dBNodesRaw["DBNodeClass"])

	objectRaw, err = polardbServiceV2.DescribeAIClusterListTagResources(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	tagsMaps := tagResourceChildRaw.(map[string]interface{})
	d.Set("tags", tagsToMap(tagsMaps))

	d.Set("db_cluster_id", d.Id())

	return nil
}

func resourceAliCloudPolardbAIClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "ModifyAIDBClusterDescription"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if d.HasChange("db_cluster_description") {
		update = true
	}
	request["DBClusterDescription"] = d.Get("db_cluster_description")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
	update = false
	action = "TagResources"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ResourceId.1"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("tags") {
		update = true
	}
	if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	request["ResourceType"] = "aicluster"
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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

	return resourceAliCloudPolardbAIClusterRead(d, meta)
}

func resourceAliCloudPolardbAIClusterDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteAIDBCluster"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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

	action = "UntagResources"
	request = make(map[string]interface{})
	request["ResourceId.1"] = d.Id()
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	request["ResourceType"] = "aicluster"
	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
