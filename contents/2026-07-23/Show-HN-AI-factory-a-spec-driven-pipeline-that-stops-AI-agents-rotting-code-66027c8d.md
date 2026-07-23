---
source: "https://github.com/highflame-ai/ai-factory"
hn_url: "https://news.ycombinator.com/item?id=49028457"
title: "Show HN: AI-factory – a spec-driven pipeline that stops AI agents rotting code"
article_title: "GitHub - highflame-ai/ai-factory: An opinionated, config-driven AI-SDLC toolkit for AI-assisted software engineering. Skills, agents, hooks, and a role-based capability model. The 1-100 toolkit for AI coding agents: spec-driven pipeline, adversarial review bench, deterministic gates. · GitHub"
author: "sharathr"
captured_at: "2026-07-23T21:57:24Z"
capture_tool: "hn-digest"
hn_id: 49028457
score: 1
comments: 0
posted_at: "2026-07-23T21:45:05Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI-factory – a spec-driven pipeline that stops AI agents rotting code

- HN: [49028457](https://news.ycombinator.com/item?id=49028457)
- Source: [github.com](https://github.com/highflame-ai/ai-factory)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T21:45:05Z

## Translation

タイトル: Show HN: AI-factory – AI エージェントによるコードの腐敗を阻止する仕様主導のパイプライン
記事のタイトル: GitHub - highflame-ai/ai-factory: AI 支援ソフトウェア エンジニアリングのための独自の構成主導型 AI-SDLC ツールキット。スキル、エージェント、フック、および役割ベースの機能モデル。 AI コーディング エージェント用の 1-100 ツールキット: 仕様主導のパイプライン、敵対的レビュー ベンチ、決定論的ゲート。 · GitHub
説明: AI 支援ソフトウェア エンジニアリングのための、独自の構成主導型 AI-SDLC ツールキット。スキル、エージェント、フック、および役割ベースの機能モデル。 AI コーディング エージェント用の 1-100 ツールキット: 仕様主導のパイプライン、敵対的レビュー ベンチ、決定論的ゲート。 - GitHub - highflame-ai/ai-factory: 意見の強い人
[切り捨てられた]

記事本文:
GitHub - highflame-ai/ai-factory: AI 支援ソフトウェア エンジニアリングのための、独自の構成主導型 AI-SDLC ツールキット。スキル、エージェント、フック、および役割ベースの機能モデル。 AI コーディング エージェント用の 1-100 ツールキット: 仕様主導のパイプライン、敵対的レビュー ベンチ、決定論的ゲート。 · GitHub
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
アカウントを切り替えました

別のタブまたはウィンドウでカウントされます。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ハイフレイムアイ
/
アイファクトリー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット .aif/ context .aif/ context .claude .claude .github/ workflows .github/ workflows エージェント エージェント bin bin コンパイル済み コンパイル済みサンプル サンプル フック フック パック パック リファレンス リファレンス ロール ロール スキル スキル ツール ツール .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md バージョン バージョンcatalog.sh category.sh install.sh install.sh Tips.md Tips.md workspace-CLAUDE.md workspace-CLAUDE.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AIコーディングエージェントによる0→1は簡単です。 1→100 が難しいところです。コードがレビューに生き残り、デモを超えてスケ​​ールし、何か月も経っても静かにコードベースを腐らせることがありません。そこでai-factoryの登場です。
AI コーディング エージェント (スキル、エージェント、決定論的フック、ロールベースの機能モデル) の完全な仕様駆動型開発プラクティス。一度インストールすると、フォークではなく設定を通じて組織にパラメータ化されます。
より高速 — Highflame での独自の開発全体で、仕様主導のパイプラインと並列スプリント オーケストレーションにより、アドホックな AI 支援コーディングと比較して、機能/PR の出荷、バグの解決、マージまでの時間の 3 ～ 5 倍が定期的に実現されます。 (当社の内部経験 - あなたの走行距離は異なります)
より少ないスロップ — 決定論的ゲート、敵対的なマルチエージェントレビューベンチ、および検証は交渉不可の精神: レビューされていない AI 出力は main に到達できません。
より無駄のないコンテキスト

— 独自のコンテキスト ウィンドウを持つサブエージェントで大規模な調査とレビューが実行されます。メインセッションは集中力を維持します。
Claude Code のスタンドアロン — スラッシュ コマンド ( /spec → /architect → /proceed → /ship ) として 41 のスキルが、すべてのセッションにライブでシンボリックリンクされます。クローンを作成し、 ./install.sh を実行します。
codeoid の下では、エージェントごとの ID とセッション間メモリを備えた codeoid のマルチセッション ランタイムの宣言型パックと同じ方法論です: codeoid run --pack aif-sdlc 。
AIF パイプライン — 並列スプリント オーケストレーション、マルチエージェント レビュー ベンチ、および aif health CLI を備えた仕様主導の開発ライフサイクル (仕様 → アーキテクト → 検証 → 実装 → 反映 → レビュー → 出荷)。
組織ワークフロー層 - 日常のエンジニアリング スキル (課題から PR へのオーケストレーション、デバッグ、非推奨、Git 規約、セッションのハンドオフ、環境トリアージ)、レビュー エージェント、および決定論的フック。組織固有のすべては .aif/config.yml を通じて解決されます。ハードコーディングされたものは何もなく、すべてのスキルは設定キーが存在しない場合の動作を示しています。
パック レジストリ (packs/) — 宣言的な codeoid パックとしての方法論 (スキーマ: codeoid/pack@v1 )。パックはデータのみです (pack.yaml + 機能ロール + 構成)。 codeoid は、その管理されたフェーズ パイプラインを実行します。パックを 1 回提供すると、どのチームでもそれを選択できます。 Packs/README.md — 最初のパック: aif-sdlc を参照してください。
git clone https://github.com/highflame-ai/ai-factory.git
cd ai-factory && ./install.sh
install.sh はべき等であり、修復可能です。シンボリックリンクします:
また、 aif CLI shim を ~/bin ( Doctor 、 check 、 Agents render 、 COMPILE 、 Renumber ) に書き込み、 ~/.claude/aif/config.yml をスキャフォールド (デフォルトで委任オフ) し、すべての失敗に対してコピー＆ペースト可能な修正を出力する完全な環境健全性チェックである aif Doctor で終了します。 ./install.sh --dry-run p

レビュー; --クローンを移動した後に修復再スタンプします。
手順: インストール → workspace-CLAUDE.md とマシン構成 ( ~/.claude/aif/config.yml ) に記入 → claude を実行 → 各コード リポジトリで /init を実行して .aif/ 構造をブートストラップ → 次に、そのリポジトリの .aif/config.yml に記入します (オプションでスキル/プリセット/スターターから開始)。その過程で、導入組織は次の 5 つの項目をカスタマイズします (詳細は example/README.md を参照)。
リポジトリごとの .aif/config.yml — リポジトリ間の調整を選択するときに /init (skill/templates/config-template.yml から) によってスキャフォールディングされるか、プリセットから後でコピーされる — スタック、兄弟リポジトリ、および組織ワークフロー スキルを推進する org: / Environmental: / mcp: / local_stack: / regression: / tenancy: セクション。
workspace-CLAUDE.md — 常にロードされるプラットフォーム コンテキスト テンプレート (サービス マップ、認証コントラクト、規則)。
参照の記入 — MCP チートシート、回帰マーカー、既知のワート レジストリ、テナンシー/JWT チェックリスト。
設定をフックする — org.commit_prefix_regex (または AIF_COMMIT_PREFIX_REGEX ) を使用して、CI のコミット ゲートをミラーリングします。組織のキープレフィックスをhooks/secret-scan.shに追加します。
あなた自身のスキルとエージェント - 例のスケルトンとハウスルール/ 。
AIF パイプライン (仕様主導のライフサイクル)
スキル
説明
/初期化
リポジトリ内のブートストラップ .aif/ 構造
/オンボード
既存のコードベースにツールキットを採用する — サービス マップ、スターター エキスパート、規約、レビュー用の構成の草案を作成します。
/測定
ツールキットが成果をあげているかどうかを測定します — git、PR、.aif/アーティファクトからの配信/品質メトリクスと、正直な注意点
/スペック
機能リクエストから要件仕様を記述する
/建築家
アーキテクチャを設計し、要件をタスクに分割する
/検証
続行する前に AIF 位相出力を検証します。
/進む
エンドツーエンドのパイプライン: 検証 → アーキテクト → 実装 → R

反映 → レビュー → PR → まとめ
/スプリント
並列パイプライン オーケストレーター — REQ にわたる複数の /proceed 実行 ( --workflow エンジン)
/反射する
導入後の正式レビュー前の自己レビュー
/レビュー
マルチエージェントコードレビュー (正確性、品質、アーキテクチャ、テスト、セキュリティ)
/敵対者
あらゆる成果物に対する敵対的なレビュー - それが間違っていると想定し、それを証明しようとします
/カナリア
スモーク テストを使用したカナリア デプロイメント
/まとめ
機能を終了する - アーティファクトのコミット、マージ、デプロイ、更新
/バグ修正
合理化されたバグ修正ワークフロー
/ステータス//マニフェスト
飛行中のAIF作業のローカル/リモート派生ビュー
/分析する
コードベースの健全性監査
/最適化する
API のコストとパフォーマンスのスキャナー
/テンプレートドリフト
プロジェクトの .aif/templates/ とツールキットのドリフトを検出します。
主要なワークフロー:
/spec → /validate → /architect → /validate →implement → /reflect → /review → merge → /wrapup
組織のワークフロー
スキル
説明
/使用-aif
メタスキル ルーティング テーブル — SessionStart で挿入される
/発行元
問題 → マージ対応の PR オーケストレーター (分類、ルーティング、実装、出荷、CI のベビーシッター)
/機能-準備
最初に仕様レジストリにエントリが必要ですか?テストはどこへ行くのですか?
/グリル機能
建築工事前の敵対的設計レビュー
/船
マージ前のオーケストレーター — 並列レビューアーのファンアウト → ゴー/ノーゴー → デプロイ後の検証
/トリアージ-dev
構成された可観測性の順序で共有開発環境を調査します
/ソース主導型
推測するのではなく、信頼できるソース (構成、コード、MCP) を読んでください。
/増分実装
薄い垂直スライス、一度に 1 リポジトリ
/デバッグ
ローカルな再現性と根本原因。フェーズ 1 = 高速の決定論的フィードバック ループ
/非推奨
構成されたリポジトリ全体にわたる多段階の削除
/git-ワークフロー
コミットプレフィックス (構成可能な正規表現)、アトミックコミット、分岐、ワークツリー、PR フロー
/ハンドオフ
コンプ

マルチリポジトリセッションをハンドオフドキュメントに変換します
/リリースノート
git タグまたは範囲からの複数の対象ユーザー向けのリリースノート ドラフト (顧客、エンジニアリング、エグゼクティブ)
/dep-update
精査された依存関係の更新 — 変更ログのリスクレビュー、分離されたブランチ、テストで実証済み、バッチごとに 1 つの PR
/ライセンス監査
SBOM + ライセンス コンプライアンス — すべてのコンポーネントのライセンスを解決し、org.license_policy に基づいて判断します。
/rotate-シークレット
プロアクティブなシークレット ローテーション — シークレットの有効期限ゲート: インベントリ、オーバーラップ パターン ローテーション、取り消し前のヒューマン ゲート
/監査権限
.claude/settings*.json 内の危険な状態の Claude 権限付与を確認して削除します ( aif Doctor の Permissions-audit と組み合わせます)
/doc-ドリフト
変更によって古くなったドキュメントを見つけて修正します — 差分派生サーフェス、古い分類と欠落した分類
/脅威モデル
仕様/RFC からの STRIDE 脅威モデル — 資産、信頼境界、反駁された脅威、緩和パンチリスト
/エキスパート
ドメインの専門家をキュレートする — ドキュメントから抽出されたパススコープのコンテキストは、変更がそのドメインに影響を与えた場合にのみ挿入されます。
/追加-検出器
テンプレート: ドキュメントに従って、拡張可能なサービスでモジュールをスキャフォールディングします。
[切り捨てられた]
Agents/ 内の 26 個のサブエージェント。それぞれが aif エージェント レンダー (config: ~/.claude/aif/config.yml ) によってレンダリングされる階層ベースのモデル選択を持ちます。
レビューベンチ: 正確性レビューア、セキュリティ監査人、アーキテクチャレビューア、品質レビューア、コード品質監査人、テスト監査人、リフレクタ、敵対者
スキャナー/エクスプローラー: api-cost-scanner、db-perf-scanner、latency-scanner、architecture-mapper、convention-auditor、feature-tracer、integration-explorer、delegate-pre-pass
パイプライン : タスク実装者、パイプラインランナー
組織ワークフロー: security-reviewer (invariant-driven Audit)、cross-repo-impact、migration-analyzer、pr-shepherd、gemini-reviewer、sider-policy-reviewer (Cedar を使用している場合)、local-s

タックランナー、種類回帰ランナー
エキスパート — パススコープのドメインコンテキスト
/expert はドメイン エキスパートをキュレーションします。独自のドキュメントとコードから抽出された緻密なコンテキストが .aif/experts/<name>.md に保存され、変更がそのドメインに影響を及ぼした場合にのみタスクに挿入されます (エキスパートの apply_to glob と一致する)。コンテキスト エンジニアリング — React の変更には React コンテキストを、データアクセスの変更には DynamoDB コンテキストを、無関係な作業には反応しません。 /review は、マッチングする専門家を自動的に有効にします。 workspace-CLAUDE.md はあらゆる場所に当てはまる事実を保持し、参照/は横断的なチェックリストであり、レッスンは関連性でランク付けされたインシデントであり、専門家はパス範囲のドメイン知識です。 apply_to グロブは、他のアシスタント (Copilot applyTo 、Cursor グロブ) のパス スコープ ルールの将来のコンパイル キーでもあります。その発行は次のフェーズです。
ロール — 1 つの機能定義、すべてのエージェント ランタイム
role/*.yaml は、各エージェント ロールが実行できることに関する唯一の信頼できる情報源です。レビューアが {envelope: [read, grep, glob, bash], write: false} を 1 回宣言すると、aif コンパイルはそれを次のように出力します。
Claude — エージェント/*.md ツールの検証: フロントマッター (ネイティブ層。ロールのエンベロープを超えるエージェントでは --check が失敗する)、およびヘッドレス/CI セッション用のコンパイル済み/クロード/ロールごとの設定オーバーレイ
Codex — ロールごとのコンパイル済み/codex/config.toml サンドボックス プロファイル
コパイロット — コンパイル済み/copilot/pa

[切り捨てられた]

## Original Extract

An opinionated, config-driven AI-SDLC toolkit for AI-assisted software engineering. Skills, agents, hooks, and a role-based capability model. The 1-100 toolkit for AI coding agents: spec-driven pipeline, adversarial review bench, deterministic gates. - GitHub - highflame-ai/ai-factory: An opinionate
[truncated]

GitHub - highflame-ai/ai-factory: An opinionated, config-driven AI-SDLC toolkit for AI-assisted software engineering. Skills, agents, hooks, and a role-based capability model. The 1-100 toolkit for AI coding agents: spec-driven pipeline, adversarial review bench, deterministic gates. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
highflame-ai
/
ai-factory
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .aif/ context .aif/ context .claude .claude .github/ workflows .github/ workflows agents agents bin bin compiled compiled examples examples hooks hooks packs packs references references roles roles skills skills tools tools .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md VERSION VERSION catalog.sh catalog.sh install.sh install.sh tips.md tips.md workspace-CLAUDE.md workspace-CLAUDE.md View all files Repository files navigation
0→1 with an AI coding agent is easy. 1→100 is where it gets hard — the code that survives review, scales past the demo, and doesn't quietly rot the codebase over months. That's where ai-factory comes in.
A complete spec-driven development practice for AI coding agents — skills, agents, deterministic hooks, and a role-based capability model — installed once and parameterized to your org through config, not forks.
Faster — across our own development at Highflame, the spec-driven pipeline plus parallel sprint orchestration routinely delivers 3–5x : features/PRs shipped, bugs resolved, and time-to-merge, vs. ad-hoc AI-assisted coding. (our internal experience — your mileage will vary)
Less slop — deterministic gates, an adversarial multi-agent review bench, and a verification-is-non-negotiable ethos: unreviewed AI output can't reach main .
Leaner context — heavy research and review run in subagents with their own context windows; the main session stays focused.
Standalone in Claude Code — 41 skills as slash commands ( /spec → /architect → /proceed → /ship ), symlinked live into every session. Clone, ./install.sh , go.
Under codeoid — the same methodology as a declarative pack on codeoid's multi-session runtime, with per-agent identity and cross-session memory: codeoid run --pack aif-sdlc .
The AIF pipeline — a spec-driven development lifecycle (spec → architect → validate → implement → reflect → review → ship) with parallel sprint orchestration, a multi-agent review bench, and the aif health CLI.
The org-workflow layer — day-to-day engineering skills (issue-to-PR orchestration, debugging, deprecation, git conventions, session handoffs, environment triage), review agents, and deterministic hooks. Everything org-specific resolves through .aif/config.yml — nothing is hardcoded, and every skill states what it does when a config key is absent.
The pack registry ( packs/ ) — the methodology as declarative codeoid packs ( schema: codeoid/pack@v1 ). A pack is data-only (a pack.yaml + capability roles + constitution); codeoid runs its governed phase pipeline. Contribute a pack once, and any team can select it. See packs/README.md — first pack: aif-sdlc .
git clone https://github.com/highflame-ai/ai-factory.git
cd ai-factory && ./install.sh
install.sh is idempotent and repair-capable. It symlinks:
It also writes the aif CLI shim to ~/bin ( doctor , check , agents render , compile , renumber ), scaffolds ~/.claude/aif/config.yml (delegation off by default), and finishes with aif doctor — a full environment health check that prints a copy-pasteable fix for every failure. ./install.sh --dry-run previews; --repair re-stamps after moving the clone.
The journey: install → fill in workspace-CLAUDE.md and the machine config ( ~/.claude/aif/config.yml ) → run claude → /init in each code repo to bootstrap its .aif/ structure → then fill that repo's .aif/config.yml (optionally starting from a skills/presets/ starter). Along the way, an adopting org customizes five things (details in examples/README.md ):
.aif/config.yml per repo — scaffolded by /init (from skills/templates/config-template.yml ) when you opt into cross-repo coordination, or copied in afterward from a preset — stack, sibling repos, and the org: / environments: / mcp: / local_stack: / regression: / tenancy: sections that drive the org-workflow skills.
workspace-CLAUDE.md — the always-loaded platform-context template (service map, auth contract, conventions).
Fill-in references — MCP cheatsheet, regression markers, known-warts registry, tenancy/JWT checklists.
Hook configuration — org.commit_prefix_regex (or AIF_COMMIT_PREFIX_REGEX ) to mirror your CI's commit gate; add your org's key prefixes to hooks/secret-scan.sh .
Your own skills and agents — skeletons and house rules in examples/ .
AIF pipeline (spec-driven lifecycle)
Skill
Description
/init
Bootstrap .aif/ structure in a repo
/onboard
Adopt the toolkit in an existing codebase — drafts the service map, starter experts, conventions, and config for review
/measure
Measure whether the toolkit is paying off — delivery/quality metrics from git, PRs, and .aif/ artifacts, with honest caveats
/spec
Write requirement specs from feature requests
/architect
Design architecture and break requirements into tasks
/validate
Validate any AIF phase output before advancing
/proceed
End-to-end pipeline: validate → architect → implement → reflect → review → PR → wrapup
/sprint
Parallel pipeline orchestrator — multiple /proceed runs across REQs ( --workflow engine)
/reflect
Post-implementation self-review before formal review
/review
Multi-agent code review (correctness, quality, architecture, tests, security)
/adversary
Adversarial review of any artifact — assumes it is wrong and tries to prove it
/canary
Canary deployment with smoke tests
/wrapup
Close out a feature — commit, merge, deploy, update artifacts
/bugfix
Streamlined bug fix workflow
/status / /manifest
Local / remote-derived view of in-flight AIF work
/analyze
Codebase health audit
/optimize
API cost & performance scanner
/template-drift
Detect drift between a project's .aif/templates/ and the toolkit's
Core workflow:
/spec → /validate → /architect → /validate → implement → /reflect → /review → merge → /wrapup
Org workflow
Skill
Description
/using-aif
Meta-skill routing table — injected at SessionStart
/from-issue
Issue → merge-ready PR orchestrator (classify, route, implement, ship, babysit CI)
/feature-prep
Does the spec registry need an entry first? Where do tests go?
/grill-feature
Adversarial design review before architectural work
/ship
Pre-merge orchestrator — parallel reviewer fan-out → go/no-go → post-deploy verify
/triage-dev
Investigate your shared dev environment in a configured observability order
/source-driven
Read the authoritative source (config, code, MCP) instead of guessing
/incremental-implementation
Thin vertical slices, one repo at a time
/debug
Local repro and root-cause; Phase 1 = fast deterministic feedback loop
/deprecate
Multi-stage removal across your configured repos
/git-workflow
Commit prefixes (configurable regex), atomic commits, branching, worktrees, PR flow
/handoff
Compact a multi-repo session into a handoff doc
/release-notes
Multi-audience release-notes draft (customer, engineering, executive) from a git tag or range
/dep-update
Vetted dependency updates — changelog risk review, isolated branch, test-proven, one PR per batch
/license-audit
SBOM + license compliance — resolve every component's license, judge against org.license_policy
/rotate-secrets
Proactive secret rotation — expiry gate over the secrets: inventory, overlap-pattern rotation, human gate before revoke
/audit-permissions
Review and prune risky standing Claude permission grants in .claude/settings*.json (pairs with aif doctor 's permissions-audit )
/doc-drift
Find and fix docs made stale by a change — diff-derived surfaces, stale vs missing classification
/threat-model
STRIDE threat model from a spec/RFC — assets, trust boundaries, refuted threats, mitigation punch list
/expert
Curate a domain expert — path-scoped context distilled from your docs, injected only when a change touches its domain
/add-detector
Template: scaffold a module in an extensible service, following your documen
[truncated]
26 subagents in agents/ , each with tier-based model selection rendered by aif agents render (config: ~/.claude/aif/config.yml ).
Review bench : correctness-reviewer, security-auditor, architecture-reviewer, quality-reviewer, code-quality-auditor, test-auditor, reflector, adversary
Scanners/explorers : api-cost-scanner, db-perf-scanner, latency-scanner, architecture-mapper, convention-auditor, feature-tracer, integration-explorer, delegate-pre-pass
Pipeline : task-implementer, pipeline-runner
Org workflow : security-reviewer (invariant-driven audit), cross-repo-impact, migration-analyzer, pr-shepherd, gemini-reviewer, cedar-policy-reviewer (if you use Cedar), local-stack-runner, kind-regression-runner
Experts — path-scoped domain context
/expert curates a domain expert : dense context distilled from your own docs and code, stored at .aif/experts/<name>.md , and injected into a task only when the change touches that domain (matched by the expert's applies_to globs). Context engineering — React context on a React change, DynamoDB context on a data-access change, neither on unrelated work. /review activates matching experts automatically. workspace-CLAUDE.md holds facts true everywhere, references/ are cross-cutting checklists, lessons are relevance-ranked incidents, experts are path-scoped domain knowledge. The applies_to globs are also the future compile key for path-scoped rules in other assistants (Copilot applyTo , Cursor globs) — that emission is the next phase.
Roles — one capability definition, every agent runtime
roles/*.yaml is the single source of truth for what each agent role may do. A reviewer declares {envelope: [read, grep, glob, bash], write: false} once, and aif compile emits it as:
Claude — validation of agents/*.md tools: frontmatter (the native layer; --check fails on any agent exceeding its role's envelope), plus compiled/claude/ per-role settings overlays for headless/CI sessions
Codex — compiled/codex/config.toml sandbox profiles per role
Copilot — compiled/copilot/ pa

[truncated]
