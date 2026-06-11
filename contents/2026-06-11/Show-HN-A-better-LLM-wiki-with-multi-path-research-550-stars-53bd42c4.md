---
source: "https://llm-wiki.net/"
hn_url: "https://news.ycombinator.com/item?id=48489778"
title: "Show HN: A better LLM-wiki with multi-path research [550 stars]"
article_title: "LLM Wiki — LLM-compiled knowledge bases for Claude Code, Codex, OpenCode & any LLM agent"
author: "nvk"
captured_at: "2026-06-11T13:27:45Z"
capture_tool: "hn-digest"
hn_id: 48489778
score: 2
comments: 0
posted_at: "2026-06-11T13:02:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A better LLM-wiki with multi-path research [550 stars]

- HN: [48489778](https://news.ycombinator.com/item?id=48489778)
- Source: [llm-wiki.net](https://llm-wiki.net/)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T13:02:01Z

## Translation

タイトル: HN を表示: マルチパス研究を備えたより良い LLM-wiki [星 550 個]
記事のタイトル: LLM Wiki — Claude Code、Codex、OpenCode、およびあらゆる LLM エージェント用の LLM でコンパイルされたナレッジ ベース
説明: LLM Wiki は、LLM エージェントを研究および知識ベースのエンジンに変えます。並行調査、コレクター カタログ、ソースの取り込み、Wiki の編集、トピックのアーカイブ、インベントリ追跡、データセット マニフェスト、真実を追求する監査、クエリ、およびアーティファクトの生成。 Claude Code プラグイン、OpenAI として出荷
[切り捨てられた]

記事本文:
コンテンツにスキップ
◩
llm-wiki
NVKによって
インストール
使用法
コマンド
ガイド
よくある質問
ギットハブ →
学習プロンプト
██╗ ██╗ ███╗ ███╗ ██╗ ██╗██╗██╗ ██╗██╗
██║ ██║ ████╗ ████║ ██║ ██║██║██║ ██╔╝██║
██║ ██║ ██╔████╔██║ ██║ █╗ ██║██║█████╔╝ ██║
██║ ██║ ██║╚██╔╝██║ ██║███╗██║██║██╔═██╗ ██║
███████╗███████╗██║ ╚═╝ ██║ ╚███╔███╔╝██║██║ ██╗██║
╚══════╝╚══════╝╚═╝ ╚═╝ ╚══╝╚══╝ ╚═╝╚═╝ ╚═╝╚═╝
LLM でコンパイルされたナレッジ ベース
素晴らしい出力を備えた AI エージェント。
マルチエージェントの並行研究。論文主導の調査。ソース摂取。
ウィキのまとめ。トピックのアーカイブ。在庫追跡。データセットマニフェスト。真実を追求する監査。
問い合わせ中。アーティファクトの生成。として出荷されます
Claude Code プラグイン、OpenAI Codex プラグイン、OpenCode 命令ファイル、または
ポータブル AGENTS.md 。オブシディアン対応。
走るたびに複合的になる。ソースは相互参照記事になります。
記事はレポート、スライドデッキ、学習ガイド、プレイブック、
そして実施計画。研究すればするほど強くなる
すべての出力が取得されます。
1 つのコマンドでトピック Wiki を起動し、最大 10 人のエージェントを派遣します。
保持する価値のあるものを取り込み、追跡する前に出所が豊富なカタログを収集し、古いトピックを削除せずにアーカイブし、永続的なフォローアップ状態を追跡し、インデックスを作成します。
大規模なデータセットをコピーせずに作成し、ソースを記事にコンパイルし、
に基づいて構築された成果物をあなたに渡します

p.あなたが所有するすべてのプレーンな Markdown。
5 ～ 10 人の並行エージェントが学術、技術、応用、ニュース、逆張りの角度から検索します。 --min-time 2h はラウンドを続け、各ラウンドで見つかったギャップを掘り下げます。
主張から始めましょう。エージェントは、サポート、反対、メカニズム、メタ、および隣接する角度に分かれます。出力は結論であり、要約ではありません。ラウンド 2 は確証バイアスとの戦いです。
URL、ファイル、PDF、受信トレイのドロップ、Git ドキュメント リポジトリ、MediaWiki ダンプ、メッセージ アーカイブ、および Wayback CDX スナップショット。生のソースは不変のままです。記事は上部に合成されます。
制限されたパブリック メディアを検索、重複排除、ダウンロードし、発見可能なアーティファクト、例、ミーム、ツール、エンティティ、およびソース候補をカタログ化します。エイリアス、コンテキスト内で見つかった出所、ローカル アセット パス、ハッシュ、スケール、メディア ポリシー、および在庫の適合性をキャプチャします。
項目、ソース候補、コーパス、エンティティ、未解決の質問、監視項目、次のアクションなど、ウィキが覚えておくべき永続的なものを追跡します。チャット ビューのデフォルトはコンパクトなテーブルです。
マニフェスト、サンプル、プロファイル、クエリ レシピを使用して、大規模な外部データ、変更可能なデータ、または運用データのインデックスを作成します。 Wiki がインターフェースになります。データはその場所に残ります。
トピック Wiki 全体を topic/.archive/ に移動します。保存されたナレッジは構造的に保守可能ですが、デフォルトのクエリ、コンパイル、調査、収集、出力、および保守コンテキストから外れます。
生のソースは、相互参照と信頼スコアを備えた合成記事になります。すべてのディレクトリには _index.md があり、盲目的にスキャンされることはありません。
クイック (インデックス)、標準 (記事)、またはディープ (すべて + 兄弟 Wiki)。 --resume は中断したところから再開します。
すべての記事の古さと品質をスコア付けします。 2 層スキャン: メタデータを高速にチェックし、フラグが設定された記事の詳細なコンテンツを読み取ります。チェックポイントの回復。機械可読な JSON + 人間が判読可能なレポート。
より広い範囲の t を答えてください

錆びの質問です。ライブラリアン パスを再利用し、 raw/ 、 wiki/ 、および Output/ にわたる出力を追跡し、ドリフトを検出し、来歴を検査し、ローカル証拠が十分でない場合は新たな調査を行います。
現在のセッションから学んだ教訓、つまりエラー→修正パターン、ユーザーの修正、発見を抽出します。構造化されたメモとして保存され、Wiki で後でクエリできるようになります。 --rules は、散文ではなく強制可能なルールを発行します。
Wiki に基づいた実装計画。ナレッジ ベースを読み、要件についてインタビューし、対象を絞った調査でギャップを埋め、証拠として Wiki 記事を引用しながら段階的な計画を作成します。 --format rfc|adr|spec 。
レポート、スライドデッキ、学習ガイド、プレイブック、実装計画、タイムライン、用語集、比較。次の出力が以前のすべての出力に基づいて構築されるように、Wiki に再度ファイルされます。
クロード プラグインのインストール wiki@llm-wiki
パブリック マーケットプレイスからインストールします。クロードコードを再起動して適用します。
マーケットプレイスのプラグイン。 @wiki を使用して呼び出します。
コーデックス プラグイン マーケットプレイスに nvk/llm-wiki を追加
# その後、/plugins を開き、「LLM Wiki」を有効にし、@wiki を使用します
または、ローカル チェックアウトから: ./scripts/bootstrap-codex-plugin.sh --scope user --verify 。コーデックス ツリーはクロードの信頼できる情報源の生成されたミラーであり、土地を同様に更新します。
# opencode.json 内:
{ "説明書": [
「パス/to/llm-wiki/plugins/llm-wiki-opencode/skills/wiki-manager/SKILL.md」
] }
または、 ~/.config/opencode/AGENTS.md にコピーします。 Web 検索には OPENCODE_ENABLE_EXA=1 が必要です。
説明書ファイル。ローカルモデルに最適です。
pi --instructions path/to/llm-wiki/plugins/\
llm-wiki-opencode/skills/wiki-manager/SKILL.md
Pi の 1K システム プロンプトには、32K コンテキスト ローカル モデルでの完全な Wiki スキルの余地が残されています。 OpenCodeと同じスキルファイルを使用します。
カール -sL https://raw.githubusercontent.com/nvk/llm-wiki/master/AGENTS.md \
> ~/あなたのプロジェクト/AGENTS.md
ファイルを任意のエージェントにドロップします。

コンテキストまたはプロジェクトのルート。ファイルの読み取り/書き込みと Web 検索ができるものなら何でも動作します。
クロード プラグイン更新 wiki@llm-wiki
# クロードコードを再起動して適用します
更新コマンドで新しいバージョン (古いマーケットプレイスのキャッシュ) が見つからない場合は、手動で同期します。
git clone https://github.com/nvk/llm-wiki.git # または: git -C ~/llm-wiki pull
REPO=~/llm-wiki/claude-plugin
DEST=~/.claude/plugins/cache/llm-wiki/wiki
VERSION=$(grep '"バージョン"' "$REPO/.claude-plugin/plugin.json" | grep -o '[0-9.]*')
rm -rf "$DEST"/*
mkdir -p "$DEST/$VERSION"
cp -R "$REPO/.claude-plugin" "$REPO/コマンド" "$REPO/スキル" "$DEST/$VERSION/"
コーデックス: コーデックス プラグイン マーケットプレイスのアップグレード llm-wiki 。ローカル チェックアウトの場合: ./scripts/bootstrap-codex-plugin.sh --scope user --verify を再実行します。
AGENTS.md: 上記のcurlコマンドを再実行してファイルを置き換えます。
どこからでも、1 つのコマンドでトピック Wiki を作成し、並列エージェントを起動し、1 時間調査を続けて、コンパイルされて戻ってきます。
/wiki:research "腸内マイクロバイオーム" --new-topic --min-time 1h
より一般的なフロー:
/wiki:研究「栄養」 --新しいトピック
/wiki:research "断食" --deep --min-time 2h
/wiki:research 「記事が口コミで広まる理由は何ですか?」 --新しいトピック
/wiki:research --mode 論文「繊維は SCFA を介して神経炎症を軽減する」
/wiki:query 「食物繊維は気分にどのように影響しますか?」
/wiki:query --resume
/wiki 追加 https://example.com/article # ファジールーター → 取り込み
/wiki CRISPR について何がわかっていますか? # ファジールータ → クエリ
/wiki:ingest-collection https://github.com/bitcoin/bips --wiki ビットコイン
/wiki:「ビットコイン ミーム」を収集 --wiki ビットコイン
/wiki:「ビットコイン ミーム」を収集 --スケール 中 --メディア リファレンス --インベントリ コーパス
/wiki:在庫項目「TRX-4M リングとピニオン」を追加 --wiki trx4m-1-18
/wiki:インベントリ リスト --アクションの表示 --制限 10
/wiki:dataset 「Bitcointalk 時間グラフ」を追加 --location https://figshare.c

om/articles/dataset/BitcoinTemporalGraph/26305093
/wiki:データセット リスト --スキーマの表示 --制限 10
/wiki:archive topic old-interest --reason "現在はアクティブではありません"
/wiki:アーカイブ リスト --アーカイブ済み
/wiki:アーカイブ復元古い関心事
/wiki:コンパイル
/wiki:出力レポート --トピック腸脳
/wiki:assess /path/to/my-app --wiki 栄養
/wiki:lint --fix
仕組み
~/wiki/ # ハブ — 軽量、コンテンツなし
§── wikis.json # すべてのトピック Wiki のレジストリ
§── _index.md # トピック Wiki を統計付きでリストします。
§── log.md # グローバルアクティビティログ
└── トピックス/ # 各トピックは独立した Wiki です
§── 栄養/
│ §── .obsidian/ # 黒曜石保管庫の設定
│ §── inbox/ # このトピックのドロップゾーン
│ §── 在庫/ # アイテム、候補、コーパス、ビュー
│ §── datasets/ # 大規模/外部データのマニフェスト
│ §── raw/ # 不変のソース
│ §── wiki/ # まとめ記事
│ │ §── コンセプト/
│ │ §── トピックス/
│ │ └─ 参考文献/
│ §── 出力/ # 生成されたアーティファクト
│ §── _index.md
│ §── config.md
│ └── log.md
§── 木工/ # 別の話題 wiki
└── .archive/ # アーカイブされたトピック Wiki、デフォルトでは非表示
1 つのトピック、1 つの Wiki
各研究分野は孤立しています。トピックをまたぐノイズはありません。クエリは集中し続けます。マルチ Wiki を覗いてみると、関連する場合は重複が見つかります。
Obsidian の [[wikilinks]] とその他すべての標準のマークダウン リンク。まったくビューアを使用しない場合も含め、すべてのビューアで動作します。
ソースが取り込まれると、変更されることはありません。上部に記事が合成されます。後退させると両方ともきれいに取り外せます。
/wiki: 証拠や永続的な状態になる前に、レコードのエイリアス、ソース コンテキスト、メディア URL、キャッシュされたアセット パス、ハッシュ、重複排除メモ、スケール、およびインベントリの推奨事項を収集します。
パーツ、ソースキュー、コア

ポラ、ウォッチアイテム、次のアクションはインベントリの下に保存されているため、証拠にならずにリスト化して再確認できます。
データセット/大規模データのマニフェスト、サンプル、プロファイル、クエリ レシピを保存します。 Wiki は、データをソース コーパスにコピーせずにインデックスを作成します。
アーカイブされたトピックは、 topic/.archive/ の下に保存されます。ほとんどのツールはデフォルトでこれらをスキップします。深いクエリによってインデックス ヒットが表面化する可能性があり、明示的な --include-archived によってそれらを読み取ることができます。
ホスト エージェントの組み込みツール上で完全に実行されます。プラグインはMarkdown + コマンドです。サーバーもサービスもテレメトリもありません。
すべてのコマンドは、トピック Wiki をターゲットとする --wiki <name> と、プロジェクト Wiki の --local を受け入れます。アーカイブされたトピック Wiki はデフォルトではスキップされます。 --include-archived をサポートするコマンドでは、アーカイブされたマテリアルの読み取りまたは書き込みの前に明示的なフラグが必要です。 query 、 Output 、 plan は、クロスウィキ コンテキストの --with <wiki> も受け入れます。
ゼロからコンパイルされた Wiki まで 5 分で完了します。
クロード プラグインのインストール wiki@llm-wiki
2. トピック Wiki を作成する
興味のあるトピックを選択してください:
/wiki 初期栄養
これにより、 ~/wiki/ にハブが作成され、 ~/wiki/topics/nutrition/ に最初のトピック Wiki が作成されます。
/wiki:研究「腸内微生物叢とメンタルヘルス」 --wiki 栄養学
# あるいは自然に言ってみましょう:
/wiki 腸内微生物叢とメンタルヘルスの研究
5 人の並行エージェントがさまざまな角度 (学術、技術、応用、ニュース、逆張り) から Web を検索し、最良のソースを取り込み、それらを相互参照された Wiki 記事に編集します。 2 ～ 5 分かかります。
/wiki:query "食物繊維は気分にどのように影響しますか?" --wikiの栄養
# あるいは自然に:
/wiki 食物繊維は気分にどのように影響しますか?
Wiki は、引用とともにまとめられた記事から回答します。
5. 出力を信頼する前に監査する
/wiki:audit --wiki 栄養
/wiki:audit --artifact 出力/report-gut-brain.md
監査は wiki レイヤーを再チェックし、出力を追跡します

ut の証拠チェーン、フラグが漂い、ローカル コーパスが信頼性の質問に答えるのに十分でない場合は、新たな調査が行われます。
/wiki:research "topic" --deep — 5 つのエージェントの代わりに 8 つのエージェントが追加され、歴史とデータの角度が追加されます。
/wiki:research "topic" --min-time 1h — 1 時間ラウンドで研究を続けます
/wiki:research "topic" --plan — 並行した研究パスに分解します
/wiki:audit --project Nutrition-Playbook — 出力と上流の Wiki の状態を一緒に検証します
/wiki 追加 https://example.com/article — ファジー ルーターが URL を検出し、それを取り込みます
/wiki CRISPR について何がわかっていますか? — ファジールーターが質問とクエリを検出します
/wiki:lint --fix — 構造的な問題をクリーンアップします
一般的なリサーチ セッションは、次の 4 つの段階を経て行われます。
ステージ 1: 質問するかトピックを選択する
llm-wiki は、質問しているのかトピックに名前を付けているのかを自動検出します。ダイレクト コマンドまたはファジー ルーターを使用します。
/wiki:research 「長文記事がバイラルになるのはなぜですか?」 # 直接コマンド
/wiki 研究量子コンピューティング # ファジールーター — 同じ結果
/wiki:research --mode 論文「繊維は SCFA を介して神経炎症を軽減する」 # 論文 → 証拠に賛成/反対
ステージ 2: エージェントは並行して検索します
5 人のエージェント ( --deep で 8 人、 --retardmax で 10 人) がさまざまな角度から同時に検索します。それぞれ 2 ～ 3 つの Web 検索、フルコンテンツのフェッチ、品質スコアリング (1 ～ 5)。信頼性パスは、取り込み前に重複を排除します。
ステージ 3: ソースが取り込まれてコンパイルされる
上位のソースは raw/ に保存されます (不変 - 取り込み後に変更されることはありません)。次に、コンパイル パスは、相互参照、信頼スコア、および双方向リンクを使用して、それらを wiki/concepts/ 、 wiki/topics/ 、および wiki/references/ の下の wiki 記事に合成します。
ステージ 4: ギャップレポートとフォローアップ
各ラウンドの後、何がカバーされているか、何がまだ欠けているかがわかります。

フォローアップの提案。 2 つ以上のギャップが残っている場合は、それらを並行して埋めるように提案されます。

[切り捨てられた]

## Original Extract

LLM Wiki turns any LLM agent into a research and knowledge-base engine. Parallel research, collector catalogs, source ingestion, wiki compilation, topic archiving, inventory tracking, dataset manifests, truth-seeking audits, querying, and artifact generation. Ships as a Claude Code plugin, an OpenAI
[truncated]

Skip to content
◩
llm-wiki
by nvk
Install
Usage
Commands
Guides
FAQ
GitHub →
learntoprompt
██╗ ██╗ ███╗ ███╗ ██╗ ██╗██╗██╗ ██╗██╗
██║ ██║ ████╗ ████║ ██║ ██║██║██║ ██╔╝██║
██║ ██║ ██╔████╔██║ ██║ █╗ ██║██║█████╔╝ ██║
██║ ██║ ██║╚██╔╝██║ ██║███╗██║██║██╔═██╗ ██║
███████╗███████╗██║ ╚═╝ ██║ ╚███╔███╔╝██║██║ ██╗██║
╚══════╝╚══════╝╚═╝ ╚═╝ ╚══╝╚══╝ ╚═╝╚═╝ ╚═╝╚═╝
LLM-compiled knowledge bases for
any AI agent with awesome outputs.
Parallel multi-agent research. Thesis-driven investigation. Source ingestion.
Wiki compilation. Topic archiving. Inventory tracking. Dataset manifests. Truth-seeking audits.
Querying. Artifact generation. Ships as a
Claude Code plugin, an OpenAI Codex plugin, an OpenCode instruction file, or a
portable AGENTS.md . Obsidian-compatible.
Every run compounds. Sources become cross-referenced articles.
Articles become reports, slide decks, study guides, playbooks,
and implementation plans. The more you research, the stronger
every output gets.
One command spins up a topic wiki, dispatches up to ten agents,
ingests what's worth keeping, collects provenance-rich catalogs before tracking them, archives old topics without deleting them, tracks durable follow-up state, indexes
large datasets without copying them, compiles sources into articles, and
hands you a deliverable built on top. All plain Markdown you own.
5–10 parallel agents search academic, technical, applied, news, and contrarian angles. --min-time 2h keeps going in rounds, drilling into gaps each round finds.
Start from a claim. Agents split across supporting, opposing, mechanistic, meta, and adjacent angles. Output is a verdict — not a summary. Round two fights confirmation bias.
URLs, files, PDFs, inbox drops, Git doc repos, MediaWiki dumps, message archives, and Wayback CDX snapshots. Raw sources stay immutable; articles synthesize on top.
Find, dedupe, download bounded public media, and catalog discoverable artifacts, examples, memes, tools, entities, and source candidates. Captures aliases, found-in-context provenance, local asset paths, hashes, scale, media policy, and inventory fit.
Track durable things the wiki should remember: items, source candidates, corpora, entities, open questions, watch items, and next actions. Chat views default to compact tables.
Index large, external, mutable, or operational data with manifests, samples, profiles, and query recipes. The wiki becomes the interface; the data stays where it belongs.
Move whole topic wikis to topics/.archive/ . Preserved knowledge stays structurally maintainable but out of default query, compile, research, collect, output, and maintenance context.
Raw sources become synthesized articles with cross-references and confidence scores. Every directory has an _index.md — nothing is scanned blindly.
Quick (indexes), standard (articles), or deep (everything + sibling wikis). --resume picks up where you left off.
Score every article for staleness and quality. Two-tier scan: fast metadata check, then deep content read for flagged articles. Checkpoint recovery. Machine-readable JSON + human-readable report.
Answer the broader trust question. Reuse the librarian pass, trace outputs across raw/ , wiki/ , and output/ , detect drift, inspect provenance, and do fresh research when local evidence is not enough.
Extract lessons learned from the current session — error→fix patterns, user corrections, discoveries. Saved as structured notes the wiki can query later. --rules emits enforceable rules instead of prose.
Wiki-grounded implementation plans. Reads the knowledge base, interviews you about requirements, fills gaps with targeted research, and produces a phased plan citing wiki articles as evidence. --format rfc|adr|spec .
Reports, slide decks, study guides, playbooks, implementation plans, timelines, glossaries, comparisons. Filed back into the wiki so the next output builds on every previous one.
claude plugin install wiki@llm-wiki
Installs from the public marketplace. Restart Claude Code to apply.
Marketplace plugin. Invoke with @wiki .
codex plugin marketplace add nvk/llm-wiki
# Then open /plugins, enable "LLM Wiki", use @wiki
Or from a local checkout: ./scripts/bootstrap-codex-plugin.sh --scope user --verify . The Codex tree is a generated mirror of the Claude source of truth — updates land identically.
# In opencode.json:
{ "instructions": [
"path/to/llm-wiki/plugins/llm-wiki-opencode/skills/wiki-manager/SKILL.md"
] }
Or copy to ~/.config/opencode/AGENTS.md . Web search requires OPENCODE_ENABLE_EXA=1 .
Instruction file. Best for local models.
pi --instructions path/to/llm-wiki/plugins/\
llm-wiki-opencode/skills/wiki-manager/SKILL.md
Pi's 1K system prompt leaves room for the full wiki skill on 32K context local models. Uses the same skill file as OpenCode.
curl -sL https://raw.githubusercontent.com/nvk/llm-wiki/master/AGENTS.md \
> ~/your-project/AGENTS.md
Drop the file into any agent's context or project root. Works with anything that can read/write files and search the web.
claude plugin update wiki@llm-wiki
# Restart Claude Code to apply
If the update command misses a new version (stale marketplace cache), sync manually:
git clone https://github.com/nvk/llm-wiki.git # or: git -C ~/llm-wiki pull
REPO=~/llm-wiki/claude-plugin
DEST=~/.claude/plugins/cache/llm-wiki/wiki
VERSION=$(grep '"version"' "$REPO/.claude-plugin/plugin.json" | grep -o '[0-9.]*')
rm -rf "$DEST"/*
mkdir -p "$DEST/$VERSION"
cp -R "$REPO/.claude-plugin" "$REPO/commands" "$REPO/skills" "$DEST/$VERSION/"
Codex: codex plugin marketplace upgrade llm-wiki . For a local checkout: re-run ./scripts/bootstrap-codex-plugin.sh --scope user --verify .
AGENTS.md: re-run the curl command above to replace the file.
One command, from anywhere — creates a topic wiki, launches parallel agents, keeps researching for an hour, comes back compiled.
/wiki:research "gut microbiome" --new-topic --min-time 1h
More common flows:
/wiki:research "nutrition" --new-topic
/wiki:research "fasting" --deep --min-time 2h
/wiki:research "What makes articles go viral?" --new-topic
/wiki:research --mode thesis "fiber reduces neuroinflammation via SCFAs"
/wiki:query "How does fiber affect mood?"
/wiki:query --resume
/wiki add https://example.com/article # fuzzy router → ingest
/wiki what do we know about CRISPR? # fuzzy router → query
/wiki:ingest-collection https://github.com/bitcoin/bips --wiki bitcoin
/wiki:collect "bitcoin memes" --wiki bitcoin
/wiki:collect "bitcoin memes" --scale medium --media reference --inventory corpus
/wiki:inventory add item "TRX-4M ring and pinion" --wiki trx4m-1-18
/wiki:inventory list --view actions --limit 10
/wiki:dataset add "Bitcointalk Temporal Graph" --location https://figshare.com/articles/dataset/BitcoinTemporalGraph/26305093
/wiki:dataset list --view schema --limit 10
/wiki:archive topic old-interest --reason "No longer active"
/wiki:archive list --archived
/wiki:archive restore old-interest
/wiki:compile
/wiki:output report --topic gut-brain
/wiki:assess /path/to/my-app --wiki nutrition
/wiki:lint --fix
How it works
~/wiki/ # Hub — lightweight, no content
├── wikis.json # Registry of all topic wikis
├── _index.md # Lists topic wikis with stats
├── log.md # Global activity log
└── topics/ # Each topic is an isolated wiki
├── nutrition/
│ ├── .obsidian/ # Obsidian vault config
│ ├── inbox/ # Drop zone for this topic
│ ├── inventory/ # Items, candidates, corpora, views
│ ├── datasets/ # Manifests for large/external data
│ ├── raw/ # Immutable sources
│ ├── wiki/ # Compiled articles
│ │ ├── concepts/
│ │ ├── topics/
│ │ └── references/
│ ├── output/ # Generated artifacts
│ ├── _index.md
│ ├── config.md
│ └── log.md
├── woodworking/ # Another topic wiki
└── .archive/ # Archived topic wikis, hidden by default
One topic, one wiki
Each research area is isolated. No cross-topic noise. Queries stay focused. A multi-wiki peek finds overlap when relevant.
[[wikilinks]] for Obsidian plus standard markdown links for everything else. Works in every viewer — including no viewer at all.
Once a source is ingested it is never modified. Articles synthesize on top. Retraction removes both cleanly.
/wiki:collect records aliases, source context, media URLs, cached asset paths, hashes, dedupe notes, scale, and inventory recommendations before anything becomes evidence or durable state.
Parts, source queues, corpora, watch items, and next actions live under inventory/ so they can be listed and revisited without becoming evidence.
datasets/ stores manifests, samples, profiles, and query recipes for large data. The wiki indexes data without copying it into the source corpus.
Archived topics live under topics/.archive/ . Most tools skip them by default; deep queries may surface index hits, and explicit --include-archived can read them.
Runs entirely on the host agent's built-in tools. Plugin is Markdown + commands. No servers, no services, no telemetry.
All commands accept --wiki <name> to target a topic wiki and --local for the project wiki. Archived topic wikis are skipped by default; commands that support --include-archived require that explicit flag before reading or writing archived material. query , output , and plan also accept --with <wiki> for cross-wiki context.
From zero to a compiled wiki in 5 minutes.
claude plugin install wiki@llm-wiki
2. Create a topic wiki
Pick any topic you're curious about:
/wiki init nutrition
This creates a hub at ~/wiki/ and your first topic wiki at ~/wiki/topics/nutrition/ .
/wiki:research "gut microbiome and mental health" --wiki nutrition
# or just say it naturally:
/wiki research gut microbiome and mental health
Five parallel agents search the web from different angles (academic, technical, applied, news, contrarian), ingest the best sources, and compile them into cross-referenced wiki articles. Takes 2-5 minutes.
/wiki:query "how does fiber affect mood?" --wiki nutrition
# or naturally:
/wiki how does fiber affect mood?
The wiki answers from its compiled articles with citations.
5. Audit before you trust an output
/wiki:audit --wiki nutrition
/wiki:audit --artifact output/report-gut-brain.md
Audit rechecks the wiki layer, traces the output's evidence chain, flags drift, and will do fresh research if the local corpus is not enough to answer the trust question.
/wiki:research "topic" --deep — 8 agents instead of 5, adds historical and data angles
/wiki:research "topic" --min-time 1h — keep researching in rounds for an hour
/wiki:research "topic" --plan — decompose into parallel research paths
/wiki:audit --project nutrition-playbook — verify outputs and upstream wiki state together
/wiki add https://example.com/article — fuzzy router detects the URL and ingests it
/wiki what do we know about CRISPR? — fuzzy router detects the question and queries
/wiki:lint --fix — clean up any structural issues
A typical research session flows through four stages:
Stage 1: Ask a question or pick a topic
llm-wiki auto-detects whether you're asking a question or naming a topic. Use the direct command or the fuzzy router:
/wiki:research "What makes long-form articles go viral?" # direct command
/wiki research quantum computing # fuzzy router — same result
/wiki:research --mode thesis "fiber reduces neuroinflammation via SCFAs" # thesis → for/against evidence
Stage 2: Agents search in parallel
5 agents (8 with --deep , 10 with --retardmax ) search simultaneously from different angles — 2-3 web searches each, full-content fetch, quality scoring (1-5). A credibility pass deduplicates before ingestion.
Stage 3: Sources are ingested and compiled
Top sources are saved to raw/ (immutable — never modified after ingestion). Then the compilation pass synthesizes them into wiki articles under wiki/concepts/ , wiki/topics/ , and wiki/references/ with cross-references, confidence scores, and bidirectional links.
Stage 4: Gap report and follow-up
After each round, you see what's covered, what's still missing, and suggested follow-ups. If 2+ gaps remain, you're offered to close them in parallel:

[truncated]
