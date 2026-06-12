---
source: "https://github.com/sandeepmvl/rails-skills"
hn_url: "https://news.ycombinator.com/item?id=48502415"
title: "Production-Grade Claude/AI Skills for Ruby on Rails"
article_title: "GitHub - sandeepmvl/rails-skills · GitHub"
author: "thinkingemote"
captured_at: "2026-06-12T13:18:54Z"
capture_tool: "hn-digest"
hn_id: 48502415
score: 1
comments: 1
posted_at: "2026-06-12T10:51:52Z"
tags:
  - hacker-news
  - translated
---

# Production-Grade Claude/AI Skills for Ruby on Rails

- HN: [48502415](https://news.ycombinator.com/item?id=48502415)
- Source: [github.com](https://github.com/sandeepmvl/rails-skills)
- Score: 1
- Comments: 1
- Posted: 2026-06-12T10:51:52Z

## Translation

タイトル: Ruby on Rails のための本番グレードのクロード/AI スキル
記事のタイトル: GitHub - Sandeepmvl/rails-skills · GitHub
説明: GitHub でアカウントを作成して、sandeepmvl/rails-skills の開発に貢献します。

記事本文:
GitHub - Sandeepmvl/rails-skills · GitHub
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
サンディープムvl
/
レールスキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード さらにアクション メニューを開く 折りたたむ

ファイルとファイル
5 コミット 5 コミット .github .github docs docs skill skill .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md HOW_TO_USE.md HOW_TO_USE.md LICENSE LICENSE PLAN.md PLAN.md README.md README.md すべてのファイルを表示リポジトリ ファイルのナビゲーション
Ruby on Rails の本番グレードのクロード スキル。 Rails の規約に基づいて AI コーディング エージェントと戦うのはやめましょう。Rails スキルを身につけると、Claude Code、Cursor、Codex、Gemini CLI、Antigravity、Windsurf が、上級 Rails 開発者が実際に記述する方法で Rails コードを作成できるようになります。
Claude Code または Cursor に「投稿に公開が予定されている下書きステータスを追加する」ように依頼します。これにより、移行、コールバック、コントローラーの更新が生成されます。見た目は大丈夫です。すると、次のことに気づきます。
既存の PublishPostJob を使用する代わりに before_save コールバックを追加しました。
モデルにすでに存在する列挙型の代わりにステータスの文字列列を使用しました。
PostgreSQL が新しいステータスに必要とする部分インデックスは追加されませんでした
作成されたテストは既存の Post ファクトリを使用しません
AI エージェントを使用しているすべての Rails 開発者は、今朝こんなことを経験したことがあるでしょう。エージェントは愚かではありません。Rails の規則を知らないだけです。それを教えてくれるのがrails-skillです。
独立した Claude Skills のパック (オープン スタンダード、2025 年 12 月に Anthropic によってリリース) は、上級 Rails 開発者が気にするパターン、パフォーマンス トラップ、セキュリティ ベースライン、およびデプロイメント規則をカバーしています。各スキルは以下の間で移植可能です。
多くのプロジェクトにまたがって仕事をしていますか?リポジトリごとではなくグローバルに一度インストールします。 Rails スキルは説明ゲート型です。Rails 以外のプロジェクト (React、Go、Python など) では休止状態のままで、Rails 作業の場合にのみ目覚めます。 「2 つのインストール モード」を参照してください。
cd < your-rails-app >
git clone https://github.com/sandeepmvl/rails-skills .claude/skills
グローバル — すべてのプロジェクト (複数プロジェクトの開発者に推奨):
git

クローン https://github.com/sandeepmvl/rails-skills ~ /.claude/skills/rails-skills
次にクロードコードを起動します。オーケストレーター スキル (rails-project-discovery ) はあなたにインタビューし、適切な下流スキルに自動的にルーティングされます。これは、実際に Rails で作業している場合にのみトリガーされます。
cd < your-rails-app >
git clone https://github.com/sandeepmvl/rails-skills .cursor/skills
その他のツール
OpenAI Codex、Gemini CLI、Antigravity、Windsurf の手順については、docs/install.md を参照してください。
#
スキル
何をするのか
00
レールプロジェクトの発見
Orchestrator — アプリについてインタビューし、以下の適切なスキルを導きます
01
アクティブなレコードパターン
アソシエーション、スコープ、コールバック、インクルード対プリロード対eager_load
02
nプラスワンキラー
N+1 クエリを検出、診断、排除する
03
サービスオブジェクトとファットモデルの比較
いつロジックをモデルに保持し、いつサービス オブジェクトを抽出するか
04
rspec-テストピラミッド
RSpec、FactoryBot、VCR、システム仕様 - どのレイヤーで何をテストするか
05
安全な移行
Strong_migrations ルール、ゼロダウンタイムパターン、バッチでのバックフィル
06
レール-API-デザイン
バージョン管理、シリアル化、ページネーション、レート制限、JWT 認証
07
ソリッドキューとサイドキック
冪等ジョブ、再試行、スケジュールされた作業のいずれかを選択する
08
考案者-rodauth
Rails 方式で実行される認証 + 認可パターン
09
kamal-docker-production
マルチステージ Dockerfile、docker-compose 開発、Kamal 2 デプロイ
10
レール-セキュリティ-ベースライン
強力なパラメータ、CSRF、Brakeman、シークレット、JWT、CORS、Rails 用の OWASP
11
レールキャッシュ戦略
ソリッド キャッシュ、Redis、フラグメントと低レベル、キャッシュ スタンピードの防止
12
ホットワイヤーターボ刺激
ターボドライブ/フレーム/ストリーム、スティミュラス、アクション ケーブルブロードキャスト
13
アクティブストレージアップロード
直接アップロード、バリアント、署名付き URL、検証、S3/GCS サービス構成
14
アクションメーラーベースライン
メール

er セットアップ、deliver_later、プレビュー、仕様、バウンス、冪等送信
15
可観測性ベースライン
lograge 構造化ログ、Sentry、PII スクラビング、Rails.error.report
v0.2 — 拡張 (24 スキル)
#
スキル
何をするのか
16
レールのアップグレード-7-to-8
デュアル Gemfile の次のバージョンのパターン。 7.x → 8 ホップ
17
レールのアップグレード-6-to-7
importmap / jsbundling、暗号化、トリロジーのイントロ
18
レールアップグレード-5-to-6
Zeitwerk の移行、Webpacker の決定
19
レールのアップグレード-4-to-5
APIモード、アクションケーブル、アプリケーションレコード
20
レールのアップグレード-3-to-4
強力なパラメータ、ターボリンク、暗号化されたシークレット
21
db-移行-postgres-mysql
型マッピング、JSONB→JSON、Trilogy gem スワップ
22
db-移行-mysql-postgres
UTF8MB4→UTF-8、prepared_statements + pgbouncer
23
db-移行-oracle-postgres
ora2pg、DATE→TIMESTAMPTZ、シーケンス
24
レールで反応する
Inertia.js デフォルト、vite_rails、API+SPA alt
25
vue-with-rails
イナーシャ + Vue 3 + Vite
26
角型レール付き
スタンドアロン コンポーネント、 @for / @if 、provideHttpClient
27
puma のチューニングと同時実行性
ワーカー/スレッド、jemalloc、YJIT、PumaWorkerKiller
28
資産パイプラインプロペラシャフト
プロペラシャフトとスプロケット、jsbundling、importmap
29
複数のデータベースとレプリカ
connect_to、スティッキー書き込み、シャーディングに関する注意事項
30
Webhook の処理
べき等挿入、HMAC 検証、DLQ パターン
31
ストライプウェブフック統合
construct_event、PaymentIntent、冪等キー
32
外部 API 統合
ファラデー 2、信号遮断器、VCR
33
機能フラグ付け
フリッパー、段階的ロールアウト、キルスイッチ
34
フォームオブジェクトクエリオブジェクトプレゼンター
ファットモデルを超えたアーキテクチャの抽出
35
アクションテキスト-リッチテキスト
Trix、サニタイズ、埋め込み添付ファイル、検索インデックス
36
i18n とタイムゾーン
Rails-i18n、フォールバック、Time.current、DST トラップ
37
レール検索
pg_search → Meilisearch → Elasticsearch 階層化
38
マルチテナンシー
テナントとして機能する、

サブドメインリゾルバー、require_tenant
39
コンソール-安全性-生産
サンドボックス、監査ログ、コンソール経由のタスクのレイク
v0.3 — 専門化 (18 スキル)
#
スキル
何をするのか
40
マイクロサービスを使用すべきでない場合
マイクロサービスを誘導する前に拒否します。モジュラーモノリスのデフォルト。
41
マイクロサービスの分解
境界付きコンテキスト、所有データ、JWT/mTLS、サガ
42
モノリスからサービスへの抽出
ストラングラー フィグ、ダーク起動、二重書き込み、カットオーバー
43
カフカレール
karafka 2.x、送信ボックス、スキーマ レジストリ、パーティショニング
44
うさぎmqレール
バニー + スニーカー、DLX、プリフェッチ、手動 ACK
45
redis-ストリーム-レール
XADD、XREADGROUP、XAUTOCLAIM リーパー
46
cdc-debezium-rails
Postgres 論理レプリケーション、アウトボックス + EventRouter SMT
47
イベント駆動型アーキテクチャ
ドメインと統合イベント、rails_event_store、outbox
48
分散トレースレール
OpenTelemetry、OTLP、手荷物、サンプリング、サンプル
49
可観測性-レール-アドバンスト
RED/USE、SLO、マルチバーンレートアラート、ランブック
50
ヒパレール
PHI 処理、AR 暗号化、監査ログ、BAA
51
PCI-DSS-レール
SAQ-A パス、ストライプ要素、PAN/CVV は使用しない
52
gdpr-レール
DSAR、消去、法的根拠、同意、保持
53
soc2レール
トラストサービス、監査ログ、アクセスレビュー、MFA
54
データウェアハウスの統合
Fivetran/Airbyte 取り込み、dbt 変換、リバース ETL
55
ci-cd-github-actions-rails
マトリックス シャーディング、バンドラー キャッシュ、OIDC、Kamal デプロイ
56
ci-cd-gitlab-rails
ステージ、サービス、OIDC の id_tokens、Auto DevOps の警告
57
ci-cd-jenkins-rails
Jenkins ではない場合の宣言型パイプライン、認証情報ストア
ツーリングとプロジェクト適合 (2 つのスキル)
#
スキル
何をするのか
58
Rubocop とコードの品質
RuboCop + おまかせプリセット +rubocop-rails/-performance/-rspec、SimpleCov、erb_lint、RBS vs Sorbet
59
足場プロジェクトのスキル
アプリにインタビューし、プロジェクト固有のローカル スキルを生成します (テナント、ドメイン ワークフロー、

検証ループ) — 汎用パックでは出荷できない製品対応レイヤー
完全なロードマップと理論的根拠については、PLAN.md を参照してください。
同じプロンプト — 「投稿者とともに最新の 20 件の投稿をリストしてください」 — 同じエージェントに対して。パックなし、その後パックあり。
Rails スキルなし (AI に典型的 - 本番環境でのみ使用される隠れた N+1):
# アプリ/コントローラー/posts_controller.rb
デフォルトインデックス
@posts = 投稿 。 order (created_at::desc) 。制限 ( 20 )
終わり
<% @posts.each は |post| を行います%>
<%= post.title %> — <%= post.author.name %> <%# 1 + 20 クエリ %>
<% 終了 %>
Rails スキル (n-plus-one-killer + activerecord-patterns スキル トリガー) を使用する場合:
# アプリ/コントローラー/posts_controller.rb
デフォルトインデックス
# include(:author) → 行数に関係なく、合計 2 つのクエリ。
@posts = 投稿 。 (:author) が含まれます。 order (created_at::desc) 。制限 ( 20 )
終わり
また、エージェントは、出荷前に開発中の次の N+1 をキャッチするための Bullet 構成を追加し、ここで手動結合に勝る理由を説明します。それが違いです。単に正しいコードではなく、上級 Rails 開発者が与えるような推論を備えた従来のコードです。
一般公開前に、この正確なフローの 30 秒のスクリーンキャストがここに表示されます。それまでは、上記の動作例はデモです。
魔法ではありません。 Rails-upgrade-4-to-8 スキルがアップグレードをガイドし、加速します。それは自主的に行うものではありません。 Rails のアップグレードでは、Gem の置き換えやテストの修正について人間の判断が必要です。
Rails ガイドの代替品ではありません。これらのスキルは、基本的なものではなく、AI エージェントが失敗するギャップを対象としています。
宝石ではありません。ここには、Rails アプリに必要なものは何もありません。スキルは、AI コーディング エージェントによって消費されるマークダウン ファイルです。
これをプロジェクト固有のスキルと併用する
Rails-skills はベースラインであり、製品を理解するスキルに代わるものではありません。
アプリにすでにローカル スキル エンコードがある場合

独自のルール (テナントの分離、ドメインのワークフロー、支払い/予約ロジック、テストと検証ゲート、支店の規律) を設定し、それらが真実の情報源であり続ける必要があります。一般的な Rails パックでは、リリースに何が重要であるかを知ることができず、より鋭いローカル命令と競合する広範なスキル層は望ましくありません。
グリーンフィールド / ローカル スキルがまだない → パック全体をインストールします。それは強力なデフォルトです。
独自のスキルを備えた成熟したアプリ → ローカル スキルの信頼性を維持します。このパックから真のギャップがある場合にのみ厳選してください (移行の安全性、セキュリティレビューのチェックリスト、ホットワイヤー規約、バグ→リグレッションの規律は一般的に解除されるものです)。
衝突？プロジェクトの慣例が優先されます。 RSpec + 厳密なテストファーストではなく Minitest + システム/E2E 検証を使用している場合は、独自のループに従ってください。ここでのパターンは転送され、セレモニーはオプションです。
これは上級 Rails 開発者のデフォルトの本能であり、全面的に採用しなければならないプロセスではないと考えてください。
これらのスキルは DHH に基づいています: Majestic Monolith、ファット モデル、デフォルトの Hotwire、Solid Queue/Cache/Cable on Rails 8。Rails コミュニティが正当に同意しない場合、私たちは根拠を示して自分たちの立場を述べ、代替案を認めます。私たちはすべての決定において双方の立場に立つわけではありません。
PLAN.md のすべてのスキルが必要です

[切り捨てられた]

## Original Extract

Contribute to sandeepmvl/rails-skills development by creating an account on GitHub.

GitHub - sandeepmvl/rails-skills · GitHub
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
sandeepmvl
/
rails-skills
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .github .github docs docs skills skills .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md HOW_TO_USE.md HOW_TO_USE.md LICENSE LICENSE PLAN.md PLAN.md README.md README.md View all files Repository files navigation
Production-grade Claude Skills for Ruby on Rails. Stop fighting your AI coding agent on Rails conventions — drop in rails-skills and Claude Code, Cursor, Codex, Gemini CLI, Antigravity, and Windsurf will write Rails code the way senior Rails developers actually write it.
You ask Claude Code or Cursor to "add a draft status with scheduled publishing to Post." It generates a migration, a callback, a controller update. It looks fine. Then you notice:
It added a before_save callback instead of using your existing PublishPostJob
It used a string column for status instead of the enum already on the model
It didn't add the partial index PostgreSQL needs for the new status
The test it wrote doesn't use your existing Post factory
Every Rails dev using an AI agent has had this morning. The agent isn't dumb — it just doesn't know your Rails conventions. rails-skills teaches it.
A pack of independent Claude Skills ( open standard , released by Anthropic in December 2025) covering the patterns, performance traps, security baselines, and deployment conventions that senior Rails developers care about. Each skill is portable across:
Work across many projects? Install once, globally instead of per-repo. The Rails skills are description-gated — they stay dormant on non-Rails projects (React, Go, Python…) and wake up only for Rails work. See Two install modes .
cd < your-rails-app >
git clone https://github.com/sandeepmvl/rails-skills .claude/skills
Global — all projects (recommended for multi-project devs):
git clone https://github.com/sandeepmvl/rails-skills ~ /.claude/skills/rails-skills
Then start Claude Code. The orchestrator skill ( rails-project-discovery ) interviews you and routes to the right downstream skills automatically — and only triggers when you're actually working on Rails.
cd < your-rails-app >
git clone https://github.com/sandeepmvl/rails-skills .cursor/skills
Other tools
See docs/install.md for OpenAI Codex, Gemini CLI, Antigravity, and Windsurf instructions.
#
Skill
What it does
00
rails-project-discovery
Orchestrator — interviews you about your app and routes to the right skills below
01
activerecord-patterns
Associations, scopes, callbacks, includes vs preload vs eager_load
02
n-plus-one-killer
Detect, diagnose, and eliminate N+1 queries
03
service-objects-vs-fat-models
When to keep logic in the model and when to extract a service object
04
rspec-testing-pyramid
RSpec, FactoryBot, VCR, system specs — what to test at which layer
05
safe-migrations
strong_migrations rules, zero-downtime patterns, backfill in batches
06
rails-api-design
Versioning, serialization, pagination, rate limiting, JWT auth
07
solid-queue-and-sidekiq
Choosing between them, idempotent jobs, retries, scheduled work
08
devise-pundit-rodauth
Authentication + authorization patterns done the Rails way
09
kamal-docker-production
Multi-stage Dockerfile, docker-compose dev, Kamal 2 deploys
10
rails-security-baseline
Strong params, CSRF, Brakeman, secrets, JWT, CORS, OWASP for Rails
11
rails-caching-strategy
Solid Cache, Redis, fragment vs low-level, cache stampede prevention
12
hotwire-turbo-stimulus
Turbo Drive/Frames/Streams, Stimulus, Action Cable broadcasts
13
activestorage-uploads
Direct upload, variants, signed URLs, validation, S3/GCS service config
14
actionmailer-baseline
Mailer setup, deliver_later, previews, specs, bounces, idempotent sends
15
observability-baseline
lograge structured logs, Sentry, PII scrubbing, Rails.error.report
v0.2 — Expansion (24 skills)
#
Skill
What it does
16
rails-upgrade-7-to-8
dual-Gemfile next-version pattern; 7.x → 8 hops
17
rails-upgrade-6-to-7
importmap / jsbundling, encrypts, Trilogy intro
18
rails-upgrade-5-to-6
Zeitwerk transition, Webpacker decision
19
rails-upgrade-4-to-5
API mode, ActionCable, ApplicationRecord
20
rails-upgrade-3-to-4
strong params, Turbolinks, encrypted secrets
21
db-migration-postgres-mysql
Type mapping, JSONB→JSON, Trilogy gem swap
22
db-migration-mysql-postgres
UTF8MB4→UTF-8, prepared_statements + pgbouncer
23
db-migration-oracle-postgres
ora2pg, DATE→TIMESTAMPTZ, sequences
24
react-with-rails
Inertia.js default, vite_rails, API+SPA alt
25
vue-with-rails
Inertia + Vue 3 + Vite
26
angular-with-rails
Standalone components, @for / @if , provideHttpClient
27
puma-tuning-and-concurrency
Workers/threads, jemalloc, YJIT, PumaWorkerKiller
28
asset-pipeline-propshaft
Propshaft vs Sprockets, jsbundling, importmap
29
multi-database-and-replicas
connects_to, sticky writes, sharding caveats
30
webhook-handling
Idempotent inserts, HMAC verify, DLQ patterns
31
stripe-webhook-integration
construct_event, PaymentIntent, idempotency keys
32
external-api-integration
Faraday 2, Stoplight circuit breaker, VCR
33
feature-flagging
Flipper, gradual rollout, kill switch
34
form-objects-query-objects-presenters
Architectural extractions beyond fat models
35
actiontext-richtext
Trix, sanitization, embedded attachments, search index
36
i18n-and-timezones
rails-i18n, fallbacks, Time.current, DST traps
37
rails-search
pg_search → Meilisearch → Elasticsearch tiering
38
multi-tenancy
acts_as_tenant, subdomain resolver, require_tenant
39
console-safety-production
Sandbox, audit log, rake tasks over console
v0.3 — Specialization (18 skills)
#
Skill
What it does
40
when-NOT-to-use-microservices
Refuses microservices before guiding them. Modular monolith default.
41
microservices-decomposition
Bounded contexts, owned data, JWT/mTLS, sagas
42
monolith-to-services-extraction
Strangler fig, dark-launch, dual-write, cutover
43
kafka-rails
karafka 2.x, outbox, schema registry, partitioning
44
rabbitmq-rails
bunny + sneakers, DLX, prefetch, manual ack
45
redis-streams-rails
XADD, XREADGROUP, XAUTOCLAIM reaper
46
cdc-debezium-rails
Postgres logical replication, outbox + EventRouter SMT
47
event-driven-architecture
Domain vs integration events, rails_event_store, outbox
48
distributed-tracing-rails
OpenTelemetry, OTLP, baggage, sampling, exemplars
49
observability-rails-advanced
RED/USE, SLOs, multi-burn-rate alerts, runbooks
50
hipaa-rails
PHI handling, AR Encryption, audit logs, BAAs
51
pci-dss-rails
SAQ-A path, Stripe Elements, never PAN/CVV
52
gdpr-rails
DSAR, erasure, lawful basis, consent, retention
53
soc2-rails
Trust Services, audit logs, access reviews, MFA
54
data-warehouse-integration
Fivetran/Airbyte ingest, dbt transforms, reverse-ETL
55
ci-cd-github-actions-rails
matrix sharding, bundler-cache, OIDC, Kamal deploy
56
ci-cd-gitlab-rails
stages, services, id_tokens for OIDC, Auto DevOps caveat
57
ci-cd-jenkins-rails
Declarative pipeline, credentials store, when NOT Jenkins
Tooling & project-fit (2 skills)
#
Skill
What it does
58
rubocop-and-code-quality
RuboCop + omakase preset + rubocop-rails/-performance/-rspec, SimpleCov, erb_lint, RBS vs Sorbet
59
scaffold-project-skills
Interview your app + generate project-specific local skills (tenancy, domain workflows, your verify loop) — the product-aware layer the generic pack can't ship
See PLAN.md for the full roadmap and rationale.
Same prompt — "list the 20 most recent posts with their authors" — to the same agent. Without the pack, then with it.
Without rails-skills (AI-typical — a hidden N+1 that only bites in production):
# app/controllers/posts_controller.rb
def index
@posts = Post . order ( created_at : :desc ) . limit ( 20 )
end
<% @posts.each do |post| %>
<%= post.title %> — <%= post.author.name %> <%# 1 + 20 queries %>
<% end %>
With rails-skills (the n-plus-one-killer + activerecord-patterns skills trigger):
# app/controllers/posts_controller.rb
def index
# includes(:author) → 2 queries total, regardless of row count.
@posts = Post . includes ( :author ) . order ( created_at : :desc ) . limit ( 20 )
end
The agent also adds a Bullet config to catch the next N+1 in development before it ships, and explains why includes beats a manual join here. That's the difference: not just correct code, but conventional code with the reasoning a senior Rails dev would give.
A 30-second screencast of this exact flow lands here before the public launch. Until then, the worked example above is the demo.
Not magic. A rails-upgrade-4-to-8 skill will guide and accelerate the upgrade. It will not do it autonomously. Rails upgrades require human judgment on gem replacements and test fixes.
Not a replacement for the Rails Guides. These skills target the gaps where AI agents go wrong, not the basics.
Not a gem. Nothing here gets require d into your Rails app. Skills are markdown files consumed by AI coding agents.
Using this alongside your project-specific skills
rails-skills is a baseline, not a replacement for skills that know your product.
If your app already has local skills encoding its own rules — tenant isolation, domain workflows, payment/booking logic, your test + verification gates, branch discipline — those should stay the source of truth. A general Rails pack can't know what matters for your launch, and you don't want a broad skill layer competing with sharper local instructions.
Greenfield / no local skills yet → install the whole pack; it's a strong default.
Mature app with its own skills → keep your local skills authoritative. Cherry-pick from this pack only where you have a genuine gap (migration safety, security-review checklists, Hotwire conventions, bug→regression discipline are common ones to lift).
Conflicts? Your project conventions win. If you're on Minitest + system/E2E verification rather than RSpec + strict test-first, follow your own loop — the patterns here transfer, the ceremony is optional.
Think of it as a senior Rails dev's default instincts, not a process you must adopt wholesale.
These skills lean DHH: Majestic Monolith, fat models, Hotwire by default, Solid Queue/Cache/Cable on Rails 8. Where the Rails community legitimately disagrees, we state our position with rationale and acknowledge the alternative. We don't both-sides every decision.
We need skills for everything in PLAN.md unde

[truncated]
