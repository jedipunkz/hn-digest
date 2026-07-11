---
source: "https://millfolio.com/blog/local-ai-infra-tags/"
hn_url: "https://news.ycombinator.com/item?id=48868694"
title: "Managing a small local AI budget (Mac M2 16gb)"
article_title: "Managing a small local AI budget — millfolio"
author: "winding"
captured_at: "2026-07-11T04:55:06Z"
capture_tool: "hn-digest"
hn_id: 48868694
score: 2
comments: 0
posted_at: "2026-07-11T04:17:19Z"
tags:
  - hacker-news
  - translated
---

# Managing a small local AI budget (Mac M2 16gb)

- HN: [48868694](https://news.ycombinator.com/item?id=48868694)
- Source: [millfolio.com](https://millfolio.com/blog/local-ai-infra-tags/)
- Score: 2
- Comments: 0
- Posted: 2026-07-11T04:17:19Z

## Translation

タイトル: 小規模なローカル AI 予算の管理 (Mac M2 16gb)
記事のタイトル: 小規模なローカル AI 予算の管理 — millfolio
説明: millfolio がオンデバイス モデルで何千ものレコードをクエリ可能に保つ方法 - 3 つのフレーバーのタグ システム、優先順位によるバッチ分類、タグが実際に一致するものを示すプレビュー。

記事本文:
小規模なローカル AI 予算の管理 — millfolio millfolio Docs Blog github.com/millfolio ↗ Community ↗ ← ブログ 小規模なローカル AI 予算の管理
2026 年 7 月 10 日 on-device-ai アーキテクチャがパフォーマンスをタグ付け
最初の投稿で、milfolio の分割について説明しました。フロンティア モデルはエイリアスされたスキーマ上に小さなプログラムを書き込み、ローカル モデルはデバイス上の実際のファイルを読み取ります。最後に、ラップトップを使用可能な状態に保つために、優先順位を付けてローカル モデルをバッチ モードで実行するという、地味な部分について書くことを約束しました。これがその投稿です。
問題: モデルが記録に合わせてスケールされない
数年間の銀行やカードの明細書は数千件の取引に相当します。ローカル モデルは、そのうちの 1 つである「これはサブスクリプションですか?」を読み取るのが得意です。まさに、的を絞った質問によく答えます。ラップトップの GPU で、ユーザーの質問ごとにその質問に何千回も答えることはできません。レコードごとにほんの数秒であっても、「昨年、サブスクリプションにいくら費やしましたか?」のような質問は、それは数分間の推論に変わり、次の質問ですべての作業が再び行われることになります。
したがって、設計ルールは次のようになりました。モデルの判断は一度計算して保存するものであり、質問時に実行するものではありません。
答え: タグは 3 種類あります
すべてのトランザクションは、インデックス時に、開いて編集できるプレーンテキストのルール ファイルからタグを取得します (category.txt — 1 行に 1 つのルール、ファイルが信頼できる情報源です)。 3 つのフレーバーはコスト順に並べられています。
文字列タグ — 安価で一般的なケースです。タグは、大文字と小文字を区別しない部分文字列キーワードのリストであり、オプションで一致を拒否する除外を指定できます。
電話 = ベライゾン、at&t、t-mobile、ミントモバイル
食料品 = ホールフーズ、トレーダージョー、セーフウェイ、コストコ
純粋な文字列マッチング、事実上無料で、実際の文字列のほとんどをカバーします。

支出履歴 - 販売者の文字列が反復的です。
参照タグ — 他のタグから構築されたタグ。ルールはキーワードの代わりに @tag 名を参照できるため、他の人のキーワード リストを繰り返さずにグループを構成できます。
食べ物 = @食料品、@レストラン
まだ決定的で自由です。参照タグは、参照されるタグがすでに行上にある場合に起動されます。参照サイクルは無害であり (タグは追加されるだけなので、評価は収束するだけです)、参照は AI タグを指すことが許可され、これにより 3 番目のフレーバーも構成可能になります。
AI タグ — ファジーテールのキーワードリストなしのキャプチャ。ルールは「はい/いいえ」の質問で、オンデバイス モデルが判断します。
サブスクリプション: これは定期的なサブスクリプション料金ですか?
これは、推論が必要な唯一のフレーバーです。まさにこれが、クエリ時の呼び出しではなく保存されたタグである理由です。
質問時には、3 つすべてが同一に見えます。生成されたプログラムは .tags でフィルタリングし、数千のレコードにわたる答えはモデル呼び出しではなく文字列比較です。プライバシーの恩恵もあります。フロンティア モデルにはタグ名とスコープ ノートが伝えられますが (そのため、それらをフィルタリングするプログラムを作成します)、キーワードは決して伝えられません。実際のマーチャント文字列はデバイス上に残ります。
バッチと昼寝のタイミングを知る
AI タグは引き続き 1 回計算する必要があり、「1 回」とは、既存のすべてのレコードと後で到着するすべてのレコードに対するバックフィルを意味します。その作業は、単一のデバイス上の作業オーケストレーターを通じて実行されます。つまり、再起動後も存続するディスクバックアップのキューで、一度に 1 つのジョブを排出します。インデックス作成と AI タグのバックフィルがエンジンをめぐって互いに争うことはなく、両方とも対話型の質問に屈します。
バックフィル ジョブ内で、分類はバッチ処理され、重複が排除されます。
多くの説明が 1 つのプロンプトでモデルに入力され、モデルは 1: はい、2: いいえ、… と応答します。 — 1 つの推論

この呼び出しはスライス全体を分類します。
完全に重複する記述は、モデルが認識する前に折りたたまれます。24 回出現する月額料金は 1 回分類され、判定が広がります。実際のステートメントでは、これにより作業の大部分が節約されます (統計ページには、分類されたレコードと対象となるレコードの実行カウンターが保持されます)。
小さな台帳には、各ルールがどこまで具体化されているかが記録され、単調な挿入世代に基づいてキー設定されます。そのため、完成したタグが古いレコードを再実行することはありません。新しいレコードは段階的に取得されます。
スライス間では、バックフィラーが仮眠します。仮眠の長さが優先設定です。高は 100 ミリ秒の呼吸、中は 1.2 秒、低は 5 秒です。この 1 つのノブによって、マシンが「やり遂げる」か「仕事中だから一晩中」かを決めることができます。バックグラウンド処理を無期限に一時停止することもできます。
次の統計は、Mac M2 16GB で mill run https://raw.githubusercontent.com/millfolio/vault/5915bf52f17b9baf2c4fe71695dae03e05a9ee16/privacy-box/eval/local_classifier_eval.mojo を実行することによって生成されました。 (このプログラムは集計のみを出力し、販売者の文字列や金額は出力しません。また、milfolio のインストールで実行されます。実行中にストリーミングされるライブ進行状況ラインには、milfolio ≥ 0.4.49 が必要です。)
モデル: Qwen/Qwen2.5-3B-Instruct
質問: 「各スニペットについて、食料品店やスーパーマーケットでの購入 (食品の買い物) の場合は「はい」と答え、その他のすべて (外食、レストラン、コーヒーショップ、食料品以外の料金) の場合は「いいえ」と答えてください。すべての答えは正確に「はい」または「いいえ」でなければなりません。決して「なし」であってはならない。キーワード タグ「食料品」を参考にしてください。
データ: 2930 トランザクション行、988 個の個別の説明 (3.0x 重複除去)。分類済み 400 (バランス: タグ付き 29 + タグなし 371)、2228 行をカバー
ML (行レベル、モデルとキーワードのルール):
同意 70.6% p

精度 39.0% リコール 100.0% F1 56.2%
混乱: TP 419 FP 654 FN 0 TN 1155 基本率 18.8%
(キーワード タグは不完全な真実です。不一致はモデルとルールの違いであり、認定されたモデルのエラーではありません。サンプルはバランスがとれているため、基本レート + 一致はボールトではなくサンプルを反映します)
回答の組み合わせ: 111 はい · 25 いいえ · 264 なし · 0 その他
運用:
約 40 のモデル呼び出しで 400 の記述 (呼び出しあたり約 10 のバッチ処理)、合計 260.9 秒
1.5 個別/秒 · 重複排除ファンアウト後に有効な 8.5 行/秒 · 個別の記述あたり最大 652 ミリ秒
(回答混合行に関する注記: none はバッチ プロトコルの適用されないエスケープ ハッチです。モデルはこれをほとんどの行の否定として使用します。きれいに解析され、「いいえ」としてカウントされます。)
食料品キーワード タグ (スーパー マーケット ブランドの厳選されたリスト) が参照ラベルであり、Qwen2.5-3B-Instruct は同じレコードに対して同等の「はい/いいえ」の質問を取得します。サンプルは意図的にバランスがとられており (すべてのタグ付き説明にタグなしの説明が埋め込まれています)、分類はモデル呼び出しごとに 10 個の説明にまとめられ、定期的な料金は実際のタグ バックフィルとまったく同じ方法で重複排除されます。つまり、2,930 のトランザクション行が 988 個の個別の説明に集約され、モデルが何かを認識する前に 3 倍の節約になります。
その結果、2,228 行をカバーする 400 を超える個別の説明が得られました。キーワード ルールに対する再現率は 100%、精度は 39% (F1 56%) でした。リコール番号は、モデルがリストを見ずに、キーワード リスト タグの 419 行すべてを食料品として捕らえたことを示しています。精度の数値は見た目よりも興味深いものです。654 件の「誤検知」は、モデルが食料品と述べているのに、私の 10 ブランドのキーワード リストでは何も述べていない行です。本物のモデルのエラーと食料品のリストが単に知らないプールが混在しています。この非対称性がハイブリッド設計のすべての議論です。

終端的なルールにより、正確で監査可能な頭が得られます。モデルは尾部をカバーします。プレビュー UI が存在するのは、どちらも盲目的に信頼すべきではないためです。
運用上: 個別の説明ごとに最大 650 ミリ秒、個別の説明あたり 1.5 行/秒、および重複排除によって各判定が取り消されると効果的な 8.5 行/秒。つまり、2,900 行のコンテナーは数分のバックグラウンド時間で完全な AI パスを取得します。これは、まさにバックフィル スケジューラが設計されている予算と同じです。
タグ システムは、ルールに対する信頼度に応じて決まります。そのため、UI によって 1 ブラインドをセーブされることはありません。タグを作成または編集すると、何かが永続化される前に予行演習が行われます。
文字列または参照タグの場合、保存されたトランザクションに対して編集されたルールが評価され、一致数と一致した説明の例、および一致しなかった例が表示されるため、欲張りすぎるキーワードはすぐに表示されます。
AI タグの場合、質問はタイムボックス化されたサンプル (オンデバイス モデルの約 5 秒) に対して実行され、肯定的な例と否定的な例を並べて「≈N 個のレコードが一致します」という結果が得られます。
両方の側面を見ることが重要です。誤検知は、一致したリストから明らかです。偽陰性は、UI がルールの外にあるものを表示する場合にのみ表示されます。保存するまでは何も書き込まれません。 codegen モデルは、分類の管理においてユーザーと協力します。生成されたプログラムがインラインで何かを分類する必要がある場合、再利用可能な AI タグとして質問を提案し返すことができるため、次回からは新たな推論ではなく、保存されたタグ フィルターが使用されます。
これはどれも珍しいものではありません。これは、すべてのデータ システムが最終的に成長するバッチとオンラインの分割と同じです。ローカルな工夫としては、高価なリソースはクラウドの請求書ではなく、GPU とフォアグラウンド作業であるため、スケジューラの仕事はスループットと同じくらい丁寧であるということです。ジャッジマンを計算する

それらをタグとして保存し、ファジーテールにのみ AI 予算を費やし、ルールが実行される前にルールが何を実行するかを常にユーザーに示します。

## Original Extract

How millfolio keeps thousands of records queryable with an on-device model — a three-flavor tag system, batched classification with priorities, and previews that show you what a tag would actually match.

Managing a small local AI budget — millfolio millfolio Docs Blog github.com/millfolio ↗ Community ↗ ← Blog Managing a small local AI budget
July 10, 2026 on-device-ai architecture tags performance
In the first post I described millfolio’s split: a frontier model writes a small program over an aliased schema, and a local model reads your actual files on-device. I ended with a promise to write about the unglamorous part — running the local model in batch mode, with priorities, so your laptop stays usable. This is that post.
The problem: the model doesn’t scale with your records
A couple of years of bank and card statements is thousands of transactions. The local model is good at reading one of them — “is this a subscription?” is exactly the kind of targeted question it answers well. What it cannot do is answer that question thousands of times, per user question, on a laptop GPU. Even at a fraction of a second per record, a question like “what did I spend on subscriptions last year?” would turn into minutes of inference — and the next question would do all that work again.
So the design rule became: the model’s judgment is a thing you compute once and store, not a thing you run at question time.
The answer: tags, in three flavors
Every transaction gets tags at index time, from a plain-text rule file you can open and edit ( categories.txt — one rule per line, the file is the source of truth). The three flavors are ordered by cost:
String tags — the cheap, common case. A tag is a list of case-insensitive substring keywords, with optional excludes that veto a match:
phone = verizon, at&t, t-mobile, mint mobile
groceries = whole foods, trader joe, safeway, costco
Pure string matching, effectively free, and it covers most of a real spending history — merchant strings are repetitive.
Reference tags — tags built out of other tags. A rule can reference @tag names instead of keywords, so you can compose groups without repeating anybody’s keyword list:
food = @groceries, @restaurant
Still deterministic and free — a reference tag fires when a referenced tag is already on the row. Reference cycles are harmless (tags are only ever added, so evaluation just converges), and a reference is allowed to point at an AI tag, which makes the third flavor composable too.
AI tags — the fuzzy tail no keyword list captures. The rule is a yes/no question, and the on-device model is the judge:
subscriptions : is this a recurring subscription charge?
This is the only flavor that costs inference — which is exactly why it’s a stored tag and not a query-time call.
At question time, all three look identical: the generated program filters on .tags , and the answer over thousands of records is a string comparison, not a model call. There’s also a privacy dividend: the frontier model is told the tag names and scope notes (so it writes programs that filter on them), but never the keywords — your actual merchant strings stay on the device.
Batches, and knowing when to nap
AI tags still have to be computed once, and “once” means a backfill over every existing record plus whatever arrives later. That work runs through a single on-device work orchestrator: a disk-backed queue that survives restarts, draining one job at a time — indexing and AI-tag backfill can never fight each other for the engine, and both yield to an interactive question.
Inside a backfill job, the classification is batched and deduplicated:
Many descriptions go to the model in one prompt , and it answers 1: yes, 2: no, … — one inference call classifies a whole slice.
Exact-duplicate descriptions are collapsed before the model sees them : a monthly charge that appears 24 times is classified once and the verdict fans out. On real statements this saves a large fraction of the work (the Stats page keeps a running counter of records classified vs. records covered).
A small ledger records how far each rule has been materialized, keyed on a monotonic insertion generation — so a finished tag never re-runs over old records; new records get picked up incrementally.
Between slices, the backfiller naps — and the nap length is the priority setting: high is a 100ms breath, medium is 1.2s, low is 5s. That one knob is how you decide whether the machine is “get it done” or “I’m working, take all night.” You can also pause the background processing indefinitely.
The following stats were generated by running mill run https://raw.githubusercontent.com/millfolio/vault/5915bf52f17b9baf2c4fe71695dae03e05a9ee16/privacy-box/eval/local_classifier_eval.mojo on a Mac M2 16GB. (The program prints only aggregates — no merchant strings or amounts — and runs on any millfolio install; the live progress lines it streams while running need millfolio ≥ 0.4.49.)
Model: Qwen/Qwen2.5-3B-Instruct
Question: "For each snippet: answer 'yes' if it is a grocery store or supermarket purchase (food shopping), answer 'no' for everything else (dining out, restaurants, coffee shops, and any non-grocery charge). Every answer must be exactly 'yes' or 'no' — never 'none', never anything else." vs the keyword tag 'groceries' as reference.
Data: 2930 transaction rows, 988 distinct descriptions (3.0x dedup); classified 400 (balanced: 29 tagged + 371 untagged) covering 2228 rows
ML (row-level, model vs keyword rule):
agreement 70.6% precision 39.0% recall 100.0% F1 56.2%
confusion: TP 419 FP 654 FN 0 TN 1155 base rate 18.8%
(keyword tags are imperfect truth — a disagreement is model-vs-rule, not a certified model error; the sample is BALANCED, so base rate + agreement reflect the sample, not the vault)
answer mix: 111 yes · 25 no · 264 none · 0 other
Operational:
400 descriptions in ~40 model calls (batched ~10/call), 260.9s total
1.5 distinct/s · 8.5 rows/s effective after dedup fan-out · ~652ms per distinct description
(A note on the answer mix line: none is the batch protocol’s does-not-apply escape hatch — the model uses it as its negative for most rows. It parses cleanly and counts as a “no.”)
The groceries keyword tag — a curated list of supermarket brands — is the reference label, and Qwen2.5-3B-Instruct gets the equivalent yes/no question over the same records. The sample is deliberately balanced (every tagged description, padded with untagged ones), classification is batched ten descriptions per model call, and recurring charges are deduplicated exactly the way the real tag backfill does it: 2,930 transaction rows collapse to 988 distinct descriptions, a 3× saving before the model sees anything.
The result, over 400 distinct descriptions covering 2,228 rows: 100% recall and 39% precision against the keyword rule (F1 56%). The recall number shows that the model caught every one of the 419 rows my keyword list tags as groceries, without ever seeing the list. The precision number is more interesting than it looks: the 654 “false positives” are rows where the model says groceries and my dozen-brand keyword list says nothing — a pool that mixes genuine model errors with grocers my list simply doesn’t know. That asymmetry is the whole argument for the hybrid design: deterministic rules give you a precise, auditable head; the model gives you coverage of the tail; and the preview UI exists because neither should be trusted blind.
Operationally: ~650 ms per distinct description, 1.5 distinct/s, and 8.5 rows/s effective once deduplication fans each verdict back out — so a 2,900-row vault gets a full AI pass in minutes of background time, which is exactly the budget the backfill scheduler is designed around.
A tag system is only as good as your confidence in the rules, so the UI never makes you save one blind. When you create or edit a tag, you get a dry-run before anything persists:
For a string or reference tag , the edited rules are evaluated over your stored transactions and you see the match count plus example descriptions that matched — and examples that didn’t — so a too-greedy keyword shows itself immediately.
For an AI tag , the question runs against a time-boxed sample (~5 seconds of the on-device model) and you get “≈N records would match” with positive and negative examples side by side.
Seeing both sides matters. A false positive is obvious from the matched list; a false negative only shows up when the UI also shows you what fell outside the rule. Nothing is written until you save. The codegen model partners with you in managing the taxonomy — when a generated program had to classify something inline, it can propose the question back as a reusable AI tag, so the next time it’s a stored-tag filter instead of fresh inference.
None of this is exotic — it’s the same batch-vs-online split every data system eventually grows. The local twist is that the expensive resource isn’t a cloud bill, it’s your GPU and your foreground work, so the scheduler’s job is as much politeness as throughput. Compute judgments once, store them as tags, spend the AI budget only on the fuzzy tail, and always show the user what a rule would do before it does it.
