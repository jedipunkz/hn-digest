---
source: "https://github.com/kidus-tiliksew/conveyor"
hn_url: "https://news.ycombinator.com/item?id=49399837"
title: "Show HN: I make my AI agents file paperwork before they're allowed to code"
article_title: "GitHub - kidus-tiliksew/conveyor: Intent on paper, agents on the line, and every merge remembers why. · GitHub"
image: "https://opengraph.githubassets.com/c730641d801ffd4a0aa540140ec9d453a1772a31f5c2a1c10ed8ad522463586b/kidus-tiliksew/conveyor"
author: "kidustiliksew"
captured_at: "2026-08-22T14:11:51Z"
capture_tool: "hn-digest"
hn_id: 49399837
score: 1
comments: 0
posted_at: "2026-08-22T14:01:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I make my AI agents file paperwork before they're allowed to code

- HN: [49399837](https://news.ycombinator.com/item?id=49399837)
- Source: [github.com](https://github.com/kidus-tiliksew/conveyor)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T14:01:46Z

## Translation

タイトル: HN を表示: AI エージェントにコーディングを許可する前に書類を作成させています
記事のタイトル: GitHub - kidus-tiliksew/conveyor: 紙上の意図、ライン上のエージェント、そしてすべてのマージがその理由を覚えています。 · GitHub
説明: 紙上の意図、オンラインのエージェント、そしてすべてのマージがその理由を覚えています。 - kidus-tiliksew/コンベヤー

記事本文:
GitHub - kidus-tiliksew/conveyor: 紙上の意図、オンラインのエージェント、そしてすべてのマージがその理由を覚えています。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
キドゥス・ティリクソー
/
コンベア
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
1,233 コミット 1,233 コミット フォルダーとファイル
.agents/ plugins .agents/ plugins .claude .claude .github/ workflows .github/ workflows .zed .zed cmd cmd docs docs 内部内部パック パックプラグイン/ c

onveyor プラグイン/コンベア スクリプト スクリプト テスト結果 テスト結果 Web Web .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md compose.dev.yaml compose.dev.yaml compose.yaml compose.yaml コンベア.example.yaml コンベア.example.yaml go.mod go.mod go.sum go.sum install.sh install.sh sqlc.yaml sqlc.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コードの生成がこれまでより簡単になりました。コードであることを確認すると、
製品の意図と一致することが現在ボトルネックとなっており、
エージェントの速度が上がるにつれて、状況は悪化します。監視されていないエージェントのキューが送信される可能性があります。
1 日当たりの未読コードの数が、チームが読み取れる量を超えています。
ソフトウェア ファクトリは、このような問題に対する答えです。そうですよね
すべてのネジを検査するわけではありません。検査が行われるようにプロセスを修正します。
間違いが発生する可能性のあるポイントを特定し、すべてのユニットを追跡可能にします
そうすることで、何か問題が発生したときに、他に何が影響しているかを知ることができます。
Conveyor は、エージェントが作成したコードのためのソフトウェア ファクトリです。
要件、システム設計文書、意思決定からの作業をキューに入れます。
人間のオペレーターが文書を確認し、必要に応じて計画を承認します。
マシン上のエージェントは、その作業を計画、実装、およびレビューします。
Conveyor は 2026 年 7 月以来、このプロセスを使用して自社を構築してきました。
コンベヤーは、各変更を文書、タスク、レビュー、およびテストの証拠にリンクします
その後ろに。
フローチャート LR
Intent["確認された要件<br/>と設計"] --> task["タスク"]
タスク --> 配信["配信された変更"]
インテント -.-> check{"位置ずれチェック"}
配達 -.-> チェック
リポジトリ["監視されたリポジトリ"] -.-> チェック
チェック -->|不一致が見つかりました|シグナル["シグナル"]
シグナル --> フォローアップ[「判断またはゲート付きフォローアップ」]
フォローアップ -->||工場に再入|タスク
読み込み中
コンベヤーは確認された要求に対して各配送をチェックします

政府と統治
デザイン。配信とその確認された意図が一致しない場合、Conveyor は
信号。リポジトリのドリフトやマージ後の障害も信号を発生させます。コンベア
コードやドキュメントを単独で書き換えることはありません。オペレータは、
通常のゲートを介して信号を送るか、フォローアップ作業を送信します。
ブラウザ内のオペレータ Codex や Claude などのエージェント
| |
React ダッシュボード MCP 作業指示サーバー
ボード、タスク、ドキュメントの請求、計画、レビュー
| |
+-----------------+-------------------+
|
Goで伝達される
REST API、ダッシュボード、イベントログ
|
PostgreSQL
イベント、ドキュメント、リンク、キュー
|
機械上のコンベヤー作業者
エージェントの CLI を監視します
|
Git ワークツリー、リポジトリ、PR
コンベアは 1 つの Go バイナリです。 PostgreSQL はイベント ログ、ドキュメント、
血統投影とリバーキュー。ワーカーはエージェント CLI を起動します。
ローカル資格情報。
ファクトリ ホストには PostgreSQL 15 以降、Git、認証された gh CLI、
OpenAI 互換モデルのエンドポイントの API キー、およびエージェント CLI
走る計画を立てる。
最新のコンベアおよびコンベアされたバイナリを ~/.local/bin にインストールします。
カール -fsSL https://raw.githubusercontent.com/kidus-tiliksew/conveyor/main/install.sh |しー
インストーラはバイナリを置き換える前にリリース チェックサムを検証します。
sudo は必要ありません。レビュー済みバージョンを固定し、ソースから構築し、
アップグレードについては、「インストール」で説明します。
そこからGetting starting (solo)が立ち上がる
工場でのエンドツーエンドの 1 台のマシン、および
はじめに (マルチプレイヤー) のカバー
共有チームサーバー。
完全なドキュメントは docs/ にあります。まだドキュメント サイトはありません。
インストール : リリースインストーラー、ソースビルド、前提条件
入門 (ソロ) : 1 人、1 台のマシン、エンドツーエンド
はじめに (マルチプレイヤー) : 共有チームサーバー
CLI リファレンス : すべてのコンベアとコンベアコマンド
オーセン

ティケーション: サインイン、トークン、ロール、GitHub ID
構成: 3 つの構成サーフェスとすべての環境変数
MCP リファレンス : エージェントがタスクを実行するために使用するツール
概念: ソフトウェア ファクトリ、ナレッジ グラフ、明暗ファクトリ パターン
文書コーパス: 要件、システム設計、意思決定、提案と確認のサイクル
タスク: 作業がどのように作成され、コンテキストが与えられ、実行され、レビューされ、リンクされるか
不整合 : ドリフト、失効、保留中の提案
ワーカーの操作: 永続的なワーカーの登録、サービスのインストール、リカバリ
GitHub ライフサイクル : 問題、PR、レビュー ステータスがどのように GitHub に投影されるか
既知の制限: 現在の実装で受け入れられている境界
プレイブック (コンベア スキルのインストールでエージェント スキルとしてインストール可能)
計画: ローカル エージェント セッションからドキュメントのドラフトとプッシュを行う
タスクファイリング: 適切な形式のタスクと依存関係チェーンをファイルします。
タスクの作業: 請求、チェックアウト、送信、レビューのライフサイクル
コンベヤは鋭意開発中です。そのイベント ログには欠陥が記録され、
マージの成功と同時に調整作業が行われます。
Conveyor は MIT ライセンスの下で利用可能です。
紙上の意図、電話中のエージェント、そしてすべてのマージがその理由を思い出します。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Intent on paper, agents on the line, and every merge remembers why. - kidus-tiliksew/conveyor

GitHub - kidus-tiliksew/conveyor: Intent on paper, agents on the line, and every merge remembers why. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
kidus-tiliksew
/
conveyor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1,233 Commits 1,233 Commits Folders and files
.agents/ plugins .agents/ plugins .claude .claude .github/ workflows .github/ workflows .zed .zed cmd cmd docs docs internal internal pack pack plugins/ conveyor plugins/ conveyor scripts scripts test-results test-results web web .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md compose.dev.yaml compose.dev.yaml compose.yaml compose.yaml conveyor.example.yaml conveyor.example.yaml go.mod go.mod go.sum go.sum install.sh install.sh sqlc.yaml sqlc.yaml View all files Repository files navigation
Generating code is easier than ever. Checking that the code
matches product intent is now the bottleneck, and it is a bottleneck that
gets worse as agents get faster: a queue of unsupervised agents can ship
more unread code per day than a team can read.
A software factory is the answer to this shape of problem. You do
not inspect every screw; you fix the process so that inspection happens at
the points where mistakes can happen, and you make every unit traceable
so that when something is wrong you know what else is affected.
Conveyor is a software factory for agent-written code.
It queues work from Requirements, System Design documents, and Decisions.
Human operators confirm the documents, and approve plans when required.
Agents on your machines plan, implement, and review that work.
Conveyor has used this process to build itself since July 2026.
Conveyor links each change to the documents, task, review, and test evidence
behind it.
flowchart LR
intent["Confirmed requirements<br/>and designs"] --> task["Task"]
task --> delivery["Delivered change"]
intent -.-> check{"Misalignment checks"}
delivery -.-> check
repository["Observed repository"] -.-> check
check -->|mismatch found| signal["Signal"]
signal --> followup["Judgment or gated follow-up"]
followup -->|re-enters the factory| task
Loading
Conveyor checks each delivery against confirmed requirements and governing
designs. When a delivery and its confirmed intent disagree, Conveyor raises a
signal. Repository drift and post-merge failures raise signals too. Conveyor
never rewrites code or documents on its own. An operator can acknowledge the
signal or send follow-up work through the normal gates.
Operators in a browser Agents such as Codex or Claude
| |
React dashboard MCP work-order server
board, tasks, docs claim, plan, review
| |
+------------------+-------------------+
|
conveyord in Go
REST API, dashboard, event log
|
PostgreSQL
events, documents, links, queue
|
conveyor worker on your machine
supervises your agent CLIs
|
Git worktrees, repositories, PRs
conveyord is one Go binary. PostgreSQL stores the event log, documents,
lineage projection, and River queue. The worker launches agent CLIs with your
local credentials.
A factory host needs PostgreSQL 15 or newer, Git, an authenticated gh CLI,
an API key for an OpenAI-compatible model endpoint, and the agent CLIs you
plan to run.
Install the latest conveyor and conveyord binaries into ~/.local/bin :
curl -fsSL https://raw.githubusercontent.com/kidus-tiliksew/conveyor/main/install.sh | sh
The installer verifies the release checksum before replacing either binary
and does not need sudo . Pinning a reviewed version, building from source,
and upgrades are covered in Installation .
From there, Getting started (solo) stands up
a factory end to end on one machine, and
Getting started (multiplayer) covers
a shared team server.
Full docs live in docs/ . There's no docs site yet.
Installation : release installer, source builds, prerequisites
Getting started (solo) : one person, one machine, end to end
Getting started (multiplayer) : a shared team server
CLI reference : every conveyor and conveyord command
Authentication : sign-in, tokens, roles, GitHub identity
Configuration : the three config surfaces and every environment variable
MCP reference : the tools an agent uses to work a task
Concepts : the software factory, the knowledge graph, light and dark factory patterns
The document corpus : requirements, System Designs, decisions, and the propose-confirm cycle
Tasks : how work is created, given context, executed, reviewed, and linked
Misalignment : drift, staleness, and pending proposals
Worker operations : durable worker enrollment, service install, recovery
GitHub lifecycle : how issues, PRs, and review statuses are projected onto GitHub
Known limitations : accepted boundaries of the current implementation
Playbooks (installable as agent skills with conveyor skills install )
Planning : draft and push documents from a local agent session
Task filing : file well-formed tasks and dependency chains
Working a task : the claim, checkout, submit, review lifecycle
Conveyor is under active development. Its event log records defects and
reconciliation work alongside successful merges.
Conveyor is available under the MIT License .
Intent on paper, agents on the line, and every merge remembers why.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
