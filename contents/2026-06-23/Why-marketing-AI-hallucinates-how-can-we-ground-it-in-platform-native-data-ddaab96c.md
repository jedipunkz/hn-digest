---
source: "https://www.fuse.is/blog/how-fuse-talks-to-tiktok-ads"
hn_url: "https://news.ycombinator.com/item?id=48646427"
title: "Why marketing AI hallucinates: how can we ground it in platform-native data"
article_title: "How Fuse Talks to TikTok Ads | Cross-Channel AI Marketing Intelligence"
author: "rkovashikawa"
captured_at: "2026-06-23T16:06:07Z"
capture_tool: "hn-digest"
hn_id: 48646427
score: 1
comments: 0
posted_at: "2026-06-23T15:18:30Z"
tags:
  - hacker-news
  - translated
---

# Why marketing AI hallucinates: how can we ground it in platform-native data

- HN: [48646427](https://news.ycombinator.com/item?id=48646427)
- Source: [www.fuse.is](https://www.fuse.is/blog/how-fuse-talks-to-tiktok-ads)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T15:18:30Z

## Translation

タイトル: マーケティング AI が幻覚を起こす理由: プラットフォーム固有のデータに基づいて考えるにはどうすればよいか
記事のタイトル: Fuse が TikTok 広告と対話する方法 |クロスチャネル AI マーケティング インテリジェンス
説明: Fuse が TikTok Ads を独自の AI ブレインに接続し、それを Meta、Google、LinkedIn、SEO と相互参照して、信頼できるクロスチャネル インテリジェンスを実現する方法。

記事本文:
Fuse が TikTok 広告と対話する方法 |クロスチャネル AI マーケティング インテリジェンス
製品 AI テクノロジー ソリューション 価格 プレス インテグレーション 企業セキュリティ ログイン 無料で始める ログイン はじめる インテグレーション 2026 年 6 月 11 日 Fuse が TikTok 広告とどのように対話するか (そして、すべてのプラットフォームが独自の AI ブレインを取得する理由)
すべてのマーケティング プラットフォームに独自の AI ブレインを提供した理由と、それによって FUSE が 1 つの質問でスタック全体に対して TikTok を相互参照できるようになった理由。
TL;DR: Fuse はすべての広告プラットフォームに独自の専用 AI ブレインを提供します。 TikTok について平易な英語で質問すると、アナリスト級の回答が得られます。これには、実際に保存されたデータのみから構築された、クロスチャネルのコンテキストと予算に関する推奨事項が含まれています。設計により幻覚ゼロ: AI が数値を発明することはありません。
‍
広告プラットフォームへの接続は解決された問題のように思えます。アクセス トークンを取得し、API を呼び出して、番号を取得します。実際には、すべてのプラットフォームは独自の方言を話しており、すべてのプラットフォームの上に AI アシスタントを配置しようとすると、その小さな方言の違いが実際のエンジニアリング上の決定に変わります。
これは、Fuse を TikTok Ads にどのように接続するか、TikTok が独自のパズルである理由、そして平易な英語で質問し、返されるすべての数字を信頼できるように私たちがたどり着いたアーキテクチャについての物語です。
‍
すべての広告プラットフォームは異なる方言を話します
TikTokを追加する前にすでにMetaとGoogleをサポートしていたので、見た目はほぼ同じだろうと想定して導入しました。そうではありませんでした。
TikTok が独自のビートに合わせて行進する例をいくつか挙げます。
「それは私です」の別の言い方。ほとんどのプラットフォームでは、資格情報が標準的な方法で提示されることを期待しています。 TikTokは独自の形式で引き渡されることを望んでいる。細かいことですが、自分が作った配管を他の人のために再利用することはできないということです。
ログイン

帽子には期限がありません。ほとんどのプラットフォームでは、接続時に取得したキーは自動的に期限切れになり、スケジュールに従って更新する必要があります。 TikTok は、誰かが明示的に取り消すまで有効期限が切れません。これは便利ですが、接続を長期にわたって保存し信頼する方法について、これまでとは異なる考え方が必要になることも意味します。
アカウントは事前に渡されます。 TikTok に接続すると、後で個別にアカウントを見つけに行くのではなく、接続時にどの広告主アカウントにアクセスできるかが表示されます。
本番前の練習室。 TikTok は別のサンドボックス環境を提供しているため、顧客の実際の広告支出に何かを指摘する前に、現実的ではあるが偽のデータに対してリハーサルを行うことができます。
これらはどれも取引を妨げるものではありません。しかし、これらを総合すると、「広告プラットフォームへの接続」が決して 1 つの作業ではないことを思い出させてくれます。これはどのチャネルにとっても新鮮な翻訳の問題であり、その認識がその後のすべてを形作りました。
本当の問題: 実際に信頼できる平易な英語の質問
マーケティング担当者が実際に望んでいることは、言うは易く伝えるは難しいものです。平易な英語で質問し（「先月のTikTokキャンペーンの成果はどうでしたか？」など）、実際に行動できるバックナンバーを入手してください。
素朴なアプローチは、AI に広告プラットフォームと直接対話させ、即興で動作させることです。私たちはそれをしませんでした。理由は次の 3 つです。
遅いですね。キャンペーン データを毎回ライブで取得して処理すると、待ち時間が発生します。
高価です。これらの API を使用して、すべての質問に対して大規模なモデルを実行すると、結果が速くなります。
それは物事を補うことができます。生データを通じてモデルを自由形式のままにしておくと、実際ではない数値を自信を持って報告できます。マーケティング上の意思決定において、「確信を持って間違っている」というのは最悪の結果であるため、私たちはそれが不可能になるようにシステム全体を設計しました。
つまり、1 つの AI の代わりに、私は

すべてを即興で実現するために、私たちは各プラットフォームに専用の頭脳を与えました。
アイデア: 各プラットフォームは独自のコネクタと独自の AI ブレインを取得します。
内部では、私たちがサポートするすべてのプラットフォーム (TikTok、Meta、Google、LinkedIn、Klaviyo、Shopify など) に独自の自己完結型モジュールが組み込まれています。つまり、そのプラットフォームの方言を認識するコネクタと、そのプラットフォームに関する質問に答える方法を知っている AI の「頭脳」が組み合わされています。
これらの脳はそれぞれ、次の 2 つの速度で動作します。
即座にキャッシュされた回答 (無料)。キャンペーンのパフォーマンスの概要などの一般的な質問は、事前に計算されキャッシュされた結果から提供されます。これらはすぐに戻ってきますし、クレジットもかかりません。
より深いリサーチモード。予測しにくい質問については、プラットフォームの AI 頭脳が実際の調査作業を行い、どのデータを取得するかを正確に把握します。
この設計の大きな利点は一貫性です。新しいプラットフォームを追加するということは、アシスタント全体を再配線することを意味するわけではありません。それは、他のすべての脳と同じスロットに差し込む自己完結型の脳をもう 1 つ構築することを意味します。先ほど述べた TikTok の癖は、漏れ出て他のすべてを複雑にするのではなく、TikTok モジュール内にきちんと含まれています。
見返り: AI アナリストが独自のクエリを作成 (数値をでっち上げることはありません)
ここで、Fuse はダッシュボードであることをやめ、アナリストになり始めます。そのより深い調査モードでは、プラットフォームの頭脳が独自の AI マーケティング アナリストのように機能します。事前に作成されたレポートの短いメニューから選択するのではなく、質問に必要なデータを正確に把握し、独自のクエリを作成してそれを取得し、答えとそれに対する対処方法を返します。
どれだけハードに動作するかについても賢明です。
繰り返しの質問を認識します。同様の質問が以前に回答されている場合は、最初から再発明するのではなく、実証済みのアプローチを再利用します。これにより、処理が速くなり、

信頼できる鉱石。
適切な量​​の努力を選択します。率直な質問には、すぐに直接アクセスできます。本当に複雑なものについては、より慎重で段階的な調査が行われます。
常にクリーンなデータを返します。どのような結果が得られたとしても、結果は整然と保存されたデータセットであるため、すべての答えは、再検討したり、図表を作成したり、構築したりできる実際の数値に基づいています。
強力になるところ: すべてを一度に相互参照する
ここで、各プラットフォームに独自の頭脳を与えることが本当に効果を発揮します。すべてのチャネルが同じ形状でクリーンで保存されたデータセットを生成するため、AI はそれらをまとめて、同時にすべてのデータセットを推論できます。
したがって、TikTok 広告を単独で見る必要はありません。同じ質問で、接続している他のすべてのものを取り込むことができます。
残りの有料チャンネル (Meta、Google、LinkedIn など) なので、TikTok の活動をすでに運営しているチャンネルと比較できます。
SEO と検索の可視性データ (オーガニック検索と AI の回答でどのように表示されているか) は、有料のパフォーマンスのすぐ隣にあります。
ファネルのさらに下にあるプラットフォーム (電子メールやストアなど) なので、1 つのチャネルでの支出がその後の実際の出来事につながる可能性があります。
これにより、単純なレポートが真のクロスチャネル インテリジェンスに変わります。 「TikTok はどうでしたか?」の代わりに、より難しく、より有益な質問をすることができます。
「TikTokは実際に新しい視聴者を取り込んでいるのか、それともMetaやGoogleですでにお金を払っているのと同じ人たちにリーチしているだけなのか？」
「TikTok の支出が増えると、ブランド検索やオーガニックの可視性もそれに伴って変化しますか?」
「先月の結果を実際に評価しているのはどのチャンネルでしょうか。次のお金はどこに送られるのでしょうか?」
実際の動作は次のとおりです。 「先月のTikTokキャンペーンの成果はどうでしたか？新しい顧客は見つかっていますか？」と尋ねてください。ヒューズはそれだけではありません

メトリックをエコーバックします。分析結果は次のとおりです。TikTok 広告の ROAS は 4.1 倍で、最高の有料チャンネルであり、リーチしたユーザーの 41% は、Meta や Google ですでにお金を払っているオーディエンスではなく、まったく新しいユーザーでした。そのため、より多くの予算を TikTok の見込み客に振り向けることを推奨しています。その答えのすべての数字は保存されたデータから直接読み取られたものであり、AI が数字を作り上げたことはありません。
このような質問に答えるには、それぞれが独自の方言を話す複数のプラットフォームからのデータを並べて正規化し、全体像を推論する必要があります。これこそまさに、プラットフォームごとの頭脳設計が可能にするものであり、Fuse をレポート ツールからマーケティング スタック全体の真に高度な AI アナリストに変えるものです。
マーケティングを行う場合にこれが重要な理由
違いを感じるためにアクセス トークンやサンドボックスを気にする必要はありません。
どの数字も信頼できます。答えは実際に保存されたデータに基づいており、推測されることはありません。設計上、幻覚はゼロです。
単なるチャートではなく、アナリストが得られます。 Fuse は何が起こったかを報告するだけでなく、新規視聴者の分析から予算の推奨事項まで、次に何をすべきかを示します。
よくある質問はすぐに無料で回答されます。いつも尋ねたことは、クレジットを消費せずにすぐに返されます。
一度に 1 つのチャンネルではなく、全体像。すべてのプラットフォームが同じクリーンなデータセットをフィードするため、単一の質問で TikTok を他の有料チャンネルと相互参照したり、SEO の可視性を相互参照したりできます。通常、アナリストと大量のスプレッドシートが必要な種類の分析です。
これを自分のデータで確認したいですか?統合の完全なリストを参照し、価格を比較し、デモを予約します。
マーケティングを転換するための信頼できるソリューション
データを収益につなげます。

## Original Extract

How Fuse connects TikTok Ads to its own AI brain and cross-references it against Meta, Google, LinkedIn, and SEO for trustworthy cross-channel intelligence.

How Fuse Talks to TikTok Ads | Cross-Channel AI Marketing Intelligence
Product AI Tech Solutions Pricing Press Integrations Company Security Log In Start for Free Log In Get Started Integrations June 11, 2026 How Fuse Talks to TikTok Ads (and Why Every Platform Gets Its Own AI Brain)
Why we gave every marketing platform its own AI brain, and how that lets FUSE cross-reference TikTok against your entire stack in a single question.
TL;DR: Fuse gives every ad platform its own dedicated AI brain. Ask about TikTok in plain English and get an analyst-grade answer, complete with cross-channel context and budget recommendations, built entirely from your real, saved data. Zero hallucinations by design: the AI never invents a number.
‍
Connecting to an advertising platform sounds like it should be a solved problem. You get an access token, you call an API, you get your numbers back. In practice, every platform speaks its own dialect, and the moment you try to put an AI assistant on top of all of them, those little dialect differences turn into real engineering decisions.
This is the story of how we connect Fuse to TikTok Ads, why TikTok was its own kind of puzzle, and the architecture we landed on so that you can ask a question in plain English and trust every number that comes back.
‍
Every ad platform speaks a different dialect
We already supported Meta and Google before we added TikTok, so we went in assuming it would look roughly the same. It didn't.
A few examples of where TikTok marches to its own beat:
A different way of saying "it's me." Most platforms expect your credentials to be presented one standard way. TikTok wants them handed over in its own format. Small detail, but it means you can't just reuse the plumbing you built for everyone else.
Logins that don't expire. On most platforms, the keys you get when you connect quietly expire and have to be refreshed on a schedule. TikTok's don't expire until someone explicitly revokes them. That's convenient, but it also means you have to think differently about how connections are stored and trusted over the long haul.
Your accounts are handed to you up front. When you connect TikTok, it tells you which advertiser accounts you have access to right at connection time, rather than making you go discover them separately afterward.
A practice room before the real thing. TikTok offers a separate sandbox environment, so we can rehearse against realistic-but-fake data before pointing anything at a customer's real ad spend.
None of these are dealbreakers. But added together, they're a good reminder that "connect to the ad platform" is never one job. It's a fresh translation problem for every channel, and that realization shaped everything that came next.
The real problem: plain-English questions you can actually trust
What marketers actually want is simple to say and hard to deliver: ask a question in plain English (like "how did my TikTok campaigns do last month?") and get back numbers you can actually act on.
The naive approach is to let the AI talk directly to the ad platform and improvise. We didn't do that, for three reasons:
It's slow. Pulling and crunching campaign data live, every single time, makes you wait.
It's expensive. Hitting these APIs and running a large model on every question adds up fast.
It can make things up. A model left to free-form its way through raw data can confidently report a number that isn't real. For marketing decisions, "confidently wrong" is the worst possible outcome, so we designed the whole system to make it impossible.
So instead of one AI improvising across everything, we gave each platform its own dedicated brain.
The idea: each platform gets its own connector and its own AI brain
Under the hood, every platform we support (TikTok, Meta, Google, LinkedIn, Klaviyo, Shopify, and more) gets its own self-contained module: a connector that knows that platform's dialect, paired with an AI "brain" that knows how to answer questions about it.
Each of these brains works at two speeds:
Instant cached answers (free). Common questions, like a campaign performance overview, are served from pre-computed, cached results. These come back fast and don't cost you any credits.
A deeper research mode. For the less predictable questions, the platform's AI brain does real investigative work to figure out exactly what data to pull.
The big win of this design is consistency. Adding a new platform doesn't mean rewiring the whole assistant. It means building one more self-contained brain that plugs into the same slots as all the others. The TikTok quirks we mentioned earlier stay neatly contained inside the TikTok module, instead of leaking out and complicating everything else.
The payoff: an AI analyst that writes its own queries (and never makes up a number)
This is where Fuse stops being a dashboard and starts being an analyst. In that deeper research mode, the platform's brain works like your own AI marketing analyst: instead of picking from a short menu of pre-built reports, it figures out exactly what data your question needs, writes its own query to go get it, and comes back with the answer and what to do about it.
It's smart about how hard it works, too:
It recognizes repeat questions. If a similar question has been answered before, it reuses the proven approach instead of reinventing it from scratch, which is faster and more reliable.
It picks the right amount of effort. Straightforward questions get a quick, direct path. Genuinely complex ones get a more careful, step-by-step investigation.
It always hands back clean data. However it gets there, the result is a tidy, saved dataset, so every answer is grounded in real numbers you can revisit, chart, and build on.
Where it gets powerful: cross-referencing everything at once
Here's where giving each platform its own brain really pays off. Because every channel produces clean, saved datasets in the same shape, the AI can pull them together and reason across all of them at the same time.
So TikTok Ads never has to be looked at in isolation. The same question can pull in everything else you've connected:
The rest of your paid channels (Meta, Google, LinkedIn, and more), so you can compare what TikTok is doing against the channels you're already running.
Your SEO and search visibility data (how you're showing up in organic search and AI answers), sitting right next to your paid performance.
The platforms further down the funnel (like email and your store), so spend on one channel can be connected to what actually happened afterward.
That turns simple reporting into genuine cross-channel intelligence. Instead of "how did TikTok do?", you can ask the harder, more useful questions:
"Is TikTok actually bringing in new audiences, or just reaching the same people I'm already paying for on Meta and Google?"
"When my TikTok spend goes up, does my branded search and organic visibility move with it?"
"Which channel is really earning the credit for last month's results, and where should the next dollar go?"
Here's what that looks like in practice. Ask "how did my TikTok campaigns do last month, and are they finding new customers?" and Fuse doesn't just echo a metric back. It comes back with the analysis: your TikTok ads returned a 4.1x ROAS, your best paid channel, and 41% of the people they reached were brand-new, not audiences you're already paying for on Meta or Google, so it recommends shifting more budget into TikTok prospecting. Every figure in that answer is read straight from your saved data, and the AI never made one up.
Answering questions like that means lining up data from several platforms that each speak their own dialect, normalizing it, and reasoning over the whole picture. That's exactly what the per-platform-brain design makes possible, and it's what turns Fuse from a reporting tool into a genuinely advanced AI analyst for your entire marketing stack.
Why this matters if you run marketing
You don't need to care about access tokens or sandboxes to feel the difference:
You can trust every number. Answers are grounded in your real, saved data, never guessed at. Zero hallucinations, by design.
You get an analyst, not just a chart. Fuse doesn't only report what happened, it tells you what to do next, from net-new audience analysis to budget recommendations.
Common questions are instant and free. The things you ask all the time come back immediately without spending credits.
The whole picture, not one channel at a time. Because every platform feeds the same clean datasets, you can cross-reference TikTok against the rest of your paid channels and your SEO visibility in a single question, the kind of analysis that normally takes an analyst and a pile of spreadsheets.
Want to see this on your own data? Browse the full list of integrations , compare pricing , or book a demo .
The trusted solution for converting marketing
data into revenue.
