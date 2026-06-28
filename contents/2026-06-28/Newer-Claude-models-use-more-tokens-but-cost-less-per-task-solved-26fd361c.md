---
source: "https://signoz.io/blog/claude-code-model-comparison/"
hn_url: "https://news.ycombinator.com/item?id=48707468"
title: "Newer Claude models use more tokens but cost less per task solved"
article_title: "Are Claude’s Models Actually Getting Better? I Instrumented Claude Code to Find Out | SigNoz"
author: "gkarthi2800"
captured_at: "2026-06-28T14:33:42Z"
capture_tool: "hn-digest"
hn_id: 48707468
score: 1
comments: 0
posted_at: "2026-06-28T14:15:50Z"
tags:
  - hacker-news
  - translated
---

# Newer Claude models use more tokens but cost less per task solved

- HN: [48707468](https://news.ycombinator.com/item?id=48707468)
- Source: [signoz.io](https://signoz.io/blog/claude-code-model-comparison/)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T14:15:50Z

## Translation

タイトル: 新しいクロード モデルはより多くのトークンを使用しますが、解決されるタスクあたりのコストは低くなります
記事のタイトル: クロードのモデルは実際に改良されていますか?クロードのコードをインストルメント化して調べてみました |シグノズ
説明: ベンチマーク スコアは、モデルがタスクを解決したかどうかを示すものであり、そこに到達するまでのコストを示します。 OpenTelemetry と SigNoz を使用して Claude Code を計測し、精度、解決済みタスクあたりのコスト、トークン、キャッシュ使用率、エラー率、アクティブ時間に関して Claude Sonnet 4.6、Opus 4.7、および Opus 4.8 を Fi 上で比較しました。
[切り捨てられた]

記事本文:
クロードのモデルは実際に改善されているのでしょうか?クロードのコードをインストルメント化して調べてみました | SigNoz SigNoz 製品ドキュメント リソース 価格 ドキュメントの検索... K サインイン 開始する - 無料 メイン メニューを開く
クロードのモデルは実際に改善されているのでしょうか?クロードコードをインストルメント化して調べてみました
クロードのモデルは実際に改善されているのでしょうか?クロードコードをインストルメント化して調べてみました
数か月ごとに新しいモデル、新しいベンチマーク数値が発表され、最新モデルがこれまでで最高であるという投稿が掲載されます。私は彼らのことをほとんど信じています。スコアが上がります。しかし、リリースが決まったときに私が気にするのは「スコアが上がる」ということではありません。問題は、新しいモデルが私の仕事をより良くするかどうか、そしてそれを調べるのにどれくらいのコストがかかるかということです。
これらは異なる質問であり、それらの間の差はリーダーボードが示しているよりも大きいです。
ベンチマーク スコアは、モデルがタスクを解決したかどうかを示します。モデルがそこに到達するまでに費やした金額はわかりません。
3 倍の費用または 2 倍の時間がかかりながら、わずかに多くのタスクを解決するモデルが、明らかにより良い選択であるとは限りません。それは最適化の対象によって異なります。
なので、クロードの最新モデルが出たときは、「良くなっているのか」という実践編を追いかけました。同じエージェント、同じタスク、3 つのモデル、すべての実行で完全なインストルメンテーションを使用したため、それぞれが完了したかどうかだけでなく、どのように完了したかを確認できました。データの収集と読み取りには OpenTelemetry と SigNoz を使用しました。
簡単に言うと、「はい、良くなってきています」です。より長い答えは、「より良い」は、測定できるようになると多くのものに分かれ、それらはすべて同じ方向を向いているわけではないということです。
それは数字だけではわかりません。信頼できるタスク セットと、実行ごとのメトリクスの 2 つが必要です。
最初の本能は、自分自身のタスクを書くことです。それを試してみて捨てました。

手作りのタスクには信頼性の問題があります。それはあなたがデザインしたものであるため、どのような結果であっても、「自分の主張を伝えるためにテストを作成したのではないか」と批判される可能性があります。それはまた、何が正しいとみなされるかを決定することを意味し、それはまさにベンチマークを議論の余地のあるものにする主観です。
そこで、ターミナルベンチを使用しました: https://github.com/harbor-framework/terminal-bench 。これは、実際の端末で実際の作業を行うエージェントのための公開ベンチマークであり、重要な特性がいくつかあります。すべてのタスクには以下が付属しています。
独自のコンテナ化された環境
合否を決めるテストスクリプト
解決可能であることを証明する参照ソリューション
評決は脚本のものであり、私のものではありません。また、クロード コードをエージェントとして直接実行することもサポートされているため、エージェントの意図された動作方法を測定しています。
テスト スイート全体を実行すると、途方もない量のトークンと時間がかかるため、実行しませんでした。代わりに、私は 10 個のタスクを選択しました。これらは、「コーディング エージェント」が実際に意味すると私が考えるコードの理解と生成を実際に実行できるように選択しました。エージェントが正しいものを書く前に、既存のコードを読んで理解する必要があるタスク。
10 個はその方向に傾いており、私はそれらを意図的に難易度全体に分散させました。すべてのモデルが通過すると予想できるほど簡単なもの (フロア)、誰も通過しないと予想できるほど難しいもの (ヘッドルーム)、そしてモデルを分離するために真ん中に数人がいます。
countdown-game : ターゲットに対して正確に評価される、固定された数値セットから算術式を構築します。
cpp-compatibility : C++ テンプレート関数をダウングレードして、同じインターフェイスを維持しながら C++11 でコンパイルできるようにします。
count-call-stack : プロファイラー スタック トレースを解析し、頻度別に上位 10 件の一意の呼び出しサイトをレポートします。
debug-long-program : 2 つの HTTP エンドポイントだけを使用して、目に見えないプログラム内のすべてのバグを見つけて修正します

■ ( /checkpoint および /fix )。
ode-solver-rk4 : 厳密なステップ サイズと精度制限の下で数値 IVP ソルバー (RK4 または類似) を作成します。SciPy は許可されません。
Schemelike-metacircular-eval : Scheme のようなインタプリタを読み取り、テスト プログラムとそれ自体を解釈できるその言語でエバリュエータを作成します。
リバースエンジニアリング : バイナリのパスワードを回復して実行します。
write-compressor : デコンプレッサーのみを指定すると、正確なターゲット テキストに解凍される 2500 バイト未満のファイルが生成されます。
cancel-async-tasks : 実行がキャンセルされた場合でも各タスクのクリーンアップを実行する、同時実行数の上限を設定した非同期ランナーを作成します。
Polyglot-rust-c : Rust と C++ の両方として正しくコンパイルおよび実行され、k 番目のフィボナッチ数を出力する 1 つのソース ファイルを作成します。
どれも推測に報いるものではありません。読むか失敗するかです。そして、そのうち 3 人はすべてより難しい層に属し、最終的にモデルを分離する作業のほとんどを行うことになりました。
何かを実行する前、そして単一の結果を見る前に、これら 10 個を選択しました。スコアを確認してからタスクを選択すると、たとえ最善の意図があったとしても、ストーリーに合わせてセットを調整することができます。最初に選んでそう言うことは、それをテーブルから外します。
これら 10 個のタスクを、Claude Sonnet 4.6、Claude Opus 4.7、Claude Opus 4.8 の 3 つのモデルを使用して Claude Code を通じて実行しました。実行間で変更された唯一の点は --model フラグでした。 (最近のアクセス変更を受けて、Fable は省略しました。)
Terminal-Bench README に従ってセットアップを行いました。次に、各モデルを同じコマンドで駆動し、 --model 値のみを交換しました。
tb run --agent claude-code --model anthropic/claude-opus-4-8 \
--タスクID カウントダウンゲーム \
--タスクID cpp-互換性 \
--タスクIDカウントコールスタック\
--タスクIDデバッグロングプログラム\
--タスクID ode-solver-rk4 \
--タスク ID スキーム様-メタサーキュラー-eval \
--タスクID r

エバースエンジニアリング \
--タスクID ポリグロット-Rust-c \
--タスク ID 書き込みコンプレッサー \
--タスクIDキャンセル非同期タスク\
--dataset-path 元のタスク
API をループで実行するのではなく、実際のエージェントを駆動する理由は、エージェントが人々が使用するものだからです。ツールの呼び出し、再試行、コンテキスト管理、ファイルをいつ読み取るか編集するだけかについての小さな決定。これらすべてが、コーディング作業のモデルを比較する際の比較対象の一部です。エージェントを取り除くと、関心のあるものよりも狭いものを測定することになります。
Claude Code は OpenTelemetry によるモニタリングをネイティブにサポートしているため、追跡可能です。実行ごとに、トークン、コスト、ツール呼び出し、モデルリクエストなどのメトリクスとトレースが生成されます。すべてを SigNoz に送信し、モデルとセッションのテンプレート変数を含む 1 つのダッシュボードを構築したので、表示しているモデルに関係なく同じパネルが再レンダリングされます。
私がダッシュボードに載せたものと、それぞれがスポットを獲得した理由:
タイプ別のトークン。合計だけではありません。 Claude Code は、新規入力、出力、キャッシュ読み取り、キャッシュ作成の 4 種類を個別にレポートします。この分割は私が予想していた以上に重要であり、それが、後の調査結果の 1 つが目に見える理由です。
コスト (ドル単位)。支出の唯一の正直な尺度。トークンはその代理ですが、すべてを伝えるわけではありません。
キャッシュの使用率。プロンプト キャッシュから取得された入力と新たに処理された入力の割合。トークンの内訳から派生するため、モデル固有でもあります。
LLM リクエスト。モデルへの往復。タスクの往復にかかった時間を大まかに測定します。
コード行。エージェントが作成して保持したコードの量。単調ですが、出力ボリュームのプロキシとして使用できます。
ツール呼び出し。エージェントが考えたことに対してどれだけのことをしたか。読み取り、編集、bash コマンドなど
アクティブな時間。エージェントが実際に作業に費やした時間 (アイドル状態は除く)。
エラー率と

誤差範囲。ツールの実行が失敗した頻度と、実際の失敗の表。
Claude Code のテレメトリは、一部の信号にモデル名を付加しますが、その他の信号には付加しません。コスト、トークン使用量、キャッシュ使用率、LLM リクエスト、およびコード行はすべて属性としてモデルを保持するため、これらのパネルは完全に 1 つのモデルにフィルターされます。残りの部分 (ツール呼び出し、アクティブ時間、エラー率、エラースパン) にはまったく影響しません。これらは実行ごとに発行されます。
コストの数値も 1 つ修正が必要でした。 Claude Code は、安価な内部作業のために、より小規模なモデル Haiku 4.5 を使用しており、その支出は生の合計に表示されます。これはテレメトリ内で分離可能であり (query_source 属性で補助呼び出しをマークします)、コストが平均で約 1% になったため、除外することにしました。
すべての指標が同じ結論を示しているわけではないため、この指標を一度に 1 つずつ見ていきます。 1 つの数字を単独で見ると、自信はあるものの、部分的に間違った結論が得られるでしょう。興味深いのは、これらを並べると何が起こるかです。 (すべての数値を一度に確認したい場合は、最後にまとめた表を参照してください。)
ここでは、モデルごとに 1 つずつ、計 3 つのダッシュボードを示します。パネルは同一です。モデルフィルターのみがそれらの間で変更されました。モデルを切り替えると、同じパネルが再レンダリングされます。
各モデルが 10 個のタスクのうち何個を解決したか:
ここでやめてしまうと、記事は 3 文になり、結論は「Opus 4.8 を使用する」になってしまいます。それは間違いではありません。しかし、精度は決して疑問の余地のない唯一の数値であり、追加の解決策に何を支払うか、または待つかについては何も示しません。
注意すべき点は、失敗の種類がすべて同じではないということです。これが重要です。ソネットのミスのうち 2 つはタイムアウトで、時計がまだ動いていませんでした。 Opus 4.8 の 1 つのミスは異なりました。すぐに終わり、シムでした。

プライが間違っています。
ここで、きれいなストーリーが中断される最初の場所です。 10 個のタスクにわたる合計支出:
順序を見てください。中間モデルである Opus 4.7 は、3 つのモデルの中で最も高価で、より新しく、より多くの解決策を備えた 4.8 よりも高価でした。 「機能が高ければコストもかかる」という単純な直感は間違いです。最新モデルは以前のモデルよりも性能が良く、安価でした。
そして、原価ベースで最も安い Sonnet が安いのは、タスクの半分が失敗し、そのうち 2 つはタイムアウトによって失敗したためです。そこで、原材料コストでは答えられないという疑問が生じます。何あたりで一番安いですか？
私が信頼する数字は、解決されたタスクあたりのコスト、つまり総支出を完了したタスクで割ったものです。
最新モデルは、実行コストが最も安いわけではありませんが、実際にタスクを完了するための最も安価な方法です。
「実行するのに最も安いもの」と「作業を完了するための最も安いもの」との間のギャップが、結果当たりのコストがトークン当たりのコストを上回る理由のすべてです。 Sonnet は、何が終わったかを考慮するまでは予算の選択肢のように見えますが、終わった時点で利点はほとんど消えてしまいます。
トークンの合計だけでモデルをランク付けすると、合理的に聞こえるストーリーが得られますが、誤解を招く可能性があります。
Sonnet 4.6: ~ 288 万トークン
Opus 4.7: ~ 588 万トークン
Opus 4.8: ~ 606 万トークン
したがって、ソネットはどちらの Opus モデルも半分未満しか使用しませんでした。トークンが直接コストを意味する場合、価格は半額以下になります。 4.8 の $8.06 に対して $6.30 ではありませんでした。トークンの半分、コストの約 4 分の 3。これらの事実は、トークンを分解するまで適合しません。これがまさに、タイプごとの内訳をダッシュ​​ボードに記載した理由です。
まずはキャッシュ。使用率は 3 つすべてでほぼ 100% に固定されており、ほぼすべての入力トークンが新たに読み取られたのではなく、プロンプト キャッシュから取得されたことを意味します。 Opus で実行されるこれら 600 万のトークンは、ほとんどがキャッシュ読み取りであり、このモデルが行う処理としては最も安価です。

安価なカテゴリが膨大であるため、総数は膨大です。それを「使用量」として数えても、請求額についてはほとんどわかりません。
第二に、コストは高価なトークン、つまり出力によって決まります。 Sonnet の高価なトークン、生成される出力と作成されるキャッシュは、Opus の実行と同じ規模であるため、そのコストもそこに影響します。トークンの総量は少なく、同等の高価なトークン、同様のコスト。
単純なトークンカウンターからは、Sonnet の効率が 2 倍以上であることがわかります。コスト委員会はそうではないと述べています。どちらが正しいかを知る唯一の方法は、トークンをタイプごとに分割することです。これは、要約番号を信頼するのではなくコーディング エージェントをインストルメント化することが非常に重要である明確な理由です。
単純に考えると、「4.8 は最も遅く、最も長く待つことになる」ということになります。それは間違いです。
このパネルは、アクティブ時間を追跡します。これは、エージェントが作業に費やした時間、アイドル状態を除外した時間、同時実行中のタスクの合計時間です。それは壁時計ではなく、人間がどれだけ待つかでもありません。つまり、4.8 の 49.75 分は「応答が遅くなった」という意味ではなく、「より多くの作業が行われた」ことを意味します。
それを他の努力シグナルと並べると、彼らは同意します。困難なタスクでは、4.8 があらゆる努力の尺度でリードしました。
最も多くのモデル リクエスト (Sonnet の 68 に対して 119)
最新モデルは難しいタスクにさらに力を入れました。
余分な労力が最高の精度と最高の解決コストに変換されます。それはさらに効果があり、その努力は報われました。
ウィ

[切り捨てられた]

## Original Extract

Benchmark scores tell you whether a model solved a task, not what it cost to get there. I instrumented Claude Code with OpenTelemetry and SigNoz to compare Claude Sonnet 4.6, Opus 4.7, and Opus 4.8 across accuracy, cost per solved task, tokens, cache utilization, error rate, and active time, on a fi
[truncated]

Are Claude’s Models Actually Getting Better? I Instrumented Claude Code to Find Out | SigNoz SigNoz Product Docs Resources Pricing Search docs... K Sign In Get Started - Free Open main menu
Are Claude’s Models Actually Getting Better? I Instrumented Claude Code to Find Out
Are Claude’s Models Actually Getting Better? I Instrumented Claude Code to Find Out
Every few months there’s a new model, new benchmark numbers, and a post telling you the latest one is the best yet. I believe them, mostly. The scores go up. But “scores go up” isn’t the question I care about when a release lands. The question is whether the new model does my work better, and what it costs me to find out.
Those are different questions, and the gap between them is bigger than the leaderboards let on.
A benchmark score tells you whether a model solved a task. It doesn’t tell you what the model spent getting there:
A model that solves slightly more tasks while spending three times as much, or taking twice as long, isn’t obviously the better choice. It depends on what you’re optimizing for.
So when the latest Claude models came out, I went after the practical version of “are they getting better.” Same agent, same tasks, three models, full instrumentation on every run, so I could see not just whether each one finished but how . I used OpenTelemetry and SigNoz to collect and read the data.
The short answer is yes, they’re getting better. The longer answer is that “better” splits into many things once you can measure them, and those things don’t all point the same direction.
You can’t see that with just a number. You need two things: a task set worth trusting, and metrics on every run.
The first instinct is to write your own task. I tried that and threw it out. A homemade task has a credibility problem: you designed it, so any result can be criticized by saying: “well, you built the test to make your point.” It also means you decide what counts as correct, which is exactly the subjectivity that makes a benchmark arguable.
So I used Terminal-Bench: https://github.com/harbor-framework/terminal-bench . It’s a public benchmark for agents doing real work in a real terminal, with a few properties that matter. Every task ships with:
its own containerized environment
a test script that decides pass or fail
a reference solution that proves it’s solvable
The verdict is a script’s, not mine. And it supports running Claude Code as the agent directly, so I’m measuring the agent the way it’s meant to be driven.
I didn’t run the whole test suite, since running that would have taken an absurd amount of tokens and time. Instead, I picked 10 tasks, chosen so they actually exercise code comprehension and generation, the two things I think “coding agent” really means. Tasks where the agent has to read existing code and understand it before it can write anything correct.
The 10 lean that way, and I spread them across difficulty on purpose: some easy enough that I expected every model to pass (a floor), some hard enough that I expected none to (headroom), and a handful in the middle to separate the models.
countdown-game : build an arithmetic expression from a fixed set of numbers that evaluates exactly to a target.
cpp-compatibility : downgrade a C++ template function so it compiles under C++11 while keeping the same interface.
count-call-stack : parse profiler stack traces and report the top 10 unique call sites by frequency.
debug-long-program : find and fix every bug in a program you can’t see, using only two HTTP endpoints ( /checkpoint and /fix ).
ode-solver-rk4 : write a numeric IVP solver (RK4 or similar) under strict step-size and accuracy limits, no SciPy allowed.
schemelike-metacircular-eval : read a Scheme-like interpreter and write an evaluator in that language that can interpret the test programs and itself.
reverse-engineering : recover a binary’s password and run it.
write-compressor : given only a decompressor, produce a file under 2500 bytes that decompresses to an exact target text.
cancel-async-tasks : write an async runner with a concurrency cap that still runs each task’s cleanup when a run is cancelled.
polyglot-rust-c : write one source file that compiles and runs correctly as both Rust and C++, printing the kth Fibonacci number.
None of them reward guessing. You read, or you fail. And three of them, all in the harder tiers, ended up doing most of the work of separating the models.
I picked these 10 before running anything and before seeing a single result. If you choose tasks after seeing scores, even with the best intentions, you can nudge the set toward whatever flatters your story. Picking first and saying so takes that off the table.
I ran those 10 tasks through Claude Code with three models: Claude Sonnet 4.6, Claude Opus 4.7, Claude Opus 4.8. The only thing that changed between runs was the --model flag. (I left Fable out, following the recent access changes around it.)
I followed the Terminal-Bench README to get set up. Each model was then driven with the same command, swapping only the --model value:
tb run --agent claude-code --model anthropic/claude-opus-4-8 \
--task-id countdown-game \
--task-id cpp-compatibility \
--task-id count-call-stack \
--task-id debug-long-program \
--task-id ode-solver-rk4 \
--task-id schemelike-metacircular-eval \
--task-id reverse-engineering \
--task-id polyglot-rust-c \
--task-id write-compressor \
--task-id cancel-async-tasks \
--dataset-path original-tasks
The reason to drive the actual agent, rather than hit the API in a loop, is that the agent is what people use. Its tool-calling, retries, context management, its little decisions about when to read a file versus just edit it: all of that is part of what you’re comparing when you compare models for coding work. Strip the agent out and you’re measuring something narrower than what you care about.
Claude Code supports monitoring with OpenTelemetry natively, which made this trackable. Every run produces metrics and traces for tokens, cost, tool calls, model requests, and more. I sent all of it to SigNoz and built one dashboard with template variables for model and session, so the same panels re-render for whichever model I’m viewing.
What I put on the dashboard, and why each earned a spot:
Tokens, by type. Not just total. Claude Code reports four kinds separately: fresh input, output, cache reads, cache creation. This split matters more than I expected, and it’s why one of the later findings is visible at all.
Cost, in dollars. The only honest measure of spend. Tokens are a proxy for it, but don’t tell the whole story.
Cache utilization. What fraction of input came from the prompt cache versus being processed fresh. Derived from the token breakdown, so model-specific too.
LLM requests. Round-trips to the model. A rough measure of how much back-and-forth a task took.
Lines of code. How much code the agent wrote and kept. Blunt, but a usable proxy for output volume.
Tool calls. How much the agent did versus thought . Reads, edits, bash commands, etc
Active time. Time the agent spent actually working, idle excluded.
Error rate and error spans. How often tool executions failed, plus a table of the actual failures.
Claude Code’s telemetry attaches the model name to some signals but not others. Cost, token usage, cache utilization, LLM requests, and lines of code all carry the model as an attribute, so those panels are genuinely filtered to one model. The rest (tool calls, active time, error rate, error spans) don’t carry it at all. They’re emitted per run.
The cost numbers also needed one correction. Claude Code uses a smaller model, Haiku 4.5, for cheap internal work, and that spend shows up in raw totals. It’s separable in the telemetry (a query_source attribute marks auxiliary calls) and came to about 1% of cost on average so I decided to exclude them.
I’ll walk through this one metric at a time, because not all the metrics point to the same conclusion. Look at one number by itself and you’ll come away with a confident but partly wrong conclusion. The interesting part is what happens when you put them side by side. (If you’d rather see every number at once, the consolidated table is at the end.)
Here are the three dashboards, one per model. The panels are identical; only the model filter changed between them. Switch models to see the same panels re-render.
How many of the 10 tasks each model solved:
If I stopped here, the post would be three sentences and the conclusion would be “use Opus 4.8.” That isn’t wrong, exactly. But accuracy is the one number that was never in doubt, and it says nothing about what you’d pay or wait for those extra solves.
One thing to note: the failures weren’t all the same kind, which matters. Two of Sonnet’s misses were timeouts, where it ran out the clock still working. Opus 4.8’s single miss was different: it finished quickly and was simply wrong.
Here’s the first place the clean story breaks. Total spend across the 10 tasks:
Look at the ordering. Opus 4.7, the middle model, was the most expensive of the three, more than 4.8, which is newer and solved more. The tidy “more capable costs more” intuition is wrong. The newest model was better and cheaper than the one before it.
And Sonnet, cheapest in raw dollars, is cheap partly because it did less: it failed half the tasks, two by timing out. Which raises the question raw cost can’t answer. Cheapest per what?
The number I trust is cost per solved task: total spend divided by tasks completed.
The newest model is the cheapest way to get a task actually done , even though it’s not the cheapest to run.
That gap between “cheapest to run” and “cheapest to get work done” is the whole reason cost-per-outcome beats cost-per-token. Sonnet looks like the budget option until you account for what it didn’t finish, at which point the advantage mostly evaporates.
If you rank the models by total tokens alone, you get a story that sounds reasonable but can be misleading.
Sonnet 4.6: ~ 2.88 million tokens
Opus 4.7: ~ 5.88 million tokens
Opus 4.8: ~ 6.06 million tokens
So Sonnet used less than half of either Opus model. If tokens directly meant cost, it’d be less than half the price. It wasn’t: $6.30 against 4.8’s $8.06 . Half the tokens, roughly three-quarters of the cost. Those facts don’t fit until you break the tokens apart, which is exactly why I put the per-type breakdown on the dashboard.
First, cache. Utilization was pinned near 100% across all three, meaning almost every input token came from the prompt cache rather than being read fresh. Those 6 million tokens on the Opus runs are mostly cache reads, the cheapest thing the model does. The total count is huge because the cheap category is huge. Counting it as “usage” tells you almost nothing about the bill.
Second, cost is driven by the expensive tokens: output. Sonnet’s expensive tokens, the output generated and cache created, are in the same ballpark as the Opus runs, so its cost lands there too. Fewer total tokens, comparable expensive tokens, similar cost.
The simple token counter would have told you Sonnet was more than twice as efficient. The cost panel says it wasn’t. The only way to know which is true is to split tokens by type. This is a clear reason why instrumenting your coding agent instead of trusting a summary number is very important.
The naive read is “4.8 is slowest, you’ll wait longest.” That’s wrong.
This panel tracks active time: time the agent spent working, idle excluded, summed across concurrently running tasks. It’s not wall-clock and not how long a human waits. So 4.8’s 49.75 minutes doesn’t mean “slower to respond,” it means “did more work.”
Line it up with the other effort signals and they agree. On the hard tasks, 4.8 led on every measure of effort:
the most model requests ( 119 against Sonnet’s 68 )
The newest model worked harder on the hard tasks.
The extra effort converted into the highest accuracy and the best cost-per-solved. It worked more, and the work paid off.
Whi

[truncated]
