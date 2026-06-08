---
source: "https://github.com/ekorbia/xword-pipeline"
hn_url: "https://news.ycombinator.com/item?id=48440123"
title: "Show HN: Dense NYT style crossword puzzle generation with Rust and Claude"
article_title: "GitHub - ekorbia/xword-pipeline: Dense NYT-style crossword puzzle pipeline: Rust fill-engine + Claude clue writer/QA/explain process · GitHub"
author: "ekorbia"
captured_at: "2026-06-08T01:00:07Z"
capture_tool: "hn-digest"
hn_id: 48440123
score: 2
comments: 0
posted_at: "2026-06-08T00:36:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Dense NYT style crossword puzzle generation with Rust and Claude

- HN: [48440123](https://news.ycombinator.com/item?id=48440123)
- Source: [github.com](https://github.com/ekorbia/xword-pipeline)
- Score: 2
- Comments: 0
- Posted: 2026-06-08T00:36:51Z

## Translation

タイトル: Show HN: Rust と Claude による濃密な NYT スタイルのクロスワード パズル生成
記事タイトル: GitHub - ekorbia/xword-pipeline: Dense NYT スタイルのクロスワード パズル パイプライン: Rust フィルエンジン + クロード クルー ライター/QA/説明プロセス · GitHub
説明: 高密度 NYT スタイルのクロスワード パズル パイプライン: Rust フィル エンジン + クロード クルー ライター/QA/説明プロセス - ekorbia/xword-pipeline
HN テキスト: 私は毎日ニューヨーク タイムズのクロスワード プレーヤーであり、独自の緻密なニューヨーク タイムズ スタイルのクロスワード パズルを作成することを目標として xword パイプラインを構築しました。パイプラインは、Rust グリッド フィル エンジンと Claude ベースのクルー ライターで構成されます。フィル エンジンは、候補グリッド スケルトン (デフォルトは 200) を生成し、Crossword Nexus Collaborative Word List の単語でそれらを埋めます。手がかり作成者はクロードを使用して、多層の手がかり (簡単、中程度、難しい)、説明を生成し、生成されたパズルを QA レビューします。独自の手がかりを作成したい場合は、手がかりライターの説明や QA ステップをスキップしたり、手がかりライターを完全にスキップしたりできます。また、パイプライン ( https://wordfuzz.com/ ) によって生成されたパズルのテストに使用できるクロスワード パズル Web アプリも作成しました。アーカイブ内のパズルはすべてパイプラインで生成されています。 QA レビューでは、これらのパズルの一部に問題があることが示されましたが、それらは解決できるはずであり、パイプラインの品質を正確に判断できるように、そのままにしておきました。自分でパイプラインを実行し、テスト ページ ( https://wordfuzz.com/test ) を使用して生成したパズルをテストできます。テスト ページは LocalStorage を使用するため、サーバーには何もアップロードされません。注: 手がかりライターには Anthropic API キーが必要で、15x15 のグリッドを複数層の手がかりと説明で埋めるのに平均約 0.50 ドルかかります。パイプラインを試してみたい場合は、

複数のパズルを生成する前に、最初に 1 つのパズルを実行し、コストが許容できるかどうかを確認してください。パイプラインに関するフィードバックをいただければ幸いです。楽しくてプレイアブルなクロスワードの生成に成功したかどうか、またはパイプラインを改善するための提案があればぜひお聞かせください。ありがとう！

記事本文:
GitHub - ekorbia/xword-pipeline: 密集した NYT スタイルのクロスワード パズル パイプライン: Rust フィルエンジン + Claude 手がかりライター/QA/説明プロセス · GitHub
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
エコルビア
/
xword パイプライン
公共
通知
通知設定を変更するにはサインインする必要があります

s
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .github .github クルーライター クルーライター docs/ イメージ docs/ イメージ fill-engine fill-engine .gitattributes .gitattributes .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md generated-batch.sh generate-batch.sh run-pipeline.sh run-pipeline.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
密度の高い NYT スタイルのクロスワード パズルを生成するためのエンドツーエンドのパイプライン:
Rust制約ソルバーフィルエンジン、多層AI手がかり書き込み（簡単→
エキスパート）、オプションの回答ごとの説明、敵対的な AI 論説
出版の門を通過します。
テーマアイデア ─▶ フィルエンジン ─▶ 手がかり ─▶ QA ⇄ 手がかり --revise
(オプション) (Rust: グリッド (クロード: (クロード: (クロード: 修正)
テーマセットの生成 + レビューの書き込み) フラグ付きの手がかり)
スコアフィル) 手がかり)
実際の動作を確認してください: wordfuzz.com がサービスを提供しています
このパイプラインによって生成されたパズル — 層ピッカーとポストソルブを含む
説明は以下に説明します。
自分で生成したパズルをプレイします。パイプライン出力 JSON を次の場所にドロップします。
wordfuzz.com/test にアクセスしてブラウザで再生します。
ファイルがマシンから離れることはありません。ページはクライアント側のレンダラーです。
このリポジトリの下には 2 つのコンポーネントが並行して存在します。
xword-パイプライン/
§── run-pipeline.sh # コマンド 1 つ: 生成 → 手がかり → QA
§── fill-engine/ # Rust: 高密度グリッド生成 + 品質スコア付きフィル
§──clude-writer/ # TypeScript: Claude 手がかりの作成、QA、テーマのアイデア作成
└── out/ # 生成されたすべての JSON (gitignored)
§── ライブラリ/ #grid-library.json / theme-library.json (fill-engine)
§── パズル/ # *.clued.json, *.qa.json, *.revized.json (clue/qa/revise)
└── テーマ/ # テーマアイデアの出力
フィルエンジンは純粋なRustで動作します

■オフライン（無料）。手がかり/QA/テーマのステップ
Claude Opus 4.7 を呼び出し、API キーが必要です。
#1. Rust エンジンを構築する
(CD フィルエンジン && カーゴビルド --release)
# 2. クロード層をインストールする
(CD クルーライター && npm インストール)
# 3. Anthropic API キーを設定します (取得方法については、clude-writer/README を参照してください)。
エクスポート ANTHROPIC_API_KEY=sk-ant-...
クイック スタート - 1 つのコマンドでパイプライン全体を実行
run-pipeline.sh は、generate→clude→qa を実行し、QA の判定を出力します。すべて
出力は out/ に到達します。引数なしで実行 ( ./run-pipeline.sh )
インタラクティブなウィザードに入り、サイズ、階層モード、説明、
そして品質レバー — テーマのないパズルのみ。テーマを設定した実行ではフラグを使用します
直接的に。 (調査結果 QA レポートの修正は、別の手動手順です。「QA レポートを修正する」を参照してください。
QA の結果を修正します。)
# テーマなし(金・土風)パズル
./run-pipeline.sh --mode テーマレス --日 土曜日
# ミニパズル - すぐに埋めて解くことができる小さなグリッド (テストに最適)
./run-pipeline.sh --mode themeless --size 5 --day Easy # 5×5 mini
./run-pipeline.sh --mode themeless --size 10 --day Medium # 10×10 midi
# テーマ別 (月曜～木曜スタイル) パズル — テーマの答えを入力します
./run-pipeline.sh --mode テーマ \
--テーマ WAITINGFORGODOT、ROCKETSCIENCE、TROMBONE --日 水曜日
グリッドステップは無料です。スクリプトは、支払われたクロードがステップを実行する前に一時停止します。
ANTHROPIC_API_KEY は設定されていません (ライブラリはまだ生成されています)。完成したら、
3 つのアーティファクト パスと判定を出力します。
旗
デフォルト
注意事項
--mode テーマなし|テーマあり
テーマのない
テーマには --theme が必要です
--サイズN
15
テーマなしのみ。グリッドの寸法 — 例:ミニの場合は 5 または 10。テーマグリッドは今のところ 15×15 です
-- テーマ A、B、C
—
テーマの回答 (A ～ Z、スペースは含まれません)。 1 ～ 4 つの回答。各 3 ～ 15 文字ですが、正確に 12 文字になることはありません (12 文字の答えを 15 文字の中に 1 つの横に置くことはできません)

幅の広いグリッド)
--ブロック N
面積の ~16% ± ~10% / 44 ± ~10%
黒い四角。デフォルトはグリッド領域の約 16% で、シード値 ±10% のジッター (15→32-40、10→14-18、5→3-5) です。テーマ グリッドのデフォルトは 44 で、同じジッター形状と 42 ブロックのフロアがあります。ジッターを無効にするには、明示的に N を渡します。 --seed 経由で再現可能
--日 DAY
土曜・水曜
手がかりの難しさ。曜日 (月曜…土曜) またはわかりやすい単語を受け入れます: 簡単 = 月、中 = 水、トリッキー = 木、難しい = 金、エキスパート = 土
--グリッド N
0
どのライブラリ グリッドを手がかりにするか (0 = 最高品質)
--keep-mean F
78
上質な床。平均解答スコア ≥ F のグリッドのみを保持します
--max-iffy N
0
重要なフィル品質レバー — 以下を参照
--候補者N
200
生成してスクリーニングするランダムグリッドの数
--時間SECS
2
グリッドごとの埋め込み予算
--トップN
20
ライブラリに保存する最適なグリッドの数
--explain-model <id>
クロード俳句-4-5
ソルブ後の説明のモデル ( --explain でのみ使用)。 claude-opus-4-7 を渡して、以前の高コストの動作を復元します。
--名前ID
<モード>-グリッド<グリッド>
出力ファイルのプレフィックスのオーバーライド。ファイルの衝突を防ぐために、実行全体で一意の名前を使用します (これはバッチ オーケストレーターによって自動的に行われます。手動で指定するのは、アドホックなマルチパズル実行の場合のみです)。
--日付 YYYY-MM-DD
—
--name が指定されていない場合は、 --nameパズル-<DATE> が自動的に導出されます。プレイヤーが正しいマニフェスト日付にパズルを挿入できるように、推奨される import-puzzle コマンドにも追加されます。
品質レバー ( --max-iffy 、 --keep-mean )
塗りつぶしの品質は、パズルの良さを左右する唯一の最大の要因です。
手がかりとQAの評決が判明した。 2 つのフラグで制御します。
--max-iffy N — 「iffy」エントリとは、単語リストのスコアが 50 未満のエントリです。
1 ～ 100 のスケール (50 = 通常の公正な回答、50 未満は疑わしいクロスワード質問、
あいまいな略語、ぎこちない部分表現）。 - マキシ

fff N はグリッドのみを保持します
最大 N 個のそのようなエントリ。
--max-iffy 0 が推奨されるデフォルトであり、ここで最も重要なフラグです。
すべての回答のスコアが 50 以上であるグリッドのみを保持し、ほとんどの回答を排除します。
単一の手がかりが見つかる前に、ソースでのすべてのカテゴリを満たす QA 結果
書かれた。 QA レポートで弱い塗りつぶしのフラグが立てられ続ける場合は、これで修正できます。
トレードオフ: ゲートが厳格になるとより多くのグリッドが拒否されるため、厳密な --max-iffy を組み合わせます
より高い --candidates (例: --max-iffy 0 --candidates 400 ) なので十分です
グリッドは生き残る。テーマ付きグリッドをきれいに埋めるのは難しいため、次のことが必要になる場合があります。
テーマモードの場合は --max-iffy 4 ～ 8 に緩和します。
--keep-mean F — 平均的な回答の質を向上させます (フロアだけでなく)。
78 ～ 80 では、洗練された生き生きとしたグリッドが得られます。通過するグリッドが少なすぎる場合は、70 に向かって下げます。
# 厳密: 完璧な塗りつぶしグリッドのみ (補うためにより多くの候補を選別する)
./run-pipeline.sh --mode テーマレス --max-iffy 0 --keep-mean 80 --candidates 400
ブロックリスト — 特定の単語を禁止する
一部のジャンク エントリは、単語リストの上位に誤ってスコア付けされています (例: TOONIEBAR は
スコアは 100) なので、 --keep-mean / --max-iffy 設定を使用してもそれらを防ぐことはできません。のために
これらを編集するには、fill-engine/data/blocklist.txt — 1 行に 1 単語ずつです。エンジン
スコアに関係なく、すべてのフィルからそれらを除外します。
# フィルエンジン/データ/ブロックリスト.txt
トゥーニーバー
イートスリップ
大文字と小文字、スペース、句読点は無視されます ( EAT SLIP == EATSLIP )。エンジン
ブロックリストを出力します: ロード時に N ワードを除外します。時間の経過とともにこのリストを増やしていきます
QA の結果 — これは、決して使用しない単語を永久に廃止する最も安価な方法です
見たいです。少数の単語をブロックしても、広告掲載率には測定可能な影響はありません。
スクリプトは QA で停止します。評決の準備ができていない場合、それを修正するかどうかは状況によって異なります。
各検出結果のカテゴリ — そして、カテゴリは 2 つの異なるレイヤーに存在します。
最初にグリッド層の検出結果を修正します - 新しい

グリッドは答えを変更します。これによりスローされます
手がかりとなる仕事をすべて遠ざけます。
グリッドレベルの修正 ( fill 、重複回答)
# 同じライブラリで次に最適なグリッドを試し、その後、手がかりを再確認し、QA を再実行します
./run-pipeline.sh --mode themeless --grid 1 --day 土曜日
# …または、弱いフィルが入り込まないようにライブラリを再生成します (その後、--grid 0)
./run-pipeline.sh --mode テーマレス --max-iffy 0 --keep-mean 80 --candidates 400
手がかりレベルの修正 (スタイル、手がかりの重複、精度、難易度)
QA レポートをフィードバックします。フラグが立てられた手がかりのみが書き換えられます。グリッドレベル
修正できない検出結果は未解決として報告されます (上記の手順に戻ります)。
CD 手がかりライター
npm 手がかりを実行 -- ../out/puzzles/themeless-grid0.clued.json \
--revise ../out/puzzles/themeless-grid0.qa.json \
--out ../out/puzzles/主題なし-グリッド0.改訂.json
npm run qa -- ../out/puzzles/themeless-grid0.revized.json
評決が完了するまで、修正→qaを繰り返します/マイナー修正。
テーマのアイデア作成 (オプション、パイプラインの前面)
テーマに行き詰まっていませんか?グリッド制約に従うセットをクロードにブレインストーミングしてもらいます。それ
テーマ化されたパイプラインに渡すことができる、すぐに実行できるコマンドを出力します。
CD 手がかりライター
npm run テーマアイデア -- --トピック " 隠されたボディパーツ " --count 3 --answers 3
# → セットを選択し、次のようにします。
cd .. && ./run-pipeline.sh --mode テーマ --主題 < A > 、 < B > 、 < C >
複数の難易度のパズルと解決後の解説
プレイヤーは、最大 4 つの手がかりセットを含む 1 つのパズルに挑戦できます。
(簡単 / 中程度 / 難しい / エキスパート) と、オプションの 1 行の説明
答えが得られるため、ソルバーは難易度を選択し、「なぜこれが当てはまるのか」という短いメモを取得します。
解決した後。各層は日ごとに調整されます。
run-pipeline.sh --tiers は 1 つのコマンドですべてを実行します。多層で
モードでは、QA は各段階をその日のルーブリックに対して個別にレビューします。
経験値に関する QA からの「土曜日には簡単すぎる」発見

ert 層には意味があります (それは
土曜日に調整された専門家の手がかりが見られます）。層ごとの出力は次のようになります。
out/puzzles/${name}.qa.${tier}.json 。
# 4 つの層すべて + 解決後の説明、エンドツーエンド:
./run-pipeline.sh --mode テーマレス --tiers easy、medium、hard、explain --explain
# --tiers を削除し、従来の 1 層パズルには --day を使用します。
スクリプトは最後に、すぐに貼り付けられる import-puzzle コマンドを出力します。の
プレイヤーの import-puzzle は、マニフェスト難易度ラベルを
バンドルの階層上限 (イージーのみ ⇒ 「イージー」、エキスパートまで ⇒
「エキスパート」)、バッジは利用可能な最も難しいコンテンツを忠実に反映しています。
ティアピッカーと説明を立ち上がらずに確認したい
プレイヤー？ *.clued.*.json ファイル (およびオプションの *.explained.json ) をドロップします。
wordfuzz.com/test へ — 同じエンジンです
毎日のパズルを強化し、ブラウザ内で完全に実行します。
コスト: ~N× 手がかり + N× QA + 1× Opus 4.7 のパズルごとの説明 (N は
層の数。下位互換性 — 古いパズルをインポートした 1 層パズル
位置指定方法 ( <clued.json> ) は引き続き機能します。
1 つのコマンドでパズルのバッチを生成する
generate-batch.sh は run-pipeline.sh をラップして、N 個の日付のパズルを生成します
1 つの調整された実行で、日付ごとに一意のファイル名が付けられます。通常の使用は1週間です
日次コンテンツ (したがって以下の例) ですが、パターンの長さは
任意 — 1 つのパズルからどこでも

[切り捨てられた]

## Original Extract

Dense NYT-style crossword puzzle pipeline: Rust fill-engine + Claude clue writer/QA/explain process - ekorbia/xword-pipeline

I’m a daily New York Times crossword player and I’ve built xword-pipeline with the goal of creating my own dense New York Times style crossword puzzles. The pipeline consists of a Rust grid fill-engine and a Claude based clue-writer. The fill-engine generates candidate grid skeletons (default is 200) and fills them with words from the Crossword Nexus Collaborative Word List. The clue-writer uses Claude to generate multi-tiered clues (easy, medium, hard), explanations, and QA review the generated puzzles. The clue-writer explanation and/or QA steps can be skipped or the clue-writer can be skipped entirely if you want to write your own clues. I’ve also created a crossword puzzle web app that can be used to test puzzles generated by the pipeline ( https://wordfuzz.com/ ). The puzzles in the archive have all been generated with the pipeline. QA reviews did show issues with some of these puzzles but they should be solvable and I’ve left them as is so the quality of the pipeline can be accurately judged. You can run the pipeline yourself and use the test page ( https://wordfuzz.com/test ) to test the puzzles you generate. The test page uses LocalStorage so nothing is uploaded to the server. Note: The clue-writer requires an Anthropic API key and costs me an average of around $0.50 to fill a 15x15 grid with multiple-tiered clues and explanations. If you’d like to give the pipeline a try, run a single puzzle first and verify the cost is acceptable to you before generating multiple puzzles. Any feedback on the pipeline would be appreciated. I would be great to hear if you’ve had success generating fun, playable crosswords or if you have any suggestions for improving the pipeline. Thanks!

GitHub - ekorbia/xword-pipeline: Dense NYT-style crossword puzzle pipeline: Rust fill-engine + Claude clue writer/QA/explain process · GitHub
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
ekorbia
/
xword-pipeline
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github .github clue-writer clue-writer docs/ images docs/ images fill-engine fill-engine .gitattributes .gitattributes .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md generate-batch.sh generate-batch.sh run-pipeline.sh run-pipeline.sh View all files Repository files navigation
An end-to-end pipeline for generating dense, NYT-style crossword puzzles :
a Rust constraint-solver fill engine, multi-tier AI clue writing (Easy →
Expert), optional per-answer explanations, and an adversarial AI editorial
pass that gates publication.
theme-idea ─▶ fill-engine ─▶ clue ─▶ qa ⇄ clue --revise
(optional) (Rust: grid (Claude: (Claude: (Claude: fix
theme sets generation + write review) flagged clues)
scored fill) clues)
See it in action: wordfuzz.com serves
puzzles generated by this pipeline — including the tier picker and post-solve
explanations described below.
Play puzzles you generate yourself: drop your pipeline output JSON at
wordfuzz.com/test to play it in the browser.
Files never leave your machine — the page is a client-side renderer.
Two components live side by side under this repo:
xword-pipeline/
├── run-pipeline.sh # one command: generate → clue → qa
├── fill-engine/ # Rust: dense grid generation + quality-scored fill
├── clue-writer/ # TypeScript: Claude clue writing, QA, theme ideation
└── out/ # ALL generated JSON (gitignored)
├── libraries/ # grid-library.json / theme-library.json (fill-engine)
├── puzzles/ # *.clued.json, *.qa.json, *.revised.json (clue/qa/revise)
└── themes/ # theme-idea output
The fill engine is pure Rust and runs offline (free). The clue/QA/theme steps
call Claude Opus 4.7 and need an API key.
# 1. Build the Rust engine
(cd fill-engine && cargo build --release)
# 2. Install the Claude layer
(cd clue-writer && npm install)
# 3. Set your Anthropic API key (see clue-writer/README for how to get one)
export ANTHROPIC_API_KEY=sk-ant-...
Quick start — the whole pipeline in one command
run-pipeline.sh runs generate → clue → qa and prints the QA verdict. All
output lands in out/ . Run with no arguments ( ./run-pipeline.sh ) to
enter an interactive wizard that walks you through size, tier mode, explainer,
and the quality levers — themeless puzzles only; themed runs use the flags
directly. (Fixing the findings QA reports is a separate manual step — see
Fixing QA findings .)
# Themeless (Fri/Sat-style) puzzle
./run-pipeline.sh --mode themeless --day Saturday
# Mini puzzles — small grids that are quick to fill & solve (great for testing)
./run-pipeline.sh --mode themeless --size 5 --day Easy # 5×5 mini
./run-pipeline.sh --mode themeless --size 10 --day Medium # 10×10 midi
# Themed (Mon–Thu-style) puzzle — supply the theme answers
./run-pipeline.sh --mode themed \
--themes WAITINGFORGODOT,ROCKETSCIENCE,TROMBONE --day Wednesday
The grid step is free; the script pauses before the paid Claude steps if
ANTHROPIC_API_KEY isn't set (your library is still generated). On completion it
prints the three artifact paths and the verdict.
Flag
Default
Notes
--mode themeless|themed
themeless
Themed requires --themes
--size N
15
Themeless only. Grid dimension — e.g. 5 or 10 for minis. Themed grids are 15×15 for now
--themes A,B,C
—
Theme answers (A–Z, no spaces). 1–4 answers; each 3–15 letters but never exactly 12 (a 12-letter answer can't be placed as a single across in a 15-wide grid)
--blocks N
~16% of area ± ~10% / 44 ± ~10%
Black squares. Default is ~16% of the grid area with a seeded ±10% jitter (15→32-40, 10→14-18, 5→3-5); themed grids default to 44 with the same jitter shape and a 42-block floor. Pass an explicit N to disable jitter. Reproducible via --seed
--day DAY
Saturday / Wednesday
Clue difficulty. Accepts a day (Monday…Saturday) or a friendly word: Easy =Mon, Medium =Wed, Tricky =Thu, Hard =Fri, Expert =Sat
--grid N
0
Which library grid to clue ( 0 = highest quality)
--keep-mean F
78
Quality floor. Keep only grids whose mean answer-score ≥ F
--max-iffy N
0
The key fill-quality lever — see below
--candidates N
200
How many random grids to generate & screen
--time SECS
2
Per-grid fill budget
--top N
20
How many of the best grids to keep in the library
--explain-model <id>
claude-haiku-4-5
Model for the post-solve explainer (used only with --explain ). Pass claude-opus-4-7 to restore the previous, higher-cost behavior
--name ID
<mode>-grid<GRID>
Output-file prefix override. Use unique names across runs to prevent file collisions (the batch orchestrator does this automatically; specify manually only for ad-hoc multi-puzzle runs)
--date YYYY-MM-DD
—
If --name is not given, auto-derives --name puzzle-<DATE> . Also appended to the suggested import-puzzle command so the player slots the puzzle into the right manifest date
The quality levers ( --max-iffy , --keep-mean )
Fill quality is the single biggest driver of how good the puzzle — and therefore
the clues and the QA verdict — turns out. Two flags control it:
--max-iffy N — an "iffy" entry is one scoring below 50 on the wordlist's
1–100 scale (50 = a normal, fair answer; below 50 is questionable crosswordese,
obscure abbreviations, awkward partials). --max-iffy N keeps only grids with
at most N such entries.
--max-iffy 0 is the recommended default and the most important flag here.
It keeps only grids where every answer scores ≥ 50, which eliminates almost
all fill -category QA findings at the source — before a single clue is
written. If your QA reports keep flagging weak fill, this is the fix.
Trade-off: stricter gates reject more grids, so pair a strict --max-iffy
with a higher --candidates (e.g. --max-iffy 0 --candidates 400 ) so enough
grids survive. Themed grids are harder to fill cleanly, so you may need to
relax to --max-iffy 4 – 8 for themed mode.
--keep-mean F — raises the average answer quality (not just the floor).
78–80 yields polished, lively grids; lower it toward 70 if too few grids pass.
# Strict: only flawless-fill grids (screen more candidates to compensate)
./run-pipeline.sh --mode themeless --max-iffy 0 --keep-mean 80 --candidates 400
Blocklist — banning specific words
Some junk entries are mis-scored high in the wordlist (e.g. TOONIEBAR is
scored 100), so no --keep-mean / --max-iffy setting can keep them out. For
those, edit fill-engine/data/blocklist.txt — one word per line; the engine
excludes them from every fill regardless of score:
# fill-engine/data/blocklist.txt
TOONIEBAR
EATSLIP
Case, spaces, and punctuation are ignored ( EAT SLIP == EATSLIP ). The engine
prints blocklist: excluding N word(s) on load. Grow this list over time from
your QA findings — it's the cheapest way to permanently retire words you never
want to see. Blocking a handful of words has no measurable effect on fill rate.
The script stops at QA. If the verdict isn't ready , fixing it depends on the
category of each finding — and the categories live at two different layers :
Fix grid-layer findings first — a new grid changes the answers, which throws
away any clue work.
Grid-level fixes ( fill , duplicate-answers)
# Try the next-best grid in the same library, then re-clue + re-QA
./run-pipeline.sh --mode themeless --grid 1 --day Saturday
# …or regenerate the library so weak fill can't get in (then --grid 0)
./run-pipeline.sh --mode themeless --max-iffy 0 --keep-mean 80 --candidates 400
Clue-level fixes ( style , duplicate-in-clue, accuracy, difficulty)
Feed the QA report back; only the flagged clues are rewritten. Grid-level
findings it can't fix are reported as unresolved (go back to the step above).
cd clue-writer
npm run clue -- ../out/puzzles/themeless-grid0.clued.json \
--revise ../out/puzzles/themeless-grid0.qa.json \
--out ../out/puzzles/themeless-grid0.revised.json
npm run qa -- ../out/puzzles/themeless-grid0.revised.json
Repeat revise → qa until the verdict is ready / minor-revisions .
Theme ideation (optional, front of pipeline)
Stuck for a theme? Have Claude brainstorm sets that obey the grid constraints; it
prints ready-to-run commands you can hand to the themed pipeline.
cd clue-writer
npm run theme-idea -- --topic " hidden body parts " --count 3 --answers 3
# → pick a set, then:
cd .. && ./run-pipeline.sh --mode themed --themes < A > , < B > , < C >
Multi-difficulty puzzles & post-solve explanations
The player can take a single puzzle that carries up to four clue sets
(Easy / Medium / Hard / Expert) plus an optional one-line explanation per
answer, so the solver picks a difficulty and gets a short "why this fits" note
after solving. Each tier is day-calibrated:
run-pipeline.sh --tiers does the whole thing in one command. In multi-tier
mode, QA reviews each tier independently against its own day rubric, so a
"too easy for Saturday" finding from QA on the expert tier is meaningful (it
sees expert clues calibrated to Saturday). Per-tier outputs land at
out/puzzles/${name}.qa.${tier}.json .
# All four tiers + post-solve explanations, end-to-end:
./run-pipeline.sh --mode themeless --tiers easy,medium,hard,expert --explain
# Drop --tiers and use --day for legacy single-tier puzzles.
The script prints a ready-to-paste import-puzzle command at the end. The
player's import-puzzle derives the manifest difficulty label from the
tier ceiling of the bundle (easy-only ⇒ "Easy", up through expert ⇒
"Expert"), so the badge truthfully reflects the hardest content available.
Want to see the tier picker and explanation reveal without standing up the
player? Drop your *.clued.*.json files (and optional *.explained.json )
onto wordfuzz.com/test — it's the same engine
that powers the daily puzzles, running entirely in your browser.
Cost: ~N× clue + N× QA + 1× explain per puzzle on Opus 4.7, where N is the
number of tiers. Backward compatible — single-tier puzzles imported the old
positional way ( <clued.json> ) still work.
Generating a batch of puzzles in one command
generate-batch.sh wraps run-pipeline.sh to produce N dated puzzles in
one orchestrated run, with unique filenames per date. Common use is a week
of daily content (hence the example below), but the pattern length is
arbitrary — anywhere from one puzzl

[truncated]
