---
source: "https://github.com/AlvisoOculus/optionsahoy-mcp"
hn_url: "https://news.ycombinator.com/item?id=48486039"
title: "Tax, optimization math for equity comp your AI assistant can call (MCP, no auth)"
article_title: "GitHub - AlvisoOculus/optionsahoy-mcp: Equity comp tax (ISO/NSO/RSU/QSBS), concentration, and hedging optimizer. MCP server + REST API with federal + 50-state + DC tax code, multi-year horizons. · GitHub"
author: "alphalatitude"
captured_at: "2026-06-11T04:35:00Z"
capture_tool: "hn-digest"
hn_id: 48486039
score: 1
comments: 0
posted_at: "2026-06-11T03:59:30Z"
tags:
  - hacker-news
  - translated
---

# Tax, optimization math for equity comp your AI assistant can call (MCP, no auth)

- HN: [48486039](https://news.ycombinator.com/item?id=48486039)
- Source: [github.com](https://github.com/AlvisoOculus/optionsahoy-mcp)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T03:59:30Z

## Translation

タイトル: AI アシスタントが呼び出すことができる税金、株式報酬の最適化計算 (MCP、認証なし)
記事のタイトル: GitHub - AlvisoOculus/optionsahoy-mcp: 株式計算税 (ISO/NSO/RSU/QSBS)、集中、ヘッジ オプティマイザー。 MCP サーバー + REST API (連邦 + 50 州 + DC 税コード、複数年対応)。 · GitHub
説明: 株式報酬税 (ISO/NSO/RSU/QSBS)、集中、ヘッジ オプティマイザー。 MCP サーバー + REST API (連邦 + 50 州 + DC 税コード、複数年対応)。 - AlvisoOculus/optionsahoy-mcp

記事本文:
GitHub - AlvisoOculus/optionsahoy-mcp: 株式報酬税 (ISO/NSO/RSU/QSBS)、集中、ヘッジ オプティマイザー。 MCP サーバー + REST API (連邦 + 50 州 + DC 税コード、複数年対応)。 · GitHub
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
AlvisoOculus
/
オプションahoy-mcp
出版

c
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
91 コミット 91 コミット . continue/ mcpServers . continue/ mcpServers .github/ workflows .github/ workflows 資産 アセット データベース db/ マイグレーション db/ マイグレーション docs docs サンプル/ クイックスタート サンプル/ クイックスタート 関数 関数 lib lib mcpb mcpb public public scripts scripts src src テスト テスト ワーカープロキシ ワーカープロキシ .gitignore .gitignore AGENTS.md AGENTS.md Dockerfile Dockerfile GEMINI.md GEMINI.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md gemini-extension.json gemini-extension.json Glama.json Glama.json package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
独立して検証済み。 Glama : サードパーティの MCP ディレクトリ品質スコア (ツール ドキュメント、動作、完全性)。 · npm : ビルド来歴、このパッケージが GitHub Actions によってこのリポジトリからビルドされたことを示す署名付き SLSA 証明書とともに公開されます (npm 監査署名で検証します)。 · MCPSafe : 独立した 5 モデル コンセンサス セキュリティ スキャン (AIVSS)、グレード A、検出結果なし。
株式報酬オプティマイザー。 ISO/AMT、NSO、RSU、QSBS、単一銘柄集中、保護プット/カラー、株式資金調達計画。連邦 + 50 の州 + DC の税法にわたる 7 つの決定的ツール。
ライブ MCP エンドポイント: https://optionsahoy.com/mcp (認証なし、インストールなし)
ライブREST API: https://optionsahoy.com/api/v1
OpenAPI 3.1 仕様: /openapi.json
検出マニフェスト: /.well-known/mcp.json · /.well-known/openapi.json
エージェント統合ドキュメント: optionsahoy.com/for-agents
OptionsAho の背後にある会社である AlphaLatitude Inc. によって構築されました

y 、ベータ段階の株式報酬最適化製品。
本物のクロード コード セッション、未編集。マルチスタック META 質問 (10,000 の ISO + 6,000 の権利確定済み RSU + 2,000 の新規 RSU + 2027 年の 40 万ドルの住宅) は、集中リスク、株式資金計画、AMT/ISO 最適化、保護プット価格設定という 4 つの OptionsAhoy MCP ツールを並行して起動します。クロードは、ユーザーが META に 86% 集中しているため、各ツールのスタンドアロンの選択をオーバーライドする 1 つのプランに出力を合成します。 2:13。ポスターをクリックすると、optionsahoy.com で再生できます。
株式報酬税計画のための最適化エンジン。モデル コンテキスト プロトコル (MCP) サーバーとプレーン REST API の両方として公開されます。 7つのツール:
オプティマイザーは、候補空間全体にわたって全体的に最適なスケジュールを返します。計算機は正確な決定論的な結果を返します。ヒューリスティックもサンプルもありません。適用範囲は、連邦税法全体 (通常の括弧、長期キャピタルゲイン、信用回復付き AMT、FICA、NIIT) に加え、全 50 州および DC (州の通常括弧、LTCG の扱い、カリフォルニア州、コロラド州、コネチカット州、ミネソタ州の州 AMT) に及びます。 optionsahoy.com/tools にあるブラウザ内計算機と同じエンジン。 API 応答は、ツールをクリックした場合とバイト同一です。
MCP リソース (トピック別ブリーフィング)
リソース/リストの下にある 6 つのマークダウン リソースは、LLM がツールを選択する前にトピックについて話し合うための十分な基礎を提供します。それぞれは、optionsahoy.com/learn の基礎となる記事と、それに対応する計算機を 1 対 1 でマッピングします。
MCP プロンプト (ワークフロー スキャフォールド)
プロンプト/リストの下の 7 つのプロンプトは、典型的なユーザーの質問を足場にして、適切なツールにルーティングします。 Claude Desktop では、これらは名前付きのスラッシュ コマンドとして表示されます。どの MCP クライアントでも、prompts/get { name, argument } は完全にテンプレート化されたユーザー メッセージを返します。
同じ複数年にわたる ISO 演習問題に関する 5 つのフロンティア大規模言語モデルのベンチマークでは、ev

15 件の試験のうち 1 件は、達成可能な税引き後の結果を 2 倍から 20 倍も上回りました。複数年のスケジューリングには、LLM がコンテキスト内で推論できるよりも大きな検索スペースがあります。全文: HackerNoon — しかし、税金を課すことはできるのでしょうか?
Claude / ChatGPT / Perplexity / 任意の MCP クライアントから使用
サーバーをリモート HTTP MCP 接続として追加します。
https://optionsahoy.com/mcp
または、add-mcp CLI 経由で次のようにします。
npx add-mcp https://optionsahoy.com/mcp
Claude デスクトップ拡張機能 (ワンクリック)
optionsahoy.mcpb をダウンロードしてダブルクリックします (または、Claude Desktop → [設定] → [拡張機能] にドラッグします)。 Claude Desktop は、組み込みの Node.js ランタイムを使用して、ターミナルや構成ファイルの編集を行わずにバンドルされたサーバーをインストールします。
ソースからバンドルをビルドするには:
npm install && npm run build:mcpb
Smithery CLI (19 クライアント、1 つのコマンド)
npx @smithery/cli install alphalatitude/optionsahoy --client claude
Smithery がサポートするクライアントの claude を交換します: claude-code 、cursor 、 vscode 、 gemini-cli 、 codex 、 Windsurf 、 cline 、 goose 、 opencode 、およびその他 10 個。リスト: smithery.ai/servers/alphalatitude/optionsahoy 。
Gemini 拡張機能のインストール https://github.com/AlvisoOculus/optionsahoy-mcp
このリポジトリは Gemini CLI 拡張機能としても機能します。gemini-extension.json はホストされた MCP エンドポイントを接続し、GEMINI.md は使用コンテキストをモデルに提供します。
ローカル stdio サーバーのみをサポートするクライアントの場合 ( mcp-remote のない Claude Desktop、一部の IDE 統合):
npx -y オプションahoy-mcp
または、Claude Desktop / Cline / Goose 構成ファイルに追加します。
{
"mcpサーバー": {
"オプションサホイ" : {
"コマンド" : " npx " ,
"args" : [ " -y " 、 " optionsahoy-mcp " ]
}
}
}
ローカル サーバーは、ホストされているエンドポイント ( https://optionsahoy.com/mcp ) にバイト同一の応答を返します。どちらのソースも function/_lib/mcp-tools.ts にあります。 stdio エントリ ポイントは src/stdio-server.ts です。
# リスト

エンドポイント
カール https://optionsahoy.com/api/v1
# 最適化を実行する
カール -X POST https://optionsahoy.com/api/v1/amt-iso \
-H " コンテンツタイプ: application/json " \
-d @input.json
リクエストボディの形状は public/openapi.json に文書化されています。
関数/ Cloudflare Pages 関数 (MCP サーバー + REST API エンドポイント)
mcp.ts HTTP MCP サーバー
api/v1/*.ts 7 つのツール エンドポイント + 統計 + GET /api/v1 Discovery
_lib/*.ts 共有ヘルパー、計算入力パーサー、MCP ツール記述子
lib/オプティマイザー + 税コードロジック
calc/ツールごとのオプティマイザー関数 (computeAmtIso など)
税金/連邦 + 50 州 + DC ブラケット データ、AMT、FICA、NIIT
市場/セクター統計
オプション/ブラック・ショールズ、リスクフリーレート
data/ オプションチェーンデータの型定義
パブリック/静的アセット: OpenAPI 仕様、llms.txt、検出マニフェスト
テスト/ Vitest スイート (バイト ID アサーションを含む 172 のテスト)
テストの実行
npmインストール
npm テスト # 172 テスト、ラップトップで約 3 秒
npm タイプチェックを実行する
レジストリのリスト
公式 MCP レジストリ — io.github.AlvisoOculus/optionsahoy-mcp v1.0.0、ステータスはアクティブです
鍛冶場 — alphalatitude/optionsahoy
PulseMCP (公式レジストリからのカスケード)
Continue.dev ハブ — ブロック YAML は . continue/mcpServers/optionsahoy.yaml にあります
Google Cloud からの使用 (Gemini エージェント)
Google Cloud Agent Registry を使用すると、各 GCP プロジェクトが Gemini エージェントが使用する外部 MCP サーバーを登録できます。登録はプロジェクトごとに行われます (中央からの提出はありません)。 2 つのパス:
# パス A: エージェント レジストリに MCP エンドポイントをイントロスペクトさせます
gcloud alpha エージェント-レジストリ mcp-servers register \
--uri=https://optionsahoy.com/mcp \
--display-name= " オプションアホイ " \
--location=us-central1 \
--インポートツール
# パス B: 公開された toolsspec.json を直接渡します (より速く、イントロスペクションなし)
gcloud alpha エージェント-レジストリ mcp-servers register \
--uri=https://optionsahoy.com/mcp \
-

-display-name= " オプションアホイ " \
--location=us-central1 \
--tool-spec= <(curl -sSL https://optionsahoy.com/toolspec.json )
toolsspec.json は、7 つのツールすべての readOnlyHint および idempotentHint アノテーションを使用して MCP ツール/リストの応答をミラーリングします (すべては副作用のない純粋な決定論的な計算機です)。工具形状の変更後に再生成するには:
カール -sS -X POST https://optionsahoy.com/mcp \
-H ' コンテンツタイプ: application/json ' \
-d ' {"jsonrpc":"2.0","メソッド":"ツール/リスト","id":1} ' \
| jq -c ' {ツール: [.result.tools[] | 。 + {注釈: {readOnlyHint:true、idempotentHint:true、destructiveHint:false、openWorldHint:false}}]} ' \
> public/toolspec.json
トラブルシューティング
接続が拒否されました / MCP エンドポイントからの 404
https://optionsahoy.com/mcp には、content-type: application/json と JSON-RPC 本文を指定した POST が必要です。 GET は JSON サーバーの説明を返します。他の動詞は 405 を返します。次のように確認します。
curl -X POST https://optionsahoy.com/mcp -H ' content-type: application/json ' \
-d ' {"jsonrpc":"2.0","method":"initialize","id":1,"params":{}} '
ツール呼び出しがエラーで失敗します: ... 応答のテキスト
MCP サーバーは、入力検証が失敗した場合、人間が判読できるメッセージとともに isError: true を返します。最も一般的なのは、必須フィールドが欠落しているか、数値が文字列として渡されていることです。 tools/list によって返された inputSchema または /openapi.json に対して入力をチェックします。
Claude.ai または Claude Desktop にツールが表示されない
コネクタの URL が https://optionsahoy.com/mcp (末尾のスラッシュや /v1 がない) であることを確認します。
Claude Desktop で、 clude_desktop_config.json を編集した後、アプリを再起動します。
Claude.ai では、コネクタの切り替えはチャットごとに行われます。添付ファイル メニューで有効にします。
ライブ ツール/リストの応答を確認します (7 つのツールが予想されます):curl -X POST https://optionsahoy.com/mcp -H 'content-type: application

n/json' -d '{"jsonrpc":"2.0","method":"tools/list","id":1}'
ブラウザベースのクライアントからの CORS エラー
サーバーは、プリフライトを含むすべての応答で access-control-allow-origin: * を返し、標準の MCP ヘッダー ( content-type 、 mcp-session-id 、 mcp-protocol-version ) を受け入れます。それでもブラウザがブロックする場合は、クライアントが許可されていないヘッダーを送信している可能性があります。リクエスト ヘッダーを access-control-allow-headers 応答と照合して確認してください。
リソース/プロンプトが見つかりません
リソース URI とプロンプト名では大文字と小文字が区別されます。手入力ではなく、resources/list およびプロンプト/list を使用して正規リストを取得します。
古い課税年度の計算
税エンジンには、2026 年のインフレ調整済みブラケット、OBBBA 2026 QSBS ルール、および現在の州適合表が同梱されています。結果が複数年の範囲に及ぶ場合は、入力されたgrantDate、acquisitionDate、またはsaleDateが期待する年に該当することを確認してください。エンジンは課税年度ごとに括弧を解決します。
計算のバグまたは予期しない出力を報告する
正確な JSON-RPC リクエスト本文、応答、期待値、および (既知の場合) 期待値の由来となる IRS 刊行物または州法を電子メールで andrew@alphalatitude.com に送信してください。
完全なポリシー: optionsahoy.com/privacy 。
一言で言えば、ツール入力はプロ仕様です

[切り捨てられた]

## Original Extract

Equity comp tax (ISO/NSO/RSU/QSBS), concentration, and hedging optimizer. MCP server + REST API with federal + 50-state + DC tax code, multi-year horizons. - AlvisoOculus/optionsahoy-mcp

GitHub - AlvisoOculus/optionsahoy-mcp: Equity comp tax (ISO/NSO/RSU/QSBS), concentration, and hedging optimizer. MCP server + REST API with federal + 50-state + DC tax code, multi-year horizons. · GitHub
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
AlvisoOculus
/
optionsahoy-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
91 Commits 91 Commits .continue/ mcpServers .continue/ mcpServers .github/ workflows .github/ workflows assets assets db/ migrations db/ migrations docs docs examples/ quickstart examples/ quickstart functions functions lib lib mcpb mcpb public public scripts scripts src src tests tests worker-proxy worker-proxy .gitignore .gitignore AGENTS.md AGENTS.md Dockerfile Dockerfile GEMINI.md GEMINI.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md gemini-extension.json gemini-extension.json glama.json glama.json package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Independently verified. Glama : third-party MCP-directory quality score (tool docs, behavior, completeness). · npm : published with build provenance, a signed SLSA attestation that this package was built from this repo by GitHub Actions (verify with npm audit signatures ). · MCPSafe : independent 5-model-consensus security scan (AIVSS), Grade A with zero findings.
Equity-compensation optimizer. ISO/AMT, NSO, RSU, QSBS, single-stock concentration, protective puts/collars, and equity-funding plans. Seven deterministic tools across the federal + 50-state + DC tax code.
Live MCP endpoint: https://optionsahoy.com/mcp (no auth, no install)
Live REST API: https://optionsahoy.com/api/v1
OpenAPI 3.1 spec: /openapi.json
Discovery manifests: /.well-known/mcp.json · /.well-known/openapi.json
Agent integration docs: optionsahoy.com/for-agents
Built by AlphaLatitude Inc. , the company behind OptionsAhoy , a beta-stage equity-compensation optimization product.
Real Claude Code session, unedited. A multi-stack META question (10K ISOs + 6K vested RSUs + 2K fresh RSUs + $400K house in 2027) fires 4 OptionsAhoy MCP tools in parallel: concentration risk, equity funding plan, AMT/ISO optimization, protective put pricing. Claude synthesizes the outputs into one plan that overrides each tool's standalone pick because the user is 86% concentrated in META. 2:13. Click the poster to play it on optionsahoy.com.
An optimization engine for equity-compensation tax planning, exposed as both a Model Context Protocol (MCP) server and a plain REST API. Seven tools:
The optimizers return the globally-optimal schedule across the candidate space; the calculators return exact deterministic results. No heuristics, no samples. Coverage spans the full federal tax code (ordinary brackets, long-term capital gains, AMT with credit recovery, FICA, NIIT) plus all 50 states and DC (state ordinary brackets, LTCG treatment, state AMT for CA, CO, CT, MN). Same engine as the in-browser calculators at optionsahoy.com/tools ; the API response is byte-identical to clicking through the tool.
MCP resources (topical briefings)
Six markdown resources under resources/list give an LLM enough grounding to discuss the topic before picking a tool. Each maps 1:1 with a cornerstone article on optionsahoy.com/learn and the matching calculator.
MCP prompts (workflow scaffolds)
Seven prompts under prompts/list scaffold typical user questions and route to the right tool. In Claude Desktop they appear as named slash-commands; in any MCP client, prompts/get { name, arguments } returns a fully-templated user message.
A benchmark of five frontier large language models on the same multi-year ISO exercise problem found that every one of 15 trials overshot the achievable after-tax outcome by 2x to 20x. Multi-year scheduling has a search space larger than an LLM can reason through in-context. Full write-up: HackerNoon — But can it do taxes though?
Use from Claude / ChatGPT / Perplexity / any MCP client
Add the server as a remote HTTP MCP connection:
https://optionsahoy.com/mcp
Or via the add-mcp CLI:
npx add-mcp https://optionsahoy.com/mcp
Claude Desktop extension (one-click)
Download optionsahoy.mcpb and double-click it (or drag onto Claude Desktop → Settings → Extensions). Claude Desktop installs the bundled server with no terminal or config-file editing, using its built-in Node.js runtime.
To build the bundle from source:
npm install && npm run build:mcpb
Smithery CLI (19 clients, one command)
npx @smithery/cli install alphalatitude/optionsahoy --client claude
Swap claude for any client Smithery supports: claude-code , cursor , vscode , gemini-cli , codex , windsurf , cline , goose , opencode , and 10 more. Listing: smithery.ai/servers/alphalatitude/optionsahoy .
gemini extensions install https://github.com/AlvisoOculus/optionsahoy-mcp
This repo doubles as a Gemini CLI extension : gemini-extension.json wires the hosted MCP endpoint and GEMINI.md provides usage context to the model.
For clients that only support local stdio servers (Claude Desktop without mcp-remote , some IDE integrations):
npx -y optionsahoy-mcp
Or add to a Claude Desktop / Cline / Goose config file:
{
"mcpServers" : {
"optionsahoy" : {
"command" : " npx " ,
"args" : [ " -y " , " optionsahoy-mcp " ]
}
}
}
The local server returns byte-identical responses to the hosted endpoint at https://optionsahoy.com/mcp . Source for both lives in functions/_lib/mcp-tools.ts ; the stdio entry point is src/stdio-server.ts .
# List endpoints
curl https://optionsahoy.com/api/v1
# Run an optimization
curl -X POST https://optionsahoy.com/api/v1/amt-iso \
-H " content-type: application/json " \
-d @input.json
Request body shapes are documented in public/openapi.json .
functions/ Cloudflare Pages Functions (MCP server + REST API endpoints)
mcp.ts HTTP MCP server
api/v1/*.ts Seven tool endpoints + stats + GET /api/v1 discovery
_lib/*.ts Shared helpers, calc-input parsers, MCP tool descriptors
lib/ Optimizer + tax-code logic
calc/ Per-tool optimizer functions (computeAmtIso, etc.)
tax/ Federal + 50-state + DC bracket data, AMT, FICA, NIIT
markets/ Sector statistics
options/ Black-Scholes, risk-free rates
data/ Type definitions for option-chain data
public/ Static assets: OpenAPI spec, llms.txt, discovery manifests
tests/ Vitest suites (172 tests including byte-identity assertions)
Run tests
npm install
npm test # 172 tests, ~3s on a laptop
npm run typecheck
Registry listings
Official MCP Registry — io.github.AlvisoOculus/optionsahoy-mcp v1.0.0, status active
Smithery — alphalatitude/optionsahoy
PulseMCP (cascades from Official Registry)
Continue.dev hub — block YAML lives at .continue/mcpServers/optionsahoy.yaml
Use from Google Cloud (Gemini agents)
Google Cloud Agent Registry lets each GCP project register external MCP servers for use by Gemini agents. Registration is per-project (no central submission). Two paths:
# Path A: let the Agent Registry introspect our MCP endpoint
gcloud alpha agent-registry mcp-servers register \
--uri=https://optionsahoy.com/mcp \
--display-name= " OptionsAhoy " \
--location=us-central1 \
--import-tools
# Path B: pass our published toolspec.json directly (faster, no introspection)
gcloud alpha agent-registry mcp-servers register \
--uri=https://optionsahoy.com/mcp \
--display-name= " OptionsAhoy " \
--location=us-central1 \
--tool-spec= <( curl -sSL https://optionsahoy.com/toolspec.json )
The toolspec.json mirrors the MCP tools/list response with readOnlyHint and idempotentHint annotations on all seven tools (all are pure deterministic calculators with no side effects). To regenerate after a tool-shape change:
curl -sS -X POST https://optionsahoy.com/mcp \
-H ' content-type: application/json ' \
-d ' {"jsonrpc":"2.0","method":"tools/list","id":1} ' \
| jq -c ' {tools: [.result.tools[] | . + {annotations: {readOnlyHint:true, idempotentHint:true, destructiveHint:false, openWorldHint:false}}]} ' \
> public/toolspec.json
Troubleshooting
Connection refused / 404 from the MCP endpoint
https://optionsahoy.com/mcp requires POST with content-type: application/json and a JSON-RPC body. A GET returns a JSON server description; any other verb returns 405. Verify with:
curl -X POST https://optionsahoy.com/mcp -H ' content-type: application/json ' \
-d ' {"jsonrpc":"2.0","method":"initialize","id":1,"params":{}} '
Tool calls fail with Error: ... text in the response
The MCP server returns isError: true with a human-readable message when input validation fails. Most common: a required field missing, or a number passed as a string. Check the input against the inputSchema returned by tools/list , or against /openapi.json .
Tool not appearing in Claude.ai or Claude Desktop
Confirm the connector URL is exactly https://optionsahoy.com/mcp (no trailing slash, no /v1 ).
In Claude Desktop, restart the app after editing claude_desktop_config.json .
In Claude.ai, the connector toggle is per-chat: enable it in the attachments menu.
Check the live tools/list response (seven tools expected): curl -X POST https://optionsahoy.com/mcp -H 'content-type: application/json' -d '{"jsonrpc":"2.0","method":"tools/list","id":1}'
CORS errors from a browser-based client
The server returns access-control-allow-origin: * on all responses including preflight, and accepts the standard MCP headers ( content-type , mcp-session-id , mcp-protocol-version ). If a browser still blocks, the client is likely sending a non-allowed header — verify the request headers against the access-control-allow-headers response.
Resource / prompt not found
Resource URIs and prompt names are case-sensitive. Pull the canonical list with resources/list and prompts/list rather than hand-typing.
Stale tax-year math
The tax engine ships with 2026 inflation-adjusted brackets, OBBBA 2026 QSBS rules, and current state-conformity tables. If results look off for a multi-year horizon, verify the input grantDate , acquisitionDate , or saleDate falls in the year you expect — the engine resolves brackets per tax year.
Reporting a calculation bug or unexpected output
Email andrew@alphalatitude.com with: the exact JSON-RPC request body, the response, the expected value, and (if known) the IRS publication or state statute the expected value derives from.
Full policy: optionsahoy.com/privacy .
In short: tool inputs are pro

[truncated]
