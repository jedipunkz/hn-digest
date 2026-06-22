---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48632192"
title: "Enabling Claude Code to run code when its sandbox fails (windows only)"
article_title: ""
author: "algoth1"
captured_at: "2026-06-22T16:33:24Z"
capture_tool: "hn-digest"
hn_id: 48632192
score: 1
comments: 0
posted_at: "2026-06-22T16:14:55Z"
tags:
  - hacker-news
  - translated
---

# Enabling Claude Code to run code when its sandbox fails (windows only)

- HN: [48632192](https://news.ycombinator.com/item?id=48632192)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T16:14:55Z

## Translation

タイトル: サンドボックスが失敗したときにクロード コードがコードを実行できるようにする (Windows のみ)
HN テキスト: これは愚かな問題に対する愚かな解決策であることをあらかじめご了承ください。クロード コード/共同作業は、サンドボックス環境を使用できないことがよくあります。つまり、ファイルの読み取りと書き込みはできますが、実行できないため、ほとんど役に立たなくなります。愚かな解決策: 現在のフォルダーを x 秒ごとにチェックして Sandbox-workaround.txt があるかどうかを確認する .vbs スクリプトを作成するよう codex に依頼します。その txt の最初の行にバット ファイルへのパスが含まれ、2 行目に run=true フラグが含まれている場合は、フラグを false に切り替えながら、対応するバット ファイルを 1 回実行します。サンドボックスが失敗するたびに .txt ファイルを「準備」するようにクロードに指示します。私のユースケースでは、クロードは Python ファイルをビルドし、サンドボックスは失敗します。彼は .bat ファイルを作成し、pyrhon にそのファイルを実行し、サンドボックス txt ファイルを編集し、フラグを true に設定して待機するように指示します。醜いですが、機能し、すべてを再起動するよりも高速です。おそらく多くのセキュリティ問題も発生する可能性があるため、自己責任で使用してください。

## Original Extract

Be forewarned that this is a dumb solution to a dumb problem. Claude code/cowork will often be unable to use the sandbox environment: this means it can read and write files but can’t run them - making it mostly useless. The dumb solution: ask codex to create a .vbs script that checks the current folder every x seconds for a sandbox-workaround.txt. If that txt contains a path to a bat file as the first line, and a run=true flag as the second line, proceed to run the corresponding bat file once, while switching the flag to false. Tell claude to ‘arm’ the .txt file whenever the sandbox fails. For my use case, claude will build a python file, the sandbox will fail, he will write a .bat file telling pyrhon to run that file, edit the sandbox txt file, set the flag to true, and wait. It’s ugly but it works, and it’s faster that rebooting everything. It probably also opens a host of security issues, so use at your own risk

