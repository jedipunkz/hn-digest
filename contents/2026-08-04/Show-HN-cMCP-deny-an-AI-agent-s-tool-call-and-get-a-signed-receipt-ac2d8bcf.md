---
source: "https://github.com/agentrust-io/cmcp"
hn_url: "https://news.ycombinator.com/item?id=49172545"
title: "Show HN: cMCP, deny an AI agent's tool call and get a signed receipt"
article_title: "GitHub - agentrust-io/cmcp: cMCP: Confidential MCP Gateway. Hardware-attested policy enforcement for MCP tool calls. · GitHub"
author: "mosiddi"
captured_at: "2026-08-04T18:23:26Z"
capture_tool: "hn-digest"
hn_id: 49172545
score: 5
comments: 0
posted_at: "2026-08-04T18:06:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: cMCP, deny an AI agent's tool call and get a signed receipt

- HN: [49172545](https://news.ycombinator.com/item?id=49172545)
- Source: [github.com](https://github.com/agentrust-io/cmcp)
- Score: 5
- Comments: 0
- Posted: 2026-08-04T18:06:08Z

## Translation

タイトル: HN: cMCP を表示、AI エージェントのツール呼び出しを拒否し、署名済みの領収書を取得する
記事のタイトル: GitHub - Agentrust-io/cmcp: cMCP: Confidential MCP ゲートウェイ。 MCP ツール呼び出しに対するハードウェア認証ポリシーの適用。 · GitHub
説明: cMCP: 機密 MCP ゲートウェイ。 MCP ツール呼び出しに対するハードウェア認証ポリシーの適用。 - エージェントラスト-io/cmcp

記事本文:
GitHub - Agentrust-io/cmcp: cMCP: 機密 MCP ゲートウェイ。 MCP ツール呼び出しに対するハードウェア認証ポリシーの適用。 · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
エージェントラスト-io
/
cmcp
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
314 コミット 314 コミット .github .github ベンチマーク ベンチマーク ドキュメント ドキュメント 例 例 実験 実験

■ ガバナンス ガバナンス オーバーライド オーバーライド スキーマ スキーマ スクリプト スクリプト src src テスト テスト .gitignore .gitignore ADOPTERS.md ADOPTERS.md CHANGELOG.md CHANGELOG.md CHARTER.md CHARTER.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE LIMITATIONS.md LIMITATIONS.md MAINTAINERS.md MAINTAINERS.md NOTICE NOTICE PRIVACY.md PRIVACY.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md STATUS.md STATUS.md docker-compose.yml docker-compose.yml mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml 要件-docs.txt 要件-docs.txt robots.txt robots.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
cMCP: 機密 MCP ランタイム
TEE 内で MCP ツール ポリシーを適用します。TEE が管理するエージェントはアクセスできません。
クイックスタート ·
建築・
構成・
CLI・
変更履歴
開発者プレビュー - 2026 年 6 月 23 日の Confidential Computing Summit で開始されました。v1.0 より前に重大な変更が含まれる可能性があります。現在出荷されるものとロードマップ上のものについては、STATUS.md を参照してください。
cMCP (Confidential MCP Runtime) は、MCP を安全かつ機密に実行する方法です。これは、ハードウェアの信頼された実行環境 (TEE) 内で MCP ツール呼び出しポリシーを強制するオープンソース ゲートウェイです。すべてのツール呼び出しはインターセプトされ、Cedar ポリシー バンドルに対して評価され、管理されているプロセスが到達できない場所で強制されます。各セッションでは署名付きの TRACE クレームが生成され、検証者はオペレーターを信頼せずにチェックします。ゲートウェイが TEE で実行されている場合はハードウェアで証明され、ソフトウェア モードでは署名のみが行われます。 MCP の安全なバージョンをお探しの場合は、これがそのための AgenTrust ランタイムです。
TL;DR - エージェントに cMCP ゲートウェイを指定します。 TEE 内の Cedar ポリシーに対してすべてのツール呼び出しを評価し、何をブロックまたは編集しますか

ポリシーはこれを否定し、証拠として改ざんが明らかな TRACE クレームを発行します。 pip install cmcp-runtime を実行し、ハードウェアを必要としないソフトウェア モードで起動します。
エージェントは、Snowflake、Salesforce、多数の API を呼び出します。これらの通話で顧客のデータが漏洩するのを防ぐものは何でしょうか?規制当局が尋ねたら、そうではないと証明できますか?
エージェントがツールを呼び出します。ポリシー エンジンは許可すると言っています。ツール呼び出しは完了します。
いずれも、ポリシー エンジン自体が侵害されていないことを証明するものではありません。ソフトウェアのみの MCP ガバナンスでは次のことは保証できません。
ディスク上の Cedar ポリシーが実行されたものです。不正な管理者は、承認後にバンドルを交換することができます。ハッシュ チェックは、管理者が制御するのと同じ OS 内で実行されます。
許可/拒否の決定はメモリ内で反転されませんでした。評価者のサプライ チェーン CVE は、攻撃者と同じアドレス空間で実行されます。
監査ログには、実際に何が起こったかが反映されます。ソフトウェア署名キーを保有する当事者は、事後的に有効な監査チェーンを再構築できます。
ツール呼び出しを管理するコントロール プレーンは、管理するプロセスがアクセスできない場所で実行する必要があります。
MCP ツール呼び出しに対するハードウェア認証ポリシーの適用。すべてのツール呼び出しはインターセプトされ、Cedar ポリシー バンドルに対して評価され、信頼された実行環境 (TEE) 内で実行されているポリシー エンジンによって強制されます。ポリシー バンドル ハッシュは、コードが実行される前にハードウェア構成証明レポートで測定されます。
トンネルベースの接続ソリューションとは異なり、cMCP ランタイムは TEE 内でツール呼び出しペイロードを処理します。接続プロバイダーは平文ではなく暗号文を認識します。エンクレーブから離れるのは、署名された TRACE クレームだけです。
pip インストール cmcp-runtime
cmcp-config.yaml を作成します。
証明書:
プロバイダー：自動
enforcement_mode : Advisory # Advisory により、初回実行のチューニングが容易になります。デフォルトは「強制」です
ポリシーバンドルのパス: ./policies/
カタ

ログパス: ./catalog.json
ゲートウェイを起動します。
CMCP_DEV_MODE=1 cmcp start --config cmcp-config.yaml
ツール呼び出しを行います。
カール -X POST http://localhost:8443/mcp \
-H " Content-Type: application/json " \
-d ' {"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"salesforce.contacts","arguments":{"query":"Acme Corp"},"_cmcp":{"session_id":"s1","workflow_id":"demo-agent"}}} '
完全なウォークスルーについては、docs/quickstart.md を参照してください: Cedar ポリシー、ツール カタログ、最初の TRACE 要求、および検証 (ハードウェア TEE は必要ありません)。
エージェントは、すべてのツール呼び出しを MCP サーバーに直接送信するのではなく、cMCP ゲートウェイに送信します。
起動時に、ゲートウェイはハードウェア構成証明レポートに含まれる Cedar ポリシー バンドル ハッシュを測定します。この測定の前にコードは実行されません。
受信した各ツール呼び出しは、TEE 内で実行されている Cedar ポリシー エンジンによって評価されます。結果は、許可、拒否、または編集になります。呼び出しとその決定は、ハードウェアで保護された監査チェーンに追加されます。
セッションの終了時に、ゲートウェイは TRACE クレームを生成します。これは、実行されたツール、各呼び出しを決定したポリシー、および完全な監査チェーンを記録する、署名されたハードウェア証明済みの成果物です。検証者は、オペレータを信頼せずにこれをチェックします。
エージェント -> cMCP ランタイム -> Cedar ポリシー エンジン (TEE) -> ツール
|
GatewayClaim (TRACE プロファイル)
+-- トレース.eat_profile
+-- トレース.ランタイム.プラットフォーム + 測定
+-- トレース.ポリシー.バンドル_ハッシュ
+--trace.cnf.jwk (Ed25519 確認キー)
+-- ゲートウェイ.audit_chain (ルート/ヒント/長さ)
+-- 署名 (正規 JSON 上の Ed25519)
ハードウェアプロバイダー
プロバイダー
プラットフォーム
保証
注意事項
tpm
TPM 2.0 / vTPM (Azure、AWS、GCP Trusted Launch)
中
ローカルTPM見積もり
セブ-SNP
AMD SEV-SNP (Azure DCasv5、AWS C6a Nitro)
高
AMD KDS
tdx
インテル TDX (Azure DCedsv5、GCP C3)
高
インテル PCS
GPU-CC (v0.2)
NVIDIA H100/H200/ブラックウェル (CC モード)
高

NVIDIA リモート認証サービス (NRAS)
不透明 (オプトイン)
OPAQUE Confidential ランタイム
該当なし (まだ実装されていません)
プレースホルダー: 自動検出から除外されます。明示的に選択すると、実装されていないエラーが発生します
プロバイダーの自動検出プローブの順序: azure-cvm -> tpm -> sev-snp -> tdx 。 detect() が成功した最初のプロバイダーが選択されます。 opaque はまだ実装されていないプレースホルダーです。自動検出から除外されており、これを選択すると、黙って失敗するのではなく、明示的に ATTESTATION_PROVIDER_NOT_IMPLEMENTED が発生します。ハードウェア プロバイダーが検出されない場合、ゲートウェイは CMCP_DEV_MODE=1 (証明されていないソフトウェアのみのフォールバック) でのみ起動し、それ以外の場合は起動を拒否します。
cmcp_runtime から。構成インポート TEEProvider
# 自動検出 (デフォルト)
# attestation.provider: auto -> azure-cvm -> tpm -> sev-snp -> tdx
# (ソフトウェアのみは CMCP_DEV_MODE=1 でのみ使用されます)
# 明示的なハードウェアの選択
# attestation.provider: sev-snp
# OPAQUE マネージド ランタイム (オプトインのみ; まだ実装されていません)
# OPAQUE_ATTESTATION_URL=https://... cmcp start --config cmcp-config.yaml
強制モード
モード
行動
ユースケース
強制する
ポリシーは HTTP 403 を返すことを拒否します。電話が転送されない
生産
勧告
ポリシーの拒否はログに記録されます。通話収入
最初の展開、ポリシーの調整
沈黙
ポリシーは評価されますが、何もログに記録されず、ブロックされません。
ベースライニング
デフォルトは強制です。アドバイザリー モードを使用するには、cmcp-config.yaml で enforcement_mode: Advisory を設定します。
cmcp-config.yaml の完全なリファレンス:
証明書:
プロバイダー: 自動 # 自動 | tpm |セブ-SNP | tdx |不透明 |ソフトウェアのみ
強制モード : 強制 # 強制 |アドバイス |沈黙
validity_seconds : 86400 # 認証の鮮度ウィンドウ (デフォルト: 24 時間)
staleness_policy:fail_closed #fail_closed |警告のみ
Expected_measurement : ~ # 特定の PCR/測定を固定します (opti

オナール)
policy_bundle_path : events/ # .cedar ファイルとmanifest.jsonを含むディレクトリ
category_path : category.json # 承認されたツール カタログ
listen_addr : " 0.0.0.0:8443 "
max_response_size_bytes : 2097152 # 2 MB デフォルト
ポリシーリロード間隔秒 : 0 # 0 = 無効。ポリシーを更新するには再起動が必要です
環境変数:
コマンド
フラグ
説明
cmcp の開始
--config PATH (必須)
ゲートウェイを起動する
cmcp 検証構成
--config PATH (必須)
cmcp-config.yaml を起動せずに検証する
cmcp 検証バンドル
--bundle-path PATH (必須)、--expected-hash sha256:<hex> (必須)
導入前に Cedar バンドルのハッシュを検証する
cmcp 検証
CLAIM_FILE (必須); --policy-hash 、 --catalog-hash 、 --max-age 、 --trusted-key 、 --audit-bundle 、 --agent-manifest 、 --agent-manifest-trust-anchor
署名された TRACE クレーム (署名、スキーマ、鮮度、監査チェーン、ピン留めされたハッシュ) を検証する
トレースクレーム
GatewayClaim は、監査人、規制当局、または下流の検証者に渡される証明の単位です。これはセッションごと (または呼び出しごと、構成可能) に生成され、TEE から離れることのない鍵で署名されます。
(この表は、最もよく使用されるフィールドの概要です。)
cmcp_verify ライブラリを使用した検証では、オペレーターを信頼する必要はありません。検証者は、TEE バインドされたキーに対して署名をチェックし、承認された値に対してポリシー バンドル ハッシュをチェックし、内部一貫性について監査チェーンをチェックします。
標準的なスキーマは schemas/trace-claim.schema.json で、 docs/quickstart.md に完全な例が示されています。完全な検証プロトコルについては、docs/spec/verification-library.md および TRACE 仕様を参照してください。
標準
適用範囲
OWASP エージェント AI トップ 10
MCP10 (ツール呼び出しによるデータ漏洩)、MCP02 (未承認ツール)、MCP08 (証明可能なガバナンス)、MCP04 (サプライ チェーン)
NIST SP 800-207
内部のポリシー決定ポイント

eTシャツ。 Workload Identity に対する暗黙的な信頼がない
EU AI 法第 2 条12、15
決定ごとの監査記録 (第 12 条); TEE に基づくサイバーセキュリティ管理 (第 15 条)
ドラアート。 9
認証チェーン。 Gateway.audit_chain を介した監査ログの保持
ラット/食べる RFC 9711
GatewayClaim は EAT です。 Eat_profile フィールドは TRACE プロファイルを識別します
セキュリティ
ツール
チェック内容
ラフ
すべての PR で lint をスタイル設定してインポートする
盗賊
すべての PR での Python セキュリティ リンティング
pip監査
すべての PR での依存関係の脆弱性スキャン
マイピー
すべての PR での静的型チェック
コードQL
Python SAST、セキュリティ拡張クエリ、毎週
OpenSSF スコアカード
毎週のスコアリング、SARIF アップロード
脆弱性レポートと対応 SLA については、SECURITY.md を参照してください。 APM ペイロード キャプチャ、ランタイム設定インジェクション、フェーズ 1 で終了しない P4.1 サプライ チェーン (typosquat) の残留リスクなど、明示的な範囲の境界については、LIMITATIONS.md を参照してください。
ページ
説明
docs/quickstart.md
ゼロから最初の TRACE Claim まで 30 分以内
docs/configuration.md
すべてのフィールドとデフォルトを含む完全な構成リファレンス
docs/SPEC.md
製品仕様: 問題分類、アーキテクチャ、カバレッジ マトリックス
docs/spec/threat-model.md
STRIDE分析、敵対者モデル、残留リスク
docs/spec/cedar-policy.md
セド

[切り捨てられた]

## Original Extract

cMCP: Confidential MCP Gateway. Hardware-attested policy enforcement for MCP tool calls. - agentrust-io/cmcp

GitHub - agentrust-io/cmcp: cMCP: Confidential MCP Gateway. Hardware-attested policy enforcement for MCP tool calls. · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
agentrust-io
/
cmcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
314 Commits 314 Commits .github .github benchmarks benchmarks docs docs examples examples experiments experiments governance governance overrides overrides schemas schemas scripts scripts src src tests tests .gitignore .gitignore ADOPTERS.md ADOPTERS.md CHANGELOG.md CHANGELOG.md CHARTER.md CHARTER.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE LIMITATIONS.md LIMITATIONS.md MAINTAINERS.md MAINTAINERS.md NOTICE NOTICE PRIVACY.md PRIVACY.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md STATUS.md STATUS.md docker-compose.yml docker-compose.yml mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml requirements-docs.txt requirements-docs.txt robots.txt robots.txt View all files Repository files navigation
cMCP: Confidential MCP Runtime
Enforce MCP tool policy inside a TEE, where the agent it governs cannot reach it
Quick Start ·
Architecture ·
Configuration ·
CLI ·
Changelog
Developer Preview - launched at the Confidential Computing Summit, June 23 2026. May have breaking changes before v1.0. See STATUS.md for exactly what ships today versus what is on the roadmap.
cMCP (Confidential MCP Runtime) is the secure, confidential way to run MCP: an open-source gateway that enforces MCP tool-call policy inside a hardware Trusted Execution Environment (TEE). Every tool call is intercepted, evaluated against a Cedar policy bundle, and enforced where the process it governs cannot reach it. Each session produces a signed TRACE Claim that a verifier checks without trusting the operator, hardware-attested when the gateway runs in a TEE and signed-only in software mode. If you are looking for a secure version of MCP, this is the AgenTrust runtime for it.
TL;DR - Point your agent at the cMCP Gateway. It evaluates every tool call against a Cedar policy inside a TEE, blocks or redacts what the policy denies, and emits a tamper-evident TRACE Claim as proof. Run pip install cmcp-runtime and start in software mode with no hardware required.
Your agent calls Snowflake, Salesforce, a dozen APIs. What stops it from leaking a customer's data on one of those calls? If a regulator asks, could you prove it didn't?
An agent calls a tool. The policy engine says allow. The tool call goes through.
None of that proves the policy engine itself was not compromised. Software-only MCP governance cannot guarantee:
The Cedar policy on disk is the one that ran. A rogue admin can swap the bundle after approval; the hash check runs inside the same OS the admin controls.
The allow/deny decision was not flipped in memory. A supply chain CVE in the evaluator runs in the same address space as the attacker.
The audit log reflects what actually happened. Any party holding the software signing key can reconstruct a valid audit chain after the fact.
The control plane that governs tool calls must run where it cannot be reached by the process it governs.
Hardware-attested policy enforcement for MCP tool calls. Every tool call is intercepted, evaluated against a Cedar policy bundle, and enforced by a policy engine running inside a Trusted Execution Environment (TEE). The policy bundle hash is measured into the hardware attestation report before any code runs.
Unlike tunnel-based connectivity solutions, the cMCP Runtime processes tool-call payloads inside the TEE. The connectivity provider sees ciphertext, not plaintext. The only thing that leaves the enclave is the signed TRACE claim.
pip install cmcp-runtime
Create cmcp-config.yaml :
attestation :
provider : auto
enforcement_mode : advisory # advisory eases first-run tuning; the default is `enforcing`
policy_bundle_path : ./policies/
catalog_path : ./catalog.json
Start the gateway:
CMCP_DEV_MODE=1 cmcp start --config cmcp-config.yaml
Make a tool call:
curl -X POST http://localhost:8443/mcp \
-H " Content-Type: application/json " \
-d ' {"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"salesforce.contacts","arguments":{"query":"Acme Corp"},"_cmcp":{"session_id":"s1","workflow_id":"demo-agent"}}} '
See docs/quickstart.md for the full walkthrough: Cedar policy, tool catalog, first TRACE Claim, and verification (no hardware TEE required).
The agent sends every tool call to the cMCP Gateway instead of directly to MCP servers.
At startup the gateway measures the Cedar policy bundle hash into the hardware attestation report. No code runs before this measurement.
Each incoming tool call is evaluated by the Cedar policy engine running inside the TEE. The result is allow, deny, or redact. The call and its decision are appended to the hardware-sealed audit chain.
At the end of the session the gateway produces a TRACE Claim: a signed, hardware-attested artifact that records which tools ran, which policy decided each call, and the full audit chain. A verifier checks this without trusting the operator.
Agent -> cMCP Runtime -> Cedar Policy Engine (TEE) -> Tool
|
GatewayClaim (TRACE Profile)
+-- trace.eat_profile
+-- trace.runtime.platform + measurement
+-- trace.policy.bundle_hash
+-- trace.cnf.jwk (Ed25519 confirmation key)
+-- gateway.audit_chain (root/tip/length)
+-- signature (Ed25519 over canonical JSON)
Hardware providers
Provider
Platform
Assurance
Notes
tpm
TPM 2.0 / vTPM (Azure, AWS, GCP Trusted Launch)
Medium
Local TPM quote
sev-snp
AMD SEV-SNP (Azure DCasv5, AWS C6a Nitro)
High
AMD KDS
tdx
Intel TDX (Azure DCedsv5, GCP C3)
High
Intel PCS
gpu-cc (v0.2)
NVIDIA H100/H200/Blackwell (CC mode)
High
NVIDIA Remote Attestation Service (NRAS)
opaque (opt-in)
OPAQUE Confidential Runtime
n/a (not yet implemented)
Placeholder: excluded from auto-detect; selecting it explicitly raises a not-implemented error
Provider auto-detect probe order: azure-cvm -> tpm -> sev-snp -> tdx . The first provider whose detect() succeeds is selected. opaque is a not-yet-implemented placeholder: it is excluded from auto-detect, and selecting it explicitly raises ATTESTATION_PROVIDER_NOT_IMPLEMENTED rather than falling through silently. If no hardware provider is detected, the gateway starts only under CMCP_DEV_MODE=1 (a non-attested software-only fallback) and otherwise refuses to start.
from cmcp_runtime . config import TEEProvider
# Auto-detect (default)
# attestation.provider: auto -> azure-cvm -> tpm -> sev-snp -> tdx
# (software-only is used only under CMCP_DEV_MODE=1)
# Explicit hardware selection
# attestation.provider: sev-snp
# OPAQUE Managed Runtime (opt-in only; not yet implemented)
# OPAQUE_ATTESTATION_URL=https://... cmcp start --config cmcp-config.yaml
Enforcement modes
Mode
Behavior
Use case
enforcing
Policy denies return HTTP 403; call is not forwarded
Production
advisory
Policy denies are logged; call proceeds
First deployment, policy tuning
silent
Policy is evaluated but nothing is logged or blocked
Baselining
Default is enforcing . Set enforcement_mode: advisory in cmcp-config.yaml to use advisory mode.
cmcp-config.yaml full reference:
attestation :
provider : auto # auto | tpm | sev-snp | tdx | opaque | software-only
enforcement_mode : enforcing # enforcing | advisory | silent
validity_seconds : 86400 # attestation freshness window (default: 24 hours)
staleness_policy : fail_closed # fail_closed | warn_only
expected_measurement : ~ # pin a specific PCR/measurement (optional)
policy_bundle_path : policies/ # directory containing .cedar files and manifest.json
catalog_path : catalog.json # approved tool catalog
listen_addr : " 0.0.0.0:8443 "
max_response_size_bytes : 2097152 # 2 MB default
policy_reload_interval_seconds : 0 # 0 = disabled; restart required to update policy
Environment variables:
Command
Flags
Description
cmcp start
--config PATH (required)
Start the gateway
cmcp validate-config
--config PATH (required)
Validate cmcp-config.yaml without starting
cmcp validate-bundle
--bundle-path PATH (required), --expected-hash sha256:<hex> (required)
Verify a Cedar bundle hash before deployment
cmcp verify
CLAIM_FILE (required); --policy-hash , --catalog-hash , --max-age , --trusted-key , --audit-bundle , --agent-manifest , --agent-manifest-trust-anchor
Verify a signed TRACE Claim (signature, schema, freshness, audit chain, and pinned hashes)
TRACE Claims
A GatewayClaim is the unit of proof handed to an auditor, regulator, or downstream verifier. It is produced per session (or per call, configurable) and signed with a key that never leaves the TEE.
(This table is a summary of the most-used fields.)
Verification with the cmcp_verify library does not require trusting the operator. The verifier checks the signature against the TEE-bound key, the policy bundle hash against the approved value, and the audit chain for internal consistency.
The normative schema is schemas/trace-claim.schema.json , and docs/quickstart.md shows a complete example. See docs/spec/verification-library.md and the TRACE specification for the full verification protocol.
Standard
Coverage
OWASP Agentic AI Top 10
MCP10 (data leakage via tool calls), MCP02 (unsanctioned tools), MCP08 (provable governance), MCP04 (supply chain)
NIST SP 800-207
Policy decision point inside TEE; no implicit trust in workload identity
EU AI Act Art. 12, 15
Per-decision audit records (Art. 12); TEE-backed cybersecurity controls (Art. 15)
DORA Art. 9
Attestation chain; audit log retention via gateway.audit_chain
RATS/EAT RFC 9711
GatewayClaim is an EAT; eat_profile field identifies the TRACE profile
Security
Tool
What it checks
ruff
Style and import linting on every PR
bandit
Python security linting on every PR
pip-audit
Dependency vulnerability scan on every PR
mypy
Static type checking on every PR
CodeQL
Python SAST, security-extended queries, weekly
OpenSSF Scorecard
Weekly scoring, SARIF upload
See SECURITY.md for vulnerability reporting and response SLAs. See LIMITATIONS.md for explicit scope boundaries, including residual risks for APM payload capture, runtime config injection, and P4.1 supply chain (typosquat) that Phase 1 does not close.
Page
Description
docs/quickstart.md
From zero to first TRACE Claim in under 30 minutes
docs/configuration.md
Full config reference with all fields and defaults
docs/SPEC.md
Product specification: problem taxonomy, architecture, coverage matrix
docs/spec/threat-model.md
STRIDE analysis, adversary model, residual risks
docs/spec/cedar-policy.md
Ced

[truncated]
