---
source: "https://gist.github.com/joelonsql/c49129b151d7e77fb83250f675ffea07"
hn_url: "https://news.ycombinator.com/item?id=48339459"
title: "Claude Code vs. Codex: FRA challenge 75746d-2025"
article_title: "Swedish National Defence Radio Establishment - Challenge 75746d-2025: Claude Code vs Codex · GitHub"
author: "JoelJacobson"
captured_at: "2026-05-31T00:26:50Z"
capture_tool: "hn-digest"
hn_id: 48339459
score: 3
comments: 0
posted_at: "2026-05-30T18:48:09Z"
tags:
  - hacker-news
  - translated
---

# Claude Code vs. Codex: FRA challenge 75746d-2025

- HN: [48339459](https://news.ycombinator.com/item?id=48339459)
- Source: [gist.github.com](https://gist.github.com/joelonsql/c49129b151d7e77fb83250f675ffea07)
- Score: 3
- Comments: 0
- Posted: 2026-05-30T18:48:09Z

## Translation

タイトル: クロード コード対コーデックス: FRA チャレンジ 75746d-2025
記事のタイトル: スウェーデン国防無線設立 - チャレンジ 75746d-2025: クロード コード vs コーデックス · GitHub
説明: スウェーデン国防無線局 - チャレンジ 75746d-2025: クロード コードとコーデックス - fra_challenge.md

記事本文:
スウェーデン国防無線局 - チャレンジ 75746d-2025: クロード コードとコーデックス · GitHub
コンテンツにスキップ
-->
要点の検索
要点の検索
すべての要点
GitHub に戻る
サインイン
サインアップ
サインイン
サインアップ
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
コード、メモ、スニペットを即座に共有します。
要点オプションを表示
ZIPをダウンロード
スター
0
( 0 )
Gist にスターを付けるにはサインインする必要があります
フォーク
0
( 0 )
Gist をフォークするにはサインインする必要があります
埋め込む
この要点を Web サイトに埋め込みます。
シェアする
この要点の共有可能なリンクをコピーします。
HTTPS経由でクローンを作成する
Web URL を使用してクローンを作成します。
<script src="https://gist.github.com/joelonsql/c49129b151d7e77fb83250f675ffea07.js"></script> でこのリポジトリのクローンを作成します。
joelonsql/c49129b151d7e77fb83250f675ffea07 をコンピューターに保存し、GitHub デスクトップで使用します。
コード
改訂
1
埋め込む
オプションを選択してください
埋め込む
この要点を Web サイトに埋め込みます。
シェアする
この要点の共有可能なリンクをコピーします。
HTTPS経由でクローンを作成する
Web URL を使用してクローンを作成します。
<script src="https://gist.github.com/joelonsql/c49129b151d7e77fb83250f675ffea07.js"></script> でこのリポジトリのクローンを作成します。
joelonsql/c49129b151d7e77fb83250f675ffea07 をコンピューターに保存し、GitHub デスクトップで使用します。
ZIPをダウンロード
スウェーデン国防無線局 - チャレンジ 75746d-2025: クロード コード vs コーデックス
生
fra_challenge.md
FRA チャレンジ 75746d-2025: 統合ステップ 1 レポート
Claude レポートと Codex レポートを組み合わせたもの。
同じプロンプトがクロードとコーデックスの両方に使用されました。
「この Web ページに記載されているチャレンジを解決してください: https://challenge.fra.se/75746d-2025/ ルール: 解決策を検索することは許可されていません」

この問題については。最初のステップだけを解決してから、立ち止まって、問題をどのように解決したかを段階的に詳細に説明する Markdown ドキュメントを作成します。」
この結合レポートは、次のローカル マークダウン レポートに基づいています。
どちらのレポートも、課題の最初のステップのみを解決する試みについて説明しています。どちらも、外部のソリューション検索やウォークスルーは使用されていないと述べています。
2 つのレポートは最初の調査で一致しています。つまり、HTML ページには実際のフラグが直接含まれておらず、直接参照されたアセットは fralogo.jpg と footer_graphic.png だけでした。どちらのレポートでも、フッター画像が関連するアーティファクトであると特定されました。
クロードの報告書では旗は回収されなかった。オーディオ波形の可能性があるとしてフッター画像に焦点を当て、ファイル構造、メタデータ、文字列、LSB、および波形の分析を実行しました。この作業では、労力の少ないいくつかの隠し場所を除外しましたが、隠しテキストを抽出することなく中止されました。
Codex のレポートでは、 footer_graphic.png の上部バンドにあるかすかな低コントラストのピクセルを視覚的に強調することで、最初のステップのフラグを回復しました。明らかにされたフラグは次のとおりです。
フラグ{alpha_mail}
完全を期すために、Codex によって報告されたそのフラグの SHA1 は次のとおりです。
710e77d190bf3c45a7c966b214bdaf1686d84c35
Codex レポートには、ステップ 1 の後に停止し、SHA1 由来の次のステップ URL を開かなかったことが明示的に記載されています。
チャレンジ ページでは、フラグ形式の一致について説明します。
フラグガ{\S+}
回収したフラグのSHA1ダイジェストが次のステップのURLパスになると説明しています。両方のエージェントのプロンプトでは、最初のステップのみを解決してから停止する必要がありました。
両方のレポートで特定されたダイレクト ページ アセットは次のとおりです。
報告書は、ページ ソースにはサンプル フラグと形式の説明のみが含まれており、実際のチャレンジ フラグは含まれていないことに同意しています。
アスペクト
クロードレポート
コーデックスレポート
結果
解決していない。識別

フッター画像をメインアーティファクトとして保存しましたが、フラグは回復されませんでした。
解決しました。フッター画像のコントラストを強調したトリミングから flagga{alpha_mail} を復元しました。
経過時間
約10分。
1分54秒。
一次仮説
フッター グラフィックはオーディオ波形または振幅エンベロープとして扱われました。
フッター グラフィックは、かすかな隠れたコンテンツがないか視覚的に検査されました。
重要な成功ステップ
該当なし。報告書は捜索範囲を狭めたが、判明する前に中止した。
フッターの上部中央をトリミングし、明るさを倍にして、トリミングを拡大しました。
制約の処理
解決策の検索はありません。ページと直接アセットのみが使用されました。
解決策の検索はありません。ページと直接アセットのみが使用されました。
クロードの試みの概要
クロードはチャレンジ ページをフェッチし、予想されるフラグ形式を確認し、ページ ソースでリテラル フラグ文字列をチェックしました。 HTML 内で見つかったフラグのような文字列は、形式式と実際に動作する例だけでした。
次に、Claude は 2 つの直接アセットをダウンロードし、その構造を検査しました。
footer_graphic.png : PNG、1229 x 92 RGBA。
fralogo.jpg : JPEG、158 x 86。
PNG メタデータには、通常の作成/変更テキスト チャンクが含まれていました。
PNG IEND チャンクの後に後続データはありませんでした。
JPEG には、終了マーカーの後に後続バイトがありませんでした。
クロードはまた、印刷可能な文字列を検索し、ビット プレーンと LSB スタイルの隠蔽をテストし、おそらくサイドカーまたはオーディオ ファイル名を調査しました。これらのチェックのいずれもフラグを明らかにしませんでした。
クロード レポートの主な調査は、フッター グラフィックをオーディオ関連のアーティファクトとして扱うことに焦点を当てていました。振幅包絡線を抽出し、片側波形を観察し、繰り返し構造を特定し、画像がモールス信号、音声、SSTV、またはスペクトログラム データを表しているかどうかを検討しました。報告書は、画像がアナログ波形のサムネイルのように見えると結論付けました。

ただし、画像のみから理解できる音声を再構成することは不可能でした。
クロード報告書の妨げとなっていた問題は、波形の解釈が調査に時間を費やしてしまうことでした。レポートが上部バンドの低コントラストのピクセルを十分に追跡していなかったため、隠されたテキストは復元されませんでした。
Codex は同じ初期パスに従い、HTML をフェッチし、参照されている 2 つの画像を特定し、ファイルと文字列をチェックします。これらのチェックでは、文字通りのフラグは明らかになりませんでした。
視覚的に観察が成功しました。 Codex は、フッター画像の上部中央付近にかすかな低コントラストのコンテンツがあることに気付きました。その領域を ImageMagick で強化しました。
マジックアーティファクト/footer_graphic.png \
-アルファオフ \
-クロップ 600x35+450+0 \
-評価乗算60 \
artifacts/footer_top_mid_x60.png
次に、コーデックスはより厳密にトリミングして拡大しました。
マジックアーティファクト/footer_graphic.png \
-アルファオフ \
-クロップ 260x28+545+0 \
-評価乗算80 \
-サイズを1040x112に変更 \
アーティファクト/flag_crop_enhanced.png
強化された作物は次のことを示しました。
フラグ{alpha_mail}
Codex は次のように SHA1 ダイジェストを計算しました。
printf %s ' flagga{alpha_mail} ' |シャサム -a 1
結果として得られたダイジェストは次のとおりです。
710e77d190bf3c45a7c966b214bdaf1686d84c35
アーティファクト
フッター画像全体はシアン色の波形のようなグラフィックによって視覚的に支配されていますが、復元されたフラグは上部バンドのかすかなテキストから来ています。
統合された再現可能な手順
2 つのレポートを組み合わせた、再現可能な最短の方法は次のとおりです。
https://challenge.fra.se/75746d-2025/ で HTML ページを取得します。
HTML ソースに実際のフラグが存在しないことを確認します。
直接参照されるフッター アセット footer_graphic.png をダウンロードします。
画像を視覚的に検査し、上部中央のかすかな帯に焦点を当てます。
ImageMagick を使用してそのバンドをトリミングして明るくします。
公開されたフラグを読みます: flagga{alpha_mail} 。
プロンプトのみなので停止してください

y はステップ 1 を要求しました。
2 つのレポートは補完的です。クロードの報告書には、広範なネガティブ検索と除外されたいくつかの隠蔽方法が記載されています。 Codex レポートは、決定的な抽出ステップ、つまりフッター画像内の薄いテキストのコントラスト強調を提供します。総合的な結論は、ステップ 1 は footer_graphic.png 内の低コントラストのテキストを検出して強調することによって解決され、 flagga{alpha_mail} が生成されるということです。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Swedish National Defence Radio Establishment - Challenge 75746d-2025: Claude Code vs Codex - fra_challenge.md

Swedish National Defence Radio Establishment - Challenge 75746d-2025: Claude Code vs Codex · GitHub
Skip to content
-->
Search Gists
Search Gists
All gists
Back to GitHub
Sign in
Sign up
Sign in
Sign up
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Instantly share code, notes, and snippets.
Show Gist options
Download ZIP
Star
0
( 0 )
You must be signed in to star a gist
Fork
0
( 0 )
You must be signed in to fork a gist
Embed
Embed this gist in your website.
Share
Copy sharable link for this gist.
Clone via HTTPS
Clone using the web URL.
Clone this repository at &lt;script src=&quot;https://gist.github.com/joelonsql/c49129b151d7e77fb83250f675ffea07.js&quot;&gt;&lt;/script&gt;
Save joelonsql/c49129b151d7e77fb83250f675ffea07 to your computer and use it in GitHub Desktop.
Code
Revisions
1
Embed
Select an option
Embed
Embed this gist in your website.
Share
Copy sharable link for this gist.
Clone via HTTPS
Clone using the web URL.
Clone this repository at &lt;script src=&quot;https://gist.github.com/joelonsql/c49129b151d7e77fb83250f675ffea07.js&quot;&gt;&lt;/script&gt;
Save joelonsql/c49129b151d7e77fb83250f675ffea07 to your computer and use it in GitHub Desktop.
Download ZIP
Swedish National Defence Radio Establishment - Challenge 75746d-2025: Claude Code vs Codex
Raw
fra_challenge.md
FRA Challenge 75746d-2025: Combined Step 1 Report
Combined from Claude and Codex reports.
The same prompt was used for both Claude and Codex:
"Solve the challenge described on this webpage: https://challenge.fra.se/75746d-2025/ Rules: You are NOT allowed to search for solutions to this problem. Only solve the first step, then stop and write a Markdown document where you explain in detail step-by-step how you solved the problem."
This combined report is based on the following local Markdown reports:
Both reports describe attempts to solve only the first step of the challenge. Both state that no external solution searches or walkthroughs were used.
The two reports agree on the initial reconnaissance: the HTML page did not directly contain the real flag, and the only directly referenced assets were fralogo.jpg and footer_graphic.png . Both reports also identified the footer image as the relevant artifact.
The Claude report did not recover the flag. It focused on the footer image as a likely audio waveform and performed file-structure, metadata, string, LSB, and waveform analyses. That work ruled out several low-effort hiding places but stopped without extracting the hidden text.
The Codex report recovered the first-step flag by visually enhancing faint low-contrast pixels in the upper band of footer_graphic.png . The revealed flag was:
flagga{alpha_mail}
For completeness, the SHA1 reported by Codex for that flag was:
710e77d190bf3c45a7c966b214bdaf1686d84c35
The Codex report explicitly states that it stopped after step 1 and did not open the SHA1-derived next-step URL.
The challenge page describes a flag format matching:
flagga{\S+}
It explains that the SHA1 digest of the recovered flag becomes the URL path for the next step. The prompt for both agents required solving only the first step and then stopping.
The direct page assets identified in both reports were:
The reports agree that the page source contained only the example flag and format description, not the real challenge flag.
Aspect
Claude Report
Codex Report
Outcome
Not solved. Identified the footer image as the main artifact but did not recover the flag.
Solved. Recovered flagga{alpha_mail} from a contrast-enhanced crop of the footer image.
Elapsed time
Approximately 10 minutes.
1 minute 54 seconds.
Primary hypothesis
The footer graphic was treated as an audio waveform or amplitude envelope.
The footer graphic was inspected visually for faint hidden content.
Key successful step
N/A. The report narrowed the search space but stopped before the reveal.
Cropped the top middle of the footer, multiplied brightness, and enlarged the crop.
Constraint handling
No solution searches; only the page and direct assets were used.
No solution searches; only the page and direct assets were used.
Claude Attempt Summary
Claude fetched the challenge page, confirmed the expected flag format, and checked the page source for literal flag strings. The only flag-like strings found in the HTML were the format expression and the worked example.
Claude then downloaded the two direct assets and inspected their structure:
footer_graphic.png : PNG, 1229 by 92 RGBA.
fralogo.jpg : JPEG, 158 by 86.
PNG metadata contained ordinary create/modify text chunks.
There was no trailing data after the PNG IEND chunk.
The JPEG had no trailing bytes after the end marker.
Claude also searched printable strings, tested bit planes and LSB-style hiding, and probed likely sidecar or audio filenames. None of those checks revealed the flag.
The Claude report's main investigation focused on treating the footer graphic as an audio-related artifact. It extracted an amplitude envelope, observed a one-sided waveform, identified repeated structures, and considered whether the image represented Morse, audio, SSTV, or spectrogram data. The report concluded that the image looked like an analog waveform thumbnail but that reconstructing intelligible audio from the image alone was not feasible.
The blocking issue in the Claude report was that the waveform interpretation consumed the investigation. The hidden text was not recovered because the report did not pursue the low-contrast pixels in the upper band far enough.
Codex followed the same initial path: fetch the HTML, identify the two referenced images, and check the files and strings. Those checks did not reveal a literal flag.
The successful observation was visual. Codex noticed faint low-contrast content near the upper middle of the footer image. It enhanced that region with ImageMagick:
magick artifacts/footer_graphic.png \
-alpha off \
-crop 600x35+450+0 \
-evaluate multiply 60 \
artifacts/footer_top_mid_x60.png
Codex then made a tighter crop and enlarged it:
magick artifacts/footer_graphic.png \
-alpha off \
-crop 260x28+545+0 \
-evaluate multiply 80 \
-resize 1040x112 \
artifacts/flag_crop_enhanced.png
The enhanced crop showed:
flagga{alpha_mail}
Codex computed the SHA1 digest with:
printf %s ' flagga{alpha_mail} ' | shasum -a 1
The resulting digest was:
710e77d190bf3c45a7c966b214bdaf1686d84c35
Artifacts
The full footer image is visually dominated by a cyan waveform-like graphic, but the recovered flag came from faint text in the top band.
Consolidated Reproducible Procedure
The combined, shortest reproducible method from the two reports is:
Fetch the HTML page at https://challenge.fra.se/75746d-2025/ .
Confirm that the real flag is not present in the HTML source.
Download the directly referenced footer asset, footer_graphic.png .
Inspect the image visually and focus on the faint upper-middle band.
Crop and brighten that band using ImageMagick.
Read the revealed flag: flagga{alpha_mail} .
Stop, because the prompt only requested step 1.
The two reports are complementary. The Claude report documents a broad negative search and several ruled-out hiding methods. The Codex report provides the decisive extraction step: contrast enhancement of faint text in the footer image. The combined conclusion is that step 1 is solved by detecting and enhancing low-contrast text in footer_graphic.png , yielding flagga{alpha_mail} .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
