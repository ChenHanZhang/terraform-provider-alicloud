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

func resourceAliCloudNATGatewayNatGateway() *schema.Resource {
  return &schema.Resource{
            Create: resourceAliCloudNATGatewayNatGatewayCreate,
                Read: resourceAliCloudNATGatewayNatGatewayRead,
                Update: resourceAliCloudNATGatewayNatGatewayUpdate,
            Delete: resourceAliCloudNATGatewayNatGatewayDelete,
    Importer: &schema.ResourceImporter{
      State: schema.ImportStatePassthrough,
    },
    Timeouts: &schema.ResourceTimeout{
              Create: schema.DefaultTimeout(5 * time.Minute),
                  Update: schema.DefaultTimeout(5 * time.Minute),
                  Delete: schema.DefaultTimeout(5 * time.Minute),
        },
            Schema: map[string]*schema.Schema {
                        "access_mode": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                    Computed: true,
                                                        ForceNew: true,
                                                                                                                                                                MaxItems: 1,
                Elem: &schema.Resource {
                              Schema: map[string]*schema.Schema {
                        "tunnel_type": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                    Computed: true,
                                                        ForceNew: true,
                                                                                                    ValidateFunc: StringInSlice([]string{"geneve"}, false),
                                                                                    },
                                    "mode_value": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                    Computed: true,
                                                        ForceNew: true,
                                                                                                    ValidateFunc: StringInSlice([]string{"route","tunnel"}, false),
                                                                                    },
            },
                },
                        },
                                        "auto_pay": {
            Type:schema.TypeBool,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                        "create_time": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "deletion_protection": {
            Type:schema.TypeBool,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "description": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                        "eip_bind_mode": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                                                                                                ValidateFunc: StringInSlice([]string{"NAT","MULTI_BINDED"}, false),
                                                                                    },
                                    "enable_session_log": {
            Type:schema.TypeBool,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                        "force_delete": {
            Type:schema.TypeBool,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "forward_table_ids": {
            Type:schema.TypeList,
                                                                                                            Computed: true,
                                                                                                                                                                                                                                Elem:&schema.Schema{Type: schema.TypeString},
                                                                                        },
                                        "icmp_reply_enabled": {
            Type:schema.TypeBool,
                                                                                            Optional: true,
                                                    Computed: true,
                                                                                                                                                                                        },
                                    "internet_charge_type": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                    Computed: true,
                                                        ForceNew: true,
                                                                                                    ValidateFunc: StringInSlice([]string{"PayByLcu","PayBySpec"}, false),
                                                                                    },
                                        "log_delivery": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                                        ForceNew: true,
                                                                                                                                                                MaxItems: 1,
                Elem: &schema.Resource {
                              Schema: map[string]*schema.Schema {
                        "deliver_logs_error_message": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "log_delivery_type": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                ValidateFunc: StringInSlice([]string{"sls"}, false),
                                                                                    },
                                    "delivery_status": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "log_destination": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
            },
                },
                        },
                                        "nat_gateway_name": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                    Computed: true,
                                                                                                                                                                                        },
                                    "nat_gateway_private_info": {
            Type:schema.TypeList,
                                                            Required: true,
                                                                                                        ForceNew: true,
                                                                                                                                                                MaxItems: 1,
                Elem: &schema.Resource {
                              Schema: map[string]*schema.Schema {
                        "iz_no": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "max_session_establish_rate": {
            Type:schema.TypeInt,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "private_ip_address": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "max_session_quota": {
            Type:schema.TypeInt,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "max_bandwidth": {
            Type:schema.TypeInt,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "vswitch_id": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                        ForceNew: true,
                                                                                                                                                            },
                                    "eni_instance_id": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "eni_type": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
            },
                },
                        },
                                    "nat_type": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                        ForceNew: true,
                                                                                                    ValidateFunc: StringInSlice([]string{"Enhanced"}, false),
                                                                                    },
                                    "network_type": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                        ForceNew: true,
                                                                                                    ValidateFunc: StringInSlice([]string{"internet","intranet"}, false),
                                                                                    },
                                            "payment_type": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                    Computed: true,
                                                        ForceNew: true,
                                                                                                    ValidateFunc: StringInSlice([]string{"PayAsYouGo","Subscription"}, false),
                                                                                    },
                                    "private_link_enabled": {
            Type:schema.TypeBool,
                                                                                            Optional: true,
                                                                        ForceNew: true,
                                                                                                                                                            },
                                        "region_id": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "security_protection_enabled": {
            Type:schema.TypeBool,
                                                                                            Optional: true,
                                                                                    Deprecated: "Field 'security_protection_enabled' has been deprecated from provider version 1.270.0. Whether to enable the firewall function. Valid values:-**false** (default): not enabled.-**true**: on.",
                                                                                                                                                },
                                    "snat_table_ids": {
            Type:schema.TypeList,
                                                                                                            Computed: true,
                                                                                                                                                                                                                                Elem:&schema.Schema{Type: schema.TypeString},
                                                                                        },
                                    "spec": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                        ForceNew: true,
                                                                                                    ValidateFunc: StringInSlice([]string{"Small","Middle","Large"}, false),
                                                                                    },
                                    "status": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "vpc_id": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                        ForceNew: true,
                                                                                                                                                            },
            },
  }
}

func resourceAliCloudNATGatewayNatGatewayCreate(d *schema.ResourceData, meta interface{}) error {

        
        
                                                                                    client := meta.(*connectivity.AliyunClient)
             
        action := "CreateNatGateway"
            var request map[string]interface{}
        var response map[string]interface{}
                            query := make(map[string]interface{})
                    var err error
        request = make(map[string]interface{})
                request["RegionId"] = client.RegionId
                        request["ClientToken"] = buildClientToken(action)
            
                                                                                request["EipBindMode"] = d.Get("eip_bind_mode")
                                                                                                                            accessMode := make(map[string]interface{})
    
if v := d.Get("access_mode");  !IsNil(v)   {
                                tunnelType1, _ := jsonpath.Get("$[0].tunnel_type", v)
        if tunnelType1 != nil                && tunnelType1 != "" {
                accessMode["TunnelType"] = tunnelType1        }
                                            modeValue1, _ := jsonpath.Get("$[0].mode_value", v)
        if modeValue1 != nil                && modeValue1 != "" {
                accessMode["ModeValue"] = modeValue1        }
            
    accessModeJson, err := json.Marshal(accessMode)
    if err != nil {
    return WrapError(err)
    }
    request["AccessMode"] = string(accessModeJson)
}

                                                                                                natGatewayPrivateInfoVSwitchIdJsonPath, err := jsonpath.Get("$[0].vswitch_id", d.Get("nat_gateway_private_info"))
            if err == nil {
                                                    request["VSwitchId"] = natGatewayPrivateInfoVSwitchIdJsonPath                            }
                
                                                                                if v, ok := d.GetOk("payment_type");  ok    {
                                                                        request["InstanceChargeType"] = convertNATGatewayNatGatewayInstanceChargeTypeRequest(v.(string))
                            }
                                                                                if v, ok := d.GetOk("nat_gateway_name");  ok    {
                                                            request["Name"] = v                                        }
                                                                                if v, ok := d.GetOk("description");  ok    {
                                                            request["Description"] = v                                        }
                                                                                if v, ok := d.GetOkExists("auto_pay");  ok    {
                                                            request["AutoPay"] = v                                        }
                                                                                                request["NetworkType"] = d.Get("network_type")
                                                                                if v, ok := d.GetOkExists("icmp_reply_enabled");  ok    {
                                                            request["IcmpReplyEnabled"] = v                                        }
                                                                                if v, ok := d.GetOkExists("private_link_enabled");  ok    {
                                                            request["PrivateLinkEnabled"] = v                                        }
                                                                                if v, ok := d.GetOkExists("security_protection_enabled");  ok    {
                                                            request["SecurityProtectionEnabled"] = v                                        }
                                                                                                request["VpcId"] = d.Get("vpc_id")
                                                                                if v, ok := d.GetOk("internet_charge_type");  ok    {
                                                            request["InternetChargeType"] = v                                        }
                                                                                if v, ok := d.GetOk("spec");  ok    {
                                                            request["Spec"] = v                                        }
                                                                                                request["NatType"] = d.Get("nat_type")
                                                    wait := incrementalWait(3*time.Second, 5*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
                        response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
                    if err != nil {
                        if  IsExpectedErrors(err, []string{"IncorrectStatus.VSWITCH","OperationConflict"}) ||  NeedRetry(err) {
                wait()
                return resource.RetryableError(err)
            }
                        return resource.NonRetryableError(err)
        }
        return nil
    })
    addDebug(action, response, request)

    if err != nil {
        return WrapErrorf(err, DefaultErrorMsg, "alicloud_nat_gateway", action, AlibabaCloudSdkGoERROR)
    }
        
                        d.SetId(fmt.Sprint(response["NatGatewayId"]))        

                nATGatewayServiceV2 := NATGatewayServiceV2{client}
        stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutCreate),  5 * time.Second ,nATGatewayServiceV2.NATGatewayNatGatewayStateRefreshFunc(d.Id(), "Status", []string {}))
        if _, err := stateConf.WaitForState(); err != nil {
            return WrapErrorf(err, IdMsg, d.Id())
        }
        
        
        
        return resourceAliCloudNATGatewayNatGatewayUpdate(d, meta)
    }


    		func resourceAliCloudNATGatewayNatGatewayRead(d *schema.ResourceData, meta interface{}) error {
	    		client := meta.(*connectivity.AliyunClient)
		nATGatewayServiceV2 := NATGatewayServiceV2{client}

												objectRaw, err := nATGatewayServiceV2.DescribeNATGatewayNatGateway(d.Id())
				if err != nil  {
					if !d.IsNewResource() && NotFoundError(err) {
				log.Printf("[DEBUG] Resource alicloud_nat_gateway DescribeNATGatewayNatGateway Failed!!! %s", err)
				d.SetId("")
				return nil
			}
					return WrapError(err)
		}

				

                            
                                        d.Set("create_time", objectRaw["CreationTime"])
                                            d.Set("deletion_protection", objectRaw["DeletionProtection"])
                                            d.Set("description", objectRaw["Description"])
                                            d.Set("eip_bind_mode", objectRaw["EipBindMode"])
                                            d.Set("enable_session_log", formatBool(objectRaw["EnableSessionLog"]))
                                            d.Set("icmp_reply_enabled", objectRaw["IcmpReplyEnabled"])
                                            d.Set("internet_charge_type", objectRaw["InternetChargeType"])
                                            d.Set("nat_gateway_name", objectRaw["Name"])
                                            d.Set("nat_type", objectRaw["NatType"])
                                            d.Set("payment_type", convertNATGatewayNatGatewayNatGatewaysNatGatewayInstanceChargeTypeResponse(objectRaw["InstanceChargeType"]))
                                            d.Set("private_link_enabled", objectRaw["PrivateLinkEnabled"])
                                            d.Set("region_id", objectRaw["RegionId"])
                                            d.Set("spec", objectRaw["Spec"])
                                            d.Set("status", objectRaw["Status"])
                                            d.Set("vpc_id", objectRaw["VpcId"])
                
                                                                    accessModeMaps := make([]map[string]interface{}, 0)
        accessModeMap := make(map[string]interface{})
                accessModeRaw := make(map[string]interface{})
if objectRaw["AccessMode"] != nil {
accessModeRaw = objectRaw["AccessMode"].(map[string]interface{})}
if len(accessModeRaw) > 0 {
                                        accessModeMap["mode_value"] = accessModeRaw["ModeValue"]
                                            accessModeMap["tunnel_type"] = accessModeRaw["TunnelType"]
                
                                                                                                        

                                                                                                                                                                                                                                                                                                                                                                                                                                                            forwardTableIdRaw, _ := jsonpath.Get("$.ForwardTableIds.ForwardTableId", objectRaw)
                                    d.Set("forward_table_ids", forwardTableIdRaw)
                                                                                                                                                                                                                                            natGatewayPrivateInfoMaps := make([]map[string]interface{}, 0)
        natGatewayPrivateInfoMap := make(map[string]interface{})
                natGatewayPrivateInfoRaw := make(map[string]interface{})
if objectRaw["NatGatewayPrivateInfo"] != nil {
natGatewayPrivateInfoRaw = objectRaw["NatGatewayPrivateInfo"].(map[string]interface{})}
if len(natGatewayPrivateInfoRaw) > 0 {
                                        natGatewayPrivateInfoMap["eni_instance_id"] = natGatewayPrivateInfoRaw["EniInstanceId"]
                                            natGatewayPrivateInfoMap["eni_type"] = natGatewayPrivateInfoRaw["EniType"]
                                            natGatewayPrivateInfoMap["iz_no"] = natGatewayPrivateInfoRaw["IzNo"]
                                            natGatewayPrivateInfoMap["max_bandwidth"] = natGatewayPrivateInfoRaw["MaxBandwidth"]
                                            natGatewayPrivateInfoMap["max_session_establish_rate"] = natGatewayPrivateInfoRaw["MaxSessionEstablishRate"]
                                            natGatewayPrivateInfoMap["max_session_quota"] = natGatewayPrivateInfoRaw["MaxSessionQuota"]
                                            natGatewayPrivateInfoMap["private_ip_address"] = natGatewayPrivateInfoRaw["PrivateIpAddress"]
                                            natGatewayPrivateInfoMap["vswitch_id"] = natGatewayPrivateInfoRaw["VswitchId"]
                
                                                                                                                                                                                                                                                                                                                                                                                                                                

                                                                                                                                                                                                                                                                                                                                                                                                        snatTableIdRaw, _ := jsonpath.Get("$.SnatTableIds.SnatTableId", objectRaw)
                                    d.Set("snat_table_ids", snatTableIdRaw)
                                                                                                                                                                        
																						objectRaw, err = nATGatewayServiceV2.DescribeNatGatewayGetNatGatewayAttribute(d.Id())
							if err != nil  && !NotFoundError(err)  {
					return WrapError(err)
		}

				

                            
                                        d.Set("enable_session_log", objectRaw["EnableSessionLog"])
                                            d.Set("nat_gateway_name", objectRaw["Name"])
                                            d.Set("nat_type", objectRaw["NatType"])
                                            d.Set("network_type", objectRaw["NetworkType"])
                                            d.Set("private_link_enabled", objectRaw["PrivateLinkEnabled"])
                                            d.Set("region_id", objectRaw["RegionId"])
                                            d.Set("status", objectRaw["Status"])
                                            d.Set("vpc_id", objectRaw["VpcId"])
                
            billingConfigRawObj, _ := jsonpath.Get("$.BillingConfig", objectRaw)
billingConfigRaw := make(map[string]interface{})
if billingConfigRawObj != nil {
billingConfigRaw = billingConfigRawObj.(map[string]interface{})
}
                                        d.Set("internet_charge_type", convertNATGatewayNatGatewayBillingConfigInternetChargeTypeResponse(billingConfigRaw["InternetChargeType"]))
                                            d.Set("payment_type", billingConfigRaw["InstanceChargeType"])
                                            d.Set("spec", billingConfigRaw["Spec"])
                
                                                                    
        
                accessModeRaw := make(map[string]interface{})
if objectRaw["AccessMode"] != nil {
accessModeRaw = objectRaw["AccessMode"].(map[string]interface{})}
if len(accessModeRaw) > 0 {
                                        accessModeMap["mode_value"] = accessModeRaw["ModeValue"]
                                            accessModeMap["tunnel_type"] = accessModeRaw["TunnelType"]
                
                                                                                                        accessModeMaps = append(accessModeMaps, accessModeMap)
}
if err := d.Set("access_mode", accessModeMaps); err != nil {
return err
}
                                                                                                                                                                                                                                            forwardTableRaw := make([]interface{}, 0)
if objectRaw["ForwardTable"] != nil {
forwardTableRaw = convertToInterfaceArray(objectRaw["ForwardTable"])}

                                    d.Set("forward_table_ids", forwardTableRaw)
                                                                                                                                    logDeliveryMaps := make([]map[string]interface{}, 0)
        logDeliveryMap := make(map[string]interface{})
                logDeliveryRaw := make(map[string]interface{})
if objectRaw["LogDelivery"] != nil {
logDeliveryRaw = objectRaw["LogDelivery"].(map[string]interface{})}
if len(logDeliveryRaw) > 0 {
                                        logDeliveryMap["deliver_logs_error_message"] = logDeliveryRaw["DeliverLogsErrorMessage"]
                                            logDeliveryMap["delivery_status"] = logDeliveryRaw["DeliveryStatus"]
                                            logDeliveryMap["log_delivery_type"] = logDeliveryRaw["LogDeliveryType"]
                                            logDeliveryMap["log_destination"] = logDeliveryRaw["LogDestination"]
                
                                                                                                                                                                                                                logDeliveryMaps = append(logDeliveryMaps, logDeliveryMap)
}
if err := d.Set("log_delivery", logDeliveryMaps); err != nil {
return err
}
                                                                                                                                                    
        
                privateInfoRaw := make(map[string]interface{})
if objectRaw["PrivateInfo"] != nil {
privateInfoRaw = objectRaw["PrivateInfo"].(map[string]interface{})}
if len(privateInfoRaw) > 0 {
                                        natGatewayPrivateInfoMap["eni_instance_id"] = privateInfoRaw["EniInstanceId"]
                                            natGatewayPrivateInfoMap["iz_no"] = privateInfoRaw["IzNo"]
                                            natGatewayPrivateInfoMap["private_ip_address"] = privateInfoRaw["PrivateIpAddress"]
                                            natGatewayPrivateInfoMap["vswitch_id"] = privateInfoRaw["VswitchId"]
                
                                                                                                                                                                                                                natGatewayPrivateInfoMaps = append(natGatewayPrivateInfoMaps, natGatewayPrivateInfoMap)
}
if err := d.Set("nat_gateway_private_info", natGatewayPrivateInfoMaps); err != nil {
return err
}
                                                                                                                                                                                                                                                                                                                                                                                                                                                            snatTableRaw := make([]interface{}, 0)
if objectRaw["SnatTable"] != nil {
snatTableRaw = convertToInterfaceArray(objectRaw["SnatTable"])}

                                    d.Set("snat_table_ids", snatTableRaw)
                                                                                                                                                                        
							
        
        
		        	return nil
}


    func resourceAliCloudNATGatewayNatGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
                client := meta.(*connectivity.AliyunClient)
        var request map[string]interface{}
    var response map[string]interface{}
            var query map[string]interface{}
                        update := false
        
        
        
                                                                var err error
                                                                                action := "ModifyNatGatewayAttribute"
    request = make(map[string]interface{})
                            query = make(map[string]interface{})
                            request["NatGatewayId"] = d.Id()
            request["RegionId"] = client.RegionId    
                                                                                        if !d.IsNewResource() && d.HasChange("eip_bind_mode") {
        update = true
        }
                    request["EipBindMode"] = d.Get("eip_bind_mode")
                                                                                                                                        if d.HasChange("enable_session_log") {
        update = true
                    request["EnableSessionLog"] = d.Get("enable_session_log")}

                                                                                                                                        if !d.IsNewResource() && d.HasChange("nat_gateway_name") {
        update = true
                    request["Name"] = d.Get("nat_gateway_name")}

                                                                                                                                        if !d.IsNewResource() && d.HasChange("description") {
        update = true
                    request["Description"] = d.Get("description")}

                                                                                                                                                                            if  d.HasChange("log_delivery") {
    update = true
                logDelivery := make(map[string]interface{})
    
if v := d.Get("log_delivery");  v != nil  {
                                logDestination1, _ := jsonpath.Get("$[0].log_destination", v)
        if logDestination1 != nil&&  logDestination1 != ""  {
                            logDelivery["LogDestination"] = logDestination1                }
                                            logDeliveryType1, _ := jsonpath.Get("$[0].log_delivery_type", v)
        if logDeliveryType1 != nil&&  logDeliveryType1 != ""  {
                            logDelivery["LogDeliveryType"] = logDeliveryType1                }
            
    logDeliveryJson, err := json.Marshal(logDelivery)
    if err != nil {
    return WrapError(err)
    }
    request["LogDelivery"] = string(logDeliveryJson)
}
    }

                                                                                                        if !d.IsNewResource() && d.HasChange("icmp_reply_enabled") {
        update = true
                    request["IcmpReplyEnabled"] = d.Get("icmp_reply_enabled")}

                                                                                            if update  {
            wait := incrementalWait(3*time.Second, 5*time.Second)
        err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
                                    response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
                                        if err != nil {
                if  IsExpectedErrors(err, []string{"IncorrectStatus.NATGW"}) ||  NeedRetry(err) {
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
                
                    return resourceAliCloudNATGatewayNatGatewayRead(d, meta)
}


func resourceAliCloudNATGatewayNatGatewayDelete(d *schema.ResourceData, meta interface{}) error {
                        
                                                                        client := meta.(*connectivity.AliyunClient)
                    action := "DeleteNatGateway"
            var request map[string]interface{}
        var response map[string]interface{}
                                        query := make(map[string]interface{})
                        var err error
        request = make(map[string]interface{})
                request["NatGatewayId"] = d.Id()
            request["RegionId"] = client.RegionId    
    

                                                if v, ok := d.GetOkExists("force_delete");  ok    {
                                                            request["Force"] = v                                        }
                                            wait := incrementalWait(3*time.Second, 5*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
                        response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
                    if err != nil {
            if  IsExpectedErrors(err, []string{"InvalidStatus.NatGateway","DependencyViolation.EIPS"}) ||  NeedRetry(err) {
                wait()
                return resource.RetryableError(err)
            }
                        return resource.NonRetryableError(err)
       }
       return nil
    })
    addDebug(action, response, request)

    if err != nil {
        if IsExpectedErrors(err, []string{"InstanceNotExist.NatGatewayId"}) ||  NotFoundError(err) {
            return nil
        }
        return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
    }

        nATGatewayServiceV2 := NATGatewayServiceV2{client}
    stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete),  5 * time.Second ,nATGatewayServiceV2.NATGatewayNatGatewayStateRefreshFunc(d.Id(), "Status", []string {}))
    if _, err := stateConf.WaitForState(); err != nil {
        return WrapErrorf(err, IdMsg, d.Id())
    }
        
                        return nil
}


  


    func convertNATGatewayNatGatewayNatGatewaysNatGatewayInstanceChargeTypeResponse(source interface{}) interface{} {
    source = fmt.Sprint(source)
    switch source {
                                case "PostPaid":
                return "PayAsYouGo"
                        case "PrePaid":
                return "Subscription"
                        }
    return source
}
    func convertNATGatewayNatGatewayBillingConfigInternetChargeTypeResponse(source interface{}) interface{} {
    source = fmt.Sprint(source)
    switch source {
                                case "PostPaid":
                return "PayAsYouGo"
                        case "PrePaid":
                return "Subscription"
                        }
    return source
}
    func convertNATGatewayNatGatewayInstanceChargeTypeRequest(source interface{}) interface{} {
    source = fmt.Sprint(source)
    switch source {
                                    case "PayAsYouGo":
                return "PostPaid"
                            case "Subscription":
                return "PrePaid"
                        }
    return source
}
