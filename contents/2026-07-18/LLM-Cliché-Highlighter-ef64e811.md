---
source: "https://tools.simonwillison.net/llm-cliche-highlighter"
hn_url: "https://news.ycombinator.com/item?id=48955805"
title: "LLM Cliché Highlighter"
article_title: "LLM cliché highlighter"
author: "m_c"
captured_at: "2026-07-18T06:53:17Z"
capture_tool: "hn-digest"
hn_id: 48955805
score: 1
comments: 0
posted_at: "2026-07-18T06:48:38Z"
tags:
  - hacker-news
  - translated
---

# LLM Cliché Highlighter

- HN: [48955805](https://news.ycombinator.com/item?id=48955805)
- Source: [tools.simonwillison.net](https://tools.simonwillison.net/llm-cliche-highlighter)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T06:48:38Z

## Translation

タイトル: LLM クリシェ ハイライター
記事のタイトル: LLM 決まり文句のハイライター

記事本文:
以下にテキストを貼り付けるか、URL から読み込むと、既知の LLM 決まり文句に一致する文が強調表示されます。 「X なし、Y なし」のようなチェーン パターンには、項目をカウントするバッジが付けられます。ハイライトにカーソルを合わせると、どの決まり文句がヒットしたかを確認できます。
上にテキストを貼り付けます。入力すると分析が実行されます。
これらを Node でヘッドレスで実行します。以下のワンライナーは、標準入力でこのファイルを読み取り、スクリプト内の「impl」と「tests」の開始/終了マーカー コメントの間のコードを抽出し、それを eval() して、合格/失敗の集計を出力します。何かが失敗した場合の終了コードは 1 です。

## Original Extract

Paste text below — or load it from a URL — and it highlights sentences that match known LLM clichés. Chain patterns like “no X, no Y” get a badge counting their items; hover a highlight to see which cliché it hit.
Paste some text above — analysis runs as you type.
Run these headlessly with Node: the one-liner below reads this file on stdin, extracts the code between the “impl” and “tests” start/end marker comments in the script, eval()s it, and prints a pass/fail tally. Exit code is 1 if anything fails.
