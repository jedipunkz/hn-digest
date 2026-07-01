---
source: "https://github.com/hi-mundo/SIGNAL"
hn_url: "https://news.ycombinator.com/item?id=48746626"
title: "Signal: A draft UX framework for LLM-based systems"
article_title: "GitHub - hi-mundo/SIGNAL: SIGNAL is a pattern framework for LLM UX, human-AI interaction, conversational AI, agent UX, semantic clarity, cognitive load, and mostly user experience in LLM-based systems. · GitHub"
author: "hi-mundo"
captured_at: "2026-07-01T13:47:18Z"
capture_tool: "hn-digest"
hn_id: 48746626
score: 1
comments: 0
posted_at: "2026-07-01T13:41:34Z"
tags:
  - hacker-news
  - translated
---

# Signal: A draft UX framework for LLM-based systems

- HN: [48746626](https://news.ycombinator.com/item?id=48746626)
- Source: [github.com](https://github.com/hi-mundo/SIGNAL)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T13:41:34Z

## Translation

タイトル: Signal: LLM ベースのシステム向けのドラフト UX フレームワーク
記事のタイトル: GitHub - hi-mundo/SIGNAL: SIGNAL は、LLM UX、人間と AI の対話、会話型 AI、エージェント UX、セマンティックの明瞭さ、認知負荷、および主に LLM ベースのシステムにおけるユーザー エクスペリエンスのためのパターン フレームワークです。 · GitHub
説明: SIGNAL は、LLM UX、人間と AI の対話、会話型 AI、エージェント UX、セマンティックの明瞭さ、認知負荷、および主に LLM ベースのシステムにおけるユーザー エクスペリエンスのためのパターン フレームワークです。 - ハイムンド/シグナル

記事本文:
GitHub - hi-mundo/SIGNAL: SIGNAL は、LLM UX、人間と AI の対話、会話型 AI、エージェント UX、意味の明確さ、認知負荷、および主に LLM ベースのシステムにおけるユーザー エクスペリエンスのためのパターン フレームワークです。 · GitHub
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
ハイムンド
/
シグナル
公共
ノーティ

フィクション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット ドキュメント ドキュメント プロファイル プロファイル テンプレート テンプレート COTRIBUTING.md CONTRIBUTING.md CONTRIBUTORS.md CONTRIBUTORS.md FRAMEWORK.md FRAMEWORK.md GOVERNANCE.md GOVERNANCE.md ライセンス ライセンス README.md README.md RESEARCH_AND_BENCHMARKS.md RESEARCH_AND_BENCHMARKS.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
自然 AI 言語のセマンティック インタラクション ガイドライン
SIGNAL は、LLM UX、人間と AI の対話、会話型 AI、エージェント UX、セマンティックの明瞭さ、認知負荷、そして最も重要な、LLM ベースのシステムにおけるユーザー エクスペリエンスのためのパターン フレームワークです。
SIGNAL は現在鋭意開発中です。
このフレームワークは、人間と AI のインタラクション、会話デザイン、認知アクセシビリティ、平易な言語、生成 AI UX、AI ガバナンス、LLM ユーザー調査における隣接する研究を研究した後に作成されました。
これらの分野には強力なアイデア、ガイドライン、実証研究が含まれていますが、主に LLM ベースのシステムでのユーザー エクスペリエンス、特に通信層 (セマンティクス、インテント、グラウンディング、ナビゲーション、エージェンシー、および認知負荷) に特化した、焦点を絞ったオープンな README ファーストのパターン フレームワークを見つけることができませんでした。
SIGNAL はゼロからの発明のクレームではありません。
これは統合の試みであり、研究論文、製品ガイドライン、アクセシビリティに関するメモ、ガバナンス フレームワーク、ベンチマーク文化、プロンプト エンジニアリング、およびチャットボットの設計慣行にわたって断片化されているように見えるスペースの実用的なパターン ランゲージです。
シグナルは、情報の送信、コマンドの伝達、または警告として使用されるジェスチャ、トークン、またはアクションです。
会話がインターフェイスである場合、コミュニケーションは製品エクスペリエンスです。
ユーザーは書きたくない

完璧なプロンプト。
彼らは、短いメッセージ、曖昧な言及、比喩、慣用句、不完全な文脈、不満、手抜き、修正、間接的な要求など、人間のようにコミュニケーションをとります。
LLM は、言語の妥当な継続を生成するように最適化されています。製品エクスペリエンスには、妥当な継続以上のものが必要です。目に見える意図、根拠、進捗状況、主体性、コスト、価値が必要です。
SIGNAL は、乱雑な人間の言語と有用な AI の動作の間のコミュニケーション層を設計します。
手紙
次元
製品に関する質問
S
セマンティクス
システムは明確な意味を生み出しますか?
私は
意図
ユーザーが明示的に尋ねたことだけでなく、ユーザーが何を意味するのかを理解していますか?
G
接地
答えや行動が何に基づいているかを示していますか?
N
ナビゲーション
コンテキスト、状態、進行状況、次のステップを通じてユーザーの指示を維持できますか?
あ
代理店
プロファイル データの使用、設定の保存、状態の変更、ユーザーへの影響など、技術的なアクションを実行する前に質問しますか?
L
ロード
複雑な事実を要約し、文脈を明確に保ち、意味の一貫性を保ち、サポートされていないユーザーの意思決定の負担を回避することにより、精神的負荷が軽減されますか?
正規概念図
SIGNAL では製品に面した名前が使用されていますが、各次元は人間と AI のインタラクション、言語学、認知心理学、情報検索、エージェント研究から確立された概念に基づいています。
SIGNAL はこれらのコンセプトをゼロから発明したわけではありません。これらは、言語を通じて AI エクスペリエンスを構築するチームのための実用的なフレームワークにまとめられています。
SIGNAL では、エージェントが内部でどのように実装されるかについては説明されていません。
ここでは、SIGNAL コンポーネントがサンプル エージェント ループの上にどのように配置され、ユーザー/コンテキスト情報をプロアクティブな動作、目に見える価値、ユーザーが期待しているものについての正確なコミュニケーションに変えることができるかについて説明します。
以下に示すループは単なる例示です

たとえば、ドキュメントで example.com を使用するなどです。これは、SIGNAL の推奨事項、パターン、または必須のアーキテクチャではありません。
実装では、RAG、ツール、メモリ、ワークフロー、エージェント、MCP、データベース、またはプロンプトのみを使用できます。 UX の責任は変わりません。
これは内部アーキテクチャ図ではありません。再利用可能な割り当てモデルです。
SIGNAL は、実際の実装フローに関係なく、AI エクスペリエンスがユーザーに何を維持する必要があるかを示します。
AI の会話 UX は、表面的な品質、簡潔さ、または説明を備えた回答を生成することだけを目的とするものではありません。
優れた AI エクスペリエンスは、AI コンテキストをユーザーが伝えようとしているもの、およびユーザーが実際に必要としているものに近づけます。
理解することを事後的な責任と積極的な責任に分ける必要があります。
ユーザーが不明な点。
どのようなコンテキストがユーザーの回答に役立つか。
AIが今ユーザーに何を求めているか。
システムが能動的および受動的に実行していること。
以前の文脈から関連する可能性のあるもの。
ユーザーにコスト、リスク、または結果をもたらす可能性があるもの。
明示的に尋ねられずに AI がユーザーのためにできること。
AI が提供できるフォローアップやアクションの提案。
ユーザーが最後のメッセージと明確に結びつかないことを言った場合、システムはすぐに「どういう意味ですか?」と尋ねるべきではありません。まず、ユーザーが会話の前半で何かを言及しているか、現在の環境で表示されている何か、またはたった今起こった何かについて言及しているかどうかを確認する必要があります。
SIGNAL は、コミュニケーション UX の責任を製品の実際の動作にマッピングすることで、あらゆるプロンプト エンジニアリング、エージェント、ボット、アシスタント、ワークフロー、または AI 製品に適用できます。
次のループは、概念を説明するために使用される単なる例です。
フローチャート LR
U["ユーザー / コンテキスト"] --> A["理解する"]
A --> B[「理由」]
B --> C["行為"]
C --> D["検証"]

D --> E[「応答する」]
E --> U
読み込み中
この例では、SIGNAL は各ステージで何を保持する必要があるかを定義します。
これにより、この特定のループを必要とせずに、アーキテクチャ間で SIGNAL を再利用できるようになります。
システムが単なるプロンプトであるか、RAG アシスタントであるか、ツールを使用するエージェントであるか、MCP ワークフローであるか、データベースを利用したボットであるか、またはマルチエージェント システムであるかは関係ありません。
実装は変更される可能性があります。 UX の責任は変わりません。
理解する -> 意味と意図を保持する。
理由 -> 証拠、不確実性、状態を保存します。
行動 -> ユーザー制御を維持します。
検証 -> 正確性、範囲、リカバリを維持します。
応答 -> 明確さ、価値を維持し、労力を軽減します。
次の 5 つのステップで使用します。
製品の AI 動作をループにマッピングします。
各ステージに SIGNAL の責任を割り当てます。
失敗するステージのパターンを選択します。
プロンプト、ツール、取得、メモリ、ワークフロー、および応答のチェックを定義します。
失敗を製品の変更に変換します。表現の改善、検索の改善、状態の明確化、承認の安全性の向上、ユーザーの労力の軽減、またはアクションの境界の強化などです。
SIGNAL レビューの出力は具体的である必要があります。つまり、書き換えられた応答、より明確なアクション境界、より良いツール動作、改善された検索重複、評価チェック、および目に見える価値の受け取りです。
各 SIGNAL パターンは、繰り返し発生する AI UX 問題に対する再利用可能な回答です。
フローチャート TD
P["ユーザー通信の問題"] --> D["SIGNAL ディメンション"]
D --> C[「標準的な概念」]
C --> R[「参照に裏付けられた原則」]
R --> X[「再利用可能な UX パターン」]
X --> E[「評価基準」]
読み込み中
例:
ファイル
目的
docs/FRAMEWORK.md
完全なフレームワーク: 柱、基準、パターン、アンチパターン、成熟度モデル、テンプレート。
docs/RESEARCH_AND_BENCHMARKS.md
研究マップ、標準的な概念、隣接するフレームワーク、ベンチマークの比較、および参考資料。
docs/PATTERNS.md
実践的なパ

会話 UX のパターン。
docs/WHY_SIGNAL.md
言語優先の AI UX の背後にある理論。
docs/FOR_TEAMS.md
製品、設計、エンジニアリング、サポート、評価チームが SIGNAL を使用する方法。
テンプレート/conversation_ux_review.md
コピー可能なレビュー テンプレート。
シグナルではないもの
内部エージェントのアーキテクチャ。
製品リサーチの代替品。
あるアシスタントが他のアシスタントよりも普遍的に優れているという主張。
SIGNAL は、AI 製品エクスペリエンスのためのコミュニケーション UX レイヤーです。
AI はユーザーの意図を理解し、適切な範囲内で行動し、ユーザーの労力を軽減し、提供された価値を可視化しましたか?
SIGNAL は、LLM UX、人間と AI の対話、会話型 AI、エージェント UX、セマンティックの明瞭さ、認知負荷、そして主に LLM ベースのシステムにおけるユーザー エクスペリエンスのためのパターン フレームワークです。
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

SIGNAL is a pattern framework for LLM UX, human-AI interaction, conversational AI, agent UX, semantic clarity, cognitive load, and mostly user experience in LLM-based systems. - hi-mundo/SIGNAL

GitHub - hi-mundo/SIGNAL: SIGNAL is a pattern framework for LLM UX, human-AI interaction, conversational AI, agent UX, semantic clarity, cognitive load, and mostly user experience in LLM-based systems. · GitHub
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
hi-mundo
/
SIGNAL
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits docs docs profiles profiles templates templates CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTORS.md CONTRIBUTORS.md FRAMEWORK.md FRAMEWORK.md GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE README.md README.md RESEARCH_AND_BENCHMARKS.md RESEARCH_AND_BENCHMARKS.md View all files Repository files navigation
Semantic Interaction Guidelines for Natural AI Language
SIGNAL is a pattern framework for LLM UX, human-AI interaction, conversational AI, agent UX, semantic clarity, cognitive load, and, most importantly, user experience in LLM-based systems.
SIGNAL is under active development.
This framework was created after researching adjacent work in Human-AI Interaction, conversation design, cognitive accessibility, plain language, generative AI UX, AI governance, and LLM user research.
These fields contain strong ideas, guidelines, and empirical studies, but I could not find a focused, open, README-first pattern framework dedicated mostly to user experience in LLM-based systems, especially the communication layer: semantics, intent, grounding, navigation, agency, and cognitive load.
SIGNAL is not a claim of invention from zero.
It is a consolidation attempt: a practical pattern language for a space that appears fragmented across research papers, product guidelines, accessibility notes, governance frameworks, benchmark culture, prompt engineering, and chatbot design practices.
A signal is a gesture, token, or action used to transmit information, communicate a command, or serve as a warning.
When conversation is the interface, communication is the product experience.
Users do not want to write perfect prompts.
They communicate like people: short messages, vague references, metaphors, idioms, incomplete context, frustration, shortcuts, corrections, and indirect requests.
LLMs are optimized to generate plausible continuations of language. Product experiences require more than plausible continuation: they require visible intent, grounding, progress, agency, cost, and value.
SIGNAL designs the communication layer between messy human language and useful AI behavior.
Letter
Dimension
Product question
S
Semantics
Does the system create clear meaning?
I
Intent
Does it understand what the user means, not only what the user explicitly asked?
G
Grounding
Does it show what the answer or action is based on?
N
Navigation
Does it keep the user oriented through context, state, progress, and next steps?
A
Agency
Does it ask before taking technical actions, using profile data, storing preferences, changing state, or creating consequences for the user?
L
Load
Does it reduce mental load by summarizing complex facts, keeping context clear, preserving semantic consistency, and avoiding unsupported user decision burden?
Canonical Concept Map
SIGNAL uses product-facing names, but each dimension is grounded in established concepts from Human-AI Interaction, linguistics, cognitive psychology, information retrieval, and agent research.
SIGNAL is not inventing these concepts from scratch. It organizes them into a practical framework for teams building AI experiences through language.
SIGNAL does not explain how an agent is implemented internally.
It explains how SIGNAL components can sit on top of an example agent loop and turn user/context information into proactive behavior, visible value, and precise communication about what the user expects.
The loop shown below is only an illustrative example, like using example.com in documentation. It is not a SIGNAL recommendation, pattern, or required architecture.
The implementation can use RAG, tools, memory, workflows, agents, MCP, databases, or only prompting. The UX responsibilities stay the same.
This is not an internal architecture diagram. It is a reusable allocation model.
SIGNAL shows what an AI experience must preserve for the user, regardless of the actual implementation flow.
AI conversation UX is not only about producing an answer with cosmetic quality, concision, or explanation.
A good AI experience makes the AI context closer to what the user is trying to communicate and what the user actually needs.
It should divide understanding into reactive and proactive responsibilities.
what the user is uncertain about;
what context helps answer the user;
what the AI needs from the user now.
what the system is doing actively and passively;
what may become relevant from earlier context;
what may cost, risk, or create consequences for the user;
what AI can do for the user without being explicitly asked;
what follow-up or action suggestion the AI can offer.
If the user says something that does not clearly connect to the last message, the system should not immediately ask "what do you mean?". It should first check whether the user is referring to something earlier in the conversation, something visible in the current environment, or something that just happened.
SIGNAL can be applied to any prompt engineering, agent, bot, assistant, workflow, or AI product by mapping its communication UX responsibilities to the product's actual behavior.
The following loop is only an example used to explain the concept:
flowchart LR
U["User / Context"] --> A["Understand"]
A --> B["Reason"]
B --> C["Act"]
C --> D["Validate"]
D --> E["Respond"]
E --> U
Loading
In this example, SIGNAL defines what each stage must preserve.
This makes SIGNAL reusable across architectures without requiring this specific loop.
It does not matter whether the system is only a prompt, a RAG assistant, a tool-using agent, an MCP workflow, a database-backed bot, or a multi-agent system.
The implementation may change. The UX responsibilities remain the same:
Understand -> preserve meaning and intent.
Reason -> preserve evidence, uncertainty, and state.
Act -> preserve user control.
Validate -> preserve correctness, scope, and recovery.
Respond -> preserve clarity, value, and low effort.
Use it in five steps:
Map the product's AI behavior to the loop.
Assign SIGNAL responsibilities to each stage.
Choose patterns for the stages that are failing.
Define checks for prompts, tools, retrieval, memory, workflows, and responses.
Convert failures into product changes: better wording, better retrieval, clearer state, safer approval, lower user effort, or a stronger action boundary.
The output of a SIGNAL review should be concrete: rewritten responses, clearer action boundaries, better tool behavior, improved retrieval overlap, evaluation checks, and visible value receipts.
Each SIGNAL pattern is a reusable answer to a recurring AI UX problem.
flowchart TD
P["User communication problem"] --> D["SIGNAL dimension"]
D --> C["Canonical concept"]
C --> R["Reference-backed principle"]
R --> X["Reusable UX pattern"]
X --> E["Evaluation criterion"]
Loading
Example:
File
Purpose
docs/FRAMEWORK.md
Full framework: pillars, criteria, patterns, anti-patterns, maturity model, and templates.
docs/RESEARCH_AND_BENCHMARKS.md
Research map, canonical concepts, adjacent frameworks, benchmark comparison, and references.
docs/PATTERNS.md
Practical patterns for conversation UX.
docs/WHY_SIGNAL.md
The thesis behind language-first AI UX.
docs/FOR_TEAMS.md
How product, design, engineering, support, and eval teams can use SIGNAL.
templates/conversation_ux_review.md
Copyable review template.
What SIGNAL is not
an internal agent architecture;
a replacement for product research;
a claim that one assistant is universally better than another.
SIGNAL is a communication UX layer for AI product experiences.
Did the AI understand what the user meant, act within the right boundaries, reduce user effort, and make the delivered value visible?
SIGNAL is a pattern framework for LLM UX, human-AI interaction, conversational AI, agent UX, semantic clarity, cognitive load, and mostly user experience in LLM-based systems.
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
