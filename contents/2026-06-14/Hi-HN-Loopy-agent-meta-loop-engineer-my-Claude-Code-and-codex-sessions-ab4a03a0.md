---
source: "https://github.com/secretbuilds/loopy"
hn_url: "https://news.ycombinator.com/item?id=48524150"
title: "Hi HN: Loopy agent, meta-loop engineer my Claude Code and codex sessions"
article_title: "GitHub - secretbuilds/loopy: meta-agent for claude code to automatically loop engineer your workflow · GitHub"
author: "secretbuilds"
captured_at: "2026-06-14T04:35:15Z"
capture_tool: "hn-digest"
hn_id: 48524150
score: 1
comments: 1
posted_at: "2026-06-14T04:22:23Z"
tags:
  - hacker-news
  - translated
---

# Hi HN: Loopy agent, meta-loop engineer my Claude Code and codex sessions

- HN: [48524150](https://news.ycombinator.com/item?id=48524150)
- Source: [github.com](https://github.com/secretbuilds/loopy)
- Score: 1
- Comments: 1
- Posted: 2026-06-14T04:22:23Z

## Translation

タイトル: こんにちは、HN: ループ エージェント、メタ ループ エンジニア、私のクロード コードとコーデックス セッション
記事のタイトル: GitHub - Secretbuilds/loopy: ワークフローを自動的にループ エンジニアリングするためのクロード コード用メタエージェント · GitHub
説明: ワークフローを自動的にループエンジニアリングするクロードコード用のメタエージェント - Secretbuilds/loopy

記事本文:
GitHub - Secretbuilds/loopy: ワークフローを自動的にループエンジニアリングするためのクロード コード用のメタエージェント · GitHub
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
秘密のビルド
/
おかしな
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード

main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
59 コミット 59 コミット コマンド コマンド docs docs src src testing testing .gitignore .gitignore CLAUDE-FABLE-5.md CLAUDE-FABLE-5.md README.md README.md install.sh install.sh package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ユーザーの作業方法を監視し、パターンを見つけて、ループを作成するターミナル メタ エージェント。
「私はもうクロードにプロンプ​​トを出しません。私はクロードにプロンプ​​トを出して何をすべきかを判断するループを実行しています。私の仕事はループを書くことです。」
— Boris Cherny 、Anthropic の Claude Code 作成者
「コーディング エージェントにプロンプトを表示するのはやめてください。プロンプトを表示するループの設計を始めてください。」
— Peter Steinberger、OpenClaw (旧 Warelay) 創設者
Claude Code と Codex を使用しているトップ エンジニアは、手動でプロンプトをやり取りしていません。彼らは自律的なループ、つまり彼らに代わって観察し、決定し、行動するプログラムを構築しています。問題: ほとんどの開発者はどこから始めればよいのかわかりません。自動化する適切なものを見つけるには、観察、パターン認識、そして時間がかかります。
LOOPY は、Claude Code と並行してターミナル内で静かに実行されるローカル メタエージェントです。セッションを監視し、手作業で続けている作業を特定し、すぐにインストールできる自動化ループを提案します。各自動化ループは、独自のトリガー、操作手順、インストール レコードを備えた自己完結型のスクリプトです。
ダッシュボードで提案を確認します。あなたは意味のあるものを承認します。ループが実行されます。
すべてがマシン上に残ります。 LLM 呼び出しのみが独自の claude -p バイナリを経由します。ループは自宅に電話をかけることはありません。
なぜループエンジニアリングなのか?数字です。
ほとんどのクロード コード ユーザーは、潜在能力の 40 ～ 70% を t に残しています。

プロンプト応答モードを維持することで可能になります。ループに移行すると何が変わるかは次のとおりです。
適切に選択された 1 つのループが、月あたり 50 ～ 200 件の手動プロンプトを置き換えます。プロンプト サイクルあたり 3 秒で、3 つのアクティブなループにより 1 日あたり約 15 ～ 30 分が節約され、ループ ライブラリが増大するにつれてさらに短縮されます。
初期の頭の悪いユーザーは、手動では決して見つけられなかった自動化可能なパターンを月に 3 ～ 5 個発見したと報告しています。そのほとんどは 10 分間のインストールになります。
LOOPY は、バックグラウンドで継続的に実行される小さなローカル パイプラインです。
クロード・コードのセッション
│
▼
[watcher] — launchd デーモンを介して新しいセッションのトランスクリプトを通知します
│
▼
[ダイジェスター] — 各セッションを圧縮し、コンパクトなテキスト ダイジェストに編集します。
│
▼
[engine] — 独自の claude -p CLI にダイジェストを送信し、繰り返し発生するパターンを検索します。
│
▼
[受信箱] — 優れた候補者が提案として到着します。あなたは確認して承認します
│
▼
[インストールされたループ] — クロード コードまたはコーデックスに接続されたループ.md + トリガー + マニフェスト
エンジンは独自の Claude API クレジットを使用します。個別のサービス、サブスクリプション、クラウド コンポーネントはありません。
カール -fsSL https://raw.githubusercontent.com/secretbuilds/loopy/main/install.sh |バッシュ
次にセットアップを完了します。
おかしな設定
インストーラーはリポジトリを ~/.loopy-app に複製してビルドし、ループ状のバイナリを PATH にリンクします。同じコマンドを再実行すると、最新バージョンに更新されます。
ループ状のセットアップ --コンパニオンマニュアル # 自動ナッジなし
ループ状のセットアップ --no-daemon # バックグラウンド デーモンなしで設定します
要件: ノード ≥ 20、git、PATH 内のクロード コード CLI ( claude )、macOS。
おかしな
引数なしでloopyを実行すると、完全なターミナルハブが開きます。
ヘッダーには、エージェント (Loopy) のライブ ステータス (監視されたセッション、デーモンの状態、今日のトークンの使用量と上限) が表示されます。
受信箱 — 保留中のループ提案。 1 つを選択すると、その概要、推定される影響、証拠が表示されます。

エンスカウントと信頼スコア。承認、却下、またはスヌーズします。
ループ — インストールされているループ、トリガーの種類とターゲット ツール。
activity — ルーピーがバックグラウンドで実行したすべてのログをスクロールします。
ダッシュボードのサイズは端末に合わせて変更されます。少なくとも 60×16 が必要です。それより下には、ウィンドウを拡大するためのヒントが表示されます。
コマンド
何をするのか
おかしな
完全なターミナル ダッシュボードを開く
おかしなレビュー
提案受信箱に重点を置いたダッシュボードを開く
おかしな仲間
裸のコマンドのエイリアス
おかしな設定
設定、トリガーフック、デーモンの初期化
ループ状のセットアップ --no-daemon
バックグラウンドデーモンを使用しない構成
ループ状のスキャン
ローカルダイジェストを分析し、新しい提案を今すぐ明らかにします
おかしなリスト
インストールされているループをリストする
ループ状のアンインストール <id>
ループとそれによってインストールされたすべてのものを削除します
ループ状の一時停止
バックグラウンドデーモンを一時停止する
おかしな履歴書
バックグラウンドデーモンを再開します
ループ状態
デーモン、支出、プロポーザルのステータスを表示する
ループマーク
セッション マーカーをドロップします (トリガー フックによって使用されます)。
おかしなデーモン
ウォッチャーをフォアグラウンドで実行する
/fable — クロード・フェイブル 5 スラッシュ コマンド
また、インストーラーは /fable スラッシュ コマンドを ~/.claude/commands/ にドロップするので、どのクロード コード セッションでもグローバルに使用できるようになります。
/fable <プロンプト>
これにより、現在のセッションのインラインで、完全な Claude Fable 5 システム プロンプトとモデルを介してプロンプトがルーティングされます。セッションを終了したりモデルを変更したりすることなく、その 1 つのメッセージに対して Fable 5 に切り替えたかのように応答が返されます。
/fable このコンポーネントをより美しくします
/fable これをより会話的な口調で書き直します
/fable 技術者以外の関係者にこれを説明する
内部では claude -p --model claude-fable-5 --system-prompt-file ~/.loopy/prompts/fable.md をサブプロセスとして生成し、出力をインラインで返します。 Fable 5 システム プロンプトは ~/.loopy/prompts/fable.md (ins

自動的に計算されます - インストーラーを実行する以外にセットアップは必要ありません)。
承認されたプロポーザルは ~/.loopy/loops/<id>/ でバンドルになります。
loop.md — クロードに提供される操作指示
trigger.json — スケジュール、フック、または手動トリガーのメタデータ
manifest.json — 証拠、ターゲット ツール、インストールされているすべてのパス
state/ — ループローカル状態 (実行全体にわたって持続)
マニフェストは、loopy アンインストールを正確にするものです。マニフェストは、推測に頼ることなく、その特定のループ用に作成されたパスのみを削除します。
トランスクリプトはマシン上に残ります。いつも。
ダイジェストが LLM に送信される前に、ループ状の編集が行われます。
API キー、トークン、パスワード、ベアラー トークン
GitHub トークン、AWS キー、URL 認証情報
秘密のように見える高エントロピーの文字列
エンジンは、コンパクトに編集されたダイジェストのみを独自の claude -p プロセスに送信します。 loopy は外部サービスに接続しません。
Claude Code および Codex のパワー ユーザーで、ループ エンジニアリングへの手動プロンプトから卒業したいが、どこから始めればよいかわからない
AI 支援プロジェクトを開発し、反復的なプロンプトによる認知的オーバーヘッドを体系的に削減したいと考えている開発者
Boris のワークフローの投稿を読んで、次のように考えた人はいるでしょう。「あれは欲しいけど、自分のパターンはまだ見つけられない」
loopy は観測レイヤーを提供します。あなたが判決を下します。ループが残りを処理します。
ルーピーは自動化を提案するかもしれませんが、提案を無視したり、居眠りしたり、却下したりすることを恥じることはありません。しつこいツールよりも、静かなツールの方が優れています。受信箱、電話。
git clone https://github.com/secretbuilds/loopy.git
CD ルーピー
npmインストール
npm ビルドを実行する
npmリンク
おかしな設定
完全なテスト スイートを実行します。
npm run typecheck && npx vitest run
手動煙レンダリング (ダッシュボードのジオメトリを確認):
npx tsx スクリプト/dash-smoke.ts 120 40
ライブ提案品質評価 (claude CLI が必要):
npx tsx スクリプト/live-eval.ts
ルーピーは初期のソフトウェアです

再。 launchd 経由で macOS 上の Claude Code セッションを監視します。 Windows/Linux デーモンのサポートはロードマップにあります。
ワークフローを自動的にループエンジニアリングするためのクロードコード用メタエージェント
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

meta-agent for claude code to automatically loop engineer your workflow - secretbuilds/loopy

GitHub - secretbuilds/loopy: meta-agent for claude code to automatically loop engineer your workflow · GitHub
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
secretbuilds
/
loopy
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
59 Commits 59 Commits commands commands docs docs src src tests tests .gitignore .gitignore CLAUDE-FABLE-5.md CLAUDE-FABLE-5.md README.md README.md install.sh install.sh package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
A terminal meta-agent that watches how you work, finds the patterns, and writes the loops so you don't have to.
"I don't prompt Claude anymore. I have loops running that prompt Claude and figure out what to do. My job is to write loops."
— Boris Cherny , creator of Claude Code at Anthropic
"Stop prompting coding agents. Start designing the loops that prompt them."
— Peter Steinberger , founder of OpenClaw (formerly Warelay)
The top engineers using Claude Code and Codex aren't manually prompting back and forth. They're building autonomous loops — programs that observe, decide, and act on their behalf. The problem: most developers don't know where to start. Finding the right things to automate takes observation, pattern recognition, and time you don't have.
loopy is a local meta-agent that runs quietly in your terminal alongside Claude Code. It watches your sessions, spots work you keep doing by hand, and proposes ready-to-install automation loops — each one a self-contained script with its own trigger, operating instructions, and install record.
You review proposals in the dashboard. You approve the ones that make sense. The loop runs.
Everything stays on your machine. The only LLM calls go through your own claude -p binary — loopy never phones home.
Why loop engineering? The numbers.
Most Claude Code users are leaving 40–70% of their potential on the table by staying in prompt-response mode. Here's what shifts when you move to loops:
A single well-chosen loop replaces 50–200 manual prompts per month. At 3 seconds per prompt cycle, 3 active loops save you roughly 15–30 minutes per day — compounding as your loop library grows.
Early loopy users report discovering 3–5 automatable patterns per month they would never have spotted manually. Most of those become 10-minute installs.
loopy is a small local pipeline that runs continuously in the background:
Claude Code sessions
│
▼
[watcher] — notices new session transcripts via launchd daemon
│
▼
[digester] — compresses + redacts each session to a compact text digest
│
▼
[engine] — sends digests to your own claude -p CLI, looks for recurring patterns
│
▼
[inbox] — good candidates land as proposals; you review and approve
│
▼
[installed loop] — loop.md + trigger + manifest wired into Claude Code or Codex
The engine uses your own Claude API credits. No separate service, no subscription, no cloud component.
curl -fsSL https://raw.githubusercontent.com/secretbuilds/loopy/main/install.sh | bash
Then finish setup:
loopy setup
The installer clones the repo to ~/.loopy-app , builds it, and links the loopy binary to your PATH. Re-running the same command updates to the latest version.
loopy setup --companion manual # no automatic nudges
loopy setup --no-daemon # configure without the background daemon
Requirements: Node ≥ 20, git, Claude Code CLI ( claude ) in your PATH, macOS.
loopy
Running loopy with no arguments opens the full-terminal hub:
The header shows your agent (Loopy) with live status: sessions watched · daemon state · today's token spend vs cap.
inbox — pending loop proposals. Select one to see its summary, estimated impact, evidence count, and confidence score. Approve, dismiss, or snooze.
loops — your installed loops, with trigger kind and target tool.
activity — scrolling log of everything loopy has done in the background.
The dashboard resizes with your terminal. It needs at least 60×16 — below that it shows a hint to grow the window.
Command
What it does
loopy
Open the full-terminal dashboard
loopy review
Open dashboard focused on the proposal inbox
loopy companion
Alias for the bare command
loopy setup
Initialize config, trigger hook, and daemon
loopy setup --no-daemon
Configure without the background daemon
loopy scan
Analyze local digests and surface new proposals now
loopy list
List installed loops
loopy uninstall <id>
Remove a loop and everything it installed
loopy pause
Pause the background daemon
loopy resume
Resume the background daemon
loopy status
Show daemon, spend, and proposal status
loopy mark
Drop a session marker (used by the trigger hook)
loopy daemon
Run the watcher in the foreground
/fable — Claude Fable 5 slash command
The installer also drops a /fable slash command into ~/.claude/commands/ so it's available in any Claude Code session , globally.
/fable <your prompt>
This routes your prompt through the full Claude Fable 5 system prompt and model, inline in your current session. The response comes back as if you'd switched to Fable 5 for that one message — without leaving your session or changing your model.
/fable make this component prettier
/fable rewrite this in a more conversational tone
/fable explain this to a non-technical stakeholder
Under the hood it spawns claude -p --model claude-fable-5 --system-prompt-file ~/.loopy/prompts/fable.md as a subprocess and returns the output inline. The Fable 5 system prompt is stored at ~/.loopy/prompts/fable.md (installed automatically — no setup needed beyond running the installer).
An approved proposal becomes a bundle at ~/.loopy/loops/<id>/ :
loop.md — operating instructions fed to Claude
trigger.json — schedule, hook, or manual trigger metadata
manifest.json — evidence, target tool, every path installed
state/ — loop-local state (persists across runs)
The manifest is what makes loopy uninstall exact: it removes only the paths it created for that specific loop, with no guesswork.
Transcripts stay on your machine. Always.
Before any digest is sent to an LLM, loopy redacts:
API keys, tokens, passwords, bearer tokens
GitHub tokens, AWS keys, URL credentials
High-entropy strings that look like secrets
The engine sends only compact redacted digests to your own claude -p process. loopy does not contact any external service.
Claude Code and Codex power users who want to graduate from manual prompting to loop engineering but don't know where to start
Developers shipping AI-assisted projects who want to systematically reduce the cognitive overhead of repetitive prompting
Anyone who's read Boris's workflow post and thought: I want that, but I can't find my own patterns yet
loopy gives you the observation layer. You bring the judgment. The loops handle the rest.
loopy may suggest automation, but it never shames you for ignoring, snoozing, or dismissing a proposal. A quiet tool is better than a nagging one. Your inbox, your call.
git clone https://github.com/secretbuilds/loopy.git
cd loopy
npm install
npm run build
npm link
loopy setup
Run the full test suite:
npm run typecheck && npx vitest run
Manual smoke render (verify dashboard geometry):
npx tsx scripts/dash-smoke.ts 120 40
Live proposal-quality eval (requires claude CLI):
npx tsx scripts/live-eval.ts
loopy is early software. It watches Claude Code sessions on macOS via launchd. Windows/Linux daemon support is on the roadmap.
meta-agent for claude code to automatically loop engineer your workflow
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
