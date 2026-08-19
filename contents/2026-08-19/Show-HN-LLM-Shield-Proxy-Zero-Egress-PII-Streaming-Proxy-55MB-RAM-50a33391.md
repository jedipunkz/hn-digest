---
source: "https://github.com/ninadphalak/LLM-Shield-Proxy"
hn_url: "https://news.ycombinator.com/item?id=49360413"
title: "Show HN: LLM-Shield-Proxy Zero-Egress PII Streaming Proxy (55MB RAM)"
article_title: "GitHub - ninadphalak/LLM-Shield-Proxy: A stateless, zero-latency reverse proxy for real-time PII redaction in LLM streams. · GitHub"
image: "https://repository-images.githubusercontent.com/1319884424/7cd0908f-5128-49c1-a64d-7a0a3fcbedb1"
author: "ninadphalak"
captured_at: "2026-08-19T12:24:42Z"
capture_tool: "hn-digest"
hn_id: 49360413
score: 2
comments: 0
posted_at: "2026-08-19T12:06:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: LLM-Shield-Proxy Zero-Egress PII Streaming Proxy (55MB RAM)

- HN: [49360413](https://news.ycombinator.com/item?id=49360413)
- Source: [github.com](https://github.com/ninadphalak/LLM-Shield-Proxy)
- Score: 2
- Comments: 0
- Posted: 2026-08-19T12:06:23Z

## Translation

タイトル: HN を表示: LLM-Shield-Proxy Zero-Egress PII ストリーミング プロキシ (55MB RAM)
記事のタイトル: GitHub - ninadpolak/LLM-Shield-Proxy: LLM ストリームでのリアルタイム PII 秘匿化のためのステートレスで遅延ゼロのリバース プロキシ。 · GitHub
説明: LLM ストリームでリアルタイム PII 編集を行うためのステートレスで遅延ゼロのリバース プロキシ。 - ninadphalak/LLM-Shield-Proxy

記事本文:
GitHub - ninadphalak/LLM-Shield-Proxy: LLM ストリームでのリアルタイム PII 秘匿化のためのステートレスで遅延ゼロのリバース プロキシ。 · GitHub
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
ニナドファラク
/
LLM シールド プロキシ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
216 コミット 216 コミット フォルダーとファイル
.github .github ベンチマーク ベンチマーク チャート/ llm-shield-proxy チャート/ llm-shield-proxy デプロイ/ helm/ llm-shield-proxy デプロイ/ helm/

llm-shield-proxy ドキュメント ドキュメントの例 例 llm_shield_proxy llm_shield_proxy スクリプト スクリプト テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml ARCHITECTURE.md ARCHITECTURE.md COMPLIANCE.md COMPLIANCE.md CTRIBUTING.md COTRIBUTING.md DEPLOYMENT.md DEPLOYMENT.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md アーキテクチャ.mmd アーキテクチャ.mmd cb_log.txt cb_log.txt docker-compose.yml docker-compose.yml e2e_test.py e2e_test.py 例外ログ.txt 例外ログ.txt full_test_log.txt full_test_log.txt full_test_log2.txt full_test_log2.txt locust_metrics_Exceptions.csv locust_metrics_Exceptions.csv locust_metrics_failures.csv locust_metrics_failures.csv locust_metrics_stats.csv locust_metrics_stats.csv locust_metrics_stats_history.csv locust_metrics_stats_history.csv pyproject.toml pyproject.toml pytest.ini pytest.ini reorg_readme.py reorg_readme.py要件.txt要件.txt test_buffer.py test_buffer.py test_out.txt test_out.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
大規模言語モデル向けの安全で高速なドロップイン PII 編集およびコンテキスト保存リバース プロキシ。
リアルタイム遅延を損なうことなく、LLM ストリームの SOC 2 Type II および HIPAA に準拠します。
LLM-Shield-Proxy は、企業 VPC 内に直接デプロイされたオープンソースのゼロ出力ミドルウェア リバース プロキシです。 OpenAI 互換の LLM API リクエストをインターセプトし、個人識別情報 (PII) と生の秘密をインフラストラクチャから流出する前に編集し、超低ストリーム遅延でリアルタイムの Server-Sent Events (SSE) チャット応答を決定的に再ハイドレートします。
リアルタイム ストリームを中断することなく、エンタープライズ プライバシー コンプライアンス (SOC 2、HIPAA、HITRUST) のブロックを解除するように設計されています。

ng レイテンシー)。
上流の統合とコンテキスト
このリポジトリは、以下で提案されているように、エンタープライズ サンドボックスでの SSE ストリームの断片化を解決するためのリファレンス プロキシ アーキテクチャとベンチマーク スイートを提供します。
上流の提案: NVIDIA/OpenShell #2763
プレプリント出版物: DOI: 10.5281/zenodo.21955770
⚡ 60 秒のクイックスタートと導入
🔄 ドロップインプルーフ (ゼロ SDK 統合)
LLM-Shield-Proxy は OpenAI 仕様をネイティブに模倣しているため、アプリケーション コードを書き直す必要はありません。
SDKのbase_urlまたはcurlコマンドのエンドポイントを変更するだけです。プロキシはペイロードをインターセプトして編集し、スキーマを正しい上流プロバイダーに自動的に変換します。
# ❌ 前: 生の PHI を OpenAI に直接送信
カール https://api.openai.com/v1/chat/completions \
-H " 権限: ベアラー sk-openai-key " \
-d ' {"messages": [{"role": "user", "content": "私の SSN は 000-00-0000"}]} '
# ✅ 後: LLM-Shield 経由でペイロードを送信 (ゼロ出力)
カール http://localhost:8000/v1/chat/completions \
-H " 認可: ベアラー シールド仮想キー " \
-d ' {"messages": [{"role": "user", "content": "私の SSN は 000-00-0000"}]} '
🚀 クイック スタート (Docker) — 3 行の Bash
ゼロ出力プロキシをスピンアップし、ライブ ストリーミング PII デモを 3 行で実行します。
#1. プロキシコンテナをバックグラウンドでスピンアップする
ドッカー構成 -d
#2. ヘルスプローブを確認する
カール http://localhost:8000/healthz
# 3. ライブデモスクリプトを実行する
Python の例/demo.py
📦 インストールオプションと構成戦略
LLM-Shield-Proxy は高度にモジュール化されています。特定のコンプライアンス ROI とメモリ制約に基づいてエンジンを構成できます。
Tier 3 ONNX NER の有効化: [ner] でインストールされている場合は、.env または環境変数で ENABLE_TIER3_ONNX_NER=true を設定して、ディープ ニューラル エンティティ抽出を有効にします。

riables (およびオプションでカスタム モデルの重みを指す ONNX_MODEL_PATH を指定します)。無効になっているかインストールされていない場合、エンジンは起動オーバーヘッドなしで自動的かつ正常に Tier 3 をバイパスします。
🧠 Bring Your Own Model (BYOM): カスタム ONNX トランスフォーマー
LLM-Shield-Proxy は、単一の NER モデルに固定されていません。エンタープライズ アーキテクチャでは、 ONNX_MODEL_PATH ( tokenizer.json とともに) を指定することで、ONNX にエクスポートされたドメイン固有の Hugging Face トランスフォーマーをプラグインできます。
ヘルスケアと HIPAA: 量子化された BioBERT 、 ClinicalBERT 、または Med-BERT モデルをロードして、臨床患者のメモと医療記録を編集します。
グローバル GDPR および多言語: フランス語、ドイツ語、スペイン語、および多言語のコンテキスト エンティティ抽出用に XLM-RoBERTa または mBERT をロードします。
リーガル テックと財務: 特殊な契約、NDA、財務監査証跡用に Legal-BERT または FinBERT をロードします。
無効時のオーバーヘッドゼロ: ENABLE_TIER3_ONNX_NER=false の場合、ONNX ランタイムは完全にバイパスされ、60MB 未満の超低 RAM と 6 µs 未満のフットプリントが維持されます。
🛡️ Bring Your Own Regex (BYOR): エンタープライズ ルール インジェクション
企業のコンプライアンスでは、多くの場合、独自の内部形式 (カスタム従業員 ID、内部プロジェクトのコード名、独自の請求トークンなど) をスキャンする必要があります。 LLM-Shield-Proxy を使用すると、壊滅的な ReDoS (正規表現サービス拒否) の危険を冒さずに、Tier 1 と一緒に評価されるカスタム正規表現ルールを挿入できます。
カスタム正規表現を挿入するには、custom_regex.yaml ファイルをプロキシにマウントし、CUSTOM_REGEX_PATH をそのファイルにポイントします。
セキュリティと ReDoS 耐性:
単純なリバース プロキシは、 (a+)+$ のような不適切に記述されたバックトラッキング正規表現に対して評価されるとクラッシュする可能性があります。これを防ぐために、LLM-Shield-Proxy はすべての BYOR カスタム パターンに対して google-re2 C++ エンジンを利用します。 FastAPI l 中に Pydantic 経由で YAML 設定を解析します。

ifespan 起動イベントを実行し、 re2 を使用してすべてのパターンをコンパイルし、正規表現がどれほど複雑か、ストリーミング ペイロードがどれほど敵対的であるかに関係なく、数学的に $O(N)$ の実行時間を保証します。これにより、マイクロ秒のレイテンシーのオーバーヘッドを犠牲にすることなく、ReDoS 攻撃に対する完全な耐性が確保されます。
# カスタム正規表現.yaml
カスタムパターン:
- 名前 : INTERNAL_EMPLOYEE_ID
パターン: " (?i)EMP-[A-Z]{3}- \\ d{5} "
説明 : " Acme Corp の内部従業員 ID と一致します "
# プロキシ サーバーをポート 8000 でローカルに起動します
llm-shield-proxy --host 0.0.0.0 --port 8000 --workers 1
🐳 Docker 経由で直接実行
docker run -d -p 8000:8000 \
-e OPENAI_API_KEY= " sk-your-openai-api-key " \
-e ホスト = " 0.0.0.0 " \
-e ポート=8000 \
--name llm-shield-proxy \
ghcr.io/ninadphalak/llm-shield-proxy:最新
1 行の SDK の変更
既存の OpenAI SDK Base_url がローカル LLM-Shield-Proxy インスタンスを指すようにします。
openaiインポートからOpenAI
クライアント = OpenAI (
api_key = "あなたのopenai-APIキー" ,
base_url = "http://localhost:8000/v1" , # LLM-Shield-Proxy を指す
)
応答 = クライアント 。チャット 。完成品。作成(
モデル = "gpt-4o-mini" 、
messages = [{ "role" : "user" , "content" : "サラ・コナー（sarah@example.com または 555-0199）までご連絡ください。" }]、
ストリーム = True 、
)
応答のチャンクの場合:
print (チャンク . 選択肢 [ 0 ]. デルタ . コンテンツまたは "" 、 end = "" 、フラッシュ = True )
💥 問題と LLM-Shield-Proxy ソリューションの比較
既存のレガシー プロキシ
LLM シールド プロキシ
リアルタイム SSE ストリーミングを破棄します。スキャン前に応答全体をバッファリングし、数秒の UI 遅延の停止を引き起こします。
超低遅延ストリーミング: SSE パケット ストリームとしてデルタごとに編集および再ハイドレートします。
大量のメモリ フットプリント: 大量の spaCy または PyTorch NLP ライブラリには 1GB ～ 2GB RAM が必要です。
超軽量 <60MB RAM: マイクロ秒のコンパイル済み正規表現 + Shannon エントロで実行

py + 合成発電エンジン。
データ責任: ユーザーの PII を長期データベースに保存します。
ゼロ長期ストレージ (ゼロデータ モード): データ責任をゼロにするために構築された自己破壊型 TTL セッション ボールト。厳密な「ゼロデータ モード」で動作します。プロンプト、PII、またはコンテキスト ウィンドウが永続ディスクや外部ストレージに書き込まれることはありません。
複雑なクラウド下り: データをサードパーティの SaaS 検査 API にルーティングします。
100% ゼロエグレス VPC: すべてのスキャンは安全な企業境界内でローカルに行われます。
🏛️ 信頼と透明性を目指して構築
高度に規制されたエンタープライズ環境、ゼロトラスト ネットワーク アーキテクチャ、セキュリティ最優先のエンジニアリング チーム向けに特別に設計されています。
データは建物内に保管されます。データをサードパーティのセキュリティ会社に送信することはありません。シールドは 100% 独自のサーバー内で実行されます。
ゼロデータ ストレージ: 機密性の高いプロンプトを保存したり記録したりしません。このシステムは、マッピングを自動的に消去する「自己破壊型」メモリ保管庫を使用します。
継続的な安定性: システムは、何時間も続く大量のシミュレートされた使用状況 (数千人の同時ユーザー) の下で積極的にテストされ、AI ツールがクラッシュしたり速度が低下したりしないことが確認されています。
透過的な設計: このシステムは、機密データを検出するために隠された「ブラック ボックス」AI に依存しません。数学的に証明された透明なルールを使用して、クレジット カード、SSN、医療記録番号などのパターンを検出します。
Microsoft Presidio を他のプロキシにしない理由は何ですか?
混雑した空間です。代替手段の代わりに LLM-Shield-Proxy を導入する必要がある理由は次のとおりです。
Microsoft Presidio / spaCy: 1 GB 以上の RAM を消費し、リクエストごとに 50 ～ 150 ミリ秒の遅延でイベント ループをブロックするレガシー ライブラリ。 (正規表現のためにユニバースを一時停止することほど「リアルタイム AI」を語るものはないからです)。 LLM-Shield-Proxy は、60 MB 未満のフラットなフットプリントを使用し、6 μs 未満の遅延を実現します。

えーっと。
クラウド AI 安全性 API (Azure/AWS): VPC から生データを送信して PII をチェックすることは、目的を達成できません。 LLM-Shield-Proxy を使用すると、データが編集されずにインフラストラクチャから離れることはありません。
標準正規表現ゲートウェイ: 非同期の Server-Sent Events (SSE) で中断されます。機密トークンが 2 つのストリーミング パケットに分割されている場合、標準のゲートウェイではトークンが漏洩します。 LLM-Shield-Proxy は、スライディング ウィンドウの先読みバッファを使用して、ストリームのフォーマットを壊すことなく分割トークンをシームレスに保持します。
LiteLLM / LangChain: LLM-Shield-Proxy はモデル ルーターまたはオーケストレーション フレームワークではありません。それはそれらと並行して機能します。 LLM-Shield-Proxy をオーケストレーターの前に配置して、ルーティング前のデータ マスキングを保証します。
🤝 オーケストレーター (私たちが補完するもの)
LLM-Shield-Proxy はモデルルーターではありません。業界標準のオーケストレーション ツールと並行して「リバース プロキシ サンドイッチ」で動作するように設計されています。既存の AI ルーティング インフラストラクチャと完全にスタックし、コードの変更は必要なく、すぐに使用できる以下のものと 100% 互換性があります。
オーケストレーション フレームワーク: LangChain、LlamaIndex、セマンティック カーネル、AutoGen、CrewAI。
AI ゲートウェイとルーター: LiteLLM、Cloudflare AI ゲートウェイ、Kong AI ゲートウェイ、Portkey。 (注: LLM-Shield-Proxy を LiteLLM の前にシームレスにスタックして、

[切り捨てられた]

## Original Extract

A stateless, zero-latency reverse proxy for real-time PII redaction in LLM streams. - ninadphalak/LLM-Shield-Proxy

GitHub - ninadphalak/LLM-Shield-Proxy: A stateless, zero-latency reverse proxy for real-time PII redaction in LLM streams. · GitHub
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
ninadphalak
/
LLM-Shield-Proxy
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
216 Commits 216 Commits Folders and files
.github .github benchmarks benchmarks charts/ llm-shield-proxy charts/ llm-shield-proxy deploy/ helm/ llm-shield-proxy deploy/ helm/ llm-shield-proxy docs docs examples examples llm_shield_proxy llm_shield_proxy scripts scripts tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml ARCHITECTURE.md ARCHITECTURE.md COMPLIANCE.md COMPLIANCE.md CONTRIBUTING.md CONTRIBUTING.md DEPLOYMENT.md DEPLOYMENT.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md architecture.mmd architecture.mmd cb_log.txt cb_log.txt docker-compose.yml docker-compose.yml e2e_test.py e2e_test.py exception_log.txt exception_log.txt full_test_log.txt full_test_log.txt full_test_log2.txt full_test_log2.txt locust_metrics_exceptions.csv locust_metrics_exceptions.csv locust_metrics_failures.csv locust_metrics_failures.csv locust_metrics_stats.csv locust_metrics_stats.csv locust_metrics_stats_history.csv locust_metrics_stats_history.csv pyproject.toml pyproject.toml pytest.ini pytest.ini reorg_readme.py reorg_readme.py requirements.txt requirements.txt test_buffer.py test_buffer.py test_out.txt test_out.txt View all files Repository files navigation
Secure, fast, and drop-in PII redaction and context preservation reverse proxy for Large Language Models.
SOC 2 Type II and HIPAA compliance for LLM streams without breaking real-time latency.
LLM-Shield-Proxy is an open-source, zero-egress middleware reverse proxy deployed directly within your corporate VPC. It intercepts OpenAI-compatible LLM API requests, redacts Personally Identifiable Information (PII) and raw secrets before they leave your infrastructure, and deterministically re-hydrates real-time Server-Sent Events (SSE) chat responses with ultra-low stream latency.
Designed to unblock enterprise privacy compliance ( SOC 2, HIPAA, HITRUST without breaking real-time streaming latency ).
Upstream Integration & Context
This repository provides the reference proxy architecture and benchmark suite for resolving SSE stream fragmentation in enterprise sandboxes, as proposed in:
Upstream Proposal: NVIDIA/OpenShell #2763
Preprint Publication: DOI: 10.5281/zenodo.21955770
⚡ 60-Second Quickstart & Deployment
🔄 The Drop-In Proof (Zero-SDK Integration)
Because LLM-Shield-Proxy natively mimics the OpenAI specification, you do not need to rewrite your application code .
You simply change the base_url in your SDK or the endpoint in your curl command. The proxy intercepts the payload, redacts it, and translates the schema to the correct upstream provider automatically.
# ❌ Before: Sending raw PHI directly to OpenAI
curl https://api.openai.com/v1/chat/completions \
-H " Authorization: Bearer sk-openai-key " \
-d ' {"messages": [{"role": "user", "content": "My SSN is 000-00-0000"}]} '
# ✅ After: Sending payload through LLM-Shield (Zero Egress)
curl http://localhost:8000/v1/chat/completions \
-H " Authorization: Bearer shield-virtual-key " \
-d ' {"messages": [{"role": "user", "content": "My SSN is 000-00-0000"}]} '
🚀 Quick Start (Docker) — 3 Lines of Bash
Spin up the zero-egress proxy and run the live streaming PII demo in 3 lines:
# 1. Spin up the proxy container in background
docker compose up -d
# 2. Verify health probe
curl http://localhost:8000/healthz
# 3. Run the live demo script
python examples/demo.py
📦 Installation Options & Configuration Strategy
LLM-Shield-Proxy is heavily modular. You can configure the engine based on your specific compliance ROI and memory constraints:
Enabling Tier 3 ONNX NER: When installed with [ner] , enable deep neural entity extraction by setting ENABLE_TIER3_ONNX_NER=true in your .env or environment variables (and optionally point ONNX_MODEL_PATH to custom model weights). If disabled or not installed, the engine automatically and gracefully bypasses Tier 3 with zero startup overhead.
🧠 Bring Your Own Model (BYOM): Custom ONNX Transformers
LLM-Shield-Proxy is not locked into a single NER model. Enterprise architectures can plug in any domain-specific Hugging Face transformer exported to ONNX by pointing ONNX_MODEL_PATH (along with its tokenizer.json ):
Healthcare & HIPAA: Load quantized BioBERT , ClinicalBERT , or Med-BERT models to redact clinical patient notes and medical records.
Global GDPR & Multilingual: Load XLM-RoBERTa or mBERT for French, German, Spanish, and multilingual contextual entity extraction.
Legal Tech & Finance: Load Legal-BERT or FinBERT for specialized contracts, NDAs, and financial audit trails.
Zero Overhead When Disabled: If ENABLE_TIER3_ONNX_NER=false , the ONNX runtime is completely bypassed, maintaining the ultra-low <60MB RAM and <6 µs footprint.
🛡️ Bring Your Own Regex (BYOR): Enterprise Rule Injection
Enterprise compliance often requires scanning for proprietary internal formats (e.g., custom employee IDs, internal project codenames, or proprietary billing tokens). LLM-Shield-Proxy allows you to inject custom regex rules that are evaluated alongside Tier 1 without risking catastrophic ReDoS (Regular Expression Denial of Service).
To inject custom regexes, mount a custom_regex.yaml file into the proxy and point CUSTOM_REGEX_PATH to it.
Security & ReDoS Immunity:
Naive reverse proxies can crash when evaluated against poorly written backtracking regexes like (a+)+$ . To prevent this, LLM-Shield-Proxy leverages the google-re2 C++ engine for all BYOR custom patterns. It parses your YAML configuration via Pydantic during the FastAPI lifespan startup event, and compiles all patterns using re2 , mathematically guaranteeing $O(N)$ execution time regardless of how complex your regex or how adversarial the streaming payload is. This ensures complete immunity against ReDoS attacks without sacrificing the microsecond latency overhead.
# custom_regex.yaml
custom_patterns :
- name : INTERNAL_EMPLOYEE_ID
pattern : " (?i)EMP-[A-Z]{3}- \\ d{5} "
description : " Matches internal Acme Corp employee IDs "
# Start the proxy server locally on port 8000
llm-shield-proxy --host 0.0.0.0 --port 8000 --workers 1
🐳 Run via Docker Directly
docker run -d -p 8000:8000 \
-e OPENAI_API_KEY= " sk-your-openai-api-key " \
-e HOST= " 0.0.0.0 " \
-e PORT=8000 \
--name llm-shield-proxy \
ghcr.io/ninadphalak/llm-shield-proxy:latest
1-Line SDK Change
Point your existing OpenAI SDK base_url to your local LLM-Shield-Proxy instance:
from openai import OpenAI
client = OpenAI (
api_key = "your-openai-api-key" ,
base_url = "http://localhost:8000/v1" , # Point to LLM-Shield-Proxy
)
response = client . chat . completions . create (
model = "gpt-4o-mini" ,
messages = [{ "role" : "user" , "content" : "Contact Sarah Connor at sarah@example.com or 555-0199." }],
stream = True ,
)
for chunk in response :
print ( chunk . choices [ 0 ]. delta . content or "" , end = "" , flush = True )
💥 The Problem vs. The LLM-Shield-Proxy Solution
Existing Legacy Proxies
LLM-Shield-Proxy
Destroys Real-Time SSE Streaming: Buffers entire responses before scanning, causing multi-second UI latency stalls.
Ultra-Low Latency Streaming: Redacts and re-hydrates delta-by-delta as SSE packets stream.
Heavy Memory Footprint: Requires 1GB–2GB RAM for heavy spaCy or PyTorch NLP libraries.
Ultra-Lightweight <60MB RAM: Runs on a microsecond compiled regex + Shannon entropy + synthetic generator engine.
Data Liability: Stores user PII in long-term databases.
Zero Long-Term Storage (Zero-Data Mode): Self-destructing TTL session vault built for zero data liability. Operates in strict "Zero-Data Mode"—no prompts, PII, or context windows are ever written to persistent disk or external storage.
Complex Cloud Egress: Routes data to 3rd-party SaaS inspection APIs.
100% Zero-Egress VPC: All scanning happens locally inside your secure corporate boundary.
🏛️ Built for Trust & Transparency
Designed specifically for highly regulated enterprise environments, Zero-Trust network architectures, and security-first engineering teams.
It keeps your data in your building: I do not send your data to a third-party security company. The shield runs 100% inside your own servers.
Zero-Data Storage: I do not save or log your sensitive prompts. The system uses a "self-destructing" memory vault that erases the mappings automatically.
Continuous Stability: The system has been aggressively tested under heavy, simulated usage (thousands of concurrent users) for hours on end to ensure it never crashes or slows down your AI tools.
Transparent Design: The system doesn't rely on hidden "black box" AI to detect sensitive data. It uses mathematically proven, transparent rules to detect patterns like Credit Cards, SSNs, and Medical Record Numbers.
Why Not Microsoft Presidio any other proxy?
It's a crowded space. Here is exactly why you should deploy LLM-Shield-Proxy instead of the alternatives:
Microsoft Presidio / spaCy: Legacy libraries that consume 1GB+ of RAM and block your event loop with 50-150ms of latency per request. (Because nothing says "real-time AI" like pausing the universe for regex). LLM-Shield-Proxy uses a flat <60 MB footprint with <6 µs latency overhead.
Cloud AI Safety APIs (Azure/AWS): Checking for PII by sending raw data out of your VPC defeats the purpose. With LLM-Shield-Proxy, the data never leaves your infrastructure unredacted.
Standard Regex Gateways: They break on asynchronous Server-Sent Events (SSE). If a sensitive token is split across two streaming packets, standard gateways let it leak. LLM-Shield-Proxy uses a sliding-window lookahead buffer to seamlessly hold split tokens without breaking stream formatting.
LiteLLM / LangChain: LLM-Shield-Proxy is not a model router or orchestration framework. It works alongside them. Put LLM-Shield-Proxy in front of your orchestrator to guarantee data masking before routing.
🤝 The Orchestrators (What we complement)
LLM-Shield-Proxy is not a model router. It is designed to work in a "Reverse Proxy Sandwich" alongside industry-standard orchestration tools. It stacks perfectly with your existing AI routing infrastructure, requires no code changes, and is 100% compatible out-of-the-box with:
Orchestration Frameworks: LangChain, LlamaIndex, Semantic Kernel, AutoGen, CrewAI.
AI Gateways & Routers: LiteLLM, Cloudflare AI Gateway, Kong AI Gateway, Portkey. (Note: You can seamlessly stack LLM-Shield-Proxy in front of LiteLLM to achieve bo

[truncated]
