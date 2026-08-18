---
source: "https://leviath.dev"
hn_url: "https://news.ycombinator.com/item?id=49345132"
title: "Show HN: Leviath – Structured context regions for LLM agents, in one Rust binary"
article_title: "Leviath, an agent runtime: one file, thousands of agents"
image: "https://leviath.dev/leviath-logo.png"
author: "gemisis"
captured_at: "2026-08-18T13:36:15Z"
capture_tool: "hn-digest"
hn_id: 49345132
score: 1
comments: 0
posted_at: "2026-08-18T13:12:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Leviath – Structured context regions for LLM agents, in one Rust binary

- HN: [49345132](https://news.ycombinator.com/item?id=49345132)
- Source: [leviath.dev](https://leviath.dev)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T13:12:01Z

## Translation

タイトル: Show HN: Leviath – 1 つの Rust バイナリでの LLM エージェントの構造化コンテキスト領域
記事のタイトル: Leviath、エージェント ランタイム: 1 つのファイル、数千のエージェント
説明: 1 つのファイルに記述するエージェント ランタイム: ステージごとのモデル、ツールセット、およびコンテキスト バジェット。1 つの 28 MB バイナリに数千のエージェントが含まれます。

記事本文:
Leviath、エージェント ランタイム: 1 つのファイル、数千のエージェント Leviath Home Docs The Lair Discord GitHub ☾ ☀ Leviath
1 つのファイルにはエージェントが記述されています。 1 つのバイナリで 1 万個のバイナリが実行されます。
エージェントごとではなく、ステージごとのモデル、ツールセット、コンテキスト バジェット。 28 MB の Rust バイナリが 1 つあり、他にインストールするものはありません。
1 つのプロセスに 10,000 以上のエージェント 何もない状態から実行中のエージェントまで 64 ミリ秒 30 ジョブ中 30 ジョブがデータが大きすぎて収まらない状態で終了しました これらの測定方法 ↓
macOS Linux Windowscurl -fsSL https://leviath.dev/install.sh | sh Copy macOS コマンド 何が問題になるのか
長期にわたるエージェントの実行は、同じ 5 か所でバラバラになります。
長い距離を走ると、自分が言ったことを忘れてしまいます
50 回のツール呼び出しで、開いた 1 つの制約がファイル ダンプとスタック トレースによって絞り出されています。
ウィンドウがいっぱいです · すべてを圧縮しています…
ウィンドウがいっぱいになると、トランスクリプト全体が 1 つの損失のある要約に押しつぶされます。詳細もそれに付随するため、エージェントはすでに読み取った内容を再読み込みします。
タスク、計画、コードベースは決して移動しません。会話のみが独自の履歴領域に圧縮され、ツール呼び出しはルールに従ってウィンドウ内をスライドします。
Grepping のコストは計画のコストと同じ
1 つのモデル、1 つのツールセット、および 1 つのプロンプトでジョブ全体が実行されるため、プレミアムを支払う理由はファイルをリストすることであり、5 番目の同一の編集が進行したかどうかを判断する唯一の理由となります。
同じプロンプト、同じモデル、すべてのツール
すべてのステップは最も難しいステップのレートで請求され、進行状況を評価する唯一のことはループを作成することです。
リストを発見して読み取る 安価なモデルを実装する コードを書き込む 高価なモデルをレビューする diff を読み取る 高価なモデルを再評価する 何が失敗したかを知る 高価なモデルのヒント ヒントがスタックする: 5 回の編集、1 つのファイルの再試行 各段階で独自のモデル、ツール、およびコンテキストの予算が選択されるため、コードを作成するステップのみが高価なコストを支払います。

ね。ランタイムは編集内容もカウントするため、同じファイルに対する 5 回目の試行は進行状況として報告されず、どこかにルーティングされます。
実行中にクラッシュすると、最初からやり直すか、2 回やり直すことになります
プロセス内で何時間ものツール呼び出しが行われるとプロセスが停止し、ディスク上に何も残らないため、どのプロセスがすでにファイルを書き込んだのか、移行を実行したのか、ブランチをプッシュしたのかを知ることができません。
今日: 記録はプロセスの中に存在します
4時間続いたが、すでに起こったことのリストは何もない。
そこで、選択します。もう一度実行してブランチを 2 回プッシュするか、リポジトリを開いて 4 つのうちのどれがランディングされたかを手動で判断します。
⋯リリースメールを添付して送信
3 回リプレイされましたが、再実行されません 1 回返されました: 繰り返す前に確認してください
すべての通話は発信時と着信時に記録されるため、終了した通話は繰り返されるのではなく再生され、まだ通話中の通話だけが質問されます。設定するものは何もありません。
少数のエージェントがマシンを埋め尽くします
実行ごとに独自のプロセス、独自のメモリ、およびランタイムの独自のコピーがもたらされるため、作業が完了するずっと前にボックスが使い果たされます。
エージェントが 3 人いるということは、すべてのコピーが 3 つあるということです。タスク リストが完了するずっと前に、マシンが停止します。
すべてのエージェントは 1 つの 28 MB Rust バイナリ内に存在し、一度に数千個存在します。
エージェントを中断することはできますが、そのサブエージェントは中断できません
3 つ下のレベルでは、従業員はあなただけが答えられる質問をし、代わりに推測します。提供される会話は最上位のエージェントとのみであるためです。
現在: トップエージェントのみがあなたの声を聞くことができます
サブエージェント: どのスキーマが正規ですか?誰も尋ねないから推測する
サブエージェント、深さ 3 レベル: どのスキーマが正規ですか?それはあなたに直接尋ねます
メッセージは、ターンの途中で話したかのように、推論呼び出しの間に到着します。正しい方向を向いて走りは続く。
あなたが書き込む 1 つのファイルと、決して触れない 2 つのファイル
ファイルはwhです

エージェントが何をするのかを説明する前に、そのステップ、各ステップでどのモデルを取得するか、いつキーボードを手渡すか、いつ停止して再考するか、どれだけのことを念頭に置いているかを説明します。ランタイムはそのファイルをグラフとして読み取ります。これは、 The Lair で実際の実行の動きを観察したり、 lev ダッシュ を使用してターミナルから観察したりするのと同じグラフです。
Leviath には 7 つのエージェントが同梱されているため、開始するためにエージェントを作成する必要はありません。以下のファイルはその 1 つです。
[stages.plan] ← ステージモード = "interactive_points" ← ライブステアリング model = { models = [{ Provider = "anthropic", model = "claude-sonnet-5" }] } available_tools = ["read_file", "ask_user_choice", "edit_document"] max_iterations = 20 [stages.implement] model = { models = [{ Provider = "anthropic", model = "claude-opus-5" }] } available_tools = ["write_file", "edit_file", "bash"] max_iterations = 50 [stages.implement.transitions.reassess] ← ステージ条件 = "stuck" stack_after_iterations = 20 stack_after_ minutes = 15 stack_after_same_file_edits = 5 ヒント = "前進なし - ステップバックし、 reassess" [context.regions] ← コンテキスト領域 task = { kind = "pinned", Budget = "2%" } plan = { kind = "pinned", Budget = "5%" } 会話 = { kind = "sliding_window", max_items = 40, Budget = "20%" } 長さを調整してトリミングされた実際の行。
発見 計画を立てる 承認プロトタイプを実装する レビューを実施する 行き詰まった場合にエージェントが選択したものを再評価する 自動 ▤ 並列ワーカー
3 つすべてが The Lair で、バイナリの一部ではなく、このサイトの一部です。これはブラウザ内で実行され、自分のマシン上のデーモンと通信するため、監視のためにコンピュータを離れて実行する必要はありません。
このファイルに見つからない 2 つは、誰もセットアップする必要のない 2 つです。すべてのエージェントは 1 つのプロセス内に存在し、1 台のマシン上に一度に数千人が存在し、すべてのステップは発生するたびに記録されます。

殺された国連は最初からやり直すのではなく拾います。
窓がなくなったらどうなるか
同じバイナリ、同じツール、同じ権限、同じイテレーション予算。唯一異なるのは、コンテキストが 1 つの成長ウィンドウであるか、最後に検証器を備えたステージのパイプラインであるかどうかです。
レポートを生成した実行
3 つのウィンドウ サイズで、ウィンドウより大きいログ コーパスから読み取られた 17 個の正確な数値の 1 つのレポート。ここでの単一モデル ループは強化されたものです。プロンプトでは、ウィンドウがいっぱいになる前に速度を下げ、結果をきれいに保つように指示されます。これは、パイプラインが強制するすべての規則であり、命令として記述されます。誰も報告書を返しませんでした。各サイズでの 10 個のプレーンなシングルループの実行も行わず、両方のバーがスケールを共有するように中断されました。失敗は間違った答えではなく、モデルがフルウィンドウに近づくと静かになります。
納品されたレポート内の数字の捏造
両方のバーは、同じウィンドウの同じタスクの同じパイプラインです。 2 つ目には、レポートが送信される前にすべての図のコーパスを再読み取る検証ステージがあります。これがなければ、5 件中 3 件のレポートが完全で適切にフォーマットされ、事前トレーニングから入力されて返されました。つまり、時間単位のタイムスタンプ、コーパスが意図的に避けている正規の綴りです。
注意点: 私たち自身のタスクは、通常の日をサンプリングするのではなく、ウィンドウに重点を置くように構築されています。それを一般的なスコアではなくエッジとして読んでください。それぞれ 1 ラウンド、フリーズ前に 5 回実行します。私たちは構造化エージェントを構築し、実行中に 2 回改善しました。タスクに合わせた回答ではなく、構造的な変更でした。やり直さなかったジョブの 1 つは、依然として失われ続けているジョブです。単一のループには同等のパスがありません。これは、1 つのウィンドウには欠陥ではなく発見を置く場所がないためです。エージェント、タスク、Runbook は両方とも公開されています: ベンチマーク ハーネス

。ラウンドがフリーズすると、それが何であろうと、完全なランツリーが公開されます。
問題は、どのエージェントが優れているかということではありません。それは作品が 1 つのウィンドウに収まるかどうかであり、Leviath はどちらかの答えになるように構築されています。
1 つのループを使用します すべてがウィンドウ内に収まります 材料、ツール、会話が作業スペースに適合する場合は、1 つのループが適切な形状です。より安く、より速く、そして私たち自身の測定では、少なくとも同等のスコア、場合によってはそれ以上のスコアを獲得しています。 Reach for nothing more.
構造を使用する マテリアルはウィンドウより大きい コーパスが収まらなくなると、ループは何を保持するかを選択しません - ウィンドウは、年齢によって異なります。ステージでは、各境界で何が生き残るかを決定できます。これは、余地がなくなったエージェントと、そのための計画が与えられたエージェントの違いです。
構造を使用する 間違った答えは何もないより悪い 成果物が正確である必要がある場合は、草稿と、それをソースと照合することだけが唯一の仕事である読者との間に段階を置きます。パイプラインではそれが必要になる場合があります。プロンプトはそれを要求することしかできません。
どちらのシェイプも同じバイナリであり、同じブループリント ファイルです。 1 つのループは、大きなウィンドウを持つ 1 つのステージです。パイプラインは 6 つあり、その間に予算があります。間違った選択は移行ではなく編集です。そのため、複雑なオプションを使用しない場合をお知らせする余裕があります。
1 万のエージェントとランタイムの邪魔にならない
これらは 1 台のマシン (16 コアを備えた Apple M3 Max) から生成され、1 つの 28 MB バイナリからフラット 1.5 秒で応答する模擬モデルに対して、バンドルされたエージェントの混合フリートを実行します。
メモリは RSS ではなくライブ メモリであり、2 つの数字のうち小さい方、つまりカーネルが再利用できるページがプロセスから出た後もプロセスが保持しているものです。 CPU はマシン全体のシェアなので、100% は 16 コアすべてになります。
モック化されたモデルがポイントです。

ランタイムにかかるコストは、ランタイム自体のレイテンシの下に埋められるでしょう。したがって、これはワークロードの予測ではなくランタイムが追加する下限であり、実際のモデルで最初にぶつかる壁はマシンではなくレート制限です。
ピーク メモリ (実行されていたエージェントの数)
推論プールは各層で 512 に保持されるため、変わるのはエージェントの数だけです。ランタイムの支払いは 1 回限りなので、エージェントを追加するにつれてエージェントあたりのコストが下がります。10 個の場合はそれぞれ 3.4 MB、1 万個の場合はそれぞれ 0.28 MB になります。その最上位層は、マシンが停止した場所ではなく、はしごが停止した場所です。128 GB のうち 2.8 GB、CPU のピークは 44% です。
同時に飛行中に走る
同じ 512 の固定プール。このフリートのエージェントはサブエージェントを開始するため、生成された数よりも多くの実行が発生します。ある瞬間に何回重なるかは別の数字であり、1万回になるとそれは低くなります。
寒くてどこにも温かいものがないから
デーモンの起動は、シナリオから差し引かれるのではなく、各シナリオ内で測定されるため、実行を開始する 64 ミリ秒には、起動にかかる 22 ミリ秒がすでに含まれています。
1,000 エージェント実行時の CPU
実行が実際にアクティブだったスパン全体で平均したため、両端のアイドル時間はそれを上回ることはできません。すべてのプール幅で 16 コアの 6% 未満です。
推論プール幅ごとに、1,000 人のエージェントが終了するまでの時間
エージェントが生成され、推論プールが 512 に保持される
ランニング全体とランニング後の記憶
3 GB 2 GB 1 GB 0 0 300 600 900 層が開始されてからの秒数
作業がある間はすべての層が上昇し、最後の実行が枯渇するとゼロに戻ります。漏れは二度と戻ってこないラインになります。
100% 50% 0 0 300 600 900 層が開始されてからの秒数
1 万のエージェントが 16 コアの約 40% でマシンを 11 分間保持し、より小さな層はほとんど軸をマークしません。残っているのはほとんどが費やした時間だワイ

モデルに取り付けます。
推論プール幅、1,000 エージェント生成
より広いプールにより、より多くの作業が実行可能になります
1.4 GB 700 MB 0 0 100 200 300 層が開始されてからの秒数
同じ 1,000 のエージェント、4 つのプール幅。プールが広いほど一度に多くのものが実行されるため、メモリ使用量は増加し、持続時間は大幅に短くなります。128 個は 5 分間にわたって 400 MB 未満に留まり、1,024 個は 1.2 GB を超えてピークに達し、90 秒で終了します。
20% 10% 0 0 100 200 300 層が開始されてからの秒数
軸に注意してください。これはマシンの 100% ではなく 20% で最高になります。高さではなく面積を読み取ります。使用される合計 CPU は、同じ 1,000 エージェントの場合、128 マシン秒の 14.5 マシン秒から 1,024 マシン秒の 2.9 マシン秒まで、各ステップで減少します。これは主に、より広いプールがより早く終了するためです。
すべてのグラフでは 3 つの繰り返しがすべて描画されており、その平均ではありません。そのため、各色が 3 本の線でほぼ重なっています。上のバーは 3 つのバーの中央値を示しているため、ここで最も高い線はバーのわずかに上を走っています。
比較に最も近いもの
他のフレームワークのベンチマークは行っていません。私たちがどのようなベースラインを書いたとしても、誰かが負けるためにそれを書いたと言うことは間違いないので、ここでは他の誰かが実行したベンチマークを示します。それは私たちのものとは異なるものを測定します、そして両方のセットは

[切り捨てられた]

## Original Extract

An agent runtime you describe in one file: a model, a toolset and a context budget per stage, with thousands of agents in one 28 MB binary.

Leviath, an agent runtime: one file, thousands of agents Leviath Home Docs The Lair Discord GitHub ☾ ☀ Leviath
One file describes the agent. One binary runs ten thousand of them.
A model, a toolset and a context budget per stage, not per agent. One 28 MB Rust binary, nothing else to install.
10,000+ agents in one process 64 ms from nothing to a running agent 30 of 30 jobs finished on data too big to fit How these were measured ↓
macOS Linux Windows curl -fsSL https://leviath.dev/install.sh | sh Copy macOS command What goes wrong
Long agent runs fall apart in the same five places .
Long runs forget what you told them
Fifty tool calls in, the one constraint you opened with has been squeezed out by file dumps and stack traces.
window full · compacting everything…
The window fills, then the whole transcript is squashed into one lossy summary. The details go with it, so the agent re-reads what it already read.
Task, plan, and codebase never move. Only the conversation compacts into its own history region, and tool calls slide through their window by rule.
Grepping costs what planning costs
One model, one toolset and one prompt run the whole job, so the reasoning you pay a premium for is also the thing listing your files, and the only thing judging whether a fifth identical edit was progress.
same prompt · same model · every tool
Every step is billed at the rate of the hardest one, and the only thing grading progress is the loop making it.
discover lists and reads cheap model implement writes the code expensive model review reads the diff expensive model reassess knows what failed expensive model hint hint stuck: 5 edits, one file retry Each stage picks its own model, tools and context budget, so only the step that writes code pays for the expensive one. The runtime counts the edits too, so a fifth attempt at the same file routes somewhere instead of being reported as progress.
A crash mid-run means starting over, or doing it twice
The process dies with hours of tool calls inside it, and nothing left on disk can tell you which of them already wrote a file, ran a migration, or pushed a branch.
Today: the record lives in the process
Four hours of it, and no list of what already happened.
So you choose: run it again and push the branch twice, or open the repo and work out by hand which of the four landed.
⋯ sending the release mail appended
3 replayed, not re-run 1 handed back: check before repeating
Every call is written down when it goes out and again when it lands, so the ones that finished are replayed instead of repeated, and the one still in the air is the only thing you are asked about. Nothing to configure.
A handful of agents fills the machine
Every run brings its own process, its own memory, and its own copy of the runtime, so the box runs out long before the work does.
Three agents means three copies of everything. Your machine taps out long before your task list does.
Every agent is an entity inside one 28 MB Rust binary, thousands of them at a time.
You can interrupt the agent, but not its sub-agents
Three levels down, a worker hits a question only you can answer and guesses instead, because the only conversation on offer is with the agent on top.
Today: only the top agent can hear you
sub-agent: which schema is canonical? no one to ask, so it guesses
sub-agent, three levels deep: which schema is canonical? it asks you directly
Messages land between inference calls, as if you had spoken mid-turn. The run keeps going, pointed the right way.
One file you write, and two things you never touch
The file is where you say what the agent does: its steps, which model each step gets, when it should hand you the keyboard, when it should stop and rethink, and how much it keeps in mind. The runtime reads that file as a graph, and it is the same graph you watch a real run move through in The Lair , or from your terminal with lev dash .
Leviath ships with 7 agents, so you never have to write one to start. The file below is one of them.
[stages.plan] ← Stages mode = "interactive_points" ← Live steering model = { models = [{ provider = "anthropic", model = "claude-sonnet-5" }] } available_tools = ["read_file", "ask_user_choice", "edit_document"] max_iterations = 20 [stages.implement] model = { models = [{ provider = "anthropic", model = "claude-opus-5" }] } available_tools = ["write_file", "edit_file", "bash"] max_iterations = 50 [stages.implement.transitions.reassess] ← Stages condition = "stuck" stuck_after_iterations = 20 stuck_after_minutes = 15 stuck_after_same_file_edits = 5 hint = "No forward progress - step back and reassess" [context.regions] ← Context regions task = { kind = "pinned", budget = "2%" } plan = { kind = "pinned", budget = "5%" } conversation = { kind = "sliding_window", max_items = 40, budget = "20%" } Real lines, trimmed for length.
discover plan your approval prototype implement review reassess agent-chosen when stuck automatic ▤ parallel workers
All three are The Lair , which is part of this site rather than part of the binary: it runs in your browser and talks to the daemon on your own machine, so no run has to leave your computer to be watched.
The two you will not find in that file are the two nobody should have to set up: every agent lives inside one process, thousands at a time on one machine, and every step is written down as it happens, so a run that gets killed picks up instead of starting over.
What happens when the window runs out
Same binary, same tools, same permissions, same iteration budget. The only thing that differs is whether the context is one growing window or a pipeline of stages with a verifier at the end.
Runs that produced the report at all
One report of 17 exact figures, read out of a log corpus larger than the window, at three window sizes. The single-model loops here are the hardened ones: they are told in the prompt to wind down before the window fills and keep their results clean, which is every discipline the pipeline enforces, written as instructions. None of them returned a report. Neither did the ten plain single-loop runs at each size, left off so both bars share a scale. The failure is not a wrong answer but the model going quiet near a full window.
Invented figures in the delivered reports
Both bars are the same pipeline on the same task at the same window; the second one has a verify stage that re-reads the corpus for every figure before the report ships. Without it, three of five reports came back complete, well formatted, and filled in from training priors: round-hour timestamps, canonical spellings the corpus deliberately avoids.
The caveats: our own tasks, built to stress a window rather than sample a normal day - read it as the edge, not a general score. Five runs of each, one round, pre-freeze. We built the structured agent and improved it twice while these ran: structural changes, not answers tuned to a task, and the one kind of job we never reworked is one we still lose. The single loop got no equivalent pass, because there is nowhere in one window to put one - the finding rather than a flaw in it. Both agents, the tasks and the runbooks are public: the benchmark harness . The full run tree publishes when the round freezes, whatever it says.
The question is not which agent is better. It is whether the work fits in one window, and Leviath is built to be either answer.
Use one loop It all fits in the window If the material, the tools and the conversation fit with room to work, a single loop is the right shape. It is cheaper, it is faster, and in our own measurements it scores at least as well - sometimes better. Reach for nothing more.
Use structure The material is bigger than the window Once the corpus cannot fit, the loop is not choosing what to keep - the window is, by age. Stages let you decide what survives each boundary, which is the difference between an agent that runs out of room and one that was given a plan for it.
Use structure A wrong answer is worse than none When the deliverable has to be exactly right, put a stage between the draft and the reader whose only job is to check it against the source. A pipeline can require that. A prompt can only ask for it.
Both shapes are the same binary and the same blueprint file. A single loop is one stage with a big window; the pipeline is six with budgets between them. Choosing wrong is an edit, not a migration - which is why we can afford to tell you when not to use the complicated one.
Ten thousand agents, and the runtime out of the way
These came off one machine, an Apple M3 Max with 16 cores, running a mixed fleet of the bundled agents against a mocked model that answers in a flat 1.5 seconds, from one 28 MB binary.
Memory is live memory rather than RSS, which is the smaller of the two figures: what the process still holds once the pages the kernel can reclaim come out of it. CPU is a share of the whole machine, so 100% would be all sixteen cores.
The mocked model is the point: a real one would bury what the runtime costs under its own latency. So this is the floor the runtime adds rather than a forecast for your workload, and with a real model the first wall you meet is your rate limit rather than your machine.
Peak memory, by how many agents were running
Inference pool held at 512 for every tier, so the only thing changing is the number of agents. The runtime is paid for once, so the cost per agent falls as you add them: 3.4 MB each at ten, 0.28 MB each at ten thousand. That top tier is where the ladder stops rather than where the machine did: 2.8 GB of 128 GB, with CPU peaking at 44%.
Runs in flight at the same moment
Same fixed pool of 512. Agents in this fleet start sub-agents, so more runs happen than were spawned; how many overlap at one instant is a separate number, and at ten thousand it comes out lower.
From cold, with nothing warm anywhere
Starting the daemon is measured inside each scenario rather than subtracted from it, so the 64 ms to begin a run already contains the 22 ms to boot one.
CPU while 1,000 agents were running
Averaged across the span where runs were actually active, so idle time at either end cannot flatter it: under 6% of sixteen cores at every pool width.
Time for 1,000 agents to finish, by inference pool width
Agents spawned, inference pool held at 512
Memory over the whole run, and after it
3 GB 2 GB 1 GB 0 0 300 600 900 Seconds since the tier started
Every tier climbs while there is work and drops back to nothing when the last run drains. A leak would be the line that never comes back.
100% 50% 0 0 300 600 900 Seconds since the tier started
Ten thousand agents hold the machine at roughly 40% of sixteen cores for eleven minutes, and the smaller tiers barely mark the axis at all. Most of what is left is time spent waiting on a model.
Inference pool width, 1,000 agents spawned
A wider pool puts more of the work in flight
1.4 GB 700 MB 0 0 100 200 300 Seconds since the tier started
The same thousand agents, four pool widths. A wider pool runs more of them at once, so the memory is higher while it lasts and it lasts a lot less time: 128 stays under 400 MB across five minutes, 1,024 peaks past 1.2 GB and is finished in ninety seconds.
20% 10% 0 0 100 200 300 Seconds since the tier started
Note the axis: this one tops out at 20% of the machine, not 100%. Read the area rather than the height: the total CPU spent falls at every step, from 14.5 machine-seconds at 128 to 2.9 at 1,024 for the same thousand agents, mostly because the wider pool is finished sooner.
All three repetitions are drawn in every chart, not an average of them, which is why each colour is three lines almost on top of each other. The bars above report the median of the three, so the tallest line here runs slightly above its bar.
The closest thing to a comparison
We do not benchmark other frameworks. Whatever baseline we wrote, somebody could fairly say we wrote it to lose, so here is a benchmark someone else ran. It measures a different thing to ours, and both sets of

[truncated]
