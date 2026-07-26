---
source: "https://github.com/eouzoe/reel"
hn_url: "https://news.ycombinator.com/item?id=49055791"
title: "A protocol for AI agent workspace state and effect management"
article_title: "GitHub - eouzoe/reel: A protocol for AI agent workspace state and effect management. · GitHub"
author: "eouzoe"
captured_at: "2026-07-26T08:08:14Z"
capture_tool: "hn-digest"
hn_id: 49055791
score: 1
comments: 0
posted_at: "2026-07-26T08:02:03Z"
tags:
  - hacker-news
  - translated
---

# A protocol for AI agent workspace state and effect management

- HN: [49055791](https://news.ycombinator.com/item?id=49055791)
- Source: [github.com](https://github.com/eouzoe/reel)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T08:02:03Z

## Translation

タイトル: AI エージェントのワークスペースの状態と影響を管理するためのプロトコル
記事タイトル: GitHub - eouzoe/reel: AI エージェント ワークスペースの状態と効果管理のためのプロトコル。 · GitHub
説明: AI エージェント ワークスペースの状態と影響を管理するためのプロトコル。 - オイゾエ/リール

記事本文:
GitHub - eouzoe/reel: AI エージェント ワークスペースの状態と効果を管理するためのプロトコル。 · GitHub
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
コードの品質 マージ時に品質を強制する
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
おうぞえ
/
リール
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

オプションについて
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット crates crates spec spec .clippy.toml .clippy.toml .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md Deny.toml Deny.toml Rust-toolchain.toml Rust-toolchain.toml Rustfmt.toml Rustfmt.toml 表示すべてのファイル リポジトリ ファイルのナビゲーション
リールは、AI エージェントの投機的な作業と、
自社が所有していないシステム。
エージェントは物事を試みることによって推論します。一部の試みはその境界を越えます。
メッセージが投稿され、行が書き込まれ、支払いが行われます。アクションが完了すると、
クロス、その後の決定はなし - ラインを放棄する決定を含む
それを生み出した推論は、それを撤回することができます。リールは境界線を作るために存在する
トランザクションであるため、システムが破棄することを決定したアクションは決して破棄されません
最初に実行されました。
リールは、6 つのコンテンツアドレス指定型 - Hash 、 Block 、 Ref 、 Delta 、
Capability 、 View — 3 つの動詞 — fork 、 commit 、 abort — と 3 つ
不変条件: 到達可能性、破棄された不可逆効果の沈黙、および
能力の狭まり。
負荷に耐える決定は、 View の保留中の副作用自体が影響を受けるということです。
コンテンツアドレス指定された Block 。 commit は、そのブロックをコミットされたログに移動します
そしてそれを排出します。 abort は参照を削除し、その後ブロックに到達できなくなります。
そして集められます。不可逆的な効果の提出と、
コミットを起動するために必要な機能が構築されているため、コミットによって起動できます。
commit のドレイン内のみ。 abort は何も構築しないため、破棄パス上にあります
構造上、エフェクトは実行できません。
動詞、 View のスナップショット分離読み取り、機能モデル、
効果分類法はそれぞれ存在するものから取得されます

NG 作品 — Gray (1981)、Wang &
Zheng (2026)、ベレンソンら。 (1995)、ミラー、イー、シャピロ (2003)、
ガルシア＝モリーナとセイラム（1987）。リールの貢献は、以下の条件下での構成です。
単一のコンテンツアドレス表現。
単一のツールではない理由
状態、副作用、機能は同じ種類のものとして表現されます。
コンテンツアドレスブロック。したがって、1 つのトランザクション カーネルがロールバックを提供します。
それぞれを実装するのではなく、ワークスペース全体の監査と共有
機能ごとにもう一度説明します。多くの外部に作業をルーティングするシステム
リソースは、代わりにこの 1 つのカーネル上でその状態と効果の層を再構築できます。
リソースごとのトランザクション性についての推論。
仕様 ( spec/ ): 6 つのタイプ、3 つの動詞、3 つ
不変条件;エフェクトクラス。適合基準。アダプター
インターフェース。
Rust カーネル ( crates/ ):
リール仕様 — プロトコルのタイプ。名前空間は永続的であり、
コピーオンライト マップなので、フォークはそれを一定時間で共有します。
リールストア — コンテンツアドレス指定のブロックおよび参照ストア、インメモリ
バックエンドと redb-backed 永続的なバックエンド、および到達可能性ウォーカー
コレクション。
リールエフェクト — シールされたエフェクトクラスの特性とリニア
発火能力。
リールコア — カーネル。 fork (定時コピーオンライト継承)
機能の絞り込みあり) と中止が実装されます。コミットが続きます。
アダプタ/ 、リール-cli — ファイルシステムとリモートアダプタ、および
コマンド ライン サーフェス、進行中です。
ワークスペースは安定した Rust 上に構築されており、そのテスト スイートに合格しています。
カーネルの不変条件に対するプロパティベースのテスト。
commit — 前提条件の検証とバッファリングされたエフェクトの順序付けされたドレイン。
コマンドライン ツール — プロセスを実行し、そのプロセスを表示する予行ループ
保留中の効果をコミットまたは破棄します。

えっと。
エフェクトインターセプトレイヤー — プロセスのアウトバウンドエフェクトをキャプチャする
交換可能なバックエンドを使用するため、新しいサービスを分類することが唯一の
新しい統合に必要な作業。
カーゴビルド --workspace
貨物テスト --ワークスペース
Rust 1.95 (安定版)、エディション 2024。
プレリリース。仕様が決まりました。カーネルは構築中です
上で説明した道沿いにあります。インターフェースは変更される場合があります。
AI エージェントのワークスペースの状態と影響を管理するためのプロトコル。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A protocol for AI agent workspace state and effect management. - eouzoe/reel

GitHub - eouzoe/reel: A protocol for AI agent workspace state and effect management. · GitHub
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
Code Quality Enforce quality at merge
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
eouzoe
/
reel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits crates crates spec spec .clippy.toml .clippy.toml .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md deny.toml deny.toml rust-toolchain.toml rust-toolchain.toml rustfmt.toml rustfmt.toml View all files Repository files navigation
reel is a protocol for the boundary between an AI agent's speculative work and
the systems it does not own.
An agent reasons by attempting things. Some attempts cross that boundary: a
message is posted, a row is written, a payment is taken. Once an action has
crossed, no later decision — including the decision to abandon the line of
reasoning that produced it — can retract it. reel exists to make the boundary
transactional, so that an action the system decides to discard is never
performed in the first place.
reel defines six content-addressed types — Hash , Block , Ref , Delta ,
Capability , View — three verbs — fork , commit , abort — and three
invariants: reachability, the silencing of discarded irreversible effects, and
capability narrowing.
The load-bearing decision is that a View 's pending side-effects are themselves
a content-addressed Block . commit moves that block into the committed log
and drains it; abort drops the reference, after which the block is unreachable
and is collected. No path between the submission of an irreversible effect and
commit can fire it, because the capability needed to fire one is constructed
only inside commit 's drain. abort constructs none, so on the discard path
the effect is, by construction, unable to run.
The verbs, the snapshot-isolation reading of a View , the capability model and
the effect taxonomy are each taken from existing work — Gray (1981), Wang &
Zheng (2026), Berenson et al. (1995), Miller, Yee & Shapiro (2003),
Garcia-Molina & Salem (1987). reel's contribution is their composition under a
single content-addressed representation.
Why it is more than a single tool
State, side-effects and capabilities are represented as the same kind of
content-addressed block. One transactional kernel therefore provides rollback,
audit and sharing for an entire workspace, rather than each being implemented
again for each feature. A system that routes work across many external
resources can rebuild its state-and-effect layer on this one kernel instead of
reasoning about transactionality resource by resource.
The specification ( spec/ ): the six types, three verbs and three
invariants; the effect classes; the conformance criteria; the adapter
interface.
A Rust kernel ( crates/ ):
reel-spec — the protocol types. The namespace is a persistent,
copy-on-write map, so a fork shares it in constant time.
reel-store — a content-addressed block-and-ref store, with an in-memory
backend and a redb-backed persistent one, and a reachability walker for
collection.
reel-effects — the sealed effect-class traits and the linear
fire-capability.
reel-core — the kernel. fork (constant-time copy-on-write inheritance
with capability narrowing) and abort are implemented; commit follows.
adapters/ , reel-cli — the filesystem and remote adapters and the
command-line surface, in progress.
The workspace builds on stable Rust and passes its test suite, which includes
property-based tests over the kernel's invariants.
commit — precondition validation and the ordered drain of buffered effects.
The command-line tool — a dry-run loop that runs a process, displays its
pending effects, and either commits or discards them.
The effect-interception layer — capturing a process's outbound effects
through a replaceable backend, so that classifying a new service is the only
work a new integration requires.
cargo build --workspace
cargo test --workspace
Rust 1.95 (stable), edition 2024.
Pre-release. The specification is settled; the kernel is under construction
along the path described above. Interfaces may change.
A protocol for AI agent workspace state and effect management.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
