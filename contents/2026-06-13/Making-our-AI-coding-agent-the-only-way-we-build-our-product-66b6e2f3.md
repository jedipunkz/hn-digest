---
source: "https://anyframe.dev/blog/we-hired-an-intern-named-gilfoyle"
hn_url: "https://news.ycombinator.com/item?id=48517164"
title: "Making our AI coding agent the only way we build our product"
article_title: "We Hired an Intern Named Gilfoyle · AnyFrame Blog · AnyFrame"
author: "nurdtechie98"
captured_at: "2026-06-13T15:38:58Z"
capture_tool: "hn-digest"
hn_id: 48517164
score: 4
comments: 0
posted_at: "2026-06-13T13:30:29Z"
tags:
  - hacker-news
  - translated
---

# Making our AI coding agent the only way we build our product

- HN: [48517164](https://news.ycombinator.com/item?id=48517164)
- Source: [anyframe.dev](https://anyframe.dev/blog/we-hired-an-intern-named-gilfoyle)
- Score: 4
- Comments: 0
- Posted: 2026-06-13T13:30:29Z

## Translation

タイトル: 製品を構築する唯一の方法として AI コーディング エージェントを作成する
記事のタイトル: ギルフォイルという名前のインターンを採用しました · AnyFrame ブログ · AnyFrame
説明: 彼はサンドボックスに住んでいて、眠らずに、私たちの本番リポジトリにコミットしています。 AnyFrame 上に内部開発エージェントを構築し、それを AnyFrame に向けた方法。

記事本文:
ギルフォイルという名前のインターンを雇いました · AnyFrame ブログ · AnyFrame ← すべての投稿 2026 年 6 月 7 日
ギルフォイルという名前のインターンを採用しました
彼はサンドボックスに住んでいて、眠らずに、私たちの本番リポジトリにコミットしています。 AnyFrame 上に内部開発エージェントを構築し、それを AnyFrame に向けた方法。
私たちの新しいチームメンバーはギルフォイルという名前です。彼はインターンで、今週は
チームのほとんどの人間よりも多くのコードを本番環境に送りました。彼の一人
プル リクエストにより、機能全体、データベース移行などが追加されました。それを統合しました
他の誰もラインに触れることなく。あなたが読んでいるページも彼が作ったものです。そして
彼はすべてを自分自身のコミット アイデンティティの下で実行しました。
「ギルフォイル」は人間ではありません。彼は社内の開発エージェントであり、クラウド上の AI です。
コードベースがチェックアウトされ、シェルが開いたサンドボックス。私たちは彼に次の名前を付けました
シリコンバレーのキャラクター、テレビでは最も傲慢で無表情なシステムエンジニア、
インターンに降格されましたが、どういうわけかまだ当社で最も生産性の高い社員がいます。
結論: 私たちは、サンドボックス化された AI エージェント用のコントロール プレーンである AnyFrame を構築しました。
次に、AnyFrame を AnyFrame に向けました。ギルフォイルは自分が出荷する製品を実行します。
エージェントを定義します: リポジトリ、インストール コマンド、システム プロンプト、スキル
(再利用可能な Playbook)、Slack、Linear、GitHub などのツールへの接続。
次に、実際に作業を実行できる分離されたクラウド サンドボックスで起動します。チャットボットではありません
ドラフトの提案: コードを含む実際のマシンと、開くターミナル
PR、テストの実行、出荷。クロード・コードやコーデックスが存続していればよかったのにと思うなら
チームが @メンションしてプライベート リポジトリを信頼できるサーバー、それが Gilfoyle です。
私たちは、他の顧客と同じように、3 つのステップで彼をオンボーディングしました。
1. テンプレートを作成します。テンプレートはブループリントです: リポジトリ (モノリポジトリ)、
インストールコマンド、スキル、そして私たちの修道院を運ぶシステムプロンプト

イオン
そして彼の人柄。そこから紡ぎ出されたものはすべてそのロットを継承します。
テンプレート: ブループリントを一度定義すると、すべてのエージェントで再利用できます。
2. エージェントを作成します。そのテンプレートからエージェントをスピンし、クロードを選択します。
ランタイム (実際に彼を動かすモデル) を作成し、GitHub に接続してクローンを作成できるようにします。
そして彼自身のコミットアイデンティティの下にプッシュします。
Gilfoyle の構成: ランタイム、リポジトリ、スキル、トリガー、権限が 1 か所にまとめられています。
3. Discord を彼に向けます。 AnyFrame の Discord 統合をオンにして狙いを定めます
エージェントで。グルーコードもベビーシッターサービスもありません。
全体の配線: Discord サーバーがギルフォイル エージェントに向けられています。
現在のワークフローは、どのチャンネルでも彼を @メンションするだけです。彼は糸を紡ぎ出し、
クローン化されたリポジトリを使用して新しいサンドボックスを起動し、作業をストリームバックします。フォローアップ
メッセージは同じサンドボックスに送られます。ターンの間に追い出された場合、次のメッセージ
スナップショットからサイレントに再開します。 1 つのスレッド = 1 つのサンドボックス = 1 つの作業単位、
したがって、3 人が彼に一度に 3 つのタスクを担当させることができます。彼はクローンを作るインターンです
彼自身。
典型的なブリーフ: @彼に、変化を平易な英語で説明し、一日に戻ってください。
1 回のチェックアウトで会社全体の情報を彼に提供
彼が何に取り組んでいるのか、そしてどのように取り組んでいるのかという 2 つの意図的な選択が私たちに影響力をもたらしました。
それを証明しています。
AnyFrame は 1 つのリポジトリではありません。バックエンド、ダッシュボード、SDK、およびいくつかのリポジトリです。
統合は共に進化します。エージェントを 1 つに向けても、エージェントは何も知りません。
残りはすべてを 1 つの monorepo にまとめ、チェックアウトは 1 つです。
すべてのリポジトリは依然として独自のリモートにプッシュされます。 API にまたがるタスク、その
SDK は 1 つの一貫した作品であり、3 つのコンテキスト スイッチではありません。
作れない。
すべてのリポジトリはサブモジュールとして、1 つの場所から作業できます。
次にスキルです。私たちの信頼を得たのは、Proof of Workです
スキル: すべてのサンドボックスには Chromium と Playwright が付属しています。

目覚めたので、いつ
Gilfoyle は目に見えるものを変更し、開発サーバーを起動し、実際のマシンを駆動します。
ブラウザでページにアクセスし、スクリーンショットを作品にホチキスで留めます。
彼はバナーが修正されたとは言っていない。彼はそれを電話の幅で固定して示しています
ビューポート。画像だけでは十分ではない場合、AnyFrame はサンドボックスをトンネリングして、
パブリック インターネットに接続すると、回線が切断される前に、実行中のブランチにライブ URL をドロップします。
合併した。
この投稿の変更をスクリーンショットで証明する実際の PR コメント。
これにより、レビュー ループが崩壊します。終了したと発表するインターンは、
責任。スクリーンショットとリンクを渡して試す人は同僚です。
彼は、1 行の修正から機能全体に至るまで、コードを出荷します。スモールエンド:「
ファーストサンドボックスのバナーがモバイルで 3 行に折り返される、修正してください。」(コミット
b1a4382 )、またはドキュメント内の間違った TTL を修正します。反対側では PR #128、
定期的なエージェント セッションのスケジューリングは、自分自身の下で自律的に開かれました
アイデンティティとマージ: データベース テーブル、バックグラウンド スケジューラ、移行、
以下のUI。他に一行も書いた人はいなかった。
彼が導入した機能: すべてのエージェントが cron スケジュールを実行できるようになりました。私たちが最初に設定したものは、平日の毎朝 9 時に起きて、デッドコードを探し出し、ドラフト PR を開きますが、人間は関与しません。
彼は、コンテキストからも次のように答えています。「アイドル リーパーのタイムアウトをどこに設定すればよいでしょうか?」
バイブではなく、ファイルパスとコミットハッシュが返されます。
これは「きちんとしたボットを開発しました」よりも野心的なものです。ギルフォイルを
AnyFrame を開発する唯一の方法です。すべての機能、すべての修正は、
製品に住むインターン。そのため、私たちは最も要求の厳しいユーザーになります。
自分の製品を継続的に: 開発者が必要とする何かを実行できなくなった瞬間、
それは私たち自身の血で書かれたバグレポートであり、顧客が到着したのと同じ週に報告されます。
それを感じるだろう。したがって、ループは自己強化され、AnyFrame を構築するためにループを使用します。
サーフィン

彼に欠けているものを補うために、私たちは彼にそれも作ってもらいます。機能はこれから生まれます
ホワイトボードではなくループです。 (上記のスケジュールは、彼が発送するまでこのリストに載っていました)
) 次は次のとおりです。
メッセージだけでなくチケットも受け取りましょう。リニアに接続して問題をピックアップする
そして自分自身でカードを動かします。バックログが彼のキューになります。
彼のチェックアウトを常に最新の状態に保ちます。スケジュールに従ってプルとリベースを行うことで、
樹齢1週間の木の上に建物を建てています。
プロセスがそのままの状態で一時停止を生き延びます。サンドボックスを再開すると、
ディスクはありましたが、彼が実行していたものはすべて消えてしまいました。今日、彼は手動でそれを再起動します。
タスク全体にわたって覚えておいてください。永続的なメモリにより、各タスクがよりスマートに開始されます。
毎朝リセットするのではなく、毎週改善するインターン。
これらはどれも魅力的ではありません、そしてそれが重要です。相談できるエージェント
あなたのコードはパーラートリックです。本番環境への唯一の道は、
100 個の退屈なバー、アイデンティティ、認証、隔離、スナップショットをクリアするために、
許可、信頼。その制約の中で生きることで、私たちはそれらを見つけることができます。ギルフォイル
私たちのマスコットではありません。彼は私たちの強制的な役割です。
それで、あなたのギルフォイルは誰になりますか、そしてあなたは彼に何を指摘しますか？
あなた自身のインターンを雇用したいですか? anyframe.dev で、
テンプレートを作成し、エージェントをスピンオフして、Discord 統合をオンにします。あなたはそうするだろう
コーヒーが冷める前に、サンドボックス化された開発エージェントをチームに参加させてチャットしましょう。ネーミング
彼ギルフォイルはオプションですが、奨励されています。
1 回のチェックアウトで会社全体の情報を彼に提供

## Original Extract

He lives in a sandbox, doesn't sleep, and commits to our production repo. How we built an internal dev agent on AnyFrame, and pointed it at AnyFrame.

We Hired an Intern Named Gilfoyle · AnyFrame Blog · AnyFrame ← all posts June 7, 2026
We Hired an Intern Named Gilfoyle
He lives in a sandbox, doesn't sleep, and commits to our production repo. How we built an internal dev agent on AnyFrame, and pointed it at AnyFrame.
Our newest team member is named Gilfoyle. He's an intern, and this week he
shipped more code to production than most of the humans on the team. One of his
pull requests added an entire feature, database migration and all. We merged it
without anyone else touching a line. He built the page you're reading, too. And
he did all of it under his own commit identity.
"Gilfoyle" isn't a person: he's our internal developer agent, an AI in a cloud
sandbox with our codebase checked out and a shell open. We named him after the
Silicon Valley character, the most arrogant, deadpan systems engineer on TV,
demoted to intern and somehow still our most productive headcount.
The punchline: we built AnyFrame, a control plane for sandboxed AI agents,
then pointed AnyFrame at AnyFrame. Gilfoyle runs on the product he ships.
You define an agent : a repo, an install command, a system prompt, skills
(reusable playbooks), and connections to tools like Slack, Linear, or GitHub.
Then you boot it into an isolated cloud sandbox that can actually do the work. Not a chatbot that
drafts suggestions: a real machine with your code and a terminal, that opens
PRs, runs your tests, and ships. If you've wished Claude Code or Codex lived on
a server your team could @mention and trust with private repos, that's Gilfoyle.
We onboarded him the way any customer would, in three steps.
1. Create a template. A template is the blueprint: the repo (our monorepo),
the install command, the skills, and a system prompt carrying our conventions
and his personality. Anything spun up from it inherits the lot.
The template: define the blueprint once, reuse it for every agent.
2. Create the agent. Spin an agent off that template, pick the Claude
runtime (the model that actually drives him), and connect GitHub so he can clone
and push under his own commit identity.
Gilfoyle's config: runtime, repo, skills, triggers, and permissions in one place.
3. Point Discord at him. Switch on AnyFrame's Discord integration and aim it
at the agent. No glue code, no service to babysit.
The whole wiring: a Discord server, pointed at the Gilfoyle agent.
Now the workflow is just @mention him in any channel. He spins up a thread,
boots a fresh sandbox with the repo cloned, and streams the work back. Follow-up
messages go to the same sandbox; if it's evicted between turns, the next message
silently resumes from a snapshot. One thread = one sandbox = one unit of work,
so three people can put him on three tasks at once. He's an intern who clones
himself.
A typical brief: @ him, describe the change in plain English, get back to your day.
Giving him the whole company in one checkout
Two deliberate choices gave us the leverage: what he works on, and how he
proves it.
AnyFrame isn't one repo: it's a backend, a dashboard, an SDK, and a few
integrations evolving together. Point an agent at one and it's blind to the
rest, so we stitched everything into a single monorepo , one checkout with
every repo, each still pushing to its own remote. A task spanning the API, its
client, and the SDK is one coherent piece of work, not three context switches he
can't make.
Every repo as a submodule, one place to work from.
Then there are skills . The one that earned our trust is the proof-of-work
skill: every sandbox ships with Chromium and Playwright baked in, so when
Gilfoyle changes something you can see, he boots the dev server, drives a real
browser to the page, and staples a screenshot to the work.
He doesn't say the banner's fixed; he shows it fixed, on a phone-width
viewport. When a picture isn't enough, AnyFrame tunnels his sandbox to the
public internet and he drops a live URL to his running branch, before a line has
merged.
A real PR comment, proving a change with a screenshot, of this very post.
That collapses the review loop: an intern who announces he's done is a
liability; one who hands you a screenshot and a link to try is a colleague.
He ships code, from one-line fixes to whole features. Small end: "the
first-sandbox banner wraps to three lines on mobile, fix it" (commit
b1a4382 ), or correcting a wrong TTL in the docs. At the other end, PR #128,
recurring agent session scheduling , was opened autonomously under his own
identity and merged: a database table, a background scheduler, a migration, and
the UI below. Nobody else wrote a line.
A feature he shipped: any agent can now carry a cron schedule. The first one we set up wakes every weekday at 9am, hunts dead code, and opens a draft PR, no human in the loop.
He answers from context , too: "where do we set the idle-reaper timeout?"
comes back with a file path and a commit hash, not vibes.
This is more ambitious than "we have a neat bot": we want Gilfoyle to be the
only way we develop AnyFrame. Every feature, every fix, routed through the
intern who lives in the product. That makes us the most demanding user of our
own product, continuously: the instant he can't do something a developer needs,
it's a bug report written in our own blood, landing the same week a customer
would feel it. So the loop self-reinforces, we use him to build AnyFrame, that
surfaces what he's missing, we have him build that too. Features come from this
loop, not a whiteboard. (Scheduling, above, was on this list until he shipped
it.) Next up:
Take tickets, not just messages. Connect Linear so he picks up issues
and moves cards himself; the backlog becomes his queue.
Keep his checkout fresh. Pull and rebase on a schedule so he's never
building on a week-old tree.
Survive a pause with processes intact. A resumed sandbox restores the
disk, but anything he had running is gone; today he relaunches it by hand.
Remember across tasks. Persistent memory so each task starts smarter,
an intern who improves every week instead of resetting every morning.
None of this is glamorous, and that's the point. An agent that can talk about
your code is a parlor trick; one you'd make your only path to production has
to clear a hundred boring bars, identity, auth, isolation, snapshots,
permissions, trust. Living inside that constraint is how we find them. Gilfoyle
isn't our mascot. He's our forcing function.
So: who would your Gilfoyle be, and what would you point him at?
Want to hire your own intern? At anyframe.dev , create a
template, spin an agent off it, and switch on the Discord integration. You'll
have a sandboxed dev agent in your team chat before your coffee's cold. Naming
him Gilfoyle is optional but encouraged.
Giving him the whole company in one checkout
