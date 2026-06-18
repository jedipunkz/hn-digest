---
source: "https://cloudfour.com/thinks/improvements-to-web-for-ai-should-benefit-all-users/"
hn_url: "https://news.ycombinator.com/item?id=48584418"
title: "Improvements to Web for AI Should Benefit All Users"
article_title: "Improvements to Web for AI Should Benefit All Users – Cloud Four"
author: "speckx"
captured_at: "2026-06-18T13:17:42Z"
capture_tool: "hn-digest"
hn_id: 48584418
score: 2
comments: 0
posted_at: "2026-06-18T12:37:44Z"
tags:
  - hacker-news
  - translated
---

# Improvements to Web for AI Should Benefit All Users

- HN: [48584418](https://news.ycombinator.com/item?id=48584418)
- Source: [cloudfour.com](https://cloudfour.com/thinks/improvements-to-web-for-ai-should-benefit-all-users/)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T12:37:44Z

## Translation

タイトル: AI のための Web の改善はすべてのユーザーに利益をもたらすはずです
記事のタイトル: AI 向け Web の改善はすべてのユーザーに利益をもたらすべき – Cloud Four
説明: AI をサポートするために Web プラットフォームに提案されている変更では、可能な限りユーザーと支援テクノロジーもサポートする必要があります。

記事本文:
AI のための Web の改善はすべてのユーザーに利益をもたらすべき – Cloud Four
メインコンテンツにスキップ
ホーム: クラウド フォー
メインメニュー
オフラインのようです。一部のコンテンツは利用できない可能性があります。
アラートを閉じる
AI 向けの Web の改善はすべてのユーザーに利益をもたらすはずです
2026 年 6 月 15 日発行
2026 年 6 月 15 日
先週、Safari チームは WebMCP API に対する反対を正式に表明しました。 Safari が WebMCP に反対する理由についての説明の中で、彼らは私が明確に表現しようとしていたものを捉えていました。
Safari チームの立場を説明するコメントの中で、Mike Wyrzykowski は次のように書いています。
サイトのアクションがエージェントにとって使いにくい場合、それはページ自体のセマンティクスにギャップがあり、その問題を解決するには、プラットフォームの共有レイヤー (HTML および ARIA) 内でギャップを閉じることだと考えられます。そこでは、ユーザー、支援テクノロジー、エージェントのすべてが利益を得ることができます。
実際、この感情は、AI エージェントをより適切にサポートするために Web を変更するあらゆる取り組みの指針となるはずです。 AI のすべての機能強化が、Web のすべての共有レイヤーをサポートするように構成できるわけではない可能性がありますが、Web がサービスを提供すべき人間ではなく、AI のみに利益をもたらすソリューションに到達する前に、まずそうするよう努めるべきです。
私たちのより深い関心は建築的なものです。ユーザーに代わって機能するエージェントは、実質的には支援テクノロジーです。エージェントはユーザーと同じようにサイトを操作する必要があり、サイトはそのエージェントを別の処理のために選択すべきではありません。
この提案では、WebMCP は「アクセシビリティ テクノロジによる取り込みを目的として設計されていない」と述べていますが、これは同じ分岐点であり、より豊富で実用的なセマンティクスがエージェントに届く一方で、スクリーン リーダーやキーボード ユーザーの受け取りは少なくなります。
これらすべてを考えると、AI がドメイン権限を理解するのに役立つ標準案について、Web Incubator Community Group (WICG) メーリング リストで最近議論されたことを思い出します。提案された標準は設計されています

次の問題を解決するには:
AI アシスタントは、一貫した機械可読でドメイン管理された ID データ ソースがないため、ドメインを誤認したり偽ったりすることがよくあります。現在、モデルはスクレイピングされたページ、一貫性のないメタデータ、サードパーティのアグリゲーター、または古いインデックスに依存しています。ドメインが誰であるか、何を表すか、またはどのリソースが権威であるかを宣言できる正規の場所はありません。
提案されたソリューションでは、問題を解決するためにすべての Web サイトに JSON-LD ドキュメントを追加することを提案しています。
人間はドメイン権限とも格闘します。人々は、どの Web サイトが正規で、どの Web サイトが正規ではないかを判断するのが困難です。 AI に関してこの問題を解決しているのであれば、おそらくエンドユーザーにも役立つソリューションを見つけることができるでしょう。
そして、AI が私たちのすべての仕事を置き換えて超知能になるはずであると同時に、AI がウェブを使用するために特別な AI 補助輪を追加する必要があるという皮肉な話はひとまず横に置いておきましょう。
W3C 設計原則として認定された Web MCP API に対する Safari チームの対応を見てみたいと思います。おそらく次のように少し書き換えます。
サイトのアクションが AI エージェントにとって使いにくい場合、それはページ自体のセマンティクスにギャップがあるため、まずプラットフォームの共有レイヤー (HTML および ARIA) でギャップを埋めるように努める必要があります。そこでは、ユーザー、支援テクノロジー、エージェントのすべてが利益を得ることができます。
これを W3C の Constituency の優先順位のすぐ隣に置くことができます。
トレードオフが必要な場合は、常にユーザーのニーズを何よりも優先してください…
ユーザーのニーズは、Web ページ作成者のニーズよりも優先され、Web ページ作成者のニーズはユーザー エージェント実装者のニーズよりも優先され、仕様作成者のニーズよりも優先され、理論的な純粋性よりも優先されます。
AI エージェントがそのような優先順位の選挙区に入っていないことに気づくでしょうが、エージェントがそこに留まると仮定すると、それは可能性があります。

どこに置くか検討中です。
ユーザーのニーズはエージェントのニーズよりも優先されるということに全員が同意できることを願っています。 AI をサポートするために Web を変更する提案を検討する際には、このことを念頭に置く必要があります。
ジェイソン・グリグズビー
Cloud Four、Mobile Portland、および Responsive Field Day の共同創設者の 1 人です。彼は、A Book Apart の Progressive Web Apps の著者です。 @grigs で彼をフォローしてください。
私たちは、eコマース、ヘルスケア、ファッション、B2B、SaaS、非営利組織の複雑なレスポンシブ Web デザインと開発の課題を解決します。
通知を受け取る
通知がオンになっています。
通知をオフにする
更新情報をメールで受け取る
更新情報をメールで受け取る
親切、礼儀正しく、建設的な対応をしてください。
単純な HTML を使用するか、
マークダウン
あなたのコメントで。
すべてのフィールドは必須です。
私たちは、eコマース、ヘルスケア、ファッション、B2B、SaaS、非営利組織の複雑なレスポンシブ Web デザインと開発の課題を解決します。
時代の先を行きましょう！ニュースレターを購読して、Web サイトをより良くするための洞察とヒントの宝庫を手に入れましょう。
YouTube で Cloud Four を購読する
@cloudf@mastodon.social をフォローしてください
Bluesky で @cloudfour.com をフォローしてください

## Original Extract

Proposed changes to the web platform to support AI should also support users and assistive technology whenever possible.

Improvements to Web for AI Should Benefit All Users – Cloud Four
Skip to main content
Home: Cloud Four
Main Menu
You appear to be offline, some content may be unavailable.
Dismiss alert
Improvements to Web for AI Should Benefit All Users
Published on June 15th, 2026
June 15th, 2026
Last week, the Safari team formally expressed their opposition to the WebMCP API . In Safari’s explanation of why they oppose WebMCP, they captured something I’ve been trying to articulate.
In the comment explaining the Safari team position, Mike Wyrzykowski writes:
When a site’s actions are hard for an agent to use, that is a gap in the page’s own semantics, and the fix, in our opinion, is to close it in the platform’s shared layers (HTML and ARIA), where the user, assistive technology, and agents all benefit.
In fact, this sentiment should guide any efforts to change the web to support AI agents better. It’s possible that not every enhancement for AI can be structured in a way that supports all of the web’s shared layers, but we should strive to do so first before reaching for solutions that only benefit AI and not the humans the web is supposed to serve.
Our deeper concern is architectural. An agent acting on a user’s behalf is, in effect, assistive technology : it should operate a site as the user would, and the site should not single it out for different treatment.
The proposal says WebMCP “is not designed for ingestion by accessibility technology”, the same fork: richer, actionable semantics reach agents while screen-reader and keyboard users get less.
All of this reminds me of a recent discussion on the Web Incubator Community Group (WICG) mailing list about a proposed standard to help AI understand domain authority. The proposed standard is designed to solve the following problem:
AI assistants often misidentify or misrepresent domains because there is no consistent, machine-readable, domain-controlled source of identity data. Today, models rely on scraped pages, inconsistent metadata, third-party aggregators, or outdated indexes. There is no canonical place where a domain can declare who they are, what they represent, or which resources are authoritative.
The proposed solution suggests adding a JSON-LD document to every website to solve the problem.
Humans struggle with domain authority as well. People have trouble telling what websites are legitimate and which aren’t. If we’re solving this problem for AI, perhaps we can find a solution that works for end users too.
And let’s set aside for the moment the irony that AI is supposed to replace all of our jobs and become a super intelligence and at the same time we also need to add special AI training wheels for it to use the web.
I’d like to see the Safari team’s response to the Web MCP API canonized as a W3C Design Principle. Maybe a slight rewrite like this:
When a site’s actions are hard for an AI agent to use, that is a gap in the page’s own semantics, and we should first seek to close it in the platform’s shared layers (HTML and ARIA), where the user, assistive technology, and agents all benefit.
We can put it right next to the W3C’s Priority of Constituencies :
If a trade-off needs to be made, always put user needs above all…
User needs come before the needs of web page authors, which come before the needs of user agent implementors, which come before the needs of specification writers, which come before theoretical purity.
You’ll notice AI agents aren’t in that priority of constituencies, but assuming agents stick around, it might be worth considering where we’d put them.
I hope we can all agree that user needs come before agent needs. We should keep this in mind as we consider proposals to modify the web to support AI.
Jason Grigsby
is one of the co-founders of Cloud Four, Mobile Portland and Responsive Field Day. He is the author of Progressive Web Apps from A Book Apart. Follow him at @grigs .
We solve complex responsive web design and development challenges for ecommerce, healthcare, fashion, B2B, SaaS, and nonprofit organizations.
Get notifications
Notifications have been turned on.
Turn off notifications
Get Email Updates
Get Email Updates
Please be kind, courteous and constructive.
You may use simple HTML or
Markdown
in your comments.
All fields are required.
We solve complex responsive web design and development challenges for ecommerce, healthcare, fashion, B2B, SaaS, and nonprofit organizations.
Stay ahead of the curve! Subscribe to our newsletter and unlock a treasure trove of insights and tips to make your web site better.
Subscribe to Cloud Four on YouTube
Follow @cloudfour@mastodon.social
Follow @cloudfour.com on Bluesky
