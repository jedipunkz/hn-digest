---
source: "https://github.com/AML-memory/agent-memory-leaderboard"
hn_url: "https://news.ycombinator.com/item?id=49145408"
title: "Show HN: A fixed harness for comparing LLM agent memory systems"
article_title: "GitHub - AML-memory/agent-memory-leaderboard · GitHub"
author: "IreneAI"
captured_at: "2026-08-02T15:53:13Z"
capture_tool: "hn-digest"
hn_id: 49145408
score: 1
comments: 0
posted_at: "2026-08-02T15:18:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A fixed harness for comparing LLM agent memory systems

- HN: [49145408](https://news.ycombinator.com/item?id=49145408)
- Source: [github.com](https://github.com/AML-memory/agent-memory-leaderboard)
- Score: 1
- Comments: 0
- Posted: 2026-08-02T15:18:00Z

## Translation

タイトル: HN を表示: LLM エージェント メモリ システムを比較するための修正ハーネス
記事のタイトル: GitHub - AML-memory/agent-memory-leaderboard · GitHub
説明: GitHub でアカウントを作成して、AML メモリ/エージェント メモリ リーダーボードの開発に貢献します。

記事本文:
GitHub - AML メモリ/エージェント メモリ リーダーボード · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
AMLメモリ
/
エージェント-メモリ-リーダーボード
公共
通知
c にサインインする必要があります

ハンゲの通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット ビーム ビーム clbench clbench locomo-refined locomo-refined longmemeval-s longmemeval-s personamem personamem scriptmem scriptmem README.md README.md api_config.py api_config.py要件.txt要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
このディレクトリは、によって使用される回答生成および評価コントラクトを公開します。
公開データセットのエージェント メモリ リーダーボード。
スクリプトには Python 3.10 以降と、次の依存関係が必要です。
要件.txt 。サービスの URL、モデル名、および資格情報を指定する必要があります
CLI 引数または環境変数を介して。ここには資格情報はバンドルされていません。
内部短縮ラベル
説明的なパブリック名に正規化されます。解答の構成と採点
挙動は変わらない。
開示された実行時の動作には、追加/検索/応答の同時実行性、タイムアウト、
再試行、キーのローテーション、チェックポイント、および集約。内部展開コード、
サービスアドレス、資格情報、
データベース、データセット、実行アーティファクトは意図的にこのパブリックの外にあります
評価版リリース。
開示された製造パラメータ
フルタスクのタイムアウト: 72 時間。検索top_k : 100。
追加: デフォルトで 64 のグローバル ワーカー、データセットあたり 48 レコードのハード キャップ、20 メッセージ
チャンク、1,200 秒の HTTP タイムアウト、および global-largest-next-record-v1
推定チャンク数を使用したスケジューリング。
検索: ワーカー 32 人、HTTP タイムアウト 1,200 秒、リクエスト試行 6 回、最大
失敗したレコードの再キューは 3 回、5 回以降は適応型同時実行数が削減されます。
再試行可能な 5xx 応答。
回答: ワーカー 32 人、 gpt-4o-mini 、温度 0、HTTP タイムアウト 180 秒、
そして6回の試み。設定されたチェックポイント値は 128 トークンですが、
アウトバウンドプロバイダーの要求

est は意図的に max_tokens を送信しません。
ジャッジ: 構成されたエバリュエーター、360 秒の HTTP タイムアウト、およびグローバル上限
リクエストは15件。ワーカーフェーズの制限は、バイナリ評価の場合は 8、ルーブリックの場合は 15 です
評価、および BEAM イベント アラインメントの場合は 16。グローバルキャップは有効であることを意味します
裁判官リクエストの同時実行数が 15 を超えることはありません。
失敗の判定: HTTP 429、タイムアウト、およびトランスポートの失敗により、キーがローテーションされます。
プール全体の 429 枯渇では、30 秒のクールダウンと最大 3 回のクールダウンが使用されます
ラウンドします。 Qwen JSON リクエストは、enable_ Thinking=false を設定します。
スコアは内部的に [0, 1] で計算され、 [0, 100] で公開されます。
Readme アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to AML-memory/agent-memory-leaderboard development by creating an account on GitHub.

GitHub - AML-memory/agent-memory-leaderboard · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
AML-memory
/
agent-memory-leaderboard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit beam beam clbench clbench locomo-refined locomo-refined longmemeval-s longmemeval-s personamem personamem scriptmem scriptmem README.md README.md api_config.py api_config.py requirements.txt requirements.txt View all files Repository files navigation
This directory publishes the answer-generation and evaluation contracts used by
the Agent Memory Leaderboard for its public datasets.
The scripts require Python 3.10 or newer and the dependency listed in
requirements.txt . Service URLs, model names, and credentials must be supplied
through CLI arguments or environment variables; no credential is bundled here.
Internal shorthand labels
are normalized to descriptive public names; answer construction and scoring
behavior are unchanged.
The disclosed runtime behavior covers Add/Search/Answer concurrency, timeouts,
retries, key rotation, checkpoints, and aggregation. Internal deployment code,
service addresses, credentials,
databases, datasets, and run artifacts are intentionally outside this public
evaluation release.
Disclosed production parameters
Full task timeout: 72 hours; retrieval top_k : 100.
Add: 64 global workers by default, 48-record per-dataset hard cap, 20-message
chunks, 1,200-second HTTP timeout, and global-largest-next-record-v1
scheduling using estimated chunk counts.
Search: 32 workers, 1,200-second HTTP timeout, six request attempts, up to
three failed-record requeues, and adaptive concurrency reduction after five
retryable 5xx responses.
Answer: 32 workers, gpt-4o-mini , temperature 0, 180-second HTTP timeout,
and six attempts. The configured checkpoint value is 128 tokens, but the
outbound provider request intentionally does not send max_tokens .
Judge: a configured evaluator, 360-second HTTP timeout, and a global cap of
15 requests. Worker phase limits are 8 for binary evaluation, 15 for rubric
evaluation, and 16 for BEAM event alignment; the global cap means effective
Judge request concurrency never exceeds 15.
Judge failures: HTTP 429, timeout, and transport failures rotate the key.
Pool-wide 429 exhaustion uses a 30-second cooldown and at most three cooldown
rounds. Qwen JSON requests set enable_thinking=false .
Scores are computed on [0, 1] internally and published on [0, 100] .
Readme Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
