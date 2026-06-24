---
source: "https://www.apimatic.io/blog/agentic-engineering-claude-code"
hn_url: "https://news.ycombinator.com/item?id=48663385"
title: "Vibe Coding to Agentic Engineering with Claude Code"
article_title: "From Vibe Coding to Agentic Engineering with Claude Code"
author: "m3h"
captured_at: "2026-06-24T18:34:34Z"
capture_tool: "hn-digest"
hn_id: 48663385
score: 1
comments: 0
posted_at: "2026-06-24T17:52:10Z"
tags:
  - hacker-news
  - translated
---

# Vibe Coding to Agentic Engineering with Claude Code

- HN: [48663385](https://news.ycombinator.com/item?id=48663385)
- Source: [www.apimatic.io](https://www.apimatic.io/blog/agentic-engineering-claude-code)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T17:52:10Z

## Translation

タイトル: クロード コードを使用したエージェント エンジニアリングへのバイブ コーディング
記事のタイトル: クロード コードによる Vibe コーディングからエージェント エンジニアリングへ
説明: 反復可能な 3 フェーズのクロード コード ワークフロー: 構造化プロンプト、計画モード、レビュー。セッションごとにコーディング エージェントから信頼性の高いコードを取得します。

記事本文:
Claude Code を使用した Vibe コーディングからエージェント エンジニアリングまで
2026 年 6 月 23 日
21 分で読めます
このページでは
プロンプトを作成し、コーディング エージェントに渡して待ちます。数分後、あなたは戻ってきます…何か。たぶんそれはコンパイルされます。しかし、それは要件の半分を無視したり、求めてもいない抽象化を発明したり、コードベースの残りの部分が依存しているパターンを静かに破ったりします。そこで、再度プロンプトを出します。そして再度プロンプトを出します。そして、何かを出荷する前にコンテキスト ウィンドウを焼きます。
このループに見覚えがある場合は、通常、問題はモデルではありません。推測に任せますが、Claude Code セッションは、あいまいな要件、不明確な範囲、欠落している受け入れ基準、および自身の作業を検証する方法がないなど、開いたままになっていたことに気づいていないギャップを埋めてくれます。
この修正は、1 回限りの賢いプロンプトではありません。これは、プロンプトを中心とした反復可能なプロセスです。このブログでは、私がコーディング タスクに取り組むときに必ず従う 3 フェーズのクロード コード ワークフローを紹介します。
この 3 段階のプロセスのおかげで、より多くのコーディング セッションが必要なものを実装できるようになり、頻繁に発生する長いセッションの実行時間、無駄なセッション、コンテキスト ウィンドウのオーバーフローも回避できるようになりました。このプロセスを採用すると、Vibe コーディングからエージェント エンジニアリングに移行します。これは、良い結果を期待することと、それをエンジニアリングすることの違いです。
Vibe コーディング vs. エージェント エンジニアリング
バイブ コーディングは、エージェントにプロンプトを出し、結果をざっと読み、それが正しいと思われる場合はそれを受け入れ、次に進むという、現在では一般的な手法です。問題は、コードが実際のコードベース内に存在しなければならない瞬間に現れます。エージェントは曖昧さを推測で埋め、エージェントは推測したものをすべて継承します。
Agentic Engineering はその逆の立場です。コーディング エージェントを有能だが文字通りの協力者として扱い、エンジニアリング ディスクを保管します。

明確な要件、計画、テスト、検証、レビューなど、数十年にわたってその地位を獲得してきたライン。エージェントはさらに多くの入力を行います。あなたは引き続きエンジニアリングの責任を負います。
一目でわかるコントラストは次のとおりです。
以下の 3 段階のプロセスは、日常の実践におけるエージェントティック エンジニアリングの様子を示しています。
構造化コードプロンプトから始めて意図を捉える
適切なコード プロンプトは、コーディング エージェントに、明確に定義されたスコープと検証基準を備えた単一のタスクを提供します。すべてのコーディング タスクの前に、コピーして入力する必要があるプロンプト テンプレートを次に示します。
[タスクを説明する 1 つの命令文。] [1 つまたは 2 つの文
なぜ...その背後にある意図について。]
変更を加える前に、[関連するファイル、モジュール、またはパターン] をお読みください。
参考にします]。
解決策は [制約]、[制約]、[制約] でなければなりません。しないでください
[対象外項目]をタッチしてください。
[期待をテストする -- 例: 「新しいロジックの単体テスト」] を書きます。
既存のスイートが合格することを確認します。タスクは次の時点で完了します
[検証可能な基準] と [検証可能な基準]。
このプロンプト テンプレートは、1 行のタスク ステートメントの代わりに、以下について考えて明確に伝えることを奨励します。
単一の必須のタスク ステートメント。一文、能動態、曖昧さなし。 「/users エンドポイントにページネーションを追加する」ではなく、「大規模なユーザー リストを処理するのが良いでしょう」ではありません。
タスクの背後にある意図。短い「理由」は、予期せぬエッジケースに遭遇したときにエージェントが賢明な判断を下すのに役立ちます。 「…そのため、大規模なデータセットでエンドポイントがタイムアウトすることはありません。」
関連するコンテキストへのポインター。エージェントが何かに触れる前に読み取る必要がある特定のファイル、モジュール、または既存のパターンに名前を付けます。狩りに行かせないでください。
制約。ソリューションが尊重する必要がある技術的、スタイル的、または運用上の制限 — 例: 「パブリック API サーフを変更しないでください」

ace", "auth.ts のエラー処理パターンに従います。", "新しい依存関係はありません"。
明示的な範囲外の項目。合理的なエージェントが公正なゲームであると想定する隣接するものはすべて。ここで明示的に示すことで、善意による行き過ぎを防ぐことができます。
テスト可能性への期待。エージェントは新しいテストを作成するか、既存のテストを更新するか、それともスイートがまだ合格することを確認するだけでよいでしょうか?詳細を指定しないと、これは一貫性のない出力の一般的な原因となります。
検証可能な受け入れ基準。各基準は、テストの合格、特定の CURL 応答、エラーなしの lint 実行など、人間の判断なしでチェックできる必要があります。 「正しく動作する必要がある」ということは基準ではありません。
このテンプレートを使用して作成されたプロンプトの例を次に示します。
オフセットベースのページネーションを「GET /users」エンドポイントに追加します。の
現在、エンドポイントは 1 つのクエリですべてのユーザーを返すため、本番環境が発生します。
テーブルが 500k 行を超えるとタイムアウトになります。
変更を加える前に、既存の `src/routes/users.ts` を読んでください。
`src/routes/posts.ts` のページネーション パターンと `paginate()`
ヘルパーは `src/db/helpers.ts` にあります。
ソリューションはデフォルトで page=1、limit=20 にする必要があるため、既存の呼び出し元は
影響を受けない場合は、「src/lib/errors.ts」のエラー処理パターンを使用してください。
新しい依存関係は導入されません。他のエンドポイントには触れないでください。
ユーザー モデル、またはデータベース スキーマ。
新しいロジックの単体テストを `src/routes/users.test.ts` に追加し、
既存のスイートが合格することを確認します。タスクは次の場合に完了します。
- `GET /users` は、ページ、制限、合計を含むメタ オブジェクトを返します。
- `GET /users?page=2&limit=10` は正しいスライスを返します
- 無効なページパラメータに対するエラー処理が実装されました
完全なプロンプト テンプレートはここで見つけることができます。
構造化コード プロンプトは、曖昧な要件、不明確な範囲、欠落している検証手順、受け入れ基準などのよくある間違いに直接対処します。
モー

ここで重要なのは、すべてのコーディング タスクを開始する前にこれらについて考え、それらを明確にし、その結果、その明確さをコーディング エージェントに伝える必要があるということです。
最後に、関連ファイルや既存のパターンを参照することで、コーディング エージェントのコード理解が大幅に向上し、作成するコード内で従うべき重要なファイル、コード フロー、類似のパターンを発見できるようになります。
回避すべきアンチパターンを促すコード
構造化されたコーディング プロンプトは、コーディング セッションの最適な開始点となります。ただし、避けるべきアンチパターンがまだいくつかあります。
構造化コードのプロンプトに従ってアンチパターンを回避する場合でも、すぐに実装に取り掛かるのではなく、変更を計画することをお勧めします。
Claude Code の計画モードを使用して実装の調整を行う
新しいタスクを開始するときは、エージェントがコーディングを開始する前に、計画モードを使用して実装計画を作成します。プラン モードでは、エージェントは次のことができます。
コードベースを探索してコンテキストを収集する
計画内のあいまいさや仮定を確認する
要件を段階的な計画に分割する
詳細についてユーザーと調整する
ほとんどのコーディング エージェントは計画モードを提供します。これは、Claude Code では Shift + Tab を 2 回押すことによって (または /plan コマンドを使用して) トリガーできます。まず、プラン モードを有効にして、構造化コード プロンプトでコーディング セッションを開始します。
/plan オフセットベースのページネーションを `GET /users` に追加します
エンドポイント... [構造化プロンプトの残りの部分がここにあります]
コーディング エージェントは次のようなプラン ファイルを返します。
計画を手に入れたら、それを直接編集したり、コーディング エージェントに変更を依頼したりして、要件やコードの変更についての理解にエージェントを合わせることができます。
計画フェーズでは、コーディング エージェントが計画を立てるための専用の時間とスペース (トークンに関して) が提供されます。

n 実装を開始する前に。ここでは、何らかの手段で何かを機能させるだけではなく、理解して接続することに重点が置かれています。
計画フェーズなしで開始すると、タスクを完了するために必要な最低限の思考 (および計画) だけを行い、タスクの目標に達したら (またはそれに近づいたら) すぐに停止するようにエージェントに指示することになります。
したがって、計画フェーズのないコーディング セッションが要件を見逃すことがよくあることは当然のことであり、後でさらなる反復、設計の統合が必要になり、計画を立てて始めた場合にはさらに多くの労力が必要になります。
計画を作成したら、要件とコードの変更についての理解に合わせてエージェントと調整します。
要件のあいまいさ、実装の決定、または計画中に妨げとなる問題についてエージェントに依頼してください。
/plan オフセットベースのページネーションを `GET /users` に追加します
エンドポイント...[残りの構造化プロンプトがここに表示されます]。
主要な設計または実装を確認する必要がある場合
詳細、妨げとなる懸念事項、または要件の曖昧さについては、
書く前に明確な質問をするための質問ツール
最終計画。
エージェントは計画に向けて作業を進める際に質問をします。
プラン モードの最も優れた点は、実装を開始する前に編集できる PLAN.md ファイルが返されることです。
フィードバックで注目すべき点は次のとおりです。
インターフェイスの設計、ファイルの配置、または命名。
スコープに何かを追加またはスコープから削除します。
より徹底的なテスト範囲を求めます。
エージェントの仮定に基づいてエージェントを修正します。
アーキテクチャに合わせてコード設計を変更します。
フィードバックを提供すると、エージェントは更新された計画ファイルを提供します。
計画には実装の詳細がすべて含まれていない場合があります。エージェントに計画について明確な質問をし、計画に対する理解をテストします。

段階。
ここで私が尋ねたい質問がいくつかあります。
実装するすべてのインターフェイスの変更をリストします。
この実装によって影響を受けるのはどのファイルですか?
実装の複雑さを大幅に増大させる要件はありますか?
XYZ 機能をどのようにテストするか説明していただけますか?
計画に満足したら、エージェントに最終的に実装を開始するよう依頼できます。
AI コードのレビューとウォークスルーで締めくくります
実装が完了すると、会話履歴と生成されたコードがコーディング エージェントのコンテキスト ウィンドウに表示され、コーディング タスクの出力を改善するさらに別の機会が提供されます。
コーディング エージェントを使用して自動コード レビューを実行することは、実装後に行うことができる最も簡単な成果の 1 つです。これを絶対にスキップしないでください。
Claude Code または Codex に組み込まれている /code-review コマンドを使用できます。ただし、私はクロードのドキュメントから採用した次のプロンプトの方が好きです。
サブエージェントを使用して、「PLAN.md」に対する変更を確認します。チェックする
すべての要件が実装されていること、リストされたエッジケース
テストがあり、タスクの範囲外は何も変更されていません。レポート
スタイルの好みではなく、ギャップです。
これにより、コードレビューと受け入れ基準に対する検証の両方が行われます。
エージェントが問題のリストを返してきたら、明確な質問をしたり、単純に問題の解決を依頼したりできます。
見つかった問題を修正します。
なぜこれが機能するのでしょうか?
コーディング エージェントに自身のコードをレビューさせるのは奇妙に思えるかもしれません。これは、エージェントに、コードを書くことよりもコードをレビューすることに重点を置いて、新鮮な気持ち (またはむしろ新鮮なコンテキスト ウィンドウ) で始めるように依頼することと考えてください。
通常、これらの手順を使用すると、各コーディング セッション後に最小限の労力で 2 ～ 4 つの問題を解決できます。
これも、コーディング エージェントに依頼できる、労力の少ない改善です。

:
新しいコードと更新されたコードのテスト カバレッジを確認する
ファイル。次に、新しいテストを追加してギャップを埋めます。
なぜこれが機能するのでしょうか?
コーディング エージェントは CLAUDE.md または AGENTS.md ファイルを読み取り、使用しているテスト ランナーとそれを実行するコマンドを認識します。
私は主に TypeScript でコーディングしているため、テスト ランナーとして Vitest を使用します。これは、コマンド vitest --coverage によるテスト カバレッジ レポートの生成をサポートします。
コーディング エージェントはこのコマンドを自動的に実行し、観察されたギャップに基づいて、ギャップを埋めるための新しいテストを作成します。
通常、この段階で、自分でコードのレビューを開始する準備が整います。
Vibe コーディングではなくエージェント エンジニアリングを行っている場合は、生成されたコードをすべてリポジトリにコミットする前にレビューする必要があります。ただし、AI が生成したコードの量が常に追いつかないように感じるかもしれません。
AI が生成したこの問題を、もっと AI で解決しましょう!
新しいコードの導入に役立つウォークスルー チュートリアルを作成する
変化します。
- 新しい概念、インターフェース、
コード変更の一部として導入されたフローを作成し、それらを接続します
コードベース内の既存のもの。
- 作業中に発生した落とし穴や意外な点を特定します。
実装。
- 説明にコード スニペットと例を使用します。

[切り捨てられた]

## Original Extract

A repeatable three-phase Claude Code workflow: structured prompt, plan mode, and review. Get reliable code from your coding agent every session.

From Vibe Coding to Agentic Engineering with Claude Code
Jun 23, 2026
21 min read
On this page
You write a prompt, hand it to your coding agent, and wait. A few minutes later, you get back… something. Maybe it compiles. But it ignores half your requirements, invents an abstraction you never asked for, or quietly breaks a pattern the rest of the codebase depends on. So you re-prompt. And re-prompt again. And burn through the context window before you’ve shipped anything.
If that loop feels familiar, the problem usually isn’t the model. Left to guess, a Claude Code session fills the gaps you didn’t realize you’d left open: ambiguous requirements, unclear scope, missing acceptance criteria, and no way for it to verify its own work.
The fix isn’t a cleverer one-off prompt. It’s a repeatable process around the prompt. In this blog, I am going to show you the three-phase Claude Code workflow that I follow whenever I’m working on a coding task:
Thanks to this three-phase process, I am seeing more coding sessions end up in implementing what I need while also avoiding frequent long session runtimes, wasted sessions and context window overflows. Adopting this process shifts you from Vibe Coding to Agentic Engineering, and it’s the difference between hoping for good output and engineering it.
Vibe Coding vs. Agentic Engineering
Vibe Coding is the now-common practice of prompting an agent, skimming the result, accepting it if it looks about right, and moving on. The problem shows up the moment the code has to live in a real codebase: the agent fills any ambiguity with guesses, and you inherit whatever it guessed.
Agentic Engineering is the opposite stance. You treat the coding agent as a capable but literal collaborator and keep the engineering disciplines that earned their place over decades, such as clear requirements, a plan, tests, verification, and review. The agent does more of the typing; you stay responsible for the engineering.
Here’s the contrast at a glance:
The three-phase process below is what Agentic Engineering looks like in day-to-day practice.
Start with a Structured Code Prompt to Capture Intent
A good code prompt gives the coding agent a single task with a well-defined scope and verification criteria. Here’s a prompt template you should copy and fill in before every coding task:
[One imperative sentence describing the task.] [One or two sentences
on why... the intent behind it.]
Before making any changes, read [relevant files, modules, or patterns
to reference].
The solution must [constraint], [constraint], and [constraint]. Do not
touch [out-of-scope items].
Write [testing expectations -- e.g., "unit tests for any new logic"]
and ensure the existing suite passes. The task is complete when
[verifiable criterion] and [verifiable criterion].
Instead of a one-line task statement, this prompt template encourages you to think about and communicate clearly:
A single, imperative task statement. One sentence, active voice, no ambiguity. "Add pagination to the /users endpoint," not "It would be good to handle large user lists maybe."
The intent behind the task. A brief "why" helps the agent make sensible judgment calls when it hits edge cases you didn't anticipate. "…so the endpoint doesn't time out on large datasets."
Pointers to relevant context. Name the specific files, modules, or existing patterns the agent should read before touching anything. Don't make it go hunting.
Constraints. Technical, stylistic, or operational limits the solution must respect — e.g., "don't change the public API surface", "follow the error-handling pattern in auth.ts ", "no new dependencies".
Explicit out-of-scope items. Anything adjacent that a reasonable agent might assume is fair game. Being explicit here prevents well-intentioned over-reach.
Testability expectations. Should the agent write new tests, update existing ones, or just verify the suite still passes? Left unspecified, this is a common source of inconsistent output.
Verifiable acceptance criteria. Each criterion should be checkable without human judgment — a passing test, a specific curl response, a lint run with no errors. "It should work correctly" is not a criterion.
Here is an example prompt written using this template:
Add offset-based pagination to the `GET /users` endpoint. The
endpoint currently returns all users in one query, causing production
timeouts as the table has grown past 500k rows.
Before making any changes, read `src/routes/users.ts`, the existing
pagination pattern in `src/routes/posts.ts`, and the `paginate()`
helper in `src/db/helpers.ts`.
The solution must default to page=1, limit=20 so existing callers are
unaffected, use the error-handling pattern in `src/lib/errors.ts`, and
introduce no new dependencies. Do not touch any other endpoints, the
User model, or the database schema.
Add unit tests for the new logic in `src/routes/users.test.ts` and
ensure the existing suite passes. The task is complete when:
- `GET /users` returns a meta object with page, limit, and total
- `GET /users?page=2&limit=10` returns the correct slice
- Error handling is implemented for invalid page parameters
You can find the complete prompt template here .
The Structured Code Prompt directly addresses common mistakes such as ambiguous requirements, unclear scope, missing verification steps, and acceptance criteria.
More importantly, it forces you to think about these before starting every coding task, gain clarity on them, and, in turn, transfer this clarity to the coding agent.
Lastly, by referencing relevant files and existing patterns, you provide the coding agent with a significant boost in code understanding , helping it discover important files, code flows and similar patterns to follow in the code it writes.
Code Prompting Anti-Patterns to Avoid
A structured coding prompt gives you the best starting point for your coding session. However, there are still a few anti-patterns you should avoid:
Even when you follow the Structured Code Prompt and avoid anti-patterns, we recommend planning the changes rather than jumping straight into implementation.
Use Claude Code’s Plan Mode to Align on Implementation
When starting a new task, use Plan Mode to create an implementation plan before the agent starts coding. Plan Mode allows the agent to:
Explore the codebase and gather context
Check any ambiguities or assumptions in the plan
Break down the requirements into a step-by-step plan
Align with the user on the details
Most coding agents provide a Plan Mode , which in Claude Code you can trigger by pressing Shift + Tab twice (or with the /plan command). First, enable Plan Mode and then start your coding session with a Structured Code Prompt:
/plan Add offset-based pagination to the `GET /users`
endpoint... [rest of your structured prompt goes here]
Your coding agent will return a plan file like this:
Once you have the plan in hand, you can edit it directly or ask the coding agent to make changes to it, allowing you to align the agent with your understanding of requirements and code changes.
The Plan phase provides dedicated time and space (in terms of tokens) for the coding agent to plan before it begins implementation. Here, the focus is on understanding and connecting rather than just making something work by any means.
When you start without a planning phase, you are telling the agent to do just the bare minimum thinking (and planning) needed to complete the task and stop as soon as you hit the task goal (or close to it).
It is, therefore, no surprise that coding sessions without a planning phase often miss requirements, which later require further iteration, design consolidation, and simply more effort if we had started with a plan.
Once you have the plan in hand, it is time to align the agent with your understanding of requirements and code changes.
Ask the agent to run any requirements ambiguity, implementation decisions or blocking issues with you during planning:
/plan Add offset-based pagination to the `GET /users`
endpoint... [rest of your structured prompt goes here].
In case you need to confirm any major design or implementation
details, blocking concerns or requirement ambiguity, use the
question tool to ask me clarifying questions before writing
the final plan.
The agent will now ask you questions as it works towards a plan.
The best part of the plan mode is that it returns a PLAN.md file you can edit before implementation begins.
Here are some things you should focus on in your feedback:
Interface design, file placement or naming.
Adding/removing something from the scope.
Asking for more thorough test coverage.
Correcting the agent on its assumptions.
Changing the code design to fit with the architecture.
Once you have provided feedback, the agent will give you an updated plan file.
Sometimes, the plan does not cover all implementation details. Ask the agent clarifying questions about the plan and test their understanding from the planning phase.
Here are a few questions I like to ask:
List all the interface changes you're going to implement.
Which files are going to be impacted by this implementation?
Is there any requirement that significantly increases implementation complexity?
Can you explain how you will test the XYZ feature?
Once you’re satisfied with the plan, you can ask the agent to finally start implementation.
Close with an AI Code Review and Walkthrough
Once the implementation is complete, we have our conversation history and the generated code in the coding agent's context window, presenting yet another opportunity to improve the coding task output.
Doing an automated code review using your coding agent is one of the lowest-hanging fruit you can go for post-implementation. Never skip this!
You can use the /code-review command, which comes built into Claude Code or Codex. However, I prefer this prompt, adopted from Claude’s docs:
Use a subagent to review the changes against `PLAN.md`. Check
that every requirement is implemented, the listed edge cases
have tests, and nothing outside the task's scope changed. Report
gaps, not style preferences.
This gets you both a code review and a validation against the acceptance criteria.
Once the agent responds with a list of issues, you can ask it clarifying questions or simply ask it to fix the issues.
Fix the issues you found.
Why does this work?
It might seem strange to have the coding agent review its own code. Think of it as asking the agent to start with a fresh mind (or rather, a fresh context window), but with a focus on reviewing code rather than writing code.
Usually, these steps help me fix 2 to 4 issues after each coding session with minimal effort.
This is another low-effort improvement that you can ask your coding agent to do for you:
Check the test coverage for the new code and updated code
files. Then close any gaps by adding new tests.
Why does this work?
The coding agent reads your CLAUDE.md or AGENTS.md file and knows which test runner you are using, as well as the commands to run it.
Since I mainly code in TypeScript, I use Vitest as my test runner, which supports generating a test coverage report with the command vitest --coverage .
The coding agent automatically runs this command and, based on the observed gaps, writes new tests to close them.
Usually, at this stage, I am ready to start reviewing the code myself.
If you’re doing Agentic Engineering rather than Vibe Coding, you should review all generated code before committing it to your repository. However, it may feel like there is always more AI-generated code than you can keep up with.
Let's tackle this AI-generated problem with more AI!
Write a walkthrough tutorial to help me onboard onto the new code
changes.
- Make sure to introduce me to the new concepts, interfaces and
flows introduced as part of the code changes and connect them to
existing ones in the codebase.
- Identify any gotcha or surprises that came up during the
implementation.
- Use code snippets and examples in your explanations.

[truncated]
