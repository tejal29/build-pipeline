# Definition
In this exercise, we will define a pipeline for [checkoutout service](https://github.com/GoogleCloudPlatform/microservices-demo/tree/master/src/checkoutservice)

This pipeline will consists of 3 tasks
1. Run unit tests for [checkoutout service](https://github.com/GoogleCloudPlatform/microservices-demo/tree/master/src/checkoutservice).
You can run those by running:
```shell
go test github.com/GoogleCloudPlatform/microservices-demo/src/checkoutservice/...
```
2. Build the checkout service and finally
3. Deploys the checkout service in test cluster
