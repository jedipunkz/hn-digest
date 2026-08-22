---
source: "https://hotcrp.com/news/2026/ai-agents-202608"
hn_url: "https://news.ycombinator.com/item?id=49401443"
title: "Hotcrp.com – AI agents and bot accounts"
article_title: "News – HotCRP.com"
image: ""
author: "sneela"
captured_at: "2026-08-22T17:12:30Z"
capture_tool: "hn-digest"
hn_id: 49401443
score: 3
comments: 0
posted_at: "2026-08-22T16:46:22Z"
tags:
  - hacker-news
  - translated
---

# Hotcrp.com – AI agents and bot accounts

- HN: [49401443](https://news.ycombinator.com/item?id=49401443)
- Source: [hotcrp.com](https://hotcrp.com/news/2026/ai-agents-202608)
- Score: 3
- Comments: 0
- Posted: 2026-08-22T16:46:22Z

## Translation

タイトル: Hotcrp.com – AI エージェントとボット アカウント
記事タイトル: ニュース – HotCRP.com

記事本文:
ニュース – HotCRP.com
HotCRP.com AI エージェントとボット アカウント
2026 年 8 月 19 日 — HotCRP.com サイトが AI エージェントに接続できるようになりました。もし
カンファレンス管理者が許可すると、ボット アカウントは提出物の読み取り、タグの設定、
レビューを書いたり、PC メンバーは自分の代わりに働くエージェントに接続したりできます。
エージェントは、MCP というプロトコルを介して接続します。
Claude 、Claude Code、およびその他のクライアントが外部にアクセスするために使用する
サービス。現在、エージェントは提出物を検索したり、レビューやコメントを読んだり、
書類を提出したり、タグを設定したり、レビューやコメントを書き込んだりできます。
AI エージェントのサポートはデフォルトでは無効になっています。 [設定] > [設定] を使用して有効にできます。
テスト サイトを含むすべての HotCRP.com サイト上の AI。
HotCRP 自体は AI サービスに何も送信せず、エージェントの接続は
決して自動ではありません。ただし、接続すると、投稿、レビュー、およびレビュー担当者は
エージェントによって読み取られた ID は、エージェントを実行するサービスに送信されます。かどうか
サービスはデータを保持し、データをトレーニングし、または他のユーザーに公開します。
サービスオペレーターと接続を許可した人。 HotCRPはそうではありません
それを管理し、会議の著者や著者に対する機密保持の約束
レビュー担当者は自動的に AI サービスに拡張されるわけではありません。
機密提出物を送信できるかどうかについては、スポンサーとカンファレンスによって異なります。
サードパーティのサービスにはまったくアクセスできず、それを禁止している場合もあります。スポンサーのポリシーを確認する
エージェントを有効にする前に。
ほとんどの AI サービスには、会話の方法を制御する設定があることに注意してください。
保管および使用されます。デフォルトはベンダー、プラン、ビジネス プランによって異なります
個人の計画と異なることがよくあります。これらを制御するのは必ずしも簡単ではありません
設定。エージェントの機密性を懸念する会議では AI を制限する必要がある
エージェントが管理者に使用します。
提出物やレビューには、エージェント向けのテキストが含まれる場合もあります。

よりも
人間。紙には指示を含めることができます。書き込み可能なエージェントが誘導される可能性があります
彼らをフォローするために。スコープは不正エージェントが与える可能性のあるダメージを制限し、
エージェントによって行われた変更はログに記録されます。
AI エージェントの権限は、認可時に設定された権限スコープによって管理されます
時間。スコープは資格情報の権限を制限します。読み取りスコープは、たとえば、
読み取りはできますが、書き込みはできません。読み取りスコープを持つエージェントは、
会議データを変更する。スコープでは送信に名前を付けることもできます: read#10 スコープ
エージェントが提出物 #10 (レビューとコメントを含む) を読むことができます。
他には何もありません。 read#agent を使用すると、#agent タグが付いた送信を読み取ることができます。
read?q=dec:yes を指定すると、承認された送信内容を読み取ることができます。 (具体的には、
接続されているユーザーが受け入れステータスを確認できる提出物: スコープ
接続しているユーザーが他の方法で表示できない情報を公開しないでください)。
PC メンバーが役割を失った場合、またはカンファレンスでエージェントを使用できる人が狭まった場合、
関連付けられた資格情報が機能しなくなります。ユーザーは自分のエージェントを表示して取り消すことができます
[プロフィール] > [開発者] で。
エージェントに関連するスコープには、submeta:read (送信フィールド)、
ドキュメント:読み取り (PDF)、タグ:読み取り、レビュー:読み取り、コメント:読み取り、タグ:書き込み、
review:write および comment:write 。 Paper:read のようなスコープは、次のすべてを許可します。
サブメタ、ドキュメント、タグ、レビュー、コメント。
管理者は、エージェント用に設計されたボット アカウントを作成および管理できます。ボット
通常は非公開の PC メンバーであり、PC 権限を持っています。レビューを割り当てることができます。
たとえば。他のアカウントとは異なります:
ボットは、API を使用して会議管理者によって許可された場合にのみサインインします。
トークンまたは OAuth 認証。
ボットは決して匿名ではなく、AI として明示的に識別されます。
[設定] > [AI] でボットを作成します。
AI レビューは私たちのコミュニティにとって新しいものです。カンファレンスはさまざまなことを望んでおり、
セットします

エージェントが利用できるサイトの機能はさらに拡大しています。メールしてください
質問、アイデア、バグがある場合は、
問題を作成するには GitHub を使用します。

## Original Extract

News – HotCRP.com
HotCRP.com AI agents and bot accounts
19 August 2026 — HotCRP.com sites can now be connected to AI agents. If
conference administrators allow it, bot accounts can read submissions, set tags,
and write reviews, and PC members can connect agents that work on their behalf.
Agents connect over MCP , the protocol that
Claude , Claude Code, and other clients use to reach outside
services. Agents can currently search submissions, read reviews, comments, and
submitted documents, set tags, and write reviews and comments.
AI agent support is disabled by default. It can be enabled using Settings >
AI on all HotCRP.com sites, including test sites .
HotCRP sends nothing to any AI service on its own, and connecting an agent is
never automatic. But once connected, the submissions, reviews, and reviewer
identities read by an agent are sent to the service that runs the agent. Whether
the service keeps the data, trains on it, or exposes it to others is between the
service operator and the person who authorized the connection. HotCRP doesn’t
control that, and a conference’s promise of confidentiality to its authors and
reviewers does not automatically extend to AI services.
Sponsors and conferences differ on whether confidential submissions may be sent
to third-party services at all, and some forbid it. Check your sponsor’s policy
before enabling agents.
Note that most AI services have settings that govern how conversations are
retained and used. Defaults differ by vendor and by plan, and business plans
often differ from personal plans. It’s not always easy to control these
settings. Conferences concerned about agent confidentiality should limit AI
agent use to administrators.
Submissions and reviews might also contain text aimed at agents rather than
humans. A paper can contain instructions; an agent that can write may be induced
to follow them. Scopes bound the damage a rogue agent can do, and
changes made by agents are logged.
AI agent permissions are governed by permission scopes set at authorization
time. A scope limits a credential’s rights. The read scope, for example,
allows reading, but not writing; an agent with read scope is prevented from
modifying conference data. Scopes can also name submissions: the read#10 scope
allows an agent to read submission #10 (including reviews and comments), and
nothing else; read#agent lets it read submissions with tag #agent, and
read?q=dec:yes lets it read accepted submissions. (Specifically, it can read
those submissions whose acceptance status the connected user can see: scopes do
not expose information the connected user couldn’t otherwise view.)
If a PC member loses a role, or a conference narrows who may use agents, the
associated credentials stop working. Users can see and revoke their own agents
under Profile > Developer.
Scopes relevant for agents include submeta:read (submission fields),
document:read (PDFs), tag:read , review:read , comment:read , tag:write ,
review:write , and comment:write . A scope like paper:read grants all of
submeta , document , tag , review , and comment .
Administrators can create and manage bot accounts designed for agent use. Bots
are typically unlisted PC members, with PC rights; they can be assigned reviews,
for example. Unlike other accounts:
Bots sign in only as authorized by conference administrators using API
tokens or OAuth authorization.
Bots are never anonymous and they are explicitly identified as AI.
Create bots under Settings > AI.
AI review is new to our community. Conferences want different things, and the
set of site features available to agents is still expanding. Please email me
with questions, ideas, and bugs, or use
GitHub to create issues.
