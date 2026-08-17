---
source: "https://github.com/alexisfox7/PRO-LONG"
hn_url: "https://news.ycombinator.com/item?id=49329010"
title: "Programmatic memory for long-horizon LLM agents"
article_title: "GitHub - alexisfox7/PRO-LONG: Programmatic memory for long-horizon LLM agents: the harness appends everything to one log, and the agent searches it with code. 97.4% on ARC-AGI-3 (arXiv:2607.20064) · GitHub"
author: "rzk"
captured_at: "2026-08-17T11:17:01Z"
capture_tool: "hn-digest"
hn_id: 49329010
score: 1
comments: 0
posted_at: "2026-08-17T10:57:44Z"
tags:
  - hacker-news
  - translated
---

# Programmatic memory for long-horizon LLM agents

- HN: [49329010](https://news.ycombinator.com/item?id=49329010)
- Source: [github.com](https://github.com/alexisfox7/PRO-LONG)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T10:57:44Z

## Translation

タイトル: ロングホライズン LLM エージェントのプログラム メモリ
記事のタイトル: GitHub - alexisfox7/PRO-LONG: ロングホライズン LLM エージェント用のプログラム メモリ: ハーネスはすべてを 1 つのログに追加し、エージェントはコードでそれを検索します。 ARC-AGI-3 (arXiv:2607.20064) で 97.4% · GitHub
説明: 長期の LLM エージェント用のプログラム メモリ: ハーネスはすべてを 1 つのログに追加し、エージェントはそれをコードで検索します。 ARC-AGI-3 (arXiv:2607.20064) で 97.4% - alexisfox7/PRO-LONG

記事本文:
GitHub - alexisfox7/PRO-LONG: ロングホライズン LLM エージェント用のプログラム メモリ: ハーネスはすべてを 1 つのログに追加し、エージェントはコードでそれを検索します。 ARC-AGI-3 (arXiv:2607.20064) で 97.4% · GitHub
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
アレクシスフォックス7
/
プロロング
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
42 コミット 42 コミット アセット アセット docker docker prolong_agent prolong_agent release_logs/ fable5 releases

e_logs/ fable5 スコアカード スコアカード .gitignore .gitignore .python-version .python-version CITATION.cff CITATION.cff ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイル ナビゲーション
PRO-LONG: プログラムによる記憶により長期的な推論が可能に
PRO-LONG は、長期タスクを実行する LLM エージェント向けの最小限のメモリ追加です。ハーネスは、すべての観察、アクション、結果を単一の構造化された log.txt に追加し、エージェントはプログラム (grep、Python) でそれを取得して推論します。サブエージェントや特殊な取得メカニズムはなく、システム プロンプトは約 30 行です。
完全な ARC-AGI-3 パブリック ゲーム セットでは、PRO-LONG は、ログなしの同じコーディング エージェントよりも平均 18 パーセンテージ ポイント改善し、請求トークンの数が 4.2 ～ 5.8 分の 1 で特殊なハーネスと同等またはそれを上回り、合計コスト 1,750 ドルで Fable 5 で 97.4% のベスト@2 に達します。
論文: arxiv.org/abs/2607.20064
Python (3.12 推奨) と Docker が必要です。
git clone git@github.com:alexisfox7/PRO-LONG.git
cd プロロング
Python -m venv .venv
ソース .venv/bin/activate
pip install -e 。
# コーデックス バックエンド
docker build -t prolong-agent/codex-sandbox:最新の docker/codex-sandbox
docker build -t prolong-openai-proxy docker/openai-proxy
# クロードコードバックエンド
docker build -t prolong-agent/claude-sandbox:最新の docker/claude-sandbox
docker build -t prolong-anthropic-proxy docker/anthropic-proxy
.env ファイルを作成します。
ARC_API_KEY=...
CODEX_API_KEY=... # コーデックス バックエンド
CLAUDE_CODE_OAUTH_TOKEN=... # クロードコード バックエンド (デフォルト)
ANTHROPIC_API_KEY=... # --api-key を使用したクロードコード バックエンド
エージェント コンテナーはゲーム ワークスペースのみをマウントし、デフォルトではモデル API へのプロキシ以外のネットワーク アクセスを持ちません。
prolong-swarm --suite all -m gpt-5.5 --max-actions 500

prolong-swarm --suite all --backend claude-code -m claude-opus-4-6
prolong-swarm --game ls20,ft09 -m gpt-5.5
結果は、evaluation_results/ に書き込まれます。
旗
デフォルト
説明
--バックエンド
コーデックス
codex (OpenAI Codex CLI) または claude-code (Claude Code CLI)
--スイート
—
ゲームスイート (すべてのみ)
--ゲーム
—
カンマで区切られた個々のゲーム名または完全な ID
--最大アクション
500
ゲームごとの最大アクション
--モデル 、 -m
バックエンド固有
コーデックスの場合は gpt-5.5。クロード・コードのclaude-opus-4-6
--努力
高い
努力レベル (クロードコード バックエンド)
--推論-努力
なし
推論作業 (コーデックス バックエンド)
--動作モード
オンライン
オンライン/オフライン/通常
メモリの状態
エージェントのゲーム履歴へのアクセスは、 --log-window および --workspace によって制御されます。これらは論文からのアブレーション条件です。
scorecards/ には、論文 ( fable_online_scorecards.txt ) に記載されている 25 の Fable 5 実行すべてを含む、公式オンライン スコアカードが含まれています。それぞれは arcprize.org で確認できます。 release_logs/ には、Fable 5 のオンライン実行のログ (ゲーム ログ、エージェントのトランスクリプト、ワークスペース) が含まれています。残りのアブレーションのログが追加されます。
prolong_agent/
§── エージェント/
│ §──base.py # 基本アーキテクチャ
│ §── codex_agent.py # Codex CLI バックエンド
│ §── claude_code_agent.py # クロードコードバックエンド
│ §── swarm.py # CLI エントリポイント
│ §── action_queue.py # アクションの実行
│ §── game_state.py # ボード/ログのフォーマット
│ lux──prompts.py # プロンプト (~30 行)
§── 環境/
│ §── arcagi3.py # ARC-AGI-3 API ラッパー
│ §──runner.py # ゲームごとのループ
│ └── config.py
§── メトリクス/
└── ユーティリティ/
このリポジトリは、以前は Read-Grep-Bash (RGB) Agent でした。ARC-AGI-3 プレビュー ゲームに関する元のブログ投稿を参照してください。
@misc { fox2026prolong 、
title = { PRO-LONG: プログレ

ラマティックな記憶により長期的な推論が可能になる } ,
著者 = { フォックス、アレクシスとワン、ジュンリンとロス、ポールとディングラ、ブワン } 、
年 = { 2026 } 、
eprint = { 2607.20064 } 、
archivePrefix = { arXiv } 、
プライマリクラス = { cs.AI } 、
URL = { https://arxiv.org/abs/2607.20064 } 、
}
について
長期的な LLM エージェント用のプログラム メモリ: ハーネスはすべてを 1 つのログに追加し、エージェントはそれをコードで検索します。 ARC-AGI-3 で 97.4% (arXiv:2607.20064)
Readme MIT ライセンス このリポジトリを引用する アクティビティのスター
40 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Programmatic memory for long-horizon LLM agents: the harness appends everything to one log, and the agent searches it with code. 97.4% on ARC-AGI-3 (arXiv:2607.20064) - alexisfox7/PRO-LONG

GitHub - alexisfox7/PRO-LONG: Programmatic memory for long-horizon LLM agents: the harness appends everything to one log, and the agent searches it with code. 97.4% on ARC-AGI-3 (arXiv:2607.20064) · GitHub
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
alexisfox7
/
PRO-LONG
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
42 Commits 42 Commits assets assets docker docker prolong_agent prolong_agent release_logs/ fable5 release_logs/ fable5 scorecards scorecards .gitignore .gitignore .python-version .python-version CITATION.cff CITATION.cff LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
PRO-LONG: Programmatic Memory Enables Long-Horizon Reasoning
PRO-LONG is a minimal memory addition for LLM agents on long-horizon tasks. The harness appends every observation, action, and outcome to a single structured log.txt, and the agent retrieves and reasons over it programmatically (grep, Python). There are no subagents or specialized retrieval mechanisms, and the system prompt is about 30 lines.
On the full ARC-AGI-3 public game set, PRO-LONG improves over the same coding agents without the log by 18 percentage points on average, matches or exceeds specialized harnesses at 4.2–5.8x fewer billed tokens, and reaches 97.4% best@2 with Fable 5 at a total cost of $1,750.
Paper: arxiv.org/abs/2607.20064
Requires Python (3.12 recommended) and Docker.
git clone git@github.com:alexisfox7/PRO-LONG.git
cd PRO-LONG
python -m venv .venv
source .venv/bin/activate
pip install -e .
# codex backend
docker build -t prolong-agent/codex-sandbox:latest docker/codex-sandbox
docker build -t prolong-openai-proxy docker/openai-proxy
# claude-code backend
docker build -t prolong-agent/claude-sandbox:latest docker/claude-sandbox
docker build -t prolong-anthropic-proxy docker/anthropic-proxy
Create a .env file:
ARC_API_KEY=...
CODEX_API_KEY=... # codex backend
CLAUDE_CODE_OAUTH_TOKEN=... # claude-code backend (default)
ANTHROPIC_API_KEY=... # claude-code backend with --api-key
The agent container only mounts the game workspace and, by default, has no network access except a proxy to the model API.
prolong-swarm --suite all -m gpt-5.5 --max-actions 500
prolong-swarm --suite all --backend claude-code -m claude-opus-4-6
prolong-swarm --game ls20,ft09 -m gpt-5.5
Results are written to evaluation_results/ .
Flag
Default
Description
--backend
codex
codex (OpenAI Codex CLI) or claude-code (Claude Code CLI)
--suite
—
Game suite ( all only)
--game
—
Comma-separated individual game names or full IDs
--max-actions
500
Max actions per game
--model , -m
Backend-specific
gpt-5.5 for Codex; claude-opus-4-6 for Claude Code
--effort
high
Effort level (claude-code backend)
--reasoning-effort
none
Reasoning effort (codex backend)
--operation-mode
online
online / offline / normal
Memory conditions
The agent's access to game history is controlled by --log-window and --workspace . These are the ablation conditions from the paper:
scorecards/ contains the official online scorecards, including all 25 Fable 5 runs from the paper ( fable_online_scorecards.txt ); each can be verified on arcprize.org. release_logs/ contains logs for the Fable 5 online runs: game logs, agent transcripts, and workspaces. Logs for the remaining ablations will be added.
prolong_agent/
├── agent/
│ ├── base.py # base architecture
│ ├── codex_agent.py # Codex CLI backend
│ ├── claude_code_agent.py # Claude Code backend
│ ├── swarm.py # CLI entry point
│ ├── action_queue.py # action execution
│ ├── game_state.py # board/log formatting
│ └── prompts.py # prompts (~30 lines)
├── environment/
│ ├── arcagi3.py # ARC-AGI-3 API wrapper
│ ├── runner.py # per-game loop
│ └── config.py
├── metrics/
└── utils/
This repo was formerly the Read-Grep-Bash (RGB) Agent, see our original blog post on the ARC-AGI-3 preview games.
@misc { fox2026prolong ,
title = { PRO-LONG: Programmatic Memory Enables Long-Horizon Reasoning } ,
author = { Fox, Alexis and Wang, Junlin and Rosu, Paul and Dhingra, Bhuwan } ,
year = { 2026 } ,
eprint = { 2607.20064 } ,
archivePrefix = { arXiv } ,
primaryClass = { cs.AI } ,
url = { https://arxiv.org/abs/2607.20064 } ,
}
About
Programmatic memory for long-horizon LLM agents: the harness appends everything to one log, and the agent searches it with code. 97.4% on ARC-AGI-3 (arXiv:2607.20064)
Readme MIT license Cite this repository Activity Stars
40 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
