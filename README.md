## TM

TM is a professional web benchmark tools built with Golang, which supports load-testing of services with the HTTP traffic file(.har). With the power of Goroutine, it has a nice loading-test performance even in standalone testing.

What's more, it has a visually command-line UI interface, with it you can see the loading-test results directly, which is very convenient for users to use.

Compared to traditional loading-test tools, it can parse our custom traffic files, generate the loading-test code, and visit the API server with the order of traffic in the file so that it can perform very similarly to production-env. With such a feature,  it is not necessary for users to write much code before testing, what they need to do is just visit the API service, make some actions and generate some business requests and record the traffic, and then parse the result with the TM tool we provide.

In many cases, a single machine can't achieve the effect we want, and we need multiple machines to test the target service. TM also supports distributed deployment, which makes the services to be tested in a cluster way with a few simple commands, you can quickly complete the construction of a distributed testing environment.

Of course, we also provide some convenient tools: such as automatically generating a distributed deployment testing file for k8s; providing a website to show the testing result, and generating the test report for the testing.

