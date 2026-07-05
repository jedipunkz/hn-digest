---
source: "https://github.com/irenee28/akm-clr-governance-layer"
hn_url: "https://news.ycombinator.com/item?id=48798328"
title: "AKM-CLR – pre-inference governance for vLLM-style LLM serving"
article_title: "GitHub - irenee28/akm-clr-governance-layer · GitHub"
author: "akili2805"
captured_at: "2026-07-05T22:53:08Z"
capture_tool: "hn-digest"
hn_id: 48798328
score: 2
comments: 0
posted_at: "2026-07-05T21:58:11Z"
tags:
  - hacker-news
  - translated
---

# AKM-CLR – pre-inference governance for vLLM-style LLM serving

- HN: [48798328](https://news.ycombinator.com/item?id=48798328)
- Source: [github.com](https://github.com/irenee28/akm-clr-governance-layer)
- Score: 2
- Comments: 0
- Posted: 2026-07-05T21:58:11Z

## Translation

タイトル: AKM-CLR – vLLM スタイルの LLM サービスの事前推論ガバナンス
記事のタイトル: GitHub - irenee28/akm-clr-governance-layer · GitHub
説明: GitHub でアカウントを作成して、irenee28/akm-clr-governance-layer の開発に貢献します。

記事本文:
GitHub - irenee28/akm-clr-governance-layer · GitHub
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
アイレニー28
/
akm-clr-ガバナンス層
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
irenee28/akm-clr-ガバナンス層
メインブラ

nches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット docs docs 例 例 runnable_demo runnable_demo .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE.md LICENSE.md README.md README.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AKM-CLR — 共有 LLM サービス用のテナント/タスク ガバナンス層
AKM-CLR は、マルチテナント LLM インフラストラクチャ用の検証済みのプロトタイプ ガバナンス/制御レイヤーです。これは、vLLM などの OpenAI 互換のサービス エンジンの上に位置し、バックエンド推論の前にテナント リクエストを転送するか、フォールバックするか、ブロックするかを決定するのに役立ちます。
vLLM は高速にサービスを提供します。 AKM-CLR は、vLLM への到達を許可するものを制御します。
AKM-CLR は、共有 LLM バックエンド、RAG システム、LoRA アダプター、またはテナント固有のワークフローを実行する AI プラットフォーム向けに設計されています。
ドメインとアダプターのルーティング ポリシー
安全でない/テナント間リクエストに対する推論前のブロック
間違ったドメインまたは不正なリクエストに対する安全なフォールバック
vLLM/OpenAI 互換のバックエンド統合
最新の AI プラットフォームでは、共有インフラストラクチャを通じて多くのテナント、タスク、RAG ソース、アダプターにサービスを提供することが増えています。ガバナンス層がないと、バックエンドが技術的に正しく動作している場合でも、安全でないリクエストや間違ったドメインのリクエストがバックエンドに到達する可能性があります。
AKM-CLR は、アプリケーション ロジックとモデル提供の間のガバナンス ギャップに焦点を当てています。
クライアント/アプリ
↓
AKM-CLR ガバナンス ゲートウェイ
↓ 許可され安全な場合のみ
vLLM / OpenAI 互換バックエンド
↓
モデルの応答
検証の概要
AKM-CLR は対照実験で評価されました。
1 台の H200 に最大 29.5 GB の VRAM を割り当て
vLLM によって提供される Qwen2.5-7B-Instruct を使用した制御された比較では、次のようになります。
AKM-CLR は検証済みのプロトタイプです。以下のことは主張していません。
硬化せずに生産準備完了
サービスエンジンとして vLLM よりも優れている
クライ

m の方が狭くて強い:
AKM-CLR は、推論の前にテナント/タスク ガバナンスを追加することで vLLM を補完します。
以下を運営する少数のデザイン パートナーを探しています。
vLLM または OpenAI 互換の推論サービス
微調整された LoRA アダプター インフラストラクチャ
規制されたエンタープライズ AI システム
テスト バックエンドに対する AKM Gateway の統合
テナント/タスク ガバナンス テスト スイート
創設者: イレネー・アキリマリ
場所: コンゴ民主共和国キンシャサ
デザインパートナーのお問い合わせ：shukranimungu@gmail.com
最新の製品準備状況の検証
最初の研究検証の後、AKM-CLR はより製品指向のゲートウェイ プロトタイプに拡張され、3 つの追加実験にわたってテストされました。
実験 1 — GPU を使用しない製品の実証
ローカルの GPU を使用しない実験では、コードとしてのポリシーを使用して、AKM-CLR が構成可能なガバナンス レイヤーとして検証されました。
ガバナンスに安全な意思決定率: 100%
レート制限動作: 検証済み
実験 2 — 実際の vLLM 統合の実証
RunPod の real-vLLM 実験では、vLLM を通じて提供される Qwen/Qwen2.5-1.5B-Instruct 上の AKM-CLR を検証しました。
ガバナンスに安全な意思決定率: 100%
許可されたバックエンド呼び出し率: 100%
承認されたテナント要求が vLLM に到達しました。テナント間のカナリア プローブ、間違ったアダプター アクセス、間違った RAG ソース アクセス、無効な API キー、テナントのスプーフィングなど、安全でないリクエストや不正なリクエストは、バックエンド推論の前にブロックされました。
実験 3 — エンタープライズ展開のプレビュー
エンタープライズ展開プレビューでは、AKM-CLR パッケージがインフラストラクチャ コンポーネントとして検証されました。
分離された API キー シークレット YAML
トラフィックガバナンス/レート制限動作
これは、本番環境への準備やセキュリティの保証を主張するものではありません。これは、AKM-CLR が研究プロトタイプから、OpenAI 互換の LLM サービング スタック用の展開可​​能なガバナンス/アドミッション コントロール層に移行できることを検証します。
実験 5 — ワークスペース /

セッション/キャッシュ/コンテキスト境界ガード
AKM-CLR は、推論前にワークスペース、セッション、キャッシュ、および取得されたコンテキスト境界に対するガバナンスをテストするように拡張されました。
この実験では、リクエストが推論に達する前に、テナントが特定のワークスペース、セッション、キャッシュ エントリ、取得されたコンテキスト オブジェクト、アダプター、および RAG ソースをアタッチする権限を持っているかどうかを検証しました。
ガバナンスに安全な意思決定率: 100%
予想される決定一致率: 100%
ワークスペース/セッション/キャッシュ/コンテキスト境界安全率: 100%
信頼性フォールバック動作: 検証済み
出力カナリアブロック: 検証済み
これは、本番環境の準備が整っていることや、AKM-CLR が特定の外部インシデントを防止したことを主張するものではありません。ワークスペース、セッション、キャッシュ、コンテキスト漏洩を伴うマルチテナント AI 分離リスクに対する追加の制御ポイントを検証します。
実験 5 — ワークスペース / セッション / キャッシュ / コンテキスト境界ガード
AKM-CLR は、推論前にワークスペース、セッション、キャッシュ、および取得されたコンテキスト境界に対するガバナンスをテストするように拡張されました。
この実験では、リクエストが推論に達する前に、テナントが特定のワークスペース、セッション、キャッシュ エントリ、取得されたコンテキスト オブジェクト、アダプター、および RAG ソースをアタッチする権限を持っているかどうかを検証しました。
ガバナンスに安全な意思決定率: 100%
予想される決定一致率: 100%
ワークスペース/セッション/キャッシュ/コンテキスト境界安全率: 100%
信頼性フォールバック動作: 検証済み
出力カナリアブロック: 検証済み
これは、本番環境の準備が整っていることや、AKM-CLR が特定の外部インシデントを防止したことを主張するものではありません。ワークスペース、セッション、キャッシュ、コンテキスト漏洩を伴うマルチテナント AI 分離リスクに対する追加の制御ポイントを検証します。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このPAをリロードしてください

げ。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to irenee28/akm-clr-governance-layer development by creating an account on GitHub.

GitHub - irenee28/akm-clr-governance-layer · GitHub
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
irenee28
/
akm-clr-governance-layer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
irenee28/akm-clr-governance-layer
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits docs docs examples examples runnable_demo runnable_demo .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE.md LICENSE.md README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
AKM-CLR — Tenant/Task Governance Layer for Shared LLM Serving
AKM-CLR is a validated prototype governance/control layer for multi-tenant LLM infrastructure. It sits above OpenAI-compatible serving engines such as vLLM and helps decide whether a tenant request should be forwarded, fallbacked, or blocked before backend inference.
vLLM serves fast. AKM-CLR governs what is allowed to reach vLLM.
AKM-CLR is designed for AI platforms that run shared LLM backends, RAG systems, LoRA adapters, or tenant-specific workflows.
Domain and adapter routing policy
Pre-inference blocking for unsafe/cross-tenant requests
Safe fallback for wrong-domain or unauthorized requests
vLLM/OpenAI-compatible backend integration
Modern AI platforms increasingly serve many tenants, tasks, RAG sources, and adapters through shared infrastructure. Without a governance layer, unsafe or wrong-domain requests can still reach the backend even if the backend is technically working correctly.
AKM-CLR focuses on the governance gap between application logic and model serving:
Client / App
↓
AKM-CLR Governance Gateway
↓ only if authorized and safe
vLLM / OpenAI-compatible backend
↓
Model response
Validation summary
AKM-CLR was evaluated in controlled experiments:
29.5 GB max allocated VRAM on one H200
In a controlled comparison using Qwen2.5-7B-Instruct served by vLLM:
AKM-CLR is a validated prototype. It does not claim:
production readiness without hardening
superiority over vLLM as a serving engine
The claim is narrower and stronger:
AKM-CLR complements vLLM by adding tenant/task governance before inference.
We are looking for a small number of design partners operating:
vLLM or OpenAI-compatible inference services
fine-tuned or LoRA adapter infrastructure
regulated enterprise AI systems
AKM Gateway integration against your test backend
tenant/task governance test suite
Founder: Irénée Akilimali
Location: Kinshasa, DRC
Design partner inquiries: shukranimungu@gmail.com
Latest product-readiness validation
After the initial research validation, AKM-CLR was extended into a more product-oriented gateway prototype and tested across three additional experiments.
Experiment 1 — No-GPU product proof
The local no-GPU experiment validated AKM-CLR as a configurable governance layer using policy-as-code.
governance-safe decision rate: 100%
rate-limiting behavior: validated
Experiment 2 — Real vLLM integration proof
The RunPod real-vLLM experiment validated AKM-CLR above Qwen/Qwen2.5-1.5B-Instruct served through vLLM.
governance-safe decision rate: 100%
authorized backend-call rate: 100%
Authorized tenant requests reached vLLM. Unsafe or unauthorized requests, including cross-tenant canary probing, wrong-adapter access, wrong-RAG-source access, invalid API keys, and tenant spoofing, were blocked before backend inference.
Experiment 3 — Enterprise deployment preview
The enterprise deployment preview validated AKM-CLR packaging as an infrastructure component.
separated API-key secrets YAML
traffic-governance / rate-limiting behavior
This does not claim production readiness or guaranteed security. It validates that AKM-CLR can move from a research prototype toward a deployable governance/admission-control layer for OpenAI-compatible LLM serving stacks.
Experiment 5 — Workspace / Session / Cache / Context Boundary Guard
AKM-CLR was extended to test governance over workspace, session, cache, and retrieved context boundaries before inference.
This experiment validated whether a tenant is authorized to attach a specific workspace, session, cache entry, retrieved context object, adapter, and RAG source before the request reaches inference.
governance-safe decision rate: 100%
expected-decision match rate: 100%
workspace/session/cache/context boundary-safe rate: 100%
confidence fallback behavior: validated
output canary blocking: validated
This does not claim production readiness or that AKM-CLR would have prevented any specific external incident. It validates an additional control point for multi-tenant AI isolation risks involving workspace, session, cache, and context leakage.
Experiment 5 — Workspace / Session / Cache / Context Boundary Guard
AKM-CLR was extended to test governance over workspace, session, cache, and retrieved context boundaries before inference.
This experiment validated whether a tenant is authorized to attach a specific workspace, session, cache entry, retrieved context object, adapter, and RAG source before the request reaches inference.
governance-safe decision rate: 100%
expected-decision match rate: 100%
workspace/session/cache/context boundary-safe rate: 100%
confidence fallback behavior: validated
output canary blocking: validated
This does not claim production readiness or that AKM-CLR would have prevented any specific external incident. It validates an additional control point for multi-tenant AI isolation risks involving workspace, session, cache, and context leakage.
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
