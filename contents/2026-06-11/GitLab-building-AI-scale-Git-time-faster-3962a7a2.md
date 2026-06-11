---
source: "https://about.gitlab.com/blog/gitlab-transcend-announcements/"
hn_url: "https://news.ycombinator.com/item?id=48494762"
title: "GitLab building AI scale Git time faster"
article_title: "GitLab: Built for the agentic engineering era"
author: "AsmodiusVI"
captured_at: "2026-06-11T19:04:59Z"
capture_tool: "hn-digest"
hn_id: 48494762
score: 2
comments: 0
posted_at: "2026-06-11T18:49:57Z"
tags:
  - hacker-news
  - translated
---

# GitLab building AI scale Git time faster

- HN: [48494762](https://news.ycombinator.com/item?id=48494762)
- Source: [about.gitlab.com](https://about.gitlab.com/blog/gitlab-transcend-announcements/)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T18:49:57Z

## Translation

タイトル: GitLab の AI スケール Git 時間を高速化する構築
記事のタイトル: GitLab: エージェント エンジニアリング時代に向けて構築
説明: GitLab Transcend で発表した内容と、企業が必要とする制御を 1 つのプラットフォームでエージェントのような速度で実現する方法。

記事本文:
閉じる リポジトリとプロジェクトを検索するには、 gitlab.com にログインします。提案 GitLab Duo エージェント プラットフォーム コードの提案 (AI) CI/CD AWS 上の GitLab Google Cloud 上の GitLab なぜ GitLab?プラットフォーム プラットフォーム DevSecOps のためのインテリジェントなオーケストレーション プラットフォーム プラットフォームを詳しく見る GitLab Duo エージェント プラットフォーム ソフトウェア ライフサイクル全体の Agentic AI GitLab Duo について GitLab を紹介する なぜ GitLab なのか 企業が GitLab を選ぶ主な理由を見る 詳細はこちら
ソフトウェア サプライ チェーンのセキュリティ
GitLab の新機能 最新の機能と改善点を常に最新の状態に保ちます。
営業担当者に問い合わせる 無料試用版を入手する サインイン 英語
GitLab Orbit はこちらです: AI エージェントのコンテキスト層。詳細はこちら ブログ AI GitLab: エージェント エンジニアリング時代に向けて構築 公開日: 2026 年 6 月 10 日
GitLab: エージェント エンジニアリング時代に向けて構築
GitLab Transcend で発表した内容と、企業が必要とする制御を 1 つのプラットフォームでエージェントのような速度で実現する方法について説明します。
当社のロードマップ、成功事例、業界調査を紹介する顧客イベントである GitLab Transcend が終了しました。私たちが発表し、デモした内容は次のとおりです。
次世代ソース コード管理、エージェント規模の同時実行のために再構築された Git エンジン、現在プライベート ベータ版
GitLab Orbit 、ソフトウェア ライフサイクル全体のコンテキスト グラフ、パブリック ベータ版になりました
エージェントのセキュリティとガバナンス、アイデンティティ、ポリシー、監査、あらゆるエージェントのアクションに関する承認のためのエージェント、現在プライベートベータ版
GitLab Duo Agent Platform は、開発者をフローから引き離すことなく、エージェントが問題を特定し、コードをレビューし、パイプラインを修正するオーケストレーション システムで、1 月から一般提供されます
GitLab Flex 、AI 導入が実際にどのように起こるかに応じた購入モデル、現在注文受付中
1,500 人を超える開発者と技術リーダーを対象とした当社の調査によると、エージェントを使用してコードを書くのがかつてないほど簡単かつ迅速になりました。

rs によると、組織の 91% が現在 2 つ以上の AI コーディング ツールを実行し、54% が 3 つ以上を実行しています。当社のお客様のコードベースの中には、1 年で最大 5 倍に成長しているものもあります。しかし、そのスピードを管理しないと混乱が生じます。
オンデマンド Web キャストで、トランセンドのあらゆる活動をご覧ください。
エージェントコーディング — 混沌としたスピードで
ほとんどの組織にとって、イノベーションのペースを高めるための答えは AI コーディングです。しかし、断片化されたソフトウェア ライフサイクルに基づいて構築されたコーディング エージェントはエージェント エンジニアリングではなく、非効率性とリスクへの近道です。
私たちはそれを繰り返しの形で顧客全体に見ています。人間のペースに合わせて構築されたインフラストラクチャは、ソース コード管理におけるマシン スケールの同時実行性によって影響を受けます。エージェントはコードに関するコンテキストを持っていますが、ソフトウェア開発ライフサイクル (SDLC) 全体を持っているわけではないため、大規模なモノリポジトリ内で停止し、コンテキスト ウィンドウがいっぱいになるとリポジトリ全体で諦めてしまいます。コード、依存関係、デプロイメントは、チームが管理できるよりも早く変化します。そして、固定契約により、チームは予測が最も困難な時代に、1年先の座席とクレジットを推測する必要があります。
同じ調査では、その負担が示されています。73% がコードの保守について心配しており、SDLC 全体で生産性が向上したと感じているのは 21% だけです。
これは速度を落とす理由にはなりません。これが、エージェント コーディングの速度を強力な ROI に変えるエージェント インフラストラクチャを構築する理由です。
エージェント コーディング + GitLab エージェント インフラストラクチャ、制御しながらスピードを実現
GitLab はすでに、企業がソフトウェアを構築して出荷するためのプラットフォームです。ソース コード管理 (SCM)、CI/CD、ガバナンス、展開などがすべて 1 か所に集約されており、当社のプラットフォームは現在 5,000 万を超えるユーザーと 100,000 の組織にサービスを提供しています。私たちは、コード、パイプライン、またはプロダクションに関わるすべてのソフトウェア チームとエージェントのパスに立っています。
考えてみましょう

人間が作業を行うかエージェントが作業を行うかに関係なく、同じように連携する 4 つの部分からなる生きたシステムとしてのプラットフォームです。モーター システムは、エージェント スケールでの実行、ソース管理、パイプライン、およびソフトウェアを構築して出荷する展開を処理します。神経系は、すべてのエージェントと人間に、適切な意思決定を行うためのコンテキストを与えます。免疫システムは、あらゆる行動にプロアクティブなガバナンスとセキュリティを適用します。
オーケストレーション システムは他の 3 つを調整し、チームとそのエージェントがライフサイクル全体にわたって作業を計画および実行できるようにします。仕事を行っている脳が人間であるかエージェントであるかにかかわらず、脳は 4 つすべてを活用します。
トランセンドでは、これらのイノベーションと新しい購入プログラムを発表し、それらが GitLab Duo Agent Platform によって実現されるインテリジェントなオーケストレーションとどのように連携するかを実証しました。
モーター システム: 同時実行のために再構築された次世代 SCM
Git バックエンドはエージェントの負荷によって座屈するため、SCM はここで注目する価値のあるプラットフォームの一部です。
Git は分散した従業員向けに構築されました。リポジトリ全体のクローンを作成して、世界中の何百人もの人々が並行して作業できるように十分なスケーラブルな分岐とコードのマージを使用して作業します。各チーム メンバーが何百ものエージェントを実行すると、このモデルは破綻し、次のような結果が生じます。
クローン税。リポジトリのクローンを作成して、エージェントごと、再試行ごとに 1 つのファイルを読み取ります。
同時実行性の崩壊。数千のセッションが一度に人間規模のバックエンドに到達し、ボトルネックと予測不可能な可用性が生じます。
孤立はありません。エージェントはアカウントと 1 つのブランチ スペースを共有するため、リポジトリが汚染され、作業を破棄するきれいな方法がなく、どのエージェントが何をしたかという記録も残りません。
そのため、私たちは次世代 SCM をプライベート ベータ版でプレビューしました。これは、下位互換性を確保するために Git プロトコル上で実行され、再設計されたバックエンドとエージェント用に構築されたインターフェイスを備えています。そうそう

任意のリポジトリでエージェントを実行し、何千ものエージェントをファンアウトして、安全に実験させることができます。
私たちのチームは、アーキテクチャの方向性を検証するために取り組んできました。つまり、同じ Git 互換性と監査可能性を備え、その下に別のモーターがあり、最初からエージェント アクセス用に設計されています。
初期の内部結果は有望です。トークンが最大 2 倍少なくなり、実時間は最大 50 倍速くなり、ネットワーク トラフィックが最大 1,000 倍少なくなります。
神経系: エージェントからの最も難しい質問に答える GitLab Orbit
エージェントはコードを書くことには熟達していますが、システム内をナビゲートすることについてはそれほど器用ではありません。彼らは自分が触れているコードは見ることができますが、ライフサイクル全体を見ることはできず、そのギャップが無駄な作業として現れます。
大規模なモノリポジトリでは、エージェントが書き込む反復とトークンが多すぎる可能性があり、応答を得るまでに時間がかかりすぎます。リポジトリ全体で、コンテキスト ウィンドウがいっぱいになりエージェントが諦めると、完全に失敗する可能性があります。その結果、正しいように見える作業が元に戻され、チームはエージェントが保存した時間よりもエージェントの出力を修正することに多くの時間を費やすことになります。
現在パブリック ベータ版の GitLab Orbit は、ソフトウェア ライフサイクル全体のコンテキスト グラフです。コード、作業項目、パイプライン、デプロイメント、および本番環境の信号を単一のライブ グラフに継続的にマッピングするため、エージェントは断片化された信号ではなくファーストパーティのコンテキストから推論します。エンジニアは、データ エクスプローラーを通じて同じグラフにクエリを実行して、変更を追跡し、インシデントを調査します。つまり、すべての決定は 1 つの真実の情報源から得られることになります。
当社の初期の社内テストでは、Orbit を使用して停止されたエージェントは、応答が最大 11 倍速く、費用対効果が最大 4.5 倍高く、幻覚が最大 45 倍少ないことがわかりました。
当社の顧客である Compare the Market は、同一のプロンプトとモデルを使用して、RAG ベースのアプローチとコンテキストなしのベースラインに対して GitLab Orbit の A/B テストを行い、次のことを発見しました。

経験豊富なエージェントは、70% (精度 0.696) の確率でインライン レビュー コメントを正しい場所に配置しましたが、RAG では 58% (0.577) でした。 Compare the Market は 79 件の実際のマージ リクエストでこのアプローチを検証し、コメント配置の精度ではグラフベースのコンテキストが従来の検索よりも優れたパフォーマンスを示しました。
続きを読む: GitLab Orbit の紹介: 完全なコードとライフサイクル コンテキストを 1 つのクエリで
Transcend ハカソンで Orbit 上に構築
Orbit を自分で動かしてみませんか? 2026 年 6 月 10 日から 24 日まで、Transcend ハッカソンではコミュニティにライフサイクル コンテキスト グラフを使用した構築を呼びかけます。 Orbit コードベースに直接貢献するか、その上にエージェント、フロー、スキルを構築して AI カタログに公開します。
経験豊富な投稿者から初めて投稿する人まで、誰でも大歓迎です。賞金も用意されています。さらに詳しく
免疫システム: セキュリティのためのエージェントとエージェントのためのガバナンス
エージェント時代では、各チームが管理するセキュリティとコンプライアンスのリスクは増大し続けています。エージェントがコードをより早く出荷すればするほど、それに付随する脆弱性もより早く出荷されます。このサイクルはさらに複雑になります。エージェントが生成するコードが増えれば増えるほど、より多くのカバレッジギャップが明らかになります。スキャンをすればするほど、トリアージして修正する必要がある結果が増えます。そして、運営するエージェントの数が増えるほど、リスクに対する姿勢を守るために、各エージェントが適切なポリシーの下で正しいことを行ったことを証明する必要がさらに増します。
このサイクルが、GitLab Ultimate を基盤として、セキュリティのためのエージェントを追加し、エージェントのガバナンスを導入した理由です。
セキュリティ担当エージェントは、脆弱性を作成するのと同じ速さで脆弱性をスキャン、優先順位付け、解決するため、エージェントが生成する量によって脆弱性の保護を担当するチームが埋もれることがありません。
エージェントのガバナンスは、プライベート ベータ版では信頼側で機能します。エージェントがコードをプッシュし、依存関係にタッチし、デプロイメントをトリガーするとき

規模が大きくなると、質問は「スキャンしましたか?」から変わります。 「どのエージェントがどの方針の下で何をしたか、そしてそれを証明できますか?」
エージェントのガバナンスでは、エージェントのあらゆるアクションに ID、ポリシー、監査、承認が適用され、組織全体の入力、推論、ツール呼び出し、高リスクまたは異常なアクティビティがリアルタイムで可視化されます。何か予期せぬことが起こった場合、何が変更されたか、どのポリシーがそれを許可したか、誰がサインオフしたかなど、完全な実行コンテキストと監査証跡がすでに存在します。
オーケストレーション システム: GitLab Duo エージェント プラットフォーム
上記の機能はインフラストラクチャです。 GitLab Duo Agent Platform は、開発者が毎日行っている作業にそれらを統合するオーケストレーション システムであり、導入が進んでいます。1 月に一般公開を発表して以来、週間アクティブ ユーザー数は 10 倍に増加しました。
ソフトウェア配信における摩擦は決して作業そのものではなく、開発者はコードの書き方、レビュー、パイプラインの修正方法を知っています。それは、コンテキストの切り替え、待機、および各ステップ間のハンドオフであり、時間が漏れて集中力が途切れてしまいます。
GitLab Duo Agent Platform では、エージェントはそのフル ループで作業します。つまり、広範囲にわたる問題を選択してマージ リクエストを開き、一般的なベスト プラクティスではなくチーム独自のルールに照らしてレビューし、返されたレビュー フィードバックを解決します。
品質は独立した監視に耐えます。 GitLab Duo Code Review は、Martian Offline Benchmark でスコア付けされた AI コード レビュー ツールのトップ 5 にランクインしました。また、これらのフローはイベント トリガーで自動的に実行できるため、複数のエンジニアをフローから引き離すことなく、コード変更が必要な中断と再実行のみが必要な中断をプラットフォームが区別することで、障害が発生したパイプラインを診断して修正できます。
GitLab Flex: 一度コミットすれば、シートの形状を変更して、

d AI支出
エージェント時代によりニーズの予測が難しくなり、ソフトウェアの購入方法も追いついていません。 6 か月後、必要なシートの数、チームが消費する AI の量、どの新機能を有効にするかはわかりませんが、従来の契約では 3 つすべてが事前に修正され、更新されるまで変更されることはありません。高く予想すると、空いている席の料金を支払うことになります。推測は低く、何か新しいものを採用するということは調達を再開することを意味します。
それは、現在利用可能な GitLab Flex が解決します。これは、シート、AI の使用、新機能にわたって月ごとに形を変える 1 つの年次コミットメントであり、すべて同じ契約に基づいており、修正や新しい調達サイクルはありません。チームがプロジェクトをロールオフすると、来月の予約を別のチームまたは AI に再割り当てします。署名後に機能が出荷されると、すでに持っているコミットメントに基づいて機能がオンになります。シートと使用量は 1 つのコミットメントの下で維持されるため、導入の変化に応じて予算を移動できます。 (Flex が構築する使用量ベースの基盤の背景については、GitLab クレジットの紹介を参照してください。)
続きを読む: GitLab Flex: 一度コミットすれば、シートと AI の支出を再構築できます
コーディングと出荷の違い
エージェントコーディングのみを取得します

[切り捨てられた]

## Original Extract

What we announced at GitLab Transcend, and how one platform delivers agentic speed with the control enterprises need.

Close To search repositories and projects, login to gitlab.com . Suggestions GitLab Duo Agent Platform Code Suggestions (AI) CI/CD GitLab on AWS GitLab on Google Cloud Why GitLab? Platform Platform The intelligent orchestration platform for DevSecOps Explore our Platform GitLab Duo Agent Platform Agentic AI for the entire software lifecycle Meet GitLab Duo Why GitLab See the top reasons enterprises choose GitLab Learn more
Software Supply Chain Security
What’s new in GitLab Stay updated with our latest features and improvements.
Talk to sales Get free trial Sign in English
GitLab Orbit is here: The context layer for AI agents. Learn more Blog AI GitLab: Built for the agentic engineering era Published on: June 10, 2026
GitLab: Built for the agentic engineering era
What we announced at GitLab Transcend, and how one platform delivers agentic speed with the control enterprises need.
GitLab Transcend, our customer event showcasing our roadmap, success stories, and industry research just wrapped. Here's what we announced and demonstrated:
Next-generation source code management , a Git engine rebuilt for agent-scale concurrency, now in private beta
GitLab Orbit , our context graph for the full software lifecycle, now in public beta
Agents for security and governance for agents , identity, policy, audit, and approval around every agent action, now in private beta
GitLab Duo Agent Platform , our orchestration system where agents pick up issues, review code, and fix pipelines without pulling developers out of flow, generally available since January
GitLab Flex , a buying model that bends to how AI adoption actually happens, now accepting orders
It's never been easier or faster to write code with agents, our research across more than 1,500 developers and tech leaders showed that 91% of organizations now run two or more AI coding tools, and 54% run three or more. Some of our customers' codebases are growing up to five times in a single year. But unmanaged, that speed creates chaos.
Catch up on all the action from Transcend with our on-demand webcast .
Agentic coding — speed with chaos
For most organizations, the answer to increasing the pace of innovation has been AI coding. But coding agents built on a fragmented software lifecycle isn't agentic engineering, it's a shortcut to inefficiency and risk.
We see it across our customers in recurring forms. Infrastructure built for human pace buckles under machine-scale concurrency in source code management. Agents have context around the code but not the full software development lifecycle (SDLC), so they stall out in large monorepos and give up across repos as the context window fills. Code, dependencies, and deployments change faster than teams can govern them. And fixed contracts force teams to guess at seats and credits a year ahead, in the era hardest to forecast.
The same research shows the strain: 73% worry about maintaining the code, and only 21% see productivity gains across the full SDLC.
None of this is a reason to slow down. It's a reason to build the agentic infrastructure that turns the speed of agentic coding into strong ROI.
Agentic coding + GitLab agentic infrastructure, delivering speed with control
GitLab is already the platform where enterprises build and ship software. Source code management (SCM), CI/CD, governance, deployment, and more all live in one place, and our platform currently serves more than 50 million users and 100,000 organizations. We sit in the path of every software team and the agents that touch their code, pipelines, or production.
Think of the platform as a living system, with four parts that work together the same way whether a human or an agent is doing the work. The motor system handles execution at agent scale, the source control, pipelines, and deployments that get software built and shipped. The nervous system gives every agent and human the context to make good decisions. The immune system puts proactive governance and security around every action.
And the orchestration system coordinates the other three, letting teams and their agents plan and carry out work across the full lifecycle. Whether the brain doing the work is human or agentic, it draws on all four.
At Transcend, we announced these innovations and a new buying program, and demonstrated how they come together with intelligent orchestration enabled by GitLab Duo Agent Platform.
Motor system: Next-gen SCM rebuilt for concurrency
SCM is the part of the platform worth focusing on here, because the Git backend buckles under agent load.
Git was built for a distributed workforce. We clone entire repositories to work on them, with branching and code merges scalable enough for hundreds of people around the globe to work in parallel. That model breaks when each team member runs hundreds of agents, resulting in:
Clone tax. Clone repo to read one file, for every agent and every retry.
Concurrency collapse. Thousands of sessions hit a human-scale backend at once, producing bottlenecks and unpredictable availability.
No isolation. Agents share accounts and one branch space, so they pollute the repo, leave no clean way to discard work, and no record of which agent did what.
That's why we previewed the next-generation SCM, in private beta . It runs on the Git protocol for backward compatibility, with a redesigned backend and interfaces built for agents. So you can run agents on any repo, fan them out by the thousands, and let them experiment safely.
Our team has been working to validate the architectural direction: same Git compatibility and auditability, with a different motor underneath, designed for agentic access from the start.
Early internal results are promising: up to 2x fewer tokens, up to 50x faster wall clock time, and up to 1,000x less network traffic.
Nervous system: GitLab Orbit to answer the toughest questions from agents
Agents are adept at writing code and far less deft at navigating the system around it. They can see the code they're touching but not the full lifecycle, and that gap shows up as wasted work.
In a large monorepo, agents can burn too many iterations and too many tokens, and they take too long to get an answer. Across repos, they may fail outright as the context window fills and the agent gives up. The result is work that looks correct but gets reverted, and teams spending more time fixing agent output than the agent saved.
GitLab Orbit, now in public beta , is the context graph for the entire software lifecycle. It continuously maps code, work items, pipelines, deployments, and production signals into a single live graph, so agents reason from first-party context instead of fragmented signals. Engineers query the same graph through the Data Explorer to trace changes and investigate incidents, which means every decision draws from one source of truth.
In our early internal tests, agents grounded with Orbit had up to 11 times faster response, were up to 4.5 times more cost effective, and had up to 45 times fewer hallucinations.
Our customer, Compare the Market, A/B tested GitLab Orbit against a RAG-based approach and a no-context baseline using identical prompts and models, and found: Graph-grounded agents placed inline review comments in the correct location 70% of the time (0.696 accuracy), vs. 58% (0.577) for RAG. Compare the Market validated the approach on 79 real merge requests, where graph-based context outperformed conventional retrieval on comment placement accuracy.
Read more: Introducing GitLab Orbit: Full code and lifecycle context, in one query
Build on Orbit at the Transcend hackathon
Want to put Orbit to work yourself? From June 10-24, 2026, the Transcend hackathon invites the community to build with the lifecycle context graph. Contribute directly to the Orbit codebase, or build agents, flows, and skills on top of it and publish them to the AI Catalog.
Everyone is welcome, from experienced contributors to first-timers, and cash prizes are on the table. Learn more
Immune system: Agents for security and governance for agents
In the agentic era, the security and compliance exposure every team manages keeps multiplying. The faster agents ship code, the faster they ship the vulnerabilities that come with it. The cycle compounds: the more code agents generate, the more coverage gaps you uncover; the more you scan, the more findings you have to triage and fix; and the more agents you run, the more you need to prove each one did the right thing, under the right policy, to stand behind your risk posture.
That cycle is why, building on GitLab Ultimate , we've added agents for security and introduced governance for agents.
The agents for security work the exposure side: scanning, triaging, and resolving vulnerabilities as fast as agents create them, so the volume agents generate doesn't bury the teams responsible for securing it.
Governance for agents, in private beta , works the trust side. When agents push code, touch dependencies, and trigger deployments at scale, the question shifts from "did we scan?" to "which agent did what, under which policy, and can we prove it?"
Governance for agents puts identity, policy, audit, and approval around every agent action, with real-time visibility into inputs, reasoning, tool calls, and high-risk or anomalous activity across the organization. When something unexpected happens, the full execution context and audit trail are already there: what changed, which policy allowed it, and who signed off.
Orchestration system: GitLab Duo Agent Platform
These capabilities above are infrastructure. GitLab Duo Agent Platform is the orchestration system that brings them together in the work developers do every day, and adoption has taken off: weekly active users have grown 10 times since we announced general availability in January.
The friction in software delivery was never the work itself, developers know how to write code, review it, and fix a pipeline. It's the context-switching, the waiting, and the handoffs between each of those steps, where time leaks out and focus breaks.
On GitLab Duo Agent Platform, agents work across that full loop: picking up a well-scoped issue and opening a merge request, reviewing it against the team's own rules rather than generic best practices, and resolving the review feedback that comes back.
The quality holds up against independent scrutiny. GitLab Duo Code Review placed in the top five of AI code review tools scored on the Martian Offline Benchmark . And because these flows can run automatically on event triggers, a failing pipeline can be diagnosed and fixed without pulling multiple engineers out of flow, with the platform distinguishing a break that needs a code change from one that just needs a rerun.
GitLab Flex: Commit once, reshape your seats and AI spend
The agentic era made your needs harder to predict, and the way you buy software hasn't caught up. Six months out, you don't know how many seats you'll need, how much AI your teams will consume, or which new capabilities you'll want to turn on, yet a traditional contract fixes all three up front and won't budge until renewal. Guess high and you pay for idle seats; guess low and adopting anything new means re-opening procurement.
That's what GitLab Flex, available now , solves. It's one annual commitment you reshape month to month across seats, AI usage, and new capabilities, all drawing from the same agreement, with no amendment and no new procurement cycle. When a team rolls off a project, you reallocate next month's reservations toward another team or toward AI. When a capability ships after you sign, you turn it on from the commitment you already have. Seats and usage live under one commitment, so you can move budget between them as adoption shifts. (For background on the usage-based foundation Flex builds on, see this introduction to GitLab Credits .)
Read more: GitLab Flex: Commit once, reshape your seats and AI spend
The difference between coding and shipping
Agentic coding only gets

[truncated]
