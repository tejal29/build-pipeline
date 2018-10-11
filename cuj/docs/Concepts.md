# Pipeline CRD
Pipeline CRD is an open source implementation to configure and run CI/CD style pipelines for your kubernetes application.
Pipeline can be expressed as a DAG of tasks. A task can be any process you would want to run as part of your continous integration process. 
A task will run as a process inside a container on your cluster.

Pipeline CRD builds on top of [Knative Build CRD](https://github.com/knative/docs/blob/master/build/builds.md) which relies on Kubernetes Custom Build Resource, to create and run on-cluster processes to completion.
A custom resource is an extension of the Kubernetes API that defines an API object of a certain kind.
Read about the [Pipeline CRDs](../../README.md#CRDs) k8s-style resources.


