---
source: "https://github.com/HBarefoot/engram"
hn_url: "https://news.ycombinator.com/item?id=48796133"
title: "Engram – persistent memory for AI agents, in-process, no cloud"
article_title: "GitHub - HBarefoot/engram: 🧠 Persistent memory for AI agents. SQLite for agent state. Zero cloud dependencies. Local embeddings. MCP-native integration with Claude Desktop/Code, Cursor, Windsurf & more. · GitHub"
author: "barefootdifital"
captured_at: "2026-07-05T17:57:02Z"
capture_tool: "hn-digest"
hn_id: 48796133
score: 2
comments: 0
posted_at: "2026-07-05T17:37:03Z"
tags:
  - hacker-news
  - translated
---

# Engram – persistent memory for AI agents, in-process, no cloud

- HN: [48796133](https://news.ycombinator.com/item?id=48796133)
- Source: [github.com](https://github.com/HBarefoot/engram)
- Score: 2
- Comments: 0
- Posted: 2026-07-05T17:37:03Z

## Translation

タイトル: Engram – AI エージェント用の永続メモリ、インプロセス、クラウドなし
記事のタイトル: GitHub - HBarefoot/engram: 🧠 AI エージェントの永続メモリ。エージェント状態の SQLite。クラウド依存性ゼロ。ローカル埋め込み。 Claude Desktop/Code、Cursor、Windsurf などとの MCP ネイティブ統合。 · GitHub
説明: 🧠 AI エージェントの永続メモリ。エージェント状態の SQLite。クラウド依存性ゼロ。ローカル埋め込み。 Claude Desktop/Code、Cursor、Windsurf などとの MCP ネイティブ統合。 - H裸足/エングラム

記事本文:
GitHub - HBarefoot/engram: 🧠 AI エージェント用の永続メモリ。エージェント状態の SQLite。クラウド依存性ゼロ。ローカル埋め込み。 Claude Desktop/Code、Cursor、Windsurf などとの MCP ネイティブ統合。 · GitHub
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
H裸足
/
エングラム
公共

ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
113 コミット 113 コミット .github .github アセット/ ブランド アセット/ ブランド ベンチ ベンチ ビン ビン ダッシュボード ダッシュボード デスクトップ デスクトップ ドキュメント ドキュメントの例 例 モデル モデル スクリプト スクリプト スキル/ エングラムメモリ スキル/ エングラムメモリ src src 統計 統計 テスト テスト .dockerignore .dockerignore .gitignore .gitignore .npmignore .npmignore BUSINESS_MODEL.md BUSINESS_MODEL.md CHANGELOG.md CHANGELOG.md CLA.md CLA.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md DESIGN_SYSTEM.md DESIGN_SYSTEM.md Dockerfile Dockerfile LICENSE LICENSE QUICKSTART.md QUICKSTART.md README.md README.md SECURITY.md SECURITY.md Ecosystem.config.cjs Ecosystem.config.cjs engram-logo.png engram-logo.png eslint.config.js eslint.config.js Glama.json Glama.json package-lock.json package-lock.json package.json package.json server.json server.json vitest.config.js vitest.config.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント用の永続メモリ。進行中。インフラはありません。
クラウド、API キー、Docker を使用せずに、AI エージェントに長年一緒に働いてきた同僚の思い出を与えましょう。
⭐役に立ちましたか? GitHub でスターを付けます。これは、他の人が Engram を見つけるのを助ける最も簡単な方法です。
npm install -g @hbarefoot/engram
エングラムスタート
AI エージェントは長期記憶を持ちます。 2 分、セットアップなし、クラウドなし。
🧠 インプロセス — エージェントのスタック内で実行されます。個別のサーバーをデプロイする必要はなく、IPC オーバーヘッドもフォークする必要もありません。
📴 オフライン — ローカル SQLite + バンドルされた埋め込み (~23 MB)。 API キーやデータがマシンから流出することはありません。
🔌 MCP ネイティブ —

Claude Desktop、Claude Code、Cursor、Windsurf、および Cline とのファーストクラスの Model Context Protocol の統合。
🔐 デフォルトで安全 — 書き込みごとに自動シークレット検出。 API キー、秘密キー、接続文字列、JWT はデータベースにアクセスする前にブロックされます。
Engram はエージェントのプロセス内で実行されます。デプロイするサービスやアカウントはなく、マシンから何も出ません。その設計の選択は測定可能です。
Apple M4 Pro で 1,000 を超えるシード済みメモリを測定 - npm run bench で再現。これらはフットプリントとレイテンシの数値であり、精度を主張するものではありません。Engram はメモリ ベンチマークで Mem0 や Zep を上回ろうとはしません。ポイントは、操作面がないしっかりとしたリコールです。
オプションの精度リフト — 依然として 100% ローカルです。すでにローカル モデルを実行している場合は、オプトイン LLM レイヤーによってファクト抽出が強化されます。推奨される henrybarefoot1987/engram-extract モデル (qwen3:1.7b) を使用すると、エンティティ抽出の精度が 45.8% (ルールベース) から 95.8% に上昇し、デバイスから 1 バイトも外部に流出することはありません。
Engram は無料で MIT ライセンスを取得しており、今後もそのようです。ペイウォール、階層ロック機能、テレメトリはありません。すべての機能はオープンソース パッケージで提供されます。スポンサーシップは純粋に継続的な開発に資金を提供する方法であり、何かのロックを解除するものではありません。
Engram によって時間を節約できる場合は、Polar 経由でスポンサーになることができます。
エンタープライズについて。 Engram は MIT ライセンスを取得しているため、商用利用がすでに許可されており、職場で使用するためにライセンスを購入する必要はありません。エンタープライズ層では、問題に対する優先対応と、Engram をスタックに接続する専用のヘルプを購入します。 MIT ライセンスのソフトウェアに依存することをポリシーで禁止している組織の場合は、リクエストに応じてオプションの商用ライセンスのオーバーライドを利用できます。 (Engram は個人の開発者によって保守されているため、これはベストエフォート型の優先対応であり、契約上の SLA ではありません。

)
ほとんどのエージェント メモリ製品は、Postgres、Docker、クラウド アカウント、API キーなど、エージェントと一緒に実行するサービスです。 Engram は、エージェントのプロセス内に埋め込まれます。実用的なガードレールを備えた、集中的で安定した npm パッケージです。
ソース: @sunriselabs/lodis 、 Sunrise-Labs-Dot-AI/engrams 、 mem0.ai 、 github.com/getzep/zep 、 github.com/letta-ai/letta 。詳細については、docs/competitive-intel.md を参照してください。 Engram には、オプションのオンデバイス LLM 抽出 (v1.9 以降) が同梱されています。ローカル モデルで llm.* をポイントします。推奨される henrybarefoot1987/engram-extract (Qwen3-1.7B、Apache-2.0) または任意の Ollama / OpenAI 互換エンドポイントです。カテゴリ/エンティティ抽出を強化します (ルールに対してエンティティ認識 +50 ポイント — 45.8% → 95.8% — engram-extract (qwen3:1.7b) (ベンチマークでは)、依然として 100% ローカルで、デフォルトでオフになっています (ゼロ構成パスはルールベース、オフライン、インフラフリーのままです)。 Mem0/Zep/Letta はクラウド モデル経由で LLM 抽出を構築します。 Lodis は LLM フリーの読み取り/書き込みであり、より広範な機能面を備えています。私たちはそれを正直にリストしています。
TL;DR — それぞれが当てはまる場合。実用的なガードレール (秘密検出、エージェントの自動検出、デスクトップ アプリ)、シンプルな 5 カテゴリのメンタル モデル、必要に応じてオプションのオンデバイス LLM 抽出を備えた、集中的で安定したローカル ファーストのメモリ層が必要な場合は、Engram を選択してください。 14 のエンティティ タイプと時間的スーパーセッションを備えたナレッジ グラフ スタイルのメモリが必要な場合は、Lodis を選択してください。クラウド LLM 抽出を組み込み、そのための運用インフラストラクチャを気にしない場合は、Mem0/Zep/Letta を選択してください。
npm install -g @hbarefoot/engram
2. サーバーを起動します
engram start # MCP + REST + ローカルホスト上のダッシュボード:3838
engram start --mcp-only # MCP サーバーのみ、stdio モード (エージェント統合用)
3. AI エージェントを接続する
クロード mcp add engram -- engram start --mcp-only
クロード デスクトップ — ~/Library/Application Support/ に追加

クロード/claude_desktop_config.json :
{
"mcpサーバー": {
"エングラム" : {
"コマンド" : "エングラム" ,
"args" : [ " start " 、 " --mcp-only " ]
}
}
}
Cline / Cursor / Windsurf — 同じ mcpServers ブロックをエディターの MCP 構成に追加します。 http://localhost:3838 にある組み込みダッシュボードには、インストールされているエージェントを自動検出し、構成を生成する統合ウィザードがあります。
あなた: 「私たちの API は 24 時間有効期限のある JWT トークンを使用していることに注意してください。」
クロード: (engram_remember 経由で保存)
あなた: (翌日) 「どのような認証アプローチを使用していますか?」
クロード: (engram_recall によるリコール) — 「JWT トークン、有効期限は 24 時間です。」
メモリはセッション間、マシンの再起動間、さらには同じ Engram インスタンスを共有する異なる AI クライアント間でも保持されます。
時間の経過とともに改善される記憶力
ほとんどのメモリ システムは追加専用ストアです。一度書き込めば永久に取得でき、最善を尽くします。エングラムは学びます。
フィードバック ループ ( engram_フィードバック ) — エージェントが記憶を思い出したときに、あなたまたはエージェントはそれが役立つか役に立たないかを投票できます。思い出は [-1, 1] にスコアを蓄積します。役に立たない記憶が続くと、自信は自動的に低下します。
矛盾の検出 — 2 つのメモリが競合する場合 (「Fastify を優先」対「Express に切り替え」)、統合エンジンはそれらにフラグを立てます。ダッシュボードの [競合] タブには、A を維持、B を維持、両方を維持、または却下という 4 つの解決アクションが並べて表示されます。
挿入時の重複排除 - 同一のメモリ (コサイン類似度 0.95 以上) は拒否されます。ほぼ重複 (0.92 ～ 0.95) は、新しいコンテンツを既存のレコードに吸収します。手作業で剪定をしなくても店内はきれいなままです。
衰退 — 思い出されない記憶は時間の経過とともに信頼を失い、将来の結果を汚さなくなります。
エングラムを長く使用するほど、その記憶はより鮮明になります。
Engram は、標準入出力経由で AI エージェントに 6 つのツールを公開します。
事実 — オブ

セットアップ、アーキテクチャ、または構成に関する主観的な真実。
好み — ユーザーの好き嫌い、スタイルの選択。
パターン — 繰り返し発生するワークフローと習慣。
決定 — 行われた選択とその背後にある理由。
結果 — 実行されたアクションの結果。
エージェントに Engram の使用方法を教える
MCP サーバーに接続すると、エージェントに記憶ツールが提供されますが、それを適切に使用するための判断力はありません。バンドルされているエングラム記憶スキルは、その判断層です。エージェントは、毎回指示されることなく、セッションの開始時に思い出し、永続的な決定、修正、結果が発生したときに保存し、最後に結果を書き戻すように教えられます。
エングラムスキルインストール # → ~/.claude/skills/engram-memory/
engram skill install --project # → ./.claude/skills/ (チーム用にコミット)
engram skill install --platform Agents # → ~/.agents/skills/ (クロスフレームワーク)
Claude Code、Claude Desktop、Cowork、またはエージェント スキル仕様 ( .agents/skills ) を読み取る任意のフレームワークで動作します。スキルはパッケージで販売されるため、Engram でバージョンが更新され、次回の Engram スキルのインストール時に更新されます。エングラムスキルアンインストールはそれをきれいに削除します。
engram start # MCP + REST + ダッシュボードを開始します
engram start --mcp-only # MCP サーバーのみ (stdio モード)
engram start --port 3838 # カスタム REST ポート
engram remember " <content> " # 記憶を保存する (-c category -eentity -n namespace --confidence)
engram remember " <query> " # 記憶の検索 (-l limit -c category -n namespace --threshold)
engram remember < id > # IDで削除
engram list # メモリのリスト (-l limit --offset -c category -n namespace)
エングラムステータス # ヘルスチェック
engram consolidate # 重複排除、矛盾の検出、減衰
# (--重複なし / --矛盾なし / --減衰なし / --クリーンアップ-古い)
engramconflicts # 未解決の矛盾をリストする
エングラムエクスポート-

context # 厳選されたコンテキスト ブロックをエクスポートする
# (-o file -f markdown|claude|txt|json -c カテゴリ --min-confidence ...)
engram import # ローカルソースからインポートする
# (-s カーソルルール|クロード|パッケージ|git|ssh|シェル|黒曜石|env --dry-run)
engram skill install # engram-memory エージェント スキルをインストールします
# (--プロジェクト → ./.claude、--プラットフォームエージェント → ~/.agents)
engram skill uninstall # engram-memory スキルを削除します
完全なフラグ リストを表示するには、engram --help を実行します。
REST API は、デフォルトでは http://localhost:3838 で実行されます。
http://localhost:3838 にある組み込みの React ダッシュボード:
ダッシュボード — メモリ統計、最近のアクティビティ、ヘルスゲージ。
思い出 — 参照、フィルター、インライン編集、一括削除。
検索 — スコアの内訳を含むセマンティック検索。
統計 — カテゴリ、名前空間、および時間ごとのグラフ。
健康 — ワンクリックでクリーンアップできる、古くなり、決して思い出されない、フィードバックの少ない記憶。
矛盾 — 並列的な矛盾の解決。
エージェント — インストールされている AI クライアントを自動検出し、その MCP 構成を (タイムスタンプ付きのバックアップとともに) 書き込む統合ウィザード。
インポート — カーソルルール、.claude ファイル、package.json、git config、SSH config、シェル履歴、Obsidian、および .env のウィザード。
Store : engram_remember はシークレット検出を通じてコン​​テンツを実行し、all-MiniLM-L6-v2 を使用してローカルに埋め込みます (~23 MB、CPU のみ、一度ダウンロードされ ~/.engram/models/ にキャッシュされます)

[切り捨てられた]

## Original Extract

🧠 Persistent memory for AI agents. SQLite for agent state. Zero cloud dependencies. Local embeddings. MCP-native integration with Claude Desktop/Code, Cursor, Windsurf & more. - HBarefoot/engram

GitHub - HBarefoot/engram: 🧠 Persistent memory for AI agents. SQLite for agent state. Zero cloud dependencies. Local embeddings. MCP-native integration with Claude Desktop/Code, Cursor, Windsurf & more. · GitHub
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
HBarefoot
/
engram
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
113 Commits 113 Commits .github .github assets/ brand assets/ brand bench bench bin bin dashboard dashboard desktop desktop docs docs examples examples models models scripts scripts skills/ engram-memory skills/ engram-memory src src stats stats test test .dockerignore .dockerignore .gitignore .gitignore .npmignore .npmignore BUSINESS_MODEL.md BUSINESS_MODEL.md CHANGELOG.md CHANGELOG.md CLA.md CLA.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DESIGN_SYSTEM.md DESIGN_SYSTEM.md Dockerfile Dockerfile LICENSE LICENSE QUICKSTART.md QUICKSTART.md README.md README.md SECURITY.md SECURITY.md ecosystem.config.cjs ecosystem.config.cjs engram-logo.png engram-logo.png eslint.config.js eslint.config.js glama.json glama.json package-lock.json package-lock.json package.json package.json server.json server.json vitest.config.js vitest.config.js View all files Repository files navigation
Persistent memory for AI agents. In-process. No infra.
Give your AI agent the memory of a colleague who's worked with you for years — without cloud, API keys, or Docker.
⭐ Useful to you? Star it on GitHub — it's the simplest way to help others find Engram.
npm install -g @hbarefoot/engram
engram start
Your AI agent now has long-term memory. Two minutes, no setup, no cloud.
🧠 In-process — runs inside your agent's stack. No separate server to deploy, no IPC overhead, nothing to fork.
📴 Offline — local SQLite + bundled embeddings (~23 MB). No API keys, no data leaving your machine.
🔌 MCP-native — first-class Model Context Protocol integration with Claude Desktop, Claude Code, Cursor, Windsurf, and Cline.
🔐 Safety by default — automatic secret detection on every write. API keys, private keys, connection strings, JWTs blocked before they hit the database.
Engram runs inside your agent's process — no service to deploy, no account, nothing leaving your machine. That design choice is measurable:
Measured on an Apple M4 Pro over 1,000 seeded memories — reproduce with npm run bench . These are footprint and latency numbers, not an accuracy claim: Engram doesn't try to out-rank Mem0 or Zep on memory benchmarks. The point is solid recall with none of the operational surface.
Optional accuracy lift — still 100% local. If you already run a local model, the opt-in LLM layer sharpens fact extraction: entity-extraction accuracy climbs from 45.8% (rule-based) to 95.8% with the recommended henrybarefoot1987/engram-extract model (qwen3:1.7b) — +50 pts — without a single byte leaving your device.
Engram is free and MIT-licensed — and always will be. No paywalls, no tier-locked features, no telemetry. Every feature ships in the open-source package. Sponsorship is purely a way to fund continued development, not to unlock anything.
If Engram saves you time, you can sponsor it via Polar :
About Enterprise. Engram is MIT-licensed, so commercial use is already granted — you don't need to buy a license to use it at work. The Enterprise tier buys priority response on issues and dedicated help wiring Engram into your stack . For organizations whose policy precludes depending on MIT-licensed software, an optional commercial-license override is available on request. (Engram is maintained by a solo developer, so this is best-effort priority response, not a contractual SLA.)
Most agent-memory products are services you run alongside your agent — Postgres, Docker, cloud accounts, API keys. Engram embeds inside your agent's process: a focused, stable npm package with practical guardrails.
Sources: @sunriselabs/lodis , Sunrise-Labs-Dot-AI/engrams , mem0.ai , github.com/getzep/zep , github.com/letta-ai/letta . See docs/competitive-intel.md for the full breakdown. Engram ships optional, on-device LLM extraction (v1.9+): point llm.* at a local model — the recommended henrybarefoot1987/engram-extract (Qwen3-1.7B, Apache-2.0) or any Ollama / OpenAI-compatible endpoint — to sharpen category/entity extraction (entity recognition +50 pts vs rules — 45.8% → 95.8% — with engram-extract (qwen3:1.7b) in our benchmark), still 100% local and off by default (the zero-config path stays rule-based, offline, and infra-free). Mem0/Zep/Letta build LLM extraction in via a cloud model; Lodis is LLM-free read/write with a broader feature surface — we list it honestly.
TL;DR — when each one fits. Pick Engram if you want a focused, stable, local-first memory layer with practical guardrails (secret detection, agent auto-discovery, desktop app), a simple 5-category mental model, and optional on-device LLM extraction when you want it. Pick Lodis if you want a knowledge-graph-style memory with 14 entity types and temporal supersession. Pick Mem0/Zep/Letta if you want cloud-LLM extraction built in and don't mind operating infrastructure for it.
npm install -g @hbarefoot/engram
2. Start the server
engram start # MCP + REST + Dashboard on localhost:3838
engram start --mcp-only # MCP server only, stdio mode (for agent integration)
3. Connect your AI agent
claude mcp add engram -- engram start --mcp-only
Claude Desktop — add to ~/Library/Application Support/Claude/claude_desktop_config.json :
{
"mcpServers" : {
"engram" : {
"command" : " engram " ,
"args" : [ " start " , " --mcp-only " ]
}
}
}
Cline / Cursor / Windsurf — add the same mcpServers block to your editor's MCP config. The built-in dashboard at http://localhost:3838 has an Integration Wizard that auto-detects your installed agents and generates the config for you.
You: "Remember that our API uses JWT tokens with 24-hour expiry."
Claude: (stores via engram_remember)
You: (next day) "What authentication approach are we using?"
Claude: (recalls via engram_recall) — "JWT tokens, 24-hour expiry."
Memories persist across sessions, machine restarts, and even between different AI clients sharing the same Engram instance.
Memory that improves over time
Most memory systems are append-only stores: write once, retrieve forever, hope for the best. Engram learns.
Feedback loop ( engram_feedback ) — when an agent recalls a memory, you or the agent can vote it helpful or unhelpful. Memories accumulate a score in [-1, 1] ; consistently-unhelpful memories see their confidence decay automatically.
Contradiction detection — when two memories conflict ("prefers Fastify" vs "switched to Express"), the consolidation engine flags them. The dashboard's Conflicts tab shows them side-by-side with four resolution actions: keep A, keep B, keep both, or dismiss.
Deduplication on insert — identical memories (≥0.95 cosine similarity) are rejected. Near-duplicates (0.92–0.95) absorb the new content into the existing record. The store stays clean without manual pruning.
Decay — memories that aren't recalled lose confidence over time and stop polluting future results.
The longer you use Engram, the sharper its recall gets.
Engram exposes 6 tools to AI agents over stdio:
fact — Objective truths about setup, architecture, or configuration.
preference — User likes, dislikes, style choices.
pattern — Recurring workflows and habits.
decision — Choices made and the reasoning behind them.
outcome — Results of actions taken.
Teach your agent to use Engram
Connecting the MCP server gives your agent the memory tools — but not the judgment to use them well. The bundled engram-memory skill is that judgment layer: it teaches an agent to recall at the start of a session, store durable decisions, corrections, and outcomes as they happen, and write results back at the end — without being told each time.
engram skill install # → ~/.claude/skills/engram-memory/
engram skill install --project # → ./.claude/skills/ (commit it for your team)
engram skill install --platform agents # → ~/.agents/skills/ (cross-framework)
Works in Claude Code, Claude Desktop, Cowork, or any framework that reads the Agent Skills spec ( .agents/skills ). The skill is vendored in the package, so it versions with Engram and updates land on the next engram skill install ; engram skill uninstall removes it cleanly.
engram start # Start MCP + REST + dashboard
engram start --mcp-only # MCP server only (stdio mode)
engram start --port 3838 # Custom REST port
engram remember " <content> " # Store a memory (-c category -e entity -n namespace --confidence)
engram recall " <query> " # Search memories (-l limit -c category -n namespace --threshold)
engram forget < id > # Delete by ID
engram list # List memories (-l limit --offset -c category -n namespace)
engram status # Health check
engram consolidate # Deduplicate, detect contradictions, decay
# (--no-duplicates / --no-contradictions / --no-decay / --cleanup-stale)
engram conflicts # List unresolved contradictions
engram export-context # Export curated context block
# (-o file -f markdown|claude|txt|json -c categories --min-confidence ...)
engram import # Import from local sources
# (-s cursorrules|claude|package|git|ssh|shell|obsidian|env --dry-run)
engram skill install # Install the engram-memory agent skill
# (--project → ./.claude, --platform agents → ~/.agents)
engram skill uninstall # Remove the engram-memory skill
Run engram --help for the full flag list.
The REST API runs on http://localhost:3838 by default.
A built-in React dashboard at http://localhost:3838 :
Dashboard — Memory stats, recent activity, health gauge.
Memories — Browse, filter, inline-edit, bulk-delete.
Search — Semantic search with score breakdown.
Statistics — Charts by category, namespace, and time.
Health — Stale, never-recalled, low-feedback memories with one-click cleanup.
Conflicts — Side-by-side contradiction resolution.
Agents — Integration wizard that auto-detects installed AI clients and writes their MCP configs (with timestamped backups).
Import — Wizard for cursorrules, .claude files, package.json, git config, SSH config, shell history, Obsidian, and .env.
Store : engram_remember runs content through secret detection, then embeds it locally using all-MiniLM-L6-v2 (~23 MB, CPU-only, downloaded once and cached at ~/.engram/models/

[truncated]
