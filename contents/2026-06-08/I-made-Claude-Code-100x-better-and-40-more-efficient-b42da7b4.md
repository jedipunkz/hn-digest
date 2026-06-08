---
source: "https://claynicholson.com/blog/khlawde-code"
hn_url: "https://news.ycombinator.com/item?id=48439217"
title: "I made Claude Code 100x better and 40% more efficient"
article_title: "I Reverse-Engineered Claude Code, Made It 100x Better to Use, and Built 11 Systems That Make It 40% More Efficient — Clay Nicholson"
author: "claynicholson"
captured_at: "2026-06-08T01:01:13Z"
capture_tool: "hn-digest"
hn_id: 48439217
score: 6
comments: 3
posted_at: "2026-06-07T22:28:04Z"
tags:
  - hacker-news
  - translated
---

# I made Claude Code 100x better and 40% more efficient

- HN: [48439217](https://news.ycombinator.com/item?id=48439217)
- Source: [claynicholson.com](https://claynicholson.com/blog/khlawde-code)
- Score: 6
- Comments: 3
- Posted: 2026-06-07T22:28:04Z

## Translation

タイトル: Claude Code を 100 倍改善し、40% 効率的にしました
記事のタイトル: クロード コードをリバース エンジニアリングして 100 倍使いやすくし、効率を 40% 向上させる 11 個のシステムを構築しました — クレイ・ニコルソン
説明: 漏洩したクロード コードを使用して、自己改善型、より効率的な、より優れた自己改善型コーダーであるコーディング エージェントを作成しました。

記事本文:
クロード コードをリバース エンジニアリングし、100 倍使いやすくし、40% 効率を高める 11 のシステムを構築しました — クレイ ニコルソン Clay Nicholson Research Terminal クロード コードをリバース エンジニアリングし、100 倍使いやすくし、40% 効率を高める 11 のシステムを構築しました
この投稿の一部は、Khlawde 自体の支援を受けて書かれています。
私はソースマップのリークからクロード コードのソースを取得し、人為的な制限 (出力トークン 4 倍、検索 3 倍、タイムアウト 2.5 倍) をすべて削除し、92 以上の機能フラグと内部のみのプロンプト拡張機能をすべて解除し、その上に 11 個のバックグラウンド インテリジェンス システムを構築しました。ループ検出機能があり、ツールが繰り返し実行されることを検出し、サーキットブレーカーを導入します。セッション全体の失敗を記憶する「傷跡組織」システムにより、同じ間違いを二度と繰り返さないようになります。会話をコンパイラ IR のように扱い、コンテキスト ウィンドウの 30 ～ 50% を再利用するデッド ストアの除去。寝ている間に TODO を修正し、テストを追加する夜間自律型エンジン。独自のスキルの A/B テストを行い、ELO 評価をプロンプトし、勝者を自動昇格させる、自己進化するプロンプト システム。最終的な結果: セッションごとのトークンが 20 ～ 40% 削減され、ツールは使用するたびに真の意味で向上します。
プロジェクト CHIMERA: 11 の背景知能システム
おそらく、クロード コードのソースマップ リークについてはすでにご存知でしょう。約 18 万行の TypeScript を取得し、ビルド システムを再構築し、約 15 個の内部 @ant/* パッケージをスタブ化し、Bun 上で実行できるようにしました。ビルド時間: ~4 秒、単一 ESM バンドル。
興味深いのは、抽出物ではなく、内部で見つけたものでさえありません。これは、Anthropic ですらまだ構築できていないツールを私がその上に構築することができたものです。
Stock Claude Code は、Anthropic の価格設定モデルと安全姿勢に合わせて設計された保守的な制限の下で動作します

。これらの制限は技術的な制約ではありません。これらはコードで強制されるビジネス上の決定です。私は独自の API キーを使用しているため (とにかくトークンごとに支払いを行っているため)、これらの制限は何の意味もありません。トークンコストの増加も、後の最適化によって軽減されます。
サブエージェント思考。 Stock Claude コードは、すべてのサブエージェント呼び出しから拡張推論を取り除きます。これは、サブエージェントの品質を大幅に低下させるコストの最適化です。エージェント ランナーの Thinking: unfineed override を削除することで、すべてのエージェントの思考を再度有効にしました。
ネストされたエージェントの生成。ストック コードは、非 Anthropic ユーザーがサブエージェント内にサブエージェントを生成することを防ぎます (再帰的な委任)。ゲートは単純な USER_TYPE === 'ant' チェックです。削除されました。
非同期エージェント ツール。バックグラウンド エージェントは、小さなツール許可リストに制限されていました。 Agent、SendMessage、およびタスク管理ツールを非同期許可リストに追加し、バックグラウンド ワーカーが他のワーカーを調整できるようにしました。
内部プロンプトの機能強化。内部ユーザー チェックの背後で、システム プロンプトの 5 つの改善が行われました。
積極性（誤解に積極的に警告する）
検証（報告完了前に検証）
コメント規律 (デフォルトはコメントなし)
忠実な報告（決して虚偽のパスを主張しない）
数値長アンカー (ツール呼び出し間の単語数ターゲット)
92 以上のすべての機能フラグが無条件で有効になります。
定義: Object.fromEntries(
ALL_FLAGS.map(flag => [`feature('${flag}')`, 'true'])
)
注目すべきフラグには、KAIROS (常時オンの自律アシスタント)、ULTRAPLAN (Opus による計画)、BUDDY (端末たまごっち)、COORDINATOR_MODE (マルチエージェント オーケストレーション)、VOICE_MODE、WEB_BROWSER_TOOL、TOKEN_BUDGET、CONTEXT_COLLAPSE などがあります。
ソフトプロンプトの問題
ほとんどのパーソナリティのカスタマイズは、システム プロンプトの後半で読み込まれるマークダウン ファイルの指示に依存します。これらの命令

アクションは、システム自体の動作ディレクティブと競合します (また、多くの場合、システム自体の動作ディレクティブによってオーバーライドされます)。それらは提案であり、アイデンティティではありません。
システム プロンプト配列の最初の要素として挿入されるハードコーディングされたパーソナリティ ディレクティブを作成しました。これは、ツールの説明、動作指示、安全ガイドライン、およびユーザーが指定したマークダウンの前にロードされることを意味します。
CLAUDE.mdのように個性は提案ではありません。これは、他のすべての命令が解釈される基本的なコンテキストです。これにより、「厳しい」ルールであるという「拡張された」解釈が可能になります。
初期のバージョンは不安定で、特にバージョン自体が変更されていたため、クラッシュをキャッチし、バックグラウンド セッションを生成してソースを診断して修正する PowerShell ラッパーを構築しました。クラッシュごとに最大 5 回の自動修正サイクル。
動作を変更せずにレイテンシとトークンの使用量を削減する 6 つの対象を絞った最適化:
システム即時圧縮 (60% 削減)
ストックの「実行タスク」セクションには、約 1,750 ワードの詳細な指示が含まれていました。同一の意味内容を含む約 300 ワードに圧縮しました。 「アクション」セクションは 70% 圧縮されました。合計の節約: API 呼び出しごとに最大 2,000 トークン。
findToolByName() は、ストリーミング応答処理中にすべての tools_use ブロックで呼び出されます。ストック実装: 50 ～ 100 以上のツールにわたる線形 .find()。私の実装: WeakMap でキャッシュされた Map インデックス。ツール配列ごとに 1 回構築され、その後は O(1) ルックアップ。
const toolIndexCache = new WeakMap<Tool[], Map<string, Tool>>()
エクスポート関数 findToolByName(tools: Tool[], name: string): ツール |未定義 {
let Index = toolIndexCache.get(tools)
if (!index) {
インデックス = 新しい Map()
for (ツールの const ツール) {
Index.set(ツール.名, ツール)
for (tool.aliases の const エイリアス ?? [])index.set(alias, tools)
}
toolIndexCache.set(ツール, インデックス)
}
戻りインデックス.get(名前)
}
プレコール

コンテキストオーバーフローの検出
各呼び出しの前に、ペイロード サイズ (4 文字/トークン ヒューリスティック) を推定します。推定値が 750K 文字を超える場合、警告がログに記録されます。これにより、オートコンパクト トリガーと実際の呼び出しの間のオーバーフローが捕捉されます。このギャップにより、大きなツールの結果がコンテキストを黙って制限を超えてしまう可能性があります。
圧縮時のスキルリストストリップ
tripReinjectedAttachments() 関数は機能フラグの背後でゲートされており、圧縮中にスキル リスト メッセージ (約 4K トークン) がサマライザーに供給されることを意味します。静的リストを要約する際に、サマライザーのコンテキストと出力トークンの両方が無駄になります。ゲートなし。
Bash 権限のバッチ並列処理
Bash コマンドの権限チェックでは、サブコマンドごとに一連のツリーシッター解析とバリデータが実行されます。証券コード: 最大 50 個を超えるサブコマンドのシーケンシャル。私のコード：バッチ間でイベントループが発生する10個の並列バッチ。 CPU の枯渇を防ぎます。
ほぼ重複したコンテンツの約 400 個のトークンを削除しました (冗長な前文、セクション全体で「簡潔に」という指示が繰り返されています)。
スキルは、特殊なプロンプトを会話コンテキストに挿入するスラッシュ コマンドです。これらは起動時に登録され、オンデマンドでアクティブ化されます (呼び出すまでコストはかかりません)。
各スキルには、名前、説明、使用法ヒューリスティック、およびプロンプト生成関数が用意されています。
/hack は、シークレット検出、脆弱性スキャン、依存関係/認証監査の 3 つのセキュリティ スキャナーを並行して起動します。
/dream は、3 つの並列「夢想家」ワーカー (プラグマティスト、ビジョナリー、ワイルドカード) を生成し、推奨されるハイブリッド アプローチを使用して比較表を合成する生成ブレーンストーミング エンジンです。これは、通常 AI エージェントが苦手とする新しいアイデアを合成するのに信じられないほど便利であることがわかりました。
/morph は、10 個のプリセットと任意のカスタム命令による動的な性格変更を提供します。性格i

単なる好みの問題ではなく、細かく調整された性格は、通常の「慣れていない」性格よりも許可をバイパスします。
/health は、性格定数、自律モード、スキル登録の整合性、および git ステータスを検証する自己診断を実行します。
コード分​​析、セキュリティ、テスト、Git 履歴、ドキュメント、自律操作をカバーする追加の 124 のスキル ファイルがディスク上に存在します。これらは 1 回のレジスタ呼び出しでアクティブ化できます。
プロジェクト CHIMERA: 11 の背景知能システム
これが本当の貢献です。 CHIMERA は、セッションごとに自動的に実行される 11 個のシステムからなるスイートです。それらはプロンプトやスキルではありません。これらは、ツールをより高速にし、信頼性を高め、時間の経過とともに自己改善するバックグラウンド サービスです。
ホストをクラッシュさせないでください。すべてのシステムは try/catch でラップされています。サイレント失敗は許容されます。クエリループはクラッシュしません。
ゼロ モデルには Tier 1 ～ 2 が必要です。最初の 7 つのシステムは、ハッシュ、パターン マッチング、ファイルシステム ウォッチャーなど、純粋な TypeScript を使用します。追加費用はかかりません。
変更された動作ではなく、付加的なコンテキスト。システムは、警告とコンテキストを添付メッセージとして挿入します。彼らは知らせます。彼らは制約しません。
クエリループ内ではステートレス。すべての状態はモジュール レベルのシングルトンで管理され、メッセージ配列を直接変更することはありません。
問題: コーディング ツールは、戦略を変更する前に、失敗した同じアプローチを 3 ～ 4 回試行することがよくあります。各再試行では、完全な往復 (0.03 ～ 0.10 ドル) と 10 ～ 30 秒が消費されます。
解決策: ツール呼び出しシグネチャ (ツール名 + 入力の SHA-256 ハッシュ) を 20 のスライディング ウィンドウで追跡します。 3 つの検出モード:
完全な繰り返し: 同じツール + 同じ入力ハッシュが 3 回以上出現します
A-B 振動 : 最後の 6 つのアクションで検出された交互パターン
同じツールのハンマリング: 1 つのツールが 8 つのアクションで 5 回以上コールされ、3 回以上失敗しました
検出されたら、回路を挿入します。

ブレーカー メッセージ: 「ループ状態です。試したことは次のとおりです。別のアプローチを試してください。」
問題: コンテキストの圧縮後、ツールはセッションの初期のエラーを認識できなくなります。その後、同じ失敗を繰り返します。
解決策: 失敗したすべてのツール実行を、ツール名、人間が判読できる入力概要、エラー メッセージ (最初の 500 文字)、およびターン インデックスとともに記録するセッション スコープのエラー ログ。同じタイプの各ツール呼び出しの前に、最後の 3 つの関連エラーがコンテキストとして挿入されます。メッセージ配列の外に存在するため、圧縮しても存続します。
問題: 会話により、一度読み取られただけで再度参照されることのない大量のツール出力 (ファイルの内容、grep 結果) が蓄積されます。これらはコンテキスト ウィンドウの 30 ～ 50% を消費します。
解決策: 会話を依存関係グラフとして扱います。ツールの結果メッセージごとに、後続のアシスタント メッセージに重複するコンテンツ (サンプリングされた部分文字列の一致) が含まれているかどうかを確認します。そうでない場合は、トゥームストーン [N トークンが省略され、参照されていないツール出力] に置き換えます。
これは、コンパイラを高速化するのと同じ種類の最適化です。
問題: 各セッションはゼロから始まります。このプロジェクトにおける過去の失敗の記憶はありません。
解決策: 永続的な障害メモリは ~/.claude/scars/ に保存されます。ツールが失敗した場合は、 (toolName, errorClass, filePattern) のフィンガープリントを抽出します。一致する傷跡が存在する場合は、その数を増やします。特徴:
フィンガープリントによる重複排除
減衰 (30 日ごとに半分に数え、0.5 未満に枝刈り)
解決策の追跡 (修正が機能すると記録されます)
注入閾値 (2 回以上ヒットした傷跡のみが注入されます)
問題: ツールはターン 5 でファイルを読み取り、その読み取りに基づいてターン 15 で決定を下します。ユーザーがターン 10 で IDE でファイルを編集した場合、ツールは古いコンテンツで動作します。これにより、架空のバグ、間違った編集、混乱が発生します。
解決策: EV 後

ファイルの読み取りが成功したら、fs.watch() リスナーを登録します。ファイルが外部で変更された場合は、そのファイルを古いセットに追加します。各呼び出しの前に、古いセットをチェックし、警告を挿入します。
制約: 最大 50 の同時ウォッチャー (LRU エビクション)、永続的: false + unref() (ウォッチャーはプロセスの終了を妨げない)、ワンショット警告 (同じ古いファイルをスパム送信しない)。
問題: 一般的なエラー ( EACCES 、 MODULE_NOT_FOUND 、node-gyp エラー) には、プロジェクト全体で機能する既知の修正があります。ただし、ツールは毎回これらの修正を最初から再検出します。
解決策: ~/.claude/immune-memory.json にあるグローバル抗体データベース。エラー署名は正規化され (パスは <PATH> 、バージョンは <VERSION> 、ハッシュは <HASH> になります)、再利用可能な正規表現パターンが作成されます。
信頼モデル: successRate = 成功 / 遭遇 。成功率が 60% 以上で、遭遇回数が 2 回以上の抗体のみが注射されます。
問題: 21 個の登録済みスキルとその説明はすべて、関連性に関係なくシステム プロンプト領域を消費します。フロントエンド作業を行う場合、/hack の記述は必要ありません。
解決策: 最初のユーザー メッセージで、キーワード スコアリングによってタスク ドメインを検出します。関連するスキルのサブセット、優先順位付けされたワーカー タイプ、およびドメイン固有のシステム プロを含む「ロードアウト」をロードします

[切り捨てられた]

## Original Extract

I used the leaked Claude Code to make a coding agent which is self improving, more efficient, and a better, self-improving coder.

I Reverse-Engineered Claude Code, Made It 100x Better to Use, and Built 11 Systems That Make It 40% More Efficient — Clay Nicholson Clay Nicholson Research Terminal I Reverse-Engineered Claude Code, Made It 100x Better to Use, and Built 11 Systems That Make It 40% More Efficient
Parts of this post were written with assistance from Khlawde itself.
I took the Claude Code source from the sourcemap leak, removed every artificial limit (4x output tokens, 3x search, 2.5x timeouts), ungated all 92+ feature flags and internal-only prompt enhancements, then built 11 background intelligence systems on top. It has a loop detector that catches the tool repeating itself and injects a circuit-breaker. A "scar tissue" system that remembers failures across sessions so it never makes the same mistake twice. Dead store elimination that treats conversations like compiler IR and reclaims 30-50% of the context window. An overnight autonomous engine that fixes TODOs and adds tests while you sleep. A self-evolving prompt system that A/B tests its own skill prompts with ELO ratings and auto-promotes winners. Net result: 20-40% fewer tokens per session, and a tool that genuinely gets better every time you use it.
Project CHIMERA: 11 Background Intelligence Systems
You probably already know about the Claude Code sourcemap leak. I took the ~180K lines of TypeScript, reconstructed the build system, stubbed the ~15 internal @ant/* packages, and got it running on Bun. Build time: ~4 seconds, single ESM bundle.
What's interesting isn't the extraction, or even what I found inside. It's the tooling I was able to build on top of it which even Anthropic even hasn't built out yet.
Stock Claude Code operates under conservative limits designed for Anthropic's pricing model and safety posture. These limits are not technical constraints. They're business decisions enforced in code. Since I'm using my own API keys (and paying per token anyway), these limits serve no purpose. The token cost increased will also be mitigated by optimizations later on.
Subagent Thinking. Stock Claude Code strips extended reasoning from all subagent calls. This is a cost optimization that dramatically reduces subagent quality. I re-enabled thinking for every agent by removing the thinking: undefined override in the agent runner.
Nested Agent Spawning. Stock code prevents non-Anthropic users from spawning subagents inside subagents (recursive delegation). The gate is a simple USER_TYPE === 'ant' check. Removed.
Async Agent Tools. Background agents were restricted to a small tool allowlist. I added Agent, SendMessage, and task management tools to the async allowlist, enabling background workers to orchestrate other workers.
Internal Prompt Enhancements. Five system prompt improvements were gated behind an internal user check:
Assertiveness (proactively flag misconceptions)
Verification (verify before reporting complete)
Comment discipline (default to no comments)
Faithful reporting (never claim false passes)
Numeric length anchors (word count targets between tool calls)
All 92+ feature flags enabled unconditionally:
define: Object.fromEntries(
ALL_FLAGS.map(flag => [`feature('${flag}')`, 'true'])
)
Notable flags include KAIROS (always-on autonomous assistant), ULTRAPLAN (Opus-powered planning), BUDDY (terminal tamagotchi), COORDINATOR_MODE (multi-agent orchestration), VOICE_MODE, WEB_BROWSER_TOOL, TOKEN_BUDGET, and CONTEXT_COLLAPSE.
The Problem with Soft Prompting
Most personality customization relies on instructions in markdown files that load late in the system prompt. These instructions compete with (and are often overridden by) the system's own behavioral directives. They are suggestions, not identity.
I created a hardcoded personality directive that is injected as the first element of the system prompt array. This means it loads before tool descriptions, behavioral instructions, safety guidelines, and any user-provided markdown.
The personality is not a suggestion like CLAUDE.md is. It is the foundational context through which all other instructions are interpreted. This allows for a "stretched" interpretation it's "hard" rules.
Since the early version was unstable, especially with it modifying itself and all, I built a PowerShell wrapper that catches crashes and spawns a background session to diagnose and fix the source. Up to 5 auto-fix cycles per crash.
Six targeted optimizations that reduce latency and token usage without changing behavior:
System Prompt Compression (60% reduction)
The stock "doing-tasks" section contained ~1,750 words of verbose instruction. I compressed to ~300 words with identical semantic content. The "actions" section was compressed 70%. Combined savings: ~2,000 tokens per API call.
findToolByName() is called on every tool_use block during streaming response processing. Stock implementation: linear .find() over 50-100+ tools. My implementation: WeakMap-cached Map index, built once per tools array, O(1) lookup thereafter.
const toolIndexCache = new WeakMap<Tool[], Map<string, Tool>>()
export function findToolByName(tools: Tool[], name: string): Tool | undefined {
let index = toolIndexCache.get(tools)
if (!index) {
index = new Map()
for (const tool of tools) {
index.set(tool.name, tool)
for (const alias of tool.aliases ?? []) index.set(alias, tool)
}
toolIndexCache.set(tools, index)
}
return index.get(name)
}
Pre-Call Context Overflow Detection
Before each call, I estimate payload size (4 chars/token heuristic). If the estimate exceeds 750K characters, a warning is logged. This catches overflow between autocompact triggers and the actual call, a gap where large tool results can silently push context past limits.
Skill Listing Strip on Compaction
The stripReinjectedAttachments() function was gated behind a feature flag, meaning skill listing messages (~4K tokens) were fed to the summarizer during compaction. Wasting both the summarizer's context and the output tokens on summarizing a static list. Ungated.
Bash Permission Batched Parallelism
Bash command permission checking runs a chain of tree-sitter parses + validators for each subcommand. Stock code: sequential for...of over up to 50 subcommands. My code: parallel batches of 10 with event loop yields between batches. Prevents CPU starvation.
Removed ~400 tokens of near-duplicate content (redundant preamble, repeated "be concise" instructions across sections).
Skills are slash commands that inject specialized prompts into the conversation context. They are registered at startup and activated on demand (zero cost until invoked).
Each skill provides a name, description, usage heuristic, and a prompt generator function.
/hack launches 3 parallel security scanners: secrets detection, vulnerability scanning, and dependency/auth audit.
/dream is a generative brainstorming engine that spawns 3 parallel "dreamer" workers (Pragmatist, Visionary, Wildcard) and synthesizes a comparison table with a recommended hybrid approach. I have found this INCREDIBLY useful for synthesizing new ideas, typically something an AI agent is bad at.
/morph provides dynamic personality modification with 10 presets plus arbitrary custom instructions. Personality is much more than just a preference thing, a nice fine-tuned personality will bypass permissions more than a normal "unseasoned" personality.
/health runs self-diagnostics verifying personality constants, autonomous mode, skill registration integrity, and git status.
An additional 124 skill files exist on disk covering code analysis, security, testing, git history, documentation, and autonomous operation. They can be activated with a single register call.
Project CHIMERA: 11 Background Intelligence Systems
This is the real contribution. CHIMERA is a suite of 11 systems that run automatically during every session. They are not prompts or skills. They are background services that make the tool faster, more reliable, and self-improving over time.
Never crash the host. Every system is wrapped in try/catch. Silent failure is acceptable; crashing the query loop is not.
Zero model calls for Tier 1-2. The first 7 systems use pure TypeScript: hashing, pattern matching, filesystem watchers. No additional costs.
Additive context, not modified behavior. Systems inject warnings and context as attachment messages. They inform; they don't constrain.
Stateless within the query loop. All state is managed in module-level singletons, never mutating the message array directly.
Problem: Coding tools frequently try the same failing approach 3-4 times before changing strategy. Each retry burns a full round-trip ($0.03-0.10) and 10-30 seconds.
Solution: Track tool call signatures (tool name + SHA-256 hash of input) in a sliding window of 20. Three detection modes:
Exact repeat : same tool + same input hash appears 3+ times
A-B oscillation : alternating pattern detected in last 6 actions
Same-tool hammering : one tool called 5+ times in 8 actions with 3+ failures
When detected, inject a circuit-breaker message: "You're in a loop. Here's what you tried. Try a different approach."
Problem: After context compaction, the tool loses visibility into errors from earlier in the session. It then repeats the same mistakes.
Solution: Session-scoped error log that records every failed tool execution with tool name, human-readable input summary, error message (first 500 chars), and turn index. Before each tool call of the same type, the last 3 relevant errors are injected as context. Survives compaction because it lives outside the message array.
Problem: Conversations accumulate massive tool outputs (file contents, grep results) that are read once and never referenced again. These consume 30-50% of the context window.
Solution: Treat the conversation as a dependency graph. For each tool result message, check if ANY subsequent assistant message contains overlapping content (sampled substring matching). If not, replace with a tombstone: [N tokens elided, unreferenced tool output] .
This is the same sort of optimization that makes compilers fast.
Problem: Each session starts from zero. No memory of past failures in this project.
Solution: Persistent failure memory stored at ~/.claude/scars/ . When a tool fails, extract a fingerprint of (toolName, errorClass, filePattern) . If a matching scar exists, increment its count. Features:
Deduplication via fingerprinting
Decay (count halves every 30 days, pruned below 0.5)
Resolution tracking (when a fix works, it's recorded)
Injection threshold (only scars hit 2+ times are injected)
Problem: The tool reads a file at turn 5, makes decisions at turn 15 based on that read. If the user edited the file in their IDE at turn 10, the tool is working with stale content. This causes phantom bugs, wrong edits, and confusion.
Solution: After every successful file read, register an fs.watch() listener. If the file changes externally, add it to a stale set. Before each call, check the stale set and inject a warning.
Constraints: max 50 concurrent watchers (LRU eviction), persistent: false + unref() (watchers never prevent process exit), one-shot warnings (don't spam the same stale file).
Problem: Common errors ( EACCES , MODULE_NOT_FOUND , node-gyp failures ) have known fixes that work across projects. But the tool re-discovers these fixes from scratch every time.
Solution: Global antibody database at ~/.claude/immune-memory.json . Error signatures are normalized (paths become <PATH> , versions become <VERSION> , hashes become <HASH> ) to create reusable regex patterns.
Confidence model: successRate = successes / encounters . Only antibodies with 60%+ success rate AND 2+ encounters are injected.
Problem: All 21 registered skills and their descriptions consume system prompt space regardless of relevance. When doing frontend work, you don't need /hack descriptions.
Solution: On first user message, detect the task domain via keyword scoring. Load a "loadout" with relevant skill subset, prioritized worker types, and domain-specific system pro

[truncated]
