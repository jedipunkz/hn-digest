---
source: "https://www.theregister.com/ai-and-ml/2026/06/04/please-do-not-vibe-f-up-this-software-broken-backups-spark-ai-coding-row-in-rsync-project/5251189"
hn_url: "https://news.ycombinator.com/item?id=48407503"
title: "'Please do not vibe f–- up this software': Broken backups spark AI coding row"
article_title: "'Please do not vibe f--- up this software': Broken backups spark AI coding row in rsync project"
author: "Bender"
captured_at: "2026-06-05T04:34:00Z"
capture_tool: "hn-digest"
hn_id: 48407503
score: 3
comments: 1
posted_at: "2026-06-05T03:15:00Z"
tags:
  - hacker-news
  - translated
---

# 'Please do not vibe f–- up this software': Broken backups spark AI coding row

- HN: [48407503](https://news.ycombinator.com/item?id=48407503)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/06/04/please-do-not-vibe-f-up-this-software-broken-backups-spark-ai-coding-row-in-rsync-project/5251189)
- Score: 3
- Comments: 1
- Posted: 2026-06-05T03:15:00Z

## Translation

タイトル: 「このソフトウェアを興奮させないでください」: 壊れたバックアップが AI コーディング問題を引き起こす
記事のタイトル: 「このソフトウェアを興奮させないでください」: 壊れたバックアップが rsync プロジェクトの AI コーディング行を引き起こす
説明: ユーザーはバックアップの失敗を調査し、クロード支援によるコミットを見つけます。ベテランエンジニアはこう言い返す。

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
「このソフトウェアを起動しないでください」: 壊れたバックアップが rsync プロジェクトの AI コーディング行を引き起こす
ユーザーはバックアップの失敗を調査し、クロード支援によるコミットを見つけます。ベテランエンジニアは「私は単に『テストスイートをPythonに変換する』というバイブコードを書いただけではない」と反論する。
最近のアップデートの後、一部の rsync ユーザーの間で増分バックアップが失敗し始めました。プロジェクトのコミット履歴で発見されたものは、すぐに日常的なバグ探しを AI 生成コードをめぐるさらなる争いに変えました。
論争の中心は、複数の脆弱性を修正するために今年初めに公開されたセキュリティに重点を置いたリリースである rsync 3.4.3 です。アップグレードの直後、一部のユーザーは増分バックアップのワークフローが期待どおりに動作しなくなったと報告し、あるユーザーはバックアップ システムが完全バックアップ以外で失敗したと述べました。
Rsyncの作成者であるアンドリュー・トリジェル氏は、「Rsyncとアウトレイジ」と題したMediumの投稿で批判に反論し、多くのコメント投稿者はAIツールが実際にどのように使用されたかを理解せずに結論を導いていると主張した。
Rsync は、Discord サーバーで 3 人によって維持される週末のサイド プロジェクトではありません。 1990 年代に初めてリリースされ、今でも Unix と Linux の世界で最も広く使用されているファイル同期およびバックアップ ユーティリティの 1 つです。無数のバックアップ製品、スクリプト、NAS アプライアンス、IT 部門は、予期せぬ事態が発生することなく静かにその仕事を実行するこの製品に依存しています。
そのため、このプロジェクトにおける AI 支援開発の提案は、他の場所よりもはるかに議論の余地があります。
ユーザーが rsync の最近のコミット履歴を調べ始めていなければ、バックアップの問題はごく普通のバグレポートのままだったかもしれません。 rsync以来、彼らはそれを発見しました

3.4.1 では、数十のコミットが、rsync 作成者の Andrew Tridgell と Anthropic の AI アシスタントの Claude を指す「tridge and claude」によるものであると考えられています。
この発見により、「このソフトウェアをクソにしないでください」というタイトルの強い文言の GitHub 投稿が生成されました。これは、コーディング タスクを AI モデルに渡し、その結果を信頼するという、ますます一般的になっている慣行に言及したものです。
そこから、議論は Reddit や Hacker News に広がり、そこで会話はバックアップのバグから、重要なオープンソース インフラストラクチャに侵入する AI 生成コードに関するより広範な議論に移りました。
ベテラン開発者の Tridgell 氏は、rsync 3.4.3 で一部のバックアップ ワークフローに影響を与えるリグレッションが発生したことを認め、それらをプロジェクトの既存のテスト スイートではカバーされていない「有効な (しかし異常な) ユースケース」であると説明しました。 「あなたのrsyncの使用例がこれらのリグレッションの影響を受けたのであれば、お詫び申し上げます」と同氏は書いている。
しかしトリジェル氏は、単に開発をクロードに引き渡し、最善の結果を期待しただけだという考えを押し返した。
Tridgell 氏によると、最も目に見える AI 支援の作業には、セキュリティ テストを改善し、コードベースを強化する広範な取り組みの一環として、rsync の老朽化したシェルスクリプト テスト スイートを Python で書き直す作業が含まれていました。同氏は、フレームワークを自分で設計し、OpenAIのCodexやGoogleのGeminiと並行してClaudeを「単調な作業」と表現した作業に使用し、結果として得られたコードを手動でレビューしたと述べた。
「私は単に『テストスイートをPythonに変換する』というバイブコーディングをしただけではありません」と彼は書いている。 「私は 40 年の経験を持つソフトウェア エンジニアです。」
トリジェル氏はまた、保守担当者が大量のセキュリティレポートに対処することが増えており、その多くはAIによって生成されており、広く使用されているオープンソースソフトウェアを安全に保つために必要な作業負荷が劇的に増加していると主張した。
SaaS の黙示録は待てますが、Salesforce にはまだ顧客がいます。

それは彼らを望んでいます
Gemini、30,000行のコードパージと偽のリカバリレポートで告発
ライナス・トーバルズがバイブコーディングを試みる、世界は何とか無傷のまま
研究者らは雰囲気を消し去り、AI コーディングのためのより良いモデルを提案したいと考えている
「ソフトウェア エンジニアリングの世界は、ここ数か月で劇的に変化しました」と彼は書いています。 「大量の報告に直面して、IT セキュリティとソフトウェアの保守の世界は、ここ数週間で完全かつ完全に変わりました。」
Tridgell 氏は、AI 支援開発から手を引くどころか、rsync がセキュリティの向上に重点を置いた大規模な 3.5 リリースに向かう中、ツールを使い続けるつもりであることを示唆しました。同氏はまた、OpenBSD の openrsync プロジェクトに飛びつくと脅すユーザーを厳しく非難し、rsync の新しいテスト スイートは現在、代替実装に対して実行すると数十件の失敗を報告していると指摘した。
その安心感が批評家を満足させるかどうかはまだ不明だ。しかし、少なくとも、この全体は、AI 支援開発とバックアップ ソフトウェアが非常に優れた組み合わせであることを示しています。 1 つはマシンを信頼することですが、もう 1 つは人間が信頼していないために存在します。 ®
AI と ML
AIのスプリントを遅らせることが「世界にとって良いことだ」とアンスロピック氏は語る
警戒を求める嘆願は、AI最大のライバルであるOpenAIを破ってIPOを申請したのと同じ週に発表された
ピンクは、偽のヘルプデスクへの電話を使用して資格を盗む最新の悪党チームです
混沌とした犯罪集団 Lapsus$ によって広められたおなじみの戦術
ZTEとパートナーは2026年のエンジニアリング能力構築プログラムを通じて世界的なICT人材を育成
世界的な ICT 専門家が深センに集まり、最先端のエンジニアリング実践を習得し、国際協力を促進します
カナダは米国のボットから解放され、独自のAIを開発したいと考えている
別の同盟国もアメリカのAIへの依存に疑問を呈する
SaaS
AWSはイーロン・マスク氏のGrok intを隠蔽すると報じられている

o 企業需要がゼロにもかかわらず、基盤は整っている
フロンティアモデルのエナジードリンク
OpenAI のエージェントが 10 年にわたる DoS 攻撃を連鎖させ、Web サーバーを数秒でクラッシュさせた
AI と ML
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
セキュリティ
レドモンドが警察に通報、マイクロソフトに「屈辱を受けた」と不満を抱く0日ハンターが「骨が砕けるほどの衝撃」を誓う
セキュリティ
軍隊の携帯電話が位置データを外国の敵に提供した
セキュリティ
すべてのパスワードは Active Directory の説明フィールドに保存されました
パーソナルテクノロジー
カリフォルニア州、3Dプリントによるゴーストガンのアルゴリズムによる死亡を宣言する法案を可決
データ主権のトレードオフを克服する
避けられないトレードオフであるデータ主権はネットワークにとって実際に何を意味するのでしょうか?もっと詳しく知る。
プロンプトからエクスプロイトまで: LLM は API 攻撃をどのように変化させているか
最新のアプリケーションは API 主導で相互接続されており、多くの場合、過剰な権限が与えられているため、AI 支援攻撃の理想的な標的となっています。
未来の構築: Kubernetes 向けエンタープライズ データ サービスのロックを解除する
インフラストラクチャのサイロを排除し、標準化されたエンタープライズ グレードのクラウド ネイティブ プラットフォームを確立する方法を発見してください。
Microsoft 365 が見逃す高度な攻撃を Behavioral AI セキュリティで捕捉
Microsoft 365 は企業コミュニケーションのバックボーンであり、そのネイティブ セキュリティが既知のものや騒々しいものを除外します。
ランサムウェア侵害の混乱に足を踏み入れ、対応スキルをテストし、他の IT およびセキュリティの専門家とチームを組んでサイバー犯罪者を出し抜こう
仮想サイバー復旧シミュレーション
ランサムウェア攻撃の勢いは衰えていませんが、私たちも同様です。 Druva の人気イベント「Escape Ransomware」が完全にバーチャルになりました。
大規模なエージェント AI: パイロットから運用まで
AI の導入を大規模に推進することで、真の ROI を実現する方法を学びましょう。
警戒を求める嘆願は同じ週に行われる

AI最大のライバルOpenAIがIPO申請まで
混沌とした犯罪集団 Lapsus$ によって広められたおなじみの戦術
別の同盟国もアメリカのAIへの依存に疑問を呈する
科学者と業界リーダーはDNA合成スクリーニングの義務化を推進
警戒を求める嘆願は、AI最大のライバルであるOpenAIを破ってIPOを申請したのと同じ週に発表された
混沌とした犯罪集団 Lapsus$ によって広められたおなじみの戦術
別の同盟国もアメリカのAIへの依存に疑問を呈する
科学者と業界リーダーはDNA合成スクリーニングの義務化を推進
AIのスプリントを遅らせることが「世界にとって良いことだ」とアンスロピック氏は語る
警戒を求める嘆願は、AI最大のライバルであるOpenAIを破ってIPOを申請したのと同じ週に発表された
ピンクは、偽のヘルプデスクへの電話を使用して資格を盗む最新の悪党チームです
混沌とした犯罪集団 Lapsus$ によって広められたおなじみの戦術
カナダは米国のボットから解放され、独自のAIを開発したいと考えている
別の同盟国もアメリカのAIへの依存に疑問を呈する
OpenAI のエージェントが 10 年にわたる DoS 攻撃を連鎖させ、Web サーバーを数秒でクラッシュさせた
Codex が HTTP/2 爆弾を投下
AIの重鎮、自社の技術がテロリストの生物兵器開発に役立つ可能性があると警告
科学者と業界リーダーはDNA合成スクリーニングの義務化を推進
慈悲深い独裁者ザック氏、キーロギングによるプライバシー侵害からメタスタッフに30分間の休憩を与える
スタッフの活動を盗み見てAIにコンピュータの使い方を教えるテックビジネス
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

Users probe backup failures find Claude-assisted commits. Veteran engineer retorts:

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
'Please do not vibe f--- up this software': Broken backups spark AI coding row in rsync project
Users probe backup failures find Claude-assisted commits. Veteran engineer retorts: "I did not just vibe-code 'convert test suite to python'."
Incremental backups started failing for some rsync users after a recent update, and what they found in the project's commit history quickly turned a routine bug hunt into yet another fight over AI-generated code.
The controversy centers on rsync 3.4.3, a security-focused release published earlier this year to fix multiple vulnerabilities. Shortly after the upgrade, some users reported that incremental backup workflows were no longer behaving as expected, with one user saying their backup system failed on anything other than a full backup .
Rsync creator Andrew Tridgell has pushed back against the criticism in a Medium post titled " Rsync and Outrage ," arguing that many commenters have drawn conclusions without understanding how the AI tools were actually used.
Rsync is not a weekend side project maintained by three people in a Discord server. First released in the 1990s, it remains one of the most widely used file synchronization and backup utilities in the Unix and Linux world. Countless backup products, scripts, NAS appliances, and IT departments depend on it quietly doing its job without surprises.
That makes any suggestion of AI-assisted development in the project far more contentious than it might be elsewhere.
The backup issue might have remained a fairly ordinary bug report had users not started poking around in rsync's recent commit history. They found that since rsync 3.4.1, dozens of commits have been attributed to "tridge and claude," referring to rsync creator Andrew Tridgell and Anthropic's AI assistant Claude.
The discovery prompted a strongly worded GitHub post titled "Please Do Not Vibe Fuck Up This Software," a reference to the increasingly common practice of handing coding tasks to AI models and trusting the results.
From there, the discussion spread to Reddit and Hacker News , where the conversation shifted from a backup bug to a broader debate about AI-generated code finding its way into critical open source infrastructure.
Veteran developer Tridgell acknowledged that rsync 3.4.3 introduced regressions affecting some backup workflows, describing them as "valid (but unusual) use cases" that were not covered by the project's existing test suite. "I apologize if your use case of rsync was hit by these regressions," he wrote.
But Tridgell pushed back on suggestions that he had simply handed development over to Claude and hoped for the best.
According to Tridgell, the most visible AI-assisted work involved rewriting rsync's aging shell-script test suite in Python as part of a broader effort to improve security testing and harden the codebase. He said he designed the framework himself, used Claude alongside OpenAI's Codex and Google's Gemini for what he described as "grunt work," and manually reviewed the resulting code.
"I did not just vibe-code 'convert test suite to python,'" he wrote. "I'm a software engineer with 40 years experience."
Tridgell also argued that maintainers are increasingly dealing with a flood of security reports, many of them AI-generated, which has dramatically increased the workload required to keep widely used open source software secure.
The SaaS-pocalypse can wait, Salesforce still has customers where it wants them
Gemini accused of 30,000-line code purge and fake recovery report
Linus Torvalds tries vibe coding, world still intact somehow
Researchers want to kill the vibe, propose better model for AI coding
"The world of software engineering has changed dramatically in the last few months," he wrote. "The world of IT security and maintaining software in the face of the flood of reports has completely and utterly changed just in the last few weeks."
Far from backing away from AI-assisted development, Tridgell suggested he intends to continue using the tools as rsync heads toward a larger 3.5 release focused on security improvements. He also took a swipe at users threatening to jump ship to OpenBSD's openrsync project, noting that rsync's new test suite currently reports dozens of failures when run against the alternative implementation.
Whether that reassurance satisfies critics is still unclear. But if nothing else, the whole thing demonstrates that AI-assisted development and backup software make for a combustible combination. One involves trusting a machine – the other exists because people don't. ®
AI AND ML
'It would be good for the world' to slow down AI sprints, Anthropic says
The plea for caution comes the same week it beat AI archrival OpenAI to filing for an IPO
Pink is the latest goon squad to use fake helpdesk calls to steal creds
A familiar tactic popularized by chaotic crime crew Lapsus$
ZTE and partners nurture global ICT talent through 2026 engineering capacity building program
Global ICT experts gather in Shenzhen to master cutting-edge engineering practices and foster international collaboration
Canada wants to make its own AI, break free from US bots
Another ally questions reliance on American AI
SaaS
AWS reportedly to tuck Elon Musk's Grok into Bedrock, despite zero enterprise demand
The energy drink of frontier models
OpenAI's agent chained decade-old DoS attacks to crash web servers in seconds
AI and ML
Netflix wiz creates app to slash AI bills, then open sources it
Security
Disgruntled 0-day hunter 'humiliated' by Microsoft pledges 'bone shattering drop' as Redmond calls cops
Security
Troops’ phones gave away location data to foreign adversaries
SECURITY
All the passwords were stored in Active Directory description fields
Personal tech
California passes bill declaring death-by-algorithm to 3D-printed ghost guns
Overcoming the trade-offs in data sovereignty
What does data sovereignty actually mean for your network, which trade-offs are unavoidable? Learn more.
From Prompt to Exploit: How LLMs Are Changing API Attacks
Modern applications are API-driven, interconnected, and often over-permissioned, making them an ideal target for AI-assisted attacks.
Architecting the Future: Unlocking Enterprise Data Services for Kubernetes
Join us to discover how to eliminate infrastructure silos and establish a standardized, enterprise-grade cloud-native platform.
Catch the Advanced Attacks Microsoft 365 Misses with Behavioral AI Security
Microsoft 365 is the backbone of enterprise communication, and its native security filters out the known and the noisy.
Step into the chaos of a live ransomware breach, test your response skills, and team up with other IT and security pros to outsmart cybercriminals
Virtual Cyber Recovery Simulation
Ransomware attacks aren’t slowing down, and neither are we. Druva’s hit event, Escape Ransomware, is now fully virtual.
Agentic AI at Scale: From Pilot to Production
Join us to learn how to unlock real ROI by driving adoption of AI at scale.
The plea for caution comes the same week it beat AI archrival OpenAI to filing for an IPO
A familiar tactic popularized by chaotic crime crew Lapsus$
Another ally questions reliance on American AI
Scientists and industry leaders push for mandatory DNA synthesis screening
The plea for caution comes the same week it beat AI archrival OpenAI to filing for an IPO
A familiar tactic popularized by chaotic crime crew Lapsus$
Another ally questions reliance on American AI
Scientists and industry leaders push for mandatory DNA synthesis screening
'It would be good for the world' to slow down AI sprints, Anthropic says
The plea for caution comes the same week it beat AI archrival OpenAI to filing for an IPO
Pink is the latest goon squad to use fake helpdesk calls to steal creds
A familiar tactic popularized by chaotic crime crew Lapsus$
Canada wants to make its own AI, break free from US bots
Another ally questions reliance on American AI
OpenAI's agent chained decade-old DoS attacks to crash web servers in seconds
Codex drops an HTTP/2 Bomb
AI heavyweights warn their tech could help terrorists develop bioweapons
Scientists and industry leaders push for mandatory DNA synthesis screening
Benevolent dictator Zuck will give Meta staff 30-minute breaks from keylogging privacy assault
Tech biz teaching AI to use computers by slurping staff activity
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
