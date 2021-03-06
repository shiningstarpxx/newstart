先重点梳理后端部分

## 医生端后端

#### 题目

云原生的腾讯云医的微服务架构以及DevOps实践

#### AGENDA

1. 云原生的微服务架构
   1. 给一版初始的架构 (grpc直接暴露是否合适？)
   2. 重点讲的还是event log 
      1. DTS - cache 的引入
      
      2. 业务指标监控 - 洞察, 从服务可观测入手？再到业务洞察？case
      
      3. 业务event 引入
      
         参考资料
      
         之前自己翻译的那个...
      4. saga 模式?? 这个应该也可以重点讲
      
         [linked-in-first](https://www.linkedin.com/pulse/microservice-communication-arpit-jain)
      
         [confluent](https://www.slideshare.net/ConfluentInc/eventdriven-model-serving-stream-processing-vs-rpc-with-kafka-and-tensorflow-kai-waehner-confluent-kafka-summit-sf-2019-179256721)
      
         [medium](https://supunbhagya.medium.com/request-driven-vs-event-driven-microservices-7b1fe40dccde)
   3. 可观测
      1. metrics 监控
      2. cls
      3. tracing 
   4. 系统的容灾&容量
   5. 系统的整体架构图 云基础+无状态+有状态； 适应于所有c端的微服务架构，满足api + event 两种模式
2. DevOps实践
   1. 现状
   2. 流水线
   3. 多环境
   4. 脚手架
   5. coding的融合









----------



#### 效能相关 (这个部分暂定往后放)

1. 问题： 整个研发过程中遇到的瓶颈梳理 

   1. 材料： 

      1. 相关资料在质量申请获奖材料中有总结  

   2.  效能二期      

      1. 后面让小黎补充一部分 

         

#### 微服务部分

微服务域的初步划分与对应子域服务优化

1. 消息服务优化
   1. 问题：
   2. 材料：https://app.diagrams.net/#G1Nu2LFGhKsH89PxDm6fiLM6lhISldd2AV
2. 会话服务优化
   1. 问题：
   2. 材料：
3. 处方服务优化
   1. 问题：
   2. 材料：
4. 引入接入层-微服务架构层优化
   1. 问题：
   2. 材料：鑫爷答辩材料提供小部分材料
5. 订单服务优化 
   1. 重构资料 https://docs.qq.com/sheet/DUnBCTFJLYmhMdEpk?tab=rcz0bm
6. 权限+多租户



#### 云原生技术+运营容灾

##### 服务监测   

1. 总体理论支持

   1. metrics
   2. log 优化     
      1.  LOG的理论分类     
      2.  LOG存在的问题，优化理论支撑，优化过程      
      3. 材料：         
         1.  鑫爷答辩材料提供小部分
   3. tracing

   

   #### 微服务架构优化

   1. 数据层增加cache      
      1. 问题：      
      2. 材料：
   2. 跨域数据域的transaction
      1. 问题：
      2. 材料：
   3. 流式系统的应用
      1. 问题：
         1. 技术指标的监测
            1. 业务指标 --- 跟数据组有冲突
         2. DB->Cache的发布
         3. 整个业务系统异步调用 (https://git.woa.com/michaelpei/golang_source_analyses/blob/master/doc/event_log/日志和事件.md)
            1. 增加消息队列，event化，做线上，线下隔离
            2. event+流式处理系统

##### 云原生	

	1. 现存问题：     
	2. 未来方向：

##### 容灾

 	1. 服务分级   	
 	  	1. 分级理论： 
 	       	1. 实时                
 	       	2. 非实时
 	  	2. 熔断保护

##### 安全

1. 审计



#### 开源协同

1. BG的工程效能建设
2. 开源协同 -- 后台
3. 开源协同 -- 前台





### 医生端前

1. 端框架优化&性能优化
   1. 框架 mpvue -> uniapp
   2. 包大小优化
   3. 性能优化
2. 效能优化



##### 后端需要补充的能力

1. 可用性能力
   1. 定义和测量覆
   2. 盖系统中所有的关键事件点
   3. 服务分级： 至关重要，重要但非必须
   4. 风险模型

#### 灵魂深处的问题

	1. 微服务的技术难点有哪些
	2. 本质上，微服务不是设计问题？而是由于云原生，以及devopos带来的必然趋势？

#### Reference



#### Quote

> Nothiing is built on stone; all is built on sand, but we must build as if the sand were stone.
>
> ---- Jorge Luis Borges



