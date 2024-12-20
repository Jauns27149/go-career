#!/bin/bash

# 清理函数，在接收到中断信号时调用
cleanup() {
    echo "Cleaning up..."
    rm -rf "${TEMP_DIR}"  # 使用引号以防止路径中有空格的问题
    trap - SIGINT SIGTERM  # 重置陷阱
    exit 0
}

# 设置陷阱，捕获中断信号 (Ctrl+C) 和终止信号
trap 'cleanup' SIGINT SIGTERM

CHUNK_SIZE=$((1024 * 1024 * 64))
TEMP_DIR=/dev/shm/memtest_dir  # 使用 /dev/shm 作为临时目录，因为它位于内存中
MEMTEST_FILE_BASE=${TEMP_DIR}/memtest_part  # 内存测试文件路径的基础名
TOTAL_MEM=$((9 * 1024 * 1024  * 1024))  || echo "${TOTAL_MEM}"
COUNT=$((TOTAL_MEM / CHUNK_SIZE))
a=0  # 初始化计数器
while true; do
    echo "循环开始：${a}"  # 打印当前的 i 值
    ((a++))  # i 自增 1
    mkdir -p ${TEMP_DIR}

    # 循环分配小块内存
    for i in $(seq 1 $COUNT); do
        dd if=/dev/zero of=${MEMTEST_FILE_BASE}${i} bs=${CHUNK_SIZE} count=1 > /dev/null 2>&1 &
    done
    # 等待所有 dd 进程完成
    wait
    # 强制刷新缓存并清除未使用的缓存
    sync
    sudo sh -c 'echo 3 > /proc/sys/vm/drop_caches'
    rm -rf ${TEMP_DIR}
done