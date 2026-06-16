---
source: "https://github.com/NanoPrompt/samf-framework"
hn_url: "https://news.ycombinator.com/item?id=48551132"
title: "SAMF- Deterministic Moscow guardrails for LLM multi-agent loops"
article_title: "GitHub - NanoPrompt/samf-framework: Prompt Engineering framework · GitHub"
author: "nanoprompter"
captured_at: "2026-06-16T08:06:09Z"
capture_tool: "hn-digest"
hn_id: 48551132
score: 1
comments: 0
posted_at: "2026-06-16T06:00:55Z"
tags:
  - hacker-news
  - translated
---

# SAMF- Deterministic Moscow guardrails for LLM multi-agent loops

- HN: [48551132](https://news.ycombinator.com/item?id=48551132)
- Source: [github.com](https://github.com/NanoPrompt/samf-framework)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T06:00:55Z

## Translation

タイトル: SAMF - LLM マルチエージェント ループ用の決定論的モスクワ ガードレール
記事タイトル: GitHub - NanoPrompt/samf-framework: プロンプト エンジニアリング フレームワーク · GitHub
説明: プロンプトエンジニアリングフレームワーク。 GitHub でアカウントを作成して、NanoPrompt/samf フレームワークの開発に貢献してください。

記事本文:
GitHub - NanoPrompt/samf-framework: プロンプト エンジニアリング フレームワーク · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
ナノプロンプト
/
samf フレームワーク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動

e コード [その他のアクション] メニューを開く フォルダーとファイル
3 コミット 3 コミット samf samf .gitignore .gitignore README.md README.md SAMF Library.docx SAMF Library.docx example.py example.py setup.py setup.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
SAMF: SWANT エージェント MoSCoW フレームワーク
決定論的な LLM 検証のための構造的な MoSCoW 契約。
SAMF は、LLM に機械可読な検証コントラクトを提供し、プロジェクト管理 MoSCoW の優先順位付けを厳格な調整ガードに変換します。
プロンプトの不安定性を軽減: 脆弱なマルチホップ推論を排除します。
ループ違反の防止: マルチエージェントの推論ループが壊れないようにします。
トークンコストの最適化: 明示的な制約を強制することで、RAG トークンの無駄を削減します。
samf import SAMFContract 、 samf_contract から
# 厳密な実行契約を定義する
Guard_contract = SAMFContract (
must_have = [ "JSON" , "ステータス" , "user_id" ],
should_have = [ "タイムスタンプ" ],
wont_have = [ "エラー" , "失敗" ]
)
# デコレーターを使用してシームレスに強制します
@ samf_contract (コントラクト = Guard_contract )
def call_billing_agent():
# LLM 生成 / LangGraph ノード ロジックはここにあります
return '{"ステータス": "成功"、"user_id": 4002、"フォーマット": "JSON"}'
応答 = call_billing_agent ()
print ( "出力は正常に検証されました!" )
- - -
** 免責事項: ** これは、完全に個人的な時間とハードウェアに基づいて開発された、独立したオープンソースの趣味のプロジェクトです。それは私の雇用主と提携したり、雇用主によって後援されたり、承認されたりするものではありません。
について
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Prompt Engineering framework. Contribute to NanoPrompt/samf-framework development by creating an account on GitHub.

GitHub - NanoPrompt/samf-framework: Prompt Engineering framework · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
NanoPrompt
/
samf-framework
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits samf samf .gitignore .gitignore README.md README.md SAMF Library.docx SAMF Library.docx example.py example.py setup.py setup.py View all files Repository files navigation
SAMF: SAWANT Agentic MoSCoW Framework
Structural MoSCoW contracts for deterministic LLM validation.
SAMF provides machine-readable validation contracts for LLMs, translating project management MoSCoW prioritization into strict alignment guards.
Slashes Prompt Instability: Eliminates brittle multi-hop reasoning.
Prevents Loop Breaches: Keeps multi-agent reasoning loops from breaking.
Optimizes Token Costs: Reduces RAG token waste by enforcing explicit constraints.
from samf import SAMFContract , samf_contract
# Define your strict execution contract
guard_contract = SAMFContract (
must_have = [ "JSON" , "status" , "user_id" ],
should_have = [ "timestamp" ],
wont_have = [ "error" , "failed" ]
)
# Enforce it seamlessly with a decorator
@ samf_contract ( contract = guard_contract )
def call_billing_agent ():
# Your LLM generation / LangGraph node logic here
return '{"status": "success", "user_id": 4002, "format": "JSON"}'
response = call_billing_agent ()
print ( "Output validated successfully!" )
- - -
* * Disclaimer : ** This is an independent open - source hobby project developed entirely on personal time and hardware . It is not affiliated with , sponsored by , or endorsed by my employer .
About
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
