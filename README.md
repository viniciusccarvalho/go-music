go-music
========

A port of [spring-music](https://github.com/scottfrederick/spring-music) to golang


Simple app with angularjs, golang and mongodb using [mgo](http://labix.org/mgo).

Mostly an exercise as I learn #golang.

###Running on CloudFoundry

All you need is a mongodb service bound to your app. The application will search VCAP_SERVICES and detect any service with tags "document", then just push it to CF:

```
cf push go-music -b https://github.com/cloudfoundry/buildpack-go.git -m 128M
```

Check it [live](http://go-music.cfapps.io) 

###Running on your local machine

Just build the server and execute it, by default it will run on port 9000, but you can choose the port by exporting the PORT variable

