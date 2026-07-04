---
source: "https://guilhermefrj.medium.com/i-built-a-local-chatgpt-killer-on-a-single-rtx-5080-heres-everything-that-went-wrong-and-right-38343c516451"
hn_url: "https://news.ycombinator.com/item?id=48789027"
title: "Show HN: Using Wake-on-LAN for an AI Project"
article_title: "I Built a Local ChatGPT-Killer on a Single RTX 5080. Here’s Everything That Went Wrong (and Right). | by Guilherme Souza | Jul, 2026 | Medium"
author: "guilhermef"
captured_at: "2026-07-04T21:49:13Z"
capture_tool: "hn-digest"
hn_id: 48789027
score: 1
comments: 0
posted_at: "2026-07-04T20:57:20Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Using Wake-on-LAN for an AI Project

- HN: [48789027](https://news.ycombinator.com/item?id=48789027)
- Source: [guilhermefrj.medium.com](https://guilhermefrj.medium.com/i-built-a-local-chatgpt-killer-on-a-single-rtx-5080-heres-everything-that-went-wrong-and-right-38343c516451)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T20:57:20Z

## Translation

タイトル: HN の表示: AI プロジェクトでの Wake-on-LAN の使用
記事のタイトル: 1 台の RTX 5080 でローカル ChatGPT キラーを構築しました。間違っていた (そして正しかった) ことはすべてここにあります。 |ギリェルメ・ソウザ著 | 2026 年 7 月 |中
説明: 単一の RTX 5080 上にローカル ChatGPT キラーを構築しました。問題があった (および正しかった) ことはすべてここにあります。週末のプロジェクトが、CUDA コンパイラ エラー、レンガ化された問題を 1 週間にわたって徹底的に調べることになりました。
HN テキスト: ついに、Wake-on-LAN を使用して AI とエネルギー代を節約する口実ができました。

記事本文:
単一の RTX 5080 上にローカル ChatGPT キラーを構築しました。ここでは、問題があった (および正しかった) ことをすべて示します。
週末のプロジェクトは、CUDA コンパイラー エラー、ブリックされた GPU、および驚くほど高速な専門家混合モデルを介して 1 週間にわたる詳細な調査に変わりました。
私は、プライベートな自己ホスト型 AI 推論サーバーを望んでいました。それは、自分のハードウェア上で実行され、世界中のどこにいても携帯電話からアクセスでき、トークンごとの課金やデータがネットワークから流出しないものです。単一のコンシューマ GPU (16 GB の VRAM を搭載した RTX 5080) を Ryzen 7 7800X3D および 32 GB の DDR5 と組み合わせただけです。
計画は紙の上では単純に見えました。Ubuntu をインストールし、量子化モデルを取得し、 llama-server を実行するだけで、午後には完了しました。
それは午後には行われませんでした。
ラウンド 1: とにかく何かを実行する
ヘッドレス推論ボックスの場合、Ubuntu Server は抵抗が最も少ないパスです。Ollama のインストール スクリプトはそれを直接ターゲットにしており、NVIDIA の CUDA パッケージは 1 回の適切なインストールで済み、コミュニティのサポートは非常に強力です。結局、ゲーム Windows と信頼できるワーカー Kubuntu と並行してデュアルブートすることになりました。
第 2 ラウンド: CUDA コンパイラ戦争
OS を分類したら、可能な限り高速な推論スタックを探しました。ストックのllama.cppは良いです。しかし、私は TurboQuant について読んだことがあります。TurboQuant は、精度損失ゼロで 3 ビット量子化を達成し、アテンション キャッシュで最大 8 倍の高速化を達成する、極端な KV キャッシュ圧縮のための Google Research の技術です。すでに誰かが llama.cpp をフォークして統合していました。
そのフォークを構築すると、私の最初の本当の壁が表面化しました。
エラー: #error -- サポートされていない GNU バージョンです。 gcc バージョン 14 以降はサポートされていません。私のディストリビューションには GCC 15 が同梱されていました。CUDA 12.8 は GCC 14 までしかサポートしていません。修正は CUDA 13.3 にアップグレードすることで、ネイティブ GCC 15 サポートが追加されました。最先端のコンパイラ ツールチェーンのバージョンを思い出してください。

n は GPU ドライバーと同じくらい重要です。
コンパイルが完了すると、 TurboQuant KV キャッシュ タイプとして Turbo2 、 Turbo3 、 Turbo4 が利用可能になりました。 -ctkturbo3 -ctvturbo3 を適用すると、KV キャッシュが fp16 と比較して約 4.3 倍圧縮されました。これは、同じ 16GB カード上で 32K コンテキスト ウィンドウを実行する場合と 262K コンテキスト ウィンドウを実行する場合の違いです。
ラウンド 3: すべてのギガバイトを争う
16 GB のコンシューマ GPU で 270 億パラメータのモデルを実行する場合の不快な真実は次のとおりです。モデルの重みだけで VRAM 予算のほぼ全体を消費します。 Q3_K_XL 量子化で Qwen3.6–27B モデルをロードしようとした最初の試みでは、残されたヘッドルームは 2GB 未満でした。ビジョン エンコーダはおろか、意味のあるコンテキスト ウィンドウを表示するには十分ではありませんでした。
すべてのロックを解除したトリック: -nkvo (一見、欺瞞的な名前のフラグです。これは「KV オフロードなし」を表しますが、実際に行うことは、KV キャッシュを VRAM ではなくシステム RAM に保持し、モデルの重みは GPU に保持することです)。 TurboQuant 圧縮と組み合わせることで、持っていない VRAM ではなく、すでに持っていた RAM を使用してモデルの完全なネイティブ 262K コンテキスト ウィンドウを実行できるようになりました。
すべてのトークンで PCIe 経由で KV キャッシュをプルする場合の生成速度コストは実際のもので、スループット ヒットは約 40% でしたが、このハードウェアでロング コンテキストと 27B クラス モデルの両方を取得するには、これが唯一の方法でした。
第 4 ラウンド: 専門家の混合が数学を変える
高密度 27B モデルは、すべてのパラメータ、すべてのトークンを毎回処理します。専門家混合 (MoE) モデルではそうではありません。 Qwen3.6–35B-A3B には合計 350 億のパラメータがありますが、トークンごとにそのうちの約 30 億だけがアクティブになります。ルーターはどの「エキスパート」が関連するかを選択し、それらのみを実行します。
これは、消費者向けハードウェアにとって非常に重要であることがわかります。 --cpu-moe (アテンションとルーティングを GPU 上に維持しながら、すべてのエキスパートの重みをシステム RAM に保持するフラグ) を使用すると、私の VRA

モデル全体の M 使用量は 3GB 未満に低下しました。また、トークンあたりの計算を行うのはネットワークの一部だけであるため、スループットはほぼ 2 倍になり、高密度 27B モデルの約 17 トークン/秒から、MoE バリアントではほぼ 40 トークン/秒に増加しました。
コーディング、Web 検索の要約、チャットなど、テキストのみの場合は、MoE モデルが明らかなデフォルトになりました。
ラウンド 5: どこからでもアクセスできるようにする
ローカル推論サーバーは、実際にアクセスできる場合にのみ役に立ちます。私が設定したのは：
Tailscale 、パブリック インターネットへのポートを開くことなく、世界中のどこにいても携帯電話からサーバーに到達するプライベート メッシュ ネットワーク用
OpenResty (nginx と Lua スクリプト) を Raspberry Pi 上のリバース プロキシとして使用し、複数のサービス (推論サーバー、Web UI) をすべて単一ポート上のホスト名ベースの仮想ルーティングを通じてルーティングします。
Wake-on-LAN は、大規模サーバーがスリープ状態のときにリクエストがプロキシに到達するたびに、Lua スクリプト経由で自動的にトリガーされるため、実際に必要な場合にのみ電力を消費します。
アイドル シャットダウン ウォッチドッグは、アクティブな SSH セッションとライブ推論リクエストの両方をチェックし、実際に非アクティブ状態が 1 時間続くと自動的にマシンの電源をオフにします。
Wake-on-LAN を確実に動作させるには、独自の回り道が必要でした。OpenResty の組み込み UDP ソケット API は、セキュリティ上の理由からブロードキャスト パケット (WoL が依存するもの) をサイレントにブロックし、標準の ngx.socket.udp() cosocket API には SO_BROADCAST を設定する方法がありません。この修正は、LuaJIT の FFI 層に組み込まれ、生の C ソケット API を直接呼び出すことでした。これは、2 行の問題に対する、わずかに不安定ではあるものの、満足のいく解決策です。
最終的なセットアップ: 35B パラメーターの混合専門家モデル。GPU 上のアテンション レイヤーのみを使用して、ほぼ完全にシステム RAM を使い果たし、正式に 16 GB のメモリが「のみ」搭載されているカードで 39 トークン/秒を提供します。 262K-to

ken コンテキスト ウィンドウ。積極的な KV キャッシュ圧縮を利用します。自動ウェイクオンデマンドとアイドル状態のシャットダウンにより、実際に使用しているときにのみ実際の電力が消費されます。
これには H100 を購入する必要はありませんでした。これには多くの忍耐が必要でした。GPU の完全なデッド状態は再起動だけで解決できました。また、現時点でローカル AI での本当のパフォーマンスが勝つのは、より大型のハードウェアを購入することではありません。すでに所有しているハードウェアをより賢く利用することが重要です。つまり、高密度よりも疎な計算、帯域幅が許す場合は VRAM よりも RAM、精度のトレードオフが無視できる場合は圧縮です。
同様のものを構築することを考えている場合は、物理的な構築よりもソフトウェア スタックに多くの時間を割り当ててください。 GPU は簡単な部分です。
もしあなたが同じようなことに取り組んでいるのであれば、何がうまくいったのか（あるいは壊滅的にうまくいかなかったのか）をぜひ聞きたいです。
#!/bin/bash
# /opt/アイドルシャットダウン.sh
#
# IDLE_THRESHOLD 秒間アクティビティがなかった場合、マシンをシャットダウンします。
# 「アクティビティ」とは、次のいずれかを意味します。
# - アクティブな SSH 接続 (ポート 22 で確立された TCP セッション)
# - アクティブな llama-server 世代 (どのスロットでも is_processing = true)
# - ログインしているローカル ユーザー (物理コンソール/グラフィカル セッション)
# - 進行中の dpkg/apt 操作 (パッケージ ロックが保持されている)
#
# ローカル アクティビティは入力ベースではなくプレゼンス ベースです。ログイン ユーザーはカウントされます。
# アイドル状態であっても、セッションの全期間中アクティブとして扱われます。これ
# コンソールでの無人タスク (長時間にわたる apt アップグレードなど) が実行されないことを意味します
# 誰もキーボードに触れていないという理由だけでシャットダウンをトリガーします。
#
# dpkg チェックは追加のセーフティ ネットです。これにより、マシンが一定時間起動したままになります。
# パッケージ操作の継続時間 (SSH セッション経由で開始された場合も含む)
# それ以来、またはウナッテによって減少しました

nded-upgrades タイマー。
#
# いずれかの信号がアクティブな場合、アイドル タイマーがリセットされます。
IDLE_THRESHOLD=3600 # 1 時間を秒単位で表す
CHECK_INTERVAL=60 # 60 秒ごとにチェックします
IDLE_FILE="/tmp/idle_since"
LLAMA_URL="http://localhost:8080/スロット"
check_llama_idle() {
ローカルアクティブ
active=$(curl -s "$LLAMA_URL" 2>/dev/null | python3 -c "
json、sysをインポート
試してみてください:
スロット = json.load(sys.stdin)
active = sum(s.get('is_processing') == True の場合、スロット内の s に 1)
印刷(アクティブ)
例外を除く:
印刷(0)
" 2>/dev/null)
エコー「${アクティブ:-0}」
}
check_ssh_connections() {
# -H はヘッダー行を抑制します。 「状態が確立された」場合、状態列
# が削除されるため、ヘッダーは (「State」ではなく)「Recv-Q」で始まり、古いものは
# 'grep -vc "^State"' はヘッダーを接続としてカウントしました。 wc -l の
# ヘッダーのない出力では実際のカウントが得られます。
ss -Htn 状態が確立されました '( dport = :22 または Sport = :22 )' 2>/dev/null |トイレ -l
}
check_local_user() {
# ローカルの非グリーター ログイン セッション (物理コンソール/グラフィカル ユーザー) をカウントします。
# フィルター:
# Class=user -> ディスプレイマネージャーのグリーター (SDDM など) を除外します。
# それ以外の場合は永続セッションが表示され、Wou
[切り捨てられた]
-- /etc/nginx/lua/wol.lua
ローカル ffi = require("ffi")
ffi.cdef[[
int ソケット (int ドメイン、int タイプ、int プロトコル);
int setsockopt(int sockfd, int level, int optname, const void *optval, unsigned int optlen);
int sendto(int sockfd, const void *buf, size_t len, int flags, const void *dest_addr, unsigned int addrlen);
int close(int fd);
unsigned short htons(unsigned short hostshort);
int inet_pton(int af, const char *src, void *dst);
構造体 sockaddr_in {
短い sin_family;
unsigned short sin_port;
unsigned int sin_addr;
char sin_zero[8];
};
]]
ローカル AF_INET = 2
ローカル SOCK_DGRAM = 2
ローカル SOL_SOCKET = 1
ローカル SO_BROADCAST = 6
ローカル _M = {}
function _M.send(mac_hex, ブロードキャスト_ip, c

オールダウン）
クールダウン = クールダウンまたは 20
ローカルwol_state = ngx.shared.wol_state
ローカル last_wake = wol_state:get("last_wake_" .. mac_hex)
ローカルの現在 = ngx.now()
last_wake かつ (now - last_wake) < クールダウンの場合
falseを返す
終わり
ローカル ソケット = ffi.C.socket(AF_INET, SOCK_DGRAM, 0)
靴下 < 0 の場合
ngx.log(ngx.ERR、「WoL: ソケットの作成に失敗しました」)
falseを返す
終わり
ローカルブロードキャスト有効 = ffi.new("int[1]", 1)
ffi.C.setsockopt(sock, SOL_SOCKET, SO_BROADCAST, Broadcasting_enable, ffi.sizeof("int"))
ローカルアドレス = ffi.new("struct sockaddr_in")
addr.sin_family = AF_INET
addr.sin_port = ffi.C.htons(9)
ローカル ip_ptr = ffi.new("unsigned int[1]")
ffi.C.inet_pton(AF_INET、ブロードキャスト_ip、ip_ptr)
addr.sin_addr = ip_ptr[0]
ローカル mac_bytes = ""
for i = 1, #mac_hex, 2 do
mac_bytes = mac_bytes .. string.char(tonumber(mac_hex:sub(i, i + 1), 16))
終わり
ローカルマジック = string.rep("\255", 6) .. string.rep(mac_bytes, 16)
ローカル送信 = ffi.C.sendto(sock, magic, #magic, 0, addr, ffi.sizeof(addr))
ffi.C.close(ソックス)
0 未満で送信された場合
ngx.log(ngx.ERR, "WoL: 送信に失敗しました")
falseを返す
終わり
wol_state:set("last_wake_" .. mac_hex、現在)
ngx.log(ngx.INFO, "WoL: MAC にマジック パケットが送信されました ", mac_hex, " (", 送信されました, " バイト)")
trueを返す
終わり
return _M Nginx サーバー構成:
上流推論サーバー {
サーバー192.168.50.142:8080 max_fails=1fail_timeout=5s;
キープアライブ 32;
}
サーバー {
8080を聞いてください。
server_name 推論 ~^inference.*$;
client_max_body_size 100M;
キープアライブタイムアウト 10;
limit_req ゾーン = 推論バースト = 20 遅延なし;
add_header X-Frame-Options 常に SAMEORIGIN。
add_header X-Content-Type-Options nosniff は常に;
# 手動ウェイクトリガー
場所 = /wake {
content_by_lua_block {
ローカル wol = require("wol")
ローカル送信 = wol.send("ac3da132635b", "192.168.50.255", 20)
ngx.header["コンテンツタイプ"] = "テキスト/プレーン"
送信された場合
ngx.say("WoL p

承認が送信されました」)
それ以外の場合
ngx.say("WoL スキップされました (クールダウン有効)")
終わり
}
}
# メインプロキシ — すべてのリクエストに対して先制的に WOL を起動し、クールダウンでスパムを防止します
場所 / {
access_by_lua_block {
ローカル wol = require("wol")
wol.send("ac3da132635b", "192.168.50.255", 20)
}
proxy_pass http://inference_server;
proxy_http_バージョン 1.1;
proxy_set_header ホスト $host;
proxy_set_header X-Real-IP $remote_addr;
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
proxy_set_header X-Forwarded-Proto $scheme;
proxy_set_header 接続 "";
proxy_connect_timeout 3 秒;
proxy_send_timeout 300 秒;
proxy_read_timeout 300 秒;
プロキシバッファリングはオフです。
proxy_next_upstream エラー タイムアウト。
エラーページ 502 504 = @wake_and_wait;
}
場所 @wake_and_wait {
デフォルトタイプ テキスト/プレーン;
add_header 再試行 - 15 以降は常に;
return 503 "推論サーバーが起動中です。15 秒以内に再試行してください\n";
}
場所 /nginx-health {
アクセス_ログオフ;
200「健康です\n」を返します。
add_header Content-Type text/plain;
}
AI エージェント Ai --
情熱的な DevOps/SRE。 https://github.com/guilhermef

## Original Extract

I Built a Local ChatGPT-Killer on a Single RTX 5080. Here’s Everything That Went Wrong (and Right). A weekend project that turned into a week-long deep dive through CUDA compiler errors, a bricked …

I finally have an excuse to use Wake-on-LAN and save on my AI and energy bill.

I Built a Local ChatGPT-Killer on a Single RTX 5080. Here’s Everything That Went Wrong (and Right).
A weekend project that turned into a week-long deep dive through CUDA compiler errors, a bricked GPU, and a surprisingly fast Mixture-of-Experts model.
I wanted a private, self-hosted AI inference server: something running on my own hardware, reachable from my phone anywhere in the world, with no per-token billing and no data leaving my network. Just a single consumer GPU — an RTX 5080 with 16GB of VRAM — paired with a Ryzen 7 7800X3D and 32GB of DDR5.
The plan looked simple on paper: install Ubuntu, grab a quantized model, run llama-server , done in an afternoon.
It was not done in an afternoon.
Round One: Just Get Something Running
For a headless inference box, Ubuntu Server is the path of least resistance — Ollama’s install script targets it directly, NVIDIA’s CUDA packages are one apt install away, and the community support is enormous. I ended up dual-booting it alongside my Gaming Windows and my trusted worker Kubuntu.
Round Two: The CUDA Compiler Wars
Once the OS was sorted, I went looking for the fastest possible inference stack. Stock llama.cpp is good. But I'd read about TurboQuant — a Google Research technique for extreme KV-cache compression, achieving 3-bit quantization with zero accuracy loss and up to 8x speedup on the attention cache. Someone had already forked llama.cpp to integrate it.
Building that fork surfaced my first real wall:
error: #error -- unsupported GNU version! gcc versions later than 14 are not supported! My distro shipped GCC 15. CUDA 12.8 only supports up to GCC 14. The fix was upgrading to CUDA 13.3, which added native GCC 15 support — a good reminder that on the bleeding edge, your compiler toolchain version matters as much as your GPU driver.
Once compiled, I had TurboQuant KV cache types available: turbo2 , turbo3 , turbo4 . Applying -ctk turbo3 -ctv turbo3 compressed my KV cache by roughly 4.3x compared to fp16, which is the difference between running a 32K context window and running a 262K one on the same 16GB card.
Round Three: Fighting for Every Gigabyte
Here’s the uncomfortable truth about running a 27-billion-parameter model on a 16GB consumer GPU: the model weights alone eat almost the entire VRAM budget. My first attempts at loading a Qwen3.6–27B model at Q3_K_XL quantization left less than 2GB of headroom — not nearly enough for a meaningful context window, let alone a vision encoder.
The trick that unlocked everything: -nkvo (a deceptively-named flag — it stands for "no KV offload," but what it actually does is keep the KV cache in system RAM instead of VRAM, while the model weights stay on the GPU). Combined with TurboQuant compression, this let me run the model's full native 262K context window using RAM I already had, rather than VRAM I didn't.
The generation speed cost of pulling KV cache over PCIe on every token was real — roughly a 40% throughput hit — but it was the only way to get both long context and a 27B-class model on this hardware.
Round Four: Mixture-of-Experts Changes the Math
Dense 27B models process every parameter, for every token, every time. Mixture-of-Experts (MoE) models don’t. Qwen3.6–35B-A3B has 35 billion total parameters, but only activates about 3 billion of them per token — a router picks which “experts” are relevant and only runs those.
This turns out to matter enormously for consumer hardware. With --cpu-moe (a flag that keeps all expert weights in system RAM, while attention and routing stay on GPU), my VRAM usage for the entire model dropped to under 3GB. And because only a fraction of the network computes per token, throughput nearly doubled : from roughly 17 tokens/second on the dense 27B model to almost 40 tokens/second on the MoE variant.
For anything text-only — coding, web search summarization, chat — this made the MoE model the obvious default.
Round Five: Making It Reachable From Anywhere
A local inference server is only useful if you can actually reach it. I set up:
Tailscale , for a private mesh network reaching the server from my phone anywhere in the world, without opening ports to the public internet
OpenResty (nginx plus Lua scripting) as a reverse proxy on a Raspberry Pi, routing multiple services — The inference server, a web UI — all through hostname-based virtual routing on a single port
Wake-on-LAN , triggered automatically via a Lua script whenever a request hits the proxy while the big server is asleep, so it only draws power when actually needed
An idle-shutdown watchdog , checking both active SSH sessions and live inference requests, powering the machine down automatically after an hour of true inactivity
Getting Wake-on-LAN working reliably took its own detour: OpenResty’s built-in UDP socket API silently blocks broadcast packets (the ones WoL depends on) for security reasons, and the standard ngx.socket.udp() cosocket API has no way to set SO_BROADCAST . The fix was dropping into LuaJIT's FFI layer to call the raw C socket API directly — a satisfying, if slightly unhinged, solution to a two-line problem.
The final setup: a 35B-parameter Mixture-of-Experts model, running almost entirely out of system RAM with only its attention layers on GPU, delivering 39 tokens/second on a card that officially has “only” 16GB of memory. A 262K-token context window, courtesy of aggressive KV cache compression. Automatic wake-on-demand and idle shutdown so the whole thing only draws real power when I’m actually using it.
None of this required buying an H100. It required a lot of patience, one genuinely dead-in-the-water GPU state that only a reboot could fix, and the slow realization that the real performance wins in local AI right now aren’t about buying bigger hardware — they’re about making smarter use of the hardware you already have: sparse computation over dense, RAM over VRAM where bandwidth allows, and compression wherever the accuracy trade-off is negligible.
If you’re thinking about building something similar: budget more time for the software stack than the physical build. The GPU is the easy part.
If you’re working on something similar, I’d love to hear what’s worked (or catastrophically not worked) for you.
#!/bin/bash
# /opt/idle-shutdown.sh
#
# Shuts down the machine after IDLE_THRESHOLD seconds of no activity.
# "Activity" means any of:
# - An active SSH connection (established TCP session on port 22)
# - An active llama-server generation (is_processing = true on any slot)
# - A logged-in local user (physical console/graphical session)
# - An in-progress dpkg/apt operation (package lock held)
#
# Local activity is presence-based, not input-based: a logged-in user counts
# as active for the entire duration of their session, even while idle. This
# means an unattended task at the console (e.g. a long apt upgrade) will NOT
# trigger a shutdown just because nobody is touching the keyboard.
#
# The dpkg check is an additional safety net: it keeps the machine awake for
# the duration of any package operation, even one started over an SSH session
# that has since dropped, or by the unattended-upgrades timer.
#
# If any signal is active, the idle timer resets.
IDLE_THRESHOLD=3600 # 1 hour in seconds
CHECK_INTERVAL=60 # check every 60 seconds
IDLE_FILE="/tmp/idle_since"
LLAMA_URL="http://localhost:8080/slots"
check_llama_idle() {
local active
active=$(curl -s "$LLAMA_URL" 2>/dev/null | python3 -c "
import json, sys
try:
slots = json.load(sys.stdin)
active = sum(1 for s in slots if s.get('is_processing') == True)
print(active)
except Exception:
print(0)
" 2>/dev/null)
echo "${active:-0}"
}
check_ssh_connections() {
# -H suppresses the header line. With 'state established' the State column
# is dropped, so the header starts with "Recv-Q" (not "State") and the old
# 'grep -vc "^State"' counted the header as a connection. wc -l on the
# header-less output gives the true count.
ss -Htn state established '( dport = :22 or sport = :22 )' 2>/dev/null | wc -l
}
check_local_user() {
# Count local, non-greeter login sessions (physical console/graphical user).
# Filters:
# Class=user -> excludes the display-manager greeter (e.g. SDDM), which
# otherwise shows a permanent session and wou
[truncated]
-- /etc/nginx/lua/wol.lua
local ffi = require("ffi")
ffi.cdef[[
int socket(int domain, int type, int protocol);
int setsockopt(int sockfd, int level, int optname, const void *optval, unsigned int optlen);
int sendto(int sockfd, const void *buf, size_t len, int flags, const void *dest_addr, unsigned int addrlen);
int close(int fd);
unsigned short htons(unsigned short hostshort);
int inet_pton(int af, const char *src, void *dst);
struct sockaddr_in {
short sin_family;
unsigned short sin_port;
unsigned int sin_addr;
char sin_zero[8];
};
]]
local AF_INET = 2
local SOCK_DGRAM = 2
local SOL_SOCKET = 1
local SO_BROADCAST = 6
local _M = {}
function _M.send(mac_hex, broadcast_ip, cooldown)
cooldown = cooldown or 20
local wol_state = ngx.shared.wol_state
local last_wake = wol_state:get("last_wake_" .. mac_hex)
local now = ngx.now()
if last_wake and (now - last_wake) < cooldown then
return false
end
local sock = ffi.C.socket(AF_INET, SOCK_DGRAM, 0)
if sock < 0 then
ngx.log(ngx.ERR, "WoL: failed to create socket")
return false
end
local broadcast_enable = ffi.new("int[1]", 1)
ffi.C.setsockopt(sock, SOL_SOCKET, SO_BROADCAST, broadcast_enable, ffi.sizeof("int"))
local addr = ffi.new("struct sockaddr_in")
addr.sin_family = AF_INET
addr.sin_port = ffi.C.htons(9)
local ip_ptr = ffi.new("unsigned int[1]")
ffi.C.inet_pton(AF_INET, broadcast_ip, ip_ptr)
addr.sin_addr = ip_ptr[0]
local mac_bytes = ""
for i = 1, #mac_hex, 2 do
mac_bytes = mac_bytes .. string.char(tonumber(mac_hex:sub(i, i + 1), 16))
end
local magic = string.rep("\255", 6) .. string.rep(mac_bytes, 16)
local sent = ffi.C.sendto(sock, magic, #magic, 0, addr, ffi.sizeof(addr))
ffi.C.close(sock)
if sent < 0 then
ngx.log(ngx.ERR, "WoL: sendto failed")
return false
end
wol_state:set("last_wake_" .. mac_hex, now)
ngx.log(ngx.INFO, "WoL: magic packet sent for MAC ", mac_hex, " (", sent, " bytes)")
return true
end
return _M Nginx server config:
upstream inference_server {
server 192.168.50.142:8080 max_fails=1 fail_timeout=5s;
keepalive 32;
}
server {
listen 8080;
server_name inference ~^inference.*$;
client_max_body_size 100M;
keepalive_timeout 10;
limit_req zone=inference burst=20 nodelay;
add_header X-Frame-Options SAMEORIGIN always;
add_header X-Content-Type-Options nosniff always;
# Manual wake trigger
location = /wake {
content_by_lua_block {
local wol = require("wol")
local sent = wol.send("ac3da132635b", "192.168.50.255", 20)
ngx.header["Content-Type"] = "text/plain"
if sent then
ngx.say("WoL packet sent")
else
ngx.say("WoL skipped (cooldown active)")
end
}
}
# Main proxy — fires WoL preemptively on every request, cooldown prevents spam
location / {
access_by_lua_block {
local wol = require("wol")
wol.send("ac3da132635b", "192.168.50.255", 20)
}
proxy_pass http://inference_server;
proxy_http_version 1.1;
proxy_set_header Host $host;
proxy_set_header X-Real-IP $remote_addr;
proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
proxy_set_header X-Forwarded-Proto $scheme;
proxy_set_header Connection "";
proxy_connect_timeout 3s;
proxy_send_timeout 300s;
proxy_read_timeout 300s;
proxy_buffering off;
proxy_next_upstream error timeout;
error_page 502 504 = @wake_and_wait;
}
location @wake_and_wait {
default_type text/plain;
add_header Retry-After 15 always;
return 503 "Inference server is waking up, please retry in 15 seconds\n";
}
location /nginx-health {
access_log off;
return 200 "healthy\n";
add_header Content-Type text/plain;
}
} AI Agentic Ai --
A passionate DevOps/SRE. https://github.com/guilhermef
