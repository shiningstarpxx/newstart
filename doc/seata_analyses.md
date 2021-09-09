## SEATA 中的SAGA 模式代码分析

#### author: peixingxin(michaelpei)

#### email: xingxinpei@gmail.com(michaelpei@tencent.com)



### 代码组织结构

#### SAGA

整个代码的最外层目录我们从SAGA目录看起,  代码版本是： 

```xml
        <!-- seata version -->
        <revision>1.5.0-SNAPSHOT</revision>
```

#### seats-saga-engine

从目录结构上看是核心engine的部分， 从内容上看主要是各种数据结构 的定义

##### evaluation

##### exception

##### expression

##### impl

##### invoker



