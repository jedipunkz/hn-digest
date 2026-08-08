---
source: "https://github.com/GlediLami/kubetective"
hn_url: "https://news.ycombinator.com/item?id=49225051"
title: "KubeTective: Deterministic Incident Investigation for Kubernetes"
article_title: "GitHub - GlediLami/kubetective: Deterministic root causes for every Kubernetes incident you don't want to debug twice · GitHub"
author: "Gledis"
captured_at: "2026-08-08T20:17:56Z"
capture_tool: "hn-digest"
hn_id: 49225051
score: 1
comments: 0
posted_at: "2026-08-08T19:29:09Z"
tags:
  - hacker-news
  - translated
---

# KubeTective: Deterministic Incident Investigation for Kubernetes

- HN: [49225051](https://news.ycombinator.com/item?id=49225051)
- Source: [github.com](https://github.com/GlediLami/kubetective)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T19:29:09Z

## Translation

タイトル: KubeTective: Kubernetes の決定論的インシデント調査
記事のタイトル: GitHub - GlediLami/kubetective: 二度デバッグしたくないすべての Kubernetes インシデントの決定的な根本原因 · GitHub
説明: 2 回デバッグしたくないすべての Kubernetes インシデントの決定的な根本原因 - GlediLami/kubetective

記事本文:
GitHub - GlediLami/kubetective: 二度デバッグしたくないすべての Kubernetes インシデントの決定的な根本原因 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
グレディラミ
/
クベ探偵
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
46 コミット 46 コミット .githooks .githooks .github .github Formula Formula cmd cmd デプロイ デプロイ ハック ハック 内部 内部 pkg/ api pkg/ api レポート レポート シナリオ シナリオ

iOS サイト site .dockerignore .dockerignore .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md CODEOWNERS CODEOWNERS COTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md Demon.gif Demon.gif go.mod go.mod go.sum go.sum kubectl-investigate.yaml kubectl-investigate.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
KubeTective: Kubernetes インシデント調査エンジン
KubeTective は、優れた SRE と同じ方法で Kubernetes インシデントを調査します。たまる
ターゲットに関する事実を収集し、タイムラインと証拠グラフを構築し、ランク付けされたデータを生成します
説明可能で決定論的なスコアを持つ仮説と、その理由を正確に示します
考えていることを考えます。あらゆる自信のポイントは、目に見えるリストによって裏付けられています。
証拠。
CLI、kubectl プラグイン、REST サーバー、または MCP サーバーとして実行されます。毎
調査は記録されるため、再生、監査、ベンチマークとして使用できます。
ドキュメント: gledilami.github.io/kubetective · 変更履歴 · 貢献
上のデモは、記録されたインシデントの再生です
必須です。
kubectl Investigator <target> --since=30m は、「すべて問題ないようです?」というメッセージを返します。
証拠付きのランク付けされた評決の瞬間 (上の GIF は全文です)
出力;これは 5 秒バージョンです):
$ kubectl 調査デプロイ/チェックアウト --since=30m
╭─────────────────────╮
│ インシデント: デプロイメント/本番/チェックアウト │
│ ステータス: OOMKILLED │
│ 重大度: 高 │
│ 信頼度: 100 │
╰───────────────────╯
根本原因
記憶が枯渇する

on: コンテナが OOMKilled 19 回で終了しました (メモリ制限 1Gi) - 19 回の再起動
推奨事項
デプロイメント/本番/チェックアウトを正常な最後のリビジョンにロールバックします [MEDIUM]
その出力は、記録されたインシデント シナリオ/展開後の oom が再生されたものです。
エンジンを通して。 11 個のアナライザーで OOM キル、クラッシュ ループ、イメージ プルをカバー
障害、スケジューリング障害、ノードプレッシャー、プローブ障害、PVC の問題、
サービス セレクターの不一致、最大 HPA、DNS 障害、構成
回帰 (Git/GitOps、コミットまで)。
5 分でお試しいただけます。壊れたクラスターは必要ありません
git clone https://github.com/GlediLami/kubetective.git && cd kubetective
ビルドする
bin/kubetective restart scenarios/oom-after-deploy/record.jsonl # 実際のインシデントを再現
bin/kubetective ベンチマーク # 16 のシナリオとグラウンド トゥルースの比較
bin/kubetective 評価 # マークダウン精度レポート
シナリオに記録されたすべてのインシデントは、自己完結型の決定論的なデモです
すべてのアナライザーの回帰ゲートとしても機能します。
KubeTective と現状の比較
kubectl + grep + ドキュメント
LLM チャット
クベテクティブ
事実を収集して関連付けます
✗ あなたは掘る
✗ あなたが貼り付けます
✓ 自動 (k8s、Prometheus、Loki、Git/GitOps)
自信を持ってランク付けされた仮説
✗
〜雰囲気
✓ グラウンドトゥルースに対して校正済み
すべての評決は証拠を示します
✗
まれに
✓ 行ごとのスコアの内訳
録画されたリプレイで検証可能
✗
✗
✓ JSONL レコード、決定論的に再実行
変化検出（直前に何が変化したか）
✗
✗
✓ git コミット、GitOps ドリフト、ポッドの変更
シナリオスイートで回帰テスト済み
✗
✗
✓ CI の kubetective ベンチマーク
承認を得た安全な修復
✗
✗
✓ 読み取り専用プレビュー → 明示的 --yes
承認を得た安全な修復
✗
✗
✓ 読み取り専用プレビュー → 明示的 --yes
オフラインで動作し、API キーもテレメトリもありません
✓
✗
✓
アン・h

AI トリアージ ツール (k8sgpt など) に関する唯一のメモ
クラスターをスキャンし、
LLM での失敗について説明します。これらはトリアージには本当に役立ちますが、
中心的な違いは、どのような種類のアーティファクトを生成するかです。
彼らが輝いていて、私たちが（まだ）輝いていないところ: クラスター全体のスキャン、より大きなアナライザー
カタログ、LLM バックエンドの幅、および継続的な監視のためのオペレーター —
すべて私たちのロードマップにあり、私たち側からすべて再現可能です。
KubeTective の判定は、LLM ではなく、ルールベースのエンジンから得られます。それぞれの仮説
スコアは、行ごとに読み取ることができる重み付けされた証拠用語の合計です。自信というのは
シナリオのベンチマークに対して調整済み (グラウンド トゥルースを含む 16 件の記録されたインシデント)、
そして、キャリブレーションは、採用される前に、1 つずつ検証されます。オプションの
LLM レイヤーは、エンジンの判定を平易な言葉で説明するだけです。それは決して変わることはできない
スコアを付けたり、原因を発明したり、アクションを提案したりできます。確定的な出力は同じことを意味します
インシデントは、CI、デモ、および LLM チャットのリプレイで同じ評決を生成します
回帰テストにはできません。
証拠に基づいた調査。 11 個のアナライザーで最も一般的な障害をカバー
モード: OOM キル、クラッシュ ループ、イメージ プルの失敗、スケジューリング、および
スケジュール不能なポッド、ノードのプレッシャー、liveness および readiness プローブの障害、PVC
問題、サービス セレクターの不一致、最大 HPA、DNS 障害、構成
回帰 (Git および GitOps)。
コレクター。 Kubernetes (ポッド、デプロイメント、イベント、PVC、サービス、HPA、
coreDNS)、Prometheus (メモリメトリックの裏付け)、Git (
マニフェスト)、および GitOps (Flux Kustomization / HelmRelease 、ArgoCD アプリケーション)。
説明可能なスコアリング。あらゆる仮説には証拠の内訳が含まれており、
スコアは、Leave-one-out 検証を使用してシナリオ スイートに対して調整されます。
タイムライン、証拠グラフ、

そして変化検出。何が起こったのか、どのような順序で起こったのか、
何が何を所有し、事件の直前に何が変わったのか。
記録して再生します。すべての調査が追加されます
~/.kubetective/incidents/ として JSONL として。再生し、エンジンのバージョンを比較し、監査します。
シナリオベンチマークと評価レポート。 kubetective ベンチマークは、
すべてのアナライザーの回帰ゲート。 kubetective Evaluate はマークダウンをレンダリングします
評価レポート (シナリオごと、カテゴリごとの精度、キャリブレーション、
誤検知チェック）CI に適しています。
安全な修復。アクション (ロールバック、再起動) は読み取り専用でプレビューされ、
人間による明示的な承認後にのみ適用されます。適用するたびに監査レコードが書き込まれます。
RESTサーバーとMCPサーバー。 HTTP API と MCP の背後にある同じパイプライン
読み取り専用ツールを備えた stdio サーバー。
オプションの LLM 説明者。ダイジェストのみ、編集され、制限されています。どれでも動作します
OpenAI 互換エンドポイント (OpenAI、Ollama、vLLM、llama.cpp)。
kubectlプラグイン。 kubectl-investigate という名前のバイナリを PATH にドロップします。
そして、 kubectl Investigator <resource> を実行します。
シナリオを追加します。 kubetective Record --from-live を使用して新しいインシデントを記録します。
グラウンド トゥルースを追加すると、デモとベンチマークの両方のケースになります。
アナライザーを強化します。 11 個のアナライザーから 1 つを選択し、誤検知を見つけます
スイートに対してスコアを修正し、kubetective ベンチマークで証明してもらいます。
出力を改善します。レンダラーはプレーンな ANSI ボックス描画です。新しい出力
フォーマット ( --format=json が存在します。sarif または Slack レンダラーが開いています)。
新しい証拠ソース (Datadog、Grafana Cloud など) を既存の証拠ソースの背後に接続します。
コレクターインターフェイス。
デザインルールについては CONTRIBUTING.md を参照してください。
brew インストール gledilami/kubetective/kubetective
これにより、kubetective と kubectl-investigate (kubectl
プラグイン）。 Homebrew は homebrew-kubetective タップで式を自動的に見つけます。
余分なものはありません

ステップ。
github.com/GlediLami/kubetective/cmd/kubetective@latest をインストールしてください
ソースから
Go 1.26 以降と Kubernetes クラスターが必要です (ライブ調査用。
ベンチマークはそれなしで実行されます)。
git クローン https://github.com/GlediLami/kubetective.git
CDクベテクティブ
make ビルド # bin/kubective + bin/kubectl-investigate
make install # kubetective -> ~/.local/bin/kubetective
make install-plugin # kubectl-investigate -> ~/.local/bin (kubectl プラグイン)
システム全体にインストールするには、make install で PREFIX=/usr/local を指定します。
kubectl は、 PATH 上で kubectl-* という名前のプラグインを自動検出します。
インストールプラグインを作成する
kubectl Investigator pod/checkout-7f84c9 # 同じ: kubetective Investigator ...
クイックスタート
クラッシュ ループ デプロイメントを調査します (現在の kubeconfig コンテキストを使用します)。
kubectl のデプロイメント/チェックアウトの調査 --since=30m
# またはプラグインなし:
kubetective デプロイメント/チェックアウトを調査します --since=30m
ターゲット フォーム: pod/<name> 、deployment/<name> 、namespace/<name> (またはベア
name (デフォルトはポッド)、オプションで --namespace とウィンドウでスコープ指定
( --since=2h 、デフォルトは 30 分)。
kubetective 展開/チェックアウトを調査します \
--since=30m \
--prometheus-url=http://localhost:9090 \ # メトリックの確証
--loki-url=http://localhost:3100 \ # 証拠のログ (Loki)
--git-repo= ~ /code/checkout-manifests # マニフェストに触れるコミット
コマンドリファレンス
コマンド
説明
kubetective 調査 <リソース>
調査を実行します (フラグ: --since 、 --namespace 、 --no-logs 、 --format=json 、 --prometheus-url 、 --loki-url 、 --git-repo 、 --llm* )
kubetective 再生 <インシデント ID>
現在のエンジンを使用して記録された調査を再実行します (決定論的)
キューブ探偵事件
記録されたインシデント ID を新しい順にリストします。
kubetective インシデントに類似した <id> [--cluster <id>]
過去の類似事件を探す（事件発生）

nt メモリ、Jaccard オーバーラップ。 --cluster はルックアップの範囲を 1 つのクラスターに設定します)
kubetective アラート <pagerduty|grafana|slack>
Webhook アラート ペイロードから調査する (stdin または --file ; API キーはゼロ)
キューブ探偵
ヘルスチェック (バージョン、構成ファイル、キャリブレーション、クラスター接続、インシデント ストア、Prometheus/Loki の到達可能性)。失敗するとゼロ以外で終了します
kubetective アクション <インシデント ID>
修復アクションのプレビュー (読み取り専用; --apply <id> --yes 承認を得て実行)
kubetective ベンチマーク [スイート]
シナリオスイートゲートを実行します。あらゆる回帰で 1 を終了します
kubetective 評価 [スイート]
マークダウン評価レポート（カテゴリごとの精度、キャリブレーション、FP チェック）
kubetectiveserve --listen :8080
REST APIサーバー
キュベテクティブMCP
標準入出力上の MCP サーバー (読み取り専用ツール)
キューブ探偵バージョン
エンジンのバージョンを出力します
すべてのフラグに対して kubetective <command> --help を実行します。
KubeTective は、kubectl と同じ方法で kubeconfig を読み取ります ( --kubeconfig 、 --context 、
またはデフォルトの読み込みルール)。
状態ディレクトリ内の kubetective.yaml ( ~/.kubetective/kubetective.yaml ) セット
調査のデフォルト。例:
# ~/.kubetective/kubetective.yaml
コンテキスト : prod-eu
名前空間 : 支払い
以来: 30m
kubeconfig : ~/.kube/config
prometheus_url : http://localhost:9090
loki_url : http://localhost:3100
git_repo : ~/code/payments-manifests
cluster_id : prod-eu # オプション: 自動派生クラスター ID をオーバーライドします
# webhook_url: https://ops.example.com/kubetective-hook # オプトイン完了通知
# webhook_sec

[切り捨てられた]

## Original Extract

Deterministic root causes for every Kubernetes incident you don't want to debug twice - GlediLami/kubetective

GitHub - GlediLami/kubetective: Deterministic root causes for every Kubernetes incident you don't want to debug twice · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
GlediLami
/
kubetective
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
46 Commits 46 Commits .githooks .githooks .github .github Formula Formula cmd cmd deploy deploy hack hack internal internal pkg/ api pkg/ api reports reports scenarios scenarios site site .dockerignore .dockerignore .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md CODEOWNERS CODEOWNERS CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md demo.gif demo.gif go.mod go.mod go.sum go.sum kubectl-investigate.yaml kubectl-investigate.yaml View all files Repository files navigation
KubeTective: Kubernetes Incident Investigation Engine
KubeTective investigates Kubernetes incidents the way a good SRE would. It collects
facts about the target, builds a timeline and an evidence graph, generates ranked
hypotheses with explainable, deterministic scores, and shows you exactly why it
thinks what it thinks. Every point of confidence is backed by a visible list of
evidence.
It runs as a CLI, a kubectl plugin, a REST server, or an MCP server. Every
investigation is recorded so it can be replayed, audited, and used as a benchmark.
Docs: gledilami.github.io/kubetective · Changelog · Contributing
The demo above is a replay of a recorded incident
required.
kubectl investigate <target> --since=30m turns an "everything looks fine?"
moment into a ranked verdict with evidence (the GIF above is the full
output; this is the 5-second version):
$ kubectl investigate deployment/checkout --since=30m
╭──────────────────────────────────────────────────╮
│ INCIDENT: deployment/prod/checkout │
│ Status: OOMKILLED │
│ Severity: HIGH │
│ Confidence: 100 │
╰──────────────────────────────────────────────────╯
ROOT CAUSE
Memory exhaustion: container terminated with OOMKilled 19 time(s) (memory limit 1Gi) - 19 restart(s)
RECOMMENDATION
roll back deployment/prod/checkout to the last known-good revision [MEDIUM]
That output is the recorded incident scenarios/oom-after-deploy replayed
through the engine. 11 analyzers cover OOM kills, crash loops, image pull
failures, scheduling failures, node pressure, probe failures, PVC issues,
service selector mismatches, HPA at max, DNS failures, and configuration
regressions (Git/GitOps, down to the commit).
Try it in 5 minutes, no broken cluster needed
git clone https://github.com/GlediLami/kubetective.git && cd kubetective
make build
bin/kubetective replay scenarios/oom-after-deploy/record.jsonl # a real incident, replayed
bin/kubetective benchmark # 16 scenarios vs. ground truth
bin/kubetective evaluate # markdown accuracy report
Every recorded incident in scenarios/ is a self-contained, deterministic demo
and doubles as the regression gate for every analyzer.
KubeTective vs. the status quo
kubectl + grep + docs
LLM chat
KubeTective
Collects and correlates the facts
✗ you dig
✗ you paste
✓ automatic (k8s, Prometheus, Loki, Git/GitOps)
Ranked hypotheses with confidence
✗
~ vibes
✓ calibrated against ground truth
Every verdict shows its evidence
✗
rarely
✓ line-by-line score breakdown
Verifiable with a recorded replay
✗
✗
✓ JSONL records, re-run deterministically
Change detection (what changed right before)
✗
✗
✓ git commits, GitOps drift, pod changes
Regression-tested on a scenario suite
✗
✗
✓ kubetective benchmark in CI
Safe remediation with approval
✗
✗
✓ read-only preview → explicit --yes
Safe remediation with approval
✗
✗
✓ read-only preview → explicit --yes
Works offline, no API key, no telemetry
✓
✗
✓
An honest note on AI triage tools (k8sgpt etc.)
There is a popular class of "AI kubectl" tools that scan the cluster and
explain failures with an LLM. They are genuinely useful for triage — but
the core difference is what kind of artifact they produce:
Where they shine and we don't (yet): cluster-wide scanning, a bigger analyzer
catalog, LLM backend breadth, and an operator for continuous monitoring —
all on our roadmap, all reproducible from our side.
KubeTective's verdicts come from a rule-based engine, not an LLM. Each hypothesis
score is a sum of weighted evidence terms you can read line by line. Confidence is
calibrated against a scenario benchmark (16 recorded incidents with ground truth),
and the calibration is validated leave-one-out before it is adopted. The optional
LLM layer only explains the engine's verdict in plain language. It can never change
scores, invent causes, or propose actions. Deterministic output means the same
incident produces the same verdict in CI, in a demo, and in a replay an LLM chat
cannot be a regression test.
Evidence-based investigation. 11 analyzers cover the most common failure
modes: OOM kills, crash loops, image pull failures, scheduling and
unschedulable pods, node pressure, liveness and readiness probe failures, PVC
issues, service selector mismatches, HPA at max, DNS failures, and configuration
regressions (Git and GitOps).
Collectors. Kubernetes (pods, deployments, events, PVCs, services, HPAs,
coreDNS), Prometheus (memory-metric corroboration), Git (commits touching your
manifests), and GitOps (Flux Kustomization / HelmRelease , ArgoCD Application ).
Explainable scoring. Every hypothesis ships its evidence breakdown, and
scores are calibrated against the scenario suite with leave-one-out validation.
Timeline, evidence graph, and change detection. What happened, in what order,
what owns what, and what changed right before the incident.
Record and replay. Every investigation is appended to
~/.kubetective/incidents/ as JSONL. Replay it, diff engine versions, audit it.
Scenario benchmark and evaluation report. kubetective benchmark is the
regression gate for every analyzer. kubetective evaluate renders a markdown
evaluation report (per-scenario, per-category accuracy, calibration,
false-positive check) suitable for CI.
Safe remediation. Actions (rollback, restart) are previewed read-only and
applied only after explicit human approval. Every apply writes an audit record.
REST server and MCP server. The same pipeline behind an HTTP API and an MCP
stdio server with read-only tools.
Optional LLM explainer. Digest-only, redacted, constrained. Works with any
OpenAI-compatible endpoint (OpenAI, Ollama, vLLM, llama.cpp).
kubectl plugin. Drop the binary named kubectl-investigate on your PATH
and run kubectl investigate <resource> .
Add a scenario. Record a new incident with kubetective record --from-live ,
add ground truth, and it becomes both a demo and a benchmark case.
Harden an analyzer. Pick one of the 11 analyzers, find a false positive
against the suite, fix the scoring, and let kubetective benchmark prove it.
Improve the output. The renderer is plain-ANSI box drawing; new output
formats ( --format=json exists; a sarif or slack renderer is open).
Wire a new evidence source (e.g. Datadog, Grafana Cloud) behind the existing
collector interface.
See CONTRIBUTING.md for the design rules.
brew install gledilami/kubetective/kubetective
This installs both binaries: kubetective and kubectl-investigate (the kubectl
plugin). Homebrew finds the formula in the homebrew-kubetective tap automatically,
no extra steps.
go install github.com/GlediLami/kubetective/cmd/kubetective@latest
From source
Requires Go 1.26 or newer and a Kubernetes cluster (for live investigations; the
benchmark runs without one).
git clone https://github.com/GlediLami/kubetective.git
cd kubetective
make build # bin/kubetective + bin/kubectl-investigate
make install # kubetective -> ~/.local/bin/kubetective
make install-plugin # kubectl-investigate -> ~/.local/bin (kubectl plugin)
Point PREFIX=/usr/local at make install to install system-wide.
kubectl auto-discovers plugins named kubectl-* on PATH :
make install-plugin
kubectl investigate pod/checkout-7f84c9 # same as: kubetective investigate ...
Quick start
Investigate a crash-looping deployment (uses your current kubeconfig context):
kubectl investigate deployment/checkout --since=30m
# or without the plugin:
kubetective investigate deployment/checkout --since=30m
Target forms: pod/<name> , deployment/<name> , namespace/<name> (or a bare
name, which defaults to a pod), optionally scoped with --namespace and a window
( --since=2h , default 30m).
kubetective investigate deployment/checkout \
--since=30m \
--prometheus-url=http://localhost:9090 \ # metric corroboration
--loki-url=http://localhost:3100 \ # log evidence (Loki)
--git-repo= ~ /code/checkout-manifests # commits touching your manifests
Command reference
Command
Description
kubetective investigate <resource>
Run an investigation (flags: --since , --namespace , --no-logs , --format=json , --prometheus-url , --loki-url , --git-repo , --llm* )
kubetective replay <incident-id>
Re-run a recorded investigation through the current engine (deterministic)
kubetective incidents
List recorded incident ids, newest first
kubetective incidents similar <id> [--cluster <id>]
Find similar past incidents (incident memory, Jaccard overlap; --cluster scopes the lookup to one cluster)
kubetective alert <pagerduty|grafana|slack>
Investigate from a webhook alert payload (stdin or --file ; zero API keys)
kubetective doctor
Health check (version, config file, calibration, cluster connectivity, incident store, Prometheus/Loki reachability); exits non-zero on any failure
kubetective action <incident-id>
Preview remediation actions (read-only; --apply <id> --yes to execute with approval)
kubetective benchmark [suite]
Run the scenario suite gate. Exits 1 on any regression
kubetective evaluate [suite]
Markdown evaluation report (per-category accuracy, calibration, FP check)
kubetective serve --listen :8080
REST API server
kubetective mcp
MCP server over stdio (read-only tools)
kubetective version
Print the engine version
Run kubetective <command> --help for every flag.
KubeTective reads kubeconfig the same way kubectl does ( --kubeconfig , --context ,
or the default loading rules).
kubetective.yaml in the state directory ( ~/.kubetective/kubetective.yaml ) sets
defaults for an investigation. Example:
# ~/.kubetective/kubetective.yaml
context : prod-eu
namespace : payments
since : 30m
kubeconfig : ~/.kube/config
prometheus_url : http://localhost:9090
loki_url : http://localhost:3100
git_repo : ~/code/payments-manifests
cluster_id : prod-eu # optional: override the auto-derived cluster identity
# webhook_url: https://ops.example.com/kubetective-hook # opt-in completion notification
# webhook_sec

[truncated]
