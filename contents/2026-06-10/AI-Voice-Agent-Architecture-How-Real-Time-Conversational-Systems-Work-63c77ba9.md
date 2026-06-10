---
source: "https://www.faridfadaie.com/2026/06/10/ai-voice-agent-architecture/"
hn_url: "https://news.ycombinator.com/item?id=48479045"
title: "AI Voice Agent Architecture: How Real-Time Conversational Systems Work"
article_title: "AI Voice Agent Architecture: Lessons From 3 Production Builds"
author: "ffadaie"
captured_at: "2026-06-10T19:02:23Z"
capture_tool: "hn-digest"
hn_id: 48479045
score: 1
comments: 0
posted_at: "2026-06-10T16:46:44Z"
tags:
  - hacker-news
  - translated
---

# AI Voice Agent Architecture: How Real-Time Conversational Systems Work

- HN: [48479045](https://news.ycombinator.com/item?id=48479045)
- Source: [www.faridfadaie.com](https://www.faridfadaie.com/2026/06/10/ai-voice-agent-architecture/)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T16:46:44Z

## Translation

タイトル: AI 音声エージェント アーキテクチャ: リアルタイム会話システムの仕組み
記事のタイトル: AI 音声エージェント アーキテクチャ: 3 つの本番ビルドからの教訓
説明: 同じ AI 音声エージェントを 3 回構築しました。サーバー側のオーケストレーターが崩壊する理由、サーバーゲート型ターンがデッドエアを生み出す理由、そして機能するアーキテクチャ。

記事本文:
AI 音声エージェント アーキテクチャ: 3 つの本番ビルドからの教訓
コンテンツにスキップ
ファリド・ファダイエ
ブログ
ホーム >
AI 音声エージェントのアーキテクチャ: 同じエージェントを 3 回構築して学んだこと
AI 音声エージェントのアーキテクチャ: 同じエージェントを 3 回構築して学んだこと
投稿カテゴリー: AI / エンジニアリング
同じプロダクション音声エージェントを 3 回構築しました。同じ要件、同じ電話スタック、同じ音声モデルが利用可能 — 3 つの根本的に異なるアーキテクチャ。最初のものはそれ自身の結合の下で崩壊しました。すべてのバグ修正は他の何かを壊しました。 2 番目の方法は制御可能で正確でしたが、人間のように即座に応答するという音声エージェントの唯一の目的を破壊してしまいました。 3 番目は、実際の発信者との接触を生き残った唯一のものです。
この投稿は、何か月も実費も費やして苦労する前に読んでおきたかったアーキテクチャに関する記事です。すべてを決定するのは、次のターンを誰が所有するかという 1 つの質問です。
エージェントが実際の電話に応答します。一度に 3 つのことを実行する必要があり、それらは異なる方向に引っ張られます。
自然な会話をしましょう。 1 秒未満の順番交代、無表情な空気、割り込み、とりとめのない発信者、会話の途中で言語を切り替える発信者への優雅な対応。
構造化データを確実にキャプチャします。定義されたフィールドのセットは、最終的にデータベースに保存される必要があります。名前のスペルは正しく、数字は桁が完璧で、何も発明されていません。
厳しいルールを強制します。いくつかの行動は常に発生しなければなりません (電話番号の後ろの桁を一桁ずつ読み取る) が、絶対に発生してはいけない行動もいくつかあります (必要な質問に答えられないまま電話を切る、企業が許可していないことを約束するなど)。
以下の各アーキテクチャは、これら 3 つの責任がどこに存在するかに対する異なる答えです。
アーキテクチャ 1: サーバー側オーケストレーター
アーキテクチャ 1: 前夜

ry 呼び出し元はファンを並列分類子に向けます。共有フラグはすべてをすべてに結合します。
最初のビルドでは、音声モデルをマウスピースとして扱いました。サーバーがすべてを所有していました。各発信者の発話は文字に起こされ、一連の並列分類子に展開されました。1 つは通話カテゴリの決定、1 つはプロバイダー名を監視、1 つは通話終了の意図を監視、1 つは参照のみの質問の検出、さらに言語検出用に手動で管理された正規表現バンク (6 個のスクリプトで「X と話せますか?」の音声文字変換を含む) を管理しました。マージ レイヤーはその出力を「ターン分析」に結合し、ポリシー エンジンは次に何を言うかを、多くの場合スクリプト化されたプロンプトから選択します。
すべての決定は検査可能なサーバー コードです。何か問題が発生した場合は、行番号が表示されます。
各分類子は小さく、安価で、単独で単体テストが可能です。
ルールの追加は簡単そうに見えます。別の分類子を作成し、それをマージに接続します。
スクリプト化されたプロンプトを使用すると、エージェントの発言を単語レベルで制御できます。
分類器は実際には決して独立していませんでした。マージ レイヤーと共有セッション フラグの成長の網は、すべての懸念のガードが他のすべての懸念の状態を読み取ることを意味しました。プロバイダー ゲートは通話終了の状態を知る必要がありました。通話終了ロジックは、言語の切り替えが保留中かどうかを知る必要がありました。言語ロジックは、スクリプト化されたプロンプトが再生中かどうかを知るために必要でした。 1 つのガードのバグを修正すると、別のガードが依存していた不変条件が確実に破られました。これは「絡み合いすぎた」失敗です。システムにはいかなる決定に対しても単一の所有者がいなかったため、すべての決定は、余白で一致しない N 個のコードによって N 回行われました。
正規銀行は多言語の現実に耐えられません。言語検出パターンは音訳とエッジケースの博物館に成長しました

— 反証不可能でレビュー不可能であり、ライブ通話でのみ発見できる方法で間違っています。新しい言語はすべて、あらゆるパターンを再派生することを意味していました。私がルーティングしていたモデルは、これらの言語をすべてネイティブに理解していました。
レイテンシーが積み重なっています。ターンごとに複数の分類子をラウンドトリップし、次にポリシー決定を行い、次に音声合成を行います。どの作品も速かったです。パイプラインはそうではありませんでした。
ターン中盤の意見の相違。 2 つの分類器が同時に起動されたとき (呼び出し元がプロバイダーを指定し、同時に別れを告げたとき)、2 つのサブシステムは両方とも次のターンを自分たちが所有していると信じていました。聞こえるバグのほとんど (二重のプロンプト、発信者に話しかける、矛盾したフォローアップ) はまさにこれでした。
アーキテクチャ 1 からの深い教訓: 単一の所有者を持たない分散型意思決定はバグです。各コンポーネントがどれだけきれいであるかは問題ではありません。 N 個のコンポーネントがそれぞれ音声を開始できる場合、競合状態が人々の耳に伝わります。
アーキテクチャ 2: サーバーゲート型ターンテイキング
アーキテクチャ 2: 決定者は 1 人、完全な制御、そしてターンごとにサーバーによるデッドエアの往復。
「決定者が多すぎる」場合の明らかな解決策は、決定者を 1 人にすることです。 2 番目のビルドでは、サーバーがすべてのターンの単一の所有者になりました。リアルタイム音声モデルはデフォルトでミュートされ、サーバーが指示付きの応答を明示的にトリガーした場合にのみ発話されました。発信者が話す → トランスクリプト → サーバーが分類 → サーバーが何を言うべきかを決定 → サーバーがモデルにそれを言うように指示します。
最大限のコントロール。エージェントは文字通り台本なしで話すことができませんでした。 「決してXとは言わない」が強制力を持つようになった。コンプライアンスのレビューが簡単になりました。
おなじみのメンタルモデル。リクエストの入力、決定、応答の出力 - これは音声が添付された Web アプリです。推論しやすく、記録しやすく、部分ごとにテストするのも簡単です。
毎ターンに 1 人の所有者。アーキテクチャ 1 の競合状態はほとんどが問題です

が現れた。
毎ターンデッドエア。発信者は話すのをやめます。サーバーは分類と決定に 1 ～ 3 秒を費やします。その後初めて音声が開始されます。人間はその沈黙を破線として解釈します。彼らは「こんにちは?」と言い、同じことを繰り返し、エージェントの返事が遅いことについて話し合っています。そして今では、順番が遅い上に順番が乱れています。
リアルタイム モデルの超能力を放棄しました。最新の音声合成モデルは、ネイティブ韻律と即時バージイン処理を使用して 1 秒未満でターンテイクを実行します。サーバー ループの背後でそれらをゲートすると、インタラクティブ メディアが IVR に変換されます。発信者はそれに応じて動作します。
分類器は消えませんでした。サーバーは、次の動きを決定するためにすべての発話を理解する必要がありました。そのため、すべての分類機構 (およびその結合の大部分) は生き残り、現在はレイテンシー クリティカルなパス上に直接置かれています。
アーキテクチャ 2 からの深い教訓: 音声の場合、遅延はパフォーマンスの指標ではなく、製品です。ターンごとにサーバーの往復を追加するアーキテクチャは、それがどれほど正しくても、すでに失敗しています。人間であると感じるための待ち時間はおよそ 0.5 秒です。分類パイプラインはそれに適合しません。
アーキテクチャ 3: モデルが会話を所有する
アーキテクチャ 3: 会話、保証、レコードにはそれぞれ 1 人の所有者がいます。
3 番目のビルドは関係を逆転させ、これが機能します。リアルタイム音声読み上げモデルは、会話を完全に所有します。サーバー ゲートなしで、発信者の声を聞き、何を言うかを決定し、それを即座に言います。サーバーが気にするものはすべて、それぞれ 1 人の所有者を持つ 3 つの個別のチャネルを通じて表現されます。
生の会話→モデル。モデルは自動応答します。状態のキャプチャは、型指定されたツールを通じて行われます。フィールドごとに 1 つの小さく、詳細に説明されたツールが必要です。

記録に加えて、通話の分類、プロバイダーの解決、通話の終了のためのツールが含まれます。ツールの結果にはステアリング (「記録済み。次の欠落フィールドを尋ねる: X」) が含まれるため、サーバーが文を作成しなくても、モデルは常に次の動きを認識します。言語処理は無料です。モデルは言語を超えて呼び出し元にネイティブに従い、アーキテクチャ 1 の正規表現ミュージアムは削除されました。
厳格な保証 → 薄いポリシー層。サーバーは、モデルが保証できないことのみを強制します。つまり、必要な質問に答えられずに電話を切ることや、2 回の別れ話をすることを構造的に不可能にする単一の通話終了ゲート (小さなステート マシン - 1 つの canEndCall() 、1 つの終了タイマー、1 つのティアダウン タイマー)。電話番号の確定的な桁ごとのリードバック。サポートされていないロケールの言語フォールバック ルール。重要なのは、ポリシー層が介入することです。運転しません。これは、ルールが破られる場合にのみ機能します。
録音の忠実度 → ポストコールパス。ほとんどのコードを削除する秘訣は次のとおりです。データ品質をめぐるライブレースに勝とうとするのをやめてください。通話終了後、ワーカーは、ライブ コールで欠落した必須フィールドの完全なトランスクリプトに対して 1 つの LLM 抽出を実行します。バックフィルのみ (ライブ コールでキャプチャされた内容は決して上書きされません) で、証拠に基づいています (復元された値はすべて、発信者の発話を逐語的に引用する必要があります。引用も書き込みもありません)。ライブ システムは、ステアリングに十分なキャプチャのみを必要とします。オフィス向けの記録は後で照合され、完全な会話が手元にあり、遅延のプレッシャーはありません。
非決定論は現実です。モデルは、記録ツールを呼び出さずに、「わかりました」という応答を確認する場合があります。まさにこれが、呼び出し後パスが存在する理由であり、すべてのサーバー ガードがモデルの主張を信頼するのではなく検証する理由です。シングルメートル

私が出荷した最も有益なバグ: 上流コンポーネントが理由文字列で「すべての必須フィールドが収集された」と主張したため、呼び出しを終了するコード パス。確認する;決して信用しないでください。
プロンプトとツールのスキーマが API サーフェスになります。ツールの説明、ステアリング文字列、指示テキストはコードであるため、コードと同じレビュー規律が必要です。
テストは難しくなります。もう正確な発言を主張することはできません。シミュレートされた呼び出し元、設計意図に照らしてトランスクリプトを採点する LLM 審査員、およびフレークに関する統計規律が必要です。 (それは次の投稿です。)
ポリシー層が正確である必要がある少数の決定的プロンプトを挿入する重要な場合を除いて、単語完璧なスクリプトを使用することはやめます。
3 つのビルドすべてを振り返ると、ユーザーに聞こえるほぼすべての失敗の原因は同じでした。2 つのサブシステムはどちらも次のターンを所有していると信じていました。 2つのサヨナラ試合が続けて行われた。スクリプト化されたプロンプトがライブの回答について話します。発信者がすでに回答した再質問。さまざまな症状が 1 つの病気で発生します。
したがって、音声エージェントのアーキテクチャに関する問題は、「モデルをどのように制御するか?」ということではありません。それは所有権台帳です。
会話にはリアルタイム モデルという 1 つの所有者がいます。
それぞれのハード保証には 1 人の所有者がいます。つまり、小さくて退屈な単体テスト可能なポリシー モジュールです。通話終了などの横断的な状態は、連携するフラグの山ではなく、実際のステート マシンを取得します。
記録の所有者は 1 人です。誰とも競合しない通話後の調整パスです。
現在、私が音声エージェントのデフォルトとして保持しているルールは次のとおりです。
発信者とエージェントの音声の間にサーバーを往復させないでください。チェックを非同期で実行できない場合、または呼び出し後に実行できない場合、そのチェックはターン ループに属しません。
決定ごとに 1 人の所有者。 2 つのコード パスが両方とも通話を開始できる場合、または両方とも通話を終了できる場合は、すでにバグが出荷されています。

。
決して信用せず、確認してください。状態に関するモデル (およびポリシー) の主張は、単一の共有ゲートで実際のデータと照合するまでは仮説です。
バックフィルのみ、証拠に基づいたリカバリ。通話後のトランスクリプトを記録の信頼できる情報源とし、抽出されたすべての値に引用符が必要です。
言語固有のコードはありません。人々が何かを言うときの正規表現を書いている場合、ひどいことに、モデルを一度に 1 つの言語ずつ再実装することになります。
イベントだけでなく決定事項も記録します。すべてのゲート判定、すべての回復されたフィールド、すべての介入 - それぞれ 1 つの grepable 行。音声のバグは再構築されるものであり、再現されるものではありません。
3 番目のアーキテクチャは、最初の 2 つよりも賢いわけではなく、より謙虚です。ハンドルを巡ってリアルタイム モデルと争うことをやめ、サーバーのみが守れるいくつかの約束にサーバーを絞り込み、時計のない、会話全体が存在する 1 つの場所、つまり通話後への正確性の移動を行います。
このシリーズの次は、これをどのようにテストするか — シミュレートされた発信者、裁判官としての LLM スコアリング、およびテスト対象が非決定的である場合に必要な統計規律です。
Viva AI: 歯科医院向けの先駆的な AI 受付係、オフィス マネージャー、診療オプティマイザー
ファリド・ファダイエ氏が語る、なぜ AI が歯科業界の隠れた生産性危機を解決しているのか
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
こんにちは、私はファリド・ファダイエです、同僚です

[切り捨てられた]

## Original Extract

I built the same AI voice agent three times. Why server-side orchestrators collapse, why server-gated turns create dead air, and the architecture that works.

AI Voice Agent Architecture: Lessons From 3 Production Builds
Skip to content
Farid Fadaie
Blog
Home >
AI Voice Agent Architecture: What I Learned Building the Same Agent Three Times
AI Voice Agent Architecture: What I Learned Building the Same Agent Three Times
Post category: AI / Engineering
I built the same production voice agent three times. Same requirements, same telephony stack, same speech models available — three fundamentally different architectures. The first one collapsed under its own coupling: every bug fix broke something else. The second one was controllable and correct, and it destroyed the one thing a voice agent exists for: responding like a human, immediately. The third one is the only one that survived contact with real callers.
This post is the architecture write-up I wish I’d read before spending months and real money finding out the hard way. It’s about one question that turns out to decide everything: who owns the next turn?
The agent answers real phone calls. It has to do three things at once, and they pull in different directions:
Hold a natural conversation. Sub-second turn-taking, no dead air, graceful handling of interruptions, callers who ramble, and callers who switch languages mid-sentence.
Capture structured data reliably. A defined set of fields must end up in a database — names spelled right, numbers digit-perfect, nothing invented.
Enforce hard rules. A few behaviors must always happen (reading a phone number back digit by digit) and a few must never happen (hanging up while a required question is unanswered, promising things the business didn’t authorize).
Every architecture below is a different answer to where those three responsibilities live.
Architecture 1: the server-side orchestrator
Architecture 1: every caller turn fans out to parallel classifiers; shared flags couple everything to everything.
The first build treated the speech model as a mouthpiece. The server owned everything: each caller utterance was transcribed and fanned out to a battery of parallel classifiers — one deciding the call category, one watching for a provider name, one watching for end-of-call intent, one detecting reference-only questions, plus hand-maintained regex banks for language detection (including phonetic transliterations of “can you speak X to me?” in half a dozen scripts). A merge layer combined their outputs into a “turn analysis,” and a policy engine picked what to say next, often from scripted prompts.
Every decision is inspectable server code. When something goes wrong, there’s a line number.
Each classifier is small, cheap, and unit-testable in isolation.
Adding a rule looks easy: write another classifier, wire it into the merge.
Scripted prompts give you word-level control over what the agent says.
The classifiers were never actually independent. The merge layer and a growing web of shared session flags meant every concern’s guard read every other concern’s state. The provider gate needed to know about the end-of-call state; the end-of-call logic needed to know whether a language switch was pending; the language logic needed to know whether a scripted prompt was mid-playback. Fixing a bug in one guard reliably broke an invariant another guard depended on. This is the “too intertangled” failure: the system had no single owner for any decision, so every decision was made N times by N pieces of code that disagreed at the margins.
Regex banks don’t survive multilingual reality. The language-detection patterns grew into a museum of transliterations and edge cases — unfalsifiable, unreviewable, and wrong in ways you only discover on a live call. Every new language meant re-deriving every pattern. The model I was routing around already understood all of these languages natively.
Latency stacked. Multiple classifier round-trips per turn, then a policy decision, then speech synthesis. Each piece was fast; the pipeline wasn’t.
Mid-turn disagreement. When two classifiers fired at once — caller names a provider and says goodbye in the same breath — two subsystems both believed they owned the next turn. Most of the audible bugs (double prompts, talking over the caller, contradictory follow-ups) were exactly this.
The deep lesson from architecture 1: distributed decision-making without a single owner is the bug. It doesn’t matter how clean each component is; if N components can each initiate speech, you ship race conditions to people’s ears.
Architecture 2: server-gated turn-taking
Architecture 2: one decider, full control — and a server round-trip of dead air on every single turn.
The obvious fix for “too many deciders” is one decider. The second build made the server the single owner of every turn: the realtime speech model was muted by default and spoke only when the server explicitly triggered a response with instructions. Caller speaks → transcript → server classifies → server decides what should be said → server instructs the model to say it.
Maximal control. The agent literally could not speak unscripted. “Never say X” became enforceable; compliance reviews got easy.
A familiar mental model. Request in, decision, response out — it’s a web app with audio attached. Easy to reason about, easy to log, easy to test piece by piece.
One owner per turn. The race conditions from architecture 1 mostly disappeared.
Dead air, every single turn. The caller stops speaking; the server spends one to three seconds classifying and deciding; only then does audio start. Humans interpret that silence as a broken line. They say “hello?”, repeat themselves, talk over the agent’s late reply — and now the turn-taking is corrupted on top of being slow.
It threw away the realtime model’s superpower. Modern speech-to-speech models do sub-second turn-taking with native prosody and instant barge-in handling. Gating them behind a server loop converts an interactive medium back into an IVR. Callers behave accordingly.
The classifiers didn’t go away. The server still had to understand every utterance to decide the next move — so all the classification machinery (and most of its coupling) survived, now sitting directly on the latency-critical path.
The deep lesson from architecture 2: in voice, latency is not a performance metric — it’s the product. An architecture that adds a server round-trip to every turn has already failed, no matter how correct it is. The latency budget for feeling human is roughly half a second; classification pipelines don’t fit in it.
Architecture 3: the model owns the conversation
Architecture 3: conversation, guarantees, and the record each have exactly one owner.
The third build inverts the relationship, and it’s the one that works. The realtime speech-to-speech model owns the conversation outright: it hears the caller, decides what to say, and says it — immediately, with no server gate. Everything the server cares about is expressed through three separate channels, each with exactly one owner:
Live conversation → the model. The model auto-responds. State capture happens through typed tools — one small, well-described tool per field it must record, plus tools for categorizing the call, resolving a provider, and ending the call. Tool results carry steering (“recorded; ask for the next missing field: X”) so the model always knows the next move without the server composing its sentences. Language handling comes free: the model natively follows the caller across languages, and the regex museum from architecture 1 got deleted.
Hard guarantees → a thin policy layer. The server enforces only what the model cannot guarantee: a single end-of-call gate (a small state machine — one canEndCall() , one closing, one teardown timer) that makes it structurally impossible to hang up with required questions unanswered or to play two goodbyes; deterministic digit-by-digit read-back for phone numbers; language fallback rules for unsupported locales. Crucially, the policy layer intervenes ; it doesn’t drive. It acts only when a rule would otherwise break.
Record fidelity → a post-call pass. Here’s the trick that removed the most code: stop trying to win the live race for data quality. After the call ends, a worker runs one LLM extraction over the full transcript for any required field the live call missed — backfill-only (it never overwrites what the live call captured) and evidence-grounded (every recovered value must quote a verbatim caller utterance; no quote, no write). The live system only needs capture good enough for steering; the office-facing record is reconciled afterward, with the complete conversation in hand and zero latency pressure.
Non-determinism is real. The model will occasionally acknowledge an answer — “Got it” — without calling the record tool. That’s precisely why the post-call pass exists, and why every server guard verifies model claims instead of trusting them. The single most instructive bug I shipped: a code path that ended calls because an upstream component claimed “all required fields collected” in a reason string. Verify; never trust.
Your prompt and tool schemas become an API surface. Tool descriptions, steering strings, and instruction text need the same review discipline as code, because they are code.
Testing gets harder. You can’t assert exact utterances anymore. You need simulated callers, an LLM judge scoring transcripts against design intent, and statistical discipline about flakes. (That’s the next post.)
You give up word-perfect scripting — except where it matters, where the policy layer injects the handful of deterministic prompts that must be exact.
Looking back across all three builds, nearly every user-audible failure had the same root: two subsystems both believed they owned the next turn. Two goodbyes played back to back. A scripted prompt talking over a live answer. A re-asked question the caller had already answered. Different symptoms, one disease.
So the architecture question for voice agents isn’t “how do I control the model?” It’s an ownership ledger:
The conversation has one owner: the realtime model.
Each hard guarantee has one owner: a small, boring, unit-testable policy module — and cross-cutting state like end-of-call gets a real state machine, not a pile of cooperating flags.
The record has one owner: a post-call reconciliation pass that doesn’t race anyone.
Rules I now hold as defaults for any voice agent:
Never put a server round-trip between the caller and the agent’s voice. If a check can’t run async or after the call, it doesn’t belong in the turn loop.
One owner per decision. If two code paths can both initiate speech or both end the call, you’ve already shipped the bug.
Verify, never trust. Model (and policy) claims about state are hypotheses until checked against the actual data at a single shared gate.
Backfill-only, evidence-grounded recovery. Let the transcript be the source of truth for the record — after the call, with a quote required for every extracted value.
No language-specific code. If you’re writing a regex for how people say something, you’re re-implementing the model, badly, one language at a time.
Log decisions, not just events. Every gate verdict, every recovered field, every intervention — one greppable line each. Voice bugs are reconstructed, not reproduced.
The third architecture isn’t cleverer than the first two — it’s humbler. It stops fighting the realtime model for the steering wheel, narrows the server to the few promises only a server can keep, and moves correctness to the one place that has the whole conversation and no clock: after the call.
Next in this series: how we test this thing — simulated callers, LLM-as-a-judge scoring, and the statistical discipline you need when your test subject is non-deterministic.
Viva AI: Pioneering AI Receptionist, Office Manager, and Practice Optimizer for Dental Offices
Farid Fadaie on Why AI Is Fixing Dentistry’s Hidden Productivity Crisis
Save my name, email, and website in this browser for the next time I comment.
Hi, I’m Farid Fadaie, cofo

[truncated]
