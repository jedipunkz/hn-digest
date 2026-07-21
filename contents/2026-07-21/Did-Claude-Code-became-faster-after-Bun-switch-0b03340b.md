---
source: "https://claude.com/product/claude-code"
hn_url: "https://news.ycombinator.com/item?id=48989141"
title: "Did Claude Code became faster after Bun switch"
article_title: "Claude Code by Anthropic | AI Coding Agent, Terminal, IDE"
author: "preetham_rangu"
captured_at: "2026-07-21T07:33:29Z"
capture_tool: "hn-digest"
hn_id: 48989141
score: 1
comments: 0
posted_at: "2026-07-21T07:21:21Z"
tags:
  - hacker-news
  - translated
---

# Did Claude Code became faster after Bun switch

- HN: [48989141](https://news.ycombinator.com/item?id=48989141)
- Source: [claude.com](https://claude.com/product/claude-code)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T07:21:21Z

## Translation

タイトル: Bun 切り替え後、クロード コードは速くなりましたか?
記事のタイトル: Claude Code by Anthropic | AIコーディングエージェント、ターミナル、IDE
説明: Anthropic の開発者向けエージェント コーディング ツール。 Claude Code はコードベースを理解し、ファイルを編集し、コマンドを実行し、より迅速な出荷を支援します。

記事本文:
クロード・コード by Anthropic | AI コーディング エージェント、ターミナル、IDE メイン コンテンツにスキップ Claude 製品紹介 Claude
コードベースでクロードを直接操作します。ターミナル、IDE、Slack、Web などからビルド、デバッグ、配布します。
macOS 用のダウンロード ドキュメントを読む (新しいタブで開きます) macOS、Linux、および Windows で利用できます。
クロード コードのその他の使用方法:
チャット コワーク コード 新しいセッション ルーチン
もっと固定するをカスタマイズする 設定にダーク モードの切り替えを追加する スケジュール設定 毎週の依存関係監査 最近 チェックアウトの二重請求バグを修正する 支払いモジュールのテストを作成する このリポジトリの内容を説明する モバイルでレイアウトが壊れている 不安定な CI テストをデバッグする スーザンのサンプル acme-dashboard / ダーク モードの切り替えを設定に追加する ダーク モードの切り替えを設定ページに追加します。選択を覚えておき、選択するまでシステムのデフォルトに従う必要があります。アプリをラップする ThemeProvider がすでにありますが、ライトに配線されています。モードをコンテキストに引き上げて localStorage に永続化し、 [設定] → [外観] にライト/ダークのセグメント化されたコントロールを追加します。
たった今 4 つのファイルを読み取り、コードベースを検索しました。 ThemeProvider.tsx を編集しました + 18 - 2 src/theme/ThemeProvider.tsx 10 エクスポート関数 ThemeProvider ({ Children }: { Children : ReactNode }) { 11 + constprefersDark = useMediaQuery ('(prefers - color -cheme : dark )') 12 const storage = localStorage 。 getItem (' テーマ ') 13 + const [ mode , setMode ] = useState (stored ?? (prefersDark ? 'dark' : 'light' )) 14 15 + useEffect (() => { 16 + localStorage . setItem (' テーマ ', mode ) 17 + }, [ mode ]) 18 19 return ( 20 < ThemeContext . Provider value ={{ mode , setMode }}> 21 {children } 編集された settings.tsx + 21 - 0 編集された tokens.css + 24 - 0 Run 2 コマンド 完了 — [外観] セクションの右側にトグルが表示されます。

アプリ全体がこれに従う必要があります。リデュースモーションと密度の行はスタブ化されているため、ページが裸に見えません。
just now # 112 acme-dashboard claude/settings-dark-mode-Xk4fq CI Auto Opus Extra high localhost:5173 /settings 外観テーマ ダッシュボードの外観を選択します。明暗密度 パワーユーザー向けのコンパクトな列。デフォルト コンパクト モーションを減らす アニメーションとトランジションを無効にします。
個人チームとエンタープライズ プロ
Claude Code は Pro プランに含まれています。小さなコードベースでの短いコーディング スプリントに最適です。
年間サブスクリプション割引付きの月額 (200 ドルの前払い)。毎月請求の場合は 20 ドル。
Claude Code は Max プランに含まれています。大規模なコードベースで日常的に使用する場合に大きな価値があります。
Max プランにはさらに多くのクロード コードが含まれています。クロード モデルに最もアクセスできるパワー ユーザーにとって大きな価値があります。
使用制限が適用されます。表示価格には該当する税金は含まれておりません。価格とプランは Anthropic の裁量により変更される場合があります。
使用制限が適用されます。価格とプランは Anthropic の裁量により変更される場合があります。
動的なワークフロー: 数十から数百の並列サブエージェントにわたって実行し、何かが届く前にその動作をチェックすることで、最も困難なタスクに取り組みます。
ブログ 2026 年 5 月 28 日 エージェント ビュー: すべてのクロード コード セッションを 1 か所で管理します。
ブログ 2026 年 5 月 11 日 ルーチン: ルーチンを一度設定すると、スケジュールに従って、API 呼び出しから、またはイベントに応答して実行できます。
ブログ 2026 年 4 月 14 日 コンピューターの使用: クロードはアプリを開いてブラウザーを操作し、開発ツールを実行してタスクを完了します。
ブログ 2026 年 3 月 23 日 変更履歴を表示 (新しいタブで開きます) Claude Code で何ができるでしょうか?
オンボーディング トリアージの問題 コードをリファクタリングする このコードベースは初めてです。説明してもらえますか？
⏺ Excalidraw コードベース構造を分析し、概要を説明します。
⏺ 分析(何を: "E

xcalidraw コードベースの構造について説明し、その内容、主なコンポーネント、およびそれらがどのように構成されているかについての概要を示します。プロジェクトの目的、アーキテクチャ、主要な機能を理解することに重点を置いてください。」)
Excalidraw は、手書きスタイルのオープンソースの仮想ホワイトボード アプリケーションです。図、ワイヤーフレーム、共同図面を作成するために設計されています。
## 目的
- 無限のキャンバスを備えた Web ベースの描画ツール
- 形状、矢印、テキストをサポートする手描きの美しさ
- エンドツーエンドの暗号化されたコラボレーション
- PWA (Progressive Web App) としてオフラインで動作します
## アーキテクチャ
プロジェクトは、次の主要コンポーネントを含むモノリポジトリとして編成されています。
1. コアパッケージ (パッケージ/excalidraw)
- Excalidraw を任意のアプリケーションに埋め込むための React コンポーネント ライブラリ
- 描画、レンダリング、およびコア機能を処理します
- 手描きスタイルにはroughjsを使用します
- エクスポート機能（PNG、SVG）が含まれています
2. Webアプリケーション（excalidraw-app/）
- excalidraw.com のフル機能の Web アプリ
- コアパッケージの上に構築
- コラボレーション機能、ローカルストレージ、共有を追加
3. サポートパッケージ
- common - 共有ユーティリティと定数
- 要素 - 要素の操作と型の定義
- 数学 - 描画のための数学関数
- utils - 一般的なユーティリティ関数
4. 例
- 連携例（Next.js、ブラウザスクリプト）
- Excalidraw をさまざまな環境に埋め込む方法を示します
5. ドキュメント (dev-docs/)
- APIドキュメント
- 統合ガイド
- コンポーネントのドキュメント
## 主な機能
- エンドツーエンド暗号化によるリアルタイムコラボレーション
- さまざまな形式 (PNG、SVG、JSON) にエクスポート
- カスタマイズ可能な描画ツール (長方形、円、矢印、フリーハンド)
- ダークモードのサポート
- ローカリゼーション
- 形状とコンポーネントのライブラリ
- 矢印のバインディングとラベル
- 画像のサポート
## テクノロジースタック
- 反応する

r UI
-
[切り捨てられた]
Claude Code は、コードベース全体を数秒でマッピングして説明します。エージェント検索を使用して、コンテキスト ファイルを手動で選択することなく、プロジェクトの構造と依存関係を理解し​​ます。
ツール間を行き来するのをやめてください。 Claude Code は、GitHub、GitLab、およびコマンド ライン ツールと統合して、ワークフロー全体 (問題の読み取り、コードの作成、テストの実行、PR の送信) をすべて端末から処理します。
Claude Code はコードベースと依存関係を理解することで、機能する強力な複数ファイル編集を行うことができます。
超強力な端末統合。あらゆる IDE と並行して、すべての CLI ツールと連携して動作します。
VS Code (+ Cursor、Devin Desktop) および JetBrains IDE のネイティブ拡張機能。
ブラウザ、モバイル アプリ、またはデスクトップのクロードからクイック アクセス。並行作業や外出先でのコーディングに最適です。
Slack でコーディングタスクを開始する
「 Claude Code のおかげで、チームのコーディング効率が劇的に向上しました。ノートブックに EDA コードを記述して、データを取得し、モデルをトレーニングし、基本的なメトリクスで評価できるようになりました。そして、それを Metaflow パイプラインに変換するよう Claude に依頼できます。このプロセスにより、モデルごとに 1 ～ 2 日の日常的な (そしてしばしば退屈な!) 作業が節約されます。」
アントン・ビリュコフ、スタッフソフトウェアエンジニア
「 Claude のおかげで、私たちはカスタマー サービスを自動化するだけではなく、真の人間的品質にまで高めています。これにより、サポート チームはカスタマー エクスペリエンスと、インタラクションを真に価値あるものにするものについて、より戦略的に考えることができるようになります。」
「 Claude Code は私たちのチームを 1 つ上のレベルに引き上げています。私たちは何が起こる必要があるかを決定し、エンドツーエンドで構築して検証できるようにプロセスをスムーズにします。私の今の仕事の大きな部分は、できるだけ多くの Claude Code のインスタンスをビジー状態に保つことです。」
お気に入りのコマンドラインツールと接続します
ターミナルは実際の作業が行われる場所です。クロード・コードは

開発を強化するツール (展開、データベース、監視、バージョン管理)。ジャグリングするために別のインターフェイスを追加するのではなく、既存のスタックを強化します。
クロードを使い始めるにはどうすればよいですか?
Claude Code には、Claude Pro または Max プラン、Team または Enterprise プランのプレミアム シート、または Claude Console アカウントを使用してアクセスできます。クロード コードをダウンロードし、それぞれのクロードまたはコンソールの資格情報でサインインします。
クロードコードはどのような種類のタスクを処理できますか?
Claude Code は、バグ修正やテストなどの日常的な開発タスクと、コードベースの深い理解を必要とするリファクタリングや機能実装などの変革的な作業の両方に優れています。
Claude Code は既存のツールとどのように連携しますか?
Claude Code はターミナル内で実行され、ワー​​クフローを変更することなく、好みの IDE や開発ツールと連携して動作します。 Claude Code は、コマンド ライン ツール (Git など) や MCP サーバー (GitHub など) を使用して、ツールを使用して独自の機能を拡張することもできます。
はい。 Claude Code は端末内でローカルに実行され、バックエンド サーバーやリモート コード インデックスを必要とせずにモデル API と直接通信します。また、ファイルを変更したりコマンドを実行したりする前に、許可を求めます。
クロード コードを実行するためのシステム要件は何ですか?
Claude Code は macOS、Linux、および Windows で動作します。完全なシステム要件を参照してください。
クロードコードの価格はいくらですか?
Claude Console アカウントで使用すると、Claude Code は標準 API 価格で API トークンを消費します。
Claude Code は Claude デスクトップ アプリで動作しますか?
はい。 Max、Pro、Team、および Enterprise ユーザーは、Claude デスクトップ アプリで Claude Code にアクセスできます。
クロードコードの高速モードとは何ですか?
高速モードは Opus 4.8 の高速構成であり、トークンあたりのコストは高くなりますが、モデルの速度は 2.5 倍になります。高速モードが利用可能です:
Cl の研究プレビュー中

aude Code で、価格は 100 万トークンあたり 30 ドルまたは 150 ドルです。
サブスクリプション プランのユーザーの使用量クレジットを介して。
開発者向けドキュメント 一般的なワークフロー
開発者ドキュメント CLAUDE.md ファイルの使用
ブログ エージェントコーディングの紹介
ブログ Anthropic チームがクロード コードを使用する方法
ブログ Claude を使用してソフトウェアのバグをより速く修正する
ブログ 刺激的なものを作成します。必要不可欠なものを維持します。
職場でクロード コードを使用する
製品のアップデート、ハウツー、コミュニティのスポットライトなど。毎月あなたの受信箱に配信されます。
毎月の開発者ニュースレターを受け取りたい場合は、電子メール アドレスを入力してください。いつでも購読を解除できます。
データ処理契約: 米国 K-12
データ処理契約: 米国 K-12

## Original Extract

Anthropic's agentic coding tool for developers. Claude Code understands your codebase, edits files, runs commands, and helps you ship faster.

Claude Code by Anthropic | AI Coding Agent, Terminal, IDE Skip to main content Meet Claude Products Claude
Work with Claude directly in your codebase. Build, debug, and ship from your terminal, IDE, Slack, web, and more.
Download for macOS Read documentation (opens in new tab) Available for macOS, Linux, and Windows.
Other ways to use Claude Code:
Chat Cowork Code New session Routines
Customize More Pinned Add a dark mode toggle to settings Scheduled Weekly dependency audit Recents Fix the double-charge bug in checkout Write tests for the payments module Explain what this repo does Layout's broken on mobile Debug flaky CI tests Susan Sample acme-dashboard / Add a dark mode toggle to settings Add a dark mode toggle to the settings page. Should remember the choice and follow the system default until they pick one. There's already a ThemeProvider wrapping the app but it's hard-wired to light. I'll lift the mode into context, persist it to localStorage , and add a Light / Dark segmented control under Settings → Appearance .
just now Read 4 files, searched the codebase Edited ThemeProvider.tsx + 18 - 2 src/theme/ThemeProvider.tsx 10 export function ThemeProvider ({ children }: { children : ReactNode }) { 11 + const prefersDark = useMediaQuery ('( prefers - color - scheme : dark )') 12 const stored = localStorage . getItem (' theme ') 13 + const [ mode , setMode ] = useState ( stored ?? ( prefersDark ? 'dark' : 'light' )) 14 15 + useEffect (() => { 16 + localStorage . setItem (' theme ', mode ) 17 + }, [ mode ]) 18 19 return ( 20 < ThemeContext . Provider value ={{ mode , setMode }}> 21 { children } Edited settings.tsx + 21 - 0 Edited tokens.css + 24 - 0 Ran 2 command s Done — the Appearance section now has the toggle. Preview's on the right; flip it and the whole app should follow. Reduce-motion and density rows are stubbed so the page doesn't look bare.
just now # 112 acme-dashboard claude/settings-dark-mode-Xk4fq CI Auto Opus Extra high localhost:5173 /settings Appearance Theme Choose how the dashboard looks. Light Dark Density Compact rows for power users. Default Compact Reduce motion Disable animations and transitions.
Individual Team & Enterprise Pro
Claude Code is included in your Pro plan. Perfect for short coding sprints in small codebases.
Per month with annual subscription discount ( $200 billed up front). $20 if billed monthly.
Claude Code is included in your Max plan. Great value for everyday use in larger codebases.
Even more Claude Code included in your Max plan. Great value for power users with the most access to Claude models.
Usage limits apply . Prices shown don't include applicable tax. Price and plans are subject to change at Anthropic's discretion.
Usage limits apply . Price and plans are subject to change at Anthropic's discretion.
Dynamic workflows: Tackle the most challenging tasks by executing across 10s to 100s of parallel subagents, and checking its work before anything reaches you.
Blog May 28, 2026 Agent view: One place to manage all your Claude Code sessions.
Blog May 11, 2026 Routines: Configure a routine once, and it can run on a schedule, from an API call, or in response to an event.
Blog Apr 14, 2026 Computer use: Claude now opens your apps, navigates your browser, and runs your dev tools to complete tasks.
Blog Mar 23, 2026 View changelog (opens in new tab) What could you do with Claude Code?
Onboarding Triage issues Refactor code I'm new to this codebase. Can you explain it to me?
⏺ I'll analyze the Excalidraw codebase structure and provide a high-level overview.
⏺ Analyze(what: "Excalidraw codebase structure and provide a high-level overview of what it is, its main components, and how they're organized. Focus on understanding the project's purpose, architecture, and key features.")
Excalidraw is an open-source virtual whiteboard application with a hand-drawn style. It's designed for creating diagrams, wireframes, and collaborative drawings.
## Purpose
- Web-based drawing tool with infinite canvas
- Hand-drawn aesthetic with support for shapes, arrows, text
- End-to-end encrypted collaboration
- Works offline as a PWA (Progressive Web App)
## Architecture
The project is organized as a monorepo with these main components:
1. Core Package (packages/excalidraw)
- React component library for embedding Excalidraw in any application
- Handles drawing, rendering, and core functionality
- Uses roughjs for the hand-drawn style
- Includes export capabilities (PNG, SVG)
2. Web Application (excalidraw-app/)
- Full-featured web app at excalidraw.com
- Built on top of the core package
- Adds collaboration features, local storage, sharing
3. Support Packages
- common - Shared utilities and constants
- element - Element manipulation and type definitions
- math - Mathematical functions for drawing
- utils - General utility functions
4. Examples
- Integration examples (Next.js, browser script)
- Shows how to embed Excalidraw in different environments
5. Documentation (dev-docs/)
- API documentation
- Integration guides
- Component documentation
## Key Features
- Real-time collaboration with end-to-end encryption
- Export to various formats (PNG, SVG, JSON)
- Customizable drawing tools (rectangle, circle, arrows, freehand)
- Dark mode support
- Localization
- Library of shapes and components
- Arrow bindings and labels
- Image support
## Technology Stack
- React for UI
-
[truncated]
Claude Code maps and explains entire codebases in a few seconds. It uses agentic search to understand project structure and dependencies without you having to manually select context files.
Stop bouncing between tools. Claude Code integrates with GitHub, GitLab, and your command line tools to handle the entire workflow—reading issues, writing code, running tests, and submitting PRs—all from your terminal.
Claude Code's understanding of your codebase and dependencies enables it to make powerful, multi-file edits that work.
Super powerful terminal integration. Works with all your CLI tools alongside any IDE.
Native extensions for VS Code (+ Cursor, Devin Desktop) and JetBrains IDEs.
Quick access from browser, mobile app, or Claude on desktop. Great for parallel work or on-the-go coding.
Kick off coding tasks in Slack
“ Claude Code has dramatically accelerated our team's coding efficiency. I can now write EDA code in a notebook—pulling data, training a model, and evaluating it with basic metrics—and then ask Claude to convert that into a Metaflow pipeline. This process saves 1-2 days of routine (and often boring!) work per model. ”
Anton Biryukov, Staff Software Engineer
“ With Claude, we're not just automating customer service—we're elevating it to truly human quality. This lets support teams think more strategically about customer experience and what makes interactions genuinely valuable. ”
“ Claude Code is moving our team up a level: we decide what needs to happen, and smooth the process so it can build and verify end-to-end. A big part of my job now is to keep as many instances of Claude Code busy as possible. ”
Connects with your favorite command line tools
Your terminal is where real work happens. Claude Code connects with the tools that power development—deployment, databases, monitoring, version control. Rather than adding another interface to juggle, it enhances your existing stack.
How do I get started with Claude?
You can access Claude Code with a Claude Pro or Max plan, a Team or Enterprise plan premium seat, or a Claude Console account. Download Claude Code and sign in with your respective Claude or Console credentials.
What kinds of tasks can Claude Code handle?
Claude Code excels at both routine development tasks like bug fixes and testing, as well as transformative work like refactors and feature implementation that require deep codebase understanding.
How does Claude Code work with my existing tools?
Claude Code runs in your terminal and works alongside your preferred IDE and development tools without requiring you to change your workflow. Claude Code can also use command line tools (like Git) and MCP servers (like GitHub) to extend its own capabilities using your tools.
Yes. Claude Code runs locally in your terminal and talks directly to model APIs without requiring a backend server or remote code index. It also asks for permission before making changes to your files or running commands.
What are the system requirements to run Claude Code?
Claude Code works on macOS, Linux, and Windows. See full system requirements .
How much does Claude Code cost?
When used with a Claude Console account, Claude Code consumes API tokens at standard API pricing .
Does Claude Code work with the Claude desktop app?
Yes. Max, Pro, Team, and Enterprise users can access Claude Code on the Claude desktop app .
What is fast mode on Claude Code?
Fast mode is a high-speed configuration for Opus 4.8, making the model 2.5x faster at a higher cost per token. Fast mode is available:
In research preview on Claude Code, and is priced at $30/$150 per million tokens.
Via usage credits for users on subscription plans.
Developer docs Common workflows
Developer docs Using CLAUDE.md files
Blog Introduction to agentic coding
Blog How Anthropic teams use Claude Code
Blog Fix software bugs faster with Claude
Blog Create what's exciting. Maintain what's essential.
Use Claude Code where you work
Product updates, how-tos, community spotlights, and more. Delivered monthly to your inbox.
Please provide your email address if you'd like to receive our monthly developer newsletter. You can unsubscribe at any time.
Data Processing Agreement: US K-12
Data Processing Agreement: US K-12
