---
source: "https://www.lightningjar.com/blog/tool-history-footgun"
hn_url: "https://news.ycombinator.com/item?id=48830900"
title: "Deprecated Accessor Footgun in AI SDK v7"
article_title: "Deprecated Accessor Broke My Benchmark | Tool-History Footgun - LJ"
author: "kevinpeckham"
captured_at: "2026-07-08T12:15:05Z"
capture_tool: "hn-digest"
hn_id: 48830900
score: 1
comments: 0
posted_at: "2026-07-08T12:11:40Z"
tags:
  - hacker-news
  - translated
---

# Deprecated Accessor Footgun in AI SDK v7

- HN: [48830900](https://news.ycombinator.com/item?id=48830900)
- Source: [www.lightningjar.com](https://www.lightningjar.com/blog/tool-history-footgun)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T12:11:40Z

## Translation

タイトル: AI SDK v7 の非推奨のアクセサー フットガン
記事のタイトル: 非推奨のアクセサーがベンチマークを破った |工具の歴史 フットガン - LJ
説明: 私たちのベンチマークにおけるすべての複数ターンの失敗は、一行の会話配管に遡っていました。非推奨のアクセサーは依然として型チェックを行っていますが、v7 では暗黙的に意味を変更し、モデルを消去していました。

記事本文:
-->
非推奨のアクセサーがベンチマークを破りました |工具の歴史 フットガン - LJ
Lightning Jar をハックする - Web スタジオ
サイトナビゲーション
まだ型チェックを行う非推奨のアクセサーが私のベンチマーク (そしておそらくあなたのエージェント) を破りました
関連メモ: この記事は、「 We Benchmarked It および HTML as a Native Data Format for LLMs 」に続きますが、この記事はマークアップに関するものではありません。それは、小型モデルがマルチターン工具呼び出しを実行できるかどうかを静かに決定する、配管に関するたった 1 行の会話です。
以下はベンチマークの転写です。モデルはツール呼び出しを使用してノードを作成したところです。システムはそれに ID m1 を割り当て、モデルにそのように指示しました。 1 ターン後、考えられる限り最も単純な編集を要求します。
ユーザー: ID「m1」(前のステップで作成したウィジェットスロット)のノードで「requireBleed」属性をtrueに設定します。ツールを使用して変更を加えてから、「DONE」と返信します。
ツール呼び出しはありません。編集はありません。ただ：完了しました。明るく自信に満ちた嘘。
これらのうち 114 件ありました (転写物がキャプチャされた壊れたプロトコルの、新しく、完全に機器化された再実行からのカウント。元の実行には 110、実行ごとの分散、同じ疾患がありました)。 114 件のすべてが同じ方法でトランスクリプト分類されました。正しい指示が入力され、ツールは「完了」と呼びかけます。モデルが小さくなるほど、状況は悪化しました。gemini-3.5-flash はこれらのフォローアップの 96% で失敗し、claude-haiku-4.5 は 71% で失敗しましたが、gpt-5.4 と claude-sonnet-4.5 は成功しました。私たちはこの調査結果を番号付きで公開しました。マルチターン編集では、ツリー全体の書き換えが詳細なツールの呼び出しを 33 ポイント上回りました。小型モデルとマルチターンの脆弱性についてきちんとした話をしました。
話が間違っていました。データではありません。データは本物で、再現可能で、温度ゼロの本物でした。その解釈は間違っており、私たちが見つけた方法は、このベンチマークがこれまでに生み出した中で最も有益なものでした。
失敗MODなので

これは非常に奇妙だったので (ターン 1 では命令を完璧に実行し、ターン 3 では同じ命令形状を無視するモデル)、メカニズムを特定するために追跡調査を事前に登録しました。会話の深さ。緩和のプロンプト。新しいコンテキスト コントロール。 6 モデル、2,160 セル。
アブレーションは奇妙に戻ってきました。深度は動作しませんでした。一部のモデルはより多くのターンを介在させて改善しましたが、一部のモデルは一貫性のあるカーブを持たずに悪化しました。 「行動する前に指示を再度説明する」緩和策（誰もが推奨する即時衛生管理の一種）により、gpt-5.4 は 100% から 7.5% に低下しました。その間、新鮮な会話により、すべてのモデルが瞬時に、普遍的に修正されました。
一貫性のない結果は通常、自分が測定していると思っているものを測定していないことを意味します。そこで私は、送信していると想定していた会話履歴ではなく、送信している生の会話履歴を読みに行きました。
私たちのハーネスは AI SDK (v7) を使用しました。ツールを呼び出すたびに、実行中の会話に次のように結果が追加されます。
messages.push(...result.response.messages); // 正しく見えます。そうではありません。問題は、マルチステップ ツールの実行では、 result.response.messages には最終ステップのアシスタント テキストのみが含まれるということです。ツールは作成されたモデルを呼び出し、ツールが受け取った結果はステップごとに 1 レベル下に存在します。実際に欲しいものは次のとおりです。
messages.push(...result.responseMessages); // 慣用的な v7 アクセサー
// ステップごとのデータから構築された同等のもの:
messages.push(...result.steps. flatMap((s) => s.response.messages));そして、これが単純なバグではなく移行トラップである理由は、result.response.messages がまさにこの目的のために文書化された v5 パターンであるということです。 v7 ではプロパティを保持し、その意味を最終ステップのみに変更し、蓄積された履歴を新しいトップレベル アクセサー result.responseMessages に移動しました。移行ガイドに記載されていますが、サイ

通話サイトでは nt です。コンパイルエラーはありません。ランタイム信号がありません。 (このプロパティには、@deprecated JSDoc の注記があります。これは、エディター内の人間にとっては取り消し線であり、CI、生成されたコード、およびホバーしていない人には見えません。) v5 に対して書かれたコードは依然として型チェックされ、実行され、静かに別の意味を持ちます。さらに残酷: v7 で降格されたパターンは、4.0 移行ガイドが全員に採用するように指示したものです。4.0 では、 responseMessages が非推奨となり、 response.messages が使用され、v7 では蓄積された履歴が responseMessages に戻されました。この名前の意味は 2 回入れ替わりました。これを vercel/ai#16840 としてアップストリームにファイルしました。メンテナはこれを、バグではなく、文書化された通りの動作、つまりガードレールと移行ツールの機能拡張リクエストとしてトリアージしました。プロジェクトの定義によれば、これは擁護可能な主張であり、この投稿は 1 つの SDK を告発するものではありません。それは、このクラスの変更 (安定した型チェック名の背後にあるセマンティック スワップ) がエージェントの会話配管に遭遇したときに何を行うかについてです。
一行。しかし、最初のバージョンでは、実行したすべてのマルチターン会話に穴がありました。モデル自体のツール アクティビティが履歴から消去されていました。モデルの観点から見ると、過去は次のように見えました。ユーザーが編集を要求しました。私は「完了」と答えました。 insertNode 呼び出しはありません。新しい ID ではツール結果がありません。ユーザーからのリクエストと一言の回答で、明らかに全員が満足しました。
モデルの目でその記録をもう一度読むと、その動作は神秘的ではなくなります。あなたの指示を無視しているわけではありません。コンテキストに含まれる唯一のパターンを継続しています。
追跡調査はすでに構築されていたため、修正された履歴を使用して再実行すると、珍しいことが得られました。1 つの配管決定、プロトコルごとに 2,160 個のセル、同じタスク、同じモデル、同じプロンプトの完璧な A/B です。
モデルのツール呼び出しを非表示にすると、

履歴、フォローアップの実行は、モデルと構成に応じて 5% ～ 100% の範囲で行われました。正しい履歴がある場合: 100%。どのモデルも。すべての腕。あらゆる深さ。 2,160対2,160。
元のベンチマークのマルチターン タスク (ツール条件、プール):
(参照ファミリー タスク、プールされた両方のツール条件、パリティ プロンプト、プロトコルごとのモデルごとに n = 80、results/analysis-permodel-reference.txt の完全なテーブルとスライス定義。)
モデルごとのマルチターン参照編集の成功、プールされたツール条件: モデル自身のツール呼び出しが履歴から隠蔽されているか、履歴が修正されているか。
同じタスク。同じモデル。同じプロンプトが表示されます。唯一変わったのは、会話履歴にモデル独自のツール呼び出しが含まれているかどうかです。 (Gemini の残りの修正履歴障害はすべてフェーズ 1 の精度であり、最初の挿入を手探りしており、その後のドロップアウトではありません。ドロップアウト クラスはすべてのモデルでゼロになりました。)
おそらくキャッチできない理由
その表の下の 2 行を見てください。フロンティアモデルは欠陥にほとんど気づきません。歴史がどれほど奇妙に見えても、彼らは現在の指示に基づいて行動します。それは美徳のように聞こえますが、それがカモフラージュになるまでは実際にそうです。
なぜなら、エージェント機能の標準的なライフサイクルは次のとおりです。フロンティア モデルに対して構築し、評価に合格し、出荷します。 6 か月後、小規模モデルではアシスタント エンドポイントが 20 倍安くなり、評価 (フロンティア モデルで実行) はまだ合格していると誰かが指摘しました。欠陥はずっと歴史構築の中に存在し、目に見えず、モデルの交換によって爆発するのを待っています。生成されるエラーはスローされず、ログにも記録されず、エラーのようには見えません。元気な「DONE」みたいです。
そして、アブレーションによるもう 1 つの罠: 迅速な緩和は、モデル固有の壊滅的な WA における壊れた歴史と相互作用します。

ys: その restate-first 命令は、破損した履歴の下では gpt-5.4 を 100% から 7.5% に低下させましたが、正しい履歴の下では無害でした。プロンプトの微調整によって、あるモデルのツールの使用が不可解に破壊されるのを見たことがあれば、プロンプトを非難する前に履歴を監査してください。
抽象化が主張する実際の配列ではなく、ツールを使用した会話の N+1 ターン目に送信する正確なメッセージ配列をダンプします。ターン N からツール呼び出し部分とツール結果部分を探します。アシスタント ターンがすべて裸のテキストである場合は、バグがあります。
永続化レイヤーを確認してください。 SDK レベルの修正は result.responseMessages (公式の累積履歴アクセサー) または同等の step. flatMap(s => s.response.messages) ですが、同じ穴が 1 層開いています。会話ストアがアシスタントの表示テキスト (非常に一般的: ユーザーに表示されるテキスト) のみを保存する場合、SDK が何を返したかに関係なく、再構築された履歴はすべてツール不要です。
回帰テストを 1 つ追加します。2 ターンのツール会話を実行します。ターン 2 の再構築履歴には、少なくとも 1 つのツール呼び出し部分と 1 つのツール結果部分が含まれていることをアサートします。エージェントエンジニアリングの中で最も安い保険です。
最高のモデルではなく、最も安価なモデルで評価を再実行します。小さなモデルはコンテキスト欠陥のカナリアです。俳句クラスのモデルが突然フォローできなくなったとしても、ツールが下手だと結論づけないでください。私たちは公の場で、信頼区間を使ってフォローできました。
構造的な解決策もあります。モデルの作業成果物が目に見える応答 (アーティファクト全体の書き換え、パッチ形式) であるインターフェイスは、この障害クラス全体の影響を受けません。失うべき隠れたツール チャネルはありません。修正されたベンチマークでは、ツールとリライトは相互に数ポイント以内で実行されます。アーティファクト全体の設計に対する正直な議論は、もはや精度ではなく、静かに故障する可能性のある配管が厳密に少ないということです。

私たちは、このバグが生み出した数字を公表しました。「ツール上の書き換えで +5.3 ポイント」、「マルチターン編集で +33 ポイント」、「ツールの脆弱性は小規模モデルの問題」。修正: ツールとリライトは同等です (実際にはツールが 2 ポイントリードしています)。マルチターンギャップはなくなりました。 Haiku のツールスコアは、Haiku をまったく変えることなく 80.5% から 95% に跳ね上がりました。修正はベンチマーク リポジトリ内にライブであり、以前の両方の投稿に注釈が付けられており、元の主張は表示されたままになっています。これが事前登録といずれかの方法での公開の目的です。生き残った調査結果 (形式と流暢性の同等性、トークンエコノミクス、大きなツリー上で折りたたまれる位置指定の JSON パッチ、最低コストでの書き換えに一致する ID アンカー付きパッチ) は、それらが実行されたプロトコルが壊れたコード パスに触れなかったため、生き残っています。
最後に、少し不快な考えを一つ。マルチターンツールのベンチマークは現在どこにでもあり、そのどれもが結果セクションで誰も説明していないフレームワークコードで会話履歴を構築しています。 +33 ポイントのファントム効果には、信頼区間、p 値、温度 0、および 8,000 のスコアランが含まれていました。厳格さが私たちを救ったわけではありません。トランスクリプトを読んでそうしました。マルチターン評価 (または信頼) を維持している場合は、モデルが実際に何を認識しているかを確認するために午後を費やす価値があります。
更新 (2026 年 7 月): その後のセッション研究 (研究 K、140 回の 12 編集セッション) で、反対方向からの診断が確認されました。独自の編集履歴が完全に表示された状態で、ジェミニ 3.5 フラッシュは、最終セッション 3 位で自身のパッチの 12 ターンを 96.3% で追跡し、ほとんど変動しませんでしたが、フロンティア モデルは、一度だけ表示された木の古い精神的なイメージに依存し、83.8% まで低下しました。隠された歴史は病気だった。どの層においても、可視性が解決策となります。詳細: エージェントのセッションが漂流しています。
「厳格さは救わなかった」

私たち;記録を読んでみたらそうだった。」
日々ますますデジタル化が進む世界でビジネスの成功を支援します。
Lightning Jar は、国内外のビジネスクライアントにプレミアムなデジタルデザインとテクノロジーサービスを提供します。リクエストに応じてポートフォリオと参考資料を入手できます。

## Original Extract

Every multi-turn failure in our benchmark traced back to one line of conversation plumbing: a deprecated accessor that still typechecks but silently changed meaning in v7, erasing the model

-->
Deprecated Accessor Broke My Benchmark | Tool-History Footgun - LJ
hack Lightning Jar - Web Studio
Site Navigation
A Deprecated Accessor That Still Typechecks Broke My Benchmark (and Maybe Your Agent)
Companion note: This piece follows We Benchmarked It and HTML as a Native Data Format for LLMs , but this one isn't about markup. It's about a single line of conversation plumbing that quietly decides whether small models can do multi-turn tool calling at all.
Here's a transcript from our benchmark. The model has just created a node with a tool call; the system assigned it id m1 and told the model so. One turn later, we ask for the simplest edit imaginable:
User: Set the "requireBleed" attribute to true on the node with id "m1" (the widget-slot you created in the previous step). Make the changes with the tools, then reply DONE.
No tool call. No edit. Just: DONE. A cheerful, confident lie.
We had 114 of these (the count from a fresh, fully instrumented re-run of the broken protocol with transcripts captured; the original run had 110, run-to-run variance, same disease). Every one of the 114 was transcript-classified the same way: correct instruction in, zero tool calls out, "DONE." The smaller the model, the worse it got: gemini-3.5-flash failed 96% of these follow-ups, claude-haiku-4.5 failed 71%, while gpt-5.4 and claude-sonnet-4.5 sailed through. We published the finding with a number on it: whole-tree rewrite beat granular tool calling by 33 points on multi-turn edits. We had a tidy story about small models and multi-turn fragility.
The story was wrong. Not the data: the data was real, reproducible, temperature-zero real. The interpretation was wrong, and the way we found out is the most useful thing this benchmark ever produced.
Because the failure mode was so strange (models that execute an instruction flawlessly in turn one, then ignore the identical instruction shape in turn three), we pre-registered a follow-up study to isolate the mechanism. Conversation depth. Mitigation prompts. A fresh-context control. Six models, 2,160 cells.
The ablations came back weird . Depth didn't behave: some models got better with more intervening turns, some worse, no coherent curve. A "restate the instruction before acting" mitigation (the kind of prompt hygiene everyone recommends) made gpt-5.4 collapse from 100% to 7.5%. Meanwhile a fresh conversation fixed every model, instantly, universally.
Results that incoherent usually mean you're not measuring what you think you're measuring. So I went to read the raw conversation histories we were sending, not the ones I assumed we were sending.
Our harness used the AI SDK (v7). After each tool-calling turn, it appended the result to the running conversation like this:
messages.push(...result.response.messages); // looks right. isn't. Here's the thing: in a multi-step tool run, result.response.messages contains only the final step's assistant text . The tool calls the model made, and the tool results it received, live one level down, per step. What you actually want is:
messages.push(...result.responseMessages); // the idiomatic v7 accessor
// equivalent, built from the per-step data:
messages.push(...result.steps.flatMap((s) => s.response.messages)); And what makes this a migration trap rather than a plain bug: result.response.messages was the documented v5 pattern for exactly this purpose. v7 kept the property, changed its meaning to final-step-only, and moved the accumulated history to a new top-level accessor, result.responseMessages . It is documented in the migration guide, but silent at the call site. No compile error. No runtime signal. (The property does carry a @deprecated JSDoc note: a strikethrough for humans in editors, invisible to CI, to generated code, and to anyone not hovering.) Code written against v5 still typechecks, still runs, and quietly means something else. Crueler still: the pattern v7 demoted is the very one the 4.0 migration guide told everyone to adopt: 4.0 deprecated responseMessages in favor of response.messages , and v7 moved the accumulated history back to responseMessages . The name has switched meanings twice. We filed this upstream as vercel/ai#16840 ; the maintainers triaged it as working-as-documented, an enhancement request for guardrails and migration tooling, not a bug. Under the project's definitions, that's a defensible call, and this post isn't an indictment of one SDK. It's about what this class of change (a semantic swap behind a stable, typechecking name) does when it meets agent conversation plumbing.
One line. But with the first version, every multi-turn conversation we ran had a hole in it: the model's own tool activity was erased from its history. From the model's point of view, the past looked like this: user asked for an edit; I replied "DONE." No insertNode call. No tool result with the new id. Just a user request and a one-word answer that apparently satisfied everyone.
Read that transcript again with the model's eyes and the behavior stops being mysterious. It isn't ignoring your instruction. It's continuing the only pattern its context contains.
Because the follow-up study was already built, re-running it with corrected histories gave us something rare: a perfect A/B of one plumbing decision, 2,160 cells per protocol, same tasks, same models, same prompts.
With the model's tool calls hidden from history, follow-up execution ranged from 5% to 100% depending on model and configuration. With correct history: 100%. Every model. Every arm. Every depth. 2,160 for 2,160.
On the original benchmark's multi-turn tasks (tools conditions, pooled):
(Reference-family tasks, both tools conditions pooled, parity prompts, n = 80 per model per protocol; full tables and slice definition in results/analysis-permodel-reference.txt .)
Multi-turn reference-edit success by model, tools conditions pooled: the model's own tool calls hidden from history versus corrected history.
Same tasks. Same models. Same prompts. The only thing that changed is whether the conversation history contained the model's own tool calls. (Gemini's residual corrected-history failures are all phase-one accuracy, fumbling the initial insert, not follow-up dropout. The dropout class went to zero in every model.)
Why you probably wouldn't catch it
Look at the bottom two rows of that table. The frontier models barely notice the defect. They act on the current instruction no matter how odd the history looks, which sounds like a virtue, and is, right up until it becomes camouflage.
Because here's the standard lifecycle of an agent feature: you build it against a frontier model, your evals pass, you ship. Six months later someone points out the assistant endpoint would be 20× cheaper on a small model, and the evals (run on the frontier model) still pass. The defect has been sitting in your history construction the whole time, invisible, waiting for the model swap to detonate it. The failure it produces doesn't throw, doesn't log, and doesn't look like an error. It looks like a cheerful "DONE."
And one more trap from our ablations: prompt mitigations interact with the broken history in model-specific, catastrophic ways: that restate-first instruction took gpt-5.4 from 100% to 7.5% under the broken history and was harmless under the correct one. If you've ever seen a prompt tweak inexplicably nuke one model's tool use, audit your history before you blame the prompt.
Dump the exact message array you send on turn N+1 of a tool-using conversation, not what your abstraction claims, the actual array. Look for tool-call and tool-result parts from turn N. If your assistant turns are all bare text, you have the bug.
Check your persistence layer. The SDK-level fix is result.responseMessages (the official accumulated-history accessor) or the equivalent steps.flatMap(s => s.response.messages) , but the same hole opens one layer up: if your conversation store saves only the assistant's visible text (very common: that's what the user sees), every rebuilt history is tool-less no matter what the SDK returned.
Add one regression test : run a two-turn tool conversation; assert the rebuilt history for turn two contains at least one tool-call part and one tool-result part. It's the cheapest insurance in agent engineering.
Re-run your evals on your cheapest model, not your best one. Small models are your canary for context defects. If haiku-class models suddenly can't follow up, don't conclude they're bad at tools: we did, in public, with confidence intervals.
There's also a structural way out: interfaces where the model's work product is its visible reply (whole-artifact rewrite, patch formats) are immune to this entire failure class. There's no hidden tool channel to lose. In our corrected benchmark, tools and rewrite perform within a couple of points of each other; the honest argument for whole-artifact designs is no longer accuracy, it's that they have strictly less plumbing that can silently fail.
We'd published numbers this bug manufactured: "+5.3 points for rewrite over tools," "+33 points on multi-turn edits," "tool fragility is a small-model problem." Corrected: tools and rewrite are at parity (tools actually edge ahead by 2 points); the multi-turn gap is gone; haiku's tool score jumped from 80.5% to 95% without haiku changing at all. The corrections are live in the benchmark repo and annotated in both earlier posts, with the original claims left visible: that's what pre-registration and publishing-either-way are for. The findings that survive (format-fluency parity, token economics, positional JSON Patch collapsing on big trees, id-anchored patches matching rewrite at the lowest cost) survive because the protocols they ran under never touched the broken code path.
One last, slightly uncomfortable thought. Multi-turn tool benchmarks are everywhere right now, and every one of them constructs conversation histories in framework code that nobody's results section describes. Our +33-point phantom effect had confidence intervals, p-values, temperature 0, and 8,000 scored runs behind it. Rigor didn't save us; reading the transcripts did. If you maintain a multi-turn eval (or trust one), it's worth an afternoon to go look at what your models actually see.
Update (July 2026): A later session study (Study K, 140 twelve-edit sessions) confirmed the diagnosis from the opposite direction: with its own edit history fully visible, gemini-3.5-flash tracked twelve turns of its own patches at 96.3% in the final session third, barely drifting, while the frontier model leaned on a stale mental picture of a tree shown only once and fell to 83.8%. Hidden history was the disease; visibility is the cure, at every tier. Details: Your Agent's Session Is Drifting .
"Rigor didn't save us; reading the transcripts did."
Helping business thrive in a world that is more relentlessly digital every day.
Lightning Jar provides premium digital design and technology services to national and international business clients. Portfolio and references available upon request.
