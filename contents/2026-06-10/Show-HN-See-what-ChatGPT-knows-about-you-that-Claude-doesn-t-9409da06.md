---
source: "https://github.com/Thinklanceai/agentkeeper"
hn_url: "https://news.ycombinator.com/item?id=48468754"
title: "Show HN: See what ChatGPT knows about you that Claude doesn't"
article_title: "GitHub - Thinklanceai/agentkeeper: Own your AI memory — import ChatGPT, Claude and Gemini exports, see what each AI knows about you. Checkpoint/restore and cross-model continuity for agents. · GitHub"
author: "tomtom1977"
captured_at: "2026-06-10T01:00:44Z"
capture_tool: "hn-digest"
hn_id: 48468754
score: 2
comments: 0
posted_at: "2026-06-09T22:35:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: See what ChatGPT knows about you that Claude doesn't

- HN: [48468754](https://news.ycombinator.com/item?id=48468754)
- Source: [github.com](https://github.com/Thinklanceai/agentkeeper)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T22:35:01Z

## Translation

タイトル: HN を表示: ChatGPT があなたについて知っていて、クロードが知らないことを確認してください
記事のタイトル: GitHub - Thinklanceai/agentkeeper: AI メモリを所有する — ChatGPT、Claude、Gemini エクスポートをインポートし、各 AI があなたについて何を知っているかを確認します。エージェントのチェックポイント/復元とモデル間の継続性。 · GitHub
説明: AI メモリを所有します — ChatGPT、Claude、Gemini のエクスポートをインポートし、各 AI があなたについて何を知っているかを確認します。エージェントのチェックポイント/復元とモデル間の継続性。 - Thinklanceai/エージェントキーパー

記事本文:
GitHub - Thinklanceai/agentkeeper: AI メモリを所有します。ChatGPT、Claude、Gemini のエクスポートをインポートし、各 AI があなたについて何を知っているかを確認します。エージェントのチェックポイント/復元とモデル間の継続性。 · GitHub
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
シンランサイ
/
エージェントキーパー
公共
ノーティ

フィクション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
52 コミット 52 コミット .github/ workflows .github/ workflows エージェントキーパー エージェントキーパー ベンチマーク ベンチマーク ドキュメント ドキュメントの例 例 マーケティング マーケティング テスト テスト .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンスREADME.md README.md pyproject.toml pyproject.toml 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
長寿命の AI エージェントのための認知継続性インフラストラクチャ。
エージェントは、モデルの切り替え、クラッシュ、コンテキスト ウィンドウの制限、再起動の後も、以前と同じ ID、メモリ、優先順位を維持して生き残ります。
エージェントが失敗するのは、事実を忘れたからではありません。モデルが変更されるか、コンテキスト ウィンドウが埋まるか、プロセスが再起動されると、その状態、優先順位、アイデンティティがドリフトするため、認知的連続性が失われるために失敗します。
AgentKeeper は、これをメモリの問題ではなくシステムの問題として扱います。
pip インストール エージェントキーパー-ai
必要な依存関係はゼロです。外部インフラストラクチャはありません。ストレージのデフォルトはローカル SQLite です。
pip install ' エージェントキーパー-ai[anthropic] ' # クロード
pip install 'agentkeeper-ai[openai]' # GPT + OpenAI 埋め込み
pip install ' Agentkeeper-ai[gemini] ' # Gemini
pip install ' Agentkeeper-ai[semantic] ' # ローカル埋め込み (文変換)
pip install ' Agentkeeper-ai[mcp] ' # MCP サーバー (クロード デスクトップ、カーソル、コーデックス)
pip install ' Agentkeeper-ai[encrypted] ' # 保存時の暗号化ストレージ
pip install ' Agentkeeper-ai[all] ' # すべて
AgentKeeper が他にはない 5 つのこと
1. すべてを生き抜くアイデンティティ
原則と制約は保護されています - 原則と制約は免除されます

トークンの予算に関係なく、再構築されたすべてのコンテキストに注入される非常に圧縮パス。これらは、衰退、統合、矛盾の調停、モデルの切り替え、およびプロセスの再起動を経ても生き残ります。
輸入代理店キーパー
エージェント = エージェントキーパー 。 create ( エージェント ID = "aria" 、プロバイダー = "anthropic" )
代理店。 set_identity (
名前 = "アリア" 、
役割 = "EU 保険ブローカーの副操縦士" ,
原則 = [ "明示的な同意なしに PII を決して共有しない" ],
制約 = [ "EU データ常駐のみ" ],
)
代理店。原則（「予算の変更は必ず書面で確認する」）
代理店。ファクト ( "クライアント: Acme Corporation" 、重要度 = 0.95 )
代理店。イベント ( "契約署名" 、 when = "2026-05-15" )
代理店。保存()
# 100 圧縮サイクル後 - アイデンティティはそのままです。
2. メモリは同じ、モデルを切り替える
認知状態は、各モデルが期待する形式で再構築されます。クロードには XML、GPT-4 にはラベル付きセクション、ジェミニには物語的な散文、オラマには簡潔なトークン。 1 つのエージェント、4 つのランタイム、リライトはゼロ。
エージェント = エージェントキーパー 。ロード ( "アリア" 、プロバイダー = "anthropic" )
応答 = エージェント。尋ねる (「Acme について何を知っていますか?」)
代理店。 switch_provider ( "openai" )。保存()
応答 = エージェント。 ask (「同じ質問ですが、モデルが異なります。」)
# 記憶とアイデンティティはそのままです。フォーマットが変更されました。何も壊れませんでした。
3. GDPR の TTL — 自動的に期限切れになるメモリ
ファクトとグラフのトリプルは TTL を受け入れます。期限が過ぎると、 purge_expired() によってそれらが削除されます。手動クリーンアップはありません。デフォルトで準拠。
エージェント 。ファクト ( "セッショントークン: abc123" 、 ttl = "1h" )
代理店。ファクト ( "監査ログ参照: 2026 年第 1 四半期" 、 ttl = "90d" )
代理店。リンク ( "Acme" 、 "signed_contract" 、 "ThinkLanceAI" 、 ttl = "2y" )
代理店。 purge_expired () # 期限切れのものは削除し、保護されているものは保持します
4. グラフトラバーサル — 散文記憶と並行した構造化された関係
事実は散文です。トリプルはストライクです

キャプチャ。どちらも同じエージェント内に存在し、独自の保持ポリシーと TTL ポリシーが適用されます。
エージェント 。リンク ( "Acme" 、 "owns" 、 "Globex" )
代理店。リンク ( "Globex" 、 "located_in" 、 "BE" )
代理店。 link ( "Alice" 、 "works_at" 、 "Acme" 、信頼度 = 0.9 )
関連 = エージェント。 find_relative ( "Acme" 、 max_hops = 2 、方向 = "out" )
# {"グローベックス": 1、"BE": 2、"アリス": 1}
5. MCP 経由で Claude デスクトップに接続します。
AgentKeeper には MCP サーバーが付属しています。 MCP 対応クライアント (Claude Desktop、Cursor、Claude Code) は、統合コードを 1 行も記述することなく、エージェントのコグニティブ レイヤーに完全にアクセスできます。
Agentkeeper-mcp --agent-id aria --provider anthropic
クロード_デスクトップ_config.json :
{
"mcpサーバー": {
"アリア" : {
"コマンド" : " エージェントキーパー-mcp " ,
"args" : [ " --agent-id " 、 " aria " 、 " --provider " 、 " anthropic " ]
}
}
}
MCP 上で利用可能なツール: add_fact 、recall 、set_identity 、link 、find_popular 、compress 、health 、gdpr_export 、purge_expired 、checkpoint 、restore 、list_checkpoints 。
6. チェックポイント — スナップショット、復元、クラッシュ後の生存
完全な認知状態を不変のコンテンツハッシュ化された状態に凍結します。
スナップショット。クラッシュ、コンテキスト ウィンドウのオーバーフロー、モデルの後に復元する
切り替え、またはプロセスの再起動。不透明な実行状態ペイロードをアタッチします
(現在のファイル、保留中のタスク、Todo) - AgentKeeper はそれを保存して返します。
逐語的に;それを解釈したり実行したりすることはありません。
スナップ = エージェント。チェックポイント (
label = "リファクタリング前" ,
実行状態 = {
"current_file" : "auth.py" ,
"pending_task" : "RS256 移行を終了" ,
}、
)
# クラッシュ、再起動、モデルの切り替え - プロセス メモリがなくなった
エージェント = エージェントキーパー 。ロード (「アリア」)。復元 (スナップ . snapshot_id )
# ID、ファクト、グラフ、および実行状態が戻ってきました
代理人。 diff ( snap_a , snap_b ) # 事実の差分: 事実の追加/削除/変更
代理店。リストチェックポイント

s () # すべてのスナップショット、最も古いものから順
再構築は決定的であり、常に同じスナップショットによって再構築されます。
同じ認知状態であり、コンテンツ ハッシュによって検証されます。行動はそうではありません
保証: AgentKeeper は、モデルの次の状態ではなく、エージェントの状態を復元します。
決断。クラッシュ回復デモは API キーなしで実行されます。
Python の例/crash_recovery.py
7. AI メモリを所有する — ChatGPT、Claude、Gemini をインポートする
すべての AI プラットフォームは、あなたに関することを記憶しています。どれもあなたに見せません
なんと、彼らは誰もお互いに話していません。 AgentKeeper は、
3 つすべての公式データを 1 つの範囲指定されたクエリ可能なメモリにエクスポートします。
次に、各プラットフォームが知っていて他のプラットフォームが知らないことを正確に教えてくれます。
輸入代理店キーパー
エージェントキーパーから。輸入業者が輸入（
ChatGPTImporter 、ClaudeImporter 、GeminiImporter 、
比較スコープ 、 スコープインベントリー 、
)
エージェント = エージェントキーパー 。作成 ( "my-memory" )
ChatGPTImporter (「chatgpt-export.zip」)。申請（代理人）
ClaudeImporter (「claude-export.zip」)。申請（代理人）
GeminiImporter (「takeout.zip」)。申請（代理人）
print (scope_inventory (エージェント))
# {"プラットフォーム:chatgpt": {"イベント": 142, "優先順位": 9, ...},
# "プラットフォーム:クロード": {"イベント": 87, "事実": 4, ...}, ...}
result = Compare_scopes (エージェント、 "プラットフォーム:chatgpt" 、 "プラットフォーム:クロード" )
# ChatGPT のみ: ['電子音楽のツアー プラットフォームの構築']
# クロードのみ: ['EU AI 法の監視機関に取り組んでいます']
# 共有: [('私の名前はトム、ブリュッセルの開発者です', 1.0)]
インポートされたすべてのファクトは名前空間化されている (metadata["scope"] == "platform:chatgpt") ため、1 つのプラットフォームからの知識が他のプラットフォームに漏洩することはありません。
あなたがそれを宣伝しない限り、他の人のコンテキスト。マッチングは決定的です —
正規化されたコンテンツの同等性と difflib 比率、埋め込みなし、なし
LLM 呼び出し、バイトまで再現可能。 PII (アカウントの電子メール、フルネーム)
あなたがいない限り外出しない

include_pii=True を渡します。アーカイブは次で解析されます
ジップボムとパストラバーサルガードを備えたメモリ。何もない
ディスクに抽出されます。
プロンプト生成メモリ用の parse_memory_summary(text) も同梱されています
2026 年のクロスプラットフォーム移行フローで使用される概要。
┌─────────────────────────┐
│ AgentKeeper パブリック API │
│ エージェント.remember() · エージェント.recall() · エージェント.ask() │
│ エージェント.compress() · エージェント.リンク() · エージェント.find_relative() │
│ エージェント.set_identity() · エージェント.パージ_期限切れ() · エージェント.save() │
━━━━━━━━━━━━━━━━━━━━━━━━━━━┘
│
┌─────────────▼─────────────────┐
│ 認知再構成エンジン (CRE) │
│ ID インジェクション・重要度ランキング・セマンティックブースト │
│ トークンバジェット・プロファイル駆動型レンダリング │
━─┬───────┬─────┬─────┬─────────┘
│ │ │ │
┌─────▼─────┐ ┌──▼─────┐ ┌▼──────┐ ┌▼────────┐
│ 記憶 │ │ 意味論的 │ │ 認知的 │ │ クロスモデル │
│ 階層 │ │ 想起 │ │ 圧縮 │ │ 翻訳 │
│ │ │ │ │ │ │
│ 働く │

│ 埋め込み │ │ 減衰 │ │ XML (クロード) │
│ エピソード │ │ ベクトル IDX │ │ コンソール。 │ │ セクション (GPT)│
│ セマンティック │ │ sqlite-vec │ │ 矛盾 │ │ 物語 (G.)│
│ アーカイブ │ │ │ │ │ │ ミニマル (Oll.)│
━━━━━━━━━━━━━━━━━━━┘ ━━━━━━━┘ ━━━━━━━━━━━━━┘
│ │
┌─────▼─────────────────────────┐
│ グラフレイヤーストレージ (プラグ可能) │
│ トリプル、TTL、BFS SQLite · 暗号化された SQLite · Postgres* │
│ Agent.link() AGENTKEEPER_DB、AGENTKEEPER_ENC_KEY │
━━━━━━━━━━━━━━━━━━━━━━┘
│
┌───────────▼───────────┐
│ MCPサーバー・LangChain・CrewAI │
│ エージェントキーパー-mcp · langchain_system_prompt│
━━━━━━━━━━━━━━━━━━━┘
*Postgres スタブが利用可能です。 v1.2 で完全に実装されます。
エージェント 。ファクト ( "予算: 50,000 ユーロ" 、重要度 = 0.9 ) # 安定したセマンティック ファクト
代理店。イベント ( "契約署名" 、 when = "2026-05-15" ) # エピソード、時間固定
代理店。原則 (「PII を決して共有しない」) # 保護され、すべての圧縮に耐えます
代理店。 remember ( "favorite color: blue" ) # ティアは自動的に推測されます
意味的想起
結果 = エージェント。思い出してください ( "プロジェクトに割り当てられたお金

"、top_k = 5)
実際、結果のスコアは次のとおりです。
print ( f" { スコア :.2f } { ファクト . コンテンツ } " )
プラグイン可能なバックエンド: ローカルの文変換 (デフォルト、無料、オフライン)、OpenAI、または独自のもの。 sqlite-vec を介した永続的なインデックス — プロセスの再起動後も再構築せずに存続します。
レポート = エージェント。圧縮()
# 圧縮レポート(
#decayed_facts=12,
# consolidation={'clusters_found': 3, 'facts_removed': 7},
# 矛盾={'pairs_found': 2, 'resolutions': 2},
# 事実の前=120、事実の後=102、
#)
3 つの独立したパス: 減衰 (未使用のファクトの指数関数的半減期)、統合 (埋め込みベースのクラスタリング、オプションの LLM シンセサイザー)、矛盾の調停 (キーと値の相違 + 極性検出、決定的勝者)。保護された事実は不滅です。
代理店。 link ( "Acme" 、 "owns" 、 "Globex" 、 confidentc

[切り捨てられた]

## Original Extract

Own your AI memory — import ChatGPT, Claude and Gemini exports, see what each AI knows about you. Checkpoint/restore and cross-model continuity for agents. - Thinklanceai/agentkeeper

GitHub - Thinklanceai/agentkeeper: Own your AI memory — import ChatGPT, Claude and Gemini exports, see what each AI knows about you. Checkpoint/restore and cross-model continuity for agents. · GitHub
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
Thinklanceai
/
agentkeeper
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
52 Commits 52 Commits .github/ workflows .github/ workflows agentkeeper agentkeeper benchmark benchmark docs docs examples examples marketing marketing tests tests .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
Cognitive continuity infrastructure for long-lived AI agents.
Your agent survives model switches, crashes, context-window limits, and restarts — with the same identity, memory, and priorities it had before.
Agents don't fail because they forget facts. They fail because they lose cognitive continuity — their state, priorities, and identity drift the moment the model changes, the context window fills, or the process restarts.
AgentKeeper treats this as a systems problem, not a memory problem.
pip install agentkeeper-ai
Zero required dependencies. No external infrastructure. Storage defaults to local SQLite.
pip install ' agentkeeper-ai[anthropic] ' # Claude
pip install ' agentkeeper-ai[openai] ' # GPT + OpenAI embeddings
pip install ' agentkeeper-ai[gemini] ' # Gemini
pip install ' agentkeeper-ai[semantic] ' # Local embeddings (sentence-transformers)
pip install ' agentkeeper-ai[mcp] ' # MCP server (Claude Desktop, Cursor, Codex)
pip install ' agentkeeper-ai[encrypted] ' # Encrypted storage at rest
pip install ' agentkeeper-ai[all] ' # Everything
Five things AgentKeeper does that nothing else does
1. Identity that survives everything
Principles and constraints are protected — exempt from every compression pass, injected into every reconstructed context, regardless of token budget. They survive decay, consolidation, contradiction arbitration, model switches, and process restarts.
import agentkeeper
agent = agentkeeper . create ( agent_id = "aria" , provider = "anthropic" )
agent . set_identity (
name = "Aria" ,
role = "EU insurance broker copilot" ,
principles = [ "never share PII without explicit consent" ],
constraints = [ "EU data residency only" ],
)
agent . principle ( "always confirm budget changes in writing" )
agent . fact ( "client: Acme Corporation" , importance = 0.95 )
agent . event ( "contract signed" , when = "2026-05-15" )
agent . save ()
# 100 compression cycles later — identity intact.
2. Same memory, switch models
The cognitive state is reconstructed in the format each model expects. XML for Claude, labelled sections for GPT-4, narrative prose for Gemini, terse tokens for Ollama. One agent, four runtimes, zero rewrites.
agent = agentkeeper . load ( "aria" , provider = "anthropic" )
response = agent . ask ( "What do we know about Acme?" )
agent . switch_provider ( "openai" ). save ()
response = agent . ask ( "Same question, different model." )
# Memory and identity are intact. Format has changed. Nothing broke.
3. TTL for GDPR — memory that expires itself
Facts and graph triples accept a TTL. When it lapses, purge_expired() removes them. No manual cleanup. Compliant by default.
agent . fact ( "session token: abc123" , ttl = "1h" )
agent . fact ( "audit log reference: Q1-2026" , ttl = "90d" )
agent . link ( "Acme" , "signed_contract" , "ThinkLanceAI" , ttl = "2y" )
agent . purge_expired () # removes what's lapsed, keeps what's protected
4. Graph traversal — structured relations alongside prose memory
Facts are prose. Triples are structure. Both live in the same agent, with their own retention and TTL policy.
agent . link ( "Acme" , "owns" , "Globex" )
agent . link ( "Globex" , "located_in" , "BE" )
agent . link ( "Alice" , "works_at" , "Acme" , confidence = 0.9 )
related = agent . find_related ( "Acme" , max_hops = 2 , direction = "out" )
# {"Globex": 1, "BE": 2, "Alice": 1}
5. Plug into Claude Desktop via MCP
AgentKeeper ships an MCP server. Any MCP-aware client — Claude Desktop, Cursor, Claude Code — gets full access to the agent's cognitive layer without writing a line of integration code.
agentkeeper-mcp --agent-id aria --provider anthropic
claude_desktop_config.json :
{
"mcpServers" : {
"aria" : {
"command" : " agentkeeper-mcp " ,
"args" : [ " --agent-id " , " aria " , " --provider " , " anthropic " ]
}
}
}
Available tools over MCP: add_fact , recall , set_identity , link , find_related , compress , health , gdpr_export , purge_expired , checkpoint , restore , list_checkpoints .
6. Checkpoints — snapshot, restore, survive a crash
Freeze the full cognitive state into an immutable, content-hashed
snapshot. Restore it after a crash, a context-window overflow, a model
switch, or a process restart. Attach an opaque execution_state payload
(current file, pending task, todos) — AgentKeeper stores and returns it
verbatim; it never interprets or runs it.
snap = agent . checkpoint (
label = "before refactor" ,
execution_state = {
"current_file" : "auth.py" ,
"pending_task" : "finish RS256 migration" ,
},
)
# crash, restart, switch model — the process memory is gone
agent = agentkeeper . load ( "aria" ). restore ( snap . snapshot_id )
# identity, facts, graph, and execution_state are back
agentkeeper . diff ( snap_a , snap_b ) # factual diff: facts added/removed/modified
agent . list_checkpoints () # every snapshot, oldest first
Reconstruction is deterministic — the same snapshot always rebuilds the
same cognitive state, verified by a content hash. Behaviour is not
guaranteed: AgentKeeper restores the agent's state, not the model's next
decision. The crash-recovery demo runs with no API key:
python examples/crash_recovery.py
7. Own your AI memory — import ChatGPT, Claude, and Gemini
Every AI platform remembers things about you. None of them show you
what, and none of them talk to each other. AgentKeeper imports the
official data exports of all three into one scoped, queryable memory —
then tells you exactly what each platform knows that the others don't.
import agentkeeper
from agentkeeper . importers import (
ChatGPTImporter , ClaudeImporter , GeminiImporter ,
compare_scopes , scope_inventory ,
)
agent = agentkeeper . create ( "my-memory" )
ChatGPTImporter ( "chatgpt-export.zip" ). apply ( agent )
ClaudeImporter ( "claude-export.zip" ). apply ( agent )
GeminiImporter ( "takeout.zip" ). apply ( agent )
print ( scope_inventory ( agent ))
# {"platform:chatgpt": {"EVENT": 142, "PREFERENCE": 9, ...},
# "platform:claude": {"EVENT": 87, "FACT": 4, ...}, ...}
result = compare_scopes ( agent , "platform:chatgpt" , "platform:claude" )
# ChatGPT only: ['Building a touring platform for electronic music']
# Claude only: ['Working on an EU AI Act observatory']
# Shared: [('My name is Tom, developer in Brussels', 1.0)]
Every imported fact is namespaced ( metadata["scope"] == "platform:chatgpt" ), so knowledge from one platform never leaks into
another's context unless you promote it. Matching is deterministic —
normalized content equality plus a difflib ratio, no embeddings, no
LLM calls, reproducible to the byte. PII (account email, full name)
stays out unless you pass include_pii=True . Archives are parsed in
memory with zip-bomb and path-traversal guards; nothing is ever
extracted to disk.
Also ships parse_memory_summary(text) for the prompt-generated memory
summaries used by the 2026 cross-platform migration flows.
┌──────────────────────────────────────────────────────────────┐
│ AgentKeeper Public API │
│ agent.remember() · agent.recall() · agent.ask() │
│ agent.compress() · agent.link() · agent.find_related() │
│ agent.set_identity() · agent.purge_expired() · agent.save() │
└────────────────────────────┬─────────────────────────────────┘
│
┌────────────────────────────▼─────────────────────────────────┐
│ Cognitive Reconstruction Engine (CRE) │
│ Identity injection · importance ranking · semantic boost │
│ Token budget · profile-driven rendering │
└─┬───────────┬────────────┬────────────┬──────────────────────┘
│ │ │ │
┌─────▼─────┐ ┌──▼─────────┐ ┌▼──────────┐ ┌▼──────────────┐
│ Memory │ │ Semantic │ │ Cognitive │ │ Cross-Model │
│ Hierarchy │ │ Recall │ │ Compress │ │ Translation │
│ │ │ │ │ │ │ │
│ working │ │ embeddings │ │ decay │ │ XML (Claude) │
│ episodic │ │ vector idx │ │ consol. │ │ sections (GPT)│
│ semantic │ │ sqlite-vec │ │ contradic │ │ narrative (G.)│
│ archival │ │ │ │ │ │ minimal (Oll.)│
└───────────┘ └────────────┘ └───────────┘ └───────────────┘
│ │
┌─────▼────────────────────────────────────────────▼────────────┐
│ Graph Layer Storage (pluggable) │
│ Triple, TTL, BFS SQLite · Encrypted SQLite · Postgres* │
│ agent.link() AGENTKEEPER_DB, AGENTKEEPER_ENC_KEY │
└───────────────────────────────────────────────────────────────┘
│
┌─────────────────────▼───────────────────┐
│ MCP Server · LangChain · CrewAI │
│ agentkeeper-mcp · langchain_system_prompt│
└──────────────────────────────────────────┘
*Postgres stub available; full implementation in v1.2.
agent . fact ( "budget: 50k EUR" , importance = 0.9 ) # stable semantic fact
agent . event ( "contract signed" , when = "2026-05-15" ) # episodic, time-anchored
agent . principle ( "never share PII" ) # protected, survives all compression
agent . remember ( "favourite colour: blue" ) # tier inferred automatically
Semantic recall
results = agent . recall ( "money allocated to the project" , top_k = 5 )
for fact , score in results :
print ( f" { score :.2f } { fact . content } " )
Pluggable backends: local sentence-transformers (default, free, offline), OpenAI, or your own. Persistent index via sqlite-vec — survives process restarts without rebuild.
report = agent . compress ()
# CompressionReport(
# decayed_facts=12,
# consolidation={'clusters_found': 3, 'facts_removed': 7},
# contradictions={'pairs_found': 2, 'resolutions': 2},
# facts_before=120, facts_after=102,
# )
Three independent passes: decay (exponential half-life on unused facts), consolidation (embedding-based clustering, optional LLM synthesiser), contradiction arbitration (key-value divergence + polarity detection, deterministic winner). Protected facts are immortal.
agent . link ( "Acme" , "owns" , "Globex" , confidenc

[truncated]
