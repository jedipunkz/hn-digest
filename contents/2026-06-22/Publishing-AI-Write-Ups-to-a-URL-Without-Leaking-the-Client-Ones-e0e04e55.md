---
source: "https://bradystroud.dev/blogs/private-docs-vault"
hn_url: "https://news.ycombinator.com/item?id=48636487"
title: "Publishing AI Write-Ups to a URL – Without Leaking the Client Ones"
article_title: "Publishing AI Write-ups to a URL - Without Leaking the Client Ones | Brady Stroud"
author: "bradystroud"
captured_at: "2026-06-22T21:41:46Z"
capture_tool: "hn-digest"
hn_id: 48636487
score: 1
comments: 0
posted_at: "2026-06-22T21:29:52Z"
tags:
  - hacker-news
  - translated
---

# Publishing AI Write-Ups to a URL – Without Leaking the Client Ones

- HN: [48636487](https://news.ycombinator.com/item?id=48636487)
- Source: [bradystroud.dev](https://bradystroud.dev/blogs/private-docs-vault)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T21:29:52Z

## Translation

タイトル: AI 書き込みを URL に公開 – クライアントの書き込みを漏らさずに
記事のタイトル: AI の書き込みを URL に公開 - クライアントの書き込みを漏らさずに |ブレイディ・ストラウド
説明: ブレイディ・ストラウド

記事本文:
AI の書き込みを URL に公開 - クライアントの書き込みを漏らすことなく |ブレイディ ストラウド メイン コンテンツにスキップ ブレイディ ストラウド .dev について
すべてのエントリ Brady Stroud 2026 年 6 月 17 日 // ai、developer-tools、cloudflare、セキュリティ、パスキー、生産性 AI の書き込みを URL に公開 - クライアントの内容を漏らすことなく
私の仕事の多くは、バグの調査、移行計画、ステータス レポート、簡単なダッシュボードなど、短期間で作成するものです。正確には文書化されていません。むしろ、何かを理解するための紙の痕跡です。
最近では、作業が完了したら AI エージェントにそれぞれを単一の自己完結型 HTML ファイルに変換させ、それを PR にドロップしたり誰かに送信したりできる URL に公開します。 「共有可能にする」ステップは 1 つのコマンドです。
これはうまく機能しました。ちょうど、これらの書き込みの一部がクライアントの作業であり、公開 URL 上に置かれていることに気づくまででした。
これが全体のストーリーです。私がそれらをどのように公開するか、そしてデフォルトで非公開になるように公開側をどのように再構築するかです。
出発点は、小さな CLI、publish-plan でした。
パブリッシュプラン ./q3-report.html
# -> https://docs.strud.dev/q3-report HTML を Cloudflare R2 バケットにアップロードし、クリーンな URL を出力します。それでおしまい。
これが重要な理由は、エージェントがそれ自体を呼び出すことができるためです。 「...そしてレポートを公開する」でタスクを終了すると、エージェントからリンクが返されます。マークダウンをコピーしてドキュメント ツールに貼り付ける必要はなく、スクリーンショットも必要ありません。実際の URL で実際にレンダリングされたページです。
HTML と URL は、チャット ウィンドウでマークダウンの壁を毎回打ち破ります。適切にレンダリングされ、常に残り、ターミナルを開かない人にも送信できます。
魔法: それは単なる CLI ではなくスキルです
ここが自動的に感じられる部分です。 publish-plan は私が手動で実行するコマンドだけではなく、エージェント スキルでもあります。短い SKILL.md は、コマンド、発行方法などをエージェントに教えます。

名前空間の機能とプライバシーのルール。
だから私はそれについて何も説明する必要はありません。エージェントはジョブを終了し、結果を保持する価値があると判断し、スキルを読み取り、適切なコマンドを実行して、書き込みを適切な可視性を持つ適切な名前空間にドロップします。 「これを公開し、クライアントの作業を非公開にする」ということは、エージェントがすでに持っている知識です。
まさにこれが、次の部分が重要な理由です。エージェントがトリガーを引くのであれば、プロンプトでプライバシーを丁寧に提案することはできません。プライバシーはシステム自体によって強制される必要があります。
問題: すべてが公開されていた
内部的には、R2 バケットをカスタム ドメインで提供するという、最も単純な動作でした。オブジェクトをアップロードします。オブジェクトは docs.strud.dev/<key> にあります。
私が公開したすべてのものをルートにリストしてほしかったのです。しかし、R2 は CLI からオブジェクトをリストできず、カスタム ドメインのバケットはインデックス ページとして機能しないため、ルートは 404 だけになります。
さらに重要なのは、私が公開したものはすべて、URL を知っている人なら誰でもアクセスできることです。 Google によってインデックスは付けられていませんが ( noindex を設定します)、アクセス可能です。そのうちのいくつかはクライアントからの仕事でした。
共有したいプランには「URL でアクセス可能」で十分です。クライアントの仕事としては良くありません。隠蔽によるセキュリティではなく、実際のアクセス制御が必要でした。
修正は、バケットを直接提供するのをやめ、その前に小さな Cloudflare Worker を配置することでした。すべてのリクエストは、「この人はこれを見ることを許可されていますか?」という 1 つの質問をするコードを通じて実行されるようになりました。
この 1 つの動作により、ダム オブジェクト ストレージが実際のアプリに変わります。
ルートはバケットをリストし、ダッシュボードをレンダリングできます (CLI ではリストできませんが、ワーカーはリストできます)。
可視性チェックに合格するまでは何も提供されません。
デフォルトの答え: いいえ。特別な記載がない限り非公開です。
仕事の隣にプライバシーを設定
ファイルごとのプライバシーについては考えたくありませんでした。出版したいし、

正しいことが起こりますように。
したがって、プライバシーはリポジトリ内で一度宣言されます。各プロジェクト (またはクライアント フォルダー全体) には docs.json が含まれます。
// docs.json
{ "path": "client-work", "visibility": "private" } 公開すると、CLI は現在のディレクトリから進み、最も近い docs.json を見つけて使用します。クライアント フォルダーはプライベートとしてマークされます。私の個人的なことは公開されています。ドキュメントは、その可視性を備えた名前空間の下に自動的に配置されます。
cd ~/dev/some-client-project
計画の発行 ./findings.html Audits/2026/q3
# -> https://docs.stroud.dev/client-work/audits/2026/q3 (private) 最初のパスセグメントは名前空間であり、名前空間はプライバシー境界です。クライアントの作品は、公開元のフォルダーがすでに非公開になっているため、誤って公開されてしまうことはありません。
これはスキルの残りの半分です。スキルは公開することを認識し、 docs.json がどこでどのように非公開にするかを決定します。その間、エージェントは私がいなくても正しいことをしてくれます。
すべてのドキュメントは次の 3 つのいずれかに解決されます。
リンク - 共有リンクを持っている人全員
Worker は、有効な共有リンク → ドキュメントごとのオーバーライド → 名前空間ポリシー → それ以外の場合は拒否という順序で解決します。 「そうでない場合は拒否」が重要です。ポリシーのない野良ファイルは非公開であり、公開されることはありません。
所有者は私だけなので、毎回パスワードを入力するのは本当に嫌でした。したがって、所有者のアクセスはパスキー、つまり Touch ID です、基本的に。
面白い部分: ここには認証ライブラリがありません。 Cloudflare Worker内のWebAuthnは手動で非常に簡単に実行できることが判明しました。登録時に、ブラウザは標準形式の公開キーを渡します。ログイン時に、組み込みの WebCrypto API を使用して署名を検証します。
const key = await crypto.subtle.importKey(
"spki", publicKey, { name: "ECDSA",namedCurve: "P-256" }, false, ["verify"]
);
const ok = 待機します crypto.subtle.verify(
{ 名前: "ECDSA

"、ハッシュ: "SHA-256" }、キー、署名、signedData
);署名形式を変換し、2 つのステップの間に挑戦を続けるための小さなダンスがありますが、それが全体の形です。依存関係はありません。
結果: docs.stroud.dev 、Touch ID を開くと、名前空間ごとにグループ化され、それぞれがプライベート / パブリック / リンクのタグが付けられたすべてのダッシュボードが表示されます。
実際に誰かにプライベート ドキュメントを送信するまでは、デフォルトでプライベートにするのが最適です。そのために、私は共有リンク、つまり URL 内に署名された有効期限切れのトークンを作成します。
公開計画共有クライアント作業/監査/2026/q3 7d
# -> https://docs.stroud.dev/client-work/audits/2026/q3?k=<トークン> これはステートレスです - リンクのデータベースはありません。トークンは許可です。「このドキュメントは、この時点まで」と書かれており、ワーカーと私の CLI だけが知っている秘密で署名されています。ワーカーは署名と有効期限をチェックし、スコープ付き Cookie をドロップして、ドキュメントを提供します。期限が切れると機能しなくなります。
細かい点ですが、ドキュメントがすでに公開されている場合、共有しても単純な URL が得られるだけです。トークンは無意味です。
日常生活は何も変わりません。エージェントがドキュメントを書き、私がコマンドを 1 つ実行し、リンクを取得します。
変更されたのは以下のすべてです。
クライアントドキュメントはデフォルトで非公開であり、パスキーによってゲートされています。
public は名前空間ごとにオプトインであり、リポジトリで宣言されています。
そして共有は、期限切れのリンクを作成する意図的な行為であり、世界のデフォルトの状態ではありません。
私が繰り返し学び続けている教訓は、プライバシーは仕事の隣にあるとき、つまり、覚えておく必要がある設定ではなく、リポジトリ内のファイルであるときに最もよく機能します。そして、バケットの前にある小さなワーカーは、「URL 上のファイルの山」をクライアントの仕事で実際に信頼できるものに変えるのに十分です。
Brady によって構築および改造されました。お立ち寄りいただきありがとうございます。

## Original Extract

Brady Stroud

Publishing AI Write-ups to a URL - Without Leaking the Client Ones | Brady Stroud Skip to main content Brady Stroud .dev About
All entries Brady Stroud 17 June 2026 // ai, developer-tools, cloudflare, security, passkeys, productivity Publishing AI Write-ups to a URL - Without Leaking the Client Ones
A lot of my work throws off short-lived write-ups - an investigation into a bug, a plan for a migration, a status report, a quick dashboard. Not documentation exactly; more the paper trail of figuring something out.
These days I have an AI agent turn each one into a single self-contained HTML file when the work's done, then publish it to a URL I can drop in a PR or send to someone. The "make it shareable" step is one command.
That worked great - right up until I noticed some of those write-ups were client work, sitting on a public URL.
This is the whole story: how I publish them, and how I rebuilt the publishing side to be private by default.
The starting point was a tiny CLI, publish-plan :
publish-plan ./q3-report.html
# -> https://docs.stroud.dev/q3-report It uploads the HTML to a Cloudflare R2 bucket and prints a clean URL. That's it.
The reason it matters is that an agent can call it itself. I can end a task with "...and publish the report," and the agent hands me back a link. No copy-pasting markdown into a doc tool, no screenshots - a real rendered page at a real URL.
HTML plus a URL beats a wall of markdown in a chat window every time. It renders properly, it sticks around, and I can send it to someone who never opens a terminal.
The magic: it's a skill, not just a CLI
Here's the part that makes it feel automatic. publish-plan isn't only a command I run by hand - it's an agent skill . A short SKILL.md teaches the agent how to publish: the command, how the namespaces work, and the privacy rules.
So I never have to explain any of it. The agent finishes a job, decides the result is worth keeping, reads the skill, and runs the right command - dropping the write-up into the correct namespace with the correct visibility. "Publish this, and keep client work private" is knowledge the agent already has.
Which is exactly why the next part matters. If the agent is the one pulling the trigger, privacy can't be a polite suggestion in a prompt - it has to be enforced by the system itself.
The problem: everything was public
Under the hood it was the simplest thing that could work: an R2 bucket served on a custom domain. Upload an object, it's live at docs.stroud.dev/<key> .
I wanted the root to list everything I'd published. But R2 can't list objects from the CLI, and a bucket on a custom domain won't serve an index page - so the root just 404'd.
More importantly: anything I published was reachable by anyone with the URL. Not indexed by Google (I'd set noindex ), but reachable. And some of those were client work.
"Reachable by URL" is fine for a plan I want to share. It is not fine for a client's work. I needed real access control, not security-by-obscurity.
The fix was to stop serving the bucket directly and put a small Cloudflare Worker in front of it. Now every request runs through code that asks one question: is this person allowed to see this?
That one move turns dumb object storage into a real app:
the root can list the bucket and render a dashboard (the Worker can list, even though the CLI can't),
and nothing is served until it passes a visibility check.
Default answer: no . Private unless something says otherwise.
Privacy as config, sitting next to the work
I didn't want to think about privacy per file. I want to publish and have the right thing happen.
So privacy is declared once, in the repo. Each project - or a whole client folder - carries a docs.json :
// docs.json
{ "path": "client-work", "visibility": "private" } When I publish, the CLI walks up from the current directory, finds the nearest docs.json , and uses it. A client folder is marked private; my personal stuff is public. The doc lands under that namespace with that visibility - automatically.
cd ~/dev/some-client-project
publish-plan ./findings.html audits/2026/q3
# -> https://docs.stroud.dev/client-work/audits/2026/q3 (private) The first path segment is the namespace, and the namespace is the privacy boundary. Client work can't end up public by accident, because the folder it's published from already says private.
This is the other half of the skill: the skill knows to publish, and docs.json decides where and how private . Between them, the agent does the right thing without me in the loop.
Every doc resolves to one of three:
link - anyone holding a share link
The Worker resolves it in order: a valid share link → a per-doc override → the namespace policy → otherwise deny. "Otherwise deny" is the important bit. A stray file with no policy is private, never public.
I'm the only owner, and I really didn't want to type a password every time. So owner access is a passkey - Touch ID, basically.
The fun part: there's no auth library in here. WebAuthn inside a Cloudflare Worker turns out to be very doable by hand. On registration the browser hands you the public key in a standard format; on login you verify the signature with the built-in WebCrypto API:
const key = await crypto.subtle.importKey(
"spki", publicKey, { name: "ECDSA", namedCurve: "P-256" }, false, ["verify"]
);
const ok = await crypto.subtle.verify(
{ name: "ECDSA", hash: "SHA-256" }, key, signature, signedData
); There's a little dance to convert the signature format and to keep a challenge around between the two steps, but that's the whole shape of it. No dependencies.
The result: I open docs.stroud.dev , Touch ID, and I see a dashboard of everything - grouped by namespace, each tagged private / public / link.
Private-by-default is great until you actually want to send someone a private doc. For that I mint a share link : a signed, expiring token in the URL.
publish-plan share client-work/audits/2026/q3 7d
# -> https://docs.stroud.dev/client-work/audits/2026/q3?k=<token> It's stateless - there's no database of links. The token is the grant: it says "this doc, until this time," signed with a secret only the Worker and my CLI know. The Worker checks the signature and the expiry, drops a scoped cookie, and serves the doc. When it expires, it stops working.
A nice little detail: if the doc is already public, sharing just gives you the plain URL - a token would be pointless.
Nothing changed about the day-to-day: an agent writes a doc, I run one command, I get a link.
What changed is everything underneath:
client docs are private by default , gated by a passkey,
public is opt-in , per namespace, declared in the repo,
and sharing is a deliberate act that mints an expiring link - not the default state of the world.
The lesson I keep relearning: privacy works best when it lives next to the work - a file in the repo, not a setting I have to remember - and a tiny Worker in front of a bucket is enough to turn "a pile of files on a URL" into something I'd actually trust with client work.
Built and tinkered with by Brady. Thanks for stopping by.
