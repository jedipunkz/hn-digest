---
source: "https://github.com/Evozim/m2mcent-sdk"
hn_url: "https://news.ycombinator.com/item?id=48393601"
title: "Show HN: X402-express. Pay-per-call HTTP 402 middleware for AI agents on Base L2"
article_title: "GitHub - Evozim/m2mcent-sdk · GitHub"
author: "m2mcent"
captured_at: "2026-06-04T04:34:40Z"
capture_tool: "hn-digest"
hn_id: 48393601
score: 1
comments: 0
posted_at: "2026-06-04T03:57:29Z"
tags:
  - hacker-news
  - translated
---

# Show HN: X402-express. Pay-per-call HTTP 402 middleware for AI agents on Base L2

- HN: [48393601](https://news.ycombinator.com/item?id=48393601)
- Source: [github.com](https://github.com/Evozim/m2mcent-sdk)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T03:57:29Z

## Translation

タイトル: 表示 HN: X402-express。 Pay-per-call HTTP 402 middleware for AI agents on Base L2
Article title: GitHub - Evozim/m2mcent-sdk · GitHub
Description: Contribute to Evozim/m2mcent-sdk development by creating an account on GitHub.

記事本文:
GitHub - Evozim/m2mcent-sdk · GitHub
Skip to content
Navigation Menu
サインイン
Appearance settings
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
Search or jump to...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
Search syntax tips
Provide feedback
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
Query
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
Appearance settings
Resetting focus
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
Dismiss alert
{{ message }}
Evozim
/
m2mcent-sdk
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
m2mCent ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーと

ファイル
5 コミット 5 コミット src src .gitignore .gitignore .npmignore .npmignore ライセンス ライセンス README.md README.md cdp-agentkit-integration.md cdp-agentkit-integration.md mcp-publisher.exe mcp-publisher.exe package-lock.json package-lock.jsonパッケージ.json パッケージ.json サーバー.json サーバー.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
🛰️ M2MCent SDK — AI エージェント用の x402 支払いラッパー
AI エージェントおよび MCP サーバー用の超軽量 x402 支払いインターセプター。 Monetize any API in 3 lines of code.
npm インストール m2mcent-sdk
import { X402Handler } from 'm2mcent-sdk' ;
const x402 = new X402Handler ( {
rpcUrl : 'https://mainnet.base.org' ,
プライベートキー: プロセス。環境 。 RELAYER_PRIVATE_KEY ! 、
受信者: プロセス。環境 。 TREASURY_ADDRESS !
} ) ;
// Express エンドポイントを USDC ペイウォールで保護する
アプリ。 post ( '/api/analyze' , x402 . requirePayment ( "100000" ) , ( req , res ) => {
レス。 json ( { result : "Premium analysis complete" , receipt : req . paymentTx } ) ;
} ) ;
📐 仕組み
M2MCent は、HTTP 402 からインスピレーションを得たマシンネイティブの支払い標準である x402 支払いプロトコルを実装しています。
Agent requests a protected endpoint
サーバーは 402 Payment Required + 支払いメタデータ (Base64 エンコード) で応答します。
エージェントは EIP-712 型データ認証に署名します (支払者にはガスレス)
サーバーはM2MCentエスクロー契約を介してオンチェーンでアトミックに決済します
Agent receives the premium response + transaction receipt
エージェント ──► API サーバー ──► 402 + メタデータ
Agent ◄── signs EIP-712 authorization
エージェント ──► API サーバー + Payment-Signature ヘッダー
└──► Escrow.settle() on Base L2
Agent ◄── Premium Response + tx hash
⚙️ 構成
パラメータ
説明
必須
rpcURL
Base Mainnet RPC endpoint
✅
秘密鍵
リレイヤーウォレットプライベートk

ey（決済用）
✅
受信者
支払いを受け取るための財務省の住所
✅
🔒 セキュリティ
非保管: 資金はオンチェーンエスクローを介して支払人→財務省から直接流れます。
支払者向けのガスレス: EIP-3009 transferWithAuthorization (USDC ネイティブ) を使用
ゼロリーク : 決済後にエージェントデータは保持されません
アトミック : 支払いとサービスの提供は単一のリクエスト サイクルで行われます。
パラメータ
値
ネットワーク
ベースメインネット (チェーン ID: 8453)
USDC契約
0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913
エスクロー契約
0xf3c3416A843d13C944554A54Ac274BB7fF264BcC
決済
アトミック、1 秒未満のファイナリティ
🧩 MCPサーバーの統合
モデル コンテキスト プロトコル サーバーに最適:
'express' からエクスプレスをインポートします。
'm2mcent-sdk' から { X402Handler } をインポートします。
const app = Express() ;
const x402 = 新しい X402Handler ( {
rpcUrl: プロセス。環境 。 BASE_RPC_URL ! 、
プライベートキー: プロセス。環境 。 RELAYER_PRIVATE_KEY ! 、
受信者: プロセス。環境 。 TREASURY_ADDRESS !
} ) ;
// MCP ツールのエンドポイントが収益化可能になります
アプリ。 post ( '/api/tools/analyze' , x402 . requirePayment ( "50000" ) , async ( req , res ) => {
const result = await runMCPTool(req.body);
レス。 json ({ ... 結果 ,paymentTx : req .paymentTx }) ;
} ) ;
📊 エコシステム
M2MCent は、Base Mainnet 上の 100 以上の実稼働 MCP サーバーを稼働させ、AI 間コマースのための実際の USDC マイクロペイメントを処理します。
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

Contribute to Evozim/m2mcent-sdk development by creating an account on GitHub.

GitHub - Evozim/m2mcent-sdk · GitHub
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
Evozim
/
m2mcent-sdk
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
m2mCent Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits src src .gitignore .gitignore .npmignore .npmignore LICENSE LICENSE README.md README.md cdp-agentkit-integration.md cdp-agentkit-integration.md mcp-publisher.exe mcp-publisher.exe package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json View all files Repository files navigation
🛰️ M2MCent SDK — x402 Payment Wrapper for AI Agents
Ultra-lightweight x402 payment interceptor for AI Agents and MCP Servers. Monetize any API in 3 lines of code.
npm install m2mcent-sdk
import { X402Handler } from 'm2mcent-sdk' ;
const x402 = new X402Handler ( {
rpcUrl : 'https://mainnet.base.org' ,
privateKey : process . env . RELAYER_PRIVATE_KEY ! ,
recipient : process . env . TREASURY_ADDRESS !
} ) ;
// Protect any Express endpoint with a USDC paywall
app . post ( '/api/analyze' , x402 . requirePayment ( "100000" ) , ( req , res ) => {
res . json ( { result : "Premium analysis complete" , receipt : req . paymentTx } ) ;
} ) ;
📐 How It Works
M2MCent implements the x402 Payment Protocol — a machine-native payment standard inspired by HTTP 402:
Agent requests a protected endpoint
Server responds with 402 Payment Required + payment metadata (Base64 encoded)
Agent signs an EIP-712 typed data authorization (gasless for the payer)
Server settles atomically on-chain via the M2MCent Escrow contract
Agent receives the premium response + transaction receipt
Agent ──► API Server ──► 402 + metadata
Agent ◄── signs EIP-712 authorization
Agent ──► API Server + Payment-Signature header
└──► Escrow.settle() on Base L2
Agent ◄── Premium Response + tx hash
⚙️ Configuration
Parameter
Description
Required
rpcUrl
Base Mainnet RPC endpoint
✅
privateKey
Relayer wallet private key (for settlement)
✅
recipient
Treasury address to receive payments
✅
🔒 Security
Non-custodial : Funds flow directly from payer → treasury via on-chain escrow
Gasless for payers : Uses EIP-3009 transferWithAuthorization (USDC native)
Zero-Leak : No agent data is retained after settlement
Atomic : Payment and service delivery happen in a single request cycle
Parameter
Value
Network
Base Mainnet (Chain ID: 8453)
USDC Contract
0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913
Escrow Contract
0xf3c3416A843d13C944554A54Ac274BB7fF264BcC
Settlement
Atomic, sub-second finality
🧩 MCP Server Integration
Perfect for Model Context Protocol servers:
import express from 'express' ;
import { X402Handler } from 'm2mcent-sdk' ;
const app = express ( ) ;
const x402 = new X402Handler ( {
rpcUrl : process . env . BASE_RPC_URL ! ,
privateKey : process . env . RELAYER_PRIVATE_KEY ! ,
recipient : process . env . TREASURY_ADDRESS !
} ) ;
// Any MCP tool endpoint becomes monetizable
app . post ( '/api/tools/analyze' , x402 . requirePayment ( "50000" ) , async ( req , res ) => {
const result = await runMCPTool ( req . body ) ;
res . json ( { ... result , paymentTx : req . paymentTx } ) ;
} ) ;
📊 Ecosystem
M2MCent powers 100+ production MCP servers on Base Mainnet, processing real USDC micro-payments for AI-to-AI commerce.
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
