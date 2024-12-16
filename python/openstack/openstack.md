

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

```bash
openstack [command]
```

```bash
usage: openstack [--version] [-v | -q] [--log-file LOG_FILE] [-h] [--debug]
                 [--os-cloud <cloud-config-name>]
                 [--os-region-name <auth-region-name>]
                 [--os-cacert <ca-bundle-file>] [--os-cert <certificate-file>]
                 [--os-key <key-file>] [--verify | --insecure]
                 [--os-default-domain <auth-domain>]
                 [--os-interface <interface>]
                 [--os-service-provider <service_provider>]
                 [--os-remote-project-name <remote_project_name> | --os-remote-project-id <remote_project_id>]
                 [--os-remote-project-domain-name <remote_project_domain_name> | --os-remote-project-domain-id <remote_project_domain_id>]
                 [--timing] [--os-beta-command] [--os-profile hmac-key]
                 [--os-compute-api-version <compute-api-version>]
                 [--os-network-api-version <network-api-version>]
                 [--os-image-api-version <image-api-version>]
                 [--os-volume-api-version <volume-api-version>]
                 [--os-delete-token <delete-token>]
                 [--os-identity-api-version <identity-api-version>]
                 [--os-object-api-version <object-api-version>]
                 [--os-key-manager-api-version <key-manager-api-version>]
                 [--os-dns-api-version <dns-api-version>]
                 [--os-auth-type <auth-type>]
                 [--os-project-domain-id <auth-project-domain-id>]
                 [--os-protocol <auth-protocol>]
                 [--os-project-name <auth-project-name>]
                 [--os-trust-id <auth-trust-id>]
                 [--os-service-provider-endpoint <auth-service-provider-endpoint>]
                 [--os-domain-name <auth-domain-name>]
                 [--os-access-secret <auth-access-secret>]
                 [--os-user-domain-id <auth-user-domain-id>]
                 [--os-access-token-type <auth-access-token-type>]
                 [--os-code <auth-code>]
                 [--os-application-credential-name <auth-application-credential-name>]
                 [--os-identity-provider-url <auth-identity-provider-url>]
                 [--os-default-domain-name <auth-default-domain-name>]
                 [--os-access-token-endpoint <auth-access-token-endpoint>]
                 [--os-access-token <auth-access-token>]
                 [--os-domain-id <auth-domain-id>]
                 [--os-user-domain-name <auth-user-domain-name>]
                 [--os-openid-scope <auth-openid-scope>]
                 [--os-user-id <auth-user-id>]
                 [--os-application-credential-secret <auth-application-credential-secret>]
                 [--os-identity-provider <auth-identity-provider>]
                 [--os-username <auth-username>]
                 [--os-auth-url <auth-auth-url>]
                 [--os-client-secret <auth-client-secret>]
                 [--os-default-domain-id <auth-default-domain-id>]
                 [--os-discovery-endpoint <auth-discovery-endpoint>]
                 [--os-client-id <auth-client-id>]
                 [--os-project-domain-name <auth-project-domain-name>]
                 [--os-service-provider-entity-id <auth-service-provider-entity-id>]
                 [--os-access-key <auth-access-key>]
                 [--os-password <auth-password>]
                 [--os-redirect-uri <auth-redirect-uri>]
                 [--os-endpoint <auth-endpoint>] [--os-url <auth-url>]
                 [--os-consumer-key <auth-consumer-key>]
                 [--os-consumer-secret <auth-consumer-secret>]
                 [--os-token <auth-token>]
                 [--os-application-credential-id <auth-application-credential-id>]
                 [--os-passcode <auth-passcode>]
                 [--os-system-scope <auth-system-scope>]
                 [--os-project-id <auth-project-id>]

Command-line interface to the OpenStack APIs

optional arguments:
  --version             show program's version number and exit
  -v, --verbose         Increase verbosity of output. Can be repeated.
  -q, --quiet           Suppress output except warnings and errors.
  --log-file LOG_FILE   Specify a file to log output. Disabled by default.
  -h, --help            Show help message and exit.
  --debug               Show tracebacks on errors.
  --os-cloud <cloud-config-name>
                        Cloud name in clouds.yaml (Env: OS_CLOUD)
  --os-region-name <auth-region-name>
                        Authentication region name (Env: OS_REGION_NAME)
  --os-cacert <ca-bundle-file>
                        CA certificate bundle file (Env: OS_CACERT)
  --os-cert <certificate-file>
                        Client certificate bundle file (Env: OS_CERT)
  --os-key <key-file>   Client certificate key file (Env: OS_KEY)
  --verify              Verify server certificate (default)
  --insecure            Disable server certificate verification
  --os-default-domain <auth-domain>
                        Default domain ID, default=default. (Env:
                        OS_DEFAULT_DOMAIN)
  --os-interface <interface>
                        Select an interface type. Valid interface types:
                        [admin, public, internal]. (Env: OS_INTERFACE)
  --os-service-provider <service_provider>
                        Authenticate with and perform the command on a service
                        provider using Keystone-to-keystone federation. Must
                        also specify the remote project option.
  --os-remote-project-name <remote_project_name>
                        Project name when authenticating to a service provider
                        if using Keystone-to-Keystone federation.
  --os-remote-project-id <remote_project_id>
                        Project ID when authenticating to a service provider
                        if using Keystone-to-Keystone federation.
  --os-remote-project-domain-name <remote_project_domain_name>
                        Domain name of the project when authenticating to a
                        service provider if using Keystone-to-Keystone
                        federation.
  --os-remote-project-domain-id <remote_project_domain_id>
                        Domain ID of the project when authenticating to a
                        service provider if using Keystone-to-Keystone
                        federation.
  --timing              Print API call timing info
  --os-beta-command     Enable beta commands which are subject to change
  --os-profile hmac-key
                        HMAC key for encrypting profiling context data
  --os-compute-api-version <compute-api-version>
                        Compute API version, default=2.1 (Env:
                        OS_COMPUTE_API_VERSION)
  --os-network-api-version <network-api-version>
                        Network API version, default=2.0 (Env:
                        OS_NETWORK_API_VERSION)
  --os-image-api-version <image-api-version>
                        Image API version, default=2 (Env:
                        OS_IMAGE_API_VERSION)
  --os-volume-api-version <volume-api-version>
                        Volume API version, default=2 (Env:
                        OS_VOLUME_API_VERSION)
  --os-delete-token <delete-token>
                        Delete token for admin role, default=(Env:
                        OS_DELETE_TOKEN)
  --os-identity-api-version <identity-api-version>
                        Identity API version, default=3 (Env:
                        OS_IDENTITY_API_VERSION)
  --os-object-api-version <object-api-version>
                        Object API version, default=1 (Env:
                        OS_OBJECT_API_VERSION)
  --os-key-manager-api-version <key-manager-api-version>
                        Barbican API version, default=1 (Env:
                        OS_KEY_MANAGER_API_VERSION)
  --os-dns-api-version <dns-api-version>
                        DNS API version, default=2 (Env: OS_DNS_API_VERSION)
  --os-auth-type <auth-type>
                        Select an authentication type. Available types:
                        v2token, none, password, admin_token, v3oidcauthcode,
                        v2password, v3samlpassword, v3password,
                        v3adfspassword, v3oidcaccesstoken, v3oidcpassword,
                        token, v3oidcclientcredentials, v3tokenlessauth,
                        v1password, v3token, v3totp, v3applicationcredential,
                        v3oauth1, token_endpoint, noauth. Default: selected
                        based on --os-username/--os-token (Env: OS_AUTH_TYPE)
  --os-project-domain-id <auth-project-domain-id>
                        With password: Domain ID containing project With
                        v3oidcauthcode: Domain ID containing project With
                        v3samlpassword: Domain ID containing project With
                        v3password: Domain ID containing project With
                        v3adfspassword: Domain ID containing project With
                        v3oidcaccesstoken: Domain ID containing project With
                        v3oidcpassword: Domain ID containing project With
                        token: Domain ID containing project With
                        v3oidcclientcredentials: Domain ID containing project
                        With v3tokenlessauth: Domain ID containing project
                        With v3token: Domain ID containing project With
                        v3totp: Domain ID containing project With
                        v3applicationcredential: Domain ID containing project
                        (Env: OS_PROJECT_DOMAIN_ID)
  --os-protocol <auth-protocol>
                        With v3oidcauthcode: Protocol for federated plugin
                        With v3samlpassword: Protocol for federated plugin
                        With v3adfspassword: Protocol for federated plugin
                        With v3oidcaccesstoken: Protocol for federated plugin
                        With v3oidcpassword: Protocol for federated plugin
                        With v3oidcclientcredentials: Protocol for federated
                        plugin (Env: OS_PROTOCOL)
  --os-project-name <auth-project-name>
                        With password: Project name to scope to With
                        v3oidcauthcode: Project name to scope to With
                        v3samlpassword: Project name to scope to With
                        v3password: Project name to scope to With
                        v3adfspassword: Project name to scope to With
                        v3oidcaccesstoken: Project name to scope to With
                        v3oidcpassword: Project name to scope to With token:
                        Project name to scope to With v3oidcclientcredentials:
                        Project name to scope to With v3tokenlessauth: Project
                        name to scope to With v1password: Swift account to use
                        With v3token: Project name to scope to With v3totp:
                        Project name to scope to With v3applicationcredential:
                        Project name to scope to (Env: OS_PROJECT_NAME)
  --os-trust-id <auth-trust-id>
                        With v2token: Trust ID With password: Trust ID With
                        v3oidcauthcode: Trust ID With v2password: Trust ID
                        With v3samlpassword: Trust ID With v3password: Trust
                        ID With v3adfspassword: Trust ID With
                        v3oidcaccesstoken: Trust ID With v3oidcpassword: Trust
                        ID With token: Trust ID With v3oidcclientcredentials:
                        Trust ID With v3token: Trust ID With v3totp: Trust ID
                        With v3applicationcredential: Trust ID (Env:
                        OS_TRUST_ID)
  --os-service-provider-endpoint <auth-service-provider-endpoint>
                        With v3adfspassword: Service Provider's Endpoint (Env:
                        OS_SERVICE_PROVIDER_ENDPOINT)
  --os-domain-name <auth-domain-name>
                        With password: Domain name to scope to With
                        v3oidcauthcode: Domain name to scope to With
                        v3samlpassword: Domain name to scope to With
                        v3password: Domain name to scope to With
                        v3adfspassword: Domain name to scope to With
                        v3oidcaccesstoken: Domain name to scope to With
                        v3oidcpassword: Domain name to scope to With token:
                        Domain name to scope to With v3oidcclientcredentials:
                        Domain name to scope to With v3tokenlessauth: Domain
                        name to scope to With v3token: Domain name to scope to
                        With v3totp: Domain name to scope to With
                        v3applicationcredential: Domain name to scope to (Env:
                        OS_DOMAIN_NAME)
  --os-access-secret <auth-access-secret>
                        With v3oauth1: OAuth Access Secret (Env:
                        OS_ACCESS_SECRET)
  --os-user-domain-id <auth-user-domain-id>
                        With password: User's domain id With v3password:
                        User's domain id With v3totp: User's domain id With
                        v3applicationcredential: User's domain id (Env:
                        OS_USER_DOMAIN_ID)
  --os-access-token-type <auth-access-token-type>
                        With v3oidcauthcode: OAuth 2.0 Authorization Server
                        Introspection token type, it is used to decide which
                        type of token will be used when processing token
                        introspection. Valid values are: "access_token" or
                        "id_token" With v3oidcpassword: OAuth 2.0
                        Authorization Server Introspection token type, it is
                        used to decide which type of token will be used when
                        processing token introspection. Valid values are:
                        "access_token" or "id_token" With
                        v3oidcclientcredentials: OAuth 2.0 Authorization
                        Server Introspection token type, it is used to decide
                        which type of token will be used when processing token
                        introspection. Valid values are: "access_token" or
                        "id_token" (Env: OS_ACCESS_TOKEN_TYPE)
  --os-code <auth-code>
                        With v3oidcauthcode: OAuth 2.0 Authorization Code
                        (Env: OS_CODE)
  --os-application-credential-name <auth-application-credential-name>
                        With v3applicationcredential: Application credential
                        name (Env: OS_APPLICATION_CREDENTIAL_NAME)
  --os-identity-provider-url <auth-identity-provider-url>
                        With v3samlpassword: An Identity Provider URL, where
                        the SAML2 authentication request will be sent. With
                        v3adfspassword: An Identity Provider URL, where the
                        SAML authentication request will be sent. (Env:
                        OS_IDENTITY_PROVIDER_URL)
  --os-default-domain-name <auth-default-domain-name>
                        With password: Optional domain name to use with v3 API
                        and v2 parameters. It will be used for both the user
                        and project domain in v3 and ignored in v2
                        authentication. With token: Optional domain name to
                        use with v3 API and v2 parameters. It will be used for
                        both the user and project domain in v3 and ignored in
                        v2 authentication. (Env: OS_DEFAULT_DOMAIN_NAME)
  --os-access-token-endpoint <auth-access-token-endpoint>
                        With v3oidcauthcode: OpenID Connect Provider Token
                        Endpoint. Note that if a discovery document is being
                        passed this option will override the endpoint provided
                        by the server in the discovery document. With
                        v3oidcpassword: OpenID Connect Provider Token
                        Endpoint. Note that if a discovery document is being
                        passed this option will override the endpoint provided
                        by the server in the discovery document. With
                        v3oidcclientcredentials: OpenID Connect Provider Token
                        Endpoint. Note that if a discovery document is being
                        passed this option will override the endpoint provided
                        by the server in the discovery document. (Env:
                        OS_ACCESS_TOKEN_ENDPOINT)
  --os-access-token <auth-access-token>
                        With v3oidcaccesstoken: OAuth 2.0 Access Token (Env:
                        OS_ACCESS_TOKEN)
  --os-domain-id <auth-domain-id>
                        With password: Domain ID to scope to With
                        v3oidcauthcode: Domain ID to scope to With
                        v3samlpassword: Domain ID to scope to With v3password:
                        Domain ID to scope to With v3adfspassword: Domain ID
                        to scope to With v3oidcaccesstoken: Domain ID to scope
                        to With v3oidcpassword: Domain ID to scope to With
                        token: Domain ID to scope to With
                        v3oidcclientcredentials: Domain ID to scope to With
                        v3tokenlessauth: Domain ID to scope to With v3token:
                        Domain ID to scope to With v3totp: Domain ID to scope
                        to With v3applicationcredential: Domain ID to scope to
                        (Env: OS_DOMAIN_ID)
  --os-user-domain-name <auth-user-domain-name>
                        With password: User's domain name With v3password:
                        User's domain name With v3totp: User's domain name
                        With v3applicationcredential: User's domain name (Env:
                        OS_USER_DOMAIN_NAME)
  --os-openid-scope <auth-openid-scope>
                        With v3oidcauthcode: OpenID Connect scope that is
                        requested from authorization server. Note that the
                        OpenID Connect specification states that "openid" must
                        be always specified. With v3oidcpassword: OpenID
                        Connect scope that is requested from authorization
                        server. Note that the OpenID Connect specification
                        states that "openid" must be always specified. With
                        v3oidcclientcredentials: OpenID Connect scope that is
                        requested from authorization server. Note that the
                        OpenID Connect specification states that "openid" must
                        be always specified. (Env: OS_OPENID_SCOPE)
  --os-user-id <auth-user-id>
                        With password: User id With v2password: User ID to
                        login with With v3password: User ID With v3totp: User
                        ID With v3applicationcredential: User ID With noauth:
                        User ID (Env: OS_USER_ID)
  --os-application-credential-secret <auth-application-credential-secret>
                        With v3applicationcredential: Application credential
                        auth secret (Env: OS_APPLICATION_CREDENTIAL_SECRET)
  --os-identity-provider <auth-identity-provider>
                        With v3oidcauthcode: Identity Provider's name With
                        v3samlpassword: Identity Provider's name With
                        v3adfspassword: Identity Provider's name With
                        v3oidcaccesstoken: Identity Provider's name With
                        v3oidcpassword: Identity Provider's name With
                        v3oidcclientcredentials: Identity Provider's name
                        (Env: OS_IDENTITY_PROVIDER)
  --os-username <auth-username>
                        With password: Username With v2password: Username to
                        login with With v3samlpassword: Username With
                        v3password: Username With v3adfspassword: Username
                        With v3oidcpassword: Username With v1password:
                        Username to login with With v3totp: Username With
                        v3applicationcredential: Username (Env: OS_USERNAME)
  --os-auth-url <auth-auth-url>
                        With v2token: Authentication URL With password:
                        Authentication URL With v3oidcauthcode: Authentication
                        URL With v2password: Authentication URL With
                        v3samlpassword: Authentication URL With v3password:
                        Authentication URL With v3adfspassword: Authentication
                        URL With v3oidcaccesstoken: Authentication URL With
                        v3oidcpassword: Authentication URL With token:
                        Authentication URL With v3oidcclientcredentials:
                        Authentication URL With v3tokenlessauth:
                        Authentication URL With v1password: Authentication URL
                        With v3token: Authentication URL With v3totp:
                        Authentication URL With v3applicationcredential:
                        Authentication URL With v3oauth1: Authentication URL
                        (Env: OS_AUTH_URL)
  --os-client-secret <auth-client-secret>
                        With v3oidcauthcode: OAuth 2.0 Client Secret With
                        v3oidcpassword: OAuth 2.0 Client Secret With
                        v3oidcclientcredentials: OAuth 2.0 Client Secret (Env:
                        OS_CLIENT_SECRET)
  --os-default-domain-id <auth-default-domain-id>
                        With password: Optional domain ID to use with v3 and
                        v2 parameters. It will be used for both the user and
                        project domain in v3 and ignored in v2 authentication.
                        With token: Optional domain ID to use with v3 and v2
                        parameters. It will be used for both the user and
                        project domain in v3 and ignored in v2 authentication.
                        (Env: OS_DEFAULT_DOMAIN_ID)
  --os-discovery-endpoint <auth-discovery-endpoint>
                        With v3oidcauthcode: OpenID Connect Discovery Document
                        URL. The discovery document will be used to obtain the
                        values of the access token endpoint and the
                        authentication endpoint. This URL should look like
                        https://idp.example.org/.well-known/openid-
                        configuration With v3oidcpassword: OpenID Connect
                        Discovery Document URL. The discovery document will be
                        used to obtain the values of the access token endpoint
                        and the authentication endpoint. This URL should look
                        like https://idp.example.org/.well-known/openid-
                        configuration With v3oidcclientcredentials: OpenID
                        Connect Discovery Document URL. The discovery document
                        will be used to obtain the values of the access token
                        endpoint and the authentication endpoint. This URL
                        should look like https://idp.example.org/.well-known
                        /openid-configuration (Env: OS_DISCOVERY_ENDPOINT)
  --os-client-id <auth-client-id>
                        With v3oidcauthcode: OAuth 2.0 Client ID With
                        v3oidcpassword: OAuth 2.0 Client ID With
                        v3oidcclientcredentials: OAuth 2.0 Client ID (Env:
                        OS_CLIENT_ID)
  --os-project-domain-name <auth-project-domain-name>
                        With password: Domain name containing project With
                        v3oidcauthcode: Domain name containing project With
                        v3samlpassword: Domain name containing project With
                        v3password: Domain name containing project With
                        v3adfspassword: Domain name containing project With
                        v3oidcaccesstoken: Domain name containing project With
                        v3oidcpassword: Domain name containing project With
                        token: Domain name containing project With
                        v3oidcclientcredentials: Domain name containing
                        project With v3tokenlessauth: Domain name containing
                        project With v3token: Domain name containing project
                        With v3totp: Domain name containing project With
                        v3applicationcredential: Domain name containing
                        project (Env: OS_PROJECT_DOMAIN_NAME)
  --os-service-provider-entity-id <auth-service-provider-entity-id>
                        With v3adfspassword: Service Provider's SAML Entity ID
                        (Env: OS_SERVICE_PROVIDER_ENTITY_ID)
  --os-access-key <auth-access-key>
                        With v3oauth1: OAuth Access Key (Env: OS_ACCESS_KEY)
  --os-password <auth-password>
                        With password: User's password With v2password:
                        Password to use With v3samlpassword: Password With
                        v3password: User's password With v3adfspassword:
                        Password With v3oidcpassword: Password With
                        v1password: Password to use (Env: OS_PASSWORD)
  --os-redirect-uri <auth-redirect-uri>
                        With v3oidcauthcode: OpenID Connect Redirect URL (Env:
                        OS_REDIRECT_URI)
  --os-endpoint <auth-endpoint>
                        With none: The endpoint that will always be used With
                        admin_token: The endpoint that will always be used
                        With noauth: Cinder endpoint (Env: OS_ENDPOINT)
  --os-url <auth-url>   With token_endpoint: Specific service endpoint to use
                        (Env: OS_URL)
  --os-consumer-key <auth-consumer-key>
                        With v3oauth1: OAuth Consumer ID/Key (Env:
                        OS_CONSUMER_KEY)
  --os-consumer-secret <auth-consumer-secret>
                        With v3oauth1: OAuth Consumer Secret (Env:
                        OS_CONSUMER_SECRET)
  --os-token <auth-token>
                        With v2token: Token With admin_token: The token that
                        will always be used With token: Token to authenticate
                        with With v3token: Token to authenticate with With
                        token_endpoint: Authentication token to use (Env:
                        OS_TOKEN)
  --os-application-credential-id <auth-application-credential-id>
                        With v3applicationcredential: Application credential
                        ID (Env: OS_APPLICATION_CREDENTIAL_ID)
  --os-passcode <auth-passcode>
                        With v3totp: User's TOTP passcode (Env: OS_PASSCODE)
  --os-system-scope <auth-system-scope>
                        With password: Scope for system operations With
                        v3oidcauthcode: Scope for system operations With
                        v3samlpassword: Scope for system operations With
                        v3password: Scope for system operations With
                        v3adfspassword: Scope for system operations With
                        v3oidcaccesstoken: Scope for system operations With
                        v3oidcpassword: Scope for system operations With
                        token: Scope for system operations With
                        v3oidcclientcredentials: Scope for system operations
                        With v3token: Scope for system operations With v3totp:
                        Scope for system operations With
                        v3applicationcredential: Scope for system operations
                        (Env: OS_SYSTEM_SCOPE)
  --os-project-id <auth-project-id>
                        With password: Project ID to scope to With
                        v3oidcauthcode: Project ID to scope to With
                        v3samlpassword: Project ID to scope to With
                        v3password: Project ID to scope to With
                        v3adfspassword: Project ID to scope to With
                        v3oidcaccesstoken: Project ID to scope to With
                        v3oidcpassword: Project ID to scope to With token:
                        Project ID to scope to With v3oidcclientcredentials:
                        Project ID to scope to With v3tokenlessauth: Project
                        ID to scope to With v3token: Project ID to scope to
                        With v3totp: Project ID to scope to With
                        v3applicationcredential: Project ID to scope to With
                        noauth: Project ID (Env: OS_PROJECT_ID)

Commands:
  access token create  Create an access token
  acl delete     Delete ACLs for a secret or container as identified by its href. (python-barbicanclient)
  acl get        Retrieve ACLs for a secret or container by providing its href. (python-barbicanclient)
  acl submit     Submit ACL on a secret or container as identified by its href. (python-barbicanclient)
  acl user add   Add ACL users to a secret or container as identified by its href. (python-barbicanclient)
  acl user remove  Remove ACL users from a secret or container as identified by its href. (python-barbicanclient)
  address scope create  Create a new Address Scope
  address scope delete  Delete address scope(s)
  address scope list  List address scopes
  address scope set  Set address scope properties
  address scope show  Display address scope details
  aggregate add host  Add host to aggregate
  aggregate create  Create a new aggregate
  aggregate delete  Delete existing aggregate(s)
  aggregate list  List all aggregates
  aggregate remove host  Remove host from aggregate
  aggregate set  Set aggregate properties
  aggregate show  Display aggregate details
  aggregate unset  Unset aggregate properties
  availability zone list  List availability zones and their status
  bgp dragent add speaker  Add a BGP speaker to a dynamic routing agent (python-neutronclient)
  bgp dragent remove speaker  Removes a BGP speaker from a dynamic routing agent (python-neutronclient)
  bgp peer create  Create a BGP peer (python-neutronclient)
  bgp peer delete  Delete a BGP peer (python-neutronclient)
  bgp peer list  List BGP peers (python-neutronclient)
  bgp peer set   Update a BGP peer (python-neutronclient)
  bgp peer show  Show information for a BGP peer (python-neutronclient)
  bgp speaker add network  Add a network to a BGP speaker (python-neutronclient)
  bgp speaker add peer  Add a peer to a BGP speaker (python-neutronclient)
  bgp speaker create  Create a BGP speaker (python-neutronclient)
  bgp speaker delete  Delete a BGP speaker (python-neutronclient)
  bgp speaker list  List BGP speakers (python-neutronclient)
  bgp speaker list advertised routes  List routes advertised (python-neutronclient)
  bgp speaker remove network  Remove a network from a BGP speaker (python-neutronclient)
  bgp speaker remove peer  Remove a peer from a BGP speaker (python-neutronclient)
  bgp speaker set  Set BGP speaker properties (python-neutronclient)
  bgp speaker show  Show a BGP speaker (python-neutronclient)
  bgp speaker show dragents  List dynamic routing agents hosting a BGP speaker (python-neutronclient)
  bgpvpn create  Create BGP VPN resource (python-neutronclient)
  bgpvpn delete  Delete BGP VPN resource(s) (python-neutronclient)
  bgpvpn list    List BGP VPN resources (python-neutronclient)
  bgpvpn network association create  Create a BGP VPN network association (python-neutronclient)
  bgpvpn network association delete  Delete a BGP VPN network association(s) for a given BGP VPN (python-neutronclient)
  bgpvpn network association list  List BGP VPN network associations for a given BGP VPN (python-neutronclient)
  bgpvpn network association show  Show information of a given BGP VPN network association (python-neutronclient)
  bgpvpn port association create  Create a BGP VPN port association (python-neutronclient)
  bgpvpn port association delete  Delete a BGP VPN port association(s) for a given BGP VPN (python-neutronclient)
  bgpvpn port association list  List BGP VPN port associations for a given BGP VPN (python-neutronclient)
  bgpvpn port association set  Set BGP VPN port association properties (python-neutronclient)
  bgpvpn port association show  Show information of a given BGP VPN port association (python-neutronclient)
  bgpvpn port association unset  Unset BGP VPN port association properties (python-neutronclient)
  bgpvpn router association create  Create a BGP VPN router association (python-neutronclient)
  bgpvpn router association delete  Delete a BGP VPN router association(s) for a given BGP VPN (python-neutronclient)
  bgpvpn router association list  List BGP VPN router associations for a given BGP VPN (python-neutronclient)
  bgpvpn router association show  Show information of a given BGP VPN router association (python-neutronclient)
  bgpvpn set     Set BGP VPN properties (python-neutronclient)
  bgpvpn show    Show information of a given BGP VPN (python-neutronclient)
  bgpvpn unset   Unset BGP VPN properties (python-neutronclient)
  ca get         Retrieve a CA by providing its URI. (python-barbicanclient)
  ca list        List CAs. (python-barbicanclient)
  catalog list   List services in the service catalog
  catalog show   Display service catalog details
  cloud cda create  Create cloud CDA
  cloud cda delete  Delete Cloud CDA
  cloud cda list  List cloud CDA
  cloud cda set  Set Cloud CDA properties
  cloud cda show  show one cloud cda
  cloud cda snat create  Create cloud CDA snat
  cloud cda snat delete  Delete CDA SNAT
  cloud cda snat list  List CDA SNAT
  cloud cda snat show  Show CDA SNAT
  cloud cda sync  Sync cloud CDA data from ODL
  command list   List recognized commands by group
  complete       print bash completion command (cliff)
  compute agent create  Create compute agent
  compute agent delete  Delete compute agent(s)
  compute agent list  List compute agents
  compute agent set  Set compute agent properties
  compute service delete  Delete compute service(s)
  compute service list  List compute services
  compute service set  Set compute service properties
  configuration show  Display configuration details
  consistency group add volume  Add volume(s) to consistency group
  consistency group create  Create new consistency group.
  consistency group delete  Delete consistency group(s).
  consistency group list  List consistency groups.
  consistency group remove volume  Remove volume(s) from consistency group
  consistency group set  Set consistency group properties
  consistency group show  Display consistency group details.
  consistency group snapshot create  Create new consistency group snapshot.
  consistency group snapshot delete  Delete consistency group snapshot(s).
  consistency group snapshot list  List consistency group snapshots.
  consistency group snapshot show  Display consistency group snapshot details
  console log show  Show server's console output
  console url show  Show server's remote console URL
  consumer create  Create new consumer
  consumer delete  Delete consumer(s)
  consumer list  List consumers
  consumer set   Set consumer properties
  consumer show  Display consumer details
  container create  Create new container
  container delete  Delete container
  container list  List containers
  container save  Save container contents locally
  container set  Set container properties
  container show  Display container details
  container unset  Unset container properties
  credential create  Create new credential
  credential delete  Delete credential(s)
  credential list  List credentials
  credential set  Set credential properties
  credential show  Display credential details
  dedicated cloud add host  Add host to dedicated cloud
  dedicated cloud available host list  List available hosts for dedicated cloud
  dedicated cloud create  Create a new dedicated cloud
  dedicated cloud delete  Delete existing dedicated cloud(s)
  dedicated cloud host list  List all host in dedicated cloud
  dedicated cloud host set  Set host properties in dedicated cloud
  dedicated cloud host show  Display host in dedicated cloud details
  dedicated cloud host statistics  Count the resources of all hosts
  dedicated cloud list  List all dedicated clouds
  dedicated cloud remove host  Remove host from dedicated cloud
  dedicated cloud server create  create a new server in dedicated cloud
  dedicated cloud server list  List all server in dedicated cloud
  dedicated cloud server show  Display server in dedicated cloud details
  dedicated cloud set  Set dedicated cloud properties
  dedicated cloud show  Display dedicated cloud details
  dedicated host create  Create a new dedicated host
  dedicated host delete  Delete existing dedicated host
  dedicated host list  List all dedicated hosts
  dedicated host show  Display dedicated host details
  delete cancel  Cancel one or more deletion(s).
  delete confirm  Confirm one or more deletion(s).
  delete confirm resource list  Get delete confirm resources.
  delete confirm resource show  Get info for a delete confirm resource by verify id.
  detector create  Create new detector
  detector delete  Delete detector
  detector detect  Detect inside address's connection by a detector.
  detector list  List detectors
  detector set   Set detector properties
  detector show  Display detector details
  dns quota list  List quotas (python-designateclient)
  dns quota reset  Delete blacklist (python-designateclient)
  dns quota set  Set blacklist properties (python-designateclient)
  dns service list  List service statuses (python-designateclient)
  dns service show  Show service status details (python-designateclient)
  domain create  Create new domain
  domain delete  Delete domain(s)
  domain list    List domains
  domain set     Set domain properties
  domain show    Display domain details
  ec2 credentials create  Create EC2 credentials
  ec2 credentials delete  Delete EC2 credentials
  ec2 credentials list  List EC2 credentials
  ec2 credentials show  Display EC2 credentials details
  endpoint add project  Associate a project to an endpoint
  endpoint create  Create new endpoint
  endpoint delete  Delete endpoint(s)
  endpoint list  List endpoints
  endpoint remove project  Dissociate a project from an endpoint
  endpoint set   Set endpoint properties
  endpoint show  Display endpoint details
  extension list  List API extensions
  extension show  Show API extension
  federation domain list  List accessible domains
  federation project list  List accessible projects
  federation protocol create  Create new federation protocol
  federation protocol delete  Delete federation protocol(s)
  federation protocol list  List federation protocols
  federation protocol set  Set federation protocol properties
  federation protocol show  Display federation protocol details
  firewall group create  Create a new firewall group (python-neutronclient)
  firewall group delete  Delete firewall group(s) (python-neutronclient)
  firewall group list  List firewall groups (python-neutronclient)
  firewall group policy add rule  Insert a rule into a given firewall policy (python-neutronclient)
  firewall group policy create  Create a new firewall policy (python-neutronclient)
  firewall group policy delete  Delete firewall policy(s) (python-neutronclient)
  firewall group policy list  List firewall policies (python-neutronclient)
  firewall group policy remove rule  Remove a rule from a given firewall policy (python-neutronclient)
  firewall group policy set  Set firewall policy properties (python-neutronclient)
  firewall group policy show  Display firewall policy details (python-neutronclient)
  firewall group policy unset  Unset firewall policy properties (python-neutronclient)
  firewall group rule create  Create a new firewall rule (python-neutronclient)
  firewall group rule delete  Delete firewall rule(s) (python-neutronclient)
  firewall group rule list  List firewall rules that belong to a given tenant (python-neutronclient)
  firewall group rule set  Set firewall rule properties (python-neutronclient)
  firewall group rule show  Display firewall rule details (python-neutronclient)
  firewall group rule unset  Unset firewall rule properties (python-neutronclient)
  firewall group set  Set firewall group properties (python-neutronclient)
  firewall group show  Display firewall group details (python-neutronclient)
  firewall group unset  Unset firewall group properties (python-neutronclient)
  flavor create  Create new flavor
  flavor delete  Delete flavor(s)
  flavor list    List flavors
  flavor set     Set flavor properties
  flavor show    Display flavor details
  flavor unset   Unset flavor properties
  floating ip create  Create floating IP
  floating ip delete  Delete floating IP(s)
  floating ip list  List floating IP(s)
  floating ip lock  Lock floating IP
  floating ip pool list  List pools of floating IP addresses
  floating ip port forwarding create  Create floating IP port forwarding
  floating ip port forwarding delete  Delete floating IP port forwarding
  floating ip port forwarding list  List floating IP port forwarding
  floating ip port forwarding set  Set floating IP port forwarding properties
  floating ip port forwarding show  Display floating IP port forwarding details
  floating ip set  Set floating IP Properties
  floating ip show  Display floating IP details
  floating ip snat create  Create floating ip snat
  floating ip snat delete  Delete floating IP SNAT
  floating ip snat list  List floating IP SNAT
  floating ip snat set  Set floating IP SNAT Properties
  floating ip snat show  Display floating IP SNAT details
  floating ip unlock  Unlock floating IP
  floating ip unset  Unset floating IP Properties
  group add user  Add user to group
  group contains user  Check user membership in group
  group create   Create new group
  group delete   Delete group(s)
  group list     List groups
  group remove user  Remove user from group
  group set      Set group properties
  group show     Display group details
  help           print detailed help for another command (cliff)
  host list      List hosts
  host set       Set host properties
  host show      Display host details
  hypervisor list  List hypervisors
  hypervisor show  Display hypervisor details
  hypervisor stats show  Display hypervisor stats details
  identity provider create  Create new identity provider
  identity provider delete  Delete identity provider(s)
  identity provider list  List identity providers
  identity provider set  Set identity provider properties
  identity provider show  Display identity provider details
  image add project  Associate project with image
  image create   Create/upload an image
  image delete   Delete image(s)
  image list     List available images
  image quota    Show the quota information of current user
  image remove project  Disassociate project with image
  image save     Save an image locally
  image set      Set image properties
  image show     Display image details
  image unset    Unset image tags and properties
  implied role create  Creates an association between prior and implied roles
  implied role delete  Deletes an association between prior and implied roles
  implied role list  List implied roles
  ip availability list  List IP availability for network
  ip availability show  Show network IP availability details
  keypair create  Create new public or private key for server ssh access
  keypair delete  Delete public or private key(s)
  keypair list   List key fingerprints
  keypair show   Display key details
  limits show    Show compute and block storage limits
  mapping create  Create new mapping
  mapping delete  Delete mapping(s)
  mapping list   List mappings
  mapping set    Set mapping properties
  mapping show   Display mapping details
  metadata host ip create  Create new metadata host ip
  metadata host ip delete  Delete metadata host ip(s)
  metadata host ip list  List metadata host ips
  metadata host ip set  Set metadata host ip properties
  metadata host ip show  Display metadata host ip details
  module list    List module versions
  network agent add network  Add network to an agent
  network agent add router  Add router to an agent
  network agent delete  Delete network agent(s)
  network agent list  List network agents
  network agent remove network  Remove network from an agent.
  network agent remove router  Remove router from an agent
  network agent set  Set network agent properties
  network agent show  Display network agent details
  network auto allocated topology create  Create the  auto allocated topology for project
  network auto allocated topology delete  Delete auto allocated topology for project
  network create  Create new network
  network delete  Delete network(s)
  network flavor add profile  Add a service profile to a network flavor
  network flavor create  Create new network flavor
  network flavor delete  Delete network flavors
  network flavor list  List network flavors
  network flavor profile create  Create new network flavor profile
  network flavor profile delete  Delete network flavor profile
  network flavor profile list  List network flavor profile(s)
  network flavor profile set  Set network flavor profile properties
  network flavor profile show  Display network flavor profile details
  network flavor remove profile  Remove service profile from network flavor
  network flavor set  Set network flavor properties
  network flavor show  Display network flavor details
  network ip pool add address  Add an address to an IP Pool
  network ip pool create  Create network IP Pool
  network ip pool delete  Delete network IP Pool
  network ip pool list  List IP Pool
  network ip pool remove address  Remove an address from an IP Pool
  network ip pool set  Set network IP Pool
  network ip pool show  show one network IP Pool
  network ipv6 address create  Create network IPv6 address
  network ipv6 address delete  Delete network IPv6 address
  network ipv6 address list  List IPv6 address
  network ipv6 address set  Set network IPv6 address
  network ipv6 address show  show one network IPv6 address
  network list   List networks
  network log create  Create a new network log (python-neutronclient)
  network log delete  Delete network log(s) (python-neutronclient)
  network log list  List network logs (python-neutronclient)
  network log set  Set network log properties (python-neutronclient)
  network log show  Display network log details (python-neutronclient)
  network loggable resources list  List supported loggable resources (python-neutronclient)
  network meter create  Create network meter
  network meter delete  Delete network meter
  network meter list  List network meters
  network meter rule create  Create a new meter rule
  network meter rule delete  Delete meter rule(s)
  network meter rule list  List meter rules
  network meter rule show  Display meter rules details
  network meter show  Show network meter
  network qos policy create  Create a QoS policy
  network qos policy delete  Delete Qos Policy(s)
  network qos policy list  List QoS policies
  network qos policy set  Set QoS policy properties
  network qos policy show  Display QoS policy details
  network qos rule create  Create new Network QoS rule
  network qos rule delete  Delete Network QoS rule
  network qos rule list  List Network QoS rules
  network qos rule set  Set Network QoS rule properties
  network qos rule show  Display Network QoS rule details
  network qos rule type list  List QoS rule types
  network qos rule type show  Show details about supported QoS rule type
  network rbac create  Create network RBAC policy
  network rbac delete  Delete network RBAC policy(s)
  network rbac list  List network RBAC policies
  network rbac set  Set network RBAC policy properties
  network rbac show  Display network RBAC policy details
  network segment create  Create new network segment
  network segment delete  Delete network segment(s)
  network segment list  List network segments
  network segment range create  Create new network segment range
  network segment range delete  Delete network segment range(s)
  network segment range list  List network segment ranges
  network segment range set  Set network segment range properties
  network segment range show  Display network segment range details
  network segment set  Set network segment properties
  network segment show  Display network segment details
  network service provider list  List Service Providers
  network set    Set network properties
  network show   Show network details
  network subport list  List all subports for a given network trunk (python-neutronclient)
  network trunk create  Create a network trunk for a given project (python-neutronclient)
  network trunk delete  Delete a given network trunk (python-neutronclient)
  network trunk list  List all network trunks (python-neutronclient)
  network trunk set  Set network trunk properties (python-neutronclient)
  network trunk show  Show information of a given network trunk (python-neutronclient)
  network trunk unset  Unset subports from a given network trunk (python-neutronclient)
  network unset  Unset network properties
  networking path create  Create new networking path
  networking path delete  Delete networking path(s)
  networking path list  List networking paths
  networking path set  Set networking path properties
  networking path show  Display networking path details
  object create  Upload object to container
  object delete  Delete object from container
  object list    List objects
  object save    Save object locally
  object set     Set object properties
  object show    Display object details
  object store account set  Set account properties
  object store account show  Display account details
  object store account unset  Unset account properties
  object unset   Unset object properties
  policy create  Create new policy
  policy delete  Delete policy(s)
  policy list    List policies
  policy set     Set policy properties
  policy show    Display policy details
  port create    Create a new port
  port delete    Delete port(s)
  port find router  Query the router which contains the subnet of the port.
  port list      List ports
  port set       Set port properties
  port show      Display port details
  port unset     Unset port properties
  probe check address  Check inside address's connection by a router.
  project create  Create new project
  project delete  Delete project(s)
  project list   List projects
  project purge  Clean resources associated with a project
  project set    Set project properties
  project show   Display project details
  ptr record list  List floatingip ptr records (python-designateclient)
  ptr record set  Set floatingip ptr record (python-designateclient)
  ptr record show  Show floatingip ptr record details (python-designateclient)
  ptr record unset  Unset floatingip ptr record (python-designateclient)
  quota list     List quotas for all projects with non-default quota values
  quota set      Set quotas for project or class
  quota show     Show quotas for project or class
  recordset create  Create new recordset (python-designateclient)
  recordset delete  Delete recordset (python-designateclient)
  recordset list  List recordsets (python-designateclient)
  recordset set  Set recordset properties (python-designateclient)
  recordset show  Show recordset details (python-designateclient)
  region create  Create new region
  region delete  Delete region(s)
  region list    List regions
  region set     Set region properties
  region show    Display region details
  request token authorize  Authorize a request token
  request token create  Create a request token
  role add       Adds a role assignment to a user or group on a domain or project
  role assignment list  List role assignments
  role create    Create new role
  role delete    Delete role(s)
  role list      List roles
  role remove    Removes a role assignment from domain/project : user/group
  role set       Set role properties
  role show      Display role details
  router add port  Add a port to a router
  router add subnet  Add a subnet to a router
  router check address  Check inside address's connection by a router.
  router create  Create a new router
  router delete  Delete router(s)
  router list    List routers
  router remove port  Remove a port from a router
  router remove subnet  Remove a subnet from a router
  router set     Set router properties
  router show    Display router details
  router unset   Unset router properties
  secret container create  Store a container in Barbican. (python-barbicanclient)
  secret container delete  Delete a container by providing its href. (python-barbicanclient)
  secret container get  Retrieve a container by providing its URI. (python-barbicanclient)
  secret container list  List containers. (python-barbicanclient)
  secret delete  Delete a secret by providing its URI. (python-barbicanclient)
  secret get     Retrieve a secret by providing its URI. (python-barbicanclient)
  secret list    List secrets. (python-barbicanclient)
  secret order create  Create a new order. (python-barbicanclient)
  secret order delete  Delete an order by providing its href. (python-barbicanclient)
  secret order get  Retrieve an order by providing its URI. (python-barbicanclient)
  secret order list  List orders. (python-barbicanclient)
  secret store   Store a secret in Barbican. (python-barbicanclient)
  secret update  Update a secret with no payload in Barbican. (python-barbicanclient)
  security group create  Create a new security group
  security group delete  Delete security group(s)
  security group list  List security groups
  security group rule create  Create a new security group rule
  security group rule delete  Delete security group rule(s)
  security group rule list  List security group rules
  security group rule show  Display security group rule details
  security group set  Set security group properties
  security group show  Display security group details
  server add fixed ip  Add fixed IP address to server
  server add floating ip  Add floating IP address to server
  server add network  Add network to server
  server add port  Add port to server
  server add security group  Add security group to server
  server add volume  Add volume to server
  server backup create  Create a server backup image
  server create  Create a new server
  server delete  Delete server(s)
  server dump create  Create a dump file in server(s)
  server event list  List recent events of a server
  server event show  Show server event details
  server group add member  Add server to server group
  server group create  Create a new server group.
  server group delete  Delete existing server group(s).
  server group list  List all server groups.
  server group remove member  Remove server from server group
  server group show  Display server group details.
  server health check  Health check for a server
  server image create  Create a new server disk image from an existing server
  server ip set  Set server IP
  server list    List servers
  server lock    Lock server(s). A non-admin user will not be able to execute actions
  server migrate  Migrate server to different host
  server pause   Pause server(s)
  server reboot  Perform a hard or soft server reboot
  server rebuild  Rebuild server
  server remove fixed ip  Remove fixed IP address from server
  server remove floating ip  Remove floating IP address from server
  server remove network  Remove all ports of a network from server
  server remove port  Remove port from server
  server remove security group  Remove security group from server
  server remove volume  Remove volume from server
  server rescue  Put server in rescue mode
  server resize  Scale server to a new flavor.
  server restore  Restore server(s)
  server resume  Resume server(s)
  server set     Set server properties
  server set disk qos  Set the disk QoS for a server
  server shelve  Shelve server(s)
  server show    Show server details
  server ssh     SSH to server
  server start   Start server(s).
  server stop    Stop server(s).
  server suspend  Suspend server(s)
  server sync    Sync for server(s).
  server unlock  Unlock server(s)
  server unpause  Unpause server(s)
  server unrescue  Restore server from rescue mode
  server unset   Unset server properties
  server unshelve  Unshelve server(s)
  service create  Create new service
  service delete  Delete service(s)
  service list   List services
  service provider create  Create new service provider
  service provider delete  Delete service provider(s)
  service provider list  List service providers
  service provider set  Set service provider properties
  service provider show  Display service provider details
  service set    Set service properties
  service show   Display service details
  sfc flow classifier create  Create a flow classifier (python-neutronclient)
  sfc flow classifier delete  Delete a given flow classifier (python-neutronclient)
  sfc flow classifier list  List flow classifiers (python-neutronclient)
  sfc flow classifier set  Set flow classifier properties (python-neutronclient)
  sfc flow classifier show  Display flow classifier details (python-neutronclient)
  sfc port chain create  Create a port chain (python-neutronclient)
  sfc port chain delete  Delete a given port chain (python-neutronclient)
  sfc port chain list  List port chains (python-neutronclient)
  sfc port chain set  Set port chain properties (python-neutronclient)
  sfc port chain show  Display port chain details (python-neutronclient)
  sfc port chain unset  Unset port chain properties (python-neutronclient)
  sfc port pair create  Create a port pair (python-neutronclient)
  sfc port pair delete  Delete a given port pair (python-neutronclient)
  sfc port pair group create  Create a port pair group (python-neutronclient)
  sfc port pair group delete  Delete a given port pair group (python-neutronclient)
  sfc port pair group list  List port pair group (python-neutronclient)
  sfc port pair group set  Set port pair group properties (python-neutronclient)
  sfc port pair group show  Display port pair group details (python-neutronclient)
  sfc port pair group unset  Unset port pairs from port pair group (python-neutronclient)
  sfc port pair list  List port pairs (python-neutronclient)
  sfc port pair set  Set port pair properties (python-neutronclient)
  sfc port pair show  Display port pair details (python-neutronclient)
  sfc service graph create  Create a service graph. (python-neutronclient)
  sfc service graph delete  Delete a given service graph. (python-neutronclient)
  sfc service graph list  List service graphs (python-neutronclient)
  sfc service graph set  Set service graph properties (python-neutronclient)
  sfc service graph show  Show information of a given service graph. (python-neutronclient)
  snapshot create  Create new snapshot
  snapshot delete  Delete volume snapshot(s)
  snapshot list  List snapshots
  snapshot set   Set snapshot properties
  snapshot show  Display snapshot details
  snapshot unset  Unset snapshot properties
  subnet create  Create a subnet
  subnet delete  Delete subnet(s)
  subnet list    List subnets
  subnet pool create  Create subnet pool
  subnet pool delete  Delete subnet pool(s)
  subnet pool list  List subnet pools
  subnet pool set  Set subnet pool properties
  subnet pool show  Display subnet pool details
  subnet pool unset  Unset subnet pool properties
  subnet set     Set subnet properties
  subnet show    Display subnet details
  subnet unset   Unset subnet properties
  tld create     Create new tld (python-designateclient)
  tld delete     Delete tld (python-designateclient)
  tld list       List tlds (python-designateclient)
  tld set        Set tld properties (python-designateclient)
  tld show       Show tld details (python-designateclient)
  token issue    Issue new token
  token revoke   Revoke existing token
  trust create   Create new trust
  trust delete   Delete trust(s)
  trust list     List trusts
  trust show     Display trust details
  tsigkey create  Create new tsigkey (python-designateclient)
  tsigkey delete  Delete tsigkey (python-designateclient)
  tsigkey list   List tsigkeys (python-designateclient)
  tsigkey set    Set tsigkey properties (python-designateclient)
  tsigkey show   Show tsigkey details (python-designateclient)
  usage list     List resource usage per project
  usage show     Show resource usage for a single project
  user create    Create new user
  user delete    Delete user(s)
  user list      List users
  user password set  Change current user password
  user set       Set user properties
  user show      Display user details
  volume backup create  Create new volume backup
  volume backup delete  Delete volume backup(s)
  volume backup list  List volume backups
  volume backup restore  Restore volume backup
  volume backup set  Set volume backup properties
  volume backup show  Display volume backup details
  volume clean reserved time  Clean volume reserved time.
  volume create  Create new volume
  volume delete  Delete volume(s)
  volume force delete system reserved  Force delete certain system_reserved volume(s).
  volume get reserved time  Get volume reserve time.
  volume host failover  Failover volume host to different backend
  volume host set  Set volume host properties
  volume list    List volumes
  volume list system reserved  List system reserved volumes.
  volume list with reserved time  Get volumes with reserved time.
  volume lock    Lock volume(s). The volume(s) cannot be deleted until unlocked.
  volume migrate  Migrate volume to a new host
  volume qos associate  Associate a QoS specification to a volume type
  volume qos create  Create new QoS specification
  volume qos delete  Delete QoS specification
  volume qos disassociate  Disassociate a QoS specification from a volume type
  volume qos list  List QoS specifications
  volume qos set  Set QoS specification properties
  volume qos show  Display QoS specification details
  volume qos unset  Unset QoS specification properties
  volume service list  List service command
  volume service set  Set volume service properties
  volume set     Set volume properties
  volume set reserved time  Set volume reserve time.
  volume show    Display volume details
  volume snapshot create  Create new volume snapshot
  volume snapshot delete  Delete volume snapshot(s)
  volume snapshot list  List volume snapshots
  volume snapshot set  Set volume snapshot properties
  volume snapshot show  Display volume snapshot details
  volume snapshot unset  Unset volume snapshot properties
  volume system restore  Restore a system_reserved volume.
  volume transfer request accept  Accept volume transfer request.
  volume transfer request create  Create volume transfer request.
  volume transfer request delete  Delete volume transfer request(s).
  volume transfer request list  Lists all volume transfer requests.
  volume transfer request show  Show volume transfer request details.
  volume type create  Create new volume type
  volume type delete  Delete volume type(s)
  volume type list  List volume types
  volume type set  Set volume type properties
  volume type show  Display volume type details
  volume type unset  Unset volume type properties
  volume unlock  Unlock volume(s). The volume(s) can be deleted.
  volume unset   Unset volume properties
  vpn endpoint group create  Create an endpoint group (python-neutronclient)
  vpn endpoint group delete  Delete endpoint group(s) (python-neutronclient)
  vpn endpoint group list  List endpoint groups that belong to a given project (python-neutronclient)
  vpn endpoint group set  Set endpoint group properties (python-neutronclient)
  vpn endpoint group show  Display endpoint group details (python-neutronclient)
  vpn ike policy create  Create an IKE policy (python-neutronclient)
  vpn ike policy delete  Delete IKE policy (policies) (python-neutronclient)
  vpn ike policy list  List IKE policies that belong to a given project (python-neutronclient)
  vpn ike policy set  Set IKE policy properties (python-neutronclient)
  vpn ike policy show  Display IKE policy details (python-neutronclient)
  vpn ipsec policy create  Create an IPsec policy (python-neutronclient)
  vpn ipsec policy delete  Delete IPsec policy(policies) (python-neutronclient)
  vpn ipsec policy list  List IPsec policies that belong to a given project (python-neutronclient)
  vpn ipsec policy set  Set IPsec policy properties (python-neutronclient)
  vpn ipsec policy show  Display IPsec policy details (python-neutronclient)
  vpn ipsec site connection create  Create an IPsec site connection (python-neutronclient)
  vpn ipsec site connection delete  Delete IPsec site connection(s) (python-neutronclient)
  vpn ipsec site connection list  List IPsec site connections that belong to a given project (python-neutronclient)
  vpn ipsec site connection set  Set IPsec site connection properties (python-neutronclient)
  vpn ipsec site connection show  Show information of a given IPsec site connection (python-neutronclient)
  vpn service create  Create an VPN service (python-neutronclient)
  vpn service delete  Delete VPN service(s) (python-neutronclient)
  vpn service list  List VPN services that belong to a given project (python-neutronclient)
  vpn service set  Set VPN service properties (python-neutronclient)
  vpn service show  Display VPN service details (python-neutronclient)
  whitelist set  config whitelist in firewall
  whitelist show  List whitelist in firewall
  zone abandon   Abandon a zone (python-designateclient)
  zone axfr      AXFR a zone (python-designateclient)
  zone blacklist create  Create new blacklist (python-designateclient)
  zone blacklist delete  Delete blacklist (python-designateclient)
  zone blacklist list  List blacklists (python-designateclient)
  zone blacklist set  Set blacklist properties (python-designateclient)
  zone blacklist show  Show blacklist details (python-designateclient)
  zone create    Create new zone (python-designateclient)
  zone delete    Delete zone (python-designateclient)
  zone export create  Export a Zone (python-designateclient)
  zone export delete  Delete a Zone Export (python-designateclient)
  zone export list  List Zone Exports (python-designateclient)
  zone export show  Show a Zone Export (python-designateclient)
  zone export showfile  Show the zone file for the Zone Export (python-designateclient)
  zone import create  Import a Zone from a file on the filesystem (python-designateclient)
  zone import delete  Delete a Zone Import (python-designateclient)
  zone import list  List Zone Imports (python-designateclient)
  zone import show  Show a Zone Import (python-designateclient)
  zone list      List zones (python-designateclient)
  zone set       Set zone properties (python-designateclient)
  zone show      Show zone details (python-designateclient)
  zone transfer accept list  List Zone Transfer Accepts (python-designateclient)
  zone transfer accept request  Accept a Zone Transfer Request (python-designateclient)
  zone transfer accept show  Show Zone Transfer Accept (python-designateclient)
  zone transfer request create  Create new zone transfer request (python-designateclient)
  zone transfer request delete  Delete a Zone Transfer Request (python-designateclient)
  zone transfer request list  List Zone Transfer Requests (python-designateclient)
  zone transfer request set  Set a Zone Transfer Request (python-designateclient)
  zone transfer request show  Show Zone Transfer Request Details (python-designateclient)
```



## image

```
openstack image [command] 
```

### list

```bash
openstack image list 
```



```bash
usage: openstack image list [flag]

optional arguments:
  -h, --help            show this help message and exit
  --public              List only public images
  --private             List only private images
  --shared              List only shared images
  --property <key=value>
                        Filter output based on property
  --name <name>         Filter images based on name.
  --status <status>     Filter images based on status.
  --long                List additional fields in output
  --sort <key>[:<direction>]
                        Sort output by selected keys and directions(asc or
                        desc) (default: name:asc), multiple keys and
                        directions can be specified separated by comma
  --limit <num-images>  Maximum number of images to display.
  --marker <image>      The last image of the previous page. Display list of
                        images after marker. Display all images if not
                        specified. (name or ID)

output formatters:
  output formatter options

  -f {csv,json,table,value,yaml}, --format {csv,json,table,value,yaml}
                        the output format, defaults to table
  -c COLUMN, --column COLUMN
                        specify the column(s) to include, can be repeated
  --sort-column SORT_COLUMN
                        specify the column(s) to sort the data (columns
                        specified first have a priority, non-existing columns
                        are ignored), can be repeated

table formatter:
  --max-width <integer>
                        Maximum display width, <1 to disable. You can also use
                        the CLIFF_MAX_TERM_WIDTH environment variable, but the
                        parameter takes precedence.
  --fit-width           Fit the table to the display width. Implied if --max-
                        width greater than 0. Set the environment variable
                        CLIFF_FIT_WIDTH=1 to always enable
  --print-empty         Print empty table if there is no data to show.

json formatter:
  --noindent            whether to disable indenting the JSON

CSV Formatter:
  --quote {all,minimal,none,nonnumeric}
                        when to include quotes, defaults to nonnumeric

```



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



## flavor

```bash
openstack flavor [command]

commands:
	create
 	delete
	list
	set
	show
	unset
```

### list

```bash
openstack flavor list --limit 10 --public
# | 214d0281-6fb5-48ae-9ca7-64f2a11b5db6 | s6.small.1          |   1024 |    0 |         0 |     1 | True  
```

```bash
[root@gz-txjs-control-55e243e31e30 ~]# openstack flavor list --help
usage: openstack flavor list [-h] [-f {csv,json,table,value,yaml}] [-c COLUMN]
                             [--max-width <integer>] [--fit-width]
                             [--print-empty] [--noindent]
                             [--quote {all,minimal,none,nonnumeric}]
                             [--sort-column SORT_COLUMN]
                             [--public | --private | --all] [--long]
                             [--marker <flavor-id>] [--limit <num-flavors>]

List flavors

optional arguments:
  -h, --help            show this help message and exit
  --public              List only public flavors (default)
  --private             List only private flavors
  --all                 List all flavors, whether public or private
  --long                List additional fields in output
  --marker <flavor-id>  The last flavor ID of the previous page
  --limit <num-flavors>
                        Maximum number of flavors to display

output formatters:
  output formatter options

  -f {csv,json,table,value,yaml}, --format {csv,json,table,value,yaml}
                        the output format, defaults to table
  -c COLUMN, --column COLUMN
                        specify the column(s) to include, can be repeated
  --sort-column SORT_COLUMN
                        specify the column(s) to sort the data (columns
                        specified first have a priority, non-existing columns
                        are ignored), can be repeated

table formatter:
  --max-width <integer>
                        Maximum display width, <1 to disable. You can also use
                        the CLIFF_MAX_TERM_WIDTH environment variable, but the
                        parameter takes precedence.
  --fit-width           Fit the table to the display width. Implied if --max-
                        width greater than 0. Set the environment variable
                        CLIFF_FIT_WIDTH=1 to always enable
  --print-empty         Print empty table if there is no data to show.

json formatter:
  --noindent            whether to disable indenting the JSON

CSV Formatter:
  --quote {all,minimal,none,nonnumeric}
                        when to include quotes, defaults to nonnumeric
```

## network

```bsh
+--------------------------------------+-----------------------------+--------------------------------------+
[root@gz-txjs-control-55e243e31e30 ~]# openstack network --help
Command "network" matches:
  network agent add network
  network agent add router
  network agent delete
  network agent list
  network agent remove network
  network agent remove router
  network agent set
  network agent show
  network auto allocated topology create
  network auto allocated topology delete
  network create
  network delete
  network flavor add profile
  network flavor create
  network flavor delete
  network flavor list
  network flavor profile create
  network flavor profile delete
  network flavor profile list
  network flavor profile set
  network flavor profile show
  network flavor remove profile
  network flavor set
  network flavor show
  network ip pool add address
  network ip pool create
  network ip pool delete
  network ip pool list
  network ip pool remove address
  network ip pool set
  network ip pool show
  network ipv6 address create
  network ipv6 address delete
  network ipv6 address list
  network ipv6 address set
  network ipv6 address show
  network list
  network log create
  network log delete
  network log list
  network log set
  network log show
  network loggable resources list
  network meter create
  network meter delete
  network meter list
  network meter rule create
  network meter rule delete
  network meter rule list
  network meter rule show
  network meter show
  network qos policy create
  network qos policy delete
  network qos policy list
  network qos policy set
  network qos policy show
  network qos rule create
  network qos rule delete
  network qos rule list
  network qos rule set
  network qos rule show
  network qos rule type list
  network qos rule type show
  network rbac create
  network rbac delete
  network rbac list
  network rbac set
  network rbac show
  network segment create
  network segment delete
  network segment list
  network segment range create
  network segment range delete
  network segment range list
  network segment range set
  network segment range show
  network segment set
  network segment show
  network service provider list
  network set
  network show
  network subport list
  network trunk create
  network trunk delete
  network trunk list
  network trunk set
  network trunk show
  network trunk unset
  network unset
  networking path create
  networking path delete
  networking path list
  networking path set
  networking path show
```



### list

```bash
 openstack network list  --enable --status ACTIVE --limit 10
 
 # 010c543a-f741-45f5-892c-ad84f710d48b | vpc-6e72           | 3b78cddf-0a70-4b87-a5d5-e6e62c7c8520 
```



```bash
usage: openstack network list [-h] [-f {csv,json,table,value,yaml}]
                              [-c COLUMN] [--max-width <integer>]
                              [--fit-width] [--print-empty] [--noindent]
                              [--quote {all,minimal,none,nonnumeric}]
                              [--sort-column SORT_COLUMN]
                              [--external | --internal] [--long]
                              [--name <name>] [--enable | --disable]
                              [--project <project>]
                              [--project-domain <project-domain>]
                              [--share | --no-share] [--status <status>]
                              [--provider-network-type <provider-network-type>]
                              [--provider-physical-network <provider-physical-network>]
                              [--provider-segment <provider-segment>]
                              [--agent <agent-id>] [--probe]
                              [--marker <network-id>] [--limit <num-networks>]
                              [--tags <tag>[,<tag>,...]]
                              [--any-tags <tag>[,<tag>,...]]
                              [--not-tags <tag>[,<tag>,...]]
                              [--not-any-tags <tag>[,<tag>,...]]

List networks

optional arguments:
  -h, --help            show this help message and exit
  --external            List external networks
  --internal            List internal networks
  --long                List additional fields in output
  --name <name>         List networks according to their name
  --enable              List enabled networks
  --disable             List disabled networks
  --project <project>   List networks according to their project (name or ID)
  --project-domain <project-domain>
                        Domain the project belongs to (name or ID). This can
                        be used in case collisions between project names
                        exist.
  --share               List networks shared between projects
  --no-share            List networks not shared between projects
  --status <status>     List networks according to their status ('ACTIVE',
                        'BUILD', 'DOWN', 'ERROR')
  --provider-network-type <provider-network-type>
                        List networks according to their physical mechanisms.
                        The supported options are: flat, geneve, gre, local,
                        vlan, vxlan.
  --provider-physical-network <provider-physical-network>
                        List networks according to name of the physical
                        network
  --provider-segment <provider-segment>
                        List networks according to VLAN ID for VLAN networks
                        or Tunnel ID for GENEVE/GRE/VXLAN networks
  --agent <agent-id>    List networks hosted by agent (ID only)
  --probe               List networks hosted by probe agent (ID only).NOTE:
                        The argument depend on --agent.
  --marker <network-id>
                        The last network ID of the previous page. Display list
                        of networks after marker. Display all networks if not
                        specified. (Only valid when --limit is specified)
  --limit <num-networks>
                        Maximum number of networks to display. Limit should be
                        greater than 0, If limit is greater than
                        'pagination_max_limit' option of Neutron API,
                        'pagination_max_limit' will be used instead.
  --tags <tag>[,<tag>,...]
                        List networks which have all given tag(s) (Comma-
                        separated list of tags)
  --any-tags <tag>[,<tag>,...]
                        List networks which have any given tag(s) (Comma-
                        separated list of tags)
  --not-tags <tag>[,<tag>,...]
                        Exclude networks which have all given tag(s) (Comma-
                        separated list of tags)
  --not-any-tags <tag>[,<tag>,...]
                        Exclude networks which have any given tag(s) (Comma-
                        separated list of tags)

output formatters:
  output formatter options

  -f {csv,json,table,value,yaml}, --format {csv,json,table,value,yaml}
                        the output format, defaults to table
  -c COLUMN, --column COLUMN
                        specify the column(s) to include, can be repeated
  --sort-column SORT_COLUMN
                        specify the column(s) to sort the data (columns
                        specified first have a priority, non-existing columns
                        are ignored), can be repeated

table formatter:
  --max-width <integer>
                        Maximum display width, <1 to disable. You can also use
                        the CLIFF_MAX_TERM_WIDTH environment variable, but the
                        parameter takes precedence.
  --fit-width           Fit the table to the display width. Implied if --max-
                        width greater than 0. Set the environment variable
                        CLIFF_FIT_WIDTH=1 to always enable
  --print-empty         Print empty table if there is no data to show.

json formatter:
  --noindent            whether to disable indenting the JSON

CSV Formatter:
  --quote {all,minimal,none,nonnumeric}
                        when to include quotes, defaults to nonnumeric
```

## server 

```bash
server add fixed ip  Add fixed IP address to server
server add floating ip  Add floating IP address to server
server add network  Add network to server
server add port  Add port to server
server add security group  Add security group to server
server add volume  Add volume to server
server backup create  Create a server backup image
server create  Create a new server
server delete  Delete server(s)
server dump create  Create a dump file in server(s)
server event list  List recent events of a server
server event show  Show server event details
server group add member  Add server to server group
server group create  Create a new server group.
server group delete  Delete existing server group(s).
server group list  List all server groups.
server group remove member  Remove server from server group
server group show  Display server group details.
server health check  Health check for a server
server image create  Create a new server disk image from an existing server
server ip set  Set server IP
server list    List servers
server lock    Lock server(s). A non-admin user will not be able to execute actions
server migrate  Migrate server to different host
server pause   Pause server(s)
server reboot  Perform a hard or soft server reboot
server rebuild  Rebuild server
server remove fixed ip  Remove fixed IP address from server
server remove floating ip  Remove floating IP address from server
server remove network  Remove all ports of a network from server
server remove port  Remove port from server
server remove security group  Remove security group from server
server remove volume  Remove volume from server
server rescue  Put server in rescue mode
server resize  Scale server to a new flavor.
server restore  Restore server(s)
server resume  Resume server(s)
server set     Set server properties
server set disk qos  Set the disk QoS for a server
server shelve  Shelve server(s)
server show    Show server details
server ssh     SSH to server
server start   Start server(s).
server stop    Stop server(s).
server suspend  Suspend server(s)
server sync    Sync for server(s).
server unlock  Unlock server(s)
server unpause  Unpause server(s)
server unrescue  Restore server from rescue mode
server unset   Unset server properties
server unshelve  Unshelve server(s)
service create  Create new service
service delete  Delete service(s)
service list   List services
service provider create  Create new service provider
service provider delete  Delete service provider(s)
service provider list  List service providers
service provider set  Set service provider properties
service provider show  Display service provider details
service set    Set service properties
service show   Display service details

[root@gz-txjs-control-55e243e31e30 ~]# openstack server create --help
usage: openstack server create [-h] [-f {json,shell,table,value,yaml}]
                               [-c COLUMN] [--max-width <integer>]
                               [--fit-width] [--print-empty] [--noindent]
                               [--prefix PREFIX] [--image <image>]
                               [--volume <volume>] --flavor <flavor>
                               [--security-group <security-group>]
                               [--key-name <key-name>]
                               [--property <key=value>]
                               [--file <dest-filename=source-filename>]
                               [--user-data <user-data>]
                               [--availability-zone <zone-name>]
                               [--block-device-mapping <dev-name=mapping>]
                               [--nic <net-id=net-uuid,v4-fixed-ip=ip-addr,v6-fixed-ip=ip-addr,port-id=port-uuid,auto,none>]
                               [--network <network>] [--port <port>]
                               [--hint <key=value>]
                               [--config-drive <config-drive-volume>|True]
                               [--min <count>] [--max <count>] [--wait]
                               [--instance-snapshot <instance_snapshot>]
                               <server-name>

Create a new server

positional arguments:
  <server-name>         New server name

optional arguments:
  -h, --help            show this help message and exit
  --image <image>       Create server boot disk from this image (name or ID)
  --volume <volume>     Create server using this volume as the boot disk (name
                        or ID).
                        This option automatically creates a block device
                        mapping with a boot index of 0. On many hypervisors
                        (libvirt/kvm for example) this will be device vda. Do
                        not create a duplicate mapping using --block-device-
                        mapping for this volume.
  --flavor <flavor>     Create server with this flavor (name or ID)
  --security-group <security-group>
                        Security group to assign to this server (name or ID)
                        (repeat option to set multiple groups)
  --key-name <key-name>
                        Keypair to inject into this server (optional
                        extension)
  --property <key=value>
                        Set a property on this server (repeat option to set
                        multiple values)
  --file <dest-filename=source-filename>
                        File to inject into image before boot (repeat option
                        to set multiple files)
  --user-data <user-data>
                        User data file to serve from the metadata server
  --availability-zone <zone-name>
                        Select an availability zone for the server
  --block-device-mapping <dev-name=mapping>
                        Create a block device on the server.
                        Block device mapping in the format
                        <dev-name>=<id>:<type>:<size(GB)>:<delete-on-
                        terminate>
                        <dev-name>: block device name, like: vdb, xvdc
                        (required)
                        <id>: UUID of the volume or snapshot (required)
                        <type>: volume or snapshot; default: volume (optional)
                        <size(GB)>: volume size if create from snapshot
                        (optional)
                        <delete-on-terminate>: true or false; default: false
                        (optional)
                        (optional extension)
  --nic <net-id=net-uuid,v4-fixed-ip=ip-addr,v6-fixed-ip=ip-addr,port-id=port-uuid,auto,none>
                        Create a NIC on the server. Specify option multiple
                        times to create multiple NICs. Either net-id or port-
                        id must be provided, but not both. net-id: attach NIC
                        to network with this UUID, port-id: attach NIC to port
                        with this UUID, v4-fixed-ip: IPv4 fixed address for
                        NIC (optional), v6-fixed-ip: IPv6 fixed address for
                        NIC (optional), none: (v2.37+) no network is attached,
                        auto: (v2.37+) the compute service will automatically
                        allocate a network. Specifying a --nic of auto or none
                        cannot be used with any other --nic value.
  --network <network>   Create a NIC on the server and connect it to network.
                        Specify option multiple times to create multiple NICs.
                        This is a wrapper for the '--nic net-id=<network>'
                        parameter that provides simple syntax for the standard
                        use case of connecting a new server to a given
                        network. For more advanced use cases, refer to the '--
                        nic' parameter.
  --port <port>         Create a NIC on the server and connect it to port.
                        Specify option multiple times to create multiple NICs.
                        This is a wrapper for the '--nic port-id=<pord>'
                        parameter that provides simple syntax for the standard
                        use case of connecting a new server to a given port.
                        For more advanced use cases, refer to the '--nic'
                        parameter.
  --hint <key=value>    Hints for the scheduler (optional extension)
  --config-drive <config-drive-volume>|True
                        Use specified volume as the config drive, or 'True' to
                        use an ephemeral drive
  --min <count>         Minimum number of servers to launch (default=1)
  --max <count>         Maximum number of servers to launch (default=1)
  --wait                Wait for build to complete
  --instance-snapshot <instance_snapshot>
                        ID of instance snapshot to create instance, cannot use
                        with other block_device or image.

output formatters:
  output formatter options

  -f {json,shell,table,value,yaml}, --format {json,shell,table,value,yaml}
                        the output format, defaults to table
  -c COLUMN, --column COLUMN
                        specify the column(s) to include, can be repeated

table formatter:
  --max-width <integer>
                        Maximum display width, <1 to disable. You can also use
                        the CLIFF_MAX_TERM_WIDTH environment variable, but the
                        parameter takes precedence.
  --fit-width           Fit the table to the display width. Implied if --max-
                        width greater than 0. Set the environment variable
                        CLIFF_FIT_WIDTH=1 to always enable
  --print-empty         Print empty table if there is no data to show.

json formatter:
  --noindent            whether to disable indenting the JSON

shell formatter:
  a format a UNIX shell can parse (variable="value")

  --prefix PREFIX       add a prefix to all variable names
```



### create

```bash
openstack server create \
--flavor 214d0281-6fb5-48ae-9ca7-64f2a11b5db6 \
--image 24647239-7b2e-4895-8ede-1d278a3b10df \
--nic net-id=010c543a-f741-45f5-892c-ad84f710d48b \
--availability-zone  S6-PUBLIC-ZONE \
janus

```

```bash
usage: openstack server create [-h] [-f {json,shell,table,value,yaml}]
                               [-c COLUMN] [--max-width <integer>]
                               [--fit-width] [--print-empty] [--noindent]
                               [--prefix PREFIX] [--image <image>]
                               [--volume <volume>] --flavor <flavor>
                               [--security-group <security-group>]
                               [--key-name <key-name>]
                               [--property <key=value>]
                               [--file <dest-filename=source-filename>]
                               [--user-data <user-data>]
                               [--availability-zone <zone-name>]
                               [--block-device-mapping <dev-name=mapping>]
                               [--nic <net-id=net-uuid,v4-fixed-ip=ip-addr,v6-fixed-ip=ip-addr,port-id=port-uuid,auto,none>]
                               [--network <network>] [--port <port>]
                               [--hint <key=value>]
                               [--config-drive <config-drive-volume>|True]
                               [--min <count>] [--max <count>] [--wait]
                               [--instance-snapshot <instance_snapshot>]
                               <server-name>

Create a new server

positional arguments:
  <server-name>         New server name

optional arguments:
  -h, --help            show this help message and exit
  --image <image>       Create server boot disk from this image (name or ID)
  --volume <volume>     Create server using this volume as the boot disk (name
                        or ID).
                        This option automatically creates a block device
                        mapping with a boot index of 0. On many hypervisors
                        (libvirt/kvm for example) this will be device vda. Do
                        not create a duplicate mapping using --block-device-
                        mapping for this volume.
  --flavor <flavor>     Create server with this flavor (name or ID)
  --security-group <security-group>
                        Security group to assign to this server (name or ID)
                        (repeat option to set multiple groups)
  --key-name <key-name>
                        Keypair to inject into this server (optional
                        extension)
  --property <key=value>
                        Set a property on this server (repeat option to set
                        multiple values)
  --file <dest-filename=source-filename>
                        File to inject into image before boot (repeat option
                        to set multiple files)
  --user-data <user-data>
                        User data file to serve from the metadata server
  --availability-zone <zone-name>
                        Select an availability zone for the server
  --block-device-mapping <dev-name=mapping>
                        Create a block device on the server.
                        Block device mapping in the format
                        <dev-name>=<id>:<type>:<size(GB)>:<delete-on-
                        terminate>
                        <dev-name>: block device name, like: vdb, xvdc
                        (required)
                        <id>: UUID of the volume or snapshot (required)
                        <type>: volume or snapshot; default: volume (optional)
                        <size(GB)>: volume size if create from snapshot
                        (optional)
                        <delete-on-terminate>: true or false; default: false
                        (optional)
                        (optional extension)
  --nic <net-id=net-uuid,v4-fixed-ip=ip-addr,v6-fixed-ip=ip-addr,port-id=port-uuid,auto,none>
                        Create a NIC on the server. Specify option multiple
                        times to create multiple NICs. Either net-id or port-
                        id must be provided, but not both. net-id: attach NIC
                        to network with this UUID, port-id: attach NIC to port
                        with this UUID, v4-fixed-ip: IPv4 fixed address for
                        NIC (optional), v6-fixed-ip: IPv6 fixed address for
                        NIC (optional), none: (v2.37+) no network is attached,
                        auto: (v2.37+) the compute service will automatically
                        allocate a network. Specifying a --nic of auto or none
                        cannot be used with any other --nic value.
  --network <network>   Create a NIC on the server and connect it to network.
                        Specify option multiple times to create multiple NICs.
                        This is a wrapper for the '--nic net-id=<network>'
                        parameter that provides simple syntax for the standard
                        use case of connecting a new server to a given
                        network. For more advanced use cases, refer to the '--
                        nic' parameter.
  --port <port>         Create a NIC on the server and connect it to port.
                        Specify option multiple times to create multiple NICs.
                        This is a wrapper for the '--nic port-id=<pord>'
                        parameter that provides simple syntax for the standard
                        use case of connecting a new server to a given port.
                        For more advanced use cases, refer to the '--nic'
                        parameter.
  --hint <key=value>    Hints for the scheduler (optional extension)
  --config-drive <config-drive-volume>|True
                        Use specified volume as the config drive, or 'True' to
                        use an ephemeral drive
  --min <count>         Minimum number of servers to launch (default=1)
  --max <count>         Maximum number of servers to launch (default=1)
  --wait                Wait for build to complete
  --instance-snapshot <instance_snapshot>
                        ID of instance snapshot to create instance, cannot use
                        with other block_device or image.

output formatters:
  output formatter options

  -f {json,shell,table,value,yaml}, --format {json,shell,table,value,yaml}
                        the output format, defaults to table
  -c COLUMN, --column COLUMN
                        specify the column(s) to include, can be repeated

table formatter:
  --max-width <integer>
                        Maximum display width, <1 to disable. You can also use
                        the CLIFF_MAX_TERM_WIDTH environment variable, but the
                        parameter takes precedence.
  --fit-width           Fit the table to the display width. Implied if --max-
                        width greater than 0. Set the environment variable
                        CLIFF_FIT_WIDTH=1 to always enable
  --print-empty         Print empty table if there is no data to show.

json formatter:
  --noindent            whether to disable indenting the JSON

shell formatter:
  a format a UNIX shell can parse (variable="value")

  --prefix PREFIX       add a prefix to all variable names
```

### list

```bash
openstack server list --name janus
```

```bash
openstack server list [flag]

optional arguments:
  -h, --help            show this help message and exit
  --reservation-id <reservation-id> Only return instances that match the reservation
  --ip <ip-address-regex>  Regular expression to match IP addresses
  --ip6 <ip-address-regex>Regular expression to match IPv6 addresses
  --name <name-regex>   Regular expression to match names
  --instance-name <server-name> Regular expression to match instance name (admin only)
  --status <status>     Search by server status
  --flavor <flavor>     Search by flavor (name or ID)
  --image <image>       Search by image (name or ID)
  --host <hostname>     Search by hostname
  --all-projects        Include all projects (admin only)
  --project <project>   Search by project (admin only) (name or ID)
  --project-domain <project-domain>  
  					Domain the project belongs to (name or ID). This can
                        be used in case collisions between project names
                        exist.
  --user <user>         Search by user (admin only) (name or ID)
  --user-domain <user-domain>
                        Domain the user belongs to (name or ID). This can be
                        used in case collisions between user names exist.
  --long                List additional fields in output
  -n, --no-name-lookup  Skip flavor and image name lookup.
  --marker <server>     The last server of the previous page. Display list of
                        servers after marker. Display all servers if not
                        specified. (name or ID)
  --limit <num-servers>
                        Maximum number of servers to display. If limit equals
                        -1, all servers will be displayed. If limit is greater
                        than 'osapi_max_limit' option of Nova API,
                        'osapi_max_limit' will be used instead.
  --deleted             Only display deleted servers (Admin only).
  --changes-since <changes-since>
                        List only servers changed after a certain point of
                        time. The provided time should be an ISO 8061
                        formatted time. ex 2016-03-04T06:27:59Z .

output formatters:
  output formatter options

  -f {csv,json,table,value,yaml}, --format {csv,json,table,value,yaml}
                        the output format, defaults to table
  -c COLUMN, --column COLUMN
                        specify the column(s) to include, can be repeated
  --sort-column SORT_COLUMN
                        specify the column(s) to sort the data (columns
                        specified first have a priority, non-existing columns
                        are ignored), can be repeated

table formatter:
  --max-width <integer>
                        Maximum display width, <1 to disable. You can also use
                        the CLIFF_MAX_TERM_WIDTH environment variable, but the
                        parameter takes precedence.
  --fit-width           Fit the table to the display width. Implied if --max-
                        width greater than 0. Set the environment variable
                        CLIFF_FIT_WIDTH=1 to always enable
  --print-empty         Print empty table if there is no data to show.

json formatter:
  --noindent            whether to disable indenting the JSON

CSV Formatter:
  --quote {all,minimal,none,nonnumeric}
                        when to include quotes, defaults to nonnumeric
```

### delete

```bash
usage: openstack server delete [flag] <server name or id>

optional arguments:
  -h, --help  show this help message and exit
  --wait      Wait for delete to complete
```

### resize

```bash

```

```bash
[root@gz-txjs-control-55e243e31e29 ~]# openstack help server resize
usage: openstack server resize [-h] [--flavor <flavor> | --confirm | --revert]
                               [--wait]
                               <server>

Scale server to a new flavor. A resize operation is implemented by creating a
new server and copying the contents of the original disk into a new one. It is
also a two-step process for the user: the first is to perform the resize, the
second is to either confirm (verify) success and release the old server, or to
declare a revert to release the new server and restart the old one.

positional arguments:
  <server>           Server (name or ID)

optional arguments:
  -h, --help         show this help message and exit
  --flavor <flavor>  Resize server to specified flavor
  --confirm          Confirm server resize is complete
  --revert           Restore server state before resize
  --wait             Wait for resize to complete
```

### stop

### set

```bash
[root@gz-txjs-control-55e243e31e29 ~]# openstack help server set
usage: openstack server set [-h] [--name <new-name>] [--root-password]
                            [--property <key=value>] [--state <state>]
                            <server>

Set server properties

positional arguments:
  <server>              Server (name or ID)

optional arguments:
  -h, --help            show this help message and exit
  --name <new-name>     New server name
  --root-password       Set new root password (interactive only)
  --property <key=value>
                        Property to add/change for this server (repeat option
                        to set multiple properties)
  --state <state>       New server state (valid value: active, error)

```



## host

```bash
openstack host [command]

commands:
	list 
	set 
	show
```

### list

```bash
usage: openstack host list [flags]

List hosts

optional arguments:
  -h, --help            show this help message and exit
  --zone <zone>         Only return hosts in the availability zone

output formatters:
  output formatter options

  -f {csv,json,table,value,yaml}, --format {csv,json,table,value,yaml}
                        the output format, defaults to table
  -c COLUMN, --column COLUMN
                        specify the column(s) to include, can be repeated
  --sort-column SORT_COLUMN
                        specify the column(s) to sort the data (columns
                        specified first have a priority, non-existing columns
                        are ignored), can be repeated

table formatter:
  --max-width <integer>
                        Maximum display width, <1 to disable. You can also use
                        the CLIFF_MAX_TERM_WIDTH environment variable, but the
                        parameter takes precedence.
  --fit-width           Fit the table to the display width. Implied if --max-
                        width greater than 0. Set the environment variable
                        CLIFF_FIT_WIDTH=1 to always enable
  --print-empty         Print empty table if there is no data to show.

json formatter:
  --noindent            whether to disable indenting the JSON

CSV Formatter:
  --quote {all,minimal,none,nonnumeric}
                        when to include quotes, defaults to nonnumeric
```



## availability zone list 

```bash
[root@gz-txjs-control-55e243e31e29 ~]# openstack help availability zone list 
usage: openstack availability zone list [-h] [-f {csv,json,table,value,yaml}]
                                        [-c COLUMN] [--max-width <integer>]
                                        [--fit-width] [--print-empty]
                                        [--noindent]
                                        [--quote {all,minimal,none,nonnumeric}]
                                        [--sort-column SORT_COLUMN]
                                        [--compute] [--network] [--volume]
                                        [--long]

List availability zones and their status

optional arguments:
  -h, --help            show this help message and exit
  --compute             List compute availability zones
  --network             List network availability zones
  --volume              List volume availability zones
  --long                List additional fields in output

output formatters:
  output formatter options

  -f {csv,json,table,value,yaml}, --format {csv,json,table,value,yaml}
                        the output format, defaults to table
  -c COLUMN, --column COLUMN
                        specify the column(s) to include, can be repeated
  --sort-column SORT_COLUMN
                        specify the column(s) to sort the data (columns
                        specified first have a priority, non-existing columns
                        are ignored), can be repeated

table formatter:
  --max-width <integer>
                        Maximum display width, <1 to disable. You can also use
                        the CLIFF_MAX_TERM_WIDTH environment variable, but the
                        parameter takes precedence.
  --fit-width           Fit the table to the display width. Implied if --max-
                        width greater than 0. Set the environment variable
                        CLIFF_FIT_WIDTH=1 to always enable
  --print-empty         Print empty table if there is no data to show.

json formatter:
  --noindent            whether to disable indenting the JSON

CSV Formatter:
  --quote {all,minimal,none,nonnumeric}
                        when to include quotes, defaults to nonnumeric
[root@gz-txjs-control-55e243e31e29 ~]# openstack help availability zone list  -long
usage: openstack help [-h] [cmd [cmd ...]]
openstack help: error: unrecognized arguments: -long

```



## server migrate

```bash
usage: openstack server migrate [-h] [--live <hostname>]
                                [--shared-migration | --block-migration]
                                [--disk-overcommit | --no-disk-overcommit]
                                [--wait] [--migration-type <migration-type>]
                                <server>

Migrate server to different host

positional arguments:
  <server>              Server (name or ID)

optional arguments:
  -h, --help            show this help message and exit
  --live <hostname>     Target hostname
  --shared-migration    Perform a shared live migration (default)
  --block-migration     Perform a block live migration
  --disk-overcommit     Allow disk over-commit on the destination host
  --no-disk-overcommit  Do not over-commit disk on the destination host
                        (default)
  --wait                Wait for migrate to complete
  --migration-type <migration-type>
                        For local storage migration to shared storage. e.g.
                        to_<volume_type_name>(see 'cinder type-list'): migrate
                        all the disks of the instance to the volume.
```



## hypervisor

### list

```bash
openstack hypervisor list --long
```

```bash
usage: openstack hypervisor list [-h] [-f {csv,json,table,value,yaml}]
                                 [-c COLUMN] [--max-width <integer>]
                                 [--fit-width] [--print-empty] [--noindent]
                                 [--quote {all,minimal,none,nonnumeric}]
                                 [--sort-column SORT_COLUMN]
                                 [--matching <hostname>] [--long]

List hypervisors

optional arguments:
  -h, --help            show this help message and exit
  --matching <hostname>
                        Filter hypervisors using <hostname> substring
  --long                List additional fields in output

output formatters:
  output formatter options

  -f {csv,json,table,value,yaml}, --format {csv,json,table,value,yaml}
                        the output format, defaults to table
  -c COLUMN, --column COLUMN
                        specify the column(s) to include, can be repeated
  --sort-column SORT_COLUMN
                        specify the column(s) to sort the data (columns
                        specified first have a priority, non-existing columns
                        are ignored), can be repeated

table formatter:
  --max-width <integer>
                        Maximum display width, <1 to disable. You can also use
                        the CLIFF_MAX_TERM_WIDTH environment variable, but the
                        parameter takes precedence.
  --fit-width           Fit the table to the display width. Implied if --max-
                        width greater than 0. Set the environment variable
                        CLIFF_FIT_WIDTH=1 to always enable
  --print-empty         Print empty table if there is no data to show.

json formatter:
  --noindent            whether to disable indenting the JSON

CSV Formatter:
  --quote {all,minimal,none,nonnumeric}
                        when to include quotes, defaults to nonnumeric
```



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
glace help [command]
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

## image-list

```bash
glance image-list --visibility public --limie 10

# | 24647239-7b2e-4895-8ede-1d278a3b10df | Ubuntu_16.04  
```

```bash
usage: glance image-list [flag]

Optional arguments:
  --limit <int>       Maximum number of images to get.
  --page-size <SIZE>    Number of images to request in each paginated request.
  --visibility <string>	可见性: public private
  --member-status <MEMBER_STATUS>  The status of images to display.
  --owner <OWNER>       Display images owned by <OWNER>.
  --property-filter <KEY=VALUE>  Filter images by a user-defined image property.
  --checksum <CHECKSUM> Displays images that match the checksum.
  --tag <TAG>           Filter images by a user-defined tag.
  --sort-key {name,status,container_format,disk_format,size,id,created_at,updated_at}
                        Sort image list by specified fields. May be used
                        multiple times.
  --sort-dir {asc,desc} Sort image list in specified directions.
  --sort <key>[:<direction>]
                        Comma-separated list of sort keys and directions in
                        the form of <key>[:<asc|desc>]. Valid keys: name,
                        status, container_format, disk_format, size, id,
                        created_at, updated_at. OPTIONAL.
```

## image-show

```bash
glance image-show 24647239-7b2e-4895-8ede-1d278a3b10df --human-readable
```

```bash
glance image-show [flag] <IMAGE_ID>

Optional arguments:
  --human-readable      Print image size in a human-friendly format.
  --max-column-width <integer> The max column width of the printed table.
```



# nova

## 项目结构

```bash
nova
├── api-guide        # 包含 Nova API 的使用指南文档。
├── api-ref          # 包含 Nova API 参考文档，通常为自动生成的 REST API 文档。
├── contrib          # 包含贡献者添加的代码或模块，可能是一些实验性功能或是社区贡献的功能。
├── devstack         # 包含用于快速设置开发环境的脚本和配置文件，通常是通过 DevStack 工具。
├── doc              # 包含项目的各种文档，可能是开发者文档、设计文档等。
├── etc              # 包含配置文件模板，这些文件是运行 Nova 所必需的。
├── gate             # 包含用于执行自动化测试的脚本和配置，特别是与 CI/CD 流程相关的部分。
├── nova             # 包含 Nova 的主要源代码，这是项目的主代码库。
│   ├── api                # 包含 API 的实现代码，处理来自用户的 HTTP 请求。
│		├── cells              # 实现 Cells v2 功能的代码，Cells 是一种用于水平扩展 Nova 的机制。
│   ├── cmd                # 包含启动不同 Nova 组件（如 nova-api, nova-compute）的命令行脚本。
│   ├── common             # 包含多个模块共享的通用代码或工具函数。
│   ├── compute            # 包含与计算节点相关的逻辑，负责虚拟机的生命周期管理。
│   ├── conductor          # 包含 Conductor 服务的实现，Conductor 负责协调多个计算节点之间的任务。
│   ├── conf               # 包含配置文件模板或配置项定义，用于指导 Nova 的行为。
│   ├── console            # 包含控制台代理服务的代码，提供用户通过 VNC 或其他协议访问实例的能力。
│   ├── consoleauth        # 包含控制台认证服务的代码，验证用户对控制台访问的权限。
│   ├── db                 # 包含数据库访问层的代码，负责与 Nova 使用的数据库进行交互。
│   ├── hacking            # 包含开发指南、代码风格检查规则等，帮助贡献者遵循项目规范。
│   ├── image              # 包含与镜像服务（Glance）交互的代码，用于管理虚拟机的镜像。
│   ├── ipv6               # 包含与 IPv6 地址管理有关的代码。
│   ├── keymgr             # 包含密钥管理服务的接口，用于加密相关操作。
│   ├── localdisk          # 包含与本地磁盘操作相关的代码。
│   ├── locale             # 包含国际化和本地化资源，支持多语言。
│   ├── network            # 包含网络服务的代码，负责管理虚拟机的网络连接。
│   ├── notifications      # 包含事件通知机制的代码，用于记录系统中发生的事件。
│   ├── objects            # 包含 Nova 中使用的所有对象模型，通常是指数据库表对应的 ORM 模型。
│   ├── pci                # 包含 PCI 设备管理的相关代码，用于直通或 SR-IOV 等功能。
│   ├── policies           # 包含策略引擎的代码，定义了哪些用户可以执行哪些操作。
│   ├── privsep            # 包含特权分离的代码，允许某些需要更高权限的操作在更安全的环境中执行。
│   ├── scheduler          # 包含调度器的实现，决定新实例应该被放置在哪一个计算节点上。
│   ├── servicegroup       # 包含服务组管理的代码，用于监控 Nova 服务的状态。
│   ├── tests              # 包含单元测试、集成测试和其他类型的测试代码。
│   ├── vendor             # 包含外部依赖库或供应商提供的代码，有时是为了解决特定版本兼容性问题而保留的。
│   ├── virt               # 包含虚拟化驱动程序的代码，负责与不同的 hypervisor 进行交互。
│   ├── vnc                # 包含 VNC 访问的支持代码，使得用户可以通过 VNC 协议访问实例的控制台。
│   └── volume             # 包含卷服务的代码，负责管理持久存储卷，通常与 Cinder 集成。
├── placement-api-ref# 包含 Placement API 的参考文档，Placement 是负责资源追踪的服务。
├── playbooks        # 包含 Ansible Playbook 或类似的配置管理剧本，用于部署或维护操作。
├── releasenotes     # 包含版本发布说明，记录了每次版本更新的内容、改进和已知问题。
├── tools            # 包含各种工具脚本，如构建、测试或其他辅助工具。
└── venv             # 包含虚拟环境配置，用于隔离 Python 环境，确保依赖项正确安装。
```

## nova命令

```bash
usage: nova [command]
nova help command
Command-line interface to the OpenStack Nova API.

Positional arguments:
  <subcommand>
    add-secgroup                Add a Security Group to a server.
    agent-create                Create new agent build.
    agent-delete                Delete existing agent build.
    agent-list                  List all builds.
    agent-modify                Modify existing agent build.
    aggregate-add-host          Add the host to the specified aggregate.
    aggregate-create            Create a new aggregate with the specified
                                details.
    aggregate-delete            Delete the aggregate.
    aggregate-list              Print a list of all aggregates.
    aggregate-remove-host       Remove the specified host from the specified
                                aggregate.
    aggregate-set-metadata      Update the metadata associated with the
                                aggregate.
    aggregate-show              Show details of the specified aggregate.
    aggregate-update            Update the aggregate's name and optionally
                                availability zone.
    attach-usb-device
    availability-zone-list      List all the availability zones.
    backup                      Backup a server by creating a 'backup' type
                                snapshot.
    bind-keypair
    boot                        Boot a new server.
    cell-capacities             Get cell capacities for all cells or a given
                                cell.
    cell-show                   Show details of a given cell.
    cgroup-cpu-show             show cgroup vcpu of server for a tenant or a
                                server. only specify <instance_uuid> or
                                <tenant_id>.
    cgroup-cpu-update           update cgroup vcpu of server for a tenant or a
                                server. only specify <instance_uuid> or
                                <tenant_id>.
    clear-password              Clear the admin password for a server from the
                                metadata server. This action does not actually
                                change the instance server password.
    clone                       Clone a server.
    console-log                 Get console log output of a server.
    dedicated-cloud-add-host    Add the host to the specified dedicated cloud.
    dedicated-cloud-available-host-list
                                Print a list of available host for dedicated
                                cloud.
    dedicated-cloud-create      Create a new dedicated cloud with the
                                specified details.
    dedicated-cloud-delete      Delete the dedicated cloud.
    dedicated-cloud-host-list   Print a list of all hosts in dedicated cloud.
    dedicated-cloud-host-show   Show details of the specified host in
                                dedicated cloud.
    dedicated-cloud-host-statistics
                                Count the resources of all hosts.
    dedicated-cloud-host-update
                                Update the dedicated cloud's display name or
                                cpu ratio or auto deployed.
    dedicated-cloud-list        Print a list of all dedicated clouds.
    dedicated-cloud-remove-host
                                Remove the specified host from the specified
                                dedicated cloud.
    dedicated-cloud-server-boot
                                Boot a new server.
    dedicated-cloud-server-list
                                List servers in dedicated cloud.
    dedicated-cloud-server-show
                                Show details about the given server in
                                dedicated cloud.
    dedicated-cloud-show        Show details of the specified dedicated cloud.
    dedicated-cloud-update      Update the dedicated cloud's name.
    dedicated-host-create       Create a dedicated host.
    dedicated-host-delete       Delete the dedicated host.
    dedicated-host-list         Print a list of all dedicated hosts.
    dedicated-host-show         Display the details of the dedicated host.
    delete                      Immediately shut down and delete specified
                                server(s).
    detach-usb-device
    diagnostics                 Retrieve server diagnostics.
    disk-qoses                  Get disk qos for a server.
    estimate-migration-time
    evacuate                    Evacuate server from failed host.
    extra-specs                 Show details about the given server.
    flavor-access-add           Add flavor access for the given tenant.
    flavor-access-list          Print access information about the given
                                flavor.
    flavor-access-remove        Remove flavor access for the given tenant.
    flavor-create               Create a new flavor.
    flavor-delete               Delete a specific flavor
    flavor-key                  Set or unset extra_spec for a flavor.
    flavor-list                 Print a list of available 'flavors' (sizes of
                                servers).
    flavor-show                 Show details about the given flavor.
    flavor-tag-add              Add one or more tags to a flavor.
    flavor-tag-delete           Delete one or more tags from a flavor.
    flavor-update               Update the description of an existing flavor.
                                (Supported by API versions '2.55' -
                                '2.latest') [hint: use '--os-compute-api-
                                version' flag to show help message for proper
                                version]
    flavor-zone
    force-delete                Force delete a server.
    get-disk-qos                Get a volume disk qos for a server.
    get-mks-console             Get an MKS console to a server. (Supported by
                                API versions '2.8' - '2.latest') [hint: use
                                '--os-compute-api-version' flag to show help
                                message for proper version]
    get-password                Get the admin password for a server. This
                                operation calls the metadata service to query
                                metadata information and does not read
                                password information from the server itself.
    get-rdp-console             Get a rdp console to a server.
    get-serial-console          Get a serial console to a server.
    get-spice-console           Get a spice console to a server.
    get-vnc-console             Get a vnc console to a server.
    host-evacuate               Evacuate all instances from failed host.
    host-evacuate-live          Live migrate all instances of the specified
                                host to other available hosts.
    host-meta                   Set or Delete metadata on all instances of a
                                host.
    host-servers-migrate        Cold migrate all instances off the specified
                                host to other available hosts.
    hypervisor-extra-specs      Get hypervisor extra specs.
    hypervisor-get-ratio        Get hypervisor allocation ratio.
    hypervisor-hostinfo         Display the hostinfo of the specified
                                hypervisor.
    hypervisor-list             List hypervisors. (Supported by API versions
                                '2.0' - '2.latest') [hint: use '--os-compute-
                                api-version' flag to show help message for
                                proper version]
    hypervisor-local-disks      Get hypervisor local disks.
    hypervisor-servers          List servers belonging to specific
                                hypervisors.
    hypervisor-set-instances-limit
                                Set hypervisor instances limit.
    hypervisor-set-ratio        Set hypervisor allocation ratio.
    hypervisor-show             Display the details of the specified
                                hypervisor.
    hypervisor-stats            Get hypervisor statistics over all compute
                                nodes.
    hypervisor-uptime           Display the uptime of the specified
                                hypervisor.
    image-create                Create a new image by taking a snapshot of a
                                running server.
    instance-action             Show an action.
    instance-action-list        List actions on a server. (Supported by API
                                versions '2.0' - '2.latest') [hint: use '--os-
                                compute-api-version' flag to show help message
                                for proper version]
    instance-backup-create      Create an instance backup.
    instance-backup-delete      Delete an instance backup.
    instance-backup-force-delete
                                Force delete a instance backup.
    instance-backup-list        Get a list of instance backup.
    instance-backup-restore     Restore an instance backup to the previous
                                instance.
    instance-backup-show        Get an instance backup details.
    instance-backup-update      Update an instance backup.
    interface-attach            Attach a network interface to a server.
    interface-detach            Detach a network interface from a server.
    interface-list              List interfaces attached to a server.
    keypair-add                 Create a new key pair for use with servers.
    keypair-delete              Delete keypair given by its name. (Supported
                                by API versions '2.0' - '2.latest') [hint: use
                                '--os-compute-api-version' flag to show help
                                message for proper version]
    keypair-list                Print a list of keypairs for a user (Supported
                                by API versions '2.0' - '2.latest') [hint: use
                                '--os-compute-api-version' flag to show help
                                message for proper version]
    keypair-show                Show details about the given keypair.
                                (Supported by API versions '2.0' - '2.latest')
                                [hint: use '--os-compute-api-version' flag to
                                show help message for proper version]
    limits                      Print rate and absolute limits.
    list                        List servers.
    list-extensions             List all the os-api extensions that are
                                available.
    list-file-system            List file system for a server.
    list-secgroup               List Security Group(s) of a server.
    list-usb-devices
    live-migrate-volume-notify
    live-migration              Migrate running server to a new machine.
    live-migration-abort        Abort an on-going live migration. (Supported
                                by API versions '2.24' - '2.latest') [hint:
                                use '--os-compute-api-version' flag to show
                                help message for proper version]
    live-migration-force-complete
                                Force on-going live migration to complete.
                                (Supported by API versions '2.22' -
                                '2.latest') [hint: use '--os-compute-api-
                                version' flag to show help message for proper
                                version]
    live-resize-flavors         Get the flavors list that the server can
                                execute live resize.
    local-disks                 Get local disk for a server.
    lock                        Lock a server. A normal (non-admin) user will
                                not be able to execute actions on a locked
                                server.
    meta                        Set or delete metadata on a server.
    migrate                     Migrate a server.
    migration-list              Print a list of migrations. (Supported by API
                                versions '2.0' - '2.latest') [hint: use '--os-
                                compute-api-version' flag to show help message
                                for proper version]
    mount-file-system           Mount file system for a server.
    offline-compute-host        offline the compute host.
    pause                       Pause a server.
    quota-class-show            List the quotas for a quota class.
    quota-class-update          Update the quotas for a quota class.
                                (Supported by API versions '2.0' - '2.latest')
                                [hint: use '--os-compute-api-version' flag to
                                show help message for proper version]
    quota-defaults              List the default quotas for a tenant.
    quota-delete                Delete quota for a tenant/user so their quota
                                will Revert back to default.
    quota-show                  List the quotas for a tenant/user.
    quota-update                Update the quotas for a tenant/user.
                                (Supported by API versions '2.0' - '2.latest')
                                [hint: use '--os-compute-api-version' flag to
                                show help message for proper version]
    reboot                      Reboot a server.
    rebuild                     Shutdown, re-image, and re-boot a server.
    refresh-network             Refresh server network information.
    remove-secgroup             Remove a Security Group from a server.
    request-spec-az-show        show request spec of the server.
    request-spec-az-update      update request spec of the server.
    rescue                      Reboots a server into rescue mode, which
                                starts the machine from either the initial
                                image or a specified image, attaching the
                                current boot disk as secondary.
    reset-failed-build          Reset failed_builds for the compute host.
    reset-network               Reset network of a server.
    reset-state                 Reset the state of a server.
    resize                      Resize a server.
    resize-confirm              Confirm a previous resize.
    resize-revert               Revert a previous resize (and return to the
                                previous VM).
    restore                     Restore a soft-deleted server.
    resume                      Resume a server.
    server-group-add-member     Add the server to the specified server_group.
    server-group-create         Create a new server group with the specified
                                details.
    server-group-delete         Delete specific server group(s).
    server-group-get            Get a specific server group.
    server-group-has-enough-resources
    server-group-list           Print a list of all server groups.
    server-group-need-migrate
    server-group-remove-member  Remove the server from the specified
                                server_group.
    server-health-check         Health check for a server.
    server-migration-list       Get the migrations list of specified server.
                                (Supported by API versions '2.23' -
                                '2.latest') [hint: use '--os-compute-api-
                                version' flag to show help message for proper
                                version]
    server-migration-show       Get the migration of specified server.
                                (Supported by API versions '2.23' -
                                '2.latest') [hint: use '--os-compute-api-
                                version' flag to show help message for proper
                                version]
    server-snapshot-create      Create a server snapshot.
    server-snapshot-delete      Delete a server snapshot.
    server-snapshot-list        Print a list of server snapshots.
    server-snapshot-reset-state
                                Reset a server snapshot to the special state.
    server-snapshot-restore     Restore a server snapshot to the server.
    server-snapshot-show        Get a server snapshot details.
    server-snapshot-update      Update a server snapshot.
    server-tag-add              Add one or more tags to a server. (Supported
                                by API versions '2.26' - '2.latest') [hint:
                                use '--os-compute-api-version' flag to show
                                help message for proper version]
    server-tag-delete           Delete one or more tags from a server.
                                (Supported by API versions '2.26' -
                                '2.latest') [hint: use '--os-compute-api-
                                version' flag to show help message for proper
                                version]
    server-tag-delete-all       Delete all tags from a server. (Supported by
                                API versions '2.26' - '2.latest') [hint: use
                                '--os-compute-api-version' flag to show help
                                message for proper version]
    server-tag-list             Get list of tags from a server. (Supported by
                                API versions '2.26' - '2.latest') [hint: use
                                '--os-compute-api-version' flag to show help
                                message for proper version]
    server-tag-set              Set list of tags to a server. (Supported by
                                API versions '2.26' - '2.latest') [hint: use
                                '--os-compute-api-version' flag to show help
                                message for proper version]
    service-delete              Delete the service by UUID ID. (Supported by
                                API versions '2.0' - '2.latest') [hint: use
                                '--os-compute-api-version' flag to show help
                                message for proper version]
    service-disable             Disable the service. (Supported by API
                                versions '2.0' - '2.latest') [hint: use '--os-
                                compute-api-version' flag to show help message
                                for proper version]
    service-enable              Enable the service. (Supported by API versions
                                '2.0' - '2.latest') [hint: use '--os-compute-
                                api-version' flag to show help message for
                                proper version]
    service-force-down          Force service to down. (Supported by API
                                versions '2.11' - '2.latest') [hint: use
                                '--os-compute-api-version' flag to show help
                                message for proper version]
    service-list                Show a list of all running services. Filter by
                                host & binary.
    set-disk-qos                Set one disk qos of a server.
    set-password                Change the user password for a server.
    shelve                      Shelve a server.
    shelve-offload              Remove a shelved server from the compute node.
    show                        Show details about the given server.
    show-credit                 show credit of the server.
    ssh                         SSH into a server.
    start                       Start the server(s).
    stop                        Stop the server(s).
    suspend                     Suspend a server.
    sync                        Sync cache for an server.
    trigger-crash-dump          Trigger crash dump in an instance. (Supported
                                by API versions '2.17' - '2.latest') [hint:
                                use '--os-compute-api-version' flag to show
                                help message for proper version]
    unbind-keypair
    unlock                      Unlock a server.
    unmount-file-system         Mount file system for a server.
    unpause                     Unpause a server.
    unrescue                    Restart the server from normal boot disk
                                again.
    unshelve                    Unshelve a server.
    update                      Update some attr for a server.
    usage                       Show usage data for a single tenant.
    usage-list                  List usage data for all tenants.
    version-list                List all API versions.
    volume-attach               Attach a volume to a server.
    volume-attachments          List all the volumes attached to a server.
    volume-detach               Detach a volume from a server.
    volume-live-migrate
    volume-update               Update the attachment on the server. Migrates
                                the data from an attached volume to the
                                specified available volume and swaps out the
                                active attachment to the new volume.
    bash-completion             Prints all of the commands and options to
                                stdout so that the nova.bash_completion script
                                doesn't have to hard code them.
    help                        Display help about this program or one of its
                                subcommands.

Optional arguments:
  --version                     show program's version number and exit
  --debug                       Print debugging output.
  --os-cache                    Use the auth token cache. Defaults to False if
                                env[OS_CACHE] is not set.
  --timings                     Print call timing info.
  --os-region-name <region-name>
                                Defaults to env[OS_REGION_NAME].
  --service-type <service-type>
                                Defaults to compute for most actions.
  --service-name <service-name>
                                Defaults to env[NOVA_SERVICE_NAME].
  --os-endpoint-type <endpoint-type>
                                Defaults to env[NOVA_ENDPOINT_TYPE],
                                env[OS_ENDPOINT_TYPE] or publicURL.
  --os-compute-api-version <compute-api-ver>
                                Accepts X, X.Y (where X is major and Y is
                                minor part) or "X.latest", defaults to
                                env[OS_COMPUTE_API_VERSION].
  --endpoint-override <bypass-url>
                                Use this API endpoint instead of the Service
                                Catalog. Defaults to
                                env[NOVACLIENT_ENDPOINT_OVERRIDE].
  --profile HMAC_KEY            HMAC key to use for encrypting context data
                                for performance profiling of operation. This
                                key should be the value of the HMAC key
                                configured for the OSprofiler middleware in
                                nova; it is specified in the Nova
                                configuration file at "/etc/nova/nova.conf".
                                Without the key, profiling will not be
                                triggered even if OSprofiler is enabled on the
                                server side.
  --os-auth-type <name>, --os-auth-plugin <name>
                                Authentication type to use

See "nova help COMMAND" for help on a specific command.
```

### show

```bash
usage: nova show [flags] <server>

Show details about the given server.

Positional arguments:
  <server>          Name or ID of server.
  
Optional arguments:
  --minimal         Skips flavor/image lookups when showing servers.
  --deleted         show deleted servers (Admin only).
  --wrap <integer>  Wrap the output to a specified length, or 0 to disable.
```



### list

```bash
nova list --name janus
```

```bash
nova list [falg]

Optional arguments:
  --reservation-id <reservation-id>  Only return servers that match reservation-id.
  --ip <ip-regexp>              Search with regular expression match by IP address.
  --ip6 <ip6-regexp>            Search with regular expression match by IPv6 address.
  --name <name-regexp>          Search with regular expression match by name.
  --instance-name <name-regexp>	Search with regular expression match by server name.
  --status <status>             Search by server status, status: ACTIVE
  --flavor <flavor>             Search by flavor name or ID.
  --image <image>               Search by image name or ID.
  --host <hostname>             Search servers by hostname to which they are assigned (Admin only).
  --all-tenants [<0|1>]         Display information from all tenants (Admin only).
  --tenant [<tenant>]           Display information from single tenant (Admin only).
  --user [<user>]               Display information from single user (Admin only).
  --deleted                     Only display deleted servers (Admin only).
  --fields <fields>             Comma-separated list of fields to display. Use
                                the show command to see which fields are
                                available.
  --minimal                     Get only UUID and name.
  --sort <key>[:<direction>]    Comma-separated list of sort keys and
                                directions in the form of <key>[:<asc|desc>].
                                The direction defaults to descending if not
                                specified.
  --marker <marker>             The last server UUID of the previous page;
                                displays list of servers after "marker".
  --limit <limit>               Maximum number of servers to display. If limit
                                == -1, all servers will be displayed. If limit
                                is bigger than 'CONF.api.max_limit' option of
                                Nova API, limit 'CONF.api.max_limit' will be
                                used instead.
  --changes-since <changes_since>
                                List only servers changed after a certain
                                point of time. The provided time should be an
                                ISO 8061 formatted time. ex
                                2016-03-04T06:27:59Z .
  --tags <tags>                 The given tags must all be present for a
                                server to be included in the list result.
                                Boolean expression in this case is 't1 AND
                                t2'. Tags must be separated by commas: --tags
                                <tag1,tag2> (Supported by API versions '2.26'
                                - '2.latest')
  --tags-any <tags-any>         If one of the given tags is present the server
                                will be included in the list result. Boolean
                                expression in this case is 't1 OR t2'. Tags
                                must be separated by commas: --tags-any
                                <tag1,tag2> (Supported by API versions '2.26'
                                - '2.latest')
  --not-tags <not-tags>         Only the servers that do not have any of the
                                given tags will be included in the list
                                results. Boolean expression in this case is
                                'NOT(t1 AND t2)'. Tags must be separated by
                                commas: --not-tags <tag1,tag2> (Supported by
                                API versions '2.26' - '2.latest')
  --not-tags-any <not-tags-any>
                                Only the servers that do not have at least one
                                of the given tags will be included in the list
                                result. Boolean expression in this case is
                                'NOT(t1 OR t2)'. Tags must be separated by
                                commas: --not-tags-any <tag1,tag2> (Supported
                                by API versions '2.26' - '2.latest')
```

### live-migration

```bash
usage: nova live-migration [flags] <server> [<host>]

Migrate running server to a new machine.

Positional arguments:
  <server>                      Name or ID of server.
  <host>                        Destination host name.

Optional arguments:
  --block-migrate               True in case of block_migration.
                                (Default=auto:live_migration) (Supported by
                                API versions '2.25' - '2.latest')
  --force                       Force to not verify the scheduler if a host is
                                provided. (Supported by API versions '2.30' -
                                '2.latest')
  --migration-type <migration-type>
                                For local storage migration to shared storage.
                                e.g. to_<volume_type_name>(see 'cinder type-
                                list'): migrate all the disks of the instance
                                to the volume.
```

### live-migration-abort

```bash
[root@gz-txjs-control-55e243e31e30 ~]# nova help live-migration-abort
usage: nova live-migration-abort <server> <migration>

Abort an on-going live migration. (Supported by API versions '2.24' -
'2.latest') [hint: use '--os-compute-api-version' flag to show help message
for proper version]

Positional arguments:
  <server>     Name or ID of server.
  <migration>  ID of migration.
```

### server-migration-list

```bash
[root@gz-txjs-control-55e243e31e30 ~]# nova help server-migration-list
usage: nova server-migration-list <server>

Get the migrations list of specified server. (Supported by API versions '2.23'
- '2.latest') [hint: use '--os-compute-api-version' flag to show help message
for proper version]

Positional arguments:
  <server>  Name or ID of server.
```

### migration-list

```bash
usage: nova migration-list [--instance-uuid <instance_uuid>] [--host <host>]
                           [--status <status>] [--marker <marker>]
                           [--limit <limit>] [--changes-since <changes_since>]

Print a list of migrations. (Supported by API versions '2.0' - '2.latest')
[hint: use '--os-compute-api-version' flag to show help message for proper
version]

Optional arguments:
  --instance-uuid <instance_uuid>
                                Fetch migrations for the given instance.
  --host <host>                 Fetch migrations for the given host.
  --status <status>             Fetch migrations for the given status.
  --marker <marker>             The last migration of the previous page;
                                displays list of migrations after "marker".
                                Note that the marker is the migration UUID.
  --limit <limit>               Maximum number of migrations to display. Note
                                that there is a configurable max limit on the
                                server, and the limit that is used will be the
                                minimum between what is requested here and
                                what is configured in the server.
  --changes-since <changes_since>
                                List only migrations changed after a certain
                                point of time. The provided time should be an
                                ISO 8061 formatted time. ex
                                2016-03-04T06:27:59Z .
```



### reset-state

```bash
usage: nova reset-state [plags] <server> [<server> ...]

Reset the state of a server.

Positional arguments:
  <server>       Name or ID of server(s).

Optional arguments:
  --all-tenants  Reset state server(s) in another tenant by name (Admin only).
  --active       Request the server be reset to "active" state instead of
                 "error" state (the default).
  --stopped      Request the server be reset to "stopped" state
```

### instance-action-list 

```bash
usage: nova instance-action-list [--marker <marker>] [--limit <limit>]
                                 [--changes-since <changes_since>]
                                 <server>

List actions on a server. (Supported by API versions '2.0' - '2.latest')
[hint: use '--os-compute-api-version' flag to show help message for proper
version]

Positional arguments:
  <server>                      Name or UUID of the server to list actions
                                for. Only UUID can be used to list actions on
                                a deleted server.

Optional arguments:
  --marker <marker>             The last instance action of the previous page;
                                displays list of actions after "marker".
  --limit <limit>               Maximum number of instance actions to display.
                                Note that there is a configurable max limit on
                                the server, and the limit that is used will be
                                the minimum between what is requested here and
                                what is configured in the server.
  --changes-since <changes_since>
                                List only instance actions changed after a
                                certain point of time. The provided time
                                should be an ISO 8061 formatted time. ex
                                2016-03-04T06:27:59Z
```

###  instance-action

```bash
usage: nova instance-action <server> <request_id>

Show an action.

Positional arguments:
  <server>      Name or UUID of the server to show actions for. Only UUID can
                be used to show actions for a deleted server. (Supported by
                API versions '2.21' - '2.latest')
  <request_id>  Request ID of the action to get.
```

### rebuild

```bash
usage: nova rebuild [--image <image>] [--rebuild-password <rebuild-password>]
                    [--poll] [--minimal] [--preserve-ephemeral]
                    [--name <name>] [--description <description>]
                    [--meta <key=value>] [--key-name <key-name>] [--key-unset]
                    [--user-data <user-data>] [--user-data-unset]
                    [--volume <volume>]
                    <server>

Shutdown, re-image, and re-boot a server.

Positional arguments:
  <server>                      Name or ID of server.

Optional arguments:
  --image <image>               Name or ID of new image.
  --rebuild-password <rebuild-password>
                                Set the provided admin password on the rebuilt
                                server.
  --poll                        Report the server rebuild progress until it
                                completes.
  --minimal                     Skips flavor/image lookups when showing
                                servers.
  --preserve-ephemeral          Preserve the default ephemeral storage
                                partition on rebuild.
  --name <name>                 Name for the new server.
  --description <description>   New description for the server. (Supported by
                                API versions '2.19' - '2.latest')
  --meta <key=value>            Record arbitrary key/value metadata to
                                /meta_data.json on the metadata server. Can be
                                specified multiple times.
  --key-name <key-name>         Keypair name to set in the server. Cannot be
                                specified with the '--key-unset' option.
                                (Supported by API versions '2.54' -
                                '2.latest')
  --key-unset                   Unset keypair in the server. Cannot be
                                specified with the '--key-name' option.
                                (Supported by API versions '2.54' -
                                '2.latest')
  --user-data <user-data>       User data file to pass to be exposed by the
                                metadata server. (Supported by API versions
                                '2.57' - '2.latest')
  --user-data-unset             Unset user_data in the server. Cannot be
                                specified with the '--user-data' option.
                                (Supported by API versions '2.57' -
                                '2.latest')
  --volume <volume>             ID of the volume to rebuild. (Supported by API
                                versions '2.61' - '2.latest')
```

### volume-attach  

```bash
nova volume-attach janus 447522bf-5f8c-41c1-9b10-a629efac9a9c
```

```bash
usage: nova volume-attach [--tag <tag>] [--disk_bus <disk_bus>]
                          <server> <volume> [<device>]

Attach a volume to a server.

Positional arguments:
  <server>               Name or ID of server.
  <volume>               ID of the volume to attach.
  <device>               Name of the device e.g. /dev/vdb. Use "auto" for
                         autoassign (if supported). Libvirt driver will use
                         default device name.

Optional arguments:
  --tag <tag>            Tag for the attached volume. (Supported by API
                         versions '2.49' - '2.latest')
  --disk_bus <disk_bus>  Disk bus for the attached volume, can be chosen from
                         ide, sata,scsi, virtio and etc. (Supported by API
                         versions '2.49' - '2.latest')
```



# cinder

``` bash
usage: cinder [command]
Command-line interface to the OpenStack Cinder API.

Positional arguments:
  <subcommand>
    absolute-limits     Lists absolute limits for a user.
    api-version         Display the server API version information. (Supported
                        by API versions 3.0 - 3.latest)
    availability-zone-list
                        Lists all availability zones.
    backup-abort        Aborts one or more backups.
    backup-create       Creates a volume backup.
    backup-delete       Removes one or more backups.
    backup-export       Export backup metadata record.
    backup-import       Import backup metadata record.
    backup-list         Lists all backups.
    backup-reset-state  Explicitly updates the backup state.
    backup-restore      Restores a backup.
    backup-show         Shows backup details.
    calc-backup-hmac    Calculate backup hmac.
    calc-snapshot-hmac  Calculate snapshot hmac.
    calc-volume-hmac    Calculate volume hmac.
    cgsnapshot-create   Creates a cgsnapshot.
    cgsnapshot-delete   Removes one or more cgsnapshots.
    cgsnapshot-list     Lists all cgsnapshots.
    cgsnapshot-show     Shows cgsnapshot details.
    clean-reserved-time
                        Clean volume reserved time.
    clone-image-metadata
                        Clone image metadata.
    config-metadata     Sets or unsets or updates config metadata.
    consisgroup-create  Creates a consistency group.
    consisgroup-create-from-src
                        Creates a consistency group from a cgsnapshot or a
                        source CG.
    consisgroup-delete  Removes one or more consistency groups.
    consisgroup-list    Lists all consistency groups.
    consisgroup-show    Shows details of a consistency group.
    consisgroup-update  Updates a consistency group.
    create              Creates a volume.
    credentials         Shows user credentials returned from auth.
    delete              Removes one or more volumes.
    delete-cancel       Cancel one or more deletions.
    delete-confirm      Confirm one or more deletions.
    delete-volume-hmac  Delete volume hmac.
    encryption-type-create
                        Creates encryption type for a volume type. Admin only.
    encryption-type-delete
                        Deletes encryption type for a volume type. Admin only.
    encryption-type-list
                        Shows encryption type details for volume types. Admin
                        only.
    encryption-type-show
                        Shows encryption type details for a volume type. Admin
                        only.
    encryption-type-update
                        Update encryption type information for a volume type
                        (Admin Only).
    endpoints           Discovers endpoints registered by authentication
                        service.
    extend              Attempts to extend size of an existing volume.
    extra-specs-list    Lists current volume types and extra specs.
    failover-host       Failover a replicating cinder-volume host.
    force-delete        Attempts force-delete of volume, regardless of state.
    force-delete-system-reserved
                        Force delete certain system_reserved volume(s).
    freeze-host         Freeze and disable the specified cinder-volume host.
    get-capabilities    Show backend volume stats and properties. Admin only.
    get-cinder-version
    get-delete-confirm-resource
                        Show a delete confirm resource by verify_id.
    get-pools           Show pool information for backends. Admin only.
    get-qos             Shows volume qos.
    get-reserved-time   Get volume reserved time.
    get-system-reserved-volumes
                        Get system reserved (user deleted) volumes with
                        reserved time.
    get-volume-snapshots-used
                        Shows volume qos.
    get-volumes-with-reserved-time
                        Get volumes with reserved time.
    group-snapshot-list
                        Lists all group snapshots.
    image-metadata      Sets or deletes volume image metadata.
    image-metadata-show
                        Shows volume image metadata.
    interrupt-io        Interrupt or recover IO operations on volumes.
    list                Lists all volumes.
    list-delete-confirm-resources
                        Get delete confirm resources.
    live-retype         Changes the volume type for a volume.
    manage              Manage an existing volume.
    metadata            Sets or deletes volume metadata.
    metadata-show       Shows volume metadata.
    metadata-update-all
                        Updates volume metadata.
    migrate             Migrates volume to a new host.
    qos-associate       Associates qos specs with specified volume type.
    qos-create          Creates a qos specs.
    qos-delete          Deletes a specified qos specs.
    qos-disassociate    Disassociates qos specs from specified volume type.
    qos-disassociate-all
                        Disassociates qos specs from all its associations.
    qos-get-association
                        Lists all associations for specified qos specs.
    qos-key             Sets or unsets specifications for a qos spec.
    qos-list            Lists qos specs.
    qos-show            Shows qos specs details.
    quota-class-show    Lists quotas for a quota class.
    quota-class-update  Updates quotas for a quota class.
    quota-defaults      Lists default quotas for a tenant.
    quota-delete        Delete the quotas for a tenant.
    quota-show          Lists quotas for a tenant.
    quota-update        Updates quotas for a tenant.
    quota-usage         Lists quota usage for a tenant.
    rate-limits         Lists rate limits for a user.
    readonly-mode-update
                        Updates volume read-only access-mode flag.
    rename              Renames a volume.
    replication-promote
                        Promote a secondary volume to primary for a
                        relationship.
    replication-reenable
                        Sync the secondary volume with primary for a
                        relationship.
    reset-state         Explicitly updates the entity state in the Cinder
                        database.
    retype              Changes the volume type for a volume.
    service-disable     Disables the service.
    service-enable      Enables the service.
    service-list        Lists all services. Filter by host and service binary.
                        (Supported by API versions 3.0 - 3.latest)
    set-bootable        Update bootable status of a volume.
    set-qos             Set qos for volume.
    set-reserved-time   Set volume reserved time.
    show                Shows volume details.
    snapshot-cg-create  Creates a cgsnapshot.
    snapshot-cg-delete  Removes one or more cgsnapshots.
    snapshot-cg-remove  Removes one or more snapshots from a consistency
                        snapshot group.
    snapshot-cg-restore
                        Restores a single cgsnapshot.
    snapshot-cg-show    Shows cgsnapshot details.
    snapshot-create     Creates a snapshot.
    snapshot-delete     Removes one or more snapshots.
    snapshot-list       Lists all snapshots.
    snapshot-manage     Manage an existing snapshot.
    snapshot-metadata   Sets or deletes snapshot metadata.
    snapshot-metadata-show
                        Shows snapshot metadata.
    snapshot-metadata-update-all
                        Updates snapshot metadata.
    snapshot-rename     Renames a snapshot.
    snapshot-reset-state
                        Explicitly updates the snapshot state.
    snapshot-show       Shows snapshot details.
    snapshot-unmanage   Stop managing a snapshot.
    system-restore      Restore a system_reserved volume.
    thaw-host           Thaw and enable the specified cinder-volume host.
    transfer-accept     Accepts a volume transfer.
    transfer-create     Creates a volume transfer.
    transfer-delete     Undoes a transfer.
    transfer-list       Lists all transfers.
    transfer-show       Shows transfer details.
    type-access-add     Adds volume type access for the given project.
    type-access-list    Print access information about the given volume type.
    type-access-remove  Removes volume type access for the given project.
    type-create         Creates a volume type.
    type-default        List the default volume type.
    type-delete         Deletes volume type or types.
    type-key            Sets or unsets extra_spec for a volume type.
    type-list           Lists available 'volume types'.
    type-show           Show volume type details.
    type-update         Updates volume type name, description, and/or
                        is_public.
    unmanage            Stop managing a volume.
    update-volume-type  Changes the volume type for a volume to diff qos in
                        same pool.
    upload-to-image     Uploads volume to Image Service as an image.
    veri-backup-hmac    Verify snapshot hmac.
    veri-snapshot-hmac  Verify snapshot hmac.
    veri-volume-hmac    Verify volume hmac.
    version-list        List all API versions. (Supported by API versions 3.0
                        - 3.latest)
    volume-lock         Lock a volume, so the volume cannot be deleted until
                        it's unlocked.
    volume-unlock       UnLock a volume, so the volume can be deleted.
    bash-completion     Prints arguments for bash_completion.
    help                Shows help about this program or one of its
                        subcommands.
    list-extensions     Lists all available os-api extensions.

Optional arguments:
  --version             show program's version number and exit
  -d, --debug           Shows debugging output.
  --service-type <service-type>
                        Service type. For most actions, default is volume.
  --service-name <service-name>
                        Service name. Default=env[CINDER_SERVICE_NAME].
  --volume-service-name <volume-service-name>
                        Volume service name.
                        Default=env[CINDER_VOLUME_SERVICE_NAME].
  --os-endpoint-type <os-endpoint-type>
                        Endpoint type, which is publicURL or internalURL.
                        Default=env[OS_ENDPOINT_TYPE] or nova
                        env[CINDER_ENDPOINT_TYPE] or publicURL.
  --endpoint-type <endpoint-type>
                        DEPRECATED! Use --os-endpoint-type.
  --os-volume-api-version <volume-api-ver>
                        Block Storage API version. Accepts X, X.Y (where X is
                        major and Y is minor
                        part).Default=env[OS_VOLUME_API_VERSION].
  --bypass-url <bypass-url>
                        DEPRECATED! Use os_endpoint. Use this API endpoint
                        instead of the Service Catalog. Defaults to
                        env[CINDERCLIENT_BYPASS_URL].
  --os-endpoint <os-endpoint>
                        Use this API endpoint instead of the Service Catalog.
                        Defaults to env[CINDER_ENDPOINT].
  --retries <retries>   Number of retries.
  --profile HMAC_KEY    HMAC key to use for encrypting context data for
                        performance profiling of operation. This key needs to
                        match the one configured on the cinder api server.
                        Without key the profiling will not be triggered even
                        if osprofiler is enabled on server side. Defaults to
                        env[OS_PROFILE].
  --os-auth-type <name>, --os-auth-plugin <name>
                        Authentication type to use
  --os-auth-strategy <auth-strategy>
                        Authentication strategy (Env: OS_AUTH_STRATEGY,
                        default keystone). For now, any other value will
                        disable the authentication.
  --os-auth-system <os-auth-system>
                        DEPRECATED! Use --os-auth-type. Defaults to
                        env[OS_AUTH_SYSTEM].
  --os-token <token>    Defaults to env[OS_TOKEN].
  --os-url <url>        Defaults to env[OS_URL].
  --os-delete-token <delete-token>
                        Delete token for admin role, default=(Env:
                        OS_DELETE_TOKEN)

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

Run "cinder help SUBCOMMAND" for help on a subcommand.
```

## create  

```bash
cinder create --name janus --image-id 24647239-7b2e-4895-8ede-1d278a3b10df 50 
```

```bash
usage: cinder create [flags] [size]

Creates a volume.

Positional arguments:
  <size>                Size of volume, in GiBs. (Required unless snapshot-id
                        /source-volid is specified).

Optional arguments:
  --consisgroup-id <consistencygroup-id>
                        ID of a consistency group where the new volume belongs
                        to. Default=None.
  --snapshot-id <snapshot-id>
                        Creates volume from snapshot ID. Default=None.
  --source-volid <source-volid>
                        Creates volume from volume ID. Default=None.
  --source-replica <source-replica>
                        Creates volume from replicated volume ID.
                        Default=None.
  --image-id <image-id>
                        Creates volume from image ID. Default=None.
  --image <image>       Creates a volume from image (ID or name).
                        Default=None.
  --name <name>         Volume name. Default=None.
  --description <description>
                        Volume description. Default=None.
  --volume-type <volume-type>
                        Volume type. Default=None.
  --availability-zone <availability-zone>
                        Availability zone for volume. Default=None.
  --metadata [<key=value> [<key=value> ...]]
                        Metadata key and value pairs. Default=None.
  --hint <key=value>    Scheduler hint, like in nova.
  --allow-multiattach   Allow volume to be attached more than once.
                        Default=False
  --poll                Wait for volume creation until it completes.
  --count <number>      Create <number> volumes (limited by quota).
  --id <id>             Creates a volume with specific id. Default=None.
```

## list

```bash
usage: cinder list [--all-tenants [<0|1>]] [--name <name>] [--status <status>]
                   [--bootable [<True|true|False|false>]]
                   [--migration_status <migration_status>]
                   [--metadata [<key=value> [<key=value> ...]]]
                   [--marker <marker>] [--limit <limit>] [--fields <fields>]
                   [--sort <key>[:<direction>]] [--tenant [<tenant>]]

Lists all volumes.

Optional arguments:
  --all-tenants [<0|1>]
                        Shows details for all tenants. Admin only.
  --name <name>         Filters results by a name. Default=None. This option
                        is deprecated and will be removed in newer release.
                        Please use '--filters' option which is introduced
                        since 3.33 instead.
  --status <status>     Filters results by a status. Default=None. This option
                        is deprecated and will be removed in newer release.
                        Please use '--filters' option which is introduced
                        since 3.33 instead.
  --bootable [<True|true|False|false>]
                        Filters results by bootable status. Default=None. This
                        option is deprecated and will be removed in newer
                        release. Please use '--filters' option which is
                        introduced since 3.33 instead.
  --migration_status <migration_status>
                        Filters results by a migration status. Default=None.
                        Admin only. This option is deprecated and will be
                        removed in newer release. Please use '--filters'
                        option which is introduced since 3.33 instead.
  --metadata [<key=value> [<key=value> ...]]
                        Filters results by a metadata key and value pair.
                        Default=None. This option is deprecated and will be
                        removed in newer release. Please use '--filters'
                        option which is introduced since 3.33 instead.
  --marker <marker>     Begin returning volumes that appear later in the
                        volume list than that represented by this volume id.
                        Default=None.
  --limit <limit>       Maximum number of volumes to return. Default=None.
  --fields <fields>     Comma-separated list of fields to display. Use the
                        show command to see which fields are available.
                        Unavailable/non-existent fields will be ignored.
                        Default=None.
  --sort <key>[:<direction>]
                        Comma-separated list of sort keys and directions in
                        the form of <key>[:<asc|desc>]. Valid keys: id,
                        status, size, availability_zone, name, bootable,
                        created_at, reference. Default=None.
  --tenant [<tenant>]   Display information from single tenant (Admin only).
```

## delete

```bash
usage: cinder delete [--cascade] [--remove-in-db] <volume> [<volume> ...]

Removes one or more volumes.

Positional arguments:
  <volume>        Name or ID of volume or volumes to delete.

Optional arguments:
  --cascade       Remove any snapshots along with volume. Default=False.
  --remove-in-db  Remove deleted record in db at the same time. Default=False.
```

