---
source: "https://ykdojo.github.io/claude-controls-mac/"
hn_url: "https://news.ycombinator.com/item?id=48959392"
title: "Setting up your spare Mac for Claude Code to control, a step-by-step guide"
article_title: "How to set up your spare Mac for Claude Code to fully control - a step-by-step guide | claude-controls-mac"
author: "ykev"
captured_at: "2026-07-18T16:44:27Z"
capture_tool: "hn-digest"
hn_id: 48959392
score: 21
comments: 3
posted_at: "2026-07-18T16:12:08Z"
tags:
  - hacker-news
  - translated
---

# Setting up your spare Mac for Claude Code to control, a step-by-step guide

- HN: [48959392](https://news.ycombinator.com/item?id=48959392)
- Source: [ykdojo.github.io](https://ykdojo.github.io/claude-controls-mac/)
- Score: 21
- Comments: 3
- Posted: 2026-07-18T16:12:08Z

## Translation

タイトル: クロード コードを制御できるように予備の Mac をセットアップする、ステップバイステップ ガイド
記事のタイトル: クロード コードを完全に制御できるように予備の Mac をセットアップする方法 - ステップバイステップ ガイド |クロード コントロール マック
説明: 予備の Mac を常時稼働マシンに変えるためのステップバイステップのガイド。Claude コードが完全に制御できます。携帯電話から Claude アプリを介して、またはメインの Mac から SSH 経由で話しかけることができます。

記事本文:
クロード コードを完全に制御できるように予備の Mac をセットアップする方法 - ステップバイステップ ガイド
予備の Mac を常時オンにする方法についての完全なステップバイステップ ガイドは次のとおりです。
コンピューターの使用が有効になっているため、マシン クロード コードは完全に制御できます。携帯電話から通話できるようになります
Claude アプリ経由、またはメイン Mac から SSH 経由で。
クロード コードが独自に制御できる別の環境を作成したかったので、
必ずしも自分のマシンで実行したくないタスクを委任する - 特定の種類のタスク
研究タスクと開発タスク。
クロード コード、特に --dangerously-skip-permissions フラグがオンの場合、
メイン マシンで実行すると、固有のリスクが伴います。これらを排除/軽減することができます
予備の Mac に必要なものがすべて揃った別の環境を作成することによるリスク
にアクセスできるようにします。
いつでもどこでもクロード コードと会話できるという追加の特典もあります。
あなたの電話。私は個人的には、誰かと話すことを好むことが多いので、非常に便利だと感じています
モバイル アプリでは通常のクロードの代わりにクロード コード - クロード コードの方が多くの場合、
有能な。
次のガイドは、メインの Mac と設定できる予備の Mac があることを前提としています。
しかし、そこからインスピレーションを得て、あらゆることに応用できるはずです。
2台のマシンの組み合わせ。
まず、いくつかの質問に簡単に答えてみましょう。
なぜコンテナ内で実行しないのでしょうか?
私はコンテナ内で実行することを強く支持しており、コンテナを構築したこともあります
それを便利に行うための環境全体。
ただし、いくつかの制限があることがわかりました。
まず、メイン マシン上で引き続き実行されるため、完全に分離されていません。のために
たとえば、送信されるネットワーク リクエストは依然としてメイン マシンを経由します。
次に、コンテナの機能には制限があります。たとえば、私が欲しかったのは、
私のエージェントからBへ

ゲーム開発のために Unity を実行できますが、簡単な方法はありません
コンテナに入ったもの。 Mac でのみ利用できる他のアプリについても同様です。
アクセスできなくなります。これは、Claude Code に次のことを実行させたい場合に特に重要です。
これらのアプリは、クリックやドラッグなどのコンピューターの使用を通じて制御します。
OpenClaw のようなものを使用してみてはいかがでしょうか?
個人的には、Claude Code の最新機能をすべて利用できるのが気に入っています。私も
クロードアプリからコントロールできるなど、本当に便利だと思いました。
そして、Claude サブスクリプションを持っている場合は、それを使用できるようになります。
追加のボーナス。
結局のところ、幅広い権限を持つエージェントをマシン上で実行する方が安全です
失うものは何もありませんが、完全な Mac を使用できるという利点があります。
コンテナの代わりに。ここでのアプローチ:
メインの Mac ではなく、古い/予備の Mac を使用します。
個人データや Apple ID を使用せずに新しいローカル アカウントを作成します。
エージェントには連絡できる機密性の高いものは何もありません。
ローカル ネットワーク上のメイン Mac から SSH 経由で駆動し、から制御します。
あなたの電話。
いつも使っている Mac (ソース) を同じ Wi-Fi 上で。
1. ターゲット Mac で新たに開始します。
最初に消去します (個人データが含まれている場合)
エージェントにこのマシンへの完全なアクセス権を与えることになり、保存されているものすべてにアクセスできるようになります
その上で。アクセスしたくない既存のデータがある場合は、マシンを消去します
最初に:
それをサポートする Mac: システム設定 -> 一般 -> 転送またはリセット ->
すべてのコンテンツと設定を消去します。
古い Intel Mac: 再起動してリカバリを開始します (Cmd-R を押したままにします)
boot)、ディスクユーティリティを使用して内部ドライブを消去してから、macOS を再インストールします。
必要に応じて、後で最新の macOS にアップデートします (システム設定 -> 一般 ->
ソフトウェアアップデート）。
新しく独立したアカウントを作成する
新しい場所を作成する

l ユーザーアカウント (システム設定 -> ユーザーとグループ)。
Apple ID にサインインしないことをお勧めします。セットアップ中はスキップしてください。
アカウントを管理者にします (まだ管理者にしていない場合)
アカウントには管理者権限が必要です。そうでない場合、sudo は実行を拒否します。
システム設定 -> ユーザーとグループ -> アカウントをこのユーザーに許可するように設定します。
このコンピュータを管理します。
別の管理者アカウントから修復する必要がある場合は、次のようにします。
sudo dseditgroup -o edit -a <ユーザー> -t ユーザー管理者
2. ターゲット Mac でリモート ログイン (SSH) を有効にします。
ターゲットで SSH をオンにして、ソース Mac が接続できるようにします。
sudo systemsetup -setremotelogin on
コマンドが「リモート ログインをオンまたはオフにするにはフル ディスク アクセスが必要です」で失敗した場合
権限を付与するには、まずターミナル アプリにフル ディスク アクセスを付与します。
システム設定 -> プライバシーとセキュリティ -> フルディスクアクセス 。
+ をクリックし、ファイルピッカーで [アプリケーション] -> [ユーティリティ] -> [ターミナル] に移動します。
そしてそれを追加します。
ターミナルを終了して再度開き、コマンドを再実行します。
3. ターゲットアカウントのパスワードなしの sudo
これは、エージェント (および SSH コマンド) がパスワードなしで管理タスクを実行できるようにするためです。
毎回プロンプトを表示します。これをターゲット上で 1 回実行します。それ
今度はログイン パスワードを要求します。
echo "<ユーザー> ALL=(ALL) NOPASSWD: ALL" | sudo tee /etc/sudoers.d/<user>-nopasswd > /dev/null
sudo chmod 440 /etc/sudoers.d/<user>-nopasswd
sudo visudo -cf /etc/sudoers.d/<user>-nopasswd # validate - 'parsed OK' を出力する必要があります
これにより、<user> がコマンドなしで sudo を実行できることを Mac に伝える小さなルール ファイルが作成されます。
パスワードプロンプト:
1 行目では、ルールを /etc/sudoers.d/ に書き込みます。
行 2 はファイルを読み取り専用にします。それ以外の場合、sudo はファイルを無視します。
3 行目では構文を検証します。 sudoers ファイルのタイプミスにより、sudo からロックされる可能性があります
したがって、 parsed OK を出力する必要があります。
この後、プロンプトなしで sudo が実行されます。 sudo -n true でテストすると成功します。
パスワードなしの場合はサイレントに sudo

動作します。
4. ターゲットのアドレス (ホスト名または IP) を見つけます。
ホスト名または IP を使用してターゲットに到達できます。を使用することをお勧めします
ホスト名: IP は変更される可能性がありますが、同じままです。
ホスト名 (推奨)。ターゲット上で実行します。
scutil --get LocalHostName # ホスト名を出力します。 MacBook Pro
.local を追加して、アドレス <target-host>.local を形成します。から読むこともできます
[システム設定] -> [一般] -> [共有]、 [ローカル ホスト名] として表示されます。
ターゲットに一意の名前を付けます。各 Mac には一意の .local 名が必要です。
あなたのネットワーク。 2 台のマシンが名前を共有している場合、アドレスが間違った Mac を指している可能性があります。
ターゲットの名前が一意であることを確認してください。必要に応じて名前を変更します。
sudo scutil --set LocalHostName newmacbook # -> newmacbook.local
IP アドレス (推奨されません)。ターゲット上で実行します。
ipconfig getifaddr en0 # 例: 192.168.1.80
IP は、再起動後または一定の時間が経過すると変更される可能性があることに注意してください。
このガイドの残りの部分では、<user> をターゲット アカウント名に置き換えます。
<target-host> は上記のホスト名に置き換えられるため、アドレスは次のようになります。
<ユーザー>@<ターゲットホスト>.local 。代わりに IP を使用することもできます。
<ターゲットホスト>.local 。
5. ソース Mac からパスワードなしの SSH をセットアップする
ソース Mac で、SSH キーを作成します (すでに持っている場合はスキップしてください)。
ssh-keygen -t ed25519
公開キーをターゲットにインストールします。これにより、ターゲットアカウントのログインが求められます
パスワードを 1 回:
ssh-copy-id <ユーザー>@<ターゲットホスト>.local
テストしてください。これにより、パスワード プロンプトなしでターゲット ユーザー名が出力されます。
ssh <ユーザー>@<ターゲットホスト>.local whoami
6. ターゲットを覚醒させておく
デフォルトでは、macOS は電源に接続されている場合でもアイドル状態で約 10 分後にスリープします。これには時間がかかります。
それをネットワークから外します。スリープ状態にしないようにするには、これをターゲット上で (または SSH 経由で) 実行します。
出典):
sudo pmset -c sleep 0 # 接続中はシステムスリープしません (-c =

充電器上）
sudo pmset -c disablesleep 1 # 蓋を閉じた状態でのスリープも防止します (クラムシェル)
sudo pmset -c Displaysleep 0 # ディスプレイもオンのままにする
確認:
pmset -g | grep -iE 'スリープ'
sleep 0 、 SleepDisabled 1 、および出力に sleep 0 が表示されると、機能したことが確認されます。
マシンがバッテリーで動作する場合がある場合は、-c の代わりに -a を使用してすべてに適用します。
電源（バッテリーの消耗を犠牲にします）。
スクリーン セーバーが起動しても、画面がロックされることがあります。 スクリーン セーバーを停止します。
常に起動しないようにするため、自動的にロックされることはありません。
デフォルト -currentHost 書き込み com.apple.screensaver idleTime 0
7. SSH 経由のクリップボード同期
macOS には、pbcopy (クリップボードの書き込み) と pbpaste (クリップボードの読み取り) が同梱されています。 SSH経由でパイプされ、
クリップボードをマシン間で移動できます - 暗号化、ピアツーピア、Apple ID なし、または
サードパーティのサービス。
Clip.sh はこれを 2 つのサブコマンドを含む 1 つのコマンドにラップし、画像を追加します
pbcopy / pbpaste (テキストのみ) の上でのサポート。 PATH にスクリプトとしてインストールします。
ソース Mac 上で、IC_BOX を使用してターゲット ホストを指します (「ic」は
「孤立したクロード」):
カール -fsSL https://raw.githubusercontent.com/ykdojo/claude-controls-mac/main/clip.sh -o ~/.local/bin/clip
chmod +x ~/.local/bin/clip
import IC_BOX = "<user>@<target-host>.local" # ~/.zshrc に追加
使用法:
クリップ送信 - この Mac のクリップボード → ターゲット (テキストまたは画像)。画像の場合は、次のことができます
Ctrl-V を使用して、ターゲット上のクロード コード セッションに直接貼り付けます。
クリップ取得 - ターゲットのクリップボード → この Mac (テキストまたは画像)。
8. 対象の Mac に Claude Code をインストールします。
インストール コマンドを送信して実行します。ソース Mac からはまっすぐに押すことができます
ターゲットのクリップボードにコピーするか、リモートで実行します。次のコマンドは、
特定のバージョンをインストールしますが、必要に応じて最新または安定したバージョンをインストールすることもできます。
ssh <us

er>@<ターゲットホスト>.local 'curl -fsSL https://claude.ai/install.sh | bash -s -- 2.1.201'
ネイティブ インストーラーは、~/.local/bin が PATH 上にないことを警告する場合があります。ターゲット上で修正します。
それを ~/.zshenv ( ~/.zshrc ではない) に追加します - .zshenv は、以下を含むすべての zsh によって読み取られます。
非インタラクティブなもの:
ssh <ユーザー>@<ターゲットホスト>.local 'echo ' \' 'export PATH="$HOME/.local/bin:$PATH"' \' ' >> ~/.zshenv'
9. 独自のクロード コードに優しい環境をセットアップする (オプション)
このオプションのステップでは、独自のデフォルト設定を次のように適用します。
setup-claude-env.sh - シェルエイリアス、DX
プラグイン、settings.json 調整、GitHub CLI、および (オプトイン) Playwright MCP と
yt-dlp。すべての項目は切り替え可能です。の完全なリストを参照してください
クロード-env-components.md 。
ターゲット上で対話的に - すべての項目のチェックリストを表示します (コア
事前にチェックが入っており、オプトインはチェックされていません）ので、任意の組み合わせを選択できます。ダウンロード
スクリプトをターゲットに追加して実行します。
ssh -t <ユーザー>@<ターゲットホスト>.local \
'curl -fsSL https://raw.githubusercontent.com/ykdojo/claude-controls-mac/main/setup-claude-env.sh -o setup-claude-env.sh && bash setup-claude-env.sh'
非対話型 - スクリプトは、
非対話型環境であるか、少なくとも 1 つのフラグを指定します。プロンプトはありません。インストールします
デフォルトではコアのみ、またはフラグ ( --yt-dlp 、 --playwright 、
--all 、 --core ):
ssh <ユーザー>@<ターゲットホスト>.local \
'curl -fsSL https://raw.githubusercontent.com/ykdojo/claude-controls-mac/main/setup-claude-env.sh -o setup-claude-env.sh && bash setup-claude-env.sh --all'
スクリプトはべき等です (再実行しても問題ありません)。
10. クロードとGitHubにログインします。
どちらのログインも対話型であるため、SSH で次のように接続します。
ssh <ユーザー>@<ターゲットホスト>.local
次に、ターゲットに対して claude を実行します。これは Anthropic のログインにドロップされます。
アカウント。プロンプトに従います (ブラウザー/デバイスのコード フローで実行できます)。

ブラウザから終了する
メインの Mac 上)。
GitHub - オプションですが、エージェントがリポジトリを操作できるようにすることを強くお勧めします。
GH 認証ログイン
GitHub CLI をまだインストールしていない場合は、
前のステップ、
必ず最初にそうしてください。
個人的には、メインの GitHub アカウントではなく、別の GitHub アカウントを使用することをお勧めします。
メインアカウントを台無しにすることはありません。
11. SSH 経由のコンピュータの使用 (オプション)
これにより、ターゲット上のインタラクティブなクロード セッションで (スクリーンショット) を確認し、制御できるようになります。
(マウス/キーボード) SSH 経由で駆動される独自のデスクトップ。
これはそのままでは機能しません。SSH と macOS の権限モデルが邪魔をするため、
それを回避するために、以下のセットアップが存在します。
回避策が必要な理由: macOS は画面録画の背後で画面キャプチャと入力をゲートします。
GUI ログイン セッションに関連付けられたアクセシビリティ権限、つまり SSH
プロセスがディスプレイに到達できません。修正: LaunchAgent は tmux サーバーを内部で存続させます。
固定ソケット上の GUI セッション。そこで作成されたすべてのクロード セッションはそこに到達します
サーバーに接続され、GUI セッションを継承するため、ディスプレイに到達できます。 SSH 経由で接続します。
ターゲット上で setup-computer-use.sh を実行します。
ssh -t <ユーザー>@<ターゲットホスト>.local \
'curl -fsSL https://raw.githubusercontent.com/ykdojo/claude-controls-mac/main/setup-computer-use.sh -o setup-computer-use.sh && bash s

[切り捨てられた]

## Original Extract

Step-by-step guide to turning your spare Mac into an always-on machine Claude Code can fully control - talk to it from your phone through the Claude app or from your main Mac over SSH.

How to set up your spare Mac for Claude Code to fully control - a step-by-step guide
Here’s a full step-by-step guide on how to turn your spare Mac into an always-on
machine Claude Code can fully control, with computer use enabled. You’ll be able to talk to it from your phone
through the Claude app, or from your main Mac over SSH.
I wanted to create a separate environment Claude Code can control on its own, so I can
delegate tasks I don’t necessarily want to run on my own machine - certain types of
research tasks, and development tasks.
Claude Code, especially with the --dangerously-skip-permissions flag on,
carries inherent risk when run on your main machine. You can eliminate / mitigate these
risks by creating a separate environment on your spare Mac with everything it needs
to have access to.
It has an added bonus of being able to talk to Claude Code anytime, anywhere from
your phone. I’ve personally found it really useful because I often prefer to talk to
Claude Code instead of regular Claude on the mobile app - Claude Code is often more
capable.
The following guide assumes you have your main Mac as well as a spare Mac you can set
up for this, but you should be able to take inspiration from it and apply it to any
combination of two machines.
First, let’s quickly address a few questions you might have.
Why not run it in a container?
I’m a big proponent of running it in a container - I even built
an entire environment for doing so conveniently .
However, I’ve found it has a few limitations.
First, it still runs on your main machine, so it’s not completely separated. For
example, network requests it sends still go through your main machine.
Second, there are limitations to the container’s capabilities. For example, I wanted
my agent to be able to run Unity for game development, and there’s no easy way to do
that in a container. The same goes for any other app that’s only available on a Mac -
you won’t have access to it. That’s especially relevant if you want Claude Code to
control these apps through computer use - clicking, dragging, and so on.
Why not use something like OpenClaw?
I personally like having access to the full, latest features of Claude Code. I also
like being able to control it from the Claude app - I’ve found it really convenient.
And you get to use your Claude subscription usage if you happen to have one, which is
an added bonus.
At the end of the day, running an agent with broad permissions is safer on a machine
that has nothing to lose - but you get the benefit of being able to use a full Mac
instead of a container. The approach here:
Use an old/spare Mac , not your main one.
Create a fresh local account with no personal data and no Apple ID signed in, so
the agent has nothing sensitive to reach.
Drive it over SSH from your main Mac on your local network, and control it from
your phone.
Your everyday Mac (the source ), on the same Wi-Fi.
1. Start fresh on the target Mac
Wipe it first (if it has any personal data)
You’ll be giving the agent full access to this machine, so it can reach anything stored
on it. If there’s existing data you don’t want it to have access to, erase the machine
first:
Macs that support it: System Settings -> General -> Transfer or Reset ->
Erase All Content and Settings .
Older Intel Macs: restart into Recovery (hold Cmd-R at
boot), use Disk Utility to erase the internal drive, then reinstall macOS.
Optionally update to the latest macOS afterward (System Settings -> General ->
Software Update).
Create a fresh, isolated account
Create a new local user account (System Settings -> Users & Groups).
I recommend not signing into an Apple ID. Skip it during setup.
Make the account an admin (if you haven’t already)
The account needs admin rights or sudo will refuse to run.
System Settings -> Users & Groups -> set the account to Allow this user to
administer this computer .
If you ever need to repair it from another admin account:
sudo dseditgroup -o edit -a <user> -t user admin
2. Enable Remote Login (SSH) on the target Mac
On the target , turn on SSH so the source Mac can connect:
sudo systemsetup -setremotelogin on
If the command fails with Turning Remote Login on or off requires Full Disk Access
privileges , give your terminal app Full Disk Access first:
System Settings -> Privacy & Security -> Full Disk Access .
Click + , then in the file picker go to Applications -> Utilities -> Terminal
and add it.
Quit and reopen the terminal, then rerun the command.
3. Passwordless sudo for the target account
This is so the agent (and your SSH commands) can run admin tasks without a password
prompt each time. Run this once on the target. It
asks for the login password this one time:
echo "<user> ALL=(ALL) NOPASSWD: ALL" | sudo tee /etc/sudoers.d/<user>-nopasswd > /dev/null
sudo chmod 440 /etc/sudoers.d/<user>-nopasswd
sudo visudo -cf /etc/sudoers.d/<user>-nopasswd # validate - must print 'parsed OK'
This creates a small rule file telling the Mac that <user> can run sudo without a
password prompt:
line 1 writes the rule into /etc/sudoers.d/ .
line 2 makes it read-only - sudo ignores the file otherwise.
line 3 validates the syntax; a typo in a sudoers file can lock you out of sudo
entirely, so it must print parsed OK .
After this, sudo runs with no prompt - test with sudo -n true , which succeeds
silently if passwordless sudo works.
4. Find the target’s address (hostname or IP)
You can reach the target by either a hostname or an IP. I recommend using the
hostname: it stays the same, while the IP can change.
Hostname (recommended). Run on the target:
scutil --get LocalHostName # prints the hostname, e.g. MacBook-Pro
Add .local to form the address: <target-host>.local . You can also read it from
System Settings -> General -> Sharing, shown as Local hostname .
Give the target a unique name. Each Mac needs a .local name that’s unique on
your network. If two machines share a name, the address can point to the wrong Mac.
Make sure the target’s name is unique - rename it if needed:
sudo scutil --set LocalHostName newmacbook # -> newmacbook.local
IP address (not recommended). Run on the target:
ipconfig getifaddr en0 # e.g. 192.168.1.80
Note that the IP can change after a reboot or after a certain amount of time.
Throughout the rest of this guide, replace <user> with the target account name and
<target-host> with the hostname from above, so the address is
<user>@<target-host>.local . You could also instead use an IP in place of
<target-host>.local .
5. Set up passwordless SSH from the source Mac
On the source Mac, create an SSH key (skip if you already have one):
ssh-keygen -t ed25519
Install your public key on the target. This asks for the target account’s login
password once:
ssh-copy-id <user>@<target-host>.local
Test it - this should print the target username with no password prompt:
ssh <user>@<target-host>.local whoami
6. Keep the target awake
By default macOS sleeps after ~10 minutes idle, even when plugged in, which takes
it off the network. To make it never sleep, run this on the target (or over SSH from
the source):
sudo pmset -c sleep 0 # never system-sleep while plugged in (-c = on charger)
sudo pmset -c disablesleep 1 # also prevents sleep with the lid closed (clamshell)
sudo pmset -c displaysleep 0 # keep the display on too
Verify:
pmset -g | grep -iE 'sleep'
sleep 0 , SleepDisabled 1 , and displaysleep 0 in the output confirm it worked.
If the machine runs on battery sometimes, use -a instead of -c to apply to all
power sources (at the cost of battery drain).
The screen can still lock when the screen saver kicks in. Stop the screen saver
from ever starting so it never locks on its own:
defaults -currentHost write com.apple.screensaver idleTime 0
7. Clipboard sync over SSH
macOS ships pbcopy (write clipboard) and pbpaste (read clipboard). Piped over SSH,
they can move the clipboard between machines - encrypted, peer-to-peer, no Apple ID or
third-party service.
clip.sh wraps this into one command with two subcommands, and adds image
support on top of pbcopy / pbpaste (which are text-only). Install it as a script on your PATH
on the source Mac, and point it at the target host with IC_BOX (“ic” standing for
“isolated claude”):
curl -fsSL https://raw.githubusercontent.com/ykdojo/claude-controls-mac/main/clip.sh -o ~/.local/bin/clip
chmod +x ~/.local/bin/clip
export IC_BOX = "<user>@<target-host>.local" # add to ~/.zshrc
Usage:
clip send - this Mac’s clipboard → the target (text or image). For an image you can
paste it straight into a Claude Code session on the target with Ctrl-V.
clip get - the target’s clipboard → this Mac (text or image).
8. Install Claude Code on the target Mac
Send the install command over and run it. From the source Mac you can push it straight
to the target’s clipboard, or run it remotely. The following command installs a
specific version, but you can also install latest or stable if you’d like:
ssh <user>@<target-host>.local 'curl -fsSL https://claude.ai/install.sh | bash -s -- 2.1.201'
The native installer may warn that ~/.local/bin is not on PATH. Fix it on the target by
adding it to ~/.zshenv (not ~/.zshrc ) - .zshenv is read by every zsh, including
non-interactive ones:
ssh <user>@<target-host>.local 'echo ' \' 'export PATH="$HOME/.local/bin:$PATH"' \' ' >> ~/.zshenv'
9. Set up an opinionated, Claude Code-friendly environment (optional)
This optional step applies a set of opinionated defaults via
setup-claude-env.sh - shell aliases, the DX
plugin, settings.json tweaks, the GitHub CLI, and (opt-in) Playwright MCP and
yt-dlp. Every item is toggleable; see the full list in
claude-env-components.md .
Interactively on the target - shows a checklist of every item (core
pre-checked, opt-ins unchecked) so you can pick any combination. Download the
script onto the target and run it:
ssh -t <user>@<target-host>.local \
'curl -fsSL https://raw.githubusercontent.com/ykdojo/claude-controls-mac/main/setup-claude-env.sh -o setup-claude-env.sh && bash setup-claude-env.sh'
Non-interactively - the script goes into this mode if it detects a
non-interactive environment or you supply at least one flag. No prompt; installs
core only by default, or pick what you want with flags ( --yt-dlp , --playwright ,
--all , --core ):
ssh <user>@<target-host>.local \
'curl -fsSL https://raw.githubusercontent.com/ykdojo/claude-controls-mac/main/setup-claude-env.sh -o setup-claude-env.sh && bash setup-claude-env.sh --all'
The script is idempotent (OK to re-run).
10. Log in to Claude and GitHub
Both logins are interactive, so SSH in:
ssh <user>@<target-host>.local
Then run claude on the target - it drops into the login for your Anthropic
account. Follow the prompts (a browser/device-code flow you can finish from a browser
on your main Mac).
GitHub - optional, but highly recommended so the agent can work with repos:
gh auth login
If you haven’t installed the GitHub CLI from
the previous step ,
make sure to do so first.
I personally recommend using a separate GitHub account, not your main one, so it
doesn’t mess up your main account.
11. Computer use over SSH (optional)
This lets an interactive claude session on the target see (screenshots) and control
(mouse/keyboard) its own desktop, driven over SSH.
This doesn’t work out of the box - SSH and macOS’s permission model get in the way, so the
setup below exists to route around that.
Why it needs a workaround: macOS gates screen capture and input behind Screen Recording
and Accessibility permissions that are tied to the GUI login session, so an SSH
process can’t reach the display. Fix: a LaunchAgent keeps a tmux server alive inside the
GUI session on a fixed socket; every claude session created there lands on that
server and inherits the GUI session, so it can reach the display. You attach over SSH.
Run setup-computer-use.sh on the target:
ssh -t <user>@<target-host>.local \
'curl -fsSL https://raw.githubusercontent.com/ykdojo/claude-controls-mac/main/setup-computer-use.sh -o setup-computer-use.sh && bash s

[truncated]
