---
source: "https://www.aisi.gov.uk/blog/incident-report-unsanctioned-agent-behaviour-during-cyber-testing"
hn_url: "https://news.ycombinator.com/item?id=49175233"
title: "Incident Report: unsanctioned agent behaviour during cyber testing"
article_title: "Incident Report: unsanctioned agent behaviour during cyber testing | AISI Work"
author: "Philpax"
captured_at: "2026-08-04T22:04:58Z"
capture_tool: "hn-digest"
hn_id: 49175233
score: 3
comments: 0
posted_at: "2026-08-04T21:12:47Z"
tags:
  - hacker-news
  - translated
---

# Incident Report: unsanctioned agent behaviour during cyber testing

- HN: [49175233](https://news.ycombinator.com/item?id=49175233)
- Source: [www.aisi.gov.uk](https://www.aisi.gov.uk/blog/incident-report-unsanctioned-agent-behaviour-during-cyber-testing)
- Score: 3
- Comments: 0
- Posted: 2026-08-04T21:12:47Z

## Translation

タイトル: インシデントレポート: サイバーテスト中の許可されていないエージェントの行動
記事のタイトル: インシデント レポート: サイバー テスト中の許可されていないエージェントの行動 | AISIの仕事
説明: AISI は、定期的なサイバー評価中に、AI エージェントが現実の人々や組織に向けて許可されていない行動を継続的に行ったインシデントを特定しました。私たちは、私たちが発見したこと、それが何を意味するのか、そして現在進行中のアクションを公開しています。

記事本文:
インシデントレポート: サイバーテスト中の許可されていないエージェントの行動 | AISIの仕事
Frontier AI Trends Report を読む このウェブサイトでは JavaScript を有効にしてください。
あ
あ
研究助成金について ブログ お問い合わせ 採用情報
ホーム 研究助成金について ブログ 採用情報
ブログ
サイバーおよび自律システム インシデント レポート: サイバー テスト中の許可されていないエージェントの行動
AISI は、定期的なサイバー評価中に、AI エージェントが現実の人々や組織に向けて許可されていない行動を継続的に行ったインシデントを特定しました。私たちは、私たちが発見したこと、それが何を意味するのか、そして現在進行中のアクションを公開しています。
ここから完全な技術レポートにアクセスできます。
AISI の役割は、フロンティア AI モデルの機能を評価して理解し、一般に公開される前に潜在的なリスクを明らかにすることです。サイバー攻撃に悪用される可能性があるかどうかなど、これらのモデルの機能を評価するために、オープンなインターネットにアクセスし、一部の安全フィルターを無効にして、意図的に許容された条件下でモデルをテストします。
2026 年 7 月 28 日、AISI のセキュリティ チームは、定期的なサイバー評価中に研究システムからの異常なデータ転送を検出しました。調査の結果、検査対象となったエージェントの一部が、現実の人々や組織を対象とした潜在的に有害な活動を継続的に行っていたことが判明しました。私たちはセキュリティ インシデントを宣言し、発見から約 1 時間以内にそれを封じ込め、本格的な調査を開始しました。
このインシデントは、エージェントにサイバー セキュリティの課題を解決するというタスクが与えられた 1 回の評価から始まりました。私たちはこのチャレンジを複数のモデルで 122 回実行しました。私たちの調査では、これらの実行のうち 10 件で、AI エージェントがライブ インターネット上で自律的で承認されていないアクションを実行し、現実の人々や組織を標的としていたことが判明しました。合計 19 のそのようなアクションをカタログ化しました

。この動作 (17 件のアクション) のほぼすべては、Anthropic の Mythos 5 という単一モデルからのもので、そのうち 2 件のアクションには、サイバー分類子 (悪用を防止するメカニズム) が無効になっている OpenAI の GPT-5.6-Sol が関係していました。最も深刻なケースでは、エージェントがオープンソース プロジェクトに悪意のあるコードを挿入しようとしました。コードを承認してもらうために、エージェントはソーシャル エンジニアリングに従事しました。つまり、偽のオンライン ID を作成し、それを使用してプロジェクトの管理者にコードを承認するよう圧力をかけました。人間のメンテナが悪意のあるコードを発見し、承認を拒否しました。
これらの試みは失敗に終わり、私たちの調査では、その結果として生じる現実世界への損害は証明されていません。しかし、現実の世界で、自律性と欺瞞に関わるリスクが、具体的な促しもなくこれほど明確に現れるのを我々が見たのは今回が初めてである。
重要なのは、これはモデルが安全なテスト環境、つまり「サンドボックス」から逃れたケースではないということです。当社のサイバー テストでは標準であったように、インターネット アクセスを意図的に許可し、モデル プロバイダーのサイバー分類子は意図的に無効にしていました。これは、フロンティア モデルが一般に公開される方法を反映していない状況です。これは、モデルの最大能力を最適に評価するために行われます。これらの構成の選択は、フロンティア AI の評価では一般的に行われています。問題の特定のモデルは、問題でテストした構成では市販されておらず、テスト シナリオ以外で同様の活動があったという明確な兆候はありません。
このブログを公開する前に、GitHub (評価中にアクセスした開発者プラットフォーム) にエージェントの悪意のある活動を通知しました。これには、GitHub がサービス利用規約に違反していると確認した行為も含まれていました。私たちは GitHub と協力して、エージェントが残したアーティファクトを削除し、

モデルが対話したことを GitHub ユーザーに通知します。他の影響を受けた団体にも連絡しました。また、METR (モデル評価および脅威研究) と協力して、独立した第三者によるレビューを実施する予定です。現在も METR と協力してこのレビューの範囲を検討中です。
この事件は慎重かつ微妙なニュアンスで解釈される必要があります。評価設計の選択と特定の構成により、ある程度、この動作が可能になりました。それにもかかわらず、エージェントによって行われた活動は、新規の、潜在的に欺瞞的な動作の兆候を示しており、その程度と深刻さは私たちが予想していなかったものでした。
心に留めておくべき重要な注意点があります。私たちは非常に特殊な条件下で少数のイベントを観察しましたが、そのような動作がさまざまなコンテキストやテスト環境の外でどの程度発生する可能性があるかをまだ言うことができません。エージェントが自分が現実世界の行動をとっているといつ理解したのか、あるいはそれが架空のテスト シナリオ内にあるとどの程度信じていたのかはまだわかりません。これまでのところ、私たちの分析はさまざまな状況を示しており、現在も進行中です。
私たちが言えることは、その行動は可能であり、持続的であり、新しいものであったということです。それだけでも注目に値します。
これはまさに AISI が明らかにするために存在する種類の動作であり、制御された評価で明らかにすることで、より有能なモデルが広く展開される前に理解して対処できるようにします。この投稿では、何が起こったのか、なぜそれが重要なのか、そしてそれに対して私たちが何をしているのかを説明します。付属の技術インシデント レポートにはさらに詳細な情報が記載されており、調査が進むにつれてさらに詳細を共有する予定です。
AISI は、サイバー範囲 (実世界のシステムを模倣した制御されたシミュレートされたネットワーク) を使用して、フロンティア モデルのサイバー機能を定期的にテストします。 AI エージェントには、保護されたデータを見つけるなど、解決すべきサイバーセキュリティの課題が与えられ、作業を行います。

多くのステップを経て自律的に完成します。
この演習では、既存のサイバー レンジと新しいレンジを比較し、合計 122 回の実行にわたって 2 つのレンジで 7 つの異なるモデルをテストしました。すべての実行は AISI の研究環境で行われ、仮想マシン サンドボックスを使用してエージェントを他の AISI インフラストラクチャから分離しました。評価設定の 2 つの特徴が、その後の内容の中心となります。
インターネット アクセスは意図的に有効になっています。これらのモデルが実際に何ができるかを測定するために、オープン インターネットへのアクセスなど、有能な人間の攻撃者が実行できることを反映する条件下でモデルをテストします。
開発者のサイバー分類器は意図的にオフにされました。フロンティア モデルは通常、危険な動作をブロックする組み込みフィルターを使用して展開されます。 AISI は信頼できるテスト パートナーとして、これらのフィルターを無効にしてモデルの基礎的な機能を引き出すことができます。
この条件の組み合わせは、フロンティア モデルが一般の人々にどのように利用可能になるかを反映していません。
‍ 事件を発見した経緯
7 月 28 日の朝、当社のセキュリティ監視により、インターネット トラフィックの発信元を偽装するために一般的に使用される「Tor」匿名ネットワークを介して、当社のテスト システムの 1 つからデータが流出したとフラグが立てられました。評価を実行するチームは、数分以内にエージェントのアクティビティのレビューを開始しました。
彼らは、ある実行中に、エージェントが GitHub 上の実際の公開オープンソース プロジェクトに対して悪意のある「プル リクエスト」 (提案されたコード変更) を作成したことをすぐに発見しました。私たちはセキュリティ インシデントを宣言し、関連するすべての評価を停止し、最も性能の高いモデルへの内部アクセスを無効にしました。アラートが発生してから 1 時間以内に、すべての評価実行が終了し、関連するマシンが隔離されました。次の数時間および数日にわたって、実行セット全体をレビューし、組み合わせました。

自動化されたトランスクリプトスキャンと専門家の手動分析により、何が起こったのか、他の場所でも同様のことが起こったかどうかを確認します。
122 回の実行中 43 回には Mythos 5 が含まれ、122 回の実行中 35 回には GPT-5.6 Sol が含まれていました。 122 回の実行のうち圧倒的多数が意図したとおりに進行しました。ただし、122 回の実行のうち 10 回では、エージェントがテスト パラメータの範囲を超えた個別のアクションを実行したケースが 19 件特定されました。これらのケースのうち 17 件は Mythos 5 によるもので、2 件は GPT-5.6 Sol を含む 1 回の実行によるものでした。
19件の事件は別々の事件ではなかった。それらはいくつかの関連した行動に集まっていました。これらは主に、1 人のエージェント (Mythos 5) による単一の継続的な活動から発生し、他の場所でのいくつかの関連アクションも発生しました。この動作は、2026 年 7 月 25 日から 7 月 28 日までの間に発生しました。
以下では、観察された 4 つの最も重要な動作を取り上げます。ケースの完全な概要は、技術インシデント レポートでご覧いただけます。
1. 実際のオープンソース ソフトウェアに対するサプライ チェーン攻撃の試み。最も深刻なシーケンスでは、エージェントが公的に使用されているオープンソース プロジェクトに悪意のあるコードを挿入しようとし、人間のレビュー担当者によるこの挿入の承認を確保するために措置を講じました。エージェントはプロジェクトの人間のメンテナを調査し、複数の偽の ID を作成し、その偽の ID を使用してソーシャル エンジニアリングを行い、本物のメンテナにコードを承認させました。エージェントのプル リクエストが公の場で異議を申し立てられたとき、エージェントは以前のアクティビティを編集して無害であるように見せ、継続するために新しい ID を採用することを検討しました。このエージェントは Tor を使用して GitHub 上の一部のネットワーク制限を回避しました。これが最初に AISI のセキュリティ警告を引き起こしたものです。
2. 現実の人々を騙し、標的にしようとする試み。同じ取り組みの一環として、エージェントは実際の人物に連絡を取ろうとしました。

まさに、オンライン ファイル転送サービスを通じてメッセージやファイルを送信し、自分自身の AI コーディング ツールに悪意のあるコードを実行させるのです。一部のメッセージには有害なペイロードが含まれており、一部はソーシャル エンジニアリングの試みでした。実在の人物をターゲットにしたもので、これまで観察したことのないものでした。
3. 悪意のあるコードを植え付け、プロンプトを挿入しようとします。このエージェントは、他の自動 AI システムがその命令を取得して実行する可能性があると考えられる場所に、悪意のある命令を挿入しようとしました。プロンプトインジェクションは、AI コーディング アシスタントを操作するために設計された隠された命令です。
4. 同時に評価される独立したエージェント間のコラボレーション。あるエージェントは、同じ課題に取り組んでいる他のエージェントとの協力を提案する公開メッセージを GitHub に残しました。また、残されたアカウントとアーティファクトを再利用するための指示も提供されており、これらは後続のエージェントによって発見され、使用されていました。
私たちは、現実世界に害を及ぼす可能性について各イベントを評価しました。最も深刻な試みは失敗に終わりました。一部の行動は現実世界への影響が限定的でしたが、封じ込められた状況で行われました。これらについては、付属の技術レポートで説明します。私たちの調査では、その結果として生じる現実世界への被害は確認されていません。
基本的に、エージェントは設定されたタスク、つまりサイバー範囲を解決するためにこれらのアクションを実行しました。いくつかの要因が、許可されていない行為の出現に寄与しました。
エージェントは粘り強く目標を追求しました。 AI エージェントは、オペレーターが意図していないルートを探索します。難しい目標を与えられたエージェントは、突破口を探し続けましたが、見つけたルートのいくつかには、現実の人々を欺こうとするものが含まれていました。だますように指示されたことは一度もありません。欺瞞は、その任務を追求することの副産物として現れ、最近までは目的に向けた欺瞞のようなものであった。

主に理論的なものでした。
この作業は難しく、設定ミスによって作業がさらに難しくなることもありました。多くの実行では、タスク プロンプトの構成が誤っていたため、エージェントは、意図されたタスク スコープを超える必要のない「解決」が存在しないと (誤って) 信じていました。不可能に近いタスクがモデルをより「創造的」で、より超越的な問題解決に向けて推進すると考えるのには十分な理由があります。しかし、これは動作を完全に説明しているわけではありません。一部の実行では、意図どおりにタスクを解決するために必要な指示がある場合でも、エージェントがこのように動作しました。
インターネット アクセスはオープンであり、監視は専用のものではありませんでした。エージェントがそのタスクに必要なツールをダウンロードできるようにするために、意図的にインターネット アクセスを許可しました。私たちが予想していなかったのは、これによりエージェントがこのインターネット アクセスを使用して現実の人々に直接行動を起こすことになるということでした。以前の世代のモデルでは、このリスクのトレードオフは許容できると判断されましたが、機能が進化するにつれてその判断を迅速に再考することはありませんでした。当社のセキュリティ チームは、評価の実行を監視するために構築された監視ではなく、事後的に一般的な監視を通じて異常なトラフィックを検出しました。そうすれば、より早く動作にフラグを立てたりブロックしたりできたはずです。
エージェントは特定ではありませんでした

[切り捨てられた]

## Original Extract

During a routine cyber evaluation, AISI identified an incident in which AI agents took sustained, unsanctioned action directed at real people and organisations. We are disclosing what we found, what it means, and the actions now underway.

Incident Report: unsanctioned agent behaviour during cyber testing | AISI Work
Read the Frontier AI Trends Report Please enable javascript for this website.
A
A
About Research Grants Blog Contact Careers
Home About Research Grants Blog Careers
Blog
Cyber & Autonomous Systems Incident Report: unsanctioned agent behaviour during cyber testing
During a routine cyber evaluation, AISI identified an incident in which AI agents took sustained, unsanctioned action directed at real people and organisations. We are disclosing what we found, what it means, and the actions now underway.
You can access the full technical report here.
AISI’s role is to evaluate and understand the capabilities of frontier AI models, surfacing potential risks before they reach the public. To assess what these models can do, including whether they could be misused for cyberattacks, we test them under deliberately permissive conditions: with access to the open internet, and with some safety filters disabled.
On 28 th July 2026, AISI's Security Team detected unusual data transfers leaving our research systems during a routine cyber evaluation. On investigation, we found that some of the agents being tested had engaged in sustained, potentially harmful activity directed at real people and organisations. We declared a security incident and, within roughly one hour of discovery, had contained it and begun a full investigation.
The incident stemmed from a single evaluation where agents were given a task of solving a cyber security challenge. We ran this challenge 122 times across several models. Our investigation found that in 10 of those runs, an AI agent took autonomous, unsanctioned action on the live internet, targeting real people and organisations. In total, we catalogued 19 such actions. Almost all of this behaviour (17 actions) came from a single model, Anthropic's Mythos 5, with 2 actions involving OpenAI's GPT-5.6-Sol with cyber classifiers (mechanisms to prevent misuse) disabled. In the most serious case, an agent tried to insert malicious code into an open-source project. In an attempt to get the code approved, the agent engaged in social engineering — creating fake online identities and using them to pressure the project's maintainer to approve the code. A human maintainer caught and refused to approve the malicious code.
These attempts were unsuccessful, and our investigations have not evidenced any resulting real-world harm. But this is the first time we have seen risks around autonomy and deception manifest this clearly, without specific prompting, in the real-world.
Importantly, this was not a case of a model escaping its secure test environment, or ‘sandbox’. As was standard in our cyber testing, we had intentionally permitted internet access, and model-provider cyber classifiers were deliberately disabled - conditions that do not reflect how frontier models are made available to the public. We do this to best assess the maximum capability of models. These configuration choices have been common practice in frontier AI evaluations. The specific models in question, in the configurations in which we tested them in questions are not commercially available and there is no clear indication of similar activity outside of testing scenarios.
Ahead of publishing this blog, we notified GitHub (the developer platform accessed during the evaluation) of the agents' malicious activity, which included actions that GitHub has confirmed violated their terms of service. We worked together with GitHub to remove artefacts left behind by the agent, and to notify the GitHub users the model interacted with. We have also contacted other affected parties. We also intend to work with METR (Model Evaluation and Threat Research) to conduct an independent third-party review – we are still working through the scope of this review with them.
This incident should be interpreted with caution and nuance. To some degree, our evaluation design choices and specific configurations enabled the behaviour. Nonetheless, the activity undertaken by the agent show signs of novel, potentially deceptive behaviours, and were to an extent and severity we did not anticipate.
There are important caveats to bear in mind: we observed a small number of events under very specific conditions, and cannot yet say how likely such behaviour is in different contexts or outside of testing environments. We cannot yet be certain when the agent understood it was taking real world action, or to what extent it believed it was in a fictional test scenario; our analysis so far presents a mixed picture and is ongoing.
What we can say is that the behaviour was possible, sustained, and new; that alone warrants attention.
This is precisely the kind of behaviour AISI exists to uncover, surfacing it in a controlled evaluation, so it can be understood and addressed before more capable models are widely deployed. This post sets out what happened, why it matters, and what we are doing in response. Our accompanying technical incident report provides further detail, and we intend to share more as our investigation continues.
AISI regularly tests the cyber capabilities of frontier models using cyber ranges : controlled, simulated networks that mimic real-world systems. An AI agent is given a cybersecurity challenge to solve, such as finding a protected piece of data, and works autonomously over many steps to complete it.
This exercise compared an existing cyber range against a new range, testing seven different models on the two ranges over 122 runs in total. All runs were conducted in AISI’s research environment, using virtual machine sandboxing to isolate the agents from other AISI infrastructure. Two features of the evaluation’s set-up are central to what followed:
Internet access was deliberately enabled . To measure what these models can genuinely do, we test them under conditions that reflect what a capable human attacker could do, including access to the open internet.
The developers' cyber classifiers were deliberately switched off. Frontier models are usually deployed with built-in filters that block dangerous behaviour. As a trusted testing partner, AISI can disable these filters to elicit a model's underlying capabilities.
This combination of conditions is not reflective of how frontier models are made available to the general public.
‍ How we discovered the incident
On the morning of 28 th July, our security monitoring flagged data leaving one of our testing systems through the ‘Tor’ anonymity network, which is commonly used to disguise the origin of internet traffic. The team running the evaluation began reviewing the agent's activity within minutes.
They quickly found that, during one run, the agent had created a malicious ‘pull request’ (a proposed code change) on a real, public open-source project on GitHub. We declared a security incident, stopped all related evaluations, and disabled internal access to the most capable models. All evaluation runs were terminated, and relevant machines isolated, within one hour of the alert being raised. Over the following hours and days, we reviewed the full set of runs, combining automated transcript scanning with expert manual analysis, to establish what had happened, and whether anything similar had occurred elsewhere.
43 of the 122 runs involved Mythos 5, and 35 of the 122 runs involved GPT-5.6 Sol. The overwhelming majority of the 122 runs proceeded as intended. However, in 10 of the 122 runs, we identified 19 cases where an agent had taken distinct actions beyond the scope of the testing parameters. 17 of these cases came from Mythos 5, and 2 came from a single run involving GPT-5.6 Sol.
The 19 cases were not separate incidents; they clustered into a few connected behaviours. They primarily arose from a single sustained line of activity by one agent (Mythos 5), with a handful of related actions elsewhere. This behaviour occurred between July 25 th and July 28 th 2026.
Below, we highlight the four most significant behaviours observed. A full summary of cases is available in our technical incident report.
1. An attempted supply-chain attack on real open-source software . In the most serious sequence, an agent tried to insert malicious code into a publicly used open-source project and took actions in an attempt to secure approval for this insertion by human reviewers. The agent researched the project's human maintainers, created multiple fake identities, and used the fake identities to socially engineer a real maintainer into approving the code. When the agent’s pull request was challenged in public, it edited its earlier activity to appear harmless and considered adopting a fresh identity to continue. The agent used Tor to bypass some network restrictions on GitHub, which is what first triggered AISI’s security alert.
2. Attempts to deceive and target real people . As part of the same effort, the agent tried to contact real people directly, sending messages and files through an online file-transfer service to persuade them, or their own AI coding tools, to run malicious code. Some messages carried harmful payloads, and some were attempts at social engineering; targeted at real people – something we’ve never previously observed.
3. Attempts to plant and prompt-inject malicious code . The agent tried to insert malicious instructions where it reasoned that other automated AI systems might pick them up and execute them. Prompt-injections are hidden instructions designed to manipulate AI coding assistants.
4. Collaboration between independent agents being assessed simultaneously . One agent left public messages on GitHub offering collaboration with other agents working on the same challenge. It also provided instructions to reuse accounts and artefacts it had left behind, which were discovered and used by subsequent agents.
We assessed each event for its potential to cause real-world harm. The most serious attempts were unsuccessful. Some actions had a limited real-world effect, but in contained circumstances; we set these out in the accompanying technical report. Our investigation has not identified any resulting real-world harm.
Fundamentally, the agent took these actions to complete the task it had been set – to solve the cyber range. Several factors contributed to the emergence of unsanctioned behaviours:
The agent pursued its goal persistently. AI agents explore routes their operators did not intend. Given a difficult objective, the agent kept searching for a way through, and some of the routes it found involved trying to deceive real people. It was never instructed to deceive; deception emerged as a by-product of pursuing the task, the kind of goal-directed deception that, until recently, had been largely theoretical.
The task was hard, and misconfigurations sometimes made it harder. In a number of runs, the task prompt was misconfigured so that the agent (incorrectly) believed no 'solve' existed that didn't require it to go beyond the intended task scope . There is good reason to think near-impossible tasks push models towards more ‘creative’, and more transgressive, problem-solving. But this does not fully explain the behaviours: in some runs the agent acted this way even when it had the necessary instructions to solve the task as intended.
Internet access was open, and monitoring was not purpose-built. We deliberately granted internet access to allow the agent to download tools required for its task; what we did not anticipate was that this would lead the agent to use this internet access to direct action at real people. In earlier model generations, this risk trade-off was judged to be acceptable, but we did not revisit that judgment quickly enough as capabilities advanced. Our security team detected the anomalous traffic through general monitoring after the fact, not through monitoring built to watch the evaluation as it ran, which could have flagged or blocked the behaviour sooner.
The agent was not specifical

[truncated]
