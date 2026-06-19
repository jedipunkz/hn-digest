---
source: "https://simplybychris.github.io/antigravity-plugin-cc/"
hn_url: "https://news.ycombinator.com/item?id=48601044"
title: "Show HN: Open-source Antigravity plugin for Claude Code"
article_title: "agy Antigravity CLI plugin for Claude Code"
author: "simplybychris"
captured_at: "2026-06-19T18:43:08Z"
capture_tool: "hn-digest"
hn_id: 48601044
score: 2
comments: 0
posted_at: "2026-06-19T17:42:39Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open-source Antigravity plugin for Claude Code

- HN: [48601044](https://news.ycombinator.com/item?id=48601044)
- Source: [simplybychris.github.io](https://simplybychris.github.io/antigravity-plugin-cc/)
- Score: 2
- Comments: 0
- Posted: 2026-06-19T17:42:39Z

## Translation

タイトル: HN を表示: クロード コード用のオープンソース反重力プラグイン
記事のタイトル: クロード コード用の agy Antigravity CLI プラグイン
説明: Claude Code から Google にタスクを委任できるオープンソース プラグイン
HN テキスト: クロード コード、コーデックス、またはその他のエージェントと連携する人々向けのクイック シェア。エディターを離れることなく CC 内から簡単にアクセスできるようにする Claude Code プラグインを構築しました。
ネイティブ イメージの生成、詳細な調査、コード レビューなどの追加のスラッシュ コマンド。面白い部分: agy には CLI -p にモデル フラグがありません。プラグインはこれを追加します。リポジトリ: https://github.com/simplybychris/antigravity-plugin-cc レビュー、共有、貢献をお待ちしております。

記事本文:
アジ
クロードコード用プラグイン
インストール
特長
GitHub
インストール
v0.4.1 · オープンソース · MIT
Claude Code 内から Google Antigravity にタスクを委任します。
agy CLI をラップする小さなオープンソース プラグイン。
6 つのスラッシュ コマンド、転送サブエージェント、コード レビュー、詳細な調査、
エディターを離れることなく、ネイティブ イメージを生成できます。
クロード コード内でこれらを順番に実行します。各カードにはワンクリックでコピー ボタンが付いています。
Claude Code にプラグインの場所を指示します。
/プラグインマーケットプレイス追加 Simplybychris/antigravity-plugin-cc
プラグインをインストールする
v0.4.1 をユーザー スコープのプラグイン キャッシュにプルします。
/プラグインのインストール agy@antigravity-cc
リロードして確認する
スラッシュ コマンドをアクティブにし、インストール/認証チェックを実行します。
/reload-プラグイン
/agy:セットアップ
特長
各タイルは 1 つのエントリ ポイントです。クロードに自然に話しかけるか、直接スラッシュコマンドを呼び出します。
agy が箱にあり、サインインしていることを確認します。初回実行時にインストールするように求められます。
単一のプロンプトを agy に送信し、その回答をそのままクロード コードにストリーミングして返します。
個別のコーディング、リファクタリング、またはデバッグ タスクを agy:runner サブエージェントにルーティングします。 --background と --model をサポートします。
あなたのトピックを構造化された研究概要 (背景、調査結果、注意点、情報源) にまとめて、agy に調査させます。
agy の組み込み Imagen ツールによるネイティブ イメージの生成。オプションの --name および --output を使用して、プロジェクトの隣にファイルをドロップします。
現在の git diff を agy にパイプして、オプションのフォーカス テキストを使用して操作可能なレビューを実行します。
エージェントは二人。端末は1つ。タスクごとに適切なツールを選択してください。
クロード・コードに留まってください。
クロードを主なドライバーとして使用します。 Gemini の方が適している場合は、コンテキストを切り替えることなく、ターゲットを絞ったタスクを agy に引き渡します。
バックグラウンドでの長時間のジョブ。
Claude Code のサブエージェントはバックグラウンドで実行され、終了時に通知されます。 --バックグラウンドリサーチ、リファクタリング、CI に自然に適合

調査。
実行時の荷物はありません。
ローカルの agy バイナリの純粋な Bash ラッパー。ノードもブローカーも隠れたサービスもありません。
オープンソース、MIT ライセンス。
ラッパーを監査し、フォークし、変更します。問題やプルリクエストは歓迎です。

## Original Extract

Open-source plugin that lets Claude Code delegate tasks to Google

Quick share for people who work with Claude Code, Codex or other agents. I built a Claude Code plugin that gives you agy access from inside CC without leaving the editor.
Additional slash commands, including native image generation, deep research and code review. The fun part: agy doesn't have a model flag on the CLI -p. The plugin adds this. Repo: https://github.com/simplybychris/antigravity-plugin-cc Will be glad for any review, share or contribution!

agy
plugin for Claude Code
Install
Features
GitHub
Install
v0.4.1 · open source · MIT
Delegate tasks to Google Antigravity from inside Claude Code.
A small open-source plugin that wraps the agy CLI.
Six slash commands, a forwarding subagent, code review, deep research,
and native image generation without leaving your editor.
Run these inside Claude Code, in order. Each card has a one-click copy button.
Tells Claude Code where to find the plugin.
/plugin marketplace add simplybychris/antigravity-plugin-cc
Install the plugin
Pulls v0.4.1 into the user-scope plugin cache.
/plugin install agy@antigravity-cc
Reload and verify
Activates the slash commands and runs the install/auth check.
/reload-plugins
/agy:setup
Features
Each tile is one entry point. Talk to Claude naturally or call the slash command directly.
Verify agy is on the box and signed in. Offers to install it on first run.
Send a single prompt to agy and stream the answer back into Claude Code, verbatim.
Route a discrete coding, refactor, or debugging task to the agy:runner subagent. Supports --background and --model .
Wraps your topic in a structured research brief (background, findings, caveats, sources) and lets agy investigate.
Native image generation via agy 's built-in Imagen tool. Optional --name and --output to drop the file next to your project.
Pipe the current git diff into agy for a steerable review with optional focus text.
Two agents. One terminal. Pick the right tool per task.
Stay in Claude Code.
Use Claude as the primary driver. Hand off targeted tasks to agy when Gemini is the better fit, without switching context.
Long jobs in the background.
Subagents in Claude Code run in the background and notify you when they finish. A natural fit for --background research, refactors, or CI investigations.
No runtime baggage.
Pure Bash wrapper around the local agy binary. No Node, no broker, no hidden services.
Open source, MIT-licensed.
Audit the wrapper, fork it, modify it. Issues and pull requests welcome.
