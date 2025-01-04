# 命令

- netstat -ano | findstr ":0.0.0.0:<port>" : 这个端口是否使用情况

# PowerShell 

- 配置文件路径 

  ```bash
  $PROFILE
  # C:\Users\Administrator\Documents\WindowsPowerShell\Microsoft.PowerShell_profile.ps1
  ```

  ```bash
  Set-Alias etcd "C:\Janus\soft\etcd-v3.5.17-windows-amd64\etcd.exe" # 命令别名
  
  # 背景颜色
  $Host.UI.RawUI.BackgroundColor = 'Black'
  $Host.UI.RawUI.ForegroundColor = 'White'
  Clear-Host  # 刷新屏幕以应用颜色
  ```

- Start-Process powershell ：多开窗口