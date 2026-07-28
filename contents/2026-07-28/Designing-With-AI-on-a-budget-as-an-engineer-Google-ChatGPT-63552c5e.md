---
source: "https://www.kcoleman.me/2026/06/11/designing-with-ai-indie.html"
hn_url: "https://news.ycombinator.com/item?id=49087038"
title: "Designing With AI on a budget as an engineer: Google+ChatGPT"
article_title: "Designing With AI on a budget as an engineer: Google+ChatGPT | Kevin Coleman"
author: "speckx"
captured_at: "2026-07-28T17:18:58Z"
capture_tool: "hn-digest"
hn_id: 49087038
score: 1
comments: 0
posted_at: "2026-07-28T17:17:56Z"
tags:
  - hacker-news
  - translated
---

# Designing With AI on a budget as an engineer: Google+ChatGPT

- HN: [49087038](https://news.ycombinator.com/item?id=49087038)
- Source: [www.kcoleman.me](https://www.kcoleman.me/2026/06/11/designing-with-ai-indie.html)
- Score: 1
- Comments: 0
- Posted: 2026-07-28T17:17:56Z

## Translation

タイトル: エンジニアとして予算内で AI を使用して設計する: Google+ChatGPT
記事のタイトル: エンジニアとして予算内で AI を使用して設計する: Google+ChatGPT |ケビン・コールマン
説明: トークンはコード用であり、デザイン用ではありません。請求可能なトークンを一切使わずに、適切なサイズと品質で Apple、Google、ソーシャル メディア アカウント向けの美しいマーケティング資料を作成する方法を教えます。

記事本文:
ケビン・コールマン
コールマン法について 連絡先 Kevin Coleman — エンジニアリング リーダー 🥞モーニング スタック賞とプロジェクト
エンジニアとして予算内で AI を使用して設計する: Google+ChatGPT
トークンはコード用であり、デザイン用ではありません。請求可能なトークンを一切使わずに、適切なサイズと品質で Apple、Google、ソーシャル メディア アカウント向けの美しいマーケティング資料を作成する方法を教えます。
これがプロのデザイナーほど優れたものではないことは十分に承知していますが、ソフトウェア エンジニアである私が AI を使用せずに自分で行うよりもはるかに品質が高く、時間も短縮されます。信じられない場合は、以下のおまけのケーススタディで、AI、GPT4、そして今日の以前に私が行ったことを見てください。
ステップ 1: ChatGPT Web にスクリーンショットのリストを説明するよう依頼します。
サンプル アプリはデスクトップ iOS アプリになる予定でした。私はまず AI に、アプリとページの表示方法について 6 ～ 10 個のスクリーンショットのアイデアを教えてもらいました。
アプリのどのスクリーンショットを使用するかはすでにわかっていますが、何かが欠けている場合に備えて、メンタル モデルの盲点を特定するためのヒントは与えずに、AI にもう少し自由に考えてもらいます。
コード ベースにアクセスできるローカル エージェント ハーネス (Codex、Claude など) ではなく、ChatGPT Web を使用することに注意することが重要です。マーケティング資料では、解決策を明示的かつ排他的に示すのではなく、これによって解決される問題を伝える必要があります。 AI には、マーケティング資料には属さないコードベースの専門用語が含まれる傾向があります。
したがって、まったく新しい ChatGPT Web プロンプトは次のようになります。
{製品名} を構築しています。主な機能は {...} です。
MacOS App Store に含めるスクリーンショットのアイデアを 10 個教えてください。それは顧客への販売に役立ちます。
これら 10 個のアイデアから、そのページのコンテンツを説明する長いプロンプトを生成するよう AI に依頼します。
汎用の AI プロンプトを作成する

{...} のスクリーンショット画像を保存します
PNG を生成する
上記のプロンプトで ChatGPT Images 2.0 に新しい画像を生成するよう依頼します。
このスクリーンショットで説明する必要があるアプリのスクリーンショットを含めてください
たとえば、プロファイル ダッシュボードの画像を示すスクリーンショットが必要です。 AI は実際のスクリーンショットとプロンプトに記載されているマーケティング テキストを組み込んで、美しいものを生成します。場合によっては、画像生成ツールを使用していくつかの修正を行います。既存の画像を変更するよう依頼するか、さらに良いことに、編集を行わずに元の画像を白紙の状態で再生成するだけのほうが効果的です。
必要なスクリーンショットごとにこれらのベース PNG を取得すると、これらの PNG の問題は、サイズが間違っており、品質が低いことです。これは非常に鮮明で洗練されている必要があり、AI が生成したただのぼやけたクソみたいな写真であってはなりません。
chatGPT 画像を高品質のマーケティング資料に変換する
したがって、これら 3 つの画像を取得したら、次の段階で次のことを行います。
Google Stich (Google の Figma の競合他社) に移動し、これらの画像をツールにアップロードします。
_添付のスクリーンショット.png_
この画面にできるだけ近いイメージをコピーします。
その画像から HTML ページを生成します。
Google Stich では利用可能なトークンが限られているため、Google Stich を使用して HTML を生成し、後で (手動または Claude、ChatGPT Web などを使用して) HTML を調整できます。
構造の大部分はそこにありますが、フォントの色に問題があり、花輪の見た目はひどいものです。
次のステップは、手動および AI 支援による HTML の編集です。
お気に入りのテキスト エディタで .zip ファイルを開きます。これは単純な単一ファイルであるため、ChatGPT Web (有料アカウントのトークン使用量が無制限) を使用して HTML を微調整して更新できます。

コピー、スタイルなど。
デザインにとても満足したら。スクリーンショットを撮るには Firefox または Chrome を使用してください。必要な条件に合わせてビュー ウィンドウを調整するか、ターゲットに合わせて最終画像をトリミングしてサイズ変更する準備をしてください。
最終ステップ: スクリーンショット、サイズ変更、切り抜き、アップロード
お気に入りの Web ブラウザーの HTML で見栄えよく表示される場合は、ページ全体のスクリーンショットを撮り、お気に入りの画像編集ツールを使用して、各画像のサイズに合うように切り取ってサイズを変更します。
元の画像にはとても気に入ったアイコンがありましたが、Google も ChatGPT も HTML/SVG で適切に複製できませんでした。そこで、私がやったことは次のとおりです。
元の ChatGPT 画像を SVG 編集ツールで開きます
これらのベクトル点を使用してアイコン上を手動でトレースします
SVG をエクスポートし、生成された SVG を手動で作成した SVG に置き換えました。
これで、ChatBT によって生成されたアイコンを、非常に正確かつカラフルで美しい方法で複製することができました。ページを開いたら、
以下に、出会い系アプリ AvoVietnam のマーケティング画像を示します。
これらは私にはただ怠け者に見えます。シンプルなスクリーンショット、グラデーションの背景。
ChatGPT が画像を生成する前に、私は chatgpt に png モックアップに基づいて HTML テンプレートを作成するように依頼しました。残念ながら、元の画像を紛失してしまいましたが、これが出力です。
これは Google Stich がリリースされる前に作成されました。 Codex (5.3?) で画像を HTML に変換すると、フォント サイズと BLOB の位置がおかしくなってしまいました。 Codex 5.5 でもこの問題がまだ発生しています。
上記の Google Stich ワークフローと比較すると、これは高校のデザイン科の学生 (またはデザインスキルのないソフトウェア エンジニア) が作ったように見えます (笑)。

## Original Extract

Tokens are for code, not for design. I will teach you how to create beautiful marketing materials for Apple, Google, and social media accounts in the correct dimensions and quality at the cost of zero billable tokens.

Kevin Coleman
About Coleman laws Contact Kevin Coleman — Engineering Leader 🥞Morning Stack Awards & Projects
Designing With AI on a budget as an engineer: Google+ChatGPT
Tokens are for code, not for design. I will teach you how to create beautiful marketing materials for Apple, Google, and social media accounts in the correct dimensions and quality at the cost of zero billable tokens.
I fully acknowledge that this is not as good as a professional designer, but it is infinitely better quality and lower time than I, a software engineer, could do on my own without AI. If you don’t believe look at what I did pre-AI, GPT4, and today in the bonuse case study below.
Step one: ask ChatGPT Web to describe list of screenshots.
The example app was going to be a desktop iOS app. I first asked AI to give me six or ten screenshot ideas for how to display the app and the page.
While I already have an idea of which screenshots of the app that I would want to use, I let the AI think a little more freely just in case something’s missing without giving hints to indentify blind spots in my mental model.
It is important to note to use ChatGPT Web, and not any local agent harness (Codex, Claude, etc) that has access to the code base. The marketing materials should be communicate the problem this solves, not explicitly and exclusively your solution. AI has a tendency to include technical terms from the codebase that do not belong in marketing materials.
So, brand new ChatGPT web prompt:
I am building {product name}. The key features are {...}.
Give me 10 ideas for screenshots to include in the MacOS App Store. that will help me sell to my customers.
From those 10 ideas, ask AI to generate a longer prompt describing the content of that page.
write an AI prompt for generating an screenshot image of {...}
Generate the PNG
Ask ChatGPT Images 2.0 to generate a new picture with the prompt above
Include the screenshot of the app that this screenshot needs to be talking about
For example, I want to have a screenshot showing of the images for the profile dashboard. AI will incorporate the real screenshot with the marketing text described in the prompt to generate something beautiful. Sometimes I’ll make some corrections with the image generation tool, either asking it to modify the existing image or even better, just regenerating the original image, clean slate, no edits, tends to be more effective.
Once I have these base PNGs for each of the screenshots I want, the problem of these PNGs is that they are the wrong dimensions and are low quailty. We need this to be super crisp and polished and it can’t just be some blurry, shitty AI generated photo.
Converting chatGPT images to high quality marketing materials
So, the next stage we’re going to do is once we have those three images, we’re going to:
Go to Google Stich (Google’s Figma competitor) and upload these images into the tool.
_attached screenshot.png_
copy the image as closely as possible exactly this screen.
Generate an HTML page from that image.
The Google Stich has limited tokens available, so use Google Stich to generate the html and we can tune the html later (either by hand or with Claude, ChatGPT web, etc).
Most of the structure is there, but there are issues with font colors and the wreath just looks terrible.
Next step is manual and AI-assisted editing of the html
Open the .zip file in your favorite text editor. Since this is a simple single file, you can use ChatGPT web (which has unlimited token usage for paid accounts) to make minor adjustments to the html to update copy, styles, etc.
Once you’re super happy with the design. Use Firefox or Chrome to take the screenshot, be sure to either adjust the view window to match the required deminsions or be prepared to crop and resize the final image to match the target.
Final step: screenshot, resize, crop and upload
looks great in HTML in your favorite web browser, take a screenshot of the entire page and use your favorite image editing tool to crop and resize it so it fits the dimensions for each of the images.
There was an icon that I really liked in the original image that neither Google nor ChatGPT could properly replicate in HTML/SVG. So, what I did was:
Opened up the original ChatGPT image in an SVG editing tool
Manually traced over the icon with these vector points
Exported the SVG and replaced the generated svg with my manually created one.
Now I was able to replicate that icon generated by ChatBT, but in a very precise and colorful and beautiful way. Once you have your page opened up,
Below you can see the marketing images for the dating app AvoVietnam
These just look lazy to me. Simple screenshot, Gradient background.
Before ChatGPT had their image generation, I asked chatgpt to make an html template based on an png mockup. Unfortunately, I lost the original images, but this is the output.
This was created before Google Stich was released. Converting the images to html with Codex (5.3?) resulted in weird font sizes and locations of blobs. I still have this issue with Codex 5.5.
Compared to the Google Stich workflow above, this looks like a high school design student (or a software engineer with no design skills) made it haha.
