---
source: "https://github.com/anat0m1a/liblzraven"
hn_url: "https://news.ycombinator.com/item?id=49112695"
title: "I asked Claude to reimplement Apple's LZRAVEN codec in C, conformance-tested"
article_title: "GitHub - anat0m1a/liblzraven: AI-written, conformance-tested C11 implementation of Apple's LZRAVEN codec · GitHub"
author: "anat0m1a"
captured_at: "2026-07-30T17:15:51Z"
capture_tool: "hn-digest"
hn_id: 49112695
score: 2
comments: 0
posted_at: "2026-07-30T17:01:09Z"
tags:
  - hacker-news
  - translated
---

# I asked Claude to reimplement Apple's LZRAVEN codec in C, conformance-tested

- HN: [49112695](https://news.ycombinator.com/item?id=49112695)
- Source: [github.com](https://github.com/anat0m1a/liblzraven)
- Score: 2
- Comments: 0
- Posted: 2026-07-30T17:01:09Z

## Translation

タイトル: クロードに Apple の LZRAVEN コーデックを C で再実装するよう依頼しました (適合性テスト済み)
記事のタイトル: GitHub - anat0m1a/liblzraven: Apple の LZRAVEN コーデックの AI 作成、適合性テスト済み C11 実装 · GitHub
説明: AI によって作成され、適合性テスト済みの Apple の LZRAVEN コーデックの C11 実装 - anat0m1a/liblzraven

記事本文:
GitHub - anat0m1a/liblzraven: Apple の LZRAVEN コーデックの AI によって作成され、適合性テスト済みの C11 実装 · GitHub
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
anat0m1a
/
リブルズレイヴン
公共
通知
通知設定を変更するにはサインインする必要があります

イングス
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 Commit 1 Commit bindings/ python bindings/ python docs/notes docs/notes fuzz fuzz include include src src testing testing tools tools .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md SPEC.md SPEC.md すべてのファイルの表示 リポジトリ ファイルのナビゲーション
Apple の LZRAVEN コーデック ( COMPRESSION_LZRAVEN 、アルゴリズム列挙値 0xD05 )、
C11 で再実装 — あらゆるプラットフォームで OS 27 OTA ペイロードをデコードします。
リポジトリの人間の所有者からの簡単な序文:
（はい、これはクロードではなく私が話しています）
これのすべての行は（非常に明確に）クロードによって書かれました。正直に言うと、私は
これを私自身の個人的なツールで動作させる必要がありましたが、これが最も早い方法でした
それを実現するために。私はこのプロジェクトを維持するつもりはまったくありませんが、
しかし、それを必要とする他の人たちと共有する価値があると感じました。
ライセンスに従って、このコードで好きなことをしてください:) 人々がそれが役立つことを願っています
Apple OTA リリース (またはこの新しい形式で圧縮するために選択したもの) を処理するため
選択したプラットフォームが *OS ではない場合。
このため、テストをこの Readme のかなり中心的な部分にしたいと考えました。
これは 100% AI によって作成されていますが、生成されたコードをテストするためにある程度の「努力」があったことは明らかです。
クロードもここでテストを作成しましたが、テストは Apple の libcompression に基づいて採点されました。
それ自体はエミュレーションの下で実行されます。クロードがここで何かをバイト正確であると言及するときは常に、
これは、実装の出力と libcompression の比較を指します。
Apple の圧縮コーデックである LZRAVEN のクリーンルーム C 実装が導入されました
OS 27 サイクル ( COMPRESSION_LZRAVEN )。
Apple はソースもフォーマット仕様も公開していません -

まさに散文
説明。私の知る限り、他にオープンな実装はありません。
現存する既存の「lzraven サポート」は、Apple の非公開に対する cgo の呼び出しです。
libcompression 、および macOS のみです。
このライブラリは、LZRAVEN ストリームをどこでもデコードします。 C11、依存関係なし、動的なし
割り当て。形式は SPEC.md に文書化されています。
OS 27 OTA ペイロードが pbzx コンテナー (xz) から pbzm に移動され、圧縮されました
LZRAVENと。調査対象の OTA のうち、OS 27 境界ではカットはきれいです。
OS 27 では 10 個中 10 個が pbzm を使用し、OS 26 では 7 個中 7 個が pbzx を使用します。
したがって、Apple 以外のプラットフォームで Apple OTA ペイロードを解凍するものはすべて停止します。
OS 27 では libcompression が存在しないため、OS 27 では動作します。 watchOS は、
深刻なケース: OTA のみで出荷されるため、頼れる IPSW はありません。
LZRAVEN の唯一のリファレンス実装は Apple プラットフォーム上でのみ実行されます。
Linux 上で開発されました。それを解決することが第一であり、他のすべてはそれにかかっています。
オラクル。 tools/oracle.py は Apple の libcompression.dylib を
Unicorn ARM64 エミュレータとその実際のエンコーダとデコーダを直接呼び出します - グラウンド
Linux では真実ですが、Mac ではありません。そこで、データを生成することで仮説を立てました。
逆アセンブリを推論するのではなく、Apple 独自のコード。
OS 27 サイクルの Apple プラットフォームでは、tools/native_oracle.py は同じです
libcompression 自体によってサポートされるインターフェースであり、約 1,000 倍高速です。
両方の機能により、エミュレータがエミュレートするものに対してチェック可能になります。
異なる OS ビルド: すべてのベクターはプレーンテキストでコミットされます — 時点で 30
その実行時 — ネイティブ エンコーダでバイト同一に再エンコードします。
スクラッチ サイズが正確に一致し、14 の派生定数テーブルがすべて表示されます。
macOS イメージ内のバイトごと。エミュレータを介してフォーマットが反転されました。
そしてエミュレータは正しかった。
差別化

l テスト ( tools/difftest.py ) は両方のデコーダを同じ上で実行します
入力してバイトごとに比較します。結果:
20 個の彫刻された実際の OTA ペイロードはすべて、Apple の出力とバイト同一にデコードされます。
2 つのデコーダは同じであるため、比較は分類され、等しいと主張されません。
意図的に同等ではありません。私たちは Apple よりも多くのことを検証しています。 「どこで拒否するか
「Apple は受け入れます」は、実際のバグである 2 つのケースとは別にカウントされます (「私たちは
Apple が拒否する場合は受け入れる」、「両方とも異なるバイトを受け入れる」)。
ネイティブ デコーダに対するテストで、最初のクラスが実際にどこに来るのかが判明
から。 Apple のデコーダは、8 つの rANS 状態が次の時点で排出されたかどうかをチェックしません。
ブロックの終わり。それらの下半分をコード化されていない尾部として書き出します。
関係なく成功を返します。私たちはそれを必要とするので、エントロピーコーダーがそれを行わなかったブロック
end は一貫して Apple 用にデコードされ、当社によって拒否されます。表示する 2 つのストリーム
これは testing/vectors/apple-accepts/ に保存されます。不変式はすべてのブロックに適用されます
Apple 独自のエンコーダはこれまで私たちのために作成してきました。これは Apple のエンコーダのプロパティであり、
彼らのデコーダーが強制する規則ではありません。
毛羽立ち。 AFL++: デコーダで 3 億 7,700 万回、BCJ で 5,000 万回の実行
フィルター — クラッシュもハングもゼロ。 libFuzzer: 3 つでさらに約 3,500 万件
ハーネス。 ASan + UBSan 全体的にきれいです。カバレッジ: rANS コーダーの 100%、
モデルとフィルター。記号文法の97%。
エンコーダには、独自の敵対的層があります: ラウンドトリップ ハーネス (
ファザーが選択したレベルとウィンドウでのデコーダーを介したエンコーダーの出力、
バイト比較 — チェックサムレス形式では許容できないエラーは、まさに
中止するもの）と、ファインダーの動作をアサートする直接マッチファインダーハーネス
どのポジションでも契約します。時計に合わせてではなく、カバレッジのプラトーに向けて実行します。
往復653,226件、マッチファインダー実行数210万件、

所見ゼロ、
往復コーパスは、新機能が枯渇するまでに 8 倍に増加しました。まだ
コードが新しく、デコーダの累積走行距離には遠く及ばない —
fuzz/README.md にはその集計が正直に記録されています。
突然変異テスト — スイートが何も検出していないことが証明されるため
何もない。デコーダのコピーには 5 つのバグが意図的に注入されました。
５匹全員捕まえた。それが、「3 億 7,700 万回の実行、クラッシュゼロ」の価値がある理由です。
と言う。
それが見つけたもの。差動層は構造的にファジングする本当のバグを発見しました
できませんでした: このデコーダが受け入れた 2 つのストリームが Apple によって拒否されました。アップルの小切手
すべての一致の前に pos + length <= code_end がコピーされ、完全に拒否されます。これ
代わりに、コードがコメントに書き込まれて決してクランプされなかったという仮定に基づいて、クランプされていました。
テストされました。両方の再現装置が拒否フィクスチャとして維持されるように修正されました。ファジングラン
3 億 7,700 万回の実行では、メモリ エラーが発生しなかったため、緑色で表示されます —
間違った行為だけ。
デフォルトのゲート。 make テストは自己完結型であり、クレームを強制します
それらを信頼するのではなく、上記のとおりです。コミットされたすべてのベクトルは、
C ライブラリ (すべて 33、マルチブロック ストリームを含む、最小数以下)
したがって、欠落しているコーパスは空虚に通過する代わりに失敗します)。 4つすべて
C デコーダによって拒否された、予想される正確なステータスを持つ必須拒否フィクスチャ、
そして再び独立した Python リファレンス デコーダによって行われます。決定論的テスト
rANS フラッシュ境界と範囲外の拒否を固定する —
コア安全性チェック - 変異検証済みの感度。そしてコンテナと
フレームツールスイート。 C テストは、出力バッファを正確に割り当てます。
宛先の容量なので、make asan では 1 バイト オーバーランはレポートであり、レポートではありません。
パス。デフォルトでは diff ゲートを --max-stricter 0 に設定し、次の場合はハード失敗します。
神託は評決を下すことができない

、空虚な合意を報告するのではなく。
常駐の警告。これらはすべて、1 つの Apple ビルドに対して検証されます。
リポジトリに sha256 が記録されている iOS 27.0 beta 4 dylib、一度クロスチェック
macOS 27.0 イメージ (SPEC.md §14) に対して。ポイントリリースにより変更される可能性があります
形式であり、ペイロードが失敗するまでは何もわかりません。これは、ペイロードに固有のものです。
仕様なしでの再実装。来歴の規律がそれを制限するものです。
すべての主張は特定の人工物に関連付けられているため、将来の相違は次のとおりです。
神秘的ではなく、原因によるものです。
make test-paths # すべての SIMD パス、33 ベクトル
python3 testing/validate_symbols.py # 独立した Python デコーダ
python3 tools/difftest.py --full --fresh --max-stricter 0 # Apple のデコーダに対して
python3 tools/enc_roundtrip.py --big # 私たちのエンコード、Apple デコード
python3 tools/enc_roundtrip.py --far --far6 # 距離クラス 5 および 6
python3 tools/enc_real.py # 実際の OTA 平文でも同じ
asan-paths # サニタイザーを作成、すべてのパス
make -C fuzz run-decode # libFuzzer、デコーダー
make -C fuzz run-roundtrip # エンコーダのラウンドトリップファザー
--far6 は 271 MB のプレーンテキストを構築し、約 3 GB の RAM が必要です。大したことは何もない
コミットされると、ハーネスがそれを生成します。
ベクトル テストは自己完結型です。ディファレンシャル ツールには Apple のものが必要です
コーデック: Linux では、 work/ 内の libcompression.dylib のコピーを意味します。
自分で用意する必要があります — これは Apple の独自コードであり、配布されていません
ここです。 macOS 27 以降では、OS に同梱されているため、何も必要ありません。
difftest.py はそのバックエンドを自動的に選択します。
ベンチマークにはコーパスが必要です。 python3 tools/make_corpus.py はこれからビルドします
Apple のエンコーダを使用したリポジトリ独自のコンテンツ、および両方のベンチマーク ツールが受け入れます
--streams 経由で実行します。
make # build/liblzraven.a
共有する # build/liblzraven.so
#include <lzraven.h>
size_t n = lzra

ven_decode_buffer ( dst 、 dst_capacity 、 src 、 src_size );
if ( n == 0 ) { /* 拒否されました — lzraven_decode_buffer_ex() が理由を示します */ }
エンコードには呼び出し元が提供するスクラッチが必要です。デフォルトは次のとおりです。
lzraven_encode_buffer は以下を使用します。
void *Scratch = malloc ( lzraven_encode_scratch_size ());
size_t n = lzraven_encode_buffer ( dst 、 lzraven_encode_bound ( len )、 src 、 len 、
傷）;
lzraven_encode_params p ; /* または調整します */
lzraven_encode_params_default ( & p );
p 。レベル = 8 ; /* 0..9、6 = デフォルト */
p 。ウィンドウ = 32u << 20 ; /* 最大一致距離 */
スクラッチ = malloc ( lzraven_encode_scratch_size_params ( & p ));
n = lzraven_encode_buffer_params ( dst , cap , src , len ,Scratch , & p );
スクラッチ サイズは、呼び出し元がエンコードする前のパラメータから導出できます。
それを割り当てます。それはマッチファインダーによって支配され、したがってウィンドウによって支配されます。
Python からも同様です: bindings/python/ は stdlib 専用の ctypes バインディングです —
1 つのファイル、ビルドステップなし、上記の共有オブジェクトをロードします ( pip install bindings/python 、または copy lzraven.py )。デコードは解凍されたサイズを取得します
カラスの流れがそれを運ばないからです。 pbzmコンテナはそうします
(SPEC.md §10)。
ズレイヴンをインポート
プレーン = lzraven 。デコード ( blob 、 unc_size )
コード化された = lzraven 。エンコード (プレーン、レベル

[切り捨てられた]

## Original Extract

AI-written, conformance-tested C11 implementation of Apple's LZRAVEN codec - anat0m1a/liblzraven

GitHub - anat0m1a/liblzraven: AI-written, conformance-tested C11 implementation of Apple's LZRAVEN codec · GitHub
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
anat0m1a
/
liblzraven
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit bindings/ python bindings/ python docs/ notes docs/ notes fuzz fuzz include include src src tests tests tools tools .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md SPEC.md SPEC.md View all files Repository files navigation
Apple's LZRAVEN codec ( COMPRESSION_LZRAVEN , algorithm enum 0xD05 ),
reimplemented in C11 — decode OS 27 OTA payloads on any platform.
A quick foreword from the human owner of the repo:
(yes this is me talking not claude)
every line of this was (quite clearly) written by Claude. I'll be completely honest, I
needed to get this working for my own personal tooling, and this was the quickest path
to making that happen. I do not intend to maintain this project in any capacity,
but it felt worth sharing with others who might need it.
As per the license, do whatever you want with this code :) I hope that people find it useful
for dealing with Apple OTA releases (or whatever they choose to compress with this new format)
if their chosen platform is not *OS.
It's for this reason that I wanted to make testing a pretty core part of this readme, to make it
clear that although this is 100% AI-authored, there was some "effort" to test the generated code.
Claude did also write the tests here, but the tests were graded on Apple's libcompression
itself, running under emulation. Whenever claude refers to something as being byte-exact here,
it's referring to our implementation's output vs libcompression.
A clean-room C implementation of LZRAVEN , Apple's compression codec introduced
in the OS 27 cycle ( COMPRESSION_LZRAVEN ).
Apple has published no source and no format specification — only a prose
description. As far as I can tell there is no other open implementation: the one
existing "lzraven support" in the wild is a cgo call into Apple's closed
libcompression , and so is macOS-only.
This library decodes LZRAVEN streams anywhere. C11, no dependencies, no dynamic
allocation. The format is documented in SPEC.md .
OS 27 OTA payloads moved from the pbzx container (xz) to pbzm , compressed
with LZRAVEN . The cut is clean at the OS 27 boundary — of the OTAs surveyed,
10 of 10 on OS 27 use pbzm , 7 of 7 on OS 26 use pbzx .
Anything that unpacks Apple OTA payloads on a non-Apple platform therefore stops
working at OS 27, because libcompression does not exist there. watchOS is the
acute case: it ships OTA-only, so there is no IPSW to fall back on.
LZRAVEN's only reference implementation runs solely on Apple platforms, and this
was developed on Linux. Solving that came first, and everything else rests on it.
The oracle. tools/oracle.py loads Apple's libcompression.dylib into a
Unicorn ARM64 emulator and calls its real encoder and decoder directly — ground
truth on Linux, no Mac. So hypotheses were settled by generating data with
Apple's own code rather than by reasoning about disassembly.
On an Apple platform from the OS 27 cycle, tools/native_oracle.py is the same
interface backed by libcompression itself, about a thousand times faster.
Having both made the emulator checkable against the thing it emulates, on a
different OS build: every vector then committed with a plaintext — 30 at the
time of that run — re-encodes byte-identically under the native encoder, both
scratch sizes agree exactly, and all 14 derived constant tables appear
byte-for-byte in the macOS image. The format was reversed through the emulator,
and the emulator was right.
Differential testing ( tools/difftest.py ) runs both decoders over the same
inputs and compares byte-for-byte. Results:
All 20 carved real OTA payloads decode byte-identically to Apple's output.
The comparison is classified , not asserted equal, because the two decoders are
intentionally not equivalent — we validate more than Apple does. "We reject where
Apple accepts" is counted separately from the two cases that are real bugs ("we
accept where Apple rejects", "both accept, different bytes").
Testing against the native decoder found where that first class actually comes
from. Apple's decoder does not check that the eight rANS states drained at
the end of a block; it writes their low halves out as the uncoded tail and
returns success regardless. We require it, so a block whose entropy coder did not
end consistently decodes for Apple and is rejected by us. Two streams that show
it are kept in tests/vectors/apple-accepts/ . The invariant holds on every block
Apple's own encoder has ever produced for us — it is a property of their encoder,
not a rule their decoder enforces.
Fuzzing. AFL++: 377 million executions on the decoder, 50 million on the BCJ
filters — zero crashes, zero hangs. libFuzzer: ~35 million more across three
harnesses. ASan + UBSan clean throughout. Coverage: 100% of the rANS coder, the
model and the filters; 97% of the symbol grammar.
The encoder has its own adversarial tier: a round-trip harness (our
encoder's output through our decoder at fuzzer-chosen level and window,
byte-compared — the failure a checksum-less format cannot afford is exactly the
one it aborts on) and a direct match-finder harness asserting the finder's
contract at every position. Run to coverage plateau rather than to a clock:
653,226 round-trip and 2.1 million match-finder executions, zero findings, the
round-trip corpus growing eight-fold before new features dried up. Still the
younger code and nowhere near the decoder's cumulative mileage —
fuzz/README.md keeps that tally honestly.
Mutation testing — because a suite that has never caught anything proves
nothing. Five bugs were deliberately injected into copies of the decoder:
All five caught. That is what makes "377 million executions, zero crashes" worth
saying.
What it found. The differential tier caught a real bug fuzzing structurally
could not: two streams this decoder accepted that Apple rejects. Apple checks
pos + length <= coded_end before every match copy and rejects outright; this
code had clamped instead, on an assumption written into a comment and never
tested. Fixed, with both reproducers kept as rejection fixtures. Fuzzing ran
green through it for 377 million executions because it produced no memory error —
only wrong behaviour.
The default gate. make test is self-contained and enforces the claims
above rather than trusting them: every committed vector byte-exact through the
C library (all 33, the multi-block stream included, behind a minimum-count
floor so a missing corpus fails instead of vacuously passing); all four
must-reject fixtures rejected by the C decoder with the exact expected status,
and again by the independent Python reference decoder; deterministic tests
pinning the rANS flush boundary and the out-of-range-distance rejection — the
core safety check — with mutation-verified sensitivity; and the container and
framing tool suites. The C tests allocate output buffers at exactly the
destination capacity, so under make asan a one-byte overrun is a report, not
a pass. make diff gates at --max-stricter 0 by default and fails hard if
the oracle cannot produce a verdict, rather than reporting vacuous agreement.
The standing caveat. All of this validates against exactly one Apple build:
the iOS 27.0 beta 4 dylib whose sha256 the repo records, cross-checked once
against a macOS 27.0 image ( SPEC.md §14). A point release could change the
format, and nothing here would know until a payload failed — inherent to a
reimplementation with no spec. The provenance discipline is what bounds it:
every claim is tied to a specific artefact, so a future divergence will be
attributable, not mysterious.
make test-paths # every SIMD path, 33 vectors
python3 tests/validate_symbols.py # independent Python decoder
python3 tools/difftest.py --full --fresh --max-stricter 0 # against Apple's decoder
python3 tools/enc_roundtrip.py --big # ours encodes, Apple decodes
python3 tools/enc_roundtrip.py --far --far6 # distance classes 5 and 6
python3 tools/enc_real.py # the same on real OTA plaintext
make asan-paths # sanitizers, every path
make -C fuzz run-decode # libFuzzer, decoder
make -C fuzz run-roundtrip # encoder round-trip fuzzer
--far6 builds a 271 MB plaintext and wants about 3 GB of RAM; nothing large is
committed, the harness generates it.
The vector tests are self-contained. The differential tooling needs Apple's
codec: on Linux that means a copy of libcompression.dylib in work/ , which you
must supply yourself — it is Apple's proprietary code and is not distributed
here. On macOS 27 or later nothing is needed, since the OS ships it;
difftest.py picks that backend automatically.
Benchmarking needs a corpus. python3 tools/make_corpus.py builds one from this
repository's own content using Apple's encoder, and both benchmark tools accept
it via --streams .
make # build/liblzraven.a
make shared # build/liblzraven.so
#include <lzraven.h>
size_t n = lzraven_decode_buffer ( dst , dst_capacity , src , src_size );
if ( n == 0 ) { /* rejected — lzraven_decode_buffer_ex() gives the reason */ }
Encoding takes caller-provided scratch, and the defaults are what
lzraven_encode_buffer uses:
void * scratch = malloc ( lzraven_encode_scratch_size ());
size_t n = lzraven_encode_buffer ( dst , lzraven_encode_bound ( len ), src , len ,
scratch );
lzraven_encode_params p ; /* or tune it */
lzraven_encode_params_default ( & p );
p . level = 8 ; /* 0..9, 6 = default */
p . window = 32u << 20 ; /* max match distance */
scratch = malloc ( lzraven_encode_scratch_size_params ( & p ));
n = lzraven_encode_buffer_params ( dst , cap , src , len , scratch , & p );
Scratch size is derivable from the parameters before encoding, since the caller
allocates it; it is dominated by the match finder, hence by the window.
The same from Python: bindings/python/ is a stdlib-only ctypes binding —
one file, no build step, loading the shared object above ( pip install bindings/python , or copy lzraven.py ). decode takes the decompressed size
because a raven stream does not carry it; the pbzm container does
( SPEC.md §10).
import lzraven
plain = lzraven . decode ( blob , unc_size )
coded = lzraven . encode ( plain , level

[truncated]
