---
source: "https://rogs.me/2026/02/claude-code-from-the-beach-my-remote-coding-setup-with-mosh-tmux-and-ntfy/"
hn_url: "https://news.ycombinator.com/item?id=48368079"
title: "Claude Code from the Beach"
article_title: "Claude Code from the beach: My remote coding setup with mosh, tmux and ntfy - rogs"
author: "emirb"
captured_at: "2026-06-03T00:49:50Z"
capture_tool: "hn-digest"
hn_id: 48368079
score: 2
comments: 2
posted_at: "2026-06-02T09:55:18Z"
tags:
  - hacker-news
  - translated
---

# Claude Code from the Beach

- HN: [48368079](https://news.ycombinator.com/item?id=48368079)
- Source: [rogs.me](https://rogs.me/2026/02/claude-code-from-the-beach-my-remote-coding-setup-with-mosh-tmux-and-ntfy/)
- Score: 2
- Comments: 2
- Posted: 2026-06-02T09:55:18Z

## Translation

タイトル: 浜辺のクロード・コード
記事のタイトル: Claude Code from the beach: mosh、tmux、ntfy を使用した私のリモート コーディング セットアップ - rogs
説明: 私のアパートから 2 ブロックの眺め 私は最近、携帯電話からクロード コードを実行することについての Granda によるこの素晴らしい投稿を読み、これは私の人生に必要だと思いました。アイデアはシンプルです。クロード コードのタスクを開始し、携帯電話をポケットに入れ、何か楽しいことをして、クロードが助けを必要としているときに通知を受け取ります。
[切り捨てられた]

記事本文:
ロッグ ホーム
について
AI
投稿
お問い合わせ
寄付する
Claude Code from the beach: mosh、tmux、ntfy を使用したリモート コーディング セットアップ
最近、Granda による、クロード コードの実行に関する素晴らしい投稿を読みました。
電話を見て、「これは人生に必要だ」と思いました。アイデアはシンプルです。
クロード コードのタスク、携帯電話をポケットに入れて、何か楽しいことをして、通知を受け取る
クロードがあなたの助けを必要とするとき、または仕事が終わったとき。どこからでも非同期開発。
しかし、私の設定は彼のものとは少し異なります。 Tailscale やクラウド VM は使用していません。
すでに WireGuard VPN を使用してデバイス、ホーム サーバー、および
自己ホスト型 ntfy インスタンス。そこで私は自分の好みに合わせた独自のバージョンを構築しました
インフラストラクチャ。
高レベルのアーキテクチャは次のとおりです。
┌─────┐ モッシュ ┌───────┐ ssh ┌────────┐
│ 電話 │──────────── │ 自宅サーバー │──────────── │ 職場の PC │
│ (Termux) │ WireGuard │ (ジャンプボックス) │ LAN │(クロードコード)│
━━━━┘ ━━━━━━┘ ━━━━┬──────┘
▲ │
│ ntfy (HTTPS) │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━┘
ループは次のとおりです。ビーチにいて、携帯電話で cc と入力し、tmux セッションに到達します。
クロード・コードと一緒に。タスクを与えて、携帯電話をポケットに入れて、やりたいことに戻ります。
やっていた。クロードが質問するか終了すると、私の携帯電話が鳴り響きます。引っ張ります
出て、応答して、再びポケットに入れます。開発は日々の隙間にフィットします。
非同期開発ループは次のようになります。

練習:
📱電話💻職場のPC🔔ntfy
│ │ │
│──── 'cc' と入力してください ────── ▶│ │
│────クロードに任務を与える───▶│ │
│ │ │
│ ┌───────┐ │ │
│ │ ポケットフォン │ │
│ ━━━━━━━┘ │ │
│ │ │
│ │──フックファイア───────▶│
│◀── 「クロードが意見を求めています」 ───────────────────
│ │ │
│────応答する─────────▶│ │
│ │ │
│ ┌───────┐ │ │
│ │ ポケットフォン │ │
│ ━━━━━━━┘ │ │
│ │ │
│ │──フックファイア───────▶│
│◀──「タスク完了」 ─────────────────────
│ │ │
│─── 審査、承認 PR ──▶│ │
│ │ │
ブログ投稿の設定をそのまま使用しないのはなぜでしょうか?
Granda のセットアップでは、VPN に Tailscale、Vultr クラウド VM、モバイルとして Termius を使用しています
ターミナル、および通知用の Poke。きれいで、機能します。しかし、私は持っていました
さまざまな制約:
すべてのデバイスを接続するサーバー上で wg-quick を実行している WireGuard VPN がすでにあります。必要ありません
テイルスケール用。
クラウド VM にお金を払いたくありませんでした。私の職場の PC は十分すぎるほど強力です
クロードコードを実行します。
通知用に ntfy を自己ホストしているため、Poke や外部のものが必要ありません。
通知サービス。
私は Termius ではなく、Termux (オープンソース) を使用しています。
この種のインフラストラクチャをまだ持っていない場合は、Granda のアプローチ

です
おそらくもっと単純です。しかし、あなたがすでに WireGuard を持っているような人であれば、
メッシュ サービスとセルフホスト サービスを使用する場合、このガイドはあなたのためのものです。
重要な洞察は、2 つの異なるタイプの回復力が必要であるということです: モッシュ
不安定なモバイル接続 (WiFi からセルラーへの移行、デッド ゾーン、
電話はスリープ状態）、tmux はセッションの永続性を処理します（アプリを閉じて再度開きます）
数時間後、すべてがまだそこにあります）。協力してモバイル開発を行う
実際に実行可能です。
なぜダブル SSH なのでしょうか?職場の PC を WireGuard ピアにしてみてはいかがでしょうか?
「すでに WireGuard ネットワークを持っているのに、なぜ追加すればいいのでしょう」と疑問に思われるかもしれません。
職場の PC をピアとして使用し、携帯電話から直接モッシュしますか?
簡単に言うと、それは私の雇用主の機械です。監視ソフトが入ってます
インストール済み: 画面取得、エンドポイント ポリシー、機能。ワイヤーガードのインストール
これは、私の個人的なネットワークを介してトラフィックをトンネリングする VPN クライアントを実行することを意味します。
インフラストラクチャは、IT セキュリティに問題を引き起こすようなものです。私は
その会話には関わりたくない。
一方、SSH は標準の開発ツールです。 Linux 上の openssh サーバー
マシンはこれ以上に目立たないほどです。
その代わりに、私のホームサーバーはジャンプボックスとして機能します。電話が家につながります
WireGuard 経由のサーバー (すべて個人のインフラストラクチャであり、雇用主は必要ありません)
関与)、その後、ホーム サーバーがローカル経由で職場の PC に SSH 接続します。
ネットワーク。職場の PC に必要なのは SSH サーバーのみで、VPN クライアントや奇妙なトンネルは必要ありません。
監視ソフトウェアを点滅させるようなものは何もありません。
┌─────────────────────┐
│ 私のインフラストラクチャー │
│ │
│ ┌─────┐ ワイヤーガード ┌─────

────┐ │
│ │ 電話 │◀──────────▶│ WG サーバー │ │
│ │ (ピア) │ トンネル │ │ │
│ ━━━━┬─────┘ ━━━━┬─────┘ │
│ │ │ │
│ │ モッシュワイヤーガード │ │
│ │ (トンネルを通る) トンネル │ │
│ │ │ │
│ ▼ ▼ │
│ ┌─────────┐ │
│ │ ホームサーバー │◀─────────────────────
│ │ (ピア) │ │
│ ━━━━┬───────┘ │
│ │ │
━─────┼────────────────┘
│
│ssh(LAN)
│
┌─────┼────────────────┐
│ ▼ │
│ ┌───────┐ │
│ │ 仕事用パソコン │ │
│ │ (SSH のみ) │ 雇用主のインフラストラクチャ │
│ └───────┘ │
━━━━━━━━━━━━━━━━━┘
おまけに、これは仕事用 PC が公共のインターネットにさらされることがないことを意味します。それ
ローカル ネットワーク上のマシンからの SSH のみを受け入れます。多層防御。
フェーズ 1: 職場 PC 上の SSH サーバー
私の職場の PC は Ubuntu 24.04 を実行しています。最初に行うこと: SSH をインストールして強化する
サーバー。
sudo apt update && sudo apt install -y openssh-server
sudo systemctl で ssh を有効にする
注: Ubuntu 24.04 では、サービスは sshd ではなく ssh と呼ばれます。これにはつまずいた
アップ。
次に、構成を強化します。 /etc/ssh/sshd_config を作成しました

と:
PermitRootログイン番号
パスワード認証番号
KbdInteractiveAuthentication いいえ
公開鍵認証はい
エージェント転送を許可しない
X11転送なし
PAM を使用する はい
MaxAuthTries 3
クライアントアライブ間隔 60
ClientAliveCountMax 3
キーのみの認証、root ログイン、パスワード認証はありません。機械だけなので、
ローカル ネットワーク経由でアクセスできるので、十分に安全です。
自宅サーバー→職場PC接続用のSSHキーの設定
ホーム サーバーで、キー ペアをまだ持っていない場合は生成します。
ssh-keygen -t ed25519 -C "ホームサーバー->ワークPC"
デフォルトのパス ( /.ssh/id_ed25519 ) を受け入れます。次に、公開キーを
職場のPC:
ssh-copy-id roger@<work-pc-ip>
sshd を再起動します。
sudo systemctl sshを再起動します
重要: ホームサーバーを閉じる前に、ホームサーバーからの SSH 接続をテストします。
現在のセッション。自分を締め出さないでください。
# 自宅サーバーから
ssh roger@<work-pc-ip>
パスワードを要求されずにシェルに入れられたら、あなたは完璧です。
WireGuard をセットアップしていない場合は、Tailscale を使用するのが最も簡単な方法です。
プライベートネットワークが稼働中です。携帯電話と職場の PC にインストールすると、
直接お互いに会います。ジャンプ ホスト、ポート転送、ファイアウォールは不要です。
ルール。正直、この種のことには魔法です。使わない唯一の理由
これは、Tailscale が存在する前にすでに WireGuard を実行していたためです。
ここでの考え方は単純です。職場の PC に SSH 接続するたびに、着陸したいと考えています。
tmux セッションで直接。セッションがすでに存在する場合は、それに接続します。そうでない場合は、
作成します。
# マウスのサポート (携帯電話でサム操作する場合に必須)
-g マウスをオンに設定します
# ウィンドウの番号付けを 1 から開始します (電話のキーボードでアクセスしやすくなります)
set -g ベースインデックス 1
setw -g ペインベースインデックス 1
# ステータスバー
set -g ステータススタイル 'bg=colour235 fg=colour136'
set -g status-left '#[fg=colour46][#S] '
セット -g 統計

us-right '#[fg=colour166]%H:%M'
set -g ステータス左の長さ 30
# より長いスクロールバック
set -g 履歴制限 50000
# エスケープ遅延を軽減します (SSH 経由でのエディターの動作が速くなります)
set -sg エスケープタイム 10
# セッションを維持する
set -g destroy-unattached オフ
携帯電話を使用する場合、マウスのサポートは不可欠です。タップできること
ペインを選択し、指でスクロールし、サイズを変更すると、
違い。
次に、作業用 PC の ~/.config/fish/config.fish で次のようにします。
-q SSH_CONNECTION を設定した場合; -q TMUX を設定しない
tmuxattach -t クロード 2 > /dev/null;または tmux new -s claude -c ~/projects/my-app
終わり
これにより SSH_CONNECTION がチェックされるため、リモート接続している場合にのみ自動接続されます。
物理的にマシンの前にいるときは、tmux を使用せずに通常どおりターミナルを使用します。
この区別は、後の通知において重要になります。
フェーズ 3: クロード コード フック + ntfy
ここが楽しい部分です。 Claude Code にはコマンドを実行できるフック システムがあります
特定の出来事が起こったとき。次の 3 つのイベントにフックします。
AskUserQuestion : クロードには私の意見が必要です。優先度の高い通知。
停止 : クロードはタスクを終了しました。通常の優先度。
エラー : 何かが壊れました。優先度が高い。
まず、通知を送信するスクリプトです。私が作成しました
~/.claude/hooks/notify.sh :
#!/usr/bin/env bash
# SSH で開始された tmux セッションにいる場合のみ通知する
もし！ tmux show-environment SSH_CONNECTION 2>/dev/null | grep -q SSH_CONNECTION = ;それから
0番出口
フィ
EVENT_TYPE = " ${ 1 :- 不明 } "
NTFY_URL = "https://ntfy.example.com/claude-code"
NTFY_TOKEN = "tk_your_token_here"
EVENT_DATA = $(猫)
ケース「 $EVENT_TYPE 」
質問）
TITLE = "🤔 クロードには意見が必要です"
優先度 = 「高」
MESSAGE = $( echo " $EVENT_DATA " | jq -r '.tool_input.question // .tool_input.questions[0].question // "クロードに質問があります"' 2>/dev/null )
;;
やめてください）
タイトル =

「✅クロード終わった」
優先度 = "デフォルト"
MESSAGE = "タスクが完了しました"
;;
エラー)
TITLE = "❌ クロードがエラーに遭遇しました"
優先度 = 「高」
MESSAGE = $( echo " $EVENT_DATA " | jq -r '.error // "問題が発生しました"' 2>/dev/null )
;;
*)
TITLE = 「クロード・コード」
優先度 = "デフォルト"
MESSAGE = "イベント: $EVENT_TYPE "
;;
イーサック
PROJECT = $( ベース名 " $PWD " )
カール -s \
-H "認可: ベアラー $NTFY_TOKEN " \
-H "タイトル: $TITLE " \
-H "優先度: $PRIORITY " \
-H "タグ: コンピューター" \
-d "[ $PROJECT ] $MESSAGE " \
" $NTFY_URL " > /dev/null 2>& 1
chmod +x ~/.claude/hooks/notify.sh
上部の SSH_CONNECTION チェックは非常に重要です。これにより、次のような通知が送信されなくなります。
私が機械の前に座っているときに発砲します。 SSH接続するときにのみtmuxを使用するため、
リモートでは、tmux 環境に SSH_CONNECTION が設定されるのは、私がリモートにいる場合のみです。
リモート。巧妙なトリック。
次に、 ~/.claude/settings.json で次のようにします。
{
「フック」: {
"PreToolUse" : [
{
"matcher" : "AskUserQuestion" ,
「フック」: [
{
"タイプ" : "コマンド" ,
"command" : "~/.claude/hooks/notify.sh 質問"
}
]
}
]、
「停止」: [
{
「フック」: [
{
"タイプ" : "コマンド" ,
"コマンド" : "~/.claude/hooks/notify.sh 停止"
}
]
}
]
}
}
これはグローバル設定ファイルです。プロジェクトにも
.claude/settings.json 、それらはマージされます。競合はありません。
私は ntfy を自己ホストしているので、トピックとアクセス トークンを作成しました。
# ntfy サーバー/コンテナ内
ntfy トークンの追加 --expires = 30d あなたのユーザー名
ntfy アクセスあなたのユーザー名 claude-code rw
ntfy アクセス全員クロードコード拒否
ntfy トピックはオンデマンドで作成されるため、トピックをサブスクライブするだけで作成されます。で
Android ntfy アプリ、私はそれを自分の自己ホスト型インスタンスに指定し、
クロードコードのトピック。
以下を使用して全体の動作をテストできます。
echo '{"tool_input":{"question":"これをリファクタリングする必要がありますか?"}}' | ~/.claude/hooks/notify.sh の質問
エコー '{}' | ~/.c

laude/hooks/notify.sh 停止
echo '{"error":"ModuleNotFoundError: foo という名前のモジュールがありません"}' | ~/.claude/hooks/notify.sh エラー
3 つの通知、3 つの異なる優先順位。とても満足です。
代替通知システム
ntfy を自己ホストしたくない場合は、次のオプションがあります。
ntfy.sh : ntfy のパブリック インスタンス。無料、セットアップ不要、選択するだけ
ランダムっぽいトピック名。欠点は、あなたのトピック名を知っている人は誰でも
通知を送信できます。
プッシュオーバー : プラットフォームごとに 5 ドルの 1 回限りの購入。非常に信頼性が高く、優れた API。の
通知スクリプトはほぼ同じですが、curl 呼び出しが異なるだけです。
Gotify : ntfy と同様に自己ホスト型ですが、HTTP の代わりに WebSocket を使用します。良かったら
あなたはすでにそれを実行しています。
Telegram Bot API : 無料でセットアップが簡単。 BotFather でボットを作成し、取得します
チャット ID を入力し、sendMessage エンドポイントをカールします。
Poke : グランダが投稿で使用しているもの。シンプルな Webhook からプッシュへのサービス。
Termux は、Android スマートフォン上のターミナル エミュレータです。設定方法は次のとおりです。
pkg update && pkg install -y mosh openssh Fish
電話機に SSH で接続します (セットアップを簡単にするため)
これらすべてを電話のキーボードで設定するのは面倒です。 Termux で sshd をセットアップしました
PCから設定できるようにしました。
~/.config/fish/config.fish 内:
sshd 2 > /dev/null
これにより、Termux を開くたびに sshd が起動します。すでに実行されている場合は、
静かに失敗します。 Termux はデフォルトでポート 8022 で sshd を実行します。
まずはパスを設定します

[切り捨てられた]

## Original Extract

The view two blocks from my apartment I recently read this awesome post by Granda about running Claude Code from a phone, and I thought: I need this in my life. The idea is simple: kick off a Claude Code task, pocket the phone, go do something fun, and get a notification when Claude needs your help
[truncated]

rogs Home
About
AI
Posts
Contact
Donate
Claude Code from the beach: My remote coding setup with mosh, tmux and ntfy
I recently read this awesome post by Granda about running Claude Code from a
phone, and I thought: I need this in my life . The idea is simple: kick off a
Claude Code task, pocket the phone, go do something fun, and get a notification
when Claude needs your help or finishes working. Async development from anywhere.
But my setup is a bit different from his. I’m not using Tailscale or a cloud VM.
I already have a WireGuard VPN connecting my devices, a home server, and a
self-hosted ntfy instance. So I built my own version, tailored to my
infrastructure.
Here’s the high-level architecture:
┌──────────┐ mosh ┌─────────────┐ ssh ┌─────────────┐
│ Phone │───────────────▶ │ Home Server │───────────────▶ │ Work PC │
│ (Termux) │ WireGuard │ (Jump Box) │ LAN │(Claude Code)│
└──────────┘ └─────────────┘ └──────┬──────┘
▲ │
│ ntfy (HTTPS) │
└─────────────────────────────────────────────────────────────┘
The loop is: I’m at the beach, I type cc on my phone, I land in a tmux session
with Claude Code. I give it a task, pocket the phone, and go back to whatever I
was doing. When Claude has a question or finishes, my phone buzzes. I pull it
out, respond, pocket it again. Development fits into the gaps of the day.
And here’s what the async development loop looks like in practice:
📱 Phone 💻 Work PC 🔔 ntfy
│ │ │
│──── type 'cc' ────────────▶│ │
│──── give Claude a task ───▶│ │
│ │ │
│ ┌─────────────────┐ │ │
│ │ pocket phone │ │ │
│ └─────────────────┘ │ │
│ │ │
│ │── hook fires ────────────▶│
│◀── "Claude needs input" ───────────────────────────────│
│ │ │
│──── respond ──────────────▶│ │
│ │ │
│ ┌─────────────────┐ │ │
│ │ pocket phone │ │ │
│ └─────────────────┘ │ │
│ │ │
│ │── hook fires ────────────▶│
│◀── "Task complete" ────────────────────────────────────│
│ │ │
│──── review, approve PR ───▶│ │
│ │ │
Why not just use the blog post’s setup?
Granda’s setup uses Tailscale for VPN, a Vultr cloud VM, Termius as the mobile
terminal, and Poke for notifications. It’s clean and it works. But I had
different constraints:
I already have a WireGuard VPN running wg-quick on a server that connects all my devices. No need
for Tailscale.
I didn’t want to pay for a cloud VM. My work PC is more than powerful enough to
run Claude Code.
I self-host ntfy for notifications, so no need for Poke or any external
notification service.
I use Termux (open-source), not Termius.
If you don’t have this kind of infrastructure already, Granda’s approach is
probably simpler. But if you’re the kind of person who already has a WireGuard
mesh and self-hosted services, this guide is for you.
The key insight is that you need two different types of resilience : mosh
handles the flaky mobile connection (WiFi to cellular transitions, dead zones,
phone sleeping), while tmux handles session persistence (close the app, reopen
hours later, everything’s still there). Together they make mobile development
actually viable.
Why the double SSH? Why not make the work PC a WireGuard peer?
You might be wondering: if I already have a WireGuard network, why not just add
the work PC as a peer and mosh straight into it from my phone?
The short answer: it’s my employer’s machine . It has monitoring software
installed: screen grabbing, endpoint policies, the works. Installing WireGuard
on it would mean running a VPN client that tunnels traffic through my personal
infrastructure, which is the kind of thing that raises flags with IT security. I
don’t want to deal with that conversation.
SSH, on the other hand, is standard dev tooling. An openssh-server on a Linux
machine is about as unremarkable as it gets.
So instead, my home server acts as a jump box. My phone connects to the home
server over WireGuard (that’s all personal infrastructure, no employer
involvement), and then the home server SSHs into the work PC over the local
network. The work PC only needs an SSH server, no VPN client, no weird tunnels,
nothing that would make the monitoring software blink.
┌──────────────────────────────────────────────────┐
│ My Infrastructure │
│ │
│ ┌───────────┐ WireGuard ┌──────────────┐ │
│ │ Phone │◀──────────────▶│ WG Server │ │
│ │ (peer) │ tunnel │ │ │
│ └─────┬─────┘ └──────┬───────┘ │
│ │ │ │
│ │ mosh WireGuard │ │
│ │ (through tunnel) tunnel │ │
│ │ │ │
│ ▼ ▼ │
│ ┌──────────────┐ │
│ │ Home Server │◀───────────────────────────────│
│ │ (peer) │ │
│ └──────┬───────┘ │
│ │ │
└─────────┼────────────────────────────────────────┘
│
│ ssh (LAN)
│
┌─────────┼────────────────────────────────────────┐
│ ▼ │
│ ┌────────────┐ │
│ │ Work PC │ │
│ │ (SSH only) │ Employer Infrastructure │
│ └────────────┘ │
└──────────────────────────────────────────────────┘
As a bonus, this means the work PC has zero exposure to the public internet. It
only accepts SSH from machines on my local network. Defense in depth.
Phase 1: SSH server on the work PC
My work PC is running Ubuntu 24.04. First thing: install and harden the SSH
server.
sudo apt update && sudo apt install -y openssh-server
sudo systemctl enable ssh
Note: on Ubuntu 24.04 the service is called ssh , not sshd . This tripped me
up.
Then harden the config. I created /etc/ssh/sshd_config with:
PermitRootLogin no
PasswordAuthentication no
KbdInteractiveAuthentication no
PubkeyAuthentication yes
AllowAgentForwarding no
X11Forwarding no
UsePAM yes
MaxAuthTries 3
ClientAliveInterval 60
ClientAliveCountMax 3
Key-only auth, no root login, no password auth. Since the machine is only
accessible through my local network, this is plenty secure.
Setting up SSH keys for the home server → work PC connection
On the home server , generate a key pair if you don’t already have one:
ssh-keygen -t ed25519 -C "homeserver->workpc"
Accept the default path ( /.ssh/id_ed25519 ). Then copy the public key to the
work PC:
ssh-copy-id roger@<work-pc-ip>
Now restart sshd:
sudo systemctl restart ssh
Important : Test the SSH connection from your home server before closing your
current session. Don’t lock yourself out.
# From the home server
ssh roger@<work-pc-ip>
If it drops you into a shell without asking for a password, you’re golden.
If you don’t have a WireGuard setup, Tailscale is the easiest way to get a
private network going. Install it on your phone and your work PC, and they can
see each other directly. No jump host needed, no port forwarding, no firewall
rules. It’s honestly magic for this kind of thing. The only reason I don’t use it
is because I already had WireGuard running before Tailscale existed.
The idea here is simple: every time I SSH into the work PC, I want to land
directly in a tmux session. If the session already exists, attach to it. If not,
create one.
# mouse support (essential for thumbing it on the phone)
set -g mouse on
# start window numbering at 1 (easier to reach on phone keyboard)
set -g base-index 1
setw -g pane-base-index 1
# status bar
set -g status-style 'bg=colour235 fg=colour136'
set -g status-left '#[fg=colour46][#S] '
set -g status-right '#[fg=colour166]%H:%M'
set -g status-left-length 30
# longer scrollback
set -g history-limit 50000
# reduce escape delay (makes editors snappier over SSH)
set -sg escape-time 10
# keep sessions alive
set -g destroy-unattached off
Mouse support is essential when you’re using your phone. Being able to tap to
select panes, scroll with your finger, and resize things makes a massive
difference.
Then in ~/.config/fish/config.fish on the work PC:
if set -q SSH_CONNECTION; and not set -q TMUX
tmux attach -t claude 2 > /dev/null; or tmux new -s claude -c ~/projects/my-app
end
This checks for SSH_CONNECTION so it only auto-attaches when I’m remoting in.
When I’m physically at the machine, I use the terminal normally without tmux.
This distinction becomes important later for notifications.
Phase 3: Claude Code hooks + ntfy
This is the fun part. Claude Code has a hook system that lets you run commands
when certain events happen. We’re going to hook into three events:
AskUserQuestion : Claude needs my input. High priority notification.
Stop : Claude finished the task. Normal priority.
Error : Something broke. High priority.
First, the script that sends notifications. I created
~/.claude/hooks/notify.sh :
#!/usr/bin/env bash
# Only notify if we're in an SSH-originated tmux session
if ! tmux show-environment SSH_CONNECTION 2>/dev/null | grep -q SSH_CONNECTION = ; then
exit 0
fi
EVENT_TYPE = " ${ 1 :- unknown } "
NTFY_URL = "https://ntfy.example.com/claude-code"
NTFY_TOKEN = "tk_your_token_here"
EVENT_DATA = $( cat )
case " $EVENT_TYPE " in
question )
TITLE = "🤔 Claude needs input"
PRIORITY = "high"
MESSAGE = $( echo " $EVENT_DATA " | jq -r '.tool_input.question // .tool_input.questions[0].question // "Claude has a question for you"' 2>/dev/null )
;;
stop )
TITLE = "✅ Claude finished"
PRIORITY = "default"
MESSAGE = "Task complete"
;;
error )
TITLE = "❌ Claude hit an error"
PRIORITY = "high"
MESSAGE = $( echo " $EVENT_DATA " | jq -r '.error // "Something went wrong"' 2>/dev/null )
;;
* )
TITLE = "Claude Code"
PRIORITY = "default"
MESSAGE = "Event: $EVENT_TYPE "
;;
esac
PROJECT = $( basename " $PWD " )
curl -s \
-H "Authorization: Bearer $NTFY_TOKEN " \
-H "Title: $TITLE " \
-H "Priority: $PRIORITY " \
-H "Tags: computer" \
-d "[ $PROJECT ] $MESSAGE " \
" $NTFY_URL " > /dev/null 2>& 1
chmod +x ~/.claude/hooks/notify.sh
The SSH_CONNECTION check at the top is crucial: it prevents notifications from
firing when I’m sitting at the machine. Since I only use tmux when SSHing in
remotely, the tmux environment will only have SSH_CONNECTION set when I’m
remote. Neat trick.
Then in ~/.claude/settings.json :
{
"hooks" : {
"PreToolUse" : [
{
"matcher" : "AskUserQuestion" ,
"hooks" : [
{
"type" : "command" ,
"command" : "~/.claude/hooks/notify.sh question"
}
]
}
],
"Stop" : [
{
"hooks" : [
{
"type" : "command" ,
"command" : "~/.claude/hooks/notify.sh stop"
}
]
}
]
}
}
This is the global settings file. If your project also has a
.claude/settings.json , they’ll be merged. No conflicts.
I’m self-hosting ntfy, so I created a topic and an access token:
# Inside your ntfy server/container
ntfy token add --expires = 30d your-username
ntfy access your-username claude-code rw
ntfy access everyone claude-code deny
ntfy topics are created on demand, so just subscribing to one creates it. On the
Android ntfy app, I pointed it at my self-hosted instance and subscribed to the
claude-code topic.
You can test the whole thing works with:
echo '{"tool_input":{"question":"Should I refactor this?"}}' | ~/.claude/hooks/notify.sh question
echo '{}' | ~/.claude/hooks/notify.sh stop
echo '{"error":"ModuleNotFoundError: No module named foo"}' | ~/.claude/hooks/notify.sh error
Three notifications, three different priorities. Very satisfying.
Alternative notification systems
If you don’t want to self-host ntfy, here are some options:
ntfy.sh : The public instance of ntfy. Free, no setup, just pick a
random-ish topic name. The downside is that anyone who knows your topic name
can send you notifications.
Pushover : $5 one-time purchase per platform. Very reliable, nice API. The
notification script would be almost identical, just a different curl call.
Gotify : Self-hosted like ntfy, but uses WebSockets instead of HTTP. Good if
you’re already running it.
Telegram Bot API : Free, easy to set up. Create a bot with BotFather, get
your chat ID, and curl the sendMessage endpoint.
Poke : What Granda uses in his post. Simple webhook-to-push service.
Termux is the terminal emulator on my Android phone. Here’s how I set it up.
pkg update && pkg install -y mosh openssh fish
SSH into your phone (for easier setup)
Configuring all of this on a phone keyboard is painful. I set up sshd on Termux
so I could configure it from my PC.
In ~/.config/fish/config.fish :
sshd 2 > /dev/null
This starts sshd every time you open Termux. If it’s already running, it
silently fails. Termux runs sshd on port 8022 by default.
First, set a pass

[truncated]
