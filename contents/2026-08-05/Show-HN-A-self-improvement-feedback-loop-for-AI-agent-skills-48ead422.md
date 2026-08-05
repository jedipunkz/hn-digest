---
source: "https://github.com/Carlo1911/skill-evolution"
hn_url: "https://news.ycombinator.com/item?id=49189270"
title: "Show HN: A self-improvement feedback loop for AI agent skills"
article_title: "GitHub - Carlo1911/skill-evolution: Self-improvement feedback loop for AI agent skills — analyzes past sessions, drafts structured improvement proposals, and gates auto-apply behind an evaluation framework. Host-agnostic (Hermes, Claude Code, and any HostAdapter). · GitHub"
author: "Carlo1911"
captured_at: "2026-08-05T22:08:17Z"
capture_tool: "hn-digest"
hn_id: 49189270
score: 2
comments: 0
posted_at: "2026-08-05T21:28:05Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A self-improvement feedback loop for AI agent skills

- HN: [49189270](https://news.ycombinator.com/item?id=49189270)
- Source: [github.com](https://github.com/Carlo1911/skill-evolution)
- Score: 2
- Comments: 0
- Posted: 2026-08-05T21:28:05Z

## Translation

タイトル: HN を表示: AI エージェント スキルの自己改善フィードバック ループ
記事のタイトル: GitHub - Carlo1911/skill-evolution: AI エージェント スキルの自己改善フィードバック ループ - 過去のセッションを分析し、構造化された改善提案を作成し、評価フレームワークの背後でゲートが自動適用されます。ホストに依存しない (Hermes、Claude Code、および任意の HostAdapter)。 · GitHub
説明: AI エージェント スキルの自己改善フィードバック ループ — 過去のセッションを分析し、構造化された改善提案を作成し、評価フレームワークの背後でゲートが自動適用されます。ホストに依存しない (Hermes、Claude Code、および任意の HostAdapter)。 - Carlo1911/スキル進化

記事本文:
GitHub - Carlo1911/skill-evolution: AI エージェント スキルの自己改善フィードバック ループ - 過去のセッションを分析し、構造化された改善提案を作成し、評価フレームワークの背後でゲートが自動適用されます。ホストに依存しない (Hermes、Claude Code、および任意の HostAdapter)。 · GitHub
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
カルロ1911
/
スキル進化
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット。

github/ workflows .github/ workflows scripts scripts testing testing .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md SKILL.md SKILL.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイル ナビゲーション
エージェントのスキルを自律的に自己向上させます。
過去のセッションを読み取り、ロードされたスキルと照らし合わせて分析し、構造化された提案を生成し、オプションで信頼性の高い改善を自動適用します。ホストに依存しない: 新しいホスト用の拡張可能な HostAdapter インターフェイスを備えた、Hermes および Claude Code 用のアダプターが同梱されています。
デーモンはありません。外部サービスはありません。 GPUはありません。ホスト エージェント、スケジュールされたジョブ、および Python stdlib だけです。
以下のすべては適切なデフォルト値を持っていますが、次の 2 つはオプションではありません。
黙って何もしない、またはスタックするのではなく、実際にエンドツーエンドで動作するパイプライン:
それ以外はすべてオプションであり、デフォルトを変更する場合にのみ設定する必要があります。
完全なリストについては、CLAUDE.md および SKILL.md を参照してください。この表は、必要な最低限のものです。
最初の実行前に読んでください。完全なリファレンスではありません。
# スキルをインストールする
hermes スキルインストール https://raw.githubusercontent.com/Carlo1911/skill-evolution/main/SKILL.md
# ロードしてください
エルメスのスキル進化
# 分析を実行する
「最近のセッションでスキル進化分析を実行する」
クロード コード: リポジトリをクローンし、スキルをディレクトリとしてロードします。セット
SKILL_EVOLUTION_HOST=claude_code なので、読み取り/書き込みアダプターは ~/.claude にルーティングされます。
ホストのセッション データベースから未処理のセッションを取得します。
各セッションを分析してカバレッジギャップがないか確認する
./proposals/ にプロポーザル ファイルを生成します ( SKILL_EVOLUTION_PROPOSALS_DIR でオーバーライドします)。
デフォルトでは、プロポーザルはレビューのみです。自動適用を有効にする:
エクスポート SKILL_EVOLUTION_AUTO_APPLY=true
エクスポート SKILL_EVOLUTION_MIN_CONFIDENCE=0.85
使用例
Clau の完全な手動ウォークスルー

コードで、「何か新しく分析するものはありますか」から
実際に提案が適用されます。これは、スケジュールされたジョブが 1 ステップで自動化する内容を反映しています。
一度に。
# 0. 評価ゲートのホストとプロバイダーを構成します (「最小構成」を参照)
SKILL_EVOLUTION_HOST=claude_code をエクスポート
エクスポート ANTHROPIC_API_KEY=sk-ant-...
# 1. 何も処理済みとしてマークせずに、未処理のセッションがいくつ存在するかを確認します。
python3 scripts/fetch_sessions.py --dry-run --lookback-hours 72
# → 「未処理のセッションが 14 個見つかりました」
# 2. スキルがロードされたクロード コード セッションで、次のように尋ねます。
# 「最近のセッションでスキル進化分析を実行する」
#
# 内部では、そのプロンプトは、このリポジトリがスクリプトとして提供するのと同じパイプラインを実行します —
#analyze.py によってフォーマットされた fetch_sessions.py の出力と、skill_index.py のスキル
# インデックス — 両方をホスト エージェント自身の推論 (LLM 分析ステップ) にフィードします。
# はこのリポジトリのスクリプトではありません。エージェントが直接実行します):
python3 scripts/fetch_sessions.py --lookback-hours 72 | python3 スクリプト/analyze.py > /tmp/sessions.txt
python3 スクリプト/skill_index.py > /tmp/skills.json
# その後、エージェントは調査結果ごとに 1 つの提案ファイルを ./proposals/ に書き込み、マークを付けます。
# それらのセッションが処理されました。
#3. 提案された内容を検証する
python3 スクリプト/proposal.py --list
python3 scripts/proposal.py --show 20260804-001
# 4. 提案ファイルを確認し、同意する場合は手動で承認します。
#proposals/20260804-001.mdを編集し、「ステータス: 提案済み」を「ステータス: 承認済み」に変更します。
＃5.それを適用します。このための CLI は意図的にありません。apply_proposal() は、
# 分析セッション自体ではなく、人間または作成したステップによって呼び出されます。
# この呼び出しは最初に完全な評価ゲートを実行し、合格した場合にのみスキルを変更します。
Python3 -c "
システムをインポートします。 sys.path.insert(0, 'スクリプト')
プロポから

salインポートload_proposal、apply_proposal
p =load_proposal('proposals/20260804-001.md')
結果 = apply_proposal(p, min_confidence=0.85)
print('can_apply:', result['can_apply'])
print('理由:', result.get('理由', result.get('評価結果')))
」
can_apply が何を意味するかはホストによって異なります。クロード コードでは、True はスキル ファイルを意味します
すでに書き換えられていました
[切り捨てられた]
このリポジトリには既製の cron/job プロンプトは同梱されていません。これはオペレーター固有です (どのように
スケジュール、何を配信するか、どのホスト上にいるかなど）、自分の仕事に属します
このリポジトリにはない構成です。ジョブ プロンプトで実行する必要があること: 実行
scripts/fetch_sessions.py + scripts/skill_index.py 、出力を LLM 分析に渡します
SKILL_EVOLUTION_* 環境変数を使用して正しいパスを見つけて書き込みさせるステップ
proposal.py のスキーマに従った提案。 SKILL.md の「スケジュールされた実行 (Cron)」を参照してください。
完全な契約に関するセクション。
スキル進化/
§── SKILL.md # スキルファイル (ホストのスキルインストールでインストール可能)
§── README.md # このファイル
§── ライセンス # MIT
§── pyproject.toml # `optimizer` と `embeddings` の追加機能を備えた Python パッケージ
§── scripts/ # スタンドアロン Python ツール
│ §── fetch_sessions.py # ホストセッションデータベースの読み込み → NDJSON
│ §── skill_index.py # スキルをスキャン → JSON インデックス
│ §──analyze.py # LLM 用にセッションをフォーマットする
│ §──proposal.py # プロポーザルスキーマ、I/O、適用ロジック
│ §── host.py # ホストアダプタのシーム (HermesAdapter, ClaudeCodeAdapter, ...)
│ §──evaluate.py # 評価ゲート (deterministic + LLM-judge + regression + opt-in human_review +embedding_similarity)
│ §── optimize_skill.py # オプションの GEPA オプティマイザー (追加の `gepa` が必要)
│ §── skill_quality.py # 定期的なスキル品質の追跡と傾向レポート

│ §── embedding_backends.py # FastEmbed / Ollama / OpenAI / llama.cpp 埋め込みバックエンド
│ §── embedding_similarity.py # 埋め込み類似性評価器 (オプトイン)
│ §── state.py # 処理されたセッションを追跡します (ホストごと)
│ §── skill-evolution-fetch.sh # Cron ラッパー (Hermes)
│ └── skill-quality-report.sh # 品質レポート用の Cron ラッパー
└── テスト/ # pytest スイート (エバリュエーター/機能領域ごとに 1 つのファイル)
仕組み
フローチャート LR
A[fetch_sessions.py] -->|NDJSON| B[転職エージェント]
C[skill_index.py] -->|スキルインデックス| B
B -->|分析| D[提案書 .md ファイル]
D -->|レビュー| E{人間は承認しますか?}
E -->|はい| F[evaluate.py ゲート]
F -->|合格| G[ホストアダプターが適用される]
F -->|失敗| H[提案されたまま]
E -->|いいえ| I[アーカイブ]
読み込み中
しきい値を超える信頼性を持つプロポーザルは、決定的なサイズ チェック (絶対的な上限、パーセンテージと絶対バイト数の両方でのパスごとの成長と縮小、さらにスキルが開始された場所に対して測定された累積ドリフト)、LLM 判定ルーブリック スコア、そのターゲット自身の履歴に対する回帰チェック、およびオプションの対話型人間拒否拒否権 (SKILL_EVOLUTION_EVALUATORS=...,human_review) という評価ゲートを通過する必要があります。ホストアダプターがミューテーションを実行します。
絶対的な上限はラチェットです。すでに限界を超えているスキルは、それ自体よりも大きくないボディに置き換えることができるため、特大のスキルは決して悪化することなく改善可能なままですが、新しいスキルが限界を超えて作成されることはありません。サイズ比較は、プロポーザル自体について報告される「現在の値」ではなく、ディスク上にインストールされている SKILL.md に対して測定します。
提案を適用する自動化されたステップはないことに注意してください。分析の実行は提案を作成して停止します。アナライザー プロンプトは、設定が何であれ、ホストのスキル変更ツールを呼び出すことを禁止します。

アイデンティティ。 apply_proposal() は人間、またはユーザーが作成したステップによって呼び出されます。詳細および SKILL_EVOLUTION_* 環境変数については、SKILL.md を参照してください。
セッションとスキルは、次の方法で選択された HostAdapter ( scripts/host.py ) を通じて読み取られます。
SKILL_EVOLUTION_HOST (デフォルトは hermes )。
HermesAdapter は ~/.hermes/state.db および ~/.hermes/skills/<category>/<skill>/SKILL.md を読み取ります。
apply_proposal() は、 skill_manage 命令の辞書 ( apply_by: Agent ) を発行します。
ClaudeCodeAdapter は ~/.claude/skills/*/SKILL.md および ~/.claude/projects/*/*.jsonl を読み取ります。
apply_proposal() はスキル ファイルを直接書き込み ( apply_by: direct )、アーカイブします
skill/.archive/ の下のソースを非推奨/マージします。
新しいホストは、HostAdapter ABC: 3 つの読み取りメソッド ( iter_sessions 、
iter_skills 、 read_skill_body ) に加えて、supports_write フラグと具象
apply_skill_write(plan) は、ホストの突然変異計画を返します。詳細については CLAUDE.md を参照してください。
完全な契約。
裁判官は、資格情報を持っているプロバイダーに対して実行されます (5 つの stdlib のみ)
アダプター、SDK および LiteLLM なし:
各エバリュエーターはグローバル選択 ( SKILL_EVOLUTION_<EVALUATOR>_PROVIDER ) をオーバーライドできます。
ジャッジは、オプティマイザーのリフレクション ステップとは別の場所で実行できます。すべてのプロンプトは
マシンから送信される前に、秘密の編集と PII マスキングが実行されます。
フレームワークは 4 つの独立したターゲットを 1 つの共有履歴ファイルにスコア付けします: スキル テキスト
(ゲートはデフォルトで自動適用されます)、さらに提案品質、ツール呼び出し品質、
アナライザー プロンプトの品質 — 最後の 3 つのゲートは、アナライザーに追加された場合にのみ自動適用されます。
SKILL_EVOLUTION_GATE_TARGETS ;デフォルトでは、それらは可観測性のみであり、次の方法で検査可能です。
Evaluate.py --eval-target および optimize_skill.py --list-candidates --target all 。参照
完全なコマンド リファレンスについては、SKILL.md を参照してください。
オプションの 5 番目の評価子 embedding_similarity は、

ベクトル埋め込み
セマンティック チェック (重複検出、ドリフト検出、接地検証)。
SKILL_EVOLUTION_EVALUATORS=...,embedding_similarity を介してオプトインされており、必要があります
追加の埋め込み ( pip install -e ".[embeddings]" )。
pip install -e " .[optimizer] " # gepa==0.1.4 をインストールします
pip install -e " .[embeddings] " # fastembed をインストールします (類似性評価器を埋め込むため)
エクスポート SKILL_EVOLUTION_OPTIMIZER_ENABLED=true
# 読み取り専用: 最適化する価値があるほど悪いスコアを獲得したターゲットはどれですか?
python3 スクリプト/optimize_skill.py --list-candidates
# 1 つのスキルのセッション履歴に対して実際の gepa.optimize_anything() ループを実行します
python3 scripts/optimize_skill.py --skill < name > [--iterations N]
--スキルにインストールされた SKILL.md を使用して GEPA をシードし、それに対して候補者をスコア付けします。
スキル自身の記録されたセッション、および勝者からの Improvement_existing 提案の草案を作成します。
他の提案と同じレビュー/評価ゲートを通過し、直接適用されることはありません。
エクストラをインストールしたインタプリタを使用します。gepa は残りの部分では必要ありません。
パイプラインがあるため、これが無い裸の Python3 はすぐに失敗し、実用的なメッセージが表示されます。
インストールされているすべてのスキルの定期的な品質評価。
評価ゲート。各スキルは、正確さ、手順に従っていること、および
簡潔さは eval_history.jsonl に記録され、傾向レポートに集約されます。
# 評価

[切り捨てられた]

## Original Extract

Self-improvement feedback loop for AI agent skills — analyzes past sessions, drafts structured improvement proposals, and gates auto-apply behind an evaluation framework. Host-agnostic (Hermes, Claude Code, and any HostAdapter). - Carlo1911/skill-evolution

GitHub - Carlo1911/skill-evolution: Self-improvement feedback loop for AI agent skills — analyzes past sessions, drafts structured improvement proposals, and gates auto-apply behind an evaluation framework. Host-agnostic (Hermes, Claude Code, and any HostAdapter). · GitHub
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
Carlo1911
/
skill-evolution
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github/ workflows .github/ workflows scripts scripts tests tests .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md SKILL.md SKILL.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Autonomous self-improvement for your agent skills.
Reads past sessions, analyzes them against loaded skills, generates structured proposals, and optionally auto-applies high-confidence improvements. Host-agnostic: ships adapters for Hermes and Claude Code , with an extensible HostAdapter interface for new hosts.
No daemons. No external services. No GPU. Just your host agent, a scheduled job, and Python stdlib.
Everything below has a sane default, but two things are not optional if you want the
pipeline to actually work end to end rather than silently doing nothing or getting stuck:
Everything else is optional and only needs to be set to change a default:
See CLAUDE.md and SKILL.md for the full list — this table is the minimum you need to
read before your first run, not an exhaustive reference.
# Install the skill
hermes skills install https://raw.githubusercontent.com/Carlo1911/skill-evolution/main/SKILL.md
# Load it
hermes -s skill-evolution
# Run analysis
" Run skill evolution analysis on my recent sessions "
Claude Code: clone the repo and load the skill as a directory. Set
SKILL_EVOLUTION_HOST=claude_code so the read/write adapter routes to ~/.claude .
Fetch unprocessed sessions from the host's session database
Analyze each session for coverage gaps
Generate proposal files in ./proposals/ (override with SKILL_EVOLUTION_PROPOSALS_DIR )
By default, proposals are review-only. Enable auto-apply:
export SKILL_EVOLUTION_AUTO_APPLY=true
export SKILL_EVOLUTION_MIN_CONFIDENCE=0.85
Usage Example
A full manual walkthrough on Claude Code, from "do I have anything new to analyze" to
"a proposal is actually applied." This mirrors what a scheduled job automates, one step
at a time.
# 0. Configure the host and a provider for the evaluation gate (see "Minimum Configuration")
export SKILL_EVOLUTION_HOST=claude_code
export ANTHROPIC_API_KEY=sk-ant-...
# 1. Check how many unprocessed sessions exist, without marking anything as processed
python3 scripts/fetch_sessions.py --dry-run --lookback-hours 72
# → "14 unprocessed sessions found"
# 2. In a Claude Code session with the skill loaded, ask:
# "Run skill evolution analysis on my recent sessions"
#
# Under the hood, that prompt runs the same pipeline this repo ships as scripts —
# fetch_sessions.py's output formatted by analyze.py, plus skill_index.py's skill
# index — and feeds both to the host agent's own reasoning (the LLM analysis step
# is not a script in this repo; the agent performs it directly):
python3 scripts/fetch_sessions.py --lookback-hours 72 | python3 scripts/analyze.py > /tmp/sessions.txt
python3 scripts/skill_index.py > /tmp/skills.json
# The agent then writes one proposal file per finding into ./proposals/, and marks
# those sessions processed.
# 3. Inspect what got proposed
python3 scripts/proposal.py --list
python3 scripts/proposal.py --show 20260804-001
# 4. Review the proposal file and, if you agree with it, approve it by hand:
# edit proposals/20260804-001.md, change `status: proposed` to `status: approved`
# 5. Apply it. There is no CLI for this on purpose — apply_proposal() is meant to be
# invoked by a human or by a step you write, never by the analysis session itself.
# This call runs the full evaluation gate first and only mutates the skill if it passes:
python3 -c "
import sys; sys.path.insert(0, 'scripts')
from proposal import load_proposal, apply_proposal
p = load_proposal('proposals/20260804-001.md')
result = apply_proposal(p, min_confidence=0.85)
print('can_apply:', result['can_apply'])
print('reason:', result.get('reason', result.get('evaluation_results')))
"
What can_apply means depends on the host: on Claude Code, True means the skill file
was already rewritten on
[truncated]
This repo doesn't ship a ready-made cron/job prompt — that's operator-specific (how you
schedule it, what it delivers to, which host you're on) and belongs in your own job
configuration, not in this repo. What the job prompt needs to do: run
scripts/fetch_sessions.py + scripts/skill_index.py , hand the output to an LLM analysis
step using the SKILL_EVOLUTION_* env vars to find the right paths, and have it write
proposals following proposal.py 's schema. See SKILL.md 's "Scheduled Runs (Cron)"
section for the full contract.
skill-evolution/
├── SKILL.md # Skill file (installable via host's skill install)
├── README.md # This file
├── LICENSE # MIT
├── pyproject.toml # Python package, with `optimizer` and `embeddings` extras
├── scripts/ # Standalone Python tools
│ ├── fetch_sessions.py # Read host session database → NDJSON
│ ├── skill_index.py # Scan skills → JSON index
│ ├── analyze.py # Format sessions for LLM
│ ├── proposal.py # Proposal schema, I/O, apply logic
│ ├── host.py # Host adapter seam (HermesAdapter, ClaudeCodeAdapter, ...)
│ ├── evaluate.py # Evaluation gate (deterministic + LLM-judge + regression + opt-in human_review + embedding_similarity)
│ ├── optimize_skill.py # Optional GEPA optimizer (needs the `gepa` extra)
│ ├── skill_quality.py # Periodic skill quality tracking and trend reports
│ ├── embedding_backends.py # FastEmbed / Ollama / OpenAI / llama.cpp embedding backends
│ ├── embedding_similarity.py # Embedding-similarity evaluator (opt-in)
│ ├── state.py # Track processed sessions (per-host)
│ ├── skill-evolution-fetch.sh # Cron wrapper (Hermes)
│ └── skill-quality-report.sh # Cron wrapper for quality reports
└── tests/ # pytest suite (one file per evaluator/feature area)
How It Works
flowchart LR
A[fetch_sessions.py] -->|NDJSON| B[job agent]
C[skill_index.py] -->|skill index| B
B -->|analysis| D[Proposal .md files]
D -->|review| E{Human approves?}
E -->|Yes| F[evaluate.py gate]
F -->|passed| G[host adapter applies]
F -->|failed| H[stays proposed]
E -->|No| I[Archive]
Loading
Proposals with confidence above threshold still have to pass the evaluation gate — deterministic size checks (absolute cap, per-pass growth and shrink both as a percentage and as an absolute byte count, plus cumulative drift measured against where the skill started), an LLM-judge rubric score, a regression check against that target's own history, and an optional interactive human-rejection veto ( SKILL_EVOLUTION_EVALUATORS=...,human_review ) — before the host adapter runs the mutation.
The absolute cap is a ratchet : a skill already over the limit can still be replaced by a body no larger than itself, so an oversized skill stays improvable without ever getting worse, while a new skill is never created over the limit. Size comparisons measure against the installed SKILL.md on disk, not against the "current value" a proposal reports about itself.
Note that no automated step applies a proposal. The analysis run writes proposals and stops; the analyzer prompt forbids it from calling the host's skill-mutation tool whatever the confidence. apply_proposal() is invoked by a human, or by a step you write. See SKILL.md for the details and SKILL_EVOLUTION_* environment variables.
Sessions and skills are read through a HostAdapter ( scripts/host.py ), selected via
SKILL_EVOLUTION_HOST (default hermes ).
HermesAdapter reads ~/.hermes/state.db and ~/.hermes/skills/<category>/<skill>/SKILL.md .
apply_proposal() emits skill_manage instruction dicts ( applied_by: agent ).
ClaudeCodeAdapter reads ~/.claude/skills/*/SKILL.md and ~/.claude/projects/*/*.jsonl .
apply_proposal() writes skill files directly ( applied_by: direct ), archiving
deprecate/merge sources under skills/.archive/ .
A new host implements the HostAdapter ABC: three read methods ( iter_sessions ,
iter_skills , read_skill_body ) plus a supports_write flag and a concrete
apply_skill_write(plan) that returns the host's mutation plan. See CLAUDE.md for the
full contract.
The judge runs against whichever provider you have credentials for — five stdlib-only
adapters, no SDK and no LiteLLM:
Each evaluator can override the global choice ( SKILL_EVOLUTION_<EVALUATOR>_PROVIDER ), so
the judge can run somewhere different from the optimizer's reflection step. Every prompt is
run through secret redaction and PII masking before it leaves the machine.
The framework scores four independent targets into one shared history file: skill text
(gates auto-apply by default), plus proposal quality , tool-call quality , and
analyzer-prompt quality — the last three gate auto-apply only when added to
SKILL_EVOLUTION_GATE_TARGETS ; by default they are observability-only, inspectable via
evaluate.py --eval-target and optimize_skill.py --list-candidates --target all . See
SKILL.md for the full command reference.
An optional fifth evaluator, embedding_similarity , uses vector embeddings for
semantic checks (duplicate detection, drift detection, grounding verification).
It is opt-in via SKILL_EVOLUTION_EVALUATORS=...,embedding_similarity and requires
the embeddings extra ( pip install -e ".[embeddings]" ).
pip install -e " .[optimizer] " # installs gepa==0.1.4
pip install -e " .[embeddings] " # installs fastembed (for embedding similarity evaluator)
export SKILL_EVOLUTION_OPTIMIZER_ENABLED=true
# Read-only: which targets scored badly enough to be worth optimizing?
python3 scripts/optimize_skill.py --list-candidates
# Run a real gepa.optimize_anything() loop over one skill's session history
python3 scripts/optimize_skill.py --skill < name > [--iterations N]
--skill seeds GEPA with the skill's installed SKILL.md , scores candidates against that
skill's own recorded sessions, and drafts an improve_existing proposal from the winner —
through the same review/evaluation gate as any other proposal, never applied directly.
Use the interpreter you installed the extra into: gepa is not needed by the rest of the
pipeline, so a bare python3 without it fails fast with an actionable message.
Periodic quality assessment of all installed skills, using the same rubric as the
evaluation gate. Each skill is scored on correctness, procedure-following, and
conciseness, recorded into eval_history.jsonl , and aggregated into a trend report.
# Evalua

[truncated]
