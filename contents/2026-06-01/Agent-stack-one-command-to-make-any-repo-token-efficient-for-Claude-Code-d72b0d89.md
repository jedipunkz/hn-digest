---
source: "https://github.com/drmahdikazempour/agent-stack"
hn_url: "https://news.ycombinator.com/item?id=48350102"
title: "Agent-stack – one command to make any repo token-efficient for Claude Code"
article_title: "GitHub - drmahdikazempour/agent-stack: One command to make any repo token-efficient for Claude Code & Cursor. Built-in code map, output compression, skills & hooks, savings measured with ccusage. · GitHub"
author: "mahdikaz"
captured_at: "2026-06-01T00:24:52Z"
capture_tool: "hn-digest"
hn_id: 48350102
score: 1
comments: 0
posted_at: "2026-05-31T21:51:06Z"
tags:
  - hacker-news
  - translated
---

# Agent-stack – one command to make any repo token-efficient for Claude Code

- HN: [48350102](https://news.ycombinator.com/item?id=48350102)
- Source: [github.com](https://github.com/drmahdikazempour/agent-stack)
- Score: 1
- Comments: 0
- Posted: 2026-05-31T21:51:06Z

## Translation

タイトル: Agent-stack – クロード コードのリポジトリ トークンを効率的にするための 1 つのコマンド
記事のタイトル: GitHub - drmahdikzempour/agent-stack: クロード コードとカーソルのリポジトリ トークンを効率的にするための 1 つのコマンド。組み込みのコードマップ、出力圧縮、スキルとフック、ccusage で測定された節約。 · GitHub
説明: クロード コードとカーソルのリポジトリ トークンを効率的にする 1 つのコマンド。組み込みのコードマップ、出力圧縮、スキルとフック、ccusage で測定された節約。 - drmahdikazempour/エージェントスタック

記事本文:
GitHub - drmahdikzempour/agent-stack: クロード コードとカーソルのリポジトリ トークンを効率的にするための 1 つのコマンド。組み込みのコードマップ、出力圧縮、スキルとフック、ccusage で測定された節約。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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

ドルマディカゼムプール
/
エージェントスタック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
14 コミット 14 コミット .changeset .changeset .github/ workflows .github/ workflows bin bin 統合 統合 スキル スキル src src テンプレート テンプレート テスト test .gitignore .gitignore CREDITS.md CREDITS.md ライセンス ライセンス PRD-claude-stack-v3.md PRD-claude-stack-v3.md README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コマンドは 1 つです。すべてがセットアップされました。アクティブ化されました。測定しました。
クロード コード (+ Cursor ) 用のスキル優先の最適化ツールキット。あらゆるリポジトリを 2 分以内にゼロから完全に最適化されたトークン効率の高いエージェント セットアップに移行します。
npx @drmahdikzempour/agent-stack init --all
🤖 これをクロードコードに貼り付けてください
以下のブロックをリポジトリのルートにあるクロード コード セッションにコピーします。フルスタックのインストールとアクティブ化、結果の監査、トークン使用量のベースライン設定、メインへのプッシュをすべて一度に実行します。
このリポジトリでエージェント スタックをセットアップし、出荷します。次のすべてを実行します。
1. フルスタックをインストールしてアクティブ化します。
npx -y @drmahdikzempour/agent-stack@latest init --all --yes
これによりホストが検出され、外部ツール (rtk、code-review-graph、
組み込みのフォールバックを介して、graphify、caveman、claude-handoff、gbrain)、
ツールコーディネーターを使用して CLAUDE.md + AGENTS.md を生成し、フックを配線します。
そしてトークンのベースラインをキャプチャします。いずれかのツールで「インストールの完了」と表示された場合
ガイダンス (ツールチェーンが欠落しています) がある場合は、そのコマンドを実行して続行してください。
2. AUDIT + LINT 結果を要約し、次のようにします。
みやび @miyabibi

/エージェントスタック監査
npx @drmahdikzempour/エージェントスタックドクター
CLAUDE.md トークンのバジェット、スキル/MCP バジェット、フックの競合、および
どのツールが利用可能か、それとも利用できないか。
3. トークンの使用状況を追跡する — トークンの使用方法を説明し、ベースラインを示します。
npx @drmahdikzempour/agent-stack 測定 --since 7d
使用量が記録される 1 日あたりのベースライン トークンを教えてください
(.agent-stack/usage.jsonl)、および 1 週間後に測定を再実行する必要があること
% の削減を確認します。 Stop フックは毎ターン自動的に攻撃を記録します。
4. コミット + メインにプッシュ:
生成されたファイルをステージングし、明確なメッセージを付けてコミットし、メインにプッシュします。
機能ブランチを使用している場合は、代わりに PR を開いて main にマージします。
次に、CI が緑色であることを確認します。
各ステップの出力を段階的に説明し、次の場合にのみ停止して質問してください。
このステップでは本当に私の入力が必要です (例: インストールできないツールチェーンが不足しているなど)。
ヒント: init --all (フルスタック) を推奨します。より軽量なセットアップの場合は、ステップ 1 をプレーンな npx -y @drmahdikzempour/agent-stack@latest init --yes に置き換えます。
Claude トークン最適化エコシステムは、シェル出力コンプレッサー、コンテキスト グラフ、出力スタイル、測定、連続性などの単一レイヤーのポイント ツールに断片化されています。どれも を構成しません。今リポジトリをセットアップするということは、5 ～ 10 個のツールを厳選し、各インストール ドキュメントを読み、フックを手動でマージし、 CLAUDE.md を手書きし、Cursor にミラーリングし、節約量を自分で測定することを意味します。
Agent-Stack はこれを 1 つのコマンドで実行し、トークン切断機構を組み込んで出荷するため、インストールする架空のものは何もなく、手動で配線する必要もありません。
cd あなたのリポジトリ
# スマートなデフォルト (ホスト、プロファイル、パッケージ マネージャーを自動検出):
npx @drmahdikzempour/エージェントスタック初期化
# …またはすべてを一度にオンにします (最大プロファイル):
npx @drmahdikzempour/agent-stack init --all
何が表示されるか (クリックして拡大)
エージェントスタック v0.2.0
検出されました:
主催：クラ

コード + カーソル
リポジトリ: TypeScript / Next.js / pnpm
プロファイル: 最大 (信頼度: 高)
書きます:
20ファイル → クロード、カーソル
.claude/settings.json (2 つのフックがマージされました)
続行しますか? [はい/いいえ] はい
✓ アダプター: ccusage (インストール済み)
✓ 20個のファイルが生成されました
✓ 2 つのフックを settings.json に接続
✓ すべてのスキルがロードされ、すべてのフックが存在し、CLAUDE.md が検証済み
✓ 構築されたコードマップ → .agent-stack/graph.md
✓ ベースライン: 12,340 トークン/日 (数、過去 7 日間の平均)
完了しました。
ヒント
これをグローバルに一度インストールすると、どこでも短い Agent-stack コマンドを取得できます。
npm i -g @drmahdikzempour/agent-stack
エージェントスタックの初期化 --all
⚙️ init の仕組み
ワンショット、10 ステップ、完全に可逆的です。 --dry-run は計画後に停止します。 --yes は 1 つの確認をスキップします。
検出 ─ 計画 ─ 確認 ─ バックアップ ─ インストール ─ 生成 ─ ワイヤーフック
│ │
--予行運転 ┘ (停止) ▼
アクティブ化する
要約 ◀── ベースライン ◀── コードマップ ◀──────────────┘ │
失敗 ──▶ ロールバック
#
ステップ
何が起こるか
1
検出
ホスト、リポジトリ タイプ、フレームワーク、パッケージ マネージャー、既存の構成、git の状態
2
計画
1 画面の計画を印刷します ( --dry-run はここで停止します)
3
確認する
単一の続行? [Y/n] ( --yes はスキップします)
4
バックアップ
既存の構成を .agent-stack.bak.<ts>/ にコピーします。
5
インストール
ccuseage のみ - その他はすべて組み込まれています
6
生成する
CLAUDE.md 、スキル、サブエージェント、コマンド、 .claudeignore 、カーソルミラー、MCP スキャフォールド
7
ワイヤーフック
1 つのマージされた .claude/settings.json への書き込み (単独のライター、重複排除、あなたのものを破壊することはありません)
8
アクティブ化
各スキルのロード、各フックの存在、CLAUDE.md の存在を確認します - 失敗時にロールバックします
9
コードマップ
初期の .agent-stack/graph.md を構築します
10
ベースライン
後で比較できるように Ccusage トークンのスナップショットを記録します
🔧 内蔵トークンカッター
これらは船内にあります

パッケージ — 外部インストールや架空のものはありません。これらは実際にトークンを削減するものです。
エージェントスタックグラフの更新 # 再構築 .agent-stack/graph.md
エージェントスタックグラフクエリ < シンボル > # 定義されている場所を検索
すべてのソース ファイル → エクスポートされたシンボルをマッピングするコンパクトなインデックス。エージェントは、ディレクトリ全体を読み取るのではなく、1 つの小さなファイルを grep して、何かが存在する場所を見つけます。 SessionStart 時に自動的に更新されます。 TypeScript/JavaScript、Python、Go、Rust をサポートします。
# コードマップ
_142 ファイル、906 個のトップレベル シンボル。ソースを開く前にこれを grep してシンボルを見つけます。_
- `src/core/detect.ts`: 検出
- `src/wire-hooks.ts`: mergeHooks、wireHooks、planHooks、countOurHooks
- …
🗜️ 出力圧縮
npm 実行ビルド 2>&1 | npx @drmahdikzempour/agent-stack 圧縮
ANSI コードを除去し、重複行 ( line (×42) ) を折り畳んで、巨大な出力を先頭/末尾で省略する stdin → stdout フィルター。500 行のログでは文字数が約 60 % 減少するため、ノイズの多いコマンドのコストはコンテキストの一部で済みます。
🪶 構造的節約 (常時有効)
.claudeignore は、node_modules、ビルド出力、メディア、およびロックファイルをコンテキスト外に保ちます。
CLAUDE.md は事実に基づいて厳密に生成されます。起動時に 800 個以下のトークンが生成され、医師によって検証されます。
サブエージェント ( stack-explorer 、 stack-reviewer ) は、ファイル ダンプではなく結論を返します。
簡潔モード (最大プロファイル) では、最小限の単語での回答が強制されます。
プロファイルには、グラフ バックエンド + 圧縮 + スキル セット + フック構成がバンドルされています。 init は自動的に 1 つを選択します。後でプロファイル use と交換します。
エージェント スタック プロファイルの使用レビュー # スワップと再生成
エージェント スタック プロファイル表示 # 現在のプロファイル
📟 コマンドリファレンス
エージェントスタック初期化 [--all] [--yes] [--dry-run] [--targets クロード、カーソル]
[--profile < 名前 > ] [--no-install] [--allow-noncommercial]
[--上書き] [--強制]
トークンカッター — スタンドアロン、パイプ内、または経由

フック
エージェントスタック圧縮 # cmd 2>&1 |エージェントスタックの圧縮
エージェントスタックグラフの更新 # コードマップを再構築する
エージェントスタックグラフクエリ < term > # シンボル/ファイルを検索
メンテナンス — インストール後、オンデマンドで
エージェント スタック監査 # トークン数 + 予算レポート
Agent-Stack Optimize # 監査修正を適用します (承認を得て)
Agent-Stack Doctor # lint everything (問題がある場合は exit 1)
エージェントスタック測定 [--since 7d] # ccusage ベースラインと現在
エージェント スタック プロファイル use < name > # swap プロファイル;再生する
エージェントスタックグラフ use < name > # 外部グラフに切り替えます (インストールされている場合)
エージェントスタックハンドオフ書き込み |セッション間の連続性を再開します
エージェントスタック同期 # CLAUDE.md からカーソルミラーを再生成
エージェントスタックアンインストール # バックアップを復元し、生成されたファイルを削除します
初期化フラグの詳細
旗
効果
--すべて
一度に完全な外部スタック (最大プロファイル): rtk + code-review-graph +graphify + caveman + claude-handoff + gbrain
--はい
単一の確認プロンプトをスキップする
--ドライラン
計画を印刷し、何も書かない
--targets クロード、カーソル
ホスト リストを強制します (カーソルはポータブル サブセットを取得します: rtk + MCP グラフ ツール)
--profile <名前>
プロファイルを強制する (コードレビューマルチモーダルスペックリサーチマックス)
--インストールなし
構成のみを書き込みます。インストールする代わりにインストール ガイドを印刷する
--上書き
既存のファイルをマージする代わりに置き換えます (バックアップは引き続き行われます)
--force
すでにインストールされている場合でも再実行します
🏗️ アーキテクチャ
1 つの真実のソース、2 つの顔、ランタイム依存性なし。
@drmahdikzempour/エージェントスタック
┌───────────────────────┐
│ CLI (src/) │
│ 組み込み/グラフ (コードマップ) · 圧縮 │
│ 生成/ クロード · カーソル · mcp │
│ Wire-hooks ← settings.json の唯一の作成者 │
│アクティ

vate ← 検証、または失敗時にロールバック │
│ スキル / 5 エージェントスキル │
━━━━━━━━━━━━━━━━━━━┘
書く │ ミラー │
▼ ▼
🟠クロードコード🔵カーソル
クロード.md .cursor/rules/*.mdc
.claude/スキル · エージェント AGENTS.md
コマンド · settings.json
.claudeignore · グラフ.md
スキルがいつになるかを決定します。 CLI がその方法を決定します。スキルは内部で CLI を呼び出します。通常、init を 1 回実行すると、再度 CLI に触れる必要はありません。
エージェントスタック/
§── bin/agent-stack.js # npx エントリポイント
§── src/
│ §── cli.ts # arg 解析 + コマンドディスパッチ
│ §── constants.ts # すべての仕様値 (トークンのバジェット、制限)
│ §── core/ # 検出 · 計画 · セーフライター · バックアップ · トークン推定子
│ §──builtin/ # グラフ（コードマップ）・compress（出力圧縮）
│ §── 生成/ # クロード · カーソル · mcp · コーディネーター ファイルビルダー
│ §── アダプタ/ # レジストリ · 検出ツール · インストール · フック
│ §── Wire-hooks.ts # settings.json フックの唯一のライター
│ §── activate.ts # 書き込み後の検証チェーン
│ §── Audit.ts # トークンバジェットのリンティング
│ └── コマンド/ # init + メンテナンスコマンド
§── スキル/ # 5 エージェント スキル (stack-bootstrap、-doctor、…)
§── 統合/ # profiles.json · tools.json
§── テンプレート/ # 生成メモ
└── test/ # vitest: tmpdir 内のユニット · ゴールデン · e2e init
📉 トークンのカット方法
節約は、ブラック ボックスからではなく、コンテキスト ウィンドウに入る内容を変更することで得られます。
❌ 前 ✅ 後 (エージェントスタック)
───────────────────

────────
12 個のファイルを読み込んで検索します ──▶ grepgraph.md → 1 個のファイルを開きます
1つの機能
500 行のログを貼り付ける ──▶ パイプスルー

[切り捨てられた]

## Original Extract

One command to make any repo token-efficient for Claude Code & Cursor. Built-in code map, output compression, skills & hooks, savings measured with ccusage. - drmahdikazempour/agent-stack

GitHub - drmahdikazempour/agent-stack: One command to make any repo token-efficient for Claude Code & Cursor. Built-in code map, output compression, skills & hooks, savings measured with ccusage. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
drmahdikazempour
/
agent-stack
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
14 Commits 14 Commits .changeset .changeset .github/ workflows .github/ workflows bin bin integrations integrations skills skills src src templates templates test test .gitignore .gitignore CREDITS.md CREDITS.md LICENSE LICENSE PRD-claude-stack-v3.md PRD-claude-stack-v3.md README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
One command. Everything set up. Activated. Measured.
A skills-first optimization toolkit for Claude Code (+ Cursor ) that takes any repo from zero to a fully optimized, token-efficient agent setup — in under two minutes.
npx @drmahdikazempour/agent-stack init --all
🤖 Paste this into Claude Code
Copy the block below into a Claude Code session at the root of any repo. It installs and activates the full stack, audits the result, baselines token usage, and pushes to main — all in one go.
Set up agent-stack in this repo and ship it. Do all of the following:
1. INSTALL + ACTIVATE the full stack:
npx -y @drmahdikazempour/agent-stack@latest init --all --yes
This detects the host, installs the external tools (rtk, code-review-graph,
graphify, caveman, claude-handoff, gbrain) over the built-in fallbacks,
generates CLAUDE.md + AGENTS.md with the tool coordinator, wires the hooks,
and captures a token baseline. If any tool prints "Finish installing"
guidance (missing toolchain), run that command for me, then continue.
2. AUDIT + LINT the result and summarize:
npx @drmahdikazempour/agent-stack audit
npx @drmahdikazempour/agent-stack doctor
Report the CLAUDE.md token budget, skill/MCP budgets, hook conflicts, and
which tools are available vs absent.
3. TRACK TOKEN USAGE — explain how it works and show the baseline:
npx @drmahdikazempour/agent-stack measure --since 7d
Tell me the baseline tokens/day, where usage is logged
(.agent-stack/usage.jsonl), and that I should re-run measure after ~a week
to see the % reduction. The Stop hook logs ccusage every turn automatically.
4. COMMIT + PUSH TO MAIN:
Stage the generated files, commit with a clear message, and push to main.
If I'm on a feature branch, open a PR and merge it to main instead.
Then confirm CI is green.
Walk me through each step's output as you go, and stop to ask me only if a
step genuinely needs my input (e.g. a missing toolchain you can't install).
Tip: prefer init --all (full stack). For a lighter setup, replace step 1 with plain npx -y @drmahdikazempour/agent-stack@latest init --yes .
The Claude token-optimization ecosystem is fragmented into single-layer point tools — shell-output compressors, context graphs, output styles, measurement, continuity. None of them compose . Setting up a repo today means hand-picking 5–10 tools, reading each install doc, hand-merging hooks, hand-writing CLAUDE.md , mirroring to Cursor, and measuring savings yourself.
agent-stack does it in one command — and ships the token-cutting machinery built in , so there's nothing fictional to install and nothing to wire by hand.
cd your-repo
# Smart defaults (auto-detects host, profile, package manager):
npx @drmahdikazempour/agent-stack init
# …or turn on EVERYTHING at once (max profile):
npx @drmahdikazempour/agent-stack init --all
What you'll see (click to expand)
agent-stack v0.2.0
Detected:
Host: Claude Code + Cursor
Repo: TypeScript / Next.js / pnpm
Profile: max (confidence: high)
Will write:
20 files → claude, cursor
.claude/settings.json (2 hooks merged)
Proceed? [Y/n] y
✓ Adapters: ccusage(installed)
✓ Generated 20 files
✓ Wired 2 hooks into settings.json
✓ All skills load, all hooks present, CLAUDE.md verified
✓ Built code map → .agent-stack/graph.md
✓ Baseline: 12,340 tokens/day (ccusage, last 7d avg)
Done.
Tip
Install it once globally to get the short agent-stack command everywhere:
npm i -g @drmahdikazempour/agent-stack
agent-stack init --all
⚙️ How init works
One shot, ten steps, fully reversible. --dry-run stops after the plan; --yes skips the single confirm.
detect ─▶ plan ─▶ confirm ─▶ back up ─▶ install ─▶ generate ─▶ wire hooks
│ │
--dry-run ┘ (stop) ▼
activate
summarize ◀── baseline ◀── code map ◀───────────────────────────┘ │
fails ──▶ roll back
#
Step
What happens
1
Detect
Host(s), repo type, framework, package manager, existing configs, git state
2
Plan
Prints a one-screen plan ( --dry-run stops here)
3
Confirm
A single Proceed? [Y/n] ( --yes skips)
4
Back up
Copies any existing config into .agent-stack.bak.<ts>/
5
Install
ccusage only — everything else is built in
6
Generate
CLAUDE.md , skills, subagents, commands, .claudeignore , Cursor mirror, MCP scaffold
7
Wire hooks
One merged write to .claude/settings.json (sole writer, dedupes, never clobbers yours)
8
Activate
Verifies each skill loads, each hook is present, CLAUDE.md exists — rolls back on failure
9
Code map
Builds the initial .agent-stack/graph.md
10
Baseline
Records a ccusage token snapshot for later comparison
🔧 Built-in token cutters
These ship inside the package — no external install, nothing fictional. They are what actually reduce tokens:
agent-stack graph refresh # rebuild .agent-stack/graph.md
agent-stack graph query < symbol > # find where it's defined
A compact index mapping every source file → its exported symbols. The agent greps one small file to find where something lives instead of reading whole directories. Refreshed automatically on SessionStart . Supports TypeScript/JavaScript, Python, Go, and Rust.
# Code map
_142 files, 906 top-level symbols. Grep this to find a symbol before opening source._
- `src/core/detect.ts`: detect
- `src/wire-hooks.ts`: mergeHooks, wireHooks, planHooks, countOurHooks
- …
🗜️ Output compression
npm run build 2>&1 | npx @drmahdikazempour/agent-stack compress
A stdin → stdout filter that strips ANSI codes, folds duplicate lines ( line (×42) ), and head/tail-elides huge output — ≈ 60 % fewer characters on a 500-line log , so noisy commands cost a fraction of the context.
🪶 Structural savings (always on)
.claudeignore keeps node_modules , build output, media, and lockfiles out of context.
CLAUDE.md is generated factual-and-tight — ≤ 800 tokens at startup, verified by doctor .
Subagents ( stack-explorer , stack-reviewer ) return conclusions, not file dumps .
Terse mode (the max profile) enforces minimal-word answers.
A profile bundles a graph backend + compression + skill set + hook config. init auto-picks one; swap later with profile use .
agent-stack profile use review # swap & regenerate
agent-stack profile show # current profile
📟 Command reference
agent-stack init [--all] [--yes] [--dry-run] [--targets claude,cursor]
[--profile < name > ] [--no-install] [--allow-noncommercial]
[--overwrite] [--force]
Token cutters — standalone, in pipes, or via hooks
agent-stack compress # cmd 2>&1 | agent-stack compress
agent-stack graph refresh # rebuild the code map
agent-stack graph query < term > # find a symbol / file
Maintenance — post-install, on demand
agent-stack audit # token counts + budget report
agent-stack optimize # apply audit fixes (with approval)
agent-stack doctor # lint everything (exit 1 on issues)
agent-stack measure [--since 7d] # ccusage baseline vs current
agent-stack profile use < name > # swap profile; regenerate
agent-stack graph use < name > # swap to an external graph (if installed)
agent-stack handoff write | resume # continuity across sessions
agent-stack sync # regenerate Cursor mirror from CLAUDE.md
agent-stack uninstall # restore backup, remove generated files
init flags in detail
Flag
Effect
--all
Full external stack at once (the max profile): rtk + code-review-graph + graphify + caveman + claude-handoff + gbrain
--yes
Skip the single confirm prompt
--dry-run
Print the plan, write nothing
--targets claude,cursor
Force the host list (Cursor gets the portable subset: rtk + MCP graph tools)
--profile <name>
Force a profile ( code review multimodal spec research max )
--no-install
Write configs only; print install guidance instead of installing
--overwrite
Replace existing files instead of merging (still backs up)
--force
Re-run even if already installed
🏗️ Architecture
One source of truth, two faces, zero runtime dependencies.
@drmahdikazempour/agent-stack
┌──────────────────────────────────────────────────────┐
│ CLI (src/) │
│ builtin/ graph (code map) · compress │
│ generate/ claude · cursor · mcp │
│ wire-hooks ← sole writer of settings.json │
│ activate ← verify, or roll back on failure │
│ skills/ 5 Agent Skills │
└───────────────┬──────────────────────┬────────────────┘
writes │ mirrors │
▼ ▼
🟠 Claude Code 🔵 Cursor
CLAUDE.md .cursor/rules/*.mdc
.claude/skills · agents AGENTS.md
commands · settings.json
.claudeignore · graph.md
Skills decide when ; the CLI decides how . Skills call the CLI under the hood; you normally run init once and never touch the CLI again.
agent-stack/
├── bin/agent-stack.js # npx entrypoint
├── src/
│ ├── cli.ts # arg parsing + command dispatch
│ ├── constants.ts # all spec values (token budgets, limits)
│ ├── core/ # detect · plan · safe-writer · backup · token estimator
│ ├── builtin/ # graph (code map) · compress (output compression)
│ ├── generate/ # claude · cursor · mcp · coordinator file builders
│ ├── adapters/ # registry · detect-tools · install · hooks
│ ├── wire-hooks.ts # SOLE writer of settings.json hooks
│ ├── activate.ts # post-write verification chain
│ ├── audit.ts # token-budget linting
│ └── commands/ # init + maintenance commands
├── skills/ # 5 Agent Skills (stack-bootstrap, -doctor, …)
├── integrations/ # profiles.json · tools.json
├── templates/ # generation notes
└── test/ # vitest: unit · golden · e2e init in a tmpdir
📉 How it cuts tokens
The savings come from changing what enters the context window , not from a black box:
❌ Before ✅ After (agent-stack)
───────────────────────── ─────────────────────────────────
read 12 files to find ──▶ grep graph.md → open 1 file
one function
paste a 500-line log ──▶ pipe thr

[truncated]
