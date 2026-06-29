---
source: "https://www.forthwrite.ai/blog/how-forthwrite-learns-your-email-voice"
hn_url: "https://news.ycombinator.com/item?id=48725623"
title: "Show HN: ForthWrite – Email AI that learns your voice from every edit you send"
article_title: "How ForthWrite Learns Your Email Voice: RAG, Edit-Distance Scoring, and a Phrasing Miner - ForthWrite"
author: "curtisboortz"
captured_at: "2026-06-29T22:23:22Z"
capture_tool: "hn-digest"
hn_id: 48725623
score: 2
comments: 0
posted_at: "2026-06-29T21:37:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ForthWrite – Email AI that learns your voice from every edit you send

- HN: [48725623](https://news.ycombinator.com/item?id=48725623)
- Source: [www.forthwrite.ai](https://www.forthwrite.ai/blog/how-forthwrite-learns-your-email-voice)
- Score: 2
- Comments: 0
- Posted: 2026-06-29T21:37:49Z

## Translation

タイトル: HN を表示: ForthWrite – 送信するすべての編集内容から音声を学習する電子メール AI
記事のタイトル: ForthWrite が電子メールの音声を学習する方法: RAG、編集距離スコアリング、およびフレージング マイナー - ForthWrite
説明: ForthWrite が電子メールの音声を学習するために使用するフィードバック ループの技術的なチュートリアル: MMR 再ランキング、編集距離スコアリング、および決定論的なフレージング マイナーを備えた pgvector RAG。

記事本文:
ForthWrite が電子メールの音声を学習する方法: RAG、編集距離スコアリング、およびフレージング マイナー - ForthWrite の機能 仕組み 料金 サインイン 無料で始める 無料で始める すべての記事 音声マッチング ForthWrite が電子メールの音声を学習する方法: RAG、編集距離スコアリング、およびフレージング マイナー
ForthWrite が電子メールの音声を学習するために使用するフィードバック ループの技術的なチュートリアルです。MMR 再ランキング、編集距離スコアリング、および決定論的なフレージング マイナーを備えた pgvector RAG です。
5 min read · June 24, 2026 私は約 1 年間、ある特定の問題を追いかけてきました。それは、電子メール アシスタントはあなたの電子メールではなく、専門的な電子メールを作成するということです。このツールは、明らかにユーザーが作成したものではなく、正確で洗練されたものを生成します。それを編集するのはあなたです。あなたがそれを送ります。次のものを編集するのはあなたです。何も改善されません。
ギャップは基礎となるモデルではありません。それがフィードバックループです。ほとんどのツールにはこれがありません。
ForthWrite がそれを閉じる方法は次のとおりです。
「自分の声を学習する」ことの核心的な問題
すべての AI メール ツールは、あなたの声を学習すると主張しています。クレームでは、ツールに応じて非常に異なる内容が説明されています。
プロンプト エンジニアリングとは、自分のスタイルを説明するシステム プロンプトを作成することを意味します。 「直接的な口調。全角ダッシュなし。100 語以内。」モデルは指示に従いますが、更新されません。どのドラフトでも同じ上限に達します。
微調整により、実際にはデータのモデルの重みが更新されます。これは機能する可能性がありますが、ユーザーごとにかなりのコンピューティングコストがかかり、執筆の進化に応じて動的に更新されません。消費者規模では実現不可能。
RAG (検索拡張生成) は ForthWrite が使用するものです。生成時にすべてのプロンプトに書き込みの実際の例を取り込むため、モデルは電子メールの説明ではなく、実際に送信された電子メールを模倣します。
RAGだけでは十分ではありません。検索は良好でなければなりませんが、「良い」検索は習慣が変わるにつれて変化します。 Th

at にはフィードバック メカニズムが必要です。
ForthWrite がコーパスを構築する方法
最初の接続時に、Gmail の送信済みフォルダーがバッチインポートされます。各メッセージは正規化され (HTML が削除され、転送されたセクションと署名が削除され)、 text-embedding-3-small で埋め込まれ、pgvector を使用して Postgres に保存されます。
これにより、AI 支援による 1 つの返信を送信する前に、モデルにあなたの声の何百もの実際の例が与えられます。ほとんどのツールでは、いくつかの書き込みサンプルを手動で貼り付けるように求められます。私たちはそれを自動的に、しかも大規模に実行します。
「ドラフト」をクリックすると、ForthWrite:
受信メールスレッドを埋め込みます
送信履歴からコサイン類似一致の候補プールを pgvector にクエリします。
MMR (Maximal Marginal Relevance) を使用して再ランク付けし、冗長なシグナルを追加するほぼ重複した例が返されることを回避します。
最終出力は、システム プロンプトに挿入された数ショットのサンプルの小さなブロックです。何かを注入する前に、最小信頼しきい値を強制します。バーをクリアする一致がない場合は、モデルが無関係な古い応答を模倣する危険を冒さずに、何も注入しません。
MMRステップは重要です。コサイン類似度だけでもクラスター化する傾向があります。上位 5 つの結果は、同じ電子メールのわずかなバリエーションであることがよくあります。 MMR は類似性と多様性をトレードオフするため、少数ショットのブロックは 1 つのスレッド タイプを過剰に表現するのではなく、スタイルの範囲をより多くカバーします。
編集距離のスコアリング: フィードバック ループ
AI ドラフトを編集して [送信] をクリックするたびに、ForthWrite は両方のバージョンをキャプチャします。
生成されたドラフトと実際に送信したものとの間の正規化されたレーベンシュタイン編集距離スコアを計算します。
1.0 に近いスコア: ほぼそのまま送信しました。ドラフトは正しかった。
0 に近いスコア: ほとんどを書き直しました。ドラフトは逃しました。
これらの (生成、送信された) ペアはテーブルに蓄積されます。スコアの低いペア (大量の編集) は、非同期オプティマイザーにフィードします。

ギャップを分析し、システムに即したアップデートを提案します。高スコアのペア (ほぼそのままの送信) は、ポジティブな少数ショット シグナルとして検索コーパスに入力されるため、将来の RAG プルは、すでに機能していたものに偏ります。
RAG は状況コンテキストを適切に処理します。しかし、テキストの 95% を保持し、2 つの単語だけを交換したニアミス下書きには、別のシグナルが隠れています。
「まったく心配ありません」と「心配ありません」という違いがある 2 つの電子メールのコサイン差分は同一に見えます。これらのペアは、「悪いドラフトを見つける」パスでは決して浮上しません。全体的なスコアが高いため (ドラフトの大部分は問題ありませんでした)、標準の編集距離バケット化でもそれらは見逃されます。
そこで、別のマイナーを実行します。これは、AI ドラフトと送信されたバージョンの間のワードレベルの LCS (最長共通部分列) の調整です。すべての削除と置換を収集し、コーパス全体でどれが繰り返されるかを数えて、一貫性を測定します。
AI が「丸返し」と書いた 12 回のうち、何回削除しましたか? 「お元気でしたら幸いです」と書かれた8回のうち、何回削除しましたか?頻度と一貫性のしきい値を超えるものは、明示的な回避または優先ルールとしてシステム プロンプトに書き込まれます。
このステップには LLM は関与しません。純粋な決定論的な差分。説明することは考えられないような、繰り返し現れる小さなパターンを捉えますが、常に手作業でクリーンアップします。
これにより、明白ではないダイナミクスが作成されます。モデルが改善されると、編集の回数が減り、生成される補正信号も少なくなります。オプティマイザは成功すると自らを枯渇させます。
編集距離スコアとは独立した時間ベースのリズムでフレージング マイニングを実行することで、これに対処します。安定した状態 (スコアが高く、編集がほとんどない) にあるユーザーは、最近送信された電子メールに対して定期的な絞り込みパスを受け取ります。フィードバックループは機能しているからといって崩壊するわけではありません。
これは広い意味で未解決の問題です。

もしあなたが RLHF やまばらな肯定的なフィードバックの好みモデリングで同様のことをやっているのであれば、どのようにアプローチしたのか知りたいと思います。
Supabase (Postgres + pgvector)
Anthropic Claude をプライマリ モデルとして使用し、コスト効率を高めるためにシステム ブロックに迅速にキャッシュします。
検索用の OpenAI text-embedding-3-small
プロンプト キャッシュの部分は、思っている以上に重要です。システム ブロック (ペルソナ、表現ルール、いくつかのショットの例) は、セッション内の複数のドラフトにわたって同じです。 Anthropic のキャッシュ TTL は、呼び出しごとではなく、1 時間に 1 回そのブロックの料金を支払うことを意味します。大規模な場合、これは手頃な価格の製品と、収益よりも運用コストの方が高い製品との違いになります。
ForthWrite は、Gmail の Chrome 拡張機能です。 14 日間の無料トライアルを開始するにはカードは必要ありません。
取得アーキテクチャ、フレージングマイナー、キャッシュ設計に関する質問に喜んでお答えします。
11 分読了 · 2026 年 6 月 16 日 AI メール音声マッチング: AI があなたのように書く方法を学習する方法
ほとんどの AI 電子メール ツールは、適切な電子メールを生成します。音声マッチング AI があなたに似たメールを生成します。この区別が実際に何を意味するのか、そしてそれが思っているよりも難しい理由を以下に示します。
6 分読了 · 2026 年 5 月 12 日 AI メール ツールについて「時間の節約」では分からないこと
スピードは、あらゆる AI 電子メール ツールにとって一か八かの約束です。実際に両者を分ける問題は、すべての単語を再読せずに送信できるほど出力が優れているかどうかです。
3 min read · Apr 29, 2026 微調整するために、最後に送信した 50 件のメールを ChatGPT にコピーしました。それでもうまくいきませんでした。
多くの人は、例を貼り付けることで ChatGPT に自分の電子メール スタイルを教えようとします。実際に試してみると何が起こるのか、そしてなぜすぐに天井に達してしまうのかを以下に示します。
他の人と同じように聞こえるのをやめる準備はできていますか？
5 歳未満であなたの声をキャプチャする一人称のペルソナ プロンプトを作成します

分。アカウントは必要ありません。
あなたが書いたものとまったく同じように聞こえる下書きを電子メールで送信します。品質を犠牲にすることなく時間を節約します。

## Original Extract

A technical walkthrough of the feedback loop ForthWrite uses to learn your email voice: pgvector RAG with MMR re-ranking, edit-distance scoring, and a deterministic phrasing miner.

How ForthWrite Learns Your Email Voice: RAG, Edit-Distance Scoring, and a Phrasing Miner - ForthWrite Features How it Works Pricing Sign in Get started free Get started free All articles Voice Matching How ForthWrite Learns Your Email Voice: RAG, Edit-Distance Scoring, and a Phrasing Miner
A technical walkthrough of the feedback loop ForthWrite uses to learn your email voice: pgvector RAG with MMR re-ranking, edit-distance scoring, and a deterministic phrasing miner.
5 min read · June 24, 2026 I have been chasing a specific problem for about a year: email assistants write a professional email, not your email. The tool produces something correct, polished, and obviously not from you. You edit it. You send it. You edit the next one. Nothing improves.
The gap is not the underlying model. It is the feedback loop. Most tools do not have one.
Here is how ForthWrite closes it.
The core problem with "learns your voice"
Every AI email tool claims to learn your voice. The claim describes very different things depending on the tool.
Prompt engineering means you write a system prompt describing your style. "Direct tone. No em-dashes. Under 100 words." The model follows instructions but does not update. You hit the same ceiling on every draft.
Fine-tuning actually updates model weights on your data. It could work, but it costs significant compute per user and does not update dynamically as your writing evolves. Not viable at consumer scale.
RAG (retrieval-augmented generation) is what ForthWrite uses: pull real examples of your writing into every prompt at generation time, so the model is imitating your actual sent email rather than a description of it.
RAG alone is not enough either. The retrieval has to be good, and "good" retrieval changes as your habits change. That requires a feedback mechanism.
How ForthWrite builds your corpus
On first connect, we batch-import your Gmail Sent folder. Each message is normalized (HTML stripped, forwarded sections and signatures removed), embedded with text-embedding-3-small , and stored in Postgres with pgvector.
This gives the model hundreds of real examples of your voice before you send a single AI-assisted reply. Most tools ask you to paste a few writing samples manually. We do it automatically, and we do it at scale.
When you hit Draft, ForthWrite:
Embeds the incoming email thread
Queries pgvector for a candidate pool of cosine-similar matches from your sent history
Re-ranks using MMR (Maximal Marginal Relevance) to avoid returning near-duplicate examples that would add redundant signal
The final output is a small block of few-shot examples injected into the system prompt. We enforce a minimum confidence threshold before injecting anything. If no match clears the bar, we inject nothing rather than risk the model imitating an irrelevant old reply.
The MMR step is important. Cosine similarity alone tends to cluster: the top five results are often slight variations of the same email. MMR trades off similarity against diversity, so the few-shot block covers more of your stylistic range rather than over-representing one thread type.
Edit-distance scoring: the feedback loop
Every time you edit an AI draft and hit Send, ForthWrite captures both versions.
We compute a normalized Levenshtein edit-distance score between the generated draft and what you actually sent:
Score near 1.0 : you sent it nearly verbatim. The draft was right.
Score near 0 : you rewrote most of it. The draft missed.
These (generated, sent) pairs accumulate in a table. Low-scoring pairs (heavy edits) feed an async optimizer that analyzes the gap and suggests system-prompt updates. High-scoring pairs (near-verbatim sends) go into the retrieval corpus as positive few-shot signals, so future RAG pulls lean toward what already worked.
RAG handles situational context well. But there is a different signal hiding in near-miss drafts, the ones where you kept 95% of the text and only swapped two words.
A cosine diff on two emails that differ by "no worries at all" versus "no worries" looks identical. Those pairs never surface in a "find bad drafts" pass. Standard edit-distance bucketing misses them too, because the overall score is high (most of the draft was fine).
So we run a separate miner: a word-level LCS (longest common subsequence) alignment between the AI draft and your sent version. We collect all deletions and substitutions, count which ones recur across your corpus, and measure consistency.
Of the 12 times the AI wrote "circle back," how many did you remove? Of the 8 times it wrote "I hope this finds you well," how many did you delete? Anything above a frequency and consistency threshold gets written to your system prompt as an explicit avoid or prefer rule.
No LLM involved in this step. Pure deterministic diff. It catches the small, recurring patterns that you would never think to describe but always clean up by hand.
This creates a non-obvious dynamic: as the model gets better, you make fewer edits, which produces fewer correction signals. The optimizer starves itself on success.
We handle it by running phrasing mining on a time-based cadence independent of edit-distance score. Users in a steady state (high scores, few edits) still get periodic refinement passes over recent sent email. The feedback loop does not collapse just because it is working.
It is an open problem in a broader sense. If you are doing something similar with RLHF or preference modeling on sparse positive feedback, I would be curious how you have approached it.
Supabase (Postgres + pgvector)
Anthropic Claude as primary model with prompt caching on the system block for cost efficiency
OpenAI text-embedding-3-small for retrieval
The prompt caching piece matters more than it sounds. The system block (your persona, phrasing rules, few-shot examples) is the same across multiple drafts in a session. Anthropic's cache TTL means we pay for that block once per hour rather than on every call. At scale, this is the difference between an affordable product and one that costs more to run than it earns.
ForthWrite is a Chrome extension for Gmail. 14-day free trial, no card required to start.
Happy to answer questions about the retrieval architecture, the phrasing miner, or the caching design.
11 min read · Jun 16, 2026 AI Email Voice Matching: How AI Learns to Write Like You
Most AI email tools generate competent email. Voice-matching AI generates email that sounds like you. Here is what that distinction means in practice and why it is harder than it sounds.
6 min read · May 12, 2026 What 'Saves Time' Doesn't Tell You About an AI Email Tool
Speed is the table-stakes promise of every AI email tool. The question that actually separates them is whether the output is good enough that you'd send it without re-reading every word.
3 min read · Apr 29, 2026 I Copied My Last 50 Sent Emails Into ChatGPT to Fine-Tune It. It Still Didn't Work.
A lot of people try to teach ChatGPT their email style by pasting examples. Here's what actually happens when you try it, and why it hits a ceiling fast.
Ready to stop sounding like everyone else?
Build a first-person persona prompt that captures your voice in under 5 minutes. No account required.
Email drafts that sound exactly like you wrote them. Save time without sacrificing quality.
