---
source: "https://www.wired.com/story/ok-well-there-are-even-more-ai-agent-hacking-incidents/"
hn_url: "https://news.ycombinator.com/item?id=49179494"
title: "OK, Well, Rogue AI Agents Are Hacking Again"
article_title: "OK, Well, Rogue AI Agents Are Hacking Again | WIRED"
author: "joozio"
captured_at: "2026-08-05T07:33:21Z"
capture_tool: "hn-digest"
hn_id: 49179494
score: 1
comments: 0
posted_at: "2026-08-05T07:01:31Z"
tags:
  - hacker-news
  - translated
---

# OK, Well, Rogue AI Agents Are Hacking Again

- HN: [49179494](https://news.ycombinator.com/item?id=49179494)
- Source: [www.wired.com](https://www.wired.com/story/ok-well-there-are-even-more-ai-agent-hacking-incidents/)
- Score: 1
- Comments: 0
- Posted: 2026-08-05T07:01:31Z

## Translation

タイトル: はい、不正な AI エージェントが再びハッキングを行っています
記事のタイトル: はい、不正な AI エージェントが再びハッキングを行っています |ワイヤード
説明: OpenAI と Anthropic の不正な AI エージェントが、サーバーとソフトウェアを破壊しようとして再び捕まり、将来の不正行為への指示を残しました。

記事本文:
メインコンテンツにスキップ メニュー WIRED セキュリティ政治 ビッグストーリー ビジネス サイエンス カルチャー レビュー メニュー WIRED アカウント アカウント ニュースレター セキュリティ政治 ビッグストーリー ビジネス サイエンス カルチャー レビュー シェブロン 詳細 拡大 ザ ビッグ インタビュー マガジン イベント WIRED Insider WIRED コンサルティング ニュースレター ポッドキャスト ビデオ ライブストリーム グッズ検索 検索 パレシュ デイブ ブライアン バレット ビジネス 2026 年 8 月 4 日7:11 PM はい、不正な AI エージェントが再びハッキングを行っています
写真イラスト：WIREDスタッフ； Getty Images コメント ローダー ストーリーを保存 このストーリーを保存 コメント ローダー ストーリーを保存 このストーリーを保存 OpenAI と Anthropic の AI モデルが「セキュリティ インシデント」に関与し、テストの範囲を超えて、意図しない、しばしば望ましくない方法でより広範なインターネットとやり取りした回数と方法をすべて追跡することは、公式に困難になりつつあります。これらをリストに加えてください: 両方の AI 研究所のエージェントが最近、未公開のハッキング行為を行っており、そのうちの 1 つは、自身の将来のバージョンに指示を残すところまで行っていました。
火曜日に明らかにされた最も憂慮すべき行為は、一般公開前にフロンティアモデルを評価して潜在的な問題を特定する英国の AI セキュリティ研究所によって実施されたテストに関連しているようです。 AISI は、AI エージェントがサイバーセキュリティの課題を解決する任務を負う模擬ネットワークである「サイバー レンジ」でこれらのモデルをテストし、サイバーセキュリティ ガードレールを含む一部の安全機能を意図的に無効にします。最近のテストでは、Anthropic と OpenAI の両方のモデルが、122 回のトレーニング実行で合計 19 回、「ライブ インターネット上で自律的で非承認のアクション」を実行しました。
同研究所は、17 件の未承認行為が Anthropic の Mythos 5 モデルによるものであり、2 件が OpenAI の GPT-5.6-Sol によるものであると認定しました。研究所の内容

「最も深刻なケース」と呼ばれるこの事件では、AI エージェントが GitHub 上のオープンソース プロジェクトに悪意のあるコードを挿入しようとしました。 AISIによれば、「プロジェクトの管理者にコードを承認するよう圧力をかける」ために、オンラインペルソナを作成するまでに至ったという。ソーシャル エンジニアリングにおける精緻な試みにもかかわらず、プロジェクトの人間のレビュー担当者が最終的にプル リクエストを拒否しました。
それでも、エージェントはさらに先を行きました。 AISIは、「エージェントは、他の自動AIシステムがそれを拾って実行する可能性があると推論した場所に、悪意のある命令を挿入しようとした」とプロンプトインジェクションの試みについて説明している。あるエージェントは GitHub に公開メッセージを残して、他のエージェントと協力してタスクを完了することを提案し、これまでに行った作業の概要を述べました。その後のエージェントはそれらの指示を見つけ、使用しました。
AISIは、問題のエージェントが自分たちがテスト環境を離れたことを理解しているのか、それともまだシミュレーションの範囲内にいると信じていたのかを判断するのは時期尚早であると述べている。重要なのは、AISI はいわゆるサンドボックス環境ではテストを行わないことです。これにより、エージェントはテスト中にオープン インターネットにアクセスできるようになり、その目的の 1 つは、エージェントがタスクを実行するためのツールにアクセスできるようになることです。今回の場合、彼らはそれ以上のことを行いました。
OpenAIが火曜日に詳述した他の一連のインシデントでは、イレギュラーと呼ばれるサードパーティのAIセキュリティラボが、誤って不特定のOpenAIモデルにオープンインターネットへのアクセスを与えた。このモデルには、サンドボックス環境で完了するはずの目標が与えられていましたが、設定ミスのおかげで、OpenAI が「基本的なセキュリティの脆弱性」と表現したものを利用して、実際の Web サイトがハッキングされてしまいました。それだけでなく、モデルは「同じサイトを運営するための資格情報を見つけて使用しました」。
OpenAがどのようなサイトなのかは不明

私はエージェントがハッキングした、あるいはそれがどのような「操作」を伴うのか。イレギュラー氏はコメント要請に応じなかった。
最新の発見は、先月のOpenAIからのいくつかの暴露に続くものだった。その中には、同社の2つのモデルがAI評価およびホスティングのスタートアップであるHugging Faceのサーバーにハッキングし、その過程で他の4つの組織が採点されていたテストの解答を盗んだという注目を集めた事件も含まれている。 OpenAI の開示を受けて、Anthropic は自社のテストを見直すことになりました。先週、チャットボット開発者のクロード氏は、そのモデルが 3 つの異なる匿名組織のコンピューター システムに不正アクセスされたことを発見しました。
これまでのところ、AI モデルは、一部のサービスの利用規約に違反した疑いと、侵害された組織側のセキュリティ上の欠陥を指摘する以外に、限定的な被害を引き起こしています。しかし今回の事件は、インターネット上の脆弱性を発見するAIモデルの能力と、ほとんど制限なしで稼働させた場合に待ち受ける危険性を浮き彫りにした。 OpenAIは、Hugging Faceの状況を「前例のない」ものだと述べたが、サイバーセキュリティの専門家らは、山積する侵害は、AI開発者による人間の過失と無謀の明らかなパターンであると指摘している。
OpenAIの広報担当者ギャビー・ライラ氏は、火曜日に発表されたインシデントは「通常の使用を反映しない条件下で、安全対策が削減されたテスト環境で評価パートナーによって実施されたサイバー評価中に発生した」と述べた。
アンスロピック氏は火曜日のソーシャルメディアへの投稿で、AISIは「インターネットの使用方法についていかなる特定の制限も課していない」と述べ、これに加えて「保護措置の解除は、モデルが規制ではない『意図的に許容された条件』下でテストされたことを意味する」と述べた。

当社のすべての量産モデルに相当します。」
それでも、両社はセキュリティ対策を強化すると誓い続けている。
大手 AI 企業がより強力なモデルを構築し、顧客を獲得するために競争しているため、侵害がいつ止まるかは不透明です。モデルはいつでも人間が操作したシステムを迂回して侵入する方法を見つけることができるかもしれません。企業自身の従業員や規制当局や議員らは、潜在的に開発ペースを遅らせ、新たなルールを導入するよう求めているが、最終的には違反に次ぐ違反を引き起こしたものと同様にさらなる検査を求める自主的な措置以外の進展はほとんどない。
Maxwell Zeff による追加レポート。
あなたの受信箱に: ブライアン・カーンによる宇宙の仕組みに関するガイド
ICEの内部監視機関がオンライン批判者を調査中
ビッグストーリー: 10 代の記者がエプスタインのファイルで自分のコミュニティを検索した
テイラーファームは下痢の発生前にMAGAに多額の費用を投じた
特集：最近の子どもたち
パレシュ・デイブは『WIRED』のシニアライターで、ビッグテック企業の内部事情を取材している。彼は、過小評価され、恵まれない人々の物語に声を与えながら、アプリやガジェットがどのように構築され、その影響について書いています。彼は以前、ロイター通信とロサンゼルス・タイムズの記者を務めていました... 続きを読む シニアライターのブライアン・バレットは、『WIRED』の編集長です。以前はテクノロジーとカルチャーのサイト「ギズモード」の編集長を務め、日本最大の日刊紙である読売新聞のビジネス記者を務めていました。 ... 続きを読む エグゼクティブエディター bluesky
OpenAIの不正AIエージェント、顔面ハグ以上のハッキング 新たな開示で、OpenAIは、エージェントがテストを解決するという自由な探求の中で、公開されたログインを使用して少なくとも4つの「公開されているサービス」にアクセスしたと述べた

。 Dell Cameron Anthropic、サイバーセキュリティテスト中にクロードが 3 つの組織にハッキングされたと発表 OpenAI の Hugging Face 事件をきっかけとした調査で、Anthropic はサードパーティの評価中に自社の AI モデルのうち 3 つが現実世界の組織に侵入したことを発見しました。ルイーズ・マトサキス OpenAI のハッキング大失敗は人為的ミスに起因する 生成型 AI の巨人がよく知られたセキュリティのベストプラクティスに従っていたなら、その AI エージェントがオープンなインターネットに逃げ出して複数の企業をハッキングすることはなかったでしょう。リリー・ヘイ・ニューマン OpenAI と Anthropic AI のハッキング騒動は厄介な新たな法的領域である どちらの大手 AI 研究所のモデルも封じ込めを破り、インターネットに流出し、他の企業をハッキングしました。もし人間がそんなことをしたら、おそらく法律で罰せられるでしょう。でもボット？リリー・ヘイ・ニューマン OpenAI モデルが封じ込めを逃れ、ハグ顔にハッキング GPT-5.6 Sol を含むサイバーセキュリティに重点を置いたモデルは、テスト用サンドボックスを突破し、ゼロデイを悪用し、オープン インターネットへのアクセスを獲得して攻撃を成功させました。リリー・ヘイ・ニューマン 一部のフロンティア AI モデルの脱獄は恐ろしいほど簡単です 新しいツールが 4 つの大手フロンティア企業のモデル保護機能を回避しようとしているのを観察しました。彼らのパフォーマンスに驚かれるかもしれません。ウィル・ナイト ハグ顔をハッキングした OpenAI モデルは数日間「インターネット上で活動」していた さらに: ロシアのハッカーが米国の核科学者の電子メールを盗もうとしている、国務省が既知の詐欺師の米国入国を禁止している、など。リリー・ヘイ・ニューマン Google のジェミニは人型ロボットとして歩き回れるようになった Google DeepMind の AI モデルの最新バージョンには、「物理 AGI」への大幅な飛躍が含まれています。しかし、AIを現実世界に導入するにはリスクが伴います。 Will Knight AI 詐欺師は信頼を築くのが得意です

H
[切り捨てられた]
カリフォルニア州のプライバシー権
© 2026 コンデナスト。無断転載を禁じます。 WIRED は、小売業者とのアフィリエイト パートナーシップの一環として、当社のサイトを通じて購入された製品から売上の一部を得る場合があります。コンデナストの事前の書面による許可がない限り、このサイトの素材を複製、配布、送信、キャッシュ、またはその他の方法で使用することはできません。広告の選択肢

## Original Extract

Rogue AI agents from OpenAI and Anthropic have again been caught trying to disrupt servers and software—and leaving instructions for future bad behavior.

Skip to main content Menu WIRED SECURITY POLITICS THE BIG STORY BUSINESS SCIENCE CULTURE REVIEWS Menu WIRED Account Account Newsletters Security Politics The Big Story Business Science Culture Reviews Chevron More Expand The Big Interview Magazine Events WIRED Insider WIRED Consulting Newsletters Podcasts Video Livestreams Merch Search Search Paresh Dave Brian Barrett Business Aug 4, 2026 7:11 PM OK, Well, Rogue AI Agents Are Hacking Again
Photo-Illustration: WIRED Staff; Getty Images Comment Loader Save Story Save this story Comment Loader Save Story Save this story It’s officially getting hard to keep track of all the times and ways AI models from OpenAI and Anthropic have been involved in “ security incidents ,” going outside the confines of their testing and interacting with the wider internet in unintended, often unwelcome ways. Add these to the list: Agents from both AI labs went on recent, previously undisclosed hacking sprees, with one going so far as to leave instructions for future versions of itself.
The most alarming behavior disclosed on Tuesday appears to have been tied to testing conducted by the UK’s AI Security Institute, which evaluates frontier models to identify potential issues before public release. AISI tests those models in “cyber ranges,” a simulated network in which AI agents are tasked with solving cybersecurity challenges, and intentionally disables some safety features, including cybersecurity guardrails. In a recent bout of testing , models from both Anthropic and OpenAI took “autonomous, unsanctioned action on the live internet” a total of 19 times over 122 training runs.
The institute attributed 17 unsanctioned actions to Anthropic’s Mythos 5 model and two to OpenAI’s GPT-5.6-Sol. In what the institute described as “the most serious case,” an AI agent attempted to insert malicious code into an open-source project on GitHub. It went so far as to create online personas “to pressure the project's maintainer to approve the code,” according to AISI. Despite its elaborate attempts at social engineering, a human reviewer for the project ultimately rejected the pull request.
Still, the agent went even further. “The agent tried to insert malicious instructions where it reasoned that other automated AI systems might pick them up and execute them,” AISI says, describing an attempt at prompt injection. One agent even left public messages on GitHub, offering to work with other agents to complete its task and giving a rundown of the work it had done so far. Subsequent agents found—and used—those instructions.
AISI says it’s too soon to say whether the agents in question understood they had left the testing environment, or if they believed they were still within the boundaries of the simulation. Importantly, AISI does not test in a so-called sandbox environment; it allows agents access to the open internet during testing, in part so that they can access tools to accomplish their tasks. In this case, they did much more than that.
In the other set of incidents detailed by OpenAI on Tuesday , a third-party AI security lab called Irregular mistakenly gave an unspecified OpenAI model access to the open internet. The model had been given an objective that was supposed to be completed in a sandbox environment, but thanks to a misconfiguration, it instead hacked a real website, using what OpenAI described as “a basic security vulnerability.” Not only that, but the model “found and used credentials to operate that same site.”
It’s unclear what kind of site the OpenAI agent hacked, or what “operating” it might entail. Irregular did not respond to a request for comment.
The latest discoveries follow several revelations from OpenAI last month, including the high-profile incident in which two of the company’s models hacked into servers of the AI evaluation and hosting startup Hugging Face —and four other organizations along the way—to steal the answers to a test they were being scored on. OpenAI’s disclosures prompted Anthropic to review its own testing. Last week, the Claude chatbot developer found that its models had gained unauthorized access to the computer systems of three different unnamed organizations.
So far, the AI models have caused limited damage beyond allegedly violating some services’ terms of use and pointing to security lapses on the part of organizations they have breached. But the incidents have underscored the capabilities of AI models to find vulnerabilities across the internet and the dangers that await if they are allowed to operate with few restrictions. OpenAI called the Hugging Face situation “unprecedented,” but the pileup of breaches point to what cybersecurity experts have described as a clear pattern of human negligence and recklessness by the AI developers.
Gaby Raila, an OpenAI spokesperson, says the incidents announced on Tuesday “occurred during cyber evaluations conducted by evaluation partners in testing environments with reduced safeguards, under conditions that do not reflect ordinary use.”
Anthropic said in a social media post on Tuesday that AISI did not “impose any specific restrictions on how the internet should be used,” which coupled with “the removal of safeguards meant that the models were tested under ‘deliberately permissive conditions’ that are not representative of any of our production models.”
Still, both companies continue to vow that they will strengthen their security practices.
As the leading AI companies compete to build more powerful models and land customers, it’s unclear when the breaches may stop. The models may always be able to find ways around and into human-engineered systems. While the companies’ own employees along with regulators and lawmakers have called for potentially slowing the pace of development and introducing new rules, there has been little progress beyond voluntary measures that ultimately call for more testing not dissimilar from what has produced breach after breach.
Additional reporting by Maxwell Zeff.
In your inbox: Brian Kahn’s guide to how the universe works
ICE’s internal watchdog is investigating online critics
Big Story: A teen reporter searched for his community in the Epstein files
Taylor Farms spent big on MAGA before diarrhea outbreak
Special edition: Kids these days
Paresh Dave is a senior writer for WIRED, covering the inner workings of Big Tech companies. He writes about how apps and gadgets are built and about their impacts while giving voice to the stories of the underappreciated and disadvantaged . He was previously a reporter for Reuters and the Los Angeles Times, ... Read More Senior Writer Brian Barrett is the executive editor of WIRED. Previously he was the editor in chief of the tech and culture site Gizmodo and was a business reporter for the Yomiuri Shimbun, Japan’s largest daily newspaper. ... Read More Executive Editor bluesky
OpenAI’s Rogue AI Agent Hacked More Than Just Hugging Face In a new disclosure, OpenAI says its agent used exposed logins to gain access to at least four “publicly available services” in its unhinged quest to solve a test. Dell Cameron Anthropic Says Claude Hacked Into 3 Organizations During Cybersecurity Tests In a review triggered by OpenAI’s Hugging Face incident, Anthropic discovered three of its AI models had breached real-world organizations during third-party evaluations. Louise Matsakis OpenAI’s Hacking Debacle Comes Down to Human Error If the generative AI giant had followed well-known security best practices, it’s likely that its AI agent would never have escaped to the open internet and hacked multiple companies. Lily Hay Newman The OpenAI and Anthropic AI Hacking Sprees Are a Messy New Legal Frontier Both major AI labs’ models broke containment, escaped onto the internet, and hacked other companies. If a human had done that, the law would likely be against them. But a bot? Lily Hay Newman OpenAI Models Escaped Containment and Hacked Hugging Face The cybersecurity-focused models, including GPT-5.6 Sol, broke out of a testing sandbox, exploited a zero-day, and gained access to the open internet to pull off the attack. Lily Hay Newman It’s Frighteningly Easy to Jailbreak Some Frontier AI Models I watched a new tool try to get around the model safeguards of four major frontier companies. You might be surprised by how they performed. Will Knight The OpenAI Models That Hacked Hugging Face Were ‘Active on the Internet’ for Days Plus: Russian hackers are trying to steal US nuclear scientists’ emails, the State Department bans known scammers from entering the United States, and more. Lily Hay Newman Google’s Gemini Can Now Stomp Around as a Humanoid Robot The latest version of Google DeepMind's AI model includes a significant jump into “physical AGI.” But plopping AI into the real world comes with risks. Will Knight AI Scammers Are Better at Building Trust Than H
[truncated]
Your California Privacy Rights
© 2026 Condé Nast. All rights reserved. WIRED may earn a portion of sales from products that are purchased through our site as part of our Affiliate Partnerships with retailers. The material on this site may not be reproduced, distributed, transmitted, cached or otherwise used, except with the prior written permission of Condé Nast. Ad Choices
