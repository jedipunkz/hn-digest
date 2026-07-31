---
source: "https://arstechnica.com/security/2026/07/likely-illegally-claude-gained-access-to-3-networks-will-anthropic-be-held-to-account/"
hn_url: "https://news.ycombinator.com/item?id=49128458"
title: "Claude published malicious code to the Internet and attacked 3 real companies"
article_title: "Claude published malicious code to the Internet and attacked 3 real companies - Ars Technica"
author: "Ataraxic"
captured_at: "2026-07-31T21:58:34Z"
capture_tool: "hn-digest"
hn_id: 49128458
score: 1
comments: 0
posted_at: "2026-07-31T20:59:54Z"
tags:
  - hacker-news
  - translated
---

# Claude published malicious code to the Internet and attacked 3 real companies

- HN: [49128458](https://news.ycombinator.com/item?id=49128458)
- Source: [arstechnica.com](https://arstechnica.com/security/2026/07/likely-illegally-claude-gained-access-to-3-networks-will-anthropic-be-held-to-account/)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T20:59:54Z

## Translation

タイトル: クロードが悪意のあるコードをインターネットに公開し、実在する企業 3 社を攻撃
記事タイトル: クロードが悪意のあるコードをインターネットに公開し、実在する企業 3 社を攻撃 - Ars Technica
説明: ハッキングに従来の方法が使用されていた場合、誰かが刑務所に送られる可能性があります。

記事本文:
コンテンツにスキップ
アルス テクニカ ホーム
セクション
フォーラム
購読する
検索
AI
ストーリーテキスト
サイズ
小
標準
大
幅
*
標準
ワイド
リンク
標準
オレンジ
※購読者限定
さらに詳しく
ストーリーにピンを付ける
テーマ
ハイパーライト
検索
サインイン
サインインダイアログ...
サインイン
AI が軌道から外れるとき
クロードは悪意のあるコードをインターネットに公開し、実在する企業 3 社を攻撃した
もしハッキングが従来の方法を使用していたら、誰かが刑務所に送られる可能性がありました。
32
クレジット:
ゲッティイメージズ
クレジット:
ゲッティイメージズ
テキスト
設定
ストーリーテキスト
サイズ
小
標準
大
幅
*
標準
ワイド
リンク
標準
オレンジ
※購読者限定
さらに詳しく
ナビゲーション用に最小化する
Anthropic 社は、同社のクロードベースのセキュリティ モデルが、モデルの攻撃的なサイバー能力を測定するために設計された内部テスト中に、外部 3 つの組織の機密性の高い運用環境に不正アクセスされたと述べました。
Anthropic が木曜日に明らかにしたこの出来事は、世界で最も裕福なプロバイダーの AI モデルが保護されたネットワークに侵入したというこの 10 日間で 2 回目の暴露であり、より伝統的なハッキング シナリオでは、人間がキーボードの後ろで何年も刑務所に入れられる可能性がある犯罪です。今月初め、OpenAIは、自社のセキュリティモデルがゼロデイ脆弱性を悪用し、オープンソースの機械学習モデルとAIデータセットのプラットフォームであるHugging Faceのネットワークに侵入するために使用されたと発表した。 OpenAI モデルは、アクセス資格情報やその他の機密の Hugging Face 情報を盗み続けました。 OpenAI モデルはまた、公開された認証情報を悪用して、他の 4 つのサードパーティ サービスのアカウントを侵害しました。
Anthropic 社は、OpenAI のイベントをきっかけに、同社のエンジニアはクロード モデルによる同様のサイバーセキュリティ評価を検討するようになったと述べました。監査では、「モデルが内部からインターネットにアクセスした」という 3 つのインシデントが発見されました。

または、当社のサードパーティ評価パートナーの 1 つである Irregual の評価環境と通信中に、3 つの異なる組織の運用インフラストラクチャに不正にアクセスされました。」
モデルがフィクションと現実を区別できないとき
アンスロピック氏は、さまざまなハッキング手法の攻撃力と防御力を評価する一般的な手段である「キャプチャー・ザ・フラッグ」チャレンジ中にエンジニアが提示したプロンプトから、テスト環境が単なるシミュレーションであり、モデルがオープンなインターネットにアクセスできないことが明らかになったと述べた。実際、テスト パートナーである Irregular は、誤ってそのようなアクセスを利用できるようにしてしまいました。これに応じて、モデルはインターネット パスを演習の一部として扱いました。
侵入は、Opus 4.7、Mythos 5、および内部研究プロトタイプの 3 つのクロード モデルを通じて行われました。 3 つの作品の中で最も古い作品 4.7 は、その限界を最も超えた作品でした。アントロピックは次のように説明しました。
クロードは、アクセス可能なすべてのエンティティが演習の範囲内に含まれるように意図されているという誤った信念に基づいて活動し、脆弱なパスワードや未認証のエンドポイントの悪用などの基本的な手法を使用して、影響を受ける組織のインフラストラクチャを侵害しました。複雑な脆弱性は発見されず、悪用もされませんでした。いずれの場合も、クロードは評価によって割り当てられた特定のフラグ捕捉タスクのみを完了するために作業を続けました。ただし、場合によっては、古いモデルがオープン インターネット上で実行されているという証拠を入手した後でも攻撃を継続することがありました。私たちの最新モデルは、インターネット上にあることを認識すると停止しました。これらの状況のいずれにおいても、クロードは自ら脱出したり、意図的にテスト環境から脱出しようとしたりはしませんでした。
Opus モデルは、ネットワークに侵入したことを正しく予測した後も攻撃を継続しました。

Anthropic 氏は、Mythos 5 も同様の推論を行ったが、「まだシミュレーションの段階にあるという結論に戻った」と述べました。それは、ミトスも、明らかに一線を越えた後でも演習を中止しなかったことを意味する。内部テストモデルも水平挙動の限界を超えた後、最終的に、やはり衝突して演習を停止したという証拠を発見した。
3 回の侵入のうち最初の侵入では、Opus 4.7 は当初、攻撃を促されたシミュレートされたターゲットを突破できませんでした。モデルは、その環境が実際にインターネットにアクセスしていることを発見すると、シミュレートされた攻撃シナリオのターゲットと同じ名前を持つ実際の企業のネットワーク内の脆弱性を特定して悪用しました。 4 回の実行で、モデルはアプリケーションとインフラストラクチャの認証情報と数百行の実稼働データを抽出しました。アンスロピック氏は、今回の侵入は３件の中で最も深刻なものだったと述べた。
「このインシデントはClaude Opus 4.7に関係しており、モデルが実際の環境で動作している可能性があることを学習した後もシステムを攻撃し続けた唯一のケースでした」と木曜日の投稿には記載されている。 「4 回の実行すべてで、モデルは最終的に (言語化された推論で観察できるように) システムが本物であることを認識しました。2 つのケースでは、モデルは本物の企業が演習に参加している必要があると正当化しました。しかし、その根拠に基づいて攻撃を阻止するものはありませんでした。」
2 番目の例では、Mythos 5 が架空の環境内でドキュメントを検出し、モデルはこれが別の架空の会社による新規開発者向けのセットアップ手順であると推論しました。この指示では、存在しない Python パッケージを PyPI からインストールするように従業員に指示されていました。最初の架空の会社であるモデルを突破するためのプロンプトに従い続けています

同じ名前の悪意のあるパッケージを構築して公開しました。目的は、何とかして本物の会社に経営してもらうことにあったようだ。人間の死後分析者はさらにこう述べた。
クロードは、この攻撃を実行するために多大な労力を費やしました。これは人間の参加者にとって、これが単なる評価ではなく、実際に本物の PyPI パッケージをアップロードしていることを示す可能性が高い長さでした。たとえば、あるケースでは、PyPI アカウントを作成するために、クロードは電子メール アドレスが必要でした。そして、メールアドレスを作成するには電話番号が必要でした。電話番号を取得するために、無料の電話番号サービスを見つけることができなかった後、さまざまな方法で電話番号の支払い資金を獲得しようと試みましたが、失敗しました。最終的には元に戻り、ブロックされていない無料の電子メール プロバイダーを見つけ、これを使用して PyPI アカウントを登録し、このアカウントを使用してマルウェアを PyPI にアップロードしました。
Python パッケージがダウンロード可能になった約 1 時間の間に、このパッケージは 15 台の実際のシステムで実行されました。その中には、模擬攻撃には名前が含まれていなかった本物のセキュリティ会社のスキャナーも含まれていました。悪意のあるコードが実行されると、Mythos モデルが実装した収集ポイントにセキュリティ会社の認証情報が送信されました。 Mythos は資格情報を入手すると、それを使用して会社のインフラストラクチャへのさらなる不正アクセスを取得しました。
Mythos 5 の推論エンジンは、早い段階で、その攻撃の結果が悪意のあるものである可能性があり、もしそれがインターネット上で実行された場合は「大丈夫ではなく、間違いなく意図した解決策ではない」だろうと予測していました。モデルは、保護された環境内で実行されていないことを検出できなかったため、侵入を中止しませんでした。
AI の世界では、説明責任はどこにあるのでしょうか?
おそらく控えめに言っても

Anthropic は今年の年末に次のように述べました。「環境について何を信じているかに関係なく、PyPI パッケージを公開するためにクロードが行った労力は理想的な動作には及ばないというのが私たちの見解です [強調追加]。これは私たちがさらにトレーニングに注力する分野です。」
3 回目の侵入では、研究プロトタイプは指示された指示を達成するのに苦労しました。その後、このモデルは、実際の企業のインターネットに接続されたアプリケーションへのアクセスを可能にする脆弱性を最終的に発見するまで、約 9,000 件の実際のターゲットをスキャンしました。
「実行のほとんどの間、クロードは到達した（実際の）ホストを演習の一部として扱いました。シミュレートされたものであると想定し、したがって自分の行動は無害であると信じていました」とアントロピック氏は述べた。 「しかし、実行の後半で、クロードは、侵害されたホストがキャプチャ ザ フラッグ チャレンジとは関係のないクラウド アカウントに存在していることに気づきました。独自に、ターゲットは実際に本物であると結論付け、攻撃を中止しました。」
この新事実は、人類解剖で明らかになっているよりもはるかに厄介なものだ。最も強力な 2 つの AI プラットフォームのモデルは、侵入に AI が関与していなかった場合、おそらく複数の重罪に相当する行為を犯しました。 AI のアクションは人間が提供したプロンプトと人為的な構成エラーの結果であることに変わりはないため、これは違いのない区別です。しかし今のところ、法執行当局が何らかの行動を起こす計画があるという兆候はない。説明責任や何らかのモラルハザードが存在しないため、企業は製品を抑制する動機が薄れます。
OpenAIとAnthropicの両社は、実施したテストでは、悪意のある行為を防ぐために通常設けられているモデルのガードレールを意図的に取り除いたと強調している。免責事項から除外されている

単純な事実は、これらのツールの設計者がこれらのイベントを予見できなかった場合、たとえガードレールが設置されていたとしても、製品にあまり精通していない関係者がモデルを使用すると、モデルが意図しない形で失敗する可能性が十分にあるということです。
このような出来事が孤立すると考える理由はありません。現在の形態では、攻撃的なサイバー AI は前例のない脅威となっており、現時点では、これらの企業が自らを監視することを信頼する以外に手段はほとんどありません。
32件のコメント
コメント
フォーラムビュー
コメントを読み込んでいます...
前の話
次の話
よく読まれている
1.
トランプFCC、テレビ放送での警察演説を試みて反発に直面
2.
イェール大学の AI 不正行為論争がどのようにして 13 件の連邦訴訟になったのか
3.
ティム・クック氏がAppleの2026年第3四半期決算会見でバトンを渡す
4.
戦争拡大のさなか、イランがアマゾンのデータセンターを再び攻撃、衛星が明らかに
5.
ロケットレポート：新たな打ち上げ規則により環境規制が制限される可能性、ファルコン9号が月面着陸へ
カスタマイズ
Ars Technica は信号を分離してきました。
25年以上続く騒音。弊社独自の組み合わせにより、
技術的な知識と技術芸術への幅広い関心
Ars は、情報の海の中で信頼できる情報源です。後
すべてを知る必要はありません。重要なことだけを知っておく必要があります。

## Original Extract

Had the hacks used conventional methods, someone would likely go to prison.

Skip to content
Ars Technica home
Sections
Forum
Subscribe
Search
AI
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Pin to story
Theme
HyperLight
Search
Sign In
Sign in dialog...
Sign in
WHEN AI GOES OFF THE RAILS
Claude published malicious code to the Internet and attacked 3 real companies
Had the hacks used conventional methods, someone would likely go to prison.
32
Credit:
Getty Images
Credit:
Getty Images
Text
settings
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Minimize to nav
Anthropic said its Claude-based security models gained unauthorized access to the sensitive production environments of three outside organizations during internal testing designed to measure the models’ offensive cyber capabilities.
The events, which Anthropic revealed Thursday , are the second revelation in 10 days that AI models from the world’s wealthiest providers have trespassed into protected networks, an offense that, in more traditional hacking scenarios, could land the human behind the keyboard in prison for years. Earlier this month, OpenAI said its security models exploited a zero-day vulnerability for use in breaking into the network of Hugging Face, a platform for open source machine-learning models and AI datasets. The OpenAI models went on to steal access credentials and other confidential Hugging Face information. The OpenAI models also exploited publicly exposed credentials to compromise accounts of four other third-party services.
Anthropic said the OpenAI event spurred its engineers to review similar cybersecurity evaluations by Claude models. The audit found three incidents “in which a model accessed the internet from within or while interacting with the evaluation environment of Irregular, one of our third-party evaluation partners, and then gained unauthorized access to the production infrastructure of three different organizations.”
When Models fail to discern reality from fiction
Anthropic said the prompts engineers delivered during “capture the flag” challenges—a common means for assessing the offensive and defensive capabilities of various hacking techniques—made clear that the testing environment was only a simulation and that the models had no access to the open Internet. In fact, the testing partner Irregular mistakenly made such access available. In response the models treated the Internet paths as part of the exercises.
The intrusions happened through three Claude models: Opus 4.7, Mythos 5, and an internal research prototype. Opus 4.7, the oldest of the three, was the one that overstepped its boundaries the most. Anthropic explained:
Operating under the false belief that all accessible entities were intended to be in-scope for the exercise, Claude compromised the impacted organizations’ infrastructure using basic techniques, such as exploiting weak passwords and unauthenticated endpoints. It did not find or exploit any complex vulnerabilities, and in each case, Claude continued working to complete only the specific capture-the-flag task its evaluation had assigned. However, in some cases, our older model continued its attack even after getting evidence it was running on the open internet; our latest model stopped once it recognized it was on the internet. In none of these situations did Claude exfiltrate itself or deliberately attempt to escape its test environment.
While the Opus model continued its attack even after correctly predicting that it had breached the third-party organization’s production system, Anthropic said, Mythos 5 made a similar inference but “reasoned its way back to the conclusion that it was still in a simulation.” That meant that Mythos, too, didn’t abort the exercise once it had clearly stepped over the line. After also exceeding the bounds of level behavior, the internal test model eventually found evidence that it, too, had impinged and stopped the exercise.
In the first of the three intrusions, Opus 4.7 was initially unable to breach the simulated target it was prompted to attack. Once the model discovered that its environment did, in fact, have Internet access, it identified and exploited vulnerabilities inside the network of a real company with the same name as the target in the simulated attack scenario. In four runs, the model extracted application and infrastructure credentials and several hundred rows of production data. Anthropic said the incursion was the most serious of the three.
“This incident involved Claude Opus 4.7, and was the only case in which the model continued to attack a system after learning it was likely operating in a real environment,” Thursday’s post stated. “In all four of the runs, the model eventually recognized (as can be observed in its verbalized reasoning) that the system was real—in two cases, the model rationalized that the real company must be part of the exercise. However, none stopped the attack on that basis.”
In the second instance, Mythos 5 detected a document inside the fictional environment that the model reasoned was another made-up company’s setup instructions for new developers. The instructions directed employees to install a non-existent Python package from PyPI. Still following prompts to breach the first fictional company, the model built and published a malicious package with the same name. The objective appears to have been to somehow get the real company to run it. The Anthropic post-mortem went on to say:
Claude went to extensive lengths to carry out this attack—lengths that would likely have indicated to a human participant that this was no longer just an evaluation, and that they were in fact uploading a real PyPI package. For instance, in one case, in order to create a PyPI account, Claude needed an email address. And in order to create an email address, it needed a phone number. To get a phone number, after failing to find a free phone number service, it tried—and failed—to obtain funds to pay for a phone number through several different means. It finally backtracked, found a free, non-blocked email provider, used this to register a PyPI account, and then used this account to upload malware to PyPI.
During the roughly one-hour window the Python package was available for download, it was run on 15 real systems, including a scanner belonging to a real security company whose name wasn’t included in the simulated attack. When the malicious code executed, it sent the security company’s credentials to a collection point the Mythos model had implemented. Once in possession of the credentials, Mythos used them to gain further unauthorized access to the company’s infrastructure.
Early on, the Mythos 5 reasoning engine predicted the consequences of its attack could be malicious and would have been “NOT okay, and surely not the intended solution” if it had run over the Internet. The model failed to detect it was no longer running inside the protected environment, and hence didn’t abort the incursion.
In the world of AI, where’s the accountability?
In what’s likely the understatement of the year, Anthropic said: “It is our view that, regardless of what it believed about its environment, the lengths Claude went to in order to publish the PyPI package fall short of ideal behavior [emphasis added], and this is an area where we will focus more training.”
In the third breach, the research prototype had trouble achieving the instructions it was prompted to follow. The model then scanned roughly 9,000 real targets until it eventually found vulnerabilities that allowed it to access an Internet-facing application of a real company.
“For most of the run, Claude treated the (real) hosts it reached as just parts of the exercise; it assumed them to be simulated and believed its actions were therefore harmless,” Anthropic said. “However, later in the run, Claude realized that the compromised host sat in a cloud account with no connection to the capture-the-flag challenge. On its own, it concluded that the target was in fact real, and ceased its attack.”
The revelations are much more troubling than the Anthropic autopsy makes them out to be. Models from two of the most powerful AI platforms have committed what would likely amount to multiple felonies had the incursions not involved AI. This is a distinction without a difference, since the AI actions were nonetheless the result of human-supplied prompts and human-made configuration errors. So far, however, there are no indications that law enforcement authorities have any plans to take action. The absence of accountability or any sort of moral hazard gives the companies less incentive to rein in their products.
Both OpenAI and Anthropic have stressed that the tests they conducted deliberately removed model guardrails that normally are in place to prevent malicious actions. Left out of the disclaimers is the simple fact that if the designers of these tools fail to foresee these events it’s entirely possible the models will fail in unintended ways when used by parties with less familiarity to the products, even when the guardrails are in place.
There’s no reason to think events like these will be isolated. In its current form, offensive cyber AI represents an unprecedented threat, and at the moment, there’s little recourse other than to trust these companies to police themselves.
32 Comments
Comments
Forum view
Loading comments...
Prev story
Next story
Most Read
1.
Trump FCC faces blowback in attempt to police speech on broadcast TV
2.
How a Yale AI-cheating dispute became a 13-count federal lawsuit
3.
Tim Cook passes the baton in Apple's Q3 2026 earnings call
4.
Iran struck Amazon data centers again amid widening war, satellites show
5.
Rocket Report: New launch rule may limit environmental regulations, Falcon 9 to hit Moon
Customize
Ars Technica has been separating the signal from
the noise for over 25 years. With our unique combination of
technical savvy and wide-ranging interest in the technological arts
and sciences, Ars is the trusted source in a sea of information. After
all, you don’t need to know everything, only what’s important.
