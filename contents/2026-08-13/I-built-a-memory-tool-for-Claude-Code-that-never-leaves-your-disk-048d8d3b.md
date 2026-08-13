---
source: "https://github.com/georgedagher/deposition"
hn_url: "https://news.ycombinator.com/item?id=49288724"
title: "I built a memory tool for Claude Code that never leaves your disk"
article_title: "GitHub - georgedagher/deposition: Local-only decision log for AI agent sessions -- zero API calls, deterministic decision extraction. · GitHub"
author: "georgedagher"
captured_at: "2026-08-13T17:51:28Z"
capture_tool: "hn-digest"
hn_id: 49288724
score: 2
comments: 0
posted_at: "2026-08-13T16:53:27Z"
tags:
  - hacker-news
  - translated
---

# I built a memory tool for Claude Code that never leaves your disk

- HN: [49288724](https://news.ycombinator.com/item?id=49288724)
- Source: [github.com](https://github.com/georgedagher/deposition)
- Score: 2
- Comments: 0
- Posted: 2026-08-13T16:53:27Z

## Translation

タイトル: ディスクから離れることのないクロード コード用の記憶ツールを構築しました
記事のタイトル: GitHub - georgedagher/deposition: AI エージェント セッションのローカルのみの意思決定ログ -- API 呼び出しゼロ、決定論的な意思決定抽出。 · GitHub
説明: AI エージェント セッションのローカルのみの意思決定ログ -- API 呼び出しはなく、決定論的な意思決定抽出。 - ジョージダガー/証言録取

記事本文:
GitHub - georgedagher/deposition: AI エージェント セッションのローカルのみの意思決定ログ -- API 呼び出しなし、決定論的な意思決定抽出。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ジョージダーガー
/
堆積
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows bin bin 例 例 フック フック スキル/ デポジット-リコール スキル

ls/deposition-recall .gitignore .gitignore ライセンス ライセンス README.md README.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントのローカルのみの意思決定ログ。それは覚えていない
あなたが言ったことすべて - あなたが決めたことを覚えていて、それを返します
次回のセッションは無料で、マシンからは何も残されません。
AI コーディング エージェントを毎日使用している場合は、次のいずれかに該当していることになります。
同じ訴訟を再度訴訟します。あなたは 20 分をかけて説得しました
Mongo 上で Postgres を使用するエージェント。新しいセッション、新しいコンテキスト ウィンドウ - それ
まるで会話がなかったかのように、モンゴは再び処刑されました。あなたは説明します
もう一度自分自身を。もしかしたら3分の1かもしれない。
「それはもう言いましたね。」 3セッション前にあなたは「私たちはそうではない」と言った
ライブラリ X を使用しているため、メンテナンスされていません。」コンテキストが圧縮され、セッションが
終わったとしても、それは消えてしまい、ライブの提案として再び表示されます。
一般的な記憶ツールは音が大きすぎて読めません。メモリプラグインを試しました
それは「すべてを覚えている」ということです。現在、すべてのセッションは次の壁で始まります。
リコールされたツール呼び出し、ファイル読み取り、およびチャタリング。あなたはそれを読むのをやめます。の
記憶は存在する。それは役に立ちません。
あなたを含め、誰もその理由を覚えていません。チームメイト（または未来のあなた）が尋ねます
「なぜ Y ではなく X を選んだのですか?」正直な答えは 20 分です
チャット履歴や PR の説明を通じて、決定したことについて
数週間前の一文で。
この記録をどこにも送信することはできません。会話には、
クライアントコード、内部名、または要約したくないものだけ
メモリを取り戻すために、他の人の LLM API を使用することもできます。ほとんどの記憶ツール
まさにその呼び出しが必要です。
証言録取書は、ほぼすべての価値が存在するという観察に基づいて構築されています。
「セッションを思い出す」というのは、実際にはほんの一握りのことを思い出しているだけです。
その中で行われる決定 — ツール呼び出しやバックアンドフォートではありません

ああ、そうではありません
行き止まり。これらを理解すれば、残りをスキップすれば、次の時点で 90% のメリットが得られます。
騒音もコストもほぼ0%です。
$ deposition.sh クエリ "どのデータベースを選択しましたか" --決定
=== 思い出してください: 「どのデータベースを選択したのか」 ===
[1] 2026-08-14 · セッションデモセッション/セッション · sim=0.55 · テーマ: -
決定: 一貫性のニーズを比較検討した結果、次の方法を採用することにしました。
強力なリレーショナル保証が必要なため、新しいサービスには Postgres を使用します。
ソース: デモセッション/session.jsonl
これは、examples/ のサンプル トランスクリプトに対する実際の実行です。試してみてください。
自分で行う場合は、約 1 分かかります (以下のクイックスタートを参照)。 API キーはありません、いいえ
アカウント、ネットワーク通話なし。
ほとんどのエージェントメモリ ツールは、LLM を呼び出してすべての観察を要約します。
結果をベクトル データベースに保存します (多くの場合、オプションのクラウド同期を使用します)。
これは強力ですが、書き込みごとにコストがかかり、表面化する傾向があります
実際には返す必要のなかった多くのコンテキストが返されます。
証言録取は逆の賭けをする: 要約電話をスキップし、試行もスキップする
すべてを覚えて、次のような文章をキャッチするだけです。
彼らが通り過ぎるときの決定。
蒸着と代替品の比較
堆積
クロード・メム
メム0 / ゼップ / レッタ
それを実行するための API 呼び出し
なし
はい (LLM 要約)
あり（抽出・グラフ化）
保存するもの
決定のみ
すべてが圧縮された状態
すべて（事実/グラフ）
セットアップ
1 つの環境変数
プラグインのインストール + ワーカー
APIキー+サービス
可動部品
2 つのスクリプト
フック + ワーカー + DB + UI
ホスト型サービスまたはセルフホスト型サービス
に良い
「物事を決め直すのはやめよう」
豊富なセッション継続性
長時間実行エージェントのメモリ
ランキングではなく、さまざまな仕事です。高度な自己編集機能を備えたエージェントが必要な場合
長期記憶、claude-mem/Letta/Zep/Mem0 の方が能力が高く、より多く
成熟したツールなので、それを使用する必要があります。デポジションはより狭い範囲のジョブに使用されます。
決定を再訴訟することは決してなく、費用はかかりません。

あなたのデータをゼロにする
機械。
クイック スタート: クロード コード プラグインとして (推奨)
/plugin Marketplace add georgedagher/deposition
/プラグインインストールのデポジット
それだけです -- 設定する環境変数はありません (デフォルトは ~/.claude/projects です、Claude)
コード自体のトランスクリプト ディレクトリ)、cron ジョブはありません。停止フックは新しいインデックスを作成します
各ターンの後に決定を下し、バンドルされたスキルはクロードにチェックするよう教えます。
何かが議論されなかったと仮定したり、再提案したりする前の証言録取
あなたがすでに行った選択。
検証ステータス: クロード プラグインは両方のプラグインの検証パスを確認しました
マーケットプレイス マニフェスト、および実際のクロード プラグイン マーケットプレイスを追加 +
クロードプラグインのインストール deposition@deposition は、次のようにして正常に完了します。
プラグインが有効になっていると表示 -- マニフェスト、マーケットプレイスのリスト、および
フックの登録はすべて構造的に正しいです (これは実際のエラーをキャッチして修正しました)
バグ: 以前のバージョンではフック ファイルが二重宣言されており、ロードに失敗していました)。
基礎となる取り込み/クエリ コマンドはエンドツーエンドでテストされます (以下を参照)。
実際にライブ中にストップフックの発射とスキル発動を見てみた
この環境では会話がまだ検証されていません (認証されていません)
これが構築された場所では対話型セッションが利用可能) -- 問題が発生した場合
具体的には、1 つ開いてください。
git clone https://github.com/georgedagher/deposition.git && cd デポジット
./install.sh # 必要に応じて UV をインストールします
import DEPOSITION_TRANSCRIPT_ROOTS= " $HOME /.claude/projects " # エージェントの .jsonl トランスクリプトが存在する場所
bin/deposition.sh の取り込み
bin/deposition.sh クエリ「データベースはもう決めましたか?」
複数のルートはコロンで区切られます。
DEPOSITION_TRANSCRIPT_ROOTS="/path/one:/path/two" 。
まずはゼロセットアップで試してみませんか?代わりにバンドルされたサンプルを指定します。
import DEPOSITION_TRANSCRIPT_ROOTS= " $( pwd ) /examples/sample-t

記録」
bin/deposition.sh 取り込み && bin/deposition.sh クエリ「データベース」 --決定
定期的に取り込みを実行するように cron ジョブ (または任意のスケジューラ) を設定します。
増分的で安全に数分ごとに実行できます。変更されていないファイルはスキップされます。
すべての設定は環境変数であり、例外を除いてすべてオプションです。
DEPOSITION_TRANSCRIPT_ROOTS :
取り込みは DEPOSITION_TRANSCRIPT_ROOTS を歩き、.jsonl トランスクリプトを検索します
ファイルとグループの連続した会話は、最大 1200 文字の塊になります。
各チャンクは、意思決定マーカー正規表現 (英語および英語) を使用してスキャンされます。
ポルトガル語のパターンが含まれています)。一致する文は
has_decion フラグの横にある決定フィールド。
チャンクはローカルに埋め込まれます (ONNX MiniLM-L6-v2、ネットワーク呼び出しなし)。
ディスク上の Chroma コレクションに更新挿入されます。
取り込みは増分です。各ファイルのバイト オフセットが追跡されるため、再実行されます。
前回の実行以降に追加されたもののみを処理します。
クエリはコレクションに対してコサイン類似度検索を実行します (オプション)
has_decion = true ( --decions ) またはテーマにフィルタリング
( --theme x 、 DEPOSITION_THEMES_FILE を構成した場合)。
ホスト型サービスではありません。サインアップする必要はありません。
リッチ/自己編集型メモリ アーキテクチャではない -- 場合は Letta/MemGPT を参照してください。
それが欲しいです。
単一のエージェント ベンダー向けに調整されていません。一般的なクロードスタイルです。
トランスクリプト JSONL ( {"type": "user"|"assistant", "message": {...}} );もし
ツールの形式は異なります。bin/deposition_common.py が 1 つのファイルです。
適応する。
早いです。コア ループ (取り込み、クエリ、検証) は、エンドツーエンドでテストされます。
合成転写物 (examples/ を参照)。メンテナンスツール
(重複排除、完全な再埋め込みを行わない決定正規表現の再処理) が存在します
元の内部バージョンにあり、まだ移植されていません -- 問題をオープンしてください
もっと早く必要な場合。 Claude Code プラグイン (ワンコマンドでインストール、自動)
フック

手動 cron ジョブの代わりに) が計画されていますが、まだ構築されていません。
「テレメトリ イベントの送信に失敗しました ... Capture() は 1 つの位置引数をとりますが、一部の chromadb/posthog パッケージの stderr で 3 つが指定されました」と表示されます。
組み合わせ。これは chromadb のテレメトリ クライアントがローカル呼び出しに失敗し、
独自の例外をキャッチ -- ネットワーク呼び出しではないことを確認しました (
この依存関係チェーンでオーバーライドされた Capture() は何も行われません)。化粧品のみ。
chromadb の上流で追跡されますが、Deposition のコードが原因または可能性のあるものではありません
外側から完全に抑え込みます。無視しても安全です。
AI エージェント セッションのローカルのみの意思決定ログ -- API 呼び出しはなく、決定論的な意思決定抽出。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Local-only decision log for AI agent sessions -- zero API calls, deterministic decision extraction. - georgedagher/deposition

GitHub - georgedagher/deposition: Local-only decision log for AI agent sessions -- zero API calls, deterministic decision extraction. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
georgedagher
/
deposition
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows bin bin examples examples hooks hooks skills/ deposition-recall skills/ deposition-recall .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh View all files Repository files navigation
A local-only decision log for AI coding agents. It doesn't remember
everything you said — it remembers what you decided , and hands it back
next session, for free, with nothing leaving your machine.
If you use an AI coding agent daily, you've hit at least one of these:
You re-litigate the same call. You spent 20 minutes convincing your
agent to use Postgres over Mongo. New session, new context window — it
scaffolds Mongo again, like the conversation never happened. You explain
yourself a second time. Maybe a third.
"I already told you that." Three sessions ago you said "we're not
using library X, it's unmaintained." Context got compacted, the session
ended, whatever — it's gone, and it comes back up as a live suggestion.
General memory tools are too loud to read. You tried a memory plugin
that "remembers everything." Now every session opens with a wall of
recalled tool calls, file reads, and chatter. You stop reading it. The
memory exists; it doesn't help.
Nobody remembers why, including you. A teammate (or future-you) asks
"why did we go with X over Y?" and the honest answer is a 20-minute dig
through chat history and PR descriptions, for something that was decided
in one sentence weeks ago.
You can't send this transcript anywhere. The conversation has
client code, internal names, or just stuff you don't want summarized by
someone else's LLM API, even to get memory back. Most memory tools
require exactly that call.
Deposition is built around the observation that almost all of the value in
"remembering the session" is really just remembering the handful of
decisions made in it — not the tool calls, not the back-and-forth, not the
dead ends. Catch those, skip the rest, and you get 90% of the benefit at
close to 0% of the noise and cost.
$ deposition.sh query "what database did we pick" --decisions
=== RECALL: "what database did we pick" ===
[1] 2026-08-14 · session demo-session/session · sim=0.55 · themes: -
DECISION: After weighing consistency needs, we decided to go with
Postgres for the new service, since we need strong relational guarantees.
source: demo-session/session.jsonl
That's a real run against the sample transcript in examples/ — try it
yourself, it takes about a minute (see Quick start below). No API key, no
account, no network call.
Most agent-memory tools call an LLM to summarize every observation, then
store the result in a vector database (often with an optional cloud sync).
That's powerful, but it costs money on every write, and it tends to surface
a lot of context you didn't actually need back.
Deposition takes the opposite bet: skip the summarization call, skip trying
to remember everything, and just catch the sentences that look like a
decision as they go by.
Deposition vs. the alternatives
Deposition
claude-mem
Mem0 / Zep / Letta
API calls to run it
None
Yes (LLM summarization)
Yes (extraction/graph)
What it stores
Decisions only
Everything, compressed
Everything (facts/graph)
Setup
1 env var
Plugin install + worker
API key + service
Moving parts
2 scripts
Hooks + worker + DB + UI
Hosted or self-hosted service
Good for
"Stop re-deciding things"
Rich session continuity
Long-running agent memory
Not a ranking — different jobs. If you want an agent with deep, self-editing
long-term memory, claude-mem/Letta/Zep/Mem0 are the more capable, more
mature tools and you should use them. Deposition is for the narrower job:
never re-litigate a decision, with zero cost and zero data leaving your
machine.
Quick start: as a Claude Code plugin (recommended)
/plugin marketplace add georgedagher/deposition
/plugin install deposition
That's it -- no env var to set (it defaults to ~/.claude/projects , Claude
Code's own transcript directory), no cron job. A Stop hook indexes new
decisions after each turn, and a bundled skill teaches Claude to check
Deposition before assuming something was never discussed or re-proposing a
choice you already made.
Verification status: claude plugin validate passes for both the plugin
and marketplace manifests, and a real claude plugin marketplace add +
claude plugin install deposition@deposition completes cleanly with the
plugin showing enabled -- confirming the manifest, marketplace listing, and
hook registration are all structurally correct (this caught and fixed a real
bug: an earlier version double-declared the hooks file and failed to load).
The underlying ingest / query commands are tested end-to-end (see below).
Actually watching the Stop hook fire and the skill get invoked inside a live
conversation has not been verified yet in this environment (no authenticated
interactive session available where this was built) -- if you hit an issue
with that specifically, please open one.
git clone https://github.com/georgedagher/deposition.git && cd deposition
./install.sh # installs uv if needed
export DEPOSITION_TRANSCRIPT_ROOTS= " $HOME /.claude/projects " # where your agent's .jsonl transcripts live
bin/deposition.sh ingest
bin/deposition.sh query " did we decide on a database yet? "
Multiple roots are colon-separated:
DEPOSITION_TRANSCRIPT_ROOTS="/path/one:/path/two" .
Want to try it with zero setup first? Point it at the bundled sample instead:
export DEPOSITION_TRANSCRIPT_ROOTS= " $( pwd ) /examples/sample-transcripts "
bin/deposition.sh ingest && bin/deposition.sh query " database " --decisions
Set up a cron job (or any scheduler) to run ingest periodically -- it's
incremental and safe to run every few minutes; unchanged files are skipped.
All configuration is environment variables, all optional except
DEPOSITION_TRANSCRIPT_ROOTS :
Ingest walks DEPOSITION_TRANSCRIPT_ROOTS , finds .jsonl transcript
files, and groups consecutive dialogue turns into ~1200-character chunks.
Each chunk is scanned with a decision-marker regex (English and
Portuguese patterns included); matching sentences are stored as a
decisions field alongside a has_decision flag.
Chunks are embedded locally (ONNX MiniLM-L6-v2, no network call) and
upserted into a Chroma collection on disk.
Ingestion is incremental: each file's byte offset is tracked, so re-runs
only process what was appended since the last run.
Query does a cosine-similarity search over the collection, optionally
filtered to has_decision = true ( --decisions ) or a theme
( --theme x , if you configured DEPOSITION_THEMES_FILE ).
Not a hosted service -- there's nothing to sign up for.
Not a rich/self-editing memory architecture -- see Letta/MemGPT if you
want that.
Not tuned for any single agent vendor. It reads generic Claude-style
transcript JSONL ( {"type": "user"|"assistant", "message": {...}} ); if
your tool's format differs, bin/deposition_common.py is the one file to
adapt.
Early. The core loop (ingest, query, verify) is tested end-to-end against
a synthetic transcript (see examples/ ). Maintenance tooling
(deduplication, decision-regex reprocessing without a full re-embed) exists
in the original internal version and hasn't been ported yet -- open an issue
if you need it sooner. A Claude Code plugin (one-command install, automatic
hooks instead of a manual cron job) is planned but not built yet.
You'll see Failed to send telemetry event ... capture() takes 1 positional argument but 3 were given in stderr on some chromadb/posthog package
combinations. This is chromadb's telemetry client failing a local call and
catching its own exception -- verified it is not a network call (the
overridden capture() in this dependency chain is a no-op). Cosmetic only;
tracked upstream in chromadb, not something Deposition's code causes or can
fully suppress from the outside. Safe to ignore.
Local-only decision log for AI agent sessions -- zero API calls, deterministic decision extraction.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
