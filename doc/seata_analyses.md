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

这里是state machine实现的地方，整体来说，代码写的是比较流畅的，不过单体函数过长，单个文件过大，不确定java的风格是否就是这样。下面重点看看几个核心的函数

* 第一个是启动函数

```java
private StateMachineInstance startInternal(String stateMachineName, String tenantId, String businessKey,
                                               Map<String, Object> startParams, boolean async, AsyncCallback callback)
            throws EngineExecutionException {
      ...
    }
}

// 
```

* 接下来是创建状态机

  ```java
  private StateMachineInstance createMachineInstance(String stateMachineName, String tenantId, String businessKey,
                                                         Map<String, Object> startParams) {
    
    //
  ```



* 接下来是驱动状态机前进

  ```java
  protected StateMachineInstance forwardInternal(String stateMachineInstId, Map<String, Object> replaceParams,
                                                     boolean skip, boolean async, AsyncCallback callback)
          throws EngineExecutionException {
  
  ```

* 如果失败了自然要开始补偿

  ```java
  public StateMachineInstance compensateInternal(String stateMachineInstId, Map<String, Object> replaceParams,
                                                     boolean async, AsyncCallback callback)
          throws EngineExecutionException {
  
  ```

* 如果server挂了，能否恢复呢？目前看，有个恢复的接口

  ```java
   public StateMachineInstance reloadStateMachineInstance(String instId) {
  
     // @todo 
  ```

* 同样，要判断server的状态

  ```java
   protected boolean checkStatus(StateMachineInstance stateMachineInstance, ExecutionStatus[] acceptStatus,
                                    ExecutionStatus[] denyStatus, ExecutionStatus status, ExecutionStatus compenStatus,
                                    String operation) {
  ```

  

##### invoker



