# 通用规则和环境配置

## Python 环境配置

### Conda 环境
- 项目中已配置 conda 环境，名称为 `gva`
- 执行任何 Python 脚本前，必须先激活此环境
- 使用命令: `conda activate gva`

### Python 脚本示例
sh(eval "$(/opt/homebrew/bin/conda shell.zsh hook)" && conda activate gva && cd scripts && python read_excel.py ../temp/测试聚合层文档_分析_最终_1.xlsx 10)

### 示例
```bash
conda activate gva && python scripts/read_excel.py
```

## 其他通用规则
（待补充）