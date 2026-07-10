---
source: "https://www.tryai.dev/blog/gpt-5.6-build-off-12-models"
hn_url: "https://news.ycombinator.com/item?id=48865093"
title: "GPT-5.6, Grok 4.5, Claude, and Muse Spark build the same 4 apps"
article_title: "GPT-5.6, Grok 4.5, Claude, and Muse Spark build the same 4 apps · TryAI"
author: "hershyb_"
captured_at: "2026-07-10T21:52:28Z"
capture_tool: "hn-digest"
hn_id: 48865093
score: 62
comments: 28
posted_at: "2026-07-10T20:52:28Z"
tags:
  - hacker-news
  - translated
---

# GPT-5.6, Grok 4.5, Claude, and Muse Spark build the same 4 apps

- HN: [48865093](https://news.ycombinator.com/item?id=48865093)
- Source: [www.tryai.dev](https://www.tryai.dev/blog/gpt-5.6-build-off-12-models)
- Score: 62
- Comments: 28
- Posted: 2026-07-10T20:52:28Z

## Translation

タイトル: GPT-5.6、Grok 4.5、Claude、および Muse Spark が同じ 4 つのアプリを構築
記事のタイトル: GPT-5.6、Grok 4.5、Claude、および Muse Spark が同じ 4 つのアプリを構築 · TryAI
説明: GPT-5.6 の新しい Sol、Terra、Luna ティアは、Grok 4.5、Claude、Meta の Muse Spark、およびレイキャスター、ルービック キューブ、電卓、およびライフ ゲームを使用した無差別級チームと対戦します。ここでは、コストとレイテンシを含むすべてのビルドを示します。

記事本文:
GPT-5.6、Grok 4.5、Claude、および Muse Spark は同じ 4 つのアプリをビルドします · TryAI TryAI モデルの比較 ギャラリー ブログ サインイン 無料で試してください すべての投稿の比較コーディング ベンチマーク GPT-5.6、Grok 4.5、Claude、および Muse Spark は同じ 4 つのアプリをビルドします
GPT-5.6 の新しい Sol、Terra、および Luna ティアは、Grok 4.5、Claude、Meta の Muse Spark、およびレイキャスター、ルービック キューブ、電卓、およびゲーム オブ ライフを使用した無差別級チームと対戦します。ここでは、コストとレイテンシを含むすべてのビルドを示します。
20 分で読む 前回のビルドオフは Hacker News のトップページに掲載され、コメントは留まりませんでした。当然のことですが、その多くは良いフィードバックでした。そこで私たちはそれを採用し、GPT-5.6 を 3 つの層 (Sol、Terra、Luna) に導入し、Meta がコーディング モデル (Muse Spark 1.1) をサプライズドロップして、全体をさらに大きく再実行しました。つまり、12 のモデル、4 つのアプリ、それぞれ 5 回の試行でした。
皆様のフィードバックに基づいて変更した点:
オープンウェイトモデルをミックスに加えたいと考えていました。そこで、比較ポイントとして GLM-5.2 、 Qwen 3.7 Plus 、 DeepSeek V4 Pro 、および Kim K2.6 を追加しました。これらはすべて Fireworks 経由で提供されます。
ある試みは奇妙だった、とあなたは言いました。同意しました。すべてのモデルはタスクごとに 5 回の試行が行われるようになりました。上部では、モデルごとに 1 つのサンプル実行が得られます。各タスク テーブルには、5 つのうち実際に成功したと思われる数 (および数え方) が示され、最も気に入った試みがリンクされます。すべての試行は下部にリンクされているため、これらのモデルが実行ごとにどれだけスイングするかを確認できます。
「これは客観的ではありません。」その通りです、そして私たちはそのように装っているわけではありません。私たちは科学的な判決を下しているわけではありません。私たちは大量の成果物を生成し、それらをすべて公開しています。あなたは自分の意見を形成することができます。以下のすべては、結果を見て私たちが観察したものにすぎません。
そのままスキップして、生のビルドを確認したいですか?すべての試行にジャンプして、自分で実行します。
ラインナップ、12ストロ

ng: 新しい GPT-5.6 Sol 、 GPT-5.6 Terra 、および GPT-5.6 Luna ;メタのミューズスパーク 1.1 ; Grok 4.5 、 GPT-5.5 、 Claude Opus 4.8 、および Claude Fable 5 。加えて、オープンウェイトクルー: Qwen 3.7 Plus 、 DeepSeek V4 Pro 、 Kimi K2.6 、および GLM-5.2 。
各タスクは次のとおりです。上部にある 5 つの試行のうち私たちが選んだもの、そのすぐ下に 5 つすべてのコストと時間が表示されます。また、投稿の下部にはすべての生の試行へのリンクがあるので、自分で判断できます。
タスク 1: Doom スタイルのレイキャスター迷路
WASD で歩く一人称視点のレイキャスター、奥行きのある陰影のある壁、床と天井、衝突。
「プレイアブル」のカウント方法: 私たちが気にした唯一の質問は、実際に迷宮を歩き、移動し、方向転換できるかどうかでした。 「はい」の場合、それはカウントされます。
全体として、クロードはここで私たちが予想していたよりも悪かった。 GPT は他のどのモデルよりも優れたパフォーマンスを発揮し、Grok はその価格帯で本当に使える代替品であり、Muse Spark は実際に機能した走行で本当に驚きました。
タスク 2: 3D ルービック キューブ (スクランブル + ソルブ)
回転を視覚的にアニメーション化するスクランブル ボタンとソルブ ボタンを備えた、カラフルで 3D のようなルービック キューブを構築します。
「クリーンソルブ」のカウント方法: 立方体をスクランブルして解決し、両方のアニメーションがスムーズに実行され、不具合や色の変化がない場合にのみ試行をカウントしました。
全体として、レイキャスターで 3D が明らかにリードしていることを考えると、GPT がここでパフォーマンスを下回ったことには驚きました。クロードは再び素晴らしい仕事をしましたが、ファブルは 5 対 5 のクリーンな成績を収めましたが、オーパスは奇妙なことに、完璧な解決策を 1 つも決めることができませんでした。
任意のモデルの 5 つすべてをライブでプレイします: Grok、Opus、Qwen (他の試行では番号 1 ～ 5 を交換してください)。
数字、演算子、クリア、等しい、正しい演算子の優先順位、実際の電卓の外観。
「稼働」のカウント方法: 網羅的なものはなく、(((5 × 5) − 100) / 10) のような基本的な計算だけで、それぞれがどのように注文を処理したかを確認しました。

操作と結果のレンダリング。
全体として、これは明らかにクロードの最高の作品でした。オーパスとフェイブルの両方が 5 つすべてを釘付けにし、スタイルの点でフェイブルのものが私たちのお気に入りでした。 GPT-5.6 Sol は、スタイルをやりすぎて、Fable と同様に電卓を 3D でレンダリングしようとします。ただし、スタイルを確立することはできず、全体的なエクスペリエンスが低下します。エクスペリエンスがそのまま機能したというだけの理由で、より単純な GPT モデルのほうがより適切に機能するように見えました。 GLM-5.2 も、理屈は外れましたが、9 分間の失敗から立ち直り、わずか 1 セントの動作で動作するようになりました。
グリッド キャンバス、再生/一時停止/ステップ/ランダム化/クリア、クリックしてセルを切り替え、アニメーション化された世代。
Life では別の 5 回の試行スコアリング パスを実行しませんでした。そのため、ここでは単にコストと時間に加えて、以下の一般的な印象を示します。
Grok 4.5 はここではうまくいきましたが、より大きなポイントは、このタスクが OSS モデルにとって非常にうまくいくのに十分単純だったということです。おそらく、Game of Life 用のオープンなサンプル コードが十分に存在するため、より低コストでより優れた作業を行うことができます。 Qwen 3.7 Plus と GLM-5.2 は、このような用途には明らかに頼りになりますが、一般的にはそれらに依存しません。他のタスクを見ると、真に斬新な作業やより複雑な作業にまだ苦労していることがわかります。
領収書: 生の速度とコスト (短い回答)
別の質問、別の表。これは標準のレイテンシ ハーネス (3 つの短いプロンプト、5 回の繰り返し、400 トークンの上限) であり、ビルド タスクではありません。 tok/s は、実時間での出力トークンであり、すべてに均一です。
正直な注意点が 1 つあります。いくつかのオープンウェイトは応答全体をバーストでバッファリングし、400 トークンの上限に達したため、tok/s は上限であり、実際のデコード レートではありません。明確に読み取れます: GPT-5.6 層は、短いプロンプト (Luna は約 1 秒で応答します) で最も機敏なモデルであり、Qwen は異常なほど安くて高速であり、DeepSeek と GLM は

throwpokes は、アプリを生成する際の苦痛と一致しました。
ボーナス: 宇宙飛行士に乗った馬を描く (SVG、ベスト 5)
ワンショット SVG、ライブラリなし。各モデルのベスト オブ 5 を示します (厳密に有効な SVG、次に最も詳細な SVG を優先します)。
キミ K2.6 クロード・フェイブル 5 DeepSeek V4 プロ GLM-5.2
GPT-5.6 ソル GPT-5.6 テラ GPT-5.6 ルナミューズ スパーク 1.1
個人的には、Claude Fable は SVG レンダリングで素晴らしい仕事をしています。ほとんどの場合、面白くて、質の高い結果が得られました。 GPT-5.6 モデルは、馬や宇宙飛行士のきれいなレンダリングが含まれていなかったため、ここでは驚くほど精彩を欠いていました。 Grok 4.5 はここでもかなりうまくいきました。
ボーナス 2: イーロンとベゾスがブルー オリジンの着陸を観察 (SVG、ベスト オブ 5)
リクエストに応じて、さらにハードなシーンを追加しました。ブルー オリジンのブースターが外洋のパッドに着陸するのを見ている、2 人の有名なハイテク億万長者の風刺画です。これは構図と似たものに重点を置いています。再び、Claude Fable がフィールドを席巻します。クリーンなレンダリング、ベゾスの額の光沢のあるスポット、着陸パッドの周囲の煙など、細部まで再現されています。 GPT は再び非常に漫画的な結果を生成します。それが好みの場合は、生成されるワンショットの結果を改善してみてもよいでしょう。世代には常にいくつかの小さなエラーがあるためです。 OSS モデルの GLM-5.2 と Qwen 3.7 も、このタスクでは非常にうまく機能しました。
キミ K2.6 クロード・フェイブル 5 DeepSeek V4 プロ GLM-5.2
GPT-5.6 ソル GPT-5.6 テラ GPT-5.6 ルナミューズ スパーク 1.1
TL;DR
フロンティアは依然として困難な課題に勝つが、複雑な課題では特に及ばない。 GPT-5.6 ソルとクロード・フェイブル 5 が傑出した成績でした。ソルはレイキャスターで優れ、フェイブルはルービック キューブで優れていました。他のすべてはうまくいきましたが、最高の結果はまだフロンティアにあります。
SOTAとオープンウェイトの間には、真に新規性や完全性に関して明らかなギャップがあります。

x は動作し、raycaster と cube で確認できます。しかし、ライフ ゲームのような単純でよく踏まれているタスクでは、OSS モデルが独自の役割を果たしており、Qwen 3.7 と GLM-5.2 が数分の 1 のコストでそれを実現できる十分なサンプル コードが存在します。そのクラスの問題に使用してください。他のタスクを見ると、彼らは依然として難しいことに取り組んでいることが示されているため、一般的にはそれらに頼らないでください。
Grok 4.5 は、一部のタスクに関しては正真正銘の「Opus レベル」です。ビジネスのコストを気にする場合、私はためらうことなく二次実行モデルとしてそれを利用するでしょう。 Grok 4.5 対 Opus 4.8 および Grok 4.5 対 GPT-5.5 を参照してください。
Muse Spark 1.1 には嬉しい驚きがありました。私はあまり期待していなかったので、Grok 4.5 よりも一歩下に感じましたが、一般的にオープンウェイトよりも優れていました。本当のデビューですが、まだ手が届くようなものではありません。
この結論は、発売日以降も当てはまります。最新で最も高価なフラッグシップ製品が自動的に勝者になるわけではありません。これらのプロンプトを自分で実行したいですか?ここにあるすべてのモデルは 1 つの TryAI アカウント上にあり、従量課金制です。モデルを参照して、独自の構築を開始してください。
私たちが生成した生のビルドはすべて、12 のモデルにタスクごとに 5 回の試行を掛けたものになります。任意の番号をクリックしてその正確な試みを開き、自分でそれを突いてみましょう。試行間のスイングがポイントです。
ここで説明したすべてのモデルは、TryAI で 1 つのアカウントで利用でき、サブスクリプションなしで従量課金制で利用できます。
以前の投稿 Grok 4.5、GPT-5.5、および Claude で同じアプリ TryAI Every フロンティア AI モデルを 1 か所で構築できるようにしました。チャット、画像、ビデオ、音楽 - 使用した分だけお支払いください。
© 2026 TryAI.無断転載を禁じます。
すべての AI モデルを試してみたい人向けに構築されています。

## Original Extract

GPT-5.6's new Sol, Terra, and Luna tiers go head-to-head with Grok 4.5, Claude, Meta's Muse Spark, and the open-weights crew on a raycaster, a Rubik's cube, a calculator, and Game of Life. Here's every build, with cost and latency.

GPT-5.6, Grok 4.5, Claude, and Muse Spark build the same 4 apps · TryAI TryAI Models Compare Gallery Blog Sign in Try free All posts comparison coding benchmarks GPT-5.6, Grok 4.5, Claude, and Muse Spark build the same 4 apps
GPT-5.6's new Sol, Terra, and Luna tiers go head-to-head with Grok 4.5, Claude, Meta's Muse Spark, and the open-weights crew on a raycaster, a Rubik's cube, a calculator, and Game of Life. Here's every build, with cost and latency.
20 min read Our last build-off hit the Hacker News front page , and the comments did not hold back. Fair enough, a lot of it was good feedback. So we took it, and with GPT-5.6 landing in three tiers (Sol, Terra, Luna) and Meta surprise-dropping a coding model (Muse Spark 1.1), we ran the whole thing again, bigger: twelve models, four apps, five attempts each.
What we changed based on your feedback:
You wanted open-weights models in the mix. So we added GLM-5.2 , Qwen 3.7 Plus , DeepSeek V4 Pro , and Kimi K2.6 as comparison points, all served via Fireworks .
One attempt was weird, you said. Agreed. Every model now gets five attempts per task. Up top you get one sample run per model; each task table then says how many of the five we thought actually succeeded (and how we counted) and links the attempt we liked best; and every attempt is linked at the bottom so you can see how much these models swing run to run.
"This isn't objective." Correct, and we are not pretending it is. We are not handing down a scientific verdict. We generated a big pile of artifacts, we are publishing all of them, and you can form your own opinion. Everything below is just our observations from watching the results.
Want to skip straight to poking at the raw builds? Jump to every attempt and run them yourself.
The lineup, twelve strong: the new GPT-5.6 Sol , GPT-5.6 Terra , and GPT-5.6 Luna ; Meta's Muse Spark 1.1 ; Grok 4.5 , GPT-5.5 , Claude Opus 4.8 , and Claude Fable 5 ; plus the open-weights crew: Qwen 3.7 Plus , DeepSeek V4 Pro , Kimi K2.6 , and GLM-5.2 .
Here is each task: our pick of the five attempts playing up top, the cost and time for all five just below, and links to every raw attempt at the bottom of the post so you can judge for yourself.
Task 1: Doom-style raycaster maze
First-person raycaster you walk with WASD, shaded walls with depth, floor and ceiling, collision.
How we counted "playable": the only question we cared about was whether you could actually walk through the labyrinth, move and turn. If yes, it counted.
Overall, Claude did worse than we expected here. GPT outperformed every other model, Grok was a genuinely usable alternative at its price point, and Muse Spark was a real surprise on the runs that actually worked.
Task 2: 3D Rubik's Cube (scramble + solve)
Build a colorful, 3D-looking Rubik's Cube with Scramble and Solve buttons that visibly animate the rotations.
How we counted a "clean solve": we scrambled and solved the cube and only counted an attempt if both animations ran smooth, no glitches, no color changes.
Overall, we were surprised GPT underperformed here given its clear 3D lead in the raycaster. Claude did an amazing job again, though it was Fable carrying it with a clean five-for-five, while Opus, oddly, couldn't land a single flawless solve.
Play all five of any model live: Grok · Opus · Qwen (swap the number 1-5 for other attempts).
Digits, operators, clear, equals, correct operator precedence, real calculator look.
How we counted "working": nothing exhaustive, just basic calculations like (((5 × 5) − 100) / 10) to see how each one handled order of operations and rendered the result.
Overall, this was clearly Claude's best work: both Opus and Fable nailed all five, and Fable's was our favorite on style. GPT-5.6 Sol tries to go overboard with styles and render the calculator in 3D similar to Fable. However, it doesn't nail the styles, which leads to a worse overall experience. The simpler GPT models seemed to do a better job just because the experience worked out of the box. GLM-5.2, reasoning-off, also came back from nine-minute failures to a snappy, fraction-of-a-cent build.
Grid canvas, Play/Pause/Step/Randomize/Clear, click to toggle cells, animated generations.
We did not run a separate five-attempt scoring pass on Life, so here it is just cost and time plus the general impression below.
Grok 4.5 did well here, but the bigger takeaway is that this task was simple enough for the OSS models to do extremely well on. There is probably enough open example code for Game of Life out there that they were able to do a much better job at a way lower cost. Qwen 3.7 Plus and GLM-5.2 are clearly the go-to for something like this, but I would not rely on them generally, the other tasks show they still struggle with genuinely novel or more complex work.
The receipts: raw speed and cost (short answers)
Separate question, separate table. This is our standard latency harness (three short prompts, five reps, 400-token cap), not the build tasks. tok/s is output tokens over wall-clock, uniform for all.
One honest caveat: several open-weights buffered their whole reply in a burst and hit the 400-token cap, so their tok/s is a ceiling, not a true decode rate. The clear read: the GPT-5.6 tiers are the snappiest models here on short prompts (Luna answers in about a second), Qwen is absurdly cheap and fast, and DeepSeek and GLM are the slowpokes, which matched the agony of generating their apps.
Bonus: draw a horse riding an astronaut (SVG, best of 5)
One-shot SVG, no libraries. We show each model's best-of-five (we prefer a strictly-valid SVG, then the most detailed).
Kimi K2.6 Claude Fable 5 DeepSeek V4 Pro GLM-5.2
GPT-5.6 Sol GPT-5.6 Terra GPT-5.6 Luna Muse Spark 1.1
Personally, Claude Fable does a great job with the SVG rendering. It's funny most of the time as well and came up with good quality results. The GPT-5.6 models were surprisingly lackluster here since none of them included clean renderings of the horse or the astronaut. Grok 4.5 did pretty well here as well.
Bonus 2: Elon and Bezos watch a Blue Origin landing (SVG, best of 5)
A harder scene by request: two recognizable tech-billionaire caricatures watching a Blue Origin booster land on a pad in the open ocean. This one leans on composition and likeness. Again, Claude Fable sweeps the field: it produces great detail with a clean render, a shiny spot on Bezos' forehead, and smoke around the landing pad. GPT again produces pretty cartoony results; if that's something you prefer then maybe you can try improving the one-shot results it produces, since there are always some minor errors with the generations. The OSS models, GLM-5.2 and Qwen 3.7, also did really well for this task.
Kimi K2.6 Claude Fable 5 DeepSeek V4 Pro GLM-5.2
GPT-5.6 Sol GPT-5.6 Terra GPT-5.6 Luna Muse Spark 1.1
TL;DR
The frontier still wins the hard tasks, and it is not particularly close on the complex ones. GPT-5.6 Sol and Claude Fable 5 were the standouts: Sol excelled at the raycaster, Fable at the Rubik's cube. Everything else did decently, but the best results are still at the frontier.
There is a clear gap between SOTA and open-weights on genuinely novel or complex work, and you can see it in the raycaster and cube. But on a simple, well-trodden task like Game of Life, the OSS models hold their own, there is enough example code out there that Qwen 3.7 and GLM-5.2 nail it at a fraction of the cost. Use them for that class of problem; just do not lean on them generally, because the other tasks show they still struggle with the hard stuff.
Grok 4.5 is genuinely "Opus level" on some tasks. If I cared about cost for my business, I would reach for it as a secondary execution model without hesitation. See Grok 4.5 vs Opus 4.8 and Grok 4.5 vs GPT-5.5 .
Muse Spark 1.1 pleasantly surprised me. I was not really expecting much, and it felt a step below Grok 4.5 but generally better than the open-weights. It is a real debut, though not something I would reach for just yet.
The takeaway holds even after launch day: the newest, most expensive flagship is not the automatic winner. Want to run these prompts yourself? Every model here is on one TryAI account , pay-as-you-go. Browse the models and start your own build-off.
Every raw build we generated, twelve models times five attempts per task. Click any number to open that exact attempt and poke at it yourself; the swings between attempts are the point.
Every model mentioned here is available on TryAI with one account, pay-as-you-go, no subscription.
Older post We made Grok 4.5, GPT-5.5, and Claude build the same apps TryAI Every frontier AI model in one place. Chat, image, video, and music — pay only for what you use.
© 2026 TryAI. All rights reserved.
Built for people who want to try every AI model.
