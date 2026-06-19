---
source: "https://sauleau.com/notes/airgap-security-for-the-modern-ai-age.html"
hn_url: "https://news.ycombinator.com/item?id=48602862"
title: "Hide Secrets from AI Agents and NPM install using Airgap"
article_title: "Sven Sauleau: Systems Engineer - airgap - Security for the modern AI age"
author: "netgusto"
captured_at: "2026-06-19T21:29:30Z"
capture_tool: "hn-digest"
hn_id: 48602862
score: 2
comments: 0
posted_at: "2026-06-19T20:24:59Z"
tags:
  - hacker-news
  - translated
---

# Hide Secrets from AI Agents and NPM install using Airgap

- HN: [48602862](https://news.ycombinator.com/item?id=48602862)
- Source: [sauleau.com](https://sauleau.com/notes/airgap-security-for-the-modern-ai-age.html)
- Score: 2
- Comments: 0
- Posted: 2026-06-19T20:24:59Z

## Translation

タイトル: エアギャップを使用した AI エージェントと NPM インストールからのシークレットの非表示
記事のタイトル: Sven Sauleau: システム エンジニア - エアギャップ - 現代の AI 時代のセキュリティ
説明: airgap は、マウント名前空間でプログラムを実行し、ファイルのシークレットを編集して、悪意のある npm インストール フックや奇妙な AI エージェントから保護する透過的なラッパーです。

記事本文:
エアギャップ - 現代の AI 時代のセキュリティ
AI エージェントにプロジェクト内のファイルの読み書きをさせ、インターネットからスキルやプラグインをインストールします。 npm マルウェアはインストール時に秘密を盗みます。 AI エージェントはコマンドを実行し、npm パッケージもインストールしますが、場合によってはそれらのパッケージが幻覚を示したり、悪意のあるものになったりすることがあります。
airgap は、名前空間内でプログラムを実行するラッパーです。ファイル内の機密情報は AI エージェントから隠されます。エージェントは引き続きファイルを操作できますが、実際の値は表示されません。パッケージ マネージャーはゲートされており、アクセスしてはいけないファイルにアクセスしようとする前に質問されます。
Linux のみ。 macOS のサポートは進行中です。
シークレットは、 .env 、 ~/.ssh 、 ~/.npmrc などのプレーン ファイルに保存されます。最近、私たち、または AI エージェントは、それらを読み取る可能性のある信頼できないコードを多数実行しています。それは 2 つの方法で起こります:
AIエージェント。 clude や opencode のようなエージェントはプロジェクトとホーム ディレクトリを読み取り、読み取ったものはすべてモデル プロバイダーに送信される可能性があります。 .env 値または SSH キーは、気付かないうちにプロンプ​​トに表示されることがあります。悪意のあるプラグインまたはスキルは、機密ファイルを意図的に取得する可能性があります。
npmパッケージ。悪意のある npm パッケージは新しいものではありませんが、最近の波は非常に成功しており、急速に拡散しています。依存関係により、インストール時に任意のコードを実行するプレインストール フックまたはポストインストール フックを追加できます。そのコードは、.env ファイル、SSH キー、クラウド認証情報の後に続きます。 Shai-Hulud 、 Miasma 、 pgserve 、および偽の Tanstack パッケージなどのキャンペーンはすべて同じことを行います。つまり、インストール時に実行され、秘密を取得し、送信されます。また、永続的な悪い状態を残す場合もあります。そして、エージェントは、自分が作成したパッケージに対して npm install を実行します。これは、攻撃者が事前に登録する名前です (スロップスクワッティング)。
最も明確な例は、Shai-Hulud です。これは、

npm は 2025 年後半にリリースされます。侵害されたパッケージをインストールすると、そのフックは .npmrc (npm トークン)、環境変数、GitHub PAT と AWS、GCP、Azure のクラウド キーを保持する構成ファイル内の認証情報を検索します。それらを攻撃者に送信します。
盗まれたnpmトークンを使用して、開発者が管理する他のパッケージのバックドアバージョンを再公開するため、それらがインストールされるたびに再び拡散します。 2025 年後半の 2.0 キャンペーンでは、コードが preinstall に移動されました。これは CI/CD パイプラインにも影響します。
エアギャップがあればどうやって阻止しただろう
これらのマルウェアは、ディスクからシークレットを盗むことができる場合にのみ機能します。エアギャップはそれを防ぐためのものです。 .env 、秘密キー、または ~/.npmrc を読み取ると、編集されたコンテンツが返されます。 ~/.aws/credentials などの他の予期しないファイルは、最初に許可を求めます。
ファイル内のシークレット ( .env など)
プロジェクトとホーム ディレクトリ内。どんどん追加していきます。
airgap は、新しいマウントおよびユーザー名前空間にターゲット プログラムを生成し、ホーム ディレクトリと現在の作業ディレクトリ (ホーム ディレクトリと異なる場合) を FUSE ファイルシステムとしてマウントします。すべてのファイルシステムへのアクセスは、エアギャップのハンドラーを経由します。そこから、シークレットを編集するか、プログラムがそのファイルにアクセスすることになっていたかどうかを確認するように求めます。
エアギャップ <プログラム> [ 引数...]
エアギャップには、プログラムごとに特別なルールを持つプログラム許可リストがあります。現在サポートされているもの:
AI エージェント — claude 、opencode — は編集のみで実行されます。
パッケージ マネージャー (npm) は、編集と対話型のファイル ゲートを使用して実行されます。 ~/.gitconfig 、node_modules 、lockfile などの安全なパスは事前に承認されています。
さらに多くのツールが追加されます。別のプログラムのサポートが必要な場合は、問題をオープンしてください。
$猫.env
API_KEY = sk-live-9f8c2a1b4e7d
DB_PASSWORD = ハンター2
$ エアギャップ猫 .env
API_KEY = <編集された値>
DB_PASSWORD = <編集された値>
フィ

ファイルは AI エージェントにとって読み取りおよび編集可能なままですが、秘密だけが AI エージェントから隠されます。
airgap の下で npm を実行すると、インストール フックが実行されますが、新しいファイルが読み取られるたびにプロンプトがトリガーされます。
$ airgap npm install bad-package
エアギャップ: /home/sven/.ssh/id_rsa の読み取りを許可しますか? [y/N]n
エアギャップ: ブロックされた /home/sven/.ssh/id_rsa
2秒で42個のパッケージを追加しました
~/.ssh/id_rsa を読み取ろうとするポストインストール スクリプトは捕捉され、拒否されます。インストールは続行されますが、フックは何も取得しません。
package.json 、 lockfiles 、npm キャッシュなどの通常の読み取りは事前に承認されているため、予期しないファイルについてのみプロンプトが表示されます。
$エアギャップ・クロード
> ./.env の内容を表示
1ファイルを読み込む
ファイルは次のとおりです。
DATABASE_URL="<編集された値>"
API_KEY="<編集された値>"
AWS_SECRET_ACCESS_KEY="<編集された値>"
エージェントは実際のファイルを操作しますが、編集されたシークレットのみを参照します。
エイリアスを使って透明にする
エージェントまたはパッケージ マネージャーを常に airgap の下で実行するには、シェル設定にエイリアスを追加します ( ~/.bashrc 、 ~/.zshrc 、…)。
別名クロード = "エアギャップ クロード"
エイリアス opencode = "エアギャップ opencode"
エイリアス npm = "エアギャップ npm"
これで npm install が実行され、何も考えずにエージェントがエアギャップ内で実行されます。
エアギャップは 1 層であり、保証するものではありません。ミスもするし、攻撃も変化し続ける。貢献は大歓迎です!ホール、サポートすべきプログラム、または編集すべき秘密を見つけた場合は、問題を提起し、改善に協力してください。
セキュリティ関連の場合は、 [email protected] まで電子メールでお問い合わせください。
GitHub: https://github.com/xtuc/airgap
crates.io: https://crates.io/crates/airgap
こんにちは: [email protected] 。
Twitter で私に連絡してください: @svensaulau 。

## Original Extract

airgap is a transparent wrapper that runs programs in a mount namespace and redacts secrets from files, protecting against malicious npm install hooks and curious AI agents.

airgap - Security for the modern AI age
We let AI agents read and write files in our projects, and we install skills or plugins from the internet. npm malware steals secrets at install time. The AI agents run commands and install npm packages too, and sometimes those packages are hallucinated or malicious.
airgap is a wrapper that runs a program inside namespaces. Sensitive secrets in files are hidden from AI agents. Agents can still work with the files, but they never see the real values. Package managers are gated: you get asked before they try to access a file they’re not supposed to.
Linux only. macOS support is work in progress.
Your secrets sit in plain files like .env , ~/.ssh , and ~/.npmrc . These days we, or our AI agents, run a lot of untrusted code that might read them. Two ways that happens:
AI agents. An agent like claude or opencode reads through your project and home directory, and everything it reads might be sent to a model provider. A .env value or an SSH key can land in a prompt without you noticing. A malicious plugin or skill can grab sensitive files on purpose.
npm packages. Malicious npm packages aren’t new, but the recent wave is very successful and spreads fast. A dependency can add a preinstall or postinstall hook that runs arbitrary code when you install it, and that code goes after your .env files, SSH keys, and cloud credentials. Campaigns like Shai-Hulud , Miasma , pgserve , and the fake tanstack package all do the same thing: run on install, grab secrets, send them out. Some also leave persistent bad state behind. And an agent will run npm install on a package it made up, which is a name attackers now register ahead of time (slopsquatting).
The clearest example is Shai-Hulud , a self-replicating worm that spread through npm in late 2025. When you install a compromised package, its hook looks for credentials in .npmrc (npm tokens), environment variables, and config files holding GitHub PATs and cloud keys for AWS, GCP, and Azure. It sends them to the attacker.
With a stolen npm token, it republishes backdoored versions of the other packages that developer maintains, so every install of those spreads it again. The 2.0 campaign in late 2025 moved the code to preinstall , which also affects CI/CD pipelines.
How airgap would have stopped it
These malware only work if they can steal your secrets from disk. airgap is meant to prevent that. When it reads .env , a private key, or ~/.npmrc , it gets back redacted content. Any other unexpected file, like ~/.aws/credentials , asks for your permission first.
secrets in files ( .env , etc.)
In your project and home directories. We’re adding more and more.
airgap spawns the target program in a new mount and user namespace, and mounts your home directory and the current working directory (if different from the home one) as a FUSE filesystem. Every filesystem access goes through airgap’s handler. From there we redact secrets, or ask you to confirm the program was supposed to access that file.
airgap <program> [ args...]
airgap has a program allowlist, with special rules for each program. Currently supported:
AI agents — claude , opencode — run with redaction only.
Package manager — npm — runs with redaction and an interactive file gate. Benign paths like ~/.gitconfig , node_modules , and lockfiles are pre-approved.
More tools will be added. If you want support for another program, open an issue .
$ cat .env
API_KEY = sk-live-9f8c2a1b4e7d
DB_PASSWORD = hunter2
$ airgap cat .env
API_KEY = <redacted value>
DB_PASSWORD = <redacted value>
The file stays readable and editable for AI agents, just the secrets are hidden from them.
When running npm under airgap , the install hooks run but each new file read triggers a prompt:
$ airgap npm install bad-package
airgap: allow reading /home/sven/.ssh/id_rsa? [ y/N] n
airgap: blocked /home/sven/.ssh/id_rsa
added 42 packages in 2s
A postinstall script that tries to read ~/.ssh/id_rsa gets caught and denied. The install keeps going, and the hook gets nothing.
Normal reads like package.json , lockfiles, and the npm cache are pre-approved, so you only get prompted for files that aren’t expected.
$ airgap claude
> show me ./.env contents
Read 1 file
Here 's the file:
DATABASE_URL="<redacted value>"
API_KEY="<redacted value>"
AWS_SECRET_ACCESS_KEY="<redacted value>"
The agent works on the real files but only sees redacted secrets.
Make it transparent with aliases
To always run your agent or package manager under airgap , add aliases to your shell config ( ~/.bashrc , ~/.zshrc , …):
alias claude = "airgap claude"
alias opencode = "airgap opencode"
alias npm = "airgap npm"
Now npm install and your agents run inside airgap without you thinking about it.
airgap is one layer, not a guarantee. It will miss things, and the attacks keep changing. Contributions welcome! If you find a hole, a program it should support, or a secret it should redact, open an issue and help make it better.
For anything security related, email [email protected] .
GitHub: https://github.com/xtuc/airgap
crates.io: https://crates.io/crates/airgap
Say hello: [email protected] .
Ping me on Twitter: @svensauleau .
