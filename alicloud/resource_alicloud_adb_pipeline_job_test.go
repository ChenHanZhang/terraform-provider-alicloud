// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Adb PipelineJob. >>> Resource test cases, automatically generated.
// Case resource_PipelineJob_sls_iceberg_external 12571
func TestAccAliCloudAdbPipelineJob_basic12571(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12571)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12571)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "SLS to ADB Iceberg External Lake (OSS) Pipeline Job",
					"sink": []map[string]interface{}{
						{
							"type": "OSS",
							"lake": []map[string]interface{}{
								{
									"table_name": "test_table_sls_iceberg_external",
									"partition_specs": []map[string]interface{}{
										{
											"source_column": "project",
											"strategy":      "Identity",
											"target_column": "project",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"format_version":     "2",
										},
									},
									"db_name":        "test_db_sls_iceberg_external",
									"storage_target": "oss://adb-test-bucket/sls-iceberg-external/",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_sls_iceberg_external.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "STRING",
									"map_name": "unique_id",
									"map_type": "STRING",
									"name":     "unique_id",
								},
								{
									"type":     "STRING",
									"map_name": "consumer_group",
									"map_type": "STRING",
									"name":     "consumer_group",
								},
								{
									"type":     "STRING",
									"map_name": "project",
									"map_type": "STRING",
									"name":     "project",
								},
								{
									"type":     "STRING",
									"map_name": "fallbehind",
									"map_type": "STRING",
									"name":     "fallbehind",
								},
								{
									"type":     "STRING",
									"map_name": "shard",
									"map_type": "STRING",
									"name":     "shard",
								},
								{
									"type":     "STRING",
									"map_name": "logstore",
									"map_type": "STRING",
									"name":     "logstore",
								},
							},
							"dirty_data_handle_mode": "ERROR",
						},
					},
					"source": []map[string]interface{}{
						{
							"type": "SLS",
							"sls": []map[string]interface{}{
								{
									"start_offset_mode":    "earliest",
									"project":              "${alicloud_log_project.create_sls_project_iceberg_external.id}",
									"enable_cross_account": "false",
									"log_store":            "${alicloud_log_store.create_sls_logstore_iceberg_external.name}",
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":   "SLS to ADB Iceberg External Lake (OSS) Pipeline Job",
						"db_cluster_id": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12571 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12571(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_sls_iceberg_external" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "sls-iceberg-external-vpc"
}

resource "alicloud_log_project" "create_sls_project_iceberg_external" {
  description = "SLS project for Iceberg external lake test"
  name        = "sls-iceberg-external-project"
}

resource "alicloud_vswitch" "createVswitch_sls_iceberg_external" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_sls_iceberg_external.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.5.0/24"
  vswitch_name = "sls-iceberg-external-vsw"
}

resource "alicloud_log_store" "create_sls_logstore_iceberg_external" {
  retention_period = "200"
  shard_count      = "2"
  project          = alicloud_log_project.create_sls_project_iceberg_external.id
  name             = "sls-iceberg-external-logstore"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_sls_iceberg_external" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_sls_iceberg_external.id
  vswitch_id                    = alicloud_vswitch.createVswitch_sls_iceberg_external.id
  db_cluster_description        = "sls_iceberg_external_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}


`, name)
}

// Case resource_PipelineJob_sls_iceberg_external_modify 12572
func TestAccAliCloudAdbPipelineJob_basic12572(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12572)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12572)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "SLS to ADB Iceberg External Modify Lake (OSS) Pipeline Job",
					"sink": []map[string]interface{}{
						{
							"type": "OSS",
							"lake": []map[string]interface{}{
								{
									"table_name": "test_table_sls_iceberg_external_modify",
									"partition_specs": []map[string]interface{}{
										{
											"source_column": "project",
											"strategy":      "Identity",
											"target_column": "project",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"format_version":     "2",
										},
									},
									"db_name":        "test_db_sls_iceberg_external_modify",
									"storage_target": "oss://adb-test-bucket/sls-iceberg-external-modify/",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_sls_iceberg_external_modify.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "STRING",
									"map_name": "unique_id",
									"map_type": "STRING",
									"name":     "unique_id",
								},
								{
									"type":     "STRING",
									"map_name": "consumer_group",
									"map_type": "STRING",
									"name":     "consumer_group",
								},
								{
									"type":     "STRING",
									"map_name": "project",
									"map_type": "STRING",
									"name":     "project",
								},
								{
									"type":     "STRING",
									"map_name": "fallbehind",
									"map_type": "STRING",
									"name":     "fallbehind",
								},
								{
									"type":     "STRING",
									"map_name": "shard",
									"map_type": "STRING",
									"name":     "shard",
								},
								{
									"type":     "STRING",
									"map_name": "logstore",
									"map_type": "STRING",
									"name":     "logstore",
								},
							},
							"dirty_data_handle_mode": "ERROR",
						},
					},
					"source": []map[string]interface{}{
						{
							"type": "SLS",
							"sls": []map[string]interface{}{
								{
									"start_offset_mode":    "earliest",
									"project":              "${alicloud_log_project.create_sls_project_iceberg_external_modify.id}",
									"enable_cross_account": "false",
									"log_store":            "${alicloud_log_store.create_sls_logstore_iceberg_external_modify.name}",
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":   "SLS to ADB Iceberg External Modify Lake (OSS) Pipeline Job",
						"db_cluster_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sink": []map[string]interface{}{
						{
							"lake": []map[string]interface{}{
								{
									"partition_specs": []map[string]interface{}{
										{
											"source_column": "project",
											"strategy":      "Identity",
											"target_column": "project",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"format_version":     "2",
										},
									},
									"storage_target": "oss://adb-test-bucket/sls-iceberg-external-modify/",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"source": []map[string]interface{}{
						{
							"sls": []map[string]interface{}{
								{
									"start_offset_mode":    "earliest",
									"project":              "${alicloud_log_project.create_sls_project_iceberg_external_modify.id}",
									"enable_cross_account": "false",
									"log_store":            "${alicloud_log_store.create_sls_logstore_iceberg_external_modify.name}",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12572 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12572(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_sls_iceberg_external_modify" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "sls-iceberg-external-modify-vpc"
}

resource "alicloud_log_project" "create_sls_project_iceberg_external_modify" {
  description = "SLS project for Iceberg external modify lake test"
  name        = "sls-iceberg-external-modify-project"
}

resource "alicloud_vswitch" "createVswitch_sls_iceberg_external_modify" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_sls_iceberg_external_modify.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.15.0/24"
  vswitch_name = "sls-iceberg-external-modify-vsw"
}

resource "alicloud_log_store" "create_sls_logstore_iceberg_external_modify" {
  retention_period = "200"
  shard_count      = "2"
  project          = alicloud_log_project.create_sls_project_iceberg_external_modify.id
  name             = "sls-iceberg-external-modify-logstore"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_sls_iceberg_external_modify" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_sls_iceberg_external_modify.id
  vswitch_id                    = alicloud_vswitch.createVswitch_sls_iceberg_external_modify.id
  db_cluster_description        = "sls_iceberg_external_modify_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}


`, name)
}

// Case resource_PipelineJob_kafka_iceberg_internal 12573
func TestAccAliCloudAdbPipelineJob_basic12573(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12573)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12573)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "Kafka to ADB Iceberg Internal Lake Pipeline Job",
					"sink": []map[string]interface{}{
						{
							"type": "ADB_LAKE",
							"lake": []map[string]interface{}{
								{
									"table_name": "test_table_kafka_iceberg_internal",
									"partition_specs": []map[string]interface{}{
										{
											"target_type_format": "yyyyMMdd",
											"source_column":      "__timestamp__",
											"strategy":           "ParseAsTimeAndFormat",
											"target_column":      "dt",
											"source_type_format": "APSLiteralTimestampMilliSecond",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"format_version":     "2",
										},
									},
									"db_name":        "test_db_kafka_iceberg_internal",
									"storage_target": "${alicloud_adb_lake_storage.create_lake_storage_kafka_iceberg_internal.lake_storage_id}",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_kafka_iceberg_internal.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "BIGINT",
									"map_name": "id",
									"map_type": "BIGINT",
									"name":     "id",
								},
								{
									"type":     "STRING",
									"map_name": "name",
									"map_type": "STRING",
									"name":     "name",
								},
								{
									"type":     "DOUBLE",
									"map_name": "value",
									"map_type": "DOUBLE",
									"name":     "value",
								},
								{
									"type":     "STRING",
									"map_name": "__key__",
									"map_type": "STRING",
									"name":     "__key__",
								},
								{
									"type":     "STRING",
									"map_name": "__topic__",
									"map_type": "STRING",
									"name":     "__topic__",
								},
								{
									"type":     "INT",
									"map_name": "__partition__",
									"map_type": "INT",
									"name":     "__partition__",
								},
								{
									"type":     "LONG",
									"map_name": "__offset__",
									"map_type": "LONG",
									"name":     "__offset__",
								},
								{
									"type":     "LONG",
									"map_name": "__timestamp__",
									"map_type": "TIMESTAMP",
									"name":     "__timestamp__",
								},
							},
							"dirty_data_handle_mode": "ERROR",
						},
					},
					"advanced_config": "aps.flink.target-file-size-bytes=134217728",
					"source": []map[string]interface{}{
						{
							"type": "KAFKA",
							"kafka": []map[string]interface{}{
								{
									"start_offset_mode": "earliest",
									"message_format":    "json",
									"cloud_managed": []map[string]interface{}{
										{
											"enable_cross_account": "false",
										},
									},
									"kafka_topic":      "${alicloud_alikafka_topic.create_kafka_topic_iceberg_internal.topic}",
									"instance_type":    "CLOUD_MANAGED",
									"kafka_cluster_id": "${ alicloud_alikafka_instance.create_kafka_instance_iceberg_internal.id}",
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":     "Kafka to ADB Iceberg Internal Lake Pipeline Job",
						"db_cluster_id":   CHECKSET,
						"advanced_config": "aps.flink.target-file-size-bytes=134217728",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12573 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12573(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_kafka_iceberg_internal" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "kafka-iceberg-internal-vpc"
}

resource "alicloud_vswitch" "createVswitch_kafka_iceberg_internal" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_kafka_iceberg_internal.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.1.0/24"
  vswitch_name = "kafka-iceberg-internal-vsw"
}

resource " alicloud_alikafka_instance" "create_kafka_instance_iceberg_internal" {
  deploy_type     = "5"
  disk_type       = "0"
  vpc_id          = alicloud_vpc.createVpc_kafka_iceberg_internal.id
  spec_type       = "professional"
  paid_type       = "1"
  vswitch_id      = alicloud_vswitch.createVswitch_kafka_iceberg_internal.id
  io_max_spec     = "alikafka.hw.2xlarge"
  disk_size       = "500"
  service_version = "2.2.0"
  deploy_module   = "vpc"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_kafka_iceberg_internal" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_kafka_iceberg_internal.id
  vswitch_id                    = alicloud_vswitch.createVswitch_kafka_iceberg_internal.id
  db_cluster_description        = "kafka_iceberg_internal_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}

resource "alicloud_alikafka_topic" "create_kafka_topic_iceberg_internal" {
  instance_id = alicloud_alikafka_instance.create_kafka_instance_iceberg_internal.id
  config      = "max.message.bytes"
  topic       = "kafka-iceberg-internal-topic"
  local_topic = true
  remark      = "kafka-iceberg-internal-topic-remark"
}

resource "alicloud_adb_lake_storage" "create_lake_storage_kafka_iceberg_internal" {
  description   = "kafka_iceberg_internal_lake_storage"
  db_cluster_id = alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_kafka_iceberg_internal.id
}


`, name)
}

// Case resource_PipelineJob_sls_iceberg_internal_modify 12574
func TestAccAliCloudAdbPipelineJob_basic12574(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12574)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12574)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "SLS to ADB Iceberg Internal Lake Pipeline Job modify",
					"sink": []map[string]interface{}{
						{
							"type": "ADB_LAKE",
							"lake": []map[string]interface{}{
								{
									"table_name": "test_table_sls_iceberg_internal_modify",
									"partition_specs": []map[string]interface{}{
										{
											"source_column": "project",
											"strategy":      "Identity",
											"target_column": "project",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"format_version":     "2",
										},
									},
									"db_name":        "test_db_sls_iceberg_internal_modify",
									"storage_target": "${alicloud_adb_lake_storage.create_lake_storage_sls_iceberg_internal_modify.lake_storage_id}",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_sls_iceberg_internal_modify.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "STRING",
									"map_name": "unique_id",
									"map_type": "STRING",
									"name":     "unique_id",
								},
								{
									"type":     "STRING",
									"map_name": "consumer_group",
									"map_type": "STRING",
									"name":     "consumer_group",
								},
								{
									"type":     "STRING",
									"map_name": "project",
									"map_type": "STRING",
									"name":     "project",
								},
								{
									"type":     "STRING",
									"map_name": "fallbehind",
									"map_type": "STRING",
									"name":     "fallbehind",
								},
								{
									"type":     "STRING",
									"map_name": "shard",
									"map_type": "STRING",
									"name":     "shard",
								},
								{
									"type":     "STRING",
									"map_name": "logstore",
									"map_type": "STRING",
									"name":     "logstore",
								},
							},
							"dirty_data_handle_mode": "ERROR",
						},
					},
					"source": []map[string]interface{}{
						{
							"type": "SLS",
							"sls": []map[string]interface{}{
								{
									"start_offset_mode":  "TIMESTAMP",
									"project":            "${alicloud_log_project.create_sls_project_iceberg_internal_modify.id}",
									"log_store":          "${alicloud_log_store.create_sls_logstore_iceberg_internal_modify.name}",
									"start_offset_value": "1735689600000",
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":   "SLS to ADB Iceberg Internal Lake Pipeline Job modify",
						"db_cluster_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sink": []map[string]interface{}{
						{
							"lake": []map[string]interface{}{
								{
									"partition_specs": []map[string]interface{}{
										{
											"source_column": "project",
											"strategy":      "Identity",
											"target_column": "project",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"format_version":     "2",
										},
									},
									"storage_target": "${alicloud_adb_lake_storage.create_lake_storage_sls_iceberg_internal_modify.lake_storage_id}",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"source": []map[string]interface{}{
						{
							"sls": []map[string]interface{}{
								{
									"start_offset_mode":  "TIMESTAMP",
									"project":            "${alicloud_log_project.create_sls_project_iceberg_internal_modify.id}",
									"log_store":          "${alicloud_log_store.create_sls_logstore_iceberg_internal_modify.name}",
									"start_offset_value": "1735689600000",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12574 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12574(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_sls_iceberg_internal_modify" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "sls-iceberg-internal-vpc-modify"
}

resource "alicloud_vswitch" "createVswitch_sls_iceberg_internal_modify" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_sls_iceberg_internal_modify.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.14.0/24"
  vswitch_name = "sls-iceberg-internal-vsw-modify"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_sls_iceberg_internal_modify" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_sls_iceberg_internal_modify.id
  vswitch_id                    = alicloud_vswitch.createVswitch_sls_iceberg_internal_modify.id
  db_cluster_description        = "sls_iceberg_internal_cluster_modify"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}

resource "alicloud_log_project" "create_sls_project_iceberg_internal_modify" {
  description = "SLS project for Iceberg internal lake test modify"
  name        = "sls-iceberg-internal-project-modify"
}

resource "alicloud_adb_lake_storage" "create_lake_storage_sls_iceberg_internal_modify" {
  description   = "sls_iceberg_internal_lake_storage_modify"
  db_cluster_id = alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_sls_iceberg_internal_modify.id
}

resource "alicloud_log_store" "create_sls_logstore_iceberg_internal_modify" {
  retention_period = "200"
  shard_count      = "2"
  project          = alicloud_log_project.create_sls_project_iceberg_internal_modify.id
  name             = "sls-iceberg-internal-logstore-modify"
}


`, name)
}

// Case resource_PipelineJob_sls_paimon_external 12575
func TestAccAliCloudAdbPipelineJob_basic12575(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12575)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12575)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "SLS to ADB Paimon External Lake (OSS) Pipeline Job",
					"sink": []map[string]interface{}{
						{
							"type": "OSS",
							"lake": []map[string]interface{}{
								{
									"table_name": "test_table_sls_paimon_external",
									"partition_specs": []map[string]interface{}{
										{
											"source_column": "project",
											"strategy":      "Identity",
											"target_column": "project",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"db_name":             "test_db_sls_paimon_external",
									"storage_target":      "oss://adb-test-bucket/sls-paimon-external/",
									"table_format":        "APACHE_PAIMON",
									"paimon": []map[string]interface{}{
										{
											"table_type": "TABLE_WITHOUT_PK",
										},
									},
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_sls_paimon_external.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "STRING",
									"map_name": "unique_id",
									"map_type": "STRING",
									"name":     "unique_id",
								},
								{
									"type":     "STRING",
									"map_name": "consumer_group",
									"map_type": "STRING",
									"name":     "consumer_group",
								},
								{
									"type":     "STRING",
									"map_name": "project",
									"map_type": "STRING",
									"name":     "project",
								},
								{
									"type":     "STRING",
									"map_name": "fallbehind",
									"map_type": "STRING",
									"name":     "fallbehind",
								},
								{
									"type":     "STRING",
									"map_name": "shard",
									"map_type": "STRING",
									"name":     "shard",
								},
								{
									"type":     "STRING",
									"map_name": "logstore",
									"map_type": "STRING",
									"name":     "logstore",
								},
							},
							"dirty_data_handle_mode": "ERROR",
						},
					},
					"source": []map[string]interface{}{
						{
							"type": "SLS",
							"sls": []map[string]interface{}{
								{
									"start_offset_mode":    "earliest",
									"project":              "${alicloud_log_project.create_sls_project_paimon_external.id}",
									"enable_cross_account": "false",
									"log_store":            "${alicloud_log_store.create_sls_logstore_paimon_external.name}",
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":   "SLS to ADB Paimon External Lake (OSS) Pipeline Job",
						"db_cluster_id": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12575 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12575(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_sls_paimon_external" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "sls-paimon-external-vpc"
}

resource "alicloud_log_project" "create_sls_project_paimon_external" {
  description = "SLS project for Paimon external lake test"
  name        = "sls-paimon-external-project"
}

resource "alicloud_vswitch" "createVswitch_sls_paimon_external" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_sls_paimon_external.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.6.0/24"
  vswitch_name = "sls-paimon-external-vsw"
}

resource "alicloud_log_store" "create_sls_logstore_paimon_external" {
  retention_period = "200"
  shard_count      = "2"
  project          = alicloud_log_project.create_sls_project_paimon_external.id
  name             = "sls-paimon-external-logstore"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_sls_paimon_external" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_sls_paimon_external.id
  vswitch_id                    = alicloud_vswitch.createVswitch_sls_paimon_external.id
  db_cluster_description        = "sls_paimon_external_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}


`, name)
}

// Case resource_PipelineJob_kafka_paimon_external_modify 12576
func TestAccAliCloudAdbPipelineJob_basic12576(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12576)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12576)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "Kafka to ADB Paimon External Lake (OSS) Pipeline Job - modify",
					"sink": []map[string]interface{}{
						{
							"type": "OSS",
							"lake": []map[string]interface{}{
								{
									"table_name": "test_table_kafka_paimon_external_modify",
									"partition_specs": []map[string]interface{}{
										{
											"target_type_format": "yyyyMMdd",
											"source_column":      "__timestamp__",
											"strategy":           "ParseAsTimeAndFormat",
											"target_column":      "dt",
											"source_type_format": "APSLiteralTimestampMilliSecond",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"db_name":             "test_db_kafka_paimon_external_modify",
									"storage_target":      "oss://adb-test-bucket/kafka-paimon-external-modify/",
									"table_format":        "APACHE_PAIMON",
									"paimon": []map[string]interface{}{
										{
											"table_type": "TABLE_WITHOUT_PK",
										},
									},
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_kafka_paimon_external_modify.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "BIGINT",
									"map_name": "id",
									"map_type": "BIGINT",
									"name":     "id",
								},
								{
									"type":     "STRING",
									"map_name": "name",
									"map_type": "STRING",
									"name":     "name",
								},
								{
									"type":     "DOUBLE",
									"map_name": "value",
									"map_type": "DOUBLE",
									"name":     "value",
								},
								{
									"type":     "STRING",
									"map_name": "__key__",
									"map_type": "STRING",
									"name":     "__key__",
								},
								{
									"type":     "STRING",
									"map_name": "__topic__",
									"map_type": "STRING",
									"name":     "__topic__",
								},
								{
									"type":     "INT",
									"map_name": "__partition__",
									"map_type": "INT",
									"name":     "__partition__",
								},
								{
									"type":     "LONG",
									"map_name": "__offset__",
									"map_type": "LONG",
									"name":     "__offset__",
								},
								{
									"type":     "LONG",
									"map_name": "__timestamp__",
									"map_type": "TIMESTAMP",
									"name":     "__timestamp__",
								},
							},
							"dirty_data_handle_mode": "ERROR",
						},
					},
					"source": []map[string]interface{}{
						{
							"type": "KAFKA",
							"kafka": []map[string]interface{}{
								{
									"start_offset_mode": "earliest",
									"message_format":    "json",
									"cloud_managed": []map[string]interface{}{
										{
											"enable_cross_account": "false",
										},
									},
									"kafka_topic":      "${alicloud_alikafka_topic.create_kafka_topic_paimon_external_modify.topic}",
									"instance_type":    "CLOUD_MANAGED",
									"kafka_cluster_id": "${ alicloud_alikafka_instance.create_kafka_instance_paimon_external_modify.id}",
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":   "Kafka to ADB Paimon External Lake (OSS) Pipeline Job - modify",
						"db_cluster_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sink": []map[string]interface{}{
						{
							"lake": []map[string]interface{}{
								{
									"partition_specs": []map[string]interface{}{
										{
											"target_type_format": "yyyyMMdd",
											"source_column":      "__timestamp__",
											"strategy":           "ParseAsTimeAndFormat",
											"target_column":      "dt",
											"source_type_format": "APSLiteralTimestampMilliSecond",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"storage_target":      "oss://adb-test-bucket/kafka-paimon-external-modify/",
									"table_format":        "APACHE_PAIMON",
									"paimon": []map[string]interface{}{
										{
											"table_type": "TABLE_WITHOUT_PK",
										},
									},
								},
							},
						},
					},
					"source": []map[string]interface{}{
						{
							"kafka": []map[string]interface{}{
								{
									"start_offset_mode": "earliest",
									"message_format":    "json",
									"cloud_managed": []map[string]interface{}{
										{
											"enable_cross_account": "false",
										},
									},
									"kafka_topic":      "${alicloud_alikafka_topic.create_kafka_topic_paimon_external_modify.topic}",
									"kafka_cluster_id": "${ alicloud_alikafka_instance.create_kafka_instance_paimon_external_modify.id}",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12576 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12576(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_kafka_paimon_external_modify" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "kafka-paimon-external-modify-vpc"
}

resource "alicloud_vswitch" "createVswitch_kafka_paimon_external_modify" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_kafka_paimon_external_modify.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.13.0/24"
  vswitch_name = "kafka-paimon-external-modify-vsw"
}

resource " alicloud_alikafka_instance" "create_kafka_instance_paimon_external_modify" {
  deploy_type     = "5"
  disk_type       = "0"
  vpc_id          = alicloud_vpc.createVpc_kafka_paimon_external_modify.id
  spec_type       = "normal"
  paid_type       = "1"
  vswitch_id      = alicloud_vswitch.createVswitch_kafka_paimon_external_modify.id
  io_max_spec     = "alikafka.hw.2xlarge"
  disk_size       = "500"
  service_version = "2.2.0"
  deploy_module   = "vpc"
}

resource "alicloud_alikafka_topic" "create_kafka_topic_paimon_external_modify" {
  instance_id = alicloud_alikafka_instance.create_kafka_instance_paimon_external_modify.id
  config      = "{\"retention.ms\":\"36000000\"}"
  topic       = "kafka-paimon-external-modify-topic"
  remark      = "kafka-paimon-external-modify-topic-remark"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_kafka_paimon_external_modify" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_kafka_paimon_external_modify.id
  vswitch_id                    = alicloud_vswitch.createVswitch_kafka_paimon_external_modify.id
  db_cluster_description        = "kafka_paimon_external_modify_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}


`, name)
}

// Case resource_PipelineJob_kafka_paimon_external 12577
func TestAccAliCloudAdbPipelineJob_basic12577(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12577)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12577)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "Kafka to ADB Paimon External Lake (OSS) Pipeline Job",
					"sink": []map[string]interface{}{
						{
							"type": "OSS",
							"lake": []map[string]interface{}{
								{
									"table_name": "test_table_kafka_paimon_external",
									"partition_specs": []map[string]interface{}{
										{
											"target_type_format": "yyyyMMdd",
											"source_column":      "__timestamp__",
											"strategy":           "ParseAsTimeAndFormat",
											"target_column":      "dt",
											"source_type_format": "APSLiteralTimestampMilliSecond",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"db_name":             "test_db_kafka_paimon_external",
									"storage_target":      "oss://adb-test-bucket/kafka-paimon-external/",
									"table_format":        "APACHE_PAIMON",
									"paimon": []map[string]interface{}{
										{
											"table_type": "TABLE_WITHOUT_PK",
										},
									},
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_kafka_paimon_external.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "BIGINT",
									"map_name": "id",
									"map_type": "BIGINT",
									"name":     "id",
								},
								{
									"type":     "STRING",
									"map_name": "name",
									"map_type": "STRING",
									"name":     "name",
								},
								{
									"type":     "DOUBLE",
									"map_name": "value",
									"map_type": "DOUBLE",
									"name":     "value",
								},
								{
									"type":     "STRING",
									"map_name": "__key__",
									"map_type": "STRING",
									"name":     "__key__",
								},
								{
									"type":     "STRING",
									"map_name": "__topic__",
									"map_type": "STRING",
									"name":     "__topic__",
								},
								{
									"type":     "INT",
									"map_name": "__partition__",
									"map_type": "INT",
									"name":     "__partition__",
								},
								{
									"type":     "LONG",
									"map_name": "__offset__",
									"map_type": "LONG",
									"name":     "__offset__",
								},
								{
									"type":     "LONG",
									"map_name": "__timestamp__",
									"map_type": "TIMESTAMP",
									"name":     "__timestamp__",
								},
							},
							"dirty_data_handle_mode": "ERROR",
						},
					},
					"source": []map[string]interface{}{
						{
							"type": "KAFKA",
							"kafka": []map[string]interface{}{
								{
									"start_offset_mode": "earliest",
									"message_format":    "json",
									"cloud_managed": []map[string]interface{}{
										{
											"enable_cross_account": "false",
										},
									},
									"kafka_topic":      "${alicloud_alikafka_topic.create_kafka_topic_paimon_external.topic}",
									"instance_type":    "CLOUD_MANAGED",
									"kafka_cluster_id": "${ alicloud_alikafka_instance.create_kafka_instance_paimon_external.id}",
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":   "Kafka to ADB Paimon External Lake (OSS) Pipeline Job",
						"db_cluster_id": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12577 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12577(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_kafka_paimon_external" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "kafka-paimon-external-vpc"
}

resource "alicloud_vswitch" "createVswitch_kafka_paimon_external" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_kafka_paimon_external.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.3.0/24"
  vswitch_name = "kafka-paimon-external-vsw"
}

resource " alicloud_alikafka_instance" "create_kafka_instance_paimon_external" {
  deploy_type     = "5"
  disk_type       = "0"
  vpc_id          = alicloud_vpc.createVpc_kafka_paimon_external.id
  spec_type       = "normal"
  paid_type       = "1"
  vswitch_id      = alicloud_vswitch.createVswitch_kafka_paimon_external.id
  io_max_spec     = "alikafka.hw.2xlarge"
  disk_size       = "500"
  service_version = "2.2.0"
  deploy_module   = "vpc"
}

resource "alicloud_alikafka_topic" "create_kafka_topic_paimon_external" {
  instance_id = alicloud_alikafka_instance.create_kafka_instance_paimon_external.id
  config      = "{\"retention.ms\":\"36000000\"}"
  topic       = "kafka-paimon-external-topic"
  remark      = "kafka-paimon-external-topic-remark"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_kafka_paimon_external" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_kafka_paimon_external.id
  vswitch_id                    = alicloud_vswitch.createVswitch_kafka_paimon_external.id
  db_cluster_description        = "kafka_paimon_external_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}


`, name)
}

// Case resource_PipelineJob_sls_warehouse_modify 12578
func TestAccAliCloudAdbPipelineJob_basic12578(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12578)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12578)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "SLS to ADB Warehouse Pipeline Job modify",
					"sink": []map[string]interface{}{
						{
							"type": "ADB_WAREHOUSE",
							"warehouse": []map[string]interface{}{
								{
									"table_name":                    "test_table_sls_warehouse_modify",
									"enable_unix_timestamp_convert": "true",
									"user_name":                     "${alicloud_adb_lake_account.create_lake_account_sls_warehouse_modify.account_name}",
									"table_creation_mode":           "USE_EXISTING",
									"enable_ip_white_list":          "true",
									"unix_timestamp_convert_format": "APSLiteralTimestampMilliSecond",
									"db_name":                       "test_db_sls_warehouse_modify",
									"password":                      "${alicloud_adb_lake_account.create_lake_account_sls_warehouse_modify.account_password}",
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_sls_warehouse_modify.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "STRING",
									"map_name": "unique_id",
									"map_type": "STRING",
									"name":     "unique_id",
								},
								{
									"type":     "STRING",
									"map_name": "consumer_group",
									"map_type": "STRING",
									"name":     "consumer_group",
								},
								{
									"type":     "STRING",
									"map_name": "project",
									"map_type": "STRING",
									"name":     "project",
								},
								{
									"type":     "STRING",
									"map_name": "fallbehind",
									"map_type": "STRING",
									"name":     "fallbehind",
								},
								{
									"type":     "STRING",
									"map_name": "shard",
									"map_type": "STRING",
									"name":     "shard",
								},
								{
									"type":     "STRING",
									"map_name": "logstore",
									"map_type": "STRING",
									"name":     "logstore",
								},
							},
							"dirty_data_handle_mode": "TREAT_AS_NULL",
						},
					},
					"source": []map[string]interface{}{
						{
							"type": "SLS",
							"sls": []map[string]interface{}{
								{
									"start_offset_mode":  "TIMESTAMP",
									"project":            "${alicloud_log_project.create_sls_project_warehouse_modify.id}",
									"log_store":          "${alicloud_log_store.create_sls_logstore_warehouse_modify.name}",
									"start_offset_value": "1735689600000",
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":   "SLS to ADB Warehouse Pipeline Job modify",
						"db_cluster_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sink": []map[string]interface{}{
						{
							"warehouse": []map[string]interface{}{
								{
									"enable_unix_timestamp_convert": "true",
									"user_name":                     "${alicloud_adb_lake_account.create_lake_account_sls_warehouse_modify.account_name}",
									"table_creation_mode":           "USE_EXISTING",
									"enable_ip_white_list":          "true",
									"unix_timestamp_convert_format": "APSLiteralTimestampMilliSecond",
									"password":                      "${alicloud_adb_lake_account.create_lake_account_sls_warehouse_modify.account_password}",
								},
							},
						},
					},
					"source": []map[string]interface{}{
						{
							"sls": []map[string]interface{}{
								{
									"start_offset_mode":  "TIMESTAMP",
									"project":            "${alicloud_log_project.create_sls_project_warehouse_modify.id}",
									"log_store":          "${alicloud_log_store.create_sls_logstore_warehouse_modify.name}",
									"start_offset_value": "1735689600000",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12578 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12578(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_sls_warehouse_modify" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "sls-warehouse-vpc-modify"
}

resource "alicloud_vswitch" "createVswitch_sls_warehouse_modify" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_sls_warehouse_modify.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.18.0/24"
  vswitch_name = "sls-warehouse-vsw-modify"
}

resource "alicloud_log_project" "create_sls_project_warehouse_modify" {
  description = "SLS project for ADB Warehouse test modify"
  name        = "sls-warehouse-project-modify"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_sls_warehouse_modify" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_sls_warehouse_modify.id
  vswitch_id                    = alicloud_vswitch.createVswitch_sls_warehouse_modify.id
  db_cluster_description        = "sls_warehouse_cluster_modify"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}

resource "alicloud_log_store" "create_sls_logstore_warehouse_modify" {
  retention_period = "200"
  shard_count      = "2"
  project          = alicloud_log_project.create_sls_project_warehouse_modify.id
  name             = "sls-warehouse-logstore-modify"
}

resource "alicloud_adb_lake_account" "create_lake_account_sls_warehouse_modify" {
  account_description = "Account for SLS Warehouse pipeline test modify"
  db_cluster_id       = alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_sls_warehouse_modify.id
  account_type        = "Super"
  engine              = "AnalyticDB"
  account_name        = "sls_warehouse_user_modify"
  account_password    = "Test@12345678"
}


`, name)
}

// Case resource_PipelineJob_kafka_iceberg_internal_cdc_dataworks 12579
func TestAccAliCloudAdbPipelineJob_basic12579(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12579)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12579)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "Kafka CDC Dataworks JSON to ADB Iceberg Internal Lake Pipeline Job",
					"sink": []map[string]interface{}{
						{
							"type": "ADB_LAKE",
							"lake": []map[string]interface{}{
								{
									"table_name": "test_table_kafka_iceberg_internal_cdc_dataworks",
									"partition_specs": []map[string]interface{}{
										{
											"source_column":  "VARCHAR_TITLE",
											"strategy":       "Truncate",
											"target_column":  "VARCHAR_TITLE_TRUNCATE",
											"truncate_width": "4",
										},
										{
											"source_column": "NATIONAL_CHAR_NICK",
											"strategy":      "Bucket",
											"target_column": "NATIONAL_CHAR_NICK_BUCKET",
											"bucket_num":    "4",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"primary_key":        "UID",
											"format_version":     "2",
										},
									},
									"db_name":        "test_db_kafka_iceberg_internal_cdc_dataworks",
									"storage_target": "${alicloud_adb_lake_storage.create_lake_storage_kafka_iceberg_internal_cdc_dataworks.lake_storage_id}",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_kafka_iceberg_internal_cdc_dataworks.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "INT",
									"map_name": "UID",
									"map_type": "INT",
									"name":     "UID",
								},
								{
									"type":     "DOUBLE",
									"map_name": "EVENT_ID",
									"map_type": "DOUBLE",
									"name":     "EVENT_ID",
								},
								{
									"type":     "DOUBLE",
									"map_name": "PLAYER_ID",
									"map_type": "DOUBLE",
									"name":     "PLAYER_ID",
								},
								{
									"type":     "DOUBLE",
									"map_name": "INTEGER_SCORE",
									"map_type": "DOUBLE",
									"name":     "INTEGER_SCORE",
								},
								{
									"type":     "DOUBLE",
									"map_name": "INT_RANK",
									"map_type": "DOUBLE",
									"name":     "INT_RANK",
								},
								{
									"type":     "DOUBLE",
									"map_name": "SMALL_LEVEL",
									"map_type": "DOUBLE",
									"name":     "SMALL_LEVEL",
								},
								{
									"type":     "DOUBLE",
									"map_name": "NUMBER_LIVES",
									"map_type": "DOUBLE",
									"name":     "NUMBER_LIVES",
								},
								{
									"type":     "DOUBLE",
									"map_name": "NUMERIC_PRICE",
									"map_type": "DOUBLE",
									"name":     "NUMERIC_PRICE",
								},
								{
									"type":     "DOUBLE",
									"map_name": "DECIMAL_DISCOUNT",
									"map_type": "DOUBLE",
									"name":     "DECIMAL_DISCOUNT",
								},
								{
									"type":     "DOUBLE",
									"map_name": "FLOAT_ACCURACY",
									"map_type": "DOUBLE",
									"name":     "FLOAT_ACCURACY",
								},
								{
									"type":     "DOUBLE",
									"map_name": "DOUBLE_POWER",
									"map_type": "DOUBLE",
									"name":     "DOUBLE_POWER",
								},
								{
									"type":     "DOUBLE",
									"map_name": "REAL_DROP",
									"map_type": "DOUBLE",
									"name":     "REAL_DROP",
								},
								{
									"type":     "STRING",
									"map_name": "CHAR_CLASS",
									"map_type": "STRING",
									"name":     "CHAR_CLASS",
								},
								{
									"type":     "STRING",
									"map_name": "NCHAR_REGION",
									"map_type": "STRING",
									"name":     "NCHAR_REGION",
								},
								{
									"type":     "STRING",
									"map_name": "VARCHAR_TITLE",
									"map_type": "STRING",
									"name":     "VARCHAR_TITLE",
								},
								{
									"type":     "STRING",
									"map_name": "VARCHAR2_ITEM",
									"map_type": "STRING",
									"name":     "VARCHAR2_ITEM",
								},
								{
									"type":     "STRING",
									"map_name": "NVARCHAR2_NPC",
									"map_type": "STRING",
									"name":     "NVARCHAR2_NPC",
								},
								{
									"type":     "STRING",
									"map_name": "CHARACTER_CODE",
									"map_type": "STRING",
									"name":     "CHARACTER_CODE",
								},
								{
									"type":     "STRING",
									"map_name": "CHAR_VARYING_TAG",
									"map_type": "STRING",
									"name":     "CHAR_VARYING_TAG",
								},
								{
									"type":     "STRING",
									"map_name": "CHARACTER_VARYING_ALIAS",
									"map_type": "STRING",
									"name":     "CHARACTER_VARYING_ALIAS",
								},
								{
									"type":     "STRING",
									"map_name": "NATIONAL_CHAR_NICK",
									"map_type": "STRING",
									"name":     "NATIONAL_CHAR_NICK",
								},
								{
									"type":     "STRING",
									"map_name": "NAT_CHAR_VARYING_COMMENT",
									"map_type": "STRING",
									"name":     "NAT_CHAR_VARYING_COMMENT",
								},
								{
									"type":     "STRING",
									"map_name": "NAT_CHARACTER_VARYING_NOTE",
									"map_type": "STRING",
									"name":     "NAT_CHARACTER_VARYING_NOTE",
								},
								{
									"type":     "STRING",
									"map_name": "NCHAR_VARYING_LABEL",
									"map_type": "STRING",
									"name":     "NCHAR_VARYING_LABEL",
								},
								{
									"type":     "STRING",
									"map_name": "CLOB_STORY",
									"map_type": "STRING",
									"name":     "CLOB_STORY",
								},
								{
									"type":     "STRING",
									"map_name": "NCLOB_DIALOG",
									"map_type": "STRING",
									"name":     "NCLOB_DIALOG",
								},
								{
									"type":     "TIMESTAMP",
									"map_name": "TIMESTAMP_EVENT",
									"map_type": "TIMESTAMP",
									"name":     "TIMESTAMP_EVENT",
								},
								{
									"type":     "TIMESTAMP",
									"map_name": "DATE_EVENT",
									"map_type": "TIMESTAMP",
									"name":     "DATE_EVENT",
								},
								{
									"type":     "DOUBLE",
									"map_name": "BIT_FLAG",
									"map_type": "DOUBLE",
									"name":     "BIT_FLAG",
								},
								{
									"type":     "DOUBLE",
									"map_name": "BOOL_FLAG",
									"map_type": "DOUBLE",
									"name":     "BOOL_FLAG",
								},
								{
									"type":     "STRING",
									"map_name": "RAWID_STR",
									"map_type": "STRING",
									"name":     "RAWID_STR",
								},
							},
							"dirty_data_handle_mode": "ERROR",
						},
					},
					"advanced_config": "aps.flink.partition.value.for.empty.data=default",
					"source": []map[string]interface{}{
						{
							"type": "KAFKA",
							"kafka": []map[string]interface{}{
								{
									"start_offset_mode": "TIMESTAMP",
									"message_format":    "dataworks_json",
									"cloud_managed": []map[string]interface{}{
										{
											"enable_cross_account": "false",
										},
									},
									"start_offset_value": "1735689600000",
									"kafka_topic":        "${alicloud_alikafka_topic.create_kafka_topic_iceberg_internal_cdc_dataworks.topic}",
									"instance_type":      "CLOUD_MANAGED",
									"kafka_cluster_id":   "${ alicloud_alikafka_instance.create_kafka_instance_iceberg_internal_cdc_dataworks.id}",
									"cdc_config": []map[string]interface{}{
										{
											"sql_types":   "INSERT",
											"sync_tables": "GAMEAPP.GAME_EVENTS",
										},
									},
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":     "Kafka CDC Dataworks JSON to ADB Iceberg Internal Lake Pipeline Job",
						"db_cluster_id":   CHECKSET,
						"advanced_config": "aps.flink.partition.value.for.empty.data=default",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12579 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12579(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_kafka_iceberg_internal_cdc_dataworks" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "kafka-iceberg-internal-cdc-dataworks-vpc"
}

resource "alicloud_vswitch" "createVswitch_kafka_iceberg_internal_cdc_dataworks" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_kafka_iceberg_internal_cdc_dataworks.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.1.0/24"
  vswitch_name = "kafka-iceberg-internal-cdc-dataworks-vsw"
}

resource " alicloud_alikafka_instance" "create_kafka_instance_iceberg_internal_cdc_dataworks" {
  deploy_type     = "5"
  disk_type       = "0"
  vpc_id          = alicloud_vpc.createVpc_kafka_iceberg_internal_cdc_dataworks.id
  spec_type       = "normal"
  paid_type       = "1"
  vswitch_id      = alicloud_vswitch.createVswitch_kafka_iceberg_internal_cdc_dataworks.id
  io_max_spec     = "alikafka.hw.2xlarge"
  disk_size       = "500"
  service_version = "2.2.0"
  deploy_module   = "vpc"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_kafka_iceberg_internal_cdc_dataworks" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_kafka_iceberg_internal_cdc_dataworks.id
  vswitch_id                    = alicloud_vswitch.createVswitch_kafka_iceberg_internal_cdc_dataworks.id
  db_cluster_description        = "kafka_iceberg_internal_cdc_dataworks_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}

resource "alicloud_alikafka_topic" "create_kafka_topic_iceberg_internal_cdc_dataworks" {
  instance_id = alicloud_alikafka_instance.create_kafka_instance_iceberg_internal_cdc_dataworks.id
  config      = "{\"retention.ms\":\"36000000\"}"
  topic       = "kafka-iceberg-internal-cdc-dataworks-topic"
  remark      = "kafka-iceberg-internal-cdc-dataworks-topic-remark"
}

resource "alicloud_adb_lake_storage" "create_lake_storage_kafka_iceberg_internal_cdc_dataworks" {
  description   = "kafka_iceberg_internal_cdc_dataworks_lake_storage"
  db_cluster_id = alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_kafka_iceberg_internal_cdc_dataworks.id
}


`, name)
}

// Case resource_PipelineJob_kafka_iceberg_external 12580
func TestAccAliCloudAdbPipelineJob_basic12580(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12580)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12580)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "Kafka to ADB Iceberg External Lake (OSS) Pipeline Job",
					"sink": []map[string]interface{}{
						{
							"type": "OSS",
							"lake": []map[string]interface{}{
								{
									"table_name": "test_table_kafka_iceberg_external",
									"partition_specs": []map[string]interface{}{
										{
											"target_type_format": "yyyyMMdd",
											"source_column":      "__timestamp__",
											"strategy":           "ParseAsTimeAndFormat",
											"target_column":      "dt",
											"source_type_format": "APSLiteralTimestampMilliSecond",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"format_version":     "2",
										},
									},
									"db_name":        "test_db_kafka_iceberg_external",
									"storage_target": "oss://adb-test-bucket/kafka-iceberg-external/",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_kafka_iceberg_external.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "BIGINT",
									"map_name": "id",
									"map_type": "BIGINT",
									"name":     "id",
								},
								{
									"type":     "STRING",
									"map_name": "name",
									"map_type": "STRING",
									"name":     "name",
								},
								{
									"type":     "DOUBLE",
									"map_name": "value",
									"map_type": "DOUBLE",
									"name":     "value",
								},
								{
									"type":     "STRING",
									"map_name": "__key__",
									"map_type": "STRING",
									"name":     "__key__",
								},
								{
									"type":     "STRING",
									"map_name": "__topic__",
									"map_type": "STRING",
									"name":     "__topic__",
								},
								{
									"type":     "INT",
									"map_name": "__partition__",
									"map_type": "INT",
									"name":     "__partition__",
								},
								{
									"type":     "LONG",
									"map_name": "__offset__",
									"map_type": "LONG",
									"name":     "__offset__",
								},
								{
									"type":     "LONG",
									"map_name": "__timestamp__",
									"map_type": "TIMESTAMP",
									"name":     "__timestamp__",
								},
							},
							"dirty_data_handle_mode": "ERROR",
						},
					},
					"source": []map[string]interface{}{
						{
							"type": "KAFKA",
							"kafka": []map[string]interface{}{
								{
									"start_offset_mode": "earliest",
									"message_format":    "json",
									"cloud_managed": []map[string]interface{}{
										{
											"enable_cross_account": "false",
										},
									},
									"kafka_topic":      "${alicloud_alikafka_topic.create_kafka_topic_iceberg_external.topic}",
									"instance_type":    "CLOUD_MANAGED",
									"kafka_cluster_id": "${ alicloud_alikafka_instance.create_kafka_instance_iceberg_external.id}",
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":   "Kafka to ADB Iceberg External Lake (OSS) Pipeline Job",
						"db_cluster_id": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12580 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12580(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_kafka_iceberg_external" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "kafka-iceberg-external-vpc"
}

resource "alicloud_vswitch" "createVswitch_kafka_iceberg_external" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_kafka_iceberg_external.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.2.0/24"
  vswitch_name = "kafka-iceberg-external-vsw"
}

resource " alicloud_alikafka_instance" "create_kafka_instance_iceberg_external" {
  deploy_type     = "5"
  disk_type       = "0"
  vpc_id          = alicloud_vpc.createVpc_kafka_iceberg_external.id
  spec_type       = "normal"
  paid_type       = "1"
  vswitch_id      = alicloud_vswitch.createVswitch_kafka_iceberg_external.id
  io_max_spec     = "alikafka.hw.2xlarge"
  disk_size       = "500"
  service_version = "2.2.0"
  deploy_module   = "vpc"
}

resource "alicloud_alikafka_topic" "create_kafka_topic_iceberg_external" {
  instance_id = alicloud_alikafka_instance.create_kafka_instance_iceberg_external.id
  config      = "{\"retention.ms\":\"36000000\"}"
  topic       = "kafka-iceberg-external-topic"
  remark      = "kafka-iceberg-external-topic-remark"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_kafka_iceberg_external" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_kafka_iceberg_external.id
  vswitch_id                    = alicloud_vswitch.createVswitch_kafka_iceberg_external.id
  db_cluster_description        = "kafka_iceberg_external_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}


`, name)
}

// Case resource_PipelineJob_sls_iceberg_internal 12581
func TestAccAliCloudAdbPipelineJob_basic12581(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12581)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12581)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "SLS to ADB Iceberg Internal Lake Pipeline Job",
					"sink": []map[string]interface{}{
						{
							"type": "ADB_LAKE",
							"lake": []map[string]interface{}{
								{
									"table_name": "test_table_sls_iceberg_internal",
									"partition_specs": []map[string]interface{}{
										{
											"source_column": "project",
											"strategy":      "Identity",
											"target_column": "project",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"format_version":     "2",
										},
									},
									"db_name":        "test_db_sls_iceberg_internal",
									"storage_target": "${alicloud_adb_lake_storage.create_lake_storage_sls_iceberg_internal.lake_storage_id}",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_sls_iceberg_internal.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "STRING",
									"map_name": "unique_id",
									"map_type": "STRING",
									"name":     "unique_id",
								},
								{
									"type":     "STRING",
									"map_name": "consumer_group",
									"map_type": "STRING",
									"name":     "consumer_group",
								},
								{
									"type":     "STRING",
									"map_name": "project",
									"map_type": "STRING",
									"name":     "project",
								},
								{
									"type":     "STRING",
									"map_name": "fallbehind",
									"map_type": "STRING",
									"name":     "fallbehind",
								},
								{
									"type":     "STRING",
									"map_name": "shard",
									"map_type": "STRING",
									"name":     "shard",
								},
								{
									"type":     "STRING",
									"map_name": "logstore",
									"map_type": "STRING",
									"name":     "logstore",
								},
							},
							"dirty_data_handle_mode": "ERROR",
						},
					},
					"source": []map[string]interface{}{
						{
							"type": "SLS",
							"sls": []map[string]interface{}{
								{
									"start_offset_mode":  "TIMESTAMP",
									"project":            "${alicloud_log_project.create_sls_project_iceberg_internal.id}",
									"log_store":          "${alicloud_log_store.create_sls_logstore_iceberg_internal.name}",
									"start_offset_value": "1735689600000",
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":   "SLS to ADB Iceberg Internal Lake Pipeline Job",
						"db_cluster_id": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12581 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12581(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_sls_iceberg_internal" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "sls-iceberg-internal-vpc"
}

resource "alicloud_vswitch" "createVswitch_sls_iceberg_internal" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_sls_iceberg_internal.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.4.0/24"
  vswitch_name = "sls-iceberg-internal-vsw"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_sls_iceberg_internal" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_sls_iceberg_internal.id
  vswitch_id                    = alicloud_vswitch.createVswitch_sls_iceberg_internal.id
  db_cluster_description        = "sls_iceberg_internal_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}

resource "alicloud_log_project" "create_sls_project_iceberg_internal" {
  description = "SLS project for Iceberg internal lake test"
  name        = "sls-iceberg-internal-project"
}

resource "alicloud_adb_lake_storage" "create_lake_storage_sls_iceberg_internal" {
  description   = "sls_iceberg_internal_lake_storage"
  db_cluster_id = alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_sls_iceberg_internal.id
}

resource "alicloud_log_store" "create_sls_logstore_iceberg_internal" {
  retention_period = "200"
  shard_count      = "2"
  project          = alicloud_log_project.create_sls_project_iceberg_internal.id
  name             = "sls-iceberg-internal-logstore"
}


`, name)
}

// Case resource_PipelineJob_kafka_iceberg_internal_modify 12582
func TestAccAliCloudAdbPipelineJob_basic12582(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12582)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12582)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "Kafka to ADB Iceberg Internal Lake Pipeline Job modify",
					"sink": []map[string]interface{}{
						{
							"type": "ADB_LAKE",
							"lake": []map[string]interface{}{
								{
									"table_name": "test_table_kafka_iceberg_internal_modify",
									"partition_specs": []map[string]interface{}{
										{
											"target_type_format": "yyyyMMdd",
											"source_column":      "__timestamp__",
											"strategy":           "ParseAsTimeAndFormat",
											"target_column":      "dt",
											"source_type_format": "APSLiteralTimestampMilliSecond",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"format_version":     "2",
										},
									},
									"db_name":        "test_db_kafka_iceberg_internal_modify",
									"storage_target": "${alicloud_adb_lake_storage.create_lake_storage_kafka_iceberg_internal_modify.lake_storage_id}",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_kafka_iceberg_internal_modify.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "BIGINT",
									"map_name": "id",
									"map_type": "BIGINT",
									"name":     "id",
								},
								{
									"type":     "STRING",
									"map_name": "name",
									"map_type": "STRING",
									"name":     "name",
								},
								{
									"type":     "DOUBLE",
									"map_name": "value",
									"map_type": "DOUBLE",
									"name":     "value",
								},
								{
									"type":     "STRING",
									"map_name": "__key__",
									"map_type": "STRING",
									"name":     "__key__",
								},
								{
									"type":     "STRING",
									"map_name": "__topic__",
									"map_type": "STRING",
									"name":     "__topic__",
								},
								{
									"type":     "INT",
									"map_name": "__partition__",
									"map_type": "INT",
									"name":     "__partition__",
								},
								{
									"type":     "LONG",
									"map_name": "__offset__",
									"map_type": "LONG",
									"name":     "__offset__",
								},
								{
									"type":     "LONG",
									"map_name": "__timestamp__",
									"map_type": "TIMESTAMP",
									"name":     "__timestamp__",
								},
							},
							"dirty_data_handle_mode": "ERROR",
						},
					},
					"advanced_config": "aps.flink.target-file-size-bytes=134217728",
					"source": []map[string]interface{}{
						{
							"type": "KAFKA",
							"kafka": []map[string]interface{}{
								{
									"start_offset_mode": "earliest",
									"message_format":    "json",
									"cloud_managed": []map[string]interface{}{
										{
											"enable_cross_account": "false",
										},
									},
									"kafka_topic":      "${alicloud_alikafka_topic.create_kafka_topic_iceberg_internal_modify.topic}",
									"instance_type":    "CLOUD_MANAGED",
									"kafka_cluster_id": "${ alicloud_alikafka_instance.create_kafka_instance_iceberg_internal_modify.id}",
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":     "Kafka to ADB Iceberg Internal Lake Pipeline Job modify",
						"db_cluster_id":   CHECKSET,
						"advanced_config": "aps.flink.target-file-size-bytes=134217728",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sink": []map[string]interface{}{
						{
							"lake": []map[string]interface{}{
								{
									"partition_specs": []map[string]interface{}{
										{
											"target_type_format": "yyyyMMdd",
											"source_column":      "__timestamp__",
											"strategy":           "ParseAsTimeAndFormat",
											"target_column":      "dt",
											"source_type_format": "APSLiteralTimestampMilliSecond",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"format_version":     "2",
										},
									},
									"storage_target": "${alicloud_adb_lake_storage.create_lake_storage_kafka_iceberg_internal_modify.lake_storage_id}",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"source": []map[string]interface{}{
						{
							"kafka": []map[string]interface{}{
								{
									"start_offset_mode": "earliest",
									"message_format":    "json",
									"cloud_managed": []map[string]interface{}{
										{
											"enable_cross_account": "false",
										},
									},
									"kafka_topic":      "${alicloud_alikafka_topic.create_kafka_topic_iceberg_internal_modify.topic}",
									"kafka_cluster_id": "${ alicloud_alikafka_instance.create_kafka_instance_iceberg_internal_modify.id}",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12582 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12582(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_kafka_iceberg_internal_modify" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "kafka-iceberg-internal-modify-vpc"
}

resource "alicloud_vswitch" "createVswitch_kafka_iceberg_internal_modify" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_kafka_iceberg_internal_modify.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.11.0/24"
  vswitch_name = "kafka-iceberg-internal-modify-vsw"
}

resource " alicloud_alikafka_instance" "create_kafka_instance_iceberg_internal_modify" {
  deploy_type     = "5"
  disk_type       = "0"
  vpc_id          = alicloud_vpc.createVpc_kafka_iceberg_internal_modify.id
  spec_type       = "normal"
  paid_type       = "1"
  vswitch_id      = alicloud_vswitch.createVswitch_kafka_iceberg_internal_modify.id
  io_max_spec     = "alikafka.hw.2xlarge"
  disk_size       = "500"
  service_version = "2.2.0"
  deploy_module   = "vpc"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_kafka_iceberg_internal_modify" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_kafka_iceberg_internal_modify.id
  vswitch_id                    = alicloud_vswitch.createVswitch_kafka_iceberg_internal_modify.id
  db_cluster_description        = "kafka_iceberg_internal_modify_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}

resource "alicloud_alikafka_topic" "create_kafka_topic_iceberg_internal_modify" {
  instance_id = alicloud_alikafka_instance.create_kafka_instance_iceberg_internal_modify.id
  config      = "{\"retention.ms\":\"36000000\"}"
  topic       = "kafka-iceberg-internal-modify-topic"
  remark      = "kafka-iceberg-internal-modify-topic-remark"
}

resource "alicloud_adb_lake_storage" "create_lake_storage_kafka_iceberg_internal_modify" {
  description   = "kafka_iceberg_internal_modify_lake_storage"
  db_cluster_id = alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_kafka_iceberg_internal_modify.id
}


`, name)
}

// Case resource_PipelineJob_kafka_iceberg_internal_cdc_dataworks_modify 12583
func TestAccAliCloudAdbPipelineJob_basic12583(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12583)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12583)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "Kafka CDC Dataworks JSON to ADB Iceberg Internal Lake Pipeline Job",
					"sink": []map[string]interface{}{
						{
							"type": "ADB_LAKE",
							"lake": []map[string]interface{}{
								{
									"table_name": "test_table_kafka_iceberg_internal_cdc_dataworks_modify",
									"partition_specs": []map[string]interface{}{
										{
											"source_column":  "VARCHAR_TITLE",
											"strategy":       "Truncate",
											"target_column":  "VARCHAR_TITLE_TRUNCATE",
											"truncate_width": "4",
										},
										{
											"source_column": "NATIONAL_CHAR_NICK",
											"strategy":      "Bucket",
											"target_column": "NATIONAL_CHAR_NICK_BUCKET",
											"bucket_num":    "4",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"primary_key":        "UID",
											"format_version":     "2",
										},
									},
									"db_name":        "test_db_kafka_iceberg_internal_cdc_dataworks_modify",
									"storage_target": "${alicloud_adb_lake_storage.create_lake_storage_kafka_iceberg_internal_cdc_dataworks_modify.lake_storage_id}",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_kafka_iceberg_internal_cdc_dataworks_modify.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "INT",
									"map_name": "UID",
									"map_type": "INT",
									"name":     "UID",
								},
								{
									"type":     "DOUBLE",
									"map_name": "EVENT_ID",
									"map_type": "DOUBLE",
									"name":     "EVENT_ID",
								},
								{
									"type":     "DOUBLE",
									"map_name": "PLAYER_ID",
									"map_type": "DOUBLE",
									"name":     "PLAYER_ID",
								},
								{
									"type":     "DOUBLE",
									"map_name": "INTEGER_SCORE",
									"map_type": "DOUBLE",
									"name":     "INTEGER_SCORE",
								},
								{
									"type":     "DOUBLE",
									"map_name": "INT_RANK",
									"map_type": "DOUBLE",
									"name":     "INT_RANK",
								},
								{
									"type":     "DOUBLE",
									"map_name": "SMALL_LEVEL",
									"map_type": "DOUBLE",
									"name":     "SMALL_LEVEL",
								},
								{
									"type":     "DOUBLE",
									"map_name": "NUMBER_LIVES",
									"map_type": "DOUBLE",
									"name":     "NUMBER_LIVES",
								},
								{
									"type":     "DOUBLE",
									"map_name": "NUMERIC_PRICE",
									"map_type": "DOUBLE",
									"name":     "NUMERIC_PRICE",
								},
								{
									"type":     "DOUBLE",
									"map_name": "DECIMAL_DISCOUNT",
									"map_type": "DOUBLE",
									"name":     "DECIMAL_DISCOUNT",
								},
								{
									"type":     "DOUBLE",
									"map_name": "FLOAT_ACCURACY",
									"map_type": "DOUBLE",
									"name":     "FLOAT_ACCURACY",
								},
								{
									"type":     "DOUBLE",
									"map_name": "DOUBLE_POWER",
									"map_type": "DOUBLE",
									"name":     "DOUBLE_POWER",
								},
								{
									"type":     "DOUBLE",
									"map_name": "REAL_DROP",
									"map_type": "DOUBLE",
									"name":     "REAL_DROP",
								},
								{
									"type":     "STRING",
									"map_name": "CHAR_CLASS",
									"map_type": "STRING",
									"name":     "CHAR_CLASS",
								},
								{
									"type":     "STRING",
									"map_name": "NCHAR_REGION",
									"map_type": "STRING",
									"name":     "NCHAR_REGION",
								},
								{
									"type":     "STRING",
									"map_name": "VARCHAR_TITLE",
									"map_type": "STRING",
									"name":     "VARCHAR_TITLE",
								},
								{
									"type":     "STRING",
									"map_name": "VARCHAR2_ITEM",
									"map_type": "STRING",
									"name":     "VARCHAR2_ITEM",
								},
								{
									"type":     "STRING",
									"map_name": "NVARCHAR2_NPC",
									"map_type": "STRING",
									"name":     "NVARCHAR2_NPC",
								},
								{
									"type":     "STRING",
									"map_name": "CHARACTER_CODE",
									"map_type": "STRING",
									"name":     "CHARACTER_CODE",
								},
								{
									"type":     "STRING",
									"map_name": "CHAR_VARYING_TAG",
									"map_type": "STRING",
									"name":     "CHAR_VARYING_TAG",
								},
								{
									"type":     "STRING",
									"map_name": "CHARACTER_VARYING_ALIAS",
									"map_type": "STRING",
									"name":     "CHARACTER_VARYING_ALIAS",
								},
								{
									"type":     "STRING",
									"map_name": "NATIONAL_CHAR_NICK",
									"map_type": "STRING",
									"name":     "NATIONAL_CHAR_NICK",
								},
								{
									"type":     "STRING",
									"map_name": "NAT_CHAR_VARYING_COMMENT",
									"map_type": "STRING",
									"name":     "NAT_CHAR_VARYING_COMMENT",
								},
								{
									"type":     "STRING",
									"map_name": "NAT_CHARACTER_VARYING_NOTE",
									"map_type": "STRING",
									"name":     "NAT_CHARACTER_VARYING_NOTE",
								},
								{
									"type":     "STRING",
									"map_name": "NCHAR_VARYING_LABEL",
									"map_type": "STRING",
									"name":     "NCHAR_VARYING_LABEL",
								},
								{
									"type":     "STRING",
									"map_name": "CLOB_STORY",
									"map_type": "STRING",
									"name":     "CLOB_STORY",
								},
								{
									"type":     "STRING",
									"map_name": "NCLOB_DIALOG",
									"map_type": "STRING",
									"name":     "NCLOB_DIALOG",
								},
								{
									"type":     "TIMESTAMP",
									"map_name": "TIMESTAMP_EVENT",
									"map_type": "TIMESTAMP",
									"name":     "TIMESTAMP_EVENT",
								},
								{
									"type":     "TIMESTAMP",
									"map_name": "DATE_EVENT",
									"map_type": "TIMESTAMP",
									"name":     "DATE_EVENT",
								},
								{
									"type":     "DOUBLE",
									"map_name": "BIT_FLAG",
									"map_type": "DOUBLE",
									"name":     "BIT_FLAG",
								},
								{
									"type":     "DOUBLE",
									"map_name": "BOOL_FLAG",
									"map_type": "DOUBLE",
									"name":     "BOOL_FLAG",
								},
								{
									"type":     "STRING",
									"map_name": "RAWID_STR",
									"map_type": "STRING",
									"name":     "RAWID_STR",
								},
							},
							"dirty_data_handle_mode": "ERROR",
						},
					},
					"advanced_config": "aps.flink.partition.value.for.empty.data=default",
					"source": []map[string]interface{}{
						{
							"type": "KAFKA",
							"kafka": []map[string]interface{}{
								{
									"start_offset_mode": "TIMESTAMP",
									"message_format":    "dataworks_json",
									"cloud_managed": []map[string]interface{}{
										{
											"enable_cross_account": "false",
										},
									},
									"start_offset_value": "1735689600000",
									"kafka_topic":        "${alicloud_alikafka_topic.create_kafka_topic_iceberg_internal_cdc_dataworks_modify.topic}",
									"instance_type":      "CLOUD_MANAGED",
									"kafka_cluster_id":   "${ alicloud_alikafka_instance.create_kafka_instance_iceberg_internal_cdc_dataworks_modify.id}",
									"cdc_config": []map[string]interface{}{
										{
											"sql_types":   "INSERT",
											"sync_tables": "GAMEAPP.GAME_EVENTS",
										},
									},
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":     "Kafka CDC Dataworks JSON to ADB Iceberg Internal Lake Pipeline Job",
						"db_cluster_id":   CHECKSET,
						"advanced_config": "aps.flink.partition.value.for.empty.data=default",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sink": []map[string]interface{}{
						{
							"lake": []map[string]interface{}{
								{
									"partition_specs": []map[string]interface{}{
										{
											"source_column":  "VARCHAR_TITLE",
											"strategy":       "Truncate",
											"target_column":  "VARCHAR_TITLE_TRUNCATE",
											"truncate_width": "4",
										},
										{
											"source_column": "NATIONAL_CHAR_NICK",
											"strategy":      "Bucket",
											"target_column": "NATIONAL_CHAR_NICK_BUCKET",
											"bucket_num":    "4",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"primary_key":        "UID",
											"format_version":     "2",
										},
									},
									"storage_target": "${alicloud_adb_lake_storage.create_lake_storage_kafka_iceberg_internal_cdc_dataworks_modify.lake_storage_id}",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"source": []map[string]interface{}{
						{
							"kafka": []map[string]interface{}{
								{
									"start_offset_mode": "TIMESTAMP",
									"message_format":    "dataworks_json",
									"cloud_managed": []map[string]interface{}{
										{
											"enable_cross_account": "false",
										},
									},
									"start_offset_value": "1735689600000",
									"kafka_topic":        "${alicloud_alikafka_topic.create_kafka_topic_iceberg_internal_cdc_dataworks_modify.topic}",
									"kafka_cluster_id":   "${ alicloud_alikafka_instance.create_kafka_instance_iceberg_internal_cdc_dataworks_modify.id}",
									"cdc_config": []map[string]interface{}{
										{
											"sql_types":   "INSERT",
											"sync_tables": "GAMEAPP.GAME_EVENTS",
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12583 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12583(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_kafka_iceberg_internal_cdc_dataworks_modify" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "kafka-iceberg-internal-cdc-dataworks-modify-vpc"
}

resource "alicloud_vswitch" "createVswitch_kafka_iceberg_internal_cdc_dataworks_modify" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_kafka_iceberg_internal_cdc_dataworks_modify.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.17.0/24"
  vswitch_name = "kafka-iceberg-internal-cdc-dataworks-modify-vsw"
}

resource " alicloud_alikafka_instance" "create_kafka_instance_iceberg_internal_cdc_dataworks_modify" {
  deploy_type     = "5"
  disk_type       = "0"
  vpc_id          = alicloud_vpc.createVpc_kafka_iceberg_internal_cdc_dataworks_modify.id
  spec_type       = "normal"
  paid_type       = "1"
  vswitch_id      = alicloud_vswitch.createVswitch_kafka_iceberg_internal_cdc_dataworks_modify.id
  io_max_spec     = "alikafka.hw.2xlarge"
  disk_size       = "500"
  service_version = "2.2.0"
  deploy_module   = "vpc"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_kafka_iceberg_internal_cdc_dataworks_modify" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_kafka_iceberg_internal_cdc_dataworks_modify.id
  vswitch_id                    = alicloud_vswitch.createVswitch_kafka_iceberg_internal_cdc_dataworks_modify.id
  db_cluster_description        = "kafka_iceberg_internal_cdc_dataworks_modify_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}

resource "alicloud_alikafka_topic" "create_kafka_topic_iceberg_internal_cdc_dataworks_modify" {
  instance_id = alicloud_alikafka_instance.create_kafka_instance_iceberg_internal_cdc_dataworks_modify.id
  config      = "{\"retention.ms\":\"36000000\"}"
  topic       = "kafka-iceberg-internal-cdc-dataworks-modify-topic"
  remark      = "kafka-iceberg-internal-cdc-dataworks-modify-topic-remark"
}

resource "alicloud_adb_lake_storage" "create_lake_storage_kafka_iceberg_internal_cdc_dataworks_modify" {
  description   = "kafka_iceberg_internal_cdc_dataworks_modify_lake_storage"
  db_cluster_id = alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_kafka_iceberg_internal_cdc_dataworks_modify.id
}


`, name)
}

// Case resource_PipelineJob_sls_warehouse 12584
func TestAccAliCloudAdbPipelineJob_basic12584(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12584)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12584)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "SLS to ADB Warehouse Pipeline Job",
					"sink": []map[string]interface{}{
						{
							"type": "ADB_WAREHOUSE",
							"warehouse": []map[string]interface{}{
								{
									"table_name":                    "test_table_sls_warehouse",
									"enable_unix_timestamp_convert": "true",
									"user_name":                     "${alicloud_adb_lake_account.create_lake_account_sls_warehouse.account_name}",
									"table_creation_mode":           "USE_EXISTING",
									"enable_ip_white_list":          "true",
									"unix_timestamp_convert_format": "APSLiteralTimestampMilliSecond",
									"db_name":                       "test_db_sls_warehouse",
									"password":                      "${alicloud_adb_lake_account.create_lake_account_sls_warehouse.account_password}",
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_sls_warehouse.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "STRING",
									"map_name": "unique_id",
									"map_type": "STRING",
									"name":     "unique_id",
								},
								{
									"type":     "STRING",
									"map_name": "consumer_group",
									"map_type": "STRING",
									"name":     "consumer_group",
								},
								{
									"type":     "STRING",
									"map_name": "project",
									"map_type": "STRING",
									"name":     "project",
								},
								{
									"type":     "STRING",
									"map_name": "fallbehind",
									"map_type": "STRING",
									"name":     "fallbehind",
								},
								{
									"type":     "STRING",
									"map_name": "shard",
									"map_type": "STRING",
									"name":     "shard",
								},
								{
									"type":     "STRING",
									"map_name": "logstore",
									"map_type": "STRING",
									"name":     "logstore",
								},
							},
							"dirty_data_handle_mode": "TREAT_AS_NULL",
						},
					},
					"source": []map[string]interface{}{
						{
							"type": "SLS",
							"sls": []map[string]interface{}{
								{
									"start_offset_mode":  "TIMESTAMP",
									"project":            "${alicloud_log_project.create_sls_project_warehouse.id}",
									"log_store":          "${alicloud_log_store.create_sls_logstore_warehouse.name}",
									"start_offset_value": "1735689600000",
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":   "SLS to ADB Warehouse Pipeline Job",
						"db_cluster_id": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12584 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12584(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_sls_warehouse" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "sls-warehouse-vpc"
}

resource "alicloud_vswitch" "createVswitch_sls_warehouse" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_sls_warehouse.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.5.0/24"
  vswitch_name = "sls-warehouse-vsw"
}

resource "alicloud_log_project" "create_sls_project_warehouse" {
  description = "SLS project for ADB Warehouse test"
  name        = "sls-warehouse-project"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_sls_warehouse" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_sls_warehouse.id
  vswitch_id                    = alicloud_vswitch.createVswitch_sls_warehouse.id
  db_cluster_description        = "sls_warehouse_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}

resource "alicloud_log_store" "create_sls_logstore_warehouse" {
  retention_period = "200"
  shard_count      = "2"
  project          = alicloud_log_project.create_sls_project_warehouse.id
  name             = "sls-warehouse-logstore"
}

resource "alicloud_adb_lake_account" "create_lake_account_sls_warehouse" {
  account_description = "Account for SLS Warehouse pipeline test"
  db_cluster_id       = alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_sls_warehouse.id
  account_type        = "Super"
  engine              = "AnalyticDB"
  account_name        = "sls_warehouse_user"
  account_password    = "Test@12345678"
}


`, name)
}

// Case resource_PipelineJob_sls_paimon_external_modify 12585
func TestAccAliCloudAdbPipelineJob_basic12585(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12585)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12585)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "SLS to ADB Paimon External Lake (OSS) Pipeline Job",
					"sink": []map[string]interface{}{
						{
							"type": "OSS",
							"lake": []map[string]interface{}{
								{
									"table_name": "test_table_sls_paimon_external_modify",
									"partition_specs": []map[string]interface{}{
										{
											"source_column": "project",
											"strategy":      "Identity",
											"target_column": "project",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"db_name":             "test_db_sls_paimon_external_modify",
									"storage_target":      "oss://adb-test-bucket/sls-paimon-external-modify/",
									"table_format":        "APACHE_PAIMON",
									"paimon": []map[string]interface{}{
										{
											"table_type": "TABLE_WITHOUT_PK",
										},
									},
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_sls_paimon_external_modify.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "STRING",
									"map_name": "unique_id",
									"map_type": "STRING",
									"name":     "unique_id",
								},
								{
									"type":     "STRING",
									"map_name": "consumer_group",
									"map_type": "STRING",
									"name":     "consumer_group",
								},
								{
									"type":     "STRING",
									"map_name": "project",
									"map_type": "STRING",
									"name":     "project",
								},
								{
									"type":     "STRING",
									"map_name": "fallbehind",
									"map_type": "STRING",
									"name":     "fallbehind",
								},
								{
									"type":     "STRING",
									"map_name": "shard",
									"map_type": "STRING",
									"name":     "shard",
								},
								{
									"type":     "STRING",
									"map_name": "logstore",
									"map_type": "STRING",
									"name":     "logstore",
								},
							},
							"dirty_data_handle_mode": "ERROR",
						},
					},
					"source": []map[string]interface{}{
						{
							"type": "SLS",
							"sls": []map[string]interface{}{
								{
									"start_offset_mode":    "earliest",
									"project":              "${alicloud_log_project.create_sls_project_paimon_external_modify.id}",
									"enable_cross_account": "false",
									"log_store":            "${alicloud_log_store.create_sls_logstore_paimon_external_modify.name}",
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":   "SLS to ADB Paimon External Lake (OSS) Pipeline Job",
						"db_cluster_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sink": []map[string]interface{}{
						{
							"lake": []map[string]interface{}{
								{
									"partition_specs": []map[string]interface{}{
										{
											"source_column": "project",
											"strategy":      "Identity",
											"target_column": "project",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"storage_target":      "oss://adb-test-bucket/sls-paimon-external-modify/",
									"table_format":        "APACHE_PAIMON",
									"paimon": []map[string]interface{}{
										{
											"table_type": "TABLE_WITHOUT_PK",
										},
									},
								},
							},
						},
					},
					"source": []map[string]interface{}{
						{
							"sls": []map[string]interface{}{
								{
									"start_offset_mode":    "earliest",
									"project":              "${alicloud_log_project.create_sls_project_paimon_external_modify.id}",
									"enable_cross_account": "false",
									"log_store":            "${alicloud_log_store.create_sls_logstore_paimon_external_modify.name}",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12585 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12585(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_sls_paimon_external_modify" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "sls-paimon-external-modify-vpc"
}

resource "alicloud_log_project" "create_sls_project_paimon_external_modify" {
  description = "SLS project for Paimon external lake test"
  name        = "sls-paimon-external-modify-project"
}

resource "alicloud_vswitch" "createVswitch_sls_paimon_external_modify" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_sls_paimon_external_modify.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.16.0/24"
  vswitch_name = "sls-paimon-external-modify-vsw"
}

resource "alicloud_log_store" "create_sls_logstore_paimon_external_modify" {
  retention_period = "200"
  shard_count      = "2"
  project          = alicloud_log_project.create_sls_project_paimon_external_modify.id
  name             = "sls-paimon-external-modify-logstore"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_sls_paimon_external_modify" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_sls_paimon_external_modify.id
  vswitch_id                    = alicloud_vswitch.createVswitch_sls_paimon_external_modify.id
  db_cluster_description        = "sls_paimon_external_modify_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}


`, name)
}

// Case resource_PipelineJob_kafka_iceberg_external_modify 12586
func TestAccAliCloudAdbPipelineJob_basic12586(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_pipeline_job.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbPipelineJobMap12586)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbPipelineJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbPipelineJobBasicDependence12586)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "Kafka to ADB Iceberg External Lake (OSS) Pipeline Job Modify",
					"sink": []map[string]interface{}{
						{
							"type": "OSS",
							"lake": []map[string]interface{}{
								{
									"table_name": "test_table_kafka_iceberg_external_modify",
									"partition_specs": []map[string]interface{}{
										{
											"target_type_format": "yyyyMMdd",
											"source_column":      "__timestamp__",
											"strategy":           "ParseAsTimeAndFormat",
											"target_column":      "dt",
											"source_type_format": "APSLiteralTimestampMilliSecond",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"format_version":     "2",
										},
									},
									"db_name":        "test_db_kafka_iceberg_external_modify",
									"storage_target": "oss://adb-test-bucket/kafka-iceberg-external-modify/",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"db_cluster_id": "${alicloud_adb_db_cluster_lake_version.resource_DBClusterLakeVersion_kafka_iceberg_external_modify.id}",
					"transform": []map[string]interface{}{
						{
							"column_map": []map[string]interface{}{
								{
									"type":     "BIGINT",
									"map_name": "id",
									"map_type": "BIGINT",
									"name":     "id",
								},
								{
									"type":     "STRING",
									"map_name": "name",
									"map_type": "STRING",
									"name":     "name",
								},
								{
									"type":     "DOUBLE",
									"map_name": "value",
									"map_type": "DOUBLE",
									"name":     "value",
								},
								{
									"type":     "STRING",
									"map_name": "__key__",
									"map_type": "STRING",
									"name":     "__key__",
								},
								{
									"type":     "STRING",
									"map_name": "__topic__",
									"map_type": "STRING",
									"name":     "__topic__",
								},
								{
									"type":     "INT",
									"map_name": "__partition__",
									"map_type": "INT",
									"name":     "__partition__",
								},
								{
									"type":     "LONG",
									"map_name": "__offset__",
									"map_type": "LONG",
									"name":     "__offset__",
								},
								{
									"type":     "LONG",
									"map_name": "__timestamp__",
									"map_type": "TIMESTAMP",
									"name":     "__timestamp__",
								},
							},
							"dirty_data_handle_mode": "ERROR",
						},
					},
					"source": []map[string]interface{}{
						{
							"type": "KAFKA",
							"kafka": []map[string]interface{}{
								{
									"start_offset_mode": "earliest",
									"message_format":    "json",
									"cloud_managed": []map[string]interface{}{
										{
											"enable_cross_account": "false",
										},
									},
									"kafka_topic":      "${alicloud_alikafka_topic.create_kafka_topic_iceberg_external_modify.topic}",
									"instance_type":    "CLOUD_MANAGED",
									"kafka_cluster_id": "${ alicloud_alikafka_instance.create_kafka_instance_iceberg_external_modify.id}",
								},
							},
						},
					},
					"job_config": []map[string]interface{}{
						{
							"checkpoint_interval": "10000",
							"enable_serverless":   "true",
							"max_compute_unit":    "20",
							"min_compute_unit":    "2",
							"resource_group":      "serverless",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":   "Kafka to ADB Iceberg External Lake (OSS) Pipeline Job Modify",
						"db_cluster_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sink": []map[string]interface{}{
						{
							"lake": []map[string]interface{}{
								{
									"partition_specs": []map[string]interface{}{
										{
											"target_type_format": "yyyyMMdd",
											"source_column":      "__timestamp__",
											"strategy":           "ParseAsTimeAndFormat",
											"target_column":      "dt",
											"source_type_format": "APSLiteralTimestampMilliSecond",
										},
									},
									"table_creation_mode": "CREATE_NEW",
									"iceberg": []map[string]interface{}{
										{
											"write_distribution": "auto",
											"format_version":     "2",
										},
									},
									"storage_target": "oss://adb-test-bucket/kafka-iceberg-external-modify/",
									"table_format":   "APACHE_ICEBERG",
								},
							},
						},
					},
					"source": []map[string]interface{}{
						{
							"kafka": []map[string]interface{}{
								{
									"start_offset_mode": "earliest",
									"message_format":    "json",
									"cloud_managed": []map[string]interface{}{
										{
											"enable_cross_account": "false",
										},
									},
									"kafka_topic":      "${alicloud_alikafka_topic.create_kafka_topic_iceberg_external_modify.topic}",
									"kafka_cluster_id": "${ alicloud_alikafka_instance.create_kafka_instance_iceberg_external_modify.id}",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sink", "source"},
			},
		},
	})
}

var AlicloudAdbPipelineJobMap12586 = map[string]string{
	"create_time":     CHECKSET,
	"pipeline_job_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudAdbPipelineJobBasicDependence12586(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_kafka_iceberg_external_modify" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "kafka-iceberg-external-modify-vpc"
}

resource "alicloud_vswitch" "createVswitch_kafka_iceberg_external_modify" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_kafka_iceberg_external_modify.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.253.12.0/24"
  vswitch_name = "kafka-iceberg-external-modify-vsw"
}

resource " alicloud_alikafka_instance" "create_kafka_instance_iceberg_external_modify" {
  deploy_type     = "5"
  disk_type       = "0"
  vpc_id          = alicloud_vpc.createVpc_kafka_iceberg_external_modify.id
  spec_type       = "normal"
  paid_type       = "1"
  vswitch_id      = alicloud_vswitch.createVswitch_kafka_iceberg_external_modify.id
  io_max_spec     = "alikafka.hw.2xlarge"
  disk_size       = "500"
  service_version = "2.2.0"
  deploy_module   = "vpc"
}

resource "alicloud_alikafka_topic" "create_kafka_topic_iceberg_external_modify" {
  instance_id = alicloud_alikafka_instance.create_kafka_instance_iceberg_external_modify.id
  config      = "{\"retention.ms\":\"36000000\"}"
  topic       = "kafka-iceberg-external-modify-topic"
  remark      = "kafka-iceberg-external-modify-topic-remark"
}

resource "alicloud_adb_db_cluster_lake_version" "resource_DBClusterLakeVersion_kafka_iceberg_external_modify" {
  storage_resource              = "24ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.createVpc_kafka_iceberg_external_modify.id
  vswitch_id                    = alicloud_vswitch.createVswitch_kafka_iceberg_external_modify.id
  db_cluster_description        = "kafka_iceberg_external_modify_cluster"
  reserved_node_size            = "8ACU"
  product_form                  = "IntegrationForm"
  product_version               = "EnterpriseVersion"
  db_cluster_network_type       = "VPC"
  reserved_node_count           = "3"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
}


`, name)
}

// Test Adb PipelineJob. <<< Resource test cases, automatically generated.
