---
source: "https://temporalstore.ai/blog-context-management.html"
hn_url: "https://news.ycombinator.com/item?id=49314204"
title: "TemporalStore: A disruptive open-source engine managing your LLM memory"
article_title: "The Context Bill — LLM memory on TemporalStore"
author: "matrixarkai"
captured_at: "2026-08-15T21:12:12Z"
capture_tool: "hn-digest"
hn_id: 49314204
score: 2
comments: 1
posted_at: "2026-08-15T20:51:41Z"
tags:
  - hacker-news
  - translated
---

# TemporalStore: A disruptive open-source engine managing your LLM memory

- HN: [49314204](https://news.ycombinator.com/item?id=49314204)
- Source: [temporalstore.ai](https://temporalstore.ai/blog-context-management.html)
- Score: 2
- Comments: 1
- Posted: 2026-08-15T20:51:41Z

## Translation

タイトル: TemporalStore: LLM メモリを管理する破壊的なオープンソース エンジン
記事のタイトル: コンテキスト法案 — TemporalStore の LLM メモリ

記事本文:
TemporalStore: LLM トークンのコストを削減し、回答の品質を向上させることができる破壊的なオープンソース エンジンです。
オープンソースの Rust ネイティブの時間エンジン。ターンごとに送信するコンテキストを縮小し、答えを鮮明にし、チーム内のすべてのセッション、デバイス、エージェントにわたって記憶します。
通常のアプローチ: 毎ターン、すべてを追加します
エージェントに記憶を与える典型的な方法は、最も単純なものです。新しいターンごとに、
以前のすべてのコンテキスト (以前のプロンプトとモデル自体の断続的な出力)
次の入力ウィンドウに戻って、再度送信します。
動作するので発送します。しかし、これには 2 つの静かな問題があります。
再送信されるもののほとんどはノイズです。追加の歴史の多く、特に
断続的な出力: ツールの結果、ログ、スタック トレース、ファイル ダンプ - 単一のステップで役に立ちました
そして、ウィンドウ内に座って毎ターン再送信され、モデルが現在実際に必要としている信号を弱めます。
通話のたびに料金を支払うことになり、重要な数回線が混雑してしまいます。
そして何も学ばないのです。すべてが 1 つのセッションに閉じ込められており、セッションをまたぐことはありません
記憶がなく、歴史が深まるにつれて改善される長期的なプロファイルもありません。エージェントは決して
ユーザー、プロジェクト、またはすでに行われた決定について把握したことを蓄積します。いつ
ウィンドウがロールするかセッションが終了すると、ゼロから始まります。
したがって、より多くの料金を支払う (増大するノイズの多いトランスクリプトを再送信する) 代わりに、記憶に残るものは少なくなる (長期間にわたって生き残るものは何もない)
セッション）。それが罠です。
毎ターン、すべてを追加
ランク付けされたコンテキストパック
プロンプト
断続出力（ノイズ）
ターン1 ターン20 ターン40
毎ターン再送信 · ほとんどがノイズ · 必要な事実が埋もれている
$$$ が増加 · 1 セッション — 何も学ばなかった
イベント + ツール
一度
テンポラルストア
デノイズ・ランク・プロフィール
~1.3,000 トークン
ノイズは除去されましたが、事実は維持されました
+ クロスセッションプロファイル、改善

深さ
モデルは重要なことだけを認識します — ソースに裏付けられた
すべてを追加すると、増大するノイズの多いトランスクリプト (ほとんどが断続的な出力) が再送信され、セッション間で忘れられてしまいます。 TemporalStore はノイズを除去し、ソースに裏付けられた小さな ContextPack にランク付けし、さらに歴史が深まるにつれて改善されるクロスセッション プロファイルを追加します。
独自のスタックとなる「ソリューション」
そこで、検索を追加します。合理的な本能。しかし、「エージェントのメモリ」が運用環境で静かにどのように変化するかを見てください。
埋め込みを保存、バージョン管理し、意味的に検索するためのベクトル データベース
古いものを圧縮して取得可能なものに変える要約/抽出パイプライン、
そして、それを結合し、範囲とランクを決定するためのメモリ サービス (通常はオーダーメイドのもの)。
システムが 3 つあるのに、サイドでトランスクリプトを再生し続けることになります。実行、監視、
一貫性を保つ - すべては、ターンごとに 1 つの小さな質問に答えるためです。「このエージェントは実際に何をするのか」
今その前に必要ですか？
TemporalStore の背後にある賭けは、エージェントのメモリは 1 つの一時的な問題であり、つなぎ合わせたスタックではなく、1 つの一時エンジンが必要であるということです。
オープン ソース (Apache-2.0)、Rust ネイティブで、単一の Docker コマンドから実行され、以下のものが付属しています。
形容詞の代わりに再現可能なベンチマーク。その仕組みと、数字が得られる理由をご紹介します。
1 つのテンポラル エンジン
取り込み・ランク・ContextPack
1 つの一時インデックスと 1 つの耐久性パス
統合されたメモリ スタック: ベクトル DB、要約パイプライン、およびそれらの接着剤が 1 つの時間エンジンに集約されます。
中心となるアイデア: トランスクリプトではなく ContextPack を送信する
すっきりとした一本のラインを描くデザイン。モデルは推論して抽出します。エンジンが記憶して機能します。
# フロー — あなたがモデルを所有し、TemporalStore がメモリを所有します
生のクエリ / イベント / リソース + ヒント
→ インテント、時間、フィルター、エンティティを抽出 (よ

あなたのランタイム、または MatrixArk)
→ コンパクトなサービングレコードを保存（TemporalStore）
→ トークン予算の ContextPack (TemporalStore) を取得します。
→ 最終的な LLM = ローカル コンテキスト + ContextPack
→答えは明日の思い出になる
TemporalStore が実行しないことに注意してください。LLM や埋め込みモデルは決して呼び出されません。
ハーネス — または MatrixArk などのマネージド ランタイム
一番上にある (matrixark.ai) — ベクトルと構造化されたデータを生成します
意図。 TemporalStore はそれらを保存し、提供時に高速かつ制限された作業を実行します。
この作業により、システムの要点が生み出されます。つまり、コンパクトでソースに裏付けられたシステムです。
コンテキストパック 。トランスクリプトではありません - このために実際に重要な数百のトークン
クエリ、それぞれを生成したイベントまで追跡可能です。
そしてそれは騒音を消します。断続的な出力を再送信したことを覚えていますか? TemporalStore の取り込み
コンパクトなレコードとして保存します。ツール呼び出しは短いマーカーとして保持され、大きな結果は先頭が切り捨てられます。
つまり、パックには、ノイズの多い塊ではなく、ツールが生成したという事実が含まれます。ノイズが少ない
ウィンドウ自体がより良い答えです。モデルは、古いログを調べて、該当する 1 行を見つけません。
重要だった。
得られるものをそれぞれ 1 行で
前もっての見返り — それぞれの背後にあるメカニズムは、この部分の残りの部分にあります。
同等以上の品質でより少ないトークン — 毎ターン成長するトランスクリプトではなく、制限された ContextPack。
より良い答え、より少ない幻覚 — 再生可能な、出典に裏付けられた引用による、時効のある記憶。
数百万のエンティティにわたる低い処理遅延 — コーパスではなく、フィルタリングされた候補セットによって制限された作業。
正確な再現率、ベクトル DB なし - 管理された候補セットが正確にスコア付けされます。 ANN の損失性は決して問題にはなりません。
チーム全体のメモリ - クロスセッション、クロスデバイス、クロスエージェントで、構築により監査可能。
1 つのエンジン、自己ホスト型

— Apache-2.0、1 つの Docker コマンド、追加パスに LSM ライトアンプなし。
小さなパックが「最新の N 個のトークンをリプレイする」よりも優れている理由
これは直観に反する部分なので、主張するのではなく証明しましょう。
実際のワーキング セットを見てみましょう: 開発者の完全なローカル コーディング履歴 (ツール呼び出しを含む) —
12,384 のイベント、97 のセッション、約 170 万のトークン。それについて 6 つの本当の質問をしてください。両腕を走らせる
同じ読者モデルと同じ審査員を通じて:
再生 — 完全なローカル履歴を最新のものから順にモデルのウィンドウにフィードします。
マネージド — トークンで予算化された ContextPack を取得します。
さて、見出しよりも興味深いのは「理由」です。
答えが古い歴史に残っている質問では、パックが最も大差で勝利します。
q_storage 10.0 対 6.05、q_parity 8.05 対 5.25。最小の差で勝ちます
最新のメッセージがたまたま話題になっていたとき。
その勾配が可視化されたメカニズムです。最新性が切り捨てられたリプレイでは文字通り、
検索結果にランク付けされた、古いが関連性のある事実。 170 万トークンを支払っても、
もっと悪い答え。
また、標準のメモリ ベンチマークにも耐えます。ロコモと
LongMemEval_s — 現場でエージェントのリコールを測定する 2 つのパブリック長期記憶スイート
対 — TemporalStore のツリープラスインデックスの取得は hit@k 0.995–1.00 になります。そして、
1 つの共有オープンソース スタック (qwen2.5:7b リーダー、MiniLM 埋め込み、
Claude の判断によると)、OpenViking スタイルのベースラインを結び付ける — の再実装
最先端の技術を報告したメモリ システム、VikingMem
独自の論文で、LoCoMo の短いチャット (~83 ～ 84%) と
長期的な LongMemEval で完全に勝利しました。回答精度は 98% 対 66% でした。それでトークンは
上記の節約は、より悪い思い出を伴って購入されるものではありません。パックは小さく、検索と答えが見つかります。
正しい証拠。
記憶それは

共有 — セッションにスタックしていない
これが実際のチームにとって最も重要な差別化要因ですが、トークン計算では見落としがちです。
トランスクリプトは 1 つのセッションに閉じ込められます。ウィンドウが回転するかタブが閉じると、それは消えます。それは
エージェントのメモリの本当の上限 — 1 つのプロンプトにどれだけ詰め込めるかではなく、何かが生き残るかどうか
プロンプト間、エージェント間、人々間。
TemporalStore はメモリをエンジン内に保持するため、一度取り込まれ、スコープによってインデックス付けされます。
( テナント → ユーザー → セッション → エージェント → リソース ) — 呼び出しは現在のスレッドにバインドされません。アン
エージェントは、他のエージェントが昨日学んだこと、ユーザーが決定したことを思い出すことができます。
別のデバイス、チームが先月合意したもの。そして、思い出されるすべての事実は、
タイムスタンプが付けられ、ソース イベントまで追跡可能であるため、共有メモリは常に監査可能です。
噂ではありません。
クロスセッション、クロスデバイス、クロスエージェント、クロスチームメイト - 構造上。それは、単一転写レトリバーでは構造的に与えることができない記憶です。
サービス提供の革新: メモリを対象としたデータベースのサービス提供パス
ほとんどの「エージェント メモリ」ライブラリは、ベクトル インデックスの薄いラッパーです。 TemporalStore は次のように構築されています
データベース — そして、それはサービス提供パスに最も多く現れます。 3 つの役割を分類すると、次のようになります。
プロキシ — フロントドア: ルーティング、リクエストのバッチ処理、バックプレッシャー。
メタサーバー — 配置、リース、シャード マップ。データがどこに存在するかを決定します。ホットリードパスから外れています。
データノード — ワーカー: 耐久性のための先行書き込みログ、モデルの実行プログラム、読み取りに応答するメモリ内インデックス。
取得は総当たりのベクトル スキャンではありません。コンテキストはツリーです - テナント → ユーザー → セッション
→リソース→エンティティ。クエリは、コンパクトなセカンダリ インデックスを使用して、レイヤーごとにクエリを実行します。
タイムラインを読み取る前にプレフィルターを適用します。フィル

投稿リストを交差させることによってターを構成します
— ステータス ∩ プロジェクト ∩ 時間バケット — したがって、クエリはコーパス全体ではなく、限定されたイベントのセットに影響します。のみ
次に、スコアを付け、一時的な減衰を適用し、予算に合わせてパックします。
メリット: エンティティごとに作業が行われるため、数百万のエンティティにわたってもレイテンシは低く抑えられます。
クエリはデータセットではなく、フィルターされたセットに応じてスケールされます。 ContextPack の取得が実行されます
1.2k トークンの予算で ~17 ミリ秒 p95。
ストレージの革新: 追加構造と一時的なストレージ
その下には、メモリは一時的なものであるという 1 つの事実に基づいて構築されたストレージ エンジンがあります。データが到着
時間順に、時間によってクエリされ、時間とともに冷却されます。それがすべての層を形作ります。
追加構造。書き込みは追加します。その場で何も変化せず、エンジンがリロードされます
独自の先行書き込みログによりクラッシュセーフになります。セグメント化されたログが追加されるたびにそれ自体を再解析したとき (
O(n²) トラップ — 追加 #200 には 100 ミリ秒以上かかりました)、ノードごとのカーソルにより O(1): 109 ミリ秒になりました →
2.2 ミリ秒、ディスク上のフォーマットがバイト同一であるため、ライブ ストアは移行なしでリロードされます。
実際の階層化 - 立ち退きと昇進。ホットデータはメモリ内に残ります。働くものとして
セットは、キャッシュを通じて SSD および共有ストレージに追い出されるバジェットを超えます。ほとんどのシステムの一部
誤解: コールド データは読み取り時にメモリに戻されます。マルチユーザー ハーネスがそれを証明します —
意図的に小さいメモリ バジェットを使用してエビクションを強制し、コールド リードで昇格を強制します。
256k 読み取り、不一致ゼロ、パーティション分離はそのまま、最大 81 ユーザー、187k 書き込み。
冷めたら取り除く→
記憶
ホット・ライブインデックス
PMem
記憶に近い暖かい
SSD
ウォームブロック
共有ストレージ
耐寒性・耐久性
← 読み取り時にメモリにプロモートされて戻される
ホット状態はメモリから提供されます。冷却されると共有ストレージに追い出され、読み取り時にプロモートされて戻ります。これは、耐久性のあるメモリ優先の 1 つの階層です。
コンピューティングとストレージの分離。サ

私のコードは 3 つの方法で実行されます - ローカル
(1 つのノード)、Raft で複製された、または共有ストレージ - したがって、ラップトップで開始します
2 つの自己完結型クラスタ上で、コンピューティングとストレージが独立して拡張する分散クラスタに拡張します。
Apache-2.0 ライブラリ: MatrixCache
(キャッシュ) と MatrixRaft (レプリケーション)。
大規模な共有ストレージ層の場合、ステートレス データノードは 1 つの耐久性のあるオブジェクト ストアの読み取りと書き込みを行います。
MatrixObject 、最高のパフォーマンスの共有ストレージ向けに設計されたエンタープライズ オブジェクト ストア —
そのため、すべてのシャードをすべてのノードにコピーすることなく、トラフィックの急増に合わせてコンピューティングを拡張できます。
RAG は通常、巨大なデータベース上で近似最近傍 (ANN) 検索を実行するベクトル DB 上で実行されます。
埋め込みの平らな山 — すべてのベクトルを正確にスコアリングするのはコストが高すぎるため、速度と再現率を引き換えにします。
TemporalStore は決してそこに到達しません。階層化された仮想ファイル システムとコンパクト
セカンダリ インデックスは、スコープ、タイプ、エンティティ、時間によって、ほとんどの候補を最初にフィルタリングします。
ランク付けの時点では、セットは小さく、ANN の損失はなく、正確にスコアを付けられます (つまり、
hit@k 0.995–1.00 上記）。その同じスコープ付きツリーは、仮想ファイル システムとしても機能します。
リソース、スキル、メモリ - それぞれに個別のストアではなく、1 つの時間認識名前空間 - したがって
また、実行して同期を保つシステムが 1 つ少なくなります。
一時メモリは追加が多く、タイムスタンプをキーにした新しいイベントを常に書き込むことになります。
それはまさに何ですか

[切り捨てられた]

## Original Extract

TemporalStore: a disruptive open-source engine that can cut LLM token cost and improve answer quality.
One open-source, Rust-native temporal engine that shrinks the context you send each turn, sharpens the answers, and remembers across every session, device, and agent on your team.
The usual approach: append everything, every turn
The typical way to give an agent memory is the simplest one: on each new turn, append
all of the prior context — the earlier prompts and the model's own intermittent output —
back into the next input window, and send it again.
It works, so it ships. But it has two quiet problems.
Most of what you re-send is noise. A lot of that appended history — especially the
intermittent output : tool results, logs, stack traces, file dumps — was useful for a single step
and then just sits in the window, re-sent every turn, diluting the signal the model actually needs now.
You pay for it on every call, and it crowds out the few lines that matter.
And nothing is learned. It's all trapped in one session — there's no cross-session
memory, and no long-term profile that gets better as the history deepens. The agent never
accumulates what it figured out about the user, the project, or the decisions already made; when the
window rolls or the session ends, it starts from zero.
So you pay more (re-sending a growing, noisy transcript) and remember less (nothing survives the
session). That's the trap.
APPEND EVERYTHING, EVERY TURN
RANKED CONTEXTPACK
prompt
intermittent output (noise)
turn 1 turn 20 turn 40
re-sent every turn · mostly noise · needed fact buried
$$$ grows · one session — nothing learned
events + tools
once
TemporalStore
denoise · rank · profile
~1.3k tokens
noise dropped · facts kept
+ cross-session profile, improves with depth
the model sees only what matters — source-backed
Appending everything re-sends a growing, noisy transcript — mostly intermittent output — and forgets between sessions. TemporalStore denoises and ranks it into a small, source-backed ContextPack, plus a cross-session profile that improves as history deepens.
The "solution" that becomes its own stack
So you add retrieval. Reasonable instinct. But watch what "agent memory" quietly turns into in production:
a vector database to store, version, and semantically search the embeddings,
a summarization / extraction pipeline to compress old turns into something retrievable,
and a memory service to glue it together, decide scope, and rank — usually a bespoke one.
Three systems, and you're still replaying transcripts on the side. Three things to run, monitor,
and keep consistent — all to answer one small question on every turn: what does this agent actually
need in front of it right now?
The bet behind TemporalStore is that agent memory is one temporal problem — and deserves one temporal engine, not a stitched-together stack.
It's open source (Apache-2.0), Rust-native, runs from a single Docker command, and ships with
reproducible benchmarks instead of adjectives. Here's how it works, and why the numbers land.
One temporal engine
ingest · rank · ContextPack
one temporal index · one durability path
The memory stack, consolidated: a vector DB, a summarization pipeline, and their glue collapse into one temporal engine.
The core idea: send a ContextPack, not a transcript
The design draws one clean line. Your model reasons and extracts; the engine remembers and serves.
# the flow — you own the model, TemporalStore owns the memory
raw query / event / resource + hints
→ extract intent, time, filters, entities (your runtime, or MatrixArk)
→ store compact serving records (TemporalStore)
→ retrieve a token-budgeted ContextPack (TemporalStore)
→ final LLM = local context + ContextPack
→ the answer becomes tomorrow's memory
Notice what TemporalStore doesn't do: it never calls an LLM or an embedding model.
Your harness — or a managed runtime like MatrixArk
( matrixark.ai ) that sits on top — produces the vectors and structured
intent. TemporalStore stores them and does the fast, bounded work at serving time.
That work produces the whole point of the system: a compact, source-backed
ContextPack . Not a transcript — the few hundred tokens that actually matter for this
query, each traceable to the event that produced it.
And it denoises. Remember all that re-sent intermittent output? TemporalStore ingests
it as compact records — the tool call kept as a short marker, the bulky result head-truncated or
summarized — so the pack carries the fact the tool produced, not the noisy blob. Less noise in the
window is itself better answers: the model isn't rummaging through stale logs to find the one line that
mattered.
What you get, in one line each
The payoff up front — the mechanism behind each one is the rest of this piece:
Fewer tokens at equal-or-better quality — a bounded ContextPack instead of a transcript that grows every turn.
Better answers, less hallucination — time-valid memory with replayable, source-backed citations.
Low serving latency across millions of entities — work bounded by the filtered candidate set, not the corpus.
Exact recall, no vector DB — a controlled candidate set scored exactly; ANN's lossiness never enters the picture.
Team-wide memory — cross-session, cross-device, cross-agent, and auditable by construction.
One engine, self-hosted — Apache-2.0, one Docker command, no LSM write-amp on the append path.
Why a small pack beats "replay the newest N tokens"
This is the counterintuitive part, so let's prove it instead of asserting it.
Take a real working set: a developer's full local coding history, tool calls included —
12,384 events, 97 sessions, ~1.7M tokens . Ask six real questions about it. Run two arms
through the same reader model and the same judge:
Replay — feed the full local history, newest-first, up to the model's window.
Managed — retrieve a token-budgeted ContextPack .
Now the why , which is more interesting than the headline.
The pack wins by the widest margin on the questions whose answers live in old history —
q_storage 10.0 vs 6.05, q_parity 8.05 vs 5.25. It wins by the smallest margin
when the newest messages happened to already be on-topic.
That gradient is the mechanism, made visible: recency-truncated replay literally cannot see the
old-but-relevant facts that ranked retrieval surfaces. You pay for 1.7M tokens and still get the
worse answer.
And it holds up on the standard memory benchmarks. On LoCoMo and
LongMemEval_s — the two public long-term-memory suites the field measures agent recall
against — TemporalStore's tree-plus-index retrieval lands hit@k 0.995–1.00 . And in a
head-to-head on one shared open-source stack (a qwen2.5:7b reader, MiniLM embeddings,
Claude-judged), it ties an OpenViking-style baseline — a reimplementation of
VikingMem , the memory system that reported state-of-the-art
in its own paper — on short LoCoMo chats (~83–84%) and
wins long-horizon LongMemEval outright — 98% vs 66% answer accuracy. So the token
savings above aren't bought with worse recall: the pack is small and it finds — and answers —
the right evidence.
Memory that's shared — not stuck in a session
Here's the differentiator that matters most for real teams, and it's easy to miss under the token math.
A transcript is trapped in one session. The moment the window rolls or the tab closes, it's gone. That's
the true ceiling on agent memory — not how much you can stuff into one prompt, but whether anything survives
between prompts, between agents, between people.
Because TemporalStore keeps memory in the engine — ingested once, indexed by scope
( tenant → user → session → agent → resource ) — recall isn't bound to the current thread. An
agent can remember what another agent learned yesterday, what the user decided on a
different device , what the team agreed last month. And because every recalled fact is
timestamped and traceable to its source event, shared memory stays auditable — a citation,
not a rumor.
Cross-session, cross-device, cross-agent, cross-teammate — by construction. That's the memory a single-transcript retriever structurally cannot give you.
The serving innovation: a database's serving path, aimed at memory
Most "agent memory" libraries are a thin wrapper over a vector index. TemporalStore is built like a
database — and that shows most in the serving path. Three roles, disaggregated:
Proxy — the front door: routing, request batching, backpressure.
Metaserver — placement, leases, the shard map. It decides where data lives; it's off the hot read path.
Datanodes — the workers: a write-ahead log for durability, the model executors, the in-memory indexes that answer reads.
Retrieval is not a brute-force vector scan. Context is a tree — tenant → user → session
→ resource → entity. A query walks it layer by layer, using compact secondary indexes to
prefilter before reading any timeline. Filters compose by intersecting posting lists
— status ∩ project ∩ time-bucket — so a query touches a bounded set of events, not the whole corpus. Only
then does it score, apply temporal decay, and pack to budget.
The payoff: latency stays low even across millions of entities , because the work per
query scales with the filtered set, not the dataset. ContextPack retrieval runs
~17 ms p95 at a 1.2k-token budget.
The storage innovation: append-structured and temporal
Underneath is a storage engine built around one fact: memory is temporal. Data arrives
in time order, is queried by time, and cools with age. That shapes every layer.
Append-structured. Writes append; nothing mutates in place, and the engine reloads
crash-safe from its own write-ahead log. When the segmented log once re-parsed itself on every append (an
O(n²) trap — append #200 took 100+ ms), a per-node cursor made it O(1): 109 ms →
2.2 ms , with the on-disk format byte-identical so live stores reload with zero migration.
Real tiering — eviction and promotion. Hot data sits in memory; as the working
set outgrows the budget it's evicted down through cache to SSD and shared storage. The part most systems
get wrong: cold data is promoted back into memory on read. A multi-user harness proves it —
a deliberately tiny memory budget to force eviction, then cold reads to force promotion:
256k reads, zero mismatches , partition isolation intact, up to 81 users and 187k writes.
evict as it cools →
Memory
hot · live index
PMem
near-memory warm
SSD
warm blocks
Shared storage
cold · durable
← promoted back into memory on read
Hot state serves from memory; as it cools it's evicted toward shared storage, and promoted back on read — one durable, memory-first hierarchy.
Compute/storage separation. The same code runs three ways — local
(one node), Raft-replicated , or shared-storage — so you start on a laptop
and scale to a disaggregated cluster where compute and storage grow independently, atop two self-contained
Apache-2.0 libraries: MatrixCache
(caching) and MatrixRaft (replication).
For the shared-storage tier at scale, stateless datanodes read and write one durable object store —
MatrixObject , the enterprise object store engineered for extreme-performance shared storage —
so compute can scale for traffic spikes without copying every shard to every node.
RAG usually runs on a vector DB doing approximate nearest-neighbor (ANN) search over a giant
flat pile of embeddings — trading recall for speed, because scoring every vector exactly is too expensive.
TemporalStore never gets there. Its layered virtual file system and compact
secondary indexes filter most candidates out first — by scope, type, entity, and time — so
by the time we rank, the set is small and we score it exactly, with no ANN loss (that's the
hit@k 0.995–1.00 above). That same scoped tree doubles as a virtual file system for
resources, skills, and memory — one time-aware namespace instead of a separate store for each — so
it's also one fewer system to run and keep in sync.
Temporal memory is append-heavy — you're constantly writing new timestamp-keyed events — which
is exactly what an

[truncated]
