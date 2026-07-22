---
source: "https://github.com/Mounstroya/kindle-ssh-guide"
hn_url: "https://news.ycombinator.com/item?id=49013970"
title: "Running Claude Code on a Jailbroken Kindle over SSH via Tailscale"
article_title: "GitHub - Mounstroya/kindle-ssh-guide: The missing guide: get a working outbound SSH client on a jailbroken Kindle over Tailscale, no OpenSSH needed · GitHub"
author: "YaelMontoya"
captured_at: "2026-07-22T21:57:25Z"
capture_tool: "hn-digest"
hn_id: 49013970
score: 1
comments: 0
posted_at: "2026-07-22T21:53:05Z"
tags:
  - hacker-news
  - translated
---

# Running Claude Code on a Jailbroken Kindle over SSH via Tailscale

- HN: [49013970](https://news.ycombinator.com/item?id=49013970)
- Source: [github.com](https://github.com/Mounstroya/kindle-ssh-guide)
- Score: 1
- Comments: 0
- Posted: 2026-07-22T21:53:05Z

## Translation

タイトル: Tailscale 経由で SSH 経由で脱獄した Kindle でクロード コードを実行する
記事のタイトル: GitHub - Mounstroya/kindle-ssh-guide: 欠けているガイド: Tailscale を介してジェイルブレイクした Kindle で動作するアウトバウンド SSH クライアントを取得する、OpenSSH は必要ありません · GitHub
説明: 不足しているガイド: Tailscale 経由でジェイルブレイクした Kindle で動作する送信 SSH クライアントを取得します。OpenSSH は必要ありません - Mounstroya/kindle-ssh-guide

記事本文:
GitHub - Mounstroya/kindle-ssh-guide: 不足しているガイド: Tailscale 経由でジェイルブレイクした Kindle で動作する送信 SSH クライアントを取得します。OpenSSH は必要ありません · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
マウンストロヤ
/
Kindle-SSH-ガイド
公共
通知

ション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット パッチ パッチ スクリプト スクリプト .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
脱獄した Kindle で SSH クライアントを動作させる方法
実際に動作する SSH プロンプトを取得するための完全なゼロからのガイド
ジェイルブレイクされた Kindle — いくつかのガイドによる SSH アクセスだけではありません
すでにカバーしています。私が見つけた限り、完全な機能を公開した人は誰もいません
デバイス自体からのアウトバウンド SSH へのパス (Tailscale 経由、なし)
完全な OpenSSH クライアントをインストールします。これがそのガイドです。
最終結果: Kindle の e-ink 画面でターミナル アプリを開き、次のように入力します。
sshc youruser@your-server を実行し、任意のマシン上の実際のシェルを取得します。
テールスケールネットワーク。
これがそのままでは機能しない理由
2 つの既存のプロジェクトで、目的の 90% を達成できます。
tailscale_kual はあなたの
Kindle をテールネットに接続すると、他のデバイスからアクセスできるようになります。
kterm は実際の値を提供します
オンスクリーン キーボードを備えた Kindle 画面上のターミナル エミュレータ。
これらを組み合わせると、「 ssh user@host 」と入力するだけで済みます。それはできません —
Kindle には SSH クライアントがまったく同梱されておらず、SSH クライアントを構築/インストールする必要があります。
あなたは 2 つの別々の壁にぶつかります。
1. 静的 glibc バイナリがこのカーネルでクラッシュします。
dbclient (Dropbear の SSH クライアント) を通常のクライアントと静的にクロスコンパイルする
gcc-arm-linux-gnueabihf ツールチェーンは中止するバイナリを生成します
Kindle ですぐに:
dbclient: dl-call-libc-early-init.c:37: _dl_call_libc_early_init:
アサーション `sym != NULL' が失敗しました。
中止されました
これは、プレーンな -static と -static -no-pie の両方で同様に発生します。来る
glibc >= 2.34 の ld.so/libc.so マージから、
_dl_ca

ll_libc_early_init を非 PIE の場合でも標準起動パスに追加
静的実行可能ファイル - Kindle の古いバージョンとは相性が良くありません
lab126 4.9 カーネル。
2. ユーザー空間モードには、アプリの送信ルートがありません。
tailscale_kual の標準 (ユーザー空間) モードは、それ自体を tailscaled にします
到達可能ですが、TUN インターフェイスは追加されません — Kindle 上のアプリ (
新しく構築された dbclient ) には、テールネットへのカーネル ルートがまったくありません。カーネル
ほとんどの Kindle には tun.ko が同梱されていないため、TUN モードもオプションではありません。
これでプロキシ モード (SOCKS5) が残ります。
ジェイルブレイクされた Kindle に関する Tailscale の公式記事
カバー — しかし、Dropbear の dbclient は SOCKS を話すことができないため、ガイドは
同様の組み込みデバイスの場合 (例:
注目に値するタブレット）教えてください
代わりに完全な OpenSSH クライアントをインストールします。
このガイドでは、OpenSSH をまったく必要とせずに両方を解決します。
代わりに musl ツールチェーンを使用して dbclient /dropbearkey をクロスコンパイルします
glibcの。静的 musl バイナリには glibc の早期初期化機構がありません。
だから彼らはただ走るだけだ。
SOCKS プロキシの代わりに、tailscaled 自身のプロキシを介して dbclient をパイプします。
-J 経由の tailscale nc (Dropbear の ProxyCommand と同等)。テールスケール
モードに関係なく、独自のユーザースペースのネットスタックを介してダイヤルアウトできます
それは含まれています — tailscale nc はそれをサブプロセスとして公開するだけなので、カーネルはありません
ルートまたは TUN デバイスが必要です。
カーネル 4.9.77-lab126、Kindle ファームウェア 5.17.1 を実行している Kindle でテスト済み。
Tailscale クライアント 1.98.9、Dropbear 2025.89。
KUAL がインストールされた脱獄済みの Kindle。
クロスコンパイルするマシン (任意の Linux x86_64 ボックス、Docker コンテナ)
正常に動作し、ビルド スクリプトではそれ以上のことは何も想定されていません)。
ステップ 1 — tailscale_kual をインストールする
tailscale_kual さん自身をフォローしてください
手順: extensions/tailscale フォルダーを /mnt/us/extensions/ にドロップします。
認証キーを extensions/tailscale/config/auth_key.txt に置きます。

KUALを開いて、
Start Tailscale を実行します (標準/ユーザースペース モード — このガイドは構築されています)
このモードは Kindle で動作するモードなので、
つん.こ）。
テールネット上の別のマシンから起動していることを確認します。
テールスケールのステータス | grep Kindle
ステップ 2 — kterm をインストールする
-armhf ビルドをダウンロードします (ファームウェア > 5.16.3 が必要です)。
kterm のリリースを解凍し、
kterm フォルダーを /mnt/us/extensions/ にドロップします。に表示されます。
クアルのメニュー。
ステップ 3 — musl を使用して静的 dbclient を構築する
git clone https://github.com/Mounstroya/kindle-ssh-guide
cd Kindle-ssh-ガイド
./scripts/build-dbclient-musl.sh
これにより、musl.cc ARM ツールチェーンと Dropbear の
ソースを作成し、dbclient +dropbearkey を静的、非 PIE としてクロスコンパイルします。
ARM バイナリ。出力は ./out/ にあります。
(このスクリプトは、DROPBEAR_SVR_PASSWORD_AUTH (サーバーのみ) もオフにします。
オプションを使用すると、crypt() が取り込まれ、静的リンクが切断されます。
dbclient は決して使用しません。)
両方のバイナリを Kindle の同じフォルダー tailscale_kual にコピーします。
バイナリは次の場所に存在します。
scp out/dbclient out/dropbearkey root@ < kindle-tailscale-ip > :/mnt/us/extensions/tailscale/bin/
ステップ 4 — キーを生成して認証する
Dropbear の dbclient -i は Dropbear 独自のバイナリ キー形式のみを読み取ります。
ssh-keygen が生成する OpenSSH PEM 形式。 OpenSSH キーを使用しようとしています
不可解な dbclient: Exited: String too long で失敗します。キーを生成する
代わりに、Kindle で直接 Dropbearkey を使用します。
chmod 755 /mnt/us/extensions/tailscale/bin/dbclient /mnt/us/extensions/tailscale/bin/dropbearkey
/mnt/us/extensions/tailscale/bin/dropbearkey -t ed25519 \
-f /mnt/us/extensions/tailscale/bin/kindle_native_key
出力された ssh-ed25519 AAAA... 行を ~/.ssh/authorized_keys にコピーします。
SSH で接続するマシン。
cat > /mnt/us/extensions/tailscale/bin/sshc < scripts/ssh

c
chmod 755 /mnt/us/extensions/tailscale/bin/sshc
このファイルに ssh という名前を付けないでください。 tailscale_kual はすでに出荷しています
bin/ の下にある ssh/ という名前のディレクトリ。埋め込みファイルの保存に使用されます。
Tailscale SSH サーバーのホスト キー。 ssh という名前のファイルが静かに着陸します
そのディレクトリ内 ( scp の destination-is-a-directory 動作経由)
期待した場所に作成されるのではなく、コピーのように見えます
成功すると、実行すると「許可が拒否されました」というメッセージが表示されます。 sshc
これを完全に回避します。
sshc は、以下を実行する小さなラッパーです。
dbclient -y -i kindle_native_key -J " tailscale nc $HOST 22 " " $TARGET " " $@ "
つまり、 tailscale nc を介して SSH 接続をトンネリングするため、
ただし、Kindle にはテールネットへのカーネル ルートがありません。
ステップ6 — kterm内のPATHに配置します
patches/kterm.sh.diff にある 1 行のパッチを以下に適用します。
/mnt/us/extensions/kterm/bin/kterm.sh (または手動で行を追加するだけです)
kterm バイナリを起動する行の直前):
エクスポート PATH=/mnt/us/extensions/tailscale/bin: $PATH
ステップ 7 — 使用する
Kindle で KUAL → kterm を開き、次のようにします。
sshc youruser@100.x.y.z
これだけです。Tailscale 上の e-ink 画面上に本物のシェルが表示されます。
ターゲットの解像度は tailscale nc を通過するため、
Kindle のテールネットから到達可能なテールネット IP または MagicDNS 名。
sshc はポート 22 をハードコードします。ターゲットが別のポートを使用している場合は、スクリプトを編集します。
ポート。
dbclient -h 自体が _dl_call_libc_early_init でクラッシュした場合
アサーション、musl ではなく glibc でビルドしました — ステップ 3 を再確認してください。
String too long で認証が失敗した場合は、次のようにキーを生成しました。
Dropbearkey の代わりに ssh-keygen — ステップ 4 を参照してください。
以下を置き換えるものではなく、以下の上に構築されます。
tailscale_kual — ジェイルブレイクされた Kindle の Tailscale
kterm — Kindle 用ターミナルエミュレータ
MIT — 「ライセンス」を参照してください。 Dropbear 自体は個別にライセンスされます (MIT スタイル)。独自の を参照してください

詳細についてはこちらをご覧ください。
欠けているガイド: 脱獄した Kindle で Tailscale 経由で動作する送信 SSH クライアントを取得します。OpenSSH は必要ありません
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The missing guide: get a working outbound SSH client on a jailbroken Kindle over Tailscale, no OpenSSH needed - Mounstroya/kindle-ssh-guide

GitHub - Mounstroya/kindle-ssh-guide: The missing guide: get a working outbound SSH client on a jailbroken Kindle over Tailscale, no OpenSSH needed · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Mounstroya
/
kindle-ssh-guide
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits patches patches scripts scripts .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
How to Get a Working SSH Client on a Jailbroken Kindle
A complete, from-scratch guide to getting a real, working ssh prompt on
a jailbroken Kindle — not just SSH access into it, which several guides
already cover. As far as I could find, nobody has published a full working
path to outbound SSH from the device itself, over Tailscale, without
installing the full OpenSSH client. This is that guide.
End result: open a terminal app on your Kindle's e-ink screen, type
sshc youruser@your-server , and get a real shell on any machine on your
Tailscale network.
Why this doesn't just work out of the box
Two existing projects get you 90% of the way there:
tailscale_kual puts your
Kindle on your tailnet and makes it reachable from other devices.
kterm gives you an actual
terminal emulator on the Kindle's screen, with an on-screen keyboard.
Put them together and you'd expect to just type ssh user@host . You can't —
the Kindle ships with no SSH client at all, and building/installing one
yourself hits two separate walls:
1. Static glibc binaries crash on this kernel.
Cross-compiling dbclient (Dropbear's SSH client) statically with a normal
gcc-arm-linux-gnueabihf toolchain produces a binary that aborts
immediately on the Kindle:
dbclient: dl-call-libc-early-init.c:37: _dl_call_libc_early_init:
Assertion `sym != NULL' failed.
Aborted
This happens with plain -static and with -static -no-pie alike. It comes
from glibc >= 2.34's ld.so/libc.so merge, which wires
_dl_call_libc_early_init into the standard startup path even for non-PIE
static executables — and it doesn't get along with the Kindle's ancient
lab126 4.9 kernel.
2. Userspace mode has no outbound route for apps.
tailscale_kual 's Standard (Userspace) mode makes tailscaled itself
reachable, but it adds no TUN interface — apps on the Kindle (including a
freshly built dbclient ) have no kernel route to the tailnet at all. Kernel
TUN mode isn't an option either, because most Kindles don't ship tun.ko .
That leaves Proxy mode (SOCKS5), which the
official Tailscale writeup for jailbroken Kindles
covers — but Dropbear's dbclient can't speak SOCKS, which is why guides
for similar embedded devices (e.g. the
reMarkable tablet ) tell you
to install the full OpenSSH client instead.
This guide solves both without needing OpenSSH at all:
Cross-compile dbclient / dropbearkey with a musl toolchain instead
of glibc. Static musl binaries don't have glibc's early-init machinery,
so they just run.
Instead of a SOCKS proxy, pipe dbclient through tailscaled's own
tailscale nc via -J (Dropbear's ProxyCommand equivalent). tailscaled
can dial out over its own userspace netstack regardless of what mode
it's in — tailscale nc just exposes that as a subprocess, so no kernel
route or TUN device is needed.
Tested on a Kindle running kernel 4.9.77-lab126 , Kindle firmware 5.17.1,
Tailscale client 1.98.9, Dropbear 2025.89.
A jailbroken Kindle with KUAL installed.
A machine to cross-compile on (any Linux x86_64 box; a Docker container
works fine and is what the build script assumes nothing beyond).
Step 1 — Install tailscale_kual
Follow tailscale_kual 's own
instructions: drop the extensions/tailscale folder onto /mnt/us/extensions/ ,
put an auth key in extensions/tailscale/config/auth_key.txt , open KUAL and
run Start Tailscale (Standard/Userspace mode — this guide is built
around that mode, since it's the one that works on Kindles without
tun.ko ).
Verify it's up from another machine on your tailnet:
tailscale status | grep kindle
Step 2 — Install kterm
Download the -armhf build (firmware > 5.16.3 needs it) from
kterm's releases , unzip it,
and drop the kterm folder into /mnt/us/extensions/ . It'll show up in the
KUAL menu.
Step 3 — Build a static dbclient with musl
git clone https://github.com/Mounstroya/kindle-ssh-guide
cd kindle-ssh-guide
./scripts/build-dbclient-musl.sh
This downloads a musl.cc ARM toolchain and Dropbear's
source, and cross-compiles dbclient + dropbearkey as static, non-PIE
ARM binaries. Output lands in ./out/ .
(The script also flips off DROPBEAR_SVR_PASSWORD_AUTH — a server-only
option that otherwise pulls in crypt() and breaks the static link.
dbclient never uses it.)
Copy both binaries to the Kindle, into the same folder tailscale_kual's
binaries live in:
scp out/dbclient out/dropbearkey root@ < kindle-tailscale-ip > :/mnt/us/extensions/tailscale/bin/
Step 4 — Generate a key and authorize it
Dropbear's dbclient -i only reads Dropbear's own binary key format — not
the OpenSSH PEM format ssh-keygen produces. Trying to use an OpenSSH key
fails with a cryptic dbclient: Exited: String too long . Generate the key
with dropbearkey instead, directly on the Kindle:
chmod 755 /mnt/us/extensions/tailscale/bin/dbclient /mnt/us/extensions/tailscale/bin/dropbearkey
/mnt/us/extensions/tailscale/bin/dropbearkey -t ed25519 \
-f /mnt/us/extensions/tailscale/bin/kindle_native_key
Copy the printed ssh-ed25519 AAAA... line into ~/.ssh/authorized_keys on
the machine you want to SSH into.
cat > /mnt/us/extensions/tailscale/bin/sshc < scripts/sshc
chmod 755 /mnt/us/extensions/tailscale/bin/sshc
Don't name this file ssh . tailscale_kual already ships a
directory named ssh/ under bin/ , used to store the embedded
Tailscale SSH server's host keys. A file named ssh silently lands
inside that directory (via scp 's destination-is-a-directory behavior)
instead of being created where you expect — it'll look like the copy
succeeded and then you'll get Permission denied running it. sshc
sidesteps this entirely.
sshc is a small wrapper that runs:
dbclient -y -i kindle_native_key -J " tailscale nc $HOST 22 " " $TARGET " " $@ "
i.e. it tunnels the SSH connection through tailscale nc , so it works even
though the Kindle has no kernel route to the tailnet.
Step 6 — Put it on your PATH inside kterm
Apply the one-line patch in patches/kterm.sh.diff to
/mnt/us/extensions/kterm/bin/kterm.sh (or just add the line by hand,
right before the line that launches the kterm binary):
export PATH=/mnt/us/extensions/tailscale/bin: $PATH
Step 7 — Use it
Open KUAL → kterm on the Kindle, then:
sshc youruser@100.x.y.z
That's it — a real shell, on your e-ink screen, over Tailscale.
Target resolution goes through tailscale nc , so it needs to be a
tailnet IP or MagicDNS name reachable from the Kindle's tailnet.
sshc hardcodes port 22. Edit the script if your target uses a different
port.
If dbclient -h itself crashes with an _dl_call_libc_early_init
assertion, you built with glibc, not musl — re-check step 3.
If auth fails with String too long , you generated the key with
ssh-keygen instead of dropbearkey — see step 4.
Built on top of, not a replacement for:
tailscale_kual — Tailscale on jailbroken Kindles
kterm — terminal emulator for Kindle
MIT — see LICENSE . Dropbear itself is separately licensed (MIT-style); see its own source for details.
The missing guide: get a working outbound SSH client on a jailbroken Kindle over Tailscale, no OpenSSH needed
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
