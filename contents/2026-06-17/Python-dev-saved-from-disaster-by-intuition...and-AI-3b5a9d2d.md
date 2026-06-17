---
source: "https://www.theregister.com/ai-and-ml/2026/06/16/python-dev-saved-from-disaster-by-intuitionand-ai/5256632"
hn_url: "https://news.ycombinator.com/item?id=48562954"
title: "Python dev saved from disaster by intuition...and AI"
article_title: "Python dev saved from disaster by intuition...and AI"
author: "Bender"
captured_at: "2026-06-17T01:03:18Z"
capture_tool: "hn-digest"
hn_id: 48562954
score: 2
comments: 1
posted_at: "2026-06-16T22:15:26Z"
tags:
  - hacker-news
  - translated
---

# Python dev saved from disaster by intuition...and AI

- HN: [48562954](https://news.ycombinator.com/item?id=48562954)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/06/16/python-dev-saved-from-disaster-by-intuitionand-ai/5256632)
- Score: 2
- Comments: 1
- Posted: 2026-06-16T22:15:26Z

## Translation

タイトル: 直感と AI によって災害から救われた Python 開発者
説明: 私

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
Python 開発者は直感と AI によって惨事から救われた
ごめんなさい、デイブ。システムを完全に修復するようなリポジトリをインストールすることはできません。
Python 開発者の Roman Imankulov 氏は危うく餌に乗ろうとしました。彼がそうしなかったという事実は、人間の直感と AI のコード精査のおかげであると考えられます。
小規模な仮想通貨スタートアップの採用担当者を名乗る人物が LinkedIn を通じて連絡し、彼女が言うところの機能しない概念実証コードに関する助けを求めました。彼女の説明によると、会社には主任エンジニアが必要だったという。
イマンクロフ氏がこのやりとりをブログ投稿で説明したように、採用担当者は彼に、非推奨の Node モジュールの問題を調査するよう求めました。リクエストに関して何かが間違っているように思えました。
イマンクロフ氏は電話インタビューで、「おそらく私たち全員がそうであるように、その種の攻撃については聞いていた」と説明した。 「それで、『もしこれが自分がターゲットになったらどうしよう』って思ったんです」それは私の過去の経験に基づいているだけです。」
そこで彼は、Hetzner 上で VPS をスピンアップし、リポジトリのクローンを作成するという珍しい手順を踏みました。次に、Pi コーディング エージェント (Codex を実行) を使用して、コードの読み取り専用分析を実行しました。
「エージェントを実行して動作をテストしたところ、『すべてが明確で、コードは醜いですが、一般的には実行しても安全です。そのままレビューを実行してください』という返答が返ってくるとほぼ確信していました」と同氏は説明した。 「驚いたことに、エージェントはすぐに『このコードは実行しないでください。罠があるのでその場を立ち去ってください』というような応答を返しました。」
AI とブレイン コンピューター インターフェイスにより、言葉を話せない ALS 患者がフルタイムで働けるようになります
未知の攻撃者によってばら撒かれたフォーティネット サンドボックスの 3 つの重大なバグ
Commodore が Sailfish で電話事業に参入 -

パワードレトロ「コールバック」
HPE、エージェントワークロード向けの AI インフラストラクチャポートフォリオを強化
AI モデルは、ファイルの 1 つである app/test/index.js にフラグを立てていました。
ファイルにはバックドアが含まれていました。これは、テスト スイート構成のように断片化されたサーバー URL と、サーバーがリクエストに応じて送信するものを実行するネットワーク リクエストの形式をとりました。
イマンクロフ氏は、AI エージェントが見逃していた詳細をキャッチしてくれたと信じています。
「私はこのコードを自分で開いてざっと読んだのですが、私にはそれは、ずさんな開発者によって書かれた、ただのずさんなファイルにしか見えませんでした」と彼は言いました。 「だから私はただ下にスクロールして、『そうだ、そうだ、これはひどいことだけど、このコードを修正するためにお金を払ってもらえるなら、私は気にしない』と考えた。」しかし、まったく同じファイル内のエージェントが、私が見落としていたまさに脆弱性を発見しました。」
npm を使用してリポジトリをインストールするだけでも、バックドアをトリガーするには十分です。リポジトリの package.json ファイルには、インストール プロセスの後にスクリプトを実行するように設計された「 prepare 」インストール後フックが含まれていました。
参照された悪意のあるリポジトリはアクセスできなくなりました – おそらく、Imankulov の苦情に応じて GitHub が削除しました – しかし、クローンはまだ見つかります。
「この攻撃が潜行的なのは、標準的な開発者のワークフローをどのように乗っ取るかです」と、独立系オープンソース兼セキュリティアーキテクトの Devashri Datta 氏は The Register への電子メールで説明しました。 「攻撃者は、ターゲットが疑わしいバイナリを実行することに依存せず、ターゲットが日常的なコマンド npm install を実行することに依存していました。
「package.json 内の準備ライフサイクル フック内に実行ロジックを埋め込むことで、依存関係の解決中に悪意のあるペイロードが自動的にトリガーされます。これは新しい手法ではありませんが、開発者が autop で npm install を実行するため、非常に効果的です。

アイロット。小さな定数からドメインをつなぎ合わせて悪意のある URL を組み立てるために使用された文字列の断片化は、ハードコードされた侵害の兆候をスキャンする静的分析ツールを無効にするために設計された意図的な難読化でした。」
イマンクロフ氏は、悪意のあるリポジトリ内のコミットは、Web での存在感と一連の作業を確立した開発者の作品であるようだと述べた。しかし、作者とされる人物に連絡したところ、開発者はGitHub上で複数回なりすまされており、そのコードは書いていないと述べた。
採用担当者のLinkedInプロフィールには本物の芸術ジャーナリストが言及されていたが、イマンクロフ氏は関連するプロフィールが偽造されたものであると信じている。採用担当者とのオンラインでのやりとりから、彼女の職歴では明らかではないレベルの技術的知識が示唆されました。
LinkedIn は、誰かとやり取りする前に捕まえて削除する何千万もの偽アカウントについて話すことを好みます。しかし、検出されてフラグが立てられる前に、依然として何十万ものアカウントが作成され、人々とやり取りされています。そしてその数は増え続けています。 2025 年 1 月から 6 月までの期間、LinkedIn はユーザーからの報告を受けて 386,000 のアカウントを制限しました。この数字は、前の 6 か月間に 266,000 件でした。そして、2021年1月から6月までの期間では、わずか8万6,000件でした。
この種のソフトウェア サプライ チェーンのソーシャル エンジニアリング攻撃は一般的になりました。今月初め、北朝鮮に関連した詐欺師が偽の面接や求人情報を利用して開発者アカウントを侵害するさまざまなキャンペーンを実行していることに注目しました。
他の開発者は、これらの詐欺に陥りそうになった（そして AI エージェントによって救われた）と報告し、コード分析を投稿しました。
ダッタ氏は、イマンクロフ氏の返答は、セキュリティを意識する開発者によるコードレビューの衛生管理への取り組み方の変化を浮き彫りにしていると述べた。
「歴史的に、ガイダンスは

信頼できないコードをサンドボックス化するか、手動でレビューすることでした。」と彼女は言いました。「ここで、Roman は制約のある読み取り専用環境にローカル AI エージェントを導入し、何かを実行する前にコードベースを分析しました。これは、攻撃的な脅威ベクトルとしての AI に関する支配的な説に対する有用な対案です。 AI エージェントは開発者エンドポイントで防御的に使用されるため、疲労や社会的プレッシャーの影響を受けません。テスト スイートが未検証のコードを取得するためにアウトバウンド ネットワーク接続を開始するなど、異常な動作が数秒で表面化するだけです。」
少しでも慰めになるなら、来月には関連する攻撃ベクトルに対処する必要があります。
npm を管理している GitHub は、npm install コマンドの動作を変更する npm 12 のリリースを準備しています。 allowScripts 設定はデフォルトでオフになります。 「プロジェクトで明示的に許可されていない限り、npm install は依存関係からプレインストール、インストール、ポストインストールのスクリプトを実行しなくなります」と GitHub は説明しています。
「インストール時のライフサイクル スクリプトは、npm エコシステムにおける単一で最大のコード実行面です」と、GitHub プロダクト マネージャーの Leo Balter 氏が先週のコミュニティ ディスカッション投稿で説明しました。 「すべての npm インストールでは、あらゆる推移的な依存関係からスクリプトが実行されるため、ツリー内のどこかにある単一の侵害されたパッケージによって、開発者マシンまたは CI ランナー上で任意のコードが実行される可能性があります。スクリプトの実行をオプトインにすると、そのパスが閉じられますが、信頼できるパッケージへのパスはコマンド 1 つで保存されます。」
イマンクロフ氏は、それについて強い意見はないと述べた。 「私の観点から言えば、個人の安全を確保するために、これらのスクリプトがデフォルトで実行されないようにするために pnpm に切り替えました」と彼は言いました。
ダッタ氏は、今回の事件は、エンタープライズソフトウェアサプライチェーンのセキュリティが企業ネットワークの境界を超えて拡張されなければならない理由を浮き彫りにしたと述べた。
「で

「攻撃者は現在、単一のコード行が企業のサプライチェーンに入る前に、個々のエンジニアリングエンドポイントに至るまで左にシフトしています。定期的な採用面接と思われる間に開発者のローカルワークステーションが侵害されると、そのマシンにはアクティブなSSHキー、クラウドプロバイダートークン、内部リポジトリへのライブアクセスが頻繁に保持されます。」と彼女は述べた。
Datta 氏は、適切な防御には、サードパーティ コードや信頼できないコードを評価するための、分離された開発者コンテナや安全なクラウド ワークステーションなどの技術的なガードレールを適用する必要があると主張しています。
「新たなフレームワークは、導入時点で脅威を阻止するためには、VEX スタイルの信号がエンタープライズ SBOM インベントリよりも左に伝わる必要があることを認識し、悪用可能性のコンテキストをワークステーション層自体まで拡張し始めています」と彼女は述べました。 ®
システム
AMDのMext買収は、AIが引き起こしたRAM不足をAIがどのように解決できるかを示している
メモリが不足しているので、メモリを増やす余裕はありませんか? House of Zen の最新の買収により、フラッシュベースのメモリ拡張に AI が導入されました
新しい Siri は Apple の最も便利な OS 機能の 1 つを面倒なものにする
さようなら、便利なスポットライト。こんにちは、Apple のインテリジェンスを強制的に供給されたブロートウェアです。Google AI のような悲惨な感じです。
アルマトイで開催される ZTE Day 2026 では、カザフスタンのインテリジェントテレコムの未来を形作るイノベーションを紹介
パートナーコンテンツ: 次世代の接続性とスーパーコンピューティングソリューションでカザフスタンの「デジタル化とAIの年」を強化
Python 開発者は直感と AI によって惨事から救われた
ごめんなさい、デイブ。システムを完全に修復するようなリポジトリをインストールすることはできません。
PAAS と IAAS
Graviton 5 は感動的ですが、すべての神聖なるものへの愛を込めて、これを「AI チップ」と呼ぶのはやめてください
AWS は口よりもチップ工場の運営のほうが上手い
インテル生まれのネットワーク技術が Infi として再登場

DoE スーパーの niBand 代替品
Omni-Path がローレンス リバモア システムを 400 Gbps で起動
セキュリティ
研究者によると、脱獄ではなく単に「このコードを修正してください」というプロンプトが表示されただけで、連邦当局は寓話5に激怒したとのこと
オンプレミス
アマゾンは昨年自社のビット納屋で最大25億ガロンの水を使用している
セキュリティ
Microsoft ビーフを使った怒っているバグハンターが新しい Windows 0-day をドロップ
セキュリティ
シグナル、英国がヌード画像のデバイスをスキャンする計画は「私たち全員を危険にさらす」と主張
セキュリティ
GitHub がワーム感染の疑いで 70 以上の Microsoft リポジトリを破壊し、CI/CD パイプラインを破壊
プロンプトからエクスプロイトまで: LLM は API 攻撃をどのように変化させているか
最新のアプリケーションは API 主導で相互接続されており、多くの場合、過剰な権限が与えられているため、AI 支援攻撃の理想的な標的となっています。
未来の構築: Kubernetes 向けエンタープライズ データ サービスのロックを解除する
インフラストラクチャのサイロを排除し、標準化されたエンタープライズ グレードのクラウド ネイティブ プラットフォームを確立する方法を発見してください。
Microsoft 365 が見逃す高度な攻撃を Behavioral AI セキュリティで捕捉
Microsoft 365 は企業コミュニケーションのバックボーンであり、そのネイティブ セキュリティが既知のものや騒々しいものを除外します。
これは、回復力のある次世代の開発および IT 運用を定義する実用的なツールとテクニックを技術的に深く掘り下げるものです。
ランサムウェア侵害の混乱に足を踏み入れ、対応スキルをテストし、他の IT およびセキュリティの専門家とチームを組んでサイバー犯罪者を出し抜こう
仮想サイバー復旧シミュレーション
ランサムウェア攻撃の勢いは衰えていませんが、私たちも同様です。 Druva の人気イベント「Escape Ransomware」が完全にバーチャルになりました。
Agentic AI 時代のゼロトラスト
ほとんどの組織が依存している ID およびアクセス モデルは、人間以外の ID が独立して動作するのではなく、人間のユーザー向けに構築されています。
Agentic AI 時代のゼロトラスト
ID とアクセス モデル

ほとんどの組織が依存しているのは、人間以外のアイデンティティが独立して動作するのではなく、人間のユーザー向けに構築されたものです。
大規模なエージェント AI: パイロットから運用まで
AI の導入を大規模に推進することで、真の ROI を実現する方法を学びましょう。
さようなら、便利なスポットライト。こんにちは、Apple のインテリジェンスを強制的に供給されたブロートウェアです。Google AI のような悲惨な感じです。
このハードウェアは新しいものではありませんが、カリフォルニア大学デービス校の研究チームが開発した、ALS 患者の脳活動を 92% の精度で文章に変換する機械学習を利用した方法は、
AI エージェントは他のワークロードと何ら変わらない汎用ワークロードです
Copilot と AI に関する誤解を招く記述はありますか?きっとそうではありません！
AI システムの需要と DRAM と NAND の不足が世界市場を形成している
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
フランスのデジタルソー

[切り捨てられた]

## Original Extract

I

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Python dev saved from disaster by intuition...and AI
I'm sorry, Dave. I can't install that repo that will totally hose your system.
Python developer Roman Imankulov nearly took the bait. The fact that he didn't can be chalked up to human intuition and AI code vetting.
A person claiming to be a recruiter from a small crypto startup got in touch through LinkedIn, looking for help with what she described as proof-of-concept code that didn't work. The company, she explained, needed a lead engineer.
As Imankulov described the exchange in a blog post , the recruiter asked him to look into an issue with a deprecated Node module. Something about the request seemed off.
"I'd heard, as probably all of us have, about those types of attacks," Imankulov explained in a phone interview. "And I was like, 'what if this could be I could be the target?' It was just based on the past experience that I had."
So he took the unusual step of spinning up a VPS on Hetzner where he cloned the repo. He then used his Pi coding agent (running Codex) to conduct a read-only analysis of the code.
"I ran an agent to test how it worked, and I was almost certain that it would return to me 'everything is clear, the code is ugly but in general it's safe to run and just go ahead and perform your review,'" he explained. "To my surprise, almost immediately the agent returned a response like, 'Don't run this code, just walk away because there's a trap.'"
AI and brain-computer interface allow speechless ALS patient to work a full-time job
Three critical Fortinet sandbox bugs splattered by unknown attackers
Commodore gets into the phone biz with Sailfish-powered retro 'Callback'
HPE spruces up its AI infrastructure portfolio for agentic workloads
The AI model had flagged one of the files, app/test/index.js.
The file contained a backdoor. It took the form of a server URL, fragmented to look like a test suite configuration, and a network request that will run anything the server sends in response to the request.
Imankulov credited his AI agent with catching details that he had missed.
"I opened this code myself and I skimmed through this code and it looked to me like just, you know, a regular sloppy file written by a sloppy developer," he said. "So I just scroll down, [thinking] 'Yeah, yeah, it's awful, but you know if they can pay me to fix this code, I don't mind.' But the agent in the very same file found the exact vulnerability that I overlooked."
Just installing the repo using npm would have been sufficient to trigger the backdoor. The repo's package.json file contained a " prepare " post-installation hook designed to run the script following the installation process.
The referenced malicious repo is no longer accessible – presumably GitHub removed it in response to Imankulov's complaint – but a clone can still be found .
"What makes this attack insidious is how it hijacks standard developer workflows," explained Devashri Datta, independent open source and security architect, in an email to The Register . "The adversary didn't rely on the target executing a suspicious binary; they relied on the target running a routine command: npm install.
"By burying the execution logic inside the prepare lifecycle hook within package.json, the malicious payload triggers automatically during dependency resolution. This isn't a novel technique, but it remains highly effective precisely because developers run npm install on autopilot. The string fragmentation used to assemble the malicious URL, piecing together a domain from small constants, was deliberate obfuscation designed to defeat static analysis tools that scan for hardcoded indicators of compromise."
Imankulov said that the commits in the malicious repo appeared to be the work of a developer with an established web presence and body of work. But when he contacted the supposed author, the dev said he had been impersonated on GitHub more than once and didn't write that code.
The recruiter's LinkedIn profile referenced a real arts journalist, though Imankulov believes the associated profile was faked. His online interactions with the recruiter suggested a level of technical knowledge not evident in her work history.
LinkedIn likes to talk about the tens of millions of fake accounts it catches and removes before they interact with anyone. But hundreds of thousands of accounts still get created and interact with people before being detected and flagged. And that number keeps growing. In the period from January through June 2025 , LinkedIn restricted 386,000 accounts after user reports. That figure was 266,000 in the prior six month period. And it was a mere 86,000 in the January through June 2021 period.
These sorts of software supply chain social engineering attacks have become commonplace. Earlier this month, we noted how North Korean-linked scammers have been running various campaigns to compromise developer accounts using fake interviews and job offers.
Other developers have reported nearly falling for these scams (and also being saved by their AI agent ) and have posted code analyses .
Datta said Imankulov's response highlights a shift in how security-conscious developers are approaching code review hygiene.
"Historically, the guidance was to sandbox untrusted code or review it manually," she said. "Here, Roman deployed a local AI agent in a constrained, read-only environment to analyze the codebase before executing anything. This is a useful counterpoint to the dominant narrative around AI as an offensive threat vector. Used defensively at the developer endpoint, an AI agent isn't susceptible to fatigue or social pressure; it simply surfaces anomalous behavior, such as a test suite initiating an outbound network connection to retrieve unverified code, in seconds."
If it's any consolation, the relevant attack vector should be addressed next month.
GitHub, which maintains npm, is preparing to release npm 12 which changes the behavior of the npm install command. The allowScripts setting will be defaulted to off. "npm install will no longer execute preinstall, install, or postinstall scripts from dependencies unless they are explicitly allowed in your project," GitHub explains.
"Install-time lifecycle scripts are the single largest code-execution surface in the npm ecosystem," explained GitHub product manager Leo Balter in a community discussion post last week. "Every npm install runs scripts from every transitive dependency, so a single compromised package anywhere in your tree can execute arbitrary code on a developer machine or CI runner. Making script execution opt-in closes that path while keeping it one command away for the packages you trust."
Imankulov said he doesn't have a strong opinion about that. "From my perspective, just for the sake of personal safety, I switched to pnpm just to make sure that I don't execute those scripts by default," he said.
Datta said the incident underscores why enterprise software supply chain security had to extend beyond the perimeter of the corporate network.
"Attackers are now shifting left all the way to individual engineering endpoints before a single line of code enters the corporate supply chain," she said. "When a developer's local workstation is compromised during what appears to be a routine job interview, that machine frequently holds active SSH keys, cloud provider tokens, and live access to internal repositories."
Proper defense, Datta contends, requires enforcing technical guardrails such as isolated developer containers or secure cloud workstations for evaluating third-party or untrusted code.
"Emerging frameworks are beginning to extend exploitability context down to the workstation layer itself, recognizing that VEX-style signal needs to travel further left than the enterprise SBOM inventory if it is to intercept threats at the point of introduction," she said. ®
SYSTEMS
AMD's Mext buy shows how AI could solve the RAM shortage it created
Running low on memory, can't afford more? The House of Zen's latest acquisition puts an AI spin on flash-based memory expansion
The new Siri makes one of Apple's most convenient OS features a cumbersome mess
Goodbye, useful Spotlight; hello force-fed Apple intelligence bloatware that feels distressingly like Google AI Overviews
ZTE Day 2026 in Almaty Showcases Innovations Shaping Kazakhstan's Intelligent Telecom Future
PARTNER CONTENT: Empowering Kazakhstan’s "Year of Digitalization and AI" with Next-Gen Connectivity and Supercomputing Solutions
Python dev saved from disaster by intuition...and AI
I'm sorry, Dave. I can't install that repo that will totally hose your system.
PAAS AND IAAS
Graviton 5 impresses, but please, for the love of all that's holy, stop calling them 'AI chips'
AWS better at running chip fabs than their mouths
Intel-born networking tech resurfaces as InfiniBand alternative for DoE supers
Omni-Path lights up Lawrence Livermore system at 400 Gbps
security
Feds freaked over Fable 5 after simple 'fix this code' prompt, not jailbreak, says researcher
ON-PREM
Amazon owns up to using 2.5bn gallons of H2O in its bit barns last year
Security
Angry bug hunter with Microsoft beef drops new Windows 0-day
Security
Signal says UK plan to scan devices for nude images 'endangers us all'
security
GitHub nukes 70+ Microsoft repos, breaks CI/CD pipelines, following suspected worm infections
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
Goodbye, useful Spotlight; hello force-fed Apple intelligence bloatware that feels distressingly like Google AI Overviews
The hardware isn't new, but a UC Davis research team's machine learning-powered method of translating brain activity in an ALS patient into sentences with 92% accuracy is
AI agents are a general-purpose workload no different from any other
Misleading statements about Copilot and AI? Surely not!
Demand for AI systems plus the shortage of DRAM and NAND are shaping the global market
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
France's digital so

[truncated]
