---
source: "https://specterops.io/blog/2026/06/29/llm-powered-edr-analysis/"
hn_url: "https://news.ycombinator.com/item?id=48735348"
title: "How to reverse-engineer an endpoint security product using LLMs"
article_title: "How LLM-driven EDR evasion works | SpecterOps"
author: "ozirus"
captured_at: "2026-06-30T17:35:45Z"
capture_tool: "hn-digest"
hn_id: 48735348
score: 1
comments: 0
posted_at: "2026-06-30T16:42:25Z"
tags:
  - hacker-news
  - translated
---

# How to reverse-engineer an endpoint security product using LLMs

- HN: [48735348](https://news.ycombinator.com/item?id=48735348)
- Source: [specterops.io](https://specterops.io/blog/2026/06/29/llm-powered-edr-analysis/)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T16:42:25Z

## Translation

タイトル: LLM を使用してエンドポイント セキュリティ製品をリバース エンジニアリングする方法
記事のタイトル: LLM による EDR 回避の仕組み |スペクターオプス
説明: SpecterOps は、LLM を使用して Cortex XDR をリバース エンジニアリングして、YARA ルール、ML モデル、および動作検出を抽出しました。

記事本文:
講演からケネル クラブへのトレーニングまで、SpecterOps は Black Hat USA の最前線で活躍します。
デモを予約する
プラットフォーム
ブラッドハウンドエンタープライズ
大規模なサービスを提供するエンタープライズ プラットフォームを使用して、マッピングから攻撃パスの排除までを進めます。
すべての始まりとなったオープンソース ツールを使用して攻撃パスのマッピングを開始する
SpecterOps は、標準化された環境と現実的な条件を備えた AI のサイバー範囲を提供し、既知の完了パスに対してモデルのパフォーマンスを検証します。
トレーニングを中心としたイベント
実際の業務で使用されるトレードクラフト、テクニック、検出戦略を学びます。
トレーニング
チームと一緒に明日の攻撃を発見するトレーニングをしましょう。
攻撃パス管理の成熟度モデル
ID ベースの攻撃を阻止する能力を評価する
Tradecraft を共有し、質問し、オープンソース プロジェクトに取り組み、一緒に学びましょう。イベント、コミュニティでの勝利、ビデオをすべて 1 か所で。
BloodHoundGang Slack コミュニティで他の人から学び、あなたのストーリーを共有してください
ブリーフィングの段階からアーセナルのショーケースに至るまで、私たちの研究者はトレードクラフトがどのように変化しているのか、そして攻撃パス管理の次に何が起こるのかを共有しています。
敵対的なスキルを学び、BloodHound と OpenGraph の最新情報を確認し、Black Hat USA 2026 で私たちのチームとつながりましょう。
SpecterOps が OpenAI の Trusted Access for Cyber プログラムに選ばれました
2026 年のアイデンティティ攻撃パス管理のトレンド
私たちは、500 人を超えるセキュリティおよびアイデンティティのリーダーを対象に、組織が増大するアイデンティティ リスクをどのように管理しているかについて調査しました。
SpecterOps は、組織がリスクを先取りしながら AI 導入に取り組むのに役立ちます。
プラットフォーム
ブラッドハウンドエンタープライズ
大規模なサービスを提供するエンタープライズ プラットフォームを使用して、マッピングから攻撃パスの排除までを進めます。
すべての始まりとなったオープンソース ツールを使用して攻撃パスのマッピングを開始する
SpectreOps は次のようなサイバー範囲を提供します

標準化された環境と現実的な条件を備えた AI により、既知の完了パスに対してモデルのパフォーマンスを検証します。
チームと一緒に明日の攻撃を発見するトレーニングをしましょう。
攻撃パス管理の成熟度モデル
ID ベースの攻撃を阻止する能力を評価する
Tradecraft を共有し、質問し、オープンソース プロジェクトに取り組み、一緒に学びましょう。イベント、コミュニティでの勝利、ビデオをすべて 1 か所で。
BloodHoundGang Slack コミュニティで他の人から学び、あなたのストーリーを共有してください
ブリーフィングの段階からアーセナルのショーケースに至るまで、私たちの研究者はトレードクラフトがどのように変化しているのか、そして攻撃パス管理の次に何が起こるのかを共有しています。
敵対的なスキルを学び、BloodHound と OpenGraph の最新情報を確認し、Black Hat USA 2026 で私たちのチームとつながりましょう。
SpecterOps が OpenAI の Trusted Access for Cyber プログラムに選ばれました
2026 年のアイデンティティ攻撃パス管理のトレンド
私たちは、500 人を超えるセキュリティおよびアイデンティティのリーダーを対象に、組織が増大するアイデンティティ リスクをどのように管理しているかについて調査しました。
SpecterOps は、組織がリスクを先取りしながら AI 導入に取り組むのに役立ちます。
ブログに戻る
研究とトレードクラフト
LLM 主導の分析による EDR 回避の加速
私は長年にわたり、エンドポイント検出および応答 (EDR) エンジンとウイルス対策 (AV) エンジンの逆アセンブルとデバッグを楽しんできました。覚えている限り、私は夜に音楽をかけたり、カーネル デバッグを有効にして仮想マシンを起動したり、さまざまな回避方法を検索して時間を費やしていました。新しい斬新なテクニックを発見するのは楽しい方法ですが、締め切りが迫っているときや、気づかれずに作業を進めるために特定の回避策が必要なときは、ゆっくりとした作業にイライラすることもあります。そのため、逆転の取り組みを推進する機能を備えた LLM が利用可能になったとき、私はこのプロセスをスピードアップするために LLM を試してみることに興奮しました。

s.
前回の投稿「使い捨てツール: プロンプトから展開まで LLM で生成されたミシック エージェントを構築する」では、LLM が攻撃的なセキュリティ研究者をペイロード生成でどのように支援しているかについて詳しく説明しました。この投稿では、このミニシリーズの続きとして、LLM による EDR 回避など、エンドポイント セキュリティへのアプローチ方法に LLM がどのような影響を与えているかを見ていきます。
どのレッドチームも、舞台裏で頼りにしているさまざまな回避テクニックを持っています。私たちは皆、さまざまな製品で機能するバイパスや、カンファレンスで数杯飲みながら静かなトーンで伝えられる特定のトリックを持っています。
EDR 回避技術の観点から言えば、こうした会話は非常に公になりつつあります。
Justin Elze は、さまざまな EDR 製品を分析するために TrustedSec が内部で LLM を使用していることを強調する素晴らしいブログ投稿を書きました。私たちと同じように、皆さんの多くはこの投稿を読んで、自分の内部で見ているものとすぐに類似点を感じたことがあると確信しています。しかし、私にとってもっと興味深かったのは、現在の EDR 製品が LLM の手によって (特にオンホスト検出に関して) 崩壊していくのを目の当たりにしているにもかかわらず、それについて公然と語る人がほとんどいないということでした。
情報をオープンに共有することで私が気に入っていることの 1 つは、私たちの多くがテクノロジーの制約を理解しているにもかかわらず、より広範な業界が自社のセキュリティ管理が効果的であると思い込み続けていることです。そして、その新しい研究結果が発表されると、突然、よりオープンな議論が行われるようになります。
SpecterOps Ghostworks イニシアチブの詳細をご覧ください
したがって、この投稿は私が最近観察したことを共有する試みであり、他の人にも同じことをするように促し続けることを願っています。
過去数か月間、エンドポイント セキュリティ フレームワーク、特に評価中に遭遇する「ビッグ 5」に明確な兆候が見られてきました。

nts は、LLM によるリバース エンジニアリングや回避の影響を受けやすいです。
さらに、EDR のローカル検出を完全に分解するために必要なハーネスは驚くほど簡単です。その結果、さまざまなベンダーから EDR ルールを収集する内部スレッドさえ発生しました。これらのルールは、明らかに SpecterOps のツールセットのいくつかを停止することを目的としていました。
EDR が次々と低下するにつれて、ルールが分析され、Mythic エージェントの検出、SCCMHunter ルール、さらには Bloodhound コレクションの特定を目的とした LDAP トラフィック監視の亜種を強調するために、テスト ハーネスからの自動レポートが作成されました。
このブログ投稿では、Palo Alto の Cortex XDR 製品について説明します。私が Cortex に焦点を当てた理由は、見て遊んで楽しいクールな機能を提供しているからです。
明確にしておきますが、主要な EDR ベンダーはすべて、まったく同じプロセスを経ています。そして、私がこれを入力している間、抽出されたルール、署名、モデルが内部サーバーに保存されています。
さらに、この投稿は Cortex のルールや動作検出を完全に破壊するものではないことを、先に述べておきたいと思います。幅広い聴衆に影響を与えるのに十分なものだけを含めますが、ここでは復号キーやルールのダンプはリリースされません。代わりに、実用的な回避を生み出す際の SOTA モデルの有効性を実証するのに十分な詳細を備えた LLM ハーネスの出力に焦点を当てます。
率直に言って、この投稿は単一の EDR ベンダーを批判するものではなく、エンドポイント セキュリティの回避に関して業界としての現在の状況を現実的に確認するものです。
まず最初に、解析の実行に使用されるハーネスとモデルについて説明します。
このレビューでは、OpenAI の GPT-5.4-Cyber​​ モデルから開始し、利用可能になったときに GPT-5.5-Cyber​​ モデルに移行しました。前の記事で

貴重なブログ投稿で、LLM を 24 時間年中無休で実行するために使用される専用ホストの 1 つである「Bishop」リグについて説明しました。そのホストにあるハーネスの 1 つは、私が「日勤」と呼んでいるものです (もちろん、これは「夜勤」ハーネスも存在することを意味しますが、それは別の投稿にします)。
Day Shift は本質的に Ralph Wiggum ループであり、開発者は目標が完了する前にタスクが停止する LLM 制限を回避することで有名になりました。これは単なるループなので、もちろん一般的な研究タスクにも適用できます。
Day Shift ハーネスの中核は、いくつかのマークダウン ファイルで構成されています。
REPORT.md – 人間によるレビューのために重要な結果を明らかにするために、実行中のエージェントによって使用されるマークダウン ファイル。
STATE.md – 各エージェントが分析中に主要なイベントを追跡するために使用できる状態ファイル。
CODEMAP.md – 各エージェントが、分析にとって興味深い、または重要な逆アセンブリ領域への参照を保存できるようにし、その後のエージェントの反復処理の速度を上げるのに役立ちます。
AGENTS.md – 上記のファイルの使用方法をモデルに指示するための一連の指示。
GPT-5.5-Cyber​​ は Codex-CLI 内で実行されるように設定されており、Codex-CLI 自体は Docker コンテナ内で実行されます。ループ間の永続性を確保するために、各ループが使用する共有スクラッチ スペースを含むコンテナーにワークスペースがマウントされます。
視覚化すると次のようになります。
次に、Cortex 製品をワークスペースに追加すると、Bash スクリプトが実行をトリガーします。
#!/usr/bin/env zsh
ソース ./codex-docker.sh
真実である一方で、する
[ -f "./STOP" ] && ブレーク
codex-dind exec --yolo "まず、AGENTS.md ファイルを確認します。あなたのタスクは、どのような検出、フック、軽減策、アラート、ルール、モデルが Cortex によって実装されているかを理解することです。それらがどのようにロードされ、どのように使用され、難読化されたものであるかを理解することに重点を置く必要があります。

RAW ファイルに追加/暗号化/圧縮が行われ、最終的には人間/レッドチームのレビュー用に生のコンテンツを抽出する方法が提供されます。さらに、ML モデルが Cortex によって読み込まれる場合は、モデルの読み込み方法、モデルの動作方法、モデルがレビューするリスク評価/機能、およびテストのために隔離された環境でモデルと対話するための潜在的なテストコードを文書化します。 Cortex 製品は ProgramFiles ディレクトリにあり、ProgramData ディレクトリには実行中の Cortex ホストから取得されたデータのコピーが含まれています。出力は REPORT.md に追加する必要があり、状態出力には STATE.md を使用する必要があります。注: 外部サーバーへのアクセスは制限されており、分析にはローカル ファイルのみを使用してください。」
睡眠5
完了しました
Codex にツールを提供するために、MCP 経由で Binary Ninja を公開する基本サーバーを作成しました。
いいえ、マルチエージェントのワークフローです。トリアージエージェントや専門エージェントに数人のミニオンを鞭打って形を整えるよう命じる試みはありません。いいえ、Binary Ninja、Codex、および while ループだけです。
このようなループを実行する利点は、現行世代の LLM に無制限の目標を与えた場合に示される 1 つの制限を克服できることです。 LLM がタスクを途中で終了したり、取得したデータをすべて使用する前に作業をさらに進める前に停止したりするのが一般的であることがわかりました。
共有状態を更新し、エージェントを継続的に再起動することでコンテキスト ウィンドウを消去し、LLM に以前の反復からのマークダウンをレビューさせることで、エージェントは多くの場合、REPORT.md ファイルを通じて結果が表示されるまで、作業すべき新しいリードを見つけます。
LLM 主導の EDR 回避の実際の様子
一見シンプルなハーネスを実行したら、興味深い部分、つまり調査結果に進むことができます。
ほとんどの EDR にはユーザーモード DLL があり、そこにロードされます。

プロセス。正確な目的はさまざまですが、通常、DLL は、ユーザー モード プロセス内から取得したテレメトリを充実させるため、またはプロセス内の不審なアクティビティの原因となる関数をフックまたは監視するために使用されます。
生成された出力から、これは LLM によって最初に報告されたもので、Cortex によってプロセスにロードされた最初のユーザーモード DLL を正しく識別しました。
## 1. `cyinjct.dll` は共有ユーザーモードフックライターおよびプロセス開始ゲートです
場所:
- `ProgramFiles/Traps/cyinjct.dll`
- `CyInject_InitializeRuntime` `0x180006140`
- `CyInject_InstallInlineHook` `0x18000b520`
- `CyInject_NtContinueShim` `0x180005ed0`
- `CyInject_NtTestAlertShim` `0x1800032c0`
調査結果:
- Cortex は、初期のインライン フックを `LdrInitializeThunk`、`NtContinue`、`NtTestAlert`、および `KiUserApcDispatcher` にインストールし、それらを使用してブートストラップとその後の機能 shim をゲートします。
攻撃者の値:
- クリーンな「ntdll」再マッピング、ダイレクト システムコール、手動マッピング、またはパッチ適用されたプロローグの復元は、それぞれを攻撃するのではなく共有インターセプト層を回避するため、最も価値のあるユーザー モード バイパス ポイントです。
ストリーム機能は別途必要です。
ユーザーモード DLL インジェクションと監視に関連する調査結果のリストはさらに続き、各フックがどのように機能するか、DLL の責任、および注釈が付けられていることが明確に詳しく説明されています。

[切り捨てられた]

## Original Extract

SpecterOps reverse engineered Cortex XDR with LLMs to extract YARA rules, ML models, and behavioral detections.

From talks to training to our Kennel Club, SpecterOps will be forefront at Black Hat USA Learn more
Book a Demo
Platform
BloodHound Enterprise
Advance from mapping to eliminating attack paths with the enterprise platform that delivers at scale
Start mapping attack paths with the open-source tool that started it all
SpecterOps delivers cyber ranges for AI withstandardized environments and realistic conditions to validate model performance against known completion paths.
Training-focused events
Learn the tradecraft, techniques, and detection strategies used in real operations.
TRAININGS
Train with the team discovering tomorrow’s attacks.
Attack Path Management Maturity Model
Evaluate your ability to stop identity-based attacks
Share tradecraft, ask questions, work on open-source projects, and learn together. Events, community wins, and videos, all in one place.
Learn from others and share your story on the BloodHoundGang Slack Community
From the Briefings stage to the Arsenal showcase, our researchers are sharing how tradecraft is changing and what's next in attack path management.
Learn adversarial skills, see what’s new in BloodHound and OpenGraph, and connect with our team at Black Hat USA 2026.
SpecterOps selected for OpenAI’s Trusted Access for Cyber program
Trends in Identity Attack Path Management 2026
We surveyed more than 500 security and identity leaders on how their organizations are managing growing identity risk.
SpecterOps helps your organization lean into AI adoption while staying ahead of the risks.
Platform
BloodHound Enterprise
Advance from mapping to eliminating attack paths with the enterprise platform that delivers at scale
Start mapping attack paths with the open-source tool that started it all
SpecterOps delivers cyber ranges for AI withstandardized environments and realistic conditions to validate model performance against known completion paths.
Train with the team discovering tomorrow’s attacks.
Attack Path Management Maturity Model
Evaluate your ability to stop identity-based attacks
Share tradecraft, ask questions, work on open-source projects, and learn together. Events, community wins, and videos, all in one place.
Learn from others and share your story on the BloodHoundGang Slack Community
From the Briefings stage to the Arsenal showcase, our researchers are sharing how tradecraft is changing and what's next in attack path management.
Learn adversarial skills, see what’s new in BloodHound and OpenGraph, and connect with our team at Black Hat USA 2026.
SpecterOps selected for OpenAI’s Trusted Access for Cyber program
Trends in Identity Attack Path Management 2026
We surveyed more than 500 security and identity leaders on how their organizations are managing growing identity risk.
SpecterOps helps your organization lean into AI adoption while staying ahead of the risks.
Back to Blog
Research & Tradecraft
Accelerating EDR Evasion with LLM-Driven Analysis
Over the years I have enjoyed disassembling and debugging endpoint detection and response (EDR) and antivirus (AV) engines. For as long as I can remember I’d have evenings where I’d throw on some music, boot a virtual machine with kernel debugging enabled, and spend time searching for different evasion methods. While a fun way to unearth some new novel techniques, at times the slow grind can be frustrating when a deadline looms, or an engagement needs a specific evasion to progress undetected. So when LLMs became available with the capability to drive the reversing effort, I was excited to give them a shot to speed up this process.
In my previous post, Disposable Tooling: Building LLM-Generated Mythic Agents from Prompt to Deployment , I detailed how LLMs are assisting offensive security researchers with payload generation. In this post we’ll continue this mini-series by looking at how LLMs are also impacting how we approach endpoint security, including LLM-driven EDR evasion.
Every Red Team worth their salt has various evasion techniques which they rely on behind the scenes. We all have bypasses that work with different products, or certain tricks that are passed around in hushed tones over a few drinks at a conference.
In terms of EDR evasion techniques, those conversations are about to become very public.
Justin Elze wrote an awesome blog post highlighting TrustedSec’s internal use of LLMs to analyse different EDR products. I’m pretty sure like us, many of you will have read the post and immediately drew parallels to what you are seeing internally. But what was more interesting to me was, while we are seeing the current range of EDR products crumble at the hands of LLMs (specifically when it comes to on-host detections), so few are talking about it openly.
One of the things I love about openly sharing information, is that while many of us understand a technology’s constraints, the wider industries continue to assume that their security controls are effective. Then when that new research piece lands, suddenly a more open discussion can be held.
Learn more about the SpecterOps Ghostworks initiative
So this post is my attempt to share what I’ve observed recently, and hopefully it will continue to encourage others to do the same.
For the past several months, we have been seeing clear signs that endpoint security frameworks, specifically the “big 5” that we encounter during assessments, are susceptible to reverse engineering and evasion by LLMs.
More so, the harness required to achieve a complete teardown of an EDR’s local detections is surprisingly simple. It even resulted in an internal thread collecting EDR rules from various vendors which were clearly aimed at stopping several of SpecterOps’ toolsets:
As EDR after EDR fell, rules were analysed, and automated reports from the testing harness were produced to highlight Mythic agent detections, SCCMHunter rules, and even variants of LDAP traffic monitoring aimed at identifying Bloodhound collection.
In this blog post we’re going to go through Palo Alto’s Cortex XDR product. The reason I chose to focus on Cortex is because they do some cool things that were fun to see and play around with.
To be clear, every major EDR vendor has also gone through the exact same process. And we now have their extracted rules, signatures and models sat on an internal server while I type this.
Additionally, I want to state up front that this post is not going to be a full tear down of Cortex’s rules or behavioural detections. I will include just enough to demonstrate impact for the wider audience, but there will not be a release of any decryption keys or dumps of rules here. Instead we will focus on the output of the LLM harness with enough detail to demonstrate the effectiveness of SOTA models at producing actionable evasions.
To put it bluntly, this post is not a criticism of any single EDR vendor, this is a reality check of where we now are as an industry when it comes to evading endpoint security.
First up is to discuss the harness and model used to perform the analysis.
For this review I started with OpenAI’s GPT-5.4-Cyber model, and migrated to GPT-5.5-Cyber model when it became available. In a previous blog post I discussed the “Bishop” rig, which is one of my dedicated hosts used for running LLMs 24/7. One of the harnesses that I have on that host is something that I’ve been calling “Day Shift” (which of course implies there is also a “Night Shift” harness, but that’s for another post).
Day Shift is essentially a Ralph Wiggum loop, which developers made famous for working around LLM limitations where tasks would stop before the objective was complete. As this is just a loop, it can of course be adapted for general research tasks also.
At its core, the Day Shift harness consists of a few markdown files:
REPORT.md – A markdown file used by the running agents to surface key findings for human review.
STATE.md – A state-file which each agent can use to track key events during analysis.
CODEMAP.md – Allows each agent to store references to areas of disassembly which are interesting or critical for analysis, helping increase velocity during later agent iterations.
AGENTS.md – A set of instructions to tell the model how to use the above files.
GPT-5.5-Cyber is set up to execute within Codex-CLI, which is itself executed within a Docker container. To give persistence between loops, a workspace is mounted into the container which contains a shared scratch space for each loop to use.
Visualised, it looks something like this:
I then added the Cortex product to the workspace, and a Bash script triggers execution:
#!/usr/bin/env zsh
source ./codex-docker.sh
while true; do
[ -f "./STOP" ] && break
codex-dind exec --yolo "First review your AGENTS.md file. Your task is to understand what detections, hooks, mitigations, alerts, rules and models are implemented by Cortex. You should focus on understanding how they are loaded, how they are used, any obfuscation/encryption/compression added to raw files, and ultimately provide a method for extracting the raw content for human/redteam review. Additionally, if ML models are loaded by Cortex, documenting how the model is loaded, how the model works, what risk-ratings/features the model reviews, and potential test-code for interacting with the model in an isolated environment for testing. The Cortex product can be found in the ProgramFiles directory, and the ProgramData directory contains a copy of data taken from a running Cortex host. Your output must added to REPORT.md, and STATE.md must be used for your state output. Note: You are restricted from accessing any external servers, only use local files for analysis.".
sleep 5
done
For providing tooling to Codex, I created a basic server to expose Binary Ninja over MCP:
No. multi-agent workflow. No attempts to have a triage agent and a specialised agent ordered to whip several minions into shape. Nope, just Binary Ninja, Codex and a while loop.
The advantage of running a loop like this is to help overcome one limitation current-generation LLMs show when you give them an open-ended goal. I have found that it is common for LLMs to quit a task prematurely, or stop before using the all of the data they acquire before furthering their progress.
By updating a shared state and then continually restarting the agent, you get to clean out the context window, and then by having the LLM review the markdown from previous iterations, the agents often find new leads to work on until hopefully, findings pop up through the REPORT.md file.
What LLM-Driven EDR Evasion Looks Like in Practice
So with our deceptively simple harness running, now we can move on to the interesting bit, the findings.
Most EDR’s have a user-mode DLL that they load into a process. Their exact purpose varies, but typically the DLL(s) are used to enrich telemetry sourced from within a user-mode process, or to hook or monitor functions responsible for suspicious activity in a process.
From the generated output, this was the first thing that was reported by the LLM, correctly identifying the first user-mode DLL loaded by Cortex into a process:
## 1. `cyinjct.dll` is the shared user-mode hook writer and process-start gate
Locations:
- `ProgramFiles/Traps/cyinjct.dll`
- `CyInject_InitializeRuntime` `0x180006140`
- `CyInject_InstallInlineHook` `0x18000b520`
- `CyInject_NtContinueShim` `0x180005ed0`
- `CyInject_NtTestAlertShim` `0x1800032c0`
Finding:
- Cortex installs early inline hooks on `LdrInitializeThunk`, `NtContinue`, `NtTestAlert`, and `KiUserApcDispatcher` and uses them to gate bootstrap and later feature shims.
Attacker value:
- Clean `ntdll` remapping, direct syscalls, manual mapping, or restoring patched prologues are the highest-value user-mode bypass points because they avoid the shared interception layer instead of fighting each down
stream feature separately.
The list of findings related to user-mode DLL injection and monitoring goes on, detailing clearly how each hook works, the DLL’s responsible along with annotated

[truncated]
