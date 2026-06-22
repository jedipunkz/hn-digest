---
source: "https://github.com/kromych/badc"
hn_url: "https://news.ycombinator.com/item?id=48625919"
title: "Badc – optimizing cross-platform C compiler built with Claude"
article_title: "GitHub - kromych/badc: Former bad C compiler · GitHub"
author: "kromych"
captured_at: "2026-06-22T05:47:06Z"
capture_tool: "hn-digest"
hn_id: 48625919
score: 1
comments: 2
posted_at: "2026-06-22T05:05:42Z"
tags:
  - hacker-news
  - translated
---

# Badc – optimizing cross-platform C compiler built with Claude

- HN: [48625919](https://news.ycombinator.com/item?id=48625919)
- Source: [github.com](https://github.com/kromych/badc)
- Score: 1
- Comments: 2
- Posted: 2026-06-22T05:05:42Z

## Translation

タイトル: Badc – Claude で構築されたクロスプラットフォーム C コンパイラの最適化
記事タイトル: GitHub - kromych/badc: かつての悪い C コンパイラ · GitHub
説明: 以前は不良の C コンパイラでした。 GitHub でアカウントを作成して、kromych/badc の開発に貢献してください。

記事本文:
GitHub - kromych/badc: 以前の悪い C コンパイラ · GitHub
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
クロミッチ
/
悪い
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
master ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く F

年長者とファイル
2,197 コミット 2,197 コミット .github/ workflows .github/ workflows デモ デモの例 例 headers/ include headers/ include lib lib scripts scripts src src テスト テスト ツール ツール .gitignore .gitignore CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンスREADME.md README.md build.rs build.rs hello.c hello.c std-conformance.md std-conformance.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
badc はかなり小規模なクロスプラットフォーム最適化コンパイラ (ライブラリとしてのコンパイラでもあります)
C言語の。
C コンパイラーが何をどのように出力するかを迅速に調整する必要性から、必要に迫られて登場しました。
そして、それが軽快な実用的なツールになることができるのは魅力的でした
ニッチなハックではなく、日常的に使用するために。コーディングに対する最新のアプローチでは、
コンパイラの構築は、思ったより簡単でした :)
現在、badc は C99、C11 標準、および一部の標準の大部分を実装しています。
後の標準からの一般的なイディオムといくつかの拡張。それはすべて
サポートされている 5 つのターゲットすべてで Python 3.14 をビルドしてテストするには十分です (そして
さらに多くのデモが含まれていますので、読んでください!)。
badc の小さなフットプリントと埋め込みヘッダー (オーバーライドまたは --install が可能)
調整または検査のためのパスへのパス) の 1 つの実行可能エクスペリエンスを提供します。
携帯用ツール。適度なサイズのコンパイラのコードベースは、小規模なコードベースとして使用できます。
独立したツールチェーンとして使用することも、プロジェクトに必要な機能を提供するライブラリとして使用することもできます。
C コードをビルドしたり、それを実行したりする機能 (ライブラリとして使用する場合のデフォルト)。
楽しい拡張機能は、badc がヘッダーを自動的に追加できることです。
標準ライブラリの場合は、裸の hello.c を使用します。
int main () {
put ( "こんにちは" );
0 を返します。
}
作品:
情報: 未宣言の `puts` を自動インクルードする <stdio.h>
情報: ターゲット `macos-aarch に対してファイル hello を書き込みました

64`
badc はデバッグ情報を生成できるため、生成されるバイナリは
デバッグしたり、パフォーマンスをプロファイリングしたりできます ( -g を使用)。
badc は、-O を指定すると最適化され、より高速なコードを生成できます。
特に ARM64 では、clang -O0 よりも優れています。コード生成のアイデアを得るには
品質については、アセンブリを含む ./tests/snapshot を見てください。
テスト フィクスチャの SSA スナップショット。最適化されたバイナリは最新のあらゆる環境で動作します。
ARM64 プロセッサ、および Intel Haswell および AMD Zen より古い x86_64 プロセッサ上
(2013 年頃、オプティマイザーは FMA3 命令を使用しました)。
badc は、位置に依存しないコードと実際のネイティブ バイナリ (macOS Mach-O、
Linux ELF または Windows PE32+)、任意のホストからの 5 つのターゲットのいずれかで:
Windows ({ ARM64 、 x86_64 } x { コンソール 、 GUI 、 NT 、ドライバー })。
個別の翻訳単位 (常に ELF に翻訳される) もサポートしており、小さな
リンカー (したがって、緩和や LTO はありません)。 badc は仮定が邪魔にならないよう努めます
ランタイム ライブラリ上にあり、必要に応じて --freestand も利用可能です。 EFI
もサポートされています。
badc はプロセス内でマシンコードに JIT コンパイルすることもできるため、バイナリは書き込まれません。
ディスクに。最後に、#! として使用されていることを認識します。 C ソースコードは次のようになります
(高速) スクリプト。
デモの下にはさまざまなデモがあります:
小さなものはいくつかあります ( thread.c 、 coro_pool.c 、 hello_server.c )。
maze.c - 迷路ビルダーおよびソルバー、
gui_hello - macOS、Linux、Windows 用の GUI デモ、
wdm_driver 、 nt_hello 、 nt_loader - Windows ネイティブ (NT) 実行可能ファイル、Windows ドライバーの例
sqlite3 - 最も有名な組み込みデータベース、
miniz - 圧縮、CRC32、整数、ビットいじり、
Kissfft - 浮動小数点、高速フーリエ変換、
bzip2 - 圧縮、整数、ビットいじり、
stb - 多くの素晴らしい機能 (数学) を備えたヘッダーのみの C ライブラリ
いいえ

いくつか例を挙げると、生成、サウンド、JPEG、PNG、BMP、PSD のサポート)。
これはコンパイラ全体に大きなストレスを与えます。
tinycc - クールで小さな C ツールチェーン
TweetNaCl 、 Monocypher 、 BearSSL - 暗号化
Lua - 埋め込み可能なスクリプト言語
Quickjs - JavaScript インタープリター
これらに加えて、ホーナー方式を実装した楽しいテスト フィクスチャ、RK4、
8-クイーンなど。
最後に、IR (中間表現) を実行するオプションがあります。
ポインタのアクセスと境界を追跡して、メモリの問題を検出します。
プロジェクトが始まったばかりで、その名前が定着していたとき、badc はひどいものでした。
このドキュメントには、コンパイラ構築に関する専門用語が随所に含まれています。できます
安全にスキップして、すぐに使用法のセクションにジャンプしてください。
真のコンパイラヘッドの場合は、それぞれを出力する --dump-ssa オプションがあります。
関数の SSA IR に加えて、レジスタ アロケータの値ごとの stderr への配置が前に行われます。
下げること。
これは、Robert Swierczek の 4 つの関数を備えた非常に小さな C コンパイラの Rust 移植として始まりました。
c4、そこから成長しました。それで十分な乖離があった
オリジナルから方言を c5 と呼びます。ソースツリーにそのふざけた名前が付けられているため、
それを c5 モジュールと C5Error タイプとして詳しく説明します。
由緒ある 4 機能の c4.c コンパイラは、テスト フィクスチャおよびセルフホストとして出荷されます。
badc -O -o c4 testing/fixtures/c/c4.c # c4 をネイティブ バイナリにコンパイルします
./c4 hello.c # これにより hello.c が実行されます
そして、次のようなものを使用すると、さらに楽しみを増やすことができます
badc -O --jit テスト/フィクスチャ/c/c4.c テスト/フィクスチャ/c/c4.c テスト/フィクスチャ/c/c4.c テスト/フィクスチャ/c/c4.c
それをクアドロネストして実行するには:)
開発中に、badc コンパイラがスタックから「スパイラル」で出てきました。
IR の実行と 3 オペランド IR および SSA IR へのフロントエンドの進化、および最適化
バックエンド。
SSA 中間表現とグラフ色付けを通じて低下します。
登録する

アロケータですが、絶妙な最適化を目指していないため、
Clang、gcc、msvc などの titan ツールチェーンを実行します。結局のところ、スリムを維持するには、
マルチギガバイトのコンパイラスイートの能力を超える可能性は低いです。
マシンから最後の一滴まで搾り取っても問題ありません。
インストール方法と最初の手順
使用しているバージョンに一致するバイナリ リリース パッケージの 1 つをダウンロードできます。
ハードウェアとOS。中には小さなバイナリが 1 つあります。
badc の使用を開始するために必要なのはこれだけです。
Rust がインストールされている場合は、リポジトリを複製し、次のようにインストールします。
カーゴインストール --path 。 -- 完全な機能
または単に
カーゴインストール badc -- 完全な機能
ソースコードから構築することに興味がない場合。
--features full はコマンドライン コンパイラに必要です。
クレートのデフォルトの機能セットはホスト アーキテクチャ JIT ライブラリのみです
(したがって、カーゴはスリムな依存関係で badc プルを追加します)、および badc バイナリ
さらにネイティブ オブジェクト ライターとクロストランスレーション ユニットが必要です
リンカー。完全な機能により有効になります。
これで、badc が PATH 上で利用できるようになりました。
badc --jit hello.c # ネイティブ コードをインプロセスで実行します
こんにちは123
または
badc -O hello.c # ネイティブに最適化されたバイナリを生成します
./hello # 前の行で生成されたもの
こんにちは123
簡単なデバッグ セッションは次のとおりです。
badc -g hello.c # デバッグ情報を使用してビルドします
情報: ターゲット macos-aarch64 のファイル hello を書き込みました
次に、デバッガー ( lldb 、 gdb 、 rr ) で実行し、ブレークポイントを設定し、ローカル変数をチェックアウトします。
lldb ./こんにちは
(lldb) ターゲットは「./hello」を作成します
現在の実行可能ファイルは「/Users/krom/src/compilers/badc/hello」(arm64) に設定されています。
(lldb) b メイン
ブレークポイント 1: = hello`main + 16 at hello.c:5、アドレス = 0x00000001000006fc
(lldb) l
注: ソースはありません
(lldb) 実行
プロセス 19800 が起動されました: '/Users/krom/src/compilers/badc/hello' (arm64)
プロセス 19800 が停止しました
* スレッド #1、ク

eue = 'com.apple.main-thread'、停止理由 = ブレークポイント 1.1
フレーム #0: 0x00000001000006fc hello`main at hello.c:5
2 #include <stdlib.h>
3
4 int main() {
-> 5 int a = 123;
6 printf("こんにちは %d\n", a);
7 0 を返します。
8 }
ターゲット 0: (こんにちは) 停止しました。
(lldb)n
プロセス 19800 が停止しました
* スレッド #1、キュー = 'com.apple.main-thread'、停止理由 = ステップ オーバー
フレーム #0: 0x0000000100000704 hello`main at hello.c:6
3
4 int main() {
5 int a = 123;
-> 6 printf("こんにちは %d\n", a);
7 0 を返します。
8 }
ターゲット 0: (こんにちは) 停止しました。
(lldb)v
(int) a = 123
最初の非フラグ引数はソース ファイルです。デフォルトではbadc
これを、次の明らかなパスにあるネイティブ バイナリに下げます。
ソース ( hello.c -> POSIX ターゲットでは hello、hello.exe では
Windows ターゲット); pass -o <path> を使用して、別のパスを選択します。
フラグ ( --target=<spec> 、 --optimize / -O 、 --dump-ssa 、
--list-symbols 、 -H / --show-includes 、および VM のみ
--track-pointers / --trace ）は、
ソース。 -D NAME[=VALUE] 、 -U NAME 、 -I path 、および -include FILE は、 gcc/clang での場合と同じように機能します。ソース駆動型
ビルド フラグは #pragma に乗ります -- 「ヘッダーとバインディング」を参照してください。
以下。
.c ファイルはシバンで始まる場合があります。 PATH に badc を指定すると、
chmod +x script.c はファイルを直接実行可能にします -- で
この場合、shebang 行はモードを選択します (VM の場合は #!/usr/bin/env badc --interp、ネイティブ コンパイルの裸の形式)。
5 つのターゲットがサポートされており、任意のホストからいずれかのターゲットにクロスコンパイルできます。
1 回の badc 呼び出しで、.c ソース ファイル、.o を混在させることができます。
オブジェクト ファイルと .a アーカイブ:
badc -c foo.c bar.c # foo.o + bar.o を発行します (ネイティブ ELF64 ET_REL、ターゲットは固定されています)
badc -o app foo.o bar.o # それらを最終バイナリにリンクします
badc --ar -o libfoo.a foo.c bar.c # SysV ar(5) アーカイブにバンドルします
badc -o app main.c -L。 -l foo # libfoo.a に対するリンク、gcc スタイル

badc は独自のリンカーを同梱しています -- ld / lld / はありません
link.exe の依存関係。オブジェクト ファイルは標準の ELF64 ET_REL
リロケータブル: ネイティブ マシン コードの .text セクション、
.data / .bss（静的ストレージの場合）、.symtab / .strtab
名前テーブル用、および再配置を伝える .rela.text
各ユニットの最終位置が判明すると、リンカーが適用されます。
ターゲットは -c 時に固定され、オブジェクトも固定されます。
ld / lld でリンク可能。アーカイブは SysV スタイルの ar(5) です
シンボルインデックス。完全な貨物機能により、全体のゲートが確保されます。
パイプライン;必要のないライブラリ利用者
マルチ TU アーティファクトは、次の方法でオプトアウトできます。
default-features = false、features = ["std"] を維持するには、
設置面積がスリム。
ストレージクラスのリンケージは C99 6.2.2 に従います: static at file
スコープが内部であり、ベアまたは外部宣言が外部である、
および外部 Tx;定義宣言がない場合は、
リンカが要求から満たそうとする未解決の外部
残りのオブジェクトまたはアーカイブ メンバー。
方言が何を解析 + 低めるのか、そしてどこで解析されるのかの概要
C99 から分岐し、 std-conformance.md に存在します。ショート
バージョン: c5 は、言語の大部分と、後の標準のいくつかの機能をカバーしています。
このドキュメントには、拒否されたイディオム、相違する動作、および c5 のみの拡張機能が列挙されています。
(

[切り捨てられた]

## Original Extract

Former bad C compiler. Contribute to kromych/badc development by creating an account on GitHub.

GitHub - kromych/badc: Former bad C compiler · GitHub
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
kromych
/
badc
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
2,197 Commits 2,197 Commits .github/ workflows .github/ workflows demos demos examples examples headers/ include headers/ include lib lib scripts scripts src src tests tests tools tools .gitignore .gitignore CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md build.rs build.rs hello.c hello.c std-conformance.md std-conformance.md View all files Repository files navigation
badc is a rather small cross-platform optimizing compiler (also a compiler-as-library)
of the C language.
It had appeared out of necessity to quickly tweak how and what a C compiler emits.
Then it was captivating making it being able to become a nimble practical tool
for everyday use rather than a niche hack. Modern approaches to coding would make
building a compiler easier than that had been before I thought :)
Now badc implements a very large portion of the C99, C11 standards and some
popular idioms from the later standards as well as few extensions. All of that is
enough to build and test Python 3.14 on all of the five supported targets (and
there are more demos included, read on!).
badc 's small footprint and embedded headers (which you can override or --install
to some path for tweaking or inspecting) give a one-executable experience of the
portable tools. The compiler's codebase of moderate size can be used as a small
self-sufficient toolchain or can be used as a library giving your project the
ability to build C code or just run it (the default when using as a library).
A fun extension is that badc can automatically add the header(s)
for the standard library so the bare hello.c with
int main () {
puts ( "Hello" );
return 0 ;
}
works:
info: auto-including <stdio.h> for undeclared `puts`
info: wrote file hello for target `macos-aarch64`
badc is able to produce the debug information so that the binaries it generates
can be debugged and/or their performance can be profiled (use -g ).
badc optimizes when you specify -O and can produce code that's faster
than clang -O0 , especially on ARM64. To get an idea of the codegen
quality, take a look at ./tests/snapshots with assembly and
SSA snapshots of the test fixtures. The optimized binaries will run on any modern
ARM64 processor, and on x86_64 processors not older than Intel Haswell and AMD Zen
(circa 2013, the optimizer uses FMA3 instructions).
badc emits position-independent code and the real native binaries (macOS Mach-O,
Linux ELF, or Windows PE32+), on any of five targets, from any host:
Windows ({ ARM64 , x86_64 } x { console , GUI , NT , driver }).
It supports also separate translation units (always translated to ELF) and has a small
linker (so no relaxations or LTO). badc tries hard not to get in the way with assumptions
on the runtime library, and --freestanding as available should you need that. EFI
is supported as well.
badc can also JIT-compile into the machine code in-process so no binary is written
to the disk. Finally, it recognizes being used as #! so that C source code becomes
a (fast) script.
There are various demo's under demos :
Few small-ish ones ( threads.c , coro_pool.c , hello_server.c ),
maze.c - maze builder and solver,
gui_hello - GUI demos for macOS, Linux and Windows,
wdm_driver , nt_hello , nt_loader - examples of the Windows native (NT) executable, Windows driver,
sqlite3 - the most famous embedded database,
miniz - compression, CRC32, integers, bit twiddling,
kissfft - floating points, Fast Fourier Transform,
bzip2 - compression, integers, bit twiddling,
stb - header-only C library with lots of incredible features (math
noise generation, sound, JPEG, PNG, BMP, PSD support to name a few).
It really stresses all of the compiler.
tinycc - a cool and small C toolchain
TweetNaCl , Monocypher , BearSSL - cryptography
Lua - the embeddable scripting language
quickjs - JavaScript interpreter
Besides these, there are some fun test fixtures implementing Horner scheme, RK4,
8-Queens and more.
Finally, there's an option to run the IR (intermediate representation) with
tracking pointer access and bounds to catch memory issues.
badc used to be bad when the projects just started out and the name stuck.
There is some compiler-building jargon in this document here and there. You can
safely skip it, and jump to the usage section right away.
For the true compiler heads there is the --dump-ssa option which prints each
function's SSA IR plus the register allocator's per-value placement to stderr before
lowering.
It started out as a Rust port of Robert Swierczek's teeny-tiny C compiler in 4 functions
c4 and grew from there. There then has been enough divergence
from the original to call the dialect c5 . Due to that facetious naming the source tree
spells that out as the c5 module and C5Error type.
The venerable 4-function c4.c compiler ships as a test fixture and self-hosts:
badc -O -o c4 tests/fixtures/c/c4.c # compile c4 to a native binary
./c4 hello.c # which then runs hello.c
And you can really crank the fun up with something like
badc -O --jit tests/fixtures/c/c4.c tests/fixtures/c/c4.c tests/fixtures/c/c4.c tests/fixtures/c/c4.c
to run it quadro-nested :)
During the development, the badc compiler was "spiraling" out from the stack
IR execution and evolving frontend to the 3-operand IR and SSA IR and the optimizing
backend.
It lowers through an SSA intermediate representation and a graph-coloring
register allocator, but doesn't go for the exquisite optimization passes a
titan toolchain like clang, gcc or msvc run. All told, to stay slim,
it's unlikely to surpass the ability of multi-gigabyte compiler suites to
squeeze the last drop of perf from the machine, and that's fine.
How to install and first steps
You can download one of the binary release packages matching your
hardware and the OS. There is one small binary inside, and that's
all you should need to start using badc .
If you have Rust installed, clone the repo, and install it with
cargo install --path . --features full
or just
cargo install badc --features full
if you're not interested in building from the source code.
The --features full is required for the command-line compiler: the
crate's default feature set is the host-architecture JIT library alone
(so cargo add badc pulls in a slim dependency), and the badc binary
additionally needs the native object writers and the cross-translation-unit
linker, which the full feature enables.
Now badc is available on the PATH.
badc --jit hello.c # runs native code in-process
Hello 123
or
badc -O hello.c # Produces native optimized binary
./hello # produced by the previous line
Hello 123
Here's a quick debugging session:
badc -g hello.c # Build with the debug information
info: wrote file hello for target macos-aarch64
Now run under the debugger ( lldb , gdb , rr ), set breakpoints, check out the local variables:
lldb ./hello
(lldb) target create "./hello"
Current executable set to '/Users/krom/src/compilers/badc/hello' (arm64).
(lldb) b main
Breakpoint 1: where = hello`main + 16 at hello.c:5, address = 0x00000001000006fc
(lldb) l
note: No source available
(lldb) run
Process 19800 launched: '/Users/krom/src/compilers/badc/hello' (arm64)
Process 19800 stopped
* thread #1, queue = 'com.apple.main-thread', stop reason = breakpoint 1.1
frame #0: 0x00000001000006fc hello`main at hello.c:5
2 #include <stdlib.h>
3
4 int main() {
-> 5 int a = 123;
6 printf("Hello %d\n", a);
7 return 0;
8 }
Target 0: (hello) stopped.
(lldb) n
Process 19800 stopped
* thread #1, queue = 'com.apple.main-thread', stop reason = step over
frame #0: 0x0000000100000704 hello`main at hello.c:6
3
4 int main() {
5 int a = 123;
-> 6 printf("Hello %d\n", a);
7 return 0;
8 }
Target 0: (hello) stopped.
(lldb) v
(int) a = 123
The first non-flag argument is the source file. By default badc
lowers it to a native binary at the obvious path next to the
source ( hello.c -> hello on POSIX targets, hello.exe on
Windows targets); pass -o <path> to choose a different one.
Flags ( --target=<spec> , --optimize / -O , --dump-ssa ,
--list-symbols , -H / --show-includes , plus the VM-only
--track-pointers / --trace ) can appear anywhere before the
source. -D NAME[=VALUE] , -U NAME , -I path , and -include FILE work the same way they do on gcc / clang. Source-driven
build flags ride on #pragma s -- see "Headers and bindings"
below.
A .c file may start with a shebang. With badc on PATH ,
chmod +x script.c makes the file directly executable -- in
which case the shebang line picks the mode ( #!/usr/bin/env badc --interp for the VM, the bare form for native compilation).
Five targets are supported, and you cross-compile from any host to any of them:
A single badc invocation can mix .c source files, .o
object files, and .a archives:
badc -c foo.c bar.c # emits foo.o + bar.o (native ELF64 ET_REL, target pinned)
badc -o app foo.o bar.o # links them into a final binary
badc --ar -o libfoo.a foo.c bar.c # bundles into a SysV ar(5) archive
badc -o app main.c -L. -l foo # link against libfoo.a, gcc-style
badc ships its own linker -- there's no ld / lld /
link.exe dependency. Object files are standard ELF64 ET_REL
relocatables: a .text section of native machine code,
.data / .bss for static storage, .symtab / .strtab
for the name table, and .rela.text carrying the relocations
the linker applies once each unit's final position is known.
The target is pinned at -c time, and the objects are also
linkable by ld / lld . Archives are ar(5) with a SysV-style
symbol index. The full cargo feature gates the entire
pipeline; library consumers that don't need
multi-TU artifacts can opt out via
default-features = false, features = ["std"] to keep the
footprint slim.
Storage-class linkage follows C99 6.2.2: static at file
scope is internal, bare or extern declarations are external,
and extern T x; with no defining declaration becomes an
unresolved external that the linker tries to satisfy from the
remaining objects or archive members.
A summary of what the dialect parses + lowers, and where it
diverges from C99, lives in std-conformance.md . Short
version: c5 covers most of the language and few features of the later standards.
The doc enumerates rejected idioms, divergent behavior, and the c5-only extensions
(

[truncated]
