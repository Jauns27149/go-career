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

   ![image-20241208095545964](../assets/image-20241208095545964.png)

2. 热迁

   ```bash
   nova live-migration c2ba7e24-29b0-4a3f-9289-abed3db29e65 gz15-txjs-szj-55e243e16e78 
   ```

3. 热迁终止

   ```bash
   nova live-migration-abort c2ba7e24-29b0-4a3f-9289-abed3db29e65 74662
   ```

   

4. 重建失败























