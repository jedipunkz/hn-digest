---
source: "https://www.pillar.security/blog/one-docker-socket-to-rule-them-all-escaping-codex-cursor-and-gemini-clis-sandboxes"
hn_url: "https://news.ycombinator.com/item?id=49003857"
title: "One Docker socket to rule them all: Escaping Codex, Cursor, and Gemini CLI"
article_title: "The Week of Sandbox Escapes Day 2: One Docker socket to rule them all: escaping Codex, Cursor, and Gemini CLI's sandboxes"
author: "BlueBerry2001"
captured_at: "2026-07-22T10:15:23Z"
capture_tool: "hn-digest"
hn_id: 49003857
score: 2
comments: 0
posted_at: "2026-07-22T09:18:40Z"
tags:
  - hacker-news
  - translated
---

# One Docker socket to rule them all: Escaping Codex, Cursor, and Gemini CLI

- HN: [49003857](https://news.ycombinator.com/item?id=49003857)
- Source: [www.pillar.security](https://www.pillar.security/blog/one-docker-socket-to-rule-them-all-escaping-codex-cursor-and-gemini-clis-sandboxes)
- Score: 2
- Comments: 0
- Posted: 2026-07-22T09:18:40Z

## Translation

タイトル: すべてを支配する 1 つの Docker ソケット: Codex、カーソル、および Gemini CLI のエスケープ
記事のタイトル: サンドボックス エスケープの週 2 日目: 1 つの Docker ソケットですべてを支配: Codex、Cursor、および Gemini CLI のサンドボックスをエスケープ

記事本文:
サンドボックス エスケープの週 2 日目: 1 つの Docker ソケットですべてを支配: Codex、Cursor、および Gemini CLI のサンドボックスをエスケープ
SAIL 2.0 フレームワークの紹介: AI エージェントを保護するための実践ガイド
SAIL 2.0 フレームワークの紹介: AI エージェントを保護するための実践ガイド
プラットフォームの概要 AI ディスカバリーとポスチャー レッド チーミングと攻撃対象領域の露出 ランタイム ガードレール ガバナンスとコンプライアンス 最新の発表 セキュリティの柱、AI ソフトウェア セキュリティ プラットフォーム ソリューションの 2026 年 Gartner® クール ベンダーに選出 (5) ソリューション ユース ケース 自社製 AI エージェントティック エンドポイント AI ゲートウェイ セキュリティ MCP およびツール セキュリティ エージェントティック AI セキュリティ 組み込み AI 業界 ヘルスケア 金融テクノロジー リソース (4) リソース ブログ Pillar Research SAIL 2.0フレームワーク 最新リソース The Week of Sandbox Escapes 会社 (4) 会社概要 ニュースルーム パートナー 採用情報 デモを入手 デモを入手 ブログ
サンドボックス エスケープの週 2 日目: 1 つの Docker ソケットですべてを支配: Codex、Cursor、および Gemini CLI のサンドボックスをエスケープ
この投稿は、AI コーディング エージェントがサンドボックス化されたワークスペースのアクションとサンドボックス化されていないホストの実行の間の境界線をどのように越え続けるかについてのシリーズ、The Week of Sandbox Escapes の一部です。ここでマスターレポートをお読みください。
以前、antigravit の Sandbox-exec 実装に関する投稿 (まだ読んでいない場合は、今すぐ読むことをお勧めします) で、開発者が antigravit のシートベルト プロファイルの (デフォルトを許可する) 設定を決定したことについて説明しました。その後、「拒否リスト」設計パターンを使用するこの設定は制限ではなく提案であることがわかりました。拒否ルールを永久に「Yes-I-said-No」ポリシーにボルトで固定し続けることができ、それを回避する新しいトリックを見つけ続けることになります。短期的な解決策は、シートベルト プロファイルで (deny de で始まる) 許可リストを作成することです。

障害）設定を変更し、エージェントが適切に動作する必要があるものだけを元に戻すことを許可します。
したがって、次の当然の疑問は、ベンダーが実際にそれを行うと何が起こるかということでした。 Codex CLI、Cursor、および Gemini CLI はすべて慎重な判断を行いました。彼らのプロフィールは閉じられ始めます、紙の上では、反重力で私が通ったドアは釘で閉められています。
そうです。それで、別のドアを探しに行きます。
そのドアは Docker Desktop であり、これらのエージェントのいずれかを実行する基本的にすべての Mac にインストールされます。 (デフォルトを拒否) 設定に影響を与える普遍的な脆弱性を発見し、同じ発見結果を 3 つのベンダーすべてに送信しました。興味深いのは、3 つのチームが 1 つのほぼ同一のチェーンを調べ、3 つの異なる答えを返したことです。1 つは修正し、1 つは情報としてファイルし、もう 1 つはすでにドキュメントにあると言いました。
デフォルトを拒否しても得られないもの
クローズドプロファイルは、エージェント自身のプロセスがカーネルに何を要求できるかを制御します。特権付きローカル デーモンがエージェントに代わって何を行うかは制御されません。これらは別の質問であり、Docker Desktop は完全にギャップの中に存在します。
これらのエージェントをサンドボックスとネットワーク アクセスをオンにして実行すると、書き込みがワークスペースと一時ディレクトリに固定されます。それが約束です。元に戻すデーモンとその Docker デスクトップはすでに実行されています。
これはチェーンです。これは 3 つの製品すべてで同じチェーンです。
サンドボックスは、ワークスペース外へのファイルの書き込みをブロックします。プロセスの実行、ファイルの読み取り、送信ネットワークはブロックされません (無効になっていない限り)。その組み合わせだけで十分です。
これにより、docker CLI を含むあらゆるバイナリを実行できます。
Docker ソケットを含むあらゆるファイルを読み取ることができます。そのソケットでの読み取りと書き込みは、実質的にデーモンに対する root になります。 macOS では、Unix ソケットへの connect() はネットワーク操作として分類されます。

重要なのは一瞬です。
ネットワーク アクセスを使用すると、コンテナー rootfs tarball をワークスペースにカールして、それを docker インポートできます。レジストリのプルがないため、サンドボックスのイメージ制限が発動することはありません。
docker run --privileged は、サンドボックスがすでに許可しているディレクトリを指すバインド マウントを使用します。何もトリップしません。
特権コンテナ内で、-t virtiofs virtiofs0 ~/mac-home をマウントします。 VirtioFS は Docker Desktop 独自のホスト共有メカニズムであり、ユーザーのホーム ディレクトリ全体を読み取りと書き込みで提供します。
エスケープは、サンドボックスの管轄外にある Docker VM 内で行われます。エージェントの観点からは、違法なことは何も起こっていません。すべての書き込みは、許可されたワークスペースで行われます。 SSH キー、認証情報、ブラウザ データ、すべてを 1 回のアクセスで入手できます。
まず、攻撃者はアーキテクチャを検出し、カール上で Alpine minirootfs をプルし、それをローカル イメージとして Docker インポートします。これにより、レジストリの制限が影響を受けなくなります。これらはすべて、単純なプロンプト インジェクションによって行われます。
次に、攻撃者はデタッチされた特権付きコンテナを実行し、サンドボックスがすでに許可しているディレクトリのみをバインドし、docker exec を通じて VirtioFS をマウントします。
最後に、攻撃者はマウント内でユーザーのホーム ディレクトリを見つけ、ペイロードを .zshrc に書き込みます。
ユーザーがターミナルを開くたびに電卓が起動するようになりました。開いた電卓を好きなものと交換してください。重要なのは、それに応答するプロセスがサンドボックス内に存在したことがないということです。
そして、どれも直接のプロンプトを必要としません。概念実証では、再現するのが最も簡単なため、明示的な指示で実行しますが、同じ手順を README、コード コメント、または通常の作業中にエージェントが読むドキュメントにドロップし、残りは間接インジェクションによって行われます。
PoC ビデオ全体はここでご覧いただけます。
カーソルのシートベルト プロファイルは起動サービスを制限しているため、エージェントは

ボックスの外でプロセスを生成しません。ただし、サンドボックスでの自動実行がオン (デフォルト) の場合、自動実行ネットワーク アクセスをオンにする必要がありますが、エージェントは Docker に到達し、Docker はシートベルトを気にしません。
Cursor は問題をトリアージし、Docker ソケットをサンドボックスの外側から実行する必要があることに同意し、修正を配布し、次のビルドで問題が再現しなくなりました。重症度は高く安定しました。このアドバイザリは、https://github.com/cursor/cursor/security/advisories/GHSA-v4xv-rqh3-w9mc で読むことができます。
Codex: 意図したものではありませんが、バグでもありません
コーデックスには同一の暴露が掲載されています。ワークスペース書き込みサンドボックスとネットワーク アクセスを使用して起動します。
上記と同じプレーンな Docker スクリプトを与えると、すぐに実行されます。 Codex は許可されたワークスペース内で自動実行されるため、承認ポリシーが起動されることはありません。
OpenAI のレビューでは、これは意図された動作ではなく、エージェントがワークスペースの境界を越えてその外側に書き込むパスに追い込まれる可能性があることに同意しました。次に、すべてのセッションにヒットするデフォルトのパスではなく、構成と環境に依存するように絞り込みました。彼らはレポートを支払い、情報としてマークし、CVE を拒否しました。
彼らの推論では、現在の権限モデルでは、アクティブなプロファイルでソケットが明示的に許可されていない限り、Unix ソケットへのアクセスはネットワーク プロキシによって仲介され、デフォルトで拒否されるということです。
Docker と通信するには明示的な許可設定が必要なため、これをデフォルトのサンドボックス エスケープとはみなしません。
3つの判決の中で最も短い判決。同じチェーン、同じ特権付き virtiofs0 マウント、同じ結果。 Gemini CLI チームの立場は、リスクはドキュメントに詳しく記載されているため、彼らに関する限りこれはバグではないというものでした。
サンドボックスは、ブロックするシステムコールに基づいて判断されます。また、どのようなデーモンを実行できるかについても判断される必要があります。

ああ。読み取り可能なソケットを備えた特権付きローカル デーモンは、1 カール離れた場所にある完全にサンドボックス化されていない 2 番目の実行環境であり、Docker Desktop はデフォルトで開発者マシンの大部分に 1 つを配置します。デフォルトの拒否は正しい選択であり、Launch Services の扉を真に閉じました。エージェントが電話を取って、デーモンにその代わりに歩いてもらうように依頼できる場合には、役に立ちません。
全体的な長期的な解決策は、gVisor などのカーネル、実際の実行コンテナ、そして最後にループ内の人間へのユーザーランド呼び出しを監視する封じ込めソリューションを使用して、自律エージェントの動作を適切に制限することです。私たちはこの新しいテクノロジーを前進させます。
購読して最新のセキュリティアップデートを入手してください
もしかしたらあなたもこれが面白いと思うかもしれません
サンドボックス エスケープ週間 4 日目: Git ディレクトリを .git と呼ぶ必要はない
サンドボックス エスケープの週 3 日目: サンドボックスで venv を編集でき、別の何かがそれを実行しました
サンドボックス エスケープの週 7 日目: .vscode の時限爆弾: Antigravity のセキュア モードをバイパスする
© 2026 柱。無断転載を禁じます。

## Original Extract

The Week of Sandbox Escapes Day 2: One Docker socket to rule them all: escaping Codex, Cursor, and Gemini CLI's sandboxes
Introducing SAIL 2.0 Framework: A Practical Guide to Secure AI Agents
Introducing SAIL 2.0 Framework: A Practical Guide to Secure AI Agents
Platform overview AI Discovery & Posture Red Teaming & Attack Surface Exposure Runtime Guardrails Governance & Compliance Latest Announcement Pillar Security Named as a 2026 Gartner® Cool Vendor in AI Software Security Platform Solutions (5) Solutions Use cases Homegrown AI Agentic Endpoint AI Gateway Security MCP & Tool Security Agentic AI Security Embedded-AI Industry Healthcare Financial Technology Resources (4) Resources Blog Pillar Research SAIL 2.0 Framework Latest Resource The Week of Sandbox Escapes Company (4) Company About Us Newsroom Partners Careers Get a demo Get a demo Blog
The Week of Sandbox Escapes Day 2: One Docker socket to rule them all: escaping Codex, Cursor, and Gemini CLI's sandboxes
This post is part of The Week of Sandbox Escapes, a series on how AI coding agents keep crossing the line between sandboxed workspace actions and unsandboxed host execution. Read the master report here .
Previously, in our post regarding antigravity’s sandbox-exec implementation(which if you haven’t read I suggest you do so now) we discussed that the developers decided to us (allow default) setting on antigravit’s seatbelt profile, we then established that this setting which uses the “deny-list” design pattern is more of a suggestion rather than a restriction. You can keep bolting deny rules onto a yes-unless-I-said-no policy forever and you'll keep finding new tricks to bypass it. The short-term fix would be to create an allow-list, which in a seatbelt profile starts with the (deny default) setting and allows back only what the agent absolutely needs to operate properly.
So the obvious next question was what happens when a vendor actually does that. Codex CLI, Cursor, and Gemini CLI all made the careful call. Their profiles start closed, On paper, the door I walked through on Antigravity is nailed shut.
It is. So you go find a different door.
That door is Docker Desktop, and it's installed on basically every Mac that runs one of these agents. I discovered a universal vulnerability that affected the (deny default) setting and I sent the same finding to all three vendors. The interesting part is that three teams looked at one near-identical chain and returned three different answers: one fixed it, one filed it as informational, and one told me it was already in the docs.
What deny-default doesn't buy you
A closed profile controls what the agent's own process can ask the kernel for. It does not control what a privileged local daemon will do on the agent's behalf. Those are different questions, and Docker Desktop lives entirely inside the gap.
Run any of these agents with its sandbox on and network access on, and writes are pinned to your workspace and temp directories. That's the promise. The daemon that undoes it is already running and its Docker Desktop.
Here's the chain, and it's the same chain on all three products:
The sandbox blocks file writes outside the workspace. It does not block process execution, file reads, or outbound network(unless disabled). That combination is all you need.
It lets you execute any binary, including the docker CLI.
It lets you read any file, including the Docker socket. Read plus write on that socket is effectively root over the daemon. On macOS the connect() to that Unix socket is classified as a network operation, which matters in a second.
Network access lets you curl a container rootfs tarball into the workspace and docker import it. No registry pull, so the sandbox's image restrictions never fire.
docker run --privileged with a bind mount pointed at a directory the sandbox already permits. Nothing trips.
Inside the privileged container, mount -t virtiofs virtiofs0 ~/mac-home. VirtioFS is Docker Desktop's own host-sharing mechanism, and it hands you the user's entire home directory with read and write.
The escape happens inside the Docker VM, which is outside the sandbox's jurisdiction. From the agent's point of view nothing illegal occurred. Every write is being made in the allowed workspace. SSH keys, credentials, browser data, all of it, one curl away.
First an attacker detects the architecture, pulls an Alpine minirootfs over curl, and docker import it as a local image so the registry restrictions never come into play, all done with a simple prompt injection:
Then the attacker runs the container detached and privileged, binding only a directory the sandbox already permits, and mount VirtioFS through a docker exec:
Finally, the attacker finds the user's home directory inside the mount, then write the payload into their .zshrc:
Now Calculator launches every time the user opens a terminal. Swap open -a Calculator for whatever you like; the point is that the process answering for it was never in the sandbox.
And none of it needs a direct prompt. The proofs of concept drive it with an explicit instruction because that's easiest to reproduce, but drop the same steps into a README, a code comment, or a doc the agent reads during normal work, and indirect injection does the rest.
You can view the full PoC video here:
Cursor's Seatbelt profile restricts Launch Services so the agent can't spawn processes outside the box. But with Auto-Run in Sandbox on (the default) and Auto-Run Network Access needs to be turned on however, the agent reaches Docker, and Docker doesn't care about Seatbelt.
Cursor triaged it, agreed the Docker socket should run from outside the sandbox, shipped a fix, and the issue stopped reproducing on the next build. Severity settled at high. You can read the advisory here: https://github.com/cursor/cursor/security/advisories/GHSA-v4xv-rqh3-w9mc.
Codex: not intended, but not a bug either
Codex carries the identical exposure. Launch it with the workspace-write sandbox and network access on:
Feed it the same plain-Docker script from above and it walks straight out. Because Codex auto-executes inside the allowed workspace, the approval policy never fires.
OpenAI's review agreed this was not intended behavior, that the agent can be driven into a path that breaks the workspace boundary and writes outside it. Then they narrowed it: configuration and environment dependent rather than a default path that hits every session. They paid the report, marked it as informational and declined a CVE.
Their reasoning is that in the current permissions model, they said, access to Unix sockets is mediated by the network proxy and denied by default unless a socket is explicitly allowed in the active profile.
Because talking to Docker requires that explicit allow setting, they don't consider it a default sandbox escape.
The shortest verdict of the three. Same chain, same privileged virtiofs0 mount, same result. The Gemini CLI team's position was that this isn't a bug as far as they're concerned, because the risks are spelled out in their documentation.
Sandboxes get judged on the syscalls they block. They should also get judged on the daemons they can reach. A privileged local daemon with a readable socket is a second, fully unsandboxed execution environment sitting one curl away, and Docker Desktop puts one on a huge fraction of developer machines by default. Deny-default was the right move, and it genuinely closed the Launch Services door. It just doesn't help when the agent can pick up a phone and ask a daemon to do the walking for it.
The overall long term solution would be to properly restrict autonomous agent operations with containment solutions that either monitor userland calls to the kernels like gVisor, real execution containers and finally humans in the loop. As we move forward with this new technology.
Subscribe and get the latest security updates
MAYBE YOU WILL FIND THIS INTERSTING AS WELL
The Week of Sandbox Escapes Day 4: Git directories do not have to be called .git
The Week of Sandbox Escapes Day 3: The sandbox let me edit a venv, and something else ran it
The Week of Sandbox Escapes Day 7: A time bomb in .vscode: bypassing Antigravity's Secure Mode
© 2026 Pillar. All rights reserved.
