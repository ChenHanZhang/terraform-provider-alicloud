// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudVpnGatewaySslVpnClientCert() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudVpnGatewaySslVpnClientCertCreate,
		Read:   resourceAliCloudVpnGatewaySslVpnClientCertRead,
		Update: resourceAliCloudVpnGatewaySslVpnClientCertUpdate,
		Delete: resourceAliCloudVpnGatewaySslVpnClientCertDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"ca_cert": {
				Type:      schema.TypeString,
				Sensitive: true,
			},
			"client_cert": {
				Type:      schema.TypeString,
				Sensitive: true,
			},
			"client_config": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_key": {
				Type:      schema.TypeString,
				Sensitive: true,
			},
			"create_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_vpn_client_cert_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_vpn_server_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudVpnGatewaySslVpnClientCertCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateSslVpnClientCert"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	request["SslVpnServerId"] = d.Get("ssl_vpn_server_id")
	if v, ok := d.GetOk("ssl_vpn_client_cert_name"); ok {
		request["Name"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ssl_vpn_client_cert", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["SslVpnClientCertId"]))

	return resourceAliCloudVpnGatewaySslVpnClientCertRead(d, meta)
}

func resourceAliCloudVpnGatewaySslVpnClientCertRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vPNGatewayServiceV2 := VPNGatewayServiceV2{client}

	objectRaw, err := vPNGatewayServiceV2.DescribeVpnGatewaySslVpnClientCert(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ssl_vpn_client_cert DescribeVpnGatewaySslVpnClientCert Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("ca_cert", objectRaw["CaCert"])
	d.Set("client_cert", objectRaw["ClientCert"])
	d.Set("client_config", objectRaw["ClientConfig"])
	d.Set("client_key", objectRaw["ClientKey"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("ssl_vpn_client_cert_name", objectRaw["Name"])
	d.Set("ssl_vpn_server_id", objectRaw["SslVpnServerId"])
	d.Set("status", objectRaw["Status"])

	return nil
}

func resourceAliCloudVpnGatewaySslVpnClientCertUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "ModifySslVpnClientCert"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["SslVpnClientCertId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("ssl_vpn_client_cert_name") {
		update = true
		request["Name"] = d.Get("ssl_vpn_client_cert_name")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
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

	return resourceAliCloudVpnGatewaySslVpnClientCertRead(d, meta)
}

func resourceAliCloudVpnGatewaySslVpnClientCertDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteSslVpnClientCert"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["SslVpnClientCertId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
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
