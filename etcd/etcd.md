	ETCD 是一个高可用的分布式键值存储系统，常用于分布式系统的配置管理和服务发现。它是 CoreOS 团队开发的一个开源项目，广泛应用于  Kubernetes、Docker Swarm 等容器编排系统中。ETCD 的设计目标是提供一个简单、可靠且高性能的分布式存储解决方案。

- etcdctl : 命令行工具
  - get <key> : 获取键值
  - put <key> <value> : 设置键值
  - del --from-key / : 删除所有键数据
  - get  "" --prefix  --keys-only / : 查看所有的key值