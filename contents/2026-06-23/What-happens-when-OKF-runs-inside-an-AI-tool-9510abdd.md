---
source: "https://tenureai.dev/writing/open-knowledge-format-ai-memory-benchmark/"
hn_url: "https://news.ycombinator.com/item?id=48642106"
title: "What happens when OKF runs inside an AI tool"
article_title: "Open Knowledge Format Benchmark: What Happens When OKF Runs Inside an AI Tool | Tenure"
author: "jflynt76"
captured_at: "2026-06-23T09:03:54Z"
capture_tool: "hn-digest"
hn_id: 48642106
score: 4
comments: 0
posted_at: "2026-06-23T08:42:10Z"
tags:
  - hacker-news
  - translated
---

# What happens when OKF runs inside an AI tool

- HN: [48642106](https://news.ycombinator.com/item?id=48642106)
- Source: [tenureai.dev](https://tenureai.dev/writing/open-knowledge-format-ai-memory-benchmark/)
- Score: 4
- Comments: 0
- Posted: 2026-06-23T08:42:10Z

## Translation

タイトル: AI ツール内で OKF を実行すると何が起こるか
記事のタイトル: オープン ナレッジ フォーマット ベンチマーク: AI ツール内で OKF を実行すると何が起こるか |在職期間
説明: 私たちは Google を経営しました

記事本文:
在職期間
プラットフォーム ▼ コア機能
Open Knowledge Format ベンチマーク: AI ツール内で OKF を実行すると何が起こるか
Google は、AI ナレッジの単純なマークダウン標準として Open Knowledge Format を廃止しました。
実際に重要な部分、つまりそのバンドルが配置されたときに何が起こるかをテストしたかったのです。
AI ツール内で、モデルはどのファイルを検査するかを決定する必要があります。
これはストレージ形式としてのマークダウンのテストではありませんでした。
これは、バンドルが今日 AI ツールにドロップされるときに OKF によって暗示される実行時パターンのテストでした。
モデルは、ファイルのリスト表示とファイルの読み取りのための通常のツール アクセスを受け取りました。カスタムの取得エージェント システム プロンプトはありません。
PrecisionMemBench は、モデルが実際に読み取ったファイルに対応する信念 ID をスコア付けしました。
OKF は生のセマンティック メモリの取得よりもはるかに優れていましたが、依然として同じ中心的な問題がありました。それは、ファイルは形式であり、取得ポリシーではないということです。
OKFが好きです。 AI ツールを実際に使用する必要があるときに何が起こるかを知りたかっただけです。
私は Open Knowledge Format の方向性が好きです。 YAML を使用したマークダウン ファイルのディレクトリ
フロントマターは良い意味で退屈だ。読みやすく、差分も簡単、コミットも簡単です。
人間にとってメンテナンスが容易です。これは、独自のメモリ BLOB や隠されたメモリよりも大きな利点です。
ベンダーインデックス。
しかし、仕様を読んだ後、明らかな疑問は、マークダウンが良い交換であるかどうかではありませんでした。
形式。マークダウンは大丈夫です。問題は、実行時に何が起こるかということでした。
OKF バンドルが AI ツールに配置された瞬間、誰かがどのバンドルを決定する必要があるためです。
ファイルはモデルリクエストを入力します。おそらくモデルにはファイルがリストされていると思います。たぶんそれは1つを読むでしょう。たぶんそれは読む
いくつか。おそらく、必要な信念が実際に含まれているファイルは決して開かれないでしょう。フォーマットにより、
知識はポータブルです。それ自体では検索が正確になるわけではありません。
私たちが

テストされたのは、ストレージ レイヤーとしての Open Knowledge Format ではありませんでした。アクセスパターンをテストしました
今すぐ OKF バンドルをツールにドロップし、モデルに検査させれば、ほとんどのチームがそれを取得できるでしょう。
包装紙はあえて無地にしました。モデルには通常のアシスタント プロンプトと 2 つのツールがありました。
1 つは利用可能な OKF マークダウン ファイルを一覧表示するため、もう 1 つは特定のファイルを読み取るためです。モデルはそうではありませんでした
動作方法を指示するカスタムの取得エージェント プロンプトが与えられます。ユーザーからの問い合わせを受け取りました
そしてバンドルを検査するかどうかを決定しました。
モデルが read_file を呼び出すと、ラッパーはファイルのフロントマターを解析し、抽出した
believeId を取得し、その ID を取得された信念として PrecisionMemBench に返しました。それが、
橋全体。モデルによって読み取られたファイルは、取得された信念 ID になりました。読み取られなかったファイルはカウントされません。
PrecisionMemBench は最終的な答えが適切に聞こえるかどうかを評価しないため、これは重要です。
システムが根底にある正しい信念を取得したかどうかをスコア化します。今回の実行で問題となったのは、
単純: モデル主導のファイル アクセスは、適切な OKF ドキュメントをリクエストに取り込みましたか?
77 ケース、36 パス、18 アクティブ検索パス、平均再現率 0.91、平均待ち時間 4.4 秒。
12 セッション ターン、2 パス、1 アクティブ検索パス、平均リコール 0.45、p95 遅延 59.3 秒。
OKF は問題の形状を改善しました。
良いニュースは、OKF が生のベクトルのリコールのように動作しなかったことです。モデルにはファイル名、タイトル、
説明、タイプ、タグ、および読み取り可能なファイル本体。これにより、コサイン検索よりも多くのハンドルが得られました
不透明なメモリチャンク上。
エイリアスの解決が最も明確な勝利でした。 23 のエイリアス ケース全体で、実行の平均精度は 0.72 に達しました。
0.92 は再現率を意味します。いくつかの短い形式のクエリは、期待どおりに機能しました。のクエリ
GHA は、モデルを GitHub Actions の信念に導く可能性があります。 Mongo に対するクエリは次のような結果をもたらす可能性があります
それをMongoDBに転送します

決断。このような場合、ファイルシステム パターンはモデルに実際のパスを与えます。
右のドキュメントに。
ランキングの安定性も強そうだ。それらの訴訟は無事に通過しました。それは言う価値があるからです
つまり、OKF に対する包括的な批判ではないということです。クエリがファイル表面にきれいにマッピングされると、
フロントマターを含むマークダウン ファイルは完全に合理的な表現です。
しかし、ファイルアクセスは依然としてメモリ検索ではありません。
障害は、スコープ、スーパーセッション、タイプ ルーティングなど、メモリ システムが通常障害を起こす箇所で発生しました。
セッションドリフト。これらはマークダウンの問題ではありません。それらは国家の問題なのです。
範囲の明確化には 12 件のケースがあり、合格したのは 4 件のみでした。平均精度は0.21でした。これが古典です
Redisの問題。 Redis がコード コンテキストと書き込みコンテキストに出現する場合、モデルはそれを認識する必要があります。
現在のリクエストにどれが属しているか。ファイルのフォルダーは両方の意味を公開できますが、実際には
このターンにどれが有効であるかは強制されません。
スーパーセッションはさらに鋭かった。実行にはスーパーセッション チェーンの除外ケースが 3 件ありましたが、何もありませんでした。
合格しました。繰り返しますが、これはマークダウンが古い信念と新しい信念を保存できないからではありません。できる。
問題は、ランタイムがどの信念が現在のもので、どの信念を除外する必要があるかを認識する必要があることです。
現在のファイルの隣にある古いファイルには、取得層が認識しない限りアクセスできます。
違い。
タイプルーティングと未解決の質問の平均精度は 0.20 でした。ファジーマッチングとプレフィックスガードが登場
0.25で。予算立ち退きとキャパシティは 0.13 に達しました。ここが知識との違いです
フォーマットとメモリシステムが非常に明白になります。 OKF は知識を説明できます。それは決まりません
その知識のどれだけをモデルリクエストに入力する必要があるか。
このバンドルにより、知識を検査できるようになりました。検索を管理するものではありませんでした。それがギャップです。
セッションの実行

ここでパターンが不安定になりました。
シングルターンの結果は楽観的な見方だ。セッションの実行は人々が実際に使用する方法に近いものになります
AIツール。話題が動き回ります。早いターンではノイズが発生します。後のいくつかのターンは暗黙的です。モデル
ファイルを検査するかどうか、どのファイルが重要か、現在のターンがまだ残っているかどうかを判断し続ける必要があります
同じトピックに属します。
その実行では、12 ターン中 2 ターンだけが通過しました。平均再現率は 0.45 に低下し、アクティブなものは 1 つだけでした
回収パス。 p95 の遅延は 59.3 秒に達しました。正確な数はパターンほど重要ではありません。
取得がセッション内で繰り返されるモデル動作になると、コストと不安定性が増大します。
これは、製品内に OKF を採用する場合に最も注目する部分です。単一のクエリ
クリーンなファイルバンドルと比較すると、かなり見栄えがよくなります。実際のセッションには、継続性、除外性、最新性が必要です。
範囲と予算を毎ターン管理します。
OKF はインターチェンジを解決します。選択を解決するものではありません。
これが私にとっての主なポイントです。 OKF は交換問題に対する優れた答えです。それは人々に与える
移植可能なナレッジバンドルを作成するクリーンな方法。それは貴重なことだ。それは、チームが厳選された動きをすることができることを意味します
1 つのベンダーのメモリ データベースに依存せずにツール間のコンテキストを共有します。
しかし、メモリの検索には選択という 2 番目の問題があります。どの信念が関連していますか?どれが現行のものですか？
このユーザー、チーム、リポジトリ、顧客、またはタスクにスコープされるのはどれですか?どの事実が固定されていますか?どの事実が
古い？どの未解決の質問を表示する必要があり、どれをリクエストから除外する必要がありますか?
ファイル形式には、これらの質問に答えるために必要なフィールドを含めることができます。それだけではそれらに答えることはできません
存在する。ランタイムはそれらを強制する必要があります。
フォーマットは難しい部分ではありません。
マークダウン ファイルをツールにドロップするのは素晴らしいスタートです。透明で持ち運びが簡単で、簡単に組み立てることができます。

n
について。しかし、モデルがどのファイルを読み取るかを選択する必要があると、再び検索の世界に戻ります。
ファイル、タイトル、説明、フロントマターが重要であるため、OKF の実行は単純なメモリ取得よりも優れた結果をもたらしました。
モデルに便利なハンドルを与えます。しかし、ハードケースは決して難しいものではなかったので、ハードケースはハードのままでした。
マークダウン。それらは国家に関するものでした。
したがって、私はフォーマットとしての OKF については強気です。私はこのフォーマットを記憶システムとして扱うことに強気ではありません。
次の層はランタイムである必要があります。これは、何を取得するか、何を除外するかを決定します。
現在の内容、対象範囲、モデルが実際に知ることを許可されているもの。

## Original Extract

We ran Google

tenure
Platform ▼ Core Capabilities
Open Knowledge Format benchmark: what happens when OKF runs inside an AI tool
Google dropped the Open Knowledge Format as a simple markdown standard for AI knowledge.
We wanted to test the part that matters in practice: what happens when that bundle is placed
inside an AI tool and the model has to decide which files to inspect.
This was not a test of markdown as a storage format.
It was a test of the runtime pattern implied by OKF when a bundle is dropped into an AI tool today.
The model received normal tool access to list files and read files. No custom retrieval-agent system prompt.
PrecisionMemBench scored the belief IDs corresponding to the files the model actually read.
OKF did much better than raw semantic memory retrieval, but it still showed the same core issue: files are a format, not a retrieval policy.
I like OKF. I just wanted to see what happens when an AI tool actually has to use it.
I like the direction of the Open Knowledge Format. A directory of markdown files with YAML
frontmatter is boring in the best way. It is easy to read, easy to diff, easy to commit, and
easy for humans to maintain. That is a real advantage over proprietary memory blobs or hidden
vendor indexes.
But after reading the spec, the obvious question was not whether markdown is a good interchange
format. Markdown is fine. The question was what happens at runtime.
Because the moment an OKF bundle gets placed into an AI tool, somebody still has to decide which
files enter the model request. Maybe the model lists the files. Maybe it reads one. Maybe it reads
several. Maybe it never opens the file that actually contains the needed belief. The format makes
the knowledge portable. It does not, by itself, make retrieval precise.
The thing we tested was not Open Knowledge Format as a storage layer. We tested the access pattern
most teams would get if they dropped an OKF bundle into a tool today and let the model inspect it.
The wrapper was intentionally plain. The model got a normal assistant prompt and two tools:
one to list available OKF markdown files, and one to read a specific file. The model was not
given a custom retrieval-agent prompt telling it how to behave. It just received a user query
and decided whether to inspect the bundle.
When the model called read_file , the wrapper parsed the file's frontmatter, extracted
the beliefId , and returned that ID to PrecisionMemBench as a retrieved belief. That is the
whole bridge. Files read by the model became retrieved belief IDs. Files not read did not count.
That matters because PrecisionMemBench does not score whether the final answer sounds good.
It scores whether the system retrieved the right underlying beliefs. In this run, the question was
simple: did model-directed file access pull the right OKF documents into the request?
77 cases, 36 passes, 18 active retrieval passes, 0.91 mean recall, and 4.4s mean latency.
12 session turns, 2 passes, 1 active retrieval pass, 0.45 mean recall, and 59.3s p95 latency.
OKF improved the shape of the problem.
The good news is that OKF did not behave like raw vector recall. The model had filenames, titles,
descriptions, types, tags, and readable file bodies. That gave it more handles than a cosine search
over opaque memory chunks.
Alias resolution was the clearest win. Across 23 alias cases, the run reached 0.72 mean precision
and 0.92 mean recall. Some short-form queries worked exactly how you would hope. A query for
GHA could lead the model to the GitHub Actions belief. A query for Mongo could lead
it to the MongoDB decision. In those cases, the filesystem pattern gives the model a real path
to the right document.
Ranking stability also looked strong. Those cases passed cleanly. That is worth saying because it
means the result is not a blanket criticism of OKF. When the query maps cleanly to the file surface,
markdown files with frontmatter are a perfectly reasonable representation.
But file access is still not memory retrieval.
The failures showed up where memory systems usually fail: scope, supersession, type routing, and
session drift. These are not markdown problems. They are state problems.
Scope disambiguation had 12 cases and only 4 passed. Mean precision was 0.21. This is the classic
Redis problem. If Redis appears in a code context and a writing context, the model has to know
which one belongs in the current request. A folder of files can expose both meanings, but it does
not enforce which one is valid for this turn.
Supersession was even sharper. The run had three supersession chain exclusion cases, and none
passed. Again, that is not because markdown cannot store an old belief and a new belief. It can.
The issue is that the runtime has to know which belief is current and which belief must be excluded.
A stale file sitting next to a current file is still accessible unless the retrieval layer knows the
difference.
Type routing and open questions had 0.20 mean precision. Fuzzy matching and prefix guards landed
at 0.25. Budget eviction and capacity landed at 0.13. This is where the difference between a knowledge
format and a memory system becomes very obvious. OKF can describe the knowledge. It does not decide
how much of that knowledge should enter the model request.
The bundle made knowledge inspectable. It did not make retrieval governed. That is the gap.
The session run is where the pattern got shaky.
The single-turn result is the optimistic view. The session run is closer to how people actually use
AI tools. Topics move around. Earlier turns create noise. Some later turns are implicit. The model
has to keep deciding whether to inspect files, which files matter, and whether the current turn still
belongs to the same topic.
In that run, only 2 of 12 turns passed. Mean recall dropped to 0.45, and there was only 1 active
retrieval pass. The p95 latency reached 59.3 seconds. The exact number matters less than the pattern:
once retrieval becomes a repeated model behavior inside a session, the cost and instability compound.
This is the part I would watch most closely if I were adopting OKF inside a product. A single query
against a clean file bundle can look pretty good. A real session needs continuity, exclusion, recency,
scope, and budget control on every turn.
OKF solves interchange. It does not solve selection.
This is the main takeaway for me. OKF is a good answer to the interchange problem. It gives people
a clean way to write portable knowledge bundles. That is valuable. It means a team can move curated
context between tools without depending on one vendor's memory database.
But memory retrieval has a second problem: selection. Which beliefs are relevant? Which are current?
Which are scoped to this user, team, repo, customer, or task? Which facts are pinned? Which facts are
stale? Which open questions should be visible, and which should stay out of the request?
A file format can carry the fields needed to answer those questions. It cannot answer them by merely
existing. The runtime has to enforce them.
The format is not the hard part.
Dropping markdown files into a tool is a great start. It is transparent, portable, and easy to reason
about. But once the model has to choose which files to read, you are back in retrieval land.
The OKF run did better than naive memory retrieval because files, titles, descriptions, and frontmatter
give the model useful handles. But the hard cases remained hard because the hard cases were never about
markdown. They were about state.
So I am bullish on OKF as a format. I am not bullish on treating the format as the memory system.
The next layer has to be the runtime: the thing that decides what gets retrieved, what gets excluded,
what is current, what is scoped, and what the model was actually allowed to know.
