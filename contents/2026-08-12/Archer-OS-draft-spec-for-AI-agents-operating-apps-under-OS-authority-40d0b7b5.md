---
source: "https://github.com/coachpato/archer-os"
hn_url: "https://news.ycombinator.com/item?id=49274512"
title: "Archer OS: draft spec for AI agents operating apps under OS authority"
article_title: "GitHub - coachpato/archer-os: Draft architecture and specification for a model-independent operating environment for AI agents. · GitHub"
author: "Billie_Archer"
captured_at: "2026-08-12T16:47:16Z"
capture_tool: "hn-digest"
hn_id: 49274512
score: 2
comments: 0
posted_at: "2026-08-12T15:59:25Z"
tags:
  - hacker-news
  - translated
---

# Archer OS: draft spec for AI agents operating apps under OS authority

- HN: [49274512](https://news.ycombinator.com/item?id=49274512)
- Source: [github.com](https://github.com/coachpato/archer-os)
- Score: 2
- Comments: 0
- Posted: 2026-08-12T15:59:25Z

## Translation

タイトル: Archer OS: OS 権限でアプリを操作する AI エージェントの仕様草案
記事のタイトル: GitHub - Coachpato/archer-os: AI エージェント用のモデルに依存しないオペレーティング環境のドラフト アーキテクチャと仕様。 · GitHub
説明: AI エージェントのモデルに依存しない動作環境のアーキテクチャと仕様の草案。 - コーチパト/アーチャー-OS

記事本文:
GitHub - Coachpato/archer-os: AI エージェント用のモデルに依存しないオペレーティング環境のドラフト アーキテクチャと仕様。 · GitHub
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
コーチパト
/
アーチャーOS
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .github .github docs docs 例 例 原則 原則 rfcs rfcs 仕様 仕様 .gitattributes .gitattributes .markdownlint.json .markdownli

nt.json CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md FAQ.md FAQ.md GOVERNANCE.md GOVERNANCE.md ライセンス ライセンス LICENSE-DOCS.md LICENSE-DOCS.md PAPER.md PAPER.md PRIOR_ART.md PRIOR_ART.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SPECIFICATION.md SPECIFICATION.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Archer OS は、モデルに依存しない AI エージェントが OS 管理のポリシーに基づいてインストールされたソフトウェアを検出、権限要求、意味論的に操作するためのオペレーティング環境とアプリケーション コントラクトの草案です。
ステータス: ドラフト 0.1 アーキテクチャおよび仕様。このリポジトリは、完成したオペレーティング システム、デスクトップ アシスタント、エージェント フレームワーク、または Linux ディストリビューションではありません。
AI エージェントはすでにツールの呼び出し、GUI の駆動、API の呼び出し、ワークフローの自動化を行うことができますが、これらのメカニズムは通常、アプリケーション固有の統合、フレームワーク固有のツール定義、エージェントにコピーされた資格情報、および脆弱な画面操作にわたって断片化されています。 Archer OS は異なる境界を探ります。モデルやアプリケーションではなく、オペレーティング環境が、機能の検出、委任された権限、認証情報、実行、イベント、アプリケーション間の構成、および監査の受信を管理する必要があります。
目的は、提案を批判から守ることではありません。目標は、失敗してもいいほど具体的なものにすることです。
決定的なベンチマークは、Archer Unseen ペア テストです。
これまでになかった 1 つの Archer 互換エージェントと、これまでになかった 1 つの Archer ネイティブ アプリケーションは、それらの間に特別なペアごとの統合を行わずに、次のことができるはずです。
OS を介した機能検出を通じて相互に検出します。
型付き機能スキーマを理解する。
委任された権限の範囲内で許可を要求および受け取ります。
構造化された結果、エラー、キャンセル、およびイブを処理します

nts;
別のアプリケーションで結果を作成します。
拒否と取り消しを尊重します。そして
結果として生じるアクションに対して監査可能な領収書を作成します。
フローチャート TD
A["推論/エージェント層"] --> B["推論抽象化層"]
B --> C[「能力ブローカー」]
C --> D[「ガードレール / ポリシーの施行」]
D --> E["ユニバーサル アプリ プロトコル"]
E --> F["Archer ネイティブ アプリケーション"]
E --> G["レガシー アダプター/コネクタ"]
F --> H["既存の OS / Linux サービス"]
G --> H
H --> I["カーネル / ハードウェア"]
D --> J["資格情報および秘密ブローカー"]
D --> K["監査ログと受領書"]
読み込み中
セマンティック ガバナンスは、カーネル リソース管理とは異なります。 Archer OS は当初、代替カーネルではなく、クロスプラットフォーム セマンティック ランタイム、アプリケーション コントラクト、および Linux リファレンス デスクトップ環境をターゲットとしています。
ヒューマン エージェント機能パリティ (HAFP): Archer ネイティブ アプリケーションのインターフェイスを通じて人間が利用できるマテリアル ドメイン機能には、技術的に適切な場合に限り、セマンティック エージェントがアクセスできる同等の機能が必要であり、両方のインターフェイスは最終的に同じ基盤となるアプリケーションまたはドメイン操作を呼び出す必要があります。
Archer OS は、一般的なアイデアとして、アプリの自動化、セマンティック API、機能検出、OS IPC、エージェント ツール、または人間とエージェントの同等性を発明したとは主張していません。これは、HAFP を OS レベルのアプリケーション適合原則に昇格することを提案しています。
能力は権限ではありません: 能力を公開しても、エージェントがそれを呼び出せるわけではありません。 Archer OS は、利用可能な機能、ユーザー ID、エージェント ID、委任された権限、コンテキスト、リスク、範囲、および実際の実行権限を分離します。
モデルの互換性: モデルとエージェント フレームワークは、置き換え可能なインテリジェンス プロバイダーです。 Archer OS は、単一のモデル プロバイダーを見つけることなく、ローカル モデル、クラウド モデル、または将来のシステムを使用してエージェントを仲介できる必要があります。

日付的な。
PAPER.md : ポジションペーパーとアーキテクチャの議論。
SPECIFICATION.md : 規範仕様の草案。
原則/ : 個別の設計原則。
spec/ : 技術モジュールのドラフト。
例/ : アプリケーション契約と権限の例を示します。
docs/threat-model.md : セキュリティ脅威モデル。
docs/review-plan.md : パブリック レビューとドラフト 0.2 の優先順位。
PRIOR_ART.md : 先行技術のレビューと位置付け。
次の GitHub の問題を開いてください。
従来技術: これがすでに存在する場所を示します。
アーキテクチャの反対: 失敗例を使って仮定に異議を唱えます。
セキュリティ上の懸念: 対処可能な脆弱性については責任ある開示を行います。
仕様提案: 具体的な契約変更を提案します。
.github/ISSUE_TEMPLATE の問題テンプレートでは、批判を有益にするために必要な証拠を求めます。
機密のセキュリティ レポートについては、利用可能な場合は GitHub のプライベート脆弱性レポートを使用するか、 Hello@bid360.co.za にお問い合わせください。
プロジェクト連絡先: Hello@bid360.co.za
コードと実行可能ファイルのサンプルでは、Apache License 2.0 を使用します。ドキュメント、仕様テキスト、図、散文では CC BY 4.0 を使用します。 PAPER.md を参照してください。
AI エージェントのモデルに依存しない動作環境のアーキテクチャと仕様の草案。
Readme Apache-2.0、ライセンス ライセンスが見つかりました 行動規範
セキュリティ ポリシー このリポジトリを引用する アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Draft architecture and specification for a model-independent operating environment for AI agents. - coachpato/archer-os

GitHub - coachpato/archer-os: Draft architecture and specification for a model-independent operating environment for AI agents. · GitHub
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
coachpato
/
archer-os
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .github .github docs docs examples examples principles principles rfcs rfcs spec spec .gitattributes .gitattributes .markdownlint.json .markdownlint.json CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md FAQ.md FAQ.md GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE LICENSE-DOCS.md LICENSE-DOCS.md PAPER.md PAPER.md PRIOR_ART.md PRIOR_ART.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SPECIFICATION.md SPECIFICATION.md View all files Repository files navigation
Archer OS is a draft operating environment and application contract for model-independent AI agents to discover, request authority for, and semantically operate installed software under OS-governed policy.
Status: Draft 0.1 architecture and specification. This repository is not a completed operating system, desktop assistant, agent framework, or Linux distribution.
AI agents can already call tools, drive GUIs, invoke APIs, and automate workflows, but those mechanisms are usually fragmented across application-specific integrations, framework-specific tool definitions, credentials copied into agents, and brittle screen interaction. Archer OS explores a different boundary: the operating environment, not the model or application, should govern capability discovery, delegated authority, credentials, execution, events, inter-application composition, and audit receipts.
The goal is not to protect the proposal from criticism. The goal is to make it concrete enough to fail.
The defining benchmark is the Archer Unseen Pair Test:
One previously unseen Archer-compatible agent and one previously unseen Archer-native application, with no bespoke pairwise integration between them, should be able to:
discover one another through OS-mediated capability discovery;
understand typed capability schemas;
request and receive permission within a delegated mandate;
handle structured results, errors, cancellation, and events;
compose the result with another application;
respect denial and revocation; and
produce an auditable receipt for consequential actions.
flowchart TD
A["Reasoning / Agent Layer"] --> B["Inference Abstraction Layer"]
B --> C["Capability Broker"]
C --> D["Guardrail / Policy Enforcement"]
D --> E["Universal App Protocol"]
E --> F["Archer-Native Applications"]
E --> G["Legacy Adapters / Connectors"]
F --> H["Existing OS / Linux Services"]
G --> H
H --> I["Kernel / Hardware"]
D --> J["Credential & Secret Broker"]
D --> K["Audit Log & Receipts"]
Loading
Semantic governance is different from kernel resource management. Archer OS initially targets a cross-platform semantic runtime, an application contract, and a Linux reference desktop environment rather than a replacement kernel.
Human-Agent Functional Parity (HAFP): material domain functionality available to a human through an Archer-native application's interface should have a semantic agent-accessible equivalent wherever technically appropriate, and both interfaces should ultimately invoke the same underlying application or domain operation.
Archer OS does not claim to invent app automation, semantic APIs, capability discovery, OS IPC, agent tools, or human-agent parity as a general idea. It proposes elevating HAFP into an OS-level application conformance principle.
Ability is not Authority: exposing a capability does not mean an agent may invoke it. Archer OS separates available capability, user identity, agent identity, delegated mandate, context, risk, scope, and actual execution authority.
Model interchangeability: models and agent frameworks are replaceable intelligence providers. Archer OS should be able to broker agents using local models, cloud models, or future systems without making any one model provider foundational.
PAPER.md : position paper and architecture argument.
SPECIFICATION.md : draft normative specification.
principles/ : individual design principles.
spec/ : draft technical modules.
examples/ : illustrative application contracts and authority examples.
docs/threat-model.md : security threat model.
docs/review-plan.md : public review and Draft 0.2 priorities.
PRIOR_ART.md : prior-art review and positioning.
Please open GitHub issues for:
prior art: show where this already exists;
architecture objections: challenge assumptions with failure cases;
security concerns: use responsible disclosure for actionable vulnerabilities;
specification proposals: suggest concrete contract changes.
The issue templates in .github/ISSUE_TEMPLATE ask for the evidence needed to make criticism useful.
For sensitive security reports, use GitHub private vulnerability reporting if available, or contact Hello@bid360.co.za .
Project contact: Hello@bid360.co.za
Code and executable examples use the Apache License 2.0 . Documentation, specification text, diagrams, and prose use CC BY 4.0 ; see PAPER.md .
Draft architecture and specification for a model-independent operating environment for AI agents.
Readme Apache-2.0, License licenses found Code of conduct
Security policy Cite this repository Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
