---
source: "https://github.com/manuuuel/seh"
hn_url: "https://news.ycombinator.com/item?id=49253987"
title: "Se-harness – portable, tool-agnostic AI coding harness generator"
article_title: "GitHub - manuuuel/seh: Portable, tool-agnostic AI coding harness generator. One source of truth via AGENTS.md — a single global ruleset plus per-project, technology-aware guideline modules — with no vendor lock-in. · GitHub"
author: "manuuuel"
captured_at: "2026-08-11T06:48:36Z"
capture_tool: "hn-digest"
hn_id: 49253987
score: 1
comments: 0
posted_at: "2026-08-11T06:14:29Z"
tags:
  - hacker-news
  - translated
---

# Se-harness – portable, tool-agnostic AI coding harness generator

- HN: [49253987](https://news.ycombinator.com/item?id=49253987)
- Source: [github.com](https://github.com/manuuuel/seh)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T06:14:29Z

## Translation

タイトル: Se-harness – ポータブルでツールに依存しない AI コーディング ハーネス ジェネレーター
記事のタイトル: GitHub - manuuuel/seh: ポータブルでツールに依存しない AI コーディング ハーネス ジェネレーター。 AGENTS.md を介した 1 つの信頼できる情報源 (単一のグローバル ルールセットに加えて、プロジェクトごとのテクノロジを意識したガイドライン モジュール) はベンダー ロックインがありません。 · GitHub
説明: ポータブルでツールに依存しない AI コーディング ハーネス ジェネレーター。 AGENTS.md を介した 1 つの信頼できる情報源 (単一のグローバル ルールセットに加えて、プロジェクトごとのテクノロジを意識したガイドライン モジュール) はベンダー ロックインがありません。 - マヌエル/セー

記事本文:
GitHub - manuuuel/seh: ポータブルでツールに依存しない AI コーディング ハーネス ジェネレーター。 AGENTS.md を介した 1 つの信頼できる情報源 (単一のグローバル ルールセットに加えて、プロジェクトごとのテクノロジを意識したガイドライン モジュール) はベンダー ロックインがありません。 · GitHub
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
マニュアル
/
ああ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
93 コミット 93 コミット .claude .claude .github .github .seh .seh アセット アセット dist

dist スクリプト スクリプト src src test test .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json seh.lock seh.lock tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ポータブルでツールに依存しない AI コーディング ハーネス ジェネレーター。真実の情報源の 1 つは、
AGENTS.md をエントリポイントとして使用し、ベンダー ロックインはありません。
カール -fsSL https://raw.githubusercontent.com/manuuuel/seh/main/scripts/install.sh |しー
# マシンごとに 1 回: 統合されたグローバル ルールセット + エージェントのシンボリックリンク
seh init --global --agents クロード、コーデックス --yes
# プロジェクトごと: スタックの検出、足場 AGENTS.md + .seh/
プロジェクトの cd
seh init --tech typescript --yes
# .seh/ ソースを編集した後に生成されたファイルを再生成する
同期してる
これにより、 ~/.seh/AGENTS.md グローバル ルールセットがシンボリックリンクされます。
エージェントの構成パスとプロジェクト AGENTS.md (および CLAUDE.md 、
GEMINI.md , …) サポートされているすべてのエージェントが自動的に読み取ります。
seh は、AI コーディング エージェント (Claude Code、Codex、
Gemini、Pi、OpenCode、Copilot、…) 読んでください: 単一のグローバル ルールセット
マシンに加えて、焦点を絞ったテクノロジー固有のプロジェクトごとのインデックス
ガイドラインモジュール。ハーネス パッケージのスキルはすべてのユーザーに配布されます。
彼らをサポートするエージェント。信頼できる 1 つの情報源を編集し、 seh sync を実行し、
すべてのツール固有のファイルはロックステップで再生成されます。もうコピー＆ペーストする必要はありません。
同じルールを CLAUDE.md 、 GEMINI.md 、および .github/copilot-instructions.md に追加します。
別に。
レイヤー
場所
それは何ですか
L0 — コア
CLIにバンドルされている
作成されたソース コンテンツ (グローバル セクション + テクノロジーごとのカタログ)。アクティブなパッケージがない場合のフォールバック。
L1 — グローバル
~/.seh/
横断的なルールを備えた単一の統合された AGENTS.md。アウ

マシンごとに 1 回のスレッド。オプションのエージェントのシンボリックリンク。リポジトリにはコピーされません。
L2 — プロジェクト
<リポジトリ>/AGENTS.md + .seh/
プロジェクト + テクノロジーごとのモジュールをリンクする薄いインデックス。リポジトリにコミットします。
パッケージ
<あなたのハーネス>/
グローバル ルール、スタック モジュール、テンプレート、スキルのバージョン管理された Git リポジトリ。アクティブな場合は、L0/L1 よりも優先されます。
任意のファイルの解決順序: package → ~/.seh/ → seh バンドルコア 。
グローバルルールはどこにでも適用されます。プロジェクトのレイヤーは拡張されます。決して矛盾しないようにします。
生成されたファイルは手動で編集しないでください。 .seh/ ソースを編集して実行します
同期して再生成してください。
グローバル ( ~/.seh/AGENTS.md ) は 1 つの自己完結型ファイル (すべてのガードレール)
強制的なクラフトマンシップの原則に基づいて単一のドキュメントにインライン化される
(小さくシャープに保ち、エレガントなコードを書き、最小限のものを追求してください)
デザインがより良く機能し、スロップが発生しません）。グローバルを自動ロードするツール
指示ファイルはルールセット全体を直接取得します。
プロジェクト ( <repo>/.seh/AGENTS.md ) は正規ファイルです。
ディレクティブ プリアンブルとプログレッシブ インデックス (リンクされた目次) のポイント
.seh/ の下のフォーカスされたモジュールで、オンデマンドでロードされます。
.seh/project.md — ミッション、制約、範囲外 (リードファースト)
.seh/domain/*.md — アーキテクチャ、用語集
.seh/stack/<tech>.md — テクノロジーごとのベスト プラクティス
ツール固有のファイル ( AGENTS.md 、 CLAUDE.md 、 GEMINI.md 、
.github/copilot-instructions.md など) へのシンボリックリンクが生成されます。
.seh/AGENTS.md 、gitignored され、 seh sync によって再生成されます。
1. グローバルセットアップ (マシンごとに 1 回)
seh init --global # インタラクティブ: シンボリックリンクするエージェントを選択します
seh init --global --agents claude,codex --yes # 非対話型
作成するもの:
~/.seh/AGENTS.md — 統一されたグローバル ルールセット (最初に職人技、次に
セキュリティ、品質ゲート、テスト、コミット、分岐、依存関係、エラー
ハンドリング、オブ

保守性、データとプライバシー、ドキュメント、リファクタリング、ワークフロー、
セッションの開始、レポート、境界、コード原則)。
~/.seh/config.json — どのエージェントがシンボリックリンクされているか。
必要に応じて、エージェントを接続して自動ロードします (seh リンクを参照)。
プロジェクトの cd
seh init # テクノロジー、対話型複数選択を検出します
seh init --tech typescript,python --yes # 非対話型 (1 つ以上が必要)
作成するもの:
AGENTS.md — プロジェクト インデックス (以下のモジュールへのリンク)
.seh/project.md 、 .seh/domain/architecture.md 、 .seh/domain/glossary.md
選択したテクノロジーごとに .seh/stack/<tech>.md
seh.lock — 選択したテクノロジーを記録します (コミットします)
サポートされているテクノロジ: javascript 、 typescript 、 python 、 go 、 c 、
サビ、ジャワ。一般的なフォールバックはありません。少なくとも 1 つを選択してください。
次に、.seh/project.md と .seh/domain/* を入力して再同期します。
3. 同期 — ソースから再生成
同期してる
プロジェクト AGENTS.md インデックスと .seh/stack/* を seh.lock から書き換えます。
冪等 (ソースを変更せずに再実行しても変更なし)。
チェックしてください
生成されたファイルがソースと一致する場合は 0 を終了します。出口 1 (メッセージ付き) の場合
AGENTS.md またはスタック モジュールが古いか、見つかりません。プリコミット/CIに適しています。
5. リンク — エージェントのシンボリックリンクを管理する
seh link --add claude # symlink ~/.claude/CLAUDE.md -> ~/.seh/AGENTS.md
seh link --remove claude # 削除します
AGENTS.md は唯一の信頼できる情報源であり続けます。シンボリックリンクは純粋なポインタであるため、
エージェントごとにドリフトするコンテンツはありません。
サポートされているエージェント: claude 、 codex 、 pi 、 gemini 、 opencode 、 copilot 、 エージェント
エージェント ターゲットは、エージェント間の相互運用性パス ( ~/.agents/ 、
.agents/ ) Gemini CLI、Pi、Copilot などによってエイリアスとして使用されます。
エージェント
ターゲット
クロード
~/.claude/CLAUDE.md
コーデックス
~/.codex/AGENTS.md
円周率
~/.pi/agent/AGENTS.md
ジェミニ
~/.gemini/GEMINI.md
オープンコード
~/.config/opencode/AGENTS.md
副操縦士
~/.c

opilot/copilot-instructions.md
エージェント
~/.agents/AGENTS.md
プロジェクトの目標
エージェント
ターゲット
正規
クロード
クロード.md
.seh/AGENTS.md
コーデックス
エージェント.md
.seh/AGENTS.md
円周率
エージェント.md
.seh/AGENTS.md
ジェミニ
ジェミニ.md
.seh/AGENTS.md
オープンコード
エージェント.md
.seh/AGENTS.md
副操縦士
.github/copilot-instructions.md
.seh/AGENTS.md
エージェント
.agents/AGENTS.md
.seh/AGENTS.md
ハーネスパッケージ
ハーネス パッケージは、ユーザーが git でバージョン管理するプレーンなディレクトリです。
マシン間で共有し、チームメイトと共有します。足場を組んでそれを読みます。 git
完全に外部であり、ラップされることはありません。
私のハーネス/
§── harness.json — パッケージのメタデータ (名前、バージョン、モデルタグ)
§── CHANGELOG.md — 意思決定と推論を利用する (人間が作成)
§── グローバル/
│ §── AGENTS.md — グローバル ルールセット (~/.seh/AGENTS.md を置き換えます)
│ └── config.json — エージェントのシンボリックリンク設定
§── テンプレート/
│ §── stack/ — テクノロジーごとの構造パターン
│ └── project/ — プロジェクト全体の足場
§── プロジェクト/ — リポジ名に一致するリポジトリごとのオーバーレイ
└── スキル/ — 配布するスキル (販売または参照)
§── ブレーンストーミング/ — ベンダー スキル (パッケージ リポジトリにコミット)
└── caveman/ — 参照されたスキル (インストール時に取得)
パッケージコマンド
seh パッケージ初期化 [パス]
新しいハーネス パッケージをパス (デフォルト: ./my-harness ) にスキャフォールディングします。作成します
完全なディレクトリ構造、バンドルされたスタック モジュールを開始点としてコピーします。
harness.json と空の CHANGELOG.md を書き込みます。
seh package init ~ /my-harness
cd ~ /my-harness && git init && git add 。 && git commit -m " init ハーネス "
seh パッケージは <パス> を使用します
既存のパッケージを指します。 packagePath を ~/.seh/config.json に書き込みます。
git clone git@github.com:you/my-harness.git ~ /my-harness
seh パッケージの使用 ~ /my-harness
パッケージ

ステータス
現在アクティブなパッケージのパス、 harness.json の名前/バージョン、および
予想される各ディレクトリが存在するかどうか。
アクティブなパッケージからホスト マシンにアーティファクトをインストールします。
seh package install --harness # write ~/.seh/AGENTS.md + エージェントのシンボリックリンク
seh package install --skills # ~/.seh/skills/ + エージェント ディレクトリへのシンボリックリンク スキル
seh package install --all # 上記の両方
seh package install --all --agents claude,gemini # 非対話型エージェントの選択
旗
アクション
--ハーネス
package/global/AGENTS.md から ~/.seh/AGENTS.md を書き込みます。エージェントのシンボリックリンクを更新します
--スキル
参照されたスキル、シンボリックリンクを取得します package/skills/<name>/ → ~/.seh/skills/<name>/ → エージェント スキル ディレクトリ
--すべて
上記の両方
--agents <リスト>
スキルのシンボリックリンクを受信するカンマ区切りのエージェント (省略した場合は対話的にプロンプトが表示されます)
--force
既存のファイルを上書きする
新しいマシンのワークフロー:
git clone git@github.com:you/my-harness.git ~ /my-harness
seh パッケージの使用 ~ /my-harness
seh パッケージのインストール --all
スキル
スキル (再利用可能な SKILL.md ベースの機能パッケージ) はファーストクラスのレイヤーです
ハーネスパッケージ内。 seh はそれらをパッケージからすべてのエージェントに配布します
スキルディレクトリをサポートします。
<package>/skills/<name>/ ← ソース (ベンダーまたは参照)
↓ シンボリックリンク
~/.seh/skills/<name>/ ← このマシンの安定した中間
↓ シンボリックリンク
~/.claude/skills/<名前>/ ← クロードコード
~/.codex/skills/<名前>/ ← Codex CLI
~/.gemini/skills/<名前>/ ← Gemini CLI
~/.config/opencode/skills/<名前>/ ← OpenCode
~/.pi/agent/skills/<名前>/ ← 円周率
~/.copilot/skills/<名前>/ ← GitHub コパイロット
~/.agents/skills/<name>/ ← エージェント間の相互運用性パス
~/.seh/skills/ は安定した中間です。パッケージが移動すると、エージェントは
シンボリックリンクは、再ポイントされるまでそのまま残ります。
seh スキル add <url> [--vendor | --参照] [--参照 <br

anch>] [ルーティングフラグ]
GitHub URL からアクティブなパッケージにスキルを追加します。
SEH スキルを追加 https://github.com/JuliusBrussee/caveman --reference
SEH スキルを追加 github:you/my-skill --vendor --ref v1.2
--vendor : ファイルを <package>/skills/<name>/ にクローンします (git にコミット)
--reference : harness.json のみに URL を記録します。追加します
skill/<name>/ を .gitignore に変換します。ファイルはインストール時に取得されます
どちらのフラグも指定されていない場合は、タイプの入力を求めるプロンプトが表示されます
リポジトリ名からスキル名を推測します
ルーティング フラグは、エージェントにスキルをいつ呼び出すかを指示します。 1 つだけを使用できます。
SEH スキルは github:you/caveman --vendor --always "すべての応答" を追加します
SEH スキルは github:you/systematic-debugging --vendor --when " バグ / テスト失敗 / 予期せぬ動作 " を追加します
SEH スキルは github:you/xlsx --vendor --optional を追加します
ルーティングは harness.json に保存され、## Skills セクションにレンダリングされます。
プロジェクト AGENTS.md (seh sync による) とグローバル ~/.seh/AGENTS.md の両方
( seh package install --harness により) そのため、エージェントはどのスキルを使用すべきかを常に把握できます。
いつ呼び出すか。
参照されたスキルをソースから再フェッチします。引数がない場合はすべてを更新します
参考にしたスキル。指定されたスキルがベンダー化されている場合はエラー。
アクティブなパッケージ内のスキルとそのルーティング モードを示します。
✓ [ベンダー] は常に、実装前にブレインストーミングを行う
✓ 体系的なデバッグ [ベンダー] 次の場合: バグ / テスト失敗
✗穴居人 【参考】https://github.com/JuliusBrussee/c

[切り捨てられた]

## Original Extract

Portable, tool-agnostic AI coding harness generator. One source of truth via AGENTS.md — a single global ruleset plus per-project, technology-aware guideline modules — with no vendor lock-in. - manuuuel/seh

GitHub - manuuuel/seh: Portable, tool-agnostic AI coding harness generator. One source of truth via AGENTS.md — a single global ruleset plus per-project, technology-aware guideline modules — with no vendor lock-in. · GitHub
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
manuuuel
/
seh
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
93 Commits 93 Commits .claude .claude .github .github .seh .seh assets assets dist dist scripts scripts src src test test .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json seh.lock seh.lock tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
Portable, tool-agnostic AI coding harness generator. One source of truth,
AGENTS.md as the entrypoint, no vendor lock-in.
curl -fsSL https://raw.githubusercontent.com/manuuuel/seh/main/scripts/install.sh | sh
# Once per machine: unified global ruleset + agent symlinks
seh init --global --agents claude,codex --yes
# Per project: detect stack, scaffold AGENTS.md + .seh/
cd your-project
seh init --tech typescript --yes
# Regenerate generated files after editing .seh/ sources
seh sync
This gives you a ~/.seh/AGENTS.md global ruleset symlinked into your
agents' config paths, plus a project AGENTS.md (and CLAUDE.md ,
GEMINI.md , …) that every supported agent reads automatically.
seh produces the context files that AI coding agents (Claude Code, Codex,
Gemini, Pi, OpenCode, Copilot, …) read: a single global ruleset on your
machine plus a per-project index of focused, technology-specific
guideline modules. Skills from a harness package are distributed to every
agent that supports them. Edit one source of truth, run seh sync , and
every tool-specific file regenerates in lockstep — no more copy-pasting the
same rules into CLAUDE.md , GEMINI.md , and .github/copilot-instructions.md
separately.
Layer
Location
What it is
L0 — Core
bundled in the CLI
The authored source content (global sections + per-technology catalog). Fallback when no package is active.
L1 — Global
~/.seh/
A single unified AGENTS.md with your cross-cutting rules. Authored once per machine; optional agent symlinks. Not copied into repos.
L2 — Project
<repo>/AGENTS.md + .seh/
A thin index linking project + per-technology modules. Committed to the repo.
Package
<your-harness>/
A versioned git repo of global rules, stack modules, templates, and skills. Takes precedence over L0/L1 when active.
Resolution order for any file: package → ~/.seh/ → seh bundled core .
Global rules apply everywhere; project layers extend — never contradict them.
Generated files must NOT be hand-edited. Edit the .seh/ sources and run
seh sync to regenerate.
Global ( ~/.seh/AGENTS.md ) is one self-contained file — every guardrail
inlined in a single document, led by a forced Craftsmanship principle
(keep it small and sharp, write elegant code, seek the most minimal
better-working design, introduce no slop). Tools that auto-load a global
instructions file get the whole ruleset directly.
Project ( <repo>/.seh/AGENTS.md ) is the canonical file — a short
directive preamble plus a progressive index (linked table of contents) pointing
at focused modules under .seh/ , loaded on demand:
.seh/project.md — mission, constraints, out-of-scope (read-first)
.seh/domain/*.md — architecture, glossary
.seh/stack/<tech>.md — per-technology best practices
Tool-specific files ( AGENTS.md , CLAUDE.md , GEMINI.md ,
.github/copilot-instructions.md , etc.) are generated symlinks to
.seh/AGENTS.md , gitignored, and regenerated by seh sync .
1. Global setup (once per machine)
seh init --global # interactive: choose which agents to symlink
seh init --global --agents claude,codex --yes # non-interactive
Creates:
~/.seh/AGENTS.md — the unified global ruleset (Craftsmanship first, then
security, quality gates, testing, commits, branching, dependencies, error
handling, observability, data & privacy, documentation, refactoring, workflow,
session startup, reporting, boundaries, code principles).
~/.seh/config.json — which agents are symlinked.
Optionally wire agents to auto-load it (see seh link ).
cd your-project
seh init # detects technologies, interactive multi-select
seh init --tech typescript,python --yes # non-interactive (≥1 required)
Creates:
AGENTS.md — the project index (links to the modules below)
.seh/project.md , .seh/domain/architecture.md , .seh/domain/glossary.md
.seh/stack/<tech>.md for each selected technology
seh.lock — records the selected technologies (commit it)
Supported technologies: javascript , typescript , python , go , c ,
rust , java . No generic fallback — pick at least one.
Then fill in .seh/project.md and .seh/domain/* and re-sync.
3. Sync — regenerate from sources
seh sync
Rewrites the project AGENTS.md index and .seh/stack/* from seh.lock .
Idempotent (no change on re-run with unchanged sources).
seh check
Exit 0 if the generated files match the sources; exit 1 (with a message) if
AGENTS.md or a stack module is stale or missing. Suitable for pre-commit/CI.
5. Link — manage agent symlinks
seh link --add claude # symlink ~/.claude/CLAUDE.md -> ~/.seh/AGENTS.md
seh link --remove claude # remove it
AGENTS.md stays the single source of truth; symlinks are pure pointers, so
there is no per-agent content to drift.
Supported agents: claude , codex , pi , gemini , opencode , copilot , agents
The agents target manages the cross-agent interoperability path ( ~/.agents/ ,
.agents/ ) used as an alias by Gemini CLI, Pi, Copilot, and others.
Agent
Target
claude
~/.claude/CLAUDE.md
codex
~/.codex/AGENTS.md
pi
~/.pi/agent/AGENTS.md
gemini
~/.gemini/GEMINI.md
opencode
~/.config/opencode/AGENTS.md
copilot
~/.copilot/copilot-instructions.md
agents
~/.agents/AGENTS.md
Project targets
Agent
Target
Canonical
claude
CLAUDE.md
.seh/AGENTS.md
codex
AGENTS.md
.seh/AGENTS.md
pi
AGENTS.md
.seh/AGENTS.md
gemini
GEMINI.md
.seh/AGENTS.md
opencode
AGENTS.md
.seh/AGENTS.md
copilot
.github/copilot-instructions.md
.seh/AGENTS.md
agents
.agents/AGENTS.md
.seh/AGENTS.md
Harness Packages
A harness package is a plain directory the user versions with git, carries
between machines, and shares with teammates. seh scaffolds and reads it; git
is fully external and never wrapped.
my-harness/
├── harness.json — package metadata (name, version, modelTag)
├── CHANGELOG.md — harness decisions and reasoning (human-authored)
├── global/
│ ├── AGENTS.md — global ruleset (replaces ~/.seh/AGENTS.md)
│ └── config.json — agent symlink config
├── templates/
│ ├── stack/ — per-technology structural patterns
│ └── project/ — full project scaffolds
├── projects/ — per-repo overlays matched by repo name
└── skills/ — skills to distribute (vendored or referenced)
├── brainstorming/ — vendored skill (committed to package repo)
└── caveman/ — referenced skill (fetched on install)
Package commands
seh package init [path]
Scaffolds a new harness package at path (default: ./my-harness ). Creates
the full directory structure, copies bundled stack modules as a starting point,
writes harness.json and an empty CHANGELOG.md .
seh package init ~ /my-harness
cd ~ /my-harness && git init && git add . && git commit -m " init harness "
seh package use <path>
Points seh at an existing package. Writes packagePath into ~/.seh/config.json .
git clone git@github.com:you/my-harness.git ~ /my-harness
seh package use ~ /my-harness
seh package status
Shows the currently active package path, name/version from harness.json , and
whether each expected directory exists.
Installs artifacts from the active package onto the host machine.
seh package install --harness # write ~/.seh/AGENTS.md + agent symlinks
seh package install --skills # symlink skills into ~/.seh/skills/ + agent dirs
seh package install --all # both of the above
seh package install --all --agents claude,gemini # non-interactive agent selection
Flag
Action
--harness
Writes ~/.seh/AGENTS.md from package/global/AGENTS.md ; updates agent symlinks
--skills
Fetches referenced skills, symlinks package/skills/<name>/ → ~/.seh/skills/<name>/ → agent skill dirs
--all
Both of the above
--agents <list>
Comma-separated agents to receive skill symlinks (prompts interactively if omitted)
--force
Overwrite existing files
New machine workflow:
git clone git@github.com:you/my-harness.git ~ /my-harness
seh package use ~ /my-harness
seh package install --all
Skills
Skills (reusable SKILL.md -based capability packages) are a first-class layer
in the harness package. seh distributes them from the package to every agent
that supports a skill directory.
<package>/skills/<name>/ ← source (vendored or referenced)
↓ symlink
~/.seh/skills/<name>/ ← stable intermediate on this machine
↓ symlinks
~/.claude/skills/<name>/ ← Claude Code
~/.codex/skills/<name>/ ← Codex CLI
~/.gemini/skills/<name>/ ← Gemini CLI
~/.config/opencode/skills/<name>/ ← OpenCode
~/.pi/agent/skills/<name>/ ← Pi
~/.copilot/skills/<name>/ ← GitHub Copilot
~/.agents/skills/<name>/ ← cross-agent interoperability path
~/.seh/skills/ is the stable intermediate: if the package moves, agent
symlinks remain intact until re-pointed.
seh skills add <url> [--vendor | --reference] [--ref <branch>] [routing flag]
Adds a skill from a GitHub URL to the active package.
seh skills add https://github.com/JuliusBrussee/caveman --reference
seh skills add github:you/my-skill --vendor --ref v1.2
--vendor : clones files into <package>/skills/<name>/ (committed to git)
--reference : records the URL in harness.json only; appends
skills/<name>/ to .gitignore ; files are fetched at install time
Prompts for type if neither flag is given
Infers skill name from the repo name
Routing flags tell agents when to invoke the skill. Exactly one may be used:
seh skills add github:you/caveman --vendor --always " every response "
seh skills add github:you/systematic-debugging --vendor --when " bug / test failure / unexpected behavior "
seh skills add github:you/xlsx --vendor --optional
Routing is stored in harness.json and rendered into a ## Skills section in
both the project AGENTS.md (by seh sync ) and the global ~/.seh/AGENTS.md
(by seh package install --harness ), so agents always know which skills to
invoke and when.
Re-fetches referenced skill(s) from their source. No arguments updates all
referenced skills. Errors if the named skill is vendored.
Shows skills in the active package with their routing mode:
✓ brainstorming [vendor] always: before any implementation
✓ systematic-debugging [vendor] when: bug / test failure
✗ caveman [reference] https://github.com/JuliusBrussee/c

[truncated]
