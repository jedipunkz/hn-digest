---
source: "https://github.com/stockvaluation-io/stockvaluation_io"
hn_url: "https://news.ycombinator.com/item?id=48662309"
title: "Show HN: Empowering codex/Claude Code with Aswath Damodaran valuation thinking"
article_title: "GitHub - stockvaluation-io/stockvaluation_io: Open-source, local-first MCP DCF valuation workflow that turns a stock ticker into transparent Damodaran-style assumptions, scenarios, and valuation outputs you can inspect, challenge, and refine. · GitHub"
author: "pradeep1177"
captured_at: "2026-06-24T16:42:07Z"
capture_tool: "hn-digest"
hn_id: 48662309
score: 1
comments: 0
posted_at: "2026-06-24T16:31:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Empowering codex/Claude Code with Aswath Damodaran valuation thinking

- HN: [48662309](https://news.ycombinator.com/item?id=48662309)
- Source: [github.com](https://github.com/stockvaluation-io/stockvaluation_io)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T16:31:10Z

## Translation

タイトル: HN を表示: アスワス ダモダラン評価思考によるコーデックス/クロード コードの強化
記事のタイトル: GitHub - Stockvaluation-io/stockvaluation_io: オープンソースのローカルファースト MCP DCF 評価ワークフロー。株式ティッカーを、検査、挑戦、改良できる透明な Damodaran スタイルの仮定、シナリオ、評価出力に変換します。 · GitHub
説明: オープンソースのローカルファースト MCP DCF 評価ワークフローは、株式相場を検査、検証、調整できる透明なダモダラン スタイルの仮定、シナリオ、評価出力に変換します。 - ストックバリュエーション-io/ストックバリュエーション_io

記事本文:
GitHub - Stockvaluation-io/stockvaluation_io: オープンソースのローカルファースト MCP DCF 評価ワークフロー。株式ティッカーを、検査、挑戦、改良できる透明な Damodaran スタイルの仮定、シナリオ、評価出力に変換します。 · GitHub
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
{{

メッセージ }}
在庫評価-io
/
在庫評価_io
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
在庫評価-io/在庫評価_io
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
102 コミット 102 コミット .github/ workflows .github/ workflows docker/ local/ postgres docker/ local/ postgres docs/ media docs/ media scripts scripts valuation-agent valuation-agent valuation-service valuation-service yfinance yfinance .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTEXT.md CONTEXT.md ライセンス ライセンス README.md README.md docker-compose.local.yml docker-compose.local.yml install.sh install.sh pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
StockValuation.io は、Codex と Claude にローカル評価ワークフローを提供します。
エージェントは企業を調査し、証拠を収集し、評価に関する質問をし、教育レポートを作成できます。ローカル サービスは DCF 計算を実行し、監査可能な数値を返します。
教育用途のみ。これは経済的なアドバイスではありません。
このデモでは、SpaceX のローカル StockValuation.io 目論見書ワークフローを使用した Codex CLI を示します。 Codex は、抽出された出願証拠をレビューし、ガイド付きの評価に関する質問をし、教育的な評価ビューを作成します。
これを構築したのは、評価規律を尊重する AI ワークフローが欲しかったからです。
ダモダラン スタイルの評価では、ビジネス ストーリーから仮定、キャッシュ フローに至るまでの連鎖よりも、最終的な価値が重要ではありません。成長が上がれば、モデルは収益の道筋を示すはずです。利益率が拡大する場合は、レポートでその理由を説明する必要があります。再投資が減少した場合、ユーザーはその背後にある資本効率の主張に気づくはずです。
エージェントは、提出書類の読み取り、比較を支援します。

情報源を調べ、ビジネスを要約し、ストーリーのプレッシャーテストを行います。隠された数字をでっち上げたり、公正価値を手動で計算したりしてはなりません。
StockValuation.io では、これらのジョブを分離しています。調査・説明はエージェントが行います。ローカル ツールは評価計算を処理します。仮定を検証し、どのシナリオが信頼に値するかを判断します。
DCF 評価は、紙の上では単純に見えます。難しい部分は入力にあります。
小さなインプットの変化が評価を大きく動かす可能性があります。ツールがこれらの仮定を隠すと、出力が証拠、判断、またはモデルのショートカットからのものであるかどうかがわかりません。
StockValuation.io は仮定を可視化します。準備ができているふりをするのではなく、材料ドライバーにガイド付きの質問をし、ローカル サービスを通じてシナリオを再計算し、弱い評価ケースをマークします。
ワークフローによる責任の分担方法
ローカル評価ツールは以下を処理します。
エージェントは、評価出力のためにローカル ツールを呼び出す必要があります。評価額を手動で計算すべきではありません。
Codex と Claude の Stockvaluation.io スキル。
評価ワークフロー用のローカル MCP ツール。
評価ランタイム用の Docker サービス。
決定論的な DCF 計算とシナリオの再計算。
ベースライン値、成長アンカー、参照データのステータス、および有効な仮定。
証拠の検討のために一時停止する調査フロー。
最終レポートの前に、ガイド付きの評価に関する質問をします。
データが評価をサポートできない場合の失敗メッセージ。
ビジネスストーリーを数字に結びつける。
評価シナリオの比較。
Codex または Claude を使用して検査可能なローカル ワークフローを実行します。
エージェントネイティブの評価スタックを構築してテストします。
推奨事項を購入、販売、または保留します。
パーソナライズされた投資決定。
このプロジェクトは、あなたの目標、リスク許容度、ポートフォリオ、財務状況を知りません。
このプロジェクトは、ダモダランの結び方の習慣に従っています。

ストーリーを数字に。
企業が急速に成長できるというストーリーであれば、収益の伸び、利益率の進捗、再投資の必要性、リスクを数字で示す必要があります。数字が不可能な話を暗示している場合、エージェントはその仮定に異議を唱える必要があります。
出力は検査できる引数です。市場価格が間違っているという主張ではありません。
アスワス・ダモダランはこのプロジェクトを支持していませんし、私は彼と何の関係もありません。
デフォルトのフローでは質問が使用されます。そのパスを要求しない限り、ワンショット レポートは生成されません。
エージェントは、ローカル評価ツールが実行されていることを確認します。
エージェントはローカル サービスから決定的なベースラインを取得します。
エージェントは企業を調査し、主な評価要因に関する証拠を収集します。
エージェントは一時停止して、証拠を確認できるようにします。
エージェントは、ガイド付きの評価に関する質問をします。
ローカル サービスはシナリオを再計算します。
エージェントは最終的な教育レポートを作成します。
証拠と質問のループをスキップしたい場合にのみ、クイック実行を依頼してください。
Docker Desktop または Compose と互換性のある Docker エンジンが必要です。
./install.sh セットアップ
または、GitHub からインストーラーを実行します。
カール -fsSL https://raw.githubusercontent.com/stockvaluation-io/stockvaluation_io/main/install.sh | bash -s -- セットアップ
セットアップでは、スキルのインストールまたは更新、ローカル ツールの構成、Docker サービスの開始、およびサービス ステータスの出力が行われます。インストーラーは、デフォルトで Codex と Claude をターゲットにします。
デフォルトでは、curl インストーラーはリポジトリのクローンを ~/.local/share/stockvaluation_io に作成します。別の場所を選択するには、STOCKVALUATION_INSTALL_DIR=/path/to/dir を設定します。
./install.sh ステータス
./install.sh 開始
./install.sh 停止
./install.sh アンインストール
インストーラーは、 docker-compose.local.yml を通じてローカル評価スタックを実行します。
セットアップ後、エージェントに評価を依頼してください。
Stockvaluation.io を使用して MSFT を評価します。
Google の価値

Stockvaluation.ioを使用します。
Stockvaluation.ioを使用してMETAを評価します。
デフォルトの調査フローでは、エージェントが最初に証拠を示し、ガイド付きの仮定に基づいた質問をし、回答後にレポートを作成することを想定しています。
SEC 目論見書の提出から評価額を求めることもできます。
Stockvaluation.io を使用して、この SEC 目論見書から企業を評価します: <SEC EDGAR HTML URL>
このワークフローでは、申告事実を抽出し、証拠を確認するよう求め、その後、評価を決定する前提条件を尋ねます。ファイリングに十分なサポートが不足している場合、ツールは洗練された数値ではなく、明確な警告または失敗を返す必要があります。
ローカルファースト、完全オフラインではない
評価サービスと DCF 計算はマシン上で実行されます。
一部の入力は依然としてマシンの外部から来ます。
エージェントが使用するモデルプロバイダー
このリポジトリには、完全にローカルな LLM スタックは同梱されていません。
公的提出書類と市場データプロバイダーがデータ範囲を管理します。
ワークフローは、プライマリ ファイリング データが欠落しているかサポートされていない場合に、正規化されたフォールバック データを使用することがあります。
上流のデータが欠落している、古い、低品質、またはモデルに不適切な場合、評価は失敗する可能性があります。
米国以外、ADR、IFRS、および特殊な申請事件では、追加の情報源レビューが必要になる場合があります。
このサービスは金融セクターの企業をサポートしていません。
このサービスは、サポートされていない企業に対しては明らかな失敗を返すはずです。
成長アンカーと参照データは批評をサポートします。それらは価値を証明するものではありません。
DCF 出力は仮定に基づいて変化します。
ガイド付き質問のデフォルトはモデリングのデフォルトです。これらは投資を推奨するものではありません。
明らかな失敗は偽りの評価に勝ります。
ローカル サービスを保護する方法がわからない場合は、ローカル サービスをマシン上に残しておきます。
すべてのレポートを教材として扱います。
購入、売却、保有、目標価格、および個人向けの推奨文言は避けてください。
シナリオを信頼する前に、すべての重要な前提条件を確認してください。
@その他{st

ockvaluation_io、
著者 = {プラディープ・シン}、
title = {StockValuation.io: Codex および Claude 用のローカル株式評価ツール},
年 = {2026}、
パブリッシャー = {GitHub}、
URL = {https://github.com/stockvaluation-io/stockvaluation_io}
}
謝辞
中心となる方法論と参照データは、アスワス ダモダランの公開評価リソースを利用しています。
https://pages.stern.nyu.edu/~adamodar/New_Home_Page/data.html
このプロジェクトは Aswath Damodaran と提携または承認されていません。教育用途のみ。
オープンソースのローカルファーストの MCP DCF 評価ワークフローは、株価表示を、検査、挑戦、改良が可能な透明な Damodaran スタイルの仮定、シナリオ、評価出力に変換します。
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
11
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open-source, local-first MCP DCF valuation workflow that turns a stock ticker into transparent Damodaran-style assumptions, scenarios, and valuation outputs you can inspect, challenge, and refine. - stockvaluation-io/stockvaluation_io

GitHub - stockvaluation-io/stockvaluation_io: Open-source, local-first MCP DCF valuation workflow that turns a stock ticker into transparent Damodaran-style assumptions, scenarios, and valuation outputs you can inspect, challenge, and refine. · GitHub
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
stockvaluation-io
/
stockvaluation_io
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
stockvaluation-io/stockvaluation_io
main Branches Tags Go to file Code Open more actions menu Folders and files
102 Commits 102 Commits .github/ workflows .github/ workflows docker/ local/ postgres docker/ local/ postgres docs/ media docs/ media scripts scripts valuation-agent valuation-agent valuation-service valuation-service yfinance yfinance .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTEXT.md CONTEXT.md LICENSE LICENSE README.md README.md docker-compose.local.yml docker-compose.local.yml install.sh install.sh pyproject.toml pyproject.toml View all files Repository files navigation
StockValuation.io gives Codex and Claude a local valuation workflow.
Your agent can research a company, gather evidence, ask valuation questions, and write an educational report. The local service runs the DCF math and returns auditable numbers.
Educational use only. This is not financial advice.
The demo shows Codex CLI using the local StockValuation.io prospectus workflow for SpaceX. Codex reviews extracted filing evidence, asks guided valuation questions, and produces an educational valuation view.
I built this because I wanted an AI workflow that respects valuation discipline.
In a Damodaran-style valuation, the final value matters less than the chain from business story to assumptions to cash flows. If growth goes up, the model should show the revenue path. If margins expand, the report should explain why. If reinvestment falls, the user should see the capital-efficiency claim behind it.
An agent can help with reading filings, comparing sources, summarizing a business, and pressure-testing a story. It should not invent hidden numbers or hand-calculate a fair value.
StockValuation.io keeps those jobs separate. The agent handles research and explanation. The local tools handle valuation math. You inspect the assumptions and decide which scenario deserves trust.
DCF valuation looks simple on paper. The hard part sits in the inputs:
Small input changes can move the valuation a lot. When a tool hides those assumptions, you cannot tell whether the output came from evidence, judgment, or a model shortcut.
StockValuation.io makes the assumptions visible. It asks guided questions for the material drivers, recalculates scenarios through the local service, and marks weak valuation cases instead of pretending they are ready.
How the workflow splits responsibility
The local valuation tools handle:
The agent should call the local tools for valuation output. It should not hand-calculate valuation numbers.
A stockvaluation.io skill for Codex and Claude.
Local MCP tools for valuation workflows.
Docker services for the valuation runtime.
Deterministic DCF math and scenario recalculation.
Baseline values, growth anchors, reference-data status, and effective assumptions.
A researched flow that pauses for evidence review.
Guided valuation questions before the final report.
Failure messages when the data cannot support a valuation.
Connecting a business story to numbers.
Comparing valuation scenarios.
Running an inspectable local workflow with Codex or Claude.
Building and testing an agent-native valuation stack.
Buy, sell, or hold recommendations.
Personalized investment decisions.
This project does not know your goals, risk tolerance, portfolio, or financial situation.
This project follows the Damodaran practice of tying story to numbers.
If the story says a company can grow fast, the numbers need to show revenue growth, margin progress, reinvestment needs, and risk. If the numbers imply an impossible story, the agent should challenge the assumptions.
The output is an argument you can inspect. It is not a claim that the market price is wrong.
Aswath Damodaran does not endorse this project, and I have no affiliation with him.
The default flow uses questions. It does not produce a one-shot report unless you ask for that path.
The agent checks that the local valuation tools are running.
The agent gets a deterministic baseline from the local service.
The agent researches the company and gathers evidence for the main valuation drivers.
The agent pauses so you can review the evidence.
The agent asks guided valuation questions.
The local service recalculates scenarios.
The agent writes the final educational report.
Ask for a quick run only when you want to skip the evidence and question loop.
You need Docker Desktop or a compatible Docker Engine with Compose.
./install.sh setup
Or run the installer from GitHub:
curl -fsSL https://raw.githubusercontent.com/stockvaluation-io/stockvaluation_io/main/install.sh | bash -s -- setup
Setup installs or updates the skill, configures the local tools, starts the Docker services, and prints service status. The installer targets Codex and Claude by default.
The curl installer clones the repo to ~/.local/share/stockvaluation_io by default. Set STOCKVALUATION_INSTALL_DIR=/path/to/dir to choose another location.
./install.sh status
./install.sh start
./install.sh stop
./install.sh uninstall
The installer runs the local valuation stack through docker-compose.local.yml .
After setup, ask your agent for a valuation:
Value MSFT using stockvaluation.io.
Value GOOGL using stockvaluation.io.
Value META using stockvaluation.io.
For the default researched flow, expect the agent to show evidence first, ask guided assumption questions, and write the report after you answer.
You can also ask for a valuation from an SEC prospectus filing:
Use stockvaluation.io to value a company from this SEC prospectus: <SEC EDGAR HTML URL>
The workflow extracts filing facts, asks you to review the evidence, and then asks for the assumptions that drive the valuation. If the filing lacks enough support, the tool should return a clear warning or failure instead of a polished number.
Local-first, not fully offline
The valuation services and DCF math run on your machine.
Some inputs still come from outside your machine:
the model provider used by your agent
This repo does not ship a fully local LLM stack.
Public filings and market-data providers control data coverage.
The workflow may use normalized fallback data when primary filing data is missing or unsupported.
Valuation can fail when upstream data is missing, stale, low quality, or unsuitable for the model.
Non-US, ADR, IFRS, and unusual filing cases may need extra source review.
The service does not support financial-sector companies.
The service should return a clear failure for unsupported companies.
Growth anchors and reference data support critique. They do not prove the value.
DCF outputs move with assumptions.
Guided-question defaults are modeling defaults. They are not investment recommendations.
A clear failure beats a fake valuation.
Keep local services on your machine unless you know how to secure them.
Treat every report as educational material.
Avoid buy, sell, hold, target-price, and personalized recommendation language.
Review every material assumption before trusting a scenario.
@misc{stockvaluation_io,
author = {Pradeep Singh},
title = {StockValuation.io: Local stock valuation tools for Codex and Claude},
year = {2026},
publisher = {GitHub},
url = {https://github.com/stockvaluation-io/stockvaluation_io}
}
Acknowledgments
Core methodology and reference data draw on Aswath Damodaran's public valuation resources:
https://pages.stern.nyu.edu/~adamodar/New_Home_Page/data.html
This project is not affiliated with or endorsed by Aswath Damodaran. Educational use only.
Open-source, local-first MCP DCF valuation workflow that turns a stock ticker into transparent Damodaran-style assumptions, scenarios, and valuation outputs you can inspect, challenge, and refine.
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
11
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
