---
source: "https://github.com/las7/Guarden"
hn_url: "https://news.ycombinator.com/item?id=48450231"
title: "Show HN: Guarden – Authorization for AI agent actions powered by OPA"
article_title: "GitHub - las7/Guarden · GitHub"
author: "sakuraiben"
captured_at: "2026-06-08T21:39:18Z"
capture_tool: "hn-digest"
hn_id: 48450231
score: 1
comments: 0
posted_at: "2026-06-08T19:20:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Guarden – Authorization for AI agent actions powered by OPA

- HN: [48450231](https://news.ycombinator.com/item?id=48450231)
- Source: [github.com](https://github.com/las7/Guarden)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T19:20:01Z

## Translation

タイトル: Show HN: Guarden – OPA による AI エージェント アクションの承認
記事タイトル: GitHub - las7/Guarden · GitHub
説明: GitHub でアカウントを作成して、las7/Guarden の開発に貢献します。

記事本文:
GitHub - las7/Guarden · GitHub
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
ラス7
/
ガーデン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミ

ts 11 コミット .github .github エージェント エージェント ポリシー ポリシー スクリプト スクリプト src src テスト テスト .dockerignore .dockerignore .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md DEMO.md DEMO.md Dockerfile Dockerfileライセンス ライセンス Makefile Makefile README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml 要件-dev.txt 要件-dev.txt 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント アクションの承認。SDK ではなくアクション境界で強制されます。
焦点を絞った概念実証。テストとレッドチームを使用して、以下のパターンをエンドツーエンドで実装します。セキュリティ モデルと SECURITY.md は、脅威のモデルと範囲をカバーしています。
実際のアクション (メールの送信、送金、チケットのオープンなど) を実行するエージェントを実行する場合、リスクが集中するのは承認です。モデルは何を行うかを選択します。あなたの仕事は、許可されていることを制限することです。
同じアクションが、直接ツール呼び出し、関数呼び出しループ、モデルが記述して実行するコード (コード インタプリタ、コンピュータ使用ループ、自律型コーディング エージェント) などのいくつかのパスを介して副作用に到達します。 SDK またはツール ラッパーに関連付けられた承認は最初の 2 つを処理し、3 つ目は処理されません。モデルがコードを出力すると、ゲートするためのツール呼び出しは残りません。
ガーデンは決定をトランスポートからアクションに移した。
すべてのパスは 1 つの Action に正規化されます。ツール スキーマから生成された 1 つのポリシーは、許可または拒否を返します。アクションが SDK アダプターからのものであっても、サンドボックス内のエージェントが作成したコードからのものであっても、判定は同じです。
SDK アダプターのコード実行アダプターの呼び出し元
\/
vv
ノーマライザー -> 1 つの正規形状のアクション
|
ポリシー エンジン (OPA) 1 つの生成されたポリシー
|
ブローカー -> ツールプロキシは許可のみ
|
台帳（JSONL）の追加

d のみの監査証跡
2 つのメカニズムが設計を担っています。
単一の施行ポイント。ブローカーはツールへの唯一のパスです。アクションを正規化し、OPA をクエリし、許可された場合にのみプロキシを実行します。フェイルクローズします。ポリシー エンジンに到達できない場合、または決定が未定義の場合、アクションは拒否されます。
アンビエント資格情報ではなく、範囲指定された機能。サンドボックス化されたコードは、実際のダウンストリーム資格情報を保持することはありません。有効期間が短く、送信者制限のあるトークンと、ブローカーにバインドされたクライアントを保持します。トークンのスコープは境界であるため、個々の呼び出しをインターセプトするものはありません。これがコード実行パスをカバーするものです。
これは、エージェントが実際に影響を与える場合に適用される最小権限とコードとしてのポリシーです。おなじみのプリミティブが右のレイヤーに移動されました。
Docker、他に何も必要ありません:
git clone https://github.com/las7/Guarden.git && cd Guarden
docker-demo を作成する
ネイティブ、Python 3.11+:
セットアップを行う
全部作る
ターゲットを指定せずに make すると、すべてがリストされます。各エントリポイントは独自の OPA とブローカーを実行し、それらを破棄します。
make Demon は 3 つのケースを実行します。正当な行為は認められます。外部受信者への同じ email:send は拒否されます。これは、OAuth スコープでは表現できない引数レベルの代理決定です。エージェントがツールを呼び出す代わりにコードを作成する場合、同一のポリシーは同一のアクションを拒否します。出力と台帳スキーマについては、DEMO.md を参照してください。
コンポーネント
ソース
役割
ノーマライザー
ノーマライザー.py
すべてのインターフェースを 1 つのアクションにマップします
能力トークン
能力.py
有効期限が短く、HMAC 署名があり、送信者制限があります。組み込みのデフォルトのシークレットはありません
ポリシーエンジン
エンジン.py
OPA を介してアクションを評価します。失敗して閉じられる
ポリシーの生成
from_schema.py
ツールスキーマから生成された最小特権 Rego
ブローカー
サーバー.py
単一の施行ポイント。プロキシは許可のみ
アダプター
sdk.py 、 codeexec.py
2つのインターフェース
元帳
元帳.py

追加専用の JSONL 監査証跡
ツールの作成者は、機能とオプションの引数制約を宣言します。最小特権ポリシーはそのスキーマから生成されます。手書きのレゴはありません。
Guarden は、最小特権スコープ、引数レベルおよび代理の制約、送信者制約機能、フェイルクローズ評価、および追加専用台帳を強制します。
完全な脅威モデル、現在の範囲、開示プロセスは SECURITY.md に文書化されています。
Guarden はエンドツーエンドで次のことを実証します。
SDK およびコード実行アダプター全体にわたるインターフェイスに依存しない適用
ツールスキーマから生成された最小権限ポリシー
議論レベルの、認識上の意思決定
送信者制約の短い TTL 機能
追加専用の意思決定台帳によるフェイルクローズ評価
運用環境の展開は、発信者認証 (mTLS またはクライアント証明書)、実際のツールの統合、TLS および KMS を利用したキー管理、OS レベルのサンドボックス分離、および改ざん防止台帳を使用してこれに基づいて構築されます。 SECURITY.md を参照してください。
セットアップ、チェック ( make test lint Audit )、およびプルリクエストのプロセスについては、CONTRIBUTING.md を参照してください。参加には行動規範が適用されます。
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to las7/Guarden development by creating an account on GitHub.

GitHub - las7/Guarden · GitHub
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
las7
/
Guarden
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits .github .github agents agents policy policy scripts scripts src src tests tests .dockerignore .dockerignore .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DEMO.md DEMO.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt View all files Repository files navigation
Authorization for AI agent actions, enforced at the action boundary rather than the SDK.
A focused proof of concept. It implements the pattern below end to end, with tests and a red-team. The security model and SECURITY.md cover the threat model and scope.
If you run agents that take real actions (sending mail, moving money, opening tickets), authorization is where the risk concentrates. The model chooses what to do; your job is to bound what it is permitted to do.
The same action reaches the side effect through several paths: a direct tool call, a function-calling loop, and code the model writes and runs (code interpreters, computer-use loops, autonomous coding agents). Authorization attached to the SDK or the tool wrapper handles the first two and misses the third. Once the model emits code, there is no tool call left to gate.
Guarden moves the decision off the transport and onto the action.
Every path normalizes to one Action . A single policy, generated from your tool schemas, returns allow or deny. The verdict is identical whether the action came from the SDK adapter or from agent-written code in a sandbox.
SDK adapter code-exec adapter callers
\ /
v v
normalizer -> Action one canonical shape
|
policy engine (OPA) one generated policy
|
broker -> tool proxies only on allow
|
ledger (JSONL) append-only audit trail
Two mechanisms carry the design:
A single enforcement point. The broker is the only path to a tool. It normalizes the action, queries OPA, and proxies only on allow. It fails closed: if the policy engine is unreachable or the decision is undefined, the action is denied.
Scoped capabilities, not ambient credentials. Sandboxed code never holds a real downstream credential. It holds a short-lived, sender-constrained token and a client bound to the broker. The token's scope is the boundary, so nothing has to intercept individual calls. That is what covers the code-execution path.
This is least privilege and policy-as-code applied where agents actually cause effects. Familiar primitives, moved to the right layer.
Docker, nothing else required:
git clone https://github.com/las7/Guarden.git && cd Guarden
make docker-demo
Native, Python 3.11+:
make setup
make all
make with no target lists everything. Each entrypoint runs its own OPA and broker and tears them down.
make demo runs three cases. A legitimate action is allowed. The same email:send to an external recipient is denied, an argument-level, on-behalf-of decision that OAuth scope cannot express. The identical policy denies the identical action when the agent writes code instead of calling a tool. See DEMO.md for the output and the ledger schema.
Component
Source
Role
Normalizer
normalizer.py
maps every interface to one Action
Capability tokens
capability.py
short-lived, HMAC-signed, sender-constrained; no built-in default secret
Policy engine
engine.py
evaluates the Action via OPA; fails closed
Policy generation
from_schema.py
least-privilege Rego generated from tool schemas
Broker
server.py
single enforcement point; proxies only on allow
Adapters
sdk.py , codeexec.py
the two interfaces
Ledger
ledger.py
append-only JSONL audit trail
A tool author declares a capability and an optional argument constraint. The least-privilege policy is generated from that schema. No hand-written Rego.
Guarden enforces least-privilege scope, argument-level and on-behalf-of constraints, sender-constrained capabilities, fail-closed evaluation, and an append-only ledger.
The full threat model, current scope, and disclosure process are documented in SECURITY.md .
Guarden demonstrates, end to end:
Interface-agnostic enforcement across the SDK and code-execution adapters
Least-privilege policy generated from tool schemas
Argument-level, on-behalf-of-aware decisions
Sender-constrained, short-TTL capabilities
Fail-closed evaluation with an append-only decision ledger
A production deployment builds on this with caller authentication (mTLS or client certificates), real tool integrations, TLS and KMS-backed key management, OS-level sandbox isolation, and a tamper-evident ledger. See SECURITY.md .
See CONTRIBUTING.md for setup, the checks ( make test lint audit ), and the pull-request process. Participation is subject to the Code of Conduct .
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
