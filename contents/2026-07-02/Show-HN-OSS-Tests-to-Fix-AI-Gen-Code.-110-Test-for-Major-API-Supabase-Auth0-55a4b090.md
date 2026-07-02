---
source: "https://github.com/qualtyco/api-doctor"
hn_url: "https://news.ycombinator.com/item?id=48762860"
title: "Show HN: OSS Tests to Fix AI Gen Code. 110 Test for Major API – Supabase, Auth0"
article_title: "GitHub - qualtyco/api-doctor: AI compiles hallucinated code that pass. This fixes it before accepting it. 100% Deterministic · GitHub"
author: "Reuben_Santoso"
captured_at: "2026-07-02T16:00:59Z"
capture_tool: "hn-digest"
hn_id: 48762860
score: 1
comments: 0
posted_at: "2026-07-02T15:11:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: OSS Tests to Fix AI Gen Code. 110 Test for Major API – Supabase, Auth0

- HN: [48762860](https://news.ycombinator.com/item?id=48762860)
- Source: [github.com](https://github.com/qualtyco/api-doctor)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T15:11:11Z

## Translation

タイトル: HN を表示: AI 生成コードを修正するための OSS テスト。 110 主要 API のテスト – Supabase、Auth0
記事タイトル: GitHub - qualtyco/api-doctor: AI が合格する幻覚コードをコンパイルします。これにより、受け入れる前に修正されます。 100% 決定論的 · GitHub
説明: AI は通過する幻覚コードをコンパイルします。これにより、受け入れる前に修正されます。 100% 決定的 - qualtyco/api-doctor
HN テキスト: クロード コードのパワー ユーザーとしては、コンテキストが圧縮されると、だらだらと書き始めるでしょう。特に API 統合用のコードを作成する場合。たとえば、Supabase 認証では、クライアントがユーザー メタデータを書き込み可能にし、next.js コンポーネントにサービス ロール キーを残すことができます。これを理解するためにプロンプ​​トとskills.mdを試してみましたが、適切なテストフィードバックがないと一貫性がありません。そこで、悪いコードの統合を見つけて修正するために、公式ドキュメントから LLM を含まないテストのコレクションをオープンソース化しました。楽しむ

記事本文:
GitHub - qualtyco/api-doctor: AI は合格する幻覚コードをコンパイルします。これにより、受け入れる前に修正されます。 100% 決定論的 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
クォリティコ
/
APIドクター
公共
通知

オン
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
49 コミット 49 コミット アセット アセット スキル/ api-doctor スキル/ api-doctor src src .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE.md LICENSE.md README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
api-doctor は、AI が生成したコードをスキャンして、不正な API 統合を検出します。
決定論的な AST ルール。プロンプトではありません。毎回同じ入力、同じ出力。
→ 詳細と例は apidoctor.co にあります
# プロジェクトをスキャンします
npx @api-doctor/cli 。
# またはエージェントスキルとしてインストール (クロードコード、カーソル、ウィンドサーフィン)
npx @api-doctor/cli インストール
📦 サポートされているプロバイダー
プロバイダー
ルール
再送信
13のルール
スーパーベース
12のルール
認証0
4つのルール
ファイアベース
20のルール
愛らしい
4つのルール
ブラウザベース
11のルール
OpenAIコンピュータの使用
7つのルール
チップタップ
11のルール
イレブンラボ
10のルール
トゥイリオ
9つのルール
OpenAIリアルタイム
9つのルール
完全なルール カタログは、GitHub リポジトリの src/providers/<name>/README.md にあります。
ルールは、セキュリティ (CWE/OWASP マッピング)、正確性 (間違ったエンドポイント)、信頼性 (本番の障害モード)、および統合 (配線ギャップ) の 4 つのカテゴリをカバーしています。
AI コードを AI でテストすることはできません。 api-doctor はそのループを打ち破ります。毎回同じルール、同じ出力。モデルコールではありません。プロンプトではありません。
api-doctor は匿名の使用状況データを PostHog に送信するため、このツールが開発者が実際のバグを捕捉するのに役立っているかどうかを確認できます。
CLIバージョン、Node.jsバージョン、プラットフォーム
実行コンテキスト: ローカル、CI、またはエージェント
検出された API SDK (例: resend 、 supabase ) — プロバイダー名のみ
どのルールが起動されたか - ルール名のみ、コードなし
実行間のスコアデルタ

同じプロジェクト上 (そのプロジェクトの .api-doctor/run-history.json にローカルに保存されます)
ハッシュされたプロジェクト識別子 ( project_hash ) — パス自体ではなく、スキャンされたディレクトリ パスの SHA-256
予期せぬクラッシュ時のエラー メッセージとスタック トレースをサニタイズ (パスは編集)
RAW ファイルのパスまたはプロジェクト名
電子メール、名前、または個人を特定できる情報
ランダムな匿名 ID は ~/.api-doctor/install-id に保存されます。プロジェクトごとの実行履歴は、 <project>/.api-doctor/run-history.json に保存されます。どちらもマシン上に残ります。上記のイベント データのみが PostHog に送信されます。
npx @api-doctor/cli 。 --テレメトリなし
npx @api-doctor/cli install --no-telemetry
または、環境で API_DOCTOR_TELEMETRY=0 または DO_NOT_TRACK=1 を設定します。
AI は通過する幻覚コードをコンパイルします。これにより、受け入れる前に修正されます。 100% 決定的
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
3
フォーク
レポートリポジトリ
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI compiles hallucinated code that pass. This fixes it before accepting it. 100% Deterministic - qualtyco/api-doctor

As a claude code power user once context compacts it will start writing slop. Especially when writing code for API integrations. For example Supabase auth, it allows user metadata to be writable by clients and leaving service role keys in next.js components. I’ve tried prompting and skills.md to catch this, but it’s not consistent without good test feedback. So I open sourced a collection of test to with no LLM involved from official docs to catch and fix bad code integrations. Enjoy

GitHub - qualtyco/api-doctor: AI compiles hallucinated code that pass. This fixes it before accepting it. 100% Deterministic · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
qualtyco
/
api-doctor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
49 Commits 49 Commits assets assets skills/ api-doctor skills/ api-doctor src src .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE.md LICENSE.md README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
api-doctor scans AI-generated code for bad API integrations.
Deterministic AST rules. Not a prompt. Same input, same output, every time.
→ Full story and examples at apidoctor.co
# Scan your project
npx @api-doctor/cli .
# Or install as an agent skill (Claude Code, Cursor, Windsurf)
npx @api-doctor/cli install
📦 Supported Providers
Provider
Rules
Resend
13 rules
Supabase
12 rules
Auth0
4 rules
Firebase
20 rules
Lovable
4 rules
Browserbase
11 rules
OpenAI Computer Use
7 rules
TipTap
11 rules
ElevenLabs
10 rules
Twilio
9 rules
OpenAI Realtime
9 rules
Full rule catalogs live in the GitHub repo under src/providers/<name>/README.md .
Rules cover four categories: security (CWE/OWASP mapped), correctness (wrong endpoints), reliability (production failure modes), and integration (wiring gaps).
You can't test AI code with AI. api-doctor breaks that loop. Same rules, same output, every time. Not a model call. Not a prompt.
api-doctor sends anonymous usage data to PostHog so we can see whether the tool is helping developers catch real bugs.
CLI version, Node.js version, platform
Run context: local, CI, or agent
Which API SDKs were detected (e.g. resend , supabase ) — provider names only
Which rules fired — rule names only, no code
Score delta between runs on the same project (stored locally in that project's .api-doctor/run-history.json )
A hashed project identifier ( project_hash ) — SHA-256 of the scanned directory path, not the path itself
Sanitized error messages and stack traces on unexpected crashes (paths redacted)
Raw file paths or project names
Email, name, or any personally identifying information
A random anonymous ID is stored at ~/.api-doctor/install-id . Per-project run history is stored at <project>/.api-doctor/run-history.json . Both stay on your machine — only the event data above is sent to PostHog.
npx @api-doctor/cli . --no-telemetry
npx @api-doctor/cli install --no-telemetry
Or set API_DOCTOR_TELEMETRY=0 or DO_NOT_TRACK=1 in your environment.
AI compiles hallucinated code that pass. This fixes it before accepting it. 100% Deterministic
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
3
forks
Report repository
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
