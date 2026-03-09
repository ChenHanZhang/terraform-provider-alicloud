// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
  "encoding/json"
  "fmt"
  "log"
  "time"
  "strings"

  util "github.com/alibabacloud-go/tea-utils/service"

    "github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAliCloudEcsNetworkInterface() *schema.Resource {
  return &schema.Resource{
            Create: resourceAliCloudEcsNetworkInterfaceCreate,
                Read: resourceAliCloudEcsNetworkInterfaceRead,
                Update: resourceAliCloudEcsNetworkInterfaceUpdate,
            Delete: resourceAliCloudEcsNetworkInterfaceDelete,
    Importer: &schema.ResourceImporter{
      State: schema.ImportStatePassthrough,
    },
    Timeouts: &schema.ResourceTimeout{
              Create: schema.DefaultTimeout(5 * time.Minute),
                  Update: schema.DefaultTimeout(5 * time.Minute),
                  Delete: schema.DefaultTimeout(5 * time.Minute),
        },
            Schema: map[string]*schema.Schema {
                                "create_time": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "delete_on_release": {
            Type:schema.TypeBool,
                                                                                            Optional: true,
                                                    Computed: true,
                                                                                                                                                                                        },
                                    "description": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                        "ipv4_prefix": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                                        ForceNew: true,
                                                                                                                                                                                                    Elem:&schema.Schema{Type: schema.TypeString},
                                                                                        },
                                    "ipv6_prefix": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                                        ForceNew: true,
                                                                                                                                                                                                    Elem:&schema.Schema{Type: schema.TypeString},
                                                                                        },
                                    "ipv6_sets": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                                                                                                                                                                                                                                                                Elem: &schema.Resource{
                                  Schema: map[string]*schema.Schema {
                        "ipv6_address": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
            },
                    },
                                        },
                                            "network_interface_name": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                    Computed: true,
                                                                                                                                                                                        },
                                            "primary_ip_address": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                    Computed: true,
                                                        ForceNew: true,
                                                                                                                                                            },
                                    "private_ip_sets": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                                                                                                                                                                                                                                                                Elem: &schema.Resource{
                                  Schema: map[string]*schema.Schema {
                        "private_ip_address": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "associated_public_ip": {
            Type:schema.TypeList,
                                                                                                            Computed: true,
                                                                                                                                                                                            MaxItems: 1,
                Elem: &schema.Resource {
                                    Schema: map[string]*schema.Schema {
                        "public_ip_address": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "allocation_id": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
            },
                },
                        },
                                    "primary": {
            Type:schema.TypeBool,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
            },
                    },
                                        },
                                    "qo_s_config": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                                                                                                                                                                            MaxItems: 1,
                Elem: &schema.Resource {
                              Schema: map[string]*schema.Schema {
                        "qo_s": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                                                                                                                                                                            MaxItems: 1,
                Elem: &schema.Resource {
                                    Schema: map[string]*schema.Schema {
                        "pps_rx": {
            Type:schema.TypeInt,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "bandwidth_tx": {
            Type:schema.TypeInt,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "bandwidth_rx": {
            Type:schema.TypeInt,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "pps_tx": {
            Type:schema.TypeInt,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "concurrent_connections": {
            Type:schema.TypeInt,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
            },
                },
                        },
                                    "enable_qo_s": {
            Type:schema.TypeBool,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
            },
                },
                        },
                                    "queue_number": {
            Type:schema.TypeInt,
                                                                                            Optional: true,
                                                    Computed: true,
                                                                                                                                                                                        },
                                        "region_id": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "resource_group_id": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                    Computed: true,
                                                        ForceNew: true,
                                                                                                                                                            },
                                    "resource_type": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "security_group_ids": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                    Computed: true,
                                                                                                                                                                                                                                Elem:&schema.Schema{Type: schema.TypeString},
                                                                                        },
                                            "status": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                                                "tags": tagsSchema(),
                                                    "vswitch_id": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                        ForceNew: true,
                                                                                                                                                            },
                    },
  }
}

func resourceAliCloudEcsNetworkInterfaceCreate(d *schema.ResourceData, meta interface{}) error {

        
        
                                                                                    client := meta.(*connectivity.AliyunClient)
             
        action := "CreateNetworkInterface"
            var request map[string]interface{}
        var response map[string]interface{}
                            query := make(map[string]interface{})
                    var err error
        request = make(map[string]interface{})
                request["RegionId"] = client.RegionId
                        request["ClientToken"] = buildClientToken(action)
            
                                                                if v, ok := d.GetOkExists("delete_on_release");  ok    {
                                                            request["DeleteOnRelease"] = v                                        }
                                                                                if v, ok := d.GetOk("resource_group_id");  ok    {
                                                            request["ResourceGroupId"] = v                                        }
                                                                                	if v, ok := d.GetOk("tags");  ok  {
			tagsMap := ConvertTags(v.(map[string]interface{}))
							request = expandTagsToMap(request, tagsMap)
			}

                                                                        if v, ok := d.GetOk("description");  ok    {
                                                            request["Description"] = v                                        }
                                                                                                request["VSwitchId"] = d.Get("vswitch_id")
                                                                                if v, ok := d.GetOk("network_interface_name");  ok    {
                                                            request["NetworkInterfaceName"] = v                                        }
                                                                                	if v, ok := d.GetOk("ipv6_sets");  ok  {
			localData1, _ := jsonpath.Get("$[*].ipv6_address", v)
								localData1Array := convertToInterfaceArray(localData1)

		request["Ipv6Address"] = localData1Array
}

                                                                        	if v, ok := d.GetOk("security_group_ids");  ok  {
									securityGroupIdsMapsArray := convertToInterfaceArray(v)

		request["SecurityGroupIds"] = securityGroupIdsMapsArray
}

                                                                        if v, ok := d.GetOkExists("queue_number");  ok    {
                                                            request["QueueNumber"] = v                                        }
                                                                                	if v, ok := d.GetOk("private_ip_sets");  ok  {
			localData3, _ := jsonpath.Get("$[*].private_ip_address", v)
								localData3Array := convertToInterfaceArray(localData3)

		request["PrivateIpAddress"] = localData3Array
}

                                                                        if v, ok := d.GetOk("primary_ip_address");  ok    {
                                                            request["PrimaryIpAddress"] = v                                        }
                                                    wait := incrementalWait(3*time.Second, 5*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
                        response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
                    if err != nil {
                        if  NeedRetry(err) {
                wait()
                return resource.RetryableError(err)
            }
                        return resource.NonRetryableError(err)
        }
        return nil
    })
    addDebug(action, response, request)

    if err != nil {
        return WrapErrorf(err, DefaultErrorMsg, "alicloud_ecs_network_interface", action, AlibabaCloudSdkGoERROR)
    }
        
                        d.SetId(fmt.Sprint(response["NetworkInterfaceId"]))        

                ecsServiceV2 := EcsServiceV2{client}
        stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutCreate),  5 * time.Second ,ecsServiceV2.EcsNetworkInterfaceStateRefreshFunc(d.Id(), "Status", []string {}))
        if _, err := stateConf.WaitForState(); err != nil {
            return WrapErrorf(err, IdMsg, d.Id())
        }
        
        
        
        return resourceAliCloudEcsNetworkInterfaceUpdate(d, meta)
    }


    		func resourceAliCloudEcsNetworkInterfaceRead(d *schema.ResourceData, meta interface{}) error {
	    		client := meta.(*connectivity.AliyunClient)
		ecsServiceV2 := EcsServiceV2{client}

												objectRaw, err := ecsServiceV2.DescribeEcsNetworkInterface(d.Id())
				if err != nil  {
					if !d.IsNewResource() && NotFoundError(err) {
				log.Printf("[DEBUG] Resource alicloud_ecs_network_interface DescribeEcsNetworkInterface Failed!!! %s", err)
				d.SetId("")
				return nil
			}
					return WrapError(err)
		}

				

                            
                                        d.Set("create_time", objectRaw["CreationTime"])
                                            d.Set("delete_on_release", objectRaw["DeleteOnRelease"])
                                            d.Set("description", objectRaw["Description"])
                                            d.Set("network_interface_name", objectRaw["NetworkInterfaceName"])
                                            d.Set("primary_ip_address", objectRaw["PrivateIpAddress"])
                                            d.Set("queue_number", objectRaw["QueueNumber"])
                                            d.Set("region_id", objectRaw["RegionId"])
                                            d.Set("resource_group_id", objectRaw["ResourceGroupId"])
                                            d.Set("status", objectRaw["Status"])
                                            d.Set("vswitch_id", objectRaw["VSwitchId"])
                
                                                                                                                                                                                                                                                                                                                        ipv4PrefixSetChildRaw := make([]interface{}, 0)
if ipv4PrefixSetRaw["Ipv4PrefixSetChild"] != nil {
ipv4PrefixSetChildRaw = convertToInterfaceArray(ipv4PrefixSetRaw["Ipv4PrefixSetChild"])}

                                    d.Set("ipv4_prefix", ipv4PrefixSetChildRaw)
                                                                                                                                                                        ipv6PrefixSetChildRaw := make([]interface{}, 0)
if ipv6PrefixSetRaw["Ipv6PrefixSetChild"] != nil {
ipv6PrefixSetChildRaw = convertToInterfaceArray(ipv6PrefixSetRaw["Ipv6PrefixSetChild"])}

                                    d.Set("ipv6_prefix", ipv6PrefixSetChildRaw)
                                                                                                                                                            ipv6SetRaw, _ := jsonpath.Get("$.Ipv6Sets.Ipv6Set", objectRaw)
                                    ipv6SetsMaps := make([]map[string]interface{}, 0)
                                                if ipv6SetRaw != nil {
for _, ipv6SetChildRaw := range convertToInterfaceArray(ipv6SetRaw) {
                                ipv6SetsMap := make(map[string]interface{})
                                                        ipv6SetChildRaw := ipv6SetChildRaw.(map[string]interface{})
                                        ipv6SetsMap["ipv6_address"] = ipv6SetChildRaw["Ipv6Address"]
                
                                                                                    ipv6SetsMaps = append(ipv6SetsMaps, ipv6SetsMap)
            }
                        }
                                                if err := d.Set("ipv6_sets", ipv6SetsMaps); err != nil {
return err
}
                                                                                                                                                                                                                                                                    privateIpSetRaw, _ := jsonpath.Get("$.PrivateIpSets.PrivateIpSet", objectRaw)
                                    privateIpSetsMaps := make([]map[string]interface{}, 0)
                                                if privateIpSetRaw != nil {
for _, privateIpSetChildRaw := range convertToInterfaceArray(privateIpSetRaw) {
                                privateIpSetsMap := make(map[string]interface{})
                                                        privateIpSetChildRaw := privateIpSetChildRaw.(map[string]interface{})
                                        privateIpSetsMap["primary"] = privateIpSetChildRaw["Primary"]
                                            privateIpSetsMap["private_ip_address"] = privateIpSetChildRaw["PrivateIpAddress"]
                
                                                                    associatedPublicIpMaps := make([]map[string]interface{}, 0)
        associatedPublicIpMap := make(map[string]interface{})
                associatedPublicIpRawObj, _ := jsonpath.Get("$.AssociatedPublicIp", privateIpSetChildRaw)
associatedPublicIpRaw := make(map[string]interface{})
if associatedPublicIpRawObj != nil {
associatedPublicIpRaw = associatedPublicIpRawObj.(map[string]interface{})
}
if len(associatedPublicIpRaw) > 0 {
                                        associatedPublicIpMap["allocation_id"] = associatedPublicIpRaw["AllocationId"]
                                            associatedPublicIpMap["public_ip_address"] = associatedPublicIpRaw["PublicIpAddress"]
                
                                                                                                        associatedPublicIpMaps = append(associatedPublicIpMaps, associatedPublicIpMap)
}
privateIpSetsMap["associated_public_ip"] = associatedPublicIpMaps
                                                                                                                                                                    privateIpSetsMaps = append(privateIpSetsMaps, privateIpSetsMap)
            }
                        }
                                                if err := d.Set("private_ip_sets", privateIpSetsMaps); err != nil {
return err
}
                                                                                                                                                                                                                                                                                                                                    securityGroupIdRaw, _ := jsonpath.Get("$.SecurityGroupIds.SecurityGroupId", objectRaw)
                                    d.Set("security_group_ids", securityGroupIdRaw)
                                                                                                                                                            tagsMaps, _ := jsonpath.Get("$.Tags.Tag", objectRaw)
                d.Set("tags", tagsToMap(tagsMaps))
                                                                    
																						objectRaw, err = ecsServiceV2.DescribeNetworkInterfaceDescribeNetworkInterfaceAttribute(d.Id())
							if err != nil  && !NotFoundError(err)  {
					return WrapError(err)
		}

				

                            
                                        d.Set("delete_on_release", objectRaw["DeleteOnRelease"])
                                            d.Set("description", objectRaw["Description"])
                                            d.Set("network_interface_name", objectRaw["NetworkInterfaceName"])
                                            d.Set("queue_number", objectRaw["QueueNumber"])
                                            d.Set("resource_group_id", objectRaw["ResourceGroupId"])
                                            d.Set("status", objectRaw["Status"])
                                            d.Set("vswitch_id", objectRaw["VSwitchId"])
                
                                                                                                                                                                                                                                                        ipv6SetRaw, _ := jsonpath.Get("$.Ipv6Sets.Ipv6Set", objectRaw)
                                    ipv6SetsMaps := make([]map[string]interface{}, 0)
                                                if ipv6SetRaw != nil {
for _, ipv6SetChildRaw := range convertToInterfaceArray(ipv6SetRaw) {
                                ipv6SetsMap := make(map[string]interface{})
                                                        ipv6SetChildRaw := ipv6SetChildRaw.(map[string]interface{})
                                        ipv6SetsMap["ipv6_address"] = ipv6SetChildRaw["Ipv6Address"]
                
                                                                                    ipv6SetsMaps = append(ipv6SetsMaps, ipv6SetsMap)
            }
                        }
                                                if err := d.Set("ipv6_sets", ipv6SetsMaps); err != nil {
return err
}
                                                                                                                                    qoSConfigMaps := make([]map[string]interface{}, 0)
        qoSConfigMap := make(map[string]interface{})
                qoSConfigRaw := make(map[string]interface{})
if objectRaw["QoSConfig"] != nil {
qoSConfigRaw = objectRaw["QoSConfig"].(map[string]interface{})}
if len(qoSConfigRaw) > 0 {
                                        qoSConfigMap["enable_qo_s"] = qoSConfigRaw["EnableQoS"]
                
                                                                                                                        qoSMaps := make([]map[string]interface{}, 0)
        qoSMap := make(map[string]interface{})
                qoSRaw := make(map[string]interface{})
if qoSConfigRaw["QoS"] != nil {
qoSRaw = qoSConfigRaw["QoS"].(map[string]interface{})}
if len(qoSRaw) > 0 {
                                        qoSMap["bandwidth_rx"] = qoSRaw["BandwidthRx"]
                                            qoSMap["bandwidth_tx"] = qoSRaw["BandwidthTx"]
                                            qoSMap["concurrent_connections"] = qoSRaw["ConcurrentConnections"]
                                            qoSMap["pps_rx"] = qoSRaw["PpsRx"]
                                            qoSMap["pps_tx"] = qoSRaw["PpsTx"]
                
                                                                                                                                                                                                                                                                    qoSMaps = append(qoSMaps, qoSMap)
}
qoSConfigMap["qo_s"] = qoSMaps
                            qoSConfigMaps = append(qoSConfigMaps, qoSConfigMap)
}
if err := d.Set("qo_s_config", qoSConfigMaps); err != nil {
return err
}
                                                                                                                                                                                                                                                                                    tagsMaps, _ := jsonpath.Get("$.Tags.Tag", objectRaw)
                d.Set("tags", tagsToMap(tagsMaps))
                                                                    
							
        
        
		        	return nil
}


    func resourceAliCloudEcsNetworkInterfaceUpdate(d *schema.ResourceData, meta interface{}) error {
                client := meta.(*connectivity.AliyunClient)
        var request map[string]interface{}
    var response map[string]interface{}
            var query map[string]interface{}
                        update := false
                        d.Partial(true)
    
                ecsServiceV2 := EcsServiceV2{client}
        objectRaw, _ := ecsServiceV2.DescribeEcsNetworkInterface(d.Id())
    
        
                    initedQoSConfig.0.EnableQoS := false
if _, ok := d.GetOkExists("qo_s_config.0.enable_qo_s"); ok && d.IsNewResource() {
	initedQoSConfig.0.EnableQoS = true
}
if initedQoSConfig.0.EnableQoS || d.HasChange("qo_s_config.0.enable_qo_s") {
    var err error
    target := d.Get("qo_s_config.0.enable_qo_s").(bool)
                                        if  target == true {
                                                            action := "EnableNetworkInterfaceQoS"
    request = make(map[string]interface{})
                            query = make(map[string]interface{})
                            request["NetworkInterfaceId"] = d.Id()
            request["RegionId"] = client.RegionId    
                                                                                            if v, ok := d.GetOkExists("qo_s_config");  ok    {
                            qoSConfigQoSBandwidthRxJsonPath, err := jsonpath.Get("$[0].qo_s[0].bandwidth_rx", v)
            if err == nil && qoSConfigQoSBandwidthRxJsonPath != "" {
                                                    request["QoS.BandwidthRx"] = qoSConfigQoSBandwidthRxJsonPath                            }
                                }
                                                                                                                                            if v, ok := d.GetOkExists("qo_s_config");  ok    {
                            qoSConfigQoSBandwidthTxJsonPath, err := jsonpath.Get("$[0].qo_s[0].bandwidth_tx", v)
            if err == nil && qoSConfigQoSBandwidthTxJsonPath != "" {
                                                    request["QoS.BandwidthTx"] = qoSConfigQoSBandwidthTxJsonPath                            }
                                }
                                                                                                                                            if v, ok := d.GetOkExists("qo_s_config");  ok    {
                            qoSConfigQoSConcurrentConnectionsJsonPath, err := jsonpath.Get("$[0].qo_s[0].concurrent_connections", v)
            if err == nil && qoSConfigQoSConcurrentConnectionsJsonPath != "" {
                                                    request["QoS.ConcurrentConnections"] = qoSConfigQoSConcurrentConnectionsJsonPath                            }
                                }
                                                                                                                                            if v, ok := d.GetOkExists("qo_s_config");  ok    {
                            qoSConfigQoSPpsRxJsonPath, err := jsonpath.Get("$[0].qo_s[0].pps_rx", v)
            if err == nil && qoSConfigQoSPpsRxJsonPath != "" {
                                                    request["QoS.PpsRx"] = qoSConfigQoSPpsRxJsonPath                            }
                                }
                                                                                                                                            if v, ok := d.GetOkExists("qo_s_config");  ok    {
                            qoSConfigQoSPpsTxJsonPath, err := jsonpath.Get("$[0].qo_s[0].pps_tx", v)
            if err == nil && qoSConfigQoSPpsTxJsonPath != "" {
                                                    request["QoS.PpsTx"] = qoSConfigQoSPpsTxJsonPath                            }
                                }
                                                                                            wait := incrementalWait(3*time.Second, 5*time.Second)
        err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
                                    response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
                                        if err != nil {
                if  NeedRetry(err) {
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
                                if  target == false {
                                                            action := "DisableNetworkInterfaceQoS"
    request = make(map[string]interface{})
                            query = make(map[string]interface{})
                            request["NetworkInterfaceId"] = d.Id()
            request["RegionId"] = client.RegionId    
                                            wait := incrementalWait(3*time.Second, 5*time.Second)
        err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
                                    response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
                                        if err != nil {
                if  NeedRetry(err) {
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
            }
    
                                                                var err error
                                                                                action := "ModifyNetworkInterfaceAttribute"
    request = make(map[string]interface{})
                            query = make(map[string]interface{})
                            request["NetworkInterfaceId"] = d.Id()
            request["RegionId"] = client.RegionId    
                                                                                                            		if !d.IsNewResource() && d.HasChange("security_group_ids") {
	update = true
if v, ok := d.GetOk("security_group_ids");  (ok || d.HasChange("security_group_ids"))  {
									securityGroupIdMapsArray := convertToInterfaceArray(v)

		request["SecurityGroupId"] = securityGroupIdMapsArray
}
	}

                                                                                                                if !d.IsNewResource() && d.HasChange("delete_on_release") {
        update = true
                    request["DeleteOnRelease"] = d.Get("delete_on_release")}

                                                                                                                                        if !d.IsNewResource() && d.HasChange("queue_number") {
        update = true
                    request["QueueNumber"] = d.Get("queue_number")}

                                                                                                                                        if !d.IsNewResource() && d.HasChange("description") {
        update = true
                    request["Description"] = d.Get("description")}

                                                                                                                                        if !d.IsNewResource() && d.HasChange("network_interface_name") {
        update = true
                    request["NetworkInterfaceName"] = d.Get("network_interface_name")}

                                                                                            if update  {
            wait := incrementalWait(3*time.Second, 5*time.Second)
        err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
                                    response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
                                        if err != nil {
                if  NeedRetry(err) {
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
                                                                                action = "AssignPrivateIpAddresses"
    request = make(map[string]interface{})
                            query = make(map[string]interface{})
                            request["NetworkInterfaceId"] = d.Id()
            request["RegionId"] = client.RegionId    
                request["ClientToken"] = buildClientToken(action)
                                                                                                                    		if  d.HasChange("ipv4_prefix") {
	update = true
if v, ok := d.GetOk("ipv4_prefix");  (ok || d.HasChange("ipv4_prefix"))  {
									ipv4PrefixMapsArray := convertToInterfaceArray(v)

		request["Ipv4Prefix"] = ipv4PrefixMapsArray
}
	}

                                                                                                                                    		if !d.IsNewResource() && d.HasChange("private_ip_sets") {
	update = true
if v, ok := d.GetOk("private_ip_sets");  (ok || d.HasChange("private_ip_sets"))  {
			localData1, _ := jsonpath.Get("$[*].private_ip_address", v)
								localData1Array := convertToInterfaceArray(localData1)

		request["PrivateIpAddress"] = localData1Array
}
	}

                                                                    if update  {
            wait := incrementalWait(3*time.Second, 5*time.Second)
        err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
                                    response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
                                        if err != nil {
                if  IsExpectedErrors(err, []string{"InvalidOperation.InvalidEcsState","InvalidOperation.InvalidEniState","OperationConflict","ServiceUnavailable","InternalError"}) ||  NeedRetry(err) {
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
                                                                                action = "UnassignPrivateIpAddresses"
    request = make(map[string]interface{})
                            query = make(map[string]interface{})
                            request["NetworkInterfaceId"] = d.Id()
            request["RegionId"] = client.RegionId    
                                                                                                            		if  d.HasChange("ipv4_prefix") {
	update = true
if v, ok := d.GetOk("ipv4_prefix");  (ok || d.HasChange("ipv4_prefix"))  {
									ipv4PrefixMapsArray := convertToInterfaceArray(v)

		request["Ipv4Prefix"] = ipv4PrefixMapsArray
}
	}

                                                                                                                                    		if !d.IsNewResource() && d.HasChange("private_ip_sets") {
	update = true
if v, ok := d.GetOk("private_ip_sets");  (ok || d.HasChange("private_ip_sets"))  {
			localData1, _ := jsonpath.Get("$[*].private_ip_address", v)
								localData1Array := convertToInterfaceArray(localData1)

		request["PrivateIpAddress"] = localData1Array
}
	}

                                                                    if update  {
            wait := incrementalWait(3*time.Second, 5*time.Second)
        err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
                                    response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
                                        if err != nil {
                if  IsExpectedErrors(err, []string{"InvalidOperation.InvalidEcsState","InvalidOperation.InvalidEniState","OperationConflict","ServiceUnavailable","InternalError"}) ||  NeedRetry(err) {
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
                
                                    if !d.IsNewResource() && d.HasChange("ipv6_sets") {
                var err error
                                            oldEntry, newEntry := d.GetChange("ipv6_sets")
            removed := oldEntry
            added := newEntry

            
                        if len(removed.([]interface{})) > 0 {
                                                                                                    action := "UnassignIpv6Addresses"
    request = make(map[string]interface{})
                            query = make(map[string]interface{})
                            request["NetworkInterfaceId"] = d.Id()
            request["RegionId"] = client.RegionId    
                                                                                                                                            localData := removed.([]interface{})
    	                                                                            ipv6PrefixMapsArray := localData
        	    request["Ipv6Prefix"] = ipv6PrefixMapsArray


                                                                                                                                                                    localData1 := removed.([]interface{})
    	                                                                            ipv6AddressMapsArray := localData1
        	    request["Ipv6Address"] = ipv6AddressMapsArray


                                                                    wait := incrementalWait(3*time.Second, 5*time.Second)
        err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
                                    response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
                                        if err != nil {
                if  NeedRetry(err) {
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

                            if len(added.([]interface{})) > 0 {
                                                                                                                action := "AssignIpv6Addresses"
    request = make(map[string]interface{})
                            query = make(map[string]interface{})
                            request["NetworkInterfaceId"] = d.Id()
            request["RegionId"] = client.RegionId    
                request["ClientToken"] = buildClientToken(action)
                                                                                                                                                            localData := added.([]interface{})
    	                                                                            ipv6AddressMapsArray := localData
        	    request["Ipv6Address"] = ipv6AddressMapsArray


                                                                    wait := incrementalWait(3*time.Second, 5*time.Second)
        err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
                                    response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
                                        if err != nil {
                if  NeedRetry(err) {
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
                        }
                                        if  d.HasChange("tags") {
                ecsServiceV2 := EcsServiceV2{client}
        if err := ecsServiceV2.SetResourceTags(d, ""); err != nil {
            return WrapError(err)
        }
    }
                            d.Partial(false)
        return resourceAliCloudEcsNetworkInterfaceRead(d, meta)
}


func resourceAliCloudEcsNetworkInterfaceDelete(d *schema.ResourceData, meta interface{}) error {
                        
                                                                        client := meta.(*connectivity.AliyunClient)
                    action := "DeleteNetworkInterface"
            var request map[string]interface{}
        var response map[string]interface{}
                                        query := make(map[string]interface{})
                        var err error
        request = make(map[string]interface{})
                request["NetworkInterfaceId"] = d.Id()
            request["RegionId"] = client.RegionId    
    

                    wait := incrementalWait(3*time.Second, 5*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
                        response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
                    if err != nil {
            if  IsExpectedErrors(err, []string{"InvalidOperation.InvalidEcsState","InvalidOperation.InvalidEniState","OperationConflict","ServiceUnavailable","InternalError"}) ||  NeedRetry(err) {
                wait()
                return resource.RetryableError(err)
            }
                        return resource.NonRetryableError(err)
       }
       return nil
    })
    addDebug(action, response, request)

    if err != nil {
        if  NotFoundError(err) {
            return nil
        }
        return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
    }

        ecsServiceV2 := EcsServiceV2{client}
    stateConf := BuildStateConf([]string{}, []string{"''"}, d.Timeout(schema.TimeoutDelete),  30 * time.Second ,ecsServiceV2.EcsNetworkInterfaceStateRefreshFunc(d.Id(), "Status", []string {}))
    if _, err := stateConf.WaitForState(); err != nil {
        return WrapErrorf(err, IdMsg, d.Id())
    }
        
                        return nil
}




