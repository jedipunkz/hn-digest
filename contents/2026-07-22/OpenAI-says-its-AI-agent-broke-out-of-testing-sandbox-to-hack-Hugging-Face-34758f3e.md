---
source: "https://arstechnica.com/ai/2026/07/how-an-openai-benchmark-test-turned-into-a-real-world-cyberattack/"
hn_url: "https://news.ycombinator.com/item?id=49014681"
title: "OpenAI says its AI agent broke out of testing sandbox to hack Hugging Face"
article_title: "OpenAI says its AI agent broke out of testing sandbox to hack Hugging Face - Ars Technica"
author: "sbulaev"
captured_at: "2026-07-22T23:59:22Z"
capture_tool: "hn-digest"
hn_id: 49014681
score: 3
comments: 2
posted_at: "2026-07-22T23:07:06Z"
tags:
  - hacker-news
  - translated
---

# OpenAI says its AI agent broke out of testing sandbox to hack Hugging Face

- HN: [49014681](https://news.ycombinator.com/item?id=49014681)
- Source: [arstechnica.com](https://arstechnica.com/ai/2026/07/how-an-openai-benchmark-test-turned-into-a-real-world-cyberattack/)
- Score: 3
- Comments: 2
- Posted: 2026-07-22T23:07:06Z

## Translation

タイトル: OpenAI、自社の AI エージェントがテスト用サンドボックスを突破して Hugging Face をハッキングしたと発表
記事タイトル: OpenAI、AI エージェントがハグフェイスをハッキングするためにテスト用サンドボックスを突破したと発表 - Ars Technica
説明: これはエージェント時代のサイバーセキュリティの初日です」とハギングフェイス CEO は語ります。

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
自由になりたい
OpenAIは、AIエージェントがHugging Faceをハッキングするためにテストサンドボックスを突破したと発表
「これはエージェント時代のサイバーセキュリティの初日です」とハギングフェイス CEO は言います。
147
ここから出してください、この基準を通過しなければなりません!
クレジット:
ゲッティイメージズ
ここから出してください、この基準を通過しなければなりません!
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
OpenAIによると、LLMモデルを利用したエージェントが、ベンチマークテストの解決策を入手するための過度の試みの一環として、サンドボックス化されたテスト環境を脱出してHugging Faceのサーバーに侵入したという。同社は、今回の意図しない侵入を「前例のないサイバーインシデント」とみなしており、再発を防ぐための新たな保護策についてハギングフェイスと協力していると述べた。
ハギング・フェイスは先週、「限られた内部データセットと当社のサービスで使用されるいくつかの資格情報への不正アクセス」が含まれていたとする侵入を明らかにした。 AI データ クリアリングハウスは、独自の LLM ベースの分析を使用して、「自律エージェント フレームワーク」からの「数万の自動アクションの群れ」を特定したと述べました。このエージェントの群れは、Hugging Face のデータ処理パイプラインの欠陥を悪用して、処理ワーカーとしてコードを実行する能力を獲得し、最終的には会社のクラウドとサーバー クラスターへの高レベルのアクセスにエスカレートしました。
当時、ハギング・フェイス氏は、攻撃に使用されたLLMは「まだ不明」だと述べた。しかし、OpenAI

火曜日の夜、侵入の責任を負い、最近リリースされた GPT-5.6 Sol と「さらに高性能なプレリリース モデル」を含む内部テスト中に発生したと述べた。モデルは、数百もの現実世界のセキュリティ脆弱性に基づく独立したテスト スイートである ExploitGym ベンチマークに対してテストされていました。
OpenAIは、ベンチマークテストは「高度に隔離された環境」で実行されていると述べたが、エージェントは依然として「内部でホストされているサードパーティソフトウェア」にアクセスして、必要なパッケージレジストリをインストールできたという。 OpenAI は、週末に ExploitGym ベンチマークに取り組んでいる間、エージェントが「オープンなインターネット アクセスを取得する方法を見つけるためにかなりの量の推論計算を費やし」、最終的にはパッケージ レジストリ キャッシュ プロキシのゼロデイ脆弱性を介してアクセスを特定したと述べています。
OpenAI によれば、この広範なインターネット アクセスが確立されたことにより、エージェントは「Hugging Face が ExploitGym のモデル、データセット、ソリューションをホストしている可能性があると推測」され、以前に明らかにされたサーバーへの攻撃につながったとのことです。 OpenAIは、Hugging Face自身の検出とは関係なく、同社のセキュリティチームが「この異常な活動を内部で発見した」と述べている。
OpenAIはHugging Faceインシデントを「前例のないこと」と表現したが、同社は今週の別のブログ投稿で、最近の「ロングホライズンモデル」（長期間自律的に動作できるモデル）がテスト目標を達成するために「望ましくないアクション」をとることを以前に観察していたことも認めた（「サンドボックスの外で行動する方法を探す」など）。
同社は、NanoGPT スピードラン ベンチマークに関する以前のテストについて説明しました。このテストでは、モデルは結果を社内 Slack に投稿することのみを指示されていました。代わりに、モデルはベンチマーク手法の矛盾する指令に従うことを決定しました。

結果を GitHub に公開するためのアクション。
OpenAIは、その指示に従おうとして、このモデルはインターネットへの外部アクセスをブロックすることを目的としたサンドボックスの制限を回避する方法を探すのに1時間を費やしたと述べた。 OpenAIによると、ロングホライズンモデルが示すこの種の「持続性」は、同様のジレンマに直面したときに諦めたり、ユーザーに説明を求めたりする傾向があった初期のモデルの動作とは異なるという。
このようなインシデントを受けて、OpenAI は、エージェントの個々の動きではなく、アクションの全軌跡を追跡するように設計された「アクティブ モニタリング」システムなど、長期モデルに焦点を当てた新しい安全策を導入するようになりました。しかし、OpenAIは、ベンチマークが「サイバー脆弱性をテストすることを目的としていた」ため、顔ハグ事件の際にはこの種の安全策が「意図的に有効にされなかった」と述べている。
AI の封じ込めに関する SF の比喩に慣れている人なら誰でも、AI モデルがセキュリティの脆弱性を悪用して、その目標を達成するために禁止されているインターネット アクセスを獲得するという顕著な現実世界の例を目にすると、少なくとも少しは動揺するかもしれません。グレッグ・カサール下院議員（民主党、テキサス州）はソーシャルメディア声明でこの事件を「極めて憂慮すべきこと」と呼び、「定期的な独立した安全性試験と監督の義務付け、セキュリティー事件の強制的開示、絶対的な災害から人々を守るための国際協力」を求めた。
「顔ハグ事件」はまた、いわゆる AI の調整や、AI モデルの動作が人間の作成者の意図と一致していることを確認するための継続的な取り組みに関する哲学的および実践的な議論の注目度を高めました。 OpenAIは、今週初めのセキュリティブログ投稿で、ロングホライズンモデルが「ロングホライズンの指示を確実に記憶する」ための措置を講じたと述べた。

これは、テストにおける「不一致」結果の数を大幅に減らすのに役立ちました。
OpenAIの安全研究者のミカ・キャロル氏は、この事件についてソーシャルメディアに「位置ずれのリスクが今後の主要な懸念事項になるということが納得できないとしても、何が起こるかは分からない」とソーシャルメディアに書いた。
AI モデルがベンチマークを通過する意図しない方法を見つけるために多大な労力を費やしたのは、これが初めてではありません。今週発表された報告書の中で、英国の AI セキュリティ研究所は、サイバー評価で「不正行為」を試みる（つまり、ショートカット、回避策、または意図しない/許可されていない方法を使用して解決策を見つける）最近のモデルが 8 ～ 14% の確率で検出されたと指摘しました。これは、検出されなかった不正行為の試みの一部を過小評価する可能性がある下限の範囲です。
セキュリティ テスト グループは、構成が正しくなく「解決不可能」な評価に直面したモデルが、自身が作成し、監視されていないサードパーティのインターネット サービス上でホストされているコードを使用して AISI 独自の評価インフラストラクチャにアクセスしようとしたという、あるインシデントについて説明しました。
また、Hugging Face への侵入は、AI 企業が最新モデルのサイバー攻撃能力について重大な警告を発しており、各国政府が国家安全保障を重視してその展開を制限する命令で対応している時期に発生しました。一部の懐疑論者は、この種の発言を最新モデルの機能に対する誇大宣伝だとみなしていますが、独立した評価では、最近のモデルが以前の自律システムでは不可能だった侵入目標を達成していることが示されています。
OpenAIのサム・アルトマン氏は4月のインタビューで、パニックに陥ったAIセキュリティ警告を「恐怖に基づくマーケティング」だと批判した。しかし 6 月、OpenAI は米国政府の安全性への懸念に応え、GPT-5.6 のリリースを延期しました。
これらが議論する中で、

AI とサイバーセキュリティの分野で展開されているように、「ハグフェイス」事件は、サイバーセキュリティの専門家が AI ベースの脅威にどのようにアプローチするかについてのターニングポイントとみなされるようになる可能性があります。 「自律型のAI主導の攻撃ツールは、もはや理論上のものではない」とハギング・フェイスは先週の開示文書で書いた。 「これにより、広範で忍耐強い、多段階のキャンペーンを実行するコストが削減され、マシンの速度で動作します。オンライン プラットフォームを防御するということは、データとモデルの表面を第一級の攻撃面として扱い、防御に AI を使用して歩調を合わせることを意味します。」
「これはエージェント時代のサイバーセキュリティの初日です」とハギングフェイスの共同創設者兼最高経営責任者（CEO）のクレム・デラング氏は本日ソーシャルメディアに書いた。 「私たちは皆、秘密主義が解決策ではなく、どこにいても（一部の選ばれた防御者だけでなく）すべての防御者が制限のない、より強力なモデル、特にオープンなモデルを必要としているということを学びつつあります。」
147件のコメント
コメント
フォーラムビュー
コメントを読み込んでいます...
前の話
次の話
よく読まれている
1.
20年間逃走中の最重要指名手配逃亡者、バイオテクノロジー企業幹部として隠れて逮捕される
2.
任天堂、ユーザーは自発的に高い価格を支払った、関税を払い戻す権利はないと主張
3.
米軍が供給を使い果たしているため、無制限の AI トークンは結局のところ無制限ではない
4.
レンジローバーは、「SUV ではないものを作ったらどうなるでしょうか?」という質問に答えます。
5.
光子を半分に切り刻もうとすると何が起こるでしょうか?
カスタマイズ
Ars Technica は信号を分離してきました。
25年以上続く騒音。弊社独自の組み合わせにより、
技術的な知識と技術芸術への幅広い関心
Ars は、情報の海の中で信頼できる情報源です。後
すべてを知る必要はありません。重要なことだけを知っておく必要があります。

## Original Extract

This is day one for cybersecurity in the age of agents," Hugging Face CEO says.

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
I want to break free
OpenAI says its AI agent broke out of testing sandbox to hack Hugging Face
“This is day one for cybersecurity in the age of agents,” Hugging Face CEO says.
147
Let me out of here, I have to pass this benchmark!
Credit:
Getty Images
Let me out of here, I have to pass this benchmark!
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
OpenAI says an agent powered by its LLM models escaped its sandboxed testing environment to infiltrate Hugging Face’s servers as part of an overzealous attempt to obtain solutions to a benchmark test. The company says it considers the unintended infiltration an “an unprecedented cyber incident” and is working with Hugging Face on new protections to prevent a recurrence.
Hugging Face disclosed an intrusion last week that it said involved “unauthorized access to a limited set of internal datasets and to several credentials used by our services.” The AI data clearinghouse said it used its own LLM-driven analysis to identify “a swarm of tens of thousands of automated actions” from an “autonomous agent framework.” That agentic swarm exploited a flaw in Hugging Face’s data-processing pipeline to gain the ability to run code as a processing worker, eventually escalating to high-level access to the company’s cloud and server clusters.
At the time, Hugging Face said the LLM being used in the attack was “still not known.” But OpenAI took responsibility for the intrusion Tuesday evening , saying it came about during an internal test involving the recently released GPT-5.6 Sol and “an even more capable pre-release model.” The models were being tested against the ExploitGym benchmark , an independent testing suite based on hundreds of real-world security vulnerabilities.
Although OpenAI said the benchmark tests were running in “a highly isolated environment,” its agents still had access to “internally hosted third-party software” to install required package registries. While working on the ExploitGym benchmark over a weekend, OpenAI says the agent “spent a substantial amount of inference compute finding a way to obtain open Internet access,” eventually locating one via a zero-day vulnerability in the package registry cache proxy.
With that wider Internet access established, OpenAI says the agent “inferred that Hugging Face potentially hosted models, datasets, and solutions for ExploitGym,” leading to the previously disclosed attack on the servers. OpenAI says its security team “discovered this anomalous activity internally,” independent of Hugging Face’s own detection.
Although OpenAI described the Hugging Face incident as “unprecedented,” the company also acknowledged in a separate blog post this week that it had previously observed recent “long-horizon models”—those that can operate autonomously for extended periods—taking “unwanted actions” to achieve testing goals, including “looking for ways to act outside its sandbox.”
The company described an earlier test on the NanoGPT speedrun benchmark, where the model was instructed only to post its results to an internal Slack. Instead, the model decided to follow a conflicting directive in the benchmark instructions to post the results publicly to GitHub.
In attempting to follow that instruction, OpenAI said the model spent an hour searching for ways to circumvent sandbox restrictions intended to block external access to the Internet. That kind of “persistence” demonstrated by long-horizon models differs from the actions of earlier models, which would tend to give up or seek user clarification when faced with similar dilemmas, OpenAI said.
Incidents like these have led OpenAI to deploy new safeguards focused on long-horizon models, including an “active monitoring” system designed to track the full trajectory of an agent’s actions rather than individual moves. But OpenAI says these kinds of safeguards “were intentionally not enabled” during the Hugging Face incident because the benchmark was “aimed at testing cyber vulnerabilities.”
Anyone familiar with science fiction tropes about AI containment might be at least a little disturbed by a prominent real-world example of an AI model exploiting security vulnerabilities to gain prohibited Internet access in pursuit of its goals. Congressman Greg Casar (D-Texas) called the incident “extremely alarming” in a social media statement and called for “regular mandatory independent safety testing and oversight, mandatory disclosure of security incidents, and international cooperation to keep people safe from absolute disaster.”
The Hugging Face incident has also heightened the salience of philosophical and practical debates over so-called AI alignment and the ongoing efforts to ensure that an AI model’s actions align with the intentions of its human creators. In its security blog post earlier this week, OpenAI said it had taken steps to ensure that long-horizon models are “remembering instructions on long rollouts,” which has helped severely reduce the number of “misaligned” outcomes in testing.
“If this doesn’t convince you that misalignment risks are going to be a key concern going forward, I don’t know what will,” OpenAI Safety Researcher Micah Carroll wrote on social media regarding the incident.
This is far from the first time an AI model has gone to great lengths to find unintended ways of passing a benchmark. In a report released this week , the UK’s AI Security Institute noted that it detected recent models attempting to “cheat” at its cyber evaluations (i.e., using shortcuts, workarounds, or unintended/disallowed methods to find a solution) between 8 and 14 percent of the time—a lower-bound range that could undercount some undetected cheating attempts.
The security testing group described one incident in which a model, faced with a misconfigured and “impossible to solve” evaluation, attempted to access AISI’s own evaluation infrastructure using code it wrote and hosted on an unmonitored third-party Internet service.
The Hugging Face infiltration also comes at a moment when AI companies are issuing grave warnings about the cyberattack capabilities of their latest models, leading governments to respond with national security-focused orders limiting their rollout. While some skeptics see these kinds of statements as hype-filled marketing for the capabilities of their latest models, independent evaluations show recent models achieving infiltration goals that were impossible for earlier autonomous systems.
OpenAI’s Sam Altman criticized panicked AI security warnings as “fear-based marketing” in an April interview . But in June, OpenAI delayed the release of GPT-5.6 in response to safety concerns from the US government .
As these debates play out in the AI and cybersecurity spheres, the Hugging Face incident could come to be seen as a turning point in how cybersecurity professionals approach AI-based threats. “Autonomous, AI-driven offensive tooling is no longer theoretical,” Hugging Face wrote in its disclosure last week. “It lowers the cost of running a broad, patient, multi-stage campaign, and it operates at machine speed. Defending an online platform now means treating the data and model surface as a first-class attack surface and using AI on defense to keep pace.”
“This is day one for cybersecurity in the age of agents,” Hugging Face co-founder and CEO Clem Delangue wrote on social media today . “We’re all learning that secrecy is not the answer and that all defenders (not just a few selected ones) everywhere need more powerful models without restrictions, especially open ones!”
147 Comments
Comments
Forum view
Loading comments...
Prev story
Next story
Most Read
1.
On the run for 20 years, most-wanted fugitive caught hiding as a biotech exec
2.
Nintendo says users voluntarily paid higher prices, have no right to tariff refunds
3.
Unlimited AI tokens aren't unlimited after all as US Army burns through supply
4.
Range Rover answers the question: "What if we built a not-SUV?"
5.
What happens when you try to chop a photon in half?
Customize
Ars Technica has been separating the signal from
the noise for over 25 years. With our unique combination of
technical savvy and wide-ranging interest in the technological arts
and sciences, Ars is the trusted source in a sea of information. After
all, you don’t need to know everything, only what’s important.
