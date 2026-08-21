---
source: "https://github.com/red-orbita/bulwark-gateway"
hn_url: "https://news.ycombinator.com/item?id=49385285"
title: "Bulwark Gateway – fail-closed security proxy for LLM agents (self-hosted)"
article_title: "GitHub - red-orbita/bulwark-gateway: Security guardrail proxy for AI agents. Intercepts and enforces policies on tool calls between users and LLM backends. Multi-tenant, fail-closed, 78+ prompt injection patterns, threat intel feeds, RBAC tool policies, output filtering, SIEM integration. Zero-LLM h\n[truncated]"
image: "https://opengraph.githubassets.com/d54c2615580826a99db9ad9220b2b990030cba1309bedf153407133fe5e107a5/red-orbita/bulwark-gateway"
author: "rokitoh"
captured_at: "2026-08-21T08:27:23Z"
capture_tool: "hn-digest"
hn_id: 49385285
score: 1
comments: 0
posted_at: "2026-08-21T08:18:00Z"
tags:
  - hacker-news
  - translated
---

# Bulwark Gateway – fail-closed security proxy for LLM agents (self-hosted)

- HN: [49385285](https://news.ycombinator.com/item?id=49385285)
- Source: [github.com](https://github.com/red-orbita/bulwark-gateway)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T08:18:00Z

## Translation

タイトル: Bulwark Gateway – LLM エージェント用のフェールクローズ型セキュリティ プロキシ (自己ホスト型)
記事のタイトル: GitHub - red-orbita/bulwark-gateway: AI エージェント用のセキュリティ ガードレール プロキシ。ユーザーと LLM バックエンド間のツール呼び出しのポリシーを傍受し、適用します。マルチテナント、フェールクローズ、78 以上のプロンプト インジェクション パターン、脅威インテリジェンス フィード、RBAC ツール ポリシー、出力フィルタリング、SIEM 統合。ゼロLLM h
[切り捨てられた]
説明: AI エージェント用のセキュリティ ガードレール プロキシ。ユーザーと LLM バックエンド間のツール呼び出しのポリシーを傍受し、適用します。マルチテナント、フェールクローズ、78 以上のプロンプト インジェクション パターン、脅威インテリジェンス フィード、RBAC ツール ポリシー、出力フィルタリング、SIEM 統合。ゼロ LLM ホット パス、p95 < 40ms。株式会社アドミンポータル
[切り捨てられた]

記事本文:
GitHub - red-orbita/bulwark-gateway: AI エージェント用のセキュリティ ガードレール プロキシ。ユーザーと LLM バックエンド間のツール呼び出しのポリシーを傍受し、適用します。マルチテナント、フェールクローズ、78 以上のプロンプト インジェクション パターン、脅威インテリジェンス フィード、RBAC ツール ポリシー、出力フィルタリング、SIEM 統合。ゼロ LLM ホット パス、p95 < 40ms。管理者ポータルが含まれています。 Kubernetes ネイティブ。 · GitHub
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
レッドオービタ
/
防波堤ゲートウェイ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスターB

牧場 タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
156 コミット 156 コミット フォルダーとファイル
.github/ workflows .github/ workflows admin admin ci ci config configdeploy/ argocddeploy/ argocd docker docker docs docs helm/ bulwark-gateway helm/ bulwark-gateway k8s k8s モデル モデル監視 監視プラグイン/ サンプル プラグイン/ サンプル prometheus プロメテウス レポート/ ブログ証拠レポート/ ブログ証拠スクリプト スクリプトsdk-go sdk-go sdk sdk Secrets Secrets src src テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .semgrep.yml .semgrep.yml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLA.md CLA.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE LICENSING.md LICENSING.md Makefile Makefile README.md README.md docker-compose.yml docker-compose.yml package.json package.json pyproject.toml pyproject.toml 要件-admin.lock要件-admin.lock要件.lock要件.lock tailwind.config.js tailwind.config.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クラウド環境の AI エージェント用のセキュリティ ガードレール プロキシ。
ユーザーと LLM エージェント間のツール呼び出しのポリシーを傍受、検証、および適用します。ユーザーが敵対的になる可能性がある環境向けに設計されています (デフォルトではフェールクローズされます)。
Bulwark Gateway は、ユーザー/アプリケーションと LLM バックエンド (OpenAI、Ollama、vLLM、Azure など) の間に位置します。すべてのリクエストは、バックエンドに到達する前に複数のセキュリティ層を通過します。
認証 - JWT/API キー検証 (フェイルクローズ)
入力ガードレール — プロンプトインジェクション、ジェイルブレイク、エンコード回避を検出します。
IOC チェック — 脅威インテリジェンス フィードから悪意のある URL/IP/ドメインをスキャンします。
ツール ポリシー — テナント/エージェントごとの RBAC の適用
奥羽

tput フィルター — シークレット/PII を秘匿化し、応答内の間接的な挿入を検出します
レート リミッター — Redis を介したテナントごとのリクエスト スロットル
いずれかのレイヤーが脅威を検出した場合、リクエストは直ちにブロックされます (フェールクローズ)。
┌─────────────────────┐
│ ブルワークゲートウェイ │
│ │
ユーザーリクエスト ───► 認証 ► 入力ガードレール ► IOC チェック │
X-テナントID │ │ │
X-Agent-ID │ エージェント登録 │
│ (マルチバックエンド) │
│ │ │
│ バックエンドに転送 │
│ │ │
│ ツールポリシー ◄── 応答 ──► 出力フィルタ │
━━━━━━━━━━━━━━━━━┘
│
┌───────────────┼─────────────┐
▼ ▼ ▼
バックエンド A (RAG) バックエンド B (LLM) バックエンド C (エージェント)
コンポーネント
港
説明
プロキシ
8080
セキュリティホットパス - すべてのLLMリクエストをインターセプトします
管理者ポータル
8090
構成、監視、監査ログ用の Web UI
レディス
6379
レート制限、状態、セッション管理
プロメテウス
9090
メトリクスの収集
グラファナ
3000
ダッシュボードと視覚化
完全なアーキテクチャの詳細: docs/ARCHITECTURE.md
マルチテナント、マルチエージェント — テナント/エージェントごとに異なるバックエンドにリクエストをルーティングします。
ゼロ LLM ホット パス — 正規表現 + Pydantic + キャッシュのみ。 p95 < 40ms オーバーヘッド
400 を超える検出パターン — プロンプト インジェクション、ジェイルブレイク、エンコード回避、多言語 (ES/ZH/AR)
4 つの脅威インテリジェンス フィード — URLhaus、ThreatFox、AlienVault OTX、AbuseIPDB (+ MISP、OpenCTI、VirusTotal、Shodan)
ストリーミング ツール呼び出しバッファリング — ツール呼び出しは t を生成する前に検証されます。

クライアント
自己保護 - エージェントによるゲートウェイ構成の変更をブロックします。
ホットリロード可能 — 再起動せずにポリシー、IOC、およびエージェント レジストリをリロードします
管理者ポータル — ゲートウェイのあらゆる側面を管理するための完全な Web UI
SIEM 統合 — 13 のプラットフォーム (Wazuh、Splunk、Elastic、QRadar、Datadog など) にエクスポート
通知チャネル - Slack、Teams、Discord、PagerDuty、Opsgenie、Telegram、電子メール、Google Chat
Kubernetes ネイティブ — NetworkPolicies、HPA、PDB、Pod Security を備えた完全な K8s マニフェスト
監査証跡 — すべての管理上の変更の不変ログ
エンタープライズ シークレット — Vault、AWS SM、Azure KV、GCP SM、CyberArk、SealedSecrets
ほとんどの LLM セキュリティ ツールは、アプリ コードに埋め込むライブラリ/SDK として出荷されます。
プロンプトを送信するホスト型 SaaS として。ブルワーク・ゲートウェイは、
OpenAI 互換の前面に配置される自己ホスト型のフェイルクローズ プロキシ
バックエンド — アプリのコードは変更されず、ネットワークからのプロンプトも表示されません。
正直な範囲。 Bulwark はガードレール プロキシであり、WAF や
モデルホスティングプラットフォーム。自由形式のチャット入力における従来の SQLi/XSS は機能しません。
設計により入力層と確実に一致します。これらは、
ペイロードが実際に DB/ファイルシステムに到達するツール引数層。を参照してください。
まさにその内容についてのギャップレポートを公開しました
釣れたり釣れなかったり。ホット パスは純粋な正規表現 (~446 入力 + ~150 出力)
パターン）のため、検出は高速で監査可能ですが、セマンティクスの代替にはなりません。
あらゆるエッジケースに分類子 — ML スキャナーはオプションのレイヤーとして利用できます。
比較には、各カテゴリの共通の導入モデルが反映されています。個人的な
ツールは異なります。各ベンダーの現在の機能を確認します。
Kubernetes 1.28+ (本番環境) または Docker Compose (開発環境)
Kubernetes (本番環境 — Helm)
Helm はマネージド クラスター (AKS/EKS/GKE) の推奨パスです

。ビルドしてプッシュする
イメージをレジストリに保存し、チャートでそれらのイメージをポイントします。
<REGISTRY> はコンテナー レジストリ パスです。 myacr.azurecr.io 、
123456789012.dkr.ecr.eu-west-1.amazonaws.com 、または ghcr.io/my-org 。
# 1. イメージをビルドしてレジストリにプッシュします (最初に `docker login <REGISTRY>` を実行します)
docker build -t < REGISTRY > /bulwark-gateway-proxy:1.0.0 -f Dockerfile 。
docker build -t < REGISTRY > /bulwark-gateway-admin:1.0.0 -f docker/Dockerfile.admin 。
docker Push < REGISTRY > /bulwark-gateway-proxy:1.0.0
docker Push < レジストリ > /bulwark-gateway-admin:1.0.0
# 2. (プライベート レジストリのみ) ポッドが使用するプル シークレットを作成します。
kubectl 名前空間を作成します bulwark-gateway
kubectl シークレット docker-registry bulwark-registry を作成 \
--docker-server= < レジストリ > \
--docker-username= < ユーザー名 > \
--docker-password= < パスワード > \
-n ブルワークゲートウェイ
# 3. インストール (アプリのシークレットはチャートによって自動生成されます)
helm インストール bulwark ./helm/bulwark-gateway \
--namespace bulwark-gateway --create-namespace \
--set backend.ip= < YOUR_LLM_BACKEND_IP > \
--set proxy.image.repository= < レジストリ > /bulwark-gateway-proxy \
--set admin.image.repository= < レジストリ > /bulwark-gateway-admin \
--set proxy.image.tag=1.0.0 \
--set admin.image.tag=1.0.0 \
--set ' imagePullSecrets[0].name=bulwark-registry ' # パブリック レジストリの場合は省略します
#4. 検証する
kubectl ポッドを取得 -n bulwark-gateway
Helm テスト bulwark -n bulwark-gateway
外部 Redis、TLS/ingress については docs/DEPLOYMENT.md を参照してください。
完全なvalues.yamlリファレンス。
Kubernetes (ローカル — minikube/kind)
ローカル クラスターの場合、k8s/deploy.sh はイメージを構築し、クラスターにロードします。
そして、KusTOMize マニフェストを 1 つのステップで適用します。
#1. シークレットを生成する
./secrets/init.sh
# 2. ビルド + ロード + デプロイ (minikube/kind を自動検出し、シークレットを生成します)
./k8s/deploy.sh
# リモコンの場合

ローカルロードの代わりにレジストリを使用します。
# IMAGE_REGISTRY=<REGISTRY>/ ./k8s/deploy.sh --backend-ip <IP>
#3. 検証する
kubectl ポッドを取得 -n bulwark-gateway
Docker Compose (開発)
#1. シークレットを生成する
./secrets/init.sh
# 2. すべてのサービスを開始する
ドッカー構成 -d
#3. アクセス
# プロキシ: http://localhost:8080
# 管理者: http://localhost:8090
# グラファナ: http://localhost:3000
サービスにアクセスする
# ポートフォワード (K8s)
kubectl ポートフォワード svc/proxy 8080:8080 -n bulwark-gateway
kubectl ポートフォワード svc/admin 8090:8090 -n bulwark-gateway
# または Ingress 経由:
# プロキシ: https://bulwark-gateway.local
# 管理者: https://admin.bulwark-gateway.local
プロキシをテストする
# 健康診断
カール http://localhost:8080/health
# リクエストを送信します (API キーはベアラー トークンとして渡されます)
curl -X POST http://localhost:8080/v1/chat/completions \
-H " Content-Type: application/json " \
-H " 認証: ベアラー your-API キー " \
-H " X-テナント ID: デフォルト " \
-H " X エージェント ID: サポートボット " \
-d ' {"モデル": "gpt-4", "メッセージ": [{"役割": "ユーザー", "コンテンツ": "Hello"}]} '
完全な展開ガイド: docs/DEPLOYMENT.md
ファイル
目的
config/agents.yaml
テナント → バックエンド マッピング、認証設定
config/policies/*.yaml
テナントごとのセキュリティ ポリシー (RBAC)
config/notifications.yaml
通知チャネルの定義
config/siem/*.yaml
SIEMプラットフォームテンプレート
config/iocs.json
IOC データベース (フィードによって自動更新)
環境変数
変数
説明
BULWARK_JWT_SECRET
JWT 署名キー (または *_FILE バリアント)
BULWARK_REDIS_URL
Redis接続URL
BULWARK_REDIS_PASSWORD
Redis 認証 (または *_FILE バリアント)
BULWARK_API_KEYS
カンマ区切りの API キー (または *_FILE )
BULWARK_WEBHOOK_ALERT_URLS
従来の通知 Webhook
BULWARK_LOG_LEVEL
ログレベル (INFO、DEBUG など)
すべてのシークレットは *_FILE パターンをサポートしています。環境変数がマウントされたファイルを指します。
環境:
- な

私：BULWARK_JWT_SECRET_FILE
値: /run/secrets/jwt-secret
エージェントレジストリの例
# config/agents.yaml
テナント：
例-corp :
backend_url : " ${BULWARK_BACKEND_URL:-http://ollama:11434} "
auth_token : " ${BACKEND_AUTH_TOKEN} "
allowed_models : ["gpt-4"、"gpt-3.5-turbo"]
レート制限rpm : 60
ポリシーの例
# config/policies/example-corp.yaml
tenant_id : example-corp
ツール:
許可される:
- ウェブ検索
- コードインタープリター
ブロックされました :
- ファイルシステム
- シェル_exec
ガードレール :
最大トークン数 : 4096
block_on_injection : true
管理者ポータル
/ (ポート 8090) の Web ベースの管理インターフェイス。
ページ
機能
ダッシュボード
リアルタイムのメトリクス、最近のブロック、スパークライン
ポリシー
CRUD、検証、ホットリロード
ガードレール
パターン管理、サンドボックステスト
SIEM輸出
トランスポート構成、接続テスト
通知
アラート チャネル管理 (Slack、Teams、電子メールなど)
監査ログ
不変のアクション履歴、エクスポート
オーケストレーター
自動化されたセキュリティテスト
カバレッジマトリックス
OWASP LLM トップ 10 検出マップ
IOC
脅威インテリジェンス フィード管理
テナント
テナントの登録と構成
エージェント
バックエンドのヘルスモニタリング
アクセス制御
RBAC の役割と権限
エンリッチメント
攻撃リプレイブラウザ、回避テレメトリ、正規表現候補レビュー
スキル
導入前スキル/MCP セキュリティ スキャナー (SkillSpector)
プラグイン
プラグイン ハブ — インストール、有効化、セキュリティ監査
評価
赤チーム敵対的評価ランナー
発見
エージェント / シャドウ AI / MCP の発見とリスク評価
ステータス
システムの健全性 (Redis、プロキシ、スキャナ、テレメトリ)
デフォルトの認証情報
ユーザー
役割
デフォルトのパスワード
秘密鍵
管理者
管理者
ブルワーク管理者
ADMIN_PASSWORD
セキュリティ
セキュリティ
防波堤セキュリティ
セキュリティ_パスワード

[切り捨てられた]

## Original Extract

Security guardrail proxy for AI agents. Intercepts and enforces policies on tool calls between users and LLM backends. Multi-tenant, fail-closed, 78+ prompt injection patterns, threat intel feeds, RBAC tool policies, output filtering, SIEM integration. Zero-LLM hot path, p95 < 40ms. Admin Portal inc
[truncated]

GitHub - red-orbita/bulwark-gateway: Security guardrail proxy for AI agents. Intercepts and enforces policies on tool calls between users and LLM backends. Multi-tenant, fail-closed, 78+ prompt injection patterns, threat intel feeds, RBAC tool policies, output filtering, SIEM integration. Zero-LLM hot path, p95 < 40ms. Admin Portal included. Kubernetes-native. · GitHub
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
red-orbita
/
bulwark-gateway
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
156 Commits 156 Commits Folders and files
.github/ workflows .github/ workflows admin admin ci ci config config deploy/ argocd deploy/ argocd docker docker docs docs helm/ bulwark-gateway helm/ bulwark-gateway k8s k8s models models monitoring monitoring plugins/ examples plugins/ examples prometheus prometheus reports/ blog-evidence reports/ blog-evidence scripts scripts sdk-go sdk-go sdk sdk secrets secrets src src tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .semgrep.yml .semgrep.yml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLA.md CLA.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE LICENSING.md LICENSING.md Makefile Makefile README.md README.md docker-compose.yml docker-compose.yml package.json package.json pyproject.toml pyproject.toml requirements-admin.lock requirements-admin.lock requirements.lock requirements.lock tailwind.config.js tailwind.config.js View all files Repository files navigation
Security guardrail proxy for AI agents in cloud environments.
Intercepts, validates, and enforces policies on tool calls between users and LLM agents. Designed for environments where the user is potentially adversarial (fail-closed by default).
Bulwark Gateway sits between your users/applications and your LLM backends (OpenAI, Ollama, vLLM, Azure, etc.). Every request passes through multiple security layers before reaching the backend:
Authentication — JWT/API key validation (fail-closed)
Input Guardrail — Detects prompt injections, jailbreaks, encoding evasion
IOC Check — Scans for malicious URLs/IPs/domains from threat intel feeds
Tool Policy — RBAC enforcement per tenant/agent
Output Filter — Redacts secrets/PII, detects indirect injection in responses
Rate Limiter — Per-tenant request throttling via Redis
If any layer detects a threat, the request is blocked immediately (fail-closed).
┌──────────────────────────────────────────────┐
│ Bulwark Gateway │
│ │
User Request ─────► Auth ► Input Guardrail ► IOC Check │
X-Tenant-ID │ │ │
X-Agent-ID │ Agent Registry │
│ (multi-backend) │
│ │ │
│ Forward to backend │
│ │ │
│ Tool Policy ◄── Response ──► Output Filter │
└──────────────┼──────────────────────────────┘
│
┌─────────────────────────┼─────────────────────────┐
▼ ▼ ▼
Backend A (RAG) Backend B (LLM) Backend C (Agent)
Component
Port
Description
Proxy
8080
Security hot path — intercepts all LLM requests
Admin Portal
8090
Web UI for configuration, monitoring, audit logs
Redis
6379
Rate limiting, state, session management
Prometheus
9090
Metrics collection
Grafana
3000
Dashboards and visualization
Full architecture details: docs/ARCHITECTURE.md
Multi-tenant, multi-agent — Route requests to different backends per tenant/agent
Zero-LLM hot path — Only regex + Pydantic + cache; p95 < 40ms overhead
400+ detection patterns — Prompt injection, jailbreak, encoding evasion, multilingual (ES/ZH/AR)
4 threat intel feeds — URLhaus, ThreatFox, AlienVault OTX, AbuseIPDB (+ MISP, OpenCTI, VirusTotal, Shodan)
Streaming tool call buffering — Tool calls validated BEFORE yielding to client
Self-protection — Blocks agents from modifying gateway config
Hot-reloadable — Policies, IOCs, and agent registry reload without restart
Admin Portal — Full web UI for managing all aspects of the gateway
SIEM integration — Export to 13 platforms (Wazuh, Splunk, Elastic, QRadar, Datadog, etc.)
Notification channels — Slack, Teams, Discord, PagerDuty, Opsgenie, Telegram, Email, Google Chat
Kubernetes-native — Full K8s manifests with NetworkPolicies, HPA, PDB, Pod Security
Audit trail — Immutable log of all administrative changes
Enterprise secrets — Vault, AWS SM, Azure KV, GCP SM, CyberArk, SealedSecrets
Most LLM-security tools ship as a library/SDK you embed in your app code, or
as a hosted SaaS you send your prompts to. Bulwark Gateway is a
self-hosted, fail-closed proxy that sits in front of any OpenAI-compatible
backend — no code changes in your app, no prompts leaving your network.
Honest scope. Bulwark is a guardrail proxy , not a WAF and not a
model-hosting platform. Classic SQLi/XSS on free-form chat input is not
reliably matched by the input layer by design — those are enforced at the
tool-argument layer where the payload actually reaches a DB/filesystem. See the
published gap report for exactly what it
does and does not catch. The hot path is pure regex (~446 input + ~150 output
patterns), so detection is fast and auditable but not a substitute for a semantic
classifier on every edge case — ML scanners are available as an optional layer.
Comparison reflects the common deployment model of each category; individual
tools vary. Verify against each vendor's current capabilities.
Kubernetes 1.28+ (production) or Docker Compose (development)
Kubernetes (Production — Helm)
Helm is the recommended path for managed clusters (AKS/EKS/GKE). Build and push
the images to your registry, then point the chart at them.
<REGISTRY> is your container registry path, e.g. myacr.azurecr.io ,
123456789012.dkr.ecr.eu-west-1.amazonaws.com , or ghcr.io/my-org .
# 1. Build and push images to your registry (run `docker login <REGISTRY>` first)
docker build -t < REGISTRY > /bulwark-gateway-proxy:1.0.0 -f Dockerfile .
docker build -t < REGISTRY > /bulwark-gateway-admin:1.0.0 -f docker/Dockerfile.admin .
docker push < REGISTRY > /bulwark-gateway-proxy:1.0.0
docker push < REGISTRY > /bulwark-gateway-admin:1.0.0
# 2. (Private registry only) create the pull secret the pods use
kubectl create namespace bulwark-gateway
kubectl create secret docker-registry bulwark-registry \
--docker-server= < REGISTRY > \
--docker-username= < USERNAME > \
--docker-password= < PASSWORD > \
-n bulwark-gateway
# 3. Install (app secrets are auto-generated by the chart)
helm install bulwark ./helm/bulwark-gateway \
--namespace bulwark-gateway --create-namespace \
--set backend.ip= < YOUR_LLM_BACKEND_IP > \
--set proxy.image.repository= < REGISTRY > /bulwark-gateway-proxy \
--set admin.image.repository= < REGISTRY > /bulwark-gateway-admin \
--set proxy.image.tag=1.0.0 \
--set admin.image.tag=1.0.0 \
--set ' imagePullSecrets[0].name=bulwark-registry ' # omit for public registries
# 4. Verify
kubectl get pods -n bulwark-gateway
helm test bulwark -n bulwark-gateway
See docs/DEPLOYMENT.md for external Redis, TLS/ingress, and
full values.yaml reference.
Kubernetes (Local — minikube/kind)
For local clusters, k8s/deploy.sh builds images, loads them into the cluster,
and applies the Kustomize manifests in one step:
# 1. Generate secrets
./secrets/init.sh
# 2. Build + load + deploy (auto-detects minikube/kind, generates secrets)
./k8s/deploy.sh
# For a remote registry instead of local load:
# IMAGE_REGISTRY=<REGISTRY>/ ./k8s/deploy.sh --backend-ip <IP>
# 3. Verify
kubectl get pods -n bulwark-gateway
Docker Compose (Development)
# 1. Generate secrets
./secrets/init.sh
# 2. Start all services
docker compose up -d
# 3. Access
# Proxy: http://localhost:8080
# Admin: http://localhost:8090
# Grafana: http://localhost:3000
Access Services
# Port-forward (K8s)
kubectl port-forward svc/proxy 8080:8080 -n bulwark-gateway
kubectl port-forward svc/admin 8090:8090 -n bulwark-gateway
# Or via Ingress:
# Proxy: https://bulwark-gateway.local
# Admin: https://admin.bulwark-gateway.local
Test the Proxy
# Health check
curl http://localhost:8080/health
# Send a request (API keys are passed as a Bearer token)
curl -X POST http://localhost:8080/v1/chat/completions \
-H " Content-Type: application/json " \
-H " Authorization: Bearer your-api-key " \
-H " X-Tenant-ID: default " \
-H " X-Agent-ID: support-bot " \
-d ' {"model": "gpt-4", "messages": [{"role": "user", "content": "Hello"}]} '
Full deployment guide: docs/DEPLOYMENT.md
File
Purpose
config/agents.yaml
Tenant → backend mapping, auth config
config/policies/*.yaml
Per-tenant security policies (RBAC)
config/notifications.yaml
Notification channel definitions
config/siem/*.yaml
SIEM platform templates
config/iocs.json
IOC database (auto-updated by feeds)
Environment Variables
Variable
Description
BULWARK_JWT_SECRET
JWT signing key (or *_FILE variant)
BULWARK_REDIS_URL
Redis connection URL
BULWARK_REDIS_PASSWORD
Redis auth (or *_FILE variant)
BULWARK_API_KEYS
Comma-separated API keys (or *_FILE )
BULWARK_WEBHOOK_ALERT_URLS
Legacy notification webhooks
BULWARK_LOG_LEVEL
Logging level (INFO, DEBUG, etc.)
All secrets support the *_FILE pattern — point an env var to a mounted file:
env :
- name : BULWARK_JWT_SECRET_FILE
value : /run/secrets/jwt-secret
Agent Registry Example
# config/agents.yaml
tenants :
example-corp :
backend_url : " ${BULWARK_BACKEND_URL:-http://ollama:11434} "
auth_token : " ${BACKEND_AUTH_TOKEN} "
allowed_models : ["gpt-4", "gpt-3.5-turbo"]
rate_limit_rpm : 60
Policy Example
# config/policies/example-corp.yaml
tenant_id : example-corp
tools :
allowed :
- web_search
- code_interpreter
blocked :
- file_system
- shell_exec
guardrails :
max_tokens : 4096
block_on_injection : true
Admin Portal
Web-based management interface at / (port 8090).
Page
Function
Dashboard
Real-time metrics, recent blocks, sparklines
Policies
CRUD, validation, hot-reload
Guardrails
Pattern management, sandbox testing
SIEM Export
Transport configuration, connectivity testing
Notifications
Alert channel management (Slack, Teams, Email, etc.)
Audit Log
Immutable action history, export
Orchestrator
Automated security testing
Coverage Matrix
OWASP LLM Top 10 detection map
IOCs
Threat intel feed management
Tenants
Tenant registration and config
Agents
Backend health monitoring
Access Control
RBAC roles and permissions
Enrichment
Attack replay browser, evasion telemetry, regex-candidate review
Skills
Pre-deployment skill/MCP security scanner (SkillSpector)
Plugins
Plugin hub — install, enable, security audit
Evaluation
Red-team adversarial evaluation runner
Discovery
Agent / shadow-AI / MCP discovery and risk assessment
Status
System health (Redis, proxy, scanner, telemetry)
Default Credentials
User
Role
Default Password
Secret Key
admin
Admin
bulwark-admin
ADMIN_PASSWORD
security
Security
bulwark-security
SECURITY_PASSW

[truncated]
