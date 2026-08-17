---
source: "https://github.com/WhiskeyTuesday/omawrite-win"
hn_url: "https://news.ycombinator.com/item?id=49326581"
title: "I made Claude port omawrite to windows"
article_title: "GitHub - WhiskeyTuesday/omawrite-win: why? · GitHub"
author: "whiskeytuesday"
captured_at: "2026-08-17T05:26:46Z"
capture_tool: "hn-digest"
hn_id: 49326581
score: 1
comments: 1
posted_at: "2026-08-17T04:33:23Z"
tags:
  - hacker-news
  - translated
---

# I made Claude port omawrite to windows

- HN: [49326581](https://news.ycombinator.com/item?id=49326581)
- Source: [github.com](https://github.com/WhiskeyTuesday/omawrite-win)
- Score: 1
- Comments: 1
- Posted: 2026-08-17T04:33:23Z

## Translation

タイトル: クロードポートを Windows に omawrite してみました
記事のタイトル: GitHub - Whisky Tuesday/omawrite-win: なぜ? · GitHub
説明: なぜですか? GitHub でアカウントを作成して、Whiskey Tuesday/omawrite-win の開発に貢献してください。

記事本文:
GitHub - Whisky Tuesday/omawrite-win: なぜですか? · GitHub
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
ウィスキー火曜日
/
オマライトウィン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット スクリーンショット スクリーンショット .gitignore .gitignore Dockerfile.mingw Dockerfile.mingw README-test.md README-test.md README.md README.md build-windows.sh build-windows.shdeploy.pydeploy.py install.cmd install.cmd make-sfx.sh make-sfx.sh

run-wine.sh run-wine.sh set-dark.cmd set-dark.cmd set-light.cmd set-light.cmd windows-port.patch windows-port.patch すべてのファイルを表示 リポジトリ ファイルのナビゲーション
私の名前はオジマンディアス、王の中の王です。
私の作品を見てください、力ある者よ、そして絶望してください！
それ以外には何も残りません。
omawrite のクロスコンパイル — マークダウン
Omarchy 4 "Quattro" に同梱されているエディタ — Windows 用。
誰もこれを求めていませんが、とにかくここにあります。クロードがこれの残りを書きました
README を参照してください。全角ダッシュなどをわざわざ削除するつもりはありません。楽しむ！
ポート全体は 42 行で、そのうち 30 行が #ifndef です。一行もいない
エディター、構文ハイライター、または変更が必要な QML です。組み込まれています
15秒くらい。
アプリケーション全体で唯一のプラットフォーム固有のコードは、
src/systemtheme.cpp の xdg-desktop-portal ブロックには、すでに
Qt ネイティブのフォールバックはその直下にあります。これは、Qt がレポートするためです。
すべてのプラットフォームで QStyleHints::colorScheme() を介したシステム カラー スキーム。それで、
「ポート」は主に D-Bus 呼び出しを削除し、すでに存在していたコードをそのまま使用します。
そこでその仕事が行われます。 QProcess 、 QSaveFile 、 QStandardPaths 、および印刷
パイプラインはすべて、記述どおりに移植可能です。
Linux から Windows 用 Qt をクロスコンパイルするためのすべてのガイドは、次のような指示で始まります。
Linux から Windows 用の Qt をコンパイルするには数時間を費やします。その必要はありません。
Fedora の mingw SIG には、事前に構築されたスタック全体 ( mingw64-gcc-c++ ) が同梱されています。
mingw64-qt6-qtbase 、 mingw64-qt6-qtdeclarative 、ロット — したがって、ツールチェーンは次のようになります
dnf をコンテナーにインストールすると、独自のコードのみがコンパイルされます。
1 つのギャップは、Fedora が Windeployqt をパッケージ化していないため、deploy.py がそのままであることです。
そのためには、 objdump -p を使用して PE インポート テーブルを調べ、それぞれの名前を解決します。
sysroot は、クロージャーが完了するまで再帰的に実行されます。約60行。それは動作します。
./b

uild-windows.sh # -> dist/ (~109M フォルダー)
./make-sfx.sh # -> 2 つの単一ファイル .exe ラッパー (それぞれ約 26M)
Docker が唯一の要件です。最初のスクリプトは上流でクローンを作成し、
パッチを適用し、イメージをビルドし、クロスコンパイルし、ランタイムをデプロイして、ストリップします。
make-sfx.sh は 2 つのフレーバーを生成します。
Portable は起動するたびに解凍コストを支払いますが、見た目や見た目には問題ありません。
毎日の編集者としては疲れます。どちらもファイル引数を渡しませんので、
.md をダブルクリックしても、そのファイルにはルーティングされません。インストールされたコピーは通常の .exe です
そしてそうするだろう。
Xvfb + wine でビルドを駆動する run-wine.sh もあります。
どこにも行く前にそれが機能するかどうかを知りたかったので、スクリーンショットを撮りました
それを実行します。
wine だけでなく (他のものは何も)、実際の Windows 11 VM でテストしました。
Print は注目に値します。omawrite が実際に実行する唯一の場所だからです。
マークダウン レンダリング — 新しい QTextDocument を構築し、 setMarkdown() を呼び出します。
したがって、PDF には実際の見出し階層があります。エディター自体がこれを行うことはありません。それは
生のテキストの上に構文ハイライター: # のようなブロック マーカーは表示されたままですが、
淡色表示になり、** のようなインライン マーカーは、
負の文字間隔の 1pt の背景色なので、ゼロに折りたたまれます。
幅。見出しは太字で表示され、サイズは変更されません。それは素晴らしい効果ですが、そうではありません
ウィシウィグ。
すべてのストリップ --strip-unneeded を実行します。 Fedora はこれらの DLL を完全なバージョンで出荷します。
デバッグシンボル。 libstdc++-6.dll だけで 25M でした。
omawrite が決してインポートしない 4 つのコントロール スタイル (FluentWinUI3、
イマジン、ユニバーサル、フュージョン）。
フロアはQtです。 icudata74.dll はそれ自体で 30M あり、静的リビルドのみです
-no-icu を指定するとそれが削除され、フォルダーを保存するために何時間ものコンパイルがかかります。
とにかく無視するつもりだった。私は断りました。あなたもそうすべきです。
アップストリームの 6.11 ではなく、Qt 6.8.3 (Fedora パッケージ) に対してビルドされています。
ノティ

ソース内の ng には新しい API が必要ですが、それは仮定であり、確実なものではありません。
証拠。
署名されていないため、SmartScreen は反対します。詳細 → とにかく実行します。
Omarchy テーマの同期は Omarchy 以外では意味がありません。 Windows のライト/ダークが表示され、
他には何もありません。
Windows はアクティブ化せずに個人用設定 UI をブロックするため、テストが必要になります
テーマの切り替えが面倒。 set-light.cmd / set-dark.cmd レジストリを書き込みます
直接的に評価し、それを完全に回避します。
omawrite は MIT (© David Heinemeier Hansson) です。ここのQtはLGPLv3です、リンクされています
これはLGPLが満足している配置です。付属のiA Writer
モノラルはOFL1.1です。これを静的 Qt ビルドに切り替えると、状況が変わり、
バイナリを渡す相手に対して、再リンク可能なオブジェクトを借りていることになります。
パッチは保護されているため、Linux ビルドはビットごとに影響を受けません。検証済みです。
引き続き libQt6DBus をリンクします。 Omacom の誰かが欲しいなら、すぐそこにあります
windows-port.patch を適用し、 clean を master に適用します。
あの巨大な難破船の朽ち果てた周囲、果てしなくむき出し、孤独で平坦なもの
砂浜が遠くまで広がっています。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

why? Contribute to WhiskeyTuesday/omawrite-win development by creating an account on GitHub.

GitHub - WhiskeyTuesday/omawrite-win: why? · GitHub
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
WhiskeyTuesday
/
omawrite-win
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit screenshots screenshots .gitignore .gitignore Dockerfile.mingw Dockerfile.mingw README-test.md README-test.md README.md README.md build-windows.sh build-windows.sh deploy.py deploy.py install.cmd install.cmd make-sfx.sh make-sfx.sh run-wine.sh run-wine.sh set-dark.cmd set-dark.cmd set-light.cmd set-light.cmd windows-port.patch windows-port.patch View all files Repository files navigation
My name is Ozymandias, King of Kings;
Look on my Works, ye Mighty, and despair!
Nothing beside remains.
Cross-compiles omawrite — the markdown
editor that ships with Omarchy 4 "Quattro" — for Windows.
Nobody asked for this but here it is anyway. Claude wrote the rest of this
README and I'm not going to bother removing all the em-dashes and stuff. Enjoy!
The entire port is 42 lines , and 30 of those are #ifndef . Not one line of
the editor, the syntax highlighter, or any QML needed to change. It builds in
about fifteen seconds.
The only platform-specific code in the whole application was the
xdg-desktop-portal block in src/systemtheme.cpp and that already had a
Qt-native fallback sitting directly underneath it, because Qt reports the
system colour scheme via QStyleHints::colorScheme() on every platform. So the
"port" is mostly deleting a D-Bus call and letting the code that was already
there do its job. QProcess , QSaveFile , QStandardPaths , and the print
pipeline are all portable as written.
Every guide to cross-compiling Qt for Windows from Linux opens by telling you to
spend several hours compiling Qt for Windows from Linux. You don't have to.
Fedora's mingw SIG ships the entire stack prebuilt — mingw64-gcc-c++ ,
mingw64-qt6-qtbase , mingw64-qt6-qtdeclarative , the lot — so the toolchain is
a dnf install in a container and nothing but your own code ever gets compiled.
The one gap is that Fedora doesn't package windeployqt , so deploy.py stands
in for it: walk the PE import table with objdump -p , resolve each name against
the sysroot, recurse until the closure is complete. About 60 lines. It works.
./build-windows.sh # -> dist/ (~109M folder)
./make-sfx.sh # -> two single-file .exe wrappers (~26M each)
Docker is the only requirement. The first script clones upstream, applies the
patch, builds the image, cross-compiles, deploys the runtime, and strips it.
make-sfx.sh produces two flavours:
Portable pays the unpack cost on every launch, which is fine for a look and
tiresome as a daily editor. Neither passes file arguments through, so
double-clicking a .md won't route to it; the installed copy is a normal .exe
and will.
There's also run-wine.sh , which drives the build under Xvfb + wine and
screenshots it, because I wanted to know whether it worked before I had anywhere
to run it.
I tested it on a real Windows 11 VM not just wine (but nothing else):
Print is worth a note, because it's the only place omawrite does real
markdown rendering — it builds a fresh QTextDocument and calls setMarkdown() ,
so the PDF has actual heading hierarchy. The editor itself never does this. It's
a syntax highlighter over raw text: block markers like # stay visible but
dimmed, and inline markers like ** are hidden by painting them in the
background colour at 1pt with negative letter-spacing so they collapse to zero
width. Headings are bolded, never resized. It's a nice effect and it is not
WYSIWYG.
strip --strip-unneeded on everything. Fedora ships those DLLs with full
debug symbols; libstdc++-6.dll alone was 25M.
Dropped the four Controls styles omawrite never imports (FluentWinUI3,
Imagine, Universal, Fusion).
The floor is Qt. icudata74.dll is 30M on its own and only a static rebuild
with -no-icu removes it, which costs hours of compiling to save a folder you
were going to ignore anyway. I declined. You should too.
Built against Qt 6.8.3 (what Fedora packages), not upstream's 6.11.
Nothing in the source needs the newer API, but that's an assumption, not a
proof.
Unsigned, so SmartScreen will object. More info → Run anyway .
Omarchy theme sync is meaningless off Omarchy. You get Windows light/dark and
nothing else.
Windows blocks the Personalization UI without activation, which makes testing
the theme switch annoying. set-light.cmd / set-dark.cmd write the registry
value directly and sidestep it entirely.
omawrite is MIT (© David Heinemeier Hansson). Qt here is LGPLv3, linked
dynamically, which is the arrangement LGPL is happy with. The bundled iA Writer
Mono is OFL 1.1. If you ever switch this to a static Qt build, that changes and
you'll owe relinkable objects to anyone you hand a binary to.
The patch is guarded so the Linux build is bit-for-bit unaffected — verified, it
still links libQt6DBus . If anyone at Omacom wants it, it's right there in
windows-port.patch and applies clean to master .
Round the decay of that colossal wreck, boundless and bare, the lone and level
sands stretch far away.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
