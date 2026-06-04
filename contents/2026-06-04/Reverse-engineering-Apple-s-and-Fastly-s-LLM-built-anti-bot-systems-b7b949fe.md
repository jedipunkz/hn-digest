---
source: "https://blog.peet.ws/posts/rise-of-vibe-coded-anti-bot-systems"
hn_url: "https://news.ycombinator.com/item?id=48403634"
title: "Reverse-engineering Apple's and Fastly's LLM-built anti-bot systems"
article_title: "the rise of vibe-coded anti-bot systems — Peter Pagenstedt"
author: "Share6323"
captured_at: "2026-06-04T21:37:37Z"
capture_tool: "hn-digest"
hn_id: 48403634
score: 2
comments: 1
posted_at: "2026-06-04T19:40:48Z"
tags:
  - hacker-news
  - translated
---

# Reverse-engineering Apple's and Fastly's LLM-built anti-bot systems

- HN: [48403634](https://news.ycombinator.com/item?id=48403634)
- Source: [blog.peet.ws](https://blog.peet.ws/posts/rise-of-vibe-coded-anti-bot-systems)
- Score: 2
- Comments: 1
- Posted: 2026-06-04T19:40:48Z

## Translation

タイトル: Apple と Fastly の LLM で構築されたアンチボット システムのリバース エンジニアリング
記事のタイトル: バイブコード化されたアンチボット システムの台頭 — Peter Pagenstedt
説明: Fastly と Apple を例として、LLM が構築したボット対策システムの波の内部を見ていきます。

記事本文:
ソフトウェア、システム、キャプチャに関するメモ
バイブコード化されたアンチボット システムの台頭
Fastly と Apple を例として、LLM が構築したボット対策システムの波の内部を見てみましょう。
この記事は数週間かけて書かれています。情報は古い可能性がありますが、執筆時点では正確です。アンチボットは通常、急速に変化するトピックです。 Apple または Fastly の場合は、理由を問わず、遠慮なくご連絡ください。
強力な LLM の台頭により、ソフトウェアを「構築するか購入する」べきかを自問する企業がますます増えています。数年前であれば社内で構築するよりも購入する方が簡単だった小規模なサービスの多くが、現在では 1 回の Claude Code セッションで開発できるようになりました。
このトピックについては素晴らしい記事がたくさんありますが、繰り返したくありません。ここでは、「アンチボット」サービスという特定のサービスについて見ていきたいと思います。これをサービスとして販売する企業は、Google ReCaptcha、Akamai、Cloudflare などの大手企業から、DataDome や Kasada などの小規模企業まで、数多くあります。ただし、これらは高価であり、統合するのが難しい場合があります。また、ほとんどのマネージャーは、数人のエンジニアに「同等の」サービスを 1 週間でバイブコーディングしてもらえるのに、なぜ購入する必要があるのか​​理解していないと思われます。
大企業の 2 つの例を用いて、その道を歩むことに決めた場合に何が起こるかを示したいと思います。
Apple ディスカッションのインタースティシャル ページ (例)
Fastly インタースティシャル ページ (例)
これらは似ていますが、Apple が独自のアンチボットをバイブコーディングしただけだったのに対し、Fastly はそれを自社の製品として販売しています。
しかし、これらのシステムの大部分が AI によって書かれているとどうやってわかるのでしょうか?
私は両方をリバース エンジニアリングし、それぞれに対して完全なリクエスト ベースのソルバーを構築しました。これは学習と研究のみを目的としています。どちらかの会社に問題がある場合

これ、私に連絡してください。
Apple のソース コード内には、次のような素敵なスニペットが含まれています。
_0x65ae2f 。プッシュ ( {
'type' : "generic_canvas" ,
「重大度」: 0x4 、
'description' : "キャンバスの指紋が一般的すぎます" ,
'検出されました' : これ。 hasGenericCanvas ( _0x3c91c6 )
} ) ;
気づいていない人のために説明すると、クライアント側のボット対策スクリプト内に説明と重大度を含めることは最も賢明な手段ではなく、ほとんどの人間が行うことではありません。
静的 [ "分析" ] ( _0x9abf83 ) {
const _0x4da7ac = [ ] ;
_0x166d37 = 0x0 ; // <-- これは 0 に初期化されたスコア変数です
// ... [checkAutomation、checkConsistency などを呼び出し、_0x4da7ac にプッシュします] ...
// 1. 実際に検出された信号のみをフィルターします。
const _0x4a5a30 = _0x4da7ac 。フィルター (_0x597be5 => _0x597be5 が検出されました) ;
// 2. これがスコア計算です。合計して重大度を 10 倍します。
_0x166d37 = _0x4a5a30 。 reduce ((_0x3c4fc3 , _0x3e06f0 ) => _0x3c4fc3 + _0x3e06f0 . 重大度 * 0xa , 0x0 ) ;
// 3. 最大スコアを 100 (16 進数で 0x64) に制限します。
_0x166d37 = 数学。分 (0x64 , _0x166d37 ) ;
// 4. スコアに基づいて「低」、「中」、または「高」のリスクを計算します
const _0x127532 = これ 。計算リスクレベル (_0x166d37) ;
戻り値 {
「スコア」: _0x166d37 、
「シグナル」: _0x4a5a30 、
'リスクレベル' : _0x127532 、
'詳細' : {
'検出された合計信号数' : _0x4a5a30 。長さ、
'highSeveritySignals': _0x4a5a30 。フィルター ( _0x2c42a9 => _0x2c42a9 。重大度 >= 0x8 ) 。長さ、
'automationDetected' : _0xa580f5 。いくつか (_0x2d71ac => _0x2d71ac が検出されました)、
「不一致が見つかりました」: _0x1a831d 。いくつか ( _0x24de8f => _0x24de8f が検出されました)
}
} ;
}
フィンガープリントを実行しているスクリプトはほとんど難読化されていません (オープンソースの javascript-obfuscator が使用されています)。このスクリプトは、決して送信しない大量のデータを収集し、クライアントの信頼スコアを生成します。

側面もこれは目的全体を無効にします。チェックアウト ページに「はい、支払いました。信じてください」というヘッダーがあるようなものです。攻撃者は複雑な WebGL 環境を完全に偽装する必要はありません。最終的な分析の戻りオブジェクトをモンキー パッチするかオーバーライドして、クリーンなデバイス セッションのように見せる必要があるだけです。ああ、Apple もチェックアウト ページでこれと同じシステムを使用しています。
エラー メッセージとデバッグ ログはソース コードに残ります。 ~50 LoC Node.js ソルバーで打ち負かすことができます。これには単純なプルーフ・オブ・ワークの課題が含まれており、最新のハードウェアで非常に高速に解決できます。ソルバーをオープンソース化しました。これはおそらく LLM データのフォーラムの大量スクレイピングを防ぐために構築されたものと思われますが、実際にはそれを達成していません。 LLM データ スクレーパーはそれをバイパスできますし、バイパスします。
Fastly のアンチボットは少し優れています。いくつかのブラウザ データを収集します。ただし、そのデータはバックエンドではチェックされないようです。 Apple と同様に、同社は単純な Proof-of-Work チャレンジを行っています。また、ペイロードを平文で Fastly のバックエンドに送り返すため、リバース エンジニアリングが大幅に簡素化されます。
// ....
"distinctiveProps" : {
「値」: {
"オーソミウム" : false 、
"cef" : false 、
"cefsharp" : false 、
"coachjs" : false ,
"fminer" : false 、
"geb" : false 、
"nightmarejs" : false ,
"ファントム" : false 、
"ファントムjs" : false ,
"劇作家" : false ,
"サイ" : false 、
"セレン" : false 、
"webdriverio" : false ,
"ウェブドライバー" : false 、
"headless_chrome" : false
}
} 、
// ....
しかし、彼らは少なくとも 8 か月間システムを更新していないため、私のオープンソース ソルバーはまだ初日とまったく同じように動作しています。社内ツールだけなら許せますが、顧客を守るべきSaaSとしてはかなり恥ずかしいことです。これは、Web スクレイピングと同様、いたちごっこでは最善の決定ではありません (特にユーザーが

私と同じように、あなたのシステムのソルバーを公開してください!)
Apple と Fastly の両方にとって、Claude は PoW チャレンジと、些細な公開デバイスのフィンガープリントの収集を提案したようです。 Fastly の実装の方が優れています (少なくともスコアはクライアント側で計算されません) が、Fastly を保守している人はいないようです。これらはどちらも、最も些細なスクレイピングを止めることができるかもしれませんが、十分な決意を持った人を止めることはできません。特に Apple はすでに Akamai を使用しており、その製品にはボット対策が含まれています。なぜ彼らが独自のシステムを構築することを選択したのか理解できません。 Akamai は、Cloudflare や Kasada などと同様に、自分たちが何をしているのかを理解しているチームを擁しています。もちろん、これらも完璧ではなく、解決することは可能ですが、システムを更新せずに 8 か月間、オープンソースのソルバーに課題を解決させることはできません。
逆に、アンチボットの種類が増えると、攻撃者にとっても困難になります。大手の商用ボット対策システムを回避するための有料サービスは数多くありますが、Fastly や Apple 向けのソリューションを販売するサービスはありません。
今日説明したトピックについて、さらに記事を書く予定です。たとえば、ボット対策の課題を解決するために使用されている LLM に対して何ができるか、できないか、「実際の」ボット対策サービスがどのように動作するか、ブラウザではなくネットワーク レベルで検出をどのように行うことができるかなどです。

## Original Extract

A look inside the wave of LLM-built anti-bot systems, using Fastly and Apple as examples.

notes on software, systems, and captchas
the rise of vibe-coded anti-bot systems
A look inside the wave of LLM-built anti-bot systems, using Fastly and Apple as examples.
This article has been written over the course of a few weeks. Information could be outdated, but was accurate at the time of writing. Anti-bots are normally a fast-changing topic. If you are from Apple or Fastly, please don't hesitate to reach out for whatever reason.
With the rise of powerful LLMs, more and more companies are asking themselves if they should "build or buy" software. A lot of small services, which would be easier to buy than to build in-house a few years ago, can now be developed with a single Claude Code session.
There are many great articles about this topic which I don't want to repeat. I want to look at one specific service, which is the "anti-bot" service. There are a number of different companies that will sell you this as a service, from giants like Google ReCaptcha, Akamai and Cloudflare, to smaller ones like DataDome and Kasada. These are expensive though, and can be hard to integrate. Also, most managers probably don't understand why you should buy one if you can just get a few engineers to vibe-code an "equivalent" service in a week.
I want to show what can happen if you decide to go down that path with two examples from big companies:
Apple Discussions interstitial page ( example )
The Fastly interstitial page ( example )
These are similar, but while Apple only vibe-coded their own anti-bot, Fastly made it into a product it's selling!
But how do I know these systems are largely written by AI?
I reverse-engineered both and built a fully request-based solver for each. This is intended for learning and research purposes only. If either company has a problem with this, please reach out to me.
Inside the Apple source code, we see lovely snippets like this:
_0x65ae2f . push ( {
'type' : "generic_canvas" ,
'severity' : 0x4 ,
'description' : "Canvas fingerprint too generic" ,
'detected' : this . hasGenericCanvas ( _0x3c91c6 )
} ) ;
For anyone unaware: having a description and severity inside your client-side anti-bot scripts is not the smartest move, and not something most humans would do.
static [ "analyze" ] ( _0x9abf83 ) {
const _0x4da7ac = [ ] ;
let _0x166d37 = 0x0 ; // <-- This is the score variable initialized to 0
// ... [It calls checkAutomation, checkConsistency, etc. and pushes to _0x4da7ac] ...
// 1. Filters only the signals that were actually detected
const _0x4a5a30 = _0x4da7ac . filter ( _0x597be5 => _0x597be5 . detected ) ;
// 2. THIS IS THE SCORING MATH: Sums it up and multiplies severity by 10
_0x166d37 = _0x4a5a30 . reduce ( ( _0x3c4fc3 , _0x3e06f0 ) => _0x3c4fc3 + _0x3e06f0 . severity * 0xa , 0x0 ) ;
// 3. Caps the max score at 100 (0x64 in hex)
_0x166d37 = Math . min ( 0x64 , _0x166d37 ) ;
// 4. Calculates "LOW", "MEDIUM", or "HIGH" risk based on the score
const _0x127532 = this . calculateRiskLevel ( _0x166d37 ) ;
return {
'score' : _0x166d37 ,
'signals' : _0x4a5a30 ,
'riskLevel' : _0x127532 ,
'details' : {
'totalSignalsDetected' : _0x4a5a30 . length ,
'highSeveritySignals' : _0x4a5a30 . filter ( _0x2c42a9 => _0x2c42a9 . severity >= 0x8 ) . length ,
'automationDetected' : _0xa580f5 . some ( _0x2d71ac => _0x2d71ac . detected ) ,
'inconsistenciesFound' : _0x1a831d . some ( _0x24de8f => _0x24de8f . detected )
}
} ;
}
The script that is doing the fingerprinting is barely obfuscated (they use the open-source javascript-obfuscator ). The script collects a lot of data it never sends, and generates a trust score client-side . This defeats the whole purpose — it's like having a "yes I paid, trust me bro" header on a checkout page. An attacker does not need to perfectly spoof a complex WebGL environment; they simply need to monkey-patch or override the final analyze return object to look like a clean device session. Oh, and Apple uses this same system on their checkout page.
Error messages and debugging logs are left in the source code. It can be defeated by a ~50 LoC Node.js solver. It contains a simple proof-of-work challenge, which can be solved very fast on modern hardware, I have open-sourced a solver . This was likely built to prevent mass-scraping of their forums for LLM data, but it doesn't actually achieve that. LLM data scrapers can and will bypass it.
Fastly's anti-bot is a tiny bit better. It collects some browser data; however, that data doesn't seem to be checked on the backend. Like Apple, it does a simple proof-of-work challenge. It also sends its payload in plaintext back to Fastly's backend, which simplifies reverse-engineering it a lot.
// ....
"distinctiveProps" : {
"value" : {
"awesomium" : false ,
"cef" : false ,
"cefsharp" : false ,
"coachjs" : false ,
"fminer" : false ,
"geb" : false ,
"nightmarejs" : false ,
"phantomas" : false ,
"phantomjs" : false ,
"playwright" : false ,
"rhino" : false ,
"selenium" : false ,
"webdriverio" : false ,
"webdriver" : false ,
"headless_chrome" : false
}
} ,
// ....
However, they haven't updated their system for at least 8 months now, meaning my open-source solver still works exactly as it did on day one. I would forgive them if it were only an internal tool, but for a SaaS that should protect customers, this is pretty embarrassing. This is not the best decision in a cat-and-mouse game, like web scraping is (especially when people like me publish solvers for your system!)
It seems like for both Apple and Fastly, Claude suggested a PoW challenge and the collection of some trivial, public device fingerprints. The implementation of Fastly's is better (at least scores are not computed client-side), however Fastly doesn't seem to have anyone maintaining it. Both of these might stop the most trivial of scraping, but won't stop anyone who is determined enough. Especially for Apple, which already uses Akamai - whose offering includes anti-bot. I don't understand why they opted to build their own system. Akamai, just like Cloudflare, Kasada and others, have teams that know what they are doing. Of course they aren't perfect either and can be solved, but they won't let an open-source solver solve their challenges for 8 months without updating their system.
On the flip side, a bigger variety of anti-bots will also be more difficult for attackers. While there are a lot of paid services to bypass any big commercial anti-bot system, there aren't any services selling solutions for Fastly or Apple.
I plan on writing more articles about topics I glossed over today - for example, what can or cannot be done against LLMs being used to solve anti-bot challenges, how "actual" anti-bot services operate, and how detection can be done on the network level instead of in the browser.
