---
source: "https://itsfoss.com/news/collabora-office-26-04/"
hn_url: "https://news.ycombinator.com/item?id=48779911"
title: "Collabora Office 26.04 Keeps AI Optional and Refines Writer and Calc"
article_title: "Collabora Office 26.04 Keeps AI Optional and Refines Writer and Calc"
author: "mmarian"
captured_at: "2026-07-03T21:57:33Z"
capture_tool: "hn-digest"
hn_id: 48779911
score: 2
comments: 0
posted_at: "2026-07-03T21:01:20Z"
tags:
  - hacker-news
  - translated
---

# Collabora Office 26.04 Keeps AI Optional and Refines Writer and Calc

- HN: [48779911](https://news.ycombinator.com/item?id=48779911)
- Source: [itsfoss.com](https://itsfoss.com/news/collabora-office-26-04/)
- Score: 2
- Comments: 0
- Posted: 2026-07-03T21:01:20Z

## Translation

タイトル: Collabora Office 26.04 は AI をオプションのままにし、Writer と Calc を洗練
説明: デスクトップ スイートの最初のメジャー アップデートは、コード 26.04 から機能を引き出します。

記事本文:
ログイン
サインアップ
サインイン
🧩 クイズとパズル
Collabora Office 26.04 は AI をオプションのままにし、Writer と Calc を洗練
Collabora Productivity は昨年、Collabora Online と同じレンダリング メカニズムに基づいて構築されたオフィス スイートですが、オフラインでの使用に重点を置いた Collabora Office を立ち上げ、デスクトップ エディター市場に参入しました。
彼らは、人々がインターフェイスの使い方を再学習することなく、オンラインでもオフラインでも同じ編集体験を得ることができるようにこれを考案しました。内部的には依然として LibreOffice コードですが、従来の VCL インターフェイスの代わりに、JavaScript、CSS、WebGL、および Canvas スタックを使用します。
コードベースが共有されているということは、Collabora 製品間で更新が迅速に行われることを意味します。そのため、Collabora Office の最初のメジャー アップデートではすでに CODE 26.04 の機能が取り込まれています。
ドキュメントをロードせずに 3 つのエディターのいずれかを開くと、新しいスタート画面が表示されます。最近使用したドキュメントがファイル タイプ アイコン付きで表示され、テンプレート ギャラリーが表示され、ファイルを開いた後は AI やドキュメント署名などのクイック アクセス機能が提供されます。
もちろん、ここでさらに大きな追加となるのは AI です。他のすべてのオープンソース オフィス スイートが Clanker ペイントに浸っているのを見るのにうんざりしているのはわかりますが、Collabora はデフォルトで AI をオフにしているため、心配する必要はありません。
これをオンにするということは、独自のモデル プロバイダーを選択するか、モデルを自己ホストするか、すでに信頼しているベンダーを選択して、独自の資格情報を接続することを意味します。 Collabora は完全に蚊帳の外にあり、あなたがドキュメントを手渡さない限り、アシスタントはドキュメントにアクセスできません。
有効にすると、 Writer でテキストの下書きと書き直しが行われ、自分でエラーを探す前に Calc で壊れた数式が整理され、大まかなメモが Impress で実際のスライドデッキに変換されます。
画像の生成や文書の要約も可能

手動で確認する時間がない場合。
残りのリリースでは、Office 26.04 には CODE 26.04 リリースに付属していたものが同梱されます。
そのため Writer では、挿入/削除と移動されたコンテンツを色分けして、誰がいつ変更したかをフラグで示す、再加工されたドキュメント比較ツールを利用できます。
これらはすべて並べて表示することも、変更の追跡パネルを通じて表示することもできます。新しい複数ページ ビュー、より豊富なスタイル プレビュー、ナビゲーター検索、Markdown インポート/エクスポートも追加されています。
Calcも同様に増加しました。ユーザーごとのシート ビューを使用すると、各ユーザーが他のユーザーのフィルターやレイアウトに触れることなく独自のフィルターやレイアウトを設定できます。また、新しいテーブル デザイン タブでは、計算されたピボット フィールドとともに適切なテーブル スタイルが提供されます。
数式エラーがセル上のフローティング ヘルパー ダイアログに表示されるようになり、シートをスクロールせずに問題を調べて修正できるようになりました。 TEXTSPLIT 、 HSTACK 、 WRAPROWS などの新しい関数のバッチも導入されました。
Impress では、フォローミー プレゼンテーションを使用すると、視聴者はプレゼンテーションの人の先にジャンプすることなく、自分で以前のスライドをスクロールして戻すことができます。同様に、スライドをセクションにグループ化でき、1 つのデッキで複数のスライド サイズを混在させることができ、マルチ モニターのサポートが以前よりも向上しています。
また、プレゼンテーションのフォント埋め込みも改善され、どこで開いても一貫してレンダリングされます。
Collabora Office は、Linux (Flatpak および Snap として)、Windows (Microsoft Store 経由)、macOS (App Store 経由) を含む幅広いプラットフォームで利用できます。
開発を追跡し、ソース コードにアクセスするには、Collabora の Gerrit インスタンスにアクセスしてください。
エンタープライズ サポートを探している場合は、Collabora がそれに取り組んでおり、このリリースは今後のサポートのプレビューとして機能すると述べています。 CEOのマイケル・ミークス氏に聞いた

エンタープライズ ユーザー、特に Linux 上に Collabora Office を展開しようとしているユーザーが期待できること。
企業は今日からこれの評価を開始でき、サポートは数週間以内に提供される予定です。 Collabora Online と共有されるより魅力的で人間工学に基づいた UX、組み込みの AI サポートなどにより、よりスムーズなワークフロー、少ないトレーニングを期待できます。企業ユーザーからのフィードバックをお待ちしています。
シェアする
シェアする
シェアする
シェアする
電子メール
フィードバック
著者について
スーラブ ルドラ
オープンソース ソフトウェア、カスタム PC ビルド、モータースポーツ、そしてこの世界の無限の可能性の探求に情熱を注ぐオタクです。
AI のスロップにうんざりし、元 Microsoft エンジニアが AI を使用しない新しいメモ帳を開発
Lumo 2.0: ChatGPT と Claude に代わる Proton のプライベートな代替手段がさらに改良されました
Linux ユーザーはこれを無料で入手できます! Brave Origin は他の人には $59.99 かかります
Firefox でこれらすべてができるのでしょうか?ほとんどのユーザーが触れない 21 の機能
Canonical の新しい AI ツールでは、文字を入力する代わりに Ubuntu と会話する必要があります
FOSS Weekly #26.27: KDE Linux の開発モード、ローカル AI、脱 Google Android、無料のターミナル スターター コース、KDE ステップなど
あらゆる種類の Linux ユーザー向けの 6 つのバックアップ ツール
DNS の提供により、AI エージェントは検証済みの ID を取得できる
広告ブロッカーを使用するというお客様の選択を尊重します。 FOSS はあなたのサポートに依存する独立した出版物です。
高品質の Linux コンテンツを誰でも無料で利用できるようにするために、私たちのサポートを検討してください。
AI のスロップにうんざりし、元 Microsoft エンジニアが AI を使用しない新しいメモ帳を開発
Lumo 2.0: ChatGPT と Claude に代わる Proton のプライベートな代替手段がさらに改良されました
DNS の提供により、AI エージェントは検証済みの ID を取得できる
Purism の Librem 16 Linux ラップトップに最大 11,944 ドル支払うことができます
Linux ユーザーはこれを無料で入手できます! Brave Origin は他の人には $59.99 かかります
FOSS ウィークリー ニュースレターでは、Linux の役立つヒントを学び、アプリケーションを発見します

新しいディストリビューションを探索し、Linux 世界の最新情報を常に入手してください。
あなたをより良い Linux ユーザーに
購読する
素晴らしい！受信箱を確認してリンクをクリックしてください。
申し訳ありませんが、問題が発生しました。もう一度試してください。
ナビゲーション
🧩 クイズとパズル
フェイスブック
ツイッター
RSS
ブルースカイ
マストドン
https://tabler-icons.io/i/brand-github からさらにアイコンのバリエーションをダウンロードします
ギットハブ
https://tabler-icons.io/i/brand-instagram からさらにアイコンのバリエーションをダウンロードします
インスタグラム
レディット
https://tabler-icons.io/i/brand-telegram からさらにアイコンのバリエーションをダウンロードします
電報
https://tabler-icons.io/i/brand-youtube からさらにアイコンのバリエーションをダウンロードします
ユーチューブ
©2026 FOSSです。
Digital Ocean でホストされ ($200 の無料クレジットを獲得)、Ghost & Rinne で出版されています。
システム
ライト
暗い
素晴らしい！サインアップが完了しました。
おかえり！正常にサインインしました。
It's FOSS への購読が完了しました。
成功！サインインするためのマジック リンクがあるかどうかメールを確認してください。
成功！お支払い情報が更新されました。

## Original Extract

The desktop suite's first major update pulls features from CODE 26.04.

Log in
Sign up
Sign in
🧩 Quizzes & Puzzles
Collabora Office 26.04 Keeps AI Optional and Refines Writer and Calc
Collabora Productivity got into the desktop editor market last year when they launched Collabora Office , an office suite built on the same rendering mechanism as Collabora Online, but focused on offline use.
They came up with this so people could get the same editing experience online or offline without needing to re-learn their way around the interface. It is still LibreOffice code under the hood, but instead of the traditional VCL interface, it uses a JavaScript, CSS, WebGL, and Canvas stack.
That shared codebase means updates move fast between Collabora's products, which is why the first major update for Collabora Office already pulls in features from CODE 26.04 .
Open any one of the three editors without a document loaded, and you will land on the new start screen. It will show your recent documents with file type icons, a template gallery, and provide quick access features like AI and document signing once a file's open.
Of course the bigger addition here is AI . I know you are tired of seeing every other open source office suite dipping themselves in Clanker paint, but fret not, as Collabora has kept AI off by default .
Switching it on means picking your own model provider, self-hosting a model, or choosing a vendor you already trust, and plugging in your own credentials. Collabora stays out of the loop entirely and the assistant gets no access to your documents unless you hand it over yourself.
When enabled, it drafts and rewrites text in Writer , sorts out broken formulas in Calc before you go hunting for the error yourself, and turns rough notes into an actual slide deck in Impress .
It can also generate images and summarize documents when you don't have time to go through them manually.
For the rest of the release , Office 26.04 ships what the CODE 26.04 release came with.
So for Writer , you get a reworked document comparison tool that color-codes insertions/deletions and any moved content while flagging who made each change and when.
All of that is viewable side by side or through the tracked changes panel. There's also a new multi-page view, richer style previews, Navigator search, and Markdown import/export.
Calc picked up just as much. Per-user sheet views let each person set their own filters and layout without touching anyone else's, and a new table design tab brings proper table styles along with calculated pivot fields.
Formula errors now show up in a floating helper dialog right on the cell, so you can inspect and fix the problem without scrolling through the sheet. A batch of new functions has also landed, including TEXTSPLIT , HSTACK , and WRAPROWS .
Over in Impress , follow-me presenting lets viewers scroll back through earlier slides on their own without ever jumping ahead of whoever's presenting. Similarly, slides can be grouped into sections, a single deck can mix multiple slide sizes, and multi-monitor support is better than before.
You also get better font embedding for presentations that render consistently wherever they're opened.
Collabora Office is available for a wide range of platforms, including Linux ( as a Flatpak and Snap ), Windows ( via the Microsoft Store ), and macOS ( via the App Store ).
For tracking development and access to the source code, you can visit Collabora's Gerrit instance.
If you were looking out for enterprise support , Collabora says that they are working on it and that this release acts as a preview of what's to come. I asked Michael Meeks , the CEO of Collabora Productivity, what enterprise users, particularly those looking to deploy Collabora Office on Linux , can expect.
Enterprises can start evaluating this today, support will arrive in a few weeks. They can look forward to smoother workflows, less training with a more attractive and ergonomic UX shared with Collabora Online, built-in AI support and more. We look forward to enterprise user feedback.
Share
Share
Share
Share
Email
Feedback
About the author
Sourav Rudra
A nerd with a passion for open source software, custom PC builds, motorsports, and exploring the endless possibilities of this world.
Sick of AI Slop, Former Microsoft Engineer Built a New, AI-less Notepad
Lumo 2.0: Proton's Private Alternative to ChatGPT and Claude Just Got Better
Linux Users Get This For Free! Brave Origin Costs $59.99 For Everyone Else
Firefox Can Do All This? 21 Features Most Users Never Touch
Canonical's New AI Tool Wants You to Talk to Ubuntu Instead of Type
FOSS Weekly #26.27: Dev Mode in KDE Linux, Local AI, De-Google Android, Free Terminal Starter Course, KDE Step and More
6 Backup Tools for Linux Users of All Kind
AI Agents Could Get Verified Identities, Courtesy of DNS
We respect your choice to use an ad blocker! It's FOSS is an independent publication that relies on your support.
Consider supporting us to keep quality Linux content free for everyone.
Sick of AI Slop, Former Microsoft Engineer Built a New, AI-less Notepad
Lumo 2.0: Proton's Private Alternative to ChatGPT and Claude Just Got Better
AI Agents Could Get Verified Identities, Courtesy of DNS
You Can Spend Up to $11,944 on Purism's Librem 16 Linux Laptop
Linux Users Get This For Free! Brave Origin Costs $59.99 For Everyone Else
With the FOSS Weekly Newsletter, you learn useful Linux tips, discover applications, explore new distros and stay updated with the latest from Linux world
Making You A Better Linux User
Subscribe
Great! Check your inbox and click the link.
Sorry, something went wrong. Please try again.
Navigation
🧩 Quizzes & Puzzles
Facebook
Twitter
RSS
Bluesky
Mastodon
Download more icon variants from https://tabler-icons.io/i/brand-github
Github
Download more icon variants from https://tabler-icons.io/i/brand-instagram
Instagram
Reddit
Download more icon variants from https://tabler-icons.io/i/brand-telegram
Telegram
Download more icon variants from https://tabler-icons.io/i/brand-youtube
Youtube
©2026 It's FOSS .
Hosted on Digital Ocean (get $200 in FREE credits) & Published with Ghost & Rinne .
System
Light
Dark
Great! You’ve successfully signed up.
Welcome back! You've successfully signed in.
You've successfully subscribed to It's FOSS.
Success! Check your email for magic link to sign-in.
Success! Your billing info has been updated.
