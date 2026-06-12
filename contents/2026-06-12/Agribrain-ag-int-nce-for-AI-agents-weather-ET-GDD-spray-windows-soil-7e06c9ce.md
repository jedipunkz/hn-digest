---
source: "https://github.com/VasileiosTs/agribrain"
hn_url: "https://news.ycombinator.com/item?id=48502078"
title: "Agribrain / ag-int/nce for AI agents (weather, ET₀, GDD, spray windows, soil)"
article_title: "GitHub - VasileiosTs/agribrain: Agronomic intelligence for AI agents / weather, ET₀, GDD, spray windows, soil. Zero API keys. MCP server + core library. · GitHub"
author: "vasileiosts"
captured_at: "2026-06-12T10:34:42Z"
capture_tool: "hn-digest"
hn_id: 48502078
score: 2
comments: 1
posted_at: "2026-06-12T10:04:38Z"
tags:
  - hacker-news
  - translated
---

# Agribrain / ag-int/nce for AI agents (weather, ET₀, GDD, spray windows, soil)

- HN: [48502078](https://news.ycombinator.com/item?id=48502078)
- Source: [github.com](https://github.com/VasileiosTs/agribrain)
- Score: 2
- Comments: 1
- Posted: 2026-06-12T10:04:38Z

## Translation

タイトル: AI エージェント用の Agribrain / ag-int/nce (気象、ET₀、GDD、スプレー ウィンドウ、土壌)
記事タイトル: GitHub - VasileiosTs/agribrain: AI エージェントのための農業インテリジェンス / 天気、ET₀、GDD、スプレー ウィンドウ、土壌。 API キーはゼロです。 MCPサーバー+コアライブラリ。 · GitHub
説明: AI エージェント用の農業インテリジェンス / 天気、ET₀、GDD、スプレー ウィンドウ、土壌。 API キーはゼロです。 MCPサーバー+コアライブラリ。 - VasileiosTs/アグリブレイン

記事本文:
GitHub - VasileiosTs/agribrain: AI エージェント用の農業インテリジェンス / 天気、ET₀、GDD、スプレー ウィンドウ、土壌。 API キーはゼロです。 MCPサーバー+コアライブラリ。 · GitHub
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
ヴァシリオス
/
アグリブレイン
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット データ データの例 例 パッケージ パッケージ .gitignore .gitignore ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
AI アシスタントに認可を受けた農学者の頭脳を与えましょう。
[デモ GIF はここにあります — 15 秒: 実際のフィールド座標 → 「スプレーの状況はどうなっていますか?」
今週のオリーブの水分バランスは？」 →数字で表す本当の答え]
すべての LLM はオリーブの木についての詩を書くことができます。彼らは誰もあなたのことを知りません
オリーブミバエ三代目は火曜日から始まり、土曜日の風が吹いて
スプレーしても無意味、あるいはフィールドが水面から 12mm 遅れていることもあります。このMCP
サーバーは無料のオープンデータのみを使用してそれを修正します。 API キーはありません。アカウントがありません。
npxアグリブレイン
クロード デスクトップのセットアップ (60 秒)
{
"mcpサーバー": {
"agri" : { "command" : " npx " , "args" : [ " agribrain " ] }
}
}
次に、「38.01、23.72 — オリーブ、2024 年 3 月に植えられた現地説明を確認してください。」と尋ねます。
ツール
答えは何ですか
get_field_briefing
「今週は何を心配すればいいですか？」 — すべてのレポート
get_spray_windows
「実際にスプレーできるのはいつですか？」 — 理由付きでウィンドウをランク付け
get_water_balance
「十分に灌水していますか？」 — ET₀ 損失対受領水、純赤字
計算_gdd
「害虫の圧力はどこにありますか？」 — 学位日数、世代、予定日
get_chill_hours
「私の果樹園は冬の寒さを十分に受けましたか？」
get_agro_weather
農業用語の予測と最近の履歴
get_soil_profile
地球上のあらゆる地点の pH、テクスチャー、有機炭素
何が違うのか
これは別の天気ラッパーではありません。このリポジトリ内のすべてのモデル - 害虫
度日のしきい値、スプレーウィンドウのルール、作物のステージング —

厳選されており、
地中海で19年間勤務した認可を受けた農学者によって承認されました
農業、データ ファイル内の文献引用付き。データは安価です。
農学は難しい部分です。
データだけでなく意思決定 – スプレーウィンドウ、水不足、発電タイミング
キーゼロ、コストゼロ — Met.no、NASA POWER、ISRIC SoilGrids (無料、商用可)
引用が含まれています - すべての害虫モデルがその出典とリンクしています
リポジトリ内の評価スイート - これらのツールを使用して、LLM が実際に正しく応答することをテストします。
天気予報: ノルウェー MET (CC-BY 4.0)。
歴史的風土: NASA POWER (パブリックドメイン)。
土壌: ISRIC SoilGrids (CC-BY 4.0)。
ET₀: Hargreaves 経由で計算。 FAO-56 (Allen et al.、1998)。
ローカル トラップ データのない GDD 出力は推定値であり、そのようにラベル付けされます。
SoilGrids の解像度は 250m です。これはデフォルトであり、土壌テストの代替ではありません。
このツールは決定を通知します。地元の農学者に代わるものではありません。
製品ラベル。ここに記載されているものは、申請率に関するアドバイスではありません。
v2 — 現場の目: Sentinel-2 NDVI 時系列とゾーン異常
あらゆる圃場ポリゴンの検出 (無料のコペルニクス データ)、FAO-56 作物水
需要 (Kc × ET₀)、30 年間の気候状況。
v3 — コンプライアンス: EU 農薬承認チェック、収穫前間隔、
抵抗グループ。
Ask Oli のオープン データ レイヤーとして構築および維持されます。
— 地中海の小規模農家を対象とした AI 農学者。パートタイムで維持
ソロ創設者。問題は毎週トリアージされ、農学上の貢献（害虫モデル）
お住まいの地域 — data/pest-models.json スキーマを参照) は特に歓迎です
そして個人的にレビューしました。
AI エージェント用の農業インテリジェンス / 天気、ET₀、GDD、スプレー ウィンドウ、土壌。 API キーはゼロです。 MCPサーバー+コアライブラリ。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
エラーが発生しました

ロード中。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Agronomic intelligence for AI agents / weather, ET₀, GDD, spray windows, soil. Zero API keys. MCP server + core library. - VasileiosTs/agribrain

GitHub - VasileiosTs/agribrain: Agronomic intelligence for AI agents / weather, ET₀, GDD, spray windows, soil. Zero API keys. MCP server + core library. · GitHub
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
VasileiosTs
/
agribrain
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits data data examples examples packages packages .gitignore .gitignore LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Give your AI assistant a licensed agronomist's brain.
[DEMO GIF GOES HERE — 15s: real field coordinates → "What's the spray situation
and water balance for my olives this week?" → real answer with numbers]
Every LLM can write a poem about olive trees. None of them know that your
olive fruit fly third generation started Tuesday, that Saturday's wind makes
spraying pointless, or that your field is 12mm behind on water. This MCP
server fixes that — using only free, open data. No API keys. No accounts.
npx agribrain
Claude Desktop setup (60 seconds)
{
"mcpServers" : {
"agri" : { "command" : " npx " , "args" : [ " agribrain " ] }
}
}
Then ask: "Check the field briefing for 38.01, 23.72 — olives, planted March 2024."
Tool
What it answers
get_field_briefing
"What should I worry about this week?" — the everything report
get_spray_windows
"When can I actually spray?" — ranked windows with reasons
get_water_balance
"Am I irrigating enough?" — ET₀ loss vs. water received, net deficit
compute_gdd
"Where is the pest pressure?" — degree-days, generations, projected dates
get_chill_hours
"Did my orchard get enough winter chill?"
get_agro_weather
Forecast + recent history, in farming terms
get_soil_profile
pH, texture, organic carbon for any point on Earth
What makes this different
This is not another weather wrapper. Every model in this repo — pest
degree-day thresholds, spray-window rules, crop staging — is curated and
signed off by a licensed agronomist with 19 years in Mediterranean
agriculture, with literature citations in the data files. Data is cheap;
agronomy is the hard part.
Decisions, not just data — spray windows, water deficits, generation timing
Zero keys, zero cost — Met.no, NASA POWER, ISRIC SoilGrids (free, commercial-friendly)
Citations included — every pest model links its sources
Eval suite in the repo — we test that LLMs actually answer correctly with these tools
Weather forecasts: MET Norway (CC-BY 4.0).
Historical climate: NASA POWER (public domain).
Soil: ISRIC SoilGrids (CC-BY 4.0).
ET₀: computed via Hargreaves; FAO-56 (Allen et al., 1998).
GDD outputs without local trap data are estimates and labeled as such.
SoilGrids is 250m resolution — a default, not a substitute for a soil test.
This tool informs decisions; it does not replace your local agronomist or
the product label. Nothing here is application-rate advice.
v2 — eyes on the field: Sentinel-2 NDVI time series and zone anomaly
detection for any field polygon (free Copernicus data), FAO-56 crop water
demand (Kc × ET₀), 30-year climate context.
v3 — compliance: EU pesticide approval checks, pre-harvest intervals,
resistance groups.
Built and maintained as the open data layer of Ask Oli
— the AI agronomist for Mediterranean smallholders. Maintained part-time by a
solo founder; issues are triaged weekly, agronomy contributions (pest models
for your region — see data/pest-models.json schema) are especially welcome
and reviewed personally.
Agronomic intelligence for AI agents / weather, ET₀, GDD, spray windows, soil. Zero API keys. MCP server + core library.
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
