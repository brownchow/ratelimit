# Token_Bucket 算法实现限流



https://en.wikipedia.org/wiki/Token_bucket

一般上线一个服务之后，都会预估这个服务能承担多少请求，当单位时间内请求数量超过阈值后，就应该拒绝多余的请求。一般用 token bucket 算法和 leaky bucket 算法实现上述功能



## TODO：

- 并发控制，对共享变量加锁

## Reference

https://github.com/juju/ratelimit

https://www.jianshu.com/p/3df8c3b38fc0

