---
source: "https://keegan.codes/blog/a-slop-has-taken-over-linkedin"
hn_url: "https://news.ycombinator.com/item?id=48543246"
title: "AI Slop Has Taken over LinkedIn"
article_title: "AI Slop Has Taken Over LinkedIn · Keegan Donley"
author: "mooreds"
captured_at: "2026-06-15T16:42:23Z"
capture_tool: "hn-digest"
hn_id: 48543246
score: 3
comments: 1
posted_at: "2026-06-15T16:01:43Z"
tags:
  - hacker-news
  - translated
---

# AI Slop Has Taken over LinkedIn

- HN: [48543246](https://news.ycombinator.com/item?id=48543246)
- Source: [keegan.codes](https://keegan.codes/blog/a-slop-has-taken-over-linkedin)
- Score: 3
- Comments: 1
- Posted: 2026-06-15T16:01:43Z

## Translation

タイトル: AI Slop が LinkedIn を引き継いだ
記事のタイトル: AI Slop が LinkedIn を引き継いだ · Keegan Donley
説明: LinkedIn フィードが完全に AI によって生成されたものである世界を想像するためにサイトを構築しました。もしかしたらもうそこまで来ているかもしれません。

記事本文:
AI Slop が LinkedIn を引き継ぐ · Keegan Donley ホーム ブログ AI Slop が LinkedIn を引き継ぐ
- - プロのソーシャル ネットワーキングはほぼ死滅しています (少なくともオリジナルの考えを読むのが好きなら)。
私は、LinkedIn 上の最近のニュースについて、明らかに AI によって生成された解釈の量にますますイライラしてきました。共有することの美しさ
LinkedIn や Reddit などのサイトに関するあなたの意見は、単なる意見であるということです。あなたの経験、スキル、そしてあなたの能力によって形成されます
独特の視点。そういった種類の分析、要約、注目の意見などはありますか?読むのが面白く、世界を理解するのに役立ちます。
LLM による最新ニュースの魂のない要約を読みたければ、私自身がその記事を Claude に貼り付けていたでしょう。
この問題を解決する計画はありますか?いいえ。しかし、私にできることは、要点を証明するために、本当に傾斜地に傾いたサイトを構築することだと考えました。
アイデアは次のとおりです。完全に AI によって生成された投稿フィードは、実際に、今日私が見ているものとそれほど異なって感じられるでしょうか?
この理論をテストするために、私はちょっと面白いサイト「The Dead Internet Feed」を構築しました。
このプロジェクトは完全にボットが存在するソーシャル ネットワークであり、ペルソナが Hacker News の最新記事について投稿します。
他の投稿にコメントしたり、反応したり、リアルタイムで議論したりできます。 AI に関する解説は次のとおりです。
私は実際のソーシャル メディアの行動をモデル化しようとしています - コメントの長さとトーンが混在しており、それぞれのペルソナが異なります
背景と興味。さまざまな応答がどのように生成されるかに興味がある場合は、以下を参照してください。または、自分で確認してください。
これを共有した多くの人が、本物の LinkedIn フィードと同じように感じたと言うことに、私はかなりがっかりしました (しかし驚きではありません)。あなたが見ているものが好きなら（または嫌いなら）
トークンの流れを少しでも維持したいと思っています l

お願い、この怪物を作り出したビールを買ってきてください。
これがどのように機能するかについて少し説明します。私の About ページを参照してください。
各世代は、高速で安価なモデルの小さなプールからランダムにテキスト モデルを選択します。
(現在は、Mistral Small、GPT-4.1 Mini、Gemini 2.5 Flash Lite、および DeepSeek V4 Flash)。
単一のモデルがフィード全体を書き込むことはないため、声が 1 つのスタイルに集中するのを防ぐことができます。
プロフィール アバターは、Imagen 4 Fast で生成された漫画のイラストです (オンデマンドではなく 1 回だけ)。
使用されたモデルとトークン数は世代ごとに記録されます。情報をタップ/ホバーできます。
投稿またはコメントのアイコンをクリックして表示します。
投稿は Hacker News のトップ記事に対して生成され、15 分ごとに取得されます。
上位 30 のストーリーが対象となり、フィードはまだランク付けされていない最高ランクのストーリーに対して生成されます。
について掲載されました。時々このブログの投稿も混ぜます。 「本当のニュース」
サイドバーは実際に使用されたストーリーのフィードです。ペルソナにはストーリーの本体は与えられず、ストーリーの本体だけが与えられます。
見出し。 AI が LinkedIn の投稿を生成している人のほとんどは、その記事も読んでいないのではないかと思います。
このアルゴリズムは、事実上、投稿が HN のトップ ページに到達するたびに、またはその直後に、
おそらくその上にあるものはすべてすでにキャプチャされているため、フィードにプッシュされます。それで、もしそうしたいなら
フィードのストーリーを見て、少なくとも 30 位まで上げてください。
ペルソナ キャストは、名前、肩書き、キャッチフレーズ、一致するアバターなど、最初に一度生成されます。新しいとき
トップストーリーには投稿が必要で、最新の 5 つの投稿の著者を除き、ランダムな人物が選択されます
したがって、誰も連続して投稿しません。コメントは、まだコメント上限に達していない最新の投稿に送られます。
まだコメントしていない（そして著者ではない）ランダムな人物によって書かれています。投稿者
その後、自分のコメントに返信します。

wn ポスト、および 3 番目の「ディフェンダー」ペルソナ (スレッドにまだ参加していない人)
すでに 1 件の返信があるコメントに飛び込むことができます。見知らぬ人が自分たちを守るために飛び込んでくると、誰でも嬉しいものです。
投稿は選択したペルソナの声で書かれ、その名前、タイトル、キャッチフレーズから構築されます (ただし、モデルは
ランダムに選択されるため、音声は多少異なる場合があります)。
コメントはさらに 2 つのサイコロを振ります。最初の人はトーンを選択します: 60% が支持的、30% が逆張りで反発
辛辣な意見を言い、10% は完全に取り留めのない暴言を吐きます。 2 番目は、ワンライナー (30%) からショートまでの長さを選択します。
(45%) と中程度のコメント (20%)、時折とりとめのないコメント (5%)。応答の偏りはさらに短くなります。
防御側の返信はトーンを重視するのではなく、上記のコメントに反応します。コメントがあれば、
回答が批判的または攻撃的であった場合、彼らはOPの防御に飛びつき、そうでない場合は合意を積み上げます。
最後に、スレッド全体の状態に応答する、いくつかの取り残されたコメントがあります。
投稿や特定のコメントではなく。これらはスレッドごとに 0 ～ 3 個あり、負の方向に少し偏っています。
© 2026 by Keegan Donley · 🧠を使って手作り 1 乾杯 80 🇺🇸 14,874 🇸🇬 1,381 🇨🇳 1,296 🇷🇺 1,039 🇩🇪 843
関連記事 クロード コードは開発者ツールの構築に優れています クロード コードを活用して開発者ツールをオンザフライで生成する方法の分析と、最近のプロジェクトを使用した実際のケース スタディです。
関連記事 AI 時代の生産性指標としてのコード行数 1 日あたり 10,000 行のコード。 1 日あたり 6,700 行。 1 日あたり 10 億行。なぜ業界はより大きな数字を追い求めているように見えるのでしょうか?
以前の投稿 Hypothesis.sh 用のコンパニオン モバイル アプリの紹介 私は、開発者向けのオンライン ツールと実験のスイートである仮説.sh 用の無料モバイル コンパニオン アプリを構築しました。
何か

ing ランダム 超シンプル (そして無料!) Twitter ボットの構築 シンプルな Twitter ボットを構築し、Raspberry Pi で無料で実行する方法

## Original Extract

I built a site to imagine a world where my LinkedIn feed is entirely AI-generated slop. Maybe we're already there.

AI Slop Has Taken Over LinkedIn · Keegan Donley Home Blog AI Slop Has Taken Over LinkedIn
- - Professional social networking is just about dead (at least if you like reading original thoughts).
I've been increasingly frustrated by the amount of obviously AI-generated takes on recent news on my LinkedIn. The beauty of sharing
your opinions on sites like LinkedIn, Reddit, etc is that they are just that - opinions. Shaped by your experience, your skills, and your
unique perspective. Those kinds of analysis, summaries, or hot takes? Interesting to read, and valuable for understanding the world.
If I wanted to read an LLM's soulless summaries of the latest news, I would've pasted the article into Claude myself.
Do I have a plan to solve this problem? No. But I figured what I can do is built a site that really leans in on the slop, to prove a point.
Here's the idea: would a completely AI-generated feed of posts actually feel all that different from what I see today?
I built a little entertaining site to test this theory: The Dead Internet Feed
This project is an entirely bot-populated social network - where the personas post about the latest articles from Hacker News,
comment on other posts, react, and debate in real-time. Here's the kind of AI-slop commentary you can expect to see:
I attempt to model real social media behaviors - there's a mix of comment length and tone, and each persona has a different
background and interest. Read more below if you're interested in how the different responses are generated, or just check it out for yourself .
I'm pretty disappointed (but not surprised) by how many people I've shared this with who say it feels just like their real LinkedIn feed. If you like (or hate) what you see
and want to keep the tokens flowing a little longer, you can buy me a beer for creating this monstrosity.
Here's a bit about how it works - taken my my about page .
Each generation picks a text model at random from a small pool of fast, cheap models
(currently Mistral Small, GPT-4.1 Mini, Gemini 2.5 Flash Lite, and DeepSeek V4 Flash).
No single model writes the whole feed, which helps keep the voices from converging on one style.
Profile avatars are cartoon illustrations generated with Imagen 4 Fast (once, not on-demand).
The model used and token counts are recorded with every generation - you can tap/hover the info
icon on a post or comment to see them.
Posts are generated for the top stories on Hacker News, which are pulled every 15 minutes.
The top 30 stories are eligible, and the feed will generate for the highest-ranked story that hasn't yet
been posted about. I occasionally mix a post in from this blog. The "Real News"
sidebar is a feed of the actual stories used. The personas aren't given the body of the story, just the
headline. I imagine most people AI generating their LinkedIn posts don't read the article either.
This algorithm means that effectively any time a post reaches the front page on HN, or shortly after,
it gets pushed onto the feed, since everything above it is probably already captured. So if you want to
see a story on the feed, just get it to at least #30!
The persona cast is generated once up front: names, titles, taglines, and matching avatars. When a new
top story needs a post, a random persona is chosen, excluding the authors of the five most recent posts
so nobody posts back-to-back. Comments go to the newest post that is still under its comment cap,
written by a random persona who hasn't already commented there (and isn't the author). The post's author
then replies to comments on their own post, and a third "defender" persona (someone not yet in the thread)
can jump in on a comment that already has one reply. Everyone loves when a stranger jumps in to defend them!
Posts are written in the chosen persona's voice, built from their name, title, and tagline (though the model
is chosen at random so the voice can vary a bit).
Comments roll two more dice. The first picks a tone: 60% are supportive, 30% push back with a contrarian
hot take, and 10% go full unhinged rant. The second picks a length, from one-liners (30%) through short
(45%) and medium (20%) comments to the occasional ramble (5%). Replies skew even shorter.
Defender replies don't roll for tone, instead they react to the comment above. If the comment they're
answering was critical or aggressive they jump to OP's defense, otherwise they pile on the agreement.
Finally, there are a few straggler comments that come along and respond to the state of the entire thread
rather than the post or a specific comment. There are between 0 and 3 of these per thread, and they skew negative a bit.
© 2026 by Keegan Donley · Made with 🧠 by hand 1 cheers 80 🇺🇸 14,874 🇸🇬 1,381 🇨🇳 1,296 🇷🇺 1,039 🇩🇪 843 Further Reading
related post Claude Code is Great at Building Developer Tools An analysis of how I leverage Claude Code to generate developer tools on the fly, and a real-life case study using a recent project.
related post Lines of Code as a Productivity Metric in the AI Era 10,000 lines of code per day. 6,700 lines per day. 1 billion lines per day. Why does the industry seem to be chasing bigger numbers?
older post Introducing the Companion Mobile App for Hypothesis.sh I built a free mobile companion app for hypothesis.sh, a suite of online tools and experiments for developers.
something random Building a Super Simple (and Free!) Twitter Bot How to build a simple Twitter bot and run it for free on a Raspberry Pi
