---
source: "https://www.questionpro.com/engineering/engineering/developer%20tools/ai%20&%20machine%20learning/stop-asking-claude-to-agree-with-you/"
hn_url: "https://news.ycombinator.com/item?id=48456990"
title: "Stop Asking Claude to Agree with You"
article_title: "Stop Asking Claude to Agree With You - Engineering Blog"
author: "skyDoesWork38"
captured_at: "2026-06-09T07:22:16Z"
capture_tool: "hn-digest"
hn_id: 48456990
score: 1
comments: 0
posted_at: "2026-06-09T05:46:27Z"
tags:
  - hacker-news
  - translated
---

# Stop Asking Claude to Agree with You

- HN: [48456990](https://news.ycombinator.com/item?id=48456990)
- Source: [www.questionpro.com](https://www.questionpro.com/engineering/engineering/developer%20tools/ai%20&%20machine%20learning/stop-asking-claude-to-agree-with-you/)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T05:46:27Z

## Translation

タイトル: クロードに同意を求めるのはやめてください
記事のタイトル: クロードに同意を求めるのはやめてください - エンジニアリング ブログ
説明: 私が見つけた最高のクロード スキルは、コードを記述しません。自分が何を構築しているのかがわかるまでは始められません。

記事本文:
ソフトウェア エンジニア • フルスタック
AI 愛好家
クロードに同意を求めるのはやめてください
私が見つけたクロードの最高のスキルは、コードを書かないことです。自分が何を構築しているのかがわかるまでは始められません。
長い間、重要な機能に対する私のデフォルトの動きは、クロード コードのプラン モードでした。アイデアをスケッチし、計画の下書きを作成し、ざっと目を通し、実行ボタンを押します。
曖昧な質問は曖昧な計画を生み出します。出力は構造化され自信に満ちているように見えますが、あなたの思考のあらゆるギャップを静かに継承しています。
エージェントは実行中に漂流します。計画を承認し、実行に切り替えると、計画には記載されていないファイルやパターンが静かに追加されます。結局、すべての実行に「書かれたとおりに実装し、余分なものは何もありません」と追加することになります。
アイデアそのものに異議を唱えることは決してありません。計画モードでは、クロードがあなたのものをどのように構築するかを答えます。それを構築する価値があるかどうか、あるいは実際にそれを考え抜いたかどうかは決して問われません。 3 日も費やせば、きれいに構造化された間違いが生まれる可能性があります。
最後の部分が本当のギャップです。計画モードでは、アイデアがすでにしっかりしていて、問題は実行だけであると想定します。ほとんどの場合、その仮定は間違っています。そして、それこそが /grill-me によって埋められるギャップなのです。
Total TypeScript の背後にある TypeScript 教育者である Matt Pocock は、2026 年 2 月 3 日に個人の .claude ディレクトリを公開スキル リポジトリとして公開しました。3 か月以内に、このディレクトリは GitHub のスター数 50,000 を超え、GitHub のトレンド チャートのトップ近くに達しました。フレームワークでもラッパーでもありません。彼が毎日使用している実際のマークダウン ファイルだけです。
このリポジトリは、1 つの障害モード「不整合」を対象としています。エージェントはあなたのことを理解してくれたと思います。次に、それが何を構築したかがわかりますが、明らかにそうではありませんでした。人間の開発者と同じコミュニケーション ギャップが、マシンの速度で実現されました。
見出しの修正は /grill-me です。以前の曖昧さを取り除きます

計画を一度に 1 つずつ決定することで、あらゆるコードが存在します。スキル全体は次の 3 つの文で構成されます。
この件に関するあらゆる側面について執拗にインタビューしてください
共通の理解に達するまで計画を立てます。下を歩く
デザイン ツリーの各ブランチ、依存関係の解決
一つ一つの決断の間。
それでおしまい。賢い機械はありません。答えを強制されることにこそ価値があります。
npx スキルで mattpocock/スキルを追加
次に、/setup-matt-pocock-skills をリポジトリごとに 1 回実行して、問題トラッカーと構成を接続します。
答えたくない質問
実際の機能を取り上げます。マルチテナント SaaS におけるワークスペースごとの請求制限です。アイデアをスケッチし、 /grill-me と入力します。
クロード: ワークスペースが制限に達すると、ワークスペース全体がブロックされますか、それとも新しい使用だけがブロックされますか?
クロード: 通知を受け取るのは誰ですか? それをトリガーしたユーザー、ワークスペース管理者、あるいはその両方ですか?どのようなチャネルを通じて?
クロード: 管理者がそのメールを無視し、ユーザーがセッション中に壁にぶつかったら、どうなりますか?ハードエラー？劣化モード？優しいメッセージ？
クロード: それがあなたの劣化の物語のすべてです。これにより、API コントラクト、フロントエンド、請求ページに表示する内容が決まります。続けていきましょう。
大まかなアイデアを 1 つ。 3 日後にバグとして表面化するのではなく、実装前にアーキテクチャを再構築する少数の決定。
グリル セッションでのすべての回答は、コード レビューで覆す必要のない決定となります。
イチジク。 01 · 4 ステップのワークフロー。
全体が単一の連続セッションとして実行され、各ステップが次のステップに影響を与えます。
/グリルミー → /to-prd → /to-issues → /afk
これを機能させるためのルール: ステップ間で /clear を実行しないこと。 PRD スキルは、答え、推論、指定したエッジ ケース、受け入れたトレードオフなど、グリルからのあらゆるものに依存します。 PRD で生成されたコールドはテンプレートです。あ

グリル直後に発生するPRDは視点を持った仕様だ。
緊密な PRD 入力、緊密なチケット出力 — /to-issue は、「認証の実装」ではなく、実際のコンテキストと受け入れ基準を備えたタスクを生成します。チケットがタイトであるということは、/afk が推測することなく実行できることを意味します。
後で現れる2番目の見返りがあります。 PRD と問題は書かれた成果物であり、チャット ウィンドウだけでなくトラッカー内に存在します。したがって、セッションが途中で終了した場合、または 1 週間後に同じ機能に戻った場合、意図を記憶から再構築しているわけではありません。 PRD と未解決の問題に新たなセッションを向けると、すでに手元にある完全な推論、つまり何を決定したのか、なぜ決定したのか、何が残っているのかがわかります。グリルは一度だけ起こります。それが生成するコンテキストは、その機能に触れるすべてのセッションにわたって支払いを続けます。
最後の出力品質は、最初の誠実さに直接影響されます。
グリルが硬くなったら：/handoff
ディープ グリル セッションではコンテキストが焼き付けられます。何かのプロトタイプを作成したり、既存のスキーマがどのように動作するかを確認したりする必要があることに気付いたとき、あなたはこのスキーマに十分取り組んでいます。本能的には、同じセッション内ですぐにそれを実行することになります。
Pocock は /handoff も書きました。これはまさにこれのために構築されました。現在のセッションを文書に圧縮し、新人エージェントがコンテキスト、決定事項、意図、提案される次のステップなどを取得できるようにします。グリルはきれいなままです。サイドクエストには独自のウィンドウが表示されます。
ファイアアンドフォーゲット - 範囲外の何かが表示されます。それを新鮮なエージェントに渡し、グリルに戻します。
グリル→ハンドオフ→プロトタイプ→ハンドオフバック — 使い捨てセッションで仮説を検証し、学んだことを家に持ち帰ります。
質問が難しいほど、グリルにはより多くのコンテキストが蓄積され、グリルを保護する価値が高まります。 /handoff は、セッションを溺れずにハードに続ける方法です。
グリルが深いほど、

脱線して負けることです。ハンドオフを使用すると、接線を追いかけて、まだ鮮明なセッションに戻ることができます。
スタックに追加する価値のあるスキル
/tdd — 一度に 1 つの垂直スライスを構築する赤-緑-リファクタリング ループ: 失敗したテスト、最小限の実装、リファクタリング。グリルと PRD が徹底していれば、合格基準はすでに存在します。これは、何かが機能すると主張する前に、クロードが満たさなければならないテストになります。
/zoom-out — 馴染みのないコード部分に到達し、それが全体像にどのように適合するのかがわからない場合、これによりエージェントが 1 つ上のレベルに引き上げられます。セクションの役割、その依存関係、およびシステムの残りの部分との接続方法が説明されます。モジュールを継承する場合、古いコードに戻る場合、または何かを変更する前に状況を確認する必要がある場合に便利です。
/grill-with-docs — 既存のドメイン モデルとコードベースに対して計画をテストする /grill-me のバリアント。新たに開始するのではなく、ライブ システムを拡張する場合には正しい選択です。新しい決定を既存のものに結び付け、コンテキスト ドキュメントをインラインで更新します。
グリルを超えて: Claude Code を無駄なく実行できるようにする
グリルはワークフローの主役ですが、セットアップの残りの部分がトークンとコンテキストを静かに滲ませていない場合にのみ効果を発揮します。これらは、Claude Code を高速、低コスト、そしてシャープに保つ習慣です。
最強のモデルでグリル セッションを実行します。安価なもので実装します。
グリルは純粋な探索であり、ファイルへの書き込みやツール呼び出しのオーバーヘッドはありません。そのため、Opus のようなフロンティア モデルは、より鋭い質問をし、小規模なモデルが見逃している矛盾を見つけることでコストを稼ぎます。また、単なる会話であるため、トークンの請求額は低く抑えられます。実装 ( /afk 、 /tdd 、実行) に入ると、Sonnet のようなモデルが、数分の 1 の価格で重労働を行ってくれます。
過ごす

コードを入力する場所ではなく、意思決定が行われる場所で最適なモデルを作成します。
.claudeignore を使用すると、1 回の編集でコンテキストの負担が軽減されます。
デフォルトでは、Claude はプロジェクト内のすべてのものをスコープに取り込みます。 .claudeignore ファイル ( .gitignore と同じ構文) は、デッドウェイトの自動読み込みを停止します。 .next/ を除外するだけで、Next.js リポジトリのコンテキストを 30 ～ 40% 削減できます。これは、あなたが費やす最高のレバレッジの 2 分間です。
ノードモジュール/
距離/
ビルド/
.next/
__pycache__/
*.lock
.git/
*.db
これは、明示的に要求した場合にクロードがこれらのファイルを読み取るのを停止するものではなく、自動探索から除外するだけです。
ファイルは 200 ～ 300 行以内に収めてください。
エージェントはファイル全体を読み取ります。 20 行の関数を見つけるために 1,500 行のファイルを読み取ると、1,480 行のノイズが発生し、さらに悪いことに原因が生じます。これは即効性のある方法ではありません。これは、セッション内のすべての読み取りにわたって複合化されるコードベースの規律です。
大きなファイルは、すべてのエージェントが参照するたびに支払う税金です。
CLAUDE.md でクロードに読み取りではなく grep を教えます。
CLAUDE.md は、クロードが各セッションの開始時に読み取る 1 つのファイルです。つまり、スタック、規約、および決して再説明したくない事項など、プロジェクトの固定ルールが存在する場所です。そこに設定できる最も価値のあるルールは、 read の代わりにどのファイルを検索するかを指示することです。 2,000 を超えるエントリを含む翻訳キー ファイルがあります。それを全部読むと、コンテキストの予算がノイズで消費されてしまいます。そこで CLAUDE.md は次のように言います。
`packages/survey-type-defs/src/appTranslation/AppTranslationKeyEnum.ts` の grep
— 2000 以上のキー、ファイル全体を読み取ることはありません
たった 1 行で、Claude は 1 つのキーが必要になるたびに巨大なファイルをコンテキストに取り込むのをやめます。リポジトリ内のサイズが大きくなりすぎて避けられないすべてのファイルに対してこれを実行します。
使用していない MCP サーバーを無効にします。
従来、Claude Code は MCP ツール定義を直接ロードしていました。

コンテキストに応じて、いくつかの重いサーバー (Playwright、GitHub、Gmail など) がそれぞれ数千のトークンを静かに消費する可能性があり、ターンごとに数万の広大なセットアップが必要になります。最近の Claude Code バージョンでは、デフォルトで有効になっているツール検索によってこの問題が軽減されます。ツールは、すべてのスキーマを事前に配布するのではなく、オンデマンドで検出されてロードされます。それは役に立ちますが、大規模な MCP セットアップは依然としてオーバーヘッドと複雑さを追加します。
/context を実行して各サーバーにかかるコストを調べ、未使用のサーバーを /mcp でプルーニングします。
壊れたときではなく、半分くらい埋まった状態で早めにコンパクトにしましょう。
ウィンドウがいっぱいになると品質が低下します。反応が曖昧に感じられる頃には、あなたはすでに腐敗ゾーンに深く陥っています。実際のスイート スポットは、約 50 ～ 60% の容量を圧縮することですが、クロードには、要約するための完全な非圧縮コンテキストがまだ残っています。 90% まで待つということは、すでに劣化したビューを要約していることを意味します。
並列エージェント用の Git ワークツリー。
独立した機能を次々に実行する必要はありません。ワークツリーは各エージェントに独自のブランチと同じリポジトリから分離されたディレクトリを提供するため、3 人のエージェントが連続して実行する作業を 1 時間で実行できます。コンテキストの衝突やマージの混乱はありません。
思考が新たなボトルネックになる
構築ソフトウェアの形は反転しました。
以前はコードを書くことがボトルネックでしたが、今ではエージェントが最も得意とする部分になっています。私たちに残されているのは、実際に何を構築する価値があるのか​​を判断し、実装がずれないように十分に正確であることです。
このワークフローで最も活用度の高いスキルが /grill-me である理由はここにあります。これは、コードをまったく生成しないにもかかわらず、他の何よりも活用度を高めるスキルです。
レバレッジは上流に移動しました。生の出力よりも、特異性、エッジケースの誠実さ、厳密なコンテキスト管理が重要です。
その他すべて - 無駄のないコンテキスト、より小さなファイル、適切なモデル

適切なフェーズ — その思考をノイズから守るために存在します。
グリルが最初に来ます。コードは二番目です。
タグ:
AIコーディング 、
クロードコード、
コンテキスト管理、
開発者の経験、
グリルミー 、
生産性、
スキル、
ワークフロー
カテゴリ:
AIと機械学習、
開発者ツール、
エンジニアリング
リリース時に書き換えが失敗しない。彼らはその後何が起こるかについて失敗します。
すべての書き換えには目に見えない負債が引き継がれます。エッジケースでは、古いシステムが壊れるまで何年もの間、誰にも知られずに黙って処理されていました。
LLM コンテキストの切り捨ては取得ではない: スライス(0, N) が間違った修正だった理由
260 万文字の LLM プロンプトを、slice(0, 10000) でキャップしました。レイテンシーが低下し、ダッシュボードが緑色に変わり、モデルが間違ったコンテキストで応答し始めました。
一度に 1 つのナビ: レガシーへのフィードを停止し、置き換えを開始した方法
ストラングラー・フィグ — 私たちはそれを書き換えたわけではありません。私たちはそれを超えました。
限界: エンジニアリング上の決定は、本番環境で決定されるまで先送りされ続けます
大規模な実行により、私たちが設定していない限界について学びました。

## Original Extract

The best claude skill I’ve found writes no code. It just won’t let me start until I know what I’m building.

Software Engineer • Full-Stack
AI Enthusiast
Stop Asking Claude to Agree With You
The best claude skill I’ve found writes no code. It just won’t let me start until I know what I’m building.
For a long while, my default move with any non-trivial feature was Claude Code’s plan mode. Sketch the idea, let it draft a plan, skim it, hit go.
A fuzzy ask produces a fuzzy plan. The output looks structured and confident, but it’s quietly inheriting every gap in your thinking.
The agent drifts mid-execution. You approve a plan, switch to execute, and it quietly adds files and patterns the plan never mentioned. You end up appending “implement exactly as written, nothing extra” to every run.
It never challenges the idea itself. Plan mode answers how Claude would build your thing. It never asks whether the thing is worth building, or whether you’ve actually thought it through. You can burn three days producing a clean, well-structured mistake.
That last one is the real gap. Plan mode assumes your idea is already sound and your only problem is execution. Most of the time, that assumption is wrong — and that’s the exact gap /grill-me fills.
Matt Pocock — the TypeScript educator behind Total TypeScript — opened up his personal .claude directory as a public skills repo on February 3, 2026. Within three months it had crossed 50,000 GitHub stars, peaking near the top of GitHub’s trending charts. Not a framework, not a wrapper — just the actual markdown files he uses every day.
The repo targets one failure mode: misalignment . You think the agent understood you. Then you see what it built, and it clearly didn’t. The same communication gap you get with a human dev, now at machine speed.
The headline fix is /grill-me . It removes vagueness before any code exists, by interrogating your plan one decision at a time. The entire skill is three sentences:
Interview me relentlessly about every aspect of this
plan until we reach a shared understanding. Walk down
each branch of the design tree, resolving dependencies
between decisions one by one.
That’s it. No clever machinery — the value is entirely in being forced to answer.
npx skills add mattpocock/skills
Then run /setup-matt-pocock-skills once per repo to wire up your issue tracker and config.
The questions you didn’t want to answer
Take a real feature — per-workspace billing limits in a multi-tenant SaaS. You sketch the idea and type /grill-me .
Claude: When a workspace hits its limit, does the whole workspace get blocked, or just new usage?
Claude: Who gets notified — the user who triggered it, the workspace admin, or both? Through what channel?
Claude: If the admin ignores that email and a user hits the wall mid-session, what’s the experience? Hard error? Degraded mode? Graceful message?
Claude: That’s your entire degradation story. It shapes the API contract, the frontend, and what the billing page has to show. Let’s keep going.
One rough idea. A handful of decisions that reshape the architecture before implementation, instead of surfacing as bugs three days in.
Every answer you give in a grill session is a decision you won’t have to reverse in code review.
FIG. 01 · The four-step workflow.
The whole thing runs as a single continuous session, each step feeding the next.
/grill-me → /to-prd → /to-issues → /afk
The rule that makes it work: never /clear between steps. The PRD skill leans on everything from the grill — your answers, your reasoning, the edge cases you named, the trade-offs you accepted. A PRD generated cold is a template. A PRD generated right after a grill is a spec with a point of view.
Tight PRD in, tight tickets out — /to-issues produces tasks with real context and acceptance criteria, not “implement auth.” Tight tickets mean /afk can execute without guessing.
There’s a second payoff that shows up later. The PRD and the issues are written artifacts — they live in your tracker, not just in a chat window. So when a session dies halfway, or you come back to the same feature a week later, you’re not reconstructing intent from memory. You point a fresh session at the PRD and the open issues, and it picks up with the full reasoning already in hand: what you decided, why, and what’s left. The grill happens once; the context it produces keeps paying out across every session that touches the feature.
Output quality at the end is a direct function of honesty at the beginning.
When the grill gets hard: /handoff
Deep grill sessions burn context. You’re well into one when you realize you need to prototype something or check how an existing schema behaves. The instinct is to do it right there in the same session.
Pocock also wrote /handoff , built for exactly this. It compresses the current session into a document a fresh agent can pick up — context, decisions, intent, suggested next steps. Your grill stays clean; the side quest gets its own window.
Fire and forget — something out of scope appears. Hand it to a fresh agent, return to the grill.
Grill → handoff → prototype → handoff back — validate an assumption in a throwaway session, bring the learning home.
The harder the questions, the more context a grill accumulates — and the more valuable it becomes to protect it. /handoff is how you keep going hard without drowning the session.
The deeper the grill, the more there is to lose by derailing it. A handoff lets you chase a tangent and come back to a session that’s still sharp.
Skills worth adding to the stack
/tdd — a red-green-refactor loop that builds one vertical slice at a time: failing test, minimal implementation, refactor. If your grill and PRD were thorough, the acceptance criteria already exist — this turns them into tests Claude has to satisfy before it can claim something works.
/zoom-out — when you land in an unfamiliar stretch of code and can’t see how it fits the bigger picture, this pulls the agent up a level: it explains the section’s role, its dependencies, and how it connects to the rest of the system. Useful when you inherit a module, come back to old code, or need the lay of the land before changing anything.
/grill-with-docs — a /grill-me variant that tests your plan against the existing domain model and codebase. The right choice when you’re extending a live system rather than starting fresh; it ties new decisions to what already exists and updates context docs inline.
Beyond the grill: getting Claude Code to run lean
The grill is the star of the workflow, but it only pays off if the rest of your setup isn’t quietly bleeding tokens and context. These are the habits that keep Claude Code fast, cheap, and sharp around it.
Run grill sessions on the strongest model; implement on a cheaper one.
A grill is pure exploration — no file writes, no tool-call overhead — so a frontier model like Opus earns its cost here by asking sharper questions and catching contradictions a smaller model misses, and the token bill stays low because it’s just conversation. Once you cross into implementation — /afk , /tdd , execution — a model like Sonnet does the heavy lifting at a fraction of the price.
Spend your best model where decisions are made, not where code is typed.
.claudeignore cuts your context bill in one edit.
By default Claude pulls everything in your project into scope. A .claudeignore file (same syntax as .gitignore ) stops the dead weight from auto-loading — excluding .next/ alone can trim 30–40% off context in a Next.js repo. It’s the highest-leverage two minutes you’ll spend.
node_modules/
dist/
build/
.next/
__pycache__/
*.lock
.git/
*.db
This doesn’t stop Claude from reading these files when you explicitly ask — it just keeps them out of automatic exploration.
Keep files under 200–300 lines.
Agents read whole files. When one reads a 1,500-line file to find a 20-line function, it pays for 1,480 lines of noise — and reasons worse for it. This isn’t a prompt trick; it’s a codebase discipline that compounds across every read in a session.
A large file is a tax every agent pays, every time it looks.
Teach Claude to grep, not read — in CLAUDE.md .
CLAUDE.md is the one file Claude reads at the start of every session, so it’s where your project’s standing rules live — stack, conventions, and the things you never want to re-explain. The highest-value rule you can put there: tell it which files to search instead of read . We have a translation-keys file with over 2,000 entries; reading it whole would torch the context budget on noise. So CLAUDE.md says:
grep in `packages/survey-type-defs/src/appTranslation/AppTranslationKeyEnum.ts`
— 2000+ keys, never read the full file
One line, and Claude stops pulling a giant file into context every time it needs a single key. Do this for every oversized-but-unavoidable file in your repo.
Disable MCP servers you aren’t using.
Historically, Claude Code would load MCP tool definitions directly into context, and a few heavy servers (Playwright, GitHub, Gmail, etc.) could quietly consume thousands of tokens each — and a sprawling setup, tens of thousands per turn. Recent Claude Code versions mitigate this with Tool Search, now enabled by default: tools are discovered and loaded on demand instead of shipping every schema upfront. It helps, but large MCP setups still add overhead and complexity.
Run /context to inspect what each server is costing you, and prune unused servers with /mcp .
Compact early — around half-full, not when it breaks.
Quality degrades as the window fills; by the time responses feel hazy, you’re already deep in the rot zone. The practical sweet spot is compacting around 50–60% capacity, while Claude still has full, uncompressed context to summarize from. Waiting until 90% means you’re summarizing an already-degraded view.
Git worktrees for parallel agents.
Independent features don’t need to run one after another. Worktrees give each agent its own branch and isolated directory off the same repo, so three agents can do in an hour what would take three sequentially — no context collisions, no merge chaos.
Thinking is the new bottleneck
The shape of building software has flipped.
Writing code used to be the bottleneck — now it’s the part agents are best at. What’s left for us is deciding what’s actually worth building and being precise enough that the implementation doesn’t drift.
That’s why the highest-leverage skill in this workflow is /grill-me — a skill that generates no code at all, yet creates more leverage than anything else.
The leverage has moved upstream. Specificity, edge-case honesty, and tight context management matter more than raw output.
Everything else — lean context, smaller files, the right model for the right phase — is just there to protect that thinking from noise.
The grill comes first. The code comes second.
Tags:
ai-coding ,
claude code ,
context management ,
developer experience ,
grill-me ,
productivity ,
skills ,
workflow
Categories:
AI & Machine Learning ,
Developer Tools ,
Engineering
Rewrites Don’t Fail at Release. They Fail in What Happens After.
Every rewrite inherits invisible debt — edge cases the old system silently handled for years, unknown to anyone, until they break.
Truncating LLM Context Is Not Retrieval: Why slice(0, N) Was the Wrong Fix
We capped a 2.6M-character LLM prompt with slice(0, 10000). Latency dropped, dashboards turned green, and the model started answering with the wrong context....
One Nav at a Time: How We Stopped Feeding the Legacy and Started Replacing It
The Strangler Fig — We Didn’t Rewrite It. We Outgrew It.
Limits: The Engineering Decision You Keep Postponing—Until Production Makes It For You
What running at scale taught us about the limits we didn’t set.
