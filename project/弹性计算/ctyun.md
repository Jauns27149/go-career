

# 云计算服务模型

- **IaaS（Infrastructure as a Service）**：提供基础计算资源，用户管理操作系统及之上的一切。
- **PaaS（Platform as a Service）**：提供开发和部署平台，用户管理应用程序和数据。
- **SaaS（Software as a Service）**：提供完整应用程序，用户直接使用应用程序。
- **FaaS（Functions as a Service）**：提供无服务器计算服务，用户上传代码片段按需执行。
- **DaaS（Data as a Service）**：提供数据管理和分析服务，用户访问和使用数据。
- **CaaS（Containers as a Service）**：提供容器管理服务，用户管理容器化的应用程序。

# openStack

​	OpenStack是一个云操作系统，通过数据中心可控制大型的计算、存储、网络等资源池。所有的管理通过前端界面管理员就可以完成，同样也可以通过web接口让最终用户部署资源。

- Nova : 计算服务，负责管理虚拟机的生命周期
  - nova-api：处理 API 请求
  - nova-compute：管理虚拟机的生命周期
  - nova-scheduler：选择最适合运行虚拟机的主机
  - nova-conductor：处理数据库操作。

- Neutron: 网络服务，负责管理虚拟网络
  - neutron-server：处理 API 请求
  - neutron-linuxbridge-agent：管理 Linux 桥接网络
  - neutron-dhcp-agent：提供 DHCP 服务
  - neutron-l3-agent：管理路由和 NAT
  - neutron-metadata-agent：提供元数据服务

- Glance : 镜像服务，存储和检索虚拟机镜像
  - glance-api：处理 API 请求
  - glance-registry：管理镜像元数据

- Cinder : 块存储服务，提供持久性存储卷
  - cinder-api：处理 API 请求
  - cinder-volume：管理块存储卷
  - cinder-scheduler：选择最适合创建卷的主机

- Keystone :  身份认证服务，管理用户和权限
- Swift(对象存储服务): 提供了一个可扩展的、冗余的对象存储系统，用于存储非结构化的数据，如图片或视频文件。
- Heat(编排服务): 一个编排引擎，允许用户定义和管理复杂的云应用，通过模板描述应用的资源需求和依赖关系，Heat 自动化资源的创建和配置过程。
- Trove(数据库服务): 提供了一个数据库即服务的解决方案，使用户能够轻松地管理和部署关系型数据库。
- Horizon：Web 控制面板，提供图形界面

## nova

- 概念

1. Nova和Swift是OpenStack最早的两个组件，nova分为控制节点和计算节点。
2. 计算节点通过Nova Compute进行虚拟机创建，通过libvirt调用kvm创建虚拟机，nova之间通信通过rabbitMQ队列进行通信。
3. Nova位于Openstack架构的中心，其他服务或者组件（比如Glance、Cinder、Neutron等）对它提供支持，另外它本身的架构也比较复杂。

- 作用

1. Nova是OpenStack最核心的服务模块，负责管理和维护云计算环境的计算资源，负责整个云环境虚拟机生命周期的管理。
2. Nova是OpenStack的计算服务，负责维护和管理的网络和存储，提供计算服务。

- 组件

  ![stickPicture](assets/stickPicture-1731554014520-1.png)

![stickPicture](assets/stickPicture.png)

### 工作流程

![stickPicture](assets/stickPicture-1731403108838-10.png)

![stickPicture](assets/stickPicture-1731403218807-15.png)

![stickPicture](assets/stickPicture-1731403224722-17.png)

![stickPicture](assets/stickPicture-1731403231405-19.png)

## 虚拟机创建流程

![https://s2.51cto.com/wyfs02/M01/87/01/wKiom1fRA6PRYz_7AAQI6mXaRn4200.png](https://s2.51cto.com/wyfs02/M01/87/01/wKiom1fRA6PRYz_7AAQI6mXaRn4200.png)

1. 界面或命令行通过RESTful API向keystone获取认证信息。
2. keystone通过用户请求认证信息，并生成auth-token返回给对应的认证请求。
3. 界面或命令行通过RESTful API向nova-api发送一个boot instance的请求（携带auth-token）。
4. nova-api接受请求后向keystone发送认证请求，查看token是否为有效用户和token。
5. keystone验证token是否有效，如有效则返回有效的认证和对应的角色（注：有些操作需要有角色权限才能操作）。
6. 通过认证后nova-api和数据库通讯。
7. 初始化新建虚拟机的数据库记录。
8. nova-api通过rpc.call向nova-scheduler请求是否有创建虚拟机的资源(Host ID)。
9. nova-scheduler进程侦听消息队列，获取nova-api的请求。
10. nova-scheduler通过查询nova数据库中计算资源的情况，并通过调度算法计算符合虚拟机创建需要的主机。
11. 对于有符合虚拟机创建的主机，nova-scheduler更新数据库中虚拟机对应的物理主机信息。
12. nova-scheduler通过rpc.cast向nova-compute发送对应的创建虚拟机请求的消息。
13. nova-compute会从对应的消息队列中获取创建虚拟机请求的消息。
14. nova-compute通过rpc.call向nova-conductor请求获取虚拟机消息。（Flavor）
15. nova-conductor从消息队队列中拿到nova-compute请求消息。
16. nova-conductor根据消息查询虚拟机对应的信息。
17. nova-conductor从数据库中获得虚拟机对应信息。
18. nova-conductor把虚拟机信息通过消息的方式发送到消息队列中。
19. nova-compute从对应的消息队列中获取虚拟机信息消息。
20. nova-compute通过keystone的RESTfull API拿到认证的token，并通过HTTP请求glance-api获取创建虚拟机所需要镜像。
21. glance-api向keystone认证token是否有效，并返回验证结果。
22. token验证通过，nova-compute获得虚拟机镜像信息(URL)。
23. nova-compute通过keystone的RESTfull API拿到认证k的token，并通过HTTP请求neutron-server获取创建虚拟机所需要的网络信息。
24. neutron-server向keystone认证token是否有效，并返回验证结果。
25. token验证通过，nova-compute获得虚拟机网络信息。
26. nova-compute通过keystone的RESTfull API拿到认证的token，并通过HTTP请求cinder-api获取创建虚拟机所需要的持久化存储信息。
27. cinder-api向keystone认证token是否有效，并返回验证结果。
28. token验证通过，nova-compute获得虚拟机持久化存储信息。
29. nova-compute根据instance的信息调用配置的虚拟化驱动来创建虚拟机。

# Libvirt

​	Libvirt是一个开源项目，提供了一组API、工具、库，用于管理和控制虚拟化平台。
​	在Openstack环境中，Libvirt是一个至关重要的组件，它为各种虚拟化技术（如 KVM、QUME、Xen和LXC）提供统一的接口，使得Openstack能够和底层虚拟化技术进行交互。

- Libvirt 主要功能包括：
  1. API提供：Libvirt 提供一个C语言的API，同时也支持多种高级编程语言的绑定。这些API允许开发者编写应用程序来创建、配置和管理虚拟机。
  2. 虚拟化管理接口：Libvirt 提供了一个统一的接口，可以透明地处理不同的虚拟化技术。这意味着Openstack不需要知道具体的虚拟化实现，而是通过libvirt进行操作，简化了开发和维护工作。
  3. 安全隔离：Libvirt 支持安全策略，确保各个虚拟机之间的隔离，提高系统的安全性。
  4. 资源管理：Libvirt 可以控制和调整虚拟机的资源分配，包括CPU、内存、磁盘和网络。这对于优化虚拟化环境中的资源利用率至关重要。
  5. 网络管理：Libvirt 提供了网络抽象层，能够创建和配置网络桥联、网络过滤器等，支持虚拟网络设备的管理。
  6. 存储管理：Libvirt 支持多种存储类型，如块设备、文件系统、网络存储，以及Openstack中的Cinder存储服务。

- Openstack中，Libvirt 主要与以下服务交互：
  1. nova：作为Openstack计算服务，nova 通过调用 Libvirt 的API来执行这些操作，包括创建、启动、停止和迁移虚拟机实例。
  2. neutron：Openstack网络服务 neutron 可以利用Libvirt 来配置虚拟网络，如设置网络连接、端口安全规则和负载均衡。
  3. cinder：cinder 直接与后端存储系统交互，但 Libvirt 参与了卷的挂载和卸载，以及在虚拟机内部使用的cinder卷。

 	Libvirt 还包含了一些命令行工具，如virsh 等，允许管理员直接对虚拟机进行操作，如查看状态、编辑配置、挂载磁盘等。Libvirt 是 Openstack 架构中的关键组件，它作为中间层连接上层服务和底层虚拟化技术，提供高效、灵活和安全的虚拟化管理能力。

## virsh 命令

| 命令                           | 作用                  |
| ------------------------------ | --------------------- |
| virsh list                     | 列出所有虚拟机        |
| virsh start <domain-name>      | 启动虚拟机            |
| virsh shutdown <domain-name>   | 关闭虚拟机            |
| virsh destroy <domain-name>    | 强制关闭虚拟机        |
| virsh reboot <domain-name>     | 重启虚拟机            |
| virsh suspend <domain-name>    | 暂停虚拟机            |
| virsh resume <domain-name>     | 恢复虚拟机            |
| virsh dumpxml <domain-name>    | 查看虚拟机详细信息    |
| virsh define <xml-file>        | 创建虚拟机            |
| virsh undefine <domain-name>   | 删除虚拟机            |
| virsh domstate <domain-name>   | 查看虚拟机状态        |
| virsh vncdisplay <domain-name> | 查看虚拟机的 VNC 端口 |

# QGA（Qemu Guest Agent）

- 定义与作用：
  1. QGA是一个运行在虚拟机内部的普通应用程序（可执行文件名称默认为qemu-ga，服务名称默认为qemu-guest-agent）。
  2. 其主要目的是实现宿主机和虚拟机之间的一种不依赖于网络的交互方式，而是依赖于virtio-serial（默认首选方式）或者isa-serial。
  3. QGA通过读写串口设备与宿主机上的socket通道进行交互，交互的协议与QMP（QEMU Monitor Protocol）相同，即使用JSON格式进行数据交换。

- 功能特点：
  1. QGA提供了虚拟机内部状态信息（如文件系统信息、网络信息等）的查询和修改能力。
  2. 它可以执行一些宿主机发起的操作，如文件操作、磁盘管理、网络配置等。
  3. QGA的功能扩展较为方便，开发者可以通过修改源码来添加新的命令或功能。

#  openstack 命令

- openstack [commands]
  - image list : 查看镜像
  
  - flavor list : 查看规格
  
  - network list : 查看network
  
  - availability zone list : 查看可用区
  
  - server 
    - list : 查看虚拟机
  
      ```bash
      Status:
      	BUILD: 实例正在创建过程中。
      	ACTIVE: 实例已经成功创建并运行。这是实例正常运行的状态。
      	SHUTOFF: 实例已被关闭，但未被删除。这个状态下的实例不会消耗资源，但可以通过启动操作重新变为ACTIVE状态。
      	ERROR: 实例在创建或运行过程中遇到了错误，无法继续。这个状态通常需要管理员进行干预来解决问题。
      	DELETED: 实例已经被删除，不再存在。
      	RESIZE: 实例正在进行大小调整（例如更改CPU或内存配置）。
      	VERIFY_RESIZE: 实例已经完成了大小调整，但尚未确认新的大小是否生效。
      	PAUSED: 实例已暂停，可以暂时停止其运行。
      	SUSPENDED: 实例已挂起，类似于暂停，但可能涉及更深层次的资源管理。
      	SHELVED: 实例已被“归档”，即保存到磁盘上，以便以后恢复。
      	SHELVED_OFFLOADED: 实例已被完全从主机上移除，存储在外部存储设备上。
      ```
  
    - create
      - --image :
  
      - --flavor : 指定规格
  
      - --availability-zone :
  
      - --nic <net-id=net-uuid,v4-fixed-ip=ip-addr,v6-fixed-ip=ip-addr,port-id=port-uuid,auto,none> ：
  
      - <server-name>
  
        ```bash
         openstack server create --image centOS_6.8  --flavor s2.2xlarge.2   --availability-zone  S6-PUBLIC-ZONE --nic net-id=08048318-af62-4a15-9634-e5db4d3a015f janus-test
         # 2022年-贵州公共测试-贵州-弹性计算测试环境
        ```
    
    - delete <instance-id> : 删除虚拟机实例
    
    - show <instance-id> : 虚机详情
    
    - stop <instance-id> : 虚机关机
    
    - start <instance-id> : 虚机开机
    
    - rebuld <instance-id> --image <image-id> : 重装虚机
    
    - resize <instance-id> 
    
      - --flavor <flavor name or id> : 更换规格（需要再次确认操作）
      - --confirm <instance-id> : 确认更改
      - --revert <instance-id> : 取消变更
    
    - migrate 
    
      - --live <target-host> <instance-id> : 热迁手打指定目标节点（虚机状态应当为ACTIVE or RUNNING）
      - --host<target-hsot> <instance-id> : 冷迁（虚机状态为SHUTOFF）
    
  - hypervisor stats show : 查看计算节点的资源利用率
  
  - compute service list : 查看节点
  
    

模拟环境：2022年-贵州多AZ测试环境-POC2 ->  

2022年-贵州多AZ测试环境-POC2 -> 55.249.31.26



systemctl status openstack-nova-api 

systemctl status openstack-nova-scheduler 

systemctl status openstack-nova-conductor

# gs

内蒙08 ->  10.8.73.43
配置环境  . admin-openrc az1

## 命令

- gs : `gs [command]`

  - instance : 虚拟机操作 

    - boot : 创建虚拟机

      - --flavor : 规格

      - --image : 镜像

      - --host <hsot-id>： 宿主机

      - --zone :

      - --name : 虚拟机名称

      - --port ： 虚拟机映射端口

      - --volume <int> : 创建卷大小 ，单位G

        ```bash
        gs instance boot \
        --flavor s7.small.1 \
        --image 65639384-7cc3-4d9d-bf44-c9161ebd3d26  \
        --host 23d752acc3284677ac00e5320aa8633adfd25ca5 \
        --zone SERIES-7-ZONE  \
        --name janus \
        --port port-ustdivjyb0 \
        --volume-size 536870(单位？)
        ```

        

    - password : 更改密码

      - --password <string> 
      - --username <string>

    - list : 虚拟机列表

      - --name <string> ： 虚拟机名称

    - delete <instance-id> : 删除虚拟机（软删除）

      - --hard ：硬删除
      - --ignore : 忽略虚拟机ERROR状态（ERROR状态默认不能删除）

    - show <instance-id> : 虚拟详情

    - stop <instance-id> : 关机

    - start <instance-id> : 开机

    - rebuild : `gs instance  <instance_id> rebuild [flags]`, 重建虚拟机

      - --image <image-id> : 镜像 

    - detach-volume <instance-id> [flag] : 卸载磁盘

      -  --volume-id <string>

    -  attach-volume <isntance-id> :

  - host : 

    - show <host-id> : Host详情

- virsh
  - list : 虚拟机列表
    - --limit <int> : 限制返回条目数量
  - show <volume-id> : 卷详情
  - delete <volume-id> : 删除卷
  - create <size> :  新建卷



- cinder
  - 



```bash
ce0a144b-82b8-4eca-9bdb-2fe3c6fe46d8卷

gs instance boot \
--flavor s7.small.1 \
--image 65639384-7cc3-4d9d-bf44-c9161ebd3d26  \
--host 23d752acc3284677ac00e5320aa8633adfd25ca5 \
--zone SERIES-7-ZONE  \
--name janus \
--port port-ustdivjyb0 \
--volume-size 53,687,091,200  536870
--network vpc-t15x76ve03 

17788b41-8910-5ac7-0327-3d807bc90911 intance-id 
001a18cadd4b401e9fdeab6c411d9816  project
192.168.1.3

10.8.75.47 计算节点
 4e1f3a23-d9d3-46f2-8b92-69a859c79358 | Ubuntu-18.04-amd64-20220903  

06616f84-a543-46db-950b-4f647a417889 例子实例


206eb1d075cc2e7e7e75cfdbdd02f6db36764e44
964cc409638f68259d7f47b85f3d9818604e6323
c2b9a3ee68bda58f76901a8f6d291c0010326a11
c418b33b3b60dfeef5541d5dc224a792be3e4afc
07e9c3307196aa171edd42f8528e9702e7273f07
73f3a807f9b7584fdbeb2b80634da116ca449202
6580818821e180e5bb5b7ae5146a654c84aa012d
42da0c8abe25298c5468a136a0ea1c81de0e2c05
17b2538eb5028ea9775856101784a7b211af469b
4bb62dedca0a5b14f202f9d1a8f3837d51b65ef8
799743ff750a2f86e24a8c73e1273fca0789927c
fe2fa8b90a4b2a1d9d9b1eae169b0fce2cd716bb
fba7e74f6902347d36c559d05b7fdb92fd9a02bf
9a74917beecfd360b4e10067e028bfb78e22d7f5
e940ff6070d097adde89aab6615128493df3037e
a72f954dd0ef45bf8486be5a8be9ba7a40ed10b3
eca3ff03b465dfb488ef4a7fcf45db1da6473053
09f59e975239ff18ce3b6c3f882e465c518a235d
9fe7e23e6477d79e4daa6d33317e3b085562a648
9ae1e25cd94c93dc1d73fd305359180d3b95724b
8d0e4695605610ceb7e60ea5f301aa118ea02d79
89920d7f7a69c62c3d467e6734f5d290964c6d3c
a6c7c9897928a7773b95f3fa444765eaf84f1d3b
e502500cb576102fab53f7eb0447abba421b3134
1b0680e94d636f8bd3cf97f072a806a547041288
3d66eeef554d804ae301c037c863f18ef5ce8227
0a335591afc11b39b1e68e8b5c429e41a3d4fe95  1
b0993e183676d323bea9e1c5a1933a536701ed39
b703c9a17e7c57ad9516aead2bf6d4e896440eec
0450d7c99cd4b762115e11ca62b73354899c7f7e
a9e296b7c3c2fda4d9b8f92c58e7fd3e2ed6cf6d
c891d41058eddc63b66f9f59d4373ad751d29a7b


[root@cn-nm-region1-az1-control-10e8e73e43 ~]# gs instance list --help
list all instances

Usage:
  gs instance list [flags]

Examples:
  gs instance list [--status <status>] [--host-id <host_id>] [--host-name <hostname>]
 [--zone <zone>] [--inner-status <inner-status>] [--long]

Flags:
      --create-date int64Slice   create date of instances (default [])
      --debug                    debug
      --fields string            show fields separated by comma
      --flavor-ids strings       flavor-ids of instances
  -h, --help                     help for list
      --host-id string           host-id of instances
      --host-name string         host-name of instances
      --image-id string          image id of instances
      --inner-status             list instance inner  status(etc. --inner-status=false)
      --ip string                ip of instances
      --key-name string          key-name of instances
  -l, --limit int                limit of result
      --long                     list more message
      --marker string            marker of instances
      --name string              name of instances
      --project-id string        project id of instances
  -s, --skip int                 skip of result
      --status string            search instances by comma-separated list of status
      --update-date int64Slice   update date of instances (default [])
      --user-id string           user id of instances
      --zone string              zone of instances

```





- gs [commad]
  - instance 
    - list : 虚机列表
    - show <instance-id> : 查看虚机详细信息
    - delete <instance-id> : 删除虚机，默认软删除
      - --hard : 硬删除
      - --ignore : 忽略 error





```bash
[root@cn-nm-region1-az1-control-10e8e73e43 ~]# vnetops --help
command line interface for vnet service

Usage:
  vnetops [OPTIONS] COMMAND [args] [flags]
  vnetops [command]

Available Commands:
  api         command line interface for vnet-api service
  cbm         command line interface for cbm service
  cbmmgr      command line interface for cbm manager
  completion  Generate the autocompletion script for the specified shell
  elb         command line interface for elb service
  elbgw       A self-sufficient runtime for elb gw controller system
  elbmgr      command line interface for elb manager
  epos        command line interface for epos service
  hc          command line interface for hc service
  hcmgr       command line interface for hc manager
  help        Help about any command
  hg          command line interface for vpc service
  ibc         command line interface for ib service
  mc          command line interface for mc service
  mcmgr       command line interface for MC MANAGERv
  migration   command line interface for fast migration
  netprobe    Command line interface for netprobe server
  nfv         command line interface for nfv service
  nfvmgr      command line interface for nfv manager
  ns          command line interface for notify service
  proxy       A self-sufficient runtime for vnet proxy system
  qos         command line interface for Qos init
  roce        command line interface for roce service
  rocemgr     command line interface for roce manager
  scan        command line interface for vnet-scan service
  sdv2        command line interface for service discovery
  smoke       command line interface for smoke-online service
  vpc         command line interface for vpc service
  vpcc        A self-sufficient runtime for vpc controller system
  vpcgw       A self-sufficient runtime for vpc gw controller system
  vpcmgr      command line interface for vpc manager

Flags:




[root@cn-nm-region1-az1-control-10e8e73e43 ~]# vpcs --help
command line interface for vpc service

Usage:
  vpcs COMMAND [args] [flags]
  vpcs [command]

Available Commands:
  acl-create                                 Create a network ACL(Access Control List) for a given vpc.
  acl-delete                                 Delete a network ACL(Access Control List).
  acl-list                                   List network ACL(Access Control List)s.
  acl-modify                                 Modify attributes of a network ACL(Access Control List).
  acl-rule-create                            Create an ACL rule for a given ACL.
  acl-rule-delete                            Delete network ACL rules.
  acl-rule-list                              List network ACL rules.
  acl-rule-modify                            Modify attributes of a network ACL(Access Control List) rule.
  acl-show                                   Display network ACL(Access Control List) details.
  appgw-host-config-list                     List appgw host config
  appgw-host-config-show                     Show appgw host config
  appgw-rs-create                            Create an APPGW RS
  appgw-rs-delete                            Delete an APPGW RS
  appgw-rs-list                              List APPGW RSes
  appgw-rs-modify                            Modify APPGW RS attributes
  appgw-rs-show                              Show APPGW RS
  appgw-rule-add-binding                     Binding APPGW rule to hostgroup.
  appgw-rule-binding-list                    List APPGW rule hostgroup bindings.
  appgw-rule-create                          Create an APPGW rule
  appgw-rule-delete                          Delete an APPGW rule
  appgw-rule-list                            List APPGW rules
  appgw-rule-modify                          Modify APPGW rule attributes
  appgw-rule-modify-binding-role             Modify role of APPGW rule hostgroup binding.
  appgw-rule-remove-binding                  Unbinding APPGW rule from hostgroup.
  appgw-rule-show                            Show APPGW rule
  attached-ip-list                           List attached ips.
  az-create                                  Create a availability zone
  az-delete                                  Delete a Availability Zone
  az-list                                    list Availability Zones
  az-modify                                  Modify a Availability Zone
  az-modify-admin-status                     modify availability zone admin status
  az-show                                    show a Availability Zone
  backend-create                             Create Backend
  backend-delete                             Delete Backend
  backend-group-create                       Create Backend Group
  backend-group-delete                       Delete Backend Group
  backend-group-list                         List Backend Group
  backend-group-modify                       Modify Backend Group Attributes
  backend-group-show                         Show Backend Group
  backend-list                               List Backends
  backend-modify                             Modify Backend Attributes
  backend-modify-health-check-status         Modify health check status of a backend
  backend-show                               Show Backend
  bandwidth-add-binding                      Binding bandwidth to hostgroup.
  bandwidth-add-eip                          Add eips into bandwidth.
  bandwidth-add-ipv6                         Add ipv6 addresses into bandwidth.
  bandwidth-binding-list                     List bandwidth hostgroup bindings.
  bandwidth-create                           Create a bandwidth.
  bandwidth-delete                           Delete a bandwidth.
  bandwidth-list                             List bandwidths.
  bandwidth-modify                           Modify attributes of a bandwidth.
  bandwidth-modify-admin-status              Modify bandwidth admin status
  bandwidth-modify-binding-role              Modify role of bandwidth hostgroup binding.
  bandwidth-remove-binding                   Unbinding bandwidth from hostgroup.
  bandwidth-remove-eip                       Remove eips from bandwidth.
  bandwidth-remove-ipv6                      Remove ipv6 addresses from bandwidth.
  bandwidth-show                             Display bandwidth details.
  dhcp-options-set-create                    Create a dhcp options set.
  dhcp-options-set-delete                    Delete a given dhcp options set.
  dhcp-options-set-list                      List DHCP_OPTIONS_SETs.
  dhcp-options-set-modify                    Update given dhcp options set.
  dhcp-options-set-show                      Show dhcp options set.
  dns-host-config-associate-eip              DNS host config associate eip
  dns-host-config-disassociate-eip           DNS host config disassociate eip
  dns-host-config-list                       List dns host config
  dns-host-config-modify                     Modify dns host config
  dns-host-config-show                       Show dns host config
  dnsgw-add-hostgroup-binding                Add dnsgw hostgroup binding
  dnsgw-create                               Create dnsgw for vpc
  dnsgw-delete                               Delete dnsgw
  dnsgw-list                                 List dnsgws
  dnsgw-remove-hostgroup-binding             Remove dnsgw hostgroup binding
  dnsgw-show                                 Show dnsgw detail
  ecmp-group-add-members                     Add member list for a given EcmpGroup.
  ecmp-group-create                          Create a ecmp group.
  ecmp-group-delete                          Delete a ecmp group.
  ecmp-group-list                            List ecmp groups.
  ecmp-group-member-list                     List ecmp group members.
  ecmp-group-modify                          Modify attributes of a ecmp group.
  ecmp-group-remove-members                  Remove member list for a given EcmpGroup.
  ecmp-group-show                            Display ecmp group details.
  ecmp-group-update-members                  Update member list for a given EcmpGroup.
  eip-associate                              Associate a eip
  eip-check-address                          check eip address can be used
  eip-create                                 Create a eip
  eip-delete                                 Delete a eip
  eip-disassociate                           Disassociate a eip
  eip-list                                   list eips
  eip-modify-admin-status                    Modify eip admin status
  eip-modify-attributes                      Modify a eip
  eip-route-info-list                        list eips' route info
  eip-show                                   Show eip
  endpoint-binding-list                      List Endpoint and APPGW Rule bindings
  endpoint-create                            Create Endpoint
  endpoint-delete                            Delete Endpoint
  endpoint-list                              List Endpoint
  endpoint-modify                            Modify Endpoint Attributes
  endpoint-modify-ip-version                 Modify Endpoint IpVersion
  endpoint-replace-hostgroup-binding         Replace Endpoint Hostgroup
  endpoint-service-create                    Create Endpoint Service
  endpoint-service-delete                    Delete Endpoint Service
  endpoint-service-list                      List Endpoint Service
  endpoint-service-modify                    Modify Endpoint Service Attributes
  endpoint-service-replace-hostgroup-binding Replace Endpoint Service Hostgroup
  endpoint-service-reverse-rule-create       Create Endpoint Service Reverse Rule
  endpoint-service-reverse-rule-delete       Delete Endpoint Service Reverse Rule
  endpoint-service-reverse-rule-list         List Endpoint Service Reverse Rule
  endpoint-service-rule-create               Create Endpoint Service Rule
  endpoint-service-rule-delete               Delete Endpoint Service Rule
  endpoint-service-rule-list                 List Endpoint Service Rule
  endpoint-service-show                      Show Endpoint Service
  endpoint-service-transit-ip-create         Create Endpoint Service Transit IP
  endpoint-service-transit-ip-delete         Delete Endpoint Service Transit IP
  endpoint-show                              Show Endpoint
  eni-add-ecmp-group                         Add ecmpgroup for a given ENI.
  eni-ecmp-group-list                        List eni ecmp group bindings.
  eni-remove-ecmp-group                      Remove ecmpgroup for a given ENI.
  eni-update-ecmp-group                      Update ecmpgroup for a given ENI.
  faasgw-add-hostgroup-binding               Add faasgw hostgroup binding
  faasgw-create                              Create faasgw for vpc
  faasgw-delete                              Delete faasgw
  faasgw-list                                List faasgws
  faasgw-modify                              Modify faasgw attributes
  faasgw-modify-hostgroup-binding-role       Remove faasgw hostgroup binding
  faasgw-remove-hostgroup-binding            Remove faasgw hostgroup binding
  faasgw-rule-create                         Create faasgwRule for vpc
  faasgw-rule-delete                         Delete faasgw rule
  faasgw-rule-expand                         Expand faasgwRule port ip
  faasgw-rule-list                           List faasgw rules
  faasgw-rule-modify                         Modify faasgwRule attributes
  faasgw-rule-show                           Show faasgw rule detail
  faasgw-rule-shrink                         Shrink faasgwRule port ip
  faasgw-show                                Show faasgw detail
  global-route-rule-create                   Create a global route rule.
  global-route-rule-delete                   Delete global route rules.
  global-route-rule-list                     List global route rules.
  ha-vip-associate-port                      Modify attributes of a haVip.
  ha-vip-association-list                    list the binding of havip and port.
  ha-vip-create                              Create a haVip
  ha-vip-delete                              Delete a haVip.
  ha-vip-disassociate-port                   Modify attributes of a haVip.
  ha-vip-list                                List haVips.
  ha-vip-modify                              Modify attributes of a haVip.
  ha-vip-modify-association-role             Modify attributes role of a association.
  ha-vip-show                                Display haVip details.
  host-add-into-hostgroup                    Add host into hostgroup
  host-create                                Create a host
  host-delete                                Delete a host
  host-list                                  List hosts
  host-modify                                Modify host attributes
  host-modify-admin-status                   Modify host admin status
  host-modify-status                         Modify host status
  host-remove-from-hostgroup                 Remove host from hostgroup
  host-show                                  Show host
  hostgroup-create                           Create a hostgroup
  hostgroup-delete                           Delete a hostgroup
  hostgroup-list                             List hostgroups
  hostgroup-modify                           Modify hostgroup attributes
  hostgroup-modify-admin-status              Modify hostgroup admin status
  hostgroup-show                             Show hostgroup
  igw-add-hostgroup-binding                  Add igw hostgroup binding
  igw-add-route-table-binding                add the route table id for igw.
  igw-binding-list                           List igw bindings
  igw-create                                 Create igw for vpc
  igw-delete                                 Delete igw
  igw-list                                   List igws
  igw-modify                                 Modify igw attributes
  igw-modify-hostgroup-binding-role          Remove igw hostgroup binding
  igw-remove-hostgroup-binding               Remove igw hostgroup binding
  igw-remove-route-table-binding             remove the route table id for igw.
  igw-show                                   Show igw detail
  igw6-add-hostgroup-binding                 Add igw6 hostgroup binding
  igw6-add-route-table-binding               add the route table id for igw.
  igw6-binding-list                          List igw6 bindings
  igw6-create                                Create igw6 for vpc
  igw6-delete                                Delete igw6
  igw6-list                                  List igw6s
  igw6-modify                                Modify igw6 attributes
  igw6-modify-hostgroup-binding-role         Remove igw6 hostgroup binding
  igw6-remove-hostgroup-binding              Remove igw6 hostgroup binding
  igw6-remove-route-table-binding            remove the route table id for igw6.
  igw6-show                                  Show igw6 detail
  increase-version                           Increase Resources' version
  internal-port-create                       create internal port
  internal-port-delete                       delet internal port id
  ip-filing-rule-create                      Create filing rules for given ip.
  ip-filing-rule-delete                      Delete ip filing rules for given ip address.
  ip-filing-rule-list                        List ip filing rules.
  ip-filing-rule-show                        Show ip filing rule.
  ipv4-segment-add-agw-hostgroup-binding     Bind the agw hostgroup and ipv4 segment.
  ipv4-segment-list                          List Ipv4 Segments.
  ipv4-segment-list-agw-hostgroup-binding    List the binding of Ipv4 Segments and Agw hostgroup.
  ipv4-segment-modify                        Update given segment.
  ipv4-segment-modify-admin-status           Update the admin status of a ipv4 segment.
  ipv4-segment-remove-agw-hostgroup-binding  Remove the binding of the agw hostgroup and ipv4 segment.
  ipv4-segment-show                          Show Ipv4 Segment.
  ipv6-bandwidth-create                      Create bandwidth for ipv6.
  ipv6-bandwidth-delete                      Delete bandwidth for ipv6.
  ipv6-bandwidth-modify-qos                  Modify qos of bandwidth for ipv6.
  ipv6-block-allocations-list                list ipv6 block allocations
  ipv6-list                                  List ipv6s.
  ipv6-modify                                Modify ipv6's attributes
  ipv6-modify-admin-status                   Modify ipv6 admin status
  ipv6-pre-allocated-block-allocations-clear clean up pre allocated ipv6 block allocations
  ipv6-segment-add-agw-hostgroup-binding     ipv6 segment add agw hostgroup binding
  ipv6-segment-list                          List Segments.
  ipv6-segment-list-agw-hostgroup-binding    List the binding of Ipv6 Segments and Agw hostgroup.
  ipv6-segment-modify                        Update given segment.
  ipv6-segment-modify-admin-status           Update the admin status of a segment .
  ipv6-segment-remove-agw-hostgroup-binding  ipv6 segment remove agw hostgroup binding
  ipv6-segment-show                          Show Segment .
  ipv6-show                                  Display ipv6 details.
  l2-connection-create                       Create l2 connection for l2gw
  l2-connection-delete                       Delete l2 connection
  l2-connection-list                         List l2 connection
  l2-connection-modify                       Modify l2gw connection attributes
  l2-connection-show                         Show l2 connection detail
  l2gw-add-hostgroup-binding                 Add l2gw hostgroup binding
  l2gw-binding-list                          List l2gw bindings
  l2gw-create                                Create l2gw for vpc
  l2gw-delete                                Delete l2gw for vpc
  l2gw-list                                  List l2gws
  l2gw-modify                                Modify l2gw attributes
  l2gw-modify-hostgroup-binding-role         modify l2gw hostgroup binding cmd
  l2gw-remove-hostgroup-binding              Remove l2gw hostgroup binding
  l2gw-show                                  Show l2gw detail
  link-local-port-binding-add                Add link local port binding
  link-local-port-binding-list               List link local port bindings.
  link-local-port-binding-modify-role        Modify link local port binding role
  link-local-port-binding-remove             Add link local port binding
  link-local-port-create                     Create Link Local Port for tenant
  link-local-port-delete                     Delete a link local port
  link-local-port-list                       List link local ports.
  link-local-port-show                       Show link local port.
  list-actions                               List actions service support.
  modify-relay-network                       Modify relay network
  modify-relay-network-admin-status          Modify relay network
  mr-binding-list                            List mirror flow filter bindings.
  mr-filter-attach                           attach a mirror flow filter
  mr-filter-create                           Create a mirror  filter
  mr-filter-delete                           Delete a Mirror Filter
  mr-filter-detach                           detach a mirror flow filter
  mr-filter-list                             List mirror filter.
  mr-filter-modify                           modify mirror filter
  mr-filter-rule-create                      Create a mirror filter rule
  mr-filter-rule-delete                      Delete a mirror filter rule
  mr-filter-rule-list                        list  mirror filter rule
  mr-filter-rule-modify                      modify  mirror filter rule
  mr-filter-rule-show                        Show a mirror filter rule
  mr-filter-rules-delete                     Delete mirror filter rules
  mr-filter-show                             Show a Mirror Filter
  mr-flow-add-src-ports                      Add source ports to mirror flow.
  mr-flow-create                             Create mirror flow for vpc
  mr-flow-delete                             Delete mirror flow id
  mr-flow-list                               List mirror flow.
  mr-flow-modify                             Modify mirror flow
  mr-flow-modify-admin-status                Modify mirror flow admin status
  mr-flow-remove-src-ports                   Remove source ports from mirror flow.
  mr-flow-show                               Show mirror flow.
  mr-src-port-binding-list                   List mirror flow src port bindings.
  mrgw-add-hostgroup-binding                 Bind the hostgroup and mrgw.
  mrgw-create                                Create mrgw for vpc
  mrgw-delete                                Delete mrgw
  mrgw-list                                  List mrgws.
  mrgw-list-hostgroup-binding                List the binding of mrgw and hostgroup.
  mrgw-modify-hostgroup-binding-role         modify mrgw hostgroup binding
  mrgw-remove-hostgroup-binding              Remove the binding of the hostgroup and mrgw.
  mrgw-show                                  Show mrgw.
  natgw-add-hostgroup-binding                Add natgw hostgroup binding
  natgw-assign-private-ips                   Assign private ips to natgw.
  natgw-associate-eip                        Associate eips into natgw.
  natgw-create                               Create natgw .
  natgw-delete                               Delete natgw
  natgw-disassociate-eip                     Disassociate eips from natgw.
  natgw-dnat-create                          Create DNat rule for Nat Gateway
  natgw-dnat-delete                          Delete DNat rule
  natgw-dnat-list                            List DNat rules of Nat Gateway
  natgw-dnat-modify                          Modify DNat rule for dNat Id
  natgw-list                                 List natgws
  natgw-modify                               Modify attributes of a natgw.
  natgw-modify-admin-status                  Modify natgw admin status
  natgw-modify-hostgroup-binding-role        Remove natgw hostgroup binding
  natgw-remove-hostgroup-binding             Remove natgw hostgroup binding
  natgw-show                                 Show natgw detail
  natgw-snat-add-eips                        Add eips for SNat rule
  natgw-snat-create                          Create SNat rule for Nat Gateway
  natgw-snat-delete                          Delete SNat rules
  natgw-snat-list                            List SNat rules of given Nat Gateway
  natgw-snat-modify                          Modify SNat rule for sNat Id
  natgw-snat-remove-eips                     Remove eips from SNat rule
  natgw-unassign-private-ips                 Unassign private ips from natgw.
  port-add-security-group                    Add security group to port.
  port-assign-ipv6                           Assign ipv6 address to port.
  port-assign-secondary-private-ip           Assign secondary private ip to port.
  port-attach                                Attach port to host.
  port-binding-list                          list port bindings
  port-create                                create port
  port-delete                                delet port id
  port-delete-security-group                 Delete security group from port.
  port-detach                                detach port.
  port-list                                  list ports
  port-modify                                modify port
  port-modify-admin-status                   modify port
  port-replace-binding                       replace port binding
  port-replace-subnet                        replace port subnet
  port-security-group-binding-list           list the binding of security group and port.
  port-show                                  show port
  port-status-check                          Check port status.
  port-status-check-admin                    Check port status.
  port-unassign-ipv6                         Unassign ipv6 address from port.
  port-unassign-secondary-private-ip         Unassign secondary private ip from port.
  ports-assign-ipv6                          Assign ipv6 address to multiple ports. The maximum number of ports is 50.
  prefix-list-associations-list              list the associations resourece by PREFIX_LIST_ID.
  prefix-list-create                         Create a prefix list
  prefix-list-delete                         Delete a prefix list.
  prefix-list-list                           List prefix lists.
  prefix-list-modify                         Modify attributes of a prefix list.
  prefix-list-rule-create                    Create an prefix list rule for a given prefix list.
  prefix-list-rule-delete                    Delete prefix list rules.
  prefix-list-rule-list                      List network prefix list rules.
  prefix-list-rule-modify                    Modify attributes of a prefix list rule.
  prefix-list-show                           Display prefix list details.
  provider-list                              list providers
  psgw-add-hostgroup-binding                 Add psgw hostgroup binding
  psgw-binding-list                          List psgw bindings
  psgw-create                                Create psgw for vpc
  psgw-delete                                Delete psgw
  psgw-list                                  List psgws
  psgw-modify-hostgroup-binding-role         Remove psgw hostgroup binding
  psgw-remove-hostgroup-binding              Remove psgw hostgroup binding
  psgw-show                                  Show psgw detail
  qos-policy-create                          Create qos policy
  qos-policy-delete                          Delete qos policy
  qos-policy-list                            List qos policy
  qos-policy-modify                          Modify qos policy attributes
  qos-policy-show                            Show qos policy detail
  quota-create                               Create a quota for a given type of resource.
  quota-delete                               Delete a quota.
  quota-list                                 List quotas.
  quota-modify                               Modify a quota max limit.
  relay-network-create                       Create a Relay network.
  relay-network-delete                       Delete a relay network.
  relay-network-list                         List relay networks.
  relay-network-show                         Show relay network.
  relay-port-create                          create relay port
  relay-port-delete                          delete relay port id
  relay-port-list                            List relay ports.
  relay-port-show                            Show relay port.
  resource-list                              List Basic view of resources
  route-rules-create                         create a route rule for route table
  route-rules-delete                         delete route table rules
  route-rules-list                           list a route table route rules
  route-rules-modify                         modify a route table route rule
  route-table-create                         create route table
  route-table-delete                         delete route table
  route-table-list                           route table list
  route-table-modify                         modify route table attrs
  route-table-show                           show route table
  router-list                                List vrouters.
  security-group-associate-ports             Add security group to port.
  security-group-create                      Create a security group for a given vpc.
  security-group-delete                      Delete a security group.
  security-group-disassociate-ports          Delete security group from port.
  security-group-list                        List security groups.
  security-group-modify                      Modify attributes of a security group.
  security-group-rule-create                 Create an security group rule for a given security group.
  security-group-rule-delete                 Delete security group rules.
  security-group-rule-list                   List network security group rules.
  security-group-rule-modify                 Modify attributes of a security group rule.
  security-group-show                        Display security group details.
  segment-check-available-ip                 check segment available ip
  segment-pool-create                        Add a Segment Pool for region.
  segment-pool-delete                        Delete given segment pool.
  segment-pool-list                          List Segment Pools.
  segment-pool-modify                        Update given segment pool.
  segment-pool-modify-admin-status           Update the admin status of a segment pool.
  segment-pool-show                          Show Segment Pool.
  service-ip-allocate                        Allocate a service ip
  service-ip-list                            List service ips
  service-ip-release                         Release a service ip
  service-ip-segment-create                  Create a service ip segment
  service-ip-segment-delete                  Delete a service ip segment
  service-ip-segment-list                    List service ip segments
  service-ip-segment-modify                  Modify service ip segment attributes
  service-ip-segment-show                    Show service ip segment
  subnet-check-ip-availability               check if the IP is available in the subnet.
  subnet-create                              Create a SUBNET in VPC.
  subnet-delete                              Delete a given subnet.
  subnet-disable-ipv6                        Disable IPv6 of given SUBNET.
  subnet-disassociate-acl                    disassociate the network acl id of subnet.
  subnet-enable-ipv6                         Enable IPv6 of given SUBNET.
  subnet-list                                List SUBNETs.
  subnet-modify                              Update given subnet.
  subnet-replace-acl                         update the network acl id of subnet.
  subnet-replace-route-table                 update the route table id of subnet.
  subnet-show                                Show SUBNET.
  tunnel-network-create                      Create a Tunnel network for given service.
  tunnel-network-delete                      Delete a tunnel network.
  tunnel-network-list                        List tunnel networks.
  tunnel-network-modify                      Modify tunnel network
  tunnel-network-modify-admin-status         Modify tunnel network
  tunnel-network-show                        Show tunnel network.
  vgw-add-hostgroup-binding                  Add vgw hostgroup binding
  vgw-binding-list                           List vgw bindings
  vgw-create                                 Create vgw for vpc
  vgw-delete                                 Delete vgw
  vgw-list                                   List vgws
  vgw-modify-hostgroup-binding-role          Remove vgw hostgroup binding
  vgw-remove-hostgroup-binding               Remove vgw hostgroup binding
  vgw-show                                   Show vgw detail
  vni-range-create                           Create a VNI Range.
  vni-range-delete                           Delete a VNI Range.
  vni-range-list                             List VNI ranges.
  vni-range-show                             List VNI ranges.
  vpc-add-v2gw                               Add v2gw for given VPC.
  vpc-associate-cidrs                        Associate Secondary CIDRs to given VPC.
  vpc-associate-dhcp-options-set             Associate vpc with dhcp options set.
  vpc-connection-create                      Create VPC Connection for tenant
  vpc-connection-delete                      Delete VPC Connection
  vpc-connection-list                        List VPC Connections with filters
  vpc-connection-modify                      Modify VPC Connection attributes
  vpc-connection-modify-status               Modify VPC Connection status
  vpc-connection-show                        Show VPC Connection details
  vpc-create                                 Create a VPC(Virtual Private Cloud) for tenant.
  vpc-delete                                 Delete given vpc.
  vpc-dhcp-options-set-binding-list          List the binding of dhcp options set and vpc.
  vpc-dhcp-options-set-binding-show          Show the binding of dhcp options set and vpc.
  vpc-disable-ipv6                           Disable IPv6 of given VPC.
  vpc-disassociate-cidrs                     Disassociate Secondary CIDRs from given VPC.
  vpc-disassociate-dhcp-options-set          Disassociate vpc from dhcp options set.
  vpc-enable-ipv6                            Enable IPv6 of given VPC.
  vpc-list                                   List VPCs.
  vpc-modify                                 Update given vpc's attribute.
  vpc-peering-create
  vpc-peering-delete
  vpc-peering-list
  vpc-peering-modify
  vpc-remove-v2gw                            Remove v2gw for given VPC.
  vpc-replace-dhcp-options-set               Replace dhcp options set associated with vpc.
  vpc-show                                   Show VPC.
  vpn-connection-create                      Create VPN Connection
  vpn-connection-delete                      Delete VPN Connection
  vpn-connection-list                        List VPN Connection with filters
  vpn-connection-modify                      Modify VPN Connection attributes.
  vpn-connection-show                        Show VPN Connection details
  vpn-gateway-create                         Create VPN Gateway
  vpn-gateway-delete                         Delete VPN Gateway
  vpn-gateway-list                           List VPN Gateway with filters
  vpn-gateway-modify                         Modify VPN Gateway attributes.
  vpn-gateway-show                           Show VPN Gateway details.
  zone-associate-vpc                         Associate vpc and zone.
  zone-create                                Create zone .
  zone-delete                                Delete zone
  zone-disassociate-vpc                      Remove vpcs from zone.
  zone-list                                  List zones
  zone-modify                                Modify attributes of a zone.
  zone-modify-admin-status                   Modify zone admin status
  zone-record-create                         Create zone record.
  zone-record-delete                         Delete zone record
  zone-record-list                           List zone records
  zone-record-modify                         Modify attributes of a zone record.
  zone-record-show                           Show zone record detail
  zone-records-create                        Create zone record.
  zone-show                                  Show zone detail
  zone-vpc-association-list                  List zone vpc associations.

Flags:
  -a, --admin                 admin role
  -f, --format string         Format to print result, support table/vertical/json (default "table")
  -h, --help                  Type for help
      --server-url string     URL of vpc server api or set by env 'VPC_SERVER_URL' (default "http://10.8.73.43:32198/vpc-server")
      --tenant-id string      Tenant ID used to call api
      --trace-id string       trace id (default "trace-gtvi083wvd")
  -v, --verbose               verbose output
      --with-json-formatted   Print json in formatted, only for format 'json'
      --with-line-number      Print table with line number, only for format 'table'

Use "vpcs [command] --help" for more information about a command.

```

