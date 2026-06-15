---
source: "https://fedoramagazine.org/sandbox-ai-coding-agents-with-microvms-on-fedora-linux/"
hn_url: "https://news.ycombinator.com/item?id=48540375"
title: "Sandbox AI coding agents with microVMs on Fedora Linux"
article_title: "Sandbox AI coding agents with microVMs on Fedora Linux - Fedora Magazine"
author: "vitorsr"
captured_at: "2026-06-15T14:15:54Z"
capture_tool: "hn-digest"
hn_id: 48540375
score: 1
comments: 0
posted_at: "2026-06-15T12:38:16Z"
tags:
  - hacker-news
  - translated
---

# Sandbox AI coding agents with microVMs on Fedora Linux

- HN: [48540375](https://news.ycombinator.com/item?id=48540375)
- Source: [fedoramagazine.org](https://fedoramagazine.org/sandbox-ai-coding-agents-with-microvms-on-fedora-linux/)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T12:38:16Z

## Translation

タイトル: Fedora Linux 上の microVM を使用したサンドボックス AI コーディング エージェント
記事のタイトル: Fedora Linux 上の microVM を使用したサンドボックス AI コーディング エージェント - Fedora Magazine
説明: Linux ディストリビューションには、プロセス分離のオプションが付属しています。この記事では、microVM を使用して AI コーディング エージェントを実行する方法を説明します。

記事本文:
Fedora Linux 上の microVM を使用したサンドボックス AI コーディング エージェント - Fedora マガジン
フォストドン
Fedora Linux 上の microVM を使用したサンドボックス AI コーディング エージェント
Unsplash の Immo Wegmann による写真
Justin Wheeler が Fedora コミュニティでの成長について語る
Microsoft セキュア ブート証明書の有効期限について知っておくべきこと: パニックにならないでください。
ヤロスラフ・レズニクがセキュリティ、EU サイバーレジリエンス法、そして机の後ろから物事を行うことができない理由について語ります。
Claude コードや Codex などの AI コーディング エージェントの能力は毎月向上しています。これは生産性にとっては素晴らしいことですが、すべてのコマンドを承認するのはすぐに面倒になってしまいます。一方、エージェントが作業用マシン上で任意のコマンドを実行できるようにするのは、良いアイデアではありません。彼らは、kubectl を使用して運用クラスターを探索したり、SSH 経由で運用サーバーでリモート コマンドを実行したりすることに非常に優れています。
幸いなことに、Linux ディストリビューションには、プロセスを分離するためのオプションが豊富に用意されています。エージェントは、コンテナーまたは VM 内でまったく別のユーザーとして実行できます。この記事では、microVM を使用してコーディング エージェントを実行する方法を説明します。
AI エージェントを無人モードで実行することは、信頼できないコードを実行することに似ています。 Anthropic や Google など、これらのエージェントを背後に持つ企業は認証情報を盗もうとしているわけではありませんが、人々は事実上どこでも、スロップスクワッティングやプロンプト インジェクションなどの新しい攻撃ベクトルを考案し続けています。
コーディング エージェント自体には、たとえばここで説明されているように、プロンプト インジェクションを拒否しようとする組み込みの緩和策が付属しています。
軽量サンドボックス テクノロジは、コーディング エージェントにおけるもう 1 つの防御層です。 Linux では、bwrap が可能な実装の 1 つです。これによりハードルが高くなりますが、サンドボックスからの脱出は依然として問題です。マルチプラットフォーム サンドボックス エスケープの例として CVE-2026-39861 を見てください。
コンテナを使用してエージェントを隔離することもできます。

独自の名前空間を持ちますが、依然としてホスト カーネルを共有します。最近のカーネルの脆弱性の一部は、特権昇格 (通常のユーザーから root への切り替え) を引き起こし、コンテナーがセキュリティ境界として十分ではないことを示唆しています。
この記事の残りの部分では、microVM を使用して Fedora Linux 上でコーディング エージェントを簡単にサンドボックス化する方法について説明します。
まず、microVM とは何かを見てみましょう。他の VM と同様に、各 microVM ごとに 1 つずつ、独自のカーネルがあります。従来の VM と比較すると、非常に短時間 (数百ミリ秒) で起動しますが、完全な VM のすべての機能は提供されません。
この記事では、podman 用の krun ランタイムの使用方法について説明します。このアプローチは、コンテナーと同じよく知られたワークフローを提供しますが、すべてのコンテナーを microVM として実行するだけです。
まずランタイムをインストールします。
dnf インストール クランクルン
microVM を実行するには、ターミナルで –runtime=krun を指定して podman を実行するだけです。
ポッドマン run --runtime=krun --rm -it fedora:44 /bin/bash
注意すべきこと
microVM は通常のコンテナではないため、いくつかの動作が異なる場合があります。まず、krun アノテーションを使用して十分な CPU と RAM を割り当てます。デフォルトは小さすぎるため、OOM (メモリ不足) による強制終了が発生する可能性があります。次に、libkrun バージョンが 1.8 以上であることを確認します。古いバージョンにはバグがあり、コーディング エージェントで Enter キーを押すことができません。 3 番目に、microVM は Dockerfile に設定された USER を無視し、常に root として起動します。正しいユーザーに手動で切り替えるか、エントリポイント スクリプトに切り替えを組み込みます。
ケーススタディ: Python プロジェクトのクロード コードをサンドボックス化する
このセクションでは、 uv によって管理される Python プロジェクトの簡単なセットアップの概要を説明します。 podman-compose を使用してプロジェクトを microVM にマウントします。コンテナーと比較して、このポッドマン構成には、UID/GID 変換、SEL のための追加のアノテーションが必要です。

inux のラベル付け、および HW リソース。最終的なセットアップは、コンテナーに必要なものと非常に似ています。
公式 Fedora リポジトリから podman compose をインストールするには、次を実行します。
dnf インストール podman-compose
セットアップには 3 つの部分があります。
上で述べたように、krun ランタイムを備えた podman は引き続きコンテナーを実行しますが、それぞれのコンテナーを microVM 内に生成します。このサンプル コンテナには、uv パッケージ マネージャー、クロード コード、およびいくつかの追加の RPM パッケージが含まれています。プロジェクトの依存関係とプログラミング言語に基づいて独自のコンテナを定義します。
必ず非特権ユーザーを作成し、それをエージェントの実行に使用してください。
フェドラから:44
ARG ホスト_UID=1000
ARG ホスト_GID=1000
# ホスト UID/GID に一致するグループとユーザーを作成します
RUN groupadd -g ${HOST_GID} appuser && \
useradd -u ${HOST_UID} -g ${HOST_GID} -m appuser
RUN mkdir -p /venv && chown appuser:appuser /venv
RUN mkdir -p /home/appuser/.claude && chown appuser:appuser /home/appuser/.claude
USERアプリユーザー
# めったに変更されないツール。 dnf レイヤーの上に保持されるため、RPM リストを編集します
# 以下の方法では、これらのインストールは無効 (および再実行) されません。
実行curl -LsSf https://astral.sh/uv/install.sh |し&&\
カール -fsSL https://claude.ai/install.sh |バッシュ
ユーザールート
# 頻繁に変更される RPM。最後に保持されるため、パッケージを追加するとここから下のみが再構築されます。
RUN dnf install git make vim free libpq-devel python3-devel gcc -y && \
DNFすべてをクリーンアップ
COPY --chown=appuser エントリーポイント.sh /エントリーポイント.sh
chmod +x /entrypoint.sh を実行します。
USERアプリユーザー
WORKDIR /アプリ
# エントリポイントの PATH に .local/bin がないため、これが必要です
ENV PATH="/home/appuser/.local/bin:$PATH"
エントリーポイント ["/entrypoint.sh"]
CMD ["/bin/bash"]
docker-compose.yaml
構成ファイルは、プロジェクト ディレクトリを microVM にマウントする方法を定義します。 podman は UID/GID を変換し、S を管理する必要があるため、ほとんどの魔法がここで発生します。

ELinux ラベルを付けないと、microVM 内でファイルにアクセスできなくなるか、別のユーザーが所有することになります。
サービス:
クロード:
コンテナ名: プロジェクト名-クロード
注釈:
run.oci.handler: krun
krun.ram_mib: "4096"
krun.cpus: "4"
ユーザー: "${HOST_UID}:${HOST_GID}"
userns_mode: keep-id # オプション、ルートレスホストの場合
ビルド:
コンテキスト: 。
引数:
HOST_UID: "${HOST_UID}" # コンテナ内で作成されたファイルに正しい権限が与えられるように、ホストの UID と GID を使用します
HOST_GID: "${HOST_GID}"
ボリューム:
- ../:/app:U,z # プロジェクトをバインドマウントします
- プロジェクト名-venv-cache:/venv:U,z
- claude-config:/home/appuser/.claude:U,z # 永続的なクロード資格情報/config
作業ディレクトリ: /app
stdin_open: true
ティ：本当
環境:
- CLAUDE_CONFIG_DIR=/home/appuser/.claude
- UV_LINK_MODE=コピー
- UV_PROJECT_ENVIRONMENT=/venv/env # これはキャッシュされたボリューム内にあります
- UV_PYTHON_INSTALL_DIR=/venv/python # UV 管理の Python インストールがホームではなく /venv にキャッシュされるようにする
- TERM=xterm-256color
- COLORTERM=トゥルーカラー
ボリューム:
プロジェクト名-venv-cache:
クロード構成:
外部: true
名前: クロード構成
重要な部分は 3 つあります。
注釈 – これらは、krun をランタイムとして選択し、HW 要件を指定します。
user と userns_mode – これは、microVM で作成されたファイルが最終的にホスト上のユーザーによって所有されるように、UID/GID を変換するように podman に指示します。
ボリューム ラベル – z は、podman に共有 SELinux ラベルを使用してファイルのラベルを付け直すように指示します。そうしないと、SELinux によって、microVM 内のプロセスがボリューム内のファイルにアクセスできなくなります。 U は、podman にすべてのファイルを再帰的に chown するように指示します。
動的依存関係をコンテナー イメージに焼き付けたくないため、エントリポイントは Python プロジェクトの仮想環境を作成します。また、podman には root から通常のユーザーへの切り替えも実行されます。

krun ランタイムは、コンテナーからの USER ディレクティブを無視します。
#!/bin/bash
セット -e
echo "サンドボックスはユーザー: $(id -un)、ディレクトリ: $(pwd) として開始されました"
if [ "$(id -un)" != "appuser" ];それから
runuser -u appuser -- UV 同期
echo "アプリユーザーとして ${@} を実行しています"
exec runuser -u appuser -- "$@"
フィ
UV同期
「$@」を実行
セットアップを実行する
$ HOST_UID=$(id -u) HOST_GID=$(id -g) podman-compose -f .agent-sandbox/docker-compose.yaml build
ステップ 1/18: フェドーラから:44
...
localhost/agent-sandbox_claude:latest のタグ付けに成功しました
次に、外部ボリュームを作成し、クロード コンテナを対話的に実行します。
$ podman volume create claude-config
$ HOST_UID=$(id -u) HOST_GID=$(id -g) podman-compose -f .agent-sandbox/docker-compose.yaml run --rm クロード
サンドボックスはユーザー: root としてディレクトリ: /app で開始されました
6 ミリ秒で 3 つのパッケージを解決
1ミリ秒で3つのパッケージをチェックしました
/bin/bash を appuser として実行する
tty: ttyname エラー: デバイスの ioctl が不適切です
[appuser@3bd1234b9a77 app]$
microVM 内で uname -a を実行することで、カーネルが異なることを確認できるようになりました。
まとめ: セットアップ全体を作成する単一のスクリプト
すべてのプロジェクトに対して同じセットアップを手動で作成することは、最高のユーザー エクスペリエンスではありませんが、次のような単純なスクリプトを使用してセットアップを自動化できます。これは、上記のセットアップを 3 つの単純なコマンド オプション (init、build、run) にラップする新しい sbx コマンドをインストールします。
注意事項 — microVM は完全な境界ではありません
microVM は基準を大幅に引き上げますが、完全な分離ではないため、そのように提示するのは無責任です。セキュリティ モデルの詳細については、libkrun git リポジトリを参照してください。
危険な可能性のあるソフトウェアを実行したい場合は、完全な VM またはクラウド VM を使用することをお勧めします。
MicroVM は AI エージェントを実行するのに最適なスポットのようです。これらは、c の使い慣れたワークフローを提供します。

コンテナーですが、エージェントはハイパーバイザーの背後にある独自のカーネルで実行されます。この記事では、podman および krun ランタイムに基づくワークフローについて説明します。これは、Fedora Linux には両方がネイティブに同梱されているためですが、どのプラットフォームでも利用できる他のオプションが多数あります ( docker Sandbox など)。
AI の使用に関する注意: この記事は私自身が書きました。私は Claude (Anthropic) を使用して、文法、言葉遣い、文章の構造を大幅に改良しました。技術的な内容とすべての主張は私自身のものであり、テスト済みです。
読み込み中…
Fedora プロジェクト コミュニティ
マーティン・セーナウトカ
Discussion.fedoraproject.org でディスカッションを開始してください
Fedora Linux 44 は現在入手可能です。詳細については、リリースのお知らせをお読みください。
電子メールで Fedora マガジンを購読する
Fedora Magazine では寄稿者を募集しています。
このウェブサイトに掲載されている意見は各著者によるものであり、
著者の雇用主または Red Hat のものではありません。 Fedora Magazine が目指しているのは、
クリエイティブ コモンズ ライセンスに基づいてすべてのコンテンツを公開できますが、公開できない場合があります。
あらゆる場合にそうすること。あなたには、次のものがあることを確認する責任があります。
このサイト上の作品を再利用するために必要な許可。フェドラのロゴは、
Red Hat, Inc. の商標です。 利用規約

## Original Extract

Linux distributions come with options for process isolation. This article shows how to use microVMs to run AI coding agents.

Sandbox AI coding agents with microVMs on Fedora Linux - Fedora Magazine
fosstodon
Sandbox AI coding agents with microVMs on Fedora Linux
Photo by Immo Wegmann on Unsplash
Justin Wheeler on Growing Up in the Fedora Community
What you need to know about the Microsoft Secure Boot certificate expiration: Don’t Panic!
Jaroslav Reznik on Security, the EU Cyber Resilience Act, and Why You Can't Do Things From Behind a Desk!
AI coding agents such as Claude code or Codex get more capable every month. This is great for productivity, but approving all commands gets annoying really quickly. On the other hand, allowing agents to run any command on your work machine is not a great idea. They are really good at exploring your production cluster using kubectl or running remote commands at your production servers over SSH.
Fortunately, Linux distributions come with plenty of options for process isolation. You can run agents as a completely different user, in a container, or in a VM. This article shows how to use microVMs to run coding agents.
Running AI agents in unattended mode is like running untrusted code. Companies behind these agents, such as Anthropic or Google, are not trying to steal credentials, but people keep coming up with new attack vectors like Slopsquatting or prompt injections virtually anywhere .
The coding agents themselves ship with built-in mitigations that try to refuse prompt injections as described, for example, here .
Lightweight sandboxing technologies are another layer of defense in coding agents. On Linux, bwrap is one of the possible implementations. This raises the bar, yet sandbox escapes are still a problem. Take a look at CVE-2026-39861 as an example of multi-platform sandbox escape.
You could use containers to isolate the agent in their own namespace, but they still share the host kernel. Some of the the recent kernel vulnerabilities resulted in privilege escalation (switching from regular user to root) suggesting that containers are not enough as a security boundary .
In the rest of this article, I describe how to use microVMs to easily sandbox coding agents on your Fedora Linux.
First of all, let’s take a look at what microVMs are. Just like any VM, they have their own kernel, one per each microVM. Compared to traditional VMs they start in very short time (hundreds of milliseconds) but don’t offer all the features of full VMs.
This article explains usage of krun runtime for podman. This approach offers the same well-known workflow as containers, but simply runs every container as a microVM.
Start by installing the runtime:
dnf install crun-krun
To run a microVM, simply run podman with –runtime=krun in your terminal:
podman run --runtime=krun --rm -it fedora:44 /bin/bash
Things to watch out for
A microVM is not a regular container, so a few things might behave differently. First, allocate enough CPU and RAM with krun annotations. The defaults are too small and might result in OOM (Out Of Memory) kills. Second, make sure you have libkrun version >= 1.8. Older versions have a bug which prevents you from pressing Enter in your coding agent. Third, the microVM ignores the USER set in the Dockerfile and always boots as root. Either switch to the correct user manually or put the switch into an entrypoint script.
Case study: sandboxing Claude Code for a Python project
This section outlines a simple setup for a Python project managed by uv . It uses podman-compose to mount the project into the microVM. Compared to containers, this podman compose needs additional annotations for UID/GID translation, SELinux labeling, and HW resources. The final setup is very similar to what you would need for containers.
To install podman compose from official Fedora repositories, run:
dnf install podman-compose
The setup has 3 parts:
As mentioned above, podman with krun runtime still runs containers, but spawns each of them in a microVM. This example container includes uv package manager, claude code and a few additional RPM packages. Define your own container based on your project dependencies and programming language.
Make sure to create an unprivileged user and use it for running the agent.
FROM fedora:44
ARG HOST_UID=1000
ARG HOST_GID=1000
# Create group and user matching host UID/GID
RUN groupadd -g ${HOST_GID} appuser && \
useradd -u ${HOST_UID} -g ${HOST_GID} -m appuser
RUN mkdir -p /venv && chown appuser:appuser /venv
RUN mkdir -p /home/appuser/.claude && chown appuser:appuser /home/appuser/.claude
USER appuser
# Rarely-changing tooling. Kept above the dnf layer so editing the RPM list
# below does not invalidate (and re-run) these installs.
RUN curl -LsSf https://astral.sh/uv/install.sh | sh && \
curl -fsSL https://claude.ai/install.sh | bash
USER root
# Frequently-changing RPMs. Kept last so adding a package only rebuilds from here down.
RUN dnf install git make vim free libpq-devel python3-devel gcc -y && \
dnf clean all
COPY --chown=appuser entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
USER appuser
WORKDIR /app
# This is needed because entrypoint does not have .local/bin in the PATH
ENV PATH="/home/appuser/.local/bin:$PATH"
ENTRYPOINT ["/entrypoint.sh"]
CMD ["/bin/bash"]
docker-compose.yaml
The compose file defines how to mount the project directory into the microVM. This is where most of the magic happens, because podman needs to translate UID/GID and manage SELinux labels, otherwise the files would not be accessible inside of the microVM or they would end up being owned by a different user.
services:
claude:
container_name: project-name-claude
annotations:
run.oci.handler: krun
krun.ram_mib: "4096"
krun.cpus: "4"
user: "${HOST_UID}:${HOST_GID}"
userns_mode: keep-id # optional, for rootless host
build:
context: .
args:
HOST_UID: "${HOST_UID}" # use UID and GID from the host so that files created in the container have correct permissions
HOST_GID: "${HOST_GID}"
volumes:
- ../:/app:U,z # bind mount your project
- project-name-venv-cache:/venv:U,z
- claude-config:/home/appuser/.claude:U,z # persistent claude credentials/config
working_dir: /app
stdin_open: true
tty: true
environment:
- CLAUDE_CONFIG_DIR=/home/appuser/.claude
- UV_LINK_MODE=copy
- UV_PROJECT_ENVIRONMENT=/venv/env # This is inside the cached volume
- UV_PYTHON_INSTALL_DIR=/venv/python # So that uv-managed python installations are not in home but cached in /venv
- TERM=xterm-256color
- COLORTERM=truecolor
volumes:
project-name-venv-cache:
claude-config:
external: true
name: claude-config
There are 3 key parts:
annotations – these select krun as a runtime and specify HW requirements
user and userns_mode – this tells podman to translate UID/GID so that the files created in the microVM end up owned by your user on the host
volume labels – z tells podman to relabel the files with a shared SELinux label. Otherwise SELinux would prevent the process inside the microVM from touching the files in the volume. U tells podman to recursively chown all files.
The entrypoint creates a virtual environment for the Python project, because we don’t want dynamic dependencies baked into the container image. It also runs the switch from root to regular user because podman with krun runtime ignores the USER directive from the container.
#!/bin/bash
set -e
echo "Sandbox started as user: $(id -un) in directory: $(pwd)"
if [ "$(id -un)" != "appuser" ]; then
runuser -u appuser -- uv sync
echo "Running ${@} as appuser"
exec runuser -u appuser -- "$@"
fi
uv sync
exec "$@"
Run the setup
$ HOST_UID=$(id -u) HOST_GID=$(id -g) podman-compose -f .agent-sandbox/docker-compose.yaml build
STEP 1/18: FROM fedora:44
...
Successfully tagged localhost/agent-sandbox_claude:latest
Then create the external volume and run the claude container interactively:
$ podman volume create claude-config
$ HOST_UID=$(id -u) HOST_GID=$(id -g) podman-compose -f .agent-sandbox/docker-compose.yaml run --rm claude
Sandbox started as user: root in directory: /app
Resolved 3 packages in 6ms
Checked 3 packages in 1ms
Running /bin/bash as appuser
tty: ttyname error: Inappropriate ioctl for device
[appuser@3bd1234b9a77 app]$
You can now check that the kernel is different by running uname -a inside of the microVM.
Putting it together: single script to create the whole setup
Creating the same setup manually for every project is not the greatest user experience, but you can automate the setup using a simple script like this . It installs a new sbx command that wraps the setup described above into 3 simple command options: init, build, and run.
A word of caution — microVM is not a bulletproof boundary
A microVM raises the bar considerably, but it is not perfect isolation, and it would be irresponsible to present it as such. Take a look at the libkrun git repository to read more about the security model .
If you want to run software that might be dangerous, prefer using a full VM or even cloud VM.
MicroVMs seem like a sweet spot for running AI Agents. They provide a familiar workflow of containers, but the agents run on their own kernel behind a hypervisor. This article describes workflow based on podman and krun runtime because Fedora Linux ships both of them natively, but there are plenty of other options available for any platform (for example docker sandbox ).
Note about AI usage: I wrote this article myself. I used Claude (Anthropic) to significantly refine the grammar, wording, and sentence structure; the technical content and all claims are my own and tested.
Loading…
Fedora Project Community
Martin Sehnoutka
Start the discussion at discussion.fedoraproject.org
Fedora Linux 44 is available now . Read the release announcement for all the details.
Subscribe to Fedora Magazine via Email
Fedora Magazine is looking for contributors!
The opinions expressed on this website are those of each author,
not of the author's employer or of Red Hat. Fedora Magazine aspires to
publish all content under a Creative Commons license but may not be able
to do so in all cases. You are responsible for ensuring that you have the
necessary permission to reuse any work on this site. The Fedora logo is a
trademark of Red Hat, Inc. Terms and Conditions
