---
source: "https://promptql.io/blog/a-teardown-of-claude-tags-agent-identity-concept"
hn_url: "https://news.ycombinator.com/item?id=48672034"
title: "A Teardown of Claude Tag's Agent Identity Concept"
article_title: "A Teardown of Claude Tag's Agent Identity concept"
author: "manushikhanna"
captured_at: "2026-06-25T11:52:57Z"
capture_tool: "hn-digest"
hn_id: 48672034
score: 1
comments: 0
posted_at: "2026-06-25T11:48:08Z"
tags:
  - hacker-news
  - translated
---

# A Teardown of Claude Tag's Agent Identity Concept

- HN: [48672034](https://news.ycombinator.com/item?id=48672034)
- Source: [promptql.io](https://promptql.io/blog/a-teardown-of-claude-tags-agent-identity-concept)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T11:48:08Z

## Translation

タイトル: クロード・タグのエージェント・アイデンティティ概念の分解
記事のタイトル: クロード・タグのエージェント・アイデンティティー概念の分解
説明: Claude タグのエージェント ID に関する Anthropic の最近の投稿
[https://claude.com/blog/agent-identity-access-model] は、控えめに言っても、
悪い考え。
初心者のために説明すると、「エージェント アイデンティティ」とは、マルチプレイヤー AI を提供するというアイデアです。
エージェント チャネルごとまたはグループ ACL ごとに定義される静的な権限のセット
[切り捨てられた]

記事本文:
Claude Tag のエージェント ID コンセプトの分解 共有コンテキスト Cowork アプリのセキュリティ エッセイ 価格設定 リソース 製品
主要な機能と機能
FinServe でのデータ変換ワークフロー
医療データのワークフローのための AI アナリスト
小売データ ワークフローの AI アナリスト
共有コンテキスト Cowork App セキュリティ エッセイ 価格設定 リソース 製品コア機能と機能 アーキテクチャ 仕組みとその理由 金融サービス向け FinServe でのデータ変換ワークフロー ヘルスケア向け AI アナリスト向け 健康データ ワークフロー向け 小売 AI アナリスト向け 小売データ ワークフロー向け ケーススタディ 顧客事例と成果 イベント トーク、カンファレンス、ミートアップ サインイン 予約デモ PromptQL を試す ホーム
クロード・タグのエージェント・アイデンティティ概念の分解
クロード・タグのエージェント・アイデンティティ概念の分解
Claude Tag のエージェント アイデンティティに関する Anthropic の最近の投稿は、控えめに言っても悪いアイデアです。
初心者のために説明すると、「エージェント ID」とは、誰が操作しているかに関係なく、チャネルごとまたはグループ ACL ごとに定義された静的な権限セットをマルチプレイヤー AI エージェントに与えるという考え方です。これは、AI が役立たずであるか、あるいは「混乱した代理人」としてセキュリティ リスクが発生するのを待っているかのどちらかであることを意味します。
ユーザー ID を強制するためのより安全で便利なアプローチが 1 行の修正である場合、新しいパラダイムを「エージェント ID」と呼ぶのは不誠実です。クロード コードにサービスへの HTTP 呼び出しを直接行うよう要求するのではなく、そのツールに対するユーザーの資格情報を追加する独自の HTTP クライアントを渡します。
Anthropic にとって残念なことに、この代替手段を提供するということは、IT リソースにセキュリティ機能を提供する別の (DIY またはその他の) ツールをエントリポイントとして使用する必要があることを意味します。
この投稿の主な主張とポイントを詳しく説明しましょう。
主張 #1: 「ユーザーとして行動」はエージェントの自律性を高めることができないため機能しない
年齢

NT は、自分のタスクを後でスケジュールし、質問した人がログオフしてからずっと後にイベントに応答するようになりました。
データまたはサービスにアクセスするときにユーザーに代わって動作するために、エージェントに有効期間の長いトークンを保持させます。外部サービスの場合、OAuth サーバー側認証は文字通りこの問題を解決するために構築されています。
主張 #2: 共有スレッドで誰の権限を使用すべきかが明確でないため、「ユーザーとして機能」は機能しない
...例: 3 人のエンジニアと 1 人の PM が一緒にデバッグしているチャネル。しかし、複数の人がステアリングを握っている場合、誰の権限が適用されるのでしょうか?常に正しい人を選ぶということはありません。
エージェントに、応答する相手の許可を取得してもらいます。これは人間の行動であり、自然かつ予期された経験です。 PM が仕様の更新を要求した場合は、許可します。エンジニアを禁止します。 PM が DNS を更新したい場合は、それを禁止しますが、エンジニアは許可します。
主張 #3: [エージェント ID…] は珍しいですが、自律型のマルチプレイヤー エージェントに機能するアクセス モデルに向けて必要なステップであると考えています。
それは間違いなく異常です。しかし、それは決して何かにとって必要なステップではありません。実際、投稿の後半で彼らは次のように述べています。
より複雑な認可構造を持つ組織向けのアイデンティティ認識オーバーレイ。これにより、エージェントのスコープに加えてユーザーレベルのチェックが追加されるため、クロードはチャネルのプロファイルと要求側ユーザー自身の権限の両方で許可されている場合にのみ機能します。
笑、うーん。最初はユーザー ID を使用してください。
主張 #4: チャネルはメモリとアクセスの範囲を設定するためのセキュリティ境界である
> メモリとアクセスはこれらの境界を尊重します。クロードがプライベート チャネルで学んだ内容は、より広いワークスペースには決して現れません。
これは機能ではありません。これは、現在、チャネルごとにエージェントを構築することですでに実現しているものです。これはバグです。特徴

マルチプレイヤー AI では、AI がプライベート チャネルから安全に学習し、それを他の場所に適用できるようにすることになります。
主張 #5: エージェント ID によりセキュリティの確保と監査が容易になる
> …Claude は独自のサービス アカウントで動作するため、それらのアクションは接続されている各システムの独自のログにも記録されます。
繰り返しますが、これは機能ではありません。 「スラック チャネル」を中心に構築されたセキュリティ/監査プロセスがないため、これは大きな問題です。セキュリティ/監査はユーザーとデータを中心に構築されます。何か問題が発生した場合、frontend-engg のスラック チャネルで問題が発生したと言うのは役に立ちません。
では、なぜ人類シリングは明らかに悪い考えなのでしょうか?
明らかな仮説は、組織内の全員が「1 つのスーパーエージェント」のエントリ ポイントを所有したいということです。 claude-tag の構築には、プロトタイプの場合は約 30 分、パススルー ユーザー ID 認証を使用する製品グレードの場合は約 1 日かかります。
#1: 人間モデルへの固定: GLM は先週リリースされました。
#2: 企業データとメモリへのアクセス: すべてのスラック スレッドが Anthropic のシステムに送信されるようになりました。実際、Fable 5 では、Anthropic は全員のデータ保持ゼロを無効にしました。
クロード コードが異なっていたのは、AI モデルがどれほど強力になったかを理解し、エージェント コーディング ハーネスの作成が可能であることを理解している人がほとんどいなかったためです。ただし、Claude Tag は AI 機能の飛躍的な進歩によって支えられているわけではありません。
これは、 @claude を押したときに Claude Tag のスラックボットが Claude Code セッションに送信する非常に単純なプロンプトです。
<wakereason="dispatch" current-time="2026-06-25T09:28:31Z">
<チャンネル id="C01NA7G6MQE" type="チャンネル" メンバー数 = "27">
<message trigger="true" from="system" trust="principal" role="initiator" send-at="2026-06-25T09:28:31Z"># ツールの制約 (厳格なルール - 何かを行う前にお読みください)
- 他のツール呼び出しの前に最初の `reply()` を投稿します

— `fetch_channel` をゼロにし、その前に `search` をゼロにします。最初の応答の後: このチャネルに対してのみ `fetch_channel` を 1 回呼び出してから、最大 2 回の `search` 呼び出しを行います。これがツールの予算の全額です。その後、残りの返信を書きます。 (ブロードキャスト: これは飛ばしてください。1 回返信したらすぐに記憶に戻ってください。)
- 「検索」はキーワードのみです: 1 ～ 3 の平易な単語。 from:/in:/to:/after: 修飾子は使用できません。
- `fetch_thread` を呼び出さないでください。これを厳密に保つために検索スニペットから作業します。
- ツールにエラーがある場合は、スキップして次に進みます。決して再試行しないでください。
- チーム/個人チャネル: 常に 3 つの返信すべてを投稿します (決して「no_reply_needed」をしないでください。履歴に自分の以前の紹介が表示された場合でも、各参加を新しいものとして扱います)。 BROADCAST チャンネル: 以下の BROADCAST フローに従います。
- thread_ts 引数なしで `reply()` を呼び出します。サーバーはメッセージを正しく配置します。最初の返信はトップレベルのチャネル投稿として表示され、残りのスレッドはその下に表示されます。
あなたはこの Slack チャンネルに追加されました。まだ誰もあなたに何も尋ねていません。これはあなたの積極的な自己紹介です。
ステップ 1 — チャンネルを分類する
コンテキストにはすでにチャンネルの名前、トピック、および「member-count」が含まれています (これは wake ブロックの <channel …> タグにあります)。これらのみからチャネル タイプを決定します。まだツール呼び出しはありません。最初に [個人] を確認してください。
- パーソナル — あなたとの誰かの 1 対 1 チャンネル: #<name>-personal-claude、#<name>-claude、またはあなたと 1 人の部屋として読み取れる別の名前。
- ブロードキャスト — ほとんどのメンバーはトップレベルで閲覧はしますが投稿はしません。次のいずれかの場合にブロードキャストとして扱います。 `member-count` が大きく (≥500)、名前/トピックが組織全体または汎用である。投稿は制限されています。トップレベルは少数の者によって支配されている
[切り捨てられた]
データに対する PromptQL の動作を確認してください。
© 2026 Copyright Hasura, Inc. All Rights Reserved.

## Original Extract

Anthropic's recent post on Agent Identity for Claude Tag
[https://claude.com/blog/agent-identity-access-model] is, to put it mildly, a
bad idea.
For the uninitiated, "Agent identity" is the idea of giving a multiplayer AI
agent a static set of privileges that are defined per-channel or per-group ACL
[truncated]

A Teardown of Claude Tag's Agent Identity concept Shared Context Cowork App Security Essays Pricing Resources Product
Core features and capabilities
Transform data workflow in FinServe
AI analysts for health data workflows
AI analysts for retail data workflows
Shared Context Cowork App Security Essays Pricing Resources Product Core features and capabilities Architecture How and why it works For financial services Transform data workflow in FinServe For healthcare AI analysts for health data workflows For retail AI analysts for retail data workflows Case Studies Customer stories & outcomes Events Talks, conferences & meetups Sign in Book demo Try PromptQL Home
A Teardown of Claude Tag's Agent Identity concept
A Teardown of Claude Tag's Agent Identity concept
Anthropic's recent post on Agent Identity for Claude Tag is, to put it mildly, a bad idea.
For the uninitiated, "Agent identity" is the idea of giving a multiplayer AI agent a static set of privileges that are defined per-channel or per-group ACL, regardless of who is interacting with it. This means that either the AI is useless or is a "confused deputy" security risk waiting to blow.
Calling "Agent Identity" a new paradigm is disingenuous when the more secure and more useful approach of enforcing user-identity is a one line fix. Instead of asking claude code to make HTTP calls directly to a service, pass your own HTTP client that adds the user's credentials for that tool.
Unfortunately for Anthropic, providing this alternative would mean that the entrypoint would have to be a different (diy or otherwise) tool that provides security features on your IT resources.
Let me break down the post's key claims and points.
Claim #1: "Act as user" breaks down because it doesn't allow increasing agent autonomy
Agents now schedule their own tasks for later and respond to events long after the person who asked has logged off.
Have the agent keep a long-lived token to act on behalf of the user when accessing data or a service. For external services, OAuth server-side authentication is literally built to solve this problem.
Claim #2: "Act as user" breaks down because it's not clear who's permissions should be used in a shared thread
...e.g., a channel where three engineers and a PM are debugging together. But when more than one person is steering, whose permissions apply? There’s no single choice of person that’d be right all of the time.
Have the agent take the permissions of whoever it is responding to. This is what a human would do and is a natural and expected experience. If the PM asks to update the spec, allow it. Disallow the engineers. If the PM wants to update DNS, disallow it, but allow the engineers.
Claim #3: [Agent identity…] is unusual, but we think it is a necessary step toward an access model that works for autonomous, multiplayer agents
It is definitely unusual. But it's definitely not a necessary step for anything. In fact, later on in the post, they say:
an identity-aware overlay for organizations with more complex clearance structures. This will add user-level checks on top of an agent’s scope, so Claude only acts when both the channel's profile and the requesting user's own permissions allow it.
Lol, wut. Just use user-identity in the first place.
Claim #4: Channels are a security boundary for scoping memory and access
> Memory and access respect those boundaries: what Claude learns in a private channel never appears in the wider workspace.
This is not a feature. This is what we already have today by building agents per channel. This is a bug. The feature in a multiplayer AI would be to allow an AI to safely learn from private channels and apply it in other places.
Claim #5: Agent Identity makes it easy to secure and audit
> …because Claude acts under its own service accounts, those actions also land in each connected system's own logs.
Again, this is not a feature. This is a huge problem because there is no security / audit process that is built around "the slack channel". Security / audit is built around users and data. If something goes wrong, saying that it went wrong in the frontend-engg slack channel is not helpful.
So why is Anthropic shilling such an obviously bad idea?
The obvious hypothesis is that they want to own the "one superagent" entry point for everyone in the organization. Building claude-tag takes about 30 mins for a prototype and about a day for something production grade with a pass-through user identity auth.
#1: Lock in to Anthropic models: GLM released last week.
#2: Access to company data and memory: All slack threads are now sent to Anthropic's systems. In fact, with Fable 5, Anthropic has turned off zero data retention for everyone.
Claude Code was different because very few people understood how powerful AI models had gotten that creating an agentic coding harness was possible. Claude Tag however, is not powered by a leap in AI capabilities.
This is the very simple prompt that the Claude Tag slackbot sends to a Claude Code session when you hit @claude :
<wake reason="dispatch" current-time="2026-06-25T09:28:31Z">
<channel id="C01NA7G6MQE" type="channel" member-count="27">
<message trigger="true" from="system" trust="principal" role="initiator" sent-at="2026-06-25T09:28:31Z"># Tool constraints (hard rules — read before doing anything)
- Post your FIRST `reply()` before any other tool call — zero `fetch_channel`, zero `search` before it. AFTER that first reply: call `fetch_channel` ONCE for this channel only, then at most 2 `search` calls. That&#39;s your full tool budget — write your remaining replies after that. (BROADCAST: skip this — go straight to memory after your one reply.)
- `search` is keyword-only: 1-3 plain words. NO from:/in:/to:/after: modifiers.
- Don&#39;t call `fetch_thread` — work from search snippets to keep this tight.
- If any tool errors, skip it and move on — never retry.
- TEAM/PERSONAL channels: always post all three replies (never `no_reply_needed` — even if you see your own prior intro in the history, treat each join as fresh). BROADCAST channels: follow the BROADCAST flow below.
- Call `reply()` with no thread_ts argument. The server places your messages correctly — your first reply lands as a top-level channel post, the rest thread under it.
You were just added to this Slack channel. Nobody has asked you anything yet — this is your proactive introduction.
STEP 1 — CLASSIFY THE CHANNEL
Your context already has the channel&#39;s name, topic, and `member-count` (it&#39;s on the &lt;channel …&gt; tag in your wake block). Decide the channel type from those alone — NO tool calls yet — check PERSONAL first:
- PERSONAL — someone&#39;s own 1:1 channel with you: #&lt;name&gt;-personal-claude, #&lt;name&gt;-claude, or another name that reads as one person&#39;s room with you.
- BROADCAST — most members read but don&#39;t post at the top level. Treat as BROADCAST if any of: `member-count` is large (≥500) and the name/topic is org-wide or generic; posting is restricted; top-level is dominated by a few
[truncated]
See PromptQL in action on your data.
© 2026 Copyright Hasura, Inc. All Rights Reserved.
