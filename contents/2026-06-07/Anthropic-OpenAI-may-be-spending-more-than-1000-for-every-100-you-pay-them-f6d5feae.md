---
source: "https://ea.rna.nl/2026/06/07/anthropic-openai-may-be-spending-more-than-1000-for-every-100-you-pay-them/"
hn_url: "https://news.ycombinator.com/item?id=48434342"
title: "Anthropic/OpenAI may be spending more than $1000 for every $100 you pay them"
article_title: "Anthropic/OpenAI may be spending more than $1000 for every $100 you pay them – R&A IT Strategy & Architecture"
author: "gctwnl"
captured_at: "2026-06-07T15:35:55Z"
capture_tool: "hn-digest"
hn_id: 48434342
score: 51
comments: 63
posted_at: "2026-06-07T12:54:54Z"
tags:
  - hacker-news
  - translated
---

# Anthropic/OpenAI may be spending more than $1000 for every $100 you pay them

- HN: [48434342](https://news.ycombinator.com/item?id=48434342)
- Source: [ea.rna.nl](https://ea.rna.nl/2026/06/07/anthropic-openai-may-be-spending-more-than-1000-for-every-100-you-pay-them/)
- Score: 51
- Comments: 63
- Posted: 2026-06-07T12:54:54Z

## Translation

タイトル: Anthropic/OpenAI は、あなたが支払う 100 ドルごとに 1000 ドル以上を費やしている可能性があります
記事タイトル: Anthropic/OpenAI は、あなたが支払う 100 ドルごとに 1000 ドル以上を費やしている可能性があります – R&A IT 戦略とアーキテクチャ
説明: LLM (Claude Code、OpenAI Codex) を使用したコーディングは、生成 AI の「キラー アプリ」として提示されることがよくあります。しかし、データを見ると、欠けているパズルの 1 ピースは実際のコストであるようです。何が起こっているのかについて、より曖昧な全体像を得ようとする探究の結果、驚くべき結果が得られました。

記事本文:
コンテンツにスキップ
R&AのIT戦略とアーキテクチャ
技術的および非技術的な「デジタル」に関するあらゆるものに関する洞察
書籍とコレクション
「ChatGPTと仲間たち」コレクション
「情報革命について」コレクション
ArchiMate のマスタリング – エディション 3.2
複数の言語で提供される無料の ArchiMate 3.2 概要 PDF
EA チェス — ザ・ブック
チェスと EA の技術 — 概要
記事コレクション
「ChatGPT と仲間たち」コレクション
メニュー
検索ボックスを開く
検索ボックスを閉じる
検索:
検索
Anthropic/OpenAI は、あなたが支払う 100 ドルごとに 1000 ドル以上を費やす可能性があります
理由は隠されたままになりますが、15 か月の休止期間を経て、ジェネレーティブ AI/LLM についての執筆を再開します (2025 年 10 月の記事と 2025 年 6 月の記事は本格的な記事とはみなされません)。今日は、LLM を使用したコーディングが LLM の「キラー アプリ」として位置付けられているため、「大規模な「言語」モデルを使用したコーディング」に関する 2 つの記事のうちの最初の記事です。
Anthropic が最近リリースしたブログ投稿「AI が自らを構築するとき」についての短い余談のため、このプログラムを中断します。
Anthropic はおそらく Google のマーケティング担当者を雇ったのでしょうか?
Anthropic のブログ投稿は、示唆に富む文章のマスタークラスです。注意事項はありますが、隠されているか、より誇張的な記述の間に挟まれています。 「私たちは間違っている可能性があります」という文がありますが、それらが間違っていないとして、数千語のテキストの中の 1 つの文としてどのような役割があるのでしょうか?ベンチマークは疑わしい (人間と比較したコーディング タスクの成功率が 50% または 80% であっても、フル エージェント (ループに人間が関与しない) コーディングでは実質的にまったく役に立たない。1 日に 8 倍のコード行をチェックインするのは本当に良いことなのだろうか? 前日に OK でなかったものを毎日置き換えているとしたらどうなるだろうか? コード行の信頼性が低くなるような方法で LLM が編集されたらどうなるだろうか?

すでにこのままの措置ですか？全体として、これは Google の「Willow」QM コンピューティング チップに関する欺瞞的な話を思い出させます。
ちなみに、「1 日あたり 8 倍のコード行をチェックインするのは本当に良いことなのでしょうか?」については、私の LLM ベースのチェックインの多くはそのようなもので、率直に言って、クロード コードが失われた場合にバックトラックできるようにするためだけにチェックインの数を増やしました。中間チェックインなしですべての作業を行ったにもかかわらず、コード行数で言えば、最終的に行った変更の 7 倍の変更をコミットしました。したがって、生産性の向上というよりも、オーバーヘッドの増加を見積もる「8」は、ほぼ正しいように思えます…
これについては、後ほど「クロードによるコーディング」について詳しく説明する場合に戻ります。
そしてその間に。これは長くて曲がりくねった文章です (この点で本当に苦労しています、申し訳ありません)。そのため、以下を提供しましょう。
TL;DR — LLM コーディングが手頃な価格になるとは思えません
（ましてや「AIが自ら構築する」なんてことは）
私はいくつかの実験を行ってきました。実験は「そもそもクロード コードはどれくらい優れているのか?」というものです。その実験は現在も進行中であり、Claude Code は現在、約 40,000 行のコードと、(不完全ではあるが) 動作するアプリケーションを作成しています。その経験については近いうちにレポートしたいと思います（ただし、それははるかに難しい記事です）。その間、私はコストの問題を経験し、それが短期間の研究プロジェクトにつながり、多くの興味深い観察と結論につながりました。
重要な観察から始めましょう。Claude Code と私自身の (錆びてはいますが、十分にしっかりした) プログラミングの背景の組み合わせのおかげで、このようなアプリケーションでは作成できなかったこのアプリケーション (現時点では未完成ですが機能します) を Claude Code に作成させることができました。

短い時間ですが、時間とエネルギーを考えれば、それは「まったく」ということを意味するでしょう。経験豊富なプログラマにとって、最初の経験は非常に印象的です。なぜなら、経験豊富なプログラマは、そのようなコードを作成するのに通常どの程度の理解を自分自身で行う必要があるかを知っているからです。
しかし… LLM コーディングは、ほとんどの用途において経済的に実行可能ではありません。サブスクリプションには多額の補助金が支払われているため、現在も実行可能です。しかし、月額 100 ドルの Claude Max プランを使用し、完全な「エージェント コーディング」（つまり人間が関与する人がほとんどいない）で週の制限まで使用すると、API 価格で 1000 ドルを超える量のトークンを使用することになります。 Anthropic はその出血を止めるのに忙しいようです (Opus 4.7、4.8)。たとえそれが品質を損なうことなく成功したとしても、それは実質的な改善の終わり (つまり S 字カーブの終わり) を示しています。
そして…予算モデルまたはフロンティア モデルによる単純な会話は確かに「コストがかかりすぎて計測できない」ものになっていますが、再帰的/間接的/ツール使用/「思考」(非) モデルを必要とする本格的な用途 (コーディング、複雑な推論など) はトークンの使用が爆発的に増加しており、これらの用途は非常に高価になっています。高い労力を費やす最上位の再帰モデルによる単一タスクのコストは、API レートで約 75 ドルと推定されます。 1 つのクエリで 100 万トークンが使用されるのを見たことがありますが、これは API レートで最大 25 ドルを意味します。
つまり…世界に提示されている経済モデルは、コストを隠したり、「測定するには安すぎる」と話したりしながら、複雑なものについて良い結果を近似するために最大限の総当たり力を必要とするタスクの価値の組み合わせに基づいているように見えます。
したがって、この船が沈没しない限り音楽を楽しみ、良い救命いかだを準備してください。
これはアプリケーション「I am b」のスクリーンショットの一部です。

作る（そして楽しむ）。これは、必要な図の作成をサポートしてくれる実際のアプリケーションです (つまり、データとグラフィックの組み合わせです)。目標は「LLM を使用したコーディング」を調査することであり、私が構築しているものは、このテーマについてある程度の経験があるので使用できると考えた単なる例です。
重要な観察から始めましょう。Claude Code と私自身の (錆びてはいますが、十分にしっかりした) プログラミングの背景の組み合わせのおかげで、私は Claude Code にアプリケーション (現時点では未完成ですが、機能します) を作成させることができました。これがなければ、このような短期間で作成することはできませんでした。これは、時間とエネルギーを考慮すると、「まったく作成できない」ことを意味します。しかし、落とし穴があります。いくつかさえあります。技術的な問題と「クロード・コードは実際にどこまで『コード』するのか？」という質問への答えについては、後ほど報告します。今日私が衝撃を受けたのは経済学のことです。
私は「雰囲気コーディング」（私はこの用語が大嫌いです。LLM の有無にかかわらず、コーディングにはほとんど「雰囲気」がありません）の実験を、今では 4 か月かけて構築しました（フルタイムではないことに注意してください）。私は LLM コーディングの感触を得るために非常に小さなプロジェクトから始めましたが、最終的に中程度の作業量設定で Opus 4.6 を使用する Claude Code に落ち着きました。高い設定を使用すると、森の中で迷子になることが多くなる傾向がありました。少ないと結果も貧弱でした (コードの品質管理については後の投稿で詳しく説明します)。最後のより本格的なプロジェクト (過去 2 か月) でしばらくすると、「中程度の労力」の結果が悪くなったので、「高」に切り替えて、ほぼ同じ品質に戻りました。
コスト管理も経験の一部でした。まず、月額 20 ドルのサブスクリプションを契約しました。すぐに使用制限に達してしまいました。 5 時間ごとにリセットされる制限があります

毎週リセットされるものでは、API 価格でトークンを購入することで制限を超えることができます。そこで、追加資金を追加してみました。そして、API コストで購入したトークンを使用すると、使用制限内で使用するよりも明らかに高価であることに気づきました。まだ月額 20 ドルのプランを利用していたとき、仕事を終わらせるために追加のトークンを購入したところ、数日以内に約 80 ドルのトークンを購入していました。その時点で、最も安価なオプションを使用し、必要に応じて API コスト レベルで資金を追加するよりも、月額 100 ドルを支払う方がはるかに良い取引であることが明らかになりました。
私はすでにコストを検討していました。少し前に、トレーニングと推論 (生成 AI による結果の生成) が比較され、コストの原因はトレーニングではなく推論/生成であると結論づけられました。そう結論付けたので、私はタスクを実行するために LLM を使用するコストを調べ始めました。そして、私は—開示—この研究を支援するために実際にいくつかの LLM チャットボット (Gemini、Claude) を使用しました。さて、これらの LLM は信頼性が低すぎると確信して読むのをやめる前に、たとえ自分で作成したコンテンツに間違いがあるとしても、LLM は実際にはかなりまともな検索エンジンです。結局のところ、彼らは ArXiv の記事をすべて取り込んでおり、ページランクに基づく従来の Web 検索では実行できなかったことを、LLM は実行できます。実際のコンテンツに基づいて情報を見つけます。オフィス環境で LLM を使用している人は、これを裏付けるでしょう。
私が最初に得た答えは、サム・アルトマンの有名な「推論が「あまりにも安すぎて計測できない」ようになってしまった」のような、標準的な簡単で汚い答えでした。このような記述は真実であることが判明しましたが、後でわかるように、誤解を招く/不完全でもあります。結局のところ、ユーザーにとって重要なのは「トークンあたり」の価格ではなく、「リソル」を取得するためにいくら支払うかです。

「使用中」または「タスクが完了しました」。そして、結果は改善されているものの、特に間接/再帰的 (「思考」モデル) でトークンの量が急増していることはすでにわかっていました。[余談: それらを「思考モデル」と呼ぶのは非常に誤解を招きます。これらはそのようなことは何も行いません。彼らが行うことは、何よりも目に見えない大量の再帰、間接化、試行錯誤です。LLM が舞台裏で 20 または 100 のアプローチを作成し、それを何度も繰り返すのと同じように、いくつかの追加ツール (最初の生成など) を使用します。 Python スクリプトを作成し、そのスクリプトを実行する必要があります (そのスクリプトに明らかなバグがなくなるまで複数回実行する必要があります)。さらに、数回テストしてから、その出力を LLM への追加の入力として使用するなど、ユーザーにはほとんど目に見えないものばかりです。] そこで、「2023 年第 1 四半期から現在までの平均クエリで使用されたトークンの量」から始めて、その傾向を調査しました。トークンあたりのコストの低下と使用されるトークンの量の増加を組み合わせることができました。数回試した後、より焦点を絞った開始プロンプトで新しい会話を開始しました。
人々は 1 つのクエリの結果を望んでいないことに留意することが重要です。これは最も単純なタスクにのみ当てはまります。実際には、複数のクエリを実行することが多く、結果を受け入れるまでに前後のやり取りが発生します。したがって、解決に至るまでに平均でどれくらいのクエリがかかるかを見積もる必要もあります。
そして、目に見えないトークンもすべてあります。課金の一部ではないトークンもあれば、「ダーク トークン」もあり、誤ってラベル付けされた「思考」モデルでは、バックグラウンドで膨大な負荷の間接/再帰と信じられないほどの「試行錯誤」が行われており、私には見えないものが存在します。

入力または出力ですが、すべて大量のトークンを使用します。これらの「再帰的」モデルでは、ユーザーが結果として目にするものを矮小化するほどの量のトークンが使用されることがわかっています。
これらも目に見えないトークンですが、目に見える出力 (生成) トークン コストで請求されます。コストを把握するために、Opus 4.6 はトークンごとに支払う場合、入力 (与えられるデータ、たとえば、すべてのクエリでコード ベースから抽出するものやバックグラウンドで読み取られるデータなど) 100 万あたり 5 ドル、トークン生成データ (バックグラウンドでの「再帰的な取り組み」に含まれる世代を含む) 100 万あたり 25 ドルを請求します。
とにかく、これはほとんどすべてがかなりの見積もりです。貴重なわずかなハードデータがあります。この投稿は、私が思いつく限りの最良かつ合理的な推定に基づいた結果であり、これは部分的には確かなデータ、つまり私自身のこれまでの経験に基づいています。
使用コストの推移を示すために、ちょっとしたフレームワークを考え出しました。
最初の要素は、トークンごとのコストを測定する方法です。以下ではサブスクリプションについて見ていきますが、私たちが従う標準は、「トークンごとの API」コストが時間の経過とともにどのように変化したかです。これは実際のトークンコストについて私たちが知る最良の推定値ですが、ベンダーがこれについてオープンにしない限り、私たちは仮定をする必要があるため、これは不安定です。 API 価格設定で利益は得られますか?それとも、その使用は依然として実際には赤字であり、補助金が出ているのでしょうか？現時点ではどちらも信じられるシナリオです。そしてAPに奇妙なことが起こった

[切り捨てられた]

## Original Extract

Coding with LLMs (Claude Code, OpenAI Codex) is often presented as the 'killer app' for Generative AI. But looking at data, it seems the one piece of the puzzle missing is actual cost. A quest into getting a less muddy picture about what is going on, with surprising results.

Skip to content
R&A IT Strategy & Architecture
Insights on all things 'digital', both technical and non-technical
Books and Collections
“ChatGPT and Friends” Collection
“On the Information Revolution” Collection
Mastering ArchiMate – Edition 3.2
Free ArchiMate 3.2 Overview PDFs in multiple languages
EA Chess — The Book
Chess and the Art of EA — The Summary
Article Collections
The “ChatGPT and Friends” Collection
Menu
Open a search box
Close a search box
Search for:
Search
Anthropic/OpenAI may be spending more than $1000 for every $100 you pay them
For reasons that will remain hidden, we resume writing about Generative AI/LLM after a hiatus of 15 months ( that one from October 2025 , and the one from June 2025 , don’t really count as serious pieces). Today, the first of two articles about “coding with Large ‘Language’ Models”, as coding with LLMs is positioned as the ‘ killer app ‘ for LLMs.
We interrupt this program for a short digression on Anthropic’s recently released blog post When AI builds itself .
Did Anthropic perchance hire Google marketeers?
Anthropic’s blog post is a masterclass in suggestive writing. The caveats are there, but hidden or sandwiched between more hyperbolic statements. A sentence ‘we might be wrong’ is there, but what role has it as a single sentence in thousands of words of text assuming that they are not wrong? The benchmarks are suspect (a 50% or even 80% success rate on a coding task compared to a human is effectively completely useless in full-agentic (no human in the loop) coding. Is checking in 8 times as many lines of code per day really a good thing? What if every day you are replacing what wasn’t OK the day before? What if LLMs edit in a way that Lines of Code becomes less trustworthy a measure as it is already? All in all, it reminds me of Google’s deceptive talk about its ‘Willow’ QM computing chip .
By the way, on that “Is checking in 8 times as many lines of code per day really a good thing?”: A lot of my LLM-based checkins have been like that, and frankly, I have increased my number of checkins just to be able to backtrack if Claude Code has gotten lost. Even with all the work I did without intermediary checkins, I committed 7 times as much change than I ended up with in terms of lines of code…. So, ‘8’ as an estimate not so much of increased productivity, but of increased overhead does sound about right…
We will return to this if I get around to writing the in-depth of ‘coding with Claude’ later.
And while we’re at it. This is a long and winding piece (I am really having some trouble on this front, apologies), so let’s provide:
TL;DR — It doesn’t look like LLM-coding is going to be affordable
(let alone for “AI to build itself”)
I have been doing some experimenting. The experiment is: “How good is Claude Code anyway?”. That experiment is still running and Claude Code has by now created around 40k lines of code and a working (though incomplete) application. I hope to report on that experience in a short while (but it is a much more difficult write-up). In the meantime, I experienced the cost issue and it led to a short research project, which led to a number of interesting observations and conclusions:
Let’s start with an important observation: Thanks to the combination of Claude Code and my own (rusty, but solid enough) programming background, I have been able to let Claude Code create this application (unfinished as of now, but functional) that I would otherwise not have been able to create in such a short amount of time, which — given time and energy — would have meant: ‘not at all’. For an experienced programmer, the initial experience is extremely impressive as an experienced programmer knows how much understanding by themselves normally goes in to creating such code;
But… LLM-coding isn’t economically viable for most uses. It is viable now because the subscriptions are heavily subsidised. But if you use the $100 a month Claude Max plan, and you would use it to the weekly limit by going full ‘agentic coding’ (so almost no human in the loop) you would use an amount of tokens that would cost you more than $1000 at API-pricing. Anthropic seems to be busy (Opus 4.7, 4.8) to stop that bleeding, and even if that succeeds without a loss of quality, it does signal an end of substantial improvements (i.e. the end of an S-curve);
And… while simple conversations with either budget or frontier models have become indeed ‘too cheap to meter’, the serious uses (like coding, complex reasoning) that require the recursive/indirection/tool-using/’thinking’ (not) models have exploded so much in token use that these uses have become very expensive. A single task by a top recursive model at high effort is estimated to cost around $75 at API-rates. I have seen a single query use one million tokens, which would mean max $25 at API-rates;
So… the economic model that is being presented to the world seems based on the combination of the value of the tasks that require a maximum amount of brute force to approximate good results on anything complex, while hiding the cost or talking about ‘too cheap to meter’;
Hence: enjoy the music for as long as this ship hasn’t sunk, and prepare a good life raft.
Here is a part of a screen shot of the application I am building (and having some fun with). It’s a real application which may support me in creating diagrams I need (so it’s a combination of data and graphics). The goal has been to investigate ‘coding with an LLM’, the thing I am building is just an example I thought I could use because I am somewhat experienced on the subject:
Let’s start with an important observation: Thanks to the combination of Claude Code and my own (rusty, but solid enough) programming background, I have been able to let Claude Code create an application (unfinished as of now, but functional) that I would otherwise not have been able to create in such a short amount of time, which — given time and energy — would have meant: ‘not at all’. But there is a catch. A few even. I will report later on the technical issues and an answer to the question: “To what extent does Claude Code actually ‘code’?”. Today it is simply the economics that struck me.
I built up my ‘vibe coding’ (I really dislike that term, there is little ‘vibe’ about coding with or without an LLM) experiment over — by now — 4 months (not full time, note). I started with a very small project to get a feel of LLM-coding, finally settled on Claude Code using Opus 4.6 on medium effort setting. If I used a higher setting it tended to get lost in the woods more often. Less, and the results were poor (more on managing the quality of the code in that later post). After a while on my final more serious project (the last two months) the results of ‘medium effort’ got worse, so I switched to ‘high’, returning to approximately the same quality.
Managing cost was also part of the experience. First, I took out a $20/month subscription. I quickly ran into usage limits. You have a limit that resets every 5 hours and one that resets every week, you can go beyond the limit by buying tokens at API-pricing. So, I tried adding some extra funds. And I noticed that the use of purchased tokens at API-cost was obviously much more expensive that staying within the usage limits. When I was still on the $20/month plan, and I bought some extra tokens to get some job finished, within a few days I had bought about $80 of tokens. At that point it became clear to me that paying $100/month was a much, much better deal than using the cheapest option and adding funds at API-cost level when needed.
I had already been looking at cost, e.g. training versus inference (generating results by Generative AI) a while back, concluding that not training, but inference/generation is what drives the cost. Having concluded that, I now started to look into the cost of LLM- use to perform a task . And I — disclosure — actually used a few LLM-chatbots (Gemini, Claude) to help me in this research. Now, before you stop reading because you are convinced these LLMs are too unreliable, they really are pretty decent search engines, even if they make mistakes on content they produce themselves. After all, they have ingested all those ArXiv articles and what a classic web search based on page-rank would not have been able to do, LLMs can: they find you information based on its actual content. People using LLMs in office setting will corroborate this.
The first answers I got were the standard quick & dirty answers, like Sam Altman’s famous: inference having become ‘too cheap to meter’. Such statements turned out to be true but also misleading/incomplete, very much so, as you will see. At the end of the day, the price ‘per token’ is not what is relevant for users, it is how much you pay for getting a ‘resolution’, or a ‘task finished’. And we already knew that while the results have improved, the amount of tokens had skyrocketed, especially with the indirect/recursive (‘thinking’ models. [Aside: calling them ‘thinking models’ is extremely misleading. They do nothing of the kind . What they do is above all a massive amount of invisible recursion, indirection, and trial and error. Like an LLM creating 20 or a 100 approaches behind the scenes, doing that time and again, use some additional tools (like first generate a python script, then run that script — having to run it multiple times until there are no obvious bugs in that script), maybe even test it a few times, and then use the output of that as more input for the LLM. And so on. A lot of indirection and recursion, a lot of ‘trial and error’, everything almost invisible to the user.] So, I investigated that trend, starting with “How much tokens did the average query use between Q1 2023 and now?” so I could combine falling per-token cost with rising amount of tokens used. After a few tries, I started a new conversation with a more focused start prompt.
It is important to keep in mind that people do not want the result of a single query, that is only true in the most simple tasks. In reality, they often make multiple queries and there is a back and forth, before they accept the result. So, you also need to estimate how many queries on average it takes to get to a resolution.
And then there are all those tokens you do not see. There are tokens that aren’t part of billing, there are ‘dark tokens’, and with the mis-labeled ‘thinking’ models there is a massive load of indirection/recursion and an incredible amount of ‘trial and error’ in the background going on, things you do not see as input or output, but all using a shitload of tokens. We know that these ‘recursive’ models use amounts of tokens that dwarf what you as a user see as result.
These are also tokens that you do not see, but they are nonetheless billed to you at visible output (generated) token cost. To get an idea of cost, Opus 4.6 when you pay per-token charges you $5/million tokens input (the data you give it, e.g. what it extracts from your code base on every query and what it reads in the background), and $25/million tokens generated data (including that generation that is in those ‘recursive efforts’ in the background.
Anyway, almost all of this is a lot of estimation. There is precious little hard data. This post is the result based on whatever best and reasonable estimates I could come up with, which in part is hard data: my own experience so far.
To show the development of use-cost, I’ve come up with a little framework:
The first element is how to measure per-token cost . We will look at subscriptions below, but the standard we follow is how ‘API per-token’ costs have developed over time. This is the best estimate we have of real token cost, but it is fragile, because as long as the vendors do not open up about this we must make an assumption. Do they make a profit at API-pricing? Or is that use still actually loss-making and subsidised? Either is a believable scenario at this time. And weird things have happened to AP

[truncated]
