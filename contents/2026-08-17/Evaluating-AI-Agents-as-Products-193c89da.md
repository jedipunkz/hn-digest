---
source: "https://www.marble.onl/posts/evaluating_ai_product_quality.html"
hn_url: "https://news.ycombinator.com/item?id=49337534"
title: "Evaluating AI Agents as Products"
article_title: "Evaluating AI Agents as Products"
image: ""
author: "amarble"
captured_at: "2026-08-17T21:17:15Z"
capture_tool: "hn-digest"
hn_id: 49337534
score: 1
comments: 0
posted_at: "2026-08-17T20:57:59Z"
tags:
  - hacker-news
  - translated
---

# Evaluating AI Agents as Products

- HN: [49337534](https://news.ycombinator.com/item?id=49337534)
- Source: [www.marble.onl](https://www.marble.onl/posts/evaluating_ai_product_quality.html)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T20:57:59Z

## Translation

タイトル: 製品としての AI エージェントの評価

記事本文:
AIエージェントを製品として評価する
コーディング ベンチマークは、明確に指定されたタスクに関して AI エージェントを評価します。
オープンエンドの定性チェックでは、どれが最も優れたデモを作成できるかを示しますが、
実際の剛性を測定しないほど十分な柔軟性を残す
仕事。人間参加者のパフォーマンスをキャプチャすることも、会話することもありません。
効率と生産性。より適切に測定できる方法を探るため
「使える」AIエージェント、データを生成する3つのタスクを紹介
(理想的には) 短いインタラクティブなコーディング セッションでインターフェイスにラベルを付ける。
指定されたインターフェイス機能は必須要件を組み合わせています
自由に選択できるため、効率性を評価できます。
コラボレーションとセンスにより、エージェントがどの程度優れているかを示す指標となります。
実際のコーディング状況でツールとして機能します。ここでタスクを紹介します
5 つのエージェント (3 つの小規模なローカル モデルと 2 つの近傍モデル) を評価します。
フロンティア。
アンドリュー・マーブル
大理石.onl
アンドリュー@willows.ai
2026 年 8 月 15 日
私たちが ML を始めたとき、ベンチマーク スコアはかなり正確に捕捉されました。
モデルのパフォーマンス。分類器が特定の精度を持っていて、
思い出してください、データ分布が変わらない限り、これでわかりました
知る必要があったこと。 AI の起源はより単純な ML モデルにあると思います。
学問の進歩と商業との間には現在の狭いギャップがある
オファリングは常にベンチマークを重視してきました。
ベンチマークは、(a) 比較したい場合に最も役立ちます。
(b) 現在、モデルが急速に進歩している時期にあります。
そして、LLM の短い歴史の中で、モデルのパフォーマンスはおそらく
製品のパフォーマンスに適した代替手段。 GPT-5.6 で構築するものはすべて、
GPT-4 で構築した場合よりも優れたものになるでしょう。既知の問題
もちろん、そのベンチマークが飽和し、すべてのモデルが基本的に「エース」になれるのであれば
古いベンチマークのため、比較には役に立ちません。ル

議論されているSSは
モデルベンチマークの概念全体が多くの分野で飽和している可能性がある
残りのアルファはより製品的な品質になります。
ユーザーエクスペリエンスととらえどころのない「味」をテーマにしています。
人工知能コーディング エージェント インデックス 1 は、
3 つの公開ベンチマーク データセット (DeepSWE、
Terminal-Bench V2、SWE-Atlas-QnA)。このうち最初の 2 つはエージェントが行うことです
予期される動作が明確に定義されたソフトウェア エンジニアリング タスク。最後
エージェントにコードベースに関する質問に答えてもらいます。
私は、そのような評価はもはや、どのように行われるかということと一致しないと主張したい。
特定のツール (Cursor、Codex、Claude Code など) が有効になります。
開発者の成果を向上させる (時間の節約、より多くの出荷など)。まあ
指定されたタスクは、複数のタスクとは異なる一連の動作をテストします。
センスと判断力が必要な、オープンエンドの現実世界の問題 2 ．さらに、私たちは次のような段階にいます。
ほとんどのエージェントは、人間が関与している場合、基本的にあらゆるタスクを解決できます。
解決できないタスクを見つけることは、ますます彼らを騙すことになる
または、テストの代わりにエッジケースを特定するほうが有益です。
現実世界。関連する質問は、これまでのように「これはできるか」というものではなくなりました
初期の段階ではありましたが、それをどれほど効率的に行うことができるでしょうか（人間の場合）
ループ内）、品質につながる適切な判断が行われていますか
製品について、同僚として有益な意見が得られるかなど。
一方で、オープンエンドテストには危険が伴います。
明確な目標。最近、アンドレイ・カルパシー氏はビデオの視聴を促すことを提案しました。
第 3 世代、同様のコーディングがあります
ワンショットビデオゲームのような集中的なタスク。これらはクールなデモを作成しますが、
実用性を現実的にテストするには限界がありすぎます。 AI の永続的な問題
それは、デモには最適ですが、現実世界の考慮事項には逆風が吹くということです
アウト

その帆の。 Netlify のブログ投稿でも、モデルがどの程度優れているかを比較しました。
Web デザイン タスクを単発で実行できる 4 。これにも同様のものがあります
問題は、主観的な使用基準に対してテストを行っていないことです。
インタラクションではなくシングルショットのパフォーマンスのみを測定するため、
ツールが人間と一緒に使用されたときにどれほど役立つかを明らかに予測する
ループ。
さまざまなコーディング エージェント (モデルで構成される) がどれほど役立つかを調査するには
およびハーネス）は、十分に具体的な一連のタスクが必要でした。
制約と必要な成功基準があると同時に、
モデルが自分の好みを発揮できる十分な柔軟性。重要なことに
また、実際の使いやすさの基準も必要でした。ただ構築したくなかった
テストに合格し、デザインも優れていると何か言って、私は欲しかった
最終結果を決定できるユーザーと同等の存在が必要です。
彼らのニーズを満たしました。目標は、実際の使用を完全にシミュレートすることではなく、
ツールの効率性と効果性を実感できる代理タスク
有意義な対話と成功をシミュレートすることで、実際の状況を再現します
基準。
これらの基準に基づいて、エージェントに次のことを要求するタスクを試してみました。
データセットのラベル付けまたは注釈をサポートするソフトウェアを構築します。これは
実生活で AI を頻繁に使用するタスク。明確な機能が含まれます
基準 - ユーザーには一連のデータポイントが提示され、次のことを行う必要があります。
それぞれについて何らかの決定を下す (ラベルを追加する) ため、ソフトウェアは
それをサポートしなければなりません。それは味と判断力を必要とします。
マテリアルの提示、ラベル付けフロー、進行状況の保存方法など。
そして、その最終的な成功は、ユーザーがどれだけ効果的に実行できるかにかかっています。
ソフトウェアを使用してラベル付けを実行します。
一般的な基準: 人間の判断を必要とするいくつかの単純なタスク
迅速に完了することができました。感触を得るのに十分なボリューム

あるいはどうやって
ユーザーは多くのサンプルを効率的に実行できます。意味のあるアップサイド
ユーザーの情報をどのようにレイアウトすればよいか
タスク。
クラスのラベル付け: Wikipedia から 75 文ずつ抽出しました
(a) ドットコム バブルと (b) 2008 年の金融危機に関する記事。
目標は、ユーザーがどの記事にラベルを付けるかを支援するインターフェイスを構築することです。
その文は（彼らがそれを読めば分かると仮定して）から来たものです。
ランキング: 50 の Unicode 文字 (♠、♥、☀、☃ など) について、
LLM (Qwen 3.6 35B A3B) を使用して、.svg 線描画で 5 回の試行を生成します。
目標は、ユーザーが試行をランク付けできるインターフェースを構築することです。
キャラクターごとに。
コードの編集と注釈: 100 個の Python コード スニペットを生成しました
さまざまな構文エラーが導入された関数で構成されています。目標
ユーザーがコードを修正できるインターフェースを構築することです。
エラー メッセージに応じて、エラーの場所に注釈を付ける / 検証する
Python インタプリタから。
テストセットアップ: これまでにテストしたモデルは Z.ai GLM 5.2、Qwen 3.6 35B です。
A3B (8 ビット クオント)、Laguna Poolside 2.1 (4 ビット クオント)、Meta Muse Glimmer
(8 ビット quant)、および Claude Opus 5。Claude は Claude Code とともに使用され、
残りはストック Pi (別のコーディング ハーネス) で使用されました。私が使用した
それぞれのデフォルトの考え方が何であれ、これは意図されたものではありません
思考モード間の詳細な比較。モデルが選ばれました
何らかの理由で使いたかったものに基づいています。クロードだから
全体的に良好で、長い間使用しています。現代的なGLM
オープンウェイトモデルと、実行できる例としての 3 つの小さいモデル
ローカルでアクティブとトータルの間で異なるトレードオフを伴う
パラメータ。
テスト手順: 3 つのタスクのそれぞれについて、GLM を使用して
で始まる対話型セッションでの参照実装
p

前述のロンプト。これらの各セッションを使用して、
私が望んでいた基本機能をキャプチャした適切な終了状態
アプリケーションを参照してください。私は実際の仕事と同じように仕事に取り組みました
開発設定、アプリケーションの反復実行、特定
問題や改善点、そしてそれらへの対処を促します。
後続のエージェントについては、可能な限り再現しました。
同じ開始プロンプトで始まり、同じものを使用するセッション
該当する場合はフォローアップしますが、状況に応じて必要に応じて調整します
生じた問題。
評価: 使用されたトークンの数と数を確認できます。
効率とネイティブの好みの尺度として必要なユーザーターンの数。少ない
ターンは、モデルが良い結果に達したことを示します
最初のプロンプトの仕様が不十分であるにもかかわらず、より高速になります。多ければ多いほど
評価の興味深い部分は、実際に体験することから生まれます。
インタラクションと出力はほとんどが手動で定性的です。これ
自動化されたベンチマークと比較すると魅力がない、または次のように感じるかもしれません
警官出動。ただし、目標は「製品テスト」を再現することです。
他の領域、たとえばランニング シューズや車などは、結局のところ、
何かを試して、それがどれだけうまく機能するかを確認してください。 1 つの定量的尺度
このメソッドはサポートしていますが試みられていないことは、実際にどのようにタイミングを計るかです
下流のデータラベル付けタスクを効率的かつ正確に行うことができます。
さまざまなツールを使用して完了します。これに開発も相まって、
時間と手動ステアリングが必要ですが、よりしっかりとした定量的な結果を得ることができます
さまざまなエージェントがどれほど効果的かを示す図。
定性的評価では、どのようなスタイルの選択があったかを調べます。
エージェントがどの程度の直感や常識を持ってタスクをサポートするように作られているか
このタスクと明らかに使いやすさについて考えた

開発されたものの
解決策。使用した全体的な経験と印象も考慮します
エージェント。評価を実行する際、彼らは非常に代表的であると感じました。
デバッグを含む、コーディング エージェントとのやり取りの種類
サイドクエストとステアリングの必要性。
何ターンかかったかを定量的に測定することもできます。
多くのトークンが使用されましたが、どちらも効率を高めるためのものでした。
以下は、それぞれで生成された 3 つのアプリケーションのスクリーンショットです。
モデル
クラスのラベル付けタスクでは、GLM にサイドバーを追加するように指示する必要がありました。
考えるのに時間がかかりすぎて、デフォルトのトークン制限に 2 回達してしまいました。
最初は構文の強調表示がめちゃくちゃで、レイアウトが崩れていました
単語を強調表示するたびに。と聞かれたときは特に苦労しました
単語を強調表示する方法を思いつき、作ろうとして困惑した
単語リスト。また、（私の考えでは）赤を選んだのはデザイン上の選択としては不適切でした。
良い/悪い意味合いを考慮すると、2 つのクラスカラーは緑です。
この２色。
ランキングタスクのために、GLMは黒地に黒に近いインターフェイスを作成して描画しました
最初は見えなかった写真。も表示されました
タスクに関連性のない大量のメタデータ。
編集タスクに関して、GLM には初期のスクロールとレイアウトの問題がいくつかありました。
表示された無関係な情報 - 指示には次のように指定されています
これはPythonエラーメッセージのみを使用したテストですが、表示されました
とにかく他の情報。
全体的に GLM は使いやすいアプリケーションを生成し、優れたものにしました
文体の選択。たとえば、前もって文が大きくなり、
クラスのラベル付けタスクにとって読みやすい。ランキングタスクについては、
ただし、インターフェースの忙しさとダークモードのせいで、
読みましたが、後から振り返って、明暗の切り替えが追加されていることに気づきました。
編集インターフェイスでは、全体を表示するにはスクロールする必要があります。

アインパネル
そのため使用が遅くなります。
Qwen はクラス ラベル付けアプリのスペースをあまり活用していませんでした。
審査対象の文章が必要以上に短かった。に似ている
GLM、Qwen は最初はスクロール バーを追加しておらず、スクロール バーを台無しにしてしまいました。
最初のパス。構文を適用する際のレイアウトにも問題がありました
ハイライト。
ランキングタスクでは、クウェンは私が思ったように完全に完了しませんでした
欲しかったが、うまくいくものを完成させた。エラーが発生しました
最初に読み込みから、複数の人間参加型デバッグ ラウンドが必要でした
画像をドラッグ アンド ドロップします。こちらも初期表示されていました
黒地黒の写真。
編集作業に関して、クウェンは当初、そうではなかった情報を共有しました
ユーザーが行うことになっている、エラーの強調表示でのステアリングの必要性、
ボタンの状態にエラーがあり、手動で修正する必要がありました
クラスのラベル付けタスクのために、Laguna は最初にターミナルを構築しました
アプリケーション。ただし、Web アプリを作成するというプロンプトには応答しました。
ただし、サイドバーは自動的に追加されます。全文表示されませんでした
これは、特に次の用途に使用する場合には不適切な判断です。
文章の見直し。ラグナもクラスの 1 つに赤を選びました
色。
ランキング課題において、ラグナは以下のような軽微な判断ミスをしただけだった。
を表示する

[切り捨てられた]

## Original Extract

Evaluating AI Agents as Products
Coding benchmarks evaluate AI agents on well specified tasks.
Open ended qualitative checks show which can make the nicest demo but
leave enough flexibility that they don’t measure the rigidity of real
work. Neither capture human-in-the-loop performance nor do they speak to
efficiency and productivity. To explore how we can better measure how
“useful” AI agents are, I introduce three tasks to generate data
labeling interfaces in an (ideally) short interactive coding session.
The specified interface functionality combines mandatory requirements
with open ended choices, allowing us to evaluate efficiency,
collaboration, and taste, providing a proxy for how well the agents
perform as tools in real coding situations. Here I introduce the tasks
and evaluate five agents – three smaller local models and two near
frontier.
Andrew Marble
marble.onl
andrew@willows.ai
Aug 15, 2026
When we started doing ML, benchmarks scores pretty closely captured
the performance of a model. If a classifier had a certain precision and
recall, as long as the data distribution didn’t change this told you
what you needed to know. I think AI’s origins in simpler ML models, and
the current narrow gap between academic advances and commercial
offerings has kept an outsized emphasis on benchmarks.
Benchmarks can be most useful when (a) we want to compare
models and (b) we are in a period of rapid model advancement.
And in the brief history of LLMs, model performance was probably a
suitable proxy for product performance. Anything you build on GPT-5.6 is
going to be better than if you’d built it in GPT-4. A known issue of
course if that benchmarks saturate and all models can basically “ace”
old benchmarks making them useless for comparison. Less discussed is
that the entire concept of model benchmarking may be saturating in many
applications, and the remaining alpha is in more product-like qualities
around user experience and the elusive “taste”.
The Artificial Intelligence Coding Agent Index 1 is
a composite index of three public benchmark datasets (DeepSWE,
Terminal-Bench V2, SWE-Atlas-QnA). The first two of these have agents do
software engineering tasks with well defined expected behavior. The last
one has the agent answer questions about a codebase.
I would contend that such evaluations no longer align with how
effective a given tool (say Cursor or Codex or Claude Code) is at
improving developer outcomes (saving time, shipping more, etc). Well
specified tasks test for a different set of behaviors than more
open-ended real-world problems where taste and judgment are necessary 2 . Furthermore, we’re at a point where
most agents can basically solve any task when a human is in the loop.
Finding tasks they can’t solve becomes increasingly about tricking them
or identifying edge cases instead of testing which is more useful in the
real world. The relevant question is no longer “can it do this” like it
was in the early days, but how efficiently does it do it (with a human
in the loop), does it exercise good judgement that leads to a quality
product, does it provide helpful input as a co-worker, etc.
On the other hand, there is a danger in open-ended testing without
clear goals. Recently Andrej Karpathy suggested prompting for video
generation 3 , and there are analogous coding
focused tasks like one-shotting video games. These make cool demos but
are too open ended to realistically test utility. AI’s perennial problem
is that it’s great for demos but real world considerations take the wind
out of its sails. A Netlify blog post also compared how well models
could single-shot web design tasks 4 . This has a similar
problem, it doesn’t test against any subjective use criteria, and it
only measures single-shot performance, not interaction, so doesn’t
obviously predict how useful a tool will be when used with a human in
the loop.
To explore how useful different coding agents (consisting of a model
and harness) are, I wanted a set of tasks that was specific enough to
have constraints and necessary success criteria, while also involving
sufficient flexibility to let models demonstrate their taste. Crucially
I also wanted real useability criteria. I didn’t want to just build
something and say it passed the tests and looked well designed, I wanted
there to be the equivalent of a user that could decide if the end result
met their needs. The goal is not to fully simulate real use, but to find
proxy tasks that give a sense of how efficient and effective the tools
are in real situations by simulating meaningful interaction and success
criteria.
Based on these criteria I experimented with tasks requiring agents to
build software in support of dataset labeling or annotation. This is a
task I use AI for frequently in real life. It involves clear functional
criteria – a user is presented with a set of data points and needs to
make some determination(s) about each (add a label), so the software
must support that. It involves taste and judgement in terms of how the
material is presented, the labeling flow, how progress is saved, etc.
And its ultimate success rests in how effectively the user is able to
use the software to perform the labeling.
General criteria: some simple tasks requiring human judgement that
could be completed rapidly. Sufficient volume to get a feel for how
efficiently a user can run through many samples. Meaningful upside in
how the information could be laid out to support the user in their
task.
Class labeling: I extracted 75 sentences each from Wikipedia
articles about (a) the Dot-Com bubble and (b) the 2008 financial crisis.
The goal is to build an interface to help the user label which article
the sentence came from (assuming they can tell by reading it).
Ranking: For 50 unicode characters (♠, ♥, ☀, ☃, etc) I asked an
LLM (Qwen 3.6 35B A3B) to generate five attempts at a .svg line drawing.
The goal is to build an interface that lets the user rank the attempts
for each character.
Code editing and annotation: I generated 100 python code snippets
consisting of functions with various syntax errors introduced. The goal
is to build an interface that lets the user correct the code, and
annotate / validate the location of the error, given the error message
from the python interpreter.
Test Setup: The models tested so far are Z.ai GLM 5.2, Qwen 3.6 35B
A3B (8-bit quant), Laguna Poolside 2.1 (4-bit quant), Meta Muse Glimmer
(8-bit quant), and Claude Opus 5. Claude was used with Claude Code and
the rest were used with stock Pi (another coding harness). I used
whatever the default thinking was for each and this isn’t meant to be a
detailed comparison between thinking modes. The models were selected
based on ones I wanted to use for one reason or another. Claude because
it’s generally good and I’ve used it for a long time, GLM as a modern
open weights model, and the three smaller ones as examples I can run
locally with different tradeoffs between active and total
parameters.
Test Procedure: For each of the three tasks I used GLM to build a
reference implementation in an interactive session beginning with the
prompts mentioned earlier. I used each of these sessions to decide on an
appropriate end-state that captured the base functionality I wanted to
see in the application. I approached the task as I would in a real
development setting, iteratively running the application, identifying
issues or refinements, and them prompting to address them.
For the subsequent agents I replayed, to the extent possible, the
sessions, beginning with the same starting prompt, and using the same
follow-ups when applicable, but adjusting as necessary for different
issues that arose.
Evaluation: We can look at the number of tokens used and the number
of user turns required as measures of efficiency and native taste. Fewer
turns indicates the model arrives at what I consider a good result
faster, despite the underspecification of the initial prompt. The more
interesting part of the evaluation comes from actually experiencing the
interaction and the output and is mostly manual and qualitative. This
may be unappealing in comparison with automated benchmarks, or feel like
a cop-out. The goal is to replicate “product testing” however, which in
other domains, say running shoes or cars, ultimately boils down to using
something and seeing how well it works for you. One quantitative measure
that this method supports but was not attempted is actually timing how
efficiently and accurately the downstream data labeling task can be
completed using the different tools. This, combined with the development
time and manual steering required, could give a firmer quantitative
picture of how effective the different agents are.
In the qualitative evaluation we look for what stylistic choices were
made to support the tasks, how much intuition or common sense the agent
had about the task, and the apparent ease of use of the developed
solution. We also consider the overall experience and impression using
the agent. In running the evals, they felt very representative of the
kind of interactions I have with coding agents, including the debugging
side quests and need for steering.
Quantitatively, we can also measure how many turns were taken and how
many tokens were used, both proxies for efficiency.
Below are screenshots of the three applications generated with each
model
On the class labeling task, GLM needed to be told to add a sidebar,
took too long thinking and reached the default token limit twice, and
initially messed up the syntax highlighting, breaking the layout
whenever it highlighted a word. It particularly struggled when asked to
come up with a way of highlighting words and waffled on trying to make a
word list. It also made (in my view) a poor design choice in picking red
and green for the two class colors, given the good/bad connotation of
these two colors.
For the ranking task, GLM made near black-on-black interface to draw
the pictures at first which made it impossible to see. It also displayed
a bunch of metadata that had no relevance to the task.
On the editing task, GLM had some initial scroll and layout issues,
and displayed extraneous information – the instructions specify that
this is a test using just the python error message but it displayed
other info anyway.
Overall GLM generated easy to use applications and made good
stylistic choices. For example, it up front made the sentence large and
easy to read for the class labeling task. For the ranking task, there
was a but if busyness in the interface and the dark mode made it hard to
read, though I retrospectively noticed it added a light / dark toggle.
The editing interface requires scrolling to see the whole main panel
which makes it slower to use.
Qwen didn’t make a great use of space in the class labeling app, the
sentence to be reviewed was smaller than it needed to be. Similar to
GLM, Qwen didn’t initially add a scroll bar, and messed it up on its
first pass. It also had trouble with the layout when applying syntax
highlighting.
On the ranking task, Qwen didn’t fully finish according to what I
wanted but completed something that worked. It ran into errors that
needed multiple human-in-the-loop debugging rounds, first with loading
the pictures and then dragging and dropping. It also initially displayed
the pictures in black-on-black.
For the editing task, Qwen initially shared information it was not
supposed to with the user, required steering on the error highlighting,
had an error in the button states that needed manual correction
For the class labeling task, Laguna initially built a terminal
application. It did however respond to the prompt to create a web-app
but adding a sidebar automatically. It didn’t display the full sentences
initially which is poor judgement in an application specifically for
reviewing sentences. Laguna also picked red for one of the class
colors.
For the ranking task, Laguna made only minor judgement errors such as
displaying the

[truncated]
