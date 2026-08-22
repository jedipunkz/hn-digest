---
source: "https://github.com/torvalds/linux/commit/818bebeb63dd6bf5f4e07e145f6cdbace520a34c"
hn_url: "https://news.ycombinator.com/item?id=49395262"
title: "Linus: \"A debug session from hell, enormously helped by an AI\""
article_title: "drm/xe: Don't hand out the flat CCS storage as usable VRAM · torvalds/linux@818bebe · GitHub"
image: "https://opengraph.githubassets.com/22087a98ed668cae2b949a46863384f107467421a4bea6f1340abf5c42a38490/torvalds/linux/commit/818bebeb63dd6bf5f4e07e145f6cdbace520a34c"
author: "dnw"
captured_at: "2026-08-22T00:39:13Z"
capture_tool: "hn-digest"
hn_id: 49395262
score: 1
comments: 0
posted_at: "2026-08-22T00:08:47Z"
tags:
  - hacker-news
  - translated
---

# Linus: "A debug session from hell, enormously helped by an AI"

- HN: [49395262](https://news.ycombinator.com/item?id=49395262)
- Source: [github.com](https://github.com/torvalds/linux/commit/818bebeb63dd6bf5f4e07e145f6cdbace520a34c)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T00:08:47Z

## Translation

タイトル: ライナス: 「地獄のデバッグ セッション、AI に大いに助けられた」
記事のタイトル: drm/xe: フラット CCS ストレージを使用可能な VRAM として配布しない · torvalds/linux@818bebe · GitHub
説明: Linux カーネル ソース ツリー。 GitHub でアカウントを作成して、torvalds/linux の開発に貢献してください。

記事本文:
drm/xe: フラット CCS ストレージを使用可能な VRAM として配布しない · torvalds/linux@818bebe · GitHub
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
トーバルズ
/
リナックス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ファイルを参照する 履歴のこの時点のリポジトリを参照する ファイルを参照する drm/xe がコミットされます: フラット CCS ストレージを使用可能な VRAM として配布しないでください get_ flat_ccs_offset() は、フラット CCS ストレージのベースを
ハードウェアを使用して、有効な L3 ノードの数によってスケールし、四捨五入します。
レス

最大 128K。そのオフセットより下にあるものはすべて、
使用可能なメモリとしての VRAM アロケータ。
「使用可能なメモリがここで終了する」という制限を切り上げてパブリッシュする
実際の底と空きメモリとして丸められた底の間にあるものは何でも、
そしてそのメモリは圧縮ハードウェアに属します。スケーリングされた値
128K にアライメントする理由はありません。16 GiB の Battlemage G21 では、
ではありません:
フラット CCS ベース: 生の 0x3fafff800、丸められた 0x3fb000000
したがって、ページ 0x3fafff000 の最後の 2 KiB は、アロケーターの CCS ストレージです。
プール。そこに割り当てられたものはすべて、その末尾が上書きされます。
圧縮ハードウェアは、ページテーブルエントリやバッファオブジェクトを必要としません。
また、それを行うための GPU 送信はなく、ユーザー空間が存在する前に実行されます。
このマシンでは、Mesa VM のレベル 3 ページ テーブルがそのページに配置されました。
すべてのコールドブート。コンポジターをカバーするエントリが失われました
バッチバッファヒープのため、コンポジタの最初の送信でフェッチエラーが発生しました
そのバッチと gdm が永久に再起動しました。それ以外の場合は黒い画面が表示されます。
作業機械。次の VM のページが表示されるため、gdm を再起動するとそれがクリアされました
テーブルが別の場所に割り当てられていました。
代わりに、アロケータが動作するページ サイズに切り捨てます。
正確に 1 ページを除外するマシン。
その後、予約されたページを読むと、何が書き込まれていたのかがわかります。
[369] 0xcccc000000000000
[371] 0xcc77000000000000
[373] 0xcccc000000000000
[375] 0xcc77000000000000
圧縮メタデータ、16 あたり 2 バイト、ドライバーの場所に配置
メモリを配布するために使用されます。
これを捉えるべきアサーションは、オフセットを比較します。
GSMBASE - ccs_size が等しい場合。この値は 128K でアライメントされているため、
底が一致しない場合、切り上げオフセットと正確に一致します。
整列 -
[切り捨てられた]
ファイル ツリーを展開する ファイル ツリーを折りたたむ 差分表示設定を開く フィルター オプション ドライバー/GPU

/drm/xe xe_vram.c
ファイル ツリーを展開する ファイル ツリーを折りたたむ 差分表示設定を開く ファイルを折りたたむ drivers/gpu/drm/xe/xe_vram.c
ファイル名をクリップボードにコピー すべての行を展開: drivers/gpu/drm/xe/xe_vram.c + 18 - 5 行変更: 18 追加 & 5 削除 元のファイル行番号 差分行番号 差分行変更 @@ -89,12 +89,25 @@ static int get_ flat_ccs_offset(struct xe_gt *gt, u64 tile_size, u64 *poffset) 89 89 オフセット = オフセット_hi << 32 ; /* HW ビュー ビット 39:32 */ 90 90 offset |= offset_lo << 6 ; /* HW ビュー ビット 31:6 */ 91 91 offset *= num_enabled ; /* SW ビューに変換 */ 92 - offset =round_up ( offset , SZ_128K ); /* SW は最も近い 128K に切り上げる必要があります */ 93 92
94 - /* ホールは予想されません */ 95 - xe_assert_msg ( xe , offset == ( xe_mmio_read64_2x32 ( & gt_to_tile ( gt ) -> mmio , GSMBASE ) - 96 - ccs_size ), 97 - "CCS と GSM の間にホールがあります。\n" ); 93 + /* 94 + * このオフセットより下のすべては VRAM 95 + * アロケーターに渡されるため、96 + * 圧縮ハードウェアが所有する *最初の * アドレスを切り捨てたものでなければなりません。切り上げて 97 + * CCS ストレージを空きメモリとして公開します。 98 + */ 99 + offset =round_down ( offset , SZ_4K ); 100 + 101 + /* 102 + * CCS ストレージは GSM で実行してはなりません。古いチェックでは、103 + * オフセットを GSMBASE - ccs_size と等しいかどうか比較していましたが、104 + * では失敗するはずがありませんでした。その値は 128K にアライメントされているため、ベースが 128K 106 + * にアライメントされていない場合でも、切り上げられたオフセットと 105 + * が一致しました。まさにこのケースがこれで修正されます。 107 + */ 108 + xe_assert_msg ( xe , offset + ccs_size <= 109 + xe_mmio_read64_2x32 ( & gt_to_tile ( gt ) -> mmio , GSMBASE ), 110 + "CCS は GSM と重複します。\n" ); 98 111 } else { 99 112 reg = xe_gt_mcr_unicast_read_any ( gt , XEHP_FLAT_CCS_BASE_ADDR ); 100 113 オフセット = ( u64 ) REG_FIELD_GET ( XEHP_FLAT_CCS_PTR , reg ) * SZ_64K

;コミットコメント0件
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Linux kernel source tree. Contribute to torvalds/linux development by creating an account on GitHub.

drm/xe: Don't hand out the flat CCS storage as usable VRAM · torvalds/linux@818bebe · GitHub
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
torvalds
/
linux
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Browse files Browse the repository at this point in the history Browse files torvalds committed drm/xe: Don't hand out the flat CCS storage as usable VRAM get_flat_ccs_offset() reads the base of the flat CCS storage from the
hardware, scales it by the number of enabled L3 nodes, and rounds the
result up to 128K. Everything below that offset is then handed to the
VRAM allocator as usable memory.
Rounding a limit that means "usable memory ends here" upwards publishes
whatever lies between the real base and the rounded one as free memory,
and that memory belongs to the compression hardware. The scaled value
has no reason to be 128K aligned, and on a Battlemage G21 with 16 GiB it
is not:
flat CCS base: raw 0x3fafff800, rounded 0x3fb000000
so the last 2 KiB of page 0x3fafff000 is CCS storage, in the allocator's
pool. Whatever is allocated there gets that tail overwritten by the
compression hardware, which needs no page-table entry, no buffer object
and no GPU submission to do it, and does it before userspace exists.
On this machine a Mesa VM's level-3 page table landed on that page on
every cold boot. It lost the entry covering the compositor's
batch-buffer heap, so the compositor's first submission faulted fetching
its batch and gdm restarted it forever: a black screen on an otherwise
working machine. Restarting gdm cleared it because the next VM's page
tables were allocated somewhere else.
Round down instead, to the page size the allocator works in. On this
machine that excludes exactly one page.
Reading the reserved page afterwards shows what had been writing it:
[369] 0xcccc000000000000
[371] 0xcc77000000000000
[373] 0xcccc000000000000
[375] 0xcc77000000000000
compression metadata, two bytes per sixteen, sitting where the driver
used to hand out memory.
The assertion that should have caught this compares the offset against
GSMBASE - ccs_size for equality. That value is 128K aligned, so it
agrees with the rounded-up offset precisely when the base is not
aligned - the
[truncated]
Expand file tree Collapse file tree Open diff view settings Filter options drivers/gpu/drm/xe xe_vram.c
Expand file tree Collapse file tree Open diff view settings Collapse file ‎ drivers/gpu/drm/xe/xe_vram.c ‎
Copy file name to clipboard Expand all lines: drivers/gpu/drm/xe/xe_vram.c + 18 - 5 Lines changed: 18 additions & 5 deletions Original file line number Diff line number Diff line change @@ -89,12 +89,25 @@ static int get_flat_ccs_offset(struct xe_gt *gt, u64 tile_size, u64 *poffset) 89 89 offset = offset_hi << 32 ; /* HW view bits 39:32 */ 90 90 offset |= offset_lo << 6 ; /* HW view bits 31:6 */ 91 91 offset *= num_enabled ; /* convert to SW view */ 92 - offset = round_up ( offset , SZ_128K ); /* SW must round up to nearest 128K */ 93 92
94 - /* We don't expect any holes */ 95 - xe_assert_msg ( xe , offset == ( xe_mmio_read64_2x32 ( & gt_to_tile ( gt ) -> mmio , GSMBASE ) - 96 - ccs_size ), 97 - "Hole between CCS and GSM.\n" ); 93 + /* 94 + * Everything below this offset is handed to the VRAM 95 + * allocator, so it has to be the *first* address the 96 + * compression hardware owns, rounded down. Rounding it up 97 + * publishes CCS storage as free memory. 98 + */ 99 + offset = round_down ( offset , SZ_4K ); 100 + 101 + /* 102 + * CCS storage must not run into GSM. The old check compared 103 + * the offset against GSMBASE - ccs_size for equality, which 104 + * could not fail: that value is 128K aligned, so it agreed 105 + * with the rounded-up offset even when the base was not 128K 106 + * aligned - exactly the case this fixes. 107 + */ 108 + xe_assert_msg ( xe , offset + ccs_size <= 109 + xe_mmio_read64_2x32 ( & gt_to_tile ( gt ) -> mmio , GSMBASE ), 110 + "CCS overlaps GSM.\n" ); 98 111 } else { 99 112 reg = xe_gt_mcr_unicast_read_any ( gt , XEHP_FLAT_CCS_BASE_ADDR ); 100 113 offset = ( u64 ) REG_FIELD_GET ( XEHP_FLAT_CCS_PTR , reg ) * SZ_64K ; 0 commit comments
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
