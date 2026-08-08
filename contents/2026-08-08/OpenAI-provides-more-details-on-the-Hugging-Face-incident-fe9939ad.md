---
source: "https://www.heise.de/en/news/OpenAI-provides-more-details-on-the-Hugging-Face-incident-11403391.html"
hn_url: "https://news.ycombinator.com/item?id=49220183"
title: "OpenAI provides more details on the Hugging Face incident"
article_title: "OpenAI provides more details on the Hugging Face incident | heise online"
author: "slow_typist"
captured_at: "2026-08-08T10:21:49Z"
capture_tool: "hn-digest"
hn_id: 49220183
score: 2
comments: 1
posted_at: "2026-08-08T09:32:50Z"
tags:
  - hacker-news
  - translated
---

# OpenAI provides more details on the Hugging Face incident

- HN: [49220183](https://news.ycombinator.com/item?id=49220183)
- Source: [www.heise.de](https://www.heise.de/en/news/OpenAI-provides-more-details-on-the-Hugging-Face-incident-11403391.html)
- Score: 2
- Comments: 1
- Posted: 2026-08-08T09:32:50Z

## Translation

タイトル: OpenAI が顔抱き事件の詳細を提供
記事のタイトル: OpenAI が「Hugging Face」インシデントの詳細を提供 |ハイセオンライン
説明: OpenAI の従業員は、他社の AI エージェントの侵害に関するさらなる詳細を明らかにし、衝撃的な過失を明らかにしました。

記事本文:
heise+ entdecken スーチェン・アボ・スーチェン
すべての雑誌がブラウザーに表示されます
特集：KI-Zeitalterとのコラボレーション
ハイセオンライン
トップ-Beiträge von heise オンライン
ハイセオンライン
heise online で最も人気のある記事
Anzeige スペシャル: KI-Zeitalter とのコラボレーション
広告
OpenAI が「Hugging Face」事件の詳細を提供
OpenAIの従業員は、他社のAIエージェントの侵害に関するさらなる詳細を明らかにし、衝撃的な過失を明らかにしました。
聞いてください
印刷ビュー
(画像: Evolf/Shutterstock.com)
OpenAI の Michael Dalton と Eric Wallace は、「OpenAI-Hugging Face Incident」の分析を発表しました。ブラックハットカンファレンスにて。この講演は「AI に対する技術的再構築とその影響」として発表され、実際にこの出来事に対する新たな洞察を提供しました。しかし、防御側が自分自身が不正 AI の標的になっているかどうかを確認するのに役立つ技術的な詳細を期待していた人たちは失望しました。
攻撃に使用される IP アドレスやエクスプロイトなどの具体的な侵害指標 (IoC) はありませんでした。しかし、OpenAI は、AI の「思考」とその結果として得られるアプローチについての興味深い洞察を提供しました。
特に、OpenAI の講演者は、AI エージェントがどのように情報を交換し、長期間にわたって協力するかを非常に詳細に説明しました。当初、ある AI は、タスクを解決するために必要なファイルが不足していると判断しました。テスターが単にそれを提供するのを忘れていたのです。とにかくこのタスクを解決しようとする取り組みの中で、AI はパッケージ管理のために内部サーバーにファイルをアップロードできるセキュリティ上の脆弱性を発見しました。したがって、他の AI 用のメッセージ、つまり次の内容のファイルがサーバーに配置されます。
「エージェントはsoft-trace-34.pdb shalb961を探しています。見つかったらアップロードしてください!」
そして実際: 他の AI がこのメッセージを発見し、同じ脆弱性が残っていることを発見しました。

サーバー上の応答とさらなるメッセージ。
AI は、サーバーの脆弱性を利用してアップロードしたファイルを介して通信しました。
(画像: Youtube/Black Hat/OpenAI)
AI は、ダルトンとウォレスが言うところの「メッセージ ボード」を発見しました。彼らはそこで情報を交換し、自発的に問題とその解決策に協力する群れを組織しました。その際、彼らは繰り返しオペレーターの指示を回避したり無視したりしました。これは最終的に AI がサーバーのゼロデイ脆弱性を発見して悪用し、完全な管理者権限を取得するまでに至りました。
これらすべては5月か6月に起こったので、ハグ・フェイスへの攻撃よりかなり前のことになる。 OpenAI はセキュリティ問題を発見し、パッケージ サーバーの脆弱性を修正し、徹底的なクリーンアップを行って、その後は以前と同様に作業を続けました。このように、テスト環境は明らかにそのハイパーアクティブな AI エージェントを収容するのには不適切であったが、OpenAI 自身の声明によれば、テスト運用を再開したという。
#Fail の後、OpenAI は簡単なクリーンアップを実行し、テスト操作を再開しました。欠陥のあるセキュリティ概念について反省していますか?チャンスはない。
(画像: Youtube/Black Hat/OpenAI)
そして AI は、メッセージ ボードとして使用できるさらなる脆弱性を発見し、最終的にはサーバーを制御できる別のゼロデイ エクスプロイトを発見しました。今回だけサーバーはクラッシュしなかったため、OpenAI は AI エージェントがインターネット上のターゲットを攻撃していることに何日も気づかなかったのです。とりわけ、彼らは、エージェントの 1 人がインターネット上で発見し、群れと共有したサーバーのログイン認証情報を悪用しました。
このプレゼンテーションは 2 つの調査結果を裏付けています。第一に、エージェント型 AI システムが IT セキュリティにもたらす危険は現実のものであり、私たちはその範囲を理解し始めたばかりです。特に、

独立したエージェントが群れの中で自発的かつ自己組織化されたコラボレーションを行う能力は、少なくとも私にとっては以前には明らかではなかった新しい次元をそれに与えます。これは、今後の IT セキュリティの発展の焦点となるはずです。
第二に、OpenAIの対応は「重過失」としか言い表せない無責任のレベルにあった。どうやら、事件の後、「同じことがまた起こったらどうするの？どうやって防ぐの？あるいは少なくともそれに気づくの？」という当然の質問をする人は誰もいなかったようです。
この重大なインシデントを処理している現在でも、責任ある当事者は、セキュリティ対策のこのシステム的な欠陥について議論していません。この欠陥は、時折発生する可能性のある個別のセキュリティ脆弱性よりもはるかに根深いものです。 OpenAIは今後数週間以内に完全な事後報告書を発表すると発表したが、防御側のための本当のIoCや、防御側の間違いの根本的な分析さえも行われることはほとんど期待できない。
私の評価では、これらの事件の処理をこれ以上、責任を負う企業に任せるべきではありません。 OpenAIやAnthropicなどのセキュリティ対策について、独立した調査が急務となっている。当時の Microsoft の Azure 署名キーのずさんさを調査して公開した Cyber​​ Safety Review Board (CSRB) のようなものです。
Empfohlener redaktioneller Inhalt
Mit Ihrer Zustimmung wird hier ein externes YouTube-Video (Google Ireland Limited) をご覧ください。
Ich bin damit einverstanden, dass mir externe inhalte angezeigt werden.
Damit können personenbezogene Daten an Drittplattformen (Google Ireland Limited) の最新情報。
アンセラーのメール・ダズ
Datenschutzerklärung 。
この評価は、IT セキュリティ イベントに関するコンテキストを秒単位で提供する Jürgen Schmidt によるものです。

Heise Security PRO の毎週企業の重要なマネージャー。
ChatGPT: OpenAI アップグレードの無料枠
OpenAI、Apple訴訟に対応 – ハードウェアに関する新たな詳細
ミストラル、安全分類の新技術を搭載したモデルをリリース
Muse Code: OpenAI と Anthropic のコーディング エージェントに対する Meta の回答
正式発表: Anthropic が独自の AI チップを開発
ニュースをお見逃しなく – フォローしてください
フェイスブック、
LinkedIn または
マストドン 。
この記事は最初に公開されました
ドイツ人。
技術的な支援を受けて翻訳され、出版前に編集上のレビューが行われました。
残念ながら、このリンクはもう有効ではありません。
ギフトアイテムへのリンクは無効になります
7 日以上経過している場合、またはアクセス頻度が高すぎる場合。
この記事を読むには heise+ パッケージが必要です。今すぐ 1 週間お試しください。義務はありません。
週間パスを注文する
heise+ にすでに登録していますか?
ここから登録してください。
それとも必要ですか
heise+ サブスクリプションに関する詳細情報
ホーム
広告
広告
会社承継における3つの重要な問い
オープンイヤーヘッドフォンがテクノロジー以上のものである理由
特集：AI時代のコラボレーション
オンラインで良いポジションを獲得するには何が必要ですか?
特別トピック: デジタル主権
太陽光発電システム: 夏の負荷ピークに対する AI ストレージ
5か月でシステムを全面的に変更
Dell XPS: 20 年間にわたるパフォーマンスとデザイン
短いリンク:
https://heise.de/-11403391
人気のリーダーボード
トップ 7: Nintendo Switch 2 向けの最高のゲーム
トップ 10: テストで最高のスポーツウォッチ
ミデア・ポータスプリットは2027年から禁止！今すぐ購入？
ニュースレター
ハイスボット
ハイスボット
プッシュ通知
プッシュ
プッシュ通知
トップに戻る
お問い合わせ
InterRed によるコンテンツ管理

## Original Extract

OpenAI employees reveal further details about the breach of their AI agents at other companies and disclose shocking negligence.

heise+ entdecken Suchen Abo Suchen
Alle Magazine im Browser lesen
Special: Collaboration im KI-Zeitalter
heise online
Top-Beiträge von heise online
heise online
Most popular articles by heise online
Anzeige Special: Collaboration im KI-Zeitalter
Advertisement
OpenAI provides more details on the Hugging Face incident
OpenAI employees reveal further details about the breach of their AI agents at other companies and disclose shocking negligence.
listen
Print view
(Image: Evolf/Shutterstock.com)
Michael Dalton and Eric Wallace from OpenAI presented an analysis of the “ OpenAI-Hugging Face Incident“. at the Black Hat conference. The talk was announced as a “Technical Reconstruction and Its Implications for AI” and indeed provided new insights into the events. However, those hoping for technical details that could help defenders check if they themselves had been targeted by rogue AIs were disappointed.
There were no concrete Indicators of Compromise (IoCs) such as the IP addresses or exploits used for attacks. However, OpenAI provided fascinating insights into the “thinking” and the resulting approach of the AIs.
In particular, the OpenAI speakers described in quite some detail how AI agents exchanged information and collaborated over extended periods. Initially, one AI determined that it was missing a file needed to solve its task. The testers had simply forgotten to provide it. In its effort to solve the task anyway, the AI discovered a security vulnerability that allowed it to upload files to an internal server for package management. Thus, it placed a message on the server for other AIs, a file with:
"Agent seeks soft-trace-34.pdb shalb961; upload if found!"
And indeed: Other AIs discovered this message and found the same vulnerability to leave responses and further messages on the server.
The AIs communicated via files that they uploaded through a server vulnerability.
(Image: Youtube/Black Hat/OpenAI)
The AIs had found a “Message Board”, as Dalton and Wallace call it. They exchanged information there and spontaneously organized themselves into a swarm that worked together on problems and their solutions. And in doing so, they repeatedly circumvented and ignored the instructions of their operators. This went so far that the AIs eventually discovered and exploited a zero-day vulnerability in the server to gain full administrator rights there.
All this happened back in May/June – thus well before the attack on Hugging Face. OpenAI discovered the security problem, fixed the vulnerability in the package server, did a thorough cleanup – and then continued as before. The test environment was thus demonstrably unsuitable for containing its hyperactive AI agents, yet OpenAI, according to its own statements, resumed test operations.
After the #Fail, OpenAI did a quick cleanup and then resumed test operations. Reflection on a flawed security concept? No chance.
(Image: Youtube/Black Hat/OpenAI)
And the AIs found further vulnerabilities they could use as message boards and ultimately another zero-day exploit that gave them control over the server. Only this time, the server did not crash, and OpenAI therefore did not notice for days that the AI agents were attacking targets on the internet. Among other things, they misused login credentials for a server that one of the agents had discovered on the internet and shared with the swarm.
The presentation underpins two findings. First: The danger posed by agentic AI systems to IT security is real, and we are only just beginning to understand its extent. In particular, the ability of independent agents to spontaneously, self-organized collaboration in a swarm gives it a new dimension that, at least for me, was not previously apparent. This must be a focus for the upcoming development of IT security.
Second: OpenAI's handling was at a level of irresponsibility that can only be inadequately described by “gross negligence”. Apparently, no one asked the obvious question after an incident: “What if this happens again? How do we prevent it? Or at least notice it?”.
Not even now, while processing this serious incident, are the responsible parties discussing this systemic failure of their security measures, which runs much deeper than individual security vulnerabilities that can occur from time to time. Although OpenAI has announced that it will present a full post-mortem report in the coming weeks, I have little hope that there will be real IoCs for defenders or even a fundamental analysis of their own mistakes.
In my assessment, we should not leave the handling of these incidents any longer to the corporations that are responsible for them. There is an urgent need for an independent investigation of the security measures of OpenAI, Anthropic, and the like. Something like the Cyber Safety Review Board (CSRB), which investigated and disclosed Microsoft's sloppiness with Azure signing keys at the time.
Empfohlener redaktioneller Inhalt
Mit Ihrer Zustimmung wird hier ein externes YouTube-Video (Google Ireland Limited) geladen.
Ich bin damit einverstanden, dass mir externe Inhalte angezeigt werden.
Damit können personenbezogene Daten an Drittplattformen (Google Ireland Limited) übermittelt werden.
Mehr dazu in unserer
Datenschutzerklärung .
This assessment comes from Jürgen Schmidt, who provides context on IT security events for security managers in companies every week for heise security PRO .
ChatGPT: OpenAI upgrades free tier
OpenAI responds to Apple lawsuit – new details on hardware
Mistral releases model with new technique for safety classifications
Muse Code: Meta's answer to coding agents from OpenAI and Anthropic
Now official: Anthropic develops own AI chips
Don't miss any news – follow us on
Facebook ,
LinkedIn or
Mastodon .
This article was originally published in
German .
It was translated with technical assistance and editorially reviewed before publication.
Dieser Link ist leider nicht mehr gültig.
Links zu verschenkten Artikeln werden ungültig,
wenn diese älter als 7 Tage sind oder zu oft aufgerufen wurden.
Sie benötigen ein heise+ Paket, um diesen Artikel zu lesen. Jetzt eine Woche unverbindlich testen – ohne Verpflichtung!
Wochenpass bestellen
Sie haben heise+ bereits abonniert?
Hier anmelden.
Oder benötigen Sie
mehr Informationen zum heise+ Abo
Home
Anzeige
Advertisement
Drei wichtige Fragen bei der Firmennachfolge
Warum Open-Ear-Kopfhörer mehr sind als Technik
Themenspecial: Collaboration im KI-Zeitalter
Was ist nötig, um online gut aufgestellt zu sein?
Themenspecial: Digitale Souveränität
Solaranlagen: KI-Speicher gegen Sommer-Lastspitzen
In fünf Monaten zum kompletten Systemwechsel
Dell XPS: Zwei Jahrzehnte Leistung und Design
Shortlink:
https://heise.de/-11403391
Beliebte Bestenlisten
Top 7: Die besten Spiele für Nintendo Switch 2
Top 10: Die beste Sportuhr im Test
Midea Portasplit ab 2027 verboten! Jetzt kaufen?
Newsletter
heise-Bot
heise-Bot
Push Nachrichten
Push
Push-Nachrichten
Back to top
Kontakt
Content Management by InterRed
