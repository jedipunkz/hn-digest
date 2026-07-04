---
source: "https://daringfireball.net/2026/07/claudes_criminally_bad_mac_app_is_an_inside_job"
hn_url: "https://news.ycombinator.com/item?id=48784469"
title: "Claude's Criminally Bad Electron Mac App Is an Inside Job"
article_title: "Daring Fireball: Claude’s Criminally Bad Electron Mac App Is an Inside Job"
author: "tosh"
captured_at: "2026-07-04T11:34:29Z"
capture_tool: "hn-digest"
hn_id: 48784469
score: 1
comments: 0
posted_at: "2026-07-04T10:59:14Z"
tags:
  - hacker-news
  - translated
---

# Claude's Criminally Bad Electron Mac App Is an Inside Job

- HN: [48784469](https://news.ycombinator.com/item?id=48784469)
- Source: [daringfireball.net](https://daringfireball.net/2026/07/claudes_criminally_bad_mac_app_is_an_inside_job)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T10:59:14Z

## Translation

タイトル: クロードの犯罪的に悪質な Electron Mac アプリは内部犯行
記事タイトル: Daring Fireball: クロードの犯罪的に悪質な Electron Mac アプリは内部犯行
説明: Felix Rieseberg が、なぜ Claude が Electron アプリなのかという質問に対する答えであることは明らかです。それは、なぜ建物のすべてのネジが壁に打ち込まれているのか不思議に思った後、建設を監督した男が世界最大のハンマー製造会社を設立し共同所有していることを知ったようなものです。

記事本文:
Day One — 実際に付けている日記。チャットで始まり、日記エントリで終わります。 ⭐ 4.8 (400k)
クロードの犯罪的に悪質な Electron Mac アプリは内部犯行である
Anthropic は、2024 年 10 月に MacOS 用の Claude 「デスクトップ」アプリの最初のバージョンをリリースしました。これは、UI デザイナーに感銘を与えなかった Electron のポンコツでした。それが出たとき、私は次のように書きました。
一方、ChatGPT のネイティブ Mac アプリは、まさに
ネイティブの Mac アプリ。 Mac アプリのような見た目と Mac アプリのような操作感
それは本当にMacアプリだからです。それ以来ずっと好きです
5 月に発売されましたが、改良され続けています。そして
私は回答するための頼りになるリソースとしてこれをますます使用し続けています
質問。
私はクロードに尋ねました。「ネイティブ Mac を設計する最良の方法は何ですか?」
アプリ？次の場合、どのようなフレームワークと開発者ツールを使用する必要がありますか?
目標は素晴らしい Mac エクスペリエンスですか?」クロードの答えは
それは、SwiftUI と AppKit の間の決定であると仮定することから始まりました。
おそらく、Anthropic の Mac エンジニアはクロードにこれを尋ねるべきだったでしょう。
Electron アプリのこのクソを構築する前に同じ質問。
今年の 3 月、Claude Code と Claude Cowork が Mac を制御してエージェント タスクを実行できるという Anthropic の発表に関連して、私は同じ質問に戻りました。
Claude Mac クライアント自体は、依然として怠惰な Electron ポンコツのままです。もし
クロード・コードはとても優れているので、なぜそれを証明しないのか理解できません
それを使用して、中途半端でもまともなネイティブ Mac アプリを作成します。
これについて考えたことがあるのは私だけではありません。 Drew Breunig は今年 2 月に「 Why is Claude an Electron App? 」と書きました。
表面的には、この能力はエレクトロンの利益をもたらすはずです
時代遅れ！ 1 つの Web アプリを作成してそれぞれに配布するのではなく、
プラットフォームでは、1 つの仕様とテスト スイートを作成し、コーディングを使用する必要があります。
エージェントがネイティブ コードを各プラットフォームに配布します。このアビなら

真実は
本物で採用されているため、ユーザーは迅速でパフォーマンスの高いネイティブ アプリを
小規模で集中的なチームが幅広い市場にサービスを提供します。
しかし、私たちは依然として Electron に依存しています。人間的でさえ、
派手なエージェントを公開し続ける AI コーディング ツールのリーダー
コーディングの実績はありますが、依然として Claude デスクトップで Electron を使用しています
アプリ。そして、それは遅く、バグが多く、肥大化したアプリです。
では、なぜ私たちは依然として Electron を使用し、
エージェント主導の仕様主導型開発の未来は?
まず、コーディング エージェントは最初の 90% において非常に優れています。
開発者。しかし最後の部分は、すべてのエッジケースを特定し、
現実の世界に出会ってもサポートを継続するのは難しいままですが、
面倒で、多くのエージェントの手を握る必要があります。 [...]
今のところ、Electron はまだ意味があります。コーディングエージェントは素晴らしいです。
しかし、開発とサポートの最後の 1 マイル領域は依然として重要な領域です。
本当の懸念。
私は、最後の 10% で苦戦しているコーディング エージェントを、Mac アプリの作成に Electron を選択する正当な理由として彼が受け入れるところまで、Breunig 氏に協力しました。個人でもチームでも、多くの人が Claude Code を使用して素晴らしい新しいネイティブ Mac アプリを作成しています。私の友人でも、グレン・フライシュマン氏、レックス・フリードマン氏、そしてジェイソン・スネル氏は、ここ数カ月間、一般的な AI コーディング アシスタントだけでなく、特にクロード コードを使用して、数十年にわたる文字通りプロフェッショナルな Mac の俗物化によって鍛えられた、Mac に対する個人的な高い基準を満たす真のネイティブ Mac アプリを作成しました。 Claude Code の支援を受けて作成された Mac 対応アプリの包括的なカタログは、おそらく非常に長くなるのではないかと思います。
最後の 10% との闘いは AI コーディングとは無関係です。それはすべてのソフトウェア エンジニアリングの性質です。ウィキペディアで「99 十規則」と名付けられた有名な格言があります。

ベル研究所のトム・カーギルに敬意を表します。
コードの最初の 90 パーセントが最初の 90 パーセントを占めます
開発当時のこと。コードの残りの 10%
開発時間の残りの 90% を占めます。
カーギルの数学的にユーモラスな定式化は、最後の 10 パーセントに時間が半分かかる理由だけでなく、ソフトウェア プロジェクトに予想の 2 倍の時間がかかる傾向がある理由も説明しているため、共感を呼びます。この普遍的な真実は、コードが人間によって書かれたものであっても、AI によって生成されたものであっても、あるいはその両方が混在したものであっても当てはまります。
Breunig 氏は追記で真実に近づいており、彼の投稿について議論している Hacker News スレッドにリンクしています。 HN スレッドで最も評価の高いコメントは、Anthropic の Claude Code チームで働く Boris Cherny によるものです。チャーニーはこう書きました：
Claude Code チームの Boris です。
アプリに取り組んでいるエンジニアの一部は Electron back に取り組んでいました
日中は、非ネイティブの建物を好んだ。それも素敵ですね
コードを共有する方法により、Web 全体で機能が保証される
とデスクトップは同じ外観と操作性を持ちます。最後にクロードは、
上手です。
とはいえ、エンジニアリングはトレードオフがすべてであり、これは変わる可能性があります
将来的には！
私は、「Web とデスクトップの機能が同じルック アンド フィールを持つ」という保証を、Mac アプリが Web の制限によって制限され、Mac のネイティブ フレームワークによって提供される幅広い慣用的なネイティブ機能から切り離されていることを保証すると言い換えたいと思います。 Electron は、アプリがすべてのプラットフォームで同じように間違っていると感じることを保証します。しかし、より関連性のある情報は、次の文です。「アプリに取り組んでいるエンジニアの一部は、昔は Electron で働いていたため、非ネイティブで構築することを好みました。」つまり、クロードがどういうわけか Electron を好むのではなく、Anthropic の「一部のエンジニア」がそうしているのです。
激しい運動をしている人もいます

現在 Anthropic の Claude Cowork および Claude Code Desktop のエンジニアリング リードであり、以前は MacOS および Windows 用の Claude アプリのエンジニアリング リードを務めていた Felix Rieseberg が「エンジニアの一部」に含まれていることを考慮すると、そこに留まります。リーゼバーグ氏は単に「昔、エレクトロンに取り組んでいた」だけではありません。彼は Electron の創設に責任を負う主要人物の 1 人であり、現在も「ガバナンスとプロジェクト全体を監督する」Electron プロジェクトの管理ワーキング グループの 3 人のメンバーのうちの 1 人です。彼は文字通りエレクトロンに関する本を書きました。
Felix Rieseberg が、なぜ Claude が Electron アプリなのかという質問に対する答えであることは明らかです。それは、なぜ建物のすべてのネジが壁に打ち込まれているのか不思議に思った後、建設を監督した男が世界最大のハンマー製造会社を設立し共同所有していることを知ったようなものです。 Windows はプラスネジを使用し、Linux は六角ネジを使用し、MacOS にはトルクス (もちろん) が必要ですが、ハンマーはどのネジでも同じように機能します。それがエレクトロンです。それはリーゼベルクの赤ちゃんです。
リーゼバーグは、クロードが Electron アプリになることに関与しただけではないことが判明しました。彼の個人ホームページと LinkedIn プロフィールによると、Anthropic に入社する前は、Notion のデスクトップ チームのエンジニアリング マネージャーとして 2 年以上過ごしました。Notion の Mac 用クライアントは、518 MB の巨大な Electron アプリであり、非ネイティブなエクスペリエンスであることで悪名高いです。 1 Notion に入社する前、Rieseberg は 2016 年から 2021 年を「Slack でシニア スタッフ エンジニアおよびエンジニアリング マネージャーとして過ごしました。そこで私は、クロスプラットフォームのデスクトップ フレームワーク Electron と macOS、Windows、Linux 用の Slack デスクトップ アプリを構築する素晴らしい C++ エンジニアのチームをサポートすることができました。」 2
Electron の上級メンテナである 1 人の男性がデスクトップのチームを率いていることを知った

Slack や Notion のクライアントは、今ではクロードが、タイタニック号を操縦し、ヒンデンブルク号を操縦し、その後アメリア・イアハートの航空管制官を務めたのは、家業が蒸留所だった一人の男だったことを発見したような気分です。
指摘しておく価値があるのは、Notion がおそらく彼らのやり方の誤りに気づいているということです。 Appleは先月のWWDCでのプラットフォーム一般教書基調講演でNotionを大きく取り上げ、「以前はNotionのようなクロスプラットフォームやWebテクノロジーを使用していたアプリは、他のテクノロジーでは提供できないレベルのパフォーマンスとUIの一貫性を求めて、ユーザーインターフェイスをSwiftUIに移行している」と述べた。これは、リーゼバーグがアントロピック社に去ってからわずか 1 年後のことです。おそらくクロードも同様に、最終的にはクロード アプリから電子の悪臭を洗い流そうとするでしょう。ただし、Electron のコードベースは樹液のようなものではないかと思います。粘着性があり、汚れており、時間が経てば経つほど洗い流すのが難しくなります。 ↩︎
Slack のすべての「デスクトップ」アプリを担当するエンジニアリング マネージャーに昇進する前に、Rieseberg 氏は Windows のエンジニアリング チーム リーダーとして Slack に入社しました。このことから、彼のプラットフォームの好みが垣間見えます。 ↩︎︎
表示設定
著作権 © 2002–2026 The Daring Fireball Company LLC。

## Original Extract

Felix Rieseberg, quite obviously, is the answer to the question why Claude is an Electron app. It’s like wondering why all the screws in a building were hammered into the walls, and then finding out that the guy who oversaw construction founded and co-owns the world’s biggest hammer manufacturer.

Day One — The journal you actually keep. Start with a chat, end with a journal entry. ⭐ 4.8 (400k)
Claude’s Criminally Bad Electron Mac App Is an Inside Job
Anthropic released the first version of the Claude “desktop” app for MacOS in October 2024 — an Electron clunker that did not impress UI designers . When it came out, I wrote :
ChatGPT’s native Mac app , on the other hand, is a truly
native Mac app. It looks like a Mac app and feels like a Mac app
because it really is a Mac app. I’ve liked it ever since it
launched back in May , and it keeps getting better. And
I keep using it more and more as my go-to resource for answering
questions.
I asked Claude, “What is the best way to engineer a native Mac
app? What frameworks and developer tools should one use if the
goal is a great Mac experience?” Claude’s answer
started by positing it as a decision between SwiftUI and AppKit.
Perhaps Anthropic’s Mac engineers should have asked Claude this
same question before they built this turd of an Electron app.
In March of this year, linking to Anthropic’s announcement that Claude Code and Claude Cowork can take control of your Mac to accomplish agentic tasks, I returned to the same question :
The Claude Mac client itself remains a lazy Electron clunker. If
Claude Code is so good I don’t get why they don’t prove it by
using it to make an even halfway decent native Mac app.
I’m not the only one who has pondered this. Drew Breunig wrote “ Why is Claude an Electron App? ” in February this year:
On the surface, this ability should render Electron’s benefits
obsolete! Rather than write one web app and ship it to each
platform, we should write one spec and test suite and use coding
agents to ship native code to each platform. If this ability is
real and adopted, users get snappy, performant, native apps from
small, focused teams serving a broad market.
But we’re still leaning on Electron. Even Anthropic, one of the
leaders in AI coding tools, who keeps publishing flashy agentic
coding achievements, still uses Electron in the Claude desktop
app. And it’s a slow, buggy, and bloated app.
So why are we still using Electron and not embracing the
agent-powered, spec driven development future?
For one thing, coding agents are really good at the first 90% of
dev. But that last bit — nailing down all the edge cases and
continuing support once it meets the real world — remains hard,
tedious, and requires plenty of agent hand-holding. [...]
For now, Electron still makes sense. Coding agents are amazing.
But the last mile of dev and the support surface area remains a
real concern.
I’m with Breunig up until the point where he accepts coding agents struggling with the final 10 percent as a justification for choosing Electron to create a Mac app. Plenty of people — individuals and teams alike — are using Claude Code to create terrific new native Mac apps. Just among my friends, Glenn Fleishman , Lex Friedman , and Jason Snell man , have all in recent months used not just AI coding assistants in general, but Claude Code specifically, to create genuinely native Mac apps that meet their own personal high standards for Mac-assedness, forged through decades of literally professional Mac snobbery. A comprehensive catalog of Mac-assed apps made with the assistance of Claude Code, would, I suspect, be remarkably long.
The struggle with the last 10 percent is unrelated to AI coding. It’s the nature of all software engineering. There’s a well-known adage that Wikipedia names the “ Ninety-Ninety Rule ”, attributed to Tom Cargill of Bell Labs:
The first 90 percent of the code accounts for the first 90 percent
of the development time. The remaining 10 percent of the code
accounts for the other 90 percent of the development time.
Cargill’s mathematically humorous formulation resonates because it not only explains why the final 10 percent consumes half the time, but also why software projects tend to take twice as long as expected. This universal truth holds whether the code is human-written, AI-generated, or a mix of both.
Breunig gets closer to the truth in a postscript, linking to the Hacker News thread discussing his post . The top-rated comment in the HN thread is from Boris Cherny , who works at Anthropic on the Claude Code team. Cherny wrote :
Boris from the Claude Code team here.
Some of the engineers working on the app worked on Electron back
in the day, so preferred building non-natively. It’s also a nice
way to share code so we’re guaranteed that features across web
and desktop have the same look and feel. Finally, Claude is
great at it.
That said, engineering is all about tradeoffs and this may change
in the future!
I would rephrase the guarantee that “features across web and desktop have the same look and feel” as guaranteeing that the Mac app is restrained by the limits of the web and cut off from the breadth of idiomatically native functionality provided by the Mac’s native frameworks. Electron guarantees that an app feels just as wrong on all platforms. But the more relevant tidbit is this sentence: “Some of the engineers working on the app worked on Electron back in the day, so preferred building non-natively.” So it’s not that Claude somehow prefers Electron, but that “some of the engineers” at Anthropic do.
Some is doing some heavy lifting there, given that “some of the engineers” includes Felix Rieseberg , currently Anthropic’s engineering lead for Claude Cowork and Claude Code Desktop, and previously engineering lead for the Claude apps for MacOS and Windows. Rieseberg didn’t merely “work on Electron back in the day”. He is one of the principal people responsible for creating Electron, and remains today one of three members of the Electron project’s Administrative Working Group that “oversees the entire governance and project”. He literally wrote the book on Electron .
Felix Rieseberg, quite obviously, is the answer to the question why Claude is an Electron app. It’s like wondering why all the screws in a building were hammered into the walls, and then finding out that the guy who oversaw construction founded and co-owns the world’s biggest hammer manufacturer. Windows uses Philips head screws, Linux uses hex screws, and MacOS requires Torx ( of course ) — but a hammer works the same way with all screws. That’s Electron. That’s Rieseberg’s baby.
Rieseberg, it turns out, hasn’t only had a hand in Claude being an Electron app. Per both his personal home page and LinkedIn profile , before joining Anthropic he spent over two years as the engineering manager for the desktop team at Notion, whose client for Mac is a massive 518 MB Electron app and a notoriously non-native experience. 1 Before Notion, Rieseberg spent 2016–2021 “as a Senior Staff Engineer and Engineering Manager at Slack, where I got to support a team of amazing C++ engineers building the cross-platform desktop framework Electron — as well as Slack’s desktop apps for macOS, Windows, and Linux.” 2
Finding out that one guy — who is a senior Electron maintainer — has led the teams for the desktop clients for Slack, Notion, and now Claude is like discovering that it was one guy — whose family business was a distillery — who helmed the Titanic, piloted the Hindenburg, and then served as air traffic controller for Amelia Earhart.
Notion, it’s worth pointing out, has perhaps seen the error of their ways. Apple prominently featured Notion during the Platforms State of the Union technical keynote at WWDC last month, saying, “And apps that previously used cross-platform or web technologies like Notion are migrating their user interface to SwiftUI because they want a level of performance and UI consistency that other technologies can’t deliver.” This, just one year after Rieseberg left for Anthropic. Perhaps Claude will similarly seek to wash the Electron stink off the Claude app eventually. I suspect an Electron codebase is like sap, though — sticky, dirty, and harder to wash off the longer you wait. ↩︎
Before getting promoted to engineering manager in charge of all of Slack’s “desktop” apps, Rieseberg started at Slack as engineering team lead for Windows , which offers an inkling as to his platform taste. ↩︎︎
Display Preferences
Copyright © 2002–2026 The Daring Fireball Company LLC.
