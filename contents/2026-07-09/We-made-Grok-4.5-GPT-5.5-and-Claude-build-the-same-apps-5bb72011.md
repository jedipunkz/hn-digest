---
source: "https://www.tryai.dev/blog/grok-4.5-vs-gpt-5.5-vs-claude-build-off"
hn_url: "https://news.ycombinator.com/item?id=48838772"
title: "We made Grok 4.5, GPT-5.5, and Claude build the same apps"
article_title: "We made Grok 4.5, GPT-5.5, and Claude build the same apps · TryAI"
author: "hershyb_"
captured_at: "2026-07-09T00:00:40Z"
capture_tool: "hn-digest"
hn_id: 48838772
score: 2
comments: 0
posted_at: "2026-07-08T23:27:14Z"
tags:
  - hacker-news
  - translated
---

# We made Grok 4.5, GPT-5.5, and Claude build the same apps

- HN: [48838772](https://news.ycombinator.com/item?id=48838772)
- Source: [www.tryai.dev](https://www.tryai.dev/blog/grok-4.5-vs-gpt-5.5-vs-claude-build-off)
- Score: 2
- Comments: 0
- Posted: 2026-07-08T23:27:14Z

## Translation

タイトル: Grok 4.5、GPT-5.5、および Claude に同じアプリを構築させました
記事のタイトル: Grok 4.5、GPT-5.5、および Claude に同じアプリを構築させました · TryAI
説明: Grok 4.5 がリリースされました。そこで、GPT-5.5、Claude Opus 4.8、および Fable 5 で同じ対話型アプリをワンショットで実行し、遅延とコストを測定しました。勝ったのはこちらです。

記事本文:
Grok 4.5、GPT-5.5、および Claude で同じアプリを構築しました · TryAI TryAI モデルの比較 ギャラリー ブログ サインイン 無料で試してください すべての投稿の比較 grok コーディング Grok 4.5、GPT-5.5、および Claude で同じアプリを構築しました
Grok 4.5 がリリースされました。そこで、GPT-5.5、Claude Opus 4.8、および Fable 5 で同じ対話型アプリをワンショットで実行し、遅延とコストを測定しました。勝ったのはこちらです。
Grok 4.5 は今週リリースされ、xAI はこれをコーディングとエージェント作業のために「Cursor と一緒にトレーニングされた」これまでで最もスマートなモデルと呼んでいます。新しいコーディング モデルの雰囲気をチェックするインターネットで人気の方法は、ビルドオフです。プロンプトを 1 つ渡して、それが吐き出すアプリを確認し、クリップを投稿します。それで、私たちはそれをやりましたが、きちんとしました。
Grok 4.5 、 GPT-5.5 、 Claude Opus 4.8 、および Claude Fable 5 に、まったく同じ 3 つのプロンプトを与えました。それは、対話型アプリ用の単一の自己完結型 HTML ファイル (ライブラリなし、ネットワーク呼び出しなし) を構築するというものです。ショットは 1 つずつで、手を握ったり、プロンプトをいじったりすることはありません。次に、すべての結果を実際のブラウザにロードし、それを調べて、何が起こったかを記録しました。以下のすべてのクリップはモデルの生の出力であり、「ライブで再生」をクリックすると、実際に生成されたアプリを自分で実行できます。
失敗に関するルールが 1 つあります。アプリがまったくレンダリングされなかった場合、再試行は 1 回だけ許可され、それをいつ使用したかが通知されます。
ラウンド 1: スクランブルして解く 3D ルービック キューブ
「スクランブル」ボタンと「解決」ボタンを使って、カラフルで 3D のようなルービック キューブを組み立てましょう。立方体は回転を視覚的にアニメーション化する必要があります。
これが意地悪です。実際の 3D 計算、面ごとの状態、アニメーションをすべて 1 つのファイルに収める必要があります。
ライブで演奏: Grok 4.5 · GPT-5.5 · Opus 4.8 · Fable 5
クロードたちはこれを持って逃げた。 Opus 4.8 と Fable 5 は両方とも、アニメーション化されたターンでスクランブルおよび解決する適切な、正しく色の 3D 立方体を生成しました。まず試してください。 Grok 4.5 はここでつまづきました: 最初の再試行

タイトルとボタンは理解できましたが、立方体はまったくありませんでした (空白の空白)。そこで、許可された 1 回の再試行を費やし、2 回目の試行できれいでカラフルな 3D 立方体が生成されました。 GPT-5.5 は奇妙なもので、完全な立方体ではなく、ほとんど色を持たない単一の暗い顔だけをレンダリングしました。勝者：オーパスとフェイブル、同点。
ラウンド 2: パーティクル重力サンドボックス
キャンバス上のインタラクティブなパーティクル重力サンドボックス: トレイルを持つ数百のパーティクル。クリックすると重いアトラクターが追加されます。魅惑的なものにしてください。
ライブで演奏: Grok 4.5 · GPT-5.5 · Opus 4.8 · Fable 5
誰もが機能するサンドボックスをここに出荷したため、これは好みに応じて決まりました。 GPT-5.5 は、密集した渦巻く色の軌跡を持つ輝くネオン アトラクターという、最も真に魅惑的なものを作りました。 Grok 4.5 は、整然としたアトラクター リングとカラフルな縞模様のパーティクルを備え、クリーンかつ軌道上に到達しました。 Fable 5 は柔らかく光るオーブに傾き、Opus 4.8 は最も忙しく、最もクモの巣のような粒子フィールド、素晴らしい物理学を生み出しましたが、目の保養はわずかに減りました。勝者: GPT-5.5、バイブ付き。
ラウンド 3: プレイ可能なブレイクアウト ゲーム
キャンバス上でプレイ可能なブレイクアウト / レンガブレイカー: パドルがマウスを追いかけ、ボールがカラフルなレンガを壊し、スコアとライフを獲得します。
ライブで演奏: Grok 4.5 · GPT-5.5 · Opus 4.8 · Fable 5
デッドヒート。 4 人全員が、最初の試行で、スコア、ライフ、そして光るパドルを備えた、洗練された実際にプレイ可能なレンガブレイカーを作成しました。 Grok 4.5 と GPT-5.5 は、ネオン アーケードの外観に最も重点を置いています (GPT は、スクリプトがボールを打っている間にスコアを獲得しました)。厳密に言うと、4 つすべてが船級の品質です。勝者：全員です。
きれいなデモは別のことです。各モデルのランニングコストはいくらですか?私たちは、アプリが使用するものと同じプロバイダー パスを介して、独自のハーネス (コーディング、推論、要約にわたる 3 つの固定プロンプト、それぞれ 3 回の繰り返し、出力トークンの上限 400) を使用して自分たちで測定しました。

スループットは、すべてのプロバイダーに対して均一に測定された、合計実時間遅延に対する出力トークンです。
ここが Grok 4.5 の強みです。最初のトークンは 0.5 秒未満でヒットし、最大 110 トークン/秒 (ここでは他のすべての約 2 倍) でストリーミングされ、一連の応答の中で最も安価な応答であり、まさに xAI が行った「単位時間とコストあたりのインテリジェンス」のピッチでした。その壁時計の中央値は、冗長であり (回答ごとに最も多くのトークンを書き込んでいる)、尾部が尖っている (~9 秒 p95) ため、集団の真ん中にしか見えません。 GPT-5.5 は短答で最も機敏で、Opus 4.8 はバランスの取れた中間に位置し、Fable 5 は最も遅くて最も高価で、知能チャートのトップに支払う税金でした。
ボーナスラウンド: 何か奇妙なものを描きます
コーディングは 1 つの軸です。空間的想像力はまた別です。私たちは各モデルに、意図的に愚かなシーン、つまり月面を歩く宇宙飛行士におんぶに乗って乗馬するシーンを描いた 1 つの手書きの SVG (ラスターやライブラリなし) を作成するよう依頼しました。役割逆転、2 つのフィギュア、1 つのファイル。
生の SVG: Grok 4.5 · GPT-5.5 · Opus 4.8 · Fable 5
フェイブル 5 が注目を集めました。カウボーイハットをかぶった馬が「めまいがする、人間！」と叫びます。一方で、前かがみになった宇宙飛行士は「フフ…フフ…」と喘ぎ声を上げます。それは単なる絵を描くことではなく、コメディです。 GPT-5.5は僅差で2位となり、馬は「ウィー！」と大喜びした。 Grok 4.5 は、カラフルで読みやすいシーンで概要をきれいに完成させました。 Opus 4.8 も特徴的な馬に乗った宇宙飛行士を描いていましたが、生の SVG には厳密な SVG パーサーをトリップさせる重複属性が同梱されており (上の画像では寛大にレンダリングしました)、小さいながらも正確性には問題がありました。
Grok 4.5 は、スピードと価値を兼ね備えたモンスターです。優れたブレイクアウトと滑らかな重力シミュレーションを構築し、フィールドの約 2 倍の速度でストリーミングし、実行コストは最も安価です。唯一の本当の失敗は、最も困難なステートフル タスク (最初の試行では空のキューブ、再試行で修正される) でした。もしあなたの

ワークロードはレイテンシーとコストが複合する大量のコード生成であり、これについて議論するのは非常に困難です。 Grok 4.5 と GPT-5.5 を参照してください。
クロード オーパス 4.8 とフェイブル 5 は最も信頼できるビルダーです。最初に 3D キューブを完成させたのはこの 2 人だけで、Fable は最も面白い SVG を持っていました。特に Fable では、レイテンシと価格で対価を支払います。 Grok 4.5 と Opus 4.8 を参照してください。
GPT-5.5はキビキビとしたスタイリストです。最も美しい重力シミュレーション、最速の短い回答ですが、立方体は失敗しました。
正直な見出し: Grok 4.5 は、真新しい発売日に、購入できる最高のモデルと互角に渡り合い、スピードとコストで完全に勝利を収めました。独自のプロンプトで解決したいですか?ここにあるすべてのモデルは 1 つの TryAI アカウントで実行され、従量課金制です。モデルを参照して、独自の構築を開始してください。
ここで説明したすべてのモデルは、TryAI で 1 つのアカウントで利用でき、サブスクリプションなしで従量課金制で利用できます。
古い投稿 Meta の Muse 画像と実際に使用できる最高の画像モデル TryAI のすべてのフロンティア AI モデルを 1 か所にまとめました。チャット、画像、ビデオ、音楽 - 使用した分だけお支払いください。
© 2026 TryAI.無断転載を禁じます。
すべての AI モデルを試してみたい人向けに構築されています。

## Original Extract

Grok 4.5 just launched. So we had it, GPT-5.5, Claude Opus 4.8, and Fable 5 one-shot the same interactive apps, then measured latency and cost. Here is who won.

We made Grok 4.5, GPT-5.5, and Claude build the same apps · TryAI TryAI Models Compare Gallery Blog Sign in Try free All posts comparison grok coding We made Grok 4.5, GPT-5.5, and Claude build the same apps
Grok 4.5 just launched. So we had it, GPT-5.5, Claude Opus 4.8, and Fable 5 one-shot the same interactive apps, then measured latency and cost. Here is who won.
7 min read Grok 4.5 landed this week with xAI calling it their smartest model yet, "trained alongside Cursor" for coding and agentic work. The internet's favorite way to vibe-check a new coding model is the build-off: hand it one prompt, see what app it spits out, post the clip. So we did it, but properly.
We gave Grok 4.5 , GPT-5.5 , Claude Opus 4.8 , and Claude Fable 5 the exact same three prompts: build a single self-contained HTML file (no libraries, no network calls) for an interactive app. One shot each, no hand-holding, no prompt-fiddling. Then we loaded every result in a real browser, poked at it, and recorded what happened. Every clip below is the model's raw output, and you can click "play it live" to run the actual generated app yourself.
One rule on failures: if an app did not render at all, we allowed exactly one retry, and we tell you when we used it.
Round 1: a 3D Rubik's Cube that scrambles and solves
Build a colorful, 3D-looking Rubik's Cube with 'Scramble' and 'Solve' buttons. The cube must visibly animate its rotations.
This is the mean one. It needs real 3D math, per-face state, and animation, all in one file.
Play them live: Grok 4.5 · GPT-5.5 · Opus 4.8 · Fable 5
The Claudes ran away with this one. Opus 4.8 and Fable 5 both produced a proper, correctly colored 3D cube that scrambles and solves with animated turns, first try. Grok 4.5 stumbled here: its first attempt rendered the title and buttons but no cube at all (a blank void), so we spent our one allowed retry, and the second attempt produced a clean, colorful 3D cube. GPT-5.5 was the odd one out, rendering only a single dark face with barely any color instead of a full cube. Winner: Opus and Fable, tied.
Round 2: a particle gravity sandbox
An interactive particle gravity sandbox on a canvas: hundreds of particles with trails, clicking adds a heavy attractor. Make it mesmerizing.
Play them live: Grok 4.5 · GPT-5.5 · Opus 4.8 · Fable 5
Everyone shipped a working sandbox here, so this came down to taste. GPT-5.5 made the most genuinely mesmerizing one: glowing neon attractors with dense, swirling colored trails. Grok 4.5 went clean and orbital, with tidy attractor rings and colorful streaking particles. Fable 5 leaned into soft glowing orbs, and Opus 4.8 produced the busiest, most web-like particle field, great physics, slightly less eye-candy. Winner: GPT-5.5, on vibes.
Round 3: a playable Breakout game
A playable Breakout / brick-breaker on a canvas: paddle follows the mouse, ball breaks colorful bricks, with score and lives.
Play them live: Grok 4.5 · GPT-5.5 · Opus 4.8 · Fable 5
Dead heat. All four produced a polished, actually-playable brick-breaker with score, lives, and a glowing paddle, on the first try. Grok 4.5 and GPT-5.5 leaned hardest into the neon arcade look (GPT even racked up a score while our script batted the ball around). If we are splitting hairs, all four are ship-quality. Winner: everybody.
Pretty demos are one thing; what does each model cost to run? We measured it ourselves with our own harness (three fixed prompts spanning coding, reasoning, and summarization, three reps each, capped at 400 output tokens) through the same provider path the app uses. Throughput is output tokens over total wall-clock latency, measured uniformly for every provider.
This is where Grok 4.5 flexes. It hit first token in under half a second, streamed at ~110 tokens/second (roughly double everything else here), and was the cheapest reply of the bunch, exactly the "intelligence per unit of time and cost" pitch xAI made. Its median wall-clock only looks mid-pack because it is verbose (it wrote the most tokens per answer), and its tail was spiky (a ~9s p95). GPT-5.5 was snappiest on short answers, Opus 4.8 sat in the balanced middle, and Fable 5 was the slowest and priciest, the tax you pay for the top of the intelligence charts.
Bonus round: draw something weird
Coding is one axis; spatial imagination is another. We asked each model for a single hand-authored SVG (no raster, no libraries) of a deliberately silly scene: a horse riding piggyback on an astronaut walking on the moon. Role reversal, two figures, one file.
Raw SVGs: Grok 4.5 · GPT-5.5 · Opus 4.8 · Fable 5
Fable 5 stole the show: a cowboy-hatted horse yelling "Giddy-up, human!" while a hunched astronaut wheezes "hff... hff...". That is not just drawing, that is comedy. GPT-5.5 was a close second with a gleeful horse shouting "WHEE!". Grok 4.5 nailed the brief cleanly with a colorful, readable scene. Opus 4.8 drew a characterful horse-on-astronaut too, but its raw SVG shipped with a duplicate attribute that trips strict SVG parsers (we rendered it leniently for the image above), a small but real correctness ding.
Grok 4.5 is the speed-and-value monster. It built a great Breakout and a slick gravity sim, streams about twice as fast as the field, and is the cheapest to run. Its only real miss was the hardest stateful task (a blank cube on the first try, fixed on the retry). If your workload is high-volume codegen where latency and cost compound, it is very hard to argue with. See Grok 4.5 vs GPT-5.5 .
Claude Opus 4.8 and Fable 5 are the most reliable builders. They were the only two to nail the 3D cube first try, and Fable had the funniest SVG. You pay for it in latency and price, especially Fable. See Grok 4.5 vs Opus 4.8 .
GPT-5.5 is the snappy stylist. Prettiest gravity sim, fastest short answers, but it flubbed the cube.
The honest headline: on a brand-new launch day, Grok 4.5 went toe to toe with the best models you can buy and won on speed and cost outright. Want to settle it with your own prompts? Every model here runs on one TryAI account , pay-as-you-go. Browse the models and start your own build-off.
Every model mentioned here is available on TryAI with one account, pay-as-you-go, no subscription.
Older post Meta's Muse Image vs the best image models you can actually use TryAI Every frontier AI model in one place. Chat, image, video, and music — pay only for what you use.
© 2026 TryAI. All rights reserved.
Built for people who want to try every AI model.
