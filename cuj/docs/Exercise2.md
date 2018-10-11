# Definition
In this exercise, we will define a task which executes `go test` to test your go code. 
This task will run `go test` command for your go project.

Look at the sample [build-push task](./../build-push-task.yaml). 
You can use the go runtime image `gcr.io/google_appengine/golang`
e.g go test command
To run all tests in the kritis project, 

```shell
go test github.com/grafeas/kritis/...
```


