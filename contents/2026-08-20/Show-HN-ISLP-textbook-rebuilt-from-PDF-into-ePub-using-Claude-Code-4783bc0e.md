---
source: "https://github.com/govarthenan/islp-epub"
hn_url: "https://news.ycombinator.com/item?id=49375661"
title: "Show HN: ISLP textbook rebuilt from PDF into ePub, using Claude Code"
article_title: "GitHub - govarthenan/islp-epub: Rebuilds the ISLP statistics textbook from its PDF into a reflowable EPUB 3 for a 7-inch e-reader. Text-first mathematics, real HTML tables, 3,686 resolved cross-references. · GitHub"
image: "https://opengraph.githubassets.com/09de892dc162446dceaf0a24daeefb2cdfe31404b5e67e6db30b89fedf9abc99/govarthenan/islp-epub"
author: "govarthenan"
captured_at: "2026-08-20T15:24:32Z"
capture_tool: "hn-digest"
hn_id: 49375661
score: 1
comments: 0
posted_at: "2026-08-20T15:03:18Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ISLP textbook rebuilt from PDF into ePub, using Claude Code

- HN: [49375661](https://news.ycombinator.com/item?id=49375661)
- Source: [github.com](https://github.com/govarthenan/islp-epub)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T15:03:18Z

## Translation

タイトル: Show HN: クロード コードを使用して PDF から ePub に再構築された ISLP 教科書
記事のタイトル: GitHub - govarthenan/islp-epub: ISLP 統計教科書を PDF から 7 インチ電子書籍リーダー用のリフロー可能な EPUB 3 に再構築します。テキストファーストの数学、実際の HTML テーブル、3,686 の解決された相互参照。 · GitHub
説明: ISLP 統計教科書を PDF から 7 インチ電子リーダー用のリフロー可能な EPUB 3 に再構築します。テキストファーストの数学、実際の HTML テーブル、3,686 の解決された相互参照。 - ゴヴァルテナン/islp-epub

記事本文:
GitHub - govarthenan/islp-epub: ISLP 統計教科書を PDF から 7 インチ電子書籍リーダー用のリフロー可能な EPUB 3 に再構築します。テキストファーストの数学、実際の HTML テーブル、3,686 の解決された相互参照。 · GitHub
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
ゴバルテナン
/
islp-epub
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
62 コミット 62 コミット フォルダーとファイル
.github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE アセット 尻

ets docs docs 出力 出力 src src work work .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md DISCLAIMER.md DISCLAIMER.md LICENSE LICENSE LICENSE-DOCS LICENSE-DOCS NOTICE NOTICE README.md README.md Index.html Index.html package-lock.json package-lock.json package.json package.json pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
613 ページの統計教科書。PDF から、どの電子書籍リーダーでも実際に設定できるリフロー可能な EPUB 3 に再構築されています。
教育および実験でのみ使用してください。商用利用は禁止です。
その本は私のものではありません。このリポジトリは、統計学習入門、
Python のアプリケーションを使用して。そのテキスト、図、表、数学に関するすべての権利はそのまま残ります。
ギャレス・ジェームス、ダニエラ・ウィッテン、トレバー・ハスティ、ロバート・ティブシラニ、ジョナサン・テイラー、
スプリンガーの自然。著者は PDF を次の場所で無料で配布しています。
statlearning.com 。このプロジェクトは形式のみを変更します。それ
はコンテンツに対して何も主張せず、コンテンツに対していかなる権利も付与しません。持っていれば
本の権利があり、EPUB ファイルを削除したい場合は、問題を開くと削除されます。
削除されました。
数学の一部はAIモデルによって読み取られました。 531 の表現が次から切り取られました。
ページを作成し、AI ビジョン モデルによって転写されます。すべてが二度目のチェックを受け、サンプルが採取されました
別のモデル ファミリによって監査されました。99.8% 一致し、1 つのエラーが見つかり修正されました。
これでエラー率が測定されました。それは削除されませんでした。信頼する前に方程式を確認してください
勉強用、試験用、仕事用です。無料の PDF が権威であり、これではありません
変換。
保証はありません。ここにあるものは何も正しいとは約束されておらず、何もオープンになるとは約束されていません
あなたのデバイス。
完全な免責条項 → · 本の権利 → ·
ライセンス → · 貢献 →
1 つのビルドで 2 冊の本が作成される

PDF を 1 回通過するだけで済みます。どちらも同じ言葉を保持しており、同じです
図と同じ方程式。それらは数学の描画方法が異なるだけです。
どれが欲しいかわからない場合は、最初のものを選択してください。方程式が次のように表示される場合、
小さな空の箱、2番目のものを取ります。
Android 上の Moon+ Reader は、画像、インライン、またはページとして SVG をまったく描画しません。それは、
SVG というラベルが付いた小さなボックスに数式が配置されます。 MathML もありません。ラスター
画像が描く唯一の形式であるため、2 冊目の本が存在します。
Kobo — 出力/ISLP.epub を KOBOeReader ボリュームにコピーします。イジェクト時にピックアップされます。
kobo はプレーン EPUB を直接読み取ります。 Kobo 独自の特典、章ごとのページ数と本文
辞書検索、Calibre と KoboTouchExtended を使用して .kepub.epub に変換します
運転手。この本の内容はそれに依存するものではありません。
Kindle — Output/ISLP-raster.epub を Send-to-Kindle アドレスに送信します。アマゾンの
コンバーターは常に SVG を保持するとは限らず、PNG を保持します。
Android の Moon+ Reader — Output/ISLP-raster.epub を使用します。夜のテーマでは、
Moon+ はどちらも暗い色を宣言していないため、暗いページでは数学が黒く表示される可能性があります。
ページにスキームを適用したり、画面を反転したりします。 Moon+には夜間に画像を反転する設定があります。
Apple Books、Google Play Books、Thorium、KOReader、Calibre — 出力/ISLP.epub を開く
直接的に。
テストされたデバイス
この本は、最初は 1 台の Kobo Libra 2、7 インチの電子インク リーダー用に作られ、その後はすべての初期の
決定は 1264 × 1680 の 1 つのパネルに対して測定されました。
それだけでは十分ではありませんでした。 Moon+ を実行している Samsung Galaxy S23 から写真が返されました
Reader 、「表記法と単純な行列代数」で開きました。マトリックスがあるべき場所
小さな空箱がありました。その 1 枚の写真が 2 冊目の本が存在する理由であり、
src/make_ が理由です

project_epub.py が作成されました。これは、
同じ表現を 8 つの異なる方法で描画できるため、どのデバイスからでも 1 枚の写真を作成できます
そのデバイスが描画できるものと描画できないものを答えます。
つまり、「あらゆる電子書籍リーダーで読める」ということは、次のことを意味します: Kobo Libra 2 と Galaxy S23 で測定
Moon+ Reader を使用し、その他すべてのことについて考えました。上記に挙げた他のすべての読者は、
機能することが期待されます。どれも動作確認されていません。
読者が何か間違っていることを示した場合は、写真を送ってください。それが一番役に立つんです
誰でもこのプロジェクトに参加できます — を参照してください
デバイスレポート
および COTRIBUTING.md 。
PDF はページの画像です。幅もフォント サイズも 1 つだけで、段落が何なのかわかりません。
です。ほとんど読めない 7 インチのパネルでは、ページ全体を目を細めて見るか、
ズームされたフラグメントをパンします。
このプロジェクトは PDF を分解し、本を実際のテキスト (段落) として再構成します。
そのリフロー、フォントに合わせて拡大縮小する数学、折り返すテーブル、そして 3,686 の作業
相互参照。
2 つの文書で問題と答えが説明されています。コードの前にどちらかを読んでください。
PDF は、同じ PDF のインタラクティブ バージョンであるindex.html からレンダリングされます。
物語。 govarthenan.github.io/islp-epub でライブで読んでください。
ここでファイルリストからindex.htmlを開くと、ページではなく生のマークアップが表示されます。
ソース PDF はこのリポジトリにありません
この本は、Gareth 著『An Introduction to Statistical Learning, with Applications in Python』です。
ジェームズ、ダニエラ・ウィッテン、トレバー・ハスティ、ロバート・ティブシラニ、ジョナサン・テイラー
（スプリンガー、2023）。著者は statlearning.com でそれを配布しています。
パイプラインを実行するには、PDF を自分でダウンロードし、次のようにリポジトリのルートに置きます。
ISLP_ウェブサイト.pdf 。これは .gitignore 内にあり、コミットされることはありません。
サイトからダウンロードして、次の手順を実行します。
mv ~ /Downloads/ISLP_website.pdf 。
警官

はい、これは変換されました
パイプラインは PDF の 1 つの印刷に対して実行されました。ダウンロードしたものが同じであることを確認してください
1つ:
sha256sum ISLP_website.pdf # macOS: shasum -a 256 ISLP_website.pdf
プロパティ
値
SHA-256
278d3bdd49a8a480c2ff8e03245822caad8a3a48e81afd6d039c52c8fc13ad60
サイズ
20,053,984バイト
ページ
613
作られた
2023 年 8 月 14 日、Adobe InDesign 17.4
チェックサムが一致する場合、この README の番号がビルドに適用されます。そうでない場合は、あなたは
印刷が異なります。ビルドは引き続き実行されますが、各式がページごとに切り取られます。
とそのページ上の位置を確認するため、src/probe_structor.py から完全なパイプラインを再度実行します。
古い work/math_final.json を引き継がないでください。この README のカウントは移動する可能性があります。
uv sync # Python の依存関係
npm install # MathJax、数学を植字するため
UV 実行 Python src/build_epub.py
uv 実行 python src/validate_epub.py 出力/ISLP.epub
uv 実行 python src/validate_epub.py 出力/ISLP-raster.epub
src/build_epub.py の便利なオプション:
ラスター パスを変更した後、図面を確認します。
uv run python src/check_math_raster.py --sample 60 --contact-sheet
すべての方程式にインクが含まれており、グリフ参照に従っていることが確認されます。
ブラケットを work/math_stretchy_sheet.png に引き伸ばす 40 個の方程式を描画します。見てください
そのシートを自分の目で確認してください。レンダラーが間違うのは、括弧が伸びていることです。
このリポジトリの自動テストは、壊れたものを通常の相違点から分離します。
レンダラー。仕訳帳エントリ 020 を参照してください。
新しい PDF からの完全なパイプライン:
uv run python src/probe_structor.py # 章、インデックス、ページゾーンを見つける
uv run python src/extract_math_jobs.py # モデルが読み取る必要があるすべての式をトリミングします
# 転記と検証書き込み work/math_final.json
uv run python src/build_epub.py --out I

SLP.epub
uv 実行 python src/validate_epub.py 出力/ISLP.epub
uv run python src/build_index.py # ストーリーページを再生成
./src/make_story_pdf.sh # およびその PDF (GitHub で読むため)
何をするのか
本の一部
どのように引き継がれるのか
散文
文字ストリームから段落ごとに再構築されます。本自体の語彙に反してハイフネーションが元に戻される
インライン数学
下付き文字と上付き文字はグリフ ジオメトリから復元されるため、そのほとんどは HTML テキストのリフローのままになります
アクセント、スクリプト文字
MathJax によって植字された文字データから生成された LaTeX
分数、根号、行列
ページから切り取られ、AI ビジョン モデルによって読み取られ、ページと照合され、MathJax によってタイプセットされます。
すべてのタイプセット式
1 冊の本には SVG が含まれており、もう 1 冊の本には同じファイルが PNG として描画されています
数字
ベクター アートから 300 ppi、カラーで再レンダリング
テーブル
実際の HTML として読み戻すため、リーダーのフォントでリフローされます。
ラボコード
Jupyter の入力セルと出力セルは分離され、位置合わせは維持されます
相互参照
すべての内部リンクは EPUB アンカーに解決されます
脚注、欄外注、索引
小さなページに適した形式で保存
結果
章
15
段落
3,297
コードセル
1,137
数字
191、すべて 300 ppi で再レンダリング
テーブル
36、すべてリフロー可能なマークアップとして - 画像としてはなし
計算式を表示する
409
数学をテキストとして保存
6,690 アイテム中 5,974 (89%)
MathJax による数学タイプセット
716 の表現、926 の場所
— そのうち、ルールによって生成され、モデルはありません
185
— そのうち、AI モデルによってページから読み取られます
531
ページの絵として残された数学
0
内部リンクが解決されました
3,686 件中 3,686 件
EPUBサイズ
SVG で 8.7 MB、PNG で 11.9 MB
531 個のモデル読み取り式がそれぞれ独立して 2 回チェックされました。
印刷されたページとの比較: 481 件が同一、49 件の外観の違い、1 件が間違っている — 99.8%
合意。シ

角度の間違いを修正しました。次に、サンプルを別のモデルで監査しました
OpenAI の codex コマンド ライン ツールを介したファミリー: 45 人中 44 人が同意し、1 人は
意見の相違は現実であり、重要でした。パイプライン内のチェックではそれが検出されませんでした。 3
抽出の欠陥はそのために修正され、50 回の繰り返し監査で問題は見つかりませんでした。
意見の相違。 work/math_verification.json には詳細があり、ジャーナル エントリ 014 には
議論。
そのどれもが数学を確かなものにするものではありません。免責事項.md を参照してください。
1 つのデバイスのためではなく、読者のために選択される
以下のすべての決定は、6 インチ Kindle、カラー Kobo、電話、タブレット、またはデスクトップに適用されます。
最初に測定されたパネルだけでなく、窓にも適用されます。
本文に埋め込みフォントやフォント サイズはありません。読者独自のタイポグラフィ コントロールがそのまま残ります。
充電する。スタイルシート内のすべての長さは相対的なものです。
どちらの本でも数学のサイズは em 単位で表記されているため、テキストとともに大きくなります。これが理由です
ページからは何も切り取られません。 SVG ブックでは、各ファイルには独自の 2 つのルールが含まれています。
<img> は別のドキュメントであり、currentColor であるため、ダーク テーマに従っています。
だけではページからアクセスできません。 PNG にはルールを含めることができないため、ラスター ブックによってルールが変更されます。
代わりに、ページ スタイルシートの filter: invert(1) を使用して数学をオーバーします。 F

[切り捨てられた]

## Original Extract

Rebuilds the ISLP statistics textbook from its PDF into a reflowable EPUB 3 for a 7-inch e-reader. Text-first mathematics, real HTML tables, 3,686 resolved cross-references. - govarthenan/islp-epub

GitHub - govarthenan/islp-epub: Rebuilds the ISLP statistics textbook from its PDF into a reflowable EPUB 3 for a 7-inch e-reader. Text-first mathematics, real HTML tables, 3,686 resolved cross-references. · GitHub
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
govarthenan
/
islp-epub
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
62 Commits 62 Commits Folders and files
.github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE assets assets docs docs output output src src work work .gitignore .gitignore CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md DISCLAIMER.md DISCLAIMER.md LICENSE LICENSE LICENSE-DOCS LICENSE-DOCS NOTICE NOTICE README.md README.md index.html index.html package-lock.json package-lock.json package.json package.json pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
A 613-page statistics textbook, rebuilt from its PDF into a reflowable EPUB 3 that any e-reader can actually set.
Educational and experimental use only. No commercial use.
The book is not mine. This repository converts An Introduction to Statistical Learning,
with Applications in Python . All rights in its text, figures, tables and mathematics stay
with Gareth James, Daniela Witten, Trevor Hastie, Robert Tibshirani, Jonathan Taylor and
Springer Nature. The authors give the PDF away free at
statlearning.com . This project changes the format only. It
claims nothing over the content, and it grants no rights over the content. If you hold
rights in the book and you want the EPUB files removed, open an issue and they will be
removed.
Part of the mathematics was read by AI models. 531 expressions were cropped from the
page and transcribed by AI vision models. Every one was checked a second time, and a sample
was audited by a different model family — 99.8% agreement, one error found and corrected.
That measured the error rate. It did not remove it. Check any equation before you rely on
it for study, for an examination or for work. The free PDF is the authority, not this
conversion.
No warranty. Nothing here is promised to be correct, and nothing is promised to open on
your device.
Full disclaimer → · Rights in the book → ·
Licence → · Contributing →
One build makes two books from one pass over the PDF. Both hold the same words, the same
figures and the same equations. They differ only in how the mathematics is drawn.
If you do not know which one you want, take the first one. If the equations appear as
small empty boxes, take the second one.
Moon+ Reader on Android draws no SVG at all, in an image, inline, or as a page. It shows a
small box with the label SVG where the equation should be. It has no MathML either. A raster
image is the only form it draws, which is why the second book exists.
Kobo — copy output/ISLP.epub to the KOBOeReader volume; it is picked up on eject.
Kobo reads a plain EPUB directly. For Kobo's own extras, per-chapter page counts and in-text
dictionary look-up, convert it to .kepub.epub with Calibre and the KoboTouchExtended
driver. Nothing in the book depends on that.
Kindle — send output/ISLP-raster.epub to your Send-to-Kindle address. Amazon's
converter does not always keep SVG, and it keeps a PNG.
Moon+ Reader on Android — use output/ISLP-raster.epub . In the night theme the
mathematics may come out black on a dark page, because Moon+ neither declares a dark colour
scheme to the page nor inverts the screen. Moon+ has a setting that inverts images at night.
Apple Books, Google Play Books, Thorium, KOReader, Calibre — open output/ISLP.epub
directly.
The devices this was tested on
The book was built first for one Kobo Libra 2 , a 7-inch e-ink reader, and every early
decision was measured against that one panel at 1264 × 1680.
That was not enough. A photograph came back from a Samsung Galaxy S23 running Moon+
Reader , opened at "Notation and Simple Matrix Algebra". Where a matrix should have been
there was a small empty box. That one photograph is the reason the second book exists, and it
is the reason src/make_probe_epub.py was written: it builds a one-page EPUB that shows the
same expression drawn eight different ways, so that a single photograph from any device
answers what that device can and cannot draw.
So "reads on any e-reader" means this: measured on a Kobo Libra 2 and on a Galaxy S23 with
Moon+ Reader, and reasoned about for everything else. Every other reader named above is
expected to work. None of them is confirmed to work.
If your reader shows something wrong, send a photograph. It is the most useful thing
anybody can give this project — see
device report
and CONTRIBUTING.md .
A PDF is a picture of a page. It has one width, one font size, and no idea what a paragraph
is. On a 7-inch panel that is close to unreadable: you either squint at the whole page or you
pan around a zoomed fragment.
This project takes the PDF apart and puts the book back together as real text — paragraphs
that reflow, mathematics that scales with the font, tables that wrap, and 3,686 working
cross-references.
Two documents explain the problem and the answer. Read either one before the code.
The PDF is rendered from index.html , the interactive version of the same
story. Read it live at govarthenan.github.io/islp-epub .
Opening index.html from the file list here shows raw markup, not the page.
The source PDF is not in this repository
The book is An Introduction to Statistical Learning, with Applications in Python by Gareth
James, Daniela Witten, Trevor Hastie, Robert Tibshirani and Jonathan Taylor
(Springer, 2023). The authors give it away at statlearning.com .
To run the pipeline, download the PDF yourself and put it in the repository root as
ISLP_website.pdf . It is in .gitignore and it will never be committed.
Download it from the site, then:
mv ~ /Downloads/ISLP_website.pdf .
The copy this was converted
The pipeline was run against one printing of the PDF. Check that your download is the same
one:
sha256sum ISLP_website.pdf # macOS: shasum -a 256 ISLP_website.pdf
Property
Value
SHA-256
278d3bdd49a8a480c2ff8e03245822caad8a3a48e81afd6d039c52c8fc13ad60
Size
20,053,984 bytes
Pages
613
Made
14 August 2023, Adobe InDesign 17.4
If the checksum agrees, the numbers in this README apply to your build. If it does not, you
have a different printing. The build still runs, but each expression is cropped by its page
and its position on that page, so run the full pipeline again from src/probe_structure.py
and do not carry over an old work/math_final.json . The counts in this README can move.
uv sync # Python dependencies
npm install # MathJax, for typesetting the mathematics
uv run python src/build_epub.py
uv run python src/validate_epub.py output/ISLP.epub
uv run python src/validate_epub.py output/ISLP-raster.epub
Useful options on src/build_epub.py :
Check the drawing after any change to the raster path:
uv run python src/check_math_raster.py --sample 60 --contact-sheet
It confirms that every equation carries ink and that the glyph references are followed, and
it draws the 40 equations that stretch a bracket onto work/math_stretchy_sheet.png . Look at
that sheet with your own eyes: a stretched bracket is the thing a renderer gets wrong, and no
automatic test in this repository separates a broken one from ordinary differences between
renderers. See journal entry 020.
The full pipeline, from a fresh PDF:
uv run python src/probe_structure.py # find the chapters, the index, the page zones
uv run python src/extract_math_jobs.py # crop every expression a model must read
# transcription and verification write work/math_final.json
uv run python src/build_epub.py --out ISLP.epub
uv run python src/validate_epub.py output/ISLP.epub
uv run python src/build_index.py # regenerate the story page
./src/make_story_pdf.sh # and its PDF, for reading on GitHub
What it does
Part of the book
How it is carried over
Prose
Rebuilt paragraph by paragraph from the character stream; hyphenation undone against the book's own vocabulary
Inline mathematics
Sub- and superscripts recovered from glyph geometry, so most of it stays reflowing HTML text
Accents, script letters
LaTeX generated from the character data, typeset by MathJax
Fractions, radicals, matrices
Cropped from the page, read by an AI vision model, checked against the page, typeset by MathJax
Every typeset expression
SVG in one book, and the same file drawn as a PNG in the other
Figures
Re-rendered from the vector art at 300 ppi, in colour
Tables
Read back as real HTML, so they reflow with the reader's font
Lab code
Jupyter input and output cells kept apart, alignment preserved
Cross-references
Every internal link resolved to an EPUB anchor
Footnotes, margin notes, index
Kept, in forms that suit a small page
Results
Chapters
15
Paragraphs
3,297
Code cells
1,137
Figures
191, all re-rendered at 300 ppi
Tables
36, all as reflowable markup — none as pictures
Display equations
409
Mathematics kept as text
5,974 of 6,690 items (89%)
Mathematics typeset by MathJax
716 expressions, 926 places
— of those, generated by rule, with no model
185
— of those, read from the page by an AI model
531
Mathematics left as a picture of the page
0
Internal links resolved
3,686 of 3,686
EPUB size
8.7 MB with SVG, 11.9 MB with PNG
Every one of the 531 model-read expressions was checked a second time, independently,
against the printed page: 481 identical, 49 cosmetic differences, 1 wrong — 99.8%
agreement. The single error was corrected. A sample was then audited by a different model
family, through OpenAI's codex command line tool: 44 of 45 agreed, and the one
disagreement was real and mattered — no check inside the pipeline had caught it. Three
faults in the extraction were fixed because of it, and a repeat audit of 50 found no
disagreement. work/math_verification.json has the detail, and journal entry 014 has the
argument.
None of that makes the mathematics certain. See DISCLAIMER.md .
Choices made for the reader, not for one device
Every decision below holds on a 6-inch Kindle, a colour Kobo, a phone, a tablet or a desktop
window, and not only on the panel it was first measured against.
No embedded font, no font size on body — the reader's own typography controls stay in
charge. Every length in the stylesheet is relative.
Mathematics is sized in em in both books , so it grows with the text. This is why
nothing is cropped from the page. In the SVG book each file carries two rules of its own
that follow a dark theme, because an <img> is a separate document and currentColor
alone cannot reach it from the page. A PNG can carry no rules, so the raster book turns its
mathematics over with filter: invert(1) in the page stylesheet instead. F

[truncated]
