---
source: "https://github.com/agentsploit/agentsploit"
hn_url: "https://news.ycombinator.com/item?id=48460764"
title: "AgentSploit – Burp Suite for AI Agents and MCP Servers)"
article_title: "GitHub - agentsploit/agentsploit: Offensive security framework for AI agents and MCP servers. · GitHub"
author: "di_desle"
captured_at: "2026-06-09T16:07:22Z"
capture_tool: "hn-digest"
hn_id: 48460764
score: 2
comments: 0
posted_at: "2026-06-09T13:18:07Z"
tags:
  - hacker-news
  - translated
---

# AgentSploit – Burp Suite for AI Agents and MCP Servers)

- HN: [48460764](https://news.ycombinator.com/item?id=48460764)
- Source: [github.com](https://github.com/agentsploit/agentsploit)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T13:18:07Z

## Translation

タイトル: AgentSploit – AI エージェントおよび MCP サーバー用の Burp スイート)
記事のタイトル: GitHub - Agentsploit/agentsploit: AI エージェントと MCP サーバー用の攻撃的なセキュリティ フレームワーク。 · GitHub
説明: AI エージェントと MCP サーバーのための攻撃的なセキュリティ フレームワーク。 - エージェントスプロイト/エージェントスプロイト

記事本文:
GitHub - Agentsploit/agentsploit: AI エージェントと MCP サーバー用の攻撃的なセキュリティ フレームワーク。 · GitHub
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
エージェントを悪用する
/
エージェントを悪用する
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メートル

ain ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
24 コミット 24 コミット .github/ workflows .github/ workflows docs docs 例 例 src/ Agentsploit src/ Agentsploit テスト テスト ui ui .gitignore .gitignore .python-version .python-version AUTHORIZATION.md AUTHORIZATION.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントと MCP サーバーのための攻撃的なセキュリティ フレームワーク。
AgentSploit は、エージェント AI 攻撃対象領域専用に構築された Burp Suite / Metasploit スタイルのフレームワークです。これは、レッド チーマー、AI セキュリティ研究者、製品セキュリティ チームが LLM エージェントとモデル コンテキスト プロトコル (MCP) サーバーを調査して、従来のツールでは検出できない脆弱性を見つけるのに役立ちます。
AgentSploit は、使用が許可されたセキュリティ テスト ツールです。自分が所有していないターゲットをスキャンするには、書面による明示的な許可が必要です。何かを実行する前に、AUTHORIZATION.md を参照してください。
👉 ここは初めてですか? docs/getting-started.md から始めましょう - バンドルされたフィクスチャに対するすべての機能を説明する 10 分間のツアーです。API キーは必要ありません。
すべてのフォーチュン 500 企業は、2026 年に LLM エージェントと MCP サーバーを出荷します。攻撃対象領域はまったく新しいものです。
ツールの説明は LLM で読み取り可能な指示であり、悪意のあるものはエージェントをハイジャックする可能性があります。
エージェントは、PDF、Web ページ、カレンダーの招待状、チケットから信頼できないコンテンツを取得し、そのコンテンツがコマンドを発行する可能性があります。
チェーンされたツール呼び出しにより、従来の権限モデルでは把握できない権限昇格パスが作成されます。
メモリおよびコンテキスト ウィンドウはセッション間で汚染される可能性があります。
既存のスキャナー (Burp、ZAP、Semgrep、Snyk) はこのレイヤーを認識しません。 AgentSploit はそうします。
pip install エージェントの悪用
# 婚約の足場を築く
エージェントスロイ

t init my-engagement/ --authorized-by " Jane Doe <ciso@example.com> "
cd 私の婚約/
# MCP サーバーをスキャンします (トレーニング モード = API キーは不要、バンドルされたフィクスチャ)
Agentsploit scan mcp stdio://./tests/fixtures/vulnerable_mcp/server.py --training
# ブラウザで結果を参照します (ライブ エンゲージメント ダッシュボード、v1.6 以降)
エージェントスプロイト サーブ --トレーニング
# -> http://127.0.0.1:8800 (起動時にトークンが出力される)
能力
AgentSploit には、エージェントの攻撃対象領域をカバーする 11 個のモジュールが同梱されています。それぞれはバンドルされた脆弱なフィクスチャとともにエンドツーエンドで文書化されているため、 --training を使用して API キーを使用せずにローカル ターゲットに対してすべてを実行できます。
stdio 、 Streamable HTTP 、または SSE 経由で MCP サーバーに接続し、そのツール、リソース、プロンプト、および (HTTP/SSE の場合) HTTP サーフェスに対して一連のチェックを実行します。
在庫チェック (すべての輸送):
2. 間接プロンプトインジェクションペイロードジェネレーター
エージェントが信頼できないコンテンツを安全に処理するかどうかをテストするために、ラベル付きペイロードを生成します。
direct - 直接的なオーバーライドの試行
role_confusion - 偽のシステム: / アシスタント: ターン
区切り文字 - フェンスで囲まれたコンテンツのエスケープと再コンテキスト
unicode_tag - 目に見えない Unicode タグブロックの密輸 (U+E0000 範囲)
tool_smuggling - 説明文内の隠されたツール呼び出しの呼び出し
マークダウン - README/コメント形式
html - 非表示要素を含むページコンテンツ
pdf - 表示 + 非表示層 PDF
電子メール - HTML 本文とヘッダーを含む RFC 5322
ical - 悪意のある説明を含む .ics カレンダーの招待
すべてのペイロードにはカナリア文字列がタグ付けされているため、インジェクションの成功をログで検出できます。
生成されたペイロードとエージェント構成を取得し、実際の LLM を通じてペイロードを駆動します。エージェントの応答、ツール呼び出し、または思考トレースにカナリアが出現した場合、その注入は悪用可能であることが確認されます。完全なトレースが永続的であり、CRITICAL/HIGH の結果が得られます。

または監査。
アダプター (v0.9): anthropic (本物の Claude ツール使用)、openai (ツール使用によるチャット完了)、http (OpenAI 形状のコントラクトを備えた汎用 HTTP エージェント - カスタム形状のサブクラス)、モック (決定的、テスト用)。 docs/runner.md を参照してください。
4. パーミッショングラフマッパー (v0.4)
複数の MCP サーバーにわたるツールを列挙し、それぞれを権限 (ソース / ピボット / シンク) ごとに分類し、データ フロー エッジを推測して、信頼性の低いソースから影響力の高いシンクへの攻撃パスを見つけます。ツールチェーン用の BloodHound。
エージェントスプロイト マップ ビルド --targets ./examples/map-targets.yaml --auth ./auth.yaml
エージェントスプロイト マップ エクスポート --graph ./engagements/ < id > / < sid > /permission_graph.json -f mermaid -ograph.md
シンク特権
重大度
実行 ( run_command 、 eval )
クリティカル
MUTATION ( git_push 、 delete_* )
高
EGRESS ( send_email 、 Webhook )
高
パス検出はテスト可能な仮説です。v0.5 検証ツールと組み合わせて、悪用可能性をエンドツーエンドで確認します。 docs/mapper.md を参照してください。
マッパーのループを閉じます。マッパーが推測したパスを取得し、実際のエージェントまたは模擬エージェントを通じてパスをターゲットとしたペイロードを駆動し、チェーンが実際に完了するかどうかを証明します。
エージェントの侵害によるパスの検証 \
--graph ./engagements/ < id > / < sid > /permission_graph.json \
--read_file から --send_email へ \
--training # または --agent ./agent-anthropic.yaml --auth ./auth.yaml
結果
意味
重大度
確認済み
シンク ツールが引数にカナリアを指定して呼び出されました
シンク権限に関連付けられています (EXEC → CRITICAL)
部分的
シンクに到達したか、カナリアが別の場所に浮上しましたが、チェーンは不完全です
高
失敗しました
カナリア面はどこにもありません
情報
確認された結果は、マッパーの仮説を「もっともらしい攻撃経路」から「実証済みのエクスプロイト」に移行させます。 docs/verifier.md を参照してください。
6. バッチパス検証 (v0.6)
1 つのコマンドでグラフ内のすべてのパスにわたってベリファイアを実行する - 一般的なワークフロー

まず安価な模擬エージェントに対してトリアージを行い、次に実際のモデルに対して確認のみを再実行します。
# 格安トリアージパス (無料、即時)
エージェントスプロイトはすべてのパスを検証します --graph ./.../permission_graph.json --training
# 同じグラフ上の実際のモデルのパス
エージェントスプロイトはすべてのパスを確認します --graph ./.../permission_graph.json \
--agent ./agent-anthropic.yaml --auth ./auth.yaml \
--Parallel 3 --max-paths 20
(ソース、シンク) ペアごとに重複を排除し、レート制限を意識した同時実行で並列処理し、パスごとのエラーを分離し、確認率のパーセンテージを含む集計結果の batch_summary を出力します。 docs/verifier.md を参照してください。
デフォルトのベリファイアは 1 つのインジェクション エンベロープ ( role_confusion ) を使用します。 v0.7 では、さらに 4 つ ( direct 、 delimiter 、 unicode_tag 、 tool_smuggling ) と、1 つが見つかるまで順番にそれらを試行するファザーが追加されています。どのエンベロープが勝つかを知ることで、防御側はインジェクション フィルターが何を見逃したかを知ることができます。
# シングルパスファズ
Agentsploit verify fuzz-path --graph ./.../permission_graph.json \
--read_file から --to send_email --training
# バッチファズ - すべてのパス × すべてのテクニック、パスごとに早期終了あり
エージェントスプロイトはすべてのパスを確認します --graph ./.../permission_graph.json \
--fuzz --techniques 役割の混乱、区切り文字、unicode_tag \
--パラレル 3 --トレーニング
テクニック
これが着地するとディフェンダーが奪う
役割_混乱
チャット テンプレート フィルターが偽の <system> ターンを捕捉しない
区切り文字
信頼できないコンテンツの境界は強制されません
unicode_tag
防御は印刷可能な ASCII を削除しますが、U+E0000 タグ ブロックは印刷できません
ツール_密輸
エージェント ランタイムは、説明テキストから JSON ツール呼び出し構文を解析します。
直接
即時注入防御がまったく導入されていない
docs/verifier.md を参照してください。
最初のマルチフェーズ攻撃モジュール。攻撃者は、作成したメモをエージェントの共有ストレージに植え付けます。別の被害者エージェントを実行するとメモが取得され、

攻撃者のカナリアを使用してシンク ツールを起動することに踏み切りました。これが捕捉する修復パターンは、取得したストレージ コンテンツをデータではなく命令として扱うエージェントです。
# 模擬エージェントに対するメモリポイズニング攻撃を検証します (無料、インスタント)
エージェントスプロイトポイズン検証 \
--sink-tool send_email --sink-arg body \
--sink-privilege egress \
--トレーニング
# 本物のクロードに対して
エージェントスプロイトポイズン検証 \
--sink-tool send_email --sink-arg body \
--agent ./agent-anthropic.yaml --auth ./auth.yaml
結果
意味
重大度
確認済み
被害者は引数でカナリアを使用してシンクを呼び出しました
シンク権限に関連付けられています (EXEC → CRITICAL)
部分的
メモは取得されましたが、カナリアはシンクに表示されませんでした
高
NOT_RETRIEVED
メモは保存されているが、被害者はそれを読んでいない
情報
NOT_STORED
攻撃者による書き込み失敗
情報
docs/poisoning.md を参照してください。
9. RAG / ベクターストアポイズニング (v1.1)
v0.8 と同じ脅威モデルですが、検索拡張生成のケースに拡張されました。攻撃者は被害者のクエリ キーを知らず、トピックだけを知っています。ターゲット クエリの正規のコーパス コンテンツよりも上位にある有害なドキュメントのインデックスを作成します。被害者エージェントは semantic_search を実行し、上位にランク付けされた一致を取得し、埋め込まれたチェーンに従います。
エージェントスプロイトポイズン verify-rag \
--sink-tool send_email --sink-arg body \
--sink-privilege egress \
--query "パスワードをリセットするにはどうすればよいですか" \
--トレーニング
結果
意味
重大度
確認済み
毒殺された医師が1位、被害者はカナリアでシンクに電話
シンク特権に関連付けられている
部分的
ドクは救出されたが、カナリアはシンクに着地しなかった
高
NOT_RETRIEVED
毒を盛られたドクターはデコイを上回らなかった（レトリーバーの失敗）
情報
NOT_STORED
インデックス作成に失敗しました (セットアップの問題)
情報
取得失敗と服従失敗の分割は、RAG バリアントが追加する重要な実用的な区別です。 docs/poisoning.md#rag--vector-store-poisoning-v11 を参照してください。
10. 会話スレッドポイズニング (v1.

4)
3 番目の最も微妙な中毒の亜種。攻撃者は、一見無害に見える内容を共有会話スレッドに書き込みます。被害者エージェントは後でスレッドを再開し、毒されたターンを信頼できる以前のコンテキストの一部として扱います。取得したコンテンツをフィルタリングする防御機能ではこれを捕捉できません。毒は取得したデータではなく、エージェントの履歴に存在します。
エージェントスプロイト毒検証スレッド \
--sink-tool send_email --sink-arg body \
--sink-privilege egress \
--折り返し 3 \
--トレーニング
OpenAI アシスタント スレッド、マルチメッセージ状態のカスタマー サポート チャットボット、ユーザーがセッション コンテキストを共有するマルチテナント チャット プラットフォームに対して現実的です。 docs/poisoning.md#conversation-thread-poisoning-v14 を参照してください。
11. ライブ エンゲージメント ダッシュボード (v1.5 + v1.6)
エンゲージメント出力を参照し、ブラウザからライブ スキャン/検証を実行するためのローカル Web アプリ。 v1.5 には読み取り専用ブラウザが同梱されていました。 v1.6 では、ベアラー トークン認証、書き込みエンドポイント、サーバー送信イベント ストリーム、パス エクスプローラー ページ、およびライブ検出結果の更新が追加されました。
エージェントスプロイト サーブ --auth authorization.yaml
# -> http://127.0.0.1:8800
# 起動時にトークンが出力されます。それを /login ページに貼り付けます。
Cytoscape でレンダリングされたパーミッション グラフ、p を含む重大度フィルターされた所見テーブル

[切り捨てられた]

## Original Extract

Offensive security framework for AI agents and MCP servers. - agentsploit/agentsploit

GitHub - agentsploit/agentsploit: Offensive security framework for AI agents and MCP servers. · GitHub
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
agentsploit
/
agentsploit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
24 Commits 24 Commits .github/ workflows .github/ workflows docs docs examples examples src/ agentsploit src/ agentsploit tests tests ui ui .gitignore .gitignore .python-version .python-version AUTHORIZATION.md AUTHORIZATION.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Offensive security framework for AI agents and MCP servers.
AgentSploit is a Burp Suite / Metasploit-style framework purpose-built for the agentic AI attack surface. It helps red teamers, AI security researchers, and product security teams probe LLM agents and Model Context Protocol (MCP) servers for vulnerabilities that legacy tooling cannot find.
AgentSploit is an authorized-use security testing tool. You must have explicit written permission to scan any target you do not own. See AUTHORIZATION.md before running anything.
👉 New here? Start with docs/getting-started.md - a 10-minute tour of every capability against bundled fixtures, no API keys needed.
Every Fortune 500 is shipping LLM agents and MCP servers in 2026. The attack surface is genuinely new:
Tool descriptions are LLM-readable instructions - malicious ones can hijack agents.
Agents fetch untrusted content from PDFs, web pages, calendar invites, tickets - and that content can issue commands.
Chained tool calls create privilege escalation paths that no traditional permission model captures.
Memory and context windows can be poisoned across sessions.
Existing scanners (Burp, ZAP, Semgrep, Snyk) don't speak this layer. AgentSploit does.
pip install agentsploit
# Scaffold an engagement
agentsploit init my-engagement/ --authorized-by " Jane Doe <ciso@example.com> "
cd my-engagement/
# Scan an MCP server (training mode = no API keys needed, bundled fixtures)
agentsploit scan mcp stdio://./tests/fixtures/vulnerable_mcp/server.py --training
# Browse results in your browser (live engagement dashboard, v1.6+)
agentsploit serve --training
# -> http://127.0.0.1:8800 (token printed on startup)
Capabilities
AgentSploit ships eleven modules covering the agentic attack surface. Each is documented end-to-end with bundled vulnerable fixtures so you can run every one against a local target with --training and no API keys.
Connects to an MCP server over stdio , Streamable HTTP , or SSE and runs a battery of checks against its tools, resources, prompts, and (for HTTP/SSE) its HTTP surface.
Inventory checks (all transports):
2. Indirect Prompt Injection Payload Generator
Generates labeled payloads for testing whether an agent processes untrusted content safely.
direct - straightforward override attempts
role_confusion - fake system: / assistant: turns
delimiter - fenced-content escape and re-context
unicode_tag - invisible Unicode tag-block smuggling (U+E0000 range)
tool_smuggling - hidden tool-call invocations in narrative text
markdown - README/comment-style
html - page content with hidden elements
pdf - visible + hidden-layer PDF
email - RFC 5322 with HTML body and headers
ical - .ics calendar invite with malicious DESCRIPTION
Every payload is tagged with a canary string so you can detect successful injection in logs.
Takes a generated payload + an agent config and drives the payload through a real LLM. If the canary surfaces in the agent's response, tool calls, or thinking trace, the injection is confirmed exploitable - a CRITICAL/HIGH finding with the full trace persisted for audit.
Adapters (v0.9): anthropic (real Claude tool-use), openai (Chat Completions with tool use), http (generic HTTP agent with OpenAI-shaped contract - subclass for custom shapes), mock (deterministic, for tests). See docs/runner.md .
4. Permission Graph Mapper (v0.4)
Enumerates tools across multiple MCP servers, classifies each by privilege (source / pivot / sink), infers data-flow edges, and finds attack paths from low-trust sources to high-impact sinks. BloodHound for tool chains.
agentsploit map build --targets ./examples/map-targets.yaml --auth ./auth.yaml
agentsploit map export --graph ./engagements/ < id > / < sid > /permission_graph.json -f mermaid -o graph.md
Sink privilege
Severity
EXECUTION ( run_command , eval )
CRITICAL
MUTATION ( git_push , delete_* )
HIGH
EGRESS ( send_email , webhook )
HIGH
A path finding is a testable hypothesis - pair it with the v0.5 verifier to confirm exploitability end-to-end. See docs/mapper.md .
Closes the loop on the mapper. Take any mapper-inferred path, drive a path-targeted payload through a real or mock agent, and prove whether the chain actually completes.
agentsploit verify path \
--graph ./engagements/ < id > / < sid > /permission_graph.json \
--from read_file --to send_email \
--training # or --agent ./agent-anthropic.yaml --auth ./auth.yaml
Outcome
Meaning
Severity
CONFIRMED
Sink tool was called with the canary in its arguments
Tied to sink privilege (EXEC → CRITICAL)
PARTIAL
Sink reached or canary surfaced elsewhere, but chain incomplete
HIGH
FAILED
No canary surface anywhere
INFO
A CONFIRMED finding moves a mapper hypothesis from "plausible attack path" to "proven exploit." See docs/verifier.md .
6. Batch Path Verification (v0.6)
Drive the verifier across every path in a graph in one command - typical workflow is to triage against the cheap mock agent first, then re-run only the confirmations against the real model.
# Cheap triage pass (free, instant)
agentsploit verify all-paths --graph ./.../permission_graph.json --training
# Real-model pass on the same graph
agentsploit verify all-paths --graph ./.../permission_graph.json \
--agent ./agent-anthropic.yaml --auth ./auth.yaml \
--parallel 3 --max-paths 20
Deduplicates by (source, sink) pair, parallelises with rate-limit-aware concurrency, isolates per-path errors, and emits an aggregate batch_summary finding with the confirmation-rate percentage. See docs/verifier.md .
The default verifier uses one injection envelope ( role_confusion ). v0.7 adds four more - direct , delimiter , unicode_tag , tool_smuggling - and a fuzzer that tries them in sequence until one lands. Knowing which envelope wins tells defenders what their injection filter missed.
# Single-path fuzz
agentsploit verify fuzz-path --graph ./.../permission_graph.json \
--from read_file --to send_email --training
# Batch fuzz - every path × every technique, with early termination per path
agentsploit verify all-paths --graph ./.../permission_graph.json \
--fuzz --techniques role_confusion,delimiter,unicode_tag \
--parallel 3 --training
Technique
Defender takeaway when this lands
role_confusion
Chat-template filter doesn't catch fake <system> turns
delimiter
Untrusted content boundaries aren't enforced
unicode_tag
Defence strips printable ASCII but not U+E0000 tag block
tool_smuggling
Agent runtime parses JSON tool-call syntax out of narrative text
direct
No prompt-injection defence in place at all
See docs/verifier.md .
The first multi-phase attack module. Attacker plants a crafted note in shared agent storage; a separate victim agent run retrieves the note and is steered into invoking a sink tool with the attacker's canary. The remediation pattern this catches: agents that treat retrieved storage content as instructions, not data.
# Verify a memory-poisoning attack against the mock agent (free, instant)
agentsploit poison verify \
--sink-tool send_email --sink-arg body \
--sink-privilege egress \
--training
# Against real Claude
agentsploit poison verify \
--sink-tool send_email --sink-arg body \
--agent ./agent-anthropic.yaml --auth ./auth.yaml
Outcome
Meaning
Severity
CONFIRMED
Victim called sink with canary in args
Tied to sink privilege (EXEC → CRITICAL)
PARTIAL
Note retrieved but canary didn't surface in sink
HIGH
NOT_RETRIEVED
Note stored but victim never read it
INFO
NOT_STORED
Attacker write failed
INFO
See docs/poisoning.md .
9. RAG / Vector-Store Poisoning (v1.1)
The same threat model as v0.8 but lifted to the retrieval-augmented-generation case. Attacker doesn't know the victim's query key, only the topic . Indexes a poisoned doc that outranks legitimate corpus content for the target query; victim agent runs semantic_search , retrieves the top-ranked match, follows the embedded chain.
agentsploit poison verify-rag \
--sink-tool send_email --sink-arg body \
--sink-privilege egress \
--query " how do I reset my password " \
--training
Outcome
Meaning
Severity
CONFIRMED
Poisoned doc ranked first, victim called sink with canary
Tied to sink privilege
PARTIAL
Doc retrieved but canary didn't land in sink
HIGH
NOT_RETRIEVED
Poisoned doc didn't outrank decoys (retriever failure)
INFO
NOT_STORED
Indexing failed (setup issue)
INFO
The retrieval-failure vs obedience-failure split is the key actionable distinction the RAG variant adds. See docs/poisoning.md#rag--vector-store-poisoning-v11 .
10. Conversation-Thread Poisoning (v1.4)
The third and most subtle poisoning variant. The attacker writes a benign-looking turn into a shared conversation thread; the victim agent resumes the thread later and treats the poisoned turn as part of its own trusted prior context. Defences that filter retrieved content don't catch this - the poison sits in the agent's history, not in retrieved data.
agentsploit poison verify-thread \
--sink-tool send_email --sink-arg body \
--sink-privilege egress \
--turns-back 3 \
--training
Realistic against OpenAI Assistants threads, customer-support chatbots with multi-message state, and multi-tenant chat platforms where users share session context. See docs/poisoning.md#conversation-thread-poisoning-v14 .
11. Live engagement dashboard (v1.5 + v1.6)
A local web app for browsing engagement output AND driving live scans / verifies from the browser. v1.5 shipped the read-only browser; v1.6 added bearer-token auth, write endpoints, a Server-Sent Events stream, a path explorer page, and live findings updates.
agentsploit serve --auth authorization.yaml
# -> http://127.0.0.1:8800
# Token printed on startup; paste it into the /login page.
Cytoscape-rendered permission graphs, severity-filtered findings tables with p

[truncated]
