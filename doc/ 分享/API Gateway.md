#### API Gateway

#### author: Pei, Xingxin

#### email: xingxinpei@gmail.com, michaelpei@tencent.com



#### Why API GATEWAY



#### 选型讨论

来自网上的调研1

> ### API网关选型
>
> - 私有云开源解决方案：
>
> Netflix Zuul，zuul是spring cloud的一个推荐组件，https://github.com/Netflix/zuul
> Kong kong是基于Nginx+Lua进行二次开发的方案， https://konghq.com/
> Tyk是2014年创建的开源API网关，甚至比AWS的API网关即服务功能还要早。Tyk用Golang编写并使用Golang自己的HTTP服务器。
>
> - 公有云解决方案：
>
> Amazon API Gateway，https://aws.amazon.com/cn/api-gateway/
> 阿里云API网关，https://www.aliyun.com/product/apigateway/
> 腾讯云API网关， https://cloud.tencent.com/product/apigateway
>
> - 自开发解决方案
>
> 基于Nginx+Lua+ OpenResty的方案，可以看到Kong,orange都是基于这个方案
> 基于Netty、非阻塞IO模型。 通过网上搜索可以看到国内的宜人贷等一些公司是基于这种方案，是一种成熟的方案。
> 基于Node.js的方案。 这种方案是应用了Node.js天生的非阻塞的特性。
> 基于java Servlet的方案。 zuul基于的就是这种方案，这种方案的效率不高，这也是zuul总是被诟病的原因。
>
> 如果没有开源项目的支撑前提下，自己来做这样一套东西，是非常大的一个工作量，而且还要做 API 网关本身的高可用等，如果一旦做不好，有可能最先挂掉的不是你的其他服务，而就是这个API网关。
>
> 国外目前比较成型的是kong和tyk
>
> kong是基于ngnix+lua的，它借助于Nginx的事件驱动模型和非阻塞IO，性能方面是非常棒的，但是从公司的角度比较难于找到能去维护这种架构产品的人。 需求评估当前公司是否有这个能力去维护这个产品。
>
> SpringCloud-Zuul 社区活跃,基于 SrpingCloud 完整生态，但因为架构的原因（zull基于 Servlet 框架构建，采用的是阻塞和多线程方式，即一个线程处理一次连接请求，当出现问题时，如后端延迟或设备错误重试，活跃的连接和线程数量会增加，这会加大服务器负载并可能使集群无法承受）在高并发的情况下性能不高，同时需要去基于研究整合开源的适配zuul的监控和管理系统。较新的Spring cloud Gateway和zuul2倒是不错
>
> Tyk用Golang编写，并发性能较好，但一切均导向收费版本,免费版本第一次申请有一年的使用授权.没找到明确表示是否可以免费继续使用的说明.
>
> 扩展Tyk需要会Go语言，扩展Kong需要会写lua脚本，使用 zuul 还得会Java



#### 开源API GATEWAY选型

这个部分重点对比一下业界目前开源的api gateway

##### KONG

##### apisix

为什么需要apisix

![](https://segmentfault.com/img/remote/1460000040412331)

> - NGINX 和 Kong 都有各自不同应用场景；
> - NGINX 缺少官方集群统一管理方案；
> - Kong 的控制面不是完全云原生架构。



#### KONG VS APISIX

| **Features**         | **Apache APISIX** | **KONG** |
| :------------------- | :---------------- | :------- |
| **Dynamic upstream** | Yes               | Yes      |
| **Dynamic router**   | Yes               | Yes      |
| **Health check**     | Yes               | Yes      |
| **Dynamic SSL**      | Yes               | Yes      |
| **L4 and L7 proxy**  | Yes               | Yes      |
| **Opentracing**      | Yes               | Yes      |
| **Custom plugin**    | Yes               | Yes      |
| **REST API**         | Yes               | Yes      |
| **CLI**              | Yes               | Yes      |

| **Features**                                                 | **Apache APISIX**                                 | **Kong**                |
| :----------------------------------------------------------- | :------------------------------------------------ | :---------------------- |
| Belongs to                                                   | Apache Software Foundation                        | Kong Inc.               |
| Tech Architecture                                            | Nginx + etcd                                      | Nginx + postgres        |
| Communication channels                                       | Mail list, Wechat group, QQ group, GitHub, meetup | GitHub, freenode, forum |
| Single-core CPU, QPS(enable limit-count and prometheus plugins) | 18000                                             | 1700                    |
| Latency                                                      | 0.2 ms                                            | 2 ms                    |
| Dubbo                                                        | Yes                                               | No                      |
| Configuration rollback                                       | Yes                                               | No                      |
| Route with TTL                                               | Yes                                               | No                      |
| Plug-in hot loading                                          | Yes                                               | No                      |
| Custom LB and route                                          | Yes                                               | No                      |
| REST API <--> gRPC transcoding                               | Yes                                               | No                      |
| Tengine                                                      | Yes                                               | No                      |
| MQTT                                                         | Yes                                               | No                      |
| Configuration effective time                                 | Event driven, < 1ms                               | polling, 5 seconds      |
| Dashboard                                                    | Yes                                               | No                      |
| IdP                                                          | Yes                                               | No                      |
| Configuration Center HA                                      | Yes                                               | No                      |
| Speed limit for a specified time window                      | Yes                                               | No                      |
| Support any Nginx variable as routing condition              | Yes                                               | No                      |

##### Reference

1. [why api gateway](https://microservices.io/patterns/apigateway.html)
2. [来自网上的调研1](https://www.daimajiaoliu.com/daima/47973c73d100400)
3. [掘金-有数据对比](https://juejin.cn/post/6882952033712734216)

4. [API Gateways Apache APISIX and Kong Selection Comparison](https://www.apiseven.com/en/blog/api-gateways-apache-apisix-and-kong-selection-comparison)
5. [为什么需要apisix](https://segmentfault.com/a/1190000040412320)

