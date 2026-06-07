---
source: "https://github.com/mlguyYT/AgentCrew"
hn_url: "https://news.ycombinator.com/item?id=48436561"
title: "Show HN: AgentCrew – a Markdown-first operating system for AI coding agents"
article_title: "GitHub - mlguyYT/AgentCrew: AgentCrew is a Markdown workflow for coordinating AI agent teams with roles, playbooks, skills, handoffs, and human approval gates. · GitHub"
author: "mhjafari92"
captured_at: "2026-06-07T18:33:42Z"
capture_tool: "hn-digest"
hn_id: 48436561
score: 3
comments: 0
posted_at: "2026-06-07T16:50:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AgentCrew – a Markdown-first operating system for AI coding agents

- HN: [48436561](https://news.ycombinator.com/item?id=48436561)
- Source: [github.com](https://github.com/mlguyYT/AgentCrew)
- Score: 3
- Comments: 0
- Posted: 2026-06-07T16:50:46Z

## Translation

タイトル: Show HN: AgentCrew – AI コーディング エージェント用の Markdown ファースト オペレーティング システム
記事タイトル: GitHub - mlguyYT/AgentCrew: AgentCrew は、ロール、プレイブック、スキル、ハンドオフ、および人間の承認ゲートを使用して AI エージェント チームを調整するためのマークダウン ワークフローです。 · GitHub
説明: AgentCrew は、役割、プレイブック、スキル、ハンドオフ、および人間の承認ゲートを使用して AI エージェント チームを調整するための Markdown ワークフローです。 - GitHub - mlguyYT/AgentCrew: AgentCrew は、ロール、プレイブック、スキル、ハンドオフ、および人間の承認ゲートを使用して AI エージェント チームを調整するためのマークダウン ワークフローです。

記事本文:
GitHub - mlguyYT/AgentCrew: AgentCrew は、役割、プレイブック、スキル、ハンドオフ、および人間の承認ゲートを使用して AI エージェント チームを調整するためのマークダウン ワークフローです。 · GitHub
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
mlguyYT
/
エージェントクルー
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
48 コミット 48 コミット .claude .claude .codex .codex .cursor/ rules .cursor/ rules .github .github エージェントチーム エージェントチーム bin bin docs docs 例 例 ランタイム ランタイム .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md バージョン バージョン すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェントを規律あるチームに変えます。
AgentCrew は、会話ファースト、マークダウンファーストのエージェント コーディング方法論です。既存のコーディング エージェント (Claude Code、Codex、Cursor、OpenClaw、Hermes、またはその他の互換性のあるアシスタント) は、これをガイダンスとして読み込み、役割、ルーティング、ハンドオフ、品質ゲート、および人間の承認を伴うソフトウェア タスクを実行するために使用します。
働き方は変わりません。あなたはエージェントと話し続けます。 AgentCrew は、そのチャットの下にチーム プロセスを提供します。
方法論は主に Markdown とシェルです。オプションのエンジンは、自動マルチエージェント実行が便利な場合に、ホスト エージェントがバックグラウンドで呼び出すことができる実行可能レイヤーです。
AgentCrew はデーモンではありません。
CI に代わるものではありません。
それは魔法のような自律型エンジニアではありません。
コーディングエージェントにとって実用的なオペレーティングシステムです。
通常、単一のコーディング エージェント セッションは、スコープ設定、実装、テスト、レビュー、文書化、そして場合によっては作成したばかりの作業を承認できるかのように、あらゆるジョブを実行する 1 つのコンテキストです。
それは優れたソフトウェア チームの働き方ではありません。
バグを作成したのと同じモデルがバグの唯一のレビュー者であってはなりません。 API を推測した同じチャットが、その API が正しいかどうかの最終的な権威となるべきではありません。危険な認証変更は扱わないでください

タイプミスの修正が好きです。リファクタリングは、コードベースの無関係な部分にまたがって暗黙的に拡張すべきではありません。
ほとんどのチームは、 CLAUDE.md 、 AGENTS.md 、またはエディター ルールで、より大きなプロンプトや「注意してください」という注記を表示することで、この問題を解決しようとします。これは少しは役に立ちますが、エージェントは依然としてトンネルビジョンを行ったり、テストをスキップしたり、コンテキストを失ったり、無関係な編集を行ったり、信頼するのが難しい巨大な差分を人間に残したりする可能性があります。
AgentCrew は、コーディング エージェントがコードに触れる前にプロセスを提供するために存在します。
私たちは、コーディング エージェントで同じ問題に繰り返し遭遇した後、AgentCrew を構築しました。エージェントはコードを速く書くことができますが、自然に規律ある開発サイクルに従うわけではありません。
あらゆる重大な変更には実装以上のものが必要です。エージェントはタスクを理解し、適切なコンテキストを選択し、適切なスキルを使用し、テストを作成または更新し、それらのテストが実際に機能することを確認し、変更を文書化し、人間がレビューできるものを準備する必要があります。
この構造がなければ、同じ指示を何度も繰り返すことになります。
故障モードは理論上のものではありません。私たちは、エージェント支援のリファクタリングがコードベース全体に厄介な変更を加え、重要なコンテキストを見逃し、人間がクリーンアップに数日かかることを見てきました。あるケースでは、管理されているはずのリファクタリングが最終的にかなりの部分を破壊し、手動で修正するのに約 1 週間かかりました。問題は、モデルが役に立たないことではありませんでした。問題は、実際のプロセス、専門的な役割、スキルの境界線、および本格的なレビューゲートがないことでした。
私たちは、より大きなプロンプト、プロジェクト ルール、内部 Codex ワークフロー、プロンプト スタック、コマンド ベースのアプローチを試しました。助けてくれた人もいました。しかし、特にタスクの計画、テスト、レビュー、セキュリティへの注意、または人間の決定がいつ必要かをエージェント自体に認識させたい場合には、依然として手動による操作が多すぎます。
AgentCrew はシステムです

私たちが望んでいたのは、魔法の自動操縦ではなく、開発者の判断に代わるものでもありませんでした。これは、エージェント コーディング (役割、スキル、ルーティング、ハンドオフ、品質ゲート、メモリ、承認境界) のための規律ある方法論であるため、コーディング エージェントは信頼されやすく、レビューが容易になり、実際の開発作業でさらに役立ちます。
通常のコーディング エージェントに結果を求めます。
Google OAuth ログインをサインアップ ページに追加します。
AgentCrew がロードされていても、エージェントはすぐにファイルの編集に入るわけではありません。まず、次の方法論に従います。
リクエストを分類する
純粋な bash 分類子は、タスクの種類、リスク レベル、レーン、可能性の高い役割、関連スキル、品質ゲートを識別できます。
右の車線を選択してください
単純でリスクの低い作業には高速レーンを使用できます。リスクが高く、複数のステップが必要な作業、セキュリティが重要な作業、またはアーキテクチャを変更する作業には、より完全なワークフローが使用されます。
適切な役割から始める
エージェントは、実装前に範囲と受け入れ基準を明確にするために、アドバイザーまたはプロダクト マネージャーとして開始する場合があります。
方向性の確認と費用の見積り
意味のあるモデル トークンを消費する作業の場合、AgentCrew は実行を開始する前に推定コストを明らかにできます。
仕事を計画して発送する
ユーザーが指示を確認すると、AgentCrew は作業を準備し、ホスト エージェントまたはオプションのエンジンを通じて必要なエージェントを実行できるようになります。ユーザーは AgentCrew コマンドを手動で実行する必要はありません。
開発者として実装する
開発者の役割は、最小限の関連する変更に焦点を当てます。無関係な編集、隠れたエラー、安全でないコマンド、および時期尚早なマージ決定を回避する必要があります。
テスターとしてチェックする
テスターの役割は検証に重点を置きます。適切なテストを実行または提案し、何が成功したか、何が失敗したか、またはチェックできなかったかを報告する必要があります。
レビュアーまたはセキュリティレビュアーとしてレビューする
危険なコード パス、認証、認可、支払い、移行、秘密、

インフラストラクチャの変更には、より厳格なレビュー ガイダンスが適用されます。
ハンドオフを生成する
最終的な出力は単に「完了」しただけではありません。何が変更されたのか、何がテストされたのか、どのようなリスクが残っているのか、そして何が人間の承認を必要とするのかを含める必要があります。
人間の門で立ち止まってください
AgentCrew は、エージェントを最終承認者として扱いません。どこに着陸するかを決めるのは依然として人間です。
通常のエージェント チャットを続けます。 AgentCrew は、そのチャットにルーティング システム、役割の指示、引き継ぎ形式、状態アーティファクト、および承認ゲートを提供するため、作業は 1 つの長い非構造化セッションではなく、チーム プロセスに従います。
AgentCrew のクローンをプロジェクトの外部のどこかに作成します。
git clone https://github.com/mlguyYT/AgentCrew ~ /AgentCrew
AgentCrew をホスト エージェントに登録します。
~ /AgentCrew/bin/agentcrew インストール
設定を確認します。
~ /AgentCrew/bin/agentcrew ドクター
次に、通常のコーディング エージェントを開き、通常どおり作業を開始します。
AgentCrew は、次のような小さなローダー ブロックをホスト エージェント命令ファイルにインストールします。
~/.claude/CLAUDE.md
~/.codex/AGENTS.md
~/.hermes/SOUL.md
ローダーは、ホスト エージェントに AgentCrew 方法論を指示します。プロジェクトには新しいランタイム依存関係は必要ありません。
エージェントチーム/
§─ エージェント/アドバイザー、PM、開発者などの役割定義
│ テスター、レビューアー、セキュリティレビューアー、UXレビューアー、
│ ドキュメントエージェントなど
│
§─ ツール/CLI ヘルパー (純粋な bash タスク分類子を含む)
│
§─ プレイブック/ タスクのルーティング方法、ゲートの処理方法、PR の準備方法、
│ 手戻りを管理し、有用なメモリを節約
│
§─ レシピ/ 機能、バグ修正、リファクタリング、
│ 移行、レビュー、セキュリティに配慮した作業など
│
§─ プロトコル/ハンドオフ形式、状態アーティファクト、承認境界、
│とトークン規律ルール
│
§─ context/ ルート インデックスと車線固有のコンテキスト プロファイル
│
§─ チェックリスト/プレミー

準備完了、レビューチェック、メモリのリフレッシュ、
│その他の品質管理
│
§─ テンプレート/ タスクの概要、作業計画、コンパクトな引き継ぎ、レビュー
│ レポート、PR説明、セッション概要
│
━─ スキル／積める技術的・専門的スキル
関連する場合のみ
bin/agentcrew インストール、ドクター、ステータス、分類、コンテキスト、開始、
概要、計画、準備完了、事前パック、チェックポイント、セッションの保存、
復元セッション、検出プロジェクト、プリセット
ホスト エージェント オーケストレーション用のエンジン/オプションの実行可能レイヤー、
コストのプレビューとマルチエージェントの実行
方法論はマークダウンです。
分類子とセットアップ ヘルパーはシェル スクリプトです。
プロジェクト内にはデーモンも必要なランタイムもありません。
すべてのタスクに重いワークフローが必要なわけではありません。
AgentCrew は作業をレーンに分割するため、エージェントが単純なタスクを過剰に処理したり、危険なタスクを過少処理したりすることはありません。
小規模でリスクの低い作業に使用されます。
目標は、基本的な規律を備えたスピードです。
より強力な制御が必要な作業に使用されます。
認証または認可の変更。
マルチファイルまたはマルチフェーズの実装。
目標は、より安全な実行、より明確な引き継ぎ、より良いレビューです。
AgentCrew は、1 つの一般的なチャットにすべてを依頼するのではなく、コーディング エージェントのロール固有の指示を提供します。
アドバイザー — 実装前に説明、指導し、推論を支援します。
プロダクト マネージャー — 範囲、ユーザーの価値、受け入れ基準、トレードオフを明確にします。
開発者 — 無関係な編集を最小限に抑えて、焦点を絞った変更を実装します。
テスター — 動作を検証し、テストされた内容を報告します。
レビュー担当者 — 実装の正確性、保守性、リスクをレビューします。
Security Reviewer — 認証、権限、シークレット、安全でない操作などの機密パスをチェックします。
UX レビュー担当者 — ユーザー対応の動作とインタラクションの品質をレビューします。
文書作成エージェント —

ドキュメント、例、および使用上の注意を更新またはレビューします。
リリース マネージャー — リリース ノートの準備、準備状況のチェック、リリースのハンドオフを支援します。
重要なのは、別々の人間が存在するふりをしないことです。重要なのは、エージェントが作業のさまざまなフェーズに対して異なる制約を使用できるようにすることです。
AgentCrew は、プロジェクトローカルの作業状態を .agent-state/ に維持できます。
.エージェント状態/
§─ 現在のタスク.md
§─ 作業計画.md
§─ 人間の決断.md
§─ handoff.md
§─ テストレポート.md
§─ レビューレポート.md
§─ pr-pack.md
└─ セッションチェックポイント/
これにより、エージェントに現在のプロジェクトの制御されたメモリ領域が与えられます。
目標はすべてを保存することではありません。目標は、重要な部分を保存することです。
どのような決定がすでに行われているか。
人間の承認がまだ必要なもの。
エージェントの優れたメモリは、アーカイブではなく圧縮されている必要があります。
AgentCrew は、方法論がスタックに適合するようにプロジェクト プリセットを検出または構成できます。
プロジェクト プリセットでは、より厳格な期待値、関連スキル、または追加のレビュー ゲートを追加できます。
プロジェクトのローカル構成により、プロセスが強化されるはずです。人間の承認を削除したり、必要なゲートをスキップしたり、中核となる安全規則を弱めたりしてはなりません。
AgentCrew は、次のような単純な原則に基づいて構築されています。
エージェントは仕事の準備を手伝うことができますが、仕事を着地させる決定権は依然として人間にあります。
エージェントは自分の作業をマージしてはなりません。
エージェントはブランチ プロをバイパスしてはなりません

[切り捨てられた]

## Original Extract

AgentCrew is a Markdown workflow for coordinating AI agent teams with roles, playbooks, skills, handoffs, and human approval gates. - GitHub - mlguyYT/AgentCrew: AgentCrew is a Markdown workflow for coordinating AI agent teams with roles, playbooks, skills, handoffs, and human approval gates.

GitHub - mlguyYT/AgentCrew: AgentCrew is a Markdown workflow for coordinating AI agent teams with roles, playbooks, skills, handoffs, and human approval gates. · GitHub
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
mlguyYT
/
AgentCrew
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
48 Commits 48 Commits .claude .claude .codex .codex .cursor/ rules .cursor/ rules .github .github agent-team agent-team bin bin docs docs examples examples runtime runtime .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md VERSION VERSION View all files Repository files navigation
Turn your coding agent into a disciplined team.
AgentCrew is a conversation-first, Markdown-first methodology for agentic coding. Your existing coding agent — Claude Code, Codex, Cursor, OpenClaw, Hermes, or another compatible assistant — loads it as guidance and uses it to work through software tasks with roles, routing, handoffs, quality gates, and human approval.
You do not change how you work. You keep talking to your agent. AgentCrew gives that chat a team process underneath.
The methodology is mostly Markdown and shell. The optional engine is the executable layer host agents can call behind the scenes when automatic multi-agent execution is useful.
AgentCrew is not a daemon.
It is not a replacement for CI.
It is not a magical autonomous engineer.
It is a practical operating system for coding agents.
A single coding-agent session is usually one context doing every job: scoping, implementing, testing, reviewing, documenting, and sometimes pretending it can approve the work it just wrote.
That is not how good software teams work.
The same model that wrote the bug should not be the only reviewer of the bug. The same chat that guessed an API should not be the final authority that the API is correct. A risky authentication change should not be treated like a typo fix. A refactor should not silently expand across unrelated parts of the codebase.
Most teams try to fix this with bigger prompts or “be careful” notes in CLAUDE.md , AGENTS.md , or editor rules. That helps a little, but the agent can still tunnel-vision, skip tests, lose context, make unrelated edits, or leave humans with a giant diff that is hard to trust.
AgentCrew exists to give coding agents a process before they touch the code.
We built AgentCrew after repeatedly running into the same problem with coding agents: they can write code fast, but they do not naturally follow a disciplined development cycle.
Every serious change needs more than implementation. The agent should understand the task, pick the right context, use the right skills, write or update tests, verify that those tests actually work, document the change, and prepare something a human can review.
Without that structure, we found ourselves repeating the same instructions again and again:
The failure mode is not theoretical. We have seen agent-assisted refactors create messy changes across a codebase, miss important context, and leave humans with days of cleanup. In one case, a refactor that should have been controlled ended up breaking enough pieces that fixing it manually took about a week. The problem was not that the model was useless. The problem was that it had no real process, no specialized role, no skill boundary, and no serious review gate.
We tried bigger prompts, project rules, internal Codex workflows, prompt stacks, and command-based approaches. Some helped. But they still required too much manual steering, especially when we wanted the agent itself to recognize when a task needed planning, testing, review, security attention, or a human decision.
AgentCrew is the system we wanted: not a magical autopilot, and not a replacement for developer judgment. It is a disciplined methodology for agentic coding — roles, skills, routing, handoffs, quality gates, memory, and approval boundaries — so coding agents become easier to trust, easier to review, and more useful in real development work.
You ask your normal coding agent for an outcome:
Add Google OAuth login to the signup page.
With AgentCrew loaded, the agent does not jump straight into editing files. It first follows the methodology:
Classify the request
A pure-bash classifier can identify the task type, risk level, lane, likely role, relevant skills, and quality gates.
Pick the right lane
Simple, low-risk work can use a fast lane. Risky, multi-step, security-sensitive, or architecture-changing work uses a fuller workflow.
Start with the right role
The agent may begin as an Advisor or Product Manager to clarify scope and acceptance criteria before implementation.
Confirm direction and estimate cost
For work that will spend meaningful model tokens, AgentCrew can surface an estimated cost before execution starts.
Plan and dispatch the work
After the user confirms the direction, AgentCrew prepares the work and can let the required agents run through the host agent or optional engine. The user should not need to run AgentCrew commands manually.
Implement as Developer
The Developer role focuses on minimal, relevant changes. It must avoid unrelated edits, hidden failures, unsafe commands, and premature merge decisions.
Check as Tester
The Tester role focuses on verification. It should run or propose the right tests and report what passed, failed, or could not be checked.
Review as Reviewer or Security Reviewer
Risky code paths, authentication, authorization, payments, migrations, secrets, and infrastructure changes get stricter review guidance.
Produce a handoff
The final output is not just “done.” It should include what changed, what was tested, what risks remain, and what needs human approval.
Stop at the human gate
AgentCrew does not treat the agent as the final approver. A human still decides what lands.
You stay in your normal agent chat. AgentCrew gives that chat a routing system, role instructions, handoff formats, state artifacts, and approval gates, so the work follows a team process instead of one long unstructured session.
Clone AgentCrew somewhere outside your project:
git clone https://github.com/mlguyYT/AgentCrew ~ /AgentCrew
Register AgentCrew with your host agent:
~ /AgentCrew/bin/agentcrew install
Verify the setup:
~ /AgentCrew/bin/agentcrew doctor
Then open your normal coding agent and start working as usual.
AgentCrew installs a small loader block into your host-agent instruction file, such as:
~/.claude/CLAUDE.md
~/.codex/AGENTS.md
~/.hermes/SOUL.md
The loader points the host agent to the AgentCrew methodology. Your project does not need a new runtime dependency.
agent-team/
├─ agents/ Role definitions such as Advisor, PM, Developer,
│ Tester, Reviewer, Security Reviewer, UX Reviewer,
│ Documentation Agent, and more
│
├─ tools/ CLI helpers, including the pure-bash task classifier
│
├─ playbooks/ How to route tasks, handle gates, prepare PRs,
│ manage rework, and save useful memory
│
├─ recipes/ Workflow recipes for features, bug fixes, refactors,
│ migrations, reviews, security-sensitive work, and more
│
├─ protocols/ Handoff format, state artifacts, approval boundaries,
│ and token-discipline rules
│
├─ context/ Route index and lane-specific context profiles
│
├─ checklists/ Pre-merge readiness, review checks, memory refreshes,
│ and other quality controls
│
├─ templates/ Task briefs, work plans, compact handoffs, review
│ reports, PR descriptions, and session summaries
│
└─ skills/ Technical and professional skills that can be loaded
only when relevant
bin/agentcrew install, doctor, status, classify, context, start,
brief, plan, ready, pr-pack, checkpoint, save-session,
restore-session, detect-project, preset
engine/ optional executable layer for host-agent orchestration,
cost previews, and multi-agent execution
The methodology is Markdown.
The classifier and setup helpers are shell scripts.
There is no daemon and no required runtime inside your project.
Not every task needs a heavy workflow.
AgentCrew separates work into lanes so the agent does not over-process simple tasks or under-process risky ones.
Used for small, low-risk work:
The goal is speed with basic discipline.
Used for work that needs stronger control:
authentication or authorization changes;
multi-file or multi-phase implementation.
The goal is safer execution, clearer handoffs, and better review.
AgentCrew gives your coding agent role-specific instructions instead of asking one generic chat to do everything.
Advisor — explains, guides, and helps reason before implementation.
Product Manager — clarifies scope, user value, acceptance criteria, and tradeoffs.
Developer — implements focused changes with minimal unrelated edits.
Tester — verifies behavior and reports what was tested.
Reviewer — reviews the implementation for correctness, maintainability, and risk.
Security Reviewer — checks sensitive paths such as auth, permissions, secrets, and unsafe operations.
UX Reviewer — reviews user-facing behavior and interaction quality.
Documentation Agent — updates or reviews documentation, examples, and usage notes.
Release Manager — helps prepare release notes, readiness checks, and release handoffs.
The point is not to pretend there are separate humans. The point is to make the agent use different constraints for different phases of work.
AgentCrew can keep project-local working state in .agent-state/ .
.agent-state/
├─ current-task.md
├─ work-plan.md
├─ human-decisions.md
├─ handoff.md
├─ test-report.md
├─ review-report.md
├─ pr-pack.md
└─ session-checkpoints/
This gives the agent a controlled memory surface for the current project.
The goal is not to store everything. The goal is to preserve the important parts:
what decisions were already made;
what still needs human approval.
Good agent memory should be compressive, not archival.
AgentCrew can detect or configure project presets so the methodology fits the stack.
A project preset can add stricter expectations, relevant skills, or extra review gates.
Project-local configuration should tighten the process. It should not remove human approval, skip required gates, or weaken core safety rules.
AgentCrew is built around a simple principle:
The agent can help prepare the work, but a human still owns the decision to land it.
Agents must not merge their own work.
Agents must not bypass branch pro

[truncated]
