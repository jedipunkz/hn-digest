---
source: "https://sevangelatos.github.io/spiruxbench-results/blog/"
hn_url: "https://news.ycombinator.com/item?id=50570496"
title: "I benchmarked 22 local LLMs looking for the Qwen 3.8 35B we didn't get"
article_title: "I benchmarked 22 local LLMs looking for the Qwen 3.8 35B we didn’t get."
image: "https://sevangelatos.github.io/spiruxbench-results/blog/img/full_table.png"
author: "koukopoulos"
captured_at: "2026-11-20T05:57:51Z"
capture_tool: "hn-digest"
hn_id: 50570496
score: 2
comments: 0
posted_at: "2026-11-20T05:24:48Z"
tags:
  - hacker-news
  - translated
---

# I benchmarked 22 local LLMs looking for the Qwen 3.8 35B we didn't get

- HN: [50570496](https://news.ycombinator.com/item?id=50570496)
- Source: [sevangelatos.github.io](https://sevangelatos.github.io/spiruxbench-results/blog/)
- Score: 2
- Comments: 0
- Posted: 2026-11-20T05:24:48Z

## Translation

タイトル: 入手できなかった Qwen 3.8 35B を探して 22 個のローカル LLM をベンチマークしました
記事のタイトル: 入手できなかった Qwen 3.8 35B を探して 22 個のローカル LLM をベンチマークしました。
説明: ゲーム PC でのエージェント コーディング用の 22 個のローカル LLM の実践的なベンチマーク: Qwen の微調整、MoE と高密度、量子化効果、および私の選択。

記事本文:
Qwen 3.6 35B の後継機は見つかりましたか? Qwen 3.6 の微調整
高密度モデルと Qwen 3.8 Flash Next
Qwen 3.8 27b (16 ギガ VRAM)
私が選ぶエージェント コーディングまたは中期/長期タスク
22 個のローカル LLM のベンチマークを実行し、入手できなかった Qwen 3.8 35B を探しました。
Spiros Evangelatos 著 · 2026 年 11 月 19 日 · 10 分で読む
Qwen 3.6 35B の後継機は見つかりましたか? Qwen 3.6 の微調整
高密度モデルと Qwen 3.8 Flash Next
Qwen 3.8 27b (16 ギガ VRAM)
私が選ぶエージェント コーディングまたは中期/長期タスク
ここ 1 年ほど、私はローカル AI に魅了されてきました。自分のコンピュータが単独で問題に対処する様子を観察したり、コンピュータを対話型ヘルパーとして使用したりするのは、魔法のように非現実的なものです。私の意見では、クラウド AI にはそのような魔法はありません。まるで外部の存在にハンドルを渡しているような気分になります。いずれにせよ、私はゲーミング PC 上で現実的に動作するモデルの機能を測定する方法を探していました。より正確に言えば、実行する特定の量子化をテストすることです。また、Qwen 3.6 35B の微調整のどれかが、入手できなかった 3.8 MoE の代替品として使用できるかどうかも知りたかったのです。そこから事態はエスカレートしていきました…
そうです、私は独自の LLM ベンチマークを構築しました。
これは主にベンチマーク モデル用のバイブ コード化されたハーネスです。これは、podman コンテナーと pi を中心に構築されています。エージェントは pi で動作し、100k コンテキスト (最大出力 32k) で、時間制約の下で動作します。これには、llama.cpp サーバーにアクセスするための名前付きパイプ経由でのみホストと通信する、ネットワークから分離されたコンテナーが与えられます。課題は、非表示の自動テストに基づいて別のコンテナで採点されます。課題のほとんどは、3 つの主要な柱を中心とした Python および C++ コードの課題で構成されています。第 1 の柱は、leet コード スタイルの解決策を求めるアルゴリズムの課題です。

問題または既知のアルゴリズムの実装について。デバッグの柱は、何らかのバグを含むコード スニペットで構成されます。課題は、バグを見つけて修正することです。その後、修正は非表示の単体テストによってテストされます。最後の柱は、いくつかのミニプロジェクトの実施です。これらの範囲は、ミニ SQL 言語や BigInt 実装から、リバース エンジニアリング タスクや提供されるユーティリティの機能の複製まで多岐にわたります。フル実行には約 10 時間かかります。
私の最初の試みは、取り組むべき 100 個の異なる Python および C++ の課題を思いつき、MiMO v2.6 Flash を割り当てることから始まりました。課題は…挑戦的であるように見えました。その主張と私の最初の楽観主義は、テストに耐えられませんでした。おそらくすべてのモデルが同じ基本的な課題に基づいてトレーニングされていると思います。非常に弱いモデルであっても、それらを乗り越えることができます。非常に小さなモデルに対しても新たな敬意を抱くようになりましたが、ほとんどのモデルが合格するというベンチマークは役に立ちません。
私は Gemma 4 E4B (明らかに弱いエージェント プログラミングのモデル) を使用してすべての課題を試し、合格したものはすべて破棄しました。その副作用として、私のベンチマークは Gemma 4 ファミリに対して少し偏っている可能性があります。その後、ミニ プロジェクトのアイデアのブレーンストーミングと、GPT-Sol の助けを借りて、課題に変える最近の CVE を適応させることに基づいて、徐々に課題を再度構築していきました。私は、プロンプトや公開テストでモデルが残すであろう露骨な特典について、タスクを積極的にレビューする必要がありました。最後に、すべてのタスクが平等に行われるわけではありません。そこで、タスクに重みを付けて、合格率が低く、スコアが高いタスクを与えました。もちろんこれはヒューリスティックですが、すべてのタスクを均等に数えるよりもはるかに優れています。
全体的には結果に満足しています。これは、私自身の経験とよく一致していると思います。

コーディングタスク用のモデル。
Qwen 3.6 35B の後継機は見つかりましたか? ＃
多分。 Qwen 3.6 35B の後継機は、オリジナル モデルを微調整したものであるか、精神的な後継機である可能性があります。精神的な後継者は、実行が比較的簡単で、コーディングを対象としていますが、必ずしも Qwen アーキテクチャに従うわけではない MoE モデルになります。
まとめると、Ornith 1.5 35Bは本物だということです。これは、私が見つけたサードパーティ製の最高の微調整の 1 つで、オリジナルのものを明らかに上回り、このデータセットで有意義に高い品質を実現しています。
オリジナルのスコアを上回り、より複雑なプロジェクトで優れたパフォーマンスを発揮するもう 1 つのモデルは、 Qwen-AgentWorld-35B です。元のモデルよりも全体的な合格率は低くなりますが、より複雑な完全なプロジェクト シナリオでははるかに優れたパフォーマンスを発揮します。これはコーディングに重点を置いたモデルではないため、多少一貫性がありませんが、予想外の利点が得られます。また、元のモデルよりもトークンの消費量が大幅に少なくなります (タスクあたりの平均トークンが 43.3k 対 69.5k)。全体的には Ornith に勝るものはなく、プロジェクト レベルのタスクでは、おそらくより強力ではあるが、より遅いモデルを実行する可能性が高いと思います。
他のすべての微調整 ( Occamy 、 Apex-35B-V1.0 、 XYZ-Aquila-mini 、 MiniCPM-V-4.7 )、さらには Nex-N2.5-mini もベンチマークを開始しましたが、最初のチャレンジで十分な結果が得られず、実行を中止しましたが、元のモデルと比較してパフォーマンスが低下しました。
私の探究の中で、私はいくつかの代替専門家混合モデルをテストしました。一般的な競合製品は Gemma 4 26B です。 Gemma は、エージェント コーディングにおいて Qwen に対してあまり競争力がないことが知られています。 Gemma 4 E4B を使用してデータセットをフィルタリングすると、Gemma 4 ファミリの簡単なタスクが削除されているため、若干不利になると主張する人もいます。しかし、それはスコアにはあまり現れません。ジェマにあげたかった

彼はQ6クォントを実行することでそれを最大限に達成できました。これは他のタスクにとっては優れたモデルですが、エージェント コーディングではうまく機能しませんでした。 Qwenと比較すると大幅にパフォーマンスが劣ります。
フリントも、発売時に大きな話題を集めたモデルです。確かに Qwen 3.6 を上回り、Ornith とほぼ同等です。しかし、私の意見では、これには致命的な欠陥があり、迷惑だと思います。冗長であり、単純なタスクでは一貫性がなくなる傾向があります。より困難なプロジェクト レベルのタスクが加わると、スコアが向上します。これらのタスクでは、おそらく代わりに密なモデルを使用するでしょう。
Nemotron 3.5 Bolt 74B-A9B-24GB はスペクトルの対極にあり、ハードウェアで正常に動作する場合は、Qwen 3.6 の優れた代替品であると考えられます。ここで NVidia が引き出した量子化の魔法がうまく機能しているようです。このモデルは、単純なタスクでは比較的一貫性があり、よくまとまっています。私にとっては明らかにクウェンを上回っています。タスクが複雑になると、これでは耐えられなくなりますが、これらのタスクでは、代わりに低速の密度モデルを使用します。私の経験では、Nemotron モデルは十分にトレーニングされておらず、パフォーマンスも低い傾向にあるため、これは私にとって嬉しい驚きでした。
一方、Nemotron 3.5 Lightning 30B-A3B は、エージェント コーディングでは完全に失敗し、時代遅れの gpt-oss-20b に負けました。ただし、GPT は非常に高速でトークン効率が高いため、依然としてその役割を果たしています。 Nemotron は明らかにコーディング モデルではありませんが、それでもエージェント コーディングのためのトリックが隠されていると私は信じています。これは、ネイティブ 1M トークン コンテキスト ウィンドウを備え、比較的コンテキスト効率が高い唯一のモデルです。したがって、圧縮に役立ちます。独自のモデルで実行されている長時間実行セッションを取得し、ほぼゼロのコストで (またはオフライン中に) 圧縮し、ローカル モデルで作業を続けることができます。
高密度モジュール

els と Qwen 3.8 フラッシュの次の #
私の意見では、小規模な MoE モデルは、関心のあるコードベースへの変更を信頼できるほど信頼性がまだ高くありません。高密度モデルと QFN (Qwen Flash Next) は、変更に注意を払い最新情報を把握している限り、確実に実現します。すべての高密度モデルがトップ 10 スコアに入ります。小型の MoE モデルのうち、トップ 10 に入るのは Nemotron 3.5 Bolt 74B-A9B-24GB 、Flint および Ornith 1.5 35B だけです。
予想通り、Qwen3.8-27B-UD-Q4_K_XL (UD V3.0) と Qwen3.8-Flash-Next-UD-IQ4_XS (UD V3.1) がスコアを支配しており、フラッシュ モデルの方がわずかに高いスコアを獲得していますが、通過したチャレンジの数で負けています。 1 つ下に行くと、QFN UD-IQ3_XXX (UD V3.0) 量子化が 3 位に下がります。
ここで驚くべきことは、エキスパートの半分が削除され、32GB RAM 構成で実行可能な QFN の縮小バージョンがそれほど遅れていないことです。 Qwen3.8-Flash-Next-Coder-GSQ-RCO-IQ1_M は、非常に立派な 4 位を獲得しました。
トークン効率は、gemma-4-31B (14.118 ポイント/1,000 トークン) および Swift-1.5-Qwen3.8-27B (10.086 ポイント/1,000 トークン) でさらに優れています。 Swift は、使用する価値のある優れた微調整機能のように見えます。
Muse-Glimmer-30B はトークン効率も良く (9.897 ポイント/1,000 トークン)、全体的に良いスコアを獲得しています。私の経験では、そのモデルは少し怠惰になる傾向があります。ただし、特定の変更を処理するためにクルーズ コントロールの一種として使用すると、うまく機能します。
Qwen3.8-27B で中程度の労力を使用すると、高いトークン効率 (6.035) が達成されないようです。おそらく、ベンチマークが行われていない場合、モデルの反応は異なるでしょう。このスコアは、Gemma、Muse、Swift などのより効率的な高密度モデルのスコアよりも低くなります。この失敗率の増加は、より複雑な課題に主に影響を与えているようです。いずれにせよ、Swiftを使用すると、

Qwen をミディアムで実行する場合と比較して、トークンの効率と機能が向上します。
使用された累積トークンを見ると、より困難なタスクに直面したときのモデルの反応が異なることは明らかです。一部のモデルはトークンの使用量を大幅に拡大しますが、他のモデルは横ばいのままです。ジェマはかなり一貫して倹約しているようです。 Muse glimmer はより多くのトークンを消費しますが、タスクが難しい場合にはトークンの使用量を大幅に拡大しません。一方で、クウェン家は難しいタスクにかなり多くのトークンを費やします。他のモデル (Flint、Nemotron Bolt、Ornith など) も同様の動作をします。
Qwen 3.8 27b (16 ギガの VRAM) #
興味深い質問は、VRAM が 16 GB しかない場合、MoE を実行する場合と比較して、カードに Qwen3.8-27B を詰め込む価値があるかどうかです。そのために、VRAM 内のコンテキストに適合させるために、Q8_0 KV キャッシュ量子化を使用して Qwen3.8-27B-GGUF:UD-Q2_K_XL の 2 つの異なる量子を試しました。 (他のすべてのテストは 16 ビット KV キャッシュを使用して実行されました)。これらの唯一の違いは、一方は古い unsloth UD-V3.0 量子化メソッドを使用し、もう一方は更新された V3.1 メソッドを使用していることです。
答えは否定的なようです。どちらの量子化スコアも複数の MoE モデルを下回っています。オーニスとフリントはどちらもスコアが高く、ネモトロン ボルトは高難易度のチャレンジで負けたためスコアが少し低くなりましたが、より簡単なチャレンジではより安定しています。
ただし、新しい量子化方法の方がスコアが高かったことは注目に値します。さらに重要な違いは、量子化方法間で消費されるトークンにありました。
新しい量子化方法により消費されるトークンは約 25% 削減されます。したがって、ナマケモノが何をしたとしても、それは明らかに変化をもたらしました。他の量子化方法との違いはさらに顕著になるのではないかと想像しています。
データは、次のことを示唆しています。

より多くの探索を必要とする特定のデバッグ タスクでは、低品質の量子化に対するトークンの使用量の増加が実際に報われる可能性があります。単なる測定ノイズかもしれませんが。
エージェントティック コーディングまたは中長期タスク #
VRAM が 24GB 以上ある場合は、 Qwen 3.8 27B を選択します。 CPU と RAM をユーザー自身またはハーネスに使用できる状態にしておくと、優れたコーディング能力が得られます。同様に確実な選択肢であり、実際、VRAM が 24GB 未満の場合に唯一良い選択肢は、Qwen 3.8 Flash Next です。 RAM に適した量子化レベルを選択してください。コーダーのバリアントは依然として優れています。これらのいずれかに十分な VRAM または RAM がない場合は、 Ornith-1.5-35B を実行します。
より簡単なタスクや、対話性が必要でエージェントを細かく管理できるタスクの場合は、 Swift-1.5-Qwen3.8-27B 、 gemma-4-31B 、 Muse-Glimmer-30B 、または Nemotron-3.5-Bolt-74B をお勧めします。十分な RAM がある場合、または別の推論システムがあり、そのシステム全体を LLM 専用にしても構わない場合は、Qwen3.8-Flash-Next-GGUF がさらに優れています。 Swift-1.5 Flash Next 私はテストしていませんが、この役割にもうまく適合する可能性があると思います。これらは比較的信頼できる実装者、コードレビュー担当者、または協力者である必要があり、それぞれに独自の強みがあります。
アイデアを話し合う、対話する

[切り捨てられた]

## Original Extract

A hands-on benchmark of 22 local LLMs for agentic coding on a gaming PC: Qwen finetunes, MoE vs dense, quantization effects and my picks.

Did you find a Qwen 3.6 35B successor? Qwen 3.6 finetunes
Dense models and Qwen 3.8 Flash Next
Qwen 3.8 27b on 16 gigs of VRAM
My picks Agentic Coding or Medium/Long Horizon tasks
I benchmarked 22 local LLMs looking for the Qwen 3.8 35B we didn’t get.
By Spiros Evangelatos · November 19, 2026 · 10 min read
Did you find a Qwen 3.6 35B successor? Qwen 3.6 finetunes
Dense models and Qwen 3.8 Flash Next
Qwen 3.8 27b on 16 gigs of VRAM
My picks Agentic Coding or Medium/Long Horizon tasks
The last year or so I have been fascinated by local AI. There is something magical and unreal in watching your own computer work on a problem on its own or using it as an interactive helper. Cloud AI, in my opinion, doesn’t have that same magic. It feels like you are handing the wheel to some external entity. In any case, I was looking for a way to measure the capabilities of models that are realistic to run on a gaming PC. Or to be more exact, testing the specific quantizations that one would run. I wanted also to understand which, if any of the fine tunes of Qwen 3.6 35B could be used as a replacement of the 3.8 MoE that we never got. Things escalated from there…
That’s right, I built my own LLM benchmark .
It is a mostly vibe coded harness to benchmark models. It is built around podman containers and pi. The agent works in pi, with 100k context (32k max output) and under a time constraint. It is given a network isolated container that only communicates with the host over a named pipe for accessing the llama.cpp server. The challenges are graded in another container based on hidden automated tests. The challenges consist mostly of Python and C++ code challenges around 3 major pillars. Pillar one is algorithmic challenges that ask for a solution to a leet-code style problem or for an implementation of a known algorithm. The debugging pillar consists of code snippets given that contain some bug. The challenge is to find and fix the bug. The fix is then tested by hidden unit tests. The final pillar is the implementation of some mini projects. These range from a mini SQL language or BigInt implementation to reverse engineering tasks or the duplication of the functionality of a utility that is provided. A full run takes about 10 hours.
My initial attempt started by just assigning MiMO v2.6 Flash with coming up with 100 different python and C++ challenges to work on. The challenges appeared …challenging. That assertion and my initial optimism did not survive testing. I guess all of the models are trained on the same basic challenges. Even very weak models could ace through them. I got a new degree of respect for even the very small models but a benchmark that most models ace is useless.
I put Gemma 4 E4B (an admittedly weak model for agentic programming) to attempt all of the challenges, then simply discarded any that it passed. A side effect of that is that my benchmark might be a bit biased against the Gemma 4 family. After that, I gradually built up again the challenges based on brainstorming mini project ideas and on adapting, with the help of GPT-Sol, recent CVEs that we would turn into challenges. I had to aggressively review the tasks for blatant giveaways that the models would leave behind in the prompts or in the public tests. Finally, not all tasks are made equal. So I weighted the tasks giving tasks with a lower pass rate, a higher score. This is a heuristic of course but it’s much better than counting all tasks equally.
Overall I am happy with the result. I think it correlates well with my own experience working with these models for coding tasks.
Did you find a Qwen 3.6 35B successor? #
Maybe. A Qwen 3.6 35B successor could be either a fine-tune of the original model or a spiritual successor. A spiritual successor would be a MoE model that is relatively easy to run, targets coding but does not necessarily follow the Qwen architecture.
The summary of this is that Ornith 1.5 35B is the real deal. It’s one of the best third party fine tunes I found that clearly surpasses the original one and meaningfully delivers higher quality in this dataset.
Another model that surpasses the original in score and does better in more complex projects is Qwen-AgentWorld-35B . It has a lower overall pass rate than the original model, but it does much better in more complex full project scenarios. It is not supposed to be a coding focused model and it is somewhat inconsistent but it delivers some unexpected gains. It is also significantly less token hungry than the original model (43.3k vs 69.5k average tokens per task). It does not beat Ornith overall and I feel that for project level tasks one would most probably run a stronger, probably slower model.
All other fine-tunes ( Occamy , Apex-35B-V1.0 , XYZ-Aquila-mini , MiniCPM-V-4.7 ) and even Nex-N2.5-mini , which I started benchmarking but was doing bad enough in the first challenges that I aborted the run, underperform compared to the original model.
In my quest I tested several alternative Mixture of Experts models. A common competitor is Gemma 4 26B . Gemma is known not to be very competitive in agentic coding against Qwen. One could also argue that filtering the dataset using Gemma 4 E4B puts it at a slight disadvantage as easy tasks for the Gemma 4 family have been removed. That does not really show up in the scores though. I wanted to give Gemma the best shot it could by running a Q6 quant. It is an excellent model for other tasks but in agentic coding it did not do well. It significantly underperforms compared to Qwen.
Flint is another model that got lots of hype when it was released. It outperformed Qwen 3.6 indeed and is roughly on par with Ornith. But it has a fatal flaw in my opinion that I find annoying. It is verbose and has a tendency to be inconsistent in simpler tasks. It is picking up score when more difficult project level tasks get in the mix. For these tasks I would probably use a dense model instead anyway.
Nemotron 3.5 Bolt 74B-A9B-24GB is at the other end of the spectrum and I would consider it an excellent replacement for Qwen 3.6 if it runs well on your hardware. The quantization magic that NVidia pulled here seems to work well. The model is relatively consistent in the simpler tasks and is well rounded. It clearly beats Qwen for me. As the tasks get more complex, it does not hold up but for these tasks I would use a slower dense model instead. That was a pleasant surprise for me as in my experience Nemotron models tend to be undertrained and usually underperform.
On the other hand Nemotron 3.5 Lightning 30B-A3B fell completely flat in agentic coding, getting beaten by the outdated gpt-oss-20b . GPT still has its place though as it is extremely fast and very token efficient. While Nemotron is clearly not a coding model, I believe it still has a trick up its sleeve for agentic coding. It is the only model here that has a native 1M token context window and is relatively context efficient. So it can be useful for compaction. I can take a long running session that is running in a proprietary model, compact it at near zero cost (or while being offline) and keep working with a local model.
Dense models and Qwen 3.8 Flash Next #
In my opinion, small MoE models are not yet reliable enough to be trusted with changes to a codebase you care about. The dense models and QFN (Qwen Flash Next) are getting there as long as you keep an eye on their changes and you stay in the loop. All of the dense models make it to the top 10 scores. Of the smaller MoE models only Nemotron 3.5 Bolt 74B-A9B-24GB , Flint and Ornith 1.5 35B make it to the top 10.
Predictably, Qwen3.8-27B-UD-Q4_K_XL (UD V3.0) and Qwen3.8-Flash-Next-UD-IQ4_XS (UD V3.1) dominate the scores with the flash model getting a slightly higher score but losing in the number of challenges passed. Going one step lower, the QFN UD-IQ3_XXX (UD V3.0) quantization falls to the 3rd position.
The surprise here is that the cut down version of QFN with half the experts removed that is runnable in a 32GB RAM config is not far behind. Qwen3.8-Flash-Next-Coder-GSQ-RCO-IQ1_M gets a very respectable 4th position.
Token efficiency is even better for gemma-4-31B (14.118 points/1k tokens) as well as Swift-1.5-Qwen3.8-27B (10.086 points/1k tokens). Swift looks like an excellent fine-tune well worth using.
Muse-Glimmer-30B has good token efficiency as well (9.897 points/1k tokens) while getting a good score overall. In my experience that model tends to be a bit lazy. But if you use it as a type of cruise control, to handle specific changes that you want to do, it works well.
Using the medium effort on Qwen3.8-27B does not seem to achieve a high token efficiency (6.035). Maybe the model would react differently when not being benchmarked. The score ends up lower than that of the more efficient dense models like Gemma, Muse and Swift. That increased failure rate seems to mostly affect the more complex challenges. In any case, using Swift seems to be more token efficient and more capable compared to running Qwen on medium.
Looking at the cumulative tokens used, it is evident that models react differently when faced with more difficult tasks. Some models scale up their token usage significantly while others remain flat. Gemma seems to be fairly consistently frugal. Muse glimmer spends more tokens but does not scale up its token usage significantly when the task is difficult. On the other hand, the Qwen family spends significantly more tokens for difficult tasks. Other models (eg. Flint, Nemotron Bolt and Ornith) have a similar behavior.
Qwen 3.8 27b on 16 gigs of VRAM #
An interesting question is if you only have 16 gigabytes of VRAM, if it is worth it to try to squeeze Qwen3.8-27B in your card, compared to running a MoE. For that I tried 2 different quants of Qwen3.8-27B-GGUF:UD-Q2_K_XL using a Q8_0 KV cache quantization in order to fit the context in the VRAM. (All other tests were done with 16Bit KV cache). The only difference between these was that one was using the older unsloth UD-V3.0 quantization method while the other one was using the updated V3.1 method.
The answer seems to be negative. Both quantizations score below multiple MoE models. Both Ornith and Flint scored higher while Nemotron Bolt scored a little lower due to losing in high difficulty challenges but is more consistent in the easier ones.
It is remarkable that the newer quantization method, scored higher though. An even more significant difference was in the tokens spent between quantization methods.
The tokens spent by the newer quantization method are about 25% less! So whatever unsloth did, it clearly made a difference. I imagine that the differences between other quantization methods may be even more dramatic.
The data suggests that the higher token spend of lower-quality quantizations might actually pay off in certain debugging tasks that require more exploration. It might be just measurement noise though.
Agentic Coding or Medium/Long Horizon tasks #
If you have >= 24GB VRAM, I would pick Qwen 3.8 27B . You get excellent coding ability with your CPU and RAM remaining usable for you or the harness. An equally solid choice and in fact the only good choice if you have < 24GB VRAM is Qwen 3.8 Flash Next . Pick the quantization level that will fit your RAM. The coder variant is still good. If you don’t have enough VRAM or RAM for either of those, run Ornith-1.5-35B .
For easier tasks or tasks where you want interactivity and you can micromanage the agent, I would prefer Swift-1.5-Qwen3.8-27B , gemma-4-31B , Muse-Glimmer-30B or Nemotron-3.5-Bolt-74B . Qwen3.8-Flash-Next-GGUF is even better if you have plenty of RAM or you have a separate inference system and you are ok with dedicating that entire system to the LLM. Swift-1.5 Flash Next I have not tested but I expect it might fit well in this role too. These should be relatively reliable implementors, code reviewers or collaborators and each one has their own strengths.
Discussing ideas, interac

[truncated]
