---
source: "https://snapname.app"
hn_url: "https://news.ycombinator.com/item?id=48342306"
title: "Show HN: I made a Gemma 4 Mac app that names screenshots with local AI"
article_title: "SnapName - AI-powered screenshot naming for macOS"
author: "joas_coder"
captured_at: "2026-06-01T00:35:30Z"
capture_tool: "hn-digest"
hn_id: 48342306
score: 7
comments: 6
posted_at: "2026-05-31T01:40:56Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I made a Gemma 4 Mac app that names screenshots with local AI

- HN: [48342306](https://news.ycombinator.com/item?id=48342306)
- Source: [snapname.app](https://snapname.app)
- Score: 7
- Comments: 6
- Posted: 2026-05-31T01:40:56Z

## Translation

タイトル: Show HN: ローカル AI でスクリーンショットに名前を付ける Gemma 4 Mac アプリを作りました
記事のタイトル: SnapName - AI を利用した macOS 用のスクリーンショットの命名
説明: SnapName はスクリーンショット フォルダーを監視し、ローカル AI を使用して Mac 上で役立つファイル名を提案します。
HN テキスト: 私はバンドルされた Gemma 4 モデル、具体的には Gemma E4B モデルに同梱される初めての macOS ユーティリティ アプリを作成しました。そのため、私のアプリ DMG のサイズは 5.3 GB になりましたが、この無料のローカル モデルが提供できる能力を考えると小さいサイズだと思います。 CPU では問題なく動作しますが、Apple Silicon GPU でも実行できますが、GPU によるパフォーマンスの向上には気づきませんでした (M5 チップでテスト)。これらのローカル軽量モデルとマルチモーダル モデルは、プライバシーが不可欠な新しいソフトウェア ツールにさまざまな可能性をもたらすと思います。

記事本文:
SnapName - AI を活用した macOS 用のスクリーンショットの命名
SnapName ダウンロード プライバシー サポート ダウンロード macOS 用のローカル AI スクリーンショットの命名
スクリーンショットアプリを保存してください。
SnapName はフォルダーを監視します。
新しいイメージには、ローカルで役立つ名前が付けられます。
アップルシリコンマック。 macOS 13以降。 CPUとGPUのサポート。
スクリーンショットが届くようにします。 SnapName が名前を処理します。
SnapName はスクリーンショット フォルダーを監視し、3 つの AI 名の提案を確認して選択できます。または、自動保存を選択すると、何もしなくても最初の AI 提案が自動的に適用されます。
自動モードは、macOS スクリーンショットの場所に従い、macOS マークが付けられたスクリーンショットをインポートします。
カスタム モードは、選択したフォルダーを監視し、新しい読み取り可能な画像をすべてインポートします。
提案された名前を自分で承認するか、監視フォルダー ファイルの自動保存をオンにします。
すでに気に入っているスクリーンショット ツールを使用します。
SnapName はキャプチャ ワークフローを置き換える必要はありません。 macOS、CleanShot、Shottr、または監視フォルダーに画像を保存するその他のツールによって作成されたスクリーンショットの名前を変更できます。
PNG、JPEG、HEIC、TIFF、AVIF ファイルは、macOS がデコードできる場合にサポートされます。 SnapName 独自のメニューとホットキー キャプチャ機能は、必要なときに引き続き利用できます。
スクリーンショットは AI 命名のためにアップロードされません。
SnapName はアプリ内に llama.cpp と Google Gemma 4 モデル ファイルをバンドルしているため、スクリーンショットのコンテンツは Mac 上でローカルに分析されます。画像は名前付けのために外部 AI モデルに送信されることはなく、スクリーンショットに対する完全なプライバシーが提供されます。
macOS の監視フォルダーのスクリーンショットの名前。 StillByte Labs によって構築されました。

## Original Extract

SnapName watches your screenshot folder and uses local AI to suggest useful filenames on your Mac.

I made my first macOS utility app that ships with a bundled Gemma 4 model, specifically the Gemma E4B one. It made my app DMG have 5.3 GB in size, but I think it is a small size for the power that this free local model can provide. It runs fine on CPU, but can also run on Apple Silicon GPU, although I did not notice any performance improvements with GPU (tested on a M5 chip). I think these local lightweight and multimodal models will open multiple possibilities for new software tools where privacy is essential.

SnapName - AI-powered screenshot naming for macOS
SnapName Download Privacy Support Download Local AI screenshot naming for macOS
Keep your screenshot app.
SnapName watches the folder.
New images get useful names locally.
Apple Silicon Mac. macOS 13 or later. CPU and GPU support.
Let screenshots arrive. SnapName handles the names.
SnapName watches your screenshot folder and lets you review and choose from three AI name suggestions. Or just choose auto-save to automatically go with the first AI suggestion without having to do anything.
Automatic mode follows the macOS screenshot location and imports macOS-marked screenshots.
Custom mode watches a folder you choose and imports every new readable image.
Approve a suggested name yourself, or turn on auto-save for watched-folder files.
Use the screenshot tool you already like.
SnapName does not need to replace your capture workflow. It can rename screenshots created by macOS, CleanShot, Shottr, or any other tool that saves images into a folder you watch.
PNG, JPEG, HEIC, TIFF, and AVIF files are supported when macOS can decode them. SnapName's own menu and hotkey capture features are still there when you want them.
Your screenshots are not uploaded for AI naming.
SnapName bundles llama.cpp and Google Gemma 4 model files inside the app, so screenshot content is analyzed locally on your Mac. Your images are never sent to an external AI model for naming, providing you with total privacy over your screenshots.
Watched-folder screenshot naming for macOS. Built by StillByte Labs .
