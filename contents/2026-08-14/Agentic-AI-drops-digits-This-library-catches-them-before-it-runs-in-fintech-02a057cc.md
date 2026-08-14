---
source: "https://github.com/insightitsGit/PrismManifest"
hn_url: "https://news.ycombinator.com/item?id=49301925"
title: "Agentic AI drops digits, This library catches them before it runs in fintech"
article_title: "GitHub - insightitsGit/PrismManifest: Zero-trust tool-argument gate: signed ParameterManifest before deterministic Group 3 DAGs (PyPI: prismmanifest) · GitHub"
author: "parvaamin"
captured_at: "2026-08-14T17:45:01Z"
capture_tool: "hn-digest"
hn_id: 49301925
score: 1
comments: 0
posted_at: "2026-08-14T17:32:07Z"
tags:
  - hacker-news
  - translated
---

# Agentic AI drops digits, This library catches them before it runs in fintech

- HN: [49301925](https://news.ycombinator.com/item?id=49301925)
- Source: [github.com](https://github.com/insightitsGit/PrismManifest)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T17:32:07Z

## Translation

タイトル: Agentic AI は数字をドロップします。このライブラリはフィンテックで実行される前に数字をキャッチします。
記事タイトル: GitHub - InsightitsGit/PrismManifest: ゼロトラスト ツール引数ゲート: 決定論的なグループ 3 DAG の前に署名された ParameterManifest (PyPI: prismmanifest) · GitHub
説明: ゼロトラスト ツール引数ゲート: 決定論的なグループ 3 DAG の前に署名された ParameterManifest (PyPI: prismmanifest) - InsightitsGit/PrismManifest

記事本文:
GitHub - InsightitsGit/PrismManifest: ゼロトラスト ツール引数ゲート: 決定論的なグループ 3 DAG の前の署名された ParameterManifest (PyPI: prismmanifest) · GitHub
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
洞察力Git
/
プリズムマニフェスト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
20 コミット 20 コミット .github .github cpp/gate cpp/gate cuda/ kernels cuda/ kernels docs docs prismmanifest prismmanifest proto proto reports

レポート スキーマ スキーマ スクリプト スクリプト テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md SECURITY.md SECURITY.md azure-pipelines.yml azure-pipelines.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
決定論的な AI ツール実行のためのゼロトラスト ツール引数ゲート。
(以前の ParamGate — 同じ設計; パッケージ prismmanifest 、 CLI prismmanifest-gate 。)
PrismManifest は、確率抽出器 (LLM、OCR、テーブル パーサー) と
決定論的なグループ 3 は DAG を計算します。未検証の金額が DAG に入力されることはありません。
グループ 3 の境界をクリアする Ed25519 署名付き ParameterManifest のみ
通過が許可されます。
決定論的エンジン + 未検証の確率的入力 = 決定論的誤答
キーワード: ツール引数ゲート、ParameterManifest、桁落ち防止、ゼロトラスト LLM ツール引数、決定論的 DAG 境界、OCR マネー フィールド認証、PrismManifest
確率システム (LLM、OCR) は、数値がどこにあるかを提案することに優れています。
彼らは税務機関に入力される金額を承認することを信頼してはなりません。
台帳または引受モデル。 PrismManifest はその認可境界です。
pip インストール プリズムマニフェスト
オプションの追加機能:
pip インストール " prismmanifest[dev] "
pip install " prismmanifest[cuda] " # NVIDIA GPU + Numba
pip install " prismmanifest[kms-azure] " # Azure Key Vault エンベロープ キー
pip install " prismmanifest[kms-aws] " # AWS KMS エンベロープ (オプション)
ソースから (編集可能):
git clone https://github.com/insightitsGit/PrismManifest.git
cd プリズムマニフェスト
python -m pip install -e " .[dev] "
python -m pytest -q
インストール後の CLI: prismmanifest-gate 。
AI アーキテクトはこれをどのように使用すべきか
LLM/OCR/パーサーを信頼できないものとして扱います。
プリスを置く

mManifest は抽出後、金額を消費する計算機、引受業務、または元帳ツールの前に実行します。
ルートの結果: ACCEPT → DAG を実行 · ACCEPT_PENDING_HUMAN → レビュー · REJECT → 停止。
モデルが生成したお金のテキストをツールの議論にしないでください。証拠に拘束された、署名されたマニフェストのみを使用してください。
LLM / OCR → PrismManifest → 署名付き ParameterManifest → グループ 3 DAG
追加の LLM ベンチはオプションの証拠です。コアのゲート設計は変更されません。
実際の顧客文書を見つけることは、複数モデルの FA 調査よりも製造クレームの場合に重要です。
全文: エンジニアと AI アーキテクト向けの使用法
(および PDF: docs/pdf/01_USAGE_FOR_ENGINEERS_AND_ARCHITECTS.pdf )。
エンジニアはこれをどのように統合すべきか
from prismmanifest import KeyRing 、 PrismManifestPipeline 、 enforce_group3_boundary 、 GateDecision
プリズムマニフェストから。ルーターインポート DocumentPackage 、 IntentRouter
プリズムマニフェストから。統合インポートdemo_capital_gains_dag
キーリング = キーリング 。生成 ( key_id = "local-dev-ed25519" )
パッケージ = DocumentPackage (
doc_id = "1040.txt" ,
ページ = [
"フォーム 1040 2024 課税年度\n"
「行 1 総収入: $470,000.00 \n」
「行 11 調整済み総収入: $450,000.00 \n」
]、
form_type = "IRS_FORM_1040" ,
税年 = 2024 、
)
# ルート → インジェスト + スパン抽出 (パターン B ポインタ)
ルーティング = IntentRouter ()。実行（パッケージ）
# 検証、ゲートの決定、マニフェストへの署名
パイプライン = PrismManifestPipeline (キーリング)
結果 = パイプライン。 run_on_evidence (
証拠 = ルーティング済み。証拠、
抽出 = ルーティング済み。抽出、
)
# グループ 3 のハード境界 — これはセキュリティ ゲートです
ゲート = enforce_group3_boundary (
結果。マニフェスト、
public_keys = キーリング 、
Expected_dag_id = "capital_gains_v3" ,
)
もしゲートがあれば。決定は GateDecision です。同意します:
レシート=demo_capital_gains_dag (ゲート.マニフェスト)
印刷（再

領収書）
それ以外の場合:
print (ゲートの決定、ゲートのメッセージ)
人間のエスカレーション
プリズムマニフェストから。監査インポート EscalationQueue
prismmanifest からインポート PrismManifestPipeline 、 KeyRing
キーリング = キーリング 。生成する()
キュー = EscalationQueue ( ".escalation" )
パイプライン = PrismManifestPipeline (キーリング、エスカレーション_キュー = キュー)
# PASS_WITH_HUMAN マニフェストは署名され、キューが接続されると自動的にエンキューされます。
デコレーター (DX のみ - セキュリティ境界ではない)
prismmanifestインポートparameter_gatedから、KeyRing
キーリング = キーリング 。ロード ( ".keys" )
@parameter_gated ( public_keys = キーリング 、 Expected_dag_id = "capital_gains_v3" )
def run_dag ( * 、マニフェスト ):
マニフェストを返します。フィールド [ 0 ]。値_固定_マイクロ
運用信頼は引き続き enforce_group3_boundary / gRPC / C++ / インプロセス prismmanifest_c を通過する必要があります。
プリズムマニフェストから。 binary_codec インポート encode_manifest
プリズムマニフェストから。ゲートインポート enforce_group3_boundary
buf = encode_manifest ( signed_manifest )
result = enforce_group3_boundary ( buf 、 public_keys = キーリング 、 Expected_dag_id = "capital_gains_v3" )
スキーマ: schemas/prismmanifest.fbs 。
パイプラインは「良好に見えた」ため、enforce_group3_boundary をスキップします。
LLM で出力された $ 文字列を DAG に直接フィードします。
@parameter_gated のみを境界として扱います
cuda_sim / SKIP を CUDA 検証済みとして市場に出す
デュアル OCR コンセンサス (PDF テキスト + レイアウト再トークナイザー、オプションの Tesseract) を使用して証拠を取り込みます。
Ground は、文字通りのスパンとアンカーを形成することを主張します (お金の価値を生成することはありません)。
クォーラム、OCR フロア (≥ 0.98)、および妥当性によって PASS / PASS_WITH_HUMAN / REFUSE を決定します。
ParameterManifest (Ed25519) に署名し、必要に応じて人間によるレビューをエスカレーションします。
DAG (Python、gRPC、C++ FlatBuffer、またはインプロセス DLL) を実行する前に、グループ 3 の境界を強制します。
FinancePackBench と FinancePackBench-G4

合成 SLA / 敵対的スイートを提供します。
ステータス
意味
パス
スパン接地、妥当性 OK、異論なし、OCR ≥ 0.98、アンカー OK
PASS_WITH_HUMAN
軽微な意見の相違 (≤ 1,000)、OCR が低い、またはアンカーが未検証
拒否する
根拠がない、妥当性の欠如、または重大な不一致 (> 1,000 ドル)
GateDecision (グループ 3 境界)
決定
いつ
受け入れる
PASS 、または有効な人間の承認トークンを使用して PASS_WITH_HUMAN をクリアしました
ACCEPT_PENDING_HUMAN
PASS_WITH_HUMAN が審査待ちです
拒否する
REFUSE 、不良/クリアランスの欠落、または認証/鮮度/再生/DAG の失敗
証明、鮮度 ( signed_at_unix スキュー、デフォルト 300 秒)、およびリプレイ ( ReplayGuard ) の失敗により、 GateError が発生します。
# キー
prismmanifest-gate gen-keys --out .keys --key-id local-dev-ed25519
prismmanifest-gate gen-hsm-key --out .hsm --key-id prod-ed25519
prismmanifest-gate gen-kms-key --out .kms --mode local --key-id kms-dev
# Azure: prismmanifest-gate gen-kms-key --out .kms --mode azure --kms-key-id <vault-key-url>
# 署名/検証
prismmanifest-gate サイン --keys .keys --manifest マニフェスト.json --out signed.json
prismmanifest-gate サイン --keys .keys --manifest マニフェスト.json --out signed.fbs -- flatbuffer
prismmanifest-gate verify --keys .keys --manifest signed.json --dag-id Capital_gains_v3
# ベンチと証拠
プリズムマニフェストゲートベンチ -- パッケージ 500
prismmanifest-gate bench --require-cuda --packages 40
prismmanifest-gate bench-manifest-parity --packages 500 --require-cuda
prismmanifest-gate bench-perf --out reports/perf
prismmanifest-gate bench-customer-pdf --corpus .corpus/customer --seed-synthetic
prismmanifest-gate コンプライアンス --out レポート/コンプライアンス
prismmanifest-gate パイロットパック --out reports/pilot_pack
prismmanifest-gate g4-suite --out reports/g4 --fuzz 200
# 操作 / レビュー
prismmanifest-gate Audit-replay --store .audit --receipt < id

> --keys .keys
prismmanifest-gate エスカレーションリスト --queue .escalation
prismmanifest-gate review-ui --queue .escalation --keys .keys --bind 127.0.0.1:8766
gRPC グループ 3 サービス
プロト: proto/prismmanifest_gate.proto
サービス: prismmanifest.v1.Group3Gate — VerifyManifest 、 ExecuteDag
プリズムマニフェストから。証明書インポートキーリング
プリズムマニフェストから。 grpc_servicerインポートサーブ
serve ( KeyRing .load ( ".keys" )、bind = "[::]:50051" )
C++ / インプロセスゲート（オプション）
# Windows
powershell -ファイル scripts/build_cpp_gate.ps1
# 生成: prismmanifest_gate_enforce.exe + prismmanifest_c.dll
prismmanifest_gate_enforce signed.fbs public.pem < key_id > Capital_gains_v3
Python はプロセスを生成せずに共有ライブラリを呼び出すことができます。
プリズムマニフェストから。 cpp_bridge import enforce_fb 、bridge_status
print ( Bridge_status ()) # PRISMMANIFEST_C_DLL / ビルドが存在する場合は inprocess_available
Ops パッケージ化 (パイロット展開)
「運用パッケージ化」 = サービス内でゲートを実行する方法 (アルゴリズムではありません):
キー (ローカル / ソフトウェア HSM / Azure KV)
RBAC、タイムアウト、冪等性、メトリクス ( prismmanifest.ops )
Python パッケージの横にあるオプションの C++ / CUDA
docs/handoff/PILOT_DEPLOY.md を参照してください。
ドキュメント (Markdown + PDF)
ドクター
役割
USAGE_FOR_ENGINEERS_AND_ARCHITECTS.md
PrismManifestの使い方
PRISMMANIFEST_SYSTEM_DESIGN.md
アーキテクチャと脅威モデル (権威)
PRISMMANIFEST_IMPLEMENTATION_PLAN.md
フェーズと SLA
PRISMMANIFEST_MASTER_SPECIFICATION.md
エグゼクティブインデックス
FINANCEPACKBENCH_G4_ADVERSARIAL_SUITE.md
敵対的スイート
FINANCEPACKBENCH_PROMPT_INJECTION_SEMANTICS.md
不活性注入が依然として合格できる理由
ドキュメント/pdf/
生成された PDF パック (完全に結合された PDF を含む)
PDF を再生成します。
pip install " prismmanifest[docs] " # または: pip install markdown fpdf2
Python スクリプト/build_docs_pdf.py
レイアウト
prismmanifest/Python パッケージ

cpp/gate/ C++ 正規化 + FlatBuffer enforce + prismmanifest_c
cuda/kernels/ 実験的な .cu カーネル
スキーマ/prismmanifest.fbs
proto/gRPC Group3Gate
テスト/pytest
docs/ 仕様 + 使用方法 + PDF 出力
スクリプト/プロト、C++ ビルド、CUDA シム、PDF ビルダー
レポート/生成されたプルーフ アーティファクト (pip インストールには必要ありません)
正直な範囲と準備状況
ラベル
意味
パイロットOSS
合成 + 敵対的 FA=0 テスト中。 Py/C++/CUDA 決定パリティ。インプロセス C++ パス
製造上の主張
ライブ顧客 FAX/スキャンされたコーパスが同じ FA/SLA バーに合格する必要があります
また:
セキュリティ境界 = enforce_group3_boundary / gRPC / C++ / prismmanifest_c — @parameter_gated だけではありません。
CUDA: 実際のパリティには GPU + Numba が必要です。 cuda_sim は CI セルフチェックのみです。
HSM: gen-hsm-key は、PKCS#11 ハードウェアではなく、保存時に暗号化されるソフトウェアです。
プロンプトインジェクション: PrismManifest は実行トラストゲートです。「インジェクション防御」を主張する前に、インジェクションセマンティクスのドキュメントを参照してください。
公開ノート: docs/PUBLISHING.md 。
チャンネル
用途
ディスカッション
設計に関する質問、パイロットのフィードバック、アーキテクチャ
問題点
バグと具体的な機能リクエスト
セキュリティ.md
非公開の脆弱性レポート
貢献.md
開発セットアップと PR への期待
ライセンス
Apache License 2.0 — 「 LICENSE 」を参照してください。
著者: アミン・パルヴァ (insightits.info@gma)

[切り捨てられた]

## Original Extract

Zero-trust tool-argument gate: signed ParameterManifest before deterministic Group 3 DAGs (PyPI: prismmanifest) - insightitsGit/PrismManifest

GitHub - insightitsGit/PrismManifest: Zero-trust tool-argument gate: signed ParameterManifest before deterministic Group 3 DAGs (PyPI: prismmanifest) · GitHub
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
insightitsGit
/
PrismManifest
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
20 Commits 20 Commits .github .github cpp/ gate cpp/ gate cuda/ kernels cuda/ kernels docs docs prismmanifest prismmanifest proto proto reports reports schemas schemas scripts scripts tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md SECURITY.md SECURITY.md azure-pipelines.yml azure-pipelines.yml pyproject.toml pyproject.toml View all files Repository files navigation
Zero-trust tool-argument gate for deterministic AI tool execution.
(Formerly ParamGate — same design; package prismmanifest , CLI prismmanifest-gate .)
PrismManifest sits between probabilistic extractors (LLMs, OCR, table parsers) and
deterministic Group 3 compute DAGs. Unverified money values never enter the DAG.
Only an Ed25519-signed ParameterManifest that clears the Group 3 boundary
is allowed through.
Deterministic Engine + Unverified Probabilistic Input = Deterministic Wrong Answer
Keywords: tool argument gate, ParameterManifest, digit drop prevention, zero-trust LLM tool args, deterministic DAG boundary, OCR money field authorization, PrismManifest
Probabilistic systems (LLMs, OCR) are excellent at proposing where a number is.
They must not be trusted to authorize the dollar amount that enters a tax engine,
ledger, or underwriting model. PrismManifest is that authorization boundary.
pip install prismmanifest
With optional extras:
pip install " prismmanifest[dev] "
pip install " prismmanifest[cuda] " # NVIDIA GPU + Numba
pip install " prismmanifest[kms-azure] " # Azure Key Vault envelope keys
pip install " prismmanifest[kms-aws] " # AWS KMS envelope (optional)
From source (editable):
git clone https://github.com/insightitsGit/PrismManifest.git
cd PrismManifest
python -m pip install -e " .[dev] "
python -m pytest -q
CLI after install: prismmanifest-gate .
How an AI architect should use this
Treat LLMs / OCR / parsers as untrusted .
Put PrismManifest after extraction and before any calculator, underwriting, or ledger tool that consumes dollar amounts.
Route outcomes: ACCEPT → run DAG · ACCEPT_PENDING_HUMAN → review · REJECT → stop.
Never let model-generated money text be the tool argument — only evidence-bound, signed manifests.
LLM / OCR → PrismManifest → signed ParameterManifest → Group 3 DAG
Extra LLM benches are optional evidence . They do not change the core gate design.
Finding real customer documents matters more for production claims than multi-model FA studies.
Full write-up: Usage for Engineers & AI Architects
(and PDF: docs/pdf/01_USAGE_FOR_ENGINEERS_AND_ARCHITECTS.pdf ).
How an engineer should integrate this
from prismmanifest import KeyRing , PrismManifestPipeline , enforce_group3_boundary , GateDecision
from prismmanifest . router import DocumentPackage , IntentRouter
from prismmanifest . integrations import demo_capital_gains_dag
keyring = KeyRing . generate ( key_id = "local-dev-ed25519" )
package = DocumentPackage (
doc_id = "1040.txt" ,
pages = [
"Form 1040 Tax Year 2024 \n "
"Line 1 Gross income: $470,000.00 \n "
"Line 11 Adjusted gross income: $450,000.00 \n "
],
form_type = "IRS_FORM_1040" ,
tax_year = 2024 ,
)
# Route → ingest + span extraction (Pattern B pointers)
routed = IntentRouter (). run ( package )
# Verify, decide gate, sign manifest
pipeline = PrismManifestPipeline ( keyring )
result = pipeline . run_on_evidence (
evidence = routed . evidence ,
extraction = routed . extraction ,
)
# Group 3 hard boundary — this is the security gate
gate = enforce_group3_boundary (
result . manifest ,
public_keys = keyring ,
expected_dag_id = "capital_gains_v3" ,
)
if gate . decision is GateDecision . ACCEPT :
receipt = demo_capital_gains_dag ( gate . manifest )
print ( receipt )
else :
print ( gate . decision , gate . message )
Human escalation
from prismmanifest . audit import EscalationQueue
from prismmanifest import PrismManifestPipeline , KeyRing
keyring = KeyRing . generate ()
queue = EscalationQueue ( ".escalation" )
pipeline = PrismManifestPipeline ( keyring , escalation_queue = queue )
# PASS_WITH_HUMAN manifests are signed and auto-enqueued when a queue is attached.
Decorator (DX only — not the security boundary)
from prismmanifest import parameter_gated , KeyRing
keyring = KeyRing . load ( ".keys" )
@ parameter_gated ( public_keys = keyring , expected_dag_id = "capital_gains_v3" )
def run_dag ( * , manifest ):
return manifest . fields [ 0 ]. value_fixed_micro
Production trust must still go through enforce_group3_boundary / gRPC / C++ / in-process prismmanifest_c .
from prismmanifest . binary_codec import encode_manifest
from prismmanifest . gate import enforce_group3_boundary
buf = encode_manifest ( signed_manifest )
result = enforce_group3_boundary ( buf , public_keys = keyring , expected_dag_id = "capital_gains_v3" )
Schema: schemas/prismmanifest.fbs .
Skip enforce_group3_boundary because the pipeline “looked good”
Feed LLM-printed $ strings straight into the DAG
Treat @parameter_gated alone as the boundary
Market cuda_sim / SKIP as CUDA-validated
Ingest evidence with dual-OCR consensus (PDF text + layout re-tokenizer; optional Tesseract).
Ground claims to verbatim spans and form anchors (no generative money values).
Decide PASS / PASS_WITH_HUMAN / REFUSE via quorum, OCR floor (≥ 0.98), and plausibility.
Sign a ParameterManifest (Ed25519) and optionally escalate human review.
Enforce the Group 3 boundary before any DAG runs — Python, gRPC, C++ FlatBuffer, or in-process DLL.
FinancePackBench and FinancePackBench-G4 provide synthetic SLA / adversarial suites.
Status
Meaning
PASS
Span-grounded, plausibility OK, no disagreement, OCR ≥ 0.98, anchors OK
PASS_WITH_HUMAN
Immaterial disagreement (≤ $1k), low OCR, or anchors unverified
REFUSE
Not grounded, plausibility failure, or material disagreement (> $1k)
GateDecision (Group 3 boundary)
Decision
When
ACCEPT
PASS , or cleared PASS_WITH_HUMAN with valid human approval token
ACCEPT_PENDING_HUMAN
PASS_WITH_HUMAN awaiting review
REJECT
REFUSE , bad/missing clearance, or attestation/freshness/replay/dag failure
Attestation, freshness ( signed_at_unix skew, default 300s), and replay ( ReplayGuard ) failures raise GateError .
# Keys
prismmanifest-gate gen-keys --out .keys --key-id local-dev-ed25519
prismmanifest-gate gen-hsm-key --out .hsm --key-id prod-ed25519
prismmanifest-gate gen-kms-key --out .kms --mode local --key-id kms-dev
# Azure: prismmanifest-gate gen-kms-key --out .kms --mode azure --kms-key-id <vault-key-url>
# Sign / verify
prismmanifest-gate sign --keys .keys --manifest manifest.json --out signed.json
prismmanifest-gate sign --keys .keys --manifest manifest.json --out signed.fbs --flatbuffer
prismmanifest-gate verify --keys .keys --manifest signed.json --dag-id capital_gains_v3
# Benches & proof
prismmanifest-gate bench --packages 500
prismmanifest-gate bench --require-cuda --packages 40
prismmanifest-gate bench-manifest-parity --packages 500 --require-cuda
prismmanifest-gate bench-perf --out reports/perf
prismmanifest-gate bench-customer-pdf --corpus .corpus/customer --seed-synthetic
prismmanifest-gate compliance --out reports/compliance
prismmanifest-gate pilot-pack --out reports/pilot_pack
prismmanifest-gate g4-suite --out reports/g4 --fuzz 200
# Ops / review
prismmanifest-gate audit-replay --store .audit --receipt < id > --keys .keys
prismmanifest-gate escalation-list --queue .escalation
prismmanifest-gate review-ui --queue .escalation --keys .keys --bind 127.0.0.1:8766
gRPC Group 3 service
Proto: proto/prismmanifest_gate.proto
Service: prismmanifest.v1.Group3Gate — VerifyManifest , ExecuteDag
from prismmanifest . attestation import KeyRing
from prismmanifest . grpc_servicer import serve
serve ( KeyRing . load ( ".keys" ), bind = "[::]:50051" )
C++ / in-process gate (optional)
# Windows
powershell -File scripts/build_cpp_gate.ps1
# Produces: prismmanifest_gate_enforce.exe + prismmanifest_c.dll
prismmanifest_gate_enforce signed.fbs public.pem < key_id > capital_gains_v3
Python can call the shared library without process spawn:
from prismmanifest . cpp_bridge import enforce_fb , bridge_status
print ( bridge_status ()) # inprocess_available when PRISMMANIFEST_C_DLL / build present
Ops packaging (pilot deploy)
“Ops packaging” = how you run the gate in a service (not the algorithm):
Keys (local / software-HSM / Azure KV)
RBAC, timeouts, idempotency, metrics ( prismmanifest.ops )
Optional C++ / CUDA beside the Python package
See docs/handoff/PILOT_DEPLOY.md .
Documentation (Markdown + PDF)
Doc
Role
USAGE_FOR_ENGINEERS_AND_ARCHITECTS.md
How to use PrismManifest
PRISMMANIFEST_SYSTEM_DESIGN.md
Architecture & threat model ( authority )
PRISMMANIFEST_IMPLEMENTATION_PLAN.md
Phases & SLAs
PRISMMANIFEST_MASTER_SPECIFICATION.md
Executive index
FINANCEPACKBENCH_G4_ADVERSARIAL_SUITE.md
Adversarial suite
FINANCEPACKBENCH_PROMPT_INJECTION_SEMANTICS.md
Why inert injection can still PASS
docs/pdf/
Generated PDF pack (including full combined PDF)
Regenerate PDFs:
pip install " prismmanifest[docs] " # or: pip install markdown fpdf2
python scripts/build_docs_pdf.py
Layout
prismmanifest/ Python package
cpp/gate/ C++ canonicalize + FlatBuffer enforce + prismmanifest_c
cuda/kernels/ Experimental .cu kernels
schemas/ prismmanifest.fbs
proto/ gRPC Group3Gate
tests/ pytest
docs/ Specs + usage + PDF output
scripts/ proto, C++ build, CUDA shim, PDF builder
reports/ Generated proof artifacts (not required for pip install)
Honest scope & readiness
Label
Meaning
Pilot OSS
Synthetic + adversarial FA=0 under test; Py/C++/CUDA decision parity; in-process C++ path
Production claim
Requires your live customer fax/scanned corpus to pass the same FA/SLA bar
Also:
Security boundary = enforce_group3_boundary / gRPC / C++ / prismmanifest_c — not @parameter_gated alone.
CUDA: real parity needs GPU + Numba; cuda_sim is CI self-check only.
HSM: gen-hsm-key is software encrypted-at-rest, not PKCS#11 hardware.
Prompt injection: PrismManifest is an execution trust gate — see the injection semantics doc before claiming “injection defense.”
Publishing notes: docs/PUBLISHING.md .
Channel
Use for
Discussions
Design questions, pilot feedback, architecture
Issues
Bugs and concrete feature requests
SECURITY.md
Private vulnerability reports
CONTRIBUTING.md
Dev setup and PR expectations
License
Apache License 2.0 — see LICENSE .
Author: Amin Parva ( insightits.info@gma

[truncated]
