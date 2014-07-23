go-music
========

A port of [spring-music](https://github.com/scottfrederick/spring-music) to golang


Working just as the java counter part. It only supports mongodb. Automatic service binding using VCAP_SERVICES

```
cf push go-music -b https://github.com/cloudfoundry/buildpack-go.git -m 128M
```

Check it [live](http://go-music.cfapps.io) 
