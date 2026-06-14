---
source: "https://github.com/001TMF/harness-forge"
hn_url: "https://news.ycombinator.com/item?id=48526233"
title: "Relent less AI self-evolution"
article_title: "GitHub - 001TMF/harness-forge: Turn Claude Code into its own Meta-Harness — a skill that evolves the scaffolding around a fixed model (memory, retrieval, context, prompts) via a native propose→score→Pareto loop. Native reimplementation of Meta-Harness (Lee et al. 2026). · GitHub"
author: "proteus-design"
captured_at: "2026-06-14T12:45:51Z"
capture_tool: "hn-digest"
hn_id: 48526233
score: 1
comments: 0
posted_at: "2026-06-14T11:29:23Z"
tags:
  - hacker-news
  - translated
---

# Relent less AI self-evolution

- HN: [48526233](https://news.ycombinator.com/item?id=48526233)
- Source: [github.com](https://github.com/001TMF/harness-forge)
- Score: 1
- Comments: 0
- Posted: 2026-06-14T11:29:23Z

## Translation

タイトル: 絶え間ない AI の自己進化
記事のタイトル: GitHub - 001TMF/harness-forge: クロード コードを独自のメタハーネスに変える — ネイティブの提案→スコア→パレート ループを介して、固定モデル (メモリ、取得、コンテキスト、プロンプト) の周囲の足場を進化させるスキルです。 Meta-Harness のネイティブ再実装 (Lee et al. 2026)。 · GitHub
説明: クロード コードを独自のメタハーネスに変換します。これは、ネイティブの提案→スコア→パレート ループを介して、固定モデル (メモリ、取得、コンテキスト、プロンプト) の周りの足場を進化させるスキルです。 Meta-Harness のネイティブ再実装 (Lee et al. 2026)。 - 001TMF/ハーネスフォージ

記事本文:
GitHub - 001TMF/harness-forge: クロード コードを独自のメタハーネスに変える — ネイティブの提案→スコア→パレート ループを介して、固定モデル (メモリ、取得、コンテキスト、プロンプト) の周りの足場を進化させるスキルです。 Meta-Harness のネイティブ再実装 (Lee et al. 2026)。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。参照先へリロード

セッションをしてください。
アラートを閉じる
{{ メッセージ }}
001TMF
/
ハーネス鍛造
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .claude-plugin .claude-plugin アセット アセットの例/メモリの概要の例/メモリの概要のスキル/メタハーネスのスキル/メタハーネス .gitignore .gitignore ライセンス ライセンス README.md README.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
クロード コードを独自のメタハーネスに変換します。固定モデルの周りの足場をネイティブに進化させます。
Harness Forge は、エンドツーエンドで実行するクロード コード スキルです。
ハーネス最適化ループ — 提案 → スコア → パレートベストを維持 → 繰り返し — を改善するために
固定モデルに関するコード: そのメモリ、取得、コンテキスト構築、要約、プロンプト
テンプレートとツール選択ロジック。モデルは決して変わりません。足場が良くなります。
これはメソッドのネイティブな再実装です。
Meta-Harness: モデル ハーネスのエンドツーエンドの最適化
(リー、ナイル、チャン、リー、ハッタブ、フィン、2026)。オリジナル
リファレンス リポジトリには、約 1,260 行の Python が同梱されています
( claude_wrapper.py + meta_harness.py ) その仕事は、首なしクロードを駆動することです。
セッション、その出力の解析、ツール呼び出しの追跡、すべてのログの記録、ループ。クロードコードの中では、
ランタイムは第一級のツールとしてすでに存在しています。したがって、Harness Forge は既約領域のみを保持します
ロジック (安価なスコアラー) であり、外側のループ全体をネイティブ オーケストレーションとして表現します。全体
検索行数は最大 1,260 行ではなく、最大 75 行になります。
既存のハーネス（打ち負かすもの）でフロンティアに種を蒔く
繰り返します：
k 個の候補ハーネス バリアントを提案する ← 並列提案者エージェントがコードを書く
各インポート/型チェックを検証する

クス
Hold-out-protected eval でそれぞれスコアを付ける ← $0、決定的スコアラー
FRONTIER パレートマージ: 品質の向上、コストの削減、フロアの尊重
決勝: アンタッチドテストスプリットでフロンティアを 1 回獲得する
提案者は突然変異オペレーターです。フロンティアは検索記憶です。モデルはフリーズしています
全体を通して — これがまさに、これが固定/既製の API デプロイメントに適合する理由です。
重量を変更することはできず、ゲインはハーネスから得る必要があります。
論文の見出しの結果は、テキスト上のコンテキスト トークンが約 4 分の 1 で、精度が +7.7 ポイントでした
分類 — 純粋にハーネス側の勝利です。 Harness Forge は、その結果の形状をネイティブに再現します。
claude_wrapper.py は手動のエージェント ランタイムです。 Claude Code はエージェント ランタイムです。だから、すべての
オーケストレーション部分にはネイティブ同等のものがあり、Python ドライバーは冗長になります。
あなたがまだ書いているのは、安っぽいスコアラー + ルーブリック + 候補者インターフェイスだけです。すべて
オーケストレーション形式は無料です。
1. スキルを 1 行でインストールします。
カール -fsSL https://raw.githubusercontent.com/001TMF/harness-forge/main/install.sh |バッシュ
または、Claude Code プラグインとして (Claude Code 内):
/プラグイン マーケットプレイス 001TMF/harness-forge を追加
/プラグインのインストール harness-forge@tmf-skills
他の方法
# プロジェクトスコープ (./.claude/skills、このリポジトリのみ)
カール -fsSL https://raw.githubusercontent.com/001TMF/harness-forge/main/install.sh | bash -s -- --プロジェクト
# skill.sh 経由 (vercel-labs/skills)
npx スキル追加 001TMF/harness-forge --skill meta-harness -a claude-code
# マニュアル
git クローン https://github.com/001TMF/harness-forge.git
cp -r harness-forge/スキル/メタハーネス ~ /.claude/スキル/メタハーネス
ハーネス、足場、プロンプト システム、メモリ、または
取得ポリシー、またはサマライザー — またはメタハーネス スキルとして直接呼び出します。
2.R

動作した例 ($0、モデルなし、ネットワークなし):
cd harness-forge/examples/memory-summary
Pythonスコア_ベースライン.py
# ->baseline_incumbent fidelity=1.000 chars=269 (打ち負かすシステム)
3. 実際の検索を実行します。例のループ スクリプトを使用してワークフロー ツールを呼び出します。
ワークフロー({ scriptPath: "<abs>/examples/memory-summary/native_meta_harness_workflow.js",
args: { dir: "<abs>/examples/memory-summary"、ラウンド: 2、k: 3 } })
Proposer エージェントは Claude サブスクリプションで実行されます。スコアラーは $0 です。ソルバーモデルは存在せず、
従量制の API はありません。ラウンドが成功すると、269 文字未満の忠実度を維持するコンプレッサーが生成されます。
あなたが提供するもの（5つのブロック）
ループはネイティブです。ドメインはあなたのものです。テンプレートは入っています
スキル/メタハーネス/アセット/ ;ハウツーは入っています
参照/ビルディングブロック.md :
候補インターフェイス — 1 つのクリーンで交換可能な境界 (ABC / プロトコル)。
$0 の決定論的スコアラー + ルーブリック — 内部ループ。何百回も実行されるため、LLM も必要ありません
ネットワーク。候補者によって異なる必要があります (以下のトラップを参照)。
保持された分割を含む評価コーパス。
事前提案者 — 提案者をメカニズムレベルの変更に向けて導くミニスキル (そうではありません)
定数チューニング)、評価セットのリークを禁止します。
フロンティア + 実行ログ — 状態。スクリプト/pareto.py
フロアを考慮したフロンティアを決定論的に計算します。
素朴なハーネス検索を沈める唯一の罠
フリーズリプレイの不具合。スコアラーがキャッシュされた出力 (記録された実行、フリーズされた出力) を再生する場合、
トレース）、スキャフォールディング候補は記録された結果を変更できません。コスト軸のみが移動します。
単純な「品質を最大化し、コストを最小化する」検索は、コンテキストを空にすることで成功します。
凍結された品質スコアは決して低下せず、自信に満ちた無意味なフロンティアを生み出します。
テスト: 「大幅に異なる候補者と交換した場合、この数値は品質に影響を与えるかどうか」
理由は？」

コストのみを移動できる場合は、フリーズした出力を再生することになります。
修正: 候補者が本当に管理しているもの (検索の関連性、圧縮) を採点します。
忠実さ、反事実的な決定）、および/または一方的な害を及ぼさないフロアとしての実行品質
最大化軸よりも。このスキルがこれを実現します。加えて、堅固な規律、反グッドハートのフロア、そして
漏れ防止 — 耐荷重性。での完全な治療
参照/メソッド.md 。
ハーネス鍛造/
§── .claude-plugin/marketplace.json # クロードコードプラグインとしてインストール可能
§── install.sh # 一行curl|bashインストール
§── スキル/
│ └── メタハーネス/ # インストール可能なスキル
│ §── SKILL.md # いつ、何を、ループ、5つのブロック、ガードレール
│ §── 参考文献/ # メソッド · ネイティブ実行 · ビルディングブロック · 動作した例
│ §── アセット/ # テンプレート: ワークフロー ループ、スコアラー、インターフェイス、事前プロポーザー
│ └── scripts/pareto.py # 再利用可能な床を尊重したパレートフロンティア
└── 例/
└──memory-summary/ # 完全な実行可能な検索 ($0 のデモ + ネイティブ ループ)
これを使用する場合 (および使用しない場合)
基本モデルが固定されており、繰り返しタスクがあり、安価な測定可能な評価がある場合に使用します。
存在します (または構築できます)。つまり、ゲインはハーネスから得られる必要があります。従来のターゲット: コンテキスト
肥大化、弱い検索、損失の多い要約、脆い即時足場。
ゲインをモデルの重みから取得する必要がある場合は使用しないでください (代わりに RL / 微調整を実行します)。
安定した評価ループはありません。 Meta-Harness と RL は補完的です: 固定ベースモデルにおいて
この段階では、Harness Forge が唯一利用可能なオプティマイザであり、評価強化を後で強制的に実行します。
RL フェーズも、ほぼゼロのコストで依存します。参照
参考文献/メソッド.md §6。
このメソッドは、Yoonho Lee、Roshen Nair、Qizheng Zhang、Kangwook Lee による Meta-Harness です。

オマル
ハタブとチェルシー・フィン。このリポジトリは、クロード コードとして独立したネイティブ再実装です。
スキル;元のリポジトリからのコードは提供されません。使用する場合は、論文を引用してください。
@misc { lee2026metaharnessendtoendoptimizationmodel 、
title = { メタハーネス: モデル ハーネスのエンドツーエンドの最適化 } ,
著者 = { ユンホ・リー、ローシェン・ネール、チージェン・チャン、カンウク・リー、オマール・ハッタブ、チェルシー・フィン } 、
年 = { 2026 } 、
eprint = { 2603.28052 } 、
archivePrefix = { arXiv } 、
プライマリクラス = { cs.AI } 、
URL = { https://arxiv.org/abs/2603.28052 } 、
}
論文: https://arxiv.org/abs/2603.28052
リファレンス実装: https://github.com/stanford-iris-lab/meta-harness
クロード コードを独自のメタハーネスに変換します。これは、ネイティブの提案→スコア→パレート ループを介して、固定モデル (メモリ、取得、コンテキスト、プロンプト) の周りの足場を進化させるスキルです。 Meta-Harness のネイティブ再実装 (Lee et al. 2026)。
読み込み中にエラーが発生しました。このページをリロードしてください。
2
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Turn Claude Code into its own Meta-Harness — a skill that evolves the scaffolding around a fixed model (memory, retrieval, context, prompts) via a native propose→score→Pareto loop. Native reimplementation of Meta-Harness (Lee et al. 2026). - 001TMF/harness-forge

GitHub - 001TMF/harness-forge: Turn Claude Code into its own Meta-Harness — a skill that evolves the scaffolding around a fixed model (memory, retrieval, context, prompts) via a native propose→score→Pareto loop. Native reimplementation of Meta-Harness (Lee et al. 2026). · GitHub
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
001TMF
/
harness-forge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .claude-plugin .claude-plugin assets assets examples/ memory-summary examples/ memory-summary skills/ meta-harness skills/ meta-harness .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh View all files Repository files navigation
Turn Claude Code into its own Meta-Harness — evolve the scaffolding around a fixed model, natively.
Harness Forge is a Claude Code skill that runs an end-to-end
harness-optimization loop — propose → score → keep the Pareto-best → repeat — to improve the
code around a fixed model: its memory, retrieval, context construction, summarization, prompt
templates, and tool-selection logic. The model never changes; the scaffolding gets better.
It is a native reimplementation of the method in
Meta-Harness: End-to-End Optimization of Model Harnesses
(Lee, Nair, Zhang, Lee, Khattab & Finn, 2026). The original
reference repo ships ~1,260 lines of Python
( claude_wrapper.py + meta_harness.py ) whose job is to drive a headless Claude : spawn a
session, parse its output, track tool calls, log everything, loop. Inside Claude Code, that
runtime already exists as first-class tools. So Harness Forge keeps only the irreducible domain
logic — a cheap scorer — and expresses the entire outer loop as native orchestration. The whole
search becomes ~75 lines instead of ~1,260.
seed the frontier with the incumbent harness (the thing to beat)
repeat:
PROPOSE k candidate harness variants ← parallel proposer agents write code
VALIDATE each imports / type-checks
SCORE each on a held-out-protected eval ← a $0, deterministic scorer
FRONTIER Pareto-merge: quality up, cost down, floor-respecting
final: score the frontier once on the untouched test split
The proposer is the mutation operator. The frontier is the search memory. The model is frozen
throughout — which is exactly why this fits a fixed / off-the-shelf-API deployment, where you
can't change the weights and the gain has to come from the harness.
The paper's headline result was +7.7 accuracy points at ~4× fewer context tokens on text
classification — a pure harness-side win. Harness Forge reproduces that shape of result natively.
claude_wrapper.py is a hand-rolled agent runtime. Claude Code is an agent runtime. So every
orchestration piece has a native equivalent, and the Python driver becomes redundant:
The only thing you still write is the cheap scorer + rubric + candidate interface . Everything
orchestration-shaped is free.
1. Install the skill — one line:
curl -fsSL https://raw.githubusercontent.com/001TMF/harness-forge/main/install.sh | bash
Or as a Claude Code plugin (inside Claude Code):
/plugin marketplace add 001TMF/harness-forge
/plugin install harness-forge@tmf-skills
Other ways
# project-scoped (./.claude/skills, this repo only)
curl -fsSL https://raw.githubusercontent.com/001TMF/harness-forge/main/install.sh | bash -s -- --project
# via skills.sh (vercel-labs/skills)
npx skills add 001TMF/harness-forge --skill meta-harness -a claude-code
# manual
git clone https://github.com/001TMF/harness-forge.git
cp -r harness-forge/skills/meta-harness ~ /.claude/skills/meta-harness
It auto-triggers when you talk about optimizing a harness, scaffold, prompt system, memory or
retrieval policy, or summarizer — or invoke it directly as the meta-harness skill.
2. Run the worked example ($0, no model, no network):
cd harness-forge/examples/memory-summary
python score_baselines.py
# -> baseline_incumbent fidelity=1.000 chars=269 (the system to beat)
3. Run a real search — invoke the Workflow tool with the example's loop script:
Workflow({ scriptPath: "<abs>/examples/memory-summary/native_meta_harness_workflow.js",
args: { dir: "<abs>/examples/memory-summary", rounds: 2, k: 3 } })
Proposer agents run on your Claude subscription; the scorer is $0; there is no solver model and
no metered API . A successful round produces a compressor holding fidelity at < 269 chars .
What you supply (the five blocks)
The loop is native; the domain is yours. Templates are in
skills/meta-harness/assets/ ; how-to is in
references/building-blocks.md :
Candidate interface — one clean, swappable boundary (an ABC / Protocol).
A $0 deterministic scorer + rubric — the inner loop; runs hundreds of times, so no LLM, no
network. It must vary with the candidate (see the trap below).
An eval corpus with a held-out split.
A proposer prior — a mini-skill steering proposers toward mechanism-level changes (not
constant-tuning) and forbidding eval-set leakage.
A frontier + run log — the state. scripts/pareto.py
computes the floor-respecting frontier deterministically.
The one trap that sinks naive harness searches
The frozen-replay defect. If your scorer replays cached outputs (a recorded run, a frozen
trace), a scaffolding candidate cannot change the recorded result — only the cost axis moves.
A naive "maximize quality, minimize cost" search then wins by emptying the context while the
frozen quality score never drops, producing a confident, meaningless frontier.
Test: "If I swap in a wildly different candidate, can this number change for a quality
reason?" If only cost can move, you are replaying frozen outputs.
Fix: grade something the candidate genuinely controls (retrieval relevance, compression
fidelity, a counterfactual decision), and/or run quality as a one-sided do-no-harm floor rather
than a maximize axis. The skill makes this — plus held-out discipline, an anti-Goodhart floor, and
anti-leakage — load-bearing. Full treatment in
references/method.md .
harness-forge/
├── .claude-plugin/marketplace.json # installable as a Claude Code plugin
├── install.sh # one-line curl|bash install
├── skills/
│ └── meta-harness/ # the installable skill
│ ├── SKILL.md # what/when, the loop, the 5 blocks, the guardrails
│ ├── references/ # method · native-execution · building-blocks · worked example
│ ├── assets/ # templates: workflow loop, scorer, interface, proposer prior
│ └── scripts/pareto.py # reusable floor-respecting Pareto frontier
└── examples/
└── memory-summary/ # a complete, runnable search (the $0 demo + the native loop)
When to use this (and when not)
Use it when the base model is fixed, there are repeated tasks, and a cheap measurable eval
exists (or can be built) — i.e. the gain has to come from the harness. Classic targets: context
bloat, weak retrieval, lossy summarization, brittle prompt scaffolds.
Don't when the gain must come from the model weights (do RL / fine-tuning instead), or when
there is no stable evaluation loop. Meta-Harness and RL are complementary: in a fixed-base-model
phase, Harness Forge is the only available optimizer — and it forces the eval-hardening a later
RL phase also depends on, at near-zero cost. See
references/method.md §6.
The method is Meta-Harness by Yoonho Lee, Roshen Nair, Qizheng Zhang, Kangwook Lee, Omar
Khattab, and Chelsea Finn. This repo is an independent native reimplementation as a Claude Code
skill; it vendors no code from the original repo. If you use it, please cite the paper:
@misc { lee2026metaharnessendtoendoptimizationmodel ,
title = { Meta-Harness: End-to-End Optimization of Model Harnesses } ,
author = { Yoonho Lee and Roshen Nair and Qizheng Zhang and Kangwook Lee and Omar Khattab and Chelsea Finn } ,
year = { 2026 } ,
eprint = { 2603.28052 } ,
archivePrefix = { arXiv } ,
primaryClass = { cs.AI } ,
url = { https://arxiv.org/abs/2603.28052 } ,
}
Paper: https://arxiv.org/abs/2603.28052
Reference implementation: https://github.com/stanford-iris-lab/meta-harness
Turn Claude Code into its own Meta-Harness — a skill that evolves the scaffolding around a fixed model (memory, retrieval, context, prompts) via a native propose→score→Pareto loop. Native reimplementation of Meta-Harness (Lee et al. 2026).
There was an error while loading. Please reload this page .
2
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
