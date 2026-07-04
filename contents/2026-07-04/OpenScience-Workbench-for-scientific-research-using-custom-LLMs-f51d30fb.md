---
source: "https://github.com/synthetic-sciences/openscience"
hn_url: "https://news.ycombinator.com/item?id=48786827"
title: "OpenScience: Workbench for scientific research using custom LLMs"
article_title: "GitHub - synthetic-sciences/openscience: The open-source AI workbench for scientific research · GitHub"
author: "ignoramous"
captured_at: "2026-07-04T16:54:15Z"
capture_tool: "hn-digest"
hn_id: 48786827
score: 1
comments: 0
posted_at: "2026-07-04T16:50:25Z"
tags:
  - hacker-news
  - translated
---

# OpenScience: Workbench for scientific research using custom LLMs

- HN: [48786827](https://news.ycombinator.com/item?id=48786827)
- Source: [github.com](https://github.com/synthetic-sciences/openscience)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T16:50:25Z

## Translation

タイトル: OpenScience: カスタム LLM を使用した科学研究用ワークベンチ
記事のタイトル: GitHub - 合成科学/オープンサイエンス: 科学研究のためのオープンソース AI ワークベンチ · GitHub
説明: 科学研究用のオープンソース AI ワークベンチ - 合成科学/オープンサイエンス

記事本文:
GitHub - 合成科学/オープンサイエンス: 科学研究用のオープンソース AI ワークベンチ · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
総合科学
/
オープンサイエンス
公共
通知
あなたは署名しているに違いありません

通知設定を変更するには編集してください
追加のナビゲーション オプション
コード
合成科学/オープンサイエンス
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .github .github .husky .husky .openscience .openscience アセット アセット バックエンド/ cli バックエンド/ cli フロントエンド フロントエンド ツール ツール .editorconfig .editorconfig .gitignore .gitignore .gitleaksignore .gitleaksignore .prettierignore .prettierignore AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md bun.lock bun.lock bunfig.toml bunfig.toml install install package.json package.json tsconfig.json tsconfig.json Turbo.json Turbo.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
科学研究用のオープンソース AI ワークベンチ
目標を与えてください。文献を読み、コードを書いて実行し、実験を実行して、発見したことを書き留めます。
インストール · クイックスタート · ドキュメント · アトラス
OpenScience は、科学研究用の AI ワークベンチです。目標を与えると、有能な共同研究者が行うように研究ループを通じて機能します。重要な論文を読み、仮説を立て、コードを書いて実行し、実際のコンピューティングで実験を実行し、主要な科学データベースにクエリを実行し、結果を書き込みます。これはブラウザ内のワークスペースとして実行され、独自の API キーを使用して、Anthropic、OpenAI、Google、その他多数のプロバイダーのフロンティア モデルまたはオープンウェイト モデルと連携します。アカウントは必要ありません。
これはモデルに依存しないオープンソースであり、機械学習、生物学、物理学、化学の実際の作業を行うために構築されています。
ループ全体を実行します。文献レビュー、仮説、コード、実験、分析、執筆を 1 つの連続したシーケンスで行う

セッション。
研究エージェント。デフォルトの研究エージェントに加えて、生物学、物理学、および ml の専門家、批評および文献レビューのサブエージェント、および読み取り専用のプラン モードが含まれます。
250以上のスキル。トレーニング (DeepSpeed、PEFT、TRL)、評価、データセット作業、分子生物学および臨床生物学、化学情報学、論文および LaTeX、図、クラウド コンピューティング (Modal、Tinker など)。
ツールとしての科学データベース。 UniProt、PDB、Ensembl、ChEMBL、PubChem、arXiv、OpenAlex、Semantic Sc​​holar、その他約 30 種類がエージェントから直接クエリ可能です。
本物のワークスペース。ファイル ツリー、エディター、ターミナル、セッション履歴、分子、構造、ゲノム、プロットのインライン レンダリングを備えたブラウザー UI。
拡張可能。 LSP 統合、MCP サーバー、プラグイン、カスタム エージェントとコマンド、TypeScript SDK。
npm でインストールし、ワークスペースを開きます。
npm install -g @synsci/openscience
オープンサイエンス
コマンドは openscience で、ブラウザーでワークスペースが開きます。グローバルにインストールしたくない場合は、npx synci は同じことを 1 つのステップで実行します。
npx同期
プラットフォーム バイナリも GitHub リリースに添付されています。
任意のプロバイダー ( ANTHROPIC_API_KEY 、 OPENAI_API_KEY 、 GEMINI_API_KEY など) から API キーを設定し、ワークスペースを開始します。
エクスポート ANTHROPIC_API_KEY=sk-ant-...
オープンサイエンス
openscience はブラウザーでワークスペースを開きます。キーはマシン上に残り、リクエストはプロバイダーに直接送信されます。 [資格情報] パネルからキーを追加し、モデル セレクターからモデルを選択することもできます。特定のプロジェクトでワークスペースを開くには:
オープンサイエンス ~ /code/my-project
アトラス
Atlas は、Synthetic Sciences のマネージド プラットフォームです。プリペイド ウォレットから請求される厳選されたフロンティア モデルのセットが提供されるため、プロバイダーごとのキーは必要なく、さらに永続的なリサーチ グラフとクラウド コンピューティングも利用できます。オープンサイエンス

e は Atlas で動作しますが、Atlas を必要とすることはありません。
オープンサイエンス接続ログイン
Bring-your-own-key の使用は常に無料であり、ゲートされることはありません。 Atlas はサービスを提供するモデルのみを測定します。
OpenScience は、ワークスペース UI、エージェント ランタイム、およびツール層をホストするローカル サーバーを実行します。エージェントはリサーチ ハーネスを使用して計画を立て、ツール (シェル、エディタ、LSP、MCP サーバー、科学コネクタ、およびスキル) を呼び出し、その作業をブラウザにストリーミングして返します。モデルはリクエストごとにルーティングされるため、他に何も変更せずにプロバイダーを切り替えたり、ローカル モデルを実行したりできます。セッション、成果物、来歴はディスクに保存され、リンクとして共有できます。
グローバル構成は ~/.config/openscience/openscience.json にあります。プロジェクト構成は、リポジトリ ルート (スキーマ) の openscience.json または .openscience/ ディレクトリに存在します。カスタム エージェント、コマンド、ツール、プラグイン、およびテーマは、これらのディレクトリからロードされます。
バンインストール
パン開発者
バンランタイプチェック
bun run --cwd バックエンド/cli テスト
bun run --cwd バックエンド/cli ビルド
bun dev はソースからワークスペースを実行し、bun run --cwd backend/cli build はプラットフォーム バイナリを生成します。
システムがどのように連携するかについては ARCHITECTURE.md、貢献方法については CONTRIBUTING.md、スタイル ガイドについては AGENTS.md、コミュニティ標準については CODE_OF_CONDUCT.md を参照してください。
エージェントはサンドボックス化されていません。権限システムにより、エージェントが何を行っているかを常に把握できます。それは隔離境界ではありません。分離が必要な場合は、コンテナーまたは VM 内で実行します。プロバイダーと同期された資格情報はサブプロセス環境から除外され、出力から編集されます。脆弱性を報告するには、 SECURITY.md を参照してください。
Apache ライセンス 2.0。 「ライセンスと通知」を参照してください。
OpenScience は独立したプロジェクトです。 Anthropic との提携、承認、後援はありません。 「Claude」は Anthropic、PBC の商標であり、ここでは目的のためにのみ使用されます。

スクライブの互換性。
科学研究用のオープンソース AI ワークベンチ
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
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

The open-source AI workbench for scientific research - synthetic-sciences/openscience

GitHub - synthetic-sciences/openscience: The open-source AI workbench for scientific research · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
synthetic-sciences
/
openscience
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
synthetic-sciences/openscience
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .github .github .husky .husky .openscience .openscience assets assets backend/ cli backend/ cli frontend frontend tooling tooling .editorconfig .editorconfig .gitignore .gitignore .gitleaksignore .gitleaksignore .prettierignore .prettierignore AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md bun.lock bun.lock bunfig.toml bunfig.toml install install package.json package.json tsconfig.json tsconfig.json turbo.json turbo.json View all files Repository files navigation
The open-source AI workbench for scientific research
Give it a goal. It reads the literature, writes and runs code, runs the experiments, and writes up what it found.
Install · Quickstart · Docs · Atlas
OpenScience is an AI workbench for scientific research. You give it a goal, and it works through the research loop the way a capable collaborator would. It reads the papers that matter, forms a hypothesis, writes and runs code, runs experiments on real compute, queries the major scientific databases, and writes up the result. It runs as a workspace in your browser and works with any frontier or open-weight model from Anthropic, OpenAI, Google, and dozens of other providers, using your own API keys. No account is required.
It is model-agnostic, open source, and built to do real work in machine learning, biology, physics, and chemistry.
Runs the whole loop. Literature review, hypothesis, code, experiment, analysis, and write-up, in one continuous session.
Research agents. A research agent by default, plus biology , physics , and ml specialists, with critique and literature-review sub-agents and a read-only plan mode.
250+ skills. Training (DeepSpeed, PEFT, TRL), evaluation, dataset work, molecular and clinical biology, cheminformatics, papers and LaTeX, figures, and cloud compute (Modal, Tinker, and others).
Scientific databases as tools. UniProt, PDB, Ensembl, ChEMBL, PubChem, arXiv, OpenAlex, Semantic Scholar, and around 30 more, queryable directly by the agent.
A real workspace. A browser UI with a file tree, an editor, a terminal, session history, and inline rendering for molecules, structures, genomes, and plots.
Extensible. LSP integration, MCP servers, plugins, custom agents and commands, and a TypeScript SDK.
Install with npm, then open the workspace:
npm install -g @synsci/openscience
openscience
The command is openscience , and it opens the workspace in your browser. If you would rather not install it globally, npx synsci does the same thing in a single step:
npx synsci
Platform binaries are also attached to GitHub Releases .
Set an API key from any provider ( ANTHROPIC_API_KEY , OPENAI_API_KEY , GEMINI_API_KEY , and so on) and start the workspace:
export ANTHROPIC_API_KEY=sk-ant-...
openscience
openscience opens the workspace in your browser. Your keys stay on your machine and requests go straight to the provider. You can also add keys from the Credentials panel and pick a model from the model selector. To open the workspace in a specific project:
openscience ~ /code/my-project
Atlas
Atlas is Synthetic Sciences' managed platform. It gives you a curated set of frontier models billed from a prepaid wallet, so you do not need per-provider keys, plus a persistent research graph and cloud compute. OpenScience works with Atlas but never requires it.
openscience connect login
Bring-your-own-key usage is always free and is never gated. Atlas only meters the models it serves.
OpenScience runs a local server that hosts the workspace UI, the agent runtime, and the tool layer. The agent plans with a research harness, calls tools (shell, editor, LSP, MCP servers, scientific connectors, and skills), and streams its work back to the browser. Models are routed per request, so you can switch between providers or run local models without changing anything else. Sessions, artifacts, and provenance are stored on disk and can be shared as links.
Global config lives in ~/.config/openscience/openscience.json . Project config lives in openscience.json or a .openscience/ directory at the repo root ( schema ). Custom agents, commands, tools, plugins, and themes load from those directories.
bun install
bun dev
bun run typecheck
bun run --cwd backend/cli test
bun run --cwd backend/cli build
bun dev runs the workspace from source, and bun run --cwd backend/cli build produces the platform binaries.
See ARCHITECTURE.md for how the system fits together, CONTRIBUTING.md for how to contribute, AGENTS.md for the style guide, and CODE_OF_CONDUCT.md for community standards.
The agent is not sandboxed. The permission system keeps you aware of what the agent is doing; it is not an isolation boundary. Run inside a container or VM if you need isolation. Provider and synced credentials are filtered out of subprocess environments and redacted from output. To report a vulnerability, see SECURITY.md .
Apache License 2.0. See LICENSE and NOTICE .
OpenScience is an independent project. It is not affiliated with, endorsed by, or sponsored by Anthropic. "Claude" is a trademark of Anthropic, PBC, used here only to describe compatibility.
The open-source AI workbench for scientific research
Apache-2.0 license
Code of conduct
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
