---
source: "https://www.octopunk.io/blog/introducing-octopunk"
hn_url: "https://news.ycombinator.com/item?id=48571849"
title: "Show HN: OctoPunk – a full-featured, AI-integrated GitHub client"
article_title: "Introducing OctoPunk: who we are, what we're building, and how to try it"
author: "ldelossa"
captured_at: "2026-06-17T16:21:37Z"
capture_tool: "hn-digest"
hn_id: 48571849
score: 4
comments: 0
posted_at: "2026-06-17T15:28:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: OctoPunk – a full-featured, AI-integrated GitHub client

- HN: [48571849](https://news.ycombinator.com/item?id=48571849)
- Source: [www.octopunk.io](https://www.octopunk.io/blog/introducing-octopunk)
- Score: 4
- Comments: 0
- Posted: 2026-06-17T15:28:57Z

## Translation

タイトル: Show HN: OctoPunk – フル機能の AI 統合 GitHub クライアント
記事のタイトル: OctoPunk の紹介: 私たちが誰であるか、何を構築しているのか、そしてそれを試す方法
説明: 私たちは
HN テキスト: こんにちは皆さん。多くの優れたソフトウェアと同様、私たちはフラストレーションから OctoPunk を構築しました。私たちは、何らかの理由で、非常に長い間、専門的な仕事のための主要なコード フォージとして GitHub を使用してきた開発者の集まりです。私たちは、GitHub が成長し、停滞し、機能要求を無視し、フル機能のデスクトップ クライアントをまったく作成しないのを見てきました。最後に、もう十分だったので、独自のものを構築することに決め、OctoPunk が誕生しました。 OctoPunk は、GitKraken のような Git GUI ラッパーではなく、コード レビューをサポートしていない GitHub の公式クライアントのような中途半端に実装されたデスクトップ クライアントでもありません。また、コミットごとのレビューのような完全なワークフローを無視するものでもありません。VS Code 拡張機能 ( https://github.com/microsoft/vscode-pull-request-github/issu... )。 OctoPunk は、デスクトップ クライアントとしての GitHub の機能の 95% を表面化するフル機能の GitHub クライアントで、最新のエディターや開発者ツールにすぐに馴染みのある UI と UX を通じて表示されます。カスタマイズ可能なダッシュボード、クイック検索ウィジェット、デスクトップ通知、さらにオプションでユーザーが見ているものを理解し、コメントや回答の下書きなどの説明を支援できる AI アシスタントも提供します。 OctoPunk の名前は、私たちが構築しているソフトウェアだけではなく、会社の精神から名付けられました。私たちは小さなチームであり、大きな技術慣行に同意していません。私たちは小規模で独立し、機敏に行動し続けるつもりです。とにかく、デスクトップネイティブで生産性を重視した GitHub クライアントをお探しの場合は、ぜひ試してみてください。現在ベータ版であり、フィードバック、機能リクエスト、バグレポートを募集しています。これらはすべて便利です。

OctoPunk自体からのみ送信されます。 OctoPunk はすべてパブリック リポジトリでは無料ですが、プライベート リポジトリではサブスクリプションが必要です。これが当社の唯一の収益化モデルです。ダーク パターンやデータ収集はなく、常にシンプルに保つことをお約束します。私たちについての詳細とベータ版のダウンロード方法については、こちらをご覧ください: https://www.octopunk.io/blog/introducing-octopunk
ここでコミュニティと交流してください: https://github.com/octopunkio/community

記事本文:
OctoPunk の紹介: 私たちが誰であるか、何を構築しているのか、そしてそれを試す方法
OctoPunk について ブログ ドキュメントと機能 価格の比較
GitHub でログイン ← ブログに戻る OctoPunk の紹介: 私たちが誰であるか、何を構築しているのか、そしてそれを試す方法
私たちは10年間、ブラウザタブの中で生きてきました。私たちは出ていきたいのです。
GitHub で作業が行われるため、作業は Chrome、Firefox、または (震える) Internet Explorer で止まってしまいます。タブ 14 はプル リクエスト、タブ 22 は 3 日前にトリアージするつもりだった問題で、通知ベルは赤ですが、これはまったく意味がありません。私たちはそれを無視するように自分自身を訓練してきました。
OctoPunk は、私たちがもう我慢できなくなったときに構築したものです。 「Create PR」ボタンが追加された git GUI ではなく、GitHub プラットフォームの実際のクライアントです。全部。実際に作業するデスクトップ上。
OctoPunk が最初に修正するのは通知です。忙しいリポジトリに取り組んだことがある人なら誰でも、この状況が起こるのを見たことがあるでしょう。 GitHub は、本物のフィルターを必要とせず、実際に関心のある事柄を表面化しにくくする通知システムを提供します。たとえば、自分が所属しているチームに割り当てられたリクエストを確認します。結局、ベルを完全に消すか、点滅していないふりをすることになります。
OctoPunk では、問題の検索ですでに知っているのと同じクエリ言語であるクエリによってフィルターを構築できます。各フィルターには名前、色、および購読ボタンが与えられます。一致した通知は、色の付いた錠剤とともにダッシュボードの受信トレイに表示されるため、それぞれがどのようなものであるか一目でわかります。通知ベルに腹を立てなかったのは久しぶりです。
詳細については、通知ドキュメントを参照してください。
github.com を開くと、最初にフィードが表示されます。アルゴリズムによる星のスクロールが続き、「あなたがフォローしている誰かが 3 日前に何かをフォークしました」。 2010 年にそれは理にかなっていました。

ery Web 製品はソーシャル ネットワークになろうとしていました。あなたの仕事がコードを出荷することであれば、2026 年には意味がありません。 GitHub は Instagram ではありませんし、星空観察活動のタイムラインはツールではありません。
OctoPunk はフィードではなくダッシュボードを開きます。自分で組み立てる、カスタマイズ可能なヘッドアップ ディスプレイ。ウィジェットとは、検索クエリ、アクティビティ フィード (必要な場合)、GitHub アクティビティ グラフなどです。それらをドラッグして名前を付け、実際に関心のある作業の流れに一致するレイアウトを構築し、残りは無視します。検索ウィジェットは完全な GitHub 検索構文を使用するため、「組織内で 3 日間レビューされていないすべての開いている PR」を表現し、そのまま画面上に残すことができます。
アプリの残りの部分のナビゲーションは同じ雰囲気を保ちます。 GitHub のクロスリポジトリ検索、最近アクセスしたアイテム、開いているタブのオムニバー。コンテキスト認識型のコマンド パレットにより、課題やユーザー ページとは異なるコマンドが PR 上で提供されます。どちらもあいまい検索可能で、どちらもキーボードファーストです。
詳細については、 ダッシュボード 、 オムニバー 、およびコマンド パレットのドキュメントを参照してください。
私たちが最も重視したのはコードレビューです。 GitHub Web UI での差分を読み取るのは 10 行では問題ありませんが、200 行では悲惨です。 OctoPunk はローカル LSP をファイル ビューアに接続します。マウスを移動して型情報を表示し、定義にジャンプし、シンボルを検索するなど、すべてが同じウィンドウ内でレビュー中のコード上で機能します。 Clangd の apply_commands.json などのリポジトリごとの設定パネルがあるため、少しセットアップが必要な LSP は、ラッパー スクリプトの作成に駆り出されることなく、このパネルを入手できます。
詳細については、コード インテリジェンスのドキュメントを参照してください。
OctoPunk には AI アシスタントが付属しています。全員が1つずつ発送しています。興味深い質問は、それをどうするかということです。
Alt を押したままページ上の何かをクリックすると、アシスタントに正確なコンテキストが与えられます。コメント、PR、ファイル、

ユーザーです。そこから具体的なことを尋ねることができます。 「この PR は実際に何を変えるのでしょうか?」 「このスレッドへの返信の下書きを作成してください。」 「この会話を要約してください。」選択がなければ、ページ全体をコンテキストとして取得し、それを把握します。コードレビュー、コメントの下書き、まだ持っていないコードの取得などのためのツールがいくつかあります。独自の API キー (Anthropic、Gemini、または OpenAI) を使用することで、信頼できるプロバイダーとの会話が継続されます。
詳細については、AI アシスタントのドキュメントを参照してください。
上記の機能は、私たちが最初に強調したい機能です。それらが製品のすべてではありません。 OctoPunk は、プル リクエスト、問題、プロジェクト、ディスカッション、リリース、アクション、CI ワークフローもエンドツーエンドで処理します。これらはいずれも GitHub Web UI のクローンではありません。プラットフォームにサーフェスがある場所であればどこでも、OctoPunk は最新の開発者ツールが動作するようにそれを作り直します。大きな PR と問題を解決するための概要。コメント、ファイル、コードの間を移動するときにも残るコンテキスト。残りは今後の投稿に当てます。
GitHub に真剣に時間を費やす場合、ブラウザは間違ったインターフェイスです。 OctoPunk は、デスクトップ上で実際の G​​itHub プラットフォームを提供する唯一のツールです。優先順位を付けることができる通知。 PR を 1 つのウィンドウで確認できます。問題、プロジェクト、ディスカッション、アクション、および実際の LSP によるコードレビューの表面。ブラウザーのタブではなく、すべてがキーボードファーストです。
ホームページで、ダウンロード ボタンまでスクロールし、[ベータ版をお探しですか?] をクリックします。その下にリンクがあります。これにより、ベータ チャネルにアクセスし、Mac、Linux (フラットパック)、および Windows 用の最新のベータ ビルドが表示されます。 3つとも。
パブリックリポジトリは無料です。プライベート リポジトリには Pro サブスクリプションが必要です。これはアカウント ダッシュボードから設定できます。
入場する際に知っておくべきことがいくつかあります。
ベータ版のビルドは何かを壊します。それが契約だ。
アプリのデータ ディレクトリをリセットする場合があります

リリース間でRY。これにより設定とサインインがクリアされるため、再度ログインする必要があります。楽しみのためにやらないように努めていますが、そうなってしまうのです。
問題をコミュニティ リポジトリに記録します。バグ、機能のリクエスト、ただのご挨拶、すべて大歓迎です。
これは何で、何がそうではないのか
OctoPunk は git GUI ではありません。素晴らしいものはすでにあるので、ぜひ使ってみてください。 lategit 、 tower 、 fork は、これまでの「GitHub クライアント」よりも優れた git を実行する専用ツールです。インタラクティブなリベース、競合編集、非難サーフィン、ハンクごとのステージング、これらは機能します。それらを使用してください。
OctoPunk は、git と GitHub を別のものとして扱う開発者を対象としています。 Git はバージョン管理ツールです。 GitHub はその上のコラボレーション プラットフォームです。私たちは両方になろうとはしません。コミットとブランチを形成するときは、お気に入りの git ツールを使用してください。 PR の確認、通知のトリアージ、問題の解決、CI の監視など、プラットフォームを操作するときに OctoPunk を使用します。この 2 つは共存します。それがポイントです。
私たちは誰ですか：オクトパンクのパンク
最後に、ソフトウェア会社としての私たちについてお話しましょう。
私たちは小規模なチームであり、10 年以上にわたるテクノロジーの誇大宣伝サイクルからくる特定の意味で疲れています。暗号、NFT、Web3、メタバース、サービスとしてのエージェント、その他すべてのスタートアップは、[すべて] の未来として宣伝しています。私たちはそのすべてを座って見ました。私たちが知っている人の基本的な仕事のやり方は、どれも変わりませんでした。
暗いパターンはありません。 OctoPunk は小言を言ったり、データを販売したり、必要のない情報を要求したりしません。 GitHub に一度サインインするだけで完了です。私たちは明かりを灯し続けるために必要なものを集めます。それ以外は何も集めません。機能として見せかけたエンゲージメント指標や、ユーザーを犠牲にしてチャートを上昇させるための「成長戦略」は存在しません。
はい、AI があります。独自の誇大宣伝サイクルの真っ只中にあり、熱い空気が印象的です。それも本当にあなたです

見たことのない広大な PR を理解したり、読む時間がないスレッドを要約したりするなど、いくつかの特定の場所で役立ちます。それらは私たちが置いた場所です。バックグラウンドでループするエージェント、「AI ファースト」のブランド変更、追加料金はありません。自分のキーを持参し、役立つ場合はそれを使用し、役に立たない場合は無視します。
OctoPunk のパンクは、上記のすべてにノーと言う部分です。
何が壊れているのか教えてください。何が足りないのか教えてください。少し迷惑なことを教えてください。ベータ入力は私たちが現在持っている中で最も便利なものであり、摩擦ができる限りゼロに近づくようにファイリングワークフローをOctoPunkに組み込みました。
アプリ内からコマンド パレットを開き、「⌘+P バグを報告」を実行してバグを報告します。または ⌘+P 構築してほしい機能を提案する機能を提案します。どちらもコミュニティ リポジトリにあります。どちらも読まれます。
私たちはしばらく OctoPunk を構築するために OctoPunk をフルタイムで使用してきました。ブラウザのタブの習慣がようやく解消されたのは十分に良いことです。ベータ版は、タブの墓場にうんざりしているすべての人のためのものです。

## Original Extract

We

Hello folks. Like a lot of great software, we built OctoPunk out of frustration. We are a set of developers who, for one reason or another, have been using GitHub as our primary code forge for professional work for a very long time. We've seen GitHub grow, stagnate, ignore feature requests, and simply never produce a full-featured desktop client. Finally, we had enough, decided to build our own, and OctoPunk was born. OctoPunk is not a git GUI wrapper like GitKraken, not a half-implemented desktop client like GitHub's official client, which doesn't support code review, and not something that ignores complete workflows like per-commit reviews, looking at you, VS Code extension ( https://github.com/microsoft/vscode-pull-request-github/issu... ). OctoPunk is a full-featured GitHub client that surfaces 95% of GitHub's functionality as a desktop client, presented through a UI and UX that feels instantly familiar to modern editors and developer tools. We provide a customizable dashboard, quick search widgets, desktop notifications, and even an AI assistant that optionally understands what you're looking at and can help explain comments, draft responses, and more. OctoPunk got its name not just from the software we are building, but from the mentality of the company. We are a tiny team that doesn't subscribe to big tech practices. We plan on staying small, independent, and nimble. Anyway, if you're looking for a desktop-native and productivity-focused GitHub client, give us a try. We are currently in beta and looking for feedback, feature requests, and bug reports, all of which can be conveniently sent from within OctoPunk itself. All of OctoPunk is free for public repositories and requires a subscription for private ones. That's our only monetization model: no dark patterns, no data harvesting, and we promise to always keep it simple. Learn more about us and how to download the beta here: https://www.octopunk.io/blog/introducing-octopunk
Interact with our community here: https://github.com/octopunkio/community

Introducing OctoPunk: who we are, what we're building, and how to try it
OctoPunk About Blog Docs and features Compare Pricing
Login with GitHub ← Back to blog Introducing OctoPunk: who we are, what we're building, and how to try it
We've been living inside a browser tab for ten years. We want out.
GitHub is where the work happens, so the work has been stuck in Chrome, Firefox, or ( shivers ) Internet Explorer. Tab 14 is a pull request, tab 22 is an issue we meant to triage three days ago, the notifications bell is red and that means absolutely nothing. We've trained ourselves to ignore it.
OctoPunk is what we built when we couldn't take it anymore. Not a git GUI with a "Create PR" button bolted on, a real client for the GitHub platform. The whole thing. On the desktop, where you actually work.
The first thing OctoPunk fixes is notifications. Anyone who has worked on a busy repo has seen this play out. GitHub gives you a firehose, no real filters, and a notification system that makes the things you actually care about hard to surface. Review requests assigned to a team you're on, for example. You end up either turning the bell off entirely or pretending it isn't blinking.
OctoPunk lets you build filters by query, same query language you already know from issue search. Each filter gets a name, a color, and a subscribe button. The matched notifications land in your dashboard inbox with a colored pill so you can tell at a glance what kind of thing each one is. It is the first time in a long time we haven't been mad at the notifications bell.
More in the notifications docs .
Open github.com and the first thing you see is a feed. An algorithmic scroll of stars, follows, and "someone you follow forked something three days ago". That made sense in 2010, when every web product was trying to be a social network. It does not make sense in 2026 if your job is shipping code. GitHub is not Instagram, and a timeline of stargazer activity is not a tool.
OctoPunk opens to a dashboard, not a feed. A customizable heads-up display you build yourself. Widgets are search queries, your activity feed (if you still want one), the GitHub activity graph, things like that. Drag them around, name them, build a layout that matches the streams of work you actually care about, ignore the rest. Search widgets take the full GitHub search syntax so you can express "all open PRs in my org that haven't been reviewed in three days" and just leave it on screen.
Navigation around the rest of the app keeps the same vibe. An omnibar for cross-repo GitHub search, recently visited items, and open tabs. A context-aware command palette that gives you different commands on a PR than on an issue or a user page. Both fuzzy searchable, both keyboard-first.
More in the dashboard , omnibar , and command palette docs.
Code review is the thing we cared about most. Reading diffs in the GitHub web UI is fine for ten lines and miserable for two hundred. OctoPunk wires a local LSP into the file viewer. Hover for type info, jump to definitions, search symbols, all of it works on the code you are reviewing, in the same window. There is a per-repo settings panel for things like compile_commands.json for clangd , so the LSPs that need a bit of setup get one without sending you off to write a wrapper script.
More in the code intelligence docs .
OctoPunk ships with an AI assistant. Everyone is shipping one. The interesting question tho, is what to do with it.
Hold Alt , click something on the page, and you've given the assistant a precise piece of context. A comment, a PR, a file, a user. From there you can ask specific things. "What is this PR actually changing?" "Draft a reply to this thread." "Summarize this conversation for me." Without a selection it gets the whole page as context and figures it out. It has tools for code review, comment drafting, fetching code it does not already have, and a few other things. You bring your own API key (Anthropic, Gemini, or OpenAI) so the conversation stays on a provider you trust.
More in the AI assistant docs .
The features above are the ones we wanted to call out first. They aren't the whole product. OctoPunk also handles pull requests end-to-end, issues, projects, discussions, releases, actions and CI workflows. None of them are clones of the GitHub web UI. Anywhere the platform has a surface, OctoPunk reworks it the way modern developer tools work. Outlines for navigating big PRs and issues. Context that stays with you when you jump between comments, files, and code. The rest is for future posts.
If you spend serious time on GitHub, the browser is the wrong interface. OctoPunk is the only tool that gives you the actual GitHub platform on the desktop. Notifications you can triage. PRs you can review in one window. Issues, projects, discussions, actions, and a code-review surface with real LSPs. Keyboard-first, the whole thing, not a browser tab in sight.
On the home page, scroll to the download buttons and click the "Looking for the beta?" link underneath them. That puts you on the beta channel and shows you the latest beta builds for Mac, Linux (flatpak), and Windows. All three.
Free for public repos. Private repos need a Pro subscription, which you can set up from your account dashboard.
A couple of things to know going in.
Beta builds break things. That's the deal.
We may reset the app data directory between releases. That clears your settings and your sign-in, you'll need to log in again. We try not to do it for fun but it happens.
File issues at the community repo . Bugs, feature requests, just saying hi, all welcome.
What this is, and what it is not
OctoPunk is not a git GUI. There are great ones already, and you should be using one. lazygit , tower , fork are purpose-built tools that do git better than any "GitHub client" ever will. Interactive rebase, conflict editing, blame surfing, hunk-by-hunk staging, the works. Use them.
OctoPunk is geared for developers who treat git and GitHub as separate things. Git is the version-control tool. GitHub is the collaboration platform on top of it. We don't try to be both. Use your favorite git tool when you're shaping commits and branches. Use OctoPunk when you're working with the platform: reviewing PRs, triaging notifications, navigating issues, watching CI. The two coexist. That's the point.
Who we are: the punk in OctoPunk
Let's close this out with who we are as a software company.
We're a small team, tired in the specific way that comes from a decade-plus of tech-hype cycles. Crypto, NFTs, web3, the metaverse, agents-as-a-service, every other startup pitched as the future of [everything]. We sat through all of it. None of it changed how anyone we know fundamentally does their job.
No dark patterns. OctoPunk doesn't nag, doesn't sell your data, doesn't ask for information it doesn't need. Sign in once with GitHub, that's the deal. We collect what we need to keep the lights on and nothing else. No engagement metrics dressed up as features, no "growth tactics" that exist to make a chart go up at your expense.
Yes, we have AI. It's in the middle of its own hype cycle and the hot air is impressive. It's also genuinely useful in a few specific places, like understanding a sprawling PR you've never seen or summarizing a thread you don't have time to read. Those are the places we put it. No agents looping in the background, no "AI-first" rebrand, no charging extra for it. Bring your own key, use it when it helps, ignore it when it doesn't.
The punk in OctoPunk is the part that says no to all of the above.
Tell us what's broken. Tell us what's missing. Tell us what's slightly annoying. Beta input is the most useful thing we have right now, and we built the filing workflow into OctoPunk so the friction is as close to zero as we could make it.
From inside the app, open the command palette and run ⌘+P Report a bug to file a bug. Or ⌘+P Suggest a feature to pitch one you want us to build. Both land in our community repo. Both get read.
We've been using OctoPunk full-time to build OctoPunk for a while now. It is good enough that the browser tab habit is finally breaking. The beta is for everyone else who is sick of the tab graveyard.
