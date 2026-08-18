---
source: "https://octomind.run/blog/octomind-0-44-2-release"
hn_url: "https://news.ycombinator.com/item?id=49347637"
title: "Removing Self-Verification from AI Coding Agents in Octomind"
article_title: "Octomind 0.44.2: The Supervisor Learns to Demand Proof | Octomind"
image: "https://octomind.run/blog/octomind-0-44-2-release/index.png"
author: "donk8r"
captured_at: "2026-08-18T16:20:16Z"
capture_tool: "hn-digest"
hn_id: 49347637
score: 2
comments: 0
posted_at: "2026-08-18T15:55:55Z"
tags:
  - hacker-news
  - translated
---

# Removing Self-Verification from AI Coding Agents in Octomind

- HN: [49347637](https://news.ycombinator.com/item?id=49347637)
- Source: [octomind.run](https://octomind.run/blog/octomind-0-44-2-release)
- Score: 2
- Comments: 0
- Posted: 2026-08-18T15:55:55Z

## Translation

タイトル: Octomind の AI コーディング エージェントから自己検証を削除
記事のタイトル: Octomind 0.44.2: 監督者は証拠を要求することを学ぶ |オクトマインド
説明: Octomind 0.44 は、証拠を中心にスーパーバイザーを再構築します。つまり、再起動後も存続する検証ポリシー、検証ゲートでの条件ごとの証明、およびモデルから完全に移動された計画です。さらに、ファジー ピッカーを備えたセッション タイトル、/new および /monitor コマンド、実際の進行状況に応じてリセットされる MCP タイムアウト
[切り捨てられた]

記事本文:
Octomind 0.44.2: 監督者は証拠を要求することを学ぶ |オクトマインド
Tap Products Cloud Hub Blog Vision Docs インストール パネル GitHub ⭐ 109 ブログに戻る Octomind 0.44.2: 監督者は証拠を要求することを学ぶ
Octomind 0.44 は、再起動後も存続する検証ポリシー、検証ゲートでの条件ごとの証明、およびモデルから完全に削除された計画など、証拠を中心にスーパーバイザーを再構築します。さらに、ファジー ピッカーを備えたセッション タイトル、/new および /monitor コマンド、および実際の進行状況に応じてリセットされる MCP タイムアウト。
長いエージェント セッションはすべて同じ方法で終了します。つまり、モデルが終了を通知します。時々それは真実です。場合によっては、「完了」とは、5 つのファイルのうち 3 つが編集されたこと、「確認」されたテスト実行は古いツリーに対して 4,000 トークン前に行われたこと、検証したと主張するファイルは、独自の編集内容を読み取って確認することによって検証されたことを意味する場合があります。
これら 2 つのエンディングの違いは、製品全体にあります。再確認が必要なエージェントはデモです。 「完了」が監査後に生き残るエージェントはツールです。
0.43.0 は、レシートを保持する圧縮に関するものでした。 0.44 ポイントは、パイプラインのもう一方の端でも同じ本能を示しています。監督者は、「完了」を、項目ごと、条件ごと、出所を伴う検証に耐える必要があるクレームとして扱うようになりました。そして同じリリースでは、計画はモデル自身に対して行うものではなくなりました。計画ツールは完全にモデルの手から離れ、作業を監視しチェックリストを正直に保つ外部マネージャーに置き換えられました。
このコアを中心に、このリリースでは、長時間セッションのユーザーが求めていた一連の生活の質を向上させる機能が追加されています。ファジーな起動ピッカーを備えたセッション タイトル、よりクリーンな /new コマンド、長時間実行される監視用の /monitor、および「遅い」と「スタック」の違いを最終的に理解する MCP タイムアウトです。
0.44。

0 は 8 月 17 日に出荷され、同じ日に 0.44.1 が続き、0.44.2 でサイクルが終了しました。以下ではすべて 0.44 ラインを 1 つのリリースとして参照し、インストールするのは 0.44.2 です。
検証は今では気分ではなく方針です
まず静かな構造変化があり、他のすべてがそれに基づいて成り立っているからです。
検証状態は、スーパーバイザの検出器内のブール ラッチ、つまりプロセスとともに存続および消滅するメモリ内フラグに存在していました。タスクの途中でラップトップを再起動し、セッションを再開すると、スーパーバイザーは記憶喪失になって戻ってきました。確認すると約束した内容の記憶はなく、エージェントの言葉をすべて受け入れるという新たな意欲が戻っていました。
0.44 では、ラッチが永続的な検証ポリシーに置き換えられます。現在はセッション レベルの状態です。セッションとともにシリアル化され、再開時に復元され、ガバナンス ハッシュに組み込まれ、長く圧縮されたセッションで目標の説明に投影されます。そのため、プロセスの再起動や圧縮サイクルによっても状態を振り払うことはできません。証拠台帳 (実際に実行、変更、観察された内容の記録) は、再起動後も同様に保持されます。これにより、台帳が消去されたというだけの理由で再開されたターンが検証ギャップのように見える、ある種の誤報も無効になります。ターンの回答は専用の台帳でも追跡され、メッセージ履歴から再推測されるのではなく、履歴書のトランスクリプトから復元されます。
気づく動作: 再開したセッションは、一度も終了しなかったセッションと同じように動作します。監督者が知っていたことは、今でも知っています。借りたものは今でも集まっています。
検証ゲートは形のある証拠を必要とします
検証ゲート (エージェントが自己報告を完了したときに実行されるチェックポイント) では、サイクルの中で最も深い再作業が行われました。テーマは単純です。全体的な「見た目が良い」という判定は、もはや検証者が行うことを許可されているものではありません。
こちらです

現在のゲートの仕組み。リクエストがタスクに解決されると、分類子はそこから証拠条件を導き出します。これは、正直に言うと「完了」するために明らかに真でなければならないことの履行チェックリストです。ゲートでは、検証者は構造化タグ内の各条件に対処する必要があり、決定論的パーサーが応答を追跡します。一致しない条件は常に全体的な PASS を上回ります。必要なタグが欠落している応答には、フォーマットのみの再試行が 1 回だけ行われます。実質的なギャップは、再度通過することはできません。
検証者が見ているものは、それに合わせてより豊かなものになります。グラウンド トゥルースには、最後のコマンド出力だけでなく、最近のコマンド出力のスライディング ウィンドウと git status が表示されるようになりました。また、git diff バジェットはファイルごとに分散されます。1 つの大きなファイルがバジェットを使い果たし、他の 4 つのファイルの変更を隠すことはできなくなりました。リゾルバーの必須カバレッジ クラスは 4 つから 6 つに増加し、名前付き形式と定量化されたカバレッジが追加されました。名前を付けるかカウントするリクエストは、名前を付けてカウントし直す条件を取得します。また、新しい非列挙カテゴリの形状は、メンバーをリストせずにカテゴリを定量化するリクエストをカバーします。
そして、古い抜け穴を塞ぐ来歴ビットがあります。変更後の検証がリードバックのみで行われたかどうかがゲートに通知されるようになりました。ファイルを編集してから自身の編集内容を再読み込みするエージェントは、チェックではなく差分を生成します。そして検証者はその違いを確認し、動作上の違いを要求できるようになります。
検証者自体にも規律が求められます。
フォールバックモデルはありません。設定した verifier_model が失敗した場合、ゲートはそのターンの間閉じられません。エージェント独自のモデルを暗黙的に借用するわけではありません。これは、ゲートがキャッチするために存在する死角を正確に継承することになります。
使い果たされた反復はハードストップです。ギャップ修正バジェット ( max_iterations 、デフォルトは 2) がなくなると、t

彼は、より友好的な裁判官に落ちずにターンストップした。
暗黙的な完了はゲートされます。ツール呼び出しの生成を静かに停止するエージェントは、完了を通知するエージェントと同じチェックポイントを取得します。
解答のみのターンは免除されます。 「なぜこの関数はミューテックスを必要とするのですか?」単に答えが必要な質問に突然変異チェックを引きずり込むことはありません。
サポートされる 2 つの変更により、これらすべてがより安価で安定したものになります。スーパーバイザー呼び出しでは、スキーマが適用された JSON 出力が使用されるようになり (プロバイダーがサポートしている場合はプロバイダー側​​で強制し、サポートしていない場合は寛大な回復を使用します)、分類子は、フォールバックに直接陥るのではなく、不正な応答に対して制限付きの再試行を取得します。現在、ゲートの交換には独自のバジェット ([supervisor.gate] の下の max_tokens = 8192) があり、検証者の推論と、検証者が表示する組み立てられた成果物のサイズの両方をカバーします。
証拠に拘束された主張 (claim_check) が同じパスでより正確になりました。明示的な <evidence> 引用符は現在のターン ツールの出力で使用する必要があり、通常の URL、パス、コード サンプルは引用であると推論されなくなりました。コードベース内のフィクスチャ データが外部ソースと間違われなくなりました。
企画はモデルに任せる
これは最初に感じる変化であり、哲学的なものです。
プランニングが出荷されて以来、プランはモデルが plan(command="start", ...) 、 plan(command="next", ...) と呼ぶツールでした。これは、モデルが独自の宿題を採点することを意味します。そして、完了したいモデルには、チェックリストを進め、タスクに完了のマークを付け、計画を閉じるという、完了したように見せる方法があります。書類を仕上げることは仕事をしているわけではありませんが、自己管理のチェックリストでは違いが分かりません。
0.44 には、モデル呼び出し可能な計画ツールはありません。計画は監督者が所有します。集中した仕事の場合は何も変わりません。集中した仕事はプランフリーのままであり、

作業のみおよび観察のみのタスクでは、計画をまったく作成できなくなります。作業に実際に依存フェーズやコンテキスト損失のリスクがある場合、スペシャリストは実際の作業応答とともにまばらな隠れたシグナルを発し、別個の適応型外部計画マネージャー (独自の安価なモデル、決定ごとに 1 つの構造化された要求/応答) が、要求、専門家のコンテキスト、および記録された証拠からランタイム所有の計画を作成、推進、改訂、または最終決定します。証拠はフェーズごとにチェックポイントが作成されます。サイドカー プランは 6 つのタスクに制限されています。 /plan は引き続きチェックリストを表示します。もう内部からゲームすることはできず、検証済み PASS では、残りのフェーズが自動的に終了します。複数のフェーズを満たす 1 つの検証済み成果物は、忙しい作業の再実行を強制するのではなく、まさにそのように解決されます。
正直、計画も古くなります。すべてのプランは、最後に実行された時期を追跡し、作成、変更、復元時に更新され、最新のメッセージより前のプランは明示的に古いマーカーで表示されます。会話が過ぎ去ったチェックリストは、静かに作業を進めるのではなく、自動的にアナウンスされます。
マネージャーは構成可能であり、デフォルトでオンになっています。
【監督・企画】
有効 = true
モデル = "オクトハブ:自動"
max_tokens = 2048 # 生成された JSON 決定。コンテキストを入力しない
trajectory_max_tokens = 4096 # 有界アシスタント/ツール入力スライス
Adopt_min_actions = 8
Adopt_min_distinct_actions = 4 導入のしきい値は広範な作業のみを指定します。分類子は、計画が形成される前に回答のみのターンを拒否します。また、プランナーの失敗はユーザーのターンごとにラッチされるため、プランナーの調子が悪い日があると、無制限の再試行ループではなく「このターンは計画なし」という状態に陥ります。
ここには副次的な利点が隠されています。計画ツールのスキーマは、計画を立てているかどうかに関係なく、モデルがすべてのターンで見つめるテキストでした。

関連性があった。そのトークン税はなくなりました。
削除により、組み込み MCP サーバーが再シャッフルされます。コアはセッションネイティブのプリミティブを保持するようになりました (recall)。新しいオーケストレーション サーバーは委任とバックグラウンド作業 (タップ、スケジュール、新しいモニター) を実行します。ランタイムはハーネスの再構成を維持します。ロールまたはタップ マニフェストが server_refs をピン留めしている場合、移行ガイドには古いものから新しいものへの完全なテーブルが含まれています。
デフォルトではオフになっている 1 つの新しいスーパーバイザ検出器は、オンにする価値があります: 再読み取り検出。スーパーバイザはパスごとに成功した読み取りを追跡し、変更が介在せずに同じパスがしきい値を超えてフェッチされると、モデルはアドバイスを受け取ります。モデルは、コンテンツを使用するのではなく、すでに持っているコンテンツを再フェッチします。パスを変更するとカウンターがリセットされるため、編集と検証のループがトリップすることはありません。
[スーパーバイザー.検出器]
read_threshold = 4 # 0 = off (デフォルト) 既存の逐次勧告バジェットを共有するため、これを有効にしてもスーパーバイザーがうるさくなることはありません。
さて、物語の展開は半分です。多くのセッションを実行し、デーモン ユーザーが数十を実行している場合、障害モードがわかります。ディレクトリは session_1755301892 でいっぱいで、火曜日に中途半端に完了した移行がどのセッションに保持されているかわかりません。
セッションにタイトルが付けられるようになりました。 /rename で設定するか、 /new Auth Refactor という名前で開始します。タイトルはセッションの隣の JSON サイドカー内に存在し (160 文字まで)、 /list と /info に表示され、ターミナル ウィンドウのヘッダーにプッシュされます。最終的にどのセッションがどれであるかをウィンドウ マネージャーが認識します。
そしてピッカーもいます。裸の --resume (セッション名なし) から始めると、名前、タイトル、ロール、モデルでフィルタリングできる、あいまい検索可能なセッションのリストが代替画面に表示されます。クエリ行は実際のリードライン編集 (Ctrl-A/E、Ctrl-U/K、Ctrl-W、Home/End) をサポートしており、Ctrl-N/P でリストを参照します。再開するセッションを選択してください

それを選択するか、新たなスタートを選択してください。
/session は廃止され、 /new が代わりに使用されます。これは、名前のとおりの動作を行います。octomind 実行で使用するのと同じ YYMMDD-basename-HHMM-uuid4short 形式で名前を付け、オプションのタイトルを付けて新しいセッションを開始します。古いコマンドの switch-or-maybe-create のあいまいさはなくなりました。セッション コマンド リファレンスには全文が記載されています。
2 つの小さな追加がそれを完成させます。 /monitor を使用すると、オーケストレーション サーバーのイベント ストリーム監視がスラッシュ コマンドに行われるため、ツール呼び出しを作成せずに、長時間実行されているモニターを確認できます。そして、コストの読み出しでは、総支出がメイン モデルのシェアから分割されるようになりました。スーパーバイザーの機械にかかるコストを初めて正確に確認できます。これは、octohub:auto が機械の処理を行うため、通常、単一の未検証の「完了」のクリーンアップにかかる費用に次ぐ四捨五入誤差です。
0.43.0 では、圧縮圧力ラダーが計算エンジンに置き換えられました。 0.44 はそのエンジンのシェイクダウン クルーズであり、正直な修正が 1 つ含まれています。
成長は、出力トークンのみではなく、完全なコンテキストで測定されるようになりました。 0.43.0 の設計では、出力のみが新しいマテリアルであると考えられています。実際には、ツールの結果とインジェクションによってコンテキストも拡張され、自律セッションではそれらが支配的になります。フルコンテキストの成長を測定することで、滑走路の予測が現実と一致します

[切り捨てられた]

## Original Extract

Octomind 0.44 rebuilds the supervisor around evidence: a verification policy that survives restarts, per-condition proof at the verify gate, and planning moved out of the model entirely. Plus session titles with a fuzzy picker, /new and /monitor commands, and MCP timeouts that reset on real progress
[truncated]

Octomind 0.44.2: The Supervisor Learns to Demand Proof | Octomind
Tap Products Cloud Hub Blog Vision Docs Install Panel GitHub ⭐ 109 Back to Blog Octomind 0.44.2: The Supervisor Learns to Demand Proof
Octomind 0.44 rebuilds the supervisor around evidence: a verification policy that survives restarts, per-condition proof at the verify gate, and planning moved out of the model entirely. Plus session titles with a fuzzy picker, /new and /monitor commands, and MCP timeouts that reset on real progress.
Every long agent session ends the same way: the model announces it's done. Sometimes that's true. Sometimes "done" means three of the five files got edited, the test run it "confirmed" happened four thousand tokens ago against an older tree, and the file it claims to have verified was verified by reading its own edit back and admiring it.
The difference between those two endings is the whole product. An agent you have to re-check is a demo; an agent whose "done" survives an audit is a tool.
0.43.0 was about compression that keeps receipts. 0.44 points the same instinct at the other end of the pipeline: the supervisor now treats "done" as a claim that has to survive verification — item by item, condition by condition, with provenance. And in the same release, planning stops being something the model does to itself: the plan tool is gone from the model's hands entirely, replaced by an external manager that watches the work and keeps the checklist honest.
Around that core, the release grows a set of quality-of-life capabilities that long-session users have been asking for: session titles with a fuzzy startup picker, a cleaner /new command, /monitor for long-running watches, and MCP timeouts that finally understand the difference between "slow" and "stuck."
0.44.0 shipped on August 17, 0.44.1 followed the same day, and 0.44.2 closes out the cycle — everything below refers to the 0.44 line as one release, and 0.44.2 is the one to install.
Verification is a policy now, not a mood
The quiet structural change first, because everything else stands on it.
Verification state used to live in boolean latches inside the supervisor's detectors — in-memory flags that lived and died with the process. Restart your laptop mid-task, resume the session, and the supervisor came back with amnesia: no memory of what it had promised to verify, and a fresh willingness to take the agent's word for things.
0.44 replaces the latches with a persisted verification policy . It's session-level state now: serialized with the session, restored on resume, folded into the governance hash, and projected into goal recitation on long compacted sessions — so neither a process restart nor a compression cycle can shake it off. The evidence ledger — the record of what was actually run, changed, and observed — persists across restarts the same way, which also kills a class of false alarms where a resumed turn looked like a verification gap simply because the ledger had been wiped. Turn answers are tracked in a dedicated ledger too, restored from the transcript on resume instead of being re-guessed from message history.
The behavior you'll notice: a session you resume behaves like a session you never left. What the supervisor knew, it still knows. What it was owed, it still collects.
The verify gate wants evidence with a shape
The verify gate — the checkpoint that runs when the agent self-reports done — got the deepest rework of the cycle. The theme is simple: a holistic "looks good" verdict is no longer a thing the verifier is allowed to produce.
Here's how the gate works now. When your request is resolved into a task, the classifier derives evidence conditions from it — a fulfillment checklist of what would have to be demonstrably true for "done" to be honest. At the gate, the verifier must address each condition in a structured tag, and a deterministic parser walks the response: an unmatched condition beats a holistic PASS, every time. A response missing its required tags gets exactly one format-only retry; substantive gaps don't get to retry their way through.
What the verifier sees got richer to match. Ground truth now carries a sliding window of recent command outputs instead of just the last one, plus git status , and the git diff budget is distributed per file — one large file can no longer eat the budget and hide the changes in four others. The resolver's mandatory coverage classes grew from four to six, adding named-form and quantified coverage: requests that name things or count things get conditions that name and count them back, and a new unenumerated-category shape covers requests that quantify over a category without listing its members.
And there's a provenance bit that closes an old loophole: the gate is now told whether post-change verification happened by read-back only . An agent that edits a file and then re-reads its own edit has produced a diff, not a check — and the verifier can now see the difference and demand a behavioral one.
The verifier itself is held to discipline too:
No fallback model. If your configured verifier_model fails, the gate fails closed for the turn — it does not silently borrow the agent's own model, which would inherit exactly the blind spots the gate exists to catch.
Exhausted iterations hard-stop. When the gap-fixing budget ( max_iterations , default 2) runs out, the turn stops instead of falling through to a friendlier judge.
Implicit done is gated. An agent that just quietly stops producing tool calls gets the same checkpoint as one that announces completion.
Answer-only turns are exempt. Asking "why does this function take a mutex?" doesn't drag mutation checks into a question that just needs an answer.
Two supporting changes make all of this cheaper and steadier: supervisor calls now use schema-enforced JSON output (with provider-side enforcement where the provider supports it and lenient recovery where it doesn't), and the classifier gets a bounded retry on malformed responses instead of falling straight to a fallback. The gate's exchange has its own budget now — max_tokens = 8192 under [supervisor.gate] — covering both the verifier's reasoning and the size of the assembled deliverable it's shown.
Evidence-bound claims ( claim_check ) got more precise in the same pass: explicit <evidence> quotes must occur in current-turn tool output, and ordinary URLs, paths, and code examples are no longer inferred to be citations — fixture data in your codebase stops being mistaken for an external source.
Planning leaves the model's hands
This is the change you'll feel first, and it's a philosophical one.
Since planning shipped, the plan was a tool the model called: plan(command="start", ...) , plan(command="next", ...) . Which meant the model graded its own homework — and a model that wants to be done has a way to look done: advance the checklist, mark tasks complete, close the plan. Finishing paperwork is not doing the work, but a self-managed checklist can't tell the difference.
In 0.44 there is no model-callable plan tool. Planning is supervisor-owned. For focused work, nothing changes — focused work stays plan-free, and answer-only and observe-only tasks can no longer form plans at all. When work genuinely has dependent phases or context-loss risk, the specialist emits a sparse hidden signal alongside a real work response, and a separate adaptive external plan manager — its own cheap model, one structured request/response per decision — creates, advances, revises, or finalizes the runtime-owned plan from the request, the specialist's context, and the recorded evidence. Evidence is checkpointed per phase. Sidecar plans are capped at six tasks. /plan still shows you the checklist; it just can't be gamed from inside anymore, and on a verified PASS the remaining phases close atomically — one verified deliverable satisfying several phases is settled as exactly that, rather than forcing busywork re-runs.
Plans age honestly, too. Every plan tracks when it was last engaged — refreshed on creation, mutation, and restoration — and a plan that predates your latest message renders with an explicit stale marker. A checklist the conversation has moved past announces itself instead of quietly steering the work.
The manager is configurable and on by default:
[supervisor.plan]
enabled = true
model = "octohub:auto"
max_tokens = 2048 # generated JSON decision; not input context
trajectory_max_tokens = 4096 # bounded assistant/tool input slice
adoption_min_actions = 8
adoption_min_distinct_actions = 4 The adoption thresholds only nominate broad work — the classifier still rejects answer-only turns before a plan can form. And planner failures are latched per user turn, so a planner having a bad day degrades to "no plan this turn" instead of an unbounded retry loop.
There's a side benefit hiding here: the plan tool's schema was text the model stared at on every single turn, whether or not planning was relevant. That token tax is gone.
The removal reshuffles the builtin MCP servers. core now holds session-native primitives ( recall ); a new orchestration server takes delegation and background work ( tap , schedule , and the new monitor ); runtime keeps harness reconfiguration. If your roles or tap manifests pin server_refs , the migration guide has the full old-to-new table.
One new supervisor detector, off by default, worth turning on: re-read detection . The supervisor tracks successful reads per path, and when the same path is fetched past a threshold with no intervening mutation, the model gets an advisory — it's re-fetching content it already has instead of using it. Mutating a path resets its counter, so edit-verify loops never trip it.
[supervisor.detectors]
reread_threshold = 4 # 0 = off (default) It shares the existing sequential-advisory budget, so enabling it can't turn the supervisor into a nag.
Now the expansion half of the story. If you run many sessions — and daemon users run dozens — you know the failure mode: a directory full of session_1755301892 and no idea which one held the migration you half-finished on Tuesday.
Sessions have titles now. Set one with /rename , or start named: /new Auth Refactor . Titles live in a JSON sidecar next to the session (capped at 160 characters), show up in /list and /info , and are pushed into the terminal window header — your window manager finally knows which session is which.
And there's a picker. Start with a bare --resume (no session name) and you get a fuzzy-searchable list of your sessions, filterable by name, title, role, and model, on the alternate screen. The query line supports real readline editing — Ctrl-A/E, Ctrl-U/K, Ctrl-W, Home/End — and Ctrl-N/P walks the list. Pick a session to resume it, or choose a fresh start.
/session is retired in favor of /new , which does what its name says: starts a fresh session, named in the same YYMMDD-basename-HHMM-uuid4short format octomind run uses, with an optional title. The old command's switch-or-maybe-create ambiguity is gone. The session commands reference has the full surface.
Two smaller additions round it out. /monitor brings the orchestration server's event-stream watches to a slash command, so you can check on long-running monitors without composing a tool call. And the cost readout now splits total spend from the main model's share — for the first time you can see exactly what the supervisor machinery costs you, which, with octohub:auto doing the mechanics, is typically a rounding error next to what a single unverified "done" costs to clean up.
0.43.0 replaced the compression pressure ladder with a computed engine; 0.44 is that engine's shakedown cruise, and it includes one honest correction.
Growth is now measured on the full context, not output tokens only. The 0.43.0 design reasoned that only output is new material. In practice, tool results and injections expand the context too — and on autonomous sessions they dominate. Measuring full-context growth makes runway predictions match realit

[truncated]
