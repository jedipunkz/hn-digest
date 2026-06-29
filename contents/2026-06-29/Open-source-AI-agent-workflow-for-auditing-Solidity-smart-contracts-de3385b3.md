---
source: "https://github.com/chain-shield/ai-agent-audit"
hn_url: "https://news.ycombinator.com/item?id=48726182"
title: "Open-source AI agent workflow for auditing Solidity smart contracts"
article_title: "GitHub - chain-shield/ai-agent-audit · GitHub"
author: "chainshieldai"
captured_at: "2026-06-29T23:21:12Z"
capture_tool: "hn-digest"
hn_id: 48726182
score: 2
comments: 0
posted_at: "2026-06-29T22:29:32Z"
tags:
  - hacker-news
  - translated
---

# Open-source AI agent workflow for auditing Solidity smart contracts

- HN: [48726182](https://news.ycombinator.com/item?id=48726182)
- Source: [github.com](https://github.com/chain-shield/ai-agent-audit)
- Score: 2
- Comments: 0
- Posted: 2026-06-29T22:29:32Z

## Translation

タイトル: Solidity スマート コントラクトを監査するためのオープンソース AI エージェント ワークフロー
記事のタイトル: GitHub -chain-shield/ai-agent-audit · GitHub
説明: GitHub でアカウントを作成して、chain-shield/ai-agent-audit の開発に貢献します。

記事本文:
GitHub - チェーンシールド/ai-agent-audit · GitHub
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
チェーンシールド
/
ai-エージェント-監査
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

オプション
コード
開発 ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーとファイル
397 コミット 397 コミット .github .github サンプル サンプル スクリプト スクリプト src src テスト テスト 検証-スリーショット 検証-スリーショット ベンダー/ リグコア ベンダー/ リグコア .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md AUTO_AUDIT_CONTEXT_ARCHITECTURE.md AUTO_AUDIT_CONTEXT_ARCHITECTURE.md BENCHMARKING_WHITEPAPER.md BENCHMARKING_WHITEPAPER.md CODE4RENA_CONTEXT_PHASE2_SPEC.md CODE4RENA_CONTEXT_PHASE2_SPEC.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile.rust Dockerfile.rust Dockerfile.rust.v2 Dockerfile.rust.v2 ライセンス ライセンス POC_DEPRECATION_REMOVAL_PLAN.md POC_DEPRECATION_REMOVAL_PLAN.md README.md README.md SECURITY.md SECURITY.md VERIFY_CHECKLIST.md VERIFY_CHECKLIST.md Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI Agent Audit は、Solidity リポジトリの AI 支援セキュリティ レビュー用の Rust コマンドライン ツールです。
Solidity および EVM ベースのコードベースのセキュリティ脆弱性を発見します
重複を排除し、結果を検証します
検証された高/中結果ごとに実行可能な PoC を生成
マークダウンで検証された結果ごとに専門的な監査レポートを作成する
私はこのツールを使用して Code4rena コンテストに参加しましたが、結果は有望なものでした。
https://code4rena.com/@saraswati
リポジトリはパブリックベータ版です。これは専門家のレビューを加速することを目的としており、手動監査に代わるものではありません。
リポジトリ ソース、ドキュメント、および派生コンテキストは、構成したサードパーティ LLM プロバイダーに送信されます。
現在のデフォルトの監査パイプラインは、OpenAI アクセスに ChatGPT/Codex OAuth を使用し、 gpt-5.5 でアクティブなレビュー フローを実行します。重複排除ヘルパーは、低い推論で gpt-5.4 を使用します。
コーデックス

は、現在の検証ワークフローおよびオペレーティング モデルに推奨されるデフォルト パスです。
スタートアップは、必要に応じて 1 回限りの ChatGPT サインインを実行し、トークンの有効期限が切れるまで、その後の実行でキャッシュされたセッションを再利用します。
OPENAI_API_KEY は、 AI_AGENT_AUDIT_OPENAI_BACKEND=api を設定することで、Rust OpenAI 呼び出しの二次フォールバックとしてサポートされます。
ANTHROPIC_API_KEY 、 GEMINI_API_KEY / GOOGLE_AI_API_KEY 、および DEEPSEEK_API_KEY は引き続きエージェント層でサポートされていますが、デフォルトのレビュー パスでは必要ありません。
OpenAI/Codex で検証/レポートを維持しながら、パターン、アクター、および不変式に Google AI を使用したい場合は、src/config.rs のデフォルトを変更することで Discovery スタイルの実行を Gemini に切り替えることができます。
PoC 生成と PoC 検証は、主要な検証ワークフローである validation-three-shot を通じてサポートされます。
デフォルトで ~/Desktop/Audit の下に Foundry または Hardhat リポジトリをクローンして構築します。
監査スコープとプロトコルのドキュメントを README/構成済みエントリ ファイルから Audit-docs/ に生成します。
静的解析が成功した場合は、Slither 派生のコール グラフとセマンティック データを使用します。
Solidity ソースから継承およびインターフェイス実装のインデックスを構築します。
スコープ内のコントラクトごとにコンテキストに応じたコードブロックを生成します。
パターン ライブラリ、不変プロンプト、およびアクター指向のコンテキストを使用して、候補結果を発見します。
レポート出力を作成する前に、結果を検証して重複を排除します。
ローカルの SQLite 状態を .ai-agent-audit/ に保存します。
スマートコントラクト監査人およびセキュリティ研究者。
Solidity コードベースの内部レビューを行うプロトコル チーム。
外部モデルプロバイダーとの共有が許可されているリポジトリ上で AI 支援監査ワークフローを実験するエンジニア。
このプロジェクトはホスト型サービスではなく、あらゆる言語に対応する汎用 SAST スキャナーでも、ヒューマの代替品でもありません。

検証。
Rust の安定したツールチェーン。 Rust-lang.org/tools/install からインストールします。
ギット。 git-scm.com/downloads からインストールします。
ずるずると。 github.com/crytic/slither からインストールします。
Foundry リポジトリの Foundry ( forge )。 book.getfoundry.sh/getting-started/installation からインストールします。
Node.js 18 以降と、JavaScript/Hardhat リポジトリ用の npm / npx 、Yarn、pnpm、または Bun。 Nodejs.org から Node.js をインストールします。ターゲット リポジトリが bun.lock / bun.lockb を使用している場合は、bun.sh から Bun をインストールします。
プライベート GitHub リポジトリの場合はオプションの GITHUB_TOKEN。
このリポジトリのクローンを作成して入力します。
git clone https://github.com/chain-shield/ai-agent-audit.git
cd ai-エージェント-監査
テンプレートからローカル環境ファイルを作成します。
cp .env.example .env
.env を編集し、実際に必要な値を設定します。
RUST_LOG=情報
AI_AGENT_AUDIT_OPENAI_BACKEND=コーデックス
# Codex がコストのデフォルトとして推奨されます。オプションの Rust API フォールバック:
# AI_AGENT_AUDIT_OPENAI_BACKEND=api
# OPENAI_API_KEY=your_openai_api_key
# オプションの非 OpenAI プロバイダー キー:
# ANTHROPIC_API_KEY=...
# GEMINI_API_KEY=...
# DEEPSEEK_API_KEY=...
Codex セッションがまだキャッシュされていない場合は、Codex を使用した最初の実行では、ChatGPT でサインインするように求められます。その後、セッションは期限が切れるまで自動的に再利用されます。 AI_AGENT_AUDIT_OPENAI_BACKEND=api を設定すると、Rust OpenAI 呼び出しは代わりに OPENAI_API_KEY を使用し、Codex サインインを必要としません。
サンプル構成をコピーし、Solidity リポジトリを指定します。
cp 例/audit-config.example.yaml Audit-config.yaml
最小限の例:
リポジトリ: " https://github.com/example/protocol.git "
Audit_type : "クライアント"
コードフォルダー:
-「ソース」
ツールをビルドして実行します。
カーゴビルド --release
カーゴ実行 --release -- --config Audit-config.yaml
単純な実行に YAML を使用したくない場合は、次のようにします。
カーゴ実行 --release -- https://github.com/example/protocol.git --audit-t

そうだクライアント
src/ ではなくcontracts/の下にコントラクトを保持するリポジトリの場合は、それに応じて code_folders を設定します。
Rust のメイン実行では、初期監査アーティファクトが生成され、デフォルトで、すぐに実行できる検証 3 ショット ジョブが生成されます。この検証ワークフローでは、より詳細なフィルタリング、PoC 生成、PoC 検証、レポートの強化が行われます。
ターゲット リポジトリがプライベートの場合は、ツールを実行する前に GITHUB_TOKEN を設定します。クローン パスは、GitHub HTTPS URL のトークンを使用します。
エクスポートGITHUB_TOKEN=...
構成
--config <file> は YAML をロードし、明示的な CLI フラグが YAML 値をオーバーライドします。現在のサンプル ファイルは、examples/audit-config.example.yaml にあります。
内部監査またはクライアント スタイルの監査にはクライアントを使用します。重大度の処理とレポート言語をこれらのプラットフォームに合わせて調整したい場合は、コンテスト値を使用します。 C4 バグ報奨金プログラムには Code4renaBounty を使用します。このプログラムでは、実行可能な PoC で現在悪用可能な重大/高の問題のみが提出候補となります。実行で Immunefi プログラム ページからリポジトリ/ドキュメント/スコープを取得し、Immunefi スタイルの影響ルールに対して検証する必要がある場合は、ImmunefiBugBounty を使用します。
immunofi_bounty / --immunefi-bounty
Immunefi 報奨金ページからリポジトリ、ドキュメント、スコープを取得します。これは ImmunefiBugBounty の標準エントリポイントです。
code4rena_bounty / --code4rena-bounty
Code4rena の報奨金ページからリポジトリ、ドキュメント、スコープを取得します。これは、ツールで報奨金コンテキストを収集する場合の Code4renaBounty の標準エントリポイントです。
code4rena_contest_repo / --code4rena-contest-repo
Code4rena コンテスト コンテキストとコンテスト GitHub リポジトリ URL からの V12 ルックアップ サポートを追加します。
code4rena_contest_url / --code4rena-contest-url
code4rena_contest_repo のエイリアス。 2 つのうち 1 つだけを指定してください。
# Immunefi由来のラン
Audit_type : " ImmunefiBugBounty "
免疫フィボ

unty : " https://immunefi.com/bug-bounty/example/information/ "
# Code4rena の賞金由来の実行
Audit_type : " Code4renaBounty "
code4rena_bounty : " https://code4rena.com/bounties/example "
# Code4rena コンテストは追加のコンテスト コンテキストで実行されます
リポジトリ: " https://github.com/example/protocol.git "
Audit_type : " Code4rena "
code4rena_contest_repo : " https://github.com/code-423n4/2026-01-example "
YAML および CLI フィールド
enum のような値の場合、YAML は Code4renaBounty 、 ImmunefiBugBounty 、 HardhatYarn などの Rust スタイルの名前を使用します。 CLI フラグは、 --builder のhardhat-yarn など、実際の --help スペルを使用します。
Custom_doc 、 Audit_scope 、scoped_files 、および monorepo_folders は、ツールを実行しているマシン上で提供されたローカル ファイルから読み取られます。
subfolder 、 code_folders 、 doc_folder 、および exclude_folders は、クローンされたターゲット リポジトリを基準として解釈されます。
手動のcustom_doc、audit_scope、またはscoped_filesが提供されていない場合、ツールはaudit-docs/に <repo-folder>-docs.md 、 <repo-folder>-scope.md 、および <repo-folder>-scope.txt を生成します。
生成されたファイル名プレフィックスには、 2026-04-monetrix などの日付/コンテスト プレフィックスを含む、複製されたリポジトリ フォルダーの ID が保持されます。
context.force_regenerate は、生成されたコンテキストに対してデフォルトで true に設定されます。コンテキスト ブロックが存在しない場合、3 つの手動コンテキスト ファイルをすべてすでに提供しているレガシー YAML はそのまま残されます。
最小限の生成されたコンテキスト構成:
コンテキスト:
ファイル:
- README.md
URL: []
v12_url : " 自動 "
Output_dir : " 監査ドキュメント "
Force_regenerate : true
ファイルごとの最大トークン数: 5000
ジェネレーターは、複製されたリポジトリが存在する場合、そこからscope.txtをコピーします。 scope.txt が存在しない場合は、エントリ コンテキスト ファイルからスコープ内の Solidity パスを抽出します。デフォルトでは、エントリ ファイルは README.md のみです。 context.files はエントリ ファイルを追加または置換できます。それらのエントリ内にあるリンク

ファイルは第 2 レベルの候補として扱われ、対象範囲、既知の問題、プロトコル文書、以前の監査、または Code4rena V12 レポートに関連すると思われる場合にのみ取得されます。 Code4renaBounty の場合、ジェネレーターは Code4rena の報奨金ガイドと報奨金基準もフェッチし、グローバルな報奨金のスコープ外ルールを保持し、コントラクト名のみのスコープ テーブルをローカルの Solidity 定義にマッピングできます。フェッチされた第 2 レベルのページは、それ以上のリンクを生成しません。生成されたscope.mdまたはdocs.mdが max_tokens_per_file を超える場合、Codexはそれを制限まで要約します。ジェネレーターは最終的なマークダウンを黙って切り捨てることを拒否します。
ポク :
allow_fork : true
優先フォーク : true
rpc_env :
イーサリアム-メインネット: " MAINNET_RPC_URL "
ベースメインネット: " BASE_RPC_URL "
poc ブロックは YAML のみです。これは主に、Immunefi の報奨金検証など、明示的なネットワークとフォークの仮定を必要とするワークフロー用に生成された PoC ランタイム ガイダンスに影響します。
変数
必須
注意事項
ChatGPT/Codex サインイン
デフォルトのパイプラインの場合は「はい」
必要に応じて起動時に一度対話的に実行され、その後は有効期限が切れるまでローカルにキャッシュされます。
AI_AGENT_AUDIT_OPENAI_BACKEND
いいえ
デフォルトではコーデックス。 OpenAI API の直接課金を使用するには、API に設定します。
OPENAI_API_KEY
APIフォールバックのみ
AI_AGEの場合は必須

[切り捨てられた]

## Original Extract

Contribute to chain-shield/ai-agent-audit development by creating an account on GitHub.

GitHub - chain-shield/ai-agent-audit · GitHub
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
chain-shield
/
ai-agent-audit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
develop Branches Tags Go to file Code Open more actions menu Folders and files
397 Commits 397 Commits .github .github examples examples scripts scripts src src tests tests validation-three-shot validation-three-shot vendor/ rig-core vendor/ rig-core .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md AUTO_AUDIT_CONTEXT_ARCHITECTURE.md AUTO_AUDIT_CONTEXT_ARCHITECTURE.md BENCHMARKING_WHITEPAPER.md BENCHMARKING_WHITEPAPER.md CODE4RENA_CONTEXT_PHASE2_SPEC.md CODE4RENA_CONTEXT_PHASE2_SPEC.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile.rust Dockerfile.rust Dockerfile.rust.v2 Dockerfile.rust.v2 LICENSE LICENSE POC_DEPRECATION_REMOVAL_PLAN.md POC_DEPRECATION_REMOVAL_PLAN.md README.md README.md SECURITY.md SECURITY.md VERIFY_CHECKLIST.md VERIFY_CHECKLIST.md rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
AI Agent Audit is a Rust command-line tool for AI-assisted security review of Solidity repositories.
discovers security vulnerabilities in Solidity and EVM-based codebases
deduplicates and validates findings
generated runnable PoC for each validated High/Medium finding
create professional audit report for each validated finding in markdown
I used this tool to compete in Code4rena competitions and the results were encouraging:
https://code4rena.com/@saraswati
The repository is in public beta. It is meant to accelerate expert review, not replace manual auditing.
Repository source, docs, and derived context are sent to third-party LLM providers you configure.
The current default audit pipeline uses ChatGPT/Codex OAuth for OpenAI access and runs the active review flow on gpt-5.5 . Deduplication helpers use gpt-5.4 with low reasoning.
Codex is the recommended default path for the current validation workflow and operating model.
Startup performs a one-time ChatGPT sign-in if needed and reuses the cached session on later runs until the token expires.
OPENAI_API_KEY is supported as a secondary fallback for Rust OpenAI calls by setting AI_AGENT_AUDIT_OPENAI_BACKEND=api .
ANTHROPIC_API_KEY , GEMINI_API_KEY / GOOGLE_AI_API_KEY , and DEEPSEEK_API_KEY are still supported by the agent layer, but they are not required by the default review path.
Discovery-style runs can be switched to Gemini by changing the defaults in src/config.rs if you want to use Google AI for patterns, actors, and invariants while keeping verification/reporting on OpenAI/Codex.
PoC generation and PoC verification are supported through validation-three-shot , which is the primary validation workflow.
Clones and builds Foundry or Hardhat repositories under ~/Desktop/Audit by default.
Generates audit scope and protocol docs from README/configured entry files into audit-docs/ .
Uses Slither-derived call graph and semantic data when static analysis succeeds.
Builds inheritance and interface-implementation indexes from Solidity source.
Generates contextual codeblocks for each in-scope contract.
Uses pattern libraries, invariant prompts, and actor-oriented context to discover candidate findings.
Verifies and deduplicates findings before producing report output.
Stores local SQLite state in .ai-agent-audit/ .
Smart contract auditors and security researchers.
Protocol teams doing internal review of Solidity codebases.
Engineers experimenting with AI-assisted audit workflows on repos they are allowed to share with external model providers.
This project is not a hosted service, not a generic SAST scanner for every language, and not a substitute for human validation.
Rust stable toolchain. Install from rust-lang.org/tools/install .
Git. Install from git-scm.com/downloads .
Slither. Install from github.com/crytic/slither .
Foundry ( forge ) for Foundry repositories. Install from book.getfoundry.sh/getting-started/installation .
Node.js 18+ plus npm / npx , Yarn, pnpm, or Bun for JavaScript/Hardhat repositories. Install Node.js from nodejs.org . Install Bun from bun.sh if the target repo uses bun.lock / bun.lockb .
Optional GITHUB_TOKEN for private GitHub repositories.
Clone this repository and enter it.
git clone https://github.com/chain-shield/ai-agent-audit.git
cd ai-agent-audit
Create a local env file from the template.
cp .env.example .env
Edit .env and set the values you actually need:
RUST_LOG=info
AI_AGENT_AUDIT_OPENAI_BACKEND=codex
# Codex is the recommended default for cost. Optional Rust API fallback:
# AI_AGENT_AUDIT_OPENAI_BACKEND=api
# OPENAI_API_KEY=your_openai_api_key
# Optional non-OpenAI provider keys:
# ANTHROPIC_API_KEY=...
# GEMINI_API_KEY=...
# DEEPSEEK_API_KEY=...
The first Codex-backed run will prompt you to sign in with ChatGPT if there is no cached Codex session yet. After that, the session is reused automatically until expiry. If you set AI_AGENT_AUDIT_OPENAI_BACKEND=api , Rust OpenAI calls use OPENAI_API_KEY instead and do not require Codex sign-in.
Copy the example config and point it at a Solidity repository.
cp examples/audit-config.example.yaml audit-config.yaml
Minimal example:
repo : " https://github.com/example/protocol.git "
audit_type : " Client "
code_folders :
- " src "
Build and run the tool.
cargo build --release
cargo run --release -- --config audit-config.yaml
If you prefer not to use YAML for a simple run:
cargo run --release -- https://github.com/example/protocol.git --audit-type Client
For repos that keep contracts under contracts/ instead of src/ , set code_folders accordingly.
The main Rust run produces the initial audit artifacts and, by default, emits a ready-to-run validation-three-shot job. That validation workflow is where deeper filtering, PoC generation, PoC verification, and report hardening happen.
If the target repository is private, set GITHUB_TOKEN before running the tool. The clone path uses that token for GitHub HTTPS URLs.
export GITHUB_TOKEN=...
Configuration
--config <file> loads YAML, and explicit CLI flags override YAML values. The current example file lives at examples/audit-config.example.yaml .
Use Client for internal or client-style audits. Use the contest values when you want severity handling and report language aligned more closely with those platforms. Use Code4renaBounty for C4 bug bounty programs where only currently exploitable Critical/High issues with runnable PoCs should become submission candidates. Use ImmunefiBugBounty when the run should derive repo/docs/scope from an Immunefi program page and validate against Immunefi-style impact rules.
immunefi_bounty / --immunefi-bounty
Derives repo, docs, and scope from an Immunefi bounty page. This is the standard entrypoint for ImmunefiBugBounty .
code4rena_bounty / --code4rena-bounty
Derives repo, docs, and scope from a Code4rena bounty page. This is the standard entrypoint for Code4renaBounty when you want the tool to gather bounty context for you.
code4rena_contest_repo / --code4rena-contest-repo
Adds Code4rena contest context and V12 lookup support from a contest GitHub repo URL.
code4rena_contest_url / --code4rena-contest-url
Alias for code4rena_contest_repo . Provide only one of the two.
# Immunefi-derived run
audit_type : " ImmunefiBugBounty "
immunefi_bounty : " https://immunefi.com/bug-bounty/example/information/ "
# Code4rena bounty-derived run
audit_type : " Code4renaBounty "
code4rena_bounty : " https://code4rena.com/bounties/example "
# Code4rena contest run with extra contest context
repo : " https://github.com/example/protocol.git "
audit_type : " Code4rena "
code4rena_contest_repo : " https://github.com/code-423n4/2026-01-example "
YAML And CLI Fields
For enum-like values, YAML uses the Rust-style names such as Code4renaBounty , ImmunefiBugBounty , and HardhatYarn . CLI flags use the actual --help spellings, such as hardhat-yarn for --builder .
custom_doc , audit_scope , scoped_files , and monorepo_folders are read from local files you provide on the machine running the tool.
subfolder , code_folders , doc_folder , and exclude_folders are interpreted relative to the cloned target repository.
If no manual custom_doc , audit_scope , or scoped_files are provided, the tool generates <repo-folder>-docs.md , <repo-folder>-scope.md , and <repo-folder>-scope.txt in audit-docs/ .
The generated filename prefix preserves the cloned repo folder identity, including date/contest prefixes such as 2026-04-monetrix .
context.force_regenerate defaults to true for generated context. Legacy YAMLs that already provide all three manual context files are left alone when no context block is present.
Minimal generated-context config:
context :
files :
- README.md
urls : []
v12_url : " auto "
output_dir : " audit-docs "
force_regenerate : true
max_tokens_per_file : 5000
The generator copies scope.txt from the cloned repo when present. If no scope.txt exists, it extracts in-scope Solidity paths from entry context files. By default the only entry file is README.md ; context.files can add or replace entry files. Links found in those entry files are treated as second-level candidates and fetched only when they look relevant to scope, known issues, protocol documentation, prior audits, or Code4rena V12 reports. For Code4renaBounty , the generator also fetches Code4rena's bounty guide and bounty criteria, preserves global bounty out-of-scope rules, and can map contract-name-only scope tables to local Solidity definitions. Fetched second-level pages do not emit more links. If generated scope.md or docs.md exceeds max_tokens_per_file , Codex summarizes it down to the limit; the generator refuses to silently truncate final Markdown.
poc :
allow_fork : true
prefer_fork : true
rpc_env :
ethereum-mainnet : " MAINNET_RPC_URL "
base-mainnet : " BASE_RPC_URL "
The poc block is YAML-only. It mainly affects generated PoC runtime guidance for workflows that need explicit network and fork assumptions, such as Immunefi bounty validation.
Variable
Required
Notes
ChatGPT/Codex sign-in
Yes for the default pipeline
Performed interactively once at startup when needed, then cached locally until expiry.
AI_AGENT_AUDIT_OPENAI_BACKEND
No
codex by default. Set to api to use direct OpenAI API billing.
OPENAI_API_KEY
Only for API fallback
Required when AI_AGE

[truncated]
