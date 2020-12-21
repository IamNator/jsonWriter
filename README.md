# jsonWriter

This package writes a json response to http.ResponseWriter. 

## Getting Started

Use the command "go get github.com/IamNator/jsonWriter" to download and setup the package in your go development environment.

### Prerequisites

go 1.1 or greater, to check your go version type the following commmand in your terminal

```
go version
```

### Installing

A step by step series of examples that tell you how to get a development env running

Use the "go get command"

```
go get github.com/IamNator/jsonWriter
```

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


### Examples
```
package main

import (
	"github.com/IamNator/jsonWriter
)

    err := json.NewDecoder(req.Body).Decode(&user)
    check(err)
    
	if user.Email == "" || user.PassWord == "" {
		JsonError(w, "please Fill in fields", http.StatusBadRequest)
		return
	}
    
//This write the json {"error":"please Fill in Fields"} as a response.
//For more documentation use "go doc -all " 

```

## Built With

* [net/http and encoding/json ]() - go dependencies used
* [go mod]() - Dependency Management


## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/IamNator/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/IamNator/jsonWriter/tags). 

## Authors

* **Nator Verinumbe** - *Initial work* - [IamNator](https://github.com/IamNator)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Inspiration - Writing json Error response for API
