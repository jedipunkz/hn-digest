---
source: "https://github.com/alxshelepenok/grove"
hn_url: "https://news.ycombinator.com/item?id=49362786"
title: "Show HN: Grove, a formal workflow protocol for long-running AI coding agents"
article_title: "GitHub - alxshelepenok/grove: A formal workflow protocol that keeps AI coding agents on track through machine-enforced invariants, verified evidence, and structured context. Designed to keep deep, long-running projects coherent across sessions, agents, and months. · GitHub"
image: "https://repository-images.githubusercontent.com/1229298801/0ea5ed66-b950-4fcd-b38f-458c0dd3df2d"
author: "alxshelepenok"
captured_at: "2026-08-19T15:21:49Z"
capture_tool: "hn-digest"
hn_id: 49362786
score: 1
comments: 0
posted_at: "2026-08-19T15:19:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Grove, a formal workflow protocol for long-running AI coding agents

- HN: [49362786](https://news.ycombinator.com/item?id=49362786)
- Source: [github.com](https://github.com/alxshelepenok/grove)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T15:19:53Z

## Translation

タイトル: Show HN: Grove、長期実行 AI コーディング エージェント向けの正式なワークフロー プロトコル
記事のタイトル: GitHub - alxshelepenok/grove: 機械強制の不変条件、検証された証拠、構造化されたコンテキストを通じて AI コーディング エージェントを軌道に乗せ続ける正式なワークフロー プロトコル。セッション、エージェント、月にまたがる、深く長期にわたるプロジェクトの一貫性を維持するように設計されています。 · GitHub
説明: 機械によって強制された不変条件、検証された証拠、構造化されたコンテキストを通じて AI コーディング エージェントを軌道に乗せ続ける正式なワークフロー プロトコル。セッション、エージェント、月にまたがる、深く長期にわたるプロジェクトの一貫性を維持するように設計されています。 - アルクスシェルペノック/グローブ

記事本文:
GitHub - alxshelepenok/grove: 機械強制の不変条件、検証された証拠、構造化されたコンテキストを通じて AI コーディング エージェントを軌道に乗せ続ける正式なワークフロー プロトコル。セッション、エージェント、月にまたがる、深く長期にわたるプロジェクトの一貫性を維持するように設計されています。 · GitHub
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
アルクスシェルペノック
/
木立
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
80 コミット 80 コミット フォルダーとファイル
。

github .github .grove .grove bin bin docs docs パッケージ パッケージ テスト テスト .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス Makefile Makefile README.md README.md SECURITY.md SECURITY.md install.ps1 install.ps1 install.ps1.sig install.ps1.sig install.sh install.sh install.sh.sig install.sh.sig マニフェスト.json マニフェスト.json マニフェスト.json.sig マニフェスト.json.sig すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コードは成長します。コンテキストウィンドウはそうではありません。木立を植える 🌳
検証済みの証拠を基にしたグラフによる推論。機械によって強制された不変条件、検証された証拠、構造化されたコンテキストを通じて、AI コーディング エージェントを軌道に乗せ続ける正式なワークフロー プロトコル。セッション、エージェント、月にまたがる、深く長期にわたるプロジェクトの一貫性を維持するように設計されています。
私たちはもはやコードを書くだけではありません。私たちは、1 日 20 件のコミットで運行する列車の手綱を握り、列車が脱線しないように注意します。
原子的な進歩: すべてのタスクは、主張されるだけでなく、機械的に証明されます。
圧縮に対する階層: コンテキストは厳密に構造化されており、損失を伴う要約は決して行われません。
絶対的な再現性: すべてのアクションが記録されるため、あらゆるステップ、決定、仮定を再現して追跡することができます。
エージェント間の継続性: エージェントが目標の途中で停止した場合、別のクライアントまたはモデル上の別のエージェントが検証済みの状態から再開します。
完全な所有権: すべての目標、決定、タスク、質問、仮定は、完全にあなたに属する永続的な履歴グラフに保存されます。
Grove は、単一のチェックサム付きロックファイルに保存された決定論的な不変条件としてルールを適用します。エージェントは、反証可能な証拠なしに完了した作業を宣言したり、準備ができていないタスクを開始したり、進捗状況を幻覚したりすることはできません。
非可逆プロンプト圧縮または要約の代わりに、Grove 構造

機械チェック可能なエッジを持つ型付き推論グラフに状態を投影します。現在のステップに必要な実行パケットを正確にルーティングし、それ以上はルーティングしません。
Claude Code、Codex、Gemini CLI、Cursor、Windsurf、Cline、GitHub Copilot、およびその他のエージェントで動作します。 Grove には CLI、組み込み MCP サーバー、ドロップイン エージェント スキル バンドルが含まれているため、すぐにワークスペースに統合できます。
以下にデスクトップ アプリに関するメモを示します。これは、プロジェクトのグラフ、証拠、健全性を一目で確認できるようにするために存在しますが、まだ実験段階であり、macOS ではまだテストされていません。 CLI と MCP サーバーは、テスト済みの安定したインターフェイスです。 UI は JavaScript フレームワークを使用せずに、主にプレーンな HTML、CSS、および Tauri を使用して実装されているため、高速性が維持され、グラフは WebGL でレンダリングされます。
推論グラフをライブで探索します。
143 ノードでの同じビュー (履歴を含む)。
エージェントが必要とするのは 1 つの作業項目だけであり、それ以上は必要ありません。
作業項目は提案から完了まで追跡されます。
エリアごとの目標、作業、およびコンテンツの健全性。
クリティカル パスを備えたテーマごとにグループ化された作品。
プロジェクトが成長するにつれてエージェントはコンテキストを失う
長期にわたって活動するエージェントは状況記憶喪失に悩まされます。セッションが長くなるにつれて、意思決定、仮定、依存関係が視界から消えてしまいます。関連する失敗としては、信頼性の低い自己報告が挙げられます。エージェントは、十分な証拠なしに作業が完了したと宣言します。これは欺瞞から起こるのではなく、環境の中にそれを妨げるものがないために起こります。
標準的なタスク トラッカーは本質的に実行者を信頼します。人間が Jira のボックスにチェックを入れると、作業が完了したものとみなされます。それは人間にとって合理的なデフォルトです。自律エージェントの場合、これはサイレント障害モードです。時期尚早な「完了」宣言は、長いセッションにわたって構造的エラーを引き起こし、コードベースが誰もが信じているものと一致しなくなります。

d 特定の行が存在する理由を誰も説明できません。
これは理論的なものではありません。 Grove は、粗雑な Markdown 仮説から、現在ではブロックチェーンと暗号化に重点を置いた 20 万行以上の Rust と TypeScript のクローズドソース制作プロジェクトである Merlin Guild を管理するプロトコルに進化しました。このプロトコルは、エージェントがワークフローを正しく記憶したり従ったりすることに依存していないため、拡張性があります。
圧縮ではなく構造化
文脈健忘症に対する明らかな対応は、要約、圧縮、検索などのさらなるツールです。これらはすべてコンテキストを圧縮し、圧縮により情報が失われますが、その損失が問題にならないことを望みます。
Grove はコンテキストを圧縮しません。それを構造化します。
プロジェクトが進化するにつれて、その作業は領域、目標、質問、仮定、決定、および実行可能な作業項目に継続的に編成されます。この階層により、エージェントはプロジェクト履歴全体をコンテキスト ウィンドウに繰り返し入力する代わりに、現在のステップに関連するコンテキストを識別する方法が提供されます。
その結果、継続性が向上するだけではありません。また、トークンの使用量も削減されます。エージェントは、単に現在のタスクがどこにあるのか、なぜ存在するのかを回復するために、無関係なプロジェクト コンテキストをロードしなくなりました。
単一のプロンプトでコードベース全体をマップする
Grove がアタッチされた生のコードベースをフロンティア モデルに指定するとどうなりますか?
プロトコルがなければ、エージェントはファイルを流し読みし、簡単な要約を幻覚で見せ、数ステップ後にプロットを失います。 Grove を使用すると、領域の地図が作成されます。複雑なコードベースでの実際のテストでは、単一のプロンプトで 5 時間の自律的な Discovery セッションが開始されました。
エージェントはコンテキスト ウィンドウの制限を使い果たしましたが、その結果は、テスト カバレッジ ギャップとロジック バグを対象とした、44 のエリア、47 の目標、および 89 の作業項目という完全に入力された推論グラフになりました。単純な To Do リストをただダンプしただけではありません。それはその作戦を正式に宣言した

n 個の未知の項目を質問として抽出し、関連する取り組みをテーマにグループ化しました。
プロジェクトをブートストラップするために使用される正確なプロンプトは次のとおりです。
Grove スキルをロードし、MCP サーバーに接続します。スキルを十分に学習してプロトコルの制約を理解してください。
あなたの目的は、コードベース全体にわたって完全な検出フェーズを実行することです。実装コードはまだ作成しないでください。あなたの仕事はプロジェクト分析とグラフ構築です。
プロジェクトをエリアに分割します。 1 つのパッケージに複数の論理領域が含まれる場合があることに注意してください。
各領域を反復処理して目標を作成し、不足しているテスト カバレッジや明示的なロジックのバグに対処します。
目標ごとに、Grove プロトコルに従って厳密にフォーマットされた作業項目を作成します。
関連する作業をテーマにグループ化します。
未解決の未知のものを質問として宣言し、必要に応じてアーキテクチャ上の選択を決定として形式化します。
Grove は準備完了の定義 (DoR) を強制しているため、エージェントは進捗状況を偽装できませんでした。翌日、エージェントは体系的に配信フェーズを実行し、75 件のタスクを Done にしました。残りのタスクは、提案された状態または準備完了状態で適切に停止し、エージェントが提起した妨げとなる質問に対する人間の回答を待っていました。
1 つのプロジェクト状態、複数のインターフェイス
記憶の単位は会話ではありません。それがプロジェクトです。
進捗の単位はエージェントの主張ではありません。検証された証拠です。
Grove は、丁寧な即時指示を、厳密で機械的に強制される手順に置き換えます。エージェントは、CLI、MCP、デスクトップ インターフェイスを通じてプロジェクトの状態と対話します。プロトコル自体がルールを強制します。
証拠の記録がなければ、作業項目に完了のマークを付けることはできません。
すべての前提条件が機械によって検証されるまで、作業を開始することはできません。
目標の進捗状況を手動で更新することはできません。フィットネス デルタは、近い時間にアトミックに適用されるか、まったく適用されません。
状態は p を満たすか、

ロトコル、そうでなければグローブはそれを進めることを拒否します。エージェントは、作業が完了したと主張するだけではプロトコルの状態を進めることはできません。
エージェントに必要なコンテキストのみを提供する
Grove は、エージェントに製品を厳密に型指定された階層に分解することを強制します。タスクの生のリストの代わりに、計画コンテキストのすべての部分が明示的にノード分類に存在します。サイドドキュメントやエージェントの内部状態には何も隠されていません。
型付きエッジはこれらのノードをグラフに接続し、ブロックは非循環のままになります。この構造により、Grove はなぜこのタスクが存在するのか、両方に答えることができます。それを変更すると何が壊れますか?エージェントの内部状態に依存せずに。
このグラフが存在すると、上で約束したルーティングが機械的になります。 grove next は現在のステップを選択します。 grove パケットはそのコンテキストを正確に出力します。コンテキストウィンドウには無関係なものは何も表示されません。
コード行に触れる前に、エージェントは因果関係コーンをクエリします。 grove packet W-NN --cone は、後方コーン (トポロジー短縮順序で最初に終了する必要があるすべてのもの)、前方コーン (この項目が変更された場合の爆発半径)、および影響を受けるゴールごとの脆弱性スコアをマップします。
推論を再利用可能な知識に変える
Grove は、デュアルトラック アジャイル、仮説駆動開発、ADR、継続的発見、Cynefin、Mikado メソッドに基づいた、AI 駆動のソフトウェア開発のための統一方法論を導入しています。これらの影響は、自律 LLM エージェントの制約に基づいて設計された単一のワークフローに統合され、機械チェック可能な不変条件によって強制されます。
ほとんどの AI ワークフローは小さなウォーターフォールです。すべてを指定してから、すべてを構築します。 Grove は検出と配信を並行して実行し、各トラックが他のトラックにフィードします。
Discovery は未知の部分をオープンにして運用可能にします。質問は反証可能な仮定になります。検証された結果が厳選されたディスになる

カバーズ。
Delivery では、準備ができた作業項目が実行され、それらの前提に基づいて検証済みのコードが作成されます。
関節は機械式です。質問は目標に対して行われます。仮定は作業をターゲットにし、それをゲートします。発見は次の目標を導きます。完成した仕事は蒸留されて発見に戻ります。テストで仮定が反証されると、計画はすぐに作り直されます。依存的な仕事は壊れた基盤の上では進められません。
グローブは、蒸留された知識を永続的な真実として扱いません。プロジェクトが進化するにつれて、発見は古くなってしまう可能性があります。アンカーを再アクティブ化するには、新しいアンカーが必要です。蒸留負債は黙って蓄積されるのではなく追跡され、目標の適合性は終了するたびに再導出されます。
ダッシュボードには、意図した状態ではなく、すべての変異後の実際のプロトコルの状態が反映されます。 Grove 氏は、プロジェクトは進行し続けると想定しています。つまり、知識は変化し、仮定は無効になり、未完成の推論が蓄積されます。このプロトコルは、これらの変更をプロジェクト履歴に消すのではなく、表示できるようにします。
ディスカバリーとデリバリーは並行して実行されます (デュアルトラック・アジャイル、ケイガン)。ワークアイテムは、すべての未解決の質問とワークアイテムを妨げる未検証の仮定が Discovery で解決されるまで、配信に入ることができません。
すべての実行可能ユニットには明示的な承認がある

[切り捨てられた]

## Original Extract

A formal workflow protocol that keeps AI coding agents on track through machine-enforced invariants, verified evidence, and structured context. Designed to keep deep, long-running projects coherent across sessions, agents, and months. - alxshelepenok/grove

GitHub - alxshelepenok/grove: A formal workflow protocol that keeps AI coding agents on track through machine-enforced invariants, verified evidence, and structured context. Designed to keep deep, long-running projects coherent across sessions, agents, and months. · GitHub
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
alxshelepenok
/
grove
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
80 Commits 80 Commits Folders and files
.github .github .grove .grove bin bin docs docs packages packages tests tests .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md install.ps1 install.ps1 install.ps1.sig install.ps1.sig install.sh install.sh install.sh.sig install.sh.sig manifest.json manifest.json manifest.json.sig manifest.json.sig View all files Repository files navigation
Code grows. Context windows don't. Plant a Grove 🌳
G raph-driven R easoning O ver V erified E vidence. A formal workflow protocol that keeps AI coding agents on track through machine-enforced invariants, verified evidence, and structured context. Designed to keep deep, long-running projects coherent across sessions, agents, and months.
We no longer just write code. We hold the reins of a train running at 20 commits a day, making sure it doesn't derail.
Atomic progress: every task is mechanically proven, not just claimed.
Hierarchy over compression: context is strictly structured, never lossily summarized.
Absolute reproducibility: every action is recorded, so any step, decision, or assumption can be replayed and retraced.
Continuity across agents: if an agent stops mid-goal, another agent on another client or model resumes from verified state.
Total ownership: every goal, decision, task, question, and assumption is stored in a persistent historical graph that belongs entirely to you.
Grove enforces rules as deterministic invariants stored in a single checksummed lockfile. The agent cannot declare work done without falsifiable evidence, start unready tasks, or hallucinate progress.
Instead of lossy prompt compression or summarization, Grove structures project state into a typed reasoning graph with machine-checkable edges. It routes exactly the execution packet needed for the current step, and nothing more.
Works with Claude Code, Codex, Gemini CLI, Cursor, Windsurf, Cline, GitHub Copilot, and any other agent. Grove includes a CLI, a built-in MCP server, and a drop-in agent skill bundle so you can integrate it into your workspace immediately.
A note on the desktop app shown below. It exists to make the graph, the evidence, and the health of a project visible at a glance, but it is still experimental and has not been tested on macOS yet. The CLI and the MCP server are the stable, tested interfaces. The UI is implemented without any JavaScript framework, mostly plain HTML, CSS, and Tauri, which keeps it fast, and the graph is rendered with WebGL.
The reasoning graph, explored live.
The same view at 143 nodes, history included.
All an agent needs for one work item, nothing more.
Work items tracked from proposal to done.
Goals, work, and content health per area.
Work grouped by theme, with the critical path.
Agents lose context as projects grow
Long-running agents suffer from context amnesia. Decisions, assumptions, and dependencies drift out of view as the session grows. A related failure is unreliable self-reporting : the agent declares work complete without sufficient evidence. This happens not out of deception, but because nothing in the environment prevents it.
Standard task trackers inherently trust the executor. When a human checks a box in Jira, the work is assumed done. That is a reasonable default for humans. For autonomous agents it is a silent failure mode. Premature "done" declarations compound across long sessions into structural errors, until the codebase no longer matches what anyone believes it to be, and nobody can say why a given line exists.
This is not theoretical. Grove evolved from a crude Markdown hypothesis into a protocol that today manages Merlin Guild, a closed-source production project of 200k+ lines of Rust and TypeScript, heavy on blockchain and cryptography. It scales because the protocol does not rely on the agent remembering or obeying the workflow correctly.
Structure instead of compression
The obvious response to context amnesia is more tooling: summarization, compaction, retrieval. All of these compress context, and compression loses information while hoping the loss does not matter.
Grove does not compress context. It structures it.
As a project evolves, its work is continuously organized into areas, goals, questions, assumptions, decisions, and executable work items. This hierarchy gives the agent a way to identify the context relevant to the current step instead of repeatedly carrying the entire project history into its context window.
The result is not just better continuity; it also reduces token usage. The agent no longer loads unrelated project context simply to recover where it is and why the current task exists.
Map the entire codebase in a single prompt
What happens when you point a frontier model at a raw codebase with Grove attached?
Without a protocol, an agent skims the files, hallucinates a quick summary, and loses the plot after a few steps. With Grove, it maps the territory. In a real-world test on a complex codebase, a single prompt initiated a 5-hour autonomous Discovery session.
The agent burned through its context window limit, but the result was a fully populated reasoning graph: 44 Areas, 47 Goals, and 89 Work items targeting test coverage gaps and logic bugs. It didn't just dump a flat to-do list; it formally declared its open unknowns as Questions and grouped related efforts into Themes.
Here is the exact prompt used to bootstrap the project:
Load the Grove skill and connect to the MCP server. Study the skill fully to understand the protocol constraints.
Your objective is to run a complete Discovery phase across the entire codebase. Do not write implementation code yet, your job is project analysis and graph construction.
Break the project down into Areas. Note that a single package may contain multiple logical areas.
Iterate through each Area and create Goals to address missing test coverage and explicit logic bugs.
For each Goal, create Work items strictly formatted according to the Grove protocol.
Group related work into Themes.
Declare open unknowns as Questions and formalize architectural choices as Decisions where necessary.
Because Grove enforces a Definition of Ready (DoR), the agent couldn't fake progress. Over the next day, the agent systematically executed the delivery phase, turning 75 tasks to Done . The remaining tasks correctly stalled in proposed or ready states, waiting for human answers to the blocking Questions the agent had raised.
One project state, multiple interfaces
The unit of memory is not the conversation. It is the project.
The unit of progress is not the agent's claim. It is verified evidence.
Grove replaces polite prompt instructions with a strict, mechanically enforced protocol. Agents interact with project state through the CLI, MCP, and desktop interfaces. The protocol itself enforces the rules:
A work item cannot be marked done without an evidence record.
Work cannot start until every precondition is machine-verified.
Goal progress cannot be updated by hand; fitness deltas are applied atomically at close time or not at all.
The state either satisfies the protocol, or Grove refuses to advance it. An agent cannot advance protocol state by merely claiming that work is complete.
Give agents only the context they need
Grove forces the agent to decompose the product into a strict typed hierarchy. Instead of a raw list of tasks, every piece of planning context lives explicitly in a node taxonomy. Nothing hides in side documents or the agent's internal state:
Typed edges connect these nodes into a graph, with blocks remaining acyclic. This structure lets Grove answer both why does this task exist? and what breaks if I change it? without relying on the agent's internal state.
Once this graph exists, the routing promised above becomes mechanical. grove next picks the current step; grove packet emits exactly its context. Nothing irrelevant enters the context window.
Before touching a line of code, the agent queries the causality cone . grove packet W-NN --cone maps the backward cone (everything that must finish first, in topological contraction order), the forward cone (the blast radius if this item changes), and a fragility score per affected goal.
Turn reasoning into reusable knowledge
Grove introduces a unified methodology for AI-driven software development, informed by Dual-Track Agile, Hypothesis-Driven Development, ADRs, Continuous Discovery, Cynefin, and the Mikado method. These influences are integrated into a single workflow designed around the constraints of autonomous LLM agents and enforced through machine-checkable invariants.
Most AI workflows are tiny waterfalls: specify everything, then build everything. Grove runs discovery and delivery in parallel, each track feeding the other:
Discovery takes open unknowns and operationalizes them. A question becomes a falsifiable assumption; validated outcomes become curated discoveries.
Delivery executes ready work items and writes verified code on top of those assumptions.
The joints are mechanical. Questions are asked against goals; assumptions target work and gate it; discoveries guide the next goals; finished work distills back into discoveries. When a test falsifies an assumption, the plan reshapes at once; dependent work cannot proceed on a broken foundation.
Grove does not treat distilled knowledge as permanent truth. Discoveries can become stale as the project evolves; reactivating one requires a fresh anchor. Distillation debt is tracked rather than silently accumulating, while goal fitness is re-derived on every close.
The dashboard reflects actual protocol state after every mutation, not intended state. Grove assumes that a project keeps moving: knowledge changes, assumptions are invalidated, and unfinished reasoning accumulates. The protocol makes those changes visible instead of letting them disappear into project history.
Discovery and Delivery run in parallel (Dual-Track Agile, Cagan). A work item cannot enter Delivery until every open question and unvalidated assumption that blocks it is resolved in Discovery.
Every executable unit has explicit acceptance

[truncated]
