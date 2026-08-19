---
source: "https://ismaelroblesrazzaq.github.io/blog/llm-self-portraits/"
hn_url: "https://news.ycombinator.com/item?id=49362780"
title: "I gave frontier LLMs a canvas and told them to draw a self-portrait"
article_title: "I gave frontier LLMs a canvas and told them to draw a self-portrait"
image: "https://ismaelroblesrazzaq.github.io/blog/llm-self-portraits/media/og-card.png"
author: "ismael_rr"
captured_at: "2026-08-19T15:21:55Z"
capture_tool: "hn-digest"
hn_id: 49362780
score: 1
comments: 0
posted_at: "2026-08-19T15:19:22Z"
tags:
  - hacker-news
  - translated
---

# I gave frontier LLMs a canvas and told them to draw a self-portrait

- HN: [49362780](https://news.ycombinator.com/item?id=49362780)
- Source: [ismaelroblesrazzaq.github.io](https://ismaelroblesrazzaq.github.io/blog/llm-self-portraits/)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T15:19:22Z

## Translation

タイトル: フロンティアLLMにキャンバスを渡して自画像を描くように言いました
説明: フロンティア モデル 10 個、色鉛筆セット、各 200 ターン。彼らが何を描いたのか、どのように描いたのか、そしてその費用はいくらだったのか。

記事本文:
私はフロンティアLLMたちにキャンバスを渡し、自画像を描くように言いました。
以下は、Fable 5、GPT 5.6 Sol、Kimi K3、および Qwen 3.8 Max のセルフポートレートです (順不同)。
どれがどれだかわかりますか？
面白いですよね！ 『ザ・ファブル5』は私にとって非常に暗いと感じました。また、注目に値するのは、レオポルド・アシェンブレナーにとってこのような波瀾万丈な時期に描かれた 5.6 ソルの絵が、レオポルド・アシェンブレナーに奇妙に似ていることである。
このブログ投稿では、このベンチマークの動機、結果、仕組みについて説明します。以下のモデルで実行しました。
すべてのモデルは、可能な最大の努力を使用するように設定されました 1 。 OpenAI と Anthropic モデルはそれぞれの API を介して呼び出し、Qwen 3.8 Max は OpenRouter を介して Alibaba に呼び出し、Kimi K3 も Openrouter を介して Moonshot に呼び出しました 2 。
モナリザと星月夜を再現するフロンティア LLM のベンチマークについての tryai のブログ投稿を読んだ後、LLM がアート ツールを使用することに非常に興味を持ちました。 LLM に人間と同じようにアート ツールを扱える能力を与えることは、モデルに自転車に乗っているペリカンの SVG を盲目的に生成させるよりも、創造的能力の適切なベンチマークであることは明らかです。さらに、これは OOD タスクである可能性が高くなります。
このベンチマークの目的は 2 つあります。
フロンティア LLM 描画ベンチマークを完全に自由なクリエイティブ タスクに拡張する
モデルが自分自身をどのように描写するかを真に見るため 3
そのため、私のベンチマークは、図面の「芸術性」と「肖像画らしさ」の純粋に主観的な尺度です。本物のアーティストによる図面に見えるほど、ランクが高くなります。
各モデルの描画プロセスを 512x512 の各キャンバスで再生するビデオとして表示します。以下は、各図面の完全な転写へのリンク、図面に関する独自の説明、10 点満点の私の評価、および総コストです。
私が一番印象に残ったのは『寓話』でした。それは th でした

このモデルは、「実際の自分に正直に感じられる」ため、意図的に自分自身をより抽象的に描写する唯一のモデルです。ここで内部対話を見ることができます。その画力の高さにも驚かされました。 1 回のツール呼び出しで多くのブラーを一貫して正確に適用しますが、これは他のモデルでは困難です。
また、5.6 Terra にも感銘を受けました。なぜなら、その描画が最も様式的に感じられるからです。アーティストが実際にこの方法で自分自身を描いていると想像できるからです。
もう一つ注目したいのは、Opus 5 のプロセスです。不気味で少し恐ろしい肖像画を描くことに加えて、最も多くのターンを費やしたモデルでした 4 。図面の一部を消したり、やり直したりすることが最も多かった。 Fable がそうしなかったのは、最初に描きたいものを正しく描くことができ、やり直しの必要がなかったからではないかと推測しています。
Qwen 3.8 Max はまあまあでしたが、Kimi K3 にはがっかりしました。すべての古いモデル (GPT 5、GPT 5.4、Haiku 4.5) のパフォーマンスは非常に低かったため、過去 6 か月間であっても、視覚的な理解、推論、およびツールの呼び出しがどれほど向上したかがわかります。
全体的には楽しい練習だったと思います。モデルの機能が向上していることは明らかですが、依然として本物の芸術からは程遠いです。注意すべき点の 1 つは、ハーネスがかなり初歩的なものであるということです。より優れたハーネスを作成するために多大な努力を注げば、より良い出力が得られる可能性があると思います。
Tryai の絵画再構築ベンチマークが提供する環境から始めました。
ただし、出力はまったく図面のように見えません。何度か繰り返した結果、満足のいくハーネスが完成しました (GitHub: github.com/iroblesrazzaq/canvas-arena )。 GPT 5.6 Luna を使用したハーネスの使用前と使用後は次のとおりです。
モデルには次のツールがあります。
目標は、十分に引き分けを作成できるように十分なツールをエージェントに提供することです。

キャンバスをリアルに見せます。リアルなキャンバスの場合、顔料を暗くなるように設定し、鉛筆が実際に描く様子をシミュレートする不完全性を追加し、ピクセルの「高さ」と適用される圧力の関数として鉛筆が植物の繊維にマークを付ける方法をシミュレートするフィールドを追加しました。つまり、繊維が多いとより多くの色素を捕らえることができ、色が濃くなります。
描画を完了するために各モデルに 200 ターンを与えました。ターンは、モデルからの 1 つの応答として定義されます。それは、絵を描いたり、見たり、単に考えたりすることです。また、モデルは同じツール呼び出しで複数のドローまたはスマッジを書き込むことができます。
タスクをよく理解できなかった GPT 5 を除いて、すべてのモデルは 200 ターン未満で終了しました。実際、オーパス 5 の 191 ターンを除き、すべてのモデルの使用ターン数は 100 未満です。
制約を指定します - 200 ターン、完了したら終了を呼び出します
モデルに何を描くかを決して指示しないでください
適切な練習の 1 つは、OpenAI の SOTA が時間の経過とともに改善されるのを観察することです。 OpenAI は、2025 年 8 月に GPT 5、2026 年 3 月に GPT 5.4、2026 年 7 月に GPT 5.6 をリリースしました。以下は 3 つすべての最終バージョンであり、全体を通して顕著な進歩が見られます。無料の API クレジットにアクセスできるため、OpenAI モデルを使用してこの分析を実行しました。
もう 1 つの興味深い軸はコストです。各実行のコストの内訳は次のとおりです。
そして、私の非常に主観的な評価で測定されたパレート図は次のとおりです。
芸術における AI の役割について議論しないとしたら、私は怠慢になるでしょう。このベンチマークにもかかわらず、私は次のように強く信じています。
クリエイティブな世界には AI の居場所はありません。アートにおいて、アーティストは他者との対話を生み出します。
鑑賞者、アートコミュニティ、さらにはアートを通じて自分たち自身をも保護する重要な要素です。
AI は代替できませんし、代替すべきではありません。
さらに、芸術が自己完結したものである、つまり芸術を創造することが芸術を創造する主な目的の一つであると認識することが重要だと思います。 AIは違う

家賃。その目的は、生産性の向上や人間性の向上（あるいは株主価値の向上）です。したがって、まさに人間的な芸術の世界には AI の居場所はありません。
AI バブルから一度だけ逃れたいと考えているベイエリアの皆さん、
SF MoMA のマティス展にぜひ足を運んでみてください。とても楽しかったので、9月13日までです。
ただし、人体モデルには Ultracode の代わりに max を使用しました ↩
Moonshot に切り替える前に、名前は伏せますが、信頼できないプロバイダーで K3 を試しました。ベーステンでした。 ↩
クロードではなく、私自身の「本当に」の使い方です↩
GPT 5 以外はバカすぎた ↩

## Original Extract

Ten frontier models, a set of colored pencils, and 200 turns each. What they drew, how they drew it, and what it cost.

I gave frontier LLMs a canvas and told them to draw a self-portrait.
Here are the self-portraits for Fable 5, GPT 5.6 Sol, Kimi K3, and Qwen 3.8 Max - out of order.
Can you guess which one is which?
Interesting, right!? The Fable 5 one felt very claudey to me. Also notable is the uncanny resemblance of the 5.6 Sol drawing to Leopold Ashenbrenner at such an eventful time for him.
In this blog post, I'll lay out the motivation, results, and machinery of this benchmark. I ran it on the following models:
All models were set to use max effort possible 1 . I called OpenAI and Anthropic models through their respective APIs, Qwen 3.8 Max to Alibaba through OpenRouter, and Kimi K3 to Moonshot through Openrouter as well 2 .
After reading this blog post from tryai about benchmarking frontier LLMs on recreating the Mona Lisa and Starry Night, I became very curious about LLM's using art tools. It seems obvious that giving an LLM the ability to wield art tools like a human is a more apt benchmark of creative ability than having a model blindly generate an SVG of a pelican riding a bike. Furthermore, this is more likely an OOD task.
The purpose of this benchmark is twofold:
to extend the frontier LLM drawing benchmark to a fully open-ended creative task
to genuinely see how models choose to depict themselves 3
To that end, my benchmark is a purely subjective measure of the "artness" and "portraitness" of the drawings - the more it looks like a drawing by a real artist, the higher it'll rank.
Here are the drawing processes for each model, displayed as a video playing through each 512x512 canvas. Below are the link to the full transcript of each drawing, their own descriptions of the drawing, my rating out of 10, and the total cost.
I was most impressed by Fable. It was the only model to deliberately depict itself more abstractly because it "feels honest to what I actually am". You can see its internal dialogue here . Its drawing prowess also surprised me. It consistently and accurately applies many blurs in a single tool call, something that other models struggle with.
I was also impressed by 5.6 Terra because its drawing feels the most stylistic - you can imagine an artist actually depicting themselves this way.
Another thing I wanna note is Opus 5's process. Besides drawing an uncanny, slightly terrifying portrait, it was the model that took the most turns 4 . It erased and redid parts of its drawings the most. I speculate that the reason Fable didn't do this is because it was able to correctly draw what it wanted to draw the first time, eliminating any need for re-doing.
Qwen 3.8 Max was ok, and I was disappointed by Kimi K3. All older models - GPT 5, GPT 5.4, Haiku 4.5 - did quite poorly, which shows how much visual understanding, reasoning, and tool calling have improved even in the last 6 months.
I think overall this was a fun exercise. It's clear that the models are improving in capability, yet are still are a far cry from real art. One thing to note is that the harness is fairly rudimentary - I think that putting significant effort into creating a better harness could result in much better outputs.
I started with the environment provided by tryai's painting reconstruction benchmark.
However, the outputs don't look like drawings at all. After several iterations, I ended up with a harness I'm happy with (GitHub: github.com/iroblesrazzaq/canvas-arena ). Here's the before and after of the harness using GPT 5.6 Luna:
The model has the following tools:
The goal is to give the agent enough tools to be able to sufficiently create a drawing and make the canvas look realistic. For the realistic canvas, I set pigment to only darken, added imperfections that simulate how pencils actually draw, and added a field that simulates how pencils mark plant fiber as a function of how "high" the pixel is and how much pressure is applied. That is, a high fiber catches more pigment --> is darker.
I gave each model 200 turns to finish their drawing. A turn is defined as one response from the model - that can be drawing, looking, or just thinking. Also, the model can write multiple draws or smudges in the same tool call.
Every model finished in under 200 turns except for GPT 5, which failed to understand the task well. In fact, every model used less than 100 turns except Opus 5, with 191 turns.
state the constraints - 200 turns, call finish when done
never tell the model what to draw
One neat exercise is to observe OpenAI's SOTA improve through time. OpenAI released GPT 5 in August 2025, GPT 5.4 in March 2026, and GPT 5.6 in July 2026. Below are the final versions of all three, in which we can see marked progress throughout. I ran this analysis with OpenAI models because I have access to free API credits with them.
Another interesting axis is cost. Here is the cost breakdown for each of the runs:
And the pareto chart measured with my very subjective rating:
I would be remiss if I didn't discuss the role of AI in art. Despite this benchmark, I firmly believe that
AI has no place in the creative world. In art, artists create a dialogue with
the viewer, the art community, and even themselves through their art - a critical component that
AI cannot and should not replace.
Moreover, I think it's important to recognize art as self-sufficient - that is, creating art is one of the principal purposes of creating art. AI is different. Its end is to increase productivity or improve humanity (or increase shareholder value). Thus, AI has no place in the very human world of art.
For those of you in the bay area looking to escape the AI bubble for once,
I highly recommend you visit the Matisse exhibit at SF MoMA! I enjoyed it a lot, and it's up through September 13.
although I used max instead of ultracode for anthropic models ↩
I tried K3 with an unreliable provider who I won't name before switching to Moonshot. It was Baseten. ↩
my own use of 'genuinely', not claude ↩
besides GPT 5, which was too dumb ↩
