# goBlog

    这是采用golang + iris 框架的博客

# 博客架构


# 使用框架
用了以下框架

```
go get -u github.com/kataras/iris

```



# 目录结构

```bash
|
├─api			# 接口相关
├─web			# web服务代码
│  ├─controllers	# 控制器
│  ├─middleware 	# 中间件
│  ├─modules		# 数据模型
│  └─routers		# 路由
├─bin			# 编译好的binary
├─build			# 编译部署需要的相关文件，如dockerfile之类
├─cmd			# 程序命令，入口之类
├─configs		# 配置文件
├─docs			# 相关具体文档
├─init			# 全局初始化工作
├─static		# web相关文件
│  ├─style			# css
│  ├─img			# 图片
│  └─script			# js
├─storages		# 存储相关
│  ├─cache			# 缓存
│  ├─databases		# 数据库操作
│  └─logs			# 日志相关处理
├─test			# 测试
├─utils			# 相关工具集
└─views			# html文件
```





