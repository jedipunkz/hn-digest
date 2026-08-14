---
source: "https://pssah4.github.io/vault-operator/guides/capabilities"
hn_url: "https://news.ycombinator.com/item?id=49296964"
title: "Everyone talks about AI agents. This is what one looks from the inside"
article_title: "What Vault Operator can do | Vault Operator"
author: "pssah4"
captured_at: "2026-08-14T10:55:21Z"
capture_tool: "hn-digest"
hn_id: 49296964
score: 1
comments: 0
posted_at: "2026-08-14T10:44:12Z"
tags:
  - hacker-news
  - translated
---

# Everyone talks about AI agents. This is what one looks from the inside

- HN: [49296964](https://news.ycombinator.com/item?id=49296964)
- Source: [pssah4.github.io](https://pssah4.github.io/vault-operator/guides/capabilities)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T10:44:12Z

## Translation

タイトル: AI エージェントについて誰もが話しています。内側から見るとこんな感じです
記事のタイトル: Vault Operator にできること |保管庫オペレーター
説明: Vault Operator が Vault 内で実際に行うことと、Vault Operator が作業にどのように役立つかを簡単に説明します。

記事本文:
コンテンツにスキップ Vault Operator 検索 K メイン ナビゲーション ここから開始 ガイド コンセプト リファレンス 外観について
GitHub ☽ ☼ メニュー トップに戻る サイドバー ナビゲーション 機能
ほとんどの AI ツールはナレッジ ベースの外側にあり、コピー アンド ペーストを待ちます。その中にはVault Operatorが住んでいます。それはあなたのメモを読み、それらを結び付けるグラフに従い、あなたの習慣を見つけて、あなたに代わって行動します。
このページはショートツアーです。各セクションは詳細なガイドにリンクしています。
見るほうがいいですか？この 7 分間のアニメーション解説では、ループ、5 つのゲートを備えたハーネス、ツール、スキル、MCP、そして各レイヤーを横断する 1 つの実際のジョブなど、エージェントを内部から説明します。
エディターでのインライン AI チャット
読んでいるメモ上で直接チャットを実行できます。テキストを選択し、Cmd+Shift+I (Windows および Linux では Ctrl+Shift+I) を押すか、右クリックして [インライン AI チャット] を選択すると、エディタ上にフローティング パネルが開きます。そこから、メモを離れることなく、選択内容について質問したり、調べたり、書き直したり、翻訳したり、要約したり、アクションアイテムを抽出したりできます。
パネルがファイルへの変更を提案すると、Änderungen prüfen モーダルが最初に開きます。各編集を確認し、必要な編集のみを適用すると、エージェントがチャットにチェックポイント マーカーを書き込むため、後で Vault Operator の他の場所で使用されているのと同じチェックポイント システムを通じて変更を元に戻すことができます。
インラインチャットガイド |チャットインターフェース |安全性と制御
ブロックレベルの出自を持つソースをキャプチャ
PDF、Web クリップ、または Office ファイルをチャットにドロップします。 Vault Operator は、Vault のオントロジーに対して高速トリアージを実行し、取り込みを選択した場合は、何かを書き込む前に短いダイアログを案内します。
結果として得られる意味のあるメモには、すべての主張の最後に ↗ リンクが付いています。リンクは次の場所にジャンプします。

ソース内の正確なブロック。主張の出所に戻る道を示さずに「これについてどこかにメモがあります」と言う必要はもうありません。
これは平易な言葉でチャットから実行します。「この PDF を取り込む」は迅速なシングルパス キャプチャを実行し、「この論文の詳細な取り込みを行う」は詳細なパス (トリアージ、トピックの選択、ソース マークアップ、意味のあるメモ、バックリンク) を開始します。 5 段階のディープ フローが各質問で停止し、あなたを待ちます。
プロンプトの例: 「この研究論文を詳しく調べてください。方法論のセクションに焦点を当ててください。」
クイックインジェストチュートリアル |ディープ インジェスト チュートリアル |ナレッジインジェストガイド |ブロックレベルの来歴
Web ページをクリップして永続的なメモにします
エージェントに記事をクリップするように依頼すると、Obsidian Web Clipper と同じようにページがアーカイブされます。つまり、Markdown に変換された全文、ローカル埋め込みに書き換えられたリンクとともにボールトにダウンロードされた画像、およびエージェントがその上に書き込むヘッダー (フロントマターと独自の要約) です。これは、会話にページを読み取り、画像をリモート URL として保持し、長いテキストを切り詰める web_fetch よりもさらに進んでいます。ページ URL とすべての画像 URL は、 web_fetch と同じネットワーク ガード チェーンを渡します。
プロンプトの例: 「この記事を Sources/ にクリップし、ヘッダーに短い概要を書き込みます: https://example.com/article」
ボールトの操作 |安全性と制御
セッション全体にわたる 3 層メモリ
Vault Operator は、1 つのチャット内だけでなく、チャット全体で重要なことを覚えています。
魂は、長期にわたる好み（書き方、プロジェクトの慣例、繰り返しの選択）を保持します。
ファクトには、人、プロジェクト、トピックに関する構造化されたステートメントが含まれます。
履歴は、過去の会話の検索可能な記録です。
新しいチャットが開始されると、エージェントは 3 つのレイヤーすべてからデータを取得します。また、メモをメモリ ソースとしてマークして、その内容がスコープ内に留まるようにすることができます。
プロンプトの例: 「この件について要約してください」

前回と同じようにメモを書きます。」 (好みの形式を記憶します。)
記憶とパーソナライゼーション |メモリの仕組み |マスタリーとレシピ
「X について何を知っていますか?」と尋ねた場合、Vault Operator はファイル名を grep しません。ローカルのセマンティックインデックスを介して意味を検索し、ウィキリンクとフロントマターをたどって、見逃している可能性のあるつながりを明らかにします。
セマンティック インデックスはオプトインです。一度オンにして、いつビルドするかを選択します: ビルドしない (デフォルト)、起動時、またはエージェントの切り替え時。その後、検索はすべてのメモを読んだ図書館員のように動作します。
プロンプトの例: 「行動経済学に関連するメモをすべて見つけて、それらがどのように結びついているかを示してください。」
意味で検索するチュートリアル |知識の発見 |ナレッジレイヤーの仕組み
Word、Excel、PowerPoint ファイルを構築する (ベータ版では PPTX)
Vault Operator は、Vault コンテンツから .docx 、 .xlsx 、および .pptx ファイルを書き込みます。
DOCX および XLSX 出力はクリーンで日常使用に信頼性があります。 PPTX は単一のパイプラインとして実行されます。エージェントは最初に plan_presentation を呼び出してソース ノートを制約されたアウトラインに変換し、次に create_pptx がデックを構築します。 PPTX はベータ版です。クライアント向けのデッキでは、出力を手動で仕上げたドラフトとして扱います。
プロンプトの例: 「このメモを、適切な見出しと目次を備えた Word 文書に変換します。」 「私の会議メモから 5 枚のスライドからなる社内状況プレゼンテーションを作成します。」
Office ドキュメント ガイド (ベータ版の詳細) |オフィスのパイプラインの仕組み
Vault Health を使用して Vault をナビゲートできる状態に保つ
金庫がドリフトする。メモはさまざまなフォルダーに蓄積され、関連するアイデアは相互にリンクされなくなり、タグは分岐します。
Vault Operator はバックグラウンドで暗黙的な接続分析を実行し、意味的に近いがウィキリンクを持たないノート ペアを表示します。ボールトのヘルスチェックはさらに進んでおり、孤立したメモ、壊れたリンク、一貫性のないタグ、紛失したバックにフラグを立てます。

リンク、および「ゴッド ノード」（接続が非常に多く、有用なハブではなくボトルネックになるノート）。また、フォルダーとタグの構造をナレッジ グラフで検出したトピック クラスターと比較します。
プロンプトの例: 「保管庫でヘルスチェックを実行して、修正が必要な点を教えてください。」
ボールトのヘルスチェック |知識の発見
他の AI ツールの MCP サーバーとして実行
Vault Operator には MCP サーバーが付属しています。 Claude Desktop、ChatGPT、Perplexity、Claude Code、および MCP 互換クライアントは、ボールトを読み取り、メモリ層から取得し、会話履歴に追加できます。
すべての外部呼び出しには、source_interface タグが含まれるため、メモリと履歴はサーフェスごとに分離されたままになります。厳密なソース分離はデフォルトではオフになっています。サーフェスを残りのサーフェスから隔てたい場合は、[設定] > [Vault Operator] > [エージェント] > [メモリ] でサーフェスごとにオプトインします。
追加のツールが必要な場合は、Vault Operator を外部 MCP サーバーに向けることもできます。
プロンプトの例: ChatGPT の場合: 「価格設定戦略について Vault Operator が記憶に残した内容を思い出してください。」
コネクタガイド |統合チャットメモリ | MCP アーキテクチャ
インストールされているプラグインを検出して使用する
Vault Operator は起動時にインストールされているプラグインをスキャンし、プラグインごとにスキル ファイルを生成します。 Obsidian コマンドを実行し、プラグイン API を呼び出し、複数のプラグインを 1 つのワークフローに結合できます。
データビュー クエリ、カンバン ボード、テンプレート テンプレート、タスク、Excalidraw: プラグインがインストールされている場合、エージェントはそれを使用できます。
プロンプトの例: 「プロジェクト ノートの未解決タスクからカンバン ボードを作成します。」
スキル、ルール、ワークフロー |プラグイン検出の仕組み
自動承認画面を通じてコントロールを維持する
ファイルを変更するたびに承認が求められます。編集するたびにスナップショットが作成され、ワンクリックで元に戻すことができます。操作ログはすべてのステップを記録します。
アンダーセット

[ings] > [Vault Operator] > [Agents] > [自動承認] どのカテゴリがサイレントに実行され、どのカテゴリが要求を続けるかを決定します:読み取り、書き込み、Web、Vault、プラグイン API 読み取り、プラグイン API 書き込み、レシピ、および MCP 呼び出し。すべてのアクションを表面化したい場合は、どこでも「毎回質問」を維持することもできます。
AI モデルを選択します。何をクラウドに送信するかを決定します。クラウドへの依存をゼロにしたい場合は、Ollama または LM Studio を使用してすべてをローカルで実行します。
当然のことですが、Vault Operator は最初に差分を表示せずにファイルを変更することはできません。
安全性と制御 |ガバナンスの仕組み
エージェントがタスクを正常に完了すると、ツールのシーケンスを記憶します。数回繰り返した後、ヘルパー モデルは一致するレシピから単一の決定論的な実行を計画し、反復推論の大部分をスキップします。同じタスクの LLM 呼び出しは 8 回から 2 回に、トークンの数は数十万から数十に減少します。
複数のトピックにまたがるタスクの場合、Vault Operator はサブエージェントを生成できます。ある人は会議のメモを調べ、別の人は Web を検索できます。 3 つ目は両方を文書に縫い付けます。サブエージェントは独立して実行され、それぞれが独自の会話コンテキストを持ちます。メイン エージェントは結果を収集して結合します。
プロンプトの例: 「価格設定戦略について保管庫に記載されている内容とオンラインの最新市場調査を比較し、推奨事項を書いてください。」
適切なモデルを適切なタイミングで実行します
プロバイダーは一度構成します。 Vault Operator はモデルを検出し、それらを Budget、Main、Frontier に分類し、デフォルトで Main でチャット ループを実行します。 Consult_flagship ツールは、エージェントが困難な場合に 1 つの合成ステップを Frontier にエスカレートし、タスクあたり 3 つの呼び出しと 3000 の出力トークンに制限します。低コストのバックグラウンド作業 (コンテキストの要約、ファストパス計画、プレゼンテーション計画、レシピのプロモーション) は別のヘルプにルーティングされます。

一度選んだrモデル。
モデルの選択 |プロバイダーのリファレンス
Vault Operator 自体は無料でオープンソースです。料金は使用した AI モデルに対してのみお支払いいただきます。
コストを意識したループ (アドバイザー パターン、ヘルパー モデル ルーティング、KV キャッシュ アラインメント、コンテキストの外部化、プロンプトのスリムダウン) により、トークンの使用量が低く抑えられます。単純な検索と要約タスクでは、同じワークロードが約 634,000 トークンから約 60,000 トークンに減少します。
モデルの選択 |トークンの最適化
アパッチ2.0 |インプリント |いかなる保証も責任も負わず、現状のまま提供されます。
いかなる保証も責任も負わず、現状のまま提供されます。

## Original Extract

A quick look at what Vault Operator actually does inside your vault and how it helps you work.

Skip to content Vault Operator Search K Main Navigation Start here Guides Concepts Reference About Appearance
GitHub ☽ ☼ Menu Return to top Sidebar Navigation What it does
Most AI tools sit outside your knowledge base and wait for you to copy and paste. Vault Operator lives inside it. It reads your notes, follows the graph that ties them together, picks up your habits, and acts on your behalf.
This page is the short tour. Each section links out to the guide that goes deep.
Prefer to watch? This 7-minute animated explainer walks through the agent from the inside: the loop, the harness with its five gates, tools, skills, and MCP, then one real job crossing every layer.
Inline AI chat in the editor ​
You can run a chat directly on the note you are reading. Select text, press Cmd+Shift+I (Ctrl+Shift+I on Windows and Linux) or right-click and pick "Inline AI chat", and a floating panel opens over the editor. From there you ask a question about the selection, look something up, rewrite, translate, summarize, or extract action items without leaving the note.
When the panel proposes a change to the file, the Änderungen prüfen modal opens first. You review each edit, apply only the ones you want, and the agent writes a checkpoint marker into the chat so you can undo the change later through the same checkpoint system used elsewhere in Vault Operator.
Inline chat guide | Chat interface | Safety and control
Capture sources with block-level provenance ​
Drop a PDF, web clip, or Office file into the chat. Vault Operator runs a fast triage against your vault's ontology and, if you choose to ingest, walks you through a short dialog before it writes anything.
The resulting sense-making note carries a ↗ link at the end of every claim. The link jumps back to the exact block in the source. No more "I have a note about this somewhere" without a path back to where the claim came from.
You drive this from chat in plain language: "Ingest this PDF" runs the quick single-pass capture, "Do a deep ingest of this paper" starts the deep path (triage, topic pick, source markup, sense-making note, backlinks). The five-step deep flow stops at each question and waits for you.
Example prompt: "Deep-ingest this research paper. Focus on the methodology section."
Quick ingest tutorial | Deep ingest tutorial | Knowledge ingest guide | Block-level provenance
Clip web pages into permanent notes ​
Ask the agent to clip an article and it archives the page the way the Obsidian Web Clipper does: the full text converted to Markdown, the images downloaded into your vault with the links rewritten to local embeds, and a header the agent writes on top (frontmatter plus its own summary). This goes further than web_fetch , which reads a page into the conversation, keeps images as remote URLs, and truncates long text. The page URL and every image URL pass the same network guard chain as web_fetch .
Example prompt: "Clip this article into Sources/ and write a short summary into the header: https://example.com/article "
Vault operations | Safety and control
Three-layer memory across sessions ​
Vault Operator remembers what matters across chats, not only inside one chat.
Soul holds long-lived preferences (writing style, project conventions, recurring choices).
Facts hold structured statements about people, projects, and topics.
History is a searchable transcript of past conversations.
The agent retrieves from all three layers when a new chat starts, and you can mark any note as a memory source so its content stays in scope.
Example prompt: "Summarize this meeting note like last time." (It remembers your preferred format.)
Memory and personalization | How memory works | Mastery and recipes
When you ask "what do I know about X?", Vault Operator does not grep filenames. It searches by meaning over a local semantic index, walks wikilinks and frontmatter, and surfaces connections you may have missed.
The semantic index is opt-in. Turn it on once and pick when it builds: never (default), on startup, or on agent switch. After that, search behaves like a librarian who has read every note.
Example prompt: "Find all notes related to behavioral economics and show me how they connect."
Search by meaning tutorial | Knowledge discovery | How the knowledge layer works
Build Word, Excel, and PowerPoint files (PPTX in beta) ​
Vault Operator writes .docx , .xlsx , and .pptx files from your vault content.
DOCX and XLSX output is clean and reliable for everyday use. PPTX runs as a single pipeline: the agent first calls plan_presentation to turn your source notes into a constrained outline, then create_pptx builds the deck. PPTX is in beta, treat the output as a draft you finish manually for client-facing decks.
Example prompts: "Turn this note into a Word document with proper headings and a table of contents." "Build a five-slide internal status presentation from my meeting notes."
Office documents guide (beta details) | How the office pipeline works
Keep the vault navigable with Vault Health ​
Vaults drift. Notes pile up in different folders, related ideas stop linking to each other, tags fork.
Vault Operator runs implicit connection analysis in the background and surfaces note pairs that are semantically close but have no wikilink. The vault health check goes further: it flags orphaned notes, broken links, inconsistent tags, missing backlinks, and "god nodes" (notes with so many connections they become bottlenecks instead of useful hubs). It also compares your folder and tag structure against the topic clusters it detects in the knowledge graph.
Example prompt: "Run a health check on my vault and tell me what needs fixing."
Vault health check | Knowledge discovery
Run as an MCP server for your other AI tools ​
Vault Operator ships an MCP server. Claude Desktop, ChatGPT, Perplexity, Claude Code, and any MCP-compatible client can read your vault, retrieve from the memory layer, and append to your conversation history.
Every external call carries a source_interface tag, so memory and history stay separable per surface. Strict source isolation is off by default. You opt in per surface under Settings > Vault Operator > Agents > Memory if you want a surface walled off from the rest.
You can also point Vault Operator at external MCP servers when you need extra tools.
Example prompt: In ChatGPT: "Recall what my Vault Operator memory says about pricing strategy."
Connectors guide | Unified chat memory | MCP architecture
Discover and use your installed plugins ​
Vault Operator scans your installed plugins at startup and generates a skill file for each one. It can run Obsidian commands, call plugin APIs, and stitch multiple plugins into one workflow.
Dataview queries, Kanban boards, Templater templates, Tasks, Excalidraw: if you have the plugin installed, the agent can use it.
Example prompt: "Create a Kanban board from the open tasks in my project notes."
Skills, rules, and workflows | How plugin discovery works
Stay in control via the auto-approval surface ​
Every file change asks for your approval. Every edit creates a snapshot you can undo with one click. The operation log records every step.
Under Settings > Vault Operator > Agents > Auto-approve you decide which categories run silently and which keep asking: read, write, web, vault, plugin API reads, plugin API writes, recipes, and MCP calls. You can also keep "ask every time" everywhere if you want every action to surface.
You pick the AI model. You decide what gets sent to the cloud. If you want zero cloud dependency, run everything locally with Ollama or LM Studio.
No surprises: Vault Operator cannot change a file without showing you the diff first.
Safety and control | How governance works
When the agent completes a task successfully, it remembers the tool sequence. After a few repetitions, the helper model plans a single deterministic execution from the matching recipe and skips most of the iterative reasoning. The same task drops from eight LLM calls to two, and from hundreds of thousands of tokens to tens.
For tasks that span multiple topics, Vault Operator can spawn sub-agents. One can research meeting notes while another searches the web. A third stitches both into a document. Sub-agents run in isolation, each with their own conversation context. The main agent collects and combines their results.
Example prompt: "Compare what my vault says about pricing strategy with the latest market research online, then write a recommendation note."
It runs the right model at the right time ​
You configure a provider once. Vault Operator discovers its models, sorts them into Budget, Main, and Frontier, and runs the chat loop on Main by default. The consult_flagship tool escalates one synthesis step to Frontier when the agent struggles, capped at three calls per task and 3000 output tokens. Cheap background work (context condensing, fast-path planning, presentation planning, recipe promotion) routes to a separate helper model you pick once.
Choosing a model | Providers reference
Vault Operator itself is free and open source. You pay only for the AI model you use.
The cost-aware loop (advisor pattern, helper-model routing, KV-cache alignment, context externalization, prompt slim-down) keeps token use low. On simple search-and-summarize tasks the same workload drops from around 634K tokens to around 60K.
Choosing a model | Token optimization
Apache 2.0 | Imprint | Provided as-is, without any warranty or liability.
Provided as-is, without any warranty or liability.
