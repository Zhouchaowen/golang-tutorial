# 目录
- ch_1 channel 定义
- ch_2 channel 阻塞发送 
- ch_3 小实验: 通过 goroutine+channel 求数组和 
- ch_4 非缓冲和缓冲 channel 的对比
- ch_5 遍历和关闭 channel 
- ch_6 channel+select 控制 goroutine 退出

## Channel
channel 是解决 goroutine 的同步问题以及 goroutine 之间数据共享（数据传递）的问题



|   操作   | 一个零值nil通道     |  一个非零值但已关闭的通道    | 一个非零值且尚未关闭的通道|
| ---- | ---- | ---- | ---- |
|  关闭 |  产生恐慌    |  产生恐慌    | 成功关闭|
|   发送数据   | 永久阻塞     |  产生恐慌    | 阻塞或者成功发送|
|   接收数据   |   永久阻塞   |   永不阻塞   |阻塞或者成功接收 |


