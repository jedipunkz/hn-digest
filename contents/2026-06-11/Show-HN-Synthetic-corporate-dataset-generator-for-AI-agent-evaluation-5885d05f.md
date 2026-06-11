---
source: "https://github.com/aeriesec/orgforge"
hn_url: "https://news.ycombinator.com/item?id=48494178"
title: "Show HN: Synthetic corporate dataset generator for AI agent evaluation"
article_title: "GitHub - aeriesec/orgforge: Synthetic corporate dataset generator for AI agent evaluation. · GitHub"
author: "jflynt76"
captured_at: "2026-06-11T19:05:52Z"
capture_tool: "hn-digest"
hn_id: 48494178
score: 2
comments: 0
posted_at: "2026-06-11T18:09:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Synthetic corporate dataset generator for AI agent evaluation

- HN: [48494178](https://news.ycombinator.com/item?id=48494178)
- Source: [github.com](https://github.com/aeriesec/orgforge)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T18:09:30Z

## Translation

タイトル: Show HN: AI エージェント評価用の合成企業データセット ジェネレーター
記事のタイトル: GitHub - aeriesec/orgforge: AI エージェント評価用の合成企業データセット ジェネレーター。 · GitHub
説明: AI エージェント評価用の合成企業データセット ジェネレーター。 - aeriesec/orgforge

記事本文:
GitHub - aeriesec/orgforge: AI エージェント評価用の合成企業データセット ジェネレーター。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
エアリーセック
/
オーグフォージ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラン

hes タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
60 コミット 60 コミット .github/ workflows .github/ workflows config config eval eval scripts scripts src src testing testing .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile INSIDER_THREAT.md INSIDER_THREAT.md ライセンス ライセンス README.md README.md docker-compose.yaml docker-compose.yaml orgforge_hero.png orgforge_hero.png pyproject.toml pyproject.toml 要件-cloud.txt 要件-cloud.txt 要件-test.txt 要件-test.txt 要件.txt 要件.txt シミュレーション.log.txt シミュレーション.log.txt uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
グラウンドトゥルースのエコシステムを生成し、エンタープライズ AI エージェントを評価するための決定論的な企業シミュレーター
OrgForge は、イベント駆動型のステート マシンに基づいて、Confluence ページ、JIRA チケット、Slack スレッド、Git PR、Zoom トランスクリプト、Zendesk チケット、Salesforce レコード、電子メール、サーバー テレメトリなど、数週間にわたる現実的な企業アクティビティをシミュレートするため、LLM が事実を順序どおりに幻覚表示することはできません。
データセットは、生きたシミュレーションの結果です。エンジニアはスプリントの途中で離脱し、決定的なインシデントの引き継ぎ、チケットの再割り当て、CRM の所有権の失効が余儀なくされます。文書化が不十分なシステムが壊れると、知識のギャップが表面化します。新入社員は、模擬コラボレーションを通じて社内ネットワークを構築します。ストレスは、リアルタイムの重み付けされたソーシャル グラフを通じて伝播します。すべての成果物は、作成された時点の組織の正確な状態を反映しています。
オプション 1 — Docker 内のすべて
オプション 2 — ローカル Ollama、MongoDB のみの Docker
組織的な知識に基づいて推論する AI エージェントを構築する場合、

テストするための現実的なコーパスが必要です。広く使用されている唯一の企業データセットはエンロン電子メール コーパスです。これは 25 年前のもので、法的に機密性が高く、危機に瀕している 1 社をカバーしています。
OrgForge は、あらゆる企業、業界、組織構造に合わせてパラメータ化されたコーパスをゼロから生成します。 LLM は散文を書きますが、事実 (誰がオンコールだったか、どのチケットがオープンだったか、いつインシデントが解決したか、誰がチームを離れたばかりか、どの顧客の SLA に違反したか) はステート マシンによって厳密に管理されます。
設計上の中心的な賭け: LLM 出力を決定論的なイベント ログに固定することで、データセットが検索システムの評価に実際に役立ちます。何が、いつ、誰が関与し、組織の状態がどのようなものであったかについての正確な情報が得られるため、エージェントが単にもっともらしいコンテキストだけでなく、適切なコンテキストを明らかにしたかどうかを測定できます。
実際のシミュレーションの一部が生成するものを次に示します。 8日目にインシデントが発生します。
lack/channels/engineering-incidents.json — アラートが最初に到着し、オンコールのポケットベルが起動したミリ秒のタイムスタンプが付けられます。
{
"ts" : " 2026-03-10T14:23:07 " ,
"ユーザー": " pagerduty-bot " ,
"text" : " 🔴 P1 アラート: TitanDB 遅延スパイク — 負荷による接続プールの枯渇。オンコール: Jax。"
}
jira/IT-108.json — 数秒後に開かれ、同じ SimEvent から事実が抽出されました。
{
"id" : " IT-108 " ,
"type" : " インシデント " ,
"優先度" : " P1 " 、
"title" : " TitanDB: レイテンシのスパイク " ,
"root_cause" : "負荷による接続プールの枯渇" ,
"担当者" : " ジャックス " 、
"レポーター" : "システム" ,
"開いた" : " 2026-03-10T14:23:19 "
}
confluence/postmortems/IT-108.md — 翌日に書かれ、同じ根本原因と PR が関連付けられています。
このインシデントは、持続的な負荷による接続プールの枯渇によって引き起こされ、IT-108 で初めて表面化しました。この修正は PR #47 に反映されました (Sarah によってマージされました)。あ

TitanDB 接続管理における事前の知識のギャップ (Jordan の 12 日目の出発に起因) が、診断の遅れの原因となりました。
一方、 datadog/metrics.jsonl の時系列データは正確なレイテンシのスパイクを反映しており、影響を受ける顧客からの Zendesk サポート チケットは自動的に「緊急」にエスカレーションされ、Salesforce の商談には「リスクがある」というフラグが付けられ、月末の顧客請求書 ( invoices/ ) にはインシデントの期間に基づいて SLA クレジットが自動的に適用されます。
これはどれも偶然ではありません。すべては、すべてのダウンストリーム アーティファクトの読み取り元となる 1 つの SimEvent に遡ります。
デフォルトの 22 日間のシミュレーションでは、次の結果が生成されます。
OrgForge は LLM ラッパーではありません。 4 つの連動システムが正確さを保証します。
👉 アーキテクチャの詳細については、こちらをご覧ください。
シミュレーションで最も複雑な動作。エンジニアがスプリントの途中で離脱すると、その日の計画が実行される前に、次のものが順番に起動されます。
インシデントの引き継ぎ — 退職するエンジニアに割り当てられたアクティブなインシデントは、ダイクストラ エスカレーション ルーティングを介して (ノードがまだグラフ内にある間) チェーン内の次に対応可能な担当者に再ルーティングされます。
チケットと CRM の再割り当て — 孤立した JIRA チケットは部門リーダーに送られます。退職した従業員が所有する Salesforce アカウントとオープンな商談には、再割り当てのフラグが立てられ、クロスドメインのグラウンド トゥルースが維持されます。
グラフの再計算 — 小さい方のグラフで媒介中心性が再計算されます。離脱したノードのブリッジング負荷を吸収しているエンジニアは、比例したストレスを受けます。
知識のギャップの伝播 — 退職したエンジニアが文書化されていないドメイン (文書化された_pct によって構成) を所有していた場合、それらのギャップは SimEvent ログに登録され、後続のインシデントで要因として表面化します。
employee_Departed SimEvent — エッジ スナップショット、出発時の中心性、r で発行されます。

割り当てられたチケット、およびインシデントの引き継ぎ。検索評価のための完全なグラウンドトゥルース。
そのため、ジョーダンが12日目に出発したとき、9日目の事件の検視には彼女について言及されていません。しかし、15 日目の事後分析では、「最近の退職に起因する認証サービスにおける事前の知識のギャップが、診断の遅れに寄与した」可能性があります。この文は、LLM 推論ではなく、実際の SimEvent に基づいています。
OrgForge には、クリーンなシミュレーション パスには一切触れずに、通常のシミュレーションの上に敵対的な動作を重ねるオプションの内部者脅威モジュールが含まれています。無効にすると (デフォルト)、完全に不活性になり、オーバーヘッドも追加出力もコード パスも変更されません。
有効にすると、指定された従業員は、異常な Git アクティビティ、時間外アクセス、Slack での感情の漂流、ワークステーションでのデータ ステージング、IDP 認証の異常など、複数の側面にわたって構成可能な脅威の動作を示します。すべての脅威テレメトリは、業界標準のログ形式 (JSONL、CEF、ECS、LEEF) で別の security_telemetry/ ディレクトリに書き込まれ、通常のシミュレーション出力から完全に分離されているため、検出エージェントはそれを見つけるために動作する必要があります。
このモジュールは、インサイダー脅威検出システムを構築および評価するために設計されています。出力形式に関係なく、グラウンド トゥルースは常に JSONL に保存されます。
👉 Insider Threat のリファレンス全文はこちらでお読みください。
flow.py は、シミュレーションの主要なエントリ ポイントです。 config/config.yaml は、組織構造、ペルソナ、品質プリセットの信頼できる唯一の情報源です。
シナリオ
コマンド
注意事項
Docker のすべて
ドッカーの構成
初回実行に推奨
残りはローカルの Ollama + Docker
ドッカーはmongodb orgforgeを構成します
.env に OLLAMA_BASE_URL を設定します
クラウドプリセット (AWS Bedrock)
ドッカーはmongodb orgforgeを構成します
.env で資格情報を設定し、Ollama をスキップします
オプション 1 — すべて

Docker 内のこと (推奨)
git clone https://github.com/aeriesec/orgforge
CD オルグフォージ
ドッカーの構成
最初の実行ではモデルが自動的にプルされます (接続に応じて約 5 ～ 8 分)。後続の実行は数秒で開始されます。モデルは名前付きボリュームにキャッシュされます。
シミュレーションが終了したら、後処理アーティファクト ジェネレーターを実行します。
Pythonメール_gen.py
Python post_sim_artifacts.py
出力は ./export/ に配置されます。
オプション 2 — ローカル Ollama、MongoDB のみの Docker
OLLAMA_BASE_URL=http://host.docker.internal:11434
次に:
ドッカーはmongodb orgforgeを構成します
Linux の注意: host.docker.internal には Docker Desktop、または docker-compose.yaml の extra_hosts: host-gateway エントリ (すでに含まれています) が必要です。
オプション 3 — クラウド プリセット (AWS Bedrock + OpenAI)
最高の出力品質。ドキュメント生成には Claude Sonnet、大量のワーカー呼び出しには Bedrock 上の Llama 3.1 8B、埋め込みには OpenAI text-embedding-3-large を使用します。
config.yaml でquality_preset: "cloud" を設定し、次のようにします。
# .env
AWS_ACCESS_KEY_ID=...
AWS_SECRET_ACCESS_KEY=...
AWS_DEFAULT_REGION=us-east-1
OPENAI_API_KEY=...
pip インストール boto3 langchain-aws openai
ドッカーはmongodb orgforgeを構成します
AWS EC2上で実行
安価な EC2 + Bedrock/OpenAI (GPU は不要)
t3.small は問題なく動作します。クラウド API が面倒な作業をすべて実行します。
EC2 インスタンス (Ubuntu または Amazon Linux) を起動し、Docker をインストールします。
git clone https://github.com/aeriesec/orgforge.git && cd orgforge
cp .env.example .env を実行し、資格情報を入力します。
config/config.yaml でquality_preset: "cloud" を設定します。
docker compose up --build -d mongodb orgforge
GPU インスタンス + 70B ローカル モデル
Llama 3.3 70B を完全にローカルで使用する場合は、Deep Learning AMI で g5.2xlarge または g5.12xlarge を使用します。 docker-compose.yaml の ollama サービスの下にある GPU デプロイ ブロックのコメントを解除し、quality_preset: "local_gpu" を設定してから d

オッカーは -d を構成します。
config/config.yaml が唯一の信頼できる情報源です。ほとんどのカスタマイズでは、Python を変更する必要はありません。
品質プリセット: " local_gpu " # local_gpu |雲
プリセット
プランナー
労働者
埋め込み
最適な用途
ローカルGPU
ラマ3.3:70b-instruct-q4_KM
ラマ3.1:8b-命令
mxbai-embed-large
忠実度の高いローカル実行
雲
クロード・ソネット (岩盤)
ラマ3.1:8b (岩盤)
テキスト埋め込み-3-大
最高の出力品質
主要な設定フィールド
フィールド
目的
会社名
生成されたすべての散文に挿入される
シミュレーション日数
シミュレーションの長さ (デフォルト: 22)
レガシーシステム
インシデント、チケット、ドキュメントで参照される不安定なシステム
crm
Salesforce と Zendesk のシミュレーション統合を有効/無効にする
スプリントチケット_テーマ
スプリント計画中に抽出されたチケット タイトルのプール
adhoc_confluence_topics
通常の日に生成される自発的な Wiki ページ
知識のギャップ
退職した従業員が固定的に存在し、その欠勤により初日から書類のギャップが生じる
org_lifecycle
シミュレーション中に発生する動的な離職と雇用 (以下を参照)
役割
シミュレーションの役割 (オンコール、インシデント指揮官、人事リーダー) を部門にマッピングします。
士気
減衰率、回復率、介入閾値
組織図 + リード
社内の全員と各部門の責任者
ペルソナ
書き方、ストレスレベル、専門家

[切り捨てられた]

## Original Extract

Synthetic corporate dataset generator for AI agent evaluation. - aeriesec/orgforge

GitHub - aeriesec/orgforge: Synthetic corporate dataset generator for AI agent evaluation. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
aeriesec
/
orgforge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
60 Commits 60 Commits .github/ workflows .github/ workflows config config eval eval scripts scripts src src tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile INSIDER_THREAT.md INSIDER_THREAT.md LICENSE LICENSE README.md README.md docker-compose.yaml docker-compose.yaml orgforge_hero.png orgforge_hero.png pyproject.toml pyproject.toml requirements-cloud.txt requirements-cloud.txt requirements-test.txt requirements-test.txt requirements.txt requirements.txt simulation.log.txt simulation.log.txt uv.lock uv.lock View all files Repository files navigation
A deterministic corporate simulator for generating ground-truth ecosystems and evaluating enterprise AI agents
OrgForge simulates weeks of realistic enterprise activity — Confluence pages, JIRA tickets, Slack threads, Git PRs, Zoom transcripts, Zendesk tickets, Salesforce records, emails, and server telemetry — grounded in an event-driven state machine so LLMs can't hallucinate facts out of sequence.
The dataset is the exhaust of a living simulation. Engineers leave mid-sprint, forcing deterministic incident handoffs, ticket reassignments, and CRM ownership lapses. Knowledge gaps surface when under-documented systems break. New hires build their internal network through simulated collaboration. Stress propagates through a live, weighted social graph. Every artifact reflects the exact state of the org at the moment it was written.
Option 1 — Everything in Docker
Option 2 — Local Ollama, Docker for MongoDB Only
When building AI agents that reason over institutional knowledge, you need a realistic corpus to test against. The only widely-used corporate dataset is the Enron email corpus — 25 years old, legally sensitive, and covering one company in crisis.
OrgForge generates that corpus from scratch, parameterized to any company, industry, or org structure. LLMs write the prose, but the facts — who was on-call, which ticket was open, when the incident resolved, who just left the team, and which customer SLA was breached — are strictly controlled by the state machine.
The central design bet: grounding LLM output in a deterministic event log makes the dataset actually useful for evaluating retrieval systems. You have ground truth about what happened, when, who was involved, and what the org's state was — so you can measure whether an agent surfaces the right context, not just plausible-sounding context.
Here's what a slice of a real simulation produces. An incident fires on Day 8:
slack/channels/engineering-incidents.json — the alert arrives first, timestamped to the millisecond the on-call pager fired:
{
"ts" : " 2026-03-10T14:23:07 " ,
"user" : " pagerduty-bot " ,
"text" : " 🔴 P1 ALERT: TitanDB latency spike — connection pool exhaustion under load. On-call: Jax. "
}
jira/IT-108.json — opened seconds later, facts pulled from the same SimEvent:
{
"id" : " IT-108 " ,
"type" : " incident " ,
"priority" : " P1 " ,
"title" : " TitanDB: latency spike " ,
"root_cause" : " connection pool exhaustion under load " ,
"assignee" : " Jax " ,
"reporter" : " system " ,
"opened" : " 2026-03-10T14:23:19 "
}
confluence/postmortems/IT-108.md — written the next day, linking the same root cause and PR:
This incident was triggered by connection pool exhaustion under sustained load, first surfaced in IT-108. The fix landed in PR #47 (merged by Sarah). A prior knowledge gap in TitanDB connection management — stemming from Jordan's departure on Day 12 — contributed to the delayed diagnosis.
Meanwhile, the datadog/metrics.jsonl time-series data reflects the exact latency spike, Zendesk support tickets from affected customers are automatically escalated to 'Urgent', Salesforce opportunities are flagged as 'at-risk', and end-of-month customer invoices ( invoices/ ) automatically apply SLA credits based on the incident's duration.
None of this is coincidence — it all traces back to one SimEvent that every downstream artifact reads from.
A default 22-day simulation produces:
OrgForge is not an LLM wrapper. Four interlocking systems enforce correctness.
👉 Read the full Architecture Deep-Dive here.
The most complex behaviour in the simulation. When an engineer departs mid-sprint, the following fires in order before that day's planning runs:
Incident handoff — active incidents assigned to the departing engineer are rerouted via Dijkstra escalation routing (while the node is still in the graph) to the next available person in the chain.
Ticket & CRM reassignment — orphaned JIRA tickets go to the dept lead. Salesforce accounts and open opportunities owned by the departed employee are flagged for reassignment, maintaining cross-domain ground truth.
Graph recompute — betweenness centrality is recalculated on the smaller graph. Engineers absorbing the departed node's bridging load receive a proportional stress hit.
Knowledge gap propagation — if the departed engineer owned undocumented domains (configured via documented_pct ), those gaps are registered in the SimEvent log and surface in subsequent incidents as contributing factors.
employee_departed SimEvent — emitted with edge snapshot, centrality at departure, reassigned tickets, and incident handoffs. Full ground truth for retrieval evaluation.
So when Jordan leaves on Day 12, the postmortem on Day 9's incident doesn't mention her. But the postmortem on Day 15 might: "A prior knowledge gap in auth-service, stemming from a recent departure, contributed to the delayed diagnosis." That sentence is grounded in a real SimEvent, not LLM inference.
OrgForge includes an optional insider threat module that layers adversarial behavior on top of the normal simulation — without touching any of the clean simulation paths. When disabled (the default), it is completely inert: no overhead, no additional output, no altered code paths.
When enabled, designated employees exhibit configurable threat behaviors across multiple surfaces: anomalous git activity, off-hours access, sentiment drift in Slack, data staging on their workstation, and IDP authentication anomalies. All threat telemetry is written to a separate security_telemetry/ directory in industry-standard log formats (JSONL, CEF, ECS, LEEF), keeping it cleanly isolated from the normal simulation output so detection agents must work to find it.
The module is designed for building and evaluating insider threat detection systems. Ground truth is always preserved in JSONL regardless of output format.
👉 Read the full Insider Threat reference here.
flow.py is the main simulation entry point. config/config.yaml is the single source of truth for org structure, personas, and quality presets.
Scenario
Command
Notes
Everything in Docker
docker compose up
Recommended for first run
Local Ollama + Docker for the rest
docker compose up mongodb orgforge
Set OLLAMA_BASE_URL in .env
Cloud preset (AWS Bedrock)
docker compose up mongodb orgforge
Set credentials in .env , skip Ollama
Option 1 — Everything in Docker (Recommended)
git clone https://github.com/aeriesec/orgforge
cd orgforge
docker compose up
First run pulls models automatically (~5–8 min depending on your connection). Subsequent runs start in seconds — models are cached in a named volume.
When the simulation finishes, run the post-processing artifact generators:
python email_gen.py
python post_sim_artifacts.py
Output lands in ./export/ .
Option 2 — Local Ollama, Docker for MongoDB Only
OLLAMA_BASE_URL=http://host.docker.internal:11434
Then:
docker compose up mongodb orgforge
Linux note: host.docker.internal requires Docker Desktop, or the extra_hosts: host-gateway entry in docker-compose.yaml (already included).
Option 3 — Cloud Preset (AWS Bedrock + OpenAI)
Best output quality. Uses Claude Sonnet for document generation, Llama 3.1 8B on Bedrock for high-volume worker calls, and OpenAI text-embedding-3-large for embeddings.
Set quality_preset: "cloud" in config.yaml , then:
# .env
AWS_ACCESS_KEY_ID=...
AWS_SECRET_ACCESS_KEY=...
AWS_DEFAULT_REGION=us-east-1
OPENAI_API_KEY=...
pip install boto3 langchain-aws openai
docker compose up mongodb orgforge
Running on AWS EC2
Cheap EC2 + Bedrock/OpenAI (no GPU required)
A t3.small works fine — the cloud APIs do all the heavy lifting.
Launch an EC2 instance (Ubuntu or Amazon Linux) and install Docker
git clone https://github.com/aeriesec/orgforge.git && cd orgforge
cp .env.example .env and fill in your credentials
Set quality_preset: "cloud" in config/config.yaml
docker compose up --build -d mongodb orgforge
GPU Instance + 70B Local Models
For Llama 3.3 70B entirely locally, use a g5.2xlarge or g5.12xlarge with the Deep Learning AMI. Uncomment the GPU deploy block under the ollama service in docker-compose.yaml , set quality_preset: "local_gpu" , then docker compose up -d .
config/config.yaml is the single source of truth. No Python changes are needed for most customizations.
quality_preset : " local_gpu " # local_gpu | cloud
Preset
Planner
Worker
Embeddings
Best For
local_gpu
llama3.3:70b-instruct-q4_KM
llama3.1:8b-instruct
mxbai-embed-large
High-fidelity local runs
cloud
Claude Sonnet (Bedrock)
llama3.1:8b (Bedrock)
text-embedding-3-large
Best output quality
Key Config Fields
Field
Purpose
company_name
Injected into all generated prose
simulation_days
Length of the simulation (default: 22)
legacy_system
The unstable system referenced in incidents, tickets, and docs
crm
Enable/disable Salesforce and Zendesk simulation integrations
sprint_ticket_themes
Pool of ticket titles drawn during sprint planning
adhoc_confluence_topics
Spontaneous wiki pages generated on normal days
knowledge_gaps
Static departed employees whose absence creates documentation gaps from day one
org_lifecycle
Dynamic departures and hires that occur during the simulation (see below)
roles
Maps simulation roles (on-call, incident commander, HR lead) to departments
morale
Decay rate, recovery rate, intervention threshold
org_chart + leads
Everyone in the company and who runs each department
personas
Writing style, stress level, and expert

[truncated]
