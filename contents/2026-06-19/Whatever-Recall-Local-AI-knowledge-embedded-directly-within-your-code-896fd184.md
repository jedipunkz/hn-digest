---
source: "https://whatever-recall.com"
hn_url: "https://news.ycombinator.com/item?id=48597220"
title: "Whatever Recall – Local AI-knowledge embedded directly within your code"
article_title: "whatever-recall — The code is its own memory."
author: "ChristianMcCain"
captured_at: "2026-06-19T13:19:20Z"
capture_tool: "hn-digest"
hn_id: 48597220
score: 2
comments: 0
posted_at: "2026-06-19T11:05:44Z"
tags:
  - hacker-news
  - translated
---

# Whatever Recall – Local AI-knowledge embedded directly within your code

- HN: [48597220](https://news.ycombinator.com/item?id=48597220)
- Source: [whatever-recall.com](https://whatever-recall.com)
- Score: 2
- Comments: 0
- Posted: 2026-06-19T11:05:44Z

## Translation

タイトル: Whatever Recall – コード内に直接埋め込まれたローカル AI 知識
記事のタイトル: anything-recall — コードはそれ自体の記憶です。
Description: Code that knows itself. AI はトレーニングからシンボル名を発明し、リポジトリを grep し、ミスし、トークンを焼き付けます。 remember は、リポジトリが実際に何を呼んでいるのかを書き込み時にキャプチャするため、検索前に推測を修正し、決定とその理由をミリ秒単位でトークンゼロで読み戻します。

記事本文:
anything-recall — コードはそれ自体のメモリです。何でもリコール 仕組みの詳細 ドキュメントのインストール 価格 v1.0 のダウンロード 何でもログイン - リコール 0 1 仕組み → 0 2 詳細 → 0 3 インストール → 0 4 ドキュメント → 0 5 価格 → v1.0 のダウンロード 14 日間無料で開始 → 自己ホスト型 · カードは不要 · リコールするトークンはゼロ
GitHub でソースを表示 RC-01 — AI ネイティブ プロジェクト メモリ · ビルトイン バイエルン コード
それ自体を知っています。
コードを変更する必要はありません · 自己ホスト型 · カードは必要ありません · トークンを呼び出す必要はありません
AI がコードを実行した瞬間にコードは忘れ、その後推測します。トレーニングからシンボル名を発明し、このリポジトリを grep して、その名前が付けられていないものを探し、失敗し、トークンを焼き付けて再試行します。代わりに、recall はコードを自己認識させます。書き込み時に決定と理由が各コミットにスタンプされ、recall はリポジトリの実際の語彙をすでに知っているため、grep の前に推測された名前を修正します。埋め込みは事後的にそれを学習しません。読み返すことは即座に行え、追跡可能であり、無料です。
AI に人間のように考えさせるのはやめてください。
人間の開発者は頭の中に保持できることに制限されているため、私たちは人間の限界に対応するツールを構築しました。 AIはそうではありません。グラフ全体を保持し、0.25 ミリ秒でクレームを再検証でき、決して疲れることはありません。それに欠けているのは、すべての行の背後にある理由、元に戻してはいけない決定、物事の本当の名前などを記憶するプロジェクトです。
リコールはまさにそれを提供します。つまり、記憶、良心、実績のあるコードベースです。 AI を人間の制限から解放し、プロジェクトだけが保持できるコンテキストを AI に与え、創造的、正確、高速に実際に可能な方法でコーディングします。 AI は手です。プロジェクトは同僚です。
「…しかし、それは私のコードに何をもたらすのでしょうか？」
何もない。呼び出しでは 1 行も変更されません。
としか読み取れません。すべてが覚えています

ows はローカルの .mind/ インデックスとコミット メッセージ内に存在しますが、ソース内には存在しません。再フォーマット、リファクタリング、編集、ファイルの移動は必要ありません。 .mind/ をいつでも削除すると、リポジトリはバイトごとに元の状態になります。
知識はコードの傍ら、Wiki、チケット、チャット ウィンドウに存在します。コード自体はフラットで沈黙を保つため、変更された瞬間にドキュメントは嘘をつき始めます。
リコールはコード内の知識を保持します。つまり、それを変更したコミットに固定されます。 One source of truth, not two that drift apart.
コードには 1 つの真実があります。したがって、以下のすべてはコードに基づいています。
AIがまだ認識しているときにキャプチャされます。
高価な思考はバグが修正されている間に発生しますが、その後蒸発します。その後、誰か (または AI) がコールド構文から理由を再導出する必要があります。
AI が動作し、コンテキストがまだ頭の中にある間、リコールは決定 + 理由 + アンカーをコミットにスタンプします。安くて正確、一文。
$ git commit -m "シートチェックをアトミック RPC に切り替えます
理由を思い出してください: JS のチェックと挿入が負荷の下で競合していました。
最後の席を二重予約した。 RPC は組織行をロックします。
Recall-Anchors: Seats、confirmSeatOrRollback、orgs.ts」 理由はわかった瞬間に書かれているため、読み返すにはモデルは必要ありません。
あらゆる「なぜここにあるの？」 AI にファイル全体を grep させて再読み取りさせます。質問ごとに数万のトークンがセッションごとに再度支払われます。
Because the meaning was written at write-time, reading is a plain SQLite + FTS5 lookup — no model, no embeddings.ミリ秒未満で完成した回答。
シンボルのコードベースを grep します
3 つの候補ファイルを上から下へ開いて読み取ります
毎ターン、推論を最初から導き出し直す
スタンプされたアンカーに対する 1 つの SQLite ルックアップ
決定 + 理由 + リンクを返します
コミットに固定 — 追跡可能、鮮度チェック済み
向きを変えるのは無料です。

いつでも全体像を読む余裕があります。
最初のキー入力の前に 6 つの回答。
AI はファイルをブラインドで開きます。AI は、ファイルに関連付けられている開いているタスク、元に戻そうとしている決定、または 3 つのファイルを分離する変更内容を見ることができません。
1 つのリコール概要では、6 つの質問 (何を、どこで、なぜ、何が壊れるのか、開いているタスク、証跡) に同時に答えます。読み取り専用、トークン 0、ファイルの完全なプロファイルです。
ただ読むだけではなく、見てください — グラフの再現はライブで保持されます
▶ 探索 — コードをドラッグ、ズームインします グラフネイティブの地平線を垣間見ることができます — プロトタイプ、まだ出荷されていません。クリックしてライブで探索してください。全体像は手頃な価格で提供されるため、AI は見たこともない決定を黙って取り消すことはありません。
コードはそれが何と呼ばれるかを知っています。 AI はそうではありません。
AI はトレーニングからシンボル名を発明します — enforceSeats ?座席制限? — このリポジトリを grep しますが、失敗し (ブラインド grep では 0 of 12 が表示されます)、推測ループを再び焼き付けます。
思い出しても検索が改善されるのではなく、結果が好転するのです。このリポジトリの実際の語彙は書き込み時にすでに書き込まれているため、grep が実行される前に推測を実際の名前に修正します。
もっともらしい名前を発明します — getUserSession()
このリポジトリを grep します — 0 件のヒット
再度推測し、ファイルを再読み込みし、トークンを書き込みます
リポジトリ独自の語彙に基づいて推測をランク付けします
それを修正します — getUserSession() →solveActiveSeat()
grep の着地が最初に試行され、トークンは無駄になりません
このアイデアはサポート現場から直接得られたものです。上級エージェントはすぐに答えるのではなく、誤解されないように質問を確実に行うようにします。 remember は AI に対してそれを行います。単一のトークンが推測に費やされる前に、リポジトリの実際の言葉に照らして質問を構造化します。
私たちは答えではなく質問を解決します。コードは AI の上級サポート エージェントになります。
作業すればするほど記憶は真実になる

。
Docs rot because nobody notices when reality diverges from them.半年が経ち、Wiki のほとんどは間違っており、誰も信用しなくなりました。
すべてのメモは SHA に固定されます。その下のコードが移動すると、リコールでドリフト (🟡) にフラグが立てられ、次のコミット時に修正が提供されます。承認すると、真実は最新のままになります。フラグを立てるだけで、黙って書き換えることはありません。
メモとコードを照合するものは何もありません
AI は腐ったバージョンを読み取ります
そして自信を持って間違ったものを構築する
SHA-pinned, so drift is detectable
治癒または置き換えの提案を思い出す
あなたは承認します - 古い決定は歴史に残ります
ドリフトは捕らえられ、修復されます。そして、それが書き込み時間にフィードバックされるため、脳は複雑になります。 ↻ ループが終了します — すべてのコミットが次のコミットをよりスマートにします。
トークン請求書は、
物語全体。
AI にとって、コンテキスト空間は希少なリソースです。 grep は単語がどこにあるかを見つけます。その理由を知るには、質問ごとに数万のトークンを使用して、ファイル全体をコンテキストに読み戻す必要があります。リコールは ~56 で決定を返します。 2 つのライブ本番リポジトリ (240 コミットのアプリ、668 コミットの CMS) でコールド スタートを測定し、リコール トレーラーは植えられていません。そのギャップが製品です。
チェーン全体、
請求書で返金されます。
6 つのリンク、1 つの結果: AI はついに最高の仕事をするために必要なものを手に入れ、企業は AI の推測にお金を払うのをやめました。貸借対照表に計上されるものは次のとおりです。
同じ 3 つの回答について、測定したところ、152 対 ~214,000 トークンでした。チーム全体、年間を通じて、その読み取り時間の節約は数百万単位に拡大します。これまではモデル コールに費用がかかっていたすべてのオリエンテーションが無料になりました。
Code that gets better on its own
ループはさらに複雑になります。コミットごとに新しい知識が刻印され、古いドリフトが修復されるため、最もよく作業されたコードには最も正確な記憶が残ります。長く生きれば生きるほど、受動的に賢くなります。
より迅速な結果、より少ないミス
AIは

s briefed before it edits and warned before it repeats a known mistake — so it stops re-deriving, stops hallucinating constraints, and stops undoing decisions it never saw.
逆検索により、grep の前に推測された名前がリポジトリの実際の語彙に修正されます。トークンや壁時計を書き込む推測→ミス→推測のループがなくなりました。
爆発範囲、競合するホットスポット、因果関係の痕跡、そして新しい開発者または新しいエージェントは、1 つのコマンドでリポジトリ全体の推論を継承します。
コードとメモリがマシンから離れることはありません。クラウド、テレメトリ、ベクトル DB、API キーはありません。 1 つの小さなツールである pure stdlib は、ネットワーク ケーブルを接続した状態で実行されます。
あなたのコードは決して
マシンを離れます。
remember は小規模な自己ホスト型ツールであり、プラットフォームではありません。必要なのはプロジェクト フォルダーへのアクセスです。これが全体のフットプリントで、セットアップ全体には 2 分かかります。
リポジトリ (.mind/) 内の小さな CLI とローカル インデックス。実行するサーバーやバックグラウンドでのエージェントは必要ありません。 1 時間ごとのライセンス チェック (短い猶予期間を持つ簡単なオンライン確認) の間にオフラインで動作します。コードがマシンから離れることはありません。
何もアップロードされず、何も同期されず、テレメトリも行われません。 Web サイトはプロジェクトを読み取るだけで、ライセンスをチェックするだけです。
完全なソースは GitHub で公開されています。リポジトリを実行する前に、リポジトリに関わるすべての行を読み取ります。
$ pip install git+ https://github.com/heidrich/whatever-recall .git
$初期化を思い出してください。
# インデックス付き — あなたのリポジトリはそれ自体のメモリになります。構成、クラウド アカウント、アップロードは必要ありません。メモリはリポジトリ内に存在し、すべてのクローンに同梱されます。質問がありますか? GitHub 上のソースは 1 クリックでアクセスできます。
あなたの決断は決して
マシンを離れてください。
コードベースに関して最もデリケートなことは、その背後にある理由です。リコールにより、コードはソースではなくローカルの脳内に保持されるため、コードはクリーンな状態で出荷されます。

そして、あなたがそれを共有することを選択しない限り、あなたの推論はあなたのもののままです。
なぜ存在するのかはコードではなく脳にある
あなたがスタンプしたすべての決定は、ローカルの .mind/ インデックスに保存されます。コードの中にではなく、コードの横にあります。ソースはバイトごとにクリーンな状態で出荷されます。公開する前にスクラブするものは何もありません。
ノートごとの可視性: チームまたはプライベート
メモに非公開のマークを付けると、それはあなただけのものになります。決してエクスポートしたり、共有したりすることはありません。メモのプライベート性は高まるだけであり、プライベート性が下がることはありません。偶然に何かを広げることはできません。
漏れることなく脳を共有する
呼び出しエクスポートは、プライベートなメモをすべて取り除いた共有可能なコピーを書き込みます。チームがそれを共有するかどうかはあなたの判断です。その内容はリコールの保証です。チームの知識であり、決してプライベートなものではありません。
エクスポートは、まだプライベートなメモを保持しているブレインを書き込むのではなく中止され、チェックリークによってステージングされるコミットがブロックされます。プライベートな決定は物理的にチームメイトのクローンやパブリック リポジトリに到達することはできません。
$ リコールスタンプ「これを選んだ理由」 --private
🔒 プライベート — この脳内に残ります
$ 呼び出しエクスポート --out .mind/shared.db
# プライベートなメモは削除され、ゲートはクリーンであることが確認されました。 個人開発者は、構築によってクリーンなコードベースを出荷します。チームは必要な知識を共有しますが、共有しないものは何もありません。あなたのコードは公開成果物です。あなたの推論はあなたのものです。プライベートな知識をローカルに保つ方法 →
1つの製品。
席ごとに料金が設定されています。
階層も機能のロックも、「アップグレードによるロック解除」もありません。 CLI、ダッシュボード、パワー モード、Web-AI ブリッジ、MCP、無制限のリポジトリなど、誰もがすべての機能を利用できます。メモリを共有する人数を選択するだけです。 1 つのシートを購入してソロで実行するか、シートを追加するとチーム ツールが表示されます。 14 日間すべてを無料でお試しください。カードは必要ありません。
1 席 = 単独走行。 2 つ目を追加して、チーム ツール (メンバー、キー、請求) のロックを解除します。
カードはありません。あらゆる機能。いつでもキャンセルできます。
CLI、

ダッシュボード、MCP、git フック
パワーモードとWeb-AIブリッジ
トークンフリーのリコール — コストゼロのモデルトークンを読み取ります
1時間ごとの座席チェックの間はオフラインで動作します（小切手が届かない場合は少し猶予があります）
無制限のリポジトリ、無制限のスタンプ
自己ホスト型 — コードとメモリがマシンから離れることはありません
2 シート以上のチーム ツール: メンバー、ライセンス キー、共有請求
すべての新しいリリースは 3 年後にオープンソース (Apache 2.0) になります
開発者が 1 人であっても 100 人であっても同じ製品です。スケールするのはメモリを共有する人の数だけです。
私は、教育、そして一緒に問題を探求し議論することが、人類を素晴らしい未来に導く最善の方法であると強く信じています。
そのため、学校や大学は教育や学術研究のためにリコールを無償で利用します。これまでに手数料はかかりません。あなたが何を構築しているのかぜひお聞きください。唯一の正直な意見は、研究が収益をもたらす製品やスピンオフになった場合、その時点から定期的な計画が必要であるということです。
14 日間の試用版、フル機能、クレジット カードは不要です。どちらの方法でも自己ホスト型 - コードとメモリがマシンから離れることはありません。アカウントにはライセンスのみが含まれます。ビジネス ソース ライセンス 1.1 に基づいてライセンスされています。リリースされたすべてのバージョンは、リリースから 3 年後にオープン ソース (Apache 2.0) になるため、スタックがプロプライエタリな行き詰まりに陥ることはありません。
「私たちは常に文書を作成し、その後計画を変更しますが、決して計画を立てません」

[切り捨てられた]

## Original Extract

Code that knows itself. Your AI invents a symbol name from training, greps your repo, misses, and burns tokens. recall captured what the repo really calls things at write-time, so it fixes the guess before the search — and reads back the decision and the why in milliseconds, at zero tokens.

whatever-recall — The code is its own memory. whatever- recall How it works Deep dive Install Docs Pricing Download v1.0 Login whatever- recall 0 1 How it works → 0 2 Deep dive → 0 3 Install → 0 4 Docs → 0 5 Pricing → Download v1.0 Start free for 14 days → Self-hosted · no card needed · zero tokens to recall
View the source on GitHub RC-01 — AI-native project memory · Built in Bavaria Code that
knows itself .
Never changes your code · self-hosted · no card needed · zero tokens to recall
Your code forgets the moment the AI does — and then it guesses. It invents a symbol name from training, greps this repo for something that was never called that, misses, and burns tokens trying again. recall makes the code self-aware instead: the decision and the why are stamped into each commit at write-time, and recall already knows the repo's real vocabulary, so it fixes the guessed name before the grep. No embedding learns that after the fact. Reading it back is instant, traceable, and free.
Stop making your AI think like a human .
A human dev is bound by what they can hold in their head — so we built tools for human limits. An AI isn't. It can hold the whole graph, re-verify a claim in 0.25 ms, and never tire. What it lacks is a project that remembers : the why behind every line, the decisions it must not undo, the real names of things.
recall gives it exactly that — a codebase with memory, a conscience, and a track record. Free the AI from human restrictions, give it the context only the project can hold, and it codes the way it actually can: creative, exact, and fast. The AI are the hands; the project is the coworker.
“…but what does it do to my code?”
Nothing. recall never changes a single line .
It only reads . Everything recall knows lives in a local .mind/ index and in your commit messages — never in your source. No reformatting, no refactor, no edits, no files moved. Delete .mind/ any time and your repo is byte-for-byte what it was.
Knowledge lives beside the code — in a wiki, a ticket, a chat window. The code itself stays flat and silent, so the instant it changes, the doc starts to lie.
recall keeps the knowledge in the code : pinned to the commit that changed it. One source of truth, not two that drift apart.
One truth in the code — so everything below follows from it.
Captured when the AI still knows .
The expensive thinking happens while the bug is being fixed — then it evaporates. Later, someone (or some AI) has to re-derive the why from cold syntax.
While the AI works and the context is still in its head, recall stamps the decision + the why + the anchors onto the commit. Cheap, exact, one sentence.
$ git commit -m "switch seat check to an atomic RPC
Recall-Why: the JS check-then-insert raced under load and
double-booked the last seat; the RPC locks the org row.
Recall-Anchors: seats, confirmSeatOrRollback, orgs.ts" The why is written at the knowing moment — so reading it back needs no model.
Every “why is this here?” makes the AI grep and re-read whole files — tens of thousands of tokens per question, paid again every session.
Because the meaning was written at write-time, reading is a plain SQLite + FTS5 lookup — no model, no embeddings. The finished answer in sub-milliseconds.
grep the codebase for the symbols
open + read 3 candidate files top to bottom
re-derive the reasoning from scratch — every turn
one SQLite lookup over the stamped anchors
returns the decision + the why + the links
pinned to a commit — traceable, freshness-checked
Orienting is free — so you can afford to read the whole picture, every time.
Six answers, before the first keystroke.
An AI opens a file blind — it can't see the open task wired to it, the decision it's about to undo, or what its change will break three files away.
One recall brief answers six questions at once — what · where · why · what breaks · open tasks · the trail — read-only, 0 tokens, the full profile of a file.
See it, don't just read it — the graph recall holds, live
▶ explore — drag, zoom into the code A glimpse of the graph-native horizon — a prototype, not shipped yet. Click to explore it live. The whole picture is affordable — so an AI never silently undoes a decision it never saw.
The code knows what it's called. Your AI doesn't.
An AI invents a symbol name from its training — enforceSeats ? seatLimit ? — greps this repo, misses (blind grep lands 0 of 12 ), and burns the loop guessing again.
recall doesn't make the search better — it turns it around. It already wrote down this repo's real vocabulary at write-time, so it corrects the guess into the real name before the grep runs.
invents a plausible name — getUserSession()
greps this repo for it — 0 hits
guesses again, re-reads files, burns tokens
ranks the guess by the repo's own vocabulary
corrects it — getUserSession() → resolveActiveSeat()
the grep lands first try, no tokens wasted
The idea is straight from the support trenches: a senior agent doesn't answer faster, they make sure the question is asked so it can't be misunderstood. recall does that for your AI — it structures the question against the repo's real words before a single token is spent guessing.
We fix the question, not the answer — the code becomes the senior support agent for the AI.
The memory gets truer the more you work.
Docs rot because nobody notices when reality diverges from them. Six months in, most of the wiki is wrong and nobody trusts any of it.
Every note is pinned to a SHA. When the code beneath it moves, recall flags the drift (🟡) and offers the fix at the next commit — you approve, the truth stays current. It only flags, never silently rewrites.
nothing checks the note against the code
the AI reads the rotten version
and confidently builds the wrong thing
SHA-pinned, so drift is detectable
recall offers to heal or supersede
you approve — the old decision stays in history
Drift is caught and healed — and that feeds back into write-time, so the brain compounds. ↻ the loop closes — every commit makes the next one smarter.
The token bill is
the whole story.
For an AI, context space is the scarce resource. grep finds where a word sits — to learn why , it must read whole files back into context, tens of thousands of tokens per question. recall returns the decision in ~56. Measured cold-start on two live production repos (a 240-commit app, a 668-commit CMS), no recall trailers planted — that gap is the product.
The whole chain,
paid back in your bill .
Six links, one outcome: the AI finally has what it needs to do its best work, and the company stops paying for it to guess. Here's what lands on the balance sheet.
152 vs ~214,000 tokens for the same three answers, measured. Across a team, across a year, that read-time saving scales into the millions — every orientation that used to cost a model call is now free.
Code that gets better on its own
The loop compounds: every commit stamps new knowledge and heals old drift, so the most-worked code carries the truest memory. The longer it lives, the smarter it gets — passively.
Faster results, fewer mistakes
The AI is briefed before it edits and warned before it repeats a known mistake — so it stops re-deriving, stops hallucinating constraints, and stops undoing decisions it never saw.
Search-inversion fixes the guessed name to the repo's real vocabulary before the grep — no more guess → miss → guess loop burning tokens and wall-clock.
Blast radius, contested hotspots, the causal trail — and a new dev or a fresh agent inherits the whole repo's reasoning in one command.
Your code and memory never leave your machine — no cloud, no telemetry, no vector DB, no API key. One small tool, pure stdlib, runs with the network cable pulled.
Your code never
leaves your machine.
recall is a small, self-hosted tool — not a platform. It needs exactly one thing: access to your project folder. That's the whole footprint, and the whole setup takes two minutes.
A small CLI and a local index inside your repo (.mind/). No server to run, no agent in the background. Works offline between hourly licence checks (a brief online confirmation, with a short grace window) — your code never leaves your machine.
Nothing is uploaded, nothing is synced, no telemetry. It only reads your project — the website just checks your license.
The full source is public on GitHub. Read every line that touches your repo before you run it.
$ pip install git+ https://github.com/heidrich/whatever-recall .git
$ recall init .
# indexed — your repo is its own memory now That's it. No config, no cloud account, no upload — the memory lives in your repo and ships with every clone. Questions? The source on GitHub is one click away.
Your decisions never
leave your machine.
The most sensitive thing about a codebase is the why behind it. recall keeps it in a local brain, not in your source — so your code ships clean, and your reasoning stays yours unless you choose to share it.
The why lives in the brain, not your code
Every decision you stamp sits in the local .mind/ index — beside your code, never inside it. Your source ships byte-for-byte clean; there's nothing to scrub before you publish.
Per-note visibility: team or private
Mark a note --private and it's yours alone — never in an export, never in a shared brain. Notes can only get more private, never less. You can't widen something by accident.
Share a brain — without leaking
recall export writes a shareable copy with every private note stripped out. Whether your team shares it is your call; what's in it is recall's guarantee: team knowledge, never private.
The export aborts rather than write a brain that still holds a private note, and check-leak blocks a commit that would stage one. A private decision physically can't reach a teammate's clone or a public repo.
$ recall stamp "why we chose this" --private
🔒 private — stays in this brain
$ recall export --out .mind/shared.db
# private notes stripped · gate verified clean A solo developer ships a clean codebase by construction. A team shares the knowledge it wants to — and nothing it doesn't. Your code is the public artifact; your reasoning is yours. How private knowledge stays local →
One product.
Priced by seats .
No tiers, no feature locks, no “upgrade to unlock.” Everyone gets every feature — the CLI, the dashboard, Power Mode, the web-AI bridge, MCP, unlimited repos. You only choose how many people share the memory. Buy a single seat and run solo, or add seats and the team tools appear. Try all of it free for 14 days — no card.
1 seat = run solo. Add a 2nd to unlock team tools (members, keys, billing).
No card. Every feature. Cancel anytime.
The CLI, the dashboard, MCP & git hooks
Power Mode & the web-AI bridge
Token-free recall — reads cost zero model tokens
Works offline between hourly seat checks (a short grace if a check can't reach us)
Unlimited repositories, unlimited stamps
Self-hosted — your code & memory never leave your machine
Team tools at 2+ seats: members, license keys, shared billing
Every new release becomes open source (Apache 2.0) after 3 years
The same product whether you're one developer or a hundred — the only thing that scales is the number of people sharing the memory.
I firmly believe that education — and exploring and debating problems together — is the best way to lead humanity into a fantastic future.
So schools and universities use recall for teaching and academic research at no charge — no fees, ever . We'd love to hear what you're building. The one honest line: if the research becomes a product or spin-off that earns money, it needs a regular plan from that point on.
14-day trial, full features, no credit card. Self-hosted either way — your code and your memory never leave your machine; the account only carries your license. Licensed under the Business Source License 1.1 — every released version becomes open source (Apache 2.0) three years after its release , so your stack can never end up in a proprietary dead end.
“We document constantly, then change the plan and never up

[truncated]
