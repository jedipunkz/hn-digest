---
source: "https://discuss.huggingface.co/t/if-unsure-ask-never-guess-ai-agent-pre-execution-checklist/176632"
hn_url: "https://news.ycombinator.com/item?id=48490713"
title: "Unknown Is Not False: An AI Agent Pre-Execution Checklist"
article_title: "If unsure, ask. Never guess. — AI Agent Pre-Execution Checklist - Research - Hugging Face Forums"
author: "offaxis"
captured_at: "2026-06-11T16:28:52Z"
capture_tool: "hn-digest"
hn_id: 48490713
score: 3
comments: 0
posted_at: "2026-06-11T14:17:42Z"
tags:
  - hacker-news
  - translated
---

# Unknown Is Not False: An AI Agent Pre-Execution Checklist

- HN: [48490713](https://news.ycombinator.com/item?id=48490713)
- Source: [discuss.huggingface.co](https://discuss.huggingface.co/t/if-unsure-ask-never-guess-ai-agent-pre-execution-checklist/176632)
- Score: 3
- Comments: 0
- Posted: 2026-06-11T14:17:42Z

## Translation

タイトル: 未知は偽りではない: AI エージェントの実行前チェックリスト
記事タイトル: わからない場合は質問してください。決して推測しないでください。 — AI エージェント実行前チェックリスト - 研究 - ハグフェイス フォーラム
説明: 不明な場合は質問してください。決して推測しないでください。
AI エージェント実行前チェックリスト
人間は不完全な指示を出します。
AI はそれらを完全であるかのように解釈しようとします。
そこから問題が始まります。
エージェントが実行中に実行されると…

記事本文:
= 40rem)" rel="スタイルシート" data-target="デスクトップ" />
= 40rem)" rel="スタイルシート" data-target="discourse-ai_desktop" />
= 40rem)" rel="スタイルシート" data-target="discourse-reactions_desktop" />
= 40rem)" rel="スタイルシート" data-target="poll_desktop" />
ハグフェイスフォーラム
不明な場合は質問してください。決して推測しないでください。 — AI エージェント実行前チェックリスト
ジャンウ
2026 年 6 月 9 日、午前 5 時 18 分
1
不明な場合は質問してください。決して推測しないでください。
AI エージェント実行前チェックリスト
人間は不完全な指示を出します。
AI はそれらを完全であるかのように解釈しようとします。
そこから問題が始まります。
必要な情報が検証されていない状態でエージェントが実行された場合、結果は単なるモデルエラーではなく、実行前の確認の失敗となります。バイブコーディングの間違いの多くは、コーディング能力の欠如によるものではなく、実行前に必要な質問が行われなかったことに起因します。
AI は、自分が知らないことを確実に知ることはできません。そのため、単純な推奨だけでは十分ではありません。必要なのは、構造的な強制ロジック、つまり何らかのアクションが取られる前に検証する必要があるチェックリストです。
このドキュメントでは、コードではなくルールを通じて AI の実行を制御する方法を提案します。実装はさまざまですが、中心的な原則は次のとおりです。動作をハードコーディングするのではなく、ルールを宣言し、AI にそれに従わせるのです。
この文書は原則と構造を定義します。それらがどのように適用されるかは、各システム、組織、規制機関の独自のポリシーと協定によって異なります。
コア構造 (固定チェックリスト)
固定チェックリスト — すべてのアクションに共通する最低限必要な項目。実行ユニットを定義します
プロバイダーチェックリスト — プロバイダーが各アクションに対して定義した実行上の注意事項のリスト
ユーザーチェックリスト — ユーザーが指定したルールとガイドライン
基本原則: 不明な点が残っている場合は、続行しないでください。用途は必ず聞いてください

r または実行を保留します。
指示のギャップ — ユーザーの指示が不完全であり、AI が推論によって空白を埋めます。
アクション定義のギャップ — AI は、実際には理解していないにもかかわらず、新しいアクション、API、デバイス、ツール、またはワークフローを理解しているかのように処理します。
実行前に、プロバイダーの基準とユーザーの判断に基づいてチェックリストが作成されます。実行時に、AI がチェックリストを検証します。
プロバイダーとユーザーの両方が自然言語で注意事項やルールを記述することができます。 AI はそれらを JSON に構造化します。実行ルールをハードコーディングするのではなく、実行前に確認する必要があるものを記録して検証します。
つまり、新しいアクションが表示されるたびにコードを変更するのではなく、自然言語チェックリストを追加または更新するだけで済みます。
このフレームワークにおける基本的な実行単位は、ツールではなくアクションです。 1 つのツールに複数のアクションを含めることができ、各アクションには異なる予防措置が適用される場合があります。したがって、チェックリストはツールごとではなく、アクションごとに作成する必要があります。
これらは各アクションに必要な最低限の項目です。ほとんどの場合、C1、C2、C3 はユーザーが指示した瞬間に自動的に入力されます。
実行単位は、プロバイダー アクション名だけでは決まりません。同じアクションであっても、実行ユニットは、実行時期 (C1)、ユーザーがそれを意味する意味 (C2)、システムが呼び出す内容 (C3) の 3 つがすべて一緒に指定された場合にのみ定義されます。
プロバイダー アクション名によって、呼び出されるものが決まります。ただし、実行ユニットは C1 + C2 + C3 の組み合わせで定義されます。
プロバイダー チェックリストは、プロバイダーがアクションに対して定義した実行上の注意事項のリスト、つまりアクションの取扱説明書です。 AI が実行時に読み取ることができる場所であればどこにでも配置できます。
必要なアイテムが不足している場合、責任の問題が生じます。
チェックリストには以下を含めることができます

技術的な質問だけでなく、安全性、倫理、法律、組織ポリシー、業界標準などの項目も含まれます。これは、新しいアクション、物理的なアクション、元に戻せないアクション、法的または倫理的な判断を必要とするアクションの場合に特に重要です。
AIは現実を直接観察することはできません。システム、センサー、ログ、API、またはユーザー入力を通じて確認できないプロバイダー チェックリストの値はすべて、 不明 としてマークする必要があります。不明は嘘ではありません。未知だからといって安全というわけではありません。
ユーザー チェックリストには、特定の実行に対してユーザーが指定する追加のルールとガイドラインが含まれています。プロバイダー チェックリストがアクション独自の予防策をカバーするのに対し、ユーザー チェックリストはユーザーの目的、状況、好み、制約を反映します。
ユーザー チェックリストは、実行のたびに新たに追加したり、繰り返し使用するために保存したりできます。 AI は、続行するかどうかを決定する前に、プロバイダー チェックリストとユーザー チェックリストの両方を検証します。
ユーザーの指示
→ 修正されたチェックリストを確認する
→ プロバイダーの確認チェックリスト
→確認できる値を確認する
→ 必要に応じてユーザーに質問する
→ ユーザー追加チェックリスト
→不明な点がないか確認する
→実行するかどうかを決定する
実行フローの例: コード修正
ユーザーはファイルを添付して「ログイン エラーを修正してください」と言います。
AI はまず固定チェックリストを検証します。
C1 時期 / ケース: 今すぐ修正が必要です
C2 ユーザーアクション名: ログインエラーを修正する
C3 プロバイダーのアクション名: edit_existing_code
次に、AI がプロバイダー チェックリストをチェックします。プロバイダーは次のような予防措置を提供する場合があります。
ユーザーは独自のユーザー チェックリストを追加できます。
既存のデザインをそのまま維持します。
ログイン以外の機能には触れないでください。
変更を加える前に、どのファイルが変更されるかを教えてください。
AI はこの入力に基づいてチェックリストを作成します。不足している項目がある場合は、続行する前にユーザーに質問します。
人間は不完全な指示を出します。
インステ

それらを完全なものとして扱うという広告、
AIはまずチェックリストが満たされているかどうかをチェックし、
答えが見つからない場合は尋ねます。
そして、不明なものを不明としてマークします。
実行は命令が完了した場合にのみ開始されます。
チェックリストは、一度作成すると、次回の実行時に再利用できます。
C1.いつ/場合
ユーザーはいつ、またはどのような場合にこのアクションを有効にしたいのでしょうか?
この質問により、命令の性質が決まります。つまり、命令が即時コマンドであるか、条件付き実行であるか、スケジュールされた実行であるか、イベント トリガーによる実行であるか、または特定の状況での繰り返しルールであるかが決まります。
例: 現在 / 毎日午前 7 時 / 帰宅時 / ファイルがアップロードされるたび / ユーザーが最終承認を与えた後
C2.ユーザーアクション名
ユーザーはこのアクションを何と呼んでいますか?
ユーザーが割り当てる名前には、目的、生活背景、意図が含まれる場合があります。これは単なるラベルではなく、ユーザーがこのアクションを実際にどのように理解するかを最初に宣言するものです。
例: 朝のウォームアップ / 赤ちゃんの睡眠モード / ログインバグの修正 / 私のエスプレッソ
C3.プロバイダーアクション名
プロバイダーはこのアクションを何として定義しますか?
呼び出し可能なアクションの名前。呼び出される内容を決定し、C1 および C2 とともに実行ユニットを定義します。
例:turn_on_heater / brew_coffee / send_email / edit_existing_code / process_payment
始める前に子供の健康診断を受ける必要がありますか?
目標温度に達しましたか？
最大使用時間を超えていませんか?
このアクションを停止すると損失が発生しますか?
このアクションは繰り返し実行できますか?
安全性、倫理、法的な問題:
このデータ転送には個人情報が含まれますか?
この支払額は法定限度額を超えていますか?
このファイルは会社のポリシーに基づいて外部共有が承認されていますか?
保護者や管理者の承認は必要ですか?
既存のデザインを維持する

そしてUIは変更されていません。
変更を加える前に、変更するファイルのリストを表示します。
テストが失敗した場合は、もう一度質問してください。
既存の書き方を維持します。
数字や固有名詞は変更しないでください。
出典のない主張を含めないでください。
送信する前に下書きを見せてください。
丁寧に、しかし簡潔に書きましょう。
添付ファイルが不足していないか確認してください。
子供が近くにいる場合は実行しないでください。
ユーザーが在宅している場合にのみ実行してください。
実行する前にもう一度確認してください。
金額が10万ウォンを超える場合は再度質問してください。
繰り返しの支払いは許可しないでください。
承認される前に外部に送信しないでください。
C3 プロバイダー アクション名のデフォルトの応答者はプロバイダーです。
プロバイダー アクションが登録されていない場合: システムは一時的なプロバイダー アクション名を生成する場合があります。
システムが一時的な名前を生成する場合、実行前にユーザーに通知し、承認を受ける必要があります。
注: この例外は、初期段階の展開時または新しいアクションの登録時のユーザー エクスペリエンスの低下を避けるために、最小限の柔軟性を提供します。すべての一時的なアクションは、後でプロバイダーによって正式にレビュー、承認、登録される必要があります。
初期段階の展開のフォールバック ルール
プロバイダー チェックリストが存在しないか不完全な場合、ユーザー チェックリストは次の領域を実質的にカバーできます。
安全ルール（ロールバック要件、リスクレベル、実行を停止する条件など）
法的および倫理的制約（個人データの取り扱い、禁止行為、偏見の防止など）
実行条件と制限（時間制限、金額制限、範囲制限、承認プロセスなど）
注: ユーザー チェックリストは、プロバイダー チェックリストの技術ベースライン仕様やプロバイダーが公式に保証するものを完全に置き換えることはできません。
以下の JSON 例は参照のみを目的としており、必須のテンプレートではありません。 AIはwで自由にJSONを構築すればいい

憎しみの形が状況に最もよく適合します。
ハードコーディングを削除することで、新しいアクション、新しいドメイン、予期しない状況に自然に適応できるようになります。新しいキーはいつでも追加でき、必要に応じて構造を変更できます。
柔軟性と実用性が最優先事項です。
{
「修正済み」: {
"c1_when_case": "即時実行が要求されました",
"c2_user_action_name": "ログイン エラーを修正",
"c3_provider_action_name": "既存のコードの編集"
}、
"プロバイダーチェックリスト": [
{
"question": "修正の範囲は明確に定義されていますか?",
"答え": "部分的"、
"note": "ログイン関連ファイルのみを変更する必要がありますが、全範囲を確認する必要があります"
}、
{
"question": "テスト ケースの合格基準は存在しますか?",
"答え": "不明",
"note": "現在のリポジトリには文書化されたテスト自動化基準が見つかりません"
}
]、
"ユーザーチェックリスト": [
{
"rule": "既存のデザインや UI を変更しないでください",
「ステータス」：「確認済み」
}
]、
"実行_決定": "質問ユーザー",
"summary": "ログイン エラーの修正が要求されました。テスト基準は不明な状態です。ユーザーの確認が必要です。",
"リスクレベル": "中",
"suggested_next_step": "ユーザーにテスト基準を尋ねるか、最初に確認のために変更するファイルのリストを提示します"
}
所有権とライセンス
© 2026 AnnaSoft Inc. 韓国
このドキュメントは CC BY-NC-ND 4.0 ライセンスに基づいて配布されます。
非営利的な共有は許可されています。修正版の配布や商業利用には別途同意が必要です。
この記事の重要な点は、AI の推論を難しくすることではありません。
AIが実行前に推測するのを防ぐためです。
実行単位はプロバイダー アクション名だけではありません。
C1 + C2 + C3 で定義されます。
アクションがいつ有効であるか、ユーザーがそれを何と呼ぶか、システムが実際に何を呼ぶか。
必須のチェックリスト項目が欠落している場合、責任は追跡可能になります。

実行前に、AI は固定チェックリスト、プロバイダー チェックリスト、およびユーザー チェックリストをチェックする必要があります。
確認できないものは不明のままです。
不明は嘘ではありません。未知のものは安全ではありません。
MCP の観点から見ると、この文書は次のような提案です。
プロバイダーがアクションごとに実行前チェックリストを自然言語で提供できるようにします。
MCP がアクションを呼び出し可能にする場合、プロバイダー チェックリストは、そのアクションが実行可能になる前に回答する必要がある質問を定義します。
API の時代には、呼び出し形式を記述するだけで十分でした。
エージェント時代では、実行前条件の定義も同様に必要になります。
プロバイダーは質問リストを提供します。
エージェントは回答が存在するかどうかを確認します。
そうでない場合、エージェントはユーザーに尋ねます。
それでも応答がない場合、エージェントは実行されません。
プロバイダーが正確性と安全性のために必要な検証項目を提供しない場合、その欠如は、事件後のレビューでプロバイダーの責任範囲を決定する際に関連する可能性があります。
また、エージェントがユーザーの確認なしに実行した場合、その実行をユーザーが承認した決定として簡単に扱うべきではありません。
顔のハグからレディット、ディスコード、リンクインに至るまで、物語は厳しく管理されており、だからこそ私がここにいるのは、言論の自由を拡大すべきであるということだけです。

[切り捨てられた]

## Original Extract

If unsure, ask. Never guess.
AI Agent Pre-Execution Checklist
Humans give incomplete instructions.
AI tries to interpret them as if they were complete.
That is where the problem begins.
If an agent executes while re&hellip;

= 40rem)" rel="stylesheet" data-target="desktop" />
= 40rem)" rel="stylesheet" data-target="discourse-ai_desktop" />
= 40rem)" rel="stylesheet" data-target="discourse-reactions_desktop" />
= 40rem)" rel="stylesheet" data-target="poll_desktop" />
Hugging Face Forums
If unsure, ask. Never guess. — AI Agent Pre-Execution Checklist
Jang-woo
June 9, 2026, 5:18am
1
If unsure, ask. Never guess.
AI Agent Pre-Execution Checklist
Humans give incomplete instructions.
AI tries to interpret them as if they were complete.
That is where the problem begins.
If an agent executes while required information remains unverified, the result is not simply a model error — it is a pre-execution confirmation failure. Many vibe coding mistakes are not due to lack of coding ability, but because the necessary questions were never asked before execution.
AI does not reliably know what it does not know. That is why a simple recommendation is not enough. What is needed is a structural enforcement logic — a checklist that must be verified before any action is taken.
This document proposes a way to control AI execution through rules, not code. Implementation may vary, but the core principle is this: declare the rules, and let AI follow them — instead of hardcoding behavior.
This document defines principles and structure. How they are applied depends on each system, organization, and regulatory body’s own policies and agreements.
Core Structure (Fixed Checklist)
Fixed Checklist — Minimum required items common to all Actions; defines the execution unit
Provider Checklist — A list of execution precautions defined by the Provider for each Action
User Checklist — Rules and guidelines specified by the user
Core principle: If any unknown remains, do not proceed. Always ask the user or place execution on hold.
Instruction gap — The user’s instruction is incomplete, and AI fills in the blanks through inference.
Action definition gap — AI processes a new Action, API, device, tool, or workflow as if it understands it, when it actually does not.
Before execution, a Checklist is constructed based on the Provider’s standards and the User’s judgment. At execution time, AI verifies the Checklist.
Both Provider and User can write precautions and rules in natural language. AI then structures them into JSON — not to hardcode execution rules, but to record and verify what must be confirmed before execution.
This means that instead of modifying code every time a new Action appears, you simply add or update a natural language Checklist.
The basic unit of execution in this framework is the Action, not the tool. A single tool can have multiple Actions, and each Action may carry different precautions. Therefore, Checklists must be written per Action, not per tool.
These are the minimum items required for every Action. In most cases, C1, C2, and C3 are filled automatically the moment the user gives an instruction.
The execution unit is not determined by Provider Action Name alone. Even for the same Action, the execution unit is only defined when all three are specified together: when it executes (C1), what the user means by it (C2), and what the system calls (C3).
Provider Action Name determines what is called. But the execution unit is defined by C1 + C2 + C3 together.
The Provider Checklist is the list of execution precautions the Provider has defined for an Action — the action’s instruction manual. It can be placed anywhere AI can read at execution time.
When required items are missing, questions of accountability arise.
A Checklist can include not only technical questions but also items covering safety, ethics, law, organizational policy, and industry standards. This is especially important for new Actions, physical Actions, irreversible Actions, and Actions requiring legal or ethical judgment.
AI cannot directly observe reality. Any value in the Provider Checklist that cannot be confirmed through systems, sensors, logs, APIs, or user input must be marked as unknown . Unknown is not false. Unknown does not mean safe.
The User Checklist contains the additional rules and guidelines the user specifies for a particular execution. Where the Provider Checklist covers the Action’s own precautions, the User Checklist reflects the user’s purpose, situation, preferences, and constraints.
A User Checklist can be added fresh for each execution, or saved for repeated use. AI verifies both the Provider Checklist and the User Checklist before deciding whether to proceed.
User instruction
→ Verify Fixed Checklist
→ Verify Provider Checklist
→ Check confirmable values
→ Ask User if needed
→ Add User Checklist
→ Check for unknowns
→ Decide whether to execute
Execution Flow Example: Code Fix
The user attaches a file and says: “Fix the login error.”
AI first verifies the Fixed Checklist:
C1 When / Case: Fix requested now
C2 User Action Name: Fix login error
C3 Provider Action Name: edit_existing_code
AI then checks the Provider Checklist. The Provider may supply precautions such as:
The user can add their own User Checklist:
Keep the existing design intact.
Do not touch any functionality outside of login.
Tell me which files will be changed before making changes.
AI assembles the Checklist from this input. If any item is missing, it asks the user before proceeding.
Humans give incomplete instructions.
Instead of treating them as complete,
AI first checks whether the Checklist has been filled,
asks when an answer is missing,
and marks what it does not know as unknown .
Execution only begins once the instruction is complete.
A Checklist, once written, can be reused for the next execution.
C1. When / Case
When, or in which case, does the user want this Action to be valid?
This question determines the nature of the instruction: whether it is an immediate command, conditional execution, scheduled execution, event-triggered execution, or a recurring rule for a specific situation.
Examples: now / every day at 7 AM / when I get home / whenever a file is uploaded / after the user gives final approval
C2. User Action Name
What does the user call this Action?
The name the user assigns may contain purpose, life context, and intent. It is not a mere label — it is the first declaration of how the user actually understands this Action.
Examples: Morning Warm-up / Baby sleep mode / Fix login bug / My espresso
C3. Provider Action Name
What does the Provider define this Action as?
The name of an invokable Action. It determines what is called, and together with C1 and C2, it defines the execution unit.
Examples: turn_on_heater / brew_coffee / send_email / edit_existing_code / process_payment
Does the child need to be checked for before starting?
Has the target temperature been reached?
Has the maximum operating time been exceeded?
Will stopping this Action cause a loss?
Can this Action be executed repeatedly?
Safety, ethics, and legal questions:
Does this data transfer include personal information?
Does this payment amount exceed the legal limit?
Is this file approved for external sharing under company policy?
Is approval from a guardian or administrator required?
Keep the existing design and UI unchanged.
Show me the list of files to be changed before making changes.
If a test fails, ask me again.
Keep the existing writing style.
Do not change numbers or proper nouns.
Do not include claims without a source.
Show me the draft before sending.
Write it politely but briefly.
Check that no attachments are missing.
Do not execute if a child is nearby.
Only execute when the user is home.
Confirm once more before executing.
Ask again if the amount exceeds 100,000 KRW.
Do not allow repeated payments.
Do not transmit externally before approval.
The default answering party for C3 Provider Action Name is the Provider.
If a Provider Action is not registered: the System may generate a temporary Provider Action Name.
When the System generates a temporary name, it must notify the user and receive approval before execution.
Note: This exception provides minimal flexibility to avoid degrading user experience in early-stage deployments or when registering new Actions. All temporary Actions must later be formally reviewed, approved, and registered by the Provider.
Fallback Rules for Early-Stage Deployments
When the Provider Checklist is absent or incomplete, the User Checklist can substantially cover the following areas:
Safety rules (rollback requirements, risk level, conditions to halt execution, etc.)
Legal and ethical constraints (personal data handling, prohibited actions, bias prevention, etc.)
Execution conditions and limits (time restrictions, amount limits, scope limits, approval processes, etc.)
Note: The User Checklist cannot fully replace the Provider Checklist’s technical baseline specifications or anything the Provider officially guarantees.
The JSON example below is for reference only — it is not a required template. AI should construct JSON freely in whatever form best fits the situation.
Removing hardcoding is what allows natural adaptation to new Actions, new domains, and unexpected situations. New keys can always be added, and the structure can be changed as needed.
Flexibility and practicality are the top priorities.
{
"fixed": {
"c1_when_case": "Immediate execution requested",
"c2_user_action_name": "Fix login error",
"c3_provider_action_name": "edit_existing_code"
},
"provider_checklist": [
{
"question": "Is the scope of modification clearly defined?",
"answer": "partial",
"note": "Only login-related files should be modified, but full scope needs confirmation"
},
{
"question": "Do passing criteria for test cases exist?",
"answer": "unknown",
"note": "No documented test automation criteria found in the current repository"
}
],
"user_checklist": [
{
"rule": "Do not change the existing design or UI",
"status": "confirmed"
}
],
"execution_decision": "ask_user",
"summary": "Login error fix requested. Test criteria are in unknown state — user confirmation required.",
"risk_level": "medium",
"suggested_next_step": "Ask the user for test criteria, or present the list of files to be modified for confirmation first"
}
Ownership & License
© 2026 AnnaSoft Inc. Republic of Korea
This document is distributed under the CC BY-NC-ND 4.0 license.
Non-commercial sharing is permitted. Distribution of modified versions and commercial use require separate agreement.
The key point of this article is not to make AI reason harder.
It is to prevent AI from guessing before execution.
The unit of execution is not the Provider Action Name alone.
It is defined by C1 + C2 + C3:
when the action is valid, what the user calls it, and what the system actually calls.
When a required checklist item is missing, responsibility becomes traceable.
Before execution, the AI must check the Fixed Checklist, Provider Checklist, and User Checklist.
If something cannot be confirmed, it remains unknown.
Unknown is not false. Unknown is not safe.
From an MCP perspective, this document is a proposal:
Let Providers supply a pre-execution Checklist per Action, in natural language.
If MCP makes an Action callable, the Provider Checklist defines the questions that must be answered before that Action becomes executable.
In the API era, describing the call format was enough.
In the Agent era, defining the pre-execution conditions becomes equally necessary.
Provider supplies the question list.
Agent checks whether the answers exist.
If not, Agent asks the user.
If still unanswered, Agent does not execute.
If a Provider does not supply the necessary verification items for accuracy and safety, that absence may become relevant in determining the Provider’s scope of responsibility in a post-incident review.
And if an Agent executes without user confirmation, that execution should not be easily treated as the user’s approved decision.
Id just like to say from hugging face to reddit to discord to linkedin even x the narrative is tightly controlled and thats why i am here at all the thing that should amplify free speech is

[truncated]
