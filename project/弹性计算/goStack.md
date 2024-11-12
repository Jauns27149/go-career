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