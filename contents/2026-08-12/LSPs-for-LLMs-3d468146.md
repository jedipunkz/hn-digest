---
source: "https://ianbarber.blog/2026/08/11/lsps-for-llms/"
hn_url: "https://news.ycombinator.com/item?id=49267518"
title: "LSPs for LLMs"
article_title: "LSPs for LLMs – Ian’s Blog"
author: "matt_d"
captured_at: "2026-08-12T03:55:47Z"
capture_tool: "hn-digest"
hn_id: 49267518
score: 1
comments: 0
posted_at: "2026-08-12T03:26:10Z"
tags:
  - hacker-news
  - translated
---

# LSPs for LLMs

- HN: [49267518](https://news.ycombinator.com/item?id=49267518)
- Source: [ianbarber.blog](https://ianbarber.blog/2026/08/11/lsps-for-llms/)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T03:26:10Z

## Translation

タイトル: LLM 用の LSP
記事のタイトル: LLM のための LSP – Ian のブログ
説明: エディターにコードを入力していた暗黒時代に戻ると、壊れたコードの下の波線、クリックして定義リンクなどに助けられていました。これは、言語サーバーと型チェッカーによって強化されていました。いくつかのハーネスは、より優れたコード インテリジェンスによって実現される合理的な前提に基づいて、LSP をツールとして公開しています。
[切り捨てられた]

記事本文:
エディターにコードを入力していた暗黒時代には、壊れたコードの下の波線、クリックして定義へのリンクなどに助けられていました。これは、言語サーバーと型チェッカーによって強化されていました。いくつかのハーネスは、より優れたコード インテリジェンスがより優れたエージェントを生み出すという合理的な前提に基づいて、LSP をツールとして公開しています。
ただし、モデルは主に、常に使用できるツール (grep や ranged-read など) を使用してトレーニングされます。効果を高めながらそれらに取って代わるのは困難です。
モデルは人間とは少し異なる方法でコードベースを理解します。人間は、少数のファイルを視野に入れ、少し多めのファイルを作業メモリに保持し、時間の経過とともに、コード ベースのおおよそのメンタル モデルを構築します。
エージェントには巨大なコンテキスト ウィンドウがあり、同時に多くのファイルを理解できます。シンボルを検索することでシンボルを見つけることができます。これにより、通常は部分読み取りがトリガーされ、ファイルの一部がコンテキスト ウィンドウに取り込まれます。彼らはまた…何かを知っているだけですか？それらの重みには、非常に多くの公開コードのかなり忠実度の高いバージョンが含まれています。これは、そのコードをナビゲートしたり、他のコード ベースについて類推したりするのに役立ちます。
このプロセスで LSP がどのように役立つかを試すために、このリポジトリにある一連の実験を実行しました。 1 実験はローカル モデルと API モデルを組み合わせて、Python に対して動作しました。チェッカーには Pyrefly を使用し、静的 AST 解決は Pyrefly 2 に対して検証され、ほとんどのテストを実行しました。
私が抱いた疑問は、LSP が他の方法では見つけられないことをエージェントに伝えるのか、LSP の回答はデフォルトの読み取りよりもトークンのコストが安いのか、診断のタイミングによって結果が変わるのかということでした。
最初の質問に対する答えは、ほとんどがノーでした。 var 内の 10 個の同じ名前のオーバーライド間の解決

モデルに LSP ツールがあるかどうかに関係なく、セットアップの一部は機能しました。ただし、型アノテーションによって違いが生じました。これらを削除すると、テキストの検索が悪化しました。ハーネスの有能なモデルは、基本的にはすでにまともなタイプ チェッカーです。
効率に関しては、LSP を追加するだけでは何も起こりません。モデルは何らかのプロンプトがなければ LSP を使用しませんでした。モデルが型を解決するためにファイルを読み取る必要がある場合、LSP defn ツールはファイル読み取りツールよりも安価でした。 3 でも！多くの場合、モデルはファイルの読み取りも行うため、トークン効率に関する LSP の利点が完全に無効になってしまいます。
関連する修正を含む定義を挿入した場合でも、モデルはほぼすべての場合にファイルを再読み取りします。提供されたスパンが完了したことをモデルに伝えることは、それらのケースの数 % を節約するだけであり、プロンプト自体を提供するためにより多くのトークンがかかります。実際にモデルがファイル読み取りへの defn 呼び出しを優先するようにするには、微調整が必​​要です。 Qwen 3.6 で DAgger スタイルのラベル変更 (軌跡を生成し、読み取ったファイルを読み取った定義と交換し、結果を微調整する) を使用すると、モデルは無関係な呼び出しを回避します。 4
安ければ安いほど良いというわけではないことにも注意してください。これらは入力/事前入力トークンであり、最適化するのに低コストであまり興味がありません。 Codebase-Memory の論文では、ファイル探索エージェントが、実際のリポジトリ上で構造化グラフ ツールをトークンの約 10 倍で上回った (92% 対 83%) ことがわかりました。したがって、それらの追加のトークンが何か役立つことがある場合があります。
3番目の質問はタイミングについてでした。型注釈を正しく保ちたい場合は (実際にそうしています)、型チェッカーを使用できます。しかし、フィードバックをいつ提供するのが最適でしょうか?
エージェントに、目に見えるテストに合格した一連の変更草案に取り組んでもらいましたが、保留されたテストには不合格となり、レビューとサブスクライブを依頼しました。

ミット。放っておくと、モデルは 12 回中 11 回不正なリビジョンを受け入れました。型チェッカーが送信をゲートしたとき、12 回中 1 回しか不正な変更を受け入れませんでした。5
インタラクション モデルに関する Thinky の投稿の後、生成中に診断をライブで配信する (基本的にエージェントの波線) ことが役立つかどうかを確認したいと思いました。 6 ほとんどの場合、答えはノーです。ライブ配信は中立的であり、フィードバックはありませんでした。モデルに修正を促すことは明らかに有害であるように見えます。エージェントにツールのフィードバックに取り組むように指示することは、悪い意味で自身の判断を無効にしているようです。
ターン終了時または各編集後にフィードバックをバッチ処理することは役に立ち、全体的なトークン支出に関してはターン終了時が明らかに勝者でした。
注意点には注意点があります。タスクの多くは合成的なもので、すべてのタスクは非常に簡単で、まともなモデルはほぼすべてを解決します。タイミング結果は 7B モデルでのみテストされたため、強力なモデルを使用するとより良い結果が得られる可能性があります。問題のコードベースはコンテキスト ウィンドウ内に完全に収まりますが、実際の使用では、エージェントが積極的に全体を取り込むことはないようです。
そうは言っても、少なくとも私自身の仕事に関して、いくつかの教訓があります。
種類は良いですよ。正しいアノテーションは、ツールに関係なく、エージェントがコードベースをナビゲートするのに役立ちました。
ターン終了フックを使用してタイプを正しく保ちます。理想的には、これをゲートにして、型エラーが発生した場合は積極的に修復を要求し、それ以外の場合は沈黙するようにします。それでも、測定してください！ブロックする頻度、修復が成功する頻度、すべきではない作業を拒否する頻度を記録します。
タスクレベルでトークンの節約を測定します。単一の演算をベンチマークする場合、その演算の方が安価であると結論付けるのは簡単です。ただし、変更を実際に評価するには、完全なモデルの動作を確認する必要があります。ツールやプロンプトの有無にかかわらず正確性は同じですか、また、ツールやプロンプトを使用しても一貫したトークンの節約が得られるかどうかなどです。

使用範囲。
大規模なコードベースで LSP とプロンプトをさらに試してください。これを直接テストしたわけではありませんが、結果から直感的に感じたのは、大規模で複雑なプライベート コードベースのナビゲーション ツールをより活用できるということです。おそらく、モデルにツールを積極的に使用するように促す必要があります。デフォルトでは、モデルは知っていることに手を伸ばすことになります。
将来的には興味深い取り組みがいくつかあります。私たちが使用しているツールは、裁量を適用したり出力を無視したりできる人間向けに構築されています。これはモデルにとってはさらに困難であり、おそらくトレーニングが必要なものです。
また、コードベースに対して大量のタスクを実行することなく、ツールが役立つかどうかをコードベースで評価できるようにしたいと考えています。そのような種類の決定に役立つ可能性のある、静的に収集できるメトリクスはあるでしょうか?
最後に、より大規模なコード ベースのリファクタリングと移行タスクが必要だと思います。これは、ProgramBench 規模の問題ですが、トレーニング セットに含まれていない大規模な既存のコードベースを扱う問題です。その点で、これを書いているときに SWE-Bench ProMax が出てきました。関連性があるようですが、まだ読んでいません。
このリポジトリの大部分は Codex と Claude の功績です。レポートの記述は LLM に加えて大量の編集が加えられているため、期待を和らげてください。興味があれば、すべての数値はそこにあるアーティファクトから再現できます。 ↩︎
実際のライブラリ シンボルに関しては、両者の意見はほとんどの場合一致しました。ギャップは、Pyrefly が null を返し、リゾルバが null を返さなかった再エクスポートでした。 ↩︎
範囲読み取りのトークンコストは、defn コールよりも約 1.3 倍高くなります。 ↩︎
この種のトレーニングの 1 つのリスクは、モデルに表面的なツールの使用法を教えるだけで、判断力を教えることはできないことです。場合によっては、実際にファイルを読み取る必要がある場合があります。少し驚いたことに、それは問題ではないようでした。不十分なスパンを与えたとき、それは読み出しました。

毎回、そしてスパンが十分な場合は 2 回しか読み取られませんでした。トレーニング データの読み取りを必要とするタスクの実際の例はなかったため、条件付けによってモデルの判断が完全に無効になるわけではないことがわかります。 ↩︎
他のケースでは、バグは型チェックで明らかになりましたが、このケースでは型チェックが正常に行われたため、信号はありませんでした。 ↩︎
ライブ ストリームに結果を挿入するために私が使用した実際のアプローチは、Hooper らによって説明されているイベントの非同期注入アプローチです。 。 ↩︎
今すぐ購読して読み続け、完全なアーカイブにアクセスしてください。

## Original Extract

Back in the dark ages of typing code into editors we were aided by squigglies under broken code, click-to-definitions links, and so on. That was powered by language servers and type checkers. Several harnesses now expose an LSP as a tool, on the reasonable premise that better code intelligence makes
[truncated]

Back in the dark ages of typing code into editors we were aided by squigglies under broken code, click-to-definitions links, and so on. That was powered by language servers and type checkers. Several harnesses now expose an LSP as a tool, on the reasonable premise that better code intelligence makes for a better agent.
Models, though, are trained primarily with the tools they always have, which tend to be things like grep and ranged-read. Supplanting those while gaining effectiveness is tricky.
Models understand codebases a bit differently than people. A human can keep a handful of files in view, a slightly larger handful in their working memory and, over time, they build up an approximate mental model of the code base.
An agent has an enormous context window and can understand a lot of files at the same time. They can find symbols by searching for them, which will usually then trigger a partial read to pull part of the file into their context window. They also just… know stuff? Their weights contain reasonably high-fidelity versions of an awful lot of public code. That helps with navigating that code, or reasoning by analogy about other code bases.
To try and see how LSPs might help in this process, I ran a bunch of experiments, which are in this repo . 1 The experiments were on a mix of local and API models, working against Python. I used Pyrefly for the checker, and a static AST resolve validated against Pyrefly 2 to run most of the tests.
The questions I had were whether the LSP tells the agent anything it couldn’t otherwise find, whether LSP answers are cheaper in tokens than the default reads, and whether the timing of a diagnostic changes the outcome.
The answer to the first one was, mostly, no. Resolving between 10 same-named overrides in a variety of setups worked whether or not the model had LSP tools. The type annotations did make a difference though. If I stripped those out, the text retrieval got worse. A capable model in a harness is basically already a decent type checker.
With regards to efficiency, merely adding an LSP did nothing: the models didn’t use it without some prompting. If the model had to read a file to resolve a type, the LSP defn tool was cheaper than the file-read tool. 3 But! The models would often do the file read as well , which completely negated the benefit of the LSP for token efficiency.
Even if I injected definitions that were contained the relevant fix, the model re-read the file in almost every case. Telling the model that the span provided was complete only saved a few % of those cases, and it cost more tokens in providing the prompt itself! To actually get the model to prefer the defn call to a file read required fine-tuning. I used a DAgger-style relabel (generate trajectories, swap the file read for the definition read, fine-tune on the result) on Qwen 3.6 and then the model would avoid the extraneous call. 4
Its also worth noting that cheaper isn’t always better. These are input/prefill tokens, which are the cheaper and less interesting ones to optimize. The Codebase-Memory paper found a file-exploration agent beat a structured graph tool on real repositories (92% against 83%) at roughly ten times the tokens. So, sometimes those extra tokens are doing something useful.
The third question was about timing. If you want to keep your type annotations correct (you do) you can use a type checker. But when is best to deliver the feedback from it?
I had an agent work on a set of draft changes that passed visible tests, but failed a held out one, and asked it to review and submit. Left alone, the model accepted the bad revision 11 times out of 12. When a type checker gated the submission, it only accepted the bad change 1 time out of 12. 5
After Thinky’s post on interaction models , I wanted to see whether delivering diagnostics live during generation (agentic squiggles, basically) would help. 6 Largely, the answer is no: live delivery was neutral vs no feedback. Prompting the model to make the fix seemed actively harmful: telling an agent to work on tool feedback seems to override its own judgement in a bad way.
Batching the feedback at end of turn or after each edit did help, and for overall token spend end-of-turn was a clear winner.
Caveats caveats caveats: many of the tasks were synthetic, all the tasks were pretty easy, and every decent model solved pretty much everything. The timing results were only tested with a 7B model, so you may get better results with a strong one. The codebases in question can fit fully within context window, though in actual usage agents never seemed to proactively just ingest the whole thing!
With that said, I do have some takeaways, at least for my own work:
Types are good. Correct annotations helped the agent navigate the codebase, regardless of tooling.
Keep types correct with an end-of-turn hook. Ideally, make this a gate so it actively asks for repair on type errors and is silent otherwise. Still, measure it! Log how often it blocks, how often the repair passes, how often it rejects work it shouldn’t have.
Measure token savings at the task level. If benchmarking a single op it is easy to conclude the op is cheaper. You need to see the full model behavior to really assess the change though: is correctness the same with and without a tool or prompt, and does it yield consistent token savings across a range of usage.
Experiment more with LSPs and prompts on larger codebases. I didn’t test this directly, but my instinct from the results is that you will get more out of navigation tools for large, complex, private, codebases. You will still likely have to prompt the model to actively make use of the tools: by default, they’re going to reach for what they know.
There is some interesting future work out there. The tools we’re using were built for humans who can apply discretion/ignore output. That is trickier for a model, and something that probably needs training.
I’d also like to be able to evaluate a codebase as to whether a tool will help without having to just run a bunch of tasks over it: are there metrics we can collect statically that might inform those kind of decisions?
Finally, I think we need more large code base refactor and migration tasks: problems of a ProgramBench scale but working with a large, existing, and not-in-the-training-set codebase. On that note, SWE-Bench ProMax came out as I was writing this: seems relevant, but I haven’t yet read it!
Credit to Codex and Claude for most of this repo, the writing in the report is LLM+a bunch of editing, so temper your expectations, and all the numbers are reproducible from the artifacts there if interested. ↩︎
On real library symbols the two agreed most of the time: the gaps were re-exports where Pyrefly returned null and the resolver didn’t. ↩︎
The ranged-read cost about 1.3x more expensive in tokens than the defn call. ↩︎
One risk with this kind of training is that it teaches the model superficial tool usage, but not judgement: sometimes, you do need to actually read the file! Somewhat surprisingly that didn’t seem to be a problem: when I gave it insufficient spans it went and read every time, and when the span was sufficient it read only twice. There weren’t any actual examples of tasks requiring reads in the training data, so it suggests that the conditioning doesn’t completely kill model judgement. ↩︎
In the other cases the bugs were exposed in a type-check, but this one type-checked clean, so no signal. ↩︎
The actual approach I used for injecting results into a live stream is the async injection of events approach described by Hooper et al. . ↩︎
Subscribe now to keep reading and get access to the full archive.
