---
source: "https://louwrentius.com/the-ai-mirage-or-why-i-think-the-hype-cant-sustain-itself.html"
hn_url: "https://news.ycombinator.com/item?id=48737847"
title: "The AI Mirage or Why I Think the Hype Can't Sustain Itself"
article_title: "The AI mirage or why I think the hype can't sustain itself"
author: "louwrentius"
captured_at: "2026-06-30T19:33:19Z"
capture_tool: "hn-digest"
hn_id: 48737847
score: 4
comments: 1
posted_at: "2026-06-30T19:14:04Z"
tags:
  - hacker-news
  - translated
---

# The AI Mirage or Why I Think the Hype Can't Sustain Itself

- HN: [48737847](https://news.ycombinator.com/item?id=48737847)
- Source: [louwrentius.com](https://louwrentius.com/the-ai-mirage-or-why-i-think-the-hype-cant-sustain-itself.html)
- Score: 4
- Comments: 1
- Posted: 2026-06-30T19:14:04Z

## Translation

タイトル: AI の蜃気楼、あるいは誇大広告が持続できないと私が考える理由
記事のタイトル: AI の蜃気楼、あるいは誇大広告が持続できないと私が考える理由

記事本文:
AI の蜃気楼、あるいは誇大広告が持続できないと私が考える理由
十分にズームアウトすると、すべてがブラック ボックスになり、内部の仕組みはわかりませんが、何かを入れて出力を観察することで、ブラック ボックスについて知ることができます。
大規模な言語モデルまたは LLM がそのようなブラック ボックスであると仮定してみましょう。 LLM を観察すると、とりわけ、出力が 99% 1 の確率で「正しい」ことがわかります。
これは非常に基本的かつ重要な観察です。コンピュータはその正確さと信頼性で知られています。入力 A で関数を実行すると、何があっても出力 X が得られることがわかっています。はい、パケットが失われたり、メモリが破損したりする可能性がありますが、それらは予測可能な方法で問題が発生します。パケットをチェックして再送信し、ECC メモリを使用します。また、データは、パケットの損失やメモリの破損を最初から検出できるような方法でフォーマットされています。私たちの世界全体がそれに依存しています。
システムが信頼できないと想像してください。 1+1 が常に 2 になるとは限りません。99% の確率でのみです。このようなシステムにはどれほどの価値があるでしょうか?それはおそらく状況によって異なりますが、確かにわかっていることが 1 つあります。それは、出力は信頼できない、確認する必要があるということです。方法は関係ありませんが、正しいかどうかをチェックする必要があり、それには人が必要です。
自動運転車でもそれが見られます。何が可能なのかは本当に印象的です。しかし、それらは真の自動運転ではありません。 AI が避けられない間違いを犯したときに介入できるように、人はハンドルに座って、あたかも自分で運転するかのように交通状況に注意を集中し続ける必要があります。
この記事は自動運転車に関するものではありませんが、人間には積極的に取り組んでいないと気が散って退屈してしまう性質があることを私たちは知っています。私たち人間が自分で運転し続けるか、100% の信頼性が必要なので、

ハンドルを握って本を読みながら車が自動運転します。 100% だけで十分であり、99% では不十分です。 1% の確率で問題が発生するため、それでも車を「運転」する必要がある場合、実際に何を解決したのでしょうか?
LLM の場合、LLM の出力を人にチェックさせることで、どのくらい時間が節約されますか? AI ベンダーが現在請求している補助金付きのコストではなく、この節約された時間で LLM の実際の運用コストを正当化できるでしょうか?
コードを記述する場合、LLM はおそらく、100 人のエンジニアのチームが 1 年で検証できるより多くの機能を 1 週間で作成できます。したがって、LLM が実際にどれほど高速であっても、人は回避できないボトルネックです。
つまり、正確さ、品質、安定性などを重視する場合です。しかし、そうしないのであれば、LLM が関与しているかどうかに関係なく、そもそもなぜ何かをする必要があるのでしょうか。
これが、AI の誇大宣伝が彼らの不条理な評価額が示唆するような高水準の約束を実現するのは不可能だと私が考える理由です。 AI に価値がないかもしれないと言っているわけではありません 2 が、おそらく人々が信じているよりも桁違いに価値が低いのです。
これはオリジナルのアイデアではないことは認めますが、人々はこのアイデアをさまざまな、おそらくより簡潔な方法で表明しています。それでも、繰り返す価値はあると思います。
また、なぜ組織が LLM をプロセスの不可欠な部分にしているのに、モデルが変更され微調整され、その結果、非常に予測不可能で異なる出力が生成されることが判明するのか、私には理解できません。
LLMに光がぴったり当たって目を細めると、それが仮想通貨の形に見えることがあります。少なくとも私にはそう見えます。
この数値は何にも基づいていないため、実際にはおそらくさらに悪くなる可能性があり、100% ではないことに注意することを除いて、実際の値は重要ではありません。 ↩
私はエネルギーの「無駄」や汚染を無視しています。

知的財産の盗難、著作権侵害、AIによる自傷行為。リストはどんどん増えていきます。 ↩
ストレージメトリクス用の Grafana ダッシュボード

## Original Extract

The AI Mirage or Why I Think the Hype Can't Sustain Itself
If you zoom out enough, everything is a black box where we don't know about the inner workings, but we can learn about the black box by putting stuff in and observing the output.
Let's pretend that a large language model or LLM is such a black box. If we observe a LLM, we learn - amongst other things - that the output is 'correct' 99% 1 of the time.
This is a very fundamental and important observation. Computers are known for their correctness, for their reliability. We know that if we run a function with input A, we get output X, no matter what. Yes, packets can be lost, memory can go corrupt but those things go wrong in a predictable way. We check and retransmit packets, we use ECC memory. And data is formatted in such a way that we can detect lost packets or corrupted memory in the first place. Our entire world relies on it.
Imagine that a system can't be trusted. That 1+1 isn't always 2. Only 99% of the time. How valuable can such a system be? That probably depends on the circumstances, but we know one thing for sure: we can't trust the output, it must be checked. It doesn't matter how, but it must be checked for correctness and that requires a person.
We see it with self-driving cars. It's really impressive what is possible. But they aren't true self-driving. A person needs to be sitting at the wheel, keeping their attention focussed on traffic as if they would be driving themselves, so they can intervene when the AI makes an inevitable mistake.
Although this post isn't about self-driving cars, we know that humans have a tenancy to get distracted and bored if we aren't actively engaged. We can argue that either we people keep driving ourselves, or we need 100% reliability, so we can remove the steering wheel and read a book while the car drives itself. Only 100% is good-enough, 99% doesn't cut it. What did we actually solve if I'm still required to 'drive' the car because of the 1% chance things go wrong?
In the case of an LLM, how much time is saved by letting a person check the output of an LLM? Can that saved time justify the actual operating cost of the LLM, as opposed to the subsidized cost AI vendors are currently charging?
In case of writing code, an LLM can probably create more functionality and features in a week a team of 100 engineers can validate in a year. So no matter what how fast the LLMs really are, the people are the bottleneck we can't circumvent.
That is, if we care about correctness, about quality, about stability and so on. But if we don't, why do something in the first place, regardless of an LLM being involved or not.
So this is why I think it's impossible for the AI hype to come through on the sky high promises their absurd valuations suggest. I'm not saying that AI may not be valuable 2 but probably orders of magnitude less valuable than people want us to believe.
I admit, this isn't an original idea, people have voiced this idea in different, maybe more succinct ways . Yet I think it is worth repeating.
I also don't understand why organizations would make LLMs a very integral part of their processes, only to discover that models are changed and tweaked, resulting in wildly unpredictable and different output.
Sometimes, when the light hits an LLM just right and you squint your eyes, it takes the shape of a crypto currency. At least that's what I see.
the number is based on nothing, in practice it's probably worse and the actual value is not important except to note that it's not 100%. ↩
I'm ignoring the energy 'waste', pollution, the IP theft, the copyright infringement, the AI-induced self-harm. On and on the list goes. ↩
Grafana Dasboard for storage metrics
