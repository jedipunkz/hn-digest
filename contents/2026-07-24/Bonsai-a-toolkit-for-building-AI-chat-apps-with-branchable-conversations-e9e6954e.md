---
source: "https://github.com/Joaoha/Bonsai"
hn_url: "https://news.ycombinator.com/item?id=49043023"
title: "Bonsai – a toolkit for building AI chat apps with branchable conversations"
article_title: "GitHub - Joaoha/Bonsai: A library for context banching and context wiki distilation · GitHub"
author: "joaoha"
captured_at: "2026-07-24T23:56:00Z"
capture_tool: "hn-digest"
hn_id: 49043023
score: 1
comments: 0
posted_at: "2026-07-24T23:52:26Z"
tags:
  - hacker-news
  - translated
---

# Bonsai – a toolkit for building AI chat apps with branchable conversations

- HN: [49043023](https://news.ycombinator.com/item?id=49043023)
- Source: [github.com](https://github.com/Joaoha/Bonsai)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T23:52:26Z

## Translation

タイトル: Bonsai – 分岐可能な会話を備えた AI チャット アプリを構築するためのツールキット
記事タイトル: GitHub - Joaoha/Bonsai: コンテキスト バンチングとコンテキスト Wiki 蒸留のためのライブラリ · GitHub
説明: コンテキスト バンチングとコンテキスト Wiki 蒸留のためのライブラリ - Joaoha/Bonsai

記事本文:
GitHub - Joaoha/Bonsai: コンテキスト バンチングとコンテキスト Wiki 蒸留のためのライブラリ · GitHub
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
ジョアオハ
/
盆栽
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

イオンオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
25 コミット 25 コミット .changeset .changeset .github/ workflows .github/ workflows docs docs 例/最小ノードの例/最小ノード パッケージ パッケージ スクリプト スクリプト テスト/フィクスチャ/境界違反 テスト/フィクスチャ/境界違反 .dependency-cruiser.cjs .dependency-cruiser.cjs .dependency-cruiser.fixture.cjs .dependency-cruiser.fixture.cjs .gitignore .gitignore ライセンス ライセンス README.md README.md eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Bonsai は、AI との会話が単一の直線的なスレッドにとどまる必要のないチャット アプリを構築するためのツールキットです。
ほとんどの AI チャット アプリでは、すべての会話が 1 本の直線に沿って進みます。つまり、質問すると答え、また質問すると、履歴全体が順番に積み重なっていきます。接点を探りたい場合は、「別のアプローチを試してみたらどうなるでしょうか?」 — 会話が脱線するか、最初からやり直して、すでに構築していたコンテキストを失うことになります。
盆栽は会話を分岐させます。チャットのどの時点からでも、元のスレッドを妨げることなく分岐してアイデアを検討できます。ブランチが役立つことが判明した場合は、その概要をメインの会話にマージして戻すことができます。それが行き止まりであれば、それを放棄するだけです。それについては何も漏れません。また、会話 (または分岐) で長期保存する価値のあるものが生成された場合は、それを毎回再説明する代わりに、将来の会話で利用できる Wiki ページに抽出できます。
重要なのは、Bonsai がこれらの作業を背後で自動的に行うことは決してありません。分岐、合流、

d 蒸留は、アプリのユーザーが明示的に実行するすべてのアクションです。そして、Bonsai は各ステップで、どのような情報が組み立てられ、特定の応答のためにモデルに送信されたかを正確に示すことができるため、「なぜ AI はそう言ったのか?」常に具体的な答えがあります。
Bonsai は、ブラウザにインストールして開く完成したチャット アプリではありません。開発者が独自の製品に組み込むライブラリです。チャットベースのアプリ (サポート ツール、リサーチ アシスタント、内部コパイロット) を構築していて、構築ブロックとして分岐、マージ可能なコンテキスト、耐久性のあるメモリが必要な場合、Bonsai は基礎となる要素 (ツリー/分岐/マージ/蒸留ロジックに加えて、データの保存場所、呼び出す LLM、および過去の知識の検索方法) を交換可能なアダプターを提供するため、その配管を自分で構築する必要はありません。
ツリー モデル — すべてのプロジェクトは 1 つのメイン ブランチとして始まります。任意のメッセージで分岐すると、独立して探索できる新しい分岐が作成されます。
ContextPacket — モデルにプロンプ​​トを送信する前に (またはその代わりに)、含まれるメッセージ、マージ、Wiki ページの正確なパケットを検査できます。隠れた取得手順はありません。
マージ — 明示的なユーザー承認アクションとして、ブランチのレビュー済みの概要をその親ブランチに戻します。
蒸留 — ブランチまたはマージのトランスクリプトを、後の会話で取得できる耐久性のある Markdown Wiki ページに変換します。
Wiki — 耐久性があり人間が判読できるナレッジ ベースは、読みやすいフロントマター付きのプレーンな Markdown ファイルに書き込み、Grep を実行したり、Bonsai の外部で完全にバックアップしたりすることができます。
取得 — Wiki を検索するためのプラグイン可能なインターフェイス (すぐに使用できる全文検索。埋め込みベースの検索は後で置き換えることができます)。
それぞれの詳細な説明については概念ページまたはドキュメント サイトを、実際の例についてはクイックスタートを参照してください。
盆栽

@bonsai/core 上に構築されたプロジェクト: 右側のブランチ グラフは、独立して探索された 3 つのブランチ ( i-like-dogs 、 i-like-hotdogs 、 i-loke-motorb... ) を持つ main を示しています。そのうちの 2 つは、レビュー済みの概要とワンクリックの「Wiki への蒸留」アクションですでにマージされています。一方、左側のチャットには、 main 自体の履歴と明示的にマージされた内容のみが表示されます。
フェーズ1の足場。まだパッケージは公開されていません。
パッケージ/
core/ # @bonsai/core — フレームワークに依存しないドメイン (ストレージ/LLMProvider/WikiStore/Retriever インターフェイス)
storage-postgres/ # @bonsai/storage-postgres — Postgres の移行 + リポジトリ + FTS リトリーバー
Provider-openai/ # @bonsai/provider-openai — OpenAI 互換のチャット/ストリーミング プロバイダー
wiki-fs/ # @bonsai/wiki-fs — ディスク上のマークダウン WikiStore
server/ # @bonsai/server — オプションの薄い HTTP レイヤー
examples/ # サンプルエンベッダー (後のフェーズで追加)
前提条件
pnpm 9.15.4 ( packageManager 経由で固定)
pnpm install # ワークスペースの依存関係をインストールする
pnpm typecheck # tsc -b プロジェクト参照を介してすべてのパッケージにわたって
pnpm build # ビルド スクリプトを定義するすべてのパッケージをビルドします
pnpm test # ワークスペース全体でテストを実行します
pnpm lint # ESLint (パッケージ/コアの「no-restricted-imports」境界ルールを含む)
pnpm depcruise # dependency-cruiser — コア -> アダプターのインポートを禁止します
pnpm border:verify # 陽性テスト: フィクスチャは lint + depcruise によって拒否される必要があります
pnpm publint # publint --strict パッケージごと (リリースゲート)
pnpm attw # パッケージごとにタイプが間違っています (リリース ゲート、ESM のみのプロファイル)
pnpm release:gate # build + publint + attw (公開前に実行)
pnpm changeset # リリースのチェンジセットエントリを追加します
pnpm changeset:status # 保留中の変更セットを表示
pnpm docs:dev # ドキュメント サイトをローカルで実行します (Astro Starlight)
pnpm docs:build # 静的ドキュメント サイトを構築します

docs/dist の下
pnpm docs:preview # 構築されたドキュメント サイトをプレビューする
ドキュメントサイト
公開ドキュメントは、Astro Starlight で構築された独自の pnpm ワークスペース パッケージ ( @bonsai/docs 、プライベート) として docs/ の下に存在します。 API リファレンスは、 TypeDoc + starlight-typedoc を介して各パッケージの src/index.ts から自動生成されます。ランディング ページ、クイックスタート、コンセプト ページ、およびレシピは、現在の @bonsai/* API サーフェスに対して検証されます。
@bonsai/core は、契約によりフレームワークとアダプターに依存しません。 2 つの層がその境界を守ります。
ESLint no-restricted-imports (packages/core/** にスコープ) は、react、react-dom、next/*、tailwindcss、pg / postgres / sqlite*、fs /node:fs、net、node http クライアント、およびプロバイダー SDK を禁止します。
dependency-cruiser は、packages/core → package/{storage-*, Provider-*, wiki-fs,server} (解決されたパスと @bonsai/* モジュール名の両方) を禁止します。
pnpm border:verify は、packages/core/src/__fixtures__/ にある意図的な違反フィクスチャに対して両方のツールを実行し、どちらかのツールがそれらを受け入れると失敗します。 CI はプッシュ/PR ごとに実行します。
コンテキスト バンチングとコンテキスト Wiki 蒸留のためのライブラリ
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A library for context banching and context wiki distilation - Joaoha/Bonsai

GitHub - Joaoha/Bonsai: A library for context banching and context wiki distilation · GitHub
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
Joaoha
/
Bonsai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
25 Commits 25 Commits .changeset .changeset .github/ workflows .github/ workflows docs docs examples/ minimal-node examples/ minimal-node packages packages scripts scripts test/ fixtures/ boundary-violation test/ fixtures/ boundary-violation .dependency-cruiser.cjs .dependency-cruiser.cjs .dependency-cruiser.fixture.cjs .dependency-cruiser.fixture.cjs .gitignore .gitignore LICENSE LICENSE README.md README.md eslint.config.mjs eslint.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json View all files Repository files navigation
Bonsai is a toolkit for building chat apps where conversations with an AI don't have to stay a single, linear thread.
Most AI chat apps force every conversation down one straight line: you ask, it answers, you ask again, and the whole history piles up in order. If you want to explore a tangent — "what if we tried a different approach?" — you either derail the conversation or start over from scratch and lose the context you'd already built up.
Bonsai lets a conversation branch. From any point in a chat, you can branch off to explore an idea, without disturbing the original thread. If the branch turns out useful, you can merge a summary of it back into the main conversation. If it's a dead end, you just abandon it — nothing about it leaks back in. And when a conversation (or a branch) produces something worth keeping long-term, you can distill it into a wiki page that future conversations can draw on, instead of re-explaining it every time.
Crucially, Bonsai never does any of this automatically behind your back. Branching, merging, and distilling are all actions your app's user explicitly takes — and at every step, Bonsai can show you exactly what information was assembled and sent to the model for a given reply, so "why did the AI say that?" always has a concrete answer.
Bonsai is not a finished chat app you install and open in a browser — it's a library that developers embed into their own product. If you're building a chat-based app (a support tool, a research assistant, an internal copilot) and want branching, mergeable context, and durable memory as building blocks, Bonsai gives you the underlying pieces — the tree/branch/merge/distill logic, plus swappable adapters for where you store data, which LLM you call, and how you search past knowledge — so you don't have to build that plumbing yourself.
Tree Model — every project starts as one main branch; branching off at any message creates a new branch that can be explored independently.
ContextPacket — before (or instead of) sending a prompt to the model, you can inspect the exact packet of messages, merges, and wiki pages that would be included — no hidden retrieval step.
Merge — land a reviewed summary of a branch back onto its parent branch, as an explicit, user-approved action.
Distill — turn a branch or merge's transcript into a durable Markdown wiki page that later conversations can retrieve.
Wiki — the durable, human-readable knowledge base distill writes into: plain Markdown files with frontmatter, easy to read, grep, or back up outside of Bonsai entirely.
Retrieval — a pluggable interface for searching that wiki (full-text search out of the box; embeddings-based search can be swapped in later).
Read the Concepts pages or the docs site for the full explanation of each, or the Quickstart for a working example.
A Bonsai project, built on @bonsai/core : the branch graph on the right shows main with three branches ( i-like-dogs , i-like-hotdogs , i-loke-motorb... ) explored independently, two of them already merged back with a reviewed summary and a one-click "Distill to Wiki" action, while the chat on the left only ever sees main 's own history plus what was explicitly merged into it:
Phase 1 scaffold. No package is published yet.
packages/
core/ # @bonsai/core — framework-agnostic domain (Storage/LLMProvider/WikiStore/Retriever interfaces)
storage-postgres/ # @bonsai/storage-postgres — Postgres migrations + repositories + FTS retriever
provider-openai/ # @bonsai/provider-openai — OpenAI-compatible chat/streaming provider
wiki-fs/ # @bonsai/wiki-fs — Markdown-on-disk WikiStore
server/ # @bonsai/server — thin optional HTTP layer
examples/ # example embedders (added in later phases)
Prerequisites
pnpm 9.15.4 (pinned via packageManager )
pnpm install # install workspace dependencies
pnpm typecheck # tsc -b across all packages via project references
pnpm build # build all packages that define a build script
pnpm test # run tests across the workspace
pnpm lint # ESLint (incl. `no-restricted-imports` boundary rules on packages/core)
pnpm depcruise # dependency-cruiser — forbid core -> adapter imports
pnpm boundary:verify # positive test: fixture must be rejected by lint + depcruise
pnpm publint # publint --strict per package (release gate)
pnpm attw # arethetypeswrong per package (release gate, ESM-only profile)
pnpm release:gate # build + publint + attw (run before publishing)
pnpm changeset # add a changeset entry for a release
pnpm changeset:status # show pending changesets
pnpm docs:dev # run the docs site locally (Astro Starlight)
pnpm docs:build # build the static docs site under docs/dist
pnpm docs:preview # preview a built docs site
Docs site
Public documentation lives under docs/ as its own pnpm workspace package ( @bonsai/docs , private), built with Astro Starlight. The API Reference is auto-generated from each package's src/index.ts via TypeDoc + starlight-typedoc . Landing page, quickstart, concept pages, and recipes are verified against the current @bonsai/* API surface.
@bonsai/core is framework- and adapter-agnostic by contract. Two layers defend that boundary:
ESLint no-restricted-imports (scoped to packages/core/** ) bans react , react-dom , next/* , tailwindcss , pg / postgres / sqlite* , fs / node:fs , net , node http clients, and provider SDKs.
dependency-cruiser forbids packages/core → packages/{storage-*, provider-*, wiki-fs, server} (both by resolved path and by @bonsai/* module name).
pnpm boundary:verify runs both tools against intentional violation fixtures under packages/core/src/__fixtures__/ and fails if either tool accepts them. CI runs it on every push/PR.
A library for context banching and context wiki distilation
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
