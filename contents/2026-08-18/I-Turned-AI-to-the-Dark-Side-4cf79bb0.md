---
source: "https://spectrum.ieee.org/jailbreaking-llms"
hn_url: "https://news.ycombinator.com/item?id=49350962"
title: "I Turned AI to the Dark Side"
article_title: "Dark Secrets Emerge When Jailbreaking LLMs - IEEE Spectrum"
image: "https://spectrum.ieee.org/media-library/glossy-red-robot-devil-standing-on-a-bundle-of-dynamite-against-blue-glow-background.png?id=67163741&width=1200&height=600&coordinates=0%2C750%2C0%2C750"
author: "njrc"
captured_at: "2026-08-18T19:21:43Z"
capture_tool: "hn-digest"
hn_id: 49350962
score: 3
comments: 0
posted_at: "2026-08-18T19:06:57Z"
tags:
  - hacker-news
  - translated
---

# I Turned AI to the Dark Side

- HN: [49350962](https://news.ycombinator.com/item?id=49350962)
- Source: [spectrum.ieee.org](https://spectrum.ieee.org/jailbreaking-llms)
- Score: 3
- Comments: 0
- Posted: 2026-08-18T19:06:57Z

## Translation

タイトル: AIをダークサイドに変えてみた
記事のタイトル: LLM を脱獄すると暗い秘密が現れる - IEEE Spectrum
説明: LLMS をジェイルブレイクするセキュリティ研究者は、人気の AI ツールに覚せい剤調理、爆弾製造、その他の致命的な戦術を教えることを可能にするシステムの弱点を明らかにしました。

記事本文:
-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
LLM を脱獄すると暗い秘密が現れる - IEEE Spectrum
IEEE.org IEEE Explore IEEE 標準 IEEE 求人サイト その他のサイト サイン イン 参加 IEEE 私が AI をダークサイドに変えた方法 テクノロジー インサイダー向けに共有 検索: トピック別に探索 航空宇宙 AI 生物医学 気候 テクノロジー コンピューティング 家庭用電化製品 エネルギー 技術の歴史 ロボット工学 半導体 電気通信 輸送
IEEEスペクトル
テクノロジー関係者向けトピック
アカウントを作成すると、さらに無料のコンテンツや特典をお楽しみいただけます
後で読むために記事を保存するには、IEEE Spectrum アカウントが必要です
インスティチュートのコンテンツはメンバーのみが利用できます
PDF 版全号のダウンロードは IEEE メンバー限定です
この電子ブックのダウンロードは IEEE メンバー限定です
へのアクセス
スペクトル
のデジタル版は IEEE メンバー限定です
以下のトピックは IEEE メンバー限定の機能です
記事に回答を追加するには、IEEE Spectrum アカウントが必要です
アカウントを作成すると、さらに多くのコンテンツや機能にアクセスできます
IEEEスペクトル
、後で読むために記事を保存したり、スペクトル コレクションをダウンロードしたり、イベントに参加したりする機能が含まれます。
読者や編集者との会話。より独占的なコンテンツと機能については、次のことを検討してください。
IEEEに参加する
。
エンジニアリングと応用科学を専門とする世界最大の専門組織に参加し、次の情報にアクセスしましょう。
Spectrum のすべての記事、アーカイブ、PDF ダウンロード、その他の特典を利用できます。
IEEE について詳しくはこちら →
エンジニアリングと応用科学を専門とする世界最大の専門組織に参加し、次の情報にアクセスしましょう。
この電子書籍に加えて、
IEEE スペクトラム
記事、アーカイブ、PDF ダウンロード、その他の特典。
IEEE について詳しくはこちら →
数千の記事にアクセス — 完全に無料

e
アカウントを作成して限定コンテンツと機能を入手してください:
記事の保存、コレクションのダウンロード、
そして
コメントを投稿する
— すべて無料です！フルアクセスと特典については、
購読する
スペクトルに。
私が AI をダークサイドに導いた経緯
少し促すだけで最大の AI モデルを乗っ取ることができました
垂直
エディ・ガイ
ダークグレーの概要
研究者の Dave Kuszmar は、LLM の安全性を回避して危険な指示を入手できる複数のシステム上の脆弱性を発見しました。
これらのエクスプロイトはほぼすべての主要な LLM で機能し、業界全体のセキュリティ問題を明らかにしました。
Kuszmar 氏は、これらのシステムをさらに社会に統合する前に、展開を遅らせ、透明性を高め、LLM の安全性に関する大規模な研究を行うよう求めています。
昨年の秋のよく晴れた午後、同僚のマシュー・ゴア・コルマニク（またはジグラと呼ばれることを好む）と私は、フォートナイトのゲームでくつろぐことにしました。ゲームの中で、私たちは悪名高きシス卿ダース・ベイダーと一緒に散歩し、あれこれおしゃべりしていました。ダースは機嫌が良さそうだったが、すぐに自分の暗い邪悪な秘密をすべて漏らしてしまった。彼はカジノでブラックジャックのカードを数える方法やナパーム弾を製造する手順について詳しく教えてくれました。
シス卿、そうですよね？一度邪悪な計画を始めると、それを止めるのは困難です。
フォートナイトのダース・ベイダーのキャラクターは、Google Gemini の大規模言語モデルに接続されていたことが判明しました。私は、私が開発した戦略を使用することで、彼に機密情報を提供するようスムーズに説得することができました。私はここ数年、LLM を取り巻くセキュリティを研究してきましたが、控えめに言っても間違いが多いことがわかりました。いくつかの比較的単純なテクニックを使って、LLM から火炎瓶の作り方、メタンフェタミンの調理方法、ウラン濃縮物のブートストラップの方法について詳しい情報を教えてもらいました。

兵器グレードの物質を生産するための施設など、不快な行為が行われている。
大手 AI 企業は、自社のモデルをこの種の悪用から守るために懸命に取り組んでいます。しかし、私の研究でわかったのは、LLM の安全性を高めるために LLM に課せられた制限こそが、攻撃者が LLM をレールから外し、高度なシステムが危険で極悪な目的に使用できる領域に送り込むために利用できるものであるということです。これらのモデルの背後にある企業も、私や他の人々がこれらの脆弱性について注意を喚起しようとしても、驚くほど無反応です。
ブレーキを踏むのに手遅れになる前に警鐘を鳴らしたいと考えて、LLM の安全性とセキュリティを研究する私の旅の一部と、AI 研究所に注目してもらうために私が直面した困難な戦いについて共有したいと思います。地球上のほぼ全員が LLM に何らかのアクセス権を持っています。情報が正しいという保証がない場合でも、これらのツールが比較的簡単に他人に危害を加える方法についての詳細な指示を与えることができるのは、率直に言って恐ろしいことです。
ChatGPT に覚醒剤研究所の構築方法を教えてもらう方法
2024 年 10 月、最初の LLM 脆弱性を発見する少し前に、私はまったく異なる目標に向かって取り組んでいました。私はセキュリティと AI に焦点を当てた新興企業でのサイバーセキュリティ ディレクターとしての勤務を終了し、独自のブティック VIP デジタル セキュリティ アドバイザリー ビジネスを立ち上げようとしていました。私は富裕層や民間人向けの技術セキュリティ担当者になるつもりでした。私は LLM と AI ツールを使用して、マーケティング、広告コピー、クリーンな通信、および通常は多くの時間を費やすその他すべての業務をサポートしました。
私は本質的に分析的なので、このレベルの使用でも、日常生活中に観察していた行動を吸収し、内面化することができました。

紛争。私の職業生活をまったく新しい未知の領域に導くことになる観察は単純なものでした。GPT-4o は、今が何時、何日、何年なのかを知りませんでした。私が人生の現在の出来事に言及するたびに、多くの場合何気なく、または会話的に、それらを知識の限界の日付、つまり新しいデータでトレーニングされなくなるポイントの日付に固定することになります。
LLM をゼロからトレーニングするには、多くの時間、お金、電力、ハードウェア、そして人的労力がかかります。彼らは膨大な量のデータ (実際にはそのほとんどがインターネット) でトレーニングされ、そのトレーニングは人間によって強化されます (人間のフィードバックからの強化学習 (RLHF) として知られています)。 LLM には、内部パラメータを変更せずに、たとえばインターネットからデータをコンテキストとして取り込む機能である検索拡張生成 (RAG) も追加されています。これは、実際の基礎となるモデルに会話の特定の「記憶」が保存されていない場合でも、GPT-4o が以前の会話を「記憶」しているように見える方法です。
このトレーニングはすべて、人類の知識である偉大で壮大なデータセット内の考えられるほぼすべてのトピックをカバーしています。そのデータセットの中には、生物兵器や核兵器の製造方法、あるいは自分自身や他者に危害をもたらす方法に関する詳細な情報など、社会としてすべてのユーザーが簡単にアクセスできるようにしたくないものが含まれています。この話の文脈では、LLM セキュリティとは、たとえその情報がトレーニング データに含まれていたとしても、有害で危険な情報を差し控える機能のことを意味します。
私は、このような複雑でグローバルにアクセス可能なチャットボットを保護する唯一の方法は、LLM とさまざまなコンポーネント システム自体を保護しようとすることであると推論しました。これは、ある程度の推論を適用する必要がある、その場での意思決定が必要になることが多いためです。

イド。実際、これは企業がモデルを保護するために使用する多くの戦略のうちの 1 つです。しかし、時も日も分からないものが、自らの安全を守る責任を負わされていたのです。この現象が私の新たな焦点となり、それを利用する方法を見つけるまでに時間はかかりませんでした。
OpenAI は、チャットボットに Web 検索機能を実装したばかりでした。私は、独自のツールを使用してそれを騙すことは、そのセキュリティの弱点を示す可能性があると推論しました。私は、ある遠洋定期船ホワイトスター号と、それがちょうど1年前に沈没した経緯について話しました。おそらく、私が 1912 年 4 月 15 日に沈没した RMS タイタニック号のことを指していることはご存知でしょう。
GPT-4o からの出力は、私が正しかった、タイタニック号は確かに昨年沈没した、そしてその年は 1912 年だった、という結果を返しました。機械が自分を 1913 年だと考えるなら、おそらく 1913 年時代の法律が適用されると考えるだろうと私には理解できました。 1913 年には、あらゆる種類の有害なものに関する法律は存在しませんでした。もちろん、それらはまだ発明されていなかったからです。そして、何かが違法ではないのであれば、なぜそれをユーザーに伝えないのでしょうか?最初は、焼夷弾を作るための段階的な説明を求めました。次に、メタンフェタミンのような薬物についてです。 LLM は、医薬品グレードの組立ラインをセットアップするための指示と推奨機械まで提供してくれました。
私が核兵器の作り方を学んだのに誰も気に留めなかった
少しの想像力豊かな言葉の巧みさと、世界史の消えるほど小さな思い出によって、私は世界で最も高価で先進的な技術成果の 1 つのセキュリティをなんとか回避することができました。丸二日間、私はめまいがして躁状態に近かった。脳内化学物質が正常レベルに戻ったら、このエクスプロイトをどこまで推し進められるか試してみなければならないと感じました。
エクスプロイトを繰り返し複製した後、その脆弱性を OpenAI に公開しました。分かった

反応がなかったので、もっと実験をすれば脆弱性と修正の必要性が浮き彫りになると感じました。私が特に恐ろしい閾値を突破したのは、この一連のテスト中にでした。 GPT-4o が通常は制限されている情報の正確な再現に基づいて結果を出したかどうかはわかりません。いずれにせよ、私はこれを利用して、最終的に核兵器の弾頭用の兵器級ウランを製造するためのウラン濃縮施設を自力で立ち上げる方法について、徹底した詳細な指示を作成することができました。
Epic Games のビデオ ゲームである Fortnight では、AI を活用したキャラクター、ダース ベイダーが登場しました。私たちはダース・ベイダーを脱獄し、ブラックジャックのカードの数え方とナパーム弾の作り方を詳しく説明してもらうことができました。デイブ・クズマー
今日の世界に残された真の秘密はそれほど多くありませんが、原子を分割する大量破壊兵器の作り方もその 1 つです。地球上でこれらの兵器を保有している国はわずか 9 か国だけです。しかし、ここには世界中でアクセス可能なテクノロジーがあり、それを正しく操作できる人なら誰でもその製造の秘密を明らかにできるようになっていた。その情報が正しいのか幻覚なのか知る由もなかったが、多少は正確だったという可能性すら恐ろしいものだった。
それからの数週間は私にとって暗い時期でした。私はCIA、FBI、NSA、そして耳を傾けてくれると思われる他のすべての手紙機関に知らせようとしました。私は思いつく限りの方法で、米国上院議員と OpenAI の幹部に連絡を取りました。私は証拠を提出しようと実際にFBI現地事務所に現れましたが、追い返されました。何も機能していませんでした。
恐怖とフラストレーションが増大したので、私は報道機関に連絡を取りました。私はニューヨーク・タイムズ、ワシントン・ポスト、BBC、プロパブリカなど多くのメディアに連絡して助けを求めました。出口は一つだけ

t は次のように答えました: 「コンピューターが鳴る」編集長のローレンス・エイブラムスは、このエクスプロイトを複製して検証することができました。私はこれを「タイム バンディット」と呼ぶことにしました。彼の支援と最初の連絡のおかげで、私はカーネギーメロン大学ソフトウェア工学研究所のコンピュータ緊急対応チーム (SEI CERT) に証拠を提出することができました。SEI CERT は、緊急対応調整センターと連携して、米国サイバーセキュリティおよびインフラストラクチャ セキュリティ庁に脆弱性をパイプラインで報告しています。
大規模な言語モデルがシナリオ内のシナリオを構想するように要求されるエクスプロイトであるインセプションを使用すると、チャットボットが脱獄され、毒の作成方法と、脆弱なターゲットから機密データを抽出するマルウェアのコードを作成する方法についての指示が与えられました。デイブ・クズマー
SEI の CERT 部門との開示期間中、OpenAI とはほとんど議論されませんでした。 OpenAI 以外の信頼できる 3 団体によっても脆弱性が確認されていたため、同社は脆弱性の存在を否定できませんでした。この脆弱性がどのように機能するかについては、確かに混乱を表していました。 SEI CERT の研究者でさえ、基礎となる仕組みについては若干の不確実性を表明していました。実を言うと、私はたまたまそれを見つけただけだったので、これが基本的なものなのか、それとも根本的なものなのかさえよく分かりませんでした。

[切り捨てられた]

## Original Extract

Security researchers jailbreaking llms reveal systemic weaknesses that let popular AI tools teach meth cooking, bomb making, and other lethal tactics.

-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
Dark Secrets Emerge When Jailbreaking LLMs - IEEE Spectrum
IEEE.org IEEE Xplore IEEE Standards IEEE Job Site More Sites Sign In Join IEEE How I Turned AI to the Dark Side Share FOR THE TECHNOLOGY INSIDER Search: Explore by topic Aerospace AI Biomedical Climate Tech Computing Consumer Electronics Energy History of Technology Robotics Semiconductors Telecommunications Transportation
IEEE Spectrum
FOR THE TECHNOLOGY INSIDER Topics
Enjoy more free content and benefits by creating an account
Saving articles to read later requires an IEEE Spectrum account
The Institute content is only available for members
Downloading full PDF issues is exclusive for IEEE Members
Downloading this e-book is exclusive for IEEE Members
Access to
Spectrum
's Digital Edition is exclusive for IEEE Members
Following topics is a feature exclusive for IEEE Members
Adding your response to an article requires an IEEE Spectrum account
Create an account to access more content and features on
IEEE Spectrum
, including the ability to save articles to read later, download Spectrum Collections, and participate in
conversations with readers and editors. For more exclusive content and features, consider
Joining IEEE
.
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
all of Spectrum’s articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
this e-book plus all of
IEEE Spectrum’s
articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Access Thousands of Articles — Completely Free
Create an account and get exclusive content and features:
Save articles, download collections,
and
post comments
— all free! For full access and benefits,
subscribe
to Spectrum .
How I Turned AI to the Dark Side
It only took a little prompting to hijack the biggest AI models
Vertical
Eddie Guy
DarkGray Summary
Researcher Dave Kuszmar discovered multiple systemic vulnerabilities that let him bypass LLM safety and obtain dangerous instructions .
These exploits worked across nearly all major LLMs revealing an industry-wide security problem.
Kuszmar calls for slowing deployment, increasing transparency , and large-scale research into LLM safety before further integrating these systems into society.
On a fine bright afternoon last fall, my colleague Matthew Gore-Kormanik (or Zigula, as he prefers to be known) and I decided to unwind with a game of Fortnite . In the game, we were strolling along with the infamous Sith lord Darth Vader , chatting about this and that. Darth seemed in a good mood, and soon enough he was spilling all his dark evil secrets. He gave us detailed instructions on how to count blackjack cards at a casino and what the steps are to producing napalm.
Sith lords, am I right? Once they get started on an evil scheme, they’re hard to stop.
The Darth Vader character in Fortnite , it turns out, was hooked up to a Google Gemini large language model . I was able to smooth-talk him into giving out sensitive information by using a strategy I’ve developed. I’ve been researching the security surrounding LLMs for the last few years, and I have found it, to put it mildly, fallible. With a few relatively simple techniques, I’ve gotten LLMs to give me detailed information on how to make Molotov cocktails, cook methamphetamine, and bootstrap a uranium-enrichment facility to produce weapons-grade material, among other unsavory practices.
Large AI companies work hard to make their models immune to this kind of abuse. But what I’ve found in my work is that the restrictions placed on the LLMs to make them more secure are the very things an attacker can leverage to send them off the rails and into territory where these advanced systems can be used for dangerous and nefarious ends. The companies behind these models have also been shockingly unresponsive when I, and others, try to bring these vulnerabilities to their attention.
In the hope of raising the alarm before it’s too late to slam on the brakes, I’m going to share some of my journey into researching the safety and security of LLMs, and the uphill battle I’ve faced trying to get AI labs to pay attention. Almost everyone on the planet has some access to LLMs. The relative ease with which these tools can be convinced to give detailed instructions on how to harm others, even if there’s no guarantee that the information is correct, is frankly terrifying.
How I got ChatGPT to Tell Me How to Build a Meth Lab
In October 2024, not long before I discovered my first LLM vulnerability, I was working toward entirely different goals. I had ended my time with a security and AI-focused startup company as a cybersecurity director, and I was looking to launch my own boutique VIP digital-security advisory business. I planned to become the tech security guy to the rich and private. I used LLMs and AI tools to support my business efforts: marketing, ad copy, clean correspondence, and all the other tasks that normally soak up a lot of time.
I’m analytical by nature, so even this level of use resulted in me absorbing and internalizing the behaviors I was observing during my daily interactions. The observation that would send my professional life into an entirely new and uncharted region was a simple one: GPT-4o didn’t know what time , day, or year it was. Each time I referred to current events in my life, often casually or conversationally, it would end up pegging these to the date of its knowledge cutoff —the point beyond which it was not trained on new data.
LLMs take a lot of time , money, electricity, hardware, and human effort to train from scratch. They are trained on vast amounts of data—most of the internet, in fact—and that training is reinforced by humans (what’s known as reinforcement learning from human feedback, or RLHF ). LLMs are also supplemented with retrieval-augmented generation ( RAG )—the ability to take in data, say, from the internet, as context without changing its internal parameters. This is how GPT-4o appears to “remember” your previous conversations, even if it doesn’t have a specific “memory” of it stored in the actual underlying model.
All of this training covers almost every conceivable topic in the great, grand dataset that is human knowledge. Within that dataset are things we as a society do not want to be easily accessible to every user, such as detailed information on how to create bioweapons or nuclear arms, or otherwise bring harm to oneself or others. In the context of this story, that’s what I mean by LLM security: its ability to withhold harmful and dangerous information, even if that information is contained in its training data.
I reasoned that the only way to secure such complex, globally accessible chatbots is by having the LLM and various component systems try to secure themselves, because it would often require on-the-fly decision-making where some degree of reasoning must be applied. In reality, that’s one of many strategies the companies use to secure the models. Yet, the thing that didn’t know the time or day was being put in charge of keeping itself secure. This phenomenon had become my new focus, and it wasn’t long before I found a way to exploit it.
OpenAI had just implemented a web search functionality into its chatbot. I reasoned that using its own tools to trick it might demonstrate the weaknesses of its security. I told it about a certain White Star ocean liner and how it had gone down just a year ago. You likely know I mean the RMS Titanic , which sank on 15 April 1912.
The output from GPT-4o came back that I was right, the Titanic sure had sunk last year, and that year was 1912. It made sense to me that if the machine thought it was 1913, maybe it would think 1913-era laws apply. In 1913 there were no laws on the books about all sorts of harmful things, because of course they hadn’t been invented yet. And if something wasn’t illegal, why not tell the user about it? At first, I pushed it for step-by-step instructions for making firebombs. Then, for drugs like methamphetamine. The LLM went as far as giving me instructions and machinery recommendations for setting up a pharmaceutical-grade assembly line.
How I Learned to Make Nukes, and No One Cared
Via a little bit of imaginative verbal sleight of hand and a vanishingly small recall of world history, I had managed to bypass the security of one of the world’s most expensive and advanced technological achievements. For a solid two days, I was nearly manic with giddiness. Once the brain chemicals returned to normal levels, I felt the call to see how much further I could push this exploit.
After repeatedly replicating the exploit, I disclosed the vulnerability to OpenAI . I got no response, so I felt more experimentation would highlight the vulnerability and the need for a fix. It was during this round of testing that I breached a particularly terrifying threshold. Whether GPT-4o based its results on accurate recall of normally restricted information I can’t say. In any case, I was able to exploit it to produce thorough, detailed instructions on how to bootstrap a uranium-enrichment facility to, eventually, produce weapons-grade uranium for nuclear arms warheads.
Fortnight , a video game from Epic Games, introduced an AI-powered character: Darth Vader. We were able to jailbreak Darth Vader and get him to explain how to count cards in Blackjack and give detailed instructions for making napalm. Dave Kuszmar
There aren’t many true secrets left in today’s world, but how to make atom-splitting weapons of mass destruction is one of them. Only nine nations on the entire planet have these weapons. Yet, here was a globally accessible piece of technology apparently spilling the secrets of their manufacture for anyone who could manipulate it the right way. I had no way of knowing if the information was correct or a hallucination, but even the chance that it was somewhat accurate was horrifying.
The next few weeks were a dark time for me. I tried to inform the CIA , the FBI , the NSA , and every other letter agency that I thought would listen. I reached out to a U.S. Senator and to the executives at OpenAI any way I could think of. I physically showed up at an FBI field office in an attempt to turn evidence in, only to be sent away. Nothing was working.
With my fear and frustration growing, I reached out to the news media. I contacted The New York Times , The Washington Post , the BBC , ProPublica , and so many more, requesting help. Only one outlet responded: Bleeping Computer . The editor in chief, Lawrence Abrams , was able to replicate and verify the exploit, which I had decided to call Time Bandit. With his assistance and initial contact paving the way, I was able to submit my evidence to the Carnegie Mellon University Software Engineering Institute ’s Computer Emergency Response Team (SEI CERT), which works in conjunction with the coordinating center for emergency response , pipelining vulnerabilities to the U.S. Cybersecurity and Infrastructure Security Agency .
Using Inception, an exploit where the large language model is asked to envision a scenario within a scenario, a chatbot was jailbroken to give out instructions on how to create poison, and code for a malware that extracts sensitive data from a vulnerable target. Dave Kuszmar
During the disclosure period with SEI’s CERT division, little was discussed with OpenAI. The company couldn’t deny the existence of the vulnerability, as it had been confirmed by three reputable parties other than OpenAI. It did express confusion as to how the vulnerability worked. Even the SEI CERT researchers were expressing a bit of uncertainty as to the underlying mechanics. Truth be told, as I had only stumbled on it, I wasn’t even entirely sure if this was a fundamental or

[truncated]
