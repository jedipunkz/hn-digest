---
source: "https://ctrlnode.ai/news/fable-claude-model-audit-experiment/"
hn_url: "https://news.ycombinator.com/item?id=48768491"
title: "We Ran a Complex Task – A LangChain Repo Analysis with Claude Fable Models"
article_title: "We Ran a Complex Task — A LangChain Repo Analysis with Five Claude Models | CTRL NODE"
author: "ctrlnode-ai"
captured_at: "2026-07-02T23:09:20Z"
capture_tool: "hn-digest"
hn_id: 48768491
score: 1
comments: 0
posted_at: "2026-07-02T23:01:41Z"
tags:
  - hacker-news
  - translated
---

# We Ran a Complex Task – A LangChain Repo Analysis with Claude Fable Models

- HN: [48768491](https://news.ycombinator.com/item?id=48768491)
- Source: [ctrlnode.ai](https://ctrlnode.ai/news/fable-claude-model-audit-experiment/)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T23:01:41Z

## Translation

タイトル: 複雑なタスクを実行しました – クロード・フェイブル・モデルを使用した LangChain リポジトリ分析
記事のタイトル: 複雑なタスクを実行しました — 5 つのクロード モデルを使用した LangChain リポジトリ分析 | CTRLノード
説明: 制御された実験: 1 つのプリンシパル エンジニア監査プロンプト、5 つのクロード層、1 つの LangChain クローン。 CTRL NODE でどのように実行したか、Fable が何を提供したか、そしてなぜ単一のモデルが勝てないのか。

記事本文:
複雑なタスクを実行しました — 5 つのクロード モデルを使用した LangChain リポジトリ分析 | CTRLノード
CTRL NODE 製品ワークフロー カンバン ルーチン ブリッジ プロバイダー リソース 変更履歴 ドキュメント クイック スタート GitHub ニュース
サインイン
無料で始める
製品のワークフロー カンバン ルーチン ブリッジ プロバイダー リソース 変更履歴 ドキュメント クイック スタート GitHub ニュース
サインイン
無料で始める
エンジニアリング · 2026 年 7 月 2 日 · 11 分で読みました 複雑なタスクを実行しました — 5 つのクロード モデルを使用した LangChain リポジトリ分析
Anthropic は Claude Fable を出荷しました。私たちは実践的な質問に対する本当の答えを求めていました。
Opus、Fable、Sonnet、Haiku で同じ複雑なエンジニアリング タスクを実行すると、実際には何が得られるでしょうか?
ベンチマークスコアではありません。雰囲気チェックではありません。証拠、重大度ラベル、および実行計画を含む、運用環境のオープンソース モノリポジトリの完全なプリンシパル エンジニアによる監査。
CTRL NODE 内でその実験を実行しました: 1 つのプロンプト、5 つのエージェント、5 つのモデル、1 つのクローン リポジトリ。
1. 目標: 1 つの難しいタスク、5 つのモデル
私たちはすべてのモデルに同じ 4 フェーズの監査プロンプトと同じターゲット、つまり LangChain Python モノリポジトリ (おもちゃのリポジトリではなく、大規模で成熟したライブラリ エコシステム) を与えました。
リポジトリ マップ — 最初に探索し、次に判断します
監査レポート — アーキテクチャ、セキュリティ、テスト、パフォーマンス、deps、DX、ドキュメント (ファイル:行の引用付き)
改善戦略 — テーマ、トレードオフ、測定可能な「完了」基準
タスク計画 - マイルストーン M0 ～ M3、クイック ウィン、各項目の労力/リスク/デプス
すべての発見は証拠に基づいていなければなりません。推測は明示的に禁止されています。
これは、何千ものファイル、実際の CI 構成、セキュリティに配慮した逆シリアル化パス、ホット コード パス上の神クラスのモジュールなど、非常に重いタスクです。これは通常、複数の上級エンジニアにまたがる作業チームのようなものです。
Fable は st として位置付けられます。

長く構造化された作業のための適切な推論モデル。以下のものと一緒に含めました。
この仮説は「寓話がすべてに勝つ」というものではありませんでした。それは、各層が異なることを認識しており、結果を出荷可能なバックログに変えるのに Fable が最も優れている可能性があるということです。
完全なプロンプトはカタログに langchain-prompt.md として存在します。主要な指示 (省略):
あなたは世界クラスのプリンシパル エンジニア レベルのソフトウェア エンジニアであり、技術監査の専門家です。
このコード リポジトリの詳細な分析を実行し、正直な監査レポートを提供します。
そして、優先順位が付けられた実行可能な改善計画を提案します。
発見 → 監査 → 戦略 → タスク計画の 4 つのフェーズを順番に実行します。
すべての判決では、実際のファイル パスと行番号を引用する必要があります。推測しないでください。
実行ごとに要求される成果物:
Audit-report-<model>.md — 完全なマークダウン レポート
Audit-report-<model>.html — インタラクティブなダークテーマ ダッシュボード (タブ: 概要、マップ、監査、戦略、タスク)
プロンプトの概要:resumen-langchain-prompt.md 。
2. CTRL NODEでの設定方法
プロンプトを 5 つのブラウザー タブに貼り付けませんでした。私たちはチームが行う方法でそれを実行しました。つまり、実マシン上にブリッジし、クローンを指すプロジェクト作業ディレクトリを配置し、モデル層ごとに 1 つのエージェントを配置しました。
ブリッジ (ctrlnode) がインストールされ、ペアリングされました。「ブリッジのセットアップ」を参照してください。
~/.ctrlnode/.env に設定された Claude SDK API キー (プロバイダーは自動的にロードされます。PROVIDERS フラグは必要ありません):
ANTHROPIC_API_KEY=sk-ant-...
BASE_PATH=/ホーム/あなた/ワークスペース
LangChain は、Bridge ホストの BASE_PATH の下にクローン作成されました (CTRL NODE は git クローンを作成しません。作業ディレクトリは既存のフォルダーを指します)。
作業ディレクトリは、エージェントが WORK DIRECTORY タスク モードで完全なツリーを読み取ることができるものであり、スタッフ エンジニアが必要とする範囲と同じです。
チーム → + エージェントの追加 — 同じプロジェクトに 5 つのエージェントを作成しました。
モデルは MODEL コンボボックスで選択されます (同期

オンライン時に Bridge から d を入力するか、手動で入力します。 Fable は、Bridge モデル マニフェスト (v2026.2.4 以降) では claude-fable-5 として表示されます。
オプションのエージェント システム命令は最小限のままにしました。エージェントごとのペルソナの違いではなく、タスク プロンプトに仕様を伝える必要がありました。
各エージェントに対して、同じ手順を実行します。
タイトル : LangChain プリンシパル監査 — <モデル>
手順: langchain-prompt.md の内容全体を貼り付けます。
ASSIGN TO AGENT : 一致するエージェントチップを選択します
出力モード : WORK DIRECTORY (フル リポジトリ スコープ、オプションのフォーカス パスは空のまま)
新しいタスク → タスクがバックログに追加されます
RUN → Bridge にディスパッチ → エージェントは In progress に移動
Bridge は repositoryPaths とリポジトリ ディスパッチ コンテキストを使用してタスクを配信するため、Claude SDK はディスク上の LangChain ツリーに対して実行されます。出力 (audit-report-*.md / .html) はエー​​ジェントの作業ディレクトリから収集され、マーケティング カタログ フォルダーにコピーされました。
再現性のためのヒント: 実行ごとに同じコミット SHA を使用します。私たちのレポートでは、記載されている箇所で、LangChain マスター (2b47357) を参照しています。
Fable はリポジトリを A- と評価しました。これは Opus と同じ校正であり、Haiku が自己評価した A よりも誠実です。
複雑さの集中 - 5 つのファイルが 1,800 行を超えています。 runnables/base.py は 6,574 LOC です。すべての呼び出し/ストリーム パスでの爆発半径が大きい。
デフォルトでは安全ではない逆シリアル化 — langchain_core.load のデフォルトは allowed_objects='core' であり、信頼できないマニフェストでは安全ではないことが文書化されています。安全なオプションは存在しますが、オプトインです。
タイプ セーフティ エスケープ ハッチ — 208 タイプ: langchain-core のみのコメントを無視します。 disallow_any_generics=false は、パブリック API コントラクトを弱めます。
次のメジャー バージョンでは、逆シリアル化のデフォルトを安全な許可リスト ( 'messages' ) に切り替えます。
パークされた lint TODO ( BLE 、 ANN401 、 ERA ) を焼き消します - 施行インフラはすでに存在します。
変更されていない公開ファイルの背後にある最上位の神ファイルを分解する

çades (ゼロ API ブレーク)。
Fable の差別化要因は、セキュリティの見出しに対する注目度の高さではありませんでした。それはフェーズ 3 とフェーズ 4 でした。
4 つの戦略的テーマ (複雑さ、スイッチオフのガードレール、デフォルトで安全な信頼境界、ワークスペースの衛生状態)
明示的な非目標 (例: このサイクルではベンダー提供のMustache.py を書き換えないでください。代わりにプロパティ テストを追加します)
マイルストーン M0 ～ M3 とワークロード バッジ (S/M/L/XL)、リスク、依存関係、および許容基準
午後に出荷できる迅速な成果 (監査アーティファクトの .gitignore、callbacks/usage.py の飲み込まれた AttributeError の logger.debug、type:ignore count の CI ラチェット)
ほぼ独占的な寓話の発見:
独自のセキュリティ面を備えたベンダー製 704 行 Mustache エンジン (Mustache.py)
McCabe C90 の複雑さの lint が明示的に無効化 - ゴッドファイルの増大に対する自動バックプレッシャーなし
langchain_v1/agents/factory.py の狭いテスト範囲と複雑さ (56 のテスト ファイルと 1,891 行のファクトリー)
Fable では、他のモデルで見つかったいくつかの問題が表面化しませんでした。
SSRF パスでの TOCTOU / DNS 再バインド (Opus)
デフォルトでの ShellToolMiddleware ホスト実行 (Opus)
SSRF トランスポートは 2 つの呼び出しサイトのみで採用 + 保護されていないgraph_mermaid.py フェッチ (Sonnet 5)
CI _lint.yml のコメント化されたロックファイル チェック (Sonnet 4.6)
壊れた README モデルの例 / SECURITY.md が見つからない (Sonnet 4.6)
このギャップが重要です。Fable はマルチモデル パイプラインの代替品ではありません。
完全なレポート: Audit-report-fable.md · インタラクティブ ダッシュボード: Audit-report-fable.html
5. 5 つのモデルの比較
*Haiku の A は紙の上では自信を持って見えます。 Sonnet 4.6 とのクロスチェックにより、CI でのロックファイルの検証に関する誤った主張が示されました。
独自の所見マトリックス (選択)
実際に使用するパイプライン
Haiku → 高速マップと建築ホットスポット
Sonnet 5 → 一次監査 + セキュリティ導入のギャップ
Sonnet 4.6 → CI、ドキュメント

、地雷のオンボーディング
Opus → エージェントに面するサーフェスの脅威レビュー
寓話 → 優先順位付けされた 1 つのバックログにマージ
人間 → チェックアウトで _lint.yml、load.py、README を確認する
このチェーンに代わる単一のモデルはありません。 Opus のみ、または Fable のみを支払うと盲点が残ります。
詳細: 比較モデルレポート.md
また、ビデオ ウォークスルー用の 14 スライドのプレゼンター デッキも構築しました:model-comparison-presentation.html (←/→ ナビゲート、F フルスクリーン)。
6. CTRL NODE ユーザーにとってこれが意味すること
モデルの選択はワークフローの決定であり、バニティ層の選択ではありません。同じプロジェクトと作業ディレクトリ上で、Haiku を使用して偵察し、Sonnet を監査し、Opus を脅威に、Fable を使用して計画を立てます。
このようなタスクでは、WORK DIRECTORY モードが重要です。出力専用のサンドボックスでは、CI、コア、パートナー パッケージ全体で file:line の引用が生成されません。
ファブルは、ソネットやオーパスの代わりにではなく、発見後にスロットを獲得します。そのA-グレードはOpusに匹敵しました。成果物の形 (マイルストーン、ラチェット、非目標) が最も実行可能でした。
リポジトリで実験を再実行します。 Bridge BASE_PATH の下でクローンを作成し、そこに Claude プロジェクトを指定し、異なる MODEL 値でタスクを 5 回複製します。
完全な実験 (すべてのプロンプト、モデルごとのレポート、および比較デッキ) は、この記事のサポート資料として以下に公開されています。
プロンプトは、すべてのモデルに .md + .html のペアの出力を要求します。このバッチのすべてのモデルは両方のフォーマットを生成しました。
無料で始めましょう — Claude プロジェクトを作成し、Bridge をペアにします。
Bridge マシン上で必要なリポジトリのクローンを作成します。作業ディレクトリを設定します。
異なる MODEL 値 ( claude-fable-5 、 claude-opus-4-8 など) でエージェントを登録します。
監査プロンプトを INSTRUCTIONS 、 assign 、 RUN 、compare の出力に貼り付けます。
ご質問がありますか、またはスタック上でこれを実行してほしいですか? info@ctrlnode.ai
実験日: 2026 年 6 月 17 日

· CTRL ノード — クロード、副操縦士、ジェミニ、カーソルなどを 1 つのコントロール プレーンから調整します。
AI エージェント用のリモート タスク ランチャー。
クロード、副操縦士、ジェミニ、コーデックス、
カーソル、ヘルメス、オープンクロウ。
info@ctrlnode.ai
@ctrlnodeai
@ctrlnodeAI
製品

## Original Extract

A controlled experiment: one principal-engineer audit prompt, five Claude tiers, one LangChain clone. How we ran it in CTRL NODE, what Fable delivered, and why no single model wins.

We Ran a Complex Task — A LangChain Repo Analysis with Five Claude Models | CTRL NODE
CTRL NODE Product Workflows Kanban Routines Bridge Providers Resources Changelog Documentation Quick Start GitHub News
Sign In
Start for free
Product Workflows Kanban Routines Bridge Providers Resources Changelog Documentation Quick Start GitHub News
Sign In
Start for free
Engineering · Jul 2, 2026 · 11 min read We Ran a Complex Task — A LangChain Repo Analysis with Five Claude Models
Anthropic just shipped Claude Fable . We wanted a real answer to a practical question:
If you run the same complex engineering task on Opus, Fable, Sonnet, and Haiku — what do you actually get back?
Not a benchmark score. Not a vibe check. A full principal-engineer audit of a production open-source monorepo — with evidence, severity labels, and an execution plan.
We ran that experiment inside CTRL NODE : one prompt, five agents, five models, one cloned repository.
1. The goal: one hard task, five models
We gave every model the same four-phase audit prompt and the same target : the LangChain Python monorepo (a large, mature library ecosystem — not a toy repo).
Repository Map — explore first, judge second
Audit Report — architecture, security, tests, performance, deps, DX, docs (with file:line citations)
Improvement Strategy — themes, trade-offs, measurable “done” criteria
Task Plan — milestones M0–M3, quick wins, effort/risk/deps on each item
Every finding must be evidence-based . Guessing is explicitly forbidden.
That is a genuinely heavy task: thousands of files, real CI configs, security-sensitive deserialization paths, and god-class modules on hot code paths. It is the kind of work teams normally spread across several senior engineers.
Fable is positioned as a strong reasoning model for long, structured work. We included it alongside:
The hypothesis was not “Fable wins everything.” It was: each tier sees different things , and Fable might be the best at turning findings into a shippable backlog .
The full prompt lives in our catalog as langchain-prompt.md . Core instruction (abbreviated):
You are a world-class, principal-engineer-level software engineer and technical audit expert.
Perform an in-depth analysis of this code repository, provide an honest audit report,
and offer a prioritized, actionable improvement plan.
Follow four phases in order: Discovery → Audit → Strategy → Task Plan.
All judgments must cite real file paths and line numbers. Do not guess.
Deliverables requested per run:
audit-report-<model>.md — full Markdown report
audit-report-<model>.html — interactive dark-theme dashboard (tabs: Overview, Map, Audit, Strategy, Tasks)
Summary of the prompt: resumen-langchain-prompt.md .
2. How we set it up in CTRL NODE
We did not paste the prompt into five browser tabs. We ran it the way a team would : Bridge on a real machine, a project work directory pointing at the clone, one agent per model tier.
Bridge ( ctrlnode ) installed and paired — see Bridge setup .
Claude SDK API key set in ~/.ctrlnode/.env (providers load automatically — no PROVIDERS flag needed):
ANTHROPIC_API_KEY=sk-ant-...
BASE_PATH=/home/you/workspace
LangChain cloned on the Bridge host under BASE_PATH (CTRL NODE does not git-clone for you; the work directory points at an existing folder).
The work directory is what lets agents read the full tree in WORK DIRECTORY task mode — the same scope a staff engineer would need.
Team → + ADD AGENT — we created five agents on the same project:
Models are selected in the MODEL combobox (synced from Bridge when online) or typed manually. Fable appears as claude-fable-5 in the Bridge model manifest (v2026.2.4+).
Optional AGENT SYSTEM INSTRUCTIONS were left minimal — we wanted the task prompt to carry the spec, not per-agent persona drift.
For each agent, same procedure:
TITLE : LangChain principal audit — <model>
INSTRUCTIONS : paste full contents of langchain-prompt.md
ASSIGN TO AGENT : pick the matching agent chip
OUTPUT MODE : WORK DIRECTORY (full repo scope; optional focus paths left empty)
NEW TASK → task lands in Backlog
RUN → dispatches to Bridge → agent moves to In progress
Bridge delivers the task with repositoryPaths and repo dispatch context so the Claude SDK runs against the LangChain tree on disk. Outputs ( audit-report-*.md / .html ) were collected from the agent’s work directory and copied into our marketing catalog folder.
Tip for reproducibility: use the same commit SHA for every run. Our reports reference LangChain master at 2b47357 where noted.
Fable graded the repo A− — the same calibration as Opus, more honest than Haiku’s self-awarded A .
Complexity concentration — five files exceed 1,800 lines; runnables/base.py is 6,574 LOC . High blast radius on every invoke/stream path.
Unsafe-by-default deserialization — langchain_core.load defaults to allowed_objects='core' , documented as unsafe for untrusted manifests. Safe options exist but are opt-in.
Type-safety escape hatches — 208 type: ignore comments in langchain-core alone; disallow_any_generics=false weakens the public API contract.
Flip deserialization default to a safe allowlist ( 'messages' ) on the next major version.
Burn down parked lint TODOs ( BLE , ANN401 , ERA ) — enforcement infra already exists.
Decompose the top god files behind unchanged public façades (zero API break).
Fable’s differentiator was not a hotter take on security headlines. It was Phase 3 and Phase 4 :
Four strategic themes (complexity, switched-off guardrails, safe-by-default trust boundaries, workspace hygiene)
Explicit non-goals (e.g. don’t rewrite vendored mustache.py this cycle — add property tests instead)
Milestones M0–M3 with workload badges (S/M/L/XL), risk, dependencies, and acceptance criteria
Quick wins you could ship in an afternoon ( .gitignore for audit artifacts, logger.debug on swallowed AttributeError in callbacks/usage.py , CI ratchet on type: ignore count)
Near-exclusive Fable findings:
Vendored 704-line Mustache engine ( mustache.py ) with its own security surface
McCabe C90 complexity lint explicitly disabled — no automated backpressure on god-file growth
Thin test breadth vs complexity for langchain_v1/agents/factory.py (56 test files vs 1,891-line factory)
Fable did not surface several issues other models caught:
TOCTOU / DNS rebinding on SSRF paths (Opus)
ShellToolMiddleware host execution by default (Opus)
SSRF transport adopted in only two call sites + unprotected graph_mermaid.py fetch (Sonnet 5)
Commented lockfile check in CI _lint.yml (Sonnet 4.6)
Broken README model example / missing SECURITY.md (Sonnet 4.6)
That gap is the point: Fable is not a replacement for a multi-model pipeline.
Full report: audit-report-fable.md · Interactive dashboard: audit-report-fable.html
5. How the five models compare
*Haiku’s A looks confident on paper. Cross-checking against Sonnet 4.6 showed a wrong claim about lockfile validation in CI.
Exclusive findings matrix (selected)
The pipeline we’d actually use
Haiku → fast map & architecture hotspots
Sonnet 5 → primary audit + security adoption gaps
Sonnet 4.6 → CI, docs, onboarding landmines
Opus → threat review for agent-facing surfaces
Fable → merge into one prioritized backlog
Human → verify _lint.yml, load.py, README in your checkout
No single model replaces this chain. Paying only for Opus — or only for Fable — leaves blind spots.
Deep dive: comparison-models-report.md
We also built a 14-slide presenter deck for video walkthroughs: model-comparison-presentation.html (←/→ navigate, F fullscreen).
6. What this means for CTRL NODE users
Model choice is a workflow decision , not a vanity tier pick. Use Haiku to scout, Sonnet to audit, Opus for threats, Fable to plan — on the same project and work directory .
WORK DIRECTORY mode matters for tasks like this. An output-only sandbox would not have produced file:line citations across CI, core, and partner packages.
Fable earns a slot after discovery , not instead of Sonnet or Opus. Its A− grade matched Opus; its deliverable shape (milestones, ratchets, non-goals) was the most actionable.
Re-run the experiment on your repo — clone under Bridge BASE_PATH , point a Claude project at it, duplicate the task five times with different MODEL values.
The full experiment — every prompt, per-model report, and the comparison deck — is published below as supporting material for this article.
The prompt asks every model for paired .md + .html outputs. Every model in this batch produced both formats.
Start free — create a Claude project and pair Bridge.
Clone the repo you care about on the Bridge machine; set WORK DIRECTORY .
Register agents with different MODEL values ( claude-fable-5 , claude-opus-4-8 , …).
Paste the audit prompt into INSTRUCTIONS , assign, RUN , compare outputs.
Questions or want us to run this on your stack? info@ctrlnode.ai
Experiment date: 17 June 2026 · CTRL NODE — orchestrate Claude, Copilot, Gemini, Cursor, and more from one control plane.
Remote task launcher for AI agents.
Claude, Copilot, Gemini, Codex,
Cursor, Hermes & OpenClaw.
info@ctrlnode.ai
@ctrlnodeai
@ctrlnodeAI
Product
