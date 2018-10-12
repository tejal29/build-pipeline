# Definition
In this exercise, we will define a task which executes `go test` to test your go code. 
This task will run `go test` command for your go project.

You can use the go runtime image `gcr.io/google_appengine/golang`

To run all tests inside inside this container image for the [golang/archive](https://golang.org/pkg/archive/) package, 

```shell
docker run -i gcr.io/google_appengine/golang go test archive/...
ok  	archive/tar	0.029s
ok  	archive/zip	7.997s
```


