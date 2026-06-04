---
source: "https://github.com/HYPD-AI/openai-ads-mcp"
hn_url: "https://news.ycombinator.com/item?id=48396767"
title: "Show HN: MCP for the ChatGPT Ads API – Query ChatGPT Ads from Claude and Codex"
article_title: "GitHub - HYPD-AI/openai-ads-mcp: MCP server for OpenAI Ads / ChatGPT Ads. Manage campaigns, ad groups, ads, creatives, and insights through the OpenAI Ads Advertiser API. · GitHub"
author: "cionut"
captured_at: "2026-06-04T13:15:49Z"
capture_tool: "hn-digest"
hn_id: 48396767
score: 1
comments: 1
posted_at: "2026-06-04T10:45:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: MCP for the ChatGPT Ads API – Query ChatGPT Ads from Claude and Codex

- HN: [48396767](https://news.ycombinator.com/item?id=48396767)
- Source: [github.com](https://github.com/HYPD-AI/openai-ads-mcp)
- Score: 1
- Comments: 1
- Posted: 2026-06-04T10:45:15Z

## Translation

タイトル: HN を表示: ChatGPT 広告 API の MCP – Claude と Codex から ChatGPT 広告をクエリする
記事タイトル: GitHub - HYPD-AI/openai-ads-mcp: OpenAI 広告 / ChatGPT 広告用の MCP サーバー。 OpenAI Ads Advertiser API を通じて、キャンペーン、広告グループ、広告、クリエイティブ、インサイトを管理します。 · GitHub
説明: OpenAI 広告/ChatGPT 広告用の MCP サーバー。 OpenAI Ads Advertiser API を通じて、キャンペーン、広告グループ、広告、クリエイティブ、インサイトを管理します。 - HYPD-AI/openai-ads-mcp
HN テキスト: 私は ChatGPT Ads 広告主 API 用の小さな MCP サーバーを構築しました。これは OpenAI が数週間前に一般公開しました。現時点では、アカウント、キャンペーン、広告グループ、広告、インサイト (合計 11 のツール) は読み取り専用です。 Claude、Codex、Cursor などのキャンペーンを平易な言語でクエリできます。
次は書き込み (「アクション」) です。ここで 2 つのアイデアがあります。
- 決定論的ツールでラップします (json in/out)
- ループに人間を追加します (おそらくクエリとツール/JSON 出力の間の検証/一貫性チェック ステップも) Claude Code がほとんどの作業を実行しましたが、それはワンショットではありませんでした。 OpenAI API 仕様は非常に簡潔で理解しやすいものです。何らかの理由で、Claude は OpenAI 自身のドキュメントを読むことができなかったので、私は手動で「コピー＆ペースト」することになりました。また、少し繰り返して、ローカルの stdio インストールから npm に移行し、それを MCP レジストリで利用できるようにしました。特に書き込み/引き出し側およびリモート バージョンのホストに関するフィードバックは歓迎です。 PS: これは (明らかに) 非公式であり、OpenAI とは関係ありません。

記事本文:
GitHub - HYPD-AI/openai-ads-mcp: OpenAI 広告 / ChatGPT 広告用の MCP サーバー。 OpenAI Ads Advertiser API を通じて、キャンペーン、広告グループ、広告、クリエイティブ、インサイトを管理します。 · GitHub
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
HYPD-AI
/
openai-ads-mcp
公共
通知
あなたはきっとsiでしょう

通知設定を変更するためにログインしました
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット .github/ workflows .github/ workflows src src test test .env.example .env.example .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json COTRIBUTING.md COTRIBUTING.md ライセンスライセンス README.md README.md eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
OpenAI Ads (広告主) API の Model Context Protocol (MCP) サーバー。これにより、MCP 互換クライアント (Claude Desktop、Cursor、VS Code など) は、自然言語を通じて OpenAI Ads キャンペーン、広告グループ、広告、パフォーマンスに関する分析情報を読み取ることができます。
読み取り専用。この最初のリリースはデータの読み取りのみを行い、何も作成、編集、一時停止することはなく、予算を消費することもありません。書き込みアクションはロードマップにあります。
非公式。これはコミュニティ プロジェクトであり、OpenAI と提携または承認されていません。免責事項を参照してください。
OpenAI Ads API は、広告主のアカウント、キャンペーン、広告グループ、広告、レポートを公開します。このサーバーは、その API の読み取りエンドポイントを MCP ツールとしてラップするため、AI アシスタントは次のような質問に答えることができます。
「私の OpenAI Ads API キーは機能していますか?どのアカウントに関連付けられていますか?」
「アクティブなキャンペーンとその予算をリストします。」
「過去 30 日間のキャンペーン cmp_123 の費用、クリック数、CTR を日ごとに表示します。」
「広告グループ adg_456 内のどの広告がまだ審査待ちですか?」
アカウント、キャンペーン、広告グループ、広告、あらゆるレベルの分析情報をカバーする 11 個の読み取り専用ツール。
忠実な応答 - API の JSON がそのまま返されるため、

翻訳中に何かが失われます。
エラーのクリア — HTTP ステータスと API エラー本文は、飲み込まれるのではなく、モデルに表示されます。
Micros 対応 — すべてのツールの説明には Micros の規則が説明されているため、アシスタントは人間が判読できる通貨を表示できます。
カーソルのページネーション パススルー ( l​​imit 、 order 、 after 、 before )。
npx 経由のゼロインストール — npx -y @hypd-ai/openai-ads-mcp 、クローンやビルドはありません。
ツール
何をするのか
広告アカウントの取得
構成されたキーの広告アカウントを取得します。接続チェックとして最適です。
list_campaigns
キャンペーンをリストします (目的、予算、ターゲット国)。
ゲットキャンペーン
ID によって 1 つのキャンペーンを取得します。
リスト広告グループ
広告グループをリストします。オプションでキャンペーンごとにフィルタリングできます。
get_ad_group
ID によって 1 つの広告グループを取得します (入札設定、コンテキスト ヒント)。
リスト広告
リスト広告。オプションで広告グループごとにフィルタリングできます。
get_ad
ID (クリエイティブ + レビュー ステータス) によって 1 つの広告を取得します。
get_account_insights
アカウント全体のパフォーマンスに関する洞察。
get_campaign_insights
1 つのキャンペーンのパフォーマンスに関する分析情報。
get_ad_group_insights
1 つの広告グループのパフォーマンスに関する分析情報。
get_ad_insights
1 つの広告のパフォーマンスに関する分析情報。
Insights ツールは、レポート ウィンドウの [since / until (YYYY-MM-DD)] に加えて、 time_granularity ( daily / none )、 aggregation_level 、fields 、 sort 、 filters 、limit (1–10000)、およびカーソルの後 / 前を受け入れます。
OpenAI Ads API キー。 ads.openai.com (現在は米国のみ) で広告アカウントを作成し、[設定] → [ads.openai.com/settings] からキーを発行します。クイックスタートと認証のドキュメントを参照してください。各キーのスコープは 1 つの広告アカウントに限定されます。
MCP クライアントはサーバーをサブプロセスとして起動し、環境変数を介して API キーを渡します。
@hypd-ai/openai-ads-mcp として npm に公開 — npx がそれを取得するため、クローンを作成したりビルドしたりする必要はありません。最新の未リリースのメインを実行するには

代わりに、 @hypd-ai/openai-ads-mcp を github:HYPD-AI/openai-ads-mcp に置き換えてください (最初の起動はソースからビルドされます。ソースからの実行を参照してください)。
以下にクライアント用のスニペットを追加します。
claude_desktop_config.json を編集します。
macOS: ~/ライブラリ/Application Support/Claude/claude_desktop_config.json
Windows: %APPDATA%\Claude\claude_desktop_config.json
{
"mcpサーバー": {
"openai-ads" : {
"コマンド" : " npx " ,
"args" : [ " -y " , " @hypd-ai/openai-ads-mcp " ],
"環境" : {
"OPENAI_ADS_API_KEY" : " your-openai-ads-API-key "
}
}
}
}
Claude Desktop を再起動し、「openai-ads ツールを使用して広告アカウントを検索してください。」と尋ねます。
~/.cursor/mcp.json (グローバル) または .cursor/mcp.json (プロジェクトごと) に追加します。
{
"mcpサーバー": {
"openai-ads" : {
"コマンド" : " npx " ,
"args" : [ " -y " , " @hypd-ai/openai-ads-mcp " ],
"環境" : {
"OPENAI_ADS_API_KEY" : " your-openai-ads-API-key "
}
}
}
}
VSコード
.vscode/mcp.json に追加します。 VS Code はキーの入力を要求し、それを入力経由でシークレットとして保存できます。
{
「入力」: [
{
"type" : "プロンプト文字列" ,
"id" : " openai_ads_api_key " ,
"説明" : " OpenAI Ads API キー " ,
「パスワード」: true
}
]、
「サーバー」: {
"openai-ads" : {
"コマンド" : " npx " ,
"args" : [ " -y " , " @hypd-ai/openai-ads-mcp " ],
"環境" : {
"OPENAI_ADS_API_KEY" : " ${input:openai_ads_api_key} "
}
}
}
}
他の MCP クライアント
標準入出力経由で MCP を話すクライアントはすべて機能します。環境に OPENAI_ADS_API_KEY を設定して、 npx -y @hypd-ai/openai-ads-mcp (またはノード /path/to/dist/index.js ) を実行します。
変数
必須
デフォルト
説明
OPENAI_ADS_API_KEY
はい
—
OpenAI Ads API キー。Bearer トークンとして送信されます。
OPENAI_ADS_BASE_URL
いいえ
https://api.ads.openai.com/v1
API ベース URL をオーバーライドします (テストまたはプロキシに役立ちます)。
.env.example を参照してください。
名前が _micros で終わるフィールド — たとえば、キャンペーンの Lifetime_spend_limit_micros や

広告グループの max_bid_micros — マイクロ単位で表されます。
1,000,000 マイクロ = アカウントの通貨の 1 単位 (例: $1.00 = 1,000,000 マイクロ)
したがって、lifetime_spend_limit_micros の 25000000 は $25.00 です。人間の量を表示するには _micros 値を 1,000,000 で除算し、逆に変換するには 1,000,000 を乗算します。
インサイトのメトリクスはミクロではありません。 Spend 、 cpc 、 cpm などのレポート値は、すでにアカウントの通貨で小数点として表示されています (例: Spend: 42.75 は $42.75 を意味します)。
このリリースでは読み取り ( GET ) ツールのみが登録されており、各ツールには MCP readOnlyHint の注釈が付けられているため、行儀の良いクライアントは状態を変更できないことがわかります。ここには、何かを作成、編集、一時停止、削除できるツールや、予算を消費できるツールはありません。書き込みアクションは、意図的に個別にレビューされたステップとして導入されます (「ロードマップ」を参照)。
git clone https://github.com/hypd-ai/openai-ads-mcp.git
cd openai-ads-mcp
npmインストール
npm ビルドを実行する
次に、MCP クライアントに構築されたエントリ ファイルを指定します。
{
"mcpサーバー": {
"openai-ads" : {
"コマンド" : "ノード" ,
"args" : [ " /absolute/path/to/openai-ads-mcp/dist/index.js " ],
"環境" : {
"OPENAI_ADS_API_KEY" : " your-openai-ads-API-key "
}
}
}
}
MCP Inspector で試してみる
OPENAI_ADS_API_KEY=あなたのキー npx @modelcontextprotocol/インスペクター ノード dist/index.js
開発
npm install # 依存関係をインストールする
npm run dev # 変更時にリビルド (tsup --watch)
npm run typecheck # tsc --noEmit
npm 実行 lint # eslint
npm 実行形式 # prettier --write
npm テスト # vitest
npm run build # dist/ へのバンドル
プロジェクトのレイアウト:
ソース/
Index.ts # bin エントリ: 構成をロードし、サーバーを構築し、stdio に接続します
server.ts # buildServer(): McpServer + すべてのツールを登録
client.ts # OpenAIAdsClient: 認証、URL 構築、エラー
config.ts # 環境の解析と検証
schemas.ts # 共有 zod シェイプ (ページネーション、in)

観光スポット）+マイクロノート
tools/ # リソースごとに 1 つのファイル (アカウント、キャンペーン、広告グループ、広告、インサイト)
test/ # vitest 仕様 (構成、クライアント、メモリ内サーバー)
ツールが API にどのようにマッピングされるか
すべてのエンドポイントはベース URL (デフォルト https://api.ads.openai.com/v1 ) の下にあります。
✍️ アクションの作成 - キャンペーン、広告グループ、広告の作成と更新 ( POST 経由)、さらに専用の状態遷移 ( POST .../activate 、 .../pause 、 .../archive )。 HTTP クライアントはすでに POST をサポートしています。これらは配信と支出を変更するため、明示的なオプトインの背後でゲートされます。
🖼️ クリエイティブのアップロード — POST /upload (JSON image_url または multipart/form-data ) を使用して画像を広告クリエイティブに添付します。
🌍 キャンペーンのターゲティング — 国を含める/除外します (targeting.locations.countries )。
🌐 ホストされた展開のためのリモート/HTTP トランスポート。
📦 npm リリースが公開され、npx -y openai-ads-mcp がそのまま動作するようになりました。
貢献は大歓迎です! COTRIBUTING.md をお読みください。つまり、重要な変更について議論するためにイシューを開いて、npm run lint && npm run typecheck && npm test green を維持し、新しい動作のテストを追加します。
これは非公式のコミュニティ構築プロジェクトです。 OpenAI と提携、承認、後援されているわけではありません。 「OpenAI」および関連する名前およびロゴは、OpenAI の商標です。このツールを介した OpenAI Ads API の使用には、OpenAI の利用規約とポリシーが適用されます。このツールは「現状のまま」提供され、いかなる種類の保証もありません。ライセンスを参照してください。
OpenAI 広告 / ChatGPT 広告用の MCP サーバー。 OpenAI Ads Advertiser API を通じて、キャンペーン、広告グループ、広告、クリエイティブ、インサイトを管理します。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
1
フォーク
レポートリポジトリ
リリース
2
v0.1.1 — 現在は npm にあります
最新の
2026 年 6 月 3 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました

。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

MCP server for OpenAI Ads / ChatGPT Ads. Manage campaigns, ad groups, ads, creatives, and insights through the OpenAI Ads Advertiser API. - HYPD-AI/openai-ads-mcp

I built a small MCP server for the ChatGPT Ads advertiser API, which OpenAI opened to the public a couple of weeks ago. It's read-only for now: accounts, campaigns, ad groups, ads and insights (11 tools in total). You can query campaigns from Claude, Codex, Cursor and so on in plain language.
Writes ("actions") are next; I have two ideas here:
- wrap them in a deterministic tool (json in/out)
- add a human in the loop (maybe even a verification/consistency check step between query and tool/json output) Claude Code did most of the work but it wasn't a one-shot. The OpenAI API spec is very clean though and easy to understand; for some reason Claude couldn't read OpenAI's own docs, so I ended up with some manual "copy&paste". Also I iterated a bit and went from a local stdio install to npm and then made it available in the MCP registry. Feedback welcome, especially on the write/elicitation side and on hosting a remote version. PS: It's (obviously) unofficial and not affiliated with OpenAI.

GitHub - HYPD-AI/openai-ads-mcp: MCP server for OpenAI Ads / ChatGPT Ads. Manage campaigns, ad groups, ads, creatives, and insights through the OpenAI Ads Advertiser API. · GitHub
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
HYPD-AI
/
openai-ads-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits .github/ workflows .github/ workflows src src test test .env.example .env.example .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
A Model Context Protocol (MCP) server for the OpenAI Ads (Advertiser) API . It lets MCP-compatible clients — Claude Desktop, Cursor, VS Code, and others — read your OpenAI Ads campaigns, ad groups, ads, and performance insights through natural language.
Read-only. This first release only reads data — it never creates, edits, or pauses anything and never spends budget. Write actions are on the roadmap .
Unofficial. This is a community project and is not affiliated with or endorsed by OpenAI . See the disclaimer .
The OpenAI Ads API exposes an advertiser's account, campaigns, ad groups, ads, and reporting. This server wraps the read endpoints of that API as MCP tools so an AI assistant can answer questions like:
"Is my OpenAI Ads API key working? What account is it tied to?"
"List my active campaigns and their budgets."
"Show spend, clicks, and CTR for campaign cmp_123 over the last 30 days, by day."
"Which ads in ad group adg_456 are still pending review?"
11 read-only tools covering the account, campaigns, ad groups, ads, and insights at every level.
Faithful responses — the API's JSON is returned as-is, so nothing is lost in translation.
Clear errors — HTTP status and the API error body are surfaced to the model instead of being swallowed.
Micros-aware — every tool description explains the micros convention so the assistant can present human-readable currency.
Cursor pagination passthrough ( limit , order , after , before ).
Zero-install via npx — npx -y @hypd-ai/openai-ads-mcp , no clone or build.
Tool
What it does
get_ad_account
Fetch the ad account for the configured key. Great as a connectivity check.
list_campaigns
List campaigns (objective, budget, country targeting).
get_campaign
Fetch a single campaign by ID.
list_ad_groups
List ad groups, optionally filtered by campaign.
get_ad_group
Fetch a single ad group by ID (bidding config, context hints).
list_ads
List ads, optionally filtered by ad group.
get_ad
Fetch a single ad by ID (creative + review status).
get_account_insights
Performance insights for the whole account.
get_campaign_insights
Performance insights for one campaign.
get_ad_group_insights
Performance insights for one ad group.
get_ad_insights
Performance insights for one ad.
Insights tools accept since / until (YYYY-MM-DD) for the reporting window, plus time_granularity ( daily / none ), aggregation_level , fields , sort , filters , limit (1–10000), and after / before cursors.
An OpenAI Ads API key. Create an Ads account at ads.openai.com (currently US-only ), then issue a key from Settings → ads.openai.com/settings . See the quickstart and authentication docs. Each key is scoped to a single ad account .
MCP clients launch the server as a subprocess and pass your API key via an environment variable.
Published on npm as @hypd-ai/openai-ads-mcp — npx fetches it for you, so there's nothing to clone or build. To run the latest unreleased main instead, replace @hypd-ai/openai-ads-mcp with github:HYPD-AI/openai-ads-mcp (its first launch builds from source — see Running from source ).
Add the snippet for your client below.
Edit your claude_desktop_config.json :
macOS: ~/Library/Application Support/Claude/claude_desktop_config.json
Windows: %APPDATA%\Claude\claude_desktop_config.json
{
"mcpServers" : {
"openai-ads" : {
"command" : " npx " ,
"args" : [ " -y " , " @hypd-ai/openai-ads-mcp " ],
"env" : {
"OPENAI_ADS_API_KEY" : " your-openai-ads-api-key "
}
}
}
}
Restart Claude Desktop, then ask: "Use the openai-ads tools to look up my ad account."
Add to ~/.cursor/mcp.json (global) or .cursor/mcp.json (per-project):
{
"mcpServers" : {
"openai-ads" : {
"command" : " npx " ,
"args" : [ " -y " , " @hypd-ai/openai-ads-mcp " ],
"env" : {
"OPENAI_ADS_API_KEY" : " your-openai-ads-api-key "
}
}
}
}
VS Code
Add to .vscode/mcp.json . VS Code can prompt for the key and store it as a secret via inputs :
{
"inputs" : [
{
"type" : " promptString " ,
"id" : " openai_ads_api_key " ,
"description" : " OpenAI Ads API key " ,
"password" : true
}
],
"servers" : {
"openai-ads" : {
"command" : " npx " ,
"args" : [ " -y " , " @hypd-ai/openai-ads-mcp " ],
"env" : {
"OPENAI_ADS_API_KEY" : " ${input:openai_ads_api_key} "
}
}
}
}
Other MCP clients
Any client that speaks MCP over stdio works. Run npx -y @hypd-ai/openai-ads-mcp (or node /path/to/dist/index.js ) with OPENAI_ADS_API_KEY set in the environment.
Variable
Required
Default
Description
OPENAI_ADS_API_KEY
Yes
—
Your OpenAI Ads API key, sent as a Bearer token.
OPENAI_ADS_BASE_URL
No
https://api.ads.openai.com/v1
Override the API base URL (useful for testing or a proxy).
See .env.example .
Fields whose names end in _micros — for example a campaign's lifetime_spend_limit_micros or an ad group's max_bid_micros — are expressed in micros:
1,000,000 micros = 1 unit of the account's currency (e.g. $1.00 = 1,000,000 micros)
So a lifetime_spend_limit_micros of 25000000 is $25.00 . Divide a _micros value by 1,000,000 to display a human amount, or multiply by 1,000,000 to convert the other way.
Insights metrics are not micros. Reporting values like spend , cpc , and cpm are already in the account's currency as decimals (e.g. spend: 42.75 means $42.75).
This release registers only read ( GET ) tools — and each one is annotated with the MCP readOnlyHint , so well-behaved clients know it cannot mutate state. There is no tool here that can create, edit, pause, or delete anything, and nothing that can spend budget. Write actions will arrive as a deliberate, separately reviewed step (see Roadmap ).
git clone https://github.com/hypd-ai/openai-ads-mcp.git
cd openai-ads-mcp
npm install
npm run build
Then point your MCP client at the built entry file:
{
"mcpServers" : {
"openai-ads" : {
"command" : " node " ,
"args" : [ " /absolute/path/to/openai-ads-mcp/dist/index.js " ],
"env" : {
"OPENAI_ADS_API_KEY" : " your-openai-ads-api-key "
}
}
}
}
Try it with the MCP Inspector
OPENAI_ADS_API_KEY=your-key npx @modelcontextprotocol/inspector node dist/index.js
Development
npm install # install dependencies
npm run dev # rebuild on change (tsup --watch)
npm run typecheck # tsc --noEmit
npm run lint # eslint
npm run format # prettier --write
npm test # vitest
npm run build # bundle to dist/
Project layout:
src/
index.ts # bin entry: load config, build server, connect stdio
server.ts # buildServer(): McpServer + register all tools
client.ts # OpenAIAdsClient: auth, URL building, errors
config.ts # environment parsing & validation
schemas.ts # shared zod shapes (pagination, insights) + micros note
tools/ # one file per resource (account, campaigns, ad-groups, ads, insights)
test/ # vitest specs (config, client, in-memory server)
How tools map to the API
All endpoints are under the base URL (default https://api.ads.openai.com/v1 ).
✍️ Write actions — create & update (via POST ) campaigns, ad groups, and ads, plus the dedicated state transitions ( POST .../activate , .../pause , .../archive ). The HTTP client already supports POST ; these will be gated behind an explicit opt-in, since they change delivery and spend.
🖼️ Creative uploads — POST /upload (JSON image_url or multipart/form-data ) to attach images to ad creatives.
🌍 Campaign targeting — country include/exclude ( targeting.locations.countries ).
🌐 Remote/HTTP transport for hosted deployments.
📦 Published npm release so npx -y openai-ads-mcp works out of the box.
Contributions are welcome! Please read CONTRIBUTING.md . In short: open an issue to discuss substantial changes, keep npm run lint && npm run typecheck && npm test green, and add tests for new behavior.
This is an unofficial , community-built project. It is not affiliated with, endorsed by, or sponsored by OpenAI . "OpenAI" and related names and logos are trademarks of OpenAI. Your use of the OpenAI Ads API through this tool is subject to OpenAI's terms and policies. The tool is provided "as is", without warranty of any kind — see the license .
MCP server for OpenAI Ads / ChatGPT Ads. Manage campaigns, ad groups, ads, creatives, and insights through the OpenAI Ads Advertiser API.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
2
v0.1.1 — now on npm
Latest
Jun 3, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
