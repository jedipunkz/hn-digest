---
source: "https://github.com/OnPoint-Dev-Tools/crewcode"
hn_url: "https://news.ycombinator.com/item?id=49364088"
title: "Show HN: CrewCode – Open-source Mission Control for AI coding agents"
article_title: "GitHub - OnPoint-Dev-Tools/crewcode: CrewCode is a free, open-source desktop ACE for developers who already work repo-first and agent-first. · GitHub"
image: "https://opengraph.githubassets.com/d792cde8c1b577f4ad0e85f64107c153ec169e458e21d59f6cf2ee4e746555e1/OnPoint-Dev-Tools/crewcode"
author: "CjLogix"
captured_at: "2026-08-19T17:18:47Z"
capture_tool: "hn-digest"
hn_id: 49364088
score: 1
comments: 0
posted_at: "2026-08-19T16:56:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: CrewCode – Open-source Mission Control for AI coding agents

- HN: [49364088](https://news.ycombinator.com/item?id=49364088)
- Source: [github.com](https://github.com/OnPoint-Dev-Tools/crewcode)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T16:56:22Z

## Translation

タイトル: Show HN: CrewCode – AI コーディング エージェント用のオープンソース ミッション コントロール
記事のタイトル: GitHub - OnPoint-Dev-Tools/crewcode: CrewCode は、すでにリポジトリ ファーストおよびエージェント ファーストで作業している開発者向けの、無料のオープンソース デスクトップ ACE です。 · GitHub
説明: CrewCode は、すでにリポジトリ ファーストおよびエージェント ファーストで作業している開発者向けの、無料のオープンソース デスクトップ ACE です。 - OnPoint-Dev-Tools/クルーコード
HN テキスト: 私はエージェントのコーディングをかなり実験してきましたが、コードを書くよりもエージェントを管理することに気付きました。それを解決するために CrewCode を構築します。これは、Git ワークツリーで作業を分離しながら、複数の AI コーディング エージェントを並行して実行するためのオープンソース デスクトップ アプリケーションです。 Git ワークツリー → エージェントの分離 → プロバイダーの独立 → 差分レビュー → オーケストレーション。特にマルチエージェント/ワークツリー モデルに関するフィードバックをいただければ幸いです。本日 Product Hunt でリリースしました。開発者からのフィードバックをお待ちしています。

記事本文:
GitHub - OnPoint-Dev-Tools/crewcode: CrewCode は、すでにリポジトリ ファーストおよびエージェント ファーストで作業している開発者向けの、無料のオープンソース デスクトップ ACE です。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
OnPoint-Dev-Tools
/
クルーコード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
18 コミット 18 コミット フォルダーとファイル
.github/ ワークフロー .github/ ワークフロー ビン ビン ビルド

ビルド ドキュメント ドキュメント 例/プラグイン 例/プラグイン パッケージ パッケージ スキーマ スキーマ スクリプト スクリプト サービス/ローカル音声サービス/ローカル音声 src src .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md LICENSEライセンスに関する通知 通知 README.md README.md SECURITY.md SECURITY.md crewcoder.json crewcoder.json electric.vite.config.ts electric.vite.config.ts package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json tsconfig.node.tsbuildinfo tsconfig.node.tsbuildinfo tsconfig.web.json tsconfig.web.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
CrewCode は、マルチエージェント ソフトウェア開発のためのコントロール センターです。
リポジトリの制御を失うことなく、git ワークツリー全体で複数の AI コーディング エージェントを実行、監視、レビューします。
CrewCode は、すでにリポジトリ ファーストおよびエージェント ファーストで作業している開発者向けの、無料のオープンソース デスクトップ ACE です。
CrewCode は、ターミナル、ワークツリー、PR ページ、ブラウザ タブ、チャット スレッドを複数のツール間でやりくりするのではなく、完全なワークフローを 1 か所にまとめます。
CrewCode と CrewCoder を組み合わせて、統合されたエージェント開発ワークフローを実現します。 CrewCode には、ACP を介したファーストクラスの CrewCoder サポートが含まれているため、アプリから直接 CrewCoder を実行および監視し、ワークフローに合わせたカスタム CrewCode プラグインの作成をアプリに依頼できます。セットアップの詳細については CrewCoder プロバイダー ガイドを、CrewCode プラグインの詳細についてはプラグイン ガイドを参照してください。
複数のコーディング エージェントを実行すると、すぐにワークフローが混乱してしまいます。
各エージェントが何をしているのかが十分に把握できない
CrewCode は、ワークフローを監視し、検査可能にし、安全に着陸しやすくするために構築されています。
MacOS をターゲットとしています (そうであればありがたいです)

meone は、Mac、Linux、Windows のテストをお手伝いします。
1 つのワークスペースで複数のエージェントを実行する
CrewCoder エージェント、Claude Code エージェント、Codex 、OpenCode 、pi 、Ollama 、Hermes 、OpenRouter および Grok Build を 1 つのアプリから実行
プロバイダーに応じて、構造化エージェント ブリッジと PTY ベースの端末を混在させる
ワークスペースごとに複数のチャット タブ セッションを保持し、それぞれに独自のエージェント、モデル、モード、および作業量を設定します。
毎回最初から開始するのではなく、永続的なセッションを再開します
ローカル プラグイン、カスタム パネル、MCP サーバー、ブラウザ/Git アクション、プラグインを利用したエージェント プロバイダーを使用して CrewCode を拡張します。
セッション中にプロバイダーを切り替えると、新しいエージェントがチャットを離れてコンテキストを失うことなく作業を継続できるように、引き継ぎ概要が作成されます。
ワークベンチ モード: 同じワークツリー上で複数のチャットとターミナルを一度に実行します。
/compact はいつでも必要に応じて実行でき、最近の会話の概要が作成され、内部で新しいセッションで作業を続けることができます。
リポジトリのクローンをローカル ワークスペースに直接作成する
新しいプロジェクトを最初から初期化する
アプリ内で git ワークツリーを作成、切り替え、マージ、削除する
スタッシュやブランチのジャグリングを行わずに、並行作業を分離して維持します。
乗組員が結合する前に、正確なファイルの重複とファイル間のコントラクトの衝突の可能性を明らかにする
再起動後にクルーのワークツリー/コミットの所有権と中断されたマージ監査状態を回復します
エージェントのチームを並行して立ち上げる
さまざまな役割、エージェント、モデル、作業レベルでレーンを構成する
分離されたワークツリーまたは共有ワークスペース パターンでエージェントを実行する
レーンを個別に再起動、ミュート、再ブロードキャスト、検査する
コーディング エージェントが実際の永続的なチャット セッションを内部から作成して実行できるようにします。
順番に、「回帰テストを実行するために別のスレッドでいくつかのエージェントを起動してください、私を待ってください」
updated」は、自分で開いて読んで続行できるスレッドを生成します。
これはディです

クルーサーフェスからのこだわり。クルーレーンは一時的なもので、1 つのタブ内に存在し、
乗組員の走行中は存在します。委任されたスレッドは通常の Session です。
トランスクリプト、アーカイブ、および引き出しの行は、作成されたターンよりも長く存続します。
監視用のコントロールセンター
プロジェクト全体でアクティブなエージェントを 1 つのダッシュボードで確認する
プロジェクト、ステータス、タイプ、またはワークツリー別にアクティビティをグループ化する
ライブクロスアプリアクティビティフィードをフォローする
リポジトリのステータス、履歴、ブランチ、ワークツリーを表示する
ステージング、ステージング解除、差分、コミット、プル、プッシュ、フェッチ、同期
アプリ内で AI によって生成された変更を確認する
マージ競合を処理し、解決をエージェントに委任します。
GitHub CLI 統合による PR の作成、承認、マージ
エージェントのチャットの横でマークダウンまたはプレーンテキストのドキュメントの下書きとレビューを行う
DOCX および PDF ファイルを、オリジナルを上書きせずに編集可能な Markdown 作業コピーに変換します
エージェントの編集を検出し、承認または拒否する前にそれらを Pierre テキストの差分としてレビューします
承認された作業コピーを衝突安全な DOCX または PDF ファイルにエクスポート
CrewCode 内の埋め込みブラウザ タブを開く
選択した要素からページのコンテキストを取得する
選択した要素からクラス名を取得します
ブラウザーの選択内容のスクリーンショットをキャプチャする
ブラウザの調査結果をエージェント チャットに直接送信します
リアルタイムの音声と作曲家のディクテーション
CrewCode には、自然に会話するためのプロバイダーに依存しないリアルタイム音声レイヤーが含まれています
アクティブなチャットでコーディング エージェントと。
選択したコーディング エージェントとの音声会話には音声 Orb を使用します
エージェントに送信する前に、各音声トランスクリプトを確認、編集、またはキャンセルします。
既存の CrewCode セッション、ワークスペース、モード、および権限の境界を通じて音声リクエストをルーティングします。
コードブロック、差分、テーブル、ログ、URL、長いパスを省略した完全な散文応答を聞いてください。
代わりに、GPT リアルタイム、xAI 音声、または完全なローカル音声を選択してください。

1 つのプロバイダーにロックされている
文字起こしには NVIDIA Parakeet TDT 0.6B v2 を使用してローカル音声を実行し、音声には am_michael 音声を使用して Kokoro-82M を実行します。
[自動]、[GPU]、または [CPU ローカル推論] を選択します。アイドル状態のモデルはアンロードされますが、軽量の音声サイドカーは利用可能なままです
ホストされたプロバイダー キーを Electron のメイン プロセスに保持します。永続キーはレンダラー状態に入ることはありません
別のコンポーザ マイクは音声からテキストへの変換のみに使用します。レビュー用にキャレットにテキストが挿入されます。
リアルタイム音声はデフォルトではオフになっています。ホストされているプロバイダーには独自の API キーが必要です
そして請求。ローカル音声には Python 3.11 環境が必要です。
ネイティブの依存関係。 docs/realtime-voice.md を参照してください。
セットアップ、アーキテクチャ、プロバイダーの可用性、セキュリティの詳細については。
再利用可能なプロンプトのローカル ライブラリを維持する
再利用可能なスキル/手順のローカル ライブラリを維持する
使用状況を追跡し、アプリ内でプロンプト/スキルを編集します
ライブターミナルペインとエージェントデーモンプロセスを監視する
内蔵システムモニターでCPUとメモリの使用状況を検査
エージェントが終了したとき、または注意が必要なときに通知を受け取る
— 同じワークツリー上で複数のチャットとターミナルを同時に実行します。
TypeScript/LSP または言語ごとのサービスによる言語対応補完。
タブ ローカル エージェントまたはホスト型 API プロバイダーによる補完用のタブ オートコンプリート
独自のプラグインが必要ですか? CrewCoder に、ワークフローに従って作成するよう指示します。
テーマ、ドキュメント アウトライン、TypeScript と JavaScript インテリジェンス、問題引き出し、コード アクション、安全性とパフォーマンスなど。
CrewCode には、信頼できる開発者自動化のためのローカルファーストのプラグイン システムが含まれるようになりました。
~/.crewcode/plugins/<plugin-id>/crewcode.plugin.json からプラグインをインストールします
分離されたプラグイン パネル、サイドバー パネル、ステータス アイテム、エディター/チャット アクション、Git レンズ、ブラウザー アクション、ターミナル ウォッチャー アクション、MCP サーバー宣言を追加します。

レーション、およびカスタム エージェント プロバイダー
crewcode-plugin:// サンドボックス化された iframe を通じてロードされる静的 HTML/JS アセットとしてプラグイン UI を構築します。
CrewCode への安全な postMessage 呼び出しには、型指定された crewcode-plugin-api ブラウザ パッケージを使用します
専用のプラグイン ページから、権限の承認/取り消し、プラグインのグローバルな有効化/無効化、ワークスペースごとのプラグインの有効化/無効化を行うことができます。
マニフェストの検証、アセットの読み込み、iframe ランタイム エラー、機能の拒否、プロバイダーの生成の失敗、HTTP プロバイダーの失敗などの分類されたログを使用してプラグインをデバッグします
プラグイン エージェント プロバイダーは、mock 、 exec 、および http ランタイムをサポートしているため、プラグイン UI に Electron または Node への直接アクセスを与えずに、ローカル CLI または内部 HTTP エージェント ゲートウェイに接続できます。現在の v0 コントラクトとテンプレートについては、docs/plugins.md および example/plugins を参照してください。
gh を介した GitHub 認証および PR ワークフロー
リモート ターゲットの SSH 到達可能性テスト
自動更新チャネル: 安定版、ベータ版、毎晩
構成可能なエージェント起動パスとシェル設定
リモート ホスティング: リモート ファイル編集、ターミナル、ホストへのエージェント接続
CrewCode は現在、次の点で最も強力です。
実際のリポジトリで複数のコーディング エージェントを実行する個人開発者
技術的な創設者がエージェントを並行協力者として使用している
小規模のエンジニアリング チームが監視付きマルチエージェント ワークフローを探索
ワークツリーを作成または切り替える
1 つ以上のエージェント セッションを開始する
チャット、ターミナルペイン、Mission Control での出力の監視
差分の検査、変更のコミット、PR のオープンまたはマージ、または作業の破棄
npm 実行開発
npm ビルドを実行する
npm 実行プレビュー
npm タイプチェックを実行する
技術スタック
CrewCode は貢献を歓迎しています。問題、バグ レポート、プル リクエストはすべて歓迎されます。より大きな変更を計画している場合は、まず問題をオープンして、構築する前にアプローチについて話し合ってください。
CrewCode は、Apache License バージョン 2.0 に基づいてライセンスされています。 「LI」を参照

全文はCENSE。
つまり、著作権とライセンス通知を保持し、加えた重要な変更を明記する限り、商用製品やクローズドソース製品を含め、CrewCode を自由に使用、変更、再配布できます。 Apache-2.0 は、貢献者からの明示的な特許ライセンスも付与します。
バンドルされている一部のコンポーネントは異なるライセンスの下にあり、Apache-2.0 の対象外です。完全なリストについては、「通知」を参照してください。
最も注目すべき点は、 src/renderer/src/assets/bearded-icons/ に設定されているファイル ツリー アイコンは BeardedBear による Bearded Icons であり、 GPL-3.0 に基づいてライセンスされています。これらのアセットは変更されずに使用され、CrewCode の Apache ライセンス ソースに組み込まれずに集約されます。 CrewCode をフォークし、完全に許可されたスタックが必要な場合は、そのディレクトリを許可されたアイコン セットに置き換えます。
著作権 © 2026 OnPoint ツール。
CrewCode は、すでにリポジトリ ファーストおよびエージェント ファーストで作業している開発者向けの、無料のオープンソース デスクトップ ACE です。
Readme Apache-2.0 ライセンス
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

CrewCode is a free, open-source desktop ACE for developers who already work repo-first and agent-first. - OnPoint-Dev-Tools/crewcode

I've been experimenting heavily with coding agents and found myself managing agents more than writing code. I build CrewCode to solve that. Its an open source desktop application for running multiple AI coding agents in parallel while isolating their work with Git worktrees. Git worktrees → agent isolation → provider independence → diff review → orchestration. I'd especially appreciate feedback on the multi-agent/worktree model. I launched it on Product Hunt today and would love feedback from developers.

GitHub - OnPoint-Dev-Tools/crewcode: CrewCode is a free, open-source desktop ACE for developers who already work repo-first and agent-first. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
OnPoint-Dev-Tools
/
crewcode
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
18 Commits 18 Commits Folders and files
.github/ workflows .github/ workflows bin bin build build docs docs examples/ plugins examples/ plugins packages packages schemas schemas scripts scripts services/ local-voice services/ local-voice src src .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md crewcoder.json crewcoder.json electron.vite.config.ts electron.vite.config.ts package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json tsconfig.node.tsbuildinfo tsconfig.node.tsbuildinfo tsconfig.web.json tsconfig.web.json vitest.config.ts vitest.config.ts View all files Repository files navigation
CrewCode is the control center for multi-agent software development.
Run, supervise, and review multiple AI coding agents across git worktrees without losing control of your repo.
CrewCode is a free, open-source desktop ACE for developers who already work repo-first and agent-first.
Instead of juggling terminals, worktrees, PR pages, browser tabs, and chat threads across several tools, CrewCode keeps the full workflow in one place.
Pair CrewCode with CrewCoder for an integrated agentic development workflow. CrewCode includes first-class CrewCoder support through ACP, so you can run and supervise CrewCoder directly from the app and ask it to create custom CrewCode plugins tailored to your workflow. See the CrewCoder provider guide for setup details and the plugin guide to learn more about CrewCode plugins.
As soon as you run more than one coding agent, the workflow gets messy fast:
not enough visibility into what each agent is doing
CrewCode is built to make that workflow supervised, inspectable, and easier to land safely.
It targets MacOS (I'd appreciate if someone can help with testing Mac), Linux, and Windows.
Run multiple agents in one workspace
Run CrewCoder agent , Claude Code agent , Codex , OpenCode , pi , Ollama , Hermes , OpenRouter and Grok Build from one app
Mix structured agent bridges and PTY-backed terminals depending on provider
Keep multiple chat tabs sessions per workspace , each with its own agent, model, mode, and effort
Resume persistent sessions instead of starting from scratch every time
Extend CrewCode with local plugins, custom panels, MCP servers, browser/git actions, and plugin-powered agent providers
Switch Providers Mid Session and it creates an hand-off Summary for new agent to continue work with out leaving chat and losing context
Workbench Mode: Run multiple chats and terminals all at once on same worktree
/compact when ever you want and it creates a summary of your recent conversation then you can continue with your work in new session under the hood
Clone repos directly into a local workspace
Initialize new projects from scratch
Create, switch, merge, and remove git worktrees inside the app
Keep parallel work isolated without stashing or branch juggling
Surface exact file overlap and likely cross-file contract collisions before crew merges
Recover crew worktree/commit ownership and interrupted merge audit state after restart
Launch a crew of agents in parallel
Configure lanes with different roles, agents, models, and effort levels
Run agents in isolated worktrees or shared workspace patterns
Restart, mute, rebroadcast to, and inspect lanes independently
Lets a coding agent create and drive real, persistent chat sessions from inside a
turn, so "spin up some agents in another thread to run a regression test, keep me
updated" produces threads you can open, read, and continue yourself.
This is distinct from Crew Surface. Crew lanes are ephemeral, live in one tab, and
exist for the duration of a crew run. Delegated threads are ordinary Session s with
transcripts, archiving, and drawer rows — they outlive the turn that made them.
Control Center for supervision
See active agents across projects in one dashboard
Group activity by project, status, type, or worktree
Follow a live cross-app activity feed
View repo status, history, branches, and worktrees
Stage, unstage, diff, commit, pull, push, fetch, and sync
Review AI-generated changes inside the app
Handle merge conflicts and delegate resolution back to an agent
Create, approve, and merge PRs through GitHub CLI integration
Draft and review Markdown or plain-text documents beside an agent chat
Convert DOCX and PDF files into editable Markdown working copies without overwriting originals
Detect agent edits and review them as Pierre text diffs before accepting or denying
Export approved working copies to collision-safe DOCX or PDF files
Open embedded browser tabs inside CrewCode
Grab page context from selected elements
Grab class names from selected elements
Capture screenshots of browser selections
Send browser research directly into an agent chat
Realtime voice and composer dictation
CrewCode includes a provider-neutral realtime voice layer for talking naturally
with the coding agent in the active chat.
Use the voice orb for voice conversations with your selected coding agent
Review, edit, or cancel each voice transcript before it is sent to the agent
Route spoken requests through the existing CrewCode session, workspace, mode, and permission boundaries
Hear the complete prose reply while omitting code blocks, diffs, tables, logs, URLs, and long paths
Choose GPT Realtime, xAI Voice, or fully local voice instead of being locked to one provider
Run local speech with NVIDIA Parakeet TDT 0.6B v2 for transcription and Kokoro-82M with the am_michael voice for speech
Choose Automatic, GPU, or CPU local inference; idle models unload while the lightweight voice sidecar stays available
Keep hosted provider keys in Electron's main process; permanent keys never enter renderer state
Use the separate composer microphone for speech-to-text only—it inserts text at the caret for review.
Realtime voice is off by default. Hosted providers require their own API keys
and billing; local voice requires a Python 3.11 environment and the documented
native dependencies. See docs/realtime-voice.md
for setup, architecture, provider availability, and security details.
Maintain a local library of reusable prompts
Maintain a local library of reusable skills/instructions
Track usage and edit prompts/skills in-app
Watch live terminal panes and agent daemon processes
Inspect CPU and memory usage with the built-in system monitor
Get notifications when agents finish or need attention
— Run multiple chats and terminals at once on the same worktree.
Language-aware completions via TypeScript/LSP or per-language services.
Tab Tab Autocomplete with local agents or Hosted API Provider for Completions
Want your own plugin? tell CrewCoder to create one for you according to your workflow
Themes, Document Outline, TypeScript and JavaScript intelligence, Problems drawer, Code Actions , Safety and performance and more!
CrewCode now includes a local-first plugin system for trusted developer automation.
Install plugins from ~/.crewcode/plugins/<plugin-id>/crewcode.plugin.json
Add isolated plugin panels, sidebar panels, status items, editor/chat actions, git lenses, browser actions, terminal watcher actions, MCP server declarations, and custom agent providers
Build plugin UIs as static HTML/JS assets loaded through crewcode-plugin:// sandboxed iframes
Use the typed crewcode-plugin-api browser package for safe postMessage calls into CrewCode
Approve/revoke permissions, enable/disable plugins globally, and enable/disable plugins per workspace from the dedicated Plugins page
Debug plugins with categorized logs for manifest validation, asset loading, iframe runtime errors, capability denials, provider spawn failures, and HTTP provider failures
Plugin agent providers support mock , exec , and http runtimes, so you can connect local CLIs or internal HTTP agent gateways without giving plugin UI direct Electron or Node access. See docs/plugins.md and examples/plugins for the current v0 contract and templates.
GitHub auth and PR workflows via gh
SSH reachability testing for remote targets
Auto-update channels: stable, beta, nightly
Configurable agent launch paths and shell preferences
Remote hosting: Remote file editing,terminal, and agents connection to host
CrewCode is strongest today for:
solo developers running multiple coding agents on real repos
technical founders using agents as parallel collaborators
small engineering teams exploring supervised multi-agent workflows
create or switch to a worktree
start one or more agent sessions
watch output in chat, terminal panes, and Mission Control
inspect diffs, commit changes, open or merge PRs, or discard work
npm run dev
npm run build
npm run preview
npm run typecheck
Tech stack
CrewCode is open to contributions — issues, bug reports, and pull requests are all welcome. If you're planning a larger change, open an issue first so we can talk through the approach before you build it.
CrewCode is licensed under the Apache License, Version 2.0 . See LICENSE for the full text.
In short: you are free to use, modify, and redistribute CrewCode, including in commercial and closed-source products, provided you retain the copyright and license notices and state any significant changes you made. Apache-2.0 also grants an explicit patent license from contributors.
Some bundled components are under different licenses and are not covered by Apache-2.0. See NOTICE for the full list.
Most notably, the file-tree icon set in src/renderer/src/assets/bearded-icons/ is Bearded Icons by BeardedBear, licensed under GPL-3.0 . These assets are used unmodified and are aggregated with — not incorporated into — CrewCode's Apache-licensed source. If you fork CrewCode and need a fully permissive stack, replace that directory with a permissively licensed icon set.
Copyright © 2026 OnPoint Tools.
CrewCode is a free, open-source desktop ACE for developers who already work repo-first and agent-first.
Readme Apache-2.0 license Contributing
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
