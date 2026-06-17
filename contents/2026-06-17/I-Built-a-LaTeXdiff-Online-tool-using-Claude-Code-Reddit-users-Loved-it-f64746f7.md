---
source: "https://thelatexlab.com/latexdiff-online/"
hn_url: "https://news.ycombinator.com/item?id=48567336"
title: "I Built a LaTeXdiff Online tool – using Claude Code (Reddit users Loved it)"
article_title: "LaTeXdiff Online - Compare Two .tex Files in Your Browser"
author: "sahiltll"
captured_at: "2026-06-17T10:39:05Z"
capture_tool: "hn-digest"
hn_id: 48567336
score: 2
comments: 2
posted_at: "2026-06-17T08:13:22Z"
tags:
  - hacker-news
  - translated
---

# I Built a LaTeXdiff Online tool – using Claude Code (Reddit users Loved it)

- HN: [48567336](https://news.ycombinator.com/item?id=48567336)
- Source: [thelatexlab.com](https://thelatexlab.com/latexdiff-online/)
- Score: 2
- Comments: 2
- Posted: 2026-06-17T08:13:22Z

## Translation

タイトル: LaTeXdiff オンライン ツールを構築しました – クロード コードを使用してみました (Reddit ユーザーに好評)
記事のタイトル: LaTeXdiff オンライン - ブラウザーで 2 つの .tex ファイルを比較する
説明: ブラウザでオンラインで latexdiff を実行します。 2 つの .tex ファイルを貼り付け、追跡された変更を含む差分を取得します。 Perlもターミナルもインストールも必要ありません。無料。

記事本文:
LaTeXdiff オンライン - ブラウザーで 2 つの .tex ファイルを比較する
コンテンツにスキップ
Word から LaTex PDF から LaTex サービス 出版社の書式設定 Springer Elsevier IEEE arXiv MPDI ACM 論文の書式設定 論文の書式設定 オーストラリアの大学 米国の大学 ドイツの大学 その他 LaTex エラーの修正 LaTex 植字価格設定ツール LaTex ツール LaTeXdiff オンラインでの変更追跡 LaTeX 変更の追跡 識別子 → BibTeX DOI から BibTeX PubMed から BibTeX arXiv から BibTeX ISBN から BibTeX ファイル → BibTeX RIS から BibTeX EndNote から BibTeX NBIB から BibTeX ブログについて 見積もりを取得 見積もりを取得
★★★★★
5.0 · 500以上のプロジェクトを納品
LaTeXdiff オンライン
無料ツール、Perl インストールなし、ターミナルなし
LaTexdiff Online を使用して 2 つの .tex ファイルを比較し、何が変更されたかを即座に確認します。オリジナルおよび改訂された LaTeX ファイルを貼り付けて、ジャーナル、監修者、または共著者向けにマークアップされたバージョンを生成します。
.zip
ファイルを選択するか、ここにファイルをドロップします
クリックして置き換えます
改訂後.zip
.zip
ファイルを選択するか、ここにファイルをドロップします
クリックして置き換えます
詳細オプション
マーキングスタイル
下線変更バー CFONT CTRADITIONAL BOLD PDFCOMMENT
削除スタイル
カラー DVIPSCOL FONTSTRIKE セーフ
フロートの取り扱い
同一のフロートセーフ 従来型セーフ
数学マークアップ
粗い 細かい 全体をオフにする
差分ファイルの生成は非公開のままです。 diff を処理し、入力をすぐに破棄します。
差分の準備ができました。また、提出用の LaTeX フォーマットと組版も行います。
latexdiff は、2 つの LaTeX ソース ファイルを比較し、2 つのマクロ (挿入には \DIFadd{...}、削除には \DIFdel{...}) を使用してマークアップされた変更を含む 3 番目の LaTeX ファイルを生成する Perl スクリプトです。 3 番目のファイルを pdflatex でコンパイルすると、追加部分は青色で表示され、削除部分は赤色の取り消し線で表示されます。 ulem パッケージは取り消し線を処理します。
テキストの差分ではありません。普通の差分ツールw

ファイルを 1 行ずつ比較すると、LaTeX 構造を完全に見逃してしまう可能性があります。たとえ 1 単語しか変更していない場合でも、名前を変更した \section{} に完全な書き換えとしてフラグが立てられます。 latexdiff は LaTeX を理解します。段落、文、方程式、引用を個別の単位として差分し、文書を反対側でコンパイル可能な状態に保ちます。
これが重要なのは、ほとんどの研究者が次の 3 つの理由のいずれかで latexdiff を使用しているためです。それは、共著者と共同作業する際の草稿間での変更点の追跡、ジャーナル再投稿用の改訂ファイルの生成、または論文の改訂パス中の変更の健全性チェックです。多くのジャーナルは、著者にクリーンな原稿と、すべての変更を示すマークアップされた改訂版の両方をアップロードすることを要求しています。2 番目のファイルを作成する標準的な方法は latexdiff です。これらのワークフローはいずれも、一般的なテキストの差分では機能しません。
問題はインストールです。 latexdiff は、ほとんどの Linux および Mac セットアップに TeX Live および MiKTeX とともに同梱されていますが、Windows ユーザーはほぼ常に壁にぶつかります。latexdiff は Perl スクリプトであり、Perl は Windows にプリインストールされていません。結局、Strawberry Perl をインストールし、TeX Live を指定して、パスの問題に対処することになります。 Overleaf ユーザーは別の壁にぶつかります。latexdiff は Overleaf の変更履歴機能に組み込まれていますが、その機能は有料のプレミアム プランの背後に制限されています。 1 回限りの比較では、どちらのパスも労力を費やす価値はありません。それがこのツールが埋めるギャップです。インストールもサブスクリプションも必要なく、どのブラウザでも動作する無料のオンライン latexdiff です。
いいえ、Perl はサーバー側で実行されます。 2 つのファイルを貼り付けるだけで結果が得られます。
いいえ、diff は一般的なテキスト比較ツールであり、LaTeX 構造を理解できません。 latexdiff は LaTeX を解析し、段落レベルと数式レベルで diff を実行し、マークアップ マクロを含むコンパイル可能な LaTeX ファイルを書き込みます。この 2 つは互換性がありません

e.
最も一般的な原因は、ローカル インストールで ulem パッケージが欠落していることです。 latexdiff は、差分ファイルのプリアンブルに \RequirePackage[normalem]{ulem} を追加します。これがないと、その行でコンパイルが失敗します。 TeX ディストリビューションからインストールすると、ファイルがコンパイルされます。
2 番目に多い原因は、テンプレートの互換性がないことです。一部のジャーナル クラス ( elsarticle 、 acmart 、 sn-jnl ) は、 \DIFadd および \DIFdel と競合するコマンドを再定義します。サブタイプ オプションを SAFE に切り替えて、再試行してください。
はい、Overleaf プロジェクトが単一の .tex ファイルの場合は可能です。 Overleaf から両方のバージョンをダウンロードし (メニュー → ダウンロード → ソース)、.tex ファイルを開いて、それぞれを上のボックスに貼り付けます。このツールは、サブスクリプションなしで、Overleaf の有料の変更履歴機能と同じ差分を生成します。
2 つの完成したファイル バージョンを比較する場合は、はい。 Overleaf の Track Changes は、プレミアム プランのライブでインタラクティブなエディターです。このツールは、サブスクリプションなしで無料で、任意の 2 つの .tex ファイルの静的な差分を数秒で生成します。 Overleaf プロジェクトが \input{} または \include{} 経由で複数のファイルを使用している場合、現時点では最初にそれらをフラット化する必要があります。このツールの複数ファイル バージョンがリリースされる予定です。
はい、無料です。小規模なサーバー側サービスで実行され、月に数千のリクエストを快適に処理します。サインアップも有料枠もありません。
ジャーナル、研究論文、論文は、Springer、IEEE、arXiv、Elsevier、ACM、および大学のテンプレートにわたって正常にフォーマットされました。
IEEE、Springer、Elsevier、ACM、またはその他の雑誌テンプレート用にフォーマットされた原稿が必要ですか?
提出可能な LaTeX 組版を 49 ドルから扱います。ソースファイルを送信し、きれいに編集されジャーナルの仕様に一致する論文を返してください。
世界中の研究者や学者向けのプロフェッショナルな Word から LaTeX および PDF から LaTeX への変換。

## Original Extract

Run latexdiff online in your browser. Paste two .tex files, get a diff with tracked changes. No Perl, no terminal, no install. Free.

LaTeXdiff Online - Compare Two .tex Files in Your Browser
Skip to content
Word to LaTex PDF to LaTex Services Publisher Formatting Springer Elsevier IEEE arXiv MPDI ACM Thesis Formatting Thesis Formatting Australian Universities USA Universities German Universities Others Fix LaTex Errors LaTex Typesetting Pricing Tools LaTex Tools LaTeXdiff Online Track Changes LaTeX Changes Identifier → BibTeX DOI to BibTeX PubMed to BibTeX arXiv to BibTeX ISBN to BibTeX File → BibTeX RIS to BibTeX EndNote to BibTeX NBIB to BibTeX Blog About Get a Quote Get a Quote
★★★★★
5.0 · 500+ Projects Delivered
LaTeXdiff Online
Free Tool, No Perl install, No Terminal
Use LaTexdiff Online to compare two .tex files and instantly see what changed. Paste your original and revised LaTeX files, then generate a marked-up version for journals, supervisors, or co-authors.
.zip
Choose file or drop file here
Click to replace
revised.zip After
.zip
Choose file or drop file here
Click to replace
Advanced options
Marking style
UNDERLINE CHANGEBAR CFONT CTRADITIONAL BOLD PDFCOMMENT
Deletion style
COLOR DVIPSCOL FONTSTRIKE SAFE
Float handling
IDENTICAL FLOATSAFE TRADITIONALSAFE
Math markup
coarse fine whole off
Generate Diff Files stay private. We process the diff and discard the input immediately.
Your diff is ready. We also do submission-ready LaTeX formatting and typesetting.
latexdiff is a Perl script that compares two LaTeX source files and produces a third LaTeX file with the changes marked up using two macros: \DIFadd{...} for insertions and \DIFdel{...} for deletions. When you compile that third file with pdflatex, the additions appear in blue and the deletions appear in red strikethrough. The ulem package handles the strikethrough.
It is not a text diff. A plain diff tool would compare your files line by line and miss the LaTeX structure entirely — it would flag a renamed \section{} as a complete rewrite even if you only changed one word. latexdiff understands LaTeX. It diffs paragraphs, sentences, equations, and citations as separate units, and it leaves the document compilable on the other side.
This matters because most researchers use latexdiff for one of three reasons: tracking what changed between drafts when working with co-authors, generating a revision file for journal resubmission, or sanity-checking changes during a thesis revision pass. Many journals require authors to upload both a clean manuscript and a marked-up revision showing all changes — latexdiff is the standard way to produce that second file. None of these workflows work with a generic text diff.
The catch is installation. latexdiff ships with TeX Live and MiKTeX on most Linux and Mac setups, but Windows users almost always hit a wall — latexdiff is a Perl script, and Perl is not preinstalled on Windows. You end up installing Strawberry Perl, then pointing TeX Live at it, then dealing with path issues. Overleaf users hit a different wall: latexdiff is built into Overleaf's Track Changes feature, but that feature is gated behind their paid Premium plan. For a one-off comparison, neither path is worth the effort. That is the gap this tool fills — a free, online latexdiff that works in any browser, with no install and no subscription.
No. Perl runs server-side on our end. You just paste the two files and get a result.
No. diff is a generic text comparison tool that does not understand LaTeX structure. latexdiff parses the LaTeX, then diffs at the paragraph and equation level, then writes a compilable LaTeX file with markup macros. The two are not interchangeable.
The most common cause is a missing ulem package on your local install. latexdiff adds \RequirePackage[normalem]{ulem} to the diffed file's preamble — if you do not have it, the compile fails on that line. Install it from your TeX distribution and the file will compile.
The second most common cause is template incompatibility. Some journal classes ( elsarticle , acmart , sn-jnl ) redefine commands that conflict with \DIFadd and \DIFdel . Switch the subtype option to SAFE and try again.
Yes, if your Overleaf project is a single .tex file. Download both versions from Overleaf (Menu → Download → Source), open the .tex files, and paste each into the box above. The tool produces the same diff Overleaf's paid Track Changes feature would, without the subscription.
For comparing two completed file versions, yes. Overleaf's Track Changes is a live, interactive editor on the Premium plan. This tool generates a static diff of any two .tex files in seconds, free, without a subscription. If your Overleaf project uses multiple files via \input{} or \include{} , you currently need to flatten them first — a multi-file version of this tool is coming.
Yes, free. It runs on a small server-side service and serves a few thousand requests a month comfortably. No signup, no paid tier.
Journals, research papers, and theses formatted successfully across Springer, IEEE, arXiv, Elsevier, ACM, and university templates.
Need the manuscript formatted for IEEE, Springer, Elsevier, ACM, or another journal template?
We handle submission-ready LaTeX typesetting from $49. Send the source files, get back a paper that compiles cleanly and matches the journal spec.
Professional Word to LaTeX and PDF to LaTeX conversion for researchers and academics worldwide.
