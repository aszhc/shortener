# 短链接


## 查看短链接

### 缓存版

有两种方式，
1. 使用自己实现的缓存，     surl -> lurl，能够节省缓存空间，缓存数据量小
2. 使用go-zero自带的缓存， surl -> 数据行，不需要自己实现，开发量小

这里使用第2种方案：
1. 添加缓存配置
- 配置文件
- 配置config结构体
2. 删除旧的model层代码
- 删除 shorturlmapmodel.go
3. 重新生成model层代码
```bash
goctl model mysql datasource -url="root:root@tcp(127.0.0.1:3306)/shortener" -table="short_url_map"  -dir="./model" -c
```
4. 修改svccontext层代码
