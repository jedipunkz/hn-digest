---
source: "https://github.com/sofumel/claude-handoff-revive"
hn_url: "https://news.ycombinator.com/item?id=48657973"
title: "Show HN: Resume Claude Code from a handoff file auto-save, snapshots, PR sharing"
article_title: "GitHub - sofumel/claude-handoff-revive: Resume Claude Code work after rate/usage/context limits without replaying the prior transcript. Auto-saves at 90%/95% usage. Plugin-installable, 10 languages. · GitHub"
author: "sofumel"
captured_at: "2026-06-24T12:01:44Z"
capture_tool: "hn-digest"
hn_id: 48657973
score: 2
comments: 0
posted_at: "2026-06-24T11:05:17Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Resume Claude Code from a handoff file auto-save, snapshots, PR sharing

- HN: [48657973](https://news.ycombinator.com/item?id=48657973)
- Source: [github.com](https://github.com/sofumel/claude-handoff-revive)
- Score: 2
- Comments: 0
- Posted: 2026-06-24T11:05:17Z

## Translation

タイトル: HN の表示: ハンドオフ ファイルの自動保存、スナップショット、PR 共有からのクロード コードの再開
記事のタイトル: GitHub - sofumel/claude-handoff-revive: レート/使用量/コンテキストの制限後に、以前のトランスクリプトを再生せずにクロード コードの作業を再開します。 90%/95% の使用率で自動保存されます。プラグインでインストール可能、10言語。 · GitHub
説明: レート/使用量/コンテキストの制限後に、以前のトランスクリプトを再生せずにクロード コードの作業を再開します。 90%/95% の使用率で自動保存されます。プラグインでインストール可能、10言語。 - ソフメル/クロード・ハンドオフ・リバイブ
HN テキスト: これはまだ進行中の作業であり、改善できる点はたくさんあると思います。皆様のご意見、ご提案、建設的なご批判をお待ちしております。

記事本文:
GitHub - sofumel/claude-handoff-revive: レート/使用量/コンテキストの制限後に、以前のトランスクリプトを再生せずにクロード コードの動作を再開します。 90%/95% の使用率で自動保存されます。プラグインでインストール可能、10言語。 · GitHub
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
ソフメル
/
クロード・ハンドオフ・リバイブ

公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows 資産 アセット プラグイン/ handoff-revive プラグイン/ handoff-revive .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md HOOK_SETUP.md HOOK_SETUP.md LICENSE LICENSE README.md README.md install.ps1 install.ps1 install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
レート/使用量/コンテキストの制限後も、以前のセッションを再生せずにクロード コードの作業を続行します。通常は数万のトークン、長いセッションでは 100k 以上のトークンが必要になります。
🇺🇸 英語 ·
🇯🇵 日本語 ·
🇨🇳 中文 ·
🇹🇼繁體中文 ·
🇰🇷 한국어 ·
🇪🇸スペイン語 ·
🇵🇹ポルトガル人 ·
🇩🇪 ドイツ語 ·
🇫🇷 フランセ ·
🇹🇷トゥルクチェ
Claude Code が「限界に達しました · リセット ...」を表示すると、標準のリカバリ (claude --resume または claude -c) によって、以前の会話全体がコンテキストにリロードされます。通常、中程度のセッションでは、1 つの質問をする前に数万のトークンが消費され、長いセッションの場合は 10 万以上のトークンが消費されることもあります。
続行するために必要な最小限のもののみが保存され、 .claude/handoff/current.md に構造化されます (~1 ～ 3,000 トークン)。
目標 — 達成しようとしていたこと
Done / Wip / todo — 完了したタスク、進行中のタスク、まだ開始されていないタスク
next_action — 具体的な次のステップ (file:line + 正確なコマンド)
touch_files — どのファイルがタッチされたのか、なぜタッチされたのか
意思決定 — それぞれの背後にある理由を含むデザインの選択
Less_learned — 失敗した試みとその教訓 (オプション)
再開するには、新しいセッションを開始します ( --resume は使用しないでください)。スキルはそれだけを読み取ります

1 つのファイルを作成し、中断したところから再開します。
プラグインのインストール (推奨) — クロード コード内でこれらを実行します。
/plugin Marketplace add sofumel/claude-handoff-revive # 1. マーケットプレイスに登録
/plugin install handoff-revive@handoff-revive-marketplace # 2. プラグインをインストールします
スキル、スラッシュコマンド、フックは自動で発動します。 settings.json の編集は必要ありません。
/plugin はこの環境では利用できませんか?クロード コードのバージョンが古すぎます。 claude --version で確認し、brew upgrade claude-code (Homebrew) または npm update -g @anthropic-ai/claude-code (npm) で更新します。アップデートできない場合は、以下の手動インストールを使用してください。
手動インストール (常に機能し、プラグインのサポートは必要ありません):
git clone https://github.com/sofumel/claude-handoff-revive.git
cd クロード・ハンドオフ・リバイブ
./install.sh /path/to/your-project # Linux/macOS/WSL/Git-Bash
# .\install.ps1 -Target C:\path\to\proj # Windows PowerShell
# またはすべてのプロジェクトに対してグローバルに:
./install.sh --global # .\install.ps1 -Global
手動インストールの場合は、オプションで HOOK_SETUP.md のスニペットを .claude/settings.json にマージすることでフックを有効にします (プラグインのインストールではこれが自動的に行われます)。
.claude/handoff/ を git にコミットする必要がありますか?どちらも機能します。慎重にどちらかを選択してください。
ソロ使用 (デフォルト) — すべて無視します。揮発性の状態とスナップショットは履歴に属しません。
.claude / ハンドオフ /
チーム共有 — 現在のハンドオフのみをコミットします。スナップショットとマーカー ファイルをローカルに保持します。
.claude / ハンドオフ / *
！ .claude / ハンドオフ / current.md
ドットマーカー ファイル ( .turn 、 .usage-flag など) やhistory/ は決してコミットしないでください。これらはマシンローカルの状態とスナップショットのノイズです。
PR に共有する場合、本文はサニタイズされたコピーから構築されます。 author_email 行が削除され、プロジェクト/ホームの絶対パスが <project-root> / ~ になり、既知のプレフィックス API キー/トークン (例: sk- 、 ghp_ 、 AKI) になります。

、秘密キー ブロック、JWT) がスキャンされます。何らかのものが検出された場合、投稿は中止されます (何も黙って編集されることはありません)。サニタイズはベストエフォート型であり、セキュリティを保証するものではありません。一般的なパスワードとプレフィックスのないランダムな文字列は検出されません。プレビューの確認はお客様の責任となります。
クロード コードに次のコマンドを入力します。
/handoff-revive:保存
スキルはバックグラウンドで次の手順を自動的に実行します (可能な場合はすべてゼロ LLM トークン)。
メッセージから言語を検出します (サポートされている 10 言語: en / ja / zh / zh-TW / ko / es / pt / de / fr / tr)
変更されたファイルを git ステータスから取得します - どのファイルを編集したかを覚えておく必要はありません
スキーマに準拠したハンドオフを .claude/handoff/current.md に書き込みます
スキーマを検証し、無効なパスをクリーンアップし、空のセクションを削除します — 純粋なシェル、トークンなし
チャットに推定節約額を表示します
リセット時間が経過すると、使用状況ウィンドウが再び開きます。
クロード
新しいセッションが開始されると、プラグインは最近のハンドオフを自動的に検出します。ただ実行してください:
/handoff-revive:再開
クロードはその 1 つの小さなファイルだけを読み取って、中断したところから再開します。会話全体をやり直す必要はありません。
チェックポイントへの定期的なリマインダー。デフォルトではオフになっています。 HANDOFF_CHECKPOINT_EVERY をリマインダー間のターン数 (例: 15 ) に設定して有効にします。
エクスポート HANDOFF_CHECKPOINT_EVERY=15
有効にすると、N ターンごとにクロードは次のように出力します。
[handoff-revive] ターン 15 — チェックポイント期限。 /handoff-revive:save を実行して保存します。
使用率 90% / 95% で自動保存 (UI 通知と同期)
使用量モニター フック (PostToolUse) は、rate_limits.five_hour.used_percentage を読み取ります。これは、Claude Code の「使用量制限が近づいています」通知を送信するのと同じ値です。 90% を超えると、クロードは次の応答の前に完全なハンドオフを自動保存します (要求はありません)

、中断なし）。 95% になると、緊急通知で再び節約されます。
現在のスレッドで自動保存を望まない場合は、セッションごとにオプトアウトします。
/handoff-revive:auto off # このセッションを無効にします
/handoff-revive:auto on # 再度有効にする
/handoff-revive:auto status # 現在の状態 + しきい値を表示
/handoff-revive:auto off は現在のセッションにのみ影響します。新しいセッションは ON に戻ります。自動保存を永続的にオフにするには、シェル プロファイルで環境変数としてこれらを設定します (例: ~/.zshrc または ~/.bashrc に追加して、新しいターミナルを開きます)。
エクスポート HANDOFF_AUTO_SAVE_PERCENT=無効
エクスポート HANDOFF_URGENT_PERCENT=無効
しきい値を無効にする代わりにシフトするには、同じ方法で HANDOFF_AUTO_SAVE_PERCENT=80 (より早く起動) を設定するか、 settings.json でフックごとにしきい値を構成します ( HOOK_SETUP.md を参照)。
方法
トークンをリプレイして再開する
クロード --再開
数万、多くの場合10万以上
クロード -c
数万、多くの場合10万以上
ハンドオフ復活
1,000～3,000
内部 (自動、トークンゼロ)
すべての保存と再開では、これらが純粋なシェルで実行されます。追加の入力やトークンはありません。
メタデータの共有 — 保存するたびに author 、 Branch 、 Base_commit 、 created_at がフロントマターに挿入されるため、チームメイト (または将来のあなた) は、誰が、どのブランチで、どのコミットから保存したかを確認できます。 HANDOFF_HIDE_EMAIL=1 は電子メールを省略します。
再開時の鮮度チェック — そのメタデータを現在の git 状態と比較し、ハンドオフが別のブランチの背後または上で N コミットになると警告します。また、ファイルが HANDOFF_STALE_DAYS (デフォルトは 7、0 は無効) より古い場合にも警告します。年齢チェックには git が必要ないため、プレーンなディレクトリでも機能します。純粋に情報提供。履歴書を妨げることはありません。
スナップショット履歴 — 各保存では、コピーも .claude/handoff/history/<timestamp>.md にアーカイブされます。 /handoff-revive:list はそれらを表示します、/handoff-revive:rest

ore <timestamp> brings one back (non-destructive — the current handoff is archived first), /handoff-revive:diff [timestamp] compares. Snapshots past HANDOFF_HISTORY_RETENTION_DAYS (default 30) are pruned automatically. A snapshot is the structured work state, not the conversation history.
They answer different questions — don't write the same thing into both:
Writing work state into CLAUDE.md taxes every future session with stale context; keeping it in the handoff costs nothing until you actually resume. Opt-in helper (appends a one-line guidance comment to CLAUDE.md, idempotent, never edits existing content):
bash .claude/skills/handoff-revive/scripts/setup-claude-md.sh
License
You've hit your limit · resets ... という制限通知が表示された後、 claude --resume や -c で再開すると、 それまでの会話履歴がまるごと コンテキストに再ロードされます。中規模のセッションでも数万トークン、長いセッションでは 10 万を超えるトークン分が、まだ何も質問していない段階で消費されてしまいます。
作業の続きに必要な最小限の情報だけを構造化して .claude/handoff/current.md に保存します（約 1〜3k トークン）。保存される項目:
done / wip / todo — 完了済み / 進行中 / 未着手のタスク
next_action — 実行可能な次の一手（ file:line + 具体的なコマンド）
lessons_learned — 失敗した試行と、そこから学んだこと（任意）
再開するときは claude --resume を使わず、 新規セッション を起動するだけです。skill がこの 1 ファイルだけを読み込み、すぐに作業を引き継ぎます。
プラグインインストール (推奨) — Claude Code 内で実行:
/plugin

marketplace add sofumel/claude-handoff-revive # 1. マーケットプレイス登録
/plugin install handoff-revive@handoff-revive-marketplace # 2. プラグインインストール
Skill・スラッシュコマンド・hook がすべて自動で有効化されます。 settings.json の編集は不要。
/plugin isn't available in this environment と表示された場合 、Claude Code のバージョンが古い可能性があります。 claude --version で確認し、 brew upgrade claude-code (Homebrew) または npm update -g @anthropic-ai/claude-code (npm) で更新してください。それでも使えない場合は下のマニュアルインストールを使用してください。
git clone https://github.com/sofumel/claude-handoff-revive.git
cd claude-handoff-revive
./install.sh /path/to/your-project # Linux/macOS/WSL/Git-Bash
# .\install.ps1 -Target C:\path\to\proj # Windows PowerShell
# 全プロジェクト共通でインストール:
./install.sh --global # .\install.ps1 -Global
マニュアルインストールの場合のみ、 HOOK_SETUP.md のスニペットを .claude/settings.json にマージして hooks を有効化してください（プラグインインストールでは自動）。
.claude/handoff/ を git にコミットすべき? どちらでも動きます。用途に合わせて選んでください:
個人利用（デフォルト推奨） — すべて ignore。揮発的な作業状態とスナップショットは履歴に残しません:
.claude / handoff /
チーム共有 — 現在の handoff だけコミットし、スナップショットとマーカーはローカルに留めます:
.claude / handoff / *
! .claude / handoff / current.md
.turn や .usage-flag などの内部ファイルと、 history/ （スナップショッ

トの置き場）はコミットしないでください。これらはあなたの環境だけで使う一時ファイルで、共有しても意味がありません。
PR に投稿する本文は、共有用に 自動で整えたコピー から作られます。具体的には、メールアドレスの行（ author_email ）を削除し、あなたの PC の絶対パスを <project-root> / ~ に置き換え、API キーやトークンらしき文字列（ sk- 、 ghp_ 、 AKIA 、秘密鍵、JWT など、よくある形式）が紛れ込んでいないか検査します。見つかった場合は投稿を 中止 します（勝手に伏せ字にすることはありません）。ただし これはあくまで補助で、安全を保証するものではありません 。ふつうのパスワードや、決まった形式を持たないランダムな文字列は検出できないので、投稿前のプレビュー確認はご自身でお願いします。
Claude Code で、以下のコマンドを入力してください:
/handoff-revive:save
裏では skill が以下を自動実行します（できる限りトークンを使わずに）:
言語を検出 (10 言語対応: en / ja / zh / zh-TW / ko / es / pt / de / fr / tr) — あなたのメッセージから判定
git status から変更ファイルを抽出 — どのファイルを編集したか思い出す必要なし
スキーマ通りのハンドオフを .claude/handoff/current.md に書き込み
形式チェック + 存在しないファイルへの参照を自動削除 + 空の項目を除去 — すべてシェルスクリプトだけで処理し（Claude を呼びません）、 トークンを消費しません
resets ... の時刻を過ぎたら、使用枠が回復します。
claude
新しい

セッションを開くと、プラグインが自動で最近のハンドオフを検知します。以下のコマンドを入力してください:
/handoff-revive:resume
Claude はその 小

[切り捨てられた]

## Original Extract

Resume Claude Code work after rate/usage/context limits without replaying the prior transcript. Auto-saves at 90%/95% usage. Plugin-installable, 10 languages. - sofumel/claude-handoff-revive

This is still a work in progress, and I'm sure there are many things I can improve. I'd love to hear your thoughts, suggestions, and constructive criticism.

GitHub - sofumel/claude-handoff-revive: Resume Claude Code work after rate/usage/context limits without replaying the prior transcript. Auto-saves at 90%/95% usage. Plugin-installable, 10 languages. · GitHub
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
sofumel
/
claude-handoff-revive
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows assets assets plugins/ handoff-revive plugins/ handoff-revive .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md HOOK_SETUP.md HOOK_SETUP.md LICENSE LICENSE README.md README.md install.ps1 install.ps1 install.sh install.sh View all files Repository files navigation
Continue Claude Code work after a rate / usage / context limit without replaying your prior session — typically tens of thousands of tokens, often 100k+ for long sessions.
🇺🇸 English ·
🇯🇵 日本語 ·
🇨🇳 中文 ·
🇹🇼 繁體中文 ·
🇰🇷 한국어 ·
🇪🇸 Español ·
🇵🇹 Português ·
🇩🇪 Deutsch ·
🇫🇷 Français ·
🇹🇷 Türkçe
When Claude Code shows You've hit your limit · resets ... , the standard recovery — claude --resume or claude -c — reloads the entire prior conversation into context. A medium session typically burns tens of thousands of tokens, often 100k+ for long sessions, before you've asked a single question.
It saves only the minimum needed to continue, structured into .claude/handoff/current.md (~1–3k tokens):
goal — what you were trying to accomplish
done / wip / todo — completed, in-progress, not-yet-started tasks
next_action — the concrete next step ( file:line + exact command)
touched_files — what files were touched and why
decisions — design choices, with the reason behind each
lessons_learned — failed attempts and what they taught (optional)
To resume, start a new session (don't use --resume ). The skill reads only that one file and picks up where you left off.
Plugin install (recommended) — run these inside Claude Code:
/plugin marketplace add sofumel/claude-handoff-revive # 1. register marketplace
/plugin install handoff-revive@handoff-revive-marketplace # 2. install plugin
Skill, slash commands, and hooks auto-activate. No settings.json editing required.
/plugin isn't available in this environment ? Your Claude Code version is too old. Check with claude --version and update via brew upgrade claude-code (Homebrew) or npm update -g @anthropic-ai/claude-code (npm). If you can't update, use the manual install below.
Manual install (always works, no plugin support required):
git clone https://github.com/sofumel/claude-handoff-revive.git
cd claude-handoff-revive
./install.sh /path/to/your-project # Linux/macOS/WSL/Git-Bash
# .\install.ps1 -Target C:\path\to\proj # Windows PowerShell
# Or globally for all projects:
./install.sh --global # .\install.ps1 -Global
For the manual install, optionally enable hooks by merging the snippet from HOOK_SETUP.md into your .claude/settings.json (the plugin install does this automatically).
Should .claude/handoff/ be committed to git? Both work — pick one deliberately:
Solo use (default) — ignore it all; volatile state and snapshots don't belong in history:
.claude / handoff /
Team sharing — commit only the current handoff; keep snapshots and marker files local:
.claude / handoff / *
! .claude / handoff / current.md
Never commit the dot-marker files ( .turn , .usage-flag , …) or history/ — they are machine-local state and snapshot noise.
When sharing to a PR, the body is built from a sanitized copy : the author_email line is removed, absolute project/home paths become <project-root> / ~ , and known-prefix API keys/tokens (e.g. sk- , ghp_ , AKIA , private-key blocks, JWTs) are scanned — if any are detected the post is aborted (nothing is ever silently redacted). Sanitization is best-effort, not a security guarantee : generic passwords and unprefixed random strings are NOT detected; reviewing the preview remains your responsibility.
In Claude Code, type the following command:
/handoff-revive:save
Behind the scenes, the skill runs these steps automatically (all zero-LLM-token where possible):
Detects your language (10 supported: en / ja / zh / zh-TW / ko / es / pt / de / fr / tr) from your message
Pulls changed files from git status — you don't have to remember which files you edited
Writes a schema-conformant handoff to .claude/handoff/current.md
Validates the schema, cleans up dead paths, strips empty sections — pure shell, no tokens
Displays the estimated savings in the chat
When the time on resets ... passes, your usage window opens again.
claude
When the new session starts, the plugin automatically detects the recent handoff. Just run:
/handoff-revive:resume
Claude reads only that one small file and picks up right where you left off — no need to replay the whole conversation.
A periodic reminder to checkpoint. It is off by default — enable it by setting HANDOFF_CHECKPOINT_EVERY to the number of turns between reminders (e.g. 15 ):
export HANDOFF_CHECKPOINT_EVERY=15
When enabled, every N turns Claude prints:
[handoff-revive] Turn 15 — checkpoint due. Run /handoff-revive:save to save.
Auto-save at 90% / 95% usage (synced with the UI notification)
The usage-monitor hook (PostToolUse) reads rate_limits.five_hour.used_percentage — the same value that drives Claude Code's "approaching usage limit" notification. When you cross 90% , Claude auto-saves a full handoff before its next response (no asking, no interruption). At 95% , it saves again with an urgent notice.
Per-session opt-out if you don't want auto-save in the current thread:
/handoff-revive:auto off # disable for THIS session
/handoff-revive:auto on # re-enable
/handoff-revive:auto status # show current state + thresholds
/handoff-revive:auto off only affects the current session — a new session is back to ON. To turn auto-save off permanently , set these as environment variables in your shell profile (e.g. add to ~/.zshrc or ~/.bashrc , then open a new terminal):
export HANDOFF_AUTO_SAVE_PERCENT=disabled
export HANDOFF_URGENT_PERCENT=disabled
To shift the thresholds instead of disabling, set HANDOFF_AUTO_SAVE_PERCENT=80 (fires earlier) the same way, or configure them per-hook in settings.json (see HOOK_SETUP.md ).
Method
Tokens replayed to resume
claude --resume
tens of thousands, often 100k+
claude -c
tens of thousands, often 100k+
handoff-revive
1,000–3,000
Under the hood (automatic, zero-token)
Every save and resume runs these in pure shell — no extra input, no tokens:
Sharing metadata — each save injects author , branch , base_commit , created_at into the frontmatter, so a teammate (or future you) can see who saved it, on which branch, from which commit. HANDOFF_HIDE_EMAIL=1 omits the email.
Freshness check on resume — compares that metadata against your current git state and warns when the handoff is N commits behind or on a different branch. It also warns when the file is older than HANDOFF_STALE_DAYS (default 7; 0 disables) — the age check needs no git, so it works in plain directories too. Purely informational; it never blocks the resume.
Snapshot history — each save also archives a copy to .claude/handoff/history/<timestamp>.md . /handoff-revive:list shows them, /handoff-revive:restore <timestamp> brings one back (non-destructive — the current handoff is archived first), /handoff-revive:diff [timestamp] compares. Snapshots past HANDOFF_HISTORY_RETENTION_DAYS (default 30) are pruned automatically. A snapshot is the structured work state, not the conversation history.
They answer different questions — don't write the same thing into both:
Writing work state into CLAUDE.md taxes every future session with stale context; keeping it in the handoff costs nothing until you actually resume. Opt-in helper (appends a one-line guidance comment to CLAUDE.md, idempotent, never edits existing content):
bash .claude/skills/handoff-revive/scripts/setup-claude-md.sh
License
You've hit your limit · resets ... という制限通知が表示された後、 claude --resume や -c で再開すると、 それまでの会話履歴がまるごと コンテキストに再ロードされます。中規模のセッションでも数万トークン、長いセッションでは 10 万を超えるトークン分が、まだ何も質問していない段階で消費されてしまいます。
作業の続きに必要な最小限の情報だけを構造化して .claude/handoff/current.md に保存します（約 1〜3k トークン）。保存される項目:
done / wip / todo — 完了済み / 進行中 / 未着手のタスク
next_action — 実行可能な次の一手（ file:line + 具体的なコマンド）
lessons_learned — 失敗した試行と、そこから学んだこと（任意）
再開するときは claude --resume を使わず、 新規セッション を起動するだけです。skill がこの 1 ファイルだけを読み込み、すぐに作業を引き継ぎます。
プラグインインストール (推奨) — Claude Code 内で実行:
/plugin marketplace add sofumel/claude-handoff-revive # 1. マーケットプレイス登録
/plugin install handoff-revive@handoff-revive-marketplace # 2. プラグインインストール
Skill・スラッシュコマンド・hook がすべて自動で有効化されます。 settings.json の編集は不要。
/plugin isn't available in this environment と表示された場合 、Claude Code のバージョンが古い可能性があります。 claude --version で確認し、 brew upgrade claude-code (Homebrew) または npm update -g @anthropic-ai/claude-code (npm) で更新してください。それでも使えない場合は下のマニュアルインストールを使用してください。
git clone https://github.com/sofumel/claude-handoff-revive.git
cd claude-handoff-revive
./install.sh /path/to/your-project # Linux/macOS/WSL/Git-Bash
# .\install.ps1 -Target C:\path\to\proj # Windows PowerShell
# 全プロジェクト共通でインストール:
./install.sh --global # .\install.ps1 -Global
マニュアルインストールの場合のみ、 HOOK_SETUP.md のスニペットを .claude/settings.json にマージして hooks を有効化してください（プラグインインストールでは自動）。
.claude/handoff/ を git にコミットすべき? どちらでも動きます。用途に合わせて選んでください:
個人利用（デフォルト推奨） — すべて ignore。揮発的な作業状態とスナップショットは履歴に残しません:
.claude / handoff /
チーム共有 — 現在の handoff だけコミットし、スナップショットとマーカーはローカルに留めます:
.claude / handoff / *
! .claude / handoff / current.md
.turn や .usage-flag などの内部ファイルと、 history/ （スナップショットの置き場）はコミットしないでください。これらはあなたの環境だけで使う一時ファイルで、共有しても意味がありません。
PR に投稿する本文は、共有用に 自動で整えたコピー から作られます。具体的には、メールアドレスの行（ author_email ）を削除し、あなたの PC の絶対パスを <project-root> / ~ に置き換え、API キーやトークンらしき文字列（ sk- 、 ghp_ 、 AKIA 、秘密鍵、JWT など、よくある形式）が紛れ込んでいないか検査します。見つかった場合は投稿を 中止 します（勝手に伏せ字にすることはありません）。ただし これはあくまで補助で、安全を保証するものではありません 。ふつうのパスワードや、決まった形式を持たないランダムな文字列は検出できないので、投稿前のプレビュー確認はご自身でお願いします。
Claude Code で、以下のコマンドを入力してください:
/handoff-revive:save
裏では skill が以下を自動実行します（できる限りトークンを使わずに）:
言語を検出 (10 言語対応: en / ja / zh / zh-TW / ko / es / pt / de / fr / tr) — あなたのメッセージから判定
git status から変更ファイルを抽出 — どのファイルを編集したか思い出す必要なし
スキーマ通りのハンドオフを .claude/handoff/current.md に書き込み
形式チェック + 存在しないファイルへの参照を自動削除 + 空の項目を除去 — すべてシェルスクリプトだけで処理し（Claude を呼びません）、 トークンを消費しません
resets ... の時刻を過ぎたら、使用枠が回復します。
claude
新しいセッションを開くと、プラグインが自動で最近のハンドオフを検知します。以下のコマンドを入力してください:
/handoff-revive:resume
Claude はその 小

[truncated]
