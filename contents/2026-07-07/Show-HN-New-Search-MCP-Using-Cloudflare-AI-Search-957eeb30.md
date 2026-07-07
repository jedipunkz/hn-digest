---
source: "https://github.com/skyphusion-labs/search-mcp"
hn_url: "https://news.ycombinator.com/item?id=48818002"
title: "Show HN: New Search MCP Using Cloudflare AI Search"
article_title: "GitHub - skyphusion-labs/search-mcp: Open-source Cloudflare AI Search toolkit: bearer-gated MCP server, streaming /ask query Worker, embeddable docs widget, and git-to-R2 corpus sync with extension remapping (AGPL). · GitHub"
author: "skyphusion"
captured_at: "2026-07-07T14:07:39Z"
capture_tool: "hn-digest"
hn_id: 48818002
score: 1
comments: 0
posted_at: "2026-07-07T14:00:27Z"
tags:
  - hacker-news
  - translated
---

# Show HN: New Search MCP Using Cloudflare AI Search

- HN: [48818002](https://news.ycombinator.com/item?id=48818002)
- Source: [github.com](https://github.com/skyphusion-labs/search-mcp)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T14:00:27Z

## Translation

タイトル: HN の表示: Cloudflare AI 検索を使用した新しい検索 MCP
Article title: GitHub - skyphusion-labs/search-mcp: Open-source Cloudflare AI Search toolkit: bearer-gated MCP server, streaming /ask query Worker, embeddable docs widget, and git-to-R2 corpus sync with extension remapping (AGPL). · GitHub
Description: Open-source Cloudflare AI Search toolkit: bearer-gated MCP server, streaming /ask query Worker, embeddable docs widget, and git-to-R2 corpus sync with extension remapping (AGPL). - skyphusion-labs/search-mcp
HN text: Hello, So I got sick of my agents not having ready access to information they needed about projects that I'm working on ready and at hand so I built a new AGPL-3.0 (free forever) MCP that utilizes Cloudflare AI Search to index your repos and anything else you want to feed into it. It also includes a nice little web widget that allows end users to search your repos for documentation, etc., and allows end users to ask an AI Chat Bot questions regarding your projects.ウィジェットを表示したい場合は、私が取り組んでいる別のプロジェクトでウィジェットを有効にしました: https://vivijure.com 。 I hope people will find this new piece of software useful, it's been very useful for me so far and I'm sure there's at least one other person out there who could use it :).ありがとう、-スカイプフュージョン

記事本文:
GitHub - skyphusion-labs/search-mcp: オープンソースの Cloudflare AI 検索ツールキット: ベアラーゲート MCP サーバー、ストリーミング /ask クエリ ワーカー、埋め込み可能なドキュメント ウィジェット、および拡張機能再マッピング (AGPL) を使用した git-to-R2 コーパス同期。 · GitHub
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
ありました

読み込み中のエラーです。このページをリロードしてください。
スカイプフュージョンラボ
/
検索-mcp
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット .github .github docs docs public public scripts scripts src src .gitattributes .gitattributes .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンスREADME.md README.md SECURITY.md SECURITY.md Index.test.ts Index.test.ts package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts wrangler.mcp.toml.example wrangler.mcp.toml.example wrangler.toml.example wrangler.toml.example すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Cloudflare AI Search用のオープンソースツールキット:
MCP ワーカー ( src/mcp.ts ) -- エージェント用の検索ツールを備えたベアラーゲート型 Streamable-HTTP MCP。
クエリ ワーカー ( src/index.ts ) -- CORS + Turnstile + レート制限された POST /ask で、ブラウザ ウィジェットの回答をストリーミングします。
コーパス同期 ( scripts/sync.mjs 、 scripts/sync-runner.mjs ) -- git で追跡されたソースを R2 に送信します。拡張子は再マッピングされているため、TypeScript、Dockerfile、その他のテキスト AI 検索はインデックス作成をスキップします。
git リポジトリ -> sync.mjs -> R2 バケット -> AI Search インスタンス -> /ask + /mcp
クイックスタート
npmインストール
cp wrangler.toml.example wrangler.toml
cp wrangler.mcp.toml.example wrangler.mcp.toml
cp スクリプト/targets.json.example スクリプト/targets.json
# アカウント、インスタンス、バケット、リポジトリの 3 つのファイルを編集します
npm タイプチェックを実行する
npmテスト
R2 + AI Search をプロビジョニングし、コーパスを同期し、両方のワーカーをデプロイします。ステップバイステップ: docs/DEPLOY.md 。
労働者
エントリー
エンドポイント
認証
クエリ
ラングラー.toml
POST /ask 、 GET /health
回転式改札口 (オプション)

) + CORS 許可リスト
MCP
wrangler.mcp.toml
POST /mcp 、 GET /health
権限: Bearer (フェイルクローズ)
必要に応じて、ブラウザー トラフィックとエージェント トラフィックが異なる AI Search インスタンスをバインドできるように、個別にデプロイします。
npm rundeploy #クエリワーカー
npm rundeploy:mcp # MCP ワーカー
wrangler シークレット put MCP_TOKEN -c wrangler.mcp.toml
Wrangler シークレットに TURNSTILE_SECRET # オプションを追加します。設定を解除すると検証をスキップします
MCPクライアントの配線
{
"mcpサーバー": {
"検索-mcp" : {
"type" : " http " ,
"url" : " https://YOUR_MCP_HOST/mcp " ,
"headers" : { "Authorization" : " Bearer YOUR_TOKEN " }
}
}
}
MCP_TOKEN は、ログ内の消費者ごとの属性として、単一のトークンまたはカンマ区切りの name=token ペアを受け入れます。
エクスポート R2_ACCESS_KEY_ID=... R2_SECRET_ACCESS_KEY=... CLOUDFLARE_ACCOUNT_ID=...
import CORPUS_GIT_ORG=your-org GITHUB_TOKEN=... # sync-runner クローン認証用
npm run sync:dry # デフォルトの `corpus` ターゲットのアップロードを計画します
npm run sync # アップロード + プルーン
npm run sync:run # 分離されたクローン ルート、すべてのターゲットを同期、オプションでインデックスを再作成
同期により、非ネイティブ拡張子 ( .ts 、 .tsx 、拡張子なしの Dockerfile 、 .service など) が .txt キーに再マッピングされるため、AI Search はこれらの拡張子にインデックスを付けます。 scripts/sync-ingest.mjs を参照してください。
ターゲットに public という名前を付けると、オプションの境界チェックによって制限付きリポジトリとの重複が拒否され、アップロード前に GitHub リポジトリの可視性を確認できます。
public/ask-widget.js と public/ask-widget.css をドキュメント サイトにコピーします。
< div id =" docs-ask " > </ div >
< script defer src =" /ask-widget.js "
データエンドポイント =" https://search.example.com/ask "
データターゲット = #docs-ask "
data-label =" ドキュメントに問い合わせてください "
data-sitekey =" YOUR_TURNSTILE_SITEKEY " > </ script >
ライセンス
オープンソースのCloudflare AI検索ツールキット: ベアラーゲートMCPサーバー、ストリーミング/askクエリワーカー、埋め込み可能なドキュメントウィジェット、拡張機能remappinによるgit-to-R2コーパス同期

g (AGPL)。
開発者.cloudflare.com/ai-search/
トピックス
AGPL-3.0ライセンス
行動規範
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
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open-source Cloudflare AI Search toolkit: bearer-gated MCP server, streaming /ask query Worker, embeddable docs widget, and git-to-R2 corpus sync with extension remapping (AGPL). - skyphusion-labs/search-mcp

Hello, So I got sick of my agents not having ready access to information they needed about projects that I'm working on ready and at hand so I built a new AGPL-3.0 (free forever) MCP that utilizes Cloudflare AI Search to index your repos and anything else you want to feed into it. It also includes a nice little web widget that allows end users to search your repos for documentation, etc., and allows end users to ask an AI Chat Bot questions regarding your projects. If you want to see the widget, I enabled it on another project I'm working on here: https://vivijure.com . I hope people will find this new piece of software useful, it's been very useful for me so far and I'm sure there's at least one other person out there who could use it :). Thanks, -Skyphusion

GitHub - skyphusion-labs/search-mcp: Open-source Cloudflare AI Search toolkit: bearer-gated MCP server, streaming /ask query Worker, embeddable docs widget, and git-to-R2 corpus sync with extension remapping (AGPL). · GitHub
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
skyphusion-labs
/
search-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .github .github docs docs public public scripts scripts src src .gitattributes .gitattributes .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md index.test.ts index.test.ts package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts wrangler.mcp.toml.example wrangler.mcp.toml.example wrangler.toml.example wrangler.toml.example View all files Repository files navigation
Open-source toolkit for Cloudflare AI Search :
MCP Worker ( src/mcp.ts ) -- bearer-gated Streamable-HTTP MCP with a search tool for agents.
Query Worker ( src/index.ts ) -- CORS + Turnstile + rate-limited POST /ask that streams answers for a browser widget.
Corpus sync ( scripts/sync.mjs , scripts/sync-runner.mjs ) -- git-tracked sources to R2, with extension remapping so TypeScript, Dockerfiles, and other text AI Search would otherwise skip get indexed.
git repos -> sync.mjs -> R2 bucket -> AI Search instance -> /ask + /mcp
Quick start
npm install
cp wrangler.toml.example wrangler.toml
cp wrangler.mcp.toml.example wrangler.mcp.toml
cp scripts/targets.json.example scripts/targets.json
# edit the three files for your account, instance, bucket, and repos
npm run typecheck
npm test
Provision R2 + AI Search, sync your corpus, deploy both Workers. Step-by-step: docs/DEPLOY.md .
Worker
Entry
Endpoint
Auth
Query
wrangler.toml
POST /ask , GET /health
Turnstile (optional) + CORS allowlist
MCP
wrangler.mcp.toml
POST /mcp , GET /health
Authorization: Bearer (fail closed)
Deploy separately so browser traffic and agent traffic can bind different AI Search instances if you want.
npm run deploy # query Worker
npm run deploy:mcp # MCP Worker
wrangler secret put MCP_TOKEN -c wrangler.mcp.toml
wrangler secret put TURNSTILE_SECRET # optional; skips verification when unset
MCP client wiring
{
"mcpServers" : {
"search-mcp" : {
"type" : " http " ,
"url" : " https://YOUR_MCP_HOST/mcp " ,
"headers" : { "Authorization" : " Bearer YOUR_TOKEN " }
}
}
}
MCP_TOKEN accepts a single token or comma-separated name=token pairs for per-consumer attribution in logs.
export R2_ACCESS_KEY_ID=... R2_SECRET_ACCESS_KEY=... CLOUDFLARE_ACCOUNT_ID=...
export CORPUS_GIT_ORG=your-org GITHUB_TOKEN=... # for sync-runner clone auth
npm run sync:dry # plan upload for the default `corpus` target
npm run sync # upload + prune
npm run sync:run # isolated clone root, sync all targets, optional reindex
The sync remaps non-native extensions ( .ts , .tsx , extensionless Dockerfile , .service , etc.) to .txt keys so AI Search indexes them. See scripts/sync-ingest.mjs .
If you name a target public , optional boundary checks refuse overlap with restrictedRepos and can verify GitHub repo visibility before upload.
Copy public/ask-widget.js and public/ask-widget.css to your docs site:
< div id =" docs-ask " > </ div >
< script defer src =" /ask-widget.js "
data-endpoint =" https://search.example.com/ask "
data-target =" #docs-ask "
data-label =" Ask the docs "
data-sitekey =" YOUR_TURNSTILE_SITEKEY " > </ script >
License
Open-source Cloudflare AI Search toolkit: bearer-gated MCP server, streaming /ask query Worker, embeddable docs widget, and git-to-R2 corpus sync with extension remapping (AGPL).
developers.cloudflare.com/ai-search/
Topics
AGPL-3.0 license
Code of conduct
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
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
