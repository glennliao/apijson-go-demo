access_ext
- ? deleteCheck // 删除检查

request_ext
- then 成功后续  清理缓存、发送消息等
- catch 失败后续

- before  操作前处理 (数据处理(function/js)) + 特殊权限校验
- after  操作前处理 成功/失败 都会执行的操作

