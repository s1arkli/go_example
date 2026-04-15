import subprocess
import os
from datetime import date

base = "/Users/cc/note/日志"
now = date.today().strftime("%Y/%m月/%m-%d") # 2026/4月/04-14
output = f'{base}/{now}.md'

os.makedirs(os.path.dirname(output), exist_ok=True)
for repo in ['carutoo-server', 'carl_server']:
    print(f'{repo}总结中。。。')
    result = subprocess.run(
        ['codex', 'exec', '帮我总结今日git提交记录，生成日报，我在中国，按照中国的时区进行git命令。要求：只需要总结的相关内容，如果未产生git提交记录那么只返回今日无提交记录即可'],
        cwd=f'/Users/cc/go_project/{repo}',
        capture_output=True, text=True
    )
    with open(output, 'a') as f:
        f.write(f'\n## {repo}\n\n')
        f.write(result.stdout)
print('总结完毕')

#  echo "alias daily='python3 /Users/cc/go_project/go_example/py/summary.py' " >> ~/.zshrc
# 把命令加入配置文件，使用daily运行