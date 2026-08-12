---
source: "https://github.com/lilcipherx/drift"
hn_url: "https://news.ycombinator.com/item?id=49277475"
title: "Drift – Intent-driven versioning for AI coding agents"
article_title: "GitHub - lilcipherx/drift: Drift — Intent-Driven Versioning. Git tracks what changed; Drift tracks why. Signed intents, semantic blame, MCP tools for AI agents. (Node 24, zero native deps) · GitHub"
author: "lilcipherx"
captured_at: "2026-08-12T19:56:33Z"
capture_tool: "hn-digest"
hn_id: 49277475
score: 1
comments: 0
posted_at: "2026-08-12T19:32:41Z"
tags:
  - hacker-news
  - translated
---

# Drift – Intent-driven versioning for AI coding agents

- HN: [49277475](https://news.ycombinator.com/item?id=49277475)
- Source: [github.com](https://github.com/lilcipherx/drift)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T19:32:41Z

## Translation

タイトル: Drift – AI コーディング エージェント向けのインテント駆動型バージョニング
記事のタイトル: GitHub - lilcipherx/drift: ドリフト — インテント駆動型バージョニング。 Git は何が変更されたかを追跡します。ドリフトはその理由を追跡します。 AI エージェント用の署名付きインテント、セマンティック責任、MCP ツール。 (ノード 24、ネイティブ デプスはゼロ) · GitHub
説明: ドリフト — インテント主導型のバージョニング。 Git は何が変更されたかを追跡します。ドリフトはその理由を追跡します。 AI エージェント用の署名付きインテント、セマンティック責任、MCP ツール。 (ノード 24、ネイティブ DEPS はゼロ) - lilcipherx/drift

記事本文:
GitHub - lilcipherx/drift: ドリフト — インテント主導のバージョニング。 Git は何が変更されたかを追跡します。ドリフトはその理由を追跡します。 AI エージェント用の署名付きインテント、セマンティック責任、MCP ツール。 (ノード 24、ネイティブ デプスはゼロ) · GitHub
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
リリシフェルクス
/
ドリフト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
69 コミット 69 コミット .claude-plugin .claude-plugin .codex-plugin .codex-plugin .cursor-plugin .cursor-plu

gin .factory-plugin .factory-plugin .github .github .opencode .opencode .plugin .plugin docs docs eval eval 例 例 マイグレーション マイグレーション パッケージ パッケージ プロンプト プロンプト スクリプト スキル/ドリフト スキル/ドリフト テスト テスト .gitignore .gitignore AUDIT.md AUDIT.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Drift (Intent-Driven Versioning).md Drift (Intent-Driven Versioning).md ライセンス ライセンス NEXT_STEPS.md NEXT_STEPS.md README.md README.md SECURITY.md SECURITY.md action.yml action.yml gemini-extension.json gemini-extension.json mcp.json mcp.json mcp_config.json mcp_config.json package-lock.json package-lock.json package.json package.json plugin.json plugin.json tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Git は何が変更されたかを追跡します。ドリフトはその理由を追跡します。
Drift は、Git をラップするセマンティックなバージョン管理レイヤーです。すべてのコミットは
意図 : 変更を生成したプロンプト、その背後にあるエージェント モデル、
AST レベルの突然変異、エージェントの認知状態のオプションのチェックポイント、および
暗号化 Ed25519 署名 — すべては監査可能で再生可能なグラフにリンクされています。
AI 時代に向けて構築されました。コードの 80% 以上が生成されると、テキストの差分が
レビューには役に立ちません。何が変更されたのかが示されており、その理由は示されていません。ドリフト拒否が壊れた
履歴に入る前の構文は、「なぜこの関数が存在するのか?」という質問に答えます。と
元のプロンプトを表示し、クラッシュしたエージェントを最後のチェックポイントから再開できるようにします。
エージェントにドリフトを与えます: クロード コード 、反重力 、
コーデックス アプリ 、コーデックス CLI 、カーソル 、
ファクトリードロイド 、 Gemini CLI 、
GitHub Copilot CLI、キミコード、
オープンコード、Pi 。
エージェントを使用しないほうがいいですか? CLI、GitHub アプリ、
GitHub アクション 、または VS Code 。 5分間の「なるほど」が欲しい
まず

?デモ リポジトリをシードし、ドリフト非難を実行します。
ドキュメント: クイックスタート (5 分で開始) ·
API リファレンス (CLI + MCP ツール) ·
アーキテクチャ (Drift が内部でどのように動作するか)
Windows 11 でライブ検証済み (ノード v24.18.0、2026 年 8 月 6 日): 新しいクローン →
最初のドリフト非難は約 8.1 秒以内、10/10 チェックはパス、レジストリ 404 なし — そして
npm パス (空のディレクトリにインストールされたパックされた @drift/* チェーン) が MCP に応答します
6 つのツールすべてとのハンドシェイクは 1 秒以内に完了します。を参照してください。
完全な測定テーブル。
これは、drift init を実行した瞬間から始まります。 Drift は .drift/ — SQLite DAG を作成します。
構成、およびリポジトリごとの Ed25519 キーペア — そしてそれ以降、すべてのコミットは次のようになります。
意図。
あなた (またはあなたのエージェント) がdrift Realize -p "<prompt>" を実行すると、Driftは単に
コミットします。変更を意味的に解析し、構文が次の場合はコミットを拒否します。
壊れた (出口 2 - 壊れたコードは履歴には入りません)、あなたの秘密を編集します。
プロンプト、AST デルタを計算します (追加 / 変更 / 削除 / 移動 / 名前変更)、
インテントに署名し、コンテンツアドレスを指定して .drift/objects/ に保存します。
Drift-Intent: トレーラーを使用してコミットします。
その後、ドリフト非難は任意のラインを歩くか、プロンプトに戻ることができます。
ファイルが作成されると、ドリフト コンテキストがファイルの最後のインテントをハイドレートするため、エージェントは
編集前に自らを接地し、記録されたものをドリフト検証で再実行します。
検証コマンド。クラッシュしたエージェントがドリフト リプレイ --checkout を実行して再開する
まさに中断したところから。
これらは MCP ツールであるため、コーディング エージェントはこれらを直接使用できます。
git commit の代わりにdrift_realizeを使用します。
さらに詳しく読む: 完全なコマンド リファレンスは次の場所にあります。
docs/api.md (CLI フラグ、終了コード、JSON スキーマ、MCP ツール)
入力)、docs/architecture.md で説明されています。
ストレージ モデル、保存時の暗号化、Webhook アプリ、およびセキュリティ
境界線。
ハーネスにより取り付けが異なります。 mを使用する場合

複数の場合は、Drift をインストールしてください
それぞれ別に。すべてのハーネスは同じ 6 つのツールを公開します。
drift_realize、drift_context、drift_replay、drift_blame、drift_verify、
ドリフト_ログ 。
ステータス: @drift/* npm パッケージはまだ公開されていません。毎
以下のセクションではクローン パスが示されています。これは現在、
このリポジトリをチェックアウトします。 npx -y @drift/mcp / npx -y @drift/cli
ワンライナーは、パッケージが npm に到達すると自動的に有効になります。まで
その場合、404 が返されるため、最初に示した clone コマンドを使用します。
以下のすべてのコマンドは、このリポジトリ内の実際のマニフェストによってサポートされています
( .claude-plugin/plugin.json 、 .plugin/plugin.json 、
.cursor-plugin/plugin.json 、.codex-plugin/plugin.json 、
gemini-extension.json 、 plugin.json 、 .factory-plugin/ 、
package.json → pi ）または既製の設定
例/harness-configs/ 。現在、インストールには Node.js ≥ 24 が必要です。
npm とこのリポジトリのクローン (MCP サーバーは直接実行されます)
パッケージ/drift-mcp/dist/index.js ;ビルド手順は必要ありません)。一度
@drift/* パッケージが公開され、同じ構成が経由で動作します
npx -y @drift/mcp (クローンなし)。
Drift マーケットプレイスからプラグインとしてインストールします (Superpowers などのプラグイン スタイル)。
/プラグイン マーケットプレイス lilcipherx/drift を追加
/プラグインインストールdrift@drift
または、クローンから Drift MCP サーバーを直接追加します (プロジェクト スコープ)。
クロード mcp adddrift --env DRIFT_REPO=/abs/path/to/your/repo --node /path/to/drift/packages/drift-mcp/dist/index.js
パッケージが公開されると、npx (クローンなし) 経由で同じコマンドが機能します。
クロード mcp ドリフトを追加 --env DRIFT_REPO=/abs/path/to/your/repo -- npx -y @drift/mcp
または、既製の構成をコピーします。
cp 例/harness-configs/claude-code/.mcp.json .mcp.json
claude mcp list で確認します。6 つのツールを備えたドリフトが表示されるはずです。
マーケットプレイスのマニフェストは .clau にあります

この中の de-plugin/marketplace.json
リポジトリ (github-source lilcipherx/drift 、strict プラグイン →
同じリポジトリ内の .claude-plugin/plugin.json)。
このリポジトリから Drift をプラグインとしてインストールします。
agy プラグインのインストール https://github.com/lilcipherx/drift
Antigravity はプラグインのセッション開始フックを実行するため、Drift は
最初のメッセージ。同じコマンドで再インストールして更新します。
Codex アプリで、[設定] → [MCP サーバー] を開き、以下を追加します。
コマンド:node (クローンパス) — またはパッケージが公開されたら npx
引数 : /path/to/drift/packages/drift-mcp/dist/index.js — または公開後の -y @drift/mcp
環境: DRIFT_REPO=/abs/path/to/your/repo
Drift MCP サーバーを ~/.codex/config.toml に追加します (クローンから)。
[ mcp_servers .ドリフト】
コマンド = " ノード "
args = [ " /path/to/drift/packages/drift-mcp/dist/index.js " ]
env = { DRIFT_REPO = " /abs/path/to/your/repo " }
公開されると、同じサーバーが npx (クローンなし) 経由で実行されます。
[ mcp_servers .ドリフト】
コマンド = " npx "
args = [ " -y " , " @drift/mcp " ]
env = { DRIFT_REPO = " /abs/path/to/your/repo " }
Codex を再起動し、任意のセッションでdrift_blame /drift_contextを要求します。
既製の設定をコピーし、カーソル設定で MCP サーバーを有効にします。
cp 例/harness-configs/cursor/mcp.json .cursor/mcp.json
次に、チャットでdrift_blameについて尋ねます。
droid プラグイン マーケットプレイスを追加 https://github.com/lilcipherx/drift
droid プラグインのインストールdrift@drift
gemini 拡張機能のインストール https://github.com/lilcipherx/drift
Gemini 拡張機能の更新ドリフト
Drift MCP サーバーを (クローンから) 追加します。
copilot mcp adddrift -e DRIFT_REPO=/abs/path/to/your/repo -- ノード /path/to/drift/packages/drift-mcp/dist/index.js
パッケージが公開されると、npx (クローンなし) 経由で同じコマンドが機能します。
copilot mcp adddrift -e DRIFT_REPO=/abs/path/to/your/repo -- npx -y @drift/mcp
またはコピーしてください

既製の設定を .github/mcp.json にコピーし、Copilot を再起動します。
Drift は、 Kim Code のプラグイン マーケットプレイスで入手できます。
Kim Code のプラグイン マネージャーを開きます。
マーケットプレイス > Drift に移動してインストールします。
または、このリポジトリから直接インストールします。
/plugins インストール https://github.com/lilcipherx/drift
詳細ドキュメント: docs/README.kimi.md
OpenCode は独自のプラグインのインストールを使用します。 Drift を個別にインストールする場合でも、
すでに別のハーネスで使用しています。
https://raw.githubusercontent.com/lilcipherx/drift/main/.opencode/INSTALL.md から手順を取得して実行します。
詳細ドキュメント: docs/README.opencode.md
このリポジトリから Drift を Pi パッケージとしてインストールします。
pi インストール git:github.com/lilcipherx/drift
ローカル開発の場合は、このチェックアウトを一時パッケージとしてロードして Pi を実行します。
Drift MCP サーバーを .vscode/mcp.json に追加します (ネイティブ VS Code MCP サポート)。
cp 例/harness-configs/claude-code/.mcp.json .vscode/mcp.json
VS Code はウィンドウのリロード時にそれを取得します。
git clone https://github.com/lilcipherx/drift.git && cd ドリフト
npmインストール
ノードパッケージ/drift-cli/dist/cli.js --help
CLI が npm に公開されると、npx はクローンを作成せずに機能します。
npx -y @drift/cli --help
自分のリポジトリ上:
ノード /path/to/drift/packages/drift-cli/dist/cli.js init
# … ファイルを編集する …
node /path/to/drift/packages/drift-cli/dist/cli.js Realize -p " トークン更新時の競合状態を修正 " --agent --model claude-3-5-sonnet
ノード /path/to/drift/packages/drift-cli/dist/cli.js ログ
GitHub アプリ
@drift/app をインストールして、すべてのプル リクエストのインテント サマリーを取得します。
PR コミットから Drift-Intent: トレーラーを読み取り、インテント オブジェクトをハイドレートします。
.drift/objects/ を PR ヘッドに配置し、セマンティックな概要コメントを投稿します —
2,000 行の差分ではなく、意図を確認してください。コメントはべき等です: アプリ
独自のマーカー コメントをその場で更新するため、コメントが蓄積されることはありません

食べた。
Webhook サーバーとして実行:drift-app start (packages/drift-app/app.yml を参照)
アプリ マニフェストの場合は scripts/webhook-proxy.sh (ローカル デバッグの場合)。
- 使用: lilcipherx/drift@v0.3.0
付き:
コマンド：log # または：doctor / verify <intent-id>
デモ
5 分間の「なるほど」 - CLI 自体によって生成された本物のドリフト履歴:
bash スクリプト/seed-demo.sh
CD サンプル/デモリポジトリ
ノード ../../packages/drift-cli/dist/cli.js ログ
ノード ../../packages/drift-cli/dist/cli.js の責任 src/auth.ts --function freshToken
blame は、関数のプロンプト、モデル、および有効な署名を出力します。
src/auth.ts:12 (リフレッシュトークン)
エージェント@ドリフトデモ
モデル: クロード-3-5-ソネット
プロンプト: 実行中のリフレッシュの重複を排除して、トークンのリフレッシュにおける競合状態を修正します。
意図: Did_2941b4547b4ed505a7c37190247768a7
コミット: 087c492f… 署名: 有効
基本的なワークフロー
init — .drift/ (SQLite DAG、config、Ed25519 キーペア) を作成します。決して書き換えない
歴史; .drift/ を削除すると、完全に機能する git リポジトリが残ります。
実現 — 意図を持ってコミットする。構文ゲート (出口 2)、シークレット編集、
AST デルタ、Ed25519 署名、ドリフトインテント: トレーラー。これは git commit です
AIの時代。
log — インテントのタイムライン: ID、作成者 (エージェントと人間)、モデル、プロンプト。
非難 / コンテキスト — 非難 --line|-- 関数は、シンボルを元の場所に戻します。
オリ

[切り捨てられた]

## Original Extract

Drift — Intent-Driven Versioning. Git tracks what changed; Drift tracks why. Signed intents, semantic blame, MCP tools for AI agents. (Node 24, zero native deps) - lilcipherx/drift

GitHub - lilcipherx/drift: Drift — Intent-Driven Versioning. Git tracks what changed; Drift tracks why. Signed intents, semantic blame, MCP tools for AI agents. (Node 24, zero native deps) · GitHub
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
lilcipherx
/
drift
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
69 Commits 69 Commits .claude-plugin .claude-plugin .codex-plugin .codex-plugin .cursor-plugin .cursor-plugin .factory-plugin .factory-plugin .github .github .opencode .opencode .plugin .plugin docs docs eval eval examples examples migrations migrations packages packages prompts prompts scripts scripts skills/ drift skills/ drift tests tests .gitignore .gitignore AUDIT.md AUDIT.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Drift (Intent-Driven Versioning).md Drift (Intent-Driven Versioning).md LICENSE LICENSE NEXT_STEPS.md NEXT_STEPS.md README.md README.md SECURITY.md SECURITY.md action.yml action.yml gemini-extension.json gemini-extension.json mcp.json mcp.json mcp_config.json mcp_config.json package-lock.json package-lock.json package.json package.json plugin.json plugin.json tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json View all files Repository files navigation
Git tracks what changed. Drift tracks why.
Drift is a semantic version-control layer that wraps Git. Every commit becomes an
Intent : the prompt that produced the change, the agent model behind it, the
AST-level mutations, an optional checkpoint of the agent's cognitive state, and a
cryptographic Ed25519 signature — all linked into an auditable, replayable graph.
Built for the AI era. When more than 80% of code is generated, text diffs are
useless for review: they show what changed, never why . Drift rejects broken
syntax before it enters history, answers "why does this function exist?" with
the originating prompt, and lets a crashed agent resume from its last checkpoint.
Give your agent Drift: Claude Code , Antigravity ,
Codex App , Codex CLI , Cursor ,
Factory Droid , Gemini CLI ,
GitHub Copilot CLI , Kimi Code ,
OpenCode , Pi .
Prefer no agent? Use the CLI , the GitHub App , the
GitHub Action , or VS Code . Want the 5-minute "aha"
first? Seed the demo repo and run drift blame .
Documentation: Quickstart (5-minute start) ·
API reference (CLI + MCP tools) ·
Architecture (how Drift works under the hood)
Verified live on Windows 11 (Node v24.18.0, 2026-08-06): fresh clone →
first drift blame in ~8.1 s, 10/10 checks pass, no registry 404 — and the
npm path (packed @drift/* chain installed into an empty dir) answers the MCP
handshake with all six tools in ~1 s. See the
full measured tables .
It starts the moment you run drift init . Drift creates .drift/ — a SQLite DAG,
a config, and a per-repo Ed25519 keypair — and from then on every commit becomes
an intent.
When you (or your agent) run drift realize -p "<prompt>" , Drift doesn't just
commit. It parses the change semantically, rejects the commit if the syntax is
broken (exit 2 — broken code never enters history), redacts secrets from your
prompt, computes an AST delta (ADDED / MODIFIED / DELETED / MOVED / RENAMED),
signs the intent, and stores it content-addressed in .drift/objects/ before
committing with a Drift-Intent: trailer.
After that, drift blame can walk any line or function back to the prompt that
created it, drift context hydrates the last intents for a file so an agent
grounds itself before editing, and drift verify re-runs the recorded
verification command. A crashed agent runs drift replay --checkout and resumes
exactly where it left off.
And because these are MCP tools, your coding agent can use them directly —
drift_realize instead of git commit .
Deeper reading: the full command reference lives in
docs/api.md (CLI flags, exit codes, JSON schemas, MCP tool
inputs), and docs/architecture.md explains the
storage model, encryption at rest, the webhook app, and the security
boundaries.
Installation differs by harness. If you use more than one, install Drift
separately for each one. All harnesses expose the same six tools:
drift_realize , drift_context , drift_replay , drift_blame , drift_verify ,
drift_log .
Status: the @drift/* npm packages are not published yet. Every
section below leads with the clone path — it works right now from a
checkout of this repository. The npx -y @drift/mcp / npx -y @drift/cli
one-liners activate automatically once the packages land on npm; until
then they return a 404, so use the clone command shown first.
Every command below is backed by a real manifest in this repository
( .claude-plugin/plugin.json , .plugin/plugin.json ,
.cursor-plugin/plugin.json , .codex-plugin/plugin.json ,
gemini-extension.json , plugin.json , .factory-plugin/ ,
package.json → pi ) or a ready-made config in
examples/harness-configs/ . Today, installation needs Node.js ≥ 24,
npm and a clone of this repository (the MCP server runs straight from
packages/drift-mcp/dist/index.js ; no build step needed). Once the
@drift/* packages are published, the same configs work via
npx -y @drift/mcp with no clone.
Install as a plugin from the Drift marketplace (plugin-style, like Superpowers):
/plugin marketplace add lilcipherx/drift
/plugin install drift@drift
Or add the Drift MCP server directly (project scope) — from a clone:
claude mcp add drift --env DRIFT_REPO=/abs/path/to/your/repo -- node /path/to/drift/packages/drift-mcp/dist/index.js
Once the packages are published, the same command works via npx (no clone):
claude mcp add drift --env DRIFT_REPO=/abs/path/to/your/repo -- npx -y @drift/mcp
Or copy the ready-made config:
cp examples/harness-configs/claude-code/.mcp.json .mcp.json
Verify with claude mcp list — you should see drift with its six tools.
The marketplace manifest lives at .claude-plugin/marketplace.json in this
repository (github-source lilcipherx/drift , strict plugin →
.claude-plugin/plugin.json in the same repo).
Install Drift as a plugin from this repository:
agy plugin install https://github.com/lilcipherx/drift
Antigravity runs the plugin's session-start hook, so Drift is active from the
first message. Reinstall with the same command to update.
In the Codex app, open Settings → MCP servers and add:
Command : node (clone path) — or npx once the packages are published
Args : /path/to/drift/packages/drift-mcp/dist/index.js — or -y @drift/mcp after publication
Env : DRIFT_REPO=/abs/path/to/your/repo
Add the Drift MCP server to ~/.codex/config.toml (from a clone):
[ mcp_servers . drift ]
command = " node "
args = [ " /path/to/drift/packages/drift-mcp/dist/index.js " ]
env = { DRIFT_REPO = " /abs/path/to/your/repo " }
Once published, the same server runs via npx (no clone):
[ mcp_servers . drift ]
command = " npx "
args = [ " -y " , " @drift/mcp " ]
env = { DRIFT_REPO = " /abs/path/to/your/repo " }
Restart Codex, then ask for drift_blame / drift_context in any session.
Copy the ready-made config and enable MCP servers in Cursor settings:
cp examples/harness-configs/cursor/mcp.json .cursor/mcp.json
Then ask for drift_blame in chat.
droid plugin marketplace add https://github.com/lilcipherx/drift
droid plugin install drift@drift
gemini extensions install https://github.com/lilcipherx/drift
gemini extensions update drift
Add the Drift MCP server (from a clone):
copilot mcp add drift -e DRIFT_REPO=/abs/path/to/your/repo -- node /path/to/drift/packages/drift-mcp/dist/index.js
Once the packages are published, the same command works via npx (no clone):
copilot mcp add drift -e DRIFT_REPO=/abs/path/to/your/repo -- npx -y @drift/mcp
Or copy the ready-made config to .github/mcp.json and restart Copilot.
Drift is available in Kimi Code's plugin marketplace.
Open Kimi Code's plugin manager:
Go to Marketplace > Drift and install it.
Or install directly from this repository:
/plugins install https://github.com/lilcipherx/drift
Detailed docs: docs/README.kimi.md
OpenCode uses its own plugin install; install Drift separately even if you
already use it in another harness.
Fetch and follow instructions from https://raw.githubusercontent.com/lilcipherx/drift/main/.opencode/INSTALL.md
Detailed docs: docs/README.opencode.md
Install Drift as a Pi package from this repository:
pi install git:github.com/lilcipherx/drift
For local development, run Pi with this checkout loaded as a temporary package:
Add the Drift MCP server to .vscode/mcp.json (native VS Code MCP support):
cp examples/harness-configs/claude-code/.mcp.json .vscode/mcp.json
VS Code picks it up on window reload.
git clone https://github.com/lilcipherx/drift.git && cd drift
npm install
node packages/drift-cli/dist/cli.js --help
Once the CLI is published to npm, npx works without cloning:
npx -y @drift/cli --help
On your own repository:
node /path/to/drift/packages/drift-cli/dist/cli.js init
# … edit a file …
node /path/to/drift/packages/drift-cli/dist/cli.js realize -p " Fix race condition in token refresh " --agent --model claude-3-5-sonnet
node /path/to/drift/packages/drift-cli/dist/cli.js log
GitHub App
Install @drift/app to get intent summaries on every pull request:
Reads Drift-Intent: trailers from PR commits, hydrates the intent objects from
.drift/objects/ at the PR head, and posts a semantic summary comment —
review the intent, not 2,000 lines of diff. Comments are idempotent: the app
updates its own marker comment in place, so they never accumulate.
Runs as a webhook server: drift-app start (see packages/drift-app/app.yml
for the app manifest, scripts/webhook-proxy.sh for local debugging).
- uses : lilcipherx/drift@v0.3.0
with :
command : log # or: doctor / verify <intent-id>
Demo
The 5-minute "aha" — a real Drift history, generated by the CLI itself:
bash scripts/seed-demo.sh
cd examples/demo-repo
node ../../packages/drift-cli/dist/cli.js log
node ../../packages/drift-cli/dist/cli.js blame src/auth.ts --function refreshToken
blame prints the prompt, model and a valid signature for the function:
src/auth.ts:12 (refreshToken)
AGENT @ Drift Demo
model: claude-3-5-sonnet
prompt: Fix race condition in token refresh by de-duplicating in-flight refreshes
intent: did_2941b4547b4ed505a7c37190247768a7
commit: 087c492f… signature: valid
The Basic Workflow
init — Creates .drift/ (SQLite DAG, config, Ed25519 keypair). Never rewrites
history; deleting .drift/ leaves a fully functional git repo.
realize — Commit with intent. Syntax gate (exit 2), secret redaction,
AST delta, Ed25519 signature, Drift-Intent: trailer. This is git commit for
the AI era.
log — Timeline of intents: id, author (agent vs human), model, prompt.
blame / context — blame --line|--function walks a symbol back to its
ori

[truncated]
