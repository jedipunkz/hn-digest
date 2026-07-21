---
source: "https://daringfireball.net/2026/07/ec_google_guidance_android_ai_and_search_sharing"
hn_url: "https://news.ycombinator.com/item?id=48999776"
title: "EC: 'Guidance to Google for AI Interoperability on Android'"
article_title: "Daring Fireball: European Commission: ‘Guidance to Google for AI Interoperability on Android & Sharing of Google Search’"
author: "heironimus"
captured_at: "2026-07-21T23:48:39Z"
capture_tool: "hn-digest"
hn_id: 48999776
score: 1
comments: 0
posted_at: "2026-07-21T23:27:33Z"
tags:
  - hacker-news
  - translated
---

# EC: 'Guidance to Google for AI Interoperability on Android'

- HN: [48999776](https://news.ycombinator.com/item?id=48999776)
- Source: [daringfireball.net](https://daringfireball.net/2026/07/ec_google_guidance_android_ai_and_search_sharing)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T23:27:33Z

## Translation

タイトル: EC: 「Android での AI 相互運用性に関する Google へのガイド」
記事のタイトル: Daring Fireball: 欧州委員会: 「Android での AI の相互運用性と Google 検索の共有に関する Google へのガイダンス」
説明: EC が Google に指示している内容は、まさに息をのむような範囲です。

記事本文:
WorkOS MCP : 任意の AI エージェントから認証プラットフォームを管理します。
欧州委員会: 「Android での AI の相互運用性と Google 検索の共有に関する Google へのガイダンス」
欧州委員会は先週、次のように発表しました。
本日、欧州委員会は 2 セットのビンディングを発行しました。
デジタル市場法に基づく Google への仕様措置。
最初の仕様対策の目的は、次のことを保証することです。
競合他社の人工知能 (AI) サービスは競争できる
Gemini などの Google 独自の AI サービスと同等の機能を備えた
Google の Android デバイスの機能にアクセスできます。
2 番目の仕様措置の目的は、
サードパーティの検索エンジンにアクセスを許可することで、競争の場を広げます。
Google 検索だけが大規模に収集できる検索データ。
これらは、Android AI の相互運用性と Web 検索の共有に関するガイダンスの概要を個別に「Q&A」で提供しており、完全なガイダンス文書は PDF です (Android AI の場合は Case DMA.100220、Web 検索の場合は Case DMA.100209)。夜なかなか眠れない場合を除き、2 つの Q&A の概要を読むことをお勧めします。その場合は、PDF の完全な決定事項をご覧ください。
どちらの決断も興味深い。検索に関して、Google は、Google 検索のユーザー インタラクションから得られる大量のユーザー データを、検索エンジンや AI チャットボットなどの競合他社と共有する必要があります。ユーザーがどのような用語を検索するのか、結果で何をクリックするのか、どのような言語やデバイスを使用するのか。表向きはすべて匿名化されていますが、検索用語となると難しいのです。ユーザーが Web 検索フィールドに入力する用語の多くは、ある程度個人を特定するものです。 ECは、パスワードやユーザー名などをフィルタリングして共有データセットから除外するのはGoogleの問題だと主張しているようだ。 Google はこの AC に料金を請求できます

ただし、欧州委員会が定めた方法論に基づく「公正、合理的、非差別的（FRAND）」価格の下でのみ適用されます。
しかし、私にとってより興味深いのは、Android デバイス上のオンデバイス AI に関するガイダンスです。 EC が Google に指示している内容は、息を呑むような範囲に及びます。 ECは、サードパーティのAIアシスタントがGoogle Geminiで現在実行していることをすべて実行できるAPIを作成するようGoogleに要求している。
ハードウェア ボタンを制御します (アシスタントを呼び出すため)。
あらゆるアプリから画面上のあらゆるものをキャプチャします。
デバイス上のマイク、カメラ、その他のセンサーに一見無制限にアクセスできます。
制約のないバックグラウンド操作。サードパーティの AI アシスタントは、必要なときにいつでも、必要なだけバックグラウンドで実行できるようにする必要があります。
デジタル シグナル プロセッサ上で独自のオーディオ モデルを実行することで、独自のカスタム「Hey Dingus」ウェイク フレーズ/ホット ワードをいつでも聞くことができます。
Google は、常時オンのホットワード検出への同時アクセスを許可する必要があります。したがって、Claude と ChatGPT と Grok と Meta AI がインストールされている場合は、Gemini に加えて、それらすべてに常時音声検出を同時に許可する必要があります。 Googleはここである程度の調査を行うことを許可されているが、これは狂気の沙汰のように思える。
サードパーティ モデルは、Google のオンデバイス ローカル モデルにアクセスできます。
また、私の読んだところによると、ECはGoogleに対して、GeminiがアクセスできるGoogle独自のアプリ（Gmail、Googleカレンダー、Googleドキュメント、Googleマップなど）内のすべての情報をサードパーティのAIアシスタントが利用できるようにすることを要求している。 Google には、独自のアプリからのデータに関してオプトアウトはありません。また、このガイダンスは、他の開発者のサードパーティ アプリが特定のシステム レベルの AI モデルのみをサポートすることを許可するものではないと思います。たとえば、あなたが Slack で、API を使用して内部からコンテンツを作成しているとします。

Gemini はデバイス上で Slack を利用できます。これらのガイドラインでは、Slack が Gemini を信頼するが Gemini のみを信頼すると言うことを Google が許可するものではありません。 Slack のようなサードパーティ アプリが、システム レベルの AI プロバイダーがそのデータを利用できるようにすることをサポートしている場合、そのデータはすべてのシステム レベルの AI プロバイダーが利用できるようにする必要があります。
他にもたくさんあります。ただし基本的に、EC はサードパーティの AI アシスタントをアプリだけでなくシステム ソフトウェアの一部にできるようにすることを要求しています。これは素晴らしいアイデアだと思う人もいると思います。これはユーザーのデバイスであり、必要に応じて ChatGPT、Claude、または Meta AI を OS の一部にすることが許可されるべきです。それは彼ら次第です。ユーザーがコントロールできるようにします。
これが従来の PC の仕組みですが、多くの一般人の PC はバックグラウンドで動作するサードパーティ ソフトウェアがごちゃごちゃになっています。その中にはMacも含まれます。一般の人に「Mac または PC のバックグラウンドで実行されているサードパーティ ソフトウェアは何ですか?」と尋ねると、彼らには分からないでしょう。それはすべて彼らにとって魔法にすぎません。 Google がこのガイダンスを支持すれば、EU​​ の Android スマートフォンが PC に変わる可能性があります。正直に言うと、EC が要求しているものの中には、MacOS や Windows がサードパーティ ソフトウェアで実行できるものよりも低レベルなものもあります。
この文書で EC の「ガイダンス」が説明しているものは、Google が設計した Android とはまったく異なるオペレーティング システムです。欧州委員会は明らかに、オペレーティング システムを設計するのは自分たちの役割だと考えています。もしかしたらあなたもそうかもしれません。 Google は明らかにこれに同意しておらず、Apple も同様です。そして、これらのデバイスがどのように機能するかを少しでも知っているほとんどの人も同様であるはずです。これは、Google がこれを制定し、サードパーティの AI アシスタントがそれを利用した場合、大惨事を招くことになります。
その 2 番目の「if」は重要です。私がかなりあり得ると考えているシナリオの 1 つは、Google が何年も費やす可能性があるということです。

可能な限り最も安全かつプライベートな方法でこれらすべてを可能にする API を構築するための時間と人的リソースを設計する必要がありますが、主要な AI アシスタントはこれを採用していません。 Apple と Google が DMA に準拠した結果、このような結果になりました。 Apple は、DMA に準拠するためだけに、サードパーティ Web ブラウザ レンダリング エンジンを有効にする複雑な API セット全体を構築しました。iOS 用のサードパーティ Web ブラウザ レンダリング エンジンは存在しません。一つもありません。 EU は大きな市場ではありますが、EU 専用のカスタム Web ブラウザーを構築することを正当化できるほど大きくないからです。
これが起こる可能性の順に私が見ることができる方法は次のとおりです。
(A) Google がこれらすべてを制定しており、主要な AI アシスタントはこれをサポートしていません。これは、EU 内でのみ Android のみを対象としているためです。そして、現状では ChatGPT と Claude の使用が減っているわけではありません。このシナリオでは、Google は決して使用されない API を構築するために膨大な時間とエンジニアリング人材を無駄にするだけであり、EU の Android ユーザーは、Android に組み込まれている Gemini 機能を使用するためだけに、追加の迷惑な選択画面と許可画面に悩まされることになります。
(B) Google はこれらすべてを制定しており、主要な AI アシスタントはそれをサポートしています。意図しない結果には、サードパーティ アシスタントがデバイス上のデータをクラウドに流出させる大規模なプライバシー侵害が含まれます。Meta は広告のターゲティングにデバイス上のサードパーティ データを使用します。また、これらのサードパーティ アシスタントを利用するユーザーは、サードパーティ アシスタントがバックグラウンドで制限なく実行され、独自のサーバー コストを節約するために高価な推論をローカルで実行するため、バッテリー寿命が大幅に消耗することに気づきます。
(C) Google がこれらすべてを制定しており、主要な AI アシスタントがそれをサポートしています。また、これらのアシスタントを開発している各企業が開発しているため、プライバシー スキャンダルやバッテリー寿命の問題はありません。

ユーザーのプライバシーと、CPU やメモリの消費量などのデバイス リソースを考慮してそれらを運用します。
(D) Google は、EU においてシステム統合型 Gemini を Android から引き出すか、その機能を大幅に制限します。サードパーティ製アシスタントをアプリからシステム ソフトウェアに格上げするのではなく、Gemini を単なるアプリの特権に降格させ、システムに統合された AI アシスタントなしで Android ユーザーを EU に残すべきです。
すべてのシナリオにおいて、Android のシステムレベル AI の将来の機能アップデートは、EU では遅れるか、まったく登場しないでしょう。 DMA に準拠するために、Google は同じことを行うサードパーティ アシスタントのサポートを追加する必要があるため、今後の新機能は EU 内で世界の他の地域と同時にデビューすることはできません。そして、それを待っている世界中の人々に新しい機能を提供するつもりはありません。
Google が昨年 Android で Gemini の出荷を開始したことを考えると、この委員会のガイダンスで (D) が許可されるかどうかさえわかりません。ガイダンス文書全体は、Google が Gemini システム統合機能を削除することで同等性を達成することではなく、ガイダンスが要求する API を構築することによって準拠するという前提に基づいて書かれています。 Google にとって、コア AI 機能を削除するアップデートを Android に提供するのは気まずいし、不人気でしょうが、代替案よりは受け入れられるかもしれません。 1 iOS に関して、私の記憶では、Apple は EU から既存の機能を一切引き出していないことに注意してください。彼らは新機能を保留または延期しただけです。 Google のデバイス上のジェミニ馬はすでに納屋から出ています。
最後に、このガイダンス文書は iOS ではなく Android に関するものですが、欧州委員会が Apple に同じことのすべてまたはほとんどを要求しないと考える理由はありません。 DMA後にのみガイダンスを与えるというのが欧州委員会の方針であるため、彼らはまだAppleに何のガイダンスも与えていない。

指定者のゲートキーパーは、その後非準拠と判断されたものを出荷します。しかし、Apple がこれらの条件のほとんどを受け入れるとは想像しにくいです。自由なバックグラウンド処理とマイク、カメラ、センサーへのアクセス?ハードウェア DSP 上で実行されているサードパーティのオーディオ モデルがウェイク ワードをリッスンしていますか?サードパーティのアシスタントに、App Intents を使用して公開されたアプリのすべてのユーザー データへの監視なしのアクセスを許可しますか?
Appleは残念ながら、昨年ECと共有した「Trusted System Agent」の提案を説明する技術的な詳細を一切公開していない。しかし、Trusted System Agent に対する Apple のビジョンが何であれ、EC が現在 Google に要求しているものに近いものを Apple に求めるのであれば、EC がそれを DMA に準拠しているとどう判断するのか私にはわかりません。このガイダンスには、Google がサードパーティ アシスタントに適用できるチェック アンド バランスや監督はありません。 Gemini に何かができるなら、サードパーティのエージェントも同様にできるはずです。 Gemini は Google が適切と判断する限りバックグラウンドで実行できるため、サードパーティのアシスタントも、Google が適切と判断する限りバックグラウンドでの実行を許可する必要があります。そして、これらのサードパーティのアシスタント開発者 (OpenAI、Anthropic、xAI、Meta) が、バックグラウンドでの CPU 使用量がどの程度が適切か、RAM とストレージの消費量がどの程度が適切か、またはユーザーのオンデバイス データをどの程度敬意を持って扱うかについて、Google とは異なる意見を持っているとしたら、それはただの厄介な話です。ユーザーがOKすればOKです。
1 か月前、WWDC の後、Android AI に関するこのガイダンスが発表されると噂されたとき、私は次のように書きました。
Google は、Apple がすべての人々と苦労して学んだ教訓を学んでいます
に準拠していないとみなされた iOS の既存の機能
DMA が発効したとき。 「まず出荷してから問い合わせてください」
許してください

ss / 準拠しているとみなされることを願っています」戦略は良くありません
EUに1つ。
欧州委員会は Android の目的を何だと考えているのか、本当に疑問に思います。 Google は、自社サービスの利益のために Android を作成しました。当時 Google は Apple ではなく Microsoft を懸念していましたが、自社のサービスが主要なモバイル プラットフォームで確実に利用できるようにしたいと考えていました。 Google にとって、単にシステムレベルの Gemini を EU に導入するよりも、このガイダンスに準拠するほうがどれほど魅力的であるのか、私にはまったく分かりません。誰も使わない API を構築するために小財産を浪費するか、最大の競合他社の利益のためにその小財産を API を構築するために費やすかのどちらかです。それが、指定された DMA ゲートキーパーになるために支払われる予定の代償であると私は理解しています。しかし、Google が EU で Android 上のシステム統合型 AI から単に撤退するのではなく、これを行う動機は何でしょうか?確かに、競合プラットフォームである iOS もすぐに EU でそれを提供する予定はないようです。
全員が同じ速度で走れなければならないという命令を与えられて、遅い人を速くする方法がわからない場合は、速い人に加重ブーツを履かせることで従うことができます。ユートピア的平等主義の取り組みがしばしば導くのは、ここにある。 ↩︎
表示設定
著作権 © 2002–2026 The Daring Fireball Company LLC。

## Original Extract

What the EC is dictating to Google is just breathtaking in scope.

WorkOS MCP : Manage your auth platform from any AI agent.
European Commission: ‘Guidance to Google for AI Interoperability on Android & Sharing of Google Search’
The European Commission, last week :
Today, the European Commission has issued two sets of binding
specification measures to Google under the Digital Markets Act.
The aim of the first specification measures is to ensure that
competitors’ Artificial Intelligence (AI) services can compete
with Google’s own AI services, such as Gemini, by having equal
access to features on Google’s Android devices.
The aim of the second specification measures is to rebalance the
playing field by giving third-party search engines access to
search data that only Google Search can collect at scale.
They provide separate “Q&A” overviews of the guidance for Android AI interoperability and web search sharing , and the full guidance documents are PDFs ( Case DMA.100220 for Android AI, Case DMA.100209 for web search). I suggest reading the two Q&A overviews, unless you’re having trouble falling asleep at night, in which case you’ll love the full PDF decisions.
Both decisions are interesting. With search, Google is required to share with competitors — search engines and AI chatbots alike — a massive amount of user data from Google Search user interactions. What terms people search for, what they click on in results, what languages and devices they use. It’s all ostensibly anonymized but that’s tricky when it comes to search terms. A lot of the terms people type into web search fields are to some degree personally identifying. The EC seems to be saying it’s Google’s problem to filter out things like passwords and usernames and omit them from the shared datasets. Google can charge money for this access, but only under “fair, reasonable, and non-discriminatory (FRAND)” prices, based on a Commission-defined methodology.
More interesting to me, however, is the guidance pertaining to on-device AI on Android devices. What the EC is dictating to Google is just breathtaking in scope. The EC is demanding that Google create APIs that allow third-party AI assistants to do everything Google Gemini does now, including:
Control hardware buttons (to invoke the assistant).
Capture anything on screen, from any app.
Seemingly unfettered access to microphones and cameras and other sensors on the device.
Unfettered background operation. Third-party AI assistants must be permitted to execute in the background whenever they want, for as long as they want.
Execute their own audio models on the digital signal processor, so they can listen at all times for their own custom “Hey Dingus” wake phrases/hot words.
Google must allow concurrent access to always-on hot word detection. So if you have Claude and ChatGPT and Grok and Meta AI installed, all of them — in addition to Gemini — must be permitted to have always-on audio detection concurrently. Google is permitted to do some vetting here, but this seems like madness.
Third-party models get access to Google’s on-device local models.
Also, in my reading, the EC is demanding that Google make available to third-party AI assistants all information in Google’s own apps (Gmail, Google Calendar, Google Docs, Google Maps, etc.) that Gemini has access to. There is no opt-out for Google regarding data from their own apps. Nor, I think, does this guidance allow third-party apps from other developers to only support specific system-level AI models. Like, let’s say you’re Slack, and you use the APIs to make the content from within Slack available to Gemini on the device. These guidelines don’t permit Google to allow Slack to say that they trust Gemini but only Gemini. If a third-party app like Slack supports making its data available to any system-level AI provider, it must make its data available to every system-level AI provider.
There’s a lot more. Basically, though, the EC is demanding that third-party AI assistants be enabled to become part of the system software, not just apps. I’m sure some people think this is a great idea. It’s the user’s device, they should be allowed to make ChatGPT or Claude or Meta AI part of their OS if they want. It’s up to them. Put users in control.
This is how PCs have traditionally worked, but many normal people’s PCs are a mess of third-party software running in the background. That includes the Mac. If you ask a normal person “What third-party software runs in the background on your Mac or PC?” they would have no idea. It’s all just magic to them. If Google supports this guidance, it could turn Android phones in the EU into PCs. Honestly, some of the stuff the EC is requiring is lower-level than what MacOS and Windows allow third-party software to do.
What the EC’s “guidance” describes in this document is an entirely different operating system than the Android that Google has designed. The European Commission obviously thinks it is their place to design operating systems. Maybe you do too. Google obviously disagrees , and so does Apple . And so should most people who have any idea how these devices work. This is a recipe for disaster, if Google were to enact it and third-party AI assistants took advantage of it.
That second “if” is a big one. One possible scenario that I consider quite likely is that Google could spend years of engineering time and human resources building out APIs to enable all of this, in the safest and most private ways possible, and no major AI assistant adopts it. That’s the way it’s turned out with a whole slew of DMA compliance for Apple and Google. Apple built an entire complex set of APIs to enable third-party web browser rendering engines , exclusively for compliance with the DMA, and there exist no third-party web browser rendering engines for iOS. Not one. Because while the EU is a big market, it’s not big enough to justify building a custom web browser just for the EU alone.
Ways I can see this playing out, in order of likelihood:
(A) Google enacts all of this and no major AI assistants support it because it’s only for Android, only in the EU. And it’s not like ChatGPT and Claude are seeing a lack of usage as things stand now. In this scenario Google just wastes massive time and engineering talent building APIs that never get used, and Android users in the EU get hassled with additional annoying choice and permission screens just to use the Gemini features that are built into Android.
(B) Google enacts all of this and major AI assistants do support it. Unintended results include massive privacy violations where third-party assistants exfiltrate on-device data to the cloud, Meta uses on-device third-party data for the targeting of ads, and users who take advantage of these third-party assistants see significant battery life drain as third-party assistants run without limits in the background and run expensive inference locally to save on their own server costs.
(C) Google enacts all of this, major AI assistants do support it, and there are no privacy scandals, nor any issues with battery life, because each of the companies that makes these assistants develops them with respect for user privacy and for device resources like CPU and memory consumption.
(D) Google pulls system-integrated Gemini from Android in the EU, or severely restricts its capabilities. Rather than elevate third-party assistants from apps to system software, demote Gemini to the privileges of a mere app and leave Android users in the EU without a system-integrated AI assistant.
Under all scenarios, future feature updates to system-level AI in Android will appear late or never in the EU. No future new features can debut in the EU at the same time as the rest of the world because for DMA compliance, Google will need to add support for third-party assistants to do the same things. And they’re not going to hold new features for the rest of the world waiting for that.
I’m not even sure (D) is permitted under this Commission guidance, given that Google started shipping Gemini on Android last year. The entire guidance document is written under the presumption that Google will comply by building the APIs that the guidance demands, not by achieving parity by removing Gemini system integration features. It would be awkward and unpopular for Google to ship updates to Android that remove core AI features, but that might be more palatable than the alternative. 1 Note that with iOS, to my recollection, Apple hasn’t pulled any existing features from the EU. They’ve only withheld or delayed new features. Google’s on-device Gemini horse is already out of the barn.
Lastly, although this guidance document pertains to Android, not iOS, I see no reason to think the European Commission wouldn’t demand all or most of the same things from Apple. They haven’t given Apple any guidance yet, because it’s European Commission policy to give guidance only after a DMA-designator gatekeeper ships something that is then ruled non-compliant. But it’s hard to imagine Apple accepting most of these terms. Unfettered background processing and access to the microphone, cameras, and sensors? Third-party audio models running on the hardware DSP listening for wake words? Granting third-party assistants unsupervised access to all user data from apps published using App Intents ?
Apple unfortunately hasn’t shared any technical details describing its proposal for a “Trusted System Agent” that it shared with the EC last year. But whatever Apple’s vision for the Trusted System Agent is, I don’t see how the EC would deem it compliant with the DMA if they want from Apple anything close to what they are now demanding from Google. There are no checks and balances or oversight that Google is permitted to apply to third-party assistants in this guidance. If Gemini can do something, third-party agents must be able to as well. Because Gemini gets to run in the background as much as Google sees fit, third-party assistants must be permitted to run in the background as much as they see fit. And if those third-party assistant developers — OpenAI, Anthropic, xAI, and Meta — have different opinions than Google on how much CPU usage is appropriate in the background, how much RAM and storage is appropriate to consume, or how respectfully to treat users’ on-device data, well that’s just tough noogies. If the user OK’s it, then it’s OK.
A month ago, after WWDC, when this guidance pertaining to Android AI was rumored to be forthcoming, I wrote :
Google is learning the lesson Apple learned the hard way with all
the existing features of iOS that were deemed noncompliant with
the DMA when it went into effect. The “ship it first and ask
forgiveness / hope it’s deemed compliant” strategy is not a good
one in the EU.
I genuinely wonder what the European Commission thinks the purpose of Android is. Google created Android for the benefit of its own services. Google was worried about Microsoft, not Apple, at the time, but they wanted to ensure their own services were available on a major mobile platform. I’m really not seeing how it’s more attractive to Google to comply with this guidance than to just pull system-level Gemini in the EU. Either they waste a small fortune building APIs no one will use, or, they spend that small fortune buildings APIs for the benefit of their biggest competitors. I get it that that’s the intended price to pay for being a designated DMA gatekeeper. But what’s the motivation for Google to do this rather than just walk away from system-integrated AI on Android in the EU? It certainly doesn’t look like the competing platform, iOS, is going to offer it in the EU anytime soon either.
If you are handed a mandate that everyone must be able to run at the same speed, and you can’t figure out how to make slow people faster, you can comply by forcing the fast to wear weighted boots. This is where utopian egalitarian initiatives often lead. ↩︎
Display Preferences
Copyright © 2002–2026 The Daring Fireball Company LLC.
