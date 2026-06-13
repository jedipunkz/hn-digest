---
source: "https://github.com/SawyerHood/omegacode"
hn_url: "https://news.ycombinator.com/item?id=48519317"
title: "Show HN: Omegacode – an agent agnostic version of Claude Workflows"
article_title: "GitHub - SawyerHood/omegacode: Code based orchestration for any coding agent. · GitHub"
author: "sawyerjhood"
captured_at: "2026-06-13T18:35:14Z"
capture_tool: "hn-digest"
hn_id: 48519317
score: 1
comments: 0
posted_at: "2026-06-13T17:18:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Omegacode – an agent agnostic version of Claude Workflows

- HN: [48519317](https://news.ycombinator.com/item?id=48519317)
- Source: [github.com](https://github.com/SawyerHood/omegacode)
- Score: 1
- Comments: 0
- Posted: 2026-06-13T17:18:53Z

## Translation

タイトル: Show HN: Omegacode – クロード ワークフローのエージェントに依存しないバージョン
記事のタイトル: GitHub - SawyerHood/omegacode: あらゆるコーディング エージェントのためのコード ベースのオーケストレーション。 · GitHub
説明: あらゆるコーディング エージェント向けのコード ベースのオーケストレーション。 GitHub でアカウントを作成して、SawyerHood/omegacode の開発に貢献してください。

記事本文:
GitHub - SawyerHood/omegacode: あらゆるコーディング エージェント向けのコード ベースのオーケストレーション。 · GitHub
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
ソーヤーフッド
/
オメガコード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブランチ タグ G

o ファイルへ コード 「その他のアクション」メニューを開く フォルダーとファイル
55 コミット 55 コミット .agents/ skill/ remotion-best-practices .agents/ skill/ remotion-best-practices .claude/ skill .claude/ skill .github/ workflows .github/ workflows ビルトイン ビルトイン サンプル サンプル omega-logos omega-logos plan プラン スクリプト scripts skill skill src src test test viewer viewer viewer .gitignore .gitignore DESIGN.md DESIGN.md LICENSE LICENSE PARITY.md PARITY.md README.md README.md package-lock.json package-lock.json package.json package.json review-掃引.workflow.js review-掃引.workflow.js スキルロック.json スキルロック.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code のワークフローのエージェントに依存しない実装。 omegacode は JavaScript ワークフローを実行します
小規模な決定論的 DSL を使用してコーディング エージェントのフリートを調整するファイル — Agent() /
Parallel() / Pipeline() / Phase() — ワーカーはプラグイン可能です。同じワークフローを使用できます。
Claude Code 、 Codex 、 OpenCode 、および pi を 1 回の実行で実行します。
npm install -g omegacode
オメガコードのインストールスキル
install-skill は、スキルをコピーすることで、エージェントにワークフローの作成方法と実行方法を教えます。
~/.claude/skills/ (クロード コード) および ~/.agents/skills/ (コーデックスおよびその他のエージェント)。パス
--claude または --agents を 1 つだけにインストールします。
Node 20+ と少なくとも 1 つのワーカーがインストールされている必要があります: codex (デフォルトのプロバイダー)、claude 、
opencode (≥ 1.16.2) および/または pi (≥ 0.79.1、 npm i -g @earendil-works/pi-coding-agent )。走る
omegacode ドクターがチェックします — ワーカーが最小バージョン未満のバイナリにフラグを立てます。
実行時に拒否します。
opencode/pi サンドボックスに関する注意: どちらの CLI も限定されたサンドボックスを強制できないため、omegacode
明示的なサンドボックス「danger-full-access」（呼び出しごとまたは経由）でのみそれらを受け入れます。
--砂

箱 ）。デフォルトの読み取り専用サンドボックスは、救済策の名前を指定するエラーで拒否されます。
意図的なフェイルクローズされた選択。モデル文字列はそのままバックエンドに渡されます (例:
Agent("…", { プロバイダー: "opencode"、モデル: "openrouter/anthropic/claude-sonnet-4.5"、サンドボックス: "danger-full-access" }) )。
スキルをインストールしたら、エージェントに次のように尋ねます。
omegacode を使用して、クロード コードとコーデックスの両方でこの PR を敵対的にレビューします
ワークフローを作成します。ファインダーは並行してファンアウトし、クロスプロバイダーの懐疑的なパスは、
それぞれの発見に反論し、残ったものを合成装置がマージし、それを実行して報告します。ランは
ジャーナリングと再開が可能で、omegacodeserve は動作中にすべてのエージェントのライブ ダッシュボードを開きます。
import const meta = { name : "adversarial-review" , description : "バグを見つけて反対尋問する" }
// FINDINGS と VERDICT はプレーンな JSON スキーマであり、ここでは省略されています
フェーズ (「検索」)
const 結果 = 並列待機 (
[ 「正確性」 、 「セキュリティ」 、 「パフォーマンス」 ] 。マップ ( ( レンズ ) => ( ) =>
エージェント (` ${ レンズ } レンズを通して差分を確認します。具体的な問題をリストします。` , { スキーマ : FINDINGS } ) ) 、
)
フェーズ (「検証」)
return await パイプライン (
所見。フィルター (ブール値) 。 flatMap (( f ) => f . issues )、
( issue ) => エージェント ( `反論してみます: ${ issue . desc } ` 、 { プロバイダー : "claude-code" 、モデル : "claude-fable-5" 、スキーマ : VERDICT } ) 、
)
プレーンな JavaScript、インポートなし - DSL が挿入されます。各 Agent() は実際の Codex を生成します、クロード
コード、OpenCode、または pi エージェント。実行が開始されたものを継承するプロバイダー/モデルを省略します。
( --provider --model 、デフォルトの codex )、またはクロスプロバイダーが必要な場合は呼び出しごとに固定します
多様性。プロバイダーとモデルは、すべてのサイトで両方またはどちらでもない (呼び出しごと、メタデフォルト、
CLI フラグ): 単独のプロバイダー: またはモデル: は拒否されるため、目的のモデルは

1 つのプロバイダーでできること
他のプロバイダーの通話に黙って乗らないでください。
omegacode を実行 <file.workflow.js | name> # ワークフローを実行します (ライブ ビューアを自動起動します)
omegacode は、すべての実行にわたって # 読み取り専用ダッシュボードを提供します
omegacode run <name> --resume <runId> # 再開 — 変更されたサフィックスのみが再実行されます
omegacode Doctor # codex/claude/opencode/pi の可用性とバージョンを確認します
omegacode guide # 完全なオーサリングガイドを印刷する
run は保存されたワークフロー名も受け入れます。 6 つの組み込み機能がパッケージに同梱されています。
deep-research 、 code-review 、および 2 つを組み合わせた 4 つのマルチプロバイダー ワークフロー
モデルの非相関エラーが機能するようにする — マルチプロバイダーレビュー (両方とも同じレビューを行う)
独立して分岐し、その後合成で両方がマージされます）、ベイクオフ（両方とも
分離されたワークツリーで同じタスクを実行し、盲目的なクロスプロバイダーの審査員が勝者を選びます)、
プロバイダー討論（提案→攻撃→Nラウンドの反論、その後裁判官の裁定）、および
セカンドオピニオン（両方とも安いと答える；合意は統合され、意見の相違はエスカレートする）
深い努力と裁定を行ってください）。 omegacode run deep-research --args '"your question"' 、または omegacode ワークフローを試してリストを表示します。詳細については、オメガコード ガイドを参照してください。
完全なオーサリング リファレンス。
あらゆるコーディング エージェント向けのコード ベースのオーケストレーション。
読み込み中にエラーが発生しました。このページをリロードしてください。
4
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Code based orchestration for any coding agent. Contribute to SawyerHood/omegacode development by creating an account on GitHub.

GitHub - SawyerHood/omegacode: Code based orchestration for any coding agent. · GitHub
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
SawyerHood
/
omegacode
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
55 Commits 55 Commits .agents/ skills/ remotion-best-practices .agents/ skills/ remotion-best-practices .claude/ skills .claude/ skills .github/ workflows .github/ workflows builtins builtins examples examples omega-logos omega-logos plans plans scripts scripts skill skill src src test test viewer viewer .gitignore .gitignore DESIGN.md DESIGN.md LICENSE LICENSE PARITY.md PARITY.md README.md README.md package-lock.json package-lock.json package.json package.json review-sweep.workflow.js review-sweep.workflow.js skills-lock.json skills-lock.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts View all files Repository files navigation
An agent-agnostic implementation of Claude Code's Workflows . omegacode runs JavaScript workflow
files that orchestrate fleets of coding agents with a small deterministic DSL — agent() /
parallel() / pipeline() / phase() — and the workers are pluggable: the same workflow can
drive Claude Code , Codex , OpenCode , and pi in a single run.
npm install -g omegacode
omegacode install-skill
install-skill teaches your agents how to author and run workflows by copying the skill into
~/.claude/skills/ (Claude Code) and ~/.agents/skills/ (Codex and other agents). Pass
--claude or --agents to install to just one.
You'll need Node 20+ and at least one worker installed: codex (the default provider), claude ,
opencode (≥ 1.16.2), and/or pi (≥ 0.79.1, npm i -g @earendil-works/pi-coding-agent ). Run
omegacode doctor to check — it flags binaries below the minimum versions, which the workers
refuse at runtime.
Note on opencode/pi sandboxing: neither CLI can enforce a confined sandbox, so omegacode
accepts them only with an explicit sandbox: "danger-full-access" (per call or via
--sandbox ). The default read-only sandbox is rejected with an error naming the remedy —
a deliberate fail-closed choice. Model strings pass through verbatim to the backend (e.g.
agent("…", { provider: "opencode", model: "openrouter/anthropic/claude-sonnet-4.5", sandbox: "danger-full-access" }) ).
With the skill installed, just ask your agent:
use omegacode to adversarially review this PR with both claude code and codex
It will author a workflow — finders fan out in parallel, a cross-provider skeptic pass tries to
refute each finding, a synthesizer merges what survives — then run it and report back. Runs are
journaled and resumable, and omegacode serve opens a live dashboard of every agent as it works.
export const meta = { name : "adversarial-review" , description : "find bugs, cross-examine them" }
// FINDINGS and VERDICT are plain JSON Schemas, elided here
phase ( "Find" )
const findings = await parallel (
[ "correctness" , "security" , "performance" ] . map ( ( lens ) => ( ) =>
agent ( `Review the diff through the ${ lens } lens. List concrete issues.` , { schema : FINDINGS } ) ) ,
)
phase ( "Verify" )
return await pipeline (
findings . filter ( Boolean ) . flatMap ( ( f ) => f . issues ) ,
( issue ) => agent ( `Try to refute: ${ issue . desc } ` , { provider : "claude-code" , model : "claude-fable-5" , schema : VERDICT } ) ,
)
Plain JavaScript, no imports — the DSL is injected. Each agent() spawns a real Codex, Claude
Code, OpenCode, or pi agent; omit provider / model to inherit whatever the run was started with
( --provider --model , default codex ), or pin them per call when you want cross-provider
diversity. Provider and model are both-or-neither at every site (per-call, meta defaults,
CLI flags): a lone provider: or model: is rejected, so a model meant for one provider can
never silently ride another provider's call.
omegacode run <file.workflow.js | name> # run a workflow (auto-starts the live viewer)
omegacode serve # read-only dashboard over all runs
omegacode run <name> --resume <runId> # resume — only the changed suffix re-runs
omegacode doctor # check codex/claude/opencode/pi availability + versions
omegacode guide # print the full authoring guide
run also accepts saved workflow names. Six built-ins ship with the package:
deep-research , code-review , and four multi-provider workflows that put the two
models' decorrelated errors to work — multi-provider-review (both review the same
branch independently, then a synthesis merges both), bake-off (both implement the
same task in isolated worktrees, blind cross-provider judges pick a winner),
provider-debate (propose → attack → rebut for N rounds, then a judge rules), and
second-opinion (both answer cheap; agreement returns merged, disagreement escalates
to deep effort and adjudicates). Try omegacode run deep-research --args '"your question"' , or omegacode workflows to list them. See omegacode guide for the
complete authoring reference.
Code based orchestration for any coding agent.
There was an error while loading. Please reload this page .
4
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
