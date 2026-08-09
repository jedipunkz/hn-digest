---
source: "https://github.com/JustVugg/lumabri"
hn_url: "https://news.ycombinator.com/item?id=49236781"
title: "Show HN: Lumabri – What if LLMs worked like Napster?"
article_title: "GitHub - JustVugg/lumabri: Run huge MoE models from a swarm of peers, with the colibri engine. Pure C. · GitHub"
author: "vforno"
captured_at: "2026-08-09T23:20:51Z"
capture_tool: "hn-digest"
hn_id: 49236781
score: 7
comments: 1
posted_at: "2026-08-09T22:24:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Lumabri – What if LLMs worked like Napster?

- HN: [49236781](https://news.ycombinator.com/item?id=49236781)
- Source: [github.com](https://github.com/JustVugg/lumabri)
- Score: 7
- Comments: 1
- Posted: 2026-08-09T22:24:24Z

## Translation

タイトル: 表示 HN: Lumabri – LLM が Napster のように機能したらどうなるでしょうか?
記事のタイトル: GitHub - JustVugg/lumabri: colibri エンジンを使用して、ピアの群れから巨大な MoE モデルを実行します。 Pure C. · GitHub
説明: colibri エンジンを使用して、ピアの群れから巨大な MoE モデルを実行します。 Pure C. - JustVugg/ルマブリ
HN テキスト: 少し前、私は通常のコンピューター上で巨大な LLM を実行できるかどうかを確認するために Colibrì に取り組み始めました。このプロジェクトは、主に HackerNews コミュニティのおかげで、私の予想をはるかに超えて成長しました。そこで私は新たな疑問を抱きました。1 台のコンピューターについて考えるのをやめたらどうなるでしょうか?これがルマブリの背後にある考え方です。 Lumabri は、巨大なモデル全体を保存して実行するために単一のマシンを必要とするのではなく、通常のコンピューターのネットワークをリソースの共有プールとして扱います。 1 台のマシンがディスク領域を提供し、別のマシンがコンピューティングを提供し、別のマシンがモデルの異なる部分を提供する場合があります。必要なブロックまたはエキスパートがローカルで利用できない場合、システムはそれをピア上で取得または実行できます。これは、専門家混合モデルにとって特に興味深いものです。モデルには数千億のパラメーターを含めることができますが、各トークンでアクティブになるのはほんの一部だけです。 Lumabri は、巨大なエキスパートの重みをネットワーク上で移動するのではなく、すでにエキスパートを持つピアに小さなアクティベーションを送信し、実行させることができます。目標は、マシンが余裕のあるリソースをすべて提供し、残りのリソースには swarm を使用することです。このアイデアはピアツーピア システムから非常にインスピレーションを得ており、ユーザーはインフラストラクチャです。特にネットワーク遅延とセキュリティなど、明らかに大きな課題があります。私はピア検証、SHA-256 検証、署名されたモデルの状態、レプリカの選択、フェイルオーバー、決定論的実行を実験しています。ルマブリはまだ初期の実験です。データセンターも巨大な GPU CL も持っていません

そこで、私は手持ちのハードウェアを使ってそれを構築し、そのアイデアが実際に意味があるかどうかを確認しようとしています。 Colibrì で私は「1 台の通常のコンピュータで巨大な LLM を実行できるでしょうか?」と尋ねました。 Lumabri で私はこう問います。多数の通常のコンピューターが 1 台の巨大なコンピューターになれるとしたらどうなるでしょうか?フィードバックは歓迎です。リポジトリ: https://github.com/JustVugg/lumabri

記事本文:
GitHub - JustVugg/lumabri: colibri エンジンを使用して、ピアの群れから巨大な MoE モデルを実行します。 Pure C. · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ジャストヴグ
/
ルマブリ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
27 コミット 27 コミット Engine_patches Engine_patches Expert_engines Expert_engines .gitignore .gitignore DEPLOY.md DEPLOY.md DESIGN.md DESIGN.md LICENSE LICENSE Makefile Makefile README.md README

.md RESULTS_PHASE2.md RESULTS_PHASE2.md assign_test.sh assign_test.sh chat_proto_test.sh chat_proto_test.sh concurrency_test.sh concurrency_test.shデモ_ディープシーク.sh デモ_ディープシーク.sh Expert_node.c Expert_node.c logo.svg logo.svg lumabri.c lumabri.c lumabri_client.h lumabri_client.h lumabri_proto.h lumabri_proto.h lumabri_sha.h lumabri_sha.h lumabri_sign.h lumabri_sign.h lumashim.c lumashim.c lumabri_client.h lumabri_client.h mainer.c mainer.c make_tiny_inkling.py make_tiny_inkling.py make_tiny_kimi.py make_tiny_kimi.py make_tiny_olmoe.py make_tiny_olmoe.pyphase2_bench.shphase2_bench.shphase2_deepseek_test.shphase2_deepseek_test.shphase2_glm_test.shphase2_glm_test.sh phase2_inkling_test.shphase2_inkling_test.shphase2_kimi_test.shphase2_kimi_test.shphase2_test.shphase2_test.shphase3_test.shphase3_test.shphase4_test.shphase4_test.shphase5_test.shphase5_test.shprobe_rtt.cprobe_rtt.csecurity_test.shsecurity_test.shselftest.shselftest.sh Sign_test.sh Sign_test.sh swarm.pub swarm.pub test_shim.c test_shim.c tracker.c tracker.c すべてのファイルを表示 リポジトリ ファイルのナビゲーション
多数のピアから専門家が混在した大規模なモデルを実行します。
コリブリエンジン。純粋なC、いいえ
依存関係。
1 台のマシンがモデルを共有します。他のマシンはそれとチャットします。何もありません。
前もってダウンロードされているため、推論が実際に触れるバイトは、
最初の使用時にピアし、ローカルミラーに留まります。 2番目の質問が提供されます
ローカルディスクからフルスピードで。エンジンのバイナリは変更されていません。
創設原則: GPU の有無にかかわらず、どのマシンも参加できる。エンジンは
最初に CPU と SSD 用に構築されました。 GPU は高速化するだけで、決して変わりません。
どちらの方法でも出力はバイト同一です。 GPU をまったく持たない swarm は、
働く群れ。 GPU をプールするネットワークは少数から採用します。ルマブリ
r

みんなからのエクリュ。
作る
モデルがあるマシン (任意の colibri モデル ディレクトリ):
./lumabriserve --model /path/to/model
チャットしたいマシン上で (エンジン用に colibri ビルドが必要です):
./lumabri chat --tracker < サーバー IP > :7300 --engines-dir /path/to/colibri/c
以上です。最初の答えは、ワーキング セットが次の領域を通過する間は遅くなります。
ネットワーク。その後、~/.lumabri 内のミラーは、たとえ
サーバーがオフラインになります。
チャット内: /swarm はネットワークがライブで匿名であることを示します (ピアは
番号付き、名前なし: 保持されているモデル、GB、提供されたバイト数、ハートビート)、/model
swarm 上のモデルをリストし、それらを切り替えて、
オンザフライエンジン。
TUI しか開かない人向け
ルマブリ
引数はありません。群れのアドレスを尋ね、さらに一度オペレーターのアドレスも尋ねます。
公開鍵。エンジン自体を見つけます。そしてそれはすべてを覚えています
~/.lumabri/config なので、2 回目は Enter、Enter で完了です。
それは利便性よりも重要です。かつては要求されるものはすべてフラグだった、
そして、一つ間違えても「無効な引数」、つまり欠落した引数が生成されることはありません。
--engines-dir では 299 GB のダウンロードが生成され、キーが欠落していると
誰も検証していないモデル。フラグが与えられた場合でも優先されるため、スクリプトは決して有効ではありません
誰かが保存した回答を継承します。
参加: チャットするか、何かを持参してください
lumabri チャットは途中で 1 回尋ねます。Enter は「チャットするだけ」を意味するため、
せっかちなパスが鍵の 1 つです。
エントリー・ネロ・シャームに来ますか？
1 個のソロチャットが非コンディヴィディニエンテ
2 チャット イー ドニ ディスコ ティエニ アン ペッツォ ディ グルム パー ロスシアメ
3 会話をしながら、すべてのことを話し合う
4 提出期限
invio = ソロチャタレ
2 を選択すると、何 GB かを尋ねられます (Enter を押すと空き領域の 4 分の 1 が使用されます)。
上限付き)、その予算でメンテナを開始します。トラッカーが予算を割り当てます。
最も複製の少ないファイルが最初に、

すべてのバイトを検証してそれらをプルします
オペレーターの署名を付けてサービスを提供します。 3 を選択するとエキスパートが開始されます
そのモデルのエンジンのノード。どちらもチャットの子として実行され、次の場合に停止します。
あなたはそれを閉じます - それは、から提供されたものの正当な寿命です
開いている端末。セッションを超えて存続する必要があるドナーは、
ルマブリ サーブ --join 。
コンピューティングを寄付するには、ディスク上のモデルが必要です (エキスパート ノードは、
そこから重みを取得するため)、オプション 3 は --model-dir DIR でのみ提供されます。
ディスクを寄付するのに何も必要はありません。空から始めると、群れがいっぱいになります。
スクリプトは質問をスキップします: --role chat|disk|compute|all 、 --donate GB を使用
および --model-dir DIR 。
1 つのトラッカーはモデル サーバーではなくインデックスであるため、同じ数のモデルを保持します。
1 台のマシンまたは複数のマシンからそれをポイントします。
ルマブリ サーブ --model /models/glm --port 7300
lumabriserve --model /models/olmoe --join 127.0.0.1:7300 --port 7310
lumabriserve --model /models/deepseek --join 127.0.0.1:7300 --port 7320
各サーブは独自のメンテナー (バイト) と独自のエキスパート ノードを持ちます。
(コンピューティング)、それぞれがそのモデルの名前で登録されます。おしゃべりが見える
3 モデルとスイッチ /model <name> : エンジンは
新しいモデルに対して再起動され、エンジン バイナリが次から選択されるため、
そのモデルの model_type 、異なるアーキテクチャ間の切り替えは機能します
また、GLM から OLMoE、DeepSeek、1 つのクライアント、1 つのトラッカー。
ドナーもモデルごとに存在します。マシンは 1 つのモデルのスライスを保持でき、
専門家を別の人のために処刑する。トラッカーが二人を引き離し、おしゃべりをする
通信しているモデルのピアのみを検出します。
スペースを寄付するかどうかはサーバーが決定します: 空のディレクトリを持つマシンは、
バイトバジェットを提供すると、トラッカーはそれに保持するスライスを割り当てます。
最も希少なものから順に、寄付されたすべてのギガバイトの土地

群れが最も薄い場所。
提供者は群れから自分のスライスを取り出して提供します。
./lumabriserve --model ./slice --join TRACKER:7300 \
--モデル名 tiny_olmoe --donate 5
NAT フロア: 保守者はトラッカーへのアウトバウンド制御接続を 1 つ保持します。
また、直接ダイヤルが失敗すると、バイトが直接ダイヤルを通じて中継されます。後ろの仲間
すべてのホーム NAT はルーター構成なしで機能します。直接ピアツーピア
が第一選択のままです。セルフテストにより、リレー パス上のバイトの同一性が証明されます。
も。プライベート swarm: すべてのマシンで LUMABRI_TOKEN=S を設定します — サービスを提供します
それをトラッカーに渡しますが、トラッカーとすべてのメンテナの両方が拒否します
認証されていない接続なので、トークンはバイトだけでなくバイトも保護します。
インデックス。メンテナーの宣伝
別のマシンの localhost は、トラッカーによってアドレスが修正されます。
つながりが示すもの。
チュートリアル: 5 分で最初の群れ
手元にモデルがありませんか?小さな合成ファイル (python3 + numpy) を生成し、
やってみてください。以下の各ステップは実際のものですが、ほんの小さなものです。
make && make fixture # すべてをビルド + tiny_olmoe/
./lumabriserve --model ./tiny_olmoe
2 番目のターミナルでは、おしゃべりです。からのエンジンバイナリが必要です
コリブリビルド:
./lumabri chat --engines-dir /path/to/colibri/c
順番に見る必要があるものは次のとおりです。
エンジンは、システム上に存在しないディレクトリに対して起動します。
おしゃべりのディスク: 設定とトークナイザーが swarm から到着します。
最初のタッチ。
何か質問してください。最初の答えは遅いです - ネット MB を見てください
作業セットがワイヤーを横切る間にカウンタークライムします。
もう一度尋ねてください。ミラーカルド、ゼロリテ：2番目の答えはから提供されます
~/.lumabri/<model>/cache をローカル速度で実行します。
/swarm は匿名のネットワークを示します。 /model は、他に保持するものをリストします。
サーブターミナルを強制終了してチャットを続けます。温かいミラーは次のように応答します。
すべての仲間が死んだ。つまり

make test の pass 3 は、read ではなく live になりました。
それを成長させるには、別のマシン上の友人とチャットします
--tracker <your-ip>:7300 、またはモデルが存続するようにディスクを寄付します
マシンの電源を切ります。
./lumabriserve --model ./slice --join < あなたの IP > :7300 \
--モデル名 tiny_olmoe --donate 2
トラッカーは、複製が最も少ないファイルを最初にドナーに割り当てます。ドナー
群れからそれらを取り出して提供します。 /swarm をチャットに今すぐ追加
は 2 つのピアを示しています。
2 段階: 最初は 1 台のマシン上ですべてを実行し、次に複数のマシン上で実行します。何もない
以下はシミュレーションです。単一マシンのバージョンでは同じバイナリが実行されます。
同じソケット上で。
モデルのあるマシン上。フェーズ 2 の両方の半分を構築します (おしゃべり
はパッチが適用されたエンジン、エンジンはエキスパート ノード)、モデルを確認します
ネットワークが関与する前にまったく実行されません。
make Phase2-all ENGINE=/path/to/colibri/c && sudo make install
lumabri チャット --local /path/to/model --engines-dir /path/to/colibri/c
次に、群れはまだ 1 台のマシン上にあります。
lumabri key --out swarm # 一度、swarm.key を安全に保ちます
lumabriserve --model /path/to/model --key swarm.key --exec-cache 256
他の人がアクセスできる swarm の場合は、 --advertise <マシンのパブリック IP> を追加します。ピアは与えられたアドレスを公開し、それなしではピアは公開します。
127.0.0.1 — このマシンには適切ですが、他のマシンには役に立ちません。奉仕する
失敗モードはリモートのおしゃべりなので、忘れると大声で言う
これはリレーにフォールバックし、フェーズ 2 をオンにすることはありません。これは次のようになります。
「設定が間違っている」というよりは「遅い」です。
順番に期待します: ハッシュの進行状況 (最初の開始のみ - 大きな場合は数分)
モデル)、ORIGIN: N ファイルの真実に署名し、:7302 で EXEC を提供しています … トラッカーに登録されました。 2 番目の端末で:
LUMABRI_PUBKEY= $( cat swarm.pub ) ルマブリ チャット --tracker 127.0.0.1:7300 --engines-dir /path/to/c

オリブリ/C
探すべきラインは [lumabri] Phase 2 active です。それがなければ、
チャッターはストックエンジンを実行しているため、代わりにエキスパートウェイトをダウンロードします
ピアに実行を依頼する方法 — du -sh ~/.lumabri はもう 1 つの Tell: です。
その上のフェーズ2は、密な部分によって成長し、停止します。
マシンの追加。 7300 ～ 7302 (追加モデルごとに +10) を開きます。
ファイアウォール — クラウド ホスト上、プロバイダーのコンソール内、および ufw 内。
他のすべてのマシンには lumabri と colibri チェックアウトが必要です。
Phase2-all ENGINE=… を作成します。そこから 3 つの役割が生まれます。
/swarm が到着したチャット ショーに参加します。群れが本当にいるという証拠
仕事を運ぶ: 応答が生成されている間に、ドナーを殺害します - あなたはそれを手に入れます
フェイルオーバーラインとトークンは継続し、同一です。
concurrency_test.sh は N 個のチャッターから同じ世代を同時に実行します
そして、最も速いものと最も遅いものの間の広がりを報告します。
「答え」は簡単な質問ですが、「飢える人はいますか」は本当の質問です。
1 つの 6 コア ボックス tiny_olmoe 上で、すべて (サーバー、ピア、クライアント) を共有
それらのコア:
誰も飢えているわけではありません。絶対時間が増加してもスプレッドは横ばいのままです。
これは CPU 競合の様子であり、ロックコンボイの様子ではありません
みたいな。すでに稼働している 6 つのコア上の 4 つのクライアント

[切り捨てられた]

## Original Extract

Run huge MoE models from a swarm of peers, with the colibri engine. Pure C. - JustVugg/lumabri

A while ago I started working on Colibrì to see if it was possible to run huge LLMs on a normal computer. The project grew far beyond what I expected, thanks in large part to the HackerNews community. That led me to a new question: What if we stopped thinking about one computer? This is the idea behind Lumabri. Instead of requiring a single machine to store and run an entire huge model, Lumabri treats a network of normal computers as a shared pool of resources. One machine might provide disk space, another compute, another a different part of the model. If a required block or expert isn’t available locally, the system can retrieve or execute it on a peer. This is particularly interesting for Mixture-of-Experts models. A model can have hundreds of billions of parameters, while only a fraction are activated for each token. Rather than moving huge expert weights over the network, Lumabri can send the small activation to a peer that already has the expert and let it execute it. The goal is for machines to contribute whatever resources they can afford while using the swarm for the rest. The idea is very much inspired by peer-to-peer systems: users are the infrastructure. There are obviously major challenges, especially network latency and security. I’m experimenting with peer verification, SHA-256 verification, signed model state, replica selection, failover, and deterministic execution. Lumabri is still an early experiment. I don’t have a datacenter or a huge GPU cluster, so I’m building it with the hardware I have and trying to find out whether the idea actually makes sense. With Colibrì I asked: Can one normal computer run a huge LLM? With Lumabri I’m asking: What if many normal computers could become one huge computer? Feedback welcome. Repo: https://github.com/JustVugg/lumabri

GitHub - JustVugg/lumabri: Run huge MoE models from a swarm of peers, with the colibri engine. Pure C. · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
JustVugg
/
lumabri
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
27 Commits 27 Commits engine_patches engine_patches expert_engines expert_engines .gitignore .gitignore DEPLOY.md DEPLOY.md DESIGN.md DESIGN.md LICENSE LICENSE Makefile Makefile README.md README.md RESULTS_PHASE2.md RESULTS_PHASE2.md assign_test.sh assign_test.sh chat_proto_test.sh chat_proto_test.sh concurrency_test.sh concurrency_test.sh demo_deepseek.sh demo_deepseek.sh expert_node.c expert_node.c logo.svg logo.svg lumabri.c lumabri.c lumabri_client.h lumabri_client.h lumabri_proto.h lumabri_proto.h lumabri_sha.h lumabri_sha.h lumabri_sign.h lumabri_sign.h lumashim.c lumashim.c lumibri_client.h lumibri_client.h maintainer.c maintainer.c make_tiny_inkling.py make_tiny_inkling.py make_tiny_kimi.py make_tiny_kimi.py make_tiny_olmoe.py make_tiny_olmoe.py phase2_bench.sh phase2_bench.sh phase2_deepseek_test.sh phase2_deepseek_test.sh phase2_glm_test.sh phase2_glm_test.sh phase2_inkling_test.sh phase2_inkling_test.sh phase2_kimi_test.sh phase2_kimi_test.sh phase2_test.sh phase2_test.sh phase3_test.sh phase3_test.sh phase4_test.sh phase4_test.sh phase5_test.sh phase5_test.sh probe_rtt.c probe_rtt.c security_test.sh security_test.sh selftest.sh selftest.sh sign_test.sh sign_test.sh swarm.pub swarm.pub test_shim.c test_shim.c tracker.c tracker.c View all files Repository files navigation
Run huge mixture-of-experts models from a swarm of peers, with the
colibri engine. Pure C, no
dependencies.
One machine shares a model. Any other machine chats with it: nothing is
downloaded up front, the bytes an inference actually touches arrive from the
peer on first use and stay in a local mirror. The second question is served
from local disk at full speed. The engine binary is unmodified.
The founding principle: any machine may join, GPU or not. The engine was
built for CPU and SSD first; a GPU only makes it faster, never different,
and the output is byte-identical either way. A swarm with no GPU at all is
a working swarm. Networks that pool GPUs recruit from the few; lumabri
recruits from everyone.
make
On the machine that has a model (any colibri model directory):
./lumabri serve --model /path/to/model
On the machine that wants to chat (needs a colibri build for the engine):
./lumabri chat --tracker < server-ip > :7300 --engines-dir /path/to/colibri/c
That is all. The first answer is slower while the working set crosses the
network; afterwards the mirror in ~/.lumabri keeps serving even if the
server goes offline.
In the chat: /swarm shows the network live and anonymous (peers are
numbered, never named: model held, GB, bytes served, heartbeat), /model
lists the models on the swarm and switches between them, restarting the
engine on the fly.
For someone who only ever opens the TUI
lumabri
No arguments. It asks for the swarm's address and, once, for the operator's
public key; it finds the engines itself; and it remembers all of it in
~/.lumabri/config , so the second time it is Enter, Enter, and you are in.
That matters more than convenience. Everything it asks for used to be a flag,
and getting one wrong does not produce "invalid argument" — a missing
--engines-dir produces a 299 GB download, and a missing key produces a
model nobody verified. Flags still win when given, so a script never
inherits somebody's saved answers.
Joining: chat, or bring something
lumabri chat asks once, on the way in, and Enter means "just chat" so the
impatient path is one key:
come entri nello sciame?
1 solo chattare non condividi niente
2 chatti e doni disco tieni un pezzo di glm per lo sciame
3 chatti e doni calcolo esegui esperti per gli altri
4 tutti e due
invio = solo chattare
Pick 2 and it asks how many GB (Enter takes a quarter of the free space,
capped), then starts a maintainer with that budget: the tracker assigns it
the least-replicated files first, it pulls them verifying every byte against
the operator's signature, and serves them. Pick 3 and it starts the expert
node for that model's engine. Both run as children of the chat and stop when
you close it — which is the honest lifetime for something offered from a
terminal you have open. A donor that should outlive the session is
lumabri serve --join .
Donating compute needs the model on your disk (an expert node reads the
weights from there), so option 3 is offered only with --model-dir DIR .
Donating disk needs nothing: you start empty and the swarm fills you.
Scripts skip the question: --role chat|disk|compute|all , with --donate GB
and --model-dir DIR .
One tracker is an index, not a model server, so it holds as many models as
you point at it — from one machine or many:
lumabri serve --model /models/glm --port 7300
lumabri serve --model /models/olmoe --join 127.0.0.1:7300 --port 7310
lumabri serve --model /models/deepseek --join 127.0.0.1:7300 --port 7320
Each serve brings its own maintainer (the bytes) and its own expert node
(the compute), and each registers under its model's name. A chatter sees
3 modelli sullo sciame and switches with /model <name> : the engine is
restarted against the new model, and since the engine binary is chosen from
that model's model_type , switching between different architectures works
too — GLM to OLMoE to DeepSeek, one client, one tracker.
Donors are per model as well: a machine can hold a slice of one model and
execute experts for another. The tracker keeps them apart, and a chatter
only ever discovers the peers for the model it is talking to.
Donating space, the server decides: a machine with an empty directory can
offer a byte budget and the tracker assigns it the slice to hold,
rarest-first, so every donated gigabyte lands where the swarm is thinnest.
The donor pulls its slice from the swarm, then serves it.
./lumabri serve --model ./slice --join TRACKER:7300 \
--model-name tiny_olmoe --donate 5
NAT floor: maintainers keep one outbound control connection to the tracker
and, when a direct dial fails, bytes are relayed through it. A peer behind
any home NAT serves with zero router configuration; direct peer-to-peer
stays the first choice. The selftest proves byte identity on the relay path
too. Private swarms: set LUMABRI_TOKEN=S on every machine — serve
passes it to its tracker, and both the tracker and every maintainer refuse
unauthenticated connections, so the token guards the bytes, not just the
index. A maintainer advertising
localhost from another machine gets its address corrected by the tracker to
what the connection shows.
Tutorial: a first swarm in five minutes
No model at hand? Generate the tiny synthetic one (python3 + numpy) and
serve it — every step below is the real thing, just small.
make && make fixture # builds everything + tiny_olmoe/
./lumabri serve --model ./tiny_olmoe
In a second terminal, the chatter. It needs the engine binaries from a
colibri build:
./lumabri chat --engines-dir /path/to/colibri/c
What you should see, in order:
The engine boots against a directory that does not exist on the
chatter's disk: the config and tokenizer arrive from the swarm on
first touch.
Ask something. The first answer is the slow one — watch the net MB
counter climb while the working set crosses the wire.
Ask again. mirror caldo, zero rete : the second answer is served from
~/.lumabri/<model>/cache at local speed.
/swarm shows the network, anonymous; /model lists what else it holds.
Kill the serve terminal and keep chatting: the warm mirror answers with
every peer dead. That is pass 3 of make test , lived instead of read.
To grow it, a friend on another machine chats with
--tracker <your-ip>:7300 , or donates disk so the model survives you
turning your machine off:
./lumabri serve --model ./slice --join < your-ip > :7300 \
--model-name tiny_olmoe --donate 2
The tracker assigns the donor the least-replicated files first; the donor
pulls them from the swarm, then serves them. /swarm in your chat now
shows two peers.
Two stages: everything on one machine first, then more machines. Nothing
below is a simulation — the single-machine version runs the same binaries
over the same sockets.
On the machine with the model. Build both halves of phase 2 ( chatters
is the patched engines, engines is the expert nodes) and check the model
runs at all before any network is involved:
make phase2-all ENGINE=/path/to/colibri/c && sudo make install
lumabri chat --local /path/to/model --engines-dir /path/to/colibri/c
Then the swarm, still on one machine:
lumabri key --out swarm # once, keep swarm.key safe
lumabri serve --model /path/to/model --key swarm.key --exec-cache 256
For a swarm anyone else can reach, add --advertise <the machine's public IP> : peers publish the address they are given, and without it they publish
127.0.0.1 — right for this machine, useless for everyone else. serve
says so loudly if you forget, because the failure mode is a remote chatter
that falls back to the relay and never turns phase 2 on, which reads as
"slow" rather than "misconfigured".
Expect, in order: the hashing progress (first start only — minutes on a big
model), ORIGIN: signed the truth of N files , and serving EXEC on :7302 … registered with tracker . In a second terminal:
LUMABRI_PUBKEY= $( cat swarm.pub ) lumabri chat --tracker 127.0.0.1:7300 --engines-dir /path/to/colibri/c
The line to look for is [lumabri] phase 2 active . Without it the
chatter is running the stock engine and will download expert weights instead
of asking peers to run them — du -sh ~/.lumabri is the other tell: with
phase 2 on it grows by the dense part and stops.
Adding machines. Open 7300-7302 (and +10 per extra model) in the
firewall — on a cloud host, in the provider's console as well as in ufw .
Every other machine needs lumabri and a colibri checkout, then
make phase2-all ENGINE=… . From there the three roles:
/swarm in any chat shows who arrived. The proof that the swarm is really
carrying the work: while a reply is generating, kill a donor — you get one
failover line and the tokens continue, identical.
concurrency_test.sh runs the same generation from N chatters simultaneously
and reports the spread between the fastest and the slowest, because "does it
answer" is the easy question and "does anyone get starved" is the real one.
On one 6-core box, tiny_olmoe, everything (server, peers, clients) sharing
those cores:
Nobody is starved — the spread stays flat while the absolute time grows,
which is what CPU contention looks like and not what a lock convoy looks
like. Four clients on six cores that are alrea

[truncated]
