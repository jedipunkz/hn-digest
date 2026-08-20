---
source: "https://tablecanon.app/"
hn_url: "https://news.ycombinator.com/item?id=49376991"
title: "Show HN: Building Table Canon, an AI Campaign Memory Engine for TTRPGs"
article_title: "Table Canon — TTRPG campaign memory for game masters"
image: "https://tablecanon.app/og-image.png"
author: "schillingderek"
captured_at: "2026-08-20T17:20:46Z"
capture_tool: "hn-digest"
hn_id: 49376991
score: 2
comments: 1
posted_at: "2026-08-20T16:40:43Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Building Table Canon, an AI Campaign Memory Engine for TTRPGs

- HN: [49376991](https://news.ycombinator.com/item?id=49376991)
- Source: [tablecanon.app](https://tablecanon.app/)
- Score: 2
- Comments: 1
- Posted: 2026-08-20T16:40:43Z

## Translation

タイトル: Show HN: Building Table Canon、TTRPG 用 AI キャンペーン メモリ エンジン
記事タイトル: Table Canon — ゲームマスターのためのTTRPGキャンペーン記念
説明: あなたのテーブル、覚えています。 Table Canon は、セッション音声を永続的なキャンペーンの記憶に変えます。つまり、テーブルトップ RPG 全体での総括、世界資料、そして「Ask」です。
HN テキスト: こんにちは、HN!私は、プレイグループが常に遭遇し続けていた問題を解決するために Table Canon を構築しました。3 ～ 4 時間のテーブルトップ ゲーム セッションでは大量の音声録音が残りますが、標準的な会議メモ作成者はすべてのセッションを孤立した孤島、肉屋の空想用語のように扱い、誰が話しているのかわかりません。数か月にわたるゲームにわたって長期的な状態を追跡するエンジンが必要だったので、セッション全体でエンティティの更新、オープン クエスト フック、およびキャラクター プロミスを抽出するパイプラインを構築しました。技術スタック: * 転写: Whisper-large-v3-turbo
* ダイアライゼーション: スピーカーの埋め込みと音声プロファイルのマッチングのための pyannote
抽出とメモリ: 構造化出力を備えた OpenAI API (状態更新のための JSON スキーマの適用)
* TTS & オーディオの要約: Kokoro / Chatterbox Turbo
音楽生成: セッションの概要を歌詞/バラードにレンダリングするための ACE-Step-v1.5-XL-Turbo いくつかのエンジニアリングのレッスンと課題: * 状態デルタ抽出とコンテキストの爆発: 20 個の以前のセッションのトランスクリプトをコンテキスト ウィンドウにフィードすると、すぐにコストが高くなり、ノイズが多くなります。生の履歴を再読み取りする代わりに、各セッションはアトミック状態デルタ (NPC ドシエの更新、新しい場所、解決された約束) をデータベースに出力します。キャンペーンがセッション 30 を超えて拡大するときにコンテキストを制限し続けることは、アーキテクチャ上の最も難しいハードルの 1 つです。
* カスタムプレ辞書: 一般的な STT モデルは、自作の固有名詞 (空想の名前を標準の辞書の単語に変換する) に苦労します。プリパスファンタジー用語辞書をプロンプトコンに挿入する

テキストの初回パスのスペルが大幅に改善されました。
* VAD とオーディオ チャンキング: 4 時間の生のオーディオ ファイルを Pyannote/Whisper に直接渡すと、メモリ リークとプロセス タイムアウトが発生します。モデルに触れる前に、音声アクティビティ検出 (VAD) と確定的チャンキングによる前処理が必要でした。現在の制限事項と進行中の困難な問題: * エンティティ エイリアスの解決: プレイヤーがさまざまなエイリアスまたは非公式の略記 (例: 「レッド ビショップ」を「アーサー」または「あのカルト リーダーの男」に一致させる) を使用する場合、異なる NPC を誤ってマージすることなく、セッション全体でエンティティを一致させます。私はこれに部分的に対処しますが、ユーザーが事後にエイリアスを編集したり、エンティティを結合または分割したりできるようにします。
* クエストとフックの解決ロジック: LLM を微調整して、約束、未解決の謎、またはクエストが実際に解決されたのか、未解決のままか暗黙的に放棄されたのかを確実に判断します。他の人がこの種の問題にどのように対処しているか、またはそれを試している人へのメモについてのフィードバックをお待ちしています。初回ログインは不要で、6 時間のアップロードを試すことができます。

記事本文:
Table Canon — ゲームマスター向け TTRPG キャンペーン記念

## Original Extract

Your table, remembered. Table Canon turns session audio into durable campaign memory — recaps, world dossiers, and Ask across your tabletop RPG.

Hey HN! I built Table Canon to solve a problem my playgroup kept running into: 3-4 hour tabletop gaming sessions leave behind massive audio recordings, but standard meeting note-takers treat every session as an isolated island, butcher fantasy terms, and don't know who is speaking. I wanted an engine that tracks long-term state across months of games, so I built a pipeline to extract entity updates, open quest hooks, and character promises across sessions. The Tech Stack: * Transcription: whisper-large-v3-turbo
* Diarization: pyannote for speaker embeddings & voice profile matching
Extraction & Memory: OpenAI API with Structured Outputs (JSON Schema enforcement for state updates)
* TTS & Audio Recaps: Kokoro / Chatterbox Turbo
Music Generation: ACE-Step-v1.5-XL-Turbo for rendering session summaries into lyrics/ballads A Few Engineering Lessons & Challenges: * State Delta Extraction vs. Context Explosions: Feeding 20 prior session transcripts into context windows quickly becomes cost-prohibitive and noisy. Instead of re-reading raw history, each session outputs an atomic state delta (updates to NPC dossiers, new locations, resolved promises) to a database. Keeping context bounded as campaigns stretch past session 30+ has been one of the trickiest architectural hurdles.
* Custom Pre-Lexicons: General STT models struggle with homebrew proper nouns (turning fantasy names into standard dictionary words). Injecting a pre-pass fantasy term dictionary into prompt context significantly improved first-pass spelling.
* VAD & Audio Chunking: Passing a 4-hour raw audio file directly to Pyannote/Whisper leads to memory leaks and process timeouts. Pre-processing with Voice Activity Detection (VAD) and deterministic chunking was necessary before touching the models. Current Limitations & Active Hard Problems: * Entity Alias Resolution: Matching entities across sessions when players use varying aliases or informal shorthand (e.g., matching "The Red Bishop" to "Arthur" or "that cult leader guy") without accidentally merging distinct NPCs. I address this, partially, but allowing the user to Edit aliases, merge or split entities after-the-fact.
* Quest & Hook Resolution Logic: Fine-tuning the LLM to reliably determine whether a promise, open mystery, or quest has actually been resolved versus remaining open or implicitly abandoned. I'd love feedback on how others are handling these sorts of issues - or any notes for folks who try it out! No initial login required with 6 hours of upload available to try.

Table Canon — TTRPG campaign memory for game masters
