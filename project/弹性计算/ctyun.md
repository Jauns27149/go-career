

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

# gs

内蒙08 -> 10.8.73.43
配置环境  . admin-openrc az1

- - -

## virsh

- virsh
  - list : 虚拟机列表
    - --limit <int> : 限制返回条目数量
    - --status <string> ：按状态过滤卷，常见的状态包括 `available`, `in-use`, `error` 
    - --tenant <tenant-id> : 按租户ID过滤卷（虚拟机的 project_id 为 tenant_id）
  - console <instance-id> : 进入虚拟机

## cinder

- cinder
  - list : 查看卷列表
    - --limit <int> : 限制返回条目数量
    - --name <string> : 指定卷名称
  - show <volume-id> : 卷详情
  - delete <volume-id> : 删除卷
  - create [flag] <size> :  新建卷
    - --name <string> : 卷名称

## vnetops

- vnetops
  - vpc
    - port-detach <port-id>
    - -- tenant-id <project-id> 

# gostack执行流程

项目结构

- gostack
  - agent
  - api
  - engine
  - scheduler
  - cron

```bash
[root@cn-nm-region1-az1-control-10e8e73e43 etc]# kubectl get pods | grep etcd
[root@cn-nm-region1-az1-control-10e8e73e43 etc]# kubectl get svc gostack-etcd -n az1 -o wide
NAME           TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)             AGE    SELECTOR
gostack-etcd   ClusterIP   10.96.27.16   <none>        2379/TCP,2380/TCP   378d   app.kubernetes.io/instance=gostack-etcd
[root@cn-nm-region1-az1-control-10e8e73e43 etc]# kubectl get pods -n az1 -l app.kubernetes.io/instance=gostack-etcd -o wide
NAME                                               READY   STATUS    RESTARTS   AGE   IP           NODE                                   NOMINATED NODE   READINESS GATES
cn-nm-region1-az1-gostack-etcd2-855574798d-q572f   1/1     Running   0          13d   10.8.73.45   cn-nm-region1-az1-control-10e8e73e45   <none>           <none>
[root@cn-nm-region1-az1-control-10e8e73e43 etc]# kubectl exec -it cn-nm-region1-az1-gostack-etcd2-855574798d-q572f  -n az1 -- /bin/sh
# ls
bin  boot  dev  etc  etcd-data  home  lib  lib64  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
# cd etcd-d
/bin/sh: 2: cd: can't cd to etcd-d



 etcdctl get "/template" --prefix --user=root:2021CTyu
 
 
 
 [root@cn-nm-region1-az1-control-10e8e73e43 gostack]# cat gsinitrc.yml-az1
#AZ1
AppInfo:
  Username: "root"
  Password: "2021CTyun!"
  EndPoint: "gostack-etcd.az1.svc.cluster.net:2379"
  BakEndPoint: "gostack-etcd-backup.az1.svc.cluster.net:2379"
  MaxCallSendMsgSize: 10485760
  MaxCallRecvMsgSize: 107374182400
  KeepAliveTime: 30
  KeepAliveTimeout: 10
MongoConf:
  EndPoint: "gostack-mongos.az1.svc.cluster.net:27017"
  EndPoints:
    - "gostack-mongos.az1.svc.cluster.net:27017"
  Username: "root"
  Password: "test"
  Timeout: 300
  Worker: 10

```

# glance

```bash
usage: glance [--version] [-d] [-v] [--get-schema] [-f]
              [--os-image-url OS_IMAGE_URL]
              [--os-image-api-version OS_IMAGE_API_VERSION]
              [--profile HMAC_KEY] [--key-file OS_KEY] [--ca-file OS_CACERT]
              [--cert-file OS_CERT] [--os-region-name OS_REGION_NAME]
              [--os-auth-token OS_AUTH_TOKEN]
              [--os-service-type OS_SERVICE_TYPE]
              [--os-endpoint-type OS_ENDPOINT_TYPE] [--insecure]
              [--os-cacert <ca-certificate>] [--os-cert <certificate>]
              [--os-key <key>] [--timeout <seconds>] [--os-auth-type <name>]
              [--os-auth-url OS_AUTH_URL] [--os-system-scope OS_SYSTEM_SCOPE]
              [--os-domain-id OS_DOMAIN_ID] [--os-domain-name OS_DOMAIN_NAME]
              [--os-project-id OS_PROJECT_ID]
              [--os-project-name OS_PROJECT_NAME]
              [--os-project-domain-id OS_PROJECT_DOMAIN_ID]
              [--os-project-domain-name OS_PROJECT_DOMAIN_NAME]
              [--os-trust-id OS_TRUST_ID]
              [--os-default-domain-id OS_DEFAULT_DOMAIN_ID]
              [--os-default-domain-name OS_DEFAULT_DOMAIN_NAME]
              [--os-user-id OS_USER_ID] [--os-username OS_USERNAME]
              [--os-user-domain-id OS_USER_DOMAIN_ID]
              [--os-user-domain-name OS_USER_DOMAIN_NAME]
              [--os-password OS_PASSWORD]
              <subcommand> ...

Command-line interface to the OpenStack Images API.

Positional arguments:
  <subcommand>
    backend-list        Show the rbd backends .
    explain             Describe a specific model.
    image-create        Create a new image.
    image-create-via-import
                        EXPERIMENTAL: Create a new image via image import.
    image-deactivate    Deactivate specified image.
    image-delete        Delete specified image.
    image-download      Download a specific image.
    image-import        Initiate the image import taskflow.
    image-list          List images you can access.
    image-quota         Show the quota of current user.
    image-reactivate    Reactivate specified image.
    image-show          Describe a specific image.
    image-stage         Upload data for a specific image to staging.
    image-sync          Add a location (and related metadata) to an image.
    image-tag-delete    Delete the tag associated with the given image.
    image-tag-update    Update an image with the given tag.
    image-update        Update an existing image.
    image-upload        Upload data for a specific image.
    import-info         Print import methods available from Glance.
    location-add        Add a location (and related metadata) to an image.
    location-delete     Remove locations (and related metadata) from an image.
    location-update     Update metadata of an image's location.
    md-namespace-create
                        Create a new metadata definitions namespace.
    md-namespace-delete
                        Delete specified metadata definitions namespace with
                        its contents.
    md-namespace-import
                        Import a metadata definitions namespace from file or
                        standard input.
    md-namespace-list   List metadata definitions namespaces.
    md-namespace-objects-delete
                        Delete all metadata definitions objects inside a
                        specific namespace.
    md-namespace-properties-delete
                        Delete all metadata definitions property inside a
                        specific namespace.
    md-namespace-resource-type-list
                        List resource types associated to specific namespace.
    md-namespace-show   Describe a specific metadata definitions namespace.
    md-namespace-tags-delete
                        Delete all metadata definitions tags inside a specific
                        namespace.
    md-namespace-update
                        Update an existing metadata definitions namespace.
    md-object-create    Create a new metadata definitions object inside a
                        namespace.
    md-object-delete    Delete a specific metadata definitions object inside a
                        namespace.
    md-object-list      List metadata definitions objects inside a specific
                        namespace.
    md-object-property-show
                        Describe a specific metadata definitions property
                        inside an object.
    md-object-show      Describe a specific metadata definitions object inside
                        a namespace.
    md-object-update    Update metadata definitions object inside a namespace.
    md-property-create  Create a new metadata definitions property inside a
                        namespace.
    md-property-delete  Delete a specific metadata definitions property inside
                        a namespace.
    md-property-list    List metadata definitions properties inside a specific
                        namespace.
    md-property-show    Describe a specific metadata definitions property
                        inside a namespace.
    md-property-update  Update metadata definitions property inside a
                        namespace.
    md-resource-type-associate
                        Associate resource type with a metadata definitions
                        namespace.
    md-resource-type-deassociate
                        Deassociate resource type with a metadata definitions
                        namespace.
    md-resource-type-list
                        List available resource type names.
    md-tag-create       Add a new metadata definitions tag inside a namespace.
    md-tag-create-multiple
                        Create new metadata definitions tags inside a
                        namespace.
    md-tag-delete       Delete a specific metadata definitions tag inside a
                        namespace.
    md-tag-list         List metadata definitions tags inside a specific
                        namespace.
    md-tag-show         Describe a specific metadata definitions tag inside a
                        namespace.
    md-tag-update       Rename a metadata definitions tag inside a namespace.
    member-create       Create member for a given image.
    member-delete       Delete image member.
    member-list         Describe sharing permissions by image.
    member-update       Update the status of a member for a given image.
    stores-info         Print available backends from Glance.
    task-create         Create a new task.
    task-list           List tasks you can access.
    task-show           Describe a specific task.
    bash-completion     Prints arguments for bash_completion.
    help                Display help about this program or one of its
                        subcommands.

Optional arguments:
  --version             show program's version number and exit
  -d, --debug           Defaults to env[GLANCECLIENT_DEBUG].
  -v, --verbose         Print more verbose output.
  --get-schema          Ignores cached copy and forces retrieval of schema
                        that generates portions of the help text. Ignored with
                        API version 1.
  -f, --force           Prevent select actions from requesting user
                        confirmation.
  --os-image-url OS_IMAGE_URL
                        Defaults to env[OS_IMAGE_URL]. If the provided image
                        url contains a version number and `--os-image-api-
                        version` is omitted the version of the URL will be
                        picked as the image api version to use.
  --os-image-api-version OS_IMAGE_API_VERSION
                        Defaults to env[OS_IMAGE_API_VERSION] or 2.
  --profile HMAC_KEY    HMAC key to use for encrypting context data for
                        performance profiling of operation. This key should be
                        the value of HMAC key configured in osprofiler
                        middleware in glance, it is specified in glance
                        configuration file at /etc/glance/glance-api.conf and
                        /etc/glance/glance-registry.conf. Without key the
                        profiling will not be triggered even if osprofiler is
                        enabled on server side. Defaults to env[OS_PROFILE].
  --key-file OS_KEY     DEPRECATED! Use --os-key.
  --ca-file OS_CACERT   DEPRECATED! Use --os-cacert.
  --cert-file OS_CERT   DEPRECATED! Use --os-cert.
  --os-region-name OS_REGION_NAME
                        Defaults to env[OS_REGION_NAME].
  --os-auth-token OS_AUTH_TOKEN
                        Defaults to env[OS_AUTH_TOKEN].
  --os-service-type OS_SERVICE_TYPE
                        Defaults to env[OS_SERVICE_TYPE].
  --os-endpoint-type OS_ENDPOINT_TYPE
                        Defaults to env[OS_ENDPOINT_TYPE].
  --os-auth-type <name>, --os-auth-plugin <name>
                        Authentication type to use

API Connection Options:
  Options controlling the HTTP API Connections

  --insecure            Explicitly allow client to perform "insecure" TLS
                        (https) requests. The server's certificate will not be
                        verified against any certificate authorities. This
                        option should be used with caution.
  --os-cacert <ca-certificate>
                        Specify a CA bundle file to use in verifying a TLS
                        (https) server certificate. Defaults to
                        env[OS_CACERT].
  --os-cert <certificate>
                        Defaults to env[OS_CERT].
  --os-key <key>        Defaults to env[OS_KEY].
  --timeout <seconds>   Set request timeout (in seconds).

Authentication Options:
  Options specific to the password plugin.

  --os-auth-url OS_AUTH_URL
                        Authentication URL
  --os-system-scope OS_SYSTEM_SCOPE
                        Scope for system operations
  --os-domain-id OS_DOMAIN_ID
                        Domain ID to scope to
  --os-domain-name OS_DOMAIN_NAME
                        Domain name to scope to
  --os-project-id OS_PROJECT_ID, --os-tenant-id OS_PROJECT_ID
                        Project ID to scope to
  --os-project-name OS_PROJECT_NAME, --os-tenant-name OS_PROJECT_NAME
                        Project name to scope to
  --os-project-domain-id OS_PROJECT_DOMAIN_ID
                        Domain ID containing project
  --os-project-domain-name OS_PROJECT_DOMAIN_NAME
                        Domain name containing project
  --os-trust-id OS_TRUST_ID
                        Trust ID
  --os-default-domain-id OS_DEFAULT_DOMAIN_ID
                        Optional domain ID to use with v3 and v2 parameters.
                        It will be used for both the user and project domain
                        in v3 and ignored in v2 authentication.
  --os-default-domain-name OS_DEFAULT_DOMAIN_NAME
                        Optional domain name to use with v3 API and v2
                        parameters. It will be used for both the user and
                        project domain in v3 and ignored in v2 authentication.
  --os-user-id OS_USER_ID
                        User id
  --os-username OS_USERNAME, --os-user-name OS_USERNAME
                        Username
  --os-user-domain-id OS_USER_DOMAIN_ID
                        User's domain id
  --os-user-domain-name OS_USER_DOMAIN_NAME
                        User's domain name
  --os-password OS_PASSWORD
                        User's password
```

## imge-list

```bash
glance image-list [--limit <LIMIT>] [--page-size <SIZE>]
                         [--visibility <VISIBILITY>]
                         [--member-status <MEMBER_STATUS>] [--owner <OWNER>]
                         [--property-filter <KEY=VALUE>]
                         [--checksum <CHECKSUM>] [--tag <TAG>]
                         [--sort-key {name,status,container_format,disk_format,size,id,created_at,updated_at}]
                         [--sort-dir {asc,desc}] [--sort <key>[:<direction>]]
```

