#!/bin/bash

# 定义要检查的端口列表
ports=(5775 6831 6832 5778 16686 14268 9411 9090 3001 9200 9300 5601 2181 9092 8980 33069 36379)

# 循环遍历端口数组
for port in "${ports[@]}"; do
    # 使用ss命令检查端口
    result=$(sudo ss -tuln | grep ":$port ")
    if [ -z "$result" ]; then
        echo "Port $port is available."
    else
        echo "Port $port is in use:"
        echo "$result"
    fi
done

