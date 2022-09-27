# kafkacli
----------------
一个简单的kafkacli工具希望能够帮你在日常的运维过程中快速确认kafka集群是否可以正常生产消费消息 🔧
更详尽的kafkacli相关工具请移步 [Kafkactl](https://github.com/deviceinsight/kafkactl)

此工具的诞生只是为了展示生产消息，及消费消息到kafka的耗时，同时也可以测试Kafka集群是否正常。

## 安装

您可以通过下载源代码，自行编译，也可以直接下载我编译后的二进制文件.

### 安装预编译的二进制文件


从[发布页面](https://github.com/vlinux/kafkacli/releases)下载预编译的二进制文件并复制到所需位置。


### 从源代码编译

```bash'
git clone git@github.com:vlinux/kafkacli.git
cd kafkacli && vim Makefile
#修改build中您自己的操作系统和内核情况
#例：linux amd64
build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o ${BINARY} main.go
#修改后执行make可编译二进制文件
make help && make
```

## 使用说明

该Cli包含的子命令非常简单，所有使用方法均在这里：

#### 生产消息

可以通过以下方式来生产消息,例kafka地址为192.168.1.1:9092,发送2条消息：

> 支持单borker也支持多borker，例：-s 192.168.1.1:9092.192.168.1.2:9092,192.168.1.3:9092
>
> 消息体已内定，默认消息的key为`hello`，value为`vlinux`,如果你会go开发，可以自行修改
>
> Topic已内定，默认topic为test，程序可以自动创建，默认打印消费耗时
>
> kafka测试版本：V0_11_0_1

```bash
./kafkacli produce -s 192.168.1.1:9092 -n 2
```

#### 消费消息

可以通过以下方式使用来自主题的消息：

>默认从最新一条开始消费，默认打印消费耗时

```bash
./kafkacli conmuser -s 192.168.1.1:9092
```

### 帮助

```bash
./kafkacli help
```



