---
source: "https://minor.gripe/posts/2026-07-13-the_ai_whalefall_and_open_source/"
hn_url: "https://news.ycombinator.com/item?id=48900231"
title: "The AI Whale Fall and Open Source"
article_title: "Minor Gripe - The AI Whale Fall and Open Source"
author: "ai_critic"
captured_at: "2026-07-13T23:43:30Z"
capture_tool: "hn-digest"
hn_id: 48900231
score: 1
comments: 0
posted_at: "2026-07-13T23:16:42Z"
tags:
  - hacker-news
  - translated
---

# The AI Whale Fall and Open Source

- HN: [48900231](https://news.ycombinator.com/item?id=48900231)
- Source: [minor.gripe](https://minor.gripe/posts/2026-07-13-the_ai_whalefall_and_open_source/)
- Score: 1
- Comments: 0
- Posted: 2026-07-13T23:16:42Z

## Translation

タイトル: AI クジラの秋とオープンソース
記事のタイトル: 小さな不満 - AI クジラの秋とオープンソース

記事本文:
マイナーな不満 - AI クジラの秋とオープンソース
軽度の不満
AI クジラの秋とオープンソース
オープンソースの人間中心主義に反対
Millwright: エージェントのエクスペリエンスからより賢明なツール選択
グレムリン: 思考と作動のアイデア
ヒューゴによって提供される |毒をテーマにした
© 2026 .無断転載を禁じます。
AI クジラの秋とオープンソース
フロンティア ラボは AI の使用に補助金を出していますが、それが維持されている間、オープンソースの状態を改善するために AI を利用しないのは愚かなことでしょう。
クジラはフロンティアラボに似ていますか?
クジラの秋は、クジラの死骸が沈んで海底に着地し、餌を提供し、新しい生命の誕生を促す生態系の開花です。
現在、商業化されている AI を批判する人たちは、何らかの方法で返済しなければならない巨額の資金 (場合によっては数千億ドル) を考えると、Anthropic や OpenAI のような研究所が存続できるはずがない、と言うでしょう。彼らは、これらの組織を支えている金融工学が市場原理から一時的に猶予を与えている、そしてそれが、その特権のために鼻からお金を払わずに粗利を作ることに時間を浪費できる唯一の理由だと言うでしょう。それが事実であるかどうかは完全にはわかりませんが、それが真実であると考えてみましょう。
個人的な経験だけでなく、多くの経験豊富な開発者の経験からも、AI はコードを作成したり保守したりするときに便利なツールとなり得ます。これについてはまた別の機会に詳しく書きますが、明らかにそこには存在があり、そうでないと言う人は別の現実に生きています。
AI は限られた期間だけ利用できるツールです。
さて、個人的には、期間限定のことはでたらめだと思いますが、ここでは修辞上の理由から許可します。
私がこれまでに話をしたことのある反LLMや反AIの人々のほとんどは、GLM 5のようなモデルの話題については奇妙なことに沈黙している。

.2 やその他の最新のオープンウェイト モデルは、現時点では補​​助として「十分」ですが、1 年後には現在の Fable や 5.6 と同じくらい優れている可能性があります。これは、彼らがラーパーであり、実際にテクノロジーを追跡していないためではないかと思います。
これらの同じ人々は、バブルがはじけた場合、AI/LLM はすべて自然に消滅し、宇宙の集合ストレージから自らを削除し、おそらくそれらを実行できる高価なグラフィックス カード、CPU、FPGA が後に続くと信じているようです。
そうそう、2008年のサブプライム崩壊後、誰もが家に住むのをやめたのはよく知られています。
クジラの死骸は海底に永遠に残るわけではありませんし、過剰に活用されているフロンティア研究所も同様です。しかし、それまでの間、残りの私たちにはそれらのトークンをきれいに拾えるチャンスがあります。
トークンの饗宴、もし私たちが食べることを選ぶなら
おそらく、古典的な XKCD イメージのいくつかのバージョンを見たことがあるでしょう。そこには、ソフトウェアのバベルの塔があり、それをすべてサポートしている重要なジェンガ部分が、人里離れたどこかの静かな開発者によって維持されています。そういったプロジェクトはたくさんありますが、正直なところ、私たちはこれ以上若くなっていません。
NixOS のような理論的には「若い」（この場合は明らかな誤りですが、ご容赦ください）プロジェクトであっても、最終的には数万件前半のオープン PR (問題は言うまでもありません!) が発生する可能性があります。これらの同じプロジェクトでも、利用可能な人材の不足について頻繁に苦情を言うようになります。
今のところ、AI を信頼できないものもあります。たとえば、他のプロジェクトの構成要素となるプロジェクトの大規模なアーキテクチャの書き換えや機能の追加は、雰囲気に任せるのが最善であるとは思いません。
機械的な作業 - バージョンの変更、失敗したテストの修正、ドキュメントのチェックなど

不一致については注意 – ポンコツはその価値を証明しています (LLM が導入される前でさえ、dependabot を覚えている人はいますか?)。これらのプロジェクトの技術的負債 (および開発者の人間工学的負債) の大部分は、これらのモデルを使用して改善できるものです。
NixOS の nixpkgs のようなプロジェクトでは、私たちはすでに自動化 (r-ryantm とその仲間たち) と機械的検証と CI を多用しています。これらの決定論的システムにより、テストが改善され、マイナーな PR リファクタリングなどの提案がさらに安全になります。これにより、ポンコツの貢献をさらに信頼しやすくなります。
私たちは皆、太陽が輝いているうちに干し草を作るべきです。
私が長年メンバーとして参加してきたもう 1 つのコミュニティ、Elixir コミュニティは、ここではやや最前線に立っていました。
Jose Valim の取り組み (たとえば、Andrew Kelley や彼の Zig への対応とは著しく対照的) や Elixir コミュニティの他の人々 (Chris McCord、Isaac 米本、Zach Daniels、私、その他) の取り組みを見ると、これらのツールを明らかに受け入れていることがわかります。
それは無批判に行われているわけではなく、ヴァリム自身もカンファレンス中にこれについてコメントしているが、それが実際に行われているということは、5年前でさえ想像できなかったものをはるかに超える成果を上げており、今後もその恩恵をもたらすことになるだろう。
クローボットからのランダムなドライブバイ PR は、一部のメンテナー (私たちが助けようと努めるべき) にとっては大きな懸念事項であり、ラーパー (彼らはボゾなので無視すべき) にとっては便利な隠れ蓑です。 「このままではPRを続けるのがやっとです！」 「ボットと議論したくないのです!」 「問題トラッカーは人間専用です!」など、イライラする可能性があることは否定しません。
もちろん、これを解決する方法は自動化されたプロセスと機械的なガードレールです。スタイルの自動適用、自動テスト、lint およびフォーマット設定はこれまでと同様です。コミュニティのレビュープロセスが耐えられない場合は、

ポンコツだったので、PR の統合を本気で望んでいる喧嘩好きな初投稿者たちの洪水に耐えられるものではありませんでした。
クジラの落下トークンを使用してその機械を構築できます。そうすることで、今後のポンコツなコラボレーションの管理が容易になります。そして、アーミッシュの足跡をたどることを決定したプロジェクト以外には、ある程度のポンコツが発生しない未来はありません（そして、アーミッシュですら、イギリス人のやり方の良い部分を選択的にゆっくりと採用するでしょう！）。
クジラの骨は何十年も保存され、秋にきれいに採取された後でも、日和見主義者の子孫にとって貴重なインフラとして機能します。私たちはクジラの落下を利用して将来を楽にするためのツールを構築する必要がありますが、それは外部のポンコツ投稿者のために解決しようとしなくても行うことができます。
この投稿について議論したいですか? |
メールで話し合う |
ハッカーニュース |
ロブスター |
クジラはフロンティアラボに似ていますか?
トークンの饗宴、もし私たちが食べることを選ぶなら

## Original Extract

Minor Gripe - The AI Whale Fall and Open Source
Minor Gripe
The AI Whale Fall and Open Source
Against Open Source's Anthropocentrism
Millwright: Smarter Tool Selection From Agent Experience
Gremlins: Ideas for Cogitation and Actuation
powered by Hugo | themed with poison
© 2026 . All rights reserved.
The AI Whale Fall and Open Source
Frontier labs are subsidizing AI usage, and we’d be fools not to use it to improve the state of open source while that holds out.
How are whales like frontier labs?
A whale fall is the bloom of an ecosystem when the carcass of a whale sinks and lands on the ocean floor, providing food and spurring the creation of new life.
Right now, critics of AI as currently being commercialized will tell you that labs like Anthropic and OpenAI cannot possibly last given the huge sums of money (hundreds of billions in some cases) that somehow must be paid back. They will say that the financial engineering underpinning these orgs has temporarily given a reprieve from market forces, and that is the only reason you can waste time making slop without paying through the nose for the privilege. I’m not entirely sure that that is the case, but let’s take it as true.
From personal experience as well as that of many experienced developers , AI can be a useful tool when writing or maintaining code. I’ll write more on this some other time, but there is very clearly a there there and anybody saying otherwise is living in some different reality.
AI is a tool that will only be available for a limited time.
Now, personally , the limited-time thing I think is bullshit–but, I’ll allow it for rhetorical reasons here.
Most-all of the people I’ve talked to that are strongly anti-LLM and anti-AI are oddly silent on the subject of models like GLM 5.2 and other modern open-weight models that are “good enough” for assistance now, and in another year will likely be as good as Fable or 5.6 today. I suspect this is because they are larpers and aren’t actually keeping track of the technology.
These same people seem to believe that AI/LLMs, if the bubble bursts, will all just spontaneously disappear and delete themselves from the collective storage of the universe, to presumably be followed by the expensive graphics cards and CPUs and FPGAs that can run them.
Ah yes, it is well-known that after the 2008 sub-prime collapse everybody stopped living in houses.
The whale carcass will not last forever on the ocean floor, nor too shall the over-leveraged frontier lab. But, in the meantime, there’s a chance for the rest of us to pick those tokens clean.
A feast of tokens, if we but choose to eat
You’ve probably seen some version of the classic XKCD image , where there’s a tower of babel of software and the critical jenga piece it all happens to be supported by is maintained by some quiet developer in the middle of nowhere. There are a lot of those projects, and honestly, we ain’t getting any younger.
Even in theoretically “young” (flagrantly false in this case, but bear with me) projects like NixOS, you can end up with open PRs (not to mention issues !) numbering in the low tens of thousands . These same projects will also then frequently complain about the lack of available manpower.
There are things that I wouldn’t trust to an AI right now. I don’t think, for example, that large architectural rewrites or feature additions for projects that are meant to be building blocks for other projects are best left to the vibes.
For mechanical work–things like bumping versions, fixing failing tests, checking documentation for inconsistencies–the clankers have proven their worth (even before LLMs, anybody remember dependabot?)! A large part of the technical debt (and debt in developer ergonomics) for these projects is something that could be remediated using these models.
In projects like NixOS’ nixpkgs, we already make heavy use of automation ( r-ryantm and friends) and mechanical verification and CI. Those deterministic systems make improving tests and suggesting minor PR refactors and things even safer . It makes it even easier to trust the contributions of clankers.
We should all be making hay while the sun shines.
Another community I’ve long been a member of–the Elixir community–has been somewhat at the forefront here.
If you look at work done by Jose Valim (in marked contrast to, say, Andrew Kelley and how he’s handling Zig) and other folks in the Elixir community (Chris McCord, Isaac Yonemoto, Zach Daniels, myself, others), there’s a clear embrace of these tools.
It isn’t done uncritically– Valim himself commented on this during a conference –but that it is being done at all has yielded and will yield dividends far beyond what we could’ve imagined even five years ago.
Random drive-by PRs from clawbots are a real concern for some maintainers (who we should try to help) and are convenient cover for the larpers (who are bozos and should be ignored). “We can barely keep on top of the PRs as it is!” “We don’t want to argue with a bot!” “Issue trackers are for humans only!”, etc. I won’t deny that can be frustrating.
The way to fix that, of course, is automated processes and mechanical guardrails. Automatic enforcement of style, automatic testing and linting and formatting–same as ever. If your community review process can’t stand up to a clanker, it wasn’t going to stand up to a flood of quarrelsome first-time contributors that really wanted their PRs merged.
We can use the whale fall tokens to build that machinery . Doing so makes it easier to manage the clanker collaboration down the road–and there’s no future where some degree of clanking does not occur, outside of projects that have decided to follow in the footsteps of the Amish (and even the Amish will selectively, slowly, adopt some good bits of what the English do!).
Whale bones last for decades and even after the fall has been picked clean they function as valuable infrastructure for the opportunists’ descendants. We need to use the whale fall to build the tools to make our futures easier, and that can be done even without trying to solve for external clanker contributors.
Want to discuss this post? |
Discuss via email |
Hacker News |
Lobsters |
How are whales like frontier labs?
A feast of tokens, if we but choose to eat
