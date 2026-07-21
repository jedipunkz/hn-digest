---
source: "https://github.com/TryRekon/Rekon"
hn_url: "https://news.ycombinator.com/item?id=48998468"
title: "Open Source AI Harness Profiler – discover where tf your tokens are going"
article_title: "GitHub - TryRekon/Rekon · GitHub"
author: "TreDub"
captured_at: "2026-07-21T21:55:56Z"
capture_tool: "hn-digest"
hn_id: 48998468
score: 1
comments: 1
posted_at: "2026-07-21T21:17:28Z"
tags:
  - hacker-news
  - translated
---

# Open Source AI Harness Profiler – discover where tf your tokens are going

- HN: [48998468](https://news.ycombinator.com/item?id=48998468)
- Source: [github.com](https://github.com/TryRekon/Rekon)
- Score: 1
- Comments: 1
- Posted: 2026-07-21T21:17:28Z

## Translation

タイトル: オープンソース AI Harness Profiler – トークンの行き先を発見する
記事タイトル: GitHub - TryRekon/Rekon · GitHub
説明: GitHub でアカウントを作成して、TryRekon/Rekon の開発に貢献します。

記事本文:
GitHub - TryRekon/Rekon · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
トライレコン
/
レコン
公共
通知
通知設定を変更するにはサインインする必要があります
追加

ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .claude .claude 移行 移行スクリプト スクリプト共有 共有 src src Web Web .dev.vars.example .dev.vars.example .env.example .env.example .gitignore .gitignore .stylelintrc.cjs .stylelintrc.cjs AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md DCO DCO ライセンス ライセンスに関する通知 README.md README.md drizzle.config.ts drizzle.config.ts eslint.config.js eslint.config.js Index.html Index.html package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vite.config.ts vite.config.ts wrangler.jsonc wrangler.jsonc すべてのファイルを表示 リポジトリ ファイルのナビゲーション
トークンを記録する Anthropic API および OpenAI API の透過プロキシ
使用状況を確認し、ダッシュボードに表示します。クライアントにそれを指示し、そのまま作業を続けてください
通常、リクエストごとのトークン、再構築されたセッション、ツールレベルを参照
帰属と推定コスト。
Cloudflare Workers、D1、Hono、React に基づいて構築されています。
ホストされたインスタンスは tryrekon.com で実行されます。
追加遅延ゼロ - 応答は直接ストリーミングされます。録音が行われる
クリティカルパスから外れています。
キーは決して保存されません - 認証ヘッダーは転送され、決して保持されないため、使用法は
常に呼び出し元自身のキーまたはサブスクリプションに費やされます。
セッション ツリー — チャット API はステートレスであるため、プロキシはそれぞれのセッション ツリーを再構築します。
クライアントシグナル、フォークなどからの会話。
ツールレベルのアトリビューション — どのツールが実際にトークンを書き込むかを確認します。
ダッシュボードでシステムを作成すると、独自のプロキシ URL が取得されます。ポイントしてください
クライアントでは 1 つの環境変数を使用します。
import ANTHROPIC_BASE_URL=https:// < ホスト > /s/ < システム ID >
# OpenAI クライアントは /openai/v1 を追加します:
import OPENAI_BASE_URL=https:// < ホスト > /s/ < システム ID > /openai/v1
リクエストはそのまま転送されます (認証ヘッダー

含まれます）、使用状況は次のように記録されます。
帰り道。システム ID は取り込みキーとしても機能します。不明な ID は拒否されます。
したがって、デプロイメントは決してオープンリレーではありません。 URL を秘密として扱います。
git クローン https://github.com/TryRekon/Rekon.git
cd レコン
npmインストール
cp .dev.vars.example .dev.vars # SESSION_SECRET + 1 つの OAuth プロバイダー (Google または GitHub) を設定します
npm run typegen
npm 実行 db:移行:ローカル
npm run dev # http://localhost:5173
サインインし、オンボーディング画面からシステムの URL をコピーし、クライアントに次の URL を指定します。
それ。記録は最初のリクエストから始まります。
独自のデータベースをデプロイするには: D1 データベースを作成します ( wrangler d1 create token-profiler-db )。
その ID を wrangler.jsonc に database_id として貼り付けて、
カスタムドメインパターンをあなたが制御するゾーンにルーティングします（または無料で削除します）
works.dev URL)。 Wrangler Secret put ( SESSION_SECRET 、
OAuth プロバイダーのペア、および分析が必要な場合は POSTHOG_KEY) を実行します。
npm run db:merge:remote 、次に npm rundeploy 。デプロイは行われないことに注意してください。
移行を実行します。
コストは、shared/pricing.ts の定価からの見積もりです。未知のモデルショー
費用はかかりません。
アーキテクチャ — セッションの再構築、トークンの帰属、エンドポイント リスト、
ファイルごとのレイアウトは AGENTS.md にあります。
アパッチ 2.0 。このリポジトリは、今後も完全にオープン ソースであり続けます。の
商用モデルはホストされた展開であり、このコード上のペイウォールではありません。
貢献は DCO に基づいて受け入れられます - でコミットをサインオフします。
git commit -s 。
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

Contribute to TryRekon/Rekon development by creating an account on GitHub.

GitHub - TryRekon/Rekon · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
TryRekon
/
Rekon
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .claude .claude migrations migrations scripts scripts shared shared src src web web .dev.vars.example .dev.vars.example .env.example .env.example .gitignore .gitignore .stylelintrc.cjs .stylelintrc.cjs AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md DCO DCO LICENSE LICENSE NOTICE NOTICE README.md README.md drizzle.config.ts drizzle.config.ts eslint.config.js eslint.config.js index.html index.html package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vite.config.ts vite.config.ts wrangler.jsonc wrangler.jsonc View all files Repository files navigation
A transparent proxy for the Anthropic and OpenAI APIs that records token
usage and shows it in a dashboard. Point your client at it, keep working as
normal, and see per-request tokens, reconstructed sessions, tool-level
attribution, and estimated cost.
Built on Cloudflare Workers, D1, Hono, and React.
A hosted instance runs at tryrekon.com .
Zero added latency — responses stream straight through; recording happens
off the critical path.
Keys never stored — auth headers are forwarded, never held, so usage is
always spent on the caller's own key or subscription.
Session trees — the chat APIs are stateless, so the proxy rebuilds each
conversation from client signals, forks and all.
Tool-level attribution — see which tools actually burn the tokens.
Create a system in the dashboard and it gets its own proxy URL. Point your
client at it with one environment variable:
export ANTHROPIC_BASE_URL=https:// < host > /s/ < system-id >
# OpenAI clients add /openai/v1:
export OPENAI_BASE_URL=https:// < host > /s/ < system-id > /openai/v1
Requests are forwarded untouched (auth headers included) and usage is recorded on
the way back. The system id doubles as the ingest key — unknown ids are rejected,
so a deployment is never an open relay. Treat the URL as a secret.
git clone https://github.com/TryRekon/Rekon.git
cd Rekon
npm install
cp .dev.vars.example .dev.vars # set SESSION_SECRET + one OAuth provider (Google or GitHub)
npm run typegen
npm run db:migrate:local
npm run dev # http://localhost:5173
Sign in, copy your system's URL from the onboarding screen, and point a client at
it. Recording starts with the first request.
To deploy your own: create a D1 database ( wrangler d1 create token-profiler-db )
and paste its id into wrangler.jsonc as database_id , then change the
routes custom-domain pattern to a zone you control (or remove it for a free
workers.dev URL). Set your secrets with wrangler secret put ( SESSION_SECRET ,
an OAuth provider pair, and POSTHOG_KEY if you want analytics), run
npm run db:migrate:remote , then npm run deploy . Note that deploy does not
run migrations.
Costs are estimates from list prices in shared/pricing.ts ; unknown models show
no cost.
The architecture — session reconstruction, token attribution, the endpoint list,
and a file-by-file layout — is in AGENTS.md .
Apache 2.0 . This repository is, and will remain, fully open source; the
commercial model is a hosted deployment, not a paywall over this code.
Contributions are accepted under the DCO — sign off your commits with
git commit -s .
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
