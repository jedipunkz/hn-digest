---
source: "https://github.com/RobertBateman/thoughttree-framework"
hn_url: "https://news.ycombinator.com/item?id=48701254"
title: "Open handoff: Thought Tree, a markup/spec idea for modular LLM workflows"
article_title: "GitHub - RobertBateman/thoughttree-framework: An open framework for cognitive programming with LLMs: structured, reusable, inspectable workflows built from Data Units, Operations, Modules, Contracts and Traces. · GitHub"
author: "xavier1764"
captured_at: "2026-06-27T20:23:00Z"
capture_tool: "hn-digest"
hn_id: 48701254
score: 1
comments: 0
posted_at: "2026-06-27T20:06:28Z"
tags:
  - hacker-news
  - translated
---

# Open handoff: Thought Tree, a markup/spec idea for modular LLM workflows

- HN: [48701254](https://news.ycombinator.com/item?id=48701254)
- Source: [github.com](https://github.com/RobertBateman/thoughttree-framework)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T20:06:28Z

## Translation

タイトル: オープンハンドオフ: 思考ツリー、モジュラー LLM ワークフローのマークアップ/仕様のアイデア
記事のタイトル: GitHub - RobertBateman/thoughttree-framework: LLM を使用したコグニティブ プログラミングのためのオープン フレームワーク: データ ユニット、オペレーション、モジュール、コントラクト、トレースから構築された、構造化され再利用可能で検査可能なワークフロー。 · GitHub
説明: LLM を使用したコグニティブ プログラミングのオープン フレームワーク: データ ユニット、オペレーション、モジュール、コントラクト、トレースから構築された、構造化され、再利用可能で、検査可能なワークフロー。 - RobertBateman/thoughttree-framework

記事本文:
GitHub - RobertBateman/thoughttree-framework: LLM を使用したコグニティブ プログラミングのオープン フレームワーク: データ ユニット、オペレーション、モジュール、コントラクト、トレースから構築された、構造化され再利用可能で検査可能なワークフロー。 · GitHub
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
ロバート・ベイトマン
/
あなた

ghtree フレームワーク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
RobertBateman/thoughttree-framework
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット ドキュメント ドキュメントのサンプル サンプル プロトタイプ プロトタイプ 仕様 仕様 COTRIBUTING.md CONTRIBUTING.md GOVERNANCE.md GOVERNANCE.md HANDOFF.md HANDOFF.md LICENSE LICENSE LICENSE.md LICENSE.md README.md README.md ROADMAP.md ROADMAP.md STATUS.md STATUS.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM を使用したコグニティブ プログラミングのためのオープン フレームワーク。
Thought Treeは、LLMを利用した複雑な作業を、構造化され、実行可能で検査可能なワークフローとして記述するためのフレームワークです。
単一のプロンプト、不透明なエージェント ループ、またはモデル呼び出しの非公式なチェーンに依存する代わりに、思考ツリーは次のように定義します。
どのような中間成果物を生成する必要があるか。
どのような操作によってあるアーティファクトが別のアーティファクトに変換されるか。
出力をどのようにレビューまたは検証するか。
どのような最終出力を返す必要があるか。
そして、どの実行トレースを保存する必要があるか。
このフレームワークの中核では、次のような単純なパターンが使用されます。
データ単位 → 操作 → データ単位
思考ツリー プログラムはこれを再帰的に適用します。複雑な認知作業は、名前付きの成果物、変換、契約、モジュール、およびトレースに分解されます。
長期的な目標は、LLM 支援の作業を、孤立したプロンプトからコグニティブ プログラミング (アイデア、文書、計画、要件、レビュー、コンセプト、およびそれらの間の変換を主な目的とする、再利用可能で検査可能なモデルに依存しないプログラム) に移行することです。
Thought Tree Frameworkは概念的に開発され、部分的にプロトタイプ化されていますが、完成した実稼働システムではありません。
原作者は家族との約束や介護の責任を理由に手を引いています。

ork のコミットメント、燃え尽き症候群からの回復、そしてプロジェクトには今や 1 人を超えるスキル、時間、専門知識が必要であるという事実。
このリポジトリは、他の人が次のことができるようにオープンにリリースされています。
既存のエージェントおよびワークフロー フレームワークと比較します。
あるいはプロジェクトをより良い方向に導くこともできます。
このリリースの目標は、フレームワークが完成したと主張することではありません。
目的は、アイデアが一人の人間に囚われないようにすることです。
LLM は、文書、要件、メモ、ストーリー、計画、レビュー、ポリシー、設計、意思決定など、人間のあいまいな概念を扱うのに強力です。
しかし、現在の LLM の使用のほとんどは依然として次の方法で行われています。
セマンティック検証を行わないワークフローの自動化。
これらのアプローチは便利ですが、実稼働システムに期待される品質が欠けていることがよくあります。
検査可能な中間結果。
ワークフロー定義とモデルプロバイダーの分離。
最終的な出力の出所を明確にします。
小さなタスクの場合は、プロンプトで十分な場合があります。
より大きな認知作業では、プロセスが重要です。
思考ツリー プログラムは、認知作業を名前付きの成果物に対する変換として説明します。
入力
↓
運営
↓
中間品
↓
レビュー/検証
↓
最終出力
↓
実行トレース
たとえば、LLM に質問する代わりに、次のようにします。
「これらのメモから技術設計文書を作成します。」
思考ツリー モジュールでは次のように定義できます。
デザインノート
↓
ソースダイジェストの作成
↓
ソースダイジェスト
↓
要件の抽出
↓
要件登録
↓
デザインシステム分解
↓
システム分解
↓
ドラフトセクション
↓
ドラフト_tdd_sections
↓
ドラフトを組み立てる
↓
ドラフト_tdd
↓
ドラフトのレビュー
↓
見直しと修正計画
↓
最終文書の改訂
↓
テクニカルデザインドキュメント
各ステップでは、検査、検証、置換、再利用、または追跡できるアーティファクトが生成されます。
##データ単位
データユニットは使用または生成される個別の成果物です

ワークフローによって。
現在の TTML ドラフトでは、データ ユニットは多くの場合ファイルとして表されますが、概念的には、アドレス可能な任意の成果物である可能性があります。
オペレーションは、入力データ ユニットを出力データ ユニットに変換します。
ソースノート → ソースダイジェストの作成 → ソースダイジェスト
ドラフト_章 + フィードバック → 改訂章 → 改訂_章
セクション→ConcatenateFiles→compiled_document
操作は次の方法で実行できます。
モジュールは再利用可能な認知プログラムです。
オプションの契約と検証の期待。
モジュールはスタンドアロン ワークフローとして実行することも、別のモジュールによって呼び出すこともできます。
コグニティブ エンジンは、思考ツリー プログラムのコンパイラーおよびランタイムです。
思考ツリー定義をロードします。
変数、入力、参照、コレクション、イテレータの解決。
ワークフローを実行可能な変換グラフにコンパイルします。
依存関係エラーと出力衝突を検出します。
LLM、関数、ツール、サブモジュール、人間によるレビュー手順の呼び出し。
中間成果物を保存する。
LLM はシステム全体ではありません。
コグニティブ エンジンは両方を調整します。
Thought Tree Markup Language (TTML) は、Thought Tree モジュールの現在のドラフト XML ベースのソース形式です。
最小限の TTML スタイルのワークフローは次のようになります。
< TTML バージョン = " 0.12.0 " >
< プロジェクト
id = "記事概要モジュール"
desc = " 草案、レビュー、改訂を通じて記事を要約します。 " />
<入力>
< ファイル ID = " source_article " フォルダー = " /inputs " 拡張子 = " txt " />
</ 入力 >
＜操作＞
< 操作方法
id = " ドラフトサマリー "
type = " テキストコンプリーション "
desc = " ソース記事の簡潔な要約を作成します。 " >
< FileRef id = " source_article " />
<出力>
< ファイル ID = "draft_summary" 拡張子 = "txt" />
</ 出力 >
〈操作〉
< 操作方法
id = "レビュー概要"
type = " テキストコンプリーション "
desc = " 草案の概要を以下の内容に照らしてレビューします。

ソース記事。欠落、不正確、および裏付けのない主張を特定します。 " >
< FileRef id = " source_article " />
< FileRef id = "draft_summary " />
<出力>
< ファイル ID = " summary_review " 拡張子 = " txt " />
</ 出力 >
〈操作〉
< 操作方法
id = "改訂概要"
type = " テキストコンプリーション "
desc = " レビューを使用して概要を修正します。最終的な概要のみを作成します。 " >
< FileRef id = " source_article " />
< FileRef id = "draft_summary " />
< FileRef id = " summary_review " />
<出力>
< ファイル ID = " Final_summary " 拡張子 = " txt " />
</ 出力 >
〈操作〉
</ 操作 >
<出力>
< FileRef id = " Final_summary " />
</ 出力 >
</TTML>
TTML がフレームワーク全体ではありません。
これは、基礎となる思考ツリー プログラム モデルの 1 つのソース表現です。
将来のソース形式には、YAML、JSON、ビジュアル グラフ、またはより高レベルのオーサリング ツールが含まれる可能性があります。
Thought Treeと他のアプローチの違い
プロンプトはモデルに答えを求めます。
思考ツリーは、答えを作成、レビュー、修正、追跡するプロセスを定義します。
プロンプト チェーンはモデル呼び出しを接続します。
モデルに依存しない実行。
自律エージェントとの比較
エージェントは柔軟ですが、予測、デバッグ、監査が難しい場合があります。
思考ツリーは明示的な構造を好みます。
定義されたプロセス
→ 検査可能プラン
→ 制御された実行
→ 保存された工芸品
→ 追跡可能な出力
動的計画は引き続き可能ですが、生成された計画またはサブモジュールを検証して記録する必要があります。
従来のワークフローツールとの比較
従来のワークフロー ツールは、決定論的なプロセスを得意としています。
Thought Treeは、ハイブリッド認知作業用に設計されています。
セマンティック変換のための LLM。
決定論的変換のためのコード。
外部機能用のツール。
判断と承認のための人間。
認知

e オーケストレーションとトレーサビリティのためのエンジン。
Thought Treeは、品質、トレーサビリティ、イテレーションが重要となる複雑なコグニティブ・プロダクション・タスクを対象としています。
散らばったメモを正式な文書に変換する。
技術設計文書の作成。
レガシープロジェクトの知識を回復する。
コンプライアンスまたは監査準備パックの作成。
クリエイティブなメディアの生成とレビュー。
ゲームコンテンツと伝承の制作。
長編小説の草稿と改訂。
ソース文書から要件を抽出する。
定期的な調査レポートを作成する。
他の思考ツリーモジュールの生成、検証、改善。
このハンドオフ パッケージには以下が含まれているか、または含まれる予定です。
思考ツリー フレームワークの完全な説明。
TTML の概念と例の草案。
思考ツリープログラムモデル。
コグニティブ エンジン アーキテクチャに関するメモ。
初期の概念実証
C#/Unity プロトタイプは、LLM ワークフローが操作間でテキスト ファイルを変数として渡すことができることを実証しました。
このプロトタイプでは TTML は使用されませんでした。ハードコードされたプロンプトの配列が使用されていました。
ユーザーが投稿した 500 語の説明から 50,000 語の小説を計画、起草、レビューすることに成功しました。
部分的な認知エンジンのプロトタイプ
その後の C#/Unity プロジェクトでは、認知エンジンの実装が開始されました。
以下に接続できます:
基本的な TextCompletion の実行をサポートします。
安定した TTML インポーター、完全なグラフ コンパイル、コントラクト、検証ゲート、実稼働ワークスペース/トレース システムはまだ含まれていません。
このフレームワークにはまだ大幅な作業が必要です。
イテレータとコレクションの解決。
セキュリティとガバナンスのモデル。
これは基礎であり、完成品ではありません。
貢献者に推奨される次のステップ
役立つ貢献分野には次のようなものがあります。
思考ツリープログラムモデルを改良する。
実行トレーススキーマを定義します。
セマンティックコントラクトスキーマを定義します。
最小限の CLI コグニティブ エンジンを実装します。
r

データユニットとFileRefを解決します。
TextCompletion 操作を実行します。
決定論的な関数を実行します。
中間成果物を保存します。
より大きなワークフローの例を作成します。
技術文書、クリエイティブライティング、コンプライアンス、調査、モジュール改善のためのサンプルを作成します。
思考ツリーを LangGraph、AutoGen、CrewAI、セマンティック カーネル、ワークフロー エンジン、エージェント システムと比較します。
フレームワークがどこで重複しているか、どこが異なっているか、または相互運用できるかを特定します。
最小限のリファレンス エンジン ターゲット
最小限の認知エンジンは次のことができる必要があります。
ドキュメントの順序で操作を実行します。
登録された決定論的関数をサポートします。
より高度なエンジンでは、次のものが追加される可能性があります。
自動化された改善ワークフロー。
思考ツリーフレームワーク/
README.md
ハンドオフ.md
ステータス.md
ロードマップ.md
貢献.md
ガバナンス.md
ライセンス.md
ドキュメント/
ONE_PAGE_OVERVIEW.md
WHY_THIS_MATTERS.md
アーキテクチャ.md
PROGRAM_MODEL.md
COGNITIVE_ENGINE.md
EXECUTION_SEMANTICS.md
SEMANTIC_CONTRACTS.md
オーソリング_ガイド.md
ThoughtTreeFramework.pdf
スペック/
TTML_DRAFT.md
TTMLSchema.xsd
EXECUTION_TRACE_SCHEMA_DRAFT.json
CONTRACT_SCHEMA_DRAFT.json
例/
記事-概要-レビュー/
新しい世代のパイプライン/
ビデオゲーム-tdd/
コンプライアンスギャップ分析/
モジュールの改善/
試作品/
Unity-novel-poc/
Unity-コグニティブ-エンジン/
このプロジェクトは以下に関連する可能性があります:
AI

[切り捨てられた]

## Original Extract

An open framework for cognitive programming with LLMs: structured, reusable, inspectable workflows built from Data Units, Operations, Modules, Contracts and Traces. - RobertBateman/thoughttree-framework

GitHub - RobertBateman/thoughttree-framework: An open framework for cognitive programming with LLMs: structured, reusable, inspectable workflows built from Data Units, Operations, Modules, Contracts and Traces. · GitHub
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
RobertBateman
/
thoughttree-framework
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
RobertBateman/thoughttree-framework
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits docs docs examples examples prototypes prototypes spec spec CONTRIBUTING.md CONTRIBUTING.md GOVERNANCE.md GOVERNANCE.md HANDOFF.md HANDOFF.md LICENSE LICENSE LICENSE.md LICENSE.md README.md README.md ROADMAP.md ROADMAP.md STATUS.md STATUS.md View all files Repository files navigation
An open framework for cognitive programming with LLMs.
Thought Tree is a framework for describing complex LLM-assisted work as structured, executable, inspectable workflows.
Instead of relying on a single prompt, an opaque agent loop, or an informal chain of model calls, a Thought Tree defines:
what intermediate artefacts should be produced;
what operations transform one artefact into another;
how outputs should be reviewed or validated;
what final outputs should be returned;
and what execution trace should be preserved.
At its core, the framework uses a simple pattern:
Data Units → Operations → Data Units
A Thought Tree program applies this recursively. Complex cognitive work is decomposed into named artefacts, transformations, contracts, modules and traces.
The long-term aim is to move LLM-assisted work from isolated prompting toward cognitive programming: reusable, inspectable, model-independent programs whose primary objects are ideas, documents, plans, requirements, reviews, concepts and transformations between them.
The Thought Tree Framework is conceptually developed and partially prototyped, but it is not a finished production system.
The original author is stepping back due to family commitments, caring responsibilities, work commitments, burnout recovery and the fact that the project now needs skills, time and expertise beyond one person.
This repository is being released openly so that others can:
compare it with existing agent and workflow frameworks;
or take the project in a better direction.
The goal of this release is not to claim the framework is complete.
The goal is to prevent the idea from being trapped with one person.
LLMs are powerful at working with ambiguous human concepts: documents, requirements, notes, stories, plans, reviews, policies, designs and decisions.
But most current LLM use still happens through:
workflow automations without semantic validation.
These approaches can be useful, but they often lack qualities expected from production systems:
inspectable intermediate results;
separation between workflow definition and model provider;
and clear provenance for final outputs.
For small tasks, a prompt may be enough.
For larger cognitive work, the process matters.
A Thought Tree program describes cognitive work as transformations over named artefacts.
Inputs
↓
Operations
↓
Intermediate Artefacts
↓
Review / Validation
↓
Final Outputs
↓
Execution Trace
For example, instead of asking an LLM:
"Write a technical design document from these notes."
a Thought Tree module might define:
design_notes
↓
CreateSourceDigest
↓
source_digest
↓
ExtractRequirements
↓
requirements_register
↓
DesignSystemDecomposition
↓
system_decomposition
↓
DraftSections
↓
draft_tdd_sections
↓
AssembleDraft
↓
draft_tdd
↓
ReviewDraft
↓
review_and_correction_plan
↓
ReviseFinalDocument
↓
technical_design_document
Each step produces an artefact that can be inspected, validated, replaced, reused or traced.
##Data Unit
A Data Unit is a discrete artefact used or produced by a workflow.
In the current TTML draft, Data Units are often represented as files, but conceptually they may be any addressable artefact.
An Operation transforms input Data Units into output Data Units.
source_notes → CreateSourceDigest → source_digest
draft_chapter + feedback → ReviseChapter → revised_chapter
sections → ConcatenateFiles → compiled_document
Operations may be executed by:
A Module is a reusable cognitive program.
optional contracts and validation expectations.
A Module may run as a standalone workflow or be called by another Module.
A Cognitive Engine is the compiler and runtime for Thought Tree programs.
loading a Thought Tree definition;
resolving variables, inputs, references, collections and iterators;
compiling the workflow into an executable transformation graph;
detecting dependency errors and output collisions;
invoking LLMs, functions, tools, submodules and human review steps;
storing intermediate artefacts;
The LLM is not the whole system.
The Cognitive Engine coordinates both.
Thought Tree Markup Language, or TTML, is the current draft XML-based source format for Thought Tree Modules.
A minimal TTML-style workflow might look like:
< TTML version = " 0.12.0 " >
< Project
id = " ArticleSummaryModule "
desc = " Summarise an article through draft, review and revision. " />
< Inputs >
< File id = " source_article " folder = " /inputs " extension = " txt " />
</ Inputs >
< Operations >
< Operation
id = " DraftSummary "
type = " TextCompletion "
desc = " Create a concise summary of the source article. " >
< FileRef id = " source_article " />
< Output >
< File id = " draft_summary " extension = " txt " />
</ Output >
</ Operation >
< Operation
id = " ReviewSummary "
type = " TextCompletion "
desc = " Review the draft summary against the source article. Identify omissions, inaccuracies and unsupported claims. " >
< FileRef id = " source_article " />
< FileRef id = " draft_summary " />
< Output >
< File id = " summary_review " extension = " txt " />
</ Output >
</ Operation >
< Operation
id = " ReviseSummary "
type = " TextCompletion "
desc = " Revise the summary using the review. Produce the final summary only. " >
< FileRef id = " source_article " />
< FileRef id = " draft_summary " />
< FileRef id = " summary_review " />
< Output >
< File id = " final_summary " extension = " txt " />
</ Output >
</ Operation >
</ Operations >
< Output >
< FileRef id = " final_summary " />
</ Output >
</ TTML >
TTML is not the entire framework.
It is one source representation of the underlying Thought Tree Program Model.
Future source formats could include YAML, JSON, visual graphs or higher-level authoring tools.
How Thought Tree Differs from Other Approaches
A prompt asks a model for an answer.
A Thought Tree defines the process by which an answer should be produced, reviewed, revised and traced.
Prompt chains connect model calls.
and model-independent execution.
Compared with autonomous agents
Agents can be flexible, but may be difficult to predict, debug or audit.
Thought Tree favours explicit structure:
defined process
→ inspectable plan
→ controlled execution
→ preserved artefacts
→ traceable output
Dynamic planning is still possible, but generated plans or submodules should be validated and recorded.
Compared with traditional workflow tools
Traditional workflow tools are good at deterministic processes.
Thought Tree is designed for hybrid cognitive work:
LLMs for semantic transformation;
code for deterministic transformation;
tools for external capabilities;
humans for judgement and approval;
a Cognitive Engine for orchestration and traceability.
Thought Tree is intended for complex cognitive production tasks where quality, traceability and iteration matter.
transforming scattered notes into formal documentation;
generating technical design documents;
recovering legacy project knowledge;
producing compliance or audit preparation packs;
generating and reviewing creative media;
producing game content and lore;
drafting and revising long-form fiction;
extracting requirements from source documents;
producing recurring research reports;
generating, validating and improving other Thought Tree Modules.
This handoff package includes or is intended to include:
the full Thought Tree Framework explainer;
draft TTML concepts and examples;
the Thought Tree Program Model;
Cognitive Engine architecture notes;
Early proof of concept
A C#/Unity prototype demonstrated that an LLM workflow could pass text files between operations as variables.
This prototype did not use TTML. It used a hardcoded array of prompts.
It successfully planned, drafted and reviewed a 50,000-word novel from a 500-word user-submitted description.
Partial Cognitive Engine prototype
A later C#/Unity project began implementing a Cognitive Engine.
It can connect to:
It supports basic TextCompletion execution.
It does not yet include a stable TTML importer, full graph compilation, contracts, validation gates or a production workspace/trace system.
The framework still needs substantial work.
iterator and collection resolution;
security and governance model;
This is a foundation, not a finished product.
Suggested Next Steps for Contributors
Useful contribution areas include:
refine the Thought Tree Program Model;
define execution trace schema;
define semantic contract schema.
implement a minimal CLI Cognitive Engine;
resolve Data Units and FileRefs;
execute TextCompletion operations;
execute deterministic functions;
preserve intermediate artefacts;
create larger workflow examples;
create examples for technical documentation, creative writing, compliance, research and module improvement.
compare Thought Tree with LangGraph, AutoGen, CrewAI, Semantic Kernel, workflow engines and agent systems;
identify where the framework overlaps, differs or can interoperate.
Minimal Reference Engine Target
A minimal Cognitive Engine should be able to:
execute operations in document order;
support registered deterministic functions;
A more advanced engine could add:
automated improvement workflows.
thoughttree-framework/
README.md
HANDOFF.md
STATUS.md
ROADMAP.md
CONTRIBUTING.md
GOVERNANCE.md
LICENSE.md
docs/
ONE_PAGE_OVERVIEW.md
WHY_THIS_MATTERS.md
ARCHITECTURE.md
PROGRAM_MODEL.md
COGNITIVE_ENGINE.md
EXECUTION_SEMANTICS.md
SEMANTIC_CONTRACTS.md
AUTHORING_GUIDE.md
ThoughtTreeFramework.pdf
spec/
TTML_DRAFT.md
TTMLSchema.xsd
EXECUTION_TRACE_SCHEMA_DRAFT.json
CONTRACT_SCHEMA_DRAFT.json
examples/
article-summary-review/
novel-generation-pipeline/
video-game-tdd/
compliance-gap-analysis/
module-improvement/
prototypes/
unity-novel-poc/
unity-cognitive-engine/
This project may be relevant to:
AI

[truncated]
