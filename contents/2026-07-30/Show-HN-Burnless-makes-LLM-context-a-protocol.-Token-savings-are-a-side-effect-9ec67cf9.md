---
source: "https://github.com/Rudekwydra/burnless"
hn_url: "https://news.ycombinator.com/item?id=49112480"
title: "Show HN: Burnless makes LLM context a protocol. Token savings are a side effect"
article_title: "GitHub - Rudekwydra/burnless: Stop replaying the transcript: capsule session state + rolling memory for LLM CLIs (Claude Code, Codex). Measured: −90.3% vs no-cache / −30% vs cached replay at turn 10, reproducible bench. MIT, provider-agnostic. · GitHub"
author: "rudekwydra"
captured_at: "2026-07-30T17:16:45Z"
capture_tool: "hn-digest"
hn_id: 49112480
score: 1
comments: 1
posted_at: "2026-07-30T16:45:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Burnless makes LLM context a protocol. Token savings are a side effect

- HN: [49112480](https://news.ycombinator.com/item?id=49112480)
- Source: [github.com](https://github.com/Rudekwydra/burnless)
- Score: 1
- Comments: 1
- Posted: 2026-07-30T16:45:01Z

## Translation

タイトル: Show HN: Burnless は LLM コンテキストをプロトコルにします。トークンの節約は副作用です
記事のタイトル: GitHub - Rudekwydra/burnless: トランスクリプトの再生を停止: LLM CLI のカプセル セッション状態 + ローリング メモリ (クロード コード、コーデックス)。測定値: ターン 10 で、キャッシュなしの場合に対して -90.3% / キャッシュされたリプレイに対して -30%、再現可能なベンチ。 MIT、プロバイダーに依存しない。 · GitHub
説明: トランスクリプトの再生を停止します: カプセル セッション状態 + LLM CLI のローリング メモリ (クロード コード、コーデックス)。測定値: ターン 10 で、キャッシュなしの場合に対して -90.3% / キャッシュされたリプレイに対して -30%、再現可能なベンチ。 MIT、プロバイダーに依存しない。 - ルデクウィドラ/バーンレス
HN テキスト: Burnless を作成したのは、そこに存在する必要のないものにコンテキスト ウィンドウを使用していることに気づいたからです。すでに起こって次のステップには必要なくなったものでトークンを書き込むのをやめたかったのです。それは後になって発見されたものではありません。 Burnless は、物を分離するために生まれました。私がまだ知らなかったのは、問題の大きさ、利益の大きさ、そしてそれが他の皆がたどってきた道とどれほど違うのかということでした。最近私が目にしているのは、誰もがコンテキスト ウィンドウを大きくしたいと考えているということです。もちろんそれは役に立ちます。しかし私の意見では、その概念は間違っています。ウィンドウが大きくなれば、その中にさらに多くのものを入れることができるだけですが、会話、記憶、意思決定、実行、未解決の質問、仕事の履歴、そしてすでに行われたことのすべての古い記録に同じ場所を使用しているという事実は解決されません。 Burnless はそれを層に分離します。会話、考え、議論、質問、作業履歴はインデックス付けされ、ディスクに保存されます。メイン コンテキストには、作業を継続するために引き続き必要な決定事項と情報のみが含まれます。しかし、必要に応じて実際の言葉はまだディスク内にあります

チェックする必要があります（正直に言うと、その必要さえありません）。また、私たちが何かを言い、誰かが繰り返すように頼むとき、私たちは決して同じ言葉を使うことはありません。なぜ AI がそんなことをする必要があるのでしょうか? (元のファイルはまだディスク内にあるので、心配しないでください…) タスクは独立した軽量ワーカーに委任されます。彼らはクリーンなタスクを受け取り、コードを書き、ファイルを操作し、コマンドを実行し、実行した内容をディスクに記録します。そして、はい、ワーカーは同じバイト キャッシュ ヘッダーを使用します (これについても心配する必要はありません) その後、Burnless が結果を監査します。作業者がファイルを作成または変更したと主張した場合、Burnless はタスクを完了としてマークする前に、そのファイルが実際に存在することを確認します。つまり、コンテキストはすべてが存在しなければならない場所ではなくなります。今注目すべきものだけを置いておく場所になります。だからこそ、私は Burnless をプロトコルと比較し、最初は TCP/IP とさえ比較しました。それは同じ歴史の側面を持っているからではなく、動きが似ているからです。つまり、責任を層ごとに、労働者の能力ごとに、そしてこの生きているシステム間のコミュニケーションを定義する便利でクリーンな方法に分離しているからです。会話中に記憶が持ち込まれなくなります。実行はメイン モデルの役割ではなくなります。プロジェクト履歴は、作品のメモリを永続的に占有することはなくなります。 /clear コマンドの後やモデル間の切り替え後でも、プロジェクトは続行できます。トークンの貯蓄が最初の動機でしたが、最終的にはアーキテクチャの結果となりました。実際のタスクでは、ワーカーは約 140 万個のトークンを処理しました。しかし、完全なタスクを完了した後、脳のアクティブなコンテキストには 1,590 個のトークンしかありませんでした。私は、Burnless がまったく同じ 140 万トークンのテキストを取得して 1,590 トークンに圧縮したとは主張しません。実際に何が起こったのかというと、その作業はすべて、

1 つは、アクティブなコンテキスト内に蓄積された履歴を残さずに記録、検証、継続されるものです。最終的な稼働率は約 908 対 1 でした。(そうです、908 倍) この数字を見せると、信じられないように思えます。そのとき、問題は当初想像していたよりもはるかに大きいことに気づきました。一般的な方向は、ますます大きなコンテキスト ウィンドウを構築することですが、私は逆の方向に進み、現時点で考慮する必要のないものをすべてコンテキストから削除しました。このリポジトリには、完全なアーキテクチャ、ベンチマーク、計算、制限事項、数値を再現するために必要なコマンドを含む llms.txt ファイルも含まれています。短いコメントですべてを適切に説明することはできないため、より深いレベルで何が起こっているのかを理解し、主張を検証したい人のために、llms.txt ファイルが用意されています。これを使用すると、AI がその内容を簡単に理解し、質問に応答できるようになります。

記事本文:
GitHub - Rudekwydra/burnless: トランスクリプトの再生を停止します: カプセル セッション状態 + LLM CLI のローリング メモリ (クロード コード、コーデックス)。測定値: ターン 10 で、キャッシュなしの場合に対して -90.3% / キャッシュされたリプレイに対して -30%、再現可能なベンチ。 MIT、プロバイダーに依存しない。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
別のタブまたはウィンドウでアカウントを切り替えました。リロード先

セッションをリフレッシュしてください。
アラートを閉じる
{{ メッセージ }}
ルデクウィドラ
/
燃えない
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
639 コミット 639 コミット .claude .claude .codex/ コマンド .codex/ コマンド .github .github _design _design bench ドキュメント docs サンプル サンプル スクリプト スクリプト サイト サイト src/ burnless src/ burnless テンプレート テンプレート テスト テスト .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md BENCHMARK_T5_FINDINGS.md BENCHMARK_T5_FINDINGS.md BURNLESS_FOR_LLMS.md BURNLESS_FOR_LLMS.md CHANGELOG.md CHANGELOG.md COUNTDOWN_DISCIPLINE.md COUNTDOWN_DISCIPLINE.md GEMINI.md GEMINI.mdライセンス ライセンス MATH.md MATH.md PAPER_2026-05-20.md PAPER_2026-05-20.md PLUGIN_PROTOCOL.md PLUGIN_PROTOCOL.md PROTOCOL.md PROTOCOL.md README.md README.md RELEASE.md RELEASE.md SECURITY.md SECURITY.md SECURITY_AUDIT.md SECURITY_AUDIT.md SPEC.md SPEC.md VISION.md VISION.md llms.txt llms.txt pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
カプセルベースのセッション状態、プレフィックス キャッシュの再利用、およびファイルシステム優先監査を備えた多層 LLM オーケストレーション レイヤー。マサチューセッツ工科大学
AI アシスタントとレビュー担当者: llms.txt には、コンパクトな形式で技術概要が記載されています。
pip インストール バーンレス
cd <あなたのプロジェクト>
burnless setup # CLI/キーを検出し、.burnless/config.yaml を書き込みます
バーンレスドクター # ヘルスチェック — 配線が正しい場合は緑色で終了
次に、実際のセッションを 1 つ実行し、 /clear を実行します。セッションがスレッドとともに戻ってきます
トランスクリプトから再生されるのではなく、ディスクから復元されたままです。
Python 3.10以降。 MCP サーバーのサポートは追加です: pip install 'burnless[mcp]' 。
Codex のセットアップ、ソースからのインストール、アンインストールは [インストール] の下にあります。
AI AG を実行する

重いループに陥ると、通常は 2 つの悪のどちらかを選択することになります。
トランスクリプトを再生します。毎ターン、それ以前のすべてを再送信します。
キャッシュなしループ。支払った料金の最大 90% はモデル自体の再読み取りです。
そして端末は決して読むことのない出力の壁に埋もれてしまいます。
モデルを簡潔にプロンプ​​トハックします。 「言葉を多く言わない」ルールと
積極的な要約はモデルの動作を壊すことでトークンを節約します
コンテキスト: 要約には、次のステップで必要な技術的な詳細が正確に記載されています。
Burnless は、行動面ではなく構造面でその選択を拒否します。脳
コンパクト カプセルを読み取ります (ディスク上の .jsonl、1 ターンあたり最大 80 文字)。労働者が実行する
単独で、完全な出力をディスクに書き込み、1 行の結果をレポートして終了します。
UI スパムはありません。ワーカーがディスクにストリーミングすると、画面にワンライナーが表示されます。
即時ハッキングはありません。 「簡潔に」という呪文は不要です - 冗長性は自由です
コンテキスト内ではなくディスク上に着陸します。
測定されたものであり、約束されたものではありません。キャッシュなしのリプレイに対して -90.3%、キャッシュなしのリプレイに対して -30%
すでにキャッシュされたベースライン (実際の支出の $5.76、 bench/run.py )。残りの部分
この README の数字には、逸話としてラベルが付けられています。
このプロジェクトの履歴に関するメモ (2026-05-08)。 Burnless は 2026 年 5 月 3 日に初めて PyPI に公開され、プロジェクトの新規性と節約を誇張するドキュメントが含まれていました。具体的には、TCP/IP との類似性は、アーキテクチャ上の同等性を示唆しています (実際にはそうではありません)。 「16 倍​​安い」という数字は、普遍的な主張として提示された個人のワークロードの逸話でした。また、プレフィックス キャッシュがモデル間で共有されるという主張は技術的に間違っていました。Anthropic のプレフィックス キャッシュはモデルごとにキー付けされており、共有されません。これらの主張は Claude と共同で書かれたものであり (git log の Co-Authored-By: trailers に表示されます)、今では調整されたアソシというよりも RLHF によって引き起こされた熱意と認識しています。

メッセージ。領収書: git log --pretty=fuller は、インフレ期間 (2026-05-03 から 2026-05-05) と 2026-05-08 の再調整を示します。この修正は 0.7.3 で出荷され、それ以降のすべてのリリースで同じ規律が維持されています。履歴はそのまま残されます。書き直しやカバーはありません。以下のアーキテクチャは、防御可能な実装の選択肢の 1 つであり、基本的なプロトコルの画期的な進歩ではありません。
2026 年 7 月 27 日更新。上記の撤回は正しい。効果が実際にあったとしてもフレーミングは間違っていた。それ以来 2 か月で何が変わったかというと、取り下げられた請求が、測定された再現可能な形式で存在するようになりました。 1 億 2,100 万個以上のトークンがコンテキストから切り離され、作成者のマシン上の追加専用 JSONL にターンバイターンで記録されます。実際の有料 API は、キャッシュなしのリプレイに対して -90.3%、キャッシュ済みのベースラインに対して -30% で実行されます (実際の支出の $5.76、 bench/run.py )。本番環境の委任は、200k ～ 365k トークンの作業を定期的に圧縮して、最大 300 トークン カプセル (700 ～ 1050 倍) に達します。メイ氏の主張が撤回されたのは、効果が現実的ではなかったからではなく、領収書なしで普遍的であると記載されていたためだ。領収書は現在存在しています。それが違いであり、その違いを維持することがこのプロジェクトの全体的な規律です。
Burnless は、Python でのリファレンス実装を備えたプロトコル仕様です。コントラクト (ワーカー エンベロープ JSON、カプセル形式、層セマンティクス、監査ゲート、プラグイン フック) は PROTOCOL.md に書き込まれ、このコードベースから独立しています。バーンレス CLI はその実装の 1 つです。 (LSP の意味でのプロトコル — 他の人が実装できる指定されたインターフェイス — TCP/IP スケールの基盤の主張ではありません。上記の注を参照してください。) 2 番目の独立した実装と準拠スイートは v1.0 の基準であり、出荷された事実ではありません。
リファレンス実装は、AI アシスタント (または独自のコード) と MOD の間に配置されます。

プロバイダー。具体的には次の 3 つのことを行います。
タスクをモデル層 (ゴールド/シルバー/ブロンズ) にルーティングします。層はハードコードされたモデルではなくコマンドであり、任意の CLI を介した任意のプロバイダーです。これらは出荷時のデフォルトから開始され、チャットごとに、またはバーンレス メニュー/バーンレス モデル セットを介してグローバルに変更されます (プロジェクトには独自の階層マップがありません)。
毎ターン完全なトランスクリプトを再生するのではなく、セッション状態をコンパクト カプセルとしてディスク ( .burnless/ ) に保存し、システム プロンプト プレフィックスをバイト同一に保つため、プロバイダーのプロンプト キャッシュはウォームな状態を保ちます。
ファイルシステム (QTP-A) に対してワーカーの出力を監査します。ワーカーがファイルを書き込んだと言う場合、Burnless は成功を報告する前にファイルの存在とサイズが一致していることを確認します。
それが製品全体です。この README のその他の内容はすべて、構成、例、および著者自身の使用法による正直な測定値です。
理論的には新しい画期的な進歩ではありません。階層ルーティング、プロンプト キャッシュ、および状態の要約はすべて、他のツール (LangGraph、AutoGen、CrewAI、Aider など) に存在します。 Burnless の貢献は、小さな CLI としてパッケージ化されたカプセル + ファイルシステム監査 + プラグイン プロトコルという特定の実装の選択です。
魔法コストエリミネーターではありません。すべてのワークロードの漸近形状は変わりません。費用を節約できるかどうかは、セッションの長さ、モデルの組み合わせ、および既存のセットアップがすでにどれだけ積極的にキャッシュしているかによって決まります。
すべての代替案に対してベンチマークを行っているわけではありません。以下の数値は、特定の単純なベースライン (全履歴再生、キャッシュなし) および作成者自身の個人的なワークロードに対して測定されています。それらを普遍的な主張としてではなく、「私が観察したこと」として扱います。
履歴を圧縮するとニュアンスが失われませんか?
これが標準的な反対意見であり、デザインは構造的にそれに答えています。純粋な表面、意味論的履歴 — 人間の記憶の仕組みに基づいてモデル化されています (sa 以降)。

一度何かを言うと、そのアイデアを繰り返しますが、正確な言葉は決して出ません）。モデルは常に現在のプロンプトを完全に受け取ります。人間は常に応答を完全に受け取ります。運ばれた履歴だけが意味論的です。そして何も破壊されません。完全なトランスクリプトはディスク上に残り、インデックスが付けられ、burnless read/log/capsule <id> で逐語的に取得可能になります。 Ref はポインタであり、再生されるペイロードではありません。ニュアンスを逃す最悪のケースは、あと 1 回の検索です。人間の記憶と同じことを、より良いインデックスを使用して行うことができます。
長い複数ターンのセッションでは、成長するトランスクリプトをターンごとに再生する必要がありますが、カプセル + ホット プレフィックス キャッシュにより、入力トークンが大幅に削減されます。著者の日常では、これにより、数日にわたるワークロードにわたる API 支出が大幅に削減されました。走行距離はさまざまです。実際にどのような条件で測定されたかについては、以下の「数値」セクションを参照してください。
セッションが短い場合 ( N ≤ 3 ターン )、ワンショット スクリプト、またはキャッシュと状態を処理するフレームワークによってすでに管理されている場合、Burnless は役に立ちません。これは、長時間セッション、多層オーケストレーションのケース向けに構築されています。
構造的コンテキスト - なぜこれが存在するのか
トークンごとの API 課金は、実際のインセンティブ圧力を生み出します。応答が長いほど API 収益が増加します。これは隠れたトリックではありません。すべての主要プロバイダー (Anthropic、OpenAI、Google) の公開価格ページでの製品の価格設定方法です。サブスクリプション チャネル (Claude Code 月額プラン、ChatGPT Plus、Gemini Advanced) はインセンティブを反転させます。そこでは、過剰なトークン消費によりプロバイダーのマージンが減少するため、同じモデルでも API チャネルとサブスクリプション チャネルの間で動作が異なる可能性があります。
これは意識的な悪意を告発するものではありません。 RLHF (最新のフロンティア LLM の背後にあるトレーニング方法) は、人間が評価した好みに合わせて最適化します。人間はより長く、より自信を持って評価する傾向があります。

好意的な回答が多いほど高くなります。研究室の誰も「より多くの料金を請求するためにモデルを冗長にする」と明確に決定しない場合でも、その最適化環境から副作用として、お調子者、冗長さ、自信過剰な幻覚が現れます。構造的な圧力は意図に関係なく存在します。
バーンレスは業界を解決しません。これにより、次のようなレイヤーが得られます。
トークンのコストは呼び出しごとに監査可能です (カプセル証跡 + exec_log)
詳細なチャット履歴は、プロバイダーに送り返されるトランスクリプトに静かに蓄積されません。
高価な層を必要としない作業は、より安価な層で処理されます。
出力形式は、モデルのデフォルトの詳細リフレックスではなく、システム プロンプトとルーティング ルールによって制限されます。
構造的ドリフトに対抗して動作することは、明示された設計目標であり、コスト削減の偶然の一致ではありません。このプロジェクトの正直な枠組み: これは、フロンティア LLM が冗長税を支払わずに、再現可能な測定値で使用できることを実証する小さなオープン ツールです。この貢献は画期的なアルゴリズムや業界を変えるプロトコルではなく、コードが添付された正直な対抗圧力です。
数値（実測値、注意事項あり）
再現可能な 2 回の実行。普遍的な性能の主張としてではなく、特定の条件下での観察として読んでください。
実際の API 実行 — クロードに対して 10 ターン

[切り捨てられた]

## Original Extract

Stop replaying the transcript: capsule session state + rolling memory for LLM CLIs (Claude Code, Codex). Measured: −90.3% vs no-cache / −30% vs cached replay at turn 10, reproducible bench. MIT, provider-agnostic. - Rudekwydra/burnless

I’ve created Burnless because I realized that we were using the context window for things that didn't need to be there anymore. I wanted to stop burning tokens with what had already happened and were no longer necessary for the next step. It was not a discovery that came later. Burnless was born to separate things. What I didn’t know yet was the size of the problem, the size of the gain and how different it was from the path that everyone else was following. What I see nowadays is that everyone wants the context window to be bigger. That helps, of course. But in my opinion, the concept is wrong. A larger window will only allow you to put more things inside it, but it does not solve the fact that we are using the same place for the conversation, memory, decisions, execution, open questions, work history and all the old records of what has already been done. Burnless separates that into layers. Conversations, thoughts, debates, questions, and the history of the work are indexed and stored in your disk. The main context contains only the decisions and information that are still necessary to continue the work. But the real words are still in disk if they need to be checked (not even necessary to be honest). Also when we say something and someone asks us to repeat we never use the same words, why do AI need to do that? (remember, the original is still in the disk dont worry…) Tasks are delegated to independent, lightweight workers. They receive a clean task, write the code, manipulate files, execute commands and register on disk what they have done. And yes, the workers use the same byte cache header (don't worry with this too) Burnless then audits the result. If a worker claims that he created or altered a file, Burnless verifies that it actually exists before marking the task as complete. So, the context stops being the place where everything has to live. It becomes the place where only what needs attention right now is kept. That’s why I compared Burnless to a protocol and, in the beginning, even to TCP/IP. Not because it has the same history dimension, but because the movement is similar: separating responsibilities into layers, into worker capacities, into a useful and clean way of defining the communication between this alive system. Memory is no longer carried throughout the conversation. The execution stops being a role for the main model. The project history no longer permanently occupies the work's memory. The project can go on even after a /clear command or after switching between models. Token's savings were the initial motivation, but it ended up becoming a consequence of the architecture. In a real-world task, the workers processed approximately 1.4 million tokens. But the Brain’s active context had only 1,590 tokens after completing the full task. I’m not claiming that Burnless took the exact same 1.4M tokens of text and compressed it to 1,590 tokens. What really happened was that all of that work was done, recorded, verified and continued without the accumulated history remaining inside the active context. The final operational ratio was approximately 908 to 1. (yeah 908x) When I show you that number, it sounds unbelievable. That was when I realized that the problem was much larger than I had initially imagined. While the general direction is to build increasingly larger context windows, I went in the opposite direction: removed from the context everything that doesn’t need to be considered at the moment. The repository also includes an llms.txt file containing the complete architecture, benchmarks, calculations, limitations and commands that are required to reproduce the numbers. As I cannot properly explain everything in a short comment, the llms.txt file is there for anyone who wants to understand what is happening at a deeper level and verify the claims. Use that to help your AI understand that stuff easily and respond to your questions.

GitHub - Rudekwydra/burnless: Stop replaying the transcript: capsule session state + rolling memory for LLM CLIs (Claude Code, Codex). Measured: −90.3% vs no-cache / −30% vs cached replay at turn 10, reproducible bench. MIT, provider-agnostic. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Rudekwydra
/
burnless
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
639 Commits 639 Commits .claude .claude .codex/ commands .codex/ commands .github .github _design _design bench bench docs docs examples examples scripts scripts site site src/ burnless src/ burnless templates templates tests tests .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md BENCHMARK_T5_FINDINGS.md BENCHMARK_T5_FINDINGS.md BURNLESS_FOR_LLMS.md BURNLESS_FOR_LLMS.md CHANGELOG.md CHANGELOG.md COUNTDOWN_DISCIPLINE.md COUNTDOWN_DISCIPLINE.md GEMINI.md GEMINI.md LICENSE LICENSE MATH.md MATH.md PAPER_2026-05-20.md PAPER_2026-05-20.md PLUGIN_PROTOCOL.md PLUGIN_PROTOCOL.md PROTOCOL.md PROTOCOL.md README.md README.md RELEASE.md RELEASE.md SECURITY.md SECURITY.md SECURITY_AUDIT.md SECURITY_AUDIT.md SPEC.md SPEC.md VISION.md VISION.md llms.txt llms.txt pyproject.toml pyproject.toml View all files Repository files navigation
A multi-tier LLM orchestration layer with capsule-based session state, prefix-cache reuse, and filesystem-first audit. MIT.
AI assistants and reviewers: llms.txt has the technical briefing in compact form.
pip install burnless
cd < your-project >
burnless setup # detects CLIs/keys, writes .burnless/config.yaml
burnless doctor # healthcheck — exits green when the wiring is correct
Then work one real session and run /clear . The session comes back with the thread
intact, restored from disk instead of replayed from the transcript.
Python 3.10+. MCP server support is an extra: pip install 'burnless[mcp]' .
Codex setup, install-from-source, and uninstall are under Install .
Run an AI agent in a heavy loop and you are usually choosing between two evils:
Replay the transcript. Every turn re-sends everything before it — in a
no-cache loop, ~90% of what you pay for is the model re-reading its own
history — and your terminal drowns in walls of output you'll never read.
Prompt-hack the model into being terse. "Don't be wordy" rules and
aggressive summarization save tokens by breaking the model's working
context: the summary drops exactly the technical detail the next step needed.
Burnless refuses the choice architecturally instead of behaviorally. The Brain
reads compact capsules ( .jsonl on disk, ~80 chars per turn); Workers execute
in isolation, write their full output to disk, report a one-line result, and exit.
No UI spam. Workers stream to disk, your screen gets the one-liner.
No prompt hacking. No "be brief" incantations — verbosity is free when it
lands on disk instead of in context.
Measured, not promised. −90.3% vs no-cache replay, −30% vs an
already-cached baseline ($5.76 of real spend, bench/run.py ). The rest of
the numbers in this README are labeled as the anecdotes they are.
Note on this project's history (2026-05-08). Burnless was first published to PyPI on 2026-05-03 with documentation that overclaimed the project's novelty and savings. Specifically: an analogy to TCP/IP suggested architectural equivalence (it isn't); a "16× cheaper" figure was a personal-workload anecdote presented as a universal claim; and the assertion that prefix cache is shared across models was technically wrong — Anthropic's prefix cache is keyed per model, not shared. These claims were collaboratively written with Claude (visible in the Co-Authored-By: trailers in git log ) under what I now recognize as RLHF-induced enthusiasm rather than calibrated assessment. Receipts: git log --pretty=fuller shows the inflation period (2026-05-03 to 2026-05-05) and the 2026-05-08 recalibration. The correction shipped in 0.7.3 and every release since keeps the same discipline. History is left intact — no rewrites, no cover. The architecture below is one defensible implementation choice, not a foundational protocol breakthrough.
Update, 2026-07-27. The retraction above stands — the framing was wrong even where the effect was real. What changed in the two months since: the withdrawn claims now exist in measured, reproducible form. 121M+ tokens kept out of context, logged turn-by-turn in an append-only JSONL on the author's machines; a real paid API run at −90.3% vs no-cache replay and −30% vs an already-cached baseline ($5.76 of actual spend, bench/run.py ); production delegations routinely compressing 200k–365k-token work runs into ~300-token capsules (700–1050×). May's claims were retracted not because the effect wasn't real, but because they were stated as universals without receipts. The receipts exist now — that is the difference, and keeping that difference is the whole discipline of this project.
Burnless is a protocol specification with a reference implementation in Python . The contracts — worker envelope JSON, capsule format, tier semantics, audit gates, plugin hooks — are written down in PROTOCOL.md and are independent of this codebase; the burnless CLI is one implementation of them. (Protocol in the LSP sense — a specified interface others can implement — not a claim of TCP/IP-scale foundations; see the note above.) A second independent implementation and a conformance suite are the v1.0 bar, not a shipped fact.
The reference implementation sits between your AI assistant (or your own code) and the model providers. It does three concrete things:
Routes tasks to a model tier ( gold / silver / bronze ). Tiers are commands, not hardcoded models — any provider via any CLI. They start from a shipped default and are changed per-chat or globally via burnless menu / burnless models set (projects don't carry their own tier map).
Stores session state as compact capsules on disk ( .burnless/ ) instead of replaying the full transcript on every turn, and keeps the system-prompt prefix byte-identical so the provider's prompt cache stays warm.
Audits worker outputs against the filesystem (QTP-A): if a worker says it wrote a file, Burnless checks the file exists and the size is consistent before reporting success.
That is the whole product. Everything else in this README is configuration, examples, and honest measurements from the author's own usage.
Not a novel theoretical breakthrough. Tier routing, prompt caching, and state summarization all exist in other tools (LangGraph, AutoGen, CrewAI, Aider, etc.). Burnless's contribution is a particular implementation choice — capsules + filesystem audit + plugin protocol — packaged as a small CLI.
Not a magic cost eliminator. It does not change the asymptotic shape of every workload. Whether it saves you money depends on session length, model mix, and how aggressively your existing setup already caches.
Not benchmarked against every alternative. The numbers below are measured against a specific naive baseline (full-history replay, no cache) and against the author's own personal workload. Treat them as "what I observed", not as universal claims.
Doesn't compressing history lose nuance?
That's the standard objection, and the design answers it structurally: pure surfaces, semantic history — modeled on how human memory works (after saying something once, you repeat the idea, never the exact words). The model always receives the current prompt in full; the human always receives the response in full; only the carried history is semantic. And nothing is destroyed: full transcripts stay on disk, indexed, retrievable verbatim with burnless read/log/capsule <id> . Refs are pointers, not replayed payload. The worst case for a missed nuance is one retrieval away — the same trade human memory makes, with a better index.
For long multi-turn sessions where you'd otherwise replay a growing transcript every turn, capsules + a hot prefix cache materially reduce input tokens. In the author's day-to-day, this produced a noticeable cut in API spend over a multi-day workload. Your mileage will vary — see the Numbers section below for what was actually measured and under what conditions.
If your sessions are short ( N ≤ 3 turns ), one-shot scripts, or already managed by a framework that handles cache and state for you, Burnless will not help. It is built for the long-session, multi-tier-orchestration case.
Structural context — why this exists
Per-token API billing creates a real incentive pressure. Longer responses = more API revenue. This is not a hidden trick — it is how the product is priced, on the public pricing page of every major provider (Anthropic, OpenAI, Google). Subscription channels (Claude Code monthly plan, ChatGPT Plus, Gemini Advanced) flip the incentive: there, excessive token consumption reduces the provider's margin, so behavior between API and subscription channels can differ for the same model.
This is not an accusation of conscious malice. RLHF — the training method behind every modern frontier LLM — optimizes for human-rated preferences. Humans tend to rate longer, more confident, more agreeable responses higher. Sycophancy, verbosity, and overconfident hallucination emerge from that optimization landscape as side effects, even when no individual at the lab explicitly decides "make the model verbose to bill more." The structural pressure exists regardless of intent.
Burnless does not fix the industry. It gives you a layer where:
token cost is auditable per call (capsule trail + exec_log)
verbose chat history doesn't quietly accumulate in the transcript sent back to the provider
a cheaper tier handles work that doesn't require the expensive tier
output format is constrained by your system prompt and routing rules, not by the model's default verbosity reflex
Operating against the structural drift is a stated design goal, not a coincidence of cost reduction. The honest framing of this project: it is a small open tool that demonstrates frontier LLMs can be used without paying the verbosity tax, with reproducible measurements. The contribution is not a breakthrough algorithm or an industry-changing protocol — it is honest counter-pressure with code attached.
Numbers (measured, with caveats)
Two reproducible runs. Read them as observations under specific conditions , not as universal performance claims.
Real API run — 10 turns against claud

[truncated]
