---
source: "https://machinelearningmastery.com/ai-agent-tool-design-what-works-and-what-doesnt/"
hn_url: "https://news.ycombinator.com/item?id=48540345"
title: "AI Agent Tool Design: What Works and What Doesn't"
article_title: "AI Agent Tool Design: What Works and What Doesn't"
author: "eigenBasis"
captured_at: "2026-06-15T14:16:10Z"
capture_tool: "hn-digest"
hn_id: 48540345
score: 1
comments: 0
posted_at: "2026-06-15T12:35:12Z"
tags:
  - hacker-news
  - translated
---

# AI Agent Tool Design: What Works and What Doesn't

- HN: [48540345](https://news.ycombinator.com/item?id=48540345)
- Source: [machinelearningmastery.com](https://machinelearningmastery.com/ai-agent-tool-design-what-works-and-what-doesnt/)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T12:35:12Z

## Translation

タイトル: AI エージェント ツールの設計: 何が機能し、何が機能しないのか
説明: この記事では、AI エージェント ツールが適切に機能する理由と、失敗の原因となる一般的な設計ミスについて説明します。ツールの設計がエージェントのタスクを正確かつ一貫して完了する能力にどのように影響するかを学びます。

記事本文:
AI エージェント ツールの設計: 何が機能し、何が機能しないのか
ナビゲーション
開発者を機械学習で優れたものにする
開発者を機械学習で優れたものにする
コード アルゴリズム 機械学習アルゴリズムを最初から実装します。
ディープラーニング (keras) ディープラーニング
時系列予測のためのニューラルネット時系列ディープラーニング
LSTM 長短期記憶ネットワーク
開発者を機械学習で優れたものにする
開発者を機械学習で優れたものにする
コード アルゴリズム 機械学習アルゴリズムを最初から実装します。
ディープラーニング (keras) ディープラーニング
時系列予測のためのニューラルネット時系列ディープラーニング
LSTM 長短期記憶ネットワーク
AI エージェント ツールの設計: 何が機能し、何が機能しないのか
シェアする
ポスト
シェアする
この記事では、ほとんどの AI エージェントの失敗の根本原因がモデルの機能ではなくツールの設計であることと、それを修正するために適用できる具体的な設計パターンについて学びます。
単一責任ツール、厳密なスキーマ、構造化されたエラーリターンなど、エージェントの信頼性を向上させるツール設計プラクティス。
フィルターされていない API の公開、サイレントな部分的な成功、実際のワークロードを中断するツール名の重複などの一般的な障害モード。
ツール境界での幻覚や信頼性の低い動作を軽減するスキーマとエラー処理パターン。
AI エージェント ツールの設計: 何が機能し、何が機能しないのか
AI エージェントの失敗のほとんどは、間違ったツールを選択したり、間違った引数を渡したり、エラーを誤って処理したりするなど、モデルの間違いのように見えます。しかし実際には、モデルは通常、与えられたインターフェイスで動作します。多くの場合、根本的な問題はツールの設計自体にあります。
モデルは、ツール インターフェイスを通じて公開される情報 (ツール名、その説明、パラメーター スキーマ、およびパラメーターの説明) からのみ推論できます。その詳細はね

モデルがどのように意図を解釈し、アクションを計画し、タスクを実行するかを説明します。ツールの設計が不明確、不完全、または構造が緩い場合、障害は偶然ではなく予測可能になります。
あいまいな名前、あいまいな命令、一貫性のないスキーマ、脆弱なパラメータ定義、貧弱なエラー処理などの問題はすべて、失敗の可能性を高めます。モデルを強化すると、一部の間違いは軽減できますが、欠陥のあるインターフェイスを確実に補うことはできません。この記事の内容は次のとおりです。
信頼性を向上させるツール設計の実践
デモでは問題ないように見えても、実際のワークロードでは壊れる障害モード
ツール境界での幻覚を軽減するスキーマとエラー設計
設計が失敗する理由を理解することは、それを何に置き換えるべきかを知ることと同じくらい重要であるため、各パターンは失敗した対応するパターンとペアになります。
AI エージェント ツールの設計で機能するもの
1. 1 つのツール、1 つの責任
ほとんどのエージェント システムでは、ツールは単一の明確な操作を表す必要があります。 1 つのツールがアクション パラメーターを通じて複数の動作を処理する場合、モデルは実際のタスクを解決する前に、まずどのモードを呼び出すかを判断する必要があります。
マルチアクション ツールと専用の単一目的ツールを比較すると、その違いがより明確になります。
⚠️ 注 : これは、普遍的なルールではなく、便利なデフォルトです。シェル、ファイルシステム、ブラウザ、カレンダー ツールなどの一部のドメインでは、アクション空間自体が基礎となる抽象化の一部であるため、制約付きマルチアクション インターフェイスの恩恵を受ける可能性があります。
2. 無効な状態を不可能にするスキーマ
ツール呼び出しエージェント では、モデルはスキーマから推論してツール呼び出し引数を構築します。
緩いスキーマとは、モデルが制約を推測することを意味します。
厳密なスキーマはこれらの制約をエンコードするため、推測する必要はありません。
列挙型はフィールドに特に役立ちます

妥当ではあるが無効な出力のクラスが排除されるため、有効な値の小さなセットが使用されます。検証の失敗は、下流の不可解なエラーとしてではなく、ツールの境界で表面化します。
3. 目的だけでなく範囲を定義する説明
ツールの説明はモデル向けのドキュメントです。彼らは 2 つのことを行う必要があります。1 つはツールをいつ使用するかを説明すること、もう 1 つはツールを使用しない場合を説明することです。ほとんどの説明は最初の説明のみを行います。
曖昧さを明確にしないと、モデルはツール名のみからスコープを推測します。これは多くの場合、大規模な選択エラーの信頼できる原因となります。優れたツール定義には、使用手順だけでなく、他のツールとの明確な境界が含まれています。
4. 構造化されたアクション可能なエラーが返される
ツールが失敗すると、モデルはエラーを読み取り、次に何をすべきかを決定します。未処理の例外またはスタック トレースにより、ノイズ主導のフォローアップ動作が発生します。構造化エラーは、モデルに分岐のきっかけを与えます。
構造化エラーは、何が失敗したかを報告するだけでなく、エージェントが次に何をすべきかを決定するのにも役立ちます。適切なエラー形式では、再試行動作が明確になり、モデルに明確な回復パスが与えられます。
回復可能フラグと示唆されたアクションフィールドは、エージェントの動作を変更するものです。これらがないと、モデルは再試行不可能なエラーを再試行するか、回復可能なエラーを破棄します。
5. 冪等な状態変更操作
状態を変更するすべてのツール (レコードの作成、メッセージの送信、資金の送金) は、2 回呼び出しても安全でなければなりません。実際には、エージェントが再試行し、ネットワークに障害が発生し、最初のコールの確認が届かなかったため、LLM ループが 2 番目のコールを発行する可能性があります。
重複した副作用を防ぐ簡単な方法は、すべての書き込み操作に冪等性キーを要求することです。
冪等性が保証されていないと、一時的な障害が簡単に重複したアクションに変わってしまう可能性があります。
AI エージェント ツールの設計で機能しないもの
1. 薄いラッパー

フィルタリングされていない API について
エージェントに REST API を指定し、それをツールとして表示することは、最も一般的な近道であり、実稼働エラーの最も一般的な原因です。開発者向けに構築された API は、多くの場合、エージェントが実際に必要とするよりもはるかに詳細な情報を公開します。回答には、関連するフィールドがほんの一部であっても、数百ものフィールドが詰め込まれています。これらはページネーションに依存し、文脈上の意味がほとんどない不透明な内部 ID を使用し、解釈するには深いドメイン知識を必要とするエラー コードを返します。
専用のラッパーは内部でページネーションを処理し、エージェントが必要とするフィールドのみを投影し、API エラーを上で説明した構造化された ToolError 形式にマップします。エージェントは API パスを構築したり、ページを管理したりすることはありません。推論できる型付きオブジェクトを受け取ります。
とはいえ、過剰な包装も有害である可能性があります。すべてのエンドポイントが共有構造を持たない個別の狭義のツールになると、ツールの表面が断片化し、モデルの操作が困難になる可能性があります。目標は、最大限の抽象化ではなく、一貫したエージェントフレンドリーな抽象化レイヤーです。
2. すべてのツールをすべてのコンテキストにロードする
ツールカタログが増えると精度が低下します。長いコンテキストにわたるツール呼び出しのパフォーマンスに関する 2025 年の調査である LongFuncEval では、128K コンテキスト ウィンドウを持つモデルであっても、ツール カタログ サイズが増加するとパフォーマンスが大幅に低下することがわかりました。すべてのツールをすべてのシステム プロンプトにロードすると、タスク コンテンツが処理される前にトークン バジェットが消費されるため、これがさらに悪化します。
動的なツールの読み込みにより、両方の問題が解決されます。現在のステップに関連するツールを特定し、次のツールのみを含めます。
部分的な成功は、ツールが要求された作業の一部だけを完了しても、完全に成功したように見える応答を返す場合に問題になります。エージェントは、システムの不完全なビューまたは誤解を招くビューで実行を続行します。

テムの状態。
これは通常、ツールが内部エラーを抑制し、結果の成功部分のみを返す場合に発生します。
Partial_success フラグは、失敗した項目を再試行するか、部分的な結果をユーザーに表示するか、ワークフローを停止するかなど、モデルに分岐の基準を与えます。
4. 重複するツール名と説明
2 つのツールが同様のことを実行する場合、モデルは呼び出しごとにどちらを使用するかを検討します。この推論ではトークンが消費され、エラーが発生します。一般的な例としては次のようなものがあります。
同じ目的を持つ search_documents と find_documents
get_user と fetch_user_profile の違いが不明瞭
1 つの操作のための 3 つのツールとしての create_task 、 add_task 、および new_task
このような場合、名前を変更するだけでは解決できません。すべてのツールには、セット内の他のツールを参照せずに説明できる目的が必要です。 「X とは異なり、これは…」という説明が意味をなす必要がある場合、それは設計上の問題です。ツールのスプロール（範囲が重複するツールが多すぎる）は、企業展開におけるエージェントの動作が信頼できない原因になります。
5. 確認ゲートのない破壊行為
記録の削除、実際のユーザーへのメッセージ送信、金融取引の実行など、取り消しできないアクションを実行するツールには、即時に「よろしいですか?」というメッセージではなく、構造的な 2 段階の確認が必要です。段階的なアプローチにより、明示的な確認境界が導入され、偶発的または不正な実行のリスクが軽減されます。
最も安全なパターンは、ステージングを実行から分離し、2 つのステップの間に有効期間の短い確認トークンを要求することです。
確認ゲートのない破壊行為
⚠️ 注: ただし、多くのシステムでは、2 段階の安全フローだけでは十分ではないことがよくあります。ステージングと確認が使用される場合でも、有効期間が短く使い捨てのトークンなどの追加の保護手段が必要です。

厳密なセッション バインディングとリプレイ保護は、意図された安全境界をバイパスする可能性のあるトークンの再利用、漏洩、またはクロスセッション実行を防ぐために必要です。
AI エージェント ツールの設計上の決定の概要
各行は、AI エージェント ツールの設計における重要な決定を表しています。
AI エージェント用の効果的なツールを作成する - Anthropic の AI エージェントを使用することは、ツール設計の参考になります。
シェアする
ポスト
シェアする
このトピックの詳細
実際の機能スケーリング: 何が機能し、何が機能しないのか
エラー回復を備えたマルチツール Gemma 4 エージェントの構築
Weka で最初の実験を設計して実行する
畳み込みニューラル ネットワークの設計を理解する
安定拡散のインテリアデザイン（8日間ミニコース）
知っておくべき 7 つのエージェント AI 設計パターン
返信を残す 返信をキャンセルするにはここをクリックしてください。
メールアドレス（公開されません）（必須）
ようこそ！
私はジェイソン・ブラウンリー博士です
また、開発者が機械学習で結果を得るのを支援しています。
続きを読む
電子書籍カタログはここにあります
本当に良いものが見つかります。
Machine Learning Mastery は、人々がテクノロジーを理解できるよう支援することに重点を置いた大手デジタル メディア パブリッシャーである Guiding Tech Media の一部です。当社の使命とチームについて詳しくは、当社の企業 Web サイトをご覧ください。
© 2026 ガイディングテックメディア全著作権所有

## Original Extract

In this article, we explore what makes AI agent tools work well and the common design mistakes that cause failures. Learn how tool design affects an agent's ability to complete tasks accurately and consistently.

AI Agent Tool Design: What Works and What Doesn't
Navigation
Making developers awesome at machine learning
Making Developers Awesome at Machine Learning
Code Algorithms Implementing machine learning algorithms from scratch.
Deep Learning (keras) Deep Learning
Neural Net Time Series Deep Learning for Time Series Forecasting
LSTMs Long Short-Term Memory Networks
Making developers awesome at machine learning
Making Developers Awesome at Machine Learning
Code Algorithms Implementing machine learning algorithms from scratch.
Deep Learning (keras) Deep Learning
Neural Net Time Series Deep Learning for Time Series Forecasting
LSTMs Long Short-Term Memory Networks
AI Agent Tool Design: What Works and What Doesn’t
Share
Post
Share
In this article, you will learn how tool design — not model capability — is the root cause of most AI agent failures, and what concrete design patterns you can apply to fix it.
Tool design practices that improve agent reliability, including single-responsibility tools, tight schemas, and structured error returns.
Common failure modes such as unfiltered API exposure, silent partial success, and overlapping tool names that break real-world workloads.
Schema and error handling patterns that reduce hallucination and unreliable behavior at the tool boundary.
AI Agent Tool Design: What Works and What Doesn’t
Most AI agent failures look like model mistakes: choosing the wrong tool, passing bad arguments, or mishandling errors. But in practice, the model is usually working with the interface it was given. The underlying issue is often the tool design itself.
A model can only reason from the information exposed through the tool interface: the tool name, its description, the parameter schema, and the parameter descriptions. Those details shape how the model interprets intent, plans actions, and executes tasks. When the tool design is unclear, incomplete, or loosely structured, failures become predictable rather than accidental.
Problems like vague naming, ambiguous instructions, inconsistent schemas, weak parameter definitions, and poor error handling all increase the likelihood of failures. Stronger models can reduce some mistakes, but they cannot reliably compensate for a flawed interface. This article covers:
Tool design practices that improve reliability
Failure modes that look fine in demos but break under real workloads
Schema and error design that reduces hallucination at the tool boundary
Each pattern is paired with its failure counterpart, because understanding why a design fails is as important as knowing what to replace it with.
What Works in AI Agent Tool Design
1. One Tool, One Responsibility
In most agent systems, a tool should represent a single, clear operation . When one tool handles multiple behaviors through an action parameter, the model must first figure out which mode to invoke before it can solve the actual task.
The difference becomes clearer when comparing a multi-action tool against dedicated single-purpose tools:
⚠️ Note : This is a useful default rather than a universal rule. Some domains — such as shell, filesystem, browser, or calendar tools — may benefit from a constrained multi-action interface because the action space itself is part of the underlying abstraction.
2. Schemas That Make Invalid States Impossible
In tool-calling agents , the model constructs tool call arguments by reasoning from your schema.
A loose schema means the model guesses at constraints.
A tight schema encodes those constraints so no guessing is needed.
Enums are particularly useful for fields with a small set of valid values because they eliminate a class of plausible-but-invalid outputs. Validation failures surface at the tool boundary rather than as cryptic downstream errors.
3. Descriptions That Define Scope, Not Just Purpose
Tool descriptions are model-facing documentation. They need to do two things: explain when to use the tool, and explain when not to . Most descriptions only do the first.
Without the disambiguation, the model infers scope from the tool name alone, which is often a reliable source of selection errors at scale. A good tool definition includes clear boundaries from other tools, not just usage instructions.
4. Structured, Actionable Error Returns
When a tool fails, the model reads the error and decides what to do next. An unhandled exception or stack trace produces noise-driven follow-up behavior. A structured error gives the model something to branch on.
Structured errors should not only report what failed but also help the agent decide what to do next. A good error format makes retry behavior explicit and gives the model a clear recovery path:
The recoverable flag and suggested_action field are what change agent behavior. Without them, models retry non-retryable errors or abandon recoverable ones.
5. Idempotent State-Changing Operations
Every tool that mutates state — creates a record, sends a message, transfers funds — must be safe to call twice. In practice, agents retry, networks fail, and the LLM loop may issue a second call because confirmation of the first never arrived.
A simple way to prevent duplicate side effects is to require an idempotency key for every write operation:
Without idempotency guarantees, transient failures can easily turn into duplicate actions.
What Doesn’t Work in AI Agent Tool Design
1. Thin Wrappers Around Unfiltered APIs
Pointing an agent at a REST API and surfacing it as a tool is the most common shortcut and the most common source of production failures. APIs built for developers often expose far more detail than agents actually need . Responses come packed with hundreds of fields, even when only a handful are relevant. They rely on pagination, use opaque internal IDs with little contextual meaning, and return error codes that require deep domain knowledge to interpret.
A purpose-built wrapper handles pagination internally, projects only the fields the agent needs, and maps API errors to the structured ToolError format discussed above. The agent never constructs API paths or manages pages; it receives typed objects it can reason about.
That said, over-wrapping can also be harmful. If every endpoint becomes a separate, narrowly defined tool with no shared structure, the tool surface can become fragmented and harder for the model to navigate. The goal is not maximal abstraction, but a consistent, agent-friendly abstraction layer.
2. Loading All Tools Into Every Context
Accuracy degrades as the tool catalog grows. LongFuncEval , a 2025 study on tool-calling performance across long contexts, found performance drops substantially as the tool catalog size increased — even in models with 128K context windows. Loading every tool into every system prompt compounds this by consuming token budget before any task content is processed.
Dynamic tool loading addresses both problems. Determine which tools are relevant to the current step and include only those:
Partial success becomes a problem when a tool completes only part of the requested work but returns a response that looks fully successful. The agent continues execution with an incomplete or misleading view of the system state.
This usually happens when tools suppress internal failures and return only the successful portion of the result:
The partial_success flag gives the model something to branch on: retry the failed items, surface the partial result to the user, or halt the workflow.
4. Overlapping Tool Names and Descriptions
When two tools do similar things, the model reasons about which to use on every call. That reasoning costs tokens and introduces errors. Some common examples include:
search_documents and find_documents with identical purpose
get_user and fetch_user_profile with unclear differences
create_task , add_task , and new_task as three tools for one operation
In such cases, renaming alone isn’t the fix. Every tool needs a purpose that can be described without reference to other tools in the set. If a description requires “unlike X, this one…” to make sense, that’s a design problem. Tool sprawl — too many tools with overlapping scope — is a source of unreliable agent behavior in enterprise deployments.
5. Destructive Actions Without a Confirmation Gate
Any tool that takes an irreversible action — deleting records, messaging real users, executing financial transactions — needs a structural two-step confirmation, not an in-prompt “are you sure?” A staged approach introduces an explicit confirmation boundary that reduces the risk of accidental or unauthorized execution.
The safest pattern is to separate staging from execution and require a short-lived confirmation token between the two steps:
Destructive Actions Without a Confirmation Gate
⚠️ Note : Two-step safety flows, however, are often not sufficient on their own in many systems. Even when staging and confirmation are used, additional safeguards — such as short-lived, single-use tokens, strict session binding, and replay protection — are necessary to prevent token reuse, leakage, or cross-session execution that can bypass the intended safety boundary.
AI Agent Tool Design Decisions at a Glance
Every row represents a key decision in AI agent tool design:
Writing effective tools for AI agents — using AI agents from Anthropic is a useful reference on tool design.
Share
Post
Share
More On This Topic
Feature Scaling in Practice: What Works and What Doesn’t
Building a Multi-Tool Gemma 4 Agent with Error Recovery
Design and Run your First Experiment in Weka
Understanding the Design of a Convolutional Neural Network
Interior Design with Stable Diffusion (8-day mini-course)
7 Must-Know Agentic AI Design Patterns
Leave a Reply Click here to cancel reply.
Email (will not be published) (required)
Welcome!
I'm Jason Brownlee PhD
and I help developers get results with machine learning .
Read more
The EBook Catalog is where
you'll find the Really Good stuff.
Machine Learning Mastery is part of Guiding Tech Media, a leading digital media publisher focused on helping people figure out technology. Visit our corporate website to learn more about our mission and team.
© 2026 Guiding Tech Media All Rights Reserved
