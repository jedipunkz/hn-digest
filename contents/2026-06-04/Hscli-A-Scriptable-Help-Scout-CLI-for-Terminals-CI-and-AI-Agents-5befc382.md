---
source: "https://wavyx.github.io/hscli/"
hn_url: "https://news.ycombinator.com/item?id=48394847"
title: "Hscli – A Scriptable Help Scout CLI for Terminals, CI, and AI Agents"
article_title: "hscli — Help Scout from your terminal"
author: "wavyx"
captured_at: "2026-06-04T07:45:28Z"
capture_tool: "hn-digest"
hn_id: 48394847
score: 1
comments: 0
posted_at: "2026-06-04T06:34:11Z"
tags:
  - hacker-news
  - translated
---

# Hscli – A Scriptable Help Scout CLI for Terminals, CI, and AI Agents

- HN: [48394847](https://news.ycombinator.com/item?id=48394847)
- Source: [wavyx.github.io](https://wavyx.github.io/hscli/)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T06:34:11Z

## Translation

タイトル: Hscli – ターミナル、CI、AI エージェント用のスクリプト可能な Help Scout CLI
記事のタイトル: hscli — 端末から Scout をヘルプする
説明: Help Scout 用の高速でスクリプト可能な CLI — ターミナル、CI パイプライン、AI エージェント用に構築されています。

記事本文:
hs cli ガイド コマンド エージェント 変更ログ ドキュメントの検索 ⌘ K オープン ソース · Help Scout CLI Help Scout、
あなたの端末から。
Help Scout 用の高速でスクリプト可能な CLI — ターミナル、CI パイプライン、Claude Code や Codex などの AI エージェント向けに構築されています。すべてのコマンドは JSON を読み上げ、クリーンな終了コードを返します。
GitHub でスターを付ける
$ npm i -g @wavyx/hscli zsh — hscli $ hscli conv list --status active --output table ID 対象顧客ステータス 4831 注文の返金 #2207 amir@acme.io active 4830 パスワードをリセットできません lou@hey.com active 4827 請求書の質問 kai@studio.co active $ # タグ付けされたものをすべて閉じます"解決済み" $ hscli convBulk-status --status active --tag selected \ --set Closed --yes ✓ 12 の会話が終了しました · exit 0 なぜ hscli なのか
クリックではなく自動化のために構築されています。
hscli をスクリプトまたはエージェントに安全に渡すための 4 つの設計上の決定事項。
JSON 出力、確定的な終了コード、生の hscli API ハッチにより、エージェントとパイプラインが Help Scout を直接操作できます。
すべてのコマンドは、 --output table|json|yaml|csv に加えて、インライン フィルタリングとプロジェクション用の --jq および --field をサポートします。
OAuth トークンは OS キーチェーン内にのみ存在し、平文でディスクに書き込まれることはありません。 hscli API は Help Scout ホストにロックされています。
hscli バックアップは、増分更新、再開、削除検出、添付ファイルを使用して、アカウント全体を JSON にダンプします。
実際の作業は、一度に 1 行ずつです。
hscli は、すでに使用しているツール ( jq 、 xargs 、 cron 、および CI ランナー) で構成されます。
トリアージ エージェントにキューをクリアさせます
アクティブな会話をモデルにパイプし、その決定に基づいて動作します。統合コードは必要ありません。
CI での Nightly アカウントのスナップショットのバックアップ
デフォルトで増分、再開可能、削除対応。 cron ジョブまたは GitHub アクションにドロップします。
レポート メトリクスをスプレッドシートに取り込む
レポートを CSV としてストリーミングし、直接パイプ処理します

BI ツールまたは Google スプレッドシートに入力します。
どこからでも Help Scout を起動します。
一度インストールしてください。すべてをスクリプト化します。残りをエージェントに渡します。
コマンドリファレンス hs cli MIT ライセンスを取得していますが、Help Scout ガイドとは関係ありません コマンド GitHub npm Esc Try conv list 、backup 、または authentication

## Original Extract

A fast, scriptable CLI for Help Scout — built for terminals, CI pipelines, and AI agents.

hs cli Guide Commands Agents Changelog Search docs ⌘ K Open source · Help Scout CLI Help Scout,
from your terminal .
A fast, scriptable CLI for Help Scout — built for terminals, CI pipelines, and AI agents like Claude Code and Codex. Every command speaks JSON and returns clean exit codes.
Star on GitHub
$ npm i -g @wavyx/hscli zsh — hscli $ hscli conv list --status active --output table ID SUBJECT CUSTOMER STATUS 4831 Refund for order #2207 amir@acme.io active 4830 Can't reset my password lou@hey.com active 4827 Invoice question kai@studio.co active $ # close everything tagged "resolved" $ hscli conv bulk-status --status active --tag resolved \ --set closed --yes ✓ 12 conversations closed · exit 0 Why hscli
Built for automation, not clicking.
Four design decisions that make hscli safe to hand to a script — or an agent.
JSON output, deterministic exit codes, and a raw hscli api hatch let agents and pipelines drive Help Scout directly.
Every command supports --output table|json|yaml|csv , plus --jq and --fields for inline filtering and projection.
OAuth tokens live only in your OS keychain — never written to disk in plaintext. hscli api is locked to the Help Scout host.
hscli backup dumps your whole account to JSON with incremental refresh, resume, deletion detection, and attachments.
Real work, one line at a time.
hscli composes with the tools you already use — jq , xargs , cron, and your CI runner.
Triage Let an agent clear the queue
Pipe active conversations into your model and act on its decisions — no integration code.
Backup Nightly account snapshot in CI
Incremental by default, resumable, and deletion-aware. Drop it in a cron job or a GitHub Action.
Report Pull metrics into a spreadsheet
Stream any report as CSV and pipe it straight into your BI tool or a Google Sheet.
Drive Help Scout from anywhere.
Install once. Script everything. Hand the rest to an agent.
Command reference hs cli MIT licensed · not affiliated with Help Scout Guide Commands GitHub npm Esc Try conv list , backup , or authentication
