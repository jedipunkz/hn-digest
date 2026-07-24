---
source: "https://rentry.co/organic-code-spec-v1-0-md"
hn_url: "https://news.ycombinator.com/item?id=49039130"
title: "Specification for a 100% AI-Free Software Supply Chain"
article_title: "...:::The Organic Code Movement: Specification for a 100% AI-Free Software Supply Chain:::..."
author: "open-reality"
captured_at: "2026-07-24T18:11:53Z"
capture_tool: "hn-digest"
hn_id: 49039130
score: 2
comments: 0
posted_at: "2026-07-24T17:39:20Z"
tags:
  - hacker-news
  - translated
---

# Specification for a 100% AI-Free Software Supply Chain

- HN: [49039130](https://news.ycombinator.com/item?id=49039130)
- Source: [rentry.co](https://rentry.co/organic-code-spec-v1-0-md)
- Score: 2
- Comments: 0
- Posted: 2026-07-24T17:39:20Z

## Translation

タイトル: 100% AI フリーのソフトウェア サプライ チェーンの仕様
記事のタイトル: ...:::オーガニック コード運動: 100% AI フリー ソフトウェア サプライ チェーンの仕様:::...
説明: 機械時代はソフトウェア技術の魂を盗みました。かつて人間の推論の究極の技術、つまり思考から論理を作り上げる技術であったものは、自動化されたスロップ、統計的推測装置、幻覚的な構文によって氾濫しました。私たちは、感染するツールチェーンによってコンパイルされた、人間が本当に理解できないコードに溺れています...

記事本文:
-->
...:::オーガニック コード運動: 100% AI フリー ソフトウェア サプライ チェーンの仕様:::...
機械時代はソフトウェア技術の魂を奪いました。かつて人間の推論の究極の技術、つまり思考から論理を作り上げる技術であったものは、自動化されたスロップ、統計的推測装置、幻覚的な構文によって氾濫しました。私たちは、マシンに感染したツールチェーンによってコンパイルされ、アルゴリズムによって設計されたハードウェア上で実行される、人間が本当に理解できないコードに溺れています。
私たちは船の引き渡しを拒否します。私たちは、真のソフトウェアは人間の知性、意図、責任の行為であることを宣言します。人間の著作物の途切れることのない連鎖を維持するために、私たちはシリコンに線を引きます。私たちは総合的なシフトを拒否します。
このガイドでは、OS レベルから人間の作成者に至るまで、AI の関与を 100% 排除したコードを記述および構築するための要件の概要を説明します。経験則は次のとおりです。誰も信用しないでください。書かれているコードをすべて見たわけではないので、AI が使用されていないとは断言できません。これは、ツールをコンパイルしたコンパイラ、OS、および関連するすべての知識に当てはまります。このルールは AI の使用にのみ関係していることを思い出してください。コンパイラー/ツールチェーン/OS に悪意のあるバックドアがある可能性があります。AI を排除しているだけで、すべての危険を排除しているわけではありません。
最新の IDE や OS などには、ある時点で AI で記述されたコードが含まれている可能性があるため、古いビルドを厳密に使用し、更新しない必要があります。推奨される OS は、Debian 9 (Stretch) や Ubuntu 16.04.3 LTS など、2017 年以降のものです。アップデートは自分で作成することも、信頼できる人から受け取ることもできます。
AI以前のコードに基づくプロジェクトは、すべてのパッチに、以前の経験や有効な参考文献、または他の形式の検証に基づいた決定を説明する音声解説が伴う、人間が検証したコードベースの維持を開始する可能性がありますが、誰も論理的に信頼できません。

cはそれには適用できません。たとえそれが現れてそれを使用したとしても、コードに AI が含まれていないと断言することはできません。
最新のハードウェアも、シリコンやメモリから、CPU や BIOS で実行されるマイクロコードに至るまで、AI を使用して開発されている可能性があり、それらは 2017 年以前のものであるはずです。
2017 以前のリリースにロックされているエディタのみが許可されます。例:
厳禁: 新しいバージョンの VS Code、VSCodium、Cursor、Windsurf、JetBrains IDE、および Visual Studio 2019/2022。
ツールチェーン バイナリもそのブートストラップ ルーチンも生成モデルを使用してビルドまたはリファクタリングされていないことを保証するには、コンパイラとビルド システムを 2017 以前のバージョンに凍結する必要があります。例:
C/C++: GCC 7.2.0 または LLVM 5.0.0
ビルド システム: CMake 3.9.6 と Ninja 1.8.2
最大限の保証を得るには: ツールチェーンのバイナリとコンパイラは、読み取り専用の物理メディア (アーカイブされたリリース CD/DVD など) から取得するか、インターネット アーカイブ (ウェイバック マシン) 上の検証済みの履歴スナップショットから直接取得する必要があります。ダウンロードされたすべてのアセットは、オリジナルの 2017 年および初期の SHA-256 チェックサムと PGP 署名を使用して検証する必要があります。
許可された参考資料:
物理的な書籍: 2017 年 12 月 31 日より前に印刷および発行された技術マニュアル、教科書、およびドキュメント。
オフライン ドキュメントセット: ローカル システムのマニュアル ページ、ISO/POSIX 標準 PDF (2017 年 12 月 31 日より前)、または 2017 年 12 月 31 日より前に生成されたオフライン ドキュメントセット。
ウェイバック マシンによって確認された、編集されていないコンテンツを含む古いフォーラムの投稿。
スタック オーバーフロー: LLM によって生成されたコード応答のため禁止されています。
最新の Web 検索エンジンとブログ: AI で生成されたチュートリアル、SEO で作られた記事、自動化されたコードの要約のため禁止されています。
ライブ パッケージ マネージャーがない: リモート レジストリに接続する npm install、pip install、cargo build などのコマンドは禁止されています。
販売：すべてサードパー

ty ライブラリは、2017 年 12 月 31 日より前に公開されたバージョンから入手する必要があります。
可能であれば、入手した各ツール、依存関係、およびソースをリリース チェックサムで検証して、信頼性を確認してください。
技術独占企業に合成バベルの塔を建てさせましょう。目に見えないバグでシステムが崩壊するまで、大衆に幻覚のようなロジックを貼り付けさせましょう。
私たちは困難な道を選びます。私たちは工芸品を選びます。コンパイラをフリーズさせ、メディアを物理的に、そして心を鋭く保ちます。チェックサムを検証し、ツールチェーンを保護し、コードのすべての行を人間の思考の記念碑として位置づけます。
機械を信用しないでください。自分の未来を自分で書きましょう(๑•̀ㅂ•́)و✧
上記のコンテンツ警告が設定されたリンクにアクセスしようとしています。続行しますか?

## Original Extract

The machine age has stolen the soul of software craft. What was once the ultimate art of human reasoning—forging logic from thought—has been flooded by automated slop, statistical guessers, and hallucinated syntax. We are drowning in code no human truly understands, compiled by toolchains infecte...

-->
...:::The Organic Code Movement: Specification for a 100% AI-Free Software Supply Chain:::...
The machine age has stolen the soul of software craft. What was once the ultimate art of human reasoning—forging logic from thought—has been flooded by automated slop, statistical guessers, and hallucinated syntax. We are drowning in code no human truly understands, compiled by toolchains infected by machines, running on hardware designed by algorithms.
We refuse to surrender the craft. We declare that true software is an act of human intellect, intent, and accountability. To preserve the unbroken chain of human authorship, we draw a line in the silicon. We reject the synthetic shift.
This guide outlines the requirements to write and build code that is 100% free from AI involvement from the OS level up to the human author. The rule of thumb is: Trust no one, you didn't see all the code being written, so you cannot affirm it didn't use AI! This applies for the compilers your tools were compiled with, your OS, and all the knowledge involved. A reminder that this rule relates to AI usage only, maybe your compiler/toolchain/OS have a malicious backdoor on it, who knows, you are getting rid of AI, not every danger.
Modern IDEs, OSs, and such can contain code that was written with AI at some point, you need to strictly use an older build and do not update. Recommended OSs are from 2017, like Debian 9 (Stretch) or Ubuntu 16.04.3 LTS. You can write the update yourself or receive them from people you trust.
While a project based on Pre-AI code could start to maintain a human verified codebase where all patches are accompanied by audio commentary explaining decisions based on prior experience and valid references, or other form of verification, the trust no one logic cannot be applied to that. If that ever appears and you use it, you cannot affirm your code is AI-free.
Modern hardware also could have been developed with AI, from the silicon and memory, to the microcode running on your CPU and your BIOS, they are should be from 2017 or earlier.
Only editors locked to 2017 and prior releases are permitted, e.g.:
Strictly Prohibited: New versions VS Code, VSCodium, Cursor, Windsurf, JetBrains IDEs, and Visual Studio 2019/2022.
Compilers and build systems must be frozen to 2017 and prior versions to guarantee that neither the toolchain binaries nor their bootstrap routines were built or refactored using generative models, e.g.:
C/C++: GCC 7.2.0 or LLVM 5.0.0
Build Systems: CMake 3.9.6 with Ninja 1.8.2
For Maximum Assurance: Toolchain binaries and compilers should be acquired from read-only physical media (such as archived release CDs/DVDs) or sourced directly from verified historical snapshots on the Internet Archive (Wayback Machine). All downloaded assets must be validated using original 2017 and early SHA-256 checksums and PGP signatures.
Permitted Reference Materials:
Physical Books: Technical manuals, textbooks, and documentation printed and published before December 31, 2017.
Offline Docsets: Local system man pages, ISO/POSIX standard PDFs (pre December 31, 2017), or offline docsets generated prior to December 31, 2017.
Old forum posts with non-edited content confirmed by the Wayback Machine.
Stack Overflow: Prohibited due to LLM-generated code answers.
Modern Web Search Engines & Blogs: Prohibited due to AI-generated tutorials, SEO-farmed articles, and automated code summaries.
No Live Package Managers: Commands like npm install, pip install, or cargo build connecting to remote registries are forbidden.
Vendoring: All third-party libraries must be sourced from versions published prior to December 31, 2017.
If possible, verify each tool, dependency and source you get with their release checksums to be sure of the authenticity.
Let the tech monopolies build their towers of synthetic Babel. Let the masses paste hallucinated logic until their systems crumble under unseen bugs.
We choose the hard road. We choose the craft. Keep your compilers frozen, your media physical, and your mind sharp. Verify your checksums, guard your toolchain, and let every line of code stand as a monument to human thought.
Trust no machine. Author your own future (๑•̀ㅂ•́)و✧
You are about to visit a link which has been flagged with the above content warnings. Do you wish to continue?
