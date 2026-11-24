---
source: "https://harny.dev/"
hn_url: "https://news.ycombinator.com/item?id=50625336"
title: "Show HN: Harny – AI coding agent for VSCodium/VS Code tuned for small/local LLMs"
article_title: "Harny — AI Coding Agent inside VSCodium & VS Code"
image: "https://harny.dev/assets/img/harny_wall.jpg"
author: "spltek"
captured_at: "2026-11-24T08:41:56Z"
capture_tool: "hn-digest"
hn_id: 50625336
score: 1
comments: 0
posted_at: "2026-11-24T08:26:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Harny – AI coding agent for VSCodium/VS Code tuned for small/local LLMs

- HN: [50625336](https://news.ycombinator.com/item?id=50625336)
- Source: [harny.dev](https://harny.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-11-24T08:26:49Z

## Translation

タイトル: Show HN: Harny – 小規模/ローカル LLM 向けに調整された VSCodium/VS Code 用の AI コーディング エージェント
記事のタイトル: Harny — VSCodium および VS Code 内の AI コーディング エージェント
説明: Harny は、アクティビティ バーに常駐する完全な AI コーディング エージェントです。分析でプレビューし、ビルドで適用し、独自のモデルをクラウドまたはローカルに持ち込んでください。 1 つのプロトコル: OpenAI Chat Completions。

記事本文:
コンテンツにスキップ
ハーニー
LLM AI ハーネス
デモ
なぜ
モード
パラレル
コンテキスト
ツール
見る
ガード
MCP
スキル
プロバイダー
インストール
ハーニーをゲット
SYS.HARNY_CORE // オンライン
SPLTEKが開発
CORR_HARVESTER · アクティブ
--:--:-- ローカル
SPLTEK・LLM AIハーネス
アクティビティ バーに常駐する完全なエージェント。プロジェクトを読み取り、ファイルを編集し、
ツール呼び出しを使用して、コマンドを実行し、Web を検索し、フォルダーごとに 1 つの永続的な会話を維持します。
そしてライブストリーミング。
ご自身のモデルをご持参ください。クラウド、またはラップトップ上のモデル。 1 つのプロトコル:
OpenAI チャット完了 API。
Open VSX からインストールする
VS コード マーケットプレイス
GitHub のソース
ツール 13
プロバイダー 8
単体テスト 3000+
MITライセンス
ハーニー
? + ⚙
▸
セッション
? + ⚙
ハーニー・デヴ ×
src/llm/http.ts でバックオフを使用して再試行し、スイートを実行します。
ビルド · deepseek-v4.1-フラッシュ · 6.9k 入力 · 74 出力 · ⚡ 6.1k ·
経過 2.3 秒 · 最初のトークン 676 ミリ秒
▸ 考え中…
このパスでは再試行は行われません。429 または 5xx はユーザーに直接返されます。ループは次のようにする必要があります
クライアントに留まるので、予算は私たちのものとなり、ターンの期限前に停止する必要があります。
▾
読む
src/llm/http.ts
入力
{
"パス" : "src/llm/http.ts" ,
「オフセット」 : 42 、
「限界」：9
}
出力
非同期リクエスト (url: string 、init: RequestInit ) {
for (試行 = 0 ; ;試行++) {
const res = フェッチを待ちます (url, init);
if (res.ok || ! isRetryable (res.status)) return res;
await sleep ( Math . min ( 2000 , 250 * 2 ** 試行));
}
▸ シェル
npmテスト
テスト ファイル 133 が合格しました ( 133 )
テスト 2988 に合格 ( 2988 )
継続時間 20.39 秒
ビルド · 257.4k イン · 775 アウト · ⚡ 257.3k · 経過 9.8 秒
フォーカス すべて折りたたむ フォローする
プロンプトをキューに挿入
ハーニーに聞いて… (Enter で送信、Esc で中止)
ビルドする
⌄
パラレル
AI モデル: deepseek-v4.1-flash · 1M · out 32k ⌄
⚙
トークン
259.3k
/616k/1M
(42%)
⚡ 259.2k
(100%)
コンパクト
自動80％ctx・キープ

10%
⚙
デモ / 1 セッション
あなたが読んでいるサイトは、
エージェントによって構築されました。
「ハーニーは独自のウェブサイトを作成します」 - このページを作成したセッションが記録されています
端から端まで。
遊ぶ
youtu.be/DKJm4pCaEIs
·
@splteklab
OPENAI ◆ ANTHROPIC ◆ OPENROUTER ◆
OLLAMA ◆ LM STUDIO ◆ LLAMA.CPP ◆
MLX ◆ OPENCODE GO ◆
/v1/チャット/完了 ◆
01 / なぜハーニーなのか
道具のように作られており、
デモではありません。
Harny におけるすべての設計上の決定は、リポジトリに書き込むエージェントの障害モードに答えます。
尋ねます。これがあなたに買うものです。
2 段階: プレビュー、その後適用
Analyze では、すべてのファイル変更ツールが予行演習で実行され、
正確な差分を取得し、何も書き込みません。何かを触る前に何が変わるかがわかります。に切り替えます
幸せなときに構築します。
OpenAI、Anthropic、OpenRouter、または Ollama、LM Studio、または llama.cpp を使用した独自のマシン上のモデル。
セッションごとに 1 つのドロップダウン。
Parallel をオンにすると、エージェントはオーケストレーターになり、タスクを次のように分割します。
独立したサブタスクを作成し、サブエージェントとして実行します。それぞれに独自のストリームと独自の停止ボタンがあります。
動作している間は入力を続けてください。キューに入る、⌘ / Ctrl + ↵ ステアリング
飛行中。 Esc はすべてを停止します。
コンテキストがいっぱいになると、Harny はターン境界で圧縮し、最近の部分をそのまま保持し、
作業の状態を前進させます。
すべてのターンで実際のトークンの使用状況とタイミングが表示されます。すべてのツール呼び出しには、その入力、出力、およびその出力が表示されます。
差分。推論は起こるごとに流れます。フォーカスはエージェントとしてすべてのセクションを開きます
入って開いたままにします。ランは上から下まで読める状態のままで、方向を変えることができます。
何かが間違っているように見える瞬間。 Collapse all はその反対の言葉です。
あなたの歴史はプロジェクトの外に存在します
セッションは、リポジトリではなく、中央の SQLite データベースに保存されます。エージェントツールは決して読み取ることができません。
破損したトランスクリプト、

プロジェクトをクリーニングしても、チャット履歴は削除されません。
実行はそのモードで生成され、最後のツール結果まで維持されます。薬は代わりにその理由を説明します
クリック音を飲み込む — 停止は中断する方法です。
新しいセッションは Analyze で開始されます。最初の実行はプレビューであり、驚くことではありません。
書き込みのみが保留されます。読み取り、grep、および検索は引き続き実行されます。
プレビューの準備ができると、ホストはビルドに切り替えるかどうかを一度尋ねます。
実行中はモードとパラレルがロックされます。
ハーニー
? + ⚙
▸
セッション
? + ⚙
ハーニー・デヴ ×
今日のリトライガードはどうなっているでしょうか?
分析 · deepseek-v4.1-flash · 4.1k 入力 · 26 出力 · 経過 1.1 秒
▸ 読む
src/llm/http.ts
if (res.ok) は res を返します。
新しいエラーをスローします ( `HTTP ${res.status}` )
retry-with-backoff を src/llm/retry.ts に追加します。
分析 · deepseek-v4.1-flash · 6.2k 入力 · 41 出力 ·
1.4秒経過
▾
書く
src/llm/retry.ts
+24 −0
入力
{
"パス" : "src/llm/retry.ts" ,
"コンテンツ" : "エクスポート const isRetryable = …"
}
差分
@@ -0,0 +1,22 @@
分析します · 何も書き込まれません — 1 つのファイルがビルドを待機します
フォーカス すべて折りたたむ フォローする
プロンプトをキューに挿入
ハーニーに聞いて… (Enter で送信、Esc で中止)
分析する
⌄
パラレル
AI モデル: deepseek-v4.1-flash · 1M · out 32k ⌄
⚙
トークン
6.2k
/ 616k/1M (1%)
コンパクト
自動 80% ctx · 10% 維持
⚙
ビルドに切り替えますか?
変更はプレビューとして準備されており、まだディスクには何も書き込まれていません。
ビルドに切り替えて適用します。
ヒント: チャットを再開します。クリーンなビルド コンテキストは、この分析トランスクリプトを続行するよりも確実に適用されます。
互いに素なファイルは並行して実行されます。重複するライター、または同じリソースのライターとリーダーは、
すべてのデリゲートがリソースを正直に宣言するため、自動的にシリアル化されます。
src/sessions/compactor.ts の vitest カバレッジを作成する
README.md の圧縮セクションとテーブルを更新
tsconfig strを修正

src/core/llm の ict エラー
Stop は実行全体を中止します。 ⏹ 1 つのバブルを停止します — キャップの後ろで実行または待機しています。
停止したサブエージェントは aborted を報告するため、オーケストレーターはそれを完了したものとして扱いません。
サブエージェントは、ループ ガード、ツール バジェット、およびスキル カタログを継承します。
実際に実行されている上の図: 1 つのオーケストレーター、3 つのデリゲート、重複するファイル
オンザフライで連載中。
遊ぶ
ハーニー並列解析
·
@splteklab
04 / ロングセッション
窓は天井です。
崖ではありません。
トランスクリプトにはコンテキスト ウィンドウが表示されます。区切り線はコンテキストの開始位置を示し、
チェックポイントは、概要が添付された 1 枚のカードです。
カットは完全なターンの間にあり、入力→最終的な答え、その間のすべての要素が含まれます。
そのターンだけで予算を超えた場合にのみ、そのターン内でのみ。
トリガーのデフォルトは、使用可能なコンテキストの 80% です。最新の共有はそのまま残ります。
2 つの独立したノブ、計画の正直なプレビュー。
後の圧縮では、要約を要約するのではなく、前の要約が更新されます。実際には、
世代ごとに失われない。
古いツールの出力は、ツールに名前を付ける 1 行のスタブとして表示されます。成績証明書、店舗および支店
すべてを保管してください。
コンテキスト長エラーは 1 回の圧縮と再試行をトリガーします。プロバイダーが示す制限は次のとおりです。
セッションウィンドウとして採用されました。
配線前に 8 つのルールがチェックされます。違反は毒物ではなく名指しで拒否される
その後のリクエストごとに。
完了: src/llm/retry.ts の withRetry() がエクスポートされました。
開く: サイトを呼び出します — http.ts:42 、 models.ts:88 。
次に、 npm test を実行します。2 回の失敗が予想されます。
事実: 上限は 2000 ミリ秒。 429 および 5xx で再試行します。
証拠 / 2 つの請求期間
同じサブスクリプション、同じモデル、同じ種類の作業を 2 か月間継続します。たった一つのこと
変更されたのは、各リクエストにどれだけのコンテキストが適用されるかです。
安価なモデルではない

。どちらの期間も ~100% DeepSeek V4.1 Flash : トークン
ダウンしましたが、モデルはダウンしませんでした。
複合を停止するウィンドウ。ターンごとにコンテキスト全体が再送信されるため、同じ
100,000 の場合、1 ターン目と 40 ターン目では同じトークンがかかります。代わりにウィンドウに概要が表示されます
吸収された回転数に加えて、一行のスタブにマスクされた古いツールの出力により、
次のリクエストは、より大きいものではなく、より小さいものになります。
アドバイスはメカニズムです。貼り付けられたログ、ファイル全体、および再読み取りによって、
窓。チェックポイントは事実を保持し、サイクルは 94% で終了します。
116% でオーバーシュート。
上のメーターは実際に満たされています: 完全なターンの間のカット、チェックポイント カード、
ウィンドウが進みます。
遊ぶ
調和のとれたコンテキストの圧縮
·
@splteklab
05 / ハーネス
すべてのツールは、小さなインターフェイスの背後にある独自のファイル内のクラスです。エージェントを追加すると、エージェントが登録されます。
そしてプロバイダーは決して変わりません。 MCP サーバーは mcp__<server>__<tool> として到着します。
小さなモデルは推論とツール呼び出しをループします。ループは飲み込むべきエラーではありません。
モデルを押し出すナッジを取得します。
12 個の調整可能なパラメータ · セッションごとに永続化
ガードには独自のセクションがあります →
作品はページにあります。
それ以上ではありません。
ハーニーはこれから何をしようとしているのかを説明し、それを目の前で実行します。答え、ツール、そして
コントロールは上から下まで 1 つの列を共有します。読んでいる行の上には何も表示されません。
パネルが手の下に飛び込むことはありません。
再試行が開始され、パーサーがそれを受け取ります。スイートを実行してから、変更ログ エントリを実行します。
フォローはハーニーが働いている間のみ存在します。何も実行されていないときは、電源が切れて暗くなります。
何も到着していないので、従うべきものはありません。実行中は常に最新の状態を維持します
視野内のライン。オフにすると、置いた場所に留まり、元に戻ることはありません。唯一のことは

で
指を離すと会話が戻ります。スイッチひとつ、ひとつ
行動、驚くべきジャンプはありません。
フォーカスはセクションを開きますが、セクションを閉じることはありません。ハーニーが入るセクションは、
カーソルの下に何も閉じず、フォーカスをオンにしても何も折りたたまれません。
ロットを閉じるには、一言で「すべてを崩壊」します。
ただ止まることはありません。
それは操縦します。
小規模で局所的なモデルは 3 つの形で失敗します: 自分の尻尾を追う推論、同じツールの呼び出し
逐語的に再試行され、答えが繰り返されます。ガードは流れを遮断し、
ナッジ — モデルがループから抜け出すためのループを対象とした修正
シャットダウンされる代わりに。
範囲内の 4 つのプリセット ポイント
12 個のノブをクランプし、セッションごとに編集
単なる検出ではなく重複排除。同一の失敗したコールは再実行されずに拒否され、そのコールは
以前の結果が返されます。同じ引数を使用して同じ呼び出しが再試行されると、モデルは次のようになります。
Glob 、 Grep 、 Read — 古典的な Shell ループはそこで終了します。
ループの両方の読み取り値が捕捉されます: 連続 (A/A/A) と交互 (A/B/A/B、サブエージェント)
Web ループ)、ローリング ウィンドウでは近隣の出現だけでなく出現もカウントされるためです。
回答も保護できます。独自のしきい値を使用してオプトインするため、最終的な返信は開始されます。
繰り返しは最後までお金を払うのではなく、短くなります。
サブエージェントはセッションのガードを継承し、実行のたびに同じオプションが読み取られます。
モデルはどのバブルでも不安定です。
一度インストールしてください。
会話ごとの範囲。
サーバーはマシンに関する事実です。会話で使用できるツールはその会話です
決断。すべてのツールは mcp__<server>__<tool> として提供されますが、どれも実行されません
言われない限り分析してください。
stdio — コマンド、引数、環境 — または http — URL、ヘッダー。から追加されました
パネル

または、 config.json に直接アクセスし、ステータスとともにリストされ、クラッシュ後に再接続されます。
削除されずにオフになりました。
新しい会話は、すべての MCP サーバーがオフの状態で開始されます。このタスクに必要なもののスイッチをオンにし、
隣の会話に触れることなく、サーバーまたは組み込みのツールを 1 つだけオフにします。
Analyze の読み取り専用は除外リストではなく包含リストです。宣言したツールのみが実行できます。
そこに。サーバーが明日追加するツールは、プレビュー中にサイレントに書き込むことは許可されていません。
宣言は他のものと同様、セッションごとに行われます。
// config.json — インストールはマシンレベルであり、スコープはセッションごとです
{
"mcpサーバー": {
"ファイル" : {
"コマンド": "npx",
"args": ["-y", "@modelcontextprotocol/サーバーファイルシステム", "."],
// ⟵ このサーバーの唯一のツールが分析で許可されています
"analyzeTools": ["read_file", "list_directory"]
}
}
}
09 / スキル
プロンプトではなくカタログです。
オンデマンドで読んでください。
スキルは、 .agents/skills 内にある SKILL.md のあるフォルダーです。
.claude/skills または .pi/skills — プロジェクトまたはユーザー レベル。システムプロンプトには次のようなメッセージが表示されます。
その名前と説明のみ。ボディはスキルがアクティブ化される前ではなく、アクティブ化されたときに読み取られます。
プロンプトの開始時に名前で呼ばれます。 /skill <名前> <タスク> を注入します
タスクに先立ってそのスキルの指示が行われます。モデル自体のアクティブ化の決定はバイパスされます。
したがって、「こうしてください」は指示であり、希望ではありません。
またはモデルが単独でそれを手に取ります: カタログ

[切り捨てられた]

## Original Extract

Harny is a full AI coding agent that lives in your Activity Bar. Preview in Analyze, apply in Build, bring your own model — cloud or local. One protocol: OpenAI Chat Completions.

Skip to content
HARNY
LLM AI HARNESS
Demo
Why
Modes
Parallel
Context
Tools
Watch
Guard
MCP
Skills
Providers
Install
Get Harny
SYS.HARNY_CORE // ONLINE
DEVELOPED BY SPLTEK
CORR_HARVESTER · ACTIVE
--:--:-- LOCAL
SPLTEK · LLM AI HARNESS
A full agent that lives in your Activity Bar . It reads your project, edits files,
runs commands, searches the web and keeps one persistent conversation per folder — with tool calling
and live streaming.
Bring your own model. Cloud, or a model on your laptop. One protocol: the
OpenAI Chat Completions API.
Install from Open VSX
VS Code Marketplace
Source on GitHub
Tools 13
Providers 8
Unit tests 3000+
License MIT
HARNY
? + ⚙
▸
Sessions
? + ⚙
Harny Dev ×
Retry-with-backoff in src/llm/http.ts , then run the suite.
Build · deepseek-v4.1-flash · 6.9k in · 74 out · ⚡ 6.1k ·
elapsed 2.3s · first token 676ms
▸ Thinking…
No retry on this path: a 429 or a 5xx goes straight back to the user. The loop has to
stay on the client so the budget is ours, and it has to stop before the turn's deadline.
▾
Read
src/llm/http.ts
Input
{
"path" : "src/llm/http.ts" ,
"offset" : 42 ,
"limit" : 9
}
Output
async request (url: string , init: RequestInit ) {
for ( let attempt = 0 ; ; attempt++) {
const res = await fetch (url, init);
if (res.ok || ! isRetryable (res.status)) return res;
await sleep ( Math . min ( 2000 , 250 * 2 ** attempt));
}
▸ Shell
npm test
Test Files 133 passed ( 133 )
Tests 2988 passed ( 2988 )
Duration 20.39 s
Build · 257.4k in · 775 out · ⚡ 257.3k · elapsed 9.8s
Focus Collapse all Follow
Insert prompt in QUEUE
Ask Harny… (Enter to send, Esc to abort)
Build
⌄
Parallel
AI Model: deepseek-v4.1-flash · 1M · out 32k ⌄
⚙
Tokens
259.3k
/ 616k/1M
(42%)
⚡ 259.2k
(100%)
Compact
Auto 80% ctx · keep 10%
⚙
DEMO / ONE SESSION
The site you are reading,
built by the agent.
“Harny creates its own website” — the session that produced this page, recorded
end to end.
Play
youtu.be/DKJm4pCaEIs
·
@splteklab
OPENAI ◆ ANTHROPIC ◆ OPENROUTER ◆
OLLAMA ◆ LM STUDIO ◆ LLAMA.CPP ◆
MLX ◆ OPENCODE GO ◆
/v1/chat/completions ◆
01 / WHY HARNY
Built like a tool,
not a demo.
Every design decision in Harny answers a failure mode of agents that write into your repo without
asking. Here is what that buys you.
Two stages: preview, then apply
In Analyze , every file-modifying tool runs in dry-run — it returns the
exact diff and writes nothing. You see what would change before anything is touched. Switch to
Build when you are happy.
OpenAI, Anthropic, OpenRouter, or a model on your own machine with Ollama, LM Studio or llama.cpp.
One dropdown, per session.
Turn on Parallel and the agent becomes an orchestrator: it splits the task into
independent sub-tasks and runs them as sub-agents — each with its own stream and its own stop button.
Keep typing while it works. Enter queues, ⌘ / Ctrl + ↵ steers
mid-flight. Esc stops everything.
When the context fills, Harny compacts at turn boundaries , keeps the recent part verbatim and
carries the state of the work forward.
Every turn shows real token usage and timings. Every tool call shows its input, its output and its
diff. Reasoning streams as it happens. Focus opens every section as the agent
enters it and leaves it open — the run stays readable top to bottom, and you can steer the
moment something looks wrong. Collapse all is the one word for the opposite.
Your history lives outside your project
Sessions are stored in a central SQLite database — not in your repo. Agent tools can never read or
corrupt transcripts, and cleaning your project never deletes your chat history.
A run is born with its mode and keeps it to its last tool result. The pills explain why instead of
swallowing the click — Stop is the way to interrupt.
New sessions start in Analyze — your first run is a preview, not a surprise.
Only the write is withheld: reads, greps and searches still run.
When previews are ready, the host asks once whether to switch to Build.
Mode and Parallel are locked while a run is going.
HARNY
? + ⚙
▸
Sessions
? + ⚙
Harny Dev ×
How does the retry guard look today?
Analyze · deepseek-v4.1-flash · 4.1k in · 26 out · elapsed 1.1s
▸ Read
src/llm/http.ts
if (res.ok) return res;
throw new Error ( `HTTP ${res.status}` )
Add retry-with-backoff to src/llm/retry.ts .
Analyze · deepseek-v4.1-flash · 6.2k in · 41 out ·
elapsed 1.4s
▾
Write
src/llm/retry.ts
+24 −0
Input
{
"path" : "src/llm/retry.ts" ,
"content" : "export const isRetryable = …"
}
Diff
@@ -0,0 +1,22 @@
Analyze · nothing written — 1 file awaits Build
Focus Collapse all Follow
Insert prompt in QUEUE
Ask Harny… (Enter to send, Esc to abort)
Analyze
⌄
Parallel
AI Model: deepseek-v4.1-flash · 1M · out 32k ⌄
⚙
Tokens
6.2k
/ 616k/1M (1%)
Compact
Auto 80% ctx · keep 10%
⚙
Switch to Build?
Changes were prepared as previews — nothing has been written to disk yet.
Switch to Build to apply them.
Tip: restart the chat — a clean Build context applies more reliably than continuing this Analyze transcript.
Disjoint files run in parallel. Overlapping writers — or a writer and a reader of the same resource — are
serialized automatically, because every Delegate declares its resources honestly.
Write vitest coverage for src/sessions/compactor.ts
Update README.md compaction section + table
Fix tsconfig strict errors in src/core/llm
Stop aborts the whole run. ⏹ stops one bubble — running or queued behind the cap.
A stopped sub-agent reports aborted , so the orchestrator never treats it as done.
Sub-agents inherit the loop guard, the tool budget and the skills catalog.
The diagram above, running for real: one orchestrator, three delegates, overlapping files
serialized on the fly.
Play
Harny Parallel Analysis
·
@splteklab
04 / LONG SESSIONS
The window is a ceiling.
Not a cliff.
The context window is visible in the transcript: a divider marks where the context begins, and the
checkpoint is one card with its summary attached.
The cut falls between complete turns — your input → the final answer, everything in between — and
inside one only when that turn alone exceeds the budget.
The trigger defaults to 80% of the usable context; the most recent share stays verbatim.
Two independent knobs, one honest preview of the plan.
A later compaction updates the previous summary instead of summarizing a summary — a fact is
not lost per generation.
Old tool outputs go as one-line stubs that name the tool. The transcript, the store and any branch
keep everything.
A context-length error triggers one compaction and a retry — and the limit the provider states is
adopted as the session window.
Eight rules are checked before the wire. A violation is refused by name , instead of poisoning
every later request.
Done: withRetry() in src/llm/retry.ts , exported.
Open: call sites — http.ts:42 , models.ts:88 .
Next: run npm test , expect 2 failures.
Facts: cap 2000ms ; retry on 429 and 5xx .
EVIDENCE / TWO BILLING PERIODS
Two months of the same subscription, the same model, the same kind of work. The one thing
that changed is how much context rides on every request.
Not a cheaper model. Both periods are ~100% DeepSeek V4.1 Flash : the tokens
went down, the model did not.
A window that stops compounding. Every turn re-sends the whole context, so the same
100k costs the same tokens on turn one and on turn forty. A summary in the window — instead
of the turns it absorbed — plus old tool output masked to one-line stubs is what makes the
next request smaller rather than bigger.
The advice is the mechanism. Pasted logs, whole files and re-reads are what fill a
window; the checkpoint keeps the facts, and the cycle closes at 94% instead of
overshooting at 116% .
The meter above, filling up for real: the cut between complete turns, the checkpoint card, the
window moving on.
Play
Harny Context Compaction
·
@splteklab
05 / THE HARNESS
Every tool is a class in its own file behind a small interface. Adding one is registering it — the agent
and the providers never change. MCP servers arrive as mcp__<server>__<tool> .
Small models loop in their reasoning and in their tool calls. A loop is not an error to swallow: it
gets a Nudge that pushes the model past it.
12 tunable parameters · persisted per session
The guard has a section of its own →
The work is on the page.
Not over it.
Harny tells you what it is about to do, then does it in front of you. The answer, the tools and
the controls share one column, top to bottom: nothing floats over the lines you are reading,
and the panel never jumps under your hand.
The retry is in and the parser takes it now. Running the suite, then the changelog entry.
Follow only exists while Harny is working. When nothing is running it is off and dark:
nothing is arriving, so there is nothing to follow. During a run, on, it keeps the newest
line in view; off, it stays where you left it and never pulls you back. The only thing that
can hold it is your finger on the conversation — let go, and it comes home. One switch, one
behaviour, no surprise jumps.
Focus opens sections — it never closes them. The section Harny walks into opens on its
own and stays open: nothing shuts under your cursor, and turning Focus on folds nothing away.
To close the lot there is one word, Collapse all .
It never just stops.
It steers.
Small and local models fail in three shapes: reasoning that chases its own tail, the same tool call
retried verbatim, an answer that repeats itself. The guard cuts the stream and injects a
Nudge — a correction aimed at the loop, so the model gets out of it
instead of being shut down.
Four presets points inside the ranges
Twelve knobs clamped · edited per session
Dedup, not just detection. An identical failing call is refused instead of re-run and its
previous result comes back; the same call retried with identical arguments points the model at
Glob , Grep and Read — the classic Shell loop ends there.
Both readings of a loop are caught: back-to-back (A/A/A) and alternating (A/B/A/B, the sub-agent
web loop), because a rolling window counts occurrences, not just neighbours.
The answer can be guarded too — opt-in, with its own threshold, so a final reply that starts
repeating is cut short instead of paid for to the end.
Sub-agents inherit the guard of their session, with the same options read at each run: a flaky
model is flaky in every bubble.
Install once.
Scope per conversation.
A server is a fact about the machine; which of its tools a conversation may use is that conversation's
decision. Every tool arrives as mcp__<server>__<tool> — and none of them runs in
Analyze unless you say so.
stdio — command, arguments, environment — or http — url, headers. Added from
the panel or straight into config.json , listed with its status, reconnected after a crash,
and switched off without being deleted.
A new conversation starts with every MCP server off . Switch on the one this task needs, and
switch off single tools — of a server or a built-in — without touching the conversation next door.
Read-only in Analyze is an include list, not an exclusion: only the tools you declare can run
there. A tool the server adds tomorrow is not silently allowed to write during a preview — and the
declaration is per session, like everything else.
// config.json — install is machine-level, the scope is per session
{
"mcpServers" : {
"files" : {
"command": "npx",
"args": ["-y", "@modelcontextprotocol/server-filesystem", "."],
// ⟵ the ONLY tools of this server allowed in Analyze
"analyzeTools": ["read_file", "list_directory"]
}
}
}
09 / SKILLS
A catalog, not a prompt.
Read on demand.
A skill is a folder with a SKILL.md , in .agents/skills ,
.claude/skills or .pi/skills — project or user level. The system prompt carries
only its name and description; the body is read when the skill is activated, and not before.
Called by name, at the start of the prompt. /skill <name> <task> injects
that skill's instructions ahead of your task — the model's own activation decision is bypassed,
so “do it this way” is an instruction, not a hope.
Or the model picks it up alone : the catalog

[truncated]
