无法连接Linux虚拟机的Mysql排查

1. 确保虚拟机和云服务提供商的防火墙允许通过3306端口的流量

2. 确保登录用户有远程登录权限以及密码正确

   登录MySQL并检查用户信息。确保您使用的账户`Host`字段是`%`或目标主机IP(用于连接的主机)

   ```bash
   mysql> SELECT User, Host FROM mysql.user;
   +------------------+-----------+
   | User             | Host      |
   +------------------+-----------+
   | janus            | %         |
   | debian-sys-maint | localhost |
   | mysql.infoschema | localhost |
   | mysql.session    | localhost |
   | mysql.sys        | localhost |
   | root             | localhost |
   +------------------+-----------+
   # Host字段为允许连接的主机，%代表任意主机
   ```
   
3. Mysql是否支持远程连接

   检查配置文件`mysqld.cnf`，确保`bind-address`绑定的是`0.0.0.0`或目标主机(用于连接的主机)

   ```bash
   # Mysql的配置文件一般在 /etc/mysql/mysql.conf.d/
   janus@janus:/etc/mysql/mysql.conf.d$ cat mysqld.cnf | grep bind-address
   bind-address		= 0.0.0.0
   mysqlx-bind-address	= 127.0.0.1
   
   # 不同版本的Mysql配置文件位置可能略不同，但基本都在 /etc/mysql/目标，也可以用grpe进行递归搜索排查
   janus@janus:/etc/mysql$ sudo grep -r bind-addres
   ./mysql.conf.d/mysqld.cnf:bind-address		= 0.0.0.0
   ./mysql.conf.d/mysqld.cnf:mysqlx-bind-address	= 127.0.0.1
   # 0.0.0.0 代表任意主机
   ```