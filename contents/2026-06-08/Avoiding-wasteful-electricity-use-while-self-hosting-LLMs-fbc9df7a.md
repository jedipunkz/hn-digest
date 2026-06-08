---
source: "https://patrickmccanna.net/fixing-a-stuck-ollama-runner-and-building-a-gpu-watchdog/"
hn_url: "https://news.ycombinator.com/item?id=48451624"
title: "Avoiding wasteful electricity use while self hosting LLMs"
article_title: "Fixing a stuck Ollama runner and building a GPU watchdog – blog"
author: "0o_MrPatrick_o0"
captured_at: "2026-06-08T21:38:00Z"
capture_tool: "hn-digest"
hn_id: 48451624
score: 1
comments: 0
posted_at: "2026-06-08T20:38:51Z"
tags:
  - hacker-news
  - translated
---

# Avoiding wasteful electricity use while self hosting LLMs

- HN: [48451624](https://news.ycombinator.com/item?id=48451624)
- Source: [patrickmccanna.net](https://patrickmccanna.net/fixing-a-stuck-ollama-runner-and-building-a-gpu-watchdog/)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T20:38:51Z

## Translation

タイトル: セルフホスティング LLM での無駄な電力使用の回避
記事のタイトル: スタックした Ollama ランナーを修正し、GPU ウォッチドッグを構築する – ブログ

記事本文:
スタックした Ollama ランナーを修正し、GPU ウォッチドッグを構築する – ブログ
コンテンツにスキップ
ブログ
モバイル セキュリティのメンターシップとテクノロジー分野で働くよう子供たちを鼓舞する
スタックした Ollama ランナーを修正し、GPU ウォッチドッグを構築する
LLM を自己ホストしている場合、ある時点で GPU が長時間にわたって大量の電力を消費するスタックに遭遇する可能性があります。これは、問題の発見と、問題を修正する自動化の実装についての私の要約です。これらの問題を自動化して発見して修正してほしいと考えています。不必要な電力消費に対する最前線にはなりたくないのです。
私は自宅のハードウェア、つまりミニ PC (統合 GPU を搭載した AMD Ryzen AI MAX+ 395) 上の Ollama サーバーで大規模な言語モデルを実行しています。セルフホスティングにより、プライバシーが得られ、サブスクリプションが不要で、インターネット接続なしで動作するモデルが得られます。その代わりに、私は自分の運用チームになるということです。何か問題が発生した場合、サービスの停止を監視するプロバイダーは存在しません。
今週末、何か問題が起こりました。
AI サーバーの冷却ファンがフルスピードで停止することなく動作していることに気づきました。私はそれを使用していませんでした。そして私の家族の使用はかなり限られていました。システムに対するアクティブなリクエストがなかった可能性があります。サーバーが侵害されて暗号通貨のマイニングが行われたか、プロセスがハングして無駄に電力が消費されたかのどちらかです。このシステムは最大 85 W を消費します。フルパワーで無期限に固定したままにしておくと、電気代に反映されてしまいます。
2 つのことを行う必要がありました。まず、マシンが実際の作業を行っているか、停止しているかを判断します。次に、この種の問題を自動的に検出するソリューションを構築します。長時間実行されるファン アクティビティの監視では、平均解決時間の信頼性が低くなります。
診断。 ollam ps でモデルが何をしているのかを確認したところ、モデルが Stopping... 状態でスタックしていることがわかりました。それは su でした。

GPU をアンロードして解放することを提案しましたが、基礎となるプロセスが終了することはありませんでした。カーネルから GPU 使用率ゲージ ( /sys/class/drm/card1/device/gpu_busy_percent ) を読み取って条件を確認しました: ~89% ビジー、~85W。約 20 時間この状態が続き、推論リクエストはありませんでした。通常のシャットダウン コマンド ( ollam stop ) は効果がありませんでした。サービスの再起動 ( sudo systemctl restart ollama ) によりこれがクリアされました。根本的な原因は、GPU が何もせずにフル稼働状態に留まる Ollama の既知のバグです。
番犬。この時点で、Ollama システムがハングしており、リセットする必要があることを確認しました。しかし、この問題は以前にも発生しました。おそらくまた起こるだろう。サーバーから発せられる予期せぬゼファーを検出する能力に依存するのをやめる必要がありました。この特定の障害が発生した場合にのみ警告を発するウォッチドッグを作成しました。誤検知を回避するには、ウォッチドッグのロジックを狭くする必要があります。システムの基本は次のとおりです。
cron ジョブは 5 分ごとに 2 つの読み取りでシステムをサンプリングします: モデル プロセスは読み込まれていますか? GPU のビジー状態はどれくらいですか?
サンプルが異常であるとみなされるのは、モデルがロードされ、GPU の使用率が 70% 以上であるという両方の条件が当てはまる場合のみです。推論中に GPU 使用率が高くなるのは正常です。これは、アイドル時の使用率が高いことをターゲットにしています。
ウォッチドッグが警告を発するまで、異常な状態が 15 分間継続する必要があります。これは、正当なリクエストを除外するには十分な長さです。
正常な読み取り値があるとカウンターがリセットされるため、ウォッチドッグは一時的なスパイクではなく、継続的なスタック状態でのみ起動します。
起動すると、診断と 1 行の修正コマンドを含むプッシュ通知が ntfy.sh 経由で携帯電話に送信されます。
送信後、1 時間はそれ以降のアラートが抑制されます。一連の通知ではなく、1 つの通知が必要です。
デザインは2種類からお選びいただけます

注目に値する:
ウォッチドッグは警告を出しますが、モデルを自動的に再起動することはありません。人は行き詰まったランナーが正当な長時間の仕事であるかどうかを数秒で判断できます。自動化によって実際の作業が台無しになることは望ましくありませんでした。
アラートの送信に失敗した場合、ウォッチドッグはイベントを処理済みとしてマークしないため、次回の実行はサイレントになる代わりに再試行されます。
通信会社の友人への余談です。私は当初、AT&T の電子メールから SMS へのゲートウェイを介してアラートを SMS として送信する予定でしたが、そのサービスは 2025 年半ばに閉鎖されました。 ntfy.sh プッシュ通知に切り替えました。これはよりシンプルで、保存された資格情報を必要としないことがわかりました。通信業界は大騒ぎになった。プッシュ通知はキャリアのインフラストラクチャであるべきですが、代わりに OTP サービスです。業界は、さまざまなネットワーク間で通話と SMS をルーティングする方法を理解できましたが、プッシュ通知を異なるネットワーク間でルーティングする方法を理解できなかったでしょうか?あなたの家すべてに痘瘡が発生しています。 IMS は SIP ルーティング以上のものであるべきでした。
スタックした Ollama ランナーが発見され、評価され、リセットされました。先週までは、ハングしたプロセスが 20 時間以上サイレントで実行され、電気代が高くなる危険がありました。システムは、使用量が多い場合は 15 分以内に自己報告するようになりました。モデルが GPU を再度固定すると、マシンの前にいるかどうかに関係なく、修正方法に関する明確なガイダンスを含む電話通知が届きます。ウォッチドッグは無料で実行でき、非常にシンプルです。再起動後も存続する cron ジョブと 2 つのタイムスタンプ ファイル。
cron セットアップは /etc/cron.d/ollama-watchdog にあります (ユーザー crontab ではなく、システム cron ドロップイン)。インストーラーは次のように記述します。
SHELL=/bin/bash
PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
# m h dom mon dow ユーザーコマンド
*/5 * * * * ユーザー /opt/ollama-watchdog/watchdog.sh >> /opt/ollama-watchdog/state/watchdog.log 2>&1
操作上の注意事項:
crontabコマンドはありません

関与した。これはドロップイン ファイルであるため、/etc/cron.d/ollama-watchdog を直接編集または削除して管理します。
自動的に取得されるため、リロードは必要ありません。ファイルはモード 644、所有ルートである必要があります。
再インストール (sudo bash /home/$USER/install-ollama-watchdog.sh) すると毎回このファイルが書き換えられるため、必要に応じてインストーラーを編集します。
スケジュールを永久に変更します。 1 回限りの調整の場合は、cron ファイルを直接編集します。
検出頻度とアラート遅延: 5 分の間隔が解像度を設定します。 15 分間の維持しきい値と組み合わせた、最悪の場合
ランナーウェッジ後のアラートまでの時間は、およそ 15 ～ 20 分です (900 秒を超えるには、最大 3 回の連続した不良サンプルが必要です)。インストールされていることを確認し、動作することを確認します。
cat /etc/cron.d/ollama-watchdog # エントリを確認します
tail -f /opt/ollama-watchdog/state/watchdog.log # 5 分間の決定ごとに監視します
ウォッチドッグのインストール スクリプト:
#!/usr/bin/env bash
#
# インストール-ollam-watchdog.sh
# --------------------------------------
# 電話通知をプッシュする cron 駆動のウォッチドッグをインストールします (ntfy.sh 経由)
# 「オラマ ランナー」が持続的に GPU を固定するとき (
# 「永久回転ファン」障害が発生しました)。
#
# - 5 分ごとのサンプル (cron)。
# - 悪い状態が 15 分以上継続した場合にのみアラートが発行されます。
# - スタックしたままの場合は、レートを 1 時間あたり最大 1 回のプッシュに制限します。
# - 警告のみ: ollam が再起動されることはありません。
#
# トランスポート: ntfy.sh。トピック名はシークレット/資格情報なので、それを作成します
# 意図的に不透明にします。 ntfy アプリまたはで購読してください。
# https://ntfy.sh/<トピック>。アカウント、API キー、.env は必要ありません。
#
# sudo bash install-ollama-watchdog.sh で実行します。
#
set -euo パイプ失敗
BASE=/opt/ollam-watchdog
STATE="$BASE/州"
RUN_USER=_USER_ # cron ジョブはこのユーザーとして実行されます (状態/ログを所有)。ユーザーアカウントを追加してください！
if [ "$(id -u)" -ne 0 ];それから
エッチ

o "このインストーラーは root として実行する必要があります: sudo bash $0" >&2
出口1
フィ
echo "==> $BASE を作成しています"
mkdir -p "$BASE" "$STATE"
# ----------------------------------------------------------------------
# watchdog.conf (設定; ここでしきい値/トピックを編集し、再インストールは不要)
# ----------------------------------------------------------------------
echo "==> $BASE/watchdog.conf の書き込み"
cat > "$BASE/watchdog.conf" <<'CONF_EOF'
# ollama-watchdog 設定 (watchdog.sh から取得)
# --- アラートの送信先 (ntfy.sh) ---
# トピック名は意図的に不透明です。この URL を知っている人は誰でも読むことができ、かつ
# アラートを偽装しても、監視対象については何も明らかにされません。治療してください
# パスワードのようなもの。 ntfy 電話アプリでこれと同じトピックを購読します。
NTFY_URL="https://ntfy.sh/_YOURNTFY.SH_TOPIC"
NTFY_PRIORITY="high" # min|low|default|high|urgent (緊急はサイレントモードを回避できます)
NTFY_TITLE="オラマ警報"
NTFY_TAGS="rotating_light,fire" # 通知に絵文字が表示されます
# --- 検出 ---
GPU_BUSY_FILE="/sys/class/drm/card1/device/gpu_busy_percent"
GPU_BUSY_
[切り捨てられた]

## Original Extract

Fixing a stuck Ollama runner and building a GPU watchdog – blog
Skip to content
blog
Mobile Security Mentorship & Inspiring Kids to work in tech
Fixing a stuck Ollama runner and building a GPU watchdog
If you self-host LLMs, at some point, you’ll likely experience a stuck GPU consuming significant electricity for long periods of time. This is my summary of discovering the problem & then implementing automation that corrects the issue. I want automation to discover & correct these issues. I don’t want to be the first line of defense against unnecessary electricity consumption.
I run a large language model on hardware at home — an Ollama server on a mini PC (AMD Ryzen AI MAX+ 395 with an integrated GPU). Self-hosting gives me privacy, no subscription, and a model that works without an internet connection. The tradeoff is that I’m my own ops team. When something goes wrong, there’s no provider watching for outages.
This weekend, something went wrong.
I noticed the cooling fan on my AI server running at full speed without letting up. I hadn’t been using it- and my family’s use of it is pretty limited. There likely were no active requests for the system. Either the server was compromised and mining cryptocurrency, or a process was hung and burning power for nothing. The system draws up to 85W — leaving it pinned at full power indefinitely would show up on my electricity bill.
I needed to do two things. First, determine whether the machine was doing real work or stuck. Second, build a solution that catches this class of problem automatically: Monitoring for long running fan activity has an unreliable Mean Time To Resolution.
Diagnosis. I checked what the model was doing with ollama ps and found a model stuck in a Stopping... state — it was supposed to unload and free the GPU, but the underlying process never exited. I confirmed the conditions by reading the GPU utilization gauge from the kernel ( /sys/class/drm/card1/device/gpu_busy_percent ): ~89% busy, ~85W. It had been in this state for roughly 20 hours with zero inference requests. The normal shutdown command ( ollama stop ) had no effect. A service restart ( sudo systemctl restart ollama ) cleared it. The root cause is a known Ollama bug where the GPU stays at full utilization with no work to do.
Watchdog. At this point, I confirmed that the Ollama system was hung and needed to be reset. But this problem has happened before! It would likely happen again. I needed to stop relying on my ability to detect unexpected zephyrs emanating off my server. I created a watchdog that alerts me only when this specific failure occurs. The logic of the watchdog needs to be narrow to avoid false positives. The basics of the system are as follows:
A cron job samples the system every 5 minutes with two reads: is a model process loaded , and how busy is the GPU ?
A sample only counts as unhealthy when both conditions are true: a model is loaded and the GPU is at or above 70% utilization . High GPU usage during inference is normal; this targets high usage while idle.
The unhealthy state must hold continuously for 15 minutes before the watchdog alerts. That’s long enough to rule out any legitimate request.
Any healthy reading resets the counter, so the watchdog only fires on a sustained stuck condition, not a transient spike.
When it fires, it sends a push notification to my phone via ntfy.sh with the diagnosis and the one-line fix command.
After sending, it suppresses further alerts for one hour. I want one notification- not a stream of them.
There are two design choices worth noting:
The watchdog alerts but never restarts the model automatically. A person can tell a stuck runner from a legitimate long job in seconds. I didn’t want automation killing real work.
If the alert fails to send, the watchdog doesn’t mark the event as handled, so the next run retries instead of going silent.
An aside for my friends in Telco: I originally planned to send alerts as SMS through a AT&T’s email-to-SMS gateway, but that service was shut down in mid-2025. I switched to ntfy.sh push notifications, which turned out simpler and requires no stored credentials. The telco industry whiffed badly. Push notifications should be carrier infrastructure- but instead it’s an OTP service. The industry could figure out how to route calls and SMS across different networks, but it couldn’t figure out how to route push notifications across them? A pox on all your houses. IMS should have been more than SIP routing!
The stuck Ollama runner was discovered, evaluated and reset. Until last week- I ran the risk that a hung process would run silently for 20+ hours in ways that could raise my electricity bill. The system now self-reports within 15 minutes of high usage. When a model pins the GPU again, I’ll get a phone notification with clear guidance on how to fix it, whether I’m at the machine or not. The watchdog is free to run and is very simple. A cron job and two timestamp files that survive reboots.
The cron setup lives in /etc/cron.d/ollama-watchdog (a system cron drop-in, not a user crontab). The installer writes it as:
SHELL=/bin/bash
PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
# m h dom mon dow user command
*/5 * * * * USER /opt/ollama-watchdog/watchdog.sh >> /opt/ollama-watchdog/state/watchdog.log 2>&1
A few operational notes:
No crontab command involved. Because it’s a drop-in file, you manage it by editing/removing /etc/cron.d/ollama-watchdog directly.
It’s picked up automatically — no reload needed. File should be mode 644, owned root.
Reinstalling (sudo bash /home/$USER/install-ollama-watchdog.sh) rewrites this file each time, so edit the installer if you want to
change the schedule permanently. For a one-off tweak, edit the cron file directly.
Detection cadence vs. alert latency: the 5-min interval sets the resolution. Combined with the 15-min sustain threshold, worst-case
time-to-alert after a runner wedges is roughly 15–20 minutes (you need ~3 consecutive bad samples to cross 900s). Verify it’s installed and watch it work:
cat /etc/cron.d/ollama-watchdog # confirm the entry
tail -f /opt/ollama-watchdog/state/watchdog.log # watch each 5-min decision
Watchdog Installation Script:
#!/usr/bin/env bash
#
# install-ollama-watchdog.sh
# ---------------------------
# Installs a cron-driven watchdog that pushes a phone notification (via ntfy.sh)
# when an `ollama runner` pegs the GPU for a sustained period (the
# "permaspinning fan" failure seen ).
#
# - Samples every 5 min (cron).
# - Alerts only after the bad condition has held continuously for >15 min.
# - Rate-limits to at most one push per hour while it stays stuck.
# - ALERT ONLY: it never restarts ollama for you.
#
# Transport: ntfy.sh. The topic name IS the secret/credential, so make it
# deliberately opaque. Subscribe to it in the ntfy app or at
# https://ntfy.sh/<topic>. No account, no API key, no .env needed.
#
# Run with: sudo bash install-ollama-watchdog.sh
#
set -euo pipefail
BASE=/opt/ollama-watchdog
STATE="$BASE/state"
RUN_USER=_USER_ # cron job runs as this user (owns state/log). Add your user account!
if [ "$(id -u)" -ne 0 ]; then
echo "This installer must be run as root: sudo bash $0" >&2
exit 1
fi
echo "==> Creating $BASE"
mkdir -p "$BASE" "$STATE"
# ---------------------------------------------------------------------------
# watchdog.conf (settings; edit thresholds/topic here, then no reinstall)
# ---------------------------------------------------------------------------
echo "==> Writing $BASE/watchdog.conf"
cat > "$BASE/watchdog.conf" <<'CONF_EOF'
# ollama-watchdog configuration (sourced by watchdog.sh)
# --- where the alert goes (ntfy.sh) ---
# The topic name is opaque on purpose: anyone who knows this URL can read AND
# spoof your alerts, and it reveals nothing about what it monitors. Treat it
# like a password. Subscribe to this same topic in the ntfy phone app.
NTFY_URL="https://ntfy.sh/_YOURNTFY.SH_TOPIC"
NTFY_PRIORITY="high" # min|low|default|high|urgent (urgent can bypass Do Not Disturb)
NTFY_TITLE="Ollama ALERT"
NTFY_TAGS="rotating_light,fire" # emoji shown on the notification
# --- detection ---
GPU_BUSY_FILE="/sys/class/drm/card1/device/gpu_busy_percent"
GPU_BUSY_
[truncated]
