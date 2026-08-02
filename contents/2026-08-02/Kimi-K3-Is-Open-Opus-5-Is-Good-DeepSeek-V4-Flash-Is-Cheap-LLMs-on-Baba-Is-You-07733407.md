---
source: "https://quesma.com/blog/baba-kimi-k3-opus-5/"
hn_url: "https://news.ycombinator.com/item?id=49146588"
title: "Kimi K3 Is Open, Opus 5 Is Good, DeepSeek V4 Flash Is Cheap: LLMs on Baba Is You"
article_title: "Kimi K3 is Open, Opus 5 is Good, DeepSeek V4 Flash is Cheap: LLMs on Baba Is You - Quesma Blog"
author: "stared"
captured_at: "2026-08-02T17:53:35Z"
capture_tool: "hn-digest"
hn_id: 49146588
score: 1
comments: 1
posted_at: "2026-08-02T17:43:54Z"
tags:
  - hacker-news
  - translated
---

# Kimi K3 Is Open, Opus 5 Is Good, DeepSeek V4 Flash Is Cheap: LLMs on Baba Is You

- HN: [49146588](https://news.ycombinator.com/item?id=49146588)
- Source: [quesma.com](https://quesma.com/blog/baba-kimi-k3-opus-5/)
- Score: 1
- Comments: 1
- Posted: 2026-08-02T17:43:54Z

## Translation

タイトル: キミ K3 はオープン、オーパス 5 は良好、ディープシーク V4 フラッシュは安い: Baba の LLM はあなたです
記事タイトル: Kim K3 はオープン、Opus 5 は良好、DeepSeek V4 Flash は安い: LLMs on Baba Is You - Quesma Blog
説明: パズル ゲーム Baba Is You に基づいた LLM エージェント ベンチマークである Baba Is Bench で、2026 年 7 月の新規リリース Kimi K3、Claude Opus 5、Grok 4.5、Gemini 3.6 Flash、および DeepSeek V4 Flash 0731 を評価し、パス率、速度、コストを Claude Fable 5 および GPT-5.6 と比較します。

記事本文:
Kimi K3 はオープン、Opus 5 は良好、DeepSeek V4 Flash は安い: Baba Is You の LLM - Quesma Blog ブログベンチ ベンチマーク About Contact
GitHub
Kimi K3 はオープン、Opus 5 は良好、DeepSeek V4 Flash は安い: Baba Is You の LLM
PNG をダウンロード いくつかのエキサイティングなモデル リリースがあります: Kim K3 、 Grok 4.5 、 Gemini 3.6 Flash 、 Claude Opus 5 、および DeepSeek V4 Flash 0731 。実りの多い7月でした！
これらはすべて、Artificial Analysis Intelligence Index でかなり高いため、Baba Is Bench で再実行しました。これは、エージェントが素敵なパズル ゲーム Baba Is You をプレイするベンチマークであり、OpenAI が ARC-AGI-3 の結果について議論する際に言及したベンチマークです。
このゲームはすでに 7 年前のものですが、軌道は解決策を知っている兆候を示しておらず、SWE ベンチ検証済みリークとははっきりと対照的でした。
これを 2 つのラウンドで行います。最初に、元の投稿と同じ Terminus-2 ハーネスを使用して、初期段階である Intro を 3 回繰り返します。
次に、カスタム ハーネスを使用して、これらを次のステージである湖で実行します。
速度とコストを簡単にチェックするために、最も単純な方法から始めましょう。私たちは、これらのモデルがほとんどのタスクを解決するとほぼ想定していました。わかりやすくするために、他の行を淡色表示にし、新しいモデルを強調表示しました。
モデル 00 馬場さんは 01 どこに行きますか? 02 さて、これは何ですか？ 03 手の届かない 04 まだ手の届かない 05 火山 06 立ち入り禁止 07 芝生ヤード パス@1 ↓ パス@3 ↓ ターン ↓ 出力トークン ↓ 総コスト ↑ 人間的クロード Opus 5 100% 100% 7 23k $17.15 OpenAI GPT-5.5 100% 100% 8 15k $13.62 人間的クロードFable 5 100% 100% 6 29k $40.98 OpenAI GPT-5.6 Sol 100% 100% 11 10k $11.91 キミキミ K3 96% 100% 9 29k $12.56 Anthropic Claude Opus 4.8 96% 100% 35 46k $63.14 Z.ai GLM-5.2 96% 100% 12 75k $6.20 Google Gemini 3.1 Pro 92% 100% 8 38k $12.41 Google Gemini 3.6 Flash 88% 100% 92 8

2k $124.31 Google Gemini 3.5 Flash 88% 100% 111 74k $98.31 OpenAI GPT-5.6 Terra 88% 100% 41 55k $34.28 Grok Grok 4.5 75% 88% 86 68k $50.66 DeepSeek DeepSeek V4 Flash 0731 75% 88% 22 156k $1.16 Anthropic Claude Sonnet 5 67% 88% 22 151k $42.21 OpenAI GPT-5.6 Luna 63% 88% 118 129k $59.32 Hunyuan Hy3 54% 75% 25 76k $3.23 Minimax MiniMax M3 46% 63% 96 173k $34.43 Qwen Qwen3.6 27B 29% 38% 24 82k $5.59 DeepSeek DeepSeek V4 Pro 29% 38% 37 183k $11.49 Qwen Qwen3.7 Max 25% 38% 21 132k $16.98 パスタイムエージェントステップ移動が解決されました タイムアウト 間違った答えが解決されませんでした
Claude Opus 5 はすべての試みを解決しただけでなく、Fable 5 の 2 倍以上安価でした (それでも GLM-5.2 のほぼ 3 倍高価です)。 Kimi K3 は Opus よりも安く、合格率はほぼ同じでした。
Gemini 3.6 Flash…私たちは、途方もなく高価だった前モデルである Gemini 3.5 Flash よりも費用対効果が高いと期待していました。それはわずかに優れていましたが、さらに高価でした。
Grok 4.5 は安くも効率的でもありませんでした。さらに、あるレベルで失敗しました。
そして、DeepSeek V4 Flash 0731 の新しいアップデートは驚異的です。すべてのイントロ レベルを上回るわけではありませんが、Sonnet 5、GPT-5.6 Luna、Grok 4.5 のレベルにおよそ 1/40 のコストで到達します。真の費用対効果の驚異、パレート辺境の英雄。
トークンごとの価格は総コストと同じではありません。一部のモデルではさらに多くの回転が必要です。以下はエージェントのターン数とコストのグラフです。
$1 $2 $5 $10 $20 $50 $100 5 10 20 50 100 エージェント ターン (試行ごとの平均) * 1 レベルの未解決の総コスト ← より意図的であるほどトリガーハッピーになる → ← 安価であるほど高価である → Anthropic Claude Opus 5 OpenAI GPT-5.5 Anthropic Claude Fable 5 OpenAI GPT-5.6 Sol Kim Kim K3 Anthropic Claude Opus 4.8 Z.ai GLM-5.2 Google Gemini 3.1 Pro Google Gemini 3.6 Flash Google Gemini 3.5 Flash OpenAI GPT-5.6 Terra Grok Grok 4.5* DeepSee

k DeepSeek V4 Flash 0731* Anthropic Claude Sonnet 5* OpenAI GPT-5.6 Luna*
更新された DeepSeek V4 Flash が他の LLM よりわずかに安いと言っても過言ではありません。
以前の結果についての議論に興味がある場合は、元の投稿「Baba Is Solved by Fable 5 and GPT-5.6 Sol, but at whatcost?」を参照してください。 。
私たちの方法論に従って、100% pass@3 のモデルである Opus 5、Kimi K3、および Gemini 3.6 Flash を開発しました。
「Baba Is Solved by Fable 5」と「GPT-5.6 Sol」の場合と同様に、どのようなコストがかかりますか? 、コストを合理的に保つために、すべてのレベルを 1 回だけ実行します。
各モデルに対して、ネイティブ ハーネスを選択しました。Opus 5 には Claude Code、Kimi K3 には kimi-cli、そして… さて、Gemini には、以前と同様に問題がありました。
最初に Antigravity CLI を試しましたが、すでにご存知のとおり、他のすべてのモデルを評価するために使用する OpenRouter とは互換性がありません。その後、私 (PM) は Gemini のサブスクリプションを使用しましたが、すぐに使い切ってしまいました。それから私は個人の Google アカウントに請求しました…ただ、（ウェブ検索を行うため）速い場合もあれば、ループ状態になる場合もあることに気づきました。最後に、標準の Terminus-2 ハーネスを使用して実行しました。
前世代と同様に、Gemini 3.6 Flash は高速 (1 秒あたり 135 トークン) ですが、これは同時にお金を非常に早く使ってしまうことを意味します。解決できないレベルでは、進行せずにループし、何千回ものツール呼び出し (1 レベルの場合 1870 回) が行われ、あるケースではコンテキスト ウィンドウの 500,000 トークン以上を消費しました。
ピーク時には、解決できない 9 つのレベルを目的もなくループすることで、1 分あたり 7 ドル (推定: 1 時間あたり 400 ドル以上) を消費しました。 260ドルを支払った後、私（PG）はこの無謀なドライバーにスピード違反の切符を切らせ、財布にこれ以上の被害が出る前に止めさせました。
0m 3m 10m 30m 1h 2h 4h 6h 8h 10h 12h 14h 16h 18h 0 2 4 6 8 10 12 14 累積壁時計リニア・ログレベル解決 人間の君

Google OpenAI OpenAI Anthropic Anthropic Z.ai Google Google クロード オーパス 5 キミ K3 ジェミニ 3.6 フラッシュ GPT-5.6 ソル GPT-5.5 クロード フェイブル 5 クロード オーパス 4.8 GLM-5.2 ジェミニ 3.1 プロ ジェミニ 3.5 フラッシュ
実時間で測定すると、Opus 5 は Fable 5 よりも遅く、Kimi K3 は前回の投稿の GLM-5.2 よりも遅かったです。
注意点が 1 つあります。リリース直後、API クォータが制限されていたときに Kim K3 をテストしました。今はもっと速くなっているかもしれません。
$0 $0.3 $1 $3 $10 $20 $30 $40 $60 $80 0 2 4 6 8 10 12 14 累積コスト線形・ログレベル解決 Anthropic キミ Google OpenAI OpenAI Anthropic Anthropic Z.ai Google Google クロード オーパス 5 キミ K3 ジェミニ 3.6 フラッシュ GPT-5.6 ソル GPT-5.5 クロード ファブル 5 クロードオーパス 4.8 GLM-5.2 ジェミニ 3.1 プロ ジェミニ 3.5 フラッシュ
興味深いのはコストです。
Opus 5 は、最も困難な 2 つのタスクを除いて、安価な GPT-5.6 Sol とより高価な Fable 5 の間にあり、Fable 5 よりもコストがかかります。
したがって、最も難しいタスクを追求していない限り、Opus 5 が最適かもしれません。
Kim K3 は Opus 4.8 と GPT-5.5 の間に位置し、かなり高価です。
したがって、オープン ウェイトは無料または安価であることを意味するものではありません。オープン ウェイトを実行するための計算コストは​​、決して小さくありません。
Gemini 3.6 Flash はテストしたモデルの中で最もパフォーマンスが悪く、残りのレベルを解決するために費やした 260 ドル以上は考慮されていません。このグラフでは、これは合計支出額の 10 ドルにすぎませんが、未完了のタスクのすべてのループを考慮しているわけではありません。
Claude Opus 5 は、Fable 5 とほぼ同じ品質でありながら、より安価です。
キミ K3 は明らかに良いモデルであり、それがハギングフェイスで無差別級になったことは真の奇跡です。
Grok 4.5 は一概には言えません。公式ベンチマークでは良好ですが、少し異なることを実行するとスコアが急落します。それがベンチマックス化なのか、それとも狭い範囲のセットに焦点を当てているだけなのかを判断するのは困難です。

スキル。
ジェミニは落ちました。 Gemini 2.5、3、および 3.1 のリリースは真のランドマークでしたが、最新のものは、生のスキル、価格設定、またはツールの点で実現できません。私たちは双子座モデルの思考力が大好きだったので、それが復活することを願っています。
別の話になりますが、コミュニティはベンチマーク用に Terminus-2 よりも優れたニュートラル ハーネスを必要としています。 Pi と OpenCode ハーネスもテストしましたが、効率はさらに低かったです。 Frontier-Bench v0.1 (以前は Terminal-Bench 3 と呼ばれていました) にネイティブ ハーネスがリストされているのには理由があるのか​​もしれません。
フロンティアの 2026 年 7 月モデルについてはどう思いますか?
LinkedIn 、 X 、 Hacker News 、 Reddit で議論してください。
前 DeepSeek-V4 の論文に隠された再試行に関するレッスン 今後の投稿とリリースにご期待ください
RSS で購読する 関連記事
同様のトピックを引き続き探索する
Baba は Fable 5 と GPT-5.6 Sol によって解決されますが、その代償はどのようなものでしょうか?
私たちはパズル ゲーム Baba Is You を Harbor フレームワークに移植し、Claude、GPT、Gemini、GLM、DeepSeek などの現在のモデルをベンチマークしました。人間の Twitch ストリーマーは、Claude Fable 5 よりも 4 倍高速です。
HN Qwen3.6 27B はローカル開発のスイートスポットです
Qwen3.6 27B は、llama.cpp と OpenCode を使用して、MacBook または NVIDIA RTX でのコーディングに使用できるスマート モデルです。
Qwen3.6 27B 量子化はペリカンを破壊しますか?
私たちは、バイク、ギア、ターミナルベンチ 2.1、および AIME-120 のペリカンを使用して、Hugging Face の Unsloth による Qwen3.6 27B 量子化をテストしました。
Baba は Fable 5 と GPT-5.6 Sol によって解決されますが、その代償はどのようなものでしょうか?
私たちはパズル ゲーム Baba Is You を Harbor フレームワークに移植し、Claude、GPT、Gemini、GLM、DeepSeek などの現在のモデルをベンチマークしました。人間の Twitch ストリーマーは、Claude Fable 5 よりも 4 倍高速です。
HN Qwen3.6 27B はローカル開発のスイートスポットです
Qwen3.6 27B は、ついに MacBook でのコーディングに使用できるスマートなモデルです。

NVIDIA RTX - llama.cpp および OpenCode を使用。
Qwen3.6 27B 量子化はペリカンを破壊しますか?
私たちは、バイク、ギア、ターミナルベンチ 2.1、および AIME-120 のペリカンを使用して、Hugging Face の Unsloth による Qwen3.6 27B 量子化をテストしました。

## Original Extract

We evaluate July 2026 fresh releases Kimi K3, Claude Opus 5, Grok 4.5, Gemini 3.6 Flash, and DeepSeek V4 Flash 0731 on Baba Is Bench, an LLM agent benchmark based on the puzzle game Baba Is You, comparing pass rate, speed, and cost with Claude Fable 5 and GPT-5.6.

Kimi K3 is Open, Opus 5 is Good, DeepSeek V4 Flash is Cheap: LLMs on Baba Is You - Quesma Blog Blog Bench Benchmarks About Contact
GitHub
Kimi K3 is Open, Opus 5 is Good, DeepSeek V4 Flash is Cheap: LLMs on Baba Is You
Download PNG There are a few exciting model releases: Kimi K3 , Grok 4.5 , Gemini 3.6 Flash , Claude Opus 5 , and DeepSeek V4 Flash 0731 . It was a fruitful July!
Since they all are reasonably high in the Artificial Analysis Intelligence Index , we re-ran these on our Baba Is Bench , a benchmark in which agents play the lovely puzzle game Baba Is You , and which got mentioned by OpenAI when discussing their ARC-AGI-3 results .
While the game is already 7 years old, trajectories showed no signs of knowing the solution, a sharp contrast with SWE-bench Verified leaks .
We do it in two rounds: first, three repetitions for the initial stage, the Intro , using the same Terminus-2 harness as in the original post.
Then we run these on the next stage, the Lake , with custom harnesses.
Let’s start with the simplest one, to have a quick check of speed and cost. We pretty much assumed these models would solve most tasks. For clarity, we dimmed other rows, highlighting the new models.
Model 00 baba is you 01 where do i go? 02 now what is this? 03 out of reach 04 still out of reach 05 volcano 06 off limits 07 grass yard pass@1 ↓ pass@3 ↓ Turns ↓ Output tokens ↓ Total cost ↑ Anthropic Claude Opus 5 100% 100% 7 23k $17.15 OpenAI GPT-5.5 100% 100% 8 15k $13.62 Anthropic Claude Fable 5 100% 100% 6 29k $40.98 OpenAI GPT-5.6 Sol 100% 100% 11 10k $11.91 Kimi Kimi K3 96% 100% 9 29k $12.56 Anthropic Claude Opus 4.8 96% 100% 35 46k $63.14 Z.ai GLM-5.2 96% 100% 12 75k $6.20 Google Gemini 3.1 Pro 92% 100% 8 38k $12.41 Google Gemini 3.6 Flash 88% 100% 92 82k $124.31 Google Gemini 3.5 Flash 88% 100% 111 74k $98.31 OpenAI GPT-5.6 Terra 88% 100% 41 55k $34.28 Grok Grok 4.5 75% 88% 86 68k $50.66 DeepSeek DeepSeek V4 Flash 0731 75% 88% 22 156k $1.16 Anthropic Claude Sonnet 5 67% 88% 22 151k $42.21 OpenAI GPT-5.6 Luna 63% 88% 118 129k $59.32 Hunyuan Hy3 54% 75% 25 76k $3.23 Minimax MiniMax M3 46% 63% 96 173k $34.43 Qwen Qwen3.6 27B 29% 38% 24 82k $5.59 DeepSeek DeepSeek V4 Pro 29% 38% 37 183k $11.49 Qwen Qwen3.7 Max 25% 38% 21 132k $16.98 pass time agent steps moves solved timeout wrong answer not solved
Claude Opus 5 not only solved all attempts, but also was more than twice as cheap as Fable 5 (yet, still almost 3x as expensive as GLM-5.2). Kimi K3 was cheaper than Opus, with almost the same pass rate.
Gemini 3.6 Flash… We expected it to be more cost-effective than its ridiculously expensive predecessor, Gemini 3.5 Flash. It was slightly better, yet even more expensive.
Grok 4.5 was neither cheap nor efficient, and even failed at one level.
And the new update of DeepSeek V4 Flash 0731 is a wonder. While it does not beat all intro levels, it reaches the level of Sonnet 5, GPT-5.6 Luna and Grok 4.5 at roughly 1/40 of their cost! A true cost-effective marvel, a Pareto frontier hero .
Price per token is not the same as the total cost. Some models need many more turns. Here is a chart of agent turns and cost.
$1 $2 $5 $10 $20 $50 $100 5 10 20 50 100 agent turns (mean per attempt) * one level unsolved total cost ← more deliberate more trigger-happy → ← cheaper more expensive → Anthropic Claude Opus 5 OpenAI GPT-5.5 Anthropic Claude Fable 5 OpenAI GPT-5.6 Sol Kimi Kimi K3 Anthropic Claude Opus 4.8 Z.ai GLM-5.2 Google Gemini 3.1 Pro Google Gemini 3.6 Flash Google Gemini 3.5 Flash OpenAI GPT-5.6 Terra Grok Grok 4.5* DeepSeek DeepSeek V4 Flash 0731* Anthropic Claude Sonnet 5* OpenAI GPT-5.6 Luna*
To say that the updated DeepSeek V4 Flash is slightly cheaper than other LLMs is an understement!
If you are interested in discussion of previous results, see the original post Baba Is Solved by Fable 5 and GPT-5.6 Sol, but at what cost? .
Per our methodology, we advanced the models with 100% pass@3: Opus 5, Kimi K3, and Gemini 3.6 Flash.
As in our Baba Is Solved by Fable 5 and GPT-5.6 Sol, but at what cost? , we run all levels just once, to keep costs reasonable.
For each model, we chose their native harness: Claude Code for Opus 5, kimi-cli for Kimi K3, and… well, with Gemini, as previously, there were issues.
First I tried Antigravity CLI, but, as we already knew, it is incompatible with OpenRouter, which we use to evaluate all other models. Then I (PM) used my Gemini subscription, but it ate it in no time. Then I charged my personal Google account… just to discover that sometimes it is quick (because it web-searches) and other times it goes in loops. Finally I ran it with a standard Terminus-2 harness.
Like its predecessor, Gemini 3.6 Flash is fast (135 tokens per second), but this also means it spends your dollars extremely quickly. On levels it couldn’t solve, it looped without making progress, made thousands of tool calls (1870 in the case of one level) and in one case consumed over 500k tokens of context window.
At peak, the run consumed $7 per minute (extrapolated: over $400 per hour) by aimlessly looping on 9 levels it couldn’t solve. After $260, I (PG) gave this reckless driver a speeding ticket and stopped it before it did more harm to my wallet.
0m 3m 10m 30m 1h 2h 4h 6h 8h 10h 12h 14h 16h 18h 0 2 4 6 8 10 12 14 cumulative wall clock linear · log levels solved Anthropic Kimi Google OpenAI OpenAI Anthropic Anthropic Z.ai Google Google Claude Opus 5 Kimi K3 Gemini 3.6 Flash GPT-5.6 Sol GPT-5.5 Claude Fable 5 Claude Opus 4.8 GLM-5.2 Gemini 3.1 Pro Gemini 3.5 Flash
Measuring against the wall clock, Opus 5 was slower than Fable 5, and Kimi K3 slower than the previous post’s GLM-5.2.
One caveat: we tested Kimi K3 right after release, when API quotas were throttling it. Who knows, it may be faster now.
$0 $0.3 $1 $3 $10 $20 $30 $40 $60 $80 0 2 4 6 8 10 12 14 cumulative cost linear · log levels solved Anthropic Kimi Google OpenAI OpenAI Anthropic Anthropic Z.ai Google Google Claude Opus 5 Kimi K3 Gemini 3.6 Flash GPT-5.6 Sol GPT-5.5 Claude Fable 5 Claude Opus 4.8 GLM-5.2 Gemini 3.1 Pro Gemini 3.5 Flash
Cost is where it gets interesting.
Opus 5 is between the cheaper GPT-5.6 Sol and the more expensive Fable 5, except for the two hardest tasks, for which it was costlier than Fable 5.
So unless you are after the hardest tasks, Opus 5 might be the way to go.
Kimi K3 falls between Opus 4.8 and GPT-5.5, which is rather expensive.
So, open weights do not mean free or even cheap: compute to run it is a non-trivial cost.
Gemini 3.6 Flash performs the worst amongst the tested models, and that doesn’t even count the over $260 it spent trying to solve the remainder of the levels. While on this chart it is just $10 of total spend, it does not account for all the loops on unfinished tasks.
Claude Opus 5 delivers — almost as good as Fable 5, yet cheaper.
Kimi K3 is clearly a good model, and it is a true miracle that it is now open-weight on Hugging Face .
Grok 4.5 does not generalize: good on official benchmarks, but when doing something slightly different, its score plummets. It is hard to tell whether it is benchmaxxing or just focused on a narrow set of skills.
Gemini fell off. While the releases of Gemini 2.5, 3 and 3.1 were true landmarks, the newest ones do not deliver — not in raw skill, pricing, or tooling. We loved Gemini models for thinking, so we hope they will rebound.
On another note, the community needs a better neutral harness for benchmarks than Terminus-2. We also tested the Pi and OpenCode harnesses, but these were even less efficient. Maybe there is a reason why Frontier-Bench v0.1 (previously called Terminal-Bench 3) lists native harnesses.
And what is your take on the frontier July 2026 models?
Discuss on LinkedIn , X , Hacker News or Reddit .
Previous A lesson about retries, hidden in the DeepSeek-V4 paper Stay tuned for future posts and releases
Subscribe via RSS Related Articles
Continue exploring similar topics
Baba Is Solved by Fable 5 and GPT-5.6 Sol, but at what cost?
We ported the puzzle game Baba Is You to the Harbor framework, and benchmarked current models, including Claude, GPT, Gemini, GLM, and DeepSeek. A human Twitch streamer is 4x faster than Claude Fable 5.
HN Qwen3.6 27B is the sweet spot for local development
Qwen3.6 27B is finally a smart model we can use for coding on MacBook or NVIDIA RTX - with llama.cpp and OpenCode.
Do Qwen3.6 27B quantizations break the pelican?
We tested Qwen3.6 27B quantizations by Unsloth on Hugging Face, with pelicans on bikes, gears, Terminal-Bench 2.1, and AIME-120.
Baba Is Solved by Fable 5 and GPT-5.6 Sol, but at what cost?
We ported the puzzle game Baba Is You to the Harbor framework, and benchmarked current models, including Claude, GPT, Gemini, GLM, and DeepSeek. A human Twitch streamer is 4x faster than Claude Fable 5.
HN Qwen3.6 27B is the sweet spot for local development
Qwen3.6 27B is finally a smart model we can use for coding on MacBook or NVIDIA RTX - with llama.cpp and OpenCode.
Do Qwen3.6 27B quantizations break the pelican?
We tested Qwen3.6 27B quantizations by Unsloth on Hugging Face, with pelicans on bikes, gears, Terminal-Bench 2.1, and AIME-120.
