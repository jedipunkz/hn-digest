---
source: "https://github.com/tkngate/tkngate"
hn_url: "https://news.ycombinator.com/item?id=48538373"
title: "The Cloudflare for Autonomous AI Agents"
article_title: "GitHub - tkngate/tkngate: Cloudflare for AI Agents · GitHub"
author: "kilopalisme"
captured_at: "2026-06-15T11:12:14Z"
capture_tool: "hn-digest"
hn_id: 48538373
score: 2
comments: 1
posted_at: "2026-06-15T08:42:05Z"
tags:
  - hacker-news
  - translated
---

# The Cloudflare for Autonomous AI Agents

- HN: [48538373](https://news.ycombinator.com/item?id=48538373)
- Source: [github.com](https://github.com/tkngate/tkngate)
- Score: 2
- Comments: 1
- Posted: 2026-06-15T08:42:05Z

## Translation

タイトル: 自律型 AI エージェントのための Cloudflare
記事タイトル: GitHub - tkngate/tkngate: AI エージェント向け Cloudflare · GitHub
説明: AI エージェント用の Cloudflare。 GitHub でアカウントを作成して、tkngate/tkngate の開発に貢献してください。

記事本文:
GitHub - tkngate/tkngate: AI エージェント用 Cloudflare · GitHub
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
tkngate
/
tkngate
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 他のアクションを開く

メニュー フォルダーとファイル
53 コミット 53 コミット .github/ workflows .github/ workflows cmd cmd docs docs 例 例 内部 内部 .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfileライセンス ライセンス README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum main.go main.go tkngate.example.yaml tkngate.example.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
自律型 AI エージェント用の Cloudflare
ゼロナレッジ リバース プロキシ、マルチプロバイダー フェイルオーバー、P2P トークン メッシュ
インストールします。
Tkngateを選ぶ理由 •
特徴 •
構成 •
ドキュメント
OpenAI がダウンすると、自律エージェントがクラッシュします。不正ループによって 200 ドルのトークンが消費されると、予算が失われます。 API キーは、.env ファイル内のプレーン テキストに保存されます。
Tkngate はこれらすべてを修正します。エージェントと AI プロバイダーの間に位置し、キー、予算、セキュリティ、フェイルオーバーを自動的に管理するインテリジェントなシールドとして機能します。
git クローン https://github.com/tkngate/tkngate.git
cd tkngate
cp tkngate.example.yaml tkngate.yaml # キーを使用して編集します
import TKNGATE_MASTER_KEY= " あなたの-32文字-秘密キー-ここに-!! "
メインを実行してください。サーブしてください
エージェントは、 api.openai.com ではなく http://localhost:7477/openai/chat/completions をポイントするようになりました。
単一のエンドポイントから OpenAI、Anthropic、DeepSeek、Kimi、Groq を介してルーティングします。 1 つのプロバイダー (HTTP 500/502/503) がダウンした場合、Tkngate は自動的に次のプロバイダーにフェイルオーバーします。
プロンプト インジェクション攻撃をブロックし、PII (クレジット カード、SSN、API キー) がプロバイダーに到達する前に自動的に編集します。
ステークアンドスラッシュの評判 (メッシュ)
暗号化された不正防止機能を使用して、寄付された API キーを悪用から保護します。 WAF をバイパスして悪意のあるプロンプトをルーティングするノードは、Stak によってペナルティを受けます。

e-and-Slash の信頼台帳と永久にブラックリストに登録されます。
チーム用に安全な tkngate-sk-... エンタープライズ仮想キーを生成します。各キーはハードバジェットキャップを持つ分離されたサンドボックスとして機能し、物理的なアップストリームキーを保護します。
厳格なレート制限 (トークンバケット)
超低遅延のインメモリ トークン バケット レート リミッタにより、自律エージェントの「バースト ループ」から予算とプロバイダーを保護します。
グリーン→アンバー→レッドゾーンによるリアルタイムの支出追跡。グローバル制限、セッションごとの上限を設定し、予算がなくなった場合のリクエストの自動ブロックを設定します。
分散セマンティック キャッシュ (Redis)
同一のプロンプトが分散 Redis キャッシュから提供され、複数の Tkngate プロキシ ノードにわたる水平スケーリングが可能になります。フリート全体にわたって世界中でトークンとお金を節約します。キャッシュ キーは、正規化されたモデル + メッセージ ハッシュから計算されます。
プロンプト内の Go、Python、JavaScript のコード ブロックを自動的に圧縮し、コメントと空白を削除して、トークンの使用量を最大 40% 削減します。
世界初の LLM API 用の BitTorrent スタイルのトークン プール。予備の API キーをメッシュに寄付すると、停止中にネットワークの容量に優先的にアクセスできるようになります。 AES-256 ゼロ知識暗号化によって保護されています。
本番トラフィックの一部を代替プロバイダー (DeepSeek など) にサイレントにミラーリングし、プライマリ リクエストに対するレイテンシの影響をゼロにして、コスト削減を評価します。
完全なリファレンスについては、tkngate.example.yaml を参照してください。
トピック
リンク
予算制度と信号機
docs/budgeting.md
DRR トークンメッシュとレピュテーション
docs/drr-mesh.md
エンタープライズ仮想キー
docs/virtual-keys.md
厳格なレート制限
docs/rate-limiting.md
ゼロ知識セキュリティ
docs/zero-knowledge-security.md
シャドウモード
docs/shadow-mode.md
建築
エージェントリクエスト
│
▼
┌───────────

───────┐
│ TKNGATE プロキシ │
│ │
│ バジェットガード → AI-WAF/DLP │
│ → コンテキスト圧縮器 │
│ → セマンティックキャッシュ │
│ → 自動再試行 (3 回) │
│ → ユニバーサルルーター（フェイルオーバー） │
│ → シャドウモード（非同期ミラー） │
│ → トークンカウンター → 台帳 │
━━━━━━━━━━━━━━┘
│
▼
OpenAI / Anthropic / DeepSeek / キミ / Groq
ライセンス
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
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

Cloudflare for AI Agents. Contribute to tkngate/tkngate development by creating an account on GitHub.

GitHub - tkngate/tkngate: Cloudflare for AI Agents · GitHub
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
tkngate
/
tkngate
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
53 Commits 53 Commits .github/ workflows .github/ workflows cmd cmd docs docs examples examples internal internal .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum main.go main.go tkngate.example.yaml tkngate.example.yaml View all files Repository files navigation
The Cloudflare for Autonomous AI Agents
Zero-knowledge reverse proxy · Multi-provider failover · P2P token mesh
Install •
Why Tkngate •
Features •
Configuration •
Docs
Your autonomous agents crash when OpenAI goes down. Your budget bleeds when a rogue loop burns $200 in tokens. Your API keys sit in plain text in .env files.
Tkngate fixes all of this. It sits between your agents and the AI providers, acting as an intelligent shield that manages keys, budgets, security, and failover — automatically.
git clone https://github.com/tkngate/tkngate.git
cd tkngate
cp tkngate.example.yaml tkngate.yaml # edit with your keys
export TKNGATE_MASTER_KEY= " your-32-char-secret-key-here-!! "
go run main.go serve
Your agents now point to http://localhost:7477/openai/chat/completions instead of api.openai.com .
Route through OpenAI, Anthropic, DeepSeek, Kimi, and Groq from a single endpoint. If one provider goes down (HTTP 500/502/503), Tkngate automatically fails over to the next.
Block prompt injection attacks and automatically redact PII (credit cards, SSNs, API keys) before they reach the provider.
Stake-and-Slash Reputation (Mesh)
Protect donated API keys from abuse using cryptographic Fraud Proofs. Nodes that route malicious prompts bypassing the WAF are penalized via a Stake-and-Slash trust ledger and permanently blacklisted.
Generate secure tkngate-sk-... enterprise virtual keys for your teams. Each key acts as an isolated sandbox with hard budget caps, shielding your physical upstream keys.
Strict Rate Limiting (Token Bucket)
Protect your budget and providers from autonomous agent "burst loops" with an ultra-low latency, in-memory Token Bucket rate limiter.
Real-time spend tracking with Green → Amber → Red zones. Set global limits, per-session caps, and automatic request blocking when budgets are exhausted.
Distributed Semantic Cache (Redis)
Identical prompts are served from a distributed Redis cache, enabling horizontal scaling across multiple Tkngate proxy nodes. Save tokens and money globally across your fleet. Cache keys are computed from normalised model + messages hashes.
Automatically compresses Go, Python, and JavaScript code blocks in prompts — stripping comments and whitespace to reduce token usage by up to 40%.
The world's first BitTorrent-style token pool for LLM APIs. Donate spare API keys to the mesh, and get priority access to the network's capacity during outages. Protected by AES-256 zero-knowledge encryption.
Silently mirror a fraction of production traffic to an alternative provider (e.g., DeepSeek) to evaluate cost savings — with zero latency impact on the primary request.
See tkngate.example.yaml for the full reference.
Topic
Link
Budget System & Traffic Lights
docs/budgeting.md
DRR Token Mesh & Reputation
docs/drr-mesh.md
Enterprise Virtual Keys
docs/virtual-keys.md
Strict Rate Limiting
docs/rate-limiting.md
Zero-Knowledge Security
docs/zero-knowledge-security.md
Shadow Mode
docs/shadow-mode.md
Architecture
Agent Request
│
▼
┌──────────────────────────────────────┐
│ TKNGATE PROXY │
│ │
│ Budget Guard → AI-WAF/DLP │
│ → Context Compressor │
│ → Semantic Cache │
│ → Auto-Retry (3x) │
│ → Universal Router (Failover) │
│ → Shadow Mode (async mirror) │
│ → Token Counter → Ledger │
└──────────────────────────────────────┘
│
▼
OpenAI / Anthropic / DeepSeek / Kimi / Groq
License
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
