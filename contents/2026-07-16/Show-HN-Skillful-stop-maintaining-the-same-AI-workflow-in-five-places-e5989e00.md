---
source: "https://skillful.md/"
hn_url: "https://news.ycombinator.com/item?id=48932688"
title: "Show HN: Skillful, stop maintaining the same AI workflow in five places"
article_title: "Skillful | Reusable AI workflows, managed on disk."
author: "mastermindzh"
captured_at: "2026-07-16T11:30:15Z"
capture_tool: "hn-digest"
hn_id: 48932688
score: 1
comments: 0
posted_at: "2026-07-16T10:35:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Skillful, stop maintaining the same AI workflow in five places

- HN: [48932688](https://news.ycombinator.com/item?id=48932688)
- Source: [skillful.md](https://skillful.md/)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T10:35:53Z

## Translation

タイトル: Show HN: 上手、同じ AI ワークフローを 5 か所で維持するのはやめよう
記事タイトル: 上手 |再利用可能な AI ワークフローはディスク上で管理されます。
説明: Skillful は、再利用可能な AI スキルとエージェントを実際のフォルダーとして構築し、それらを Claude Code、Codex、Cursor、Copilot、Gemini、Junie、および opencode にインストールするためのローカル ファースト デスクトップ アプリです。
HN テキスト: 皆さん、ここ 1 年で、Claude コードや Codex などの複数の AI ツールで同じ AI ワークフローを維持していることに気づきました。 1 つのワークフローを改善するたびに、それをコピーした他の場所を覚えておかなければなりませんでしたが、ほとんどの場合覚えていませんでした...そこで、すぐにすべてを一緒にシンボリックリンクする「ドットファイル」ルートに進みました。そこで、私は CLI に慣れているのでこれでうまくいきますが、他の人はそれほどうまくいかないかもしれないと思いました。いろいろな人と話したり、ディスカッションしたりして、最終的にSkillfulを作ることになりました。 Skillful は、プロンプトやスキルを単一の AI ツール内に存在するものとして扱うのではなく、ディスク上の通常のフォルダーとして扱います。スキルは、SKILL.md または AGENT.md と、必要なもの、例、スクリプト、スクリーンショット、ドキュメントなどで構成されます。このアプリケーションでは、次のことが可能です。 - 1 つのローカル ライブラリからすべてを管理
- コピーの代わりにリンクを使用して、同じワークフローを複数の AI ツールにインストールします
- Github を使用してライブラリを複数のコンピューターに同期します
- スキルを含むパブリック GitHub リポジトリをインポートする
- 壊れたインストールを検出して修復します。すべてがマシン上に残ります。アカウント、テレメトリ、バックエンドはありません。これは基本的に、シンボリックリンクを備えたファイルシステムを中心とした GUI です。アプリケーションがいつか消えてしまったとしても、ファイルが引き続き役立つようにしたかったからです。また、誰でも公開または投稿できる再利用可能なスキル パックを備えた「素晴らしい」リポジトリも開始しました: https://github.c

om/Mastermindzh/awesome-killed いくつかの点についてフィードバックをいただければ幸いです。「1 つのソース、多数の AI ツール」というアプローチは、あなたにとっても理にかなっていますか?
これらをプレーンな Git リポジトリとして管理したいですか?
次にサポートされるべき AI ツールは何ですか?
すでに再利用可能なスキルを使用している場合、現在はそれらをどのように整理していますか?ご意見をお待ちしております! PS: 私が Skillful を構築している間に、Vercel Labs が「npx スキル」も構築したことは知っています。これは代替手段だと思います :P

記事本文:
上手
再利用可能なワークフロー、プレーン ファイル。
特長
ワークフロー
プライバシー
よくある質問
ダウンロード
所有するファイルから再利用可能な AI ワークフローを構築します。
プロンプト、エージェント、スクリーンショット、スクリプト、メモを巧みに変換します。
ファイルシステムライブラリを検査および保守できます。フォルダーを直接編集したり、マークダウンをプレビューしたり、
公開スキル リポジトリをインポートし、同じアイテムを必要とするすべての AI ツールにインストールします。
ダウンロード
macOS
窓
Linux
ソース
無料
同じスキルまたはエージェントを Claude Code、Codex、Cursor、Copilot、Gemini、
Junie、opencode、またはフォルダーベースの宛先。コピーもドリフトもありません。
プロンプト、スクリーンショット、スクリプト、サポートノートがスキルの隣に表示されます。
それらを使用します。
「SKILL.md」と「AGENT.md」のワークフローをフィルターとバッジとともに 1 つのライブラリに保持
ワークスペースを分割せずに違いを明確にします。
壊れたシンボリックリンクを見つけてクリックするだけで修復し、スキルが壊れたらきれいに削除します。
その使用期間を超えてしまいました。
作業に応じて変化するワークフロー向けに構築されています。
エージェントは 1 つのツール内に存在します。プロンプトがあちこちに散らばっている
~/notes/ 、 ~/dotfiles/ 、および 4 つの Obsidian 保管庫。
Codex には、そのうちの半分の独自のコピーがあります。
それらのコピーのうち 2 つは期限切れです。どれか覚えていないでしょう。
スクリーンショットとヘルパー スクリプトは、プロンプトから 3 つ離れたフォルダーにあります。
それらを使用します。
図書館がひとつ。すべてのスキルまたはエージェントは、読み取り可能なフォルダーに必要なものを加えたものです。
各項目を、それが表示されるツールにリンクします。
Skillful でソースを調整すると、リンクされたインストールが自動的に更新されます。
壊れたリンクがサイドバーに表示され、クリックするだけで修復されます。
Skillful はソースが入手可能で、無料です。スキル、エージェント、メモはそのまま残ります
あなたのマシン: アカウントもテレメトリもサーバーもありません。あなたがフォルダーの所有者であり、
上手に消えても、ライブラリはすべてローカルにあるため、まだ機能します。
人々が尋ねる質問

ダウンロードする前に。
はい。このソースは、MIT の Functional Source License 1.1 に基づいて入手できます。
将来のライセンス。個人、社内、教育、研究、および専門サービス
使用は許可されています。競合する商用製品やサービスは存在しません。
いいえ、すべてがローカル ファイル システムに残ります。 Skillful にはバックエンドがなく、
分析。唯一の送信トラフィックは GitHub 宛てです。
リリース フィード、およびインポートするために明示的に選択したソース アーカイブ。を参照してください。
完全なリストについては、プライバシー セクションを参照してください。
Skillful は、Claude Code、Codex、Cursor、GitHub 用のワンクリック インストーラーを同梱しています
Copilot、Gemini CLI、Junie、オープンコード。独自のカスタム インストールを追加することもできます
ターゲット: Skillful をディスク上の任意のフォルダーにポイントすると、それを
内蔵のもの。スキルごとまたはコレクションごとに、好きなだけ。
Skillful は、GitHub Releases で Windows、macOS、Linux ビルドを公開しています。 Linux
リリースには、AppImage、deb、rpm、pacman、Snap、および Flatpak アーティファクトが含まれているため、
システムに合ったインストール スタイルを選択できます。
はい。 Skillful は直接インポート リンクをサポートしているため、リポジトリの README でアプリを開くことができます
GitHub リポジトリはインポート ダイアログにすでに入力されています。それが公になるのです
スキル ライブラリは、URL を手でコピーすることなく試すのがはるかに簡単です。
発電機を調べてみる
自分自身のものを作成します。
ただし、ファイルはすでにバックアップされています。ライブラリは単なるフォルダーですので、
git init を実行してプライベート リポジトリにプッシュし、通常のリポジトリと同期します
バックアップ ツールを使用するか、コピーを USB ドライブにドラッグします。独占的なものは何もありません
エクスポートするだけで、マシンを切り替えたときに移行するものは何もありません。
上手はローカルファースト。 Skillful アカウントも Skillful サーバーも存在しません。
あらゆる種類の分析またはテレメトリー。アプリは、小さくて予測可能な一連の要素を作成します。
アウトバウンド ネットワーク リクエスト、すべて GitHub へ、すべて

あなたが尋ねたときだけ
またはアプリを最新の状態に保つには:
アップデートチェック。デスクトップアプリのチェック
新しいバージョンについては、github.com/Mastermindzh/skilled/releases を参照してください。
[更新] タブを開くか、起動時に開きます。それ以降は識別情報は送信されません
GitHub が HTTPS リクエストに関してすでにログに記録しているもの。
GitHub のインポート。公開リポジトリをインポートする場合、Skillful
リポジトリのソース アーカイブをから直接ダウンロードします。
codeload.github.com 。何もアップロードされていません。
ディープリンク。 skilled://import?... リンクを開く
インポートダイアログをローカルで事前入力するだけです。クリックするまでリクエストは行われません
インポートします。
それがリスト全体です。スキルフォルダー、ドラフト、設定、ライブラリのパス
決してディスクから離れないでください。これを読むことで確認できます
の
ソースから、または選択したツールを使用してネットワーク トラフィックを監視することによって。
AI ワークフローを再利用可能にする準備はできていますか?
Skillful をダウンロードしてライブラリにポイントすると、5 か所で同じワークフローを維持するのがやめられます。

## Original Extract

Skillful is a local-first desktop app for building reusable AI skills and agents as real folders, then installing them into Claude Code, Codex, Cursor, Copilot, Gemini, Junie, and opencode.

Hey all, Over the past year I noticed I was maintaining the same AI workflows in multiple AI tools such as Claude code and Codex. Whenever I improved one workflow, I had to remember everywhere else I'd copied it and most of the time I didn't... so I quickly went the "dotfiles" route where I symlinked everything together. Then I thought, this works for me since I'm used to my CLI but others might not have such a good time. Had a few talks with people, had some discussions, and in the end I ended up building Skillful. Instead of treating prompts or skills as something that lives inside a single AI tool, Skillful treats them as ordinary folders on disk. A skill consists of a SKILL.md or AGENT.md together with whatever else it needs, examples, scripts, screenshots, documentation, etc. The application lets you: - manage everything from one local library
- install the same workflow into multiple AI tools using links instead of copies
- Sync your library to multiple computers with Github
- import public GitHub repositories containing skills
- detect and repair broken installs Everything stays on YOUR machine. There are NO accounts, NO telemetry, and NO backend. It's essentially a GUI around a filesystem with symlinks, because I wanted the files to remain useful even if the application disappeared one day. I also started an "awesome" repository with reusable skill packs that anyone can publish or contribute to: https://github.com/Mastermindzh/awesome-skillful I'd love feedback on a few things: Does the "one source, many AI tools" approach make sense for you too?
Would you rather manage these as plain Git repositories?
What AI tools should be supported next?
If you already use reusable skills, how are you organising them today? Looking forward to hearing what you think! PS: I know that Vercel Labs also built `npx skills` whilst I was building Skillful, I think of this as an alternative :P

Skillful
Reusable workflows, plain files.
Features
Workflow
Privacy
FAQ
Download
Build reusable AI workflows from files you own.
Skillful turns prompts, agents, screenshots, scripts, and notes into a
filesystem library you can inspect and maintain. Edit folders directly, preview markdown,
import public skill repos, and install the same item into every AI tool that needs it.
Download Skillful
macOS
Windows
Linux
Source
Free
Install the same skill or agent into Claude Code, Codex, Cursor, Copilot, Gemini,
Junie, opencode, or any folder-based destination. No copies, no drift.
Prompts, screenshots, scripts, and supporting notes live next to the skill that
uses them.
Keep `SKILL.md` and `AGENT.md` workflows in one library, with filters and badges
that make the difference clear without splitting your workspace.
Spot broken symlinks, repair them in a click, and remove cleanly when a skill has
outlived its use.
Built for workflows that change as you work.
Your agents live inside one tool. Your prompts are scattered across
~/notes/ , ~/dotfiles/ , and four Obsidian vaults.
Codex has its own copy of half of them.
Two of those copies are out of date. You don't remember which.
Screenshots and helper scripts are three folders away from the prompt that
uses them.
One library. Every skill or agent is a readable folder plus whatever it needs.
You link each item into the tools that should see it.
Tweak the source in Skillful, and linked installs automatically update.
Broken links surface in the sidebar and repair in a click.
Skillful is source-available and free. Your skills, agents, and notes stay on
your machine: no account, no telemetry, no servers. You own the folder, and if
Skillful ever disappears your library still works, because it's all local.
Questions people ask before downloading.
Yes. The source is available under the Functional Source License 1.1 with an MIT
future license. Personal, internal, education, research, and professional-services
use are allowed; competing commercial products or services are not.
No. Everything stays on your local filesystem. Skillful has no backend and no
analytics. The only outbound traffic is to GitHub: update checks against the
release feed, and the source archives you explicitly choose to import. See the
Privacy section for the full list.
Skillful ships one-click installers for Claude Code, Codex, Cursor, GitHub
Copilot, Gemini CLI, Junie, and opencode. You can also add your own custom install
targets: point Skillful at any folder on disk and it'll treat it like the
built-in ones. As many as you like, per skill or per collection.
Skillful publishes Windows, macOS, and Linux builds on GitHub Releases. Linux
releases include AppImage, deb, rpm, pacman, Snap, and Flatpak artifacts so you
can pick the install style that matches your system.
Yes. Skillful supports direct import links, so a repository README can open the app
with a GitHub repository already filled into the import dialog. That makes public
skill libraries much easier to try without copying URLs around by hand.
Check out the generator
to create your own.
However you already back up files. Your library is just a folder, so
git init and push it to a private repo, sync it with your usual
backup tool, or drag a copy onto a USB drive. There's nothing proprietary to
export and nothing to migrate when you switch machines.
Skillful is local-first. There is no Skillful account, no Skillful server, and no
analytics or telemetry of any kind. The app makes a small, predictable set of
outbound network requests, all of them to GitHub, all of them only when you ask
for them or to keep the app up to date:
Update checks. The desktop app checks
github.com/Mastermindzh/skillful/releases for new versions when you
open the Updates tab or on launch. No identifying information is sent beyond
what GitHub already logs for any HTTPS request.
GitHub imports. When you import a public repository, Skillful
downloads the repository's source archive directly from
codeload.github.com . Nothing is uploaded.
Deep links. Opening a skillful://import?... link
only prefills the import dialog locally. No request is made until you click
Import .
That's the entire list. Your skill folders, drafts, settings, and library paths
never leave your disk. You can verify this by reading
the
source or by watching network traffic with your tool of choice.
Ready to make your AI workflows reusable?
Download Skillful, point it at your library, and stop maintaining the same workflow in five places.
