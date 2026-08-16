---
source: "https://github.com/aleks-ent/pump-dump-crypto-screener"
hn_url: "https://news.ycombinator.com/item?id=49323256"
title: "An open-source, candle-only crypto pump detector–looking for ML contributors"
article_title: "GitHub - aleks-ent/pump-dump-crypto-screener: Pump and dump crypto screener with Telegram alerts · GitHub"
author: "alekshenderson"
captured_at: "2026-08-16T21:12:06Z"
capture_tool: "hn-digest"
hn_id: 49323256
score: 2
comments: 0
posted_at: "2026-08-16T20:16:30Z"
tags:
  - hacker-news
  - translated
---

# An open-source, candle-only crypto pump detector–looking for ML contributors

- HN: [49323256](https://news.ycombinator.com/item?id=49323256)
- Source: [github.com](https://github.com/aleks-ent/pump-dump-crypto-screener)
- Score: 2
- Comments: 0
- Posted: 2026-08-16T20:16:30Z

## Translation

タイトル: オープンソースのキャンドル専用クリプト ポンプ検出器 – ML 貢献者を募集しています
記事のタイトル: GitHub - aleks-ent/pump-dump-crypto-screener: Telegram アラートを使用したポンプおよびダンプ暗号スクリーナー · GitHub
説明: Telegram アラートを備えたポンプおよびダンプ暗号スクリーナー - aleks-ent/pump-dump-crypto-screener

記事本文:
GitHub - aleks-ent/pump-dump-crypto-screener: Telegram アラートを使用したポンプおよびダンプ暗号スクリーナー · GitHub
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
アレックス・エント
/
ポンプダンプ暗号スクリーナー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
69 コミット 69 コミット .cursor/ rules .cursor/ rules apps apps config config docs docs package package scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE PUMP_DETECT

ION_RULES.md PUMP_DETECTION_RULES.md README.md README.md config.example.js config.example.js エコシステム.config.cjs エコシステム.config.cjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml stop.sh stop.sh tsconfig.base.json tsconfig.base.json update.sh update.sh vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Binance および Bybit から 1m/5m 市場統計とバルク kline アーカイブを取得し、ポンプ/ダンプ レジームを検出し、Telegram 経由でアラートを送信するための TypeScript pnpm モノレポ。
ライブボットを試してみましょう: Telegram で @pumpdumpscreenerautobot を開きます
要件: Node.js 20 以降、pnpm 9 以降、pm2 (実稼働用のプロセス マネージャー)。
API キーの交換はありません。 Binance と Bybit のデータは、パブリック一括アーカイブとパブリック REST エンドポイントから取得されます。
2 つの無料アカウントが必要です (それぞれ約 2 分)。
1. 依存関係をインストールして構成する
pnpmインストール
cp config.example.js config.js # Turso + Telegram を入力します (上の表を参照)
config.js を編集します。少なくとも、 database 、 telegramBotToken 、およびプライベート classifierTelegramChatId を構成します。 /start を送信する人は誰でも自動的にアラートに登録され、 /stats 、 /runs 、および /about を使用できます。構成されたチャットのみがアラートを分類できます。ポンプ ルックバックとスキャン設定はポンプの下に存在します (デフォルトは config.example.js )。
データベース: {
URL : "libsql://screener-....turso.io" ,
authToken : "..." 、
} 、
telegramBotToken : "123456789:ABC..." 、
classifierTelegramChatId : "36772199" 、
ウェブ: {
port : 3000 , // ローカルアプリサーバー; nginx がパブリック HTTP/HTTPS を終了する
ホスト: "127.0.0.1" 、
} 、
ポンプ : {
days : 5 , // ダウンロードとスキャンのカレンダー日数をルックバックします
最小スコア : 80 、
scanCache : true 、
requireCalmPrePump : false , // 機能フラグ: ポンプの前に穏やかな 2 時間の期間が必要です
} 、
2. データベースを構築して初期化する
pnpmビルド
プンプ

m db:ブートストラップ
コードを変更するたびに pnpm build を再度実行します。PM2 は自動的にリビルドしません。
本番環境の実行は、リポジトリ ルートの Ecosystem.config.cjs で定義されます。次の 3 つのプロセスが開始されます。
pm2 起動エコシステム.config.cjs
PM2 はすべてのプロセスを存続させます。ポンプモニターはパイプラインの実行を終了すると終了します。 PM2 はすぐに次の実行を開始します ( autorestart: true )。これは手動の cron ループを置き換えます。
最初の実行では、pump.day 分のキャンドルが data/market_stats/ にダウンロードされます — ネットワーク + ディスクが必要です。何時間もかかることもあります。市場データはリポジトリにありません。
Web ページは、デフォルトでは 127.0.0.1:3000 の Pump-Web によって提供されます。パブリック HTTP/HTTPS の場合は、nginx をその前に置きます。 docs/nginx_letsencrypt.md を参照してください。
pm2保存
pm2 起動 # 出力されたコマンドに従います (サーバーごとに 1 回)
ポンプ イベント レビュアー
/review の手動レビュー担当者は、検出されたポンプを人間がラベル付けしたデータセットに変換します。
スクリーナーの評価と改善。高速かつキーボードフレンドリーな操作性を実現するように設計されています。
ワークフロー:
保存されているポンプ イベントをステータス、カテゴリ、交換、シンボル、日付ごとに参照およびフィルタリングします。
Telegram 購読者の投票を、
スクリーナーの検出時間。チャートは、利用可能な場合はローカル履歴を使用し、それ以外の場合はロードします
イベントのパブリック Binance または Bybit API からの 4 時間のウィンドウ。
各イベントをウィックスパイク、弱いポンプ、持続的な動き、ボリュームのみ、
非液体ノイズ、または不明瞭。オプションで信頼度とコメントを記録します。
キーボード ショートカットを使用して保存して先に進み、既存のラベルに再度アクセスし、レビューを追跡します
進行状況を確認し、ラベル付きデータセットを JSON または CSV としてエクスポートします。
人間による注釈は元のポンプ記録とは別に保存されるため、レビューは
検出器の出力を変更しないでください。ワークスペースは既存の Turso データベースを使用します。
ブラウザからリクエストを交換し、単純な HT で保護できます。

TP 基本認証。
セットアップについては、ポンプレビューオペレーターガイドを参照してください。
アクセス制御、展開、リリースのチェック。
PM2ステータス
pm2 ログ # すべてのアプリ
pm2 ログ ポンプ モニター # ダウンロード + スキャン パイプライン
pm2 ログポンプボット # テレグラムボット
pm2 ログ ポンプウェブ # HTTP ページ
pm2 restart ecosystem.config.cjs # config.js またはコード変更後 (最初に再構築)
pm2 restart Pump-monitor # パイプラインのみを再起動します
pm2 restart Pump-bot # ボットのみを再起動します
pm2 restart Pump-web # HTTP ページのみを再起動します
pm2 stop エコシステム.config.cjs
pm2 エコシステム.config.cjs を削除します
config.js を編集した後、影響を受けるプロセスを再起動します ( pm2 restart Pump-monitor 、 Pump-bot 、または Pump-Web )。再起動した場合、構成のみの変更には PM2 のリロードは必要ありません。
分類ボタンはポンプボットを実行する必要があります。これは Ecosystem.config.cjs に含まれており、運用環境ではオプションではありません。
PM2 駆動の各実行はエンドツーエンドのパイプラインです。
ダウンロード — Binance と Bybit からの 1m/5m ローソク足の最後の Pump.days (アーカイブ + REST フォールバック)。 docs/fetch_all.md を参照してください。
スキャン - ポンプ/ダンプ検出。 data/market_stats/reports/pump_events.ndjson を出力します。
持続 + アラート — Turso へのポンプ エピソードをアップサートします。新しい現在のポンプごとのテレグラム メッセージ ( Coin|pump_start_utc )。前回の正常な監視サイクルが開始される前に終了したエピソードは履歴バックフィルです。これらは /review 用に保存されますが、新しいアラートとしてブロードキャストされません。
交換記号の世界は次のときに再発見されます。
data/market_stats/reports/symbol_universe.json に到達
Pump.universeRefreshDays 古い (デフォルト: 4 日)。これにより、新しいリストが追加され、削除されます
その後のフェッチとスキャンからリストから除外された機器。
ディスク上: data/market_stats/ ( archives/ 、 api_fallback/raw/ 、 reports/ )。キャッシュされたシリーズは繰り返し実行ではスキップされます。自動的にプルーニングされることはありません。「ディスクの使用量と保持」を参照してください。
アラート

classifierTelegramChatId へのエントリには、Pump |ダンプ |ボタンはありません。 Pump-bot は、 Pumps.classification を書き込む前に、コールバックがそのチャットからのものであることを確認します。他の加入者は、分類ボタンなしで同じアラートを受信します。
turso db 作成スクリーナー
turso db ショー スクリーナー --url
turso db トークン作成スクリーナー
pnpm db:ブートストラップ
テレグラムの詳細: docs/telegram_setup.md 。
data/market_stats/reports/scan_cache/ にあるコインごとの結果 (基礎となる CLI の --cache-dir でオーバーライド)。エントリは、検出器のバージョン、スキャン パラメータ、ウィンドウの開始、交換ごとに読み込まれたキャンドル ID が変更されていない場合 (ファイルの mtime ではない) に再利用されます。新しい 5 メートルのバーは、増分テール再スキャン (~400 バー) をトリガーします。 UTC 真夜中のウィンドウ ロールにより、完全な再スキャンがトリガーされます。
キャッシュを無効にします。 config.js で Pump.scanCache: false を設定するか、 scan_cache/ を削除してから、pm2 で Pump-monitor を再起動します。
スキャンでは、ワーカー スレッド プール (自動検出された CPU コア) を使用します。実稼働スキャン ステップでは、ネイティブ ワーカー スレッド (Linux/VDS で必要) でコンパイルされた dist/cli.js を使用します。
パイプラインでは何も削除されず、実行されるたびに追加のみが行われます。長期間存続する VDS では、data/market_stats/ は際限なく増加します (数か月で 100 GB 以上が正常です)。
通常、reports/pump_detector/ が大部分を占めます。スキャンごとに新しい実行ディレクトリが作成され、PM2 は終了時にポンプ モニターを再起動します。そのため、コード パスが読み取られないデバッグ出力は 1 日あたり数 GB になります。
スキャン自体には、ろうそくのポンプ.日分だけが必要です。過去のエピソードのチャートを描画するために archives/ 、 api_fallback/ 、および extracted/ を読み取る /review UI の保存期間は長くなります。保存期間を過ぎると、これらのチャートは TradingView にフォールバックしますが、Turso に保存されているポンプ行は影響を受けません。
scripts/prune-market-data.sh は両方の年齢別にプルーニングします。デフォルトでは予行演習されます。
./scripts/prune-mar

ket-data.sh # プレビュー: 一致したディレクトリとサイズ
./scripts/prune-market-data.sh --apply # 削除
デフォルトでは、60 日間のキャンドルと 2 日間の実行ディレクトリが保持されます。 RETAIN_DAYS 、 RUN_RETAIN_DAYS 、 DATA_DIR でオーバーライドします。ポンプモニターがライブである間は安全に実行できます。実行ディレクトリのフロアは実行中の実行を維持し、パイプラインは今日の日付パーティションのみを書き込みます。
毎日の cron で戻ってこないようにします。
(crontab -l 2> /dev/null ; echo " 17 4 * * * cd /path/to/repo && ./scripts/prune-market-data.sh --apply >> /tmp/prune-market-data.log 2>&1 " ) |クロンタブ -
PM2 自身のログは data/ とは別に蓄積されます。 du -sh ~/.pm2/logs で確認してください。 pm2 flash はそれらをクリアし、pm2 install pm2-logrotate は再発を防ぎます。
パイプライン CLI は、手動で 1 回実行する必要がある場合 (デバッグのみ) にフラグを受け入れます。
pnpm ポンプ:モニター -- --no-telegram --cache-dir /path/to/cache
通常の動作は PM2 のままでなければなりません。
ローカル キャッシュ カバレッジ — ルックバック ウィンドウのディスク上のデータがどの程度完全であるか:
pnpm レポート:カバレッジ
pnpm レポート:カバレッジ -- --exchanges binance --days 7
pnpm レポート:カバレッジ -- --start 2026-06-03T00:00:00Z --end 2026-06-04T00:00:00Z --json
出力例:
窓口：2026-06-03 → 2026-06-04（1日）、5m間隔
取引所に存在するすべてのコイン: Binance 82.9% キャッシュ (1133/1367)、Bybit 80.0% キャッシュ (...)
便利なフラグ: --days 、 --exchanges 、 --quote-currency 、 --discover 、 --json 。
Ecosystem.config.cjs — PM2 プロセス定義 ( Pump-monitor 、 Pump-bot 、 Pump-Web )
config.js — Turso、Telegram ボット トークン、pump.* 、 fetch.intervals ( config.example.js からコピー)
scripts/prune-market-data.sh — data/market_stats/ の年齢ベースのクリーンアップ
パッケージ/コア — HTTP クライアント、構成、プル ウィンドウ
パッケージ/交換 — 交換アダプター
パッケージ/ストレージ — NDJSON、ギャップ、マニフェスト I/O
パッケージ/大学

erse — シンボル ユニバースとタスクの解決
パッケージ/アーカイブ — アーカイブ プランナーとダウンロード ランナー
apps/fetch-market-archives — ユニバース アーカイブ プル CLI (+ カバレッジ レポート)
Packages/db — スクリーナー データベース (Turso/libSQL)
Packages/pump-detector — ポンプレジーム検出 ( PUMP_DETECTION_RULES.md )
apps/run-pump-detector — ワーカー プールと Pump_events.ndjson をスキャンします
apps/pump-monitor — フェッチ + スキャン + Turso + Telegram アラート + HTTP ポンプ ページ
apps/fetch-market-stats — REST キャンドル プル CLI (スタンドアロン、 Pump-monitor では使用されません)
このプロジェクトは MIT ライセンスの下でオープンソースです。
Telegram アラートを備えたポンプ アンド ダンプ暗号スクリーナー
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Pump and dump crypto screener with Telegram alerts - aleks-ent/pump-dump-crypto-screener

GitHub - aleks-ent/pump-dump-crypto-screener: Pump and dump crypto screener with Telegram alerts · GitHub
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
aleks-ent
/
pump-dump-crypto-screener
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
69 Commits 69 Commits .cursor/ rules .cursor/ rules apps apps config config docs docs packages packages scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE PUMP_DETECTION_RULES.md PUMP_DETECTION_RULES.md README.md README.md config.example.js config.example.js ecosystem.config.cjs ecosystem.config.cjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml stop.sh stop.sh tsconfig.base.json tsconfig.base.json update.sh update.sh vitest.config.ts vitest.config.ts View all files Repository files navigation
TypeScript pnpm monorepo for pulling 1m/5m market statistics and bulk kline archives from Binance and Bybit, detecting pump/dump regimes, and alerting via Telegram.
Try the live bot: Open @pumpdumpscreenerautobot on Telegram
Requirements: Node.js 20+, pnpm 9+, pm2 (process manager for production).
No exchange API keys. Binance and Bybit data comes from public bulk archives and public REST endpoints.
You need two free accounts (~2 minutes each):
1. Install dependencies and configure
pnpm install
cp config.example.js config.js # fill in Turso + Telegram (see table above)
Edit config.js — at minimum, configure database , telegramBotToken , and your private classifierTelegramChatId . Anyone who sends /start is automatically subscribed to alerts and can use /stats , /runs , and /about ; only the configured chat can classify alerts. Pump lookback and scan settings live under pump (defaults in config.example.js ):
database : {
url : "libsql://screener-....turso.io" ,
authToken : "..." ,
} ,
telegramBotToken : "123456789:ABC..." ,
classifierTelegramChatId : "36772199" ,
web : {
port : 3000 , // local app server; nginx terminates public HTTP/HTTPS
host : "127.0.0.1" ,
} ,
pump : {
days : 5 , // lookback calendar days for download + scan
minScore : 80 ,
scanCache : true ,
requireCalmPrePump : false , // feature flag: require a calm 2h period before pumps
} ,
2. Build and init the database
pnpm build
pnpm db:bootstrap
Run pnpm build again after every code change — PM2 does not rebuild for you.
Production runs are defined in ecosystem.config.cjs at the repo root. It starts three processes:
pm2 start ecosystem.config.cjs
PM2 keeps all processes alive. When pump-monitor finishes a pipeline run it exits; PM2 immediately starts the next run ( autorestart: true ). That replaces a manual cron loop.
First run downloads pump.days of candles into data/market_stats/ — network + disk required; can take hours. Market data is not in the repository.
The web page is served by pump-web on 127.0.0.1:3000 by default. Put nginx in front of it for public HTTP/HTTPS; see docs/nginx_letsencrypt.md .
pm2 save
pm2 startup # follow the printed command (once per server)
Pump Event Reviewer
The manual reviewer at /review turns detected pumps into a human-labeled dataset for
evaluating and improving the screener. It is designed for a fast, keyboard-friendly
workflow:
Browse and filter stored pump events by status, category, exchange, symbol, and date.
Inspect Telegram subscriber votes alongside a 1m or 5m OHLCV chart centered on the
screener's detection time. Charts use local history when available and otherwise load
the four-hour window from the event's public Binance or Bybit API.
Classify each event as a wick spike, weak pump, sustained move, volume only,
illiquid noise, or unclear; optionally record confidence and a comment.
Save and advance with keyboard shortcuts, revisit existing labels, track review
progress, and export labeled datasets as JSON or CSV.
Human annotations are stored separately from the original pump records, so review does
not modify detector output. The workspace uses the existing Turso database, makes no
browser-to-exchange requests, and can be protected with simple HTTP Basic authentication.
See the pump review operator guide for setup,
access control, deployment, and release checks.
pm2 status
pm2 logs # all apps
pm2 logs pump-monitor # download + scan pipeline
pm2 logs pump-bot # Telegram bot
pm2 logs pump-web # HTTP page
pm2 restart ecosystem.config.cjs # after config.js or code changes (rebuild first)
pm2 restart pump-monitor # restart pipeline only
pm2 restart pump-bot # restart bot only
pm2 restart pump-web # restart HTTP page only
pm2 stop ecosystem.config.cjs
pm2 delete ecosystem.config.cjs
After editing config.js , restart the affected process ( pm2 restart pump-monitor , pump-bot , or pump-web ). No PM2 reload is needed for config-only changes if you restart.
Classification buttons need pump-bot running — it is included in ecosystem.config.cjs , not optional in production.
Each PM2-driven run is an end-to-end pipeline:
Download — last pump.days of 1m/5m candles from Binance and Bybit (archives + REST fallback). See docs/fetch_all.md .
Scan — pump/dump detection; output data/market_stats/reports/pump_events.ndjson .
Persist + alert — upsert pump episodes to Turso; Telegram message per new, current pump ( coin|pump_start_utc ). Episodes ending before the previous successful monitor cycle began are historical backfill: they are stored for /review but are not broadcast as fresh alerts.
The exchange symbol universe is re-discovered when
data/market_stats/reports/symbol_universe.json reaches
pump.universeRefreshDays old (default: 4 days). This adds new listings and removes
delisted instruments from subsequent fetches and scans.
On disk: data/market_stats/ ( archives/ , api_fallback/raw/ , reports/ ). Cached series are skipped on repeat runs. Nothing is ever pruned automatically — see Disk usage and retention .
Alerts sent to classifierTelegramChatId have Pump | Dump | None buttons. pump-bot verifies the callback came from that chat before writing pumps.classification ; other subscribers receive the same alerts without classification buttons.
turso db create screener
turso db show screener --url
turso db tokens create screener
pnpm db:bootstrap
Telegram details: docs/telegram_setup.md .
Per-coin results under data/market_stats/reports/scan_cache/ (override with --cache-dir on the underlying CLI). Entries are reused when detector version, scan params, window start, and loaded candle identity per exchange are unchanged — not file mtimes. New 5m bars trigger an incremental tail rescan (~400 bars); UTC midnight window rolls trigger a full rescan.
Disable cache: set pump.scanCache: false in config.js , or delete scan_cache/ , then pm2 restart pump-monitor .
Scanning uses a worker thread pool (auto-detected CPU cores). Production scan step uses compiled dist/cli.js with native worker threads (required on Linux/VDS).
Nothing in the pipeline deletes anything — every run only appends. On a long-lived VDS data/market_stats/ grows without bound (100 GB+ over a couple of months is normal).
reports/pump_detector/ is usually the bulk of it. A fresh run directory is created on every scan, and PM2 restarts pump-monitor the moment it exits, so this is several GB/day of debug output that no code path reads.
Scanning itself only needs pump.days of candles. Longer retention exists for the /review UI, which reads archives/ , api_fallback/ , and extracted/ to draw charts for past episodes — below the retention horizon those charts fall back to TradingView, while the stored pump rows in Turso are unaffected.
scripts/prune-market-data.sh prunes both by age. It is dry-run by default:
./scripts/prune-market-data.sh # preview: matched directories and size
./scripts/prune-market-data.sh --apply # delete
Defaults keep 60 days of candles and 2 days of run directories; override with RETAIN_DAYS , RUN_RETAIN_DAYS , DATA_DIR . Safe to run while pump-monitor is live — the run-directory floor keeps the in-flight run, and the pipeline only writes today's date partitions.
Keep it from coming back with a daily cron:
(crontab -l 2> /dev/null ; echo " 17 4 * * * cd /path/to/repo && ./scripts/prune-market-data.sh --apply >> /tmp/prune-market-data.log 2>&1 " ) | crontab -
PM2's own logs accumulate separately from data/ . Check with du -sh ~/.pm2/logs ; pm2 flush clears them and pm2 install pm2-logrotate prevents recurrence.
The pipeline CLIs accept flags if you need a single manual run (debugging only):
pnpm pump:monitor -- --no-telegram --cache-dir /path/to/cache
Normal operation should stay on PM2.
Local cache coverage — how complete on-disk data is for the lookback window:
pnpm report:coverage
pnpm report:coverage -- --exchanges binance --days 7
pnpm report:coverage -- --start 2026-06-03T00:00:00Z --end 2026-06-04T00:00:00Z --json
Example output:
Window: 2026-06-03 → 2026-06-04 (1 days), intervals 5m
For all coins present on exchange: Binance 82.9% cached (1133/1367), Bybit 80.0% cached (...)
Useful flags: --days , --exchanges , --quote-currencies , --discover , --json .
ecosystem.config.cjs — PM2 process definitions ( pump-monitor , pump-bot , pump-web )
config.js — Turso, Telegram bot token, pump.* , fetch.intervals (copy from config.example.js )
scripts/prune-market-data.sh — age-based cleanup for data/market_stats/
packages/core — HTTP client, config, pull window
packages/exchanges — exchange adapters
packages/storage — NDJSON, gaps, manifest I/O
packages/universe — symbol universe and task resolution
packages/archive — archive planners and download runner
apps/fetch-market-archives — universe archive pull CLI (+ coverage reports)
packages/db — screener database (Turso/libSQL)
packages/pump-detector — pump regime detection ( PUMP_DETECTION_RULES.md )
apps/run-pump-detector — scan worker pool and pump_events.ndjson
apps/pump-monitor — fetch + scan + Turso + Telegram alerts + HTTP pump page
apps/fetch-market-stats — REST candle pull CLI (standalone; not used by pump-monitor )
This project is open source under the MIT License .
Pump and dump crypto screener with Telegram alerts
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
