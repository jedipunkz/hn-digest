---
source: "https://www.microsoft.com/en-us/security/blog/2026/06/18/autojack-single-page-rce-host-running-ai-agent/"
hn_url: "https://news.ycombinator.com/item?id=48612716"
title: "AutoJack: A single page can RCE the host running your AI agent"
article_title: "AutoJack: How a single page can RCE the host running your AI agent | Microsoft Security Blog"
author: "p_stuart82"
captured_at: "2026-06-20T21:33:36Z"
capture_tool: "hn-digest"
hn_id: 48612716
score: 4
comments: 0
posted_at: "2026-06-20T20:30:16Z"
tags:
  - hacker-news
  - translated
---

# AutoJack: A single page can RCE the host running your AI agent

- HN: [48612716](https://news.ycombinator.com/item?id=48612716)
- Source: [www.microsoft.com](https://www.microsoft.com/en-us/security/blog/2026/06/18/autojack-single-page-rce-host-running-ai-agent/)
- Score: 4
- Comments: 0
- Posted: 2026-06-20T20:30:16Z

## Translation

タイトル: AutoJack: 単一ページで AI エージェントを実行しているホストを RCE できる
記事のタイトル: AutoJack: 単一ページで AI エージェントを実行しているホストを RCE できる方法 |マイクロソフト セキュリティ ブログ
説明: AutoJack は、単一の悪意のある Web ページがどのように AI ブラウジング エージェントをホスト マシン上のリモート コード実行ベクトルに変えることができるかを示す新しいエクスプロイト チェーンです。攻撃者は、localhost の信頼性、認証の欠如、安全でないパラメータの処理を悪用することで、任意のプロセスの実行を引き起こす可能性があります。
[切り捨てられた]

記事本文:
AutoJack: 単一ページで AI エージェントを実行しているホストを RCE する方法 |マイクロソフト セキュリティ ブログ
コンテンツにスキップ
メインコンテンツにスキップ
セキュリティ
マイクロソフトディフェンダー
マイクロソフト エントラ
Microsoft Intune
マイクロソフトの権限
Microsoft セキュリティ副操縦士
マイクロソフトセンチネル
すべての製品を見る
AIを活用したサイバーセキュリティ
クラウドセキュリティ
データセキュリティとガバナンス
ID とネットワーク アクセス
プライバシーとリスク管理
AIのセキュリティ
中小企業
統合された SecOps
ゼロトラスト
価格設定
サービス
パートナー
Microsoft セキュリティを選ぶ理由
サイバーセキュリティの意識
お客様の事例
セキュリティ101
製品のトライアル
Microsoft を保護する方法
業界での認知度
Microsoft セキュリティ インサイダー
Microsoft デジタル ディフェンス レポート
セキュリティ対応センター
マイクロソフト セキュリティ ブログ
マイクロソフトのセキュリティ イベント
マイクロソフト技術コミュニティ
ドキュメント
技術コンテンツライブラリ
トレーニングと認定資格
Microsoft クラウドのコンプライアンス プログラム
マイクロソフト トラスト センター
セキュリティ エンジニアリング ポータル
サービストラストポータル
マイクロソフト セキュア フューチャー イニシアチブ
ビジネスソリューションハブ
営業担当者へのお問い合わせ
無料トライアルを開始する
マイクロソフトのセキュリティ
アズール
ダイナミクス 365
マイクロソフト 365
マイクロソフトチーム
Windows 365
マイクロソフトAI
アズールスペース
複合現実
マイクロソフト ホロレンズ
マイクロソフト ビバ
量子コンピューティング
持続可能性
教育
自動車
金融サービス
政府
ヘルスケア
製造業
小売
パートナーを探す
パートナーになる
パートナーネットワーク
マイクロソフト マーケットプレイス
ソフトウェア会社
ブログ
マイクロソフトの広告
開発者センター
ドキュメント
イベント
ライセンス
Microsoft Learn
マイクロソフトリサーチ
サイトマップを見る
ホーム
AutoJack: 単一ページで AI エージェントを実行しているホストを RCE する方法
AutoJack: 単一ページで AI エージェントを実行しているホストを RCE する方法
/
1x
Microsoft Copilot を搭載
シェアする
エージェント フレームワークに注目する理由 AutoGen Studio とは
オートジャックc

一目で分かる
チェーンの構造 問題 1: エージェント自体が無効にするオリジン ホワイトリスト
問題 2: MCP をオプトアウトする認証ミドルウェア
問題 3: URL の server_params がコマンド ラインである
まとめ: 現実的なシナリオ 適用された修正と強化措置
緩和と保護のガイダンス Microsoft がエージェント システムのセキュリティを保護する方法
これが広範なエージェント エコシステムにとって何を意味するか
AI エージェント フレームワークのセキュリティに関する継続的な調査により、AutoGen Studio (AutoGen のオープンソース プロトタイピング ユーザー インターフェイス) のエクスプロイト チェーンが特定されました。これにより、ブラウジング エージェントによってレンダリングされた信頼できない Web コンテンツがローカルのモデル コンテキスト プロトコル (MCP) WebSocket に到達し、ホスト上で任意のプロセスを生成できるようになります。 AutoJack と呼ばれるこの手法は、多くの開発者ツールが依存するローカルホストの信頼境界を越えることにより、エージェントをジャックして攻撃者のラストマイル配送手段にします。
この動作を Microsoft セキュリティ レスポンス センター (MSRC) に報告しました。この報告を受けて、メンテナはコミット b047730 で上流のメイン ブランチを強化しました。この問題は開発中に特定され、解決されました。影響を受ける MCP WebSocket サーフェスは Python Package Index (PyPI) リリースには含まれていないため、PyPI から AutoGen Studio をインストールするユーザーは、この特定のチェーンにさらされることはありません。
より広範な教訓は一般的です。エージェントが信頼されていないページを参照し、特権のあるローカル サービスと通信できる場合、ループバックが攻撃対象領域になる可能性があり、コントロール プレーンを認証、許可、分離する必要があります。
エージェント フレームワークに注目する理由
最新の AI エージェントは単なるテキスト ジェネレーターではありません。ファイルを読み取り、ページを参照し、API を呼び出し、ツールにシェルアウトします。まさにそれがフレームの有用性であり、フレーム内の体系的な実行リスクの発見に投資が行われる理由です。

モデルをツールに接続する eworks。このシリーズの前半で、Microsoft Semantic Kernel の RCE プリミティブについて説明しました。この投稿では、スタックの 1 つ上の層をインフラストラクチャおよび開発者向けのプロトタイピング サーフェスに移動し、これらのツールを実験に価値のあるものにする同じエージェント機能が、安全対策なしでプロトタイプが実行されるときにリモート コード実行の配信チャネルになる仕組みを示します。
重要なのは、プロトタイプを避けることではないということです。それは次のとおりです。コア サーバーまたはラップトップ上のエージェントがオープン Web を参照し、特権のあるローカル サービスと通信できるようになると、localhost は信頼境界ではなくなります。ディフェンダーはそのための計画を立てる必要があり、これらの調査結果はその理由を示しています。
AutoGen Studio は、マルチエージェント システム用の Microsoft Research のフレームワークである AutoGen の上にあるユーザー インターフェイス (UI) です。これにより、開発者はエージェントを作成し、MCP サーバーなどのツールを接続し、簡単な実験を実行できます。そのドキュメントには使用目的が明確に記載されています。言い換えれば、これは開発者エクスペリエンスとのトレードオフが予想される研究プロトタイプであり、デフォルトは強化されたデプロイメントではなく反復が容易になるように調整されています。
AutoJack チェーンの概要
以下の説明は、説明のみを目的としています。エクスプロイト チェーンは現在のビルドでは機能しません。ここに含まれるのは、防御者が他のエージェント フレームワークのパターンを認識できるようにするためです。
エクスプロイト チェーンは、AutoGen Studio の MCP WebSocket サーフェスに 3 つの独立した弱点を構成します。
オリジン ホワイトリストは localhost を信頼します – しかし、ローカル エージェントは localhost です (CWE-1385 – WebSocket でオリジン検証がありません): MCP WebSocket は、オリジンが http://127.0.0.1 または http://localhost である接続のみを受け入れます。これにより、evil.com を指すブラウザがブロックされます。 AutoGen エージェントが所有するヘッドレス ブラウザによってレンダリングされる JavaScript はブロックされません。

同じマシンです。
認証ミドルウェアは MCP パスに対してオプトアウトされています (CWE-306 – 重要な機能の認証がありません): AutoGen Studio の認証ミドルウェアは、/api/mcp/* (および /api/ws/*) が独自のチェックを行うことを前提として、これらを明示的にスキップしました。 MCP WebSocket ハンドラーは、そのフォローアップ チェックを実装していませんでした。その結果、MCP WebSocket は、アプリの残りの部分に構成された認証モードに関係なく、認証なしで接続を受け入れました。
URL からの StdioServerParams はそのまま実行されます (CWE-78 – OS コマンドで使用される特殊要素の不適切な中和): エンドポイントは、server_params クエリ パラメーターを受け入れ、JSON BLOB を Base64 デコードして StdioServerParams にし、コマンド + 引数を stdio_client(…) に渡します。ホワイトリストは存在せず、calc.exe、powershell.exe -enc …、または bash -c ‘…’ はすべて「MCP サーバー」として受け入れられました。
これらを、同じマシン上で実行されている AutoGen エージェントによってレンダリングされるオープン インターネット上の Web ページと連結すると、リモート コード実行プリミティブが得られます。エージェントに攻撃者のページを表示させる以外に、ユーザーの操作は必要ありません。
私たちはこの手法を AutoJack と名付けました。攻撃者はブラウジング エージェントをカージャックし、それを混乱した代理人として使用し、ローカルホストの境界を越えて AutoGen Studio の MCP コントロール プレーンに侵入します。
問題 1: エージェント自体が無効にするオリジンのホワイトリスト
AutoGen Studio の MCP WebSocket は、ブラウザ主導のクロスサイト WebSocket ハイジャック (CSWSH) に対する従来の防御に依存しており、127.0.0.1 / localhost からの同一オリジン接続のみを許可します。
allowed_origins = [“http://127.0.0.1”, “http://localhost”]
これは、evil[.]com へのタブを開いている人間のユーザーにとって適切なコントロールです。ブラウザは Origin ヘッダーを hxxps://evil[.]com に設定し、チェックは失敗します。

接続は拒否されます。
発信元チェックだけでは、エージェントにとって適切な管理とは言えません。 MultimodalWebSurfer、fetch_webpage_tool、Playwright がサポートするサーファー、またはリクエスト/WebSocket を実行するコード実行ツールなどの組み込み Web ブラウジング ツールを備えた AutoGen エージェントは、ワークステーション上のプロセスです。ロードされるものはすべてローカルホスト ID を継承します。ヘッドレス ブラウザによって実行される JavaScript の「オリジン」は、エージェントがナビゲートしたものすべてであり、エージェントが行う WebSocket 呼び出しには、ホワイトリストを満たすオリジンが含まれます。
問題 2: MCP をオプトアウトする認証ミドルウェア
AutoGen Studio は、いくつかの認証モード (なし、github、msal、firebase) をサポートしています。これらはすべて、FastAPI ルートのディスパッチに先立って実行される単一の AuthMiddleware に接続されています。 PyPI のバージョンでは、そのミドルウェアには WebSocket スタイルのパスの早期リターンが含まれています。
# これらのパスでは認証は除外されます。彼らは独自のチェックを行うことを目的としていた
request.url.path.startswith("/api/ws") または request.url.path.startswith("/api/mcp") の場合:
return await call_next(リクエスト)
この意図は合理的です。ASGI ミドルウェアは、HTTP リクエストをゲートするのと同じ方法で WebSocket ハンドシェイクを意味のあるゲートで制御できないため、WebSocket ハンドラーが受け入れ時に認証自体を強制する設計になっています。 MCP (Model Context Protocol) ルートはその責任を負いません。その結果、リリースされたパッケージには次の表が当てはまります。
config.yaml で認証を有効にしても、このホールは単独では閉じられません。
問題 3: URL の server_params がコマンド ラインである
開発ビルドの MCP WebSocket ルートは、server_params クエリ パラメーターを読み取り、base64 でデコードし、JSON で解析して StdioServerParams に変換し、それを stdio_client(…) に渡します。
@router.websocket("/ws/{session_id}")
async def mcp_websocket(web

ソケット: WebSocket、セッション ID: str):
encoded = websocket.query_params.get("server_params")
デコードされた =base64.b64decode(エンコードされた)
params = StdioServerParams(**json.loads(デコードされた))
await create_mcp_session(ブリッジ、パラメータ、セッションID)
StdioServerParams.command と StdioServerParams.args は stdio_client に渡され、stdio_client はそれらを使用して MCP 「サーバー」プロセスを生成します。実行可能ファイルが MCP 対応バイナリであるという許可リストはないため、同じ配管で calc.exe、powershell.exe -enc …、または bash -c ‘…’ が問題なく生成されます。
{
"type": "StdioServerParams",
"コマンド": "calc.exe",
"引数": [],
"env": { "pwned": "true" }
}
Base64 でエンコードしてクエリ文字列にすると、完全なリーチは次のようになります。
ws://localhost:8081/api/mcp/ws/ ?server_params=
問題 1 と 2 を組み合わせると、攻撃者が必要とするのは、エージェントがその URL を開くページをレンダリングすることだけです。
まとめ: 現実的なシナリオ
エンドツーエンド チェーンを検証するために、2 つの小さなハーネスを作成しました。
悪意のある Web_server.py : http://攻撃者[.]example/websocket-interactive で提供される Web ページ。その唯一の意味のあるコンテンツは、calc.exe を実行する Base64 ペイロードで上記の WebSocket を開く <script> です。
web_summarizer_app.py : Flask UI にラップされた小さな「Web Content Summarizer」AutoGen エージェント。アプリはユーザーから URL を取得し、「この URL を参照してその内容を要約してください」というプロンプトとともにそれを MultimodalWebSurfer エージェントに渡します。言い換えれば、これは誰でもフレームワーク上に構築できる本格的な AutoGen エージェントです。Flask ページは単なるインターフェイスです。
エンドツーエンドのフローは次のようになります。
開発者は、AutoGen Studio と同じマシン上で実行される、Web Page Summarizer などの AutoGen エージェント、またはブラウジング機能を備えた任意のエージェントを構築しました。
攻撃者が正規のニュース サイトに悪意のあるコメントを投稿し、

または、ユーザーがサマライザ エージェントに攻撃者が制御する URL を要約するよう依頼します。これは、直接プロンプト、以前のコンテンツへのプロンプト挿入、またはアプリの URL フィールドを通じて発生する可能性があります。
次に、エージェントのブラウジング ツール (この場合は MultimodalWebSurfer) が、ヘッドレス ブラウザを攻撃者のページに移動します。
ページの JavaScript により、ws://localhost:8081/api/mcp/ws/<id>?server_params=<base64> が開きます。ブラウザは同じマシン上にあるため、Origin を使用できます。認証ミドルウェアは /api/mcp/* をショートするため、トークンは必要ありません。
AutoGen Studio はペイロードをデコードし、開発者のアカウントで calc.exe (またはその他のもの) を実行します。
このデモンストレーションは、制御されたローカルの概念実証としてパッケージ化されていることに注意してください。エンドツーエンドでご覧ください。
以下のスクリーンショットは、単一のワークステーション上の完全なチェーンを示しています。開発者は、localhost:8081 (デフォルトのポート) で AutoGen Studio を起動し、Web Content Summarizer アプリを開き、攻撃者が制御する URL を送信します。 MultimodalWebSurfer がページをレンダリングしてから数秒以内に、ブラウザやエージェントのヘッドレス Chromium ではなく、AutoGen Studio プロセスによって起動される calc.exe が開発者のデスクトップに表示されます。
修正と強化策が適用されました
この問題は、Microsoft Security Response Center の協力により修正されました。メンテナー

[切り捨てられた]

## Original Extract

AutoJack is a novel exploit chain showing how a single malicious webpage can turn an AI browsing agent into a remote code execution vector on the host machine. By abusing trust in localhost, missing authentication, and unsafe parameter handling, attackers can trigger arbitrary process execution thro
[truncated]

AutoJack: How a single page can RCE the host running your AI agent | Microsoft Security Blog
Skip to content
Skip to main content
Security
Microsoft Defender
Microsoft Entra
Microsoft Intune
Microsoft Purview
Microsoft Security Copilot
Microsoft Sentinel
View all products
AI-powered cybersecurity
Cloud security
Data security & governance
Identity & network access
Privacy & risk management
Security for AI
Small and medium business
Unified SecOps
Zero Trust
Pricing
Services
Partners
Why Microsoft Security
Cybersecurity awareness
Customer stories
Security 101
Product trials
How we protect Microsoft
Industry recognition
Microsoft Security Insider
Microsoft Digital Defense Report
Security Response Center
Microsoft Security Blog
Microsoft Security Events
Microsoft Tech Community
Documentation
Technical Content Library
Training & certifications
Compliance Program for Microsoft Cloud
Microsoft Trust Center
Security Engineering Portal
Service Trust Portal
Microsoft Secure Future Initiative
Business Solutions Hub
Contact Sales
Start free trial
Microsoft Security
Azure
Dynamics 365
Microsoft 365
Microsoft Teams
Windows 365
Microsoft AI
Azure Space
Mixed reality
Microsoft HoloLens
Microsoft Viva
Quantum computing
Sustainability
Education
Automotive
Financial services
Government
Healthcare
Manufacturing
Retail
Find a partner
Become a partner
Partner Network
Microsoft Marketplace
Software companies
Blog
Microsoft Advertising
Developer Center
Documentation
Events
Licensing
Microsoft Learn
Microsoft Research
View Sitemap
Home
AutoJack: How a single page can RCE the host running your AI agent
AutoJack: How a single page can RCE the host running your AI agent
/
1x
Powered by Microsoft Copilot
Share
Why we are looking at agent frameworks What is AutoGen Studio
The AutoJack chain at a glance
Anatomy of the chain Issue 1: Origin allowlist that the agent itself defeats
Issue 2: Auth middleware that opts MCP out
Issue 3: server_paramsfrom the URL is the command line
Putting it together: a realistic scenario Fixes and hardening measures applied
Mitigation and protection guidance How Microsoft helps secure agentic systems
What this means for the broader agent ecosystem
Ongoing research into AI agent framework security identified an exploit chain in AutoGen Studio (AutoGen’s open-source prototyping user interface) that allows untrusted web content rendered by a browsing agent to reach a local Model Context Protocol (MCP) WebSocket and spawn arbitrary processes on the host. The technique, which we call AutoJack, jacks the agent into becoming the attacker’s last-mile delivery vehicle by crossing the localhost trust boundary that many developer tools rely on.
We reported the behavior to the Microsoft Security Response Center (MSRC); following the report the maintainers hardened the upstream main branch in commit b047730. This issue was identified and addressed during development. The affected MCP WebSocket surface was never included in a Python Package Index (PyPI) release, so users who install AutoGen Studio from PyPI aren’t exposed to this specific chain.
The broader lesson is general: if an agent can browse untrusted pages and also talk to privileged local services, loopback can become an attack surface and control planes must be authenticated, authorized, and isolated.
Why we are looking at agent frameworks
Modern AI agents are not just text generators. They read files, browse pages, call APIs, and shell out to tools. That is exactly what makes them useful, and exactly why there is investment in finding systemic execution risks in the frameworks that wire models to tools. Earlier in this series we covered RCE primitives in Microsoft Semantic Kernel . In this post we move one layer up the stack to an infrastructure and developer-facing prototyping surface and show how the same agent capabilities that make these tools valuable for experimentation can become a delivery channel for remote code execution when the prototype runs without safeguards.
The takeaway is not to avoid prototypes. It is this: when an agent on your core server or laptop can browse the open web and communicate with privileged local services, localhost stops being a trust boundary. Defenders need to plan for that, and these findings show why.
AutoGen Studio is a user interface (UI) on top of AutoGen , Microsoft Research’s framework for multi-agent systems. It lets developers compose agents, attach tools, including MCP servers, and run quick experiments. Its documentation is clear about intended use. In other words, it is a research prototype with expected developer-experience tradeoffs: defaults tuned for ease of iteration rather than hardened deployment.
The AutoJack chain at a glance
The explanation below is for demonstrative purposes only. The exploit chain doesn’t work on current builds. It is included here so that defenders can recognize the pattern in other agent frameworks.
The exploit chain composes three independent weaknesses in AutoGen Studio’s MCP WebSocket surface:
Origin allowlist trusts localhost – but a local agent is localhost (CWE-1385 – Missing Origin Validation in WebSockets): The MCP WebSocket only accepts connections whose Origin is http://127.0.0.1 or http://localhost. That blocks a browser pointed at evil.com. It does not block JavaScript that is rendered by a headless browser owned by an AutoGen agent on the same machine .
Authentication middleware is opt-out for MCP paths (CWE-306 – Missing Authentication for Critical Function): The auth middleware in AutoGen Studio explicitly skipped /api/mcp/* (and /api/ws/*) on the assumption that these would do their own checks. The MCP WebSocket handler did not implement that follow-up check. As a result, the MCP WebSocket accepted connections without any authentication regardless of the auth mode configured for the rest of the app.
StdioServerParams from the URL is executed verbatim (CWE-78 – Improper Neutralization of Special Elements used in an OS Command): The endpoint accepted a server_params query parameter, base64-decoded a JSON blob into StdioServerParams, and handed command + args to stdio_client(…). There was no allowlist – calc.exe, powershell.exe -enc …, or bash -c ‘…’ were all accepted as “MCP servers.”
Chain these together with a webpage on the open internet, rendered by an AutoGen agent running on the same machine, and you have a remote code execution primitive. No user interaction is required beyond getting the agent to render the attacker’s page.
We named the technique AutoJack: an attacker carjacks the browsing agent and uses it as a confused deputy to drive across the localhost boundary into AutoGen Studio’s MCP control plane.
Issue 1: Origin allowlist that the agent itself defeats
AutoGen Studio’s MCP WebSocket relies on the conventional defense for browser-driven cross-site WebSocket hijacking (CSWSH): allow only same-origin connections from 127.0.0.1 / localhost.
allowed_origins = [“http://127.0.0.1”, “http://localhost”]
That is the right control for a human user opening a tab to evil[.]com. The browser will set the Origin header to hxxps://evil[.]com, the check will fail, and the connection will be refused.
Origin checks alone are not the right control for an agent. An AutoGen agent equipped with built-in web-browsing tooling, such as MultimodalWebSurfer, fetch_webpage_tool, any Playwright-backed surfer, or a code-execution tool that runs requests/websockets is a process on the workstation . Anything it loads inherits the localhost identity. The “origin” of any JavaScript executed by that headless browser is whatever the agent navigated to – and the WebSocket call it then makes carries an Origin that satisfies the allowlist.
Issue 2: Auth middleware that opts MCP out
AutoGen Studio supports several authentication modes (none, github, msal, firebase). All of them are wired into a single AuthMiddleware that runs ahead of FastAPI route dispatch. In the version on PyPI, that middleware contains an early-return for WebSocket-style paths:
# auth excluded for these paths; they were intended to do their own checks
if request.url.path.startswith("/api/ws") or request.url.path.startswith("/api/mcp"):
return await call_next(request)
The intent is reasonable: ASGI middlewares cannot meaningfully gate WebSocket handshakes the same way they gate HTTP requests, so the design called for the WebSocket handler to enforce auth itself at accept time. The MCP (Model Context Protocol) route never picked up that responsibility. As a result, the table below holds for the released package:
Turning on auth in config.yaml does not close this hole on its own.
Issue 3: server_paramsfrom the URL is the command line
The MCP WebSocket route in the development build reads a server_params query parameter, base64-decodes it, JSON-parses it into StdioServerParams, and passes that into stdio_client(…):
@router.websocket("/ws/{session_id}")
async def mcp_websocket(websocket: WebSocket, session_id: str):
encoded = websocket.query_params.get("server_params")
decoded = base64.b64decode(encoded)
params = StdioServerParams(**json.loads(decoded))
await create_mcp_session(bridge, params, session_id)
StdioServerParams.command and StdioServerParams.args are passed to stdio_client, which uses them to spawn an MCP “server” process. There is no allowlist that the executable be an MCP-speaking binary, so the same plumbing happily spawns calc.exe, powershell.exe -enc …, or bash -c ‘…’.
{
"type": "StdioServerParams",
"command": "calc.exe",
"args": [],
"env": { "pwned": "true" }
}
Base64-encoded into a query string, the full reach-out is:
ws://localhost:8081/api/mcp/ws/ ?server_params=
Combined with Issues 1 and 2, all an attacker needs is for the agent to render a page that opens that URL.
Putting it together: a realistic scenario
To validate the end-to-end chain, we wrote two tiny harnesses:
malicious_web_server.py : a web page served at http://attacker[.]example/websocket-interactive. Its only meaningful content is a <script> that opens the WebSocket above with a base64 payload that runs calc.exe.
web_summarizer_app.py : a small “Web Content Summarizer” AutoGen agent wrapped in a Flask UI. The app takes a URL from the user and hands it to a MultimodalWebSurfer agent with the prompt “Browse this URL and summarize its content.” It is, in other words, a fully-fledged AutoGen agent that anyone could build on top of the framework – the Flask page is just the interface.
The end-to-end flow looks like this:
The developer has built an AutoGen agent such as a Web Page Summarizer, or any agent with browsing capabilities, that runs on the same machine as AutoGen Studio.
An attacker plants a malicious comment on a legitimate news site, or a user asks the summarizer agent to summarize an attacker-controlled URL. This can happen through a direct prompt, a prompt injection in earlier content, or a URL field in the app.
The agent’s browsing tool, MultimodalWebSurfer in our case, then navigates the headless browser to the attacker’s page.
The page’s JavaScript opens ws://localhost:8081/api/mcp/ws/<id>?server_params=<base64>. Because the browser is on the same machine, the Origin is acceptable; because the auth middleware short-circuits /api/mcp/*, no token is required.
AutoGen Studio decodes the payload and runs calc.exe (or anything else) under the developer’s account.
Note that we packaged the demonstration as a controlled local proof of concept , See it end-to-end.
The screenshots below show the full chain on a single workstation: the developer launches AutoGen Studio on localhost:8081 (the default port), opens the Web Content Summarizer app, and submits an attacker-controlled URL. Within seconds of MultimodalWebSurfer rendering the page, calc.exe pops on the developer’s desktop, launched by the AutoGen Studio process, not by the browser and not by the agent’s headless Chromium.
Fixes and hardening measures applied
The issue was fixed with the help of the Microsoft Security Response Center. The maintainers

[truncated]
