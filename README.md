# Diary for learning Go

## stack
* api interface: ~~gin gionic~~ [gorilla mux](https://github.com/gorilla/mux)
* middleware chainning: [alice](https://github.com/justinas/alice)
* persistence: ~~mysql~~ [mgo](https://gopkg.in/mgo.v2)
* pattern: [DDD](https://en.wikipedia.org/wiki/Domain-driven_design)
* unit testing: mocking interface (no framework)
* api documentation: [go-swagger](https://github.com/go-swagger/go-swagger)
* configuration: [env-config](https://github.com/kelseyhightower/envconfig)
* event consumer: [consumer](github.com/Shopify/sarama) & [group](github.com/wvanbergen/kafka/consumergroup) 
* graceful shutdown methodology: [golang-gracefully-stop](https://medium.com/@kpbird/golang-gracefully-stop-application-23c2390bb212) but with db closing included

## walkthrough
1. cd `<root-folder>/cmd/<one-of-cmd>`
1. `go install`
1. run `<one-of-cmd>` for example: `jajakhttp`
1. cd test & `go test` this will compile and run the test in folder test or alternatively if you want just re-test without re-compile the test code, then from `<root-folder>` run ```go test github.com/toshim45/jajak/test```
1. install go-swagger, generate the swagger-spec json, see [here](https://goswagger.io)
1. run `swagger serve <swagger-spec.json> --port=<swagger-port> --host=<api-ip-address>` --no-ui --no-open
1. deploy swagger ui (if you use docker, available at docker hub), configure your swagger-ui to consume this <swagger-spec.json> running in port <swagger-port>

## structure
for more info see [here](https://talks.golang.org/2014/organizeio.slide#9)