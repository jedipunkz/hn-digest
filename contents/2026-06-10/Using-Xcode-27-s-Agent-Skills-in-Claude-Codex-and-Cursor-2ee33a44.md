---
source: "https://www.avanderlee.com/ai-development/using-xcode-27s-agent-skills-in-claude-codex-and-cursor/"
hn_url: "https://news.ycombinator.com/item?id=48475011"
title: "Using Xcode 27's Agent Skills in Claude, Codex, and Cursor"
article_title: "Using Xcode 27's Agent Skills in Claude, Codex, and Cursor - SwiftLee"
author: "ianhxu"
captured_at: "2026-06-10T13:19:26Z"
capture_tool: "hn-digest"
hn_id: 48475011
score: 1
comments: 0
posted_at: "2026-06-10T12:00:25Z"
tags:
  - hacker-news
  - translated
---

# Using Xcode 27's Agent Skills in Claude, Codex, and Cursor

- HN: [48475011](https://news.ycombinator.com/item?id=48475011)
- Source: [www.avanderlee.com](https://www.avanderlee.com/ai-development/using-xcode-27s-agent-skills-in-claude-codex-and-cursor/)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T12:00:25Z

## Translation

タイトル: クロード、コーデックス、およびカーソルでの Xcode 27 のエージェント スキルの使用
記事のタイトル: クロード、コーデックス、カーソルで Xcode 27 のエージェント スキルを使用する - SwiftLee
説明: Xcode 27 で公式エージェント スキルがどのように導入され、革新的な機能でアプリ開発エクスペリエンスが向上するかをご覧ください。

記事本文:
無料の mDevCamp チケットを手に入れましょう。10 秒以内に参加してください。プレゼントへ→
プレゼント: 無料の mDevCamp チケットプレゼント。
-->
コース
Swift 同時実行プレイブック
無料の 5 日間の電子メール コースでは、iOS 開発者が App Store の拒否につながる async/await で犯す 5 つの最大の間違いを明らかにします (Swift を何年も書いている場合でも) 移行プロジェクトには数日ではなく数か月かかります
5レッスン |完全無料 | iOS開発者向け
RocketSim: 必須の開発者ツール
Appleが推奨するとおり
SwiftLee > AI 開発 > Claude、Codex、および Cursor での Xcode 27 のエージェント スキルの使用
サードパーティ IDE で使用するエージェント スキルのエクスポート
Xcode 27 に同梱されるエージェント スキル
その他の人気のあるエージェント スキルを検討してください
Claude、Codex、および Cursor での Xcode 27 のエージェント スキルの使用
Apple は WWDC'26 中に Xcode 27 を発表し、公式エージェントのスキルを含むエージェント開発の多数の改善を導入しました。私の「適切なエージェント スキルを選択するための 9 ステップのフレームワーク」から学んだように、信頼できる情報源からスキルを選択することが重要です。もちろん、Apple も間違いなく 1 社として数えられます。
公式エージェント スキルは、新しい Xcode を使用する場合でも、Claude、Codex、Cursor などの既存の IDE を使用する場合でも、アプリの構築方法に影響を与えます。以前、AI を使用してより良いビューを構築するために SwiftUI エージェント スキルをリリースしましたが、今日正式な代替手段を確認できて、正直にとてもうれしく思っています。 Cursor、Claude、ChatGPT、または Codex を使用して RevenueCat、Sentry、Linear、GitHub、Supabase、または Astro を使用するすべての AI アプリで MCP の構成を停止しますか? MCP Beast は、Mac 上の iOS スタックに 1 つのプライベート MCP ゲートウェイを提供します。各サーバーに 1 回接続し、macOS キーチェーンに認証情報を保持し、すべての AI クライアントが同じ信頼できるエンドポイントを使用できるようにします。 MCP Beastを無料でダウンロードしてください。
サードパーティ IDE で使用するエージェント スキルのエクスポート
その間

エージェント スキルは Xcode 27 ですでに利用可能であるため、グローバルに利用できるようにしたい場合があります。 Codex、Claude、Cursor、またはエージェント スキルをサポートするその他の IDE を使用するかどうか。
Xcode に対する多くの改善の一環として、チームは、 Agent と呼ばれる新しい CLI ツールも導入しました。
新しいエージェント スキル エクスポート コマンドを備えた Xcode のエージェント CLI。
CLI は主に Xcode に付属のモデル コンテキスト プロトコルに焦点を当てていますが、Xcode に付属のエージェント スキルもエクスポートできるようになりました。
Xcode 27 エージェント スキルをエクスポートして、Claude、Codex、Cursor、またはその他の IDE で使用する方法は次のとおりです。 Terminal.app を開きます。
コマンド xcrun を実行します。エージェント スキルのエクスポート ~/.agents/skills 一般的なエージェント IDE の多くは、エージェント スキルを検出するための ~/.agents/ フォルダーをサポートしています。
Agent IDE を開き、必要に応じてスキルを更新します。スキルが貯まらない場合は、IDE を再起動してください。
エージェント スキルを正常にエクスポートしたら、選択した IDE 内でそれらを使用できるようになります。たとえば、Cursor 内のオートコンプリートは次のとおりです。
Xcode の 27 のエージェント スキルは、カーソル プロンプト フィールド内で利用できます。
コード内で AI エージェントを使用する方法を推測するのはやめましょう
AI エージェントを操作するための明確でツールに依存しないシステム (コンテキスト、指示、検証ループをカバー) を学習することで、使用するツールやモデルに関係なく、技術的負債を蓄積することなくより迅速に出荷できるようになります。
ウェイティングリストに参加する
Xcode 27 に同梱されるエージェント スキル
ターミナルでコマンドを実行すると、Xcode 27 内で現在どのエージェント スキルが使用できるかがわかります。重要な注意事項: これらのスキルのすべてが Xcode の外で役立つわけではありません (これについては後で詳しく説明します)。
avanderlee@AJs-MacBook-Pro MacOS % xcrun エージェント スキル エクスポート ~/.agents/skills
7 つのスキルを /Users/avanderlee/.agents/skills にエクスポートしました
✓ う

ikit-アプリの最新化
✓ デバイス間の相互作用
✓ swiftui-whats-new-27
✓ swiftui スペシャリスト
✓ テストモダナイザー
✓ C-bounds-safety
✓ 監査-xcode-セキュリティ設定
名前だけを読むだけで興奮してしまうかもしれませんが、それぞれが何をするのかについてもう少し詳しく見てみましょう。
uikit-app-モダナイゼーション
UIScreen.main や InterfaceOrientation などの従来の共有状態 API を置き換えることにより、マルチウィンドウ環境向けに UIKit アプリを最新化します。 Swift と Objective-C の両方のコードベースにわたって、シーンのライフサイクル移行と非対称セーフ エリア インセットを処理します。
デバイスの相互作用
スクリーンショット、UI 階層の検査、合成されたタッチ インタラクションを使用して、実際のデバイスまたはシミュレーター上でアプリの動作を検証します。サブエージェントとして実行して、機能が実際に動作することを確認し、視覚的または機能的なバグを捕捉します。
swiftui-新着情報-27
2027 OS リリース (iOS、macOS、watchOS、tvOS、visionOS 27) 全体で導入された新しい SwiftUI API、動作、非推奨について文書化します。 @State マクロの移行、ドラッグして並べ替え、新しいツールバー API、スワイプ アクション、ドキュメント ベースのアプリ、およびその他のソースと互換性のない変更について説明します。
swiftui-スペシャリスト
SwiftUI コードの作成、レビュー、リファクタリングに関する Apple の信頼できるベスト プラクティスと慣用的なパターンを提供します。ビューの構造、データ フロー、環境の使用法、修飾子、ローカリゼーション、アニメーション、ForEach ID、およびソフトで非推奨となった API について説明します。
テストモダナイザー
既存の XCTest スイートを Swift Testing に移行し、古い Swift Testing コードを再構築して新しい機能を採用します。アサーションを期待値にマッピングし、setUp/tearDown を init/deinit に変換し、特性、確認、パラメーター化されたテストを導入します。
C-bounds-safety
境界外のメモリ アクセスを防止するために、C -fbounds-safety 言語拡張機能の導入をガイドします。言語をカバーします

ゲージ モデル、ポインタ アノテーション、コンパイラ ビルド設定、および境界違反のランタイム デバッグ。
監査-xcode-セキュリティ設定
Xcode プロジェクトのセキュリティ体制を監査し、ビルド設定の強化、コンパイラー警告、静的アナライザー チェッカーを段階的に有効にします。強化されたセキュリティの権利と、ポインター認証、型付きアロケーター、スタックのゼロ初期化などの機能を構成します。
今後の Xcode の更新では、新しいスキルや既存のスキルの更新が提供されると予想されるため、コマンドを頻繁に実行することをお勧めします。
これらのスキルはすべて、Xcode 内から実行せずに使用できますか?
これらのスキルはすべて、Xcode 内から実行せずに使用できますか?
ほとんどの場合、そうですが、すべてではありません。スキルは実際には単なる指示のセットであり、指示はエージェントがそれに基づいて動作するツールを持っている場合にのみ役に立ちます。 swiftui-specialist 、 swiftui-whats-new-27 、 test-modernizer などの知識主導のスキルは、ソース ファイルの読み取りと編集のみを行うため、どこにでも移動できます。私は Xcode をまったく実行せずに、Cursor でそれらを喜んで使用します。
例外は、コードの外部の何かと通信する必要があるスキルです。 Audit-xcode-security-settings は、Xcode 固有のツールがビルド設定の読み取りと書き込みを行うことを想定しているため、他の場所での手動の .pbxproj 編集に低下します。そして、デバイス インタラクションは最も明確なケースです。これは、Xcode 27 エージェントのみが提供するツールを通じて実際のデバイスを駆動するため、Xcode の外部には呼び出すものがまったくありません。私の経験則: スキルがソース ファイルのみを操作する場合、そのスキルはどこでも機能します。プロジェクト設定や実行中のデバイスが必要になった瞬間に、Xcode 内から実行したくなるでしょう。
エージェントと一緒にシミュレーターを制御したいですか?シミュレーターを使用したエージェント開発を確認してください。
その他の人気のあるエージェント スキルを検討してください
アジャンを初めて使用する場合

t スキル。使用するエージェント スキルをさらに知りたいと思われるかもしれません。私のスキルを見つける方法は、私の個人的なエージェント スキルの概要のように、skills.sh を使用することです。特定のエージェントのスキルに関する背景ストーリーを詳しく知りたい場合は、次の関連記事を参照してください。
Core Data Agent Skill: オープンソースで利用可能になりました
6 つのエージェント スキルを使用した Xcode ビルドの最適化
SwiftUI エージェント スキル: AI を使用してより良いビューを構築する
Swift Testing Agent Skill: AI を使用して高品質のテストを作成する
または、 AGENTS.md を置き換えるのが初めての場合は、「エージェント スキルの説明: AGENTS.md を再利用可能な AI ナレッジで置き換える」を参照してください。これらは全体として、Xcode 27 に付属するエージェント スキルを補完します。
Cursor、Claude、ChatGPT、または Codex を使用して RevenueCat、Sentry、Linear、GitHub、Supabase、または Astro を使用するすべての AI アプリで MCP の構成を停止しますか? MCP Beast は、Mac 上の iOS スタックに 1 つのプライベート MCP ゲートウェイを提供します。各サーバーに 1 回接続し、macOS キーチェーンに認証情報を保持し、すべての AI クライアントが同じ信頼できるエンドポイントを使用できるようにします。 MCP Beastを無料でダウンロードしてください。結論
Apple が公式エージェント スキルを提供することは、コミュニティにとってありがたいことです。 Apple のスキルはコンパクトで、必要なものすべてをカバーしていないか、更新が遅いために古くなっている可能性があるため、GitHub 上にはオープンソースのサードパーティ エージェント スキルを活用する余地がまだあります。エージェントのコード結果を常に監視し、エージェントが使用するスキルにギャップがあるかどうかを判断することが重要です。学習内容をオープンソースのスキルにまとめるのは、Xcode リリースからの不確実なアップデートを待つよりも簡単かつ迅速になります。
エージェント スキルの選択は AI の基礎の一部であり、私の専用コース「開発者のためのエージェント コーディングの基礎」にぜひお越しください。
会いましょう?
SwiftLee の注目の求人
M7 Health のリード ソフトウェア エンジニア • $105,000 - $185,000
テスト

Ordergroove のエンジニアリング生産性、オートメーション リード • $140,000 - $170,000
2010 年から iOS 開発者であり、元 WeTransfer のスタッフ iOS エンジニアであり、現在は SwiftLee のフルタイムのインディー開発者兼創設者です。 Swift、iOS、Xcode に関連する新しいブログ投稿を毎週書いています。定期的な講演者およびワークショップの主催者。
あなたの情熱を利益に変えるための私の実証済みの手順を学びましょう。
6 つのエージェント スキルを使用した Xcode ビルドの最適化
Xcode のシミュレーターとエージェントを使用したネットワーク リクエストの最適化
Xcode Instruments Time Profiler: AI によるパフォーマンスの向上
AI は問題なく見える Swift コードを書きました…Xcode Instruments が公開するまでは
YouTube ビデオ VVViSXV4QTZNLVE5QXRLZ1FKETN6TkhnLmhERk95LXluSjZJ
AI は問題なく見える Swift コードを書きました…Xcode Instruments が公開するまでは
AI エージェントが iOS シミュレーターを制御できるようになりました
AI エージェントを使用してネットワーク リクエストを最適化する

## Original Extract

Discover how Xcode 27 introduces official agent skills, improving your app development experience with innovative features.

Free mDevCamp ticket up for grabs, enter in 10 seconds. Go to the giveaway →
Giveaway: Free mDevCamp ticket giveaway.
-->
Courses
The Swift Concurrency Playbook
A FREE 5-day email course revealing the 5 biggest mistakes iOS developers make with async/await that lead to App Store rejections And migration projects taking months instead of days (even if you've been writing Swift for years)
5 lessons | 100% free | For iOS developers
RocketSim: An Essential Developer Tool
as recommended by Apple
SwiftLee > AI Development > Using Xcode 27’s Agent Skills in Claude, Codex, and Cursor
Exporting Agent Skills to use with 3rd party IDEs
The Agent Skills that ship with Xcode 27
Other popular Agent Skills to explore
Using Xcode 27’s Agent Skills in Claude, Codex, and Cursor
Apple launched Xcode 27 during WWDC’26, introducing a bunch of agentic development improvements, including official agent skills. As you’ve learned from my 9-Step Framework for Choosing the Right Agent Skill , it’s important to pick skills from authoritative sources. Apple absolutely counts like one, of course!
Official Agent Skills will impact how we build apps, whether with the new Xcode or with existing IDEs like Claude, Codex, or Cursor. I’ve earlier launched a SwiftUI Agent Skill to build better views with AI , but I’m honestly super happy to see an official alternative today. Stop configuring MCPs in every AI app Using RevenueCat, Sentry, Linear, GitHub, Supabase, or Astro with Cursor, Claude, ChatGPT, or Codex? MCP Beast gives your iOS stack one private MCP gateway on your Mac. Connect each server once, keep credentials in macOS Keychain, and let every AI client use the same trusted endpoint. Download MCP Beast for free .
Exporting Agent Skills to use with 3rd party IDEs
While the Agent Skills are already available in Xcode 27, you might want them to be available globally. Whether you use Codex, Claude, Cursor, or any other IDE that supports Agent Skills.
As part of the many improvements to Xcode, the team also introduced a new CLI tool called agent :
Xcode’s agent CLI that comes with a new Agent Skills export command.
While the CLI is primarily focused on the Model Context Protocol that ships with Xcode, it now also lets you export the Agent Skills that come with Xcode.
Here’s how you can export the Xcode 27 Agent Skills to use in Claude, Codex, Cursor, or any other IDEs: Open the Terminal.app
Execute the command xcrun agent skills export ~/.agents/skills Many popular Agent IDEs support the ~/.agents/ folder to discover Agent Skills.
Open your Agent IDE, refresh the skills if needed. If the skills don’t shop up, relaunch the IDE.
After you’ve successfully exported the agent skills, you’re ready to use them inside any IDE of your choice. For example, here’s the autocomplete inside Cursor:
Xcode’s 27 Agent Skills are available inside the Cursor prompt field.
Stop Guessing How to Use AI Agents in Your Code
Learn a clear, tool-agnostic system for working with AI agents — covering context, instructions, and validation loops — so you can ship faster without accumulating tech debt , no matter which tools or models you use.
Join the Waitlist
The Agent Skills that ship with Xcode 27
After you’ve executed the command in the terminal, you’ll notice which Agent Skills are currently available inside Xcode 27. Important note: not all of these skills are useful outside of Xcode (more on this later).
avanderlee@AJs-MacBook-Pro MacOS % xcrun agent skills export ~/.agents/skills
Exported 7 skills to /Users/avanderlee/.agents/skills
✓ uikit-app-modernization
✓ device-interaction
✓ swiftui-whats-new-27
✓ swiftui-specialist
✓ test-modernizer
✓ c-bounds-safety
✓ audit-xcode-security-settings
You might get enthusiastic by reading just the names, but let’s also dive into a bit more details on what each does.
uikit-app-modernization
Modernizes UIKit apps for multi-window environments by replacing legacy shared-state APIs like UIScreen.main and interfaceOrientation . Handles scene lifecycle migration and asymmetric safe area insets across both Swift and Objective-C codebases.
device-interaction
Verifies your app’s behavior on a real device or simulator using screenshots, UI hierarchy inspection, and synthesized touch interactions. Runs as a subagent to confirm that features actually work and to catch visual or functional bugs.
swiftui-whats-new-27
Documents the new SwiftUI APIs, behaviors, and deprecations introduced across the 2027 OS releases (iOS, macOS, watchOS, tvOS, visionOS 27). Covers the @State macro migration, drag-to-reorder, new toolbar APIs, swipe actions, document-based apps, and other source-incompatible changes.
swiftui-specialist
Provides Apple’s authoritative best practices and idiomatic patterns for writing, reviewing, and refactoring SwiftUI code. Covers view structure, data flow, environment usage, modifiers, localization, animations, ForEach identity, and soft-deprecated APIs.
test-modernizer
Migrates existing XCTest suites to Swift Testing and restructures older Swift Testing code to adopt newer features. Maps assertions to expectations, converts setUp/tearDown to init/deinit, and introduces traits, confirmations, and parameterized tests.
c-bounds-safety
Guides adoption of the C -fbounds-safety language extension to prevent out-of-bounds memory access. Covers the language model, pointer annotations, compiler build settings, and runtime debugging of bounds violations.
audit-xcode-security-settings
Audits an Xcode project’s security posture and progressively enables hardening build settings, compiler warnings, and static analyzer checkers. Configures Enhanced Security entitlements and features like pointer authentication, typed allocators, and stack zero-initialization.
I expect future Xcode updates to ship new skills and updates to existing ones, so it’s recommended to run the command frequently.
Can I use all these skills without running them from inside Xcode?
Can I use all these skills without running them from inside Xcode?
Mostly, yes, but not all of them. A skill is really just a set of instructions, and instructions are only useful if the agent has the tools to act on them. The knowledge-driven skills like swiftui-specialist , swiftui-whats-new-27 , and test-modernizer travel anywhere, since they only read and edit your source files. I happily use those in Cursor without Xcode running at all.
The exceptions are the skills that need to talk to something outside your code. audit-xcode-security-settings expects Xcode-specific tools to read and write build settings, so it degrades to manual .pbxproj edits elsewhere. And device-interaction is the clearest case: it drives a real device through tools only the Xcode 27 agent provides, so outside of Xcode there is simply nothing to call. My rule of thumb: if a skill only touches source files, it works everywhere; the moment it needs your project configuration or a running device, you’ll want to run it from inside Xcode.
Looking to control the Simulator with your agent? Check out Agentic Development with the Simulator .
Other popular Agent Skills to explore
If you’re new to Agent Skills, you might now be interested in finding out more Agent Skills to use. My way of finding skills is to use skills.sh, like this overview from my personal Agent Skills . If you want more background story on specific agent skills, here are a few related articles:
Core Data Agent Skill: Now available open-source
Xcode Build Optimization using 6 Agent Skills
SwiftUI Agent Skill: Build better views with AI
Swift Testing Agent Skill: Write high quality tests with AI
Or, if you’re new to replacing your AGENTS.md , you’ll enjoy this one: Agent Skills explained: Replacing AGENTS.md with reusable AI knowledge . Altogether, they complement the Agent Skills that come with Xcode 27.
Stop configuring MCPs in every AI app Using RevenueCat, Sentry, Linear, GitHub, Supabase, or Astro with Cursor, Claude, ChatGPT, or Codex? MCP Beast gives your iOS stack one private MCP gateway on your Mac. Connect each server once, keep credentials in macOS Keychain, and let every AI client use the same trusted endpoint. Download MCP Beast for free . Conclusion
Apple providing official Agent Skills is a blessing for the community. There’s still a place for 3rd-party agent skills open-sourced on GitHub, as Apple’s skills might be compact, not cover everything you need, or be outdated due to slower updates. It’s important to keep an eye on agents’ code results and determine whether there are gaps in the skills they use. Compiling learnings into open-source skills will be easier and faster than waiting for uncertain updates from the Xcode releases.
Picking Agent Skills is part of AI Fundamentals, and I’d love to welcome you to my dedicated course: Agentic coding fundamentals for developers .
See you there?
Featured SwiftLee Jobs
Lead Software Engineer @ M7 Health • $105K - $185K
Test Automation Lead, Engineering Productivity @ Ordergroove • $140K - $170K
iOS Developer since 2010, former Staff iOS Engineer at WeTransfer and currently full-time Indie Developer & Founder at SwiftLee. Writing a new blog post every week related to Swift, iOS and Xcode. Regular speaker and workshop host.
Learn my proven steps to transform your passion into profit.
Xcode Build Optimization using 6 Agent Skills
Network Requests Optimization using Xcode’s Simulator & Agents
Xcode Instruments Time Profiler: Improve performance with AI
AI Wrote Swift Code That Looked Fine… Until Xcode Instruments Exposed It
YouTube Video VVViSXV4QTZNLVE5QXRLZ1FKeTN6TkhnLmhERk95LXluSjZJ
AI Wrote Swift Code That Looked Fine… Until Xcode Instruments Exposed It
AI Agents Can Now Control the iOS Simulator
Using AI Agents to Optimize Network Requests
