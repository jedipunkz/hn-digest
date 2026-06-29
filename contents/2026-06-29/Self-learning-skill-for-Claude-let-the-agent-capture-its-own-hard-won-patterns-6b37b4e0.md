---
source: "https://github.com/Kulaxyz/self-learning-skills"
hn_url: "https://news.ycombinator.com/item?id=48713137"
title: "Self-learning skill for Claude: let the agent capture its own hard-won patterns"
article_title: "GitHub - Kulaxyz/self-learning-skills: A self-improving skill for AI coding agents (Claude Code, Cursor, AGENTS.md): recognize a hard-won golden path in a session and harvest it into a reusable skill/rule for next time. · GitHub"
author: "kulaxyz"
captured_at: "2026-06-29T00:31:02Z"
capture_tool: "hn-digest"
hn_id: 48713137
score: 1
comments: 0
posted_at: "2026-06-29T00:04:54Z"
tags:
  - hacker-news
  - translated
---

# Self-learning skill for Claude: let the agent capture its own hard-won patterns

- HN: [48713137](https://news.ycombinator.com/item?id=48713137)
- Source: [github.com](https://github.com/Kulaxyz/self-learning-skills)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T00:04:54Z

## Translation

タイトル: クロードの自己学習スキル: エージェントが苦労して獲得した独自のパターンを捕捉させます
記事のタイトル: GitHub - Kulaxyz/self-learning-skills: AI コーディング エージェントのための自己改善スキル (Claude Code、Cursor、AGENTS.md): セッションで苦労して獲得した黄金の道を認識し、それを次回のために再利用可能なスキル/ルールに収穫します。 · GitHub
説明: AI コーディング エージェント用の自己改善スキル (クロード コード、カーソル、AGENTS.md): セッションで苦労して獲得した黄金の道を認識し、それを次回のために再利用可能なスキル/ルールに収集します。 - Kulaxyz/自己学習スキル

記事本文:
GitHub - Kulaxyz/self-learning-skills: AI コーディング エージェント用の自己改善スキル (Claude Code、Cursor、AGENTS.md): セッションで苦労して獲得した黄金の道を認識し、それを次回のために再利用可能なスキル/ルールに収集します。 · GitHub
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
クラクシズ
/
s

エルフの学習スキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .claude-plugin .claude-plugin .cursor/ rules .cursor/ rules skill/ self-learning skill/ self-learning .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md skill.sh.json skill.sh.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントの自己改善スキル。クロードコード、カーソル、
および AGENTS.md/stand-instructions ファイルを読み取るエージェント。
セッションごとにハードなデバッグを行ったり、同じことを再発見したりします。どうすればよいですか
製品 DB に到達しますか?信条はどこに住んでいますか？デプロイコマンドとは何ですか？どうすればいいですか
これをライブで検証しますか？ — そして苦労して得た知識はセッションが終わると蒸発してしまいます
終わります。次のセッションはゼロからスタートして再学習します。
独学でそれを解決します。エージェントにその瞬間を認識するよう教えます。
再利用可能なゴールデン パスを取得し、ツールが保存する場所にそれを永続化します。
次回はそれを自動ロードします - そのため、次のセッションはすでにルートを認識して開始されます
それを再発見するのではなく。
これはメタスキルです。仕事を実行するのではなく、仕事がどのように行われたかを捉えるのです。
完了 — 既知のデッドエンドの次のセッションをスキップしたため、失敗を含む
多くの場合、勝利そのものよりも価値があります。
瞬間を認識する — 何度か試して初めてうまくいくタスク、
自明ではないコマンド、事前に知らなかったプロジェクトの事実、運用上の注意事項
ワークフローが繰り返し発生する可能性があるか、単に「これを覚えておいてください」と言っているだけです。
それをキャプチャします。プロンプトは必要ありません。すぐにキューに作用し、キューを選択します。
スコープ/名前自体が表示され、後で通知されます。プロシージャはキャプチャされます (キャプチャされません)
1 回限りの回答)、さらに「うまくいかなかったこと」のメモ。
再利用 — 次へ

セッションでは、スキル/ルールによってエントリが自動的にロードされます
説明、または指示ファイルが常に読み取られるためです。
ツールごとに異なるのは、知識がどこに保持され、どのように保持されるかだけです。
自動ロード:
npx — 推奨 (70 以上のエージェントで動作)
コミュニティ スキル CLI を使用します。
検出したエージェント (Claude Code、Cursor、Codex、Cline、
OpenCode など:
npx skill add kulaxyz/self-learning-skills # このプロジェクト (エージェントを自動検出)
npx skill add kulaxyz/self-learning-skills -g # global — すべてのプロジェクト
npx スキル追加 kulaxyz/self-learning-skills -a claude-code # 特定のエージェント
インストールせずに一度試してみてください。
npx スキルは kulaxyz/self-learning-skills --skill self-learning を使用します。クロード
クロードコードプラグイン
/プラグイン マーケットプレイス add kulaxyz/self-learning-skills
/プラグインのインストール self-learning@self-learning-skills
マニュアル
自分でファイルを所定の場所にコピーする
git clone https://github.com/kulaxyz/self-learning-skills
# Claude コード — グローバル (または、git 経由で共有するプロジェクトの .claude/skills/ に)
cp -R 自己学習-スキル/スキル/自己学習 ~ /.claude/skills/
# Cursor — .cursor/rules/ を自動ロードします (収集されたルールは .cursor/rules/learned/ に配置されます)
mkdir -p .cursor/rules
cp self-learning-skills/.cursor/rules/self-learning.mdc .cursor/rules/
# 任意の AGENTS.md エージェント (Codex、Zed、Aider、Gemini CLI など)
カール https://raw.githubusercontent.com/kulaxyz/self-learning-skills/main/AGENTS.md >> AGENTS.md
トリアージ: スキル、記憶、それともスキップ?
ワンライナーで構成が肥大化することはありません。各レッスンは次のようにルーティングされます。
収集されたスキル/ルールはコミットされて共有されるため、決して共有されないように構築されています。
シークレット値を書き込みます。パスワード、トークン、接続文字列、API キーは書き込みません。
シークレットの場所 (環境変数名、クライアント/セレクター) のみを記録します。
機能、MCP ツール、シークレット マネージャー)。 R

秘密を共有に再生成する
ファイルが漏洩します。
自己学習スキル/
§── AGENTS.md # ループの汎用、クロスツール バージョン
§── skill.sh.json # `npx skill` のレジストリ マニフェスト / skill.sh
§── .claude-plugin/
│ ━──marketplace.json # クロードコードプラグインマニフェスト
§── スキル/
│ └─ 自己学習/ # エージェントスキル標準（クロードコード＋クライアント）
│ §── SKILL.md # 瞬間認識 + 収穫手順
│ §── 参考文献/
│ │ └── skill-authoring.md # 優れたスキルを作成するためにライターがロードする凝縮された仕様
│ └── 資産/
│ └── SKILL.template.md # 採取したスキルの記入テンプレート
━━ .cursor/
└── ルール/
§── self-learning.mdc # カーソルアダプタ（常に適用されるルール）
└── 学習/ # 収穫 カーソル ルールがここに配置されます
オープンなエージェント スキル標準に基づいて構築されています。
AI コーディング エージェントの自己改善スキル (クロード コード、カーソル、AGENTS.md): セッションで苦労して獲得した黄金の道を認識し、それを次回のために再利用可能なスキル/ルールに収集します。
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

A self-improving skill for AI coding agents (Claude Code, Cursor, AGENTS.md): recognize a hard-won golden path in a session and harvest it into a reusable skill/rule for next time. - Kulaxyz/self-learning-skills

GitHub - Kulaxyz/self-learning-skills: A self-improving skill for AI coding agents (Claude Code, Cursor, AGENTS.md): recognize a hard-won golden path in a session and harvest it into a reusable skill/rule for next time. · GitHub
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
Kulaxyz
/
self-learning-skills
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .claude-plugin .claude-plugin .cursor/ rules .cursor/ rules skills/ self-learning skills/ self-learning .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md skills.sh.json skills.sh.json View all files Repository files navigation
A self-improving skill for AI coding agents. Works with Claude Code, Cursor,
and any agent that reads an AGENTS.md / standing-instructions file.
Every session you do hard debugging or rediscover the same thing — how do I
reach the prod DB? where do the creds live? what's the deploy command? how do I
verify this live? — and that hard-won knowledge evaporates when the session
ends. The next session starts from zero and re-learns it.
self-learning fixes that. It teaches your agent to recognize the moment it
has just earned a reusable golden path and persist it where the tool will
auto-load it next time — so the next session starts already knowing the route
instead of rediscovering it.
It's a meta-skill : it doesn't do the work, it captures how the work got
done — including the failures , since skipping a known dead-end next session
is often worth more than the win itself.
Recognize the moment — a task that only worked after several tries, a
non-obvious command, a project fact you didn't know up front, an operational
workflow likely to recur, or you simply saying "remember this" .
Capture it, no prompt needed — it acts on the cue immediately, picks the
scope/name itself, and tells you afterward. The procedure is captured (not
a one-off answer), plus a "what didn't work" note.
Reuse — next session the entry loads automatically, by skill/rule
description or because the instructions file is always read.
What differs per tool is only where knowledge is persisted and how it's
auto-loaded:
npx — recommended (works with 70+ agents)
Uses the community skills CLI, which
installs into whatever agents it detects — Claude Code, Cursor, Codex, Cline,
OpenCode, and more:
npx skills add kulaxyz/self-learning-skills # this project (auto-detects agents)
npx skills add kulaxyz/self-learning-skills -g # global — all your projects
npx skills add kulaxyz/self-learning-skills -a claude-code # a specific agent
Try it once without installing:
npx skills use kulaxyz/self-learning-skills --skill self-learning | claude
Claude Code plugin
/plugin marketplace add kulaxyz/self-learning-skills
/plugin install self-learning@self-learning-skills
Manual
Copy the files into place yourself
git clone https://github.com/kulaxyz/self-learning-skills
# Claude Code — global (or into a project's .claude/skills/ to share via git)
cp -R self-learning-skills/skills/self-learning ~ /.claude/skills/
# Cursor — auto-loads .cursor/rules/ (harvested rules land in .cursor/rules/learned/)
mkdir -p .cursor/rules
cp self-learning-skills/.cursor/rules/self-learning.mdc .cursor/rules/
# Any AGENTS.md agent (Codex, Zed, Aider, Gemini CLI, …)
curl https://raw.githubusercontent.com/kulaxyz/self-learning-skills/main/AGENTS.md >> AGENTS.md
Triage: skill, memory, or skip?
It won't bloat your config with one-liners. Each lesson is routed:
Harvested skills/rules get committed and shared, so this is built to never
write secret values — no passwords, tokens, connection strings, or API keys.
It records only where to find a secret (env var name, a client/selector
function, an MCP tool, a secret manager). Reproducing a secret into a shared
file leaks it.
self-learning-skills/
├── AGENTS.md # generic, cross-tool version of the loop
├── skills.sh.json # registry manifest for `npx skills` / skills.sh
├── .claude-plugin/
│ └── marketplace.json # Claude Code plugin manifest
├── skills/
│ └── self-learning/ # Agent Skills standard (Claude Code + clients)
│ ├── SKILL.md # recognize-the-moment + harvest procedure
│ ├── references/
│ │ └── skill-authoring.md # condensed spec the writer loads to author a good skill
│ └── assets/
│ └── SKILL.template.md # fill-in template for harvested skills
└── .cursor/
└── rules/
├── self-learning.mdc # Cursor adapter (always-applied rule)
└── learned/ # harvested Cursor rules land here
Built on the open Agent Skills standard.
A self-improving skill for AI coding agents (Claude Code, Cursor, AGENTS.md): recognize a hard-won golden path in a session and harvest it into a reusable skill/rule for next time.
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
