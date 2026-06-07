---
source: "https://github.com/ninoxAI/nightwatch"
hn_url: "https://news.ycombinator.com/item?id=48438180"
title: "Show HN: Nightwatch, The open-source, read-only AI SRE"
article_title: "GitHub - ninoxAI/nightwatch: Open-source, local-first, read-only AI SRE: clusters alert storms, investigates root cause over your live systems, proposes human-gated fixes. · GitHub"
author: "egorferber"
captured_at: "2026-06-07T21:30:58Z"
capture_tool: "hn-digest"
hn_id: 48438180
score: 2
comments: 0
posted_at: "2026-06-07T20:24:36Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Nightwatch, The open-source, read-only AI SRE

- HN: [48438180](https://news.ycombinator.com/item?id=48438180)
- Source: [github.com](https://github.com/ninoxAI/nightwatch)
- Score: 2
- Comments: 0
- Posted: 2026-06-07T20:24:36Z

## Translation

タイトル: Show HN: Nightwatch、オープンソースの読み取り専用 AI SRE
記事のタイトル: GitHub - ninoxAI/nightwatch: オープンソース、ローカルファースト、読み取り専用 AI SRE: アラート ストームをクラスター化し、ライブ システムの根本原因を調査し、人間による修正を提案します。 · GitHub
説明: オープンソース、ローカルファースト、読み取り専用 AI SRE: アラート ストームをクラスター化し、稼働中のシステムの根本原因を調査し、人間による修正を提案します。 - ninoxAI/ナイトウォッチ
HN テキスト: nightwatch は、モニタリングの上にあるローカルファーストの読み取り専用レイヤーです。アラート ストームをインシデントにグループ化し、ノイズの多いチェックにフラグを立て、稼働中のシステムを調査できるエージェントを備えています。たとえば、次のようにすることができます。インシデントからエージェントに直接ジャンプします。この週末プロジェクトを行う理由は、Kubernetes のアップグレードが失敗し、ある時点でロールバックが不可能になったため、いくつかの問題が重なっている間に夜間にライブで修正する必要があったためです。私たちはオンプレミスやいくつかの Kubernetes クラスターなど、さまざまなシステムを多数実行しています。そのような状況では、実際に何がどこで壊れているのかを把握することにほとんどの時間を費やします。そこで、各システムの暗闇の中に「脳」と会話できる目があれば、とても素晴らしいだろうと思いました。そこで、赤ちゃんフクロウをそれぞれの環境に置くというアイデアです。各フクロウはシステムが存在する場所で実行され、その環境の資格情報をローカルに保ち、中央の脳に発信方向にのみダイヤルするため、本番環境への受信ホールはありません。これにより一連の読み取り専用スキルが公開され、エージェントはそれらを使用して証拠を収集し、根本原因の仮説を立てるため、オンコール エンジニアはゼロからではなく有利なスタートを切ることができます。今のところ読み取り専用です。私はまだ本番環境ではそれを信頼していません。正直に言って、あなたもそうすべきではありません。 llocal-first はセルフホスティングを容易にし、資格情報を側に置いておきます。クラスタリングと

推奨事項は、llm をまったく使用せずに完全にオフラインで実行されます。エージェントにはツールを呼び出す llm が必要です。完全にオフラインにしたい場合は、リモート ツールまたはセルフホスト ツール (ollam など) をポイントすることができます。非セルフホスターの場合: すべてのリモート llm 呼び出しの前に、nightwatch は実際のシークレット (復元不可能) を削除し、ips、ホスト名、パスなどの識別子を可逆的なプレースホルダーと交換します。そのため、モデルはマスクされたデータのみを参照し、実際の値は提案されたコマンドとツール呼び出しでのみ復元されます。あなたのシステムで試してみていただければ幸いです。

記事本文:
GitHub - ninoxAI/nightwatch: オープンソース、ローカルファースト、読み取り専用 AI SRE: アラート ストームをクラスター化し、稼働中のシステムの根本原因を調査し、人間による修正を提案します。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ニノックスAI
/
夜回り
公共
通知
あなたは署名している必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
87 コミット 87 コミット .github .github デプロイ/ k8s デプロイ/ k8s ドキュメント ドキュメント ラボ ラボ src/ ninoxai src/ ninoxai テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.embeddings Dockerfile.embeddings ライセンス ライセンス通知 README.md README.md alembic.ini alembic.ini docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
オープンソースの読み取り専用 AI SRE。
ninoxAI は、本番環境に一切触れることなく、アラートの嵐をインシデントに変え、稼働中のシステムの根本原因を調査し、人間が承認した修正を提案します。
クイックスタート
・AI SRE
・デモラボ
· ドキュメント
・不和
監視していると、何かが壊れていることがわかります。午前 3 時に 1 回の停止に対して 50 件のアラートが通知され、難しい部分はあなたに任せられます。
何が壊れたのか、なぜ壊れたのか、そして次に何をすべきか?
ninoxAI は、その質問に答える、ローカルファーストでモニタリングに依存しない薄い AI SRE レイヤーです。 Checkmk、Prometheus、Icinga2、Zabbix、Webhook、Docker、Kubernetes、AWS、Grafana、GitHub、Git、およびプレーン VM の上に位置します。
🌊 アラートの洪水をインシデントに変換します。症状ごとに 1 ページではなく、停止ごとに 1 つのインシデントが「N 個のツールによって確認」されます。
🔇 うるさい小切手 (バタバタする、過敏になる、何も行われない) を証拠とともに見つけます。
🤖 根本原因を調査 — ツールを呼び出す AI エージェントが稼働中のシステムを読み取り、根本原因の仮説を立てます。
🧰 機密扱いの修正を提案します。コピー＆ペースト可能で、リスクと爆発範囲によってランク付けされ、人間がゲートできるようにします。
ninoxAI は観察し、理由付けし、推奨します - 決して元に戻ることはありません

何でも実行します。コマンドの実行、アラートの確認、しきい値の変更、運用環境へのライトバックはありません。すべての修正は、人間が承認するコピー可能な成果物です。ゲート付きで管理された修復がロードマップに載っています。無条件自動実行はそうではありません。
60 秒以内にお試しください — LLM、API キーなし、完全オフライン:
cp .env.example .env # set NINOXAI_SECRET_KEY (ワンライナーがファイル内にあります)
docker compose up --build # → http://127.0.0.1:8765
# ライブモニタリングはありませんか?合成アラート ノイズを優先順位付けする様子を見てください。
docker compose exec ninoxai ninoxai 生成モック
docker compose exec ninoxai ninoxai import data/mock_alerts.json
docker compose exec ninoxai ninoxai 再処理
# → /recommendations は、根拠のあるしきい値とフラッピングの修正を表示するようになりました
ローカル Python インストール (開発用)
python -m venv .venv && ソース .venv/bin/activate # Windows: .venv\Scripts\Activate.ps1
pip install -e " .[dev,embeddings] "
Python -m ninoxai 生成モック
python -m ninoxai import data/mock_alerts.json
Python -m ninoxaiの再処理
python -m ninoxai サーブ # → http://127.0.0.1:8765
次に、AI SRE を起動します。ninoxAI をツール呼び出し LLM (Anthropic / OpenAI / Mistral / ローカル Ollama) にポイントし、システムを直接または ninox ランナー経由で接続します (下記)。完全なエンドツーエンドのシナリオ (実際に障害が発生したワークロードに対する実際の監視ツールとライブ調査機能 (Docker/Kubernetes/host/AWS/Grafana/GitHub)) は、 lab/ にあります。
取り込み → 正規化 → クラスタ → スコアノイズ → 推奨 → ダッシュボード
↓
エージェント、読み取り専用の根本原因調査者
ステージ
何が起こるか
摂取する
読み取り専用アダプターは、各ソース + JSON/CSV インポートから非 OK アラートを取得します。
正規化する
すべてのソースを 1 つのスキーマ + メッセージ フィンガープリントにマッピングします。
クラスター
ホスト/サービス/重大度/時間枠ごとにグループ化します。セマンティック埋め込みはオプションです。
いいえ

せ
周波数、ACK レート、チケット レート、ショート リカバリ、フラッピング → 0 ～ 1 のスコア。
勧める
理論的根拠と証拠を備えたルールベースのチューニング推奨事項。
調査する
ツールを呼び出す LLM は、生きた証拠 → 根本原因仮説 + 分類された修正を収集します。
ツール間の相関関係: すべてのツールで同じ障害が発生します。インシデント ビューでは、(ホスト、重大度、時間枠) を共有するクラスターが 1 つのインシデントにグループ化されます (「N ツールで確認済み」)。読み取り専用でマージはありません。
ninoxAIの卓越した機能。ツール呼び出し LLM は、読み取り専用機能の型指定されたホワイトリスト (ネイティブ関数呼び出しの ReAct ループ - 理由→動作→観察) を駆動し、生きた証拠から根本原因仮説を構築し、人間が承認する分類された修正を提案します。
すべてのアクションは、読み取り専用、可逆、不可逆 + スコープ (爆発範囲) に分類されます。不明な場合は不可逆的な強制が行われます。決して黙って自動化されることはありません。
事前接地: エージェントは環境の簡潔な概要から開始するため、再検出ではなく診断が行われます。
強化: 信頼できないログ/差分は注入から保護され、秘密は一方向にスクラブされ、主張が証拠に裏付けられていない場合は接地ゲートによって信頼が制限されます。
エージェント コンソール ( /agent ) または CLI からライブ ストリーミングで実行します。 → 捜査官内部事情
🦉 分散型 ninox — エージェントの目、どこにいても
エージェントは、直接アクセスできないシステムを調査できます。 ninox は、1 つの環境 (クラスター、VPC、オンプレミス セグメント) 内に存在し、その環境の認証情報をローカルに保持し、頭脳にダイヤル ホームする、送信専用のシン ランナーです。受信ファイアウォール ホールはありません。これは、脳がローカルであるかのように呼び出す読み取り専用機能サーフェスをアドバタイズします。
┌───────────┐ ┌───────────┐
│ にん

oxAI Brain │ ◀─ 発信専用─── │ ninoxランナー │
│ ダッシュボード · API │ (ninox ダイヤル │ k8s/Docker/ 内 │
│ 事件・RCA │ ホーム;インバウンドなし │ AWS/オンプレミス/VM │
│ AI SRE 調査員│ ファイアウォールの穴）│ 資格情報の保持│
━━━━━━━━━┘ ◀── 読み取り専用証拠 │ ローカル │
━━━━━━━━━┘
機能は環境によって自動的に選択されます。つまり、1 つのバイナリと、そのバイナリが配置されるボックスに適したツールです。接続されたニノックスがフクロウの議会 ( /parliament ) に現れます。 → 導入とオンプレミス
すべてのアダプターは読み取り専用です。ACK、ダウンタイム、ライトバックはありません。 UI ( /connections ) で構成され、資格情報は Fernet で暗号化されます。
AI SRE にスタック (Jira、Sentry、Postgres など) を読み取るように教えたいですか?任意の MCP サーバーにポイントしたり、Python 機能プラグインを作成したり、ランナー プロトコル経由でツールを公開したりできます。すべての外部ツールは同じ安全シェル (名前空間、インジェクション スキャン、分類強制) を通じて実行されます。 → 機能の拡張
デフォルトはテンプレートです。完全にオフラインです。LLM、ネットワーク、API キー、追跡はありません。概要/推奨事項についてはそのまま使用できますが、意図的にエージェントを駆動することはできません (ツールの呼び出しが必要です)。役割ごとにリモートを選択します。大量の要約には安価なモデル、まれな調査には強力なモデルです。
すべてのリモート呼び出しの前にリダクションとシークレットのスクラビングが実行されます。ホスト名、IP、UUID、電子メール、パスは決定的なプレースホルダーとなり、提案されたコマンドでのみ復元されます。資格情報は一方向にスクラブされ、返されません。 → 技術アーキテクチャ
完全な CLI リファレンス、テスト セットアップ、lint ルールは docs/development.md にあります。
すべての投稿者はフクロウです。 🦉 プルリクエスト、コネクタ適応

ター、機能プロバイダー、バグ レポートはすべて大歓迎です。 CONTRIBUTING.md を参照してください。
コミュニティ: Discord で議会に参加します。
ninoxAI は、Apache License 2.0 に基づく完全なオープン ソースであり、オープン プロジェクトでもクローズ プロジェクトでも同様に無料で使用でき、自己ホスト、フォーク、ビルドが可能です。
フクロウは観察します。人間が決めるのです。 🦉
オープンソース、ローカルファースト、読み取り専用 AI SRE: アラート ストームをクラスター化し、稼働中のシステムの根本原因を調査し、人間による修正を提案します。
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open-source, local-first, read-only AI SRE: clusters alert storms, investigates root cause over your live systems, proposes human-gated fixes. - ninoxAI/nightwatch

nightwatch is a local-first, read-only layer on top of your monitoring. it groups alert storm into incidents, flags noisy checks and has an agent that can investigate for you live systems. You can e.g. jump from the incident into the agent directly. the reason for this weekend project is that we had a kubernetes upgrade that went wrong, and at some point a rollback wasn't possible anymore, so it had to be fixed live during the night while several problems came together. We run a lot of different systems, on-prem and several Kubernetes clusters, and in a situation like that you spend most of the time just figuring out what is actually broken and where. So i thought that it would be pretty cool to have eyes in the dark in each system that can talk to your "brain". so the idea is to put a baby owl into each environment. Each owl runs where the systems live, keeps that environment's credentials local, and only dials outbound to a central brain, so there is no inbound hole into prod. It exposes a set of read-only skills, and the agent uses them to gather evidence and form a root-cause hypothesis, so the on-call engineer starts with a head start instead of from zero. read-only for now, i don't trust it near prod yet and honestly neither should you. llocal-first for easy self-hosting and to keep credentials on your side. the clustering and recommendations run fully offline with no llm at all. the agent needs a tool-calling llm, you can point it at a remote one, or self-host one (ollama etc.) if you want to stay fully offline. for non selfhosters: before every remote llm call, nightwatch strips real secrets (unrestorable) and swaps identifiers like ips, hostnames and paths for reversible placeholders, so the model only sees masked data while real values are restored only in the proposed commands and tool calls Would love if you try it in your Systems

GitHub - ninoxAI/nightwatch: Open-source, local-first, read-only AI SRE: clusters alert storms, investigates root cause over your live systems, proposes human-gated fixes. · GitHub
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
ninoxAI
/
nightwatch
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
87 Commits 87 Commits .github .github deploy/ k8s deploy/ k8s docs docs lab lab src/ ninoxai src/ ninoxai tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.embeddings Dockerfile.embeddings LICENSE LICENSE NOTICE NOTICE README.md README.md alembic.ini alembic.ini docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml View all files Repository files navigation
The open-source, read-only AI SRE.
ninoxAI turns alert storms into incidents, investigates root cause over your live systems, and proposes human-approved fixes — without ever touching production.
Quickstart
· AI SRE
· Demo lab
· Docs
· Discord
Your monitoring tells you something broke. It pages you at 3am with fifty alerts for one outage and leaves the hard part to you:
What broke, why did it break, and what should we do next?
ninoxAI is a thin, local-first , monitoring-agnostic AI SRE layer that answers that question. It sits above Checkmk, Prometheus, Icinga2, Zabbix, webhooks, Docker, Kubernetes, AWS, Grafana, GitHub, Git and plain VMs, and:
🌊 Turns alert floods into incidents — one incident per outage, "confirmed by N tools" , instead of one page per symptom.
🔇 Finds the noisy checks — flapping, over-sensitive, never-actioned — with evidence.
🤖 Investigates root cause — a tool-calling AI agent reads your live systems and forms a root-cause hypothesis.
🧰 Proposes classified fixes — copy-pasteable, ranked by risk and blast radius, for a human to gate.
ninoxAI observes, reasons, and recommends — it never executes anything. No commands run, no alerts acked, no thresholds changed, no write-back to production. Every fix is a copyable artifact a human approves. Gated, governed remediation is on the roadmap; unconditional auto-execute is not.
Try it in 60 seconds — no LLM, no API keys, fully offline:
cp .env.example .env # set NINOXAI_SECRET_KEY (one-liner is in the file)
docker compose up --build # → http://127.0.0.1:8765
# No live monitoring? Watch it triage synthetic alert noise:
docker compose exec ninoxai ninoxai generate-mocks
docker compose exec ninoxai ninoxai import data/mock_alerts.json
docker compose exec ninoxai ninoxai reprocess
# → /recommendations now shows reasoned threshold + flapping fixes
Local Python install (for development)
python -m venv .venv && source .venv/bin/activate # Windows: .venv\Scripts\Activate.ps1
pip install -e " .[dev,embeddings] "
python -m ninoxai generate-mocks
python -m ninoxai import data/mock_alerts.json
python -m ninoxai reprocess
python -m ninoxai serve # → http://127.0.0.1:8765
Then light up the AI SRE: point ninoxAI at a tool-calling LLM (Anthropic / OpenAI / Mistral / a local Ollama) and connect your systems — either directly or via a ninox runner (below). The full end-to-end scenario — real monitoring tools and live investigator capabilities (Docker/Kubernetes/host/AWS/Grafana/GitHub) against a genuinely failing workload — lives in lab/ .
ingest → normalize → cluster → score noise → recommend → dashboard
↓
agentic, read-only root-cause investigator
Stage
What happens
ingest
Read-only adapters pull non-OK alerts from each source + JSON/CSV import.
normalize
Maps every source onto one schema + message fingerprint.
cluster
Groups by host / service / severity / time-window. Semantic embeddings optional.
noise
Frequency, ack-rate, ticket-rate, short-recovery, flapping → one 0–1 score.
recommend
Rule-based tuning recommendations with rationale + evidence.
investigate
A tool-calling LLM gathers live evidence → root-cause hypothesis + classified fixes.
Cross-tool correlation: the same fault fires in every tool. The Incidents view groups clusters that share (host, severity, time-window) into one incident — "confirmed by N tools" — read-only, no merge.
ninoxAI's standout capability. A tool-calling LLM drives a typed allowlist of read-only capabilities (a ReAct loop on native function-calling — reason → act → observe), builds a root-cause hypothesis from live evidence, and proposes classified fixes a human approves.
Every action is classified read_only · reversible · irreversible + a scope (blast radius). Unknown coerces to irreversible — never silently auto.
Pre-grounded: the agent starts with a compact brief of your environment, so it diagnoses instead of rediscovering.
Hardened: untrusted logs/diffs are injection-shielded, secrets are one-way scrubbed, and a grounding gate caps confidence when claims aren't backed by evidence.
Run it live-streaming in the agent console ( /agent ) or from the CLI. → Investigator internals
🦉 Distributed ninoxes — the agent's eyes, anywhere
The agent can investigate systems it can't reach directly. A ninox is a thin, outbound-only runner that lives inside one environment (cluster, VPC, on-prem segment), holds that environment's credentials locally , and dials home to the brain — no inbound firewall hole . It advertises a read-only capability surface the brain calls as if local.
┌────────────────────┐ ┌────────────────────┐
│ ninoxAI brain │ ◀── outbound only ─── │ ninox runner │
│ dashboard · API │ (the ninox dials │ inside k8s/Docker/ │
│ incidents · RCA │ home; no inbound │ AWS/on-prem/VM │
│ AI SRE investigator│ firewall hole) │ credentials stay │
└────────────────────┘ ◀── read-only evidence │ local │
└────────────────────┘
Capabilities self-select by environment — one binary, the right tools for the box it lands on. Connected ninoxes show up in the Parliament of Owls ( /parliament ). → Deployment & on-prem
All adapters are read-only — no ack, no downtime, no write-back. Configured in the UI ( /connections ), credentials Fernet-encrypted.
Want to teach the AI SRE to read your stack (Jira, Sentry, Postgres…)? Point it at any MCP server , write a Python capability plugin , or expose tools via the runner protocol — every external tool runs through the same safety shell (namespaced, injection-scanned, classification-coerced). → Extending capabilities
Default is template — fully offline : no LLM, no network, no API keys, no tracking. It works out of the box for summaries/recommendations but deliberately can't drive the agent (that needs tool-calling). Pick a remote per role — a cheap model for high-volume summaries, a strong one for the rare investigation:
Redaction + secret-scrubbing run before every remote call — hostnames, IPs, UUIDs, emails, paths become deterministic placeholders, restored only in proposed commands; credentials are one-way scrubbed and never returned. → Technical architecture
Full CLI reference, test setup, and lint rules live in docs/development.md .
Every contributor is an Owl. 🦉 Pull requests, connector adapters, capability providers, and bug reports are all welcome — see CONTRIBUTING.md .
Community: Join the parliament on Discord .
ninoxAI is fully open source under the Apache License 2.0 — free to use, self-host, fork, and build on, in open or closed projects alike.
The owl observes; the human decides. 🦉
Open-source, local-first, read-only AI SRE: clusters alert storms, investigates root cause over your live systems, proposes human-gated fixes.
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
