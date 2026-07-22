---
source: "https://dylancastillo.co/posts/pelicanmaxxing.html"
hn_url: "https://news.ycombinator.com/item?id=49010129"
title: "Are AI Labs Pelicanmaxxing?"
article_title: "Are AI labs pelicanmaxxing? – Dylan Castillo"
author: "dcastm"
captured_at: "2026-07-22T18:04:27Z"
capture_tool: "hn-digest"
hn_id: 49010129
score: 3
comments: 0
posted_at: "2026-07-22T17:17:54Z"
tags:
  - hacker-news
  - translated
---

# Are AI Labs Pelicanmaxxing?

- HN: [49010129](https://news.ycombinator.com/item?id=49010129)
- Source: [dylancastillo.co](https://dylancastillo.co/posts/pelicanmaxxing.html)
- Score: 3
- Comments: 0
- Posted: 2026-07-22T17:17:54Z

## Translation

タイトル: AI 研究所は Pelicanmaxxing を行っていますか?
記事タイトル：AI研究所はペリカンマックス化しているのか？ – ディラン・カスティージョ
説明: AI ラボが Simon Willison のペリカン自転車に乗るベンチマークに基づいてトレーニングしているかどうかをテストするために、7 つのフロンティア モデルにわたって 1,000 以上の SVG を生成しました。

記事本文:
証拠 #1: 自転車に乗ったペリカンの見た目はそれほど良くありません
証拠 #2: 研究室はペリカンを描くのが得意ではありません
証拠 #3: 研究室の方が自転車を描くのが得意ではない
証拠 #4: 自転車に乗ったペリカンを描くのは、難易度を調整しても研究室のほうが上手ではありません。
証拠 #5: ペリカン自転車のシーンは記憶に残っていないように見える
過去数年間、Simon Willison はすべてのメジャー LLM リリースを同じプロンプト「自転車に乗っているペリカンの SVG を生成する」でテストしてきました。
冗談めいたベンチマークとして始まったベンチマークは、AI で最も有名な非公式ベンチマークの 1 つになりました。サイモンの自転車に乗ったペリカンの結果は、AI 研究所からの新しいリリースを発表するハッカー ニュース スレッドで最も賛成票を集めたコメントの 1 つとなります。
このベンチマークは現在、その有用性について、また AI ラボがベンチマーク 1 をベンチマークしているかどうかについて、多くの議論が行われているほど有名です。何十億ドル、さらには何兆ドルもかかっていて、強力な結果がユーザーを説得するのに役立つ可能性がある場合、モデルを少しでもペリカンマックス化したくなりませんか?
それを知りたかったので、ちょっとした実験をしてみました。 7 つのフロンティア モデルにわたって 1,008 個の SVG を生成し、LLM ジャッジでスコアを付け、分析には Claude Fable 5 を使用しました。
この記事ではその結果を紹介します。すべてのコードは Github で入手できます。
8 匹の動物 × 6 台の車両 = 48 個のプロンプトのグリッドを作成しました。有名なプロンプトは 1 つのセルです。
動物 : ペリカン、フラミンゴ、サギ、カワウソ、アライグマ、アンテロープ、クジラ、猫
乗り物 : 自転車、一輪車、スケートボード、スクーター、飛行機、ボート
すべてのプロンプトでは、動物と乗り物が切り替わっているだけで、サイモンのものとほぼ同じフレーズが使用されています。動物と乗り物の選択はそれほど厳密な方法では行われませんでしたが、元のプロンプトとの類似性と難易度の両方を変えるように努めました。フラミンゴとh

エロンはペリカンによく似ています。猫、アライグマ、カワウソは簡単なケースです。アンテロープは硬いです。そしてクジラはあなたが得ることができる限り異なります。
OpenRouter を通じて 7 つのモデル (GPT-5.6 Terra、Claude Sonnet 5、Gemini 3.5 Flash、Grok 4.5、Qwen3.7-Max、GLM-5.2、および DeepSeek V4 Pro) をテストしました。温度 1.0 でプロンプトごとに 3 つのサンプルを生成し、すべてのモデルに同じ推論作業を要求しました。その結果、SVG は 1,008 個になりました。
次に、3 段階のパイプラインを通じて各イメージを実行しました。
レンダリング : 各 SVG は PNG にレンダリングされます。モデルが SVG を返さないか、レンダリングに失敗した場合は、有効な SVG が生成されるまで再生成し、試行回数を記録します。 1,008 世代にわたって再試行はわずか 11 回でした。
判定 : GPT-5.6 ルナは、動物、乗り物、およびアクションの一貫性について 1 ～ 5 の評価で各画像を採点します。以下で動物や乗り物をランク付けするときは、一致する評価を単独で使用します。画像ごとに 1 つの数値が必要な場合は、3 つの平均を使用します。これを審査員スコアと呼びます。
特徴抽出 : より詳細な分析のために、各レンダリング画像を Gemini 3.1 Flash-Lite に渡し、認識した動物と乗り物、被写体の向き、およびシーン要素の無制限のリストを記録しました。
私の仮説は、研究室がベンチマークでトレーニングした場合、ペリカンの行のスコアが動物の値を上回っているか、自転車の列のスコアが車両の値を上回っているか、特定のペリカンと自転車のセルがその両方を上回っているかの組み合わせで現れるはずだ、というものです。
証拠 #1: 自転車に乗ったペリカンの見た目はそれほど良くありません
採点する前に、最も簡単なテストは画像を自分で見ることです。ラボを選択すると、そのラボが描いたすべてが表示され、各画像の下に審査員のスコアが表示されます (クリックするとフルサイズが開きます)。
実行する前に自分で画像を調べました

彼の分析は以下の通り。 Nothing jumped out at me.ペリカン自転車の画像が、そのモデルのグリッドの他の部分よりも著しく優れているという例は見つかりませんでした。おそらく GLM-5.2 の最初のサンプルでは残りのサンプルよりもわずかに優れていると感じましたが、そのバッチではスケートボードに乗った非常にクールなサギも生成されたため、確かなことは言えません。それ以外の点では、各モデルが描いた残りの部分と同じように見えます。自転車に乗ったペリカンを上手に描いた研究室は、他の動物と乗り物の組み合わせも上手に描いています。
しかし、このテストを再現するのは難しく、人によって意見が異なるでしょう。したがって、より定量的なものが必要だったので、上記で詳しく説明した方法を選択しました。
証拠 #2: 研究室はペリカンを描くのが得意ではありません
すべてのモデルにわたってプールされた、動物ごとの平均評価は次のとおりです。
ペリカンは、ネコ、クジラ、アライグマ、サギ、アンテロープに次ぐ 8 頭中 6 位です。 AI ラボがベンチマークに基づいてトレーニングを行っていた場合、ペリカンがトップになることが予想されます。 Instead they’re in the bottom half. 7 つのラボはすべて、ペリカンよりもネコ、クジラ、アライグマの方が上手に描けています。
もちろん、ペリカンは猫よりも描くのが難しいだけかもしれません。研究室がペリカンを使って訓練しても、ペリカンが扱いやすい動物を追い越せない可能性があるため、このランキングだけでその可能性を排除することはできません。 I’ll adjust for difficulty in Evidence #4.
証拠 #3: 研究室の方が自転車を描くのが得意ではない
Bicycles fare even worse.彼らは最後から 2 番目に位置し、最後に来る飛行機とほぼ同点です。
研究室がベンチマークに基づいてトレーニングを行っている場合、自転車がこのランキングの上位に入ることが予想されます。そうではありません。 However, the same caveat applies here.自転車はスケートボードよりも描くのが難しいです。自転車には、対応する 2 つの車輪、両方の車軸に届くフレーム、ハンドルバー、シート、ペダルが必要です。裁判官は、b の 2/3 の 1 つが行方不明または切断されているとフラグを立てます。

アイサイクルの画像。自転車の画像を使ってトレーニングすることはできますが、単純な車両と比べて優れた成果を上げることはできません。
ただし、飛行機に関する注意点が 1 つあります。モデルは幾何学的に解釈することが多いため、「飛行機」ではなく「飛行機」を選択するべきでした。彼らは、飛行機を飛ばす代わりに、平らな面に立つ動物を描きました。飛行機は、特徴抽出機能によって車両がまったく検出されないことがある唯一の乗り物であり (168 枚の画像中 25 枚、他の 5 枚はゼロ)、飛行機画像の 20% が乗り物評価で 1 または 2 のスコアを獲得しましたが、自転車については 5% で、ボート、スクーター、またはスケートボードについてはまったくありませんでした。
証拠 #4: 自転車に乗ったペリカンを描くのは、難易度を調整しても研究室のほうが上手ではありません。
2 つを組み合わせると、「自転車に乗ったペリカン」はランキングの最下位近く、48 件中 42 位になります。
ただし、繰り返しになりますが、一部の組み合わせは他の組み合わせよりも描くのが難しい場合があります。
これを説明するために、1,008 枚の画像すべてに対して固定効果回帰を当てはめました。スコア ~ 研究室 + 動物 × 車両、およびペリカン、自転車、ペリカン自転車セルの研究室ごとの交互作用項をロバストな標準誤差で当てはめました。動物×乗り物の用語は、48 通りの組み合わせすべてに特有の難しさを吸収します。この相互作用は、信頼区間を使用して、平均的なラボと比較した各ラボのベンチマーク固有の向上を測定します。
ラボごとのすべてのペリカン効果 (6 台すべてのペリカンに対するラボのブースト) は、-0.11 ～ +0.14 の判定ポイントの間に収まり、どれも有意性に近いものはありません (最小 p = 0.25)。
ラボごとの自転車効果 (8 匹すべての動物にわたるラボの自転車効果) は、Grok 4.5 の -0.18 (p=0.11) から Gemini 3.5 Flash の +0.27 (p=0.022) まで続きます。 p < 0.05 をクリアしているのは双子座だけであり、7 人は両方向を向いています。
ペリカン自転車セル効果なし (特定の組み合わせでの追加ブースト、

研究室のペリカン効果と自転車効果のトップ）は p < 0.05 をクリアします。最大のプラスは GLM-5.2 の +0.35 (p=0.12) で、これは先ほど述べたものです。これはこの実験の信号に最も近いものですが、それでも可能性の範囲内です。
ラボごとの完全な推定値は次のとおりです。 pelicanmaxxing ラボでは、行全体にわたってゼロ線の右側にドットが表示されます。
すべてのペリカン間隔とすべてのセル間隔にはゼロが含まれます。自転車コラムの Gemini 3.5 Flash は、そうではありません。しかし、p < 0.05 の 21 回の検査では、確率だけで約 1 回の偽陽性 (21 × 0.05 ≈ 1.05) が予測され、その 1 回がまさにその結果となります。また、多重比較補正にも耐えられません。21 のテストにわたるボンフェローニ閾値は 0.05/21 ≈ 0.002 で、その p 値は 0.022 です。推定値と p 値の完全な表は、リポジトリにあります。
ただし、これらの間隔は広く、平均で約 ±0.6 ジャッジ ポイントです。それより小さいブーストはこのテストでは捕捉されません。
証拠 #5: ペリカン自転車のシーンは記憶に残っていないように見える
自転車に乗ったペリカンが記憶された構図のように見えると示唆する人もおり、ペリカンが常に右を向いているなどの繰り返しのパターンや、太陽やスカーフなどの繰り返しの要素を指摘しています。それで、これが本当かどうか知りたかったのです。
方向 : 7 つのラボすべてにわたる 21 枚のペリカン自転車の画像はすべて右を向いています。これほどの動物と乗り物の組み合わせは他にはありません。
ただし、右向きであることは一般的であり、全 1,008 枚の画像の 60% が右向きです。どの程度一般的かは動物と乗り物によって異なりますが、自転車はその傾向が最も強い 2 つの乗り物のうちの 1 つです。
ペリカンも右を向く傾向がある動物です。
ペリカンや自転車を正面を向いて描くのは難しいため、モデルはほとんどの場合、左または右を向いて横から描きます。それが彼らの画像が非常に少ない理由です

曖昧です。他の組み合わせもほぼ一致しています。スクーターに乗ったカモシカとスクーターに乗ったペリカンは 21 点中 20 点で着地し、自転車に乗ったアオサギは 21 点中 19 点でした。したがって、21 点中 21 点は異常値ではないようです。
シーン要素: エクストラクターに画像内にある要素に名前を付けさせます。カウントは次のとおりです。
記憶されたシーンは、次から次へと繰り返される同じ要素のセットとして表示されます。それを探してみたところ、いくつかの組み合わせは毎回同じ要素を生成する傾向があることがわかりました。ボートに乗っているフラミンゴはどれも太陽を持っています。飛行機に乗るカワウソは 38% の確率でスカーフを着用しています。自転車に乗った猫は 38% の確率でカゴを手に入れます。
自転車に乗ったペリカンは特に変わったところはないようだ。他の動物と乗り物の組み合わせと同様に、より頻繁に現れる要素がいくつかあるだけです。
採点には 1 人の LLM ジャッジを使用します。ここでのすべてのスコアは、一度に 1 つの画像を調べた 1 つのモデル GPT-5.6 Luna から得られています。調整はあまり行わず、再実行時にどの程度一致するかは確認しませんでした。モデルが図面を確実に判断できない場合、上記の数値はどれもあまり意味がありません。審査員も出場者の一人である GPT-5.6 Terra と同じ家族の出身です。ただし、どのラボも 48 通りの組み合わせをすべて引くため、あるラボのスタイルをたまたま気に入った審査員が、そのラボのグリッド全体を一度に引き上げます。ただし、この分析ではラボ内の違いのみが考慮されるため、結果は変わりません。
SVGmaxing。 SVG 生成全体 (または車両に乗った動物などのサブセット) を最適化したラボは、すべてのセルで同時に立ち上がり、単に優れたラボと同じように見えます。 Google/DeepMind などの一部のラボでは、これを公然と行っています。この実験ではそれを検出できません。
限られた予算。実験全体は約 80 ドルの API クレジットを使用して実行されました

s.これにより、セルあたり 3 つのサンプル、1 人の審査員、および 7 つのモデルが上限となりました。これにより、「飛行機」と「飛行機」の場合のように、プロンプトとパイプラインを何度も繰り返すこともなくなりました。
HN 嫌いの皆さん、申し訳ありませんが、AI 研究所がペリカンマックスを行っているという証拠はほとんどありません。少なくとも、彼らはそれを明白な方法で行っているわけではありません。
ペリカンは他の動物よりも上手に描かれているわけではありません。自転車は他の乗り物と比べてうまく描かれているわけではありません。そして、ペリカンと自転車がすでに予測している以上にその組み合わせをうまく引き出している研究室はありません。 GLM-5.2 はこれに最も近いもので、まさにペリカン自転車のセルを最大限に強化しており、その最初のペリカン自転車サンプルが私の目に留まりました。ただし、効果は小さく、重要ではないため、あまり重視しません。
もう一つ際立っているのは、シーン構成の方向性です。 21 枚のペリカン自転車の画像はすべて右を向いており、グリッド内ですべての画像が一致する唯一の組み合わせです。しかし、それはそれほど奇妙ではないようです。実験全体では右向きが標準です。他の 3 つの組み合わせは 90% 以上に達し、そのうち 48 個あるので、21 個中 21 個に達したのは驚くことではありません。
よりもっともらしい話は、Google/DeepMind のような SVGmaxxing です。他の研究室はもっと静かにやっているかもしれない。残念ながら、この実験を誰が行っているのかはわかりません。でも少なくとも今夜はこれを知って眠れます

[切り捨てられた]

## Original Extract

I generated 1,000+ SVGs across 7 frontier models to test whether AI labs are training on Simon Willison’s pelican-riding-a-bicycle benchmark.

Evidence #1: The pelicans on bicycles don’t look any better
Evidence #2: Labs are not better at drawing pelicans
Evidence #3: Labs are not better at drawing bicycles
Evidence #4: Labs are not better at drawing pelicans on bicycles, even adjusting for difficulty
Evidence #5: The pelican-bicycle scenes don’t look memorized
For the past few years, Simon Willison has tested every major LLM release with the same prompt: “Generate an SVG of a pelican riding a bicycle”.
What began as a tongue-in-cheek benchmark has become one of the most famous informal benchmarks in AI. Simon’s pelican-on-a-bicycle results are often among the most upvoted comments on Hacker News threads announcing new releases from AI labs.
The benchmark is now famous enough that there’s plenty of discussion about its usefulness and about whether AI labs might be benchmaxxing 1 on it. When billions or even trillions of dollars are at stake, and a strong result could help persuade users, wouldn’t it be tempting to pelicanmaxx your model just a bit?
I wanted to find out, so I put together a small experiment. I generated 1,008 SVGs across seven frontier models, scored them with an LLM judge, and used Claude Fable 5 for the analysis.
This article presents the results. All the code is available on Github .
I built a grid of 8 animals × 6 vehicles = 48 prompts, where the famous prompt is one cell:
Animals : pelican, flamingo, heron, otter, raccoon, antelope, whale, cat
Vehicles : bicycle, unicycle, skateboard, scooter, plane, boat
Every prompt uses almost identical phrasing to Simon’s, only switching the animal and vehicle. The animal and vehicle selection wasn’t done in a very rigorous manner, but I tried to vary both similarity to the original prompt and difficulty. Flamingo and heron are quite similar to pelicans; cat, raccoon, and otter are easy cases; antelope is hard; and whale is as different as you can get.
I tested seven models through OpenRouter : GPT-5.6 Terra , Claude Sonnet 5 , Gemini 3.5 Flash , Grok 4.5 , Qwen3.7-Max , GLM-5.2 , and DeepSeek V4 Pro . I generated 3 samples per prompt, at temperature 1.0, requesting the same reasoning effort from every model. That resulted in 1,008 SVGs.
Then I ran each image through a three-stage pipeline:
Rendering : Each SVG is rendered to PNG. If a model returns no SVG or one that fails to render, I regenerate until it produces a valid one, and record the number of attempts. There were only 11 retries across the 1,008 generations.
Judging : GPT-5.6 Luna scores each image with 1-5 ratings for the animal, the vehicle, and the coherence of the action. When I rank animals or vehicles below, I use the matching rating on its own. When I need one number per image, I use the average of the three, which I call the judge score.
Feature extraction : For a more detailed analysis, I also passed each rendered image to Gemini 3.1 Flash-Lite , which recorded the animal and vehicle it recognized, which way the subject faces, and an open-ended list of scene elements.
My hypothesis is that if a lab trained on the benchmark, it should show up in some combination of the pelican row scoring above what the animal deserves, the bicycle column scoring above what the vehicle deserves, or the specific pelican-bicycle cell beating both.
Evidence #1: The pelicans on bicycles don’t look any better
Before any scoring, the simplest test is to look at the images yourself. Pick a lab to see everything it drew, with the judge’s score under each image (click to open full size):
I looked through the images myself before running the analysis below. Nothing jumped out at me. I couldn’t find a case where the pelican-bicycle images looked noticeably better than the rest of that model’s grid. Maybe in GLM-5.2’s first sample it felt slightly better than the rest, but that batch also produced a pretty cool heron on a skateboard, so I cannot say for sure. Otherwise they look like the rest of what each model draws, and the labs that draw good pelicans on bicycles also do a good job drawing other animal-vehicle combinations.
But this test is hard to replicate, and everyone will have a different opinion. So I wanted something more quantitative, which is why I opted for the method detailed above.
Evidence #2: Labs are not better at drawing pelicans
Here’s the mean animal rating per animal, pooled across all models:
The pelican is 6th of 8, behind cat, whale, raccoon, heron, and antelope. If AI labs were training on the benchmark, you’d expect pelicans at the top. Instead they’re in the bottom half. All seven labs draw cats, whales, and raccoons better than pelicans.
Of course, a pelican may simply be harder to draw than a cat. A lab could train on pelicans and still not push them past the easy animals, so this ranking alone can’t rule that out. I’ll adjust for difficulty in Evidence #4.
Evidence #3: Labs are not better at drawing bicycles
Bicycles fare even worse. They sit second from last, in a near-tie with planes, which come in last:
If labs were training on the benchmark, you’d expect bicycles near the top of this ranking. They’re not. However, the same caveat applies here. A bicycle is harder to draw than a skateboard: it needs two matching wheels, a frame that reaches both axles, handlebars, a seat, and pedals. The judge flags a missing or disconnected one of those on 2/3 of the bicycle images. You can train on bicycle images and still not do a great job relative to simpler vehicles.
One note on the plane, though: I should’ve picked “airplane” instead of “plane” because models often read it geometrically. They drew the animal standing on a flat surface instead of flying an aircraft. The plane is the only vehicle where the feature extractor sometimes found no vehicle at all (25 of 168 images, against zero for the other five), and 20% of plane images scored a 1 or 2 on the vehicle rating, against 5% for bicycles and none at all for boats, scooters, or skateboards.
Evidence #4: Labs are not better at drawing pelicans on bicycles, even adjusting for difficulty
Put the two together and the “pelican on a bicycle” ends up near the bottom of the ranking, at #42 of 48:
But again, some combinations might be just harder to draw than others.
To account for that, I fit a fixed-effects regression on all 1,008 images: score ~ lab + animal × vehicle, plus per-lab interaction terms for pelican, bicycle, and the pelican-bicycle cell, with robust standard errors. The animal × vehicle terms absorb the inherent difficulty of all 48 combinations. The interactions measure each lab’s benchmark-specific boost relative to the average lab, with confidence intervals.
Every per-lab pelican effect (the lab’s boost on pelicans across all six vehicles) lands between -0.11 and +0.14 judge points, and none comes close to significance (smallest p = 0.25).
The per-lab bicycle effects (the lab’s boost on bicycles across all eight animals) run from Grok 4.5 at -0.18 (p=0.11) to Gemini 3.5 Flash at +0.27 (p=0.022). Only Gemini clears p < 0.05, and the seven point in both directions.
No pelican-bicycle cell effect (the extra boost on the specific combination, on top of the lab’s pelican and bicycle effects) clears p < 0.05. The largest positive is GLM-5.2 at +0.35 (p=0.12), which is the one I mentioned earlier. It’s the closest thing to a signal in this experiment, but still within chance.
Here are the full per-lab estimates. A pelicanmaxxing lab would show dots to the right of the zero line across its whole row:
Every pelican interval and every cell interval contains zero. Exactly one doesn’t: Gemini 3.5 Flash in the bicycle column. But with 21 tests at p < 0.05, chance alone predicts about one false positive (21 × 0.05 ≈ 1.05), and one is exactly what came up. It also doesn’t survive a multiple-comparisons correction: the Bonferroni threshold across the 21 tests is 0.05/21 ≈ 0.002, and its p-value is 0.022. The full table of estimates and p-values is in the repo .
But these intervals are wide, about ±0.6 judge points on average. Any boost smaller than that won’t be captured by this test.
Evidence #5: The pelican-bicycle scenes don’t look memorized
Some have suggested that the pelican on a bicycle looks like a memorized composition, pointing to recurring patterns such as the pelican always facing right, or recurring elements like a sun or a scarf. So I wanted to know if this was true.
Direction : All 21 pelican-bicycle images, across all seven labs, face right. No other animal/vehicle combination does that.
However, facing right is common: 60% of all 1,008 images do it. How common depends on the animal and the vehicle, and bicycles are one of the two vehicles where it’s strongest:
Pelicans are also among the animals that tend to face right:
It’s hard to draw a pelican or a bicycle facing the viewer, so models almost always draw them from the side, facing left or right. That’s why so few of their images are ambiguous. Other combinations also come close to unanimous: antelope on a scooter and pelican on a scooter land at 20 of 21, and heron on a bicycle at 19 of 21. So 21 out of 21 doesn’t seem like an outlier.
Scene elements : I let the extractor name any element it saw in the image. These are the counts:
A memorized scene would show up as the same set of elements recurring picture after picture. I went looking for that, and found some combinations do tend to produce the same elements every time. Every single flamingo on a boat has a sun in it. Otters on planes wear scarves 38% of the time. Cats on bicycles get a basket 38% of the time.
The pelican on a bicycle doesn’t seem to have anything particularly different about it. It just has some elements that appear more frequently, like every other animal-vehicle combination.
Using a single LLM judge for scoring. Every score here comes from one model, GPT-5.6 Luna , looking at one image at a time. I didn’t do much alignment and didn’t check how often it agrees with itself on a re-run. If a model just can’t judge a drawing reliably, none of the numbers above mean much. The judge is also from the same family as one of the contestants, GPT-5.6 Terra . However, every lab draws all 48 combinations, so a judge that happens to like one lab’s style lifts that lab’s whole grid at once. But that doesn’t change the results because this analysis only cares about the within-lab differences.
SVGmaxxing. A lab that optimized SVG generation as a whole (or a subset such as animals on vehicles) rises on every cell at once and looks identical to a lab that’s just good. Some labs, such as Google/DeepMind, openly do this. This experiment can’t detect that.
Limited budget. The whole experiment ran on roughly $80 of API credits. That capped it at 3 samples per cell, a single judge, and 7 models. This also prevented me from iterating too much on the prompts and pipeline, as with the “plane” vs. “airplane” case.
Sorry, HN haters, but there’s little evidence that AI labs are pelicanmaxxing. Or at least they’re not doing it in a plainly obvious manner.
Pelicans aren’t drawn any better than other animals. Bicycles aren’t drawn any better than other vehicles. And no lab draws the combination better than its pelicans and bicycles already predict. GLM-5.2 comes closest: it has the largest boost on the exact pelican-bicycle cell, and and its first pelican-on-bicycle sample caught my eye. But the effect is small and not significant, so I wouldn’t put too much weight on it.
The other thing that stands out is direction in the scene composition. All 21 pelican-bicycle images face right, the only combination in the grid where every image agrees. But it doesn’t seem that strange. Facing right is the norm across the experiment. Three other combinations land at 90% or above, and with 48 of them, I’m not surprised one reached 21 out of 21.
The more plausible story is SVGmaxxing like Google/DeepMind does. Other labs might be doing it more quietly. Sadly, this experiment can’t say who’s doing it. But at least you can sleep tonight knowing that

[truncated]
