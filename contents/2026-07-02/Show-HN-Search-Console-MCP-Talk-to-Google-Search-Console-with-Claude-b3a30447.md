---
source: "https://github.com/sudomichael/search-console-mcp"
hn_url: "https://news.ycombinator.com/item?id=48764726"
title: "Show HN: Search Console MCP – Talk to Google Search Console with Claude"
article_title: "GitHub - sudomichael/search-console-mcp: Google Search Console for Claude, Cursor & any MCP client — one sign-in, 30 seconds, no Google Cloud project. Read-only by construction. · GitHub"
author: "ushercakes"
captured_at: "2026-07-02T17:52:01Z"
capture_tool: "hn-digest"
hn_id: 48764726
score: 1
comments: 0
posted_at: "2026-07-02T17:30:19Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Search Console MCP – Talk to Google Search Console with Claude

- HN: [48764726](https://news.ycombinator.com/item?id=48764726)
- Source: [github.com](https://github.com/sudomichael/search-console-mcp)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T17:30:19Z

## Translation

タイトル: HN を表示: Search Console MCP – クロードと Google Search Console と話す
記事のタイトル: GitHub - sudomichael/search-console-mcp: クロード、カーソル、および任意の MCP クライアント用の Google Search Console — 1 回のサインイン、30 秒、Google Cloud プロジェクトなし。構築により読み取り専用になります。 · GitHub
説明: Claude、Cursor、および任意の MCP クライアント用の Google Search Console — 1 回のログイン、30 秒、Google Cloud プロジェクトなし。構築により読み取り専用になります。 - sudomichael/search-console-mcp
HN本文：無料でOSSです！ GSC は貴重な SEO データの宝庫ですが、データは手動で取得するのが面倒です。いくつかの異なるページに移動し、いくつかの異なる場所から「エクスポート」ボタンをクリックする必要があります。 MCP を使用すると、ダウンロードの煩わしさを感じることなく、最初から認証するだけで簡単にデータとチャットできます。使ったり、嫌いになったり、気に入ったり - より良いものにするために私を手伝ってください!

記事本文:
GitHub - sudomichael/search-console-mcp: クロード、カーソル、および任意の MCP クライアント用の Google Search Console — 1 回のサインイン、30 秒、Google Cloud プロジェクトなし。構築により読み取り専用になります。 · GitHub
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
スドマイケル
/
検索コンソール-mcp
公共
通知

ns
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
sudomichael/search-console-mcp
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
16 コミット 16 コミット .serena .serena src src .gitignore .gitignore ライセンス ライセンス README.md README.md lhm.plugin.json lhm.plugin.json package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json すべてのファイルを表示リポジトリ ファイルのナビゲーション
Search Console はダッシュボード用に構築されました。これにより、代わりに会話が可能になります。
Claude、Cursor、または任意の MCP クライアントから Google Search Console に話しかけます。ログインは 1 回、30 秒で、Google Cloud プロジェクトは必要ありません。
Googleサーチコンソール
│
▼
search-console-mcp ← マシン上で読み取り専用で実行されます
│
▼
クロード / カーソル / 任意の MCP クライアント
│
▼
「トラフィックが減少したのはなぜですか?」
他のすべての Search Console MCP サーバーでは、Google Cloud プロジェクトを作成し、API を有効にし、サービス アカウント JSON をラングリングします。これはそうではありません:
npx search-console-mcp-server ログイン
ブラウザが開き、Google でサインインすれば完了です。トークンは Google によってあなたのマシンに直接作成され、~/.search-console-mcp/ にのみ保存されます。誰のサーバーも通過しません。
あなた: 私のページの中で、ページ 1 に最も近いのはどれですか?
クロード: 重要な印象に応じた 3 つの実際の機会:
1. /pricing — 「スコアボード アプリの価格設定」は 11.2 位にあります。
2,400 インプレッション/月、クリックはほとんどありません。 1つのポジショングループ
ジャンプは月に約 90 クリックの価値があります。タイトルには含まれていません
クエリ;そこから始めてください。
2. /templates/trivia — 「トリビア スコアボード」の 12.8 位、
1,100 インプレッション/月。ページはランク付けされていますが、H2 一致がありません
という意図。
3. /blog/keep-score-online — 9.4 位、すでに 1 ページ目
いくつかの日。 /features と共食いしている — 実行

触れる前にcannibalization_checkしてください。
(出力例 — 形状と数値は、ツールが実際のデータから返すものです。)
5 つの既製の分析が MCP プロンプト (それらをサポートするクライアントのスラッシュ コマンド) として出荷されます。 API 呼び出しではなく結果:
npx search-console-mcp-server ログイン
2. MCP クライアントに以下を追加します。
クロード mcp add gsc -- npx -y search-console-mcp-server
クロード デスクトップ / カーソル / その他 (mcpServers 設定):
{
"mcpサーバー": {
"gsc" : { "コマンド" : " npx " , "args" : [ " -y " , " search-console-mcp-server " ] }
}
}
3. 次のように尋ねます。
「今月と先月で最もクリック数が減ったページはどれですか?」
「実際のインプレッションを含む 8 ～ 15 位にある私のクエリを表示します。1 ページ目に最も近いものはどれですか?」
「https://mysite.com/pricing はインデックスに登録されていますか? インデックスに登録されていない場合は、なぜですか?」
「Google Discover のトラフィックはどのくらいですか?」
ツール
何をするのか
リストプロパティ
GSC プロパティ (最初に呼び出して、正確な siteUrl 値を提供します)
クエリ検索_分析
クエリ、ページ、日付、国、デバイス別のクリック数 / インプレッション数 / CTR / 順位 - フィルタ、ページネーション、Google Discover を含むすべてのトラフィック ソース
比較期間
最大の変動要因を自動的に計算: この期間と前の期間、ページごとまたはクエリごとのデルタ、変更ごとに並べ替え
検査URL
URL の Google のインデックス ステータス: インデックス済みかどうか、その理由、正規として選択されたもの、最後のクロール
検査URL
同じバッチ処理 — インデックス作成監査の 1 回の呼び出しで最大 10 個の URL
list_sitemaps
送信されたサイトマップとステータス
構造による読み取り専用: このツールが要求する Google スコープ ( webmasters.readonly ) は、プロパティの変更、サイトマップの送信、または何も変更できません。「デフォルトで無効」ではありません。不可能。
Compare_periods 呼び出しがモデルに渡す内容 (モデルは日付演算を行いません)。
{ "キー" : " https://example.com/pricing " ,
"現在" : { "クリック数" : 1040 , "インプレッション数"

: 20502 , "位置" : 6.4 },
"前" : { "クリック数" : 1070 、 "インプレッション数" : 17176 、 "掲載順位" : 5.8 },
"deltaClicks" : -30 、 "deltaImpressions" : 3326 、 "deltaPosition" : 0.6 }
プライバシー
読み取り専用の Google スコープ ( webmasters.readonly )。
トークンは、マシン ( chmod 600 ) の ~/.search-console-mcp/credentials.json に存在します。 npx search-console-mcp-server logout で削除されます。
テレメトリもプロキシもありません。API 呼び出しはマシンから Google に完全に送信されます。
独自の Google Cloud プロジェクトをご希望ですか? SEARCH_CONSOLE_MCP_CLIENT_ID / SEARCH_CONSOLE_MCP_CLIENT_SECRET を設定すると、あなたのものが使用されます。
Google 認証情報を作成せずにログインはどのように機能しますか?
CLI には、Google 「デスクトップ アプリ」OAuth クライアントが同梱されています。これは、gcloud と GitHub の CLI が使用するものと同じサポートされているモデルです。トークンは引き続き Google によってマシンに直接作成されます (PKCE + localhost コールバック)。彼らは私たちのサーバーには決して触れません。認証パス全体は src/auth.ts の約 200 行です。読んでください。何も共有しないようにしたいですか? SEARCH_CONSOLE_MCP_CLIENT_ID / SEARCH_CONSOLE_MCP_CLIENT_SECRET を独自の資格情報で設定します。
Google アカウントを切り替えますか?
npx search-console-mcp-server 再度ログイン — Google にアカウント ピッカーが表示されます。
MCP クライアントで「サインインしていません」エラーが発生しますか?
クライアントを実行しているのと同じユーザー アカウントからログインを実行します。認証情報は ~/.search-console-mcp/ にあります。
list_properties にプロパティがありません。
ログインしている Google アカウントには、少なくとも Search Console で「制限付き」アクセスが必要です。
昨日がデータから欠落しているのはなぜですか?
Google Search Console のデータには最大 2 ～ 3 日の遅れがあります。それは Google であり、ツールではありません。歴史は約 16 か月前に遡ります。
アクセスを取り消しますか?
npx search-console-mcp-server logout はローカル トークンを削除します。 myaccount.google.com/permissions は付与自体を取り消します。
積極的なものをお探しですか?
この MCP サーバーは、調査に役立ちます。

sk、答えます。
1 週間後、同じ質問を繰り返していることに気づくかもしれません。「何が落ちたのか?」 1 ページ目に最も近いものは何ですか?何を直せばいいのでしょうか？
パーシーは残りの半分です。彼は毎週自分で Search Console をチェックし、戦略を立て、変更すべき内容を正確にメールで送信します。その後、ライブ ページを再チェックして、それが機能したかどうかを測定します。
ここで探検。そこで処刑。同じ哲学: 実際のデータ、平易な英語、ダッシュボードなし。
MITライセンス取得済み。問題やPRを歓迎します。パーシーによって建てられました。
Claude、Cursor、および任意の MCP クライアント用の Google Search Console — 1 回のログイン、30 秒、Google Cloud プロジェクトなし。構築により読み取り専用になります。
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

Google Search Console for Claude, Cursor & any MCP client — one sign-in, 30 seconds, no Google Cloud project. Read-only by construction. - sudomichael/search-console-mcp

It's free and OSS! GSC is a goldmine on valuable SEO data, but - the data is manual and kind of annoying to grab. You have to go to a few different pages, and click the "Export" button from several different locations. MCP makes it easy to just oauth in, chat with the data right from the get-go without any download annoyances. Use it, hate it, love it - help me make it better!

GitHub - sudomichael/search-console-mcp: Google Search Console for Claude, Cursor & any MCP client — one sign-in, 30 seconds, no Google Cloud project. Read-only by construction. · GitHub
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
sudomichael
/
search-console-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
sudomichael/search-console-mcp
main Branches Tags Go to file Code Open more actions menu Folders and files
16 Commits 16 Commits .serena .serena src src .gitignore .gitignore LICENSE LICENSE README.md README.md lhm.plugin.json lhm.plugin.json package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json View all files Repository files navigation
Search Console was built for dashboards. This gives you a conversation instead.
Talk to your Google Search Console from Claude, Cursor, or any MCP client — one sign-in, 30 seconds, no Google Cloud project.
Google Search Console
│
▼
search-console-mcp ← runs on your machine, read-only
│
▼
Claude / Cursor / any MCP client
│
▼
"Why did my traffic drop?"
Every other Search Console MCP server makes you create a Google Cloud project, enable APIs, and wrangle service-account JSON. This one doesn't:
npx search-console-mcp-server login
Your browser opens, you sign in with Google, done. Tokens are minted by Google directly to your machine and stored only in ~/.search-console-mcp/ — nothing passes through anyone's servers.
You: Which of my pages are closest to page one?
Claude: Three real opportunities, by impressions at stake:
1. /pricing — "scoreboard app pricing" sits at position 11.2 with
2,400 impressions/mo and almost no clicks. One position-group
jump is worth ~90 clicks/mo. The title doesn't contain the
query; start there.
2. /templates/trivia — position 12.8 on "trivia scoreboard",
1,100 impressions/mo. The page ranks but has no H2 matching
the intent.
3. /blog/keep-score-online — position 9.4, already page one on
some days. It's cannibalizing with /features — run
cannibalization_check before touching it.
(Illustrative output — shape and numbers are what the tools return from your real data.)
Five ready-made analyses ship as MCP prompts — slash-commands in clients that support them. Outcomes, not API calls:
npx search-console-mcp-server login
2. Add to your MCP client:
claude mcp add gsc -- npx -y search-console-mcp-server
Claude Desktop / Cursor / anything else ( mcpServers config):
{
"mcpServers" : {
"gsc" : { "command" : " npx " , "args" : [ " -y " , " search-console-mcp-server " ] }
}
}
3. Ask:
"Which of my pages lost the most clicks this month vs last?"
"Show my queries sitting at position 8–15 with real impressions — what's closest to page 1?"
"Is https://mysite.com/pricing indexed? If not, why?"
"How much Google Discover traffic do I get?"
Tool
What it does
list_properties
Your GSC properties (call first — gives exact siteUrl values)
query_search_analytics
Clicks / impressions / CTR / position by query, page, date, country, device — filters, pagination, and every traffic source including Google Discover
compare_periods
Biggest movers, computed for you: this period vs the prior one, per-page or per-query deltas, sorted by change
inspect_url
Google's index status for a URL: indexed or not, why, chosen canonical, last crawl
inspect_urls
The same, batched — up to 10 URLs in one call for indexing audits
list_sitemaps
Submitted sitemaps with status
Read-only by construction: the Google scope this tool requests ( webmasters.readonly ) cannot modify your properties, submit sitemaps, or change anything — not "disabled by default." Impossible.
What a compare_periods call hands your model (the model never does date arithmetic):
{ "key" : " https://example.com/pricing " ,
"current" : { "clicks" : 1040 , "impressions" : 20502 , "position" : 6.4 },
"previous" : { "clicks" : 1070 , "impressions" : 17176 , "position" : 5.8 },
"deltaClicks" : -30 , "deltaImpressions" : 3326 , "deltaPosition" : 0.6 }
Privacy
Read-only Google scope ( webmasters.readonly ).
Tokens live in ~/.search-console-mcp/credentials.json on your machine, chmod 600 . npx search-console-mcp-server logout deletes them.
No telemetry, no proxy — API calls go from your machine to Google, full stop.
Prefer your own Google Cloud project? Set SEARCH_CONSOLE_MCP_CLIENT_ID / SEARCH_CONSOLE_MCP_CLIENT_SECRET and it uses yours.
How can the login work without me creating Google credentials?
The CLI ships a Google "Desktop app" OAuth client — the same supported model gcloud and GitHub's CLI use. Your tokens are still minted by Google directly to your machine (PKCE + localhost callback); they never touch our servers. The entire auth path is ~200 lines in src/auth.ts — read it. Want zero shared anything? Set SEARCH_CONSOLE_MCP_CLIENT_ID / SEARCH_CONSOLE_MCP_CLIENT_SECRET with your own credentials.
Switch Google accounts?
npx search-console-mcp-server login again — Google shows the account picker.
"Not signed in" errors in my MCP client?
Run the login from the same user account your client runs under; credentials live in ~/.search-console-mcp/ .
A property is missing from list_properties .
The signed-in Google account needs at least "Restricted" access to it in Search Console.
Why is yesterday missing from the data?
Google Search Console data lags ~2–3 days. That's Google, not the tool. History goes back ~16 months.
Revoke access?
npx search-console-mcp-server logout deletes local tokens; myaccount.google.com/permissions revokes the grant itself.
Looking for something proactive?
This MCP server helps you investigate — you ask, it answers.
After a week you may notice you keep asking the same questions: What dropped? What's closest to page one? What should I fix?
Percy is the other half: he checks your Search Console every week on his own, builds a strategy, emails you exactly what to change — then re-checks the live page and measures whether it worked.
Exploration here. Execution there. Same philosophy: real data, plain English, no dashboards.
MIT licensed. Issues and PRs welcome. Built by Percy .
Google Search Console for Claude, Cursor & any MCP client — one sign-in, 30 seconds, no Google Cloud project. Read-only by construction.
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
