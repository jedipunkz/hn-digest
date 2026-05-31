---
source: "https://github.com/peymanvahidi/awesome-claude-dynamic-workflows"
hn_url: "https://news.ycombinator.com/item?id=48336168"
title: "Workflow patterns behind Claude Dynamic Workflows"
article_title: "GitHub - peymanvahidi/awesome-claude-dynamic-workflows · GitHub"
author: "peymanvahidi"
captured_at: "2026-05-31T00:30:10Z"
capture_tool: "hn-digest"
hn_id: 48336168
score: 1
comments: 0
posted_at: "2026-05-30T13:50:17Z"
tags:
  - hacker-news
  - translated
---

# Workflow patterns behind Claude Dynamic Workflows

- HN: [48336168](https://news.ycombinator.com/item?id=48336168)
- Source: [github.com](https://github.com/peymanvahidi/awesome-claude-dynamic-workflows)
- Score: 1
- Comments: 0
- Posted: 2026-05-30T13:50:17Z

## Translation

タイトル: Claude Dynamic Workflow の背後にあるワークフロー パターン
記事のタイトル: GitHub - peymanvahidi/awesome-claude-dynamic-workflows · GitHub
説明: GitHub でアカウントを作成して、peymanvahidi/awesome-claude-dynamic-workflows の開発に貢献します。

記事本文:
GitHub - peymanvahidi/awesome-claude-dynamic-workflows · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ペイマンバヒディ
/
素晴らしい-クロード-ダイナミック-ワークフロー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

イオンオプション
コード
peymanvahidi/awesome-claude-dynamic-workflows
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット ダイナミックワークフロースキル ダイナミックワークフロースキル ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
素晴らしい Claude ダイナミック ワークフロー
コミュニティ README では、dynamic-workflows-skill/SKILL.md で明らかになった主要なアイデア、つまり動的ワークフローとは何か、いつ使用されるか、ランタイム プリミティブがどのように機能するか、およびどのオーケストレーション パターンが最も重要であるかを説明しています。
このリポジトリは非公式であり、Claude Dynamic Workflows スキルの教育的詳細を目的としています。
目標は、資料をざっと見て、理解し、議論しやすくすることです。
Dynamic-workflows-skill/SKILL.md — このリポジトリに収集されたワークフロー スキル ファイル。
README.md — そのスキルの最も重要な部分についてのガイド付き説明。
このスキルは、制御された決定的な方法で多くのサブエージェントを調整するためのワークフロー ランタイムを記述します。すべてに対して単一のモデル呼び出しを使用するのではなく、ワークフローを、分解、検証、合成、およびスケーリングのための実行可能なオーケストレーション ロジックとして提示します。
実際には、このシステムは、調査、レビュー、移行、監査、および 1 つのコンテキスト ウィンドウまたは 1 つのエージェントでは不十分なその他のジョブなどのタスク向けに設計されているようです。
このスキルは長いですが、7 つの概念レイヤーに分割すると、はるかに理解しやすくなります。
目的とオプトイン ルール - ワークフローが存在する理由と、ワークフローがいつ許可されるか。
Ultracode モード — 徹底的なオーケストレーションが有効になっているときにシステムがどのように変化するか。
ワークフロー ファイルの構造 — メタとフェーズを使用してスクリプトを宣言する方法。
ランタイム プリミティブ — Agent() 、 Pipeline() 、およびParallel() などのコア関数。
実行ルール —

同時実行の制限、障壁、引数、予算、およびネストされたワークフロー。
パターン — 敵対的検証やループ・アンティル・ドライなどの再利用可能なオーケストレーション形状。
再開セマンティクス — 中断後にキャッシュされたプレフィックスから実行を続行する方法。
冒頭のセクションでは、決定論的なマルチエージェント オーケストレーションのためのツールとしてワークフローを組み立てます。これらは、並行分解、段階的な検証、または広い作業面にわたる広範囲のカバーを必要とするタスクを対象としています。
主要なテーマは明示的なオプトインです。このスキルは、ユーザーがそのレベルのオーケストレーションを明確に要求した場合、ワークフロー固有のコマンドを呼び出した場合、または上位レベルのモードがすでにそれを承認している場合にのみ、ワークフローを実行する必要があると繰り返し述べています。
これは、ワークフローが高価であると見なされているため重要です。ワークフローは多くのサブエージェントにファンアウトし、大量のトークンを消費する可能性があります。したがって、システムはこれらをすべてのタスクに対するデフォルトの回答ではなく、強力なモードとして扱います。
このスキルの最も興味深い部分の 1 つは、 Ultracode の説明です。このモードでは、ワークフロー オーケストレーションが実質的なタスクのデフォルトとなり、トークン コストは徹底性よりも二次的なものとして扱われます。
このスキルは、Ultracode が単に「より深く考える」だけではないことを示唆しています。これは、特に理解 -> 設計 -> 実装 -> レビューなどのより困難なタスクの場合、デフォルトで複数ステップのワークフローの作成と実行を許可する永続的な権限モデルに近いように見えます。
これは重要な概念の変化です。システムは、単一のアシスタントが順番に応答するだけではなく、プログラム可能なランタイムに似てきます。
このスキルは、各ワークフロー スクリプトがリテラルのexport const meta = { ... } ブロックで始まることを示しています。このメタデータは、ワークフローに名前、説明、およびオプションのフェーズ定義を与えます。
重要な実装の詳細は、メタ オブジェクトが純粋な li でなければならないということです。

テラル。つまり、変数、関数呼び出し、またはテンプレート式から動的に組み立てるべきではありません。
フェーズシステムも重要です。 meta.phases のフェーズ タイトルは、進行状況 UI が関連する作業を正しくグループ化できるように、ランタイムの Phase() 呼び出しと正確に一致することが期待されます。
このスキルは、オーケストレーション モデルを定義するランタイム プリミティブの小さなセットを中心に展開します。
Agent() は、サブエージェントを生成するための基本単位であるようです。スキーマがアタッチされている場合は、フリー テキストまたは構造化出力のいずれかを返すことができます。
スキルに示されているオプションは、ラベル、フェーズの割り当て、モデルのオーバーライド、ワークツリーの分離、さらにはカスタム サブエージェント タイプなど、かなり豊富なランタイムを暗示しています。
Pipeline() はおそらく、このドキュメントの中で最も重要なプリミティブです。スキルは、それをマルチステージ作業のデフォルトとして明示的に構成します。
その主なアイデアは、他のすべてのアイテムが前のステージを完了するのを待たずに、各アイテムが独立してステージを通過するということです。つまり、別のブランチのレビュー中に、あるブランチの検証を開始できるということです。
Parallel() はまったく異なる方法で説明されています。それは、 バリア です。実行を続行するには、バッチ内のすべてのタスクが完了する必要があります。
このスキルは、グローバル重複排除や項目間の比較など、次のステップで以前の結果の完全なセットが実際に必要な場合にのみこれを使用する必要があることを強調しています。
Phase() は作業を UI の進捗グループに編成するようですが、log() はより高レベルのステータス更新を発行します。ランタイムは単に作業を実行するだけではないため、これらのプリミティブは重要です。また、その作品の構造をユーザーに公開することになります。
args 、 Budget 、および workflow()
このスキルには、次の 3 つのオーケストレーション レベルのコントロールも表示されます。
実際の JSON 値を使用してワークフローをパラメータ化するための args。
トークンを意識したスケーリングとハードシーリングのための予算。
ワークフロー()

別のワークフローをサブステップとして呼び出します。
これらを組み合わせると、システムはプロンプト テンプレートというよりもプログラム可能な実行環境のように感じられます。
5. 実行ルールと設計思想
スキルの主な設計原則は単純で、デフォルトは Pipeline() です。障壁は通常ではなく例外として扱われます。
このテキストでは、後の段階で以前の結果をすべて一度に必要とする場合にのみ、バリアが正当化されるかどうかを明確にテストします。例としては、高価な検証の前にすべての結果を重複排除すること、合計結果がゼロの場合の早期終了、または 1 つの項目を残りの項目と明示的に比較するプロンプトが含​​まれます。
このスキルは、不要なバリアがなぜ悪いのかも説明します。遅いタスクを待機している間、速いタスクを強制的にアイドル状態にしておくことで、実時間を無駄にします。
もう 1 つの重要なルールは、スケールの制限です。同時実行のagent()呼び出しはワークフローごとに制限され、エージェントの総数はワークフローの存続期間全体で制限されます。これは、ユーザーが非常に大きな項目リストを渡した場合でも、ランタイムが制御不能な爆発を防ぐように設計されていることを示唆しています。
このスキルの最も興味深い部分はパターン カタログです。これらのパターンは、オーケストレーション システムの背後にある考え方を明らかにするため、役立ちます。
ワークフローは発見を額面通りに受け入れるのではなく、独立した懐疑論者を生み出し、彼らに反論するよう求めます。主張は、十分な検証者が反論できなかった場合にのみ存続します。
これは、検証がオプションのクリーンアップ パスではなく、第一級の懸念事項として扱われていることを示す強力な兆候です。
このスキルは、冗長検証と多様な検証を区別します。 3 人の同一の批評家に同じことを尋ねるのではなく、正確性、セキュリティ、パフォーマンス、再現性などの異なるレンズを使用することを推奨しています。
これは微妙ですが重要な考え方です。評価者の多様性により、異なる失敗が検出される可能性があります。

同一のチェックを繰り返すモードは失敗します。
もう一つのパターンは審査員団です。このワークフローは、いくつかの独立した試行を生成し、それらを採点し、次点者から良いアイデアを借用しながら、勝者から合成します。
これは、ワークフローがバグを見つけるためだけのものではないことを示唆しています。また、明白な最初の試みが 1 つも存在しないソリューション スペースを探索するためにも使用できます。
このパターンは、無制限の検出タスク用です。システムは、一定のカウント後に停止するのではなく、いくつかの連続したラウンドで新しいものが何も生成されないまで実行を続けます。
これは、実際の検出結果の総数が事前に不明な監査、バグ ハント、および調査のスイープに役立ちます。
マルチモーダルなスイープと完全性の批評家
このスキルでは、マルチモーダル スイープと完全性の批評パターンについても説明します。 1 つ目は、さまざまな視点または検索角度にわたって検索を分散します。 2 番目の質問では、未読のソース、未検証の主張、スキップされたモダリティなど、何が見逃されたかを尋ねます。
これらのパターンを総合すると、単一の検索パスを信頼せずにカバレッジを最大化しようとするアーキテクチャが示されています。
最後の主要なセクションでは、再開可能性について説明します。ワークフローの実行は runId を使用して再開できるようで、前の Agent() 呼び出しの変更されていないプレフィックスはキャッシュされた結果をすぐに返すことができます。
これは、反復プロセスとしてのワークフロー開発を示唆しているため、特に興味深いです。必ずしもグラフ全体を再実行するとは限りません。安定したプレフィックスを保持し、編集した末尾のみを再実行できます。
このスキルでは、このレンズを通していくつかの珍しい制限についても説明します。Date.now() や Math.random() などの API は、決定的な再開動作を壊すためブロックされます。
最大のポイントは、最新のエージェント システムが「1 つのモデル、1 つの応答」から、明示的なオーケストレーション フレームワークに向かって進化している可能性があるということです。

分解、検証、予算編成、履歴書のセマンティクスが組み込まれた rks。
Dynamic-workflows-skill/SKILL.md を初めて開く場合は、これが最速のパスです。
ワークフローの目的とオプトインに関する冒頭のセクションを読んでください。
メタ ブロックとランタイム プリミティブの定義をざっと読んでください。
Pipeline() とParallel() に注意深く注目してください。
品質パターンのセクションをお読みください。
「履歴書」セクションで終了します。
このシーケンスにより、認知的過負荷が最小限に抑えられた最良のメンタル モデルが得られます。
この README は説明ガイドであり、正式な仕様ではありません。
このリポジトリは、コミュニティのメモ作成および分析プロジェクトとして扱うのが最適です。その価値は、スキルを公式の公的標準として扱うことではなく、整理、解釈、例によってもたらされます。
読み込み中にエラーが発生しました。このページをリロードしてください。
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

Contribute to peymanvahidi/awesome-claude-dynamic-workflows development by creating an account on GitHub.

GitHub - peymanvahidi/awesome-claude-dynamic-workflows · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
peymanvahidi
/
awesome-claude-dynamic-workflows
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
peymanvahidi/awesome-claude-dynamic-workflows
master Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit dynamic-workflows-skill dynamic-workflows-skill LICENSE LICENSE README.md README.md View all files Repository files navigation
Awesome Claude Dynamic Workflows
A community README that explains the key ideas surfaced in dynamic-workflows-skill/SKILL.md : what dynamic workflows are, when they are used, how their runtime primitives work, and which orchestration patterns appear most important.
This repository is unofficial and intended as an educational breakdown of a Claude Dynamic Workflows skill.
The goal is to make the material easier to scan, understand, and discuss.
dynamic-workflows-skill/SKILL.md — the workflow skill file collected in this repository.
README.md — a guided explanation of the most important parts of that skill.
The skill describes a workflow runtime for orchestrating many subagents in a controlled, deterministic way. Rather than using a single model call for everything, it presents workflows as executable orchestration logic for decomposition, verification, synthesis, and scaling.
In practical terms, the system appears designed for tasks such as research, review, migration, audits, and other jobs where one context window or one agent is not enough.
The skill is long, but it becomes much easier to understand if you split it into seven conceptual layers:
Purpose and opt-in rules — why workflows exist and when they are allowed.
Ultracode mode — how the system changes when exhaustive orchestration is enabled.
Workflow file structure — how a script is declared with meta and phases.
Runtime primitives — the core functions such as agent() , pipeline() , and parallel() .
Execution rules — concurrency limits, barriers, arguments, budgets, and nested workflows.
Patterns — reusable orchestration shapes like adversarial verification or loop-until-dry.
Resume semantics — how runs can continue from cached prefixes after interruption.
The opening section frames workflows as a tool for deterministic multi-agent orchestration. They are meant for tasks that need parallel decomposition, staged verification, or broad coverage across a large work surface.
A major theme is explicit opt-in . The skill repeatedly says workflows should only run when the user clearly requested that level of orchestration, invoked a workflow-specific command, or when a higher-level mode already authorizes it.
This matters because workflows are presented as expensive: they can fan out into many subagents and consume substantial tokens. The system therefore treats them as a powerful mode, not the default answer to every task.
One of the most interesting parts of the skill is the description of Ultracode . In this mode, workflow orchestration becomes the default for substantive tasks, and token cost is treated as secondary to thoroughness.
The skill suggests that Ultracode is not just “think harder.” It appears closer to a standing permission model that allows multi-step workflow authoring and execution by default, especially for harder tasks like understand -> design -> implement -> review.
That is an important conceptual shift: the system starts to resemble a programmable runtime, not just a single assistant responding turn by turn.
The skill shows that each workflow script begins with a literal export const meta = { ... } block. This metadata gives the workflow a name, description, and optional phase definitions.
A key implementation detail is that the meta object must be a pure literal. In other words, it should not be dynamically assembled from variables, function calls, or template expressions.
The phase system is also important. Phase titles in meta.phases are expected to match the runtime phase() calls exactly so the progress UI can group related work correctly.
The skill revolves around a small set of runtime primitives that define the orchestration model.
agent() appears to be the basic unit for spawning a subagent. It can return either free text or structured output when a schema is attached.
The options shown in the skill imply a fairly rich runtime: labels, phase assignment, model override, worktree isolation, and even custom subagent types.
pipeline() is arguably the most important primitive in the document. The skill explicitly frames it as the default for multi-stage work.
Its main idea is that each item moves through stages independently, without waiting for all other items to finish the previous stage. That means verification for one branch can begin while another branch is still reviewing.
parallel() is described very differently: it is a barrier . All tasks in the batch must complete before execution continues.
The skill emphasizes that this should only be used when the next step truly needs the full set of prior results together, such as global deduplication or cross-item comparison.
phase() seems to organize work into progress groups for the UI, while log() emits higher-level status updates. These primitives matter because the runtime is not just executing work; it is also exposing the structure of that work to the user.
args , budget , and workflow()
The skill also shows three orchestration-level controls:
args for parameterizing workflows with real JSON values.
budget for token-aware scaling and hard ceilings.
workflow() for invoking another workflow as a sub-step.
Together, these make the system feel more like a programmable execution environment than a prompt template.
5. Execution rules and design philosophy
A major design principle in the skill is simple: default to pipeline() . Barriers are treated as exceptional, not normal.
The text gives a clear test for when a barrier is justified: only when a later stage truly needs all previous results at once. Examples include deduplicating all findings before expensive verification, early exit when total results are zero, or prompts that explicitly compare one item to the rest.
The skill also explains why unnecessary barriers are bad. They waste wall-clock time by forcing fast tasks to sit idle while waiting for slower ones.
Another important rule is bounded scale. Concurrent agent() calls are capped per workflow, and total agent count is capped across the workflow lifetime. That suggests the runtime is designed to prevent uncontrolled explosion even when users pass in very large item lists.
The most interesting part of the skill is the pattern catalog. These patterns are useful because they reveal the mindset behind the orchestration system.
Instead of accepting a finding at face value, the workflow spawns independent skeptics and asks them to refute it. A claim survives only if enough verifiers fail to refute it.
This is a strong sign that verification is treated as a first-class concern, not an optional cleanup pass.
The skill distinguishes between redundant verification and diverse verification. Rather than asking three identical critics the same thing, it recommends using different lenses such as correctness, security, performance, or reproducibility.
This is a subtle but important idea: diversity of evaluators may catch different failure modes that repeated identical checks would miss.
Another pattern is the judge panel. The workflow generates several independent attempts, scores them, and synthesizes from the winner while borrowing good ideas from the runners-up.
This suggests workflows are not just for finding bugs; they can also be used to explore solution spaces where there is no single obvious first attempt.
This pattern is for open-ended discovery tasks. Instead of stopping after a fixed count, the system keeps running until several consecutive rounds produce nothing new.
That is useful for audits, bug hunts, and research sweeps where the total number of real findings is unknown in advance.
Multi-modal sweep and completeness critic
The skill also describes multi-modal sweep and completeness critic patterns. The first spreads search across different perspectives or search angles; the second asks what was missed, such as unread sources, unverified claims, or skipped modalities.
Together, these patterns show an architecture that tries to maximize coverage without trusting any single search path.
The last major section covers resumability. A workflow run can apparently be resumed using a runId , and unchanged prefixes of prior agent() calls can return cached results immediately.
This is particularly interesting because it hints at workflow development as an iterative process. You do not always rerun the entire graph; you can keep the stable prefix and only re-execute the edited tail.
The skill also explains some unusual restrictions through this lens: APIs like Date.now() and Math.random() are blocked because they would break deterministic resume behavior.
The biggest takeaway is that modern agent systems may be evolving away from “one model, one response” and toward explicit orchestration frameworks with decomposition, verification, budgeting, and resume semantics built in.
If you are opening dynamic-workflows-skill/SKILL.md for the first time, this is the fastest path:
Read the opening section on workflow purpose and opt-in.
Skim the meta block and runtime primitive definitions.
Focus carefully on pipeline() vs parallel() .
Read the quality patterns section.
Finish with the Resume section.
That sequence gives the best mental model with the least cognitive overload.
This README is an explanatory guide, not an official specification.
This repository is best treated as a community note-taking and analysis project. Its value comes from organization, interpretation, and examples rather than from treating the skill as an official public standard.
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
