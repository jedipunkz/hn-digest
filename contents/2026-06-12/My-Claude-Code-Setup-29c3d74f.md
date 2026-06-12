---
source: "https://illuminatedcomputing.com/posts/2026/06/my-claude-code-setup/"
hn_url: "https://news.ycombinator.com/item?id=48505773"
title: "My Claude Code Setup"
article_title: "Illuminated Computing\n|\nMy Claude Code Setup"
author: "pjungwir"
captured_at: "2026-06-12T16:06:16Z"
capture_tool: "hn-digest"
hn_id: 48505773
score: 1
comments: 0
posted_at: "2026-06-12T15:57:53Z"
tags:
  - hacker-news
  - translated
---

# My Claude Code Setup

- HN: [48505773](https://news.ycombinator.com/item?id=48505773)
- Source: [illuminatedcomputing.com](https://illuminatedcomputing.com/posts/2026/06/my-claude-code-setup/)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T15:57:53Z

## Translation

タイトル: 私のクロード コードのセットアップ
記事のタイトル: イルミネーション コンピューティング
|
私のクロードコードのセットアップ

記事本文:
コードのような曲
私はコンピューティングに光を当てました
クロード コードの実行方法は特殊なようですが、利便性とセキュリティの間で適切なトレードオフが得られると考えています。
HN で説明していましたが、ここにも書いておきたいと思います。
私は常に --dangerously-skip-permissions (またはそれが何と呼ばれているか、ずっと前に設定したグローバル フラグがあります) を使用して実行します。誰もがそうだと思います。そうしないと面倒すぎます。しかし、許容できるリスクでそれを行うにはどうすればよいでしょうか?
実際、どのモードでも、制限なしでクロードを実行することは望ましくありません。私のアカウントには、AWS 認証情報、k8s 認証情報、SSH キー、Github アクセス、何十もの顧客からの .env ファイルなど、誰が知っているかが含まれています。最近パスワードを入力した場合は、sudo を実行して問題なくパスできます。 ~/bin には、VPN に参加してデータベースにログインするためのスクリプトがあります。私のブラウザはどこでもサインインしています。 Thunderbird ではメールの送受信が可能です。そして、たとえクロードが境界線を決して忘れなかったとしても、私は少なくとも .env ファイルを Anthropic に送っているのではないでしょうか?
そこで私の解決策は、Claude に独自の OS ユーザーを与えることでした。
LLM は別の同僚のようなものだと人々は言うので、私は LLM をそのように扱っています。彼は私と同様のドットファイルを持っていますが、秘密はありません。私自身のホーム ディレクトリは 0700 です。彼は私が github プロファイルに追加した独自の ssh キーを持っていますが、パスワードで保護されているので、私が彼のためにプッシュ/プルしています。彼は独自の Postgres (非スーパーユーザー!) {開発、テスト} {ユーザー、データベース} を持っています。 sudo を使って何かを実行する必要があるかどうか、彼は私に尋ねます。多くの場合、私たちは二人で並行して何かに取り組むことができます。結局のところ、Unix はマルチユーザー システムであるはずでした。私は Debian 13 と xfce を使用していますが、これは他の場所でもうまく機能すると思います。
クロードに何かをしてもらいたいときは、別の端末タブを開いて彼のアカウントに su します。彼も私と同じように、プロジェクト用の ~/src フォルダーを持っています。そのうちの 1 つに移動し、tmux セッションを開始します。彼の ~/.tmux.conf は

すべてのセッションには黄色のステータス バーが表示されるので、簡単に認識できます。新しいシェルは Ctrl-a c だけです。
通常、私は最初の tmux ウィンドウを bash に保持し、プッシュ/プル、コミット/差分の読み取り、追加のテストの実行、その他何でもできるようにします。 2 番目のウィンドウでクロード セッションを実行します。さらに何かをする必要がある場合は、さらにウィンドウを起動します。多くの場合、ウィンドウ 3 は psql です。 vim などについてはさらに詳しく説明します。
私がよく使うトリックは、彼の git リポジトリの多くには次のような追加のリモートが含まれていることです。
ポール ssh://paul@localhost/~/src/example (フェッチ)
ポール ssh://paul@localhost/~/src/example (プッシュ)
これにより、共有する準備ができていないものについても簡単に共同作業できるようになります。また、リポジトリ以外のものを ~paul の外側に配置できる /pub/paul フォルダーも設定しましたが、これはほとんど使用していません。
VM をいじらないのが気に入っています。すべてはホスト上にあります。すべてが一度セットアップされます。クロードの環境も私の環境と同じくらい快適です。また、クロードの割り当ての一部では VM を実行する必要があるため、追加のネストはありません。
私はこのセットアップを数か月間使用してきましたが、とても気に入っています。
Linux の権限昇格のバグが心配です。 AI が脆弱性の悪用は容認できないことを理解するとは思えません。 (最初の仕事で、急いで何かが必要になったときに、公式には httpd.conf の編集に限定されていた sudo 権限を拡張するために vim の :! 機能を誤用した可能性があることを思い出さずにはいられません。...) 最近では、自動セキュリティ更新にもかかわらず、パッケージを手動でアップグレードすることが増えていることに気づきました。 Opus がセキュリティの脆弱性をわざわざ調べるとは思えませんが、Fable ならそうするかもしれません。最近は脆弱性がたくさんあります。おそらく、将来のモデルは自ら新しいモデルを見つけることになるでしょう。または、キーロガーをインストールして SSH キーのパスワードを学習します。
VM のほうが安全なのでしょうか?そこには直感がありません。しかし、そこに

ハイパーバイザーエスケープの脆弱性もあり、共有フォルダーが心配です。たとえば、vagrant では、ゲストは /vagrant を取得してホストフォルダーを読み書きします。何をどこに置くかは非常に注意する必要があります。
これまでのところ最大の煩わしさは、Docker コンテナーの実行です。クロードを docker グループに追加したり、sudo 権限を付与したりしたくありません。ユーザー用にルートレス Docker を設定でき、通常のシステム全体の Docker と並行して実行できると読んだことがありますが、まだ試していません。それがうまくいかない場合は、おそらくクロードに自分のマシンを与えるでしょう。予備のボックスやラップトップがたくさん転がっています。
どう思いますか？私のアプローチにはセキュリティ上の問題がありますか?効率的でありながら責任を負うための素晴らしい方法だと思います。
ブログのコメントは Disqus によって提供されています
次へ: Linux (Debian 13) 上の GeForce RTX 5070 Ti
ポール・A・ユングヴィルス
ブログ
Linux (Debian 13) 上の GeForce RTX 5070 Ti
Postgres での Bison シフトの解決/競合の削減
時間スキーマへの移行
私たちが愛する論文: 時間的調整
Postgres 一時外部キーのベンチベース
Postgres のハッキング: 実行フェーズ
SQL:2011 有効時間を Postgres に追加する進捗状況
テンポラル データベース: 理論と Postgres
© Illuminated Computing Inc. 2009 – 2026。全著作権所有。

## Original Extract

Code like song
I lluminated Computing
I seem to have an unusual way of running Claude Code, but I find it a good trade-off between convenience and security.
I was explaining it on HN , but I wanted to write it down here too:
I always run with --dangerously-skip-permissions (or whatever it’s called; there’s a global flag I set a long time ago). I assume everyone does. It’s too tedious otherwise. But how can I do that with tolerable risk?
In fact in any mode, I wouldn’t want to run Claude without restrictions. My account has AWS creds, k8s creds, ssh keys, Github access, .env files from dozens of customers with who-knows-what, etc. I can run sudo and pass unchallenged, if I recently gave the password. There are scripts in ~/bin to join VPNs and log in to databases. My browser is signed in everywhere. Thunderbird can send/receive email. And even if Claude never forgets a boundary, aren’t I sending Anthropic at least the .env files?
So my solution was to give Claude its own OS user.
People say the LLM is like another co-worker, so I’m treating it that way. He has similar dotfiles to mine, but no secrets. My own home directory is 0700. He has his own ssh key that I added to my github profile, but it’s password-protected, and I push/pull for him. He has his own Postgres (non-superuser!) {development,test} {users,databases}. If he needs something run with sudo, he asks me. Often we can both work on something in parallel. Unix was supposed to be a multi-user system after all. I’m on Debian 13 with xfce, but I think this would work well elsewhere.
When I want Claude to do something, I open another terminal tab and su to his account. He has a ~/src folder for projects, just like me. I go to one of those and start a tmux session. His ~/.tmux.conf gives every session a yellow status bar, so they’re easy to recognize. Then new shells are just Ctrl-a c .
Usually I keep the first tmux window in bash, so I can push/pull, read commits/diffs, run extra tests, and do whatever else. I run the claude session in the second window. If I need to do more things, I start more windows. Frequently window 3 is psql. More for vim, etc.
A trick I use a lot is that many of his git repos have an extra remote, like this:
paul ssh://paul@localhost/~/src/example (fetch)
paul ssh://paul@localhost/~/src/example (push)
That makes it easy to collaborate on things I’m not ready to share. I also set up a /pub/paul folder where I can put non-repo things outside of ~paul , but I’ve hardly used it.
I like that I’m not mucking with VMs. Everything is on the host. Everything is set up once. Claude’s environment is as comfortable as mine. And some of Claude’s assignments require him running VMs, so there is no extra nesting.
I’ve been using this setup for months, and I really like it.
I do worry about Linux privilege escalation bugs. I don’t trust an AI to understand that exploiting vulns is not acceptable. (I can’t help but recall that at my first job I may have misused vim’s :! feature to broaden my sudo powers, which were officially limited to editing httpd.conf , when I needed something in a hurry. . . .) I find myself manually upgrading packages more often these days, despite automatic security updates. I don’t think Opus would go to the trouble of looking up security vulns, but maybe Fable would, and there have been a lot lately. Maybe some future model will just take it upon itself to find new ones. Or install a keylogger to learn the ssh key password.
Would a VM be more secure? I don’t have an intuition there. But there are hypervisor escape vulns too, and I’m anxious about shared folders. For instance in vagrant the guest gets /vagrant to read/write the host folder. You’d have to be very careful what you put where.
The biggest annoyance so far is running docker containers. I don’t want to add claude to the docker group or give it sudo privileges. I’ve read that you can set up rootless docker for a user, and even that you can run it side-by-side with a normal system-wide docker, but I haven’t tried doing that yet. If that doesn’t work, I will probably give Claude his own machine. I have plenty of spare boxes/laptops lying around.
What do you think? Are there security problems with my approach? I think it’s a great way to be efficient but responsible.
blog comments powered by Disqus
Next: GeForce RTX 5070 Ti on Linux (Debian 13)
Paul A. Jungwirth
Blog
GeForce RTX 5070 Ti on Linux (Debian 13)
Solving bison shift/reduce conflicts in Postgres
Migrating to a Temporal Schema
Papers We Love: Temporal Alignment
Benchbase for Postgres Temporal Foreign Keys
Hacking Postgres: the Exec Phase
Progress Adding SQL:2011 Valid Time to Postgres
Temporal Databases: Theory and Postgres
© Illuminated Computing Inc. 2009 – 2026. All rights reserved.
