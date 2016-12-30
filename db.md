#Jajak MONGO 

##initialization

* enter mongo shell  
`Macko2:github.com artikow$ mongo`  

* create poll collection, insert mock
```
MongoDB shell version: 3.2.11
connecting to: test
Server has startup warnings: 
2016-12-30T03:53:58.704+0700 I CONTROL  [initandlisten] 
2016-12-30T03:53:58.704+0700 I CONTROL  [initandlisten] ** WARNING: soft rlimits too low. Number of files is 256, should be at least 1000
> use jajak
switched to db jajak
> db.poll.insert({title:"programming language",creator:"artikow@gmail.com",items:["java","go","c++"]})
WriteResult({ "nInserted" : 1 })
> db.poll.find()
{ "_id" : ObjectId("586578e7b67a434d4954b6de"), "title" : "programming language", "creator" : "artikow@gmail.com", "items" : [ "java", "go", "c++" ] }
> db.poll.find(creator:"artikow@gmail.com")
2016-12-30T03:58:52.038+0700 E QUERY    [thread1] SyntaxError: missing ) after argument list @(shell):1:20

> db.poll.find({creator:"artikow@gmail.com"})
{ "_id" : ObjectId("586578e7b67a434d4954b6de"), "title" : "programming language", "creator" : "artikow@gmail.com", "items" : [ "java", "go", "c++" ] }
> db.poll.find({creator:"artiko@gmail.com"})
> db.poll.insert({title:"e-commerce",creator:"mefisya@gmail.com",items:["tokopedia","salestock","shopee","blibli"]})
WriteResult({ "nInserted" : 1 })
> db.poll.count()
2
> db.poll.find({creator:"artiko@gmail.com"})
> db.poll.find({creator:"artikow@gmail.com"})
{ "_id" : ObjectId("586578e7b67a434d4954b6de"), "title" : "programming language", "creator" : "artikow@gmail.com", "items" : [ "java", "go", "c++" ] }
> quit()
```