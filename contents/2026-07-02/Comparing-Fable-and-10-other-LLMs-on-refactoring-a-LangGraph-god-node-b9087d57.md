---
source: "https://wtf.korridzy.com/twilight-of-the-gods/"
hn_url: "https://news.ycombinator.com/item?id=48761132"
title: "Comparing Fable and 10 other LLMs on refactoring a LangGraph god node"
article_title: "Twilight of the Gods. Fable and 10 more LLMs on a Code Reorganization Task. Comparison. - wtf"
author: "Korridzy"
captured_at: "2026-07-02T13:30:32Z"
capture_tool: "hn-digest"
hn_id: 48761132
score: 1
comments: 0
posted_at: "2026-07-02T13:19:56Z"
tags:
  - hacker-news
  - translated
---

# Comparing Fable and 10 other LLMs on refactoring a LangGraph god node

- HN: [48761132](https://news.ycombinator.com/item?id=48761132)
- Source: [wtf.korridzy.com](https://wtf.korridzy.com/twilight-of-the-gods/)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T13:19:56Z

## Translation

タイトル: LangGraph ゴッドノードのリファクタリングに関する Fable と他の 10 個の LLM の比較
記事のタイトル: 神々の黄昏。 Fable とさらに 10 個の LLM がコード再編成タスクに参加しています。比較。 - なんてことだ
説明: Korridzy による毎週のトラッカー フィード。

記事本文:
神々の黄昏。 Fable とさらに 10 個の LLM がコード再編成タスクに参加しています。比較。 - なんてことだ
コンテンツにスキップ
なんと！
神々の黄昏。 Fable とさらに 10 個の LLM がコード再編成タスクに参加しています。比較。
英語
材料
材料
材料と生データ
提案
提案
Deepseek 4 プロの提案
レビュー
レビュー
ディープシーク 4 プロレンジ
最優秀アナリスト
最優秀アナリスト
フェイブル5 Ba範囲
論文
論文
ラン1
ラン1
評価者共通論文ランキング
評価者好感度ランキング
ラン2
ラン2
評価者のコンセンサスランキング
ラン3
ラン3
評価者のコンセンサスランキング
カテゴリー
カテゴリー
AIエージェント
もともとの問題
計画ノードが実際に行ったこと
ステージ1。モデルが提案を生成する
提案テーブル
各提案についてもう少し詳しく
寓話-5
ステージ2。モデルが提案を評価する
レビュー表
各レビューについてもう少し詳しく
寓話-5
ステージ3。濡れた水着コンテスト
1つに近づきます。スコアは一致していますか?
最良の提案: 概要
2にアプローチします。論文ごとのレビューの比較
3にアプローチします。意見の中心とメドイド
デウス・エクス・マキナ
アナリストの評価方法
インデックスに戻る
コリッジ
メタデータ
2026 年 6 月 30 日
で
ランググラフ 、
LLM、
AIエージェント、
リファクタリング 、
パイソン
神々の黄昏。 Fable とさらに 10 個の LLM がコード再編成タスクに参加しています。比較。 ¶
Эта статья также доступна на русском: Гибель богов 。
11 個のモデル提案すべて、クロスレビュー、論文の実行、およびランキング スクリプトは、「材料とこの実験の再現」で公開されています。
これは、ある実験の詳細な記録です。私は実際の LangGraph エージェントからゴッド ノードを取得し、5 人のアメリカ人モデルと 6 人の中国人モデルに、最初にそれを解きほぐす方法を提案してもらい、次にお互いの提案を評価してもらいました。その後、どれを信頼すべきかを判断するために 3 つの異なる方法を試しました。

その問題。
本来の問題 ゴッドノードとは何か、なぜ危険なのか
計画ノードが実際に行ったこと
レモンから作るレモネード なぜこのようなことになったのか、そして実験はどのように行われたのか
ステージ1。モデルが提案を生成する
ステージ2。モデルが提案を評価する
ステージ3。濡れ水着コンテスト 上手さを決める
1つに近づきます。スコアは一致していますか?最適な提案を選択する
2にアプローチします。論文ごとにレビューを比較する 最適なアナリストを選択する
3にアプローチします。意見の中心とmedoid 最高のアナリストを再び選出
デウス・エクス・マキナ。もう一度最高のアナリストを選ぶ
要点 どのモデルを生成器として使用するか、どのモデルを評価器として使用するか、そしてどこに心が安らぎを見つけることができるか。
それがどうなるかはご存知でしょう。Data Sanity のコースで仲間たちと練習用 AI エージェントを構築しているとします。そして、急速に増えていく機能のカラフルな渦の中で、プロジェクトの内部エージェントの 1 つが次のような状態グラフ (LangGraph) を持っていることに突然気づきました。
フローチャート TD
planner_start([START]) --> 計画[計画]
計画 -->|検索|検索[検索]
計画 -->|ask_user| ask_user[ask_user / 割り込み]
計画 -->|反省|反射する[反射する]
計画 -->|計算|計算する[計算する]
計画 -->|終了|仕上げ[仕上げ]
検索 -->|last_observation|観察する[観察する]
検索 -->|ヒットなし / バックエンド障害|計画
観察 --> 計画
計算 --> 計画
ask_user --> 観察ユーザー[観察ユーザー]
観察ユーザー --> 計画
反省 --> 計画
終了 --> planner_end([END])
一見すると、これはただのかわいいタコですが、心配する必要はありません。しかし、このタコの控えめな 8 本足の頭にどれだけの論理が含まれているかを知れば、私たちが見ているのはアンチパターンであることがすぐに明らかになります。この場合、それをゴッドノードと呼びます。
プラン ノードは、反復チェック、ブートを含む約 350 行のロジックを非表示にします。

地域と通貨、スキーマの準備、取得タスクのルーティング、LLM 呼び出し、その後の決定の修正などに関する質問をストラップします。
問題は関数のサイズだけではありません。重要なオーケストレーションが単一ノード内に隠蔽されると、グラフはシステムを表現できなくなります。説明するのは難しく、デバッグもテストも難しく、変更するのはより危険です。したがって、明らかなタスクは単に「大きな関数を細かく切り分ける」ことではなく、隠れた制御ロジックをグラフ レベルまで引き上げて、結果として得られるアーキテクチャがより明確になり、さらなる開発に適したものになるようにすることです。
計画ノードが実際に行ったこと ¶
このグラフが説明しようとしていたエージェントは、大まかに言って、下流の計算のためにさまざまなパラメータを収集する業務に携わっていました。これらのパラメータの一部は、インターネット上で巧妙に検索されました。いくつかはユーザーに尋ねました。そして、特定の会話のコンテキストに応じて、同じパラメータを取得する適切な方法が大幅に異なる可能性があるため、これらすべてを完全に決定的ではないアルゴリズムによって実行しました。以下は、プラン ノードにパックされていた実際の関数のセットです。
そのため、私たちは最近、多くの人がよく認識している状況に陥っています。ムッシュ・クロードはスパゲッティから小さな傑作を私たちに彫刻してくれました。スパゲッティを徹底的に調べるのは、まったくレビューせずに次から次へと特集を追加することほど楽しいことではありません。しかし、これは、エントロピーを増大させたのと同じツールを使用してエントロピーを低減できることに気づくまでの話です。そして、それははるかに楽しい見通しです。
しかし、半分のチャンスさえあれば大喜びでコードを解くモデルそのものに、コードのもつれが解けるということを信頼できるでしょうか?
それに答えるために、さまざまなモデルからいくつかの独立した建築提案を収集することにしました。

、そして彼らがアドバイスするものを比較してください。
11人のモデルが審査のためにランウェイに招待されました。
まず、それぞれが分割案を独自に提案しました。次に、モデルは評価者モードに切り替わり、完成した提案のセット全体を読み取り、それらをランク付けしました。
1 つの幸運なテキストを繰り返し語るのではなく、独立した意見を確実に収集するために、次の条件が強制されました。
提案が生成されている間、モデルはお互いの作品を見ることができませんでした。
分析の生成中、彼らはすべての提案を確認しましたが、他の分析は確認しませんでした。
各実行は新しいセッションで行われました。
すべての作業は、Oh My Openagent プラグインを使用して OpenCode で実行され、モデルごとに最大の推論努力が行われました。
ステージ1。モデルは提案を生成します ¶
最初の段階では、各モデルが計画ノードのロジックをグラフ レベルまで引き上げる独自の方法を提案しました。各プロポーザルの生成に使用されるプロンプト (出力ファイル名のみが変更されました):
docs/planner-graph-ref/current-graph.md を参照してください。 「plan」ノードに含まれるロジックが多すぎるようです。このロジックを <model>-proposal.md でグラフ レベルに移動する方法を提案します。
提案テーブル ¶
フローチャート TD
planner_start([START]) --> ティック[ティック]
ティック -->|中止 / max_iters|仕上げ[仕上げ]
チェックマーク -->|地域または通貨がありません| ask_user[ask_user / 割り込み]
チェックを入れます -->|それ以外|準備する[準備する]
準備 --> 選択[選択]
select -->|決定的な決定が見つかりました|ガード[ガード]
選択 -->|決定なし|決定[決定 / LLM]
決める→守る
ガード -->|検索|検索[検索]
ガード -->|ask_user| ask_user
ガード -->|反射|反射する[反射する]
ガード -->|計算|計算する[計算する]
ガード -->|フィニッシュ|仕上げる
検索 -->|最新の観測結果|観察する[観察する]
検索 -->|ヒットまたはバックエンド障害がありません|カチカチ
観察 --> チェックを入れる
およそ

計算 --> チェック
ask_user --> 観察ユーザー[観察ユーザー]
観察ユーザー --> チェックを入れる
反映 --> チェック
終了 --> planner_end([END])
Fable-5 は、隠しロジックを 5 つのステージに分散することを提案しました。 tiny は反復の開始を引き受けます: ステップカウンター、status == "aborted" stop、max_iters 制限、そして地域と通貨に関する最初の質問。
準備層は prepare に入ります。プロアクティブな分解、後のデータ収集のためのレシピの収集、および compose_ready_fields の呼び出しです。
次に、select は、メインのモデル呼び出しを行わずに、決定論的な選択のチェーンを実行します。ここは、計算機がチェックする場所であり、収集済みのデータの高速パス、自動終了、およびブロックされた計算に続くデータ収集ブランチも同様です。
select で決定的な答えが見つからない場合、制御は Decide に渡され、モデルを呼び出して構造化された決定を構築するだけです。
次に、guard は、派生フィールドの決定のリダイレクト、オンラインでより適切に検索できるフィールドを検索するための強制切り替え、繰り返しアクションの制限、 search 、 ask_user 、reflect 、 Calculate 、またはfinish への最終ルーティング前の調整を 1 か所に収集します。以前は、これらすべての書き換えと制限は計画の末尾に塗りつぶされていました。アクション ノードはプロセス中に変更されません。それらは単に古い計画ノードに戻るのではなく、ティックに戻るだけです。
この分割により、アーキテクチャがより観察しやすくなります。グラフから、どこで決定的な決定がなされたか、どこで LLM が必要か、そして、決定が制限と修正の層を通過する場所はすでにわかります。最も重要な検出結果は、 Decision_origin です。これにより、共有ガードはどの決定が LLM から来たもので、どれが決定論的に見つかったのかを伝えることができるため、同じポリシーがすべてのブランチに無差別に適用されることはありません。で

sign には弱点もあります。すべての LLM 呼び出しが Decide に移動されるわけではなく、guard はかなり高密度のノードのままです。
フローチャート TD
開始 --> チェックを入れます
チェックを入れます -->|続行|ブートストラップゲート
チェックを入れます -->|ターミナル|仕上げる
bootstrap_gate -->|地域/通貨が必要| ask_user
bootstrap_gate -->|準備完了|コンテキストの準備
prepare_context --> 取得ゲート
取得ゲート -->|決定的な決定|意思決定方針
取得ゲート -->|LLM が必要| llm_plan
llm_plan --> 決定_ポリシー
意思決定ポリシー -->|検索|検索
意思決定ポリシー -->|ask_user| ask_user
意思決定方針 -->|反映|反映する
意思決定ポリシー -->|計算|計算する
意思決定ポリシー -->|終了|仕上げる
検索 -->|last_observation|観察する
検索 -->|観察なし|カチカチ
観察 --> チェックを入れる
計算 --> チェック
反映 --> チェック
ask_user --> 観察ユーザー
観察ユーザー --> チェックを入れる
終了 --> 終了
GPT-5.4 は、Fable-5 とほぼ同じように配置されています。また、ループエントリノード、コンテキストの準備、モデルの前の決定論的フォーク、別のモデル呼び出し、その後の単一の意思決定層もあります。
主な違いは、最初の地域と通貨の質問が check 内に留まらず、別の bootstrap_gate に取り込まれることです。そのため、tick はステップの開始、反復カウンター、および停止チェックのみを担当しますが、 bootstrap_gate は続行できるかどうか、または最初にユーザーを ask_user に送信する必要があるかどうかを決定します。
残りのステージは、実質的には Fable-5 とほとんど一致しますが、より明示的に名前が付けられているだけです。
GPT-5.4 は明確な作業指針を提供します。アンチパターンに関するセクションにより、目標を正確に把握してリファクタリングを適切に実行できるようになります。
フローチャート TD
planner_start([START]) --> enter_iteration
enter_iteration -->|中止または max_iters|仕上げる
enter_iteration --> prepare_schema
prepare_schema -->|リージョンがありません|必要な領域
準備_

スキーマ -->|通貨がありません|必要な通貨
prepare_schema -->|計算機の上限ヒット|仕上げる
prepare_schema -->|現在の計算機|仕上げる
準備スキーマ --> 取得ゲート
require_region --> ask_user
require_currency --> ask_user
取得ゲート -->|コンポーネント タスク|正規化_決定
acquisition_gate -->|すべての生の入力が準備完了|たぶん_計算する
取得ゲート -->|LLM が必要|計画
計画 --> 正規化_決定
正規化_決定 --> 再試行ゲート
retry_gate --> might_calculate
might_calculate -->|計算機が必要です|計算する
たぶん_計算 -->|検索|検索
might_calculate -->|ask_user| ask_user
might_calculate -->|反映|反映する
たぶん_計算 -->|終了|仕上げる
検索 -->|last_observation|観察する
検索 -->|ヒットしません| Enter_iteration
観察 --> enter_iteration
計算 --> 反復入力
ask_user --> 観察ユーザー
observ_user --> enter_iteration
反映 --> Enter_iteration
終了 --> planner_end([END])
GPT-5.5 は、最初の 2 つよりもさらに進んで、ループのブックキーピング部分をより積極的に分割します。 enter_iteration はエントリ ノードのままです。カウンタをインクリメントし、停止をチェックします。次に、 prepare_schema は状態の準備を収集し、地域と通貨を個別にチェックし、計算機の制限の早期終了も保持します。

[切り捨てられた]

## Original Extract

Weekly Tracker Feed by Korridzy.

Twilight of the Gods. Fable and 10 more LLMs on a Code Reorganization Task. Comparison. - wtf
Skip to content
wtf
Twilight of the Gods. Fable and 10 more LLMs on a Code Reorganization Task. Comparison.
English
Materials
Materials
Materials & raw data
Proposals
Proposals
Deepseek 4 Pro Proposal
Reviews
Reviews
Deepseek 4 Pro Range
Best analyst
Best analyst
Fable 5 Ba Range
Theses
Theses
Run 1
Run 1
Evaluator Common Thesis Ranking
Evaluator Preferability Ranking
Run 2
Run 2
Evaluator Consensus Ranking
Run 3
Run 3
Evaluator Consensus Ranking
Categories
Categories
AI agents
The original problem
What the plan node actually did
Stage one. The models generate proposals
The proposals table
A bit more on each proposal
Fable-5
Stage two. The models evaluate the proposals
The reviews table
A bit more on each review
Fable-5
Stage three. The wet swimsuit contest
Approach one. Do the scores agree?
Best proposals: the summary
Approach two. Comparing reviews by theses
Approach three. Center of opinion and medoid
Deus ex machina
Methods for evaluating the analysts
Back to index
Korridzy
Metadata
June 30, 2026
in
LangGraph ,
LLM ,
AI agents ,
Refactoring ,
Python
Twilight of the Gods. Fable and 10 more LLMs on a Code Reorganization Task. Comparison. ¶
Эта статья также доступна на русском: Гибель богов .
All 11 model proposals, the cross-reviews, the theses runs, and the ranking script are published here: Materials & reproduce this experiment .
This is a detailed write-up of one experiment. I took a god node from a real LangGraph agent and asked 5 American and 6 Chinese models first to propose how to untangle it, then to evaluate each other's proposals. After that, I tried three different ways to figure out which of them to trust on the matter.
The original problem What a god node is and why it's dangerous
What the plan node actually did
Lemonade from lemons Why all this, and how the experiment is set up
Stage one. The models generate proposals
Stage two. The models evaluate the proposals
Stage three. The wet swimsuit contest Deciding who's good at what
Approach one. Do the scores agree? Picking the best proposal
Approach two. Comparing reviews by theses Picking the best analyst
Approach three. Center of opinion and medoid Picking the best analyst again
Deus ex machina . Picking the best analyst one more time
Takeaways Which model to use as a generator, which as an evaluator, and where your heart will find peace.
You know how it goes: you're building a practice AI agent with the fellas on a course by Data Sanity , and amid the colorful whirl of rapidly accreting features you suddenly notice that one of the project's internal agents has a state graph (LangGraph) that looks like this:
flowchart TD
planner_start([START]) --> plan[plan]
plan -->|search| search[search]
plan -->|ask_user| ask_user[ask_user / interrupt]
plan -->|reflect| reflect[reflect]
plan -->|calculate| calculate[calculate]
plan -->|finish| finish[finish]
search -->|last_observation| observe[observe]
search -->|no hits / backend failure| plan
observe --> plan
calculate --> plan
ask_user --> observe_user[observe_user]
observe_user --> plan
reflect --> plan
finish --> planner_end([END])
At first glance this is just a cute little octopus — nothing to worry about. But once you know how much logic this octopus has to hold in its modest eight-legged head, it becomes clear right away that we're looking at an anti-pattern. In this case, let's call it a god node.
The plan node hides about 350 lines of logic, including iterative checks, bootstrap questions about region and currency, schema preparation, acquisition-task routing, the LLM call, the subsequent correction of the decision, and so on.
The problem isn't just the size of the function. When important orchestration is hidden inside a single node, the graph stops being a representation of the system. It's harder to explain, harder to debug, harder to test, and more dangerous to change. So the obvious task isn't merely to "chop a big function into pieces" but to lift the hidden control logic up to the graph level, so that the resulting architecture becomes clearer and more amenable to further development.
What the plan node actually did ¶
The agent this graph was meant to describe was, broadly, in the business of collecting various parameters for downstream calculations. Some of these parameters it cleverly searched for on the internet; some it asked the user about. And it did all this by a not-fully-deterministic algorithm, because depending on the context of a particular conversation, the right way to obtain the same parameter could vary considerably. Here is the set of real functions that had been packed into the plan node:
So, we find ourselves in a situation many people will recognize all too well these days. Monsieur Claude has sculpted us a little masterpiece out of spaghetti. Combing out that spaghetti is nowhere near as much fun as bolting on feature after feature with no review whatsoever. But that's only true until you realize you can reduce the entropy with the same tool that raised it. And that's a far more pleasant prospect.
But can you trust the untangling of code to the very model that, given half a chance, so gleefully tangles it up?
To answer that, I decided to collect several independent architectural proposals from different models, and compare what they'd advise.
Eleven models were invited onto the runway for judging:
First, each of them made its own proposal for splitting plan . Then the models switched into evaluator mode, read the entire set of finished proposals, and ranked them.
To make sure I was collecting independent opinions rather than a relay of retellings of one lucky text, the following conditions were enforced:
While the proposals were being generated, the models couldn't see one another's work.
While the analyses were being generated, they saw every proposal but none of the other analyses.
Each run took place in a fresh session.
All of the work was done in OpenCode with the Oh My Openagent plugin, at maximum reasoning effort for every model.
Stage one. The models generate proposals ¶
In the first stage, each model proposed its own way to lift the plan node's logic up to the graph level. The prompt used to generate each proposal (only the output filename changed):
look at docs/planner-graph-ref/current-graph.md. Looks like "plan" node contains too many logic in it. give a proposal of how to move this logic to graph level in <model>-proposal.md
The proposals table ¶
flowchart TD
planner_start([START]) --> tick[tick]
tick -->|aborted / max_iters| finish[finish]
tick -->|region or currency missing| ask_user[ask_user / interrupt]
tick -->|otherwise| prepare[prepare]
prepare --> select[select]
select -->|deterministic decision found| guard[guard]
select -->|no decision| decide[decide / LLM]
decide --> guard
guard -->|search| search[search]
guard -->|ask_user| ask_user
guard -->|reflect| reflect[reflect]
guard -->|calculate| calculate[calculate]
guard -->|finish| finish
search -->|last_observation present| observe[observe]
search -->|no hits or backend failure| tick
observe --> tick
calculate --> tick
ask_user --> observe_user[observe_user]
observe_user --> tick
reflect --> tick
finish --> planner_end([END])
Fable-5 proposed spreading the hidden logic across five stages. tick takes on the start of the iteration: the step counter, the status == "aborted" stop, the max_iters limit, and the initial questions about region and currency .
The preparatory layer goes into prepare : proactive decomposition, gathering recipes for later data collection, and the compose_ready_fields call.
Next, select runs the chain of deterministic choices without the main model call. This is where the calculator checks land, along with the fast pass over already-collected data, auto-finish, and the data-collection branches that follow a blocked calculation.
If select finds no deterministic answer, control passes to decide , which only calls the model and builds a structured decision.
Then guard collects in one place the redirection of decisions for derived fields, the forced switch to search for fields better looked up online, the limits on repeated actions, and the adjustment before the final routing to search , ask_user , reflect , calculate , or finish . Previously all these rewrites and limits were smeared across the tail of plan . The action nodes don't change in the process; they simply return not to the old plan node but back to tick .
This split makes the architecture more observable. From the graph you can already tell where a deterministic decision is made, where the LLM is needed, and where a decision passes through the layer of limits and fixes. The most important find is decision_origin . It lets the shared guard tell which decision came from the LLM and which was found deterministically, so it doesn't apply the same policy to every branch indiscriminately. The design has weak spots too: not all LLM calls are moved into decide , and guard remains a fairly dense node.
flowchart TD
START --> tick
tick -->|continue| bootstrap_gate
tick -->|terminal| finish
bootstrap_gate -->|needs region/currency| ask_user
bootstrap_gate -->|ready| prepare_context
prepare_context --> acquisition_gate
acquisition_gate -->|deterministic decision| decision_policy
acquisition_gate -->|needs LLM| llm_plan
llm_plan --> decision_policy
decision_policy -->|search| search
decision_policy -->|ask_user| ask_user
decision_policy -->|reflect| reflect
decision_policy -->|calculate| calculate
decision_policy -->|finish| finish
search -->|last_observation| observe
search -->|no observation| tick
observe --> tick
calculate --> tick
reflect --> tick
ask_user --> observe_user
observe_user --> tick
finish --> END
GPT-5.4 is arranged almost exactly like Fable-5. It also has a loop-entry node, context preparation, a deterministic fork before the model, a separate model call, and a single decision-fixing layer after it.
The key difference is that the initial region and currency questions don't stay inside tick but are pulled into a separate bootstrap_gate . So tick is responsible only for the start of the step, the iteration counter, and the stop checks, while bootstrap_gate decides whether you can proceed or must first send the user to ask_user .
The remaining stages mostly match Fable-5 in substance, just named more explicitly.
GPT-5.4 gives clear working guidance. Its section on anti-patterns makes it possible to carry out the refactor well, with a real grasp of the goals.
flowchart TD
planner_start([START]) --> enter_iteration
enter_iteration -->|aborted or max_iters| finish
enter_iteration --> prepare_schema
prepare_schema -->|region missing| require_region
prepare_schema -->|currency missing| require_currency
prepare_schema -->|calculator cap hit| finish
prepare_schema -->|calculator current| finish
prepare_schema --> acquisition_gate
require_region --> ask_user
require_currency --> ask_user
acquisition_gate -->|component task| normalize_decision
acquisition_gate -->|all raw inputs ready| maybe_calculate
acquisition_gate -->|LLM needed| plan
plan --> normalize_decision
normalize_decision --> retry_gate
retry_gate --> maybe_calculate
maybe_calculate -->|calculator required| calculate
maybe_calculate -->|search| search
maybe_calculate -->|ask_user| ask_user
maybe_calculate -->|reflect| reflect
maybe_calculate -->|finish| finish
search -->|last_observation| observe
search -->|no hits| enter_iteration
observe --> enter_iteration
calculate --> enter_iteration
ask_user --> observe_user
observe_user --> enter_iteration
reflect --> enter_iteration
finish --> planner_end([END])
GPT-5.5 goes further than the first two and splits the loop's bookkeeping part more aggressively. enter_iteration remains the entry node: it increments the counter and checks for a stop. Then prepare_schema gathers the state preparation, separately checks for region and currency , and also holds the early exits for the calculator limit a

[truncated]
