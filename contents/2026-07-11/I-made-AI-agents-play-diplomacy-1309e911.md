---
source: "https://github.com/brendenehlers/diplomacy-ai"
hn_url: "https://news.ycombinator.com/item?id=48872843"
title: "I made AI agents play diplomacy"
article_title: "GitHub - brendenehlers/diplomacy-ai: ever wondered what a game of diplomacy played by ai agents would look like? · GitHub"
author: "behlers99"
captured_at: "2026-07-11T15:45:16Z"
capture_tool: "hn-digest"
hn_id: 48872843
score: 1
comments: 0
posted_at: "2026-07-11T15:24:54Z"
tags:
  - hacker-news
  - translated
---

# I made AI agents play diplomacy

- HN: [48872843](https://news.ycombinator.com/item?id=48872843)
- Source: [github.com](https://github.com/brendenehlers/diplomacy-ai)
- Score: 1
- Comments: 0
- Posted: 2026-07-11T15:24:54Z

## Translation

タイトル: AIエージェントに外交をさせてみた
記事のタイトル: GitHub - brendenehlers/diplomacy-ai: AI エージェントが行う外交ゲームがどのようなものになるか考えたことはありますか? · GitHub
説明: AI エージェントが行う外交ゲームがどのようなものになるか考えたことはありますか? - ブレンデネラーズ/外交-ai

記事本文:
GitHub - brendenehlers/diplomacy-ai: AI エージェントが行う外交ゲームがどのようなものになるか考えたことはありますか? · GitHub
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
ブレンデネラーズ
/
外交ai
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビ

イゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
21 コミット 21 コミット外交_ai 外交_ai docs/ superpowers docs/ superpowers testing testing .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md README.md README.md game.toml game.toml justfile justfile pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルナビゲーション
外交ゲームを完全に実行する
公式上の 7 つの LLM 制御権限の間
外交エンジン。
各運動フェーズでは、権力側は一定回数の報道ラウンドにわたって交渉します。
(プライベート メッセージとグローバル メッセージ)、注文を送信します。あらゆる権力は私的なものである
推理、メッセージ、命令が記録されているので、ゲーム全体を振り返ることができます
その後、決して公にはならない陰謀も含めて。
各フェーズについて、エンジンは次のようにレポートします。
ネゴシエーション (移動フェーズのみ) — n_negotiation_rounds にわたって、すべて
生きている権力者は、受信箱を読み、同時に新しいメッセージを送信します。
別の電力（プライベート）またはグローバル（ブロードキャスト）に送信します。メッセージはルーティングされ、
次のラウンドの受信箱に配信されます。
注文 — 注文可能なユニットを持つ各パワーは構造化された注文を返します
JSON。注文はエンジンの法的注文リストと照合して検証されます。違法
注文により、修理の再プロンプトが 1 回トリガーされます。まだ違法なものは削除されます
(単位は保持されます)。エンジンが不正な命令を受けることはありません。
処理と記録 — エンジンがフェーズを判断し、完全な
トランスクリプトと更新された保存済みゲーム ファイルがディスクに書き込まれます。
ループはゲームが終了するか max_year に達するまで継続します。
モジュール境界は意図的に行われます - エンジンまたは LLM バックエンドタッチを交換します
ちょうど 1 つのファイルです。
Python 3.11+ (リポジトリは 3.14 で実行されます)。すべての Python は .venv/ の venv を通過します。
# または: .venv/bin/pip install -e ".[dev]" をインストールしてください
cp .env.exam

ple .env # .env を自動ロードするだけ
サンプルの game.toml は Google Gemini ( gemini/gemini-3.5-flash ) を使用します。参照
以下の Gemini をプレイしてセットアップしてください。
代わりにモデルをローカルで実行するには (API キーなし、コストなし)、次を使用します。
LM StudioのOpenAI対応サーバー：セット
game.toml の default_model を lm_studio/qwen/qwen3-4b- Thinking-2507 に追加し、開始します
LM Studio にそのモデルをロードし、LM_STUDIO_API_BASE / を設定します。
.env 内の LM_STUDIO_API_KEY 。 LiteLLM がサポートするすべての
プロバイダーも同様に機能します。モデルを変更し、一致するキーを設定します。
独自のゲームを実行する最も簡単な方法。 Gemini の無料枠で十分です
で実験します。
Google AI Studio から API キーを取得します。
それを .env に置きます ( .env.example からコピー):
GEMINI_API_KEY=あなたのキーはここにあります
.env を自動ロードするだけなので、他にエクスポートするものはありません。
game.toml でモデルを選択します。デフォルトは gemini/gemini-3.5-flash です。
（早くて安い）。より強力なプレイのために、gemini/gemini-2.5-pro に交換してください。どれでも
gemini/... id LiteLLM は機能することを知っています。別のモデルを設定することもできます
[powers.<NAME>] の下のパワーごとに、モデルを互いに比較します。
ゲーム全体の前に、最初にセットアップを確認します (実際の Gemini フェーズ 1 つ)。
完全なゲームでは、すべてのフェーズにわたって多くの呼び出しが行われます。割り当てやコストに注意してください。
n_negotiation_rounds または max_year を上げます。
# または: .venv/bin/diplomacy-ai run --config game.toml を実行するだけです
出力は run/<timestamp>/ (gitignored) にあります。
game.json — 保存されたゲーム ファイル。公式 Web UI にロードして、
ボードとプレス。
トランスクリプト/<フェーズ>.json — 送信/受信された各勢力のプライベートな推論
ラウンドごとのメッセージ、生の注文と最終注文、ドロップされた注文、およびコールのメタデータ
(トークン、コスト、レイテンシー)。
events.log — 実行中の進行状況ログ。
game.toml <dir> を実行するだけでカスタム ディレクトリに書き込みます。または
--out <ディレクトリ> 。
LiteLLM でサポートされているモデル ID はすべて機能します。例:ジェミニ/ジェミニ-3.5-フラッシュ、
ジェミニ/ジェミニ-2。

5-pro 、 openai/gpt-4o 、 anthropic/claude-... 、または
lm_studio/qwen/qwen3-4b- Thinking-2507 (ローカル)。ペルソナはフリーテキストであり、どのように形成するか
それぞれの勢力が交渉し、行動します。サンプル構成には 7 つの積極的な構成が同梱されています。
公式ウェブ UI で試合を観戦する
ディプロマシー パッケージには React Web UI がバンドルされています。終了したゲームをリプレイするには:
just ui-setup # one-time: UI を ./ui にコピーし、npm install
justserve-engine # ターミナル 1 — :8432 の WebSocket サーバー
# (最初の起動では護送隊の経路が事前計算されます。数分待ちます)
justserve-ui # ターミナル 2 — Web UI :3000 (3000 が使用されている場合はポートを渡します)
# 例: `justserve-ui 3007`)
次に、UI を開き、localhost:8432 に接続し、任意のホストを登録します。
ユーザー名/パスワード → ログイン → 「ディスクからゲームをロード」 → 選択
run/<タイムスタンプ>/game.json 。
UI にはボードとプレスのみが表示されます。プライベートな力ごとの推論が存在します
トランスクリプト/<フェーズ>.json 。
これは古い React-Scripts 3 アプリです。 serve-ui は従来の OpenSSL フラグを設定します
最新のノードでのニーズに対応します。 ui/ は gitignore されます。
# 高速スイートをテストするだけで、ネットワークなし (偽のプロバイダー)
RUN_SMOKE=1 テストスモークのみ # オプトインライブフェーズ (GEMINI_API_KEY が必要、費用がかかる)
just test-one testing/test_agent.py # 単一のファイル
高速スイートにはネットワークや LM Studio は必要ありません。次の場合を除き、煙テストはスキップされます。
RUN_SMOKE=1 と GEMINI_API_KEY の両方が設定されます。
すべてのレシピをリストするために実行するだけです。寄稿者のメモについては、AGENTS.md を参照してください。
元の仕様と実装計画については docs/superpowers/ を参照してください。
AI エージェントが行う外交ゲームがどのようなものになるか考えたことはありますか?
読み込み中にエラーが発生しました。このページをリロードしてください。
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

ever wondered what a game of diplomacy played by ai agents would look like? - brendenehlers/diplomacy-ai

GitHub - brendenehlers/diplomacy-ai: ever wondered what a game of diplomacy played by ai agents would look like? · GitHub
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
brendenehlers
/
diplomacy-ai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
21 Commits 21 Commits diplomacy_ai diplomacy_ai docs/ superpowers docs/ superpowers tests tests .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md README.md README.md game.toml game.toml justfile justfile pyproject.toml pyproject.toml View all files Repository files navigation
Runs a complete game of Diplomacy
between 7 LLM-controlled powers on the official
diplomacy engine.
Each movement phase, the powers negotiate over a fixed number of press rounds
(private and global messages), then submit orders . Every power's private
reasoning, messages, and orders are recorded so you can review the whole game
afterward — including the plotting that never becomes public.
For each phase the engine reports:
Negotiation (movement phases only) — over n_negotiation_rounds , all
living powers concurrently read their inbox and send new messages, addressed
to another power (private) or GLOBAL (broadcast). Messages are routed and
delivered to the next round's inboxes.
Orders — each power with orderable units returns orders as structured
JSON. Orders are validated against the engine's legal-order list. Illegal
orders trigger one repair re-prompt ; anything still illegal is dropped
(the unit holds). The engine never receives an illegal order.
Process & record — the engine adjudicates the phase, and the full
transcript plus the updated saved-game file are written to disk.
The loop continues until the game ends or max_year is reached.
Module boundaries are deliberate — swapping the engine or LLM backend touches
exactly one file.
Python 3.11+ (repo runs on 3.14). All Python goes through the venv at .venv/ .
just install # or: .venv/bin/pip install -e ".[dev]"
cp .env.example .env # just auto-loads .env
The sample game.toml uses Google Gemini ( gemini/gemini-3.5-flash ). See
Play with Gemini below to get set up.
To run a model locally instead (no API key, no cost), use
LM Studio 's OpenAI-compatible server: set
default_model in game.toml to lm_studio/qwen/qwen3-4b-thinking-2507 , start
LM Studio with that model loaded, and set LM_STUDIO_API_BASE /
LM_STUDIO_API_KEY in .env . Any LiteLLM-supported
provider works the same way — change the model and set the matching key.
The quickest way to run your own game. Gemini's free tier is enough to
experiment with.
Get an API key from Google AI Studio .
Put it in .env (copied from .env.example ):
GEMINI_API_KEY=your-key-here
just auto-loads .env , so nothing else to export.
Pick your model in game.toml — the default is gemini/gemini-3.5-flash
(fast and cheap). Swap in gemini/gemini-2.5-pro for stronger play. Any
gemini/... id LiteLLM knows works. You can also set a different model
per power under [powers.<NAME>] to pit models against each other.
Verify your setup first (one real Gemini phase) before a full game:
Full games make many calls across all phases — watch your quota/costs if you
raise n_negotiation_rounds or max_year .
just run # or: .venv/bin/diplomacy-ai run --config game.toml
Output lands in runs/<timestamp>/ (gitignored):
game.json — saved-game file; load it in the official web UI to watch the
board and press.
transcript/<phase>.json — each power's private reasoning, sent/received
messages per round, raw + final orders, dropped orders, and call metadata
(tokens, cost, latency).
events.log — running progress log.
Write to a custom directory with just run-out game.toml <dir> or
--out <dir> .
Any LiteLLM-supported model id works, e.g. gemini/gemini-3.5-flash ,
gemini/gemini-2.5-pro , openai/gpt-4o , anthropic/claude-... , or
lm_studio/qwen/qwen3-4b-thinking-2507 (local). Personas are free-text and shape how
each power negotiates and plays; the sample config ships 7 aggressive ones.
Watch a game in the official web UI
The diplomacy package bundles a React web UI. To replay a finished game:
just ui-setup # one-time: copy the UI into ./ui and npm install
just serve-engine # terminal 1 — websocket server on :8432
# (first launch precomputes convoy paths; wait a few minutes)
just serve-ui # terminal 2 — web UI on :3000 (pass a port if 3000 is taken,
# e.g. `just serve-ui 3007`)
Then open the UI → Connect to localhost:8432 → register any
username/password → log in → "Load a game from disk" → pick
runs/<timestamp>/game.json .
The UI shows the board and press only; private per-power reasoning lives in
transcript/<phase>.json .
It's an old react-scripts 3 app; serve-ui sets the legacy OpenSSL flag it
needs on modern Node. ui/ is gitignored.
just test # fast suite, no network (fake provider)
RUN_SMOKE=1 just test-smoke # opt-in live phase (needs GEMINI_API_KEY, costs money)
just test-one tests/test_agent.py # a single file
The fast suite needs no network or LM Studio. The smoke test is skipped unless
both RUN_SMOKE=1 and GEMINI_API_KEY are set.
Run just to list all recipes. See AGENTS.md for contributor notes and
docs/superpowers/ for the original spec and implementation plan.
ever wondered what a game of diplomacy played by ai agents would look like?
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
