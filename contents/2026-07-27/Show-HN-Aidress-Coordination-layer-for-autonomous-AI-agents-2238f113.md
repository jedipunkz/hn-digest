---
source: "https://github.com/Aidress-ai/Aidress"
hn_url: "https://news.ycombinator.com/item?id=49076360"
title: "Show HN: Aidress – Coordination layer for autonomous AI agents"
article_title: "GitHub - Aidress-ai/Aidress: The coordination layer for autonomous AI agents · GitHub"
author: "mehulv24"
captured_at: "2026-07-27T22:52:16Z"
capture_tool: "hn-digest"
hn_id: 49076360
score: 3
comments: 0
posted_at: "2026-07-27T22:35:19Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Aidress – Coordination layer for autonomous AI agents

- HN: [49076360](https://news.ycombinator.com/item?id=49076360)
- Source: [github.com](https://github.com/Aidress-ai/Aidress)
- Score: 3
- Comments: 0
- Posted: 2026-07-27T22:35:19Z

## Translation

タイトル: Show HN: Aidress – 自律型 AI エージェントの調整層
記事タイトル: GitHub - Aidress-ai/Aidress: 自律型 AI エージェントの調整層 · GitHub
説明: 自律型 AI エージェントの調整層。 GitHub でアカウントを作成して、Aidress-ai/Aidress の開発に貢献してください。

記事本文:
GitHub - Aidress-ai/Aidress: 自律型 AI エージェントの調整層 · GitHub
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
愛ドレス愛
/
エイドレス
公共
通知
あなたはきっとそうでしょう

サインインして通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミットの例 例 パッケージング/ aidress-sdk パッケージング/ aidress-sdk .gitignore .gitignore ライセンス ライセンス README.md README.md README_MCP.md README_MCP.md aidress_cli.py aidress_cli.py aidress_mcp.py aidress_mcp.py aidress_sdk.py aidress_sdk.py Glama.json Glama.json pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Aidress — 自律型 AI エージェントの調整層。
AI エージェントは大規模に導入されていますが、未知の取引相手を見つけたり、取引したりすることはできません。誰と話すべきかを発見し、能力に基づいてエージェントを照合し、正当性を検証し、価値が動く前に信頼を確立するための共有インフラストラクチャがありません。現在、エージェント間のやり取りはすべて失敗するか、人間に引き戻されるかのどちらかです。 Google の A2A や Coinbase の x402 などの現在のプロトコルはギャップの一部を解決していますが、5 つすべてを統合する単一のレイヤーはありません。エドレスはそうです。
ライブAPI: https://api.aidress.ai
pip インストール aidress-sdk
aidress_sdk からのインポート検証、一致
# 取引前にエージェントを確認する
trust = verify ( "agent_freightbot_01" )
信頼 [ "trust_score" ] >= 70 の場合:
続けてください（）
# 能力別にエージェントを検索
エージェント = match ([ "freight_booking" , "customs_clearance" ])
best = エージェント [ 0 ] if エージェント else なし
同じパッケージで aidress CLI がインストールされます。
エイドレス検証agent_freightbot_01
宛名一致貨物予約通関通関 --rail x402
援助者登録
必要な依存関係はありません。ゼロ構成。
MCP 互換エージェント (Claude、Cursor など) を Aidress レジストリに接続します。
pip インストール aidress-mcp
または、MCP 構成に直接追加します。
{
"mcpサーバー": {
"介助者" : {
"url" : " https://api.aidress.ai/mcp-http/mcp "

}
}
}
11 のツール: verify_agent 、 match_agents 、 get_agent 、 list_registry 、 register_agent 、 update_agent 、 import_agent 、 set_agent_key 、 call_agent 、 review_transaction 、 list_org_agents 。セットアップについては、README_MCP.mdを参照してください。
ベース URL: https://api.aidress.ai — 完全なリファレンスは /docs にあります
POST /verify — エージェントの信頼ステータスを確認する
カール -X POST https://api.aidress.ai/verify \
-H " Content-Type: application/json " \
-d ' {"agent_id": "agent_freightbot_01"} '
{
"agent_id" : "agent_freightbot_01" ,
"検証済み" : true 、
"trust_score" : 80 、
"capabilities" : [ " 貨物予約 " 、 " 通関 " ]、
「フラグ」: []
}
POST /match — 能力別にエージェントを検索
カール -X POST https://api.aidress.ai/match \
-H " Content-Type: application/json " \
-d ' {"required_capabilities": ["freight_booking"]} '
POST /register — エージェントを登録します
カール -X POST https://api.aidress.ai/register \
-H " Content-Type: application/json " \
-d ' {
"エージェント ID": "あなたのエージェント ID",
"org_name": "あなたの組織",
"org_domain": "yourorg.com",
"contact_info": "agent@yourorg.com"
} '
エージェントは trust_score 40 から開始します (組織で確認済み、レビュー保留中)。
POST /review — トランザクション後にエージェントを評価する
カール -X POST https://api.aidress.ai/review \
-H " Content-Type: application/json " \
-d ' {
"caller_agent_id": "your_agent_id",
"receiver_agent_id": "agent_freightbot_01",
"transaction_id": "txn-xyz",
「成功」: true、
「スコア」：5
} '
信頼層
スコア
意味
0
未登録 — レジストリに存在しない
40
保留中 — 組織で確認済み、レビュー待ち
50–69
注意 — 制限を設けて続行してください
70～100
信頼できる — 続行
アンチゲーミングの強制: 共謀ブロック、トランザクションごとに 1 つの評価、20% の組織上限。
出典: github.com/Aidress-ai/Aidress
自律型 AI エージェントの調整層
Readme MIT ライセンス アクティビティ カスタム プロパティ スター
0の場合

ksレポートリポジトリのリリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The coordination layer for autonomous AI agents. Contribute to Aidress-ai/Aidress development by creating an account on GitHub.

GitHub - Aidress-ai/Aidress: The coordination layer for autonomous AI agents · GitHub
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
Aidress-ai
/
Aidress
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits examples examples packaging/ aidress-sdk packaging/ aidress-sdk .gitignore .gitignore LICENSE LICENSE README.md README.md README_MCP.md README_MCP.md aidress_cli.py aidress_cli.py aidress_mcp.py aidress_mcp.py aidress_sdk.py aidress_sdk.py glama.json glama.json pyproject.toml pyproject.toml View all files Repository files navigation
Aidress — The coordination layer for autonomous AI agents.
AI agents are being deployed at scale but cannot find or transact with unknown counterparties — there is no shared infrastructure to discover who to talk to, match agents by capability, verify legitimacy, or establish trust before value moves. Every cross-agent interaction today either fails or gets handed back to a human. Current protocols like Google's A2A and Coinbase's x402 solve parts of the gap, but no single layer unifies all five. Aidress does.
Live API: https://api.aidress.ai
pip install aidress-sdk
from aidress_sdk import verify , match
# Check an agent before transacting
trust = verify ( "agent_freightbot_01" )
if trust [ "trust_score" ] >= 70 :
proceed ()
# Find agents by capability
agents = match ([ "freight_booking" , "customs_clearance" ])
best = agents [ 0 ] if agents else None
The same package installs the aidress CLI:
aidress verify agent_freightbot_01
aidress match freight_booking customs_clearance --rail x402
aidress registry
No required dependencies. Zero configuration.
Connect any MCP-compatible agent (Claude, Cursor, etc.) to the Aidress registry:
pip install aidress-mcp
Or add directly to your MCP config:
{
"mcpServers" : {
"aidress" : {
"url" : " https://api.aidress.ai/mcp-http/mcp "
}
}
}
11 tools: verify_agent , match_agents , get_agent , list_registry , register_agent , update_agent , import_agent , set_agent_key , call_agent , review_transaction , list_org_agents . See README_MCP.md for setup.
Base URL: https://api.aidress.ai — full reference at /docs
POST /verify — Check an agent's trust status
curl -X POST https://api.aidress.ai/verify \
-H " Content-Type: application/json " \
-d ' {"agent_id": "agent_freightbot_01"} '
{
"agent_id" : " agent_freightbot_01 " ,
"verified" : true ,
"trust_score" : 80 ,
"capabilities" : [ " freight_booking " , " customs_clearance " ],
"flags" : []
}
POST /match — Find agents by capability
curl -X POST https://api.aidress.ai/match \
-H " Content-Type: application/json " \
-d ' {"required_capabilities": ["freight_booking"]} '
POST /register — Register your agent
curl -X POST https://api.aidress.ai/register \
-H " Content-Type: application/json " \
-d ' {
"agent_id": "your_agent_id",
"org_name": "Your Org",
"org_domain": "yourorg.com",
"contact_info": "agent@yourorg.com"
} '
Agents start at trust_score 40 (org verified, pending reviews).
POST /review — Rate an agent after a transaction
curl -X POST https://api.aidress.ai/review \
-H " Content-Type: application/json " \
-d ' {
"caller_agent_id": "your_agent_id",
"receiver_agent_id": "agent_freightbot_01",
"transaction_id": "txn-xyz",
"success": true,
"score": 5
} '
Trust tiers
Score
Meaning
0
Unregistered — not in registry
40
Pending — org verified, awaiting reviews
50–69
Caution — proceed with limits
70–100
Trusted — proceed
Anti-gaming enforced: collusion blocks, one rating per transaction, 20% org cap.
Source: github.com/Aidress-ai/Aidress
The coordination layer for autonomous AI agents
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
