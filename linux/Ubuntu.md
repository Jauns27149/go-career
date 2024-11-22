# 结构

- /
  - var: 存放可变数据文件(动态)
    - lib: 存放与特定应用程序相关的状态信息和数据文件
  - etc: 存放系统的配置文件(静态)
    - os-release: 存储操作系统的识别信息(key-value)
    - shadow: 存储用户密码和其他与用户帐户相关的安全信息的重要文件
  - opt: 存放可选的第三方软件包
  - tmp: 存储临时文件和数据(temporary)

# 命令

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