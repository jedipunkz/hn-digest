---
source: "https://tim-schipper.nl/en/blog/claude-code-context-management"
hn_url: "https://news.ycombinator.com/item?id=49211443"
title: "Claude Code context management: when to /clear and when to /compact"
article_title: "Claude Code context management: when to /clear and when to /compact | Tim Schipper"
author: "opwizardx"
captured_at: "2026-08-07T15:44:58Z"
capture_tool: "hn-digest"
hn_id: 49211443
score: 2
comments: 1
posted_at: "2026-08-07T14:55:24Z"
tags:
  - hacker-news
  - translated
---

# Claude Code context management: when to /clear and when to /compact

- HN: [49211443](https://news.ycombinator.com/item?id=49211443)
- Source: [tim-schipper.nl](https://tim-schipper.nl/en/blog/claude-code-context-management)
- Score: 2
- Comments: 1
- Posted: 2026-08-07T14:55:24Z

## Translation

タイトル: Claude コードのコンテキスト管理: /clear と /compact のタイミング
記事のタイトル: Claude コードのコンテキスト管理: /clear のタイミングと /compact のタイミング |ティム・シッパー
説明: ウィンドウが大きくなると、Claude Code でのコンテキスト管理が難しくなりました。 /context が実際に何を測定するのか、/compact がセッションの非可逆再エンコードである理由、および代わりにクリアする場合に使用するルール。

記事本文:
コンテンツへスキップ tim-schipper.nl ホーム ブログ サービス エクスペリエンス プロジェクト お問い合わせ $ ask anything… Ctrl K ⇄ NL ≡ クロード コードのコンテキスト管理: /clear のタイミングと /compact のタイミング ~/blog / series/claude-pro July 31, 2026 · 7m read time · $ share Copy link クロード コードのコンテキスト管理: /clear のタイミングと /compact のタイミング
ウィンドウが大きくなると、Claude Code でのコンテキスト管理が難しくなります。 /context が実際に何を測定するのか、/compact がセッションの非可逆再エンコードである理由、および代わりにクリアする場合に使用するルール。
朝食から続いているセッションで /context を実行します。私の報告によると、空きスペースの壁が存在します。警告も圧縮通知も何も赤くならない。
セッションは午前9時よりもさらに悪化している。メーターにはそれを言う方法がありません。
劣化はウィンドウが埋まるずっと前から始まります
Chroma のコンテキスト腐敗作業では 18 のフロンティア モデルをテストし、単語を繰り返すという単純なタスクでは、入力が増加するにつれてどのモデルも信頼性が低下することがわかりました。彼らの LongMemEval 実行は最も明確な例です。同じ質問が、焦点を絞った ~300 トークンのプロンプトから、または埋め込まれた ~113k トークンの完全な会話から回答されています。両方に同じ情報が存在します。答えに大きな差がある。
NoLiMa は開始位置に番号を付けます。少なくとも 128,000 のコンテキストをアドバタイズする 13 のモデルのうち、11 のモデルは 32,000 トークンという短いコンテキストのベースラインの半分を下回りました。 GPT-4o は 99.3% から 69.7% になりました。
32,000 トークン。 1M ウィンドウで、進入方向の 3% です。
Anthropic はこのメカニズムをアテンションバジェットとして説明しています。すべてのトークンは他のすべてのトークンに対応するため、n 個のトークンは n² 個のペア関係を意味し、追加する新しいトークンはそれぞれ有限プールから消費されます。彼ら自身の指針は、

仕事が終わりました。 100 万トークンのウィンドウと並行して公開するのは奇妙ですが、それは正しいことでもあります。
長いセッションがコードに与える影響
SlopCodeBench (arXiv 2603.24755、2026 年 3 月) は、これに関する私が待ち望んでいた研究です。ワンショットのベンチマークの代わりに、エージェントは 93 のチェックポイントにわたって独自の以前のコードを拡張します。仕様は進化し、ターン間で修正されたリファレンス実装は渡されません。 Sonnet 4.5 から Opus 4.6 および GPT-5.x Codex バリアントまでの 11 モデル。
正確性の数値自体は厳しいものです。単一の問題をエンドツーエンドで解決したエージェントは存在せず、最高のチェックポイント解決率は 17.2% でしたが、最後のチェックポイントまでに厳密な解決率は 0.5% に低下しました。コストは問題全体で 2.9 倍に増加しましたが、正確性は改善されませんでした。
品質に関する数値は、あなたの習慣を変える必要があるものです。
構造侵食は軌道の 80% で増加しました。高複雑度の関数はコードベースあたり 4.1 から 37.0 に、ピーク循環複雑度は 27.1 から 68.2 になりました。
冗長性は 89.8% で増加し、構造的重複はほとんどの軌道で 66% 増加しました。
人間が管理するリポジトリは、同じ基準でほぼ横ばいを維持します。エージェントの軌道は、ほぼすべての反復で悪化します。
これらは、ダッシュボードでカバレッジと突然変異テストを超えて書いたものと同じ手段です: 複雑さの集中とクローンの検出。エージェントに長いセッションを指示すると、下り坂であることがわかります。
同紙は明らかな修正も試みた。品質を意識したプロンプトにより、GPT-5.4 では初期の冗長性が 34.5% 削減され、劣化の傾きは平行のままでした。より良いスタート地点が得られますが、同じように衰退します。セッションの長さが鍵となります。
長時間セッションが発生するケースは実際にあります
2026 年 3 月の論文では、コーディング エージェントが効果的なロングコンテキスト プロセッサであり、公開されているスタッドを上回ると主張しています。

最大 3 兆トークンのコーパスで、芸術の 17.3% が増加します。これは良い仕事であり、私の側に立つ前に読む価値があります。
ただし、それがどのように勝利するかを読んでください。エージェントはテキストをファイル システムに編成し、通常のツールを使用して操作します。コーパスはディスク上に残り、ウィンドウにはワーキング セットが保持されます。この結果は、肥大化したウィンドウが問題ないという証拠ではなく、ファイル システムがバルクに属する場所であるという証拠です。
これは、別の方向から見ても同じ結論です。
/context が実際に測定するもの
Claude Code 2.1.220 の /context は、ウィンドウをシステム プロンプト、システム ツール、MCP ツール、カスタム エージェント、メモリ ファイル、メッセージ、空き領域、および Autocompact バッファと呼ばれるほとんどの人がスクロールして通過する行などの名前付きカテゴリに分割します。ステータスの読み出しには、自動圧縮されるまでのパーセンテージが示され、さらにその下に「コンテキストが低い」という警告が表示されます。
そのバッファラインが便利です。これはあなたが費やすことのできない予約されたスペースなので、その瞬間が来たときにモデル自体の概要を書く余地がまだあります。使用可能なウィンドウは、箱に記載されている数字よりも小さいです。これは常にそうなっていました。
/doctor コマンドは、この残りの半分、つまり、何かを入力する前に、インストールによってすべてのセッションに読み込まれる内容を監査します。これはディスクベースの推定に範囲を設定し、ライブ測定の /context を明示的に示します。インストールの健全性とセッションの健全性は別の問題であり、リンターがあるのはそのうちの 1 つだけです。
/compact はそれ自体を要約したモデルです
圧縮プロンプトをバイナリから取り出すと、デザインがすぐそこにあります。 3 つのバリエーションがあります。1 つは会話全体を要約するもの、1 つは以前のプレフィックスがそのまま保持されている場合に最近の部分のみを要約するもの、もう 1 つは継続セッションの先頭に置かれるように書かれたものです。
それらはすべて同じ 8 つのセクションを要求します。一次要求

エストと意図。主要な技術概念。ファイルとコード セクション (完全なスニペット付き)。エラーと修正。問題解決。すべてのユーザーメッセージ。保留中のタスク。現在の仕事。
スペックは充実しています。サマライザーは、すでに劣化していると思われるコンテキストを使用して独自のセッションを評価するモデルであり、8 つのバケットのいずれかにノミネートされなかったものはすべて失われます。
このプロンプトに含まれる 2 つの詳細は注目に値します。サマライザは、セキュリティ関連のユーザー指示をそのまま保存して、その後も有効になるように指示されます。これは、そうしないと、再エンコードで制約が蒸発する可能性があることを明確に示しています。また、概要は、次のウィンドウに偽の指示をロンダリングしようとするあらゆるものにとってのソフトターゲットであるため、アシスタントメッセージ内のトランスクリプト形式のテキストをユーザーに帰属させないように警告されています。
次に、再開プロンプトは、あたかも中断が起こらなかったかのように、概要を確認することなくモデルに応答するように指示します。スムーズなデザイン。継ぎ目を感じることはありません。だからこそ、自動操縦で継ぎ目を発生させてはいけません。
早めにクリアし、めったにコンパクトにまとめ、自分で概要を書きます
部屋の外に出そうになると、自動圧縮が開始されます。したがって、サマリーは、精度が低下し始める 32k マークをはるかに超えた、セッションの最悪のバージョンによって書かれています。コンパクト化する場合は、リコールがまだ良好なうちに、意図的にヘッドルームを残して実行してください。
より良い: 必要ありません。タスクの境界を明確にし、残すべきものを要約ではなくファイルに入れます。 CLAUDE.md はすべてのセッションを超えて存続するメモリであり、自分で作成した圧縮の概要とは異なり、読み取ることができ、問題のある行を削除することもできます。エージェントがリポジトリから取得できない規則、落とし穴、および理論的根拠に従わないと、最終的に必要なメモリが残ってしまいます。

6週間後、あなたにこっそり嘘をつきます。
この概要を書くのに 2 分かかりますが、どの 4 つのことが重要かがわかり、モデルは 8 つを推測しているため、生成された要約よりも優れた成果物となります。
3 つのタスクを同時に実行するまでは、タスクごとに 1 つのクリーン セッションを実行するのが簡単なアドバイスです。これが git worktree の目的です。ただし、git worktree はファイルを分離するだけで、それ以外はほとんど分離しないことに注意してください。
ここでの測定の問題は、生成されたコードの場合と同じです。楽器は存在しており、無料ですが、まだ何も壊れていないため、誰も楽器を開けません。
複雑さと重複は、コードベースが漂流していることを示しています。 /context はセッションであることを示します。どちらも、見るのをやめた瞬間に沈黙し、答えがまだ退屈なうちに行動するのが最も安上がりです。
2026 年 5 月 27 日 トークンセーバー税: 穴居人のアドバイスを撤回する
6 日前、私は Caveman をコンテキストモードの上に重ねるように言いました。トークノミクス演算は、正直に実行された場合、インストール後は存続しません。ここでは私が見逃していたものと、今お勧めしたいものを紹介します。
2026 年 5 月 21 日 穴居人 vs コンテキストモード: 口が小さいか、部屋が小さいか?
ある Claude Code プラグインには 63,000 個のスターがあり、穴居人のように話すように求められます。もう 1 つは 15,000 個のスターとツールの出力をサンドボックス化します。インターネットは面白いものを選びました。必要かどうかは、実際にどのトークン リークを修正しようとしているかによって異なります。
2026 年 5 月 14 日 エージェントに Markdown を記述させるのをやめる
AI エージェントが Markdown ではなく HTML を生成することに誰もが興奮しています。出力は美しく見えます。しかし、エージェントの応答がすべて使い捨ての Web ページになった場合に、それにどれくらいの費用がかかるのか、あるいは何を失うのかを誰も知りません。

## Original Extract

Context management in Claude Code got harder when the window got bigger. What /context actually measures, why /compact is a lossy re-encode of your session, and the rule I use for when to clear instead.

Skip to content tim-schipper.nl Home Blog Services Experience Projects Contact $ ask anything… Ctrl K ⇄ NL ≡ Claude Code context management: when to /clear and when to /compact ~/blog / series/claude-pro July 31, 2026 · 7m read time · $ share Copy link Claude Code context management: when to /clear and when to /compact
Context management in Claude Code got harder when the window got bigger. What /context actually measures, why /compact is a lossy re-encode of your session, and the rule I use for when to clear instead.
Run /context in a session that has been going since breakfast. Mine reports a wall of free space. No warning, no compaction notice, nothing red.
The session is still worse than it was at nine in the morning. The meter just has no way to say so.
Degradation starts long before the window fills ​
Chroma's context-rot work tested 18 frontier models and found every one of them gets less reliable as input grows, on tasks as simple as repeating words back. Their LongMemEval run is the cleanest illustration: the same question, answered from a focused ~300-token prompt or from the full ~113k-token conversation it was buried in. Same information present in both. Large gap in the answers.
NoLiMa puts a number on where it starts. Across 13 models that all advertise at least 128k of context, 11 dropped below half their short-context baseline at 32k tokens . GPT-4o went from 99.3% to 69.7%.
Thirty-two thousand tokens. On a 1M window that is three percent of the way in.
Anthropic describes the mechanism as an attention budget: every token attends to every other token, so n tokens means n² pairwise relationships, and each new token you add spends from a finite pool. Their own guidance is to find the smallest set of high-signal tokens that gets the job done. That is a strange thing to publish alongside a million-token window, and it is also correct.
What a long session does to the code ​
SlopCodeBench (arXiv 2603.24755, March 2026) is the study I have been waiting for on this. Instead of one-shot benchmarks it makes agents extend their own prior code across 93 checkpoints, with evolving specs and no corrected reference implementation handed back between turns. Eleven models, from Sonnet 4.5 up to Opus 4.6 and the GPT-5.x Codex variants.
The correctness numbers are grim on their own: no agent solved a single problem end to end, the best checkpoint solve rate was 17.2%, and by the final checkpoint strict solve rates collapsed to 0.5%. Cost grew 2.9× across a problem while correctness did not improve.
The quality numbers are the ones that should change your habits:
Structural erosion rose in 80% of trajectories. High-complexity functions went from 4.1 to 37.0 per codebase, peak cyclomatic complexity from 27.1 to 68.2.
Verbosity rose in 89.8% , with structural duplication up 66% across most trajectories.
Human-maintained repositories stay roughly flat on the same measures. Agent trajectories deteriorate with nearly every iteration.
Those are the same instruments I wrote about in the dashboard beyond coverage and mutation testing : complexity concentration and clone detection. Point them at a long agent session and they read as a downward slope.
The paper also tried the obvious fix. Quality-aware prompting cut initial verbosity by 34.5% on GPT-5.4, and the degradation slopes stayed parallel. You get a better starting point and the same decline. Session length is the lever.
The case for long sessions is real ​
A March 2026 paper argues that coding agents are effective long-context processors, beating published state of the art by 17.3% on corpora up to three trillion tokens. It is good work and worth reading before you take my side of this.
Read how it wins, though. The agents organise the text into a file system and manipulate it with ordinary tools. The corpus stays on disk and the window holds a working set. That result is not evidence that a bloated window is fine, it is evidence that the filesystem is where bulk belongs.
Which is the same conclusion from the other direction.
What /context actually measures ​
/context in Claude Code 2.1.220 breaks your window into named categories: system prompt, system tools, MCP tools, custom agents, memory files, messages, free space, and a line most people scroll past called Autocompact buffer . The status readout tells you the percentage until auto-compact, and there is a "Context low" warning further down the line.
That buffer line is the useful one. It is space you cannot spend, reserved so the model still has room to write a summary of itself when the moment comes. Your usable window is smaller than the number on the box, and it always was.
The /doctor command audits the other half of this: what your install loads into every session before you type anything. It scopes itself to disk-based estimates and explicitly points you at /context for the live measurement. Install hygiene and session hygiene are separate problems, and only one of them has a linter.
/compact is a model summarising itself ​
Pull the compaction prompts out of the binary and the design is right there. Three variants: one that summarises the whole conversation, one that summarises only the recent portion when an earlier prefix is being kept intact, and one written to sit at the head of a continuing session.
All of them demand the same eight sections. Primary request and intent. Key technical concepts. Files and code sections, with full snippets. Errors and fixes. Problem solving. All user messages. Pending tasks. Current work.
The spec is thorough. The summariser is a model grading its own session with the context you already suspect is degraded, and everything that did not get nominated into one of those eight buckets is gone.
Two details in that prompt are worth your attention. The summariser is told to preserve security-relevant user instructions verbatim so they remain in effect afterwards, which tells you plainly that constraints can otherwise evaporate in the re-encode. And it is warned never to attribute transcript-shaped text inside assistant messages to the user, because a summary is a soft target for anything that wants to launder a fake instruction into your next window.
Then the resumption prompt tells the model to pick up without acknowledging the summary, as if the break never happened. Smooth by design. You will not feel the seam, which is exactly why you should not let it happen on autopilot.
Clear early, compact rarely, write the brief yourself ​
Auto-compaction fires when you are nearly out of room. The summary is therefore written by the worst version of the session, long past the 32k mark where accuracy starts sliding. If you are going to compact, do it with headroom left, deliberately, while recall is still good.
Better: do not need it. Clear on task boundaries, and put what should survive into a file instead of a summary. CLAUDE.md is the memory that outlives every session , and unlike a compaction summary you wrote it, you can read it, and you can delete a line that has gone wrong. Keep it to conventions, pitfalls and rationale that the agent cannot derive from the repo, or you end up with memory that confidently lies to you six weeks later .
Writing that brief takes two minutes and it is a better artefact than any generated summary, because you know which four things mattered and the model is guessing at eight.
One clean session per task is easy advice until you want three tasks running at once. That is what git worktrees are for , with the caveat that they isolate your files and remarkably little else.
The measurement problem here is the same one as with generated code. The instruments exist, they are free, and nobody opens them because nothing has broken yet.
Complexity and duplication tell you the codebase is drifting. /context tells you the session is. Both go quiet the moment you stop looking, and both are cheapest to act on while the answer is still boring.
May 27, 2026 The token-saver tax: walking back my Caveman advice
Six days ago I told you to stack Caveman on top of context-mode. The tokenomics arithmetic, done honestly, doesn't survive the install. Here's what I missed, and what I'd recommend now.
May 21, 2026 Caveman vs context-mode: small mouth, or smaller room?
One Claude Code plugin has 63k stars and asks you to talk like a caveman. The other has 15k stars and sandboxes your tool output. The internet picked the funny one. Whether you should depends on which token leak you are actually trying to fix.
May 14, 2026 Stop letting your agents write Markdown
Everyone is excited about AI agents generating HTML instead of Markdown. The output looks beautiful. But nobody is asking what it costs, or what we lose when every agent response becomes a single-use webpage.
