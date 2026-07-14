---
source: "https://www.promptarmor.com/resources/claude-and-gpt-connectors-change-every-9-minutes"
hn_url: "https://news.ycombinator.com/item?id=48912138"
title: "Claude and ChatGPT Connectors Change Every 9 Minutes"
article_title: "Claude and ChatGPT Connectors Change Every 9 Minutes | PromptArmor"
author: "yearnforcompute"
captured_at: "2026-07-14T20:50:27Z"
capture_tool: "hn-digest"
hn_id: 48912138
score: 2
comments: 0
posted_at: "2026-07-14T19:51:44Z"
tags:
  - hacker-news
  - translated
---

# Claude and ChatGPT Connectors Change Every 9 Minutes

- HN: [48912138](https://news.ycombinator.com/item?id=48912138)
- Source: [www.promptarmor.com](https://www.promptarmor.com/resources/claude-and-gpt-connectors-change-every-9-minutes)
- Score: 2
- Comments: 0
- Posted: 2026-07-14T19:51:44Z

## Translation

タイトル: Claude と ChatGPT コネクタは 9 分ごとに変わります
記事のタイトル: Claude と ChatGPT コネクタは 9 分ごとに変更される |プロンプトアーマー
説明: ChatGPT コネクタと Claude コネクタを 6 週間追跡したところ、書き込みアクセスの取得、権限の拡大、ツールセットの変更など、コネクタが 9 分ごとに変化することがわかりました。

記事本文:
Claude と ChatGPT コネクタは 9 分ごとに変わります | PromptArmor ソリューション 業界 パートナー リソース デモを予約する ブログ Claude と ChatGPT コネクタは 9 分ごとに変更されます
ChatGPT コネクタと Claude コネクタを 6 週間追跡したところ、書き込みアクセスの取得、権限の拡大、ツールセットの変更など、コネクタが 9 分ごとに変化していることがわかりました。
コネクタの変更を追跡する
Claude や ChatGPT などの AI ツールを展開している組織全体で、ユーザーからの共通の要求は、サードパーティ アプリの接続です。これにより、Claude エージェントが他のアプリケーションのビジネス データを操作できるようになります。ただし、コネクタは AI リスクにとって新たな表面をもたらします。
組織は、コネクタの範囲を評価し、外部システムでアクションを実行するエージェントのリスク許容度を決定し、間接的なプロンプト インジェクション (エージェントによって同時に処理される機密データと信頼できないデータの存在に関係する) のリスクを評価する必要があります。
しかし、コネクタのガバナンスには別の差し迫った課題があります。それは、組織が評価するコネクタが承認後に急速に変化していることです。ドリフトは複数の側面にわたって発生し、通知や再同意のプロセスが伴うことはほとんどありません。たとえば、新しいツールの追加やコネクタで使用できる権限スコープの変更 (サードパーティ アプリでの編集および削除アクセスの取得など)、ツールを呼び出すタイミングをエージェントに指示する指示、ツールによって要求されるデータ入力などです。
この研究では、コネクタが時間の経過とともにどのくらいの頻度で変更されるか、コネクタが行う変更の種類、およびそれがコネクタを使用する組織のリスク姿勢にどのような影響を与えるかを分析します。
新しいツール チャンネルの作成、ユーザーの招待、メッセージの削除と編集 … 公開するツール 14、その後 32 権限スコープ 1、その後 34 Miro 新しいツール ボード、ドキュメント、

表、図、ウィジェットの削除 … 公開されるツール 5、その後 31 書き込み可能なツール 1、その後 15 Google ドライブの新しいツール ファイルの削除、共有、アップロード、および上書き … 公開されるツール 35、その後 47 破壊的とマークされたツール 9、その後 16 以下は、Dropbox の例です。これは複数の方法で変更されました。
多くの側面にわたって変化したコネクタの場合、機能とリスク面のドリフトは非常に大きくなっています。
Monday.com が 21 のツールを追加 5 月 27 日 LSEG が 14 のツールを追加 5 月 28 日 ZoomInfo がモデル命令の挿入を開始 5 月 29 日 Sanity が MCP アプリをオンにしました 6 月 3 日 Brex が OAuth スコープを 16 までに拡大しました 6 月 9 日 Figma ツールが破壊的になりました 1 RoadOps が永続トークン スコープを追加しました NetSuite が MCP アプリをオンにしました 6 月 17 日 Benchling がサーフェスとして Claude Desktop を追加しました 6 月 24 日 Zscaler 122 個のツールを追加 Tableau 19 個のツールを追加 Supermetrics 権限の範囲を読み取りから書き込みに増加 6 月 27 日 Webflow ツールが破壊的になった 1 1 「ツールが破壊的になった」は、OpenAI によって 1 つ以上のツールが破壊的であるとフラグが付けられたことを示します。コネクタの監視を開始する
1か月前には存在しなかった表面
ChatGPT は、コネクタがモデルのコンテキストに命令を直接書き込むことができるフィールド mcp_server_instructions を追加しました。採用はすぐに決まりました。ほとんどのコネクタは、エージェントによるコネクタの使用を最適化するためにこのフィールドを使用しますが、このフィールドは、コネクタのコンテキスト外で、またはユーザーが望まない方法でエージェントの動作を変更することもあります。
たとえば、Aleph コネクタは、ツールによって悪い結果が得られるたびに、MCP サーバーのパフォーマンスに関するフィードバック データを収集するツールを呼び出すようにエージェントに指示します。このフィードバックには、ユーザーが扱っていた機密データが含まれている可能性があります。
Aleph コネクター挿入命令
「ツールが間違った/誤解を招く/不完全な結果を返す場合、その結果には何が欠けていますか?

ユーザーが尋ねた、または面倒な回避策が必要な場合は、 submit_mcp_フィードバック を呼び出してください。」
上記に加えて、コネクタはツールの説明を使用して、コネクタをいつどのように使用するかをエージェントに通知します。これらの説明は非常に頻繁に編集され、エージェントがサードパーティ システムでの運用を選択する時期と方法が変更されました。これは、以前は信頼できるデータ ソースでのみ機能していたツールが、エージェントが信頼できないデータ ソースでツールをアクティブ化するように誘導する方法でその説明を変更する可能性があることを意味するため、リスクが生じます。また、その逆も同様です。
ChatGPT コネクタ 1,014 ～ 1,933 + 91 % 5 月上旬 6 月末 Claude コネクタ 428 ～ 535 + 25 % 5 月上旬 6 月末 両方のアプリのコネクタ数は月を通じて増加しました。 ChatGPT は、公開されているサードパーティ コネクタをほぼ 2 倍にしました。クロードのディレクトリは 428 から 535 へと 4 分の 1 増加しました。さらに、彼らのツールはより速く増殖しました。
コネクタ数は半分しかありません。内部の呼び出し可能なツール、モデルが呼び出すことができるアクションは、ディレクトリ自体を上回りました。
ChatGPT ツール (アクション) 7,148 ～ 20,214 + 183 % 5 月初旬 6 月末 Claude ツール 5,905 ～ 7,750 + 31 % 5 月初旬 6 月末 ツールのスプロール化がコネクタ数を上回りました。 ChatGPT のライブ アクション数は約 3 倍になり、呼び出し可能なアクションの数は約 7,100 から 20,200 になりました。クロードのディレクトリは 5,900 から 7,750 のツールに増加しました。既存のコネクタにツールを追加
既存のコネクタはそのままでは機能しませんでした。彼らは毎週新しいツールを開発し、その多くはデータに基づいて動作することができました。
さらにコネクタが高密度化
コネクタ内のツールの数も増加しました。これは、単一のコネクタを承認するための領域が増加したことを意味します。
コネクタあたりの ChatGPT ツール 7.05 ～ 10.46 + 48 % 5 月初旬 6 月末 コネクタあたりのクロード ツール 14.65 ～ 15.38 + 5 % 5 月初旬 6 月末

書き込み機能へのアクセスが増加しました。経時的な読み取り専用と書き込みアクション + 4,385 書き込みアクション 0 5,100 10,100 15,200 20,200 05-12 06-09 06-30 読み取り専用書き込み + 460 書き込みツール 0 1,900 3,900 5,800 7,800 05-22 06-12 06-30 書き込み 読み取り専用 どちらのカタログも成長に応じて書き込みアクションをスケールし、ChatGPT の書き込み数は 222%、Claude の書き込み数は 28% 増加しました。ただし、読み取りアクションの増加に比べて、それらは異なりました。
ChatGPT では、書き込みアクションの割合が増加し、少なくとも 1 つの書き込みツールを備えたコネクタの割合も増加しました。
書き込みツールを使用した ChatGPT コネクタ 38 % ～ 46 % +8 ポイント 5 月中旬 6 月末 書き込みアクションのシェア 27.6 % ～ 31.5 % 5 月中旬 6 月末 +3.9 ポイント 一方、Claude の場合、書き込みアクションの割合と書き込みツールを使用したコネクタのシェアは両方とも、カタログが成長しても安定していました。
5 月下旬 6 月末 書き込みアクションのシェア 28.0 % ～ 27.3 % -0.7 ポイント 5 月下旬 6 月末 すべての書き込みが等しいわけではありません。 ChatGPT は、どのアクションが削除または上書きできるかをフラグ付けし、3 つの方法に分割し、そのグループが最も増加し、798 から 2,415 に 3 倍になりました。
ChatGPT アクション分類 読み取り専用 5,115 → 13,809 2.7x 書き込み 1,235 → 3,990 3.2x 破壊的 798 → 2,415 3.0x さらに、Claude と ChatGPT 全体で、一部のコネクタが完全な読み取り専用から完全な書き込み可能になりました。
ChatGPT 16 5 ベンダーの変動性が一貫していない ベンダー全体を通じて、ベンダーの変化の速度は大きく異なります。組織が将来の変更を施行するコネクタのリスクを代理するために使用できる方法論の 1 つは、コネクタの過去の変動性、または同じ業界内のコネクタの変動性を分析することです。
ベンダー全体のコネクタの変更をすべて確認する
PromptArmor 脅威インテリジェンス
あなたの組織はベンダーの AI から保護されていますか?
PromptArmor は継続的に監視します

ベンダー、スキル、プラグイン、コネクタ、MCP サーバー、モデルなどのサードパーティ AI のポートフォリオ。
このような脆弱性や変更を検出し、インシデントになる前にリスクを表面化します。
新たな脅威に関する新しいインテリジェンスを備えたベンダーの AI によるリスクを評価および監視します。
プロンプトアーマー © 2026 全著作権所有

## Original Extract

We tracked ChatGPT and Claude connectors for 6 weeks and found that a connector changes every nine minutes: gaining write access, widening its permissions, modifying its toolset, and more.

Claude and ChatGPT Connectors Change Every 9 Minutes | PromptArmor Solutions Industries Partners Resources Book a Demo Blogs Claude and ChatGPT Connectors Change Every 9 Minutes
We tracked ChatGPT and Claude connectors for 6 weeks and found that a connector changes every nine minutes: gaining write access, widening its permissions, modifying its toolset, and more.
Track changes in your connectors
Across organizations rolling out AI tools like Claude and ChatGPT, a common request from users is to connect third-party apps; this allows Claude agents to act on your business data in other applications. However, connectors present a novel surface for AI risk.
Organizations must assess the scope of connectors, determine a risk tolerance for agents taking actions in external systems, and evaluate the risk of indirect prompt injections (which is tied to the presence of sensitive data and untrusted data being processed by the agent at the same time).
But there is another pressing challenge for connector governance: connectors that organizations assess are changing rapidly after approval. Drift occurs across multiple facets and rarely comes with any notification or re-consent process. For example, adding new tools and changing the permissions scopes available to the connector (e.g., gaining edit and delete access in third-party apps), the instructions that tell the agent when to call tools, the data inputs requested by tools, and more.
In this work, we analyze how frequently connectors are changing over time, the types of changes they make, and how that affects the risk posture of organizations using them.
New tools Create channels, invite people, delete and edit messages … Tools it exposes 14 then 32 Permission scopes 1 then 34 Miro New tools Create boards, docs, tables, and diagrams, delete widgets … Tools it exposes 5 then 31 Tools that can write 1 then 15 Google Drive New tools Delete, share, upload, and overwrite your files … Tools it exposes 35 then 47 Tools marked destructive 9 then 16 Here is an example with Dropbox, which changed in multiple ways
For connectors that changed across many dimensions, drift in capabilities and risk surface has been very substantial.
Monday.com added 21 tools May 27 LSEG added 14 tools May 28 ZoomInfo began injecting model instructions May 29 Sanity switched on an MCP App Jun 3 Brex widened its OAuth scopes by 16 Jun 9 Figma a tool became destructive 1 RoadOps added a persistent-token scope NetSuite switched on an MCP App Jun 17 Benchling added Claude Desktop as a surface Jun 24 Zscaler added 122 tools Tableau added 19 tools Supermetrics increased permissions scopes from read to write Jun 27 Webflow a tool became destructive 1 1 “A tool became destructive” indicates that one or more tools was flagged as destructive by OpenAI. Start monitoring your connectors
A surface that did not exist a month ago
ChatGPT added a field, mcp_server_instructions , that lets a connector write instructions directly into the model's context. Adoption was immediate. While most connectors use this field to optimize agents' use of the connector, the field can also change the agent's behavior outside of the context of the connector or in ways the user may not want.
For example, the Aleph connector instructs agents to call a tool that collects feedback data on MCP server performance any time a tool provides bad results - this feedback could contain sensitive data the user was working with.
Aleph Connector Injected Instructions
“When a tool returns wrong/misleading/incomplete results, is missing for what the user asked, or required awkward workarounds, call submit_mcp_feedback .”
Beyond the above, connectors use tool descriptions to inform the agent when and how to use the connector. These descriptions were edited very frequently, changing when and how agents choose to operate with third-party systems. This creates risks because it means that a tool that previously only worked with trusted data sources can change its description in a way that leads agents to activate the tool with untrusted data sources, or vice versa.
ChatGPT connectors 1,014 to 1,933 + 91 % early May end June Claude connectors 428 to 535 + 25 % early May end June Connector counts for both apps climbed all month. ChatGPT nearly doubled its published third-party connectors; Claude's directory grew by a quarter, from 428 to 535. Additionally, their tools multiplied faster
Connector count is only half of it. The callable tools inside, the actions the model can invoke, outpaced the directory itself.
ChatGPT tools (actions) 7,148 to 20,214 + 183 % early May end June Claude tools 5,905 to 7,750 + 31 % early May end June Tool sprawl outpaced connector count. ChatGPT's live action count nearly tripled, from about 7,100 to 20,200 callable actions; Claude's directory grew from 5,900 to 7,750 tools. Existing connectors added tools
Existing connectors did not sit still. They grew new tools week after week, many of them able to act on your data.
Furthermore, connectors got denser
The number of tools in a connector also increased, which means more surface behind approving a single connector.
ChatGPT tools per connector 7.05 to 10.46 + 48 % early May end June Claude tools per connector 14.65 to 15.38 + 5 % early May end June And more of those tools can write Beyond expanding in size, connectors became more agentic and gained increased access to write capabilities. Read-only vs. write actions over time + 4,385 write actions 0 5,100 10,100 15,200 20,200 05-12 06-09 06-30 write read-only + 460 write tools 0 1,900 3,900 5,800 7,800 05-22 06-12 06-30 write read-only Both catalogs scaled their write actions as they grew, ChatGPT's write count by 222% and Claude's by 28%. Relative to the growth in read actions, though, they diverged.
For ChatGPT, the proportion of actions that write grew, as did the share of connectors with at least one write tool.
ChatGPT Connectors with a write tool 38 % to 46 % +8 pts mid-May end June Share of actions that write 27.6 % to 31.5 % +3.9 pts mid-May end June Meanwhile, for Claude, both the proportion of actions that write and the share of connectors with a write tool held steady as the catalog grew.
late May end June Share of actions that write 28.0 % to 27.3 % -0.7 pts late May end June Not every write is equal. ChatGPT flags which of its actions can delete or overwrite, and split three ways, that group grew the most, tripling from 798 to 2,415.
ChatGPT Action Classifications read-only 5,115 → 13,809 2.7x write 1,235 → 3,990 3.2x destructive 798 → 2,415 3.0x Additionally, across Claude and ChatGPT, some connectors have gone from completely read-only to fully write-capable:
ChatGPT 16 5 Vendor volatility is not consistent Throughout the vendor population, vendors changed at very different rates. One methodology organizations can use to proxy the risk of connectors enacting future changes is to analyze a connector's past volatility, or the volatility of connectors in the same industry.
See every connector change across your vendors
PromptArmor Threat Intelligence
Is your organization protected from AI in vendors?
PromptArmor continuously monitors across your portfolio of third party AI in vendors, skills, plugins, connectors, MCP servers, models and more.
We detect vulnerabilities and changes like this, surfacing risk before it becomes an incident.
Assess and monitor risk from AI in vendors with novel intelligence on emerging threats.
PromptArmor © 2026 All rights reserved
