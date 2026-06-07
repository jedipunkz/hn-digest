---
source: "https://yeri.be/giving-ai-ssh-access/"
hn_url: "https://news.ycombinator.com/item?id=48435545"
title: "Giving AI SSH Access"
article_title: "Giving AI SSH access – Yeri Tiete"
author: "yakkomajuri"
captured_at: "2026-06-07T15:34:24Z"
capture_tool: "hn-digest"
hn_id: 48435545
score: 3
comments: 0
posted_at: "2026-06-07T15:05:02Z"
tags:
  - hacker-news
  - translated
---

# Giving AI SSH Access

- HN: [48435545](https://news.ycombinator.com/item?id=48435545)
- Source: [yeri.be](https://yeri.be/giving-ai-ssh-access/)
- Score: 3
- Comments: 0
- Posted: 2026-06-07T15:05:02Z

## Translation

タイトル: AI SSH アクセスの付与
記事のタイトル: AI SSH アクセスの付与 – Yeri Tiete
説明: 共有 Docker ジャンプボックスを介した LLM エージェントの SSH アクセスの監査: ForceCommand、監査スクリプト、および BetterStack への OTLP。

記事本文:
イェリ・ティエテ
検索
政治
プライバシー
AI に SSH アクセスを与える
2026 年 6 月 7 日
· イェリ・ティエテ著
· ソフトウェア、Linux、ネットワーキング
私は、必要なユーザーはすべて root であることが多いため、最小限の監視で Claude (および以前は Gravity ) をフリートに SSH 接続できるようにしてきました。それは完全に問題ないと思いますが、そうでなくなるまでは… 😉
ここ数週間、私は AI の強化だけでなく、適切なツールへのアクセス (OrbStack、別の Mac で実行されている Claude Remote 経由) のセットアップも試みてきました。私は、AI が本番環境に移行する前にすべてを (ローカルで) テストできるようにしたいと考えています (たとえば、Cloudflare Workers やゼロトラストでは常にそうであるわけではありません) が、その動作を監査したいとも考えています。
クロード ユーザーの SSH キーをすべてのサーバーにゆっくりと追加していくうちに、傷つき、間違っていると感じ始めました。そこで、Box ジャンプボックスの目的を更新する時期が来ました。何年も前にセットアップした Docker シェル サーバーは、今でも毎日 (モッシュを介して) すべてのサーバーとルーターにアクセスするために使用しています。
エージェントは、私がまだ持っていないものを何も取得していませんでした (特定のデバイスに SSH キーを追加していなかったとき、時々ジャンプ ホストとしてボックスを使用するようにすでにエージェントに指示していたので、エージェントはジャンプ ホストの使用をすでに文書化していました)。
SSH の問題は、それが目に見えずに実行され、紙の痕跡が残らないことです。つまり、AI は驚くべきもので、ここ数週間、数か月ほど「永遠に未処理」だった仕事でこれほど大きな進歩を遂げることができたのは初めてですが、同時に、AF は恐ろしいものでもあります。信頼できない新入社員に核爆弾の暗号を与えるようなものだ。
私は簡単なものを望んでいました。エージェントが (SSH 経由で) 実行するすべてのコマンド: キャプチャされ、属性が付けられ、(Betterlogs に、えーっと、BetterStack のことです 😶‍🌫️ に) 送信され、検索可能です。そして、ボックス内の独自のユーザーに対してこれを行う必要がありました (自分のセッションをログに記録する必要はありません)。
コツは止めることです

ログインを使用します。エージェントに独自のユーザー alfons を与え、authorized_keys に独自のキーを設定しました。私自身のログインはまったく手つかずのままでした (私自身ではなく AI のみを監査したいのです)。
次に、 sshd は、そのユーザーのみを対象とした ForceCommand を取得します。
ユーザー アルフォンスと一致する
ForceCommand /usr/local/bin/ssh-audit
ssh-audit は、プロキシされたワンショット コマンド ( $SSH_ORIGINAL_COMMAND 、例: ssh Ocean systemctl status nginx ) とセッションのオープン/クローズをログに記録し、実際のシェルまたはコマンドを実行するため、ダウンストリームは何も変更されません。エージェントが実際の対話型シェルを開いている間、/etc/profile.d/box-audit.sh の bash DEBUG トラップが、入力された各コマンドを捕捉します。
すべては OpenTelemetry 経由で BetterLogs に送信されます。小さな box-log スクリプトは、jq を使用して OTLP/HTTP ペイロードを構築し、それを専用のソースに POST します。カールはバックグラウンドで行われ、3 秒のタイムアウトがあり、すべての失敗は無視されます。 Audit が SSH セッションをハングしたり強制終了したりすることは決して許可されません。 BetterLogs が動作を停止した場合、(いくつかの) ログ行が失われますが、AI は続行できます。
そしてもちろん、AI の SSH キーをすべてのデバイスから削除し、ジャンプ ホストを使用するように強制しました。
エージェント (モルゴス) --ssh alfons@box "ssh sea w"-->
sshd → 一致するユーザー alfons → ForceCommand /usr/local/bin/ssh-audit
→ ssh-audit: box-log proxied_command "ssh オーシャン w"
→ ssh-audit: exec bash -c "ssh オーシャン w"
→ box-log は BetterStack OTLP /v1/logs に対してバックグラウンドでカールを起動します
→ コマンドは以前と同様にダウンストリームで実行され、終了ステータスは影響を受けません。
Morgoth は、現在実行されている専用の Mac Mini m1 です。 Ocean はサーバーの例です。
各レコードには、 box.user=alfons 、リテラル コマンド、 ssh.connection (ソース IP/ポート)、および host.name が含まれています。最後の 1 つはコンテナーのホスト名から設定されているため、2 つのジャンプボックス ( box と boxnl ) が 1 つの区別できないストリームに崩壊することはありません。
eBPF は CAP_BPF を必要としています。

ジャンプボックス コンテナーにカーネル レベルの監査権限を与えること自体は、私が防御しようとしているものよりも悪い取引です。
コンテナーは --network host で実行されるため、コンテナー内の eBPF エージェントはホスト上のすべてのプロセス (隣接する約 20 個の無関係なコンテナーを含む) を確認します。
ここの AI は仲介されており信頼されており、積極的に私から逃れようとしていません。 ForceCommand はサーバー側で強制され、事後に execve syscall からインテントを再構築する代わりに、インテント ( ssh sea w ) を簡単にキャプチャします。
eBPF はフリート全体のランタイム セキュリティには適切なツールかもしれませんが、この使用例ではそうではありません。
ジャンプ ホストの使用方法は、プロンプトで説明されています。 SSH 構成では、ジャンプ ホスト (およびプルの場合は Gitlab) へのアクセスのみが許可されます。
すべてのコマンドの前に Gitlab 問題 ID を付けてログに記録する必要があるフィールドを追加することを検討しています (これにより、実行されるすべてのコマンド、たとえば問題 #123 を識別できるようになります)。
また、ブラックリストに登録されたコマンドのリスト (つまり、 rm -rf ) を追加する可能性があります。これがどのように機能するかはまだ完全にはわかりませんが、検討する価値はあります。

## Original Extract

Auditing an LLM agent's SSH access through a shared Docker jumpbox: ForceCommand, an audit script, and OTLP to BetterStack.

Yeri Tiete
Search
Politics
Privacy
Giving AI SSH access
7 June 2026
· by Yeri Tiete
· in software , linux , networking
I’ve been letting Claude (and previously Gravity ) SSH into my fleet with minimal oversight, as whatever user it needed, often root. Which I guess is perfectly fine, until it isn’t… 😉
Last couple of weeks I’ve been trying to set up hardening for AI, but also the proper tooling access (via OrbStack, Claude Remote running on a separate Mac). I want AI to be able to test everything (locally) before pushing to prod (not always a given with Cloudflare Workers and Zero Trust for example) but also want to audit what it does.
As I was slowly adding the SSH key of my Claude user to all my servers, I started to hurt and feel… wrong. So it was time to update the purpose of my Box jumpbox ; the Docker shell server I set up years ago and still use daily (over mosh) to reach all my servers and routers.
The agent wasn’t getting anything I didn’t already have (when I hadn’t added its SSH key to a particular device, I already told it to use box as a jump host at times, so it already had documented the use of a jump host).
The problem with SSHing, is that it was just doing it invisibly: leaving no papertrail. I mean, AI is amazing and I’ve never been able to make so much progress on stuff that was “forever in the backlog” as the last few weeks/months, but at the same time, it’s scary AF. Like giving an unreliable new employee the the codes to a nuclear bomb.
I wanted something easy: every command the agent runs (over SSH): captured, attributed, shipped (to Betterlogs , errrr, I mean BetterStack 😶‍🌫️), and searchable. And, it had to do this for its own user in box (no need to log my own sessions).
The trick is to stop sharing a login. I gave the agent its own user, alfons , with its own key in authorized_keys . My own logins stayed completely untouched (I only want to audit the AI, not myself).
Then sshd gets a ForceCommand , scoped to that user only:
Match User alfons
ForceCommand /usr/local/bin/ssh-audit
ssh-audit logs the proxied one-shot command ( $SSH_ORIGINAL_COMMAND , e.g. ssh ocean systemctl status nginx ), plus session open/close, then exec s the real shell or command so nothing downstream changes. For the times the agent opens an actual interactive shell, a bash DEBUG trap in /etc/profile.d/box-audit.sh catches each command as it’s typed.
Everything ships over OpenTelemetry to BetterLogs. A little box-log script builds an OTLP/HTTP payload with jq and POST s it to a dedicated source. The curl is backgrounded, has a 3 second timeout, and every failure is swallowed. Audit is never allowed to hang or kill an SSH session. If BetterLogs stops working, I lose a (few) log line(s), but the AI can proceed.
And I of course removed the AI’s SSH key from all the devices, forcing it to use the jump host.
agent (morgoth) --ssh alfons@box "ssh ocean w"-->
sshd → Match User alfons → ForceCommand /usr/local/bin/ssh-audit
→ ssh-audit: box-log proxied_command "ssh ocean w"
→ ssh-audit: exec bash -c "ssh ocean w"
→ box-log fires curl in background to BetterStack OTLP /v1/logs
→ command runs downstream as before, exit status unaffected
Morgoth is the dedicated Mac Mini m1 it now runs on. Ocean is an example server.
Each record carries box.user=alfons , the literal command, ssh.connection (source IP/port) and host.name — that last one set from the container’s hostname so my two jumpboxes ( box and boxnl ) don’t collapse into one indistinguishable stream.
eBPF wants CAP_BPF ; handing the jumpbox container kernel-level powers to audit itself is a worse trade than the thing I’m trying to defend against.
The container runs with --network host , so an in-container eBPF agent would see every process on the host (including the ~20 unrelated containers next to it).
The AI here is brokered and trusted(ish), not actively trying to evade me. ForceCommand is enforced server-side and trivially captures the intent ( ssh ocean w ) instead of me reconstructing it from execve syscalls after the fact.
eBPF may be the right tool for fleet-wide runtime security, but not in this use case.
The use of the jump host is doocumentend in the prompt . The SSH config only gives access to the jump host (and Gitlab, for pulls).
I’m considering adding a field where it needs to precede every command with a Gitlab issue ID and log that as well (so I can identify all commands it run for, for example, issue #123).
And potentially adding a list of blacklisted commands (ie rm -rf ); not entirely sure yet how this would work, but worth pondering about.
