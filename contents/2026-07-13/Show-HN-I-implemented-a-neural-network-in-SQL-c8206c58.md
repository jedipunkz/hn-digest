---
source: "https://github.com/xqlsystems/xarray-sql/blob/claude/xarray-sql-mnist-demo/benchmarks/nn.py"
hn_url: "https://news.ycombinator.com/item?id=48897975"
title: "Show HN: I implemented a neural network in SQL"
article_title: "xarray-sql/benchmarks/nn.py at claude/xarray-sql-mnist-demo · xqlsystems/xarray-sql · GitHub"
author: "alxmrs"
captured_at: "2026-07-13T20:51:23Z"
capture_tool: "hn-digest"
hn_id: 48897975
score: 14
comments: 1
posted_at: "2026-07-13T20:00:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I implemented a neural network in SQL

- HN: [48897975](https://news.ycombinator.com/item?id=48897975)
- Source: [github.com](https://github.com/xqlsystems/xarray-sql/blob/claude/xarray-sql-mnist-demo/benchmarks/nn.py)
- Score: 14
- Comments: 1
- Posted: 2026-07-13T20:00:10Z

## Translation

タイトル: Show HN: SQL でニューラル ネットワークを実装しました
記事のタイトル: xarray-sql/benchmarks/nn.py (claude/xarray-sql-mnist-demo · xqlsystems/xarray-sql · GitHub)
説明: SQL を使用して Xarray データセットをクエリする実験。 GitHub でアカウントを作成して、xqlsystems/xarray-sql の開発に貢献してください。
HN テキスト: 2 週間前、私はベビームーンでギリシャのコルフ島にいました。移動中、私は GSoC インターンが重要な機能を配列データベース ライブラリである Xarray-SQL に送信するのを監督していました。彼は `to_dataset()` を追加しました。これにより、表形式モデル内の配列データをグリッド ラスターとして同時に考える間の往復処理が完了しました (プロジェクトの前提は、すべての Nd 配列を 2 次元にマッピングできるということであり、Nd 配列の直交する次元は表形式表現の主キーにすぎません)。この機能が存在したので、このデータ モデルが機能することを証明するためにどのようなデモを作成できるかについてチャットで話し合いました。熱波の最中に暖かいビーチで休憩し、冷たい塩水が新鮮なアイデアを与えてくれたので、私はアイデアを思いつきました。地理と気候のクエリの包括的な概要として、Coiled の地理空間ベンチマークのディスカッションを使用したらどうなるでしょうか。これらの一般的な操作はすべて、データ モデルが間違っているだけで、秘密裏にリレーショナルに行われているのでしょうか?ビーチでクロード コードを使用すると、これが事実であると思われることを確認できます。クロードと私は、地理科学および気候科学のすべての一般的な操作 (100 TB 範囲) が実際には秘密裏にリレーショナル操作であることを示すベンチマークを公開しています: https://github.com/xqlsystems/xarray-sql/blob/main/docs/geos... 。これらの例から最も驚くべきことは、コア操作であるリグリッドが単なるスパース行列とベクトルの積であるということでした。クロードは、このデータ モデルでは matmul は単なる `SUM(val * val) ... JOIN .. GROUP BY` であると私に指摘しました。これは e と直接的に類似しています

insum 表記法ですが、(おそらく) エレガントな SQL 構文で表現できます。この能力は、各部分の合計よりも優れているように見えました。イオニアン号の冷たい水に戻って、私はこれが何を意味するのかをより深く考えました。私は、Coiled ベンチマークのすべてが、深いところでは数値/配列コードで発生する_ポストプロセス_シミュレーションであったことを思い出しました。 SQL で計算できるのに、なぜこれらの物理計算をデータベースにもプッシュダウンできないのでしょうか?そのとき、私は思いつきました。線形代数に加えて、SQL で微積分も実行できるのであれば、できるかもしれません。 https://bsky.app/profile/al.merose.com/post/3mpbods7wts2y その後、JAX の実装に基づいて DataFusion の訪問者パターンの上に autograd を実装しました。私の簡略化された配列モデルでは、ヤコビアンの対角での偏微分のみを考慮していることがわかりました。これは、`grad()`、`jvp`、および `vjp` が単なる行単位の演算であることを意味します。次に、勾配を必要とするコイル状ベンチマークから一般的な物理計算を実装しました。ここから、データベースで自動採点できるのに、なぜニューラル ネットワークを作成できないことに気づきました。家に帰ってから、いくつかのスライドを作成し、この作品を DataFusion の最初のショーケースで発表しました: https://www.youtube.com/watch?t=1511&v=5o-4hL8vGPw&feature=y... この合成で、SQL は必ずしもニューラル ネットワークを作成するためのおもちゃの言語ではないことに気づきましたが、実際には、リレーショナル データベースの基本原則により、将来的には非常に望ましい可能性があることがわかりました。論理層は物理層から独立している必要があるということです。層。この性質が当てはまり、ニューラル ネットワークが一連の関係である場合、トレーニング用の SOTA 分散システムをより簡単に作成できるでしょうか?たとえば、データフローのグローバル論理プランが 1 つある場合、1,000 個以上の GPU で作業をより適切に分散できるでしょうか?セブ

eral の科学者やエンジニアと私は、https://xql.systems でリレーショナル配列のこの奇妙な世界を探索するために協力しています (参加したい場合は下部にある discord リンク)。

記事本文:
xarray-sql/benchmarks/nn.py (claude/xarray-sql-mnist-demo · xqlsystems/xarray-sql · GitHub)
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
xqlシステム
/
xarray-sql
公共
通知
「いいえ」を変更するにはサインインする必要があります

ティフィケーション設定
追加のナビゲーション オプション
コード
クロード/xarray-sql-mnist-demo ブレッドクラム
パスをコピーする もっとファイル アクションを責める もっとファイル アクションを責める 最新のコミット
履歴 履歴 488 行 (446 loc) · 18.8 KB claude/xarray-sql-mnist-demo ブレッドクラム
コピー パス トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード シンボル パネルを開く 編集および raw アクション 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100 101 102 103 104 105 106 107 108 109 110 111 112 113 114 115 116 117 118 119 120 121 122 123 124 125 126 127 128 129 130 131 132 133 134 135 136 137 138 139 140 141 142 143 144 145 146 147 148 149 150 151 152 153 154 155 156 157 158 159 160 161 162 163 164 165 166 167 168 169 170 171 172 173 174 175 176 177 178 179 180 181 182 183 184 185 186 187 188 189 190 191 192 193 194 195 196 197 198 199 200 201 202 203 204 205 206 207 208 209 210 211 212 213 214 215 216 217 218 219 220 221 222 223 224 225 226 227 228 229 230 231 232 233 234 235 236 237 238 239 240 241 242 243 244 245 246 247 248 249 250 251 252 253 254 255 256 257 258 259 260 261 262 263 264 265 266 267 268 269 270 271 272 273 274 275 276 277 278 279 280 281 282 283 284 285 286 287 288 289 290 291 292 293 294 295 296 297 298 299 300 301 302 303 304 305 306 307 308 309 310 311 312 313 314 315 316 317 318 319 320 321 322 323 324 325 326 327 328 329 330 331 332 333 334 335 336 337 338 339 340 341 342 343 344 345 346 347 348 349 350 351 352 353 354 355 356 357 358 359 360 361 362 363 364 365 366 367 368 369 370 371 372 373 374 375 376 377 378 379 380 381 382 383 384 385

386 387 388 389 390 391 392 393 394 395 396 397 398 399 400 401 402 403 404 405 406 407 408 409 410 411 412 413 414 415 416 417 418 419 420 421 422 423 424 425 426 427 428 429 430 431 432 433 434 435 436 437 438 439 440 441 442 443 444 445 446 447 448 449 450 451 452 453 454 455 456 457 458 459 460 461 462 463 464 465 466 467 468 469 470 471 472 473 474 475 476 477 478 479 480 481 482 483 484 485 486 487 488 # /// script # require-python = ">=3.12" # dependency = [ # "xarray_sql", # "xarr
[切り捨てられた]
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An experiment to query Xarray datasets with SQL. Contribute to xqlsystems/xarray-sql development by creating an account on GitHub.

Two weeks ago I was on my babymoon in Corfu, Greece. While in transit, I was overseeing a GSoC intern submit an important feature to my array database library, Xarray-SQL. He added `to_dataset()`, which completed the roundtrip between thinking of array data in a tabular model simultaneously as gridded rasters (the premise of the project is that every Nd array can be mapped to 2d, where orthogonal dims of the Nd array are just primary keys of a tabular representation). We discussed in chat, now that this feature existed, what demos could we make that would prove this data model works? With down time on a warm beach during a heatwave, cool salty water giving me fresh ideas, I had an idea: what if we used Coiled's Geospatial benchmark discussion as a comprehensive overview of geo and climate queries. Are all of these common operations secretly relational, just with the wrong data model? Using Claude Code on the beach, I can confirm that this seemed to be the case: Claude and I publish a benchmark that illustrated how every common operation in geo and climate sciences (at the 100 TB range) were actually secretly relational operations: https://github.com/xqlsystems/xarray-sql/blob/main/docs/geos... . Most surprisingly of all, from these examples was that a core operation, regridding, was just a sparse matrix-vector product. Claude had pointed out to me that in this data model, matmul was just a `SUM(val * val) ... JOIN .. GROUP BY`. This has a direct parallel to einsum notation, but can be expressed in (arguably) elegant SQL syntax! This capability seemed to be greater than the sum of it's parts. Back in the cool water of the Ionian, I thought about the implications of this more deeply. I reflected that, all of the Coiled benchmarks did, deep down, was _post process_ simulations that happen in numerical/array code. Why couldn't these physics calculations be push down into the database also, if we could so matmul in SQL? Then it hit me: maybe they could, if in addition to linear algebra, if SQL could do calculus! https://bsky.app/profile/al.merose.com/post/3mpbods7wts2y Later on, I implemented autograd on top of DataFusion's visitor pattern based on JAX's implementation. In my simplified array model, it turns out that we only care about partial differentiation on the diagonal of the Jacobian, meaning that `grad()`, `jvp` and `vjp` are just row-wise operations! I then implemented a common physics calculation from the coiled benchmark that required gradients. From here, I realized if I can autograd in the database, why can't I create a neural network? As I came back home, I created some slides, and presented this work to DataFusion's inaugural showcase: https://www.youtube.com/watch?t=1511&v=5o-4hL8vGPw&feature=y... I realized in this synthesis that SQL is not necessarily a toy language for writing neural networks, but in fact, may be highly desirable in the future due to the fundamental principles of relational databases: the logical layer should be independent from the physical layer. If that property holds, and a neural network is a series of relations, could we create a SOTA distributed system for training more easily? For example, if we had one global logical plan of dataflow, could we better distribute work on 1000+ GPUs? Several scientists and engineers and I are working together to explore this weird world of relational arrays at https://xql.systems (discord link at the bottom if you want to get involved).

xarray-sql/benchmarks/nn.py at claude/xarray-sql-mnist-demo · xqlsystems/xarray-sql · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
xqlsystems
/
xarray-sql
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
claude/xarray-sql-mnist-demo Breadcrumbs
Copy path Blame More file actions Blame More file actions Latest commit
History History 488 lines (446 loc) · 18.8 KB claude/xarray-sql-mnist-demo Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Open symbols panel Edit and raw actions 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100 101 102 103 104 105 106 107 108 109 110 111 112 113 114 115 116 117 118 119 120 121 122 123 124 125 126 127 128 129 130 131 132 133 134 135 136 137 138 139 140 141 142 143 144 145 146 147 148 149 150 151 152 153 154 155 156 157 158 159 160 161 162 163 164 165 166 167 168 169 170 171 172 173 174 175 176 177 178 179 180 181 182 183 184 185 186 187 188 189 190 191 192 193 194 195 196 197 198 199 200 201 202 203 204 205 206 207 208 209 210 211 212 213 214 215 216 217 218 219 220 221 222 223 224 225 226 227 228 229 230 231 232 233 234 235 236 237 238 239 240 241 242 243 244 245 246 247 248 249 250 251 252 253 254 255 256 257 258 259 260 261 262 263 264 265 266 267 268 269 270 271 272 273 274 275 276 277 278 279 280 281 282 283 284 285 286 287 288 289 290 291 292 293 294 295 296 297 298 299 300 301 302 303 304 305 306 307 308 309 310 311 312 313 314 315 316 317 318 319 320 321 322 323 324 325 326 327 328 329 330 331 332 333 334 335 336 337 338 339 340 341 342 343 344 345 346 347 348 349 350 351 352 353 354 355 356 357 358 359 360 361 362 363 364 365 366 367 368 369 370 371 372 373 374 375 376 377 378 379 380 381 382 383 384 385 386 387 388 389 390 391 392 393 394 395 396 397 398 399 400 401 402 403 404 405 406 407 408 409 410 411 412 413 414 415 416 417 418 419 420 421 422 423 424 425 426 427 428 429 430 431 432 433 434 435 436 437 438 439 440 441 442 443 444 445 446 447 448 449 450 451 452 453 454 455 456 457 458 459 460 461 462 463 464 465 466 467 468 469 470 471 472 473 474 475 476 477 478 479 480 481 482 483 484 485 486 487 488 # /// script # requires-python = ">=3.12" # dependencies = [ # "xarray_sql", # "xarr
[truncated]
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
