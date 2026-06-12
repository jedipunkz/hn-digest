---
source: "https://github.com/nochinxx/SwiPR"
hn_url: "https://news.ycombinator.com/item?id=48498866"
title: "Built SwiPR – swipe-to-review GitHub PRs with AI context"
article_title: "GitHub - nochinxx/SwiPR · GitHub"
author: "nochinxx"
captured_at: "2026-06-12T04:35:29Z"
capture_tool: "hn-digest"
hn_id: 48498866
score: 1
comments: 0
posted_at: "2026-06-12T01:49:26Z"
tags:
  - hacker-news
  - translated
---

# Built SwiPR – swipe-to-review GitHub PRs with AI context

- HN: [48498866](https://news.ycombinator.com/item?id=48498866)
- Source: [github.com](https://github.com/nochinxx/SwiPR)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T01:49:26Z

## Translation

タイトル: 構築された SwiPR – AI コンテキストを使用してスワイプして GitHub PR をレビュー
記事タイトル: GitHub - nochinxx/SwiPR · GitHub
説明: GitHub でアカウントを作成して、nochinxx/SwiPR の開発に貢献します。

記事本文:
GitHub - nochinxx/SwiPR · GitHub
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
ノチンクス
/
SwiPR
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
50℃

省略 50 コミット .claude .claude app app コンポーネント コンポーネント db db フック フック lib lib public public scripts スクリプト スタイル スタイル .env.example .env.example .gitignore .gitignore ライセンス ライセンス MIGRATION.md MIGRATION.md README.md README.md コンポーネント.json コンポーネント.json drizzle.config.ts drizzle.config.ts next-env.d.ts next-env.d.ts next.config.mjs next.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml postcss.config.mjs postcss.config.mjs tsconfig.json tsconfig.json すべてのファイルを表示リポジトリ ファイルのナビゲーション
AI コンテキストを使用して GitHub PR をスワイプして確認します。クロード MCP プラグインとして機能します。
ライブデモを試す→ ・Claudeデスクトップに追加↓ ・smithery.aiにリストされる
X の Zeno Rocha (Resend の CEO):
「PR を開設するコストはゼロになりました。現在、大量のドラフト PR が有限の (そして非常に貴重な) リソースである注意を待っています。ボトルネックはもはや作成ではなく、レビューであることがわかりました。」
— 2026 年 2 月
「以前は、メイン リポジトリには 1 日に平均約 20 ～ 40 のオープン PR がありましたが、現在は平均約 130 ～ 200 のオープン PR があります。」
— 2026 年 5 月
PR レビューは文脈の問題です。メンテナがトリアージを行っている場合でも、AI エージェントがマージするかどうかを決定している場合でも、問題は同じです。「これは出荷しても安全ですか?」ということです。それに適切に答えるには、単に変更された行を読み取るだけでなく、リスク レベル、貢献者の履歴、同様の過去の変更、および差分をカバーするテストを知る必要があります。
SwiPR は、人間にとってはスワイプ UI、エージェントにとっては MCP サーバーとしてそのコンテキストを表示します。
パブリック GitHub リポジトリを貼り付けます — SwiPR はオープン PR を取得し、埋め込みとともに保存します
右にスワイプして承認、左にスワイプして変更をリクエスト、下にスワイプしてスキップ — または J / F / Space を使用します
右側のパネルには、リスク スコア、AI の概要、同様の過去の PR、貢献者の履歴が表示されます。
「なぜこれが危険なのですか?」を押します。

、「発信者を表示」、または「これをカバーするテストは何ですか?」より深いコンテキストをオンデマンドで確認するには
チャットで何でも質問してください - AI は完全な差分とコードベースのコンテキストにアクセスできます
AI コンテキストは PR ごとにキャッシュされます。すべてのカード ビューで API 呼び出しが繰り返されることはありません。
SwiPR は、MCP サーバーと同じ PR コンテキストを公開します。これを Claude Desktop または Cursor に追加し、チャットから直接 PR を確認します。ブラウザは必要ありません。
~/Library/Application Support/Claude/claude_desktop_config.json に追加します。
{
"mcpサーバー": {
"スワイパー" : {
"コマンド" : " npx " ,
"args" : [ " -y " 、 " supergateway " 、 " --streamableHttp " 、 " https://v0-swipr-build.vercel.app/api/mcp " ]
}
}
}
次にクロードに尋ねます。
「再送信/再送信ノード PR #1247 を調べてください。リスクは何ですか? 過去に同様の変更はありますか?」
「vercel/next.js 内のどのオープン PR がルーターに影響しますか?」
「変更されたファイルをカバーするテストを PR #892 で見つけてください。」
12 の MCP ツールはすべて純粋なデータベース読み取りであり、ホストされたサーバーの使用時に消費される AI クレジットはゼロです。
SwiPR は smithery.ai および mcp.so にリストされています。
ツール
返されるもの
ルックアップ_pr
所有者/リポジトリ番号を解決 → 内部 ID
分析_pr
完全な PR データ: ファイル、パッチ、リスク スコア、貢献者の統計、キャッシュされた AI の概要
リスクスコア
0 ～ 100 のヒューリスティック スコアと理由
類似変更の検索
意味的に類似した同じリポジトリ内の過去の PR
get_contributor_history
PR 数、マージ率、最初の投稿日
検査ファイル
HEAD の生ファイルの内容
find_callers
関数名の使用法をパッチで検索する
find_関連テスト
変更されたコードをカバーする可能性のあるテスト ファイル
git_blame_summary
ファイルパスへの最近の投稿者
比較する
任意の git ref のファイルの内容
記録_決定
キャプチャ承認/変更/スキップ
要約セッション
セッション終了時の統計
自分の AI キーを持参する
チャット機能は、これらの無料枠プロバイダーのいずれでも動作します — ⌘ ボタンをクリックしてください

ヘッダーに n を入力してキーを貼り付けます。これはブラウザの localStorage にのみ保存され、サーバーには送信されません。
リスク ヒューリスティックは完全に lib/scoring.ts 内に存在します。プレーンな TypeScript、ML、外部呼び出しはありません。チームの標準に合わせて編集してください。
ルールを追加するには、 computeRiskScore に if ブロックを追加します。
const touchAuth = ファイル 。 some ( ( f ) => f . filename . include ( "lib/auth" ) ) ;
if ( touchAuth ) {
スコア += 25 ;
理由。 Push ( "認証モジュールにタッチします — セキ​​ュリティレビューが必要です" ) ;
}
スコアの上限は 100 です。色のしきい値 (緑/黄/赤) は 40 と 70 です。 app/swipe/_components/view-helpers.ts で調整します。
3 つのサービスが必要です。すべてに無料利用枠があります。
git clone https://github.com/nochinxx/SwiPR.git
cd SwiPR
pnpmインストール
2. アカウントを作成する
サービス
目的
無料利用枠
ネオン
Postgres + pgvector
512MB
Vercel AI ゲートウェイ
クロード + 埋め込み
5ドルのクレジット
GitHub PAT (オプション)
より高い API レート制限
無料
3. 環境の設定
cp .env.example .env.local
DATABASE_URL = postgresql://... # Neon プールされた接続文字列
AI_GATEWAY_API_KEY = ... # Vercel AI ゲートウェイ キー
GITHUB_TOKEN = ... # オプション — GitHub のレート制限を 5000/hr に引き上げます
4. データベースの初期化
-- Neon SQL エディターで実行
ベクトルが存在しない場合は拡張子を作成します。
pnpm db:プッシュ
-- 高速な類似性検索のための HNSW インデックスの追加
存在しない場合はインデックスを作成します。 prs_embedding_idx ON prs USING hnsw (vector_cosine_ops を埋め込む);
存在しない場合はインデックスを作成します。 hnsw を使用して pr_files_embedding_idx ON pr_files (vector_cosine_ops を埋め込む);
5. 実行と展開
pnpm dev # ローカル開発
pnpm build # デプロイ前に検証する
pnpm precache # デモ用のリポジトリをプリロードする
GitHub にプッシュし、Vercel プロジェクトに接続し、Vercel 設定で同じ環境変数を設定します。 Claude デスクトップ構成を独自の展開 URL に指定します。
Neon の無料利用枠は継続されます

～512MB。取り込まれた各 PR は最大 100 KB を使用します。これは、アップグレードが必要になるまでに約 5,000 PR または 40 ～ 50 の中規模リポジトリに相当します。取り込みは、リポジトリあたり 100 個のオープン PR に制限されます。
Next.js 16 — アプリルーター、React 19
Tailwind v4 + shadcn/ui — スタイリング
Framer Motion — カードアニメーション
Neon Postgres + pgvector — PR ストレージと類似性検索
Drizzle ORM — スキーマとクエリ
Vercel AI ゲートウェイ — Claude Sonnet 4.6 (分析)、text-embedding-3-small (ベクター)
@ai-sdk/anthropic · @ai-sdk/google · @ai-sdk/groq — BYOK チャット サポート
mcp-handler — /api/mcp にある MCP サーバー
スタックは Next.js App Router + Drizzle + Neon + Vercel AI SDK であり、他の抽象化はありません。
lib/scoring.ts でのリスク スコアリング — ヒューリスティックな改善は特に歓迎されます。
GitHub OAuth や実際の PR 投稿を追加しないでください。読み取り専用アクセスは意図的です。
PR を開く前に pnpm build を実行します。
v0-swipr-build.vercel.app
リソース
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.1.0 — 初期リリース
最新の
2026 年 6 月 12 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to nochinxx/SwiPR development by creating an account on GitHub.

GitHub - nochinxx/SwiPR · GitHub
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
nochinxx
/
SwiPR
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
50 Commits 50 Commits .claude .claude app app components components db db hooks hooks lib lib public public scripts scripts styles styles .env.example .env.example .gitignore .gitignore LICENSE LICENSE MIGRATION.md MIGRATION.md README.md README.md components.json components.json drizzle.config.ts drizzle.config.ts next-env.d.ts next-env.d.ts next.config.mjs next.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml postcss.config.mjs postcss.config.mjs tsconfig.json tsconfig.json View all files Repository files navigation
Swipe-to-review GitHub PRs with AI context. Works as a Claude MCP plugin.
Try the live demo → · Add to Claude Desktop ↓ · Listed on smithery.ai
Zeno Rocha (CEO of Resend ) on X:
"the cost of opening a PR has dropped to zero. now, we have tons of draft PRs waiting for a finite (and ultra precious) resource: attention. turns out the bottleneck is no longer creation. it's reviewing."
— Feb 2026
"before our main repo had an average of ~20–40 open PRs on any given day. now, we average ~130–200 open PRs."
— May 2026
PR review is a context problem. Whether a maintainer is triaging or an AI agent is deciding whether to merge, the question is the same: is this safe to ship? Answering it well requires knowing the risk level, contributor history, similar past changes, and which tests cover the diff — not just reading the lines changed.
SwiPR surfaces that context as a swipe UI for humans and an MCP server for agents.
Paste any public GitHub repo — SwiPR fetches open PRs and stores them with embeddings
Swipe right to approve, left to request changes, down to skip — or use J / F / Space
The right panel surfaces: risk score, AI summary, similar past PRs, contributor history
Hit "Why is this risky?" , "Show me callers" , or "What tests cover this?" for deeper context on demand
Ask anything in the chat — the AI has access to the full diff and codebase context
AI context is cached per PR — no repeated API calls on every card view.
SwiPR exposes the same PR context as an MCP server. Add it to Claude Desktop or Cursor and review PRs directly from chat — no browser required.
Add to ~/Library/Application Support/Claude/claude_desktop_config.json :
{
"mcpServers" : {
"swipr" : {
"command" : " npx " ,
"args" : [ " -y " , " supergateway " , " --streamableHttp " , " https://v0-swipr-build.vercel.app/api/mcp " ]
}
}
}
Then ask Claude:
"Look up resend/resend-node PR #1247 — what's the risk and are there similar past changes?"
"Which open PRs in vercel/next.js touch the router?"
"Find tests that cover the changed files in PR #892."
All 12 MCP tools are pure database reads — zero AI credits consumed when using the hosted server.
SwiPR is listed on smithery.ai and mcp.so .
Tool
What it returns
lookup_pr
Resolve owner/repo#number → internal ID
analyze_pr
Full PR data: files, patches, risk score, contributor stats, cached AI summary
risk_score
0–100 heuristic score with reasons
find_similar_changes
Past PRs in the same repo with semantic similarity
get_contributor_history
PR count, merge rate, first contribution date
inspect_file
Raw file content at HEAD
find_callers
Search patches for usages of a function name
find_related_tests
Test files that likely cover the changed code
git_blame_summary
Recent contributors to a file path
compare_with
File content at an arbitrary git ref
record_decision
Capture approve / changes / skip
summarize_session
End-of-session stats
Bring your own AI key
The chat feature works with any of these free-tier providers — click the ⌘ button in the header and paste your key. It's stored in your browser's localStorage only, never sent to the server.
The risk heuristic lives entirely in lib/scoring.ts — plain TypeScript, no ML, no external calls. Edit it to match your team's standards.
To add a rule, add an if block in computeRiskScore :
const touchesAuth = files . some ( ( f ) => f . filename . includes ( "lib/auth" ) ) ;
if ( touchesAuth ) {
score += 25 ;
reasons . push ( "Touches auth module — requires security review" ) ;
}
Score is capped at 100. Color thresholds (green/yellow/red) are at 40 and 70 — adjust in app/swipe/_components/view-helpers.ts .
Requires three services — all have free tiers.
git clone https://github.com/nochinxx/SwiPR.git
cd SwiPR
pnpm install
2. Create accounts
Service
Purpose
Free tier
Neon
Postgres + pgvector
512 MB
Vercel AI Gateway
Claude + embeddings
$5 credits
GitHub PAT (optional)
Higher API rate limits
Free
3. Configure environment
cp .env.example .env.local
DATABASE_URL = postgresql://... # Neon pooled connection string
AI_GATEWAY_API_KEY = ... # Vercel AI Gateway key
GITHUB_TOKEN = ... # Optional — raises GitHub rate limit to 5000/hr
4. Initialize the database
-- Run in Neon SQL editor
CREATE EXTENSION IF NOT EXISTS vector;
pnpm db:push
-- Add HNSW indexes for fast similarity search
CREATE INDEX IF NOT EXISTS prs_embedding_idx ON prs USING hnsw (embedding vector_cosine_ops);
CREATE INDEX IF NOT EXISTS pr_files_embedding_idx ON pr_files USING hnsw (embedding vector_cosine_ops);
5. Run and deploy
pnpm dev # local dev
pnpm build # verify before deploying
pnpm precache # pre-load repos for a demo
Push to GitHub, connect to a Vercel project, set the same env vars in Vercel settings. Point your Claude Desktop config at your own deployment URL.
Neon's free tier holds ~512 MB. Each ingested PR uses ~100 KB. That's around 5,000 PRs or 40–50 mid-sized repos before needing an upgrade. Ingest is capped at 100 open PRs per repo.
Next.js 16 — App Router, React 19
Tailwind v4 + shadcn/ui — styling
Framer Motion — card animations
Neon Postgres + pgvector — PR storage and similarity search
Drizzle ORM — schema and queries
Vercel AI Gateway — Claude Sonnet 4.6 (analysis), text-embedding-3-small (vectors)
@ai-sdk/anthropic · @ai-sdk/google · @ai-sdk/groq — BYOK chat support
mcp-handler — MCP server at /api/mcp
Stack is Next.js App Router + Drizzle + Neon + Vercel AI SDK — no other abstractions.
Risk scoring in lib/scoring.ts — heuristic improvements especially welcome.
Don't add GitHub OAuth or actual PR posting. Read-only access is intentional.
Run pnpm build before opening a PR.
v0-swipr-build.vercel.app
Resources
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.1.0 — initial release
Latest
Jun 12, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
