---
source: "https://github.com/mohitjandwani/analyst-kit"
hn_url: "https://news.ycombinator.com/item?id=48621014"
title: "Analyst Kit (YC W23): Turn your Claude / Codex into an investment analyst (Free)"
article_title: "GitHub - mohitjandwani/analyst-kit · GitHub"
author: "mohitjandwani"
captured_at: "2026-06-21T18:29:15Z"
capture_tool: "hn-digest"
hn_id: 48621014
score: 2
comments: 3
posted_at: "2026-06-21T17:58:40Z"
tags:
  - hacker-news
  - translated
---

# Analyst Kit (YC W23): Turn your Claude / Codex into an investment analyst (Free)

- HN: [48621014](https://news.ycombinator.com/item?id=48621014)
- Source: [github.com](https://github.com/mohitjandwani/analyst-kit)
- Score: 2
- Comments: 3
- Posted: 2026-06-21T17:58:40Z

## Translation

タイトル: アナリスト キット (YC W23): クロード/コーデックスを投資アナリストに変身させます (無料)
記事タイトル: GitHub - mohitjandwani/analyst-kit · GitHub
説明: GitHub でアカウントを作成して、mohitjandwani/analyst-kit の開発に貢献します。

記事本文:
GitHub - mohitjandwani/analyst-kit · GitHub
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
モヒチャンドワニ
/
アナリストキット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く

フォルダーとファイル
86 コミット 86 コミット .claude-plugin .claude-plugin .devcontainer .devcontainer .github/ workflows .github/ workflows bin bin plan plan plugins plugins scripts scripts skill skill src src test test tools/ ta-tester tools/ ta-tester .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md README.md README.md VERSION VERSION COMPATIATION.MD COMPATIATION.MD package.json package.json registry.json registry.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェント向けのインストール可能なヘッジファンド グレードの株式調査スキル。
各スキルは、命令が含まれた自己完結型のフォルダーです (便利な場合は実行可能です)。
スクリプト）をエージェントがオンデマンドでロードします。これらをプラグインとして Claude Code にインストールします。
または、バンドルされたインストーラーを使用してエージェント ランタイムにコピーします。
スキルのフロントマターは唯一の真実の情報源です - レジストリ、プラグイン
マニフェストとインストーラーはすべてマニフェストから派生します。
スキルは機能 (1 つのアトミックなジョブ、データ ソース、エンジン、
成果物、または再利用可能なナレッジ）とワークフロー（エンゲージメントのエントリ ポイント）
これは、requires: ) を介して機能を調整します。 「ニーズ」列にはランタイムがリストされます。
API キーと必要なスキル。
1 つのコマンド — macOS、Linux、および Windows で動作します (WSL2 内 — 以下を参照)。それ
すべてのスキルを選択したランタイムにインストールし、エージェントのシステム/共通プロンプトに接続します。
ノード 18 以降のみが必要です (OS が検出され、正しいパスにインストールされます)。
npx github:mohitjandwani/analyst-kit claude-code # または: codex · openclaw · cowork
クロードコードを codex 、 openclaw 、または cowork に交換します。現在のプロジェクトにインストールする --scope プロジェクトを追加します
ホーム ディレクトリの代わりにプロジェクト ( ./.claude/skills , …) を使用します。すでにリポジトリのクローンを作成しましたか?ノード bin/analyst-kit.js クロードコードは行います

同じです (さらに list 、 Doctor 、 uninstall 、または install <skill|persona> を 1 つだけ追加します)。
Claude Cowork の場合、コマンドはアプリ内ステップを出力し、貼り付けるために cowork-global-instructions.md を書き込みます。
「設定」→「Cowork」→「グローバル指示」に移動 — Cowork はプラグインを通じてスキル自体をインストールします
マーケットプレイス（下）。
Windows の場合: WSL2 内で実行します。ネイティブ Windows (PowerShell/cmd) は、スキルがサポートされていないためサポートされていません。
ランタイムはPOSIX/bashです。以下の「Windows: WSL2 を使用する」を参照してください。
マーケットプレイス プラグイン — クローンなし、ノードなし (Claude Code & Cowork)
Claude Code と Claude Cowork (Anthropic のデスクトップ アプリ) の両方
同じプラグイン マーケットプレイスからインストールします。
クロードコード:
/プラグイン マーケットプレイス追加 mohitjandwani/analyst-kit
/plugin install us-stock-analyst@analyst-kit # または international-analyst / taiwan-stock-analyst
Claude Cowork (デスクトップ アプリ): カスタマイズ → プラグイン → 個人用プラグイン → + → マーケットプレイスの追加 →
mohitjandwani/analyst-kit 、 us-stock-analyst プラグインを追加し、設定 → を有効にします
機能 → コード実行。
インストール後、トリガーフレーズ (例: 「NVDA の詳細」) を尋ねると、一致するスキルがロードされます。クローンから
すべてのプラットフォームでインストーラーをセルフテストすることもできます。
npm run test:integration プラットフォームごとの実際のインストール数 + path.win32 パス層チェック
Codex (任意の OS、ChatGPT ログインは必要ありません) — API キーを使用してスキルにアクセスできることを確認します。
CODEX_API_KEY=sk-... codex exec --json " sec-filings スキルを使用して NVDA の最新の 8-K をリストします "
各ランタイムが下で何を行うかについては、compatibility.md を参照してください。スキルがどこに到達するか、
ルーティング テーブル、および Windows の仕様。
Analyst Kit のスキルは POSIX/bash ランタイム (analyst-kit-core) を実行するため、Windows では
WSL2 内でのみサポートされます。これは、エージェント自体が要求するものと一致します。
Claude Code のサンドボックスは、macOS、Linux、

d WSL2 (ネイティブ Windows は
未サポート)、Codex の Linux モードも WSL2 です。
推奨: WSL2 をインストールし、
次に、WSL2 内で Claude Code (または Codex) とこのインストーラーを実行します。
ディストリビューション - すべてが Linux とまったく同じように動作します。
ネイティブ Windows (PowerShell/cmd) は bash ランタイムをまったく実行できません。ネイティブ
Windows + Git Bash はスクリプトを実行しますが、.env ファイルは強制されません
アクセス許可 (chmod は NTFS では何も操作しません)、サンドボックスなしではサポートされません。
ネイティブ Windows で実行すると、node bin/analyst-kit.js Doctor --platform claude-code が警告を発します。
3 つのペルソナ プラグインは、さまざまな市場のリサーチ ワークフローをバンドルします。すべて
3 つは調査ワークフロー (詳細分析、テーマ分析、技術分析、
企業 Wiki) とそのサポート機能 (グラフ作成、レポート作成、
ウィキビルダー、企業ユニバースマネージャー、財務モデリング準備、市場インテリジェンス、
財務諸表の分析、財務モデルの作成、データ分析）。の
市場の違いは、FinMind (台湾のデータ) と SEC の申請が含まれるかどうかです。
プラグインの正確な内容を確認するには、node bin/analyst-kit.js list --persona <name> を実行します。
キーは環境または git-ignored .env から読み取られます。インストーラ
(analyst-kit env /analyst-kit install) は、対話的に実行すると、不足しているものがあればプロンプトを表示します。
.env.example を参照してください。
13f 分析にはキーは必要ありません。SEC EDGAR を直接読み取ります (オプションの
SEC_EDGAR_UA SEC フェアアクセスの儀礼としての連絡先文字列)。
コードを実行するスキルは、最初の使用時に独自の依存関係をブートストラップします。ランタイムとは、
スキルごとの前提条件はインストーラーによってインストールされません: Python (finmind,
会社の世界のマネージャー。 13f 分析は標準ライブラリのみ) および Bun
(ウィキビルダー)。
~/.analyst-kit データ ホーム、分析、更新
すべてのスキルは共有ランタイム (analyst-kit-core、install) 上で実行されます。

として自動的に編集されます
依存関係）、すべてのユーザーごとの状態を 1 つの固定場所 ~/.analyst-kit/ に保持します。
.env — API キー (chmod 600、プロジェクト間で共有)
config — 設定 (analyst-kit-core/bin/analyst-kit-config get|set|list )
Analytics/skill-usage.jsonl — ローカル使用状況ログ: どのスキルがいつ実行されたか、
結果、期間
learnings.jsonl — セットアップや設定についてスキルが学んだこと、
間違いが繰り返されないように
テレメトリはデフォルトでオンになっています (匿名、オプトアウト)。スキル名のみを送信しますが、
バージョン、結果、期間 — 名前、ファイル パス、ティッカー、コンテンツをリポジトリにしないでください —
これにより、どのスキルが壊れたり動作が遅くなったりするのかがわかるので、それを直接オンにし続けます
エクスペリエンスが向上します。最初の実行時に一度それについて通知されます。階層: コミュニティ
(デフォルト、安定した匿名 ID)、匿名 (ID なし)、オフ。いつでもオプトアウト:
~ /.claude/skills/analyst-kit-core/bin/analyst-kit-config テレメトリをオフに設定します
更新: スキルは公開バージョンを最大 1 日に 1 回チェックし、
新しいリリースがリリースされたときのガイド付きアップグレード (拒否すると 1 週間保留されます。無効にします)
Analyst-kit-config では update_check false を設定します)。
各スキルは、YAML フロントマターを含む skill/<name>/SKILL.md です。
---
名前: 13f-分析 # ケバブケース;フォルダ名と等しくなければなりません
タイプ: 能力 # 能力 |ワークフロー
description : > # 動作内容 + トリガーフレーズの「Triggers:」句
... トリガー: 「X の 13F を入手」、「<fund> は何を所有しているか」、...
Required : [ ... ] # これが構築する能力スキル (ワークフローを必要とするものは何もないかもしれません)
env : [ FMP_API_KEY ] # スキルに必要な API キー
---
エージェントは name + description のみを読み取るため、トリガー フレーズは
説明。 type 、 require 、および env はインストーラーとバリデーターを駆動します。
npm run validate # lint スキル + プラグイン マニフェスト (+ プリアンブル同期チェック)
npm 実行ビルド

ld:registry # フロントマターから registry.json を再生成
npm run check:registry # registry.json が同期していることを確認します
npm run sync:preamble # すべての SKILL.md 内のanalyst-kit-core ブロックを再生成します
registry.json が生成されます。スキル フロントマターを編集して再構築します。の
各 SKILL.md 内の <!-- Analyst-kit:preamble/epilogue --> ブロックも生成されます — 編集
skill/analyst-kit-core/templates/ と再同期;マーカー間は決して編集しないでください。同じ
プッシュおよびプル リクエストごとにチェックが CI ( .github/workflows/validate.yml ) で実行されます。
計画されたスキル、まだ利用可能ではありません: LBO モデル (借金スケジュール、資金調達、
IRR/MOIC) および PDF レポート分析。
財務諸表の分析と DCF + 感度ツール
財務モデルの作成はカスタム財務スキルからインスピレーションを受けました
Anthropic のクロード料理本では、
その後、このリポジトリのスキル コントラクト (フロントマター、スクリプト/レイアウト) に修正されました。
入力ガードとテストスイートで強化されています。
財務モデル作成における M&A の増加/希薄化モデルは次のとおりです。
joe-neary/MergerDealSimulator からインスピレーションを受けました。
財務計算式は最初から再実装されました (コードはコピーされませんでした)。
そのプロジェクトの動作例は、
スキルのテストスイート。
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

Contribute to mohitjandwani/analyst-kit development by creating an account on GitHub.

GitHub - mohitjandwani/analyst-kit · GitHub
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
mohitjandwani
/
analyst-kit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
86 Commits 86 Commits .claude-plugin .claude-plugin .devcontainer .devcontainer .github/ workflows .github/ workflows bin bin plans plans plugins plugins scripts scripts skills skills src src test test tools/ ta-tester tools/ ta-tester .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md README.md README.md VERSION VERSION compatibility.md compatibility.md package.json package.json registry.json registry.json View all files Repository files navigation
Installable, hedge-fund-grade equity-research skills for AI coding agents.
Each skill is a self-contained folder of instructions (and, where useful, runnable
scripts) that an agent loads on demand. Install them into Claude Code as a plugin,
or copy them into any agent runtime with the bundled installer.
The skill frontmatter is the single source of truth — the registry, the plugin
manifests, and the installer all derive from it.
The skills split into capabilities (one atomic job — a data source, an engine,
a deliverable, or reusable knowledge) and workflows (an engagement entry point
that orchestrates capabilities via requires: ). The "Needs" column lists runtimes,
API keys, and required skills.
One command — works on macOS, Linux, and Windows (inside WSL2 — see below ). It
installs all the skills into your chosen runtime and wires them into the agent's system/common prompt.
Needs only Node ≥ 18 (which detects your OS and installs to the right paths):
npx github:mohitjandwani/analyst-kit claude-code # or: codex · openclaw · cowork
Swap claude-code for codex , openclaw , or cowork ; add --scope project to install into the current
project ( ./.claude/skills , …) instead of your home directory. Already cloned the repo? node bin/analyst-kit.js claude-code does the same (plus list , doctor , uninstall , or install <skill|persona> for just one).
For Claude Cowork , the command prints the in-app steps and writes cowork-global-instructions.md to paste
into Settings → Cowork → Global instructions — Cowork installs the skills themselves through its plugin
marketplace (below).
On Windows: run inside WSL2 — native Windows (PowerShell/cmd) is unsupported because the skill
runtime is POSIX/bash. See Windows: use WSL2 below.
Marketplace plugin — no clone, no Node (Claude Code & Cowork)
Both Claude Code and Claude Cowork (Anthropic's desktop app)
install from the same plugin marketplace:
Claude Code:
/plugin marketplace add mohitjandwani/analyst-kit
/plugin install us-stock-analyst@analyst-kit # or international-analyst / taiwan-stock-analyst
Claude Cowork (desktop app): Customize → Plugins → Personal plugins → + → Add marketplace →
mohitjandwani/analyst-kit , add the us-stock-analyst plugin, then enable Settings →
Capabilities → Code execution .
After installing, ask a trigger phrase (e.g. "deep dive on NVDA") and the matching skill loads. From a clone
you can also self-test the installers across every platform:
npm run test:integration # real installs per platform + the path.win32 path-layer check
Codex (any OS, no ChatGPT login needed) — confirm a skill is reachable with an API key:
CODEX_API_KEY=sk-... codex exec --json " use the sec-filings skill to list NVDA's latest 8-K "
See compatibility.md for what each runtime does underneath — where skills land, the
routing table, and Windows specifics.
Analyst Kit's skills run a POSIX/bash runtime ( analyst-kit-core ), so on Windows they are
supported only inside WSL2 . This matches what the agents themselves require:
Claude Code's sandbox runs on macOS, Linux, and WSL2 (native Windows is
unsupported), and Codex's Linux mode is WSL2 as well.
Recommended: install WSL2 ,
then run Claude Code (or Codex) and this installer inside your WSL2
distribution — everything then behaves exactly like Linux.
Native Windows (PowerShell/cmd) cannot run the bash runtime at all. Native
Windows + Git Bash will run the scripts but without enforced .env file
permissions ( chmod is a no-op on NTFS) and without sandboxing — unsupported.
node bin/analyst-kit.js doctor --platform claude-code warns when run on native Windows.
Three persona plugins bundle the research workflows for different markets. All
three include the research workflows (deep dive, thematic, technical analysis,
company wiki) plus their supporting capabilities (charting, reporting,
wiki-builder, company-universe-manager, financialmodellingprep, market-intelligence,
analyzing-financial-statements, creating-financial-models, data-analysis); the
market difference is whether FinMind (Taiwan data) and SEC filings are included.
Run node bin/analyst-kit.js list --persona <name> to see a plugin's exact contents.
Keys are read from the environment or a git-ignored .env . The installer
( analyst-kit env / analyst-kit install ) prompts for anything missing when run interactively.
See .env.example .
13f-analysis needs no key — it reads SEC EDGAR directly (set the optional
SEC_EDGAR_UA contact string as an SEC fair-access courtesy).
Skills that run code bootstrap their own dependencies on first use. Runtimes are a
per-skill prerequisite the installer does not install for you: Python (finmind,
company-universe-manager; 13f-analysis is standard-library only) and Bun
(wiki-builder).
The ~/.analyst-kit data home, analytics & updates
Every skill runs on a shared runtime ( analyst-kit-core , installed automatically as a
dependency) that keeps all per-user state in one fixed place, ~/.analyst-kit/ :
.env — your API keys (chmod 600, shared across projects)
config — settings ( analyst-kit-core/bin/analyst-kit-config get|set|list )
analytics/skill-usage.jsonl — local usage log: which skill ran, when,
outcome, duration
learnings.jsonl — things the skills learned about your setup and preferences,
so mistakes aren't repeated
Telemetry is on by default (anonymous, opt-out). It sends only skill name,
version, outcome, and duration — never repo names, file paths, tickers, or content —
and is what tells us which skills break or run slow, so keeping it on directly
improves your experience. You're told about it once on first run. Tiers: community
(default, stable anonymous id), anonymous (no id), off . Opt out any time:
~ /.claude/skills/analyst-kit-core/bin/analyst-kit-config set telemetry off
Updates: skills check the published version at most once a day and offer a
guided upgrade when a new release is out (declining snoozes it for a week; disable
with analyst-kit-config set update_check false ).
Each skill is skills/<name>/SKILL.md with YAML frontmatter:
---
name : 13f-analysis # kebab-case; must equal the folder name
type : capability # capability | workflow
description : > # what it does + a "Triggers:" clause of trigger phrases
... Triggers: "get the 13F for X", "what does <fund> own", ...
requires : [ ... ] # capability skills this one builds on (nothing may require a workflow)
env : [ FMP_API_KEY ] # API keys the skill needs
---
Agents read only name + description , so trigger phrases live inside the
description. type , requires , and env drive the installer and validator.
npm run validate # lint skills + plugin manifests (+ preamble sync check)
npm run build:registry # regenerate registry.json from frontmatter
npm run check:registry # verify registry.json is in sync
npm run sync:preamble # regenerate the analyst-kit-core blocks in every SKILL.md
registry.json is generated — edit skill frontmatter, then rebuild it. The
<!-- analyst-kit:preamble/epilogue --> blocks in each SKILL.md are also generated — edit
skills/analyst-kit-core/templates/ and re-sync; never edit between the markers. The same
checks run in CI ( .github/workflows/validate.yml ) on every push and pull request.
Planned skills, not yet available: an LBO model (debt schedule, cash sweep,
IRR/MOIC) and PDF report analysis.
analyzing-financial-statements and the DCF + sensitivity tooling in
creating-financial-models were inspired by the custom financial skills
in Anthropic's claude-cookbooks ,
then reworked to this repo's skill contract (frontmatter, scripts/ layout)
and hardened with input guards and a test suite.
The M&A accretion/dilution model in creating-financial-models was
inspired by joe-neary/MergerDealSimulator .
The financial formulas were reimplemented from scratch (no code was copied);
that project's worked example serves as an independent cross-check in the
skill's test suite.
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
