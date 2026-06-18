---
source: "https://www.motionscript.dev/"
hn_url: "https://news.ycombinator.com/item?id=48584300"
title: "The Open Source Alternative to Claude Design"
article_title: "MotionScript — Open Source Motion Design Tool | MotionScript"
author: "kevinoliveira"
captured_at: "2026-06-18T13:17:56Z"
capture_tool: "hn-digest"
hn_id: 48584300
score: 3
comments: 0
posted_at: "2026-06-18T12:26:17Z"
tags:
  - hacker-news
  - translated
---

# The Open Source Alternative to Claude Design

- HN: [48584300](https://news.ycombinator.com/item?id=48584300)
- Source: [www.motionscript.dev](https://www.motionscript.dev/)
- Score: 3
- Comments: 0
- Posted: 2026-06-18T12:26:17Z

## Translation

タイトル: Claude Design に代わるオープンソース
記事のタイトル: MotionScript — オープンソースのモーション デザイン ツール |モーションスクリプト
説明: コードを使用して素晴らしいモーション グラフィックスを作成します。 Web を利用したオープンソースの After Effects の代替品。

記事本文:
メイン コンテンツにスキップ モーション スクリプト ドキュメント API ブログ GitHub モーション スクリプト機能 ドキュメント API ブログ GitHub 機能 ドキュメント API ブログ GitHub Discord はじめに モーション グラフィックス
コード付き
Manim などのツールからインスピレーションを得たオープンソースのモーション デザイン ツールで、開発者や教育者がすべてコードから素晴らしいアニメーションを作成できるようにします。
正方形アニメーション 100% エクスポート 0 100 200 300 400 500 600 700 800 1x 00:00:000 Rect Text Ellipse 0f 20f 40f 60f 80f エディターの両方の長所
マウスを使用すると簡単にできることもあります。お気に入りの IDE を使用して TypeScript でアニメーションを作成します。 Web ベースのエディタを使用してオーディオと同期します。
Vite を利用すると、アニメーションのリアルタイム プレビューが変更に応じて自動的に更新されます。
ドキュメントを読む 手順 変更の手順
コードの実行によってアニメーションを定義します。何が起こるかを説明するジェネレータ関数を段階的に記述します。
コピー エクスポート クラス HelloScene extends Scene { * build() { this . set ( { 塗りつぶし : "#0D0F15" 、パディング : 80 } ) ;
const カード = createRef < Rect > ( ) ;
これ 。 add (< Rect ref = {card } width = { 240 } height = { 120 } fill = " #5ea8d8 " > < Text text = " Hello " fontSize = { 32 } fill = "white " /> </ Rect > );
// コードの実行によりアニメーションが定義されます yield * wait ( 0.5 ) ; yield * sequence (card().to({cornerRadius:60},0.8,easeOutBack()),card().to({rotation:360},0.8),);作成するために必要なものがすべて揃っています
モーション デザイナーとクリエイティブ開発者のための完全なツールキット。ベクター シェイプから GPU シェーダーまで、すべてコードによって駆動されます。
ロジックを説明するのと同じようにモーションを説明します。完全なタイプ セーフティを備えた TypeScript でアニメーションを作成し、コンポーネントを再利用して、Vite がすべての変更をライブ プレビューにストリーミングできるようにします。
使い慣れたフレックスボックス モデルを使用してシーンをレイアウトする — rows, co

塊、隙間、折り返し — コンテンツがアニメーションするにつれてすべてがリフローするのを観察します。
モーション用に構築されたフルテキスト エンジン: 可変フォント軸をアニメーション化し、グラデーションと画像の塗りつぶしでグリフをペイントし、ストロークとダッシュのレターフォームを作成し、フィットするように自動サイズ調整します。
SkSL で GPU シェーダ エフェクトを作成し、任意のノードに適用します。周辺光量や波紋から色収差まで、フレームごとにアニメーション化するリアルタイムのビジュアル。
MotionScript はオープンソースであり、コミュニティの貢献によって発展しています。バグの修正、機能の追加、ドキュメントの改善など、あらゆる貢献が重要です。
GitHub のスター 寄稿ガイドを読む Motion Script API Docs ブログ © 2026 MotionScript. MITライセンスに基づくオープンソース。

## Original Extract

Create stunning motion graphics with code. An open-source After Effects alternative powered by the web.

Skip to main content Motion Script Docs API Blog GitHub Motion Script Features Docs API Blog GitHub Features Docs API Blog GitHub Discord Get Started Motion graphics
with CODE
An open-source motion design tool, inspired by tools like Manim to help developers and educators create stunning animations, all from code!
Square Animation 100% Export 0 100 200 300 400 500 600 700 800 1x 00:00:000 Rect Text Ellipse 0f 20f 40f 60f 80f Editor Best of Both Worlds
Some things are easier with a mouse. Write animations in TypeScript with your favorite IDE; use a web-based editor to sync them with audio.
Powered by Vite, a real-time preview of your animation automatically updates upon any changes.
Read the Docs Procedural Procedural for a Change
Let the execution of your code define the animation. Write generator functions that describe what should happen — step by step.
Copy export class HelloScene extends Scene { * build ( ) { this . set ( { fill : "#0D0F15" , padding : 80 } ) ;
const card = createRef < Rect > ( ) ;
this . add ( < Rect ref = { card } width = { 240 } height = { 120 } fill = " #5ea8d8 " > < Text text = " Hello " fontSize = { 32 } fill = " white " /> </ Rect > ) ;
// The execution of your code defines the animation yield * wait ( 0.5 ) ; yield * sequence ( card ( ) . to ( { cornerRadius : 60 } , 0.8 , easeOutBack ( ) ) , card ( ) . to ( { rotation : 360 } , 0.8 ) , ) ; } } Features Everything you need to CREATE
A complete toolkit for motion designers and creative developers. From vector shapes to GPU shaders, all driven by code.
Describe motion the way you describe logic. Compose animations in TypeScript with full type-safety, reuse components, and let Vite stream every change into a live preview.
Lay out scenes with a familiar flexbox model — rows, columns, gaps, wrapping — and watch everything reflow as your content animates.
A full text engine built for motion: animate variable-font axes, paint glyphs with gradient and image fills, stroke and dash letterforms, and autosize to fit.
Author GPU shader effects in SkSL and apply them to any node. From vignettes and ripples to chromatic aberration — real-time visuals that animate frame by frame.
MotionScript is open source and thrives on community contributions. Whether you're fixing bugs, adding features, or improving docs — every contribution matters.
Star on GitHub Read Contributing Guide Motion Script API Docs Blog © 2026 MotionScript. Open source under the MIT License.
