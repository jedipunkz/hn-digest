---
source: "https://github.com/TheArcForge/Hades"
hn_url: "https://news.ycombinator.com/item?id=48475362"
title: "Show HN: Hades – Unity context graph for Claude Code"
article_title: "GitHub - TheArcForge/Hades: Unity-aware AI infrastructure for Claude Code — a knowledge graph + 88 MCP tools that let your AI agent know your project, not just grep its files. · GitHub"
author: "mikekuharuk"
captured_at: "2026-06-10T13:18:41Z"
capture_tool: "hn-digest"
hn_id: 48475362
score: 1
comments: 0
posted_at: "2026-06-10T12:34:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Hades – Unity context graph for Claude Code

- HN: [48475362](https://news.ycombinator.com/item?id=48475362)
- Source: [github.com](https://github.com/TheArcForge/Hades)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T12:34:01Z

## Translation

タイトル: HN を表示: Hades – クロード コードの Unity コンテキスト グラフ
記事のタイトル: GitHub - TheArcForge/Hades: Claude Code 用の Unity 対応 AI インフラストラクチャ — ファイルを grep するだけでなく、AI エージェントにプロジェクトを認識させるナレッジ グラフ + 88 MCP ツール。 · GitHub
説明: Claude Code 用の Unity 対応 AI インフラストラクチャ — ファイルを grep するだけでなく、AI エージェントにプロジェクトを知らせるナレッジ グラフ + 88 MCP ツール。 - アークフォージ/ハデス
HN テキスト: Unity と Claude Code をペアで動作させるための新しいオープンソース プロジェクトを紹介しました。簡単に言うと、Claude は通常、複雑なシーン、ネストされたプレハブ、さらにはそれぞれに数十のプロパティを持つ数十のコンポーネントが接続されたゲームオブジェクトなど、Unity プロジェクトをナビゲートするのに苦労しています。単純な grep-and-guess は、効率の点で非常に弱く、そして最も重要なことに、信頼性の点で非常に弱いです。クロードはゲーム内のエンティティ間の関係を見逃しています、それが苦い真実です。実際の比較に興味がある場合は、具体的な例を示します。エージェントに「EnemyAI を変更するとどのプレハブとシーンが壊れますか?」と尋ねました。 Stock Claude コードは YAML の約 200,000 トークンを読み取り、1 つのプレハブを見つけ、3 つのバリアントを見逃し、継承を壊す変更を提案しました。 Hades を使用すると、7 回のツール呼び出しで正しいマップ (4 つのプレハブ、3 つのシーン) が返され、27% 安くなりました。再現と完全に並べて表示:
https://github.com/TheArcForge/Hades/blob/main/Documentation... 私は仕事の一環として毎日 Unity と Claude を使用するため、解決策を考え出しました。 Hades はプロジェクトのエンティティからグラフを構築し、Claude は MCP ツールを介してグラフをクエリできます。これは簡単なバージョンです。詳細が必要な場合は、Medium に全文を書きました。
https://medium.com/@mike.kuharuk/unity-context-graph-for-cla... これまでのところ、macOS と Unity 6 でのみ実証テストを行っています。その範囲

一歩ずつ成長していきます。また、v1 は主に Claude Code を中心に構築されているため、Cursor およびその他すべての AI エージェントのサポートは発売後に追加されます。技術的なフィードバックや提案をいただければ幸いです。

記事本文:
GitHub - TheArcForge/Hades: Claude Code 用の Unity 対応 AI インフラストラクチャ — ファイルを grep するだけでなく、AI エージェントにプロジェクトを認識させるナレッジ グラフ + 88 MCP ツール。 · GitHub
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
ザアークフォージ
/
ハデス
公共
通知
あなたはきっとそうでしょう

サインインして通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
38 コミット 38 コミット .claude-plugin .claude-plugin .github .github Bridge~ Bridge~ Dashboard~ Dashboard~ Documentation Documentation Editor Editor Fixtures~/ TestProject/ Assets Fixtures~/ TestProject/ Assets Scanner~ Scanner~ Tests Tests ThirdParty ThirdParty コマンド コマンド スクリプト スクリプト スキル スキル .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CHANGELOG.md.meta CHANGELOG.md.meta CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md.meta CODE_OF_CONDUCT.md.meta COTRIBUTING.md COTRIBUTING.md COTRIBUTING.md.meta CONTRIBUTING.md.meta Documentation.meta Documentation.meta Editor.meta Editor.meta LICENSE LICENSE LICENSE.meta LICENSE.meta LIMITATIONS.md LIMITATIONS.md LIMITATIONS.md.meta LIMITATIONS.md.meta README.md README.md README.md.meta README.md.meta SECURITY.md SECURITY.md SECURITY.md.meta SECURITY.md.meta Tests.meta Tests.meta ThirdParty.meta ThirdParty.meta コマンド.meta コマンド.meta package.json package.json package.json.meta package.json.meta scripts.meta scripts.meta skill.meta skill.meta すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Unity プロジェクトのアンダーワールドでは、Hades から何も隠されていません。これで、AI エージェントから何も隠されることはなくなりました。
Hades は、Claude Code 用の Unity 対応 AI インフラストラクチャです。 Unity プロジェクト全体 (すべてのシーン、プレハブ、スクリプト、アセット、依存関係) のクエリ可能なナレッジ グラフを構築するため、AI エージェントはプロジェクトの構造を推測するのではなく認識できます。箱から出してすぐに使えるのは、88 個の MCP ツール、22 個のスキル、6 個のコマンドです。すべてがローカルで実行され、すべてがバージョン管理可能です。
1 つのプロンプト

— 「EnemyAI を変更するとどのプレハブとシーンが壊れますか?」
左：ストッククロードコード。右：ハデスと。
→ 完全な内訳を並べて表示
Stock Claude コードは YAML の約 200,000 トークンを読み取り、1 つのプレハブを見つけ、3 つをミスします
バリアントを破り、それを壊すコードを追加するように指示します。ハデスの答え
正しく — 4 つのプレハブ、3 つのシーン — 7 つのツール呼び出しでコストが 27% 削減されます。
そのギャップがプロジェクト全体です。
ほとんどの AI ツールは検索と予測を行います。関連すると思われるテキストを grep し、残りをモデルに推測させます。答えは確率的なものであり、目に見えない形で間違っていることもよくあります。
Hades はエージェントに情報を提供し、分析します。 「何が PlayerController を参照しているのか」を尋ねると、散在するスニペットからの推測ではなく、グラフから構造的な事実を読み取ります。依存関係分析は実際のエッジを追跡します。同じ質問を2回すると、同じ答えが得られます。 1 つのグラフ クエリで数十のファイル読み取りが置き換えられます。また、エージェントがプロジェクトについて 2 度説明することはありません。
Hades が AI エージェントに与えるもの
レイヤー
何をするのか
グラフ
Unity プロジェクトのセマンティック ナレッジ グラフ (シーン、プレハブ、スクリプト、アセット、およびそれらの依存関係)。エージェントは、プロジェクトのファイルだけでなく、プロジェクトの構造も確認します。
カロン
完全な可観測性 — すべてのツール呼び出し、グラフ クエリ、メモリ操作がトレースされます。ローカル ダッシュボード経由で検査します (Hades > Unity で Charon ダッシュボードを開く)。
アスフォデル
バージョン管理されたマークダウンの永続的なプロジェクト メモリ ( .arcforge/memory/ )。決定事項、パターン、慣例を一度把握します。エージェントは、セッションごとにコンテキストに応じたアドバイスを得るためにそれらを読みます。
22 スキル
アーキテクチャの決定、ワークフローのガイダンス、ネットワーキング、オーディオ、UI、シェーダー、ECS、テストなどの分野の専門知識。
88 MCP ツール
21 個のグラフ/カロン/メモリ ツール + 67 個のエディター アクション ツール (シーン、コンポーネント、プレハブ、マテリアル、アニメーション、アセット)。
6 コマンド

s
/hades:status 、 /hades:rebuild-graph 、 /hades:show-traces 、 /hades:validate-memory 、 /hades:show-proposals 、 /hades:export-traces
ピースがどのように組み合わされるか
フローチャート TD
エージェント["AIエージェント<br/>(クロード・コード)"]
エージェント <-->|"88 MCP ツール (ブリッジ経由)"|ハデス
サブグラフ Hades["Hades — Unity エディター内"]
グラフ["グラフ<br/>プロジェクトナレッジグラフ"]
アスフォデル[「アスフォデル<br/>永続的な記憶」]
Charon[「カロンの<br/>可観測性」]
終わり
グラフ -->|"インデックス"| Unity["Unity プロジェクト<br/>シーン · プレハブ · スクリプト · アセット"]
アスフォデル -->|"保管場所"|メモリ[".arcforge/memory/*.md"]
グラフ -.->|"追跡"|カロン
アスフォデル -.->|"追跡者"|カロン
読み込み中
エージェントは、88 MCP ツール (Bridge ランチャーとハブを介して実行) を通じて Hades と通信します。 Unity エディター内で、Graph はプロジェクト構造のインデックスを作成し、Asphodel はバージョン管理されたマークダウンとしてプロジェクト メモリを永続化し、Charon はすべての操作をトレースするため、何も隠されていません。
ほとんどの AI for Unity ツールは、2 つのグループのいずれかに分類されます。アクション ブリッジを使用すると、エージェントはエディター アクションを実行できますが、プロジェクトのモデルはありません。コードグラフ/RAG ツールはコードを理解しますが、Unity のアセット層 (プレハブ、シーン、GUID、シリアル化された参照) については認識しません。 Hades は両方を実行し、Unity ネイティブです。
Unity プロジェクト ディレクトリから Claude Code を開き、次のように尋ねます。
このプロジェクトについて教えてください
エージェントはグラフを使用して、一般的な概要ではなく、プロジェクト固有の概要を提供します。
PlayerController はどこで使用するのでしょうか?
テキスト grep だけでなく、シーン、プレハブ、スクリプトにわたる構造検索。
OldNetworkManager を削除したいと考えています。何が壊れるでしょうか？
何かを変更する前に、完全なプロジェクト グラフを通じて参照を追跡する依存関係分析。
証拠が欲しいですか？ Hades がある場合とない場合: 1 つのプロンプトを並べて表示 - 同じタスクが同じ条件下で 2 回実行されます。ストック

k Claude Code は 3 つのプレハブ バリアントを見逃しており、継承を壊すような変更を推奨しています。 Hades は、27% 少ないコストで正しいインパクト マップを返します。完全なノーカット録音と再生ステップが含まれています。
何を信頼すべきか（そして何を検証すべきか）
Hades は自身の確実性について正直です。すべての結果には信頼のシグナルが含まれており、ツールはそれらに依存すべきでない時期を教えてくれます。経験則として:
Hades は神託ではなくナビゲーターです。プロジェクトを迅速かつ構造的に理解できるようにし、独自の盲点を明らかにするので、あなた (およびエージェント) は、何か破壊的なことが起こる前に最新情報を把握できます。各信頼信号の意味については「結果の解釈」と、設計上存在する境界の制限を参照してください。
macOS — 現在テストされている唯一のプラットフォーム。 Windows と Linux はテストされていません。報告は大歓迎です。
Claude Code — 他の MCP クライアントは v1 ではテストされていません
Unity のパッケージ マネージャーで、[git URL からパッケージを追加] をクリックし、次のように入力します。
https://github.com/TheArcForge/Hades.git
ローカル フォルダーから (テストまたはオフラインで使用するため):
Unity のパッケージ マネージャーで、[ディスクからパッケージを追加...] をクリックし、ローカルの Hades フォルダー内の package.json を選択します。
インストール後に初めて開くと、Hades はプロジェクトのナレッジ グラフを自動的に構築します。これは 1 回限りのステップ (進行状況バーの後ろ) です。一般的なプロジェクトでは数秒、非常に大規模なプロジェクトでは最大で数分かかります。その後、更新は増分的かつほぼ瞬時に行われます。
オプション A: 永続インストール (推奨)
/plugin マーケットプレイスに TheArcForge/hades-plugin を追加
/プラグインインストールハデス
オプション B: セッションごと
クロード --plugin-dir /path/to/hades-plugin
それだけです。 Unity プロジェクト ディレクトリから Claude Code を開くと、ツールがすぐに利用可能になります。
初めてですか？各ステップでの検証を含む段階的なチュートリアルについては、完全なスタート ガイドを参照してください。

ステップ。
Claude Code は標準入出力を介して軽量ランチャーに接続し、HTTP リクエストを Hades Hub にルーティングし、次に Hades Hub がツール呼び出しを適切な Unity Editor インスタンスに転送します。ハブはマシンごとに 1 回実行され、複数プロジェクトのルーティングを自動的に処理します。すべてのデータはローカルに残ります。クラウド サービス、テレメトリ、ベンダー ロックインはありません。アーキテクチャの詳細については、Documentation/arcforge-hades-architecture.md を参照してください。
症状
修正
クロードコードにはツールが表示されません
Unityは起動していますか？登録されたインスタンスについては、~/.arcforge/hades-hub/hub.json を確認してください。
再コンパイル後にツールが消える
10 秒ほど待ちます — Unity のドメインのリロード中に Hub ツールの呼び出しがバッファされます。
間違ったプロジェクトがツール呼び出しを受信する
正しいプロジェクト ディレクトリから Claude Code を起動します。
プロジェクト情報が古いようです
/hades:rebuild-graph を実行してナレッジ グラフを再生成します。
完全なトラブルシューティング ガイドについては、Documentation/troubleshooting.md を参照してください。
Hades は v1 です。大規模な実稼働 Unity プロジェクトでフィールドテストされていますが、まだ若いです。静的解析には既知の境界があり (「制限事項」を参照)、多くのプロジェクト、Unity バージョン、またはプラットフォームにわたってまだ実行されていません。ここがあなたの出番です。プロジェクトの結果が間違っていると思われる場合は、問題を開いてください。実際のプロジェクトでの具体的な再現は、まさに v1 が安定する方法です。このツールは、不確実な場合にそれを通知するように構築されているため、信頼性のシグナルを信頼し、有害なことが起こる前に検証してください。
結果の解釈 - それぞれの信頼シグナルの意味とそれに基づいて行動する方法
制限 — 設計上存在する境界
アーキテクチャ — システム設計、データフロー、コンポーネントの責任
プラグインマニフェスト — ツールとスキルのリファレンス
ロードマップ — 開発段階とステータス
ビジョン — 長期的な目標と設計哲学
Claude Code 用の Unity 対応 AI インフラストラクチャ — 既知

wedge グラフ + 88 個の MCP ツールにより、ファイルを grep するだけでなく、AI エージェントにプロジェクトを認識させることができます。
MIT、不明なライセンスが見つかりました
ライセンスが見つかりました
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
1
フォーク
レポートリポジトリ
リリース
1
Hades v1.0.0 — パブリックリリース
最新の
2026 年 6 月 9 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Unity-aware AI infrastructure for Claude Code — a knowledge graph + 88 MCP tools that let your AI agent know your project, not just grep its files. - TheArcForge/Hades

I brought you my new open source project for Unity and Claude Code to work in pair. Long story short: Claude usually has a hard time navigating Unity projects - complex scenes, nested prefabs, and even GameObjects with dozens of components attached, each with dozens of properties. Straightforward grep-and-guess is very weak in terms of efficiency, and — most importantly — in terms of reliability. Claude misses the relations between entities in your game, that's the bitter truth. If you're interested in an actual comparison, concrete example: I asked the agent "which prefabs and scenes break if I change EnemyAI?" Stock Claude Code read ~200k tokens of YAML, found 1 prefab, missed 3 variants, and suggested a change that breaks inheritance. With Hades it returned the correct map — 4 prefabs, 3 scenes — in 7 tool calls, 27% cheaper. Full side-by-side with repro:
https://github.com/TheArcForge/Hades/blob/main/Documentation... Since I work with Unity and Claude daily as part of my job, I came up with a solution. Hades builds a graph out of your project's entities, and Claude can query it via MCP tools. That's the simple version — if you want the details, I wrote up the full story on Medium:
https://medium.com/@mike.kuharuk/unity-context-graph-for-cla... So far I've only proof-tested it on macOS and Unity 6; that scope will grow step by step. Also v1 is mainly built around Claude Code, so Cursor and all other AI agents supports will be added post launch. I'd appreciate technical feedback and suggestions.

GitHub - TheArcForge/Hades: Unity-aware AI infrastructure for Claude Code — a knowledge graph + 88 MCP tools that let your AI agent know your project, not just grep its files. · GitHub
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
TheArcForge
/
Hades
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
38 Commits 38 Commits .claude-plugin .claude-plugin .github .github Bridge~ Bridge~ Dashboard~ Dashboard~ Documentation Documentation Editor Editor Fixtures~/ TestProject/ Assets Fixtures~/ TestProject/ Assets Scanner~ Scanner~ Tests Tests ThirdParty ThirdParty commands commands scripts scripts skills skills .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CHANGELOG.md.meta CHANGELOG.md.meta CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md.meta CODE_OF_CONDUCT.md.meta CONTRIBUTING.md CONTRIBUTING.md CONTRIBUTING.md.meta CONTRIBUTING.md.meta Documentation.meta Documentation.meta Editor.meta Editor.meta LICENSE LICENSE LICENSE.meta LICENSE.meta LIMITATIONS.md LIMITATIONS.md LIMITATIONS.md.meta LIMITATIONS.md.meta README.md README.md README.md.meta README.md.meta SECURITY.md SECURITY.md SECURITY.md.meta SECURITY.md.meta Tests.meta Tests.meta ThirdParty.meta ThirdParty.meta commands.meta commands.meta package.json package.json package.json.meta package.json.meta scripts.meta scripts.meta skills.meta skills.meta View all files Repository files navigation
In the underworld of your Unity project, nothing is hidden from Hades. And now, nothing is hidden from your AI agent.
Hades is Unity-aware AI infrastructure for Claude Code. It builds a queryable knowledge graph of your entire Unity project — every scene, prefab, script, asset, and dependency — so your AI agent knows your project's structure instead of guessing at it. Out of the box you get 88 MCP tools, 22 skills, and 6 commands. Everything runs locally, and everything is version-controllable.
One prompt — "which prefabs and scenes break if I change EnemyAI ?"
Left: stock Claude Code. Right: with Hades.
→ see the full side-by-side breakdown
Stock Claude Code reads ~200k tokens of YAML, finds 1 prefab, misses 3
variants, and tells you to add code that breaks them. Hades answers
correctly — 4 prefabs, 3 scenes — in 7 tool calls for 27% less cost.
That gap is the whole project.
Most AI tools search and predict : they grep for text that looks relevant and let the model infer the rest. The answers are probabilistic — and often wrong in ways you can't see.
Hades lets your agent know and analyze . When it asks "what references PlayerController ," it reads a structural fact from the graph, not a guess from scattered snippets. Dependency analysis traces real edges. Ask the same question twice, get the same answer. One graph query replaces a dozen file reads — and the agent never makes you explain your project twice.
What Hades gives your AI agent
Layer
What it does
Graph
A semantic knowledge graph of your Unity project — scenes, prefabs, scripts, assets, and their dependencies. The agent sees your project's structure, not just its files.
Charon
Full observability — every tool call, graph query, and memory operation is traced. Inspect via the local dashboard ( Hades > Open Charon Dashboard in Unity).
Asphodel
Persistent project memory in version-controlled markdown ( .arcforge/memory/ ). Capture decisions, patterns, and conventions once; the agent reads them for context-aware advice every session.
22 Skills
Architecture decisions, workflow guidance, and domain expertise — networking, audio, UI, shaders, ECS, testing, and more.
88 MCP Tools
21 graph/charon/memory tools + 67 editor-action tools (scenes, components, prefabs, materials, animation, assets).
6 Commands
/hades:status , /hades:rebuild-graph , /hades:show-traces , /hades:validate-memory , /hades:show-proposals , /hades:export-traces
How the pieces fit together
flowchart TD
Agent["AI Agent<br/>(Claude Code)"]
Agent <-->|"88 MCP tools (via Bridge)"| Hades
subgraph Hades["Hades — inside the Unity Editor"]
Graph["Graph<br/>project knowledge graph"]
Asphodel["Asphodel<br/>persistent memory"]
Charon["Charon<br/>observability"]
end
Graph -->|"indexes"| Unity["Unity Project<br/>scenes · prefabs · scripts · assets"]
Asphodel -->|"stored as"| Mem[".arcforge/memory/*.md"]
Graph -.->|"traced by"| Charon
Asphodel -.->|"traced by"| Charon
Loading
Your agent talks to Hades through the 88 MCP tools (carried over the Bridge launcher and hub). Inside the Unity Editor, Graph indexes your project's structure, Asphodel persists project memory as version-controlled markdown, and Charon traces every operation so nothing is hidden.
Most AI-for-Unity tooling falls into one of two camps. Action bridges let an agent execute editor actions but have no model of your project. Code-graph / RAG tools understand code but are blind to Unity's asset layer — prefabs, scenes, GUIDs, serialized references. Hades does both, and it's Unity-native.
Open Claude Code from your Unity project directory and ask:
Tell me about this project
The agent uses the graph to give a project-specific overview — not a generic summary.
Where do we use PlayerController?
Structural search across scenes, prefabs, and scripts — not just text grep.
I want to remove OldNetworkManager. What would break?
Dependency analysis that traces references through the full project graph before you change anything.
Want proof? With and without Hades: one prompt, side by side — the same task run twice under identical conditions. Stock Claude Code misses 3 prefab variants and recommends a change that would break inheritance; Hades returns the correct impact map for 27% less cost. Includes full uncut recordings and reproduction steps.
What to trust (and what to verify)
Hades is honest about its own certainty — every result carries a confidence signal, and the tools tell you when not to rely on them. As a rule of thumb:
Hades is a navigator, not an oracle : it makes understanding your project fast and structural, and it surfaces its own blind spots so you (and your agent) stay in the loop before anything destructive. See Interpreting results for what each confidence signal means, and Limitations for the boundaries that are there by design.
macOS — currently the only tested platform. Windows and Linux are untested; reports welcome.
Claude Code — other MCP clients are untested in v1
In Unity's Package Manager, click Add package from git URL and enter:
https://github.com/TheArcForge/Hades.git
From local folder (for testing or offline use):
In Unity's Package Manager, click Add package from disk... and select the package.json inside your local Hades folder.
On first open after install, Hades automatically builds the project knowledge graph. This is a one-time step (behind a progress bar) — a few seconds on a typical project, up to a few minutes on a very large one. After that, updates are incremental and near-instant.
Option A: Persistent install (recommended)
/plugin marketplace add TheArcForge/hades-plugin
/plugin install hades
Option B: Per-session
claude --plugin-dir /path/to/hades-plugin
That's it. Open Claude Code from your Unity project directory and the tools are available immediately.
First time? See the full Getting Started guide for a step-by-step walkthrough with verification at each step.
Claude Code connects over stdio to a lightweight launcher, which routes HTTP requests to the Hades Hub, which in turn forwards tool calls to the correct Unity Editor instance. The Hub runs once per machine and handles multi-project routing automatically. All data stays local — no cloud services, no telemetry, no vendor lock-in. See Documentation/arcforge-hades-architecture.md for full architectural details.
Symptom
Fix
No tools appear in Claude Code
Is Unity running? Check ~/.arcforge/hades-hub/hub.json for a registered instance.
Tools disappear after recompile
Wait ~10 seconds — the Hub buffers tool calls during Unity's domain reload.
Wrong project receives tool calls
Launch Claude Code from the correct project directory.
Project info seems stale
Run /hades:rebuild-graph to regenerate the knowledge graph.
See Documentation/troubleshooting.md for the full troubleshooting guide.
Hades is v1 — field-tested on a large production Unity project, but young. Static analysis has known boundaries (see Limitations ), and it hasn't yet been exercised across many projects, Unity versions, or platforms. That's where you come in: if a result looks wrong on your project, please open an issue — concrete repros on real projects are exactly how v1 gets solid. The tools are built to tell you when they're uncertain, so trust the confidence signals and verify before anything destructive.
Interpreting results — what each confidence signal means and how to act on it
Limitations — the boundaries that are there by design
Architecture — system design, data flow, component responsibilities
Plugin Manifest — tool and skill reference
Roadmap — development phases and status
Vision — long-term goals and design philosophy
Unity-aware AI infrastructure for Claude Code — a knowledge graph + 88 MCP tools that let your AI agent know your project, not just grep its files.
MIT, Unknown licenses found
Licenses found
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
1
Hades v1.0.0 — Public Release
Latest
Jun 9, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
