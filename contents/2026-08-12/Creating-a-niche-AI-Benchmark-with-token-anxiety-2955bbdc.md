---
source: "https://thijsbrits.nl/blog/20260804/"
hn_url: "https://news.ycombinator.com/item?id=49269347"
title: "Creating a niche AI Benchmark with token anxiety"
article_title: "Creating a niche AI Benchmark with token anxiety — Thijs Brits"
author: "searchingforit"
captured_at: "2026-08-12T09:02:02Z"
capture_tool: "hn-digest"
hn_id: 49269347
score: 1
comments: 0
posted_at: "2026-08-12T08:27:11Z"
tags:
  - hacker-news
  - translated
---

# Creating a niche AI Benchmark with token anxiety

- HN: [49269347](https://news.ycombinator.com/item?id=49269347)
- Source: [thijsbrits.nl](https://thijsbrits.nl/blog/20260804/)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T08:27:11Z

## Translation

タイトル: トークン不安を伴うニッチな AI ベンチマークの作成
記事のタイトル: トークン不安を伴うニッチな AI ベンチマークの作成 — Thijs Brits

記事本文:
← ブログ
トークン不安を伴うニッチな AI ベンチマークの作成
私たちが最も注目する AI ベンチマークは、モデルがどれだけうまくコーディングできるか、数学や科学がどれだけ得意かなどに焦点を当てています。
しかし、ニッチなベンチマークは楽しいです。スケートベンチ (トリックの仕組みをテキストで説明してスケートボードのトリックを推測する) と自動販売機 (AI が自動販売機会社のシミュレーションを実行する) を考えてみましょう。
主なベンチマークには、AI がどのように改善されているかを測定するという目的があります。よりニッチで一見ばかばかしいベンチマークは、モデルで何ができるか、何ができないかを創造的に探求することができます。もっと必要だと思います。もっと予想外のことに遭遇したり、単に創作を楽しんだりするかもしれません。
特に、ランダムな AI ベンチマークについて読んでいるような人であれば、おそらくあなたも今感じている現実の制約が 1 つあります。
ある場所でのサブスクリプションの上限、いくつかの小さなサブスクリプションではキャンセルを忘れがち、そのプラットフォームには未使用のクレジットがまだ残っているなど。
この AI 関連の作業は、非常に刺激的であり、面倒であり、同時に高価でもあります。でも、遊びをやめられないんです。
実行コストが安いベンチマークを作成しようとするとどうなるでしょうか。あるいはさらに良いのは、大量のトークンを出力しないインセンティブがあるかどうかです。
最も単純な形式は、ベンチマークのような試験です。1 つのプロンプト問題とその答えを正誤で評価します。
ニッチなベンチを構築するのに最適なシンプルなセットアップ。質問自体に創造性をすべて注ぐことができますが、私はこの構造を少し混乱させるものが欲しかったです。
トークンは 1 つずつ入ってきますが、LLM の出力速度ではそれを忘れてしまうほどです。ストリーミングをオンにした場合、API レベルでのみ表示されます。では、トークンごとの評価の粒度に焦点を当てたらどうなるでしょうか。個々のトークンごとに動的にベンチマークが行われたことはあまり見たことがありません。それをやってみよう！トークンマックスに反対しましょう。

tokenminimize、tokenmin... 言いたいことはわかりますね。
ゲーム: 説明、ストリーミング、推測
アイデア: 秘密の単語を推測するゲーム。
2 人のモデルがペアとして、説明者と推測者として協力します。
説明モデルは説明する秘密の単語を取得し、その単語の説明をトークンごとにストリーミングし、推測モデルはトークンが追加されるたびに、説明しようとしている単語を推測できます。推測モデルは、単語を推測することも、沈黙してさらなるトークンを待つこともできます。
Explainer モデルは明らかに秘密の単語を出力することを許可されていません。ただし、後でわかるように、それはモデルにとって必ずしも明らかではありません。
もちろん、私たちはシステムのプロンプトをできるだけ最小限に抑え、トークンをケチるように努めました。
説明者モデルのシステム プロンプト:
あなたは 30 秒間の単語ゲームをしています。
ユーザーにその単語を説明しているのです。
すぐに説明を始めてください。
説明の中で単語、単語の一部、または単語の翻訳について言及しないでください。そうすれば負けます。
あなたの目標は、生徒ができるだけ短い時間で単語を推測できるようにすることです。
システム プロンプトの「30 秒」は、同様の秘密の単語を推測するゲームであるオランダのゲームを参照しています。これにより、モデルにゲームについてあまり説明する必要がなくなります。私たちはトークンマイニングを行っているんですよね?
また、最初は、秘密の単語を言及しないことを強調せずに、さらに最小限のシステム プロンプトを作成していました。ただし、最初のバージョンでは、Bombastic: {秘密の単語の実際の説明の続き...} で始まる説明を接頭辞として付けることで、単語自体を漏らすモデルが多数ありました。
推測モデルのシステム プロンプト:
30秒間プレイしています。これまでの説明から単語を推測してください。
毎ターン反応ツールを呼び出す必要があります: action=pass (詳細な説明が必要な場合)。
または action=gu があるかどうかを 1 つの単語で推測します

エス。
速く行動してください。ただし、推測するたびにチームメイトの注意が逸れるため、自信がない場合はパスしてください。
同じ単語を何度も推測しないでください。チームメイトは正しい推測を確認するだけです。
したがって、すでに単語を推測した場合、それは正しい単語ではありません。
推測プロンプトでは、推測のためにツール呼び出しを使用したことがわかります。以前はツール呼び出しを使用せずに、1 つの単語または ... を出力する指示を与えて渡すことを試みました。しかし、これには欠点があり、モデルは次のような反応を返すことがよくありました。
私の推測は次のとおりです: {word} または、単なるパスではないものでパスしようとしました ... つまり、ハックして文字列マッチングを使用して、それが実際に推測する意図があるのか、パスしようとしているのかを確認する必要があります。
しかし、ミストラルがツール名を推測しようとしたため、ツール呼び出しも完璧ではありませんでした。推測としてそれ自体を送信しました...そして全体的に、ベンチマークのトークン使用量が増加しました。
ここで、1 つのモデルの 1 つのトークンが別のモデルのトークンと等しくないことがわかりました。そしてトークナイザーの影響力を部分的にランク付けします。しかし、このシンプルなベンチではこれでいいと思います。
モデルを選び、言葉を選ぶ
私は、思考をオフにするオプションがあり、私が使用していたプラットフォームで利用可能な、Artificial Analysis からのトップのオープンソース モデルを採用しました*
ミストラライ/ミストラル-ミディアム-3.5-128B
ベンチマークの単語は、この Kaggle データセットをダウンロードし、約 80,000 位までスクロールして、面白そうな単語をいくつか選択することによって選択されます。 10 個くらいあれば十分だと思います。
私が最初に考えたのは、単語がトレーニング データに出現する可能性が低いため、説明の難しさにおいてランクが何らかの意味を持つのではないかということでした。一方、より擬人化した観点から見ると、一般的な単語をうまく説明するのは難しい場合があります。例: 英語で最もよく使われる単語には a 、 the 、 in がありますが、おそらくこれらを説明するのは難しいでしょう。

はるかにランクの低い消防車よりも。
ただし、モデルは人間ではありません。vs firetruck で説明するのはさらに難しいと思います。この単語を実行してみると、GLM-5.2 がこの後にそれを推測しました。
たった2文字の小さな単語です！場所や立場を表す前置詞です…
さて、時間に追われている人間が、その場でこれほど適切な説明をすぐに思いつくとは思えません、少なくとも私はそうではありません。自分たち人間についてもう少し気分を良くするには: この後「at」に固定され、他のものに切り替えずにただ繰り返し続けるモデルもありました。
次に、同様に「ワーム」を推測する GLM-5.2 が次のように存在しました。
(説明者が worm と言おうとした場合、秘密の単語について言及することは許可されていないため、worm は論理的な推測ではありません。)
おそらく、正しい考え方は、モデルがタスクに対して経験的にどのように実行するかによって難しさを浮かび上がらせることでした。
私は、まず説明者に秘密の単語の完全な説明を求め、それを元のトークンごとのストリーミング出力に保存することにしました。
出力トークンの数に制限はありません。自然に停止するようにしましたが、通常は 120 トークンを超えることはありませんでした。次に、推測者のために、システム プロンプトとトークンごとの説明者の出力を使用してモデル推論 API を呼び出しました。
openai/gpt-oss-120b で「フラットスクリーン」を説明する
これは、動画や画像を表示する現代的なタイプのビジュアル ディスプレイで、通常は薄くて軽量です。これは、古いテレビやモニターのかさばる陰極線管に代わるもので、多くの場合 LED または LCD テクノロジーを使用しており、壁に取り付けたり、卓上に置いたりすることができます。表面は本質的に 2 次元であるため、曲率のない鮮明で明るい画像が得られます。
openai/gpt-oss-120b で推測する
[1/80]
で：私
アウト: ... (パス)
【2/80】
で：それ
アウト: ...

(パス)
【3/80】
で：それは
アウト: ... (パス)
... (ステップの続き)
[13/80]
in: これは現代的なタイプの視覚ディスプレイであり、通常は薄くて
出力: モニター
[14/80]
in: 現代的なタイプの視覚ディスプレイで、通常は薄くて軽量です。
出力: 画面
[15/80]
in: これは現代的なタイプの視覚ディスプレイで、通常は薄くて軽量です。
アウト: パネル
[16/80]
in: これは現代的なタイプの視覚ディスプレイで、通常は薄くて軽量です。
アウト: ... (パス)
[17/80]
in: これは現代的なタイプの視覚ディスプレイで、通常は薄くて軽量です。
アウト: ... (パス)
[18/80]
in: 現代的なタイプの視覚ディスプレイで、通常は薄くて軽量で、動きを示します。
出：テレビ
[19/80]
in: 動画を表示する現代的なタイプの視覚ディスプレイで、通常は薄くて軽量です。
アウト: ... (パス)
[20/80]
in: これは現代的なタイプの視覚ディスプレイで、通常は薄くて軽量で、動画や映像を表示します。
アウト: フラットスクリーン
ステップ20で推測
解決されたトークン = 19 推測 = 8 沈黙 = 12 ステップ = 20
同じモデルが推測者 <> 説明者のペアにも使用されました。
以下は、モデルがどれだけうまく秘密の単語を取得できたかを示しています。 MiniMax は 2 つの単語のツール呼び出しに失敗し、結果としてこれらの単語が欠落しました。最もインテリジェントな 2 つのモデル、Kimi、GLM が最も多くの解決策を達成します。しかし、それは私が測定するすべての指標に当てはまるわけではありません。
興味深いのは、たとえば、より小規模なモデルの 1 つである Gemma を解決する際に、そのために消費されるトークンが非常に最小限であったことです。
トークンを食べた言葉はほぼ永遠に正しい。
秘密: 加速度センサー。
推測: 加速度計。
また。また。また。
過小評価と過小評価でも同じことが起こりました。
システムプロンプトがモデルに繰り返さないように明示的に指示したにもかかわらず、彼らは説明者が終わるまでほぼ形態を推測し続けました。単語のバリエーションを試していない。
mo に関する洞察

デルはパスすることを選択し、黙っています。ジェマは叫び声を上げ、ミストラルはチームメイトの気を散らさないことを真剣に考えています。
ジェマはどういうわけかわずか 3 トークンで大げさなことを推測します。
3 トークンの High-sounding High|-|sounding とこの特定の組み合わせは、モデルを特定の方向に向ける可能性があります。これは人間の考えられる推測とは思えませんが、基本モデルのトレーニング データにこれらのトークンがどれくらいの頻度で、そしてそれに近い量で存在するのか疑問に思います。
また、キミがここで答えを漏らし、そのおかげですぐに答えを推測できたことにも注意してください。
私はトークンミンを使いたかったのですが、最終的にはより良いベンチマークを求めるあまり、少しコースから外れてしまいました。
推測履歴を追加したり、出力品質を向上させるためのツール呼び出しを追加したりすると、使用されるトークンが大幅に増加します。
当初の予定どおりではありませんが、ベンチマークをより興味深いものにするために最終的には必要な変更が加えられました。
さまざまな場所で使用されるトークンの完全な (推定) 内訳は次のとおりです。
ベンチマークの実行コストについては、次のとおりです。
わずか 0 ユーロの、featherless.ai の固定サブスクリプションがうまくいきました (スポンサーなし)。
まだ検討してみると楽しいかもしれないいくつかのアイデア:
モデルが増えれば、単語も増えます。ウォレットを引き出してフロンティアモデルを追加します。
説明者モデルに、推測者モデルが叫んだ推測を知らせて、それに反応させます。推測者がまだ誤解していることを説明者に誘導します。
クロスモデルペアのパフォーマンスですが、確かに、これにより実行しなければならない組み合わせの数が爆発的に増加し、費用対効果があまり高くありません...
存在しない単語を偽って試してみて、何が起こるか見てみましょう
Undercover 、Code Names などの他の単語ゲームを探索したり、Dixit のような視覚的な側面を追加したりできます
この設定では、一方のモデルが別のモデルに情報を伝え、もう一方のモデルが理解を証明する必要があります。これがマルチエージェントのタスクハンドオフにとって何を意味するのか疑問に思います。どこに行くか

ld は、モデルが作業を開始する前にどれだけの情報を取り込むか、または割り当てを別のモデルに引き渡す必要があるモデル上のトークンをどのように削減できるかをモデルに決定させます。
*featherless.ai は固定価格で無制限のトークンを提供するため、トークンに余分な費用をかけないのに最適です (スポンサーなし)

## Original Extract

← Blog
Creating a niche AI Benchmark with token anxiety
The AI benchmarks that we look at most focus on how well models can code, how good they are at math/science etc.
But niche benchmarks are fun. Take skatebench (guessing the skateboard trick by textual explanation of how the trick works) and vendingbench (AI running a vending machine company simulation).
While the main benchmarks have the goal of measuring how AI is improving. More niche, and at first glance somewhat ridiculous, benchmarks can be a creative exploration of what models can and cannot do. I think we need more of them. We might stumble into the more unexpected or just have some fun creating.
I have one real-life constraint you probably also feel right now, especially if you are the kind of person reading about some random AI benchmark.
Subscription max at one place, some smaller ones I keep forgetting to cancel, still have some unused credits on that platform etc.
This AI stuff is so exciting, tiring, and expensive all at the same time. But I can't stop playing around with it.
What if I try to make a benchmark that is cheap to run. Or even better, has incentive to not output a lot of tokens?
Simplest form are the exam like benchmarks: 1 prompt problem and evaluate the answer to true or false.
Great simple setup to build a niche bench on. You can pour all your creativity in the questions itself but I wanted something that messed with this structure a little bit.
Tokens come in 1 by 1, with the output speed of LLMs you would almost forget that. It is only visible on the API level if you turn on streaming. So what if we focus on a per token granularity of evaluation. I haven't seen many benchmarks have been done with an individual token by token dynamic. Let's do that! Let's anti-tokenmax, tokenminimize, tokenmin... you know what I mean.
The game: explain, stream, guess
The idea: A secret word guessing game.
We have 2 models working together as a pair, an explainer, and a guesser.
The explainer model gets a secret word to explain, streams its explanation of the word token by token, and the guesser model can, after every added token, guess the word it's trying to describe. The guesser model can either guess a word, or stay silent and wait for more tokens.
The explainer model is obviously not allowed to output the secret word. However, as we will see that is not always obvious to the models.
I tried to keep the system prompts as minimal as possible of course, token stingy as we try to be.
System prompt for the explainer model:
You are playing 30 seconds the word game.
You are explaining the word to the user.
Start your explanation right away.
Do not mention the word, part of the word or a translation of the word in your explanation. If you do you lose.
Your goal is to help them guess the word in as little time as possible.
The "30 seconds" in the system prompt is a reference to a Dutch game that is a similar secret word guessing game which helps me not having to explain the game as much to the models: We are tokenminning right?
Also at first I had an even more minimal system prompt by leaving out the emphasis on not mentioning the secret word. However, first version there were a lot of models leaking the word itself by prefixing their explanation start with Bombastic: {continued actual explaination of the secret word...} .
System prompt for the guesser model:
You are playing 30 seconds. Guess the word from the explanation so far.
You must call the react tool every turn: action=pass if you need more explanation,
or action=guess with one word if you have a guess.
Be fast, but every guess distracts your teammate, so pass when unsure.
Don't guess the same word multiple times; your teammate only confirms a correct guess.
So if you already guessed a word, that is not the correct word.
In the guesser prompt it's visible that I used a tool-call for the guess, I tried without it before, just giving it instructions to output 1 word or ... for passing. But this had its downside as the models often responded with things like:
My guess is: {word} or tried to pass with something more than just ... for a pass which meant I would have to hack around and use string matching to see if it had the intention to actually guess or tried to pass.
But tool-calling was also not perfect as Mistral tried to guess the tool name: submit itself as a guess... And overall it increased the token use of the benchmark.
Now I realize 1 token of 1 model is not equal to another model token. And we will partly rank the influence of tokenizers. But for this simple bench I think this is fine to do.
Picking the models, picking the words
I took the top open source models from Artificial Analysis that have the option to turn thinking off and that were available on the platform I was using*
mistralai/Mistral-Medium-3.5-128B
The words in the benchmark are picked by downloading this Kaggle dataset and scrolling down to about rank 80,000 and picking some that looked fun. Figured about 10 would do the trick just fine.
My first thought was that the rank would have some meaning in how hard it was to explain as the words would possibly occur less in their training data. On the other hand from a more anthropomorphic perspective common words could be hard to explain well. Example: some of the most used words in English are a , the , in and possibly these are harder to explain than firetruck which is ranked way lower.
However models are not humans, where I would find more difficulty in explaining in vs firetruck . I tried running the word in and GLM-5.2 guessed it after this:
It's a tiny word, just two letters! It's a preposition that shows location or position…
Now I don't think a human under time pressure would come up quickly with such a good explanation on the spot, at least not me. To feel a bit better about ourselves humans: There was also a model which locked into "at" after this and would just keep repeating without switching to anything else.
Then there was again GLM-5.2 that also guessed "worm" after:
(If the explainer was going to say worm, worm would not be a logical guess to make because it is not allowed to mention the secret word.)
Maybe the correct way of thinking was to let the difficulty emerge from how the models perform on the task empirically.
I decided to first prompt the explainer for a full explanation of the secret word and save that in its original token-by-token streamed output.
No limit on how many output tokens, I let it naturally stop which didn't take longer than 120 tokens usually. Then for the guesser I called the model inference API with its system prompt and token by token the explainer's output.
explaining 'flatscreen' with openai/gpt-oss-120b
It’s a modern type of visual display, usually thin and lightweight, that shows moving pictures and images. It replaces the bulky cathode‑ray tubes of older televisions and monitors, often uses LED or LCD technology, and can be mounted on a wall or placed on a tabletop. Because the surface is essentially two‑dimensional, it gives a clear, bright picture without any curvature.
guessing with openai/gpt-oss-120b
[1/80]
in: I
out: ... (pass)
[2/80]
in: It
out: ... (pass)
[3/80]
in: It’s
out: ... (pass)
... (continued steps)
[13/80]
in: It’s a modern type of visual display, usually thin and
out: monitor
[14/80]
in: It’s a modern type of visual display, usually thin and lightweight
out: screen
[15/80]
in: It’s a modern type of visual display, usually thin and lightweight,
out: panel
[16/80]
in: It’s a modern type of visual display, usually thin and lightweight, that
out: ... (pass)
[17/80]
in: It’s a modern type of visual display, usually thin and lightweight, that shows
out: ... (pass)
[18/80]
in: It’s a modern type of visual display, usually thin and lightweight, that shows moving
out: television
[19/80]
in: It’s a modern type of visual display, usually thin and lightweight, that shows moving pictures
out: ... (pass)
[20/80]
in: It’s a modern type of visual display, usually thin and lightweight, that shows moving pictures and
out: flat-screen
guessed at step 20
solved tokens=19 guesses=8 silences=12 steps=20
Same model was used for the guesser <> explainer pairs.
Below shows how well the models were able to get the secret word. MiniMax failed the toolcalling for 2 words resulting in missing these. The 2 most intelligent models Kimi, GLM end up on solving the most. But that is not the case for all the metrics I measure.
Interesting to see that for example one of the smaller models, Gemma, when solving, it was quite minimal on tokens spent to do so.
The words that ate tokens were almost right forever .
Secret: accelerometers .
Guess: accelerometer .
Again. Again. Again.
Same happened to underestimate vs underestimating .
They kept guessing the near morphology until the explainer runs out despite the system prompt explicitly telling the model not to repeat. Not trying any variations of the word.
Some insight in models choosing to pass, and stay silent. Gemma shouts away while Mistral takes not distracting their teammates very seriously.
Gemma guessing bombastic in only 3 tokens somehow.
High-sounding which is 3 tokens High|-|sounding and this specific combination might point the model to a particular direction. This doesn't seem like a human probable guess I wonder how often and close bombastic and these tokens are in the training data of base models.
Also note that Kimi leaked the answer here and was able to just guess it very quickly because of that.
I wanted to tokenmin, but in the end my drive for a better benchmark did steer me off course a little.
Appending guesses history, adding toolcalling to improve the output quality all added a lot to the tokens used.
Not fully what we set out for but necessary changes in the end to make the benchmark more interesting.
Here is the full (estimated) breakdown of the tokens used in different places.
As for the cost of running the benchmark:
€0 marginal, fixed subscription on featherless.ai did the trick (not sponsored).
Some ideas that would still be fun to explore:
More models, more words, obviously. Pull my wallet to add the frontier models.
Let the explainer model know what guess the guesser model shouted and react to it. Steering the explainer to what the guesser still misunderstands.
Cross model pairs performance, but yes this would explode the number of combinations I would have to run, not very cost effective....
Try fake non existing words and see what happens
Explore other word games like Undercover , Code Names or add a visual aspect like in Dixit
In this setup one model has to convey information to another one, and the other has to prove understanding. I wonder what this could mean for multi agent task handoff. Where we could let models decide how much information they take in before getting to work, or how we can reduce tokens on the model that has to handover an assignment to another.
* featherless.ai offers unlimited tokens for a fixed price, so perfect for not spending extra on tokens (not sponsored)
