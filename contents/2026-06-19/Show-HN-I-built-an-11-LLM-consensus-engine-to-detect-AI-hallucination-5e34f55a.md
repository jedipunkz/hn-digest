---
source: "https://github.com/jaquelinejaque/quorum-saas-starter"
hn_url: "https://news.ycombinator.com/item?id=48596771"
title: "Show HN: I built an 11-LLM consensus engine to detect AI hallucination"
article_title: "GitHub - jaquelinejaque/quorum-saas-starter: The only Next.js SaaS boilerplate with 11 LLM providers in semantic consensus, HSP fail-closed gate (Patent Pending), and EU AI Act audit-grade compliance. Production code from api.quorum-ai.dev. · GitHub"
author: "jaquelinejaque"
captured_at: "2026-06-19T10:35:00Z"
capture_tool: "hn-digest"
hn_id: 48596771
score: 3
comments: 0
posted_at: "2026-06-19T09:53:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I built an 11-LLM consensus engine to detect AI hallucination

- HN: [48596771](https://news.ycombinator.com/item?id=48596771)
- Source: [github.com](https://github.com/jaquelinejaque/quorum-saas-starter)
- Score: 3
- Comments: 0
- Posted: 2026-06-19T09:53:50Z

## Translation

タイトル: Show HN: AI 幻覚を検出する 11-LLM コンセンサス エンジンを構築しました
記事のタイトル: GitHub - jaquelinejaque/quorum-saas-starter: セマンティック コンセンサスに基づく 11 の LLM プロバイダー、HSP フェールクローズ ゲート (特許出願中)、および EU AI 法の監査グレードのコンプライアンスを備えた唯一の Next.js SaaS ボイラープレート。 api.quorum-ai.dev からの製品コード。 · GitHub
説明: セマンティック コンセンサス、HSP フェールクローズ ゲート (特許出願中)、および EU AI 法の監査グレードのコンプライアンスにおける 11 の LLM プロバイダーを備えた唯一の Next.js SaaS ボイラープレート。 api.quorum-ai.dev からの製品コード。 - jaquelinejaque/quorum-saas-starter

記事本文:
GitHub - jaquelinejaque/quorum-saas-starter: セマンティック コンセンサスの 11 社の LLM プロバイダー、HSP フェールクローズ ゲート (特許出願中)、および EU AI 法の監査グレードのコンプライアンスを備えた唯一の Next.js SaaS ボイラープレート。 api.quorum-ai.dev からの製品コード。 · GitHub
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
{{ メートル

メッセージ }}
ジャクリーンジャク
/
クォーラム-saas-スターター
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ジャクリーンジャク/quorum-saas-starter
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット docs docs lib lib scripts scripts templates templates .gitignore .gitignore DEMO_VIDEO_SCRIPT.md DEMO_VIDEO_SCRIPT.md GUMROAD_LISTING.md GUMROAD_LISTING.md LICENSE.md LICENSE.md README.md README.mdデモ.gif デモ.gif デモ.mp4デモ.mp4 デモ.テープ デモ.テープ パッケージ.json パッケージ.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
セマンティック コンセンサス、EU AI 法の監査グレードの準拠、およびすぐに使用できる 13 の自己進化ループに基づく 14 の LLM プロバイダーが同梱されている、実稼働対応の唯一のボイラープレートです。 api.quorum-ai.dev を動作させるのと同じコードに基づいて構築されています。
市場にあるすべての AI SaaS ボイラープレートには、1 つの LLM プロバイダー (通常は OpenAI、場合によっては Anthropic) が付属しています。それは簡単です。難しいのはその周りのすべてです。
爆発的なコストを発生させずに複数のプロバイダー間でクエリをルーティングする
単一のモデルを信頼するのではなく、意味論的な合意を通じて幻覚を検出する
EU AI法第12条/第13条のレビューに合格した監査証跡の生成
リスクの高い自律的なアクションのためのフェイルクローズされたゲートの構築
ユーザーごとのメモリ、RLHF フィードバック、および適応型ルーティング データを分離した状態に保つ
このキットには、実稼働環境ですでに配線、展開、実戦テスト済みのものすべてが含まれています。あなたは、他の人が立ち寄るオーケストレーション層の再構築ではなく、自分の製品を何が他と違うものにするかにエンジニアリング時間を費やします。
すぐに使える 14 のプロバイダー統合: OpenAI、Anthropic、Google Gemini、xAI Grok、Mistral、Cohere、NVIDIA、DeepSeek、Replicate、DashScope/Qwen、Zhipu/GLM、Moonshot/Kimi、Hermes (Nous)

ローカル、ラマ オラマ経由のローカル
埋め込みのコサイン類似度による意味的一致スコアリング — 語彙の重複や多数決ではありません
すべての回答で返される不一致の追跡 - ユーザーは、どのモデルが同意し、どのモデルが反対し、その理由を正確に確認できます。
クエリごとのコスト ログにより、API トークンの各コンセンサスにかかるコストをクライアントに示すことができます
自己進化ループ (合計 13)
ループ
学習内容
RLHFトラッカー
サムアップ/ダウンによるユーザーごと、クエリクラスごとの重み更新
ELOコンテスト
ペアワイズ モデルの勝利 → クエリ クラスごとのグローバル ランキング
ヘビアン共活性化
どのモデルが互いに一致するか (冗長なファンアウトを削減)
MoEルーター
毎回すべてではなく、クエリごとにプロバイダーの適切なサブセットを選択します
ヘビアンメモリー
セマンティックリコールを備えたユーザーごとのベクトルメモリ
遺伝的即時進化
評価セットで最も高いスコアを持つシステム プロンプトを変更して選択します
敵対的なループ
レッドチームは、ユーザーが脱獄する前に脱獄を発見するよう促す
A/B テスト
統計的に既存政策を上回る場合にのみ、新しい政策を推進する
蒸留
社内データに基づいて微調整された Llama チェックポイントを促進します
合成データ
実際の運用トラフィックからトレーニング ペアを生成します
アーキテクチャの検索
世代を超えてループトポロジを進化させる
フェデレーションラーニング
生データを共有せずにテナント全体で更新を集約します
ウェブ学習者
4 ソースの Web 取り込み (DDG + Wikipedia + HackerNews + arXiv) をベクター KB に取り込む
コンプライアンスと安全性 (すべてオプトイン)
HSP フェールクローズ ゲート (オプトイン) — 高リスク関数を requireHspApproval() でラップすると、人間が承認した Webhook がサインオフしない限り拒否されます。コンプライアンスは必要ありませんか？包まないでください。キットはこれを強制しません。依存関係によってラップされている場合でも、HSP_ENABLED=false を使用してグローバルにオフに切り替えます。
EU AI 法第 12 条監査証跡 (オプトイン) — ビジネス ロジックから writeAuditLine() を呼び出して、

g 決定。監査証明書ジェネレーターがこのログを読み取ります
EU AI 法第 13 条監査証明書 (オプトイン) — SHA-256 ハッシュチェーン PDF ジェネレーター、監査期間ごとに 1 回の呼び出し:generateAuditCert({...}) 。規制対象業種 (法律、フィンテック、医療) に販売する場合に役立ちます。それ以外の場合は無視できます。
ホスト型とローカルのポスチャ分割 — ホスト型デプロイメントはインフラマーカー検出によってフェールクローズされます (Cloud Run、k8s、Lambda)。 HSP_GATE_DEV_MODE=1 によるローカルリサーチモードのオプトイン
CUSTOMER_KEYS_ENCRYPTION_KEY — Fernet で暗号化された BYOK プロバイダー キー。ログには記録されず、漏洩もありません。エンドユーザーに独自の API キーを提供させる場合にのみ関係します。
認証 — 電子メール マジック リンク + OAuth 対応 (Apple、Google)、セッション Cookie、API 用の JWT
Stripe 請求 — Pro レベルのサブスクリプション、無料レベルのメータリング、署名検証付き Webhook ハンドラー (Stripe Event SDK)
電子メールの再送信 — トランザクション テンプレート: ようこそ、請求、監査アラート
Firestore の永続性 — API キー、顧客層、使用状況カウンター、暗号化された BYOK キー
Cloud Run デプロイ - 単一コマンド、0→N 自動スケーリング、リージョン フェイルオーバー対応
GitHub Webhook レシーバー — HMAC SHA-256 署名が検証され、プッシュ時に監査ジョブをディスパッチする準備ができています
Cloud Scheduler ターゲット — 共有シークレットによってゲートされる cron エンドポイント、夜間のバッチ ジョブの準備ができています
git clone https://github.com/jaquelinejaque/quorum-saas-starter
cd クォーラム-saas-starter
./scripts/setup.sh # インタラクティブ: キーの収集、Stripe プロダクトの作成、Cloud Run の構成
./scripts/deploy.sh # gcloud rundeploy + Firestore init + Stripe Webhook 登録
30 分以内に本番環境に出荷します。
このキットにはオーケストレーション レイヤーが含まれています。垂直方向を追加します。
AI Tax Assistant — プロンプト + このキットのコンセンサス = 製品
AI Code Reviewer — 評価 + このキットの RLHF = 製品
AIコンプライアンス監査人

— あなたのドメイン知識 + このキットの HSP ゲート = 製品
AI Legal Research — 訴訟データベース + このキットの Web 学習者 = 製品
インフラストラクチャのハード作業は完了しました。あなたは、顧客、プロンプト、そしてあなただけが持っている特定の分野のデータに焦点を当てます。
単一プロジェクトの展開。修正された Apache-2.0 (HSP 商用制限) に基づくソース コード。 3ヶ月ぶりの更新。サポートはありません。
事前トレーニングされた MoE ルーターの重み (10,000 を超える実稼働クエリがすでにポリシー テーブルを形成)
遺伝子プロンプト進化キット + 一般的な SaaS カテゴリにわたる 50 の進化したシード プロンプト
毎月録画された「オフィスアワー」ビデオ (ライブ通話なし - 録画のみ、すべての Pro+ 購入者がアクセス可能)
商用/ホワイトラベルライセンス — 独自の製品として再販
EU AI 法第 12/13 条監査キット (フォーム、テンプレート、証明書ジェネレーター)
オフィスアワーのアーカイブ (過去と未来のすべて)
すべての層: ソース コードのみ、マネージド ホスティング、ライブ サポートなし。使用するために許可を求めなければならないものではなく、構築できるものを購入してください。
このキットのオーケストレーション レイヤーは、本番環境で安定するまでに 18 か月と 60 以上のコミットを要しました。監査グレードのコンプライアンス ワークフローで使用される Quorum API を実行します。価格は上級エンジニアがゼロから再構築する場合にかかる費用を反映しています。標準的な請負業者料金 (約 25,000 ポンド) でフルタイムで約 6 週間の作業が必要です。 497ポンドから2,497ポンドで、その時点で買い戻すことになります。
マネージド ホスティングが必要な場合は、クォーラム ホスティング (月額 49 ポンド) を使用してください。
ライブサポートが必要な場合 - これはコードのみのライセンスです
AI アプリをバイブコーディングして 2 時間以内に出荷したい場合は、代わりに Lovable または Vercel テンプレートを使用してください
あなたはまだ実際の製品のアイデアを持っていません - このキットは学習用ではなく出荷用です
キットが環境に正常に展開されなかった場合、問答無用で 30 日間経過します。導入が成功した後は返金はありません。ソース コードは手元にあります。
ジャクリーン M によって構築されました

artins / Sovereign Chain Ltd. 特許出願中: HSP プロトコル (PCT/US26/11908)。
セマンティックコンセンサスの 11 社の LLM プロバイダー、HSP フェールクローズ ゲート (特許出願中)、および EU AI 法の監査グレードのコンプライアンスを備えた唯一の Next.js SaaS ボイラープレート。 api.quorum-ai.dev からの製品コード。
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

The only Next.js SaaS boilerplate with 11 LLM providers in semantic consensus, HSP fail-closed gate (Patent Pending), and EU AI Act audit-grade compliance. Production code from api.quorum-ai.dev. - jaquelinejaque/quorum-saas-starter

GitHub - jaquelinejaque/quorum-saas-starter: The only Next.js SaaS boilerplate with 11 LLM providers in semantic consensus, HSP fail-closed gate (Patent Pending), and EU AI Act audit-grade compliance. Production code from api.quorum-ai.dev. · GitHub
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
jaquelinejaque
/
quorum-saas-starter
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
jaquelinejaque/quorum-saas-starter
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits docs docs lib lib scripts scripts templates templates .gitignore .gitignore DEMO_VIDEO_SCRIPT.md DEMO_VIDEO_SCRIPT.md GUMROAD_LISTING.md GUMROAD_LISTING.md LICENSE.md LICENSE.md README.md README.md demo.gif demo.gif demo.mp4 demo.mp4 demo.tape demo.tape package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
The only production-ready boilerplate that ships with 14 LLM providers in semantic consensus , EU AI Act audit-grade compliance , and 13 self-evolution loops out of the box. Built on the same code that powers api.quorum-ai.dev .
Every AI SaaS boilerplate on the market ships with one LLM provider — usually OpenAI, sometimes Anthropic. That's the easy part. The hard part is everything around it:
Routing queries across multiple providers without exploding costs
Detecting hallucination via semantic consensus instead of trusting a single model
Generating audit trails that pass EU AI Act Article 12/13 review
Building fail-closed gates for high-risk autonomous actions
Keeping per-user memory, RLHF feedback, and adaptive routing data isolated
This kit gives you all of that already wired, deployed, and battle-tested in production. You spend your engineering time on what makes your product different — not on rebuilding the orchestration layer everyone else stops at.
14 provider integrations out of the box: OpenAI, Anthropic, Google Gemini, xAI Grok, Mistral, Cohere, NVIDIA, DeepSeek, Replicate, DashScope/Qwen, Zhipu/GLM, Moonshot/Kimi, Hermes (Nous) local, Llama local via Ollama
Semantic agreement scoring via cosine similarity on embeddings — not lexical overlap, not majority vote
Disagreement trace returned with every answer — your users see exactly which models agreed, which dissented, and why
Per-query cost log so you can show clients what each consensus cost in API tokens
Self-evolution loops (13 total)
Loop
What it learns
RLHF tracker
Per-user, per-query-class weight updates from thumbs-up/down
ELO competition
Pairwise model wins → global ranking per query class
Hebbian co-activation
Which models agree with each other (reduces redundant fan-out)
MoE router
Picks the right subset of providers per query, not all every time
Hebbian memory
Per-user vector memory with semantic recall
Genetic prompt evolution
Mutates and selects system prompts that score highest on your eval set
Adversarial loop
Red-team prompts to catch jailbreaks before users do
A/B testing
Promotes new policies only when they beat the incumbent statistically
Distillation
Promotes Llama checkpoints fine-tuned on your in-house data
Synthetic data
Generates training pairs from real production traffic
Architecture search
Evolves loop topology over generations
Federated learning
Aggregates updates across tenants without sharing raw data
Web learner
4-source web ingestion (DDG + Wikipedia + HackerNews + arXiv) into a vector KB
Compliance & safety (all opt-in)
HSP fail-closed gate (opt-in) — wrap any high-risk function with requiresHspApproval() and it refuses unless a human-approved webhook signs off. Don't need compliance? Don't wrap. The kit doesn't force this on you. Toggle off globally with HSP_ENABLED=false even if a dependency does wrap it.
EU AI Act Article 12 audit trail (opt-in) — call writeAuditLine() from your business logic to log decisions; the audit cert generator reads this log
EU AI Act Article 13 audit certificate (opt-in) — SHA-256 hash-chained PDF generator, one call per audit period: generateAuditCert({...}) . Useful when selling to regulated verticals (legal, fintech, health); ignorable otherwise.
Hosted vs local posture split — hosted deployment is fail-closed by infra-marker detection (Cloud Run, k8s, Lambda); local research mode opt-in via HSP_GATE_DEV_MODE=1
CUSTOMER_KEYS_ENCRYPTION_KEY — Fernet-encrypted BYOK provider keys, never logged, never leaked. Only relevant if you're letting your end users supply their own API keys.
Auth — email magic links + OAuth-ready (Apple, Google), session cookies, JWT for API
Stripe billing — Pro tier subscription, Free tier metering, webhook handler with signature verification (Stripe Event SDK)
Resend email — transactional templates: welcome, billing, audit alerts
Firestore persistence — API keys, customer tiers, usage counters, encrypted BYOK keys
Cloud Run deploy — single command, autoscaling 0→N, regional failover ready
GitHub webhook receiver — HMAC SHA-256 signature verified, ready to dispatch audit jobs on push
Cloud Scheduler target — cron endpoint gated by shared secret, ready for nightly batch jobs
git clone https://github.com/jaquelinejaque/quorum-saas-starter
cd quorum-saas-starter
./scripts/setup.sh # interactive: collects keys, creates Stripe products, configures Cloud Run
./scripts/deploy.sh # gcloud run deploy + Firestore init + Stripe webhook registration
You ship to production in under 30 minutes.
The kit gives you the orchestration layer. You add the vertical:
AI Tax Assistant — your prompts + this kit's consensus = product
AI Code Reviewer — your evals + this kit's RLHF = product
AI Compliance Auditor — your domain knowledge + this kit's HSP gate = product
AI Legal Research — your case database + this kit's web learner = product
The hard infrastructure work is done. You focus on customers, prompts, and the niche-specific data only you have.
Single project deployment. Source code under modified Apache-2.0 (HSP commercial restriction). 3 months of updates. No support.
Pre-trained MoE router weights (10,000+ production queries already shaped the policy table)
Genetic prompt evolution kit + 50 evolved seed prompts across common SaaS categories
Monthly recorded "office hours" video (no live calls — recordings only, accessible to all Pro+ buyers)
Commercial / white-label license — resell as your own product
EU AI Act Article 12/13 audit kit (forms, templates, cert generator)
Office hours archive (all past + future)
All tiers: source code only, no managed hosting, no live support. Buy what you can build on, not what you have to ask permission to use.
The orchestration layer in this kit took 18 months and 60+ commits to get production-stable. It runs the Quorum API used in audit-grade compliance workflows. Pricing reflects what it would cost a senior engineer to rebuild from scratch: roughly 6 weeks of full-time work at standard contractor rates (~£25,000). At £497–£2,497 you're buying back that time.
You want managed hosting — go use Quorum hosted (£49/mo)
You want live support — this is code-only license
You want to vibe-code an AI app and ship in 2 hours — use Lovable or Vercel templates instead
You don't have a real product idea yet — this kit is for shipping, not learning
30 days, no questions asked, if the kit doesn't deploy successfully on your environment. After successful deployment, no refunds — you have the source code.
Built by Jaqueline Martins / Sovereign Chain Ltd. Patent Pending: HSP Protocol (PCT/US26/11908).
The only Next.js SaaS boilerplate with 11 LLM providers in semantic consensus, HSP fail-closed gate (Patent Pending), and EU AI Act audit-grade compliance. Production code from api.quorum-ai.dev.
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
