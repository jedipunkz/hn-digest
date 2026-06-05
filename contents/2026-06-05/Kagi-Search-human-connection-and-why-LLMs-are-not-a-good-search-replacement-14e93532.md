---
source: "https://www.alexselimov.com/posts/small_search_human_contact/"
hn_url: "https://news.ycombinator.com/item?id=48411015"
title: "Kagi Search, human connection, and why LLMs are not a good search replacement"
article_title: "Kagi Search, human connection, and why LLMs are not a good search replacement | Alex Selimov"
author: "aselimov3"
captured_at: "2026-06-05T13:09:43Z"
capture_tool: "hn-digest"
hn_id: 48411015
score: 3
comments: 0
posted_at: "2026-06-05T11:34:34Z"
tags:
  - hacker-news
  - translated
---

# Kagi Search, human connection, and why LLMs are not a good search replacement

- HN: [48411015](https://news.ycombinator.com/item?id=48411015)
- Source: [www.alexselimov.com](https://www.alexselimov.com/posts/small_search_human_contact/)
- Score: 3
- Comments: 0
- Posted: 2026-06-05T11:34:34Z

## Translation

タイトル: Kagi Search、人間関係、LLM が検索の代替として適していない理由
記事のタイトル: Kagi Search、人間関係、LLM が検索の代替として適していない理由 |アレックス・セリモフ
説明: おそらく 1 ～ 2 か月前に Kagi 検索に切り替えましたが、生活の質が大幅に向上したことに気づきました。
本当の答えにたどり着くまでにクリックスルーしなければならない AI の無駄な Web ページの量が大幅に減りました。
残念なことに、特定の検索では依然として AI スロップの概要ページが表示されますが、これは
[切り捨てられた]

記事本文:
Kagi Search、人間関係、LLM が検索の代替として適さない理由 |アレックス・セリモフ アレックス・セリモフ
Kagi Search、人間関係、LLM が検索の代替として適さない理由
おそらく 1 ～ 2 か月前に Kagi 検索に切り替えましたが、生活の質が大幅に向上したことに気づきました。
本当の答えにたどり着くまでにクリックスルーしなければならない AI の無駄な Web ページの量が大幅に減りました。
残念なことに、特定の検索では依然として AI スロップの概要ページが表示されますが、切り替えてからははるかに管理しやすくなりました。
私にとって最大の驚きは、その過程でできた友達です。
Kagi 検索により、さらに多くのブログベースのソリューションを利用できるようになりました
Kagi の技術的な問題の解決策を検索すると、さらに多くのブログにたどり着くことがわかりました。
最近、私は主に Rust/C++ のバックグラウンドを持って Zig を勉強しています。
他の言語と同様に、私はジグの「哲学」を確実に理解するために、さまざまな問題を解決するための慣用的なアプローチを探してきました。
Kagi のおかげで、実際の人間の苦闘や解決策を読むことができるブログが増えたような気がします。
最近の例の 1 つは、Zig での JSON、特に std.json.Parsed 型の周りでの解析に苦労したことです。
Karl Seguin 1 からの素晴らしい投稿は、Parsed 型についての良い概要とより深い洞察を提供してくれました。
最も役に立ったのは、彼も T の代わりに Parsed(T) を渡すことに同じ嫌悪感を持っているようで、まともな解決策を示してくれたことです。
これは、ChatGPT が推奨したものよりも優れたソリューションであり、本当に感謝しています。
ブログを通じた人のつながり
人間の実際のユニークな声や経験を読むことに加えて、ほとんどのブログ (私のブログを含む) には電子メールによる返信オプションがあることがわかりました。
ソースを保存しておけばよかったのですが、少し前にブログの記事を読みました。

良いブログ投稿を見つけたら電子メールで返信するように人々に指示しないでください。
非常に役立つ投稿を見つけたら、それを実行するように努めています。
ほとんどの場合、それは短い「X への投稿ありがとうございます。問題を解決するのに本当に役立ちました!」というだけです。
これを行うたびに、ブログ所有者から、自分のブログが役立ったことをうれしく思うという短い返信が届きます。
これは人間同士の接触のほんの小さなものですが、インターネットの荒野を明るくするほんの小さな前向きな相互作用であると私は感じています。
さて、私のおそらくより物議を醸す見解について説明します。
LLM が検索の代替として適切ではないと思う理由
私は LLM をほぼフルタイムで使用し、Agentic Engineering™️ のウサギの穴に挑戦した後、LLM に関するシンプルな哲学を洗練させました 2 。
私の考えでは、LLM の使用は現在の速度と長期的な成長の間の直接のトレードであり、依存し続けると気づかないうちに平均値まで下がってしまいます。
この取引は明らかに単純なものではありません。
速度を維持することが重要であるため、私は仕事で LLM とエージェント ツールを毎日使用しています。
(良くも悪くも) 問題を見つけるために、1 週間かけてオンラインで検索したりコードを 100 回読んだりできる時代は終わりました。
しかし、私はそのタスクに成長の可能性があり、費やす時間の予算がある瞬間を引き出すようにしています。
一般に、LLM を最も安全に使用するのは、必要なものがすでに正確にわかっていて、実装するか、欠けている情報のほんの一部を見つけ出すのにポンコツが必要なだけの場合だと思います。
Zig を学習するという観点から見ると、自分が何を望んでいるのかわからないので、LLM を使用することに満足していません。
特定の JSON ペイロードを解析する必要があることはわかっているかもしれませんが、それを実装する正しい方法はわかりません。
それを念頭に置くと、ChatGPT/Codex に「どうすればよいですか？」と尋ねることを躊躇します。

「この特定の JSON を解析してこの結果を取得します」の理由は次のとおりです。
LLM はユーザーを納得させるように設計されています。 LLM が基本的に誰にでも何でも納得させることができることを示す多くの研究 3 が存在します。
LLM の使用に関して強い意見を持っていない場合は、おそらく LLM ソリューションをそのまま受け入れるでしょう。問題があり、解決策を改良しようとしても、LLM やあなたの知識によって提供される解決策の領域に閉じ込められたままになります。
その結合された解空間は、必ずしも正しい解空間であるとは限りません。
だからこそ、実際の検索を行って実際のソースを見つけたほうが、潜在的により大きなソリューション領域を探索することになるため、最適なソリューションが得られる可能性が高いと私は考えています。
また、おしゃべりマシンによって自分の理解を誤魔化されてしまうという問題全体を回避できます。
ただし、これは、検索プロバイダーが、優れた広告を提供したり、LLM インタラクションを強制したりするのではなく、優れた検索を提供することに重点を置いている場合にのみ機能します 4 。
カギさん、ありがとうございました。
本当に気にしていないことをしなければならない状況はたくさんあります。
場合によっては、1 回限りのスクリプトをすばやく作成したり、悪質なファイル形式用のカスタム ファイル パーサーを作成したりする必要があるかもしれません。あるいは、たとえそれが最良の答えではなくても、何らかの答えが必要なだけかもしれません。
そのような状況では、私はためらうことなく LLM を使用します。
しかし、LLM を使用すると利便性のために成長が犠牲になるということは常に自分自身に正直です。
そして、この問題の特に厄介な点は、自分が何を諦めようとしているのかが分かっていないことです。
ある特定の問題の解決策を研究しているときに、別の問題に役立つ可能性のある革新的なアイデアや技術コンセプトに偶然出会うことがあります。
https://www.openmymind.net/Reading-A-Json-Config-In-Zig/ および https://www.openmymind.net/Zigs-std-json-Parsed/ ↩︎
これは私の降下と回復について説明したブログ投稿です。

AI ドゥーマー スパイラル: https://www.alexselimov.com/posts/let_no_crisis_go_to_waste/ ↩︎
理論的証明: https://arxiv.org/abs/2602.19141 および実際の例: https://www.science.org/doi/10.1126/sciadv.adw5578 。 ↩︎
Google の検索はデフォルトで AI モードになっています ↩︎
#ソフトウェア開発
#人生の最適化
#小さなウェブ

## Original Extract

I switched to Kagi search maybe 1 to 2 months ago, and I’ve noticed a significant life quality improvement.
The amount of AI slop web pages I have to click through before getting to a real answer has dropped significantly.
Sadly, for certain searches I still get the AI slop summary pages, but it’s a
[truncated]

Kagi Search, human connection, and why LLMs are not a good search replacement | Alex Selimov Alex Selimov
Kagi Search, human connection, and why LLMs are not a good search replacement
I switched to Kagi search maybe 1 to 2 months ago, and I’ve noticed a significant life quality improvement.
The amount of AI slop web pages I have to click through before getting to a real answer has dropped significantly.
Sadly, for certain searches I still get the AI slop summary pages, but it’s a lot more manageable since switching.
The biggest surprise for me has been the friends I made along the way.
Kagi search exposes me to a lot more blog based solutions
I’ve found that searching for solutions to technical problems in Kagi brings me to many more blogs.
Recently I’ve been learning Zig, coming from a mainly Rust/C++ background.
As with any language, I’ve been searching for the idiomatic approach to solving different problems to make sure I understand the Zig “philosophy.”
With Kagi I feel like I get taken to more blogs where I can read the struggles and solutions of real human beings.
One recent example was my struggle with parsing JSON in Zig, particularly around the std.json.Parsed type.
Nice posts from Karl Seguin 1 gave me a good overview as well as a deeper dive into the Parsed type.
Most helpful was that it seemed like he had the same distaste about passing around a Parsed(T) instead of T and gave a decent solution.
This was a better solution than what ChatGPT recommended which I really appreciated.
Human connection through blogs
In addition to reading the actual unique voice and experience of a human being, I found that most blogs (including mine) have a reply by email option.
I really wish I had saved the source, but some time ago I read a blog post telling people to respond via email when you see a good blog post.
I’ve been making an effort to do it when I find a post that has been very helpful.
Most times it’s just a short “Thank you for your post on X, it really helped me out with my issues!”
Every time I’ve done this, I get a short response from the blog owner mentioning that they are happy their blog was helpful.
It’s a pretty small bit of human contact, but I feel like it’s just a small positive interaction that brightens the internet wasteland.
Now for my maybe more controversial take.
Why I don’t think LLMs are a good search replacement
I’ve refined a simple philosophy around LLMs after using them pretty much full time and trying to go down the Agentic Engineering™️ rabbit hole 2 .
In my view, use of LLMs is a direct trade between current velocity and long term growth, with continual reliance dropping you down to the mean without you even realizing it.
This trade is obviously not straightforward.
I use LLMs and agentic tools daily in my job because keeping up velocity is important.
Gone are the days when you can spend 1 week searching online/reading your code for the 100th time to try and find the issue (for better or worse).
But I try to pull out the moments where the task offers growth potential and I have the time budget to spend.
In general, I think the safest use of LLMs is when you already know exactly what you want and you just need the clanker to implement or find the small tidbit of missing info.
In the context of learning Zig, I’m not satisfied using LLMs because I don’t know what I want.
I may know that I need to parse a specific JSON payload, but I don’t know what the right way to implement that is.
With that in mind, I hesitate to just ask ChatGPT/Codex “How do I parse this specific JSON to get this result” for the following reasons:
LLMs are designed to convince you. Many studies 3 exist that show an LLM can basically convince anyone of anything.
If you don’t have a strong opinion coming into LLM usage, you likely will just accept the LLM solution. Even if you have issues and try to refine the solution, you are still locked into the solution space provided by the LLM/your knowledge.
That combined solution space is not necessarily the correct solution space.
That’s why I think doing real search and finding real sources has a better chance of giving you an optimal solution, since you are exploring a potentially larger solution space.
You also avoid the whole issue of getting your understanding psyoped by the sycophancy machine.
This only works though if the search provider is focused on providing good search instead of providing good ads or forcing LLM interactions 4 .
Thank goodness for Kagi.
There are plenty of situations where I have to do something that I really don’t care about.
Maybe I need to write a one-off script quickly, create a custom file parser for an atrocious file format, or just need some answer even if it isn’t the best answer.
In those situations I use LLMs without hesitation.
But I’m always honest with myself that using LLMs is sacrificing growth for convenience.
And the particularly insidious part of this, is that you don’t know what you are giving up.
While researching solutions for one specific problem, you may accidentally come across a transformative idea or a technical concept that may help with another problem.
https://www.openmymind.net/Reading-A-Json-Config-In-Zig/ and https://www.openmymind.net/Zigs-std-json-Parsed/ ↩︎
Here’s a blog post describing my descent and recovery from an AI doomer spiral: https://www.alexselimov.com/posts/let_no_crisis_go_to_waste/ ↩︎
Theoretical proof: https://arxiv.org/abs/2602.19141 and real world example: https://www.science.org/doi/10.1126/sciadv.adw5578 . ↩︎
Google is defaulting to AI mode for search ↩︎
#software development
#life optimization
#small web
