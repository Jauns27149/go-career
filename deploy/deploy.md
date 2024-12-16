# Kubernetes

## kubectl 命令

```bash
kubectl <command> [flag...]

Commands:
  get           Display one or many resources
  edit          Edit a resource on the server
  delete        Delete resources by filenames, stdin, resources and names, or by resources and label selector
  
flags:
	-o <wode|json|yaml>	选择输出格式
	-n <namespace>	命令空间
```

### get

```bash
kubectl get [command]

commands:
	svc <service-name>	查看指定服务的相关信息
```

#### pods

```bash
kubectl get pods [flag]

flags:
	-n <namespace>	命令空间
	-l --selector <selector>	根据上设置的键值对标签来选择特定的资源
```

```bash
kubectl exec -it -n az3 gostack-mongos-0 -- /bin/sh #进入pod内部
```

```bash
MongoConf:
  EndPoint: "gostack-mongos-0.gostack-mongos.az3.svc.cluster.net:27017"
  EndPoints:
    - "gostack-mongos-0.gostack-mongos.az3.svc.cluster.net:27017" 
  Username: "root"
  Password: "testaz3"
  Timeout: 300

```

```bash
mongo -u "root" -p "testaz3" --authenticationDatabase "admin"
```

