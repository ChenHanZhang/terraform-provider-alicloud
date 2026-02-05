// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudPolardbEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPolardbEndpointCreate,
		Read:   resourceAliCloudPolardbEndpointRead,
		Update: resourceAliCloudPolardbEndpointUpdate,
		Delete: resourceAliCloudPolardbEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(11 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"auto_add_new_nodes": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"Enable", "Disable"}, false),
			},
			"db_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_endpoint_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_endpoint_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"endpoint_config": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"endpoint_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"nodes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"read_write_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"ReadWrite", "ReadOnly"}, false),
			},
		},
	}
}

func resourceAliCloudPolardbEndpointCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateDBClusterEndpoint"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("db_cluster_id"); ok {
		request["DBClusterId"] = v
	}

	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("read_write_mode"); ok {
		request["ReadWriteMode"] = v
	}
	if v, ok := d.GetOk("auto_add_new_nodes"); ok {
		request["AutoAddNewNodes"] = v
	}
	if v, ok := d.GetOk("db_endpoint_description"); ok {
		request["DBEndpointDescription"] = v
	}
	if v, ok := d.GetOk("endpoint_config"); ok {
		request["EndpointConfig"] = v
	}
	if v, ok := d.GetOk("nodes"); ok {
		request["Nodes"] = v
	}
	request["EndpointType"] = d.Get("endpoint_type")
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_polardb_endpoint", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v", request["DBClusterId"]))

	polardbServiceV2 := PolardbServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Custom"}, d.Timeout(schema.TimeoutCreate), 60*time.Second, polardbServiceV2.DescribeAsyncPolardbEndpointStateRefreshFunc(d, response, "$.Items[*].EndpointType", []string{}))
	if jobDetail, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id(), jobDetail)
	}

	return resourceAliCloudPolardbEndpointRead(d, meta)
}

func resourceAliCloudPolardbEndpointRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	polardbServiceV2 := PolardbServiceV2{client}

	objectRaw, err := polardbServiceV2.DescribePolardbEndpoint(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_polardb_endpoint DescribePolardbEndpoint Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("auto_add_new_nodes", objectRaw["AutoAddNewNodes"])
	d.Set("db_endpoint_description", objectRaw["DBEndpointDescription"])
	d.Set("endpoint_config", objectRaw["EndpointConfig"])
	d.Set("endpoint_type", objectRaw["EndpointType"])
	d.Set("nodes", objectRaw["Nodes"])
	d.Set("read_write_mode", objectRaw["ReadWriteMode"])
	d.Set("db_cluster_id", objectRaw["DBClusterId"])
	d.Set("db_endpoint_id", objectRaw["DBEndpointId"])

	return nil
}

func resourceAliCloudPolardbEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "ModifyDBClusterEndpoint"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBEndpointId"] = parts[1]
	request["DBClusterId"] = parts[0]

	request["DBEndpointId"] = buildClientToken(action)
	if d.HasChange("read_write_mode") {
		update = true
		request["ReadWriteMode"] = d.Get("read_write_mode")
	}

	if d.HasChange("auto_add_new_nodes") {
		update = true
		request["AutoAddNewNodes"] = d.Get("auto_add_new_nodes")
	}

	if d.HasChange("db_endpoint_description") {
		update = true
		request["DBEndpointDescription"] = d.Get("db_endpoint_description")
	}

	if d.HasChange("endpoint_config") {
		update = true
		request["EndpointConfig"] = d.Get("endpoint_config")
	}

	if d.HasChange("nodes") {
		update = true
		request["Nodes"] = d.Get("nodes")
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

	return resourceAliCloudPolardbEndpointRead(d, meta)
}

func resourceAliCloudPolardbEndpointDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteDBClusterEndpoint"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["DBEndpointId"] = parts[1]
	request["DBClusterId"] = parts[0]

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
		if IsExpectedErrors(err, []string{"InvalidDBClusterId.NotFound"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
