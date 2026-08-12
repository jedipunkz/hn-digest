---
source: "https://newsletter.pragmaticengineer.com/p/stop-being-skeptical-about-ai-for"
hn_url: "https://news.ycombinator.com/item?id=49275503"
title: "Stop being skeptical about AI for development"
article_title: "Stop being skeptical about AI for development with Charity Majors"
author: "duck"
captured_at: "2026-08-12T17:52:20Z"
capture_tool: "hn-digest"
hn_id: 49275503
score: 2
comments: 0
posted_at: "2026-08-12T17:02:54Z"
tags:
  - hacker-news
  - translated
---

# Stop being skeptical about AI for development

- HN: [49275503](https://news.ycombinator.com/item?id=49275503)
- Source: [newsletter.pragmaticengineer.com](https://newsletter.pragmaticengineer.com/p/stop-being-skeptical-about-ai-for)
- Score: 2
- Comments: 0
- Posted: 2026-08-12T17:02:54Z

## Translation

タイトル: 開発用 AI について懐疑的になるのをやめる
記事のタイトル: 慈善メジャーによる開発のための AI について懐疑的になるのをやめる
説明: Honeycomb の CTO 兼共同創設者である Charity Majors が、AI がソフトウェア エンジニアリングをどのように再構築し、信頼性と検証がこれまで以上に重要になっているかについて説明します。

記事本文:
Charity Majors で AI の開発に懐疑的になるのはやめましょう
実践的なエンジニア
購読 サインイン プラグマティック エンジニア Charity Majors による開発のための AI について懐疑的になるのをやめる 66 1 1× 0:00 現在時刻: 0:00 / 合計時刻: -1:25:32 -1:25:32 お使いのブラウザではオーディオ再生がサポートされていません。アップグレードしてください。 Charity Majors で AI の開発に懐疑的になるのはやめましょう
Gergely Orosz 2026 年 8 月 12 日 66 1 シェア トランスクリプト 最新のエピソードをストリーミング
YouTube 、Apple 、Spotify で今すぐ聴いたり見たりできます。このページの上部にあるエピソードのトランスクリプトと下部にあるエピソードのタイムスタンプをご覧ください。
• アンチテーゼ – 積極的なフォールト挿入下でシステム全体を実行することによる、システムのターボチャージ テスト。 Jane Street、Fly.io、etcd コミュニティなどのチームが Antithesis を利用しているのには十分な理由があります。もっと詳しく知る。
• Buildkite – OpenAI、Anthropic、Cursor、Meta、Uber、NVIDIA、Airbnb などによって信頼されている CI プラットフォーム。 CI ボリュームがアーキテクチャの問題になる場合は、より優れた CI が必要になります。現在、来年、そしてそれ以降も、コーディング エージェントがビルド キューに投入するあらゆるものを確実に管理できるように設計されています。もっと詳しく知る。
• WorkOS – SSO、SCIM、RBAC などを使用して、アプリとエージェントをエンタープライズ対応にします。始めましょう。
2025 年には AI について懐疑的になるのが合理的でしたが、2026 年には AI が業界全体を変えつつあることは明らかであり、懐疑的な余地はますます少なくなっています。この意見は、ソフトウェアの信頼性と可観測性に関する私のお気に入りの意見の 1 つである、Charity Majors 氏、Honeycomb の CTO 兼共同創設者、Observability Engineering の共著者によるものです。 (注: 『Observability Engineering』の第 2 版が出版されていますが、これは本書をほぼ完全に書き直したものです。信頼性の高いシステムを構築している場合は、本書を入手することをお勧めします)
で

このエピソードでは、私は Charity と対談し、AI に関する彼女の考え方がどのように進化してきたのか、AI がソフトウェア エンジニアリングの基礎的な部分になりつつあると彼女が考える理由、そしてそれがチームがソフトウェアを構築、レビュー、出荷する方法に何を意味するのかについて話し合いました。
AI がコード生成の経済性をどのように変えているのか、なぜ信頼性と検証がますますボトルネックになっているのか、なぜ非決定論的システムの台頭によりより多くのエンジニアリング規律が必要なのかを探ります。 Charity は、コード レビュー、可観測性、DevOps、リーダーシップ、そして AI 懐疑論者と愛好家の両方が重要なことを正しく理解している理由について、彼女の見解を共有しています。
慈善団体との会話から得たポイント
Charity との会話で私が特に興味深いと感じた 13 の部分を紹介します。
1. 2025 年 3 月、Charity は SREcon の聴衆にバイブコーディングを試すよう呼びかけましたが、当時の反応は不平不満でした。慈善団体の言いたいことは、AI に懐疑的な人でも、AI の使い方を学ぶべきだということだった。なぜなら、AI を学べばもっと文句を言えるようになるからだ。この時点で、Charity は AI が新しいプログラミング言語よりも大きな影響力を持つとまだ考えていましたが、AI が世代に影響を与えるかどうかについては懐疑的でした。
2. Charity が AI を世代交代とみなすターニングポイントは 2025 年 11 月でした。これは Opus 4.5 によるものですが、Charity はコーディング ハーネス (Claude Code) がより大きな違いをもたらしたと主張しています。クロード コードのおかげで、ハーネスは単なるシェル スクリプトから本格的なインフラストラクチャになったからです。
3. 2025 年の業界に対する AI の影響は、2010 年のクラウドの影響と同様でした。振り返ってみると、Charity は次のように言うことに自信があります。2010 年には、クラウド コンピューティングが確実に主流になり、インフラ層を変えるであろうことが明らかになりました。 2025 年以降、AI がソフトウェア構築のインフラストラクチャに同様の影響を与えることも明らかです

再。
4. 2025 年まで AI に懐疑的だったエンジニア: 彼らにはそうなるのには十分な理由がありました。これは、ソフトウェア業界に変革をもたらすと期待されながらも、その後は期待に届かなかったテクノロジーやイノベーションをこれまでに数多く見てきたからです。例としては、COBOL (ソフトウェア作成にプログラマーが不要になると約束されているテクノロジー)、ニューラル ネット、ノーコード ツールやローコード ツールなどがあります。
5. エンジニアが答える必要がある質問: 読んでいないコードを完全に快適に出荷するには何が必要ですか?慈善団体は、プロのソフトウェア エンジニアが一度も見ていない、つまり理解していないコードを本番環境に出荷するのは「いつ」であって、「もし」ではないと考えています。エンジニアリング部門は、このコードを検証し、完全な自信を持って出荷できるシステムを構築しています。
6. AI は、2010 年代にコンピューティング インフラストラクチャが経験した「ペット」から「牛」への変化をソフトウェア業界にもたらす可能性があります。これまで、ソフトウェアを最初から作成するのは、既存のソフトウェアを編集するよりもはるかに高価でした。しかし今では、関数の何百ものバリアントを生成する方が、一度手書きで書くよりも早く実行できるようになりました。
慈善団体は、私たちはハードウェア インフラストラクチャ層で起こった「ペット」から「牛」への移行の始まりにいるかもしれないと考えています。 2010 年代以前は、個々のサーバーの構成と修復が一般的に行われていました。しかし、Terraform や Kubernetes などのツールを使用すると、問題のある個々のサーバーは修正されなくなり、代わりに再作成されます。慈善団体は、遅かれ早かれ同じことがコードでも起こるかもしれないと考えています。コードに問題がある場合は、それを解決し、正しいことが検証された新しいコードを生成します。
7. 彼女の逆説的見解: コードレビューは過大評価されており、人間の仕事の最も価値の低い部分である

ソフトウェアエンジニアリングに追加します。慈善団体は、人間は会話をしたり、何を構築するかを決定するのが得意であり、コードを読んで正確さ、構文、バグを確認することは得意ではないと述べています。
8. 20 年間の DevOps に対する慈善団体の評決: それは失敗でした。 DevOps のフィードバックは、コードを作成する人々と本番環境で実行されているコードを結びつけるフィードバック ループを作成しようとするものでした。彼女は、「作戦担当者はコーディングを学びましょう!」と考えています。 wave は機能しましたが、「ソフトウェア エンジニア: 本番環境のコードを理解する」ことは、今日に至るまで失敗しました。
9. 非決定論的なシステムには、以前よりも多くのエンジニアリング規律が必要です。 AI によって書かれたコードでは、コードの信頼性が低下しているため (コードを書かなくなったため)、開発プロセスの他の部分での信頼性を高める必要があります。具体的には、テスト、評価、適合テストなどの検証時です。
10. エンジニアリング ディレクターに対する慈善団体からのキャリア アドバイス: 波に向かって走り、すぐに履歴書に AI を記載してください。 AI によるあらゆる変化のおかげで、テクノロジー業界で働くには不安な時期が来ています。慈善活動は、不安と興奮は生理学的にはほぼ同じであるが、違いは主体性であることを思い出させます。主体性がないときは不安になる可能性が高く、主体性があるときは興奮する可能性が高くなります。
そこで、不安を抱えるエンジニアリング ディレクターへの彼女のアドバイスは、より自由度の高い IC の仕事に戻ることを検討してくださいということです。 IC の取り組みは高く評価されており、IC への復帰はこれまでになく簡単になりましたが、そのための窓口は閉まりつつあります。彼女は次のように述べています。
「次回就職面接を受けるとき、AI の経験がない場合は除外されます。」
11. AI 疲労について: 小さな行動でコントロールを取り戻しましょう!私たちはさまざまなタイプの AI 疲労について話しました。AI の失敗を振り返ること、AI の誇大広告にうんざりすること、そして

AI CEO による「破滅荒らし」によって、その価値は損なわれています。慈善団体は、AI ツールを活用して仕事のコントロールを取り戻す小さな行為を見つけます。たとえば、Honeycomb チームは水曜日には AI を使用しません。
12. 慈善団体は、「AI に苦しめられた」陣営と「反 AI」陣営の両方がより良くストーリーを伝えることを望んでいます。彼女はこう言いました。
「現在、ソフトウェアでは本当に信じられないようなことが起こっています。たとえば、書き換えや自動化による労力の削減などです。私が話をした中で、AI の使用をやめようとする人は一人もいませんでした。
しかし、半数の人は成果を実感しており、それをコストに結び付けていないため、同僚は自動化が消滅することをただ恐れているだけだと考えています。
それで、これを聞いている皆さんにお願いです。すべてのストーリーを話してください!費用についても相談してください。私たち全員が一緒に取り組んでいます。」
13. AI ライティングに関する慈善団体のルール: 自分自身が全文を読んでいないメッセージや電子メールを人間に送信しないでください。彼女はまた、あなたが送ったものは何であれ、あなたが送信したものを読むのに、あなたがそれを作成するのにかかる時間よりも時間がかかるだろうとも言います。
このエピソードに関連するプラグマティック エンジニアの詳細
• 詳細調査: テクノロジー企業 10 社が次世代の開発ツールをどのように選択しているか
• なぜ Meta はエンジニアリング組織を破壊するのでしょうか?
• AI がほぼすべてのコードを作成すると、ソフトウェア エンジニアリングはどうなるでしょうか?
• AI エージェントは実際に私たちの速度を低下させているのでしょうか?
• 観察可能性: 慈善専攻による現在と未来
• ソフトウェア エンジニアリングの第 3 回黄金時代 – AI のおかげで、Grady Booch 氏とともに
02:56 Parse がどのようにして Honeycomb に至ったか
06:00 個人の生産性指標の限界
09:08 AI に対する慈善団体の視点はどのように進化したか
13:50 コードの書き換えとコードの編集
19:20 開発段階としての生産
26:56 非決定論的システム
44:4

0 AI がソフトウェアの構築に非常に適している理由
1:00:40 コンテキストのオーバーロードの処理
1:01:56 Observability Engineering 第 2 版の新機能
1:07:45 効果的なリーダーシップとはどのようなものか
1:10:25 エンジニアリング管理: 何が変化していますか?
• LinkedIn: https://www.linkedin.com/in/charity-majors
• 可観測性エンジニアリング、第 2 版: https://www.oreilly.com/library/view/observability-engineering-2nd/9781098179915
• ハニカム: https://www.honeycomb.io
• リンデンラボ: https://lindenlab.com
• セカンドライフ: https://secondlife.com
• 解析: https://en.wikipedia.org/wiki/Parse,_Inc 。
• スキューバ: https://research.facebook.com/publications/scuba-diving-into-data-at-facebook
• 本当に個々の開発者の生産性を測定できますか? - EM に質問します: https://blog.pragmaticengineer.com/can-you-measure-developer-productivity
• エージェントティック開発について話しましょう: Spotify x Anthropic ライブ: https://engineering.atspotify.com/2026/4/anthropic-agentic-development
• 疑わしいアドバイス: エンジニアリングの生産性は測定できますか?:
• 2025 年は AI にとって、2010 年はクラウドにとっては同じものでした。
• AI には、より多くのエンジニアリング規律が必要です。それ以上:
• フェニックスのアーキテクチャ: https://aicoding.leaflet.pub
• ソフトウェア エンジニアリングの第 3 黄金時代 – AI のおかげで、Grady Booch 氏とともに: https://newsletter.pragmaticengineer.com/p/the-third-golden-age-of-software
• Grady Booch によるソフトウェア アーキテクチャ: https://newsletter.pragmaticengineer.com/p/software-architecture-with-grady-booch
• Anders Hejlsberg 氏による TypeScript、C#、Turbo Pascal: https://newsletter.pragmaticengineer.com/p/typescript-c-and-turbo-pascal-with
• LinkedIn の David Poll: https://www.linkedin.com/in/depoll
• インターコム: https://www.intercom.com
• AI がプル リクエストを承認しています: 安全性を確保する方法は次のとおりです: https://www

.intercom.com/blog/ai-is-approving-our-pull-requests-heres-how-we-made-it-safe
• AI がソフトウェア エンジニアリングをどのように変えるか – Martin Fowler 氏: https://newsletter.pragmaticengineer.com/p/martin-fowler
• HackerRank は ATS をオープンソース化しました。私の履歴書は 90/100 点でした。ああ、お待ちください 74。いいえ – 88: https://news.ycombinator.com/item?id=48713832
• AI 愛好家は時間との戦い、AI 懐疑論者はエントロピーとの戦いです。
• 第 1 話#89、ソフトウェアはキラーアプリです (0xide Computer の Bryan Cantrill 氏): https://www.honeycomb.io/resources/podcasts/ep-89-bryan-cantrill-software-is-the-killer-app
• LinkedIn への Eric Riddoch の投稿: https://www.linkedin.com/posts/eric-riddoch_the-observability-engineering-book-has-share-7475807056285814785-pw4J
• 従来の可観測性が AI エージェントの障害を見逃す理由: https://www.dataiku.com/blog/traditional-observability-misses-ai-agent-failure
• 効果的なリーダーに関する慈善団体の LinkedIn 投稿: https://www.linkedin.com/posts/charity-majors_the-most-effective-leaders-are-kind-caring-share-7477160924928233472-qcLw
• 大惨事の倫理: 厳しい選択の世界で上手に選択する方法: https://www.amazon.com/dp/0593471970
• More Everything Forever: AI オーバーロード、宇宙帝国、人類の運命をコントロールするシリコンバレーの聖戦: https://www.amazon.com/More-Everything-Forever-Overlords-Humanity/dp/1541619595
制作・販売はペンネーム。
このエピソードはプラグムによって提供されます

[切り捨てられた]

## Original Extract

Charity Majors, CTO and co-founder of Honeycomb, explains how AI is reshaping software engineering and making reliability and verification more important than ever.

Stop being skeptical about AI for development with Charity Majors
The Pragmatic Engineer
Subscribe Sign in The Pragmatic Engineer Stop being skeptical about AI for development with Charity Majors 66 1 1× 0:00 Current time: 0:00 / Total time: -1:25:32 -1:25:32 Audio playback is not supported on your browser. Please upgrade. Stop being skeptical about AI for development with Charity Majors
Gergely Orosz Aug 12, 2026 66 1 Share Transcript Stream the latest episode
Listen and watch now on YouTube , Apple , and Spotify . See the episode transcript at the top of this page, and timestamps for the episode at the bottom.
• Antithesis – turbocharge testing of your systems by running your whole system under aggressive fault injection. There’s good reason teams like Jane Street, Fly.io, and the etcd community rely on Antithesis. Learn more.
• Buildkite – the CI platform trusted by OpenAI, Anthropic, Cursor, Meta, Uber, NVIDIA , Airbnb and many more. When CI volume becomes an architecture problem, you deserve better CI. Engineered to reliably manage whatever your coding agents throw at the build queue: today, next year, and beyond. Learn more.
• WorkOS – make your app and agents Enterprise Ready, with SSO, SCIM, RBAC, and more. Get started .
In 2025, it was rational to be skeptical about AI, but in 2026 it’s clear that AI is changing all of the industry, and there’s less and less place for skepticism. This take is from one of my favorite voices in software reliability and observability: Charity Majors, CTO and cofounder of Honeycomb, co-author of Observability Engineering . (Note: the second edition of Observability Engineering is out, and it’s pretty much a full rewrite of the book, I recommend grabbing it if you’re building reliable systems)
In this episode, I sat down with Charity to discuss how her thinking on AI has evolved, why she believes it is becoming a foundational part of software engineering, and what that means for how teams build, review, and ship software.
We explore how AI is changing the economics of code generation, why reliability and verification are increasingly the bottlenecks, and why the rise of non-deterministic systems requires more engineering discipline. Charity shares her views on code reviews, observability, DevOps, leadership, and why both AI skeptics and enthusiasts are getting important things right.
Takeaways from the conversation with Charity
Here are 13 parts I found especially interesting, talking with Charity:
1. In March 2025, Charity told the audience at SREcon to try vibe coding, and back then, the response was grumbling. Charity’s point was that people who are skeptical of AI should still learn to use it, because you can complain better if you’ve learned it. At this time, Charity still saw AI having a bigger impact than a new programming language, but was skeptical that it would have a generational impact.​
2. Charity’s turning point in seeing AI as a generational change was in November 2025. This was due to Opus 4.5, but Charity argues that the coding harness (Claude Code) made the bigger difference. Because thanks to Claude Code, harnesses went from being more of a shell script to serious infrastructure.​
3. The impact of AI on the industry in 2025 was similar to the impact of the cloud in 2010. Looking back, Charity is comfortable saying this: in 2010, it became clear that cloud computing was certainly going mainstream and would change the infra-layer. After 2025, it’s also clear that AI will have a similar impact on the infrastructure of building software.
4. Engineers who were skeptical of AI up to 2025: they had good reason to be so. This was because we’ve seen plenty of technologies and innovations in the past that all promised to transform the software industry, but later fell short. Examples include COBOL (a technology promising that programmers would no longer be needed to create software), neural nets, no-code and low-code tools.​
5. The question engineers need to answer: what would it take for you to be fully comfortable shipping code you have not read? Charity believes it is a “when” and not an “if” that professional software engineers will ship code they never looked at – and thus do not understand – to production. Engineering is building the systems that validate this code, and allow shipping with full confidence.
6. AI could have the software industry go through the “pets” to “cattle” change that compute infra went through in the 2010s. Up to now, writing software from scratch was far more expensive than editing existing software. But now, generating hundreds of variants of a function can be done faster than how long it would take you to hand-write it once.
Charity believes that we might be at the beginning of the transition from “pets” to “cattle” that happened at the hardware infrastructure layer. Before the 2010s, configuring and repairing individual servers was commonly done. But with tools like Terraform and Kubernetes, individual servers having issues are no longer fixed up: they are re-created instead. Charity thinks the same might happen with code, sooner rather than later. When there’s an issue with the code, generate new code that solves it, and is verifyably correct.​
7. Her contrarian take: code review is overrated, and the least valuable part of what humans add to software engineering. Charity says that humans are good at conversations and deciding what to build, not reading code to check for correctness, syntax and bugs.​
8. Charity’s verdict of 20 years of DevOps: it failed. The DevOps feedback was about trying to create a feedback loop that connected people writing the code to the code running in production. She thinks that the “ops people: learn to code!” wave worked, but the “software engineers: understand your code in production” failed, to this day.
9. Non-deterministic systems require more engineering discipline versus before. With code written by AI, we’re reducing the trust in the code (because we no longer wrote it), so we need to increase trust at the other part of the development process. Specifically, at validation: with things like tests, evals, and conformance testing.​
10. Charity’s career advice for engineering directors: run towards the waves, and get AI on your resume, immediately. It’s an anxious time to work in tech, thanks to all the change, driven by AI. Charity reminds us that anxiety and excitement are physiologically almost the same, but the difference is agency. When you have no agency, you’re more likely to get anxious, and when you do, you’re more likely to get excited.
So her advice to anxious engineering directors: consider going back to IC work, where you’ll have far more agency. IC work is well-respected, getting back to it has never been easier, but the window to do so is closing. As she put it:​
“The next time you’ll have a job interview, you’ll be filtered out if you don’t have AI experience.”​
11. On AI fatigue: take back control with small acts! We talked about various types of AI fatigue: reviewing AI slop, getting tired of the AI hype, and getting worn down by “doom trolling” by AI CEOs. Charity finds small acts of taking control back in your work from AI tools help. For example, none of the Honeycomb team uses AI on Wednesdays.​
12. Charity would like to see both the “AI-pilled” and the “anti-AI” camps tell the stories better. As she put it:
“There are some really incredible things happening in software right now, for example, with rewrites and with automating away toil. Not a single person that I’ve talked to would give up using AI.
But half of the people are seeing the wins, and they’re not connecting it to the cost , which makes them think that their coworkers are just afraid of getting automated out of existence.
So that’s my beg to everyone who listens to this: tell the whole story! Talk about the costs as well. We’re all in it together.”
13. Charity’s rule on AI writing: do not send any message/email to a human that you yourself have not read in full. She also says that it would take them longer to read whatever you send than it took you to produce it: it’s probably slop!
The Pragmatic Engineer deepdives relevant for this episode
• Deepdive: How 10 tech companies choose the next generation of dev tools
• Why is Meta destroying its engineering organization?
• When AI writes almost all code, what happens to software engineering?
• Are AI agents actually slowing us down?
• Observability: the present and future , with Charity Majors
• The third golden age of software engineering – thanks to AI , with Grady Booch
02:56 How Parse led to Honeycomb
06:00 The limits of individual productivity metrics
09:08 How Charity’s perspective on AI has evolved
13:50 Rewriting code vs. editing code
19:20 Production as a stage of development
26:56 Non-deterministic systems
44:40 Why AI works so well for building software
1:00:40 Handling context overload
1:01:56 What’s new in Observability Engineering’s 2nd edition
1:07:45 What effective leadership looks like
1:10:25 Engineering management: what is changing?
• LinkedIn: https://www.linkedin.com/in/charity-majors
• Observability Engineering, 2nd Edition: https://www.oreilly.com/library/view/observability-engineering-2nd/9781098179915
• Honeycomb: https://www.honeycomb.io
• Linden Lab: https://lindenlab.com
• Second Life: https://secondlife.com
• Parse: https://en.wikipedia.org/wiki/Parse,_Inc .
• Scuba: https://research.facebook.com/publications/scuba-diving-into-data-at-facebook
• Can You Really Measure Individual Developer Productivity? - Ask the EM: https://blog.pragmaticengineer.com/can-you-measure-developer-productivity
• Let’s Talk Agentic Development: Spotify x Anthropic Live: https://engineering.atspotify.com/2026/4/anthropic-agentic-development
• Questionable Advice: Can Engineering Productivity Be Measured?:
• 2025 was for AI what 2010 was for cloud:
• AI demands more engineering discipline. Not less:
• The Phoenix Architecture: https://aicoding.leaflet.pub
• The third golden age of software engineering – thanks to AI, with Grady Booch: https://newsletter.pragmaticengineer.com/p/the-third-golden-age-of-software
• Software architecture with Grady Booch: https://newsletter.pragmaticengineer.com/p/software-architecture-with-grady-booch
• TypeScript, C# and Turbo Pascal with Anders Hejlsberg: https://newsletter.pragmaticengineer.com/p/typescript-c-and-turbo-pascal-with
• David Poll on LinkedIn: https://www.linkedin.com/in/depoll
• Intercom: https://www.intercom.com
• AI is approving our pull requests: Here’s how we made it safe: https://www.intercom.com/blog/ai-is-approving-our-pull-requests-heres-how-we-made-it-safe
• How AI will change software engineering – with Martin Fowler: https://newsletter.pragmaticengineer.com/p/martin-fowler
• HackerRank open sourced its ATS. My resume scored 90/100. Oh wait 74. No – 88: https://news.ycombinator.com/item?id=48713832
• AI enthusiasts are in a race against time, AI skeptics are in a race against entropy:
• Ep. #89, Software is the Killer App with Bryan Cantrill of 0xide Computer: https://www.honeycomb.io/resources/podcasts/ep-89-bryan-cantrill-software-is-the-killer-app
• Eric Riddoch’s post on LinkedIn: https://www.linkedin.com/posts/eric-riddoch_the-observability-engineering-book-has-share-7475807056285814785-pw4J
• Why traditional observability misses AI agent failure: https://www.dataiku.com/blog/traditional-observability-misses-ai-agent-failure
• Charity’s LinkedIn post on effective leaders: https://www.linkedin.com/posts/charity-majors_the-most-effective-leaders-are-kind-caring-share-7477160924928233472-qcLw
• Catastrophe Ethics: How to Choose Well in a World of Tough Choices: https://www.amazon.com/dp/0593471970
• More Everything Forever: AI Overlords, Space Empires, and Silicon Valley’s Crusade to Control the Fate of Humanity: https://www.amazon.com/More-Everything-Forever-Overlords-Humanity/dp/1541619595
Production and marketing by Pen Name .
THIS EPISODE IS PRESENTED BY The Pragm

[truncated]
