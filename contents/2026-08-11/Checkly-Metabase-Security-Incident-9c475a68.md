---
source: "https://www.checklyhq.com/blog/metabase-security-incident/"
hn_url: "https://news.ycombinator.com/item?id=49253669"
title: "Checkly: Metabase Security Incident"
article_title: "Metabase Security Incident"
author: "sangeeth96"
captured_at: "2026-08-11T05:52:36Z"
capture_tool: "hn-digest"
hn_id: 49253669
score: 1
comments: 0
posted_at: "2026-08-11T05:16:54Z"
tags:
  - hacker-news
  - translated
---

# Checkly: Metabase Security Incident

- HN: [49253669](https://news.ycombinator.com/item?id=49253669)
- Source: [www.checklyhq.com](https://www.checklyhq.com/blog/metabase-security-incident/)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T05:16:54Z

## Translation

タイトル: Checkly: メタベース セキュリティ インシデント
記事のタイトル: メタベース セキュリティ インシデント
説明: 2026 年 8 月 3 日、攻撃者は、当社が社内で使用しているサードパーティ分析ツールである Metabase のゼロデイ脆弱性を悪用し、…への読み取りアクセスを取得しました。

記事本文:
メタベースのセキュリティ インシデント Checkly - ホーム製品の検出
稼働時間の監視 デジタル フットプリントの可用性を測定する ハートビートの監視 サイレントに失敗する cron ジョブとバックアップを検出する 合成監視 スタック全体で実際のユーザー インタラクションをシミュレートする テスト AI を活用したテスト レポーターで本番前に問題を検出する コミュニケーション
ステータス ページ アプリの可用性を顧客に通知 アラート チームにすぐに通知するコンテキスト アラート 解決
AI 根本原因分析 AI エージェントによる自動根本原因分析 トレース より深い洞察のための強力な OTel トレース
開発者 アプリとともにデプロイされたリポジトリ内の TypeScript をチェックインする SRE およびプラットフォーム 可観測性スタックの合成レイヤー QA エンジニア Playwright スイートを運用モニターとして実行する エンジニアリング マネージャー すべてのチームに独自のモニターの所有権を付与する ユースケース
重要なユーザー フロー 実際のブラウザからログイン、チェックアウト、サインアップを監視 API とバックエンドの監視 スケジュールに従ってエンドポイント、チェーン、認証を検証 実稼働モニターへのテスト タグ付き仕様を CI からスケジュールされたチェックに昇格 可観測性の統合 APM に固定された合成モジュールを置き換える AI 生成コードの信頼性 ユーザーが実行する前にエージェントが出荷するものを確認する
電子商取引 チェックアウトと収益経路を 24 時間保護 金融サービス 厳格な SLA に基づいて稼働時間と遅延を証明 SaaS および B2B ソフトウェア すべてのテナントのコア ワークフローを検証し続ける 別のツールから移行しますか?
ドキュメント 主要な概念と機能を学ぶ API ドキュメント Checkly REST API CLI ドキュメントで構築 モニタリングを端末に導入する クイックスタート 最初のチェックを迅速に設定する ガイド Playwright および OTel の詳細ガイド MCP サーバー Checkly を AI ツールに接続する リファレンス
エージェントスキルテラフォ

rm Pulumi すべての統合変更ログの注目のリソース
ブログ Checkly で最新ニュースを読む Playwright などを学ぶためのヒントとベスト プラクティス ウェビナー オンデマンドでウェビナーに登録または表示する イベント チームと直接会う コミュニティ
パブリック ロードマップ パブリック ロードマップの機能を読んだり、投票したり、追加したりする コミュニティ Slack Checkly コミュニティと接続する お客様 価格 ログイン 無料で始める オープン ナビゲーション 製品 ソリューション 開発者向けリソース お客様 価格 無料で始める ログイン ブログ / セキュリティ メタベース セキュリティ インシデント
2026 年 8 月 3 日、攻撃者は、当社が社内で使用しているサードパーティ分析ツールである Metabase のゼロデイ脆弱性を悪用し、Checkly 運用データのコピーを保持するデータベースへの読み取りアクセスを取得しました。攻撃者は認証をバイパスし、Metabase Cloud インスタンス上の管理者セッションを取得しました。それ以来、Metabase は攻撃をブロックし、脆弱性にパッチを当て、セキュリティ アップデートを公開しました。攻撃者は当社の実稼働プラットフォームにアクセスしませんでした。
認証値をチェック構成に直接入力する場合、またはトレース収集に OTEL API キーを使用する場合、または暗号化された状態でアクセスされる可能性のあるデータに対してさらに安全性を高めたい場合は、今すぐ資格情報をローテーションしてください。
この投稿の残りの部分では、何が暴露され、何が暴露されなかったのか、そしてそれに対して私たちが何をしているのかについて説明します。
このようなことが起こって申し訳ありません。この脆弱性はベンダーの製品にありましたが、お客様のデータを保護するのが私たちの仕事であり、今回のインシデントによりその一部が危険にさらされました。
Metabase は、製品の使用状況と運用指標を分析するために使用するビジネス インテリジェンス ツールです。これは、運用プラットフォームとは別に、Checkly 運用データのコピーを保持するデータ ウェアハウスに接続します。
ゼロデイにより、攻撃者は有効な権限を持たずに管理者セッションを作成することができました。

資格。パスワードは盗まれず、Checkly システムが直接侵害されることもありませんでした。このセッションでは、攻撃者は 2026 年 8 月 3 日に約 26 分間、メタベースを介して接続されたデータ ウェアハウスにクエリを実行しました。すべてのアクティビティは読み取り専用でした。
Metabase は同日、Metabase Cloud に対する攻撃を発見し、攻撃者が使用したエンドポイントをブロックし、脆弱性を修正しました。彼らは 8 月 6 日に顧客と法執行機関に通知し、サードパーティのフォレンジック会社と協力しました。 8 月 10 日、影響を受けるシステムの独自の分析を完了するために Metabase が提供した詳細なログを使用して、インスタンスにおける攻撃者のアクティビティの調査を終了しました。
データ ウェアハウスには小切手の構成データとアカウント データが含まれており、小切手で送信する場合は認証キーやヘッダーが含まれる可能性があります。具体的には:
Checkly シークレットとして保存した認証情報は保存時に暗号化され、実行時にのみ挿入されます。これらは平文では公開されませんでした。
Checkly は、チェックの構成に直接追加されたカスタム値 (カスタム ヘッダー、Cookie、クエリ文字列パラメータ、または認証ヘッダー) をプレーンテキストで保存します。これらの値は公開されていると考える必要があります。
トレース収集に使用される OTEL API キーの暗号化ハッシュにもアクセスされました。
アクセスは読み取り専用で、8 月 3 日には約 26 分間続きました。攻撃者は、チェック、アラート、または設定を変更、無効化、作成していません。
攻撃者はアカウント、ユーザー、API キーに変更を加えておらず、新しいキーも作成していません。
攻撃者は当社の実稼働プラットフォームにアクセスしませんでした。運用データベースも Checkly プラットフォームも、チェック実行インフラストラクチャとアラートを含めて、侵害されたことはありません。
クラウド監査ログを確認したところ、既知の分析ツールの外部から誰かがデータ ウェアハウスにアクセスした形跡は見つかりませんでした。
回転させます

チェック構成 (カスタム ヘッダー、Cookie、クエリ文字列パラメーター、または承認ヘッダー) に直接追加される資格情報とシークレット。
使用中のすべての OTEL API キーをローテーションします。
2026 年 8 月 3 日以降、これらの資格情報が不正使用から保護しているシステムを確認します。
Vercel、Prometheus、Coralogix の統合を再認証する
チェックが認証に Checkly シークレットのみを使用する場合は、シークレットをローテーションする必要はありません。将来のセキュリティを確保するために、シークレットと機密データを Checkly Secrets に保存することをお勧めします。
Metabase がデータ ウェアハウスにアクセスするために使用する認証情報をローテーションしました。
すべての Metabase ユーザー パスワードをリセットし、侵害された管理者アカウントを削除し、すべての Metabase API キーをローテーションしました。
インスタンスの管理者アカウントを監査したところ、攻撃者が何も作成していないことが確認されました。
クラウド監査ログを通じて、既知の分析ツールのみがウェアハウスの認証情報を使用していることを確認しました。
影響を受けるお客様には直接ご連絡させていただいております。
資格情報をローテーションすると、当面の問題が解決されます。これは問題の原因を解決するものではありません。分析環境には機密データが多く含まれており、必要以上に広範なアクセスが行われていました。そこで、次のような変更を加えています。
分析および処理ツールを広範に監査して、潜在的に機密データの露出面を制限します。
データ アクセスの範囲を絞り、必要な情報のみにアクセスすることで、分析処理用のチェック構成とデータを保存する際のサニタイズ プロセスを改善します。
ベンダーのセキュリティ通知は、当社のオンコール エンジニアに直接連絡します。
潜在的に機密データがチェック構成で公開された場合に、顧客に積極的に通信することを中心とした将来の製品機能。
Metabase による進行中のフォレンジック調査または独自の調査により、上記のガイダンスを変更するものが判明した場合は、この投稿を更新します。
あなたが持っているなら

問題が発生した場合、または小切手が接触したシステム上で不審なアクティビティが発生した場合は、support@checklyhq.com までご連絡ください。
次の場合、アカウントの機密データが漏洩する可能性があります。
API チェックのカスタム リクエスト ヘッダー、クエリ文字列パラメータ、認証ヘッダー/Cookie などのチェック設定、およびブラウザ チェックとマルチステップ チェック スクリプトで機密情報をハードコーディングしました。
OpenTelemetry と Checkly の統合をセットアップします。トレース収集に使用される OTEL API キーの暗号化ハッシュがアクセスされました。
チェック構成 (カスタム ヘッダー、Cookie、クエリ文字列パラメーター、または承認ヘッダー) に直接追加された資格情報とシークレットをローテーションします。
使用中のすべての OTEL API キーをローテーションします。
2026 年 8 月 3 日以降、これらの資格情報が不正使用から保護しているシステムを確認します。
統合の再認証 (Vercel、Prometheus、Coralogix などを含む)
サービス API キー、ユーザー API キー、またはプライベート ロケーション API キーが公開されましたか?
いいえ。ただし、安全のためにローテーションすることをお勧めします。
Checkly アカウントのパスワードを変更する必要がありますか?
チェックごとの環境変数とシークレットは公開されましたか?
チェックごとの環境変数が公開されました。これらは、チェック構成に直接作成して保存した環境変数です。
チェックごとの環境シークレットと「ロックされた」変数が暗号化された形式で公開されました。それらは平文では公開されませんでした。量子コンピューティングなどのより高度なメカニズムが利用可能になった後、安全を確保し、将来的に暗号化解除されないように保護するために、これらの認証情報をローテーションすることをお勧めします。
グローバル環境変数と秘密は漏洩しましたか?
グローバル環境変数とシークレット値は公開されませんでした。グローバル環境変数と秘密のメタデータ (名前、タイムスタンプなど) が小規模に漏洩した可能性があります。
いいえ、私たちは

フック URL、Slack トークン、PagerDuty キー、およびアラート電子メール アドレスは公開されませんでした。安全を期すために、これらの資格情報をローテーションすることをお勧めします。

## Original Extract

On 3 August 2026, an attacker exploited a zero-day vulnerability in Metabase, the third-party analytics tool we use internally, and gained read access to a…

Metabase Security Incident Checkly - Home Product Detect
Uptime Monitoring Measure the availability of your digital footprint Heartbeat Monitoring Catch cron jobs and backups that fail silently Synthetic Monitoring Simulate real user interactions across your stack Testing Catch issues before production with an AI-powered test reporter Communicate
Status Pages Communicate app availability to your customers Alerts Contextual alerting to notify the team right away Resolve
AI Root Cause Analysis Automated root cause analysis powered by AI agents Traces Powerful OTel tracing for deeper insights Getting Started
Developers Checks in TypeScript, in your repo, deployed with your app SRE & Platform The synthetic layer for your observability stack QA Engineers Run your Playwright suite as production monitors Engineering Managers Give every team ownership of its own monitors Use Cases
Critical user flows Watch login, checkout, and signup from real browsers API & backend monitoring Validate endpoints, chains, and auth on a schedule Tests to production monitors Promote tagged specs from CI to scheduled checks Observability consolidation Replace the synthetic module bolted onto your APM Reliability for AI-generated code Verify what your agents ship before your users do Industries
E-commerce Protect checkout and revenue paths around the clock Financial services Prove uptime and latency against strict SLAs SaaS & B2B software Keep every tenant's core workflows verified Moving from another tool?
Documentation Learn key concepts and features API Docs Build on the Checkly REST API CLI Docs Bring monitoring to your terminal Quickstart Set up your first check fast Guides In-depth Playwright & OTel guides MCP Server Connect Checkly to your AI tools Reference
Agent Skills Terraform Pulumi All Integrations Changelog Resources Featured
Blog Read about the latest news at Checkly Learn Tips and best practices for learning Playwright and more Webinars Register or view webinars on-demand Events Meet the team in person Community
Public Roadmap Read, vote, or add on features in a public roadmap Community Slack Connect with the Checkly Community Customers Pricing Login Start for free Open Navigation Product Solutions Developers Resources Customers Pricing Start for free Login Blog / Security Metabase Security Incident
On 3 August 2026, an attacker exploited a zero-day vulnerability in Metabase , the third-party analytics tool we use internally, and gained read access to a database holding a copy of Checkly operational data. The attacker bypassed authentication and obtained an administrator session on our Metabase Cloud instance. Metabase has since blocked the attack, patched the vulnerability, and published a security update . The attacker did not access our production platform.
If you put authentication values directly in your check configurations, or you use OTEL API keys for trace collection, or you want to be extra safe about potential data being accessed in an encrypted state, please rotate your credentials now.
The rest of this post explains what was exposed, what was not, and what we are doing about it.
We are sorry this happened. The vulnerability was in a vendor's product, but protecting your data is our job, and this incident put some of it at risk.
Metabase is a business intelligence tool we use to analyze product usage and operational metrics. It connects to a data warehouse that holds a copy of Checkly operational data, separate from our production platform.
The zero-day allowed the attacker to create an administrator session without valid credentials. No password was stolen, and no Checkly system was directly breached. With that session, the attacker queried the connected data warehouse through Metabase for about 26 minutes on August 3rd, 2026. All activity was read-only.
Metabase discovered the attack against Metabase Cloud the same day, blocked the endpoints the attacker used, and patched the vulnerability. They notified customers, law enforcement, and engaged a third-party forensics firm on August 6th. On August 10th, we concluded our investigation of the attacker's activity in our instance using detailed logs Metabase provided to complete our own analysis of the affected systems.
The data warehouse contained check configuration and account data, potentially including authorization keys or headers if your checks send them. Specifically:
Credentials you store as Checkly secrets are encrypted at rest and injected only at run time. These were not exposed in plaintext .
Checkly stores custom values added directly in a check's configuration (custom headers, cookies, query string parameters, or authorization headers) in plaintext. You should consider these values exposed.
Cryptographic hashes of OTEL API keys used for trace collection were also accessed.
Access was read-only and lasted about 26 minutes on 3 August. The attacker did not modify, disable, or create any checks, alerts, or settings.
The attacker made no changes to your account, users, or API keys, and created no new ones.
The attacker did not access our production platform; neither the production database nor the Checkly platform, including check execution infrastructure and alerting, were ever breached.
We reviewed our cloud audit logs and found no evidence of anyone reaching the data warehouse from outside our known analytics tooling.
Rotate any credentials and secrets added directly to a check configuration (custom headers, cookies, query string parameters, or authorization headers).
Rotate all OTEL API keys in use.
Review the systems those credentials protect for unauthorized use since 3 August 2026.
Reauthorize Vercel, Prometheus, and/or Coralogix integrations
If your checks only use Checkly secrets for authentication, you do not need to rotate them. For future security, we recommend storing secrets and sensitive data with Checkly Secrets .
We rotated the credentials Metabase uses to reach the data warehouse.
We reset all Metabase user passwords, removed the compromised administrator account, and rotated all Metabase API keys.
We audited the admin accounts on the instance and confirmed the attacker created none.
We verified through cloud audit logs that only our known analytics tooling used the warehouse credentials.
We are contacting affected customers directly.
Rotating credentials fixes the immediate problem. It does not fix the reason this hurt: our analytics environment held more sensitive data and had broader access than it needed. So we are making these changes:
Extensive audit of analytics and processing tools to limit the exposure surface of any potentially sensitive data.
Improve our sanitization process when storing check configurations and data for analytic processing by scoping down data access so only required information is accessed.
Vendor security notices will page our on-call engineers directly.
Future product features around proactively communicating to customers when potentially sensitive data is exposed in their check configurations.
We will update this post if Metabase's ongoing forensic investigation or our own review turns up anything that changes the guidance above.
If you have questions or see suspicious activity on systems your checks touch, contact us at support@checklyhq.com.
Sensitive data from your account was potentially exposed if:
You hard-coded sensitive information in check configurations, like API check custom request headers, query string parameters, authorization headers / cookies, and in Browser check and Multistep check scripts.
You set up an OpenTelemetry integration with Checkly. Cryptographic hashes of OTEL API keys used for trace collection were accessed.
Rotate any credentials and secrets added directly to a check configuration (custom headers, cookies, query string parameters, or authorization headers).
Rotate all OTEL API keys in use.
Review the systems those credentials protect for unauthorized use since 3 August 2026.
Reauthorize integrations (including Vercel, Prometheus, Coralogix, etc)
Were service API keys, user API keys, or private location API keys exposed?
No. However, we recommend you rotate them to be on the safe side.
Do I need to change my Checkly account password?
Were per-check environment variables and secrets exposed?
Per-check environment variables were exposed. These are environment variables that you created and saved directly in your check configuration.
Per-check environment secrets and ‘locked’ variables were exposed in their encrypted form. They were not exposed in plain text. We recommend rotating those credentials to be on the safe side and protect against decrypting them in the future, once more sophisticated mechanisms like quantum computing are available.
Were global environment variables and secrets exposed?
Global environment variable and secret values were not exposed. Global environment variable and secret metadata (name, timestamp, etc.) was potentially exposed at a small scale.
No. Webhook URLs, Slack tokens, PagerDuty keys, and alert email addresses were not exposed. We recommend rotating those credentials to be on the safe side.
