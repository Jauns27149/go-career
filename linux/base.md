# 结构

- /
  - var : 存放可变数据文件(动态)
    - lib: 存放与特定应用程序相关的状态信息和数据文件
    - log : 存放系统和应用程序的日志文件
  - etc : 存放系统的配置文件(静态)
    - os-release : 存储操作系统的识别信息(key-value)
    - shadow : 存储用户密码和其他与用户帐户相关的安全信息的重要文件
  - opt : 存放可选的第三方软件包
  - tmp : 存储临时文件和数据(temporary)
  - dev : 表示设备文件
    - /null : 任何写入 `/dev/null` 的数据都会被丢弃

# 命令

## 目录&文件操作

### journalctl

```bash
journalctl [OPTIONS...] [MATCHES...]
# 访问和查看由 systemd 管理的系统日志，包含系统启动信息、服务日志、内核消息、应用程序日志等
```

- options : 

  - -S --since=DATE : 显示date之后的日志

  - -U --until=DATE : 显示date之前的日志

  - -u --unit=UNIT : 查看特定服务的日志

  - -f : 实时查看日志

  - xe: 查看当前登录会话的日志

  - -p : 按日志的优先级（日志级别）进行过滤

    ```bash
    0 - emergency（紧急）
    1 - alert（警报）
    2 - critical（严重）
    3 - error（错误）
    4 - warning（警告）
    5 - notice（通知）
    6 - info（信息
    7 - debug（调试）
    ```

    

### find

```bash
find [path] [expression]
# 在文件系统中搜索文件和目录，默认目录是当前目录
```

- expression

  - -name : 指定名字文件，可以使用正则

  - -type : 指定类型

    ```bash
    d : 目录
    f :	文件
    l : 符号连接
    ```

  - size : 按大小寻找

    ```bash
    +100M : 大于100M
    -100M : 小于100M
    ```

  - -exec : 对找到的文件执行某个命令

### grep

```bash
grep [options...] "pattern" filename
```

- options :
  - -i : 忽略大小写
  - -n : 显示行号
  - -v : 反向匹配（不显示匹配的行）
  - -w : 匹配完整单词
  - -c : 显示匹配行的数量
  - -l : 显示包含匹配模式的文件名
  - -r  : 递归搜索目录
  - -E : 使用扩展正则表达式
  - -P : 使用 Perl 正则表达式。

### 基本正则

- .（点）: 匹配任何单个字符（除了换行符），a.c 可以匹配 "abc"、"axc" 
- `*`（星号）: 匹配前面的字符或子表达式零次或多次,，a*b 匹配 "b"、"ab"、"aab"、"aaab" 等
- ^（脱字符）: 匹配行的开头，^abc 匹配以 "abc" 开头的行
- $（美元符号）: 匹配行的结尾，abc$ 匹配以 "abc" 结尾的行
- []（方括号）: 匹配方括号内的任何一个字符，可以使用范围， [a-z] 匹配任何小写字母，[abc] 匹配 "a", "b" 或 "c"
- [^]（否定方括号）: 匹配不在方括号内的任何字符，`[^0-9] `匹配任何非数字字符
- \（反斜杠）: 用于转义特殊字符，使其失去特殊含义，`\.` 匹配一个实际的句点，而不是“任意字符”。

### cp

```
cp [OPTIONS...] source destination
```

- options :
  - -r : 递归
  - -f : 强制
  - -v : 输出过程
  - -u : 更新

### wc

```bash
wc [options] [flie...]
# word count 用于统计文件的字数、行数、字符数和字节数
```

- options
  - `-l` :  统计行数（Line count）
  - `-w : 统计单词数（Word count）
  - `-c : 统计字符数（Character count）
  - `-m` : 统计字符数（按字符计）
  - `-L` : 统计文件中最长一行的长度

## 系统操作

### systemctl

```bash
systemctl [OPTIONS...] COMMAND ...
# 管理和控制 systemd 系统和服务管理器的命令行工具
```

- command

  - status <server-name> : 查看特定服务的当前状态
  - start <server-name> : 启动服务
  - restart <server-name> : 重启服务
  - reload <server-name> : 重新加载服务配置
  - stop <server-name> : 停止服务
  - enable <server-name> : 系统启动时自动启动
  - disable <server-name> : 禁用服务
  - daemon-reload: 重新加载 systemd 系统和服务管理器的配置文件

- options : 

  - --type : 指定要列出的单位类型（unit type）。常见的单位类型包括 service, socket, target, mount

  - --state : 过滤和显示具有特定状态的服务单元

    ```bash
    active: 单元正在运行或已完成
    inactive: 单元未运行且未激活
    reloading: 单元正在重新加载其配置
    failed: 单元失败了，可能是因为启动过程中出现错误
    activating: 单元正在启动中
    deactivating: 单元正在停止中
    ```

  - --force : 强制执行

### lsblk

```bash
lsblk
# list block devices，用于列出系统中所有可用的块设备信
```

## awk

```bash
awk [选项] '模式 {动作}' 文件名
# 文本处理工具
```

- 选项: 可选参数，如 -F 指定字段分隔符
- 模式: 指定何时执行动作的条件
- 动作: 当模式匹配时要执行的操作
  - print: 用于打印整行或指定字段
  - $[n]: 字段变量
- 文件名: 需要处理的文件名

## install

```bash
install [选项] 源文件 目标文件
install [选项] 源文件... 目标目录
install -d 目标目录...
# 安装或复制文件及目录，并设置适当的权限和所有权
```

- 选项

  - -d 或 --directory: 创建目录。如果目录已存在，则不会报错

  - -m 或 --mode: 设置文件或目录的权限模式。模式可以是八进制数字或符号表示法

## curl

```bash
curl [选项] [URL]
#用于从服务器传输数据或向服务器传输数据的命令行工具
```

- 选项
  - -f: 禁止返回 HTTP 错误页面，如果请求失败，`curl` 会返回一个错误而不是显示错误页面
  - -s: 静默模式，不显示进度条或其他输出
  - -S: 即使使用 `-s` 选项，仍然在发生错误时显示错误信息
  - -L: 跟随重定向。如果服务器返回一个重定向响应，`curl` 会自动跟随重定向
  - -o: 指定下载的文件保存的路径和文件名

## chomd

```bash
chmod [选项] 权限模式 文件或目录
# 更改文件或目录权限
```

- 选项
  - -R 或 --recursive: 递归地更改目录及其内容的权限
  - -v 或 --verbose: 显示详细的输出信息
- 权限模式
  - 符号表示法
  - 八进制表示法

## tee

```bash
tee [选项] 文件...
# 将命令的输出同时显示在终端上并保存到文件中
```

- 选项
  - -a: 追加到文件末尾而不是覆盖文件

## sudo

```bash
# sudo - execute a command as another user
sudo [optins]

Options:
  -i, --login                   run login shell as the target user; a command may also be specified
```

## source

​	在当前shell环境中执行一个指定的脚本文件（和 . 同效）。

## journalctl

​	journalctl 是 systemd 系统和服务管理器的一部分，用于查询和显示系统日志。systemd 日志系统（也称为 journald）收集来自内核、系统服务和用户进程的日志信息，并将其存储在一个二进制日志文件中。journalctl 命令提供了多种选项来过滤和格式化这些日志信息，使其更易于阅读和分析。

```bash
journalctl [OPTIONS...] [MATCHES...]

Query the journal.

Flags:
     --system              Show the system journal
     --user                Show the user journal for the current user
  -M --machine=CONTAINER   Operate on local container
  -S --since=DATE          Show entries not older than the specified date
  -U --until=DATE          Show entries not newer than the specified date
  -c --cursor=CURSOR       Show entries starting at the specified cursor
     --after-cursor=CURSOR Show entries after the specified cursor
     --show-cursor         Print the cursor after all the entries
  -b --boot[=ID]           Show current boot or the specified boot
     --list-boots          Show terse information about recorded boots
  -k --dmesg               Show kernel message log from the current boot
  -u --unit=UNIT           Show logs from the specified unit
  -t --identifier=STRING   Show entries with the specified syslog identifier
  -p --priority=RANGE      Show entries with the specified priority
  -e --pager-end           Immediately jump to the end in the pager
  -f --follow              Follow the journal
  -n --lines[=INTEGER]     Number of journal entries to show
     --no-tail             Show all lines, even in follow mode
  -r --reverse             Show the newest entries first
  -o --output=STRING       Change journal output mode (short, short-iso,
                                   short-precise, short-monotonic, verbose,
                                   export, json, json-pretty, json-sse, cat)
     --utc                 Express time in Coordinated Universal Time (UTC)
  -x --catalog             Add message explanations where available
     --no-full             Ellipsize fields
  -a --all                 Show all fields, including long and unprintable
  -q --quiet               Do not show privilege warning
     --no-pager            Do not pipe output into a pager
  -m --merge               Show entries from all available journals
  -D --directory=PATH      Show journal files from directory
     --file=PATH           Show journal file
     --root=ROOT           Operate on catalog files underneath the root ROOT
     --interval=TIME       Time interval for changing the FSS sealing key
     --verify-key=KEY      Specify FSS verification key
     --force               Override of the FSS key pair with --setup-keys

Commands:
  -h --help                Show this help text
     --version             Show package version
  -F --field=FIELD         List all values that a specified field takes
     --new-id128           Generate a new 128-bit ID
     --disk-usage          Show total disk usage of all journal files
     --vacuum-size=BYTES   Reduce disk usage below specified size
     --vacuum-time=TIME    Remove journal files older than specified date
     --flush               Flush all journal data from /run into /var
     --header              Show journal header information

```



# 权限模式

##  四位八进制表示法

​	八进制表示法使用三位八进制数字来表示文件或目录的权限。每一位数字代表一组用户的权限，分别是文件所有者（user）、文件所属组（group）和其他用户（others）。

- 位数作用

  - 第一位: 特殊权限位

  - 第二位: 文件所有者的权限

  - 第三位: 文件所属组的权限

  - 第四位: 其他用户的权限

- 权限位的含义

  - 4: 读权限（read, r）

  - *: 写权限（write, w）

  - 1: 执行权限（execute, x）

- 权限组合

  - 7: 读、写、执行（rwx）

  - 6: 读、写（rw-）

  - 5: 读、执行（r-x）

  - 4: 读（r--）

  - 3: 写、执行（-wx）

  - 2: 写（-w-）

  - 1: 执行（--x）

  - 0: 无权限（---）

##  符号表示法

符号表示法使用字符来表示用户类别和权限。

- 用户类别: 
  - u: 文件所有者（user）
  - g: 文件所属组（group）
  - o: 其他用户（others）
  - a: 所有用户（all，等同于 `ugo`）
- 操作符: 
  - +: 添加权限
  - -: 移除权限
  - =: 设置权限
- 权限: 
  - r: 读权限（read）
  - w: 写权限（write）
  - x: 执行权限（execute）

# shell 特殊符号

| 符号 | 作用                                         |
| ---- | -------------------------------------------- |
| \|   | 将前一个命令的输出作为下一个命令的输入       |
| >    | 将命令的输出重定向到文件，覆盖文件内容       |
| >>   | 将命令的输出重定向到文件，追加到文件末尾     |
| &    | 将命令放入后台执行                           |
| &&   | 只有前一个命令执行成功时，才会执行下一个命令 |
| *    | 匹配任意数量的字符（包括零个字符）           |
| ？   | 匹配一个字符                                 |

