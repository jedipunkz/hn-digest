---
source: "https://www.tryai.dev/blog/ai-drawing-arena-colored-pencils-claude-gpt-grok"
hn_url: "https://news.ycombinator.com/item?id=48998404"
title: "\"Drawing\" the Mona Lisa with GPT-5.6, Claude, Gemini, and Grok"
article_title: "\"Drawing\" the Mona Lisa with GPT-5.6, Claude, Gemini, and Grok · TryAI"
author: "hershyb_"
captured_at: "2026-07-21T21:56:04Z"
capture_tool: "hn-digest"
hn_id: 48998404
score: 11
comments: 3
posted_at: "2026-07-21T21:13:20Z"
tags:
  - hacker-news
  - translated
---

# "Drawing" the Mona Lisa with GPT-5.6, Claude, Gemini, and Grok

- HN: [48998404](https://news.ycombinator.com/item?id=48998404)
- Source: [www.tryai.dev](https://www.tryai.dev/blog/ai-drawing-arena-colored-pencils-claude-gpt-grok)
- Score: 11
- Comments: 3
- Posted: 2026-07-21T21:13:20Z

## Translation

タイトル: GPT-5.6、クロード、ジェミニ、グロクでモナリザを「描く」
記事タイトル: GPT-5.6、Claude、Gemini、Grok でモナリザを「描く」 · TryAI
説明: 4 つのフロンティア モデル、空白のキャンバス、色鉛筆。彼らがモナリザを描こうとしたとき、私たちはすべてのストローク、金額、出力を追跡しました。

記事本文:
GPT-5.6、クロード、ジェミニ、およびグロクでモナリザを「描く」 · TryAI TryAI モデルの比較 ギャラリー ブログ サインイン 無料で試してみる すべての投稿 比較エージェント GPT-5.6、クロード、ジェミニ、およびグロクでモナリザを「描く」
4 つのフロンティア モデル、空白のキャンバス、色鉛筆。彼らがモナリザを描こうとしたとき、私たちはすべてのストローク、金額、出力を追跡しました。
11 min read 私たちは描画アリーナを構築しました。モデルに空白の白いキャンバスと色鉛筆ツールのセットを手渡し、その後は邪魔をしません。モデルは色、先端の幅、筆圧を設定し、ストロークのバッチを配置し、ブレンドするために汚れを付け、消去し、view_canvas を呼び出して自身の作業を確認し、何を修正するかを決定します。ターゲット イメージを再現するか、テキスト プロンプトから描画します。
4 つのビジョン モデル、 GPT-5.6 Sol 、Claude Fable 5 、Grok 4.5 、および Gemini 3.6 Flash を 2 つのターゲット (モナ リザとゴッホの星月夜、どちらも客観的にスコア付け) と 5 つのオープンエンド プロンプトで実行し、合計 28 枚の描画を行いました。ツールの使用、コスト、出力、モデルによって実際に作業が改善されたかどうかについて説明し、最後に私たちの意見を記載します。
私たちがこれを行う理由について簡単に説明します。なぜなら、私たちの最後の作品 (ミュージック ビデオを作成するモデル) が良い議論のきっかけとなり、モデルの創造性や芸術的能力を促進していると解釈する人もいるからです。これらはモデルの機能を客観的にテストするものではありません。これらは意図的に無制限で曖昧なタスクです。私たちが個人的にこれらを実行し続ける理由は次のとおりです。
それらは楽しくて、本当に有益です。モデルが緩やかで無制限のタスクに取り組むのを見ることは、別のベンチマーク数値よりも興味深い能力の視覚的な指標となります。
これらは、フロンティアと未開拓の本当のギャップを明らかにします。安価なオープンウェイト モデルは、多くの実行作業のフロンティアを完全に置き換えることができ、私たちはそれについて引き続きレポートしていきます。 B

世の中にはベンチマックスもたくさんありますが、このようなタスクはそれを打ち破ります。これらはフロンティア能力をその他の能力から完全に分離します。ご覧のとおり、Grok 4.5 はこの「基本的な」描画タスクが苦手で、私たちが試したオープンウェイト モデルは使用すらできず、いくつかは空白のキャンバスを返すだけでした。 (これは、完全にオープンソース化されたら報告する Kim K3 では変わる可能性があります)
これらは、長時間実行されるタスクの実際のコストを示します。 Claude Fable 5 はほとんどの場合、他のものよりもはるかに長い時間がかかり、はるかに多額の費用がかかり、結果はさらに悪化しました。 Fable は、ミュージック ビデオ チャレンジを含む多くのタスクで GPT-5.6 を上回りましたが、今回のタスクではオーバーヘッドが大幅に増加し、さらに悪かったです。ユースケースによっては、そのトレードオフが重要になります。
私たちは議論が大好きです。ディスカッションは本当に興味深いものでした。セットアップや新しいタスクについてご提案があれば、それを考慮に入れます。これらは実行するのがとても楽しく、出力を見るのは非常にエキサイティングです。
すべてのモデルはまったく同じ色鉛筆ツールセットで動作しました。
plan : ステップ間の思考と計画のための操作不要のスクラッチパッド。
view_target : ターゲット画像をもう一度見てください。
view_canvas : 現在のページをレンダリングして表示します。これは、ターゲットに対してドローを得点するときでもあります。
set_color / set_brush / set_pressure : 現在の鉛筆の色、ペン先の幅、筆圧 (不透明度 0 ～ 1、柔らかいトーンの場合は低い筆圧)。
draw : 1 回の呼び出しでマーク (ストローク、線、形状のアウトライン、ドット) のバッチを配置します。塗りつぶしはなく、本物の鉛筆のようにトーンと色を重ねて構築します。
smudge : ブレンド切り株のような長方形の領域をブレンド/ソフト化します。
Erase / Clear_canvas : 領域を白に戻す / ページ全体をリセットします。
ハーネス全体は github.com/hershalb/canvas-arena でオープンソースです。 P

ターゲット画像またはテキストプロンプトにそれを入力し、自分で実行します。
モデルは常にターゲットを認識でき、それとの構造的類似性 (SSIM) でスコア付けされました。以下にSSIMスコアを示します。
完全なトランスクリプト: GPT-5.6 Sol · Claude Fable 5 · Grok 4.5 · Gemini 3.6 Flash
完全なトランスクリプト: GPT-5.6 Sol · Claude Fable 5 · Grok 4.5 · Gemini 3.6 Flash
目標もスコアもありません。ただテキストの概要と空白のページがあるだけです。
「年老いた漁師の風化した顔、暖かい午後遅くの日差し、深いシワと無精ひげ」
完全なトランスクリプト: GPT-5.6 Sol · Claude Fable 5 · Grok 4.5 · Gemini 3.6 Flash
「穏やかな海、オレンジから紫の空、シルエットの地平線に沈む美しい夕日」
完全なトランスクリプト: GPT-5.6 Sol · Claude Fable 5 · Grok 4.5 · Gemini 3.6 Flash
「暗い背景にガラスの花瓶に入った一輪の赤いバラ、劇的なレンブラントの照明」
完全なトランスクリプト: GPT-5.6 Sol · Claude Fable 5 · Grok 4.5 · Gemini 3.6 Flash
「午後の日差しの中で窓辺で丸まって眠っているトラ猫」
完全なトランスクリプト: GPT-5.6 Sol · Claude Fable 5 · Grok 4.5 · Gemini 3.6 Flash
「暖炉が灯り、温かみのある琥珀色の光と柔らかな影を備えた居心地の良いキャビンのインテリア」
完全なトランスクリプト: GPT-5.6 Sol · Claude Fable 5 · Grok 4.5 · Gemini 3.6 Flash
モデルごとの 7 つの図面すべての平均と合計。
4 つのモデルはまったく同じツールセットをまったく異なる方法で使用しました。
Grok 4.5 ジェミニ 3.6 フラッシュ
GPT-5.6 Sol は set_color、set_brush、または set_pressure を一度も呼び出すことはなく、代わりに描画呼び出しごとにそれらをインラインで設定するため、その呼び出しはほぼすべてdraw、smudge、view_canvas になります。 Grok 4.5 はその逆を行いました。1,349 のツール呼び出しの 65% は set_color / set_brush / set_pressure でした。そのため、描画あたり平均 99 ステップでした。 Claude Fable 5 はスミッジ (123 コール) に依存し、常にレビューされました。 Gemini 3.6 Flash が最も多かった

何よりも執拗なレビュアーであり、その呼び出しの 3 分の 1 近くが view_canvas (図面ごとに約 23 のセルフレビュー) であり、Sol と同様に set_* ツールに触れることはありませんでした。
Claude Fable 5 の 7 つの描画の費用は推定 160 ドルで、GPT-5.6 Sol (7.74 ドル)、Grok 4.5 (9.21 ドル)、または Gemini 3.6 Flash (12.87 ドル) の約 20 倍に相当します。これは今となっては驚くべきことではありません。 Grok はこれまでで最も多くのトークン (3,400 万) を使用しましたが、~98% が入力レートの一部でキャッシュされた読み取りであり、Grok のレートが 4 つの中で最も低いため、低価格のままでした。また、Gemini はキャッシュされた読み取りに大きく依存していたので、2,770 万トークンであっても控えめなままでした。
実際にモデルは改善されましたか?
すべての view_canvas でターゲットに対してキャンバスのスコアが再設定されるため、実行中の進捗状況を監視できます。両方のターゲットにわたって 2 つのパターンが保持されます。
まず、早期に停滞します。 Claude Fable 5 はモナリザを 27 回レビューしましたが、その類似性は 5 回目くらいのレビューではほぼ横ばいになりました。第 2 に、8 つのターゲット実行すべてにおいて、最終的な描画スコアはモデルが実行中に達成した最高スコアを下回りました。 GPT-5.6 Sol のモナリザは、0.352 SSIM でピークに達し、0.325 で終了しました。 Starry Night は 0.179 でピークに達し、0.130 で終了しました。 Gemini 3.6 Flash は最も多くレビューされ (描画あたり約 23 回)、バッチ全体の最高値であるモナリザの 0.449 に達しましたが、最後には 0.337 まで下がりました。レビューを増やしても仕上がりが良くなるわけではありません。モデルは自身の最高のパフォーマンスを超えて編集を続けました。
客観的な類似性 (ターゲット実行)
SSIM は 0 ～ 1 で、高いほど目標に近づきます。 RMSEは0～255で、低いほど目標に近づきます。
生の SSIM では、Gemini 3.6 Flash は、最終スコアとピークの両方のターゲットで最高のスコアを獲得しましたが、最高のフレームと最終のフレームの間で最も振れも大きかったです。 Gemini 3.6 Flash はまともに見えましたが、明らかに tar に最も近いようには見えません

出力を取得します。しかし、生のスコアがより高いことが判明したのは非常に興味深いです。いつものように、SSIM は構造の近さを測定するものであり、人にとって図面がどれだけ美しく見えるかを測定するものではありません (以下の見解を参照)。
4 つのモデル、2 つのスコアリングされたターゲットとそれぞれ 5 つのプロンプト、合計 28 の描画。すべての実行で同じ色鉛筆ツールセットを使用しました。どのモデルも推理力を高めに設定した。
SSIM/RMSE は、両方の画像のサイズを 256x256 に変更し、比較することによって計算されます。これらは、芸術的な品質ではなく、構造/ピクセルの近さを測定し、ターゲットの実行にのみ適用されます (プロンプトにはスコアの参照がありません)。
実時間には、モデル自身の思考とツールの往復時間が含まれます。
GPT-5.6ソルが暴走リーダーとなった。その「星月夜」と「薔薇」は、このバッチ全体の中で私たちのお気に入りでした。空白のキャンバスに一筆一筆描いたことを考えると、これはワイルドです。また、ほぼすべての絵の細部において他の 2 人を上回っていますが、唯一の例外は年配の漁師です。
Grok 4.5 は基本的にゴミでした。使えるものはほとんど生成されませんでした。ツールを必死に使用したにもかかわらず、視覚的な出力を確実に理解したり判断したりできるとは私はまだ信じていません。
Gemini 3.6 Flash は、コストがわずかに増加しましたが、Grok よりも明らかに優れていましたが、Flash モデルをフロンティア バージョンと比較するのは少し不公平に思えます。 Gemini 4 がどのようなパフォーマンスを発揮するか非常に楽しみです。これは私たちもすでに知っていましたが、トークンの出力速度は非常に印象的です。
Claude Fable 5 は品質では 2 位でしたが、コストと時間では最下位でした。このタスクの場合、これは 3 つのタスクの中で最も魅力のないトレードオフでした。Sol に追いつかない出力のため、はるかに遅く、はるかに高価 (他のタスクの約 20 倍) になります。
ここで説明したすべてのモデルは、TryAI で 1 つのアカウントで利用でき、サブスクリプションなしで従量課金制で利用できます。
古い投稿 $100 AI ミュージック ビデオ: Claude Fable 5 vs. GPT-5.6 Sol

TryAI あらゆるフロンティア AI モデルを 1 か所にチャット、画像、ビデオ、音楽 - 使用した分だけお支払いください。
© 2026 TryAI.無断転載を禁じます。
すべての AI モデルを試してみたい人向けに構築されています。

## Original Extract

Four frontier models, a blank canvas, and colored pencils. We tracked every stroke, dollar, and output as they tried to draw the Mona Lisa.

"Drawing" the Mona Lisa with GPT-5.6, Claude, Gemini, and Grok · TryAI TryAI Models Compare Gallery Blog Sign in Try free All posts comparison agents drawing "Drawing" the Mona Lisa with GPT-5.6, Claude, Gemini, and Grok
Four frontier models, a blank canvas, and colored pencils. We tracked every stroke, dollar, and output as they tried to draw the Mona Lisa.
11 min read We built a drawing arena: hand a model a blank white canvas and a set of colored-pencil tools, then get out of the way. The model sets a color, tip width, and pressure, lays down batches of strokes, smudges to blend, erases, and calls view_canvas to see its own work and decide what to fix. It either reproduces a target image or draws from a text prompt.
We ran four vision models, GPT-5.6 Sol , Claude Fable 5 , Grok 4.5 , and Gemini 3.6 Flash , across two targets (the Mona Lisa and Van Gogh's Starry Night, both scored objectively) and five open-ended prompts, for 28 drawings total. We'll cover tool use, cost, output, and whether the models actually improved their work, with our opinion at the end.
A quick note on why we're doing this, because our last one of these (models making music videos) kicked off a good discussion where some folks read it as us promoting the creativity or artistic ability of models. These are not objective tests of model capability. They are deliberately open-ended, fuzzy tasks. Here is why we personally keep running them:
They are fun, and genuinely informative. Watching a model tackle a loose, open-ended task is a more interesting visual indicator of capability than yet another benchmark number.
They expose the real frontier-vs-open gap. Cheaper open-weight models can absolutely replace the frontier for a lot of execution work, and we will keep reporting on that. But there is also a lot of benchmaxxing out there, and tasks like this cut through it. They genuinely separate frontier capability from the rest. Grok 4.5, as you will see, was rough at this "basic" drawing task, and the open-weight models we tried were not even usable, several just returned a blank canvas. (That may change with Kimi K3 which we will report on once it is fully open-sourced)
They show what long-running tasks actually cost. Claude Fable 5 almost always took much longer than the others, for far more money, and here produced worse output. Fable has beaten GPT-5.6 for us on plenty of tasks, including the music-video challenge , but on this one it was worse for a lot more overhead. Depending on your use case, that trade-off matters.
We're loving the discussions. The discussions have genuinely been interesting, and if you have suggestions for the setup or new tasks, we will take them into account. These have been a blast to run, and the outputs are super exciting to see.
Every model worked with the exact same colored-pencil toolset:
plan : a no-op scratchpad for thinking and planning between steps.
view_target : look at the target image again.
view_canvas : render the current page and see it. This is also when we score the drawing against the target.
set_color / set_brush / set_pressure : the current pencil color, tip width, and pressure (0 to 1 opacity, low pressure for softer tone).
draw : lay down a batch of marks (strokes, lines, shape outlines, dots) in one call. There are no solid fills, tone and color are built by layering, like a real pencil.
smudge : blend/soften a rectangular region, like a blending stump.
erase / clear_canvas : lift a region back to white / reset the whole page.
The whole harness is open source at github.com/hershalb/canvas-arena . Point it at any target image or a text prompt and run it yourself.
The models could see the target the whole time and were scored on structural similarity (SSIM) to it. We show the SSIM scores below.
Full transcripts: GPT-5.6 Sol · Claude Fable 5 · Grok 4.5 · Gemini 3.6 Flash
Full transcripts: GPT-5.6 Sol · Claude Fable 5 · Grok 4.5 · Gemini 3.6 Flash
No target, no score, just a text brief and a blank page.
"An elderly fisherman's weathered face, warm late-afternoon sun, deep wrinkles and stubble"
Full transcripts: GPT-5.6 Sol · Claude Fable 5 · Grok 4.5 · Gemini 3.6 Flash
"A beautiful sunset over a calm ocean, orange-to-violet sky, silhouetted horizon"
Full transcripts: GPT-5.6 Sol · Claude Fable 5 · Grok 4.5 · Gemini 3.6 Flash
"A single red rose in a glass vase against a dark background, dramatic Rembrandt lighting"
Full transcripts: GPT-5.6 Sol · Claude Fable 5 · Grok 4.5 · Gemini 3.6 Flash
"A tabby cat curled asleep on a windowsill in afternoon sun"
Full transcripts: GPT-5.6 Sol · Claude Fable 5 · Grok 4.5 · Gemini 3.6 Flash
"A cozy cabin interior with a lit fireplace, warm amber light and soft shadows"
Full transcripts: GPT-5.6 Sol · Claude Fable 5 · Grok 4.5 · Gemini 3.6 Flash
Averages and totals across all seven drawings per model.
The four models used the exact same toolset in completely different ways.
Grok 4.5 Gemini 3.6 Flash
GPT-5.6 Sol never called set_color, set_brush, or set_pressure once , it set those inline on each draw call instead, so its calls are almost all draw, smudge, and view_canvas. Grok 4.5 did the opposite : 65% of its 1,349 tool calls were set_color / set_brush / set_pressure, which is why it averaged 99 steps per drawing. Claude Fable 5 leaned on smudge (123 calls) and reviewed constantly. Gemini 3.6 Flash was the most obsessive reviewer of all, nearly a third of its calls were view_canvas (about 23 self-reviews per drawing), and like Sol it never touched the set_* tools.
Claude Fable 5 cost an estimated $160 for its seven drawings, roughly 20x GPT-5.6 Sol ($7.74), Grok 4.5 ($9.21), or Gemini 3.6 Flash ($12.87). This shouldn't really be a surprise by now. Grok used by far the most tokens (34M) but stayed cheap because ~98% were cached reads at a fraction of the input rate, and Grok's rates are the lowest of the four. Gemini also leaned heavily on cached reads, so even at 27.7M tokens it stayed modest.
Did the models actually improve?
Every view_canvas re-scored the canvas against the target, so we can watch progress over a run. Two patterns held across both targets:
First, they plateau early. Claude Fable 5 reviewed its Mona Lisa 27 times, but its similarity was essentially flat after about the fifth review. Second, in all eight target runs, the final drawing scored below the best the model reached mid-run. GPT-5.6 Sol's Mona Lisa peaked at 0.352 SSIM and ended at 0.325; its Starry Night peaked at 0.179 and ended at 0.130. Gemini 3.6 Flash reviewed the most (about 23 times per drawing) and reached the highest peak of the whole batch, 0.449 on the Mona Lisa, then slid all the way back to 0.337 by the end. More reviewing did not translate into a better finish. The models kept editing past their own best performance.
Objective similarity (target runs)
SSIM is 0 to 1, higher is closer to the target. RMSE is 0 to 255, lower is closer to the target.
On raw SSIM, Gemini 3.6 Flash scored highest on both targets, on final score and on peak, though it also swung the most between its best frame and its final one. Gemini 3.6 Flash looked decent, but definitely does not look the closest to the target output. Very interesting that the raw score turned out to be higher though. As always, SSIM measures structural closeness, not how good the drawing looks to a person (see our take below).
Four models, two scored targets plus five prompts each, 28 drawings total. All runs used the same colored-pencil toolset. Reasoning was set to high for every model.
SSIM/RMSE are computed by resizing both images to 256x256 and comparing. They measure structural/pixel closeness, not artistic quality, and only apply to the target runs (prompts have no reference to score against).
Wall-clock time includes the model's own thinking and tool round-trips.
GPT-5.6 Sol was the runaway leader. Its Starry Night and its rose were our favorites of the whole batch, which is wild considering it drew them stroke by stroke on a blank canvas. It also beat the other two on detail in almost every drawing, the one exception being the elderly fisherman.
Grok 4.5 was basically garbage here. It rarely produced anything usable, and despite its frantic tool use I would not trust it to reliably understand or judge visual output yet.
Gemini 3.6 Flash was definitely better than Grok at a slight cost increase, but comparing the Flash model to frontier versions seems a little unfair. I'm definitely excited to see how Gemini 4 will perform. We also already knew this, but the token output speed is just so impressive.
Claude Fable 5 landed second on quality but at the far end on cost and time. For this task it was the least appealing trade-off of the three: much slower and far more expensive (about 20x the others) for output that did not keep up with Sol.
Every model mentioned here is available on TryAI with one account, pay-as-you-go, no subscription.
Older post $100 AI Music Video: Claude Fable 5 vs. GPT-5.6 Sol TryAI Every frontier AI model in one place. Chat, image, video, and music — pay only for what you use.
© 2026 TryAI. All rights reserved.
Built for people who want to try every AI model.
