---
source: "https://github.com/nauta-ai/holster-scan"
hn_url: "https://news.ycombinator.com/item?id=48483102"
title: "Holster-scan – catch AI-hallucinated package imports before agents run"
article_title: "GitHub - nauta-ai/holster-scan: Local-first scanner for AI-hallucinated / typosquatted (slopsquat) packages + agent boundary preflight — catches what generic SAST misses, before an agent runs or a repo is shared. · GitHub"
author: "davidnauta"
captured_at: "2026-06-10T21:44:12Z"
capture_tool: "hn-digest"
hn_id: 48483102
score: 1
comments: 0
posted_at: "2026-06-10T21:39:51Z"
tags:
  - hacker-news
  - translated
---

# Holster-scan – catch AI-hallucinated package imports before agents run

- HN: [48483102](https://news.ycombinator.com/item?id=48483102)
- Source: [github.com](https://github.com/nauta-ai/holster-scan)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T21:39:51Z

## Translation

タイトル: ホルスタースキャン – エージェントが実行される前に AI 幻覚を起こしたパッケージのインポートを捕捉
記事のタイトル: GitHub - nauta-ai/holster-scan: AI 幻覚 / タイポスクワット (スロップスクワット) パッケージ用のローカルファースト スキャナー + エージェント境界プリフライト — エージェントが実行されるかリポジトリが共有される前に、汎用 SAST が見逃しているものをキャッチします。 · GitHub
説明: AI 幻覚 / タイポスクワット (スロップスクワット) パッケージ用のローカルファースト スキャナー + エージェント境界プリフライト — エージェントが実行される前、またはリポジトリが共有される前に、一般的な SAST が見逃したものをキャッチします。 - ナウタアイ/ホルスタースキャン

記事本文:
GitHub - nauta-ai/holster-scan: AI 幻覚 / タイポスクワット (スロップスクワット) パッケージ + エージェント境界プリフライト用のローカルファーストスキャナー — エージェントが実行されるかリポジトリが共有される前に、一般的な SAST が見逃したものをキャッチします。 · GitHub
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
なうたあい
/
ホルスト

er-スキャン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット holster_scan holster_scan テスト テスト .gitignore .gitignore MANIFEST.in MANIFEST.in README.md README.md action.yml action.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントが実行する前、またはリポジトリが共有される前に、AI エージェントが何を継承するかを確認し、幻覚やタイプポスクワッティングされたパッケージを検出します。地元。無料。サインアップはありません。
コーディング エージェント (Claude Code、Cursor、Codex) と MCP サーバーがファイル、シェル、トークン、ツールにアクセスできるようになりました。 holster-scan は、リポジトリとエージェントの構成をローカルで読み取り、2 つのことを明確に示します。これは実行しても安全か、共有しても安全か、さらに、誰もチェックしていない AI 固有のリスク、幻覚やずさんなパッケージのインポートです。
別の秘密スキャナーではありません。 AI エージェント作業の事前実行/事前共有境界チェック。
pip install holster-scan # またはソースから: pipx install git+https://github.com/nauta-ai/holster-scan
ホルスタースキャン。 # 現在のリポジトリと任意のエージェント/MCP 構成をスキャンします
完全にマシン上で実行されます。コード、構成、シークレットがそこから離れることはありません。アカウントもテレメトリもありません。
幻覚を起こした/ずさんなパッケージ — AI が作成したインポート ( reqeusts 、 langchain-utils 、 panda ) および実際のスキャナーが見逃した公開されたタイポスクワット。 (10 の主要なリアル リポジトリで 95% の再現率、0.36% の誤検知率。)
エージェントが継承するもの - シェル環境、ファイル スコープ、トークン、MCP ツール、クラウド認証。
安全に共有できるギャップ — ライブ認証情報ポインター、Git 履歴内のシークレット、許可リストのない MCP ツール。
爆発範囲 — 誤用されたエージェントまたは共有構成が実際に到達する可能性のある範囲。
修正優先の順序 — 200 の警告ではなく、5 つのステップ。
$ホー

lster-scan 。
エージェント境界レポート · client-portal-agent
走っても安全ですか？ ⚠ はい、制限あり (2 つ修正)
共有しても安全ですか? ✗ いいえ (ブロッカー 3 個)
リスク高
パッケージ
✗ リクエストの幻覚 — 「リクエスト」のタイプミス (PyPI にはありません)
✗ langchain-utils 幻覚を起こした langchain の「-utils」ヘルパー
境界
⚠ 実行ラッパーは完全なシェル環境を継承します (AWS_*、STRIPE_* がエージェントに表示されます)
⚠ MCP fs-server スコープ = $HOME、プロジェクトではありません
✗ docker-compose.override.yml → ライブキーパス /Users/.../ストライプ/live.key
✗ .env が git 履歴に存在する (回復可能)
✗ MCP ツール「shell.exec」には許可リストがありません
まず修正してください
1 参照されたライブキーを回転します 2 ラッパー環境を分離します
3 MCP fs + Shell.exec を制限します。 4 履歴から .env をスクラブします。
5 再実行 → ターゲット: 安全に実行 + 安全に共有
あなたのマシンには何も残りませんでした。
CI (GitHub アクション) 内
- 使用: nauta-ai/holster-scan@v0
付き:
failed-on : high # 重大度の高い結果に基づいてビルドを失敗させます
format : sarif # GitHub コード スキャンで結果が表示される
なぜ違うのか
AI 固有。汎用 SAST で Snyk/Semgrep と戦うのではなく、AI が生成したコードとエージェントのセットアップに特有のリスクを捉えます。
ローカルファーストの設計。分析はマシン上で実行されます。 --オフラインでも動作します。不明なパッケージにはフラグが付けられ、サイレントに渡されることはありません。
秘密だけではなく、境界線。クリーンなローカル実行は、クリーンな共有可能な成果物であることを証明するものではありません。両方をチェックします。
プライベート パッケージを許可リストに登録する — .holster.yml は内部/ベンダー インデックス パッケージを抑制するため、ノイズはゼロに近くなります。
# .holster.yml
allow : [ "torch_mlu", "internal-*" ] # プライベート/ベンダー パッケージ、フラグなし
フェイルオン: 高
レジストリ : on # PyPI の存在 + メンテナンス チェック
オープンコア。ローカルおよび CI 内で永久に無料で実行できます。
実際のクライアント リポジトリやエージェントのセットアップを詳細にレビューする人間が必要ですか?私たちは書面によるエージェント境界レビューを行います — 安全に実行できる/安全に実行できる

o-share 判定、継承マップ、優先順位付けされた修正 — 1 つのリポジトリ/構成用。ファウンダー価格は500ドル。オプションのライブウォークスルー。 → nautaai.com/holster · Nauta AI のホルスター。
AI 幻覚 / タイポスクワット (スロップスクワット) パッケージ用のローカルファースト スキャナー + エージェント境界プリフライト — エージェントが実行される前、またはリポジトリが共有される前に、一般的な SAST が見逃したものをキャッチします。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
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

Local-first scanner for AI-hallucinated / typosquatted (slopsquat) packages + agent boundary preflight — catches what generic SAST misses, before an agent runs or a repo is shared. - nauta-ai/holster-scan

GitHub - nauta-ai/holster-scan: Local-first scanner for AI-hallucinated / typosquatted (slopsquat) packages + agent boundary preflight — catches what generic SAST misses, before an agent runs or a repo is shared. · GitHub
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
nauta-ai
/
holster-scan
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits holster_scan holster_scan tests tests .gitignore .gitignore MANIFEST.in MANIFEST.in README.md README.md action.yml action.yml pyproject.toml pyproject.toml View all files Repository files navigation
See what an AI agent would inherit — and catch hallucinated/typosquatted packages — before the agent runs or the repo gets shared. Local. Free. No signup.
Coding agents (Claude Code, Cursor, Codex) and MCP servers now reach your files, shell, tokens, and tools. holster-scan reads your repo + agent config locally and tells you two things plainly: is this safe to run, and is this safe to share — plus the AI-specific risk nobody else checks: hallucinated / slopsquatted package imports.
Not another secret scanner. A pre-run / pre-share boundary check for AI-agent work.
pip install holster-scan # or from source: pipx install git+https://github.com/nauta-ai/holster-scan
holster-scan . # scan the current repo + any agent/MCP config
Runs entirely on your machine. Your code, configs, and secrets never leave it. No account, no telemetry.
Hallucinated / slopsquatted packages — AI-invented imports ( reqeusts , langchain-utils , panda ) and published typosquats that real scanners miss. (95% recall, 0.36% false-positive rate on 10 major real repos.)
What the agent inherits — shell env, file scope, tokens, MCP tools, cloud creds.
Safe-to-share gaps — live-credential pointers, secrets in git history, MCP tools with no allow-list.
Blast radius — what a misused agent or shared config could actually reach.
A fix-first order — five steps, not two hundred warnings.
$ holster-scan .
Agent Boundary Report · client-portal-agent
Safe to run? ⚠ yes, with restrictions (2 to fix)
Safe to share? ✗ no (3 blockers)
Risk HIGH
PACKAGES
✗ reqeusts hallucinated — typo of "requests" (not on PyPI)
✗ langchain-utils hallucinated "-utils" helper of langchain
BOUNDARY
⚠ run wrapper inherits full shell env (AWS_*, STRIPE_* visible to agent)
⚠ MCP fs-server scope = $HOME, not project
✗ docker-compose.override.yml → live key path /Users/.../stripe/live.key
✗ .env present in git history (recoverable)
✗ MCP tool "shell.exec" has no allow-list
FIX FIRST
1 rotate the referenced live key 2 isolate the wrapper env
3 restrict MCP fs + shell.exec 4 scrub .env from history
5 re-run → target: safe to run + safe to share
Nothing left your machine.
In CI (GitHub Action)
- uses : nauta-ai/holster-scan@v0
with :
fail-on : high # fail the build on high-severity findings
format : sarif # results show in GitHub code scanning
Why it's different
AI-specific. It's not fighting Snyk/Semgrep on generic SAST — it catches the risks unique to AI-generated code and agent setups.
Local-first by design. Analysis runs on your machine. --offline works; unknown packages are flagged, never silently passed.
Boundary, not just secrets. A clean local run doesn't prove a clean shareable artifact. It checks both.
Allow-list your private packages — .holster.yml suppresses internal/vendor-index packages so the noise stays near zero.
# .holster.yml
allow : [ "torch_mlu", "internal-*" ] # private/vendor packages, not flagged
fail_on : high
registry : on # PyPI existence + maintenance checks
Open-core. Free to run locally and in CI, forever.
Need a human to review a real client repo or agent setup in depth? We do a written Agent Boundary Review — safe-to-run / safe-to-share verdict, inheritance map, prioritized fixes — for one repo/config. Founder price $500. Optional live walkthrough. → nautaai.com/holster · Holster by Nauta AI .
Local-first scanner for AI-hallucinated / typosquatted (slopsquat) packages + agent boundary preflight — catches what generic SAST misses, before an agent runs or a repo is shared.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
