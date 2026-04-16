"""
HTML 富文本 → 纯文本转换
读取脚本同级目录下的 a.txt（内容为 HTML），输出 a-plain.txt。

规则：
1. 去除所有 HTML 标签
2. 块级元素（h1-h6, p, li, div 等）之间保留换行
3. 列表项前加 "- " 提示
4. HTML 实体解码（&nbsp; → 空格等）
5. 合并多余空白行
"""
import os
import re
import html
from bs4 import BeautifulSoup, NavigableString

# === 定位同级目录下的输入/输出文件 ===
SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
INPUT_PATH = os.path.join(SCRIPT_DIR, 'a.txt')
OUTPUT_PATH = os.path.join(SCRIPT_DIR, 'a-plain.txt')

# === 读取源文件 ===
with open(INPUT_PATH, 'r', encoding='utf-8') as f:
    html_content = f.read()

soup = BeautifulSoup(html_content, 'html.parser')

# 块级标签：渲染前后加换行
BLOCK_TAGS = {'h1', 'h2', 'h3', 'h4', 'h5', 'h6', 'p', 'div', 'ul', 'ol', 'blockquote', 'pre', 'br'}
LIST_ITEM_TAGS = {'li'}


def extract_text(node, lines):
    """递归提取文本，保留块级结构"""
    if isinstance(node, NavigableString):
        text = str(node)
        if text.strip():
            if lines and not lines[-1].endswith('\n'):
                lines[-1] += text
            else:
                lines.append(text)
        return

    tag = node.name

    if tag == 'br':
        lines.append('\n')
        return

    if tag in LIST_ITEM_TAGS:
        lines.append('\n- ')
        for child in node.children:
            extract_text(child, lines)
        return

    if tag in BLOCK_TAGS:
        if lines and lines[-1] and not lines[-1].endswith('\n'):
            lines.append('\n')
        for child in node.children:
            extract_text(child, lines)
        lines.append('\n')
        return

    # 行内元素（a, span, strong, em 等）直接提取
    for child in node.children:
        extract_text(child, lines)


lines = []
extract_text(soup, lines)

text = ''.join(lines)

# HTML 实体解码
text = html.unescape(text)

# 清理空白
text = re.sub(r'[ \t]+', ' ', text)       # 多空格合一
text = re.sub(r' *\n *', '\n', text)      # 行首行尾空格去除
text = re.sub(r'\n{3,}', '\n\n', text)    # 最多连续两个换行
text = text.strip()

# 输出
with open(OUTPUT_PATH, 'w', encoding='utf-8') as f:
    f.write(text)

# 预览
print(f'输入: {INPUT_PATH}')
print(f'输出: {OUTPUT_PATH}')
print('--- 预览前 500 字 ---')
print(text[:500])
print('---')
print(f'总字数: {len(text)}')