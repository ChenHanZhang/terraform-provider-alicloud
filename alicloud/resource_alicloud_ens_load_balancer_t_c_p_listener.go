// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudEnsLoadBalancerTCPListener() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudEnsLoadBalancerTCPListenerCreate,
		Read:   resourceAliCloudEnsLoadBalancerTCPListenerRead,
		Update: resourceAliCloudEnsLoadBalancerTCPListenerUpdate,
		Delete: resourceAliCloudEnsLoadBalancerTCPListenerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"backend_server_port": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: IntBetween(0, 65535),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"eip_transmit": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"on", "off"}, false),
			},
			"established_timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntBetween(10, 900),
			},
			"health_check_connect_port": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntBetween(0, 65535),
			},
			"health_check_connect_timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntBetween(0, 300),
			},
			"health_check_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"health_check_http_code": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"http_2xx", "http_3xx", "http_4xx", "http_5xx"}, false),
			},
			"health_check_interval": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: IntBetween(0, 50),
			},
			"health_check_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"tcp", "http"}, false),
			},
			"health_check_uri": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"healthy_threshold": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntBetween(2, 10),
			},
			"listener_port": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"persistence_timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntBetween(0, 3600),
			},
			"scheduler": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"wrr", "wlc", "rr", "sch", "qch", "iqch"}, false),
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"Running", "Stopped", "Starting", "Configuring", "Stopping"}, false),
			},
			"unhealthy_threshold": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntBetween(2, 10),
			},
		},
	}
}

func resourceAliCloudEnsLoadBalancerTCPListenerCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateLoadBalancerTCPListener"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("load_balancer_id"); ok {
		request["LoadBalancerId"] = v
	}
	if v, ok := d.GetOk("listener_port"); ok {
		request["ListenerPort"] = v
	}

	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("scheduler"); ok {
		request["Scheduler"] = v
	}
	if v, ok := d.GetOkExists("persistence_timeout"); ok {
		request["PersistenceTimeout"] = v
	}
	if v, ok := d.GetOkExists("established_timeout"); ok && v.(int) > 0 {
		request["EstablishedTimeout"] = v
	}
	if v, ok := d.GetOkExists("healthy_threshold"); ok && v.(int) > 0 {
		request["HealthyThreshold"] = v
	}
	if v, ok := d.GetOkExists("unhealthy_threshold"); ok && v.(int) > 0 {
		request["UnhealthyThreshold"] = v
	}
	if v, ok := d.GetOkExists("health_check_connect_timeout"); ok && v.(int) > 0 {
		request["HealthCheckConnectTimeout"] = v
	}
	if v, ok := d.GetOkExists("health_check_interval"); ok && v.(int) > 0 {
		request["HealthCheckInterval"] = v
	}
	if v, ok := d.GetOk("health_check_domain"); ok {
		request["HealthCheckDomain"] = v
	}
	if v, ok := d.GetOk("health_check_http_code"); ok {
		request["HealthCheckHttpCode"] = v
	}
	if v, ok := d.GetOk("health_check_type"); ok {
		request["HealthCheckType"] = v
	}
	if v, ok := d.GetOkExists("backend_server_port"); ok {
		request["BackendServerPort"] = v
	}
	if v, ok := d.GetOkExists("health_check_connect_port"); ok && v.(int) > 0 {
		request["HealthCheckConnectPort"] = v
	}
	if v, ok := d.GetOk("eip_transmit"); ok {
		request["EipTransmit"] = v
	}
	if v, ok := d.GetOk("health_check_uri"); ok {
		request["HealthCheckURI"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Ens", "2017-11-10", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ens_load_balancer_t_c_p_listener", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["LoadBalancerId"], request["ListenerPort"]))

	ensServiceV2 := EnsServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Stopped"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, ensServiceV2.EnsLoadBalancerTCPListenerStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudEnsLoadBalancerTCPListenerUpdate(d, meta)
}

func resourceAliCloudEnsLoadBalancerTCPListenerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	ensServiceV2 := EnsServiceV2{client}

	objectRaw, err := ensServiceV2.DescribeEnsLoadBalancerTCPListener(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ens_load_balancer_t_c_p_listener DescribeEnsLoadBalancerTCPListener Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("backend_server_port", objectRaw["BackendServerPort"])
	d.Set("description", objectRaw["Description"])
	d.Set("eip_transmit", objectRaw["EipTransmit"])
	d.Set("established_timeout", objectRaw["EstablishedTimeout"])
	d.Set("health_check_connect_port", objectRaw["HealthCheckConnectPort"])
	d.Set("health_check_connect_timeout", objectRaw["HealthCheckConnectTimeout"])
	d.Set("health_check_domain", objectRaw["HealthCheckDomain"])
	d.Set("health_check_http_code", objectRaw["HealthCheckHttpCode"])
	d.Set("health_check_interval", objectRaw["HealthCheckInterval"])
	d.Set("health_check_type", objectRaw["HealthCheckType"])
	d.Set("health_check_uri", objectRaw["HealthCheckURI"])
	d.Set("healthy_threshold", objectRaw["HealthyThreshold"])
	d.Set("persistence_timeout", objectRaw["PersistenceTimeout"])
	d.Set("scheduler", objectRaw["Scheduler"])
	d.Set("status", objectRaw["Status"])
	d.Set("unhealthy_threshold", objectRaw["UnhealthyThreshold"])
	d.Set("listener_port", objectRaw["ListenerPort"])

	parts := strings.Split(d.Id(), ":")
	d.Set("load_balancer_id", parts[0])

	return nil
}

func resourceAliCloudEnsLoadBalancerTCPListenerUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	ensServiceV2 := EnsServiceV2{client}
	objectRaw, _ := ensServiceV2.DescribeEnsLoadBalancerTCPListener(d.Id())

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
				action := "StartLoadBalancerListener"
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				request["ListenerPort"] = parts[1]
				request["LoadBalancerId"] = parts[0]

				request["ListenerProtocol"] = "tcp"
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.RpcPost("Ens", "2017-11-10", action, query, request, true)
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
				ensServiceV2 := EnsServiceV2{client}
				stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, ensServiceV2.EnsLoadBalancerTCPListenerStateRefreshFunc(d.Id(), "Status", []string{}))
				if _, err := stateConf.WaitForState(); err != nil {
					return WrapErrorf(err, IdMsg, d.Id())
				}

			}
			if target == "Stopped" {
				parts := strings.Split(d.Id(), ":")
				action := "StopLoadBalancerListener"
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				request["ListenerPort"] = parts[1]
				request["LoadBalancerId"] = parts[0]

				request["ListenerProtocol"] = "tcp"
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.RpcPost("Ens", "2017-11-10", action, query, request, true)
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
				ensServiceV2 := EnsServiceV2{client}
				stateConf := BuildStateConf([]string{}, []string{"Stopped"}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, ensServiceV2.EnsLoadBalancerTCPListenerStateRefreshFunc(d.Id(), "Status", []string{}))
				if _, err := stateConf.WaitForState(); err != nil {
					return WrapErrorf(err, IdMsg, d.Id())
				}

			}
		}
	}

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "SetLoadBalancerTCPListenerAttribute"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["LoadBalancerId"] = parts[0]
	request["ListenerPort"] = parts[1]

	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	if !d.IsNewResource() && d.HasChange("scheduler") {
		update = true
		request["Scheduler"] = d.Get("scheduler")
	}

	if !d.IsNewResource() && d.HasChange("persistence_timeout") {
		update = true
		request["PersistenceTimeout"] = d.Get("persistence_timeout")
	}

	if !d.IsNewResource() && d.HasChange("established_timeout") {
		update = true
		request["EstablishedTimeout"] = d.Get("established_timeout")
	}

	if !d.IsNewResource() && d.HasChange("healthy_threshold") {
		update = true
		request["HealthyThreshold"] = d.Get("healthy_threshold")
	}

	if !d.IsNewResource() && d.HasChange("unhealthy_threshold") {
		update = true
		request["UnhealthyThreshold"] = d.Get("unhealthy_threshold")
	}

	if !d.IsNewResource() && d.HasChange("health_check_connect_timeout") {
		update = true
		request["HealthCheckConnectTimeout"] = d.Get("health_check_connect_timeout")
	}

	if !d.IsNewResource() && d.HasChange("health_check_connect_port") {
		update = true
		request["HealthCheckConnectPort"] = d.Get("health_check_connect_port")
	}

	if !d.IsNewResource() && d.HasChange("health_check_interval") {
		update = true
		request["HealthCheckInterval"] = d.Get("health_check_interval")
	}

	if !d.IsNewResource() && d.HasChange("health_check_domain") {
		update = true
		request["HealthCheckDomain"] = d.Get("health_check_domain")
	}

	if !d.IsNewResource() && d.HasChange("health_check_http_code") {
		update = true
		request["HealthCheckHttpCode"] = d.Get("health_check_http_code")
	}

	if !d.IsNewResource() && d.HasChange("health_check_type") {
		update = true
		request["HealthCheckType"] = d.Get("health_check_type")
	}

	if !d.IsNewResource() && d.HasChange("eip_transmit") {
		update = true
		request["EipTransmit"] = d.Get("eip_transmit")
	}

	if !d.IsNewResource() && d.HasChange("health_check_uri") {
		update = true
		request["HealthCheckURI"] = d.Get("health_check_uri")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ens", "2017-11-10", action, query, request, true)
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

	return resourceAliCloudEnsLoadBalancerTCPListenerRead(d, meta)
}

func resourceAliCloudEnsLoadBalancerTCPListenerDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteLoadBalancerListener"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["ListenerPort"] = parts[1]
	request["LoadBalancerId"] = parts[0]

	request["ListenerProtocol"] = "tcp"
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Ens", "2017-11-10", action, query, request, true)
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
		if IsExpectedErrors(err, []string{"ListenerNotFound"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	ensServiceV2 := EnsServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 5*time.Second, ensServiceV2.EnsLoadBalancerTCPListenerStateRefreshFunc(d.Id(), "ListenerPort", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
