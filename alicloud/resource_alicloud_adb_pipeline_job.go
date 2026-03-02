// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudAdbPipelineJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAdbPipelineJobCreate,
		Read:   resourceAliCloudAdbPipelineJobRead,
		Update: resourceAliCloudAdbPipelineJobUpdate,
		Delete: resourceAliCloudAdbPipelineJobDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"advanced_config": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"job_config": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"checkpoint_interval": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"enable_serverless": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"max_compute_unit": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"min_compute_unit": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"resource_group": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"pipeline_job_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sink": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"warehouse": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"table_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"enable_unix_timestamp_convert": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"user_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"table_creation_mode": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"enable_ip_white_list": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"unix_timestamp_convert_format": {
										Type:         schema.TypeString,
										Optional:     true,
										ForceNew:     true,
										ValidateFunc: StringInSlice([]string{"APSLiteralTimestampMilliSecond", "APSLiteralTimestampMicroSecond", "APSLiteralTimestampSecond"}, false),
									},
									"db_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"password": {
										Type:      schema.TypeString,
										Optional:  true,
										ForceNew:  true,
										Sensitive: true,
									},
								},
							},
						},
						"lake": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"table_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"partition_specs": {
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"target_type_format": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"source_column": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"strategy": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"target_column": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"source_type_format": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"bucket_num": {
													Type:     schema.TypeInt,
													Optional: true,
													ForceNew: true,
												},
												"truncate_width": {
													Type:     schema.TypeInt,
													Optional: true,
													ForceNew: true,
												},
											},
										},
									},
									"table_creation_mode": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"iceberg": {
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"write_distribution": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"primary_key": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"format_version": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
											},
										},
									},
									"db_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"storage_target": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"table_format": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"paimon": {
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"table_type": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"primary_key": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"source": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"kafka": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_offset_mode": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"network_gateway": {
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sasl_password": {
													Type:      schema.TypeString,
													Optional:  true,
													ForceNew:  true,
													Sensitive: true,
												},
												"vpc_id": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"security_protocol": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"bootstrap_servers": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"vswitch_id": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"sasl_username": {
													Type:      schema.TypeString,
													Optional:  true,
													ForceNew:  true,
													Sensitive: true,
												},
											},
										},
									},
									"message_format": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"cloud_managed": {
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enable_cross_account": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"across_role": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"across_uid": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
											},
										},
									},
									"kafka_topic": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"start_offset_value": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"instance_type": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"kafka_cluster_id": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"cdc_config": {
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sql_types": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
												"sync_tables": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
												},
											},
										},
									},
								},
							},
						},
						"sls": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_offset_mode": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"project": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"enable_cross_account": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"log_store": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"start_offset_value": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"across_role": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"across_uid": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
					},
				},
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"Created", "Running", "Stopped", "Failed"}, false),
			},
			"transform": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"column_map": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"map_name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"map_type": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
						"dirty_data_handle_mode": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceAliCloudAdbPipelineJobCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreatePipelineJob"
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

	source := make(map[string]interface{})

	if v := d.Get("source"); !IsNil(v) {
		kafka := make(map[string]interface{})
		cdcConfig := make(map[string]interface{})
		sqlTypes1, _ := jsonpath.Get("$[0].kafka[0].cdc_config[0].sql_types", d.Get("source"))
		if sqlTypes1 != nil && sqlTypes1 != "" {
			cdcConfig["SqlTypes"] = sqlTypes1
		}
		syncTables1, _ := jsonpath.Get("$[0].kafka[0].cdc_config[0].sync_tables", d.Get("source"))
		if syncTables1 != nil && syncTables1 != "" {
			cdcConfig["SyncTables"] = syncTables1
		}

		if len(cdcConfig) > 0 {
			kafka["CdcConfig"] = cdcConfig
		}
		kafkaTopic1, _ := jsonpath.Get("$[0].kafka[0].kafka_topic", d.Get("source"))
		if kafkaTopic1 != nil && kafkaTopic1 != "" {
			kafka["KafkaTopic"] = kafkaTopic1
		}
		networkGateway := make(map[string]interface{})
		vswitchId1, _ := jsonpath.Get("$[0].kafka[0].network_gateway[0].vswitch_id", d.Get("source"))
		if vswitchId1 != nil && vswitchId1 != "" {
			networkGateway["VswitchId"] = vswitchId1
		}
		saslUsername1, _ := jsonpath.Get("$[0].kafka[0].network_gateway[0].sasl_username", d.Get("source"))
		if saslUsername1 != nil && saslUsername1 != "" {
			networkGateway["SaslUsername"] = saslUsername1
		}
		bootstrapServers1, _ := jsonpath.Get("$[0].kafka[0].network_gateway[0].bootstrap_servers", d.Get("source"))
		if bootstrapServers1 != nil && bootstrapServers1 != "" {
			networkGateway["BootstrapServers"] = bootstrapServers1
		}
		saslPassword1, _ := jsonpath.Get("$[0].kafka[0].network_gateway[0].sasl_password", d.Get("source"))
		if saslPassword1 != nil && saslPassword1 != "" {
			networkGateway["SaslPassword"] = saslPassword1
		}
		securityProtocol1, _ := jsonpath.Get("$[0].kafka[0].network_gateway[0].security_protocol", d.Get("source"))
		if securityProtocol1 != nil && securityProtocol1 != "" {
			networkGateway["SecurityProtocol"] = securityProtocol1
		}
		vpcId1, _ := jsonpath.Get("$[0].kafka[0].network_gateway[0].vpc_id", d.Get("source"))
		if vpcId1 != nil && vpcId1 != "" {
			networkGateway["VpcId"] = vpcId1
		}

		if len(networkGateway) > 0 {
			kafka["NetworkGateway"] = networkGateway
		}
		cloudManaged := make(map[string]interface{})
		enableCrossAccount1, _ := jsonpath.Get("$[0].kafka[0].cloud_managed[0].enable_cross_account", d.Get("source"))
		if enableCrossAccount1 != nil && enableCrossAccount1 != "" {
			cloudManaged["EnableCrossAccount"] = enableCrossAccount1
		}
		acrossUid1, _ := jsonpath.Get("$[0].kafka[0].cloud_managed[0].across_uid", d.Get("source"))
		if acrossUid1 != nil && acrossUid1 != "" {
			cloudManaged["AcrossUid"] = acrossUid1
		}
		acrossRole1, _ := jsonpath.Get("$[0].kafka[0].cloud_managed[0].across_role", d.Get("source"))
		if acrossRole1 != nil && acrossRole1 != "" {
			cloudManaged["AcrossRole"] = acrossRole1
		}

		if len(cloudManaged) > 0 {
			kafka["CloudManaged"] = cloudManaged
		}
		startOffsetValue1, _ := jsonpath.Get("$[0].kafka[0].start_offset_value", d.Get("source"))
		if startOffsetValue1 != nil && startOffsetValue1 != "" {
			kafka["StartOffsetValue"] = startOffsetValue1
		}
		instanceType1, _ := jsonpath.Get("$[0].kafka[0].instance_type", d.Get("source"))
		if instanceType1 != nil && instanceType1 != "" {
			kafka["InstanceType"] = instanceType1
		}
		startOffsetMode1, _ := jsonpath.Get("$[0].kafka[0].start_offset_mode", d.Get("source"))
		if startOffsetMode1 != nil && startOffsetMode1 != "" {
			kafka["StartOffsetMode"] = startOffsetMode1
		}
		kafkaClusterId1, _ := jsonpath.Get("$[0].kafka[0].kafka_cluster_id", d.Get("source"))
		if kafkaClusterId1 != nil && kafkaClusterId1 != "" {
			kafka["KafkaClusterId"] = kafkaClusterId1
		}
		messageFormat1, _ := jsonpath.Get("$[0].kafka[0].message_format", d.Get("source"))
		if messageFormat1 != nil && messageFormat1 != "" {
			kafka["MessageFormat"] = messageFormat1
		}

		if len(kafka) > 0 {
			source["Kafka"] = kafka
		}
		type1, _ := jsonpath.Get("$[0].type", v)
		if type1 != nil && type1 != "" {
			source["Type"] = type1
		}
		sls := make(map[string]interface{})
		acrossRole3, _ := jsonpath.Get("$[0].sls[0].across_role", d.Get("source"))
		if acrossRole3 != nil && acrossRole3 != "" {
			sls["AcrossRole"] = acrossRole3
		}
		startOffsetValue3, _ := jsonpath.Get("$[0].sls[0].start_offset_value", d.Get("source"))
		if startOffsetValue3 != nil && startOffsetValue3 != "" {
			sls["StartOffsetValue"] = startOffsetValue3
		}
		startOffsetMode3, _ := jsonpath.Get("$[0].sls[0].start_offset_mode", d.Get("source"))
		if startOffsetMode3 != nil && startOffsetMode3 != "" {
			sls["StartOffsetMode"] = startOffsetMode3
		}
		project1, _ := jsonpath.Get("$[0].sls[0].project", d.Get("source"))
		if project1 != nil && project1 != "" {
			sls["Project"] = project1
		}
		enableCrossAccount3, _ := jsonpath.Get("$[0].sls[0].enable_cross_account", d.Get("source"))
		if enableCrossAccount3 != nil && enableCrossAccount3 != "" {
			sls["EnableCrossAccount"] = enableCrossAccount3
		}
		logStore1, _ := jsonpath.Get("$[0].sls[0].log_store", d.Get("source"))
		if logStore1 != nil && logStore1 != "" {
			sls["LogStore"] = logStore1
		}
		acrossUid3, _ := jsonpath.Get("$[0].sls[0].across_uid", d.Get("source"))
		if acrossUid3 != nil && acrossUid3 != "" {
			sls["AcrossUid"] = acrossUid3
		}

		if len(sls) > 0 {
			source["Sls"] = sls
		}

		sourceJson, err := json.Marshal(source)
		if err != nil {
			return WrapError(err)
		}
		request["Source"] = string(sourceJson)
	}

	jobConfig := make(map[string]interface{})

	if v := d.Get("job_config"); !IsNil(v) {
		checkpointInterval1, _ := jsonpath.Get("$[0].checkpoint_interval", v)
		if checkpointInterval1 != nil && checkpointInterval1 != "" {
			jobConfig["CheckpointInterval"] = checkpointInterval1
		}
		maxComputeUnit1, _ := jsonpath.Get("$[0].max_compute_unit", v)
		if maxComputeUnit1 != nil && maxComputeUnit1 != "" {
			jobConfig["MaxComputeUnit"] = maxComputeUnit1
		}
		resourceGroup1, _ := jsonpath.Get("$[0].resource_group", v)
		if resourceGroup1 != nil && resourceGroup1 != "" {
			jobConfig["ResourceGroup"] = resourceGroup1
		}
		enableServerless1, _ := jsonpath.Get("$[0].enable_serverless", v)
		if enableServerless1 != nil && enableServerless1 != "" {
			jobConfig["EnableServerless"] = enableServerless1
		}
		minComputeUnit1, _ := jsonpath.Get("$[0].min_compute_unit", v)
		if minComputeUnit1 != nil && minComputeUnit1 != "" {
			jobConfig["MinComputeUnit"] = minComputeUnit1
		}

		jobConfigJson, err := json.Marshal(jobConfig)
		if err != nil {
			return WrapError(err)
		}
		request["JobConfig"] = string(jobConfigJson)
	}

	sink := make(map[string]interface{})

	if v := d.Get("sink"); !IsNil(v) {
		warehouse := make(map[string]interface{})
		password1, _ := jsonpath.Get("$[0].warehouse[0].password", d.Get("sink"))
		if password1 != nil && password1 != "" {
			warehouse["Password"] = password1
		}
		enableIpWhiteList1, _ := jsonpath.Get("$[0].warehouse[0].enable_ip_white_list", d.Get("sink"))
		if enableIpWhiteList1 != nil && enableIpWhiteList1 != "" {
			warehouse["EnableIpWhiteList"] = enableIpWhiteList1
		}
		unixTimestampConvertFormat1, _ := jsonpath.Get("$[0].warehouse[0].unix_timestamp_convert_format", d.Get("sink"))
		if unixTimestampConvertFormat1 != nil && unixTimestampConvertFormat1 != "" {
			warehouse["UnixTimestampConvertFormat"] = unixTimestampConvertFormat1
		}
		enableUnixTimestampConvert1, _ := jsonpath.Get("$[0].warehouse[0].enable_unix_timestamp_convert", d.Get("sink"))
		if enableUnixTimestampConvert1 != nil && enableUnixTimestampConvert1 != "" {
			warehouse["EnableUnixTimestampConvert"] = enableUnixTimestampConvert1
		}
		dbName1, _ := jsonpath.Get("$[0].warehouse[0].db_name", d.Get("sink"))
		if dbName1 != nil && dbName1 != "" {
			warehouse["DbName"] = dbName1
		}
		userName1, _ := jsonpath.Get("$[0].warehouse[0].user_name", d.Get("sink"))
		if userName1 != nil && userName1 != "" {
			warehouse["UserName"] = userName1
		}
		tableName1, _ := jsonpath.Get("$[0].warehouse[0].table_name", d.Get("sink"))
		if tableName1 != nil && tableName1 != "" {
			warehouse["TableName"] = tableName1
		}
		tableCreationMode1, _ := jsonpath.Get("$[0].warehouse[0].table_creation_mode", d.Get("sink"))
		if tableCreationMode1 != nil && tableCreationMode1 != "" {
			warehouse["TableCreationMode"] = tableCreationMode1
		}

		if len(warehouse) > 0 {
			sink["Warehouse"] = warehouse
		}
		lake := make(map[string]interface{})
		localData, err := jsonpath.Get("$[0].lake[0].partition_specs", v)
		if err != nil {
			localData = make([]interface{}, 0)
		}
		localMaps := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(localData) {
			dataLoopTmp := make(map[string]interface{})
			if dataLoop != nil {
				dataLoopTmp = dataLoop.(map[string]interface{})
			}
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["TruncateWidth"] = dataLoopTmp["truncate_width"]
			dataLoopMap["TargetColumn"] = dataLoopTmp["target_column"]
			dataLoopMap["SourceTypeFormat"] = dataLoopTmp["source_type_format"]
			dataLoopMap["BucketNum"] = dataLoopTmp["bucket_num"]
			dataLoopMap["TargetTypeFormat"] = dataLoopTmp["target_type_format"]
			dataLoopMap["Strategy"] = dataLoopTmp["strategy"]
			dataLoopMap["SourceColumn"] = dataLoopTmp["source_column"]
			localMaps = append(localMaps, dataLoopMap)
		}
		lake["PartitionSpecs"] = localMaps

		tableFormat1, _ := jsonpath.Get("$[0].lake[0].table_format", d.Get("sink"))
		if tableFormat1 != nil && tableFormat1 != "" {
			lake["TableFormat"] = tableFormat1
		}
		iceberg := make(map[string]interface{})
		writeDistribution1, _ := jsonpath.Get("$[0].lake[0].iceberg[0].write_distribution", d.Get("sink"))
		if writeDistribution1 != nil && writeDistribution1 != "" {
			iceberg["WriteDistribution"] = writeDistribution1
		}
		formatVersion1, _ := jsonpath.Get("$[0].lake[0].iceberg[0].format_version", d.Get("sink"))
		if formatVersion1 != nil && formatVersion1 != "" {
			iceberg["FormatVersion"] = formatVersion1
		}
		primaryKey1, _ := jsonpath.Get("$[0].lake[0].iceberg[0].primary_key", d.Get("sink"))
		if primaryKey1 != nil && primaryKey1 != "" {
			iceberg["PrimaryKey"] = primaryKey1
		}

		if len(iceberg) > 0 {
			lake["Iceberg"] = iceberg
		}
		paimon := make(map[string]interface{})
		tableType1, _ := jsonpath.Get("$[0].lake[0].paimon[0].table_type", d.Get("sink"))
		if tableType1 != nil && tableType1 != "" {
			paimon["TableType"] = tableType1
		}
		primaryKey3, _ := jsonpath.Get("$[0].lake[0].paimon[0].primary_key", d.Get("sink"))
		if primaryKey3 != nil && primaryKey3 != "" {
			paimon["PrimaryKey"] = primaryKey3
		}

		if len(paimon) > 0 {
			lake["Paimon"] = paimon
		}
		tableCreationMode3, _ := jsonpath.Get("$[0].lake[0].table_creation_mode", d.Get("sink"))
		if tableCreationMode3 != nil && tableCreationMode3 != "" {
			lake["TableCreationMode"] = tableCreationMode3
		}
		tableName3, _ := jsonpath.Get("$[0].lake[0].table_name", d.Get("sink"))
		if tableName3 != nil && tableName3 != "" {
			lake["TableName"] = tableName3
		}
		storageTarget1, _ := jsonpath.Get("$[0].lake[0].storage_target", d.Get("sink"))
		if storageTarget1 != nil && storageTarget1 != "" {
			lake["StorageTarget"] = storageTarget1
		}
		dbName3, _ := jsonpath.Get("$[0].lake[0].db_name", d.Get("sink"))
		if dbName3 != nil && dbName3 != "" {
			lake["DbName"] = dbName3
		}

		if len(lake) > 0 {
			sink["Lake"] = lake
		}
		type3, _ := jsonpath.Get("$[0].type", v)
		if type3 != nil && type3 != "" {
			sink["Type"] = type3
		}

		sinkJson, err := json.Marshal(sink)
		if err != nil {
			return WrapError(err)
		}
		request["Sink"] = string(sinkJson)
	}

	transform := make(map[string]interface{})

	if v := d.Get("transform"); !IsNil(v) {
		localData1, err := jsonpath.Get("$[0].column_map", v)
		if err != nil {
			localData1 = make([]interface{}, 0)
		}
		localMaps1 := make([]interface{}, 0)
		for _, dataLoop1 := range convertToInterfaceArray(localData1) {
			dataLoop1Tmp := make(map[string]interface{})
			if dataLoop1 != nil {
				dataLoop1Tmp = dataLoop1.(map[string]interface{})
			}
			dataLoop1Map := make(map[string]interface{})
			dataLoop1Map["Type"] = dataLoop1Tmp["type"]
			dataLoop1Map["MapName"] = dataLoop1Tmp["map_name"]
			dataLoop1Map["MapType"] = dataLoop1Tmp["map_type"]
			dataLoop1Map["Name"] = dataLoop1Tmp["name"]
			localMaps1 = append(localMaps1, dataLoop1Map)
		}
		transform["ColumnMap"] = localMaps1

		dirtyDataHandleMode1, _ := jsonpath.Get("$[0].dirty_data_handle_mode", v)
		if dirtyDataHandleMode1 != nil && dirtyDataHandleMode1 != "" {
			transform["DirtyDataHandleMode"] = dirtyDataHandleMode1
		}

		transformJson, err := json.Marshal(transform)
		if err != nil {
			return WrapError(err)
		}
		request["Transform"] = string(transformJson)
	}

	if v, ok := d.GetOk("advanced_config"); ok {
		request["AdvancedConfig"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_adb_pipeline_job", action, AlibabaCloudSdkGoERROR)
	}

	DataPipelineJobIdVar, _ := jsonpath.Get("$.Data.PipelineJobId", response)
	d.SetId(fmt.Sprintf("%v:%v", request["DBClusterId"], DataPipelineJobIdVar))

	return resourceAliCloudAdbPipelineJobUpdate(d, meta)
}

func resourceAliCloudAdbPipelineJobRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	adbServiceV2 := AdbServiceV2{client}

	objectRaw, err := adbServiceV2.DescribeAdbPipelineJob(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_adb_pipeline_job DescribeAdbPipelineJob Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("advanced_config", objectRaw["AdvancedConfig"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("description", objectRaw["Description"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("status", objectRaw["Status"])
	d.Set("db_cluster_id", objectRaw["DBClusterId"])
	d.Set("pipeline_job_id", objectRaw["PipelineJobId"])

	jobConfigMaps := make([]map[string]interface{}, 0)
	jobConfigMap := make(map[string]interface{})
	jobConfigRaw := make(map[string]interface{})
	if objectRaw["JobConfig"] != nil {
		jobConfigRaw = objectRaw["JobConfig"].(map[string]interface{})
	}
	if len(jobConfigRaw) > 0 {
		jobConfigMap["checkpoint_interval"] = jobConfigRaw["CheckpointInterval"]
		jobConfigMap["enable_serverless"] = jobConfigRaw["EnableServerless"]
		jobConfigMap["max_compute_unit"] = jobConfigRaw["MaxComputeUnit"]
		jobConfigMap["min_compute_unit"] = jobConfigRaw["MinComputeUnit"]
		jobConfigMap["resource_group"] = jobConfigRaw["ResourceGroup"]

		jobConfigMaps = append(jobConfigMaps, jobConfigMap)
	}
	if err := d.Set("job_config", jobConfigMaps); err != nil {
		return err
	}
	sinkMaps := make([]map[string]interface{}, 0)
	sinkMap := make(map[string]interface{})
	sinkRaw := make(map[string]interface{})
	if objectRaw["Sink"] != nil {
		sinkRaw = objectRaw["Sink"].(map[string]interface{})
	}
	if len(sinkRaw) > 0 {
		sinkMap["type"] = sinkRaw["Type"]

		lakeMaps := make([]map[string]interface{}, 0)
		lakeMap := make(map[string]interface{})
		lakeRaw := make(map[string]interface{})
		if sinkRaw["Lake"] != nil {
			lakeRaw = sinkRaw["Lake"].(map[string]interface{})
		}
		if len(lakeRaw) > 0 {
			lakeMap["db_name"] = lakeRaw["DbName"]
			lakeMap["storage_target"] = lakeRaw["StorageTarget"]
			lakeMap["table_creation_mode"] = lakeRaw["TableCreationMode"]
			lakeMap["table_format"] = lakeRaw["TableFormat"]
			lakeMap["table_name"] = lakeRaw["TableName"]

			icebergMaps := make([]map[string]interface{}, 0)
			icebergMap := make(map[string]interface{})
			icebergRaw := make(map[string]interface{})
			if lakeRaw["Iceberg"] != nil {
				icebergRaw = lakeRaw["Iceberg"].(map[string]interface{})
			}
			if len(icebergRaw) > 0 {
				icebergMap["format_version"] = icebergRaw["FormatVersion"]
				icebergMap["primary_key"] = icebergRaw["PrimaryKey"]
				icebergMap["write_distribution"] = icebergRaw["WriteDistribution"]

				icebergMaps = append(icebergMaps, icebergMap)
			}
			lakeMap["iceberg"] = icebergMaps
			paimonMaps := make([]map[string]interface{}, 0)
			paimonMap := make(map[string]interface{})
			paimonRaw := make(map[string]interface{})
			if lakeRaw["Paimon"] != nil {
				paimonRaw = lakeRaw["Paimon"].(map[string]interface{})
			}
			if len(paimonRaw) > 0 {
				paimonMap["primary_key"] = paimonRaw["PrimaryKey"]
				paimonMap["table_type"] = paimonRaw["TableType"]

				paimonMaps = append(paimonMaps, paimonMap)
			}
			lakeMap["paimon"] = paimonMaps
			partitionSpecsRaw := lakeRaw["PartitionSpecs"]
			partitionSpecsMaps := make([]map[string]interface{}, 0)
			if partitionSpecsRaw != nil {
				for _, partitionSpecsChildRaw := range convertToInterfaceArray(partitionSpecsRaw) {
					partitionSpecsMap := make(map[string]interface{})
					partitionSpecsChildRaw := partitionSpecsChildRaw.(map[string]interface{})
					partitionSpecsMap["bucket_num"] = partitionSpecsChildRaw["BucketNum"]
					partitionSpecsMap["source_column"] = partitionSpecsChildRaw["SourceColumn"]
					partitionSpecsMap["source_type_format"] = partitionSpecsChildRaw["SourceTypeFormat"]
					partitionSpecsMap["strategy"] = partitionSpecsChildRaw["Strategy"]
					partitionSpecsMap["target_column"] = partitionSpecsChildRaw["TargetColumn"]
					partitionSpecsMap["target_type_format"] = partitionSpecsChildRaw["TargetTypeFormat"]
					partitionSpecsMap["truncate_width"] = partitionSpecsChildRaw["TruncateWidth"]

					partitionSpecsMaps = append(partitionSpecsMaps, partitionSpecsMap)
				}
			}
			lakeMap["partition_specs"] = partitionSpecsMaps
			lakeMaps = append(lakeMaps, lakeMap)
		}
		sinkMap["lake"] = lakeMaps
		warehouseMaps := make([]map[string]interface{}, 0)
		warehouseMap := make(map[string]interface{})
		warehouseRaw := make(map[string]interface{})
		if sinkRaw["Warehouse"] != nil {
			warehouseRaw = sinkRaw["Warehouse"].(map[string]interface{})
		}
		if len(warehouseRaw) > 0 {
			warehouseMap["db_name"] = warehouseRaw["DbName"]
			warehouseMap["enable_ip_white_list"] = warehouseRaw["EnableIpWhiteList"]
			warehouseMap["enable_unix_timestamp_convert"] = warehouseRaw["EnableUnixTimestampConvert"]
			warehouseMap["password"] = warehouseRaw["Password"]
			warehouseMap["table_creation_mode"] = warehouseRaw["TableCreationMode"]
			warehouseMap["table_name"] = warehouseRaw["TableName"]
			warehouseMap["unix_timestamp_convert_format"] = warehouseRaw["UnixTimestampConvertFormat"]
			warehouseMap["user_name"] = warehouseRaw["UserName"]

			warehouseMaps = append(warehouseMaps, warehouseMap)
		}
		sinkMap["warehouse"] = warehouseMaps
		sinkMaps = append(sinkMaps, sinkMap)
	}
	if err := d.Set("sink", sinkMaps); err != nil {
		return err
	}
	sourceMaps := make([]map[string]interface{}, 0)
	sourceMap := make(map[string]interface{})
	sourceRaw := make(map[string]interface{})
	if objectRaw["Source"] != nil {
		sourceRaw = objectRaw["Source"].(map[string]interface{})
	}
	if len(sourceRaw) > 0 {
		sourceMap["type"] = sourceRaw["Type"]

		kafkaMaps := make([]map[string]interface{}, 0)
		kafkaMap := make(map[string]interface{})
		kafkaRaw := make(map[string]interface{})
		if sourceRaw["Kafka"] != nil {
			kafkaRaw = sourceRaw["Kafka"].(map[string]interface{})
		}
		if len(kafkaRaw) > 0 {
			kafkaMap["instance_type"] = kafkaRaw["InstanceType"]
			kafkaMap["kafka_cluster_id"] = kafkaRaw["KafkaClusterId"]
			kafkaMap["kafka_topic"] = kafkaRaw["KafkaTopic"]
			kafkaMap["message_format"] = kafkaRaw["MessageFormat"]
			kafkaMap["start_offset_mode"] = kafkaRaw["StartOffsetMode"]
			kafkaMap["start_offset_value"] = kafkaRaw["StartOffsetValue"]

			cdcConfigMaps := make([]map[string]interface{}, 0)
			cdcConfigMap := make(map[string]interface{})
			cdcConfigRaw := make(map[string]interface{})
			if kafkaRaw["CdcConfig"] != nil {
				cdcConfigRaw = kafkaRaw["CdcConfig"].(map[string]interface{})
			}
			if len(cdcConfigRaw) > 0 {
				cdcConfigMap["sql_types"] = cdcConfigRaw["SqlTypes"]
				cdcConfigMap["sync_tables"] = cdcConfigRaw["SyncTables"]

				cdcConfigMaps = append(cdcConfigMaps, cdcConfigMap)
			}
			kafkaMap["cdc_config"] = cdcConfigMaps
			cloudManagedMaps := make([]map[string]interface{}, 0)
			cloudManagedMap := make(map[string]interface{})
			cloudManagedRaw := make(map[string]interface{})
			if kafkaRaw["CloudManaged"] != nil {
				cloudManagedRaw = kafkaRaw["CloudManaged"].(map[string]interface{})
			}
			if len(cloudManagedRaw) > 0 {
				cloudManagedMap["across_role"] = cloudManagedRaw["AcrossRole"]
				cloudManagedMap["across_uid"] = cloudManagedRaw["AcrossUid"]
				cloudManagedMap["enable_cross_account"] = cloudManagedRaw["EnableCrossAccount"]

				cloudManagedMaps = append(cloudManagedMaps, cloudManagedMap)
			}
			kafkaMap["cloud_managed"] = cloudManagedMaps
			networkGatewayMaps := make([]map[string]interface{}, 0)
			networkGatewayMap := make(map[string]interface{})
			networkGatewayRaw := make(map[string]interface{})
			if kafkaRaw["NetworkGateway"] != nil {
				networkGatewayRaw = kafkaRaw["NetworkGateway"].(map[string]interface{})
			}
			if len(networkGatewayRaw) > 0 {
				networkGatewayMap["bootstrap_servers"] = networkGatewayRaw["BootstrapServers"]
				networkGatewayMap["sasl_password"] = networkGatewayRaw["SaslPassword"]
				networkGatewayMap["sasl_username"] = networkGatewayRaw["SaslUsername"]
				networkGatewayMap["security_protocol"] = networkGatewayRaw["SecurityProtocol"]
				networkGatewayMap["vpc_id"] = networkGatewayRaw["VpcId"]
				networkGatewayMap["vswitch_id"] = networkGatewayRaw["VswitchId"]

				networkGatewayMaps = append(networkGatewayMaps, networkGatewayMap)
			}
			kafkaMap["network_gateway"] = networkGatewayMaps
			kafkaMaps = append(kafkaMaps, kafkaMap)
		}
		sourceMap["kafka"] = kafkaMaps
		slsMaps := make([]map[string]interface{}, 0)
		slsMap := make(map[string]interface{})
		slsRaw := make(map[string]interface{})
		if sourceRaw["Sls"] != nil {
			slsRaw = sourceRaw["Sls"].(map[string]interface{})
		}
		if len(slsRaw) > 0 {
			slsMap["across_role"] = slsRaw["AcrossRole"]
			slsMap["across_uid"] = slsRaw["AcrossUid"]
			slsMap["enable_cross_account"] = slsRaw["EnableCrossAccount"]
			slsMap["log_store"] = slsRaw["LogStore"]
			slsMap["project"] = slsRaw["Project"]
			slsMap["start_offset_mode"] = slsRaw["StartOffsetMode"]
			slsMap["start_offset_value"] = slsRaw["StartOffsetValue"]

			slsMaps = append(slsMaps, slsMap)
		}
		sourceMap["sls"] = slsMaps
		sourceMaps = append(sourceMaps, sourceMap)
	}
	if err := d.Set("source", sourceMaps); err != nil {
		return err
	}
	transformMaps := make([]map[string]interface{}, 0)
	transformMap := make(map[string]interface{})
	transformRaw := make(map[string]interface{})
	if objectRaw["Transform"] != nil {
		transformRaw = objectRaw["Transform"].(map[string]interface{})
	}
	if len(transformRaw) > 0 {
		transformMap["dirty_data_handle_mode"] = transformRaw["DirtyDataHandleMode"]

		columnMapRaw := transformRaw["ColumnMap"]
		columnMapMaps := make([]map[string]interface{}, 0)
		if columnMapRaw != nil {
			for _, columnMapChildRaw := range convertToInterfaceArray(columnMapRaw) {
				columnMapMap := make(map[string]interface{})
				columnMapChildRaw := columnMapChildRaw.(map[string]interface{})
				columnMapMap["map_name"] = columnMapChildRaw["MapName"]
				columnMapMap["map_type"] = columnMapChildRaw["MapType"]
				columnMapMap["name"] = columnMapChildRaw["Name"]
				columnMapMap["type"] = columnMapChildRaw["Type"]

				columnMapMaps = append(columnMapMaps, columnMapMap)
			}
		}
		transformMap["column_map"] = columnMapMaps
		transformMaps = append(transformMaps, transformMap)
	}
	if err := d.Set("transform", transformMaps); err != nil {
		return err
	}

	return nil
}

func resourceAliCloudAdbPipelineJobUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	adbServiceV2 := AdbServiceV2{client}
	objectRaw, _ := adbServiceV2.DescribeAdbPipelineJob(d.Id())

	if d.HasChange("status") {
		var err error
		target := d.Get("status").(string)

		currentStatus, err := jsonpath.Get("Status", objectRaw)
		if err != nil {
			return WrapErrorf(err, FailedGetAttributeMsg, d.Id(), "Status", objectRaw)
		}
		if fmt.Sprint(currentStatus) != target {
			if target == "Running" {
				parts := strings.Split(d.Id(), ":")
				action := "StartPipelineJob"
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				request["DBClusterId"] = parts[0]
				request["PipelineJobId"] = parts[1]
				request["RegionId"] = client.RegionId
				request["ClientToken"] = buildClientToken(action)
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
				adbServiceV2 := AdbServiceV2{client}
				stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, adbServiceV2.AdbPipelineJobStateRefreshFunc(d.Id(), "Status", []string{}))
				if _, err := stateConf.WaitForState(); err != nil {
					return WrapErrorf(err, IdMsg, d.Id())
				}

			}
			if target == "Stopped" {
				parts := strings.Split(d.Id(), ":")
				action := "StopPipelineJob"
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				request["DBClusterId"] = parts[0]
				request["PipelineJobId"] = parts[1]
				request["RegionId"] = client.RegionId
				request["ClientToken"] = buildClientToken(action)
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
				adbServiceV2 := AdbServiceV2{client}
				stateConf := BuildStateConf([]string{}, []string{"Stopped"}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, adbServiceV2.AdbPipelineJobStateRefreshFunc(d.Id(), "Status", []string{}))
				if _, err := stateConf.WaitForState(); err != nil {
					return WrapErrorf(err, IdMsg, d.Id())
				}

			}
		}
	}

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "UpdatePipelineJob"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["PipelineJobId"] = parts[1]
	request["DBClusterId"] = parts[0]
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("source") {
		update = true
		source := make(map[string]interface{})

		if v := d.Get("source"); v != nil {
			kafka := make(map[string]interface{})
			cdcConfig := make(map[string]interface{})
			sqlTypes1, _ := jsonpath.Get("$[0].kafka[0].cdc_config[0].sql_types", d.Get("source"))
			if sqlTypes1 != nil && sqlTypes1 != "" {
				cdcConfig["SqlTypes"] = sqlTypes1
			}
			syncTables1, _ := jsonpath.Get("$[0].kafka[0].cdc_config[0].sync_tables", d.Get("source"))
			if syncTables1 != nil && syncTables1 != "" {
				cdcConfig["SyncTables"] = syncTables1
			}

			if len(cdcConfig) > 0 {
				kafka["CdcConfig"] = cdcConfig
			}
			kafkaTopic1, _ := jsonpath.Get("$[0].kafka[0].kafka_topic", d.Get("source"))
			if kafkaTopic1 != nil && kafkaTopic1 != "" {
				kafka["KafkaTopic"] = kafkaTopic1
			}
			networkGateway := make(map[string]interface{})
			vswitchId1, _ := jsonpath.Get("$[0].kafka[0].network_gateway[0].vswitch_id", d.Get("source"))
			if vswitchId1 != nil && vswitchId1 != "" {
				networkGateway["VswitchId"] = vswitchId1
			}
			saslUsername1, _ := jsonpath.Get("$[0].kafka[0].network_gateway[0].sasl_username", d.Get("source"))
			if saslUsername1 != nil && saslUsername1 != "" {
				networkGateway["SaslUsername"] = saslUsername1
			}
			bootstrapServers1, _ := jsonpath.Get("$[0].kafka[0].network_gateway[0].bootstrap_servers", d.Get("source"))
			if bootstrapServers1 != nil && bootstrapServers1 != "" {
				networkGateway["BootstrapServers"] = bootstrapServers1
			}
			saslPassword1, _ := jsonpath.Get("$[0].kafka[0].network_gateway[0].sasl_password", d.Get("source"))
			if saslPassword1 != nil && saslPassword1 != "" {
				networkGateway["SaslPassword"] = saslPassword1
			}
			securityProtocol1, _ := jsonpath.Get("$[0].kafka[0].network_gateway[0].security_protocol", d.Get("source"))
			if securityProtocol1 != nil && securityProtocol1 != "" {
				networkGateway["SecurityProtocol"] = securityProtocol1
			}
			vpcId1, _ := jsonpath.Get("$[0].kafka[0].network_gateway[0].vpc_id", d.Get("source"))
			if vpcId1 != nil && vpcId1 != "" {
				networkGateway["VpcId"] = vpcId1
			}

			if len(networkGateway) > 0 {
				kafka["NetworkGateway"] = networkGateway
			}
			cloudManaged := make(map[string]interface{})
			enableCrossAccount1, _ := jsonpath.Get("$[0].kafka[0].cloud_managed[0].enable_cross_account", d.Get("source"))
			if enableCrossAccount1 != nil && enableCrossAccount1 != "" {
				cloudManaged["EnableCrossAccount"] = enableCrossAccount1
			}
			acrossUid1, _ := jsonpath.Get("$[0].kafka[0].cloud_managed[0].across_uid", d.Get("source"))
			if acrossUid1 != nil && acrossUid1 != "" {
				cloudManaged["AcrossUid"] = acrossUid1
			}
			acrossRole1, _ := jsonpath.Get("$[0].kafka[0].cloud_managed[0].across_role", d.Get("source"))
			if acrossRole1 != nil && acrossRole1 != "" {
				cloudManaged["AcrossRole"] = acrossRole1
			}

			if len(cloudManaged) > 0 {
				kafka["CloudManaged"] = cloudManaged
			}
			startOffsetValue1, _ := jsonpath.Get("$[0].kafka[0].start_offset_value", d.Get("source"))
			if startOffsetValue1 != nil && startOffsetValue1 != "" {
				kafka["StartOffsetValue"] = startOffsetValue1
			}
			startOffsetMode1, _ := jsonpath.Get("$[0].kafka[0].start_offset_mode", d.Get("source"))
			if startOffsetMode1 != nil && startOffsetMode1 != "" {
				kafka["StartOffsetMode"] = startOffsetMode1
			}
			kafkaClusterId1, _ := jsonpath.Get("$[0].kafka[0].kafka_cluster_id", d.Get("source"))
			if kafkaClusterId1 != nil && kafkaClusterId1 != "" {
				kafka["KafkaClusterId"] = kafkaClusterId1
			}
			messageFormat1, _ := jsonpath.Get("$[0].kafka[0].message_format", d.Get("source"))
			if messageFormat1 != nil && messageFormat1 != "" {
				kafka["MessageFormat"] = messageFormat1
			}

			if len(kafka) > 0 {
				source["Kafka"] = kafka
			}
			sls := make(map[string]interface{})
			acrossRole3, _ := jsonpath.Get("$[0].sls[0].across_role", d.Get("source"))
			if acrossRole3 != nil && acrossRole3 != "" {
				sls["AcrossRole"] = acrossRole3
			}
			startOffsetValue3, _ := jsonpath.Get("$[0].sls[0].start_offset_value", d.Get("source"))
			if startOffsetValue3 != nil && startOffsetValue3 != "" {
				sls["StartOffsetValue"] = startOffsetValue3
			}
			startOffsetMode3, _ := jsonpath.Get("$[0].sls[0].start_offset_mode", d.Get("source"))
			if startOffsetMode3 != nil && startOffsetMode3 != "" {
				sls["StartOffsetMode"] = startOffsetMode3
			}
			project1, _ := jsonpath.Get("$[0].sls[0].project", d.Get("source"))
			if project1 != nil && project1 != "" {
				sls["Project"] = project1
			}
			enableCrossAccount3, _ := jsonpath.Get("$[0].sls[0].enable_cross_account", d.Get("source"))
			if enableCrossAccount3 != nil && enableCrossAccount3 != "" {
				sls["EnableCrossAccount"] = enableCrossAccount3
			}
			logStore1, _ := jsonpath.Get("$[0].sls[0].log_store", d.Get("source"))
			if logStore1 != nil && logStore1 != "" {
				sls["LogStore"] = logStore1
			}
			acrossUid3, _ := jsonpath.Get("$[0].sls[0].across_uid", d.Get("source"))
			if acrossUid3 != nil && acrossUid3 != "" {
				sls["AcrossUid"] = acrossUid3
			}

			if len(sls) > 0 {
				source["Sls"] = sls
			}

			sourceJson, err := json.Marshal(source)
			if err != nil {
				return WrapError(err)
			}
			request["Source"] = string(sourceJson)
		}
	}

	if !d.IsNewResource() && d.HasChange("job_config") {
		update = true
		jobConfig := make(map[string]interface{})

		if v := d.Get("job_config"); v != nil {
			checkpointInterval1, _ := jsonpath.Get("$[0].checkpoint_interval", v)
			if checkpointInterval1 != nil && checkpointInterval1 != "" {
				jobConfig["CheckpointInterval"] = checkpointInterval1
			}
			maxComputeUnit1, _ := jsonpath.Get("$[0].max_compute_unit", v)
			if maxComputeUnit1 != nil && maxComputeUnit1 != "" {
				jobConfig["MaxComputeUnit"] = maxComputeUnit1
			}
			resourceGroup1, _ := jsonpath.Get("$[0].resource_group", v)
			if resourceGroup1 != nil && resourceGroup1 != "" {
				jobConfig["ResourceGroup"] = resourceGroup1
			}
			enableServerless1, _ := jsonpath.Get("$[0].enable_serverless", v)
			if enableServerless1 != nil && enableServerless1 != "" {
				jobConfig["EnableServerless"] = enableServerless1
			}
			minComputeUnit1, _ := jsonpath.Get("$[0].min_compute_unit", v)
			if minComputeUnit1 != nil && minComputeUnit1 != "" {
				jobConfig["MinComputeUnit"] = minComputeUnit1
			}

			jobConfigJson, err := json.Marshal(jobConfig)
			if err != nil {
				return WrapError(err)
			}
			request["JobConfig"] = string(jobConfigJson)
		}
	}

	if d.HasChange("sink") {
		update = true
		sink := make(map[string]interface{})

		if v := d.Get("sink"); v != nil {
			warehouse := make(map[string]interface{})
			password1, _ := jsonpath.Get("$[0].warehouse[0].password", d.Get("sink"))
			if password1 != nil && password1 != "" {
				warehouse["Password"] = password1
			}
			enableIpWhiteList1, _ := jsonpath.Get("$[0].warehouse[0].enable_ip_white_list", d.Get("sink"))
			if enableIpWhiteList1 != nil && enableIpWhiteList1 != "" {
				warehouse["EnableIpWhiteList"] = enableIpWhiteList1
			}
			unixTimestampConvertFormat1, _ := jsonpath.Get("$[0].warehouse[0].unix_timestamp_convert_format", d.Get("sink"))
			if unixTimestampConvertFormat1 != nil && unixTimestampConvertFormat1 != "" {
				warehouse["UnixTimestampConvertFormat"] = unixTimestampConvertFormat1
			}
			enableUnixTimestampConvert1, _ := jsonpath.Get("$[0].warehouse[0].enable_unix_timestamp_convert", d.Get("sink"))
			if enableUnixTimestampConvert1 != nil && enableUnixTimestampConvert1 != "" {
				warehouse["EnableUnixTimestampConvert"] = enableUnixTimestampConvert1
			}
			userName1, _ := jsonpath.Get("$[0].warehouse[0].user_name", d.Get("sink"))
			if userName1 != nil && userName1 != "" {
				warehouse["UserName"] = userName1
			}
			tableCreationMode1, _ := jsonpath.Get("$[0].warehouse[0].table_creation_mode", d.Get("sink"))
			if tableCreationMode1 != nil && tableCreationMode1 != "" {
				warehouse["TableCreationMode"] = tableCreationMode1
			}

			if len(warehouse) > 0 {
				sink["Warehouse"] = warehouse
			}
			lake := make(map[string]interface{})
			localData, err := jsonpath.Get("$[0].lake[0].partition_specs", v)
			if err != nil {
				localData = make([]interface{}, 0)
			}
			localMaps := make([]interface{}, 0)
			for _, dataLoop := range convertToInterfaceArray(localData) {
				dataLoopTmp := make(map[string]interface{})
				if dataLoop != nil {
					dataLoopTmp = dataLoop.(map[string]interface{})
				}
				dataLoopMap := make(map[string]interface{})
				dataLoopMap["TruncateWidth"] = dataLoopTmp["truncate_width"]
				dataLoopMap["TargetColumn"] = dataLoopTmp["target_column"]
				dataLoopMap["SourceTypeFormat"] = dataLoopTmp["source_type_format"]
				dataLoopMap["BucketNum"] = dataLoopTmp["bucket_num"]
				dataLoopMap["TargetTypeFormat"] = dataLoopTmp["target_type_format"]
				dataLoopMap["Strategy"] = dataLoopTmp["strategy"]
				dataLoopMap["SourceColumn"] = dataLoopTmp["source_column"]
				localMaps = append(localMaps, dataLoopMap)
			}
			lake["PartitionSpecs"] = localMaps

			tableFormat1, _ := jsonpath.Get("$[0].lake[0].table_format", d.Get("sink"))
			if tableFormat1 != nil && tableFormat1 != "" {
				lake["TableFormat"] = tableFormat1
			}
			iceberg := make(map[string]interface{})
			writeDistribution1, _ := jsonpath.Get("$[0].lake[0].iceberg[0].write_distribution", d.Get("sink"))
			if writeDistribution1 != nil && writeDistribution1 != "" {
				iceberg["WriteDistribution"] = writeDistribution1
			}
			formatVersion1, _ := jsonpath.Get("$[0].lake[0].iceberg[0].format_version", d.Get("sink"))
			if formatVersion1 != nil && formatVersion1 != "" {
				iceberg["FormatVersion"] = formatVersion1
			}
			primaryKey1, _ := jsonpath.Get("$[0].lake[0].iceberg[0].primary_key", d.Get("sink"))
			if primaryKey1 != nil && primaryKey1 != "" {
				iceberg["PrimaryKey"] = primaryKey1
			}

			if len(iceberg) > 0 {
				lake["Iceberg"] = iceberg
			}
			paimon := make(map[string]interface{})
			tableType1, _ := jsonpath.Get("$[0].lake[0].paimon[0].table_type", d.Get("sink"))
			if tableType1 != nil && tableType1 != "" {
				paimon["TableType"] = tableType1
			}
			primaryKey3, _ := jsonpath.Get("$[0].lake[0].paimon[0].primary_key", d.Get("sink"))
			if primaryKey3 != nil && primaryKey3 != "" {
				paimon["PrimaryKey"] = primaryKey3
			}

			if len(paimon) > 0 {
				lake["Paimon"] = paimon
			}
			tableCreationMode3, _ := jsonpath.Get("$[0].lake[0].table_creation_mode", d.Get("sink"))
			if tableCreationMode3 != nil && tableCreationMode3 != "" {
				lake["TableCreationMode"] = tableCreationMode3
			}
			storageTarget1, _ := jsonpath.Get("$[0].lake[0].storage_target", d.Get("sink"))
			if storageTarget1 != nil && storageTarget1 != "" {
				lake["StorageTarget"] = storageTarget1
			}

			if len(lake) > 0 {
				sink["Lake"] = lake
			}

			sinkJson, err := json.Marshal(sink)
			if err != nil {
				return WrapError(err)
			}
			request["Sink"] = string(sinkJson)
		}
	}

	if !d.IsNewResource() && d.HasChange("transform") {
		update = true
		transform := make(map[string]interface{})

		if v := d.Get("transform"); v != nil {
			localData1, err := jsonpath.Get("$[0].column_map", v)
			if err != nil {
				localData1 = make([]interface{}, 0)
			}
			localMaps1 := make([]interface{}, 0)
			for _, dataLoop1 := range convertToInterfaceArray(localData1) {
				dataLoop1Tmp := make(map[string]interface{})
				if dataLoop1 != nil {
					dataLoop1Tmp = dataLoop1.(map[string]interface{})
				}
				dataLoop1Map := make(map[string]interface{})
				dataLoop1Map["Type"] = dataLoop1Tmp["type"]
				dataLoop1Map["MapName"] = dataLoop1Tmp["map_name"]
				dataLoop1Map["MapType"] = dataLoop1Tmp["map_type"]
				dataLoop1Map["Name"] = dataLoop1Tmp["name"]
				localMaps1 = append(localMaps1, dataLoop1Map)
			}
			transform["ColumnMap"] = localMaps1

			dirtyDataHandleMode1, _ := jsonpath.Get("$[0].dirty_data_handle_mode", v)
			if dirtyDataHandleMode1 != nil && dirtyDataHandleMode1 != "" {
				transform["DirtyDataHandleMode"] = dirtyDataHandleMode1
			}

			transformJson, err := json.Marshal(transform)
			if err != nil {
				return WrapError(err)
			}
			request["Transform"] = string(transformJson)
		}
	}

	if !d.IsNewResource() && d.HasChange("advanced_config") {
		update = true
		request["AdvancedConfig"] = d.Get("advanced_config")
	}

	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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

	return resourceAliCloudAdbPipelineJobRead(d, meta)
}

func resourceAliCloudAdbPipelineJobDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeletePipelineJob"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["DBClusterId"] = parts[0]
	request["PipelineJobId"] = parts[1]
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
		if IsExpectedErrors(err, []string{"InvalidInput"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
