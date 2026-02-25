// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudPolardbApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPolardbApplicationCreate,
		Read:   resourceAliCloudPolardbApplicationRead,
		Update: resourceAliCloudPolardbApplicationUpdate,
		Delete: resourceAliCloudPolardbApplicationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"application_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"architecture": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"auto_create_polar_fs": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"auto_renew": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"auto_use_coupon": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"components": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"security_groups": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"security_group_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"security_group_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"region_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"net_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"component_type": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"component_replica": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"topology": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parents": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"children": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"layer": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"scale_max": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"component_max_replica": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"scale_min": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"security_ip_arrays": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"security_ip_array_tag": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"security_ip_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"security_ip_list": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"security_ip_array_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"security_ip_net_type": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
						"component_class_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"component_replica_group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"component_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"component_class": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"db_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"endpoints": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"endpoint_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"net_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"modify_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"payment_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"PayAsYouGo", "Subscription"}, false),
			},
			"period": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"Year", "Month"}, false),
			},
			"polar_fs_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"promotion_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"security_group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"security_group_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"region_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"net_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"security_ip_arrays": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"security_ip_array_tag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"security_ip_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"security_ip_list": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"security_ip_array_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"security_ip_net_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchema(),
			"used_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vswitch_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudPolardbApplicationCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateApplication"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("components"); ok {
		componentsMapsArray := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(v) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["SecurityGroups"] = dataLoopTmp["security_groups"]
			dataLoopMap["ScaleMin"] = dataLoopTmp["scale_min"]
			dataLoopMap["ScaleMax"] = dataLoopTmp["scale_max"]
			dataLoopMap["SecurityIPList"] = dataLoopTmp["security_ip_arrays_security_ip_list"]
			dataLoopMap["SecurityIPNetType"] = dataLoopTmp["security_ip_arrays_security_ip_net_type"]
			dataLoopMap["ComponentType"] = dataLoopTmp["component_type"]
			dataLoopMap["SecurityIPArrayName"] = dataLoopTmp["security_ip_arrays_security_ip_array_name"]
			dataLoopMap["ComponentReplica"] = dataLoopTmp["component_replica"]
			dataLoopMap["ComponentMaxReplica"] = dataLoopTmp["component_max_replica"]
			dataLoopMap["ComponentClass"] = dataLoopTmp["component_class"]
			dataLoopMap["SecurityIPType"] = dataLoopTmp["security_ip_arrays_security_ip_type"]
			componentsMapsArray = append(componentsMapsArray, dataLoopMap)
		}
		componentsMapsJson, err := json.Marshal(componentsMapsArray)
		if err != nil {
			return WrapError(err)
		}
		request["Components"] = string(componentsMapsJson)
	}

	if v, ok := d.GetOk("security_groups"); ok {
		securityGroupsSecurityGroupIdJsonPath, err := jsonpath.Get("$.security_group_id", v)
		if err == nil && securityGroupsSecurityGroupIdJsonPath != "" {
			request["SecurityGroupId"] = securityGroupsSecurityGroupIdJsonPath
		}
	}
	request["PayType"] = convertPolardbApplicationPayTypeRequest(d.Get("payment_type").(string))
	if v, ok := d.GetOk("vswitch_id"); ok {
		request["VSwitchId"] = v
	}
	if v, ok := d.GetOk("used_time"); ok {
		request["UsedTime"] = v
	}
	if v, ok := d.GetOk("endpoints"); ok {
		endpointsMapsArray := make([]interface{}, 0)
		for _, dataLoop1 := range convertToInterfaceArray(v) {
			dataLoop1Tmp := dataLoop1.(map[string]interface{})
			dataLoop1Map := make(map[string]interface{})
			dataLoop1Map["Description"] = dataLoop1Tmp["description"]
			endpointsMapsArray = append(endpointsMapsArray, dataLoop1Map)
		}
		endpointsMapsJson, err := json.Marshal(endpointsMapsArray)
		if err != nil {
			return WrapError(err)
		}
		request["Endpoints"] = string(endpointsMapsJson)
	}

	if v, ok := d.GetOk("polar_fs_instance_id"); ok {
		request["PolarFSInstanceId"] = v
	}
	request["ApplicationType"] = d.Get("application_type")
	if v, ok := d.GetOkExists("auto_renew"); ok {
		request["AutoRenew"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("promotion_code"); ok {
		request["PromotionCode"] = v
	}
	request["Architecture"] = d.Get("architecture")
	if v, ok := d.GetOkExists("auto_create_polar_fs"); ok {
		request["AutoCreatePolarFs"] = v
	}
	if v, ok := d.GetOkExists("auto_use_coupon"); ok {
		request["AutoUseCoupon"] = v
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		request["VpcId"] = v
	}
	request["DBClusterId"] = d.Get("db_cluster_id")
	if v, ok := d.GetOk("period"); ok {
		request["Period"] = v
	}
	if v, ok := d.GetOk("zone_id"); ok {
		request["ZoneId"] = v
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_polardb_application", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["ApplicationId"]))

	return resourceAliCloudPolardbApplicationUpdate(d, meta)
}

func resourceAliCloudPolardbApplicationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	polardbServiceV2 := PolardbServiceV2{client}

	objectRaw, err := polardbServiceV2.DescribePolardbApplication(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_polardb_application DescribePolardbApplication Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("application_type", objectRaw["ApplicationType"])
	d.Set("architecture", objectRaw["Architecture"])
	d.Set("db_cluster_id", objectRaw["DBClusterId"])
	d.Set("description", objectRaw["Description"])
	d.Set("payment_type", convertPolardbApplicationPayTypeResponse(objectRaw["PayType"]))
	d.Set("polar_fs_instance_id", objectRaw["PolarFSInstanceId"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("status", objectRaw["Status"])
	d.Set("vswitch_id", objectRaw["VSwitchId"])
	d.Set("vpc_id", objectRaw["VPCId"])
	d.Set("zone_id", objectRaw["ZoneId"])

	componentsRaw := objectRaw["Components"]
	componentsMaps := make([]map[string]interface{}, 0)
	if componentsRaw != nil {
		for _, componentsChildRaw := range convertToInterfaceArray(componentsRaw) {
			componentsMap := make(map[string]interface{})
			componentsChildRaw := componentsChildRaw.(map[string]interface{})
			componentsMap["component_class"] = componentsChildRaw["ComponentClass"]
			componentsMap["component_class_description"] = componentsChildRaw["ComponentClassDescription"]
			componentsMap["component_id"] = componentsChildRaw["ComponentId"]
			componentsMap["component_max_replica"] = componentsChildRaw["ComponentMaxReplica"]
			componentsMap["component_replica"] = componentsChildRaw["ComponentReplica"]
			componentsMap["component_replica_group_name"] = componentsChildRaw["ComponentReplicaGroupName"]
			componentsMap["component_type"] = componentsChildRaw["ComponentType"]
			componentsMap["status"] = componentsChildRaw["Status"]

			securityGroupsRaw := componentsChildRaw["SecurityGroups"]
			securityGroupsMaps := make([]map[string]interface{}, 0)
			if securityGroupsRaw != nil {
				for _, securityGroupsChildRaw := range convertToInterfaceArray(securityGroupsRaw) {
					securityGroupsMap := make(map[string]interface{})
					securityGroupsChildRaw := securityGroupsChildRaw.(map[string]interface{})
					securityGroupsMap["net_type"] = securityGroupsChildRaw["NetType"]
					securityGroupsMap["region_id"] = securityGroupsChildRaw["RegionId"]
					securityGroupsMap["security_group_id"] = securityGroupsChildRaw["SecurityGroupId"]
					securityGroupsMap["security_group_name"] = securityGroupsChildRaw["SecurityGroupName"]

					securityGroupsMaps = append(securityGroupsMaps, securityGroupsMap)
				}
			}
			componentsMap["security_groups"] = securityGroupsMaps
			securityIPArraysRaw := componentsChildRaw["SecurityIPArrays"]
			securityIPArraysMaps := make([]map[string]interface{}, 0)
			if securityIPArraysRaw != nil {
				for _, securityIPArraysChildRaw := range convertToInterfaceArray(securityIPArraysRaw) {
					securityIPArraysMap := make(map[string]interface{})
					securityIPArraysChildRaw := securityIPArraysChildRaw.(map[string]interface{})
					securityIPArraysMap["security_ip_array_name"] = securityIPArraysChildRaw["SecurityIPArrayName"]
					securityIPArraysMap["security_ip_array_tag"] = securityIPArraysChildRaw["SecurityIPArrayTag"]
					securityIPArraysMap["security_ip_list"] = securityIPArraysChildRaw["SecurityIPList"]
					securityIPArraysMap["security_ip_net_type"] = securityIPArraysChildRaw["SecurityIPNetType"]
					securityIPArraysMap["security_ip_type"] = securityIPArraysChildRaw["SecurityIPType"]

					securityIPArraysMaps = append(securityIPArraysMaps, securityIPArraysMap)
				}
			}
			componentsMap["security_ip_arrays"] = securityIPArraysMaps
			topologyMaps := make([]map[string]interface{}, 0)
			topologyMap := make(map[string]interface{})
			topologyRaw := make(map[string]interface{})
			if componentsChildRaw["Topology"] != nil {
				topologyRaw = componentsChildRaw["Topology"].(map[string]interface{})
			}
			if len(topologyRaw) > 0 {
				topologyMap["layer"] = topologyRaw["Layer"]

				topologyMap["children"] = childrenRaw

				topologyMap["parents"] = parentsRaw
				topologyMaps = append(topologyMaps, topologyMap)
			}
			componentsMap["topology"] = topologyMaps
			componentsMaps = append(componentsMaps, componentsMap)
		}
	}
	if err := d.Set("components", componentsMaps); err != nil {
		return err
	}
	endpointsRaw := objectRaw["Endpoints"]
	endpointsMaps := make([]map[string]interface{}, 0)
	if endpointsRaw != nil {
		for _, endpointsChildRaw := range convertToInterfaceArray(endpointsRaw) {
			endpointsMap := make(map[string]interface{})
			endpointsChildRaw := endpointsChildRaw.(map[string]interface{})
			endpointsMap["description"] = endpointsChildRaw["Description"]
			endpointsMap["endpoint_id"] = endpointsChildRaw["EndpointId"]
			endpointsMap["ip"] = endpointsChildRaw["IP"]
			endpointsMap["net_type"] = endpointsChildRaw["NetType"]
			endpointsMap["port"] = endpointsChildRaw["Port"]
			endpointsMap["port_description"] = endpointsChildRaw["PortDescription"]

			endpointsMaps = append(endpointsMaps, endpointsMap)
		}
	}
	if err := d.Set("endpoints", endpointsMaps); err != nil {
		return err
	}
	securityGroupsRaw := objectRaw["SecurityGroups"]
	securityGroupsMaps := make([]map[string]interface{}, 0)
	if securityGroupsRaw != nil {
		for _, securityGroupsChildRaw := range convertToInterfaceArray(securityGroupsRaw) {
			securityGroupsMap := make(map[string]interface{})
			securityGroupsChildRaw := securityGroupsChildRaw.(map[string]interface{})
			securityGroupsMap["net_type"] = securityGroupsChildRaw["NetType"]
			securityGroupsMap["region_id"] = securityGroupsChildRaw["RegionId"]
			securityGroupsMap["security_group_id"] = securityGroupsChildRaw["SecurityGroupId"]
			securityGroupsMap["security_group_name"] = securityGroupsChildRaw["SecurityGroupName"]

			securityGroupsMaps = append(securityGroupsMaps, securityGroupsMap)
		}
	}
	if err := d.Set("security_groups", securityGroupsMaps); err != nil {
		return err
	}
	securityIPArraysRaw := objectRaw["SecurityIPArrays"]
	securityIPArraysMaps := make([]map[string]interface{}, 0)
	if securityIPArraysRaw != nil {
		for _, securityIPArraysChildRaw := range convertToInterfaceArray(securityIPArraysRaw) {
			securityIPArraysMap := make(map[string]interface{})
			securityIPArraysChildRaw := securityIPArraysChildRaw.(map[string]interface{})
			securityIPArraysMap["security_ip_array_name"] = securityIPArraysChildRaw["SecurityIPArrayName"]
			securityIPArraysMap["security_ip_array_tag"] = securityIPArraysChildRaw["SecurityIPArrayTag"]
			securityIPArraysMap["security_ip_list"] = securityIPArraysChildRaw["SecurityIPList"]
			securityIPArraysMap["security_ip_net_type"] = securityIPArraysChildRaw["SecurityIPNetType"]
			securityIPArraysMap["security_ip_type"] = securityIPArraysChildRaw["SecurityIPType"]

			securityIPArraysMaps = append(securityIPArraysMaps, securityIPArraysMap)
		}
	}
	if err := d.Set("security_ip_arrays", securityIPArraysMaps); err != nil {
		return err
	}

	objectRaw, err = polardbServiceV2.DescribeApplicationListTagResources(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	tagsMaps := tagResourceChildRaw.(map[string]interface{})
	d.Set("tags", tagsToMap(tagsMaps))

	return nil
}

func resourceAliCloudPolardbApplicationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "ModifyApplicationDescription"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ApplicationId"] = d.Id()

	if !d.IsNewResource() && d.HasChange("description") {
		update = true
	}
	request["Description"] = d.Get("description")
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
	action = "UntagResources"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ResourceId.1"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("tags") {
		update = true
		if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
			tagsMap := ConvertTags(v.(map[string]interface{}))
			request = expandTagsToMap(request, tagsMap)
		}
	}

	request["ResourceType"] = "application"
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
	action = "ModifyApplicationWhitelist"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ApplicationId"] = d.Id()

	if d.HasChange("security_ip_arrays") {
		update = true
		securityIPArraysSecurityIPArrayNameJsonPath, err := jsonpath.Get("$.security_ip_array_name", d.Get("security_ip_arrays"))
		if err == nil {
			request["SecurityIPArrayName"] = securityIPArraysSecurityIPArrayNameJsonPath
		}
	}

	if !d.IsNewResource() && d.HasChange("security_groups") {
		update = true
		securityGroupsSecurityGroupIdJsonPath, err := jsonpath.Get("$.security_group_id", d.Get("security_groups"))
		if err == nil {
			request["SecurityGroups"] = securityGroupsSecurityGroupIdJsonPath
		}
	}

	if v, ok := d.GetOk("modify_mode"); ok {
		request["ModifyMode"] = v
	}
	if d.HasChange("security_ip_arrays") {
		update = true
		securityIPArraysSecurityIPListJsonPath, err := jsonpath.Get("$.security_ip_list", d.Get("security_ip_arrays"))
		if err == nil {
			request["SecurityIPList"] = securityIPArraysSecurityIPListJsonPath
		}
	}

	if d.HasChange("components") {
		update = true
		componentsComponentIdJsonPath, err := jsonpath.Get("$.component_id", d.Get("components"))
		if err == nil {
			request["ComponentId"] = componentsComponentIdJsonPath
		}
	}

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
	action = "CreateApplicationEndpointAddress"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ApplicationId"] = d.Id()

	if d.HasChange("endpoints") {
		update = true
	}
	endpointsEndpointIdJsonPath, err := jsonpath.Get("$.endpoint_id", d.Get("endpoints"))
	if err == nil {
		request["EndpointId"] = endpointsEndpointIdJsonPath
	}

	if d.HasChange("endpoints") {
		update = true
	}
	endpointsNetTypeJsonPath, err := jsonpath.Get("$.net_type", d.Get("endpoints"))
	if err == nil {
		request["NetType"] = endpointsNetTypeJsonPath
	}

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

	d.Partial(false)
	return resourceAliCloudPolardbApplicationRead(d, meta)
}

func resourceAliCloudPolardbApplicationDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	enableDelete := false
	if v, ok := d.GetOkExists("payment_type"); ok {
		if InArray(fmt.Sprint(v), []string{"PayAsYouGo"}) {
			enableDelete = true
		}
	}
	if enableDelete {
		action := "DeleteApplication"
		var request map[string]interface{}
		var response map[string]interface{}
		query := make(map[string]interface{})
		var err error
		request = make(map[string]interface{})
		request["ApplicationId"] = d.Id()

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
			if IsExpectedErrors(err, []string{"InvalidApplicationId.NotFound"}) || NotFoundError(err) {
				return nil
			}
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}

	}
	return nil
}

func convertPolardbApplicationPayTypeResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "Postpaid":
		return "PayAsYouGo"
	case "Prepaid":
		return "Subscription"
	}
	return source
}

func convertPolardbApplicationPayTypeRequest(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "PayAsYouGo":
		return "Postpaid"
	case "Subscription":
		return "Prepaid"
	}
	return source
}
