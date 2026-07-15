---
source: "https://github.com/grigoriitropin/vehir-platform"
hn_url: "https://news.ycombinator.com/item?id=48914954"
title: "Show HN: Vehir – a platform built for AI agents: compiler, microkernel, CAS"
article_title: "GitHub - grigoriitropin/vehir-platform: An AI-native, self-hosting platform for autonomous agents, with a user-space microkernel, worker runtime, native compiler, and content-addressed generations. · GitHub"
author: "dewdgi"
captured_at: "2026-07-15T01:23:37Z"
capture_tool: "hn-digest"
hn_id: 48914954
score: 2
comments: 0
posted_at: "2026-07-15T01:00:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Vehir – a platform built for AI agents: compiler, microkernel, CAS

- HN: [48914954](https://news.ycombinator.com/item?id=48914954)
- Source: [github.com](https://github.com/grigoriitropin/vehir-platform)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T01:00:50Z

## Translation

タイトル: Show HN: Vehir – AI エージェント用に構築されたプラットフォーム: コンパイラー、マイクロカーネル、CAS
記事のタイトル: GitHub - grigoriitropin/vehir-platform: ユーザー空間のマイクロカーネル、ワーカー ランタイム、ネイティブ コンパイラー、およびコンテンツ アドレス指定された世代を備えた、自律エージェント用の AI ネイティブのセルフホスティング プラットフォーム。 · GitHub
説明: ユーザー空間のマイクロカーネル、ワーカー ランタイム、ネイティブ コンパイラー、およびコンテンツ アドレス指定された世代を備えた、自律エージェント用の AI ネイティブのセルフホスティング プラットフォームです。 - グリゴリトロピン/車両プラットフォーム

記事本文:
GitHub - grigoriitropin/vehir-platform: ユーザー空間のマイクロカーネル、ワーカー ランタイム、ネイティブ コンパイラー、およびコンテンツ アドレス指定された世代を備えた、自律エージェント用の AI ネイティブのセルフホスティング プラットフォーム。 · GitHub
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
グリゴリトロピン
/
車両プラットフォーム
P

ユーブリック
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット ARCHITECTURE.md ARCHITECTURE.md INSTALL.md INSTALL.md LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Vehir Platform は、エージェントとコンピューターの対話を中心に設計された実験的な AI ネイティブ コンピューティング環境です。従来のシステムが人間とコンピューターの対話原理を使用して人々がコンピューティングにアクセスできるようにするのと同じように、Vehir は AI エージェントのニーズに基づいてソフトウェア環境全体を設計します。
そのインターフェイスとツールは、AI エージェントが直接使用できる構造化された検出可能な入出力を中心に設計されています。 Vehir 氏は、エージェント サポートを既存のワークフロー上の追加インターフェイスとしてのみ扱うのではなく、ソフトウェア環境自体を機械推論にとってより便利にする方法を模索しています。
現在の実装は、構造化エージェント ツール、IPM パッケージ、ネイティブ コンパイルとロード、コンテンツ アドレス指定された生成、調整、監視対象サービスなど、この環境の基礎を提供します。このプラットフォームは現在も活発に開発中です。
4 GB 以上の RAM を推奨
以下をコピーして AI エージェントに送信します。
完全なインストール ガイドをお読みください。
https://github.com/grigoriitropin/vehir-platform/blob/main/INSTALL.md
何らかの行動を起こす前に。最初から最後までフォローして、質問してください。
ホスト システムに関する選択や情報が必要なときはいつでも、ユーザーに連絡してください。
なぜ AI ネイティブなのか
Vehir はマシン間の対話のために構築されています。ラッピングの代わりに
AI インターフェイスを備えた既存の人間向けツールを使用して、あらゆるレイヤーを設計します
エージェントが直接使用できるようにするには:
構造化された入出力の削減

自由形式のテキスト解析に依存しています。
自己検出可能なツール — システムはエージェントに何ができるかを伝えます。
修復ヒントを含む機械可読エラー。
ソース言語 (IPM) は、手動による編集ではなく、機械による書き込みと検証用に設計されています。
宣言型ビルド — エージェントは望ましい状態を宣言します。システムが調和します。
一部の機能はまだ開発中であり、まだ完成していない可能性があります。
Vehir は単なるエージェント ツール サーバーではありません。いくつかのレイヤーを組み合わせて、
1 つのセルフホスティング プラットフォーム:
ユーザー空間のマイクロカーネルとワーカー ランタイム。カーネルが起動し、
アリーナ、共有メモリ、ドアベルベースを使用して従業員を管理します
コーディネート。マネージド サービスは同じワーカー モデルを使用します。
自己ホスト型のネイティブ コンパイラ。 IPM パッケージは直接コンパイルされます。
x86-64 ネイティブ コード。コンパイラーはフィックスポイント検証を使用して自動的にビルドします。
世代を超えて。
ネイティブ読み込み。 Vehir はオブジェクトのクロージャを解決し、再配置を適用し、
従来のコンポーネントに依存せずに、独自にコンパイルされたアーティファクトをロードします。
アプリケーションのランタイム。
コンテンツアドレス型ストレージ。ソースとコンパイルされたオブジェクトが扱われます
ダイジェストによって重複が除去され、不変の世代に記録されます。
宣言的和解。宣言されたビルド状態とパッケージのグラフ
検証された世代に変換され、その後アトミックに変換されます。
電流スイッチ。
コードから派生したエージェント ツール。 MCP の操作、契約のリクエスト、ヘルプ、および
リクエストの検証は、ツールの実装から派生します。
個別の手書きインターフェイスとして維持されます。
全宇宙的な強制。クロスパッケージ分析による構造チェック
世代が昇格する前のポリシー不変条件。
この GitHub リポジトリにはドキュメントのみが含まれています。リリースアーカイブは、
バイナリのみのディストリビューションではない: 公開されたすべての世代のアーカイブ
( N.tar.gz ) には完全な IPM ソースが含まれています

ネイティブと一緒に木を作る
正確な世代をブートストラップするために必要なアーティファクト。ソースは次のとおりです。
Vehir コンポーネントが開始される前に検査されます。
Vehir は実験的なものであり、現在開発中です。インターフェースと
形式はリリース間で変更される可能性があります。
INSTALL.md — AI エージェントのインストール ガイド
ARCHITECTURE.md — 技術アーキテクチャ
何かを取り外す前に、実際の Vehir の設置場所を特定してください
ユーザーに確認を求めます。ユーザー サービスを停止して削除します。
systemctl --user disable --now vehir-kernel.service
rm ~ /.config/systemd/user/vehir-kernel.service
systemctl --user デーモン-リロード
エージェント クライアントの構成から Vehir MCP エントリのみを削除し、
インストール中に追加された永続的な Vehir 命令は、すべてを保持します。
関係のない構成。
インストール ディレクトリは、完全な Vehir コンテンツ アドレス ストアです。
これには、すべての世代、コンパイルされたオブジェクト、およびソース オブジェクトが含まれます。その後
ユーザーは、この状態がどれも必要ないことを確認し、確認された絶対的な状態を削除します
設置場所:
rm -rf -- <確認済みの絶対インストール場所>
について
ユーザー空間のマイクロカーネル、ワーカー ランタイム、ネイティブ コンパイラー、およびコンテンツ アドレス指定された世代を備えた、自律エージェント用の AI ネイティブのセルフホスティング プラットフォーム。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
3
2486世代
最新の
2026 年 7 月 15 日
+ 2 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An AI-native, self-hosting platform for autonomous agents, with a user-space microkernel, worker runtime, native compiler, and content-addressed generations. - grigoriitropin/vehir-platform

GitHub - grigoriitropin/vehir-platform: An AI-native, self-hosting platform for autonomous agents, with a user-space microkernel, worker runtime, native compiler, and content-addressed generations. · GitHub
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
grigoriitropin
/
vehir-platform
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits ARCHITECTURE.md ARCHITECTURE.md INSTALL.md INSTALL.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Vehir Platform is an experimental AI-native computing environment designed around agent–computer interaction. In the same way that conventional systems use human–computer interaction principles to make computing accessible to people, Vehir designs the entire software environment around the needs of AI agents.
Its interfaces and tooling are designed around structured, discoverable input and output that AI agents can use directly. Vehir explores how the software environment itself can be made more convenient for machine reasoning, rather than treating agent support only as an additional interface over existing workflows.
The current implementation provides the foundations of this environment, including structured agent tools, IPM packages, native compilation and loading, content-addressed generations, reconciliation, and supervised services. The platform is still under active development.
4 GB of RAM or more recommended
Copy the following and send it to your AI agent:
Read the complete installation guide at
https://github.com/grigoriitropin/vehir-platform/blob/main/INSTALL.md
before taking any action. Follow it from beginning to end, and ask the
user whenever it requires a choice or information about the host system.
Why AI-native
Vehir is built for machine-to-machine interaction. Instead of wrapping
existing human-facing tools with AI interfaces, it designs every layer
to be directly usable by agents:
Structured input and output reduce reliance on free-form text parsing.
Self-discoverable tools — the system tells the agent what it can do.
Machine-readable errors with remediation hints.
The source language (IPM) is designed for machine writing and verification, not manual editing.
Declarative build — the agent declares desired state; the system reconciles.
Some features are still under development and may not yet be complete.
Vehir is more than an agent tool server. It combines several layers into
one self-hosting platform:
User-space microkernel and worker runtime. The kernel launches and
manages workers using arenas, shared memory, and doorbell-based
coordination. Managed services use the same worker model.
Self-hosting native compiler. IPM packages are compiled directly to
x86-64 native code. The compiler builds itself, with fixpoint verification
across generations.
Native loading. Vehir resolves object closures, applies relocations,
and loads its own compiled artifacts without relying on a conventional
application runtime.
Content-addressed storage. Sources and compiled objects are addressed
by digest, deduplicated, and recorded in immutable generations.
Declarative reconciliation. A declared build state and package graph
are transformed into a verified generation, followed by an atomic
current switch.
Code-derived agent tools. MCP operations, request contracts, help, and
request validation are derived from the tool implementations rather than
maintained as separate handwritten interfaces.
Whole-universe enforcement. Cross-package analysis checks structural
and policy invariants before a generation is promoted.
This GitHub repository contains documentation only. Release archives are
not binary-only distributions: every published generation archive
( N.tar.gz ) contains the complete IPM source tree together with the native
artifacts needed to bootstrap that exact generation. The source can be
inspected before any Vehir component is started.
Vehir is experimental and under active development. Interfaces and
formats may change between releases.
INSTALL.md — installation guide for AI agents
ARCHITECTURE.md — technical architecture
Before removing anything, identify the actual Vehir installation location
and ask the user to confirm it. Stop and remove the user service:
systemctl --user disable --now vehir-kernel.service
rm ~ /.config/systemd/user/vehir-kernel.service
systemctl --user daemon-reload
Remove only the Vehir MCP entry from the agent client's configuration and
the persistent Vehir instruction added during installation, preserving all
unrelated configuration.
The installation directory is the complete Vehir content-addressed store.
It contains all generations, compiled objects, and source objects. After the
user confirms that none of this state is needed, remove the confirmed absolute
installation location:
rm -rf -- < confirmed-absolute-installation-location >
About
An AI-native, self-hosting platform for autonomous agents, with a user-space microkernel, worker runtime, native compiler, and content-addressed generations.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
3
Generation 2486
Latest
Jul 15, 2026
+ 2 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
