---
source: "https://www.bulaev.net/p/i-had-11-ai-subagents-test-every"
hn_url: "https://news.ycombinator.com/item?id=49008916"
title: "I Had 11 AI Subagents Test Every PPTX Skill for Claude Code. Here's What Works"
article_title: "I Had 11 AI Subagents Test Every PPTX Skill for Claude Code. Here's What Actually Works."
author: "sbulaev"
captured_at: "2026-07-22T16:15:45Z"
capture_tool: "hn-digest"
hn_id: 49008916
score: 1
comments: 0
posted_at: "2026-07-22T16:02:48Z"
tags:
  - hacker-news
  - translated
---

# I Had 11 AI Subagents Test Every PPTX Skill for Claude Code. Here's What Works

- HN: [49008916](https://news.ycombinator.com/item?id=49008916)
- Source: [www.bulaev.net](https://www.bulaev.net/p/i-had-11-ai-subagents-test-every)
- Score: 1
- Comments: 0
- Posted: 2026-07-22T16:02:48Z

## Translation

タイトル: 11 人の AI サブエージェントにクロード コードのすべての PPTX スキルをテストしてもらいました。これが機能します
記事のタイトル: 11 人の AI サブエージェントにクロード コードのすべての PPTX スキルをテストしてもらいました。実際に機能するものは次のとおりです。
説明: 同じ 5 スライドのピッチデッキブリーフ、11 個のプレゼンテーションスキル、11 個の並列クロードサブエージェント、ゼロ慈悲。スクリーンショットも含まれています。

記事本文:
11 人の AI サブエージェントにクロード コードのすべての PPTX スキルをテストしてもらいました。実際に機能するものは次のとおりです。
Serge のサブスタック
11 人の AI サブエージェントにクロード コードのすべての PPTX スキルをテストしてもらいました。実際に機能するものは次のとおりです。
同じ 5 スライドのピッチデッキブリーフ、11 個のプレゼンテーションスキル、11 個の並列クロードサブエージェント、ゼロ慈悲。スクリーンショットも含まれています。
Serge Bulaev 2026 年 7 月 22 日 Share Claude コードのスキルは、約 6 か月で好奇心からエコシステムになりました。今すぐ GitHub で「pptx スキル」を検索すると、Anthropic の公式ドキュメント スキルから、ハーベイ ボールを使用したマッキンゼーのコスプレ デザイン システム、画像モデルを使用して古いデッキの外観をクローンするパイプラインまで、あらゆるものが見つかります。
誰も答えない質問: 実際に発表する資料を作成し、会議に向かう飛行機の中で PowerPoint で編集できるのは誰ですか?
私は AI ツールについてよく記事を書いていますが、私のパブリッシング パイプライン自体はエージェントによって実行されているため、エージェントのスキルをエージェントと一緒にテストすることは適切だと感じました。方法:
全員に簡単な説明を 1 つずつ。 Publora (実際の製品: ソーシャル パブリッシング API、10 プラットフォーム、1 アカウントあたり 2.99 ドル) の 5 スライドの投資家向けピッチ。固定された事実、固定されたスライド構造、1 つの難しい罠: スライド 4 には、写真ではなく、実際の表またはグラフが含まれている必要があります。
11 のスキル、11 の孤立したサブエージェント。各エージェントは 1 つのリポジトリを複製し、その SKILL.md を読み取り、フリーランスではなくスキル独自の文書化された方法論に従いました。それぞれがレポートを作成しました。何が問題でしたか (正確なエラーを含む)、何が回避されたか、使いやすさ、出力品質、編集可能性、およびドキュメントに関して 1 ～ 10 のスコアを付けました。
独立したレンダリング。結果として得られたすべてのデッキは、同じ LibreOffice → PNG パイプラインを通過し、私自身がスライドを視覚的にレビューしました。
合計: ~130 万のエージェント トークン、実測時間は約 1 時間、ディスク上には 11 デッキ。
を設定する 1 つの数字

シーン: 1 つのスキルの指示ファイルは 82,000 トークン (slides_maker の SKILL.md) まで実行できるようになりました。これはほとんどの中編小説よりも長く、エージェントは概要に触れる前にそのすべてを読みます。
アスタリスク: claude-office-skills は、Anthropic のリリース前の pptx スキルのそのままのスナップショットであることが判明しました。独自の README には、アップストリームを使用するように指示されています。行 1 ～ 2 を 1 つのエントリとして扱います。
claude-office-skills: 公式スキルのリリース前のスナップショット — ファミリーの類似性が示されています。
最も重要な発見
個々のレビューの前に、フィールド全体にわたるパターンを次に示します。
ほぼすべてのサードパーティ スキルは、テーブルを構築するのではなく、テーブルを描画します。
ブリーフでは、スライド 4 に実際の比較表を要求しました。11 のスキルのうち、ちょうど 2 つが、独自のパイプラインを通じてネイティブ OOXML テーブル/チャート オブジェクトを生成しました。公式の Anthropic スキルと slides_maker です。 2 つの最も優れたジェネレーターを含む他のユーザーは、テキスト ボックスと四角形から「表」を描画します。スクリーンショットでは同じように見えます。次に、CFO がファイルを開いて行を追加しようとすると、テーブルがあるべき場所に 62 個のゆるい図形が見つかりました。
私のサブエージェントは、概要を満たすために python-pptx 経由で実際のテーブルにパッチを当て、すべてのレポートで同じフラグが立てられました。現在、美しさと編集可能性はトレードオフであり、熟練した作成者のほとんどは美しさを選択しました。
ネイティブで編集可能なテーブル - 公式スキル。
ゴージャスで、ゆったりとした形で作られています。マッキンゼー-pptx。
1. 人間学/スキル — 退屈な正解 (6/9/9/9)
公式ドキュメント スキル バンドルは、.pptx をレンダリング ターゲットではなくドキュメント形式として扱う唯一のエントリです。実際のテーブル、実際のチャート、デザイン ルールブック、XSD バリデータ、および実行中のすべてのレイアウト欠陥を検出した必須のレンダリング アンド ルック QA ループ。
問題点: 主要な作成パット

h は、pptxgenjs がプリインストールされていることを前提としています (これは Anthropic 独自のサンドボックスにあります。私たちのサンドボックスにはありませんでした — npm ENOTCACHED )。エージェントは python-pptx に戻りましたが、スキルのツールに依存しないルールブックにより、エージェントは白紙の状態に戻されました。これは、適切に設計されたスキルの兆候です。方法論はツールを使用しても存続します。
他の人が編集するデッキが必要な場合に使用します。それはほとんどのデッキです。
報告書が水面下で明らかにしたこと: SKILL.md は苦労して獲得した約 240 行の傷跡組織であり、「# のない 16 進数の色」や「破損したファイルをオフセットするネガティブ シャドウ」などの 20 数個の落とし穴が文書化されています。 render-QA ラッパーは、GCC を使用して小さな C shim をオンザフライでコンパイルし、それを LD_PRELOAD して、LibreOffice のソケットがサンドボックス内で動作するようにします。デザイン セクションでは、タイトルの下のアクセント線、装飾的なカラー バー、クリーム色の背景など、古典的な AI スライドの指示を明示的に禁止しています。ある人間的な瞬間: プレースホルダー チェックが TODO を grep し、デッキにフラグを立てました。その一致は、ブリーフにある実際のプラットフォームである「Mas todo n」でした。
2. slides_maker — ダークホース (6/9/9/7)
独自のエンジンを通じてネイティブ テーブルとネイティブの編集可能なグラフの両方を提供する唯一のサードパーティ スキル。そのビルド→strict-lint→レンダリング→render-lint→criticループは、人間がファイルを見る前に実行中に30以上の実際の欠陥を検出しました。出力 (dark_tech プリセット) は、Python-pptx デモではなく、実際のスタートアップ デッキのように見えます。
鋭いエッジが存在します。2 つの互換性のない実行コンテンツ形式と、リンターがキャッチするまでフローチャートのテキストを白地に白で表示するダークテーマのパレット反転です。
その背後にあるスケール: 約 60 のコンポーネント (ネイティブ テーブルとチャート、KPI タイル、ダイアグラム キット)、14 のデザイン プリセット、および 1 つのコマンドですべての依存関係を検証する check_env.py プリフライトを含む、約 5,800 行の python-pptx ヘルパー ライブラリ。最初のビルドではそうではなかった

注意してください。厳密なリンターは、スライド、障害、修正に名前を付けた 19 個の CRITICAL を含むファイルの保存を拒否しました。レンダー リンターのメッセージは、測定ツールを使用するデザイナーのように読み取れます。 見えないテキスト: 塗りつぶし #FFFFFF 上の 'Call 1' インク #E8EDF5 — コントラスト 1.18:1、判読できません。これは、ほとんどの人間の査読者が与えないレベルのフィードバックです。
これは次の場合に使用します。 QA 規律が組み込まれた、 slides_maker のまさにニッチな領域 (論文/ドキュメント/コードを入力し、ネイティブに編集可能なデッキを出力) が必要な場合。
3. dashi-ppt-skill — 美しさ (7/9/6/8)
星が 4,000 個あり、視覚的な出力がその理由を説明しています。Publora のカバー スライドは、デザイナーが 2,000 ドルを請求したように見えます。エージェントのワークフロー (goal.json → レイアウト クエリ → 3 つのバリデータ → CLI エクスポート) は本当に快適で、すべてがローカルで実行され、有料 API はありません。
正直な注意点が 2 つあります。独自のエクスポーターはテキスト ボックスのみを出力します。エンジンにはネイティブのテーブルやチャートは存在しません。また、私たちの独立した LibreOffice レンダリングでは、ブラウザーのプレビューでは見られなかった、スライドの端でのテキストのクリッピングが示されました。デザインは独自のレンダラーに合わせて調整されており、ファイルが移動すると忠実度が低下します。
これは、デッキが編集ではなく表示され、レンダリングされるマシンを制御する場合に使用します。
内部では、これは実際の規律を備えたテンプレート オーケストレーターです。12 のテーマ、約 1,020 の事前構築済みレイアウト ページ、20 のセマンティック ロール、およびレイアウトごとのコントラクト ( copyKeys 、文字バジェット、配列形状) は、エージェントがデッキ ブラインドを満たし、最初のレンダリングで 3 つのバリデーターすべてに合格するほど正確です。哲学は「テンプレートをロックし、コピーを埋める」です。スタイル設定は意図的に禁止されています。導入する前に知っておく価値のある詳細: SKILL.md は中国語専用であり、文書化されたブラウザのエクスポートは設計によりスクリプトに対して 403 Forbidden を返します (人間によるクリックを想定しています)、CLI フォールバックは機能します。

彼がエクスポートした比較「テーブル」は、引かれたグリッド線の上におよそ 40 個の絶対位置に配置されたテキスト ボックスとして到着しました。つまり、5 つのスライドにまたがる 186 個の編集可能なテキスト オブジェクトであり、実際のテーブルはありません。
4. mckinsey-pptx — 効果的なコンサルタントのコスプレ (9/8/6/7)
この分野で最高の使いやすさスコア: 計画 → 40 のテンプレートから選択 → スクリプトの生成 → 構築。最初のビルドは成功しました。出力は、悪びれることなく MBB です。アクション タイトル、シェブロン プロセス フロー、ハーベイ ボール比較行列、および下部のソース行です。テストのスライド 4 は、実際のエンゲージメント資料に組み込むことができます。
注意: README は韓国語のみです (英語のテンプレート カタログが役立ちます)。また、ネイティブのテーブル/チャート オブジェクトは含まれていません。すべては形です。
コンサルティング言語がすぐに必要で、編集は PowerPoint ではなくコードを介して行われる場合に使用してください。
エンジンが手を抜いている場合でも、方法論について 2 つの詳細を読んで納得しました。まず、バンドルされたサブエージェント定義により、モデルはテンプレートの選択ごとに 1 行の根拠を記述することを強制され、885 行のカタログからの隣接するテンプレートからモデルを保護します。この 1 つのルールにより、最も怠惰なスライドの選択が無効になります。次に、5 スライドのビルド全体が Python で約 1 秒で実行されます。コーナーカットも同様に具体的です。Harvey ボールは小さな白い長方形でマスクされた円であり、フッター ルーチンはテキストの折り返しに関係なく、脚注ごとに固定 y += 0.18 ずつ進めます。これはまさに、スライド 4 で、必須のレンダリング検査ループがキャッチするまでソース行の上に脚注をスタンプする方法と同じです。
5.Academic-pptx-skill — あなたが雇う最高の編集者 (5/8/8/7)
テストの驚き。このスキルは実際にはビルダーではなく、編集方法論です。アクションタイトルの規律、「ゴーストデッキ」のアウトラインテスト、注釈ルールの展示などです。スタートアップのピッチで指摘された

(意図的にプロファイルを外して)、編集レイヤーはほぼ 1:1 で転送し、バッチの中で最も議論に富んだ健全なデッキを作成しました。
取引: 実際のファイル構築を Anthropic の組み込み pptx スキル (バンドルされていない) に委任し、その「ファンディング ピッチ」モードは空のポインターです。私たちのエージェントは技術層自体を再構築する必要がありました。
その代表的なツールは「ゴーストデッキ」テストです。デッキからアクションタイトルだけを取り除き、それらがまだ単独でストーリーを語っているかどうかをチェックします。私たちのPubloraのピッチは、2つのタイトルが書き直された後でのみ合格しました。また、展示の注釈 (すべてのチャートはラベルではなく文として要点を伝えます) も強制します。これは、読みやすいデッキと装飾的なデッキを最も分ける唯一の習慣です。それが生成した 2 つのレイアウト欠陥 (タイトルとサブタイトルの重なりと、折り返された見出しによる区切り) は両方ともハードコードされたパターン座標を追跡し、両方とも独自の修正と検証ループによって捕捉されました。
次の場合に使用します: スキル #1 と組み合わせてください。 Anthropic のビルダーとこのスキルの編集頭脳の組み合わせが、今回の実行ではどちらか一方のみを上回りました。
レベル 2: 興味深いですが、予約は必要です
thesis-defense-pptx (4/7/9/8) — ネイティブ テンプレートを再利用し、コンテンツを適切に置き換えるという時代遅れのことを行うことで、この分野で 2 番目に良い編集性スコアを獲得しました。そのダンプファースト/スキャン/検査方法は、驚くほどうまく一般化されます。ただし、パイプラインの半分は Windows PowerPoint-COM のみで、入力は PDF/LaTeX である必要があり、中国の大学のロゴがスライド レイアウトに埋め込まれています。ニッチな製品に閉じ込められた素晴らしい方法論。
Mck-ppt-design-skill (7/8/5/6) — 本格的に設計された、機械読み取り可能な QA ゲートを備えた 5 段階のハーネスで、最初のパスでデッキのスコアが 97/100 でした。ただし、文書化されたレイアウト参照層がリポジトリに単に欠落しているだけです。エンジンのパスは

~/.workbuddy にハードコーディングされており、一部のレイアウトではハードコーディングされた中国語のラベルが英語のデッキに漏洩します。
ppt-agent-skill (4/7/6/6) — 26 のスタイルは本物です。パイプラインはすぐに使えるわけではありません。SVG コンバーターはグラデーション背景を静かに削除し (スキル独自のデフォルト スタイルから白地に白のスライドを生成します)、ポリゴン ハンドラーがないため、独自の推奨矢印が消えます。 「18 のグラフ」は HTML/CSS レシピであり、グラフ オブジェクトではありません。有効なデッキを完成させるまでに 55 分間の修正が必要で、これはテストされたスキルの中で最も長い時間です。
レベル 3: 何を購入するかを把握する
ppt-image-first (4/6/2/7) — スライド上の画像を意図的に生成する規律ある設計プロトコル (3 つのプレビュー スタイル方向、ロックされた仕様)。編集可能性スコア: 2。これは欠陥ではなく、哲学です。しかし、外部画像モデルへの強い依存と組み合わせると、これは広く販売されている限定的なツールです。
gpt-image2-ppt-skills (3/4/2/6) — OpenAI キーがないと何も生成されません。画像生成は出力の 100% です。スタイル ルールは簡体字中国語テキストをハードコードします。結果は、PPTX ラッパー内のフラット PNG になります。あなたが中国語ユーザーで、アート志向の画像スライドが必要で、API コストを所有している場合は、その約束を果たします。他の人にとっては、ここには何もありません。
1. 人間性/スキル → ビルダー (ネイティブ オブジェクト、検証、QA ループ)
2.学術-

[切り捨てられた]

## Original Extract

The same 5-slide pitch deck brief, eleven presentation skills, eleven parallel Claude subagents, zero mercy. Screenshots included.

I Had 11 AI Subagents Test Every PPTX Skill for Claude Code. Here's What Actually Works.
Serge’s Substack
Subscribe Sign in I Had 11 AI Subagents Test Every PPTX Skill for Claude Code. Here's What Actually Works.
The same 5-slide pitch deck brief, eleven presentation skills, eleven parallel Claude subagents, zero mercy. Screenshots included.
Serge Bulaev Jul 22, 2026 Share Claude Code skills went from a curiosity to an ecosystem in about six months. Search GitHub for "pptx skill" today, and you'll find everything from Anthropic's official document skills to a McKinsey-cosplay design system with Harvey balls to a pipeline that clones the look of your old decks with an image model.
The question nobody answers: Which of them produce a deck you'd actually present — and then edit in PowerPoint on the flight to the meeting?
I write about AI tooling a lot, and my publishing pipeline is itself agent-run — so testing agent skills with agents felt right. The method:
One brief for everyone. A 5-slide investor pitch for Publora (a real product: social publishing API, 10 platforms, $2.99/account). Fixed facts, fixed slide structure, one hard trap: Slide 4 must contain a real table or chart — not a picture of one.
Eleven skills, eleven isolated subagents. Each agent cloned one repo, read its SKILL.md, and followed the skill's own documented methodology — no freelancing. Each wrote a report: what broke (with exact errors), what it worked around, scores 1–10 on ease of use, output quality, editability, and docs.
Independent rendering. Every resulting deck went through the same LibreOffice → PNG pipeline, and I reviewed the slides visually myself.
Total: ~1.3M agent tokens, about an hour of wall-clock time, and eleven decks on disk.
One number that sets the scene: a single skill's instruction file can now run to 82,000 tokens (slides_maker's SKILL.md) — longer than most novellas — and the agents read all of it before touching the brief.
The asterisk: claude-office-skills turned out to be a verbatim snapshot of Anthropic's pre-release pptx skill. Its own README tells you to use the upstream. Treat rows 1–2 as one entry.
claude-office-skills: a pre-release snapshot of the official skill — the family resemblance shows.
The one finding that matters most
Before the individual reviews, here's the pattern that cuts across the whole field:
Almost every third-party skill paints tables instead of building them.
The brief demanded a real comparison table on slide 4. Out of eleven skills, exactly two produced native OOXML table/chart objects through their own pipeline: the official Anthropic skill and slides_maker. Everyone else — including the two prettiest generators — draws "tables" out of text boxes and rectangles. It looks identical in a screenshot. Then your CFO opens the file, tries to add a row, and finds sixty-two loose shapes where a table should be.
My subagents patched real tables in via python-pptx to satisfy the brief, and every report flagged the same thing: beauty and editability are currently a trade-off, and most skilled authors chose beauty.
A native, editable table — the official skill.
Gorgeous — and made of loose shapes. mckinsey-pptx.
1. anthropics/skills — the boring, correct answer (6/9/9/9)
The official document-skills bundle is the only entry that treats a .pptx as a document format rather than a rendering target. Real tables, real charts, a design rulebook, an XSD validator, and a mandatory render-and-look QA loop that caught every layout defect in our run.
The catch: its primary create path assumes pptxgenjs is preinstalled (it is in Anthropic's own sandbox; it wasn’t in ours — npm ENOTCACHED ). The agent fell back to python-pptx and the skill's tool-agnostic rulebook carried it to a clean slate anyway. That's the sign of a well-designed skill: the methodology survives the tooling.
Use it when: you need decks people will edit. Which is most decks.
What the reports surfaced beneath the surface: the SKILL.md is ~240 lines of hard-won scar tissue — twenty-odd documented gotchas like "hex colors without # " and "negative shadow offsets corrupt files." The render-QA wrapper goes as far as compiling a small C shim on the fly with GCC and LD_PRELOADing it so LibreOffice's sockets work inside sandboxes. The design section explicitly bans the classic AI-slide tells — accent lines under titles, decorative color bars, and cream backgrounds. One human moment: its placeholder check greps for TODO and flagged our deck — the match was "Mas todo n," an actual platform in the brief.
2. slides_maker — the dark horse (6/9/9/7)
The only third-party skill that shipped both a native table and a native editable chart through its own engine. Its build → strict-lint → render → render-lint → critic loop caught 30+ real defects in our run before a human ever saw the file. The output (dark_tech preset) looks like a real startup deck, not a python-pptx demo.
Sharp edges exist: two incompatible runs content formats, and a dark-theme palette flip that made flowchart text white-on-white until the linter caught it.
The scale behind it: a ~5,800-line python-pptx helper library with ~60 components (native tables and charts, KPI tiles, and diagram kits), 14 design presets, and a check_env.py preflight that verifies every dependency in one command. The first build didn't just warn — the strict linter refused to save the file with 19 CRITICALs, each naming the slide, the fault and the fix. The render-linter's messages read like a designer with measurement tools: INVISIBLE TEXT: 'Call 1' ink #E8EDF5 on fill #FFFFFF — contrast 1.18:1, unreadable. That's the level of feedback most human reviewers don't give.
Use it when: you want slides_maker 's exact niche — papers/docs/code in, natively editable deck out — with QA discipline built in.
3. dashi-ppt-skill — the beauty (7/9/6/8)
4,000 stars , and the visual output explains why: our Publora cover slide looks like a designer charged $2k for it. The agent workflow (goal.json → layout queries → three validators → CLI export) is genuinely pleasant, everything runs locally, no paid APIs.
Two honest caveats. The proprietary exporter emits text boxes only — no native tables or charts exist in the engine. And our independent LibreOffice render showed text clipping at slide edges that the browser preview didn't: the design is tuned to its own renderer, and fidelity degrades once the file travels.
Use it when: the deck will be presented, not edited — and you control the machine it renders on.
Under the hood it's a template orchestrator with real discipline: 12 themes, ~1,020 pre-built layout pages, 20 semantic roles, and per-layout contracts ( copyKeys , character budgets, array shapes) precise enough that our agent filled a deck blind and passed all three validators on the first render. The philosophy is "lock the template, fill the copy" — styling is deliberately off-limits. Details worth knowing before you adopt it: the SKILL.md is Chinese-only, the documented browser export returns 403 Forbidden to scripts by design (it assumes a human clicking), the CLI fallback works, and the exported comparison "table" arrived as roughly forty absolutely-positioned text boxes over drawn grid lines — 186 editable text objects across five slides, zero real tables.
4. mckinsey-pptx — the consultant cosplay that works (9/8/6/7)
Highest ease-of-use score of the field: plan → pick from 40 templates → generate script → build. First build succeeded. The output is unapologetically MBB — action titles, chevron process flows, Harvey-ball comparison matrices, and a sources line at the bottom. Slide 4 of our test could be dropped into a real engagement deck.
Caveats: README is Korean-only (the English template catalog saves you), and — again — zero native table/chart objects; everything is shapes.
Use it when: you need the consulting dialect fast and your edits will go through code, not PowerPoint.
Two details sold me on the methodology even where the engine cuts corners. First, the bundled subagent definition forces the model to write a one-line rationale for every template pick, defending it against neighboring templates from the 885-line catalog — that one rule kills most lazy slide choices. Second, the whole 5-slide build executes in about one second of Python. The corner-cutting is equally concrete: the Harvey balls are circles masked by little white rectangles, and the footer routine advances a fixed y += 0.18 per footnote regardless of text wrapping — which is exactly how our slide 4 got a footnote stamped over the source line until the mandatory render-inspect loop caught it.
5. academic-pptx-skill — the best editor you'll hire (5/8/8/7)
The surprise of the test. This skill is not really a builder — it's an editorial methodology : action-title discipline, a "ghost deck" outline test, exhibit annotation rules. Pointed at a startup pitch (off-profile on purpose), the editorial layer transferred almost 1:1 and produced the most argumentatively sound deck of the batch.
The trade: it delegates the actual file-building to Anthropic's built-in pptx skill (not bundled), and its "funding pitch" mode is an empty pointer. Our agent had to rebuild the technical layer itself.
Its signature tool is the "ghost deck" test: strip the deck to just the action titles and check whether they still tell the story on their own — our Publora pitch passed only after two titles were rewritten. It also enforces exhibit annotation (every chart carries its takeaway as a sentence, not a label), which is the single habit that most separates readable decks from decorative ones. The two layout defects it produced — a title/subtitle overlap and a divider through a wrapped headline — both traced to hardcoded pattern coordinates, and both were caught by its own fix-and-verify loop.
Use it when: paired with skill #1. The combination — Anthropic's builder + this skill's editorial brain — beat either alone in our run.
Tier 2: interesting, with reservations
thesis-defense-pptx (4/7/9/8) — earned the field's second-best editability score by doing the unfashionable thing: reusing a native template and replacing content in place. Its dump-first/scan/inspect methodology generalizes surprisingly well. But half the pipeline is Windows PowerPoint-COM-only, input must be PDF/LaTeX, and a Chinese university logo is embedded in the slide layouts. A great methodology trapped in a niche product.
Mck-ppt-design-skill (7/8/5/6) — genuinely engineered: a five-stage harness with machine-readable QA gates that scored our deck 97/100 on the first pass. But the documented layout-reference tier is simply missing from the repo, the engine path is hardcoded to ~/.workbuddy , and some layouts leak hardcoded Chinese labels into your English deck.
ppt-agent-skill (4/7/6/6) — the 26 styles are real. The pipeline out-of-the-box is not: its SVG converter silently drops gradient backgrounds (producing white-on-white slides from the skill's own default style) and has no polygon handler, so its own recommended arrows vanish. The "18 charts" are HTML/CSS recipes, not chart objects. Fifty-five minutes of fixes to reach a valid deck — the most of any skill tested.
Tier 3: know what you're buying
ppt-image-first (4/6/2/7) — a disciplined design protocol (three previewed style directions, a locked spec) that deliberately produces images-on-slides. Editability score: 2. That's not a defect, it's the philosophy — but combined with a hard dependency on an external image model, it's a narrow tool being marketed broadly.
gpt-image2-ppt-skills (3/4/2/6) — without an OpenAI key it generates nothing: image generation is 100% of the output. The style rules hardcode Simplified Chinese text. The result is flat PNGs in a PPTX wrapper. If you're a Chinese-language user who wants art-directed image slides and owns the API cost, it delivers its promise; for everyone else there's nothing here.
1. anthropics/skills → the builder (native objects, validation, QA loop)
2. academic-

[truncated]
