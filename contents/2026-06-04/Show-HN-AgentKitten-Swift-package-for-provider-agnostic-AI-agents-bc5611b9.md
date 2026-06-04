---
source: "https://github.com/fbeeper/agentkitten"
hn_url: "https://news.ycombinator.com/item?id=48399345"
title: "Show HN: AgentKitten: Swift package for provider-agnostic AI agents"
article_title: "GitHub - fbeeper/agentkitten: Swift package for building provider-agnostic AI agents on Apple platforms. · GitHub"
author: "fbeeper"
captured_at: "2026-06-04T16:13:07Z"
capture_tool: "hn-digest"
hn_id: 48399345
score: 8
comments: 1
posted_at: "2026-06-04T14:37:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AgentKitten: Swift package for provider-agnostic AI agents

- HN: [48399345](https://news.ycombinator.com/item?id=48399345)
- Source: [github.com](https://github.com/fbeeper/agentkitten)
- Score: 8
- Comments: 1
- Posted: 2026-06-04T14:37:55Z

## Translation

タイトル: HN を表示: AgentKitten: プロバイダーに依存しない AI エージェント用の Swift パッケージ
記事のタイトル: GitHub - fbeeper/agentkitten: Apple プラットフォーム上でプロバイダーに依存しない AI エージェントを構築するための Swift パッケージ。 · GitHub
Description: Swift package for building provider-agnostic AI agents on Apple platforms. - fbeeper/エージェントキトゥン

記事本文:
GitHub - fbeeper/agentkitten: Apple プラットフォーム上でプロバイダーに依存しない AI エージェントを構築するための Swift パッケージ。 · GitHub
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
ピーピー
/
エージェントの子猫
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
55 コミット 55 コミット .claude .claude .github .github ドキュメント/リソース ドキュメント/リソース スクリプト スクリプト ソース ソース テスト テスト .gitignore .gitignore .spi.yml .spi.yml .swiftformat .swiftformat .swiftlint.yml .swiftlint.yml AGENTS.md AGENTS.md AUTHORS.txt AUTHORS.txt CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md CONTRIBUTORS.txt CONTRIBUTORS.txt ライセンス ライセンス Makefile Makefile 通知 通知 Package.resolved Package.resolved Package.swift Package.swift README.md README.md VISION.md VISION.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Apple プラットフォーム上でプロバイダーに依存しない AI エージェントを構築するための Swift パッケージ。
AgentKitten gives you standard building blocks for easily creating agents on Apple platforms without you having to reinvent their wheels.
さらに良いことに、あらゆる具体的なプロバイダー API から抽象化されています。
完全にカスタマイズ可能でプロバイダーに依存しない AI エージェントを簡単に構築します。
以下を簡単にサポートします:
ランタイムツールの権限とフック、
また、エージェントをデバッグ、テスト、評価するための簡単な監査可能なトレースも用意されています。
はい...最近では、毎回全体をバイブコーディングできますが、トークンを節約してプロジェクトの責任を軽減してみてはいかがでしょうか? 🙃
大規模な言語モデルに対してエージェント ハーネスが実行できることを示す、非常に安定した共有言語がすでに存在します。
説明すると、ツールは、ほとんどの LLM/エージェント プロバイダー API/SDK が完全にサポートしており、モデルが物事 (API 呼び出し、ファイル操作、または接続するカスタム ロジック) を実行できるようにする概念です。ただし、プロバイダーの特性があり、通常、実行時の実行許可とフックに関する標準はありません。
しかし、より広く共有されている概念がたくさんあるため、再実装する必要があるかもしれません。

o 強力なエージェントを動かします。たとえば、ツールの可用性の制御、ツール実行のモデル理論的根拠の取得、コンテキスト圧縮の実行、セッション状態の維持、さらにはモデルの出力に関する検証ループの定義などです。
それに加えて、プロバイダーの変更にかかるコストを回避したい、または回避する必要がある場合もあります。
機能の検討段階では、どのモデルがエージェントにとって最適なパフォーマンスを発揮するかまだわからない可能性があるためです。
プロダクション エージェントの場合、新しいモデルが頻繁に登場し、価格が変更され、ビジネス ポリシーが進化し、関係が変化するため、API も進化します...
AgentKitten は、選択した具体的なプロバイダーに関係なく、これらの基本エージェント機能を実装するための簡単な方法 (および可能な場合は段階的に開示) を提供することで支援することを目的としています。
そしてもちろん、LLM の確率的な性質を考慮して、AgentKitten はデバッグ、テスト、評価のための第一級のエージェント機能としてトレーサビリティも考慮しています。
InferenceProvider : エージェント モデルが存在する場所。 Claude API、基盤モデルを介したローカル Apple Intelligence、または独自の推論モデルが考えられます。ステートレスかステートフルか、オンデバイスかリモートかに関係なく。プロバイダーの固有性はすべて残しておきます。 1 つを選択すると、エージェントの実装に無料で交換できます。
エージェント : あなたのベースコントロールです。ツールを使用して構成し、基本動作と使用するプロバイダーを設定します。エージェントはこの構成を保持します。それだけです。
AgentSession : エージェントからセッションを開始します。各セッションは独立しています。同じエージェント、異なるセッション。複数ターンの会話を行い、お互いにステップを踏むことなく並列スレッドを実行します。デフォルトでは軽量で同時安全です。
トレース : すべてのセッションは、ツール呼び出し、結果、圧縮イベントなど、各ターンで起こったことの詳細な記録を保持します。デバッグ用の主要なリソース

NG、テスト、評価。
AgentTool : エージェントが動作できる対象を定義します。各ツールは、モデルが呼び出すことができる型付きのスキーマ記述された関数です (API 呼び出し、ファイル アクセス、アプリ統合など、さまざまな名前があります)。そして @Tool マクロを使用して、最小限の定型文で Swift 関数を接続するだけです。
ToolDefinition : エージェントが呼び出す可能性のあるツールと、それらを管理するポリシーおよびフックをバンドルします。
ToolExecutionPolicy : 選択したランタイム コンテキストに基づいて、実行前にツール呼び出しを承認、拒否、または一時停止します。
ToolHook : ツール入力を実行前に変換し、結果がモデルに到達する前に再形成またはサニタイズします。フックは宣言順に実行されます。
AgentBehavior および ToolBehavior : エージェントのシステム プロンプト、推論設定、圧縮ポリシーなど、エージェントが開始するデフォルトを設定します。ツールの可用性、ステップの予算、ツールのモデル ガイダンス。
TurnOverrides : ターンごとにこれらの動作のデフォルトをオーバーライドします。プロバイダーを交換したり、推論設定を調整したり、ツールの選択を制限したり、コンテキストを先頭に追加したりできます。また、エージェント設定に結合せずに、ツールの承認、フック、推論を通じてカスタム型指定された値 (ExecutionConfigurationKey サブスクリプト経由) をスレッドする場所でもあります。
Validator / JudgeValidator : アシスタントの応答の受け入れ基準を定義します。検証が失敗した場合、AgentKitten はフィードバックとともに自動的に再試行します。それが通過するか、裁判官モデルがそれを承認するまで。
以下は、アプリ用のシンプルかつ強力な自動圧縮検索エージェントの最小限のサンプルです。
AgentKitten をインポート
AgentKittenAnthropicInference をインポート
プロバイダー = InferenceProvider にします。人間性 ( )
let toolConfig = ToolDefinition ( tools : [
AnyAgentTool ( MySearchTool () ) 、
])
let 動作 = AgentBehavior (
systemPrompt : " あなたは検索アシスタントです。 " ,
デフォルトオートム

aticCompactionPolicy : 。有効 (
トリガー: 。コンテキストウィンドウのパーセント ( 0.5 )
)
)
let エージェント = エージェント (
動作: 動作、
プロバイダー : プロバイダー、
ツール定義:toolConfig、
)
セッション = エージェントとします。 makeSession ( )
また、チャット ボットからテキストを取得できます。
letturn = セッションを待ってみてください。 send (「ダウンタウン近くの公園を探してください。」)
順番に await イベントを試してください。イベント {
場合がございます。 textDelta (let text) = イベント。種類 {
print ( テキスト , ターミネータ : " " )
}
}
または、アプリを強化するための構造化された出力:
letturn = セッションを待ってみてください。 generate < [ PointOfInterest ] > (「ダウンタウン近くの公園を見つけてください。」)
順番に await イベントを試してください。イベント {
場合がございます。 result (let pois ) = イベント 。種類 {
DidLoad (ポイズ)
}
}
構築できるもの (およびいくつかのサンプル)
空には限界がありますが、2026 年には言わせてください...
アプリの検索バーにある古き良きあいまい一致やフィルターだけでは、ユーザーはもう十分ではないと感じているかもしれません。今日は、彼らが探しているものを説明してもらうことができます。また、エージェントは、適切な結果をもたらす検索パラメータとフィルターを組み合わせることができます。ピッカーをいじったり、万能のパワーユーザー機能を事前に知っておく必要はもうありません。
アプリには、アプリを特別なものにする多くの設定がありますが、美しく、丹念に整理された階層設定は、ユーザーにとって依然として時間のかかる切り替えの壁となっています。最近では、エージェントが、ユーザーが 6 つの異なる機能を一度に簡単にカスタマイズできるように支援してくれる可能性があります。
そして...なぜそうではないのでしょうか？あなたのアプリやゲームには、なぞなぞで話す風変わりなキャラクターがいて、夕食に適切な食材を与えることができれば城の鍵をくれるかもしれません。
おもちゃの例を swift run でご覧ください。 Playground Chicken
ギミックの多いサンプルはさておき、AgentKitten を使用して簡単に実行できるものをいくつか示します。
オンデバイスモデルでは継続時間が制限されています

外部ウィンドウ。ただし、それによって会話が制限される必要はありません (カスタム実装の話に突入する必要もありません)。 AgentKitten はセッションの増大に応じてコンテキストを自動的に圧縮できるため、ユーザーは手動でメモリを管理しなくても、無限のマルチターン エクスペリエンスを得ることができます。
早速実行してみてください Playground chat --compaction --show-usage
エージェントにツールを提供することは強力ですが、ユーザーは最新情報を常に把握しておきたい、または必要な場合があります。 AgentKitten を使用すると、各ツール呼び出しを実行前に簡単に承認または拒否できるようになり、各ツール呼び出しを使用する理由についてモデル独自の理論的根拠を持たせることもできます。また、人間参加が必要ない場合は、カスタムのサイレント非対話型ポリシーを定義することもできます。
早速実行してみてください Playground chat --tool-policy ask
一部のタスクでは、実行フェーズの前に思考フェーズを行うことでメリットが得られます。ターンごとのオーバーライドを使用すると、エージェントを計画モードに移行できます。ツールを制限し、ガイダンスを変更し、実行のためにツールを再度開きます。ハーネスのコーディングにおけるプラン モードの背後にある同じパターン。
即時実行 Playground プランモードで試してみる
プライバシーは多くのアプリにとって最重要事項です。 AgentKitten を使用すると、ユーザーとの編集コントラクトを簡単に構築できます。入力から機密データを削除し、推論を実行し、ツールを実行する直前にデータを再ハイドレートします。したがって、ユーザーデータは完全にモデルの外に保持されます。
swift run Playground pii で試してください
すべてのセッションでエージェントのライフサイクルの詳細なトレース (ツール呼び出しとその結果を含む) が生成されるため、毎回ハーネスを再構築することなくエージェントを評価できます。プロンプトを交換したり、スキャフォールディングを調整したり、プロバイダーを変更して結果を比較します。面倒な部分を使わずに素早く反復するのに最適です。
早速実行してみてください Playground chat --trace
ユーザーが作成したコンテンツは、定義上予測不可能です。イメージ機能を持つエージェント

ファイル モデルとその画像を提供する豊富なツールは、アプリが公開される前にアップロードを管理したり、代替テキストの説明を事前に入力したり、写真から構造化データを抽出したりするのに役立ちます。ユーザーが独自の画像をアップロードするあらゆるアプリ向けの実用的なパターン。
Playground tools --image-demo --provider anthropic を迅速に実行して試してください
機能の追加を夢見る前に、最初のプレリリース後、最も重要なマイルストーンは次のとおりです。
テストスイートとプレイグラウンドを改善します。
基本的なエージェント/ツールの動作とオーバーライド (API ブレーク) の間の対称性を追求してみてください。
さらに、AgentKitten をより良くするための可能性の長いリストがあります。 VISION.md でアイデアの最初の流出を見つけます。
macOS 15 以降、iOS 17 以降、tvOS 17 以降、watchOS 10 以降、visionOS 1 以降。スイフト6。
。パッケージ ( URL : " https://github.com/fbeeper/agentkitten "、 から: " 0.0.1 " )
次に、「AgentKitten」をターゲットの依存関係に追加します。
プロジェクトには、次の機能を実行できるプレイグラウンドが含まれています。
Playground --help をすぐに実行して、すべてのオプションを調べてください。
このフレームワークは私によって入念に設計、レビュー、手作業でテストされました。ただし、AI を書き込みツールとして使用して非常に意図的に構築されています。私はあらゆる面で AI コーディングのハーネスを徹底的に操縦しました

[切り捨てられた]

## Original Extract

Swift package for building provider-agnostic AI agents on Apple platforms. - fbeeper/agentkitten

GitHub - fbeeper/agentkitten: Swift package for building provider-agnostic AI agents on Apple platforms. · GitHub
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
fbeeper
/
agentkitten
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
55 Commits 55 Commits .claude .claude .github .github Docs/ Resources Docs/ Resources Scripts Scripts Sources Sources Tests Tests .gitignore .gitignore .spi.yml .spi.yml .swiftformat .swiftformat .swiftlint.yml .swiftlint.yml AGENTS.md AGENTS.md AUTHORS.txt AUTHORS.txt CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTORS.txt CONTRIBUTORS.txt LICENSE LICENSE Makefile Makefile NOTICE NOTICE Package.resolved Package.resolved Package.swift Package.swift README.md README.md VISION.md VISION.md View all files Repository files navigation
Swift package for building provider-agnostic AI agents on Apple platforms.
AgentKitten gives you standard building blocks for easily creating agents on Apple platforms without you having to reinvent their wheels.
And, even better, abstracted from any and all concrete provider APIs.
Effortlessly build fully customizable and provider-agnostic AI agents.
With straightforward support for:
Runtime Tool Permissions and Hooks,
And have simple auditable traces to debug, test, and evaluate your agents.
Yes... these days you could vibe-code the entire thing every time, but why not save tokens and reduce the responsibilities of your project? 🙃
There already is a quite stable shared language of what an agent harness can do for a large language model:
To illustrate, Tools are a concept that most LLM/agent provider APIs/SDKs fully support to allow a model to execute things (API calls, file operations, or any custom logic you wire up). However, there are provider specificities, and there typically isn't an standard for execution permissions and hooks at runtime.
But there are so many more widely shared concepts that you may have to reimplement to get a powerful agent going. For example, controlling tool availability , obtaining model rationale for tool execution, performing context compaction , keeping session state , or even defining validation loops around the model's output.
Besides that, you may also want/need to avoid the cost of changing providers:
During the exploratory phase of features, because you may not yet know which model performs best for your agent.
For your production agents, because new models come out often, pricing changes, business policies evolve, and relationships change, APIs evolve...
AgentKitten aims to help by offering simple ways (and with progressive disclosure where possible) to implement those base agent features independently of the concrete provider you choose.
And, of course, given the stochastic nature of LLMs, AgentKitten has also considered traceability as a first-class agent feature for debugging, testing, and evaluation.
InferenceProvider : Where the agent model lives. Could be the Claude API, local Apple Intelligence via Foundation Models, or your own inference model. Whether stateless or stateful, on-device or remote. You leave all provider specificities behind. Pick one, swap it out at no cost to your agent implementation.
Agent : Your base control. Configure it with tools, set up its base behavior, and the providers it may use. The agent keeps this configuration. That's all.
AgentSession : Start a session from your Agent. Each session is independent. Same agent, different sessions. Have multi-turn conversations and run parallel threads without stepping on each other. Lightweight and concurrent-safe by default.
Trace : Every session keeps a detailed record of what happened in each turn: tool calls, results, compaction events, and more. Your primary resource for debugging, testing, and evaluation.
AgentTool : Define what your agent can act on. Each tool is a typed, schema-described function the model can invoke (API calls, file access, app integrations, you name it...). And use the @Tool macro to just wire up a Swift function with minimal boilerplate.
ToolDefinition : Bundle the tools the agent may invoke together with the policy and hooks that govern them:
ToolExecutionPolicy : Approve, deny, or suspend any tool call before it runs, based on any runtime context you choose.
ToolHook : Transform tool inputs before execution, and reshape or sanitize results before they reach the model. Hooks run in declaration order.
AgentBehavior and ToolBehavior : Set the defaults your agent starts with, including the system prompt, inference settings, and compaction policy for the agent; tool availability, step budget, and model guidance for the tools.
TurnOverrides : Override any of those behavior defaults on a per-turn basis. Swap providers, adjust inference settings, restrict tool selection, or prepend context. Also the place to thread custom typed values (via ExecutionConfigurationKey subscript) through tool approval, hooks, and inference without coupling them to your agent setup.
Validator / JudgeValidator : Define acceptance criteria for the assistant's response. If validation fails, AgentKitten retries automatically with feedback. Until it passes, or a judge model approves it.
Here's a minimal sample of a simple but powerful auto-compacting search agent for your app:
import AgentKitten
import AgentKittenAnthropicInference
let provider = InferenceProvider . anthropic ( )
let toolConfig = ToolDefinition ( tools : [
AnyAgentTool ( MySearchTool ( ) ) ,
] )
let behavior = AgentBehavior (
systemPrompt : " You are a search assistant. " ,
defaultAutomaticCompactionPolicy : . enabled (
trigger : . percentOfContextWindow ( 0.5 )
)
)
let agent = Agent (
behavior : behavior ,
provider : provider ,
toolDefinition : toolConfig ,
)
let session = agent . makeSession ( )
And you can obtain text à la chat bot:
let turn = try await session . send ( " Find me parks near downtown. " )
for try await event in turn . events {
if case . textDelta ( let text ) = event . kind {
print ( text , terminator : " " )
}
}
Or structured output to power your app:
let turn = try await session . generate < [ PointOfInterest ] > ( " Find me parks near downtown. " )
for try await event in turn . events {
if case . result ( let pois ) = event . kind {
didLoad ( pois )
}
}
What You Could Build (and some Samples)
Sky is the limit, but let me say that in 2026...
Good old fuzzy matching and filters on your app's search bar may not feel enough to your users anymore. Today, you can let them describe what they are looking for. And your agent can put together the search parameters and filters that yield the right results. No more fiddling with pickers or having to know the almighty power-user features ahead of time.
Your app has many settings that make it special, but your beautiful and painstakingly organized hierarchical settings are still a time consuming wall of toggles for the user. These days, an agent could be there to help your users customize 6 different features on one go with much less fuss.
And... why not? Your app/game could have a quirky character that talks in riddles and will give you the key to the castle if you can give them the right ingredients for their dinner.
See a toy example at swift run Playground chicken
Gimmicky samples aside, here are some things that are simple to reach for with AgentKitten:
On-device models have limited context windows. But that doesn't have to limit the conversation (or send you into a custom implementation tangent). AgentKitten can compact context automatically as the session grows, so users get an infinite multi-turn experience without you managing memory manually.
Try it at swift run Playground chat --compaction --show-usage
Giving an agent tools is powerful, but users sometimes want/need to stay in the loop. With AgentKitten you can easily let them approve or deny each tool call before it runs, and will even have the model's own rationale for why it wants to use each tool call. And you can also define custom silent non-interactive policies when human-in-the-loop isn't needed.
Try it at swift run Playground chat --tool-policy ask
Some tasks benefit from a thinking phase before an acting phase. With per-turn overrides you can shift the agent into a planning mode. Restricting tools, changing guidance, then opening them back up for execution. The same pattern behind plan mode in coding harnesses.
Try it at swift run Playground plan-mode
Privacy is a first-class concern for many apps. AgentKitten makes it straightforward to build a redaction contract with your users: strip sensitive data from the input, run inference, then rehydrate it right before tools execute. Thus, keeping user data entirely outside the model.
Try it at swift run Playground pii
Because every session produces a detailed trace of Agent lifecycle (including tool calls and their results,) you can eval your agents without rebuilding the harness each time. Swap prompts, tweak scaffolding, or change providers and compare outcomes. Great for iterating quickly without the menial parts.
Try it at swift run Playground chat --trace
User-generated content is unpredictable by definition. An agent with an image-capable model and your rich tool providing it images can help your app moderate uploads before they go live, help pre-populate alt-text descriptions, or extract structured data from photos. Practical patterns for any app where users upload their own images.
Try it at swift run Playground tools --image-demo --provider anthropic
Before dreaming of adding more features, after its initial pre-release, the most important milestones reside on:
Improve the testing suite and Playground.
Try pursue symmetry between base Agent/Tools behavior and turn overrides (API breaking).
Past that, there is a long list of possibilities to make AgentKitten better. Find an initial spill of ideas in VISION.md .
macOS 15+, iOS 17+, tvOS 17+, watchOS 10+, visionOS 1+. Swift 6.
. package ( url : " https://github.com/fbeeper/agentkitten " , from : " 0.0.1 " )
Then add "AgentKitten" to your target's dependencies.
The project includes playground of its functionality you can exercise:
Explore all its options with swift run Playground --help .
This framework has been painstakingly designed, reviewed, and hand tested by me. But very deliberately built using AI as the writing tool. I very thoroughly steered AI coding harnesses in all thing

[truncated]
