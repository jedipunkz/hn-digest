---
source: "https://github.com/dogtorjonah/context-warp-drive"
hn_url: "https://news.ycombinator.com/item?id=48722788"
title: "Show HN: Context Warp Drive – deterministic folding for LLM agents"
article_title: "GitHub - dogtorjonah/context-warp-drive · GitHub"
author: "VoxxoAI"
captured_at: "2026-06-29T18:38:39Z"
capture_tool: "hn-digest"
hn_id: 48722788
score: 2
comments: 0
posted_at: "2026-06-29T18:02:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Context Warp Drive – deterministic folding for LLM agents

- HN: [48722788](https://news.ycombinator.com/item?id=48722788)
- Source: [github.com](https://github.com/dogtorjonah/context-warp-drive)
- Score: 2
- Comments: 0
- Posted: 2026-06-29T18:02:38Z

## Translation

タイトル: HN を表示: コンテキスト ワープ ドライブ – LLM エージェントの決定論的折りたたみ
記事のタイトル: GitHub - Dogtorjonah/context-warp-drive · GitHub
説明: GitHub でアカウントを作成して、dogtorjonah/context-warp-drive の開発に貢献します。

記事本文:
GitHub - ドッグトルジョナ/context-warp-drive · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ドッグトルジョナ
/
コンテキストワープドライブ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ドッグトルジョナ/コンテキストワープドライブ
メインブランチタグ

s ファイルコードに移動 その他のアクションメニューを開く フォルダーとファイル
43 コミット 43 コミット .github/ workflows .github/ workflows docs docs サンプル サンプル スクリプト スクリプト src src test test .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md FILE-LIST.txt FILE-LIST.txt ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェントの記憶を要約するのはやめてください。すべての圧縮呼び出しでは、モデルのラウンドトリップが書き​​込まれ、プレフィックスが書き換えられるため、プロバイダー プロンプト キャッシュがコールドになり、エージェントが必要とする正確な識別子が静かに削除されます。代わりに決定的に折り畳んでください。
無限コンテキスト ワープ エンジン。 LLM 要約呼び出しやセッションを終了せずに、関数呼び出しエージェント セッションをコンテキスト ウィンドウで長時間維持し、プロバイダー プロンプト キャッシュをホットに保ち、エージェントが再度タッチした瞬間にページが折りたたまれたコンテンツを元に戻します。
決定論的。ゼロLLM。純粋な CPU、ゼロ I/O、同一入力に対するバイト同一出力。プロバイダーに依存しない: 人間のコンテンツ ブロック、OpenAI ツール呼び出し、および Gemini パーツ。
実稼働マルチエージェント システムから抽出され、すべてのモデルおよび長時間実行されるエージェント ワークロードにわたってコンテキストが継続的に折り畳まれます。
コア エンジンは、ローリング フォールド、リコール、フリーズ、統合にわたる 380 以上の決定論的テストに合格します。
以下のすべての数値は推定ではなく測定されたものです。クロード プロバイダーの使用状況台帳からの実稼働キャッシュ レート、クロードに対して再現可能なライブ (ANTHROPIC_API_KEY=… npx tsx Examples/benchmark-live.ts 、実際のモデル + 実際のサマライザー)、および正確な o200k_base BPE トークン カウントによるオフライン ( npx tsx Examples/benchmark.ts 、d)

決定論的、キーなし)。
出所に関する注記: このパブリック パッケージは運用環境から派生したものです。これは、プライベート マルチエージェント システム内でライブ実行されるエンジンのポータブル ディストリビューションであるため、一般的な WARP_* 環境名、パッケージに依存しないサンプル、生の履歴の回復の文言、およびツールに依存しない音声マイニングを意図的に使用します。バイト同一の不変式はこのパッケージにローカルなものであり、同一の入力が同一の折り畳まれたビューを生成しますが、プライベート統合レイヤーとのビットごとのパリティを主張するものではありません。
実稼働環境で測定 — 実際のクロード ワークロード、プロバイダー キャッシュ テレメトリ
重要な数値は、このエンジンが搭載する実稼働マルチエージェント システムからのものです。プロバイダー独自の使用状況台帳 (キャッシュ読み取りトークン ÷ 合計入力トークン) から測定された、フォールド/フリーズ エンジンを数百ターンにわたって継続的に実行する実際のクロード ワークロードです。
すべての入力トークンの約 90% は、これらの高ターン クロード ワークロード全体でキャッシュから提供されます。つまり、バイト同一の凍結フォールド プレフィックスが、ターンごとにその仕事を実行し、新規入力あたり $3.00/MTok ではなく、$0.30/MTok のキャッシュ読み取り (ソネット レート) で処理されます。再要約コンパクターはプレフィックスを書き換えるので、これを維持することはできません。切り詰めるとウィンドウがスライドして壊れます。これは、実際に測定された経済的議論の全体です。
あなた自身でそれを再現してください — 生きて、クロードと対戦してください
ANTHROPIC_API_KEY=sk-ant-... npx tsx Examples/benchmark-live.ts # デフォルトの claude-haiku-4-5
Real Claude は、Anthropic cache_control ブレークポイント、本物の Claude サマライザー (すべての識別子を保持するよう指示されています - 公正な戦い)、およびプロバイダー独自の cache_read_input_tokens / cache_creation_input_tokens を使用して毎ターン呼び出します。 16 ターンの短いデモでは、本番キャッシュ レートが控えめに示されています (キャッシュには 1024 トークン以上のトークン プレフィックスが必要であり、CWD の利点はセッションが長いとさらに大きくなります)。

t は実際のテレメトリのメカニズムを示しています。CWD はキャッシュから読み取り、切り捨てと要約によってプレフィックスを再構築します。
オフラインの決定論的デモ (API キーなし、実行ごとにバイトが同一)
npx tsx Examples/benchmark.ts — 16 ターンの停止デバッグ セッション、正確な o200k_base BPE トークン カウント (ポータブル プロキシ。Claude のトークナイザーは公開されていません)、claude-haiku-4-5 のリスト価格。これはCI煙テストです。キャッシュ列はターンオーバーターンのバイトプレフィックスプロキシであり、サマライザは透過的な決定論的な代役です（ヘッドカットオフを超えて埋められたIDをドロップします。これは、座標クローゼットが回避するために存在する障害モードです）。
CWD は最も安価で (Claude-haiku レートで要約に対して -71%、切り捨てに対して -62%)、追加のモデル呼び出しはゼロで、保持に関しては切り捨てに決定的に勝ります。 (適切なプロンプトを備えた実際のサマライザーは、より高いコストでの保持に匹敵します。CWD の永続的なエッジは、コスト + ゼロコール + 決定性 + ホット キャッシュです。) このエンジンはプロバイダーに依存しません。WARP_BENCH_MODEL (およびリストされていないモデルの場合は WARP_BENCH_PRICE_*) を設定して、OpenAI を含む任意のモデルに対してベンチマークを設定します。
エージェント セッションが長くなると、必ず同じ壁にぶつかります。つまり、コンテキスト ウィンドウがいっぱいになります。通常の答えはダメです:
切り詰めると履歴の途中が削除されます。エージェントは何をしていたか忘れてしまいます。
LLM 要約 (「圧縮」) はモデル呼び出しにコストがかかり、待ち時間が追加され、非決定的であり、プレフィックスを書き換えるたびにプロバイダー プロンプト キャッシュを破壊します。
Context Warp Drive はどちらも行いません。古いターンを決定論的に折りたたんでコンパクトな構造スケルトン (ツール呼び出しごとに 1 行 + 保持された推論) にし、顕著な正確な識別子 (UUID、SHA、パス、ポート) を予算に応じた座標クローゼットに保存し、折りたたまれたプレフィックスを凍結して、プロバイダー キャッシュがウォームである間はバイト同一で再利用され、ページが折りたたまれたコンテンツになります。

エージェントがパスに再タッチすると、自動的に戻ります。モデルコールはありません。切り捨てはありません。キャッシュはホットなままです。
まだnpmでは公開されていません。今すぐソースからインストールします。
git clone https://github.com/dogtorjonah/context-warp-drive.git
cd コンテキストワープドライブ
npm install # `prepare` を実行 -> dist/ を自動的に構築
# オプション — 参照 SQLite エピソード ストアのみ:
npm install better-sqlite3
コア (context-warp-drive/fold) にはランタイム依存関係がありません。 better-sqlite3 は、参照エピソード ストアでのみ必要なオプションのピアです。
dist/ は gitignore されるため、別のプロジェクトからパッケージを使用する前にビルドしてください。ローカル パッケージのインストールの場合:
npm run build # 明示的なフォールバック
npmパック
# 使用するプロジェクトから:
npm install /path/to/context-warp-drive/context-warp-drive- * .tgz
最初の npm 公開後のインストールは次のようになります。
npm install context-warp-drive
AIに配線してもらうと
ソース チェックアウトまたはローカル tarball から context-warp-drive を追加し、各モデル呼び出しの前に FoldSession.prepare() で関数呼び出しメッセージ履歴をラップします。生の履歴を個別に保存します。準備されたメッセージ ビューのみをプロバイダーに送信します。ログ記録には、cacheHot と stats を使用します。
次に、プロバイダー キャッシュ ノブを追加します。
コンテキスト ワープ ドライブは、プレフィックスのバイト同一性を維持します。プロバイダー SDK 呼び出しは引き続きプロバイダー固有のキャッシュ設定を所有します。
'context-warp-drive' から { FoldSession } をインポートします。
// 会話ごとに 1 つ。アクティブなウィンドウを超えて折りたたまれ、プロバイダーのキャッシュがホットに保たれます。
const session = new FoldSession() ;
// プロバイダー形式の完全な履歴 (Anthropic / OpenAI / Gemini メッセージ オブジェクト)。
const 履歴 = [
{ role : 'user' , content : 'src/parser.ts で失敗したテストを調査する' } ,
// ... 毎ターン成長します ...
] ;
// 毎ターン、モデルを呼び出す前に:
const { メッセージ 、cacheHot

, 統計 } = セッション。準備 (履歴、{
// オプションですが推奨: 実際のプロバイダー/リレー入力トークン テレメトリを渡します
// 前のターンから。デフォルトでは 240k で、FoldSession は新しいものを強制します。
// ホット再利用する代わりに、エポックを特大プロンプトに折りたたむ。
測定された入力トークン : 前の使用状況 ?。 input_tokens 、
} ) ;
// `messages` は送信する圧縮されたビューです。 `cacheHot` が true の場合、プレフィックスは次のようになります。
// 最後のターンとバイトが同一であるため、プロバイダー プロンプト キャッシュが再利用されます。
callYourModel (メッセージ) を待ちます。 // Anthropic / OpenAI / Gemini — メッセージの形状は変更されずに通過します
コンソール。 log ( `送信済み ${ メッセージ . 長さ } msgs ·cacheHot= ${cacheHot } · Saving= ${ stats . SavingPercent ?? 0 } %` ) ;
それが見出し全体です。継続的な常時リーン折り畳みの場合は、 ALWAYS_ON_FOLD_CONFIG を渡します。プロバイダーの実際のキャッシュ TTL と一致させるには、フリーズ: {enabled:true,ttlMs:3_600_000,maxTailChars:150_000} を設定します。測定されたトークンのプレッシャー ガードのデフォルトは DEFAULT_FOLD_PRESSURE_CEILING_TOKENS (240,000) です。無効にするには pressureCeiling: false を渡し、調整するには pressureCeiling: 120_000 を渡します。
完全なツール ループについては、examples/anthropic-loop.ts および example/openai-loop.ts を参照してください。
ハードエポック再生シードパリティ
FoldSession.prepare() には、Voxxo リレーで使用されるポータブル ハード エポック パスが含まれています。これは、プロバイダーに表示されるビューを 1 つの決定論的な復活シード メッセージに置き換え、トリガーとなるライブ ユーザー ターンを 1 回だけマージし、そのコンパクト シードを次の凍結プレフィックスとして再封印します。実際に測定された入力トークンが pressureCeiling に達すると、自動的に起動されます。ハーネスは同じパスを直接強制することもできます。
const 結果 = セッション 。準備 (履歴、{
ハードエポック : true 、
hardEpochSeed : renderMyHostRebirthPackage ( ) 、 // オプション;省略 = 生のトレースシード
測定された入力トークン : 前の使用状況

?。 input_tokens 、
} ) ;
Anthropic の場合、毎ターン、outcome.sealedBoundary をプロバイダー ヘルパーにフィードします。
import { prepareAnthropicCachedRequest } から 'context-warp-drive/providers/anthropic' ;
const キャッシュ = prepareAnthropicCachedRequest ( {
メッセージ: 結果。 AnthropicMessage [ ] としてのメッセージ、
sealedBoundary : 結果 。封印された境界 、
システム: SYSTEM_PROMPT 、
ツール : ツール 、
} ) ;
クライアントを待ちます。メッセージ。作成(
{モデル、最大トークン: 8192、...キャッシュされました。リクエスト } 、
キャッシュされた。 requestOptions 、
) ;
カスタム ハーネスのパリティ チェックリスト:
生の履歴は追加のみを保持し、完全な生のトレースを prepare() に渡します。
MeasurementInputTokens には測定されたプロバイダー トークン テレメトリを使用します。キャラクターからのプレッシャーを推定しないでください。
意図的な同じインスタンスの再誕生/リセットの場合は、hardEpoch: true に加えてレンダリングされたホスト シードを渡すか、パッケージに History から生のシードを計算させます。
タスク レール、ファイル クレーム、ワークスペースの状態、チャット、エピソード カードなどのホスト専用コンテキストを自分で永続化し、リレーのようなウェイク テキストが必要なときにそれらのセクションを renderRawRebirthSeed() に渡します。
クローン/モデル固有の ID デルタが安定したキャッシュされたプレフィックスに含まれないようにします。 Anthropic ヘルパーは、デフォルトで ## Your Identity の前にシステム プロンプトを分割します。より安価なクローンのファンアウト、PU

[切り捨てられた]

## Original Extract

Contribute to dogtorjonah/context-warp-drive development by creating an account on GitHub.

GitHub - dogtorjonah/context-warp-drive · GitHub
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
dogtorjonah
/
context-warp-drive
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
dogtorjonah/context-warp-drive
main Branches Tags Go to file Code Open more actions menu Folders and files
43 Commits 43 Commits .github/ workflows .github/ workflows docs docs examples examples scripts scripts src src test test .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md FILE-LIST.txt FILE-LIST.txt LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
Stop summarizing your agent's memory. Every compaction call burns a model round-trip, rewrites your prefix so the provider prompt cache goes cold, and quietly drops the exact identifiers your agent needs. Fold it deterministically instead.
The Infinite Context Warp Engine. Keep long function-calling agent sessions under the context window without LLM summarization calls and without ending the session — while keeping provider prompt caches hot — and page folded content back in the moment the agent touches it again.
Deterministic. Zero-LLM. Pure CPU, zero I/O, byte-identical output for identical inputs. Provider-agnostic: Anthropic content blocks, OpenAI tool_calls , and Gemini parts .
Extracted from a production multi-agent system, where it folds context continuously across every model and long-running agent workloads.
The core engine passes 380+ deterministic tests across rolling fold, recall, freeze, and integration.
Every number below is measured, not estimated — production cache rates from the Claude provider usage ledger, reproducible live against Claude ( ANTHROPIC_API_KEY=… npx tsx examples/benchmark-live.ts , real model + real summarizer) and offline with exact o200k_base BPE token counts ( npx tsx examples/benchmark.ts , deterministic, no key).
Provenance note: this public package is production-derived. It is the portable distribution of an engine that runs live inside a private multi-agent system, so it deliberately uses generic WARP_* environment names, package-neutral examples, raw-history recovery wording, and tool-agnostic voice mining. The byte-identical invariant is local to this package — identical inputs produce identical folded views — and is not a claim of bit-for-bit parity with any private integration layer.
Measured in production — real Claude workloads, provider cache telemetry
The numbers that matter are from the production multi-agent system this engine powers — real Claude workloads running the fold/freeze engine continuously across hundreds of turns , measured from the provider's own usage ledger (cache-read tokens ÷ total input tokens):
~90% of all input tokens are served from cache across these high-turn Claude workloads — that is the byte-identical frozen-fold prefix doing its job, turn after turn, at $0.30/MTok cache reads instead of $3.00/MTok fresh input (Sonnet rates). A re-summarizing compactor rewrites the prefix and can never sustain this; truncation slides the window and breaks it. This is the entire economic argument, measured live.
Reproduce it yourself — live, against Claude
ANTHROPIC_API_KEY=sk-ant-... npx tsx examples/benchmark-live.ts # default claude-haiku-4-5
Real Claude calls every turn with Anthropic cache_control breakpoints, a real Claude summarizer (told to preserve every identifier — a fair fight), and the provider's own cache_read_input_tokens / cache_creation_input_tokens . A short 16-turn demo understates the production cache rate (caching needs a ≥1024-token prefix and CWD's advantage compounds over long sessions) — but it shows the mechanism on real telemetry, with CWD reading from cache while truncation and summarization rebuild their prefix.
Offline deterministic demo (no API key, byte-identical every run)
npx tsx examples/benchmark.ts — a 16-turn outage-debugging session, exact o200k_base BPE token counts (a portable proxy; Claude's tokenizer isn't public), claude-haiku-4-5 list pricing. This is the CI smoke test; the cache column is a turn-over-turn byte-prefix proxy and the summarizer is a transparent deterministic stand-in (it drops ids buried past its head cutoff — the failure mode the Coordinate Closet exists to avoid).
CWD is cheapest ( −71% vs summarization, −62% vs truncation at Claude-haiku rates), makes zero extra model calls, and beats truncation decisively on retention. (A well-prompted real summarizer can match retention at higher cost — CWD's durable edge is cost + zero calls + determinism + a hot cache.) The engine is provider-agnostic: set WARP_BENCH_MODEL (and WARP_BENCH_PRICE_* for an unlisted model) to benchmark against any model, including OpenAI.
Every long agent session hits the same wall: the context window fills up. The usual answers are bad:
Truncation drops the middle of your history — the agent forgets what it was doing.
LLM summarization ("compaction") costs a model call, adds latency, is non-deterministic, and busts your provider prompt cache every time it rewrites the prefix.
Context Warp Drive does neither. It deterministically folds old turns into compact structural skeletons (one line per tool call + retained reasoning), conserves the salient exact identifiers (UUIDs, SHAs, paths, ports) in a budget-scored Coordinate Closet, freezes the folded prefix so it's reused byte-identical while the provider cache is warm, and pages folded content back in automatically when the agent re-touches a path. No model calls. No truncation. Cache stays hot.
Not published on npm yet. Install from source today:
git clone https://github.com/dogtorjonah/context-warp-drive.git
cd context-warp-drive
npm install # runs `prepare` -> builds dist/ automatically
# optional — only for the reference SQLite episode store:
npm install better-sqlite3
The core ( context-warp-drive/fold ) has zero runtime dependencies . better-sqlite3 is an optional peer needed only by the reference episodic store.
dist/ is gitignored, so build before consuming the package from another project. For a local package install:
npm run build # explicit fallback
npm pack
# from your consuming project:
npm install /path/to/context-warp-drive/context-warp-drive- * .tgz
After the first npm publish, installation becomes:
npm install context-warp-drive
If you ask an AI to wire it in
Add context-warp-drive from the source checkout or local tarball, then wrap our function-calling message history with FoldSession.prepare() before each model call. Preserve raw history separately; send only the prepared messages view to the provider. Use cacheHot and stats for logging.
Then add the provider cache knob:
Context Warp Drive keeps the prefix byte-identical. The provider SDK call still owns provider-specific cache settings.
import { FoldSession } from 'context-warp-drive' ;
// One per conversation. Folds past the active window + keeps the provider cache hot.
const session = new FoldSession ( ) ;
// Your full provider-shaped history (Anthropic / OpenAI / Gemini message objects).
const history = [
{ role : 'user' , content : 'Investigate the failing test in src/parser.ts' } ,
// ... grows every turn ...
] ;
// Every turn, before you call the model:
const { messages , cacheHot , stats } = session . prepare ( history , {
// Optional but recommended: pass real provider/relay input-token telemetry
// from the previous turn. At 240k by default, FoldSession forces a fresh
// fold epoch instead of hot-reusing into an oversized prompt.
measuredInputTokens : previousUsage ?. input_tokens ,
} ) ;
// `messages` is the compacted view to send. When `cacheHot` is true the prefix is
// byte-identical to last turn, so the provider prompt cache is reused.
await callYourModel ( messages ) ; // Anthropic / OpenAI / Gemini — the message shapes pass through unchanged
console . log ( `sent ${ messages . length } msgs · cacheHot= ${ cacheHot } · savings= ${ stats . savingsPercent ?? 0 } %` ) ;
That's the whole headline. For continuous always-lean folding, pass ALWAYS_ON_FOLD_CONFIG ; to match your provider's real cache TTL, set freeze: { enabled: true, ttlMs: 3_600_000, maxTailChars: 150_000 } . The measured-token pressure guard defaults to DEFAULT_FOLD_PRESSURE_CEILING_TOKENS (240,000); pass pressureCeiling: false to disable it or pressureCeiling: 120_000 to tune it.
See examples/anthropic-loop.ts and examples/openai-loop.ts for full tool loops.
Hard-epoch rebirth seed parity
FoldSession.prepare() includes the portable hard-epoch path used by the Voxxo relay: it replaces the provider-visible view with one deterministic rebirth seed message, merges the triggering live user turn exactly once, and reseals that compact seed as the next frozen prefix. It fires automatically when real measured input tokens reach pressureCeiling ; a harness can also force the same path directly:
const outcome = session . prepare ( history , {
hardEpoch : true ,
hardEpochSeed : renderMyHostRebirthPackage ( ) , // optional; omitted = raw trace seed
measuredInputTokens : previousUsage ?. input_tokens ,
} ) ;
For Anthropic, feed outcome.sealedBoundary to the provider helper every turn:
import { prepareAnthropicCachedRequest } from 'context-warp-drive/providers/anthropic' ;
const cached = prepareAnthropicCachedRequest ( {
messages : outcome . messages as AnthropicMessage [ ] ,
sealedBoundary : outcome . sealedBoundary ,
system : SYSTEM_PROMPT ,
tools : TOOLS ,
} ) ;
await client . messages . create (
{ model , max_tokens : 8192 , ... cached . request } ,
cached . requestOptions ,
) ;
Parity checklist for a custom harness:
Keep raw history append-only and pass the full raw trace to prepare() .
Use measured provider token telemetry for measuredInputTokens ; do not estimate pressure from characters.
For intentional same-instance rebirth/reset, pass hardEpoch: true plus your rendered host seed, or let the package compute the raw seed from history .
Persist host-only context such as task rails, file claims, workspace state, chat, and episode cards yourself, then pass those sections into renderRawRebirthSeed() when you need relay-like wake text.
Keep clone/model-specific identity deltas out of the stable cached prefix. The Anthropic helper splits the system prompt before ## Your Identity by default; for cheaper clone fanout, pu

[truncated]
