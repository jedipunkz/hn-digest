---
source: "https://blog.documentfoundation.org/blog/2026/11/24/updated-and-more-restrictive-policy-for-llm-usage-in-libreoffice-development/"
hn_url: "https://news.ycombinator.com/item?id=50624984"
title: "More restrictive policy for LLM usage in LibreOffice development"
article_title: "Updated (and more restrictive) policy for LLM usage in LibreOffice development - TDF Community Blog"
image: ""
author: "mksaunders2"
captured_at: "2026-11-24T08:42:28Z"
capture_tool: "hn-digest"
hn_id: 50624984
score: 2
comments: 0
posted_at: "2026-11-24T07:44:42Z"
tags:
  - hacker-news
  - translated
---

# More restrictive policy for LLM usage in LibreOffice development

- HN: [50624984](https://news.ycombinator.com/item?id=50624984)
- Source: [blog.documentfoundation.org](https://blog.documentfoundation.org/blog/2026/11/24/updated-and-more-restrictive-policy-for-llm-usage-in-libreoffice-development/)
- Score: 2
- Comments: 0
- Posted: 2026-11-24T07:44:42Z

## Translation

タイトル: LibreOffice 開発における LLM の使用に対するより制限的なポリシー
記事のタイトル: LibreOffice 開発における LLM の使用に関する更新された (そしてより制限された) ポリシー - TDF コミュニティ ブログ
説明: LibreOffice – スイート – には AI 機能が組み込まれていません。コンテンツを変更したり、データを共有したり、ユーザーの作業を使用してモデルをトレーニングしたりすることはありません。 AI 機能が必要な場合は、拡張機能を介して追加できますが、ソフトウェア自体には AI 機能は含まれていません。ソフトウェア開発について：
[切り捨てられた]

記事本文:
LibreOffice 開発における LLM の使用に関する更新された (そしてより制限的な) ポリシー - TDF Community Blog
コンテンツにスキップ
2026年11月24日火曜日
TDF コミュニティ ブログ
LibreOffice開発におけるLLMの使用に関する更新された(そしてより制限された)ポリシー
LibreOffice (スイート) には AI 機能は組み込まれていません。コンテンツを変更したり、データを共有したり、モデルをトレーニングするために作業内容を使用したりすることはありません。 AI 機能が必要な場合は、拡張機能を介して追加できますが、ソフトウェア自体には AI 機能は含まれていません。
ソフトウェア開発に関して: ソース コードに取り組む際の大規模言語モデル (LLM) の使用についての議論が進行中です。一部のフリー オープン ソース ソフトウェア (FOSS) プロジェクトは、LLM 支援コードの提出を処理する方法を検討する一方で、一般からの投稿を完全に締め切ることを選択しており、事実上「招待専用」クラブとなっています。私たちは、このような劇的な変化が起こらないことを願っています。
LibreOffice のような大規模なソフトウェア プロジェクトでは、AI 支援コーディング ツールの影響を避けることは不可能です。セキュリティ問題の報告では LLM の使用がますます増えており、LLM の承認を拒否することはソフトウェアやユーザーにとって利益になりません。サードパーティの依存関係、および依存するコンパイラにも、AI 支援コードが含まれています。私たちは、ある時点で LLM が使用されたという理由だけで、ソフトウェアをすべての人にとってより良いものにする貴重なコードの貢献を拒否したくありません。
しかし、私たちは人間が生成したコードに重点を置いており、開発における LLM の使用に関するポリシーをさらに厳格にしました。要約すると:
仕事で LLM を使用する前に許可を得る必要があります
コミットで LLM の使用を開示する必要があります
パッチの助けを借りて作成したパッチ内の変更されたすべてのコード行を、自分の言葉で説明できなければなりません。

LLM
LLM を使用して、コミット メッセージ、コメント、バグ レポート、および他の人へのメッセージを書き込むことはできません。
LLM はコードベースの特定の領域では許可されません
ポリシー全文を読むにはここをクリックしてください
LibreOffice 用の短編ビデオ (および新しい TikTok チャンネル)
行政への質問
LibreOffice と Google Summer of Code 2026: 結果
コメントを投稿するにはログインする必要があります。
LibreOffice 27.2 Alpha1 はテスト用に利用可能です
Windows 上の qt6 vcl プラグインのステータスと潜在的な次のステップ by Michael Weghorn
デバッグを高速化するための GDB インデックス キャッシュ
LibreOffice 用の短編ビデオ (および新しい TikTok チャンネル)
行政への質問
LibreOffice と Google Summer of Code 2026: 結果
Fb-cpp 1.0.0: Firebird 用の完全な機能を備えた最新の C++ ラッパー
LibreOffice 26.8.3の発表
LibreOffice 26.2.7 がダウンロード可能です
LibreOffice 26.8.2 は安定性、マルチメディア、互換性を向上させます
LibreOffice 26.8 と相互運用性: まだ読み取れないものを保存する
LibreOffice 26.8 は、プロフェッショナルなタイポグラフィー、世界の書記体系のより深いサポートを実現し、人工知能は使用しません
「最適な」インターフェイスはなく、適切なものだけが存在します
特に指定のない限り、この Web サイト上のすべてのテキストと画像は、クリエイティブ コモンズ 表示 - 継承 3.0 ライセンスに基づいてライセンスされています。
これには、Mozilla Public License (MPLv2) に基づいてライセンスされている LibreOffice のソース コードは含まれません。
LibreOffice および The Document Foundation は、対応する登録所有者の登録商標であるか、1 つ以上の国で商標として実際に使用されています。それぞれのロゴとアイコンも国際著作権法の対象となります。その使用については、当社の商標ポリシーで説明されています。
当社のウェブサイトをご覧いただき、当社のサービスにご興味をお持ちいただきありがとうございます

そして製品。お客様の個人データの保護は当社にとって重要な懸念事項であるため、「詳細情報」リンクをクリックして当社のプライバシー ポリシー ページにアクセスしてください。このページは別のブラウザ タブで開きます。そこでは、お客様が当社の Web サイトにアクセスした際に当社がどのような情報を収集するか、その情報がどのように処理されるか、またどのように使用されるかについて説明されています。
当社のプライバシー ポリシー ページをよく読んだ後、ブラウザのタブを閉じてこのページに戻り、このテキストの下にある [設定を保存] ボタンをクリックして同意し、ダイアログを閉じて Web サイトに戻ります。
当社は、お客様の個人データを損失や悪用から保護するために、必要な技術的および組織的なセキュリティ対策をすべて講じています。お客様のデータは、一般の人がアクセスできない安全な操作環境に保存されます。

## Original Extract

LibreOffice – the suite – includes no AI features built in. It doesn’t try to change your content, share your data, or use your work to train models. If you want AI features, you can add them via extensions – but the software itself has zero AI features inside. Regarding software development: there
[truncated]

Updated (and more restrictive) policy for LLM usage in LibreOffice development - TDF Community Blog
Skip to content
Tuesday, November 24, 2026
TDF Community Blog
Updated (and more restrictive) policy for LLM usage in LibreOffice development
LibreOffice – the suite – includes no AI features built in. It doesn’t try to change your content, share your data, or use your work to train models. If you want AI features, you can add them via extensions – but the software itself has zero AI features inside.
Regarding software development: there are ongoing discussions about the use of large language models (LLMs) when working on source code. Some free and open source software (FOSS) projects have chosen to close contributions from the public altogether, while they work out how to handle LLM-assisted code submissions – effectively becoming “invite-only” clubs. We hope to never have to make such a drastic change.
Avoiding any influence of AI-assisted coding tools in a large software project like LibreOffice is not possible. Security issue reports involve more and more use of LLMs, and refusing to acknowledge them would not benefit the software or our users. Our third party dependencies – and even the compilers we depend on – include AI-assisted code. We don’t want to reject any valuable code contributions that make the software better for everyone, just because an LLM was used at one point.
But our focus is strongly on human-generated code, and we have made our policy for using LLMs in development even stricter. In summary:
You must get permission before using an LLM in your work
You must disclose the use of an LLM in commits
You must be able to explain – in your own words – every changed line of code in a patch you created with the help of an LLM
You cannot use an LLM to write commit messages, comments, bug reports, and messages to others
LLMs are not permitted in certain areas of the codebase
Click here to read the full policy
Short-form videos (and new TikTok channel) for LibreOffice
Questions for Public Administrations
LibreOffice and Google Summer of Code 2026: The results
You must be logged in to post a comment.
LibreOffice 27.2 Alpha1 is available for testing
Status and potential next steps for qt6 vcl plugin on Windows by Michael Weghorn
GDB index cache for faster debugging
Short-form videos (and new TikTok channel) for LibreOffice
Questions for Public Administrations
LibreOffice and Google Summer of Code 2026: The results
Fb-cpp 1.0.0: A feature-complete modern C++ wrapper for Firebird
Announcement of LibreOffice 26.8.3
LibreOffice 26.2.7 is available for download
LibreOffice 26.8.2 improves stability, multimedia and compatibility
LibreOffice 26.8 and interoperability: preserving what cannot yet be read
LibreOffice 26.8 brings professional typography, deeper support for the world’s writing systems, and no artificial intelligence
There Is No “Best” Interface, Only the Right Fit
Unless otherwise specified, all text and images on this website are licensed under the Creative Commons Attribution-Share Alike 3.0 License .
This does not include the source code of LibreOffice, which is licensed under the Mozilla Public License ( MPLv2 ).
LibreOffice and The Document Foundation are registered trademarks of their corresponding registered owners or are in actual use as trademarks in one or more countries. Their respective logos and icons are also subject to international copyright laws. Use thereof is explained in our Trademark Policy .
Thank you for visiting our website and your interest in our services and products. As the protection of your personal data is an important concern for us, please click on the "More information" link to access our Privacy Policy page - which will open in a separate browser tab - where we explain what information we collect during your visit to our website, how it is processed, and whether or how it may be used.
Once you have carefully read our Privacy Policy page, close the browser tab to return to this page and click on the "Save Preferences" button under this text to acknowledge it, close the dialogue and return to the website.
We take all the necessary technical and organisational security measures to protect your personal data from loss and misuse. Your data is stored in a secure operating environment that is not accessible to the public.
