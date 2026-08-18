---
source: "https://github.com/bleugreen/unigram"
hn_url: "https://news.ycombinator.com/item?id=49349738"
title: "Show HN: Unigram: encode bytes as words that cost one LLM token"
article_title: "GitHub - bleugreen/unigram: A bijective codec between bytes and words that cost exactly one LLM token · GitHub"
image: "https://opengraph.githubassets.com/e560c0172edcb839156ebeffbac73b41a15f8e455ef67f9634c65edaa538c457/bleugreen/unigram"
author: "bleugreenlab"
captured_at: "2026-08-18T18:22:30Z"
capture_tool: "hn-digest"
hn_id: 49349738
score: 2
comments: 0
posted_at: "2026-08-18T17:58:17Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Unigram: encode bytes as words that cost one LLM token

- HN: [49349738](https://news.ycombinator.com/item?id=49349738)
- Source: [github.com](https://github.com/bleugreen/unigram)
- Score: 2
- Comments: 0
- Posted: 2026-08-18T17:58:17Z

## Translation

タイトル: HN を表示: Unigram: バイトを 1 つの LLM トークンを消費するワードとしてエンコードします
記事のタイトル: GitHub - bleugreen/unigram: ちょうど 1 つの LLM トークンを必要とするバイトとワード間の全単射コーデック · GitHub
説明: ちょうど 1 つの LLM トークンを必要とするバイトとワード間の全単射コーデック - bleugreen/unigram

記事本文:
GitHub - bleugreen/unigram: ちょうど 1 つの LLM トークンを必要とするバイトとワードの間の全単射コーデック · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ブルーグリーン
/
ユニグラム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
15 コミット 15 コミット フォルダーとファイル
.github/ workflows .github/ workflows src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md verify-alphabe

t.py verify-alphabet.py verify-claude.py verify-claude.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ちょうど 1 つの LLM トークンを必要とする、バイトとワードの間の全単射コーデック。
貨物追加ユニグラム · crates.io ·
docs.rs · 変更履歴
a14ed61a -> パスワード電子メール共有の構築
8623a771b764ce50bb85371ff65aebe9 -> リンク変更ポイント高ランダムが見つかりました
シーズン イベント リージョン ライト const ケース
ユーザーフィールドテーブルのサポート
識別子は読み取り可能なものになります。大声で言って、部屋中運んでください
または 2 つの窓の間で、隣の窓と一目で区別し、認識します。
再び 1 時間後 — 名前が与える普通のこと。イドは一生を過ごす
プロンプト、ログ、エラー メッセージが調べられます。これにより無料になります。
1 つのワードは 1 バイトと 1 つのトークンであるため、値にはその値とまったく同じ数のトークンがかかります。
バイトを運びます。すべての値に対してフラットで、単語間のスペースにはコストがかかりません。
上記の 4 つのワードは 4 つのトークンで 32 ビットを運びます。 16 は 16 で 128 を運びます。
unigram を使用します:: { UnigramId , CheckedUnigramId } ;
let id : UnigramId < 4 > = UnigramId :: try_random ( ) ? ; // 32 の新しいビット、4 つのトークン
プリントイン！ ( "{id}" ) ; // "パスワードメール共有の構築"
返された = UnigramId :: < 4 > :: parse (& text ) ? ; // 正規: 正確
let salvages = UnigramId :: < 4 > :: Recovery (& text ) ? ; // 寛容: 往復を許容します
// 変更された値が有効な値として渡されない場合の CRC-8 の 1 つの追加ワード。
let checked : CheckedUnigramId < 4 > = CheckedUnigramId :: try_random ( ) ? ;
バイトは値です。単語は、それが表示および解析される方法です。それを抱えて
way は、長さが型の一部であり、等価性がバイト等価性であり、何も存在しないことを意味します。
指定された値がどのような形式であるかという問題 - 文字列形式の API ではできない問題
答えて推測する必要があります。
神父

ee 関数 ( encode 、 decode 、 decode_recovered 、 try_mint ) は、
可変長ペイロード。
parse は正規です。小文字のアルファベットの単語、単一のスペース、その他は何も含まれません。 1 件が受け入れられました
値ごとのスペル。値が信頼される対象となるものです。
Recovery は、モデルの往復処理 (大文字小文字、区切り文字、行) を許容します。
ラッピング。入力全体を読み取るため、最初に候補を分離します。
どちらも未知の単語を拒否し、それに名前を付けます。
1 つのワードは 1 バイトと 1 つのトークンであるため、N バイトの値にはちょうど N 個のトークンがかかります。
あらゆる価値観に対して。クロードの下での平均トークン、最悪の場合は 200 の決定論的ペイロード
括弧内：
括弧内の数値は平均値と同じくらい重要です。他のエンコードごとにコストが変動する
したがって、それに基づいて構築される予算は最悪のケースを想定する必要があります。これは知られています
値が鋳造される前に。
ヘックスは、どこでも、どんな規模でも、どんな家族でも負けます。 GPTの語彙は暗記されています
Base64 フラグメント。4 バイトを超えるとランキングが変更されます — o200k 、base64url の下
フラット 32 に対して 32 バイトで平均 29.5 トークンですが、それでもユニグラムは 4 で勝ちます
バイト (4.0 対 4.5)。 Nonce と correlation-id の幅は、これが構築された目的です。
32 バイトのダイジェストは 224 文字であり、GPT ではトークン マージンが残らないため、適合度はさらに低くなります。
バイトごとに 1 つのトークンがスペース接頭辞付きと bare を保持するため、値のコストはちょうど N になります。
JSON の文字列の先頭、スペースの後、文の途中。唯一の追加料金は、
直前の句読点。理想値 4 に対して 4 バイト値を測定した場合、
256 件のエントリーすべてをオープニングポジションとクロージングポジションでスイープし、最悪の状態を維持しました:
つまり、バイトごとに 1 つのトークンと、その直前の句読点に最大 1 つのトークンが追加されます。
定数、ペイロードに応じてスケーリングしない、コンテキストがで終わる負の値
スペース

e 値が吸収されます。
それはテーブルの特性であり、無料ではありませんでした。 0.2.0 出荷 22 エントリのコスト計算
2 つまたは 3 つのトークンがベアであるため、開始時に評議会コスト N+2 でバリュー オープニングが行われます。
string — そしてその検証者は、冒頭の単語がたまたま安かった 1 つのペイロードをテストしました。
どちらも修正されています。このスイープが、クレームに例外リストを必要としない理由です。
BIP39、Diceware、PGP 単語リスト、および what3words はすべて、このマップおよびすべてのマップよりも古いものです。
データを言葉に。トークナイザーには何も選択されていないことがわかります。 BIP39が最も近い
比較 — 2048 ワード。すべて単一トークンの場合、それぞれ 11 ビットになります。
BIP39 の 2048 のうち 349 のみがフィルターを通過し、クロードは次のバインディング制約になります。
366. 349 を 2 の累乗に切り捨てると、BIP39 由来のエンコーディングが正確に一致します。
256 のエントリとトークンごとに正確に 8 ビット — 同じ密度、以下のリストから
ベアコストや周囲のコンテキストの保証はありません。
BIP39 は別のことを最適化し、それをうまく実行します: ユニークな 4 文字
紙から読み取ったシードフレーズの接頭辞と人間による転写距離。つまり
持つ価値がある。単語に 1 トークンのコストがかかるわけではありません。
トークナイザーの語彙は、正規の単語エントリにスペース接頭辞を付けて保持します。
2 つの単語の間にある単語は、その後に続く単語に吸収され、費用はかかりません。いいえ
その他の区切り文字は無料です。 5 つのファミリーすべてにわたって測定された、8 バイトの値:
結合にはペイロードとほぼ同じコストがかかります。エンコードされた値は引用符内で移動します
実際には文字列。埋め込まれたスペースは自由であり、recover はすべてのスペースを受け入れます。
とにかくこれらの区切り文字を削除するため、別の方法で結合されて戻ってくる値は失われません。
5 つの制約のもとで、4 ～ 10 文字の小文字 ASCII 英語の 256 エントリ:
すべてのトークナイザーの下に、スペース接頭辞が付いたベアの 1 つのトークンが検証者によって固定されます。
OpenAI

の r50k_base 、 p50k_base 、 cl100k_base 、 o200k_base ;の
hf-internal-testing/llama-tokenizer リビジョン d02ad6cb の SentencePiece アーティファクト。
ctok 1.0.0 の「5.0」カウンター、クロードのトークナイザーをオフラインで再構築したもの
Anthropic 独自のものではなく、Anthropic の公式 count_tokens と照合してチェックされます。
すべての 256 エントリの claude-opus-5 上のエンドポイント (スペースがあり、一致する場合)
正確に (verify-claude.py が再実行します)。それらの正確な成果物は主張であって、そうではありません
名前を共有するすべてのモデル、特にトークン化される Llama 3 は除く
ここで確認した SentencePiece モデルではなく tiktoken です。
1 つの文字編集内に 2 つのエントリを含めることはできません。また、接頭辞や接尾辞の派生も含めることはできません
別の。抜けた文字、脱落した接尾語、または完成した単語が着地する
別の有効なエントリではなく、アルファベットの外側にあります。
死、暴力、人種、性別、宗教、政治などは何も起訴されていません。これら
文字列は、トランスクリプト、ログ、およびユーザーが直面するエラーで無断で表面化します。
機能語はありません。それから作られた値は、損傷していると読み取られます
名前としてではなく散文。
凍った。バイト n は ALPHABET[n] で、256 個のスロットすべてが占有されており、
エントリは、以前に発行されたすべての値をデコードする内容を変更します。テストでテーブルを固定する
ダイジェスト。エンコードされた値には、どのテーブルがその値を生成したかが示されていないため、システムは
ストアには、FORMAT_VERSION を一緒に記録する必要があります。
クレートは、実行時またはテスト中において OS CSPRNG のみに依存しており、決して依存しません。
トークン化します。カーゴ テストでは、コーデックとテーブルの構造 (ソート、一意、
長さ、編集距離、プレフィックスとサフィックスの関係、凍結されたダイジェスト、および
チェックワードに対するすべての単一単語置換を徹底的にスイープします。それは言う
コストについては何もありません。
このページのすべての数字は verify-alphabet.py によって出力されます。

アルファベットを読むもの
src/lib.rs から直接取り出して、すべてのエントリを 5 つのファミリーすべてに対して再測定します。
スペース接頭辞があり、裸で、開始と終了を通じて 256 個のエントリすべてをスイープします
すべてのコンテキストの位置 - 依存関係とトークナイザーのリビジョンが固定されている
正確に:
UV 実行 verify-alphabet.py
テーブルを編集した後に実行します。グリーン テスト スイートだけでは、何も確立できません。
この木箱の名前が付けられています。
verify-claude.py は、再構築である 1 つの測定の監査です。
語彙ではなく、クロードのコラムを Anthropic の公式と照らし合わせて再チェックします。
count_tokens エンドポイントを指定し、この 2 つが一致しないエントリを報告します。ニーズ
ANTHROPIC_API_KEY ; --bare を使用した呼び出しは約 300 回。最後はすべてのエントリでクリーンに実行されました。
ちょうど 1 つの LLM トークンを必要とする、バイトとワード間の全単射コーデック
crates.io/crates/unigram リソース
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A bijective codec between bytes and words that cost exactly one LLM token - bleugreen/unigram

GitHub - bleugreen/unigram: A bijective codec between bytes and words that cost exactly one LLM token · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
bleugreen
/
unigram
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
15 Commits 15 Commits Folders and files
.github/ workflows .github/ workflows src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md verify-alphabet.py verify-alphabet.py verify-claude.py verify-claude.py View all files Repository files navigation
A bijective codec between bytes and words that cost exactly one LLM token.
cargo add unigram · crates.io ·
docs.rs · CHANGELOG
a14ed61a -> password email share building
8623a771b764ce50bb85371ff65aebe9 -> links change points high random found
season events region light const case
users field table support
An identifier becomes something you can read. Say it out loud, carry it across a room
or between two windows, tell it apart from its neighbour at a glance, recognise it
again an hour later — the ordinary things a name affords. Ids spend their lives in
prompts, logs, and error messages, being looked at; this makes that free.
One word is one byte and one token, so a value costs exactly as many tokens as it
carries bytes — flat, for every value, with the spaces between words costing nothing.
The four words above carry 32 bits in 4 tokens; the sixteen carry 128 in 16.
use unigram :: { UnigramId , CheckedUnigramId } ;
let id : UnigramId < 4 > = UnigramId :: try_random ( ) ? ; // 32 fresh bits, 4 tokens
println ! ( "{id}" ) ; // "password email share building"
let returned = UnigramId :: < 4 > :: parse ( & text ) ? ; // canonical: exact
let salvaged = UnigramId :: < 4 > :: recover ( & text ) ? ; // tolerant: forgives a round trip
// One extra word of CRC-8, when a mutated value must not pass as a valid one.
let checked : CheckedUnigramId < 4 > = CheckedUnigramId :: try_random ( ) ? ;
The bytes are the value; the words are how it is displayed and parsed. Holding it that
way means the length is part of the type, equality is byte equality, and there is no
question of what format a given value is in — the question a string-shaped API cannot
answer and has to guess at.
Free functions ( encode , decode , decode_recovered , try_mint ) are there for
variable-length payloads.
parse is canonical: lowercase alphabet words, single spaces, nothing else. One accepted
spelling per value, which is what belongs where a value is about to be trusted.
recover forgives what a round trip through a model does — case, separators, line
wrapping. It reads the whole input, so isolate the candidate first.
Both refuse an unknown word and name it.
One word is one byte and one token, so an N-byte value costs exactly N tokens, the same
for every value. Mean tokens under Claude, with the worst of 200 deterministic payloads
in parentheses:
The parenthesised figure matters as much as the mean. Every other encoding's cost swings
with the value, so a budget built on one has to assume its worst case; this one is known
before the value is minted.
Hex loses everywhere, at every size, in every family. The GPT vocabularies have memorised
base64 fragments, which changes that ranking above 4 bytes — under o200k , base64url
averages 29.5 tokens for 32 bytes against a flat 32, while unigram still wins at 4
bytes (4.0 against 4.5). Nonce and correlation-id widths are what this was built for;
a 32-byte digest is a worse fit, at 224 characters and no token margin left under GPT.
One token per byte holds space-prefixed and bare , so a value costs exactly N at
the start of a string, after a space, in JSON, and mid-sentence. The only surcharge is
punctuation immediately before it. Measured for a 4-byte value against an ideal of 4,
sweeping all 256 entries through the opening and closing positions, worst kept:
So: one token per byte, plus at most one for punctuation immediately before it — a
constant, never scaling with the payload, and negative where the context ends in a
space the value absorbs.
That is a property of the table, and it was not free. 0.2.0 shipped 22 entries costing
two or three tokens bare, so a value opening with council cost N+2 at the start of a
string — and its verifier tested one payload whose opening word happened to be cheap.
Both are fixed. The sweep is why the claim needs no exception list.
BIP39, Diceware, the PGP word list, and what3words all predate this and all map
data to words. None was chosen for tokenizers, and it shows. BIP39 is the closest
comparison — 2048 words, which would be 11 bits each if they were all single tokens:
Only 349 of BIP39's 2048 survive the filter, and Claude is the binding constraint at
366. Round 349 down to a power of two and a BIP39-derived encoding lands on exactly
256 entries and exactly 8 bits per token — the same density, from a list that also has
no bare-cost or surrounding-context guarantee.
BIP39 optimises for a different thing, and does it well: unique four-character
prefixes and human-transcription distance, for seed phrases read off paper. That is
worth having. It is not what makes a word cost one token.
Tokenizer vocabularies hold their canonical word entries space-prefixed, so the space
between two words is absorbed into the word that follows it and costs nothing. No
other separator is free. Measured across all five families, an eight-byte value:
The join would cost almost as much as the payload. Encoded values travel inside quoted
strings in practice, where embedded spaces are free — and recover accepts every one
of those separators anyway, so a value that comes back joined differently is not lost.
256 entries of lowercase ASCII English, 4 to 10 characters, under five constraints:
One token, space-prefixed and bare, under every tokenizer the verifier pins:
OpenAI's r50k_base , p50k_base , cl100k_base , o200k_base ; the
hf-internal-testing/llama-tokenizer SentencePiece artifact at revision d02ad6cb ;
and ctok 1.0.0's "5.0" counter, an offline reconstruction of Claude's tokenizer
rather than Anthropic's own — checked against Anthropic's official count_tokens
endpoint on claude-opus-5 for all 256 entries, spaced and bare, where it agrees
exactly ( verify-claude.py reruns it). Those exact artifacts are the claim, not
every model that shares a name, and in particular not Llama 3, which tokenizes with
tiktoken rather than the SentencePiece model checked here.
No two entries within one character edit, and none a prefix or suffix-derivative
of another. A slipped character, a dropped suffix, or a completed word lands
outside the alphabet rather than on a different valid entry.
Nothing charged — no death, violence, race, gender, religion, or politics. These
strings surface unbidden in transcripts, logs, and user-facing errors.
No function words. A value made of that , which , and would reads as damaged
prose rather than as a name.
Frozen. Byte n is ALPHABET[n] , all 256 slots are occupied, and changing an
entry changes what every previously issued value decodes to. A test pins the table's
digest. Nothing in an encoded value says which table produced it, so a system that
stores these must record FORMAT_VERSION alongside them.
The crate depends on nothing but the OS CSPRNG, at runtime or under test, and never
tokenizes. cargo test covers the codec and the table's structure — sorted, unique,
lengths, edit distance, prefix and suffix relationships, the frozen digest, and an
exhaustive sweep of every single-word substitution against the check word. It says
nothing about cost.
Every number on this page is printed by verify-alphabet.py , which reads the alphabet
straight out of src/lib.rs , re-measures every entry against all five families both
space-prefixed and bare, and sweeps all 256 entries through the opening and closing
positions of every context — with dependencies and the tokenizer revision pinned
exactly:
uv run verify-alphabet.py
Run it after any edit to the table. A green test suite alone establishes none of what
this crate is named for.
verify-claude.py is the audit for the one measurement that is a reconstruction rather
than a vocabulary: it re-checks the Claude column against Anthropic's official
count_tokens endpoint and reports any entry where the two disagree. Needs
ANTHROPIC_API_KEY ; roughly 300 calls with --bare . It last ran clean on every entry.
A bijective codec between bytes and words that cost exactly one LLM token
crates.io/crates/unigram Resources
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
