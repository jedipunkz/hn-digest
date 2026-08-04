---
source: "https://github.com/mountainowl/bubo"
hn_url: "https://news.ycombinator.com/item?id=49164862"
title: "Bubo: AI code-reviewer that learns from review comments"
article_title: "GitHub - mountainowl/bubo: Agentic AI code review for GitLab MRs and GitHub PRs, with the LLM of your choice. Posts only actionable findings as inline review threads. · GitHub"
author: "mt_owl"
captured_at: "2026-08-04T06:24:48Z"
capture_tool: "hn-digest"
hn_id: 49164862
score: 1
comments: 1
posted_at: "2026-08-04T06:12:11Z"
tags:
  - hacker-news
  - translated
---

# Bubo: AI code-reviewer that learns from review comments

- HN: [49164862](https://news.ycombinator.com/item?id=49164862)
- Source: [github.com](https://github.com/mountainowl/bubo)
- Score: 1
- Comments: 1
- Posted: 2026-08-04T06:12:11Z

## Translation

タイトル: Bubo: レビューコメントから学習する AI コードレビューワー
記事のタイトル: GitHub - Mountainowl/bubo: 選択した LLM を使用した、GitLab MR および GitHub PR のための Agentic AI コード レビュー。実用的な調査結果のみをインライン レビュー スレッドとして投稿します。 · GitHub
説明: 選択した LLM を使用した、GitLab MR および GitHub PR の Agentic AI コード レビュー。実用的な調査結果のみをインライン レビュー スレッドとして投稿します。 - マウンテンフクロウ/横痃

記事本文:
GitHub - Mountainowl/bubo: 選択した LLM を使用した、GitLab MR および GitHub PR のための Agentic AI コード レビュー。実用的な調査結果のみをインライン レビュー スレッドとして投稿します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ヤマフクロウ
/
横痃
公共
ノーティ

fications
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
183 コミット 183 コミット .github .github アセット アセット ベンチマーク ベンチマーク bin bin config デプロイ/ テンプレート デプロイ/ テンプレート docs docs プラグイン/ スーパーパワー プラグイン/ スーパーパワー プロンプト プロンプト スキル/ コードレビューアー スキル/ コードレビューアー src/ bubo src/ bubo テスト テスト ui ui .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .release-please-manifest.json .release-please-manifest.json CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md action.yml action.yml pyproject.toml pyproject.toml release-please-config.json release-please-config.json uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
選択した LLM を使用した Agentic AI コード レビュー。 Bubo reviews your GitLab
MR と GitHub PR は実行したモデルを使用し、価値のある調査結果のみを投稿します
インライン スレッドとして動作します。チャットボットのノイズ、賞賛、要約はありません。
自己ホスト型 — コード、差分、レビュー データはインフラストラクチャ上に留まります
Bring-your-own-LLM — Codex、Claude、または CLI が駆動する任意のモデル
SCM — 現在、Gitlab と Github をサポートしています
所見 — インライン、またはクリーンな場合は「すべて良好」
ガバナンス、来歴、監査可能なオンプレミス レポート - SBOM を使用した共同署名リリース
完全なドキュメント → Mountainowl.github.io/bubo
uv ツール install bubo # または: pipx install bubo
bubo init # idempotent;シード構成 + ワークスペース + DB
横痃医師 # 最初のポーリングの前に確認します
bubo-poller # 1 つのポーリング サイクル — デフォルトでは予行演習を行い、何も投稿しません
Prefer a cont

ainer? docker pull ghcr.io/mountainowl/bubo (マルチアーチ;
review-agent CLI は BYO です)。 Continue with the
recipes and quickstart .
すべてはドキュメント サイトにあります。この README は玄関口にすぎません。
ポーリングによる GitLab および GitHub の投稿 — 成果メトリクスでの本番パス
parity. [scm].provider = "github" (または BUBO_PROVIDER=github ) を設定します。
MCP サーバー ( bubo-mcp ) — 読み取り専用メトリクス + トリガーされたレビュー。 stdio or HTTP.
Codex または Claude — Bubo はエージェントのラッパーを通じてレビューを実行します
CLI; Codex ships pre-wired.
Webhook 主導のトリガー - まだです。ポーリングが唯一のパスです。
レビューの実行は設計により CI/CD の外部にあります。それをポーラーとして実行します。
existing pipelines.
config/env.toml は gitignored であり、トークンを保持します。実際の値を出力したりコミットしたりしないでください。
Review-agent の標準出力が編集されています ( GITLAB_TOKEN= 、 OPENAI_API_KEY= 、 glpat-… 、
sk-… 、認証済みの Git URL) がレポート、ログ、データベースにアクセスする前に実行されます。
レビューアーのサブプロセスは厳密な環境許可リストの下で実行されます。ホスト シークレットは許可されません。
LLM代理店に卸しました。
リリースは Sigstore キーレス OIDC を介して署名され、リリースごとに SBOM が付いています。
SECURITY.md に従って脆弱性を報告してください。
貢献 · セキュリティ ポリシー · サポート · 行動規範 · ライセンス: MIT
選択した LLM を使用した、GitLab MR および GitHub PR の Agentic AI コード レビュー。実用的な調査結果のみをインライン レビュー スレッドとして投稿します。
Mountainowl.github.io/bubo/ トピック
Readme MIT ライセンスの行動規範
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Agentic AI code review for GitLab MRs and GitHub PRs, with the LLM of your choice. Posts only actionable findings as inline review threads. - mountainowl/bubo

GitHub - mountainowl/bubo: Agentic AI code review for GitLab MRs and GitHub PRs, with the LLM of your choice. Posts only actionable findings as inline review threads. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
mountainowl
/
bubo
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
183 Commits 183 Commits .github .github assets assets benchmarks benchmarks bin bin config config deploy/ templates deploy/ templates docs docs plugins/ superpowers plugins/ superpowers prompts prompts skills/ code-reviewer skills/ code-reviewer src/ bubo src/ bubo tests tests ui ui .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .release-please-manifest.json .release-please-manifest.json CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md action.yml action.yml pyproject.toml pyproject.toml release-please-config.json release-please-config.json uv.lock uv.lock View all files Repository files navigation
Agentic AI code review with the LLM of your choice. Bubo reviews your GitLab
MRs and GitHub PRs with the model you run, and posts only the findings worth
acting on as inline threads — no chatbot noise, no praise, no summaries.
Self-hosted — code, diffs, and review data stay on your infrastructure
Bring-your-own-LLM — Codex, Claude, or any model your CLI drives
SCM — Currently supports Gitlab and Github
findings — Inline or "all good" if clean
Governance, provenance & an auditable on-prem report - cosign-signed releases with SBOMs
Full documentation → mountainowl.github.io/bubo
uv tool install bubo # or: pipx install bubo
bubo init # idempotent; seeds config + workspace + DB
bubo doctor # verify before the first poll
bubo-poller # one poll cycle — dry-run by default, posts nothing
Prefer a container? docker pull ghcr.io/mountainowl/bubo (multi-arch; the
review-agent CLI is BYO). Continue with the
recipes and quickstart .
Everything lives on the docs site — this README is just the front door.
GitLab & GitHub posting via polling — production path, at outcome-metric
parity. Set [scm].provider = "github" (or BUBO_PROVIDER=github ).
MCP server ( bubo-mcp ) — read-only metrics + triggered reviews; stdio or HTTP.
Codex or Claude — Bubo runs the review through a wrapper around your agent
CLI; Codex ships pre-wired.
Webhook-driven triggering — not yet; polling is the only path.
Review execution sits outside CI/CD by design — run it as a poller beside your
existing pipelines.
config/env.toml is gitignored and holds tokens. Do not print or commit real values.
Review-agent stdout is redacted ( GITLAB_TOKEN= , OPENAI_API_KEY= , glpat-… ,
sk-… , credentialed Git URLs) before it touches reports, logs, or the database.
The reviewer subprocess runs under a strict env allowlist — host secrets aren't
handed wholesale to the LLM agent.
Releases are cosign-signed via Sigstore keyless OIDC, with an SBOM on every release.
Report vulnerabilities per SECURITY.md .
Contributing · Security policy · Support · Code of conduct · License: MIT
Agentic AI code review for GitLab MRs and GitHub PRs, with the LLM of your choice. Posts only actionable findings as inline review threads.
mountainowl.github.io/bubo/ Topics
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
