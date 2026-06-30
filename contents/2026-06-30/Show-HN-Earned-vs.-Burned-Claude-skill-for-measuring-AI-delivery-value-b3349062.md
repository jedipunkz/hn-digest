---
source: "https://github.com/harveer10x/earned-vs-burned-skill"
hn_url: "https://news.ycombinator.com/item?id=48726771"
title: "Show HN:Earned vs. Burned, Claude skill for measuring AI delivery value"
article_title: "GitHub - harveer10x/earned-vs-burned-skill: A Claude skill for measuring AI delivery value. Replaces tokens, story points & velocity with one honest metric: outcomes earned vs effort burned. · GitHub"
author: "harveer10x"
captured_at: "2026-06-30T00:31:20Z"
capture_tool: "hn-digest"
hn_id: 48726771
score: 1
comments: 0
posted_at: "2026-06-29T23:42:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN:Earned vs. Burned, Claude skill for measuring AI delivery value

- HN: [48726771](https://news.ycombinator.com/item?id=48726771)
- Source: [github.com](https://github.com/harveer10x/earned-vs-burned-skill)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T23:42:51Z

## Translation

タイトル: AI 配信価値を測定するための HN: Earned vs. Burned、Claude スキルを表示
記事タイトル: GitHub - harveer10x/earned-vs-burned-skill: AI 提供価値を測定するためのクロード スキル。トークン、ストーリー ポイント、ベロシティを、獲得した成果と消費した努力という 1 つの正直な指標に置き換えます。 · GitHub
説明: AI 配信価値を測定するためのクロード スキル。トークン、ストーリー ポイント、ベロシティを、獲得した成果と消費した努力という 1 つの正直な指標に置き換えます。 - harveer10x/獲得スキルと燃焼スキル

記事本文:
GitHub - harveer10x/earned-vs-burned-skill: AI 提供価値を測定するための Claude スキル。トークン、ストーリー ポイント、ベロシティを、獲得した成果と消費した努力という 1 つの正直な指標に置き換えます。 · GitHub
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
ハーバー10x
/
獲得したスキルと消費したスキル
出版

ic
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
harveer10x/獲得スキル vs 燃焼スキル
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット アセット アセット リファレンス リファレンス ライセンス ライセンス README.md README.md SKILL.md SKILL.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
獲得 vs 燃焼 — クロード スキル
AI ネイティブの配信において実際に何が重要かを測定するためのオープンソースの Claude スキル。
ハービア・シン著 |著者、データが移動するとき |リズ・ワイヤレス創設者
AI を使用して構築されているすべてのチームは、間違ったものを測定しています。
トークンの使用法。コード行。ストーリーポイント。速度。何時間も燃えた。これらは燃焼指標であり、価値ではなく労力を測定します。 100 トークンで幻覚を起こすモデルが、10,000 トークンで問題を解決するモデルよりも優れているわけではありません。出荷されない 500 行のコードは何も得られません。本番環境に到達しない作業でクローズされた 40 のストーリー ポイントでは、ビジネス価値はゼロになります。
重要な唯一の尺度は、それを獲得できたかどうかです。
価値は、具体的で検証可能なビジネス成果が達成された場合にのみ得られます。その時点までは、すべての努力は無駄になるだけです。
これは、Earned vs Burned フレームワークです。元々は、100 万以上の SKU を処理する Sysco Foods データ ファクトリー用に、2010 年に Deloitte で構築されました (この用語が誕生する前からあった人間と AI のハイブリッド ワークフロー)。 AI ネイティブのソフトウェア配信向けに一般化されました。
レベル
ゲート
獲得した
それが何を意味するか
0
未開始
0.00
バックログ。燃焼ゼロ、収益ゼロ。
1
進行中
0.00
努力の積み重ね。まだ収入ゼロです。
2
開発完了
0.25
コードが作成され、単体テストが行われた。部分的。
3
QAに合格しました
0.60
テストされ、受け入れられました。意味はありますが、本番環境ではありません。
4
本番環境へのデプロイメント
0.85
本番環境で実行中。閉じる — しかし確認されていません。
5
検証された結果
1.00
KPI が移動しました。ユーザーが確認しました。収益に影響が出ました。の

完全に稼ぐだけです。
主要な指標
メトリック
計算式
ターゲット
獲得率 %
L5 の結果 ÷ 総ストーリー数
70%以上
獲得/時間
獲得した合計 ÷ 合計時間
>0.10
AIトークンごとに獲得
獲得合計 ÷ トークン合計
上昇傾向
結果の検証%
L5 検証済み ÷ L4+L5 導入済み
上昇傾向
AI トークンあたりの収益は、業界がまだ持っていない指標です。トークンのボリュームを完全に置き換えます。
このスキルをクロードにインストールすると、次のことが可能になります。
任意のプロジェクト ツール (Linear、Asana、GitHub Issues、Jira、Azure DevOps、または貼り付けられた/CSV リスト) からタスクを取得します。
5 レベルの結果階層に基づいて各タスクをスコアリングします
すべての指標を計算します — 獲得率、E/H 比率、トークンごとの獲得率、結果検証 %、チームの有効性
アーンド バリュー レポートを生成する — 1 ページ、3 つの数字がベロシティ レポートの代わりとなります。
チームを指導します — すべてのレポートの最後に、配信文化を変える質問、つまり、この作業が完了したことを確認する検証可能な結果は何ですか? という質問を付けます。
FTE チーム、アウトソーシング/オフショア配信、AI エージェントのワークフロー、またはあらゆる組み合わせで機能します。
オプション 1 — .skill ファイルをインストールする (Claude Desktop / Cowork)
「リリース」ページから「earned-vs-burned.skill」をダウンロードします。
Claude Desktop または Cowork の場合: [設定] → [機能] → [スキルのインストール]
オプション 2 — 手動インストール (クロード コード)
# リポジトリのクローンを作成する
git clone https://github.com/[your-org]/earned-vs-burned-skill
# スキルディレクトリでクロードをポイントします
# .claude/settings.json に追加します。
# "skillPaths": ["./earned-vs-burned"]
使用例
インストールしたら、クロードに自然に話しかけるだけです。
「Earned vs Burned フレームワークに対してリニア スプリントをスコアリングする」
「これが私の Jira タスクです。このスプリントで実際にどれくらいの価値を獲得できましたか?」
「今月は 400 時間と 200 万トークンを消費しました。私たちは何を稼いだでしょうか?」
「私はアウトソーシング ベンダーです。結果を使ってクライアントに当社の価値を証明できるように手伝ってください」

メトリクス」
「GitHub の問題を取得して、アーンド バリュー レポートを送ってください。」
獲得 vs 燃焼/
§── SKILL.md # クロードのコアスキル説明書
§── README.md # このファイル
§── 参考文献/
│ §── Framework.md # 完全なフレームワーク、起源のストーリー、階層の詳細
│ §── metrics.md # 各メトリクスの「だから何」解釈ガイド
│ └── integrations.md # Linear、Asana、GitHub、Jira、ADO からプルする方法
└── 資産/
└── tracker_template.csv # 入力してアップロードする空のトラッカー
貢献する
これはオープンなフレームワークです。貢献を歓迎します:
新しい統合 — Monday.com、Shortcut、Notion、ClickUp
スプリントの例 — フレームワークの動作を示す実際の匿名化されたデータ
翻訳 — 他言語のフレームワーク
ロケールのバリエーション — 特定の業界 (医療、金融、政府) への適応
PR または問題を開きます。フレームワークは Harveer Singh の IP です。スキルの実装には MIT ライセンスが適用されます。
フレームワーク IP: © Harveer Singh. Earned vs Burned フレームワーク、結果階層、および関連する概念は、Harveer Singh の独自の知的財産です (2010 年にシスコ/デロイトによって設立、2026 年に正式化)。
スキルコード：MITライセンス。実装を自由に使用、フォーク、適応、構築してください。出典を感謝します。
Harveer Singh は、元フォーチュン 500 最高データ責任者 (信託銀行、ウェスタン ユニオン)、Rizz Wireless の創設者、および『When Data Moves』の著者です。
彼は、「AI ネイティブ」という言葉が誰もが使うようになる前の 2010 年に、デロイトでオリジナルの Earned vs Burned モデルを構築し、データとテクノロジ プログラムはアクティビティではなく成果を提供するために存在することを 25 年かけて証明してきました。
リンクトイン |ニュースレター: データが移動するとき |本
AI の提供価値を測定するためのクロード スキル。トークン、ストーリー ポイント、ベロシティを 1 つの正直な指標、つまり結果に置き換えます。

獲得した労力と消費した労力。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
獲得スキル vs 燃焼スキル v1.0
最新の
2026 年 6 月 29 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A Claude skill for measuring AI delivery value. Replaces tokens, story points & velocity with one honest metric: outcomes earned vs effort burned. - harveer10x/earned-vs-burned-skill

GitHub - harveer10x/earned-vs-burned-skill: A Claude skill for measuring AI delivery value. Replaces tokens, story points & velocity with one honest metric: outcomes earned vs effort burned. · GitHub
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
harveer10x
/
earned-vs-burned-skill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
harveer10x/earned-vs-burned-skill
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits assets assets references references LICENSE LICENSE README.md README.md SKILL.md SKILL.md View all files Repository files navigation
Earned vs Burned — Claude Skill
An open source Claude skill for measuring what actually matters in AI-native delivery.
By Harveer Singh | Author, When Data Moves | Founder, Rizz Wireless
Every team building with AI is measuring the wrong things.
Token usage. Lines of code. Story points. Velocity. Hours burned. These are burned metrics — they measure effort, not value. A model that hallucinates in 100 tokens is not better than one that solves the problem in 10,000. 500 lines of code that don't ship earn nothing. 40 story points closed on work that never reaches production delivers zero business value.
The only measure that matters: did we earn it?
Value is only EARNED when a tangible, verifiable business outcome is achieved. Until that point, all effort is only BURNED .
This is the Earned vs Burned Framework — originally built at Deloitte in 2010 for a Sysco Foods data factory processing 1M+ SKUs (a human-AI hybrid workflow that predated the term). Now generalised for AI-native software delivery.
Level
Gate
Earned
What it means
0
Not Started
0.00
Backlog. Zero burn, zero earn.
1
In Progress
0.00
Effort accumulating. Still zero earned.
2
Dev Complete
0.25
Code written and unit-tested. Partial.
3
QA Passed
0.60
Tested and accepted. Meaningful but not in prod.
4
Deployed to Prod
0.85
Running in production. Close — but not confirmed.
5
Outcome Verified
1.00
KPI moved. User confirmed. Revenue impacted. The only full earn.
The Key Metrics
Metric
Formula
Target
Earn Rate %
L5 outcomes ÷ Total stories
70%+
Earned / Hours
Total Earned ÷ Total Hours
>0.10
Earned per AI Token
Total Earned ÷ Total Tokens
Trending up
Outcome Verification %
L5 verified ÷ L4+L5 deployed
Trending up
Earned per AI Token is the metric the industry doesn't have yet. It replaces token volume entirely.
Install this skill into Claude and it can:
Pull tasks from any project tool — Linear, Asana, GitHub Issues, Jira, Azure DevOps, or a pasted/CSV list
Score each task against the 5-level Outcome Hierarchy
Calculate all metrics — Earn Rate, E/H Ratio, Earned per Token, Outcome Verification %, Team Effectiveness
Generate an Earned Value Report — one page, three numbers, replaces your velocity report
Coach your team — ends every report with the question that changes delivery culture: what is the verifiable outcome that will confirm this work is done?
Works for FTE teams, outsourced/offshored delivery, AI-agent workflows, or any mix.
Option 1 — Install the .skill file (Claude Desktop / Cowork)
Download earned-vs-burned.skill from the Releases page
In Claude Desktop or Cowork: Settings → Capabilities → Install Skill
Option 2 — Manual install (Claude Code)
# Clone the repo
git clone https://github.com/[your-org]/earned-vs-burned-skill
# Point Claude at the skill directory
# Add to your .claude/settings.json:
# "skillPaths": ["./earned-vs-burned"]
Usage Examples
Once installed, just talk to Claude naturally:
"Score my Linear sprint against the Earned vs Burned framework"
"Here are my Jira tasks — how much value did we actually earn this sprint?"
"We burned 400 hours and 2M tokens this month. What did we earn?"
"I'm an outsourcing vendor — help me prove our value to the client using outcome metrics"
"Pull our GitHub issues and give me an Earned Value Report"
earned-vs-burned/
├── SKILL.md # Core skill instructions for Claude
├── README.md # This file
├── references/
│ ├── framework.md # Full framework, origin story, hierarchy detail
│ ├── metrics.md # "So what" interpretation guide for each metric
│ └── integrations.md # How to pull from Linear, Asana, GitHub, Jira, ADO
└── assets/
└── tracker_template.csv # Blank tracker to fill in and upload
Contributing
This is an open framework. Contributions welcome:
New integrations — Monday.com, Shortcut, Notion, ClickUp
Example sprints — Real anonymised data showing the framework in action
Translations — The framework in other languages
Locale variants — Adaptations for specific industries (healthcare, finance, gov)
Open a PR or issue. The framework is IP of Harveer Singh; the skill implementation is MIT licensed.
Framework IP: © Harveer Singh. The Earned vs Burned Framework, Outcome Hierarchy, and associated concepts are original intellectual property of Harveer Singh (established 2010, Sysco/Deloitte; formalised 2026).
Skill code: MIT License. Use, fork, adapt, and build on the implementation freely. Attribution appreciated.
Harveer Singh is a former Fortune 500 Chief Data Officer (Truist Bank, Western Union), founder of Rizz Wireless , and author of When Data Moves .
He built the original Earned vs Burned model at Deloitte in 2010 — before "AI-native" was a phrase anyone used — and has spent 25 years proving that data and technology programs exist to deliver outcomes, not activity.
LinkedIn | Newsletter: When Data Moves | Book
A Claude skill for measuring AI delivery value. Replaces tokens, story points & velocity with one honest metric: outcomes earned vs effort burned.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
Earned vs Burned Skill v1.0
Latest
Jun 29, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
