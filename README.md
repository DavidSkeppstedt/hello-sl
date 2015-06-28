#Hello-SL
Backend webservice which serves a REST-API based on information from SL:s API. 

Written in Go using standardlibs like net/http etc.


At the moment it supports:

+ [SL-Platsuppslag](https://www.trafiklab.se/api/sl-platsuppslag) json is served at ```/place?search="<search string>" ```

Plan is to also support:

+ [SL Realtidsinformation 3](https://www.trafiklab.se/api/sl-realtidsinformation-3)

##API-Keys
###Basic setup
Create a map called **keys** and place it the root of the project. 
### API 1: SL Platsuppslag
Inside **keys** you would create a file called _plats.key_ and paste your api key  which you can fetch at [trafiklab](http://trafiklab.se).

The key should be on the first row/line and be consistent. No spaces or line-breaks.

##First time start
Based on the assumption that you have setup Go in the right way and you have cloned this project into your workspace.

You should be able to run:

###SL-Platsuppslag 
```go run places/*.go```
you must stand in the root project folder for it to start. It would otherwise complain that it can not find the keys folder. 

Then you shoulde be able to visit ```http://localhost:8080/place?search=%22slussen%22``` and get something like this returned ```{"places":[{"Name":"Slussen (Stockholm)","SiteId":"9192"},{"Name":"Stora mossen (Stockholm)","SiteId":"9111"},{"Name":"Bagarmossen (Stockholm)","SiteId":"9141"},{"Name":"Blekholmsterrassen (Stockholm)","SiteId":"1006"},{"Name":"Hökmossen (Stockholm)","SiteId":"1630"},{"Name":"Hökmossens gård (Stockholm)","SiteId":"1787"},{"Name":"Terrassen (Stockholm)","SiteId":"1535"},{"Name":"Mossens idrottsplats (Stockholm)","SiteId":"3646"},{"Name":"Musseronvägen (Vallentuna)","SiteId":"2516"},{"Name":"Prästgårdsmossen (Värmdö)","SiteId":"4429"}]}```