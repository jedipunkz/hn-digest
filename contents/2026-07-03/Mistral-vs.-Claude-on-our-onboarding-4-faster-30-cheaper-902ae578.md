---
source: "https://squidler.io/blog/eu-models-1-discovery-mistral"
hn_url: "https://news.ycombinator.com/item?id=48775768"
title: "Mistral vs. Claude on our onboarding: 4× faster, 30% cheaper"
article_title: "Mistral vs. Claude on our onboarding: 4× faster, 30% cheaper | Squidler"
author: "tidbeck"
captured_at: "2026-07-03T15:50:42Z"
capture_tool: "hn-digest"
hn_id: 48775768
score: 2
comments: 0
posted_at: "2026-07-03T14:51:57Z"
tags:
  - hacker-news
  - translated
---

# Mistral vs. Claude on our onboarding: 4× faster, 30% cheaper

- HN: [48775768](https://news.ycombinator.com/item?id=48775768)
- Source: [squidler.io](https://squidler.io/blog/eu-models-1-discovery-mistral)
- Score: 2
- Comments: 0
- Posted: 2026-07-03T14:51:57Z

## Translation

タイトル: ミストラル vs. クロードのオンボーディング: 4 倍高速、30% 安価
記事のタイトル: オンボーディングでのミストラル vs. クロード: 4 倍高速、30% 安価 |イカドラー
説明: 私たちは、独自のワークロードで、オンボーディング エージェントの背後にあるモデルである Claude Sonnet 4.6 と Mistral Medium 3.5 を比較しました。 Mistral は、ほぼ同等の品質で、約 4 倍の速度で 30% 安く登場しました。この特定の仕事では、オンボーディングの仕組みに特有の理由により、これが驚くほど適していることが判明しました。
[切り捨てられた]

記事本文:
オンボーディングでのミストラル対クロード: 4 倍高速、30% 安価 | Squidler の使用例 価格 Replit パートナーについて サインイン ← ブログ オンボーディングにおけるミストラル対クロード: 4 倍高速、30% 安価
私たちは、独自のワークロード上で、オンボーディング エージェントの背後にあるモデルである Claude Sonnet 4.6 と Mistral Medium 3.5 を比較しました。 Mistral は、ほぼ同等の品質で、約 4 倍の速度で 30% 安く登場しました。この特定のジョブでは、オンボーディングが実際にどのように実行されるかに特有の理由により、これが驚くほど適していることがわかりました。
私たちはマルメに拠点を置いており、ヨーロッパの技術スタックを常に望んでいました。たとえば、これが Hetzner を使用する理由の 1 つです。目標は全ヨーロッパのスタックであり、私たちはそれを推進していきます。しかし、素晴らしい製品を作りたいという意欲も同様に強いので、欧州のオプションではまだ対応できない場合には、入手可能な最高のモデルを採用し、それが EU の法律に基づいて動作することを確認します。 EU の規則に基づいてデータを処理しない場合、多くの潜在的な顧客にとって、当社は選択肢にさえなりません。
しかし、AI に関しては、フロンティアの大部分はアメリカ人であり、私たちの初期の調査では、スタック内のほとんどのモデルが Anthropic の Claude にたどり着きました。私たちは極めて速いスピードで今日の地位に到達しましたが、今は速度を落として、どのようにして全ヨーロッパの目標を達成できるかを検討するときです。
Squidler スペシャリストを雇うときの最初の会話は、実際に必要なことを解決する新人研修エージェントによって進められます。システムの他の部分からかなり分離して実行されるため、最初に使用するのに適した候補のように感じられます。
現在、オンボーディングは Claude Sonnet 4.6 で実行されており、タスクの品質、速度、価格のバランスが取れています。当初、Anthropic の低価格帯である Haiku 4.5 も評価しましたが、期待する品質は得られませんでした。それに匹敵する欧州モデルを探すにはw

公開されているベンチマークを抽出しました (機能と速度については人工分析、価格についてはベンダー)。一人の候補者が浮かび上がった。
各ベンダー独自の API の 2026 年 7 月 2 日時点の数値。新しいベンチマークが着陸すると、彼らは漂います。
ミストラル ミディアム 3.5 は、インテリジェンスに関してソネットより 6 ポイント遅れており、価格は半額です。最初のトークンまでの時間、つまり画面に何かが表示されるまでの待機時間は Sonnet よりも少し遅いですが、その出力速度は Haiku と同等であり、Sonnet の 2 倍であるため、完全な応答をより速く生成できるはずです。理論上、ミストラルはテストすべきものであるのは明白です。フランスで作成され、EU でホストされているデータを使用して独自の API を通じて提供されます。品質が保てれば、全ヨーロッパのオプション。
私たちはモデルを雰囲気で判断しません。オンボーディングには独自の評価スイートがあります。17 個のスクリプト化された会話フィクスチャがそれぞれ 3 回実行され、重要な点については、promptfoo でスコア付けされます。エージェントは具体的な問題を表面化し、きちんと引き継ぎ、複数の言語を扱い、話をでっち上げるのではなく地に足を着いたままでいますか?
それぞれのチェックは、全体として 1 つのフィクスチャの実行で合否を採点されたものであり、クロード モデルによる主観的なものであり、どちらかといえば現職者に有利なバイアスです。 51 回の実行における 2 つのチェックのギャップは、実行間のノイズの範囲内にあります。速度とコストは、モデルごとに 8 ターンのセッションを再生することで得られます。
品質に関しては、公開されている数値と一致し、Mistral Medium 3.5 はテストで Sonnet とほぼ同等の結果を出しました。驚くべきは速度です。1 ターンあたりの速度は約 4 倍で、スペック シートの約束の 2 倍です。 eval から生の応答を取得すると、それが説明されます。会話が始まると、Mistral はターンあたり平均約 170 の出力トークンを使用しましたが、Sonnet は約 500 を使用しました。同じ仕事、言葉の3分の1。以下の例のように:
次の返信はモデルの、両方とも同じ会話です
それは本当に問題です - 何か重要なものが欠けているのです

埋もれてるのがストレスになる。主な苦痛は、行動する必要があるもの (返信、フォローアップ、何かをすること) を逃していることですか、それとも後で必要なときに情報が見つからないことですか?
それは本当に苦痛のようですね。 Gmail または他の電子メール サービスを使用していますか?
読者の中には、ここでのソネットのより詳しい返答を好む人もいるでしょうが、それは当然です。どちらも同じチェックに合格します。
注目に値する: ターンごとの数値は合計時間であり、体感的な待ち時間は測定された待ち時間と同じではありません。最初に感じるのは、最初のトークンの時間です。仕様書ではミストラルは最初の単語までが遅いとされていますが、私たちはそれを感じたことはありません。ミストラルのターンの中央値は、仕様書に記載されている最初のトークンまでの時間よりも短い時間で終了しました。公表されているレイテンシの数値を答えではなく出発点として取り上げます。
コストに関して言えば、モデルカードを読み取るほど簡単ではありません。他の 2 つの要素、キャッシュと出力トークンが考慮されます。ミストラルとアンスロピックの価格は両方とも、通常のレートの 10 分の 1 で入力をキャッシュしました。問題は、キャッシュはヒットした場合にのみ役立つということです。私たちの経験では、Claude はこの点で優れていましたが、Mistral は、独自のドキュメントによる最善の努力を尽くして、私たちが直接試したときに、キャッシュされたトークンを一度も提供しませんでした。オンボーディングの場合、ミスによって結果が反転することはありません。キャッシュは入力を割引するだけで、Sonnet の請求額は主に長い応答で、Mistral はすべての入力トークンの全額を支払ってもセッションあたり約 30% 安くなります。ミストラルのキャッシュが配信を開始すれば、その数値はさらに向上します。
驚くほどうまくいきました。当社独自のオンボーディング ワークロードでは、ヨーロッパ モデルがクロードと同等の品質を維持し、低コストで提供され、待ち時間が著しく短くなりました。私たちが毎日実行しているジョブにとって、これは予想していなかった真のオプションです。
ほぼ同じ性能で、4 倍高速、30% 安く、しかもヨーロッパ製です。
ソネット 5: Anthropic sh についてのメモ

これを実行したのと同じ週に作成し、それに対する評価を再実行しませんでした。現在オンボーディングで実際に実行されているものと比較します。
© Squidler AB · プライバシー · 利用規約 · ブログ · Cookie 設定 · Replit パートナー

## Original Extract

We put Mistral Medium 3.5 up against Claude Sonnet 4.6, the model behind our onboarding agent, on our own workload. Mistral came out about 4× faster and 30% cheaper, at near-equal quality. For this particular job it turned out to be a surprisingly good fit, for reasons specific to how onboarding act
[truncated]

Mistral vs. Claude on our onboarding: 4× faster, 30% cheaper | Squidler Use cases Pricing About Replit partner Sign in ← Blog Mistral vs. Claude on our onboarding: 4× faster, 30% cheaper
We put Mistral Medium 3.5 up against Claude Sonnet 4.6, the model behind our onboarding agent, on our own workload. Mistral came out about 4× faster and 30% cheaper, at near-equal quality. For this particular job it turned out to be a surprisingly good fit, for reasons specific to how onboarding actually runs.
We're based in Malmö, and we've always wanted a European tech stack. It's part of why we run on Hetzner , for one. The goal is an all-European stack , and we'll push for it. But the drive to build an awesome product is equally strong, so where a European option can't yet carry the job, we'll take the best model available and make sure it runs under EU legislation. We are not even an option for many of our potential customers if we don't handle their data under EU rules.
But with AI, most of the frontier is American , and our early investigations landed us on Anthropic 's Claude for most of the models in our stack. It has taken us to where we are today with extreme speed, but now it's time to slow down and see how we can reach that all-European goal.
The first conversation you have when you hire a Squidler specialist is driven by an onboarding agent that works out what you actually need. As it runs fairly isolated from the rest of the system, it feels like a good candidate to start with.
Onboarding runs on Claude Sonnet 4.6 today, as it strikes a good balance between quality, speed, and price for the task. Initially we also evaluated Haiku 4.5, Anthropic's cheaper tier, but didn't get the quality we wanted. To find a European model that could stand next to it, we pulled the published benchmarks ( Artificial Analysis for capability and speed, the vendors for price). One candidate stood out.
Numbers as of July 2, 2026, for each vendor's own API. They drift as new benchmark runs land.
Mistral Medium 3.5 sits six points behind Sonnet on intelligence, at half the price. Its time to first token, the wait before anything shows on screen, is a bit slower than Sonnet's, but its output speed matches Haiku's, twice Sonnet's, so it should produce the complete reply quicker. On paper, Mistral is the obvious one to test: created in France, and served through their own API with data hosted in the EU. The all-European option, if the quality holds.
We don't judge models on vibes. Onboarding has its own eval suite: 17 scripted conversation fixtures, each run three times, scored with promptfoo on the things that matter. Does the agent surface a concrete problem, hand off cleanly, handle multiple languages, and stay grounded instead of making things up?
Each check is one fixture run graded pass/fail as a whole, the subjective ones by a Claude model, a bias that favors the incumbent if anything. A two-check gap on 51 runs is within run-to-run noise. Speed and cost come from replayed eight-turn sessions per model.
On quality, Mistral Medium 3.5 came out almost as good as Sonnet in our tests, in line with the published numbers. The surprise is speed: about 4× as quick per turn , double what the spec sheet promised. Pulling the raw replies from the eval explains it: once the conversation gets going, Mistral averaged roughly 170 output tokens per turn where Sonnet used about 500 . Same job, a third of the words. As in the example below:
The next reply is the model's, same conversation for both
That's a real problem — missing something important because it's buried is stressful. Is the main pain that you're missing things you need to act on (reply, follow up, do something), or more that you can't find information when you need it later?
That sounds like a real pain. Do you use Gmail or another email service?
Some readers will prefer Sonnet's fuller reply here, and that's fair. Both pass the same checks.
Worth noting: the per-turn figure is total time, and felt latency is not the same as measured latency. What you feel first is time to first token. The spec sheet has Mistral slower to the first word, but we never felt it: a median Mistral turn finished in less time than the spec sheet's quoted time to first token. Take published latency numbers as a starting point, not an answer.
When it comes to cost, it is not as easy as reading the model card. Two other factors weigh in: caching and output tokens. Both Mistral and Anthropic price cached input at a tenth of the normal rate. The catch is that a cache only helps if it hits: Claude has been outstanding at this in our experience, while Mistral's, best-effort by its own docs, never gave us a single cached token when we tried it head-on. For onboarding the miss doesn't flip the result: caching only discounts input, Sonnet's bill is mostly its long replies, and Mistral lands about 30% cheaper per session even paying full price for every input token. If Mistral's cache ever starts delivering, that number only improves.
It turned out surprisingly well. On our own onboarding workload, a European model held its own with Claude on quality, came in at a lower cost, and answered with noticeably lower latency. For a job we run every day, that's a genuine option we didn't expect to have.
Almost as good, 4× faster, 30% cheaper, and it's European.
A note on Sonnet 5: Anthropic shipped it the same week we ran this, and we didn't rerun the eval against it. We compare against what actually runs onboarding today.
© Squidler AB · Privacy · Terms · Blog · Cookie settings · Replit partner
