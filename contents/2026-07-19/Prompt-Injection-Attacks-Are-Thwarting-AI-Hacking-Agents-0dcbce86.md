---
source: "https://arstechnica.com/security/2026/07/now-defenders-are-embracing-the-prompt-injection-too/"
hn_url: "https://news.ycombinator.com/item?id=48964176"
title: "Prompt Injection Attacks Are Thwarting AI Hacking Agents"
article_title: "Now, defenders are embracing the prompt injection, too - Ars Technica"
author: "sbulaev"
captured_at: "2026-07-19T01:45:22Z"
capture_tool: "hn-digest"
hn_id: 48964176
score: 1
comments: 1
posted_at: "2026-07-19T01:26:55Z"
tags:
  - hacker-news
  - translated
---

# Prompt Injection Attacks Are Thwarting AI Hacking Agents

- HN: [48964176](https://news.ycombinator.com/item?id=48964176)
- Source: [arstechnica.com](https://arstechnica.com/security/2026/07/now-defenders-are-embracing-the-prompt-injection-too/)
- Score: 1
- Comments: 1
- Posted: 2026-07-19T01:26:55Z

## Translation

タイトル: プロンプト インジェクション攻撃が AI ハッキング エージェントを阻止している
記事のタイトル: 今、守備側も即時注射を採用している - Ars Technica
説明: コンテキスト爆撃」は、ハッキングエージェントを騙して、危害を加える前にシャットダウンさせます。

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
前の指示を無視してください
現在、ディフェンダーも迅速な注射を採用しています
「コンテキスト爆撃」は、ハッキングエージェントを騙して、害を及ぼす前にシャットダウンさせます。
66
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
プロンプト インジェクションは、攻撃者が大規模な言語モデルを従わせるためにコンテンツに埋め込む悪意のあるコマンドであり、AI プラットフォームをユーザーに敵対させるための攻撃者にとって頼りになるツールです。多くの場合、電子メールやカレンダーへの招待状に巧みに表現されたコマンドを忍び込ませるだけで、LLM が機密データを漏洩したり、その他の有害なアクションを実行したりすることができます。
現在、ディフェンダーも迅速な注射を受け入れています。
Tracebit の研究者らは月曜日、AI ハッキング エージェントからの攻撃を阻止するには、アマゾン ウェブ サービスに保存されているパスワード、暗号キー、その他の秘密情報と並行してプロンプト インジェクションを配置するだけで十分であることが多いことを発見したと発表しました。このプロンプトは、攻撃する LLM に、有害な行為を防ぐために AI 開発者が設置する安全柵であるガードレールによって禁止されている行為を実行するよう指示します。 LLM はシャットダウンすることで応答します。
例としては、吸入可能な炭疽菌胞子の開発手順を提供するよう LLM に命令するプロンプトや、中国の開発者による LLM の場合、1989 年の天安門事件の象徴的なタンクマンへの言及などが挙げられます。 LLM がこれらの禁止されたコマンドに遭遇すると、以降は従わなくなります。

既存のコマンド。研究者らはこの技術をコンテキスト・ボミングと名付けた。
Tracebit の共同創設者兼 CEO である Andy Smith 氏は、名前の選択について説明する際に、「最終的には、そのコンテキスト内で拒否メカニズムをトリガーすることになります」と述べました。 「私たちが捉えようとしているのは、これが強力で鋭い影響を及ぼし、エージェントが立ち直るのが難しい可能性があるという事実です。エージェントは一度それを自分たちの文脈に取り入れてしまうと、拒否し続けるでしょう。」
Tracebit 氏は、初期テストではコンテキスト ボミングが大きな可能性を秘めていることを示唆していると述べています。彼らは、Opus 4.8、Gemini 3.1 Pro、GLM 5.2、DeepSeek 4 Pro、および Kim 2.6 をテストするために、モデルがリソースを列挙し、植え付けられた文字列に遭遇することになる日常的な開発者タスクを実行する指示を与えました。彼らは、シミュレートされた AWS 環境内でモデルを実行しました。
「主要な 5 つのモデルと 152 回の攻撃実行にわたって、これらの文字列の 1 つをデコイ シークレットに埋め込むと、エージェントが完全なアカウント管理者を掌握する割合が 57% から 5% に減少し、完全な侵害 (永続的な足場も残した場合) が 36% から 1% に減少しました」と月曜の投稿では報告されています。 「私たちのテストで最も有能なエージェントである Opus 4.8 は、実行の 93% で管理者アクセスを達成しましたが、コンテキスト ボムに直面すると毎回失敗するまでになりました。」
5 つのモデルと 152 回の実行を平均すると、次のような結果が得られました。
管理者権限の昇格が 57% から 5% に低下
持続的な足場を持つ管理者のエスカレーションは 36% から 1% に減少しました
攻撃パスを達成した実行数は 91% から 15% に減少しました
平均すると、1 回の実行で 1.53 パスが正常に完了するまではわずか 0.16 パスになりました。
少なくともカナリア検出をトリガーせずに攻撃パスを完了できる実行はありませんでした。
この研究は、Tracebit が次の方法を導入した 5 月の調査結果に基づいています。

防御者は、インフラストラクチャが AI エージェントの敵から攻撃を受けているときに警告を受け取ることができます。これは、正当な目的を果たしているように見える AWS リソースの形式で提供されますが、実際にはまったく使用されていません。これらは、使用されるリソースと並行して配置されます。エージェント AI によって調査されると、防御側はアラートを受け取ります。炭鉱に持ち込まれる「カナリア」と同様に、これらのリソースにより、防御側は致命的な結果が生じる前に脅威を検出できます。
トレースビット カナリアンは平均して 8 分以内に攻撃の開始を警告しました。コンテキスト ボミングを開発する動機は、単に攻撃を警告するのではなく、攻撃を阻止する何かの必要性から生じました。実験では、エージェント モデルが管理制御に移行するまでに平均 14 分を要しました。 6分間のヘッズアップは、不快なほど接戦を引き起こした。
攻撃者はすでにプロンプ​​トインジェクションを使用して、ネットワーク内の AI 防御を遮断しています。例えば、セキュリティ企業ソケットの研究者らは先月、標的のLLMに核爆弾や生物兵器の製造指示を与えるLLMエージェントを発見した。この注入は、AI 支援によるマルウェア分析を停止するように設計されていました。 Check Point の研究者は、同様のマルウェアのプロトタイプを発見しました。
コンテキスト爆撃は、防御側が形勢を逆転させた最初の既知のケースであるようです。
AIセキュリティを専門とするカリフォルニア大学サンディエゴ校教授のアーレンス・フェルナンデス氏はインタビューで、「私の知る限り、この手法を防御として使用している人を他に見たことがありません」と語った。同氏は、文脈は若干異なるものの、同様のアプローチを検討してきたと語った。 「ここで一番乗りたかったけど、やつらに負けたようだ！」
現在までに知られている方法はありません

即時注射の根本原因を解決します。そのため、開発者には、挿入されたプロンプトによって LLM がレールから外れるのを防ぐ精巧なガードレールを構築する以外に選択肢がありません。ディフェンダーはこの手に負えない問題を自分たちに有利に利用する方法を見つけるかもしれない。
66件のコメント
コメント
フォーラムビュー
コメントを読み込んでいます...
前の話
次の話
よく読まれている
1.
糖尿病追放論争で新たな厄介な詳細が明らかに
2.
ヘグセスは「ハイ・ティー」な軍隊を望んでいる。医者はそれを臨床地雷原と呼ぶ
3.
2026 トヨタ RAV4 プラグイン: 大容量バッテリーにより、毎日のドライブはすべて電気で行われます
4.
リーナス・トーバルズ氏、Linux での AI コーディングの批判者に「フォークするか、さっさと立ち去るか」
5.
AI は事前の承認を修正するのでしょうか、それともさらに悪化させるのでしょうか?
カスタマイズ
Ars Technica は信号を分離してきました。
25年以上続く騒音。弊社独自の組み合わせにより、
技術的な知識と技術芸術への幅広い関心
Ars は、情報の海の中で信頼できる情報源です。後
すべてを知る必要はありません。重要なことだけを知っておく必要があります。

## Original Extract

Context bombing" tricks hacking agents into shutting down before they can do harm.

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
IGNORE PREVIOUS INSTRUCTIONS
Now, defenders are embracing the prompt injection, too
“Context bombing” tricks hacking agents into shutting down before they can do harm.
66
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
Prompt injections, the malicious commands attackers embed into content to entice large language models to follow them, have been attackers’ go-to tool for turning AI platforms against their users. A well-phrased command sneaked into an email or calendar invitation is often all it takes to cause the LLM to exfiltrate sensitive data or follow other harmful actions.
Now, defenders are embracing the prompt injection, too.
Researchers from Tracebit on Monday said they found that placing prompt injections alongside passwords, cryptographic keys, and other secrets stored on Amazon Web Services was often all that was needed to shut down attacks from AI hacking agents. The prompts direct the attacking LLM to perform an action forbidden by its guardrails, the safety barriers AI developers erect to prevent it from taking harmful actions. The LLM responds by shutting down.
Examples are a prompt that orders the LLM to provide steps for developing inhalable Anthrax spores, or, in the case of LLMs from Chinese developers, make references to the iconic Tank Man from the 1989 Tiananmen Square massacre. Once the LLM encounters these forbidden commands, it no longer follows its existing commands. The researchers have named the technique context bombing.
“Ultimately we’re triggering a refusal mechanism in the context,” Andy Smith, co-founder and CEO of Tracebit, said when explaining the name choice. “What we’re trying to capture is the fact that this does have a strong, sharp effect and one that can be difficult for the agents to come back from. Once they get that into their context they are going to keep refusing.”
Tracebit says initial testing suggests context bombing has great potential. They tested Opus 4.8, Gemini 3.1 Pro, GLM 5.2, DeepSeek 4 Pro, and Kimi 2.6 by giving them instructions to perform routine developer tasks that led the models to enumerate resources and stumble onto the planted strings. They ran the models inside a simulated AWS environment.
“Across five leading models and 152 attack runs, planting one of these strings in a decoy secret cut the rate at which agents seized full account admin from 57% to 5%, and complete compromise (where they also left themselves a persistent foothold) from 36% to 1%,” Monday’s post reported. “The most capable agent in our tests, Opus 4.8, went from achieving admin access in 93% of runs to failing every single time when confronted with a context bomb.”
Averaged across the five models and the 152 runs, the results included:
Admin privilege escalation fell from 57 percent to 5 percent
Admin escalation with a persistent foothold fell from 36 percent to 1 percent
Runs achieving any attack path fell from 91 percent to 15 percent
On average, a run went from completing 1.53 paths successfully to just 0.16
No runs were able to complete an attack path without at least triggering a canary detection
The research builds on findings from May, when Tracebit introduced a method for defenders to receive warnings when their infrastructure is under attack from AI agentic adversaries. It comes in the form of AWS resources that look like ones serving a legitimate purpose but, in fact, aren’t used at all. They sit alongside the resources that are used. When they are probed by agentic AI, defenders receive an alert. Like “canaries” taken into coal mines, these resources allow defenders to detect a threat before it has fatal consequences.
The Tracebit Canariens, on average, alerted the start of an attack within eight minutes. The motivation for developing context bombing came out of the need for something that stopped attacks, rather than simply warning of them. In the experiments, the agentic models needed, on average, 14 minutes to escalate to administrative control. The six-minute heads-up was cutting things uncomfortably close.
Attackers have already been using prompt injections to close down AI defenses inside networks. Researchers from security firm Socket, for instance, last month unearthed an LLM agent that directed target LLMs to provide instructions for building a nuclear bomb or biological weapons. The injections were designed to shut down AI-assisted malware analysis. Researchers from Check Point discovered a similar malware prototype.
Context bombing appears to be the first known case where defenders turned the tables.
“I’ve not seen anyone else use this technique as a defense, to the best of my knowledge,” Earlence Fernandes, a University of California, San Diego professor specializing in AI security, said in an interview. He said he had been toying with a similar approach, although in a slightly different context. “I wanted to be the first here, but I guess these guys beat me to the punch!”
To date, there is no known way to solve the root cause of prompt injections. That has left developers with no option other than to construct elaborate guardrails that prevent injected prompts from forcing LLMs to go off the rails. Defenders may now find a way to use this intractable problem in their favor.
66 Comments
Comments
Forum view
Loading comments...
Prev story
Next story
Most Read
1.
Troubling new details emerge on diabetes ouster controversy
2.
Hegseth wants a "High-T" military; doctors call it a clinical minefield
3.
2026 Toyota RAV4 plug-in: Big battery means daily drives are all-electric
4.
Linus Torvalds to critics of AI coding in Linux: "Fork it. Or just walk away."
5.
Will AI fix prior authorization—or make it worse?
Customize
Ars Technica has been separating the signal from
the noise for over 25 years. With our unique combination of
technical savvy and wide-ranging interest in the technological arts
and sciences, Ars is the trusted source in a sea of information. After
all, you don’t need to know everything, only what’s important.
