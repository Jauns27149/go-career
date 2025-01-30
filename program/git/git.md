Git提交标题格式

- feat: 新增功能
- fix: 修复 bug
- docs: 文档更新
- style: 格式化代码，不影响功能
- refactor: 重构代码，不添加新功能也不修复 bug
- test: 添加或修改测试用例
- chore: 其他维护工作，如更新依赖包

# 修改最近一次commit的message

```bash
git commit --amend -m "fix# 快照统计功能调试"
```

# fatal: Cannot prompt because user interactivity has been disabled

```bash
git config --global credential.helper manager-core
git credential-manager-core store # enter两次退出
```

```bash
protocol=https
host=work.ctyun.cn
username=wueq@mail.asiainfo.com
password=4.jVHjy&Leh1
```

