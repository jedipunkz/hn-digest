---
source: "https://grigio.org/sandbox-bwrap-nix-a-lightweight-sandbox-for-ai-agents-and-experiments/"
hn_url: "https://news.ycombinator.com/item?id=49046290"
title: "Sandbox-bwrap-Nix: A lightweight sandbox for AI agents and experiments"
article_title: "sandbox-bwrap-nix: A lightweight sandbox for AI agents and experiments"
author: "grigio"
captured_at: "2026-07-25T11:02:09Z"
capture_tool: "hn-digest"
hn_id: 49046290
score: 1
comments: 0
posted_at: "2026-07-25T10:17:13Z"
tags:
  - hacker-news
  - translated
---

# Sandbox-bwrap-Nix: A lightweight sandbox for AI agents and experiments

- HN: [49046290](https://news.ycombinator.com/item?id=49046290)
- Source: [grigio.org](https://grigio.org/sandbox-bwrap-nix-a-lightweight-sandbox-for-ai-agents-and-experiments/)
- Score: 1
- Comments: 0
- Posted: 2026-07-25T10:17:13Z

## Translation

タイトル: Sandbox-bwrap-Nix: AI エージェントと実験用の軽量サンドボックス
記事のタイトル: Sandbox-bwrap-nix: AI エージェントと実験用の軽量サンドボックス
説明: マシン上で AI コーディング ツールを実行している場合は、その恐怖がわかります。エージェントが、実行すべきではないものを rm -rf することにした場合はどうなるでしょうか?ホーム ディレクトリ、設定ファイル、プロジェクトが台無しになったらどうなるでしょうか?
コンテナは役に立ちますが、Docker と Podman は重いです。必要です

記事本文:
llm
注目の
Sandbox-bwrap-nix: AI エージェントと実験用の軽量サンドボックス
マシン上で AI コーディング ツールを実行している場合は、その恐怖がわかります。エージェントが、実行すべきではないものを rm -rf することにした場合はどうなるでしょうか?ホーム ディレクトリ、設定ファイル、プロジェクトが台無しになったらどうなるでしょうか?
コンテナは役に立ちますが、Docker と Podman は重いです。デーモンを実行する必要があります。画像とレイヤーを管理する必要があります。 sudo が必要です。そして、画像がプルされるまで待ちます。
Sandbox-bwrap-nix は、bubblewrap と Nix を組み合わせた最小限のサンドボックスで、コマンド、AI エージェント、実験を実行するための隔離された環境を提供します。 1 秒以内に起動します。デーモン、イメージのプル、root 権限は必要ありません。スクリプトを実行するだけで、クリーンで分離されたシェル内にいることになります。
これは、bubblewrap がインストールされ、フレークが有効になっている Nix を備えた Linux マシンで動作します。
プロジェクトはほんの数個のファイルです。メインのエントリ ポイントは、慎重に選択されたフラグのセットを使用して bwrap を呼び出すシェル スクリプトです。ホストの /nix/store を読み取り専用としてマウントするので、任意の Nix パッケージを使用できます。これにより、サンドボックスに独自の /tmp、独自の /home、および独自のプロセス名前空間が与えられます。ホスト ファイルシステムはサンドボックス内からは見えません。
名前空間が設定されると、スクリプトはサンドボックス内で nix Development を実行します。これにより、git、bun、uv、gnumake、opencode などのツールがすでにインストールされている開発シェルがアクティブになります。見覚えのある bash プロンプトが表示されますが、完全に含まれています。
この設定が真価を発揮する状況がいくつかあります。
まず、AI コーディング エージェントを実行します。 OpenCode などのツールは、コードベースの参照、コマンドの実行、パッケージのインストール、ファイルの編集を行うことができます。効果的なものである必要がありますが、ガードレールも必要です。 Sandbox-bwrap-nix を使用すると、エージェントは実際のホーム ディレクトリを見ることなく仕事を行うことができます。

教会。不正になった場合、被害はサンドボックス内にとどまります。
2 つ目は、Nix 自体を実験することです。 Nix を学習している場合、または新しいフレークをテストしている場合は、ホスト環境を汚染することを心配せずにサンドボックス内で実行できます。後片付けをせずに、何かを試したり、壊したり、新たに始めることができます。
3 番目に、信頼できないコードを実行します。誰かがインターネットからスクリプトを送信します。サンドボックスで実行して、その動作を確認できます。ネットワークにアクセスできます (サンドボックスはネットワークを共有します) が、ファイルにアクセスすることはできません。
サンドボックスは 3 層の保護を提供します。
ディレクトリ。特定のパスのみが表示されます。 Nix ストアは読み取り専用です。ホーム ディレクトリは、いつでもリセットできる Sandbox-home フォルダに置き換えられます。 /tmp は新しい tmpfs です。ファイルシステムの残りの部分は、サンドボックス内には存在しません。
プロセス。サンドボックスは独自の PID 名前空間を使用します。ホスト プロセスは表示されず、ホスト プロセスはサンドボックスのプロセスを表示できません。これにより、エージェントが実行中の他のソフトウェアに干渉するのを防ぎます。
環境。サンドボックスは --clearenv で始まります。これは、ホストから環境変数が漏洩しないことを意味します。環境変数に秘密はなく、迷走した PATH エントリも、ホスト ツールへの誤ったアクセスもありません。サンドボックスは、独自のクリーンな環境を最初からセットアップします。
これは完全なコンテナではありません。イメージ管理、レイヤー キャッシュ、レジストリはありません。ホストのネットワークとカーネルを共有します。重要なのは、Docker を置き換えることではありません。重要なのは、AI エージェントと実験を実行するという特定のユースケース向けに、より軽量なものを用意することです。
ホストにも Nix が必要です。 Nix をまだ使用していない場合は、追加でインストールする必要があります。しかし、Nix エコシステムに属している場合、この設定は自然に感じられます。
リポジトリのクローンを作成し、開始スクリプトを実行します。
git clone https://gith

ub.com/grigio/sandbox-bwrap-nix
cd サンドボックス-bwrap-nix
./start-sandbox.sh
シェル構成にエイリアスを追加して、短いコマンドを使用してどこからでもサンドボックスを起動できるようにすることもできます。
エイリアス sss=/path/to/sandbox-bwrap-nix/start-sandbox.sh
中に入ると、通常のユーザーとして Nix パッケージをインストールできます。オープンコードの実行、ファイルの編集、プロジェクトのビルドなど、通常行うことをすべて行うことができます。違いは、実際のシステムに影響を与えるものがないことです。
Sandbox-home ディレクトリには、git ブランチ情報、Nix ストアからの bash 補完、便利なエイリアスを含む優れたプロンプトを表示する .bashrc が事前に設定されています。フレークと nix-command を有効にする nix.conf があります。また、スキルとプロンプト手順が記載された、すぐに使用できるオープンコード構成フォルダーもあります。
開発シェルには、すぐに使えるこれらのツールが含まれています。
bashプログラム可能な補完を備えたインタラクティブ
flake.nix を編集し、nix flake update を実行することでさらに追加できます。
Sandbox-bwrap-nix を通常の nix 開発の隣に置くと、違いは明らかです。 nix 開発を単独で使用すると、エージェントはファイルシステムに完全にアクセスでき、すべての環境変数を継承し、PID 名前空間を共有し、実際のホーム ディレクトリを使用します。コマンドを 1 つ間違えると、混乱を一掃することになります。
Sandbox-bwrap-nix を使用すると、エージェントは、許可されたものだけを表示します。独自のソフトウェアをインストールする AI ツールの場合、そのコントラストは特に顕著です。サンドボックスがなければ、それらのパッケージは実際の家に入ります。サンドボックスを使用すると、サンドボックス ホームに留まり、リセットすると消えます。
Docker と比較すると、この特定のワークフローでは、sandbox-bwrap-nix の方がシンプルかつ高速です。書き込む Dockerfile がありません。構築するイメージがありません。起動するデーモンがありません。名前空間とバインドマウントだけです。しかし、移植性も低くなります。bwrap がインストールされている Linux でのみ動作し、機能しません。

別の OS やカーネルを提供することはできません。
リポジトリは小さいです。数分ですべてを理解できます。
start-sandbox.sh がエントリ ポイントです。適切なフラグを指定して bwrap を呼び出してから、nix Development を実行します。
flake.nix は、すべてのツールを備えた開発シェルを定義します。
flake.lock は nixpkgs リビジョンを固定します。
Sandbox-home は、サンドボックスにマウントされる分離されたホーム ディレクトリです。
それだけです。 3 つのアクティブ ファイルとホーム ディレクトリ。
Sandbox-bwrap-nix は、特定の問題を簡単な方法で解決します。 AI コーディング エージェントを実行している場合、Nix を実験している場合、または単に何かを試すための使い捨て環境が必要な場合は、一見の価値があります。
高速で、抑制され、邪魔になりません。クローンを作成して試してみて、ワークフローに適合するかどうかを確認してください。
リポジトリは https://github.com/grigio/sandbox-bwrap-nix にあります。
La Bolla dell'AI: スコッピアレによる元外国人ビッグテックの信条を語る
注目の
Oltre la Democrazia: il Modello Cinese di Governance Algoritmica Secondo GLM
注目の
「Il modello non è ciò che conta di più」: la lezione di Zhilin Yang (Kimi CEO)

## Original Extract

If you run AI coding tools on your machine, you know the fear. What if the agent decides to rm -rf something it should not? What if it messes up your home directory, your config files, or your projects?
Containers can help, but Docker and Podman are heavy. You need

llm
Featured
sandbox-bwrap-nix: A lightweight sandbox for AI agents and experiments
If you run AI coding tools on your machine, you know the fear. What if the agent decides to rm -rf something it should not? What if it messes up your home directory, your config files, or your projects?
Containers can help, but Docker and Podman are heavy. You need a daemon running. You need to manage images and layers. You need sudo. And you wait while images pull.
sandbox-bwrap-nix is a minimal sandbox that combines bubblewrap and Nix to give you an isolated environment for running commands, AI agents, and experiments. It boots in under a second. It needs no daemon, no image pulls, and no root privileges. You just run a script and you are inside a clean, isolated shell.
It works on any Linux machine that has bubblewrap installed and Nix with flakes enabled.
The project is just a few files. The main entry point is a shell script that calls bwrap with a carefully chosen set of flags. It mounts your host's /nix/store as read-only so you can use any Nix package. It gives the sandbox its own /tmp, its own /home, and its own process namespace. Your host filesystem is invisible from inside the sandbox.
Once the namespaces are set up, the script runs nix develop inside the sandbox. This activates a dev shell with tools like git, bun, uv, gnumake, and opencode already installed. You land in a bash prompt that looks familiar but is completely contained.
There are a few situations where this setup really shines.
First, running AI coding agents. Tools like OpenCode can browse your codebase, run commands, install packages, and edit files. You want them to be effective, but you also want guardrails. With sandbox-bwrap-nix, the agent can do its job without ever seeing your real home directory. If it goes rogue, the damage stays inside the sandbox.
Second, experimenting with Nix itself. If you are learning Nix or testing a new flake, you can do it inside the sandbox without worrying about polluting your host environment. You can try things, break things, and start fresh with no cleanup.
Third, running untrusted code. Someone sends you a script from the internet. You can run it in the sandbox and see what it does. It has network access (the sandbox shares your network), but it cannot touch your files.
The sandbox gives you three layers of protection.
Directories. Only specific paths are visible. The Nix store comes in as read-only. Your home directory is replaced with a sandbox-home folder that you can reset any time. /tmp is a fresh tmpfs. The rest of your filesystem simply does not exist from inside the sandbox.
Processes. The sandbox uses its own PID namespace. You cannot see host processes, and host processes cannot see the sandbox's processes. This prevents agents from interfering with other running software.
Environment. The sandbox starts with --clearenv, which means no environment variables leak from your host. No secrets in env vars, no stray PATH entries, no accidental access to host tools. The sandbox sets up its own clean environment from scratch.
This is not a full container. There is no image management, no layer caching, no registry. You share the host's network and kernel. The point is not to replace Docker. The point is to have something much lighter for the specific use case of running AI agents and experiments.
You also need Nix on your host. If you are not using Nix already, that is an extra thing to install. But if you are in the Nix ecosystem, this setup feels natural.
Clone the repo and run the start script.
git clone https://github.com/grigio/sandbox-bwrap-nix
cd sandbox-bwrap-nix
./start-sandbox.sh
You can also add an alias to your shell config so you can launch the sandbox from anywhere with a short command.
alias sss=/path/to/sandbox-bwrap-nix/start-sandbox.sh
Once you are inside, you can install any Nix package as a normal user. You can run opencode, edit files, build projects, and do everything you would normally do. The difference is, none of it touches your real system.
The sandbox-home directory comes preconfigured with a .bashrc that gives you a nice prompt with git branch info, bash completion from the Nix store, and useful aliases. There is a nix.conf that enables flakes and nix-command. And there is a ready-to-use opencode config folder with skills and prompt instructions.
The dev shell includes these tools out of the box.
bashInteractive with programmable completion
You can add more by editing flake.nix and running nix flake update.
If you put sandbox-bwrap-nix next to a regular nix develop, the difference is clear. With nix develop alone, your agent has full access to your filesystem, inherits all your environment variables, shares your PID namespace, and uses your real home directory. One wrong command and you are cleaning up a mess.
With sandbox-bwrap-nix, the agent sees only what you let it see. The contrast is especially stark for AI tools that install their own software. Without the sandbox, those packages go into your real home. With the sandbox, they stay in sandbox-home and disappear when you reset it.
Compared to Docker, sandbox-bwrap-nix is simpler and faster for this specific workflow. No Dockerfile to write. No image to build. No daemon to start. Just namespaces and bind mounts. But it is also less portable -- it only works on Linux with bwrap installed, and it does not give you a different OS or kernel.
The repo is small. You can understand all of it in a few minutes.
start-sandbox.sh is the entry point. It calls bwrap with the right flags and then runs nix develop.
flake.nix defines the dev shell with all the tools.
flake.lock pins the nixpkgs revision.
sandbox-home is the isolated home directory that gets mounted into the sandbox.
That is it. Three active files and a home directory.
sandbox-bwrap-nix solves a specific problem in a simple way. If you run AI coding agents, if you experiment with Nix, or if you just want a throwaway environment for trying things, it is worth a look.
It is fast, it is contained, and it does not get in your way. Clone it, try it, and see if it fits your workflow.
The repo is at https://github.com/grigio/sandbox-bwrap-nix
La Bolla dell'AI: Perché un Ex-Ingegnere Big Tech Crede che Stia per Scoppiare
Featured
Oltre la Democrazia: il Modello Cinese di Governance Algoritmica Secondo GLM
Featured
"Il modello non è ciò che conta di più": la lezione di Zhilin Yang (Kimi CEO)
