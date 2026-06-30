---
source: "https://www.elliotcsmith.com/autoresearch-claude-and-constrained-optimization/"
hn_url: "https://news.ycombinator.com/item?id=48740094"
title: "Autoresearch, Claude and Constrained Optimization"
article_title: "Autoresearch, Claude and Constrained Optimization"
author: "smitec"
captured_at: "2026-06-30T23:27:32Z"
capture_tool: "hn-digest"
hn_id: 48740094
score: 1
comments: 0
posted_at: "2026-06-30T22:29:36Z"
tags:
  - hacker-news
  - translated
---

# Autoresearch, Claude and Constrained Optimization

- HN: [48740094](https://news.ycombinator.com/item?id=48740094)
- Source: [www.elliotcsmith.com](https://www.elliotcsmith.com/autoresearch-claude-and-constrained-optimization/)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T22:29:36Z

## Translation

タイトル: 自動リサーチ、クロードと制約付き最適化
説明: はじめに
AI を使用して何十人もの人々の仕事を行っているという主張を、遠くまで探す必要はありません。私は、証拠なしに改善について議論する主張には懐疑的な傾向があります。私はその疑念を受け止め、実践してみることにしました。これ

記事本文:
サインイン
購読する
エリオット・スミス著
で
AI
—
2026 年 7 月 1 日
自動リサーチ、クロード、制約付き最適化
AI を使用して何十人もの人々の仕事を行っているという主張を、遠くまで探す必要はありません。私は、証拠なしに改善について議論する主張には懐疑的な傾向があります。私はその疑念を受け止め、実践してみることにしました。これは、X に関する「ループ」の議論全体と若干の重複がありましたが、これは偶然です。
ここ数週間、私はカパルセイの「オートリサーチ」をテーマにしたプロジェクトをまとめてきました。私は、従来の機械学習や数値最適化の問題ではなく、ある程度の成功の客観的な尺度が得られる問題を選びたかったのです。
私がこの種の問題を選択したのは、私が取り組んできたプロジェクトや製品の多くがそのように構造化されているためです。変更 (増減) したいメトリクスがあり、理想的にはそれを測定する方法があります。おそらく、いくつかの制約もあるでしょう。この機能では、ページの読み込み時間が 500 ミリ秒を超えることはできません。
私はまだ、未知から成功への道が機械学習に似た明確な勾配最適化であるこのような問題に取り組んだことがありません。多くの場合、何らかの作業を完了し、それを「現実世界」でテストし、そのパフォーマンスを確認してから、次のステップについて決定します。すべての変更が良い結果をもたらすわけではなく、局所的に最適な結果をもたらすパスを深く掘り下げることは簡単です。
私は、ほとんど監視されていない方法で AI エージェントに大きな作業を任せる方法について直感を与える実験をしたかったのです。 Ralph Loops や現在クロード コードに組み込まれている /goal コマンドなど、この結果を達成しようとする他のメカニズムがすでに存在します。この設定の違いは、定量化可能な数値を主要な尺度として選択することです。

成功の可否を判断し、問題をいくつかの合否制約で制限しました。
物事を複雑にしすぎたくないので、ファイル圧縮の問題を選択しました。目的と制約がシンプルだったのでこれを選びました。最終的なファイル サイズが小さいほど、圧縮アルゴリズムが優れています。この問題に 2 つの制約を追加しました。1 つは非圧縮ファイルが完全に一致する必要があるということ、もう 1 つは圧縮も解凍も 300 秒を超えてはいけないということです。意図的に速度を最適化していませんでしたが、タイムアウトが発生して無限ループが発生することを承知の上で、時間を制限し、プロセスがほぼ監視なしで実行できるようにしたいと考えていました。
ファイル圧縮に関するもう 1 つの優れた点は、最終的なベンチマークに使用できる既存のツールが多数あることです。これが小規模な概念実証であったことを考えると、新しい最高級のアルゴリズムを作成するとは予想していませんでした。
それにもかかわらず、この自家製バージョンが既存のツールに対してどの程度優れたパフォーマンスを示したかを知ることは、私たちが図書館や既製のソリューションからどの程度離れていくかについてのデータポイントを提供するのにも役立ちます。以前は外部依存によって解決されていた問題をエージェントが迅速かつ確実に解決できるのであれば、社内ソリューションの価値がサプライ チェーン攻撃などのリスクを超える点があるはずです。これは 1 回の実験で答えが得られるものではありませんが、これをさらに調べる価値があるかどうかを判断するのに役立ちます。
まず、ここでの目標は、特定のモデルをベンチマークすることではなく、このアプローチが実行可能かどうかを確認することであったことを思い出してください。
次に、本題に入る前に、このプロジェクトのすべてのコードはここから入手できます: https://github.com/smitec/agent-compression
この作業では、Sonnet 4.6 のデフォルト設定で Claude Code を使用しました。モデルが異なれば、動作も異なるはずだと確信しています。

別の日に運動しましょう。
エージェントが関与する前に、プロジェクトの基本的な足場をセットアップしました。 Rust を選んだのは、「関数のシグネチャを変更しない」などの暗黙の制約の一部が型システムを介して簡単に強制できるためです。両方ともバイトをコピーするだけの圧縮関数と解凍関数のスタブをまとめました。これは「機能しました」が、どのデータにも圧縮がありませんでした。
次に、文字列と単純なファイルの両方に対する圧縮と解凍の往復をテストするために、いくつかの基本的な単体テストを導入しました。これらのテストは網羅的なものではありませんでしたが、圧縮および解凍機能がほぼ完璧な往復という目標を遵守していることを検証しました。
そこからベンチマーク スクリプトを作成しました。このスクリプトは、ビデオ、オーディオ、テキストにわたるいくつかのパブリック ドメイン ファイル サンプルを取得し、さまざまなサイズのランダム データで満たされたいくつかのファイルを作成しました。これらのファイルの多くは、すでにある程度圧縮されている形式であったため、圧縮率の低い形式に変換する手順を追加しました。これにより、全体的な圧縮ベンチマークと並んで、ファイルに関する優れたベンチマークが得られます。
このサンプル セットがあるということは、高エントロピー ファイル形式と低エントロピー ファイル形式が混在していることを意味します。優れた圧縮アルゴリズムでは、低エントロピー形式は縮小され、高エントロピー形式はほとんど変更されません。形式固有のバイトによりファイル サイズが多少変化することは予想されますが、全体的にはファイル サイズが意味のある形で増加することは望ましくありません。
サンプル セット内の最大のファイルは約 150MB でした。圧縮は、ファイルがさらに大きい場合にはより意味があると考えられますが、特に後のステップでテスト ループが非常に遅くなる可能性があります。
ベンチマーク スクリプトは各ファイルをループし、個別に圧縮してから解凍しました。スクリプトは解凍されたファイルをチェックしました

ile は元のファイルとビット単位で一致し、サイズの変更と圧縮および解凍の手順にかかる時間を記録しました。主に偶発的な無限ループをチェックするために、各ファイルのステップに 300 秒のタイムアウトが適用されました。
このスクリプトは、ファイルごとの変更の概要を説明する debug.csv ファイルを生成し、改善があった場合は、主要なメトリクスを results.csv ファイルに書き込みます。注目すべき点の 1 つは、結合された圧縮メトリックが (圧縮されたバイトの合計) / (元のバイトの合計) であることです。サンプルセット全体の平均圧縮率を取ることも検討しました。この選択の違いと影響については、後ほど説明します。
これらすべてのセットアップが完了したら、スタブ実装のベンチマークを実行し、実験を実行する準備ができていると判断しました。
物事を比較的適切に制御するために、各反復の前にクロード コンテキストをクリアし、モデルに「現在のコードベースを確認して、改善の別の反復を試みてください」というプロンプトを出しました。クロード コードをデフォルトでプラン モードに設定しているので、プランを待ち、簡単に確認した後、プランを受け入れ、エージェントを独自に実行させました。
この実験では、完全に自律的な選択ができるようにするため、意図的に計画を変更しませんでした。介入が役に立ったと思う場面も何度かありましたが、それは教訓になりました。
10 回の反復を実行し、データ固有の最適化を制御するために、いくつかの一般的な圧縮ツールと新しいデータセットに対する最終拡張ベンチマークを完了しました。これらの反復は約 2 週間にわたって実行され、通常は開始されて、他のことをしている間に実行され続けました。この延長された期間はトライアルの設計上の特徴ではなく、主に他の作業中にクロード コードの制限を使い果たしてしまうのを避けるためでした。

NG。
最初の反復中に、エージェントは、かなり標準的でよく知られた圧縮方法であるカスタム LZSS 実装を生成しました。次の 9 回の反復はこのメソッドの拡張であり、新しいエントロピー チェックとエントロピーの除去を試みるエンコード技術を追加しました。
各ループは、所要時間と使用されるトークンが大きく異なりました。クロード コードの /usage コマンドに基づくと、平均すると、1 回の反復のコストは約 4 米ドルです。繰り返しになりますが、これはデフォルト設定であったため、モデルごとに価格がどれだけ異なるかを考えると、価格についてはあまり深く読み込んでいません。
興味深いことに、モデルは特定の反復で複数セットの変更を行うことはありませんでした。仮説を立て、コードを追加し、ベンチマークを実行して、それ自体を「完了」と呼びます。これは、/goal コマンドを使用しないというプロンプト設定が原因である可能性があります。
以下の結果は、モデルが圧縮率を継続的に改善できたことを示しています。特に「圧縮可能」比率に注目すると、私の意見では、タスクがどれほど緩やかであったかを考えると、結果はかなり印象的でした。
最終結果を評価するために、同じデータセットに対して複数の圧縮ツールを実行しました。これらのツールが選択されたのは、たまたま既にインストールされていたためです。これはベンチマークを選択する最も堅牢な方法ではありませんが、一般的なツールとの比較を反映しています。
全体的に、カスタム アルゴリズムはかなり良好なパフォーマンスを示し、オーディオとビデオの圧縮では優れていましたが、他のカテゴリではわずかに劣るか、同等でした。最適化に使用される指標を考慮すると、オーディオとビデオのスコアが低いことは驚くべきことではありません。これらのファイル タイプは、ほとんどのバイトが圧縮されているため、合計スコアはそこでの勝利によって最も大きく変動しました。
このプロジェクトの目標に戻ると、これは画期的な圧縮アルゴリズムを見つけるという探求ではなく、タスクについての直感を養うことでした。

最適化ソフトウェアを備えたエージェント。
最後に、この投稿が長すぎる場合に読み飛ばしていただけるように、このプロジェクトからのハイレベルな要点をいくつか紹介します。全体として、最適化するための堅牢で測定可能で十分に制約された指標を見つけることができれば、この自動調査/ループ スタイルの作業は理にかなっていると思います。それらの 1 つを見つけるのは難しいことがよくあります。
セットアップを見たりレビューしたりしたときに私が感じた全体的な印象は、できるだけ早く「完了」したいということでした。これに基づいて、このセットアップの現実世界バージョンでは、明示的なループ メカニズムのセットアップが重要であると思います。
目的関数の選択が重要
私が得たもう 1 つの観察は、300 秒の時間パラメーターは制約があまりにも緩すぎる可能性があるということでした。これは変更によるマイナス面を抑えるのには役立ちましたが、モデルは圧縮のみを最適化していました。ミッチェル・ハシモト氏がこの X スレッドで最近捉えた現象:
フレーム時間 (および測定するテスト) を最小限に抑えることを目的として、レンダラーを最適化するループにエージェントを入れています。時間は 88 ミリ秒から 2 ミリ秒に短縮され、割り当ては約 150K から 500 に減少しました。いい感じですね?間違っている。これがまさにエージェント精神病が大きな問題である理由だ。
として…
このメソッドを実際に適用するには、より複雑な「スコア」を最適化するか、後で速度向上のための最適化に切り替える必要があります。コードの長さ、メモリ使用量などの他の二次的な指標についても同じことが言えます。
これは決して新しい問題ではなく、エージェント ベースのコーディングに特有の問題でもありません。 「成功」と「完了」の尺度を選択することは、エンジニアリング組織において長年の課題でした。現実的には、どのようなメトリクスまたはメトリクスの組み合わせにもトレードオフが伴います。おそらく、その事実に慣れ、時間の経過とともに焦点を移す必要があるだけです。

変化が必要です。
最近、PostHog が新しい PostHog Code 製品でこの分野で何らかの取り組みを行っているのを知りました。ユーザーがコーディング エージェントのコンテキストに製品分析を組み込んで、より適切な意思決定を行えるようにします。まだテストしていませんが、方向性は正しいと感じています。
現実世界の目標はこれほど簡単に測定できるものではありません
メトリクスについて議論する際には、この手法が「現実世界」でどのように異なるかを考慮する価値があります。圧縮ツールには非常に高速なフィードバック ループがあります。ファイルを取得し、圧縮、解凍し、結果を比較できます。この変更がより広範なものである場合、たとえば「チェックアウトのコンバージョン率を向上させる」場合は、サンプルの収集にさらに多くの時間が必要となり、データ内のノイズの影響をさらに受けやすくなります。
ここでの 1 つの解決策は、コンバージョン率が向上するという期待/仮説を立てて、プロキシ メトリクスを最適化することです。それは、「ページの読み込み速度の向上」や「チェックアウトに必要なクリック数の削減」などです。これは確かにもっと簡単に反復できるかもしれませんが、最終目標との相関が緩やかな代理メトリクスに基づいて過剰に最適化する危険性があります。これは、より複雑なメトリクスと完全かつ線形に相関するプロキシ メトリクスを見つける速度です。
それで

[切り捨てられた]

## Original Extract

Introduction
You don't need to look far to find claims that folks have been using AI to do the work of dozens of people. I tend to be skeptical of any claim that discusses improvements without evidence. I decided to take that skepticism and put it to work. This

Sign in
Subscribe
By Elliot Smith
in
AI
—
01 Jul 2026
Autoresearch, Claude and Constrained Optimization
You don't need to look far to find claims that folks have been using AI to do the work of dozens of people. I tend to be skeptical of any claim that discusses improvements without evidence. I decided to take that skepticism and put it to work. This had a minor overlap with the whole 'loops' discussion on X but that's coincidental.
Over the last few weeks I have put together a project in the theme of Kaparthay's 'Autoresearch'. I wanted to choose a problem that was not a traditional machine learning or numerical optimization problem but one that still had some objective measure of success.
I chose this kind of problem because many of the projects or products I have worked on are structured that way. You have some metric that you want to change (up or down) and ideally some way to measure it. You likely also have some constraints e.g. we can't let the page load time exceed 500ms for this feature.
I have yet to work on a problem like this where the path from unknown to success is a clear, gradient optimization akin to machine learning. More often you complete some work, test it in the 'real world', look at how it performed and then make a decision about next steps. Not all changes result in a positive outcome and it's easy to go deep down a path that results in a locally optimal outcome.
I wanted an experiment that would give me some intuition about how to task AI agents with bigger pieces of work in a mostly unsupervised way. There are already other mechanisms to try and achieve this outcome, such as Ralph Loops and the /goal command that's now in Claude Code. The difference in this setup is that I would pick a quantifiable number as the primary measure of success and bound the problem with some pass-fail constraints.
Not wanting to over complicate things I chose the problem of file compression. I picked it because the objective and the constraints were simple. A compression algorithm is better if the final file size is smaller. I added two constraints to the problem, one being that the uncompressed file needed to match perfectly and the other that neither compression or decompression could exceed 300 seconds. I was deliberately not optimizing for speed but wanted to cap the time and ensure the process could run mostly unsupervised with the knowledge that a timeout would catch and infinite loops.
The other nice thing about file compression is that there are many existing tools I could use for a final benchmark. Given this was a small proof of concept I wasn't expecting to create a new top-of-the-line algorithm.
Despite that, knowing how well this home cooked version performed against existing tools also helps provide a data point on how much we might move away from libraries and off the shelf solutions. If an agent can quickly and reliably solve a problem previously solved by an external dependency there must be some point at which the value of an in house solution exceeds the risk of things like supply chain attacks. This isn't something one single experiment would answer but it would help determine if this was worth looking at more.
First, a reminder that the goal here was to see if this approach was viable rather than to benchmark any particular model.
Second, before we get into it, all the code for this project is available here: https://github.com/smitec/agent-compression
For this work I used Claude Code with default settings on Sonnet 4.6. I am certain different models would have done things differently, that's an exercise for another day.
Prior to any agent involvement I setup a basic scaffold for the project. I picked Rust because some of the implicit constraints like "don't modify the function signature" were easily enforceable via the type system. I put together a stub of the compress and decompress function which both just copied the bytes across. This 'worked' but provided zero compression to any of the data.
I then put in place a couple of basic unit tests to test the compress-decompress round trip on both a string and a simple file. These tests weren't exhaustive but did validate that the compress and decompress function were adhering to their goal of a bit perfect round trip.
From there I put together a bench-marking script. This script fetched some public domain file samples across video, audio and text as well as created some files filled with random data of various sizes. Many of these files were in formats that were already somewhat compressed so I added a step to convert them to less compressed formats. This gives a good file wise benchmark alongside the overall compression benchmarks.
Having this sample set meant that there were a mix of high and low entropy file formats. A good compression algorithm will shrink low entropy formats and leave high entropy formats mostly unchanged. You can expect some minor change in file size due to format specific bytes but overall you don't want file size to increase in a meaningful way.
The largest file in the sample set was around 150MB. While compression is likely more meaningful on even larger files it would have resulted in a very slow test loop, especially in later steps.
The bench-marking script looped through each of the files, compressed them individually and then decompressed them. The script checked the decompressed file was a bitwise match to the original and noted down the change in size and how long the compress and decompress steps took. There was a 300 second timeout applied to each file's steps mainly to check for accidental infinite loops.
The script produced a debug.csv file which outlined the changes per file and, if there was an improvement, wrote the key metrics to a results.csv file. One thing of note was that the combined compression metric was (total compressed bytes) / (total original bytes) . I had also considered taking the average percentage compression across the sample set. I'll get into the differences and impact of this choice a little later.
Once all of this was setup I ran the benchmark for the stub implementation and considered the experiment ready to run.
To keep things relatively well controlled I cleared the Claude context before each iteration and prompted the model with "Review the current codebase and attempt another iteration of improvement." I have Claude Code set to plan mode by default so I waited for the plan and then after a quick review accepted the plan and let the agent run on its own.
I intentionally didn't modify any of the plans in this experiment, wanting to let it make fully autonomous choices. There were a few times where I think an intervention would have been useful but that’s a lesson learned.
I ran ten iterations and then completed a final extended benchmark against some common compression tools and on a new dataset to control for any data-specific optimization. These iterations were run over the course of about two weeks usually kicked off and left to run while I was doing other things. This extended time period wasn't a design feature of the trial, it was mostly to avoid exhausting my Claude Code limits while working on other things.
During the first iteration the agent produced a custom LZSS implementation, a fairly standard and well known method of compression. The next nine iterations were extensions to this method, adding new entropy checks and encoding techniques to try and remove entropy.
Each loop varied a lot in time taken and tokens used. On average, based on the /usage command in Claude Code, a single iteration cost about $4 USD. Again this was on the default settings so I am not reading too much into the price given how much that varies per model.
Interestingly the model never made more than one set of changes in a given iteration. It would form a hypothesis, add the code, run the benchmark and call itself 'complete'. This may come down to the prompting setup of not using the /goal command.
The results below show that the model was able to continue to make improvements to the compression factor. Looking in particular at the 'compressible' ratio the results were, in my opinion, pretty impressive given how loose the task was.
To assess the final results I ran several compression tools over the same dataset. These tools were chosen because they happened to already be installed. This is not the most robust method of choosing a benchmark but it does reflect a comparison to common tooling.
Overall the custom algorithm performed fairly well, it excelled at audio and video compression and was slightly worse or on par in other categories. The lower scores in audio and video aren't surprising given the metric used to optimise. These file types represented most of the bytes being compressed so the combined score was moved most by wins there.
Coming back to the goal of this project, this wasn't a quest to find a breakthrough compression algorithm but instead to develop some intuition about tasking an agent with optimizing software.
To wrap this up, and give folks something to skip to if this post is too long, here are some high level take-aways from this project. Overall, I think if you can find a robust, measurable and well constrained metric to optimise then this auto-research/loop style work makes sense. Finding one of those is often tricky.
The overall feeling I had while watching/reviewing the setup was that it wanted to be 'done' as quickly as possible. Based on this I think having some explicit looping mechanism setup would be important for a real world version of this setup.
The choice of objective function is key
Another observation I had was that the 300 second time parameter was likely far too loose a constraint. It was useful for capping the downside of a change but the model was only ever optimising for compression. A phenomenon recently captured by Mitchell Hashimoto in this X thread:
I've got an agent in a loop optimizing a renderer with the goal to minimize frame times (and tests to measure). It got times down from 88ms to 2ms and allocations down from ~150K to 500. Sounds good, right? Wrong. This is exactly why agent psychosis is a big fucking problem.
As…
A real world application of this method would either need a more complex 'score' to optimize or to later switch to an optimisation for speed. The same can be said for other secondary metrics like code length, memory usage etc.
This is by no means a new issue or one that is unique to agent based coding. Choosing measures of 'success' and 'done' has long been a challenge in engineering organisations. Realistically any metric or combination of metrics is going to come with trade offs. You probably just need to get comfortable with that fact and be willing to shift your focus over time as the needs change.
I saw recently that PostHog was doing some work in this space with their new PostHog Code product. Allowing users to bring product analytics into their coding agent context to better guide decisions. I'm yet to test it out but it feels like the right direction.
Real world objectives are rarely as simple to measure
While discussing metrics it's worth considering how this technique might differ in the 'real world'. A compression tool has a very fast feedback loop. You can take a file, compress it, decompress it and compare the results. If this change was more broad, say "Improve the checkout conversion rate," you'd need a lot more time to gather samples and you'd be a lot more susceptible to noise in the data.
One solution here is to optimize a proxy metric with the hope/hypotheses that it will improve the conversion rate. That might be something like 'improve page load speed' or 'reduce the number of clicks needed to checkout'. This could certainly be more easily iterated on but you then run the risk of over-optimising on a proxy metric that only loosely correlates with your final goal. It is rate to find a proxy metric that is perfectly and linearly correlated with a more complex one.
So

[truncated]
