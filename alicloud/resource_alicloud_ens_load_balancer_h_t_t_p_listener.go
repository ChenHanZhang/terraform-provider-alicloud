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

func resourceAliCloudEnsLoadBalancerHTTPListener() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudEnsLoadBalancerHTTPListenerCreate,
		Read:   resourceAliCloudEnsLoadBalancerHTTPListenerRead,
		Update: resourceAliCloudEnsLoadBalancerHTTPListenerUpdate,
		Delete: resourceAliCloudEnsLoadBalancerHTTPListenerDelete,
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
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"health_check": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringInSlice([]string{"on", "off"}, false),
			},
			"health_check_connect_port": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntBetween(0, 65535),
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
				ValidateFunc: IntBetween(0, 50),
			},
			"health_check_method": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"head", "get"}, false),
			},
			"health_check_timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntBetween(0, 300),
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
			"idle_timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntBetween(0, 60),
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
			"request_timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntBetween(0, 180),
			},
			"scheduler": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"wrr", "rr"}, false),
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
			"x_forwarded_for": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"on", "off"}, false),
			},
		},
	}
}

func resourceAliCloudEnsLoadBalancerHTTPListenerCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateLoadBalancerHTTPListener"
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
	request["HealthCheck"] = d.Get("health_check")
	if v, ok := d.GetOk("health_check_domain"); ok {
		request["HealthCheckDomain"] = v
	}
	if v, ok := d.GetOk("health_check_uri"); ok {
		request["HealthCheckURI"] = v
	}
	if v, ok := d.GetOkExists("healthy_threshold"); ok && v.(int) > 0 {
		request["HealthyThreshold"] = v
	}
	if v, ok := d.GetOkExists("unhealthy_threshold"); ok && v.(int) > 0 {
		request["UnhealthyThreshold"] = v
	}
	if v, ok := d.GetOkExists("health_check_timeout"); ok && v.(int) > 0 {
		request["HealthCheckTimeout"] = v
	}
	if v, ok := d.GetOkExists("health_check_connect_port"); ok {
		request["HealthCheckConnectPort"] = v
	}
	if v, ok := d.GetOkExists("health_check_interval"); ok && v.(int) > 0 {
		request["HealthCheckInterval"] = v
	}
	if v, ok := d.GetOk("health_check_http_code"); ok {
		request["HealthCheckHttpCode"] = v
	}
	if v, ok := d.GetOkExists("idle_timeout"); ok && v.(int) > 0 {
		request["IdleTimeout"] = v
	}
	if v, ok := d.GetOkExists("request_timeout"); ok && v.(int) > 0 {
		request["RequestTimeout"] = v
	}
	if v, ok := d.GetOk("health_check_method"); ok {
		request["HealthCheckMethod"] = v
	}
	if v, ok := d.GetOk("x_forwarded_for"); ok {
		request["XForwardedFor"] = v
	}
	if v, ok := d.GetOkExists("backend_server_port"); ok {
		request["BackendServerPort"] = v
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ens_load_balancer_h_t_t_p_listener", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["LoadBalancerId"], request["ListenerPort"]))

	ensServiceV2 := EnsServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Stopped"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, ensServiceV2.EnsLoadBalancerHTTPListenerStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudEnsLoadBalancerHTTPListenerUpdate(d, meta)
}

func resourceAliCloudEnsLoadBalancerHTTPListenerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	ensServiceV2 := EnsServiceV2{client}

	objectRaw, err := ensServiceV2.DescribeEnsLoadBalancerHTTPListener(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ens_load_balancer_h_t_t_p_listener DescribeEnsLoadBalancerHTTPListener Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("backend_server_port", objectRaw["BackendServerPort"])
	d.Set("description", objectRaw["Description"])
	d.Set("health_check", objectRaw["HealthCheck"])
	d.Set("health_check_connect_port", objectRaw["HealthCheckConnectPort"])
	d.Set("health_check_domain", objectRaw["HealthCheckDomain"])
	d.Set("health_check_http_code", objectRaw["HealthCheckHttpCode"])
	d.Set("health_check_interval", objectRaw["HealthCheckInterval"])
	d.Set("health_check_method", objectRaw["HealthCheckMethod"])
	d.Set("health_check_timeout", objectRaw["HealthCheckTimeout"])
	d.Set("health_check_uri", objectRaw["HealthCheckURI"])
	d.Set("healthy_threshold", objectRaw["HealthyThreshold"])
	d.Set("idle_timeout", objectRaw["IdleTimeout"])
	d.Set("request_timeout", objectRaw["RequestTimeout"])
	d.Set("scheduler", objectRaw["Scheduler"])
	d.Set("status", objectRaw["Status"])
	d.Set("unhealthy_threshold", objectRaw["UnhealthyThreshold"])
	d.Set("x_forwarded_for", objectRaw["XForwardedFor"])
	d.Set("listener_port", objectRaw["ListenerPort"])

	parts := strings.Split(d.Id(), ":")
	d.Set("load_balancer_id", parts[0])

	return nil
}

func resourceAliCloudEnsLoadBalancerHTTPListenerUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	ensServiceV2 := EnsServiceV2{client}
	objectRaw, _ := ensServiceV2.DescribeEnsLoadBalancerHTTPListener(d.Id())

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

				request["ListenerProtocol"] = "http"
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
				stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, ensServiceV2.EnsLoadBalancerHTTPListenerStateRefreshFunc(d.Id(), "Status", []string{}))
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

				request["ListenerProtocol"] = "http"
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
				stateConf := BuildStateConf([]string{}, []string{"Stopped"}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, ensServiceV2.EnsLoadBalancerHTTPListenerStateRefreshFunc(d.Id(), "Status", []string{}))
				if _, err := stateConf.WaitForState(); err != nil {
					return WrapErrorf(err, IdMsg, d.Id())
				}

			}
		}
	}

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "SetLoadBalancerHTTPListenerAttribute"
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

	if !d.IsNewResource() && d.HasChange("health_check") {
		update = true
	}
	request["HealthCheck"] = d.Get("health_check")
	if !d.IsNewResource() && d.HasChange("health_check_domain") {
		update = true
		request["HealthCheckDomain"] = d.Get("health_check_domain")
	}

	if !d.IsNewResource() && d.HasChange("health_check_uri") {
		update = true
		request["HealthCheckURI"] = d.Get("health_check_uri")
	}

	if !d.IsNewResource() && d.HasChange("healthy_threshold") {
		update = true
		request["HealthyThreshold"] = d.Get("healthy_threshold")
	}

	if !d.IsNewResource() && d.HasChange("unhealthy_threshold") {
		update = true
		request["UnhealthyThreshold"] = d.Get("unhealthy_threshold")
	}

	if !d.IsNewResource() && d.HasChange("health_check_timeout") {
		update = true
		request["HealthCheckTimeout"] = d.Get("health_check_timeout")
	}

	if !d.IsNewResource() && d.HasChange("health_check_connect_port") {
		update = true
		request["HealthCheckConnectPort"] = d.Get("health_check_connect_port")
	}

	if !d.IsNewResource() && d.HasChange("health_check_interval") {
		update = true
		request["HealthCheckInterval"] = d.Get("health_check_interval")
	}

	if !d.IsNewResource() && d.HasChange("health_check_http_code") {
		update = true
		request["HealthCheckHttpCode"] = d.Get("health_check_http_code")
	}

	if !d.IsNewResource() && d.HasChange("idle_timeout") {
		update = true
		request["IdleTimeout"] = d.Get("idle_timeout")
	}

	if !d.IsNewResource() && d.HasChange("request_timeout") {
		update = true
		request["RequestTimeout"] = d.Get("request_timeout")
	}

	if !d.IsNewResource() && d.HasChange("health_check_method") {
		update = true
		request["HealthCheckMethod"] = d.Get("health_check_method")
	}

	if !d.IsNewResource() && d.HasChange("x_forwarded_for") {
		update = true
		request["XForwardedFor"] = d.Get("x_forwarded_for")
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

	return resourceAliCloudEnsLoadBalancerHTTPListenerRead(d, meta)
}

func resourceAliCloudEnsLoadBalancerHTTPListenerDelete(d *schema.ResourceData, meta interface{}) error {

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

	request["ListenerProtocol"] = "http"
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
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
