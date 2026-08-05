---
source: "https://github.com/moon2022alakhali-commits/agentguard"
hn_url: "https://news.ycombinator.com/item?id=49188739"
title: "AgentGuard – fail-closed approval gateway for AI agent tool calls"
article_title: "GitHub - moon2022alakhali-commits/agentguard: Fail-closed approval gateway for AI agent tool calls · GitHub"
author: "yhouseff_dev"
captured_at: "2026-08-05T20:57:28Z"
capture_tool: "hn-digest"
hn_id: 49188739
score: 1
comments: 0
posted_at: "2026-08-05T20:43:29Z"
tags:
  - hacker-news
  - translated
---

# AgentGuard – fail-closed approval gateway for AI agent tool calls

- HN: [49188739](https://news.ycombinator.com/item?id=49188739)
- Source: [github.com](https://github.com/moon2022alakhali-commits/agentguard)
- Score: 1
- Comments: 0
- Posted: 2026-08-05T20:43:29Z

## Translation

タイトル: AgentGuard – AI エージェント ツール呼び出し用のフェールクローズ承認ゲートウェイ
記事タイトル: GitHub - Moon2022alakhali-commits/agentguard: AI エージェント ツール呼び出しのフェールクローズ承認ゲートウェイ · GitHub
説明: AI エージェント ツール呼び出しのフェールクローズ承認ゲートウェイ - Moon2022alakhali-commits/agentguard

記事本文:
GitHub - Moon2022alakhali-commits/agentguard: AI エージェント ツール呼び出しのフェールクローズ承認ゲートウェイ · GitHub
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
Moon2022alakhali-コミット
/
エージェントガード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット README.md README.md エージェントガード_デモ.py エージェントガード_デモ.py デモ.gif デモ.gif すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント向けのフェールクローズ承認ゲートウェイ

t ツール呼び出し。
エージェントが実際の処理 (データの削除、支払いの送信、アクセスの取り消し、副作用のある API のヒット) を実行するツールを呼び出すことができる場合、最終的には、人間が承認するまで危険な呼び出しを一時停止し、何も承認されない場合はデフォルトで拒否するレイヤーが必要になります。 「ログと希望」ではありません。実はブロックされてる。
このリポジトリは、そのゲートウェイの最小限の、依存関係のないデモです。これはまだ完全な SDK 統合ではありません。これは、内部で実行するスクリプトから抽出されたコア ロジックであり、どのように動作するかを正確に確認できるようにパッケージ化されています。
SAFE_TOOLS -- ホワイトリストに登録され、自動承認され、人間は必要ありません。
DESTRUCTIVE_TOOLS -- 実行前に常に人間による明示的な承認が必要です。
それ以外のもの (不明なツール、承認なし、承認タイムアウト) -- 自動的に拒否されます。フェールオープンではなくフェールクローズです。
python3 エージェントガード_デモ.py
依存関係はありません。まさに標準ライブラリ。
def authorization_gateway (call : ToolCall , * , allowed_by_human : bool = False ) -> 決定 :
電話すれば。 SAFE_TOOLS の名前:
決定を返す ( allowed = True 、reason = "whitelisted_safe_tool" )
電話すれば。 DESTRUCTIVE_TOOLS にある名前ですが、 allowed_by_human ではありません:
return 決定 (許可 = False 、理由 = "destructive_action_unapproved" )
人間によって承認された場合:
return 決定 ( allowed = True 、reason = "explicit_human_approval" )
# デフォルト: 不明なツール、承認なし -> 拒否。これはフェールクローズ部分です。
決定を返します (許可 = False 、理由 = "unknown_tool_no_approval")
なぜ
ほとんどのエージェント フレームワークでは承認フックを追加できますが、何も応答しない場合のデフォルトは通常、許可 (フェールオープン) またはログ行だけです。それは破壊的なものにとっては逆行するものです。これにより、デフォルトが反転されます。
早い。これは抽出されたコアであり、まだパッケージ化されたライブラリではありません。LangChain/OpenAI Agents SDK アダプター、永続性、非同期承認フロー (Slack/Telegram ボットなど) はありません。

）ここに含まれています。それが役立つ場合は、問題やディスカッションを開いて、これをさらに構築する価値があるかどうか、および優先的に統合する必要があるかどうかを判断してください。
LangChain や特定のエージェント フレームワークには依存しません。すでにあるツール呼び出しループの前に同じパターンをドロップします。
AI エージェント ツール呼び出しのフェールクローズ承認ゲートウェイ
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Fail-closed approval gateway for AI agent tool calls - moon2022alakhali-commits/agentguard

GitHub - moon2022alakhali-commits/agentguard: Fail-closed approval gateway for AI agent tool calls · GitHub
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
moon2022alakhali-commits
/
agentguard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit README.md README.md agentguard_demo.py agentguard_demo.py demo.gif demo.gif View all files Repository files navigation
A fail-closed approval gateway for AI agent tool calls.
If your agent can call tools that do real things (delete data, send payments, revoke access, hit an API with side effects), you eventually want a layer that pauses risky calls until a human approves them -- and denies by default if nothing approves it. Not "log and hope." Actually blocked.
This repo is a minimal, dependency-free demo of that gateway. It's not a full SDK integration yet -- it's the core logic, extracted from a script I run internally, packaged so you can see exactly how it behaves.
SAFE_TOOLS -- whitelisted, auto-approved, no human needed.
DESTRUCTIVE_TOOLS -- always requires explicit human approval before running.
Anything else (unknown tool, no approval, approval times out) -- denied automatically . Fail-closed, not fail-open.
python3 agentguard_demo.py
No dependencies. Just the standard library.
def approval_gateway ( call : ToolCall , * , approved_by_human : bool = False ) -> Decision :
if call . name in SAFE_TOOLS :
return Decision ( allowed = True , reason = "whitelisted_safe_tool" )
if call . name in DESTRUCTIVE_TOOLS and not approved_by_human :
return Decision ( allowed = False , reason = "destructive_action_unapproved" )
if approved_by_human :
return Decision ( allowed = True , reason = "explicit_human_approval" )
# Default: unknown tool, no approval -> deny. This is the fail-closed part.
return Decision ( allowed = False , reason = "unknown_tool_no_approval" )
Why
Most agent frameworks let you add approval hooks, but the default when nothing responds is usually allow (fail-open) or just a log line. That's backwards for anything destructive. This flips the default.
Early. This is the extracted core, not a packaged library yet -- no LangChain/OpenAI Agents SDK adapter, no persistence, no async approval flow (Slack/Telegram bot, etc.) included here. If that's useful to you, open an issue or a discussion -- trying to figure out if this is worth building out further and what the priority integration should be.
No dependency on LangChain or any specific agent framework. Drop the same pattern in front of whatever tool-calling loop you already have.
Fail-closed approval gateway for AI agent tool calls
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
