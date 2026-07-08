---
source: "https://parchmint.app"
hn_url: "https://news.ycombinator.com/item?id=48832387"
title: "Show HN: Parchmint – a Markdown editor that shows what the AI actually reads"
article_title: "Parchmint — the Markdown editor for documents AI reads"
author: "mazheru"
captured_at: "2026-07-08T15:11:54Z"
capture_tool: "hn-digest"
hn_id: 48832387
score: 1
comments: 0
posted_at: "2026-07-08T14:24:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Parchmint – a Markdown editor that shows what the AI actually reads

- HN: [48832387](https://news.ycombinator.com/item?id=48832387)
- Source: [parchmint.app](https://parchmint.app)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T14:24:24Z

## Translation

タイトル: Show HN: Parchmint – AI が実際に何を読み取るかを示す Markdown エディター
記事のタイトル: Parchmint — AI が読み取るドキュメント用の Markdown エディター
説明: Parchmint は、LLM が実際に見る正直なプレビュー (すべてのトークン、すべての警告) を表示するため、プロンプト、スキル ファイル、ドキュメントは書かれたとおりに正確に表示されます。無料、オープンソース、ローカルファースト。

記事本文:
パーチミント
アイデア
特長
ダウンロード
出典 ↗
v0.1.0 · MIT
マシンリーダーのマークダウン
マークダウン、
モデルがそれを読むと。
Parchmint は、LLM が実際に見る正直なプレビューを示します —
すべてのトークン、すべての警告 - プロンプト、スキル ファイル、ドキュメントが表示される
まさに書かれているとおりです。推論時に予期せぬ書式設定が行われることはありません。
無料 & オープンソース · ローカルファースト · アカウントなし、テレメトリなし
---
名前：正直プレビュー
説明 : あなたが見ているものは何ですか
モデルが取得します
---
## モデルが読み取るもの
- [x] すべてのトークンがカウントされます
- [ ] 何も隠されていません
**1 回** を書き込み、クリーンな状態で出荷します。
モデルは読みます
モデルが読み取るもの
すべてのトークンがカウントされる
あなたが見ているのは
モデルが得られるもの。
ほとんどのエディターは、目に見えるように Markdown をレンダリングします - 美しい見出し、非表示
空白、沈黙の仮定。しかし、読者は言語モデルです。パーチミント
トークナイザーが文書を解析する方法で文書をレンダリングし、トークンの予算を明らかにします。
そして、プロンプトを静かに破る構造にフラグを立てます。あなたはそれが何であるかを推測するのをやめます
モデルが見ます。ただ見てください。
## コンテキスト
あなたは **正確な**編集者です。
- [ ] それぞれの主張を確認してください
→
トークン化されたものとして
## コンテキスト あなたは正確な編集者です。それぞれの主張を検証する
02 — 中身は何ですか
書類作成のワークショップ
機械の聴衆と一緒に。
トークナイザーが解析したものと一致するレンダリングです。飾り立てられた幻想ではありません。疑似タグと生の HTML がわかりやすく表示されます。
入力時にライブトークンがカウントされ、文書が作成対象の予算を超過する前にトークンレベルの警告が表示されます。
文書タイプの検出と構造的リンティングにより、プロンプトを静かに脱線させる壊れたアウトラインや不正な前付部分を検出します。
すべての見出しのライブアウトライン。カーソルを任意の場所にジャンプするか、プレビューを右クリックして、クリックした正確な行を編集します。
図、数学、および構文が強調表示されたフェンスはインラインでレンダリングされます。 3 つのビュー モード — サワー

ce、分割、プレビュー — キーストロークごとに異なります。
Tauri と Rust をベースに構築されています。ファイルはディスク上に残り、外部からの変更は検出されますが、何も通知されません。これまで。
無料、オープンソース、MIT ライセンス。プラットフォームを選択してください。すべてのビルドは最新リリースからのものです。
またはソースからビルドします: npm install && npm run tauri build
羊皮紙に作られ、機械で読み取られます。
© 2026 アブ・マゼール・ウディン。 MITライセンスに基づいてリリースされています。

## Original Extract

Parchmint shows the honest preview an LLM actually sees — every token, every warning — so your prompts, skill files, and docs land exactly as written. Free, open source, local-first.

Parchmint
The idea
Features
Download
Source ↗
v0.1.0 · MIT
Markdown for machine readers
Markdown,
as the model reads it.
Parchmint shows the honest preview an LLM actually sees —
every token, every warning — so your prompts, skill files, and docs land
exactly as written. No formatting surprises at inference time.
Free & open source · Local-first · No account, no telemetry
---
name : honest-preview
description : what you see is what
the model gets
---
## What the model reads
- [x] every token counted
- [ ] nothing hidden
Write **once** , ship clean.
the model reads
What the model reads
every token counted
What you see is
what the model gets.
Most editors render Markdown for your eyes — pretty headings, hidden
whitespace, silent assumptions. But your reader is a language model. Parchmint
renders the document the way a tokenizer parses it, surfaces the token budget,
and flags the structure that quietly breaks prompts. You stop guessing what the
model sees. You just look.
## Context
You are a **precise** editor.
- [ ] verify each claim
→
As tokenized
## Context You are a precise editor . verify each claim
02 — What's inside
A workshop for documents
with a machine audience.
A render that matches what the tokenizer parses — not a prettied-up illusion. Pseudo-tags and raw HTML are shown plainly.
Live token count as you type, with token-level warnings before a document blows past the budget you're writing for.
Document-type detection plus structural linting catches the broken outlines and malformed frontmatter that quietly derail prompts.
A live outline of every heading. Jump the cursor anywhere — or right-click the preview and edit the exact line you clicked.
Diagrams, math, and syntax-highlighted fences render inline. Three view modes — source, split, preview — a keystroke apart.
Built on Tauri and Rust. Your files stay on disk, external changes are detected, and nothing phones home. Ever.
Free, open source, MIT-licensed. Pick your platform — every build comes from the latest release.
Or build from source: npm install && npm run tauri build
Crafted on parchment · read by machines.
© 2026 Abu Mazher Uddin. Released under the MIT License.
