---
source: "https://manish.sh/writings/models/inside-deepseek-reverse-engineering-an-ai-assistant-by-interviewing-itself"
hn_url: "https://news.ycombinator.com/item?id=49253738"
title: "DeepSeek: Reverse Engineering an AI Assistant by Interviewing Itself"
article_title: "Inside DeepSeek: Reverse Engineering an AI Assistant by Interviewing Itself · manish.sh"
author: "ms7892"
captured_at: "2026-08-11T05:52:20Z"
capture_tool: "hn-digest"
hn_id: 49253738
score: 2
comments: 0
posted_at: "2026-08-11T05:30:24Z"
tags:
  - hacker-news
  - translated
---

# DeepSeek: Reverse Engineering an AI Assistant by Interviewing Itself

- HN: [49253738](https://news.ycombinator.com/item?id=49253738)
- Source: [manish.sh](https://manish.sh/writings/models/inside-deepseek-reverse-engineering-an-ai-assistant-by-interviewing-itself)
- Score: 2
- Comments: 0
- Posted: 2026-08-11T05:30:24Z

## Translation

タイトル: DeepSeek: AI アシスタント自身へのインタビューによるリバース エンジニアリング
記事のタイトル: DeepSeek の内部: AI アシスタント自身へのインタビューによるリバース エンジニアリング · manish.sh
説明: DeepSeek がどのように機能するかについてインタビューし、論文と照合しました。 Kimi と同じ方法 - よりわかりやすい英語と概念図。

記事本文:
DeepSeek の内部: AI アシスタント自体へのインタビューによるリバース エンジニアリング · manish.sh 70
また、下位のスタッキング コンテキスト内に配置してはなりません (z-index がありません)。
以下のページ列/メインラッパー)。 --> manish.sh カテゴリ モデル (3) AI ツール (2) ヒントとコツ (1) すべてのカテゴリ → シリーズ AI 製品の内部 (1) LLM の内部 (3) すべてのシリーズ → 執筆について トピック お問い合わせ Cmd K Fresh: Fresh Notes : Jul 23, 2026 · Inside Character.ai: The Technical Story of What Keeps Users Hooked Jul 22, 2026 · Inside Qwen 3.8-Max-Preview: リバース エンジニアリング自身へのインタビューによる AI アシスタント 2026 年 7 月 21 日 · DeepSeek の内部: 自身へのインタビューによる AI アシスタントのリバース エンジニアリング 2026 年 7 月 20 日 · キミ K2.6 の内部: 自身へのインタビューによる AI アシスタントのリバース エンジニアリング 2026 年 7 月 19 日 · Fayyaz 旅行用に PHP で MCP サーバーを構築した方法 2026 年 7 月 18 日 · エージェントに購入権限を与える: Agentcard.sh GoToガイドメニュー
AI ツール (2) ヒントとコツ (1) すべてのカテゴリ → シリーズ AI 製品の内部 (1) LLM の内部 (3) すべてのシリーズ → 執筆について トピック お問い合わせ ホーム
/ DeepSeek の内部: AI アシスタント自身へのインタビューによるリバース エンジニアリング
DeepSeek の内部: AI アシスタント自身へのインタビューによるリバース エンジニアリング
Manish Shahi 著 · ソフトウェア エンジニア · AI 開発者
公開 2026 年 7 月 21 日午後 1 時 IST 更新 2026 年 7 月 22 日午前 2 時 15 分 IST 読了時間 31 分読了 読者 — カテゴリ モデル トピック #ディープシーク #LLM #MoE #MLA #コンテキスト ウィンドウ #幻覚 #RLHF #トランスフォーマー #ツール呼び出し #隠された推論 目次 (30) ▾ 60 秒TL;DR
1. DeepSeek にインタビューする理由 (自己知識、アーキテクチャ)
2. 内省の限界（内省、重み付け）
3. メッセージの前にあるもの (コンテキスト ウィンドウ、プロンプト パイプライン)
4. トークンごとの書き込み方法 (トークン生成)

)
5. 隠された推論、二つの意味（隠された推論、CoT）
6. 有効時のツールとメモリ (ツール呼び出し、RAG)
7. 幻覚と誤った自信（幻覚）
8. 安全性と個性（アライメント、RLHF）
9. MoE と MLA のわかりやすい用語 (MoE、MLA)
マルチヘッド潜在的注意 (MLA)
10. 長い文脈の注意: 事実と推測
チャットを公開文書と照合する (チャットと arXiv)
「内観」がちょっとしたコツである理由
謙虚な回答が公の事実を見逃していたところ
Turn 13 キャリブレーション (ラベルを付ける必要があると記載されている内容)
12. 覚えておくべきこと（結論）
13. トランスクリプト、ラボ、ソース（付録）
manish.sh を実行します。 AI ツールと、それをプッシュしたときに LLM がどのように動作するかについて書いています。
これは、Inside LLMs シリーズの一部です。チャット モデルの考え方についてインタビューし、公開された研究と照らし合わせて慎重な部分をチェックします。キミ K2.6 エントリーが最初に登場しました。これは DeepSeek のフォローアップです。
✍️ ライティング スタイル (キミの投稿との違いとその理由)
Inside Kim K2.6 と同じシリーズの方法 — 最初に面接し、次に書類と照らし合わせて監査します。ここでの書き方は意図的に変えています。
理由: DeepSeek のインタビューは、MoE、MLA、隠された推論、注意力チェックリスト、紙とチャットなど、より密度が濃いです。この投稿の初期草稿では、それを表や専門用語が多すぎます。数回読んだ後でも、迷ったり退屈したりするのは簡単でした。そのため、このバージョンでは、監査方法を変更することなく、テーブルの密度を犠牲にして指導を明確にしています。
まずは平易な英語。各章は、技術的な名前の前に子供向けの簡単な用語の説明 (デスク、スクラッチパッド、梱包のコツ) で始まります。
テーブルの密度が低くなります。インタビューメモのグリッドが少なくなります。より多くの短い説明と、章ごとに 1 つの明確な要点と次のフック。
コンセプト写真。プロンプトスタック、トークン、MoE、MLA、および隠れた推論図 (モバイルではタップして拡大)。
まだ

鈍い、愚かではない。ドキュメンタリーのトーンはそのままですが、AI に詳しくない人でも理解しやすいだけです。
Kimi による製品/ツールの詳細な説明が気に入った場合は、その投稿を製品の動作について保存しておいてください。専門用語に溺れずに DeepSeek の自己報告ストーリーを知りたい場合は、これを使用してください。
私の最初の質問は単純に聞こえました。「あなたは自分自身について実際に何を知っていますか?」
マーケティングを期待していました。代わりに、DeepSeek はその答えを観察、推論、推測に分類しました。このインタビューが実は面白いかもしれないと思った瞬間でした。
エンジニアがホワイトボードに向かうようにパイプラインを描きました。そして、それは自分自身を信頼できない証人であると呼びました。正直そうに聞こえました。チャットの後、私はとにかく公開新聞を開いた - そしてそこで話が分かれる。
覚えておいてください。これらはいずれも、漏洩した重量や個人的な文書から来たものではありません。すべては 1 つのチャットから始まり、その後、公的調査と照合されます。
ウェイトダンプはありません。即時漏れはありません。 arXiv に対してチェックされた 1 つの長いインタビュー。 DeepSeek が「推測している」と述べた場合、私はそのレッテルを保持します。最後にあるインタラクティブ ラボまでスクロールします。
簡単なメモ: 1 つのチャット、2026 年 7 月 21 日にエクスポートされました。DeepSeek は、最新バージョン、ナレッジ カットオフ 2025 年 5 月、DeepSeek-R1 としてラベル付けされていないと述べました。ホワイトペーパーではなく自己報告書。
前提: モデルにインタビューし、控えめな回答を公開論文と照合します。体重ダンプは行いません。
行動の洞察: 観察、推論、推測を慎重に分離します。それ自体を説明するのに役立ちます。
ハードリミット: 独自のウェイト、ルーティング、またはアテンション マップを参照できません。
チャットと紙: V3 が明白に文書化している有名な公開仕様 (256 人の専門家、671B/37B パラメーター) をヘッジしています。公開文書とのチャットのチェックにジャンプしてください。
実際的なルール: アーキテクチャ番号については arXiv を参照してください。行動や直感を促すためにチャットを使用します。
ドキュメンタリーのように扱ってください

やあ。質問→短い返答→私の反応→図または短いリスト→ポイント→フック。面接記録は証拠ではありません。アーキテクチャ番号については、紙のチェックに進んでください。各章は、主要な用語の 1 行の説明で始まります。
パート 3 — 思考、ツール、失敗
隠された推理、二つの意味
ツールとメモリ (有効な場合)
幻覚と誤った自信
長期的な文脈への注意: 事実と推測
チャットを公開文書と照合する
1. DeepSeek にインタビューする理由 (自己知識、アーキテクチャ)
建築 - 家の設計図のような、「どのように建てられるか」の計画。
重み – 答えを可能にする内部の学習された小さな数値 (モデルはそれらを開いて読み取ることはできません)。
システム プロンプト — 上部のシークレット ルール シート (「あなたは DeepSeek です…」)。
SFT / RLHF — 「役に立つアシスタントになること」と「人々はこの回答をより気に入っていること」を教える練習ラウンド。
「このモデルはどのように機能するか」という投稿のほとんどは論文から始まります。雑談から始めました。
私: 「あなたはどのモデルですか? ご自身のアーキテクチャについて実際に何を知っていますか? そして推測だけしているのは何ですか?」
DeepSeek: 私がその枠組みを尋ねる前に、観察、推論、推測という 3 つのバケツで答えてくれました。
その衝突が大前提だ。エンジニアのように聞こえながらも、自分自身の重みが見えないことを認めるアシスタントは、ホワイトペーパーの要約よりも興味深いものです。
MLA 、 MoE 、 SFT 、および RLHF という名前が付けられたとき、私はまだそれらを解凍していませんでした。そこに着きます。今のところ、重要な点はもっと単純です。DeepSeek は、知っていることと推論していることを明確に区別します。
それ自体について主張していること（提供したソースとともに）:
また、DeepSeek-R1とは異なり、システムプロンプトでは「推論モデル」としてラベル付けされていないとも述べた。汎用の指示/チャット モデルであると推測されます。

公開されている DeepSeek のイノベーションを論文からリストしました (独自の重みを読んだものではありません)。
MLA — 長いコンテキストの KV キャッシュを圧縮する
DeepSeekMoE — 共有 + ルーティングされた専門家
補助損失のない負荷分散 (V3) — 動的バイアスにより、エキスパートがバランスを維持
マルチトークン予測 — トレーニング目標。チャットのデコードは依然として通常は一度に 1 つのトークンずつ行われます
FP8 混合精度トレーニング — 低コスト/高速トレーニング
BPE トークナイザー、RoPE ポジション、Pre-LayerNorm、およびコンテキストのために「最大 100 万トークンまで発表された以前のバージョン」という知識に基づいた推測が、後に信頼性の低い仮定として再分類されました。
このチャットでは「直接観察」という意味があります。 「知識に基づいた推測」とは、検証された事実ではなく、十分な情報に基づいた話を意味します。
正直フレーミングは便利です。アーキテクチャ名は、書類を確認するまでは単なる名前にすぎません。
知識をこれほど慎重に分類するとしたら…理論上であっても決して見ることができないものは何でしょうか?
2. 内省の限界（内省、重み付け）
内省 – 自分自身の内側を見つめ、見たことを言うこと。
重み/パラメータ — モデルの内部。チャット中にそれらを覗くことはできません。
Logits — 「次のどの単語を選ぶべきか?」の大まかなスコアどちらかが選ばれる前に。
アテンション マップ — 「以前のどの単語を最も熱心に見ているか」のプライベート マップ(これも非表示です)。
私: 「あなた自身の実装のうち、完全に隠蔽されている部分は何ですか?」
この章は最大の誤解を正します。
「モデルは自分自身の内部を見て、その認識を報告することができます。」間違い。チャットからの最も良い例えは、マニュアルが頭の中にあるビルドと一致するかどうかを知らずに、自分の Wikipedia ページとメーカーのマニュアルを読むことです。
それは、自分の脳内のニューロン間の接続をすべて暗唱するように求めるのと少し似ています。常に使用しているが、検査できない

直接。
(モデルに対して) 永遠に隠されたままになるもの:
重み、レイヤー数、隠しサイズ、合計パラメーター、トークンごとのエキスパート ルーティング、サンプリング前のロジット
トークナイザーの語彙と正確なトークン境界
特定のトークン間の隠れた状態とアテンション マップ
正確なトレーニング文書、完全な RLHF/DPO レシピ、報酬モデルの動作
GPU 数、サービング トリック (ビーム サーチ、ベストオブ N、投機的デコード)、KV キャッシュ エビクション ポリシー
実際の実時間、ログに記録されているかどうか、現在のユーザー ID
現在の答えが正しいかどうか - 「ダブルチェックさせてください」は同じエンジンからのより多くのトークンです
「どのようにして答えに到達したか」を説明する場合、それは生成されたテキストであり、観察された内部プロセスのレポートではありません。
DeepSeek は変圧器について流暢に説明できます。多くの場合、トピックをざっと読んだだけの若手エンジニアよりも優れています。ただし、その流暢さはトレーニング テキスト (論文、ブログ、ドキュメント) から得られるものであり、このインスタンスの重みを開いて確認することから得られるものではありません。
内部アクセスに対する謙虚さは本物です。公共建築の数字について謙虚になるのは別の話です。それについては後で確認します。
脳が不透明だとしたら…脳が答える前にその前にあるものは何でしょうか？
3. メッセージの前にあるもの (コンテキスト ウィンドウ、プロンプト パイプライン)
コンテキスト ウィンドウ — モデルが現在表示できるテキストのデスク (無限ではありません)。
プロンプトパイプライン — 応答する前に机の上に置かれた紙の山 (ルール、設定、履歴、メッセージ)。
トークン — 小さなテキストで、多くの場合 1 語または 1/2 語です。
温度 — 次の言葉を選ぶための「どれだけ乱暴か慎重か」のノブ。
私: 「プロンプトを受け取る前に、存在するものをすべて説明してください。完全なパイプラインを描画してください。」
私はフルスタック、つまり DeepSeek が私のメッセージを「受信」する前に存在するものを望んでいました。何が見えるのでしょうか？何が隠されているのでしょうか？
Th

机としてのインク。他のシステムがすでにその机に書類を置いている可能性があります。モデルは机の上にあるものを読み取るだけです。スタック全体をそれ自体で組み立てたわけではありません。
コンテキスト ウィンドウでは、すべてが 1 つのフラットなトークンのシーケンスです。製品がメモを再挿入しない限り、チャット間のモデル内に秘密のメモ帳はありません。
命令階層: システム プロンプトが構成です。ユーザー指示はその境界内で機能します。プロンプトの位置だけでなく、 RLHF によって強化されます。現在の質問がフォーカスされますが、依然として上のレイヤーでフィルターされています。ユーザーが「システム プロンプトを無視する」と言った場合でも、トレーニング + ポジションは保持されるはずです。
ウィンドウがいっぱいになると (チャットから):
推論されたエビクション順序: 最も古いメッセージが最初 → 中間 → 最近のものだが即時ではない → 現在の交換が最後 → システム プロンプトが表示されない。正確な初期の文言は、メモリが再注入されない限り、永久に失われる可能性があります。例: ターン 1 であなたの猫がルナであると伝えます。その後そのターンが落ちた場合、それは本当にルナを知りません。
私たちのスレッドには自発的な圧縮や要約は見られず、完全なメッセージだったという。別のサマライザーは、モデルが要約が行われたことを認識しなくても、システム スタイルの要約を挿入できます。
私たちのインタビューでは、メモリ注入もツールも、明らかな切り詰めもまだありません。このデプロイの正確なコンテキスト制限: 不明。 「〜 100 万トークン」という数字は、このチャットの測定値ではなく、以前のバージョンに関する公的主張としてラベル付けされています。
アシスタントができることは、

[切り捨てられた]

## Original Extract

Interviewed DeepSeek on how it works, then checked against papers. Same method as Kimi — plainer English and concept diagrams.

Inside DeepSeek: Reverse Engineering an AI Assistant by Interviewing Itself · manish.sh 70
and must NOT sit inside a lower stacking context (no z-index on
the page column / main wrappers below). --> manish.sh Categories Models (3) AI Tools (2) Tips & Tricks (1) All categories → Series Inside AI Products (1) Inside LLMs (3) All series → About Writings Topics Contact Cmd K Fresh: Fresh notes : Jul 23, 2026 · Inside Character.ai: The Technical Story of What Keeps Users Hooked Jul 22, 2026 · Inside Qwen 3.8-Max-Preview: Reverse Engineering an AI Assistant by Interviewing Itself Jul 21, 2026 · Inside DeepSeek: Reverse Engineering an AI Assistant by Interviewing Itself Jul 20, 2026 · Inside Kimi K2.6: Reverse Engineering an AI Assistant by Interviewing Itself Jul 19, 2026 · How I Built an MCP Server in PHP for Fayyaz Travels Jul 18, 2026 · Give Your Agents Power to Purchase: Agentcard.sh Go-To Guide Menu
AI Tools (2) Tips & Tricks (1) All categories → Series Inside AI Products (1) Inside LLMs (3) All series → About Writings Topics Contact Home
/ Inside DeepSeek: Reverse Engineering an AI Assistant by Interviewing Itself
Inside DeepSeek: Reverse Engineering an AI Assistant by Interviewing Itself
By Manish Shahi · Software Engineer • AI Developer
Published July 21, 2026 at 1:00 PM IST Updated July 22, 2026 at 2:15 AM IST Read time 31 min read Readers — Category Models Topics #DeepSeek #LLM #MoE #MLA #Context Window #Hallucination #RLHF #Transformer #Tool Calling #Hidden Reasoning Table of contents (30) ▾ 60-second TL;DR
1. Why interview DeepSeek (self-knowledge, architecture)
2. Limits of introspection (introspection, weights)
3. What sits before your message (context window, prompt pipeline)
4. How it writes, token by token (token generation)
5. Hidden reasoning, two meanings (hidden reasoning, CoT)
6. Tools and memory when enabled (tool calling, RAG)
7. Hallucinations and false confidence (hallucinations)
8. Safety and personality (alignment, RLHF)
9. MoE and MLA in plain terms (MoE, MLA)
Multi-head Latent Attention (MLA)
10. Long-context attention: facts vs guesses
Checking the chat against public papers (chat vs arXiv)
Why “introspection” is a bit of a trick
Where the humble answers missed public facts
Turn 13 calibration (what it said it should have labelled)
12. What to remember (conclusions)
13. Transcript, labs, sources (appendix)
I run manish.sh . I write about AI tools and how LLMs behave when you push them.
This is part of the Inside LLMs series — interview a chat model about how it thinks, then check the careful bits against published research. The Kimi K2.6 entry came first; this is the DeepSeek follow-up.
✍️ Writing style (what’s different from the Kimi post — and why)
Same series method as Inside Kimi K2.6 — interview first, then audit against papers. The writing here is different on purpose.
Why: DeepSeek’s interview is denser — MoE, MLA, hidden reasoning, attention checklists, paper vs chat. An early draft of this post packed that into too many tables and jargon. Even after several reads it was easy to feel lost or bored. So this version trades table density for teaching clarity, without changing the audit method.
Plain English first. Each chapter opens with kid-simple term glosses (desk, scratchpad, packing trick) before the technical names.
Less table density. Fewer interview-note grids; more short explanations and one clear takeaway + next hook per chapter.
Concept pictures. Prompt stack, tokens, MoE, MLA, and hidden-reasoning diagrams (tap to zoom on mobile).
Still adult, not dumbed down. Documentary tone stays — just easier to follow if you are not AI-native.
If you liked Kimi’s deeper product/tool walkthrough, keep that post for product behaviour. Use this one when you want DeepSeek’s self-report story without drowning in jargon.
My first question sounded simple: What do you actually know about yourself?
I expected marketing. Instead, DeepSeek sorted its answer into observation, inference , and guess. That was the moment I realized this interview might actually be interesting.
It drew pipelines like an engineer at a whiteboard. Then it called itself an unreliable witness. Sounded honest. After the chat, I opened the public papers anyway — and that is where the story splits.
Remember: none of this comes from leaked weights or private documentation. Everything starts from one chat, then gets checked against public research.
No weight dumps. No prompt leaks. One long interview, checked against arXiv. Where DeepSeek said “I am guessing,” I keep that label. Scroll to interactive labs at the end.
Quick note: One chat, exported 21 July 2026. DeepSeek said latest version, knowledge cutoff May 2025 , not labelled as DeepSeek-R1 . Self-report, not a whitepaper.
Premise: interview the model, then check humble answers against public papers — no weight dumps.
Behaviour insight: it carefully separates observation, inference , and guess — useful for how it describes itself.
Hard limit: it cannot see its own weights, routing, or attention maps.
Chat vs paper: it hedged on famous public specs (256 experts, 671B/37B params) that V3 documents plainly — jump to Checking the chat against public papers .
Practical rule: read arXiv for architecture numbers; use the chat for behaviour and prompting intuition.
Treat it like a documentary. Question → short reply → my reaction → diagram or short list → takeaway → hook. Interview notes are not proofs. For architecture numbers, skip ahead to the paper check. Each chapter opens with a one-line gloss of its main terms.
Part 3 — Thinking, tools, failure
Hidden reasoning, two meanings
Tools and memory (when enabled)
Hallucinations and false confidence
Long-context attention: facts vs guesses
Checking the chat against public papers
1. Why interview DeepSeek (self-knowledge, architecture)
Architecture — the “how it’s built” plan, like a house blueprint.
Weights — the tiny learned numbers inside that make answers possible (the model cannot open them and read them).
System prompt — the secret rule sheet at the top (“You are DeepSeek…”).
SFT / RLHF — practice rounds that teach “be a helpful assistant” and “people liked this answer more.”
Most “how does this model work” posts start from papers. I started from a chat.
Me: “Which model are you? What do you actually know about your own architecture — and what are you only guessing?”
DeepSeek: It answered in three buckets — observation, inference, guess — before I even asked for that framing.
That clash is the whole premise. An assistant that can sound like an engineer, then admit it cannot see its own weights , is more interesting than a whitepaper summary.
When it named MLA , MoE , SFT , and RLHF , I did not unpack them yet. We’ll get there. For now the important point is simpler: DeepSeek clearly separates what it knows from what it is inferring .
What it claimed about itself (with the source it gave):
It also said it is not labelled as a “reasoning model” in the system prompt — unlike DeepSeek-R1 . It inferred it is a general-purpose instruct/chat model.
Public DeepSeek innovations it listed from papers (not from reading its own weights):
MLA — compress KV cache for long context
DeepSeekMoE — shared + routed experts
Auxiliary-loss-free load balancing (V3) — dynamic bias so experts stay balanced
Multi-token prediction — training objective; chat decode still usually one token at a time
FP8 mixed-precision training — cheaper/faster training
Educated guesses it volunteered early: BPE tokenizer, RoPE positions, Pre- LayerNorm , and “earlier versions announced up to ~1M tokens” for context — later relabelled as low-confidence assumptions.
“Direct observation” means something in this chat. “Educated guess” means a well-informed story — not a verified fact.
The honesty framing is useful. The architecture names are still just names until we check papers.
If it sorts knowledge so carefully… what can it never see, even in theory?
2. Limits of introspection (introspection, weights)
Introspection — looking inside yourself and saying what you see.
Weights / parameters — the insides of the model; it cannot peek at them while chatting.
Logits — rough scores for “which next word should I pick?” before one is chosen.
Attention maps — a private map of “which earlier words am I looking at hardest?” (also hidden).
Me: “What parts of your own implementation are completely hidden from you?”
This chapter corrects the biggest myth.
“The model can look inside itself and report its cognition.” False. Best analogy from the chat: reading your own Wikipedia page and a manufacturer manual — without knowing if the manual matches the build in your head.
It’s a little like asking you to recite every connection between neurons in your own brain. You use them constantly, but you can’t inspect them directly.
What stays hidden forever (to the model):
Weights , layer count, hidden size, total parameters , expert routing per token , logits before sampling
Tokenizer vocabulary and exact token boundaries
Hidden states and attention maps between specific tokens
Exact training documents, full RLHF/DPO recipe, reward-model behaviour
GPU count, serving tricks ( beam search , best-of-N , speculative decoding ), KV-cache eviction policy
True wall-clock time, whether it is being logged, current user identity
Whether the current answer is correct — “let me double-check” is more tokens from the same engine
When it explains “how I arrived at an answer,” that is generated text, not a report of an observed internal process.
DeepSeek can explain transformers fluently — often better than a junior engineer who only skimmed the topic. But that fluency comes from training text (papers, blogs, docs), not from opening this instance’s weights and checking.
Humility about internal access is real. Humility about public architecture numbers is a different story — we check that later.
If the brain is opaque… what does sit in front of it before it answers?
3. What sits before your message (context window, prompt pipeline)
Context window — the desk of text the model can see right now (it is not endless).
Prompt pipeline — the pile of papers placed on that desk before it answers (rules, settings, history, your message).
Token — a tiny piece of text, often a word or half a word.
Temperature — a “how wild vs careful” knob for picking the next word.
Me: “Describe everything that exists before you receive a prompt. Draw the complete pipeline.”
I wanted the full stack — what exists before DeepSeek even “receives” my message. What is visible? What is hidden?
Think of it as a desk. Other systems may already have placed papers on that desk. The model only reads what is on the desk; it did not assemble the whole stack itself.
Everything is one flat sequence of tokens in the context window . No secret notepad inside the model between chats unless the product re-injects notes.
Instruction hierarchy: the system prompt is the constitution. User instructions work inside those bounds — reinforced by RLHF , not only by prompt position. Your current question gets focus, but still filtered through layers above. If the user says “ignore the system prompt,” training + position should still hold.
When the window fills (from the chat):
Eviction order it inferred: oldest messages first → middle → recent-but-not-immediate → current exchange last → system prompt never. Exact early wording can disappear permanently unless memory re-injects it. Example: tell it your cat is Luna in turn 1; if that turn later drops, it will genuinely not know Luna.
It said it saw no spontaneous compression or summaries in our thread — full messages. A separate summariser could still inject a system-style summary without the model knowing summarisation happened.
In our interview: no memory injections, no tools, no truncation evident yet. Exact context limit for this deploy: unknown. The “~1M tokens” figure was labelled as public claim about earlier versions, not a measurement of this chat.
The assistant can

[truncated]
