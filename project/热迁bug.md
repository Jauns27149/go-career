# 问题复现

1. 创建虚拟机

   ```bash
   openstack server create \
   --flavor 214d0281-6fb5-48ae-9ca7-64f2a11b5db6 \
   --image 24647239-7b2e-4895-8ede-1d278a3b10df \
   --nic net-id=010c543a-f741-45f5-892c-ad84f710d48b \
   --availability-zone  S6-PUBLIC-ZONE \
   janus
   
   # 环境: 2022年-贵州公共测试-贵州-弹性计算测试环境 55.243.31.30
   openstack server create \
   --flavor 214d0281-6fb5-48ae-9ca7-64f2a11b5db6 \
   --volume 01423f2a-d1c8-4ded-850b-664c8910a9c2 \
   --network  00072616-0c0b-4551-b4a3-01465e2c19a6 \
   --availability-zone  S6-PUBLIC-ZONE \
   --wait \
   janus
   
   openstack server create \
   --flavor 214d0281-6fb5-48ae-9ca7-64f2a11b5db6 \
   --volume 01423f2a-d1c8-4ded-850b-664c8910a9c2 \
   --network  00072616-0c0b-4551-b4a3-01465e2c19a6 \
   --availability-zone  S6-PUBLIC-ZONE \
   --wait \
   janus_confuse_volume
   
   nova volume-attach janus 447522bf-5f8c-41c1-9b10-a629efac9a9c
   ```

   ```bash
   [root@gz-txjs-control-55e243e31e30 ~]# openstack host list --zone S6-PUBLIC-ZONE
   +----------------------------+---------+----------------+
   | Host Name                  | Service | Zone           |
   +----------------------------+---------+----------------+
   | gz-txjs-szj-55e243e16e33   | compute | S6-PUBLIC-ZONE |
   | gz15-txjs-szj-55e243e16e36 | compute | S6-PUBLIC-ZONE |
   +----------------------------+---------+----------------+
   ```

   ![image-20241208095545964](assets/image-20241208095545964.png)

2. 热迁

   ```bash
   nova live-migration c2ba7e24-29b0-4a3f-9289-abed3db29e65 gz15-txjs-szj-55e243e16e78 
   ```

3. 热迁终止

   ```bash
   nova live-migration-abort c2ba7e24-29b0-4a3f-9289-abed3db29e65 74662
   ```

   

4. 重建失败









```bash
[api_database]
connection = CTYUN_SECRET@IyHuWZgPOfEgnZ5trJl7skIPeHQdbTmcQ+J2ekhaPLixjDeMAis5iQ84rMJ0aWk8K4Y98UvBMs8qEkL8yXolcTimf0PfBSy+ev/Py4a13LM=
```

```bash

[root@gz-txjs-control-55e243e31e30 Encrypt_tools]# python decrypt.py --decrypt_data  CTYUN_SECRET@IyHuWZgPOfEgnZ5trJl7skIPeHQdbTmcQ+J2ekhaPLixjDeMAis5iQ84rMJ0aWk8K4Y98UvBMs8qEkL8yXolcTimf0PfBSy+ev/Py4a13LM=
mysql+pymysql://nova:S43NGltg0817@mariadb.cty.os:10024/nova_api
[root@gz-txjs-control-55e243e31e30 Encrypt_tools]#


python /root/zyc/Encrypt_tools/decrypt.py --decrypt_data "CTYUN_SECRET@IyHuWZgPOfEgnZ5trJl7skIPeHQdbTmcQ+J2ekhaPLixjDeMAis5iQ84rMJ0aWk8K4Y98UvBMs8qEkL8yXolcTimf0PfBSy+ev/Py4a13LM="


echo "mysql+pymysql://nova:S43NGltg0817@mariadb.cty.os:10024/nova_api" | cut -d'/' -f3 | cut -d'@' -f1 | cut -d':' -f2
```

```bash
[root@gz-txjs-control-55e243e31e30 etc]# cat hosts |grep db
55.243.31.11 mariadb.cty.os
[root@gz-txjs-control-55e243e31e30 etc]# mysql -unova -pS43NGltg0817 -P 10024 -h 55.243.31.11


nova ->  block_device_mapping
```



```sql
select volume_id, instance_uuid
	from block_device_mapping 
	where deleted=0
	and volume_id in ( select volume_id from block_device_mapping where deleted=0 group by volume_id having count(*)>1 )
    and boot_index=0;
	
select count(*) from block_device_mapping group by volume_id having count(*)>1

1352852d-41e1-480e-87fa-d30be21ad382
select volume_id, instance_uuid from block_device_mapping where volume_id='1352852d-41e1-480e-87fa-d30be21ad382';

select volume_id from block_device_mapping where deleted=0 group by count(*)>1 ;


```

