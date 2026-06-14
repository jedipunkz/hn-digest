---
source: "https://github.com/DietrichGebert/ponytail"
hn_url: "https://news.ycombinator.com/item?id=48527946"
title: "Ponytail – make your AI agent think like the laziest senior dev in the room"
article_title: "GitHub - DietrichGebert/ponytail: Makes your AI agent think like the laziest senior dev in the room. The best code is the code you never wrote. · GitHub"
author: "mellosouls"
captured_at: "2026-06-14T15:37:15Z"
capture_tool: "hn-digest"
hn_id: 48527946
score: 3
comments: 0
posted_at: "2026-06-14T15:08:17Z"
tags:
  - hacker-news
  - translated
---

# Ponytail – make your AI agent think like the laziest senior dev in the room

- HN: [48527946](https://news.ycombinator.com/item?id=48527946)
- Source: [github.com](https://github.com/DietrichGebert/ponytail)
- Score: 3
- Comments: 0
- Posted: 2026-06-14T15:08:17Z

## Translation

タイトル: ポニーテール – AI エージェントにその部屋で最も怠惰な上級開発者のように思わせる
記事のタイトル: GitHub - DietrichGebert/ponytail: AI エージェントをその場で最も怠惰な上級開発者のように考えさせます。最良のコードは、一度も書いたことのないコードです。 · GitHub
説明: AI エージェントに、その部屋で最も怠惰な上級開発者のように考えさせます。最良のコードは、一度も書いたことのないコードです。 - ディートリッヒ・ゲベルト/ポニーテール

記事本文:
GitHub - DietrichGebert/ponytail: AI エージェントをその場で最も怠惰な上級開発者のように考えさせます。最良のコードは、一度も書いたことのないコードです。 · GitHub
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
ディートリッヒ・ゲベルト
/
ポニーテール
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
37 コミット 37 コミット .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .clinerules .clinerules .codex-plugin .codex-plugin .cursor/ rules .cursor/ rules .github .github .kiro/ステアリング .kiro/ステアリング .opencode .opencode .windsurf/ ルール .windsurf/ ルール アセット アセット ベンチマーク ベンチマーク コマンド コマンド ドキュメント ドキュメントのサンプル サンプル フック フック pi-extension pi-extension スクリプト スクリプト スキル スキル テスト テスト .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md ライセンス ライセンス README.md README.md opencode.json opencode.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
彼は何も言いません。彼は一行書きます。それは動作します。
コードが 80 ～ 94% 少ない · 3 ～ 6 倍高速 · 47 ～ 77% 安い
中央値 10 は、Haiku、Sonnet、Opus にまたがります。自分で再現してください。
あなたは彼を知っています。長いポニーテール。楕円形のメガネ。バージョン管理よりも長く会社に勤務しています。あなたは彼に50行を見せました。彼はそれらを見て、何も言わず、それらを別のものに置き換えます。
ポニーテールは彼を AI エージェントの中に入れます。
日付ピッカーを要求します。エージェントは flatpickr をインストールし、ラッパー コンポーネントを作成し、スタイルシートを追加して、タイムゾーンに関するディスカッションを開始します。
<!-- ポニーテール: ブラウザにはポニーテールがあります -->
< 入力タイプ =" 日付 " >
例ではさらに多くの生存者がいます/ 。
5 つの日常タスク (電子メール検証、デバウンス、CSV 合計、カウントダウン タイマー、レート リミッター)、3 つのモデル、3 つのアーム: スキルなし、穴居人のスキル、およびポニーテール。セルあたり 10 回の実行、中央値が報告されました。
すべてのモデルで、スキルを持たないエージェントに比べて、コード量が 80 ～ 94%、コストが 47 ～ 77% 削減され、3 ～ 6 倍高速になります。ポニーテールが使用するすべてのショートカットは、コード内でポニーテールでマークされます。アップグレード パスを指定するコメントです。自分で再現してみよう

: npx promptfoo eval -c benchmarks/promptfooconfig.yaml 。方法と生の数値: ベンチマーク/ 。制約のないエージェントがはるかに肥大化する運用グレードのタスクは、 benchmarks/results/ に書き込まれます。
コードを記述する前に、エージェントは以下を保持する最初の行で停止します。
1. これは存在する必要がありますか? →いいえ：スキップします（ヤグニ）
2. Stdlib は可能ですか? →使ってください
3. ネイティブプラットフォーム機能? →使ってください
4. 依存関係がインストールされていますか? →使ってください
5. 一行? →一行
6. そのときのみ: 機能する最小限のもの
怠惰ではなく、怠慢ではありません。信頼境界の検証、データ損失の処理、セキュリティ、およびアクセシビリティは決して妨げにはなりません。
ポニーテールにはこれまでで最も労力がかかります:
/プラグインマーケットプレイス追加 DietrichGebert/ポニーテール
/プラグインのインストール ponytail@ponytail
コーデックス
コーデックス プラグイン マーケットプレイスに DietrichGebert/ポニーテールを追加
コーデックス
/plugins を開き、Ponytail マーケットプレイスを選択して、Ponytail をインストールします。それから
/hooks を開き、その 2 つのライフサイクル フックを確認して信頼し、新しいスレッドを開始します。
pi インストール git:github.com/DietrichGebert/ponytail
オープンコード
このリポジトリのチェックアウトから OpenCode を実行し (プラグインはフック/ とスキル/ を再利用します)、 opencode.json に追加します。
{ "プラグイン" : [ " ./.opencode/plugins/ponytail.mjs " ] }
アクティブレベルで毎ターンルールセットを挿入します。 /ponytail 、 /ponytail-review 、および /ponytail-audit を追加します。 OpenCode はこのリポジトリの AGENTS.md も自動ロードするため、プラグインがなくてもルールは適用されます。このプラグインは、lite/full/ultra/off レベルを追加します。
それはそれでした。彼は誇りに思うだろう。彼はそれを言わないでしょう。
すべてのセッションがアクティブになります。 /ponytail-review は diff 内で削除するものを見つけます。/ponytail-audit はリポジトリ全体に対して同じことを行います。 /ponytail Ultra は、コードベースによって個人的に不当な扱いを受けた場合に備えて存在します。 /ponytail-help で残りの説明を説明します。
Codex では、 @ponytail 、 @ponytail-review 、としてスキルを呼び出します。
@ポニーテール

udit 、および @ponytail-help 。起動およびモード変更のテキストには、
現在のモード。
Cursor、Windsurf、Cline、Copilot、Aider、Kiro: このリポジトリから一致するルール ファイルをコピーします ( .cursor/rules/ 、 .windsurf/rules/ 、 .clinerules/ 、 .github/copilot-instructions.md 、 AGENTS.md 、 .kiro/steering/ )。
Kiro: .kiro/steering/ponytail.md をプロジェクト内の ~/.kiro/steering/ (グローバル) または .kiro/steering/ にコピーします。
どのファイルがどのエージェントにマップされるか: エージェントの移植性。
コンパクト ルール テキストを変更するときは、エージェントのコピーを揃えてください。
ノードスクリプト/check-rule-copies.js
よくある質問
設定ファイルが必要なのでしょうか？
いいえ。
120 行のキャッシュ クラスが本当に必要な場合はどうすればよいですか?
あなたはしない。とにかく主張すれば、彼はそれを構築します。ゆっくり。そうです。あなたを見ながら。
スケールはありますか?
あなたが書いたことのないコードは無限に拡張できます。バグ、CVE ゼロ、稼働率 100% が永遠に続きます。
なぜ「ポニーテール」なのか？
その理由はあなたにもよくわかります。
マサチューセッツ工科大学機能する最も短いライセンス。
AI エージェントに、その部屋で最も怠惰な上級開発者のように考えさせます。最良のコードは、一度も書いたことのないコードです。
読み込み中にエラーが発生しました。このページをリロードしてください。
217
フォーク
レポートリポジトリ
リリース
4
v4.2.0: OpenCode で遅延が発生するようになりました
最新の
2026 年 6 月 13 日
+ 3 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Makes your AI agent think like the laziest senior dev in the room. The best code is the code you never wrote. - DietrichGebert/ponytail

GitHub - DietrichGebert/ponytail: Makes your AI agent think like the laziest senior dev in the room. The best code is the code you never wrote. · GitHub
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
DietrichGebert
/
ponytail
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
37 Commits 37 Commits .agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .clinerules .clinerules .codex-plugin .codex-plugin .cursor/ rules .cursor/ rules .github .github .kiro/ steering .kiro/ steering .opencode .opencode .windsurf/ rules .windsurf/ rules assets assets benchmarks benchmarks commands commands docs docs examples examples hooks hooks pi-extension pi-extension scripts scripts skills skills tests tests .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md opencode.json opencode.json package.json package.json View all files Repository files navigation
He says nothing. He writes one line. It works.
80-94% less code · 3-6× faster · 47-77% cheaper
Median of 10 runs across Haiku, Sonnet, and Opus. Reproduce it yourself.
You know him. Long ponytail. Oval glasses. Has been at the company longer than the version control. You show him fifty lines; he looks at them, says nothing, and replaces them with one.
Ponytail puts him inside your AI agent.
You ask for a date picker. Your agent installs flatpickr, writes a wrapper component, adds a stylesheet, and starts a discussion about timezones.
<!-- ponytail: browser has one -->
< input type =" date " >
More survivors in examples/ .
Five everyday tasks (email validator, debounce, CSV sum, countdown timer, rate limiter), three models, three arms: no skill, the caveman skill, and ponytail. Ten runs per cell, median reported.
80-94% less code, 47-77% less cost, and 3-6× faster than a no-skill agent, on every model. Every shortcut ponytail takes is marked in the code with a ponytail: comment naming its upgrade path. Reproduce it yourself: npx promptfoo eval -c benchmarks/promptfooconfig.yaml . Method and raw numbers: benchmarks/ . Production-grade tasks, where an unconstrained agent bloats far more, are written up in benchmarks/results/ .
Before writing code, the agent stops at the first rung that holds:
1. Does this need to exist? → no: skip it (YAGNI)
2. Stdlib does it? → use it
3. Native platform feature? → use it
4. Installed dependency? → use it
5. One line? → one line
6. Only then: the minimum that works
Lazy, not negligent: trust-boundary validation, data-loss handling, security, and accessibility are never on the chopping block.
The most effort ponytail will ever ask of you:
/plugin marketplace add DietrichGebert/ponytail
/plugin install ponytail@ponytail
Codex
codex plugin marketplace add DietrichGebert/ponytail
codex
Open /plugins , select the Ponytail marketplace, and install Ponytail. Then
open /hooks , review and trust its two lifecycle hooks, and start a new thread.
pi install git:github.com/DietrichGebert/ponytail
OpenCode
Run OpenCode from a checkout of this repo (the plugin reuses its hooks/ and skills/ ), and add to opencode.json :
{ "plugin" : [ " ./.opencode/plugins/ponytail.mjs " ] }
Injects the ruleset every turn at the active level; adds /ponytail , /ponytail-review , and /ponytail-audit . OpenCode also auto-loads this repo's AGENTS.md , so the rules hold even without the plugin. The plugin adds the lite/full/ultra/off levels.
That was it. He'd be proud. He won't say it.
Active every session. /ponytail-review finds what to delete in your diff, /ponytail-audit does the same for the whole repo. /ponytail ultra exists for when the codebase has wronged you personally. /ponytail-help explains the rest.
In Codex, invoke the skills as @ponytail , @ponytail-review ,
@ponytail-audit , and @ponytail-help . Startup and mode-change text shows the
current mode.
Cursor, Windsurf, Cline, Copilot, Aider, Kiro: copy the matching rules file from this repo ( .cursor/rules/ , .windsurf/rules/ , .clinerules/ , .github/copilot-instructions.md , AGENTS.md , .kiro/steering/ ).
Kiro: copy .kiro/steering/ponytail.md to ~/.kiro/steering/ (global) or .kiro/steering/ in your project.
Which files map to which agent: Agent portability .
When changing the compact rule text, keep the agent copies aligned:
node scripts/check-rule-copies.js
FAQ
Does it need a config file?
No.
What if I really need the 120-line cache class?
You don't. Insist anyway and he'll build it. Slowly. Correctly. While looking at you.
Does it scale?
The code you never wrote scales infinitely. Zero bugs, zero CVEs, 100% uptime since forever.
Why "ponytail"?
You know exactly why.
MIT . The shortest license that works.
Makes your AI agent think like the laziest senior dev in the room. The best code is the code you never wrote.
There was an error while loading. Please reload this page .
217
forks
Report repository
Releases
4
v4.2.0: lazy in OpenCode now
Latest
Jun 13, 2026
+ 3 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
