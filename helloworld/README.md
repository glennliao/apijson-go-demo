# hello-world

## 最简demo
```go
package main

//...

func main() {
    framework.Init()
    
    config.AccessVerify = false // 全局配置验证权限开关
    
    s := g.Server()
    
    s.Group("/", handler.Bind)
    
    s.Run()
}
```

启动后访问: http://127.0.0.1:8090/get (Post application/json)

> 可以使用postman或类似相关的工具访问

## example
1. 导入db.sql
2. 将config.toml.example 重命名成config.toml , 然后修改下config.toml的database配置
3. 执行test.http的例子