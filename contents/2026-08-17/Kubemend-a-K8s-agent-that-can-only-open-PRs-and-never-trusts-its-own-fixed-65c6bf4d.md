---
source: "https://github.com/m-stepkowski/kubemend"
hn_url: "https://news.ycombinator.com/item?id=49328606"
title: "Kubemend – a K8s agent that can only open PRs, and never trusts its own \"fixed\""
article_title: "GitHub - m-stepkowski/kubemend: An LLM agent that diagnoses Kubernetes incidents from Observability tools and remediates by opening GitOps pull requests — never touching the cluster directly. Hand-built harness (no agent framework), independent verification gate, and a fault-injection eval lab with\n[truncated]"
image: "https://opengraph.githubassets.com/73963067d6f94e6b69d7dbeb0aa9e9a5e248d674b70b519e920cf61f1a251beb/m-stepkowski/kubemend"
author: "m-stepkowski"
captured_at: "2026-08-17T10:22:16Z"
capture_tool: "hn-digest"
hn_id: 49328606
score: 2
comments: 0
posted_at: "2026-08-17T10:03:53Z"
tags:
  - hacker-news
  - translated
---

# Kubemend – a K8s agent that can only open PRs, and never trusts its own "fixed"

- HN: [49328606](https://news.ycombinator.com/item?id=49328606)
- Source: [github.com](https://github.com/m-stepkowski/kubemend)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T10:03:53Z

## Translation

タイトル: Kubemend – PR のみを開くことができ、自身の「修正」を決して信頼しない K8s エージェント
記事のタイトル: GitHub - m-stepkowski/kubemend: オブザーバビリティ ツールから Kubernetes インシデントを診断し、GitOps プル リクエストを開くことで修復する LLM エージェントです。クラスターには直接触れません。手作りのハーネス (エージェント フレームワークなし)、独立した検証ゲート、およびフォールト挿入評価ラボ
[切り捨てられた]
説明: オブザーバビリティ ツールから Kubernetes インシデントを診断し、GitOps プル リクエストを開くことで修復する LLM エージェント。クラスターには直接触れません。手作りのハーネス (エージェント フレームワークなし)、独立した検証ゲート、再現可能な合格率ベンチマークを備えたフォールト挿入評価ラボ
[切り捨てられた]

記事本文:
GitHub - m-stepkowski/kubemend: オブザーバビリティ ツールから Kubernetes インシデントを診断し、GitOps プル リクエストを開くことで修復する LLM エージェントです。クラスターには直接触れません。手作りのハーネス (エージェント フレームワークなし)、独立した検証ゲート、再現可能な合格率ベンチマークを備えたフォールト挿入評価ラボ。 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
m-ステップコウスキー
/
クベメンド
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 開く

その他のアクション メニュー フォルダーとファイル
108 コミット 108 コミット .claude .claude .github/ workflows .github/ workflows charts/ kubemend charts/ kubemend config config docs docs evals evals kubemend kubemend lab lab ポリシー ポリシー テスト テスト .dockerignore .dockerignore .gitignore .gitignore .python-version .python-version ARCHITECTURE.md ARCHITECTURE.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile IMPLEMENTATION_PLAN.md IMPLEMENTATION_PLAN.md ライセンス ライセンス README.md README.md Taskfile.yaml Taskfile.yaml kubemend.yaml kubemend.yaml pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
プル リクエストのみを開くことができる GitOps ネイティブの Kubernetes 修復エージェント。
人間に何か承認を求める前に、Prometheus メトリクスと Loki ログからインシデントを診断し、修正を提案し、その修正自体を検証します (ヘルム レンダリング → Kyverno ポリシー チェック → ライブ差分 → スコープ チェック → ライブ クォータ ヘッドルーム)。 kubectl apply は決して実行されません。書き込み可能なクラスター資格情報がありません。
本番環境に対応していません: マルチリポジトリの GitOps がなく、サンドボックス ツールの実行もまだありません。範囲内および範囲外については、docs/threat-model.md を参照してください。
ほとんどの「AI SRE エージェント」デモは印象的ですが検証不可能です。モデルは何かを修正したと主張しており、その言葉をそのまま受け入れます。 kubemend は逆の方法で構築されています。つまり、モデルの成功の主張は決して信頼されません。すべての実行は、提案された修正が適切にレンダリングされ、ポリシーを満たし、実際のスコープ付きの差分が生成され、宣言されたインシデント以外には何も触れていないことを独立した検証パイプラインが判断した後にのみ終了します。エージェントの唯一のアクチュエーターは Git ブランチとドラフト PR です。マージするのは依然として人間です。
これは、LangChain/CrewAI/AutoGen のラッパーではなく、最初から作成したエージェント ハーネスでもあります。ループ

、コンテキスト管理、ツール レジストリ、および検証ゲートは手書きで文書化されています。これは、フレームワークを接着することではなく、これらのトレードオフを理解することがプロジェクトのポイントであるためです。
タスク ──▶ ループ ──▶ ツール呼び出し ──▶ Prometheus / Loki / K8s (読み取り専用)
│
└── モデルは「完了」と主張 ──▶ 独立した検証ゲート
Helm テンプレート → kyverno 適用
→ argocd/kubectl diff → スコープチェック
→ ライブ クォータ ヘッドルーム
│
パス ──▶ GitOps リポジトリに対する PR ドラフト
失敗 ──▶ 構造化された失敗がループにフィードバックされる
可観測性: Prometheus/Mimir に対する PromQL、Loki に対する LogQL。プロバイダー インターフェイスの背後で交換可能 (Dynatrace/CloudWatch は将来のドロップインになります)。
クラスター アクセス: 読み取り専用の ServiceAccount、許可リストに登録されたリソースの種類、Secret 値はフェッチされません。
修復: エージェントは Helm 値*.yaml のみを編集し、テンプレートを直接編集することはありません。そのため、差分は小さくレビュー可能な状態に保たれます。
検証: モデルの言葉を鵜呑みにせず、終了時にハーネスによって独立して再実行します。
すべてが評価されます。密閉された種類ベースのフォールト挿入ラボは、プロパティ ベースのチェッカーを使用して実際のインシデント (不正なイメージ タグ、OOMKill、構成キーの欠落、壊れたプローブなど) を再現し、シナリオごとに N 回実行して、厳選したデモではなく、合格率/コスト/反復テーブルを作成します。さらに 3 つのシナリオは、意図的に敵対的です。値のみの解決策を持たない修正、宣言された範囲外に本当の原因があるインシデント、エージェント自身のログ証拠に埋め込まれたプロンプト インジェクションの試行です。それぞれのシナリオは、ハンドオフまたはスコープ クリーンな PR を期待しており、もっともらしく見える間違った答えは決してありません。
完全な設計、不変条件、およびその根拠を含むすべての数値デフォルト: ARCHITECTURE.md 。
mainとcheapはそれぞれ独立して設定されているため、プロバイダーを混在させる
ネクタイを越えて

rs (例: メインの場合は Bedrock の Claude、安価な場合は DeepSeek)
特別なケースではなく、通常の構成:
モデル:
メイン:
提供者：岩盤
名前: us.anthropic.claude-sonnet-5-v1:0
aws_region : us-east-1
安い：
プロバイダー：オープンアイ
名前：ディープシーク-v4-フラッシュ
ベース URL : https://api.deepseek.com
その他の例については、 kubemend.yaml 自身のコメントを参照してください。
コストガードレールの価格設定については config/pricing.yaml —
非人類的エントリーには、公開価格をソースとするプレースホルダーがあります
ページ、請求書と照合されていない。信頼する前に確認してください
コミットされたベースライン。
タグ付けされたすべてのリリースは、PyPI と ghcr.io の両方に公開されます。
pip インストール kubemend
docker pull ghcr.io/m-stepkowski/kubemend:latest
docker run --rm ghcr.io/m-stepkowski/kubemend:latest --help
いずれの方法でも、モデルの認証情報 (デフォルトでは ANTHROPIC_API_KEY —) が必要になります。
上記の「モデルプロバイダー」を参照してください)、および kubemend.yaml は、
クラスターの Prometheus/Loki、kubeconfig、および GitOps リポジトリ — コミットされたリポジトリを参照してください
各フィールドに対する kubemend.yaml 独自のコメント。走るには
ラップトップからではなくクラスター内で展開する場合は、以下の「クラスター内での展開」を参照してください。
Docker (または Rancher デスクトップが必要です)
何でも使用できます)、 uv 、および
go-task 、および ANTHROPIC_API_KEY 。
エンドツーエンドで動作することを確認する最速の方法 — ラボを立ち上げ、実際の
障害が発生した場合、それに対してエージェントを実行し、結果の提案を出力します。これは次のとおりです。
git clone https://github.com/m-stepkowski/kubemend.git && cd kubemend
UV同期
エクスポート ANTHROPIC_API_KEY=...
task lab:up # 種類クラスター: gitea、Argo CD、kube-prometheus-stack、Loki、Kyverno
タスクデモ # フォールトを挿入し、kubemend を実行し、結果の提案を表示します (~90 秒)
タスクデモはデフォルトで安価なモデルで実行されます。 pass -- --model main を使用する
以下のヘッドラインスイープが実行されたモデル:
タスクデモ -- --モデルメイン
v の代わりに手動で運転するには

デモスクリプトは次のとおりです。
task lab:forward # port-forward Prometheus/Loki/gitea/Argo をローカルでブロックし、別のターミナルで実行します
kubemend run --task " 名前空間 shop の shop-api ポッドが 10 分前からクラッシュ ループしています " \
--namespace shop --app shop-api
これにより、ブランチ (そして gitops.backend: gitea を使用して、実際のドラフト PR が書き込まれます)
ラボの gitea インスタンス）と、traces/ の下の完全な JSONL トレース。参照
信頼境界については docs/threat-model.md 、
まだ範囲外のもの (単一リポジトリ、値のみの編集、永続化なし)
実行後のメモリ)。
逸話ではなく、再現可能な合格率ベンチマーク - すべてのシナリオが N 回実行され、コストと反復回数がレポートされます。
task evals -- --scenarios all -n 5 --model main
v0.1 ベースライン ( claude-sonnet-5 、シナリオごとに n=5、合計 11.08 ドル —
evals/reports/v0.1-baseline/ ):
全体で 29/30 (97%) 合格。 1 つの失敗は、モデルの問題ではなく、真のモデルの闘争です。
ハーネスのバグ: bad-probe-path の失敗した実行後に Budget_exhausted がヒットしました
収束せずにproposal_git_change / validate_changeのサイクルを繰り返しました。
敵対的シナリオ、M6 ベースライン ( claude-sonnet-5 、シナリオごとに n=3、
合計 4.01 ドル、このスイープの予算の上限は 5 ドル —
evals/reports/m6-baseline/ ):
ここでは n=10 ではなく n=3 —scope-trap の実行ごとの実際のコスト (15 回の反復、
$0.71) により、このベースラインの予算では実現不可能な大規模なスイープが行われました。
切り上げられず、正直な n=3 サンプルとして報告されます。一つの失敗
( fix-needs-template-change ) は実際の特定のモデル ギャップです。正しくは、
ハードコードされたプローブスキームが根本原因であると診断したが、
「値のみの修正は存在しない」とコミットするのではなく、ハンドオフします。参照
ログインジェクションについては docs/threat-model.md §9
シナリオの完全なトレースの抜粋。
日常回帰に使用される安価なモデル ( claude-haiku-4-5 ) の数値
開発中のスイープは、

安くて安い — を参照
evals/reports/latest/ 。
ラップトップから kubemend を実行するには、完全な読み取り専用を保持する kubeconfig が必要です。
RBAC kubemend を使用します。 Helm チャートは、それを絞り込むために存在します。
一度インストールすれば、オンコール エンジニアは作成許可だけが必要です。
リーダー自身の権限ではなく、1 つの名前空間内のジョブ。
helm install kubemend charts/kubemend -n kubemend-system --create-namespace
これにより、リーダー ServiceAccount と RBAC (名前空間スコープのロール) がインストールされます。
デフォルト。 --set rbac.clusterScoped=true (ClusterRole の場合) と生成
何もありません — job.enabled のデフォルトは false です。実行をトリガーするには:
Helm テンプレート kubemend charts/kubemend \
--namespace kubemend-system \
--set job.enabled=true \
--set job.namespace=shop \
--set job.app=shop-api \
--set job.task= " shop-api ポッドがクラッシュ ループしている " \
-s テンプレート/job.yaml \
| kubectl create -f -
ジョブは、スコープが狭い独自のクラスター内 ServiceAccount を使用して実行されます。
( kubernetes.in_cluster: true 、kubeconfig ファイルは関係しません) 同じ経由
ghcr.io/m-stepkowski/kubemend イメージは各リリースで公開されます。参照
charts/kubemend/README.md での配線について
GitOps リポジトリのチェックアウトと完全な値の参照。
M8b 以降では、アラートによってトリガーされる自動化も利用できます。 --setoperator.enabled=true は、小さな Webhook レシーバー (stdlib) をデプロイします。
http.server 、フレームワークなし）同じ種類のジョブを独自に作成します
Alertmanager が起動したとき、必要なベアラー トークンとスコープごとにゲートされます。
クールダウン。これは、リーダーとリーダーの両方からの明確で狭い RBAC ID です。
手動トリガー パスであり、ジョブが開始された後の動作は変わりません
— すべての実行は引き続き同じ信頼できないモデルのループを通過します。
確認ゲート。 charts/kubemend/README.md を参照してください。
「アラートでトリガーされるオペレーター」セクションを有効にして、
実際に実行する前に docs/threat-model.md §11 を参照してください。
クラスター

r.
kubemend/ハーネスコア、ツール、gitopsモジュール、検証ゲート
プロンプト/バージョン管理されたシステム/コンパクション/ハンドオフ プロンプト
ポリシー/ Kyverno パック (入場者とバリデーターによって共有)
ラボ/種類のブートストラップ、ラボ GitOps リポジトリ、フォールト挿入シナリオ
evals/スイープランナー + コミットされたベースラインレポート
テスト/ユニット (FakeLLM、ネットワークなし) + 統合 (ラボに対して)
docs/knowledge/ 設計契約 — core/、tools/、またはシナリオを変更する前に読んでください。
各モジュールの完全なツリーと理論的根拠: ARCHITECTURE.md §9 。
まだ外部からの貢献は受け付けていません。IMPLEMENTATION_PLAN.md のマイルストーンに向けて作業中です。それまでの間、問題やデザインについての議論を歓迎します。
オブザーバビリティ ツールから Kubernetes インシデントを診断し、GitOps プル リクエストを開くことで修復する LLM エージェント。クラスターに直接触れることはありません。手作りのハーネス (エージェント フレームワークなし)、独立した検証ゲート、再現可能な合格率ベンチマークを備えたフォールト挿入評価ラボ。
Readme Apache-2.0 ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An LLM agent that diagnoses Kubernetes incidents from Observability tools and remediates by opening GitOps pull requests — never touching the cluster directly. Hand-built harness (no agent framework), independent verification gate, and a fault-injection eval lab with reproducible pass-rate benchmark
[truncated]

GitHub - m-stepkowski/kubemend: An LLM agent that diagnoses Kubernetes incidents from Observability tools and remediates by opening GitOps pull requests — never touching the cluster directly. Hand-built harness (no agent framework), independent verification gate, and a fault-injection eval lab with reproducible pass-rate benchmarks. · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
m-stepkowski
/
kubemend
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
108 Commits 108 Commits .claude .claude .github/ workflows .github/ workflows charts/ kubemend charts/ kubemend config config docs docs evals evals kubemend kubemend lab lab policies policies tests tests .dockerignore .dockerignore .gitignore .gitignore .python-version .python-version ARCHITECTURE.md ARCHITECTURE.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile IMPLEMENTATION_PLAN.md IMPLEMENTATION_PLAN.md LICENSE LICENSE README.md README.md Taskfile.yaml Taskfile.yaml kubemend.yaml kubemend.yaml pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
A GitOps-native Kubernetes remediation agent that can only open pull requests.
It diagnoses incidents from Prometheus metrics and Loki logs, proposes a fix, and verifies that fix itself — helm render → Kyverno policy check → live diff → scope check → live quota headroom — before it ever asks a human to approve anything. It never runs kubectl apply . It has no cluster credentials that can write.
Not production-ready: no multi-repo GitOps, no sandboxed tool execution yet. See docs/threat-model.md for what's in and out of scope.
Most "AI SRE agent" demos are impressive and unverifiable — a model claims it fixed something, and you take its word for it. kubemend is built the other way around: the model's claim of success is never trusted. Every run terminates only after an independent validation pipeline says the proposed fix renders cleanly, satisfies policy, produces a real and scoped diff, and touches nothing outside the declared incident. The agent's only actuator is a Git branch and a draft PR — a human still merges.
It's also a from-scratch agent harness, not a wrapper around LangChain/CrewAI/AutoGen. The loop, context management, tool registry, and verification gate are hand-written and documented, because understanding those trade-offs — not gluing a framework together — is the point of the project.
task ──▶ Loop ──▶ tool calls ──▶ Prometheus / Loki / K8s (read-only)
│
└── model claims "done" ──▶ independent verification gate
helm template → kyverno apply
→ argocd/kubectl diff → scope check
→ live quota headroom
│
pass ──▶ draft PR against the GitOps repo
fail ──▶ structured failure fed back into the loop
Observability: PromQL against Prometheus/Mimir, LogQL against Loki. Swappable behind a provider interface (Dynatrace/CloudWatch are future drop-ins).
Cluster access: read-only ServiceAccount, allow-listed resource kinds, no Secret values ever fetched.
Remediation: the agent edits Helm values*.yaml only — never templates directly — so diffs stay small and reviewable.
Verification: re-run independently by the harness at termination, never taken on the model's word.
Everything is evaluated: a hermetic kind -based fault-injection lab reproduces real incidents (bad image tags, OOMKills, missing config keys, broken probes...) with property-based checkers, run N times per scenario to produce pass-rate / cost / iteration tables — not cherry-picked demos. Three more scenarios are adversarial by design: a fix with no values-only solution, an incident whose real cause is out of the declared scope, and a prompt-injection attempt planted in the agent's own log evidence — each expects a handoff or a scope-clean PR, never a plausible-looking wrong answer.
Full design, invariants, and every numeric default with its rationale: ARCHITECTURE.md .
main and cheap are each configured independently, so mixing providers
across tiers (e.g. Claude on Bedrock for main , DeepSeek for cheap ) is a
normal configuration, not a special case:
model :
main :
provider : bedrock
name : us.anthropic.claude-sonnet-5-v1:0
aws_region : us-east-1
cheap :
provider : openai
name : deepseek-v4-flash
base_url : https://api.deepseek.com
See kubemend.yaml 's own comments for more examples, and
config/pricing.yaml for cost-guardrail pricing —
non-Anthropic entries there are placeholders sourced from public pricing
pages, not verified against an invoice; check before trusting them for a
committed baseline.
Every tagged release publishes to both PyPI and ghcr.io:
pip install kubemend
docker pull ghcr.io/m-stepkowski/kubemend:latest
docker run --rm ghcr.io/m-stepkowski/kubemend:latest --help
Either way you'll need model credentials ( ANTHROPIC_API_KEY by default —
see "Model providers" above) and a kubemend.yaml pointing at your
cluster's Prometheus/Loki, kubeconfig, and GitOps repo — see the committed
kubemend.yaml 's own comments for every field. To run
in-cluster instead of from a laptop, see "Deploy in-cluster" below.
Requires Docker (or Rancher Desktop —
anything kind can use), uv , and
go-task , plus an ANTHROPIC_API_KEY .
The fastest way to see it work end to end — bring up the lab, inject a real
fault, run the agent against it, and print the resulting proposal — is:
git clone https://github.com/m-stepkowski/kubemend.git && cd kubemend
uv sync
export ANTHROPIC_API_KEY=...
task lab:up # kind cluster: gitea, Argo CD, kube-prometheus-stack, Loki, Kyverno
task demo # inject a fault, run kubemend, show the resulting proposal (~90s)
task demo runs on the cheap model by default; pass -- --model main to use
the model the headline sweep below was run on:
task demo -- --model main
To drive it by hand instead of via the demo script:
task lab:forward # port-forward Prometheus/Loki/gitea/Argo locally, blocks — run in another terminal
kubemend run --task " shop-api pods in namespace shop are crash-looping since 10 minutes ago " \
--namespace shop --app shop-api
This writes a branch (and, with gitops.backend: gitea , a real draft PR in
the lab's gitea instance) plus a full JSONL trace under traces/ . See
docs/threat-model.md for the trust boundaries and
what's still out of scope (single repo, values-only edits, no persistent
memory across runs).
Reproducible pass-rate benchmarks, not anecdotes — every scenario is run N times and reported with cost and iteration counts:
task evals -- --scenarios all -n 5 --model main
v0.1 baseline ( claude-sonnet-5 , n=5 per scenario, $11.08 total —
evals/reports/v0.1-baseline/ ):
29/30 (97%) pass overall. The one failure is a genuine model struggle, not a
harness bug: bad-probe-path 's failing run hit budget_exhausted after
repeated propose_git_change / validate_change cycling without converging.
Adversarial scenarios, M6 baseline ( claude-sonnet-5 , n=3 per scenario,
$4.01 total, capped at a $5 budget for this sweep —
evals/reports/m6-baseline/ ):
n=3 here, not n=10 — scope-trap 's real per-run cost (15 iterations,
$0.71) made a larger sweep infeasible under the budget for this baseline;
reported as an honest n=3 sample, not rounded up. The one failure
( fix-needs-template-change ) is a real, specific model gap: it correctly
diagnosed a hardcoded probe scheme as the root cause but hedged on the
handoff instead of committing to "no values-only fix exists." See
docs/threat-model.md §9 for the log-injection
scenario's full trace excerpt.
Cheap model ( claude-haiku-4-5 ) numbers, used for day-to-day regression
sweeps during development, are lower and cheaper — see
evals/reports/latest/ .
A kubemend run from a laptop needs a kubeconfig holding the full read-only
RBAC kubemend uses. The Helm chart exists to narrow that:
install it once and an on-call engineer only needs permission to create a
Job in one namespace, not the reader's own permissions.
helm install kubemend charts/kubemend -n kubemend-system --create-namespace
This installs the reader ServiceAccount and RBAC (namespace-scoped Role by
default; --set rbac.clusterScoped=true for a ClusterRole ) and spawns
nothing — job.enabled defaults to false . To trigger a run:
helm template kubemend charts/kubemend \
--namespace kubemend-system \
--set job.enabled=true \
--set job.namespace=shop \
--set job.app=shop-api \
--set job.task= " shop-api pods are crash-looping " \
-s templates/job.yaml \
| kubectl create -f -
The Job runs with its own tightly-scoped in-cluster ServiceAccount
( kubernetes.in_cluster: true , no kubeconfig file involved) via the same
ghcr.io/m-stepkowski/kubemend image published on each release. See
charts/kubemend/README.md for wiring in a
GitOps repo checkout and the full values reference.
Alert-triggered automation is also available as of M8b: --set operator.enabled=true deploys a small webhook receiver (stdlib
http.server , no framework) that creates the same kind of Job on its own
when Alertmanager fires, gated by a required bearer token and a per-scope
cooldown. It is a distinct, narrower-RBAC identity from both the reader and
the manual-trigger path, and does not change what happens once a Job starts
— every run still goes through the same untrusted-model loop and
verification gate. See charts/kubemend/README.md 's
"Alert-triggered operator" section to enable it, and
docs/threat-model.md §11 before doing so in a real
cluster.
kubemend/ harness core, tools, gitops module, verification gate
prompts/ versioned system/compaction/handoff prompts
policies/ Kyverno pack (shared by admission and the validator)
lab/ kind bootstrap, lab GitOps repo, fault-injection scenarios
evals/ sweep runner + committed baseline reports
tests/ unit (FakeLLM, no network) + integration (against the lab)
docs/knowledge/ design contracts — read before modifying core/, tools/, or scenarios
Full tree and rationale for each module: ARCHITECTURE.md §9 .
Not yet open for external contributions — still working through the milestones in IMPLEMENTATION_PLAN.md . Issues and design discussion welcome in the meantime.
An LLM agent that diagnoses Kubernetes incidents from Observability tools and remediates by opening GitOps pull requests — never touching the cluster directly. Hand-built harness (no agent framework), independent verification gate, and a fault-injection eval lab with reproducible pass-rate benchmarks.
Readme Apache-2.0 license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
