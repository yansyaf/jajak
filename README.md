# Repository for learning Go & Mongodb

## stack
* api interface: ~~gin gionic~~ gorilla mux
* middleware chainning: alice
* persistence: mgo (mongodb)
* pattern: handler - domain (DDD)
* unit testing: mocking interface (no framework)

## walkthrough
1. cd `<root-folder>`
1. `go build`
1. cd test & `go test` this will compile and run the test in folder test or alternatively if you want just re-test without re-compile the test code, then from `<root-folder>` run ```go test github.com/toshim45/jajak/test```