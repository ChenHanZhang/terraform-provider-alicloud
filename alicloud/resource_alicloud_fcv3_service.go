package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAlicloudFcv3Service() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudFcv3ServiceCreate,
		Read:   resourceAlicloudFcv3ServiceRead,
		Update: resourceAlicloudFcv3ServiceUpdate,
		Delete: resourceAlicloudFcv3ServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_access": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"vpc_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"vswitch_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"security_group_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"nas_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"group_id": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mount_points": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server_addr": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"mount_dir": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"log_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"logstore": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"log_begin_rule": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"enable_request_metrics": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"enable_instance_metrics": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"tags": tagsSchema(),
			"created_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_modified_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAlicloudFcv3ServiceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/2023-03-30/services")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("service_name"); ok {
		request["serviceName"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["description"] = v
	}
	if v, ok := d.GetOk("role"); ok {
		request["role"] = v
	}
	if v, ok := d.GetOkExists("internet_access"); ok {
		request["internetAccess"] = v
	}

	vpcConfig := make(map[string]interface{})
	if v := d.Get("vpc_config"); !IsNil(v) {
		vpcId1, err := GetInterfaceItemFromList(v, 0, "vpc_id")
		if err != nil {
			return WrapError(err)
		}
		if vpcId1 != nil && vpcId1.(string) != "" {
			vpcConfig["vpcId"] = vpcId1
		}
		vswitchIds1, err := GetInterfaceItemFromList(v, 0, "vswitch_ids")
		if err != nil {
			return WrapError(err)
		}
		if vswitchIds1 != nil {
			vpcConfig["vswitchIds"] = vswitchIds1
		}
		securityGroupId1, err := GetInterfaceItemFromList(v, 0, "security_group_id")
		if err != nil {
			return WrapError(err)
		}
		if securityGroupId1 != nil && securityGroupId1.(string) != "" {
			vpcConfig["securityGroupId"] = securityGroupId1
		}

		request["vpcConfig"] = vpcConfig
	}

	nasConfig := make(map[string]interface{})
	if v := d.Get("nas_config"); !IsNil(v) {
		userId1, err := GetInterfaceItemFromList(v, 0, "user_id")
		if err != nil {
			return WrapError(err)
		}
		if userId1 != nil && userId1.(int) != 0 {
			nasConfig["userId"] = userId1
		}
		groupId1, err := GetInterfaceItemFromList(v, 0, "group_id")
		if err != nil {
			return WrapError(err)
		}
		if groupId1 != nil && groupId1.(int) != 0 {
			nasConfig["groupId"] = groupId1
		}

		request["nasConfig"] = nasConfig
	}

	logConfig := make(map[string]interface{})
	if v := d.Get("log_config"); !IsNil(v) {
		project1, err := GetInterfaceItemFromList(v, 0, "project")
		if err != nil {
			return WrapError(err)
		}
		if project1 != nil && project1.(string) != "" {
			logConfig["project"] = project1
		}
		logstore1, err := GetInterfaceItemFromList(v, 0, "logstore")
		if err != nil {
			return WrapError(err)
		}
		if logstore1 != nil && logstore1.(string) != "" {
			logConfig["logstore"] = logstore1
		}
		logBeginRule1, err := GetInterfaceItemFromList(v, 0, "log_begin_rule")
		if err != nil {
			return WrapError(err)
		}
		if logBeginRule1 != nil && logBeginRule1.(string) != "" {
			logConfig["logBeginRule"] = logBeginRule1
		}
		enableRequestMetrics1, err := GetInterfaceItemFromList(v, 0, "enable_request_metrics")
		if err != nil {
			return WrapError(err)
		}
		if enableRequestMetrics1 != nil {
			logConfig["enableRequestMetrics"] = enableRequestMetrics1
		}
		enableInstanceMetrics1, err := GetInterfaceItemFromList(v, 0, "enable_instance_metrics")
		if err != nil {
			return WrapError(err)
		}
		if enableInstanceMetrics1 != nil {
			logConfig["enableInstanceMetrics"] = enableInstanceMetrics1
		}

		request["logConfig"] = logConfig
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RoaPost("FC", "2023-03-30", action, query, nil, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, query, request)
		return nil
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_fcv3_service", action, AlibabaCloudSdkGoERROR)
	}

	id := fmt.Sprint(request["serviceName"])
	d.SetId(id)

	if d.Get("tags").(map[string]interface{}) != nil && len(d.Get("tags").(map[string]interface{})) > 0 {
		fcv3ServiceV2 := Fcv3ServiceV2{client}
		err := fcv3ServiceV2.SetResourceTags(d, "service")
		if err != nil {
			return WrapError(err)
		}
	}

	return resourceAlicloudFcv3ServiceRead(d, meta)
}

func resourceAlicloudFcv3ServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	fcv3ServiceV2 := Fcv3ServiceV2{client}

	object, err := fcv3ServiceV2.DescribeFcv3Service(d.Id())
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("service_name", object["serviceName"])
	d.Set("description", object["description"])
	d.Set("role", object["role"])
	d.Set("internet_access", object["internetAccess"])
	d.Set("created_time", object["createdTime"])
	d.Set("last_modified_time", object["lastModifiedTime"])
	d.Set("service_id", object["serviceId"])
	d.Set("service_arn", object["serviceArn"])

	if v, ok := object["vpcConfig"]; ok {
		vpcConfig := v.(map[string]interface{})
		var vpcConfigs []map[string]interface{}
		mapping := map[string]interface{}{
			"vpc_id":            vpcConfig["vpcId"],
			"vswitch_ids":       vpcConfig["vswitchIds"],
			"security_group_id": vpcConfig["securityGroupId"],
		}
		vpcConfigs = append(vpcConfigs, mapping)
		if err := d.Set("vpc_config", vpcConfigs); err != nil {
			return WrapError(err)
		}
	}

	if v, ok := object["nasConfig"]; ok {
		nasConfig := v.(map[string]interface{})
		var nasConfigs []map[string]interface{}
		mapping := map[string]interface{}{
			"user_id":      nasConfig["userId"],
			"group_id":     nasConfig["groupId"],
			"mount_points": nasConfig["mountPoints"],
		}
		nasConfigs = append(nasConfigs, mapping)
		if err := d.Set("nas_config", nasConfigs); err != nil {
			return WrapError(err)
		}
	}

	if v, ok := object["logConfig"]; ok {
		logConfig := v.(map[string]interface{})
		var logConfigs []map[string]interface{}
		mapping := map[string]interface{}{
			"project":                 logConfig["project"],
			"logstore":                logConfig["logstore"],
			"log_begin_rule":          logConfig["logBeginRule"],
			"enable_request_metrics":  logConfig["enableRequestMetrics"],
			"enable_instance_metrics": logConfig["enableInstanceMetrics"],
		}
		logConfigs = append(logConfigs, mapping)
		if err := d.Set("log_config", logConfigs); err != nil {
			return WrapError(err)
		}
	}

	// Get tags
	tags, err := fcv3ServiceV2.DescribeFcv3ServiceTags(d.Id())
	if err != nil {
		log.Printf("[ERROR] DescribeFcv3ServiceTags error: %s", err.Error())
	} else {
		d.Set("tags", tags)
	}

	return nil
}

func resourceAlicloudFcv3ServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	if d.HasChange("tags") {
		fcv3ServiceV2 := Fcv3ServiceV2{client}
		err := fcv3ServiceV2.SetResourceTags(d, "service")
		if err != nil {
			return WrapError(err)
		}
	}

	update := false
	request := make(map[string]interface{})
	if d.HasChange("description") {
		update = true
		request["description"] = d.Get("description").(string)
	}
	if d.HasChange("role") {
		update = true
		request["role"] = d.Get("role").(string)
	}
	if d.HasChange("internet_access") {
		update = true
		request["internetAccess"] = d.Get("internet_access").(bool)
	}

	if d.HasChange("vpc_config") {
		update = true
		vpcConfig := make(map[string]interface{})
		if v := d.Get("vpc_config"); !IsNil(v) {
			vpcId1, err := GetInterfaceItemFromList(v, 0, "vpc_id")
			if err != nil {
				return WrapError(err)
			}
			if vpcId1 != nil && vpcId1.(string) != "" {
				vpcConfig["vpcId"] = vpcId1
			}
			vswitchIds1, err := GetInterfaceItemFromList(v, 0, "vswitch_ids")
			if err != nil {
				return WrapError(err)
			}
			if vswitchIds1 != nil {
				vpcConfig["vswitchIds"] = vswitchIds1
			}
			securityGroupId1, err := GetInterfaceItemFromList(v, 0, "security_group_id")
			if err != nil {
				return WrapError(err)
			}
			if securityGroupId1 != nil && securityGroupId1.(string) != "" {
				vpcConfig["securityGroupId"] = securityGroupId1
			}
		}
		request["vpcConfig"] = vpcConfig
	}

	if d.HasChange("nas_config") {
		update = true
		nasConfig := make(map[string]interface{})
		if v := d.Get("nas_config"); !IsNil(v) {
			userId1, err := GetInterfaceItemFromList(v, 0, "user_id")
			if err != nil {
				return WrapError(err)
			}
			if userId1 != nil && userId1.(int) != 0 {
				nasConfig["userId"] = userId1
			}
			groupId1, err := GetInterfaceItemFromList(v, 0, "group_id")
			if err != nil {
				return WrapError(err)
			}
			if groupId1 != nil && groupId1.(int) != 0 {
				nasConfig["groupId"] = groupId1
			}
		}
		request["nasConfig"] = nasConfig
	}

	if d.HasChange("log_config") {
		update = true
		logConfig := make(map[string]interface{})
		if v := d.Get("log_config"); !IsNil(v) {
			project1, err := GetInterfaceItemFromList(v, 0, "project")
			if err != nil {
				return WrapError(err)
			}
			if project1 != nil && project1.(string) != "" {
				logConfig["project"] = project1
			}
			logstore1, err := GetInterfaceItemFromList(v, 0, "logstore")
			if err != nil {
				return WrapError(err)
			}
			if logstore1 != nil && logstore1.(string) != "" {
				logConfig["logstore"] = logstore1
			}
			logBeginRule1, err := GetInterfaceItemFromList(v, 0, "log_begin_rule")
			if err != nil {
				return WrapError(err)
			}
			if logBeginRule1 != nil && logBeginRule1.(string) != "" {
				logConfig["logBeginRule"] = logBeginRule1
			}
			enableRequestMetrics1, err := GetInterfaceItemFromList(v, 0, "enable_request_metrics")
			if err != nil {
				return WrapError(err)
			}
			if enableRequestMetrics1 != nil {
				logConfig["enableRequestMetrics"] = enableRequestMetrics1
			}
			enableInstanceMetrics1, err := GetInterfaceItemFromList(v, 0, "enable_instance_metrics")
			if err != nil {
				return WrapError(err)
			}
			if enableInstanceMetrics1 != nil {
				logConfig["enableInstanceMetrics"] = enableInstanceMetrics1
			}
		}
		request["logConfig"] = logConfig
	}

	if update {
		action := fmt.Sprintf("/2023-03-30/services/%s", d.Id())
		var response map[string]interface{}
		var err error
		query := make(map[string]*string)

		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPut("FC", "2023-03-30", action, query, nil, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, query, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}

	return resourceAlicloudFcv3ServiceRead(d, meta)
}

func resourceAlicloudFcv3ServiceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/2023-03-30/services/%s", d.Id())
	var request map[string]interface{}
	query := make(map[string]*string)
	request = make(map[string]interface{})

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err := resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, requestErr := client.RoaDelete("FC", "2023-03-30", action, query, nil, request, true)
		if requestErr != nil {
			if NeedRetry(requestErr) {
				wait()
				return resource.RetryableError(requestErr)
			}
			return resource.NonRetryableError(requestErr)
		}
		addDebug(action, response, query, request)
		return nil
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}
	return nil
}
