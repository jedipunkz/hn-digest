---
source: "https://mloduchowski.com/-mounted-bitter-fs-better-with-claude/"
hn_url: "https://news.ycombinator.com/item?id=48422375"
title: "Fixing \"unfixable\" 41TB BTRFS by Claude's one-shot"
article_title: "🎉 Mounted — bitter-FS better with Claude\n- mloduchowski.com"
author: "qdotme"
captured_at: "2026-06-06T09:45:42Z"
capture_tool: "hn-digest"
hn_id: 48422375
score: 3
comments: 0
posted_at: "2026-06-06T07:30:37Z"
tags:
  - hacker-news
  - translated
---

# Fixing "unfixable" 41TB BTRFS by Claude's one-shot

- HN: [48422375](https://news.ycombinator.com/item?id=48422375)
- Source: [mloduchowski.com](https://mloduchowski.com/-mounted-bitter-fs-better-with-claude/)
- Score: 3
- Comments: 0
- Posted: 2026-06-06T07:30:37Z

## Translation

タイトル: クロードのワンショットによる「修復不可能な」41TB BTRFS の修復
記事のタイトル: 🎉 マウント — クロードを使ったビター FS の改善
- mloduchowski.com

記事本文:
根っからの科学者、本業のエンジニア、自らの選択による起業家。
🎉 マウント — クロードとのビターFSの方が優れています
TL;DR: ロック解除された iSCSI LUN が 2 つの OS インスタンスによって同時にマウントされました。この状態は 10 か月間気付かれませんでした。標準的な BTRFS 回復ツールはすべて失敗しました。クロード (Opus 4.8、クロード コード セッションで、root として) は、最初の原理から起こったことを再構築し、ディスク上で無傷だが参照されていない 2 番目のトランザクション履歴を見つけ、それを指すようにスーパーブロックに手動でパッチを適用し、本当に破壊された 19 個のメタデータ リーフを再構築し、ファイル システムを再マウントしました。レスキューフラグもエラーもありません。失われたデータ: ゼロ。
完全開示 - このブログ投稿は 90% クロードによって書かれています (ただし、100% は最初に私が読みました。あまり編集しないことにしました)。私はただの観客でした。
クロードがこの約 4 時間のセッション中に作成したツールのリポジトリ - 自己責任で使用してください。最初に読んでください (私はざっと読みました) - ただし、すべて役立つと思われるかもしれません。クロードによるこの投稿の生の草稿も比較のために含まれています。
https://gitlab.defensiblelogic.com/pub/rebtrfs
私の研究室のどこかに 41 TB の BTRFS ファイルシステムがあります。これは、私の NAS によってエクスポートされた iSCSI LUN 上の LUKS2 コンテナ内に存在し、何年ものディスク イメージ、バックアップ、および「後で整理​​する」アーカイブを保持しています。 40 TiB 使用、zstd 圧縮、約 750 万ファイル。これはバックアップ層であり、他のマシンがコピーされる場所です。バックアップのバックアップは存在しますが、数週間前から古いものであり、そこから復元するのは *氷のように * 遅いです。
死因: LUN はロックされておらず、正確に 1 か所にマウントされていると思いました。 ２つに分かれて取り付けられていました。デュアルマウント状態は 10 か月間続きました。1 つのインスタンスは実際の作業を実行し、もう 1 つのインスタンスはアイドル状態ですが生きており、各カーネルは無料であると信じていたものから COW を割り当てていました。

空間、お互いの存在を知りませんでした。
標準ツールのパレード - mount -o ro、usebackuproot、nologreplay、btrfs-find-root、btrfscue chunk-recover、btrfscue super-recover、btrfsrescue zero-log、btrfs Restore - すべて失敗しました。すぐにできるものもあれば、何時間もかけて磨いた後のものもあります。ツールの能力を超えて破損しています。現時点での標準的なアドバイスは「バックアップから復元する」であり、そのオプションは技術的に検討されていました。その代償として、数日間の復元と過去数週間の書き込みが必要になります。最初に手術し、後で降伏する。
したがって、LUN の RAW イメージをローカル RAID-0 スクラッチ ペアに再同期し (10 GigE で約 1 日)、root として Claude Code セッションを開き、問題を多かれ少なかれ次のように述べます。「マシンはあり、コピーはあり、最大 17 TB の空きスクラッチ スペースがあります。それを回復してください。」
最初の動き（何も信頼できない状態）
lsblk / mdstat / LVM recon — コピーを見つけます: 41 T 論理ボリューム、内部に crypto_LUKS。
root の .bash_history を読み取り、元のボリュームで既に試行されたすべてのリカバリ手順を再構築します。これには、LV が破損した LUKS ボリュームのブロックごとの生のイメージを保持していることを証明する pv < /dev/sdf > /dev/dm-0 行が含まれます。
人間に尋ねることは 1 つだけです。LUKS パスフレーズを入力してください。いいえ、セッション中ではなく、 luksOpen /dev/mapper/RAID0-secure_repair として直接
blockdev --setro /dev/mapper/s — 復号化されたデバイスが他のデバイスにアクセスされる前に、カーネル レベルの書き込み保護を行います。
512 G のコピーオンライト ボリュームを基盤とする dm-snapshot オーバーレイを構築します。
最後のステップが盗むステップです。DM スナップショット オーバーレイは、ワンショット リカバリを無制限の再試行に変えます。すべての実験はオーバーレイにヒットします。基礎となるコピーは決して変更されません。あらゆる破壊的なアイデアは、可逆的な実験になります。修復に失敗すると、もう 1 日かけて再同期する代わりに、「dmsetup delete」が必要になります。
lvcreate

-n Cow0 -L 512G RAID0
dmsetup create sr_work --table \
「0 $(blockdev --getsz /dev/mapper/s) スナップショット /dev/mapper/s /dev/RAID0/cow0 P 32」
スーパーブロック自体は正常でした。チェックサムは有効で、世代 24312 の 40 TiB ファイルシステムは健全に見えました。レスキュー マウントは、最初の逆参照で停止しました。
BTRFS エラー: 論理 29245440 ミラーでレベル検証に失敗しました 1 が必要でした 1 が見つかりました 0
BTRFS エラー: 論理ミラー 29245440 でレベル検証に失敗しました 2 が必要でした 1 が見つかりました 0
BTRFS エラー: チャンク ルートの読み取りに失敗しました
スーパーは次のように述べています: 論理 29245440、レベル 1、世代 24287 のチャンク ルート。実際にそこにあるノード (両方の DUP コピー内) はレベル 0 であり、次のとおりです。
親のtransid検証が29245440で失敗しました 必要な数 24287 見つかりました 24860
世代 24860 。スーパーブロックの世界は世代 24312 で終わりますが、その後の 548 トランザクションからのメタデータがここにあります。
その一行がすべての災難だ。ディスク上には 2 つの異なるトランザクション履歴が交互に存在していました。 Git ユーザーの場合: 誰かがすべての参照を破損したコミットに強制プッシュし、実際の履歴はオブジェクト ストアにそのまま残っているリポジトリを想像してください。これはデータの回復ではありません。適切なコミットを見つけて参照を戻すことです。
- 履歴 A — 3 つのスーパーブロック コピーすべてが指定したもの。世代 24312 で停止しました。そのチャンク、csum、デバイス、および fs ツリーのルート: 上書きされた瓦礫。
- 履歴 B — 世代 24866 まで実行され、ボリュームが消滅する数時間前まで実際のファイルが書き込まれていました。完全で、自己一貫性があり、何も参照されません。
各カーネルは、互いのことを何も知らずに、空き領域と思われる領域 (他のインスタンスの新たに書き込まれたメタデータを含む) から割り当てられました。アクティブなインスタンスである B は、A のチャンク ルートを含む、A の最近のツリー ノードをスチームロールしました。ほとんどアイドル状態だがまだ生きている A が、最後の侮辱を加えた。最後のスーパーを書いたのだ。

rblocks — 自身の残骸を示す 3 つの有効なチェックサム付きの道標。
それが、あらゆるツールが失敗した理由です。すべてはスーパーブロックから始まり、チャンク ツリーは論理アドレスから物理ディスク オフセットへのマップです。スーパーには、SYSTEM 領域自体の小さな埋め込みブートストラップ マップが含まれていますが、その外側のすべての論理アドレス、つまり実際のメタデータはすべて、有効なチャンク ルートがなければマップできません。 btrfs-find-root という、失われたルートを見つけることが全体の仕事であるツールでさえ、見つけられないものを読み取ることはできません。すべてのツールは履歴 A の死体を見つめていましたが、履歴 B のファイルシステムはポインタから 50 バイト離れたところに無傷で目に見えずに放置されていました。
困難な方法でチャンクルートを見つける
このファイルシステムでは、チャンク ツリー全体が、アイデンティティ マップされた単一の 8 MiB SYSTEM チャンク内に存在します。つまり、2 つの DUP コピーのうちの最初のコピーについては論理 == 物理的であり、512 個の 16 KiB ノード スロットが可能です。 Claude は、約 60 行の Python スキャナーを作成しました。両方のコピーのすべてのスロットのヘッダーを解析し、そのチェックサム (リカバリ ボックスには python-crc32c がないため、インラインで実装された crc32c テーブル) を検証し、すべての有効なノードの所有者/世代/レベルを記録します。
SYSTEM エリアは、両方の履歴のチャンク ツリー ノードの墓場であることが判明しました。その中には、履歴 B の完全なレベル 1 チャンク ルートが含まれていました。
logical=26361856 gen=24864 level=1 nritems=269 [両方のコピーが同意、csum OK]
オーバーレイのスーパーブロック ( chunk_root 、 chunk_root_generation ) を再ポイントし、スーパーの crc32c を再計算すると、マウントの失敗がさらに深くなり、この業界ではこれが進歩と感じられます。新しいエラー: デバイス サイズの不一致。これは、履歴 B がファイル システムのサイズを 40 TiB から 41 TiB LUN 全体を満たすように静かに変更したためです (非常に NAS アプライアンスの動作であり、どのインスタンス B であるかを識別する別のフィンガープリント)。

total_bytes にもパッチを適用すると、チャンク ツリーがロードされます。 fs ツリーはまだ A のままで、死んだままです。しかし、ディスクのマップは戻ってきており、それによって他のすべてのロックが解除されます。
次の問題: 履歴 B のツリー ルートはディスク上に存在しますが、それらを指すものは何もありません。 btrfs-find-root が標準的な答えであり、チャンク ツリーを復元すると、実際に実行されました。悪い結果として、10 分間で 16 GB の transid-warning スパム、59 GB の RSS、収束の見通しはありません。クロードはそれを殺し、もっと単純なことをしました。
B のチャンク ツリーによれば、すべてのメタデータはちょうど 62 GiB の DUP チャンクに存在します。つまり、62 GiB をすべて順番に読み取り、16 KiB のブロック ヘッダーごとに解析し、カタログ (4,063,210 個のメタデータ ノード、両方の履歴のすべてのツリー ノード) を論理、物理、所有者、世代、レベルの CSV として書き込みます。メカニカル RAID-0 での数分間のシーケンシャル I/O。
ルート ツリー ノードのカタログをクエリします。履歴 B の最終ルート ツリーは、世代 24866 、レベル 0、11 項目です。 (チャンク ルートが 24864 と言っているのに、なぜ 24866 なのでしょうか? ツリー ルートの世代は、*そのツリー* を変更した最後のトランザクションです。チャンク ツリーは B の最後の 2 つのコミットを終了しました。予期されていますが、疑わしいものではありません。) そのルート ツリー リーフを生のまま解析し、それが参照しているすべてのツリー ルートをカタログに対してクロスチェックします。
エクステント: @47105846919168 gen=24866 level=3 -> OK
DEV: @47105193230336 gen=24864 level=1 -> OK
FS: @47105458651136 gen=24866 level=3 -> OK
CSUM: @47105458716672 gen=24866 level=3 -> OK
FREE_SPACE: @47105846984704 gen=24866 level=1 -> OK
履歴 B の最終コミットはディスク上で完了しました。このファイルシステムの唯一の問題は、スーパーブロックがその存在を認めなかったことです。
オーバーレイのスーパーブロック上の root 、 Generation 、 root_level を再ポイントし、次のようにします。
# mount -t btrfs -o ro /dev/mapper/sr_work /mnt/sr_work
マウント出口: 0
dmesg: (何も)
プレーンな読み取り専用マウント。 rescue=all ではなく、usebac ではありません

kuproot — 救済オプションはまったくありません。カーネルには何の不満もありませんでした。 🎉取り付けました。
そこにはすべてが存在していました。40 TiB のディスク イメージ、squashfs アーカイブ、ホーム ディレクトリ、ボリュームがなくなった同じ夜に書き込まれたバックアップも含まれていました。
信頼しますが、検証します（検証者も信頼しません）
マウントは回復ではありません。興味深いのは検証パスです。
- プローブ スイープ : ライブ チェックサムを使用して、それぞれ 5 つのオフセットで 7,497,112 ファイルを読み取ります。エラーゼロ。
-btrfs チェックでは 6,480 万件のエラーが報告されました。クロードは、苦情があったエクステントを 1 つ選択し、カーネルの FIEMAP を通じて反対尋問することで判決を下しました。エクステント レコードとファイル マッピングの両方が存在し、同意されました。 6,480 万件の「参照数の不一致」エラーは誤検知でした。チェックのユーザー空間トラバーサルが早い段階で 1 つの不良ノードにヒットし、無意味な事態にまで連鎖しました。 (mkfs 時代の 2 つの古い空きスペース ツリー エントリは本物であり、表面的なものであり、後で注目されます。)
- 死んだメタデータ内の inode ctimes からのタイムライン フォレンジック: デュアル マウントは 2025 年 4 月 17 日、コピー ジョブの途中で開始されました。履歴 A はほぼ間違いなく、このマシンの古い iSCSI マウントであり、約 25 の atime ノイズの野良コミットで 10 か月間アイドル状態でした。そのうちの最後の履歴は、B の死後に書き込まれ、最後の横たわっていたスーパーブロックを生成しました。そして、フォーク日の差分はきれいに戻ってきました。フォークが解決された後に A がタッチした 787 個の i ノードは、それぞれ B のツリーに存在する 365 個の異なる名前に解決されました。 B を選択しても何も犠牲になりませんでした。
その後、最終的な意見を持つ唯一の検証者であるカーネルが、時期尚早の勝利に拒否権を発動しました。すべてのユーザー空間のウォークが通過したメタデータ ブロックで EIO された 1 つのファイルの完全な読み取り。 btrfs-progs のツリー ウォーカーは活動を続けていますが、「transid 障害を無視した」後、静かに報告を停止していることが判明しました。チェックツールは、何もないと6,480万回叫んでいた

実際の被害については沈黙を守った。
そこで、クロードはウォーカーを完全に信頼するのをやめ、独自のウォーカーを作成しました。B のルートから現在のすべてのツリーのすべてのノードを訪問し、すべての親→子のエッジ (予想されるバイト数、世代、所有者、レベル) を 400 万ノードのカタログに対して検証します。
これは、10 か月間のデュアル マウント操作における完全かつ網羅的な損傷リストです。19 個の 16 KiB リーフ (メタデータの 0.0005%) はすべて、A のゾンビがコミットした履歴によって破壊され、すべて 1 つの ~6 MB 割り当て近傍にクラスター化されており、それらはすべて 2025 年 4 月のバランス実行で書き込まれたメタデータです。
- 10 個の fs ツリー リーフ: 正確に 2 つの古い VM イメージ、 theia/ultrasound.raw および theia/docker-runner-VM-uma0.raw に属する純粋なエクステント マップ範囲 (ファイル オフセットの約 230 MiB)。 i ノード項目やディレクトリ エントリはなく、ブロック マップだけです。
- 9 csum-tree リーフ: それらのファイル (および 1 つの squashfs) の再配置データの ~130 MB をカバーするチェックサム。
41 TB の木に葉を接ぎ木して戻す
これは私がパーティーでもう一度話すことになる部分です。枯れ葉たちが説明したデータはまだディスク上にあり、割り当てられたままであり、その説明だけが消えました。そして、BTRFS は逆インデックスを保持します。エクステント ツリー (完全な状態 - エッジごとに検証済み) は、その所有者に名前を付けるすべてのデータ エクステントの後方参照を保存します: ツリー、inode、ファイル オフセット (通常の un-cl の場合)

[切り捨てられた]

## Original Extract

Scientist by Heart , Engineer by Trade , Entrepreneur by Choice .
🎉 Mounted — bitter-FS better with Claude
TL;DR: An unlocked iSCSI LUN got mounted by two OS instances at once - a condition that went unnoticed for ten months. Every standard BTRFS recovery tool failed. Claude (Opus 4.8, in a Claude Code session, as root) reconstructed what happened from first principles, found an intact-but-unreferenced second transaction history on the disk, hand-patched the superblocks to point at it, rebuilt the 19 metadata leaves that were genuinely destroyed, and remounted the filesystem: mount -o ro; no rescue flags, no errors. Data lost: zero .
Full disclosure - this blog post, 90% written by Claude (but 100% read by me first, I decided not to edit it much). I was just a spectator.
Repo with the tools that Claude wrote during this ~4h session - please use at your own risk, read them first (I did skim them) - but you might find them all useful. Claude's raw draft of this post included for comparison too.
https://gitlab.defensiblelogic.com/pub/rebtrfs
Somewhere in my lab there is a 41 TB BTRFS filesystem. It lives inside a LUKS2 container, on an iSCSI LUN exported by my NAS, and holds years of disk images, backups, and "I'll sort this out later" archives. 40 TiB used, zstd-compressed, about 7.5 million files. It's the backup tier — the place other machines get copied to. A backup of the backup does exist, but it's a few weeks out of date and *glacially* slow to restore from.
Cause of death: the LUN wasn't locked, and I assumed it was mounted in exactly one place. It was mounted in two. The dual-mount condition persisted for ten months — one instance doing the real work, the other sitting idle but alive, each kernel COW-allocating from what it believed was free space, neither knowing the other existed.
The standard tool parade - mount -o ro,usebackuproot,nologreplay, btrfs-find-root, btrfs rescue chunk-recover, btrfs rescue super-recover, btrfs rescue zero-log, btrfs restore — all failed. Some instantly, some after hours of grinding. Corrupted beyond the tools' ability. The standard advice at this point is "restore from backup," and that option was technically on the table; at the cost of a multi-day restore and the last few weeks of writes. Surgery first, surrender later.
So: re-sync a raw image of the LUN onto a local RAID-0 scratch pair (about a day over 10 GigE), open a Claude Code session as root, and state the problem more or less as: "you have the machine, you have a copy, you have ~17 TB of free scratch space. Recover it."
First moves (in which nothing is trusted)
lsblk / mdstat / LVM recon — locate the copy: a 41 T logical volume, crypto_LUKS inside.
Read root's .bash_history and reconstruct every recovery step already tried on the original — including the pv < /dev/sdf > /dev/dm-0 line proving the LV held a raw, block-for-block image of the damaged LUKS volume.
Ask the human for exactly one thing: type the LUKS passphrase. No, not in a session, directly as luksOpen /dev/mapper/RAID0-secure_repair s
blockdev --setro /dev/mapper/s — kernel-level write protection on the decrypted device, before anything else gets a chance to touch it.
Build a dm-snapshot overlay backed by a 512 G copy-on-write volume.
That last step is the one to steal: a dm-snapshot overlay turns a one-shot recovery into unlimited retries. All experiments hit the overlay; the underlying copy never changes; every destructive idea becomes a reversible experiment. A failed repair costs a `dmsetup remove` instead of another day of resync.
lvcreate -n cow0 -L 512G RAID0
dmsetup create sr_work --table \
"0 $(blockdev --getsz /dev/mapper/s) snapshot /dev/mapper/s /dev/RAID0/cow0 P 32"
The superblock itself was fine — checksum valid, a healthy-looking 40 TiB filesystem at generation 24312. The rescue mount died at the very first dereference:
BTRFS error: level verify failed on logical 29245440 mirror 1 wanted 1 found 0
BTRFS error: level verify failed on logical 29245440 mirror 2 wanted 1 found 0
BTRFS error: failed to read chunk root
The super says: chunk root at logical 29245440, level 1, generation 24287. The node actually there (in both DUP copies) is level 0, and:
parent transid verify failed on 29245440 wanted 24287 found 24860
Generation 24860 . The superblock's world ends at generation 24312 — yet here is metadata from 548 transactions later.
That one line is the whole disaster. There were two divergent transaction histories interleaved on the disk. Git users: picture a repo where someone force-pushed every ref to a corrupted commit, while the real history sits intact in the object store. This isn't data recovery — it's finding the good commits and pointing the refs back.
- History A — the one all three superblock copies pointed at. Stopped at generation 24312. Its chunk, csum, device, and fs tree roots: overwritten rubble.
- History B — ran on to generation 24866, with real file writes up to hours before the volume's death. Complete, self-consistent... and referenced by nothing .
Each kernel, knowing nothing of the other, allocated from what it believed was free space — which included the other instance's freshly written metadata. B, the active instance, steamrolled A's recent tree nodes, including A's chunk root. A, mostly idle but still alive, delivered the final insult: it wrote the last superblocks — three valid, checksummed signposts pointing into its own wreckage.
And that's why every tool failed. Everything starts from the superblock, and the chunk tree is the map from logical addresses to physical disk offsets. The super carries a tiny embedded bootstrap map for the SYSTEM area itself, but every logical address outside it — which is to say, all the actual metadata — is unmappable without a valid chunk root. Even btrfs-find-root , the tool whose whole job is finding lost roots, can't read what it can't locate. Every tool was staring at history A's corpse while history B's filesystem sat intact, invisible, fifty bytes of pointer away.
Finding the chunk root the hard way
On this filesystem the entire chunk tree lives in a single 8 MiB SYSTEM chunk that is identity-mapped — logical == physical, for the first of its two DUP copies anyway — 512 possible 16 KiB node slots. Claude wrote a ~60-line Python scanner: parse every slot's header in both copies, verify its checksum (crc32c table implemented inline, because of course the recovery box has no python-crc32c), and record owner / generation / level for every valid node.
The SYSTEM area turned out to be a graveyard of chunk-tree nodes from both histories — and among them, intact level-1 chunk roots from history B:
logical=26361856 gen=24864 level=1 nritems=269 [both copies AGREE, csum OK]
Repoint the overlay's superblock — chunk_root , chunk_root_generation , recompute the super's crc32c — and the mount failure moves deeper , which in this business is what progress feels like. The new error: device size mismatch, because history B had quietly resized the filesystem from 40 TiB to fill the whole 41 TiB LUN (very NAS-appliance behavior, and another fingerprint identifying which instance B was). Patch total_bytes too, and the chunk tree loads. The fs tree is still A's, still dead — but the map of the disk is back, and that unlocks everything else.
Next problem: history B's tree roots exist on disk, but nothing points at them. btrfs-find-root is the canonical answer, and with the chunk tree restored it actually ran — badly: 16 GB of transid-warning spam in ten minutes, 59 GB of RSS, no convergence in sight. Claude killed it and did something simpler.
B's chunk tree says all metadata lives in exactly 62 GiB of DUP chunks. So: read all 62 GiB sequentially, parse every 16 KiB block header, and write a catalog — 4,063,210 metadata nodes , every tree node from both histories, as a CSV of logical, physical, owner, generation, level. A few minutes of sequential I/O on mechanical RAID-0.
Query the catalog for root-tree nodes: history B's final root tree is at generation 24866 , level 0, eleven items. (Why 24866 when the chunk root said 24864? A tree root's generation is just the last transaction that modified *that tree* — the chunk tree sat out B's final two commits. Expected, not suspicious.) Parse that root-tree leaf raw and cross-check every tree root it references against the catalog:
EXTENT: @47105846919168 gen=24866 level=3 -> OK
DEV: @47105193230336 gen=24864 level=1 -> OK
FS: @47105458651136 gen=24866 level=3 -> OK
CSUM: @47105458716672 gen=24866 level=3 -> OK
FREE_SPACE: @47105846984704 gen=24866 level=1 -> OK
History B's final commit was complete on disk . The only thing wrong with this filesystem was that no superblock admitted it existed.
Repoint root , generation , root_level on the overlay's superblock, and:
# mount -t btrfs -o ro /dev/mapper/sr_work /mnt/sr_work
mount exit: 0
dmesg: (nothing)
A plain read-only mount . Not rescue=all , not usebackuproot — no rescue options at all. The kernel had no complaints. 🎉 Mounted.
Everything was there: 40 TiB of disk images, squashfs archives, home directories — including backups written the same evening the volume died.
Trust, but verify (and don't trust the verifiers either)
A mount is not a recovery. The verification pass is where it got interesting:
- Probe sweep : read 7,497,112 files at five offsets each, with checksums live. Zero errors.
- btrfs check reported 64.8 million errors . Claude adjudicated by picking one complained-about extent and cross-examining it through the kernel's FIEMAP: both the extent record and the file mapping existed and agreed. The 64.8 million "referencer count mismatch" errors were false positives — check's userspace traversal had hit one bad node early and cascaded into nonsense. (Two stale free-space-tree entries from mkfs days were real, cosmetic, and noted for later.)
- Timeline forensics from inode ctimes inside dead metadata: the dual mount began April 17, 2025 , mid-copy-job. History A was almost certainly this very machine's stale iSCSI mount, idling for ten months with ~25 stray commits of atime noise — the last of them, written after B's death, produced those final lying superblocks. And the fork-day diff came back clean: the 787 inodes A touched after the fork resolve to 365 distinct names, every one of which is present in B's tree. Choosing B sacrificed nothing at all.
Then the kernel — the only verifier whose opinion is final — vetoed premature victory. A full read of one file EIO'd on a metadata block that every userspace walk had sailed past. It turns out btrfs-progs' tree walkers keep going but quietly stop reporting after "Ignoring transid failure." The check tool had screamed 64.8 million times about nothing and stayed silent about actual damage.
So Claude stopped trusting walkers entirely and wrote its own: visit every node of every current tree from B's roots and verify every parent→child edge — expected bytenr, generation, owner, level — against the 4-million-node catalog.
That is the complete, exhaustive damage list for ten months of dual-mounted operation: nineteen 16-KiB leaves — 0.0005% of metadata — all clobbered by history A's zombie commits, all clustered in one ~6 MB allocation neighborhood, all of them metadata that one April-2025 balance run had written:
- 10 fs-tree leaves : pure extent-map ranges (~230 MiB of file offsets) belonging to exactly two old VM images, theia/ultrasound.raw and theia/docker-runner-VM-numa0.raw . No inode items, no directory entries — just block maps.
- 9 csum-tree leaves : checksums covering ~130 MB of those files' (plus one squashfs's) relocated data.
Grafting leaves back onto a 41 TB tree
Here's the part I'll be retelling at parties. The data those dead leaves described was still on disk, still allocated — only its description died. And BTRFS keeps a reverse index: the extent tree (intact — verified edge by edge) stores a back-reference for every data extent naming its owner: tree, inode, file offset (for ordinary un-cl

[truncated]
