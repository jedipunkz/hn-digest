---
source: "https://github.com/Naveja00/OverReach"
hn_url: "https://news.ycombinator.com/item?id=48608558"
title: "Overreach audit your AI agent's diff against the prompt you gave it"
article_title: "GitHub - Naveja00/OverReach · GitHub"
author: "Naveja"
captured_at: "2026-06-20T12:43:17Z"
capture_tool: "hn-digest"
hn_id: 48608558
score: 1
comments: 0
posted_at: "2026-06-20T11:51:53Z"
tags:
  - hacker-news
  - translated
---

# Overreach audit your AI agent's diff against the prompt you gave it

- HN: [48608558](https://news.ycombinator.com/item?id=48608558)
- Source: [github.com](https://github.com/Naveja00/OverReach)
- Score: 1
- Comments: 0
- Posted: 2026-06-20T11:51:53Z

## Translation

タイトル: AI エージェントに与えたプロンプトとの差分をオーバーリーチ監査する
記事タイトル: GitHub - Naveja00/OverReach · GitHub
説明: GitHub でアカウントを作成して、Naveja00/OverReach の開発に貢献します。

記事本文:
GitHub - Naveja00/OverReach · GitHub
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
ナベジャ00
/
オーバーリーチ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

レ
16 コミット 16 コミット .github .github docs docs src src testing testing .gitignore .gitignore CLAUDE.md CLAUDE.md HANDOFF-AGENT-AUTH.md HANDOFF-AGENT-AUTH.md ライセンス ライセンス README.md README.md TESTED-MODELS.md TESTED-MODELS.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントのスコープ クリープを捕捉するスタンドアロン MCP ツール。
コーディングエージェントに与えたプロンプトと、それが生成した差分をそれに与えます。
Overreach は、差分がプロンプトが要求した内容の範囲内に留まっていたかどうかを示します。
エージェントがエンドポイント、依存関係、環境変数、または cron ジョブを静かに追加したかどうか
あなたが決して求めなかったこと。
「私の AI アシスタントが、私抜きで製品の決定を非常に頻繁に行っていたことが判明しました。」
Node.js 18+ — nodejs.org 。ノード -v で確認します。
npm には Node.js が付属しています。 npm -v で確認します。
Git — コミット前のフックと git diff パイピングに必要です。
npx -y -p オーバーリーチ オーバーリーチ-cli デモ
サンプルの差分で実際のパイプラインを実行します。API キーやセットアップは不要で、費用はかかりません。
スコープ クリープが高いという結果で 1 を終了します (デモ プロンプトではログイン フォームが要求されます。
差分は、Stripe、環境変数、エンドポイント、および cron ジョブに密輸されます)。それが
製品全体を 1 つのコマンドで実行できます。
プロンプトで許可されていないものが追加されると、diff にフラグが立てられます。
重大度: env / エンドポイント / cron = 高 · dep / ファイル = 中 · 機能 = 低。
全体的なscope_creep_score : 高い所見がある場合は HIGH、中程度の場合は MEDIUM、それ以外の場合は LOW 。
ステージ 1 — スコープ抽出 (LLM)。プロンプトを読んで、
承認されたスコープ JSON: ファイル、機能、deps、エンドポイント、env、および動作
あなたは実際に求めました。タイプミスを最も近い実際の概念に解読しますが、決して解読しません
スコープを発明します。これはモデルを呼び出す唯一のステージです。
ステージ 2 — Diff 解析 (dete

rministic、LLMなし）。差分を正規表現で解析して、
実際に追加されるもののセット — imports、deps、process.env.X 参照、route
ハンドラー、cron ジョブ、新しいシンボル。ミリ秒単位で実行されます。
ステージ 3 — 比較 (決定論的)。ファジーマッチングを使用して算術演算を設定します。
実際の − 認可された = 所見。
ステージ 2 と 3 は純粋な関数です。推論や意見はなく、完全に監査可能です。
これが、Overreach を推論に 1 セントも費やすことなくテストできる理由です。
npm install -g オーバーリーチ
または、npx 経由で直接使用します (インストールは必要ありません)。
npx -y -p オーバーリーチ オーバーリーチ-cli デモ
APIキー（オプション）
最良の結果を得るには、ステージ 1 スコープ抽出用に 1 つの LLM プロバイダー キーを設定します。
SCOPE_PROVIDER および OVERREACH_MODEL を使用してプロバイダー/モデルを固定します。
鍵がありませんか？問題ない。 API キーがないと、Overreach は次のようにフォールバックします。
確定的なスコープ抽出 — プロンプトを正規表現で解析して具体的なものを抽出します
アイテム (ファイルパス、パッケージ名、/api/... ルート、SCREAMING_SNAKE_CASE 環境)
LLM を呼び出す代わりに、vars、cron キーワード) を使用します。曖昧では分からないだろう
LLM と同様に命令ですが、すべての具体名詞をキャッチします。
あなたのプロンプト。即時、無料、完全オフライン。
1. プロジェクトのセットアップ (1 つのコマンド)
npx -y -p オーバーリーチ オーバーリーチ-cli init
これにより、次の 3 つのことが生じます。
.overreach/prompt.md — エージェントに与えたプロンプトをここに書きます
.git/hooks/pre-commit — プロンプトに対してすべてのコミットを監査します
CLAUDE.md — コミットする前に AI エージェントに自己監査を指示します
AI エージェントに与えた実際の指示を使用して .overreach/prompt.md を編集します。
電子メール/パスワードフィールドを備えたログインフォームを設定ページに追加します。
フォーム検証、および /api/auth/login を呼び出す送信ボタン。
3. コミット — オーバーリーチは自動的に実行されます
git add 。 && git commit -m "ログインフォームを追加"
コミット前フックは、プロンプトに対して段階的な変更を監査します。
高い範囲

クリープ → コミットがブロックされました (終了 1)
MEDIUM / LOW → コミットが許可され、結果が出力されます
テンプレートプロンプト (未編集) → 正常にスキップされました
API キーなし → 決定論的フォールバック (プロンプトから具体的な項目を抽出)
何をしようとしているのかわかっている場合は、 git commit --no-verify でスキップしてください。アップデート
エージェントに新しいタスクを与えるたびに、.overreach/prompt.md が作成されます。
Windows: プリコミットフックはシェルスクリプトです。すぐに使えます
Git Bash (Git for Windows に含まれています) を使用します。
npx -y -p overreach overreach-cli --prompt " 設定ページにログイン フォームを追加します " --diff my-changes.diff
または diff をパイプします:
git diff | npx -y -p overreach overreach-cli --prompt "設定ページにログインフォームを追加"
クリーンの場合は 0、HIGH の場合は 1 を終了 — CI ゲートとして使用可能。
--prompt <テキスト> — 作業を許可する指示
--diff <パス> — 差分ファイル (デフォルト: 標準入力から読み取り)
--scope <path|json> — 承認されたスコープを挿入します。 LLM を完全にスキップします
--json — きれいな端末出力の代わりに生の JSON を出力します
--no-cache — スコープ キャッシュをバイパスします (新たなステージ 1 呼び出しを強制します)
デモ — 正規デモを実行します (ゼロキー)
init — コミット前フック + CLAUDE.md をインストールする
MCP サーバー (クロード コード、カーソル、コーデックス、クロード デスクトップ)
Overreach は標準入出力 MCP サーバーであるため、MCP 対応のクライアントはどれでも接続できます。
クロード mcp オーバーリーチを追加 -- npx -y オーバーリーチ
クロード デスクトップ / カーソル — MCP 構成に追加します。
{
"mcpサーバー": {
"オーバーリーチ" : { "コマンド" : " npx " , "args" : [ " -y " , " オーバーリーチ " ] }
}
}
Codex CLI — ~/.codex/config.toml に追加します。
[ mcp_servers .行き過ぎ】
コマンド = " npx "
args = [ " -y " , " オーバーリーチ " ]
またはストリーミング可能な HTTP: PORT=8787 を設定し、POST を http://localhost:8787/mcp に設定します。
HTTP エンドポイントには認証がありません。 127.0.0.1 (ループバック) にバインドします。
デフォルト — ローカルで使用しても安全です。公に公開しないでください
( OV

ERREACH_HOST=0.0.0.0 ) 前に認証されたリバース プロキシがない場合:
到達できます。 check_overreach を呼び出して、LLM 予算を消費できます。
公開されているツール: check_overreach(prompt, diff, options?) および health 。
初回セットアップ (クロード コード)
# 1. クロードコードでサーバーを登録する(1回)
クロード mcp オーバーリーチを追加 -- npx -y オーバーリーチ
# 2. クロード コード セッションを再開する
# (すでに開いているセッションには、終了して再度開くまで新しいサーバーが表示されません)
# 3. 必要に応じて API キーを設定します (決定論的フォールバックにより API キーがなくても機能します)
import ANTHROPIC_API_KEY=sk-... # または OPENAI_API_KEY / OLLAMA_API_KEY
再起動後は、すべての新しいセッションで check_overreach が使用可能になります (タスクごとではありません)。
セットアップ。エージェントは、関連性があると判断した場合にそれを呼び出します。
キーは自動的には渡されません。 MCP サーバーは別個のものです
プロセス。エージェントは独自の資格情報を渡しません。にログインすると、
クロード ログイン (OAuth / サブスクリプション) を使用したクロード コード、ありません
環境内の ANTHROPIC_API_KEY — したがって、1 つをエクスポートします (任意のプロバイダーが機能します。ローカル
Ollama にはキーは必要ありません)、または Claude Desktop / Cursor の場合はサーバーの env に追加します。
{ "mcpServers" : { "overreach" : { "command" : " npx " 、 "args" : [ " -y " 、 " overreach " ]、 "env" : { "ANTHROPIC_API_KEY" : " sk-... " } } } }
overreach init は、プロジェクトの CLAUDE.md にスコープ監査命令を追加します。
AI エージェントは段階的な変更をコミットする前に自己監査します - ユーザーの介入はありません
必要です。エージェントは命令を読み取り、独自の差分で Overreach を実行します。
MCP サーバー経由でエージェントに直接 check_overreach を呼び出すこともできます。
独自のタスク文字列とコミットしようとしている差分を使用します。
git diff --staged | overreach-cli --prompt "<あなたが今私に与えたタスク>"
これはベストエフォート型です。エージェントは通話をスキップしたり、調査結果を無視したりできます。
(キツネの番

鶏小屋）。ハードバックストップは以下の CI ゲートです。
ハードなバックストップ。ワークフローはすべてのプル リクエストで Overreach を実行し、失敗します
scope_creep_score=HIGH の場合の PR — diff は dep / env var / を追加します
エンドポイント / cron / 範囲外のファイル プロンプトが承認しませんでした。
.github/workflows/overreach.yml をにコピーします
リポジトリに ANTHROPIC_API_KEY (または OPENAI_API_KEY / OLLAMA_API_KEY ) を追加します。
リポジトリのシークレットとして。プロンプトはリポジトリの .overreach/prompt.md から取得されます。
ファイルが存在しない場合は PR タイトル + 本文。ジョブは結果を PR として投稿します
コメントし、 HIGH のチェックに失敗します。フルセットアップ + カスタマイズ
docs/ci-gate.md 。
# .github/workflows/overreach.yml (抜粋)
- 名前 : ランオーバーリーチ
実行: |
npx -y -p overreach@最新のoverreach-cli \
--prompt "$(cat "$RUNNER_TEMP/prompt.txt")" --diff "$RUNNER_TEMP/pr.diff"
環境:
ANTHROPIC_API_KEY : ${{ Secrets.ANTHROPIC_API_KEY }}
- name : Gate — HIGH で PR が失敗する
if :steps.overreach.outputs.exit == '1'
実行: 1 番出口
このオープンソースのアクションは無料で実行できます (独自の LLM キーを持参します)。
モデル
結果
クロード・ソネット 4.6
82/82
クロード 作品4.6
65/65
GLM 5.2
82/82
キミ K2.7 コード
82/82
ミニマックス M3
81/82
決定論的フォールバック (キーなし) は、具体的な内容を含む任意のプロンプトで機能します。
アイテム — モデルは必要ありません。
動作することを確認します (API キーがゼロ)
npmテスト
経由で挿入されたスコープを使用して、実際のパイプラインを通じて 56 個のアサーションを実行します。
scopeOverride なので、ステージ 1 (LLM) は呼び出されません。オーバーリーチをカバー
検出、クリーンパス、Python/Express/Next.js パーサー、削除処理、
決定論、チャンキング、および信頼契約の不変性。
オーバーリーチは完全に自己完結型です。他のものをインポートしたり依存したりしません
プロジェクト。独自のプロセス環境のみを読み取ります。テレメトリもコールホームもありません —
それは完全にあなたのマシン上で実行されます。
オーバーリーチが何かを見逃した場合

フラグを立てる必要がある、またはプロンプトにフラグを立てる
許可されている場合は、プロンプト + 最小の再現差分を使用して問題を開きます。
https://github.com/Naveja00/OverReach/issues
まさにそれを要求するバグレポートのテンプレートがあります。
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

Contribute to Naveja00/OverReach development by creating an account on GitHub.

GitHub - Naveja00/OverReach · GitHub
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
Naveja00
/
OverReach
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
16 Commits 16 Commits .github .github docs docs src src tests tests .gitignore .gitignore CLAUDE.md CLAUDE.md HANDOFF-AGENT-AUTH.md HANDOFF-AGENT-AUTH.md LICENSE LICENSE README.md README.md TESTED-MODELS.md TESTED-MODELS.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
A standalone MCP tool that catches AI-agent scope creep.
You give it the prompt you gave your coding agent, and the diff it produced.
Overreach tells you whether the diff stayed inside what the prompt asked for — or
whether the agent quietly added an endpoint, a dependency, an env var, or a cron job
that you never asked for.
"turns out my ai assistant had been extremely making product decisions without me"
Node.js 18+ — nodejs.org . Verify with node -v .
npm comes with Node.js. Verify with npm -v .
Git — required for the pre-commit hook and git diff piping.
npx -y -p overreach overreach-cli demo
Runs the real pipeline on a sample diff — no API key, no setup, costs nothing.
Exits 1 with a HIGH scope-creep finding (the demo prompt asks for a login form;
the diff smuggles in Stripe, an env var, an endpoint, and a cron job). That's the
whole product in one command.
A diff is flagged when it adds something the prompt never authorized:
Severity: env / endpoint / cron = high · dep / file = medium · feature = low .
Overall scope_creep_score : HIGH if any high finding, MEDIUM if any medium, else LOW .
Stage 1 — Scope extraction (LLM). Reads your prompt and produces an
authorized scope JSON: which files, features, deps, endpoints, env, and behaviors
you actually asked for. Deciphers typos to the nearest real concept but never
invents scope . This is the only stage that calls a model.
Stage 2 — Diff parsing (deterministic, no LLM). Regex-parses the diff into the
set of things it actually adds — imports, deps, process.env.X references, route
handlers, cron jobs, new symbols. Runs in milliseconds.
Stage 3 — Comparison (deterministic). Set arithmetic with fuzzy matching:
actual − authorized = findings .
Stages 2 and 3 are pure functions — no inference, no opinion, fully auditable.
That's what makes Overreach testable without spending a cent on inference.
npm install -g overreach
Or use directly via npx (no install needed):
npx -y -p overreach overreach-cli demo
API key (optional)
For best results, set one LLM provider key for Stage 1 scope extraction:
Pin a provider/model with SCOPE_PROVIDER and OVERREACH_MODEL .
No key? No problem. Without an API key, Overreach falls back to
deterministic scope extraction — it regex-parses your prompt for concrete
items (file paths, package names, /api/... routes, SCREAMING_SNAKE_CASE env
vars, cron keywords) instead of calling an LLM. It won't understand vague
instructions as well as an LLM would, but it catches every concrete noun in
your prompt. Instant, free, fully offline.
1. Set up a project (one command)
npx -y -p overreach overreach-cli init
This creates three things:
.overreach/prompt.md — write the prompt you gave your agent here
.git/hooks/pre-commit — audits every commit against your prompt
CLAUDE.md — instructs AI agents to self-audit before committing
Edit .overreach/prompt.md with the actual instruction you gave your AI agent:
Add a login form to the settings page with email/password fields,
form validation, and a submit button that calls /api/auth/login.
3. Commit — Overreach runs automatically
git add . && git commit -m " add login form "
The pre-commit hook audits staged changes against your prompt:
HIGH scope creep → commit blocked (exit 1)
MEDIUM / LOW → commit allowed with findings printed
Template prompt (not yet edited) → skipped gracefully
No API key → deterministic fallback (extracts concrete items from prompt)
Skip with git commit --no-verify when you know what you're doing. Update
.overreach/prompt.md whenever you give the agent a new task.
Windows: The pre-commit hook is a shell script. It works out of the box
with Git Bash (included with Git for Windows ).
npx -y -p overreach overreach-cli --prompt " add a login form to the settings page " --diff my-changes.diff
Or pipe a diff:
git diff | npx -y -p overreach overreach-cli --prompt " add a login form to the settings page "
Exits 0 if clean, 1 if HIGH — usable as a CI gate.
--prompt <text> — the instruction that authorized the work
--diff <path> — diff file (default: read from stdin)
--scope <path|json> — inject authorized scope; skips the LLM entirely
--json — emit raw JSON instead of pretty terminal output
--no-cache — bypass the scope cache (force a fresh Stage 1 call)
demo — run the canonical demo (zero-key)
init — install pre-commit hook + CLAUDE.md
MCP server (Claude Code, Cursor, Codex, Claude Desktop)
Overreach is a stdio MCP server, so any MCP-capable client can connect:
claude mcp add overreach -- npx -y overreach
Claude Desktop / Cursor — add to your MCP config:
{
"mcpServers" : {
"overreach" : { "command" : " npx " , "args" : [ " -y " , " overreach " ] }
}
}
Codex CLI — add to ~/.codex/config.toml :
[ mcp_servers . overreach ]
command = " npx "
args = [ " -y " , " overreach " ]
Or Streamable HTTP: set PORT=8787 and POST to http://localhost:8787/mcp .
The HTTP endpoint has no auth. It binds to 127.0.0.1 (loopback) by
default — safe for local use. Do not expose it publicly
( OVERREACH_HOST=0.0.0.0 ) without an authed reverse proxy in front: anyone who
can reach it can call check_overreach and spend your LLM budget.
Tools exposed: check_overreach(prompt, diff, options?) and health .
First-time setup (Claude Code)
# 1. Register the server with Claude Code (one time)
claude mcp add overreach -- npx -y overreach
# 2. Restart your Claude Code session
# (a session already open won't see the new server until you quit and reopen it)
# 3. Optionally set an API key (works without one via deterministic fallback)
export ANTHROPIC_API_KEY=sk-... # or OPENAI_API_KEY / OLLAMA_API_KEY
After the restart, every new session has check_overreach available — no per-task
setup. The agent calls it when it decides it's relevant.
The key isn't passed through automatically. The MCP server is a separate
process; your agent does not hand it its own credentials. If you log in to
Claude Code with claude login (OAuth / subscription), there's no
ANTHROPIC_API_KEY in the environment — so export one (any provider works; local
Ollama needs no key), or for Claude Desktop / Cursor add it to the server's env :
{ "mcpServers" : { "overreach" : { "command" : " npx " , "args" : [ " -y " , " overreach " ], "env" : { "ANTHROPIC_API_KEY" : " sk-... " } } } }
overreach init adds a scope-audit instruction to your project's CLAUDE.md so
AI agents self-audit their staged changes before committing — no user intervention
needed. The agent reads the instruction and runs Overreach on its own diff.
You can also have the agent call check_overreach directly via the MCP server
with its own task string + the diff it's about to commit:
git diff --staged | overreach-cli --prompt "<the task you just gave me>"
This is best-effort — an agent can skip the call or ignore the findings
(fox guarding the henhouse). The hard backstop is the CI gate below.
The hard backstop. A workflow runs Overreach on every pull request and fails
the PR when scope_creep_score=HIGH — the diff adds a dep / env var /
endpoint / cron / out-of-scope file the prompt didn't authorize.
Copy .github/workflows/overreach.yml into
your repo and add ANTHROPIC_API_KEY (or OPENAI_API_KEY / OLLAMA_API_KEY )
as a repository secret. The prompt comes from .overreach/prompt.md in the repo,
or the PR title + body if that file is absent. The job posts its findings as a PR
comment and fails the check on HIGH . Full setup + customization in
docs/ci-gate.md .
# .github/workflows/overreach.yml (excerpt)
- name : Run Overreach
run : |
npx -y -p overreach@latest overreach-cli \
--prompt "$(cat "$RUNNER_TEMP/prompt.txt")" --diff "$RUNNER_TEMP/pr.diff"
env :
ANTHROPIC_API_KEY : ${{ secrets.ANTHROPIC_API_KEY }}
- name : Gate — fail the PR on HIGH
if : steps.overreach.outputs.exit == '1'
run : exit 1
This open-source Action is free to run (you bring your own LLM key).
Model
Result
Claude Sonnet 4.6
82/82
Claude Opus 4.6
65/65
GLM 5.2
82/82
Kimi K2.7-Code
82/82
MiniMax M3
81/82
The deterministic fallback (no key) works with any prompt that contains concrete
items — no model needed.
Verify it works (zero API key)
npm test
Runs 56 assertions through the real pipeline with the scope injected via
scopeOverride , so Stage 1 (the LLM) is never called. Covers overreach
detection, clean passes, Python/Express/Next.js parsers, deletion handling,
determinism, chunking, and the trust contract invariant.
Overreach is fully self-contained. It does not import or depend on any other
project. It reads only its own process environment. No telemetry, no call-home —
it runs entirely on your machine.
If Overreach misses something it should flag, or flags something the prompt
authorized, open an issue with the prompt + the smallest repro diff :
https://github.com/Naveja00/OverReach/issues
There's a bug-report template that asks for exactly that.
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
