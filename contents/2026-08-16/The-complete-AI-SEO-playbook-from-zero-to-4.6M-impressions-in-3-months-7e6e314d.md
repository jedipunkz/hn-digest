---
source: "https://github.com/TraceCohenTech/ai-seo-playbook"
hn_url: "https://news.ycombinator.com/item?id=49315620"
title: "The complete AI SEO playbook: from zero to 4.6M impressions in 3 months"
article_title: "GitHub - TraceCohenTech/ai-seo-playbook: The complete AI SEO playbook: methodology, scripts, and safety guards behind a 4.6M-impression content engine. GSC feedback loops, multi-model agent orchestration, quality gates, and build cost control. · GitHub"
author: "t_cohen"
captured_at: "2026-08-16T00:41:12Z"
capture_tool: "hn-digest"
hn_id: 49315620
score: 2
comments: 0
posted_at: "2026-08-16T00:19:59Z"
tags:
  - hacker-news
  - translated
---

# The complete AI SEO playbook: from zero to 4.6M impressions in 3 months

- HN: [49315620](https://news.ycombinator.com/item?id=49315620)
- Source: [github.com](https://github.com/TraceCohenTech/ai-seo-playbook)
- Score: 2
- Comments: 0
- Posted: 2026-08-16T00:19:59Z

## Translation

タイトル: 完全な AI SEO ハンドブック: 3 か月でゼロから 460 万インプレッションまで
記事のタイトル: GitHub - TraceCohenTech/ai-seo-playbook: 完全な AI SEO プレイブック: 460 万インプレッションのコンテンツ エンジンの背後にある方法論、スクリプト、安全対策。 GSC フィードバック ループ、マルチモデル エージェント オーケストレーション、品質ゲート、およびビルド コスト管理。 · GitHub
説明: 完全な AI SEO プレイブック: 460 万インプレッションのコンテンツ エンジンの背後にある方法論、スクリプト、安全対策。 GSC フィードバック ループ、マルチモデル エージェント オーケストレーション、品質ゲート、およびビルド コスト管理。 - TraceCohenTech/ai-seo-playbook

記事本文:
GitHub - TraceCohenTech/ai-seo-playbook: 完全な AI SEO プレイブック: 460 万インプレッションのコンテンツ エンジンの背後にある方法論、スクリプト、安全対策。 GSC フィードバック ループ、マルチモデル エージェント オーケストレーション、品質ゲート、およびビルド コスト管理。 · GitHub
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
トレースコーエンテック
/
ai-seo-プレイブック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .github .github config conf

ig docs docs 例 例 サンプル サンプル スキーマ スキーマ スクリプト スクリプト .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
3 か月で 0 インプレッションから 460 万インプレッションまで、実際にランク付けされる AI を活用したコンテンツ エンジンを構築するための完全なプレイブック。
これは、AI エージェント、GSC フィードバック ループ、自動品質ゲートを使用して ValueAddVC.com でコンテンツ エンジンを構築することで得た方法論、ツールキット、そして苦労して得た教訓です。 14 の診断スクリプト、9 つの実戦テスト済み構成 (安全ガード、エージェント オーケストレーション、品質ゲート、AI 対策検出)、構造化データ スキーマ、CI 自動化など、システムを複製するために必要なものがすべて含まれています。
理論ではありません。プロンプトではありません。サイトの実際のオペレーティング システムでは、月間インプレッション数が 604,000 から 462,000 に増加しました。
Trace Cohen によって ValueAddVC.com で構築されました。
Content Engine — AI エージェント オーケストレーション (マルチモデル パイプライン: 計画には Opus/Fable、執筆には Sonnet、単調な作業には Haiku)、5 形式のコンテンツ ローテーション、音声トレーニング、AI 対策指紋検出
GSC フィードバック ループ — 毎週の自動レポート、タイトル書き換え候補、共食い検出、クエリ ギャップ マイニング、射程距離の最適化
品質システム — 9 つのパブリッシュ ゲート、テンプレート フレーズ ブロックリスト、ソース検証、ファクト チェック、構造化データ検証
安全層 — リポロック、リベースガード、ビルドコスト管理 ([nobuild] タグ、デプロイティック)、自己修復ハートビート、git からのコンテンツライターの分離
成長ループ — キーワードの予測 (需要が急増する前に公開)、ライブ ページの更新、内部リンク グラフの最適化、ニュース サイトマップ + WebSub によるインスタント クロール
フローチャート LR
GSC[Google Search Console API] --> レポート[Wee]

クライレポート】
レポート --> トリアージ{Triage}
トリアージ --> リライト[タイトルリライト]
トリアージ --> マージ[人食い人種のマージ]
トリアージ --> 更新[古いページを更新]
トリアージ --> キル[デッドウェイトをキル]
コンテンツ[コンテンツ パイプライン] --> ゲート[品質ゲート]
ゲート --> |パス|公開[公開]
ゲート --> |失敗|修正[修正して再試行]
発行 --> GSC
書き換え --> GSC
マージ --> GSC
更新 --> GSC
スタイル GSC 塗りつぶし:#0d7377、カラー:#fff
スタイル ゲート塗りつぶし:#0d7377、カラー:#fff
スタイル パブリッシュ 塗りつぶし:#15803d、カラー:#fff
スタイルキルフィル:#b91c1c、カラー:#fff
読み込み中
フィードバック ループ: GSC データが診断スクリプトにフィード → スクリプトが修正が必要な点を明らかにする → AI エージェントが品質ゲートを通過して修正を行う → ランキングの向上によりより良い GSC データが生成される → 繰り返し。システムは毎週賢くなっています。
スクリプト
何をするのか
gsc-rewrite-candidates.mjs
GSC データからタイトル書き換えの機会を見つける — インプレッション数は高いが CTR が低いページ ランキング 4 ～ 20 位
テンプレート検出器.mjs
コンテンツをスキャンして AI テンプレート フィンガープリント（拡大コンテンツの不正使用を Google に通知する繰り返しのフレーズ）を検出します
共食い検出器.mjs
同じクエリで競合し、権限が分割され、1 つの統合ページよりもランクが下がっているサイト上のページを見つけます。
週報.mjs
トレンドのクエリ、削除されたページ、CTR トリアージ候補、およびクエリの独占を含む週次 SEO パフォーマンス レポートを生成します。
orphan-finder.mjs
インバウンド内部リンクがゼロのページを検索します。Google のリンクグラフ クローラーには表示されません。
コンテンツ監査.mjs
GSC データとコンテンツ品質に基づいて、すべてのページを KILL / MERGE / UPDATE / PROMOTE / KEEP バケットにスコア付けします
リダイレクトチェッカー.mjs
200 ではなく 301/302/308 を返す URL をサイトマップ内で検索します。これらは GSC 検証を破り、クロール予算を無駄にします。
リフレッシュトラッカー.mjs
更新されていないトラフィックの多いページを特定します。

最近 — 「リフレッシュドリップ」戦略の候補者
クエリ-ギャップ-マイナー.mjs
遡及的なキーワード検出エンジン - 専用ページがない場合でも、実際の需要があるクエリを見つけます。 Google は何を書くべきかを教えてくれます。
射程距離.mjs
実際のインプレッション数が 5 ～ 20 位にランク付けされているページを検索します。SEO では最も安価なページが勝ちます。改善された場合はクリックゲインを推定します。
書き換え-measurer.mjs
タイトル書き換えの前後の追跡。ベースラインを取得し、変更を加えて、2 ～ 4 週間後の影響を測定します。
websub-ping.mjs
フィードが変更されたことを Google のハブに通知します。何時間も待たずにすぐにクロールをトリガーします。パブリッシュするたびに実行します。
インデックス作成-submitter.mjs
URL を Google の Indexing API に送信して、ほぼ瞬時にクロールできるようにします。 1 日あたり 200 URL の割り当て。
壊れたリンクチェッカー.mjs
送信リンクのすべてのコンテンツをスキャンし、404、タイムアウト、およびリダイレクト チェーンをチェックします。 CI の場合はゼロ以外で終了します。
構成 ( /config )
ファイル
目的
フォーマット回転.json
5 つのフォーマットのコンテンツ システム: Deep Explainer、ニュース分析、ランク付きリスト、質問主導、逆張りテイク - フォーマットごとのワード数、チャート要件、および選択の重み付き
品質ゲート.json
パブリッシュゲートルール: カニバリゼーションチェック、ソース URL 検証、テンプレートフレーズ検出、共有クローザー検出、タイプチェック
アンチAIルール.json
AI コンテンツを人間らしく見せるための AI テンプレート フレーズとスタイル ルールの完全なブロックリスト
リフレッシュルール.json
更新ドリップ戦略のルール - コンテンツ タイプ別の古さのしきい値、更新トリガー、更新チェックリスト
キーワード-予想.json
イベント カレンダー手法 — IPO、収益、資金調達ラウンド、規制の前にコンテンツを公開することで、需要が急増したときにランク付けされます。
ヘルスチェック.json
ライブサイトのヘルスチェック: 漏洩したテンプレート変数、破損した OG 画像、挿入された広告リンク、内容の薄いコンテンツ、DEA

dページ
コンテンツパイプラインガード.json
安全対策: リポロック、リベースガード、カニバリゼーションチェック、ビルドコスト管理、自己修復ハートビート
エージェントオーケストレーション.json
マルチモデル AI パイプライン ルール: 計画には Opus/Fable、執筆には Sonnet、機械的タスクには Haiku。最大 3 人の同時エージェント。
スキーマの例 ( /schemas )
ファイル
スキーマの種類
著者付き記事.json
記事 + 人物著者エンティティ (E-E-A-T 財団)
よくある質問ページ.json
ブログ投稿の FAQPage — FAQ のリッチリザルトを促進します
アイテムリスト.json
ランキング/比較ページのItemList — スポンサーが望むフォーマット
ニュース記事.json
リアルタイム コンテンツ用のニュース記事 + ニュース サイトマップ テンプレート
例 ( /examples )
sitemap.ts — 正直な lastmod 日付を含む Next.js 動的サイトマップ
news-sitemap.ts — Google News/Discoverの48時間ローリングニュースサイトマップ
Internal-link-component.tsx — 関連投稿の React コンポーネント + ビルド時の内部リンク インサーター
vercel-ignore.sh — Vercel のスキップ ロジックの構築: [nobuild] タグ、コンテンツのみの検出、デプロイ ティック パターン ($$$ の節約)
すべてのスクリプトにはサンプル出力ファイルが含まれているため、何かを実行する前に何が起こるかを確認できます。
Weekly-report.json — クエリの傾向、ページの削除、CTR のトリアージを含む完全な週次レポート
rewrite-candidates.json — クエリごとの診断によるタイトル書き換えの機会
content-audit.json — バケット割り当ての KILL/MERGE/UPDATE/PROMOTE/KEEP
cannibal-clusters.json — インプレッション推定が無駄になった共食いクラスター
template-scan.json — ファイルごとのフレーズの場所を使用した AI 指紋スキャン
orphan-pages.json — 孤立した、低リンク、行き止まりのページ レポート
setup-gsc.md — ステップバイステップの Google Search Console API セットアップ (ローカル認証 + CI のサービス アカウント)
自動化 ( .github/workflows )
Weekly-seo-report.yml — 毎週の繰り返しを実行する GitHub アクション

毎週日曜日に ort を実行し、結果をコミットし、必要に応じて概要を含む GitHub 問題を作成します
# リポジトリのクローンを作成する
git clone https://github.com/TraceCohenTech/ai-seo-playbook.git
CD あいそプレイブック
# 依存関係をインストールする
npmインストール
# Google Search Console API アクセスを設定する
# (Search Console API が有効になっている Google Cloud プロジェクトが必要)
gcloud auth アプリケーション - デフォルトのログイン \
--scopes=https://www.googleapis.com/auth/webmasters.readonly
# タイトル書き換えの機会を見つける
npm run rewrite-candidates -- --site sc-domain:yoursite.com
# AI テンプレートのフィンガープリントをスキャンします
npm run template-scan -- --dir ./your-content-directory
# カニバリゼーションクラスターを見つける
npm run find-cannibals -- --site sc-domain:yoursite.com
# 完全なコンテンツ監査を実行する
npm run content-audit -- --site sc-domain:yoursite.com --dir ./your-content-directory
# 孤立したページを検索する (内部リンクなし)
npm run find-orphans -- --dir ./your-content-directory
# 週次レポートを生成する
npm run Weekly-report -- --site sc-domain:yoursite.com
# すでにランキングされているが、ページ ターゲティングが設定されていないキーワードを発見する
npm run query-gaps -- --site sc-domain:yoursite.com --dir ./your-content-directory
# 小さなナッジ = 大きなクリックが得られる「ほぼページ 1」のページを検索します
npm run 打撃距離 -- --site sc-domain:yoursite.com
# 更新が必要な古いページを見つける
npm runfresh-tracker -- --site sc-domain:yoursite.com --dir ./your-content-directory
# サイトマップ内のリダイレクトの問題を確認する
npm run check-redirects -- --site sc-domain:yoursite.com --sitemap https://yoursite.com/sitemap.xml
# Google に ping を送信して、更新されたフィードをすぐにクロールします
npm run websub-ping -- --feeds https://yoursite.com/sitemap.xml,https://yoursite.com/feed.xml
GSC API は初めてですか?ステップバイステップのセットアップ ガイドについては、docs/setup-gsc.md を参照してください。
これらのツールは 1 ヘクタールです

システムの lf。これらの特定の指標が重要な理由、結果の解釈方法、コンテンツ エンジンを自己改善させるフィードバック ループの構築方法などの方法論は、完全ガイドに記載されています。
AI SEO ハンドブック: AI を使用して 3 か月で 460 万インプレッションを達成するコンテンツ エンジンを構築した方法
コンテンツ エンジンの構築 (アーキテクチャ、音声トレーニング、フォーマット ローテーション)
GSC の計算 (AI の概要の発見、タイトルの書き換え、カニバリゼーション)
反復ループ (キーワードの予測、生きているページ、技術的な SEO のバグ)
システム（品質ゲート、週次レビュー、コスト管理）
これらのツールは、ValueAddVC.com で 3 か月かけて構築され、改良されました。
*インプレッションが最大 8 倍に増加したため、CTR は 0.4% です。これは主に、本質的にクリックを生成しない AI 概要引用 (GEO トラフィック) によるものです。人間の意図による CTR が向上しました。ランク付けされたリストは 6.8% に達し、質問主導の投稿は 3.2% に達しました。成長曲線はほぼ垂直で、8 月 13 日だけでインプレッション数は 127,000 回、クリック数は 854 回に達しました。
オプション 1: GitHub アクション (推奨)
Search Console API にアクセスできる Google Cloud サービス アカウントを作成する
サービス アカウント JSON を GSC_CREDENTIALS という名前の GitHub シークレットとして追加します
リポジトリ変数 GSC_SITE を GSC プロパティ (例: sc-domain:yoursite.com ) に設定します。
CONTENT_DIR をコンテンツ ディレクトリ パス (例: ./src/app/bl) に設定します。

[切り捨てられた]

## Original Extract

The complete AI SEO playbook: methodology, scripts, and safety guards behind a 4.6M-impression content engine. GSC feedback loops, multi-model agent orchestration, quality gates, and build cost control. - TraceCohenTech/ai-seo-playbook

GitHub - TraceCohenTech/ai-seo-playbook: The complete AI SEO playbook: methodology, scripts, and safety guards behind a 4.6M-impression content engine. GSC feedback loops, multi-model agent orchestration, quality gates, and build cost control. · GitHub
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
TraceCohenTech
/
ai-seo-playbook
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github .github config config docs docs examples examples samples samples schemas schemas scripts scripts .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package.json package.json View all files Repository files navigation
The complete playbook for building an AI-powered content engine that actually ranks — from zero to 4.6M impressions in 3 months.
This is the methodology, the toolkit, and the hard-won lessons from building a content engine on ValueAddVC.com using AI agents, GSC feedback loops, and automated quality gates. 14 diagnostic scripts, 9 battle-tested configs (safety guards, agent orchestration, quality gates, anti-AI detection), structured data schemas, and CI automation — everything you need to replicate the system.
Not theory. Not prompts. The actual operating system behind a site that went from 604K to 4.62M monthly impressions.
Built by Trace Cohen at ValueAddVC.com .
The Content Engine — AI agent orchestration (multi-model pipelines: Opus/Fable for planning, Sonnet for writing, Haiku for grunt work), 5-format content rotation, voice training, anti-AI fingerprint detection
The GSC Feedback Loop — Weekly automated reports, title rewrite candidates, cannibalization detection, query gap mining, striking distance optimization
The Quality System — 9 publish gates, template phrase blocklists, source verification, fact-checking, structured data validation
The Safety Layer — Repo locks, rebase guards, build cost control ([nobuild] tags, deploy-tick), self-healing heartbeats, content writer isolation from git
The Growth Loop — Keyword anticipation (publish before demand spikes), living page refreshes, internal link graph optimization, news sitemap + WebSub for instant crawling
flowchart LR
GSC[Google Search Console API] --> Report[Weekly Report]
Report --> Triage{Triage}
Triage --> Rewrite[Title Rewrites]
Triage --> Merge[Merge Cannibals]
Triage --> Refresh[Refresh Stale Pages]
Triage --> Kill[Kill Dead Weight]
Content[Content Pipeline] --> Gates[Quality Gates]
Gates --> |Pass| Publish[Publish]
Gates --> |Fail| Fix[Fix & Retry]
Publish --> GSC
Rewrite --> GSC
Merge --> GSC
Refresh --> GSC
style GSC fill:#0d7377,color:#fff
style Gates fill:#0d7377,color:#fff
style Publish fill:#15803d,color:#fff
style Kill fill:#b91c1c,color:#fff
Loading
The feedback loop: GSC data feeds diagnostic scripts → scripts surface what needs fixing → AI agents make the fixes through quality gates → improved rankings produce better GSC data → repeat. Every week the system gets smarter.
Script
What It Does
gsc-rewrite-candidates.mjs
Finds title rewrite opportunities from GSC data — pages ranking position 4–20 with high impressions but low CTR
template-detector.mjs
Scans your content for AI template fingerprints — the repeated phrases that signal scaled-content-abuse to Google
cannibalization-detector.mjs
Finds pages on your site competing for the same queries, splitting authority and ranking worse than one consolidated page would
weekly-report.mjs
Generates a weekly SEO performance report with trending queries, dropping pages, CTR triage candidates, and query monopolies
orphan-finder.mjs
Finds pages with zero inbound internal links — invisible to Google's link-graph crawler
content-audit.mjs
Scores every page into KILL / MERGE / UPDATE / PROMOTE / KEEP buckets based on GSC data + content quality
redirect-checker.mjs
Finds URLs in your sitemap that return 301/302/308 instead of 200 — these break GSC validation and waste crawl budget
refresh-tracker.mjs
Identifies high-traffic pages that haven't been updated recently — candidates for the "refresh drip" strategy
query-gap-miner.mjs
The retroactive keyword discovery engine — finds queries with real demand where you have no dedicated page. Google is telling you what to write.
striking-distance.mjs
Finds pages ranking position 5-20 with real impressions — the cheapest wins in SEO. Estimates click gain if improved.
rewrite-measurer.mjs
Before/after tracking for title rewrites. Take a baseline, make changes, measure impact 2-4 weeks later.
websub-ping.mjs
Notifies Google's hub that your feeds changed — triggers immediate crawl instead of waiting hours. Run after every publish.
indexing-submitter.mjs
Submits URLs to Google's Indexing API for near-instant crawling. 200 URLs/day quota.
broken-link-checker.mjs
Scans all content for outbound links and checks for 404s, timeouts, and redirect chains. Exits non-zero for CI.
Configuration ( /config )
File
Purpose
format-rotation.json
The 5-format content system: Deep Explainer, News Analysis, Ranked List, Question-Led, Contrarian Take — with per-format word counts, chart requirements, and selection weights
quality-gates.json
Publish gate rules: cannibalization check, source URL verification, template phrase detection, shared closer detection, typecheck
anti-ai-rules.json
The complete blocklist of AI template phrases + style rules for making AI content sound human
refresh-rules.json
Rules for the refresh drip strategy — staleness thresholds by content type, refresh triggers, and a refresh checklist
keyword-anticipation.json
Event calendar methodology — publish content before IPOs, earnings, funding rounds, regulations so you're ranked when demand spikes
health-checks.json
Live-site health checks: leaked template variables, broken OG images, injected ad links, thin content, dead pages
content-pipeline-guards.json
Safety guards: repo locks, rebase guards, cannibalization checks, build cost control, self-healing heartbeats
agent-orchestration.json
Multi-model AI pipeline rules: Opus/Fable for planning, Sonnet for writing, Haiku for mechanical tasks. Max 3 concurrent agents.
Schema Examples ( /schemas )
File
Schema Type
article-with-author.json
Article + Person author entity (the E-E-A-T foundation)
faq-page.json
FAQPage for blog posts — drives FAQ rich results
item-list.json
ItemList for ranking/comparison pages — the format sponsors want
news-article.json
NewsArticle + news sitemap template for real-time content
Examples ( /examples )
sitemap.ts — Next.js dynamic sitemap with honest lastmod dates
news-sitemap.ts — 48-hour rolling news sitemap for Google News/Discover
internal-link-component.tsx — React component for related posts + a build-time internal link inserter
vercel-ignore.sh — Build skip logic for Vercel: [nobuild] tags, content-only detection, deploy-tick pattern (saves $$$)
Every script has a sample output file so you can see what to expect before running anything:
weekly-report.json — Full weekly report with trending queries, dropping pages, CTR triage
rewrite-candidates.json — Title rewrite opportunities with per-query diagnosis
content-audit.json — KILL/MERGE/UPDATE/PROMOTE/KEEP bucket assignments
cannibal-clusters.json — Cannibalization clusters with wasted impression estimates
template-scan.json — AI fingerprint scan with per-file phrase locations
orphan-pages.json — Orphan, low-link, and dead-end page reports
setup-gsc.md — Step-by-step Google Search Console API setup (local auth + service account for CI)
Automation ( .github/workflows )
weekly-seo-report.yml — GitHub Action that runs the weekly report every Sunday, commits results, and optionally creates a GitHub issue with the summary
# Clone the repo
git clone https://github.com/TraceCohenTech/ai-seo-playbook.git
cd ai-seo-playbook
# Install dependencies
npm install
# Set up Google Search Console API access
# (requires a Google Cloud project with Search Console API enabled)
gcloud auth application-default login \
--scopes=https://www.googleapis.com/auth/webmasters.readonly
# Find title rewrite opportunities
npm run rewrite-candidates -- --site sc-domain:yoursite.com
# Scan for AI template fingerprints
npm run template-scan -- --dir ./your-content-directory
# Find cannibalization clusters
npm run find-cannibals -- --site sc-domain:yoursite.com
# Run a full content audit
npm run content-audit -- --site sc-domain:yoursite.com --dir ./your-content-directory
# Find orphan pages (no internal links)
npm run find-orphans -- --dir ./your-content-directory
# Generate weekly report
npm run weekly-report -- --site sc-domain:yoursite.com
# Discover keywords you're already ranking for but have no page targeting
npm run query-gaps -- --site sc-domain:yoursite.com --dir ./your-content-directory
# Find "almost page 1" pages where a small nudge = big click gains
npm run striking-distance -- --site sc-domain:yoursite.com
# Find stale pages that need refreshing
npm run refresh-tracker -- --site sc-domain:yoursite.com --dir ./your-content-directory
# Check for redirect problems in your sitemap
npm run check-redirects -- --site sc-domain:yoursite.com --sitemap https://yoursite.com/sitemap.xml
# Ping Google to crawl your updated feeds immediately
npm run websub-ping -- --feeds https://yoursite.com/sitemap.xml,https://yoursite.com/feed.xml
New to the GSC API? See docs/setup-gsc.md for a step-by-step setup guide.
These tools are one half of the system. The methodology — why these specific metrics matter, how to interpret the results, and how to build the feedback loop that makes your content engine self-improving — is in the full guide:
The AI SEO Playbook: How I Used AI to Build a Content Engine That Hit 4.6M Impressions in 3 Months
Building the content engine (architecture, voice training, format rotation)
The GSC reckoning (the AI-overview discovery, title rewrites, cannibalization)
The iteration loop (keyword anticipation, living pages, technical SEO bugs)
The system (quality gates, weekly reviews, cost control)
These tools were built and refined on ValueAddVC.com over 3 months:
*CTR is 0.4% because impressions grew ~8x — largely from AI-overview citations (GEO traffic) that don't produce clicks by nature. Human-intent CTR improved: ranked lists hit 6.8%, question-led posts hit 3.2%. The growth curve is near-vertical: Aug 13 alone hit 127K impressions and 854 clicks.
Option 1: GitHub Actions (recommended)
Create a Google Cloud service account with Search Console API access
Add the service account JSON as a GitHub secret named GSC_CREDENTIALS
Set the repository variable GSC_SITE to your GSC property (e.g., sc-domain:yoursite.com )
Set CONTENT_DIR to your content directory path (e.g., ./src/app/bl

[truncated]
