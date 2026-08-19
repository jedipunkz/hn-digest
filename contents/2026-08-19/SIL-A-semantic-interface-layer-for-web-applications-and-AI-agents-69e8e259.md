---
source: "https://github.com/ais-space/sil"
hn_url: "https://news.ycombinator.com/item?id=49359448"
title: "SIL: A semantic interface layer for web applications and AI agents"
article_title: "GitHub - ais-space/sil: Open specification for a machine-readable web. Self-hosting (written in its own format), 3,305 lines, 5 conformance profiles, 40 test vectors. CC-BY-4.0. · GitHub"
image: "https://opengraph.githubassets.com/bc79339aa8ff154d3b5da5b6ebd2f128146fcdf91e9bb3512ac5d92a7dcfee59/ais-space/sil"
author: "vladimir_si"
captured_at: "2026-08-19T10:19:27Z"
capture_tool: "hn-digest"
hn_id: 49359448
score: 1
comments: 0
posted_at: "2026-08-19T10:17:25Z"
tags:
  - hacker-news
  - translated
---

# SIL: A semantic interface layer for web applications and AI agents

- HN: [49359448](https://news.ycombinator.com/item?id=49359448)
- Source: [github.com](https://github.com/ais-space/sil)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T10:17:25Z

## Translation

タイトル: SIL: Web アプリケーションと AI エージェントのためのセマンティック インターフェイス レイヤー
記事のタイトル: GitHub - ais-space/sil: 機械可読 Web のオープン仕様。セルフホスティング (独自の形式で記述)、3,305 行、5 つの適合プロファイル、40 のテスト ベクター。 CC-BY-4.0。 · GitHub
説明: 機械可読 Web のオープン仕様。セルフホスティング (独自の形式で記述)、3,305 行、5 つの適合プロファイル、40 のテスト ベクター。 CC-BY-4.0。 - ais-space/sil

記事本文:
GitHub - ais-space/sil: 機械可読 Web のオープン仕様。セルフホスティング (独自の形式で記述)、3,305 行、5 つの適合プロファイル、40 のテスト ベクター。 CC-BY-4.0。 · GitHub
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
AIスペース
/
シル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
7 コミット 7 コミット フォルダーとファイル
例 例 文法 文法スキーマ スキーマ CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md COD

E_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス ライセンスコード ライセンスコード README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SIL_Concept.md SIL_Concept.md SIL_Open_仕様.sil SIL_Open_仕様.silすべてのファイルを表示 リポジトリ ファイルのナビゲーション
SIL — セマンティック インターフェイス層
インテリジェント エージェント用のテキストベースのインターフェイス プロトコル。
SIL (Semantic Interface Layer) は、AI エージェントに Web アプリケーションへの直接的なセマンティック インターフェイスを提供するためのオープン仕様です。
従来の Web アプリケーションは主に人間との対話を目的として設計されています。そのインターフェイスは、視覚的なレイアウト、ナビゲーション、コントロール、クライアント側の動作、その他のプレゼンテーション メカニズムを通じて情報を公開します。エージェントはそのインターフェイスを使用できますが、多くの場合、それらのプレゼンテーション成果物からアプリケーションの構造と意味を再構築する必要があります。
SIL は、同じアプリケーションに別のインターフェイスを追加します。これは、アプリケーションの状態、使用可能なオブジェクト、関係、ナビゲーション、機能、およびアクションの構造化されたテキスト表現です。
ヒューマンインターフェースは残ります。 SIL はこれに代わるものではありません。
目標は、エージェントにインターフェースのリバース エンジニアリングを要求するのではなく、明確に定義されたインターフェースと同等のものをエージェントに提供することです。
1 つのアプリケーション、2 つのインターフェース
2 つのインターフェイスは、同じバックエンド、ビジネス ロジック、データ、認証、認可、およびアプリケーションの状態を共有できます。
┌─────────────┐
│ アプリケーション │
│ │
│ ビジネスロジック │
│ データと状態 │
│ 権限 │
━─────┬─────┘
│
┌─────────┴─────────┐
│ │
▼ ▼
ヒューマン インターフェイス セマンティック インターフェイス
HTML SIL / STF
│ │
▼ ▼
人間エージェント
これは私

必ずしも情報が重複しているわけではありません。これは、消費者の能力と要件に応じてプレゼンテーションを分離することです。
エージェントはすでに Web ページ、API、MCP ツール、その他のインターフェイスと対話できるようになります。 SIL は別の問題に対処します。
エージェントは、プレゼンテーション指向のアーティファクトから情報を再構築せずに、アプリケーション自体 (その構造、現在の状態、利用可能なアクション、ナビゲーション、境界) をどのようにして理解できるでしょうか?
たとえば、HTML ページには、周囲の要素、スタイル、JavaScript ハンドラー、アプリケーションの状態、および視覚的なコンテキストによって意味が決定されるボタンが含まれる場合があります。
SIL 表現では、その意味を直接表現できます。
ダウンロードボタン
タイプ: ボタン
役割: プライマリアクション
ID: btn_download
キャプション: ダウンロード
アクション: アクティブ化
違いは、一方には情報が含まれ、もう一方には情報が含まれていないということではありません。
違いは、どこで解釈が行われるかです。
従来のインターフェイスでは、エージェントはプレゼンテーションからセマンティクスを推測する必要がある場合があります。 SIL を使用すると、アプリケーションはこれらのセマンティクスを明示的に公開できます。
SIL は、インテリジェント エージェントのアプリケーション語彙とインターフェイス プロトコルを定義します。
SIL v1.1.0 以降、基礎となる構文とセマンティック コアはインラインで定義されなくなりました。 SIL は現在、STF コアの拡張機能として仕様化されています。これは、 github.com/ais-space/stf で独立して公開されている別個のスタンドアロン仕様 (セマンティック テキスト形式) です。
STFコア（スタンドアロン仕様）
§── 構文 (§2)
§── セマンティックモデル (§3)
§── 構造上の安全性（§4）
└── データ/シリアル化モデル (§5、§6)
SIL v1.1.0 (本仕様)
§── SIL Vocabulary — STF コアのアプリケーション拡張 (§3)
§── 応用語彙（§4）
§── SIL プロトコル — 発見と対話 (§5)

)
§── インタラクションパターン (§6)
§── セキュリティ固有のガイダンス (§7)
§── 適合プロファイル（§8）
└── 拡張と今後の方向性（§9）
STF コアは、バージョン固定なしで、標準的に参照によって参照されます。つまり、どの SIL ドキュメントも構造上、有効な STF ドキュメントです。 STF コアが進化すると、SIL はアプリケーション層でそれを追跡します。 STF コアは意図的にテキスト形式で階層化されており、言語モデルが解釈しやすいと同時に人間が読みやすいように設計されています。
SIL 仕様自体は STF で書かれており、単に形式に関する文書ではなく、自己記述型の適合文書となっています。
SIL 対応アプリケーションは、HTTP を通じてセマンティック ページを公開します。
典型的な対話は次のようになります。
エージェントSILサーバー
│ │
│ /revizor.sil を取得 │
│─────────────>│
│ │
│ STF ドキュメント │
│<───────────────│
│ │
│ POST /revizor.sil │
│ {"インテント":"アクティブ化", │
│ "ターゲット":"btn_download"} │
│─────────────>│
│ │
│ STF ドキュメントを更新しました │
│<───────────────│
返されるドキュメントには次のものが含まれる場合があります。
権限と対話の制約。
エージェント専用のアプリケーション スペース。
使用できる正確な機能は、サーバーによって実装されている SIL 適合プロファイルによって異なります。
オブジェクトは、エージェントがプレゼンテーションから推測するのではなく、そのタイプ、役割、状態、および利用可能な機能を記述します。
SIL は、言語モデルによる直接解釈用に設計されています。タスク固有の微調整や訪問は必要ありません。

アプリケーションの基本構造を理解するためだけにモデルを使用します。
SIL 応答は、アプリケーションの現在の状態を表します。したがって、単に静的なドキュメントとして機能するのではなく、対話のために使用できます。
アクションは明示的に表現され、SIL プロトコルを通じて送信されます。エージェントは、任意の JavaScript を合成したり、UI イベント ハンドラーを再構築したりする必要はありません。
SIL は、アプリケーション定義の知識とユーザー提供のデータを構造レベルで区別します。
これにより、実装では、アプリケーションの動作を定義する情報と、データとして扱う必要がある情報との間に明示的な信頼境界を確立できます。
SIL は、アプリケーションを本質的に安全にするとは主張しません。認証、許可、検証、トランスポート セキュリティ、ビジネス ルール、および実装固有のセキュリティは、引き続きアプリケーションの責任となります。
SIL インターフェイスは HTML ページと 1 対 1 で対応する必要はありません。
アプリケーションは、エージェントとの対話専用に存在し、人間に面する対応物を持たないエージェント スペースを公開できます。
SIL は、制限された読み取り専用インターフェイスから実装を開始し、対話機能を段階的に追加できるように適合プロファイルを定義します。
この実装は、SIL v1.1.0: コア + フォーム + イベントに準拠しています。
SIL は、いくつかのメカニズムを通じて検出できます。
実装では、HTTP リンク ヘッダー、HTML <link> 要素、/.well-known/sil エンドポイント、または仕様で定義されているその他のメカニズムを通じて SIL インターフェイスをアドバタイズできます。
リンク: </.sil>; rel="シル"; type="テキスト/シル"
既知の検出エンドポイントは、利用可能な SIL リソースに関する情報を提供する場合があります。
GET /.well-known/sil
エージェントが SIL インターフェイスに入ると、SiteNavigation はアプリケーションのナビゲーション構造をセマント内に直接提供できます。

ic 表現。
検出はオプションであり、実装に依存します。サイトは、エージェントがこれらのメカニズムを使用できるように、これらのメカニズムを明示的に提供する必要があります。
カール https://ais-platform.dev/.sil
エージェントスペースを探索する
https://ais-platform.dev/フィードバック.sil
これは、フォーム、ページネーション、イベントを備えた SIL 専用のアプリケーション スペースです。
カール https://raw.githubusercontent.com/ais-space/sil/main/SIL_Open_Specific.sil
SIL と他のテクノロジーの比較
テクノロジー
主な目的
SILとの関係
HTML
人に向けた Web インターフェイス
SIL はエージェントに追加のセマンティック インターフェイスを提供します
REST API
アプリケーションリソースへのプログラムによるアクセス
SIL は、アプリケーションのセマンティクスと対話を公開することで API を補完できます
MCP
ツールとリソースへの標準化されたアクセス
SIL と MCP は異なるレイヤーに対応しており、一緒に使用できます
OpenAPI / JSON スキーマ
API とデータ構造について説明する
SIL は対話型アプリケーションとそのセマンティック状態を記述します。
構造化データ/メタデータ
機械可読情報をドキュメントに追加する
SIL は、分離されたメタデータではなく、完全なエージェント向けインターフェイスを提供します
したがって、SIL は、HTML、REST、OpenAPI、または MCP を置き換えることを目的としたものではありません。
それは別の層を占めます。
正規の仕様は次のとおりです。
SIL_Open_Specific.sil — STF の正規仕様
SIL_Concept.md — 概念的な背景
grammar/stf.ebnf — 正式な STF 文法
schema/intent.schema.json — インテント リクエストの JSON スキーマ
この仕様は自己ホスト型であり、それ自体が有効な STF ドキュメントとして表されます。
この仕様は実用的な実装に基づいて作成され、複数の AI モデルを使用した反復的なレビューを通じて改良されました。
実装の経験は最終仕様に先立って行われました。その後、仕様は実装のアーキテクチャ、技術スタック、その他のプロの機能から分離されました。

ジェクト固有の詳細。
シル/
§── SIL_Open_仕様.sil
§── SIL_Concept.md
§── README.md
§── CHANGELOG.md
§── ROADMAP.md
§── ライセンス
§── ライセンスコード
§── 投稿.md
§── SECURITY.md
§── CODE_OF_CONDUCT.md
§── 文法/
│ └── stf.ebnf
§── スキーマ/
│ └── インテント.スキーマ.json
└── 例/
適合性
実装は 1 つ以上の SIL プロファイルへの適合を主張する場合があります。
アプリケーション定義のセキュリティ境界。
読み取り専用のセマンティック インターフェイス。
フォームフィールド、入力処理、および検証。
対応する HTML を持たない SIL インターフェイス。
規格適合要件については、仕様の §8 を参照してください。
検索エンジン最適化フォーマット。
アプリケーションのビジネス ロジックを複製する要件。
言語モデルは特別な生物学的言語または「ネイティブ」言語を持っているという主張。
安全でないアプリケーションを自動的に安全にするセキュリティ メカニズム。
SIL は、アプリケーションのセマンティクスをインテリジェント エージェントが直接利用できるようにすることを目的としたインターフェイス層です。
ウェブは歴史的に、主に人間の対話を中心に構築されてきました。
SIL は単純な拡張機能を検討します。
インテリジェント エージェントがユーザーとして Web を使用する場合は、インテリジェント エージェント用に設計されたインターフェイスも必要です。
人間はビジュアル インターフェイスを使い続けることができます。
エージェントはセマンティック インターフェイスを使用できます。
どちらも同じアプリケーション上で動作できます。
仕様: クリエイティブ コモンズ表示 4.0 インターナショナル (CC-BY-4.0)
参照コード、テスト、ツール: Apache License 2.0
実装はこれらのライセンスを継承せず、他のライセンスに基づいてリリースされる場合があります。
仕様の問題: GitHub の問題
セキュリティ: security@ais-platform.dev
機械可読ウェブのオープン仕様。セルフホスティング (独自の形式で記述)、3,305 行、5 コ

情報プロファイル、40 のテスト ベクトル。 CC-BY-4.0。
Readme ライセンス、Apache-2.0 l

[切り捨てられた]

## Original Extract

Open specification for a machine-readable web. Self-hosting (written in its own format), 3,305 lines, 5 conformance profiles, 40 test vectors. CC-BY-4.0. - ais-space/sil

GitHub - ais-space/sil: Open specification for a machine-readable web. Self-hosting (written in its own format), 3,305 lines, 5 conformance profiles, 40 test vectors. CC-BY-4.0. · GitHub
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
ais-space
/
sil
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
7 Commits 7 Commits Folders and files
examples examples grammar grammar schema schema CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE LICENSE-CODE LICENSE-CODE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SIL_Concept.md SIL_Concept.md SIL_Open_Specification.sil SIL_Open_Specification.sil View all files Repository files navigation
SIL — Semantic Interface Layer
A text-based interface protocol for intelligent agents.
SIL (Semantic Interface Layer) is an open specification for providing AI agents with a direct, semantic interface to web applications.
A conventional web application is primarily designed for human interaction. Its interface exposes information through visual layout, navigation, controls, client-side behavior, and other presentation mechanisms. An agent can use that interface, but often has to reconstruct the application's structure and meaning from those presentation artifacts.
SIL adds another interface to the same application: a structured textual representation of its application state, available objects, relationships, navigation, capabilities, and actions.
The human interface remains. SIL does not replace it.
The goal is to give an agent the equivalent of a well-defined interface rather than requiring it to reverse-engineer one.
One application, two interfaces
The two interfaces can share the same backend, business logic, data, authentication, authorization, and application state.
┌─────────────────────┐
│ Application │
│ │
│ business logic │
│ data & state │
│ permissions │
└──────────┬──────────┘
│
┌─────────────┴─────────────┐
│ │
▼ ▼
Human Interface Semantic Interface
HTML SIL / STF
│ │
▼ ▼
Humans Agents
This is not necessarily duplication of information. It is a separation of presentation according to the capabilities and requirements of the consumer.
Agents can already interact with web pages, APIs, MCP tools, and other interfaces. SIL addresses a different problem:
How can an agent understand an application itself — its structure, current state, available actions, navigation, and boundaries — without reconstructing that information from presentation-oriented artifacts?
For example, an HTML page may contain a button whose meaning is determined by its surrounding elements, styling, JavaScript handlers, application state, and visual context.
A SIL representation can state that meaning directly:
DownloadButton
Type: Button
Role: PrimaryAction
Id: btn_download
Caption: Download
Actions: Activate
The difference is not that one contains information and the other does not.
The difference is where the interpretation has to happen .
With a conventional interface, the agent may have to infer semantics from presentation. With SIL, the application can expose those semantics explicitly.
SIL defines an application vocabulary and interface protocol for intelligent agents.
As of SIL v1.1.0, the underlying syntax and semantic core are no longer defined inline. SIL is now specified as an extension of STF core — a separate, standalone specification (Semantic Text Format) published independently at github.com/ais-space/stf .
STF core (standalone specification)
├── syntax (§2)
├── semantic model (§3)
├── structural security (§4)
└── data / serialization model (§5, §6)
SIL v1.1.0 (this specification)
├── SIL Vocabulary — application extension of STF core (§3)
├── Application Vocabulary (§4)
├── SIL Protocol — discovery & interaction (§5)
├── Interaction Patterns (§6)
├── Security-specific guidance (§7)
├── Conformance profiles (§8)
└── Extensions & Future Directions (§9)
STF core is referenced normatively, by reference, without version pinning : any SIL document is, by construction, a valid STF document. When STF core evolves, SIL tracks it at the application layer. STF core is intentionally textual and hierarchical, designed to remain readable by humans while being straightforward for language models to interpret.
The SIL specification is itself written in STF, making it a self-describing, conformant document rather than merely a document about the format.
A SIL-enabled application exposes semantic pages through HTTP.
A typical interaction looks like this:
Agent SIL Server
│ │
│ GET /revizor.sil │
│────────────────────────────>│
│ │
│ STF document │
│<────────────────────────────│
│ │
│ POST /revizor.sil │
│ {"intent":"activate", │
│ "target":"btn_download"} │
│────────────────────────────>│
│ │
│ Updated STF document │
│<────────────────────────────│
The returned document can contain:
permissions and interaction constraints;
agent-only application spaces.
The exact capabilities available depend on the SIL conformance profile implemented by the server.
Objects describe their type, role, state, and available capabilities instead of requiring the agent to infer them from presentation.
SIL is designed for direct interpretation by language models. It does not require task-specific fine-tuning or a vision model merely to understand the application's basic structure.
A SIL response represents the application's current state. It can therefore be used for interaction rather than serving merely as static documentation.
Actions are represented explicitly and submitted through the SIL protocol. The agent does not need to synthesize arbitrary JavaScript or reconstruct UI event handlers.
SIL distinguishes application-defined knowledge from user-provided data at the structural level.
This allows an implementation to establish explicit trust boundaries between information that defines application behavior and information that must be treated as data.
SIL does not claim to make an application inherently secure. Authentication, authorization, validation, transport security, business rules, and implementation-specific security remain the responsibility of the application.
A SIL interface does not have to correspond one-to-one with an HTML page.
An application can expose Agent Spaces that exist specifically for agent interaction and have no human-facing counterpart.
SIL defines conformance profiles so that an implementation can start with a limited read-only interface and add interaction capabilities progressively.
This implementation conforms to SIL v1.1.0: Core + Forms + Events.
SIL can be discovered through several mechanisms.
An implementation may advertise its SIL interface through HTTP Link headers, HTML <link> elements, the /.well-known/sil endpoint, or other mechanisms defined by the specification.
Link: </.sil>; rel="sil"; type="text/sil"
A well-known discovery endpoint may provide information about available SIL resources:
GET /.well-known/sil
Once an agent enters a SIL interface, SiteNavigation can provide the application's navigation structure directly within the semantic representation.
Discovery is optional and implementation-dependent: a site must explicitly provide these mechanisms for an agent to use them.
curl https://ais-platform.dev/.sil
Explore an Agent Space
https://ais-platform.dev/feedback.sil
This is a SIL-only application space with forms, pagination, and events.
curl https://raw.githubusercontent.com/ais-space/sil/main/SIL_Open_Specification.sil
SIL compared with other technologies
Technology
Primary purpose
Relationship to SIL
HTML
Human-facing web interface
SIL provides an additional semantic interface for agents
REST APIs
Programmatic access to application resources
SIL can complement APIs by exposing application semantics and interaction
MCP
Standardized access to tools and resources
SIL and MCP address different layers and can be used together
OpenAPI / JSON Schema
Describe APIs and data structures
SIL describes an interactive application and its semantic state
Structured data / metadata
Add machine-readable information to documents
SIL provides a complete agent-facing interface rather than isolated metadata
SIL is therefore not intended to replace HTML, REST, OpenAPI, or MCP.
It occupies a different layer.
The canonical specification is:
SIL_Open_Specification.sil — canonical specification in STF
SIL_Concept.md — conceptual background
grammar/stf.ebnf — formal STF grammar
schema/intent.schema.json — JSON Schema for intent requests
The specification is self-hosting: it is itself represented as a valid STF document.
The specification originated from a working implementation and was refined through iterative review with multiple AI models.
The implementation experience preceded the final specification. The specification was then separated from the implementation's architecture, technology stack, and other project-specific details.
sil/
├── SIL_Open_Specification.sil
├── SIL_Concept.md
├── README.md
├── CHANGELOG.md
├── ROADMAP.md
├── LICENSE
├── LICENSE-CODE
├── CONTRIBUTING.md
├── SECURITY.md
├── CODE_OF_CONDUCT.md
├── grammar/
│ └── stf.ebnf
├── schema/
│ └── intent.schema.json
└── examples/
Conformance
An implementation may claim conformance to one or more SIL profiles.
application-defined security boundaries.
Read-only semantic interfaces.
Form fields, input handling, and validation.
SIL interfaces without an HTML counterpart.
See §8 of the specification for the normative conformance requirements.
a search-engine optimization format;
a requirement to duplicate an application's business logic;
a claim that language models possess a special biological or "native" language;
a security mechanism that makes unsafe applications safe automatically.
SIL is an interface layer intended to make an application's semantics directly available to intelligent agents.
The web has historically been built primarily around human interaction.
SIL explores a simple extension:
If intelligent agents are going to use the web as users, they should have an interface designed for them too.
Humans can continue using the visual interface.
Agents can use the semantic interface.
Both can operate on the same application.
Specification: Creative Commons Attribution 4.0 International (CC-BY-4.0)
Reference code, tests, and tooling: Apache License 2.0
Implementations do not inherit these licenses and may be released under other licenses.
Specification issues: GitHub Issues
Security: security@ais-platform.dev
Open specification for a machine-readable web. Self-hosting (written in its own format), 3,305 lines, 5 conformance profiles, 40 test vectors. CC-BY-4.0.
Readme License, Apache-2.0 l

[truncated]
