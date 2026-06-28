---
source: "https://github.com/thefullnacho/hestia"
hn_url: "https://news.ycombinator.com/item?id=48707227"
title: "Hestia – a local-first Home Assistant that trusts timers over the LLM"
article_title: "GitHub - thefullnacho/hestia: Local-first, self-hosted home & records assistant: one LLM brain, scoped tools, durable memory. AGPL-3.0. · GitHub"
author: "thefullnacho"
captured_at: "2026-06-28T14:34:04Z"
capture_tool: "hn-digest"
hn_id: 48707227
score: 2
comments: 0
posted_at: "2026-06-28T13:46:44Z"
tags:
  - hacker-news
  - translated
---

# Hestia – a local-first Home Assistant that trusts timers over the LLM

- HN: [48707227](https://news.ycombinator.com/item?id=48707227)
- Source: [github.com](https://github.com/thefullnacho/hestia)
- Score: 2
- Comments: 0
- Posted: 2026-06-28T13:46:44Z

## Translation

タイトル: Hestia – LLM よりもタイマーを信頼するローカルファーストのホーム アシスタント
記事のタイトル: GitHub - thefullnacho/hestia: ローカルファースト、自己ホスト型のホーム & レコード アシスタント: 1 つの LLM ブレイン、範囲指定されたツール、耐久性のあるメモリ。 AGPL-3.0。 · GitHub
説明: ローカルファーストの自己ホスト型ホームおよび記録アシスタント: 1 つの LLM ブレイン、スコープ付きツール、耐久性のあるメモリ。 AGPL-3.0。 - ザフルナチョ/ヘスティア

記事本文:
GitHub - thefullnacho/hestia: ローカルファーストの自己ホスト型ホームおよび記録アシスタント: 1 つの LLM ブレイン、スコープ付きツール、耐久性のあるメモリ。 AGPL-3.0。 · GitHub
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
ザフルナチョ
/
ヘスティア
公共
通知
通知設定を変更するにはサインインする必要があります

s
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット ベンチマーク ベンチマーク ブレイン ブレイン クライアント クライアント デプロイ デプロイ メモリ メモリ .gitignore .gitignore.md ARCHITECTURE.md AUDIT.md AUDIT.md LICENSE LICENSE MARKET.md MARKET.md MEMORY-DESIGN.md MEMORY-DESIGN.md README.md README.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ローカルファーストの、ご自宅用の自己ホスト型アシスタント。 1 つのステートフルな「頭脳」がローカル LLM を実行します。
あなたが所有するハードウェアとそのすべての窓 - 携帯電話、端末、キッチンのマイク、自宅
アシスタント — 同じ脳に話しかけます。クラウド内では何も実行されず、何も公開されません。
インターネットに接続すれば、データが家から出ることはありません。
それはそのアイデアに基づいて構築されています。ほとんどの「家庭用 AI」は最悪の部分をモデルに指摘します
at: スケジュールを記憶し、しきい値を監視し、適切な分にリマインダーを起動します。ヘスティア
逆のことをします。決定的なことなら何でも - 家事が終わる、土は乾いている、ゴミは出る
火曜日 — タイマー、レコード、データベースの行など、愚かで信頼できるものに渡されます。の
LLM は、判断と会話という本当に得意なことだけを行うことになります。の
目標は決して賢い頭脳ではありませんでした。より信頼性の高いものです。 ( ARCHITECTURE.md は
ロングバージョン。 MEMORY-DESIGN.md にはメモリ プランが含まれています。)
ブレイン ( Brain/ ) — をラップする OpenAI 互換エンドポイント ( POST /v1/chat/completions )
エージェント ループを備えたローカル LLM (Ollama、qwen3:14b)。すべてのクライアントは 1 つの方言を話します。
8 つのスコープ付きツール — ホーム (ホーム アシスタントの制御)、メディア (Plex + *arr)、メモリ、
記録、リマインダー、検索、ステータス、天気。意図的にシェルツールはありません。
脳は家の中で活動することはできますが、任意のコマンドを実行することはできません。
M

成長する記憶 — マークダウンのソフトファクトと、あなたの人生の出来事の SQLite 記録
（ペット、庭、野生動物、家事）、そして永続的な事実を提案する背景メモテイカー
黙って書くのではなく、承認してください。
メディア アプライアンス — Plex + *arr スタック + Bazarr 字幕 + qBittorrent の背後にあります
フェールクローズされた VPN キルスイッチ。
音声 — Home Assistant の Assist パイプラインまたはブラウザを通じて話しかけます。
なんとそうではありません。クラウド サービス、他人の API のラッパー、またはあなたが置くべきものなら何でも
公共のインターネット上で。それはあなた自身のボックス上で根無し草で動作し、家に電話をかけることはありません。
⚠️ 実行する前に SECURITY.md を読んでください。脳には何も組み込まれていない
認証し、デバイスを制御できるため、デバイスはプライベート ネットワーク (Tailscale または
LAN）。これは意図的なトレードオフであり、見落としではありません。ドキュメントでは信頼モデルについて説明しています。
Hestia は、 forager_ml と並んで、Forager / Homesteader Labs コンステレーションの一部です。
forager-field-station 、および Homesteader Labs サイト。
フェーズ 0 — リーチ + 頭脳 ✅ — 携帯電話からホーム モデルに話しかけます (詳細は以下を参照)。
フェーズ 1 — メディア アプライアンス ✅ — Plex + qBittorrent + Gluetun VPN キルスイッチ (検証済み) +
*arr オートメーション レイヤー (Prowlarr/Sonarr/Radarr + FlareSolverr + Bazarr 字幕)。フルループ:
検索→ダウンロード（VPN経由）→ハードリンク→Plex。
フェーズ 2 — ハウス (ホーム アシスタント) ✅ — HA の実行。ホームツール経由でアクセス可能なライトとデバイス。
フェーズ 3 — 音声 ✅ — HA の Assist パイプラインとブラウザーの音声ループを通じて Hestia に話しかけます。
フェーズ 4 — 継ぎ目 (記憶 + ツール) ✅ コアは定位置にあり、まだ成長中 — 脳は
上記 8 つのツールと決定論的なスキル インジェクション、および HA を備えたツール呼び出しエージェント
会話エージェントがヘスティアを指しているので、アシストと音声が脳内を通過します (

どれができるか
コントロール HA バック)。また、ノートテイカーによって時間の経過とともに賢くなっていきます (「記憶と学習」を参照)。
次はビジョン（目）です。
Win: 携帯電話からホーム モデルに話しかけます。
Brain ( Brain/ ) は、Ollama に対する OpenAI 互換の薄型プロキシです。すべてのクライアント —
端末、電話、キッチンマイク — 1 つの方言 ( POST /v1/chat/completions ) を話します。
フェーズ 0 では、選択したモデルを強制し、ヘスティアのシステム プロンプト (ペルソナ +
ベンチマーク A/B からの強化された安全ルール）を取得し、応答をストリームバックします。
メモリとツールは、この同じ URL の背後でフェーズ 4 に到達します。
サービス
何
バインド
GPU
ヘスティア・オラマ
オラマ推論エンジン
127.0.0.1:11434 (ローカルホストのみ)
RTX 5080のみ
ヘスティア脳
ヘスティア /v1 プロキシ
0.0.0.0:8730 (Tailscale 経由で到達可能)
—
どちらもユーザーの systemd サービス (root なし) であり、deploy/systemd/ で定義され、
~/.config/systemd/user/ にインストールされます。 Linger が有効になっているため、生き残ることができます
ログアウト/再起動します。 Ollama は 5080 ( CUDA_VISIBLE_DEVICES ) に固定されており、
ベンチマークの判定に従って、フェーズ 3 (ウィスパー/パイパー) では 4060 Ti が無料。
モデル: qwen3:14b (常駐、思考停止) — モデル評価後の現在の選択
( Brain/eval_models.py ; qwen2.5:14b はフォールバックとしてディスク上に保持されます)。 MODEL_EVAL.md を参照してください。
日常的に、deploy/hestiactl ( ~/.local/bin にシンボリックリンクされている) を使用します。コマンドは 1 つだけです。
資産全体については、GPU ボックスから実行します。
hestiactl ステータス # 脳の健康状態 + ローカル ユニット + hl-relay 上のすべてのコンテナ
hestiactl health # raw /health JSON
ヘスティアクトルアップ |ダウン |再起動 X # X: 脳オラマ | arr サービス |プレックス qbit は adguard ... |すべて
hestiactl logs X [-f] #journalctl (ローカル) または docker ログ (リモート)
hestiactl vpn # qBittorrent キルスイッチを確認します
すべては Hestia が管理する部分 (ローカル ユニット + arr スタック) のみをカバーします。コア
コンテナ (AdGuard = ハウス DNS、gluetun、HA) は一度に 1 つずつ制御されます

時間と
停止する前に確認を求めてください。
直接必要な場合の基礎となるコマンド:
# ステータス / ログ
systemctl --user status hestia-ollam hestia-brain
journalctl --user -u hestia-brain -f
# ブレインコードまたはサービスファイルを編集した後に再起動する
systemctl --user daemon-reload # .service を編集した場合のみ
systemctl --user restart hestia-brain
# health (Ollama up + model present?) — 脳はローカルホストではなく、Tailscale IP をバインドします
カール -s 127.0.0.1:8730/health | jq
#話しかけてください
curl -s 127.0.0.1:8730/v1/chat/completions -H ' content-type: application/json ' \
-d ' {"メッセージ":[{"ロール":"ユーザー","コンテンツ":"こんにちはヘスティア"}]} ' | jq -r .choices[0].message.content
deploy/systemd/*.service ファイルを編集した場合は、それを次の場所に再コピーします。
~/.config/systemd/user/ daemon-reload の前。
電話からアクセスする (Tailscale)
Tailscale は root が必要な部分であるため、自動インストールされません。 GPU上で
ボックス:
カール -fsSL https://tailscale.com/install.sh |しー
sudo テールスケールアップ
次に電話で: Tailscale アプリをインストールし、同じテールネットにサインインします。の
これで、任意のアプリから http://<gpu-box-tailscale-name>:8730/v1 で Brain にアクセスできるようになります。
OpenAI を話すもの (それをベース URL として設定します。任意の API キー文字列が機能します — Ollama
無視します)。公共のインターネットには何も公開されません。
脳/
hestia.py # エージェント ループ: /v1/chat/completions + /health、ツール、メモリ、メモ取りフック
config.py # パスの単一ソース + シークレットの読み込み;脳を再配置可能にする
プロンプト.py # SYSTEM_PROMPT — ペルソナ + 強化された安全ルール
records_store.py /memory_store.py # SQLite エンティティ + イベント / マークダウン ソフト ファクト
note_taker.py # 背景「時間の経過とともに賢くなる」エクストラクター
review_notes.py # メモ作成者の提案をレビューおよび促進するための CLI
tools/ # ホーム、メディア、メモリ、記録、リマインダー、検索、ステータス、天気

r (+スキルルーター)
テスト/ # pytest: ストア、ディスパッチ、メモテイカー (実行: uv run --project Brain pytest)
pyproject.toml # deps + dev (pytest) + pytest config (UV 管理、分離された venv)
再配置可能。すべてのパスは config.py 自体の場所から派生するため、移動または
リポジトリを新しいパスに復元する場合、編集は必要ありません。必要に応じて、HESTIA_ROOT がオーバーライドされます。すべて
サービス URL、トークン、およびしきい値は、それらを使用するツールの隣で環境上書き可能なままになります。
フェーズ 1 — メディア アプライアンス (Dell Micro = hl-relay )
勝利: メディア スタックが脳とは独立して実行されます。
これらのほとんどは、Hestia よりも前の Micro にすでに存在していました: Plex ( hl-plex )、
フェールクローズ VPN を使用した Gluetun (Surfshark、OpenVPN、NL) の背後にある qBittorrent
kill-switch に加えて、AdGuard、MQTT、および Home Assistant。キルスイッチが検証されます。
qBittorrent のトラフィックは、ホストの IP ではなく、VPN データセンターの IP 経由で送信されます。やめてください
docker は既存の /opt/home/compose.yml を盲目的に作成します - そのボリュームパス
リテラルの /path/to/... は、実行中のコンテナーが依存するホスト ディレクトリです。
Hestia は、不足していた自動化レイヤーを別個の分離されたスタックとして追加しました
(deploy/media/compose.yml 、 /opt/home/arr/ にデプロイ): Prowlarr (:9696,
インデクサー マネージャー)、Sonarr (:8989、TV)、Radarr (:7878、映画)。すべて到達可能
テイルスケールを超えて。
また、Prowlarr が Cloudflare で保護された領域に到達できるように、FlareSolver (:8191) も追加されました。
インデクサー。Prowlarr インデクサープロキシとして接続されています (タグ Flaresolverr )。
API 経由で接続: ルート フォルダーは既存の Plex ライブラリを指します
( /data/TV 番組、/data/映画 );リモート パス マッピング ( /downloads →
/data/downloads ) により、Sonarr/Radarr を qBittorrent のダウンロードからハードリンクできます。
ライブラリ (インスタント、コピーなし - どちらも /mnt/media の下にある 1 つのファイルシステム)。プロウラー
Sonarr + Radarr ( fullSync ) に接続されています。 5 つの評判の良い公開インデックス作成者が追加されました
(海賊湾、ナブ

en、LimeTorrents、さらに 1337x + EZTV via FlareSolver) と同期
アプリに至るまで。 YTS は意図的に除外しました (ユーザー データを著作権荒らしに提供した歴史)。
qBittorrent は両方の Sonarr (カテゴリ tv-sonarr ) でダウンロード クライアントとして接続されています
と Radarr (レーダー)、テストは OK です。完全なループが機能します: 検索→ダウンロード
VPN → Plex ライブラリにハードリンクします。どちらのアプリも健康に関する警告を報告しません。
⚠️ メディアは現在、Micro の 98 GB ルート ディスク (最大 66 GB の空き) にあります。始めても大丈夫です。
ライブラリが大きくなる前に、専用のディスクまたは NAS を計画してください。
cd /opt/home/arr
ドッカー構成PS
docker compose pull && docker compose up -d # update *arr
フェーズ 4 — 継ぎ目: HA 会話エージェント → ヘスティア
deploy/ha/custom_components/hestia/ はシン カスタム HA 統合です。
各発話をヘスティアに転送する会話エージェント (conversation.hestia)
/v1 と応答を読み上げます。ヘスティアはループ (メモリ + ツール、制御を含む) を所有します。
HA 戻る); HA は単なるインプット + ツールです。これは現実化された建築の要石です。
hl-relay での配線 (このリポジトリにはありません - HA の構成にあります):
統合ファイルは /opt/home/ha_config/custom_components/hestia/ にインストールされます。
構成エントリは http://127.0.0.1:8730/v1/chat/completions をポイントします (Hestia
テールスケール上。 HA コンテナはそれに到達できます)。
好ましい方

[切り捨てられた]

## Original Extract

Local-first, self-hosted home & records assistant: one LLM brain, scoped tools, durable memory. AGPL-3.0. - thefullnacho/hestia

GitHub - thefullnacho/hestia: Local-first, self-hosted home & records assistant: one LLM brain, scoped tools, durable memory. AGPL-3.0. · GitHub
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
thefullnacho
/
hestia
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits benchmarks benchmarks brain brain clients clients deploy deploy memory memory .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md AUDIT.md AUDIT.md LICENSE LICENSE MARKET.md MARKET.md MEMORY-DESIGN.md MEMORY-DESIGN.md README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
A local-first, self-hosted assistant for your home. One stateful "brain" runs a local LLM on
hardware you own, and every window into it — your phone, a terminal, the kitchen mic, Home
Assistant — talks to that same brain. Nothing runs in the cloud, nothing is exposed to the
internet, and your data never leaves the house.
The idea it's built on. Most "AI for the home" points the model at the things it's worst
at: remembering a schedule, watching a threshold, firing a reminder at the right minute. Hestia
does the opposite. Anything deterministic — a chore is due, the soil is dry, the trash goes out
Tuesday — is handed to something dumb and reliable: a timer, a record, a row in a database. The
LLM is left to do the one thing it's genuinely good at, which is judgment and conversation. The
goal was never a smarter brain. It's a more reliable one. ( ARCHITECTURE.md is
the long version; MEMORY-DESIGN.md covers the memory plan.)
A brain ( brain/ ) — an OpenAI-compatible endpoint ( POST /v1/chat/completions ) wrapping a
local LLM (Ollama, qwen3:14b ) with an agent loop. Every client speaks one dialect.
Eight scoped tools — home (control Home Assistant), media (Plex + *arr), memory ,
records , reminder , search , status , weather . There is deliberately no shell tool :
the brain can act in your house but cannot run arbitrary commands.
Memory that grows — markdown soft-facts plus a SQLite record of the things in your life
(pets, garden, wildlife, chores), and a background note-taker that proposes durable facts for
you to approve rather than writing them silently.
A media appliance — Plex + the *arr stack + Bazarr subtitles + qBittorrent behind a
fail-closed VPN kill-switch.
Voice — talk to it through Home Assistant's Assist pipeline or the browser.
What it isn't. A cloud service, a wrapper around someone else's API, or anything you should put
on the public internet. It runs rootless on your own box and never phones home.
⚠️ Read SECURITY.md before running it. The brain has no built-in
authentication and can control your devices, so it must stay on a private network (Tailscale or
LAN). That's a deliberate trade-off, not an oversight — the doc explains the trust model.
Hestia is part of the Forager / Homesteader Labs constellation, alongside forager_ml ,
forager-field-station , and the Homesteader Labs site.
Phase 0 — Reach + brain ✅ — talk to your home model from your phone (details below).
Phase 1 — Media appliance ✅ — Plex + qBittorrent + gluetun VPN kill-switch (verified) + the
*arr automation layer (Prowlarr/Sonarr/Radarr + FlareSolverr + Bazarr subtitles). Full loop:
search → download (via VPN) → hardlink → Plex.
Phase 2 — House (Home Assistant) ✅ — HA running; lights and devices reachable via the home tool.
Phase 3 — Voice ✅ — speak to Hestia through HA's Assist pipeline and a browser voice loop.
Phase 4 — The seam (memory + tools) ✅ core in place, still growing — the brain is a
tool-calling agent with the eight tools above plus deterministic skill injection, and HA's
conversation agent points at Hestia , so Assist and voice route through the brain (which can
control HA back). It also gets smarter over time via the note-taker (see Memory & learning ).
Next: vision (Eyes).
Win: talk to your home model from your phone.
The brain ( brain/ ) is a thin OpenAI-compatible proxy onto Ollama. Every client —
terminal, phone, kitchen mic — speaks one dialect ( POST /v1/chat/completions ).
In Phase 0 it forces the chosen model, injects Hestia's system prompt (persona +
the hardened safety rules from the benchmark A/B), and streams the reply back.
Memory and tools land in Phase 4 behind this same URL.
Service
What
Bind
GPU
hestia-ollama
Ollama inference engine
127.0.0.1:11434 (localhost only)
RTX 5080 only
hestia-brain
Hestia /v1 proxy
0.0.0.0:8730 (reachable over Tailscale)
—
Both are user systemd services (no root), defined in deploy/systemd/ and
installed into ~/.config/systemd/user/ . Linger is enabled, so they survive
logout/reboot. Ollama is pinned to the 5080 ( CUDA_VISIBLE_DEVICES ), leaving the
4060 Ti free for Phase 3 (Whisper/Piper) per the benchmark verdict.
Model: qwen3:14b (resident, thinking off) — the current pick after the model eval
( brain/eval_models.py ; qwen2.5:14b kept on disk as a fallback). See MODEL_EVAL.md .
Day to day, use deploy/hestiactl (symlinked into ~/.local/bin ) — one command
for the whole estate, run from the GPU box:
hestiactl status # brain health + local units + every container on hl-relay
hestiactl health # raw /health JSON
hestiactl up | down | restart X # X: brain ollama | arr services | plex qbit ha adguard ... | all
hestiactl logs X [-f] # journalctl (local) or docker logs (remote)
hestiactl vpn # verify the qBittorrent kill-switch
all covers only the Hestia-managed pieces (local units + arr stack); core
containers (AdGuard = house DNS, gluetun, HA) are controlled one at a time and
ask for confirmation before stopping.
The underlying commands, for when you need them directly:
# status / logs
systemctl --user status hestia-ollama hestia-brain
journalctl --user -u hestia-brain -f
# restart after editing brain code or a service file
systemctl --user daemon-reload # only if you edited a .service
systemctl --user restart hestia-brain
# health (Ollama up + model present?) — brain binds the Tailscale IP, not localhost
curl -s 127.0.0.1:8730/health | jq
# talk to it
curl -s 127.0.0.1:8730/v1/chat/completions -H ' content-type: application/json ' \
-d ' {"messages":[{"role":"user","content":"hello Hestia"}]} ' | jq -r .choices[0].message.content
If you edit a deploy/systemd/*.service file, re-copy it into
~/.config/systemd/user/ before daemon-reload .
Reach it from the phone (Tailscale)
Tailscale is the one piece that needs root, so it isn't auto-installed. On the GPU
box:
curl -fsSL https://tailscale.com/install.sh | sh
sudo tailscale up
Then on the phone: install the Tailscale app, sign in to the same tailnet. The
brain is then reachable at http://<gpu-box-tailscale-name>:8730/v1 from any app
that speaks OpenAI (set that as the base URL; any API key string works — Ollama
ignores it). Nothing is exposed to the public internet.
brain/
hestia.py # the agent loop: /v1/chat/completions + /health, tools, memory, note-taker hook
config.py # single source of paths + secret loading; makes the brain relocatable
prompt.py # SYSTEM_PROMPT — persona + hardened safety rules
records_store.py / memory_store.py # SQLite entities+events / markdown soft facts
note_taker.py # background "gets smarter over time" extractor
review_notes.py # CLI to review + promote the note-taker's proposals
tools/ # home, media, memory, records, reminder, search, status, weather (+ skill router)
tests/ # pytest: stores, dispatch, note-taker (run: uv run --project brain pytest)
pyproject.toml # deps + dev (pytest) + pytest config (uv-managed, isolated venv)
Relocatable. Every path derives from config.py 's own location, so moving or
restoring the repo to a new path needs no edits; HESTIA_ROOT overrides if needed. All
service URLs, tokens, and thresholds stay env-overridable next to the tools that use them.
Phase 1 — Media appliance (Dell Micro = hl-relay )
Win: the media stack runs, independent of the brain.
Most of this already existed on the Micro before Hestia: Plex ( hl-plex ),
qBittorrent behind gluetun (Surfshark, OpenVPN, NL) with a fail-closed VPN
kill-switch , plus AdGuard, MQTT, and Home Assistant. The kill-switch is verified:
qBittorrent's traffic egresses via the VPN datacenter IP, not the host's. Don't
docker compose up the existing /opt/home/compose.yml blindly — its volume paths
are literal /path/to/... host dirs that the running containers depend on.
Hestia added the missing automation layer as a separate, isolated stack
( deploy/media/compose.yml , deployed to /opt/home/arr/ ): Prowlarr (:9696,
indexer manager), Sonarr (:8989, TV), Radarr (:7878, movies). All reachable
over Tailscale.
Also added FlareSolverr (:8191) so Prowlarr can reach Cloudflare-protected
indexers, wired as a Prowlarr indexer-proxy (tag flaresolverr ).
Wired via API: root folders point at the existing Plex library
( /data/TV Shows , /data/Movies ); a remote-path mapping ( /downloads →
/data/downloads ) lets Sonarr/Radarr hardlink from qBittorrent's downloads into
the library (instant, no copy — both are one filesystem under /mnt/media ); Prowlarr
is connected to Sonarr + Radarr ( fullSync ). Five reputable public indexers added
(The Pirate Bay, Knaben, LimeTorrents, plus 1337x + EZTV via FlareSolverr) and synced
down to the apps. YTS deliberately excluded (history of feeding user data to copyright trolls).
qBittorrent is wired as the download client in both Sonarr (category tv-sonarr )
and Radarr ( radarr ), tested OK. The full loop works: search → download through the
VPN → hardlink into the Plex library. Both apps report no health warnings.
⚠️ Media currently lives on the Micro's 98 GB root disk (~66 GB free). Fine to start;
plan a dedicated disk or NAS before the library grows.
cd /opt/home/arr
docker compose ps
docker compose pull && docker compose up -d # update *arr
Phase 4 — the seam: HA conversation agent → Hestia
deploy/ha/custom_components/hestia/ is a thin custom HA integration: it registers a
conversation agent ( conversation.hestia ) that forwards each utterance to Hestia's
/v1 and speaks the reply. Hestia owns the loop (memory + tools, incl. controlling
HA back); HA is just input + a tool. This is the architecture's keystone made real.
Wiring on hl-relay (not in this repo — lives in HA's config):
Integration files installed to /opt/home/ha_config/custom_components/hestia/ .
A config entry points it at http://127.0.0.1:8730/v1/chat/completions (Hestia
over Tailscale; the HA container can reach it).
The preferr

[truncated]
