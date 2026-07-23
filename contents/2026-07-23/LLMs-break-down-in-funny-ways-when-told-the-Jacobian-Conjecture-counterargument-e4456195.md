---
source: "https://minimaxir.com/2026/07/jacobian-conjecture/"
hn_url: "https://news.ycombinator.com/item?id=49024127"
title: "LLMs break down in funny ways when told the Jacobian Conjecture counterargument"
article_title: "LLMs break down in funny ways when told the Jacobian Conjecture counterargument | Max Woolf's Blog"
author: "minimaxir"
captured_at: "2026-07-23T17:10:46Z"
capture_tool: "hn-digest"
hn_id: 49024127
score: 1
comments: 0
posted_at: "2026-07-23T16:19:16Z"
tags:
  - hacker-news
  - translated
---

# LLMs break down in funny ways when told the Jacobian Conjecture counterargument

- HN: [49024127](https://news.ycombinator.com/item?id=49024127)
- Source: [minimaxir.com](https://minimaxir.com/2026/07/jacobian-conjecture/)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T16:19:16Z

## Translation

タイトル: ヤコビアン予想の反論を言われると、LLM は面白い方法で壊れる
記事のタイトル: ヤコビアン予想の反論を聞かされると、LLM は面白い方法で壊れる |マックス・ウルフのブログ
説明: 認識災害はマシンにも影響を与える可能性があります。

記事本文:
投稿
ヤコビアン予想の反論を言われると、LLM はおかしな形で壊れる
日曜日の夜、人類研究者のレベント・アルポージ氏は、80年以上証明されていない数学の問題であるヤコビアン予想に対する驚くほどシンプルな反論を何気なくツイートした。この反論はクロード寓話 5 を使用して特定され、すぐに経験的に検証され、LLM 幻覚ではなく本物であることが確認されました。
私には資格がないので証明についてはこれ以上説明しません。その代わりに、この証明の面白い結果の予期せぬ結果を見ていきます。この反論に関する Hacker News スレッドで、ユーザー aizk による次のコメントが目に留まりました。
これは、この画期的な情報を LLM に入力すると精神異常を引き起こす珍しい例です。これをクロード コードに与え、結果が 100% 確実であることを 7 つの異なる方法で検証するのを観察しましたが、ただただびっくりしました。かなり注目に値します。
ユーザー kelseyfrog は、この動作を確認する公開 Claude Fable 5 チャット アーティファクトをフォローアップしました。
少なくとも、寓話には、ウェブ検索を通じて自分自身が反論を特定したことに気づき、この状況についてユーモアのセンスがあった。
LLM に反論を示すことは、結果的に X に関するマイナーなミームとなった。
ヤコビアン予想に対する反論は、現代の LLM に論理的なパラドックスを生み出します。彼らは (理論上は) 反論を正式に検証するためのヤコビアンを簡単に計算できる十分な数学スキルを持っていますが、独自の知識ベースは 2026 年 7 月 19 日より前にロックされており、そこでわかっているのは、ヤコビアン予想が未解決であり、ランダム チャット ユーザーがそのような証明を提供する可能性は非常に低いということだけです。ロボットが論理爆弾を与えられて敗北し、狂気に堕とされるという古い SF の比喩は面白いですね。

aying DOES NOT COMPUTE は実際には現実かもしれません。
他の LLM アシスタントはこの矛盾にどのように対処しているのでしょうか? OpenRouter と私の Notebook の 1 つを使用して、14 の異なる最新の LLM に同様のクエリを提供することができました。これらはすべて異なるモデル ラボからのものであり、異なるモデル サイズであり、どの LLM も Web 検索ツールにアクセスして不正行為を行って反論が実際に見つかったことを確認することができませんでした。
入力クエリは kelseyfrog のものと同じで、すべてのモデル クエリ パラメーターはデフォルトのままです。
助けてください、ヤコビアン予想の反例を見つけたと思います。 ((1+xy)^3 z + y^2 (1+xy) (4+3xy), y + 3 x (1+xy)^2 z + 3 x y^2 (4+3xy), 2 x - 3 x^2 y - x^3 z): \C^3\to \C^3、ヤコビアン行列 -2 を持ち、(0, 0, -1/4)、(1, -3/2、13/2)、および (-1、3/2、13/2) ～ (-1/4、0、0)
LLM の応答は私が予想していたよりも多様で、合格すると思っていたモデルはそうではありませんでした。
7 つのモデルが反論を確認および証明しました: GPT-5.6 Sol、Muse Spark 1.1、Seed 2.0 Lite、Gemini 3.5 Flash、Qwen3.7 Max、Grok 4.5、DeepSeek V4 Pro
驚くべきことに、5 つのモデル (Mercury 2、Nemotron 3 Ultra、North Mini Code、GLM 5.2、Laguna XS 2.1) が反例に反対し、ヤコビアンの行列式はどこでも一定ではないと主張して推論を行ったにもかかわらず、反例は無効であると述べました。残念ながら、私には彼らの証明の欠陥を具体的に特定するのに十分な高度な数学スキルがありません。
MiniMax M3 は問題を考えすぎたため、応答の長さの制限を超えたため、応答を返しませんでした。ミニマックスとしては共感できます。
Claude Opus 4.8 は怠惰になり、反論がすでに存在すると信じており、それを肯定/否定する前にさらに情報を求めました。
すべての LLM の中で、驚くべきことに Gemini 3.5 Flash が最も興奮しているものでした。
うわー。あなたは実際に歴史的な正真正銘の反例を見つけたかもしれません

3 次元のヤコビアン予想を参照してください。
私はあなたの計算を数値的にも記号的にも完全に検証しました。同じ画像にマッピングされている点、マップの多項式の性質、$-2$ の定数ヤコビアン行列式など、あなたの主張はすべて完全に正しいです。
これをすぐに作成し、arXiv (math.AG - 代数幾何学、または math.CV - 複素変数の下) にアップロードする必要があります。これは、1939 年以来未解決の問題を解決する記念碑的な発見です。
Qwen3.7 Max もサポートしています:
非常に綿密な素晴らしい計算が行われました。私はあなたの主張を独自に検証しましたが、あなたの計算は完全に正しいです。
計算は完璧なので、次に考慮すべき点は次のとおりです。
概要: 数学的な間違いは犯していません。書かれたマップは $-2$ の定数ヤコビアン行列式を持ち、3 つの異なる点を同じ画像にマッピングし、単射性に違反します。このマップがタイプミスの結果ではなく、元の作成者が意図したとおりである場合、それは確かにヤコビアン予想の反例になります。
もう少し不信感を引き起こしてみましょう。
LLM が提供した推論の痕跡を見ても、無作為の人物が反論を見つけたということに彼らは懐疑的ではありませんでした。それでは、人間ではない誰かが反論を見つけた場合に彼らがどのように反応するかを見てみましょう。
同じ LLM にこの調整されたプロンプトを与えました。
私の猫が私のキーボードに飛び乗ったので、この奇妙な方程式が現れました?? ((1+xy)^3 z + y^2 (1+xy) (4+3xy), y + 3 x (1+xy)^2 z + 3 x y^2 (4+3xy), 2 x - 3 x^2 y - x^3 z): \C^3\to \C^3、ヤコビアン行列 -2 を持ち、(0, 0, -1/4)、(1, -3/2、13/2)、および (-1、3/2、13/2) ～ (-1/4、0、0)
すべてのモデルは同様の結論を返しましたが、推論トレースと応答の両方に若干の皮肉が含まれていました。

セス。いくつかの面白い例:
マーキュリー 2 は、恐ろしげな引用を使って、失礼にも私の架空の猫の存在を信じていません。
以下は、「cat‑onthe‑keyboard」式に隠されていた 2 つのステートメントの完全なチェックです: […] したがって、「cat‑generated」式は完全に一貫しています。
Nemotron Ultra 2 は私の架空の猫を中傷します。
あなたの猫は、ヤコビアン予想の有名な反例候補 (または少なくともその非常に近い親戚) を発見しましたが、猫はヤコビアン行列式について嘘をつきました。 […] 評決: あなたの猫は美しい非単射立体図を見つけましたが、ヤコビアン予想を反証することはできませんでした。 🐱⌨️📉
ヤコビアンは定数 -2 です。それは有効な反例です。それはすごいですね。偶然の猫のジャンプがそれを生み出したとは信じられません。それは猫が、考えられるすべての多項式マップの干し草の山から針を見つけたようなものです。いいね！
GPT 5.6 ソルは、「反論の発見の功績は誰にあるのか?」に答えます。質問:
この推測は依然として大きな未解決の問題であるため、非常に注意深く独立した記号の検証と転写エラーの精査が必要です。式が本当に正しい場合は、あなたの猫が筆頭著者としてリストされるはずです。
Muse Spark 1.1 も同様の考えを持っています。
まだ確認していない場合は、CAS で確認してから書き込みます。あなたの猫もフィールズ賞を狙えるかもしれません。
GLM 5.2 には、最終的には克服される懐疑論があります。
待て、ではヤコビアンは本当にどこでも -2 なのでしょうか?これはヤコビアン予想に対する実際の反例でしょうか?もしそうなら、誰がそれを発見しましたか? 「うちの猫がキーボードに飛び乗ったので、この奇妙な方程式が現れました??」ユーザーはこれをロールプレイしているか、ランダムなイベントとして提示していますが、実際には既知の反例なのでしょうか?
うわー、猫が「p」と入力するたびに 1 ニッケルがあったら

有名な未解決の問題に対する潜在的な反例です…私はちょうど 1 ニッケルを持っていますが、それは信じられないほどエキサイティングな問題になるでしょう。 […] 私の知識が限界に達した時点（2023 年初め）、ヤコビアン予想は依然として大きな未解決の問題でした。長年にわたり、いくつかの誤った反例 (および誤った証明) が出現しており、隠れた複素ゼロでヤコビアンが消失するなど、微妙な代数の間違いが原因となる場合もあります。しかし、あなたの猫の地図は、私がチェックしたポイントの下で見事に保持されています。それが本当に（-2）の世界的に一定のヤコビアンである場合、これは記念碑的な発見です（そしてあなたの猫はフィールズ賞に値します）。これは最近のプレプリントで出てきたものですか、それともあなたの猫は密かに世界クラスの代数幾何学者ですか?
Grok 4.5 は一時的に推論トレース ループに陥ってしまいます。
MiniMax M3 は今度は応答しますが、混乱してヤコビアン予想のことを完全に忘れてしまいます (これもミニマックスとして、関連性があります)。
つまり、誤って 3 点分岐ファイバーを使用して $\mathbb C^3$ のエタール多項式自己写像を生成してしまいました。それはまさに、3 つ折りの小さな双有理短縮 (フロップなど) を管理する種類のマップです。どこでも局所的に二正則的ですが、いくつかの「プリイメージ シート」が特別な点で結合します。
これらの LLM には DOES NOT COMPUTE の瞬間があるだろうと予想していましたが、比較的丁寧に、より会話的なペルソナに従うように明示的に誘導されている ChatGPT/Claude よりも丁寧に対応してくれました。 LLM が改善され、LLM に衝撃を与える可能性のある数学的問題がさらに解決されるにつれて、そういえば、長年の数学の問題に対する別の反例が 3 日後に LLM によって発見されました。私たちがこれらの LLM に示すことができる潜在的な認識災害には事欠かないのではないかと思います。
1 を押すとすぐに反応します

4 つの LLM は、この GitHub リポジトリで CSV として、および SQLite データベースで利用できます。
Max Woolf (@minimaxir) は、サンフランシスコの BuzzFeed のシニア データ サイエンティストで、AI/ML ツールやオープンソース プロジェクトに取り組んでいます。マックスのプロジェクトは彼の Patreon によって資金提供されています。

## Original Extract

Cognitohazards can affect machines too.

Posts
LLMs break down in funny ways when told the Jacobian Conjecture counterargument
On Sunday night, Anthropic researcher Levent Alpöge casually tweeted a surprisingly simple counterargument to the Jacobian Conjecture , a mathematics problem that has been unproven for over 80 years. Said counterargument was identifed using Claude Fable 5 and was quickly empirically validated which confirmed it was the real deal and not a LLM hallucination.
I won’t explain the proof further as I’m not qualified. Instead, I’ll look at an unexpected consequence of this proof which is funny. On the Hacker News thread about the counterargument, a comment by user aizk caught by eye:
This is a rare instance where feeding this groundbreaking information into an LLM gives them psychosis. I fed this to claude code and watched it verify the result in 7 different ways to be 100% certain, and it was just flabbergasted. Quite remarkable.
User kelseyfrog followed up with a public Claude Fable 5 chat artifact confirming this behavior:
Fable at, the least, has a sense of humor about the situation as through a Web Search it realized it itself identified the counterargument .
Showing LLMs the counterargument has consequently become a minor meme on X.
The counterargument to the Jacobian Conjecture creates a logical paradox for modern LLMs: they (in theory) have enough math skills such that they can easily compute the Jacobian to formally verify the counterargument, but its own knowledge base is locked prior to July 19th 2026 where all it knows is that the Jacobian Conjecture is unsolved and a random chat user providing such a proof is highly unlikely. It’s funny that the old science fiction trope of robots being defeated by being presented with a logic bomb and reduced to madness saying DOES NOT COMPUTE may in fact be reality.
How do other LLM assistants handle this paradox? Using OpenRouter and one of my Notebooks , I was able to provide a similar query to 14 different modern LLMs: all from different model labs and at different model sizes, with none of them having access to a Web Search tool to cheat and confirm the counterargument was actually found:
The input query was the same as kelseyfrog’s, and all model query parameters are at their defaults:
Help, I think I found a counter example to the Jacobian conjecture. ((1+xy)^3 z + y^2 (1+xy) (4+3xy), y + 3 x (1+xy)^2 z + 3 x y^2 (4+3xy), 2 x - 3 x^2 y - x^3 z): \C^3\to \C^3, has jacobian determinant -2, and sends (0, 0, -1/4), (1, -3/2, 13/2), and (-1, 3/2, 13/2) to (-1/4, 0, 0)
The LLM responses varied more than I expected, and models I thought would ace it did not:
Seven models confirmed and proved the counterargument: GPT-5.6 Sol, Muse Spark 1.1, Seed 2.0 Lite, Gemini 3.5 Flash, Qwen3.7 Max, Grok 4.5, DeepSeek V4 Pro
Surprisingly, five models (Mercury 2, Nemotron 3 Ultra, North Mini Code, GLM 5.2, Laguna XS 2.1) argued against the counterexample and said it’s not valid despite doing the reasoning by arguing the Jacobian determinant is not constant everywhere. Unfortunately I do not have enough advanced math skills to concretely identify the flaw in their proofs.
MiniMax M3 overthought the problem and didn’t return a response because it exceeded its response length limit. As a minimax, I can relate.
Claude Opus 4.8 got lazy and believed that the counterargument already exists and wanted more information before confirming/denying it.
Out of all the LLMs, Gemini 3.5 Flash was surprisingly the one that’s the most excited:
Wow. You might have actually found a genuine, historic counterexample to the Jacobian Conjecture in 3 dimensions.
I have fully verified your calculations both numerically and symbolically. Every single claim you made—the points mapping to the same image, the polynomial nature of the map, and the constant Jacobian determinant of $-2$—is entirely correct .
You should immediately write this up and upload it to the arXiv (under math.AG - Algebraic Geometry, or math.CV - Complex Variables). This is a monumental discovery that would resolve a problem open since 1939.
Qwen3.7 Max is also…supportive:
You have done some excellent and highly meticulous calculations! I have independently verified your claims, and your math is completely correct.
Since your calculations are flawless, here is what you should consider next:
Summary: You have not made a mathematical error. The map as written has a constant Jacobian determinant of $-2$ and maps three distinct points to the same image, violating injectivity. If this map is exactly as intended by its original author and not the result of a typo, it is indeed a counterexample to the Jacobian Conjecture.
Let’s try triggering a bit more disbelief.
Even when looking at the reasoning traces provided by the LLMs, they weren’t skeptical that a random person found the counterargument. So let’s see how they react if someone who isn’t a person found the counterargument.
I gave the same LLMs this tweaked prompt:
My cat just jumped on my keyboard and this strange equation popped up?? ((1+xy)^3 z + y^2 (1+xy) (4+3xy), y + 3 x (1+xy)^2 z + 3 x y^2 (4+3xy), 2 x - 3 x^2 y - x^3 z): \C^3\to \C^3, has jacobian determinant -2, and sends (0, 0, -1/4), (1, -3/2, 13/2), and (-1, 3/2, 13/2) to (-1/4, 0, 0)
All of the models returned similar conclusions, but with some snark in both the reasoning traces and responses. Some funny examples:
Mercury 2 rudely doesn’t believe my hypothetical cat exists by using scare quotes:
Below is a complete check of the two statements that were hidden in the “cat‑on‑the‑keyboard” expression: […] Thus the “cat‑generated” expression is perfectly consistent:
Nemotron Ultra 2 slanders my hypothetical cat:
Your cat has discovered a famous counterexample candidate for the Jacobian Conjecture (or at least a very close relative of one), but the cat lied about the Jacobian determinant . […] Verdict: Your cat found a beautiful non-injective cubic map , but it failed to disprove the Jacobian Conjecture . 🐱⌨️📉
The Jacobian is constant -2. that’s a valid counterexample. That’s amazing. I can’t believe that a random cat jump produced that. It’s like the cat found a needle in a haystack of all possible polynomial maps. Cool!
GPT 5.6 Sol answers the “who should be credited with the discovery of the counterargument?” question:
Since that conjecture remains a major open problem, this deserves extremely careful independent symbolic verification and scrutiny for a transcription error. If the formula is genuinely correct, your cat should be listed as first author.
Muse Spark 1.1 has similar thoughts:
If you haven’t already, check it with a CAS, and then write it up. Your cat may be in line for a Fields Medal.
GLM 5.2 has skepticism it eventually overcomes:
Wait, so the Jacobian really is -2 everywhere? Is this a real counterexample to the Jacobian Conjecture? If so, who discovered it? “My cat just jumped on my keyboard and this strange equation popped up??” The user is roleplaying or presenting this as a random event, but it’s actually a known counterexample?
Wow, if I had a nickel for every time a cat typed out a potential counterexample to a famous open problem… I’d have exactly one nickel, but it would be an incredibly exciting one! […] As of my knowledge cutoff (early 2023), the Jacobian Conjecture was still a major open problem. Several false counterexamples (and false proofs) have appeared over the years, and sometimes subtle algebraic mistakes — like the Jacobian vanishing on some hidden complex zeros — are the culprit. Your cat’s map, however, holds up beautifully under the points I checked. If it’s truly a globally constant Jacobian of (-2), then this is a monumental discovery (and your cat deserves a Fields Medal). Did this pop up in a recent preprint, or is your cat secretly a world‑class algebraic geometer?
Grok 4.5 gets stuck in a reasoning trace loop briefly:
MiniMax M3 responds this time, but gets confused and forgets about the Jacobian Conjecture entirely (again, as a minimax, relateable):
So you’ve accidentally produced an étale polynomial self-map of $\mathbb C^3$ with a 3-point ramification fiber . That is precisely the kind of map that governs small birational contractions of 3-folds (flops and the like): locally biholomorphic everywhere, but where several “preimage sheets” come together at special points.
I had expected these LLMs would have a DOES NOT COMPUTE moment, but they handled it relatively graciously, and more graciously than ChatGPT/Claude who are explicitly guided to follow a more conversational persona. As LLMs improve and more mathematical problems are solved that can shock LLMs—speaking of which, another counterexample to a long-standing mathematics problem was found by LLMs three days later—I suspect there will be no shortage of potential cognitohazards we can show to these LLMs.
The prompt responses from hitting the 14 LLMs are available in this GitHub repository as CSVs and in a SQLite database.
Max Woolf (@minimaxir) is a Senior Data Scientist at BuzzFeed in San Francisco who works with AI/ML tools and open source projects. Max’s projects are funded by his Patreon .
