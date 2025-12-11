// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/blues/jsonata-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudAliKafkaInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAliKafkaInstanceCreate,
		Read:   resourceAliCloudAliKafkaInstanceRead,
		Update: resourceAliCloudAliKafkaInstanceUpdate,
		Delete: resourceAliCloudAliKafkaInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(33 * time.Minute),
			Update: schema.DefaultTimeout(55 * time.Minute),
			Delete: schema.DefaultTimeout(55 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"config": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"confluent_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ksql_storage": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"kafka_rest_proxy_cu": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"zoo_keeper_replica": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: IntInSlice([]int{0, 3}),
						},
						"kafka_replica": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"connect_replica": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"control_center_storage": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"schema_registry_cu": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"kafka_cu": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"connect_cu": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ksql_cu": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"control_center_replica": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: IntInSlice([]int{0, 1}),
						},
						"ksql_replica": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"schema_registry_replica": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"kafka_rest_proxy_replica": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"zoo_keeper_cu": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"zoo_keeper_storage": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"control_center_cu": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"kafka_storage": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cross_zone": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default_topic_partition_num": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"deploy_module": {
				Type:     schema.TypeString,
				Required: true,
			},
			"deploy_type": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: IntInSlice([]int{0, 4, 5}),
			},
			"disk_size": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"disk_type": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"domain_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"duration": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"eip_max": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"eip_model": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_auto_group": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_auto_topic": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_point": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_left": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"group_used": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"instance_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"io_max_spec": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_eip_inner": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_force_selected_zones": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_partition_buy": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_set_user_and_password": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"notifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"order_id": {
				Type:      schema.TypeString,
				Optional:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"paid_type": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntInSlice([]int{0, 1, 3, 4}),
			},
			"partition_left": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"partition_num_of_buy": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"partition_used": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"password": {
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
			"sasl_domain_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_group": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"selected_zones": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"serverless_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reserved_publish_capacity": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"reserved_subscribe_capacity": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"service_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spec_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_domain_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": tagsSchema(),
			"topic_left": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"topic_num_of_buy": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"topic_used": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"update_default_topic_partition_num": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"user_phone_num": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vswitch_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vswitch_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceAliCloudAliKafkaInstanceCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	if v, ok := d.GetOk("paid_type"); ok && InArray(fmt.Sprint(v), []string{"1", "3"}) {
		action := "CreatePostPayInstance"
		var request map[string]interface{}
		var response map[string]interface{}
		query := make(map[string]interface{})
		var err error
		request = make(map[string]interface{})
		request["RegionId"] = client.RegionId

		if v, ok := d.GetOkExists("disk_type"); ok {
			request["DiskType"] = v
		}
		request["DeployType"] = d.Get("deploy_type")
		if v, ok := d.GetOk("resource_group_id"); ok {
			request["ResourceGroupId"] = v
		}
		if v, ok := d.GetOkExists("paid_type"); ok {
			request["PaidType"] = v
		}
		if v, ok := d.GetOk("io_max_spec"); ok {
			request["IoMaxSpec"] = v
		}
		if v, ok := d.GetOkExists("disk_size"); ok {
			request["DiskSize"] = v
		}
		serverlessConfig := make(map[string]interface{})

		if v := d.Get("serverless_config"); !IsNil(v) {
			reservedSubscribeCapacity1, _ := jsonpath.Get("$[0].reserved_subscribe_capacity", v)
			if reservedSubscribeCapacity1 != nil && reservedSubscribeCapacity1 != "" {
				serverlessConfig["ReservedSubscribeCapacity"] = reservedSubscribeCapacity1
			}
			reservedPublishCapacity1, _ := jsonpath.Get("$[0].reserved_publish_capacity", v)
			if reservedPublishCapacity1 != nil && reservedPublishCapacity1 != "" {
				serverlessConfig["ReservedPublishCapacity"] = reservedPublishCapacity1
			}

			serverlessConfigJson, err := json.Marshal(serverlessConfig)
			if err != nil {
				return WrapError(err)
			}
			request["ServerlessConfig"] = string(serverlessConfigJson)
		}

		if v, ok := d.GetOkExists("eip_max"); ok {
			request["EipMax"] = v
		}
		if v, ok := d.GetOk("spec_type"); ok {
			request["SpecType"] = v
		}
		if v, ok := d.GetOk("tags"); ok {
			tagsMap := ConvertTags(v.(map[string]interface{}))
			request = expandTagsToMap(request, tagsMap)
		}

		if v, ok := d.GetOkExists("partition_num_of_buy"); ok {
			request["PartitionNum"] = v
		}
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
			return WrapErrorf(err, DefaultErrorMsg, "alicloud_alikafka_instance_v2", action, AlibabaCloudSdkGoERROR)
		}

		id, _ := jsonpath.Get("$.Data.InstanceId", response)
		d.SetId(fmt.Sprint(id))

	}

	if v, ok := d.GetOk("paid_type"); ok && InArray(fmt.Sprint(v), []string{"0", "4"}) {
		action := "CreatePrePayInstance"
		var request map[string]interface{}
		var response map[string]interface{}
		query := make(map[string]interface{})
		var err error
		request = make(map[string]interface{})
		request["RegionId"] = client.RegionId

		if v, ok := d.GetOkExists("duration"); ok {
			request["Duration"] = v
		}
		if v, ok := d.GetOkExists("disk_type"); ok {
			request["DiskType"] = v
		}
		confluentConfig := make(map[string]interface{})

		if v := d.Get("confluent_config"); !IsNil(v) {
			zooKeeperReplica1, _ := jsonpath.Get("$[0].zoo_keeper_replica", v)
			if zooKeeperReplica1 != nil && zooKeeperReplica1 != "" {
				confluentConfig["ZooKeeperReplica"] = zooKeeperReplica1
			}
			kafkaStorage1, _ := jsonpath.Get("$[0].kafka_storage", v)
			if kafkaStorage1 != nil && kafkaStorage1 != "" {
				confluentConfig["KafkaStorage"] = kafkaStorage1
			}
			schemaRegistryReplica1, _ := jsonpath.Get("$[0].schema_registry_replica", v)
			if schemaRegistryReplica1 != nil && schemaRegistryReplica1 != "" {
				confluentConfig["SchemaRegistryReplica"] = schemaRegistryReplica1
			}
			controlCenterReplica1, _ := jsonpath.Get("$[0].control_center_replica", v)
			if controlCenterReplica1 != nil && controlCenterReplica1 != "" {
				confluentConfig["ControlCenterReplica"] = controlCenterReplica1
			}
			zooKeeperStorage1, _ := jsonpath.Get("$[0].zoo_keeper_storage", v)
			if zooKeeperStorage1 != nil && zooKeeperStorage1 != "" {
				confluentConfig["ZooKeeperStorage"] = zooKeeperStorage1
			}
			ksqlReplica1, _ := jsonpath.Get("$[0].ksql_replica", v)
			if ksqlReplica1 != nil && ksqlReplica1 != "" {
				confluentConfig["KsqlReplica"] = ksqlReplica1
			}
			ksqlStorage1, _ := jsonpath.Get("$[0].ksql_storage", v)
			if ksqlStorage1 != nil && ksqlStorage1 != "" {
				confluentConfig["KsqlStorage"] = ksqlStorage1
			}
			connectReplica1, _ := jsonpath.Get("$[0].connect_replica", v)
			if connectReplica1 != nil && connectReplica1 != "" {
				confluentConfig["ConnectReplica"] = connectReplica1
			}
			controlCenterStorage1, _ := jsonpath.Get("$[0].control_center_storage", v)
			if controlCenterStorage1 != nil && controlCenterStorage1 != "" {
				confluentConfig["ControlCenterStorage"] = controlCenterStorage1
			}
			kafkaRestProxyReplica1, _ := jsonpath.Get("$[0].kafka_rest_proxy_replica", v)
			if kafkaRestProxyReplica1 != nil && kafkaRestProxyReplica1 != "" {
				confluentConfig["KafkaRestProxyReplica"] = kafkaRestProxyReplica1
			}
			kafkaReplica1, _ := jsonpath.Get("$[0].kafka_replica", v)
			if kafkaReplica1 != nil && kafkaReplica1 != "" {
				confluentConfig["KafkaReplica"] = kafkaReplica1
			}
			kafkaCu, _ := jsonpath.Get("$[0].kafka_cu", v)
			if kafkaCu != nil && kafkaCu != "" {
				confluentConfig["KafkaCU"] = kafkaCu
			}
			zooKeeperCu, _ := jsonpath.Get("$[0].zoo_keeper_cu", v)
			if zooKeeperCu != nil && zooKeeperCu != "" {
				confluentConfig["ZooKeeperCU"] = zooKeeperCu
			}
			controlCenterCu, _ := jsonpath.Get("$[0].control_center_cu", v)
			if controlCenterCu != nil && controlCenterCu != "" {
				confluentConfig["ControlCenterCU"] = controlCenterCu
			}
			schemaRegistryCu, _ := jsonpath.Get("$[0].schema_registry_cu", v)
			if schemaRegistryCu != nil && schemaRegistryCu != "" {
				confluentConfig["SchemaRegistryCU"] = schemaRegistryCu
			}
			connectCu, _ := jsonpath.Get("$[0].connect_cu", v)
			if connectCu != nil && connectCu != "" {
				confluentConfig["ConnectCU"] = connectCu
			}
			ksqlCu, _ := jsonpath.Get("$[0].ksql_cu", v)
			if ksqlCu != nil && ksqlCu != "" {
				confluentConfig["KsqlCU"] = ksqlCu
			}
			kafkaRestProxyCu, _ := jsonpath.Get("$[0].kafka_rest_proxy_cu", v)
			if kafkaRestProxyCu != nil && kafkaRestProxyCu != "" {
				confluentConfig["KafkaRestProxyCU"] = kafkaRestProxyCu
			}

			confluentConfigJson, err := json.Marshal(confluentConfig)
			if err != nil {
				return WrapError(err)
			}
			request["ConfluentConfig"] = string(confluentConfigJson)
		}

		request["DeployType"] = d.Get("deploy_type")
		if v, ok := d.GetOk("resource_group_id"); ok {
			request["ResourceGroupId"] = v
		}
		if v, ok := d.GetOkExists("paid_type"); ok {
			request["PaidType"] = v
		}
		if v, ok := d.GetOk("io_max_spec"); ok {
			request["IoMaxSpec"] = v
		}
		if v, ok := d.GetOkExists("disk_size"); ok {
			request["DiskSize"] = v
		}
		if v, ok := d.GetOkExists("eip_max"); ok {
			request["EipMax"] = v
		}
		if v, ok := d.GetOk("spec_type"); ok {
			request["SpecType"] = v
		}
		if v, ok := d.GetOk("tags"); ok {
			tagsMap := ConvertTags(v.(map[string]interface{}))
			request = expandTagsToMap(request, tagsMap)
		}

		if v, ok := d.GetOkExists("default_topic_partition_num"); ok {
			request["PartitionNum"] = v
		}
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
			return WrapErrorf(err, DefaultErrorMsg, "alicloud_alikafka_instance_v2", action, AlibabaCloudSdkGoERROR)
		}

		id, _ := jsonpath.Get("$.Data.InstanceId", response)
		d.SetId(fmt.Sprint(id))

	}

	action := "StartInstance"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	request["RegionId"] = client.RegionId

	request["VpcId"] = d.Get("vpc_id")
	request["VSwitchId"] = d.Get("vswitch_id")
	request["DeployModule"] = d.Get("deploy_module")
	if v, ok := d.GetOkExists("is_eip_inner"); ok {
		request["IsEipInner"] = v
	}
	if v, ok := d.GetOkExists("is_set_user_and_password"); ok {
		request["IsSetUserAndPassword"] = v
	}
	if v, ok := d.GetOk("username"); ok {
		request["Username"] = v
	}
	if v, ok := d.GetOk("password"); ok {
		request["Password"] = v
	}
	if v, ok := d.GetOk("security_group"); ok {
		request["SecurityGroup"] = v
	}
	if v, ok := d.GetOk("service_version"); ok {
		request["ServiceVersion"] = v
	}
	if v, ok := d.GetOk("config"); ok {
		request["Config"] = v
	}
	if v, ok := d.GetOk("notifier"); ok {
		request["Notifier"] = v
	}
	if v, ok := d.GetOk("user_phone_num"); ok {
		request["UserPhoneNum"] = v
	}
	if v, ok := d.GetOkExists("is_force_selected_zones"); ok {
		request["IsForceSelectedZones"] = v
	}
	if v, ok := d.GetOk("kms_key_id"); ok {
		request["KMSKeyId"] = v
	}
	if v, ok := d.GetOkExists("cross_zone"); ok {
		request["CrossZone"] = v
	}
	if v, ok := d.GetOk("zone_id"); ok {
		request["ZoneId"] = v
	}
	if v, ok := d.GetOk("vswitch_ids"); ok {
		vSwitchIdsMapsArray := convertToInterfaceArray(v)

		request["VSwitchIds"] = vSwitchIdsMapsArray
	}

	if v, ok := d.GetOk("selected_zones"); ok {
		request["SelectedZones"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"Instance.NotFound", "AUTH_RESOURCE_OWNER_ERROR", "VpcVSwitch.NotFound"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_alikafka_instance_v2", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["InstanceId"]))

	aliKafkaServiceV2 := AliKafkaServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"5"}, d.Timeout(schema.TimeoutCreate), 3*time.Minute, aliKafkaServiceV2.AliKafkaInstanceStateRefreshFunc(d.Id(), "ServiceStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudAliKafkaInstanceUpdate(d, meta)
}

func resourceAliCloudAliKafkaInstanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	aliKafkaServiceV2 := AliKafkaServiceV2{client}

	objectRaw, err := aliKafkaServiceV2.DescribeAliKafkaInstance(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_alikafka_instance_v2 DescribeAliKafkaInstance Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("config", objectRaw["AllConfig"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("default_topic_partition_num", objectRaw["DefaultPartitionNum"])
	d.Set("deploy_type", objectRaw["DeployType"])
	d.Set("domain_endpoint", objectRaw["DomainEndpoint"])
	d.Set("eip_max", objectRaw["EipMax"])
	d.Set("enable_auto_group", objectRaw["AutoCreateGroupEnable"])
	d.Set("enable_auto_topic", convertAliKafkaInstanceInstanceListInstanceVOAutoCreateTopicEnableResponse(objectRaw["AutoCreateTopicEnable"]))
	d.Set("end_point", objectRaw["EndPoint"])
	d.Set("instance_name", objectRaw["Name"])
	d.Set("io_max_spec", objectRaw["IoMaxSpec"])
	d.Set("kms_key_id", objectRaw["KmsKeyId"])
	d.Set("paid_type", objectRaw["PaidType"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("resource_group_id", objectRaw["ResourceGroupId"])
	d.Set("sasl_domain_endpoint", objectRaw["SaslDomainEndpoint"])
	d.Set("security_group", objectRaw["SecurityGroup"])
	d.Set("spec_type", objectRaw["SpecType"])
	d.Set("ssl_domain_endpoint", objectRaw["SslDomainEndpoint"])
	d.Set("status", objectRaw["ServiceStatus"])
	d.Set("vpc_id", objectRaw["VpcId"])
	d.Set("zone_id", objectRaw["StandardZoneId"])
	d.Set("instance_id", objectRaw["InstanceId"])

	upgradeServiceDetailInfoRawObj, _ := jsonpath.Get("$.UpgradeServiceDetailInfo", objectRaw)
	upgradeServiceDetailInfoRaw := make(map[string]interface{})
	if upgradeServiceDetailInfoRawObj != nil {
		upgradeServiceDetailInfoRaw = upgradeServiceDetailInfoRawObj.(map[string]interface{})
	}
	d.Set("service_version", upgradeServiceDetailInfoRaw["Current2OpenSourceVersion"])

	confluentConfigMaps := make([]map[string]interface{}, 0)
	confluentConfigMap := make(map[string]interface{})
	confluentConfigRaw := make(map[string]interface{})
	if objectRaw["ConfluentConfig"] != nil {
		confluentConfigRaw = objectRaw["ConfluentConfig"].(map[string]interface{})
	}
	if len(confluentConfigRaw) > 0 {
		confluentConfigMap["connect_cu"] = confluentConfigRaw["ConnectCU"]
		confluentConfigMap["connect_replica"] = confluentConfigRaw["ConnectReplica"]
		confluentConfigMap["control_center_cu"] = confluentConfigRaw["ControlCenterCU"]
		confluentConfigMap["control_center_replica"] = confluentConfigRaw["ControlCenterReplica"]
		confluentConfigMap["control_center_storage"] = confluentConfigRaw["ControlCenterStorage"]
		confluentConfigMap["kafka_cu"] = confluentConfigRaw["KafkaCU"]
		confluentConfigMap["kafka_replica"] = confluentConfigRaw["KafkaReplica"]
		confluentConfigMap["kafka_rest_proxy_cu"] = confluentConfigRaw["KafkaRestProxyCU"]
		confluentConfigMap["kafka_rest_proxy_replica"] = confluentConfigRaw["KafkaRestProxyReplica"]
		confluentConfigMap["kafka_storage"] = confluentConfigRaw["KafkaStorage"]
		confluentConfigMap["ksql_cu"] = confluentConfigRaw["KsqlCU"]
		confluentConfigMap["ksql_replica"] = confluentConfigRaw["KsqlReplica"]
		confluentConfigMap["ksql_storage"] = confluentConfigRaw["KsqlStorage"]
		confluentConfigMap["schema_registry_cu"] = confluentConfigRaw["SchemaRegistryCU"]
		confluentConfigMap["schema_registry_replica"] = confluentConfigRaw["SchemaRegistryReplica"]
		confluentConfigMap["zoo_keeper_cu"] = confluentConfigRaw["ZooKeeperCU"]
		confluentConfigMap["zoo_keeper_replica"] = confluentConfigRaw["ZooKeeperReplica"]
		confluentConfigMap["zoo_keeper_storage"] = confluentConfigRaw["ZooKeeperStorage"]

		confluentConfigMaps = append(confluentConfigMaps, confluentConfigMap)
	}
	if err := d.Set("confluent_config", confluentConfigMaps); err != nil {
		return err
	}
	serverlessConfigMaps := make([]map[string]interface{}, 0)
	serverlessConfigMap := make(map[string]interface{})

	serverlessConfigMap["reserved_publish_capacity"] = objectRaw["ReservedPublishCapacity"]
	serverlessConfigMap["reserved_subscribe_capacity"] = objectRaw["ReservedSubscribeCapacity"]

	serverlessConfigMaps = append(serverlessConfigMaps, serverlessConfigMap)
	if err := d.Set("serverless_config", serverlessConfigMaps); err != nil {
		return err
	}
	tagsMaps, _ := jsonpath.Get("$.Tags.TagVO", objectRaw)
	d.Set("tags", tagsToMap(tagsMaps))
	vSwitchIdsRaw, _ := jsonpath.Get("$.VSwitchIds.VSwitchIds", objectRaw)
	d.Set("vswitch_ids", vSwitchIdsRaw)

	e := jsonata.MustCompile("$.DiskSize != null ? $.DiskSize  : 500")
	evaluation, _ := e.Eval(objectRaw)
	d.Set("disk_size", evaluation)
	e = jsonata.MustCompile("$.DiskType != null ? $.DiskType  : 0")
	evaluation, _ = e.Eval(objectRaw)
	d.Set("disk_type", evaluation)
	e = jsonata.MustCompile("$.PaidType = 4 ? $map($split($substring($.VSwitchId, 0, $length($.VSwitchId)-2), ','), function($v, $i, $a) {$string($substring($v, 2, $length($v)-0))})  : $.VSwitchId")
	evaluation, _ = e.Eval(objectRaw)
	d.Set("vswitch_id", evaluation)
	checkValue00 := d.Get("paid_type")
	if InArray(fmt.Sprint(checkValue00), []string{"0", "1"}) {
		objectRaw, err = aliKafkaServiceV2.DescribeInstanceGetQuotaTip(d.Id())
		if err != nil && !NotFoundError(err) {
			return WrapError(err)
		}

		d.Set("group_left", objectRaw["GroupLeft"])
		d.Set("group_used", objectRaw["GroupUsed"])
		d.Set("is_partition_buy", objectRaw["IsPartitionBuy"])
		d.Set("partition_left", objectRaw["PartitionLeft"])
		d.Set("partition_num_of_buy", objectRaw["PartitionNumOfBuy"])
		d.Set("partition_used", objectRaw["PartitionUsed"])
		d.Set("topic_left", objectRaw["TopicLeft"])
		d.Set("topic_num_of_buy", objectRaw["TopicNumOfBuy"])
		d.Set("topic_used", objectRaw["TopicUsed"])

	}

	d.Set("instance_id", d.Id())

	return nil
}

func resourceAliCloudAliKafkaInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	aliKafkaServiceV2 := AliKafkaServiceV2{client}
	objectRaw, _ := aliKafkaServiceV2.DescribeAliKafkaInstance(d.Id())

	if d.HasChange("status") {
		var err error
		target := d.Get("status").(string)

		currentStatus, err := jsonpath.Get("ServiceStatus", objectRaw)
		if err != nil {
			return WrapErrorf(err, FailedGetAttributeMsg, d.Id(), "ServiceStatus", objectRaw)
		}
		if fmt.Sprint(currentStatus) != target {
			if target == "15" {
				action := "StopInstance"
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				request["InstanceId"] = d.Id()
				request["RegionId"] = client.RegionId
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
				aliKafkaServiceV2 := AliKafkaServiceV2{client}
				stateConf := BuildStateConf([]string{}, []string{"15"}, d.Timeout(schema.TimeoutUpdate), 5*time.Minute, aliKafkaServiceV2.AliKafkaInstanceStateRefreshFunc(d.Id(), "ServiceStatus", []string{}))
				if _, err := stateConf.WaitForState(); err != nil {
					return WrapErrorf(err, IdMsg, d.Id())
				}

			}
			if target == "5" {
				action := "ReopenInstance"
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				request["InstanceId"] = d.Id()
				request["RegionId"] = client.RegionId
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
				aliKafkaServiceV2 := AliKafkaServiceV2{client}
				stateConf := BuildStateConf([]string{}, []string{"5"}, d.Timeout(schema.TimeoutUpdate), 5*time.Minute, aliKafkaServiceV2.AliKafkaInstanceStateRefreshFunc(d.Id(), "ServiceStatus", []string{}))
				if _, err := stateConf.WaitForState(); err != nil {
					return WrapErrorf(err, IdMsg, d.Id())
				}

			}
		}
	}

	var err error
	action := "ModifyInstanceName"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("instance_name") {
		update = true
	}
	request["InstanceName"] = d.Get("instance_name")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
	objectRaw, _ = aliKafkaServiceV2.DescribeAliKafkaInstance(d.Id())
	enableUpgradePostPayOrder1 := false
	checkValue00 := objectRaw["PaidType"]
	if InArray(fmt.Sprint(checkValue00), []string{"1", "3"}) {
		enableUpgradePostPayOrder1 = true
	}
	action = "UpgradePostPayOrder"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("disk_size") {
		update = true
		request["DiskSize"] = d.Get("disk_size")
	}

	if !d.IsNewResource() && d.HasChange("partition_num_of_buy") {
		update = true
		request["PartitionNum"] = d.Get("partition_num_of_buy")
	}

	if !d.IsNewResource() && d.HasChange("io_max_spec") {
		update = true
		request["IoMaxSpec"] = d.Get("io_max_spec")
	}

	if !d.IsNewResource() && d.HasChange("spec_type") {
		update = true
		request["SpecType"] = d.Get("spec_type")
	}

	if !d.IsNewResource() && d.HasChange("eip_max") {
		update = true
		request["EipMax"] = d.Get("eip_max")
	}

	if !d.IsNewResource() && d.HasChange("serverless_config") {
		update = true
		serverlessConfig := make(map[string]interface{})

		if v := d.Get("serverless_config"); v != nil {
			reservedPublishCapacity1, _ := jsonpath.Get("$[0].reserved_publish_capacity", v)
			if reservedPublishCapacity1 != nil && reservedPublishCapacity1 != "" {
				serverlessConfig["ReservedPublishCapacity"] = reservedPublishCapacity1
			}
			reservedSubscribeCapacity1, _ := jsonpath.Get("$[0].reserved_subscribe_capacity", v)
			if reservedSubscribeCapacity1 != nil && reservedSubscribeCapacity1 != "" {
				serverlessConfig["ReservedSubscribeCapacity"] = reservedSubscribeCapacity1
			}

			serverlessConfigJson, err := json.Marshal(serverlessConfig)
			if err != nil {
				return WrapError(err)
			}
			request["ServerlessConfig"] = string(serverlessConfigJson)
		}
	}

	if v, ok := d.GetOkExists("eip_model"); ok {
		request["EipModel"] = v
	}
	if update && enableUpgradePostPayOrder1 {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"ScheduledTask.AlreadyHasSameTaskType"}) || NeedRetry(err) {
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
		aliKafkaServiceV2 := AliKafkaServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"5"}, d.Timeout(schema.TimeoutUpdate), 5*time.Minute, aliKafkaServiceV2.AliKafkaInstanceStateRefreshFunc(d.Id(), "ServiceStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	objectRaw, _ = aliKafkaServiceV2.DescribeAliKafkaInstance(d.Id())
	enableUpgradePrePayOrder1 := false
	checkValue00 = objectRaw["PaidType"]
	if InArray(fmt.Sprint(checkValue00), []string{"0", "4"}) {
		enableUpgradePrePayOrder1 = true
	}
	action = "UpgradePrePayOrder"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("disk_size") {
		update = true
		request["DiskSize"] = d.Get("disk_size")
	}

	if !d.IsNewResource() && d.HasChange("io_max_spec") {
		update = true
		request["IoMaxSpec"] = d.Get("io_max_spec")
	}

	if !d.IsNewResource() && d.HasChange("eip_max") {
		update = true
		request["EipMax"] = d.Get("eip_max")
	}

	if !d.IsNewResource() && d.HasChange("spec_type") {
		update = true
		request["SpecType"] = d.Get("spec_type")
	}

	if !d.IsNewResource() && d.HasChange("confluent_config") {
		update = true
		confluentConfig := make(map[string]interface{})

		if v := d.Get("confluent_config"); v != nil {
			kafkaCu, _ := jsonpath.Get("$[0].kafka_cu", v)
			if kafkaCu != nil && kafkaCu != "" {
				confluentConfig["KafkaCU"] = kafkaCu
			}
			kafkaStorage1, _ := jsonpath.Get("$[0].kafka_storage", v)
			if kafkaStorage1 != nil && kafkaStorage1 != "" {
				confluentConfig["KafkaStorage"] = kafkaStorage1
			}
			kafkaReplica1, _ := jsonpath.Get("$[0].kafka_replica", v)
			if kafkaReplica1 != nil && kafkaReplica1 != "" {
				confluentConfig["KafkaReplica"] = kafkaReplica1
			}
			zooKeeperCu, _ := jsonpath.Get("$[0].zoo_keeper_cu", v)
			if zooKeeperCu != nil && zooKeeperCu != "" {
				confluentConfig["ZooKeeperCU"] = zooKeeperCu
			}
			zooKeeperStorage1, _ := jsonpath.Get("$[0].zoo_keeper_storage", v)
			if zooKeeperStorage1 != nil && zooKeeperStorage1 != "" {
				confluentConfig["ZooKeeperStorage"] = zooKeeperStorage1
			}
			controlCenterCu, _ := jsonpath.Get("$[0].control_center_cu", v)
			if controlCenterCu != nil && controlCenterCu != "" {
				confluentConfig["ControlCenterCU"] = controlCenterCu
			}
			controlCenterStorage1, _ := jsonpath.Get("$[0].control_center_storage", v)
			if controlCenterStorage1 != nil && controlCenterStorage1 != "" {
				confluentConfig["ControlCenterStorage"] = controlCenterStorage1
			}
			schemaRegistryCu, _ := jsonpath.Get("$[0].schema_registry_cu", v)
			if schemaRegistryCu != nil && schemaRegistryCu != "" {
				confluentConfig["SchemaRegistryCU"] = schemaRegistryCu
			}
			schemaRegistryReplica1, _ := jsonpath.Get("$[0].schema_registry_replica", v)
			if schemaRegistryReplica1 != nil && schemaRegistryReplica1 != "" {
				confluentConfig["SchemaRegistryReplica"] = schemaRegistryReplica1
			}
			connectCu, _ := jsonpath.Get("$[0].connect_cu", v)
			if connectCu != nil && connectCu != "" {
				confluentConfig["ConnectCU"] = connectCu
			}
			connectReplica1, _ := jsonpath.Get("$[0].connect_replica", v)
			if connectReplica1 != nil && connectReplica1 != "" {
				confluentConfig["ConnectReplica"] = connectReplica1
			}
			ksqlCu, _ := jsonpath.Get("$[0].ksql_cu", v)
			if ksqlCu != nil && ksqlCu != "" {
				confluentConfig["KsqlCU"] = ksqlCu
			}
			ksqlStorage1, _ := jsonpath.Get("$[0].ksql_storage", v)
			if ksqlStorage1 != nil && ksqlStorage1 != "" {
				confluentConfig["KsqlStorage"] = ksqlStorage1
			}
			ksqlReplica1, _ := jsonpath.Get("$[0].ksql_replica", v)
			if ksqlReplica1 != nil && ksqlReplica1 != "" {
				confluentConfig["KsqlReplica"] = ksqlReplica1
			}
			kafkaRestProxyCu, _ := jsonpath.Get("$[0].kafka_rest_proxy_cu", v)
			if kafkaRestProxyCu != nil && kafkaRestProxyCu != "" {
				confluentConfig["KafkaRestProxyCU"] = kafkaRestProxyCu
			}
			kafkaRestProxyReplica1, _ := jsonpath.Get("$[0].kafka_rest_proxy_replica", v)
			if kafkaRestProxyReplica1 != nil && kafkaRestProxyReplica1 != "" {
				confluentConfig["KafkaRestProxyReplica"] = kafkaRestProxyReplica1
			}
			controlCenterReplica1, _ := jsonpath.Get("$[0].control_center_replica", v)
			if controlCenterReplica1 != nil && controlCenterReplica1 != "" {
				confluentConfig["ControlCenterReplica"] = controlCenterReplica1
			}
			zooKeeperReplica1, _ := jsonpath.Get("$[0].zoo_keeper_replica", v)
			if zooKeeperReplica1 != nil && zooKeeperReplica1 != "" {
				confluentConfig["ZooKeeperReplica"] = zooKeeperReplica1
			}

			confluentConfigJson, err := json.Marshal(confluentConfig)
			if err != nil {
				return WrapError(err)
			}
			request["ConfluentConfig"] = string(confluentConfigJson)
		}
	}

	if !d.IsNewResource() && d.HasChange("paid_type") {
		update = true
		request["PaidType"] = d.Get("paid_type")
	}

	if v, ok := d.GetOkExists("eip_model"); ok {
		request["EipModel"] = v
	}
	if !d.IsNewResource() && d.HasChange("default_topic_partition_num") {
		update = true
		request["PartitionNum"] = d.Get("default_topic_partition_num")
	}

	if update && enableUpgradePrePayOrder1 {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"ScheduledTask.AlreadyHasSameTaskType"}) || NeedRetry(err) {
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
		aliKafkaServiceV2 := AliKafkaServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"5"}, d.Timeout(schema.TimeoutUpdate), 5*time.Minute, aliKafkaServiceV2.AliKafkaInstanceStateRefreshFunc(d.Id(), "ServiceStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "ChangeResourceGroup"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ResourceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if _, ok := d.GetOk("resource_group_id"); ok && !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
	}
	request["NewResourceGroupId"] = d.Get("resource_group_id")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
	action = "EnableAutoGroupCreation"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("enable_auto_group") {
		update = true
	}
	request["Enable"] = d.Get("enable_auto_group")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
	action = "EnableAutoTopicCreation"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("enable_auto_topic") {
		update = true
		request["Operate"] = d.Get("enable_auto_topic")
	}

	if !d.IsNewResource() && d.HasChange("default_topic_partition_num") {
		update = true
		request["PartitionNum"] = d.Get("default_topic_partition_num")
	}

	if v, ok := d.GetOkExists("update_default_topic_partition_num"); ok {
		request["UpdatePartition"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
	action = "UpdateInstanceConfig"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("config") {
		update = true
	}
	request["Config"] = d.Get("config")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
		aliKafkaServiceV2 := AliKafkaServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"5"}, d.Timeout(schema.TimeoutUpdate), 5*time.Minute, aliKafkaServiceV2.AliKafkaInstanceStateRefreshFunc(d.Id(), "ServiceStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "ConvertPostPayOrder"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if v, ok := d.GetOkExists("duration"); ok {
		request["Duration"] = v
	}
	if !d.IsNewResource() && d.HasChange("paid_type") {
		update = true
		request["PaidType"] = d.Get("paid_type")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
	action = "UpgradeInstanceVersion"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("service_version") {
		update = true
	}
	request["TargetVersion"] = d.Get("service_version")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
		aliKafkaServiceV2 := AliKafkaServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"5"}, d.Timeout(schema.TimeoutUpdate), 5*time.Minute, aliKafkaServiceV2.AliKafkaInstanceStateRefreshFunc(d.Id(), "ServiceStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	if d.HasChange("tags") {
		aliKafkaServiceV2 := AliKafkaServiceV2{client}
		if err := aliKafkaServiceV2.SetResourceTags(d, "INSTANCE"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAliCloudAliKafkaInstanceRead(d, meta)
}

func resourceAliCloudAliKafkaInstanceDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	enableDelete := true
	if v, ok := d.GetOkExists("paid_type"); ok {
		if InArray(fmt.Sprint(v), []string{"0", "4"}) {
			enableDelete = false
			log.Printf("[WARN] Cannot destroy resource alicloud_alikafka_instance_v2 which paid_type valued 0,4. Terraform will remove this resource from the state file, however resources may remain.")
		}
	}
	if enableDelete {
		action := "ReleaseInstance"
		var request map[string]interface{}
		var response map[string]interface{}
		query := make(map[string]interface{})
		var err error
		request = make(map[string]interface{})
		request["InstanceId"] = d.Id()
		request["RegionId"] = client.RegionId

		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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

	}

	enableDelete = true
	if v, ok := d.GetOkExists("paid_type"); ok {
		if InArray(fmt.Sprint(v), []string{"0", "4"}) {
			enableDelete = false
			log.Printf("[WARN] Cannot destroy resource alicloud_alikafka_instance_v2 which paid_type valued 0,4. Terraform will remove this resource from the state file, however resources may remain.")
		}
	}
	if enableDelete {
		action := "DeleteInstance"
		var request map[string]interface{}
		var response map[string]interface{}
		query := make(map[string]interface{})
		var err error
		request = make(map[string]interface{})
		request["InstanceId"] = d.Id()
		request["RegionId"] = client.RegionId

		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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

		aliKafkaServiceV2 := AliKafkaServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 5*time.Minute, aliKafkaServiceV2.AliKafkaInstanceStateRefreshFunc(d.Id(), "ServiceStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}

	}
	return nil
}

func convertAliKafkaInstanceInstanceListInstanceVOAutoCreateTopicEnableResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "true":
		return "enable"
	case "false":
		return "disable"
	}
	return source
}
