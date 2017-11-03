# Jajak MONGO 

## initialization

* enter mongo shell  
`Macko2:github.com artikow$ mongo`  

* create surveys collection, insert mock
```
MongoDB shell version: 3.2.11
> use jajak
switched to db jajak
> db.surveys.insert({title:"programming language",creator:"artikow@gmail.com",options:["java","go","c++"]})
WriteResult({ "nInserted" : 1 })
> db.surveys.find()
{ "_id" : ObjectId("586578e7b67a434d4954b6de"), "title" : "programming language", "creator" : "artikow@gmail.com", "options" : [ "java", "go", "c++" ] }
> db.surveys.find(creator:"artikow@gmail.com")
2016-12-30T03:58:52.038+0700 E QUERY    [thread1] SyntaxError: missing ) after argument list @(shell):1:20

> db.surveys.find({creator:"artikow@gmail.com"})
{ "_id" : ObjectId("586578e7b67a434d4954b6de"), "title" : "programming language", "creator" : "artikow@gmail.com", "options" : [ "java", "go", "c++" ]}
> db.surveys.find({creator:"artiko@gmail.com"})
> db.surveys.insert({title:"e-commerce",creator:"artiko-1@gmail.com",options:["tokopedia","salestock","shopee","blibli"]})
WriteResult({ "nInserted" : 1 })
> db.surveys.count()
2
> db.surveys.find({creator:"artiko@gmail.com"})
> db.surveys.find({creator:"artikow@gmail.com"})
{ "_id" : ObjectId("586578e7b67a434d4954b6de"), "title" : "programming language", "creator" : "artikow@gmail.com", "options" : [ "java", "go", "c++" ] }
> db.surveys.update({_id: BinData(0,"OC8/NBVXT7Gf6TB8A6ZP7Q==")},{$set: {creator:"artiko@gmail.com"}})
> db.surveys.update({_id: BinData(0,"OC8/NBVXT7Gf6TB8A6ZP7Q==")},{$set: {polls: { xyz: "bubur", abc: "roti"}}}) //reset the map
> db.surveys.update({_id: BinData(0,"OC8/NBVXT7Gf6TB8A6ZP7Q==")},{$set: {"polls.mno": "nasi goreng", "poll.tuv" : "bubur"}}) //append the map
> quit()
```