---
source: "https://millfolio.app/blog/millwright-ui-layers/"
hn_url: "https://news.ycombinator.com/item?id=48916095"
title: "Threading the AI in the UI for a personal app"
article_title: "Millwright: the three layers of a malleable UI — millfolio"
author: "winding"
captured_at: "2026-07-15T04:41:33Z"
capture_tool: "hn-digest"
hn_id: 48916095
score: 2
comments: 0
posted_at: "2026-07-15T03:53:29Z"
tags:
  - hacker-news
  - translated
---

# Threading the AI in the UI for a personal app

- HN: [48916095](https://news.ycombinator.com/item?id=48916095)
- Source: [millfolio.app](https://millfolio.app/blog/millwright-ui-layers/)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T03:53:29Z

## Translation

タイトル: 個人アプリの UI で AI をスレッド化する
記事のタイトル: Millwright: 柔軟な UI の 3 つのレイヤー — millfolio
説明: ブラインド AI を使用してダッシュボードを形成します。入力されたウィジェットの結果、バージョン管理されたボード仕様、ページ、各レイヤーはレンダリング前に検証されます。

記事本文:
Millwright: 順応性のある UI の 3 つの層 — millfolio millfolio Docs Blog github.com/millfolio ↗ Community ↗ ← ブログ Millwright: 順応性のある UI の 3 つの層
2026 年 7 月 14 日 · 2026 年 7 月 14 日更新 ミルライト アーキテクチャ オンデバイス AI
ローカル AI 予算に関する投稿は、ローカル GPU から価値を引き出す方法についてのものでした。これは、データからさらに多くのことを確認すること、つまりミルフォリオがどのようにレンダリングするかについてです。
モデルがマークアップ、スタイル、
DOM。この機能は Millwright と呼ばれ、3 つのネストされた構造として構築されています
各レイヤーは、クライアントがレンダリング前に検証するプレーン データ コントラクトです。
レイヤー 1: ウィジェット — マークアップではなく、入力された結果
ウィジェットのコンテンツは、同じサンドボックス化された小さなプログラムの出力です。
Ask での 1 回限りの質問に答えるプログラム。プログラムが HTML を返すことはありません。
マークダウン。結果仕様 (型付きブロックのバージョン管理された JSON エンベロープ) を返します。
{ "v" : 1 ,
"text" : "過去 3 か月間の上位 5 件の販売者。" 、
"データ" : [{ "種類" : "テーブル" ,
"headers" : [ "販売者" , "使用済み" ],
"行" : [[{ "タイプ" : "テキスト" , "値" : "ホールフーズ" },
{ "タイプ" : "お金" 、 "生" : 612.4 、 "テキスト" : "$612.40" }]] }] }
すべての値は入力されます。お金は {raw, text} として境界を越えます。
グラフの軸には数値が使用され、ラベルには正確にフォーマットされた文字列が使用されます。の
クライアントは表示文字列から「$612.40」を解析し返すことはありません。テーブル以外にも
KPI、時間/カテゴリシリーズ、全体の割合、およびオフラインがあります。
比例シンボルマップ。レンダラーはデータからビジュアライゼーションを選択します。
形状;プログラムはデータが何であるかを示すだけです。
これは、trusted-chrome の不変条件です。プログラムはデータを生成し、chrome は
インタラクションを管理します。テーブル列が販売者としてタグ付けされている場合
またはタグ、プログレではなくクロム

ram — セルをディープリンクに変換し、
保管庫の記録。生成されたプログラムは、リンクを挿入することはできません。
コントラクトにはスクリプト タグを挿入する場所がないため、スクリプト タグを挿入します。
レイヤ 2: ボード — セマンティックでバージョン管理された仕様
ボード自体も別の単純なドキュメントです。つまり、ウィジェットの順序付きリストです。
ID、タイトル、サイズヒント、および計算を行うプログラムへのポインタを含む
それ。編集が面白いところです。すべての変更 - タイルのサイズ変更、タイルの編集
プログラム、ウィジェットを削除 — 新しいコンテンツ アドレス バージョンのウィジェットを生成します。
spec (そのバイトの 16 進数の FNV-1a)、バージョン ログに追加されます。 「現在」
board はそのログへの単なるポインタであるため、 undo はポインタの移動であり、
何かを壊すような編集を行っても、以前のバージョンのボードを破壊することはできません。
候補仕様が受け入れられる前 — それがインラインからのものかどうか ✎
エディターまたはモデルから — バリデーターを渡します。ウィジェット ID は次のとおりである必要があります。
パスセーフで一意の参照プログラムが存在する必要があり、リモート URL は
完全に拒否され、構造上の制限が強制されます。このモデルは次のように提案しています。
バリデーターが破棄します。
レイヤー 3: ページ — 付加的なナビゲーション
ウィジェットのグループをページに昇格させることができ、独自のトップレベルを取得します。
「Ask」と「Vault」の横にあるナビゲーション ボタンをクリックし、ページをディゾルブするとそのウィジェットが返されます。
取締役会に。ページは同じ仕様ドキュメント (pages[] セクション) であり、同じです。
バージョン管理、同じバリデータ — 追加のルールが 1 つあります: ナビゲーションの変更は
添加剤のみ。生成された編集によりページを追加できます。決して名前を変更できない、または
組み込みのタブを削除します。何を検査するために依存する UI の部分
モデル自体はモデルによって編集できません。
モデルに UI を記述させる代わりにレイヤーを使用する理由
明らかな代替案は、モデルに HTML/JSX を出力させ、それをサンドボックス化することです。すること

このアプローチの欠点は、差分や検証ができないことです
構造上、ポインターを移動して元に戻すことはできず、すべてのレンダリングが
セキュリティ上の決定。 3 つのデータ レイヤーは、すべてのモデルで逆のトレードを行います。
貢献は、検査、バージョン管理、検証、拒否できる文書です。
そしてピクセルは常にアプリに同梱されているコードによって描画されます。
同じ階層化により公開デモが行われます
公開しても安全: デモボードは合成ボード上の同一の機械です。
vault、編集内容はブラウザの localStorage に保存されます - 同じハッシュ、同じ
バリデーター、同じクロム。
試してみましょう:demo.millfolio.app — ウィジェットを編集します
✎、壊して元に戻します。

## Original Extract

Shaping your dashboard with a blind AI — typed widget results, a versioned board spec, and pages, each layer validated before it renders.

Millwright: the three layers of a malleable UI — millfolio millfolio Docs Blog github.com/millfolio ↗ Community ↗ ← Blog Millwright: the three layers of a malleable UI
July 14, 2026 · updated July 14, 2026 millwright architecture on-device-ai
The local AI budget post was about how to get value out of your local GPU. This one is about seeing more from your data — how millfolio renders
model-generated analytics without ever letting a model touch markup, styles, or
the DOM. The feature is called Millwright , and it’s built as three nested
layers, each one a plain data contract the client validates before rendering.
Layer 1: the widget — typed results, not markup
A widget’s content is the output of a small program — the same sandboxed
programs that answer one-off questions in Ask. A program never returns HTML or
markdown. It returns a result spec : a versioned JSON envelope of typed blocks —
{ "v" : 1 ,
"text" : "Your top 5 merchants over the last 3 months." ,
"data" : [{ "kind" : "table" ,
"headers" : [ "Merchant" , "Spent" ],
"rows" : [[{ "type" : "text" , "value" : "WHOLE FOODS" },
{ "type" : "money" , "raw" : 612.4 , "text" : "$612.40" }]] }] }
Every value is typed — money crosses the boundary as {raw, text} so the
chart axis uses the number and the label uses the exact formatted string; the
client never parses "$612.40" back out of a display string. Besides tables
there are KPIs, time/category series, share-of-whole pies, and an offline
proportional-symbol map. The renderer picks the visualization from the data’s
shape; the program only says what the data is .
This is the trusted-chrome invariant: programs produce data, the chrome
manages interactions. When a table column is tagged as a merchant
or tag, the chrome — not the program — turns cells into deep links into your
Vault records. A generated program can’t inject a link any more than it can
inject a script tag, because there is nowhere in the contract to put one.
Layer 2: the Board — a semantic, versioned spec
The Board itself is another plain document: an ordered list of widgets, each
with an id, a title, a size hint, and a pointer to the program that computes
it. Editing is where it gets interesting. Every change — resize a tile, edit a
program, remove a widget — produces a new content-addressed version of the
spec (a 16-hex FNV-1a of its bytes), appended to a version log. The “current”
board is just a pointer into that log, so undo is a pointer move , and an
edit that breaks something can’t destroy the earlier version of the board.
Before any candidate spec is accepted — whether it came from the inline ✎
editor or from the model — it passes a validator: widget ids must be
path-safe and unique, referenced programs must exist, remote URLs are
rejected outright, and structural limits are enforced. The model proposes;
the validator disposes.
Layer 3: pages — additive navigation
A group of widgets can be promoted into a page — it gets its own top-level
nav button next to Ask and Vault, and dissolving the page returns its widgets
to the Board. Pages are the same spec document (a pages[] section), the same
versioning, the same validator — with one extra rule: navigation changes are
additive-only . Generated edits can add a page; they can never rename or
remove the built-in tabs. The parts of the UI you rely on to inspect what
the model did are not themselves editable by the model.
Why layers instead of letting the model write UI
The obvious alternative would be to have the model emit HTML/JSX and sandbox it. The downside of that approach is that you can’t diff it, you can’t validate
it structurally, you can’t revert it by moving a pointer, and every render is
a security decision. Three data layers give the opposite trade: every model
contribution is a document you can inspect, version, validate, and refuse —
and the pixels are always drawn by code that shipped with the app.
The same layering is what makes the public demo
safe to expose: the demo board is the identical machinery over a synthetic
vault, with edits kept in your browser’s localStorage — same hashing, same
validator, same chrome.
Try it: demo.millfolio.app — edit a widget with
✎, break it, and revert.
