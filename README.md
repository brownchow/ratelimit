# Token_Bucket 算法实现限流



https://en.wikipedia.org/wiki/Token_bucket

一般上线一个服务之后，都会预估这个服务能承担多少请求，当单位时间内请求数量超过阈值后，就应该拒绝多余的请求。一般用 token bucket 算法和 leaky bucket 算法实现上述功能

在分布式系统中也有应用，限制某个微服务的相应速率

Token Bucket 支持突发流量





## TODO

- 并发控制，对共享变量加锁



## Others

Nginx 对某个 IP 限制请求速率

iptables 对某个 ip 限流， hashlimit, limit







## Reference

https://github.com/juju/ratelimit

https://www.jianshu.com/p/3df8c3b38fc0

