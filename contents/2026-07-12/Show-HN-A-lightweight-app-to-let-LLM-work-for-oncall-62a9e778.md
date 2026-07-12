---
source: "https://www.neatcontext.com/"
hn_url: "https://news.ycombinator.com/item?id=48882433"
title: "Show HN: A lightweight app to let LLM work for oncall"
article_title: "NeatContext | Local context for better LLM answers"
author: "tanglearncode"
captured_at: "2026-07-12T16:46:17Z"
capture_tool: "hn-digest"
hn_id: 48882433
score: 2
comments: 1
posted_at: "2026-07-12T16:39:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A lightweight app to let LLM work for oncall

- HN: [48882433](https://news.ycombinator.com/item?id=48882433)
- Source: [www.neatcontext.com](https://www.neatcontext.com/)
- Score: 2
- Comments: 1
- Posted: 2026-07-12T16:39:28Z

## Translation

タイトル: HN を表示: LLM をオンコールで機能させるための軽量アプリ
記事のタイトル: NeatContext |より良い LLM の答えを得るためにローカルのコンテキスト
説明: LLM ローカル ドメインの知識と内部システムへの読み取り専用アクセスを提供する軽量のデスクトップ ワークスペース。
HN テキスト: 現在の SRE エージェントにおける最大のギャップは、ドメインの知識が欠如していることです。そのため、オンコールのインシデント処理に関して正確な結果を提供することはできません。問題を正確に解決する軽量で適切なソリューションが見つかりませんでした。そこで、デスクトップ アプリケーション NeatContext を作成しました。これは、LLM にインシデントを適切に処理するためのドメイン知識を知らせることができるだけです。

記事本文:
ニートコンテキスト |より良い LLM の回答のためのローカル コンテキスト NeatContext 仕組み ドキュメント コミュニティの価格設定 ブログ 連絡先 メニュー 仕組み ドキュメント コミュニティの価格設定 ブログ 連絡先 ローカルファースト コンテキスト ワークスペース ドメインの知識。
よりシャープな LLM の答え。
チームがすでに信頼しているコンテキストを使用して問題を分析します。 NeatContext は、デスクトップからローカルの知識と内部システムを LLM にもたらします。
macOS (Apple Silicon) · macOS (Intel) でも利用可能 ドキュメントを読む 重いセットアップは不要 お使いのモデル、お使いのデータ MCP で拡張可能 例 01 · GROUNDED VS GENERIC
同じ事件です。とても2つ
異なる答え。
09:07 UTC: checkout-api が 5xx のスローを開始し、INC-1001 が開きます。オンコールのインフラ チームは、NeatContext と一般的な SRE エージェントに同じ質問をします。「最初に何を確認する必要がありますか? 安全なアクションは何ですか?」
根本原因が見つかりました: 08:58Z プール サイズのカット。 RELOAD で元に戻します。
読み取り専用ツールでチェックされた証拠
デプロイ ログには、08:58Z (最初の checkout-api 5xx の分前) の pgbouncer でのdefault_pool_size 100 → 40 が示されています。
ログは、クライアントのキューイングによりプールが 40/40 で飽和していることを確認します。これは、事後検証 INC-0977 の正確な署名です。
Postgres プライマリは正常であり、max_connections=200 をはるかに下回っています。上限はデータベースではなくプールの設定です。
default_pool_size を 100 に戻し、pgbouncer を再ロードします — ドロップされた接続はゼロです。
プライマリをフェイルオーバーしないでください。TSG は飽和状態になるとフェイルオーバーを禁止します。飛行中のすべての接続が切断され、障害が拡大する可能性があります。
プールの排水を監視し、checkout-api 5xx がベースラインに戻っていることを確認します。
5xx の上昇にはさまざまな原因が考えられます。一般的なチェックリストを次に示します。
最近のアプリケーションのデプロイメントを確認し、疑わしいものがあればロールバックします。 ⚠ アプリのデプロイには問題はありません - 重要な変更は 08:58Z pgboun

チームのデプロイ ログとランブックのみが参照する cer プール サイズの削減。
データベースの健全性をチェックします。接続がスタックしている場合は、データベースを再起動するか、レプリカにフェイルオーバーします。 ⚠ 飽和しているが正常なプライマリをフェイルオーバーすると、すべてのサービスにわたるすべての実行中の接続が切断されます。db-failover-tsg.md はまさにこれを禁止しています。
圧力を軽減するには、Postgres の max_connections を増やします。 ⚠ 間違ったノブ: pgbouncer プールの天井が制約になっています。 max_connections を増やすと、障害がデータベースに移動されるだけです。
ログでエラーを確認し、問題が解決しない場合は適切なチームにエスカレーションしてください。
例 02 · チームごとの知識
異なるチーム、異なる知識、
違う正解。
同じINC-1001でもチームごとに扱いが異なります。 NeatContext は、チーム独自の知識に基づいて、それぞれに適切な措置を講じます。支払いの場合は証拠を提供し、インフラの場合は修正を行います。
支払いの根本原因ではありません。証拠を添えてインフラチームに引き渡してください。
読み取り専用ツールでチェックされた証拠
09:02:08Z の最初のプール エラーは、09:05:00Z checkout-api のデプロイより前に発生しました。ロールバックしても何も解決されません。
Stripe は正常です: ステータスは緑色、正常性プローブは 200 OK — 決済プロバイダーは除外されています。
主なエラーは pgbouncer: プール「請求」がいっぱいです (40/40) — 私たちが消費しているが所有していない共有インフラストラクチャです。
インシデント リンク、最初のプール エラー ログ行、展開タイムラインを使用してインフラ チームにエスカレーションします。これは、INC-0931 を 19 分で解決したのと同じ手順です。
暫定的な救済のみ: Stripe_webhook_auto_retry を無効にします。これはフラグの反転であり、再起動は必要ありません。
10:00Z より前に invoice-worker を再起動しないでください。毎日の決済は途中で行われており、中断すると顧客に二重請求される可能性があります。
根本原因: 08:58Z pgbouncer でのdefault_pool_size 100 → 40 のカット。
読み取り専用ツールでチェックされた証拠
デモリスト_デプル

oyments は、08:58Z でのdefault_pool_size 100 → 40 の変更を示しています。最初の checkout-api 5xx は 09:02:08Z に続きます。
billing-postgres ログには、クライアントがキューに入れられており、プール「billing」がいっぱい (40/40) であることが示されています (正確な INC-0977 シグネチャ)。
プライマリは正常です (CPU ~60%、 max_connections=200 を十分に下回っています)。設定したプールの上限が制約です。
default_pool_size を 100 に戻し、pgbouncer を再ロードします。INC-0977 では、ドロップされた接続はなく、プールは約 90 秒で空になりました。
billing-postgres プライマリをフェイルオーバーしないでください。TSG は飽和状態になるとフェイルオーバーを禁止します。それは停電を拡大することになるだろう。
checkout-api 5xx がベースラインに戻ったことを確認し、プール使用率が 80% を超えた場合にアラートを発行します。
「ログを確認し、最近の変更を確認し、データベースの健全性を確認し、必要に応じて適切なチームにエスカレーションします。」
各チームに 1 つのチェックリスト。 Payments チームやインフラチームの知識がなければ、Payments のように所有権を決定したり、Infra のように正確なプール サイズを元に戻す名前を指定したりすることはできません。
LLM に知識を与えてください
なくなってしまった。
すでに使用しているモデルの柔軟性を維持します。チームだけが知っている運用コンテキストを追加します。
各サービス、製品、顧客、または運用領域に焦点を当てた Markdown プロファイルを作成します。
ローカル フォルダーと読み取り専用の内部ツールを追加します。 NeatContext は、それらが存在する場所にインデックスを付けます。
設定されている LLM に問い合わせてください。ローカル ソースと目に見えるツール アクティビティを使用して、根拠のある分析を取得します。
設計により軽量です。
本質的に柔軟です。
単一のプロファイルと Runbook のフォルダーから始めます。必要な場合にのみモデル、ナレッジ、内部ツールを追加します。
OpenAI 互換、Anthropic、Ollama、または会社独自のゲートウェイを使用します。
システムの進化に合わせて、作成、レビュー、バージョンアップ、変更が簡単に行えます。
チームが依存する運用システム用の読み取り専用 MCP ツールを追加します。
厳密なカタログはありません。権利を構成する

各ワークフローとドメインのコンテキスト。
あなたのコンテキストはそのまま残ります
あなたのコントロール下にあります。
プロファイル、インデックス、会話、およびツールのログは、デフォルトでマシン上に残ります。 NeatContext は、知識が別の場所に存在することなく、構成されたモデルを調整します。
ローカルプロファイルとナレッジインデックス
ホストされた NeatContext モデルは必要ありません
地元の知識を活用して
役に立つ答え。
軽量のデスクトップ アプリをダウンロードして、独自のモデルを持ち込んでください。
NeatContext LLM ツール用のローカルファースト コンテキスト ワークスペース。

## Original Extract

A lightweight desktop workspace that gives your LLM local domain knowledge and read-only access to internal systems.

The biggest gap today in SRE agents is missing domain knowledge. So they cannot give accurate results for oncall incident handling. I just didn't find any lightweight and suitable solution to exactly solve the problem. So I created a desktop application NeatContext. It does nothing but can let LLMs know your domain knowledge to handle the incidents right.

NeatContext | Local context for better LLM answers NeatContext How it works Docs Community Pricing Blog Contact Menu How it works Docs Community Pricing Blog Contact Local-first context workspace Your domain knowledge.
Sharper LLM answers.
Analyze issues with the context your team already trusts. NeatContext brings local knowledge and internal systems to any LLM, right from your desktop.
Also available for macOS (Apple Silicon) · macOS (Intel) Read the docs No heavy setup Your model, your data Extensible with MCP EXAMPLE 01 · GROUNDED VS GENERIC
Same incident. Two very
different answers.
09:07 UTC: checkout-api starts throwing 5xx and INC-1001 opens. The Infra Team on-call asks NeatContext and a general SRE agent the same question: “What should we check first, and what is the safe action?”
Root cause found: the 08:58Z pool-size cut. Revert it with a RELOAD .
Evidence, checked with read-only tools
The deploy log shows default_pool_size 100 → 40 on pgbouncer at 08:58Z — minutes before the first checkout-api 5xx.
Logs confirm the pool saturated at 40/40 with clients queuing — the exact signature of postmortem INC-0977.
The Postgres primary is healthy and far below max_connections=200 — the ceiling is the pool setting, not the database.
Revert default_pool_size to 100 and RELOAD pgbouncer — zero dropped connections.
Do not fail over the primary — the TSG forbids failover for saturation; it would drop every in-flight connection and widen the outage.
Watch the pool drain and confirm checkout-api 5xx returns to baseline.
Elevated 5xx can have many causes — here is a general checklist.
Review recent application deployments and roll back if anything looks suspicious. ⚠ The app deploy is innocent — the change that matters is the 08:58Z pgbouncer pool-size cut, which only the team's deploy log and runbook point to.
Check database health; if connections are stuck, restart the database or fail over to a replica. ⚠ Failing over a saturated-but-healthy primary drops every in-flight connection across every service — db-failover-tsg.md forbids exactly this.
Increase max_connections on Postgres to relieve the pressure. ⚠ Wrong knob: the pgbouncer pool ceiling is the constraint. Raising max_connections just moves the failure into the database.
Check the logs for errors and escalate to the appropriate team if the issue persists.
EXAMPLE 02 · KNOWLEDGE PER TEAM
Different team, different knowledge,
a different right answer.
The same INC-1001, handled differently by each team. Grounded in a team's own knowledge, NeatContext gives each its right move: for Payments a hand-off with evidence, for Infra the fix.
Not a Payments root cause — hand off to Infra Team, with evidence.
Evidence, checked with read-only tools
First pool error at 09:02:08Z predates our 09:05:00Z checkout-api deploy — rolling back would fix nothing.
Stripe is healthy: status green, health probe 200 OK — the payment provider is ruled out.
Dominant error is pgbouncer: pool 'billing' is full (40/40) — shared infrastructure we consume but do not own.
Escalate to Infra Team with the incident link, the first pool-error log line, and the deploy timeline — the same play that resolved INC-0931 in 19 minutes.
Interim relief only: disable stripe_webhook_auto_retry — a flag flip, no restart needed.
Do not restart invoice-worker before 10:00Z — daily settlement is mid-flight, and interrupting it can double-charge customers.
Root cause: the 08:58Z default_pool_size 100 → 40 cut on pgbouncer.
Evidence, checked with read-only tools
demo_list_deployments shows our default_pool_size 100 → 40 change at 08:58Z ; the first checkout-api 5xx follows at 09:02:08Z .
billing-postgres logs show pool 'billing' is full (40/40) with clients queued — the exact INC-0977 signature.
The primary is healthy (CPU ~60%, well under max_connections=200 ) — the pool ceiling we set is the constraint.
Revert default_pool_size to 100 and RELOAD pgbouncer — in INC-0977 the pool drained in ~90 seconds with zero dropped connections.
Do not fail over the billing-postgres primary — the TSG forbids failover for saturation; it would widen the outage.
Confirm checkout-api 5xx returns to baseline, then alert on pool utilization above 80%.
“Check the logs, review recent changes, verify database health, and escalate to the appropriate team if needed.”
One checklist for every team. Without payments-team or infra-team knowledge it cannot decide ownership like Payments — or name the exact pool-size revert like Infra.
Give your LLM the knowledge
it's been missing.
Keep the flexibility of the models you already use. Add the operational context only your team knows.
Create focused Markdown profiles for each service, product, customer, or operational area.
Add local folders and read-only internal tools. NeatContext indexes them where they live.
Ask your configured LLM. Get grounded analysis with local sources and visible tool activity.
Lightweight by design.
Flexible by nature.
Start with a single profile and a folder of runbooks. Add models, knowledge, and internal tools only when you need them.
Use OpenAI-compatible, Anthropic, Ollama, or your company's own gateway.
Easy to write, review, version, and change as your systems evolve.
Add read-only MCP tools for the operational systems your team relies on.
No rigid catalog. Compose the right context for each workflow and domain.
Your context stays
under your control.
Profiles, indexes, conversations, and tool logs stay on your machine by default. NeatContext orchestrates your configured models without becoming another place your knowledge has to live.
Local profiles and knowledge indexes
No hosted NeatContext model required
Turn local knowledge into
useful answers.
Download the lightweight desktop app and bring your own model.
NeatContext Local-first context workspace for LLM tools.
