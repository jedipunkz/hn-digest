---
source: "https://github.com/erickong/penecho"
hn_url: "https://news.ycombinator.com/item?id=48955040"
title: "PenEcho: An Open-Source Canvas with AI"
article_title: "GitHub - erickong/penecho: Think with AI beyond the chat box. A shared canvas for handwriting, equations, diagrams, and spatial reasoning. · GitHub"
author: "mkw5053"
captured_at: "2026-07-18T04:39:56Z"
capture_tool: "hn-digest"
hn_id: 48955040
score: 2
comments: 1
posted_at: "2026-07-18T03:51:37Z"
tags:
  - hacker-news
  - translated
---

# PenEcho: An Open-Source Canvas with AI

- HN: [48955040](https://news.ycombinator.com/item?id=48955040)
- Source: [github.com](https://github.com/erickong/penecho)
- Score: 2
- Comments: 1
- Posted: 2026-07-18T03:51:37Z

## Translation

タイトル: PenEcho: AI を備えたオープンソース キャンバス
記事タイトル: GitHub - erickong/penecho: チャット ボックスを超えて AI で考える。手書き、方程式、図、空間推論のための共有キャンバス。 · GitHub
説明: チャット ボックスを超えて AI を使って考えます。手書き、方程式、図、空間推論のための共有キャンバス。 - エリッコン/ペネチョ

記事本文:
GitHub - erickong/penecho: チャット ボックスを超えて AI で考える。手書き、方程式、図、空間推論のための共有キャンバス。 · GitHub
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
エリクコン
/
ペネチョ
公共
通知
通知設定を変更するにはサインインする必要があります

s
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット .github .github docs docs public public test test .gitattributes .gitattributes .gitignore .gitignore COMMERCIAL-LICENSE.md COMMERCIAL-LICENSE.md COTRIBUTING.md COTRIBUTING.md COTRIBUTOR-LICENSE-AGREEMENT.md Contributor-LICENSE-AGREEMENT.md ライセンス ライセンス通知 NOTICE README.md README.md TRADEMARKS.md TRADEMARKS.md api-config.js api-config.js claude-cli.js claude-cli.js cli.js cli.js codex-cli.js codex-cli.js configure-ui.js configure-ui.js package-lock.json package-lock.json package.json package.json server.js server.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
チャット ボックスを超えて AI を使って考えましょう。
PenEcho は、手書き、方程式、図、空間コンテキストが会話の一部となる共有キャンバスです。
質問、方程式、図、または途中で形成されたアイデアをキャンバス上の任意の場所に置き、一時停止します。 PenEcho はあなたのマークとその空間的関係を読み取り、その横で答えます。すべてのステップをチャット メッセージに変換したり、厳密な図表ツールを使用して再構築したりしなくても、問題に対処できます。
答え、ヒント、説明、続き、公式、プロット、図をキャンバス上で直接入手できます。
すべての AI ドラフトが作業の一部になる前に、移動、サイズ変更、承認、または破棄を行ってください。
スタイラスまたはマウスで自然に描画し、まばらな 20,000 x 20,000 のキャンバス上でパンやズームを行います。
確認されたインクの周りにフリーハンドのなげなわを描いて、ローカルで移動、サイズ変更、または色変更します。選択を受け入れたりキャンセルしたりしても、AI リクエストはトリガーされません。
調査している問題の種類に合わせて、難解、SF、またはリサーチ モードを選択します。
軽量のスナップショットをブラウザにローカルに保存します。新しいキャンバスを開始すると、現在のスナップショットが上書きされ、

新しいコピー、または保存せずに続行します。未確認の AI ドラフトは決して含まれません。
PenEcho はローカル ランタイムを小さくし、インクが存在する場所に 512 x 512 タイルのみを割り当てるため、巨大な論理キャンバスが巨大なビットマップになることはありません。
フローチャート LR
User["手書き、方程式、スケッチ"] --> Canvas["ブラウザキャンバス<br/>確認済みのまばらなタイル"]
キャンバス --> アトラス[「トリミングされたビジュアル アトラス<br/>プラス ジオメトリ」]
Atlas --> サーバー["PenEcho サーバー<br/>検証とプロンプト"]
サーバー --> エグゼキュータ{"構成されたエグゼキュータ"}
Executor --> API["API モード<br/>OpenAI 互換または Anthropic"]
実行者 --> Codex["Codex CLI モード<br/>ローカル codex exec"]
実行者 --> クロード["クロード CLI モード<br/>ローカル クロード -p"]
API --> ドラフト["構造化された編集可能なドラフト"]
コーデックス --> ドラフト
クロード --> ドラフト
ドラフト --> キャンバス
読み込み中
ブラウザは、関連するキャンバスのトリミングとジオメトリのみを送信します。サーバーはリクエストを検証し、選択されたエグゼキューターを使用して、ユーザーがそれを受け入れるまで確認されたインクとは分離された可動ドラフトを返します。
Node.js 18.17 以降と、API キー、認証された Codex CLI 、または認証された Claude Code CLI のいずれかが必要です。
npm install -g ペネチョ
ペネチョ設定
ペネチョ
penecho configure は対話型構成センターを開きます。メイン メニューには、LLM ソース、設定、および終了が含まれています。矢印キーと Enter キーを使用して移動します。
LLM ソース -> Claude CLI は、検出されたモデル、推奨モデル、デフォルトのモデル、または手動で入力されたモデルとエフォート レベルを選択します。 Opus 4.8 以降を推奨します。 Sonnet と Opus 4.6 は応答できますが、キャンバスの結果が弱くなる可能性があります。
LLM ソース -> Codex CLI はモデルとエフォートを選択します。良好な結果を得るには GPT-5.5 以降が必要で、gpt-5.6-sol が推奨され、xhigh は Codex の取り組みの中で最も上位にリストされています。
LLM ソース -> API は OpenAI 互換または Anthropic/Claude-co を選択します

互換性のあるリクエスト形式を使用して、URL、モデル、エフォート、および隠しキーを要求します。既存の値がデフォルトとして提供され、空のキーは保存されたキーを保持します。
設定では、統合モデルのタイムアウト、すべてのモデル エグゼキューターに送信されるイメージ形式、リクエストの記録と保持、リスニング インターフェイスとポート、および初期自動 AI 遅延を制御します。 WebP がデフォルトです。 PNGも利用可能です。遅延はキャンバス上で変更することもできます。
すべての LLM ページは Test & Save で終わり、PenEcho はチェックする前に常に保存します。 Codex CLI は高速オフライン チェックを使用します。実行可能ファイルとログインを検証してから、Codex デバッグ モデル --bundled を読み取り、選択したモデルが存在することを確認します。推論の実行、画像の添付、オンライン カタログの更新、モデル トークンの消費は行われません。 Claude CLI および API 構成は、選択されたエンドポイント/モデル設定を確認するために 1 つの小さな実際のリクエストを送信します。チェックが成功しても失敗しても、構成は保存されたままになり、UI は明確な診断を表示して親メニューに戻ります。
デフォルトの構成は ~/.penecho/config.env です。 API 資格情報は、このローカル ファイルではプレーンテキストであり、POSIX システムでは所有者のみのアクセス許可を受け取り、ブラウザ コードに送信されることはありません。他の認証情報と同様に保護してください。このファイルが存在する前に penecho が起動された場合、対話型ターミナルで構成センターが自動的に開きます。
必要に応じて、特定の起動に別の env スタイルの構成ファイルを使用します。
penecho configure --config ./team.env
ペネチョ --config ./team.env
明示的な --config ファイルは、そのコマンドのデフォルトのグローバル ファイルを置き換えます。 PenEcho は、プロジェクト ディレクトリ .env またはパッケージ ディレクトリ .env を自動的に読み取りません。
Codex デスクトップ アプリを単独でインストールしても、Codex 実行可能ファイルがシェル PATH で利用できるとは限りません。インストールして認証する

Codex を選択する前に、CLI を個別に作成します。
npm install -g @openai/codex@latest
ハッシュ -r
コーデックス --バージョン
コーデックスのログインステータス
必要に応じて、 codex login を実行します。 Claude CLI モードでも同様に、インストールおよび認証された Claude Code CLI (通常は claude auth login を通じて) が必要です。
PenEcho は選択された CLI をローカルで使用するため、そのソースの API キーは必要ありません。通常の起動では、モデル トークンを消費せずに実行可能ファイルとログインがチェックされます。 Codex Test & Save は、モデル要求を行わずに、選択したモデルをインストールされている CLI のバンドル カタログに対してさらに検証します。 Claude Test & Save は小さな実際のリクエストを送信します。
Codex を介した Canvas リクエストは codex exec --json を使用します。 Codex が最終エージェント メッセージとturn.completed を発行するとすぐに PenEcho が返されます。 CLI プロセスがその後も存続している場合、キャンバスの応答を遅らせる代わりに、CLI プロセスは終了され、バックグラウンドでクリーンアップされます。
Claude CLI リクエストは、ツール、エージェント、MCP、プロンプト提案、セッション永続性、およびその他の重要でないバックグラウンド トラフィックを無効にして、1 つの分離された claude -p ターンを使用します。 PenEcho は、これらのレイテンシーに敏感なキャンバス ターンに対して MAX_THINKING_TOKENS=0 を設定するため、クロード コードは隠れた拡張思考の予算を大量に消費しません。エフォートが選択されると、PenEcho はクロードの --effort フラグとプロセスごとの設定オーバーライドの両方を適用するため、ユーザーレベルの CLAUDE_CODE_EFFORT_LEVEL が暗黙的にエフォートを置き換えることはできません。 PenEcho はストリームを段階的に検証し、Claude が成功した最終結果を出力するとすぐに戻ります。ツールを使用しようとするとリクエストは中止されますが、結果が終了した後も存続する CLI プロセスはバックグラウンドで終了およびクリーンアップされます。
一時的な起動オーバーライドは引き続き利用可能です。
ペネチョドクター --コーデックス
penecho --codex --model gpt-5.6-sol --effort xhigh
ペネチョ --クロード --モデル作品 --efo

最大RT
ペネチョ -- ポート 4000
--model 、 --effort 、および --port はそのプロセスにのみ適用され、選択した構成ファイルよりも優先されます。保存された選択または基礎となる CLI デフォルトを使用する場合は、それらを省略します。他のモデル固有のエフォート文字列は受け入れられ、渡されます。
このソース ディレクトリから実行します
依存関係をインストールし、npm を介してこのチェックアウトの penecho 実行可能ファイルを公開し、構成して、同じ運用エントリ ポイントを通じて開始します。
npmインストール
npmリンク
ペネチョ設定
ペネチョ
npm link はローカル コマンド リンクを作成します。パッケージは公開されません。個別のビルド手順はなく、ローカル開発では npm start を使用しません。
http://localhost:3888 を開きます。同じ信頼できる LAN 上の他のデバイスは http://<this-computer-LAN-IP>:3888 を使用できます。
以下は例示的な見積もりであり、強制的な PenEcho トークンの予算ではありません。リクエストで 10,000 個の入力トークンと 1,000 個の出力トークンが使用されると仮定すると、標準のショートコンテキスト API コストは次のようになります。
gpt-5.6-sol : 10,000 x $5.00 / 1M + 1,000 x $30.00 / 1M = $0.080
gpt-5.6-terra : 10,000 x $2.50 / 1M + 1,000 x $15.00 / 1M = $0.040
gpt-5.6-luna : 10,000 x $1.00 / 1M + 1,000 x $6.00 / 1M = $0.016
これらの例の量では、リクエストあたり約 1.6 ～ 8 セントになります。実際の入力、推論、出力の使用法は、キャンバスのコンテンツ、モデル、プロバイダー、および再試行動作によって異なります。価格は変更される可能性があるため、現在の料金については OpenAI API の価格ページを確認してください。
ChatGPT を使用して Codex にサインインすると、PenEcho は API キーの代わりにプランに含まれる Codex の使用法を使用します。含まれる制限はプランによって異なり、追加の使用には ChatGPT クレジットが必要になる場合があります。現在のプランと制限については、Codex の価格を参照してください。 Claude CLI モードも同様に、Claude Code によって認証されたアカウントを使用します。 Anthropic API の課金とは異なります。
PenEcho はモデル選択をサポートします

API、Codex CLI、および Claude CLI の実行に対して独立して実行されます。モデルの動作は依然として異なります。モデル固有の問題を見つけた場合は、エグゼキューター、モデル名、再現可能なキャンバスの例、予想される結果と実際の結果、およびシークレットが削除されたスクリーンショットを使用して問題を開いてください。
PenEcho はデフォルトで 0.0.0.0:3888 をリッスンするため、ローカルホストと信頼できる LAN アクセスはすぐに機能します。エグゼキューターに一致するデプロイメント境界を選択します。
Codex CLI および Claude CLI モード: ローカル マシンまたは信頼できる直接接続された LAN でのみ使用します。有効なリクエストはローカル CLI プロセスを開始するため、どちらのモードもパブリック インターネットや信頼できないリバース プロキシに直接公開しないでください。どちらも、パブリック オリジン設定なしで、ローカルホストおよび LAN アドレスからすぐに動作します。 PenEcho は、選択された CLI を起動する前に、ホスト、クライアント ネットワーク、正確なオリジン、プロセス存続期間セッション Cookie、および JSON コンテンツ タイプをチェックします。有効な新しいリクエストはそれぞれ、前のリクエストを即座に置き換えます。キューで待機したり、ビジー応答を返したりすることはありません。
API モード: ローカル、LAN、プロキシ、およびリモートのリクエストは、PenEcho レベルのホストまたはオリジンの制限なしで意図的に受け入れられます。公開する場合は、HTTPS の背後に配置し、認証します。

[切り捨てられた]

## Original Extract

Think with AI beyond the chat box. A shared canvas for handwriting, equations, diagrams, and spatial reasoning. - erickong/penecho

GitHub - erickong/penecho: Think with AI beyond the chat box. A shared canvas for handwriting, equations, diagrams, and spatial reasoning. · GitHub
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
erickong
/
penecho
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .github .github docs docs public public test test .gitattributes .gitattributes .gitignore .gitignore COMMERCIAL-LICENSE.md COMMERCIAL-LICENSE.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTOR-LICENSE-AGREEMENT.md CONTRIBUTOR-LICENSE-AGREEMENT.md LICENSE LICENSE NOTICE NOTICE README.md README.md TRADEMARKS.md TRADEMARKS.md api-config.js api-config.js claude-cli.js claude-cli.js cli.js cli.js codex-cli.js codex-cli.js configure-ui.js configure-ui.js package-lock.json package-lock.json package.json package.json server.js server.js View all files Repository files navigation
Think with AI beyond the chat box.
PenEcho is a shared canvas where handwriting, equations, diagrams, and spatial context become part of the conversation.
Put a question, equation, diagram, or half-formed idea anywhere on the canvas and pause. PenEcho reads your marks and their spatial relationships, then answers beside them. You can work through a problem without translating every step into a chat message or rebuilding it with rigid diagram tools.
Get answers, hints, explanations, continuations, formulas, plots, and diagrams directly on the canvas.
Move, resize, accept, or discard every AI draft before it becomes part of your work.
Draw naturally with a stylus or mouse, then pan and zoom across a sparse 20,000 x 20,000 canvas.
Draw a freehand lasso around confirmed ink to move, resize, or recolor it locally; accepting or cancelling a selection never triggers an AI request.
Choose Arcane, Sci-fi, or Research mode to match the kind of problem you are exploring.
Save lightweight snapshots locally in your browser. Starting a new canvas can overwrite the current snapshot, save a new copy, or continue without saving; unconfirmed AI drafts are never included.
PenEcho keeps a small local runtime and only allocates 512 x 512 tiles where ink exists, so the huge logical canvas does not become a huge bitmap.
flowchart LR
User["Handwriting, equations, and sketches"] --> Canvas["Browser canvas<br/>sparse confirmed tiles"]
Canvas --> Atlas["Cropped visual atlas<br/>plus geometry"]
Atlas --> Server["PenEcho server<br/>validation and prompt"]
Server --> Executor{"Configured executor"}
Executor --> API["API mode<br/>OpenAI-compatible or Anthropic"]
Executor --> Codex["Codex CLI mode<br/>local codex exec"]
Executor --> Claude["Claude CLI mode<br/>local claude -p"]
API --> Draft["Structured editable draft"]
Codex --> Draft
Claude --> Draft
Draft --> Canvas
Loading
The browser sends only the relevant canvas crop and geometry. The server validates the request, uses the selected executor, and returns a movable draft that stays separate from confirmed ink until you accept it.
You need Node.js 18.17+ and one of the following: an API key, an authenticated Codex CLI , or an authenticated Claude Code CLI .
npm install -g penecho
penecho configure
penecho
penecho configure opens the interactive configuration center. Its main menu contains LLM source , Settings , and Exit . Use the arrow keys and Enter to navigate:
LLM source -> Claude CLI selects a detected, recommended, default, or manually entered model and an effort level. Opus 4.8 or newer is recommended; Sonnet and Opus 4.6 can respond but may produce weaker canvas results.
LLM source -> Codex CLI selects a model and effort. GPT-5.5 or newer is required for good results, gpt-5.6-sol is recommended, and xhigh is the highest listed Codex effort.
LLM source -> API selects the OpenAI-compatible or Anthropic/Claude-compatible request format, then asks for the URL, model, effort, and hidden key. Existing values are offered as defaults and a blank key keeps the saved key.
Settings controls the unified model timeout, the image format sent to every model executor, request recording and retention, listening interface and port, and initial Auto AI delay. WebP is the default; PNG is also available. The delay can also be changed on the canvas.
Every LLM page ends with Test & Save , and PenEcho always saves before checking. Codex CLI uses a fast offline check: it verifies the executable and login, then reads codex debug models --bundled to confirm the selected model exists. It does not run inference, attach an image, refresh the online catalog, or consume model tokens. Claude CLI and API configuration still send one small real request to verify the selected endpoint/model settings. Whether a check passes or fails, the configuration remains saved and the UI returns to the parent menu with a clear diagnostic.
The default configuration is ~/.penecho/config.env . API credentials are plaintext in this local file, receive owner-only permissions on POSIX systems, and are never sent to browser code. Protect it like any other credential. If penecho is started before this file exists, it opens the configuration center automatically in an interactive terminal.
Use a different env-style configuration file for a particular launch when needed:
penecho configure --config ./team.env
penecho --config ./team.env
An explicit --config file replaces the default global file for that command. PenEcho does not automatically read a project-directory .env or a package-directory .env .
Installing the Codex desktop app alone does not guarantee that a codex executable is available on the shell PATH . Install and authenticate the CLI separately before selecting Codex:
npm install -g @openai/codex@latest
hash -r
codex --version
codex login status
If needed, run codex login . Claude CLI mode similarly requires an installed and authenticated Claude Code CLI, normally through claude auth login .
PenEcho uses the selected CLI locally and does not need an API key for that source. Normal startup checks the executable and login without consuming model tokens. Codex Test & Save additionally verifies the selected model against the installed CLI's bundled catalog without making a model request; Claude Test & Save sends a small real request.
Canvas requests through Codex use codex exec --json . PenEcho returns as soon as Codex emits the final agent message and turn.completed ; if the CLI process remains alive afterward, it is terminated and cleaned up in the background instead of delaying the canvas response.
Claude CLI requests use one isolated claude -p turn with tools, agents, MCP, prompt suggestions, session persistence, and other nonessential background traffic disabled. PenEcho sets MAX_THINKING_TOKENS=0 for these latency-sensitive canvas turns so Claude Code does not spend a large hidden extended-thinking budget. When an effort is selected, PenEcho applies both Claude's --effort flag and a per-process settings override so a user-level CLAUDE_CODE_EFFORT_LEVEL cannot silently replace it. PenEcho incrementally validates the stream and returns as soon as Claude emits its successful final result ; any attempted tool use aborts the request, while a CLI process that remains alive after the result is terminated and cleaned up in the background.
Transient launch overrides remain available:
penecho doctor --codex
penecho --codex --model gpt-5.6-sol --effort xhigh
penecho --claude --model opus --effort max
penecho --port 4000
--model , --effort , and --port apply only to that process and take precedence over the selected configuration file. Omit them to use the saved choice or the underlying CLI default. Other model-specific effort strings are accepted and passed through.
Run from this source directory
Install dependencies, expose this checkout's penecho executable through npm, configure it, and start it through the same production entry point:
npm install
npm link
penecho configure
penecho
npm link creates the local command link; it does not publish the package. There is no separate build step and local development does not use npm start .
Open http://localhost:3888 . Other devices on the same trusted LAN can use http://<this-computer-LAN-IP>:3888 .
The following is an illustrative estimate, not an enforced PenEcho token budget. Assuming a request uses 10,000 input tokens and 1,000 output tokens, the standard short-context API cost would be:
gpt-5.6-sol : 10,000 x $5.00 / 1M + 1,000 x $30.00 / 1M = $0.080
gpt-5.6-terra : 10,000 x $2.50 / 1M + 1,000 x $15.00 / 1M = $0.040
gpt-5.6-luna : 10,000 x $1.00 / 1M + 1,000 x $6.00 / 1M = $0.016
At those example quantities, that is about 1.6 to 8 cents per request. Actual input, reasoning, and output usage varies by canvas content, model, provider, and retry behavior. Prices can change, so check the OpenAI API pricing page for current rates.
If you sign in to Codex with ChatGPT, PenEcho uses the Codex usage included with your plan instead of an API key. Included limits vary by plan, and additional usage may require ChatGPT credits. See Codex pricing for current plans and limits. Claude CLI mode similarly uses the account authenticated by Claude Code; it is distinct from Anthropic API billing.
PenEcho supports model selection independently for API, Codex CLI, and Claude CLI execution. Model behavior still varies. If you find a model-specific issue, please open an issue with the executor, model name, a reproducible canvas example, expected and actual results, and a screenshot with secrets removed.
PenEcho listens on 0.0.0.0:3888 by default so localhost and trusted-LAN access work immediately. Choose the deployment boundary that matches your executor:
Codex CLI and Claude CLI modes: use them only on the local machine or a trusted, directly connected LAN. A valid request starts a local CLI process, so do not expose either mode directly to the public internet or an untrusted reverse proxy. Both work immediately from localhost and LAN addresses without a public-origin setting. PenEcho checks the Host, client network, exact Origin, process-lifetime session cookie, and JSON content type before launching the selected CLI. Each valid new request immediately supersedes the prior request; it never waits in a queue or returns a busy response.
API mode: local, LAN, proxy, and remote requests are intentionally accepted without PenEcho-level Host or Origin restrictions. If you expose it publicly, place it behind HTTPS, authenticati

[truncated]
