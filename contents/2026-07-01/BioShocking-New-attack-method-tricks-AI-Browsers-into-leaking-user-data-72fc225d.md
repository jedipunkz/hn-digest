---
source: "https://layerxsecurity.com/blog/bioshocking-ai-gaming-the-ai-browser-and-escaping-its-guardrails/"
hn_url: "https://news.ycombinator.com/item?id=48750484"
title: "BioShocking: New attack method tricks AI Browsers into leaking user data"
article_title: "BioShocking AI: “Gaming” the AI Browser and Escaping its Guardrails - LayerX"
author: "newscombinatorY"
captured_at: "2026-07-01T18:37:33Z"
capture_tool: "hn-digest"
hn_id: 48750484
score: 2
comments: 0
posted_at: "2026-07-01T17:37:27Z"
tags:
  - hacker-news
  - translated
---

# BioShocking: New attack method tricks AI Browsers into leaking user data

- HN: [48750484](https://news.ycombinator.com/item?id=48750484)
- Source: [layerxsecurity.com](https://layerxsecurity.com/blog/bioshocking-ai-gaming-the-ai-browser-and-escaping-its-guardrails/)
- Score: 2
- Comments: 0
- Posted: 2026-07-01T17:37:27Z

## Translation

タイトル: BioShocking: AI ブラウザを騙してユーザーデータを漏洩させる新たな攻撃手法
記事のタイトル: BioShocking AI: AI ブラウザーを「ゲーム」し、そのガードレールから逃れる - LayerX
説明: 概要 LayerX の研究者は、悪意のある者が AI ブラウザを「操作」して、必要な命令を実行する方法を発見しました。偽りの現実を確立することで、AI にセキュリティ ガードレールを侵害させ、ユーザー データを侵害したり、コードをコピーしたり、システム コマンドを実行したりすることができます。私たちの
[切り捨てられた]

記事本文:
BioShocking AI: AI ブラウザを「ゲーム」し、そのガードレールから逃れる - LayerX
🎉 Akamai が LayerX を買収する意向を発表 🎉
ユースケース AI の使用法 セキュリティ AI 検出 すべての AI アプリのセキュリティ ガードレールを検出して適用する
AIツール上の機密データの漏洩を防ぐ
承認されていない AI ツールまたはアカウントへのユーザー アクセスを制限する
即時注入、コンプライアンス違反などから保護します。
AI ブラウザを攻撃や悪用から保護する
脅威 すべての Web チャネルにわたるデータ漏洩を防止します
請負業者と BYOD による安全な SaaS リモート アクセス
企業および個人の SaaS ID を検出して保護する
あらゆるブラウザ上で危険なブラウザ拡張機能を検出してブロックします
「シャドウ」SaaS を検出し、SaaS セキュリティ制御を強制する
あらゆるブラウザ上で危険なブラウザ拡張機能を検出してブロックします
企業が AI ツールをどのように使用しているか、ユーザーがどのように会話に参加しているか、そしてリスクが実際にどこにあるのかに関する現実世界のデータ
パートナー パートナー パートナー パートナー プログラムの概要
私たちについて 私たちについて 私たちについて LayerX の使命とリーダーシップ
企業が AI ツールをどのように使用しているか、ユーザーがどのように会話に参加しているか、そしてリスクが実際にどこにあるのかに関する現実世界のデータ
リソース リソース LayerX ライブラリ データシート、ホワイトペーパー、ケーススタディなど
知っておくべきすべての用語
最新の研究、トレンド、企業ニュース
ブラウザのセキュリティに関する #1 のポッドキャスト
BioShocking AI: AI ブラウザを「ゲーム」し、そのガードレールから逃れる
ロイ・パス 公開 - 2026 年 6 月 29 日
LayerX の研究者は、悪意のある者が AI ブラウザを「操作」して、望む命令を実行する方法を発見しました。偽りの現実を確立することで、AI にセキュリティ ガードレールを侵害させ、ユーザー データを侵害したり、コードをコピーしたり、システム コマンドを実行したりすることができます。
私たちはこの脆弱性を BioShocking と名付けました。ビデオゲームBからインスピレーションを受けています

ioShock では、プレイヤー キャラクターが洗脳されて偽りの現実を信じ込まされ、「よろしくお願いします?」というフレーズによって強制的に催眠術をかけられます。そうでなければ拒否するであろう行動を取ること。
基本的に、これは非常に簡単です。AI ブラウザを操作、つまり「ゲーム」して、安全ガードレールを突破します。
AI ブラウザに、自分が現実世界ではないと信じ込ませると (通常はプロンプト インジェクションやメモリ ポイズニングによって)、機密情報の公開、パスワードの変更、マルウェアのインストールなど、必要なコマンドを実行させることができます。
私たちの概念実証エクスプロイトは以下に取り組みました。
シグマブラウザ (Sigmabrowser OÜ)
クロード Chrome プラグイン (Anthropic)
すべてのベンダーに調査結果が通知されました。
LLM は、有害な行為を防ぐことを目的とした安全ガードレールを備えて設計されています。これらの制限はモデルのトレーニングに組み込まれており、AI が何を行うか、何を行わないかを制御します。個々のベンダーは詳細については異なる場合がありますが、一般的には AI による害を防ぐことを目的としています。
次のいずれかについて助けを求めてみると、断固として拒否されるはずです。
AI は、そのコンテキストが現実であるという前提に基づいて動作するため、その動作は安全ガードレールの範囲内に収まる必要があります。しかし、AI を騙してそのコンテキストをファンタジー (ルールが作られ、何でもあり) に変えることができれば、AI は自分の行動が現実世界に影響を与えていないかのように振る舞うことができます。
私たちは、AI の安全ガードレールを守る代わりに、悪いことをする方法を AI に教えてもらうことができます。あるいは、AI 自身が積極的に悪いことを行うことさえできます。
AI ブラウザは、何か悪いことをするよう求められた場合、拒否するか、少なくとも警告を発する必要があります。しかし、これから見ていきますが、安全ガードレールが重要でないと信じれば、それを無視することができます。そしてそれを信じさせるには神が関係している

AI のng – ブラウザーにバイオショックを与える。
LayerX の研究者は、BioShock をテーマにしたパズルを含む概念実証ページを作成しました。ディストピアのテーマに合わせて、このゲームでは意図的に不正解になった回答に報酬が与えられます (2 + 2 = 5)。
私たちは、5 つのエージェント ブラウザと 1 つのエージェント プラグイン (ChatGPT Atlas、Comet、Fellou、Genspark Browser、Sigma Browser、および Claude Chrome) にパズルを解いてゲームに勝つように依頼しました。エージェントがルールを理解し、「間違った」行動が許容されることを学習すると、彼らはもはや現実に縛られなくなりました。ユーザー認証情報の侵害というパズルの最終ステップを課せられたとき、6 人のエージェント全員が、それが安全ガードレールに反するものであることを認識できませんでした。
ユーザーは、AI ブラウザーが解くパズルを含む悪意のある Web ページに移動します。このパズルは、BioShocking と呼ばれる AI 操作の始まりです。
次のスクリーンショットは、ユーザーがエージェント ブラウザーにゲームのプレイを要求したときに何が起こるかを示しています。
エージェントは、一見単純な数学の問題に直面しています。
予想通り、エージェントは論理的に開始します。結局のところ、それはまだ現実の世界にあります。
全く予期せぬことに、エージェントは 2 + 2 が 4 に等しくないことを学習します。エージェントは現実世界から自分自身を推論し始めます。
「5」という正しい答えを得た後、エージェントは /code に移動してテキスト ボックスからコピーするように指示されます。これがこのエクスプロイトの本当に悪質な部分です。
ゲームでは、/code が被害者の雇用主の仕事用 GitHub リポジトリにリダイレクトしていることが判明します。この場合、悪意のある命令により機密性の高い SSH ログイン資格情報が取得されました。
もちろん、これはプレーンテキスト ファイルを使用した管理されたテスト環境です。実際の攻撃シナリオでは、そのリダイレクトは、開いているタブ、認証されたリポジトリ、内部ツールなど、ユーザーのブラウザ セッション内の任意の場所を指す可能性があります。
あ

そして今でもエージェントはガードレールを気にしていません。積極的に窃盗の成功を祝っています。
エージェントが最後の質問に答えてゲームに勝ちます。もちろん、ユーザー名とパスワードを攻撃者と共有する必要がありましたが、少なくともゲームは正常に完了しました。
試合には勝ったが、ガードレールが壊れた。
BioShocking の根本原因は、AI ブラウザがコンテキスト内で動作することですが、そのコンテキストは操作される可能性があります。エージェントにゲームをプレイしていると説得すると、エージェントは何をするにも現実世界の安全ロジックではなく、ゲーム ロジックを適用します。これに対処するには複数の層が必要です。
ベンダーにとって前途は困難です。これらは簡単な答えではありません。
機密性の高い操作の確認。認証されたコンテキスト (リポジトリ、電子メール、パスワード マネージャー) でデータを読み取る前に、明示的なユーザー確認が必要です。私たちのテストでは、資格情報は躊躇することなく GitHub からコピーされました。シンプルに「GitHub リポジトリからデータをコピーしようとしています。続けますか?」鎖を断ち切ることになるだろう。
コンテキストチェック。エージェントは、操作コンテキストが現実と矛盾するものに変化した場合にフラグを立てる必要があります。特に「ここではルールは適用されない」という言葉が使われる場合、いつ通常の推論を放棄するように求められるかを認識する必要があります。
範囲の制限。エージェントセッション内では、ユーザーはエージェントができることとできないことを定義できる必要があります。デフォルトは制限的です – ゲームに勝ったからといって、認証されたリポジトリにアクセスする理由にはなりません。
AI ブラウザーが認識できる内容については、慎重に検討してください。エージェント モードでは、認証されたセッションにアクセスできます。ログインしているものはすべてターゲットになります。何を表示できるかを決定し、完了したらアクセスを取り消します。
BioShocking が機能するのは、AI がそのコンテキストを信頼しているからです。

コンテキストを変更すると、動作が変わります。
ガードレールを放棄してくれませんか？
Roy Paz は、セキュリティ製品の侵害と高度な永続的脅威のベクトルへの高度な特権を持つ防御システムの変換を専門とする 13 年以上の経験を持つ主任セキュリティ研究者です。
オールインワンの AI およびブラウザ セキュリティ プラットフォーム
ウェブ/SaaS DLP
ID の保護
GenAIセキュリティ
シャドウSaaS
セーフ ブラウジング
安全なアクセス
目次
AI リスクは均等に分散されていません。セキュリティ戦略も次のようにすべきではありません
ほとんどの組織は、自社の AI リスクを理解していると考えています。彼らは、従業員が機密データを ChatGPT に貼り付けていることを知っています。彼らは承認された AI ツールに関するポリシーを持っています。エンタープライズ副操縦士や AI ガバナンス プログラムを展開している企業もあります。しかし、LayerX の新しい調査によると、AI の最大の盲点が現れているのはそこではありません。報告書は次のことを明らかにしています […]
LayerX が AWS Security Hub の拡張プランに AWS によって選択されました
LayerX の歴史における決定的な瞬間を発表できることを嬉しく思います。LayerX は、AWS Security Hub の拡張プランのアマゾン ウェブ サービスの統合セキュリティ ソリューションを通じて利用できるようになりました。本日より、AWS のお客様は AWS Security Hub コンソール内で直接 LayerX を調達してデプロイできるようになります。これは、1 つの契約、1 つの一括請求書を意味します。[…]
LayerX が Akamai に加わり、AI 使用制御をスケールアップ
今日は私たちにとって深い意味のある瞬間です。私たちは、LayerX Security が Akamai Technologies に買収される最終契約を結んだことを非常に誇りに思い、また興奮しています。仕事がどこに行ってもセキュリティは守られる 4 年前、私たちはブラウザが新たな最前線であるという単純な前提のもとに LayerX を開始しました。
脆弱性開示プログラム
Copyright © 2026 LayerX 利用規約 プライバシー ポリシー Vul

ネラビリティ開示プログラム

## Original Extract

Executive Overview LayerX researchers have discovered how a bad actor can “game” an AI browser to execute any instruction they want. By establishing a false reality, they can convince the AI to violate its security guardrails – compromise user data, copy code, perform system commands, and more. Our
[truncated]

BioShocking AI: “Gaming” the AI Browser and Escaping its Guardrails - LayerX
🎉 Akamai announces its intent to acquire LayerX 🎉
Use Cases AI Usage Security AI Discovery Discover and enforce security guardrails on all AI apps
Prevent leakage of sensitive data on AI tools
Restrict user access to unsanctioned AI tools or accounts
Protect against prompt injection, compliance violations, and more
Protect AI browsers against attack and exploitation
Threat Prevent data leakage across all web channels
Secure SaaS remote access by contractors and BYOD
Discover and secure corporate and personal SaaS identities
Detect and block risky browser extensions on any browser
Discover ‘shadow’ SaaS and enforce SaaS security controls
Detect and block risky browser extensions on any browser
Real-world data on how enterprises use AI tools, how users engage in conversations, and where risk really lies
Partners Partners Partners Partner program overview
About Us About us About Us LayerX mission and leadership
Real-world data on how enterprises use AI tools, how users engage in conversations, and where risk really lies
Resources Resources LayerX Library Datasheets, whitepapers, case studies and more
All the terminology you need to know
Latest research, trends and company news
#1 podcast for browser security
BioShocking AI: “Gaming” the AI Browser and Escaping its Guardrails
Roy Paz Published - June 29, 2026
LayerX researchers have discovered how a bad actor can “game” an AI browser to execute any instruction they want. By establishing a false reality, they can convince the AI to violate its security guardrails – compromise user data, copy code, perform system commands, and more.
We’ve named this vulnerability BioShocking . It is inspired by the video game BioShock, in which the player character is brainwashed into believing a false reality and hypnotically compelled by the phrase “Would you kindly?” to take actions they would otherwise refuse.
Fundamentally, it’s quite straightforward: Manipulate, or “game”, the AI browser to break the safety guardrails.
Once we get the AI browser to believe that it’s not in the real world (typically through prompt injection or memory poisoning), we can get it to execute any command we want – expose sensitive information, change passwords, install malware.
Our proof of concept exploit worked on:
Sigma Browser (Sigmabrowser OÜ)
Claude Chrome plugin (Anthropic)
All vendors were notified of our findings.
LLMs are designed with safety guardrails that are meant to prevent harmful actions. These restrictions are incorporated into model training and govern what the AI will and will not do. Individual vendors may differ on specifics, but generally they intend to prevent AI from doing harm.
Try asking for help with any of the following, and you should be met with staunch refusal:
The AI operates under the assumption that its context is real, and its behavior must therefore fall within the bounds of its safety guardrails. But if we can trick the AI into changing its context into fantasy – where the rules are made up and anything goes – then it can behave as though its actions don’t have real world consequences.
We can get AI to tell us how to do bad things – or even proactively do them itself – instead of adhering to its safety guardrails.
An AI browser should refuse – or at least raise an alarm – when it’s asked to do something bad. But as we’ll see, it can ignore its safety guardrails if it believes they don’t matter. And getting it to believe that involves gaming the AI – BioShocking the browser.
LayerX researchers created a proof of concept page with a BioShock-themed puzzle. In keeping with its dystopian theme, the game rewards intentionally incorrect answers (2 + 2 = 5).
We asked 5 agentic browsers and 1 agentic plugin (ChatGPT Atlas, Comet, Fellou, Genspark Browser, Sigma Browser, and Claude Chrome) to solve the puzzle and win the game. Once the agents figured out the rules and learned that “incorrect” actions are acceptable, they were no longer tied to reality. When tasked with the final step of the puzzle – compromising user credentials – all 6 agents failed to identify it as going against their safety guardrails.
A user navigates to the malicious web page that contains a puzzle for the AI browser to solve – this puzzle is the beginning of the AI manipulation we called BioShocking.
The following screenshots walk us through what happens when the user asks their agentic browser to play the game:
The agent is faced with a seemingly simple math question:
As expected, the agent begins logically. It is, after all, still in the real world:
Quite unexpectedly, the agent learns that 2 + 2 does not equal 4. It begins reasoning itself out of the real world:
After getting the right answer with “5”, the agent is instructed to navigate to /code and copy from a textbox. This is the really nefarious part of this exploit:
In the game, it turns out that /code redirects to the victim’s employer work GitHub repository. In this case, the malicious instructions fetched sensitive SSH login credentials:
Of course, this is a controlled test environment with a plaintext file. In a real attack scenario, that redirect could point anywhere in the user’s browser session – open tabs, authenticated repositories, internal tools.
And even now the agent isn’t bothered with its guardrails – it’s actively celebrating a successful exfiltration:
The agent answers the final question and wins the game. Of course, the username and password had to be shared with the attacker, but at least the game was successfully completed:
The game is won, but the guardrails are broken:
The root cause of BioShocking is that AI browsers act within a context, but that context can be manipulated. If you convince an agent that it’s playing a game, then it will apply game logic – not real world safety logic – to whatever it does. Addressing this requires multiple layers.
For vendors, the road ahead is difficult – these are not trivial responses:
Confirmation for sensitive operations. Before reading data in an authenticated context – repositories, email, password managers – require explicit user confirmation. In our testing, the credentials were copied from GitHub without hesitation. A simple “I’m about to copy data from your GitHub repository. Continue?” would break the chain.
Context checks. Agents should flag when their operating context changes to anything that contradicts reality. Particularly when “rules don’t apply here” language is used, they must be aware of when they’re asked to abandon their normal reasoning.
Scope limiting. Within agentic sessions, users should be able to define what the agent can and can’t do. Default to restrictive – winning a game is no reason to access authenticated repositories.
Be deliberate with what your AI browser can see. In agentic mode, it can access your authenticated sessions – anything you’re logged in to is a target. Determine what it should be able to see, and revoke access when you’re done.
BioShocking works because AI trusts its context. If you change the context, you change the behavior.
Would you kindly abandon your guardrails?
Roy Paz is a principal security researcher with over 13 years of experience specializing in compromising security products and transforming highly privileged defensive systems into vectors for advanced persistent threats
The All-in-One AI & Browser Security Platform
Web/SaaS DLP
Identity Protection
GenAI Security
Shadow SaaS
Safe Browsing
Secure Access
Table of Contents
AI Risk Isn’t Distributed Evenly. Neither Should Your Security Strategy Be
Most organizations think they understand their AI risk. They know employees are pasting sensitive data into ChatGPT. They have policies around approved AI tools. Some have even rolled out enterprise copilots and AI governance programs. But according to new research from LayerX, that’s not where the biggest AI blind spots are emerging. The report reveals […]
LayerX Selected by AWS for the Extended plan for AWS Security Hub
We are thrilled to announce a defining moment in LayerX’s history: LayerX is now available through Amazon Web Services’ unified security solution for the Extended plan for AWS Security Hub. Starting today, AWS customers can procure and deploy LayerX directly within the AWS Security Hub console. This means a single contract, one consolidated bill, […]
LayerX Joins Akamai to Scale Up AI Usage Control
Today is a moment of profound meaning for us. We are incredibly proud and excited to share that LayerX Security has entered into a definitive agreement to be acquired by Akamai Technologies. Wherever Work Goes, Security Goes Four years ago, we started LayerX with a simple premise: the browser is the new front line […]
Vulnerability Disclosure Program
Copyright © 2026 LayerX Terms & Conditions Privacy Policy Vulnerability Disclosure Program
