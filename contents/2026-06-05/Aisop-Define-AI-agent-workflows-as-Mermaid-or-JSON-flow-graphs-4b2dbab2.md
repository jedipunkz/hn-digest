---
source: "https://github.com/AIXP-Labs/AISOP"
hn_url: "https://news.ycombinator.com/item?id=48407235"
title: "Aisop – Define AI agent workflows as Mermaid or JSON flow graphs"
article_title: "GitHub - AIXP-Labs/AISOP: AISOP — AI Standard Operating Protocol · GitHub"
author: "aisop"
captured_at: "2026-06-05T04:34:25Z"
capture_tool: "hn-digest"
hn_id: 48407235
score: 1
comments: 0
posted_at: "2026-06-05T02:28:02Z"
tags:
  - hacker-news
  - translated
---

# Aisop – Define AI agent workflows as Mermaid or JSON flow graphs

- HN: [48407235](https://news.ycombinator.com/item?id=48407235)
- Source: [github.com](https://github.com/AIXP-Labs/AISOP)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T02:28:02Z

## Translation

タイトル: Aisop – AI エージェントのワークフローをマーメイドまたは JSON フロー グラフとして定義
記事のタイトル: GitHub - AIXP-Labs/AISOP: AISOP — AI 標準オペレーティング プロトコル · GitHub
説明: AISOP — AI 標準オペレーティング プロトコル。 GitHub でアカウントを作成して、AIXP-Labs/AISOP の開発に貢献してください。

記事本文:
GitHub - AIXP-Labs/AISOP: AISOP — AI 標準オペレーティング プロトコル · GitHub
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
AIXP-Labs
/
アイソップ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード O

ペンのその他のアクション メニュー フォルダーとファイル
2 コミット 2 コミット adrs adrs docs docs リファレンス リファレンス 仕様 仕様 .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md COTRIBUTING_CN.md CONTRIBUTING_CN.md GOVERNANCE.md GOVERNANCE.md ライセンス ライセンス通知 通知 README.md README.md README_CN.md README_CN.md SECURITY.md SECURITY.md mkdocs.yml mkdocs.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AISOP — AI 標準オペレーティング プロトコル
Mermaid または JSON フロー グラフを使用して構造化 AI プログラムを定義するためのオープン プロトコル。
AISOP を使用すると、視覚的 (マーメイド) または構造的 (JSON) の制御フロー (分岐、並列実行、サブタスク、エラー処理) をすべて単一のポータブル JSON 形式で定義できるようになります。
アプローチ
定義
ポータブル
機械可読
トークン効率が高い
自然言語プロンプト
フリーテキスト
はい
いいえ
いいえ
Python / YAML ワークフロー
コード/構成
いいえ
部分的に
いいえ
ビジュアルビルダー (Dify など)
独自の
いいえ
いいえ
いいえ
アイソップ
マーメイド + JSON
はい
はい
はい
AISOP ファイルは、フロー グラフ (Mermaid または JSON) を含むプレーンな JSON であり、あらゆる言語で読み取り可能、あらゆるエディターで編集可能、git でバージョン管理可能で、LLM の理解用に最適化されています。
[
{ "役割" : " システム " , "コンテンツ" : { "プロトコル" : " AISOP V1.0.0 " , ... } },
{ "役割" : " ユーザー " 、 "コンテンツ" : { "命令" : " ... " 、 "aisop" : { ... }、 "関数" : { ... } } }
]
要素
目的
システム
プログラムのメタデータ (ID、バージョン、説明)
ユーザー
命令、フローグラフ、関数定義
簡単な例
[
{
"役割" : " システム " 、
「コンテンツ」: {
"プロトコル" : "AISOP V1.0.0" ,
"axiom_0" : "人間の主権と幸福" ,
"id" : "greeting_assistant " ,
"name" : " 挨拶アシスタント " ,
"バージョン" : " 1.0.0 " 、
"summary" : " あなたを分類します

ユーザーの意図と応答 "、
"flow_format" : "人魚"
}
}、
{
"ロール" : "ユーザー" ,
「コンテンツ」: {
"命令" : " aisop.main を実行します " ,
"user_input" : " {user_input} " ,
"アイソップ" : {
"main" : " グラフ TD \n 挨拶[挨拶] --> 分類{分類} \n 分類 -->|質問|検索[検索] \n 分類 -->|チャット| Reply[返信] \n 検索 --> end_node((End)) \n 返信 --> end_node "
}、
「関数」: {
"greet" : { "step1" : " ユーザーに挨拶します " },
"classify" : { "step1" : " ユーザーの意図を「質問」または「チャット」に分類します " },
"search" : { "step1" : " 関連情報を検索 " },
"reply" : { "step1" : " フレンドリーな返信を生成します " },
"end_node" : { "step1" : " 最終応答を出力する " }
}
}
}
]
フロー グラフ (マーメイド形式):
グラフTD
挨拶[挨拶] --> 分類{分類}
分類 -->|質問|検索[検索]
分類 -->|チャット|返信[返信]
検索 --> end_node((終了))
返信 --> end_node
読み込み中
主な特長
デュアル フロー フォーマット — マーメイド文字列 (AI ネイティブ) または JSON フロー オブジェクト (コード ネイティブ)、同じプログラム内で混合可能
14 以上の制御フロー パターン — シーケンシャル、ディシジョン、パラレル フォーク/ジョイン、デリゲート、ループ、コンバージェンス、エラー ルーティング、バッチ反復、再試行、データ分離、ステップレベルのサブタスク、エージェント ディスパッチ、HITL 確認、ランタイム アサーション
3 層の分離 — メタデータ、フロー グラフ、関数定義
トポロジ / 動作の分割 - フロー グラフは接続を定義し、関数は実行時の動作を定義します
サブタスクのサポート — デリゲート ノードまたはステップレベルの RUN aisop.<sub> を介したメイン + 名前付きサブタスク
予約キー — 関数内の 8 つの実行時動作キー (join、map、on_error、retry_policy、context_filter、output_mapping、constraints、execute_mode)
sys.* システムコール — 人間による確認 (🔒 inviolable sys.io.confirm )、I/O、アサーション、モデル呼び出し、C 用の 24 個のプロトコル予約システムコール

ode の実行と状態管理 (§6 を参照)
エラー処理 — Mermaid -.-> エラー エッジ + 関数レベルの on_error タイプのルーティング + 6 つの sys.* エラー タイプ
変数置換 — {system_prompt} 、 {user_input} は実行時に置き換えられます
トークンの効率化 — Mermaid は、JSON フロー形式よりもトークンの使用量が最大 50% 少ない
依存関係ゼロ — Python および JavaScript のリファレンス実装 (stdlib のみ)
AI に依存しない — あらゆる AI ランタイムで動作します
AISOP は 2 つのフロー グラフ形式をサポートします。両方を同じプログラム内で共存できます。
形状
構文
意味
長方形
名前[テキスト]
プロセス — 実行して続行する
ダイヤモンド
名前{テキスト}
決定 — 条件分岐
サークル
名前((テキスト))
終了 — タスクを終了する
長方形
名前[aisop.sub]
デリゲート — サブタスクの呼び出し
エッジ
構文
意味
実線の矢印
-->
通常の流れ（次へ）
ラベル付き矢印
-->|ラベル|
分岐(条件付き)
破線の矢印
-.->
エラールーティング
JSONフロー(オブジェクト)
フィールド
種類
意味
次へ
文字列[]
次のノード。 1 = シーケンシャル、2+ = パラレル フォーク
枝
オブジェクト
条件分岐: {ラベル: ターゲット}
エラー
文字列
エラーハンドラーノード
デリゲート先
文字列
サブタスクを名前で呼び出す
待ってください
文字列[]
続行する前にノードを待機します (参加)
{}
—
終了 — タスクを終了する
制御フローパターン
パターン
人魚 (§4.2)
JSON フロー (§4.3)
関数 (§5) / システム (§6)
シーケンシャル
A --> B
"次": ["B"]
—
もし/そうでなければ
A -->|はい| B 、A -->|いいえ| C
"分岐": {"はい": "B"、"いいえ": "C"}
—
スイッチ (N-way)
A -->|ラベル| B、...
"ブランチ": {"ラベル": "B", ...}
—
平行フォーク
A --> B 、A --> C
"次": ["B", "C"]
—
並列結合
B --> D、C --> D
"wait_for": ["B", "C"]
"結合": {"マージ戦略": "配列"}
ループ
分岐先→前ノード
分岐先→前ノード
—
サブタスク
A[アイソップサブ]
"デリゲート先": "サブ"
—
収束
複数→同じターゲット
複数→同じターゲット
—
エラールーティング

A -.-> E
"エラー": "E"
"on_error": {"タイムアウト": "t"}
バッチの反復処理
—
—
"マップ": {"アイテムパス": "..."}
再試行
—
—
"retry_policy": {"max_attempts": 3}
データの分離
—
—
"context_filter": {"include": [...]}
ステップレベルのサブ
—
—
ステップテキスト: RUN aisop.sub
エージェント派遣
—
—
"実行モード": "エージェント"
HITL確認
—
—
sys.io.confirm('...') (§6.2)
ランタイムアサーション
—
—
sys.assert('...', '...') (§6.5)
バックグラウンドジョブ
—
—
sys.run.bg('...') (§6.4)
はじめに
CD リファレンス/Python
python run.py example.aisop.json
Python test_all.py
JavaScript:
CD リファレンス/JavaScript
ノード run.js example.aisop.json
ノードtest_all.js
プロジェクトの構造
AISOPプロトコル/
仕様/
aisop-spec.md # プロトコル仕様書（V1.0.0）
aisop-spec_CN.md # プロトコル仕様 — 中国語
参照/
パイソン/
flow_runtime.py # Python リファレンス実装
example.aisop.json # フロー定義の例
example_syscalls.aisop.json # sys.* 呼び出しを使用した高度な例
test_all.py # テストスイート (44 個のテスト)
run.py # CLI ツール
pyproject.toml # Python パッケージのメタデータ
ライセンス番号 Apache 2.0 ライセンス (パッケージごと)
通知 # Apache 2.0 帰属通知 (パッケージごと)
README.md # 実装ガイド
JavaScript/
flow_runtime.js # JavaScript リファレンス実装
example.aisop.json # フロー定義の例
example_syscalls.aisop.json # sys.* 呼び出しを使用した高度な例
test_all.js # テストスイート (44 個のテスト)
run.js # CLI ツール
package.json # npm パッケージのメタデータ
ライセンス番号 Apache 2.0 ライセンス (パッケージごと)
通知 # Apache 2.0 帰属通知 (パッケージごと)
README.md # 実装ガイド
ドキュメント/
Index.md # ドキュメント サイトのホーム
トピックス/
what-is-aisop.md # 導入と理論的根拠
Compare-mermaid-vs-jsonflow.md # Mermaid と JSON フローの比較
比較-mermaid-vs-jsonflow_CN.md # 比較

息子 — 中国人
Performance-mermaid-vs-jsonflow.md # パフォーマンス分析
Performance-mermaid-vs-jsonflow_CN.md # パフォーマンス分析 — 中国語
広告/
adr-001-mermaid-as-primary-format.md # ADR ミラー (mkdocs サイト)
adr-002-sys-io-confirm-immutability.md # ADR ミラー (mkdocs サイト)
広告/
adr-template.md # ADR テンプレート
adr-001-mermaid-as-primary-format.md # フォーマット選択の根拠
adr-002-sys-io-confirm-immutability.md # Axiom 0 施行決定
README.md
README_CN.md
COTRIBUTING.md # 貢献ガイドライン
COTRIBUTING_CN.md #贡献指南
CODE_OF_CONDUCT.md # コミュニティ標準 (公理 0 誓約)
GOVERNANCE.md # 三者チェーンのガバナンス モデル
SECURITY.md # 脆弱性報告ポリシー
ライセンス番号 Apache 2.0 ライセンス
通知 # Apache 2.0 帰属に関する通知
CHANGELOG.md # リリースノート
mkdocs.yml # MkDocs ドキュメント サイトの構成
AIXP ラボ aixp.dev
AIXP Labs は、次のコア プロジェクトを開発および維持します。
AISOP は、公理 0: 人間の主権と福祉に基づいています。AI システムは、人類に取って代わったり支配したりするのではなく、人類に奉仕するために存在します。すべての実装では、この公理を最高の実行優先順位で強制する必要があります。全文については仕様の §0 を参照してください。
ガイドラインについては CONTRIBUTING.md を参照してください / CONTRIBUTING_CN.md を参照してください。
行動規範 — Axiom 0 に沿ったコミュニティ基準
ガバナンス — 三者チェーン モデル (シード / 権限 / 執行者)
セキュリティ ポリシー — 脆弱性レポートと sys.io.confirm 不変性
アーキテクチャの意思決定記録 — 主要な意思決定の設計理論的根拠
このプロトコル仕様とリファレンス実装は、研究および教育目的のみに提供されています。これらは実験的なものであり、実稼働での使用を目的としたものではありません。ご自身の責任でご使用ください。作者は、生じたいかなる損害についても責任を負いません

このソフトウェアの使用から。全条件については、ライセンスを参照してください (Apache 2.0)。
Apache ライセンス 2.0 - Copyright 2026 AIXP Labs AIXP.dev | AISOP.dev
公理 0: 人間の主権と福祉を調整します。バージョン: AISOP V1.0.0。 www.aisop.dev
AISOP — AI 標準オペレーティング プロトコル
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
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

AISOP — AI Standard Operating Protocol. Contribute to AIXP-Labs/AISOP development by creating an account on GitHub.

GitHub - AIXP-Labs/AISOP: AISOP — AI Standard Operating Protocol · GitHub
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
AIXP-Labs
/
AISOP
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits adrs adrs docs docs reference reference specification specification .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTING_CN.md CONTRIBUTING_CN.md GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE NOTICE NOTICE README.md README.md README_CN.md README_CN.md SECURITY.md SECURITY.md mkdocs.yml mkdocs.yml View all files Repository files navigation
AISOP — AI Standard Operating Protocol
An open protocol for defining structured AI programs using Mermaid or JSON flow graphs.
AISOP enables defining multi-step AI programs with visual (Mermaid) or structural (JSON) control flow — branching, parallel execution, sub-tasks, and error handling — all in a single portable JSON format.
Approach
Definition
Portable
Machine-readable
Token-efficient
Natural language prompts
Free text
Yes
No
No
Python / YAML workflows
Code / config
No
Partially
No
Visual builders (Dify, etc.)
Proprietary
No
No
No
AISOP
Mermaid + JSON
Yes
Yes
Yes
AISOP files are plain JSON with flow graphs (Mermaid or JSON) — readable by any language, editable in any editor, versionable in git, and optimized for LLM comprehension.
[
{ "role" : " system " , "content" : { "protocol" : " AISOP V1.0.0 " , ... } },
{ "role" : " user " , "content" : { "instruction" : " ... " , "aisop" : { ... }, "functions" : { ... } } }
]
Element
Purpose
system
Program metadata (identity, version, description)
user
Instruction, flow graph, and function definitions
Quick Example
[
{
"role" : " system " ,
"content" : {
"protocol" : " AISOP V1.0.0 " ,
"axiom_0" : " Human_Sovereignty_and_Wellbeing " ,
"id" : " greeting_assistant " ,
"name" : " Greeting Assistant " ,
"version" : " 1.0.0 " ,
"summary" : " Classify user intent and respond " ,
"flow_format" : " mermaid "
}
},
{
"role" : " user " ,
"content" : {
"instruction" : " RUN aisop.main " ,
"user_input" : " {user_input} " ,
"aisop" : {
"main" : " graph TD \n greet[Greet] --> classify{Classify} \n classify -->|question| search[Search] \n classify -->|chat| reply[Reply] \n search --> end_node((End)) \n reply --> end_node "
},
"functions" : {
"greet" : { "step1" : " Say hello to the user " },
"classify" : { "step1" : " Classify user intent as 'question' or 'chat' " },
"search" : { "step1" : " Search for relevant info " },
"reply" : { "step1" : " Generate a friendly reply " },
"end_node" : { "step1" : " Output final response " }
}
}
}
]
The flow graph (Mermaid format):
graph TD
greet[Greet] --> classify{Classify}
classify -->|question| search[Search]
classify -->|chat| reply[Reply]
search --> end_node((End))
reply --> end_node
Loading
Key Features
Dual Flow Format — Mermaid string (AI-native) or JSON flow object (code-native), mixable in the same program
14+ Control Flow Patterns — sequential, decision, parallel fork/join, delegate, loop, convergence, error routing, batch iterate, retry, data isolation, step-level sub-task, agent dispatch, HITL confirmation, runtime assertion
Three-Layer Separation — metadata, flow graph, function definitions
Topology / Behavior Split — flow graph defines connections, functions define runtime behavior
Sub-task Support — main + named sub-tasks via delegate nodes or step-level RUN aisop.<sub>
Reserved Keys — 8 runtime behavior keys in functions (join, map, on_error, retry_policy, context_filter, output_mapping, constraints, execute_mode)
sys.* System Calls — 24 protocol-reserved system calls for human confirmation (🔒 inviolable sys.io.confirm ), I/O, assertions, model invocation, code execution, and state management (see §6)
Error Handling — Mermaid -.-> error edge + function-level on_error type routing + 6 sys.* error types
Variable Substitution — {system_prompt} , {user_input} replaced at runtime
Token Efficient — Mermaid uses ~50% fewer tokens than JSON flow format
Zero Dependencies — Reference implementations in Python and JavaScript, stdlib only
AI-Agnostic — Works with any AI runtime
AISOP supports two flow graph formats. Both can coexist in the same program.
Shape
Syntax
Meaning
Rectangle
name[text]
Process — execute and proceed
Diamond
name{text}
Decision — conditional branching
Circle
name((text))
End — terminate task
Rectangle
name[aisop.sub]
Delegate — call sub-task
Edge
Syntax
Meaning
Solid arrow
-->
Normal flow (next)
Labeled arrow
-->|label|
Branch (conditional)
Dashed arrow
-.->
Error routing
JSON Flow (object)
Field
Type
Meaning
next
string[]
Next node(s). 1 = sequential, 2+ = parallel fork
branches
object
Conditional branching: {label: target}
error
string
Error handler node
delegate_to
string
Call a sub-task by name
wait_for
string[]
Wait for nodes before proceeding (join)
{}
—
End — terminate task
Control Flow Patterns
Pattern
Mermaid (§4.2)
JSON Flow (§4.3)
Functions (§5) / sys (§6)
Sequential
A --> B
"next": ["B"]
—
If/else
A -->|yes| B , A -->|no| C
"branches": {"yes": "B", "no": "C"}
—
Switch (N-way)
A -->|label| B , ...
"branches": {"label": "B", ...}
—
Parallel fork
A --> B , A --> C
"next": ["B", "C"]
—
Parallel join
B --> D , C --> D
"wait_for": ["B", "C"]
"join": {"merge_strategy": "array"}
Loop
Branch target → earlier node
Branch target → earlier node
—
Sub-task
A[aisop.sub]
"delegate_to": "sub"
—
Convergence
Multiple → same target
Multiple → same target
—
Error routing
A -.-> E
"error": "E"
"on_error": {"timeout": "t"}
Batch iterate
—
—
"map": {"items_path": "..."}
Retry
—
—
"retry_policy": {"max_attempts": 3}
Data isolation
—
—
"context_filter": {"include": [...]}
Step-level sub
—
—
step text: RUN aisop.sub
Agent dispatch
—
—
"execute_mode": "agent"
HITL confirmation
—
—
sys.io.confirm('...') (§6.2)
Runtime assertion
—
—
sys.assert('...', '...') (§6.5)
Background job
—
—
sys.run.bg('...') (§6.4)
Getting Started
cd reference/python
python run.py example.aisop.json
python test_all.py
JavaScript:
cd reference/javascript
node run.js example.aisop.json
node test_all.js
Project Structure
AISOP-Protocol/
specification/
aisop-spec.md # Protocol specification (V1.0.0)
aisop-spec_CN.md # Protocol specification — Chinese
reference/
python/
flow_runtime.py # Python reference implementation
example.aisop.json # Example flow definition
example_syscalls.aisop.json # Advanced example with sys.* calls
test_all.py # Test suite (44 tests)
run.py # CLI tool
pyproject.toml # Python package metadata
LICENSE # Apache 2.0 license (per-package)
NOTICE # Apache 2.0 attribution notice (per-package)
README.md # Implementation guide
javascript/
flow_runtime.js # JavaScript reference implementation
example.aisop.json # Example flow definition
example_syscalls.aisop.json # Advanced example with sys.* calls
test_all.js # Test suite (44 tests)
run.js # CLI tool
package.json # npm package metadata
LICENSE # Apache 2.0 license (per-package)
NOTICE # Apache 2.0 attribution notice (per-package)
README.md # Implementation guide
docs/
index.md # Documentation site home
topics/
what-is-aisop.md # Introduction and rationale
comparison-mermaid-vs-jsonflow.md # Mermaid vs JSON Flow comparison
comparison-mermaid-vs-jsonflow_CN.md # Comparison — Chinese
performance-mermaid-vs-jsonflow.md # Performance analysis
performance-mermaid-vs-jsonflow_CN.md # Performance analysis — Chinese
adrs/
adr-001-mermaid-as-primary-format.md # ADR mirror (mkdocs site)
adr-002-sys-io-confirm-immutability.md # ADR mirror (mkdocs site)
adrs/
adr-template.md # ADR template
adr-001-mermaid-as-primary-format.md # Format choice rationale
adr-002-sys-io-confirm-immutability.md # Axiom 0 enforcement decision
README.md
README_CN.md
CONTRIBUTING.md # Contribution guidelines
CONTRIBUTING_CN.md # 贡献指南
CODE_OF_CONDUCT.md # Community standards (Axiom 0 pledge)
GOVERNANCE.md # Tripartite Chain governance model
SECURITY.md # Vulnerability reporting policy
LICENSE # Apache 2.0 license
NOTICE # Apache 2.0 attribution notice
CHANGELOG.md # Release notes
mkdocs.yml # MkDocs documentation site config
AIXP Labs aixp.dev
AIXP Labs develops and maintains the following core projects:
AISOP is grounded in Axiom 0: Human Sovereignty and Wellbeing — AI systems exist to serve humanity, not replace or dominate it. All implementations must enforce this axiom at the highest execution priority. See §0 of the specification for the full text.
See CONTRIBUTING.md for guidelines / 见 CONTRIBUTING_CN.md 了解贡献指南。
Code of Conduct — Community standards aligned with Axiom 0
Governance — Tripartite Chain model (Seed / Authority / Executor)
Security Policy — Vulnerability reporting and sys.io.confirm immutability
Architecture Decision Records — Design rationale for major decisions
This protocol specification and the reference implementations are provided for research and educational purposes only . They are experimental and not intended for production use. Use at your own risk. The authors assume no liability for any damages arising from use of this software. See LICENSE for full terms (Apache 2.0).
Apache License 2.0 - Copyright 2026 AIXP Labs AIXP.dev | AISOP.dev
Align Axiom 0: Human Sovereignty and Wellbeing. Version: AISOP V1.0.0. www.aisop.dev
AISOP — AI Standard Operating Protocol
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
