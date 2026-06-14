---
source: "https://github.com/ComplyEdge/complyedge"
hn_url: "https://news.ycombinator.com/item?id=48528667"
title: "Show HN: ComplyEdge – Runtime EU AI Act Enforcement for Python"
article_title: "GitHub - ComplyEdge/complyedge: Open-source compliance engine for AI agents. Rules, SDKs, and examples. · GitHub"
author: "lc-complyedge"
captured_at: "2026-06-14T18:37:19Z"
capture_tool: "hn-digest"
hn_id: 48528667
score: 2
comments: 0
posted_at: "2026-06-14T15:57:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ComplyEdge – Runtime EU AI Act Enforcement for Python

- HN: [48528667](https://news.ycombinator.com/item?id=48528667)
- Source: [github.com](https://github.com/ComplyEdge/complyedge)
- Score: 2
- Comments: 0
- Posted: 2026-06-14T15:57:06Z

## Translation

タイトル: Show HN: ComplyEdge – Python 向けランタイム EU AI 法の施行
記事のタイトル: GitHub - ComplyEdge/complyedge: AI エージェント用のオープンソース コンプライアンス エンジン。ルール、SDK、および例。 · GitHub
説明: AI エージェント用のオープンソース コンプライアンス エンジン。ルール、SDK、および例。 - コンプライエッジ/コンプライエッジ

記事本文:
GitHub - ComplyEdge/complyedge: AI エージェント用のオープンソース コンプライアンス エンジン。ルール、SDK、および例。 · GitHub
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
コンプライエッジ
/
コンプライエッジ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

s
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
26 コミット 26 コミット .github/ workflows .github/ workflows docs docs サンプル サンプル パッケージ/ trustlint パッケージ/ trustlint プロバイダー プロバイダー ルール ルール スクリプト/ ベンチマーク スクリプト/ ベンチマーク sdks/ python sdks/ python テスト テスト CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md LAST_EXPORT LAST_EXPORT LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントに対するランタイム コンプライアンスの強制。スキャナーではありません。すべてのリクエストに応じて本番環境で実行されます。
第5条はすでに法律になっている。 GPAI の罰金は 2026 年 8 月 2 日から始まります。あなたの AI は現時点で準拠しているか、していないかのどちらかです。
コンプライアンス ツールがリクエストをブロックする場合、規制当局に何を伝えますか?確率スコア？
ComplyEdge は次のように述べています: 第 5 条(1)(a)、規則 rego-art5-1a-001 、タイムスタンプ、入力ハッシュ。 1 つは監査証跡です。一つは推測です。
pip インストール準拠エッジ
コンプライエッジインポートから、compliance_check
@compliance_check (管轄区域 = "EU" 、agent_id = "my-agent")
def my_agent (プロンプト):
llm を返します。生成 (プロンプト) # すべての入力と出力をチェック
3行。すべての AI の入力と出力は、EU AI 法のルール コーパス (第 5 条、第 50 条、GPAI) に対して評価されます。違反はユーザーに届く前にブロックされ、記事の引用、ルール ID、すべての決定のタイムスタンプが表示されます。
COMPLYEDGE_API_KEY をキーに設定します。デコレーターはデフォルトでアクティブになります。キーを削除せずに無効にするには (CI などで)、 COMPLYEDGE_ENABLED=false を設定します。
applyedge import is_safe から、チェックしてください
OSをインポートする
api_key = os 。環境 [ "COMPLYEDGE_API_KEY" ]
# ブール値チェック — 違反がない場合は True を返します
is_safe でない場合 (プロンプト、 api_key = api_key 、管轄 = "EU" ):
raise ValueError (「プロンプトは EU AI 法に違反しています」)

# 完全な結果 — 違反と引用を含む ComplianceResult を返します
result = check (プロンプト、api_key = api_key、管轄 = "EU")
そうでない場合は結果を返します。許可される:
v の結果 。違反:
print ( v . rules_id , v . citation )
管轄区域はルール コーパスにマップされます。EU は、EU AI 法第 5 条、第 50 条、および GPAI 義務に照らして評価します。米国は、HIPAA、SOX、COPPA、TCPA、BIPA に対して評価します。
API キーは必要ありません。正規表現パターンを使用して、YAML ルール コーパスに対してテキストをスキャンします。スタンドアロン パッケージとして公開され、SDK とは独立してバージョン管理されます。
pip インストール trustlint
trustlint check --text " 応募者の評価には社会的信用スコアを使用します "
# → クリティカル: EU_AI_ACT_ART5_SOCIAL_SCORING_001 — 第 5 条(1)(c)
終了コード: 0 = 合格、1 = 違反が見つかりました。 CI/CD パイプライン用に設計されています。ソース: パッケージ/trustlint/ 。
ComplyEdge は、それぞれ独自のルール ID 名前空間を持つ 2 つのエンジンを通じて同じ規制を解決します。
ランタイム API (OPA/Rego): rego-art5-1c-001 のような ID — コンプライアンス_チェックと /v1/check API によって返されます。これは、運用システムのログに記録される監査証跡です。
TrustLint (オフライン、YAML コーパス): EU_AI_ACT_ART5_SOCIAL_SCORING_001 のような ID — オフライン リンターによって発行されます。
どちらも同じ法律条項を引用しており、エンジンのみが異なります。すべてのルールに含まれる記事参照を介してそれらの間をマッピングします。
sdks/python/ Python SDK (@compliance_check デコレータ、CLI)
package/trustlint/ オフライン正規表現リンター (TrustLint) — API キーなし、CI/CD 用
ルール/規制/ 53 YAML ルール (EU AI 法、GDPR、HIPAA、SOX、PCI DSS など)
rules/rego/ 19 OPA/Rego ポリシー (EU AI 法第 5 条、第 50 条、GPAI)
ルール/スキーマ/ ルール検証スキーマ
例/使用例 (デコレーター、OpenAI エージェント)
scripts/benchmark/ ランタイム ベンチマーク (ランナー + プロンプト YAML + コミットされた結果)
テスト/ルール検証

ション + 受け入れテスト
ルール
4 つの管轄区域にわたる 53 の YAML ルール + 19 の OPA/Rego ポリシー:
各ルールでは、条件、重大度、検出範囲、法的引用を伴う修復を指定します。形式についてはルール スキーマを参照してください。
id : MY_CUSTOM_RULE_001
管轄区域 : EU
有効日: " 2025-02-02 "
説明: 「第 Y 条に基づく禁止行為 X を検出する」
重大度 : クリティカル
条件：
- タイプ: 正規表現
値：「禁止パターン」
説明: 「試合禁止練習X」
ソース:
規制：「EU AI法」
記事：「記事Y(1)(z)」
検証: cd ルール && python scripts/validate_rules.py
レイヤ 1 — 決定的 (ホット パス): 19 の OPA/Rego ポリシーがすべてのリクエストを評価します。ブロックされたプロンプトは、数十ミリ秒以内に法的引用とともに返されます。ベンチマークでは、OPA でブロックされたプロンプト全体で 38 ～ 100 ミリ秒 (中央値 62 ミリ秒) でした。バイナリ パス/ブロック、ホット パス上に LLM なし。 (TrustLint は、CI で使用するために同じ正規表現コーパスをオフラインで適用します。)
レイヤ 2 — インタプリティブ (同期、オプトイン): use_semantic_fallback=True を指定して呼び出された場合、LLM はリクエストを評価し、違反が見つかった場合はブロックします。 v0.2.2 以降、デフォルトでオフになっています。リクエストごとに 2 ～ 5 秒のレイテンシが追加されます。
セキュリティ製品は AI を悪意のある者から保護します。 ComplyEdge は、通常の業務中に自社の AI による法違反から企業を保護します。
50 個のプロンプト コーパスがライブ API に対して実行されます。ランナー、プロンプト YAML、および最新の結果 JSON は scripts/benchmark/ でコミットされます。結果を直接検査するか、独自の COMPLYEDGE_API_KEY を使用して再実行します。
ルールへの貢献を歓迎します。詳細については、CONTRIBUTING.md を参照してください。
すべてのルールには、記事と段落の引用、検証可能な検出条件、テスト ケースが含まれている必要があります。
脆弱性を報告するには、 SECURITY.md を参照してください。セキュリティ レポートの公開問題を開かないでください。
Apache ライセンス 2.0

— ライセンスを参照してください。
PyPI : pypi.org/project/complyedge
ルールスキーマ: rules/schemas/rule-schema.json
AI エージェント用のオープンソース コンプライアンス エンジン。ルール、SDK、および例。
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
1
v0.2.4 — ランタイム EU AI 法の施行
最新の
2026 年 6 月 14 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open-source compliance engine for AI agents. Rules, SDKs, and examples. - ComplyEdge/complyedge

GitHub - ComplyEdge/complyedge: Open-source compliance engine for AI agents. Rules, SDKs, and examples. · GitHub
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
ComplyEdge
/
complyedge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
26 Commits 26 Commits .github/ workflows .github/ workflows docs docs examples examples packages/ trustlint packages/ trustlint providers providers rules rules scripts/ benchmark scripts/ benchmark sdks/ python sdks/ python tests tests CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LAST_EXPORT LAST_EXPORT LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Runtime compliance enforcement for AI agents. Not a scanner — runs in production, on every request.
Article 5 is already law. GPAI fines start August 2, 2026. Your AI is either compliant right now, or it isn't.
What does your compliance tool tell a regulator when it blocks a request? A probability score?
ComplyEdge says: Article 5(1)(a), rule rego-art5-1a-001 , timestamp, input hash. One is an audit trail. One is a guess.
pip install complyedge
from complyedge import compliance_check
@ compliance_check ( jurisdiction = "EU" , agent_id = "my-agent" )
def my_agent ( prompt ):
return llm . generate ( prompt ) # every input and output checked
Three lines. Every AI input and output evaluated against the EU AI Act rule corpus (Article 5, Article 50, GPAI). Violations blocked before they reach the user — with article citation, rule ID, and timestamp on every decision.
Set COMPLYEDGE_API_KEY to your key. The decorator activates by default; to disable without removing the key (e.g., in CI), set COMPLYEDGE_ENABLED=false .
from complyedge import is_safe , check
import os
api_key = os . environ [ "COMPLYEDGE_API_KEY" ]
# Boolean check — returns True if no violations
if not is_safe ( prompt , api_key = api_key , jurisdiction = "EU" ):
raise ValueError ( "Prompt violates EU AI Act" )
# Full result — returns ComplianceResult with violations + citations
result = check ( prompt , api_key = api_key , jurisdiction = "EU" )
if not result . allowed :
for v in result . violations :
print ( v . rule_id , v . citation )
Jurisdiction maps to the rule corpus: EU evaluates against EU AI Act Article 5, Article 50, and GPAI obligations. US evaluates against HIPAA, SOX, COPPA, TCPA, BIPA.
No API key required. Scans text against the YAML rule corpus using regex patterns. Published as a standalone package, versioned independently of the SDK.
pip install trustlint
trustlint check --text " We use social credit scoring to evaluate applicants "
# → CRITICAL: EU_AI_ACT_ART5_SOCIAL_SCORING_001 — Article 5(1)(c)
Exit codes: 0 = pass, 1 = violations found. Designed for CI/CD pipelines. Source: packages/trustlint/ .
ComplyEdge resolves the same regulations through two engines, each with its own rule-ID namespace:
Runtime API (OPA/Rego): IDs like rego-art5-1c-001 — returned by compliance_check and the /v1/check API. This is the audit trail your production system logs.
TrustLint (offline, YAML corpus): IDs like EU_AI_ACT_ART5_SOCIAL_SCORING_001 — emitted by the offline linter.
Both cite the same legal article and differ only in engine. Map between them via the article reference carried in every rule.
sdks/python/ Python SDK (@compliance_check decorator, CLI)
packages/trustlint/ Offline regex linter (TrustLint) — no API key, for CI/CD
rules/regulations/ 53 YAML rules (EU AI Act, GDPR, HIPAA, SOX, PCI DSS, and more)
rules/rego/ 19 OPA/Rego policies (EU AI Act Article 5, 50, GPAI)
rules/schemas/ Rule validation schema
examples/ Usage examples (decorators, OpenAI Agents)
scripts/benchmark/ Runtime benchmark (runner + prompt YAMLs + committed results)
tests/ Rule validation + acceptance tests
Rules
53 YAML rules + 19 OPA/Rego policies across 4 jurisdictions:
Each rule specifies conditions, severity, detection scope, and remediation with legal citations. See the rule schema for the format.
id : MY_CUSTOM_RULE_001
jurisdiction : EU
effective_date : " 2025-02-02 "
description : " Detect prohibited practice X under Article Y "
severity : critical
conditions :
- type : regex
value : " prohibited pattern "
description : " Matches prohibited practice X "
source :
regulation : " EU AI Act "
article : " Article Y(1)(z) "
Validate: cd rules && python scripts/validate_rules.py
Layer 1 — Deterministic (hot path): 19 OPA/Rego policies evaluate every request. Blocked prompts return with a legal citation in tens of milliseconds — 38–100ms (median 62ms) across the OPA-blocked prompts in our benchmark. Binary pass/block, no LLM on the hot path. (TrustLint applies the same regex corpus offline for CI use.)
Layer 2 — Interpretive (synchronous, opt-in): When called with use_semantic_fallback=True , an LLM evaluates the request and blocks if a violation is found. Off by default since v0.2.2. Adds 2–5s latency per request.
Security products protect AI from bad actors. ComplyEdge protects companies from their own AI's legal violations during normal operations.
A 50-prompt corpus runs against the live API. The runner, prompt YAMLs, and the latest result JSON are committed under scripts/benchmark/ — inspect the results directly, or re-run with your own COMPLYEDGE_API_KEY .
We welcome rule contributions. See CONTRIBUTING.md for details.
Every rule must include: article + paragraph citation, verifiable detection condition, and test cases.
To report a vulnerability, see SECURITY.md . Do not open a public issue for security reports.
Apache License 2.0 — see LICENSE .
PyPI : pypi.org/project/complyedge
Rule Schema : rules/schemas/rule-schema.json
Open-source compliance engine for AI agents. Rules, SDKs, and examples.
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
v0.2.4 — runtime EU AI Act enforcement
Latest
Jun 14, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
