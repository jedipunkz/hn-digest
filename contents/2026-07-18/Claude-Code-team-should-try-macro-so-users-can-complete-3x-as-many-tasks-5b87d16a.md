---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48957109"
title: "Claude Code team should try macro so users can complete 3x as many tasks"
article_title: ""
author: "yohji1984"
captured_at: "2026-07-18T11:44:57Z"
capture_tool: "hn-digest"
hn_id: 48957109
score: 1
comments: 0
posted_at: "2026-07-18T11:37:31Z"
tags:
  - hacker-news
  - translated
---

# Claude Code team should try macro so users can complete 3x as many tasks

- HN: [48957109](https://news.ycombinator.com/item?id=48957109)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T11:37:31Z

## Translation

タイトル: クロード コード チームは、ユーザーが 3 倍のタスクを完了できるようにマクロを試してください
HN text: つまり、アイデアは非常にシンプルです。 CC がファイルを変更してテストを実行する必要がある場合、CC は次のことを行う必要があります。
# ターン 1 — パッチを適用します (パッケージ ファイルを変更します)
# ターン 2 — パッチを適用します (バグを修正します)
# ターン 3 — パッチを適用します (テスト スクリプトを編集します)
# ターン 4 — 構築
# ターン 5 — テストと lint の実行
# ターン 6 — *Playwright の場合、CC はメディアを読み取るために別のラウンドが必要です。コンセプトは、マクロ コマンドを使用して、プロセス全体を 1 つのターンで RAG に入れることができるということです。たとえば次のようになります: step1: env を検査する
step2: パッチ適用(パッケージファイル変更)
step2: パッチを適用する(バグを修正する)
step2: パッチを適用する (テストスクリプトを編集する)
ステップ3: ビルドする
ステップ4: テストと lint を実行する
step5: メディアを読み取る Antropic の利用規約のため、この実装をテストできませんでした。ただし、他のコーディング エージェントとプロバイダーでテストしたところ、llm ターンアラウンドが 40% ～ 80% 削減され、ほぼ同じ量のトークンが節約できました。 （投稿がボットによって削除されるのを避けるために、** を使用してプロバイダー名を非表示にする必要があります）。このアプローチが役に立ち、私のテスト方法が合理的であると思われる場合は、賛成票を投じてください。ここに実装とテストのベンチマークがあり、MD にあります。 DeepSWE とリポジトリ全体の書き換えケースでテストしました: https://github.com/Tura-AI/tura

## Original Extract

So the idea is really simple. If CC need to change a file and run some test, CC needs to:
# Turn 1 — apply patch (change package file)
# Turn 2 — apply patch (fix the bug)
# Turn 3 — apply patch (edit the testing script)
# Turn 4 — build
# Turn 5 — run tests and lint
# Turn 6 — *and if is playwright CC needs another round to read media The concept is that we can use macro command to put the entire process into a RAG in one single turn. For example like this: step1: inspect env
step2: apply patch (change package file)
step2: apply patch (fix the bug)
step2: apply patch (edit the testing script)
step3: build
step4: run tests and lint
step5: read medias I could not test this implementation due to Antropic terms of use. But I tested on other coding agent and providers this can reduce llm turns around by 40% - 80%, and saving nearly same amount of tokens. （I need to use ** to hide the provider's name in order to avoid the post being deleated by the bot). If you think this approch could help and my testing method is reasonable, you can give me an upvote! Thx Here is the implementation and testing benchmark can be found in the md. I tested with DeepSWE and entire repo rewrite cases: https://github.com/Tura-AI/tura

