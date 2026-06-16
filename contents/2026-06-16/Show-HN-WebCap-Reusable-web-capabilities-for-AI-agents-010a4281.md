---
source: "https://github.com/edgestorage/web-cap"
hn_url: "https://news.ycombinator.com/item?id=48554207"
title: "Show HN: WebCap – Reusable web capabilities for AI agents"
article_title: "GitHub - edgestorage/web-cap: Web-Capability: Script-first web capabilities for AI agents. Run in-page scripts, save workflows as reusable capabilities, and generate AI-native userscripts. · GitHub"
author: "huadream5827"
captured_at: "2026-06-16T13:57:38Z"
capture_tool: "hn-digest"
hn_id: 48554207
score: 3
comments: 0
posted_at: "2026-06-16T12:31:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: WebCap – Reusable web capabilities for AI agents

- HN: [48554207](https://news.ycombinator.com/item?id=48554207)
- Source: [github.com](https://github.com/edgestorage/web-cap)
- Score: 3
- Comments: 0
- Posted: 2026-06-16T12:31:14Z

## Translation

タイトル: Show HN: WebCap – AI エージェント向けの再利用可能な Web 機能
記事のタイトル: GitHub -edgestorage/web-cap: Web 機能: AI エージェント向けのスクリプトファースト Web 機能。ページ内スクリプトを実行し、ワークフローを再利用可能な機能として保存し、AI ネイティブのユーザー スクリプトを生成します。 · GitHub
説明: Web 機能: AI エージェント用のスクリプトファースト Web 機能。ページ内スクリプトを実行し、ワークフローを再利用可能な機能として保存し、AI ネイティブのユーザー スクリプトを生成します。 - エッジストレージ/ウェブキャップ

記事本文:
GitHub -edgestorage/web-cap: Web 機能: AI エージェント用のスクリプトファースト Web 機能。ページ内スクリプトを実行し、ワークフローを再利用可能な機能として保存し、AI ネイティブのユーザー スクリプトを生成します。 · GitHub
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
エッジストレージ
/
ウェブキャップ
公共
通知

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
46 コミット 46 コミット .github/ workflows .github/ workflows extension extension lib lib scripts スクリプト 共有 共有スキル/ Web-cap スキル/ Web-cap テスト テスト タイプ タイプ .gitignore .gitignore .npmrc .npmrc AGENTS.md AGENTS.md ライセンス ライセンス README.md README.md README.zh-CN.md README.zh-CN.md eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml rollup.config.mjs rollup.config.mjs tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts wxt.config.ts wxt.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント向けのスクリプトファースト Web 機能。ページ内スクリプトを実行し、ワークフローを再利用可能な機能として保存し、AI ネイティブのユーザー スクリプトを生成します。
Web-Capability は、エージェント向けのローカルファーストのブラウザ自動化ツールキットです。これにより、エージェントは実際のブラウザーのタブを検査し、再利用可能なページ内スクリプトを実行し、後でコマンドラインで使用できるように成功したワークフローを保存し、自然言語のブラウザーリクエストを AI ネイティブのユーザースクリプトに変換できます。
エージェントは、Web-Cap CLI を介して Web-Capability と対話します。 CLI は必要なローカル ランタイムを自動的に管理するため、ユーザーは別の起動コマンドを必要としません。
スキル CLI を使用して Web Cap スキルをインストールします。
npx スキルでedgestorage/web-capを追加
このスキルには、Web-cap CLI のインストールとエージェントの接続チェックのワークフローが含まれています。
Web Cap ブラウザ拡張機能をインストールします。
Web Cap リリース ページを開きます。
*chrome*.zip のような名前の Chrome 拡張機能 zip アセットをダウンロードします。
ダウンロードした拡張アセットを解凍します。
Chrome で chrome://extensions を開きます。
解凍した拡張機能フォルダーを拡張機能ページにドラッグします。
Web Cap 拡張機能の詳細を開き、「許可」を有効にします。

ユーザースクリプト。
CLI がブラウザー ランタイムを認識できることを確認します。
Hacker News で Web Cap Hub スクリプトを再利用する
web-cap-hub から再利用可能なスクリプトを実行して、現在のページの最初の 5 件の Hacker News 投稿に対するコメントを要約します。これにより、ページの探索が減り、トークンが減り、実行が高速化されます。
YouTube セクションを 1 文で非表示にする
YouTube Gaming のトップ ライブ ゲーム ブロックを 1 つの文で非表示にし、今後のアクセスでも非表示のままにします。
エージェント ワークフローの場合、Web Cap スキルは推奨される CLI セットアップ パスを提供します。 CLI を直接インストールするには、npm を使用します。
npm install -g web-capability
インストールされているコマンドは web-cap です。
web-cap --ヘルプ
Web キャップ セッション ステータス
特長
実際の Chromium ベースのブラウザ タブ用のブラウザ拡張ランタイム。
スクリプトの実行、登録、タブの作成、およびユーザーハンドオフの監視のためのコマンドラインインターフェイス。
検査、待機、クリック、入力、クエリ、テキスト読み取りなどの一般的な操作のための Playwright スタイルのページ ヘルパー。
再利用可能なブラウザワークフローのためのローカルスクリプトレジストリ。
永続的なページ固有のブラウザ変更のための AI ネイティブのユーザースクリプト生成。
エージェント ワークフローのブラウザ タブの作成およびイベント監視コマンド。
デフォルトではローカルファーストの状態ストレージ。
Web Cap は、ローカルの .web-cap/ ディレクトリから再利用可能な機能スクリプトを実行できます。共有 Web Cap Hub リポジトリは、一般的な Web サイトですぐに使用できるスクリプトを収集し、新しいサイト固有のワークフローを作成するためのサンプルを提供します。
ハブからスクリプトを再利用するには:
git clone https://github.com/edgestorage/web-cap-hub.git
cd ウェブキャップハブ
Web キャップ セッション ステータス
Web キャップ スクリプトの実行 \
--タブID < タブID > \
--script-file .web-cap/github.com/read-repository-summary.js \
--input ' {"owner":"edgestorage","repo":"web-cap"} '
現在のスクリプト コレクションと貢献ガイドラインについては、Web Cap Hub README を参照してください。
多くのbr

OWSE 自動化ツールは、直接アクションの固定セットを公開します。つまり、このセレクターをクリックし、その入力を入力し、このテキストを読み、スクリーンショットを撮ります。 Web Cap は代わりにスクリプトファーストのアプローチを採用します。
エージェントは、Playwright スタイルのヘルパーを使用してページ内で JavaScript を実行し、有用なスクリプトを再利用可能なブラウザ スキルとして登録できます。これにより、Web Cap は、エージェントがページ構造を検査し、製品固有の UI に適応し、成功した操作を後で再度実行できるものに変える必要があるワークフローにさらに適しています。
Web Cap は、エージェントが毎回同じブラウザ ワークフローを再発見できるように設計されていません。その中心的な価値は、検証済みのブラウザ操作を再利用可能なスクリプトと再利用可能なワークフローに変えることです。
繰り返しのページやタスクの場合、エージェントはページを繰り返し読んだり、各ステップを計画したり、適切なコントロールを見つけたり、間違いから回復したりする代わりに、安定したワークフローを再利用できます。これにより、トークンの使用量とブラウザーの繰り返しの探索にかかる時間を削減しながら、精度と実行速度を向上させることができます。
この意味で、Web Cap は、Codex、Claude Code、またはその他のローカル エージェント ツールのブラウザー機能レイヤーとして適切に機能します。モデルは目標の理解と意思決定に集中でき、安定したブラウザーの操作はローカルの再利用可能な自動化によって処理されます。
アクションファーストのブラウザ ツールと比較して、Web Cap は次の点に重点を置いています。
ページ内実行。スクリプトは DOM とページの状態を直接操作できます。
再利用可能な機能により、成功したスクリプトを保存して再度実行できます。
ページの検査と対話のための Playwright スタイルのページ ヘルパー。
オプションの実行後観察により、証拠収集が有効になっている場合にスクリプトを実行すると、ページ上で何が変更されたかに関する証拠を返すことができます。
ローカル永続性により、エージェントが学習したワークフローは 1 回の実行後も存続できます。
CLI アクセスにより、エージェントは同じブラウザ キャップを使用できます

通常のコマンドライン ワークフローの機能。
Web Cap は、証拠収集が有効になっている場合、スクリプト実行に関するページを観察できます。スクリプトの実行前に表示要素のスナップショットを取得し、実行中に DOM の変更を追跡し、その後変更された領域のスナップショットを取得して、追加された項目、削除された項目、および更新された項目を含む表示要素の差分を返します。実行証拠には、開いたタブ、URL の変更、リロード、スクロールの変更、管理されたクリック、キーボード入力、スクリプト呼び出しなどのブラウザー側のイベントも含まれる場合があります。
つまり、エージェントはスクリプトで宣言された JSON 結果を取得するだけではありません。また、スクリプトの後にブラウザが目に見えて実行したことを検査することもできます。これは、検証、回復、および新しく成功したスクリプトを再利用可能な機能として登録するかどうかの決定に役立ちます。
ページ ターゲティング: スクリプト定義にはターゲット サイト、URL パターン、ページ ヒント、タグ、タイプ、ステータス、バージョンが含まれるため、エージェントは適切な機能を選択して、間違ったページでスクリプトを実行することを回避できます。
2 つのスクリプト タイプ: read スクリプトはページの状態を検査または抽出し、act スクリプトはページ上で動作するかブラウザ側の変更をトリガーします。
ユーザーハンドオフの観察: wait-events は、ユーザーがブラウザーのアクションを完了するまで待機し、結果の対話パスを JSON ラインとしてストリーミングします。エージェントがユーザーのアクションを必要とするステップに到達し、ユーザーが次に何をしたかを推測するために観察されたクリック、入力/変更/送信アクティビティ、URL の変更、または読み込み状態が必要な場合に使用します。
ローカル実行履歴: インライン スクリプトは、ステータスと結果のメタデータとともにローカルで追跡されます。一時スクリプト ID は、最新のローカル履歴エントリ内にある間は呼び出し可能のままです。
成功ゲート登録: --register は、実行結果に ok: true が含まれる場合にのみスクリプトを永続化します。これは、再利用可能なスクリプト レジストリをクリーンな状態に保つのに役立ちます。
タブ-

認識された実行: コマンドは特定の --tab-id をターゲットにできますが、デフォルトの実行はアクティブな接続されたブラウザー タブに従います。
このロードマップは、Web Cap と Web Cap Hub の計画された開発方向の概要を示しています。
再利用可能なスクリプトの迅速なインストールとダウンロードのサポートを提供します。
Firefox ブラウザ拡張機能のサポートを提供します。
クライアントの構築と配布の改善
Node.js および npm 環境への依存を減らし、よりシンプルなインストール、ビルド、配布パスを検討します。
ブラウザサイド AI チャットとローカル AI ツールの統合
実際の実行のために Codex や Claude Code などのローカル ツールに接続するブラウザ内 AI チャット エントリ ポイントを提供します。
スクリプトのコンパイルをクライアントに移動する
拡張機能のサイズと複雑さを軽減するために、TypeScript のコンパイル関連のより重い責任をブラウザ拡張機能からクライアントに移します。
エージェント
|
| CLIコマンド
v
ウェブキャップ CLI
|
v
マネージドローカルランタイム
|
|ウェブソケット
v
ブラウザ拡張機能
|
v
実際のブラウザタブ
ブラウザ拡張機能はローカル ランタイムに接続し、通常のブラウザ タブに対してコマンドを実行します。エージェントは CLI を呼び出し、CLI はランタイムの起動と接続の詳細を自動的に処理します。
extension/ - ブラウザ拡張機能のエントリポイントとランタイム コード。
lib/ - CLI、ローカル ランタイム、スクリプト レジストリ、およびオーケストレーション ロジック。
共有/ - 共有プロトコル、スクリプト スキーマ、および検証ヘルパー。
skill/ - スキル CLI を使用してインストールできるエージェント スキル。
testing/ - CLI、ランタイム動作、ブラウザ コマンド コントラクト、および拡張ヘルパーに関する Vitest のカバレッジ。
scripts/ - プロジェクト ユーティリティと生成されたランタイム ヘルパー。
現在の拡張機能ランタイム用の Chromium ベースのブラウザー
pnpmインストール
拡張機能の開発ビルドを開始します。
pnpm開発
生成された拡張機能を WXT の出力ディレクトリからロードし、通常の http または https ページを開きます。
ソースを実行する

e 開発中の CLI:
pnpm cli セッションステータス
一般的なエージェントのフローは次のとおりです。
script-execute を使用して、接続されたブラウザに対してスクリプト コードを実行します。
成功したインライン スクリプトを再利用可能にする必要がある場合は、--register を script-execute に追加します。
選択したブラウザタブでスクリプトコードを実行します。スクリプトは 1 つのオブジェクト引数を受け取り、1 つの JSON オブジェクトを返します。
script-execute は、 --timeout-ms 、 --script-file 、 --input-file 、 --no-evidence 、 --register などのオプションの実行設定を受け入れます。実行中に、スクリプトは挿入された Playwright スタイルのページ ヘルパーを使用できます。 --register は、 ok: true で実行が成功した後にのみインライン スクリプトを保存します。
Web Cap には、タブ コントロールが必要なエージェント ワークフロー、またはユーザーがブラウザ ステップを完了して結果のアクション パスを検査するのを待つ必要があるエージェント ワークフロー用の、 browser-new-tab 、 session-status 、 wait-events などのコマンドも含まれています。
スクリプトは、JSON 互換の入力と出力を備えた JavaScript 関数です。
デフォルトの非同期関数をエクスポート (入力) {
const title = ページを待ちます。タイトル （ ） ;
const text = ページを待ちます。ロケータ (入力セレクタ) 。 textContent() ;
戻り値 {
わかりました：本当です、
タイトル、
テキスト、
} ;
}
ランタイムは、スクリプトの実行中に Playwright スタイルのページ ヘルパーを挿入します。

[切り捨てられた]

## Original Extract

Web-Capability: Script-first web capabilities for AI agents. Run in-page scripts, save workflows as reusable capabilities, and generate AI-native userscripts. - edgestorage/web-cap

GitHub - edgestorage/web-cap: Web-Capability: Script-first web capabilities for AI agents. Run in-page scripts, save workflows as reusable capabilities, and generate AI-native userscripts. · GitHub
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
edgestorage
/
web-cap
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
46 Commits 46 Commits .github/ workflows .github/ workflows extension extension lib lib scripts scripts shared shared skills/ web-cap skills/ web-cap tests tests types types .gitignore .gitignore .npmrc .npmrc AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md README.zh-CN.md README.zh-CN.md eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml rollup.config.mjs rollup.config.mjs tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts wxt.config.ts wxt.config.ts View all files Repository files navigation
Script-first web capabilities for AI agents. Run in-page scripts, save workflows as reusable capabilities, and generate AI-native userscripts.
Web-Capability is a local-first browser automation toolkit for agents. It lets agents inspect real browser tabs, run reusable in-page scripts, save successful workflows for later command-line use, and turn natural-language browser requests into AI-native userscripts.
Agents interact with Web-Capability through the web-cap CLI. The CLI manages the required local runtime automatically, so users do not need a separate startup command.
Install the Web Cap skill with the skills CLI:
npx skills add edgestorage/web-cap
The skill includes the web-cap CLI installation and connection-check workflow for agents.
Install the Web Cap browser extension:
Open the Web Cap Releases page.
Download the Chrome extension zip asset, named like *chrome*.zip .
Unzip the downloaded extension asset.
Open chrome://extensions in Chrome.
Drag the unzipped extension folder into the extensions page.
Open the Web Cap extension details and enable Allow User Scripts .
Check that the CLI can see the browser runtime:
Reuse a Web Cap Hub script on Hacker News
Run a reusable script from web-cap-hub to summarize the comments on the first five Hacker News posts from the current page with less page exploration, fewer tokens, and faster execution.
Hide a YouTube section with one sentence
Hide the Top live games block on YouTube Gaming with one sentence, and keep it hidden on future visits.
For agent workflows, the Web Cap skill provides the recommended CLI setup path. To install the CLI directly, use npm:
npm install -g web-capability
The installed command is web-cap :
web-cap --help
web-cap session-status
Features
Browser extension runtime for real Chromium-based browser tabs.
Command-line interface for script execution, registration, tab creation, and user handoff observation.
Playwright-style page helpers for common operations such as inspect, wait, click, fill, query, and text reading.
Local script registry for reusable browser workflows.
AI-native userscript generation for persistent, page-specific browser changes.
Browser tab creation and event watching commands for agent workflows.
Local-first state storage by default.
Web Cap can run reusable capability scripts from a local .web-cap/ directory. The shared Web Cap Hub repository collects ready-to-use scripts for common websites and provides examples for writing new site-specific workflows.
To reuse scripts from the hub:
git clone https://github.com/edgestorage/web-cap-hub.git
cd web-cap-hub
web-cap session-status
web-cap script-execute \
--tab-id < tab-id > \
--script-file .web-cap/github.com/read-repository-summary.js \
--input ' {"owner":"edgestorage","repo":"web-cap"} '
See the Web Cap Hub README for the current script collection and contribution guidelines.
Many browser automation tools expose a fixed set of direct actions: click this selector, fill that input, read this text, take a screenshot. Web Cap takes a script-first approach instead.
Agents can run JavaScript inside the page with Playwright-style helpers and register useful scripts as reusable browser skills. This makes Web Cap better suited for workflows where an agent needs to inspect page structure, adapt to product-specific UI, and turn a successful operation into something it can run again later.
Web Cap is not designed to make agents rediscover the same browser workflow every time. Its core value is turning verified browser operations into reusable scripts and reusable workflows.
For recurring pages and tasks, agents can reuse stable workflows instead of repeatedly reading the page, planning each step, finding the right controls, and recovering from mistakes. This can improve accuracy and execution speed while reducing token usage and time spent on repeated browser exploration.
In this sense, Web Cap works well as a browser capability layer for Codex, Claude Code, or other local agent tools: the model can focus on understanding goals and making decisions, while stable browser operations are handled by local reusable automation.
Compared with action-first browser tools, Web Cap focuses on:
In-page execution, so scripts can work directly with the DOM and page state.
Reusable capabilities, so successful scripts can be saved and run again.
Playwright-style page helpers for page inspection and interaction.
Optional post-execution observation, so script runs can return evidence about what changed on the page when evidence collection is enabled.
Local persistence, so agent-learned workflows can survive beyond a single run.
CLI access, so agents can use the same browser capabilities from normal command-line workflows.
Web Cap can observe the page around script execution when evidence collection is enabled. It snapshots visible elements before a script runs, tracks DOM mutations while it runs, then snapshots changed areas afterward and returns a visible-elements diff with added , removed , and updated items. Execution evidence can also include browser-side events such as opened tabs, URL changes, reloads, scroll changes, managed clicks, keyboard input, and script calls.
That means an agent does not only get a script's declared JSON result. It can also inspect what the browser visibly did after the script, which is useful for verification, recovery, and deciding whether a newly successful script should be registered as a reusable capability.
Page targeting: script definitions include target sites, URL patterns, page hints, tags, type, status, and version, so agents can select the right capability and avoid running a script on the wrong page.
Two script types: read scripts inspect or extract page state, while act scripts operate on the page or trigger browser-side changes.
User handoff observation: wait-events waits while a user completes a browser action, then streams the resulting interaction path as JSON Lines. Use it when an agent has reached a step that requires user action and needs the observed clicks, input/change/submit activity, URL changes, or loading state to infer what the user did next.
Local execution history: inline scripts are tracked locally with status and result metadata. Temporary script ids remain callable while they are in the latest local history entries.
Success-gated registration: --register only persists a script when its execution result includes ok: true , which helps keep the reusable script registry clean.
Tab-aware execution: commands can target a specific --tab-id , while default execution follows the active connected browser tab.
This roadmap outlines the planned development directions for Web Cap and Web Cap Hub.
Provide quick installation and download support for reusable scripts.
Provide Firefox browser extension support.
Client Build and Distribution Improvements
Reduce dependency on the Node.js and npm environment, and explore simpler installation, build, and distribution paths.
Browser-Side AI Chat and Local AI Tool Integration
Provide an in-browser AI chat entry point that connects to local tools such as Codex and Claude Code for actual execution.
Move Script Compilation to the Client
Move heavier TypeScript compilation-related responsibilities from the browser extension to the client to reduce extension size and complexity.
Agent
|
| CLI command
v
Web Cap CLI
|
v
Managed local runtime
|
| WebSocket
v
Browser extension
|
v
Real browser tab
The browser extension connects to the local runtime and executes commands against normal browser tabs. Agents call the CLI, and the CLI handles runtime startup and connection details automatically.
extension/ - browser extension entrypoints and runtime code.
lib/ - CLI, local runtime, script registry, and orchestration logic.
shared/ - shared protocol, script schema, and validation helpers.
skills/ - Agent Skills installable with the skills CLI.
tests/ - Vitest coverage for CLI, runtime behavior, browser command contracts, and extension helpers.
scripts/ - project utilities and generated-runtime helpers.
A Chromium-based browser for the current extension runtime
pnpm install
Start the extension development build:
pnpm dev
Load the generated extension from WXT's output directory, then open a normal http or https page.
Run the source CLI during development:
pnpm cli session-status
A typical agent flow is:
Use script-execute to run script code against the connected browser.
Add --register to script-execute when a successful inline script should become reusable.
Execute script code in the selected browser tab. Scripts receive one object argument and return one JSON object.
script-execute accepts optional execution settings such as --timeout-ms , --script-file , --input-file , --no-evidence , and --register . During execution, scripts can use the injected Playwright-style page helper. --register saves the inline script only after execution succeeds with ok: true .
Web Cap also includes commands such as browser-new-tab , session-status , and wait-events for agent workflows that need tab control, or need to wait for a user to complete a browser step and inspect the resulting action path.
Scripts are JavaScript functions with JSON-compatible inputs and outputs:
export default async function ( input ) {
const title = await page . title ( ) ;
const text = await page . locator ( input . selector ) . textContent ( ) ;
return {
ok : true ,
title ,
text ,
} ;
}
The runtime injects a Playwright-style page helper while the script execute

[truncated]
