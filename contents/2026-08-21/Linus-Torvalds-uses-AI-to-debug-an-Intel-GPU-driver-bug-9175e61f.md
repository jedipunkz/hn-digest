---
source: "https://lists.freedesktop.org/archives/dri-devel/2026-August/590630.html"
hn_url: "https://news.ycombinator.com/item?id=49390035"
title: "Linus Torvalds uses AI to debug an Intel GPU driver bug"
article_title: "drm: xe: Kernel-submitted job timed out"
image: ""
author: "signa11"
captured_at: "2026-08-21T16:20:51Z"
capture_tool: "hn-digest"
hn_id: 49390035
score: 6
comments: 0
posted_at: "2026-08-21T15:57:39Z"
tags:
  - hacker-news
  - translated
---

# Linus Torvalds uses AI to debug an Intel GPU driver bug

- HN: [49390035](https://news.ycombinator.com/item?id=49390035)
- Source: [lists.freedesktop.org](https://lists.freedesktop.org/archives/dri-devel/2026-August/590630.html)
- Score: 6
- Comments: 0
- Posted: 2026-08-21T15:57:39Z

## Translation

タイトル: Linus Torvalds が AI を使用して Intel GPU ドライバーのバグをデバッグ
記事のタイトル: drm: xe: カーネルが送信したジョブがタイムアウトしました

記事本文:
drm: xe: カーネルから送信されたジョブがタイムアウトしました
前のメッセージ (スレッド別): drm: xe: カーネルが送信したジョブがタイムアウトしました
次のメッセージ (スレッド別): [PATCH v4 0/5] powervr: MT8173 GPU サポート
メッセージを次の基準で並べ替えます:
[日付]
[ スレッド ]
【件名】
[ 著者 ]
2026 年 8 月 18 日火曜日、23:55、Linus Torvalds
< torvalds at linux-foundation.org > は次のように書きました。
>
> DRM ジョブのタイムアウトの問題を自由に生成できることがわかりました。
>
> [ +1.222075] xe 0000:4b:00.0: [drm] Tile0: GT0: エンジン リセット:
わかりました。今日はマージの合間にこれを追いかけるのに一日中費やしました
ついに「再現可能」になったので、ウィンドウでの作業を行いました。
そして、デバッグがあったとしても、修正は基本的にワンライナーで終わります。
そこに到達するためのセッションはありませんでした。
Xe ドライバは、盗まれたメモリに使用されているメモリ アドレスを取得します。
CCS を使用し、最も近い 128kB 領域に切り上げます。
それは非常に間違っています。xe ドライバーが
次に、128kB 境界にないメモリの部分を *使用*します。そして
HW エンジンも同様に書き込みを行います。
そして、そのメモリが GPU ページ テーブルに使用されなくなると、非常に悪い結果になります。
物事が起こる。
これは、私が経験した時折のランダムな画面破損も説明していると思います。
表示 - ページなど重要なものにメモリが使用されていない場合
テーブルの場合、ランダム ビットマップ メモリなどが破損する「だけ」です。
私は修正をコミット 818bebeb63dd (「drm/xe: 配布しないでください)」としてコミットしました
フラット CCS ストレージを使用可能な VRAM として使用します)。
コミット 37173392741c (「drm/xe/vram: ccs を修正)」から何人かのユーザーを追加
オフセット計算") を cc に送信します。問題はそこにあるからです。
から。 2年前。なぜそれが私にとってそれほど再現可能になったのかわかりません
現在は、ユーザー空間の動作の変更が明らかにトリガーとなっています。
ほとんどすべてのブートで。
ライナス
前のメッセージ (スレッド別): drm: xe: カーネルが送信したジョブがタイムアウトしました
次のメッセージ (スレッドごと):

[パッチ v4 0/5] powervr: MT8173 GPU のサポート
メッセージを次の基準で並べ替えます:
[日付]
[ スレッド ]
【件名】
[ 著者 ]

## Original Extract

drm: xe: Kernel-submitted job timed out
Previous message (by thread): drm: xe: Kernel-submitted job timed out
Next message (by thread): [PATCH v4 0/5] powervr: MT8173 GPU support
Messages sorted by:
[ date ]
[ thread ]
[ subject ]
[ author ]
On Tue, 18 Aug 2026 at 23:55, Linus Torvalds
< torvalds at linux-foundation.org > wrote:
>
> It turns out that now I can generate those drm job timedout issues at will.
>
> [ +1.222075] xe 0000:4b:00.0: [drm] Tile0: GT0: Engine reset:
Ok. I've spent all day today on chasing this down in between merge
window work, since it finally *was* repeatable.
And the fix ends up being basically a one-liner, even if the debug
session to get there was not.
The Xe driver takes the memory address used for the memory stolen for
CCS, and rounds it up to the nearest 128kB area.
And that is very VERY wrong, because it means that the xe driver will
then *use* that part of memory that wasn't at a 128kB boundary. And
the HW engine will too and write to it.
And when that memory happens ot be used for GPU page tables, very bad
things happen.
I bet this also explains some occasional random screen corruption I've
seen - when the memory isn't used for something as important as a page
table, it "only" corrupts random bitmap memory and the like.
I committed the fix as commit 818bebeb63dd ("drm/xe: Don't hand out
the flat CCS storage as usable VRAM").
Adding some people from commit 37173392741c ("drm/xe/vram: fix ccs
offset calculation") to the cc, because that's where the problem came
from. Two years ago. I'm not sure why it became so repeatable for me
now, but some user space behavior change clearly triggered it now
pretty much every single boot.
Linus
Previous message (by thread): drm: xe: Kernel-submitted job timed out
Next message (by thread): [PATCH v4 0/5] powervr: MT8173 GPU support
Messages sorted by:
[ date ]
[ thread ]
[ subject ]
[ author ]
