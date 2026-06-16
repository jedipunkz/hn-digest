---
source: "https://pypi.org/project/subagent-fleet/"
hn_url: "https://news.ycombinator.com/item?id=48547977"
title: "Show HN: Subagent-fleet – AI coding subagents across local Ollama machines"
article_title: "subagent-fleet · PyPI"
author: "akarnam37"
captured_at: "2026-06-16T01:09:30Z"
capture_tool: "hn-digest"
hn_id: 48547977
score: 2
comments: 0
posted_at: "2026-06-15T22:37:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Subagent-fleet – AI coding subagents across local Ollama machines

- HN: [48547977](https://news.ycombinator.com/item?id=48547977)
- Source: [pypi.org](https://pypi.org/project/subagent-fleet/)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T22:37:15Z

## Translation

タイトル: HN を表示: サブエージェント フリート – ローカル Ollama マシン全体でサブエージェントをコーディングする AI
記事のタイトル: サブエージェントフリート · PyPI
説明: ローカル モデル フリート全体で Claude Code スタイルのサブエージェントを実行します。

記事本文:
メインコンテンツにスキップ
モバイル版に切り替える
警告
サポートされていないブラウザを使用しているため、新しいバージョンにアップグレードしてください。
PyPIを検索
検索フォーカス#フォーカス検索フィールド"
data-search-focus-target="検索フィールド">
検索
ヘルプ
pip install サブエージェントフリート
PIP 命令のコピー
ローカル モデル フリート全体で Claude Code スタイルのサブエージェントを実行します。
ライセンス表記: MIT
SPDX
ライセンスの表現
ローカル モデル フリート全体で Claude Code スタイルのサブエージェントを実行します。
subagent-fleet は、コーディング サブエージェントを所有する最適な Ollama モデルとマシンにマッピングし、LiteLLM および Claude Code スタイルのエージェント構成を生成するための、config-first Python CLI です。
クイックスタート • 構成 • 生成されたファイル • セキュリティ • ロードマップ
ローカル モデル ユーザーは、ラップトップ、Mac mini、ワークステーション、ホーム サーバー、予備の GPU ボックスなど、複数の便利なマシンを所有していることがよくあります。ほとんどのコーディング ハーネスは依然として 1 つのモデル エンドポイントを指します。
subagent-fleet は、そのセットアップをプライベート ローカル サブエージェント フリートに変えます。
プランナー -> 軽量ノード上の小型高速モデル
実装者 -> より大きなノード上のより大きなコーディング モデル
レビュー担当者 -> より大きなノード上のより大きなコーディング モデル
サマライザ -> コントローラ上の小さなローカル モデル
Ollama、LiteLLM、または Claude Code を置き換えるものではありません。それらの間に接着剤が生成されます。
クロードコード/コーディングハーネス
|
v
サブエージェントフリートによって生成された LiteLLM ゲートウェイ
|
+-- Ollama ノード: ラップトップ
+-- Ollama ノード: Mac mini 64GB
+-- Ollama ノード: ワークステーション
特長
宣言的な フリート.yaml を検証します。
/api/tags を介して構成された Ollama ノードからモデルを検出します。
ollama_chat/routes を使用して litellm_config.yaml を生成します。
クロード コード スタイルの .claude/agents/*.md ファイルを生成します。
Claude Code/LiteLLM 環境変数の .env.subagent-fleet を生成します。
keep_alive を使用してウォーム構成された Ollama モデル。
ノードの健全性とエージェントのルーティング テーブルを表示する

エス。
1 台のオフライン マシンがワークフロー全体をクラッシュさせないように、到達不能なノードを隔離しておいてください。
サブエージェントフリートの初期化
サブエージェントフリートの検証
サブエージェントフリートの検出
サブエージェントフリートの生成
サブエージェントフリートのウォームアップ
サブエージェントフリートのステータス
サブエージェントフリートドクター
サブエージェントフリートクリーン
サブエージェントフリートのスキルリスト
サブエージェントフリートスキルのインストール
サブエージェントフリートプラグインのインストール
インストール
以下のインストール パスのいずれかを選択します。
CLI を PyPI から直接インストールします。
python -m pip install サブエージェントフリート
または、 pipx を使用して独立したコマンドとしてインストールします。
pipx インストール サブエージェント フリート
確認:
サブエージェントフリート --help
開発チェックアウト
プロジェクトに貢献するときにこれを使用します。
git クローン https://github.com/adityak74/subagent-fleet.git
cd サブエージェント フリート
python -m pip install -e ".[dev]"
テストを実行します。
Python -m pytest
最初のクロード コード プラグイン
最初に Claude Code からプラグインをインストールし、次にバンドルされたブートストラップ スキルで CLI をインストールします。
/plugin マーケットプレイスを追加 https://github.com/adityak74/subagent-fleet
/plugin install サブエージェントフリート
インストール後、Claude Code に次のように尋ねます。
サブエージェント フリート ブートストラップ スキルを使用して CLI をインストールし、このリポジトリを設定します。
ブートストラップ スキルは以下を実行または推奨します。
python -m pip install サブエージェントフリート
サブエージェントフリートスキルのインストール
コーデックスプラグインファースト
このリポジトリをローカル Codex マーケットプレイスとしてインストールします。
コーデックス プラグイン マーケットプレイスを追加します。
コーデックス プラグインの追加 subagent-fleet@subagent-fleet
次にコーデックスに質問してください:
サブエージェント フリート ブートストラップ スキルを使用して CLI をインストールし、このリポジトリを設定します。
クイックスタート
サブエージェントフリートの初期化
Ollama ノードのエンドポイントとモデル名を使用して フリート.yaml を編集し、それを検証します。
サブエージェントフリートの検証
どのノードが到達可能であるかを確認します。
サブエージェントフリートの検出
LiteLLM、Claude エージェント、および環境ファイルを生成します。
サブエージェントフリートの生成
LiteLLM を開始します。
エクスポート LITELLM_MASTER_KEY = "sk

-ローカル開発」
リテルム \
--config ./litellm_config.yaml \
--ホスト 127 .0.0.1 \
--ポート 4000
ローカル ゲートウェイでクロード コードを指定します。
ソース .env.subagent-fleet
クロード
構成
subagent-fleet は フリート.yaml によって駆動されます。
プロジェクト:
名前 : ローカル開発
ゲートウェイ:
プロバイダー：litellm
ホスト : 127.0.0.1
ポート: 4000
master_key_env : LITELLM_MASTER_KEY
ノード:
m5ローカル:
エンドポイント: http://localhost:11434
タグ : [ コントローラー 、 ローカル 、 高速 ]
m4-mini-64gb:
エンドポイント: http://192.168.1.50:11434
タグ : [ ヘビー 、 プログラマー 、 レビュアー ]
m4-mini-16gb :
エンドポイント: http://192.168.1.51:11434
タグ : [ スモール 、 プランナー 、 サマライザー ]
モデル:
ヘビーコーダー:
ノード: m4-mini-64gb
ollama_model : qwen2.5-coder:32b
litellm_alias : クロード・ソネット・ローカル
コンテキスト : 32768
タイムアウト: 600
最大並列 : 1
小規模コーダー:
ノード: m4-mini-16gb
ollama_model : qwen2.5-coder:7b
litellm_alias : クロード-俳句-ローカル
コンテキスト : 8192
タイムアウト: 300
最大並列 : 1
エージェント:
プランナー：
モデル：スモールコーダー
説明 : 計画、ファイルの検出、タスクの分解、および要約に使用します。
ツール: [読み取り、Grep、Glob]
プロンプト: |
あなたは、迅速な現地計画エージェントです。
ファイルを編集しないでください。
次のような簡潔な応答を返します。
- 計画
- 関連ファイル
- リスク
- 次に推奨されるエージェント
実装者:
モデル：ヘビーコーダー
説明 : 実装、バグ修正、リファクタリング、パッチ作成に使用します。
ツール: [読み取り、Grep、Glob、編集、マルチエディット、Bash]
査読者：
モデル：ヘビーコーダー
description : 実装後に差分、テスト、回帰、保守性を確認するために使用します。
ツール: [読み取り、Grep、Glob、Bash]
生成されたファイル
サブエージェントフリートの生成
作成します:
litellm_config.yaml
.claude/agents/planner.md
.claude/agents/implementer.md
.claude/agents/reviewer.md
.env.サブエージェントフリート
LiteLLM ルートの例:
モデルリスト:
- モデル名 : クロード-ソネット-

地元の
Litellm_params :
モデル: ollama_chat/qwen2.5-coder:32b
api_base : http://192.168.1.50:11434
api_key : オラマ
タイムアウト: 600
モデル情報:
max_input_tokens : 32768
クロード エージェントの例:
---
名前：プランナー
説明: 計画、ファイル検出、タスクの分解、および要約に使用します。
モデル: クロード-俳句-ローカル
ツール: 読み取り、Grep、Glob
---
あなたは、迅速な現地計画エージェントです。
ファイルを編集しないでください。
次のような簡潔な応答を返します。
- 計画
- 関連ファイル
- リスク
- 次に推奨されるエージェント
コマンド
JSON 出力は検出とステータスに使用できます。
サブエージェントフリートの発見 --json
サブエージェントフリートのステータス --json
アシスタントスキル
subagent-fleet は、Claude Code、Codex、OpenCode、および同様のツールに、リポジトリ内からフリートをセットアップして操作する方法を教えるアシスタント向けスキルを提供します。
バンドルされたスキルとサポートされているターゲットをリストします。
サブエージェントフリートのスキルリスト
サポートされているすべてのターゲットにバンドルされているすべてのスキルをインストールします。
サブエージェントフリートスキルのインストール
これは次のように書きます:
.claude/skills/subagent-fleet-setup/SKILL.md
.claude/skills/subagent-fleet-operations/SKILL.md
.codex/skills/subagent-fleet-setup/SKILL.md
.codex/skills/subagent-fleet-operations/SKILL.md
.opencode/skills/subagent-fleet-setup/SKILL.md
.opencode/skills/subagent-fleet-operations/SKILL.md
特定のアシスタント用にインストールします。
サブエージェントフリートスキルのインストール --target codex
サブエージェントフリートスキルのインストール --target claude-code
サブエージェントフリートのスキル install --target opencode
バンドルされたスキルを 1 つインストールします。
サブエージェント フリート スキル インストール --skill サブエージェント フリート セットアップ
--force を渡さない限り、既存のスキル ファイルは上書きされません。
このリポジトリにはプラグイン マーケットプレイス メタデータも同梱されているため、ユーザーは最初にアシスタント スキルをインストールし、次にそのスキルで Python CLI をインストールして検証できます。
.claude-plugin/marketplace.json
.agents/plugins/marketplace.js

に
plugins/subagent-fleet/.claude-plugin/plugin.json
plugins/subagent-fleet/.codex-plugin/plugin.json
plugins/subagent-fleet/skills/subagent-fleet-bootstrap/SKILL.md
plugins/subagent-fleet/skills/subagent-fleet-setup/SKILL.md
plugins/subagent-fleet/skills/subagent-fleet-operations/SKILL.md
ブートストラップ スキルは、Claude Code または Codex に CLI のインストール方法を教えます。
python -m pip install サブエージェントフリート
次に、リポジトリ ローカル アシスタント スキルをインストールします。
サブエージェントフリートスキルのインストール
Claude Code プラグインのインストール フロー:
/plugin マーケットプレイスを追加 https://github.com/adityak74/subagent-fleet
/plugin install サブエージェントフリート
コーデックス ローカル マーケットプレイスのフロー:
コーデックス プラグイン マーケットプレイスを追加します。
コーデックス プラグインの追加 subagent-fleet@subagent-fleet
同じマーケットプレイス/プラグイン バンドルを別のディレクトリに生成するには:
サブエージェントフリートプラグインのインストール --out /path/to/marketplace-root
ターゲットを 1 つだけインストールします。
サブエージェントフリートプラグインのインストール --target claude-code
サブエージェントフリートプラグインのインストール --target codex
--force を指定しない限り、既存のプラグイン マーケットプレイス ファイルは上書きされません。
各ワーカー マシンで、コントローラーからアクセス可能なプライベート インターフェイスで Ollama を実行します。
launchctl setenv OLLAMA_HOST "0.0.0.0:11434"
launchctl setenv OLLAMA_KEEP_ALIVE "-1"
launchctl setenv OLLAMA_NUM_PARALLEL "1"
launchctl setenv OLLAMA_MAX_LOADED_MODELS "1"
キオール・オラマ
開く -a オラマ
コントローラーから:
カール http://NODE_IP:11434/api/tags
セキュリティ
サブエージェント フリートはプライベート ローカル ネットワークを想定しています。
LAN、ファイアウォール ルール、Tailscale、WireGuard、またはプライベート サブネットを使用します。
LiteLLM アクセス用に LITELLM_MASTER_KEY を設定したままにします。
生成された .env.subagent-fleet ファイルをローカル開発者構成として扱います。
オラマを公共のインターネットに直接公開します。
認証なしで LiteLLM を公開します。
実際の API キー、LAN シークレット、またはマシン固有のプライベートをコミットする

.env ファイル。
サブエージェントフリートドクター
ローカル設定と安全リマインダー用。
python -m pip install -e ".[dev]"
テストを実行します。
Python -m pytest
焦点を絞ったテストを実行します。
python -m pytest テスト/test_config.py
CLI の配線を確認します。
python -m subagent_fleet.cli --help
プロジェクトのレイアウト
src/サブエージェントフリート/
cli.py
config.py
Discovery.py
プラグイン.py
ウォームアップ.py
status.py
スキル.py
発電機/
スキル_テンプレート/
テンプレート/
例/
プラグイン/
テスト/
ロードマップ
/api/tags を介した Ollama モデルの検出
推奨されるエージェントからノードへの割り当て
テールスケール対応のノード検出
OpenAI互換ハーネス例
vLLM、LM Studio、llama.cpp、OpenRouter、およびクラウド API のサポート
問題やプルリクエストは大歓迎です。
より堅牢な Ollama エラー レポート
実際の複数マシンのセットアップに関するドキュメント
Python -m pytest
これは何ではないのか
パブリックモデルホスティングプラットフォーム
これは、プライベート ローカル サブエージェント オーケストレーションのための小さなワークフロー層です。
ライセンス表記: MIT
SPDX
ライセンスの表現
発売履歴
リリースのお知らせ |
RSSフィード
ご使用のプラットフォーム用のファイルをダウンロードします。どれを選択すればよいかわからない場合は、パッケージのインストールについて詳しくご覧ください。
名前、インタープリター、ABI、プラットフォームでファイルをフィルターします。
ファイル名の形式がわからない場合は、 ホイール ファイル名 について学習してください。
現在のフィルターへの直接リンクをコピーします。
コピー
アップロードされました
2026 年 6 月 15 日
パイソン3
ファイル subagent_fleet-0.0.1.tar.gz の詳細。
ダウンロード URL: subagent_fleet-0.0.1.tar.gz
信頼できる出版を使用してアップロードしますか?いいえ
アップロード経由: uv/0.11.8 {"インストーラー":{"名前":"uv","バージョン":"0.11.8","サブコマンド":["公開"]},"python":null,"実装":{"名前":null,"バージョン":null},"ディストリビューション":{"名前":"macOS","vers ion":null、"id":null、"libc":null}、"system":{"name":null、"release":null}、"cpu":null、"openssl_version":null、"setuptools_version":null、"rustc_version":null、"ci":null}
もっと見る

ハッシュの使用の詳細については、こちらをご覧ください。
ファイル subagent_fleet-0.0.1-py3-none-any.whl の詳細。
ダウンロード URL: subagent_fleet-0.0.1-py3-none-any.whl
信頼できる出版を使用してアップロードしますか?いいえ
アップロード経由: uv/0.11.8 {"インストーラー":{"名前":"uv","バージョン":"0.11.8","サブコマンド":["公開"]},"python":null,"実装":{"名前":null,"バージョン":null},"ディストリビューション":{"名前":"macOS","vers ion":null、"id":null、"libc":null}、"system":{"name":null、"release":null}、"cpu":null、"openssl_version":null、"setuptools_version":null、"rustc_version":null、"ci":null}
ハッシュの使用の詳細については、こちらをご覧ください。
ステータス:
すべてのシステムが稼働中
Python コミュニティによって Python コミュニティのために開発および保守されています。
今すぐ寄付してください！
「PyPI」、「Python Package Index」、およびブロックのロゴは、Python Software Foundation の登録商標です。
© 2026 Python ソフトウェア財団
サイトマップ

## Original Extract

Run Claude Code-style subagents across your local model fleet.

Skip to main content
Switch to mobile version
Warning
You are using an unsupported browser, upgrade to a newer version.
Search PyPI
search-focus#focusSearchField"
data-search-focus-target="searchField">
Search
Help
pip install subagent-fleet
Copy PIP instructions
Run Claude Code-style subagents across your local model fleet.
License Expression: MIT
SPDX
License Expression
Run Claude Code-style subagents across your local model fleet.
subagent-fleet is a config-first Python CLI for mapping coding subagents to the best Ollama model and machine you own, then generating LiteLLM and Claude Code-style agent configuration.
Quickstart • Configuration • Generated Files • Security • Roadmap
Local model users often have more than one useful machine: a laptop, a Mac mini, a workstation, a home server, or a spare GPU box. Most coding harnesses still point at one model endpoint.
subagent-fleet turns that setup into a private local subagent fleet:
planner -> small fast model on a lightweight node
implementer -> larger coding model on a bigger node
reviewer -> larger coding model on a bigger node
summarizer -> small local model on the controller
It does not replace Ollama, LiteLLM, or Claude Code. It generates the glue between them:
Claude Code / coding harness
|
v
LiteLLM gateway generated by subagent-fleet
|
+-- Ollama node: laptop
+-- Ollama node: Mac mini 64GB
+-- Ollama node: workstation
Features
Validate a declarative fleet.yaml .
Discover models from configured Ollama nodes via /api/tags .
Generate litellm_config.yaml with ollama_chat/ routes.
Generate Claude Code-style .claude/agents/*.md files.
Generate .env.subagent-fleet for Claude Code/LiteLLM environment variables.
Warm configured Ollama models with keep_alive .
Show node health and agent routing tables.
Keep unreachable nodes isolated so one offline machine does not crash the whole workflow.
subagent-fleet init
subagent-fleet validate
subagent-fleet discover
subagent-fleet generate
subagent-fleet warmup
subagent-fleet status
subagent-fleet doctor
subagent-fleet clean
subagent-fleet skills list
subagent-fleet skills install
subagent-fleet plugins install
Install
Choose one of the install paths below.
Install the CLI directly from PyPI:
python -m pip install subagent-fleet
Or install it as an isolated command with pipx :
pipx install subagent-fleet
Verify:
subagent-fleet --help
Development Checkout
Use this when contributing to the project:
git clone https://github.com/adityak74/subagent-fleet.git
cd subagent-fleet
python -m pip install -e ".[dev]"
Run tests:
python -m pytest
Claude Code Plugin First
Install the plugin first from Claude Code, then let the bundled bootstrap skill install the CLI:
/plugin marketplace add https://github.com/adityak74/subagent-fleet
/plugin install subagent-fleet
After install, ask Claude Code:
Use the subagent-fleet bootstrap skill to install the CLI and set up this repo.
The bootstrap skill will run or recommend:
python -m pip install subagent-fleet
subagent-fleet skills install
Codex Plugin First
Install this repository as a local Codex marketplace:
codex plugin marketplace add .
codex plugin add subagent-fleet@subagent-fleet
Then ask Codex:
Use the subagent-fleet bootstrap skill to install the CLI and set up this repo.
Quickstart
subagent-fleet init
Edit fleet.yaml with your Ollama node endpoints and model names, then validate it:
subagent-fleet validate
Check which nodes are reachable:
subagent-fleet discover
Generate LiteLLM, Claude agent, and environment files:
subagent-fleet generate
Start LiteLLM:
export LITELLM_MASTER_KEY = "sk-local-dev"
litellm \
--config ./litellm_config.yaml \
--host 127 .0.0.1 \
--port 4000
Point Claude Code at the local gateway:
source .env.subagent-fleet
claude
Configuration
subagent-fleet is driven by fleet.yaml .
project :
name : local-dev
gateway :
provider : litellm
host : 127.0.0.1
port : 4000
master_key_env : LITELLM_MASTER_KEY
nodes :
m5-local :
endpoint : http://localhost:11434
tags : [ controller , local , fast ]
m4-mini-64gb :
endpoint : http://192.168.1.50:11434
tags : [ heavy , coder , reviewer ]
m4-mini-16gb :
endpoint : http://192.168.1.51:11434
tags : [ small , planner , summarizer ]
models :
heavy-coder :
node : m4-mini-64gb
ollama_model : qwen2.5-coder:32b
litellm_alias : claude-sonnet-local
context : 32768
timeout : 600
max_parallel : 1
small-coder :
node : m4-mini-16gb
ollama_model : qwen2.5-coder:7b
litellm_alias : claude-haiku-local
context : 8192
timeout : 300
max_parallel : 1
agents :
planner :
model : small-coder
description : Use for planning, file discovery, task decomposition, and summarization.
tools : [ Read , Grep , Glob ]
prompt : |
You are a fast local planning agent.
Do not edit files.
Return a concise response with:
- plan
- relevant files
- risks
- next recommended agent
implementer :
model : heavy-coder
description : Use for implementation, bug fixes, refactors, and patch creation.
tools : [ Read , Grep , Glob , Edit , MultiEdit , Bash ]
reviewer :
model : heavy-coder
description : Use after implementation to review diffs, tests, regressions, and maintainability.
tools : [ Read , Grep , Glob , Bash ]
Generated Files
subagent-fleet generate
creates:
litellm_config.yaml
.claude/agents/planner.md
.claude/agents/implementer.md
.claude/agents/reviewer.md
.env.subagent-fleet
Example LiteLLM route:
model_list :
- model_name : claude-sonnet-local
litellm_params :
model : ollama_chat/qwen2.5-coder:32b
api_base : http://192.168.1.50:11434
api_key : ollama
timeout : 600
model_info :
max_input_tokens : 32768
Example Claude agent:
---
name: planner
description: Use for planning, file discovery, task decomposition, and summarization.
model: claude-haiku-local
tools: Read, Grep, Glob
---
You are a fast local planning agent.
Do not edit files.
Return a concise response with:
- plan
- relevant files
- risks
- next recommended agent
Commands
JSON output is available for discovery and status:
subagent-fleet discover --json
subagent-fleet status --json
Assistant Skills
subagent-fleet ships assistant-facing skills that teach Claude Code, Codex, OpenCode, and similar tools how to set up and operate the fleet from inside a repository.
List bundled skills and supported targets:
subagent-fleet skills list
Install all bundled skills for all supported targets:
subagent-fleet skills install
This writes:
.claude/skills/subagent-fleet-setup/SKILL.md
.claude/skills/subagent-fleet-operations/SKILL.md
.codex/skills/subagent-fleet-setup/SKILL.md
.codex/skills/subagent-fleet-operations/SKILL.md
.opencode/skills/subagent-fleet-setup/SKILL.md
.opencode/skills/subagent-fleet-operations/SKILL.md
Install for a specific assistant:
subagent-fleet skills install --target codex
subagent-fleet skills install --target claude-code
subagent-fleet skills install --target opencode
Install one bundled skill:
subagent-fleet skills install --skill subagent-fleet-setup
Existing skill files are not overwritten unless you pass --force .
This repository also ships plugin marketplace metadata so users can install the assistant skill first, then let that skill install and verify the Python CLI.
.claude-plugin/marketplace.json
.agents/plugins/marketplace.json
plugins/subagent-fleet/.claude-plugin/plugin.json
plugins/subagent-fleet/.codex-plugin/plugin.json
plugins/subagent-fleet/skills/subagent-fleet-bootstrap/SKILL.md
plugins/subagent-fleet/skills/subagent-fleet-setup/SKILL.md
plugins/subagent-fleet/skills/subagent-fleet-operations/SKILL.md
The bootstrap skill teaches Claude Code or Codex how to install the CLI:
python -m pip install subagent-fleet
and then install repo-local assistant skills:
subagent-fleet skills install
Claude Code plugin install flow:
/plugin marketplace add https://github.com/adityak74/subagent-fleet
/plugin install subagent-fleet
Codex local marketplace flow:
codex plugin marketplace add .
codex plugin add subagent-fleet@subagent-fleet
To generate the same marketplace/plugin bundle into another directory:
subagent-fleet plugins install --out /path/to/marketplace-root
Install only one target:
subagent-fleet plugins install --target claude-code
subagent-fleet plugins install --target codex
Existing plugin marketplace files are not overwritten unless you pass --force .
On each worker machine, run Ollama on a private interface reachable from your controller:
launchctl setenv OLLAMA_HOST "0.0.0.0:11434"
launchctl setenv OLLAMA_KEEP_ALIVE "-1"
launchctl setenv OLLAMA_NUM_PARALLEL "1"
launchctl setenv OLLAMA_MAX_LOADED_MODELS "1"
killall Ollama
open -a Ollama
From the controller:
curl http://NODE_IP:11434/api/tags
Security
subagent-fleet assumes private local networking.
Use LAN, firewall rules, Tailscale, WireGuard, or a private subnet.
Keep LITELLM_MASTER_KEY set for LiteLLM access.
Treat generated .env.subagent-fleet files as local developer configuration.
Expose Ollama directly to the public internet.
Expose LiteLLM without authentication.
Commit real API keys, LAN secrets, or machine-specific private .env files.
subagent-fleet doctor
for local setup and safety reminders.
python -m pip install -e ".[dev]"
Run tests:
python -m pytest
Run a focused test:
python -m pytest tests/test_config.py
Check CLI wiring:
python -m subagent_fleet.cli --help
Project Layout
src/subagent_fleet/
cli.py
config.py
discovery.py
plugins.py
warmup.py
status.py
skills.py
generators/
skill_templates/
templates/
examples/
plugins/
tests/
Roadmap
Ollama model discovery via /api/tags
Recommended agent-to-node assignment
Tailscale-aware node discovery
OpenAI-compatible harness examples
Support for vLLM, LM Studio, llama.cpp, OpenRouter, and cloud APIs
Issues and pull requests are welcome.
More robust Ollama error reporting
Documentation for real multi-machine setups
python -m pytest
What This Is Not
a public model hosting platform
It is a small workflow layer for private local subagent orchestration.
License Expression: MIT
SPDX
License Expression
Release history
Release notifications |
RSS feed
Download the file for your platform. If you're not sure which to choose, learn more about installing packages .
Filter files by name, interpreter, ABI, and platform.
If you're not sure about the file name format, learn more about wheel file names .
Copy a direct link to the current filters
Copy
Uploaded
Jun 15, 2026
Python 3
Details for the file subagent_fleet-0.0.1.tar.gz .
Download URL: subagent_fleet-0.0.1.tar.gz
Uploaded using Trusted Publishing? No
Uploaded via: uv/0.11.8 {"installer":{"name":"uv","version":"0.11.8","subcommand":["publish"]},"python":null,"implementation":{"name":null,"version":null},"distro":{"name":"macOS","version":null,"id":null,"libc":null},"system":{"name":null,"release":null},"cpu":null,"openssl_version":null,"setuptools_version":null,"rustc_version":null,"ci":null}
See more details on using hashes here.
Details for the file subagent_fleet-0.0.1-py3-none-any.whl .
Download URL: subagent_fleet-0.0.1-py3-none-any.whl
Uploaded using Trusted Publishing? No
Uploaded via: uv/0.11.8 {"installer":{"name":"uv","version":"0.11.8","subcommand":["publish"]},"python":null,"implementation":{"name":null,"version":null},"distro":{"name":"macOS","version":null,"id":null,"libc":null},"system":{"name":null,"release":null},"cpu":null,"openssl_version":null,"setuptools_version":null,"rustc_version":null,"ci":null}
See more details on using hashes here.
Status:
all systems operational
Developed and maintained by the Python community, for the Python community.
Donate today!
"PyPI", "Python Package Index", and the blocks logos are registered trademarks of the Python Software Foundation .
© 2026 Python Software Foundation
Site map
