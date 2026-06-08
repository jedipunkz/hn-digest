---
source: "https://www.bleepingcomputer.com/news/microsoft/hands-on-with-intelligent-terminal-an-ai-powered-windows-terminal/"
hn_url: "https://news.ycombinator.com/item?id=48448582"
title: "Hands on with Intelligent Terminal, an AI-Powered Windows Terminal"
article_title: "Hands on with Intelligent Terminal, an AI-powered Windows Terminal"
author: "thunderbong"
captured_at: "2026-06-08T18:57:24Z"
capture_tool: "hn-digest"
hn_id: 48448582
score: 1
comments: 0
posted_at: "2026-06-08T17:47:08Z"
tags:
  - hacker-news
  - translated
---

# Hands on with Intelligent Terminal, an AI-Powered Windows Terminal

- HN: [48448582](https://news.ycombinator.com/item?id=48448582)
- Source: [www.bleepingcomputer.com](https://www.bleepingcomputer.com/news/microsoft/hands-on-with-intelligent-terminal-an-ai-powered-windows-terminal/)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T17:47:08Z

## Translation

タイトル: AI を搭載した Windows ターミナル、インテリジェント ターミナルの実践
記事のタイトル: AI を活用した Windows ターミナル、インテリジェント ターミナルの実践
説明: Microsoft は、Windows ターミナルのオープンソース フォークを作成しました。

記事本文:
AI を活用した Windows ターミナルであるインテリジェント ターミナルを実際に使ってみよう
CISA: ハッカーは現在、SolarWinds Serv-U の欠陥を悪用してサーバーをクラッシュさせています
Brave Software、有料で肥大化しないブラウジング体験を提供する Origin をリリース
シスコ、パッチが適用されていない SD-WAN ゼロデイが攻撃に悪用されていることを警告
クレジットカード盗難キャンペーンがStripeを悪用して盗まれた支払い情報をホスト
WhatsApp、新たなNSOスパイウェアフィッシング攻撃を阻止したと発表
Gogs はリモートコード実行を可能にする重要なゼロデイパッチを適用します
UniFi OS の重大なバグにより、ハッカーが認証なしで root を取得できる
Wazuh Cloud でセキュリティ運用の複雑さを軽減
Tor ブラウザを使用してダークウェブにアクセスする方法
Windows 11 でカーネルモードのハードウェア強制スタック保護を有効にする方法
Windows レジストリ エディタの使用方法
Windows レジストリをバックアップおよび復元する方法
Windows をセーフ モードで起動する方法
トロイの木馬、ウイルス、ワーム、その他のマルウェアを削除する方法
Windows 7で隠しファイルを表示する方法
Windows で隠しファイルを確認する方法
AI を活用した Windows ターミナルであるインテリジェント ターミナルを実際に使ってみよう
AI を活用した Windows ターミナルであるインテリジェント ターミナルを実際に使ってみよう
Microsoft は、「インテリジェント ターミナル」と呼ばれる Windows ターミナルのオープンソース フォークを作成しました。これにより、通常のセッションを妨げることなく、ターミナル内で直接 AI を使用できるようになります。
Microsoft は、インテリジェント ターミナルを、ターミナルを離れることなく、エラーの説明、コマンドの下書き、問題の修正を支援できる組み込みアシスタントであると説明しています。
まず、エージェントは端末で何が起こっているかを常に認識し、コマンドが失敗した場合に支援します。次に、アクティブなエージェント セッションと過去のエージェント セッションを記憶できるため、場所を失うことなく以前の作業に戻ることができます。
インテリジェント ターミナルを初めて開くと、ターミナル ペインの AI エージェントを選択できます。
私のスクリーンでは

enshot には、GitHub Copilot、Claude、Codex、および Gemini がリストされています。 GitHub Copilot は「インストールされる」と表示されますが、他のものはすでにインストールされています。
自動エラー検出と自動エラー提案には個別のトグルもあります。
エラー検出をオンにすると、ターミナルは失敗したコマンドを認識できるようになります。同様に、エラー提案はさらに進んで、選択された AI エージェントにエラーを送信し、修正を依頼します。
もう 1 つのオプションであるセッション管理により、インテリジェント ターミナルはアクティブなエージェント セッションと過去のエージェント セッションを追跡できます。これにより、以前のエージェントの作業を再開できるようになります。
Terminal AI を設定してしまえば、使い方は非常に簡単です。ターミナルが開き、シェルの下に AI ペインが表示され、「インテリジェント ターミナルへようこそ」と表示されます。
私のハンズオンでは、ターミナル AI モデルとしてクロードを選択しました。これが、ペイン内でクロード コードが実行されている理由です。コーディング タスクを計画し、編集を自動的に承認するか、編集を手動で承認するか、計画を継続するかを尋ねることができます。
左側では、エージェント パネルの表示または非表示を選択し、そのアイコンでエラー検出をオンまたはオフにすることができます。右側には、セッション管理パネルとエージェント ステータス バーを開くエージェント管理アイコンが表示されます。
インテリジェント ターミナルの再開セッションは、その最高の機能の 1 つです
開発者として、私はヘルプとして Windows ターミナルの Claude コードをよく使用します。これはうまく機能しますが、唯一の問題は、Claude の組み込みの再開スキルを使用しない限り、標準のターミナルでセッションを再開できないことです。これにより、モデルのパフォーマンスが低下することがよくあります。
現在の Windows ターミナルには、以前に閉じたタブを開くことができるトグルがありますが、以前のセッションは復元されません。
インテリジェント ターミナルは、セッションを再開する機能によってこれらの問題に対処します。

以前のエージェントの作業を行ったり来たりします。
Terminal AI は素晴らしいアイデアですが、すべての人を対象としたものではありません。Microsoft もそれを理解しています。そのため、ターミナル AI は別個のアプリであり、Windows インストールにはまだ含まれていません。
ご興味がございましたら、Microsoft Store または Github からインテリジェント ターミナルをダウンロードしてください。
攻撃者が行う前にすべての層をテストする
セキュリティ チームは成功した攻撃の 54% を記録し、警告を発するのはわずか 14% です。残りは目に見えずに環境内を移動します。
Picus のホワイトペーパーでは、侵害と攻撃のシミュレーションで SIEM と EDR ルールをテストし、脅威が検出されないようにする方法を示しています。
Microsoft、KB5089549 Windows セキュリティ更新プログラムのインストールの問題を修正
パフォーマンスが向上した Windows 11 KB5089573 更新プログラムがリリースされました
Microsoftは2026年にWindows 11ドライバーの品質を向上させる予定
Microsoft、制限付き Windows ネットワークにおけるパッチ適用の問題を確認
Windows 11 KB5089549 および KB5087420 の累積的な更新プログラムがリリースされました
驚くべき点は、ターミナル内の AI が松葉杖になる可能性があるということです。エラーが自動修正されれば、基本を忘れて grok コマンドの動作が少なくなる可能性があります。確かに便利ですが、本当のテストは、私たちがまだ第一原理から推論することを学ぶかどうかです。
まだメンバーではありませんか?今すぐ登録
シスコ、パッチが適用されていない SD-WAN ゼロデイが攻撃に悪用されていることを警告
CISA: ハッカーは現在、SolarWinds Serv-U の欠陥を悪用してサーバーをクラッシュさせています
中国のAPTが新たなマルウェアを導入し、ハッキングされたネットワークへのアクセスを維持
Wazuh でサイバー回復力を構築: プロアクティブな保護のためのオープンソース SIEM および XDR
新しいウェビナー: デバイス コード フィッシング キットの舞台裏
最後に侵入テストを行ったのは 345 日前です。それ以来何が変わりましたか？
AI ツールが機密データを漏洩しています。無料の監査を受けてください。
2026 年の医療資格情報漏えいの現状: (非ゲート) レポートを読む
利用規約 - プライバシー ポリシー - 倫理声明 - アフィリエイト

開示
著作権 @ 2003 - 2026 Bleeping Computer ® LLC - 全著作権所有
まだメンバーではありませんか?今すぐ登録
どのようなコンテンツが禁止されているかについては、投稿ガイドラインをお読みください。

## Original Extract

Microsoft has created an open-source fork of Windows Terminal called

Hands on with Intelligent Terminal, an AI-powered Windows Terminal
CISA: Hackers now exploit SolarWinds Serv-U flaw to crash servers
Brave Software releases Origin for a paid, bloat-free browsing experience
Cisco warns of unpatched SD-WAN zero-day exploited in attacks
Credit card theft campaign abuses Stripe to host stolen payment info
WhatsApp says it disrupted new NSO spyware phishing attacks
Gogs patches critical zero-day enabling remote code execution
Critical UniFi OS bug lets hackers gain root without authentication
Reducing security operations complexity with Wazuh Cloud
How to access the Dark Web using the Tor Browser
How to enable Kernel-mode Hardware-enforced Stack Protection in Windows 11
How to use the Windows Registry Editor
How to backup and restore the Windows Registry
How to start Windows in Safe Mode
How to remove a Trojan, Virus, Worm, or other Malware
How to show hidden files in Windows 7
How to see hidden files in Windows
Hands on with Intelligent Terminal, an AI-powered Windows Terminal
Hands on with Intelligent Terminal, an AI-powered Windows Terminal
Microsoft has created an open-source fork of Windows Terminal called "Intelligent Terminal," and it allows you to use AI directly inside Terminal without interfering with the regular session.
Microsoft describes the Intelligent Terminal as a built-in assistant that can help you explain errors, draft commands, and fix problems without leaving the terminal.
First, the agent can stay aware of what is happening in your terminal and help when a command fails. Second, it can remember active and past agent sessions, so you can return to earlier work without losing your place.
When you open Intelligent Terminal for the first time, it lets you choose the AI agent for the Terminal pane.
In my screenshot, it lists GitHub Copilot, Claude, Codex, and Gemini. GitHub Copilot is shown as “will be installed,” while the others are already installed.
There are also separate toggles for Automatic error detection and Automatic error suggestion.
When you turn on error detection, Terminal can notice failed commands. Similarly, error suggestion goes further and sends the error to the selected AI agent for a possible fix.
There's another option, Session management, that lets Intelligent Terminal track active and past agent sessions. This is what allows you to reopen previous agent work.
Once you've configured Terminal AI, it's quite easy to use. Terminal opens with an AI pane below the shell, where it says "Welcome to Intelligent Terminal."
In my hands-on, I selected Claude as my Terminal AI model, which is why Claude Code is running inside the pane. It could plan a coding task and then ask whether I wanted to auto-accept edits, manually approve edits, or keep planning.
On the left side, you can choose to show or hide the agent panel and turn error detection on or off through its icon. On the right, you'll see the agent management icon that opens your session management panel and agent status bar.
Intelligent Terminal's Resume session is one of its best features
As a developer, I use Claude Code in Windows Terminal a lot for help, and while it does the job well, the only issue is that you can't resume sessions in the standard Terminal unless you're willing to use Claude's built-in resume skill, which often makes the model perform worse.
Current Windows Terminal does have a toggle that allows it to open previously closed tabs, but that doesn't restore your previous sessions.
Intelligent Terminal addresses these concerns with the ability to resume sessions, so you can always go back and forth between your earlier agent work.
Terminal AI is a great idea, but it's not meant for everyone, and Microsoft understands that, which is why it's a separate app, and it's not included with Windows installations yet.
If you're interested, you can download Intelligent Terminal from the Microsoft Store or Github .
Test every layer before attackers do
Security teams log 54% of successful attacks and alert on just 14%. The rest move through your environment unseen.
The Picus whitepaper shows how breach and attack simulation tests your SIEM and EDR rules so threats stop slipping by detection.
Microsoft fixes KB5089549 Windows security update install issues
Windows 11 KB5089573 update released with performance improvements
Microsoft plans to improve Windows 11 driver quality in 2026
Microsoft confirms patching issues in restricted Windows networks
Windows 11 KB5089549 & KB5087420 cumulative updates released
The surprising part: AI inside Terminal could become a crutch. If it auto-fixes errors, we might forget the basics and grok command behavior less. Handy, yes, yet the real test is whether we still learn to reason from first principles.
Not a member yet? Register Now
Cisco warns of unpatched SD-WAN zero-day exploited in attacks
CISA: Hackers now exploit SolarWinds Serv-U flaw to crash servers
Chinese APT deploys new malware to keep access to hacked networks
Build cyber resilience with Wazuh: The open-source SIEM & XDR for proactive protection
New webinar: Behind-the-scenes of device code phishing kits
Your last pentest was 345 days ago. What changed since then?
Your AI tools are leaking sensitive data. Get a free audit.
The State of Healthcare Credential Exposure in 2026: Read the (Ungated) Report
Terms of Use - Privacy Policy - Ethics Statement - Affiliate Disclosure
Copyright @ 2003 - 2026 Bleeping Computer ® LLC - All Rights Reserved
Not a member yet? Register Now
Read our posting guidelinese to learn what content is prohibited.
