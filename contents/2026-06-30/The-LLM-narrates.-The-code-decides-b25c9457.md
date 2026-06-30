---
source: "https://www.irinobservability.com/blog/llm-narrates-code-decides"
hn_url: "https://news.ycombinator.com/item?id=48733367"
title: "The LLM narrates. The code decides"
article_title: "The LLM Narrates. The Code Decides. | Irin Observability"
author: "jmlarry"
captured_at: "2026-06-30T14:50:27Z"
capture_tool: "hn-digest"
hn_id: 48733367
score: 1
comments: 0
posted_at: "2026-06-30T14:41:30Z"
tags:
  - hacker-news
  - translated
---

# The LLM narrates. The code decides

- HN: [48733367](https://news.ycombinator.com/item?id=48733367)
- Source: [www.irinobservability.com](https://www.irinobservability.com/blog/llm-narrates-code-decides)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T14:41:30Z

## Translation

Title: The LLM narrates. The code decides
記事のタイトル: LLM が語る。 The Code Decides. | Irin Observability
説明: Python は固定列挙型に分類されます。 LLM は 1 つの文を書きます。この分割により、出力が十分にクリーンに保たれて集約され、モデルがクリティカル パスから外されるようにする方法。

記事本文:
The LLM Narrates. The Code Decides. | Irin Observability
Product ▾
Services
The LLM narrates. The code decides.
私が現在目にしている「可観測性のための AI」の仕事のほとんどは、言語モデルに判断を委ねています。それは逆だと思います。アラートを与え、いくつかの指標を与え、何が問題なのか、何をすべきかを尋ね、判断を下させます。言語モデルを使用した経験に基づいて、プロセスを逆にしたほうがより良い結果が得られると判断しました。
短いバージョン: 私のアラート パイプラインでは、許容される分類のセットが決定論的な Python で固定されており、モデルはそこから選択する必要があります。 LLM の唯一の仕事は、構造化された評決を理解しやすい文章に変換することです。何かが悪いことなのか、どの程度悪いことなのか、あるいはどのようなカテゴリーの問題なのかを決定することはありません。コードがすでにロックダウンしている決定空間内で語られます。
私は小規模なマネージド監視サービスを運営しています。 Alertmanager が起動し、Webhook が着地し、歴史的にその Webhook はホスト Web-VM 上の HighMemoryUsage、重大度警告のような行を生成しました。これは正確ではありますが、あまり役に立ちません。これを読む人は、HighMemoryUsage が何を意味するのか、このホストが常に高温になっているのか、気にする必要があるのか​​を知る必要があります。アラート配信プロセスを変更せずに、平易な英語のコンテキストをアラートに添付したいと考えていました。
明らかな行動は、アラート全体を LLM に投げて、説明を求めることでした。この実験の最初の繰り返しで、ある程度正確であることを期待して試してみましたが、完全に信頼できるわけではありませんでしたが、期待を裏切りませんでした。モデルは自信を持って一貫性がありませんでした。同じアラートが 3 回発生すると、3 つの異なる「根本原因」カテゴリが生成されました。あるテスト アラートは「構成またはセットアップの問題」と呼ばれ、次の実行では「構成」と呼ばれました。

/Testing、次はまた別のことです。後で何らかの集計を行うためにその出力を保存している場合 (私はそうです。同じ週に 3 つの異なるクライアントが同じクラスの問題に遭遇したときを知りたいのです)、自由形式モデルの出力は断片化してノイズになります。モデルがランダムに変更されるフィールドでグループ化することは機能しません。
素晴らしい結果は得られないし、見た目よりも難しいだろうということは最初から分かっていました。そこで私はリサーチを始め、デザインを変えることにしました。私の頭の片隅にあった小さな声はずっと正しかった、モデルに判断を任せてはいけない。
真ん中に硬い壁があるパイプラインを作成しました。
決定論的な側面では、Python が分類を行います。出力は、8 つの値の列挙型 (memory_pressure 、 cpu_saturation 、disk_pressure 、service_unavailability 、network_issue 、configuration_error 、external_dependency 、unknown ) に制限されます。このフィールドは 8 つの文字列のうちの 1 つしか取り得ないため、このフィールドで集計します。何も当てはまらない場合、答えは不明ですが、それ自体が幻覚や変化した推測ではなく、有用な信号です。
説明側では、LLM ( llama3:8b 、私自身の LAN 上のボックスでローカルに実行、データ/ネットワークが安全) は、その固定 8 値セットから分類を選択する必要があり、その横に 2 つの短いフィールド (アラートが何であるか (平易な英語で) と、それが運用上何を意味するか) を書き込みます。コードは答えの形を定義します。モデルは既存のスロットにのみ埋められます。修正を提案したり、原因をでっち上げたりしないように明示的に指示されているため、分析ではなく翻訳が実行されます。
プロンプトは文法制限された厳密な JSON を返すため、毎回 {what, means, possible_cause_class} が返され、途中で enum 値が許可されたセットに対して検証されます。 If the model re

何かをリストから外す場合は、行を保存する代わりにバグをキャプチャします。
ナレーションのステップさえ単純にすると、人を騒がせるような散文になってしまいます。システムをチューニングしているときに DiskFillPredicted アラートを受け取りましたが、一見すると調査する価値があるように見えました。次に、問題のホストを調べたところ、数か月間、ディスク使用率のベースラインが横ばいでした。この予測は丸めのアーティファクトであり、「ディスクがもうすぐいっぱいになる」という予測は誤解を招くものです。モデルはアラートだけではそれを知る方法がなかったので、ただ何かを書き込んだだけでした。
人間が反応する前に見るのと同じコンテキストをモデルに与えることで、この問題を修正しました。 LLM 呼び出しの前に、Python はそのホストの最近のベースラインのメトリクス バックエンドに対して高速検索を実行します。プロンプトには明示的なルールが含まれています。つまり、履歴コンテキストが提供されている場合は、アラートのリテラル テキストよりもそれを重視します。数か月にわたって安定したベースラインを持つホスト上のディスク使用量の予測は情報提供であり、緊急ではありません。
このルックアップの待ち時間バジェットは 5 秒です。 LLM 呼び出し自体には約 80 秒かかります。これは、CPU のみの控えめなハードウェア (16 GB RAM を搭載し GPU を搭載していない第 4 世代 Intel i7) で意図的に実行されるためです。これは選択であり、私が謝る制約ではありません。サービス全体の姿勢として、LAN からは何も出ないので、遅いローカル モデルは高速なリモート モデルに勝ります。そしてその80秒は、警報を受けている人には決して届きません。アノテーターは既存のパスに沿って移動するため (詳細は後述)、生のアラートは即座に Slack と電子メールに送信されます。ナレーション付きバージョンは、1 分ほど後に別の注釈として表示されます。モデルを待っているものは何もありません。その 80 秒の呼び出しに対して、5 秒間のプリフェッチは 7% 未満のオーバーヘッドであり、目に見えません。最も重要なことは、メトリクスのバックエンドが遅いか到達できないかどうかです。

ble の場合、システムはすぐに機能しなくなり、非水和パスにフォールバックするため、濃縮ステップによって結果が妨げられることはありません。
そのフォールバックの本能がプロセス全体に貫かれており、アノテーターは付加的であり、「あると便利」であるという精神が根付いています。Alertmanager は continue: true でアノテーターにルーティングされるため、既存の Slack および電子メール配信プロセスと並行して配置され、それらをブロックすることはありません。 LLM ボックスがダウンしている場合や JSON の形式が不正な場合でも、Webhook は常に 200 を返し、静的なフォールバック アノテーションが代わりに使用されます。最悪のケースは、エンドユーザーがそれほど派手ではないアラートを受け取るものの、決して見逃さないことです。ナレーションは、配信パスにまったく依存しない、配信パスの上に重ねられたアメニティです。
// 価値のあるアドバイス
「分析ではなくモデルに語らせる」というのはレッスンの簡単なバージョンであり、それは事実ですが、難しい部分ではありません。難しい部分は強制です。出力を検証可能な固定セットに制限し、モデルをクリティカル パスから遠ざけることで、失敗した場合に散文が発生し、アラートが表示されなくなります。モデルの価値は判断力ではなく流暢さであり、流暢さはスタックの中で最も置き換え可能なものです。実際のすべての決定をテスト可能なコードにプッシュし、後で列挙型に集約するすべてのものを制約し、モデルをパイプラインの最後の最も置き換え可能なステージとして扱います。明日モデルを交換でき、データがクリーンなままであれば、正しい場所に線を引いたことになります。
私の場合は、数週間にわたってライブインフラストラクチャに対して実行されてきました。散文は良いです。ハードウェアがアップグレードされ、より強力なモデルが導入されれば、より良くなると確信していますが、私がそれを（暫定的に）信頼する理由は、散文が機能していないからです。

## Original Extract

Python classifies into a fixed enum. The LLM writes one sentence. How that split keeps the output clean enough to aggregate and the model off the critical path.

The LLM Narrates. The Code Decides. | Irin Observability
Product ▾
Services
The LLM narrates. The code decides.
Most of the “AI for observability” work I see right now hands the language model the judgment. I think that’s backwards. Feed it the alert, feed it some metrics, ask it what’s wrong, what should be done, and let it make the judgement call. Based on my experience working with language models, I decided that inverting the process provides better results.
The short version: in my alerting pipeline, the set of allowable classifications is fixed in deterministic Python, and the model has to pick from it. The LLM’s only job is to turn a structured verdict into an easily digestible sentence. It never decides whether something is bad, how bad it is, or what category of problem it is. It narrates within a decision space the code has already locked down.
I run a small managed monitoring service. Alertmanager fires, a webhook lands, and historically that webhook produced a line like HighMemoryUsage on host web-vm, severity warning, which is accurate, but not terribly helpful. The person reading it still has to know what HighMemoryUsage implies, whether this host always runs hot, and whether to care. I wanted plain-English context attached to the alerts without altering the alert delivery process.
The obvious move was to throw the whole alert at an LLM and ask it to explain. I tried that in the first iteration of this experiment, expecting it to be somewhat accurate, but not entirely reliable, and it did not disappoint, the model was confidently inconsistent. The same alert, fired three times, produced three different “root cause” categories. One run called a test alert a “Configuration or setup issue,” the next called it “Configuration/Testing,” the next something else again. If you’re storing that output to do any kind of aggregation later (I am, I want to know when three different clients hit the same class of problem in the same week), free-form model output fragments into noise. Grouping on a field the model changes at random won’t work.
I kind of knew from the onset that I wouldn’t get amazing results, and that it would be harder than it looked. So I started doing some research, and decided to flip the design. The little voice in the back of my head was right all along, don’t let the model make the decisions.
I created a pipeline that has a hard wall down the middle.
On the deterministic side, Python does the classifying. The output is constrained to an eight-value enum: memory_pressure , cpu_saturation , disk_pressure , service_unavailability , network_issue , configuration_error , external_dependency , unknown . I aggregate on that field because it can only ever be one of eight strings. If nothing fits, the answer is unknown , which is itself a useful signal rather than a hallucinated/variant guess.
On the narrative side, the LLM ( llama3:8b , running locally on a box on my own LAN, data/network secure) must choose its classification from that fixed eight-value set, and it writes two short fields alongside it: what the alert is, in plain English, and what it means operationally. The code defines the shape of the answer; the model only fills in a slot that already exists. It is explicitly instructed not to suggest fixes and not to invent a cause, so it performs translation instead of analysis.
The prompt returns strict JSON, grammar-constrained, so I get {what, means, likely_cause_class} every time and the enum value is validated against the allowed set on the way out. If the model returns something off-list, I capture the bug instead of storing a row.
A naive version of even the narration step gets you alarmist prose. When tuning the system I received a DiskFillPredicted alert, which on its face looked worth investigating. Then I looked at the host in question, which has had a flat disk-utilization baseline for months. The prediction was a rounding artifact, “your disk is about to fill” is actively misleading. The model had no way to know that from the alert alone, so it just wrote something.
I fixed it by giving the model the same context a human would look at before reacting. Prior to the LLM call, Python does a fast lookup against the metrics backend for that host’s recent baseline, and the prompt carries an explicit rule: if historical context is provided, weigh it over the alert’s literal text. A predicted-disk-fill on a host with a stable months-long baseline is informational, not urgent.
The latency budget for that lookup is five seconds. The LLM call itself takes about eighty seconds, because this runs deliberately on modest CPU-only hardware, a 4th generation Intel i7 with 16GB RAM and no GPU. That is a choice, not a constraint I am apologizing for: the whole posture of the service is that nothing leaves my LAN, so a slow local model beats a fast remote one. And the eighty seconds never reaches the person being alerted. Because the annotator rides alongside the existing path (more on that below), the raw alert lands in Slack and email instantly; the narrated version shows up as a separate annotation a minute or so later. Nothing is ever waiting on the model. Against that eighty-second call, five seconds of pre-fetch is under seven percent overhead and invisible. What mattered most was that if the metrics backend is slow or unreachable, the system fails immediately and falls back to the un-hydrated path, so the enrichment step can never block the result.
That fallback instinct runs through the whole process, the ethos is that the annotator is additive, a ‘nice-to-have.’ Alertmanager routes to it with continue: true , so it sits alongside the existing Slack and email delivery processes, and can never block them. The webhook always returns 200, even when the LLM box is down or when the JSON is malformed, and a static fallback annotation gets used instead. The worst case is that the end-user gets a less flowery alert, but never misses one. The narration is an amenity layered on top of a delivery path that doesn’t depend on it at all.
// ADVICE, FOR WHAT IT’S WORTH
“Let the model narrate, not analyze” is the easy version of the lesson, and it’s true, but it isn’t the hard part. The hard part is the enforcement: constraining the output to a fixed set you can validate, and keeping the model off the critical path so its failures cost you prose and never an alert. The model’s value is fluency, not judgment, and fluency is the most replaceable thing in the stack. Push every actual decision into code you can test, constrain anything you’ll later aggregate down to an enum, and treat the model as the last, most replaceable stage in the pipeline. If you can swap the model out tomorrow and your data stays clean, you’ve drawn the line in the right place.
Mine’s been running against live infrastructure for a few weeks. The prose is good, I’m sure when the hardware is upgraded and a more powerful model is put in place it will be better, but the reason I trust (tentatively) it is that the prose isn’t doing the work.
