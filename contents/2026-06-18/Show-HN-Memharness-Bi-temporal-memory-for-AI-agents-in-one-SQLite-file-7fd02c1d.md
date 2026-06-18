---
source: "https://github.com/las7/memharness"
hn_url: "https://news.ycombinator.com/item?id=48581267"
title: "Show HN: Memharness – Bi-temporal memory for AI agents, in one SQLite file"
article_title: "GitHub - las7/memharness: Bi-temporal agent long-term memory: SQLite-backed MCP server with recall ranking, hybrid vector+FTS recall, and a source-staleness signal · GitHub"
author: "sakuraiben"
captured_at: "2026-06-18T07:49:17Z"
capture_tool: "hn-digest"
hn_id: 48581267
score: 1
comments: 1
posted_at: "2026-06-18T05:43:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Memharness – Bi-temporal memory for AI agents, in one SQLite file

- HN: [48581267](https://news.ycombinator.com/item?id=48581267)
- Source: [github.com](https://github.com/las7/memharness)
- Score: 1
- Comments: 1
- Posted: 2026-06-18T05:43:48Z

## Translation

タイトル: Show HN: Memharness – 1 つの SQLite ファイルでの AI エージェントのバイテンポラル メモリ
記事のタイトル: GitHub - las7/memharness: バイテンポラル エージェントの長期記憶: リコール ランキング、ハイブリッド ベクター + FTS リコール、およびソースの古さシグナルを備えた SQLite ベースの MCP サーバー · GitHub
説明: バイテンポラル エージェント長期メモリ: リコール ランキング、ハイブリッド ベクター + FTS リコール、およびソースの古さシグナルを備えた SQLite ベースの MCP サーバー - las7/memharness

記事本文:
GitHub - las7/memharness: バイテンポラル エージェントの長期記憶: リコール ランキング、ハイブリッド ベクター + FTS リコール、およびソースの古さシグナルを備えた SQLite ベースの MCP サーバー · GitHub
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
ラス7
/
メムハーネス
公共
通知
Cha にサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
19 コミット 19 コミット .github/ workflows .github/ workflows アセット アセット サンプル サンプル パッケージ パッケージ .gitignore .gitignore .nvmrc .nvmrc BACKLOG.md BACKLOG.md DOGFOOD.md DOGFOOD.md EDGE-CASES.md EDGE-CASES.md LICENSE LICENSE PLAN.md PLAN.md README.md README.md STALENESS-SIGNAL.md STALENESS-SIGNAL.md biome.json biome.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント向けの、バイテンポラルで来歴を保持するメモリ プリミティブ。 1 つ
SQLite ファイル。ストレージ層では LLM またはネットワーク呼び出しはありません。あらゆるエージェントにさらされる
MCP経由。
ほとんどのエージェントのメモリは文字列の袋です。 memharness は事実を保存し、
既存企業が分割する傾向にある 3 つのセマンティクスを組み合わせています。
バイテンポラル : あらゆる事実は、世界でいつ真実になったかを記録します。
( valid_from / valid_to ) エージェントが学習したときとは別に
( tx_at )。そこで、「3 月 1 日に何を信じましたか?」と尋ねることができます。
置き換え、決して削除しない: 修正は古い事実を閉じ、それをリンクします
その後継者に。 「私が訂正する前はどう思っていましたか？」持っています
答えてください。
事実ごとの出所: すべての記憶は、誰が、どこで、いつ言ったかを引用します。
「なぜそれを信じるのですか？」答えがあります。 「すべてを忘れる」も同様です
そのセッション。」
ストレージ層は決定的です。LLM、ネットワーク、バックグラウンド デーモンはありません。
これはプレーンな SQLite なので、どのクライアントでもファイルを開くことができます。
自分で実行してください: cd サンプル && npm install && npm run デモ
これを使用する場合 (および使用しない場合)
memharness は魔法の精度のアップグレードではありません。それについては正直です。もしあなたの
エージェントのメモリは小さく静的であるため、

コンテキスト ウィンドウに快適にフィットします。
CLAUDE.md ファイル (または単にプロンプトに履歴を詰め込む) の方が簡単です。
短い履歴では、完全なコンテキストは、どの外部メモリ システムにも匹敵するか、それを上回ります。
歴史は窓を超えて広がります : 数ヶ月にわたる事実、多くの主題、あなた以上のもの
すべてのプロンプトに貼り付けたい (または貼り付けることができます)。
監査証跡が必要です。「エージェントは何を信じてこれを行ったのか」
decision?" ( as_of )、「月曜日から何が変わりましたか?」 ( diff ), "why does it
believe this?" （ なぜ ）。これらは、文字列のバッグでは答えられないクエリです。
来歴をスコープとした削除が必要です。「そこからのことはすべて忘れてください」
セッション/ファイル/ソース」を 1 回の呼び出しで実行します (文字列検索ではなく、GDPR 形式)。
信念は時間の経過とともに変化します。黙ってではなく、修正が優先されるべきです
上書きするので、古い推論が説明可能なままになります。
Honest, and pointed at the thing memharness actually does differently: it is a
抽出サービスではなく、決定的で監査可能なストレージ層です。
他が勝つのは明らかです: mem0 と Zep は自動ファクト抽出を行います
from raw conversation, which memharness deliberately does not (the write path
stays model-free;何を覚える価値があるかはクライアントまたはスキルによって決まります)。プレーン
CLAUDE.mdはインストールする必要がありません。 memharness は、必要なときにその場所を獲得します。
他のクエリでは提供されていない時間的および来歴のクエリ。
パッケージ
それは何ですか
@memharness/core
TypeScript ライブラリ: スキーマ、移行、書き込みパス、リコール ランキング。 No model, no network.
@memharness/mcp
MCP server (stdio) exposing the seven tools to any MCP client.
@memharness/embed
オプション。ハイブリッド（セマンティック）リコールのためのローカル埋め込みモデル。デフォルトではインストールされません。
Quick start (MCP)
The default install is small (SQLite plus the MCP SDK); the embedding model is
オプトインについては、「ハイブリッド リコール」を参照してください。
クロード mcp memharness を追加 -- npx -y @memharness/mcp
Clau

デスクトップ ( ~/Library/Application Support/Claude/claude_desktop_config.json )
とカーソル ( ~/.cursor/mcp.json ) は同じ JSON 形状を使用します。
{
"mcpサーバー": {
"memharness" : { "コマンド" : " npx " , "args" : [ " -y " , " @memharness/mcp " ] }
}
}
Codex ( ~/.codex/config.toml ) は、JSON ではなく TOML を使用します。
[ mcp_servers .メムハーネス]
コマンド = " npx "
args = [ " -y " , " @memharness/mcp " ]
データベースは ~/.memharness/memory.db にあります ( MEMHARNESS_DB でオーバーライドします。
Linux では XDG_DATA_HOME が受け入れられます)。オンにしないと何も書き込まれません
オプションのデバッグ ログ。
上記のコマンドのいずれかを使用してサーバーを追加し、クライアントを再起動します。
したがって、新しい MCP サーバーを選択します。
会話の中で、永続的な事実をエージェントに渡します。 「私がいることを思い出してください
このプロジェクトを Fly.io でデプロイします。」これは remember を呼び出します。
後で (または新しいセッションで) 「デプロイ方法について何を知っていますか?」と尋ねます。それ
呼び出しは記憶から呼び出され、応答は記憶から行われます。それを修正すると、 Revise が呼び出されます。
古い信念は歴史となり、 as_of / Why / diff でクエリ可能になります。
API キーもサインアップもネットワークもありません。最初の記憶では SQLite ファイルが作成されます
それがセットアップ全体です。エージェントなしでツールがエンドツーエンドで動作するのを確認するには、
デモを実行します: cd example && npm install && npm run Demon 。
オプション: リコールを自動化します
デフォルトでは、エージェントがいつリコールを呼び出すかを決定します。関連するメモリを押し込むには
代わりに、すべてのセッションの開始時に (モデルを期待するよりも信頼性が高くなります)
忘れずに見てください)、を実行する Claude Code SessionStart フックを追加します。
バンドルされた memharness-context ツール。標準出力がコンテキストに挿入されます。
{
「フック」: {
"セッション開始" : [
{ "フック" : [ { "タイプ" : " コマンド " ,
"コマンド" : " npx -y -p @memharness/mcp memharness-context --subject user " } ] }
】
}
}
最も関連性のある現在の信念のコンパクトなダンプを出力します (そして静かに終了します)
もし

まだ何もありません)。そのため、エージェントはすでに
永続的な事実。複数のエンティティを挿入するには、 --subject を複数回渡します。
ツール
何をするのか
テストする論文
覚えておいてください
原子的な事実を信頼と来歴とともに保存する
事実 > ブロブ
思い出す
現在の信念をランク付けします。 as_of は過去の瞬間の信念を返します
バイテンポラル
修正する
信念に取って代わり、歴史を残す
置き換え > 削除
差分
日付以降に変更された内容 (学習/修正/撤回)
監査デモ
なぜ
ファクトの出典 + 完全なリビジョン チェーン
信頼/監査
忘れる
ID またはソースごとの墓石 (出所ベースの削除)
GDPR 型
統計
カウント、サブジェクト、スキーマのバージョン
—
図書館の利用
"@memharness/core" から {Memharness } をインポートします。
const mem = メムハーネス 。開ける （ ） ; // ~/.memharness/memory.db
// 今何かを学び、その後、それが以前は実際に真実であったことを学びます。
const {id} = mem 。覚えておいてください ( {
件名 : "ユーザー" 、
事実：「大阪に住んでいます」、
ソース参照: "セッション-2026-06-09" 、
} ) ;
メモリ。 Revise ( { oldFactId : id , newFact : "東京在住" , validFrom : "2026-05-01" } ) ;
メモリ。呼び出し ( { クエリ : "lives" } ) 。事実 [ 0 ] 。事実 ; // 「東京に住んでいます」（現在の信念）
メモリ。 diff ( { 以来 : "2026-06-01" } ) ; // { 学習、修正、撤回 }
メモリ。なぜ (id) ; // { 事実、祖先、子孫 }
remember は、裸の文字列ではなく、RecallResult ( {actuals: ScoredFact[]; asOf; truncated; usedFallback } ) を返します。 asOf time-travels: mem.recall({ query: "lives", asOf: "2026-04-15" }) は、その日付に開催されたと信じられていたものを返します。
これはトランザクション時間を考慮するため、今日学んだ事実はクエリには表示されません。
過去について。
リコール ランキングは、FTS5 BM25 に対する相互ランクの融合です (さらに、
ハイブリッド リコールが有効になっている場合)、信頼度の回数、回数
リーセンシー減衰 (90 日の半減期、構成可能)、sco

SQLでは赤。オプションの
maxTokens のバジェットは、コンテキスト ウィンドウの出力を制限します。部分文字列フォールバックがキャッチする
FTS 専用モードとハイブリッド モードの両方で、部分的な単語とタイプミス。
デフォルトでは、再現率は FTS5 キーワード検索に最新性/信頼性ランキングを加えたものです: いいえ
モデル、完全にオフライン。ハイブリッド リコールは、ローカル経由でセマンティック レッグを追加します。
埋め込みモデル (BGE-small、約 130MB、HuggingFace ハブから 1 回ダウンロード)
その後完全にオフラインになります: API キーもクエリごとのネットワークもありません)。 2 つの手順で有効にします。
オプションの埋め込みパッケージをサーバーと一緒にインストールします。 npx の場合:
npx -y -p @memharness/mcp -p @memharness/embed memharness-mcp
(または、グローバル インストールの場合は npm i -g @memharness/embed)。
サーバーの環境で MEMHARNESS_HYBRID=1 を設定します。
その後、サーバーは保存されたファクトを自動的に埋め込み続けます。つまり、ユーザーが覚えているファクトです。
次回のリコール時に、個別のバックフィルなしで意味的に検索可能になります。
ステップ。最初のハイブリッド リコールでは、モデルの実行中にダウンロードの進行状況が標準エラー出力に出力されます。
負荷がかかります。パッケージがインストールされていない場合、サーバーはその旨を通知し、FTS のみのままになります。それ
必ず閉じてください。
ライブラリ レベルでは、リコールは埋め込みプロバイダーに依存しません。独自のクエリを渡します。
呼び出してドキュメント ベクトルを添付するベクトル ({ queryVector })
setEmbedding(...) 、任意のモデルから。
数週間離れた 2 つのセッション。エージェントは好みを学習し、ユーザーは後で学習します
それを修正し、下流の質問でエージェントがその時点で何を信じたかを尋ねます。
時間：
// 6 月 9 日: エージェントはデプロイ ターゲットを学習し、それに基づいて動作します。
const {id} = mem 。覚えておいてください ( {
件名 : "プロジェクト:acme" 、
事実: "Heraku 経由でデプロイ" ,
ソース参照: "セッション-2026-06-09" 、
} ) ;
// 6 月 16 日: チームが 6 月 1 日にフライバックに移籍したことが判明。
メモリ。修正します ( {
oldFactId : id 、
newFact : "Fly.io 経由でデプロイ" ,
validFrom : "2026-06-01" ,
ソース参照: "セッション-2026-06-16" 、
} ) ;
メモリ。思い出してください ( { サブジェクト

ect: "プロジェクト:acme" } ) 。事実 [ 0 ] 。事実 ; // "Fly.io 経由でデプロイ"
// 「6 月 9 日に作成した CI 構成が Heroku をターゲットにしたのはなぜですか?」
メモリ。 remember ({ subject : "project:acme" , asOf : "2026-06-09" } ) 。事実 [ 0 ] 。事実 ;
// 「Heraku 経由でデプロイ」: その日エージェントが正直に信じたこと。
メモリ。なぜ (id) ; // チェーン全体: Heroku (Fly.io に取って代わられる)、ソース付き。
メモリ。 diff ( { 以来 : "2026-06-15" } ) ; // Heroku -> Fly.io リビジョンを表示します。
バッグオブストリングメモリは上書きされているため、as_of の質問に答えることはできません。
Heroku が Fly.io を学習した瞬間。
プロパティ スイートはプロジェクトの中心です。ランダム化された一連のプロパティです。
覚えている/修正している/忘れている、recall({asOf: T}) は生成された信念セットと等しくなければなりません
イベントごとに調査される、SQL を使用しない単純なイベント ログの再生による
タイムスタンプ ±1ms。メインにプッシュされるたびに 10,000 件のケースが実行されます。
開発者で 100,000 のファクト (10% の改訂チェーン、2% の撤回) でベンチマークを実施
ラップトップ (Apple Silicon): 全体的なリコールは 10 ミリ秒の予算に対して p95 ～ 1.3 ミリ秒、
4 つのクエリ形式 (2 つの用語のキーワード、キーワード + 件名、件名のみ、および
as_of + キーワード)。 pnpm ベンチはデータベースをシードし、バジェットをアサートします。
数値は引用ではなく再現可能です。
実行者1名

[切り捨てられた]

## Original Extract

Bi-temporal agent long-term memory: SQLite-backed MCP server with recall ranking, hybrid vector+FTS recall, and a source-staleness signal - las7/memharness

GitHub - las7/memharness: Bi-temporal agent long-term memory: SQLite-backed MCP server with recall ranking, hybrid vector+FTS recall, and a source-staleness signal · GitHub
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
las7
/
memharness
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
19 Commits 19 Commits .github/ workflows .github/ workflows assets assets examples examples packages packages .gitignore .gitignore .nvmrc .nvmrc BACKLOG.md BACKLOG.md DOGFOOD.md DOGFOOD.md EDGE-CASES.md EDGE-CASES.md LICENSE LICENSE PLAN.md PLAN.md README.md README.md STALENESS-SIGNAL.md STALENESS-SIGNAL.md biome.json biome.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
A bi-temporal, provenance-carrying memory primitive for AI agents. One
SQLite file. No LLM or network calls in the storage layer. Exposed to any agent
via MCP.
Most agent memory is a bag of strings. memharness stores facts , and
combines three semantics that incumbents tend to split apart:
Bi-temporal : every fact records when it became true in the world
( valid_from / valid_to ) separately from when the agent learned it
( tx_at ). So you can ask: "what did you believe on March 1st?"
Supersession, never deletion : corrections close the old fact and link it
to its successor. "What did you think before I corrected you?" has an
answer.
Provenance per fact : every memory cites who said it, where, and when.
"Why do you believe that?" has an answer. So does "forget everything from
that session."
The storage layer is deterministic: no LLM, no network, no background daemon.
It's plain SQLite, so you can open the file with any client.
Run it yourself: cd examples && npm install && npm run demo
When to use this (and when not to)
memharness is not a magic accuracy upgrade, and it is honest about that. If your
agent's memory is small and static and comfortably fits the context window, a
CLAUDE.md file (or just stuffing the history into the prompt) is simpler, and
on short histories full context will match or beat any external memory system.
History outgrows the window : months of facts, many subjects, more than you
want to (or can) paste into every prompt.
You need an audit trail : "what did the agent believe when it made this
decision?" ( as_of ), "what changed since Monday?" ( diff ), "why does it
believe this?" ( why ). These are queries a bag of strings cannot answer.
You need provenance-scoped deletion : "forget everything from that
session/file/source" in one call (GDPR-shaped, not a string search).
Beliefs change over time : corrections should supersede, not silently
overwrite, so old reasoning stays explainable.
Honest, and pointed at the thing memharness actually does differently: it is a
deterministic, auditable storage layer rather than an extraction service.
Where the others win, plainly: mem0 and Zep do automatic fact extraction
from raw conversation, which memharness deliberately does not (the write path
stays model-free; a client or skill decides what is worth remembering). Plain
CLAUDE.md needs no install at all. memharness earns its place when you need the
temporal and provenance queries the others don't offer.
Package
What it is
@memharness/core
TypeScript library: schema, migrations, write path, recall ranking. No model, no network.
@memharness/mcp
MCP server (stdio) exposing the seven tools to any MCP client.
@memharness/embed
Optional. A local embedding model for hybrid (semantic) recall. Not installed by default.
Quick start (MCP)
The default install is small (SQLite plus the MCP SDK); the embedding model is
opt-in, see Hybrid recall .
claude mcp add memharness -- npx -y @memharness/mcp
Claude Desktop ( ~/Library/Application Support/Claude/claude_desktop_config.json )
and Cursor ( ~/.cursor/mcp.json ) use the same JSON shape:
{
"mcpServers" : {
"memharness" : { "command" : " npx " , "args" : [ " -y " , " @memharness/mcp " ] }
}
}
Codex ( ~/.codex/config.toml ) uses TOML, not JSON:
[ mcp_servers . memharness ]
command = " npx "
args = [ " -y " , " @memharness/mcp " ]
The database lives at ~/.memharness/memory.db (override with MEMHARNESS_DB ;
XDG_DATA_HOME is honored on Linux). Nothing else is written unless you turn on
the optional debug log .
Add the server with one of the commands above, then restart your client
so it picks up the new MCP server.
In a conversation, hand the agent a durable fact, e.g. "remember that I
deploy this project with Fly.io." It calls remember .
Later (or in a fresh session) ask "what do you know about how I deploy?" It
calls recall and answers from memory. Correct it and it calls revise ;
the old belief becomes history, queryable with as_of / why / diff .
No API key, no signup, no network. The first remember creates the SQLite file
and that's the whole setup. To watch the tools work end to end without an agent,
run the demo: cd examples && npm install && npm run demo .
Optional: make recall automatic
By default the agent decides when to call recall . To push relevant memory in
at the start of every session instead (more reliable than hoping the model
remembers to look), add a Claude Code SessionStart hook that runs the
bundled memharness-context tool, whose stdout is injected into context:
{
"hooks" : {
"SessionStart" : [
{ "hooks" : [ { "type" : " command " ,
"command" : " npx -y -p @memharness/mcp memharness-context --subject user " } ] }
]
}
}
It prints a compact dump of the most relevant current beliefs (and exits quietly
if there's nothing yet), so the agent starts each session already knowing the
durable facts. Pass --subject more than once to inject several entities.
Tool
What it does
The thesis it tests
remember
store an atomic fact with confidence + provenance
facts > blobs
recall
ranked current beliefs; as_of returns beliefs at a past instant
bi-temporal
revise
supersede a belief, keep history
supersession > deletion
diff
what changed since a date (learned/revised/retracted)
the audit demo
why
provenance + full revision chain for a fact
trust / audit
forget
tombstone by id or by source (provenance-based deletion)
GDPR-shaped
stats
counts, subjects, schema version
—
Library use
import { Memharness } from "@memharness/core" ;
const mem = Memharness . open ( ) ; // ~/.memharness/memory.db
// Learn something now, then learn it was actually true earlier.
const { id } = mem . remember ( {
subject : "user" ,
fact : "lives in Osaka" ,
sourceRef : "session-2026-06-09" ,
} ) ;
mem . revise ( { oldFactId : id , newFact : "lives in Tokyo" , validFrom : "2026-05-01" } ) ;
mem . recall ( { query : "lives" } ) . facts [ 0 ] . fact ; // "lives in Tokyo" (current belief)
mem . diff ( { since : "2026-06-01" } ) ; // { learned, revised, retracted }
mem . why ( id ) ; // { fact, ancestors, descendants }
recall returns a RecallResult ( { facts: ScoredFact[]; asOf; truncated; usedFallback } ), not a bare string. asOf time-travels: mem.recall({ query: "lives", asOf: "2026-04-15" }) returns what was believed as held on that date .
That honors transaction time, so a fact learned today is not visible to a query
about the past.
Recall ranking is reciprocal-rank fusion over FTS5 BM25 (plus a vector rank when
hybrid recall is enabled), times confidence, times
recency decay (90-day half-life, configurable), scored in SQL. An optional
maxTokens budget caps output for context windows. A substring fallback catches
partial words and typos, in both FTS-only and hybrid modes.
By default, recall is FTS5 keyword search plus recency/confidence ranking: no
model, fully offline. Hybrid recall adds a semantic leg via a local
embedding model (BGE-small, ~130MB, downloaded once from the HuggingFace hub
then fully offline: no API key, no per-query network). Enable it in two steps:
Install the optional embedding package alongside the server. With npx :
npx -y -p @memharness/mcp -p @memharness/embed memharness-mcp
(or npm i -g @memharness/embed for a global install).
Set MEMHARNESS_HYBRID=1 in the server's environment.
The server then keeps stored facts embedded automatically: facts you remember
become semantically searchable on the next recall , with no separate backfill
step. The first hybrid recall prints download progress to stderr while the model
loads. If the package isn't installed, the server says so and stays FTS-only; it
never fails closed.
At the library level, recall is embedding-provider-agnostic: pass your own query
vector to recall({ queryVector }) and attach document vectors with
setEmbedding(...) , from any model you like.
Two sessions, weeks apart. The agent learns a preference, the user later
corrects it, and a downstream question asks what the agent believed at the
time :
// June 9: the agent learns a deploy target and acts on it.
const { id } = mem . remember ( {
subject : "project:acme" ,
fact : "deploys via Heroku" ,
sourceRef : "session-2026-06-09" ,
} ) ;
// June 16: turns out the team moved to Fly back on June 1.
mem . revise ( {
oldFactId : id ,
newFact : "deploys via Fly.io" ,
validFrom : "2026-06-01" ,
sourceRef : "session-2026-06-16" ,
} ) ;
mem . recall ( { subject : "project:acme" } ) . facts [ 0 ] . fact ; // "deploys via Fly.io"
// "Why did the CI config you wrote on June 9 target Heroku?"
mem . recall ( { subject : "project:acme" , asOf : "2026-06-09" } ) . facts [ 0 ] . fact ;
// "deploys via Heroku": what the agent honestly believed that day.
mem . why ( id ) ; // the full chain: Heroku, superseded by Fly.io, with sources.
mem . diff ( { since : "2026-06-15" } ) ; // surfaces the Heroku -> Fly.io revision.
No bag-of-strings memory can answer the as_of question, because it overwrote
Heroku the moment it learned Fly.io.
The property suite is the heart of the project: for randomized sequences of
remember/revise/forget, recall({asOf: T}) must equal the belief set produced
by a naive, SQL-free replay of the event log, probed at every event
timestamp ±1ms. 10,000 cases run on every push to main.
Benchmarked at 100k facts (10% revision chains, 2% retractions) on a developer
laptop (Apple Silicon): overall recall p95 ~1.3ms against a 10ms budget,
across four query shapes (two-term keyword, keyword + subject, subject-only, and
as_of + keyword). pnpm bench seeds the database and asserts the budget, so
the number is reproducible rather than quoted.
One deliber

[truncated]
