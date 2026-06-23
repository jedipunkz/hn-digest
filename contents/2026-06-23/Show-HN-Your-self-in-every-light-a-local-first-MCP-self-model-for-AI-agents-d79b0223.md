---
source: "https://github.com/almakit/alma"
hn_url: "https://news.ycombinator.com/item?id=48647761"
title: "Show HN: Your self, in every light - a local-first MCP self model for AI agents"
article_title: "GitHub - almakit/alma: Your self, in every light - a local-first MCP self model for AI agents · GitHub"
author: "0set0set"
captured_at: "2026-06-23T17:35:49Z"
capture_tool: "hn-digest"
hn_id: 48647761
score: 4
comments: 0
posted_at: "2026-06-23T16:49:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Your self, in every light - a local-first MCP self model for AI agents

- HN: [48647761](https://news.ycombinator.com/item?id=48647761)
- Source: [github.com](https://github.com/almakit/alma)
- Score: 4
- Comments: 0
- Posted: 2026-06-23T16:49:55Z

## Translation

タイトル: Show HN: あらゆる光の中でのあなたの自己 - AI エージェントのためのローカルファースト MCP 自己モデル
記事のタイトル: GitHub - almakit/alma: Your self, in each light - a local-first MCP self model for AI Agents · GitHub
説明: あらゆる観点から見たあなたの自己 - AI エージェントのためのローカルファースト MCP 自己モデル - almakit/alma

記事本文:
GitHub - almakit/alma: あらゆる観点から見たあなたの自己 - AI エージェントのためのローカルファースト MCP 自己モデル · GitHub
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
アルマキト
/
アルマ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
本支店

s タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows準拠/ 2026-06準拠/ 2026-06 crates crates schemas/ 2026-06 schemas/ 2026-06 .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml Justfile Justfile LICENSEライセンス Makefile Makefile 通知 通知 README.md README.md Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
アルマは AI エージェントにあなたの記憶を与えます。
これは、繰り返し続ける事実や好みのためのローカルファーストの MCP サーバーです。
新しいチャット間: あなたの名前、役割、働き方、回答の好み、現在の
コンテキスト、原則、価値観、その他の自己モデルの部分。
目標はシンプルです。エージェントは罠に陥ることなくあなたのことを理解できるようにする必要があります。
1 つのベンダー アカウント内で、永続的な書き込みを行わずに理解できる
自分の個人的なコンテキストへのアクセス。
Almaはあなたのマシン上にデータを保持し、あなたが承認した部分のみを公開し、
すべての永続的な変更は監査可能です。
ステータス: 実験的な趣味のプロジェクト。 API と動作は変更される可能性があります。
AI エージェントは長期にわたる協力者になりつつありますが、彼らの記憶はまだ残っています
断片化された:
各ツールはあなたの個別のバージョンを学習します。
新しいチャットはゼロから始まることがよくあります。
ベンダーのメモリは、ツール間で検査、移動、共有することが困難です。
エージェントに長期記憶への直接書き込みをさせるのは過剰な信頼です。
Almaは、あなたとあなたが使用するエージェントの間の小さなローカルレイヤーです。エージェントは尋ねることができます
Almaは知ることが許可されているものを返します。Almaは、
店内いっぱい。
Almaは自己モデルをファセットと呼ばれるファクトとして保存します。ファセットは特定の値です。
person.display_name やワークスタイルの好みなどのディメンション。ファセットキャリー
ステータス、信頼性、情報源、証拠。
真実の情報源は、

追加専用のイベント ログ。現在の状態は、
ログ。変更を検査可能に保ち、再生によって元に戻せるようにします。
エージェントは事実を直接編集しません。弱い信号を観察し、記録することができます。
明示的に許可されている場合は証拠に裏付けられた行動、または、
承認する人。永続書き込みには人間の承認トークンが必要です。
アルマはエージェントを信頼できないものとして扱い、地元の人間を権威あるものとして扱います。
読み取りの範囲は許可と目的によって決まります。
デフォルトのエージェント サーフェスは最小権限です。
機密レイヤーとレンズの名前空間はオプトインです。
管理アクションは、CLI、コンパニオン、または明示的な管理サーバー モードで実行されます。
実際の結果: エージェントはあなたに適応できますが、静かに引き継ぐことはできません
あなたの記憶。
エージェントは公式を使用して MCP 経由で Alma と会話します。
Rust rmcp SDK を stdio 上で使用します。
主要な読み取りツールは alma_get_reading です。エージェントはユーザーの質問をAlmaに渡します
読み取りフォーカス内の関連するディメンションをランク付けします。エージェントが正確なキーを必要とする場合、
推測する代わりに、最初に list_dimension を呼び出します。
たとえば、あなたの名前に関する質問は person.display_name に解決されるはずです。
identity.name のような作成されたキーではありません。エージェントが未知の正規版を要求した場合
次元の場合、Almaは提案とともにUNKNOWN_DIMENSIONを返します。
書き込みの場合、通常のパスは提案と承認です。
エージェントは alma_propose_facet を呼び出します。
アルマはその人に確認の質問を返します。
ユーザーは CLI またはコンパニオンから承認します。
エージェントは alma_record_facet を使用してワンタイム トークンを引き換えます。
承認トークンも永続書き込みもありません。
Almaは現在、3つのローカルバイナリを構築しています。
alma-server : エージェントが接続する MCP サーバー。
alma : セットアップ、許可、提案、エクスポート、リセット用の CLI。
alma-companion : 自己モデルをレビューおよび管理するためのローカル TUI。
3 つすべてが ~/.alma/alma.db にある同じストアを使用します。よ

でオーバーライドできます
アルマデータベース 。
パッケージ化されたリリースはまだありません。 Cargo を使用してソースからビルドします。
CDアルマ
カーゴビルド --release
次に、ガイド付きセットアップを実行します。
./target/release/alma クイックスタート
または、ホストを手動で接続します。
./target/release/alma 接続カーソル --apply --global
./target/release/alma connect クロード --apply
ホストをリロードすると、alma MCP サーバーが表示されます。
alma-companion は、JSON を編集したくない人のためのローカル UI です。開きます
サーバーおよび CLI と同じストア。
エージェントが受け取る読み取り値をプレビューします。
助成金、リクエスト、監査イベントを確認します。
提案された思い出を承認または拒否します。
アルマ・コンパニオン
ALMA_SEED=./my-self.json アルマコンパニオン
携帯性
あなたのセルフモデルを 1 つのアプリに閉じ込めるべきではありません。
alma_export_bundle は、完全なイベント ログを署名されたコンパクトな JWS としてエクスポートします。
Ed25519。公開キーはヘッダーに埋め込まれているため、バンドルを検証できます
事前のキー交換なしで。インポートはイベントを新しいストアにリプレイします。
改ざんすると署名が無効になります。
ローカル ストアを消去して、最初からやり直すことができます。
alma delete # ローカルデータを削除する前に DELETE を要求します
alma delete --yes # 非対話型
alma delete --keep-key # 署名キーを保持し、ストアを消去します
削除は人間のみが行う操作です。 MCP ツールとしては公開されていません。
ワイヤーコントラクトはリポジトリ内に存在します。
schemas/2026-06/ : プロトコル アーティファクトの JSON スキーマ。
conformance/2026-06/ : ランニングの適合ケース
実装。
プロジェクトは固定された Rust ツールチェーンを使用します。 just はオプションですが、次の場合に推奨されます
ローカルチェック。
ただfmt-checkしてください
ただの糸くず
ただテストしてください
ただの適合
ただの取材
ただチェックしてください
CI は、ビルド、テスト、lint、適合性、カバレッジ、カーゴ監査を実行します。リリース
自動化は意図的にまだ有効になっていません。ビルドはローカルのみです。
あらゆる観点から見た自分 - ローカルファーストの MCP セルフ

AIエージェントのモデル
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
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

Your self, in every light - a local-first MCP self model for AI agents - almakit/alma

GitHub - almakit/alma: Your self, in every light - a local-first MCP self model for AI agents · GitHub
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
almakit
/
alma
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows conformance/ 2026-06 conformance/ 2026-06 crates crates schemas/ 2026-06 schemas/ 2026-06 .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml Justfile Justfile LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Alma gives AI agents a memory that belongs to you.
It is a local-first MCP server for the facts and preferences you keep repeating
across new chats: your name, role, working style, answer preferences, current
context, principles, values, and other parts of your self model.
The goal is simple: an agent should be able to understand you without trapping
that understanding inside one vendor account, and without getting permanent write
access to your personal context.
Alma keeps the data on your machine, exposes only the parts you approve, and makes
every durable change auditable.
Status: experimental hobby project. APIs and behavior may change.
AI agents are becoming long-running collaborators, but their memory is still
fragmented:
Each tool learns a separate version of you.
New chats often start from zero.
Vendor memory is hard to inspect, move, or share across tools.
Letting an agent write directly to long-term memory is too much trust.
Alma is a small local layer between you and the agents you use. The agent can ask
Alma what it is allowed to know, and Alma returns a scoped Reading instead of the
full store.
Alma stores a self model as facts called facets. A facet is a value for a specific
dimension, such as person.display_name or a work-style preference. Facets carry
status, confidence, source, and evidence.
The source of truth is an append-only event log. Current state is rebuilt from the
log, which keeps changes inspectable and reversible by replay.
Agents do not edit facts directly. They can observe weak signals, record
evidence-backed behavior when explicitly granted, or propose a new facet for the
person to approve. Durable writes require a human approval token.
Alma treats the agent as untrusted and the local human as the authority.
Reads are scoped by grant and purpose.
The default agent surface is least privilege.
Sensitive layers and lens namespaces are opt-in.
Admin actions live in the CLI, Companion, or an explicit admin server mode.
The practical result: an agent can adapt to you, but it cannot quietly take over
your memory.
Agents talk to Alma over MCP , using the official
Rust rmcp SDK over stdio.
The main read tool is alma_get_reading . Agents pass the user's question and Alma
ranks the relevant dimensions in the Reading focus. If an agent needs exact keys,
it calls list_dimensions first instead of guessing.
For example, a question about your name should resolve to person.display_name ,
not a made-up key like identity.name . If an agent asks for an unknown canonical
dimension, Alma returns UNKNOWN_DIMENSION with suggestions.
For writes, the normal path is propose and approve:
The agent calls alma_propose_facet .
Alma returns a confirmation question for the person.
The person approves from the CLI or Companion.
The agent redeems the one-time token with alma_record_facet .
No approval token, no durable write.
Alma currently builds three local binaries:
alma-server : the MCP server agents connect to.
alma : the CLI for setup, grants, proposals, export, and reset.
alma-companion : a local TUI for reviewing and managing your self model.
All three use the same store at ~/.alma/alma.db . You can override it with
ALMA_DB .
There are no packaged releases yet. Build from source with Cargo.
cd alma
cargo build --release
Then run the guided setup:
./target/release/alma quickstart
Or connect a host manually:
./target/release/alma connect cursor --apply --global
./target/release/alma connect claude --apply
Reload the host and the alma MCP server should appear.
alma-companion is the local UI for people who do not want to edit JSON. It opens
the same store as the server and CLI.
Preview the Reading an agent would receive.
Review grants, requests, and audit events.
Approve or deny proposed memories.
alma-companion
ALMA_SEED=./my-self.json alma-companion
Portability
Your self model should not be trapped in one app.
alma_export_bundle exports the full event log as a compact JWS signed with
Ed25519. The public key is embedded in the header, so the bundle can be verified
without a prior key exchange. Import replays the events into a fresh store, and
tampering invalidates the signature.
You can erase the local store and start over.
alma delete # asks for DELETE before removing local data
alma delete --yes # non-interactive
alma delete --keep-key # keep the signing key, wipe the store
Deletion is a human-only action. It is not exposed as an MCP tool.
The wire contract lives in the repo:
schemas/2026-06/ : JSON Schemas for protocol artifacts.
conformance/2026-06/ : conformance cases for running
implementations.
The project uses the pinned Rust toolchain. just is optional but recommended for
local checks.
just fmt-check
just lint
just test
just conformance
just coverage
just check
CI runs build, tests, lint, conformance, coverage, and cargo audit . Release
automation is intentionally not enabled yet; builds are local-only.
Your self, in every light - a local-first MCP self model for AI agents
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
