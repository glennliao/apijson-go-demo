
access_ext
- ? idGen // id生成策略
- ? deleteCheck // 删除检查

request_ext
- tag
- version
- method
- hooks
- after 成功后续  清理缓存、发送消息等
- fail  失败后续
- before  操作前处理 (数据处理(function/js)) + 特殊权限校验

function
- debug
- name (唯一)
- arguments
- demo
- detail
- type
- version 允许操作的最低版本
- tag(允许的操作)
- methods 允许的操作   // 使用 requestIdList 替代 version,tag, methods
- back  返回值示例  ,  系统启动时校验demo是否正常通过(demo 是不是得多个)



