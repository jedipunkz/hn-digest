---
source: "https://axesslab.com/lovable/"
hn_url: "https://news.ycombinator.com/item?id=48490304"
title: "Lovable's AI built a 100% accessible site – or did it?"
article_title: "Axess Lab | Lovable’s AI built a 100% accessible site – or did it?"
author: "robin_reala"
captured_at: "2026-06-11T16:29:45Z"
capture_tool: "hn-digest"
hn_id: 48490304
score: 2
comments: 0
posted_at: "2026-06-11T13:48:59Z"
tags:
  - hacker-news
  - translated
---

# Lovable's AI built a 100% accessible site – or did it?

- HN: [48490304](https://news.ycombinator.com/item?id=48490304)
- Source: [axesslab.com](https://axesslab.com/lovable/)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T13:48:59Z

## Translation

タイトル: Lovable の AI は 100% アクセス可能なサイトを構築しました – それとも実際にそうでしたか?
記事タイトル: Axess Lab | Lovable の AI は 100% アクセス可能なサイトを構築しました – それとも実際にそうなったのでしょうか?
説明: AI によって構築されたサイトが現時点でどの程度アクセスしやすいかを知りたいと考えていました。そこで、同僚のダニエルにサイトを試してもらいました…

記事本文:
外観
暗い
ライト
システム
Lovable の AI は 100% アクセス可能なサイトを構築しました – それとも実際にそうなったのでしょうか?
2026 年 5 月 13 日に公開
ハンプス・セスフォース著
私たちは、AI によって構築されたサイトが現時点でどの程度アクセスしやすいかを知りたかったのです。そこで、同僚のダニエルに、iPhone のスクリーン リーダーを使用して、Lovable で構築されたサイトを試してもらいました。すべて録画しましたので、ぜひご参加ください。
まず第一に、サイトの構築方法は私にとって非常に衝撃的でした。
AI 愛好家である私の友人は、通勤のバスの中でスマートフォンだけを使ってそれを構築しました。通勤中に機能的で忠実度の高いサイトを構築できるという事実は、非常に素晴らしいことです。そして、これはアクセシビリティの一形態であり、より多くの人がサイトやアプリを構築できるようにするものだと主張することもできます。
しかし、はい、簡単に言えば、このサイトは開発カンファレンスのためのものでした。そこには、講演者、日付、会場などに関する情報が含まれていました。チケットを確保し、料金を支払う方法も含まれています。つまり、非常に複雑なサイトではありませんが、単なる基本的な一方通行の情報サイトでもありません。
Lovable で構築されており、チェックアウト フローなどでいくつかのサードパーティ統合が行われています。アクセシビリティを優先するよう促されていましたが、それに関する詳細な要件は示されていませんでした。
書類上、このサイトは完璧でした。 Axe を実行する Lovable の Speed ツールでは、100% のアクセシビリティ スコアを獲得しました。ところで、Lovable にこのようなアクセシビリティ ツールが組み込まれているのは素晴らしいことです。
しかし、実際の支援技術ユーザーにとって、実生活ではどのように機能したのでしょうか?ヒントをあげておきますが、100%ではありませんでした…
コンポーネントの背後でスタックする
いくつかの問題を見てみましょう。最も重要な問題から始めます。
スクリーン リーダー ユーザーにとって最もイライラする経験の 1 つは、ビジュアル レイヤーとコード レイヤーが

同期が失われます。メニューが開いているとき、スクリーン リーダーはその下のコンテンツを読み続けました。
可能であればビデオをチェックしてください。ただし、何らかの理由で見たくない場合は、クリップの後に概要を説明します (後で他のビデオも)。
基本的に、ビデオではメニューが開いていることが示されていますが、スクリーン リーダーは多数のメニューの背後にあるオブジェクトに焦点を当てています。これは明らかに混乱を引き起こします:
メニューにいると思ったのですが、スタートページから前に読んだものがアナウンスされていました…おそらくメニューの下にありますよね？それで、それは奇妙でした！
この「ゴースト」効果はチケットモーダルでも再び発生しました。
これは Youtube のクリップです。何らかの理由で記事に埋め込むことができません。
AI サイト (Youtube) で VoiceOver を使用してチケットを購入する
ここでダニエルは最初にスワイプ ナビゲーションを使用し、右にスワイプすると次の項目に移動します。これは、新しいインターフェイスを使用しているとき、または技術にあまり詳しくないユーザーが操作する一般的な方法です。スワイプ ナビゲーションを使用すると、モーダルの背後で立ち往生してしまいます。
ただし、タッチによるナビゲーションに切り替えると、モーダルに強制的にフォーカスを移すことができます。基本的には、画面上で指をドラッグし、指の下にあるものを読み上げます。
ナビゲーション方法を切り替える必要があることに誰もが気づくわけではありません。ダニエルが最後に述べたように、キーボードや点字ディスプレイを接続していて、タッチによるナビゲーションをまったく使用しないユーザーもいるでしょう。
自動化されたセキュリティ機能と奇妙な隠し領域も、少々混乱を引き起こしました。新しいページが開くたびに、これが経験されました。
1 つは、ページの上部に「通知 alt+t」という奇妙な非表示領域があることです。
しかし、最もイライラしたのは、スクリーン リーダーが自動的に「Cloudflare の整合性」と「あなたが人間であるかどうかを確認してください」とアナウンスし始めたことです。

繰り返し。
サイトの概要を把握しようとしているユーザーにとって、セキュリティ ボットがフローを中断することは、図書館でメガホンが鳴らされるのとデジタル的には同等です。
これは、ガイドラインに従っていても実際のエクスペリエンスをテストしない場合に発生する可能性がある問題の好例です。 Cloudflareコンポーネントはそのコンテンツを更新し、それをライブリージョンに置くのが一般的な経験則です。ただし、ビデオでおそらく指摘したように、この特定のケースでは、スクリーン リーダーのエクスペリエンスに大きな悪影響を及ぼします。
シングルページ アプリケーション (SPA) で失われる
サイトは SPA として構築されているため、新しいページに「移動」しても、実際には従来のページの読み込みはトリガーされませんでした。その場合、新しいページが読み込まれるときにフォーカスが適切な場所に配置されるように、フォーカスを制御する必要があります。ただし、このサイトでは、スクリーン リーダーのフォーカスが適切に処理されませんでした。
フォーカスは決して管理されませんでした…私のスクリーン リーダーはそこで何が起こったのかを理解しようとしたため、リンクを押した視覚領域にフォーカスが置かれました。
結果？ダニエルはページの途中にある見出しに着地し、その上のコンテンツをすべて見逃していました。
「メニューの切り替え」で何も表示されない場合
ダニエルが最初に反応したのはメニューボタンでした。ボタンには「メニューの切り替え」というラベルが付いていたため、ラベルが付けられていませんでした (よくある問題)。それでいいのです！
しかし、メニュー ボタンにアクセスできるようにするより良い方法があります。それは、aria-expanded 属性を使用して、メニューが展開されているか折りたたまれているかを示すことです。ダニエルの反応を見てみましょう！
まだメニューの切り替えと表示されます…何かを展開したかどうかは通知されないため、機能するかどうかはわかりません。
したがって、これは私たちが見つけた中で最もアクセスしにくいメニュー ボタンではありませんでしたが、最高のメニュー ボタンでもありませんでした。
さらに、サイトを横向きモードまたは小型デバイスで使用している場合、次のページにスクロールすることはできません。

下部メニューのオプション:
スクリーン リーダーのユーザーは見出しを使用してセクション間を移動し、情報の階層を理解できます。ただし、AI によって構築されたサイトでは、見出しレベルがスキップされることがありました。
そのため、サイトはレベル 1 からレベル 3 に直接スキップされました。スクリーン リーダーのユーザーにとって、これは本のページが切り取られたように感じることがあります。重要な情報をスキップしたかどうかがわからないだけです。
自動ツールはスキップされた見出しレベルを簡単に見つけてテストできるため、これはおそらくこのサイトで最も驚くべき失敗でした。
次のビデオはありませんが、「チケットを取得」ボタンがあります。
ビジュアルテキストは「チケットを取得」ですが、aria-label=”チケットに登録”が含まれています。これは、アクセシビリティ対応の名前とビジュアル名が一致しないことを意味します。
なぜこれが問題になるのでしょうか?その主な理由は、一部の運動障害のあるユーザーが音声コントロールを使用してサイトを操作するためです。ユーザーは、「クリックしてチケットを入手」などのボタンの表示名を言いますが、アクセシブルな名前にそのフレーズが含まれていない場合は、何も起こりません。
これは WCAG の明確な要件です。2.5.3 ラベルという名前であり、テストは非常に簡単です。さあ、AI さん、将来はこういうこともキャッチしてくれるはずですよ！
最後に、言語スイッチャーにアクセスできなくなりました。
これは Youtube のクリップです。何らかの理由で記事に埋め込むことができません。
スクリーン リーダーのユーザーは、AI で構築されたサイト (Youtube) でどの言語が選択されているかを認識できません
「enボタン」と「svボタン」と書かれているだけで、どれがどれなのかわかりません。
さらに、サイトには言語を切り替えるコンテンツの適切な lang 属性が含まれていなかったため、スクリーン リーダーはコンテンツの正しい音声に切り替えることができませんでした。
繰り返しますが、これは簡単にテストできるものであり、それに関するガイドラインは非常に優れています

明らかだったので、もっと期待していました。
ダニエルは自分の考えを次のように要約しています。
彼の引用文の要旨は次のとおりです。
要するに、内容を読んでも大丈夫だと感じました。しかし、インタラクティブなコンテンツが登場するとすぐに、いくつかの問題が発生しました。いくつかは他のものよりも深刻です。したがって、親切にするためにはまだ改善の余地があります。
したがって、サイトが 100% アクセス可能ではなかったという点には同意できると思います。
しかし、私は前向きな気持ちで終わりたいと思っています。
おそらく、この実験で最も刺激的な部分は、見つかったエラーではなく、それらがいかに早く消えるかということでした。
講演中にこのフィードバックを共有した後、10 分間の質疑応答がありました。それが起こっている間、サイト作成者はソファに座り、スマートフォンを使用してこれらの問題の多くを修正しました。これには、フォーカス管理とメニューとモーダルの漏洩に関するいくつかのより困難な問題が含まれます。とてもクールです！
したがって、人間の専門家が関与していれば、AI ツールがアクセス可能なサイトを作成できる可能性が高くなります。
ただし、これらのツールを使用して作成されたサイトの少なくとも 99.9% には、アクセシビリティの専門家が関与しておらず、サイト作成者がアクセシビリティの重要性を促しているわけでもありません。したがって、Lovable や同様のツールが、構築したすぐに使用できるインターフェイスをデフォルトでアクセスできるようにするために熱心に取り組んでくれることを期待しています。それは素晴らしい成果ですが、現時点ではその目標にはほど遠いようです。
新しいものを書いたときに通知を受け取る
私たちは月に 1 回程度、アクセシビリティや使いやすさに関する記事を書きますが、これも同様に素晴らしいものです。 #謙虚な自慢
私たちは世界中で働いています。サポートが必要な場合は、メールでお問い合わせください。
hello@axesslab.com
アクセスラボ
バサガタン 28
111 20 ストックホルム、スウェーデン

## Original Extract

We wanted to get an indication of how accessible AI-built sites are at the moment. So I let my colleague Daniel try out a site…

Appearance
Dark
Light
System
Lovable’s AI built a 100% accessible site – or did it?
Published 13 May 2026,
by Hampus Sethfors
We wanted to get an indication of how accessible AI-built sites are at the moment. So I let my colleague Daniel try out a site built with Lovable, using the screen reader on his iPhone. And we recorded it all so you can come along for the ride!
First off: the way the site was built was pretty mind blowing to me.
A friend of mine, an AI-enthusiast, built it only using his smartphone during bus rides to and from work. The fact that you can now build a functional, high-fidelity site into existence while commuting is pretty cool. And you could argue that this is a form of accessibility, making it possible for many more people to build sites and apps.
But yeah, in brief, the site was for a dev conference. It included information about speakers, dates, venue and that sort of stuff. It also included a way secure your tickets and pay for them. So not a super complex site, but still not just a basic one-way info site.
It was built with Lovable and has some third party integrations, for instance in the checkout flow. Accessibility had been prompted to be a priority, but no more detailed requirements regarding that was given.
On paper the site was perfect. It got a 100% accessibility score in Lovable’s Speed tool, that runs Axe. By the way: great to see that Lovable has an accessibility tool like this built in!
But how did it work in real life, for an actual assistive tech user? I’ll give a hint: it wasn’t 100%…
Getting stuck behind components
Let’s look at some of the issues, and I’ll begin with the most critical ones.
One of the most frustrating experiences for a screen reader user is when the visual layer and the code layer lose sync. When the menu was open, the screen reader kept reading content underneath it.
Check out the video if you can, but I’ll summarise it (and others videos later on) after the clip if you for some reason can’t want to watch it.
Basically, the video shows the menu being open, but the screen reader focusing on some object behind the meny. Which obviously causes confusion:
I thought I was in the menu, but it announced the thing I read before from the start page… probably underneath the menu, right? So that was weird!
This “ghosting” effect happened again with the ticket modal.
Here’s the clip on Youtube, for some reason I can’t embed it here in the article:
Buying ticket using VoiceOver on AI-built site (Youtube)
Here Daniel initally uses swipe navigation, where swiping right moves to the next item. It’s a common way of navigating when you’re in new interfaces, or for less tech savvy users. Using swipe navigation, he gets stuck behind the modal.
However, he manages to force focus into the modal when he switches to navigating by touch. Basically dragging his finger across the screen and having what’s underneath his finger read to him.
Far from everyone will figure out that they need to switch ways of navigating, and like Daniel mentions at the end, some users will have keyboards or braille displays connected and not use navigation by touch at all.
Automated security features together with a strange hidden region also caused a bit of chaos. This was the experience every time a new page opened:
One was that there was a strange, hidden region at the top of the page saying “Notification alt+t”.
However, the most frustrating one was that the screen reader began automatically announcing “Cloudflare integrity” and “Verify if you are a human” repeatedly.
For a user trying to get an overview of a site, having a security bot interrupt your flow is the digital equivalent of a megaphone going off in a library.
This is a great example of the issues you can run into if you follow guidelines, but don’t test for the actual experience. The Cloudflare component does update its content, and then having it in a live-region is the general rule of thumb. However, in this specific case it hurts the screen reader experience tremendously, as you probably noted in the video.
Lost in the single-page application (SPA)
Because the site was built as an SPA, “navigating” to a new page didn’t actually trigger a traditional page load. When that’s the case, focus needs to be controlled so it lands in the proper place when new pages load. However, on this site, screen reader focus wasn’t handled well:
Focus was never managed… my screen reader tried to figure out what happened there so it put focus in that visual area where I pressed the link.
The result? Daniel landed on a heading halfway down the page, missing all the content above it.
When “Toggle Menu” tells you nothing
The first thing Daniel reacted to was the menu button. The button was labelled “Toggle menu”, so it wasn’t unlabelled (a common issue). So that’s good!
But there’s a better way to make menu-buttons accessible: using the aria-expanded attribute to indicate if the menu is expanded or collapsed. Let’s see Daniel’s reaction!
It still says toggle menu… I’m not sure if it works because it doesn’t announce if I have expanded something.
So it wasn’t the least accessible menu button we’ve come across, but not the best either.
On top of this, if you use the site in landscape mode or on small devices, it’s not possible to scroll to the bottom menu options:
Headings allow screen reader users to jump between sections and understand the hierarchy of information. The AI-built site, however, sometimes skipped heading levels:
So the site skipped from Level 1 directly to Level 3. For a screen reader user this can feels like pages have been ripped out of a book—you simply don’t know if you’ve skipped vital information.
This was probably the most surprising failure on the site, since automatic tools easily can find and test for skipped heading levels.
I don’t have a video for this next one, but there’s a “Get tickets” button.
The visual text is ‘Get tickets’, but it has an aria-label=”Register for tickets”. This means that the accessible name and the visual name don’t match.
Why is this an issue? Well, mainly because some motor impaired users will use voice control to navigate a site. They will say the visible name of the button, like “Click get tickets,” and if the accessible name doesn’t include that phrase, nothing happens.
This is a clear requirement in WCAG: 2.5.3 Label in name and quite straight forward to test for. So come on AI, you should catch these sorts of things in the future!
Finally, the language switcher, was inaccessible.
Here’s the clip on Youtube, for some reason I can’t embed it here in the article:
Screen reader user can’t tell which language is selected on AI built site (Youtube)
It just says ‘en button’ and ‘sv button’… I can’t tell which one is which really.
Additionally, the site didn’t include the proper lang attributes for the content that switched language, meaning the screen reader wouldn’t know to switch to the correct voice for the content.
Again, this is something that’s easily testable and the guidelines around it are super clear, so I was expecting more.
Daniel summarises his thoughts:
Here’s the gist of his quote in text:
In short, it felt okay to read content. But as soon as there was interactive content…there were some problems. Some more severe than others. So there’s still room for improvement to be kind.
So I think we can agree that the site wasn’t 100% accessible.
I do, however, want to end on a positive note.
Maybe the most exciting part of this experiment wasn’t the errors we found, but how fast they disappeared.
After I shared this feedback during a talk, there was a 10 minute q&a. While that was going on, the site creator sat in a sofa and fixed many of these issues using his smartphone. Including some of the more difficult issues regarding focus management and the leaking menus and modals. Very cool!
So with a human expert in the loop, AI tools are more likely to be able to create accessible sites.
However, at least 99.9% of the sites created with these tools will not have an accessibility specialist involved, nor will the site creator have prompted that accessibility is important. So I’m hoping that Lovable and similar tools work hard to make the out-of-box interfaces they build accessible by default. That would be an awesome achievement, however, we seem to be far from that place at the moment.
Get notified when we write new stuff
About once a month we write an article about accessibility or usability, that's just as awesome as this one! #HumbleBrag
We work world wide, if you need help send us an email!
hello@axesslab.com
Axess Lab
Vasagatan 28
111 20 Stockholm, Sweden
