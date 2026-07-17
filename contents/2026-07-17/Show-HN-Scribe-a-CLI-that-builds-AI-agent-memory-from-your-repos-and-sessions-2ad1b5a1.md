---
source: "https://getscribe.dev/"
hn_url: "https://news.ycombinator.com/item?id=48943417"
title: "Show HN: Scribe, a CLI that builds AI agent memory from your repos and sessions"
article_title: "scribe — your knowledge base, written by your tools"
author: "quatermain"
captured_at: "2026-07-17T04:55:08Z"
capture_tool: "hn-digest"
hn_id: 48943417
score: 1
comments: 0
posted_at: "2026-07-17T04:44:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Scribe, a CLI that builds AI agent memory from your repos and sessions

- HN: [48943417](https://news.ycombinator.com/item?id=48943417)
- Source: [getscribe.dev](https://getscribe.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T04:44:16Z

## Translation

タイトル: Show HN: Scribe、リポジトリとセッションから AI エージェント メモリを構築する CLI
記事のタイトル: scribe — ツールによって書かれた知識ベース
説明: git リポジトリ、Claude Code および Codex セッション、および自己送信リンクを、厳選された意味的に検索可能なナレッジ ベースに変える単一バイナリ CLI。クロスプロジェクト、cron 駆動、完全にローカル対応。

記事本文:
コンテンツにスキップ
s
筆記者
仕組み
特長
比較する
チーム
コスト
CLI
ループ
よくある質問
GitHub ↗
インストール
GitHub
仕組み
特長
比較する
チーム
コスト
CLI
ループ
よくある質問
GitHub
ダークモードに切り替える
ライトモードに切り替える
スクライブをインストールする
s
単一の Go バイナリ · cron 上で実行 · 自己ホスト型
ツールによって作成されたナレッジ ベース。
scribe は、あなたの Git 履歴、Claude Code および Codex セッション、および自分で送信した URL を読み取り、Wiki を作成します。そのため、次のエージェント セッションでは、あなたが何を決定したか、そしてその理由をすでに知っています。これは、エージェントが行動する前に読み取るメモリであり、維持して再度開くことのない第二の頭脳ではありません。Git でのプレーンなマークダウン、プロジェクト間、cron 駆動で、API 支出ゼロで Ollama 上で 100% ローカルで実行できます。
GitHub で見る
7,472 件の文書
メンテナの KB · 手動で入力されたゼロ
$0 /同期
100% Ollama パスで検証済み
～70秒
毎週の夢のサイクル、gemma3:12b
~/projects/scriptorium — zsh
仕組み
3 つのステージ、1 つのパイプライン。
scribe は 4 つの入力ストリームをマイニングし、LLM がノイズに触れる前にノイズをフィルターで除去し、密度の高いソースをエンティティファーストの Wiki ページにファンニングします。すべてのステップは cron で実行されます。一度設定したら忘れてしまいます。
git
セッション
URL
ファイルをドロップする
BM25
トリアージ
ボイラープレート · 拒否されました · $0
吸収する
パス1→2
qmd インデックス
エージェントが読み取ります
クロード.md · エージェント.md
1 キャプチャー
4 つのストリーム、すべて cron 上にあります。
Git リポジトリ、Claude Code および Codex セッション、自分でテキスト メッセージを送信した URL、および他のプロジェクトからのファイルのドロップ。 scribe は、どちらかの CLI で開いたすべてのコードベースを自動検出し、マニフェストを最新の状態に保ちます。
キーワード密度スコアリングでは、LLM 呼び出しの前に定型セッションが拒否されるため、安価なセッションには費用がかかりません。生存者は 2 つのパス吸収を通過します: 1 つの根拠となる原子ファクトをパスし、2 つのファンの高密度ソースを複数のエンティティファースト Wiki ページにパスします。
プレーンなマークダウンの型付きグラフ。
自動生成されたウィキリンク

s、バックリンク JSON、および検索コンテキストの段落がすべての記事に接続されているため、埋め込みが暗黙的なエンティティをキャッチします。 qmd はセマンティック検索のためにインデックスを再作成します。これは、任意の端末、任意のディレクトリ、または MCP を介してクロード コード内からアクセスできます。
scribe は RAG ではなく、Obsidian でもなく、別の LLM-on-every-session バーナーでもありません。それはそれらの間に位置し、あなたの仕事を監視し、あなたのためにメモを書き、あなたが触れるすべてのプロジェクトにわたる知識を統合します。
scribe init はハンドシェイク ブロックを ~/.claude/CLAUDE.md と ~/.codex/AGENTS.md の両方に書き込むため、すべてのプロジェクトのすべてのセッションで、ライブラリを推奨したりアーキテクチャを提案したりする前に KB がクエリされます。
時間ごとの自動コミット。 2 時間ごと: プロジェクトの抽出。 1 日あたり 3 回: セッションマイニング。 30 分ごと: キューに登録された URL。 4 時間ごと: セルフ iMessage リンク。日曜日の 02:00: Dream の完全な統合と、その間に毎日軽いホット ドメイン パスが含まれます。
プロジェクト間で 1 つの KB があり、リポジトリごとにサイロ化されていません。月曜日にプロジェクト A の oban 冪等性バグを解決します。金曜日に同じ形状がプロジェクト B に現れたとき、エージェントは修正を見つけます。
プロジェクトごとの抽出、2 パス吸収、ドリーム サイクル、評価、ディープ、セッション マイニング、関係移行: すべての LLM 演算は、ローカルの Ollama サーバーに対してエンドツーエンドで実行されます。 scribe.yaml の 1 行でパイプライン全体が反転されます。 API 支出ゼロ。
アプリよりファイル: コーパスはパイプラインよりも長く存続する必要があります。 YAML フロントマターを使用したプレーンなマークダウンの git リポジトリ。自分の GitHub、Gitea、または Forgejo にプッシュします。 Obsidian、VS Code、vim、または mdbook で開きます。 SaaS もベンダー ロックインもありません。
記事は、閉じた 10 種類の型付きエッジ スキーマ ( supersedes 、contracts 、derived_from 、specialized 、 extends 、およびその他 5 つ) を介して接続されます。 scribe 関係は、既存の関連を分類し、LLM を使用してそれにリンクします。
wiki/決定/ecto-multi.md
筆記者によって書かれました · によってタイプされたものではありません

手
---
タイプ: 決定
タグ: [ecto、postgres、マルチテナンシー]
置き換えられるもの: [[ecto-transaction]]
---
# テナントごとのスキーマよりも行レベルのテナント
スキーマ上で Postgres RLS を使用して tenant_id 列を選択しました
テナントごと: 移行はシングルパスとプールのままです
テナントごとにファンアウトしません。
理由: テナントごとのスキーマが OBAN ジョブ ルーティングを壊し、
アドバイザリーは、負荷がかかるとテーブルを曖昧にロックします。
## も参照
- [[oban-idempotency]] · 重複排除キーには tenant_id が含まれます
- [[genserver-backpressure]] · テナントごとのレート分離
実行中のスクライブではありません - スクライブの出力。 git に書き込まれたページ: YAML フロントマター、散文、および入力された [[wikilinks]] 。これらのリンクは、以下のグラフが描くエッジです。
に取って代わる
派生元
矛盾している
専門分野
伸びる
派生元
外部取引
オーバン冪等性
アドバイザリーロック
トークンバケット
生成サーバーのバックプレッシャー
ライブビューストリーム
エクトマルチ
実際のスクライブ KB のスライス: フォルダーやタグではなく、置き換えや矛盾などの入力されたエッジによってリンクされたエンティティにより、エージェントは単にキーワードを照合するだけでなく、決定が行われた理由を追跡できます。
4つの入力ストリーム
45 のサブコマンド
10種類の型付きエッジ
3つの推論モード
7つのプロバイダー
2パス吸収
0 ベクトル DB
必要な API キーは 0 個です
1 つの自己完結型バイナリ
表面全体を数値で示します。すべての数値は、マーケティングではなく、 scribe --help および scribe.yaml に対してチェック可能です。ローカルモードではキーはまったく必要ありません。デフォルトの Anthropic パスは、API キーではなく、claude -p CLI を介してサインインします。
ラグではありません。黒曜石ではありません。セッションごとに LLM を実行する別のバーナーではありません。
scribe は、ベクトル データベースではなく、コンパイルされたナレッジ ベースです。エージェントが BM25 でクエリを実行する、厳選されたマークダウン Wiki を自動で書き込みます。そのため、埋め込むものやホストするものは何もありません。 「第二の脳」に関する議論は、読んだメモに関するものです。書記はそうではない。エージェントが決定する前に読むのは記憶です: 背後にある理由

二度と開くことのない要約ではなく、選択です。これは、手動メモ ツール (Obsidian、Notion) と無制限の LLM-on-every-query アプローチ (バニラ RAG、claude-memory-compiler) の間に位置します。生のソースの上にある厳選された Wiki は逐語的に保持され、エージェントが全体を読み取るのに十分な大きさで、ほとんどのルックアップはベクトルの推測ではなくプレーンテキストの一致であるため、クエリが安価です。
↔ 表をスワイプしてすべてのツール列を表示します
これらの中で、ベクター データベースを使用せずに、エージェントが判断する前に読む、ポータブルな git バージョンのマークダウン Wiki を自動作成するのは scribe だけです。 AnythingLLM とは異なり、scribe はプレーンなマークダウンを git に保存し、ベクター データベースや実行中のサーバーを必要としません。スナップショット 2026 年 6 月 8 日。ツールの機能は変更されるため、決定する前に各プロジェクトのドキュメントを確認してください。
2026 年 6 月 8 日のスナップショット。すべては動きます。決定する前にソース リポジトリを確認してください。判定は、公開されている README、問題トラッカー、および scribe メンテナのツール評価から取得されます。
scribe は、LLM があなたのために書く wiki で、そのままの状態で保存された生のソースの上にあります。 RAG はチャンクを取得します。 scribe は、grep も可能な厳選された名前付きエンティティ Wiki を提供します。
— プロジェクトの README
1 KB。チーム全体がそれに書き込みます。
scribe はデフォルトでシングルユーザーであり、そのままです。ただし、同じコードベースを使用する小規模なチームは、すべてのマシンに 1 つの git-backed KB をポイントさせることができます。エージェント セッションの共有に関する明らかな恐怖 (秘密の漏洩、プライベート クライアント リポジトリ、1 人のチームメイトの設定変更が全員に及ぶこと) は、まさに以下のゲートが阻止するために構築されたものです。バツ印を共有 KB に保存するつもりだった知識のみ。
チームの共有KB
git · チームメイトが見るもの
決断
パターン
ソリューション
越えようとするもの
それを止める門
シークレットスキャンゲート -> 開催 -->
トランスクリプト内の AWS キー
シークレットスキャンゲート
保留 — 決してコミットしない
リモート許可リスト -> 保持 -->

チームメイトのクライアント リポジトリ
リモート許可リスト
承認なしに摂取することはありません
信頼層 -> 保持 -->
LLM を再ポイントする構成
トラストレイヤー
最後の信頼されたスナップショットに戻りました
厳選された抜粋 -> クロス -->
修正の背後にある理由
厳選されたエキス
✓
昇進・来歴保持
機密性の高いマテリアルはコミット ゲートで停止します。意図的に保持している知識のみが共有 KB に追加されます。ここでのすべてのゲートはコード内のメカニズムであり、ドキュメント内のポリシーではありません。
共有設定はデフォルトでは信頼できません
トラスト層は、共有 scribe.yaml の機密面 (プロバイダー、モデル、取り込みパス、およびシークレット スキャナー自体) を固定します。推論を新しいエンドポイントに再ポイントしたり、新しいディレクトリを KB にドレインしたりするプッシュされた変更は、人間が承認するまで、最後に信頼されたスナップショットに戻ります。
秘密は共有履歴に決して届かない
scribe は、定期的に API キーとトークンを運ぶセッション トランスクリプトをマイニングします。チーム モードでは、確定的シークレット スキャナーがコミット ゲートで実行され、共有 Git 履歴に何かが記録される前に、フラグが設定された認証情報が保持されます。 LLM なし、ネットワークなし: コミット パス上の正規表現パス。
抽出料金はラップトップごとではなく、1 回ごとに支払われます
すべてのスクライブ同期では抽出前にプル、マージ、インデックスの再作成が行われ、コミットされた台帳によって 2 台のマシンが同じ git リビジョンを 2 回マイニングすることがなくなります。推論の請求額は、KB に指定されているラップトップの数ではなく、コミットに応じて増減します。
チームメイトの無関係なリポジトリが漏洩することはありません
発見されたプロジェクトは保留状態になります。 allowed_remotes とソース フィルターは、git-remote ID によって検出をゲートし、パイプラインに入る内容をスクライブ プロジェクトの {list,approve,ignore,review} で制御するため、サイド プロジェクトやクライアントのチェックアウトが承認なしに共有 KB に到達することはありません。
非公開でキュレーションし、意図的に宣伝する
scribe プロモーション <記事> --チーム KB に

個人の KB から共有の KB にページをコピーし、出所を記録します。派生ファイルや調整ファイルはソースとして拒否されるため、チームの KB には、作業中のスクラッチではなく、公開しようとしていた内容が書き込まれます。
サーバーなしで 1 台のマシンを統合
毎週の Dream 統合は Wiki 全体を書き換え、マージし、整理するため、正確に 1 台のマシンで実行する必要があります。リポジトリ内でコミットされたリーダー リースがそのマシンを選択します。etcd もロック サーバーもなし、日曜日の 02:00 に 2 台のラップトップが同じ wiki を書き換えるために競合することはありません。
ローカル、ホスト、または人為的、そしてあなたが読める請求書。
ローカルの Ollama でパイプライン全体を $0 で実行し、ラップトップが熱いときにホストされた OpenAI 互換 API (Together、Groq、Fireworks、Hugging Face) をポイントするか、Anthropic パスの claude -p を維持します。この 3 つは共存します。 scribe.yaml で操作ごとに取引を選択すると、明示的な設定がなければ有料プロバイダーには何も届きません。次に、スクライブ コストは、プロバイダーおよび KB にわたるすべてのトークンを、プロバイダー独自のダッシュボードに対してセント単位で調整します。
~ ドルの筆記コスト
スクライブコスト — 過去 7 日間 — 2 KB (個人 KB、チーム KB)
モデル呼び出し OK キャンセルレート tmout ウォールクロック イントークン アウトトークン usd
─────────────────── ─────────
ソネット 301 281 0 0 3 6 時間 54 分 250,468 789,357 $92.80
俳句 151 144 0 0 7 2時間21分 1,251 551,761 $10.77
一緒/MiniMaxAI/MiniMax-M3 195 195 0 0 0 1時間06分 1,261,197 146,728 $0.55
オラマ/ジェマ3:12b 120 120 0 0 0 4時間57分 1,037,866 138,258 —
オラマ/qwen3:30b-a3b-instruct-2507

-q4_K_M 1,643 1,627 0 0 0 23時間39分 11,317,941 1,122,293 —
オラマ/ジェマ3:4b 921 912 0 0 0 8時間23分 2,355,608 672,157 —
─────────────────── ─────────
合計 3,331 3,279 0 0 10 47時間22分 16,224,331 3,420,554 $104.12
プロバイダー別:
プロバイダーは、in-tokens out-tokens usd を呼び出します
─------------------------------------------------------------------------------------
人間 452 251,719 1,341,118 $103.57
一緒に 195 1,261,197 146,728 $0.55
オラマ 2,684 14,711,415 1,932,708 —
KB による:
KB は、イントークン、アウトトークン、米ドルを呼び出します
─────────────────
チーム KB 1,732 7,020,329 2,211,137 $102.96
個人用KB 1,599 9,204,002 1,209,417 $1.16
カバレッジ: 3175/3331 コールにはトークン データが含まれていました。他の 156 の追加料金は約 0.62 ドルと推定されますが、上には示されていません。
usd = プロバイダーに請求される費用 (料金を含む)キャッシュ書き込みとキャッシュ読み取り。 in はキャッシュされていない入力であるため、usd は in/out リストの推定値を超えています。
cancl = 兄弟キャンセル (レート制限カスケード)。 tmout = ctx.DeadlineExceeded。
↔ 表をスワイプすると、モデルごとの完全な内訳が表示されます
実際の出力、名前は匿名化されます。同じマシン、同じ週、3 つのプロバイダー パス: Anthropic の請求 $103.57 、ホストされている Together パス $0.55 、ローカルの Ollama モデルは何もありません。パスは異なるボリュームを処理したため (ルーティングは scribe.yaml の操作ごとに行われます)、それぞれのパスを、同等のものではなく、実際に請求されたものとして読み取ってください。

ケレース。通常、1 つの API キーはマシン上の KB ごとに請求されるため、scribe コストはデフォルトで集計され、ダッシュボードに反映されます。
ローカル パスには claude -p ca はありません

[切り捨てられた]

## Original Extract

A single-binary CLI that turns your git repos, Claude Code and Codex sessions, and self-sent links into a curated, semantically searchable knowledge base. Cross-project, cron-driven, fully local-capable.

Skip to content
s
scribe
How it works
Features
Compare
Teams
Cost
CLI
Loop
FAQ
GitHub ↗
Install
GitHub
How it works
Features
Compare
Teams
Cost
CLI
Loop
FAQ
GitHub
Switch to dark mode
Switch to light mode
Install scribe
s
Single Go binary · runs on cron · self-hosted
Your knowledge base, written by your tools.
scribe reads your git history, your Claude Code & Codex sessions, and self-sent URLs, then writes the wiki for you — so the next agent session already knows what you decided and why. It’s memory your agents read before they act, not a second brain you maintain and never reopen: plain markdown in git, cross-project, cron-driven, and able to run 100% locally on Ollama for zero API spend.
View on GitHub
7,472 documents
maintainer's KB · zero typed by hand
$0 /sync
verified on 100% Ollama path
~70 s
weekly Dream cycle, gemma3:12b
~/projects/scriptorium — zsh
How it works
Three stages, one pipeline.
scribe mines four input streams, filters out the noise before any LLM touches it, then fans dense sources into entity-first wiki pages. Every step runs on cron; set it up once and forget it.
git
sessions
urls
drop files
BM25
triage
boilerplate · rejected · $0
absorb
pass 1 → 2
qmd index
agent reads
CLAUDE.md · AGENTS.md
1 Capture
Four streams, all on cron.
Git repos, Claude Code & Codex sessions, URLs you text yourself, and drop files from other projects. scribe auto-discovers every codebase you've ever opened in either CLI and keeps the manifest fresh.
Keyword-density scoring rejects boilerplate sessions before any LLM call, so cheap sessions cost nothing. Survivors go through a two-pass absorb: pass 1 grounds atomic facts, pass 2 fans dense sources into multiple entity-first wiki pages.
A typed graph of plain markdown.
Auto-generated wikilinks, backlinks JSON, and retrieval-context paragraphs spliced into every article so embeddings catch implicit entities. qmd reindexes for semantic search, reachable from any terminal, in any directory, or from inside Claude Code via MCP.
scribe isn't RAG, it isn't Obsidian, and it isn't another LLM-on-every-session burner. It sits between them: watches your work, writes the notes for you, and compounds knowledge across every project you touch.
scribe init writes a handshake block into both ~/.claude/CLAUDE.md and ~/.codex/AGENTS.md , so every session in every project queries your KB before recommending a library or proposing an architecture.
Hourly auto-commits. Every 2 hours: project extraction. 3×/day: session mining. Every 30 minutes: queued URLs. Every 4 hours: self-iMessaged links. Sundays at 02:00: the full Dream consolidation, with a lighter hot-domain pass daily in between.
One cross-project KB, not siloed per repo. Solve the oban idempotency bug in project A on Monday; the agent finds your fix on Friday when the same shape comes up in project B.
Per-project extraction, two-pass absorb, Dream cycle, assess, deep, session-mine, relations migrate: every LLM op runs end-to-end against a local Ollama server. One line in scribe.yaml flips the whole pipeline. Zero API spend.
File over app: the corpus has to outlive the pipeline. A git repo of plain markdown with YAML frontmatter. Push to your own GitHub, Gitea, or Forgejo; open in Obsidian, VS Code, vim, or mdbook. No SaaS, no vendor lock-in.
Articles connect via a closed 10-kind typed-edge schema: supersedes , contradicts , derived_from , specializes , extends , and five more. scribe relations migrate classifies existing related: links into it with an LLM.
wiki/decisions/ecto-multi.md
written by scribe · not typed by hand
---
type: decision
tags: [ecto, postgres, multi-tenancy]
supersedes: [[ecto-transaction]]
---
# Row-level tenancy over schema-per-tenant
Chose a tenant_id column with Postgres RLS over a schema
per tenant: migrations stay single-pass and the pool
doesn't fan out per tenant.
Why: schema-per-tenant broke oban job routing and made
advisory locks table-ambiguous under load.
## See also
- [[oban-idempotency]] · dedupe key includes tenant_id
- [[genserver-backpressure]] · per-tenant rate isolation
Not scribe running — scribe’s output. A page it wrote to git: YAML frontmatter, prose, and typed [[wikilinks]] . Those links are the edges the graph below draws.
supersedes
derived_from
contradicts
specializes
extends
derived_from
ecto-transaction
oban-idempotency
advisory-lock
token-bucket
genserver-backpressure
liveview-streams
ecto-multi
A slice of a real scribe KB: entities linked by typed edges like supersedes and contradicts, not folders or tags, so an agent can follow why a decision was made, not just keyword-match it.
4 input streams
45 subcommands
10 typed-edge kinds
3 inference modes
7 providers
two-pass absorb
0 vector DBs
0 required API keys
1 self-contained binary
The whole surface, in numbers — every figure is checkable against scribe --help and scribe.yaml , not marketing. Local mode needs no key at all; the default Anthropic path signs in through your claude -p CLI, not an API key.
Not RAG. Not Obsidian. Not another LLM-on-every-session burner.
scribe is a compiled knowledge base, not a vector database : it auto-writes a curated markdown wiki your agents query with BM25, so there's nothing to embed and nothing to host. The "second brain" debate is about notes you read. scribe isn't that. It's memory your agent reads before it decides: the reasons behind a choice, not summaries you'll never reopen. It sits between manual-notes tools (Obsidian, Notion) and unbounded LLM-on-every-query approaches (vanilla RAG, claude-memory-compiler): a curated wiki on top of raw sources kept verbatim, small enough for an agent to read whole, and cheap to query because most lookups are plain-text matches, not vector guesses.
↔ swipe the table to see every tool column
scribe is the only one of these that auto-writes a portable, git-versioned markdown wiki your agents read before they decide, with no vector database. Unlike AnythingLLM, scribe stores plain markdown in git and needs no vector database or running server. Snapshot 2026-06-08; tool capabilities change, so check each project's docs before deciding.
Snapshot 2026-06-08. Everything moves; check the source repos before deciding. Verdicts are pulled from public READMEs, issue trackers, and the scribe maintainer's tool evaluations.
scribe is the wiki the LLM writes for you, sitting on top of raw sources kept verbatim. RAG retrieves chunks; scribe gives you a curated, named-entity wiki you can also grep.
— project README
One KB. The whole team writes to it.
scribe is single-user by default and stays that way. But a small team on the same codebases can point every machine at one git-backed KB. The obvious fear with sharing agent sessions — a leaked secret, a private client repo, one teammate's config change reaching everyone — is exactly what the gates below are built to stop. Only knowledge you meant to keep crosses into the shared KB.
shared team KB
git · what teammates see
decisions
patterns
solutions
what tries to cross
the gate that stops it
secret-scan gate -> held -->
an AWS key in a transcript
secret-scan gate
held back — never committed
remote allowlist -> held -->
a teammate's client repo
remote allowlist
never ingested without an approve
trust layer -> held -->
config that repoints the LLM
trust layer
reverted to the last trusted snapshot
curated extract -> crosses -->
the reasoning behind a fix
curated extract
✓
promoted · provenance kept
Sensitive material stops at the commit gate. Only knowledge you deliberately keep crosses into the shared KB — and every gate here is a mechanism in the code, not a policy in a doc.
Shared config is untrusted by default
A trust layer pins the sensitive surface of a shared scribe.yaml : provider, model, ingest paths, and the secret scanner itself. A pushed change that would repoint inference to a new endpoint or drain a new directory into the KB reverts to the last trusted snapshot until a human approves it.
Secrets never reach shared history
scribe mines session transcripts, which routinely carry API keys and tokens. In team mode a deterministic secret scanner runs in the commit gate and holds flagged credentials back before anything lands in shared git history. No LLM, no network: a regex pass on the commit path.
Extraction is paid for once, not per laptop
Every scribe sync pulls, merges, and reindexes before it extracts, and a committed ledger keeps two machines from mining the same git revision twice. Your inference bill scales with commits, not with the number of laptops pointed at the KB.
A teammate's unrelated repo never leaks in
Discovered projects start pending. allowed_remotes and source filters gate discovery by git-remote identity, and scribe projects {list,approve,ignore,review} controls what enters the pipeline, so a side project or a client checkout never lands in the shared KB without an approve.
Curate privately, promote deliberately
scribe promote <article> --to team-kb copies a page from your personal KB into the shared one, provenance recorded. Derived and coordination files are refused as sources, so the team KB fills with what you meant to publish, not your working scratch.
One machine consolidates, no server
The weekly Dream consolidation rewrites, merges, and prunes the whole wiki, so exactly one machine should run it. A committed leader lease in the repo elects that machine: no etcd, no lock server, and two laptops never race to rewrite the same wiki at 02:00 Sunday.
Local, hosted, or Anthropic — and a bill you can read.
Run the whole pipeline on local Ollama for $0, point it at a hosted OpenAI-compatible API (Together, Groq, Fireworks, Hugging Face) when the laptop is running hot, or keep claude -p for the Anthropic path. The three coexist; pick the trade per op in scribe.yaml , and nothing reaches a paid provider without explicit config. scribe cost then reconciles every token across providers and KBs, to the cent against the provider's own dashboard.
~ $ scribe cost
scribe cost — last 7 days — 2 KBs (personal-kb, team-kb)
model calls ok cancl rate tmout wallclock in-tokens out-tokens usd
──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
sonnet 301 281 0 0 3 6h54m 250,468 789,357 $92.80
haiku 151 144 0 0 7 2h21m 1,251 551,761 $10.77
together/MiniMaxAI/MiniMax-M3 195 195 0 0 0 1h06m 1,261,197 146,728 $0.55
ollama/gemma3:12b 120 120 0 0 0 4h57m 1,037,866 138,258 —
ollama/qwen3:30b-a3b-instruct-2507-q4_K_M 1,643 1,627 0 0 0 23h39m 11,317,941 1,122,293 —
ollama/gemma3:4b 921 912 0 0 0 8h23m 2,355,608 672,157 —
──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
TOTAL 3,331 3,279 0 0 10 47h22m 16,224,331 3,420,554 $104.12
By provider:
provider calls in-tokens out-tokens usd
─────────────────────────────────────────────────
anthropic 452 251,719 1,341,118 $103.57
together 195 1,261,197 146,728 $0.55
ollama 2,684 14,711,415 1,932,708 —
By KB:
KB calls in-tokens out-tokens usd
──────────────────────────────────────────────────
team-kb 1,732 7,020,329 2,211,137 $102.96
personal-kb 1,599 9,204,002 1,209,417 $1.16
Coverage: 3175/3331 calls had token data; the other 156 add ~$0.62 estimated, not shown above.
usd = provider-billed spend incl. cache-write & cache-read; in is uncached input, so usd exceeds an in/out list estimate.
cancl = sibling-canceled (rate-limit cascade). tmout = ctx.DeadlineExceeded.
↔ swipe the table for the full per-model breakdown
Real output, names anonymized. Same machine, same week, three provider paths: Anthropic billed $103.57 , the hosted Together path $0.55 , the local Ollama models nothing. The paths handled different volumes (routing is per-op in scribe.yaml ), so read each as what it actually billed, not a like-for-like race. One API key usually bills every KB on a machine, so scribe cost aggregates them by default and reconciles to the dashboard.
On the local path there's no claude -p ca

[truncated]
