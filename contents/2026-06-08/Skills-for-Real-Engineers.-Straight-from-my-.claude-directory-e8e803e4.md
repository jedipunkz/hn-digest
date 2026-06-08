---
source: "https://github.com/mattpocock/skills"
hn_url: "https://news.ycombinator.com/item?id=48443601"
title: "Skills for Real Engineers. Straight from my .claude directory"
article_title: "GitHub - mattpocock/skills: Skills for Real Engineers. Straight from my .claude directory. · GitHub"
author: "thunderbong"
captured_at: "2026-06-08T10:40:04Z"
capture_tool: "hn-digest"
hn_id: 48443601
score: 1
comments: 0
posted_at: "2026-06-08T10:38:03Z"
tags:
  - hacker-news
  - translated
---

# Skills for Real Engineers. Straight from my .claude directory

- HN: [48443601](https://news.ycombinator.com/item?id=48443601)
- Source: [github.com](https://github.com/mattpocock/skills)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T10:38:03Z

## Translation

タイトル: 本物のエンジニアのためのスキル。私の .claude ディレクトリから直接
記事のタイトル: GitHub - mattpocock/スキル: 本物のエンジニアのためのスキル。私の .claude ディレクトリから直接。 · GitHub
説明: 本物のエンジニアのためのスキル。私の .claude ディレクトリから直接。 - マットポコック/スキル

記事本文:
GitHub - mattpocock/skills: 本物のエンジニアのためのスキル。私の .claude ディレクトリから直接。 · GitHub
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
マットポコック
/
スキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラン

hes タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
94 コミット 94 コミット .claude-plugin .claude-plugin .out-of-scope .out-of-scope docs/ adr docs/ adr スクリプト スクリプト スキル スキル CLAUDE.md CLAUDE.md CONTEXT.md CONTEXT.md ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
バイブコーディングではなく、実際のエンジニアリングを行うために毎日使用しているエージェントのスキル。
実際のアプリケーションを開発するのは難しいです。 GSD、BMAD、Spec-Kit などのアプローチは、プロセスを所有することで支援しようとします。しかし、そうしている間、コントロールが奪われ、プロセス内のバグの解決が困難になります。
これらのスキルは、小さく、適応しやすく、構成可能なように設計されています。どのモデルでも動作します。これらは数十年にわたるエンジニアリングの経験に基づいています。彼らと一緒にハッキングしましょう。それらを自分のものにしてください。楽しむ。
これらのスキルの変更や私が作成した新しいスキルに遅れずについていきたい場合は、私のニュースレターで他の約 60,000 人の開発者に参加してください。
npx skill@latest add mattpocock/skills
必要なスキルと、それらをインストールするコーディング エージェントを選択します。必ず /setup-matt-pocock-skills を選択してください。
エージェントで /setup-matt-pocock-skills を実行します。それは次のことを行います:
どの問題トラッカーを使用するかを尋ねます (GitHub、Linear、またはローカル ファイル)。
チケットを優先順位付けするときに、チケットにどのようなラベルを適用するかを尋ねます ( /triage はラベルを使用します)
作成したドキュメントをどこに保存するかを尋ねます
私は、Claude Code、Codex、およびその他のコーディング エージェントでよく見られる一般的な障害モードを修正する方法として、これらのスキルを構築しました。
#1: エージェントは私が望むことをしてくれませんでした
「自分が何を望んでいるのか正確にわかる人はいない」
David Thomas & Andrew Hunt、実用的なプログラマー
問題。ソフトウェア開発における最も一般的な失敗モードは、位置ずれです。開発者はあなたが何を望んでいるのかを知っていると思います。そして、彼らが構築したものを見ると、それが理解されていなかったことに気づきます。

あなたはまったく。
これはAI時代でも全く同じです。あなたとエージェントの間にはコミュニケーションギャップがあります。この問題を解決するには、徹底的なセッションを行い、エージェントに構築しているものについて詳細な質問をしてもらいます。
/grill-with-docs - /grill-me と同じですが、さらに便利な機能が追加されています (以下を参照)
これらは私の最も人気のあるスキルです。これらは、開始前にエージェントと調整し、加えようとしている変更について深く考えるのに役立ちます。変更を加えたい場合は常にこれらを使用してください。
#2: エージェントが冗長すぎる
ユビキタス言語では、開発者間の会話とコードの表現はすべて同じドメイン モデルから派生します。
エリック・エヴァンス、ドメイン駆動設計
問題 : プロジェクトの開始時、開発者とソフトウェアを構築している人々 (ドメインの専門家) は通常、異なる言語を話します。
私もエージェントたちと同じ緊張感を感じました。通常、エージェントはプロジェクトに参加させられ、作業を進めながら専門用語を理解するよう求められます。したがって、1 つで済む単語を 20 個使用します。
この問題を解決するには、共有言語を使用します。これは、エージェントがプロジェクトで使用される専門用語を解読するのに役立つ文書です。
これは、私の course-video-manager リポジトリにある CONTEXT.md の例です。どちらが読みやすいでしょうか？
BEFORE : 「コースのセクション内のレッスンが「本物」になる (つまり、ファイル システム内にスポットが与えられる) と問題が発生します。」
AFTER : 「実体化カスケードに問題があります」
この簡潔さはセッションを重ねるたびに成果を上げます。
これは /grill-with-docs に組み込まれています。これは難しいセッションですが、AI との共通言語を構築し、ADR で説明が難しい決定を文書化するのに役立ちます。
これがどれほど強力であるかを説明するのは難しいです。これは、このリポジトリの中で最もクールなテクニックかもしれません。試してみてください。
共有言語には、冗長性を軽減する以外にも多くの利点があります。
ヴァール

変数、関数、およびファイルには、共有言語を使用して一貫した名前が付けられます。
その結果、エージェントにとってコードベースのナビゲーションが容易になります。
また、エージェントはより簡潔な言語にアクセスできるため、思考に費やすトークンも少なくなります。
「常に小さな計画を立ててください。フィードバックの速度がスピードの限界です。大きすぎるタスクを決して引き受けないでください。」
David Thomas & Andrew Hunt、実用的なプログラマー
問題 : あなたとエージェントが何を構築するかについて意見が一致しているとします。エージェントがまだクソを生成した場合はどうなりますか?
フィードバック ループを見てみましょう。生成されるコードが実際にどのように実行されるかについてのフィードバックがなければ、エージェントは盲目的に行動することになります。
解決策 : フィードバック ループの通常の部分 (静的型、ブラウザー アクセス、自動テスト) が必要です。
自動テストの場合、赤、緑、リファクタリングのループが重要です。ここでは、エージェントが最初に失敗したテストを作成し、次にそのテストを修正します。これは、エージェントに一貫したレベルのフィードバックを提供するのに役立ち、結果としてはるかに優れたコードが得られます。
あらゆるプロジェクトに組み込める /tdd スキルを構築しました。これはレッド・グリーン・リファクタリングを奨励し、エージェントに良いテストと悪いテストを作るための十分なガイダンスを与えます。
デバッグのために、デバッグのベスト プラクティスを単純なループにラップする /diagnose スキルも構築しました。
「システムの設計に毎日投資してください。」
ケント・ベック、エクストリーム・プログラミングの解説
「最高のモジュールは奥深いものです。シンプルなインターフェースから多くの機能にアクセスできるようになります。」
ジョン・アウスターハウト著、ソフトウェア設計の哲学
問題 : エージェントを使用して構築されたほとんどのアプリは複雑で、変更が困難です。エージェントはコーディングを大幅に高速化できるため、ソフトウェアのエントロピーも加速します。コードベースは前例のない速度で複雑化しています。
この問題を解決するのは、AI を活用した開発に対する根本的な新しいアプローチです: cari

コードの設計についてはNGです。
これは、これらのスキルのすべての層に組み込まれています。
/to-prd は、PRD を作成する前に、どのモジュールに触れているかについて質問します。
/zoom-out は、システム全体のコンテキストでコードを説明するようにエージェントに指示します。
そして重要なことに、/improve-codebase-architecture は泥団子になったコードベースを救出するのに役立ちます。数日に一度、コードベースで実行することをお勧めします。
ソフトウェア エンジニアリングの基礎がこれまで以上に重要になっています。これらのスキルは、あなたのキャリアの中で最高のアプリをリリースするのに役立つように、これらの基本を反復可能な実践に凝縮するための私の最善の努力です。楽しむ。
コード作業で日常的に使用しているスキル。
診断 — 困難なバグとパフォーマンスの低下に対する規律ある診断ループ: 再現→最小化→仮説→計測→修正→回帰テスト。
Grille-with-docs — 既存のドメイン モデルに照らして計画に挑戦し、用語を明確にし、CONTEXT.md と ADR をインラインで更新するグリル セッション。
トリアージ — トリアージ ロールのステート マシンを通じて問題をトリアージします。
Improvement-codebase-architecture — CONTEXT.md のドメイン言語と docs/adr/ の決定に基づいて、コードベースの深化の機会を見つけます。
setup-matt-pocock-skills — 他のエンジニアリング スキルが使用するリポジトリごとの構成 (問題トラッカー、トリアージ ラベルの語彙、ドメイン ドキュメントのレイアウト) を足場にします。 to-issues 、 to-prd 、 triage 、 Diagnostic 、 tdd 、 Improvement-codebase-architecture 、またはZoom-out を使用する前に、リポジトリごとに 1 回実行します。
tdd — 赤、緑、リファクタリング ループを使用したテスト駆動開発。一度に 1 つの垂直スライスごとに機能を構築したり、バグを修正したりします。
to-issues — 垂直スライスを使用して、あらゆる計画、仕様、または PRD を独立して取得可能な GitHub 問題に分割します。
to-prd — 現在の会話コンテキストを PRD に変換し、GitHub の問題として送信します。

インタビューはありません。すでに話し合った内容をまとめるだけです。
ズームアウト — エージェントにズームアウトして、コードの馴染みのないセクションについてより広範なコンテキストやより高いレベルの視点を提供するように指示します。
プロトタイプ — 設計を具体化するための使い捨てプロトタイプを構築します。状態/ビジネス ロジックの質問用の実行可能なターミナル アプリ、または 1 つのルートから切り替え可能ないくつかの根本的に異なる UI バリエーションのいずれかです。
コード固有ではない、一般的なワークフロー ツール。
caveman — 超圧縮通信モード。完全な技術的精度を維持しながら、フィラーをドロップすることでトークンの使用量を最大 75% 削減します。
グリルミー — デシジョンツリーのすべての分岐が解決されるまで、計画や設計について執拗にインタビューを受けます。
handoff — 現在の会話をハンドオフ文書に圧縮して、別のエージェントが作業を継続できるようにします。
write-a-skill — 適切な構造、段階的な開示、バンドルされたリソースを備えた新しいスキルを作成します。
常に手元に置いているものの、めったに使用しないツール。
git-guardrails-claude-code — クロード コード フックをセットアップして、危険な git コマンド (push、reset --hard、clean など) を実行前にブロックします。
merge-to-shoehorn — テスト ファイルを as type アサーションから @total-typescript/shoehorn に移行します。
scaffold-exercises — セクション、問題、解決策、および説明を含む演習ディレクトリ構造を作成します。
setup-pre-commit — lint ステージング、Prettier、型チェック、およびテストを使用して Husky の事前コミット フックをセットアップします。
本物のエンジニアのためのスキル。私の .claude ディレクトリから直接。
読み込み中にエラーが発生しました。このページをリロードしてください。
10.6k
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Skills for Real Engineers. Straight from my .claude directory. - mattpocock/skills

GitHub - mattpocock/skills: Skills for Real Engineers. Straight from my .claude directory. · GitHub
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
mattpocock
/
skills
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
94 Commits 94 Commits .claude-plugin .claude-plugin .out-of-scope .out-of-scope docs/ adr docs/ adr scripts scripts skills skills CLAUDE.md CLAUDE.md CONTEXT.md CONTEXT.md LICENSE LICENSE README.md README.md View all files Repository files navigation
My agent skills that I use every day to do real engineering - not vibe coding.
Developing real applications is hard. Approaches like GSD, BMAD, and Spec-Kit try to help by owning the process. But while doing so, they take away your control and make bugs in the process hard to resolve.
These skills are designed to be small, easy to adapt, and composable. They work with any model. They're based on decades of engineering experience. Hack around with them. Make them your own. Enjoy.
If you want to keep up with changes to these skills, and any new ones I create, you can join ~60,000 other devs on my newsletter:
npx skills@latest add mattpocock/skills
Pick the skills you want, and which coding agents you want to install them on. Make sure you select /setup-matt-pocock-skills .
Run /setup-matt-pocock-skills in your agent. It will:
Ask you which issue tracker you want to use (GitHub, Linear, or local files)
Ask you what labels you apply to tickets when you triage them ( /triage uses labels)
Ask you where you want to save any docs we create
I built these skills as a way to fix common failure modes I see with Claude Code, Codex, and other coding agents.
#1: The Agent Didn't Do What I Want
"No-one knows exactly what they want"
David Thomas & Andrew Hunt, The Pragmatic Programmer
The Problem . The most common failure mode in software development is misalignment. You think the dev knows what you want. Then you see what they've built - and you realize it didn't understand you at all.
This is just the same in the AI age. There is a communication gap between you and the agent. The fix for this is a grilling session - getting the agent to ask you detailed questions about what you're building.
/grill-with-docs - same as /grill-me , but adds more goodies (see below)
These are my most popular skills. They help you align with the agent before you get started, and think deeply about the change you're making. Use them every time you want to make a change.
#2: The Agent Is Way Too Verbose
With a ubiquitous language, conversations among developers and expressions of the code are all derived from the same domain model.
Eric Evans, Domain-Driven-Design
The Problem : At the start of a project, devs and the people they're building the software for (the domain experts) are usually speaking different languages.
I felt the same tension with my agents. Agents are usually dropped into a project and asked to figure out the jargon as they go. So they use 20 words where 1 will do.
The Fix for this is a shared language. It's a document that helps agents decode the jargon used in the project.
Here's an example CONTEXT.md , from my course-video-manager repo. Which one is easier to read?
BEFORE : "There's a problem when a lesson inside a section of a course is made 'real' (i.e. given a spot in the file system)"
AFTER : "There's a problem with the materialization cascade"
This concision pays off session after session.
This is built into /grill-with-docs . It's a grilling session, but that helps you build a shared language with the AI, and document hard-to-explain decisions in ADR's.
It's hard to explain how powerful this is. It might be the single coolest technique in this repo. Try it, and see.
A shared language has many other benefits than reducing verbosity:
Variables, functions and files are named consistently , using the shared language
As a result, the codebase is easier to navigate for the agent
The agent also spends fewer tokens on thinking , because it has access to a more concise language
"Always take small, deliberate steps. The rate of feedback is your speed limit. Never take on a task that’s too big."
David Thomas & Andrew Hunt, The Pragmatic Programmer
The Problem : Let's say that you and the agent are aligned on what to build. What happens when the agent still produces crap?
It's time to look at your feedback loops. Without feedback on how the code it produces actually runs, the agent will be flying blind.
The Fix : You need the usual tranche of feedback loops: static types, browser access, and automated tests.
For automated tests, a red-green-refactor loop is critical. This is where the agent writes a failing test first, then fixes the test. This helps give the agent a consistent level of feedback that results in far better code.
I've built a /tdd skill you can slot into any project. It encourages red-green-refactor and gives the agent plenty of guidance on what makes good and bad tests.
For debugging, I've also built a /diagnose skill that wraps best debugging practices into a simple loop.
"Invest in the design of the system every day ."
Kent Beck, Extreme Programming Explained
"The best modules are deep. They allow a lot of functionality to be accessed through a simple interface."
John Ousterhout, A Philosophy Of Software Design
The Problem : Most apps built with agents are complex and hard to change. Because agents can radically speed up coding, they also accelerate software entropy. Codebases get more complex at an unprecedented rate.
The Fix for this is a radical new approach to AI-powered development: caring about the design of the code.
This is built in to every layer of these skills:
/to-prd quizzes you about which modules you're touching before creating a PRD
/zoom-out tells the agent to explain code in the context of the whole system
And crucially, /improve-codebase-architecture helps you rescue a codebase that has become a ball of mud. I recommend running it on your codebase once every few days.
Software engineering fundamentals matter more than ever. These skills are my best effort at condensing these fundamentals into repeatable practices, to help you ship the best apps of your career. Enjoy.
Skills I use daily for code work.
diagnose — Disciplined diagnosis loop for hard bugs and performance regressions: reproduce → minimise → hypothesise → instrument → fix → regression-test.
grill-with-docs — Grilling session that challenges your plan against the existing domain model, sharpens terminology, and updates CONTEXT.md and ADRs inline.
triage — Triage issues through a state machine of triage roles.
improve-codebase-architecture — Find deepening opportunities in a codebase, informed by the domain language in CONTEXT.md and the decisions in docs/adr/ .
setup-matt-pocock-skills — Scaffold the per-repo config (issue tracker, triage label vocabulary, domain doc layout) that the other engineering skills consume. Run once per repo before using to-issues , to-prd , triage , diagnose , tdd , improve-codebase-architecture , or zoom-out .
tdd — Test-driven development with a red-green-refactor loop. Builds features or fixes bugs one vertical slice at a time.
to-issues — Break any plan, spec, or PRD into independently-grabbable GitHub issues using vertical slices.
to-prd — Turn the current conversation context into a PRD and submit it as a GitHub issue. No interview — just synthesizes what you've already discussed.
zoom-out — Tell the agent to zoom out and give broader context or a higher-level perspective on an unfamiliar section of code.
prototype — Build a throwaway prototype to flesh out a design — either a runnable terminal app for state/business-logic questions, or several radically different UI variations toggleable from one route.
General workflow tools, not code-specific.
caveman — Ultra-compressed communication mode. Cuts token usage ~75% by dropping filler while keeping full technical accuracy.
grill-me — Get relentlessly interviewed about a plan or design until every branch of the decision tree is resolved.
handoff — Compact the current conversation into a handoff document so another agent can continue the work.
write-a-skill — Create new skills with proper structure, progressive disclosure, and bundled resources.
Tools I keep around but rarely use.
git-guardrails-claude-code — Set up Claude Code hooks to block dangerous git commands (push, reset --hard, clean, etc.) before they execute.
migrate-to-shoehorn — Migrate test files from as type assertions to @total-typescript/shoehorn.
scaffold-exercises — Create exercise directory structures with sections, problems, solutions, and explainers.
setup-pre-commit — Set up Husky pre-commit hooks with lint-staged, Prettier, type checking, and tests.
Skills for Real Engineers. Straight from my .claude directory.
There was an error while loading. Please reload this page .
10.6k
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
