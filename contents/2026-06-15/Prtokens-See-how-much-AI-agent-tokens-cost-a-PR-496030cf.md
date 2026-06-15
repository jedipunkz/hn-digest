---
source: "https://github.com/SamuelZ12/prtokens/"
hn_url: "https://news.ycombinator.com/item?id=48545854"
title: "Prtokens – See how much AI agent tokens cost a PR"
article_title: "GitHub - SamuelZ12/prtokens: See how much your PR costs in LLM tokens (agent usage) · GitHub"
author: "SamuelZ12"
captured_at: "2026-06-15T21:56:40Z"
capture_tool: "hn-digest"
hn_id: 48545854
score: 2
comments: 0
posted_at: "2026-06-15T19:24:15Z"
tags:
  - hacker-news
  - translated
---

# Prtokens – See how much AI agent tokens cost a PR

- HN: [48545854](https://news.ycombinator.com/item?id=48545854)
- Source: [github.com](https://github.com/SamuelZ12/prtokens/)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T19:24:15Z

## Translation

タイトル: Prtokens – AI エージェント トークンの PR コストを確認する
記事タイトル: GitHub - SamuelZ12/prtokens: LLM トークンでの PR コストの確認 (エージェントの使用量) · GitHub
説明: LLM トークンで PR コストがいくらかを確認します (エージェントの使用量) - SamuelZ12/prtokens

記事本文:
GitHub - SamuelZ12/prtokens: LLM トークンで PR にかかる費用を確認する (エージェントの使用量) · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
サミュエルZ12
/
プロトークン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店 Ta

gs ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
170 コミット 170 コミット docs/ superpowers/ spec docs/ superpowers/ spec スクリプト スクリプト src src テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェント トークンの使用を、それを送信した GitHub プル リクエストに関連付けます。
prtokens は、ローカルのクロード コード、コーデックス、および OpenCode トランスクリプトを読み取り、トークンの使用状況を PR ブランチのコミットに関連付け、単一の推定コスト コメントを投稿します。集計された数値のみがマシンから出力されます。
GitHub CLI を 1 回認証します: gh auth login
PR が開いているブランチ上のリポジトリから、次のコマンドを実行します。
npxプロトークン
prtokens は、開いている PR を検索し、ローカルのトランスクリプトを読み取り、コメントを投稿または更新します。いつでも再実行して更新してください。
合計の推定コスト、トークンとセッションの数、使用されたモデル、エージェントのコスト、および折りたたみ可能なコミットごとのテーブルを示す 1 つのコメント。同じコメントは、後の実行時にその場で更新されます。複数の寄稿者がそれぞれ独自のラベル付きセクションを取得します。
セッションは、トップレベルのコーディング エージェント セッションをカウントします。 OpenCode の場合、子/サブエージェント セッションは親セッションの下にグループ化されます。
例 (コミットテーブルはデフォルトで折りたたまれています):
🪙 この PR にはトークンで $4.12 かかります
980k 入力 / 42k 出力 · 7 セッション
モデル: クロード-ソネット-4-6 、gpt-5-codex
エージェント: クロードコード ~$3.40 · コーデックス ~$0.72
コマンド
何をするのか
プロトークン
現在のブランチのオープン PR を解決し、ローカルの使用状況を読み取り、コメントを投稿または更新します。
プロトークン --pr <n>
ブランチからの自動検出ではなく、ターゲット PR 番号 <n> を指定します。
プロトークン --ドライラン
レンダリングされたコメントを標準出力に出力します。何も投稿しない。
プロトークン

--json
JSON ペイロード (レンダリングされたマークダウン、属性、価格設定、エージェントごとの合計、ソースごとの診断) を出力します。何も投稿しない。
プロトークン --verbose
また、ソースごとのリーダー診断を標準エラー出力に出力します。
プロトークンの初期化
オプションのグローバル プリプッシュ フックをインストールまたは更新します (下記を参照)。
prtokens init --dry-run
ファイルを書き込まずにフックの変更をプレビューします。
プロトークンのステータス
保留中、ブロックされている、失敗した、および最近完了した自動 PR コメント ジョブを表示します。
prtokens pr create -- <gh args>
gh pr create <gh args> を実行し、PR の作成が成功した後に prtokens コメントの投稿を試みます。
要件
GitHub CLI は gh auth ログインで認証されました
少なくとも 1 人のサポートされているエージェントからのトランスクリプト:
クロード コード — ~/.claude/projects
コーデックス — ~/.codex/sessions または ~/.codex/archived_sessions
OpenCode — ~/.local/share/opencode にある SQLite データベース
自動実行 (オプションのプリプッシュフック)
グローバルにインストールし、setup コマンドを実行して、 git Push ごとに自動的にコメントを投稿します。
npm i -g プロトークン
プロトークンの初期化
最初に prtokens init --dry-run を使用して変更をプレビューします。
prtokens init は、マネージド ブロックをグローバル プリプッシュ フックに書き込みます。絶対グローバル core.hooksPath がすでに設定されている場合、prtokens はそこにフックを書き込みます。それ以外の場合は、~/.config/git/hooks/pre-push が作成され、core.hooksPath がグローバルに設定されます。相対的な core.hooksPath は拒否されます。
prtokens はフックの PATH 上に存在する必要があります。そうでない場合は、prtokens init によって焼き付けられた絶対パスが有効なままでなければなりません。
ローカル core.hooksPath を持つリポジトリは、グローバル フックをバイパスします。
すでにプッシュされているブランチ上で gh pr create を実行しても、フックはトリガーされません。
プッシュにオープン PR がある場合、コメントはバックグラウンドですぐに投稿されます。 PR が存在する前にプッシュされた場合、prtokens は保留中のジョブを記録し、約 30 分以内に再試行します。
最も信頼性の高いワークフローを実現するには、次の方法で PR を作成します。
プリト

kens pr create -- --title " 私の PR " --body " 説明 "
-- 以降はすべて gh pr create に渡されます。 PR の作成が成功すると、prtokens は終了する前にコメントを投稿します。
手動セットアップの場合は、prtokens init --dry-run を実行して正確なマネージド ブロックとそのターゲット パスを確認し、それを自分で配置します。
トランスクリプトがマシンから離れることはありません。 PR コメントには、トークン数、推定金額、セッション数、モデル名、PR 内にすでに表示されているコミット メタデータなどの集計数値のみが含まれています。自動キューには、リポジトリ パス、リモート/ブランチ名、ヘッド SHA、タイムスタンプ、ジョブ ステータスのみが保存され、トランスクリプト、プロンプト、レンダリングされたコメント テキストは保存されません。
状況
結果
オープンな PR はありません
ヒントを出力します。 0 から出る
トランスクリプトなし
ヒントを出力します。 0 から出る
gh が見つからないか認証されていない
セットアップ手順を印刷します。 1番出口
コメント投稿失敗
手動で貼り付けるためにレンダリングされたマークダウンを印刷します。 0 から出る
PR がまだない状態でプッシュ (フック)
レコード保留中のジョブ。 0 から出る
prtokens pr create — gh が失敗する
gh 終了コードを返します。投稿しない
prtokens pr create — コメントは失敗します
エラーを出力します。 0 から出る
価格設定
コストには、可能な場合はエージェントから報告された値が使用されます。それ以外の場合、prtoken は、Claude モデル (ファーストパーティ、Bedrock、Vertex) と OpenAI GPT-5 コーディング エージェント エイリアスをカバーするバンドルされた LiteLLM 価格スナップショットを使用して API レートで見積もります。サブスクリプション ユーザーの限界費用はゼロである可能性があります。この数字はコスト認識の見積もりとして扱います。 npm run update-pricing を使用してスナップショットを更新します。
PR の LLM トークンにかかる費用を確認します (エージェントの使用量)
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
12
v0.4.7
最新の
2026 年 6 月 15 日
+ 11 リリース
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

See how much your PR costs in LLM tokens (agent usage) - SamuelZ12/prtokens

GitHub - SamuelZ12/prtokens: See how much your PR costs in LLM tokens (agent usage) · GitHub
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
SamuelZ12
/
prtokens
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
170 Commits 170 Commits docs/ superpowers/ specs docs/ superpowers/ specs scripts scripts src src tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Attribute coding-agent token usage to the GitHub pull request that shipped it.
prtokens reads your local Claude Code, Codex, and OpenCode transcripts, attributes token usage to the commits on your PR branch, and posts a single estimated-cost comment. Only aggregate numbers leave your machine.
Authenticate the GitHub CLI once: gh auth login
From a repository, on a branch with an open PR, run:
npx prtokens
prtokens finds the open PR, reads your local transcripts, and posts or updates the comment. Run it again anytime to refresh.
A single comment showing total estimated cost, token and session counts, models used, agent costs, and a collapsible per-commit table. The same comment is updated in place on later runs; multiple contributors each get their own labeled section.
Sessions counts top-level coding-agent sessions. For OpenCode, child/subagent sessions are grouped under their parent session.
Example (commit table collapsed by default):
🪙 This PR cost ~$4.12 in tokens
980k in / 42k out · 7 sessions
Models: claude-sonnet-4-6 , gpt-5-codex
Agents: claude-code ~$3.40 · codex ~$0.72
Command
What it does
prtokens
Resolve the current branch's open PR, read local usage, and post or update the comment.
prtokens --pr <n>
Target PR number <n> instead of auto-detecting from the branch.
prtokens --dry-run
Print the rendered comment to stdout; post nothing.
prtokens --json
Print a JSON payload (rendered markdown, attribution, pricing, per-agent totals, and per-source diagnostics); post nothing.
prtokens --verbose
Also print per-source reader diagnostics to stderr.
prtokens init
Install or update the optional global pre-push hook (see below).
prtokens init --dry-run
Preview the hook changes without writing files.
prtokens status
Show pending, blocked, failed, and recently completed automatic PR comment jobs.
prtokens pr create -- <gh args>
Run gh pr create <gh args> and attempt to post the prtokens comment after successful PR creation.
Requirements
GitHub CLI authenticated with gh auth login
Transcripts from at least one supported agent:
Claude Code — ~/.claude/projects
Codex — ~/.codex/sessions or ~/.codex/archived_sessions
OpenCode — SQLite databases under ~/.local/share/opencode
Automatic runs (optional pre-push hook)
Install globally and run the setup command to post comments automatically on every git push :
npm i -g prtokens
prtokens init
Preview changes first with prtokens init --dry-run .
prtokens init writes a managed block into your global pre-push hook. If an absolute global core.hooksPath is already set, prtokens writes the hook there; otherwise it creates ~/.config/git/hooks/pre-push and sets core.hooksPath globally. A relative core.hooksPath is rejected.
prtokens must be on PATH for hooks, or the absolute path baked in by prtokens init must stay valid.
A repository with a local core.hooksPath bypasses the global hook.
gh pr create on an already-pushed branch won't trigger the hook.
When a push has an open PR, the comment is posted immediately in the background. When pushed before the PR exists, prtokens records a pending job and retries within ~30 minutes.
For the most reliable workflow, create PRs through:
prtokens pr create -- --title " My PR " --body " Description "
Everything after -- is passed to gh pr create . If PR creation succeeds, prtokens posts the comment before exiting.
For manual setup, run prtokens init --dry-run to see the exact managed block and its target path, then place it yourself.
Transcripts never leave your machine. The PR comment contains only aggregate numbers — token counts, dollar estimates, session counts, model names, and commit metadata already visible in the PR. The automatic queue stores only repository path, remote/branch names, head SHA, timestamps, and job status — never transcripts, prompts, or rendered comment text.
Situation
Outcome
No open PR
Prints a hint; exits 0
No transcripts
Prints a hint; exits 0
Missing or unauthenticated gh
Prints setup instructions; exits 1
Comment post failure
Prints rendered markdown for manual paste; exits 0
Push with no PR yet (hook)
Records pending job; exits 0
prtokens pr create — gh fails
Returns gh exit code; does not post
prtokens pr create — comment fails
Prints error; exits 0
Pricing
Costs use agent-reported values when available. Otherwise, prtokens estimates at API rates using a bundled LiteLLM pricing snapshot covering Claude models (first-party, Bedrock, and Vertex) and OpenAI GPT-5 coding-agent aliases. Subscription users may have zero marginal cost — treat the number as a cost-awareness estimate. Refresh the snapshot with npm run update-pricing .
See how much your PR costs in LLM tokens (agent usage)
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
12
v0.4.7
Latest
Jun 15, 2026
+ 11 releases
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
