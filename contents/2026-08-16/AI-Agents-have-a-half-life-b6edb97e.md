---
source: "https://construct.computer/blog/agent-task-half-life/"
hn_url: "https://news.ycombinator.com/item?id=49317630"
title: "AI Agents have a half-life"
article_title: "Why AI Agents Keep Failing on Long Multi-Step Tasks"
author: "ankushKun"
captured_at: "2026-08-16T07:23:07Z"
capture_tool: "hn-digest"
hn_id: 49317630
score: 1
comments: 0
posted_at: "2026-08-16T07:12:58Z"
tags:
  - hacker-news
  - translated
---

# AI Agents have a half-life

- HN: [49317630](https://news.ycombinator.com/item?id=49317630)
- Source: [construct.computer](https://construct.computer/blog/agent-task-half-life/)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T07:12:58Z

## Translation

タイトル: AI エージェントには半減期がある
記事のタイトル: なぜ AI エージェントは長い複数ステップのタスクで失敗し続けるのか
説明: AI エージェントが長い複数ステップのジョブで失敗し続ける理由: 95% 信頼できるエージェントは、48 ステップを 8.5% の確率で完了します。この修正は再開可能な実行であり、より良いモデルではありません。

記事本文:
AI エージェントが長い複数ステップのタスクで失敗し続ける理由 メインコンテンツにスキップ コンピューターの価格設定を構築する
Construct、AI エージェント、開発者ツールについて書いています。
2026 年 8 月 11 日公開 · @ankushKun_
Valve は、目に見えないところに隠れている物理学のジョークにちなんで、ゲームを「Half Life」と名付けました。「半分がなくなるまでに何かがどれだけ続くか」というものです。あなたのエージェントはそのうちの 1 つを持っています。ただ、防護服やバールが取引から除外されるわけではない。
個々の手順の 95% を正しく実行しているエージェントを例に挙げます。それは良いエージェントです。次に、10 ステップのジョブを実行します。
10回中6回で終わります。残りの 4 回は途中で失敗します。文字通り。半生。
何も壊れませんでした。ステップは後退しませんでした。 0.95 の 10 乗は 0.60、これですべてです。ステップごとに優れているように見える信頼性は、ジョブごとに平凡であり、何かをスケジュールして立ち去るまで、ほとんどの人が感じない方法で悪化します。
エージェントが複数のステップのジョブの途中で失敗し続ける理由を尋ねた場合、通常はこれが答えであり、人々が探している答えではありません。
エージェントが長時間のタスクで失敗し続ける理由
長いタスクはより困難なタスクである、というのが魅力的な説明です。長くなると、曖昧さが増し、目標を読み間違える可能性が高まり、混乱がさらに大きくなります。
その説明は間違っていますが、その理由を示す明確な作品があります。 「AI エージェントの成功率には半減期はありますか?」の中で、Toby Ord は、タスクの長さにわたるエージェントのパフォーマンスが非常に単純なモデル、つまり人間がタスクを実行するのにかかる 1 分ごとの一定の失敗率によって説明されることを示しています (Ord、arXiv:2505.05115)。混乱を蓄積させない。ジョブの実行中ずっと一定のハザード率が発生します。
物理学者はすでにその速度に関する報告書を持っています。それは λ、N(t) = N₀・e^(−λt) の減衰定数、およびその半分です。

寿命はちょうど t1/2 = ln(2) / λ です。 Valve は同じ理由で、ゴードン フリーマンのスーツと Half-Life のロゴに同じ λ を付けました。それは、何かが消滅するまでの期間を決定する定数です。 Ord の結果、あなたのエージェントも持っているということになりました。
これにより、すべてのエージェントに半減期、つまり成功確率が 50% に達するタスクの期間が与えられます。 Ord の枠組みでは、より長いタスクは失敗します。これは、タスクに含まれるサブタスクのセットがますます大きくなり、どれか 1 つが失敗すると全体が失敗するためです。 λを上げる（毎分の信頼性が悪化する）か、曝露量を増やす（作業が長くなる）と、同位体のサンプルと同じように残存率が下がります。これにはバールは必要ありません。露出を減らす必要があります。
これは問題を有益に再構成します。失敗が困難に関するものであれば、よりスマートなモデルで問題を解決するでしょう。失敗が露出単位あたりの割合である場合、露出を減らすことで修正します。これらはまったく異なるエンジニアリング プログラムであり、ほぼ全員が最初のプログラムを実行しています。
マルチステップのエージェントワークフローが失敗する理由: 48 ステップ、8.5% の成功
十段はおもちゃです。実際の定期的なビジネス作業はループするため、さらに悪いものになります。
これは、小規模な代理店から常に聞かれる仕事です。それは、月次のクライアント報告です。クライアントは8名。それぞれについて、分析を取得し、広告支出を取得し、CRM の取引の動きを取得し、概要を作成し、クライアントのテンプレートにレンダリングして電子メールで送信します。 6 ステップ、8 回。四十八歩。
95% のエージェント、つまり本当に有能なエージェントは、そのジョブの開始から終了までを 12 回に 1 回程度完了します。そして、障害モードは醜いものです。クライアント 6 で停止し、すでに 5 件のレポートを電子メールで送信しています。トランスクリプト全体を読まないと、何がどのような状態にあるのかわかりません。 1 つの継続的な実験が、予算外の混乱に連鎖していきます。ブラック メサ エネルギー、スプレッドシート版。
それで、みんな

つまり、エージェントは実際の操作では機能しないと言えます。実際にうまくいかないのは、48 ステップを 1 つの中断のない賭けとして実行することです。
より良いモデルはエージェントの信頼性を修正しますか?
本物の反論があります。モデルは長時間のタスクにおいて急速に改良されています。 METR は、人間の専門家が同じ作業に必要な時間を使用して、エージェントが半分の時間で成功すると予測されるタスクの継続時間である 50% の期間を測定します。この期間は 6 年間でおよそ 7 か月ごとに 2 倍になり、最近のデータではより速くなっていることが示唆されています (METR、長いタスクを完了する AI 能力の測定、METR、時間範囲)。
それは本当の進歩であり、減速することはありません。ビジネスプロセスを計画することも大変なことです。ホライズンは信頼性 50% で定義されており、これはコイン投げに相当し、倍増曲線は来週火曜日の顧客レポートについては何も示しません。 Gartner は、コストの高騰、不透明なビジネス価値、不適切なリスク管理を理由に、エージェント AI プロジェクトの 40% 以上が 2027 年末までにキャンセルされると予測しています (Gartner、2025 年 6 月)。それらのプロジェクトの大部分は、まさにこれによって消滅することになります。1 度だけ機能したデモは毎週スケジュールされ、ほとんどの週は静かに失敗します。 cron で配信される予期せぬ結果。
モデルを待つのは納期のない戦略です。ランの短縮が可能になりました。
途中で失敗したエージェントを修正する方法: 継ぎ目でジョブをカットする
これら 48 のステップを 8 つの独立した 6 ステップの実行と算術反転に分割します。
ステップあたり 95% の場合、1 人のクライアントの 6 ステップの実行は 73.5% の確率で終了します。したがって、8 つのクライアントを通過すると、そのうちのおよそ 6 つが成功し、失敗した 2 つは名前が付けられ、分離され、再試行可能になります。再試行を実行すると、作業の 93% が完了します。残りのギャップは、特定のクライアントではなく 2 つの特定のクライアントを確認できることです。

不透明な死んだ仕事は一つもありません。
同じモデルです。ステップごとの信頼性は同じです。総作業量は同じです。唯一の変更は、実行を停止および再開できる場所です。 λを下げていません。 1 回の実行が存続するまでの時間を短縮しました。
問題は、これはピース間で状態が存続する場合にのみ機能するということです。独自のコンテキストを最初から再構築し、クライアントのテンプレートがどのようになるかを再取得し、すでに取得したものを再取得する必要がある 6 つのステップの実行は短縮されていません。短くしてからパッドを詰めて戻しています。チェックポイントが低コストになるのは、チェックポイントが耐久性のある場所に書き込まれている場合のみです。
『Half-Life 2』では、レジスタンスが壁に「λ」をスプレーペイントして、後で見つけることができる補給場所をマークしました。エージェントにも同じ種類のマークが必要です。ターン終了時に蒸発するメモリではなく、「この反復はすでに終了しました」というディスク上のファイルです。これが、エージェントにコンテキスト ウィンドウの代わりにコンピューターを提供するための実際の議論です。
Construct が同じジョブを実行する方法
Construct のエージェントには永続的なワークスペースがあります。つまり、ターン終了後も存続するファイル システム、セッション間で保持される長期記憶、オンデマンドまたはカレンダーから実行される再利用可能なワークフローです。これら 3 つの要素によって、顧客レポートの仕事が 1 つの長い賭けから収束する仕事に変わります。 HEV スーツのコスプレではなく、実際の危険環境に対応するギアを使用します。実行を生き残らなければならない状態は、何かが始まる前にディスク上に存在します。
具体的には、レポート ジョブは 8 つではなく 1 つのスケジュールされたワークフローになります。
重要な行は 3 行目です。命令はディスク上のレポートを持たずにクライアントに限定されているため、クライアント 6 で実行が停止しても、巻き戻しが必要なほどの大惨事にはなりません。次の実行では、ディレクトリが読み取られ、5 つのレポートが表示され、残りの 3 つが処理されます。何も二度送信されることはありませんし、二度送信された人もいません

どこで止まったか覚えておいてください。
2 つの正直な境界線。現在、Construct のワークフローは直線的です。分岐、ファンアウト、サブワークフローはサポートされていないため、クライアントは 8 つの並列ジョブとしてディスパッチされるのではなく、1 回の実行内で処理されます。また、スケジュール間隔は再試行の頻度であるため、1 時間以内に終了する必要があるジョブには、忍耐ではなくオンデマンドでの再実行が必要です。ユーザーごとのファイルシステムを実行間で放置できるほど安価にするインフラストラクチャは、エージェントがほとんど料金を支払わない実際のコンピュータをどのように入手するかにあります。
再開可能なエージェント ジョブを作成するためのチェックリスト
このパターンは、過去のクライアントのレポートを一般化したものです。 「X ごとに、いくつかのことを行う」という形のものが候補になります。
ループを見つけてください。ほぼすべての定期的な運用ジョブには、クライアントごと、請求書ごと、候補者ごと、リポジトリごとに 1 つあります。そのループの境界が継ぎ目です。まずそこをカットします。これは、1 つのロング ベットを多くのショート ベットに変えるカットだからです。
各反復でアーティファクトを残すようにします。ファイル、CRM レコード、送信された電子メール。後で実行すると存在を確認できるものがあります。反復によって何も残らない場合、再試行と重複を区別することはできません。壁に「λ」を残します。
反復では完了した作業をスキップします。 「今月日付のレポートをまだ作成していないクライアントに対してレポートを作成する」は再開可能です。 「レポートを書く」ではありません。通常、この 1 行が、自分でスケジュールを設定できる仕事と、ベビーシッターをしなければならない仕事の違いになります。
曖昧な判断は独自のステップに置きます。判定ステップは、機械的なステップに比べてステップごとの成功率が低くなります。それらを孤立させるということは、間違った判断により、逃げるのではなく、そのステップが犠牲になることを意味します。
スケジュールを設定する前に、最初の数回の実行をオンデマンドで監視します。どのステップが弱点であるかを探していますが、それを正しく推測することはできません。
何かが失敗した場合は、軌跡を確認してください。定数

ruct のアクティビティ フィードは、制限されたアクションの概要とベストエフォートの理由を保持し、チャットは制限されたツールのレコードを保持するため、失敗した反復は雰囲気ではなくステップまで追跡できます。
3 番目の点については、検討する価値があります。ほとんどの人は、エージェントの指示を仕事の説明として書きます。スケジューリングに残った指示は、残った作業の説明として書かれます。
チェックポイント設定で修正できないもの
チェックポイントを使用すると、失敗のコストが削減されます。失敗率が下がるわけではありませんし、その区別が重要となる仕事もあります。同じ λ がまだあります。サンプル全体を 1 回の連続露光に賭けることを拒否するだけです。
あなたのステップが最後まで完全に順序に依存しており、自然なループや安全な停止ポイントがない場合、カットする継ぎ目はなく、1 つの長い賭けに戻ります。ステップに不可逆的な外部副作用、支払い、公開投稿、顧客へのメッセージがある場合、再試行は無料ではなく、冪等性を想定するのではなく設計する必要があります。 Construct を使用すると、実行中のタスクを検査したり、ターンの途中でタスクを中断したり、エージェントが提起した質問に答えたりすることができますが、現時点ではすべての外部副作用の前に必須の承認ゲートを挿入していないため、これらのステップを無人で実行する前に実際の監視が必要です。
記憶にも境界がある。 Construct はワークスペース ファイルと選択された長期メモリをセッション間で保持しますが、自動メモリは現在、アップロードされたファイル、プライベート アプリの結果、ターミナル出力、またはライブ ブラウザーの実行にインデックスを付けません。継ぎ目を越えて伝達する必要があるコンテキストは、トランスクリプトに一度出現したからといって存続すると仮定するのではなく、ファイルに書き込むか明示的に保存する必要があります。
そして、作品が真に決定論的であれば、途中で判断することなく、毎回同じトリガーと同じ変換が行われます。

つまり、ルールベースの自動化プラットフォームは、どのエージェントよりも安価に、より予測可能にそれを実行します。私たちは、AI エージェントと Zapier の自動化において、その境界線がどこに該当するかを説明しました。
「どのモデルを使用すべきか」ではなく、これを質問してください
「どのモデルを使用すべきか」は誰もが抱く疑問です。通常、これはバインディング制約ではありません。モデルを 1 段階高くすると、ステップごとの信頼性が数ポイント向上し、指数は変更されません。 48 ステップの実行を 8 つの 6 ステップの実行に再構成すると、指数が移動し、そこにレバレッジがかかります。
代わりに、このジョブがステップ 30 で失敗した場合、ディスクには何が残っているのか、再試行では何をやり直す必要があるのか​​を尋ねてください。答えが「何もない」または「すべて」である場合、モデルは問題ではありません。適切な人を間違った場所に配置する必要はありません。ターンを超えて存続するワークスペースが必要です。
まず、再利用可能なステップとスケジュールがどのように組み合わされるかについての AI ワークフローの自動化、実行間で引き継がれる内容とその修正方法についての AI エージェント メモリ、他のツールと比較して評価する場合の AI エージェント プラットフォームの選択方法について説明します。基礎となるアイデアが新しい場合は、AI 従業員とは何かということから始めます。
ガイド · 2026年7月20日 · ニシャル
再利用可能な線形 AI ワークフローを作成し、オンデマンドで実行したり、繰り返しのスケジュールを設定したりできます。

[切り捨てられた]

## Original Extract

Why AI agents keep failing on long multi-step jobs: a 95% reliable agent finishes 48 steps 8.5% of the time. The fix is a resumable run, not a better model.

Why AI Agents Keep Failing on Long Multi-Step Tasks Skip to main content Construct Computer Pricing
Writes about Construct, AI agents, and developer tools.
Published August 11, 2026 · @ankushKun_
Valve named a game Half-Life after the physics joke hiding in plain sight: how long something lasts before half of it is gone. Your agent has one of those. It just does not get a hazmat suit or a crowbar out of the deal.
Take an agent that gets 95% of its individual steps right. That is a good agent. Now give it a ten-step job.
It finishes six times out of ten. The other four times, it fails halfway. Literally. A half-life.
Nothing broke. No step regressed. 0.95 to the tenth power is 0.60, and that is the whole story. Reliability that looks excellent per step is mediocre per job, and it gets worse in a way most people do not feel until they schedule something and walk away.
If you have been asking why your agent keeps failing halfway through a multi-step job, this is usually the answer, and it is not the one people go looking for.
Why your agent keeps failing on long tasks
The tempting explanation is that long tasks are harder tasks. Longer means more ambiguity, more chances to misread the goal, more compounding confusion.
That explanation is wrong, and there is a clean piece of work showing why. In "Is there a half-life for the success rates of AI agents?", Toby Ord shows that agent performance across task lengths is explained by an extremely simple model: a constant rate of failing during each minute a human would take to do the task ( Ord, arXiv:2505.05115 ). Not accumulating confusion. A flat hazard rate, ticking the entire time the job runs.
Physicists already have a letter for that rate. It is λ, the decay constant in N(t) = N₀·e^(−λt), and the half-life is just t½ = ln(2) / λ. Valve put the same λ on Gordon Freeman's suit and in the Half-Life logo for the same reason: it is the constant that decides how long something lasts before it is gone. Ord's result is that your agent has one too.
That gives every agent a half-life: the task duration at which its success probability hits 50%. Longer tasks fail, in Ord's framing, because they contain increasingly large sets of subtasks where failing any one fails the whole thing. Raise λ (worse per-minute reliability) or raise exposure (longer jobs), and the surviving fraction falls the same way a sample of isotope does. You do not need a crowbar for this. You need less exposure.
This reframes the problem usefully. If failure were about difficulty, you would fix it with a smarter model. If failure is a rate per unit of exposure, you fix it by reducing exposure. Those are very different engineering programs, and almost everyone is running the first one.
Why multi-step agent workflows fail: 48 steps, 8.5% success
Ten steps is a toy. Real recurring business work is worse, because it loops.
Here is a job we hear about constantly from small agencies: monthly client reporting. Eight clients. For each one, pull analytics, pull ad spend, pull the CRM's deal movement, write the summary, render it into the client's template, email it. Six steps, eight times. Forty-eight steps.
A 95% agent, which is a genuinely capable agent, completes that job start to finish about one time in twelve. And the failure mode is the ugly one: it dies at client six, having already emailed five reports, and you cannot tell what state anything is in without reading the whole transcript. One continuous experiment, cascading into a mess you did not budget for. Black Mesa energy, spreadsheet edition.
So people conclude agents do not work for real operations. What actually does not work is running forty-eight steps as one uninterrupted bet.
Will a better model fix agent reliability?
There is a genuine counterargument: models are getting better at long tasks fast. METR measures the 50% time horizon, the task duration at which an agent is predicted to succeed half the time, using how long human experts need for the same work. That horizon has been doubling roughly every seven months for six years, with recent data suggesting faster ( METR, measuring AI ability to complete long tasks ; METR, time horizons ).
That is real progress and it is not slowing. It is also a terrible thing to plan a business process around. The horizon is defined at 50% reliability, which is a coin flip, and the doubling curve tells you nothing about next Tuesday's client reports. Gartner predicts that over 40% of agentic AI projects will be canceled by the end of 2027, citing escalating costs, unclear business value, and inadequate risk controls ( Gartner, June 2025 ). A large share of those projects will die from exactly this: a demo that worked once, scheduled weekly, quietly failing most weeks. Unforeseen consequences, delivered on a cron.
Waiting for the model is a strategy with no delivery date. Shortening the run is available now.
How to fix an agent that fails partway: cut the job at the seams
Split those forty-eight steps into eight independent six-step runs and the arithmetic inverts.
At 95% per step, one client's six-step run finishes 73.5% of the time. So a pass over eight clients lands roughly six of them, and the two that failed are named, isolated, and retryable. Run the retry and you are at 93% of the work done, with the remaining gap being two specific clients you can look at rather than one opaque dead job.
Same model. Same per-step reliability. Same total work. The only change is where the run is allowed to stop and be resumed. You have not lowered λ. You have lowered how long any one run has to survive it.
The catch is that this only works if state survives between the pieces. A six-step run that has to rebuild its own context from scratch, re-derive what the client's template looks like, re-fetch what it already fetched, has not been shortened. It has been shortened and then padded back out. Checkpointing is only cheap when the checkpoint is written somewhere durable.
In Half-Life 2 , the Resistance spray-painted λ on walls to mark supply caches you could find later. Your agent needs the same kind of mark: a file on disk that says "this iteration already finished," not a memory that evaporates when the turn ends. That is the actual argument for giving an agent a computer instead of a context window.
How Construct runs the same job
Construct's agent has a persistent workspace: a filesystem that outlives the turn, long-term memories it carries across sessions, and reusable workflows that run on demand or from its Calendar. Those three pieces are what turn the client-reporting job from one long bet into a job that converges. Less HEV suit cosplay, more actual hazardous-environment gear: the state that has to survive the run is sitting on disk before anything else starts.
Concretely, the reporting job becomes one scheduled workflow rather than eight:
The important line is the third one. Because the instruction is scoped to clients without a report on disk, a run that dies at client six is not a disaster that needs unwinding. The next run reads the directory, sees five reports, and works on the remaining three. Nothing is sent twice, and nobody has to remember where it stopped.
Two honest boundaries. Construct's workflows are linear today: branching, fan-out, and subworkflows are not supported, so the clients are worked through inside a run rather than dispatched as eight parallel jobs. And the schedule interval is your retry cadence, so a job that must finish within the hour needs an on-demand rerun rather than patience. The infrastructure that makes a per-user filesystem cheap enough to leave lying around between runs is in how our agents get real computers we mostly do not pay for .
A checklist for building a resumable agent job
The pattern generalizes past client reports. Anything shaped as "for each X, do a few things" is a candidate.
Find the loop. Almost every recurring ops job has one: per client, per invoice, per candidate, per repo. That loop boundary is your seam. Cut there first, because it is the cut that turns one long bet into many short ones.
Make each iteration leave an artifact. A file, a CRM record, a sent email. Something a later run can check for existence. If an iteration leaves nothing behind, you cannot tell a retry from a duplicate. Leave a λ on the wall.
Make iterations skip completed work. "Write the report for any client that does not already have one dated this month" is resumable. "Write the reports" is not. This one line is usually the difference between a job you can schedule and a job you have to babysit.
Put the ambiguous judgment in its own step. Judgment steps have a lower per-step success rate than mechanical ones. Isolating them means a bad judgment call costs you that step, not the run.
Supervise the first several runs on demand before scheduling. You are looking for which step is your weak link, and you will not guess it correctly.
Check the trail when something fails. Construct's Activity feed keeps bounded action summaries and best-effort reasons, and chat retains bounded tool records, so a failed iteration can be traced to a step rather than to a vibe.
That third point is worth sitting with. Most people write agent instructions as a description of the work. The instructions that survive scheduling are written as a description of the work that remains.
What checkpointing does not fix
Checkpointing lowers the cost of failure. It does not lower the failure rate, and there are jobs where that distinction matters. You still have the same λ. You just refuse to stake the whole sample on one continuous exposure.
If your steps are genuinely order-dependent all the way through, with no natural loop and no safe stopping point, there is no seam to cut and you are back to one long bet. If a step has an irreversible external side effect, a payment, a public post, a message to a customer, then a retry is not free and idempotency has to be designed in rather than assumed. Construct lets you inspect a running task, interrupt it mid-turn, and answer a question the agent raises, but it does not currently insert a mandatory approval gate before every external side effect, so those steps still need real supervision before they run unattended.
Memory has a boundary too. Construct keeps workspace files and selected long-term memories across sessions, but automatic memory does not currently index uploaded files, private app results, terminal output, or live browser runs. Context that has to carry across a seam should be written to a file or stored explicitly, not assumed to survive because it appeared once in a transcript.
And if the work is genuinely deterministic, the same trigger and the same transformation every time with no judgment anywhere in the middle, a rule-based automation platform will do it more cheaply and more predictably than any agent. We laid out where that line falls in AI agent vs Zapier automation .
Ask this instead of "which model should I use"
"Which model should I use" is the question everyone asks. It is not usually the binding constraint. A model one tier better moves your per-step reliability a few points and leaves the exponent untouched. Restructuring a 48-step run into eight 6-step runs moves the exponent, which is where the leverage lives.
Ask instead: when this job fails at step thirty, what is left on disk, and what does the retry have to redo? If the answer is "nothing" and "all of it," the model is not your problem. You do not need the right man in the wrong place. You need a workspace that outlives the turn.
Start with AI workflow automation for how reusable steps and schedules fit together, AI agent memory for what carries across runs and how to correct it, and how to choose an AI agent platform if you are evaluating this against other tools. If the underlying idea is new, what is an AI employee is the place to begin.
guide · Jul 20, 2026 · Nischal
Create reusable linear AI workflows, run them on demand, or schedule recurring w

[truncated]
