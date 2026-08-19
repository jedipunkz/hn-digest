---
source: "https://github.com/AdanRott/magnificent-jobs-plugin"
hn_url: "https://news.ycombinator.com/item?id=49362280"
title: "Find jobs with Claude Code CLI"
article_title: "GitHub - AdanRott/magnificent-jobs-plugin: Magnificent Jobs for Claude — find live US jobs by describing the role you want (MCP server + skill) · GitHub"
image: "https://opengraph.githubassets.com/6accf8f2addbeefed32e0fb2ef1c7642ce0d59fc809c1cd2b1a6a7aa9ebe751f/AdanRott/magnificent-jobs-plugin"
author: "DamiandeGroot"
captured_at: "2026-08-19T15:22:27Z"
capture_tool: "hn-digest"
hn_id: 49362280
score: 1
comments: 0
posted_at: "2026-08-19T14:39:45Z"
tags:
  - hacker-news
  - translated
---

# Find jobs with Claude Code CLI

- HN: [49362280](https://news.ycombinator.com/item?id=49362280)
- Source: [github.com](https://github.com/AdanRott/magnificent-jobs-plugin)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T14:39:45Z

## Translation

タイトル: Claude Code CLI で仕事を探す
記事のタイトル: GitHub - AdanRott/magnificent-jobs-plugin: Magnificent Jobs for Claude — 希望する役割 (MCP サーバー + スキル) を説明して米国のライブジョブを検索 · GitHub
説明: Magnificent Jobs for Claude — 希望する役割 (MCP サーバー + スキル) を説明することで、米国のライブジョブを検索します - AdanRott/magnificent-jobs-plugin

記事本文:
GitHub - AdanRott/magnificent-jobs-plugin: Magnificent Jobs for Claude — 希望する役割 (MCP サーバー + スキル) を説明して、米国のライブジョブを検索 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アダンロット
/
素晴らしいジョブプラグイン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
16 コミット 16 コミット フォルダーとファイル
.claude-plugin .claude-plugin アセット アセット cli cli スキル/ジョブ検索スキル/ジョブ検索 .mcp.json .mcp.json

Dockerfile Dockerfile ライセンス ライセンス README.md README.md Glama.json Glama.json server.json server.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ライブジョブ: 3,334,183 · 今日追加: 211,021 · 米国
クロードコードとコーデックスを使用して、CLI で仕事を見つけます。私たちはインターネットを毎時間検索して、リンクされている求人を見つけますが、実際には見つかりません。それも無料です
npx grandjobs セットアップ # クロード コード · コーデックス
npx grandjobs # ターミナル
インストール
/plugin マーケットプレイス add AdanRott/magnificent-jobs-plugin
/plugin install grande-jobs@magnificent-jobs
または、Claude.ai / Claude Desktop で、カスタム コネクタとして追加します: [設定] → [コネクタ] → [追加]
カスタム コネクタ → https://magnificentjobs.com/mcp (認証なし)。
サーバーはリモートで読み取り専用であり、キーは必要ありません。公式 MCP レジストリには次のようにリストされています。
com.magnificentjobs/jobs 。ストリーミング可能な HTTP をサポートするクライアントに追加します。
{
"mcpサーバー": {
"magnificent-jobs" : { "type" : " http " , "url" : " https://magnificentjobs.com/mcp " }
}
}
クロード コード: claude mcp add --transport http壮麗なジョブ https://magnificentjobs.com/mcp
stdio 専用クライアント / Docker: npx grandjobs mcp は、stdio 経由で MCP を話し、ホストされているサーバーにプロキシします ( docker build -t grandjobs .&& docker run -i grandjobs )
カーソル / ウィンドサーフィン / VS コード: 上記の JSON を MCP 設定に貼り付けます。
ChatGPT: Plugins ディレクトリから Magnificent Jobs をインストールするか、開発者モードで https://magnificentjobs.com/mcp を追加します。
Claude.ai / デスクトップ: 設定 → コネクタ → カスタム コネクタの追加 → https://magnificentjobs.com/mcp 。
ツール: search_jobs · get_job · find_cities · list_states 。ドキュメント: https://magnificentjobs.com/developers 。
コンポーネント
何をするのか
MCP サーバー https://magnificentjobs.com/mcp (ストリーミング可能な HTTP、認証なし)
search_jobs (セマンティック検索;

都市半径 10/30/50 マイル、州、リモート/オンサイト、経験範囲、雇用形態)、get_job (完全な説明、給与、スキル、リンクの適用)、find_cities 、list_states 。すべてのツールは読み取り専用です。
スキルを使った仕事の検索
クエリの表現方法 (完全なタイトル + 一文 + スキル)、いつどのツールを使用するか、米国のみの報道、決して仕事を発明しない、応募すると主張しない。
試してみる
「リモートのシニア React エンジニアの仕事を探して、給与が最も高い 3 つを教えてください。」
「フレズノから 16 マイル以内にある倉庫またはフォークリフトの仕事は何ですか? 最も近い求人の応募リンクを教えてください。」
「コロラド州の初級レベルのマーケティングの仕事を表示し、上位 2 つを要約します。」
「現在、看護師の仕事が最も多い州はどこですか?」
検索は、結果を返すために壮麗なjobs.comに送信され、標準サーバーを超えて保存されることはありません
ログ。アカウントも個人データもありません。完全なポリシー: https://magnificentjobs.com/privacy ·
規約: https://magnificentjobs.com/terms · ドキュメント: https://magnificentjobs.com/developers ·
サポート: 情報 --- で --- grandjobs.com
MIT — 「ライセンス」を参照してください。プラグイン ファイルは MIT です。求人データは、その規約に基づいて、grandlyjobs.com によって提供されます。
クロードの素晴らしい求人 — 希望する役割 (MCP サーバー + スキル) を説明して、米国での現在の求人を検索します。
grandjobs.com/developers リソース
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Magnificent Jobs for Claude — find live US jobs by describing the role you want (MCP server + skill) - AdanRott/magnificent-jobs-plugin

GitHub - AdanRott/magnificent-jobs-plugin: Magnificent Jobs for Claude — find live US jobs by describing the role you want (MCP server + skill) · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
AdanRott
/
magnificent-jobs-plugin
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
16 Commits 16 Commits Folders and files
.claude-plugin .claude-plugin assets assets cli cli skills/ find-jobs skills/ find-jobs .mcp.json .mcp.json Dockerfile Dockerfile LICENSE LICENSE README.md README.md glama.json glama.json server.json server.json View all files Repository files navigation
LIVE JOBS: 3,334,183 · ADDED TODAY: 211,021 · US
find jobs with our cli using claude code and codex. We scrape the internet every hour to find you jobs linkedin and indeed cant find. also it's free
npx magnificentjobs setup # Claude Code · Codex
npx magnificentjobs # terminal
Install
/plugin marketplace add AdanRott/magnificent-jobs-plugin
/plugin install magnificent-jobs@magnificent-jobs
Or, in Claude.ai / Claude Desktop, add it as a custom connector: Settings → Connectors → Add
custom connector → https://magnificentjobs.com/mcp (no authentication).
The server is remote, read-only and needs no key. Listed in the official MCP Registry as
com.magnificentjobs/jobs . Add it to any client that supports Streamable HTTP:
{
"mcpServers" : {
"magnificent-jobs" : { "type" : " http " , "url" : " https://magnificentjobs.com/mcp " }
}
}
Claude Code: claude mcp add --transport http magnificent-jobs https://magnificentjobs.com/mcp
stdio-only clients / Docker: npx magnificentjobs mcp speaks MCP over stdio and proxies to the hosted server ( docker build -t magnificentjobs . && docker run -i magnificentjobs )
Cursor / Windsurf / VS Code: paste the JSON above into the MCP settings.
ChatGPT: install Magnificent Jobs from the Plugins directory, or in Developer mode add https://magnificentjobs.com/mcp .
Claude.ai / Desktop: Settings → Connectors → Add custom connector → https://magnificentjobs.com/mcp .
Tools: search_jobs · get_job · find_cities · list_states . Docs: https://magnificentjobs.com/developers .
Component
What it does
MCP server https://magnificentjobs.com/mcp (Streamable HTTP, no auth)
search_jobs (semantic search; city radius 10/30/50 mi, state, remote/on-site, experience band, employment type), get_job (full description, salary, skills, apply link), find_cities , list_states . All tools are read-only.
Skill find-jobs
How to phrase the query (full title + one sentence + skills), when to use which tool, US-only coverage, never invent jobs, never claim to apply.
Try
"Find remote senior React engineer jobs and show me the three best with salaries."
"What warehouse or forklift jobs are within 10 miles of Fresno? Give me the apply link for the closest one."
"Show entry-level marketing jobs in Colorado and summarize the top two."
"Which states have the most nursing jobs right now?"
Searches are sent to magnificentjobs.com to return results and are not stored beyond standard server
logs. No account, no personal data. Full policy: https://magnificentjobs.com/privacy ·
Terms: https://magnificentjobs.com/terms · Docs: https://magnificentjobs.com/developers ·
Support: info --- at --- magnificentjobs.com
MIT — see LICENSE . The plugin files are MIT; the job data is served by magnificentjobs.com under its terms.
Magnificent Jobs for Claude — find live US jobs by describing the role you want (MCP server + skill)
magnificentjobs.com/developers Resources
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
