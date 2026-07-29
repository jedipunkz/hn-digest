---
source: "https://github.com/hardrave/NIGHTRUN"
hn_url: "https://news.ycombinator.com/item?id=49098095"
title: "Show HN: NightRun, bare metal LLM inference, no OS, boots from USB"
article_title: "GitHub - hardrave/NIGHTRUN: Boot your PC straight into an LLM. Rust, UEFI-resident, no operating system underneath. · GitHub"
author: "hardrave"
captured_at: "2026-07-29T15:06:16Z"
capture_tool: "hn-digest"
hn_id: 49098095
score: 2
comments: 0
posted_at: "2026-07-29T14:36:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: NightRun, bare metal LLM inference, no OS, boots from USB

- HN: [49098095](https://news.ycombinator.com/item?id=49098095)
- Source: [github.com](https://github.com/hardrave/NIGHTRUN)
- Score: 2
- Comments: 0
- Posted: 2026-07-29T14:36:24Z

## Translation

タイトル: Show HN: NightRun、ベアメタル LLM 推論、OS なし、USB から起動
記事のタイトル: GitHub - ハードレイブ/NIGHTRUN: PC を LLM で直接起動します。 Rust、UEFI 常駐、その下にオペレーティング システムはありません。 · GitHub
説明: PC を LLM で直接起動します。 Rust、UEFI 常駐、その下にオペレーティング システムはありません。 - ハードレイブ/ナイトラン

記事本文:
GitHub -hardrave/NIGHTRUN: PC を LLM で直接起動します。 Rust、UEFI 常駐、その下にオペレーティング システムはありません。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ハードレイブ
/
ナイトラン
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
64 コミット 64 コミット .cargo .cargo アセット アセット config config クレート クレート docs docs スクリプト スクリプト テスト/フィクスチャ テスト/フィクスチャ ツール ツール Web Web .gitignore .gitignore CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md install.sh install.sh Rust-toolchain.toml Rust-toolchain.toml x86_64-nightrun-uefi.json x86_64-nightrun-uefi.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
USB から起動し、従来のオペレーティング システムなしで実行されるローカル LLM ランタイム。
リアルブート、ワンカット: ロードとプレフィルが高速化され、実際の速度で生成されます (Llama 3.2 1B、QEMU/KVM、8 コア)。
これは奇妙なソフトウェアです。 LLM を直接起動します。
その下に Linux ユーザー空間が隠れているわけではありません。マシンのファームウェアが NightRun を開始します
NightRun は量子化されたモデルを RAM に直接コピーし、独自の端末を RAM 上に描画します。
フレームバッファを使用してチャットします。カーネル、ブラウザ、ネットワーク スタック、ホスト プロセスはありません。
機械が行うことはまさに 1 つです。
Rustで書かれています。 no_std が重要な場合。 USB スティックから通常の x86_64 PC で実行
Raspberry Pi 5 では SD カードから。
インストーラーを実行します。ブート可能なイメージを構築し、それを USB スティックまたは SD カードにフラッシュします。
そのスティックからマシンを起動します。
ファームウェアが NightRun を開始します。メディア上に OS がないため、OS はロードされません。
モデル (1.3 ～ 2.4 GB) は、読み取り中にチェックサムが検証された状態で RAM にストリーミングされます。
保管は密閉されています。それ以降のディスク読み取りは、意図的にハードフォールトとなります。
チャットプロンプトが表示されます。モデルは、オフラインの CPU 上で永久に応答します。
git clone https://github.com/hardrave/NIGHTRUN.git
cd ナイトラン
less install.sh # これから実行しようとしているものを読む
./install.sh
インストーラーの散歩

ターゲットの選択 (x86_64 USB または Pi 5 SD)、モデルの選択、
検証済みのダウンロード (固定リビジョン、SHA-256)、イメージのビルド、およびフラッシュ。
警告。インストーラはリムーバブル メディアをフラッシュします。システムディスクを拒否し、リストのみを表示します
リムーバブルディスク全体デバイスを使用するため、その前に FLASH /dev/sdX をそのまま入力する必要があります。
バイトを書いています。とにかく表示されるデバイス名を読んでください。間違ったディスクをフラッシュすると、
回復可能なエラーではありません。
ハードウェアを使用せずにテストしたい場合は、Linux、Rust (安定版 + 夜間)、QEMU が必要です。
モデルとイメージに約 6 GB の空きディスク。
読み込み中
生成中
花崗岩
クウェン
リポジトリ独自のスクリーンショット ツールを使用して QEMU の実際のブートからキャプチャ
( カーゴ xtask run --shot )。何も演出も模擬もしていません。
ブートとランタイム。単一の UEFI アプリケーション: カスタム レンダラを備えた GOP フレームバッファ
およびビットマップ フォント、キャレット編集とスクロールバックを備えた USB キーボード入力、マルチコア
ファームウェア MP サービスによる推論、シリアル ポート診断。 Pi 5でもそれは
他にファンを駆動するものがないため、RP1 を介してファンを駆動します。
モデルのライフサイクル。 GGUF の検査と変換 ( nrconvert )、専用の .nrm
コンテナ、チャット開始前の完全な RAM 常駐、ファイルの実行中に CRC-32 セクションが検証される
ディスクからストリームし (個別の検証パスはありません)、その後は密閉されたストレージに保存されます。
世代がディスクに触れることはありません。
推論。手書きの量子化カーネル: x86_64 上の AVX2+FMA+F16C、Pi 上の NEON、
スカラー参照実装は両方に保持されます。 Q8_0、Q4_K、Q6_K の重みは以下で使用されます。
場所、逆量子化されたコピーはありません。プロンプト処理はバッチ処理されます (パスごとに最大 64 トークン)。
トークンアットアタイムデコードとビット同一であることが証明されています。生成ループは何も割り当てません。
正しさ。貪欲な出力は、すべてのトークンに対して llama.cpp に対してトークンごとに固定されます。
対応機種f

アミリー、変化するたびに。トークナイザーは生成されたフィクスチャに対してテストされます
公式の Hugging Face トークナイザーから、 apply_chat_template に対するチャット テンプレート。
カーネルの変更によりパリティが崩れた場合、カーネルは間違っています。このルールは本当のバグを捕らえました。
ターゲット
ステータス
ブートメディア
注意事項
x86_64 UEFI
サポートされている
USB
任意の 64 ビット UEFI マシン、セキュア ブートはオフ。 QEMU/OVMF で検証済み。実際のマシンのファームウェアの癖はさまざまですが、ブートレポートは歓迎です
ラズベリーパイ5
サポートされている
microSD
実際の D0 ステッピング 8 GB ボードで検証済み。固定されたソースから構築された UEFI ファームウェア
レガシー BIOS
サポートされていません
UEFIのみ
対応機種
モデル
クオント
サイズ
RAMが必要
実行します
ラマ 3.2 1B 命令
Q8_0
1.3GB
4GB
x86_64、Pi5
ラマ 3.2 3B 命令
Q4_K_M
1.9GB
6GB
x86_64、Pi5
花崗岩 4.1 3B
Q4_K_M
2.0GB
6GB
x86_64、Pi5
Qwen3 4B 命令 2507
Q4_K_M
2.3GB
8GB
x86_64、Pi 5 (8 GB)
3 つのモデル ファミリが実装されており、それぞれの実際の癖が忠実に処理されています。
Llama 3.2 : GQA、隣接ペア RoPE、タイド エンベディング、Llama 3 チャット テンプレート。
Qwen3 : NEOX スタイルのロープ (ハーフスプリット ペア)、ロープ前のヘッドごとの Q/K RMSNorm、BOS なし
トークン、アテンション幅 (4096) が非表示サイズ (2560) よりも広く、出力ヘッドが結合されています。
Granite 4.1 : 高密度トランスのみ。 GQA、SwiGLU、からの 4 つの muP スタイル スカラー
ヘッダー (埋め込み、アテンション、残差、ロジット)。ハイブリッド SSM/MoE Granite バリアントは次のとおりです。
名前付きの理由で変換時に拒否され、実行時に破損しませんでした。
テンソルが Q8_0、Q4_K、Q6_K、または F32 を使用する GGUF は、各テンソルの正確な値で変換されます。
dtype は保持されるため、これらのファミリの Q8_0、Q4_K_M、Q4_K_S、および Q6_K ビルドはすべて機能します。
これは「恣意的な GGUF サポート」ではありません。新しいアーキテクチャ ファミリにはエンジンの作業が必要です。
参照の検証を行うと、コンバーターがそのことを通知します。
ファームウェア(UEFI)
-> NightRun エントリ: フレームバッファ、キーボード、タイマー、MP se

サービス
-> モデルローダー: ストリーミング読み取り + CRC 検証
-> .nrm 検証: ヘッダー、テンソル テーブル、トークナイザー ペイロード
-> アリーナ割り当て: KV キャッシュ + スクラッチ、事前にサイズ設定
-> ストレージが封印されている (後のディスク読み取り = ハード障害)
-> トークナイザー + チャット テンプレート (家族ごと)
-> バッチ化されたプリフィル -> デコードループ -> サンプリング
-> ライブ統計を含むフレームバッファ チャット UI
説明する価値のあるいくつかの選択肢:
なぜ UEFI 常駐なのか。 NightRun は、UEFI ブート サービスではなく、意図的に UEFI ブート サービスに留まります。
ExitBootServices() を呼び出します。それが USB キーボード、ディスプレイ、ディスクを作るものです
読み取りは、カーネルの半分相当のドライバーを配布することなく、どのマシンでも効果的に機能します。
ファームウェアはプラットフォーム層です。その上のすべて (ローダー、フォーマット、トークナイザー、カーネル、
KV キャッシュ、サンプリング、UI) は NightRun コードです。私たちは「従来の OS を使用しない」とは言いますが、「純粋な裸の OS」とは言いません。
なぜなら、クールな響きの主張よりも精度が重要だからです。
モデルが RAM に存在する理由。 1 つのコピーが一度作成され、コピー中に検証されます。後
シールすると、生成パスは誤ってでも I/O を実行できなくなります。また、
パフォーマンスは予測可能: デコード速度はメモリ帯域幅であり、ディスク運ではありません。
なぜカスタムフォーマットなのか。 .nrm はブートに適したコンテナです: 固定ヘッダー、64 バイト
整列されたテンソルがゼロコピービューとして適切に使用され、トークナイザーとチャットテンプレートが埋め込まれています。
メタデータとデータの両方に対する CRC-32。コンバーターは自身の再解析と再チェックサムを実行します。
成功を宣言する前に出力します。不正な形式のファイルは解析時に名前付きエラーを受け取ります。の
パーサーは、切り捨て、ラップされたオフセット、および整列されていないテーブルに対してテストされます。
トークナイザーのパリティが生死に関わるものとして扱われる理由。完璧にロードするモデルですが、
トークン化はほぼ正しく行われ、微妙に間違った出力が生成されますが、これは最もデバッグしにくいものです。
故障モードがあります。したがって、トークン化とテンプレートは公式に固定されています
geによる実装

ネレートされたフィクスチャ、およびユーザー テキストをコントロール トークンにエンコードすることはできません。
戦争の物語を含む長いバージョンは docs/architecture.md にあります。
そしてプロジェクトサイト上で。
docs/benchmarks.md からの測定値、条件付き。
QEMU 行: q35、KVM、8 コア、AVX2 ホスト。単一のスクリプト実行は、次の場合に約 20% 変化します
ホストの負荷。円周率: 本物のボード。
同じマシン、同じ GGUF 上の llama.cpp に対して、貪欲: デコードはパリティです (両方とも
メモリ帯域幅制限)、バッチ化されたプリフィルは 1.15 倍から 1.4 倍の範囲内に収まります。デコードが遅くなる
アテンションはトークンごとに KV キャッシュ全体を読み取るため、コンテキストがいっぱいになります。Granite はからドロップされます。
初期は約 13 トーク/秒、384 トークンの世代では平均 11.6 トーク/秒でした。見出し番号
は常にショートコンテキストの最良のケースです。今ならわかります。
インストーラーがこれらすべてを行います。ピースが欲しい場合:
# 1. モデルを取得します (例: Llama 3.2 1B)
カール -L -o モデル/Llama-3.2-1B-Instruct-Q8_0.gguf \
https://huggingface.co/bartowski/Llama-3.2-1B-Instruct-GGUF/resolve/main/Llama-3.2-1B-Instruct-Q8_0.gguf
# 2. .nrm に変換します (独自の出力を検査、変換、再検証します)
カーゴ ラン --release -p nrconvert -- \
モデル/Llama-3.2-1B-Instruct-Q8_0.gguf モデル/model.nrm
# 3. ブート可能イメージを構築する
カーゴ xtask イメージ --model models/model.nrm # x86_64 -> nightrun.img
カーゴ xtask pi-image --model models/model.nrm # Pi 5 -> nightrun-pi5.img
# (Pi イメージには、固定されたソースからファームウェアを一度ビルドする必要があります:
# scripts/build-rpi5-firmware.sh; docs/rpi5-uefi.md を参照してください)
# 4. ハードウェアに触れる前に QEMU で試してみる
カーゴ xtask run --img --mem 4G --window
# 5. フラッシュ (ターゲットディスクを消去します)
lsblk -d -o 名前、サイズ、モデル、トラン # USB スティックを見つけます
sudo dd if=nightrun.img of=/dev/REPLACE_WITH_USB_DISK bs=4M status=progress conv=fsync
/dev/sdb のようなパーティションではなく、ディスク デバイス全体をフラッシュします。

1. EFI バイナリ自体
カーゴ xtask build でビルドします。夜間ターゲットとカスタムのハードフロートターゲットが必要です。
xタスクハンドル。
./install.sh はガイド付きパスであり、意図的に偏執的なものになっています。そうであるはずです。
リムーバブル メディアを保守的に検出します (USB/SD トランスポート上のディスク全体のみ)。何でも
バッキング / 、 /boot 、 /home または swap は、ストレージ スタック全体をウォークスルーすることによって除外されます。
(LVM と LUKS を含む)、「取り外し可能」フラグからの推測によるものではありません。
デバイスを事前に選択しないでください。空のリストは通常​​の応答です。
選択したデバイス (サイズ、モデル、シリアル) のフィンガープリントを作成し、身元を再検証します。
書き込みの直前なので、再列挙された /dev/sdX はヒットせずに中止されます。
間違ったディスクです。
2 回質問され、2 回目の確認では FLASH /dev/sdX を正確に入力します。はぐれ者
「y」は何もしません。
ダウンロードは Hugging Face リビジョンに固定され、使用前に SHA-256 によって検証されます。
ゲート モデルのトークンはサイレントに読み取られ、ログに記録されたり永続化されたりすることはありません。
書き込み後、メディアを読み取り、SHA-256 ダイジェストを比較します。 「検証済み」とは、
スティックには実際に画像が含まれています。
書き込みを途中で中断すると、メディアが不完全であることがわかります。
テスト スイート ( scripts/installer/tests/run.sh ) は安全性をテストします。

[切り捨てられた]

## Original Extract

Boot your PC straight into an LLM. Rust, UEFI-resident, no operating system underneath. - hardrave/NIGHTRUN

GitHub - hardrave/NIGHTRUN: Boot your PC straight into an LLM. Rust, UEFI-resident, no operating system underneath. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
hardrave
/
NIGHTRUN
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
64 Commits 64 Commits .cargo .cargo assets assets config config crates crates docs docs scripts scripts tests/ fixtures tests/ fixtures tools tools web web .gitignore .gitignore CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md install.sh install.sh rust-toolchain.toml rust-toolchain.toml x86_64-nightrun-uefi.json x86_64-nightrun-uefi.json View all files Repository files navigation
A local LLM runtime that boots from USB and runs without a conventional operating system.
Real boot, one cut: loading and prefill sped up, generation at actual speed (Llama 3.2 1B, QEMU/KVM, 8 cores).
This is weird software. It boots straight into an LLM.
There is no Linux userspace hiding underneath. Your machine's firmware starts NightRun
directly, NightRun copies a quantized model into RAM, draws its own terminal on the
framebuffer, and you chat. No kernel, no browser, no network stack, no host process.
The machine does exactly one thing.
Written in Rust. no_std where it counts. Runs on ordinary x86_64 PCs from a USB stick
and on a Raspberry Pi 5 from an SD card.
Run the installer. It builds a bootable image and flashes it to a USB stick or SD card.
Boot a machine from that stick.
Firmware starts NightRun. No OS loads, because there is no OS on the media.
The model (1.3 to 2.4 GB) streams into RAM with checksums verified during the read.
Storage is sealed. Any later disk read is a hard fault, on purpose.
You get a chat prompt. The model answers on your CPU, offline, forever.
git clone https://github.com/hardrave/NIGHTRUN.git
cd NIGHTRUN
less install.sh # read what you are about to run
./install.sh
The installer walks you through target choice (x86_64 USB or Pi 5 SD), model selection,
a verified download (pinned revision, SHA-256), image build, and flashing.
Warning. The installer flashes removable media. It refuses system disks, lists only
removable whole-disk devices, and demands you type FLASH /dev/sdX verbatim before
writing a byte. Read the device name it shows you anyway. Flashing the wrong disk is
not a recoverable error.
You need Linux, Rust (stable + nightly), QEMU if you want to test without hardware, and
about 6 GB of free disk for a model plus the image.
Loading
Generating
Granite
Qwen
Captured from real boots in QEMU with the repository's own screenshot tooling
( cargo xtask run --shot ). Nothing staged, nothing mocked up.
Boot and runtime. A single UEFI application: GOP framebuffer with a custom renderer
and bitmap fonts, USB keyboard input with caret editing and scrollback, multi-core
inference through firmware MP services, serial-port diagnostics. On the Pi 5 it also
drives the fan through the RP1, because nothing else is around to do it.
Model lifecycle. GGUF inspection and conversion ( nrconvert ), a purpose-built .nrm
container, full RAM residency before chat starts, CRC-32 sections verified while the file
streams from disk (there is no separate verify pass), and sealed storage afterwards.
Generation never touches the disk.
Inference. Hand-written quantized kernels: AVX2+FMA+F16C on x86_64, NEON on the Pi,
scalar reference implementations kept for both. Q8_0, Q4_K and Q6_K weights are used in
place, no dequantized copies. Prompt processing is batched (up to 64 tokens per pass) and
proven bit-identical to token-at-a-time decode. The generation loop allocates nothing.
Correctness. Greedy output is pinned token-for-token against llama.cpp for every
supported model family, on every change. Tokenizers are tested against fixtures generated
from the official Hugging Face tokenizers, chat templates against apply_chat_template .
If a kernel change breaks parity, the kernel is wrong. That rule has caught real bugs.
Target
Status
Boot media
Notes
x86_64 UEFI
supported
USB
any 64-bit UEFI machine, Secure Boot off; validated in QEMU/OVMF; real-machine firmware quirks vary, boot reports welcome
Raspberry Pi 5
supported
microSD
validated on a real D0-stepping 8 GB board; UEFI firmware built from pinned source
Legacy BIOS
not supported
UEFI only
Supported models
Model
Quant
Size
RAM needed
Runs on
Llama 3.2 1B Instruct
Q8_0
1.3 GB
4 GB
x86_64, Pi 5
Llama 3.2 3B Instruct
Q4_K_M
1.9 GB
6 GB
x86_64, Pi 5
Granite 4.1 3B
Q4_K_M
2.0 GB
6 GB
x86_64, Pi 5
Qwen3 4B Instruct 2507
Q4_K_M
2.3 GB
8 GB
x86_64, Pi 5 (8 GB)
Three model families are implemented, each with its real quirks handled faithfully:
Llama 3.2 : GQA, adjacent-pair RoPE, tied embeddings, the Llama 3 chat template.
Qwen3 : NEOX-style rope (half-split pairs), per-head Q/K RMSNorm before rope, no BOS
token, attention width (4096) wider than the hidden size (2560), tied output head.
Granite 4.1 : dense transformer only. GQA, SwiGLU, four muP-style scalars from the
header (embedding, attention, residual, logit). Hybrid SSM/MoE Granite variants are
rejected at conversion with a named reason, not mangled at runtime.
Any GGUF whose tensors use Q8_0, Q4_K, Q6_K or F32 converts with each tensor's exact
dtype preserved, so Q8_0, Q4_K_M, Q4_K_S and Q6_K builds of these families all work.
This is not "arbitrary GGUF support": a new architecture family needs engine work and
reference validation, and the converter will tell you so.
firmware (UEFI)
-> NightRun entry: framebuffer, keyboard, timers, MP services
-> model loader: streaming read + CRC verification
-> .nrm validation: header, tensor table, tokenizer payload
-> arena allocation: KV cache + scratch, sized up front
-> storage sealed (later disk reads = hard fault)
-> tokenizer + chat template (per family)
-> batched prefill -> decode loop -> sampling
-> framebuffer chat UI with live stats
Some choices worth explaining:
Why UEFI-resident. NightRun deliberately stays on UEFI Boot Services instead of
calling ExitBootServices() . That is what makes a USB keyboard, a display, and disk
reads work on effectively any machine without shipping half a kernel's worth of drivers.
Firmware is the platform layer; everything above it (loader, formats, tokenizer, kernels,
KV cache, sampling, UI) is NightRun code. We say "no conventional OS", not "pure bare
metal", because precision matters more than a cooler-sounding claim.
Why the model lives in RAM. One copy, made once, verified during the copy. After
sealing, the generation path cannot perform I/O even by accident. It also makes
performance predictable: decode speed is memory bandwidth, not disk luck.
Why a custom format. .nrm is a boot-friendly container: fixed header, 64-byte
aligned tensors used in place as zero-copy views, tokenizer and chat template embedded,
CRC-32 over both metadata and data. The converter re-parses and re-checksums its own
output before declaring success. Malformed files get named errors at parse time; the
parser is tested against truncations, wrapped offsets, and misaligned tables.
Why tokenizer parity is treated as life-or-death. A model that loads perfectly but
tokenizes almost-correctly produces subtly wrong output, which is the least debuggable
failure mode there is. So tokenization and templates are pinned to the official
implementations by generated fixtures, and user text can never encode into control tokens.
The long version, with the war stories, lives in docs/architecture.md
and on the project site .
Measured numbers from docs/benchmarks.md , conditions attached.
QEMU rows: q35, KVM, 8 cores, AVX2 host; single scripted runs that vary about 20% with
host load. Pi row: real board.
Against llama.cpp on the same machine, same GGUF, greedy: decode is at parity (both are
memory-bandwidth-bound), batched prefill lands within 1.15x to 1.4x. Decode slows as the
context fills, because attention reads the whole KV cache per token: Granite drops from
about 13 tok/s early to 11.6 tok/s averaged over a 384-token generation. Headline numbers
are always the short-context best case; now you know.
The installer does all of this for you. If you want the pieces:
# 1. Get a model (example: Llama 3.2 1B)
curl -L -o models/Llama-3.2-1B-Instruct-Q8_0.gguf \
https://huggingface.co/bartowski/Llama-3.2-1B-Instruct-GGUF/resolve/main/Llama-3.2-1B-Instruct-Q8_0.gguf
# 2. Convert to .nrm (inspects, converts, re-validates its own output)
cargo run --release -p nrconvert -- \
models/Llama-3.2-1B-Instruct-Q8_0.gguf models/model.nrm
# 3. Build a bootable image
cargo xtask image --model models/model.nrm # x86_64 -> nightrun.img
cargo xtask pi-image --model models/model.nrm # Pi 5 -> nightrun-pi5.img
# (the Pi image needs firmware built once from pinned source:
# scripts/build-rpi5-firmware.sh; see docs/rpi5-uefi.md)
# 4. Try it in QEMU before touching hardware
cargo xtask run --img --mem 4G --window
# 5. Flash (THIS ERASES THE TARGET DISK)
lsblk -d -o NAME,SIZE,MODEL,TRAN # find your USB stick
sudo dd if=nightrun.img of=/dev/REPLACE_WITH_USB_DISK bs=4M status=progress conv=fsync
Flash the whole disk device, never a partition like /dev/sdb1 . The EFI binary itself
builds with cargo xtask build ; it needs nightly and a custom hard-float target, which
xtask handles.
./install.sh is the guided path, and it is deliberately paranoid. It should be.
Detects removable media conservatively: whole disks on USB/SD transports only. Anything
backing / , /boot , /home or swap is excluded by walking the full storage stack
(LVM and LUKS included), not by guessing from a "removable" flag.
Never preselects a device. An empty list is a normal answer.
Fingerprints the chosen device (size, model, serial) and re-verifies the identity
immediately before writing, so a re-enumerated /dev/sdX aborts instead of hitting
the wrong disk.
Asks twice, and the second confirmation is typing FLASH /dev/sdX exactly. A stray
"y" does nothing.
Downloads are pinned to a Hugging Face revision and verified by SHA-256 before use.
Tokens for gated models are read silently and never logged or persisted.
After writing, it reads the media back and compares SHA-256 digests. "Verified" means
the stick actually contains the image.
Interrupting mid-write tells you the media is incomplete instead of pretending.
The test suite ( scripts/installer/tests/run.sh ) exercises the safety

[truncated]
