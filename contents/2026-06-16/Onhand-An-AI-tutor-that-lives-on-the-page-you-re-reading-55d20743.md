---
source: "https://useonhand.com/"
hn_url: "https://news.ycombinator.com/item?id=48555374"
title: "Onhand: An AI tutor that lives on the page you're reading"
article_title: "Onhand: a reading companion that points"
author: "phineas1500"
captured_at: "2026-06-16T16:39:05Z"
capture_tool: "hn-digest"
hn_id: 48555374
score: 2
comments: 1
posted_at: "2026-06-16T13:55:17Z"
tags:
  - hacker-news
  - translated
---

# Onhand: An AI tutor that lives on the page you're reading

- HN: [48555374](https://news.ycombinator.com/item?id=48555374)
- Source: [useonhand.com](https://useonhand.com/)
- Score: 2
- Comments: 1
- Posted: 2026-06-16T13:55:17Z

## Translation

タイトル: Onhand: あなたが読んでいるページに住む AI 家庭教師
記事のタイトル: Onhand: ポイントを示す読書の友
説明: Onhand は、学習と研究のためのコンテキスト AI です。行を強調表示し、余白にメモを残し、ページ内で説明します。

記事本文:
手持ち
特長
仕組み
インストール
ツール
ドキュメント
◐
GitHub
★
↗
☞
Chromeに追加
Chrome ウェブストアのライブ配信
ストア・最新ビルド
アパッチ2.0
クロム用
OpenAI Codex で構築
手持ちの無料、コーデックス、プロバイダー キー
ポイントを教えてくれる読書のお供。 ☞
Onhand は、学習と研究のためのコンテキスト AI です。ページについて質問する
すでに開いています。行を強調表示し、余白にメモを残します。
そして、あなたがいるサイドパネルでそれを説明します。 2 番目のウィンドウやコピー＆ペーストはありません。
アテンションは、クエリとキーと値のペアを出力にマップします。モデル
クエリと各キーを比較し、結果の重みを使用します
値を平均化します。
私たちは特に注意を払います
スケーリングされたドット積の注意。クエリとキーには次元がある
d k ;値の次元は d v です。
この段落では、Q/K/V セットアップに名前を付けます。つまり、クエリをキーと比較し、結果の重みで値を平均します。
実装では、注意スコアは次の乗算によって生成されます。
転置キー行列によるクエリ行列、スケーリング
sqrt(d k ) 、およびソフトマックスで正規化します。
スコア = torch.matmul(クエリ, key.transpose(-2, -1))
スコア = スコア / math.sqrt(query.size(-1))
p_attn = スコア.ソフトマックス(dim=-1)
return torch.matmul(p_attn, value)
Softmax は、スケーリングされたスコアをアテンションの重みに変換します。
出力は、値ベクトルの重み付き混合です。
スケーリングされたドット積の注意とスケーリングが必要な理由を説明します。キーコードと方程式を教えてください。
スケーリングされたドット積アテンションは、すべてのクエリをすべてのキーと比較します。
スコアを sqrt(d k ) で除算し、ソフトマックスを適用します。
次に、それらの重みを使用して値を混合します。
1
クエリでは、各トークンが何を探しているかを尋ねます。
キーは、各トークンが一致できるものを記述します。
内積により互換性スコアが得られます。
2
スケーリングにより、大きな次元がソフトマックスになるのを防ぎます

鋭すぎる。
Softmax はスコアを重みに変換します。
3
マニキュル、余白にある小さな指さしの手は、
12 世紀以来、本では「ほら、ここが重要です」と書かれてきました。
Onhand は、モデルがペンを持ったジェスチャーを戻します。
余白内で 4 つのことを実行します。
Onhand は、あなたが尋ねるまで邪魔になりません。やればそうなる
良い学習パートナーが行うべき 4 つのこと。
質問に答える正確な語句または段落を金色で強調表示し、ページをスクロールして表示します。段落全体のブロックレベルのハイライト、フレーズのインラインハイライト。
ハイライトの隣に記事自身の声で書かれた 1 ～ 2 文が表示されるので、場所を見失うことなく読み続けることができます。
すべてのセッションは、ハイライト、メモ、トランスクリプト、ページのスナップショットとともに保存されます。明日、元の状態に戻ってきてください。
サイド パネルの [音声] を押して、紙を通して話します。リアルタイム モデルでは、ハイライトと表示されるテキストを確認できます。
3 つのステップ。 2番目のウィンドウはありません。
Wikipedia、arXiv 論文、Google ドキュメント、サブスタック。 Chromium でレンダリングできるものはすべて。 Onhand はアクティブなタブにアタッチされます。
平易な言語。 Onhand は、ページ、選択内容、目に見える見出しなど、答えを確定するために必要なものをすべて読み取ります。
行を強調表示し、欄外メモを削除し、ページに戻る引用とともにサイド パネルに回答します。その後、すべてを保存します。
Onhand は Chrome ウェブストアにあります。
承認された Chrome ウェブストア ビルドをワンクリックでインストールします。から始める
Onhand Free、OpenAI Codex でサインインするか、プロバイダー API キーを持参してください。
当店では現在サービス中です
Chrome 経由で更新します。
承認されたストア ビルドは通常どおりインストールされ、Chrome を通じて更新されます。
GitHub ZIP は、手動インストールを希望する場合、または特定のファイルが必要な場合にのみ使用してください。
アーティファクトを解放します。
# クローンを作成してビルドする
$ git clone https://g

ithub.com/フィニアス1500/オンハンド
$ cd 手持ち
$npmci
$ npm run build:extension
# ビルドを確認する
$ npm run steam:browser-runtime
✓ ブラウザランタイム対応 ·
# 実際のプロバイダー呼び出しの場合
$ OPENAI_API_KEY= … npm run steam:browser-runtime -- --real-openai
Chrome ウェブストアのリストを開き、[Chrome に追加] をクリックします。
Chrome は現在サービスを提供しており、Onhand は自動的に更新されます。
手持オプション ページを開きます。 Onhand Free を選択するか、OpenAI Codex でサインインするか、プロバイダー API キーを貼り付けます。
拡張機能を固定し、任意のページを開き、アイコンをクリックしてサイド パネルを起動します。パネルで Cmd/Ctrl+K を押してコンポーザーにフォーカスします。
手動インストールをご希望ですか?ダウンロードして解凍し、開発者モードを有効にして chrome://extensions からフォルダーを読み込みます。
以上です☞ 。文章を強調表示し、質問を入力し、オンハンドに をポイントさせます。
アパッチ2.0。オンハンドはオープンソースです。
セッションはマシンの chrome.storage.local に残ります。
プロバイダー キーが API キー モードでデバイスから離れることはありません。手持ちの無料送信
匿名診断を使用して、ホストされたオンハンド ワーカーを介してリクエストをモデル化します。
信頼性、クォータ、コスト、不正行為の監視に使用されます。ヘリウムなど
chrome.debugger をサポートする Chromium ベースのブラウザも動作します。
ページ上で動作する 29 個の小さなツール。
各ツールは、アクティブなタブに対する単一の監査可能な操作です。手持ち
それらを構成します。毎ターンごとにサイドパネルのスクリプトを読みます。
ブラウザ_get_viewport_Headings
ブラウザ_ get_visible_region_image
ブラウザ_open_pdf_in_onhand_viewer
browser_ run_js (オプション、制約あり)
Onhand とは何か、そうでないもの、そしてすべての UX コールを決定する原則。まずこれをお読みください。
何が出荷されるか、次に何が行われるか、PDF のシーケンス、検索、およびより充実したセッションの復元。
gpt-realtime-2 WebRTC 音声チューターが同じページベースのツールセットに対してどのように機能するか。

## Original Extract

Onhand is a contextual AI for learning and research. It highlights the line, leaves a note in the margin, and explains it in the page you

Onhand
Features
How it works
Install
Tools
Docs
◐
GitHub
★
↗
☞
Add to Chrome
Chrome Web Store live
Store · latest build
Apache 2.0
For Chromium
Built with OpenAI Codex
Onhand Free · Codex · provider keys
A reading companion that points. ☞
Onhand is a contextual AI for learning and research. Ask about the page
you already have open. It highlights the line, leaves a note in the margin,
and explains it in the side panel, where you are . No second window, no copy-paste.
Attention maps a query and key-value pairs to an output. The model
compares the query with each key, then uses the resulting weights
to average the values.
We call our particular attention
Scaled Dot-Product Attention. Queries and keys have dimension
d k ; values have dimension d v .
This paragraph names the Q/K/V setup: compare queries to keys, then average values with the resulting weights.
In implementation, the attention scores are produced by multiplying
the query matrix by the transposed key matrix, scaling by
sqrt(d k ) , and normalizing with softmax.
scores = torch.matmul(query, key.transpose(-2, -1))
scores = scores / math.sqrt(query.size(-1))
p_attn = scores.softmax(dim=-1)
return torch.matmul(p_attn, value)
Softmax turns the scaled scores into attention weights, so the final
output is a weighted mixture of the value vectors.
Explain scaled dot-product attention and why the scaling is needed; point me to the key code and equation.
Scaled dot-product attention compares every query with every key,
divides the scores by sqrt(d k ) , applies softmax,
then uses those weights to mix the values.
1
Queries ask what each token is looking for.
Keys describe what each token can match.
The dot product gives the compatibility score.
2
Scaling keeps large dimensions from making softmax too sharp.
softmax converts scores into weights.
3
The manicule , a small pointing hand in the margin, has meant
"look here, this part matters" in books since the twelfth century.
Onhand brings the gesture back, with the model holding the pen.
Four things, done in the margin.
Onhand stays out of the way until you ask. When you do, it does
the four things a good study partner would do.
Highlights the exact phrase or paragraph that answers your question, in gold, and scrolls the page to bring it into view. Block-level highlights for whole paragraphs, inline highlights for phrases.
Drops a sentence or two next to your highlight, written in the article's own voice, so you can keep reading without losing your place.
Every session is saved with its highlights, notes, transcript, and a snapshot of the page. Come back tomorrow to the exact state you left.
Press Voice in the side panel and talk through a paper. The realtime model can see your highlights and the visible text.
Three steps. No second window.
Wikipedia, an arXiv paper, a Google Doc, a Substack. Anything Chromium can render. Onhand attaches to the active tab.
Plain language. Onhand reads whatever it needs to ground the answer: the page, the selection, the visible headings.
Highlights the line, drops a margin note, answers in the side panel with citations back to the page. Then saves it all.
Onhand is on the Chrome Web Store.
Install the approved Chrome Web Store build in one click. Start with
Onhand Free, sign in with OpenAI Codex, or bring a provider API key.
The store is currently serving
and updates through Chrome.
The approved store build installs normally and updates through Chrome.
Use the GitHub ZIP only if you prefer manual installs or need a specific
release artifact.
# Clone and build
$ git clone https://github.com/ Phineas1500/Onhand
$ cd Onhand
$ npm ci
$ npm run build:extension
# Verify the build
$ npm run smoke:browser-runtime
✓ browser runtime ready ·
# For a real provider call
$ OPENAI_API_KEY= … npm run smoke:browser-runtime -- --real-openai
Open the Chrome Web Store listing and click Add to Chrome .
Chrome is currently serving and will update Onhand automatically.
Open the Onhand options page . Choose Onhand Free , sign in with OpenAI Codex , or paste a provider API key .
Pin the extension, open any page, and click the icon to launch the side panel . Press Cmd/Ctrl+K in the panel to focus the composer.
Prefer a manual install? Download , unzip it, then load the folder from chrome://extensions with developer mode enabled.
That's it ☞ . Highlight a passage, type a question, and let Onhand point .
Apache 2.0. Onhand is open source.
Sessions stay on your machine in chrome.storage.local .
Provider keys never leave your device in API-key modes; Onhand Free sends
model requests through the hosted Onhand Worker with anonymous diagnostics
for reliability, quota, cost, and abuse monitoring. Helium and other
Chromium-based browsers that support chrome.debugger work too.
Twenty-nine small tools that act on the page.
Each tool is a single, auditable operation against the active tab. Onhand
composes them; you read the script in the side panel after every turn.
browser_ get_viewport_headings
browser_ get_visible_region_image
browser_ open_pdf_in_onhand_viewer
browser_ run_js (optional, constrained)
What Onhand is, what it isn't, and the principles that decide every UX call. Read this first.
What's shipped, what's next, and the sequencing for PDFs, search, and richer session restore.
How the gpt-realtime-2 WebRTC voice tutor works against the same page-grounded toolset.
