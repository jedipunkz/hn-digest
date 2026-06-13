---
source: "https://www.theregister.com/offbeat/2026/06/13/world-cup-ai-predictor-now-lets-users-ask-daft-what-ifs/5254853"
hn_url: "https://news.ycombinator.com/item?id=48517421"
title: "World Cup AI predictor now lets users ask daft what-ifs"
article_title: "World Cup AI predictor now lets users ask daft what-ifs"
author: "kagaherk"
captured_at: "2026-06-13T15:38:21Z"
capture_tool: "hn-digest"
hn_id: 48517421
score: 2
comments: 0
posted_at: "2026-06-13T13:59:29Z"
tags:
  - hacker-news
  - translated
---

# World Cup AI predictor now lets users ask daft what-ifs

- HN: [48517421](https://news.ycombinator.com/item?id=48517421)
- Source: [www.theregister.com](https://www.theregister.com/offbeat/2026/06/13/world-cup-ai-predictor-now-lets-users-ask-daft-what-ifs/5254853)
- Score: 2
- Comments: 0
- Posted: 2026-06-13T13:59:29Z

## Translation

タイトル: ワールドカップ AI 予測ツールにより、ユーザーは愚かな「もしも」を質問できるようになりました
説明: ネタバレ: ありません

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
ワールドカップ AI 予測ツールにより、ユーザーは愚かな「もしも」を質問できるようになりました
ネタバレ：Team Registerにとっては良い結果ではありませんでした
AI Octopus Euro 2024 予測ツールの背後にあるチームは、2026 FIFA ワールドカップに向けてシミュレーターを更新しました。今回は、ユーザーがモデルに自然言語シナリオを投げて、トーナメントがどうなるかを確認できるようになりました。
「レッドカード、重大な怪我、熱波、チームのベースキャンプ変更など、賢明な質問は有効ですが、例えば『トーナメントがラグビールールで行われたらどうなるか？』などの愚かな質問も同様です」とLuzmoのCTO兼共同創設者のハロエン・ヴァーミレン氏は述べた。
システムはシンプルです。プロンプト ボックスにシナリオを入力すると、予測機能が結果がどうなるかを吐き出します。生データには、選手情報、暑さと高度の要因、怪我のデータなどに基づいたチームの質が含まれます。トーナメントのモンテカルロ シミュレーションを使用して勝ち/負け/引き分けの確率が生成され、スコアラインは 5,000 回の試合実行から導出されます。
Euro 2024 AI Octopus の背後にあるエンジンは TypeScript で書かれています。今回、チームは Rust を使用しました。 「リアルタイム コンポーネントが追加されたため、より迅速に実行できるようにするために Rust に移行しました」と Vermylen 氏は The Register に語った。
「以前は 5 分程度実行できました。現在は、実際のシミュレーション時間の 2 ～ 3 秒以内に予測が実際に得られるようにしたいと考えています。」
OpenAI モデルはリクエストを解析して概要を生成し、エージェントはシナリオの作成または変換、計算エンジンの呼び出し、質問への回答などに使用されます。ユーザーは、質問して答えを理解するためにデータ サイエンティストである必要はありません。
結果に基づいて再計算するので確かに速いです

提案されたシナリオ（特定の世界指導者による政治的に疑わしい排出量の影響を熟考したシナリオでさえも）について。すべてのシナリオがうまくいくわけではありません。ヴァーミレン氏は、フィルタリングは冒涜的な表現を無視し、「特定のグループにとって有害なだけのシナリオを避けるために」設置されていると語った。
AI Octopus がユーロ 2024 の結果を予測: イングランドにとっては良くないようだ
Rust が IBM メインフレームに侵入するが、発生するのは夜間のみ
AFCアヤックス、欠陥によりハッカーがチケットや出場停止の管理を可能にしてボールをドロップ
物議を醸した幻覚症状の後、警察はMicrosoft Copilotを留置場に留置
さらに、AI パーサーが単にプロンプトを理解できないという長年の問題もあります。明瞭さが鍵です。自然言語の使用は、設定やスライダーを備えた UI に代わる優れた代替手段ですが、その使いやすさが誤解を招く可能性があります。
トーナメントが進むにつれて、データは洗練されていきます。この記事の執筆時点では、ベースラインではスペインが決勝でイングランドに勝つと予想されています。スペインがトロフィーを掲げる可能性は現在18％、決勝進出の可能性は26.8％となっている。もちろん、これらの数値はシナリオを入力することで変更できます。
たとえば、「スペインチームがまずいパエリアを食べたらどうしますか？」と尋ねました。その後、スペインがトーナメントで優勝する可能性は 1.5% に下がり、フランスが優勝すると予想されました。
また、イングランドチームをレジスターライターに置き換えたらどうなるかについても尋ねました。このシナリオはうまく終わらなかったと言うだけで十分だろう。
私たちはヴァーミレンに次は何なのか尋ねました。 「オリンピックもいいですね…あるいはユーロビジョンもいいですね。我々はイギリスに勝利を与えたいと思っています。」 ®
オフビート
ワールドカップ AI 予測ツールにより、ユーザーは愚かな「もしも」を質問できるようになりました
ネタバレ：Team Registerにとっては良い結果ではありませんでした
AWS はより高速で効率的なネットワーキングを目指してサイコロを振る
ハニー、データセンター ネットワークをフラット化しました
Z

TE、AI を活用したネットワーク イノベーションにより、Selular Award 2026 で 3 つの栄誉を獲得
パートナーコンテンツ: FWA、ネットワークエコシステム、ネイティブAIベースバンドにおける画期的な成果が認められ、ZTEはインドネシアの5GアドバンストとAIの経済成長の主要な推進者としての役割を強固にする
NHS 患者は Palantir のデータ プラットフォームをオプトアウトできませんが、病院はオプトアウトできます
大臣、議会が2027年2月のFDP契約更新を検討中、調達に関しては信託単独で対応できると発言
PAAS と IAAS
Graviton 5 は感動的ですが、すべての神聖なるものへの愛を込めて、これを「AI チップ」と呼ぶのはやめてください
AWS は口よりもチップ工場の運営のほうが上手い
XP時代のWindowsがロンドンの無人鉄道に出没しているのを発見
過去の爆風が通勤客を出迎えます
セキュリティ
GitHub がワーム感染の疑いで 70 以上の Microsoft リポジトリを破壊し、CI/CD パイプラインを破壊
公共部門
GOV.UK が Stripe を売却し、支払いをオランダに変更
セキュリティ
Microsoft ビーフを使った怒っているバグハンターが新しい Windows 0-day をドロップ
セキュリティ
シグナル、英国がヌード画像のデバイスをスキャンする計画は「私たち全員を危険にさらす」と主張
セキュリティ
すべての従業員のパスワードは 1 つの Excel ファイルに保存されていました
ボラティリティを乗り越えて成功する: 不確実な市場における永遠の優位性
消費ベースの運用モデルがどのように柔軟性を提供し、効率を向上させ、インフラストラクチャ投資に予測可能性をもたらすかを学びましょう。
プロンプトからエクスプロイトまで: LLM は API 攻撃をどのように変化させているか
最新のアプリケーションは API 主導で相互接続されており、多くの場合、過剰な権限が与えられているため、AI 支援攻撃の理想的な標的となっています。
未来の構築: Kubernetes 向けエンタープライズ データ サービスのロックを解除する
インフラストラクチャのサイロを排除し、標準化されたエンタープライズ グレードのクラウド ネイティブ プラットフォームを確立する方法を発見してください。
Microsoft 365 が見逃す高度な攻撃を行動 AI Securit で捕捉

y
Microsoft 365 は企業コミュニケーションのバックボーンであり、そのネイティブ セキュリティが既知のものや騒々しいものを除外します。
これは、回復力のある次世代の開発および IT 運用を定義する実用的なツールとテクニックを技術的に深く掘り下げるものです。
ランサムウェア侵害の混乱に足を踏み入れ、対応スキルをテストし、他の IT およびセキュリティの専門家とチームを組んでサイバー犯罪者を出し抜こう
仮想サイバー復旧シミュレーション
ランサムウェア攻撃の勢いは衰えていませんが、私たちも同様です。 Druva の人気イベント「Escape Ransomware」が完全にバーチャルになりました。
Agentic AI 時代のゼロトラスト
ほとんどの組織が依存している ID およびアクセス モデルは、人間以外の ID が独立して動作するのではなく、人間のユーザー向けに構築されています。
Agentic AI 時代のゼロトラスト
ほとんどの組織が依存している ID およびアクセス モデルは、人間以外の ID が独立して動作するのではなく、人間のユーザー向けに構築されています。
大規模なエージェント AI: パイロットから運用まで
AI の導入を大規模に推進することで、真の ROI を実現する方法を学びましょう。
AIエージェントは信用できないので、危険な権限を与えないでください
SKグループ会長「我々は可能な限り早く進んでいる」
GitHub は、普及されている AI を実際に使用している顧客に不意を突かれた
GPTZeroは、報告書の45件の引用のうち出典と一致したのは5件のみであると主張しており、四大AI研究がどのようにまとめられたのか疑問が生じている
電報ベースの「アウトサイダー・エンタープライズ」が数百万件の詐欺テキストを送信し、信頼できるブランドになりすましたとして告訴される
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
行ってきました

企業全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
CentOS の歴史: 生化学者の Linux 趣味プロジェクトがどのようにして企業世界のデフォルトのオペレーティング システムになったのか
Red Hat が Windows が「おそらく適切な製品」であると述べた後、コミュニティが団結したとき
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
Project Headroom は多額の費用も節約できる
OpenBSD 7.9 が登場、あらゆる鋭いエッジを誇るダイヤモンドの原石
60 番目のリリースでは、その禁欲的な努力を失うことなく、より多くのコア、遅延休止状態、および基本的な Wi-Fi 6 が追加されています
Fedora: Microsoft は全力で取り組んでいるが、Deepin は見捨てられている
Red Hat の無料ディストリビューションはデスクトップを失うが、重要な新しい友人を作る
LocalSend はスニーカーネットを廃業にします
Apple ロックインを除いた AirDrop と同様
dBase の劣化: データベースの巨人が 47 年後に黒に退色
ブログ投稿の喪の減少が、このベテランアプリのオンラインでの存在感をオフラインに落とす一因となったようだ
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

Spoiler: It doesn

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
World Cup AI predictor now lets users ask daft what-ifs
Spoiler: It doesn't end well for Team Register
The team behind the AI Octopus Euro 2024 predictor has updated its simulator for the 2026 FIFA World Cup, this time allowing users to throw natural-language scenarios at the model and see how the tournament might shake out.
"Sensible questions work – a red card, a key injury, a heat wave, a squad switching base camp – but so do the daft ones, e.g. 'What if the tournament were played with rugby rules?'" said Luzmo CTO and co-founder Haroen Vermylen.
The system is simple: enter a scenario in a prompt box, and the predictor spits out how the results might go. The raw data includes squad quality based on player information, heat and altitude factors, injury data, and so on. A Monte Carlo simulation of the tournament is used to generate win/lose/draw probabilities, and the score line is derived from 5,000 match runs.
The engine behind the Euro 2024 AI Octopus was written in TypeScript. This time around, the team used Rust. "We moved to Rust to also be able to run things more quickly, as now there is a real-time component to this," Vermylen told The Register .
"Before it could run for five minutes or so. Now we want the predictions to actually come out within two to three seconds of actual simulation time."
OpenAI models parse the request and generate summaries, and an agent is used to create or transform scenarios, call the calculation engine, answer questions, and so on. A user doesn't need to be a data scientist to ask questions and understand the answers.
It's certainly rapid, recalculating the results based on suggested scenarios (even one in which we pondered the effect of politically dubious emissions from a certain world leader). Not that all scenarios will work. Vermylen told us that filtering was in place to ignore profanities and "to avoid scenarios that would just be harmful to certain groups."
AI Octopus predicts results of Euro 2024: It isn't looking good for England
Rust stalks IBM mainframes, but only in nightly form
AFC Ajax drops ball as flaws let hackers play admin with tickets and bans
Cops put Microsoft Copilot in holding cell after controversial hallucination
And then there is the age-old issue of an AI parser simply not understanding the prompt. Clarity is key. Using natural language is a great alternative to a UI with settings and sliders, but that ease of use can result in misunderstandings.
As the tournament progresses, the data will be refined. At the time of writing, the baseline reckons that Spain will beat England in the final. Spain currently has an 18 percent chance of lifting the trophy and a 26.8 percent chance of reaching the finals. Those figures can, of course, be altered by feeding in scenarios.
For example, we asked: "What if the Spanish team eats a bad paella?" Spain's chance of winning the tournament then dropped to 1.5 percent, with France as the projected champion.
We also asked it what would happen if we replaced the England team with Register writers. Suffice to say that scenario did not end well.
We asked Vermylen what was next. "The Olympics would be nice… or the Eurovision. We'd like to give the United Kingdom a win." ®
OFFBEAT
World Cup AI predictor now lets users ask daft what-ifs
Spoiler: It doesn't end well for Team Register
AWS rolls the dice for faster, more efficient networking
Honey, I flattened the datacenter network
ZTE wins three Selular Award 2026 honors for AI-powered network innovation
PARTNER CONTENT: Recognized for breakthrough achievements in FWA, Network Ecosystem, and Native AI Baseband, ZTE solidifies its role as a key driver of Indonesia’s 5G-Advanced and AI economic growth
NHS patients can't opt out of Palantir's data platform – but their hospital can
Minister says trusts can go it alone on procurement as Parliament mulls February 2027 FDP contract renewal
PAAS AND IAAS
Graviton 5 impresses, but please, for the love of all that's holy, stop calling them 'AI chips'
AWS better at running chip fabs than their mouths
XP-era Windows spotted haunting London's driverless railway
A blast from the past greets commuters
security
GitHub nukes 70+ Microsoft repos, breaks CI/CD pipelines, following suspected worm infections
public sector
GOV.UK goes Dutch on payments as it dumps Stripe
Security
Angry bug hunter with Microsoft beef drops new Windows 0-day
Security
Signal says UK plan to scan devices for nude images 'endangers us all'
SECURITY
Every employee’s password was stored in a single Excel file
Thriving Through Volatility: The Everpure Advantage in an Uncertain Market
Learn how a consumption-based operating model provides flexibility, improves efficiency, and brings predictability to infrastructure investments.
From Prompt to Exploit: How LLMs Are Changing API Attacks
Modern applications are API-driven, interconnected, and often over-permissioned, making them an ideal target for AI-assisted attacks.
Architecting the Future: Unlocking Enterprise Data Services for Kubernetes
Join us to discover how to eliminate infrastructure silos and establish a standardized, enterprise-grade cloud-native platform.
Catch the Advanced Attacks Microsoft 365 Misses with Behavioral AI Security
Microsoft 365 is the backbone of enterprise communication, and its native security filters out the known and the noisy.
This is your technical deep-dive into the practical tools and techniques that define the next generation of resilient Dev and IT operations.
Step into the chaos of a live ransomware breach, test your response skills, and team up with other IT and security pros to outsmart cybercriminals
Virtual Cyber Recovery Simulation
Ransomware attacks aren’t slowing down, and neither are we. Druva’s hit event, Escape Ransomware, is now fully virtual.
Zero Trust for the Agentic AI Era
The identity and access models most organizations rely on were built for human users, not non-human identities operating independently.
Zero Trust for the Agentic AI Era
The identity and access models most organizations rely on were built for human users, not non-human identities operating independently.
Agentic AI at Scale: From Pilot to Production
Join us to learn how to unlock real ROI by driving adoption of AI at scale.
AI agents can't be trusted, so don't give them dangerous powers
We're moving as fast as we can, says SK Group chair
GitHub caught off guard by customers actually using the AI being evangelized
GPTZero claims only 5 of the report's 45 citations matched their sources, raising questions about how the Big Four's AI study was assembled
Telegram-based 'Outsider Enterprise' accused of sending millions of scam texts and impersonating trusted brands
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
History of CentOS: How a biochemist's Linux hobby project became the enterprise world's default operating system
When a community came together after Red Hat said Windows was 'probably the right product'
Netflix wiz creates app to slash AI bills, then open sources it
Project Headroom could save you big money, too
OpenBSD 7.9 arrives, a diamond in the rough proud of every sharp edge
Sixtieth release adds more cores, delayed hibernation, and basic Wi-Fi 6 without losing its ascetic streak
Fedora: Microsoft is all aboard, but Deepin is dumped
Red Hat’s free distro loses a desktop, but makes an important new friend
LocalSend puts your sneakernet out of business
Like AirDrop, minus the Apple lock-in
dBase debased: Database titan fades to black after 47 years
Blog post mourning decline appears to have helped knock what was left of the veteran app's online presence offline
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
