## 简介

欢迎使用易盾 Golang 开发者 SDK，SDK 是易盾 API 的配套工具。当前已经支持内容安全与业务安全产品(产品介绍见[官网](https://support.dun.163.com/?locale=zh-CN))，后续其他产品会逐渐添加支持。
SDK 中各接口调用方式基本相同，接入方式统一，并提供了一些接口调用过程中常见的异常处理方案。
为了简化 Golang 开发者调试和接入易盾 API 的成本，这里向您介绍适用于 Golang 的 SDK 的使用流程，并提供首次使用 SDK 的简单示例。
让您快速了解 SDK 的特性和功能，并快速使用起来。

### 1. 准备工作

#### 1.1 环境要求

* 建议Go: 1.2+

  > **说明** 您可以执行命令`go version` 查看本地golang版本。
  >
* 从 易盾控制台 开通账号，套餐以及相应产品和业务。
* 从 易盾控制台 获取 SecretID、SecretKey、BusinessId，如果有疑问请联系您的商务经理

#### 1.2 安装SDK

* 开发环境使用go mod管理依赖。
* 在您的项目中添加module依赖，只需在 go.mod 文件中添加以下依赖项即可。

  ```
    require github.com/yidun/yidun-golang-sdk v1.0.1
  ```

  > 注意： 这里的版本号只是举例，您可以在 git 仓库 上找到最新的版本。
  >

#### 2. 使用SDK

请参考demo模块下各接口使用方式

#### 3. 功能特性

- 完善的失败处理：

  - 接口请求提供同步重试机制，在出现调用失败时，会自动重试指定 region 内的下一个请求节点。

  - 接口请求支持熔断降级，在当前时间窗口内，指定 region 的某个节点失败请求次数到达阈值时，会触发熔断，后续请求会自动路由到region内的其他可用节点，直到超过熔断时间再对熔断节点探活。
- 丰富的签名算法：

  - 内置多种接口签名算法支持，包括MD5、SHA1、SHA256、SM3
- 灵活的接口配置：

  - 常用的http请求的各项配置，包括接口超时参数，http 连接池连接数、空闲时间等核心参数。

  - 切换请求协议，通过在请求对象中设置 protocol 字段，切换 HTTP/HTTPS 协议，默认 HTTPS。

  - 指定 region（默认杭州），根据您服务器的所在地，选择适合的 region，以达到更好的请求效果，支持的 region 列表请咨询您的商务经理。

  - 熔断窗口的各项配置，包括窗口大小、熔断时间、错误率阈值、熔断生效的请求数阈值。

#### 4. 配置支持

* http请求各项参数配置

  ```
   // 按照需求设置http请求的多个参数配置
   credential := auth.NewCredentials("accessKeyId", "accessKeySecret")
   clientProfile := client.NewClientProfile(credential)
   clientProfile.HttpClientConfig.ConnectionKeepAliveMillis = 10000
   crawlerSubmitClient := crawler.NewCrawlerClient(clientProfile)
  ```
* 请求协议
  ```
  // 覆盖默认请求协议
  request := submit.NewCrawlerResourceV3SubmitRequest()
  request.SetProtocol(http.ProtocolEnumHTTPS)
  ```
* 熔断配置

  ```
  credential := auth.NewCredentials("accessKeyId", "accessKeySecret")
  clientProfile := client.NewClientProfile(credential)
  // 按照需求设置固定窗口的多个熔断参数配置
  clientProfile.BreakerConfig.SetStatWindowMillis(10000)
  crawlerSubmitClient := crawler.NewCrawlerClient(clientProfile)
  ```
* 海外节点配置

  ```
  credential := auth.NewCredentials("accessKeyId", "accessKeySecret")
  clientProfile := client.NewClientProfile(credential)
  // 以访问新加坡节点为例，支持的 region 列表请咨询您的商务经理
  clientProfile.SetRegionCode("sg-singapore")
  livevideosolution.NewLiveVideoSolutionClient(clientProfile)
  ```

* 重试配置

  ```
  credential := auth.NewCredentials("accessKeyId", "accessKeySecret")
  clientProfile := client.NewClientProfile(credential)
  // 最大10次
  clientProfile.SetMaxRetryCount(10)
  crawlerSubmitClient := crawler.NewCrawlerClient(clientProfile)
  ```

#### 5. 详细文档
- [wiki链接](https://github.com/yidun/yidun-golang-sdk/wiki)
