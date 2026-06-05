---
source: "https://github.com/zdk/lowfat"
hn_url: "https://news.ycombinator.com/item?id=48409955"
title: "Show HN: Lowfat – pluggable CLI filter that saved 91.8% of my LLM tokens"
article_title: "GitHub - zdk/lowfat: lowfat - slim your command output. strips noise, saves tokens. · GitHub"
author: "zdkaster"
captured_at: "2026-06-05T10:25:36Z"
capture_tool: "hn-digest"
hn_id: 48409955
score: 2
comments: 0
posted_at: "2026-06-05T09:10:42Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Lowfat – pluggable CLI filter that saved 91.8% of my LLM tokens

- HN: [48409955](https://news.ycombinator.com/item?id=48409955)
- Source: [github.com](https://github.com/zdk/lowfat)
- Score: 2
- Comments: 0
- Posted: 2026-06-05T09:10:42Z

## Translation

タイトル: Show HN: Lowfat – LLM トークンの 91.8% を節約したプラグイン可能な CLI フィルター
記事のタイトル: GitHub - zdk/lowfat: lowfat - コマンド出力をスリム化します。ノイズを取り除き、トークンを節約します。 · GitHub
説明: lowfat - コマンド出力をスリム化します。ノイズを取り除き、トークンを節約します。 - ZDK/低脂肪
HN テキスト: こんにちは、HN。興味のある人がいるかどうかはわかりません。ただし、冗長な CLI 出力の一部をフィルタリングするのに役立つ「lowfat」という小さなツールをメンテナンスしていることを共有したいと思います。これは単一のバイナリであり、エージェント フックまたはシェル ラッパーとして機能します。
コマンドごとにフィルターをカスタマイズするプラグイン システムがあります。考え方は非常にシンプルです。エージェントは、意思決定を行うために完全な kubectl get -o yaml や 10,000 行のダンプを必要としません。
そのため、低脂肪が間に存在し、ノイズを取り除き、重要なものを通過させます。 2 か月間個人的に使用した後の実際のレポートは次のとおりです: 低脂肪履歴 -- すべての低脂肪プラグイン候補
───────────────────
# コマンド実行時の平均コスト削減ソース ステータス
1 kubectl 取得 101x 14.4K 1.5M 93.9% プラグイン良好
2 grep 103x 13.5K 1.4M 96.2% プラグイン良好
3 git diff 81x 995 80.6K 57.9% 組み込み良好
4 kubectl 90x 485 43.6K 33.6% プラグインは良好
5 ドッカー 127x 5.5K 693.6K 96.1% 組み込み良好
6 ls 489x 117 57.3K 56.2% 内蔵良好
7 検索 30x 16.5K 495.0K 95.5% プラグインが良好
8 git show 63x 490 30.9K 38.0% 組み込み良好
9 git 177x 368 65.2K 76.1% 組み込み良好
10 git log 86x 556 47.8K 78.5% 組み込み良好
11 kubectl ログ 5x 3.6K 17.8K 43.0% プラグインは良好
12 git status 86x 152 13.1K 58.0% 組み込み良好
13 docker ps 20x 467 9.3K 52.8% プラグインは良好
14 kubectl 記述 6x 656 3.9K 1.2% プラグインが弱い
15 個の Docker イメージ 9x 940 8.5K 61.8% 組み込み良好
16kで2x2をゲット。

1K 4.2K 35.9% プラグインは良好
17 terraform 10x 395 3.9K 32.1% プラグインは良好
18 git commit 32x 77 2.5K 0.0% 組み込みが弱い
19 docker build 8x 487 3.9K 37.6% 組み込み良好
20 docker compose 22x 979 21.5K 89.4% 組み込み良好
合計: 未加工 440 万 → 保存 410 万 (91.8%)
上記のツールセットは種類が限られていますが、私のユースケースでは中断することなく非常にうまく機能します。
会社のBedrock制限使用量のトークン制限に達しないようにし、後で使用するために外出先で節約を最適化し続けるのを手伝ってください。しかし、代替手段 ( https://github.com/zdk/lowfat#alternatives ) を使用しないのはなぜでしょうか?
答えは次のとおりです。
- 私の目標は、コアを軽量でありながらプラグインによって拡張できるようにすることです。つまり、ユーザーが出力フィルターを所有できるように、インストールされたバイナリ内のすべてのコマンドをバンドルしようとはしません。
- 独自のツールセットを使用しているため、プラグインまたはフィルター パイプラインを介してユースケースごとにカスタマイズできます。
- 非公開 CLI ツール用にカスタマイズ可能。たとえば、一部の企業では、一般公開されていない内部 CLI ツールを使用している場合があります。
- 人々は自分のデータを所有する必要があります。したがって、設計はローカルファーストであり、テレメトリーは永久にありません。
- 私は UNIX スタイルの合成パイプが大好きなので、lowfat-filter はこのスタイルを実装しました。
- フィルターの攻撃性を調整できるため、エージェントが必要とするものを削除しないように制御できます。 GitHub: https://github.com/zdk/lowfat とにかく、興味のある方はフィードバックや質問をお待ちしています。ありがとう！

記事本文:
GitHub - zdk/lowfat: lowfat - コマンド出力をスリム化します。ノイズを取り除き、トークンを節約します。 · GitHub
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
zdk
/
低脂肪
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ 移動先

ファイル コード [その他のアクション] メニューを開く フォルダーとファイル
100 コミット 100 コミット .github/ workflows .github/ workflows crates crates docs docs test-fixtures/ plugins test-fixtures/ plugins .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
lowfat は、エージェントに到達する前に不要な CLI 出力をフィルタリングすることで AI トークンのコストを削減する軽量の CLI ツールです。
軽量 — 小さな単一バイナリ、小さなコア。しかし拡張可能。
ローカルファースト — テレメトリなし。データの所有者はあなたです。
コンポーザブル — UNIX スタイルのパイプ、組み込みフィルターと独自のフィルターを組み合わせます。魔法ではありません。
ユーザー所有 — 低脂肪履歴には、最も頻繁に実行したものが表示されます。ユースケースに合わせてカスタマイズできます。
カーゴインストール低脂肪
# または
醸造インストール zdk/tools/lowfat
GitHub リリースで事前に構築されたバイナリ。
クロード コード フック — .claude/settings.json に追加します。
{
「フック」: {
"PreToolUse" : [
{
"matcher" : " Bash " 、
"フック" : [{ "タイプ" : " コマンド " , "コマンド" : " lowfat フック " }]
}
]
}
}
シェル統合 — エージェント環境内で自動的にアクティブ化します ( CLAUDECODE=1 、 CODEX_ENV )。または、 LOWFAT_ENABLE=1 を設定して任意のシェルに強制します。
echo ' eval "$(lowfatshell-init zsh)" ' >> ~ /.zshrc # または ~/.bashrc
OpenCode プラグイン — 1 つのコマンドで構成編集なし:
lowfat opencode install # ~/.config/opencode/plugins/lowfat.ts を書き込みます
OpenCode を再起動します。コマンドは実行前に透過的に書き換えられます。
lowfat opencode uninstall でいつでも削除できます。
直接使用 - 任意のコマンドに接頭辞を付けます。
低脂肪gitステータス
低脂肪ドッカー PS
低脂肪 ls -la
Pi エージェント — ~/.pi/agent/settings.json 内:
{ "shellCommandPrefix" : " eval \" $(lowfat shell-init zsh) \" ; " }
使用法のハイライト
# 設定内容と各フィルターの音量を確認する
低脂肪情報 #

ステータスバッジ + アクティブフィルター
lowfat info git # `git` のパイプライン
lowfat info --config # 完全に解決された構成
# 低脂肪があなたを救ってくれたものを見てください
低脂肪統計 # 生涯トークン節約額
lowfat stats --audit 最近のプラグイン実行数
低脂肪履歴 # 潜在的な節約によるコマンドのランク付け
# アグレッシブネスを調整する
低脂肪レベルのウルトラ最大圧縮
LOWFAT_LEVEL=lite lowfat git log # 1 回限りのオーバーライド
# プラグインを書く
lowfat プラグインの新しい terraform # scaffold ~/.lowfat/plugins/terraform/
lowfat plugin Doctor # プラグインをチェックします (そして Python deps を事前にインストールします)
# プラグインをインストールせずにサンプルに対してテストする
猫のサンプル/git-diff-full.txt |低脂肪フィルター --explain ./filter.lf --sub=diff --level=ultra
さらに詳しく
docs/ARCHITECTURE.md — 概要図: CLI、ランナー、プラグイン、ビルトイン
docs/CONFIG.md — .lowfat ファイル、環境変数、パイプライン DSL、内蔵プロセッサ、履歴ランキング
docs/PLUGINS.md — lf-filter (.lf プラグイン DSL)、シェル エスケープ ハッチ、PEP 723 + uv、AI エージェント プロンプト
このプロジェクトでは複数の AI ツールが使用されました
lowfat - コマンド出力をスリム化します。ノイズを取り除き、トークンを節約します。
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
2
フォーク
レポートリポジトリ
リリース
33
v0.6.8
最新の
2026 年 6 月 5 日
+ 32 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

lowfat - slim your command output. strips noise, saves tokens. - zdk/lowfat

Hi HN, Not sure if anyone would be interested. But, just wanted to share that I've been maintaining my small tool called 'lowfat' that helps me filters some of my verbose CLI output. It's a single binary, works as an agent hook or a shell wrapper.
It has a plugin system to customize filters per command. The idea is pretty simple: agents don't need the full kubectl get -o yaml or any 10k-line dump to make decisions.
So that lowfat sits in between, strips the noise, and passes through what matters. Here's my real report after 2 months of personal use: lowfat history --all lowfat plugin candidates
─────────────────────────────────────────────────────────
# command runs avg raw cost savings source status
1 kubectl get 101x 14.4K 1.5M 93.9% plugin good
2 grep 103x 13.5K 1.4M 96.2% plugin good
3 git diff 81x 995 80.6K 57.9% built-in good
4 kubectl 90x 485 43.6K 33.6% plugin good
5 docker 127x 5.5K 693.6K 96.1% built-in good
6 ls 489x 117 57.3K 56.2% built-in good
7 find 30x 16.5K 495.0K 95.5% plugin good
8 git show 63x 490 30.9K 38.0% built-in good
9 git 177x 368 65.2K 76.1% built-in good
10 git log 86x 556 47.8K 78.5% built-in good
11 kubectl logs 5x 3.6K 17.8K 43.0% plugin good
12 git status 86x 152 13.1K 58.0% built-in good
13 docker ps 20x 467 9.3K 52.8% plugin good
14 kubectl describe 6x 656 3.9K 1.2% plugin weak
15 docker images 9x 940 8.5K 61.8% built-in good
16 k get 2x 2.1K 4.2K 35.9% plugin good
17 terraform 10x 395 3.9K 32.1% plugin good
18 git commit 32x 77 2.5K 0.0% built-in weak
19 docker build 8x 487 3.9K 37.6% built-in good
20 docker compose 22x 979 21.5K 89.4% built-in good
total: 4.4M raw → 4.1M saved (91.8%)
My toolset above is kind limited, but it works pretty well for my usecase without any interruption
Kinda help me not reaching the token limit for my company Bedrock limit usage and keep optimizing the saving on the go for later usage. But, why not alternatives ( https://github.com/zdk/lowfat#alternatives ) ?
The answers are:
- My goal is to make the core lightweight but extensible via plugins i.e. not trying to bundle every command in the installed binary so that people own their output filters.
- Customizable per usecase via plugin or filter pipelines as I am using my own toolset.
- Customizable for non-public CLI tools, for example, some enterprise might have their interal CLI tools that public won't have access.
- People should own their data. So the design is local-first, No telemetry forever.
- I kinda love UNIX-style composible pipes, so lowfat-filter has implemented this style.
- Be able to adjust aggressiveness of the filter, so we can control that we won't strip something the agent needed. GitHub: https://github.com/zdk/lowfat Anyway, if anyone is interested, feedbacks and questions are welcome! Thanks!

GitHub - zdk/lowfat: lowfat - slim your command output. strips noise, saves tokens. · GitHub
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
zdk
/
lowfat
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
100 Commits 100 Commits .github/ workflows .github/ workflows crates crates docs docs test-fixtures/ plugins test-fixtures/ plugins .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md View all files Repository files navigation
lowfat is a lightweight CLI tool that reduces AI token costs by filtering unnecessary CLI output before it reaches your agent.
Lightweight — Small single binary, small core; but extensible.
Local-first — No telemetry; you own your data.
Composable — UNIX-style pipes, mix built-ins and your own filters; not magic.
User-owned — lowfat history shows what you run most; allow you to customize for your usecase.
cargo install lowfat
# or
brew install zdk/tools/lowfat
Pre-built binaries on GitHub Releases .
Claude Code hook — add to .claude/settings.json :
{
"hooks" : {
"PreToolUse" : [
{
"matcher" : " Bash " ,
"hooks" : [{ "type" : " command " , "command" : " lowfat hook " }]
}
]
}
}
Shell integration — auto-activates inside agent environments ( CLAUDECODE=1 , CODEX_ENV ), or set LOWFAT_ENABLE=1 to force it on any shell:
echo ' eval "$(lowfat shell-init zsh)" ' >> ~ /.zshrc # or ~/.bashrc
OpenCode plugin — one command, no config editing:
lowfat opencode install # writes ~/.config/opencode/plugins/lowfat.ts
Restart OpenCode; commands are rewritten transparently before they run.
Remove it anytime with lowfat opencode uninstall .
Direct usage — prefix any command:
lowfat git status
lowfat docker ps
lowfat ls -la
Pi agent — in ~/.pi/agent/settings.json :
{ "shellCommandPrefix" : " eval \" $(lowfat shell-init zsh) \" ; " }
Usage highlights
# See what's configured and how loud each filter is being
lowfat info # status badge + active filters
lowfat info git # pipeline for `git`
lowfat info --config # full resolved config
# See what lowfat has saved you
lowfat stats # lifetime token savings
lowfat stats --audit # recent plugin executions
lowfat history # rank commands by potential savings
# Dial the aggressiveness
lowfat level ultra # max compression
LOWFAT_LEVEL=lite lowfat git log # one-off override
# Write a plugin
lowfat plugin new terraform # scaffold ~/.lowfat/plugins/terraform/
lowfat plugin doctor # check plugins (and pre-install any Python deps)
# Test a plugin against a sample without installing it
cat samples/git-diff-full.txt | lowfat filter --explain ./filter.lf --sub=diff --level=ultra
Learn more
docs/ARCHITECTURE.md — high-level diagram: CLI, Runner, Plugins, Builtins
docs/CONFIG.md — .lowfat file, env vars, pipeline DSL, built-in processors, the history ranking
docs/PLUGINS.md — lf-filter (the .lf plugin DSL), shell escape hatches, PEP 723 + uv, AI agent prompt
Multiple AI tools were used for this project
lowfat - slim your command output. strips noise, saves tokens.
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
2
forks
Report repository
Releases
33
v0.6.8
Latest
Jun 5, 2026
+ 32 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
