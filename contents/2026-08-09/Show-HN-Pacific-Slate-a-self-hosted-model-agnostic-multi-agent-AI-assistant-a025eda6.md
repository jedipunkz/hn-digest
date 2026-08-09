---
source: "https://pacslate.com/"
hn_url: "https://news.ycombinator.com/item?id=49235865"
title: "Show HN: Pacific Slate: a self-hosted, model-agnostic multi-agent AI assistant"
article_title: "Pacific Slate"
author: "badwx"
captured_at: "2026-08-09T21:21:59Z"
capture_tool: "hn-digest"
hn_id: 49235865
score: 4
comments: 0
posted_at: "2026-08-09T21:04:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Pacific Slate: a self-hosted, model-agnostic multi-agent AI assistant

- HN: [49235865](https://news.ycombinator.com/item?id=49235865)
- Source: [pacslate.com](https://pacslate.com/)
- Score: 4
- Comments: 0
- Posted: 2026-08-09T21:04:02Z

## Translation

タイトル: Show HN: Pacific Slate: 自己ホスト型でモデルに依存しないマルチエージェント AI アシスタント
記事タイトル: パシフィック・スレート
説明: 記憶し、バックグラウンドで動作し、私が所有するハードウェア上で実行される自己ホスト型 AI システムです。それが何をするのか、そしてその背後にある設計上の決定。
HN テキスト: LLM を同期するシステムを維持するために、スタンドアロン コンピューターを購入したり、ラップトップを再利用して常に実行したくなかったので、これを構築しました。これは私のシステムの簡単な概要であり、開梱して自分で複製しやすいようにレイアウトされています。あるソリューションが別のソリューションにとって最適であるとは限らないため、プロジェクトは個別かつ一意に構成することを目的としています。むしろ、自分のプロジェクトで物事を実装する方法についてのアイデアが得られるかもしれません。ご多幸をお祈りします、ライアン。

記事本文:
">
パシフィックスレート //
ライブデモ
モニター
ソース
AI でオープン ▾
クロードでチャット
ChatGPTで開く
マークダウンとしてコピー
生のマークダウンを開く
llms.txt
パシフィック・スレート
私は MBA の 1 年目で、インターンシップに応募し、妻と 2 人の子供の世話をしながら大量の情報を管理しようとしていました。そして、その週に使用していた AI プログラムに対して自分の状況全体を再説明し続けました。私は、自動的に更新され、スロップがフィルタリングされ、データが自分のものに保たれるものが欲しかったのです。
これは、サービスをホストし、重要なファイルの暗号化されたバックアップを保持する専用マシンであるレンタル サーバー上で実行されます。私はそれを設計し、部品を選択し、毎日実行しました (技術的な詳細は下部の折りたたまれたセクションにあります)。
私のソースから自身を更新し、ノイズを分類し、私が個人的に所有するデータベースにデータを維持します。
マルチエージェント ツリーを使用して、作業を最適なモデルにルーティングします。
私の選択と好みを覚えていて、過去の結果から学びます（学んだ教訓を新しいタスクや状況に適用します）。
情報源を引用し、すべての行為のログを維持し、追跡可能性と説明責任を可能にします。
記憶している内容を信頼するのではなく、ソース文書をチェックします。
すでに使用しているツールやサービスに統合されるため、ワークフローの変更を強制されるのではなく、ワークフローに合わせて調整されます。
私は目新しさよりも耐久性、実用性、生産性を優先することをすぐに学びました。このシステムは製品ではなく、研究プロジェクトでもありません。これは私が毎日使用するツールであり、私が不在のときやメンテナンスする時間がないときでも機能する必要があります。私は当初、これを達成するためにモバイル クライアントとデスクトップ クライアントを構築していましたが、快適に投資できる以上の作業でした。それらを削除し、AI 開発全体で一般的なツール (MCP、プラグイン、フック) からアクセスできるようになりました。

意見。システムの将来性を確保することが重要であり、意図的にサービスに依存しない設計になっているため、その時点で最適なものに柔軟に対応できます。
8 人のエージェントが作業を行います。1 人はリクエストを読み取ってルーティングし、7 人の専門家は調査、コード、分析、レビュー、および関連作業を担当します。それぞれがそのジョブ用に選択されたモデルで実行されるため、短い検索と長い分析では同じコストが発生したり、同じ機械で待機したりすることはありません。
回答には出典が記載されているため、請求や通話料金を追跡できるため、支出は月末ではなく発生時に表示されます。
公開デモのワークスペース。答えは 1 つの長いスレッドとしてではなく、移動できるカードとして届き、それぞれに使用されたモデル、コスト、かかった時間がラベル付けされています。
デモ: pacslate.com/demo 。サンプルデータ上の実際のインターフェイス。
キャンバス。プロンプトを入力するか、プロンプトの実行を監視します。開く
モニター。ダッシュボード。地震データとヘッドラインはライブであり、ブラウザで取得されます。市場、航空機、その他個人的なものにはサンプル データというラベルが付けられます。開く
何が私のもので何がレンタルなのか
蓄積された部分はサーバーに残ります。モデルはレンタル中です。
このモデルは私が管理していない部品の 1 つであるため、交換可能として扱われます。より良いものに置き換えることは、構成の変更です。
保存された内容はエクスポートおよび削除可能であり、トレーニングには使用されません。送信されるリクエストは、データ保持ゼロのエンドポイントでのトレーニングや公開を行わないように構成されたプロバイダーに制限されます (アカウント設定、2026 年 8 月)。サーバーから何も残すべきでない場合は、サーバー上で実行されているモデルに作業を向けることができます。
蓄積されたものは私が保持し、リクエストの行き先は私が決定します。そのどれもがベンダーの資産にはなりません。
オープンソースと私が支払う対価
ほぼすべてがオープンソースであり、再構築されるのではなく、私自身のサーバー上でアセンブルされ、実行されます。

サービスとして提供されます。オープンソースの同等の機能が機能する場合は、それに切り替えてデータを保持しました。私が今でも支払っているのは、モデルと、交換可能で記録システムではない 1 つのホスト型メモリ サービスです。
コンピューター サイエンスとソフトウェア エンジニアリングには常に興味をそそられていましたが、私自身にコードを開発した経験はありませんでした。このシステムの開発は、現在オンラインで利用できる数多くのツールやサービスと連携することで可能になりました。アーキテクチャ自体はすべて独立して作成されています。どのコンポーネントが存在するか、それらがどのように接続され、それぞれが何を実行するか、部品が故障したときにどのように動作するか、そしてコンポーネントが「費やす」ことができるもの (各モジュールまたはプラットフォームの予算) です。 AI ネイティブで構築し、コーディング エージェントがコードの大部分を記述している間に指定とレビューを行いました。
2026 年の初めから実行されており、私は毎日使用しています。技術セクションで説明されている内容のほとんどは、計画ではなく運用から生まれました。たとえば、フォールバック層が存在するのは、フレームワーク自体のフォールバック設定が何も行われないことが判明したためです。別の例として、すべての回答のモデル ラベルは、ログには表示されなかったモデル スワップへの動作回帰をバックトレースする方法として存在します。本番環境で遭遇した問題のほとんどはバグではなく設計上のギャップであり、システムを運用して初めてそれらを発見しました。
私は働きながら大学院に通いながら Pacific Slate を構築したため、自由時間 (または自由時間の欠如) が何よりもデザインを形作りました。それは、スケジュールされた実行、自動的にマージされる低リスクの依存関係とセキュリティ更新、私を待たずに動作状態に低下する障害など、私がそこにいない時間に役立つ必要がありました。
このページおよびサンプル構成にあるモデル名は一例です。デザイン

モデルに依存せず、名簿は循環します。カウントは 2026 年 8 月現在のもので、実行中のシステムに対してチェックされました。
プラグインを使用したチャット モデルではなく、設計されたマルチ エージェント システム。
Google のエージェント開発キット (ADK) のマルチエージェント ツリー: ルートに 1 人のオペレーターと 7 人のスペシャリスト (コーダー、リサーチャー、アナリスト、生産性、レビュー担当者、評価者、およびコーダーを対象とする調査サブエージェント) が含まれます。
コストとフィット感。各役割は、その作業に最適で価格設定されたモデルにマッピングされます。
独立したレビュー。エバリュエーターは、出力をスコアリングするエージェントとは異なるモデル ファミリで実行されます。重みが異なるとバイアスも異なるとは限りませんが、自己採点よりは優れています。レビューアーは、ポリシーによって書き込みパスから切り離された別のインスタンスです。
最低限の特権。ツールは、明示的な許可リストによって役割ごとにスコープが設定されます。研究者はインフラに到達できません。レビュアーはチャーターにより読み取り専用となります。生産性エージェントの資格情報は分離されているため、そこで障害が発生してもオペレーターがダウンすることはありません。
封じ込め。レート制限やクラッシュは 1 人の専門家の中にとどまります。
ADK はサブエージェントに単一親制約を適用します。そのため、モデルとツールを共有しているにもかかわらず、コーダーの研究部門はスタンドアロンの研究者とは別のインスタンスになります。
スペシャリストは、ピアがアンビエント コンテキストとして読み取る Redis イベント ストリームに結果を公開します。より長い非同期ジョブは、同期パスを占有するのではなく、バックグラウンド オーケストレーターに引き渡されます。
モデルの独立性。サブストレートは、MCP ゲートウェイ (モデル コンテキスト プロトコルを介したツールとメモリ、MCP 対応クライアントによってマウント可能) と OpenAI 互換エンドポイントの 2 つの方法で公開されます。モデルが移動しても、メモリ、データ、ツール、ルーティング、検証は移動しません。
2 つのルールがすべてを管理します。決定的な関連性パスにより、何をロードするかが決定されます

モデルを呼び出すことなく、モデルを実行する前に。そして、思い出された記憶は、事実ではなく手がかりとして扱われます。ベンダーが 7 月に更新すると記憶があれば、エージェントはソース文書を取り出し、それに基づいて行動します。
積極的な仕事。スケジュールされたルーチンは、概要に入る前に、多くのソースから情報を取得し、各候補項目の関連性、重要性、新規性についてスコアを付けます。破棄されたアイテムはサイレントにドロップされるのではなくログに記録されるため、フィルターを監査できます。統合ステージでは、何かが保存される前に重複排除と正規化が行われ、途中で資格情報と個人データが検査されます。
まずはアルゴリズム。デフォルトのデータ アクセスと分類は、正規表現、インデックス、SQL、セット メンバーシップなどの決定的なものに設定されます。既約部分のみにモデルを使用します。複雑さスコアラー、ツール検索、およびメモリ関連性パスはすべてこれを応用したものです。
失敗を可視化すれば、コストは安くなります。各信頼性修正はガードレールとトレースを組み合わせているため、同じクラスの問題が目に見えずに再発することはありません。
ピーク時の能力を超えた回復力。この 2 つが競合する場合は、稼働時間が優先されます。
設計上の制約としてのコスト。実際の上限では、モデル、コンテキスト サイズ、モデルをいつ使用するかについて正直な選択が求められます。
ルーティング構造については、examples/model-routing.example.yaml を参照してください。デモの各カードのモデルとコスト ラベルも参照してください。
モデル プロバイダー、ツール、データ、メモリ、フロントエンドは、標準インターフェイス経由でアクセスできる 1 つの環境に統合されています。
Next.js のシングルページ アプリであるキャンバスは、ストリーミング接続を開き、プロンプトを送信します。ストリーミングでは Server-Sent Events 上で AG-UI プロトコルを使用するため、トークン、ツール呼び出し、モデル解決イベントは 1 つのストリームで到着します。
バックエンドは、キャンバスの AG-UI エンドポイントと他のクライアントの OpenAI 互換エンドポイントを公開する Starlette サービスです。
オペレーターが分類する

リクエストにインラインで応答するか、専門家に引き継ぎます。
スペシャリストはコンテキストをロードし、MCP 経由でツールを呼び出します。
復元力のあるモデル層は、その呼び出しのモデルを解決し、コストを意識したダウングレードと健全性を意識したフォールバック (セクション 3) を適用し、完了をストリーミングし、コストを測定します。
答えは、モデルとコストがラベル付けされたカードとして返されます。永続的な事実は記憶に書き込まれます。
単一のモデル呼び出しでのインライン応答と、スペシャリストが独自のツール ループを実行する委任された実行という 2 つの速度階層が生成されます。
システム図。クリックすると拡大します。
ツールレイヤー。 1 つのトークン認証ゲートウェイ、ストリーミング可能な HTTP 上の FastMCP の背後に約 24 台のツール サーバー (メモリ、取得、コード実行、git、フィード、インフラストラクチャ) があり、各サーバーは独自のプロセスを備えています。
登録します。再配線しないでください。新しい機能とは、サーバーを登録することを意味します。追加のホップとゲートウェイでの単一障害ドメインのコストがかかります。そのため、ゲートウェイには独自のヘルス チェックがあり、フリートはヘルス ポーリングされます。
検索しますが、読み込まないでください。すべてのエージェントのコンテキストですべてのツール スキーマを運ぶのはコストがかかり、注意力が薄れるため、ゲートウェイはツール カタログに対する BM25 検索を公開します。 MCP 仕様にはツールのみがリストされています。この検索は追加です。
上の図を参照してください。デモは、スクリプト化されたイベントに対して実行される同じキャンバスとプロトコルです。
稼働し続け、予算内に収まり、自身の出力をチェックし、それらのメカニズムはそれぞれ観察可能です。
2026 年 8 月の時点で 90 日間のトレース: 13,053 回の実行、26,515 回のモデル呼び出し、9 つのモデルがトラフィックを処理しました。モデル呼び出しの中央値は 3.1 秒で完了します。 90 パーセンタイルは 14 秒です。
コスト管理。通話前に適用されます。決定論的スコアラーは、単語数、キーワード クラス、コード マーカー、役割の重み、会話の深さからプロンプトの複雑さを 1 ～ 5 で評価します。n は

o モデルが関与している。そのスコアと現在の予算の状態によって層が選択されます。
支出は毎月の厳しい上限に対してリクエストごとに追跡され、アトミックな書き込みにより台帳は再起動後も存続します。
後退する。私が構築した ADK バージョンでは、モデル ラッパーがプロバイダーを直接呼び出し、基礎となるライブラリのフォールバック構成を無視したため、その設定は何も行いませんでした。モデル層をそれを実装するものでラップしました。
ヘルス トラッカーはレート制限されたモデルを記録し、違反の繰り返しに応じて増加するクールダウンを適用します。その後の呼び出しでは、同じ壁にぶつかるのではなく、冷却モデルをスキップします。
最初のトークンまでの時間ウォッチドッグは、試行の最初の応答のみを制限します。発行せずにストールしたモデルはトリップし、呼び出しは次のモデルに移ります。何も生成されないため、重複した出力はありません。ストリームが開始されると、制限なく実行され、バックストップとしてより長い外部タイムアウトが使用されます。
ツール定義を受け入れることができないフォールバック モデルは、互換性ミスとしてスキップされ、失敗として報告されません。
チェーンが使い果たされた場合、レイヤーは実行グループをクラッシュさせるのではなく、スペシャリストが利用できないことを示すメッセージを返します。
チェーンはロールごとに異なるモデルで終了するため、2 つのロールが最後の手段を共有することはありません。どれでもw

[切り捨てられた]

## Original Extract

A self-hosted AI system that remembers, works in the background, and runs on hardware I own. What it does, and the design decisions behind it.

I didn't want to buy a standalone computer or repurpose a laptop to run constantly so I could maintain a system to sync my LLMs, so I built this. It's a simple overview of my system, laid out in a way easy to unpack and replicate for yourself. The project is meant to be configured individually, and uniquely, since one solution might not be what's best for another. If anything, maybe it gives you some ideas on how to implement things for your own project. Best wishes, Ryan.

">
Pacific Slate //
live demo
monitor
source
Open with AI ▾
Chat in Claude
Open in ChatGPT
Copy as Markdown
Open raw Markdown
llms.txt
Pacific Slate
I was in the first year of an MBA, trying to manage a firehose of information while applying to internships and keeping up with my wife and two kids, and I kept re-explaining my whole situation to whatever AI program I was using that week. I wanted one that updated itself, filtered the slop, and kept the data mine.
It runs on a rented server, a dedicated machine that hosts the services and holds encrypted backups of the files I care about. I designed it, selected the parts, and run it day to day (technical details in the collapsed sections at the bottom).
Updates itself from my sources, sorts through the noise, and maintains the data in my privately owned database.
Uses a multi-agent tree to route work to the model best suited for it.
Remembers my choices and preferences, and learns from past outcomes (applying lessons learned to new tasks or situations).
Cites sources and maintains logs of everything it does, allowing for traceability and accountability.
Checks the source document rather than trusting what it remembers.
Integrates into the tools and services I already use, so it adjusts to my workflow rather than forcing me to change it.
I learned quickly to prioritize durability, utility, and productivity over novelty. The system is not a product, and it is not a research project. It is a tool I use every day, and it has to work when I’m not there or don’t have the time to maintain it. I had originally built mobile and desktop clients to reach it, but they were more work than I could comfortably invest. I deleted them and now reach it from the tools (MCP, plugins, hooks) commonplace throughout AI development. It’s been important to futureproof the system, and the design is intentionally service-agnostic so it can flex to whatever is best at the time.
Eight agents do the work: one that reads the request and routes it, and seven specialists for research, code, analysis, review, and related work. Each runs on a model chosen for that job, so a short lookup and a long analysis do not draw the same cost or wait on the same machinery.
Answers carry their sources, so a claim can be traced, and the cost of the call, so spending stays visible as it happens rather than at the end of the month.
The workspace, from the public demo. Answers arrive as cards that can be moved around rather than as one long thread, each labeled with the model used, the cost, and the time taken.
Demo: pacslate.com/demo . The real interface on sample data.
Canvas. Enter a prompt or watch it run. Open
Monitor. A dashboard. Seismic data and headlines are live, fetched in the browser; markets, aircraft, and anything personal are labeled sample data. Open
What is mine and what is rented
The parts that accumulate stay on my server. The model is rented.
The model is the one part I do not control, so it is treated as replaceable. Substituting a better one is a configuration change.
What is stored is exportable and deletable, and is not used for training. Requests that go out are restricted to zero-data-retention endpoints, at providers configured not to train on them or publish them ( account settings , August 2026). Where nothing should leave the server, work can be pointed at a model running on it.
I hold what accumulates, I decide where requests go, and none of it becomes a vendor’s asset.
Open source, and what I pay for
Almost all of it is open source, assembled and run on my own server rather than rented as a service. Where an open-source equivalent did the job, I switched to it and kept the data. What I still pay for is the model, and one hosted memory service that stays swappable and is not the system of record.
While computer science and software engineering have always been intriguing, I didn’t have any background developing code myself. The development of this system has been possible through working in partnership with the numerous tools and services currently available online. The architecture itself has all been created independently: which components exist, how they connect, what each runs on, how it behaves when a part fails, and what it is allowed to ‘spend’ (the budget for each module or platform). I built it AI-natively, specifying and reviewing while coding agents wrote most of the code.
It has run since early 2026 and I use it daily. Most of what the technical sections describe came out of operating it rather than planning it. For example, the fallback layer exists because the framework’s own fallback setting turned out to do nothing. Another instance, the model label on every answer exists as a way to backtrace a behavioral regression to a model swap that was invisible in the logs. Most of the issues I hit in production were not bugs but design gaps, and it was only through operating the system that I discovered them.
Because I built Pacific Slate while working and going to grad school, my free time (or lack thereof) shaped the design more than anything else. It had to be useful in the hours I wasn’t there: scheduled runs, low-risk dependencies and security updates that merge on their own, and failures that degrade to a working state rather than waiting for me.
Model names on this page and in the example config are illustrative; the design is model-agnostic and the roster rotates. Counts are current as of August 2026 and were checked against the running system.
A designed multi-agent system, not a chat model with plugins.
A multi-agent tree on Google’s Agent Development Kit (ADK) : one operator at the root plus seven specialists (coder, researcher, analyst, productivity, reviewer, evaluator, and a research sub-agent scoped to the coder).
Cost and fit. Each role maps to the model best suited and priced for its work.
Independent review. The evaluator runs on a different model family than the agents whose output it scores. Different weights are not guaranteed different biases, but it beats self-scoring. The reviewer is a separate instance kept off the write path by policy.
Least privilege. Tools are scoped per role by explicit allow-list. The researcher cannot reach infrastructure. The reviewer is read-only by charter. The productivity agent’s credentials are isolated, so a failure there cannot take down the operator.
Containment. A rate-limit or crash stays inside one specialist.
ADK enforces a single-parent constraint on sub-agents, which is why the coder’s research arm is a separate instance from the standalone researcher despite sharing a model and tools.
Specialists publish findings to a Redis event stream that peers read as ambient context. Longer asynchronous jobs hand off to a background orchestrator rather than occupying the synchronous path.
Model independence. The substrate is exposed two ways: an MCP gateway (tools and memory over the Model Context Protocol, mountable by any MCP-capable client) and an OpenAI-compatible endpoint . Memory, data, tools, routing, and verification do not move when the model does.
Two rules govern all of it. A deterministic relevance pass decides what to load before any model runs, with no model call. And recalled memory is treated as a lead rather than a fact: if memory says a vendor renews in July, the agent pulls the source document and acts on that.
Proactive jobs. Scheduled routines pull from many sources and score each candidate item for relevance, materiality, and novelty before it can enter a brief. Discarded items are logged rather than dropped silently, so the filter can be audited. A consolidation stage de-duplicates and normalizes before anything is stored, and screens for credentials and personal data on the way in.
Algorithm first. Default data access and classification to something deterministic: regex, an index, SQL, set membership. Use a model only for the irreducible part. The complexity scorer, the tool search, and the memory relevance pass are all applications of this.
Make failure visible, then cheap. Each reliability fix pairs a guardrail with a trace, so the same class of problem cannot recur unseen.
Resilience over peak capability. Where the two conflict, uptime wins.
Cost as a design constraint. A real ceiling forces honest choices about model, context size, and when to use a model at all.
See: examples/model-routing.example.yaml for the routing structure · the model and cost label on each card in the demo .
Model providers, tools, data, memory, and a frontend are integrated into one environment reachable over standard interfaces.
The canvas, a Next.js single-page app, opens a streaming connection and posts the prompt. Streaming uses the AG-UI protocol over Server-Sent Events , so tokens, tool calls, and model-resolution events arrive on one stream.
The backend is a Starlette service exposing the AG-UI endpoint for the canvas and an OpenAI-compatible endpoint for other clients.
The operator classifies the request and either answers inline or hands off to a specialist.
The specialist loads context and calls tools over MCP.
The resilient model layer resolves the model for that call, applies cost-aware downgrades and health-aware fallback (section 3), streams the completion, and meters cost.
The answer streams back as a card labeled with model and cost. Durable facts are written to memory.
Two speed tiers result: inline answers in a single model call, and delegated runs where a specialist works its own tool loop.
System diagram. Click to enlarge.
Tool layer. About two dozen tool-servers (memory, retrieval, code execution, git, feeds, infrastructure) behind one token-authenticated gateway, FastMCP over streamable HTTP, each server its own process.
Register, do not rewire. New capability means registering a server. The cost is an extra hop and a single failure domain at the gateway, which is why the gateway has its own health check and the fleet is health-polled.
Search, do not load. Carrying every tool schema in every agent’s context is expensive and dilutes attention, so the gateway exposes a BM25 search over the tool catalog . The MCP specification only lists tools; this search is an addition.
See: the diagram above · the demo , which is the same canvas and protocol running against scripted events.
It stays up, stays within budget, and checks its own output, and each of those mechanisms is observable.
Ninety days of traces as of August 2026: 13,053 runs, 26,515 model calls, nine models serving traffic. The median model call completes in 3.1 seconds; the 90th percentile is 14 seconds.
Cost control, applied before the call. A deterministic scorer rates prompt complexity 1 to 5 from word count, keyword classes, code markers, role weight, and conversation depth, with no model involved. That score and current budget state select the tier:
Spend is tracked per request against a hard monthly ceiling, with atomic writes so the ledger survives a restart.
Fallback. In the ADK version I built on, the model wrapper called the provider directly and ignored the underlying library’s fallback configuration, so that setting did nothing. I wrapped the model layer with one that implements it.
A health tracker records rate-limited models and applies an escalating cooldown that grows with repeat offenses. Later calls skip a cooling-down model rather than hitting the same wall.
A time-to-first-token watchdog bounds only the first response of an attempt. A model that stalls without emitting trips it and the call moves to the next model. Because nothing was yielded, there is no duplicate output. Once a stream starts it runs unbounded, with a longer outer timeout as backstop.
A fallback model that cannot accept the tool definitions is skipped as a compatibility miss, not reported as a failure.
If the chain is exhausted, the layer returns a message stating that specialist is unavailable rather than crashing the run group.
Chains terminate on different models per role, so no two roles share a last resort. Any w

[truncated]
