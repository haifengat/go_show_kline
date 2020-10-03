# show_kline

#### 介绍
K线数据展示

#### 软件架构
golang beego echarts gocharts


#### 安装教程
```yaml
version: "3.7"
services:
    go_xml_tick:
        image: haifengat/go_show_kline
        container_name: go_show_kline
        restart: always
        port: 
            # 宿主端口
            - 8080=8080
        environment:
            - TZ=Asia/Shanghai
            # postgres  future_min配置
            - pgConfig=postgres://postgres:123456@172.19.129.98:15432/postgres?sslmode=disable
```
