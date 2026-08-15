---
source: "https://github.com/Squid-Pay/Squid-Agent-Wallet-SDK"
hn_url: "https://news.ycombinator.com/item?id=49310808"
title: "Open-source wallet SDK for AI agents"
article_title: "GitHub - Squid-Pay/Squid-Agent-Wallet-SDK · GitHub"
author: "horatiucode"
captured_at: "2026-08-15T15:11:00Z"
capture_tool: "hn-digest"
hn_id: 49310808
score: 2
comments: 0
posted_at: "2026-08-15T14:19:20Z"
tags:
  - hacker-news
  - translated
---

# Open-source wallet SDK for AI agents

- HN: [49310808](https://news.ycombinator.com/item?id=49310808)
- Source: [github.com](https://github.com/Squid-Pay/Squid-Agent-Wallet-SDK)
- Score: 2
- Comments: 0
- Posted: 2026-08-15T14:19:20Z

## Translation

タイトル: AI エージェント用のオープンソース ウォレット SDK
記事タイトル: GitHub - Squid-Pay/Squid-Agent-Wallet-SDK · GitHub
説明: GitHub でアカウントを作成して、Squid-Pay/Squid-Agent-Wallet-SDK の開発に貢献します。

記事本文:
GitHub - Squid-Pay/Squid-Agent-Wallet-SDK · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
イカペイ
/
Squid-Agent-Wallet-SDK
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット dist dist 例 例 src src testing testing .DS_Store .DS_Store .gitattributes .gitattributes AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md パッケージ

e-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Squid Pay に接続する AI エージェントおよびバックエンド用の TypeScript SDK。
エージェントは、許可されたアカウントのコンテキストを読み取り、支払いや取引を提案できます。
保留を承認したり、取引に署名したり、ブロードキャストしたり、資金を移動したりすることはできません。所有者の接続されたウォレットが最終署名者として残ります。
エージェント(このSDK) → Squid検証/HOLD → 人間によるレビュー → オーナーウォレット署名 → Squid検証
インストール
cd /Users/horatiubudai/ceo/Squid/sdk_developer
npmインストール
npm ビルドを実行する
別のプロジェクト (ローカル パス) から:
npm install /Users/horatiubudai/ceo/Squid/sdk_developer
クイックスタート
Squid → 設定 → CLI ( sq_master_… ) でプラットフォーム API キーを作成します。
または、エージェントからのエージェント キー ( sq_live_… ) を使用します。
import { SquidAgentWallet } から "@squid/agent-wallet-sdk" ;
const イカ = 新しい SquidAgentWallet ( {
エンドポイント: プロセス。環境 。 SQUID_ENDPOINT || "http://localhost:4173" ,
apiKey: プロセス。環境 。 SQUID_API_KEY
} ) ;
const status = イカを待ちます。 getStatus() ;
const ウォレット = イカを待ちます。 listWallets() ;
const提案=squidを待ちます。支払いを提案する ( {
アセット : "USDC" 、
数量：25、
受信者 : "0xRecipient..." 、
チェーン: "ベース" 、
注: "ベンダー請求書 #1042" 、
冪等性キー : "invoice-1042"
} ) ;
コンソール。ログ (提案のステータス) ; // "保持" | 「財布の準備完了」
コンソール。ログ(提案 .提案Id) ; // レビュー保留 ID が必要です
コンソール。ログ (提案のメッセージ) ; // 人間が読める次のステップ
輸送
キー
デフォルトのトランスポート
エンドポイント
sq_master_… (プラットフォーム API キー)
CLI HTTP
POST /api/cli/execute
sq_live_… (エージェントキー)
MCP
POST /mcp
トランスポートを強制的に実行します。
new SquidAgentWallet ({ apiKey , Transport : "mcp" } ) ;
新しい SquidAgentWallet ( { apiKey , トランスポート : "cli" } )

;
カーソル / クロード / エルメスの MCP 設定
コンソール。 log(JSON.stringify(squid.toMcpServerConfig(),null,2));
{
"mcpサーバー": {
「イカ」: {
"url" : " http://localhost:4173/mcp " ,
"ヘッダー" : {
「認可」：「ベアラー sq_live_…」
}
}
}
}
この SDK によって公開される Squid MCP ツール:
イカ 。 getStatus()
イカ 。 listWallets ( ) / getBalances ( )
イカ 。 listHolds ( ) / getHold ( id )
イカ 。リストエージェント ( )
イカ 。リストルール ( )
イカ 。 listPaymentLinks ( ) / getPaymentLink ( id )
イカ 。 getPlatformState ()
イカ 。 getFinancialHarness ()
イカ 。医者（）
イカ 。 proposalAction({type,domain?,asset?,amount?,recipient?,chain?,note?,idempotencyKey?})
イカ 。 proposalPayment ({ 資産 , 金額 , 受取人 , チェーン ? , メモ ? , 冪等性Key ? } )
イカ 。 proposalTrade ( { 資産 , 金額 , … } )
イカ 。 executeCli ( [ "ステータス" ] )
イカ 。 callMcpTool ( "list_holds" )
イカ 。 listMcpTools ( )
イカ 。 toMcpServerConfig ()
意図的に利用できない ( SquidForbiddenError をスローする):
Squid またはこの SDK には秘密キー/シード フレーズはありません。
提案は Squid Brain / ポリシーを通過します。危険なものはレビュー保留になります。
お金は所有者のウォレットに署名し、Squid がチェーンレシートを確認した後にのみ移動します。
Squid からの構造化エラー (コード、タイプ、メッセージ) は SquidSdkError として保存されます。
エクスポート SQUID_ENDPOINT=http://localhost:4173
エクスポート SQUID_API_KEY=sq_master_…
ノードの例/quickstart.mjs
ノードの例/propose-payment.mjs
ノードの例/mcp-config.mjs
イカペイとの関係
このパッケージは、bank_squid/Squid Pay / のコンソールの隣にあります。
そこで npm run dev を使用して Squid を実行し (デフォルトは http://localhost:4173 )、そのオリジンで SDK を指定します。
npmインストール
npm ビルドを実行する
npmテスト
Node.js ≥ 18 (ネイティブフェッチ) が必要です。
MIT ライセンスの Readme アクティビティ C

カスタムプロパティスターズ
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to Squid-Pay/Squid-Agent-Wallet-SDK development by creating an account on GitHub.

GitHub - Squid-Pay/Squid-Agent-Wallet-SDK · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Squid-Pay
/
Squid-Agent-Wallet-SDK
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits dist dist examples examples src src tests tests .DS_Store .DS_Store .gitattributes .gitattributes AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
TypeScript SDK for AI agents and backends that connect to Squid Pay .
Agents can read permitted account context and propose payments or trades.
They cannot approve holds, sign transactions, broadcast, or move money. The owner’s connected wallet remains the final signer.
Agent (this SDK) → Squid validates / HOLDs → Human reviews → Owner wallet signs → Squid verifies
Install
cd /Users/horatiubudai/ceo/Squid/sdk_developer
npm install
npm run build
From another project (local path):
npm install /Users/horatiubudai/ceo/Squid/sdk_developer
Quick start
Create a Platform API key in Squid → Settings → CLI ( sq_master_… ),
or use an agent key from Agents ( sq_live_… ).
import { SquidAgentWallet } from "@squid/agent-wallet-sdk" ;
const squid = new SquidAgentWallet ( {
endpoint : process . env . SQUID_ENDPOINT || "http://localhost:4173" ,
apiKey : process . env . SQUID_API_KEY
} ) ;
const status = await squid . getStatus ( ) ;
const wallets = await squid . listWallets ( ) ;
const proposal = await squid . proposePayment ( {
asset : "USDC" ,
amount : 25 ,
recipient : "0xRecipient..." ,
chain : "Base" ,
note : "Vendor invoice #1042" ,
idempotencyKey : "invoice-1042"
} ) ;
console . log ( proposal . status ) ; // "held" | "ready_for_wallet"
console . log ( proposal . proposalId ) ; // Needs Review hold id
console . log ( proposal . message ) ; // Human-readable next step
Transports
Key
Default transport
Endpoint
sq_master_… (Platform API key)
CLI HTTP
POST /api/cli/execute
sq_live_… (agent key)
MCP
POST /mcp
Force a transport:
new SquidAgentWallet ( { apiKey , transport : "mcp" } ) ;
new SquidAgentWallet ( { apiKey , transport : "cli" } ) ;
MCP config for Cursor / Claude / Hermes
console . log ( JSON . stringify ( squid . toMcpServerConfig ( ) , null , 2 ) ) ;
{
"mcpServers" : {
"squid" : {
"url" : " http://localhost:4173/mcp " ,
"headers" : {
"Authorization" : " Bearer sq_live_… "
}
}
}
}
Squid MCP tools exposed by this SDK:
squid . getStatus ( )
squid . listWallets ( ) / getBalances ( )
squid . listHolds ( ) / getHold ( id )
squid . listAgents ( )
squid . listRules ( )
squid . listPaymentLinks ( ) / getPaymentLink ( id )
squid . getPlatformState ( )
squid . getFinancialHarness ( )
squid . doctor ( )
squid . proposeAction ( { type , domain ? , asset ? , amount ? , recipient ? , chain ? , note ? , idempotencyKey ? } )
squid . proposePayment ( { asset , amount , recipient , chain ? , note ? , idempotencyKey ? } )
squid . proposeTrade ( { asset , amount , … } )
squid . executeCli ( [ "status" ] )
squid . callMcpTool ( "list_holds" )
squid . listMcpTools ( )
squid . toMcpServerConfig ( )
Deliberately not available (throws SquidForbiddenError ):
No private keys / seed phrases in Squid or this SDK.
Proposals go through Squid Brain / policy; risky ones become Needs Review holds.
Money moves only after the owner wallet signs and Squid verifies the chain receipt.
Structured errors from Squid ( code , type , message ) are preserved as SquidSdkError .
export SQUID_ENDPOINT=http://localhost:4173
export SQUID_API_KEY=sq_master_…
node examples/quickstart.mjs
node examples/propose-payment.mjs
node examples/mcp-config.mjs
Relation to Squid Pay
This package lives next to the console at bank_squid/Squid Pay / .
Run Squid with npm run dev there (default http://localhost:4173 ), then point the SDK at that origin.
npm install
npm run build
npm test
Requires Node.js ≥ 18 (native fetch ).
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
