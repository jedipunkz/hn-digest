---
source: "https://github.com/lucasgodt/open-agent-ready-typescript"
hn_url: "https://news.ycombinator.com/item?id=48920268"
title: "Show HN: A TypeScript repo where AI agents can't break the architecture"
article_title: "GitHub - lucasgodt/open-agent-ready-typescript · GitHub"
author: "lucasgodt"
captured_at: "2026-07-15T13:12:28Z"
capture_tool: "hn-digest"
hn_id: 48920268
score: 1
comments: 1
posted_at: "2026-07-15T13:06:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A TypeScript repo where AI agents can't break the architecture

- HN: [48920268](https://news.ycombinator.com/item?id=48920268)
- Source: [github.com](https://github.com/lucasgodt/open-agent-ready-typescript)
- Score: 1
- Comments: 1
- Posted: 2026-07-15T13:06:07Z

## Translation

タイトル: Show HN: AI エージェントがアーキテクチャを破壊できない TypeScript リポジトリ
記事のタイトル: GitHub - lucasgodt/open-agent-ready-typescript · GitHub
説明: GitHub でアカウントを作成して、lucasgodt/open-agent-ready-typescript の開発に貢献します。

記事本文:
GitHub - lucasgodt/open-agent-ready-typescript · GitHub
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
ルーカスゴット
/
オープンエージェントレディタイプスクリプト
パブリックテンプレート
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
lucasgodt/open-agent-read

y タイプスクリプト
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット .claude .claude .github/ workflows .github/ workflows docs/ adr docs/ adr スクリプト スクリプト 仕様 仕様 src src テスト テスト .dependency-cruiser.cjs .dependency-cruiser.cjs .gitignore .gitignore 2026-07-13-143050-implemente-o-exerccio-de-benchmark-descrito-no-f.txt 2026-07-13-143050-implemente-o-exerccio-de-benchmark-descrito-no-f.txt 2026-07-13-144840-implemente-o-exerccio-de-benchmark-descrito-no-f.txt 2026-07-13-144840-implemente-o-exerccio-de-benchmark-descrito-no-f.txt AGENTS.md AGENTS.md CLAUDE.md CLAUDE.mdライセンス ライセンス README.md README.md eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json stryker.config.json stryker.config.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントが適切なコードを生成できるように設計された参照 TypeScript コードベース
そこには、ドキュメントとしてではなく、ガードレールとしてのアーキテクチャが含まれています。
ほとんどの AI 支援コードベースは同じように劣化します。生産性の向上と 3 つの点です。
数か月後、3 つのファイルに重複したロジックが存在し、テストではアサートが行われます。
何もなく、エージェントが作成したコードを変更するよう求められたときにそれを破壊します。
それはモデルの問題ではなく、ワークフローの問題です。散文の慣例ではそうではありません
エージェント（または疲れた人間）との接触を生き延びます。機械であるルール
する。
このリポジトリは小さいながらも実際の請求 API (ドラフト → 送信 → 支払い、金額として) です。
整数セント、適切な状態マシン) は、1 つの命題を実証するために構築されました。
重要なルールはすべて、エージェントが実行できないツールによって強制する必要があります
と議論する。
依存関係ルールはビルドの失敗であり、図ではありません。
dependency-cruiser は、不正な検証、コミットフック、および CI に失敗します。
輸入する。ソース/ドメイン

n は何もインポートしません。ノードの組み込みもインポートしません。エージェント
アーキテクチャを覚える必要はありません。違反するとマージできません。
突然変異テストは、テストを行わないテストを検出します。
報道によれば、列が走ったとのこと。ストライカーはこう語る
テストは行が壊れているかどうかを認識します。 CI は 75% の変異スコアを下回ります。
アプリケーション層のスコアは現在 100% です。
テストと仕様はエージェントから保護されます。
クロード コードの PreToolUse フックは、エージェントによる testing/** および
人間が明示的にオーバーライドを許可しない限り、specs/** となります。 「テストをしなさい
pass" は、テストを削除しても満足できなくなります (ADR 0004)。
コミットは完全な検証スイートでゲートされます。
フックは git コミットごとに typecheck + lint + deps + test を実行し、
失敗はエージェントにフィードバックされます。これは単なる壁ではなく、フィードバック ループです。
仕様駆動開発は構築ゲートであり、提案ではありません。
ユースケースごとに 1 つのマークダウン仕様 (Given/When/Then 受け入れ基準を含む)。
テストは基準ごとに 1 つの it() をミラーリングし、npm 実行仕様 (内部)
verify 、コミットフックおよび CI) は仕様のないユースケースでは失敗します。
孤立した仕様、およびミラーリングされたテストのない基準。黄金の
パスは npm run new -- <ユースケース> : 仕様 + ミラーリングされたテスト スタブ + ユースケース
一致する名前でスタブがスキャフォールドされ、ワークフローが印刷されます。
ルートにある AGENTS.md (オープンスタンダード)、
300 行未満、命令型、正確なコマンド - ネストされた AGENTS.md
src/domain 、 src/application 、および src/infrastruct 内のファイルを運ぶ
各層のローカル不変量。
CLAUDE.md は、AGENTS.md を指す 1 行です (ツール間で移植可能)。
docs/adr/ は決定の背後にある理由を記録するため、エージェントは停止します
彼らを再訴訟する。
どこにでも決定的な継ぎ目: 時間とアイデンティティはポートです ( Clock 、
IdGenerator ); new Date() とrandomUUID() は
構成ルート。テストはスリープすることも、不安定になることもありません。
T

AGENTS.md の最後には自己完結型タスクが含まれています
( cancel-invoice ) コーディング エージェントが動作できるかどうかを検証するように設計されています。
このリポジトリは正しく: 仕様 → テスト → ドメイン → ユースケース → HTTP、緑色を確認、
他には何も触れませんでした。エージェントにそれを指示して、差分を評価してください。それが
このリポジトリの要点 — これは単なる例ではなく、テストベンチです。
このプロジェクトから次のプロジェクトを始めましょう
このリポジトリはテンプレートとしても機能し、移植を行う Claude Code スキルを同梱しています。
あなた: ガードレール システム (構成、フック、CI、レイヤー構造) を維持します。
請求書の例を削除し、その中で新しいドメイン仕様を最初に拡張します。
3 つのプロジェクト形状をカバーしており、それぞれに独自のプレイブックがあります。
参照/ :
バックエンド API/サービス — このリポジトリ独自の構造 ( backend.md )
フロントエンド アプリ — 純粋なアダプターとしての UI を使用した React/Vite
React-free ドメイン。突然変異テストはドメイン/アプリケーションを対象とし、決して行わない
コンポーネント (frontend.md)
Monorepo — レイヤーが物理パッケージになる pnpm ワークスペース。
空の依存関係オブジェクトを持つパッケージ/ドメインが最もうるさいです
リポジトリ内のガードレール ( monorepo.md )
BASE=https://raw.githubusercontent.com/lucasgodt/open-agent-ready-typescript/main/.claude/skills/scaffold-agent-ready
mkdir -p ~ /.claude/skills/scaffold-agent-ready/references
カール -sL $BASE /SKILL.md -o ~ /.claude/skills/scaffold-agent-ready/SKILL.md
バックエンド フロントエンド モノリポジトリの f の場合。する
カール -sL $BASE /references/ $f .md -o ~ /.claude/skills/scaffold-agent-ready/references/ $f .md
完了しました
次に、空のディレクトリから、クロード コードで次のようにします。
新しいエージェント対応のスキャフォールディング [バックエンド | バックエンド]フロントエンド | [あなたのドメイン] の monorepo] プロジェクト
手動で行う方が良いですか?レシピはスキルが従うものと同じです:
クローンを作成し、すべての設定/フック/ワークフロー ファイルとレイヤー スケルトンを保持し、削除します

src ドメインの内容 + 仕様 + 例のテストを実行し、ドメインを再構築します
AGENTS.md の仕様優先ワークフローに従います。 GitHubの「このテンプレートを使用する」
ボタンは機械コピーでも機能します。
npmインストール
npm run verify # typecheck + lint + 依存関係ルール + テスト
npm run dev # API :3000
npm run ミューテーション # Stryker (遅い)
curl -s -X POST :3000/invoices -H ' content-type: application/json ' \
-d ' {"顧客名":"ACME Ltda","通貨":"BRL"} '
正直な限界
テスト保護フックはクロード コード固有です。他のエージェントには
同等のゲート。 CI での突然変異テストは、ツールに依存しないバックストップです。
JSON ファイルの永続化は意図的に退屈であり、複数のファイル間での同時安全性がありません。
プロセス (ADR 0003)。実際のデータベースでのスワップは 1 つのアダプタ + 1 行です
構成ルートでは、それがまさにデモンストレーションです。
生き残った変異は、エラー メッセージ文字列の変異体です。の
契約はエラーコードであり、散文ではなく、あらゆる場所で主張されます。
読み込み中にエラーが発生しました。このページをリロードしてください。
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

Contribute to lucasgodt/open-agent-ready-typescript development by creating an account on GitHub.

GitHub - lucasgodt/open-agent-ready-typescript · GitHub
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
lucasgodt
/
open-agent-ready-typescript
Public template
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
lucasgodt/open-agent-ready-typescript
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits .claude .claude .github/ workflows .github/ workflows docs/ adr docs/ adr scripts scripts specs specs src src tests tests .dependency-cruiser.cjs .dependency-cruiser.cjs .gitignore .gitignore 2026-07-13-143050-implemente-o-exerccio-de-benchmark-descrito-no-f.txt 2026-07-13-143050-implemente-o-exerccio-de-benchmark-descrito-no-f.txt 2026-07-13-144840-implemente-o-exerccio-de-benchmark-descrito-no-f.txt 2026-07-13-144840-implemente-o-exerccio-de-benchmark-descrito-no-f.txt AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json stryker.config.json stryker.config.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
A reference TypeScript codebase engineered so AI agents produce good code
in it — architecture as guardrails, not as documentation.
Most AI-assisted codebases degrade the same way: productivity jumps, and three
months later there's duplicated logic in three files, tests that assert
nothing, and an agent that breaks the code it wrote when asked to change it.
That's not a model problem — it's a workflow problem. Prose conventions don't
survive contact with agents (or with tired humans). Rules that are machines
do.
This repo is a small but real invoicing API (draft → send → pay, money as
integer cents, a proper state machine) built to demonstrate one thesis:
Every rule that matters must be enforced by a tool the agent cannot
argue with.
The dependency rule is a build failure, not a diagram.
dependency-cruiser fails verify , the commit hook and CI on any illegal
import. src/domain imports nothing — not even node builtins. The agent
doesn't need to remember the architecture; violating it is impossible to merge.
Mutation testing catches tests that don't test.
Coverage says a line ran. Stryker says the
tests would notice if the line broke . CI breaks below 75% mutation score;
the application layer currently scores 100%.
Tests and specs are protected from the agent.
A Claude Code PreToolUse hook blocks agent edits to tests/** and
specs/** unless a human explicitly grants an override. "Make the tests
pass" can no longer be satisfied by deleting the tests (ADR 0004).
Commits are gated on the full verify suite.
A hook runs typecheck + lint + deps + test on every git commit and
feeds failures back to the agent — feedback loop, not just a wall.
Spec-driven development is a build gate, not a suggestion.
One markdown spec per use case with Given/When/Then acceptance criteria;
tests mirror them one it() per criterion — and npm run specs (inside
verify , the commit hook and CI) fails on any use case without a spec,
any orphan spec, and any criterion without its mirrored test. The golden
path is npm run new -- <use-case> : spec + mirrored test stub + use-case
stub scaffolded with matching names, workflow printed.
AGENTS.md at the root (the open standard ),
under 300 lines, imperative, with exact commands — plus nested AGENTS.md
files in src/domain , src/application and src/infrastructure carrying
each layer's local invariants.
CLAUDE.md is one line pointing at AGENTS.md (portable across tools).
docs/adr/ records the why behind decisions, so agents stop
relitigating them.
Deterministic seams everywhere: time and identity are ports ( Clock ,
IdGenerator ); new Date() and randomUUID() exist only in the
composition root. Tests never sleep, never flake.
The end of AGENTS.md contains a self-contained task
( cancel-invoice ) designed to verify whether any coding agent can work in
this repo correctly: spec → tests → domain → use case → HTTP, verify green,
nothing else touched. Point your agent at it and grade the diff. That's the
whole point of this repository — it's not just an example, it's a test bench.
Start your next project from this one
This repo doubles as a template, and ships a Claude Code skill that does the transplant for
you: it keeps the guardrail system (configs, hooks, CI, layer structure),
removes the invoicing example, and grows your new domain spec-first inside it.
It covers three project shapes, each with its own playbook in
references/ :
Backend API/service — this repo's own structure ( backend.md )
Frontend app — React/Vite with the UI as an adapter around a pure,
React-free domain; mutation testing targets domain/application, never
components ( frontend.md )
Monorepo — pnpm workspaces where layers become physical packages;
packages/domain with an empty dependencies object is the loudest
guardrail in the repo ( monorepo.md )
BASE=https://raw.githubusercontent.com/lucasgodt/open-agent-ready-typescript/main/.claude/skills/scaffold-agent-ready
mkdir -p ~ /.claude/skills/scaffold-agent-ready/references
curl -sL $BASE /SKILL.md -o ~ /.claude/skills/scaffold-agent-ready/SKILL.md
for f in backend frontend monorepo ; do
curl -sL $BASE /references/ $f .md -o ~ /.claude/skills/scaffold-agent-ready/references/ $f .md
done
Then from any empty directory, in Claude Code:
Scaffold a new agent-ready [backend | frontend | monorepo] project for [your domain]
Prefer doing it by hand? The recipe is the same one the skill follows:
clone, keep every config/hook/workflow file plus the layer skeleton, delete
src domain contents + specs + tests of the example, and rebuild your domain
following the spec-first workflow in AGENTS.md. GitHub's "Use this template"
button also works for the mechanical copy.
npm install
npm run verify # typecheck + lint + dependency rules + tests
npm run dev # API on :3000
npm run mutation # Stryker (slower)
curl -s -X POST :3000/invoices -H ' content-type: application/json ' \
-d ' {"customerName":"ACME Ltda","currency":"BRL"} '
Honest limitations
The test-protection hook is Claude Code–specific; other agents need an
equivalent gate. Mutation testing in CI is the tool-agnostic backstop.
JSON-file persistence is deliberately boring and not concurrent-safe across
processes (ADR 0003). Swapping in a real database is one adapter + one line
in the composition root — which is precisely the demonstration.
Mutation survivors that remain are error-message string mutants; the
contract is the error code , asserted everywhere, not the prose.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
