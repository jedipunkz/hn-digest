---
source: "https://github.com/swe-workflow/log-decisions"
hn_url: "https://news.ycombinator.com/item?id=49264597"
title: "Show HN: Log-decisions – an append-only decision journal for AI agents"
article_title: "GitHub - swe-workflow/log-decisions: Append-only DECISIONS.md journal for the consequential calls a spec didn't settle — an Agent Skill for Claude Code, Codex, Gemini CLI, Cursor, and other skills hosts · GitHub"
author: "soulmachine"
captured_at: "2026-08-11T21:35:46Z"
capture_tool: "hn-digest"
hn_id: 49264597
score: 2
comments: 0
posted_at: "2026-08-11T21:15:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Log-decisions – an append-only decision journal for AI agents

- HN: [49264597](https://news.ycombinator.com/item?id=49264597)
- Source: [github.com](https://github.com/swe-workflow/log-decisions)
- Score: 2
- Comments: 0
- Posted: 2026-08-11T21:15:55Z

## Translation

タイトル: Show HN: Log-decions – AI エージェント向けの追加専用の意思決定ジャーナル
記事のタイトル: GitHub - swe-workflow/log-decions: 仕様が解決しなかった結果的な呼び出しの追加専用 DECISIONS.md ジャーナル — クロード コード、コーデックス、Gemini CLI、カーソル、およびその他のスキル ホストのエージェント スキル · GitHub
説明: 仕様が決着しなかった結果的な呼び出しに関する追加専用の DECISIONS.md ジャーナル — クロード コード、コーデックス、Gemini CLI、カーソル、およびその他のスキル ホストのエージェント スキル - swe-workflow/log-decions

記事本文:
GitHub - swe-workflow/log-decions: 仕様が解決しなかった結果的な呼び出しに対する追加専用の DECISIONS.md ジャーナル — クロード コード、コーデックス、Gemini CLI、カーソル、およびその他のスキル ホストのエージェント スキル · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
SWワークフロー
/
ログ決定
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コム

mits .claude-plugin .claude-plugin スキル/ログ決定 スキル/ログ決定 ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
仕様が解決されなかった結果的な呼び出しを記録する追加専用の DECISIONS.md ジャーナル。エージェントが単独で解決できるものについての決定/想定/エスカレーション ルールを使用して、何を、誰が、そしてなぜ選択したかを記録します。
汎用。このジャーナルはコーディングに特化したものではありません。調査、執筆、運用、データなど、エージェントが支援する作業はすべて、後で誰かが「なぜこのように行われたのか?」と尋ねる判断を生み出すことになります。について。このスキルにより、これらの呼び出しに、リポジトリ ルートにある耐久性のある grep 可能なホームが 1 つ与えられます。
これはエージェント スキル (純粋なマークダウン、スクリプトなし) として出荷されるため、Claude Code、Codex、Gemini CLI、Cursor、およびその他のスキル互換エージェント上で実行されます。
ユニバーサル (スキル互換性のあるエージェント)
npx スキルで swe-workflow/log-decions を追加
クロードコード（プラグイン）
/plugin マーケットプレイスに swe-workflow/log-decions を追加
/plugin install log-decions@log-decions
何をするのか
スキル ( skill/log-decions/SKILL.md ) は次のように定義します。
バー — 仕様がすでに呼び出しを許可している場合、それはジャーナルに値しません。エージェントが許可を発明しなければならなかったとしても、それはその通りです。
2×2 — 各呼び出しを決定可能 × 可逆として分類し、決定 (アーティファクトによる根拠)、仮定 (安全なデフォルト、非同期レビュー用に記録)、またはエスカレーション (停止して尋ねる) を行います。データ損失、破壊的な移行、取り消し不能な出費、パブリック インターフェイスの破損など、壊滅的な問題は常にエスカレートします。
エントリ - 決定ごとに 1 つの追加専用ブロック: 質問、検討されたオプション、選択された、決定者、正当性、結果、置き換え付き: 改訂用。決して編集したり、並べ替えたりすることはありません。
Claude Code プラグイン — 上記のコマンドを使用してインストールします。 Anthropic のプラグイン ディレクトリに送信されました (レビュー待ち)
SWワークフロー

— アイデア → PRD → 問題 → 出荷スイート — あらゆる段階でこのスキルを調整します。仕様レイヤーのグリルはそれを通じて想定される回答を記録し、出荷時にビルド ステージのエントリをワークツリーごとの DECISIONS.staged.md に反映し、終了時にプロモートされます。これは、そのリポジトリ (≤ v2.0.2) からこのスタンドアロン リポジトリに抽出されました。
仕様が解決しなかった結果的な呼び出しに関する追加専用の DECISIONS.md ジャーナル — クロード コード、コーデックス、Gemini CLI、カーソル、およびその他のスキル ホストのエージェント スキル
Readme MIT ライセンス アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Append-only DECISIONS.md journal for the consequential calls a spec didn't settle — an Agent Skill for Claude Code, Codex, Gemini CLI, Cursor, and other skills hosts - swe-workflow/log-decisions

GitHub - swe-workflow/log-decisions: Append-only DECISIONS.md journal for the consequential calls a spec didn't settle — an Agent Skill for Claude Code, Codex, Gemini CLI, Cursor, and other skills hosts · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
swe-workflow
/
log-decisions
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .claude-plugin .claude-plugin skills/ log-decisions skills/ log-decisions LICENSE LICENSE README.md README.md View all files Repository files navigation
An append-only DECISIONS.md journal recording the consequential calls a spec didn't settle — what was chosen, by whom, and why, with decide / assume / escalate rules for what an agent may settle alone.
General-purpose. The journal isn't coding-specific: any agent-assisted work — research, writing, ops, data — produces judgment calls someone will later ask "why was it done this way?" about. This skill gives those calls one durable, grep -able home at the repo root.
It ships as an Agent Skill — pure markdown, no scripts — so it runs on Claude Code, Codex, Gemini CLI, Cursor , and any other skills-compatible agent.
Universal (any skills-compatible agent)
npx skills add swe-workflow/log-decisions
Claude Code (plugin)
/plugin marketplace add swe-workflow/log-decisions
/plugin install log-decisions@log-decisions
What it does
The skill ( skills/log-decisions/SKILL.md ) defines:
The bar — if the spec already authorized the call it's not journal-worthy; if the agent had to invent the authorization, it is.
The 2×2 — classify each call as determinable × reversible , then decide (grounded by an artifact), assume (safe default, logged for async review), or escalate (stop and ask). A catastrophic floor — data loss, destructive migration, irreversible spend, public-interface breaks — always escalates.
The entry — one append-only block per decision: Question, Options considered, Chosen, Decided-by, Justification, Outcome, with Supersedes: for revisions. Never edited, never reordered.
Claude Code plugin — install with the commands above; submitted to Anthropic's Plugin Directory (pending review)
swe-workflow — the idea → PRD → issues → ship suite — orchestrates this skill across every stage: spec-layer grills journal their assumed answers through it, and ship builds stage entries to a per-worktree DECISIONS.staged.md promoted at close-out. It was extracted from that repo (≤ v2.0.2) into this standalone one.
Append-only DECISIONS.md journal for the consequential calls a spec didn't settle — an Agent Skill for Claude Code, Codex, Gemini CLI, Cursor, and other skills hosts
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
