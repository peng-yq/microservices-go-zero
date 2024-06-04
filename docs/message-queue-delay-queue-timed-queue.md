## 消息队列、延时队列和定时队列

### kafka

[Apache Kafka](https://kafka.apache.org/)是一个开源的分布式事件流平台，用于构建实时数据管道和流式应用程序。它被设计用于处理高吞吐量、容错性强和可扩展的数据流。

`Kafka`的主要特点包括：

1. **发布-订阅消息系统**：`Kafka`遵循发布-订阅模型，生产者将消息发布到主题，消费者订阅这些主题以接收消息。

2. **可扩展性**：`Kafka`可水平扩展，通过添加更多的代理到`Kafka`集群来轻松扩展规模。

3. **容错性**：`Kafka`在多个代理之间复制数据以确保容错性。如果一个代理失败，数据仍可以从其他代理上的复制数据中检索。

4. **高吞吐量**：`Kafka`能够处理每秒大量的消息，适用于需要实时数据处理的用例。

5. **数据保留**：`Kafka`会保留消息一段可配置的时间，允许消费者检索历史数据。

6. **流处理**：`Kafka Streams API`支持对存储在`Kafka`主题中的数据进行实时流处理和分析。

7. **连接器**：`Kafka Connect`允许与外部系统轻松集成，将数据导入/导出到`Kafka`中。

本项目中`filebeat`收集日志后，将消息发送至`kafka`中，再由`go-stash`消费`kafka`日志，最后输出到`es`中。此外，使用`go-queue-kq`基于`kafka`实现订单支付状态更新的消息队列。

### go-queue

[GitHub仓库](https://github.com/zeromicro/go-queue)

`go-queue`是`go-zero`官方提供的消息队列组件，可分为`kq`和`dq`：

- `kq`是基于`kafka Pub & Sub`框架的消息队列
- `dq`是基于[beanstalkd的延迟队列](https://github.com/beanstalkd/beanstalkd)

项目使用`go-queue-kq`实现根据第三方支付订单支付状态更新修改订单支付状态的消息队列。`order-mq`为消费者的实现，从`kafka`特定`topic`中取出第三方支付消息进行消费，并修改订单支付状态。`payment`中实现了生产者（`pusher`），将第三方支付消息推送至`kafka`特定`topic`中。

`go-queue`不支持定时队列。

[关于beanstalkd的中文详解](https://learnku.com/articles/69985)

[go-queue-消息队列教程](https://go-zero.dev/docs/tutorials/message-queue/kafka)

[go-queue-延时队列教程](https://go-zero.dev/docs/tutorials/delay-queue/beanstalkd)

### asynq

`Asynq`是一个`Google`提供的`Go`库，用于对任务进行排队并使用工作程序异步处理任务。`Asynq`直接基于`Redis`，可扩展且易于上手。前期如果业务量不大时可以直接使用`Asynq`，节省一个中间件。项目使用了`asynq`实取消超时未支付订单的延时队列和订阅消息通知用户订单状态更新的消息队列，以及定时队列的示例。

<img src="https://user-images.githubusercontent.com/11155743/116358505-656f5f80-a806-11eb-9c16-94e49dab0f99.jpg">

[Asynq Github官方提供了较为详解的例子和教程](https://github.com/hibiken/asynq)

### servicegroup

[使用servicegroup管理服务](../app/order/cmd/mq/README.md)