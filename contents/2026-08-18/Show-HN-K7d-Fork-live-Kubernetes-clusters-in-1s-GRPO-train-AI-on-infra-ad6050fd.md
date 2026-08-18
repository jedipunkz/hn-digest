---
source: "https://github.com/katakate/k7d"
hn_url: "https://news.ycombinator.com/item?id=49346284"
title: "Show HN: K7d – Fork live Kubernetes clusters in <1s –> GRPO-train AI on infra"
article_title: "GitHub - Katakate/k7d: ⚡ The Rust VMM that unlocked forking live Kubernetes clusters in ~100 ms ⭐ Star it if you like it! · GitHub"
image: "https://opengraph.githubassets.com/7c94a154d1e072fcf00032aef50689858137ee5a7b2d06bce9ced4dcdcad887f/Katakate/k7d"
author: "gbxk"
captured_at: "2026-08-18T15:23:00Z"
capture_tool: "hn-digest"
hn_id: 49346284
score: 2
comments: 0
posted_at: "2026-08-18T14:38:58Z"
tags:
  - hacker-news
  - translated
---

# Show HN: K7d – Fork live Kubernetes clusters in <1s –> GRPO-train AI on infra

- HN: [49346284](https://news.ycombinator.com/item?id=49346284)
- Source: [github.com](https://github.com/katakate/k7d)
- Score: 2
- Comments: 0
- Posted: 2026-08-18T14:38:58Z

## Translation

タイトル: Show HN: K7d – ライブ Kubernetes クラスターを <1 秒でフォーク –> インフラ上で GRPO トレーニング AI
記事のタイトル: GitHub - Katakate/k7d: ⚡ ~100 ミリ秒でライブ Kubernetes クラスターのフォークのロックを解除した Rust VMM ⭐ 気に入ったらスターを付けてください! · GitHub
説明: ⚡ 約 100 ミリ秒でライブ Kubernetes クラスターのフォークのロックを解除した Rust VMM ⭐ 気に入ったらスターを付けてください! - カタカテ/k7d
HN テキスト: こんにちは、HN、ゲイリーです。今日は、Apache 2.0、タイトな Rust VMM + shim である k7d を紹介したいと思います。これにより、以前は不可能だったもの、つまり、実行中の接続を維持しながらライブ実行中の仮想化マルチノード k8s クラスターの高速フォークが可能になります。 64 GB RAM ボックスでは、3 VM ノードの K8s クラスターは 105 ミリ秒でフォークされ、3 VM クラスターの 50 倍のフォークは 4.1 秒で行われます。ここでの目標は 2 つあります。1) Kubernetes インフラ上で AI の大規模な GRPO/RL トレーニングを可能にすることです。IMO は、実際に役立つ機能のトレーニングに加えて、推論トレーニングの優れた遊び場です。そして、これには高速なエピソード リセット (RL ポストトレーニング中に数万回のマルチターン実行が必要なため) が必要なだけでなく、RL トレーニング中に並列ブランチ探索、ロールバック、枝刈りを実行できるようにするための高速フォークによる大きなメリットも得られます。また、忠実なフォークにより、GRPO の G のバイト同一の開始が得られ、グループ全体の分散が削減されます。 2) 私の他のプロジェクト K7 では、ユーザーフレンドリーな CLI / API / Python SDK と Kubernetes ネイティブを使用して、大規模な VM サンドボックス用のセルフホスト型インフラを提供します。それに加えて、k7d には次の機能が装備されています。 - リソース管理用のツリー型 API: フォークするときにわかるように、最適化された CoW ページ共有であっても、リソース (メモリ + ディスク) を適切に管理する必要があるため、実行中にエビクトする方法を知る必要があります。それで私は木を持っています-

形成されたロジックを使用して、子供が親とページを共有する方法を追跡し、AI エージェントに有望な木の枝を保護させたり、見込みのない枝を排除したり、リソースが逼迫したときに LRU のようなロジックを自動的に排除させたりすることができます。このツリーベースのロジックは、単一 VM サンドボックスと、独自の Linux ブリッジ上の複数 VM クラスターの両方に適用されます。 - 正式な検証: もちろんすべてではありませんが、k7d の選択された重要なサブパートが正式に検証されます。私は安全でないパスでのメモリ演算に Kani を使用し、上で説明したツリーベースのロジックを正式に証明するために (リーン バックエンドを備えた) Aeneas を使用して、エビクションによって生きている子孫によって参照されるページが決して解放されないようにします。 - CI としてのレイテンシー: 一連の統合テストを通じてチェック/強制された最も重要な操作のレイテンシーを厳密に追跡します。私は独自の VMM を構築しないように懸命に努力し、最初に最初の「kfd」(Kata + Firecracker + Devmapper-snapshotter over LVM Thin-pool) とは別のバックエンドを K7 用に構築することになりました。これを Kata + Qemu + Longhorn として「kql」と呼びました。このスタックを知っている場合は、すぐに推測できるでしょう。Longhorn はノード間のレプリケーションに最適なので、スナップショットをノード間でレプリケートするためにこれを使用しました。そのため、「スナップショットの再開」は常に機能します/HA。ここで Qemu を使用した理由は、Longhorn のブロック ストレージ要件が、クロスノード レプリケーション ロジックを自分で構築したくないバックエンドである Devmapper-snapshotter を必要とする Firecracker と互換性がなかったためです。しかし、この「kql」バックエンドは、Longhorn の構築方法により 45 秒でフォークを生成しました。これは、K7 の高速フォークを有効にするように私に要求したユーザーにとっては遅すぎました。これが私を、Firecracker/Qemu と Kata の両方を一度に置き換える、独自のネイティブ VMM とシムである「k7 のデーモン」と呼ばれる k7d に私を駆り立てた理由です。これにより、2 ～ 3 秒以内にフォークできる K7 の VM サンドボックスが生成され、そのほとんどが

VMM レベルではウォームフォークが実際には 5 ミリ秒であるため、レイテンシは kubelet のオーバーヘッドです。セキュリティ上のトレードオフの 1 つは、デーモンを同じツリーのブランチ間で共有する必要があるということです。これは仕様です。つまり、VM ごとに Firecracker の看守が失われます。ただし、ツリーごとに同様の Jailer を再構築することはできます。これは、RL トレーニングに分岐を使用する場合など、ツリーがテナント間で共有されていない場合には十分です。これは、Firecracker が行うこととは異なるものを最適化しているだけです。コードベースは意図的に監査できるほど厳密にされており (VMM + shim の LOC が 30k 未満)、README の先頭に詳細なブログ投稿シリーズをリンクしました。皆さんも楽しんでいただければ幸いです。寄稿者や批評家の方もぜひご参加ください。 THX！

記事本文:
GitHub - Katakate/k7d: ⚡ 約 100 ミリ秒でライブ Kubernetes クラスターのフォークのロックを解除した Rust VMM ⭐ 気に入ったらスターを付けてください! · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
カタカテ
/
k7d
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
4 コミット 4 コミット フォルダーとファイル
.devin .devin 資産 アセット クレート クレート ドキュメント ドキュメント 例/cluster-tree-search

例/cluster-tree-search fuzz fuzz guest guest パッケージング utils utils verif verif CHALLENGES.md CHALLENGES.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml HACKING.md HACKING.md LATENCY_BUDGETS.md LATENCY_BUDGETS.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md Deny.toml Deny.toml Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
実行中の Kubernetes クラスターを約 100 ミリ秒でフォークします。
1 つの 64 GB ボックスで 50 コピーを実行します。
Kubernetes 環境の RL トレーニングとエージェント評価 —
クラスター、チャート、クラスター内のワークロード - 何千もの分離された、
リセット可能な世界。 1 つあたり 1 つのサンドボックスやコールド クラスタではありません。
裁判。新しい Kubernetes クラスターの起動には最大 30 秒かかり、RAM がいっぱいになります
コピーごとに。 k7d は一度ブートし、ライブクラスターをフォークします。
～100ミリ秒。フォークは分岐するまでメモリを共有するため、50 個のコピーにコストがかかります
ダーティ ページ、50 × フル ゲスト RAM ではありません。
同じエンジンは、ユニットが
単一 VM サンドボックス (VM 内のドッカーを含む): 非常に高速なウォーム フォーク
メモリ、ディスク、プロセス、ネットワークの忠実なスナップショット。のために
それを大規模に実行する — Kubernetes がサンドボックスを調整し、
エージェント用の CLI / API / Python SDK — 兄弟プロジェクトを参照
カタカテ k7 。
100% オープンソース (Apache-2.0)。技術サポートについては、hi@katakate.org までご連絡ください。
⚡ 実行中の Kubernetes クラスターを ~100 ミリ秒でフォークし、パックします
1 つの 64 GB ボックスに 50 コピー。新しいクラスターのコールドブート
所要時間は約 30 秒で、コピーごとに RAM の全額が請求されます。 k7d が 1 回起動してから、
コピーは分岐するまでメモリを共有します。
🔁 クラスターはフォーク後も実行を続けます。エージェントは再起動されません。
壊れた TLS も、「Kubernetes が戻るまでお待ちください」もありません。それぞれ
コピーは私を見て

内側からオリジナルと同じです。
🌲 多くの世界を探索する AI エージェント向けに構築 — ブランチをフォークし、
何かを試して、勝者を残し、敗者を捨てる。エージェント
何を保持するかを決定します。 k7d は RAM とディスクの割り当てを強制するため、ツリー
機械を食べません。
🧊 また、超高速の VM サンドボックス VMM — 単一ゲストをウォームフォークします
メモリ / ディスク / プロセス / ネットワークの状態が忠実であれば、最大 5 ミリ秒で完了します。
VM 内の Docker およびリセット可能が必要なワークロードに最適
k8s クラスター全体だけでなく、分離されたマシンも対象となります。とペアリングする
k8s が必要な場合は Katakate/k7
オーケストレーション + Python SDK がその上にあります。
🔬 私たちは、効果が得られる正式な方法を使用します — Kani
選択された安全でない/算術パス上、および
アエネアス →木にもたれかかる
予算/立ち退きモデル。すべてが証明されているという主張ではありません —
詳細は下記をご覧ください。
🪶 VMM + shim の Rust ≤25k 行 — 読み取るのに十分な小ささ
そして監査。意図的にキッチンシンク用の VMM ではありません。
✅ ここにあるすべての数値は CI アサーションであり、1 回限りのベンチマークではありません
ペースト。レイテンシーの主張が変動すると、テストは失敗します。方法論:
ベンチマークの書き込み。
KVM を備えた Linux amd64 / x86_64 ホストが必要です (/dev/kvm が存在します) -
同じ ISA ( amd64 は Debian 名です。tarball は x86_64 を使用します)。アームなし64
まだ構築します。ソースからビルドするには Rust と Docker が必要です。事前構築済み
release tarball は make release によって生成されます
( k7d-v*-x86_64-linux.tar.gz + install.sh )。
git clone https://github.com/Katakate/k7d && cd k7d
make release # デーモン、shim、ゲストカーネル、rootfs → dist/
sudo RUST_LOG=info ./dist/k7d & # 所有する VM + /run/k7d/k7d.sock
cd 例/cluster-tree-search
python3 run_demo.py --modebusybox --branches 4
# 密度要求: python3 run_demo.py --modebusybox --branches 50
デモでは、ライブ 3 VM クラスターをフォークし、分岐をスコアリングし、勝者を維持します。
敗者を剪定し、フォークを出力します w

オールクロック。本物を入れるには
これらの VM 内の Kubernetes ポッド ( runtimeClassName: k7 ) を参照してください。
ハッキング.md 。完全なドキュメント (API リファレンス、インストーラー)
個別に発送します — この README は製品の提案と入門書です。
GRPO / エージェントツリー検索に k7d を使用する
すでに Kubernetes タスクまたはシナリオ (Helm チャート、
YAML マニフェストのセット、kube-apiserver と通信する eval ハーネス)、
形状は次のとおりです。
シナリオを一度起動します。クラスターを起動します (または、
k7d がすでにホストしているものを実行中）、それが実行されるまで待ちます。
すべてのロールアウトを開始する状態。
そのチェックポイントでツリーを根付かせます。
各 GRPO グループ (またはツリー検索ステップ) について:
fork_batch(N) → N 個のコピーに対して N 個のポリシーを実行 → スコア
→ 勝者を保護し、敗者を排除する → デーモンに任せる
auto_evict は RAM/ディスクの予算内で実行してください。
次の勝者が必要な場合は、保護された勝者からロールフォワードします。
より良い状態から開始する世代、またはより良い状態にロールバックする世代
そうでない場合は、前のノードを使用します。
GRPO にとってバイト同一の開始が重要な理由
GRPO (およびほとんどのグループ相対メソッド) は、範囲内で報酬を比較します。
グループ。メンバー A がよりコールドなキャッシュから開始される場合、別の etcd
リビジョン、またはメンバー B よりも半分準備ができているデプロイメントの場合、報酬の差は次のとおりです。
信号ではなくノイズです。 k7d フォークはライブマシンのコピーです - 同じです
メモリ、同じディスク、同じクラスター内 TLS セッション、同じ kube-apiserver
状態。グループのすべてのメンバーは、同一のバイトから始まります。
世界は、あなたのポリシーが行ったことによってのみ分岐します。
それが、「環境をリセットした」ことと「環境をクローンした」ことの違いです。
宇宙。」
エージェントがツリーと対話する方法
トレーニング ループには報酬とポリシーがあります。 k7d は環境を所有しており、
予算。エージェントは Unix ソケットを介して JSON 行を通信します
( /run/k7d/k7d.sock )。実際に必要な動詞:
正確にカバーするシン Python クライアント

ループが存在する
例/クラスターツリー検索/ 。治療する
これは、完成品としてではなく、GRPO トレーナーを配線するためのテンプレートとして使用されます。
SDK。完全な API リファレンスはドキュメント サイトにあります。
仕組み (明らかではない部分)
k7d を使用するのに VMM エンジニアである必要はありません。知っておく必要があります
なぜ 100 ミリ秒のクラスターフォークが可能なのかというと、それが
製品。
誰かが書き込むまで、メモリは共有されます。ゲスト RAM は 1 つのファイル内に存在します。
フォークはソースを一時停止し、それ以降にどのページが変更されたかを記録します。
最後のチェックポイントでは、子のメモリをコピーオンライト ビューとしてマップします。
親のページをコピーし、それらのダーティ ページのみをコピーします。それ以外はすべてそうです
共有されました。これが、クラスターの 50 フォークが 64 GB に収まる理由です。
ベースのためではなく、ダイバージェンス。
ネットワークは一貫して存在するため、クラスターは再起動しません。
各クラスターは独自のプライベート Linux ブリッジ上に存在します。フォークは
ソースと同じゲスト IP および MAC アドレスを持つ新しいブリッジ。
ゲスト内部からは何も移動しません - 同じアドレス、同じ ARP キャッシュ、
同じ TLS 証明書、同じ確立された TCP、つまり kubelet、CNI、および
コントロールプレーンは実行を続けます。独立したブリッジは、フォークがそれぞれを認識できないことを意味します
その他。これがないと、フォークごとに kubelet の再起動が強制され、
ノードごとにエージェントが再起動するのは約 1 ～ 2 秒で、100 ミリ秒は次のようになります。
不可能です。
フォークは、デーモンが予算に従って管理するツリー内に存在します。
ベースクラスター ──► フォーク A ──► フォーク A1 (保護: 勝者)
§─► フォークB（剪定：低報酬）
└─► フォーク C ──► ロールバック ─► フォーク C'
正しくなければならなかった 3 つのトリック
これらは、「バイト同一性」を静かに破壊する種類のバグです。
CHALLENGES.md には 56 個すべてが含まれています。この3つが
見出しのもの:
デバイスの書き込みは、ハイパーバイザーのダーティ ログには表示されません。
k7d のブロック デバイスは、ユーザー空間からゲスト メモリに書き込みます。の
ハイパーバイザーのみ

y は CPU の書き込みを認識します。つまり、ディスク I/O 後にフォークが実行されます。
これらのページでは、I/O 前のバイトが復活します (サイレント破損)。
修正: デバイスのダーティ ページを自分で追跡し、それらを
ビットマップをフォークし、ビットマップが読み取られる前にインフライト I/O を排出します。
( チャレンジ.md #43 )
空のタイマー チップを備えた復元されたゲストは時間がフリーズします。後
フォーク、インターバル タイマーがプログラムされていないままになっている - タイマーなし
割り込みが発生し、CLOCK_REALTIME がスタックし、Kubernetes が静かにパークします。
修正: すべてのデバイスでタイマーを再設定 (および準仮想クロックをリセット)
復元します。 ( チャレンジ.md #40 )
同一の IP は、L2 ドメインが分離されている場合にのみ機能します。リプレイ
送信元のアドレスをフォークごとに新しいブリッジに転送します。同じ見解
内側から。フォーク間での競合はありません。それがゼロリスタートです
上記のトリック。
1 つのベアメタル ボックス (月額約 40 ユーロのベアメタル: Ryzen 5 3600、6 コア、
64 GiB、NVMe):
すべての行は統合テスト アサーションです
( LATENCY_BUDGETS.md )。完全な方法論:
ベンチマークの書き込み。
2 つのレイヤー — ほとんどの RL ユーザーは最初のレイヤーのみを気にします。
フォークされたクラスターの内部 (GRPO シナリオ)
これは、フォークした VM 内に存在する k3 です。フォークエンジン
N ノード (tree_create_cluster(vm_count) / 任意のライブ セットを採用) —
ハードコードされたものはありません。 3. ヘッドラインを証明する CI フィクスチャ
数値は、フランネル + kube-proxy と実際の 3 ノードのコントロール プレーンです。
クラスター内デプロイメント。いくつかのストック k3 アドオンはまだ無効になっています
その道を無駄なく保つためにそこにあります。以下のステータスには「API で実行可能」と「API で実行可能」が混在しています。
「フィクスチャはそれを行使します。」
明日、シナリオで CoreDNS + Ingress が必要な場合は、そう言ってください - 再度有効にします
標準のアドオンは、再設計ではなく、次の忠実度の向上です。同じ
より大きな CI フィクスチャ: 20 ノードの配線は構成 + RAM であり、新しいものではありません
フォーク機能。
Host RuntimeClass (k7d VM としてのポッド)
これは外側の層です: h 上の kubectl

ost はポッドを k7d にスケジュールします
runtimeClassName 経由の microVM: k7 。必要な場合も関連します
クラスター全体のフォークだけでなく、単一 VM のサンドボックスも。
Firecracker / Kata / E2B スタイルのサンドボックスを使用しないのはなぜですか?
爆竹
カタ
CubeSandbox / E2B スタイル
k7d
実行中のVMのウォームフォーク
スナップショット→復元
いいえ
スナップショット + N 回の復元 (~220 ミリ秒)
ライブ コピーオンライト フォーク (~5 ミリ秒)
スナップショット ツリー (フォーク / ロールバック / 保護 / バジェット)
いいえ
いいえ
サンドボックス周りのSDK
はい — デーモン API
k8s クラスター全体をフォークします
いいえ
いいえ
いいえ
はい (~105 ミリ秒)
Kubernetes RuntimeClass として実行
FCコンテナ経由
はい
いいえ
はい ( runtimeClassName: k7 )
形式的手法
監査/ファズ文化
—
—
選択したパス上の Kani + Aeneas (メモリ計算、ツリー バジェット)
彼らはサンドボックスをフォークします。 k7d は VM またはクラスター全体をフォークします。
意図的に E2B-API 互換ではありません - 別の仕事です。
1 つのデーモン、多数の VM、1 つのアドレス空間。ライブコピーオンライトフォーク
親メモリと子メモリが同じプロセス内でマッピングされる必要があります。
ゲスト→ホストの分離は依然として KVM です。兄弟フォーク間の分離
同じテナントのは Firecracker よりも弱いです
VM ごとに 1 つのジェイルされたプロセス。 k7d は独自のフリート用に構築されています
環境、フォーク間の敵対的なマルチテナント分離ではありません。
フォーク後に生き残るもの: フォークされたセット内のすべて —
クラスター内 TLS、メンバー VM 間で確立された TCP、ディスク状態。
ゲストの時計はリセットされるため、時間が逆戻りすることはありません。
何だ

[切り捨てられた]

## Original Extract

⚡ The Rust VMM that unlocked forking live Kubernetes clusters in ~100 ms ⭐ Star it if you like it! - Katakate/k7d

Hey HN, Gary here. Today I want to present k7d which is an Apache 2.0, tight Rust VMM + shim enabling something not possible before: fast forking of live running virtualized multi-node k8s clusters with surviving of in-flight connections. A 3-VM nodes K8s cluster gets forked in 105ms, and a 50x fork of a 3-VM cluster in 4.1s on a 64GB RAM box. I have two goals here: 1) enable large scale GRPO/RL training of AI on Kubernetes infra, which IMO is a great playground for reasoning training, besides training a capability that's actually useful. And this requires not only fast episode reset (as you need tens of thousand of multi-turn runs during RL post-training) but also greatly benefits from fast forking so you can do parallel branch exploration, rollback, pruning during RL training. Faithful forks also give you byte-identical starts for the G of GRPO, which gives variance reduction across the group. 2) enable <3s VM snapshot pause/resume/fork of sandboxes with Docker-in-VM, for my other project K7 which provides self-hosted infra for VM sandboxes at scale, with a user-friendly CLI / API / Python SDK, and Kubernetes native. Besides that, k7d is equipped with: - A Tree-shaped API for resource management: as you can guess when you fork, even with optimized CoW-page-sharing, you want to properly manage your resources (memory + disk) and hence you need to know how to evict while things run. So I have tree-shaped logic to keep track of how children share pages with parents, and let your AI agent protect a promising tree branch, evict an unpromising one, or let LRU-ish logic auto-evict when resources get tight. This tree-based logic applies both to single VM sandboxes, and to multi-VM clusters on their own Linux bridge. - Formal verification: of course not all of it, but selected critical subparts of k7d are formally verified: I use Kani for memory arithmetics in the unsafe paths, and Aeneas (with Lean backend) to formally prove the tree-based logic explained above so that eviction never frees a page referenced by a live descendent. - Latencies as CI: I rigorously keep track of latency for most important operations which remain checked/enforced via a suite of integration tests. I really tried hard not building my own VMM and first ended up building another backend for K7 than my initial "kfd" (Kata + Firecracker + Devmapper-snapshotter over LVM thin-pool), which I called "kql" for Kata + Qemu + Longhorn. If you know this stack you'll guess it right away: Longhorn is great for cross-node replication so I used it to have my snapshots replicated across nodes, so "snapshot resume" always works / HA. Qemu here is because Longhorn's block storage requirements was incompatible with Firecracker who wants Devmapper-snapshotter, a backend for which I would not want to build myself the cross-node replication logic. But this "kql" backend yielded forks in 45s due to how Longhorn is built, which was too slow for the users who asked me to enable fast forking for K7. So this is what pushed me towards k7d, named as "k7's daemon", its own native VMM and shim, replacing both Firecracker/Qemu and Kata at once. This yields VM sandboxes in K7 which you can fork in under 2-3s, and most of this latency is kubelet overhead, as at the VMM level the warm-fork is actually 5ms. One security trade-off: the daemon has to be shared across branches of a same tree: that's by design. So you lose Firecracker's Jailer per VM. But I could re-build a similar Jailer per tree, which would be sufficient when a tree isn't shared across tenants, such as when you use branching for RL training. That's just optimizing for something different than what Firecracker does. The codebase is intentionally tight enough to be audited (<30k LOC for VMM + shim) and I linked a deep-dive blogpost series at the top of the README. I hope you guys will enjoy it and I'd love contributors and critics. Thx!

GitHub - Katakate/k7d: ⚡ The Rust VMM that unlocked forking live Kubernetes clusters in ~100 ms ⭐ Star it if you like it! · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
Katakate
/
k7d
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
.devin .devin assets assets crates crates docs docs examples/ cluster-tree-search examples/ cluster-tree-search fuzz fuzz guest guest packaging packaging utils utils verif verif CHALLENGES.md CHALLENGES.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml HACKING.md HACKING.md LATENCY_BUDGETS.md LATENCY_BUDGETS.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md deny.toml deny.toml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Fork a running Kubernetes cluster in ~100 ms.
Run 50 copies on one 64 GB box.
RL training and agent evals whose environments are Kubernetes —
clusters, charts, in-cluster workloads — need thousands of isolated,
resettable worlds. Not one sandbox, and not a cold kind cluster per
trial. Booting a fresh Kubernetes cluster takes ~30 s and full RAM
per copy. k7d boots it once , then forks the live cluster in
~100 ms; forks share memory until they diverge, so 50 copies cost
dirty pages, not 50 × full guest RAM.
The same engine is also a great fork-first VMM when your unit is a
single VM sandbox (including docker-in-VM): blazing-fast warm forks with
faithful snapshotting of memory, disk, processes, and networking. For
running that at scale — Kubernetes orchestrating your sandboxes, plus a
CLI / API / Python SDK for agents — see the sibling project
Katakate k7 .
100% open‑source (Apache‑2.0). For technical support, write us at: hi@katakate.org
⚡ Fork a running Kubernetes cluster in ~100 ms — and pack
50 copies on one 64 GB box . Cold-booting a fresh cluster
takes ~30 s and a full RAM bill per copy; k7d boots once, then
copies share memory until they diverge.
🔁 The cluster keeps running after the fork — no agent restarts,
no broken TLS, no "please wait while Kubernetes comes back." Each
copy looks identical to the original from the inside.
🌲 Built for AI agents that explore many worlds — fork a branch,
try something, keep the winners, throw away the losers. The agent
decides what to keep; k7d enforces RAM and disk budgets so the tree
doesn't eat the machine.
🧊 Also a blazing-fast VM-sandbox VMM — warm-fork a single guest
in ~5 ms with faithful memory / disk / process / network state.
Ideal for docker-in-VM and any workload that needs resettable
isolated machines, not only whole k8s clusters. Pair with
Katakate/k7 when you want k8s
orchestration + Python SDKs on top.
🔬 We use formal methods where they pay off — Kani
on selected unsafe / arithmetic paths, and
Aeneas →Lean on the tree
budget/eviction model. Not a claim that everything is proven —
details below .
🪶 ≤25k lines of Rust for the VMM + shim — small enough to read
and audit. Deliberately not a kitchen-sink VMM.
✅ Every number here is a CI assertion — not a one-off benchmark
paste. If a latency claim drifts, a test fails. Methodology:
the benchmark write-up .
You need a Linux amd64 / x86_64 host with KVM ( /dev/kvm present) —
same ISA ( amd64 is the Debian name; tarballs use x86_64 ). No arm64
build yet. Rust and Docker are required to build from source. Prebuilt
release tarballs are produced by make release
( k7d-v*-x86_64-linux.tar.gz + install.sh ).
git clone https://github.com/Katakate/k7d && cd k7d
make release # daemon, shim, guest kernel, rootfs → dist/
sudo RUST_LOG=info ./dist/k7d & # owns VMs + /run/k7d/k7d.sock
cd examples/cluster-tree-search
python3 run_demo.py --mode busybox --branches 4
# density claim: python3 run_demo.py --mode busybox --branches 50
The demo forks a live 3-VM cluster, scores branches, keeps the winner,
prunes the losers, and prints the fork wall-clock. To put real
Kubernetes pods inside those VMs ( runtimeClassName: k7 ), see
HACKING.md . Full docs (API reference, installers) will
ship separately — this README is the product pitch + getting started.
Using k7d for GRPO / agent tree search
If you already have Kubernetes tasks or scenarios (a Helm chart, a
set of YAML manifests, an eval harness that talks to a kube-apiserver),
the shape is:
Boot the scenario once — bring up your cluster (or adopt a
running one that k7d already hosts) and wait until it is in the
state you want every rollout to start from.
Root a tree at that checkpoint.
For each GRPO group (or tree-search step):
fork_batch(N) → run your N policies against the N copies → score
→ protect the winners, prune the losers → let the daemon
auto_evict under your RAM/disk budget.
Roll forward from a protected winner when you want the next
generation to start from a better state, or rollback to an
earlier node when you don't.
Why byte-identical starts matter for GRPO
GRPO (and most group-relative methods) compare rewards within a
group . If member A starts from a colder cache, a different etcd
revision, or a half-ready Deployment than member B, the reward gap is
noise, not signal. A k7d fork is a copy of the live machine — same
memory, same disk, same in-cluster TLS sessions, same kube-apiserver
state. Every member of the group begins from a byte-identical
world, then diverges only because of what your policy did.
That is the difference between "we reset the env" and "we cloned the
universe."
How your agent talks to the tree
Your training loop owns rewards and policy. k7d owns environments and
budgets. The agent talks JSON-lines over a Unix socket
( /run/k7d/k7d.sock ). The verbs you actually need:
A thin Python client that covers exactly this loop lives in
examples/cluster-tree-search/ . Treat
it as the template for wiring your GRPO trainer — not as a finished
SDK. The full API reference will live in the docs site.
How it works (the non-obvious bits)
You do not need to be a VMM engineer to use k7d. You do need to know
why a 100 ms cluster fork is even possible, because that is the
product.
Memory is shared until someone writes. Guest RAM lives in one file.
A fork pauses the source for a moment, notes which pages changed since
the last checkpoint, maps the child's memory as a copy-on-write view of
the parent's, and copies only those dirty pages. Everything else is
shared. That is why 50 forks of a cluster fit in 64 GB: you pay for
divergence, not for the base.
The cluster does not reboot because the network lies consistently.
Each cluster lives on its own private Linux bridge. A fork gets a
new bridge with the same guest IPs and MAC addresses as the source.
From inside the guest, nothing moved — same addresses, same ARP cache,
same TLS certs, same established TCP — so kubelet, the CNI, and the
control plane keep running. Separate bridges mean forks cannot see each
other. Without this, every fork would force a kubelet restart and a
~1–2 s agent restart per node, and the 100 ms claim would be
impossible.
Forks live in a tree the daemon manages under budget:
base cluster ──► fork A ──► fork A1 (protected: winner)
├─► fork B (pruned: low reward)
└─► fork C ──► rollback ─► fork C'
Three tricks that had to be right
These are the kinds of bugs that silently break "byte-identical."
CHALLENGES.md has all 56; these three are the
headline ones:
Device writes are invisible to the hypervisor's dirty log.
Block devices in k7d write guest memory from userspace. The
hypervisor only sees CPU writes — so a fork taken after disk I/O
would resurrect pre-I/O bytes on those pages (silent corruption).
Fix: track device dirty pages ourselves and merge them into the
fork bitmap, and drain in-flight I/O before the bitmap is read.
( CHALLENGES.md #43 )
A restored guest with a blank timer chip freezes time. After
fork, the interval timer was left unprogrammed — no timer
interrupts, CLOCK_REALTIME stuck, and Kubernetes quietly parks.
Fix: re-arm the timer (and reset the paravirtual clock) on every
restore. ( CHALLENGES.md #40 )
Identical IPs only work if the L2 domains are separate. Replay
the source's addresses onto a fresh bridge per fork. Same view
from inside; no conflicts across forks. That is the zero-restart
trick above.
On one bare-metal box (~€40/month bare-metal: Ryzen 5 3600, 6 cores,
64 GiB, NVMe):
Every row is an integration-test assertion
( LATENCY_BUDGETS.md ). Full methodology:
the benchmark write-up .
Two layers — most RL users only care about the first.
Inside a forked cluster (your GRPO scenario)
This is the k3s that lives inside the VMs you fork. The fork engine
is N-node ( tree_create_cluster(vm_count) / adopt any live set) —
there is no hard-coded 3. The CI fixture that proves the headline
numbers is a 3-node control plane with flannel + kube-proxy and a real
in-cluster Deployment; several stock k3s add-ons are still disabled
there to keep that path lean. Status below mixes “API can do it” with
“fixture exercises it.”
If your scenario needs CoreDNS + Ingress tomorrow, say so — re-enabling
the stock add-ons is the next fidelity bump, not a redesign. Same for
a larger CI fixture: wiring 20 nodes is configuration + RAM, not a new
fork feature.
Host RuntimeClass (pods as k7d VMs)
This is the outer layer: kubectl on the host schedules pods into k7d
microVMs via runtimeClassName: k7 . Relevant if you also want
single-VM sandboxes, not only whole-cluster forks.
Why not Firecracker / Kata / E2B-style sandboxes?
Firecracker
Kata
CubeSandbox / E2B-style
k7d
Warm fork of a running VM
snapshot → restore
no
snapshot + N restores (~220 ms)
live copy-on-write fork ( ~5 ms )
Snapshot tree (fork / rollback / protect / budget)
no
no
SDK around sandboxes
yes — daemon API
Forks a whole k8s cluster
no
no
no
yes (~105 ms)
Runs as a Kubernetes RuntimeClass
via FC-containerd
yes
no
yes ( runtimeClassName: k7 )
Formal methods
audit/fuzz culture
—
—
Kani + Aeneas on selected paths (memory math, tree budgets)
They fork a sandbox. k7d forks a VM or an entire cluster .
Deliberately not E2B-API compatible — different job.
One daemon, many VMs, one address space. Live copy-on-write fork
requires parent and child memory to be mappings in the same process.
Guest→host isolation is still KVM; isolation between sibling forks
of the same tenant is weaker than Firecracker's
one-jailed-process-per-VM. k7d is built for fleets of your own
environments, not hostile multi-tenant isolation between forks.
What survives a fork: everything inside the forked set —
in-cluster TLS, established TCP between member VMs, disk state.
Guest clocks are reset so time does not jump backwards.
What d

[truncated]
