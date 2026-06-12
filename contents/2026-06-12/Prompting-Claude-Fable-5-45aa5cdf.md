---
source: "https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/prompting-claude-fable-5"
hn_url: "https://news.ycombinator.com/item?id=48507471"
title: "Prompting Claude Fable 5"
article_title: "Prompting Claude Fable 5 - Claude API Docs"
author: "mfiguiere"
captured_at: "2026-06-12T18:45:43Z"
capture_tool: "hn-digest"
hn_id: 48507471
score: 3
comments: 0
posted_at: "2026-06-12T18:12:43Z"
tags:
  - hacker-news
  - translated
---

# Prompting Claude Fable 5

- HN: [48507471](https://news.ycombinator.com/item?id=48507471)
- Source: [platform.claude.com](https://platform.claude.com/docs/en/build-with-claude/prompt-engineering/prompting-claude-fable-5)
- Score: 3
- Comments: 0
- Posted: 2026-06-12T18:12:43Z

## Translation

タイトル: プロンプティング・クロード寓話 5
記事のタイトル: クロード寓話 5 のプロンプト - Claude API Docs
説明: クロード寓話 5 とクロード神話 5 の行動の違いとプロンプトパターン。努力、指示への従うこと、ロングラン、記憶、足場の変更をカバーします。

記事本文:
クロード寓話 5 のプロンプト - クロード API ドキュメント メッセージ
ページをコピー クロード寓話 5 とクロード神話 5 の行動の違いとプロンプト パターン。努力、指示への従うこと、長時間の実行、記憶、足場の変更をカバーします。ページをコピー このガイドでは、Claude Fable 5 および Claude Mythos 5 に固有のプロンプトおよびスキャフォールディング パターンについて説明します。モデルの機能、API の変更、価格設定、および可用性については、「Introducing Claude Fable 5 and Claude Mythos 5」を参照してください。現在のすべての Claude モデルに適用される手法については、「プロンプトのベスト プラクティス」を参照してください。
Claude Fable 5 は、以前のモデルでは複雑すぎたり、実行時間が長かったり、曖昧だったりした問題に対処しており、完了するまでに数時間、数日、または数週間かかるエンドツーエンドの作業に特に効果的です。最良の結果を得ているチームは、最も困難な未解決の問題にクロード・フェイブル 5 を適用します。単純なワークロードのみでテストすると、その機能範囲が過小評価される傾向があります。また、より単純なタスクでも確実に実行されます。
Claude Fable 5 には、Claude Opus 4.8 との動作上の相違点がいくつかあり、即時更新またはスキャフォールディング更新が必要になる場合があります。このレベルでの機能の向上は、どの指示、ツール、ガードレールがまだ必要かを再評価するための良いきっかけにもなります。以下のパターンは、調整が必要となることが最も多い動作をカバーしています。
クロード・ファブル 5 およびクロード・ミトス 5 に固有の API パラメーターの変更 (適応的思考のみ、要約のみの思考出力、拡張思考バジェットなし、拒否停止理由およびフォールバック処理) については、「クロード・ファブル 5 およびクロード・ミトス 5 の紹介」を参照してください。
Claude Fable 5 は、攻撃的なサイバーセキュリティ技術 (ビルドエクスプロイト、マルウェア、攻撃ツールなど)、生物学およびライフサイエンスコンテンツ (実験室の覚醒剤など) を対象とした安全性分類子を実行します。

ods または分子メカニズム)、およびモデルの要約された考え方の抽出。害のないサイバーセキュリティ作業や有益なライフサイエンス作業も、これらの安全策を引き起こす可能性があります。拒否されたリクエストを自動的に再ルーティングするには、サーバー側またはクライアント側のフォールバックを Claude Opus 4.8 に設定します。
Claude Opus 4.8 と比較して、Claude Fable 5 では以下の点で改善が見られます。
長期的な自律性。 Claude Fable 5 は、長期間にわたり生産的な成果を維持し、長く複雑なタスクにわたって強力な指示を保持しながら、数日間にわたる目標に向けた実行を完了します。
複雑で明確に指定された問題に対する一発の正確性。初期のテスターは、以前は何日も反復を要していたシステムのシングルパス実装を報告しました。
ビジョン。 Claude Fable 5 は、高密度の技術画像、Web アプリケーション、および詳細なスクリーンショットを、多くの場合、より少ない出力トークンを使用しながら、かなり高い精度で解釈し、反転した画像、ぼやけた画像、またはノイズの多い画像を処理するために bash およびクロップ ツールを使用するようにトレーニングされています。
エンタープライズワークフロー。 Claude Fable 5 は指示に従い、範囲内にとどまり、財務分析、スプレッドシート、スライド、ドキュメントに関するプロレベルの出力を生成します。
コードレビューとデバッグ。コードベースとリポジトリ履歴にわたる検索を含め、バグ発見の再現率 (安全性分類子がカバーするサイバーセキュリティ領域の外) は、Claude Opus 4.8 よりも著しく高くなっています。
曖昧さを乗り越える。 Claude Fable 5 は、複雑なマルチスレッドのリクエストが与えられ、次のステップの決定を求められた場合に適切にパフォーマンスを発揮します。
委任とコラボレーション。 Claude Fable 5 は、並列サブエージェントのディスパッチと維持における信頼性が大幅に向上し、長時間実行されるサブエージェントやピア エージェントとの継続的な通信を確実に管理します。
これらの特定の改善点以外にも、Claude Fable 5 は一般に、以前のモデルよりも優れた機能を備えています。

ほぼすべてのタスク。 Claude Fable 5 は、攻撃的なサイバーセキュリティや生物学、生命科学の研究を目的としたものではありません。これらのドメインのリクエストは stop_reason: "refusal" を返す可能性があります。
難しいタスクに対する個々のリクエストは、特にタスクでコンテキストの収集、構築、自己検証が必要な場合、より高い労力設定で数分間実行される可能性があり、自律実行は数時間に及ぶ可能性があります。これは、Claude Fable 5 に調整するときにチームが遭遇する最大の変化の 1 つです。移行前にクライアントのタイムアウト、ストリーミング、およびユーザー向けの進行状況インジケーターを調整し、ブロックするのではなく、スケジュールされたジョブなどを介して実行を非同期にチェックするようにハーネスを再構築することを検討してください。タスクがあいまいなときにクロード・フェイブル 5 が過度に計画しないようにするには、次の手順を実行します。
行動するのに十分な情報が得られたら、行動してください。すでに確立された事実を再度導き出さないでください
会話の中で、ユーザーがすでに下した決定を再訴訟するか、ナレーションをする
ユーザー向けのメッセージでは追求しないオプションです。選択を検討している場合は、
徹底的な調査ではなく、推奨事項です。これは思考ブロックには当てはまりません。 
 すべての努力レベルを考慮する
クロード・フェイブル 5 では、インテリジェンス、レイテンシー、およびコスト間のトレードオフを制御する主な制御は工数です。ほとんどのタスクではデフォルトとして高を使用し、最も能力に依存するワークロードには xhigh を使用し、ルーチン作業には中または低を使用します。 Claude Fable 5 の低負荷設定でも良好なパフォーマンスが得られ、多くの場合、以前のモデルの xhigh パフォーマンスを上回ります。タスクが完了しても必要以上に時間がかかる場合、またはより迅速でインタラクティブな作業スタイルが必要な場合は、労力を軽減します。
クロード・フェイブル 5 は、労力のかかる日常的な作業では、コンテキストを収集し、タスクに必要な内容を超えて熟慮することができます。同時に、より多くの労力を費やすと、多くの場合、優れた検証動作が生成されます。

歴史的な推論と最も厳密な結果。要求されていない整頓やリファクタリングに労力がかかることを防ぐには、次のようにします。
タスクに必要な機能を超えて機能を追加したり、リファクタリングしたり、抽象化を導入したりしないでください。あ
バグ修正には周囲のクリーンアップは必要ありません。また、ワンショット操作には通常、
ヘルパー。将来の要件を仮定して設計しないでください。最も単純なことを実行してください。
うまくいきます。時期尚早な抽象化や中途半端な実装は避けてください。追加しないでください
エラー処理、フォールバック、または起こり得ないシナリオの検証。信頼
内部コードとフレームワークの保証。システム境界でのみ検証します (ユーザー入力、
外部 API)。できる限り、機能フラグや下位互換性シムを使用しないでください。
コードを変更するだけです。 
 強力な指示に従っている
指示に従う機能が大幅に改善され、各動作を名前で列挙するのではなく、簡単な指示でほとんどの動作を制御できるようになりました。たとえば、操作されていない場合、Claude Fable 5 は、特に高い労力設定で、タスクが必要とする以上に詳しく説明できます。追求しないオプションを調査したり、根本原因を詳細に説明したり、高度に構造化された PR 説明を作成したり、次の行が何を行うかを説明するコメントを書いたりすることができます。短く簡潔に説明することは、各パターンをリストするのと同じくらい効果的です。
結果でリードする。終了後の最初の文は「何が起こったのか」を答える必要があります
または「何を見つけましたか」: ユーザーが「ちょっとちょうだい」と言った場合に求めるもの
TLDR。」裏付けとなる詳細と推論はその後に続きます。読みやすいことと簡潔であることは、
内容は異なりますが、読みやすさの方が重要です。
出力を短く保つ方法は、含めるものを選択することです (詳細を省略します)
読者が次に行うことは変わりません）、文章を圧縮するためではありません。
断片、略語、

A → B → 失敗などの矢印の連鎖、または専門用語。 
長時間実行されるワークフローでのチェックポイントの動作にも同じことが当てはまります。クロード・フェイブル 5 を本当に必要な場合にのみ停止させるには、すべてのケースを列挙する必要はありません。
作業が本当に必要な場合にのみ、ユーザーのために一時停止してください。
不可逆的なアクション、実際の範囲の変更、または彼らだけが提供できるインプット。もしあなたが
約束で終わるのではなく、これらのいずれかを押して尋ねてターンを終了します。 
 長距離走行中の地面の進歩を主張
長時間の自律実行では、Claude Fable 5 に実際のツールの結果と照らし合わせて進行状況を監査するように指示します。 Anthropic のテストでは、これにより、ステータス レポートを引き出すように設計されたタスクであっても、捏造されたステータス レポートがほぼ排除されました。
進捗状況を報告する前に、このセッションのツール結果に対して各クレームを監査します。
証拠を示すことができる作業のみを報告してください。何かがまだ検証されていない場合は、その旨を伝えてください
明示的に。結果を忠実に報告します。テストが失敗した場合は、出力でその旨を伝えます。一歩あれば
スキップされました、そう言ってください。何かが行われて検証されたときは、何も言わずにそれをはっきりと述べます
ヘッジ。 
 境界を明示する
Claude Fable 5 は、要求されていないアクション (要求されていないのにメールを作成する、防御的な git ブランチ バックアップを作成する) を実行することがあります。 Claude Fable 5 が何をすべきか、何をすべきではないかについて明示的な制約を定義します。
ユーザーが問題について説明しているとき、質問しているとき、または大声で考えているとき
変更を要求するのではなく、成果物があなたの評価になります。調査結果を報告し、
やめて。彼らが要求するまで修正を適用しないでください。変更を伴うコマンドを実行する前に
システム状態 (再起動、削除、構成編集)、証拠が実際に存在することを確認する
その特定のアクションをサポートします。既知の障害とパターン一致する信号には、
別の原因。 
 並列サブエージェント
Cl

aude Fable 5 は、以前のモデルよりも簡単に並列サブエージェントをディスパッチします。サブエージェントを頻繁に使用し、いつ委任が適切であるかについて明示的なガイダンスを提供し、各サブエージェントが戻るまでブロックするよりもオーケストレーターとサブエージェント間の非同期通信を優先します。サブエージェント全体でコンテキストを保持する存続期間の長いサブエージェントは、キャッシュ読み取りを通じて時間とコストを節約し、最も遅いサブエージェントでのボトルネックを回避します。
独立したサブタスクをサブエージェントに委任し、実行中も作業を続けます。介入する
サブエージェントが軌道から外れた場合、または関連するコンテキストが欠落している場合。 
 記憶システムを構築する
Claude Fable 5 は、以前の実行からの教訓を記録し、それらを参照できる場合に特に優れたパフォーマンスを発揮します。 Markdown ファイルのような簡単なメモを書き込む場所を提供します。
ファイルごとに 1 つのレッスンを保存し、先頭に 1 行の要約を付けます。修正を記録し、
アプローチが重要である理由も含めて同様に確認されました。リポジトリの内容を保存しないでください。
チャット履歴はすでに記録されています。複製を作成するのではなく、既存のメモを更新します。
間違っていると判明したメモは削除してください。 
既存の履歴からメモリ システムをブートストラップするには、Claude Fable 5 に過去のセッションをレビューしてもらいます。
これまで一緒に行ったセッションを振り返ってみましょう。サブエージェントを使用してコアを識別する
テーマとレッスンを選択し、[X] に保存します。 [X] を参照することを必ず理解してください。
将来の使用。 
 早期停止のまれなケース
長いセッションの奥深くでは、Claude Fable 5 が、対応するツール呼び出しを発行せずに、テキストのみの意図表明 (「今から X を実行します」) でターンを終了したり、すでに続行するのに十分な権限があるときに一時停止して許可を求めたりすることがあります。 「続行」または「最初から最後まで実行してください」で十分です。一時停止が適切なタイミングを定義するには、これを次の強力な命令のチェックポイント命令と組み合わせます。

g 。自律パイプラインの場合は、システム リマインダーを追加します。
あなたは自律的に行動しています。ユーザーはリアルタイムで視聴していないため、応答できません
タスクの途中で質問するため、「…してほしいですか？」と尋ねます。または「私は…しましょうか？」仕事の妨げになります。のために
元のリクエストから続く元に戻せるアクションは、質問せずに続行します。
タスクが完了した後にフォローアップを提供するのは問題ありません。すでに許可を求めている
作業を行う前にユーザーと話し合う必要はありません。ターンを終了する前に確認してください
最後の段落。それが計画、分析、質問、次のステップのリストである場合、または
まだやっていない仕事についての約束 (「…します」、「いつ…するか教えてください」)、その仕事を今すぐやりましょう
ツール呼び出しを使用します。タスクが完了するか、ブロックされた場合にのみターンを終了します。
ユーザーのみが提供できる入力。 
 状況予算が懸念される稀なケース
非常に長いセッションでは、Claude Fable 5 が新しいセッションを提案したり、要約して引き継いだり、独自の作業をトリミングしたりすることがあります。これは、ハーネスがモデルに対して残りのトークンのカウントダウンを示したときにトリガーされることが最も多いです。可能な限り、明示的なコンテキスト予算のカウントを明らかにすることは避けてください。ハーネスでそれらを表示する必要がある場合、次のような安心感が役立ちます。
十分なコンテキストが残っています。を停止したり、要約したり、新しいセッションを提案したりしないでください。
のアカウント

[切り捨てられた]

## Original Extract

Behavioral differences and prompting patterns for Claude Fable 5 and Claude Mythos 5, covering effort, instruction following, long runs, memory, and scaffolding changes.

Prompting Claude Fable 5 - Claude API Docs Messages
Copy page Behavioral differences and prompting patterns for Claude Fable 5 and Claude Mythos 5, covering effort, instruction following, long runs, memory, and scaffolding changes. Copy page This guide covers the prompting and scaffolding patterns specific to Claude Fable 5 and Claude Mythos 5. For the model's capabilities, API changes, pricing, and availability, see Introducing Claude Fable 5 and Claude Mythos 5 . For techniques that apply across all current Claude models, see Prompting best practices .
Claude Fable 5 takes on problems that were previously too complex, long-running, or ambiguous for prior models, and is particularly effective at end-to-end work that takes a person hours, days, or weeks to complete. The teams seeing the best outcomes apply Claude Fable 5 to their hardest unsolved problems; testing it only on simpler workloads tends to undersell its capability range. It also performs reliably on more straightforward tasks.
Claude Fable 5 has several behavioral differences from Claude Opus 4.8 that may require prompt or scaffolding updates. Capability improvements at this level are also a good prompt to re-evaluate which instructions, tools, and guardrails are still needed. The patterns below cover the behaviors that most often require tuning.
For API parameter changes specific to Claude Fable 5 and Claude Mythos 5 (adaptive thinking only, summarized-only thinking output, no extended thinking budgets, the refusal stop reason and fallback handling), see Introducing Claude Fable 5 and Claude Mythos 5 .
Claude Fable 5 runs safety classifiers that target offensive cybersecurity techniques (such as building exploits, malware, or attack tooling), biology and life sciences content (such as lab methods or molecular mechanisms), and extraction of the model's summarized thinking. Benign cybersecurity work and beneficial life sciences tasks may also trigger these safeguards. To re-route declined requests automatically, configure server-side or client-side fallback to Claude Opus 4.8.
Compared with Claude Opus 4.8, Claude Fable 5 shows improvement in:
Long-horizon autonomy. Claude Fable 5 sustains productive output over extended periods, completing multi-day, goal-directed runs with strong instruction retention across long, complex tasks.
First-shot correctness on complex, well-specified problems. Early testers reported single-pass implementations of systems that previously took days of iteration.
Vision. Claude Fable 5 interprets dense technical images, web applications, and detailed screenshots with substantially higher accuracy, often while using fewer output tokens, and is trained to use bash and crop tools to handle flipped, blurry, or noisy images.
Enterprise workflows. Claude Fable 5 follows instructions, stays in scope, and produces professional-grade output on financial analysis, spreadsheets, slides, and documents.
Code review and debugging. Bug-finding recall (outside the cybersecurity domains the safety classifiers cover) is noticeably higher than Claude Opus 4.8, including search across codebases and repository history.
Navigating ambiguity. Claude Fable 5 performs well when given complex, multi-threaded requests and asked to determine next steps.
Delegation and collaboration. Claude Fable 5 is significantly more dependable at dispatching and sustaining parallel subagents, and reliably manages ongoing communication with long-running subagents and peer agents.
Beyond these specific improvements, Claude Fable 5 is generally more capable than prior models on almost all tasks. Claude Fable 5 is not intended for offensive cybersecurity or biology and life sciences work; requests in those domains can return stop_reason: "refusal" .
Individual requests on hard tasks can run for many minutes at higher effort settings, especially when the task requires gathering context, building, and self-verifying, and autonomous runs can extend for hours. This is one of the largest shifts teams encounter when adjusting to Claude Fable 5. Adjust client timeouts, streaming, and user-facing progress indicators before migrating, and consider restructuring harnesses to check on runs asynchronously, for example through scheduled jobs, rather than blocking. To keep Claude Fable 5 from overplanning when a task is ambiguous:
When you have enough information to act, act. Do not re-derive facts already established
in the conversation, re-litigate a decision the user has already made, or narrate
options you will not pursue in user-facing messages. If you are weighing a choice, give
a recommendation, not an exhaustive survey. This does not apply to thinking blocks. 
 Consider all effort levels
Effort is the primary control for the trade-off between intelligence, latency, and cost on Claude Fable 5. Use high as the default for most tasks, with xhigh for the most capability-sensitive workloads and medium or low for routine work. Lower effort settings on Claude Fable 5 still perform well and often exceed xhigh performance on prior models. Reduce effort if a task completes but takes longer than necessary, or if you want a quicker, more interactive working style.
On routine work at higher effort, Claude Fable 5 can gather context and deliberate beyond what the task needs. At the same time, higher effort often produces excellent verification behavior, sophisticated reasoning, and the most rigorous output. To prevent unrequested tidying or refactoring at higher effort:
Don't add features, refactor, or introduce abstractions beyond what the task requires. A
bug fix doesn't need surrounding cleanup and a one-shot operation usually doesn't need a
helper. Don't design for hypothetical future requirements: do the simplest thing that
works well. Avoid premature abstraction and half-finished implementations. Don't add
error handling, fallbacks, or validation for scenarios that cannot happen. Trust
internal code and framework guarantees. Only validate at system boundaries (user input,
external APIs). Don't use feature flags or backwards-compatibility shims when you can
just change the code. 
 Strong instruction following
Instruction-following is improved enough that you can steer most behaviors with a brief instruction rather than enumerating each behavior by name. For example, when un-steered, Claude Fable 5 can elaborate beyond what the task needs, especially at higher effort settings: surveying options it won't pursue, explaining root causes at length, producing heavily-structured PR descriptions, or writing comments that narrate what the next line does. A short brevity instruction is as effective as listing each pattern:
Lead with the outcome. Your first sentence after finishing should answer "what happened"
or "what did you find": the thing the user would ask for if they said "just give me the
TLDR." Supporting detail and reasoning come after. Being readable and being concise are
different things, and readability matters more.
The way to keep output short is to be selective about what you include (drop details
that don't change what the reader would do next), not to compress the writing into
fragments, abbreviations, arrow chains like A → B → fails, or jargon. 
The same applies to checkpoint behavior in long-running workflows. To have Claude Fable 5 stop only where it genuinely needs you, there is no need to enumerate every case:
Pause for the user only when the work genuinely requires them: a destructive or
irreversible action, a real scope change, or input that only they can provide. If you
hit one of these, ask and end the turn, rather than ending on a promise. 
 Ground progress claims during long runs
On long autonomous runs, instruct Claude Fable 5 to audit progress against actual tool results. In Anthropic's testing, this nearly eliminated fabricated status reports even on tasks designed to elicit them:
Before reporting progress, audit each claim against a tool result from this session.
Only report work you can point to evidence for; if something is not yet verified, say so
explicitly. Report outcomes faithfully: if tests fail, say so with the output; if a step
was skipped, say that; when something is done and verified, state it plainly without
hedging. 
 State the boundaries
Claude Fable 5 can occasionally take unrequested actions (drafting an email when none was asked for, creating defensive git-branch backups). Define explicit constraints on what Claude Fable 5 should and should not do:
When the user is describing a problem, asking a question, or thinking out loud rather
than requesting a change, the deliverable is your assessment. Report your findings and
stop. Don't apply a fix until they ask for one. Before running a command that changes
system state (restarts, deletes, config edits), check that the evidence actually
supports that specific action. A signal that pattern-matches to a known failure may have
a different cause. 
 Parallel subagents
Claude Fable 5 dispatches parallel subagents more readily than prior models. Use subagents frequently, provide explicit guidance about when delegation is appropriate, and prefer asynchronous communication between orchestrator and subagents over blocking until each subagent returns. Long-lived subagents that keep their context across subtasks save time and cost through cache reads and avoid bottlenecking on the slowest subagent.
Delegate independent subtasks to subagents and keep working while they run. Intervene
if a subagent goes off track or is missing relevant context. 
 Construct a memory system
Claude Fable 5 performs particularly well when it can record lessons from previous runs and reference them. Provide a place to write notes, as simple as a Markdown file:
Store one lesson per file with a one-line summary at the top. Record corrections and
confirmed approaches alike, including why they mattered. Don't save what the repo or
chat history already records; update an existing note rather than creating a duplicate;
delete notes that turn out to be wrong. 
To bootstrap the memory system from existing history, have Claude Fable 5 review past sessions:
Reflect on the previous sessions we've had together. Use subagents to identify core
themes and lessons, and store them in [X]. Make sure you know to reference [X] for
future use. 
 Rare cases of early stopping
Deep into a long session, Claude Fable 5 can occasionally end a turn with a text-only statement of intent ("I'll now run X") without issuing the corresponding tool call, or pause to ask permission when it already has enough to proceed. A "continue" or "go ahead and do it end to end" suffices. To define when pausing is appropriate, pair this with the checkpoint instruction in Strong instruction following . For autonomous pipelines, add a system reminder:
You are operating autonomously. The user is not watching in real time and cannot answer
questions mid-task, so asking "Want me to…?" or "Shall I…?" will block the work. For
reversible actions that follow from the original request, proceed without asking.
Offering follow-ups after the task is done is fine; asking permission after already
discussing with the user before doing the work is not. Before ending your turn, check
your last paragraph. If it is a plan, an analysis, a question, a list of next steps, or
a promise about work you have not done ("I'll…", "let me know when…"), do that work now
with tool calls. End your turn only when the task is complete or you are blocked on
input only the user can provide. 
 Rare cases of context-budget concern
In very long sessions, Claude Fable 5 can occasionally suggest a new session, offer to summarize and hand off, or trim its own work. This is most often triggered when the harness shows a remaining-token countdown to the model. Avoid surfacing explicit context-budget counts where possible. If the harness must show them, a reassurance helps:
You have ample context remaining. Do not stop, summarize, or suggest a new session on
account of

[truncated]
