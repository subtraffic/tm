## tm

tm是一款专业的由golang开发的http流量压测工具，在goroute的加持下，即使在单机的测试中，它也有非常不错的压测表现。

它有可视化的命令行ui界面，可以直接在命令行界面看到相应的压测结果，非常方便用户使用，该命令行界面由termui来实现。

此外，相对于传统的压测工具，它可以处理我们自定义的流量文件，生成相应的压测代码，还能根据流量的顺序对服务器进行压测，来达到非常真实的压测效果。有了这样的功能，用户压测试时，就不需要编写太多压测的代码，只需要访问被压测的服务，发起一些业务请求，并录制好流量，之后调用我们提供的tm工具即可。

在很多时候，单机的压测达不到我们想要的压测效果，我们需要多台机器对目标服务进行压力测试。而tm也支持分布式的压测部署方案，以集群的方式，对需要被压测的服务进行测试；只需要简单几个命令，就可以快速完成分布式压测环境的搭建。

当然我们还提供了一些方便的工具：比如自动生成基于k8s的分布式压测部署方案；提供ui界面的压测效果展示页面；将结果生成对应的测试报告。

