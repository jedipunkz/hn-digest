---
source: "https://www.theregister.com/devops/2026/07/25/how-ai-drove-shopify-back-to-clean-code/5277901"
hn_url: "https://news.ycombinator.com/item?id=49056472"
title: "AI drove Shopify back to clean code"
article_title: "How AI drove Shopify back to clean code"
author: "sbulaev"
captured_at: "2026-07-26T10:29:04Z"
capture_tool: "hn-digest"
hn_id: 49056472
score: 1
comments: 0
posted_at: "2026-07-26T10:07:07Z"
tags:
  - hacker-news
  - translated
---

# AI drove Shopify back to clean code

- HN: [49056472](https://news.ycombinator.com/item?id=49056472)
- Source: [www.theregister.com](https://www.theregister.com/devops/2026/07/25/how-ai-drove-shopify-back-to-clean-code/5277901)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T10:07:07Z

## Translation

タイトル: AI が Shopify をクリーンなコードに戻した
記事のタイトル: AI が Shopify をクリーンなコードに戻す方法
説明: 結局のところ、エージェントは人間と同じものを望んでいます: 読みやすいコード、明示的な契約、役立つフィードバック

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
AI が Shopify をクリーンなコードに戻す方法
結局のところ、エージェントは人間と同じもの、つまり読みやすいコード、明示的な契約、役立つフィードバックを望んでいるだけであることがわかりました。
電子商取引プラットフォームの Shopify は、ソースコードを読むことを再び実現したいと考えています。そして、その原点回帰的なアプローチには、ありそうもない協力者がいます。それは、AI エージェントです。
同社は顧客、別名​​「販売者」向けに新しい店頭テーマを立ち上げており、HTML の知識が少しでもあれば誰でも完全に理解できると主張している。人間とマシンの両方にとって、消化やハッキングがより簡単になるように設計されています。
Shopifyはまだこの新しいテーマに名前を付けていないが、顧客が独自のShopifyページをカスタマイズするために使用する現在のJSONを多用した基本テーマであるHorizo​​nの後継となる可能性のあるテーマとして開発者フォーラムで発表している。
新しいコードの大部分は HTML であり、同社独自のブロックベースの Liquid テンプレート言語によって補間されており、これも非常に目に優しいものです。
販売者がこの新しいテーマのテンプレート ディレクトリを覗くと、不可解な JSON の長い文字列ではなく、解析しやすいコードのプレーンテキスト ファイルが表示されます。このテーマのコード行は Horizo​​n より 93% 少なくなっています。
同社の AI サービスの人気が高まっていることが、この再アーキテクチャの再設計を推進しています。
「主要なモデルはすべて、すでに HTML を理解しています。HTML は表現力豊かで、ローカルで、トークン効率が高いです」と Shopify ストアフロント製品ディレクターの Ben Sehl 氏は X に関するメッセージで述べています。また、Shopify の Liquid テンプレート言語は HTML と完全に連携します。
小売業者からマーチャンダイザーへ
セール氏は投稿の中で、10年前、彼自身が商人だった頃を思い出し、

Shopifyは基本的なテンプレートのみを提供していました。当時初心者の Web プログラマーだった Sehl にとって、サードパーティのスターター キットを使用して変更するのは簡単でした。
時間が経つにつれて、Shopify はテンプレートを強化し、ユーザーが本格的なデザイナーになれるようにしました。オンライン ストア エディターにより、ページとそのさまざまな「ブロック」コンポーネントがどのように見えるかをより詳細に制御できるようになりました。すべてのボタンとノブはカスタマイズできます。
OpenAI では一部の顧客がチャットをエクスポートできませんが、このツールではエクスポートできます。
Codeberg がバイブコーディングされたプロジェクトに主導権を与え、人間の FLOSS を促進
モデル コンテキスト プロトコルはステートフルな過去と決別する準備をしています
AI レポートの洪水でセキュリティ チームが埋もれ、GitHub が公開バグ報奨金の支払いを削減
ただし、柔軟性は複雑さを生みます。ユーザー テンプレートの命令は JSON に変換されました。 「このトレードオフにより、販売業者はより多くの制御を行えるようになりましたが、開発者にとってはエクスペリエンスが犠牲になりました。1 つのファイルを読んだだけではページを理解できなくなりました」と Sehl 氏は書いています。
「テンプレートが JSON になった瞬間、テンプレートは優れた開発者画面ではなくなりました。テンプレートは自動保存される出力になりました」と Sehl 氏は書いています。
ボットにとっても人間にとっても良いこと
最近、Shopify は Shopify Sidekick と呼ばれる AI アシスタントを追加し、ユーザーがページのデザイン方法をさらに細かく制御できるようになりました。 HTML で #0000FF が「青」を意味することを覚えておく必要はもうありません。彼らは背景色を青であると「宣言」するだけで、エージェントがそれを青にします。
今年、Shopify 販売者の 20% がテーマの編集に Sidekick を使用しており、今年だけで 2,500 万件の編集が行われています。
その結果、Shopify は AI エージェント自体の作業を簡素化するという課題に直面しました。
エージェントも人間と同じもの、つまり読みやすいコード、明示的な契約、建設的なフィードバックを望んでいることがわかりました。そこで同社は、テーマ アーキテクチャを再編成しました。

これらの品質を第一原則として維持します (既存の Liquid テーマを「永久 API」で維持しながら)。
これらの変更により、新しい AI オーバーロードのサポートが強化される一方で、基礎となるコードも人間にとってより読みやすくなりました。
「シリアル化された構成のより深い問題は、JSON が本質的に悪いということではありません。構成の語彙が制限されていることです」と Sehl 氏は説明しました。
この新しい基本テーマでは、Liquid に新しい構文とパーサーが追加され、バックエンドでより複雑な命令を実行するための追加の規律を使用して構築されました。ただし、ユーザーに表示されるのは、通常の HTML に簡単に適合する、構成可能な型付きブロックです。
この新しいアプローチでは、ユーザーが設計した各テーマに、追加のアーティファクトを生成する方法に関する指示とポインタを含む特別なディレクトリとファイルが与えられます。さらに、展開された doc タグには、契約書とスニペットの処理に関する例と手順が含まれています。
新しいテーマの鍵となるのは、新しいコンポーザブル ブロック タグです。このタグには、開発者が値、構成、オーバーライド、その他の命令を入力できるページ コンポーネントを記述するネストされたパラメータのセットが含まれています。
React 開発者は、パラメータが React の小道具と同様の仕事をするものとして認識しますが、仮想 DOM を作成する必要はありません。 React ユーザーにはおなじみの新しい部分プリミティブも、ページ全体を再レンダリングすることなくページの特定の領域を更新できるようになります。
チームはまた、ショッピング カートの更新など、サイトのイベント トリガーのコレクションである標準アクションも追加しました。 「この 1 つのプリミティブによって、Horizo​​n から数千行のリアクティビティ コードを削除できるようになります」と Sehl 氏は書いています。
Shopify はまた、生成されたテーマが契約、構造、検証、複雑さに関する会社のポリシーを確実に満たすようにするための 20 の新しいルールを開発しました。

ネストとファイル サイズの制限。
さらに、ユーザーがさらに凝りたい場合は、Liquid でブール式、優先順位のある中置演算子、リテラル配列やオブジェクトなどの論理機能が使用できるようになりました。将来のある時点では、Tailwind CSS フレームワークもサポートされる予定です。
これらすべての改善により、販売者は自宅でカスタマイズされた店頭を組み立てることが簡素化されるはずです。
例: 宣言型の設定主導型アーキテクチャでは、開発者はレイアウトを開始する前に、使用する予定のボタンの数を予測する必要があります。 HTML を使用すると、販売者が追加のボタンが必要であると判断した場合、それらをすべて HTML div タグで囲むだけです。
AI の専門家は、AI にとって最適な食事は、少なくともコーディング以外のタスクにおいては、高密度で記号を多用したコードではないという考えに徐々に落ち着いてきています。人間と同じように、LLM もより理解しやすい散文を食べるとより良い成績を収めるようです (結局のところ、LLM は人間の言語で訓練されています)。
たとえば、ある開発者は、17 種類のツールのドメイン固有言語 (DSL) に比べて、SQL の処理が簡単であることに気づきました。 Anthropic は、プロンプトとして、意味論的に豊富な XML を推奨していますが、10 年前、この形式は業界によってほとんど無視され、より簡潔で冗長性の低い JSON が使用されました。
おそらく、AI をなだめようとする Shopify の取り組みは、さらなる傾斜ではなく、より読みやすいコードを求める動きにつながるかもしれません。 ®
科学
最後のスペースシャトルは15年前に地球に帰還しました
それで、NASA、市販の代替品はどのように機能していますか?
AI が Shopify をクリーンなコードに戻す方法
結局のところ、エージェントは人間と同じもの、つまり読みやすいコード、明示的な契約、役立つフィードバックを望んでいるだけであることがわかりました。
カンボジアのフン・マネ首相はZTEと会談し、デジタルインフラと人工知能（AI）分野での協力を深める
パートナーコンテンツ

: 小見出し 小見出し
反 AI オープンソースには共通の敵がありますが、それ以外にはほとんど何もありません
ボットフリーの代替案を構築するには、左翼、リバタリアン、文化戦士がコードを共有する必要があるかもしれない
コラムニスト
エアバスは AWS から運航します。次に何が起こるかが重要です
また自由の国へ行くにはどっちですか？
BOFH: この印刷技術者は本の裏技をすべて知っています
セキュリティ
Linux カーネル チームが 2 日間で 432 個の CVE を公開
セキュリティ
オラクル、まるで新常態であるかのように 1,449 件のセキュリティ パッチを投下
OSプラットフォーム
開発者が誤って Copilot バイナリを FreeBSD ポート リポジトリにコミットしてしまう
オフプレミス
物置、延長コード、GPU を 2 基、当座貸越を持っている人は誰でもデータセンターを構築しています。富士通は5台を降ろしたばかり
パッチ
1 年にわたるロシアの攻撃により、ユーザーは電子メールを見るとすぐに感染します
結局のところ、エージェントは人間と同じもの、つまり読みやすいコード、明示的な契約、役立つフィードバックを望んでいるだけであることがわかりました。
おまけに、データを保持する必要もありません。
クロード、このモデルを最適化してください
誰がグループレターにサインしなかったのか分かりますか?
敵の敵は味方だ
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
GNOME は Windows のように見えることができ、Flashback は拡張機能なしで実行できます
新しい「シンプルタスクバー」はオプションですが、よりシンプルで安定した方法があります
x86-32 での Debian の最終リリースに黙祷を捧げてください
新しい Debian バージョンが FOSSland に登場

13.6 と 12.15 の形式
脆弱な Joomla Web サイトで拡張機能のバグを悪用し、10 点満点を獲得した悪者を発見
iCagenda、Balbooa Forms 拡張機能の欠陥は、世界中の 100 万のサイトを支えるオープンソース CMS に影響を与える可能性があります
フレーム: 新しい X11 サーバー – アセンブリに直接実装
yserver、Phoenix、そしてもちろん XLibre、そして外れ値の Arcan に参加します
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
Linux Mint デスクトップの次期バージョンには両方の種類のディスプレイ サーバーが搭載されています
KDE Plasma ユーザーは恐ろしい変化の予兆に直面しています: 6.6.6 の登場
現在は 6.7 が最新版であり、6.8 では好むと好まざるにかかわらず Wayland を入手できるようになります。
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
次のプラットフォーム
開発クラス
ブロックとファイル
状況出版
クッキーポリシー
プライバシーポリシー
利用規約
私の個人情報を共有しないでください
同意のオプション
著作権。すべての著作権は © 1998-2026 に留保されます。

## Original Extract

Turns out, agents just want the same things as humans: easily-readable code, explicit contracts, and helpful feedback

Jump to main content
Search
TOPICS
Special Features
All Special Features
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
How AI drove Shopify back to clean code
Turns out, agents just want the same things as humans: easily-readable code, explicit contracts, and helpful feedback
E-commerce platform Shopify wants to make reading source code a thing again. And it has an unlikely ally for its back-to-the-roots approach: AI agents.
The company is launching a new storefront theme for its customers, aka “merchants,” that it claims is completely understandable by anyone with even a smidgen of HTML knowledge. It is designed to be simpler for both people and machines to digest and hack against.
Shopify hasn’t named this new theme yet, but is unveiling it in developer forums as the probable successor to Horizon, its current JSON-heavy base theme, which customers use to customize their own Shopify pages.
The new code is mostly HTML, interpolated by the company’s own block-based Liquid templating language, which is also fairly easy on the eyes.
When merchants peek into the templates directory of this new theme, they will see plain-text files of easy-to-parse code, not long strings of inscrutable JSON. This theme has 93% fewer lines of code than Horizon.
The company’s increasingly popular AI service is driving this rearchitecture redesign.
“Every leading model already understands HTML. It is expressive, local, and token-efficient,” wrote Ben Sehl , Shopify product director for storefronts, in a missive on X. And Shopify’s Liquid template language meshes perfectly with HTML.
From merchants to merchandisers
In his post, Sehl recalled a decade ago, when he was a merchant himself, and Shopify only offered a basic template. It was easy for Sehl, then a novice web coder, to modify with a third-party starter kit.
Over time, Shopify enriched the templating, allowing users to become full-fledged designers. The Online Store Editor gave them more control over how a page and its various “block” components would look like. Every button and knob could be customized.
OpenAI won't let some customers export their chats, but this tool will
Codeberg gives vibe-coded projects the toss, promotes human FLOSS
Model Context Protocol prepares to break with its stateful past
GitHub slashes public bug bounty payouts as AI report flood buries its security team
Flexibility begets complexity, however. The user template instructions were converted to JSON. “That trade-off gave merchants far more control, but it came with a developer-experience cost: you could no longer understand a page by reading one file,” Sehl wrote.
“The moment templates became JSON, they stopped being a great developer surface. They became an auto-saved output,” Sehl wrote.
Good for the bot, good for the human
Recently, Shopify added an AI assistant called Shopify Sidekick, which gave users even more control over how they designed their pages. No longer would they have to remember that #0000FF means ‘Blue’ in HTML speak. They just “declare” the background color to be blue and the agent will make it so.
This year, 20% of Shopify merchants are using Sidekick to edit their themes, making 25 million edits this year alone.
Consequently, Shopify faced the task of simplifying things for the AI agent itself.
It turns out that agents want the same things as humans: code that is easy to read, explicit contracts, and constructive feedback. So the company has reorganized its theme architecture using these qualities as first principles (while maintaining existing Liquid themes in a “forever API”).
While these changes better support our new AI overlords, the underlying code also became more readable to humans as well.
“The deeper problem with serialized configuration isn’t that JSON is inherently bad. It’s that configuration has a constrained vocabulary,” Sehl explained.
For this new base theme, Liquid gained a new syntax and parser, built with some additional discipline to execute more complex instructions on the back-end. What the user sees, however, are composable, typed blocks that fit easily alongside ordinary HTML.
In this new approach, each user-designed theme is given a special directory and files with instructions and pointers of how additional artifacts should be generated. Furthermore, an expanded doc tag contains examples and instructions on handling contracts and snippets.
Key to the new theme is a new composable block tag, which contains a set of nested parameters describing page components that the developer can feed with values, configurations, overrides, and other instructions.
React devs will recognize parameters as doing a similar job to React’s props, but without the need to create a virtual DOM. Also familiar to React users will be the new partial primitive, which allows a specific region of a page to be updated without needing to re-render the entire page.
The team also added standard actions, a collection of event triggers for the site, such as, say, updating the shopping cart. “This one primitive will let us remove thousands of lines of reactivity code from Horizon,” Sehl wrote.
Shopify also developed 20 new rules to ensure generated themes meet the company’s policies around contracts, structure, validation, complexity, nesting, and file-size limits.
And if the user wants to get fancy, Liquid now allows for logical capabilities such as Boolean expressions, infix operators with precedence, and literal arrays and objects. At some point in the future, the Tailwind CSS framework will even be supported.
All these improvements should simplify things for the merchant at home, trying to assemble a customized storefront.
For instance: In a declarative, settings-driven architecture, the developer must anticipate how many buttons they plan to use before they start the layout. With HTML, when the merchant discovers they need additional buttons, they just wrap them all in an HTML div tag.
AI experts are slowly coming around to the idea that the best diet for AIs is not dense, symbolic-heavy code — at least for non-coding tasks. Like humans, LLMs seem to do better on a diet of more easily-understandable prose (LLMs were, after all, trained on human language).
One developer, for instance, found that SQL was easier to process compared to the Domain Specific Languages (DSLs) of 17 different tools. For prompting, Anthropic recommends the semantically-rich XML, whereas a decade ago that format was largely cast aside by the industry for the leaner, less-verbose JSON.
Perhaps Shopify’s efforts to appease AI may lead not to more slop but to a movement for more readable code. ®
SCIENCE
The last Space Shuttle returned to Earth 15 years ago
So, how are those commercial replacements working out for you, NASA?
How AI drove Shopify back to clean code
Turns out, agents just want the same things as humans: easily-readable code, explicit contracts, and helpful feedback
Cambodian Prime Minister Hun Manet met with ZTE to deepen cooperation in digital infrastructure and artificial intelligence (AI)
PARTNER CONTENT: subhead subhead
Anti-AI open source has an enemy in common, but almost nothing else
Building bot-free alternatives may require lefties, libertarians, and culture warriors to share code
columnists
Airbus takes flight from AWS. What happens next is critical
Which way to the Land of the Free again?
BOFH: This printer engineer knows every trick in the book
security
Linux kernel team publishes 432 CVEs in two days
Security
Oracle drops 1,449 security patches like it's the new normal
OS PLATFORMS
Dev accidentally commits Copilot binary to FreeBSD ports repo
off-prem
Anyone with a shed, an extension cord, a couple of GPUs and an overdraft is building datacenters. Fujitsu just offloaded five
PATCHES
Year-long Russian attacks infect users as soon as they look at an email
Turns out, agents just want the same things as humans: easily-readable code, explicit contracts, and helpful feedback
And as a bonus, it doesn't require data retention
Hey Claude, optimize this model for me
Can you guess who didn't sign on to the group letter?
The enemy of my enemy is my friend
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
GNOME can look like Windows – and Flashback can do it without extensions
New 'Simple-taskbar' is an option, but there's a simpler, stabler way
A moment of silence, please, for the final release of Debian on x86-32
New Debian versions hit FOSSland in the form of 13.6 and 12.15
Baddies caught exploiting extensions bugs with perfect 10 scores on vulnerable Joomla websites
Flaws in iCagenda, Balbooa Forms extensions can impact open source CMS that powers a million sites worldwide
Frame: A new X11 server – implemented directly in assembly
Joins yserver, Phoenix, and of course XLibre – and outlier Arcan
Cinnamon 6.8 will support Wayland – if you want it
Next version of Linux Mint’s desktop has both kinds of display server
KDE Plasma users face a dire omen of change: 6.6.6 arrives
6.7 is now current, and in 6.8 you're getting Wayland whether you like it or not
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.
