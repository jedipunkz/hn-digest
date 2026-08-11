---
source: "https://huggingface.co/blog/jedisct1/lawyers-and-local-ai"
hn_url: "https://news.ycombinator.com/item?id=49255176"
title: "French lawyers are being told to use open source AI models for confidential data"
article_title: "French lawyers are being told to keep confidential data out of ChatGPT, Claude, and Gemini"
author: "jedisct1"
captured_at: "2026-08-11T09:51:46Z"
capture_tool: "hn-digest"
hn_id: 49255176
score: 2
comments: 0
posted_at: "2026-08-11T08:59:50Z"
tags:
  - hacker-news
  - translated
---

# French lawyers are being told to use open source AI models for confidential data

- HN: [49255176](https://news.ycombinator.com/item?id=49255176)
- Source: [huggingface.co](https://huggingface.co/blog/jedisct1/lawyers-and-local-ai)
- Score: 2
- Comments: 0
- Posted: 2026-08-11T08:59:50Z

## Translation

タイトル: フランスの弁護士は、機密データにオープンソース AI モデルを使用するよう指示されています
記事のタイトル: フランスの弁護士は、ChatGPT、Claude、Gemini に機密データを持ち込まないように言われています
説明: 「Hugging Face」に関する Frank Denis によるブログ投稿

記事本文:
フランスの弁護士はChatGPT、Claude、Geminiに機密データを持ち込まないよう命じられている
ハグ顔モデル
記事に戻る a]:hidden">
フランスの弁護士はChatGPT、Claude、Geminiに機密データを持ち込まないよう命じられている
コミュニティ記事が公開されました
2026 年 8 月 9 日 賛成票 1
フランスの弁護士は現在、いつになく率直な言葉で、「顧客の機密データをクロード、ChatGPT、Gemini、またはその他の独自のクラウド AI に送信しないでください」と告げられました。
そして、「トレーニングなし」を約束する企業計画でさえ問題を解決することはできません。
フランスの弁護士を代表するフランス国立弁護士評議会（CNB）は、2026年3月12～13日の総会で、クライアントや事件のデータを外部の生成AIサービスに決して開示しないという規則を明示した。
そうすることにより、弁護士が懲戒または法的制裁を受ける可能性があります。
CNB は標準的なクラウド AI 保証を拒否しています。独自のプロバイダーは、アップロードされたデータを再利用しないと約束する場合がありますが、その主張は一般に「証明可能でも検証可能でもありません」。
言い換えれば、契約書や「トレーニングなし」の保証は機密保持の証拠ではありません。
そして、データ主権は単なる地域セレクターではありません。 CNBは、企業はデータとモデルがどこにホストされているか、誰がサーバーを管理しているか、そしてそれらの企業が域外法の対象となるかどうかを検討する必要があると述べており、これは明示的に米国企業を指している。
これは法務チームのクロードに直接影響します。アンスロピック氏によると、商用データは米国に保存されており、トラフィックは米国、欧州、アジア、オーストラリアを経由する可能性があるという。したがって、サービスを法務チームに売り込んだからといって、クラウドがソブリンになるわけではありません。
したがって、CNB は、主権リスクを排除するための 2 つの構造的な方法を特定しています。1 つは、データが流出する前に匿名化または仮名化するか、または企業独自のシステムに LLM をインストールすることです。

ローカルのオープンモデルを実行しています。
しかし、真の匿名化は、単に名前を削除することよりもはるかに困難です。金額、プロジェクト名、文体を組み合わせると、クライアントを再識別できます。また、ライブ会議の文字起こしなどのワークフローでは、事前の仮名化がまったく不可能な場合があります。
機密作業に対する実際的な推奨事項は、ローカル推論です。つまり、管理するマシンやインフラストラクチャ上でオープンソース モデルまたはオープンウェイト モデルを実行します。 Claude、ChatGPT、Gemini、およびリモート API を、純粋に公開されたデータまたは不可逆的に匿名化されたデータとして保持します。
そして、この点において弁護士は特別な存在ではありません。医師、会計士、人事チーム、コンサルタント、エンジニア、研究者、公共団体も、法的、契約上、または専門的に保護されたデータを扱います。
したがって、彼らは同じアーキテクチャ上の疑問に直面しています。つまり、そのデータはシステムから流出すべきでしょうか?
これがフランスにとどまる可能性は低い。
ヨーロッパの弁護士および法律協会を代表する CCBE は、オンプレミス AI が機密保持のための最も安全なオプションであるとすでに説明しています。したがって、ヨーロッパ全体、そして秘密と信頼に基づいて構築された専門職全体にわたって、同様の監視が行われることが予想されます。
仕事で機密データや主権に関わるデータに触れる場合は、独自のクラウド AI をデフォルトとして扱うのはやめてください。代わりに、開いたモデルをローカルで実行します。
また、ヨーロッパで AI サービスを販売する場合は、真のローカル モードまたはセルフホスト モードを提供してください。
契約はデータ主権ではありません。
フランス国立弁護士評議会: 「たとえ独自のモデルの提供者が次のように述べたとしても
アップロードされたデータは再利用されないという主張ですが、この主張は一般に証明も検証もできません。
主権、隔離、機密性を保証する唯一の方法は、匿名化とオープンソース モデルによるローカル Al です。」
IA に準拠した CNB の規定
画像、音声、ビデオをアップロードするには、テキスト入力をドラッグするか、貼り付けるか、C

ここを舐めています。ここをタップまたは貼り付けて画像をアップロード コメント · コメントするにはサインアップまたはログインしてください

## Original Extract

A Blog post by Frank Denis on Hugging Face

French lawyers are being told to keep confidential data out of ChatGPT, Claude, and Gemini
Hugging Face Models
Back to Articles a]:hidden">
French lawyers are being told to keep confidential data out of ChatGPT, Claude, and Gemini
Community Article Published
August 9, 2026 Upvote 1
French lawyers have now been told, in unusually blunt terms: do not send confidential client data to Claude, ChatGPT, Gemini, or any other proprietary cloud AI.
And even enterprise plans that promise “no training” do not solve the problem.
At its General Assembly on 12–13 March 2026, France’s National Bar Council (CNB), which represents French lawyers, made the rule explicit: never disclose client or case data to external generative AI services.
Doing so may expose a lawyer to disciplinary or legal sanctions.
The CNB rejects the standard cloud AI reassurance. A proprietary provider may promise not to reuse uploaded data, but that claim is generally “neither provable nor verifiable.”
In other words, contracts and “no training” assurances are not proof of confidentiality.
And data sovereignty is not simply a region selector. The CNB says firms must consider where the data and models are hosted, who controls the servers, and whether those companies are subject to extraterritorial laws—explicitly pointing to US companies.
This directly affects Claude for Legal Teams. Anthropic says that commercial data is stored in the US and that traffic may be routed through the US, Europe, Asia, and Australia. So, marketing the service to legal teams does not make the cloud sovereign.
The CNB therefore identifies two structural ways to eliminate sovereignty risk: anonymize or pseudonymize the data before it leaves, or install the LLM on the firm’s own systems—for example, by running a local open model.
But genuine anonymization is much harder than simply removing names. Amounts, project names, and writing style can re-identify a client when combined. And for workflows such as live meeting transcription, prior pseudonymization may simply be impossible.
The practical recommendation for confidential work is local inference: run open-source or open-weight models on machines and infrastructure you control. Keep Claude, ChatGPT, Gemini, and remote APIs for genuinely public or irreversibly anonymized data.
And lawyers are not unique in this regard. Doctors, accountants, HR teams, consultants, engineers, researchers, and public bodies also handle legally, contractually, or professionally protected data.
They therefore face the same architectural question: should that data leave their systems at all?
This is also unlikely to stop at France.
The CCBE, which represents Europe’s bars and law societies, already describes on-premises AI as the most secure option for confidentiality. So, expect the same scrutiny across Europe and across professions built on secrecy and trust.
If your work touches confidential or sovereignty-sensitive data, stop treating proprietary cloud AI as the default. Instead, run open models locally.
And if you sell AI services in Europe, offer a genuinely local or self-hosted mode.
A contract is not data sovereignty.
French National Bar Council: "even if the provider of a proprietary model states
that uploaded data will not be reused, this claim is generally neither provable nor verifiable.
The only ways to guarantee sovereignty, isolation and confidentiality are anonymization and local Al with an open-source model."
Le CNB a codifié la conformité IA
Upload images, audio, and videos by dragging in the text input, pasting, or clicking here . Tap or paste here to upload images Comment · Sign up or log in to comment
