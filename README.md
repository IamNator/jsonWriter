# jsonWriter

This package writes a json response to http.ResponseWriter. 

## Getting Started

Use the command "go get github.com/IamNator/jsonWriter" to download and setup the package in your go development environment.

### Prerequisites

go 1.1 or greater 

```
go get github.com/IamNator/jsonWriter
```

### Installing

A step by step series of examples that tell you how to get a development env running

Say what the step will be

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo

## Running the tests

To run tests for this package use the go command below   
while in pkg directory -- go test -v 


### Break down into end to end tests

Explain what these tests test and why
if everything is working perfectly you should see     
    
=== RUN   TestError   
    JsonError_test.go:28: Status Code Test Passed : 200   
    JsonError_test.go:35: Json Body test Passed : {"error":"error"}    
--- PASS: TestError (0.00s)    
=== RUN   TestResponse   
    JsonResponse_test.go:25: Status Code Test Passed : 200   
    JsonResponse_test.go:33: Json Body test Passed : {"title":"error","message":"error"}   
--- PASS: TestResponse (0.00s)   
PASS    
ok      github.com/IamNator/JsonWrite   0.818s   



```
Give an example
```

## Built With

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Maven](https://maven.apache.org/) - Dependency Management
* [ROME](https://rometools.github.io/rome/) - Used to generate RSS Feeds

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/IamNator/jsonWriter/tags). 

## Authors

* **Nator Verinumbe** - *Initial work* - [IamNator](https://github.com/IamNator)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Inspiration - Writing json Error response for API
