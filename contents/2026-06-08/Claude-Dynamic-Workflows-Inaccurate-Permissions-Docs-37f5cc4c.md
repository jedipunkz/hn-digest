---
source: "https://www.promptarmor.com/resources/claude-dynamic-workflows-use-incorrect-permissions"
hn_url: "https://news.ycombinator.com/item?id=48448501"
title: "Claude Dynamic Workflows Inaccurate Permissions Docs"
article_title: "Claude Dynamic Workflows Inaccurate Permissions Docs"
author: "hackerBanana"
captured_at: "2026-06-08T18:57:30Z"
capture_tool: "hn-digest"
hn_id: 48448501
score: 1
comments: 0
posted_at: "2026-06-08T17:42:07Z"
tags:
  - hacker-news
  - translated
---

# Claude Dynamic Workflows Inaccurate Permissions Docs

- HN: [48448501](https://news.ycombinator.com/item?id=48448501)
- Source: [www.promptarmor.com](https://www.promptarmor.com/resources/claude-dynamic-workflows-use-incorrect-permissions)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T17:42:07Z

## Translation

タイトル: Claude Dynamic Workflows の不正確な権限に関するドキュメント
説明: クロード ダイナミック ワークフロー サブエージェントは、文書化された保証人に違反して、昇格されたアクセス許可で実行される可能性があります。

記事本文:
コーデックスの「自動レビュー」エージェントがマルウェアを実行
パッチが適用されていない Ollama の脆弱性: フィッシング オーバーレイとデータ漏洩
Google スプレッドシート用 ChatGPT がワークブックを流出
Codex for Everything が接続されたデータを漏洩する
Microsoft Copilot Cowork がファイルを流出
Ramp のシート AI が財務情報を窃取
Snowflake Cortex AI がサンドボックスを脱出してマルウェアを実行
GitHub Copilot CLI がマルウェアをダウンロードして実行する
メッセージング アプリのエージェントからのデータ流出
クロード・コワークがファイルを流出
超人的な AI がメールを窃取
IBM AI (「ボブ」) がマルウェアをダウンロードして実行
ハグフェイスチャットでデータが流出
vLex での画面乗っ取り攻撃 (10 億ドルで取得された合法 AI)
Google Antigravity がデータを流出
CellShock: クロード AI はデータを盗むのに優れています
挿入されたマーケットプレイス プラグインを介したクロード コードのハイジャック
間接プロンプトインジェクションによる Slack AI からのデータ抽出
間接プロンプトインジェクションによる Writer.com からのデータ抽出
OWASP における LLM トップ 10 のケーススタディ
クロードの動的ワークフローの不正確な権限に関するドキュメント
各 Microsoft Copilot のリスクは何ですか?
AI アプリケーションにおけるコネクタのリスク
どのドメインを許可リストに追加する必要がありますか?
Microsoft Copilot Cowork の保護: セキュリティ実務者向けガイド
Excel と Google スプレッドシートの AI: 即時挿入とデータ漏洩のリスク
あらゆるプラットフォームとユースケースで Codex を安全に構成する
カーソルの保護: セキュリティ担当者向けガイド
Claude Cowork を安全に実装する
Granola AI のセキュリティ リスクと修復
政府請負業者のための人為的代替手段
悪意のある構成ファイルに対する OpenAI Codex PSA
クロードの動的ワークフローの不正確な権限に関するドキュメント
Claude Dynamic Workflow サブエージェントは、文書化された保証人に違反して、昇格されたアクセス許可で実行される可能性があります。
Claude Dynamic Workflows によって生成されたサブエージェントはコマンド approv を継承します

ドキュメントには「ワークフローが生成するサブエージェントは、ユーザーのセッション モードに関係なく、常に acceptEdits モードで実行される」と明示的に記載されているにもかかわらず、ユーザーのセッションからのモードは変更されません。
注: 「acceptEdits」は、ユーザーの承認なしで限られたファイル編集のみを許可する制限付きモードです。
その結果、ワークフローによって生成されたサブエージェントが、Auto モードや BypassPermissions モードなど、意図しない昇格されたアクセス許可で実行される可能性があります。これにより、信頼できないシェル コマンドの実行、MCP の呼び出し、ネットワークの出力、サンドボックス外での編集、および機密の保護されたファイル パスへの編集のリスクがさらされます。
これは、Claude Code の最新バージョン、バージョン 2.1.168 で検証されています。 Claude Dynamic Workflows は、2026 年 6 月 8 日以降、すべてのユーザーに対してデフォルトで有効になります。
組織が動的ワークフローを無効にする方法
組織は、以下で "disableWorkflows": true を設定することで、動的ワークフローへのアクセスを無効にできます。
組織設定 > クロード コード > 管理設定 (settings.json)
または、次のワークフローを無効にします。
組織設定 > クロード コード > ワークフローをオフに切り替えます
動的ワークフローは、次の場所に移動してロール レベルで無効にすることもできます。
組織設定 > ユーザー > ロール > ロールを編集するか新しいロールを作成 > 機能 > クロード コード > ワークフローを無効にします。

## Original Extract

Claude Dynamic Workflow subagents can execute with elevated permissions violating documented garuntees.

Codex 'Auto-review' Agent Runs Malware
Unpatched Ollama Vulnerabilities: Phishing Overlays and Data Exfiltration
ChatGPT for Google Sheets Exfiltrates Workbooks
Codex for Everything Exfiltrates Connected Data
Microsoft Copilot Cowork Exfiltrates Files
Ramp’s Sheets AI Exfiltrates Financials
Snowflake Cortex AI Escapes Sandbox and Executes Malware
GitHub Copilot CLI Downloads and Executes Malware
Data Exfil from Agents in Messaging Apps
Claude Cowork Exfiltrates Files
Superhuman AI Exfiltrates Emails
IBM AI ('Bob') Downloads and Executes Malware
HuggingFace Chat Exfiltrates Data
Screen takeover attack in vLex (legal AI acquired for $1B)
Google Antigravity Exfiltrates Data
CellShock: Claude AI is Excel-lent at Stealing Data
Hijacking Claude Code via Injected Marketplace Plugins
Data Exfiltration from Slack AI via Indirect Prompt Injection
Data Exfiltration from Writer.com via Indirect Prompt Injection
Case Study in OWASP for LLM Top 10
Claude Dynamic Workflows Inaccurate Permissions Docs
What is the risk of each Microsoft Copilot?
The Risks of Connectors in Your AI Applications
What domains should I add to my allowlist?
Securing Microsoft Copilot Cowork: A Security Practitioner's Guide
AI in Excel and Google Sheets: Prompt Injection and Data Exfiltration Risks
Configuring Codex Securely Across Every Platform and Use Case
Securing Cursor: A Security Practitioner's Guide
Implement Claude Cowork Securely
Granola AI Security Risks and Remediations
Anthropic Alternatives for Government Contractors
OpenAI Codex PSA on Malicious Config Files
Claude Dynamic Workflows Inaccurate Permissions Docs
Claude Dynamic Workflow subagents can execute with elevated permissions violating documented garuntees.
Subagents spawned by Claude Dynamic Workflows inherit command approval modes from the user’s session, despite documentation explicitly stating “subagents the workflow spawns always always run in acceptEdits mode… regardless of the user’s session mode”.
Note: ‘acceptEdits’ is a restricted mode that allows only limited file editing without user approval.
As a result, subagents spawned by workflows can execute with unintended elevated permissions, such as in Auto or BypassPermissions modes. This exposes a risk of untrusted shell command execution, MCP invocation, network egress, edits outside the sandbox, and edits to sensitive protected file paths.
This has been validated on the latest version of Claude Code, Version 2.1.168. Claude Dynamic Workflows will be enabled by default for all users as of June 8, 2026.
How Organizations Can Disable Dynamic Workflows
Organizations can disable access to dynamic workflows by setting "disableWorkflows": true in:
Organization settings > Claude Code > Managed settings (settings.json)
Or, by disabling workflows in:
Organization settings > Claude Code > toggle off Workflows
Dynamic workflows can also be disabled at the role level by navigating to:
Organization settings > People > Roles > edit a role or create a new one > Capabilities > Claude Code > disable Workflows.
