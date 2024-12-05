# 框架和组件功能

![AgAAMBUoAXPt5c7EdR5OK5Tvg9nbQxeO](assets/AgAAMBUoAXPt5c7EdR5OK5Tvg9nbQxeO.png)

| 组件          | 功能                                                         |
| ------------- | ------------------------------------------------------------ |
| Agent         | 计算节点常驻程序，负责在计算节点执行任务 ，上报宿主机、虚拟机状态 |
| API           | 提供RESTfuI接口供前端调用、基础数据增删改查功能、将前端请求处理成Task并记录至数据库 |
| Engine        | 核心编排引擎，监听并处理Task执行异步处理流程                 |
| DTCD          | 存储数据，通过watch机制同步状态                              |
| Scheduler     | 调度引擎，负责处理资源调度                                   |
| Cron          | 周期性任务注册、执行组件                                     |
| Event         | 对接存储、网络外部事件驱动                                   |
| GoStack Proxy | 将SDK Proxy请求转换为GoStack API                             |
| SDK Proxy     | 云管平台上层与底层API交互转换层                              |
| GoStone       | 平台组件鉴权平台                                             |
| Go VNC Proxy  | VNC转发建联组件                                              |
| Gemini        | 宿主机HA，故障检测、隔离、恢复平台                           |

时序图

![AgAAMBUoAXOF5m3zyM1JRZCrIzeNylj8](assets/AgAAMBUoAXOF5m3zyM1JRZCrIzeNylj8.png)

<img src="assets/AgAAMBUoAXMbRp-uue9Kw6MixBrhyUCW.png" alt="AgAAMBUoAXMbRp-uue9Kw6MixBrhyUCW" style="zoom:67%;" />

![image-20241120174806266](assets/image-20241120174806266.png)

# 运行流程

1. 加载配置文件,初始化读取template
2. 上报健康状态
3. 加载已有数据并构建缓存（cache）
4. 开启事件监听服务(workflow_service)
   - WatchEvent（重点）

# 业务数理

## 虚拟机快照

- 接口
  - 获取虚拟机快照列表 
  - 获取虚拟机快照信息

- 小缺陷

  - swag关于接口描述的注解，表达不清晰。

    ![image-20241201163031071](/home/janus/.config/Typora/typora-user-images/image-20241201163031071.png)

### 获取虚拟机快照列表

- 流程：从mongo获取快照信息
- 疑问：快照信息存入的时机待探索

### 获取虚拟机快照信息



## 创建虚拟机

1. 请求体数据校验与绑定
   1. body是否存在，否--return err
   2. config_drive是否存在，否--ConfigDrive = true
   3. password 校验，无 --随即生成
   4. createInstanceRequest 格式校验(validate)
   5. Metadata校验
   6. 是否有name，无 name = hostname
   7. 处理网络，不支持pf网卡
2. keypair检查
3. 查找flavor
4. 创建虚拟机
5. 提交事务



# gostack项目运转流程

1. 结构
   - agent : 执行Job
   - api : 前端请求处理成Task并记录至数据库
   - engine : 监听并处理Task执行异步处理流程
   - scheduler : 负责处理资源调度
   - cron : 周期性任务注册、执行组件

模块之间是通过 etcd 的 watch 机制来时间进行模块之间的交互，实现类似远程调用的效果。

- 流程
  1. api模块根据请求，决定是否需要需要生成task
  2. engine模块
     1. 根据FlowName进行job编排，生成plan，以及决定是否需要scheduler进行调度
     2. plan下每个jod逐一在agent 执行，返回结果
     3. task、plan、jod执行过程的相关数据同步到mongo
     4. 有job有错误，直接返回错误响应，全部执行成功则返回成功响应
  3. agent执行jod, 

![AgAAMBUoAXOF5m3zyM1JRZCrIzeNylj8](assets/AgAAMBUoAXOF5m3zyM1JRZCrIzeNylj8.png)