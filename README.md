# Repository for learning Go & Mongodb

## stack
* api interface: ~~gin gionic~~ [gorilla mux](https://github.com/gorilla/mux)
* middleware chainning: [alice](https://github.com/justinas/alice)
* persistence: [mgo](https://gopkg.in/mgo.v2)
* pattern: [DDD](https://en.wikipedia.org/wiki/Domain-driven_design)
* unit testing: mocking interface (no framework)
* api documentation: [go-swagger](https://github.com/go-swagger/go-swagger)

## walkthrough
1. cd `<root-folder>`
1. `go build`
1. cd test & `go test` this will compile and run the test in folder test or alternatively if you want just re-test without re-compile the test code, then from `<root-folder>` run ```go test github.com/toshim45/jajak/test```
1. install go-swagger, generate the swagger-spec json, see [here](https://goswagger.io)
1. run `swagger serve <swagger-spec.json> --port=<up-to-u> --host=<api-ip-address>`