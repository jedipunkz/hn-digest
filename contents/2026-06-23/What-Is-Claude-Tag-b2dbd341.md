---
source: "https://support.claude.com/en/articles/15594475-what-is-claude-tag"
hn_url: "https://news.ycombinator.com/item?id=48649418"
title: "What Is Claude Tag?"
article_title: "What is Claude Tag? | Claude Help Center"
author: "shpat"
captured_at: "2026-06-23T19:36:24Z"
capture_tool: "hn-digest"
hn_id: 48649418
score: 3
comments: 1
posted_at: "2026-06-23T18:39:35Z"
tags:
  - hacker-news
  - translated
---

# What Is Claude Tag?

- HN: [48649418](https://news.ycombinator.com/item?id=48649418)
- Source: [support.claude.com](https://support.claude.com/en/articles/15594475-what-is-claude-tag)
- Score: 3
- Comments: 1
- Posted: 2026-06-23T18:39:35Z

## Translation

タイトル: クロードタグとは?
記事タイトル：クロードタグとは？ |クロード ヘルプセンター

記事本文:
クロード・タグとは何ですか？ |クロード ヘルプ センター メイン コンテンツにスキップ API ドキュメント リリース ノート サポートを受ける方法 English Français Deutsch Bahasa India Italiano 日本語 한국어 Português Pусский 简体中文 Español 繁體中文 English API ドキュメント リリース ノート サポートを受ける方法 English Français Deutsch Bahasa India Italiano 日本語 한국어 Português Pусский 简体中文 Español 繁體中文 English 記事の検索... すべてのコレクション
Slack のクロードは、2026 年 8 月 3 日に新しいクロード タグ エクスペリエンスに切り替わります。クロードと Slack を統合するには、代わりにクロード タグを使用してください。
Claude Tag は、Claude を操作するための新しい方法です。@Claude を会話にタグ付けすると、組織のツールとその周りの共有コンテキストを使用して、実際の作業が行われます。クロードは独自のアイデンティティの下で動作し、参加しているチャネルから関連情報を記憶することでコンテキストを構築し、独自にフォローアップできます。
クロード タグは、ベータ版のチーム プランとエンタープライズ プランで利用できます。クロード・タグは現在 Slack で働いています。
このようにして、Claude の機能を Slack に直接導入し、チームのワークスペースに AI 支援を導入しました。この統合により、Slack を離れることなく、次の 3 つの便利な画面を通じてクロードと作業できるようになります。
チャネルのタグ付け: 任意のチャネルで @Claude をタグ付けしてタスクを渡し、スレッド内での動作に従ってください。
Claude とのダイレクト メッセージ : @Claude とプライベートな会話を開始します。
AI アシスタント パネル : Slack の AI アシスタント ヘッダーのクロード アイコンをクリックすると、Slack ウィンドウの右側にパネルが開き、Slack アプリのどこからでもクロードにアクセスできるようになります。
チャネル内で @Claude をタグ付けすると、Claude は交換全体をチャネル内の全員に表示したままタスクを実行します。チャンネル内の全員が協力します

同じクロードなので、誰でも操縦したり、中断したところから再開したりできます。クロードは、ジョブが終了したときに投稿したり、スレッドが停止したときにタグ付けしたりするなど、独自にチェックインすることもできます。詳細については、「Claude タグの概要」を参照してください。
ダイレクト メッセージとアシスタント パネルでは、クロードは、Web 検索や接続ツールなど、自分のクロード アカウントで有効にした機能を利用できます。チャネルのタグ付けの仕組みは異なります。クロードは組織の ID の下でツールを使用して行動し、そのチャネル用に設定された管理者にアクセスします。作業はあなたではなく組織に請求されます。
クロード アプリのインストール後、プライマリ オーナーまたはオーナーはクロード タグを設定します。つまり、クロードの ID をプロビジョニングし、組織のツールとリポジトリに接続し、クロード タグが使用できるチャネルを選択します。チャネルの準備ができたら、チームのメンバーは個別に何も設定する必要はありません。完全なチュートリアルについては、「Claude Tag セットアップ ガイド」を参照してください。
重要: クロード タグのアクセスとチャネルを設定できるのは、プライマリ オーナーまたはオーナーのみです。管理者の役割ではできません。
クロードタグを使用できるユーザーを制御する
Slack の [組織設定] > [Claude] では、メンバー アクセスには 3 つのモードがあります。Slack ワークスペース内のすべてのユーザーに公開、クロード組織の任意のメンバーに公開、またはロールで許可されているメンバーのみに公開です。 3 番目のオプションはロールベースのアクセスで、Claude Enterprise プランで利用できます。ロールごとに制限するには、メンバーアクセスを「ロールが許可するメンバーのみ」に設定し、カスタムロールに「Slack のクロード」機能を付与します。この設定は、チャンネルのメンションとダイレクト メッセージの両方に適用されます。メンバー アクセスを設定するには、「Claude Tag が動作する場所を制限する」を参照してください。
Claude Tag の支出制限を管理する
Claude Tag は消費ベースであるため、支出は人数ではなく使用量に基づいて計算されます。プライマリオーナーまたはオーナーとして、あなたは

管理コンソールの使用設定から制御します。
組織全体の制限: すべてのチャネルにわたるクロード タグの合計支出に対するハードキャップ。支出がそれを超えることはできません。
チャネルごとの制限: 組織全体の上限に加えて、個々のチャネルに制限を設定します。新しいチャネルはデフォルトの制限を継承します。
しきい値アラート: 制限の 75% および 95% になると管理者に通知されます。
使用状況分析: チャネルごとの支出の内訳が同じページに表示されます。
注: 制限を超える作業は拒否され、黙って短縮されることはありません。ブロックされたユーザーは、Slack を離れることなく管理者に追加のリクエストを行うことができ、制限または利用可能な残高がブロックの原因かを示すアラートが表示されます。
チャンネル内でクロードをタグ付けすると、組織に請求されます。ダイレクト メッセージは、代わりに自分のクロード アカウントに請求されます。
制限を設定するには、「Claude Tag の使用制限を設定する」を参照してください。
Claude タグのアクセス権と権限を設定する
資格情報とリポジトリ アクセスを 3 つのレベルで設定することで、Claude Tag が到達できる範囲を決定します。各レベルは、その上のレベルの権限とメモリを継承します。
組織全体: Claude Tag がインストールされているすべての場所に適用される資格情報とリポジトリ。
ワークスペース: Slack ワークスペース内のすべてのパブリック チャネルに適用されるアクセス。組織全体の権限とメモリを継承します。
プライベート チャネル: ワークスペースがすでに付与しているものに加えて追加の資格情報またはリポジトリ。プライベート チャネルを使用して、小規模なグループへの機密性の高い接続を維持します。たとえば、法務業務用に設定されたチャネルでは、そのツールとメモリがエンジニアリング チャネルから分離されています。
アクセスを構成するには、「Claude Tag の ID とアクセス」を参照してください。
クロード・タグの記憶とアクティビティを確認する
Claude Tag は、チャネルごとおよびワークスペースごとにコンテキストを保持します。管理者はそのメモリを表示、編集、削除できます。
組織設定の監査ビュー

> Claude Tag > Audit は、エージェント ID を使用して行われたすべてのネットワーク呼び出しに加えて、組織全体のすべてのスケジュールされたタスクと 1 回限りのタスクをリストします。各アクションは、それが発生したツールでも追跡できます。投稿は Slack の Claude アプリから送信され、コミットとプル リクエストには、作成者として Claude GitHub アプリが表示され、それらを開始した Slack スレッドへのリンクが表示されます。どのチャンネルでも、「@Claude ここにどのようなトリガーを設定していますか?」と尋ねることができます。立ち仕事を確認してオフにします。
クロードとの Slack での会話はクロードの履歴とは別に保持されるため、プラットフォーム間で作業が整理されます。
Slack で開始された会話は、クロードのチャット履歴には表示されません。
Claude Web アプリで開始された会話には、Slack ではアクセスできません。
各プラットフォームは個別の会話履歴を保持します。
統合を切断するかアプリをアンインストールすると、会話は 30 日以内にクロードから自動的に削除されます。
Slack での会話は、組織の Slack 保持ポリシーに従います。
クロード タグはコンテキストを記憶して作業を行うため、チャネルとワークスペースのメモリは各タスクの後に破棄されずに保持されます。記憶とアクティビティはチャネルの境界を尊重し、管理者はクロードが記憶している内容を確認または削除できます。チャネルの作業は組織のクロード ID に帰属しますが、ダイレクト メッセージで行われた作業はあなた自身のアカウントで実行されます。
Claude Tag は、私がすでに使用している Slack の Claude とどう違うのですか?
Claude Tag は、同じ場所でのその体験の次世代です。使い慣れた作業方法が引き続き適用され、クロードはさらに多くのことができるようになりました。数日間にわたってコンテキストを記憶し、独自のフォローアップをスケジュールし、積極的にチェックインし、独自のアイデンティティの下で行動します。組織のプライマリ オーナーまたはオーナーが、組織の移行をオプトインします。
誰が

クロードタグを設定できるのですが、その費用は誰が支払いますか?
クロード タグのアクセスとチャネルを設定できるのは、組織のプライマリ オーナーまたはオーナーだけです。チャンネル内でクロードをタグ付けすると、組織に請求されます。ダイレクト メッセージは、代わりにその人自身のクロード アカウントに請求されます。
クロードのチャット検索とメモリを使用して以前のコンテキストを構築する Slack でクロードを使用する Chrome 管理コントロールでクロードを使用する 認証済みドメイン コネクタを企業に制限する これで質問は解決しましたか? 😞 😐 😃 目次 製品

## Original Extract

What is Claude Tag? | Claude Help Center Skip to main content API Docs Release Notes How to Get Support English Français Deutsch Bahasa Indonesia Italiano 日本語 한국어 Português Pусский 简体中文 Español 繁體中文 English API Docs Release Notes How to Get Support English Français Deutsch Bahasa Indonesia Italiano 日本語 한국어 Português Pусский 简体中文 Español 繁體中文 English Search for articles... All Collections
Claude in Slack will be switched over to the new Claude Tag experience on August 3, 2026. To integrate Claude and Slack, use Claude Tag instead.
Claude Tag is a new way to work with Claude: tag @Claude into a conversation and it takes on real work, using your organization's tools and the shared context around it. Claude works under its own identity, builds context by remembering relevant information from the channels it’s in, and can follow up on its own.
Claude Tag is available on Team and Enterprise plans in beta. Claude Tag works in Slack today.
It’s how we’ve brought Claude’s capabilities directly to Slack, bringing AI assistance into your team’s workspace. This integration allows you to work with Claude without leaving Slack through three convenient surfaces:
Channel tagging: Tag @Claude in any channel to hand it a task, and follow along as it works in the thread.
Direct message with Claude : Start a private conversation with @Claude.
AI assistant panel : Click the Claude icon in Slack's AI assistant header to open a panel on the right side of your Slack window, allowing you to access Claude from anywhere in the Slack app.
When you tag @Claude in a channel, Claude works through the task while the whole exchange stays visible to everyone in the channel. Everyone in a channel works with the same Claude, so anyone can steer it or pick up where it left off. Claude can also check in on its own, like posting when a job finishes or tagging you when a thread stalls. To learn more, see the Claude Tag overview .
In direct messages and the assistant panel, Claude has the capabilities you've enabled in your own Claude account, like web search and your connected tools. Channel tagging works differently: Claude acts under your organization's identity, using the tools and access an admin set up for that channel, and the work is billed to your organization rather than to you.
After the Claude app is installed, a Primary Owner or Owner sets up Claude Tag: provision Claude's identity, connect your organization's tools and repositories, and choose which channels Claude Tag can work in. People on your team don't need to set up anything individually once a channel is ready. For the full walkthrough, see the Claude Tag setup guide .
Important: Only a Primary Owner or Owner can set up Claude Tag's access and channels. The Admin role can't.
Control who can use Claude Tag
In Organization settings > Claude in Slack , Member Access has three modes: open to anyone in the Slack workspace, open to any member of your Claude organization, or only members whose role allows it. The third option is role-based access and is available on the Claude Enterprise plan. To restrict by role, set Member Access to "Only members whose role allows it" and grant the "Claude in Slack" capability to a custom role. This setting applies to both channel mentions and direct messages. To set member access, see Restrict where Claude Tag operates .
Manage spend limits for Claude Tag
Claude Tag is consumption-based, so spend is based on usage rather than the number of people. As a Primary Owner or Owner, you control it from the usage settings in your admin console.
Organization-wide limit: a hard cap on total Claude Tag spend across every channel. Spend can't exceed it.
Per-channel limits: set a limit on any individual channel, on top of the organization-wide cap. New channels inherit a default limit.
Threshold alerts: admins are notified at 75% and 95% of any limit.
Usage analytics: a per-channel spend breakdown lives on the same page.
Note: Work that would go over a limit is declined, never silently cut short. A blocked user can request more from their admin without leaving Slack, and the alert says whether the limit or the available balance caused the block.
Tagging Claude in a channel is billed to your organization. Direct messages are billed to your own Claude account instead.
To set limits, see Set spend limits for Claude Tag .
Set access and permissions for Claude Tag
You decide what Claude Tag can reach by setting credentials and repository access at three levels. Each level inherits the permissions and memory of the one above it.
Organization-wide: credentials and repositories that apply everywhere Claude Tag is installed.
Workspace: access that applies to every public channel inside a Slack workspace. Inherits organization-wide permissions and memory.
Private channel: extra credentials or repositories on top of what the workspace already grants. Use a private channel to keep sensitive connections to a smaller group. For example, a channel set up for legal work keeps its tools and memory separate from an engineering channel.
To configure access, see Claude Tag identity and access .
Review memory and activity for Claude Tag
Claude Tag keeps context per channel and per workspace. Admins can view, edit, and delete that memory.
An Audit view in Organization settings > Claude Tag > Audit lists every scheduled and one-time task across your organization in addition to all network calls made using Agent Identity. Each action is also traceable in the tool where it happened: posts come from the Claude app in Slack, and commits and pull requests show the Claude GitHub App as the author with a link back to the Slack thread that started them. In any channel, you can ask "@Claude what triggers do you have set up here?" to see and turn off standing work.
Your Slack conversations with Claude remain separate from your Claude history, keeping work organized across platforms.
Conversations initiated in Slack are not visible in your Claude chat history .
Conversations initiated in the Claude web app are not accessible in Slack.
Each platform maintains separate conversation histories.
Conversations are automatically deleted from Claude within 30 days if you disconnect the integration or uninstall the app.
Your conversations in Slack follow your organization's Slack retention policies.
Claude Tag remembers context to do its work, so channel and workspace memory is retained rather than discarded after each task. Memory and activity respect channel boundaries, and admins can review or delete what Claude remembers. Channel work is attributed to your organization's Claude identity, while work done in a direct message runs on your own account.
How is Claude Tag different from the Claude in Slack I already use?
Claude Tag is the next generation of that experience, in the same place. The familiar ways of working still apply, and Claude can now do more: it remembers context across days, schedules its own follow-ups, checks in proactively, and acts under its own identity. An organization’s Primary Owner or Owner opts your organization in to move over.
Who can set up Claude Tag, and who pays for it?
Only an organization's Primary Owner or Owner can set up Claude Tag's access and channels. Tagging Claude in a channel is billed to your organization. Direct messages are billed to the person’s own Claude account instead.
Use Claude’s chat search and memory to build on previous context Use Claude in Slack Claude in Chrome admin controls Restrict verified-domain connectors to your Enterprise Did this answer your question? 😞 😐 😃 Table of contents Product
