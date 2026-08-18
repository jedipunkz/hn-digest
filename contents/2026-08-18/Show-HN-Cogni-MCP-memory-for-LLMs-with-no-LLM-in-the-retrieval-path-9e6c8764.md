---
source: "https://getcogni.io/"
hn_url: "https://news.ycombinator.com/item?id=49339547"
title: "Show HN: Cogni: MCP memory for LLMs, with no LLM in the retrieval path"
article_title: "Cogni — MCP memory server for Claude, ChatGPT and AI agents"
image: "https://getcogni.io/og.png"
author: "ihamilton7"
captured_at: "2026-08-18T00:39:13Z"
capture_tool: "hn-digest"
hn_id: 49339547
score: 2
comments: 0
posted_at: "2026-08-18T00:21:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Cogni: MCP memory for LLMs, with no LLM in the retrieval path

- HN: [49339547](https://news.ycombinator.com/item?id=49339547)
- Source: [getcogni.io](https://getcogni.io/)
- Score: 2
- Comments: 0
- Posted: 2026-08-18T00:21:14Z

## Translation

タイトル: HN を表示: Cogni: LLM の MCP メモリ (取得パスに LLM なし)
記事のタイトル: Cogni — クロード、ChatGPT、AI エージェント用の MCP メモリ サーバー
説明: エンティティ グラフ拡散リコールにより、ベクトル ストアが実行できるファクトのチェーンを再構築する MCP メモリ サーバー

記事本文:
C " />
C
コグニ
比較する
仕組み
価格設定
アップデート
ドキュメント
サインイン
無料で始めましょう
LLM の認知層
記憶以上のもの。
理由づける記憶。
Cogni は、作業に応じて増加する接続メモリを LLM に提供する MCP サーバーです。それ
エンティティ グラフ拡散アクティベーションは事実の連鎖に従います。
プレーン ベクター ストアはアクセスできないため、エージェントは単に調べるのではなく、記憶している内容を推論します。
物事はうまくいきました。
Chain-recall@8: 単一のリコール コールが表面化する 4 つのドキュメント チェーンのシェア — Cogni が発見
4 のうち 3.4 、ベクトル ストアでは 2.6 が見つかります。答えの正確さではありません。
決定的で、ループ内にモデルがなく、出荷時のリコール パスでは、すべてのアカウントが次の設定を取得します。
300 枚のディストラクタ文書に対して 5 チェーン × 3 シード。両面
同じ 8 行を返すため、追加のコンテキストでマージンが購入されることはありません。
すべての数字と、同点または 2 位になる場所。
✓ あらゆるモデルで動作: クロード、GPT、ジェミニ、ローカル
✓ 取得パスに LLM も GPU もありません
✓ 実行するベクター DB がありません
✓ 貼り付ける API キーはありません
ベクトル ストアが取得します。コグニが接続します。
違いは、明確で測定された 1 つの機能です。それは、語彙横断検索のためのエンティティ グラフ拡散アクティベーションです。答えが質問と単語を共有しない質問であり、 k をどれだけ広く設定しても、単一の類似性検索のスコアは 0 になります。
質問と語彙を共有しない文書内に答えが存在するもの、つまり名前につながるバッチにつながる部品番号を尋ねてください。 1 回の取得呼び出しで、Cogni はエンティティ チェーンをたどります。類似性検索は最初のリンクと一致しますが、2 番目のリンクに移動するものがないため、1 回の呼び出しではスコアが 0 になります。
各アカウントが取得する設定で Cogni のデフォルト モードで、GPT-4o 応答、1 回のリコール コールを使用して MCP 上でエンドツーエンドで測定 - 実行 14 回、キュー 45 回の平均

それぞれステーション。ゼロは構造部分です。Cogni 独自のグラフをオフにすると、同じストア上で同じエンベディングと同じプロンプトを使用してグラフが正確に再現されるため、違いはトラバースだけであり、他には何もありません。
モデルが 1 つの呼び出しで停止するのではなく呼び出しを継続すると、ベクトル ストアはほぼ同等の状態に回復します。最終的には各リンクに行き当たります。これが結果の正直な形状であり、依然として議論の余地があります。答えはどちらの方法でも到達可能です。異なるのは、料金を支払う往復回数です。
Cogni は最初の呼び出しでチェーンを返します。代替案は反復することで到達し、各反復は、使用しているモデルに関係なく、別のリクエスト、別のコンテキスト ウィンドウ、および別の請求になります。
チェーンの長さに応じてギャップは広がります。余分なホップごとに、類似性検索で発見する必要がある往復が 1 回増え、1 つの Cogni がすでに歩いています。
グラフの下で完全なベクトル検索が実行されます。 Cogni のデフォルトのリコールでは、密ベクトル検索と拡散グラフが実行され、両方が返されます。そのため、Cogni を追加しても、すでに行っていた検索機能が失われることはありません。これは、1 回の電話で確認できる下限であり、約束ではありません。干し草の山に針を入れると、私たちが測定したすべての干し草の山の長さで、2 つは 1.00 から 1.00 で水平になります。
どのモデルでも動作します。ループ内に LLM がありません。
取得パスのどこにも言語モデルも GPU もありません。
途中であなたの記憶を書き換えたり、要約したり、判断したりするものは何もありません。
抜け出す方法。リコールは決定論的なグラフ走査です: 同じです
同じストアに関する質問は、毎回同じ行を返します。2 番目のモデルの意見はありません。
の間。これが利点が伝わる理由です。利点はモデルではなく検索に存在します。そして
それが、小規模なモデルが最も得られる理由です。きれいに接続されたコンテキストは、モデルに欠けていた部分でした。
であること

天井に向かってまっすぐです。コンテキスト内のコーパス全体を含む GPT-4o のスコアは 0.99 なので、これは同等ではありません。これは、コンテキストの一部で、以前はまったく実行できなかった作業を実行する小さなローカル モデルです。同じ 8B に正確なチェーンを手で与えると、スコアは 0.97 となり、残りのギャップは推論ではなく検索であることがわかります。
マルチホップ演繹的ベンチマーク、250 のディストラクタ文書に対する 12 チェーン × 3 シード、llama3.1:8b 応答、各アカウントが取得する設定で出荷された取得パスを通じて測定 - 結果ファイルにはそれらからの逸脱が記録されません。ベクター ベースラインには、Cogni の 8 行に 15 行が与えられました。Cogni のツールはモデルによって呼び出されるため、実際の一貫性はモデルとそのプロンプトの方法によって異なります。結果が異なる場合がある理由。
記憶力が増加しても、想起率は横ばいになります。 32,000 ワードに達する履歴にファクトを埋め込むと、履歴の保持にかかるコストに関係なく、Cogni は毎回、クエリごとに固定の 8 行でそのファクトを返します。
事実上、一度にすべてを保持できないモデル上で、クエリごとに一定のコストでローリングされる常時オンのコンテキスト ウィンドウ。
ベクター ストアでは他にできないこと
完全なベクトル検索はグラフの下で実行されるため、Cogni を追加しても、すでに行っていた検索機能が失われることはありません。私たちはそれを主張するのではなく、テストしています。今月のリリースでは、リコールで返される行の約半分がグラフ結果に置き換えられ、干し草の山から針を取り出しても、回答はすべての質問でランク 1 に留まりました。
すべての記憶には保存時にタイムスタンプが付けられ、ある瞬間やウィンドウを思い出すことができます。「先週の火曜日の時点で私は何を知っていたでしょうか?」 , 「3月に拾ったものだけ」。ウィンドウを要求すると、Cogni がその中を検索します。
類似性検索には独自の時間モデルはありません。意味の近さに基づいてランク付けされるため、順序付けと時点の調査が行われます。

オンは、誰かがその横に追加したタイムスタンプ インデックスに依存します。 Cogni は、保存されるすべての記憶にスタンプを付け、ランク付けする前にウィンドウを適用します。そのため、日付の古い質問は、ストア内の最も近い文言からではなく、ウィンドウに表示されていた内容から回答されます。
Cogni は独自の LLM や GPU を実行しません。利点は検索にあるため、どのモデルを指定しても保持されます。また、そのモデルがフロンティアであっても、机上にある 80 億のパラメーターであっても、コストは同じです。
思い出は 10 個ありますが、そのどれもがこの質問のために書かれたものではありません。答えは事実から 2 つ離れており、質問された内容とは何の言葉も交わしていません。各検索で実際に返されるものは次のとおりです。同じ店舗、同じ質問、同じ 8 行の予算です。関連性下限ドロップの一致が弱すぎて返す価値がないため、ベクトル列は短くなります。予算を拡大しても、それらの要素がさらに追加されるだけで、答えは得られません。
hop 0 Atlas の移行が完了するまで、Meridian ロールアウトはブロックされます。
hop 0 ロールアウトのスタンドアップは火曜日 9 時に変更されました。
hop 0 Priya Raman は、Meridian ロールアウトのデリバリー リーダーです。
hop 0 Helios のパイロットは 2 週間早く終了しました。
hop 1 Devonokafor はプラットフォーム インフラストラクチャ チームを率いています。
ホップ 1 マーカス・ウェッブは、顧客対応のあらゆることについてループインするよう求めました。
hop 1 新しいオンボーディングフローの設計レビューは 14 日に行われます。
hop 1 四半期計画は来月の第 1 週です。
ホップは、Cogni がそのメモリに到達するまでにたどった接続の数です。ホップ 1 にあるものはすべて、質問との照合ではなく、グラフを走査することによって見つかりました。
1 Meridian ロールアウトは、Atlas の移行が完了するまでブロックされます。
2 Priya Raman は、Meridian ロールアウトのデリバリーリーダーです。
3 マーカス・ウェッブは、顧客対応のあらゆることについてループインするよう求めました。
4 Atlas の移行は、プラットフォーム インフラストラクチャ チームが所有します。
ここで止まります。 4行目は1つのCOです

答えからは遠ざかっています。プラットフォームチームの名前は明記されていますが、それを運営する人物の名前は記載されていません。質問には「Devonokafor」に似ているものは何もないため、類似性検索を行う必要はありません。
イラストではなく実際の実行: デフォルト設定でのライブ Cogni ストア、同じ運用環境の埋め込み、MCP クライアントが受け取る同じ順序。 tools/make_demo_transcript.py から再現可能です。ベクトル列は、グラフがオフになっている Cogni 独自の高密度チャネルであるため、両側はエンベッダーを共有し、トラバーサルのみが異なります。
一度アレルギーについて言及すると、6 週間後もモデルがレストランを予約する前に Cogni がそのアレルギーを表面化します。期限、設定、参加者は、チャット間で失われることなく、すべてのセッションにわたって接続を維持します。
キーワードを共有していない 1,000 のドキュメントにわたって関連するケースを追跡し、フラット検索では組み立てることができない概要を作成します。質問と一致することではなく、エンティティ リンクをたどることでチェーンに到達します。
使用例の説明。それらの測定結果は比較ページにあります。
Cogni は LLM を実行しません。あなたのクライアントは脳です。 MCP 上で少数のツールを公開し、ベクター ストアではできない接続された検索を実行します。検索基質は完全に決定的です。
事実、メモ、観察を保存します。エンティティはグラフにリンクされます。
伝えた内容すべてに拡散アクティベーション検索が実行され、その下で完全なベクトル検索が実行されるため、Cogni を追加しても、すでに提供されていた検索機能が失われることはありません。オプションのポイントインタイムおよびウィンドウリコール。
モデルの人間による推定速度を、エージェント自身が測定した中央値に置き換えます。
候補者の計画を実際にかかる時間に基づいてランク付けするため、エージェントは真に高速なパスを選択します。
実際の期間を記録すると、将来の推定がより正確になります。
記憶しているすべてを網羅する 1 つのエンティティ グラフ

。そのグラフは、記憶することと推論することの違いです。
クラスター、GPU、ベクトル データベースはありません。 Cogni は 1 つのリモート エンドポイント https://mcp.getcogni.io/mcp であり、接続には約 1 分かかります。クライアントに応じて 2 つの方法があります。
コネクタ アプリ: URL を使用して Cogni をカスタム コネクタとして追加し、サインインします。貼り付けるキーはありません。OAuth が処理します。
1. [設定] → [コネクタ] → [カスタム コネクタを追加] 。
2. https://mcp.getcogni.io/mcp を貼り付け、OAuth フィールドを空白のままにして、「追加」をクリックします。
3. ブラウザが開いたら、「接続」をクリックしてサインインします。
無料プランでは 1 つのカスタム コネクタが使用できます。有料プランでは複数が許可されます。
1. 開発者モードをオンにします (設定 → コネクタ → 詳細)。
2. コネクタを追加し、 https://mcp.getcogni.io/mcp を貼り付けて、サインインします。
さらに、Web/デスクトップ上の Pro、Business、Enterprise、または Edu (モバイル アプリや無料ではありません)。
開発者およびエージェントのクライアントは、自分で作成した Cogni API キーを使用してサインインします。 GUI アプリ (トラック 1) を接続し、「Cogni API キーの作成」を要求します。sk-cogni-… キーが 1 回表示されます。ツールの構成を以下にドロップし、そのキーを Authorization: Bearer ヘッダーに貼り付けます。
クロード mcp add --transport http --scope user cogni \
https://mcp.getcogni.io/mcp \
--header "認証: ベアラー sk-cogni-YOUR_KEY"
カーソル⌄
~/.cursor/mcp.json に追加し、[設定] → [ツールと MCP] → [cogni] を有効にします。
{
"mcpサーバー": {
"認識": {
"url": "https://mcp.getcogni.io/mcp",
"headers": { "認可": "ベアラー sk-cogni-YOUR_KEY" }
}
}
}
クライン⌄
[MCP サーバー] → [MCP サーバーの構成] を選択し、 mcpServers の下に追加します。
{
"mcpサーバー": {
"認識": {
"タイプ": "streamableHttp",
"url": "https://mcp.getcogni.io/mcp",
"headers": { "認可": "ベアラー sk-cogni-YOUR_KEY" }
}
}
}
VS Code · Copilot が必要です ⌄
.vscode/mcp.json を作成し、サーバーを起動して Copilot Chat のエージェント モードで使用します。
{
」

サーバー": {
"認識": {
"タイプ": "http",
"url": "https://mcp.getcogni.io/mcp",
"headers": { "認可": "ベアラー sk-cogni-YOUR_KEY" }
}
}
}
ウィンドサーフィン ⌄
~/.codeium/windsurf/mcp_config.json に追加し (キーが serverUrl であることに注意してください)、カスケードで更新します。
{
"mcpサーバー": {
"認識": {
"serverUrl": "https://mcp.getcogni.io/mcp",
"headers": { "認可": "ベアラー sk-cogni-YOUR_KEY" }
}
}
}
オープンクロウ ⌄
1 つのコマンド (ヘッダー サポートの場合は v2026.5.12 以降):
openclaw mcp 認識を追加 \
--url https://mcp.getcogni.io/mcp \
--transport streamable-http \
--header "認証: ベアラー sk-cogni-YOUR_KEY"
ジェミニ CLI ⌄
~/.gemini/settings.json に追加します (キーは httpUrl です):
{
"mcpサーバー": {
"認識": {
"httpUrl": "https://mcp.getcogni.io/mcp",
"headers": { "認可": "ベアラー sk-cogni-YOUR_KEY" }
}
}
}
有料の Gemini API キーまたは Code Assist Standard/Enterprise ライセンスが必要です (Google の 2026 年 6 月の変更による)。
次に、Cogni スキルをロードします。接続するとアシスタントにツールが提供されます。スキルはそれらをいつ使用すべきかを教えるので、答える前に思い出して、毎回質問しなくても重要なことを覚えてくれます。それは 1 つの小さなファイルです。スキルをダウンロードし、クライアントに追加する方法。
すべてのクライアントに 1 つの思い出。 OAuth またはキー、GUI またはターミナル — サインインするすべての方法が同じ分離ストアに配置されます。サインイン

[切り捨てられた]

## Original Extract

An MCP memory server whose entity-graph spreading recall reassembles the chain of facts a vector store can

C " />
C
Cogni
Compare
How it works
Pricing
Updates
Docs
Sign in
Get started free
The cognition layer for LLMs
More than memory.
Memory that reasons.
Cogni is an MCP server that gives any LLM a connected memory that grows as you work. Its
entity-graph spreading activation follows the chain of facts a
plain vector store can't reach, so your agent reasons over what it remembers instead of just looking
things up.
Chain-recall@8: the share of a four-document chain a single recall call surfaces — Cogni finds
3.4 of the 4 , a vector store finds 2.6. Not answer accuracy.
Deterministic, no model in the loop, on the shipped recall path at the settings every account gets:
5 chains × 3 seeds against 300 distractor documents. Both sides
return the same eight rows , so none of the margin is bought with extra context.
Every number, and where we tie or come second.
✓ Works with any model: Claude, GPT, Gemini, local
✓ No LLM and no GPU in the retrieval path
✓ No vector DB to run
✓ No API key to paste
A vector store retrieves. Cogni connects .
The difference is one clean, measured capability: entity-graph spreading activation for cross-vocabulary retrieval — questions whose answer shares no words with the question, where a single similarity lookup scores zero no matter how wide you set k .
Ask something whose answer lives in a document sharing no vocabulary with the question — a part number that leads to a batch that leads to a name. In a single retrieval call Cogni follows the entity chain. A similarity search matches the first link and has nothing to take it to the second, so on one call it scores zero.
Measured end to end over MCP with GPT-4o answering, one recall call, on Cogni's default mode at the settings every account gets — mean of 14 runs, 45 questions each. The zero is the structural part : turning Cogni's own graph off reproduces it exactly, on the same store with the same embeddings and the same prompt, so the difference is the traversal and nothing else.
Let the model keep calling instead of stopping at one, and a vector store recovers to near-parity — it eventually stumbles onto each link. That is the honest shape of the result, and it is still the argument: the answer is reachable either way; what differs is how many round trips you pay for it.
Cogni returns the chain on the first call. The alternative gets there by iterating, and every iteration is another request, another context window, and another bill from whichever model you are using.
The gap widens with the length of the chain: each extra hop is one more round trip a similarity search has to discover, and one Cogni already walked.
A full vector search runs underneath the graph. Cogni's default recall runs a dense-vector search and the spreading graph, then returns both — so adding Cogni does not take away the retrieval you already had. It is a floor you can check in one call, not a promise: on needle-in-a-haystack recall the two are level at 1.00 to 1.00, at every haystack length we measured.
Works with any model. No LLM in the loop .
No language model and no GPU anywhere in the retrieval path.
Nothing rewrites, summarises, or judges your memories on the way in, and nothing is generated on the
way out. Recall is a deterministic graph traversal : the same
question over the same store returns the same rows, every time, with no second model's opinion in
between. That is why the advantage travels — it lives in the retrieval, not in the model. And
it is why a small model gains the most: clean connected context was the piece it was missing.
To be straight about the ceiling: GPT-4o with the entire corpus in context scores 0.99, so this is not parity — it is a small local model doing work it could not do at all before, on a fraction of the context. Feed that same 8B the exact chain by hand and it scores 0.97, which says the remaining gap is retrieval rather than reasoning.
Multi-hop deductive benchmark, 12 chains × 3 seeds against 250 distractor documents, llama3.1:8b answering, measured through the shipped retrieval path at the settings every account gets — the results file records no deviation from them. The vector baseline was given 15 rows to Cogni's 8. Because Cogni's tools are called by your model, real-world consistency depends on the model and how it's prompted. Why results may vary .
Recall stays flat as memory grows. Bury a fact in a history that grows to 32,000 words and Cogni returns it first, every time — at a fixed eight rows per query, whatever the history costs to hold.
Effectively a rolling, always-on context window at a constant per-query cost, on a model that could never hold it all at once.
What else it does that a vector store can’t
A full vector search runs underneath the graph, so adding Cogni does not take away the retrieval you already had. We test that rather than claim it: a release this month replaced roughly half the rows a recall returns with graph results, and on needle-in-a-haystack the answer stayed at rank one in every single question.
Every memory is timestamped as it is stored, and recall can be pointed at a moment or a window — "what did I know as of last Tuesday?" , "only what I picked up in March" . Ask for a window and Cogni searches inside it.
Similarity search has no time model of its own — it ranks on closeness of meaning, so ordering and point-in-time questions depend on whatever timestamp index someone bolted on beside it. Cogni stamps every memory as it is stored and applies the window before it ranks, so a dated question is answered from what was in the window, not from the closest wording anywhere in the store.
Cogni runs no LLM of its own and no GPU. The advantage is in the retrieval, so it holds whichever model you point at it — and it costs the same whether that model is frontier or 8 billion parameters on your desk.
Ten memories, none of them written for this question. The answer is two facts away and shares no words with what was asked. Here is what each retrieval actually returns — same store, same question, same eight-row budget. The vector column is shorter because a relevance floor drops matches too weak to be worth returning; widening the budget adds more of those, not the answer.
hop 0 The Meridian rollout is blocked until the Atlas migration finishes.
hop 0 Standups for the rollout moved to Tuesdays at 9.
hop 0 Priya Raman is the delivery lead for the Meridian rollout.
hop 0 The Helios pilot wrapped up two weeks early.
hop 1 Devon Okafor runs the platform infrastructure team.
hop 1 Marcus Webb asked to be looped in on anything customer-facing.
hop 1 The design review for the new onboarding flow is on the 14th.
hop 1 Quarterly planning is the first week of next month.
hop is how many connections Cogni followed to reach that memory. Everything at hop 1 was found by traversing the graph, not by matching the question.
1 The Meridian rollout is blocked until the Atlas migration finishes.
2 Priya Raman is the delivery lead for the Meridian rollout.
3 Marcus Webb asked to be looped in on anything customer-facing.
4 The Atlas migration is owned by the platform infrastructure team.
It stops here. Row 4 is one connection away from the answer — the platform team is named, and the person who runs it is not. Nothing in the question resembles “Devon Okafor”, so similarity search has no step left to take.
A real run, not an illustration: a live Cogni store at default settings, the same production embeddings, and the same ordering an MCP client receives. Reproducible from tools/make_demo_transcript.py . The vector column is Cogni’s own dense channel with the graph switched off, so both sides share the embedder and differ only in the traversal.
Mention an allergy once, and six weeks later Cogni still surfaces it before the model books the restaurant. Deadlines, preferences, and people stay connected across every session, not lost between chats.
Trace a connected case across a thousand documents that share no keywords, then write the brief a flat search can't assemble — the chain is reached by following entity links, not by matching the question.
Illustrative use cases. The measured results behind them are on the compare page .
Cogni runs no LLM . Your client is the brain. It exposes a handful of tools over MCP and does the connected retrieval a vector store can't. The retrieval substrate is fully deterministic.
Store a fact, note, or observation. Entities are linked into a graph.
Spreading-activation retrieval across everything you've told it, with a full vector search running underneath — so adding Cogni doesn't take away the retrieval you already had. Optional point-in-time and windowed recall.
Replace the model's human-speed guess with the agent's own measured median.
Rank candidate plans by how long they'll actually take, so the agent picks the genuinely faster path.
Log a real duration so future estimates get sharper.
One entity graph, spanning everything it remembers. That graph is the difference between remembering and reasoning.
No cluster, no GPU, no vector database. Cogni is one remote endpoint, https://mcp.getcogni.io/mcp , and connecting takes about a minute. There are two ways in, depending on your client.
Connector apps: add Cogni as a custom connector with the URL, then sign in. No key to paste — OAuth handles it.
1. Settings → Connectors → Add custom connector .
2. Paste https://mcp.getcogni.io/mcp , leave the OAuth fields blank, click Add.
3. Click Connect and sign in when the browser opens.
Free plan allows one custom connector; paid plans allow several.
1. Turn on Developer mode (Settings → Connectors → Advanced).
2. Add a connector, paste https://mcp.getcogni.io/mcp , sign in.
Plus, Pro, Business, Enterprise or Edu, on web/desktop (not the mobile app or Free).
Dev & agent clients sign in with a Cogni API key you mint yourself. Connect a GUI app (track 1), then ask it "create a Cogni API key" — it shows an sk-cogni-… key once . Drop your tool's config below and paste that key into the Authorization: Bearer header.
claude mcp add --transport http --scope user cogni \
https://mcp.getcogni.io/mcp \
--header "Authorization: Bearer sk-cogni-YOUR_KEY"
Cursor ⌄
Add to ~/.cursor/mcp.json , then Settings → Tools & MCP → enable cogni.
{
"mcpServers": {
"cogni": {
"url": "https://mcp.getcogni.io/mcp",
"headers": { "Authorization": "Bearer sk-cogni-YOUR_KEY" }
}
}
}
Cline ⌄
MCP Servers → Configure MCP Servers, add under mcpServers :
{
"mcpServers": {
"cogni": {
"type": "streamableHttp",
"url": "https://mcp.getcogni.io/mcp",
"headers": { "Authorization": "Bearer sk-cogni-YOUR_KEY" }
}
}
}
VS Code · needs Copilot ⌄
Create .vscode/mcp.json , then start the server and use it in Copilot Chat's Agent mode.
{
"servers": {
"cogni": {
"type": "http",
"url": "https://mcp.getcogni.io/mcp",
"headers": { "Authorization": "Bearer sk-cogni-YOUR_KEY" }
}
}
}
Windsurf ⌄
Add to ~/.codeium/windsurf/mcp_config.json (note the key is serverUrl ), then refresh in Cascade.
{
"mcpServers": {
"cogni": {
"serverUrl": "https://mcp.getcogni.io/mcp",
"headers": { "Authorization": "Bearer sk-cogni-YOUR_KEY" }
}
}
}
OpenClaw ⌄
One command (v2026.5.12+ for header support):
openclaw mcp add cogni \
--url https://mcp.getcogni.io/mcp \
--transport streamable-http \
--header "Authorization: Bearer sk-cogni-YOUR_KEY"
Gemini CLI ⌄
Add to ~/.gemini/settings.json (the key is httpUrl ):
{
"mcpServers": {
"cogni": {
"httpUrl": "https://mcp.getcogni.io/mcp",
"headers": { "Authorization": "Bearer sk-cogni-YOUR_KEY" }
}
}
}
Requires a paid Gemini API key or a Code Assist Standard/Enterprise license (per Google's June 2026 change).
Then load the Cogni skill. Connecting gives your assistant the tools; the skill teaches it when to use them, so it recalls before answering and remembers what matters without you asking each time. It is one small file. Download the skill · how to add it to your client .
One memory, every client. OAuth or key, GUI or terminal — every way you sign in lands on the same isolated store. Sign in

[truncated]
