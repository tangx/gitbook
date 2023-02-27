#!/bin/bash

for i in $(seq 1 10); do
    {
        echo "i = $i"
    }
done

for name in zhangsan lisi wangwu zhaoliu; do
    {
        echo "name = $name"
    }
done

while True; do
    {
        date
        sleep 1
    }
done
