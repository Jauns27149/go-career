

# 云计算服务模型

- **IaaS（Infrastructure as a Service）**：提供基础计算资源，用户管理操作系统及之上的一切。
- **PaaS（Platform as a Service）**：提供开发和部署平台，用户管理应用程序和数据。
- **SaaS（Software as a Service）**：提供完整应用程序，用户直接使用应用程序。
- **FaaS（Functions as a Service）**：提供无服务器计算服务，用户上传代码片段按需执行。
- **DaaS（Data as a Service）**：提供数据管理和分析服务，用户访问和使用数据。
- **CaaS（Containers as a Service）**：提供容器管理服务，用户管理容器化的应用程序。

# openStack

​	OpenStack是一个云操作系统，通过数据中心可控制大型的计算、存储、网络等资源池。所有的管理通过前端界面管理员就可以完成，同样也可以通过web接口让最终用户部署资源。

- Keystone(认证服务) :  OpenStack 的身份认证服务，提供了一个集中的身份验证机制，支持多种身份验证方法，包括用户名和密码、API 密钥、OAuth 和令牌。
- Glance(镜像服务): 管理虚拟机镜像，提供了一个注册表服务，用户可以查询和检索虚拟机镜像，这些镜像用于创建新的虚拟机实例。
- Nova(计算服务): 是 OpenStack 的计算组件，负责管理虚拟机实例的生命周期，包括创建、调度、迁移和销毁虚拟机。它支持多种虚拟化技术，如 KVM、Xen 和 VMware。
- Neutron(网络服务): 提供虚拟网络基础设施的功能，使用户能够创建和管理网络连接，支持各种网络拓扑和服务，如防火墙、负载均衡和虚拟专用网络（VPN）。
- Cinder(块储存服务): 提供持久性的块存储服务，用于创建和管理虚拟磁盘，这些磁盘可以附加到虚拟机上，提供额外的存储空间。
- Swift(对象存储服务): 提供了一个可扩展的、冗余的对象存储系统，用于存储非结构化的数据，如图片或视频文件。
- Heat(编排服务): 一个编排引擎，允许用户定义和管理复杂的云应用，通过模板描述应用的资源需求和依赖关系，Heat 自动化资源的创建和配置过程。
- Trove(数据库服务): 提供了一个数据库即服务的解决方案，使用户能够轻松地管理和部署关系型数据库。

## nova

- 概念

1. Nova和Swift是OpenStack最早的两个组件，nova分为控制节点和计算节点。
2. 计算节点通过Nova Compute进行虚拟机创建，通过libvirt调用kvm创建虚拟机，nova之间通信通过rabbitMQ队列进行通信。
3. Nova位于Openstack架构的中心，其他服务或者组件（比如Glance、Cinder、Neutron等）对它提供支持，另外它本身的架构也比较复杂。

- 作用

1. Nova是OpenStack最核心的服务模块，负责管理和维护云计算环境的计算资源，负责整个云环境虚拟机生命周期的管理。
2. Nova是OpenStack的计算服务，负责维护和管理的网络和存储，提供计算服务。

- 组件

![stickPicture](assets/stickPicture.png)

### 工作流程

![stickPicture](assets/stickPicture-1731403108838-10.png)

![stickPicture](assets/stickPicture-1731403218807-15.png)

![stickPicture](assets/stickPicture-1731403224722-17.png)

![stickPicture](assets/stickPicture-1731403231405-19.png)
