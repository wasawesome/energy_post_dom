#### Energy Post

A technical project, in order to record what I have seen in these two years.(2017-)

For project framework from font-end to back-end which has four project for them. Include Web client, Mini Programe client, node server middleware and Golang server.

#### Project base technical framework

- Web client: Vue multiple page application. (include adamnistrator monitoring system)
- Mini Program: MPVue
- Node Middleware: Nest.js
- Back-end: Golang for Beego


#### Commands

---

```
- bee run // start project server

- bee run -gendoc=true -downdoc=true  // Generate Swagger Doc

- docUrl: http://localhost:8090/swagger/

```

#### Admin monitoring
> app.conf:

``` app.conf
# Open admin monitoring and setup host:port

EnableAdmin = true
AdminAddr = "localhost"
AdminPort = 8088

```

#### checking

> To check the number of api reqeusts
[http://localhost:8088/qps](http://localhost:8088/qps)

---


> #### [Official Introduction](https://beego.me/docs/advantage/monitor.md)
