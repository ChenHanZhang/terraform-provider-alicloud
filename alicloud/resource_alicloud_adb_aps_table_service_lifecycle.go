// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudAdbApsTableServiceLifecycle() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAdbApsTableServiceLifecycleCreate,
		Read:   resourceAliCloudAdbApsTableServiceLifecycleRead,
		Update: resourceAliCloudAdbApsTableServiceLifecycleUpdate,
		Delete: resourceAliCloudAdbApsTableServiceLifecycleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"aps_job_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_cluster_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringMatch(regexp.MustCompile("amv-[A-Za-z0-9]{16}"), "ADB instance name"),
			},
			"operation_tables": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"table_names": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"database_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"process_all": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"on", "off"}, true),
			},
			"strategy_desc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"strategy_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"strategy_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"KEEP_BY_TIME"}, true),
			},
			"strategy_value": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringMatch(regexp.MustCompile("^([0-9]|[1-9][0-9]{1,3})$"), "Lifecycle Management Policy Values"),
			},
		},
	}
}

func resourceAliCloudAdbApsTableServiceLifecycleCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateApsLifecycleStrategy"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	conn, err := client.NewAdbClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["DBClusterId"] = d.Get("db_cluster_id")
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("operation_tables"); ok {
		operationTablesMaps := make([]map[string]interface{}, 0)
		for _, dataLoop := range v.([]interface{}) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["TableNames"] = dataLoopTmp["table_names"].(*schema.Set).List()
			dataLoopMap["ProcessAll"] = dataLoopTmp["process_all"]
			dataLoopMap["DatabaseName"] = dataLoopTmp["database_name"]
			operationTablesMaps = append(operationTablesMaps, dataLoopMap)
		}
		request["OperationTables"], _ = convertListMapToJsonString(operationTablesMaps)
	}

	if v, ok := d.GetOk("status"); ok {
		request["Status"] = v
	}
	request["StrategyValue"] = d.Get("strategy_value")
	request["StrategyName"] = d.Get("strategy_name")
	request["StrategyType"] = d.Get("strategy_type")
	if v, ok := d.GetOk("strategy_desc"); ok {
		request["StrategyDesc"] = v
	}
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2021-12-01"), StringPointer("AK"), query, request, &runtime)

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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_adb_aps_table_service_lifecycle", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["DBClusterId"], response["Data"]))

	return resourceAliCloudAdbApsTableServiceLifecycleRead(d, meta)
}

func resourceAliCloudAdbApsTableServiceLifecycleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	adbServiceV2 := AdbServiceV2{client}

	objectRaw, err := adbServiceV2.DescribeAdbApsTableServiceLifecycle(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_adb_aps_table_service_lifecycle DescribeAdbApsTableServiceLifecycle Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreatedTime"])
	d.Set("status", objectRaw["Status"])
	d.Set("strategy_desc", objectRaw["StrategyDesc"])
	d.Set("strategy_name", objectRaw["StrategyName"])
	d.Set("strategy_type", objectRaw["StrategyType"])
	d.Set("strategy_value", objectRaw["StrategyValue"])

	operationTables1Raw := objectRaw["OperationTables"]
	operationTablesMaps := make([]map[string]interface{}, 0)
	if operationTables1Raw != nil {
		for _, operationTablesChild1Raw := range operationTables1Raw.([]interface{}) {
			operationTablesMap := make(map[string]interface{})
			operationTablesChild1Raw := operationTablesChild1Raw.(map[string]interface{})
			operationTablesMap["database_name"] = operationTablesChild1Raw["DatabaseName"]
			operationTablesMap["process_all"] = operationTablesChild1Raw["ProcessAll"]

			tableNames1Raw := make([]interface{}, 0)
			if operationTablesChild1Raw["TableNames"] != nil {
				tableNames1Raw = operationTablesChild1Raw["TableNames"].([]interface{})
			}

			operationTablesMap["table_names"] = tableNames1Raw
			operationTablesMaps = append(operationTablesMaps, operationTablesMap)
		}
	}
	d.Set("operation_tables", operationTablesMaps)

	parts := strings.Split(d.Id(), ":")
	d.Set("db_cluster_id", parts[0])
	d.Set("aps_job_id", parts[1])

	return nil
}

func resourceAliCloudAdbApsTableServiceLifecycleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)
	parts := strings.Split(d.Id(), ":")
	action := "SwitchApsLifecycleStrategy"
	conn, err := client.NewAdbClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = parts[0]
	request["ApsJobId"] = parts[1]
	request["RegionId"] = client.RegionId
	if d.HasChange("status") {
		update = true
		request["Status"] = d.Get("status")
	}

	if update {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2021-12-01"), StringPointer("AK"), query, request, &runtime)

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
		d.SetPartial("status")
	}
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "UpdateApsLifecycleStrategy"
	conn, err = client.NewAdbClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = parts[0]
	request["ApsJobId"] = parts[1]
	request["RegionId"] = client.RegionId
	if d.HasChange("operation_tables") {
		update = true
		if v, ok := d.GetOk("operation_tables"); ok {
			operationTablesMaps := make([]map[string]interface{}, 0)
			for _, dataLoop := range v.([]interface{}) {
				dataLoopTmp := dataLoop.(map[string]interface{})
				dataLoopMap := make(map[string]interface{})
				dataLoopMap["TableNames"] = dataLoopTmp["table_names"].(*schema.Set).List()
				dataLoopMap["ProcessAll"] = dataLoopTmp["process_all"]
				dataLoopMap["DatabaseName"] = dataLoopTmp["database_name"]
				operationTablesMaps = append(operationTablesMaps, dataLoopMap)
			}
			request["OperationTables"], _ = convertListMapToJsonString(operationTablesMaps)
		}
	}

	if d.HasChange("status") {
		update = true
		request["Status"] = d.Get("status")
	}

	if d.HasChange("strategy_value") {
		update = true
	}
	request["StrategyValue"] = d.Get("strategy_value")
	if d.HasChange("strategy_type") {
		update = true
	}
	request["StrategyType"] = d.Get("strategy_type")
	if d.HasChange("strategy_name") {
		update = true
	}
	request["StrategyName"] = d.Get("strategy_name")
	if d.HasChange("strategy_desc") {
		update = true
		request["StrategyDesc"] = d.Get("strategy_desc")
	}

	if update {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2021-12-01"), StringPointer("AK"), query, request, &runtime)

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
		d.SetPartial("status")
		d.SetPartial("strategy_value")
		d.SetPartial("strategy_desc")
	}

	d.Partial(false)
	return resourceAliCloudAdbApsTableServiceLifecycleRead(d, meta)
}

func resourceAliCloudAdbApsTableServiceLifecycleDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteApsLifecycleStrategy"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	conn, err := client.NewAdbClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["DBClusterId"] = parts[0]
	request["ApsJobId"] = parts[1]
	request["RegionId"] = client.RegionId

	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2021-12-01"), StringPointer("AK"), query, request, &runtime)

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

	return nil
}
