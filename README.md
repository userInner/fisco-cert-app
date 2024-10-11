# fisco-cert-app
使用fisco-bcos对盲水印图片相关信息，进行保存溯源。

1. 目录结构
    - cert 证书相关目录
        - account 用户秘钥相关目录
        - sdk 节点秘钥文件
    - conf 配置文件目录
    - controller 控制层
    - dao 数据输入相关model
        - fisco 区块链
        - mysql 数据库
        - redis 内存数据库
    - docs swagger生成的相关文档
    - logger 日志处理目录
    - logic 业务逻辑
    - middlewares 中间件
    - models 前端输入模型
    - pkg 工具包
    - router 路由
    - setting 配置文件相关模型
    - test 单元测试 - 未写完

2. 在运行之前需要做的事项
   - 搭建的一条联盟链(FISCO-BCOS 2.0)
      - https://fisco-bcos-documentation.readthedocs.io/zh-cn/latest/docs/installation.html
   - 开通腾讯-COS盲水印服务
      - https://www.tencentcloud.com/zh/document/product/436/46325
   - 搭建Redis服务
   - 搭建MySQL服务
3. 更改conf目录里的配置文件
      