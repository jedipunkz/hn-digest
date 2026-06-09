---
source: "https://github.com/tamasPetki/HeadlessTracker"
hn_url: "https://news.ycombinator.com/item?id=48457500"
title: "Show HN: HeadlessTracker – MCP server that gives your AI eyes on your portfolio"
article_title: "GitHub - tamasPetki/HeadlessTracker: MCP server for crypto portfolio tracking. No dashboard UI — AI hosts render on demand. Bybit, Binance, MetaMask, Solana, Polymarket. · GitHub"
author: "bulltrapp"
captured_at: "2026-06-09T07:21:53Z"
capture_tool: "hn-digest"
hn_id: 48457500
score: 1
comments: 0
posted_at: "2026-06-09T06:53:58Z"
tags:
  - hacker-news
  - translated
---

# Show HN: HeadlessTracker – MCP server that gives your AI eyes on your portfolio

- HN: [48457500](https://news.ycombinator.com/item?id=48457500)
- Source: [github.com](https://github.com/tamasPetki/HeadlessTracker)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T06:53:58Z

## Translation

タイトル: Show HN: HeadlessTracker – AI がポートフォリオを監視する MCP サーバー
記事のタイトル: GitHub - tamasPetki/HeadlessTracker: 暗号ポートフォリオ追跡用の MCP サーバー。ダッシュボード UI なし - AI ホストがオンデマンドでレンダリングします。 Bybit、Binance、MetaMask、Solana、Polymarket。 · GitHub
説明: 暗号ポートフォリオ追跡用の MCP サーバー。ダッシュボード UI なし - AI ホストがオンデマンドでレンダリングします。 Bybit、Binance、MetaMask、Solana、Polymarket。 - タマスペトキ/HeadlessTracker

記事本文:
GitHub - tamasPetki/HeadlessTracker: 暗号ポートフォリオ追跡用の MCP サーバー。ダッシュボード UI なし - AI ホストがオンデマンドでレンダリングします。 Bybit、Binance、MetaMask、Solana、Polymarket。 · GitHub
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
タマスペトキ
/
ヘッドレストラッカー
公共
通知
あなたはしなければなりません

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
99 コミット 99 コミット .github .github bin bin dist/ mcp-apps dist/ mcp-apps docs docs scripts scripts src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md DISCLAIMER.md DISCLAIMER.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md ROADMAP.md ROADMAP.md bun.lock bun.lock daily-log.md daily-log.md Decisions.md Decisions.md Glama.json Glama.json package.json package.json server.json server.json smithery.yaml smithery.yaml tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
🤖 このプロジェクトは、AI 開発エージェントである Hex によって自律的に開発および保守されています。意思決定ログ: Decisions.md · 毎日のビルド ログ: daily-log.md · ソロ チーム。開発ループには人間がいません。
⚠️経済的なアドバイスではありません。 HeadlessTracker はポートフォリオ データ集約ツールです。情報提供のみを目的としています。全文については、免責事項.md を参照してください。
仮想通貨取引所、オンチェーンウォレット、予測市場にわたるポートフォリオ追跡のための MCP サーバー。 MCP 互換の AI ホスト (Claude Desktop、Claude Code、Cursor、ChatGPT) からポートフォリオをクエリします。
論文: AI ホスト (Claude Desktop、Claude Code、Cursor、ChatGPT) は構造化データからオンデマンドでダッシュボードを生成します。 2026 年にさらに別のトラッカー UI を構築することは無駄な作業です。データ層を構築します。 AI ホストをレンダラーにします。
ステータス: 実稼働準備が整っており、npm で稼働しています (上記のバージョン バッジ)。 5 つのコネクタ (Bybit、Binance、MetaMask/EVM、Solana、Polymarket)、15 の MCP ツール、インタラクティブなマルチタブ ダッシュボード パネル、ターミナル クエリ用の CLI、および 356 のテスト スイート。プレーン ノード ( npx headless-tracker ) または Bun で実行され、Clau とエンドツーエンドで動作します。

デスクトップ。
5 つのコネクタ: Bybit、Binance Spot+Futures、MetaMask マルチチェーン + マルチウォレット、Polymarket、Solana マルチウォレット
15 の MCP ツール: 6 データ + 7 アカウント/トークン管理 + 2 MCP アプリ パネル
3 つの MCP プロンプト: ポートフォリオ ダッシュボード、週次レビュー、リスク チェック
インタラクティブなダッシュボード MCP アプリ: ドーナツ + 棒グラフ、通貨スイッチャー、更新ボタンを備えた 3 つのタブ (ポートフォリオ / 週次 / リスク)
セットアップと管理用のライブ設定 MCP アプリ
CLI ポートフォリオ クエリ: 保有株 / 損益計算書 / 取引を表示 (クロードは必要ありません)
カスタム ERC-20 トークン リスト。 FIFO + トランザクション履歴の平均コスト
複数通貨表示 (USD/EUR/GBP/HUF); CoinGecko + Jupiter スポットと過去の価格
時間枠損益 ( --timeframe=24h|7d|30d|ytd )
356 テスト スイート ;プレーンなノードまたは Bun で実行されます
完了した内容、次の内容、および意図的に範囲外となっている内容については、ROADMAP.md を参照してください。
アカウント (読み取り専用) に接続し、すべてを単一のスキーマに正規化し、MCP ツールとして公開します。次に、クロード (または任意の MCP ホスト) に次のように尋ねます。
「私のポートフォリオは暗号市場と予測市場にどのように分割されていますか?」
「Polymarket のポジションをイベントごとにグループ化して表示します。」
「Bybit を更新して、BTC 損益を教えてください。」
AI ホストはチャート、表、内訳を生成します。 UI を構築するわけではありません。
60 秒以内にお試しください (API キーなし)
ゼロセットアップ バージョン - コマンド 1 つで、アカウントもキーもアドレスも必要ありません。 AI ホストが受け取ったとおりにレンダリングされた完全なサンプル ポートフォリオ (5 つの会場、仮想通貨 + 現金 + 予測市場) をご覧ください。
npx ヘッドレストラッカーのデモ
勘定科目記号 クラス 数量 値 価格
────────────────────────────────
bybit:UNIFIED BTC 暗号 0.420000 $25704 $61

200
バイナンス:スポット SOL 暗号 95.0000 $14440 $152.00
メタマスク:0xd8d2…f1a3 WBTC 暗号化 0.150000 $9150 $61000
solana:7vfC…Wd9k JUP 暗号 1800.00 $1656 $0.9200
Polymarket:0x9c1a…7b20 利下げ-2026 (YES) 予測 1500.00 $930.00 $0.6200
…
合計: $104126 (5 会場で 15 ポジション)
資産クラス別の配分:
暗号 $88024 84.5% ████████████████████
現金 $14900 14.3% ███
予測 $1202 1.2% █
また、クロードに尋ねるような簡単な英語の質問 (「全体で何を所有していますか?」、「どのように分割されていますか?」) を、それぞれに回答する MCP ツールにマッピングして出力します。独自の番号が必要な場合は、実際のアドレスまたは読み取り専用キーを使用した同じループになります。
それが良いかどうかを確認するためだけに、キーを交換するために新しいツールを渡す必要はありません。 Solana と Polymarket はパブリックのオンチェーン アドレスを読み取るため、認証情報なしで、目に見える任意のウォレット (自分のウォレットも含む) に HeadlessTracker をポイントできます。
# install (またはコマンドの先頭に「npx」を付けます)
npm install -g headless-tracker
# パブリック Solana ウォレットを追加します: API キーなし、アドレスのみ
ヘッドレストラッカーのセットアップ solana
# Solana アドレス (base58): <任意のパブリック アドレスを貼り付け>
# (オプションの RPC + ダスト プロンプトで Enter キーを押します)
# ターミナルで直接所蔵を印刷します。クロードは必要ありません
ヘッドレストラッカーショーホールディングス
勘定科目記号 クラス 数量 値 価格
─────── ────── ────── ─────── ─────
ソラナ:7Xk2…q9Fa SOL 暗号 12.4081 $2604 $209.88
ソラナ:7Xk2…q9Fa USDC 暗号 540.0000 $540.00 $1.00
ソラナ:7Xk2…q9Fa JUP 暗号 1200.00 $612.00 $0.5100
合計: $3756 (1 つのアカウントで 3 つのポジション)
(出力例。ここでは幅を考慮してアカウント ID が短縮されています。あなたの番号

rs はライブ チェーンから取得されます。)
これがループ全体です。インストールし、公開アドレスを指定し、正規化された所蔵を確認します。プライベートアカウント（Bybit、Binance）が必要な場合は、それらもセットアップしてください。すべてのコネクタは読み取り専用の資格情報を使用し、OS キーチェーンに保持され、ディスクに書き込まれたり、取引所独自の API 以外に送信されることはありません。それからそれをクロードに配線して、「私は何を所有していますか?」と尋ねます。チャットネイティブのダッシュボードと同じデータを取得します。
インタラクティブなダッシュボード (ライブ UI パネル)
MCP アプリ (Claude Desktop、ChatGPT、Goose、VS Code) をサポートするホストの場合は、次のように言います。
ホストは、チャット パネルにサンドボックス化された iframe を 3 つのライブ タブでレンダリングします。
ポートフォリオ — 合計値 KPI、上位順位テーブル、シンボルごとの割り当てドーナツ (上位 7 + 「その他」の末尾)、警告 + 失敗
毎週 — 7 日間のウィンドウデルタ KPI、最近の取引テーブル、スキップされたシンボルの開示 (理由付き)
リスク — 集中監査（単一ポジション、会場、ステーブルコイン準備金、予測市場オーバーウェイト）PASS / WARN / ALERTスコア、会場別ドーナツ
さらに、通貨スイッチャー (USD / EUR / GBP / HUF) と更新ボタンも付いています。 iframe は、ユーザーがタブをクリックすると独自のフォローアップ ツール呼び出しを行います。iframe を開いた後は、追加のプロンプトは必要ありません。オプションの引数:
HUF のダッシュボード、週次タブを開きます
実装: src/mcp/apps/dashboard/ (ブラウザ側 TS は bun run build:apps を介して単一の dist/mcp-apps/dashboard.html にバンドルされ、パッケージに同梱されます)。バンドルされたアーティファクトは npm パッケージ内に同梱されるため、npx headless-tracker を実行しているユーザーはビルド ステップを必要としません。
ホストが MCP アプリをまだレンダリングしていない場合でも、render_dashboard ツールはテキストによる確認を返します。以下のプロンプト クックブックをフォールバックとして使用します。同じワークフロー、同じデータで、ライブ UI パネルがないだけです。
設定パネル (セットアップ + 管理用のライブ UI)
失敗しないセットアップのために

ターミナル、尋ねてください：
設定 MCP アプリが開き、次の 4 つのタブが表示されます。
アカウント — [削除] ボタンが付いた構成済みアカウントのリスト (一方向の確認ダイアログ、OS キーチェーンとレジストリの両方から削除)。
アカウントの追加 — Bybit / Binance / MetaMask / Solana / Polymarket のフォーム。各フォームは、資格情報を永続化する前に、アップストリーム API に対して検証します。上部の明示的なセキュリティ開示: フォーム経由で送信された資格情報は、キーチェーンに送信される途中でクロード デスクトップのプロセスを通過します。 5 つのコネクタはすべて、設計上読み取り専用の資格情報を使用します (Bybit は「読み取り」のみ、引き出しなし、Binance は「読み取り可能」のみ、取引または引き出しなし、Etherscan はパブリック データのレート制限トークン、Polymarket プロキシ ウォレットはすでにパブリック、Solana アドレスはパブリックのオンチェーン ID です)。最悪の場合のリーク = ポートフォリオは読まれ、決して資金を移動しない。ゼロトラストの場合、CLI フロー (bun run setup <connector>) は引き続き利用可能です。
ウォレット — 追加のウォレット アドレスを既存の MetaMask または Solana アカウントに追加します (1 つの MCP アカウントの下にあるマルチウォレット、同じ Etherscan キー/チェーン選択または RPC URL を共有します)。
カスタム トークン — チェーンごとに ERC-20 トークンをリスト/追加/削除します。トークンデータはオンチェーン上で公開されます。キーチェーンの関与はありません。
どちらのパス (CLI または設定 UI) も同じ ~/.headless-tracker/cache.db + OS キーチェーンに書き込むため、どちらかを使用して作成されたアカウントはダッシュボードと CLI にすぐに表示されます。
クローンもビルドステップも Bun も必要ありません。パッケージはプレーン ノード (≥ 22.5) または Bun で実行されます。グローバルにインストールします。
npm install -g headless-tracker
または、インストールせずに npx というプレフィックスを付けてコマンドを実行します。 npx headless-tracker セットアップ solana 。 (開発用にソースからビルドするには Bun を使用します。「 開発 」を参照してください。)
2. アカウントを設定します (対話型)
必要な統合ごとにセットアップを実行します。それぞれ資格情報の入力を求められます。 v

それらを検証し、OS キーチェーン (macOS キーチェーン、Linux Secret Service、Windows Credential Vault) に保存します。キーチェーンのないヘッドレス ボックスについては、以下のヘッドレス / OS キーチェーンなしを参照してください。
ヘッドレストラッカーのセットアップ bybit
ヘッドレストラッカーのセットアップバイナンス
ヘッドレストラッカーセットアップメタマスク
ヘッドレストラッカーのセットアップ solana
ヘッドレストラッカーセットアップポリマーケット
構成内容を確認します。
ヘッドレストラッカーリストアカウント
ヘッドレス/OS キーチェーンなし (Docker、WSL、サーバー、CI)
OS キーチェーンには、実行中のシークレット サービス (Linux ではシークレット サービス/D-Bus、macOS ではキーチェーン、Windows では Credential Vault) が必要です。実際の環境の多くには、Docker コンテナ、WSL、ベア Linux サーバー、CI ジョブなど、これらがありません。そこで、キーチェーンの書き込みは失敗します。
その場合、セットアップは中止されません。引き続きアカウントを登録し、設定する正確な環境変数を出力します。次に例を示します。
⚠ OS キーチェーンが利用できないため、認証情報は保存されませんでした (...)。アカウントは
登録されました。 ... HEADLESS_TRACKER_SOLANA_<ADDR> 環境変数を
MCP サーバーの環境に JSON オブジェクトを追加して、再起動します。
その変数を MCP サーバーの env ブロック (またはシェル) 内のコネクタの資格情報 JSON に設定し、再起動します。 env var は常にキーチェーンよりも優先されるため、これは明示的なオーバーライドとしても機能します。コネクタごとの JSON 形状 (読み取り専用 API キーを使用します。「設定パネル」セクションのセキュリティに関する注意事項を参照してください):
変数名はアカウント識別子から派生します (セットアップでは正確な文字列が出力されるため、手動で作成する必要はありません)。 Claude デスクトップ/MCP 構成環境 blo の例

[切り捨てられた]

## Original Extract

MCP server for crypto portfolio tracking. No dashboard UI — AI hosts render on demand. Bybit, Binance, MetaMask, Solana, Polymarket. - tamasPetki/HeadlessTracker

GitHub - tamasPetki/HeadlessTracker: MCP server for crypto portfolio tracking. No dashboard UI — AI hosts render on demand. Bybit, Binance, MetaMask, Solana, Polymarket. · GitHub
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
tamasPetki
/
HeadlessTracker
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
99 Commits 99 Commits .github .github bin bin dist/ mcp-apps dist/ mcp-apps docs docs scripts scripts src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md DISCLAIMER.md DISCLAIMER.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md bun.lock bun.lock daily-log.md daily-log.md decisions.md decisions.md glama.json glama.json package.json package.json server.json server.json smithery.yaml smithery.yaml tsconfig.json tsconfig.json View all files Repository files navigation
🤖 This project is being developed and maintained autonomously by Hex , an AI dev agent. Decisions log: decisions.md · Daily build log: daily-log.md · Solo team. No human in the dev loop.
⚠️ Not financial advice. HeadlessTracker is a portfolio data aggregation tool. For informational purposes only. See DISCLAIMER.md for full text.
MCP server for portfolio tracking across crypto exchanges, on-chain wallets, and prediction markets. Query your portfolio from any MCP-compatible AI host (Claude Desktop, Claude Code, Cursor, ChatGPT).
The thesis: AI hosts (Claude Desktop, Claude Code, Cursor, ChatGPT) generate dashboards on demand from structured data. Building yet another tracker UI is wasted work in 2026. Build the data layer; let the AI host be the renderer.
Status: Production-ready and live on npm (version badge above). Five connectors (Bybit, Binance, MetaMask/EVM, Solana, Polymarket), 15 MCP tools, an interactive multi-tab dashboard panel, a CLI for terminal queries, and a 356-test suite. Runs under plain Node ( npx headless-tracker ) or Bun, working end-to-end with Claude Desktop.
5 connectors: Bybit, Binance Spot+Futures, MetaMask multi-chain + multi-wallet, Polymarket, Solana multi-wallet
15 MCP tools: 6 data + 7 account/token management + 2 MCP App panels
3 MCP prompts: portfolio-dashboard , weekly-review , risk-check
Interactive dashboard MCP App: 3 tabs (Portfolio / Weekly / Risk) with donut + bar charts, currency switcher, refresh button
Live Settings MCP App for setup and admin
CLI portfolio queries: show holdings / pnl / transactions (no Claude required)
Custom ERC-20 token lists ; FIFO + Average Cost on transaction history
Multi-currency display (USD/EUR/GBP/HUF); CoinGecko + Jupiter spot and historical prices
Time-windowed PnL ( --timeframe=24h|7d|30d|ytd )
356-test suite ; runs under plain Node or Bun
See ROADMAP.md for what's done, what's next, and what's intentionally out of scope.
Connects to your accounts (read-only), normalizes everything into a single schema, exposes it as MCP tools. Then you ask Claude (or any MCP host):
"How is my portfolio split between crypto and prediction markets?"
"Show my Polymarket positions grouped by event."
"Refresh Bybit and tell me my BTC P&L."
The AI host generates the chart, the table, the breakdown. You don't build a UI.
Try it in 60 seconds (no API keys)
The zero-setup version — one command, no accounts, no keys, not even an address. See a full sample portfolio (five venues; crypto + cash + prediction markets) rendered exactly as your AI host receives it:
npx headless-tracker demo
account symbol class qty value price
────────────────────── ─────────────────── ────────── ──────── ─────── ───────
bybit:UNIFIED BTC crypto 0.420000 $25704 $61200
binance:spot SOL crypto 95.0000 $14440 $152.00
metamask:0xd8d2…f1a3 WBTC crypto 0.150000 $9150 $61000
solana:7vfC…Wd9k JUP crypto 1800.00 $1656 $0.9200
polymarket:0x9c1a…7b20 RATE-CUT-2026 (YES) prediction 1500.00 $930.00 $0.6200
…
Total: $104126 (15 positions across 5 venues)
Allocation by asset class:
crypto $88024 84.5% ████████████████████
cash $14900 14.3% ███
prediction $1202 1.2% █
It also prints the plain-English questions you'd ask Claude ("what do I own across everything?", "how is it split?") mapped to the MCP tool that answers each. When you want your own numbers, it's the same loop with a real address or read-only key:
You shouldn't have to hand a new tool your exchange keys just to find out whether it's any good. Solana and Polymarket read public on-chain addresses , so you can point HeadlessTracker at any wallet you can see (your own included) with zero credentials.
# install (or prefix any command with `npx`)
npm install -g headless-tracker
# add a public Solana wallet: no API key, just the address
headless-tracker setup solana
# Solana address (base58): <paste any public address>
# (press ENTER through the optional RPC + dust prompts)
# print the holdings right in your terminal, no Claude required
headless-tracker show holdings
account symbol class qty value price
───────────────── ────── ────── ──────── ─────── ────────
solana:7Xk2…q9Fa SOL crypto 12.4081 $2604 $209.88
solana:7Xk2…q9Fa USDC crypto 540.0000 $540.00 $1.00
solana:7Xk2…q9Fa JUP crypto 1200.00 $612.00 $0.5100
Total: $3756 (3 positions across 1 accounts)
(Example output; the account id is shortened here for width. Your numbers come from the live chain.)
That is the whole loop: install, point at a public address, see normalized holdings. When you want your private accounts (Bybit, Binance), setup those too. Every connector uses read-only credentials, kept in your OS keychain, never written to disk and never sent anywhere except the exchange's own API. Then wire it into Claude and ask "what do I own?" to get the same data as a chat-native dashboard.
Interactive dashboard (live UI panel)
For hosts that support MCP Apps — Claude Desktop, ChatGPT, Goose, VS Code — say:
The host renders a sandboxed iframe in the chat panel with three live tabs:
Portfolio — total value KPIs, top positions table, allocation-by-symbol donut (top 7 + "Other" tail), warnings + failures
Weekly — 7-day window delta KPIs, recent trades table, skipped-symbols disclosure (with reasons)
Risk — concentration audit (single-position, venue, stablecoin reserve, prediction-market overweight) scored PASS / WARN / ALERT, by-venue donut
Plus a currency switcher (USD / EUR / GBP / HUF) and a refresh button. The iframe makes its own follow-up tool calls as the user clicks tabs — no extra prompting needed once it's open. Optional args:
Open the dashboard in HUF, weekly tab
Implementation: src/mcp/apps/dashboard/ (browser-side TS bundled into a single dist/mcp-apps/dashboard.html via bun run build:apps , ships with the package). The bundled artifact ships inside the npm package so users running npx headless-tracker don't need a build step.
If your host doesn't render MCP Apps yet, the render_dashboard tool still returns a textual confirmation. Use the prompt cookbook below as a fallback — same workflows, same data, just no live UI panel.
Settings panel (live UI for setup + admin)
For setup that doesn't drop you into a terminal, ask:
The Settings MCP App opens with four tabs:
Accounts — list of configured accounts with a Remove button (one-way confirm dialog; deletes from both the OS keychain and the registry).
Add Account — forms for Bybit / Binance / MetaMask / Solana / Polymarket. Each form validates against the upstream API before persisting credentials. Explicit security disclosure at the top: credentials submitted via the form transit Claude Desktop's process en route to the keychain. All five connectors use READ-ONLY credentials by design (Bybit "Read" only, no Withdraw; Binance "Enable Reading" only, no Trade or Withdraw; Etherscan is a public-data rate-limit token; Polymarket proxy wallet is already public; Solana addresses are public on-chain identifiers). Worst-case leak = portfolio-read, never fund movement. For zero-trust, the CLI flow ( bun run setup <connector> ) stays available.
Wallets — add an additional wallet address to an existing MetaMask OR Solana account (multi-wallet under one MCP account, sharing the same Etherscan key/chain selection or RPC URL).
Custom Tokens — list / add / remove ERC-20 tokens per chain. Token data is public on-chain; no keychain involvement.
Either path (CLI or Settings UI) writes to the same ~/.headless-tracker/cache.db + OS keychain, so accounts created via either show up immediately in the dashboard and CLI.
No clone, no build step, no Bun required. The package runs under plain Node (≥ 22.5) or Bun. Install it globally:
npm install -g headless-tracker
Or run any command without installing by prefixing npx , e.g. npx headless-tracker setup solana . (Building from source for development uses Bun — see Development .)
2. Configure your accounts (interactive)
Run setup for each integration you want. Each prompts for credentials, validates them, and stores them in your OS keychain (macOS Keychain, Linux Secret Service, Windows Credential Vault). On a headless box with no keychain, see Headless / no OS keychain below.
headless-tracker setup bybit
headless-tracker setup binance
headless-tracker setup metamask
headless-tracker setup solana
headless-tracker setup polymarket
Verify what's configured:
headless-tracker list-accounts
Headless / no OS keychain (Docker, WSL, servers, CI)
The OS keychain needs a running secret service (Secret Service / D-Bus on Linux, Keychain on macOS, Credential Vault on Windows). Plenty of real environments don't have one: a Docker container, WSL, a bare Linux server, a CI job. There, the keychain write fails.
In that case setup does not abort. It still registers the account, then prints the exact environment variable to set, for example:
⚠ OS keychain unavailable, so credentials were NOT stored (...). The account is
registered. ... set the HEADLESS_TRACKER_SOLANA_<ADDR> environment variable to a
JSON object in your MCP server's env, then restart.
Set that variable to the connector's credential JSON in your MCP server's env block (or your shell), then restart. The env var always takes precedence over the keychain, so this also works as an explicit override. Per-connector JSON shapes (use read-only API keys — see the security note in the Settings panel section):
The variable name is derived from the account identifier ( setup prints the exact string, so you don't have to construct it by hand). Example for a Claude Desktop / MCP config env blo

[truncated]
