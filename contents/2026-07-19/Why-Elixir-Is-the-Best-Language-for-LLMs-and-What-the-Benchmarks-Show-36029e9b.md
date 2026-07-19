---
source: "https://codemyspec.com/blog/why-elixir-is-the-best-language-for-llms"
hn_url: "https://news.ycombinator.com/item?id=48970657"
title: "Why Elixir Is the Best Language for LLMs (and What the Benchmarks Show)"
article_title: "Why Elixir Is the Best Language for LLMs"
author: "ksec"
captured_at: "2026-07-19T18:51:48Z"
capture_tool: "hn-digest"
hn_id: 48970657
score: 1
comments: 0
posted_at: "2026-07-19T18:42:58Z"
tags:
  - hacker-news
  - translated
---

# Why Elixir Is the Best Language for LLMs (and What the Benchmarks Show)

- HN: [48970657](https://news.ycombinator.com/item?id=48970657)
- Source: [codemyspec.com](https://codemyspec.com/blog/why-elixir-is-the-best-language-for-llms)
- Score: 1
- Comments: 0
- Posted: 2026-07-19T18:42:58Z

## Translation

タイトル: Elixir が LLM に最適な言語である理由 (およびベンチマークが示すもの)
記事のタイトル: Elixir が LLM にとって最適な言語である理由
説明: AutoCodeBench は 20 の言語をテストしました。LLM は他のどの言語よりも Elixir をうまく解決し、ほぼすべてのフロンティア モデルでトップの言語です。その理由は次のとおりです。

記事本文:
Elixir が LLM にとって最適な言語である理由
コードマイスペック
開発者
Elixir が LLM に最適な言語である理由 (およびベンチマークが実際に示すもの)
これまでに公開された最大の多言語コード生成ベンチマークでは、大規模な言語モデルが、テストされた他の 19 言語よりも高い率で Elixir の問題を解決します。それがショートバージョンです。 Elixir は低リソース言語であるため、長いバージョンの方が興味深いです。インターネットには、Python や JavaScript 用の Elixir コードの一部が保存されているため、モデルが学ぶべきことははるかに少なくなりました。とにかく彼らは先に出てきました。この記事では、ベンチマークで何が見つかったのか、なぜそれが起こるのか、主張がどこで密閉性を失うのか、そして AI を使用してソフトウェアを構築する場合にそれが何を意味するのかを説明します。
ベンチマークは、Tencent Hunyuan が公開している AutoCodeBench です。これには 20 の言語に均等に分散された 3,920 の問題が含まれており、それぞれ約 196 の問題があり、生成されたコードを目視ではなくサンドボックスで実行することによって解答を採点します。 30 を超えるモデルが、推論モードと非推論モードの両方で、オープンおよびクローズドで評価されました。 Tencent は Elixir に株式を持っていないが、それがこの結果が注目に値する理由の 1 つである。
1 つ目は和集合の上限です。少なくとも 1 つのモデルが解決したすべての問題を取り上げ、各言語がどの程度の割合をカバーしているかを尋ねます。 Elixir は 97.5% であり、全 20 言語の中で最も高く、その差は大きいです。
地球上で最も多くのトレーニング データを持つ言語である Python は、63.3 で最下位に終わりました。
2 番目の数値は、結合ではなく個々のモデルに注目しているため、より分かりやすくなります。ほぼすべてのフロンティア モデルにおいて、Elixir は単一言語の中で最も高いスコアを獲得しており、全体の平均より約 30 ポイントの差があります。
その表のどのモデルも最善を尽くしますw

エリクサーのオルク。推論以外の結果も同じことを物語っています。Claude Opus 4 は、推論をオフにした Elixir で 82.3 点を獲得し、やはりトップ言語となっています。したがって、この効果は、1 つのラボや 1 つのプロンプト モードに特有のものではありません。これは、ラボ全体、モデル サイズ全体、および両方のモードにわたって表示されます。
問題となるのは、トレーニング データの角度です。言語モデルは、Python と JavaScript の海と Elixir の水たまりを見たものから学習します。コーパスのサイズがパフォーマンスを左右するのであれば、Python が勝利し、Elixir は他のニッチな言語とともに最下位近くに位置するでしょう。逆のことが起こりました。 Elixir は、トレーニング データの約 10 倍で言語を上回りました。何が起こっているとしても、それは「モデルがより多くの Elixir を覚えた」ということではありません。
データのもう 1 つのスライスは、これがチャット チューニングのアーティファクトであるという簡単な反論を排除します。 AutoCodeBench には、人間のフィードバックからの命令調整や強化学習の前に、生の基本モデルを 3 ショット プロンプトで評価するバリアント Complete が含まれています。 Elixir も上半分に位置しています。DeepSeek-Coder-V2-Base のスコアは 52.0、Seed-Coder-8B-Base のスコアは 48.0 です。もしその利点が、研究室が調整中に磨き上げたものであれば、そのステップを経なかったモデルには現れないでしょう。とにかく表示され、言語自体を指します。
ベンチマークを見ると、Elixir が勝っていることが分かります。理由はわかりません。 Elixir の作成者である José Valim は、Dashbit で最も明確な説明を書きました。彼の主張は、結果は運によるものではないというものです。それは、言語モデルが必要とするものと偶然一致する設計の選択肢から外れます。そのうちの5つが重要です。
不変性は推論をローカルに保ちます。 Elixir では、データが関数に入力され、新しい値が出力されます。あなたの背後で変化するものは何もありません。ヴァリムは隠されたものが存在しないと呼ぶ

副作用は「遠くから不気味なアクションがない」こと、そしてゲーム全体である有限のコンテキスト ウィンドウを持つモデルの場合です。関数の動作を確認するためにモデルが取り込む必要がある周囲のコードが少なくなるほど、モデルが間違っている可能性が低くなります。パイプ演算子により、データ フローが 1 行で表示されます。
"Elixir は素晴らしいです" |> String.split() |> Enum.frequency() その透明度はモジュール境界までスケールアップします。これは、Phoenix コンテキストが LLM に適している理由について私が別途作成した議論です。適切に描画されたコンテキストは、アプリの残りの部分をロードせずにモデルが推論できる小さな自己完結型のサーフェスです。
文書は検証された契約書です。 Elixir は、@doc 属性で記述されたパブリック コントラクトを通常の実装コメントから分離します。これらのドキュメント内のサンプルは実行可能であり、テスト スイートの一部として実行されます。ドキュメントがコードと同期しなくなるとビルドが失敗するため、ドキュメントがコードと同期しなくなることはありません。それに加えて、エコシステムはバージョン化されたインデックス付き検索を備えた HexDocs を中心に集中化されています。実際の効果は、モデルがトレーニングされたテキストが正しく、エージェントが実行時に取得するテキストが正しいということです。
10 年間の安定性により信号をクリーンに保ちます。 Elixir は 2014 年からバージョン 1.x を使用しています。Phoenix は、下のアプリを壊すことなく 1.0 から 1.8 に移行しました。 Ecto は 2018 年から 3.x です。5 年または 8 年前に書かれたコードは今でも実行されます。これは思っている以上に重要です。同じことを行うための 3 つの相互に互換性のない方法によってトレーニング データが汚染されていないことを意味するからです。古い API が依然として最新の API であるため、モデルはどの API が最新であるかについて混乱することはありません。何かが非推奨になると、コンパイラは移行パスを含む警告を発行し、エージェントはこれに直接対処できます。
フィードバック

k ループは高速で、正直で、ノンブロッキングです。 Elixir はコンパイルし、コンパイラーは未使用の変数、未定義の関数、および到達不能な句にフラグを立てます。最近の型推論の作業では、誤報をほとんど発生させずに、これらをより多く検出します。重要なのは、それらはハードストップではなく警告として到着するため、エージェントは作業を続けて、最初の問題で停止するのではなく、パス内でそれらをクリーンアップすることができます。コンパイルとテストは両方ともコア間で並行して実行されます。エージェントは、順調に進んでいるかどうかについて、信頼できる厳密な信号を受け取ります。
操作面は小さく、実行時間は透過的です。 Elixir が Erlang と共有する仮想マシンである BEAM は、同時実行機能、ディストリビューション、パブリッシュ/サブスクライブ機能を同梱しています。典型的な Phoenix アプリケーションは、フレームワークにデータベースを加えたものであり、無秩序に広がる Redis、メッセージ キュー、およびエージェントが頭に保持する 3 つのサーバーレス機能ではありません。開発は生産を反映します。そして、ランタイムは内省可能です。すべての軽量プロセスはその状態と現在の作業を公開し、Phoenix LiveDashboard またはコードを通じて表示できます。 MCP サーバーである Tidewave は、実行中のアプリケーションをコーディング エージェントに渡し、ソースを読んで推測するのではなく、デバッグ中にライブ状態を検査できるようにします。
ヴァリム氏のさらに大きな点は、これらが複合しているということだ。不変性は、モデルが正しくなければならないコンテキストを低下させます。検証済みのドキュメントと長期にわたる安定性により、トレーニング信号がクリーンで最新の状態に保たれます。迅速なフィードバックとライブ イントロスペクションにより、エージェントは間違いが発生した場合にすぐに修正できます。最初のドラフトを適切に作成し、そこから適切に反復することは、まさにエージェント コーディングの形です。
これを信じていただきたいのですが、この主張が気密性を失うのはここです。
一つのベンチマークです。言語あたり 196 の問題は実際のサンプルであり、巨大なものではありません。問題自体は LLM と SA によって生成されました。

ndbox ループは、特定のスタイルの問題に有利に働く可能性がある方法です。 20 の言語セットのすべてが同じ手法で構築されているため、言語間の比較は公正ですが、「このベンチマークで優勝した」ということは、「史上最高であることが証明された」ということと同じではありません。結果を、終了した事件ではなく、強力で収束する証拠として扱います。
主張を伝える際にも注意が必要です。 Elixir はテーブルのすべてのセルに勝つわけではありません。弱いモデルの場合、C# や Shell などの言語の方が平均スコアが高くなることがよくあります。これはおそらく、モデルに記憶すべき定型文が多すぎるためと考えられます。精査の下でも有効である 2 つの結果は、上記の結果です。結合の上限 (97.5 パーセント、快適な 1 位) とフロンティア モデル パターン (Elixir は、Opus 4、Sonnet 4、o3、o4-mini、Grok-4、および DeepSeek-R1 のトップ言語です)。最先端のモデルの 1 つである Gemini 2.5 Pro は、C# で Elixir を追い抜いています。 「エリクサーはあらゆる場所であらゆるものに勝る」という話を平板にすれば、誰かがそのコラムを引っ張り出してくるでしょう、そして彼らは正しいでしょう。正確なバージョンは強力です。
次に、標準的な反論: Elixir はニッチであり、雇用するのは難しいです。それは人間にとっても当てはまりますが、エージェントにとっては逆です。モデルには雇用の問題がありません。小規模で安定した一貫性のある言語は、競合する 10 個の方言を含む大規模な言語よりも文脈に保持しやすいです。 Elixir を人材派遣マネージャーに売り込みにくくする特性は、コーディング エージェントにとって売り込みやすくする特性でもあります。
より優れた言語で得られるもの、得られないもの
何を使って構築するかを決定する場合に重要となる部分です。言語が改善されると、最初の草稿のレベルが上がります。この証拠に基づいて、Elixir を作成するエージェントは、Python を作成する同じエージェントよりも優れたところから開始します。それは本物であり、持つ価値があります。
それは

仕事全体ではありません。きれいな初稿は正しいアプリケーションではありません。モデルは依然として、正しく見えても微妙に間違っているコードを記述します。ターゲット、つまり当面の機能の「完了」が何を意味するかを書面で定義する必要があり、実行中のアプリケーションが実際に意図したとおりに動作するかどうかのチェックも必要です。言語によってドラフトがより良くなります。それはシステムを正しくするものではありません。このギャップは、Python を作成する場合でも Elixir を作成する場合でも存在します。 Elixir は、クロージングを開始するためのより良いドラフトを提供します。
CodeMySpec が利点をさらに高める方法
そのギャップを埋めるのが CodeMySpec です。 CodeMySpec は、Phoenix および Elixir 用に構築された仕様主導の AI 開発ハーネスであり、文字通りの意味で Elixir ネイティブです。Phoenix コンテキスト、LiveView、Ecto、および OTP はファーストクラスであり、Elixir モードが追加された汎用テンプレートではありません。スタック汎用ツールは AutoCodeBench フロア上に構築されていないため、AutoCodeBench フロア上に立つことはできません。これは、Elixir のスペック駆動開発の最終ラインでもあります。
その上に、CodeMySpec は言語では提供できない 2 つの機能を追加します。
ターゲットです。すべての機能は、Spex DSL の BDD シナリオとして記述された動作仕様として始まり、これらの仕様はオプションのドキュメントではなく必須です。モデルは、あなたが何を言おうとしているのかを推測しません。あなたが書き留めた合格基準に向かって構築されています。要件グラフは、あらゆる成果物、仕様、テスト、実装、結果を追跡し、次に何を取り組むべきかを計算します。
小切手。これは、このカテゴリ内の他のツールでは実行できない部分です。単体テストに合格し、動作仕様に合格した後、QA エージェントは実際のアプリケーションを起動し、実際のブラウザを起動し、実際のボタンをクリックし、表示される内容のスクリーンショットを撮り、実行中のアプリケーションが誤動作した場合に重大度の問題をファイルに記録します。

単体テストに合格し、仕様に合格したら、QA エージェントがボタンをクリックしてとにかくバグを見つけます。これは実際のアプリケーションでの検証であり、コードが正しく見えることを保証するものではありません。
また、トークンのマークアップなしで、独自のエージェント、独自のモデル、独自のキーを持ち込むことができます。つまり、今日から AutoCodeBench の結果に基づいて行動できるということです。Claude Opus 4 (Elixir で唯一最高の仕事をするモデル) を、ハーネス内の Elixir コードベースに指定して、目標とする仕様と、まだ間違っている点を見つけるためのライブ QA パスを提供します。
Elixir のおかげでモデルはドラフトを正しく行うことができます。このアプリが正しい理由はハーネスにあります。
Phoenix コンテキストが LLM に最適な理由
Elixir と Phoenix の仕様主導型開発
2026 年のスペック駆動開発: 完全ガイド
AutoCodeBench: 大規模言語モデルは自動コード ベンチマーク ジェネレーター (Tencent Hunyuan) です。論文: https://arxiv.org/abs/2508.09101
AutoCodeBench プロジェクト ページとリーダーボード: https://autocodebench.github.io/
AutoCodeBench コード: https://github.com/Tencent-Hunyuan/AutoCodeBenchmark
AutoCodeBench データセット: https://huggingface.co/datasets/tencent/AutoCodeBenchmark
José Valim / Dashbit、「Elixir が AI に最適な言語である理由」: https://dashbit.co/blog/why-elixir-best- language-for-ai
アッラ

[切り捨てられた]

## Original Extract

AutoCodeBench tested 20 languages: LLMs solve Elixir better than any other, and it is the top language for nearly every frontier model. Here is why.

Why Elixir Is the Best Language for LLMs
CodeMySpec
Developers
Why Elixir Is the Best Language for LLMs (and What the Benchmarks Actually Show)
On the largest multi-language code generation benchmark published so far, large language models solve Elixir problems at a higher rate than any of the other nineteen languages tested. That is the short version. The longer version is more interesting, because Elixir is a low-resource language. The internet holds a fraction of the Elixir code it holds for Python or JavaScript, so the models had far less to learn from. They came out ahead anyway. This post walks through what the benchmark found, why it happens, where the claim stops being airtight, and what any of it means if you build software with AI.
The benchmark is AutoCodeBench, published by Tencent Hunyuan. It contains 3,920 problems spread evenly across 20 languages, roughly 196 problems each, and it grades answers by running the generated code in a sandbox rather than by eyeballing it. More than 30 models were evaluated, open and closed, in both reasoning and non-reasoning modes. Tencent has no stake in Elixir, which is part of what makes the result worth paying attention to.
The first is the union upper bound: take every problem that at least one model solved, and ask what share of each language that covers. Elixir sits at 97.5 percent, the highest of all 20 languages, and the gap is wide.
Python, the language with more training data than any other on earth, lands last at 63.3.
The second number is more telling, because it looks at individual models rather than the union. For nearly every frontier model, Elixir is the single language it scores highest on, by a margin of roughly 30 points over its own overall average.
Every model in that table does its best work in Elixir. The non-reasoning results tell the same story: Claude Opus 4 scores 82.3 on Elixir with reasoning turned off, again its top language. The effect, then, is not a quirk of one lab or one prompting mode. It shows up across labs, across model sizes, and across both modes.
The part that should stop you is the training-data angle. Language models learn from what they have seen, and they have seen an ocean of Python and JavaScript and a puddle of Elixir. If corpus size drove performance, Python would win and Elixir would sit near the bottom with the other niche languages. The opposite happened. Elixir beat languages with roughly ten times its training data. Whatever is going on, it is not “the model memorized more Elixir.”
One more slice of the data rules out the easy counter-explanation that this is a chat-tuning artifact. AutoCodeBench includes a variant, Complete, that grades raw base models with three-shot prompting, before any instruction tuning or reinforcement learning from human feedback. Elixir lands in the upper half there too: DeepSeek-Coder-V2-Base scores 52.0 on it, Seed-Coder-8B-Base 48.0. If the advantage were something the labs polished in during alignment, it would not show up in models that never went through that step. It shows up anyway, which points back at the language itself.
The benchmark tells you Elixir wins. It does not tell you why. José Valim, the creator of Elixir, wrote the clearest explanation at Dashbit, and his argument is that the result is not luck. It falls out of design choices that happen to line up with what a language model needs. Five of them matter.
Immutability keeps reasoning local. In Elixir, data goes into a function and a new value comes out. Nothing mutates behind your back. Valim calls the absence of hidden side effects “no spooky action at a distance,” and for a model with a finite context window that is the whole game. The less surrounding code the model has to pull in to be sure of what a function does, the fewer chances it has to be wrong. The pipe operator makes the data flow visible on one line:
"Elixir is awesome" |> String.split() |> Enum.frequencies() That transparency scales up to the module boundary too, which is the argument I made separately about why Phoenix contexts are good for LLMs : a well-drawn context is a small, self-contained surface a model can reason about without loading the rest of the app.
Documentation is a verified contract. Elixir separates the public contract, written with the @doc attribute, from ordinary implementation comments. The examples inside those docs are runnable, and they execute as part of the test suite. The documentation cannot quietly drift out of sync with the code, because the build fails when it does. On top of that, the ecosystem centralizes on HexDocs with versioned, indexed search. The practical effect is that the text the models trained on was correct, and the text an agent retrieves at runtime is correct.
A decade of stability keeps the signal clean. Elixir has been on version 1.x since 2014. Phoenix went from 1.0 to 1.8 without breaking the app underneath you. Ecto has been on 3.x since 2018. Code written five or eight years ago still runs. That matters more than it sounds, because it means the training data is not polluted with three mutually incompatible ways to do the same thing. The model is not confused about which API is current, because the old one is still the current one. When something does get deprecated, the compiler emits a warning with a migration path, which an agent can act on directly.
The feedback loop is fast, honest, and non-blocking. Elixir compiles, and the compiler flags unused variables, undefined functions, and unreachable clauses. The recent type inference work catches more of these with few false alarms. Crucially, they arrive as warnings rather than hard stops, so an agent can keep working and clean them up in a pass instead of grinding to a halt on the first issue. Compilation and tests both run in parallel across cores. The agent gets a tight, trustworthy signal about whether it is on track.
The operational surface is small, and the runtime is transparent. The BEAM, the virtual machine Elixir shares with Erlang, ships concurrency, distribution, and pub/sub in the box. A typical Phoenix application is the framework plus a database, not a sprawl of Redis, a message queue, and three serverless functions for an agent to hold in its head. Development mirrors production. And the runtime is introspectable: every lightweight process exposes its state and current work, viewable in Phoenix LiveDashboard or through code. Tidewave, an MCP server, hands a running application to a coding agent so it can inspect live state while it debugs, rather than reading source and guessing.
Valim’s larger point is that these compound. Immutability lowers the context a model needs to be correct. Verified docs and long stability keep the training signal clean and current. Fast feedback and live introspection let an agent correct itself quickly when it does slip. Getting the first draft right and iterating well from there is the exact shape of agentic coding.
I want you to trust this, so here is where the claim stops being airtight.
It is one benchmark. 196 problems per language is a real sample, not an enormous one, and the problems were themselves generated by an LLM and sandbox loop, a method that could favor certain styles of problem. The same method built every one of the 20 language sets, so the comparison between languages is fair, but “won this benchmark” is not the same as “proven best for all time.” Treat the result as strong, converging evidence, not a closed case.
The claim also needs care in the telling. Elixir does not win every cell of the table. For weaker models, languages like C# and Shell often score higher on average, probably because there is so much boilerplate for the model to have memorized. The two findings that hold up under scrutiny are the ones above: the union upper bound (97.5 percent, comfortably first) and the frontier-model pattern (Elixir is the top language for Opus 4, Sonnet 4, o3, o4-mini, Grok-4, and DeepSeek-R1). One frontier model breaks the pattern, Gemini 2.5 Pro, which edges Elixir out with C#. Flatten the story to “Elixir beats everything everywhere” and someone will pull up that column, and they will be right to. The precise version is the strong one.
Then the standard objection: Elixir is niche, and hiring for it is hard. That is true for humans, and it cuts the other way for agents. A model does not have a hiring problem. A small, stable, consistent language is easier for it to hold in context than a large one with ten competing dialects. The property that makes Elixir a harder sell to a staffing manager is the property that makes it easier for a coding agent.
What a better language buys you, and what it does not
Here is the part that matters if you are deciding what to build with. A better language raises the floor of the first draft. On this evidence, an agent writing Elixir starts from a better place than the same agent writing Python. That is real, and it is worth having.
It is also not the whole job. A cleaner first draft is not a correct application. The model still writes code that looks right and is subtly wrong. It still needs a target, a written definition of what “done” means for the feature at hand, and it still needs a check that the running application actually behaves the way you intended. The language makes the draft better. It does not make the system correct. That gap exists whether you write Python or Elixir. Elixir hands you a better draft to start closing it from.
How CodeMySpec compounds the advantage
Closing that gap is where CodeMySpec sits. CodeMySpec is a specification-driven AI development harness built for Phoenix and Elixir, and it is Elixir-native in the literal sense: Phoenix contexts, LiveView, Ecto, and OTP are first-class, not a generic template with an Elixir mode bolted on. Stack-generic tools cannot stand on the AutoCodeBench floor, because they are not built on it. This is also the through-line of spec-driven development for Elixir .
On top of that floor, CodeMySpec adds the two things the language cannot give you.
A target. Every feature starts as a behavioral specification, written as BDD scenarios in the Spex DSL, and those specs are mandatory rather than optional documentation. The model is not guessing what you meant; it is building toward acceptance criteria you wrote down. A requirement graph tracks every artifact, the spec, the tests, the implementation, and the results, and computes what to work on next.
A check. This is the part no other tool in the category does. After the unit tests pass and the behavioral specs pass, a QA agent boots the real application, drives a real browser, clicks the actual buttons, screenshots what it sees, and files issues with severity when the running app misbehaves. Unit tests pass, specs pass, and then the QA agent clicks the button and finds the bug anyway. That is verification on the live application, not a promise that the code looked correct.
And you bring your own agent, your own model, and your own keys, with no token markup. That means you can act on the AutoCodeBench result today: point Claude Opus 4, the model that does its single best work in Elixir, at an Elixir codebase, inside a harness that gives it a spec to aim at and a live QA pass to catch what it still gets wrong.
Elixir is why the model gets the draft right. The harness is why the app is right.
Why Phoenix Contexts Are Great for LLMs
Spec-Driven Development for Elixir and Phoenix
Spec-Driven Development in 2026: The Complete Guide
AutoCodeBench: Large Language Models are Automatic Code Benchmark Generators (Tencent Hunyuan). Paper: https://arxiv.org/abs/2508.09101
AutoCodeBench project page and leaderboard: https://autocodebench.github.io/
AutoCodeBench code: https://github.com/Tencent-Hunyuan/AutoCodeBenchmark
AutoCodeBench dataset: https://huggingface.co/datasets/tencent/AutoCodeBenchmark
José Valim / Dashbit, “Why Elixir is the Best Language for AI”: https://dashbit.co/blog/why-elixir-best-language-for-ai
Alla

[truncated]
