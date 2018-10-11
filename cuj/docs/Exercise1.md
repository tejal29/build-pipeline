# Definition
In this exercise, we will go thourgh a pipeline defined for the [Hipster Shop](https://github.com/GoogleCloudPlatform/microservices-demo)

Hipster Shop, contains a 10-tier microservices application. The application is a web-based e-commerce app called “Hipster Shop” where users can browse items, add them to the cart, and purchase them.

For this exercise, you do not need to know the specifics of each service.
We will be concentrating on 3 of the 10 microservices in this exercise. 
1. Product Catalog Service
2. Checkout Service
3. Frontend Service


Now, take a look the [pipeline](./../pipeline/hipster-pipeline.yaml) which references above 3 services.

# Overview of files
1. [Tasks](./../pipeline/build-push-task.yaml) Contains the definitions of tasks used in the pipeline.
2. [PipelineParams](./../pipeline/pipelineparams.yaml) Defines pipeline parameters.
3. [Resources](./../pipeline/resources.yaml) Defines resources used in the pipeline.



