---
source: "https://tcz.hu/blog/2026/06/12/swearing-and-llms/"
hn_url: "https://news.ycombinator.com/item?id=48506127"
title: "F-bombs don't make LLMs smarter"
article_title: "F-bombs don't make LLMs smarter – Zoltan's Blog"
author: "hntcz"
captured_at: "2026-06-12T18:47:30Z"
capture_tool: "hn-digest"
hn_id: 48506127
score: 2
comments: 0
posted_at: "2026-06-12T16:26:27Z"
tags:
  - hacker-news
  - translated
---

# F-bombs don't make LLMs smarter

- HN: [48506127](https://news.ycombinator.com/item?id=48506127)
- Source: [tcz.hu](https://tcz.hu/blog/2026/06/12/swearing-and-llms/)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T16:26:27Z

## Translation

タイトル: F-bomb は LLM を賢くしない
記事タイトル: F-bombs don't make LLMs Smarter – Zoltan's Blog
説明: レゴ ブロックを踏んだ直後に誰かが数学パズルを解くように頼んだと想像してください。ベースラインよりも良くなるつもりですか、それとも悪くなりますか?関連して、LLM にレゴ ブロックを踏ませることができますか?

記事本文:
F-bomb は LLM を賢くしない |ゾルタンのブログ
について
F-bomb は LLM を賢くしない
あなたがレゴ ブロックを踏んでいるときに、誰かがあなたに数学パズルを解くように頼んだと想像してください。ベースラインよりも良くなるつもりですか、それとも悪くなりますか?関連して、LLM にレゴ ブロックを踏ませることができますか?
この投稿は肯定的な結果があればもっと面白いものになっていたでしょうが、無効な結果を公開するという精神でここまでにします。
OpenAI の 2024 o1 モデルは、科学的な質問に答えるという点で史上初めて人間を上回りました。 1 秘密のソースは、テスト時に生成された、RL で訓練された思考連鎖 (CoT) 推論でした。
CoT について考える 1 つの方法: モデルは元のプロンプトを改良します。ユーザー プロンプトは、モデルの表現空間の曖昧な場所に誘導される粗雑なクエリです。推論チェーンでは、モデルは独自のコンテキストを詳しく説明し、次のトークンを取得するための高品質でよりレーザーに焦点を当てたクエリを取得します。
LLM 順方向パスの最後のステップは、最後の残差ストリーム状態 $h$ と非埋め込み行列 $W_U$ を乗算することです。残留ストリーム状態は大きなベクトルであり、その方向は各モデル層で調整されます。これは、モデルがプロンプトと関連する「知識」を圧縮した数値のリストです。追加の制約、関連する事実、および言語化された計画を備えたより精巧なプロンプトは、最後のベクトルで「より適切な方向を向いた矢印」を作成します。
幾何学的に言えば、$h_{\text{sharp}} - h_{\text{blunt}}$ の差が CoT によって得られるものです。その結果、より高品質のトークンが選択されます。
「大声で考える」という人間の類似点がよく引用されますが、そのメカニズムはおそらく大きく異なります。しかし、思考連鎖は LLM の著しく擬人化された行動であるように見えますが、それだけではありません。
モデルはどのくらいの時間を「考える」必要があるかをどのようにして知るのでしょうか?

簡単な答えは、思考終了トークンが生成されると CoT が終了するということです。 CoT モデルは、実際の答えを生成する前に、特別なトークン ( < Thinking> と </ Thinking> など) 間の推論チェーンをラップするようにトレーニングされます。このトレーニングにより、思考の連鎖が完了した後、</ Thinking> が最高得点のトークンになります。チェーンの長さは直接制御されず、十分な推論を行った時点でモデルが学習した感覚から現れます。
しかし、生成時にチェーンの長さを制御できたらどうなるでしょうか?ムエニホフら。 s1: Simple test-time scaling というタイトルの論文で、これを両方向で行いました。彼らが「予算強制」と呼ぶプロセスでは、生成されたチェーンが制限を超えたときに単純に思考終了トークンを注入します。 CoT を短くするのではなく、長くしたい場合はどうなりますか?この論文ではいくつかの戦略が列挙されています。彼らは、最も効果的な解決策は、モデルが停止したいときに思考の連鎖に「待つ」という言葉を挿入することであることを発見しました。
stop = エンコード ( "</ Thinking>" )
トークン = エンコード (プロンプト)
範囲内の i の場合 ( n_ignore + 1 ):
out = 生成 ( トークン , stop = "</ Thinking>" )
トークン += アウト 。トークン
出ていたら。 hit_length_cap または i == n_ignore :
休憩
トークン = トークン [: - len ( stop )]
トークン += エンコード ( "待機" )
この強制的に組み込まれたトークンにより、モデルは元の推論に疑問を抱き、その動作を再確認します。結果として得られたモデルは、競合数学において o1-preview を最大 27% 上回りました。 2
これについては公開された文書はなく、推論 RL の成果物である可能性が高いですが、クロードの推論チェーンでは常に Wait とActual が使用されており、場合によっては応答にも注目せずにはいられませんでした。以下は、自己不信がいかに逆効果になるかを示す、特にばかげた例です。
予算の強制は、収入が増えたときに役立ちます。

トークンは一度に 1 つずつ増えますが、そのアルミ箔の帽子は脱ぎます。
強制的な自己疑念も驚くべき擬人化です。 EmotionPrompt で実証されているように、感情的な刺激も同様です (Li et al. 2023)。ベンチマークされたモデルは、「これは私のキャリアにとって非常に重要です」や「確信したほうがいいです」などの即時追加を加えてより良い答えを与えるようにいじめられた場合に、より良いパフォーマンスを発揮しました。
s1 論文では、注入された他のトークン (「うーん」と「代替」) の予算強制をテストしました。感情を刺激するレゴを踏む効果を得るには、もっとスパイシーなものをいくつか試しなければなりませんでした。
候補者は、5 歳児が見つけた爆笑バケツから直接選ばれたものでした。
ああ、まさに AIME24 と MATH500 で結びつきました。たわごとは、どちらかといえば、少し悪いです。待機は依然として最もパフォーマンスの高い予算強制トークンですが、試してみる新しい提案を歓迎します。
(注: 上記のテストは n_ignore=2 で実行されました。つまり、2 つの思考終了トークンを無視しました。s1 論文ではさまざまな構成をテストしましたが、どれもクソスナップのパフォーマンスを向上させるとは思えません。)
悪口には明らかな自信喪失が十分ではないという仮説に基づいて、数トークンの長さのフレーズをさらにいくつかテストしました。
「それは間違っています」の亜種は強すぎることが判明し、AIME24 の結果を大幅に落としました。
冒涜は推論の質に影響を与えません。
この投稿の実験には、Vast.ai 上で実行される GPU コンピューティングの約 734.43 ドルが必要でした。すごいね。 Cloud GPU (8xH100) は、s1 所見を再現し、独自のアブレーションを実行するために 32B Qwen モデルを実行する場合にのみ必要でした。忠実に再現するにはfloat32が必要なので、手を抜かないことにしました。
機械学習の研究、執筆、学習にかかるコストを透明にするために、私はこれらの数字を公開しておきます。最新の ML タスクでは、多くの場合、

研究者、学生、テクノロジーに精通した子供たちが常に利用できるわけではない膨大な計算能力。彼らは遅れをとり、テクノロジーエリートによって守られる危険があります。私は、計算量の多い研究や探索的なプログラミングに資金を提供するために、少額の GPU ミニ助成金を定期的に提供しています。
GPQA ダイヤモンド ベンチマークにおける生物学、物理学、化学の質問に関する博士レベルの専門家。 ↩
s1 モデルは、AIME24、MATH500、および GPQA Diamond でベンチマークが行われました。予算の強制に加えて、監視付きの微調整もパフォーマンスの向上に貢献しました。 ↩
あなたも興味があるかもしれません
特異な学習理論: AI は氷が溶けるように学習する
2026 年 2 月 26 日

## Original Extract

Imagine someone asks you to solve a math puzzle right after you stepped on a LEGO brick. Are you going to do better or worse than the baseline? Relatedly, can we make LLMs step on LEGO bricks?

F-bombs don’t make LLMs smarter | Zoltan’s Blog
About
F-bombs don't make LLMs smarter
Imagine someone asks you to solve a math puzzle right as you step on a LEGO brick. Are you going to do better or worse than the baseline? Relatedly, can we make LLMs step on LEGO bricks?
This post would have been a lot funnier with a positive result, but in the spirit of publishing null results, here it goes.
OpenAI’s 2024 o1 model was the first ever to beat humans in answering scientific questions. 1 The secret sauce was RL-trained chain-of-thought (CoT) reasoning generated at test-time.
One way to think about CoT: the model refines the original prompt. The user prompt is a crude query that drops you somewhere vague in the model’s representation space. In the reasoning chain, the model elaborates on its own context, getting a higher quality, more laser-focused query to retrieve the next token by.
The last step in an LLM forward pass is multiplying the last residual stream state $h$ with the unembedding matrix $W_U$. The residual stream state is a large vector whose direction gets refined at each model layer. It is the list of numbers the model compressed its prompt plus its relevant “knowledge” into. A more elaborate prompt with added constraints, relevant facts, and a verbalized plan makes for a “better-aimed arrow” in that last vector.
The difference $h_{\text{sharp}} - h_{\text{blunt}}$ is what CoT is buying you, geometrically speaking. The result is a higher-quality token being selected.
The human analog of “thinking out loud” is often cited but the mechanism is likely very different. Yet, chain-of-thought appears to be a remarkably anthropomorphic behavior of LLMs, and it’s not the only one.
How does the model know how long to “think” for? The simple answer is, when the end of thinking token is generated, the CoT ends. CoT models are trained to wrap the reasoning chain between special tokens (something like <thinking> and </thinking> ) before generating the actual answer. That training makes the </thinking> the highest-scoring token after the chain-of-thought has run its course. The length of the chain is not directly controlled, it emerges from the model’s learned sense of when it has done enough reasoning.
But what if you could control the chain’s length at generation time? Muennighoff et al. did this in both directions in their paper titled s1: Simple test-time scaling . In a process they call budget forcing , they simply inject an end-of-thinking token when the generated chain exceeds a limit. What happens if you want the CoT to be longer, not shorter? The paper lists several strategies. They found that the best-performing solution is injecting the word “Wait” into the thinking chain when the model wants to stop.
stop = encode ( "</thinking>" )
tokens = encode ( prompt )
for i in range ( n_ignore + 1 ):
out = generate ( tokens , stop = "</thinking>" )
tokens += out . tokens
if out . hit_length_cap or i == n_ignore :
break
tokens = tokens [: - len ( stop )]
tokens += encode ( "Wait" )
This forced-in token makes the model doubt its original reasoning and re-check its work. The resulting model beat o1-preview by up to 27% on competition math. 2
There is no public documentation on this, and it is most likely an artifact of the reasoning RL, but I couldn’t help noticing Wait and Actually all the time in Claude’s reasoning chain and sometimes the response too. Below is a particularly dumb example showing how self-doubt can be counter-productive.
Budget forcing can come in handy when your revenue increases one token at a time, but I will take that tinfoil hat off.
Forced self-doubt is another surprising anthropomorphism. So is emotional stimulus, as demonstrated by EmotionPrompt ( Li et al. 2023 ). The benchmarked models performed better when bullied into giving better answers with prompt additions like “This is very important to my career” and “You’d better be sure”.
The s1 paper tested other injected tokens (“Hmm” and “Alternatively”) for budget forcing. I had to try a few more, spicier ones for that emotionally stimulating stepping-on-LEGO effect.
The candidates came straight from the things-a-five-year-old-finds-hilarious bucket.
Fuck ties Oh exactly on AIME24 and MATH500. Shit is, if anything, slightly worse. Wait remains the best-performing budget forcing token, but I’d welcome new suggestions to try.
(Note: The tests above were run on n_ignore=2 , that is, we ignored two end-of-thinking tokens. The s1 paper tested different configurations, but I doubt that any of them would improve the fuck-shit-oh-snap performance.)
On the hypothesis that there’s not enough explicit self-doubt in the expletives, I tested a few more phrases that were several tokens long.
The “that’s wrong” variants turned out to be too strong, and dropped the AIME24 results significantly.
Profanity has no effect on reasoning quality.
The experiments in this post took about $734.43 of GPU compute, run on Vast.ai. Big oof. Cloud GPU (8xH100) was needed only to run the 32B Qwen model in order to reproduce the s1 findings and run my own ablations. The faithful reproduction required float32 and I decided not to cut corners.
I keep these numbers public to be transparent about the cost of machine learning research, writing, and learning. Modern ML tasks often require enormous compute capacity, which is not always available to researchers, students and tech-savvy kids. They risk falling behind and being gatekept by a tech-elite. I regularly offer a small GPU mini-grant to fund compute-heavy research and exploratory programming.
PhD-level experts on biology, physics, and chemistry questions in the GPQA Diamond benchmark. ↩
The s1 model was benchmarked on AIME24, MATH500, and GPQA Diamond. Besides budget forcing, supervised fine-tuning played a part in the improved performance. ↩
You might also be interested in
Singular Learning Theory: AI learns like ice melts
Feb 26, 2026
