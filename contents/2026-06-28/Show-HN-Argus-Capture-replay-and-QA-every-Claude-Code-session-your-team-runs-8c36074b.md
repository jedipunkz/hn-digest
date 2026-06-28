---
source: "https://www.arguslab.co/"
hn_url: "https://news.ycombinator.com/item?id=48708829"
title: "Show HN: Argus – Capture, replay and QA every Claude Code session your team runs"
article_title: "Argus — Quality & replay for Claude Cowork"
author: "zamtam"
captured_at: "2026-06-28T17:26:39Z"
capture_tool: "hn-digest"
hn_id: 48708829
score: 1
comments: 0
posted_at: "2026-06-28T16:26:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Argus – Capture, replay and QA every Claude Code session your team runs

- HN: [48708829](https://news.ycombinator.com/item?id=48708829)
- Source: [www.arguslab.co](https://www.arguslab.co/)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T16:26:46Z

## Translation

タイトル: Show HN: Argus – チームが実行するすべてのクロード コード セッションをキャプチャ、再生、QA
記事のタイトル: Argus — クロード・コワークの品質とリプレイ
説明: Argus は、すべての Claude Cowork セッションをキャプチャ、再生、QA レビューします。これは、Cowork ロールアウトに欠落している監査証跡と可観測性レイヤーです。

記事本文:
Argus メソッド ビュー サンプル プライバシー 対象者 学習 比較 サインイン アルファ版がオープンしています · サインインして開始してください Claude Cowork は実際にユーザーやクライアントのために機能していますか?
Argus は、チームまたは組織全体のユーザーのすべてのセッションをキャプチャし、一度に何千ものセッションを読み取り、スキルが保持されている場所、静かに破壊されている場所、そして次に構築すべきものを指し示しているパターンを明らかにします。
サンプル セッションを読む Magic-link サインイン。パスワードなし。アルファ期間中は無料。
ライブ プレビュー · Argus ダッシュボード、1 つのワークスペース
Claude Cowork の組み込みテレメトリは、スキルが呼び出されたことを示します。それが機能したかどうか、ユーザーが回答を受け入れたかどうか、明日修正版を出荷する必要があるかどうかはわかりません。
何も言わないカウンター。
Cowork の組み込みテレメトリは、スキルを Skill: 3 として記録します。 3回の呼び出し。今月 1 人のクライアントに対する 412 セッションにわたって、スキルごとに 1 つの数値 (呼び出し回数) が得られます。どのバージョンが実行されたか、何が返されたか、ユーザーが回答を受け入れたのか、2 回質問する必要があったのかについては何もありません。
週次レビューを 4 週間前に発送しました。最初の 38 セッションを通じて、ユーザーは最初のターンで回答を受け入れました。次の 9 回にわたって、彼らはツールを再質問し、言い換え、切り替えました。セッション 39 で何かが失敗し始めました。誰も気づきませんでした。コスト ラインは横ばいのままで、単独で壊れているように見えるセッションは 1 つもありませんでした。
まだ無かったスキル。
今月、チームのトラフィック全体で、「このリニアのチケットをリリースノートのエントリに変える」という言葉が 8 つの異なる表現で 14 回出てきました。それぞれが異なるその場限りの答えを受け取りました。あるユーザーは諦めました。そこにはスキルが書かれるのを待っています。あなたや他の誰かのテレメトリ サーフェスでは、決してそれを見つけることはできません。
Argus は Claude Cowork プラグインとして実行されます。すべてのセッション、つまりプロンプト、

継続的な応答、すべてのツール呼び出し、すべてのフォローアップがプレーンテキストで、ユーザーが実際に行った会話にステッチバックされます。他のすべてを可能にする定性層。
ユーザーは回答を受け入れましたか?
スキルが機能すれば会話は終わります。何度も促されたり、言い換えられたり、放棄されたりしないスキル。 Argus はユーザーの正確なフォローアップをキャプチャするため、エージェントは数百のセッションにわたって、どのスキルのどのバージョンが最初のターンに着地し、どのバージョンがそうでないかを一目で確認できます。
捕獲・使用：初ターン受付、追撃パターン
アシスタントはユーザーにその仕事をするように依頼しましたか?
スキルは、新しい質問をするのではなく、質問に答える必要があります。スキルの指定が不十分な場合、アシスタントは「明確にしていただけますか…」「どっちのことを言いたかったのですか…」と停滞し、ユーザーはスキルが行う予定だった作業を実行します。 Argus はストールをキャプチャし、エージェントがストールがどこに集まっているかを表示できるようにします。
キャプチャされ、次の用途で使用されます: ストール頻度、指定不足のプロンプト
ツールは実際に何を返しましたか?
MCP 呼び出しは成功しました。ステータス 200。しかし、ペイロードは空、または 400 行のダンプ、またはスキルが要求したものと一致しない JSON でした。 Argus はアシスタントの応答の横にあるツールのプレーンテキスト出力をキャプチャするため、エージェントはスキルが不正な入力を続けたセッションにフラグを立てることができます。
キャプチャされ、次の用途で使用されます: ツールと出力の不一致、サイレントエラー
ユーザーが求めた、どのスキルでも対応できないものは何でしょうか?
カスタマイズでまだカバーされていないプロンプト。同じエクスポート、同じルックアップ、同じラングル - たとえ何も答えがなかったとしても、そのままキャプチャされました。エージェントは、満たされていないプロンプトのコーパスを読み取り、次のスキルとなる準備ができているパターンを明らかにします。
捕捉・使用: アンメット・プロンプト・クラスタリング、スキル候補
すべてのセッションから、カタログ自体が改善されます。
5 つの手、下部を読む

m-up — 基礎に生の仕事、その上に洗練された知識。各レイヤーはその下のレイヤーの上に置かれます。このループにより、新しく洗練されたスキルが定着し、次のセッションに戻ります。
作品を評価してください。エージェントは弱いスキルを磨き、使用状況が求めるスキルをドラフトします。
スキルごとにグループ化されたリプレイ。組織全体のすべての呼び出しを分析します。
編集され、テナントが分離された後、各セッションは完全な構造化されたレコード、つまり詳細を可能にするデータ モデルに再構築されます。
プラグイン フックと OpenTelemetry は、すべてのプロンプト、ツール呼び出し、応答、メトリクスをソースでキャプチャします。
人々は毎日の Cowork セッションを実行します。これがすべての基礎となります。
Argus エージェントは、同じスキル、MCP、またはエージェントの何千ものセッションを読み取り、次の動きを伝えます。 Argus Web アプリで利用できます。 Claude Cowork プラグインとして提供され、作業セッション中に呼び出すことができます。
すべての Cowork セッションがキャプチャされ、インデックス付けされます。ユーザー、プロジェクト、スキル、またはステータスでフィルターします。任意のターンに注釈を付けます。スキルが何を生成したかをバージョンごとに比較します。
バージョンごとの、時間の経過に伴う最初のターンの受け入れ、フォローアップ率、ストール頻度、ツール出力の不一致。何かが漂い始めた週には、単一のセッションが破綻したように見える前にチャートが曲がります。
スキルごと、バージョンごと、クライアントごと
エージェントは、スキルの失敗した洗練されたセッションを読み取り、SKILL.md に対する具体的な編集を提案します。つまり、より鋭いトリガー、欠落しているツール、失敗を捕らえた例などです。あなたはパッチを承認します。新しいバージョンはプル リクエストとしてマーケットプレイスに登場します。
同じエージェントが、ワークスペース全体で満たされていないプロンプトをクラスター化し、候補者のスキルをドラフトします。チームが体系的であると認識していなかった繰り返しの質問です。推測が少なくなります。さらに出荷されました。
繰り返しの質問からスキルのドラフトまで
Argus エージェント自体がクロード カウです

orkプラグイン
これは、プラットフォーム上の他のすべてのスキルと同じ方法で実行されます。それはそれ自体を捕らえます。それ自体を見直します。それ自体が洗練されます。コンサルタントが出荷しているものは、コンサルタントが出荷しているレールの上を走っています。
作業場は作業場を利用します。
前方展開エンジニアのポートフォリオからの 1 つの実際のセッション。名前は編集されました。これはエージェントが読み取る内容です。
HubSpot と Stripe から第 2 四半期の解約を取得し、Intercom の [編集された顧客] からの NPS スコアと相互参照し、社内の週次レビューの声で解約要因のトップ 3 を書き留めます。
Stripe からキャンセルの理由を取得し、HubSpot からクローズド紛失メモを並行して取得し、同じウィンドウで Intercom が記録した NPS 応答と照合します。 3 つの Explore サブエージェントを開始 — Stripe はレート制限があるため、反復されます。
Argus は、カスタマイズで実行された会話をキャプチャします。私たちは彼らに対して注意しなければなりません。私たちが先に進むことのできない 4 つの約束。
№ 01 秘密は決して機械から離れることはありません。
キャプチャ プラグインは、エンベロープがユーザーのコンピュータから送信される前に、API キー、OAuth トークン、ベアラー ヘッダー、一般的なパスワード パターンをソースでスクラブします。 Anthropic、OpenAI、Supabase、GitHub、AWS、Slack のパターンがデフォルトでキャッチされます。自分で追加することもできます。
構築済み、プラグイン側、転送前
№ 02 一言でセッションが非公開になります。
いつでも Cowork に「/private」と入力すると、プラグインはそのセッションのキャプチャを停止し、すでに出荷されているものはすべて削除します。避難ハッチはユーザーのものであり、代理店のものではありません。サポートチケットも管理者の承認もありません。
№ 03 レビューの前に名前を編集します。
電子メール、名前、カスタム正規表現のワークスペースごとのパターンは、永続化される前に、キャプチャされたすべてのエンベロープで実行されます。レビュー担当者には、会社名ではなく [redacted-customer] が表示されます。
№ 04 暗号化され、隔離され、決してトレーニングされていません。
インジェストワーカーへの TLS、

データベース内に保存される AES-256。クエリごとに行レベルのセキュリティによってワークスペースが分離されます。 Argus は、キャプチャしたセッションを使用して、私たちのモデル、Anthropic のモデル、または他の人のモデルをトレーニングすることは決してありません。
永続的なポリシー · データベースで適用される
キャプチャされた同じセッションは 3 つの方法で読み取られます。各部屋には必要なものが表示され、必要のないものは何も表示されません。
カスタム スキルを複数のクライアントに配布します。どのバージョンがどのクライアントで動作しているのか、サポート チケットをどこで取得しようとしているのか、そしてどのユーザー プロンプトが次にビルドするものを指しているのかを知る必要があります。
・満たされていないプロンプトから次のスキルをドラフトする
クロード・コワークを大規模に展開
あなたは組織の MCP サーバーを標準化し、最初の社内スキルを公開しました。リーダーによるレビューの前に、200 人のエンジニア全体でどのパターンが機能し、どのパターンが機能しないのかを知る必要があります。
· 新しいスキルバージョンの QA ゲート
あなたは、チームの Cowork インストールが使用されていること、委託したスキルが成果を上げていることを知りたいと考えています。また、Slack スレッドではなく、よく整理された 1 つの概要を確認したいと考えています。
· 自分のプロジェクトのみを対象とする
2026 年 6 月時点での Argus の位置。プライベート ベータ版が進むにつれて数値が更新されます。
2 つの代理店、1 つの社内 IT チーム、1 人の個人コンサルタント。 3 つの視聴者タイプすべてをカバーします。
パイロット チーム全体で実行されているライブ プラグインから。
安定したプラグイン、MCP フレンドリーなスキル カタログ、バージョン違いの QA — それがハードルです。
サインインし、ワークスペースを作成し、プラグインを Claude Cowork にドロップします。コールドタブから最初にキャプチャされたセッションまで 5 分です。アルファ期間中は無料で、カードは必要ありません。
マジックリンク認証。ワンタイムリンクを送信するので、クリックすればすぐにアクセスできます。パスワードを覚えたり、サインアップフォームに記入したりする必要はありません。
ワークスペース名を選択し、プライバシーのデフォルトを選択します。
事前設定されたプラグイン バンドルをダウンロードし、Cowork にドロップします。
新鮮な牛を開く

ork セッション、こんにちは — 最初のキャプチャは数秒で完了します。
私たちはあなたのデータを販売したり、それを使ってモデルをトレーニングしたりすることはありません。また、いつでもセッションをオプトアウトできます。
クロードの上で行われる作業のレビュー層。 Web アプリ、キャプチャ プラグイン、およびエージェント - 他の人の組織内で AI を提供する人々向け。

## Original Extract

Argus captures, replays, and QA-reviews every Claude Cowork session — the missing audit trail and observability layer for Cowork rollouts.

A Argus Method Views Sample Privacy Audience Learn Compare Sign in Alpha is open · sign in to start Is Claude Cowork actually working for your users and clients?
Argus captures every session of your users across your team or organization, then reads across thousands of them at once — surfacing where skills are holding, where they're quietly breaking, and which patterns are pointing at the next thing you should build.
Read a sample session Magic-link sign in. No password. Free during alpha.
Live preview · Argus dashboard, one workspace
Claude Cowork's built-in telemetry tells you a skill was invoked. It can't tell you whether it worked, whether the user took the answer, whether you should ship a fix tomorrow.
The counter that says nothing.
Cowork's built-in telemetry logs your skill as Skill: 3 . Three invocations. Across 412 sessions for one client this month, you have a single number per skill — invocation count. Nothing about which versions ran, what they returned, whether the user accepted the answer or had to ask twice.
You shipped weekly-review four weeks ago. Across the first 38 sessions, users accepted the answer on the first turn. Across the next 9, they re-asked, rephrased, switched tools. Something started failing on session 39. Nobody noticed — the cost line stayed flat and no single session looked broken on its own.
The skill that wasn't there yet.
Across the team's traffic this month, “turn this Linear ticket into a release-notes entry” came up fourteen times in eight different phrasings. Each one got a different ad-hoc answer; one user gave up. A skill is waiting to be written there. No telemetry surface — yours or anyone else's — will ever find it.
Argus runs as a Claude Cowork plugin. It captures every session — prompts, assistant responses, every tool call, every follow-up — in plain text, stitched back into the conversation the user actually had. The qualitative layer that makes everything else possible.
Did the user accept the answer?
A skill that works ends the conversation. A skill that doesn't gets re-prompted, rephrased, abandoned. Argus captures the user's exact follow-ups so the Agent can see, at a glance across hundreds of sessions, which versions of which skill are landing on the first turn — and which aren't.
Captured · used in: first-turn acceptance, follow-up patterns
Did the assistant ask the user to do its job?
Skills should answer questions, not ask new ones. When a skill is under-specified, the assistant stalls — “could you clarify…” , “which one did you mean…” — and the user does the work the skill was meant to do. Argus captures the stalls so the Agent can show where they cluster.
Captured · used in: stall frequency, under-specified prompts
What did the tool actually return?
The MCP call succeeded. Status 200. But the payload was empty, or a 400-row dump, or a JSON that didn't match what the skill asked for. Argus captures the tool's plain-text output beside the assistant's response, so the Agent can flag the sessions where the skill kept going on bad input.
Captured · used in: tool-output mismatches, silent failures
What did users ask for that no skill could handle?
The prompts your customisations don't yet cover. Same export, same lookup, same wrangle — captured verbatim, even when nothing answered them. The Agent reads across the unmet-prompts corpus and surfaces patterns ready to become the next skill.
Captured · used in: unmet-prompt clustering, skill candidates
From every session, a catalogue that improves itself.
Five moves, read bottom-up — raw work at the foundation, refined knowledge on top. Each layer rests on the one beneath it; the loop settles new and refined skills back into the next session.
Rate the work; an agent refines weak skills and drafts the ones your usage is asking for.
Replay grouped by skill; analyse every invocation across the org.
Redacted and tenant-isolated, then each session is rebuilt into a complete, structured record — the data model that makes the detail possible.
Plugin hooks and OpenTelemetry capture every prompt, tool call, reply and metric — at the source.
People run their everyday Cowork sessions — the foundation everything rests on.
The Argus Agent reads across thousands of sessions of the same skill, MCP, or agent — and tells you the next move. Available in the Argus web app. Coming as a Claude Cowork plugin you can invoke during a working session.
Every Cowork session captured and indexed. Filter by user, project, skill, or status. Annotate any turn. Diff version-by-version what a skill produced.
First-turn acceptance · follow-up rate · stall frequency · tool-output mismatches, per version, over time. The week something starts drifting, the chart bends before any single session looks broken.
Per skill · per version · per client
The Agent reads across a skill's failed and refined sessions and proposes concrete edits to the SKILL.md — sharper trigger, missing tool, an example that would have caught the failure. You approve the patch; the new version lands in the marketplace as a pull request.
The same Agent clusters the unmet prompts across the workspace and drafts candidate skills — the recurring asks your team didn't realise were systematic. Less guesswork. More shipped.
From recurring asks to draft skills
The Argus Agent is itself a Claude Cowork plugin
It runs the same way every other skill on the platform runs. It captures itself. It reviews itself. It refines itself. The thing the consultant is shipping is running on the rails the consultant is shipping.
The workshop uses the workshop.
One real session from a forward-deployed engineer's portfolio. Names redacted. This is what the Agent reads.
Pull Q2 churn from HubSpot and Stripe, cross-reference with NPS scores from [redacted-customer] in Intercom, and write up the top three churn drivers in our internal weekly-review voice.
I'll pull the cancellation reasons from Stripe and the closed-lost notes from HubSpot in parallel, then join them against the NPS responses Intercom logged in the same window. Kicking off three Explore subagents — Stripe's rate-limited so it'll iterate.
Argus captures the conversations your customisations run. We have to be careful with them. Four commitments we won't move on.
№ 01 Secrets never leave the machine.
The capture plugin scrubs API keys, OAuth tokens, Bearer headers, and common password patterns at the source — before the envelope leaves the user's computer. Anthropic, OpenAI, Supabase, GitHub, AWS, Slack patterns are caught by default; you can add your own.
Built · plugin-side · pre-transit
№ 02 One word makes a session private.
Type /private in Cowork at any point and the plugin stops capturing that session — and deletes anything already shipped. The escape hatch is the user's, not the agency's. No support ticket, no admin approval.
№ 03 Redact names before review.
Per-workspace patterns for emails, names, and custom regex run on every captured envelope before it's persisted. Reviewers see [redacted-customer] , not the company name.
№ 04 Encrypted, isolated, never trained on.
TLS to the ingestion worker, AES-256 at rest in the database, workspace-isolated by row-level security on every query. Argus never uses your captured sessions to train any model — ours, Anthropic's, or anyone else's.
Standing policy · enforced at the database
The same captured session is read three ways. Each room sees what it needs to and nothing it doesn't.
You ship custom skills to several clients. You need to know which versions are working at which client, where you're about to get a support ticket, and which user prompts are pointing at the next thing to build.
· Draft the next skill from unmet prompts
Rolling Claude Cowork out at scale
You're standardising your org's MCP servers, you've published your first internal skills, and you need to know — across 200 engineers — which patterns work and which don't, before the leadership review.
· QA gates for new skill versions
You want to know your team's Cowork install is being used, that the skills you commissioned are landing, and you'd like to see one well-organised summary instead of a Slack thread.
· Scoped to your projects only
Where Argus stands as of June 2026. Numbers update as the private beta rolls forward.
Two agencies, one internal IT team, one solo consultant. Coverage of all three audience types.
From the live plugin running across pilot teams.
Stable plugin, MCP-friendly skill catalog, version-diff QA — that's the bar.
Sign in, create your workspace, drop the plugin into Claude Cowork — five minutes from a cold tab to your first captured session. Free during alpha, no card needed.
Magic-link auth. We send a one-time link, you click it, you're in. No password to remember, no signup form to fill.
Pick a workspace name, choose privacy defaults.
Download a pre-configured plugin bundle, drop it into Cowork.
Open a fresh Cowork session, say hi — the first capture lands in seconds.
We don't sell your data, we don't train models on it, and you can opt any session out at any time.
The review layer for the work that happens on top of Claude. A web app, a capture plugin, and an Agent — for the people delivering AI inside other people's organisations.
