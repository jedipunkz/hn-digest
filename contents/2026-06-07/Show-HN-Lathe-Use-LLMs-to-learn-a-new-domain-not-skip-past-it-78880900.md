---
source: "https://github.com/devenjarvis/lathe"
hn_url: "https://news.ycombinator.com/item?id=48433756"
title: "Show HN: Lathe – Use LLMs to learn a new domain, not skip past it"
article_title: "GitHub - devenjarvis/lathe: Generate hands-on, multi-part technical tutorials on demand, with LLM skills tuned to make content approachable. Then you work through them yourself, by hand ✋ · GitHub"
author: "devenjarvis"
captured_at: "2026-06-07T12:39:48Z"
capture_tool: "hn-digest"
hn_id: 48433756
score: 2
comments: 0
posted_at: "2026-06-07T11:16:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Lathe – Use LLMs to learn a new domain, not skip past it

- HN: [48433756](https://news.ycombinator.com/item?id=48433756)
- Source: [github.com](https://github.com/devenjarvis/lathe)
- Score: 2
- Comments: 0
- Posted: 2026-06-07T11:16:46Z

## Translation

タイトル: Show HN: Lathe – LLM を使用して新しいドメインを学習し、そこをスキップするのではありません
記事のタイトル: GitHub - devenjarvis/lathe: 内容を親しみやすくするために調整された LLM スキルを使用して、実践的な複数パートの技術チュートリアルをオンデマンドで生成します。その後、自分で手作業で作業します ✋ · GitHub
説明: コンテンツを親しみやすいものにするために調整された LLM スキルを使用して、実践的な複数パートの技術チュートリアルをオンデマンドで生成します。その後、自分で手作業で作業します ✋ - devenjarvis/lathe
HN テキスト: こんにちは、HN! Lathe は、私に代わって作業を行うのではなく、LLM を使用して何か新しいことを教える実験です。学習したい技術的なトピックについて、ソースに裏付けされた実践的なチュートリアルを生成します。次に、まさにそのために構築されたローカル UI でコードを手動で読んで入力 (gasp ) することで、自分で作業を進めます。これは、Go CLI と LLM エージェント スキル (クロード コード / カーソル / コーデックス) です。 「/lathe Erlang で 3D スライサーを構築する」のようなプロンプトを表示し、`latheserve` を実行してローカル Web アプリを起動し、それをブラウザーで読み取ります。すべてのチュートリアルには、これまで私にとって独学での学習を快適にしてくれたものが含まれています。 - スクロールすると続く目次
- 考えるきっかけとなる補足
- 読者のための演習
- より深く理解するために使用できるコンテンツをバックアップするソース チュートリアルの背後にある人間の知力の不足を補うために、コンテンツについて質問したり、別の LLM にチュートリアルが実際にコンパイルおよび実行されることを検証してもらったり、別のパートで拡張したりすることもできます (2021 年以来更新されていない「パート 4/6」はもうありません)。私は人間が書いたチュートリアルを置き換えるために旋盤を作ったわけではありません。私は人が書いたチュートリアルが大好きなので旋盤を作りましたが、まだ人が書いた優れたチュートリアルが存在しない技術領域を学びたかったのです (3D の構築)

スライサーを最初から作成し、埋め込み Zig を利用しやすくするなど)。 README には、私が PSP の自作チュートリアルを通じてプログラミングを始めた経緯と、それを LLM に負けてこれを構築するのに十分な悩みを抱えた理由についての長い記事があります。私はあなたに何かを売り込むためにここにいるわけではありません（ここにはVC支援のスタートアップに近いものは何もありません:D）。これは LLM であり、その出力は通常は良好ですが、決して完璧ではありません。これまでのところ、私の経験では、入力しているのはあなたであり、実際に取り組んでいるので、奇妙なものに気づくことができます（そして、それを押し返すこと自体が一種の学習であることがわかりました）。そして、はい、それはバイブコード化されています。なぜなら、それは範囲が狭く、リスクが低く、個人的なかゆみを掻くからです。個人的には Claude Code + macOS で実行しています。他のセットアップも動作するはずですが、まだ検証できていません。人間が書いたものを学ぶためのリソースが見つかった場合は、まずそれを読んでください。しかし、Lathe は、そうでない場合にギャップを埋めるためにここにいます。LLM が私たちの思考を低下させるのではなく、より良くするのに役立つ例として役立つことを願っています。リポジトリ: https://github.com/devenjarvis/lathe チェックしてみようと思ったら、ぜひフィードバックをいただければ幸いです。

記事本文:
GitHub - devenjarvis/lathe: コンテンツを親しみやすくするために調整された LLM スキルを使用して、実践的な複数パートの技術チュートリアルをオンデマンドで生成します。その後、自分で手作業で作業します ✋ · GitHub
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
デベンジャービス
/
旋盤
公共
通知

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
123 コミット 123 コミット .claude/ スキル .claude/ スキル .github/ ワークフロー .github/ ワークフロー cmd cmd docs docs 内部 内部 .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml CLAUDE.md CLAUDE.md ライセンスライセンス README.md README.md go.mod go.mod go.sum go.sum install.sh install.sh magefile.go magefile.go main.go main.go すべてのファイルを表示 リポジトリ ファイルのナビゲーション
自分で考えるのではなく、教えるために LLM を使用する実験。
Lathe は、コンテンツを親しみやすいものにするために調整されたスキルを備えた、実践的な複数のパートからなる技術チュートリアルをオンデマンドで生成します。次に、楽しい学習のためにゼロから構築されたローカル UI を使用して、自分で手作業で作業を進めます。 (私たちが石器時代にやったのと同じように 😎)
任意のプロンプトから実践的な技術チュートリアル (単一パートまたは複数パートのシリーズ) を生成します。
専用のローカル UI でチュートリアルを自分で進めてください
スキルを使用して質問し、チュートリアルを確認し、新しい部分でチュートリアルを拡張します
ライブラリからチュートリアルを検索、フィルタリング、管理します
すべてのチュートリアルには、そのソース、使用されたモデル、およびチュートリアルの「音声」を駆動したプロンプトが文書化されています。
Lathe は、LLM スキルと、生成されたチュートリアルの保存、管理、表示に使用される Golang CLI を組み合わせたものです。インストール後 (下記)、次のようなプロンプトを表示することで、LLM セッション (クロード コード、カーソル、およびコーデックスがサポートされている) 内でチュートリアルを生成できます。
/lathe Erlang で 3D スライサーを構築します
次に、任意の端末から旋盤を開きます。
latheserve # Webサーバーを起動し、ブラウザを開きます
心配しないでください。ダークモードもあります。
読みたいチュートリアルをクリックして学習を始めてください。
CLIにはたくさんのoがあります

他のコマンドもありますが、正直なところ、これらは LLM にチュートリアルを管理するための決定的な方法を提供するために構築されました。日常的に必要なものは上記だけで十分だと思います (私がこれまでに使用したものはこれだけです)。チュートリアルについて質問する場合、LLM に検証してもらう場合、または追加パーツで拡張する場合、UI にはこれらのそれぞれに対するアフォーダンスがあり、アクションをトリガーするために LLM に与える正確なスキル/プロンプトが提供されます。
Lathe は単一の自己完結型バイナリです。必要なのは $PATH を旋盤するだけです。の
スキルは、インタラクティブなクロード コード、カーソル、またはコーデックス セッションで実行されます。
自作 (macOS、推奨):
brew install devenjarvis/tap/lathe
カスク (ビルド済みバイナリ) として配布されるため、macOS のみです。Linux では、
スクリプトをインストールするか、以下のインストールに進みます。
カール -sSf https://raw.githubusercontent.com/devenjarvis/lathe/main/install.sh |しー
Go (Go 1.25 以降が必要):
github.com/devenjarvis/lathe@latest をインストールしてください
ソースより:
git clone https://github.com/devenjarvis/lathe
CD旋盤
旋盤を作りに行く
スキルをインストールする
スキルはバイナリにバンドルされています。旋盤を取り付けた後、それらを
プロジェクトを作成して、Claude Code (または Cursor / Codex) がそれらを検出できるようにします。
旋盤スキルのインストール # ./.claude/skills/<名前>/SKILL.md (このプロジェクト)
旋盤スキルのインストール --user # ~/.claude/skills/<名前>/SKILL.md (すべてのプロジェクト)
旋盤スキルのインストール --agent カーソル # ./.cursor/commands/<slug>.md (カーソル スラッシュ コマンド)
旋盤スキルのインストール --agent codex # ./.agents/skills/<名前>/SKILL.md (コーデックス エージェント スキル)
旋盤スキル install --agent all # クロード コード、カーソル、およびコーデックス
旋盤スキル一覧 # 同梱スキルを表示
Codex は Claude Code と同じ SKILL.md 形式を使用しているため、スキルはそのまま出荷されます。
(そして --user は ~/.agents/skills/... にインストールします)。カーソルコマンドはスラッシュで呼び出されます
/<slug> として (例: /lathe );インタラクティブ

Claude については、有効なハンドオフ モデルが文書化されています。
コードなので、Cursor と Codex では実行時の詳細がいくつか異なります。
私は 2000 年代に 10 代のころ、PSP (PlayStation Portable) 用の自作ゲームを Lua で作成し、その後 C++ でプログラミングの方法を学びました。当時私が学んだことの多くは、小さな PSP 自作コミュニティを通じてのもので、参加できたことに非常に感謝していますが、その形成的な学習の多くは、インターネット上で利用できる無料のオンライン リソースとチュートリアルのおかげでもあります (2007 年の cplusplus.com に感謝します - あのサイトには、以前よりもはるかに多くの広告が表示されています 😅)。最終的に私はプロのソフトウェア エンジニアになり、その後 10 年間、豊富な技術ブログを見つけて利用することで「スキルアップ」に費やしました (ただし、通常は、必要以上に興味深いトピックを学ぶため)。そして、私の学習スタイルにとってより重要なことは、実践的なチュートリアルです。 build-your-own-x リポジトリ 、 Crafting Interpreters などのリソース、および raytracer の構築から timeseries データベース、線形代数行列ライブラリ、およびその間のすべてを教えてくれる 1,000 の他の 1 回限りのチュートリアル (真剣に言うと、私に影響を与えた素晴らしい実践的なチュートリアルをすべてリストアップすることさえできません)。
実践的に学習することが、私にとって常に最善の学習方法です。これらのチュートリアルは、まったく新しい領域でゼロから 1 へ進むために必要な学習曲線を私に与えてくれましたが、さらに重要なことに、これらのチュートリアルは私に 1 から 2、10 まで独力で進むための足場と自信を与えてくれました。
2026 年に早送りすると、今では LLM が存在します。 LLM との複雑な関係について話が逸れるつもりはありませんが、ソフトウェアを作成する場合、LLM は興味深いものであり、多くの場合、非常に生産的です。しかし、彼らはあなたの代わりにほとんどの仕事をしてくれるので、その仕事がなくなると、私が助けてくれた部分も奪われることになります。

新しい概念や領域を確立します。場合によっては、それは問題ではありません - 私たちには出荷する製品があり、LLM がそれをより早く出荷するのに役立ちます - しかし、私にとって、そしてこの分野と趣味の喜びにとって、私はまだ「ああ、そうだ！」を切望しています。何かがついにカチッと音を立てて、それを自分のものに形作り始めるのに必要な自信が持てる瞬間です。
つまり、旋盤は、私に代わって考えるのではなく、LLM を使用して私に教える実験です。この仕事を愛することを私に教えてくれた実践的な学習の瞬間を再現し、理論的には何でも教えてくれる広範な「専門家」LLM の可能性と結びつけます。私は旋盤をきっかけとして、どうやって始めればよいか分からないプロジェクトに取り組むきっかけとして使用していますが、人が教えるための既存の資料が見つかりません。たとえば、私が最初に旋盤を思いついたのは、3D スライサー ソフトウェアを最初から書きたかったからです (G コードに関するドキュメントを見つけるだけでも大変でした。再ラップするように叫びました)。これを書いている時点では、私は Zig を使用して組み込みソフトウェア開発の世界に飛び込んでいます。どちらの場合も、旋盤は、人が書いたリソースがまだ存在しない、無名または非常に若い領域でゼロから 1 を作るのに効果的なツールでした (そして、LLM だけが読むのであれば、いつまで人間がわざわざチュートリアルを書くつもりなのかと思います...)。
しかし、幻覚についてはどうでしょうか？
旋盤のチュートリアルは人間が書いたものと同じくらい優れていますか?決してそうではありません。しかし、彼らに欠けている心、人格、アーキテクチャの健全性は、チュートリアルの作成者を準備してあなたのすべての質問に答えるのを待ち、あなたが望んでいたものではないときはいつでも喜んでチュートリアルを修正したり更新したりすることで補っており、実際に 2018 年に開始したシリーズの全 6 部を書き終えています (私たち全員がそこにいたことがあります 😁)。 Lathe は LLM であり、私はこれと同じくらい優れたものになるように構築および調整しましたが、

私はこの特定のタスクを成功させる方法を知っていますが、それでも LLM が失敗するのと同じように失敗するでしょう。これらのタスクは、プログラミング時に最適化する可能性のある反復的な機械的実行に関するものではなく、具体的なコンセプトを最初から最後まで調査、設計、説明することに重点が置かれているため、アクセスできる最大の「思考」モデル (Opus、GPT-5 Codex など) を使用することをお勧めします。
さらに、この状況での幻覚のリスクは、私の意見では大幅に低くなります。 Lathe は、ユーザーの思考を支援するために構築されており、ユーザーがこのコードを自分で入力するという期待に基づいて構築されています。ガイドを読んで入力することで、積極的に作業に取り組んでいることになり、自然に「ちょっと待って、それは意味があるの?」と尋ねることができるはずです。何か変なものに出会ったとき。その時点で、/lathe-ask (場合によっては、外部ドメインであるため、LLM が私にはなかった適切な推論を返してきて、何かを学ぶことができます) するか、LLM にチュートリアルを更新するように直接指示することができます。私にはこれを裏付ける教育学的な資格はありませんが、LLM の間違いを見つけて押し返すことで、実際には概念をよりよく内面化できているのかもしれないと思います。 YMMV。
そうは言っても、人間が書いたチュートリアルがあれば、私はいつもまずそこに手を伸ばすでしょう。あなたがそうしないよりも頻繁にそうすることを願っています。しかし、私と同じように学習し、教材をあまり使わない分野に取り組みたいのであれば、旋盤は非常に優れたツールです。それが人間ではなく LLM であることを覚えておいてください。これを助けるために、私はあなたが何を得ることができ、何が得られないのかを常に明確にするように努めています。チュートリアルを書くための旋盤のスキルは、書いた内容について確信が持てないときに教えてくれます。私はより「個人的な」声を提供していますが、デフォルトでは、暗いふりをしない声を選んでいます。

それは違います。
正直に言うと、これはあなたがバイブコーディングしたのですか？それはあなたの論文と矛盾していませんか？
そう、旋盤は「バイブコーディング」されています。この場合、旋盤の範囲とリスクは低くなります。これは個人的な学習のための生きた論文です。とはいえ、私は最近これを毎日使用しており、ツールボックスにある便利で安定したツールであることが証明されています。私はこれを使用することで多くのことを学んでおり、現時点では、他の人もそれから恩恵を受けることができるのに十分だと考えています。今後の数ポイントのリリースでは、他のユーザーにとっても安定した状態を維持できるように、意図的なコード/アーキテクチャのクリーンアップが行われることを期待しています。もちろん、私が得たフィードバックはすべて反映されます。
とはいえ、透明性を保つために、今日は私自身のユースケースで旋盤をテストします - MacOS でクロード コードを使用します。そのセットアップの外にいる場合、旋盤は動作するはずですが、私はそれを検証していません。別のセットアップで試してみてうまく機能する場合、または最終的に道にぶつかることになる場合は、どちらにせよお知らせいただけると幸いです。
さて、それはどのように機能するのでしょうか？
LLM スキル — チュートリアルの生成と操作。すべて対話型 LLM セッションで実行されます。 /lathe は part-01.md を書き込み、 /lathe-extend は次のパートを追加します。 /lathe-verify はチュートリアルを通じて動作し、コンパイルして実行されることを確認します。 /lathe-ask は Questi に答えます。

[切り捨てられた]

## Original Extract

Generate hands-on, multi-part technical tutorials on demand, with LLM skills tuned to make content approachable. Then you work through them yourself, by hand ✋ - devenjarvis/lathe

Hey HN! Lathe is an experiment in using LLMs to teach me something new, instead of doing the work for me. It generates a hands-on, source-backed tutorial for any technical topic you want to learn. Then you work through it yourself by reading and typing the code by hand ( gasp ) in a local UI built for exactly that. It's a Go CLI plus LLM agent skills (Claude Code / Cursor / Codex). You prompt something like "/lathe build a 3D slicer in Erlang", run `lathe serve` to spin up a local webapp, and read it in your browser. Every tutorial comes with the things that have made self-learning a pleasant experience for me in the past: - table of contents that follows along as you scroll
- side-notes that nudge you to think
- exercises for the reader
- sources backing up the content that you can use to take you deeper To help make up for the lack of human brainpower behind the tutorial, you can also ask questions about the content, have another LLM verify the tutorial actually compiles and runs, or extend it with another part (no more "Part 4 of 6" that hasn't seen an update since 2021). I didn't build lathe to replace human-written tutorials. I built lathe because I _love_ human-written tutorials, but wanted to learn technical domains where no good human-written tutorial exists yet (building a 3D slicer from scratch, making embedded Zig approachable, etc). There's a longer story in the README about how I got started with programming through PSP homebrew tutorials, and why losing that to LLMs bugged me enough to build this. I'm not here to sell you anything (there's nothing close to a VC-backed startup here :D). It's an LLM, and its output is usually good but not perfect by any means. So far, my experience is that because you're the one typing and actually engaged, you catch the weird stuff (and I'm finding that pushing back on it is its own kind of learning). And yes, it's vibecoded, because it's low scope, low risk, and scratching a personal itch. I run it on Claude Code + macOS personally, other setups should work but I haven't been able to verify them yet. If you can find resources to learn something that was written by a human, read that first. But Lathe is here to fill in the gaps when that isn't the case, and I hope it serves as an example where LLMs can help us think better, rather than less. Repo: https://github.com/devenjarvis/lathe Would love your feedback if you decide to check it out!

GitHub - devenjarvis/lathe: Generate hands-on, multi-part technical tutorials on demand, with LLM skills tuned to make content approachable. Then you work through them yourself, by hand ✋ · GitHub
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
devenjarvis
/
lathe
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
123 Commits 123 Commits .claude/ skills .claude/ skills .github/ workflows .github/ workflows cmd cmd docs docs internal internal .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum install.sh install.sh magefile.go magefile.go main.go main.go View all files Repository files navigation
An experiment in using LLMs to teach you, rather than think for you.
Lathe generates hands-on, multi-part technical tutorials on demand, with skills tuned to make content approachable. Then you work through them yourself, by hand, in a local UI built from the ground up for pleasant learning. (Just like we did it in the stone age 😎)
Generate hands-on technical tutorials (single-part or a multi-part series) from any prompt
Work through the tutorial yourself in a purpose-built local UI
Use skills to ask questions, verify the tutorial, and extend it with a new part
Search, filter, and manage tutorials from your library
Every tutorial documents its sources, which model was used, and what prompt drove the "voice" for the tutorial
Lathe is a combination of LLM skills and a Golang CLI used to store, manage, and view generated tutorials. After install (below), you can generate a tutorial inside any LLM session (Claude Code, Cursor, and Codex supported) by prompting something like:
/lathe build a 3D Slicer in Erlang
Then open lathe from any terminal:
lathe serve # starts the web server, opens the browser
Don't worry, we also have dark mode:
Click the tutorial you want to read and start learning!
The CLI has a bunch of other commands, but honestly those were built to give the LLM a deterministic way to manage tutorials. I expect the above to be all you need (it's all I ever use) for day-to-day. If you want to ask a question about a tutorial, have the LLM verify it, or extend it with an additional part, the UI has affordances for each of these which will give you the exact skill/prompt to give your LLM in order to trigger the action.
Lathe is a single self-contained binary. All you need is lathe on your $PATH ; the
skills run in an interactive Claude Code, Cursor, or Codex session.
Homebrew (macOS, recommended):
brew install devenjarvis/tap/lathe
Distributed as a cask (a pre-built binary), so it's macOS-only — on Linux use the
install script or go install below.
curl -sSf https://raw.githubusercontent.com/devenjarvis/lathe/main/install.sh | sh
Go (needs Go 1.25+):
go install github.com/devenjarvis/lathe@latest
From source:
git clone https://github.com/devenjarvis/lathe
cd lathe
go build -o lathe
Install the skills
The skills are bundled into the binary. After installing lathe , drop them into a
project so Claude Code (or Cursor / Codex) can discover them:
lathe skills install # ./.claude/skills/<name>/SKILL.md (this project)
lathe skills install --user # ~/.claude/skills/<name>/SKILL.md (all projects)
lathe skills install --agent cursor # ./.cursor/commands/<slug>.md (Cursor slash commands)
lathe skills install --agent codex # ./.agents/skills/<name>/SKILL.md (Codex Agent Skills)
lathe skills install --agent all # Claude Code, Cursor, and Codex
lathe skills list # show the bundled skills
Codex uses the same SKILL.md format as Claude Code, so its skills ship verbatim
(and --user installs to ~/.agents/skills/... ). Cursor commands are slash-invoked
as /<slug> (e.g. /lathe ); the interactive handoff model is documented for Claude
Code, so a few runtime details differ on Cursor and Codex.
I learned how to program as a teen in the 2000s by building homebrew games for my PSP (PlayStation Portable) in Lua, and then in C++. Lots of what I learned at the time was through the small PSP homebrew community I'm incredibly grateful I got to be a part of, but I also owe much of that formative learning to free online resources and tutorials available on the internet (shoutout to 2007 cplusplus.com - man does that site have a lot more ads now than it used to 😅). Eventually I became a professional software engineer and I spent the next decade "upskilling" (though usually to learn more interesting topics than needed for ) by finding and consuming a wealth of technical blogs, and more importantly for my learning style - hands on tutorials. Resources like the build-your-own-x repo , and Crafting Interpreters , and the 1,000 other one-off tutorials that taught me everything from building a raytracer , to a timeseries database , to a linear algebra matrix library and everything in between (seriously, I couldn't even begin to list all the amazing hands-on tutorials out there that have influenced me).
Hands on learning is how I've always learned best. These tutorials gave me the learning curve I needed to go from zero-to-one in a brand new domain, but even more importantly they gave me footing and confidence to take it from one-to-two-to-ten on my own.
Fast forward to 2026, and now we've got LLMs. I'm not going to go off topic about my complicated relationship with LLMs, but for writing software they are interesting and in many cases they can be really productive! But they do most of the work for you, and with that work gone they also take away the part that helped me learn a new concept or domain. In some cases, that doesn't matter - we've got a product to ship and LLMs help us ship it faster - but for me and my joy in this field and hobby I still crave those "ah ha!" moments where something finally clicks and I have the confidence I need to begin shaping it into my own.
So lathe is an experiment in using LLMs to teach me, rather than think for me. To recreate those moments of hands-on learning that taught me to love this work, and marry it with the potential of a broad "expert" LLM who can, in theory, teach me anything. I use lathe as a catalyst to get me started on projects I wouldn't know how to start in, and can't find any existing human written resources to teach . For example I first came up with lathe because I wanted to write a 3D Slicer Software from scratch (just finding documentation on g-code was a pain, shoutout to reprap ). At the time of writing I'm diving into the world of embedded software development with Zig. Both of these cases lathe has been an effective tool in getting me from zero-to-one in obscure or extremely young domains where the human written resources just don't exist yet (and I wonder for how long humans will still bother writing tutorials if only the LLMs read them...).
But what about hallucinations?
Are lathe tutorials as good as ones written by humans? Not in the slightest. But what they lack in heart, personality, and architectural soundness, they make up for by having the tutorial writer ready and waiting to answer all of your questions, always willing to fix or update their tutorial when it isn't exactly what you wanted, and they actually complete writing all 6 parts to that series they started in 2018 (we've ALL been there 😁). Lathe is an LLM, and while I've built and tuned it to be as good as I know how to make it for this particular task, it's still going to fail in the ways LLMs fail. I recommend using the biggest "thinking" model you have access to (Opus, GPT-5 Codex, etc) as these tasks are less about iterative mechanical execution you might optimize for when programming, and more about researching, designing, and explaining a tangible concept from start to finish.
Additionally, the risk for hallucinations in this context is, in my opinion, significantly lower. Lathe is built to help you do the thinking, and is built around the expectation that you're the one typing this code out yourself. By reading through the guide and typing it out, you are actively engaged in the work and should be well positioned to naturally ask "wait, does that make sense?" when you come across something weird. At which point you can /lathe-ask (and sometimes the LLM comes back with good reasoning I didn't have because it's a foreign domain, and I learn something) or just straight tell your LLM to update the tutorial. While I have no pedagogical credentials to back this up, I think I may be actually internalizing concepts better by catching and pushing back on perceived slip-ups of the LLM. YMMV.
All of that said, if you can find a tutorial written by a human, I'd always reach for that first. I hope more often than not you do. But if you learn the same way I do and want to dive into a domain that is light on teaching materials, lathe is a pretty cool tool. Just remember it is an LLM and not a human. To help with this, I try to make it clear at all times what you are and are not getting. The lathe skills to write tutorials will tell you when it isn't sure about something it has written, and while I offer a more "personal" voice, I've defaulted to one that doesn't pretend to be something it isn't.
Be honest, did you vibecode this? Isn't that contradictory to your thesis?
Yep, lathe is "vibecoded". In this case, the scope and risk of lathe is low. It's a living thesis, for personal learning. That said, I've been using it daily lately and it's proven to be a useful and stable tool in my toolbox. I'm learning a lot by using it, and at this point I think it's good enough that others might benefit from it too. I expect the next few point releases to be some intentional code/architecture clean up to ensure it remains stable for others, and of course incorporate any feedback I get.
That said, for the sake of transparency, today I test lathe for my own usecases - Using Claude Code on MacOS. If you are outside of that setup, lathe should work, but I've not verified it. If you're willing to try it on a different setup and it does work, or you end up hitting a bump in the road, I'd love an issue letting me know either way!
Alright then, how does it work?
LLM skills — generate and work with tutorials, all run in your interactive LLM session: /lathe writes part-01.md , /lathe-extend adds the next part, /lathe-verify works through a tutorial to confirm it compiles and runs, /lathe-ask answers questi

[truncated]
