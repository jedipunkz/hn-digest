---
source: "https://code.claude.com/docs/en/cross-session-messaging"
hn_url: "https://news.ycombinator.com/item?id=49217752"
title: "Message your other Claude Code sessions"
article_title: "Message your other Claude Code sessions - Claude Code Docs"
author: "adamfeldman"
captured_at: "2026-08-08T00:53:12Z"
capture_tool: "hn-digest"
hn_id: 49217752
score: 1
comments: 0
posted_at: "2026-08-08T00:29:23Z"
tags:
  - hacker-news
  - translated
---

# Message your other Claude Code sessions

- HN: [49217752](https://news.ycombinator.com/item?id=49217752)
- Source: [code.claude.com](https://code.claude.com/docs/en/cross-session-messaging)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T00:29:23Z

## Translation

タイトル: 他のクロード コード セッションにメッセージを送信してください
記事のタイトル: 他の Claude Code セッションにメッセージを送信する - Claude Code Docs
説明: クロードが 1 台のマシン上で他のクロード コード セッションをリストしてメッセージを送信し、他のマシンまたは Web 上のセッションにリモート コントロール経由で返信できるようにします。

記事本文:
他の Claude Code セッションにメッセージを送信する - Claude Code Docs Documentation Index
/docs/llms.txt で完全なドキュメントのインデックスを取得します。
さらに探索する前に、このファイルを使用して利用可能なすべてのページを検出します。
メイン コンテンツにスキップ Claude Code Docs ホーム ページ 英語 検索... ⌘ K Ask Assistant Claude 開発者プラットフォーム
検索... ナビゲーション エージェントと並行作業 他のクロード コード セッションにメッセージを送信する はじめに クロード コードで構築する 管理構成リファレンス エージェント SDK 新機能 リソース エージェントと並行作業
ワークツリーを使用してセッションを分離する
事前に構築されたプラグインを検出してインストールする
セッション出力をアーティファクトとして共有する
外部イベントをクロードにプッシュする
インストールとログインのトラブルシューティング
パフォーマンスと安定性のトラブルシューティング
クロスセッションメッセージングを使用する場合
別のセッションにメッセージを送信する メッセージ配信
クロードが到達できるセッションを確認する
他のマシンでのメッセージセッション
セッションが受信メッセージを処理する方法 メッセージの外観
クロスセッションメッセージングを制限する クロスマシンメッセージの承認を必要とする
クロスセッションメッセージングをオフにする
他のクロード コード セッションにメッセージを送信する
ページをコピー ページをコピー クロードが 1 台のマシン上で他のクロード コード セッションをリストしてメッセージを送信し、他のマシンまたは Web 上でリモート コントロール経由でセッションに返信できるようにします。
ページをコピー ページをコピー クロスセッション メッセージングには、Claude Code v2.1.224 以降が必要で、macOS および Linux で実行されます。セッションが要件を満たしている場合、メッセージングは​​オンになり、有効にするものは何もありません。プロバイダーの要件と、セッションにそれが存在することを確認する方法については、「可用性」を参照してください。
クロスセッション メッセージングを使用すると、クロードはクロード コード セッションの 1 つから別のセッションにメッセージを配信できます。あるセッションでの変更により、別のセッションの構築が中断されると、クロードはユーザーが気づく前にそのセッションに警告することができます。だれ

あるセッションで他のセッションで解決できなかった質問が解決された場合、クロードは答えを送信できます。
メッセージはクロードが別のクロードに書き込むテキストであり、会話履歴やファイルではありません。会話全体またはそのコンテキストを移動するには、代わりにセッションを再開します。
Claude は、このために 2 つのツールを使用します。ListAgents を使用して、どのエージェントに連絡できるかを検出し、SendMessage を使用して、そのうちの 1 人に名前でメッセージを配信します。クロードは、同じ SendMessage ツールを使用して、単一のセッションまたはチーム内のサブエージェントやエージェント チームのチームメイトにメッセージを送信することもできます。このページでは、独立したセッション間のメッセージについて説明します。
クロスセッションメッセージングを使用する場合
結果を引き渡す : あるセッションで重大な変更が発見された場合、または決定が下された場合、あなたがそこで再度説明するのではなく、影響を受ける領域に取り組んでいるセッションのためにクロードがそれを要約します。
並列ワークツリーを調整する: セッションが別々のワークツリー内の同じリポジトリで作業する場合、クロードは他のセッションに何が着地したかを伝えることができます。
長時間実行されている作業からステータスを取得します。監視しているセッションに移行またはテスト実行のレポートを返すか、そこから自分で質問します。
マシン間で返信 : セッションの 1 つから別のマシンまたは Web 上に到着したメッセージに返信します。マシンを越えて、クロードは返信することしかできません。交換を開始できません。
別の端末で 1 つの会話を継続するか、そのコンテキストを新しいセッションで共有するには、セッションを再開します
クロードが生成および監督するセッションの調整されたチームの場合は、エージェント チームを使用します。
1 か所から多くのセッションを監視および操作するには、エージェント ビューを使用します
セッション間で相互にメッセージを送信するのではなく、携帯電話や別のデバイスから自分でセッションを操作するには、リモート コントロールを使用します。
CI 結果やチャット メッセージなどの外部イベントをセッションにプッシュするには、チャネルを使用します
私の他の環境で実行されているセッションに質問してください

r ターミナル移行が完了したかどうか
実際のメッセージ自体はクロードが書き込むため、プロンプトの内容はクロードに任せることができます。このプロンプトは、文言を指示することなく概要を要求し、クロードが送信する内容は次のように異なります。
支払い API で作業するセッションに対して何を行ったか説明します。
クロードが書いたメッセージが到着したときにどのように表示されるかについては、その例も含めて、メッセージの表示を参照してください。
メッセージ配信
Delivered : クロード コードはメッセージを受信側のクロードに渡します。
保留 : クロード コードはメッセージが配信されないまま保留します。保留されたメッセージは、ユーザーが承認するか、後のモードや設定の変更で許可された場合にのみクロードに届きます。
拒否 : クロード コードはメッセージを配信せずにドロップします。
クロードが到達できるセッションを確認する
サブエージェント : 現在のセッション内で実行されているエージェント。エージェントチームのチームメイトはリストされていません。クロードはチームの名簿を通じて彼らにメッセージを送ります。
他のローカル セッション: 同じマシン上で実行されているクロード コード セッション (バックグラウンド セッションを含む)。セッションは、受信ボックスのソケットをバインドする場合にのみ表示されます。
このマシンを超えるセッション: リモート コントロールが接続されているときに表示され、「リモート コントロール」というラベルが付けられます。これらは、他のマシン上のセッションと Web セッション上のクロード コードです。クロードは、これらのセッションのいずれかと会話を開始するためのメッセージを送信できません。いずれか 1 つから届いたメッセージにのみ返信できます。 「他のマシンでのメッセージ セッション」を参照してください。
他のマシンでのメッセージセッション
セッションが受信メッセージを処理する方法
何も承認できません。別のセッションからのメッセージは決して同意としてカウントされないため、保留中の許可プロンプトにユーザーに代わって応答することはできません。
設定を変更することはできません : クロード コードは、受信側のクロードにアクセス許可設定を決して変更しないように指示します。CL

AUDE.md 、または別のセッションが要求したためのその他の設定。
コマンドは実行されません。メッセージのテキスト内のコマンド ( /compact など) はプレーン テキストとして届きます。クロードコードは決して実行しません。
権限プロンプトは引き続き表示されます。メッセージの操作に受信セッションにない権限が必要な場合は、他の作業の場合と同じプロンプトが表示されます。
スキーマの移行が完了しました。新しい列は tenant_id で、main でのリベースは安全になりました。
受信メッセージを制御する
受信セッションではアクセス許可の入力が求められます。Claude Code は各メッセージを配信します。送信セッションが許可プロンプトをバイパスしていると認識した場合にのみ、承認のために保持されます。
受信セッションは許可プロンプトをバイパスします。クロード コードは承認のために各メッセージを保持します。送信セッションが自身をバイパスしていると認識した場合にのみ、セッションを配信します。
承認はその一言をクロードに伝える。
Deny 、つまりダイアログを閉じると、ダイアログが削除されます。
DialogExpiry の期限を過ぎても応答がないまま放置すると、ダイアログが閉じ、Claude Code がメッセージを削除します。期限のデフォルトは 5 分です。
メッセージが保持されている間にこのセッションのアクセス許可モード クラスが変更された場合、クロード コードは受信ルールを再適用し、現在受け入れられるメッセージを配信し、通知を表示します。
変更により、メッセージが保留されている間に拒否が適用されるようになった場合、クロード コードは保留されているすべてのメッセージを削除し、到達可能な各送信者に拒否を報告します。
/status はピアアドレス行にそれを示します。パスには uds: という接頭辞が付きます。
クロード コードは、これを CLAUDE_CODE_MESSAGING_SOCKET 環境変数としてフックおよび Bash コマンドにエクスポートします。エクスポートは、 SessionStart を含むフックが実行される前に行われます。各セッションは独自のソケットをエクスポートします。親セッションから継承されることはありません。
自身の子メッセージ:crossSessionInbound 値 ap がない場合

つまり、Claude Code は、セッション自身のソケットにポストバックするフックや Bash コマンドなど、セッション自身の子プロセスからのものであることを確認したメッセージを配信します。 WSL 2 内を含む Linux では、すでに終了した子であっても検証できますが、macOS では投稿プロセスがまだ実行されている間のみ検証でき、クロード コードがプロセス ID 1 として実行されているコンテナではまったく検証できません。検証できない場合は、アクセス許可クラスをアサートしない他のメッセージと同様にメッセージを処理するため、アクセス許可のプロンプトをバイパスするセッションは、メッセージを承認のために保持します。
サンドボックスセッション: サンドボックスの Unix ソケット設定、 Sandbox.network.allowAllUnixSockets および Sandbox.network.allowUnixSockets を使用して、Bash コマンドがサンドボックス内からソケットに到達できるかどうかを制御します。
セッション間のメッセージングを制限する
クロスマシンメッセージには承認が必要
{
"isolatePeerMachines" : true
}
このセットを使用すると、通常の権限プロンプトをスキップする bypassPermissions モードであっても、クロード コードは、このマシンを超えたセッションへのクロードの応答が終了する前に承認を求めます。どの設定スコープの true も適用されるため、チェックインされたプロジェクト ファイルによって要件をオンにすることはできますが、オフにすることはできません。同じマシン上のセッション間のメッセージにはプロンプトが表示されません。
セッション間のメッセージングをオフにする
受信を停止します。crossSessionInbound を拒否するように設定すると、クロード コードは受信ピア メッセージを配信せずにドロップします。プロジェクトまたはローカル設定からは、拒否は他のすべてのソースに適用され、ユーザー設定からは、管理対象設定または --settings フラグが値を設定しない限り適用されます。
送信とリストを停止します。SendMessage および ListAgents という名前のアクセス許可拒否ルールを追加します。どちらも、指定子のない裸のツール名を受け取ります。
{
「権限」: {
"拒否" : [ "メッセージの送信"

、「エージェントのリスト」]
}、
"crossSessionInbound" : "拒否"
}
これを導入しても、クロード コードは引き続き各セッションの受信トレイ ソケットをバインドしますが、そこに到着するすべてのメッセージはクロードに何も配信せずにドロップされます。同じツールが両方に機能するため、SendMessage を拒否すると、サブエージェントとエージェント チームのチームメイトへのメッセージングも削除されます。拒否したセッションでは、そのセッション自体の /status や他のセッションのリストに目に見える変化が見られないため、セッションの構成から設定を確認してください。
可用性
オペレーティング システム : macOS および Linux (WSL 2 内の Linux を含む) で利用可能。Claude Code はネイティブ Windows ではクロスセッション メッセージングを提供しません。
プロバイダー : Amazon Bedrock、AWS の Claude Platform、Google Cloud のエージェント プラットフォーム、または Microsoft Foundry では利用できません。
機能フラグ評価 : CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC 、 DISABLE_TELEMETRY 、 DO_NOT_TRACK 、または DISABLE_GROWTHBOOK のいずれかが、その機能が依存する機能フラグ評価をオフにすると、クロスセッション メッセージングは​​オフのままになります。各変数の行には、どの値がそれを行うのかが示されています。該当する方の設定を解除します。これらの変数は、シェル、設定ファイルの環境マップ、または管理設定から取得できます。
/list-agents は認識されません。セッションにはセッション間のメッセージングがありません。バージョン要件の claude --version から始めて、上記の要件を実行します。
/list-agents は機能しますが、送信は到着しませんでした。メッセージングがオンになっており、より狭いものが適用されます。アクセス許可拒否ルールにより、SendMessage ツールと ListAgents ツールが削除され、受信セッションの受信コントロールは送信内容を保持またはドロップでき、このマシンを超えるセッションは応答専用になります。
プレーン テキストのみ : クロードはセッション間でプレーン テキストのみを送信します。構造化されたエージェント チーム プロトコル メッセージはチーム内に留まります。
メッセージループ

are throttled : Claude Code は、送信者ごとに繰り返されるメッセージのレートを制限し、短いウィンドウ内に到着する同一の繰り返しをドロップし、クロードが読み取るのを待機している受け入れられたメッセージをセッションごとに 50 に制限します。したがって、2 つのセッション間のメッセージ ループは自動的に停止します。
サブエージェントとエージェント チーム: 単一のセッションまたはチーム内でのメッセージング
バックグラウンド エージェント: メッセージを送信する可能性のある並列セッションをディスパッチして監視します
リモート コントロール : クロスマシン メッセージングが通過するセッションを接続します。
設定:crossSessionInbound、isolatePeerMachines、dialogExpiry
許可モード: インバウンドのデフォルトの 2 つのクラスの背後にあるモード
ツールリファレンス: ツールテーブルの ListAgents 行と SendMessage 行
エージェントを並行して実行する : Claude Code が複数のエージェントを実行する方法を比較する
はい いいえ エージェント チームの実行 動的ワークフロー ⌘ I Claude Code Docs ホーム ページ x linkedin Company

## Original Extract

Let Claude list and message your other Claude Code sessions on one machine, and reply to your sessions on other machines or on the web through Remote Control.

Message your other Claude Code sessions - Claude Code Docs Documentation Index
Fetch the complete documentation index at: /docs/llms.txt
Use this file to discover all available pages before exploring further.
Skip to main content Claude Code Docs home page English Search... ⌘ K Ask Assistant Claude Developer Platform
Search... Navigation Agents and parallel work Message your other Claude Code sessions Getting started Build with Claude Code Administration Configuration Reference Agent SDK What's New Resources Agents and parallel work
Isolate sessions with worktrees
Discover and install prebuilt plugins
Share session output as artifacts
Push external events to Claude
Troubleshoot installation and login
Troubleshoot performance and stability
When to use cross-session messaging
Message another session Message delivery
See which sessions Claude can reach
Message sessions on other machines
How a session treats an incoming message What a message looks like
Restrict cross-session messaging Require approval for cross-machine messages
Turn off cross-session messaging
Message your other Claude Code sessions
Copy page Copy page Let Claude list and message your other Claude Code sessions on one machine, and reply to your sessions on other machines or on the web through Remote Control.
Copy page Copy page Cross-session messaging requires Claude Code v2.1.224 or later and runs on macOS and Linux. When a session meets the requirements, messaging is on with nothing to enable. See Availability for provider requirements and how to confirm a session has it.
Cross-session messaging lets Claude deliver a message from one of your Claude Code sessions to another. When a change in one session breaks what another is building on, Claude can warn that session before you notice. When one session settles a question another is blocked on, Claude can send the answer across.
A message is a piece of text one Claude writes to another, never conversation history or files. To move a whole conversation or its context, resume the session instead.
Claude uses two tools for this: ListAgents to discover which agents it can reach, and SendMessage to deliver a message to one of them by name. With the same SendMessage tool, Claude can also message subagents and agent team teammates within a single session or team. This page covers messages between your independent sessions.
​ When to use cross-session messaging
Hand over a finding : when one session discovers a breaking change or makes a decision, Claude summarizes it for the session working on the affected area, instead of you re-explaining it there.
Coordinate parallel worktrees : when sessions work the same repository in separate worktrees , Claude can tell the other sessions what landed.
Get status from long-running work : have a migration or test run report back to the session you’re watching, or ask it yourself from there.
Reply across machines : answer a message that arrived from one of your sessions on another machine or on the web. Across machines, Claude can only reply. It can’t start the exchange.
To continue one conversation in another terminal, or share its context with a new session, resume the session
For a coordinated team of sessions Claude spawns and supervises, use agent teams
To watch and steer many sessions from one place, use agent view
To steer a session yourself from your phone or another device, rather than have sessions message each other, use Remote Control
To push external events, such as CI results or chat messages, into a session, use channels
Ask the session running in my other terminal whether the migration finished
Claude writes the actual message itself, so your prompt can leave the content to Claude. This prompt asks for a summary without dictating its wording, and what Claude sends varies:
Explain what we just did to the session working on the payments API
For what the message Claude writes looks like when it arrives, including an example of one, see what a message looks like .
​ Message delivery
Delivered : Claude Code passes the message to the receiving Claude.
Held : Claude Code sets the message aside undelivered. A held message reaches Claude only when you approve it or a later mode or settings change allows it.
Refused : Claude Code drops the message without delivering it.
​ See which sessions Claude can reach
Subagents : agents running inside the current session. Agent team teammates aren’t listed; Claude messages them through the team’s own roster.
Your other local sessions : Claude Code sessions running on the same machine, including background sessions . A session appears only when it binds an inbox socket .
Sessions beyond this machine : shown while Remote Control is connected and labeled Remote Control . These are your sessions on other machines and your Claude Code on the web sessions. Claude can’t send a message to start a conversation with one of these sessions. It can only reply to a message that arrived from one of them. See Message sessions on other machines .
​ Message sessions on other machines
​ How a session treats an incoming message
It can’t approve anything : a message from another session never counts as your consent, so it can’t answer a pending permission prompt on your behalf.
It can’t change configuration : Claude Code instructs the receiving Claude never to change permission settings, CLAUDE.md , or other configuration because another session asked.
Commands don’t run : a command in the message’s text, such as /compact , arrives as plain text. Claude Code never executes it.
Permission prompts still fire : if acting on the message requires a permission the receiving session doesn’t have, you see the same prompt you’d see for any other work.
Schema migration finished: the new column is tenant_id, and rebasing on main is safe now.
​ Control inbound messages
The receiving session prompts for permissions : Claude Code delivers each message. It holds one for your approval only when the sending session identifies itself as bypassing permission prompts.
The receiving session bypasses permission prompts : Claude Code holds each message for your approval. It delivers one only when the sending session identifies itself as also bypassing.
Approve delivers that one message to Claude.
Deny , or dismissing the dialog, drops it.
Left unanswered past the dialogExpiry deadline, the dialog closes and Claude Code drops the message. The deadline defaults to five minutes.
If this session’s permission-mode class changes while messages are held, Claude Code re-applies the inbound rules, delivers the messages they now accept, and shows a notice.
If a change makes refuse apply while messages are held, Claude Code drops every held message and reports a denial to each sender it can reach.
/status shows it in the Peer address row. The path is prefixed with uds: .
Claude Code exports it to hooks and Bash commands as the CLAUDE_CODE_MESSAGING_SOCKET environment variable. The export happens before any hook runs, including SessionStart . Each session exports its own socket, never one inherited from a parent session.
Own-child messages : when no crossSessionInbound value applies, Claude Code delivers a message it verifies came from the session’s own child processes, such as a hook or Bash command posting back to its own session’s socket. On Linux, including inside WSL 2, it can verify even for a child that has already exited, while on macOS it can verify only while the posting process is still running, and in containers where Claude Code runs as process ID 1 it can’t verify at all. Whenever it can’t verify, it treats the message like any other that asserts no permission class, so a session that bypasses permission prompts holds it for your approval.
Sandboxed sessions : control whether a Bash command can reach the socket from inside the sandbox with the sandbox’s Unix-socket settings, sandbox.network.allowAllUnixSockets and sandbox.network.allowUnixSockets .
​ Restrict cross-session messaging
​ Require approval for cross-machine messages
{
"isolatePeerMachines" : true
}
With this set, Claude Code asks for your approval before Claude’s reply to a session beyond this machine leaves, even in bypassPermissions mode, which skips ordinary permission prompts. A true from any settings scope applies, so a checked-in project file can turn the requirement on but not off. Messages between sessions on the same machine don’t prompt.
​ Turn off cross-session messaging
Stop receiving : set crossSessionInbound to refuse , and Claude Code drops inbound peer messages without delivering them. From project or local settings, refuse applies over every other source, and from your user settings it applies unless managed settings or the --settings flag set a value.
Stop sending and listing : add permission deny rules naming SendMessage and ListAgents . Both take the bare tool name with no specifier.
{
"permissions" : {
"deny" : [ "SendMessage" , "ListAgents" ]
},
"crossSessionInbound" : "refuse"
}
With this in place, Claude Code still binds each session’s inbox socket, but drops every message that arrives on it without delivering anything to Claude. Denying SendMessage also removes messaging to subagents and agent-team teammates, since the same tool serves both. A refusing session shows no visible change, in its own /status or in other sessions’ listings, so confirm the setting from the session’s configuration.
​ Availability
Operating system : available on macOS and Linux, including Linux inside WSL 2. Claude Code doesn’t offer cross-session messaging on native Windows.
Provider : not available on Amazon Bedrock, Claude Platform on AWS, Google Cloud’s Agent Platform, or Microsoft Foundry.
Feature-flag evaluation : when any of CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC , DISABLE_TELEMETRY , DO_NOT_TRACK , or DISABLE_GROWTHBOOK turns off the feature-flag evaluation the feature depends on, cross-session messaging stays off. Each variable’s row says which values do that. Unset whichever applies. These variables can come from your shell, from a settings file’s env map, or from managed settings.
/list-agents isn’t recognized : the session doesn’t have cross-session messaging. Work through the requirements above, starting with claude --version for the version requirement.
/list-agents works but a send didn’t arrive : messaging is on, and something narrower applies. A permission deny rule removes the SendMessage and ListAgents tools, the receiving session’s inbound controls can hold or drop what you send it, and a session beyond this machine is reply-only .
Plain text only : Claude sends only plain text across sessions. Structured agent team protocol messages stay within a team.
Message loops are throttled : Claude Code rate-limits repeated messages per sender, drops identical repeats arriving within a short window, and caps accepted messages waiting for Claude to read them at 50 per session. A message loop between two sessions therefore stops on its own.
Subagents and agent teams : messaging within a single session or team
Background agents : dispatch and monitor the parallel sessions you might message
Remote Control : connect the sessions that cross-machine messaging travels through
Settings : crossSessionInbound , isolatePeerMachines , and dialogExpiry
Permission modes : the modes behind the inbound default’s two classes
Tools reference : the ListAgents and SendMessage rows in the tools table
Run agents in parallel : compare the ways Claude Code runs multiple agents
Yes No Run agent teams Dynamic workflows ⌘ I Claude Code Docs home page x linkedin Company
