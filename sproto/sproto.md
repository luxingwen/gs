引用：<https://github.com/xjdrew/gosproto>


采用大端字节序

数据包

| len   | data  |
| ----- | ----- |
| 4字节 | bytes |



data结构

| mode                       | type   | data       |
| -------------------------- | ------ | ---------- |
| 1字节                      | 2字节  | bytes      |
| 0表示request 1表示response | 协议id | 数据包内容 |





