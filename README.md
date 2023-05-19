#### otter同步钉钉报警工具

这是一个简单的otter同步钉钉报警工具
使用如下:
```shell
./otter-dingtalk  --zk zk地址:2181 \
                  --mysql_host ottermysql地址:3306 \
                  --mysql_pass 'xxxxxxx' \
                  --token 钉钉机器人token \
                  --secret 钉钉机器人secret \
                  -c 29 -c 30 
```
- -c参数为指定监控的channelID ,可以为all，也可以分开指定要监控的channelID

![img](images/1.png)
