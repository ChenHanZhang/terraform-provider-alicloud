---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_db_instance"
description: |-
  Provides a Alicloud RDS Db Instance resource.
---

# alicloud_db_instance

Provides a RDS Db Instance resource.

Database instance resource  .

For information about RDS Db Instance and how to use it, see [What is Db Instance](https://www.alibabacloud.com/help/en/doc-detail/26092.htm).

-> **NOTE:** Available since v1.155.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_db_zones" "example" {
  engine                   = "MySQL"
  engine_version           = "8.0"
  instance_charge_type     = "PostPaid"
  category                 = "Basic"
  db_instance_storage_type = "cloud_essd"
}

data "alicloud_db_instance_classes" "example" {
  zone_id                  = data.alicloud_db_zones.example.zones.0.id
  engine                   = "MySQL"
  engine_version           = "8.0"
  category                 = "Basic"
  db_instance_storage_type = "cloud_essd"
  instance_charge_type     = "PostPaid"
}

resource "alicloud_vpc" "example" {
  vpc_name   = "terraform-example"
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "example" {
  vpc_id       = alicloud_vpc.example.id
  cidr_block   = "172.16.0.0/24"
  zone_id      = data.alicloud_db_zones.example.zones.0.id
  vswitch_name = "terraform-example"
}

resource "alicloud_security_group" "example" {
  name   = "terraform-example"
  vpc_id = alicloud_vpc.example.id
}

resource "alicloud_db_instance" "example" {
  engine                   = "MySQL"
  engine_version           = "8.0"
  instance_type            = data.alicloud_db_instance_classes.example.instance_classes.0.instance_class
  instance_storage         = data.alicloud_db_instance_classes.example.instance_classes.0.storage_range.min
  instance_charge_type     = "Postpaid"
  instance_name            = "terraform-example"
  vswitch_id               = alicloud_vswitch.example.id
  monitoring_period        = "60"
  db_instance_storage_type = "cloud_essd"
  security_group_ids       = [alicloud_security_group.example.id]
}
```

## Argument Reference

The following arguments are supported:
* `allow_major_version_upgrade` - (Optional, Available since v1.274.0) Major version upgrade for SQL Server instances [Major Version Upgrade](https://help.aliyun.com/document_detail/127458.html). Valid values:  
  - `true`: Perform the upgrade.  
  - `false` (default): Do not perform the upgrade.  

-> **NOTE:**  - When performing a major version upgrade, you must also specify the following required parameters: DBInstanceId, EngineVersion, DBInstanceClass, Category, ZoneId, and VSwitchId.  

-> **NOTE:**  - Additionally, if you want to upgrade to a High Availability or Cluster Edition instance, you must also specify the ZoneIdSlave1 parameter.  


-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `auto_create_proxy` - (Optional, Available since v1.274.0) Specifies whether to automatically create a proxy. Valid values:  
  - `true`: Enables automatic proxy creation. By default, a general-purpose proxy is created.  
  - `false`: Disables automatic proxy creation.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_pause` - (Optional, Available since v1.274.0) Specifies whether to enable intelligent pause and resume for Serverless instances. Valid values:  
* `true`: Enables the feature.  
* `false`: Disables the feature (default).  

-> **NOTE:**  This parameter applies only to MySQL and PostgreSQL Serverless instances. If no connections are established for 10 minutes, the instance enters the paused state and automatically resumes when a new connection is initiated.  


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_pay` - (Optional, Available since v1.274.0) Specifies whether to enable automatic payment. Valid values:
  - `true`: Automatic payment is enabled. You must ensure that your account balance is sufficient.
  - `false`: Only an order is generated without charging your account.



-> **NOTE:**  The default value is true. If your payment method has insufficient funds, you can set AutoPay to false. In this case, an unpaid order is generated, and you can log on to the ApsaraDB RDS console to complete the payment manually.

-> **NOTE:** .


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_renew` - (Optional, Available since v1.34.0) Specifies whether auto-renewal is enabled. Valid values:

* `true`: Enables auto-renewal.
* `false`: Disables auto-renewal.

-> **NOTE:**  * This parameter takes effect only when changing the billing method from pay-as-you-go to subscription.

-> **NOTE:**  * Any string other than `true` is treated as `false`.


-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `babelfish_config` - (Optional, ForceNew, Computed, Set, Available since v1.176.0) Configuration information for Babelfish for RDS PostgreSQL.

-> **NOTE:**  This parameter applies only to Babelfish for RDS PostgreSQL instances. For more information, see [Introduction to Babelfish](https://help.aliyun.com/document_detail/428613.html).
 See [`babelfish_config`](#babelfish_config) below.
* `babelfish_port` - (Optional, Computed, Available since v1.176.0) The TDS port number for Babelfish for ApsaraDB RDS for PostgreSQL.  

-> **NOTE:**  This parameter applies only to ApsaraDB RDS for PostgreSQL instances. For more information about Babelfish for ApsaraDB RDS for PostgreSQL, see [Introduction to Babelfish](https://help.aliyun.com/document_detail/428613.html).  


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `bpe_enabled` - (Optional, Available since v1.274.0) This parameter is deprecated and does not require configuration.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `bursting_enabled` - (Optional, Available since v1.242.0) The switch for the I/O burst performance feature of ultra disks. Valid values:
* `true`: Enable.
* `false`: Disable.

-> **NOTE:**  For more information about the I/O burst performance feature of ultra disks, see [What are ultra disks?](https://help.aliyun.com/document_detail/2340501.html).


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `business_info` - (Optional, Available since v1.274.0) Business extension parameters.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `category` - (Optional, Computed, Available since v1.187.0) Instance series. Valid values:
  - General-purpose instances
    * `Basic`: Basic series.
    * `HighAvailability`: High-availability series.
    * `cluster`: MySQL or PostgreSQL cluster series.
    * `AlwaysOn`: SQL Server cluster series.
    * `Finance`: Three-node enterprise series.
    > This parameter is required when creating SQL Server Enterprise Cluster Edition, Basic series Standard Edition, or Basic series Enterprise Edition. For example, when creating a 2022 Enterprise Cluster Edition (2022_ent) in the Basic series, you must specify `Basic` for this parameter.
  - Serverless instances
    * `serverless_basic`: Serverless Basic series. (Applies only to MySQL and PostgreSQL)
    * `serverless_standard`: Serverless High-availability series. (Applies only to MySQL and PostgreSQL)
    * `serverless_ha`: SQL Server Serverless High-availability series.

    > This parameter is required when **PayType=Serverless**.
* `cold_data_enabled` - (Optional) Switch for the [data archiving](https://help.aliyun.com/document_detail/2701832.html) feature of enhanced cloud disks. Valid values:
  - `true`: Enabled.
  - `false`: Disabled.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `compression_mode` - (Optional, Available since v1.274.0) MySQL [storage compression feature](https://help.aliyun.com/document_detail/2861985.html). Valid values:
  - `on`: Enables storage compression
  - `off`: Disables storage compression.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `config_value` - (Optional, Available since v1.274.0) The value of the configuration item to be modified, used together with ConfigName.

RDS PostgreSQL Configuration Values
  - PgBouncer feature: `true` (enabled), `false` (disabled).
  - Cloud disk encryption feature:
  - `ServiceKey`: Use a key automatically generated by Alibaba Cloud, that is, enable cloud disk encryption using the RDS-managed service key (Default Service CMK).
  - ****: Use a custom key to enable cloud disk encryption or replace the current key. For example, `494c98ce-f2b5-48ab-96ab-36c986b6****`.
  - `disabled`: Disable cloud disk encryption.
  - One-click preparation of prerequisites required for creating a DuckDB analytics instance: `duckdb_prepare_dependency`
  - Batch set primary instance databases as DuckDB columnar storage databases, in JSON format. For example: `{"dbNames": "db1,db2,db3", "accountName": "yourSuperAccountName"}`, where:
  - `dbNames`: Names of the databases to be converted to DuckDB columnar storage databases. Separate multiple database names with commas (`,`).
  - `accountName`: A high-privilege account. Only one high-privilege account needs to be specified.
  - Enable CONNECT RLS on the instance to control database visibility: `true` to enable.
  - Enable CONNECT RLS on specific databases to control database visibility: Specify the **database names** governed by CONNECT RLS. Separate multiple database names with commas (`,`), such as **testdb1,testdb2**. When connecting to a database with CONNECT RLS enabled, users will see an accurate list of databases based on whether they have CONNECT privileges to other databases.

RDS SQL Server Configuration Values

  - Error log cleanup feature: `1` (confirm cleanup).
  - Cloud disk encryption feature (**this feature cannot be disabled after it is enabled**):
  - `serviceKey`: Use a key automatically generated by Alibaba Cloud, that is, enable cloud disk encryption using the RDS-managed service key (Default Service CMK).
  - ****: Use a custom key to enable cloud disk encryption or replace the current key. For example, `494c98ce-f2b5-48ab-96ab-36c986b6****`.



  - Simple recovery model: `simple` (enable simple recovery).
  - Error log cleanup feature: `1` (confirm cleanup).
  - Cloud disk encryption feature (**this feature cannot be disabled after it is enabled**):
  - `serviceKey`: Use a key automatically generated by Alibaba Cloud, that is, enable cloud disk encryption using the RDS-managed service key (Default Service CMK).
  - ****: Use a custom key to enable cloud disk encryption or replace the current key. For example, `494c98ce-f2b5-48ab-96ab-36c986b6****`.


.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `connection_mode` - (Optional, Available since v1.274.0) Connection mode of the instance. Valid values:
* `Standard`: Standard connection mode.
* `Safe`: Database proxy mode.

By default, the RDS system assigns the connection mode.

-> **NOTE:**  SQL Server 2012, 2016, and 2017 support only the standard connection mode.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `connection_string_prefix` - (Optional, Computed, Available since v1.126.0) Read-only endpoint prefix. It must be unique, consist of lowercase letters and hyphens, start with a letter, and be no longer than 30 characters.

-> **NOTE:**  By default, the prefix is formed by concatenating the instance name and the string "rw".


-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `current_connection_string` - (Optional, Available since v1.274.0) A current connection string of the instance, which can be an internal or public endpoint, or a classic network endpoint in hybrid access mode.  

-> **NOTE:**  Read/write splitting endpoints cannot be modified.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `db_instance_description` - (Optional, ForceNew, Available since v1.274.0) Instance name, which must be 2 to 256 characters in length. It must start with a Chinese character or an English letter and can contain digits, Chinese characters, English letters, underscores (_), and hyphens (-).

-> **NOTE:**  It cannot start with `http://` or `https://`.

* `db_instance_storage` - (Required, Int, Available since v1.274.0) Instance storage capacity, in GB. The value increments in steps of 5 GB. For more information, see the [Instance types](https://help.aliyun.com/document_detail/26312.html).

-> **NOTE:**  By default, the storage capacity is the same as that of the primary instance.

* `db_instance_class` - (Required, Available since v1.274.0) The instance class after upgrade. The CPU and memory configuration must be greater than or equal to those of the original instance class. You do not need to specify this parameter when `UpgradeMode` is set to `inPlaceUpgrade` or `zeroDownTimeUpgrade`.

For example, if the original instance class is `pg.n2.small.2c` (1 vCPU and 2 GB memory), you can upgrade it to `pg.n2.medium.2c` (2 vCPUs and 4 GB memory).

-> **NOTE:**  For the instance class codes of ApsaraDB RDS for PostgreSQL, see [ApsaraDB RDS for PostgreSQL primary instance types](https://help.aliyun.com/document_detail/276990.html).

* `db_instance_net_type` - (Required, ForceNew, Available since v1.274.0) The network connection type of the instance. Valid values:  
  - `Internet`: Public network connection.  
  - `Intranet`: Internal network connection.  
* `db_instance_storage_type` - (Optional, Computed, Available since v1.68.0) The storage type of the instance. Valid values:
* `local_ssd`: High-performance local disk (recommended).
* `general_essd`: High-performance cloud disk (recommended).
* `cloud_essd`: ESSD PL1 cloud disk.
* `cloud_essd2`: ESSD PL2 cloud disk.
* `cloud_essd3`: ESSD PL3 cloud disk.
* `cloud_ssd`: SSD cloud disk (not recommended; sales have been discontinued in some regions).

The default value of this parameter is automatically determined based on the instance class specified by the `DBInstanceClass` parameter:
* If the instance class corresponds to a high-performance local disk, the default value is `local_ssd`.
* If the instance class corresponds to a cloud disk, the default value is `cloud_essd`.

-> **NOTE:**  Serverless instances support only ESSD PL1 cloud disks and high-performance cloud disks.

* `db_is_ignore_case` - (Optional, Computed, Available since v1.168.0) Specifies whether table names are case-sensitive. Valid values:  
  - `true`: Case-insensitive (default).  
  - `false`: Case-sensitive.  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `db_param_group_id` - (Optional) The parameter template ID. You can call DescribeParameterGroups to query it.  

-> **NOTE:**  This parameter is supported only for MySQL and PostgreSQL. If you do not specify this parameter, the system default parameter template is used. You can also create a custom parameter template and specify it here.


-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `db_time_zone` - (Optional, Computed, Available since v1.136.0) The time zone of the instance. This parameter takes effect only when `Engine` is set to `MySQL` or `PostgreSQL`.  
  - When `Engine` is `MySQL`:  
    - This parameter configures the UTC time zone. Valid values range from **-12:59** to **+13:00**.  
    - For high-performance local disk instances, you can use named time zones, such as Asia/Hong_Kong. For more information about named time zones, see [Named Time Zone Reference](https://help.aliyun.com/document_detail/297356.html).  
  - When `Engine` is `PostgreSQL`:  
    - This parameter configures a named time zone. UTC time zones are not supported. For more information about named time zones, see [Named Time Zone Reference](https://help.aliyun.com/document_detail/297356.html).  
    - This parameter is configurable only for PostgreSQL instances with cloud disks.  

-> **NOTE:**  - You can configure the time zone when purchasing a primary instance. Read-only instances do not support time zone configuration and inherit the time zone of the primary instance.  

-> **NOTE:**  - If you do not configure this parameter, the system selects a default time zone based on the region where you purchase the instance.


-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `deletion_protection` - (Optional, ForceNew, Available since v1.165.0) Indicates whether RDS instance release protection is enabled. Valid values:
* `true`: Enabled.
* `false`: Disabled (default).

-> **NOTE:**  This feature is supported only when the **billing method** is **Pay-As-You-Go**.

* `direction` - (Optional, Available since v1.209.1) The instance specification change type. Valid values:
  - `Up` (default): Upgrades a subscription instance or upgrades/downgrades a pay-as-you-go instance.
  - `Down`: Downgrades a subscription instance.
  - `TempUpgrade`: Specifies an elastic specification change for a subscription SQL Server instance. This value is required for elastic specification changes.
  - `Serverless`: Configures elasticity settings for a Serverless instance.

-> **NOTE:**  If you are changing only the `DBInstanceStorageType` parameter—for example, from SSD cloud disk to ESSD cloud disk—leave this parameter empty.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `effective_time` - (Optional, Available since v1.191.0) The time when the new configuration takes effect. Valid values:

-> **NOTE:**  **Changing certain configurations may affect your instance**. Carefully review the [Impact section in the feature documentation](https://help.aliyun.com/document_detail/96061.html) before setting this parameter. We recommend performing such changes during off-peak business hours.
* `Immediate` (default): Takes effect immediately.
* `MaintainTime`: Takes effect during the [maintenance window](https://help.aliyun.com/document_detail/610402.html).
* `ScheduleTime`: Switches at a specified time. The specified time must be at least 12 hours after the current time. The actual switchover time follows the formula: **EffectiveTime = ScheduleTime + SwitchTime**.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `encryption_key` - (Optional, Computed, Available since v1.109.0) The ID of the key used for disk encryption within the same region. Specifying this parameter enables disk encryption (which cannot be disabled after it is enabled) and requires you to also specify `RoleARN`.

You can view the key ID in the Key Management Service console or create a new key. For more information, see [Create a key](https://help.aliyun.com/document_detail/181610.html).

-> **NOTE:**  - For ApsaraDB RDS for MySQL, PostgreSQL, and SQL Server, you do not need to specify this parameter. You only need to provide `RoleARN` to create an instance with disk encryption enabled by using the service-managed key.

-> **NOTE:**  - RAM authorization can be configured to require that disk encryption must be enabled when a RAM user creates an instance. If disk encryption is not enabled, instance creation is denied. The following RAM policy enforces this requirement:
`{"Version":"1","Statement":[{"Effect":"Deny","Action":"rds:CreateDBInstance","Resource":"*","Condition":{"StringEquals":{"rds:DiskEncryptionRequired":"false"}}}]}`

-> **NOTE:** This configuration also affects the CreateOrder API operation used to create instances in the console..


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `engine` - (Required, ForceNew) Database engine. Valid values:
  - `MySQL`
  - `SQLServer`
  - `PostgreSQL`
  - `MariaDB`

By default, all database engines are returned.
* `engine_version` - (Required) The target database version. The valid values for this parameter depend on the value of the `Engine` parameter as follows:  
  - MySQL: **5.5/5.6/5.7/8.0**  
  - SQL Server: **2008r2 (high-performance local disk, discontinued)/08r2_ent_ha (cloud disk edition, discontinued)/2012/2012_ent_ha/2012_std_ha/2012_web/2014_std_ha/2016_ent_ha/2016_std_ha/2016_web/2017_std_ha/2017_ent/2019_std_ha/2019_ent**  
  - PostgreSQL: **10.0/11.0/12.0/13.0/14.0/15.0**  

-> **NOTE:**  In SQL Server instances, `_ent` indicates Enterprise Cluster Edition, `_ent_ha` indicates Enterprise Edition, `_std_ha` indicates Standard Edition, and `_web` indicates Web Edition.  

* `general_group_name` - (Optional, Available since v1.274.0) The name of the group to which the ApsaraDB for MyBase MySQL General-purpose instance belongs.  
* `group_policy` - (Optional, Available since v1.274.0) Custom RDS SQL Server account password policy. The following policies are supported:
  - Set only the maximum password age. After this period, the password must be changed: `{"accountSecurityPolicy": {"MaximumPasswordAge": specify_maximum_age}}`
  - Set only the minimum password age. The password cannot be changed again within this period: `{"accountSecurityPolicy": {"MinimumPasswordAge": specify_minimum_age}}`
  - Set both the maximum and minimum password ages: `{"accountSecurityPolicy": {"MaximumPasswordAge": specify_maximum_age, "MinimumPasswordAge": specify_minimum_age}}`

-> **NOTE:**  The minimum password age (valid range: 0–998) cannot be greater than the maximum password age (valid range: 0–999).

* `io_acceleration_enabled` - (Optional, Available since v1.274.0) Switch for the [Buffer Pool Extension (BPE)](https://help.aliyun.com/document_detail/2527067.html) feature of high-performance cloud disks. Valid values:

  - `1`: Enabled
  - `0`: Disabled

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `max_capacity` - (Optional, Float, Available since v1.274.0) Maximum value in the auto-scaling range for RCU (RDS Capacity Unit) of the instance. Valid values:
  - MySQL: **1–32**
  - SQL Server: **2–16**
  - PostgreSQL: **1–14**

-> **NOTE:**  The value of this parameter must be greater than or equal to `MinCapacity`, and only `integers` are supported.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `min_capacity` - (Optional, Float, Available since v1.274.0) The minimum value of the auto-scaling range for the instance's RCU. Valid values:  
  - MySQL: **0.5–32**  
  - SQL Server: **2–16** (integers only)  
  - PostgreSQL: **0.5–14**  

-> **NOTE:** The value of this parameter must be less than or equal to `MaxCapacity`.  


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `network_type` - (Optional, ForceNew, Available since v1.274.0) The network type of the read-only instance. Valid values:

* `VPC`: Virtual Private Cloud
* `Classic`: Classic network

By default, a VPC instance is created, and you must specify `VPCId` and `VSwitchId`.

-> **NOTE:** The network type of a read-only instance can differ from that of the primary instance.

* `optimized_writes` - (Optional, Computed, Available since v1.245.0) Switch for the MySQL [16K atomic write feature](https://help.aliyun.com/document_detail/2858761.html). Valid values:
  - `optimized`: Enabled
  - `none`: Disabled.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `payment_type` - (Required, Available since v1.274.0) The payment type of the instance. Valid values:
  - `Postpaid`: Pay-as-you-go.
  - `Prepaid`: Subscription (billed monthly or yearly).
  - `Serverless`: Serverless billing type. This type is not supported for MariaDB instances. For more information, see [Introduction to MySQL Serverless instances](https://help.aliyun.com/document_detail/411291.html), [Introduction to SQL Server Serverless instances](https://help.aliyun.com/document_detail/604344.html), and [Introduction to PostgreSQL Serverless instances](https://help.aliyun.com/document_detail/607742.html).

-> **NOTE:**  The system automatically generates an order and completes the payment without requiring manual confirmation.

* `period` - (Optional) Specifies whether a subscription instance is billed yearly or monthly. Valid values:
* `Year`: Yearly subscription.
* `Month`: Monthly subscription.

-> **NOTE:**  This parameter is required if the billing method is `Prepaid`.


-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `pgbouncer_port` - (Optional, Available since v1.274.0) The PgBouncer port number.  

-> **NOTE:**  This parameter applies only to ApsaraDB RDS for PostgreSQL instances. If PgBouncer is enabled, you can modify the PgBouncer port number.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `port` - (Optional, Computed, Available since v1.126.0) The connection port of the new instance.  

-> **NOTE:**  The `DBInstanceNetType` parameter determines whether this port is an internal or public endpoint.  

* `private_ip_address` - (Optional, Computed, Available since v1.125.0) No configuration required. Represents the private IP address of the target instance. The system automatically assigns an IP address based on the specified `VPCId` and `vSwitchId` by default.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `promotion_code` - (Optional, Available since v1.274.0) The coupon code.  

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `resource_group_id` - (Optional, Computed, Available since v1.86.0) Resource group ID. You can obtain it by calling the DescribeDBInstanceAttribute operation.
* `role_arn` - (Optional, Computed, Available since v1.208.0) The Amazon Resource Name (ARN) that grants the RDS service account permission to access KMS on behalf of the Alibaba Cloud account. You can use the CheckCloudResourceAuthorized API operation to view the ARN information.  

-> **NOTE:** `RoleARN` must be specified when disk encryption is enabled.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `security_ip_list` - (Required, ForceNew, Available since v1.274.0) The [IP whitelist](https://help.aliyun.com/document_detail/43185.html) for this instance. Separate multiple entries with commas (`,`), ensure no duplicates, and include up to 1,000 entries. The following formats are supported:
  - IP address format, for example: 10.10.XX.XX.
  - CIDR notation, for example: 10.10.XX.XX/24 (Classless Inter-Domain Routing; 24 indicates the prefix length, which can range from 1 to 32).

If left unspecified, the default value is the IP whitelist of the `default` group from the original instance.
* `serverless_config` - (Optional, ForceNew, Set, Available since v1.200.0) The configuration settings for an RDS Serverless instance. This parameter is required when creating a Serverless instance.  

-> **NOTE:**  Serverless-related configurations apply only to RDS MySQL instances.
 See [`serverless_config`](#serverless_config) below.
* `sql_collector_status` - (Optional, Computed, Available since v1.70.0) Enables or disables SQL Insight (SQL audit). Valid values:
  - `Enable`
  - `Disabled`.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `storage_auto_scale` - (Optional, Available since v1.129.0) The switch for automatic storage scaling. This feature is supported only for MySQL and PostgreSQL instances. Valid values:
* `Enable`: Enable.
* `Disable`: Disable (default).

-> **NOTE:** You can also adjust this setting after instance creation by calling ModifyDasInstanceConfig. For more information, see [Configure automatic storage scaling](https://help.aliyun.com/document_detail/173826.html).


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `storage_threshold` - (Optional, Int, Available since v1.129.0) The threshold (in percentage) that triggers automatic storage scaling. Valid values:
* `10`
* `20`
* `30`
* `40`
* `50`

-> **NOTE:** This parameter is required when `StorageAutoScale` is set to `Enable`.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `storage_upper_bound` - (Optional, Int, Available since v1.129.0) The upper limit of total storage space when automatic storage scaling is enabled. The total storage space of the instance will not exceed this value after automatic scaling. Unit: GB.

-> **NOTE:**  - The value must be greater than or equal to 0.

-> **NOTE:**  - This parameter is required when `StorageAutoScale` is set to `Enable`.


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `switch_time` - (Optional, Available since v1.126.0) The time at which the specification change takes effect. **We recommend performing specification changes during off-peak business hours.**  

Format: yyyy-MM-ddTHH:mm:ssZ (UTC time).  

-> **NOTE:**  - The specified time **must be later than the current time** (that is, the time when the API call is initiated). Otherwise, the specification change task fails. After a failure, you must wait until the order is automatically canceled before initiating another call.  

-> **NOTE:**  - If you only increase storage capacity or change the ESSD storage type, the specification change takes effect immediately after submission because it does not affect your workloads. In this case, you do not need to configure this parameter.  


-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `tags` - (Optional, Map) List of tags.
* `target_minor_version` - (Optional, Computed, Available since v1.126.0) Specifies the minor kernel version of the created RDS instance. This parameter is required only when creating a MySQL or PostgreSQL instance.  
Format:  
  - MySQL: `_`. Examples: `rds_20200229`, `xcluster_20200229`, or `xcluster80_20200229`. Descriptions are as follows:  
    * `rds`: High-availability Edition or Basic Edition.  
    * `xcluster`: MySQL 5.7 Enterprise Edition with three nodes.  
    * `xcluster80`: MySQL 8.0 Enterprise Edition with three nodes.  

    > You can query the numeric version number by calling the DescribeDBMiniEngineVersions operation. For differences between versions, see [AliSQL Minor Version Release Notes](https://help.aliyun.com/document_detail/96060.html).  
  - PostgreSQL: `rds_postgres_00_`. Example: `rds_postgres_1400_20220830`. Descriptions are as follows:  
    * `1400`: Major PostgreSQL version 14.  
    * `20220830`: AliPG minor kernel version, which you can query by calling the DescribeDBMiniEngineVersions operation. For differences between versions, see [PostgreSQL Minor Version Release Notes](https://help.aliyun.com/document_detail/126002.html).  

    > If Babelfish is enabled in `BabelfishConfig`, the minor version format for an RDS PostgreSQL instance is: `rds_postgres_00__babelfish`.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `used_time` - (Optional, Available since v1.274.0) The subscription duration. Valid values:  
* If `Period` is set to `Year`, `UsedTime` can be **1 to 5**.  
* If `Period` is set to `Month`, `UsedTime` can be **1 to 11**.  

-> **NOTE:**  This parameter is required if the payment type is `Prepaid`.  


-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `user_backup_id` - (Optional, Available since v1.274.0) The user backup ID. You can query this ID by calling the ListUserBackupFiles operation. If you specify this parameter, you can create an instance based on the user backup.

If you specify this parameter, the following restrictions apply:  
  - The `PayType` parameter must be `Postpaid`.  
  - The `Engine` parameter must be `MySQL`.  
  - The `EngineVersion` parameter must be **5.7**.  
  - The `Category` parameter must be `Basic`.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `vswitch_id` - (Optional) The ID of the vSwitch. The zone where the vSwitch resides must match the zone ID specified in `ZoneId`.  
  - The network type `InstanceNetworkType` must be `VPC`.  
  - If you specify ZoneSlaveId1 (the secondary zone ID), you must provide two vSwitch IDs separated by a comma (,).  

-> **NOTE:**  This parameter is required only when you upgrade the major version of an SQL Server instance (AllowMajorVersionUpgrade) or when you switch vSwitches.  

* `vpc_id` - (Optional, ForceNew, Computed, Available since v1.185.0) The VPC ID of the read-only instance. This parameter is required when `InstanceNetworkType` is set to empty or `VPC`.

-> **NOTE:**  * If the storage type of the primary instance uses high-performance local disks, the read-only instance can use any VPC.

-> **NOTE:**  * If the storage type of the primary instance uses cloud disks, the VPC of the read-only instance must be the same as that of the primary instance.

* `whitelist_template_list` - (Optional, Available since v1.274.0) The whitelist list.  
When you need to configure multiple IP addresses, separate multiple IP addresses or CIDR blocks with commas (,). Do not add spaces before or after the commas, for example, `192.168.0.1,172.16.213.9`.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `zone_id_slave_one` - (Optional, Available since v1.274.0) The zone ID of the secondary node.
  - If you specify `Auto`, the deployment spans multiple zones, and the system automatically selects a zone for the secondary node.
  - If this value is the same as `ZoneId`, the deployment is in a single zone.
  - If this value differs from `ZoneId`, the deployment spans multiple zones.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `zone_id_slave_two` - (Optional, Available since v1.274.0) Secondary zone 2.  

-> **NOTE:**  You can specify this parameter only for instances of the three-node Enterprise Edition.


-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.


### `babelfish_config`

The babelfish_config supports the following:
* `babelfish_enabled` - (Optional, ForceNew) Babelfish toggle.

-> **NOTE:**  If Babelfish is enabled when you purchase an RDS PostgreSQL instance, this parameter is fixed to `true`.

* `master_user_password` - (Optional, ForceNew) The password for the administrative account.
* `master_username` - (Optional, ForceNew) The initial administrative account.
* `migration_mode` - (Optional, ForceNew) Babelfish migration mode.  
  - **single-db**: Single-database mode.  
  - **multi-db**: Multi-database mode.  

-> **NOTE:**  For more information about Babelfish migration modes, see [Introduction to migration modes](https://help.aliyun.com/document_detail/428613.html).  


### `serverless_config`

The serverless_config supports the following:
* `switch_force` - (Optional) Forces scaling for MySQL Serverless or PostgreSQL Serverless instances. Elastic scaling of instance RCU usually takes effect immediately, but in certain special scenarios (for example, during execution of large transactions), scaling may not complete instantly. In such cases, you can enable this parameter to force scaling. Valid values:
  - `true`: Enables forced scaling.
  - `false` (default): Disables forced scaling.

-> **NOTE:**  After enabling this parameter, the instance will experience **30 to 120 seconds of service unavailability during forced scaling**. Use this option with caution based on your actual situation.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `auto_upgrade_minor_version` - The method for upgrading the minor version of the instance.
* `create_time` - The creation time of the backup data migration task.
* `db_instance_type` - The instance type.
* `maintain_time` - The maintenance window for the instance.
* `region_id` - The region ID.
* `status` - The task status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 15 mins) Used when create the Db Instance.
* `delete` - (Defaults to 15 mins) Used when delete the Db Instance.
* `update` - (Defaults to 15 mins) Used when update the Db Instance.

## Import

RDS Db Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_db_instance.example <db_instance_id>
```