---
source: "https://github.com/marcsnid/steganeur"
hn_url: "https://news.ycombinator.com/item?id=49390543"
title: "Show HN: Steganeur – Hide secret messages in LLM-generated text (Rust)"
article_title: "GitHub - marcsnid/steganeur: Hide secret messages inside LLM-generated text. Neural linguistic steganography in Rust with multiple methods. · GitHub"
image: "https://opengraph.githubassets.com/60355e4f51767fa58d1982da8a94473c8f8d3ab8d3973c0e0b3bceb32f73f47a/marcsnid/steganeur"
author: "marcsnid"
captured_at: "2026-08-21T17:20:16Z"
capture_tool: "hn-digest"
hn_id: 49390543
score: 3
comments: 0
posted_at: "2026-08-21T16:33:37Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Steganeur – Hide secret messages in LLM-generated text (Rust)

- HN: [49390543](https://news.ycombinator.com/item?id=49390543)
- Source: [github.com](https://github.com/marcsnid/steganeur)
- Score: 3
- Comments: 0
- Posted: 2026-08-21T16:33:37Z

## Translation

タイトル: HN を表示: Steganeur – LLM で生成されたテキスト内の秘密メッセージを非表示にする (Rust)
記事のタイトル: GitHub - marcsnid/steganeur: LLM で生成されたテキスト内の秘密メッセージを非表示にします。複数の方法を使用した Rust のニューラル言語ステガノグラフィー。 · GitHub
説明: LLM で生成されたテキスト内の秘密メッセージを非表示にします。複数の方法を使用した Rust のニューラル言語ステガノグラフィー。 - マークスニッド/ステガヌール
HN テキスト: こんにちは、HN!私は本当に楽しいプロジェクトを作りました。そのアイデアと実装を誇りに思っています。これが最初に思いついたのは、友人と話しているときでした。自然テキストでステガノグラフィーを行う方法について話していたところ、LLM のトークンの選択が自然なチャネルであることに気づきました。 Steganeur は、秘密のメッセージを言語モデルのトークン選択肢にエンコードし、通常の文のように読めるカバー テキストを生成することでこれを実現します。受信者は、カバー テキストとモデルのみを使用してメッセージを復元します。拒否方法を使用すると、出力は単に見た目が似ているだけでなく、統計的に通常の生成と同一になります。私が遭遇した興味深い問題: steganeur はモデルの logprob を直接読み取りますが、4 つのメソッドのうち 3 つはエンコードとデコードでまったく同じ結果が得られる必要があります。 GPU では削減は非決定的であるため、logprob は実行間で変動し、ビットが破損するのに十分です。確率ではなく ID によってトークンをビン化するため、GPU 上で生き残るメソッド (ブロック) は 1 つだけです。したがって、ブロックはどのサーバーでも動作しますが、他の 3 つは CPU を必要とします。 Rust で書かれ、デュアル ライセンスの MIT/Apache-2.0 で、llama.cpp をターゲットとします (サーバーはトークン ID フィールドを含む top_logprobs を返す必要があります)

記事本文:
GitHub - marcsnid/steganeur: LLM で生成されたテキスト内の秘密メッセージを非表示にします。複数の方法を使用した Rust のニューラル言語ステガノグラフィー。 · GitHub
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
マークスニド
/
ステガヌール
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
4 コミット 4 コミット フォルダーとファイル
例 例 src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md V

すべてのファイルを表示 リポジトリ ファイルのナビゲーション
言語モデルによって生成された自然なテキスト内に秘密のメッセージを隠します。
受信者は、カバー テキストと同じモデルのみを使用してメッセージを復元します。
「従来の暗号化は秘密メッセージを理解できない形式に暗号化しますが、ステガノグラフィーは秘密メッセージをカバー信号にエンコードすることで通信が行われていることを隠蔽します。」
ジーグラー、デン、ラッシュ (2019)
Steganeur は、ニューラル言語ステガノグラフィーの Rust 実装です。秘密
メッセージは言語モデルのトークン選択肢にエンコードされ、カバーが生成されます。
通常の散文のように読めるテキスト。フォーラム上の文章を見た第三者が、
電子メールやブログでは、隠されたメッセージが含まれていることが分かりません。受信者
モデルとコミュニティのデフォルト設定を知っている人がテキストをフィードバックします
ステガヌールを介してメッセージを回復します。
あなたはステガヌールに秘密のメッセージとコンテキストプロンプトを与えます。カバーを生み出します
プロンプトを自然に継続するテキスト。メッセージは
トークンの選択肢。受信者はカバー テキスト (およびカバー テキストのみ) を返送します。
同じモデル設定のステガヌアに送信すると、メッセージが返されます。
エコー「正午に会いましょう」\
|ステガヌールエンコード\
--メソッド拒否 --拒否ビット 2 --シード 42 \
--温度 1.0 --top-k 300 \
--context " 古い灯台は崖の端に建っており、その光は毎晩暗い海を横切っていました。" \
--llama-url " http://127.0.0.1:11434 " \
--モデル「Qwen3.6-27B-GGUF」\
--vocab-size 152064 --eos-token 151643
出力 (表紙テキスト):
古い灯台は崖の端に建っており、その光は海を横切っていた。
毎晩暗い海。しかし、何かが変わりました -- 警告灯は点滅しませんでした。
そして航海日誌を開いたままにしていた管理人もいなかった。嵐のせいではなく、光は見えなくなったが、
なぜなら、

供給元のシステムは、上で実行されているマシンのように、静かに障害を起こしていました。
代わりに牙燃料。霧が濃く感じられました。近海は空の下を滑り、
手付かずの。
デコード (カバーテキストのみ、同じモデル設定):
猫カバー.txt \
|ステガヌールのデコード \
--メソッド拒否 --拒否ビット 2 \
--温度 1.0 --top-k 300 \
--context " 古い灯台は崖の端に建っており、その光は毎晩暗い海を横切っていました。" \
--llama-url " http://127.0.0.1:11434 " \
--モデル「Qwen3.6-27B-GGUF」\
--vocab-size 152064 --eos-token 151643
# 出力: 正午に会いましょう
トークン ファイルもサイド チャネルもありません。カバーテキストがチャンネルです。
カーゴビルド --release
2. LLM サーバーを起動します
完全品質のメソッド (算術、リジェクション、ハフマン) の場合は、CPU 上で実行されるため、
logprob は決定的です (理由については、以下の「課題」を参照してください)。
ラマサーバー -m モデル.gguf --ホスト 0.0.0.0 --ポート 11434 \
--n-gpu-layers 0 \
--キャッシュ タイプ-k f16 --キャッシュ タイプ-v f16 \
--フラッシュアトネーションオフ
ブロック (GPU の非決定性を許容する) の場合、任意のサーバー構成が機能します。
上記の「機能」の例を使用してください。どの方法については、以下の「方法」を参照してください。
それぞれが必要とするフラグを選択します。
Steganeur は 4 つのエンコード方式をサポートしています。それぞれのカバーには異なるトレードオフがあります
品質、ビット密度、サーバー要件。
拒絶 (Cachin、2004) -- 推奨
モデルの分布は 2^b の等確率質量ビンに分割されます。
メッセージ ビットはターゲット ビンを選択します。エンコーダーは、以下からランダムなサンプルを抽出します。
ディストリビューション (シードされた RNG を使用してローカルで、ごとに追加のモデル呼び出しはありません)
拒否)、ターゲット ビンに該当する場合にのみトークンを発行し、再描画します
それ以外の場合は。
エコー「正午に会いましょう」\
|ステガヌールエンコード --method 拒否 --rejection-bits 2 --seed 42 \
--温度 1.0 --top-k 300 \
--context " 古い灯台

崖の端に立っていて、その光は毎晩暗い水面を横切っていた。 " \
--llama-url " http://127.0.0.1:11434 " \
--モデル「Qwen3.6-27B-GGUF」\
--vocab-size 152064 --eos-token 151643
カバーテキスト (実際の例):
古い灯台は崖の端に建っており、その光は海を横切っていた。
毎晩暗い海。しかし、何かが変わりました -- 警告灯は点滅しませんでした。
そして航海日誌を開いたままにしていた管理人もいなかった。嵐のせいではなく、光は見えなくなったが、
なぜなら、それを供給するシステムは、上で実行されているマシンのように、静かに障害を起こしていたからです。
代わりに牙燃料。霧が濃く感じられました。近海は空の下を滑り、
手付かずの。
長所: モデルの分布からの KL 発散がまったくゼロです。出力は
統計的には通常世代と同一であり、これが最も強力です
検出不可能であると主張する可能性があります。カバーの品質は 4 つの方法の中で最高です。
短所: 決定論的なサーバー (CPU) が必要です。ビンの境界は、
累積確率分布なので、logprob ドリフトによってそれらがシフトされます。エラーは発生します
カスケードではありません (各トークンのビットは独立しています) が、ビット自体はカスケードできます。
ドリフトでフリップします。
算術 (Ziegler、Deng、Rush、2019)
メッセージは 2 進数の小数として解釈されます。各ステップで、モデルの
配布パーティション [0, 1)。ビンにメッセージが含まれているトークン
端数が出て間隔が狭くなります。これにより、可変ビット密度が得られます
(トークンあたり 1 ～ 8 ビット)。
エコー「正午に会いましょう」\
|ステガヌール エンコード --メソッド算術 \
--温度 2.0 --top-k 300 \
--context " 彼女は古い手紙を慎重に開きましたが、紙は経年劣化で黄ばんでいました。 " \
--llama-url " http://127.0.0.1:11434 " \
--モデル「Qwen3.6-27B-GGUF」\
--vocab-size 152064 --eos-token 151643
カバーテキスト (実際の例):
彼女は古い手紙を注意深く開いたが、紙は経年劣化で黄ばんでいた。 「ジェニファーへ
の

カルパチア」と右上隅に書かれていました。とても特徴的で似た名前です。
純粋なアンティークから。
彼女はその手紙を読みました、
長所: 拒否と同時に最高のカバー品質。可変ビット密度は高いことを意味します
フラット配布の場合のトークンあたりの容量。
短所: 決定論的なサーバー (CPU) が必要です。 logprob からの 1 つのビン フリップ
ドリフトは間隔状態を通じてカスケードし、後続のすべてのビットを破壊します。
算術演算は、非決定論のもとでは最も脆弱な方法です。を使用します。
--arith-block-size フラグを使用して間隔を定期的にリセットし、制限します
単一ブロックへのカスケード ダメージ (回復するには --ecc と組み合わせます)。
長さ制限のあるハフマン ツリー (パッケージ マージ アルゴリズムで構築) は次のとおりです。
各ステップでのモデルの分布から構築されます。メッセージビットのトラバース
ツリーを使用してトークンを選択します。確率が高いトークンは、消費するビットが少なくなります。
エコー「正午に会いましょう」\
|ステガヌール エンコード --method huffman \
--温度 2.0 --top-k 300 \
--context " 紅葉が流れてきた " \
--llama-url " http://127.0.0.1:11434 " \
--モデル「Qwen3.6-27B-GGUF」\
--vocab-size 152064 --eos-token 151643
カバーテキスト (実際の例):
秋の紅葉は爽やかな輝きを放ち、
明るいインドの９月。感謝祭の匂いが漂ってくるのは明らかだった
で
長所: カバーの品質が良い。可変ビット密度は分布形状に適応します。
短所: 決定論的なサーバー (CPU) が必要です。ツリー構造は後でドリフトする可能性があります
算術と同様に、誤って割り当てられたトークン。
語彙は、固定ハッシュを介して 2^b 個のビンに分割されます。メッセージが分割されています
b ビットのチャンクに分割します。各チャンクについて、その中で最も確率の高いトークンが、
対応するビンが出力されます。各トークンは正確に b ビットをエンコードします。
エコー「正午に会いましょう」\
|ステガヌールエンコード --method block --block-bits 2 \
--温度 2.0 --top-k 300 \
--context " 彼女は歩きます

森を通って " \
--llama-url " http://127.0.0.1:11434 " \
--モデル「Qwen3.6-27B-GGUF」\
--vocab-size 152064 --eos-token 151643
カバーテキスト (実際の例):
彼女は誰にも気づかれずに森を歩きました。シンプルな散歩、シンプルな
目標と単純な楽しみ: 葉を吹き抜ける風、木々に差し込む陽光
枝の隙間、ベルベットの枕のように木の根元を覆う苔、そしてアブラナの花
あらゆる形が、この森の周りのあらゆる道や空き地を愛でています。
たくさんの道。
彼女のペースは相変わらずカジュアルなままだった
長所: 非決定的な GPU サーバー (ハッシュ
ビンは logprob ドリフトを無視します)。各トークンには固定数のビットが含まれるため、
デコードはシンプルかつ堅牢です。
短所: カバーの品質は 4 つの方法の中で最も低くなります。ターゲットのビンに到達するには、
エンコーダは、異常なトークンまたは確率の低いトークンの発行を強制される可能性があります。
散文が少しずれているように見えます。
方法
トークンあたりのビット数
カバーの品質
確定的サーバー (CPU)
非決定的サーバー (GPU)
担保請求
拒否
固定 (例: 2)
ベスト
信頼できる
壊れやすい（ビンシフト）
まさにゼロKL
算数
変数 (1 ～ 8)
ベスト
信頼できる
壊れやすい（カスケード）
ハフマン
変数
良い
信頼できる
壊れやすい（木の漂流）
ブロック
固定 (例: 2)
中等度
信頼できる
信頼性が高い (ハッシュ ビンはドリフトを無視します)
推奨事項:
決定論的 (CPU) サーバー: 最高のカバー品質を得るには拒否を使用します
ゼロ KL の検出不可能性、または最大ビット密度の算術。
非決定的 (GPU) サーバーの場合: --ecc を指定してブロックを使用します。そのハッシュ
ビンは logprob ドリフトを無視します。
不明なサーバー間で運用環境で使用する場合は、コミュニティのデフォルトを次のように文書化します。
決定的な (CPU) サーバー構成。
GF(2^8) 上のリードソロモン層は、次のビットエラーを許容するために利用できます。
logprob ドリフトまたはその他のノイズ。エンコード前にメッセージバイトをラップし、
デコーディ後にアンラップされる

NG。
エコー「正午に会いましょう」\
|ステガヌール エンコード --ecc --ecc-parity 10 --method block --block-bits 2 \
--温度 2.0 --top-k 300 \
--context " 彼女は森を歩いた " \
--llama-url " http://127.0.0.1:11434 " \
--モデル「Qwen3.6-27B-GGUF」\
--vocab-size 152064 --eos-token 151643
猫カバー.txt \
|ステガヌール デコード --ecc --ecc-parity 10 --method block --block-bits 2 \
--温度 2.0 --top-k 300 \
--context " 彼女は森を歩いた " \
--llama-url " http://127.0.0.1:11434 " \
--モデル「Qwen3.6-27B-GGUF」\
--vocab-size 152064 --eos-token 151643
--ecc-parity N は、ペイロードに N パリティ バイトを追加します。デコーダで修正できる
フロア(N/2)バイトエラーまで。 ECC は、エラーが発生するメソッドで最も役立ちます。
局所的 (ブロック、拒否) であり、算術の効率が低くなります。
ビン フリップは無制限のバーストにカスケードされます。
温度は、各トークンが伝送できるビット数を制御します。より高い温度
確率質量を分散し、メッセージ ビットのためのスペースを増やします。
演算方式は2.0以上を使用してください。ブロック方式は少ない
温度に敏感。
エコー「こんにちは」 | steganeur encode --context " 何らかのコンテキスト " --dummy
ダミー LM にはトークン文字列がないため、テキストベースのデコードは行われません。

[切り捨てられた]

## Original Extract

Hide secret messages inside LLM-generated text. Neural linguistic steganography in Rust with multiple methods. - marcsnid/steganeur

Hi HN! I made a project that I found really fun and I'm proud of the idea and implementation. This first came up when speaking to a friend, we were talking around how to do steganography in natural text, and realized an LLM's token choices are a natural channel. Steganeur accomplishes this by encoding a secret message into the token choices of a language model, producing cover text that reads like a normal sentence. A recipient recovers the message using only the cover text and the model. With the rejection method, the output is statistically identical to normal generation, not just something that looks similar. The interesting problem I hit: steganeur reads the model's logprobs directly, and three of the four methods need them to come out exactly the same on encode and decode. On GPU reductions are non-deterministic, so the logprobs drift between runs, which is enough to corrupt the bits. Only one method (block) survives on GPU, because it bins tokens by their id rather than by probability. So block works on any server, the other three need CPU. Written in Rust, dual-licensed MIT/Apache-2.0, targets llama.cpp (the server must return top_logprobs with token id fields)

GitHub - marcsnid/steganeur: Hide secret messages inside LLM-generated text. Neural linguistic steganography in Rust with multiple methods. · GitHub
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
marcsnid
/
steganeur
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
examples examples src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md View all files Repository files navigation
Hide secret messages inside natural-looking text generated by a language model.
A recipient recovers the message using only the cover text and the same model.
"Whereas traditional cryptography encrypts a secret message into an unintelligible form, steganography conceals that communication is taking place by encoding a secret message into a cover signal."
Ziegler, Deng, Rush (2019)
Steganeur is a Rust implementation of neural linguistic steganography. A secret
message is encoded into the token choices of a language model, producing cover
text that reads like normal prose. A third party who sees the text on a forum,
in an email, or on a blog cannot tell it carries a hidden message. A recipient
who knows the model and the community's default settings feeds the text back
through steganeur and recovers the message.
You give steganeur a secret message and a context prompt. It produces cover
text that continues the prompt naturally, with the message encoded into the
token choices. The recipient gives the cover text (and only the cover text) back
to steganeur with the same model settings, and gets the message back.
echo " Meet me at noon " \
| steganeur encode \
--method rejection --rejection-bits 2 --seed 42 \
--temperature 1.0 --top-k 300 \
--context " The old lighthouse stood at the edge of the cliff, its beam sweeping across the dark waters each night. " \
--llama-url " http://127.0.0.1:11434 " \
--model " Qwen3.6-27B-GGUF " \
--vocab-size 152064 --eos-token 151643
Output (the cover text):
The old lighthouse stood at the edge of the cliff, its beam sweeping across the
dark waters each night. But something changed -- no warning lights blinked on,
and no keeper kept the logbook open. The beam was blind, not from storm, but
because the system that fed it was failing silently, like a machine running on
fang fuel instead. The fog felt thick. Nearshore waters glided under the sky,
untouched.
Decode (from the cover text alone, same model settings):
cat cover.txt \
| steganeur decode \
--method rejection --rejection-bits 2 \
--temperature 1.0 --top-k 300 \
--context " The old lighthouse stood at the edge of the cliff, its beam sweeping across the dark waters each night. " \
--llama-url " http://127.0.0.1:11434 " \
--model " Qwen3.6-27B-GGUF " \
--vocab-size 152064 --eos-token 151643
# Output: Meet me at noon
No tokens file, no side channel. The cover text is the channel.
cargo build --release
2. Start your LLM server
For the full-quality methods (arithmetic, rejection, huffman), run on CPU so
the logprobs are deterministic (see "Challenges" below for why):
llama-server -m model.gguf --host 0.0.0.0 --port 11434 \
--n-gpu-layers 0 \
--cache-type-k f16 --cache-type-v f16 \
--flash-attn off
For block (which tolerates GPU non-determinism), any server config works.
Use the examples in "What it does" above. See "Methods" below for which method
to pick and the flags each one requires.
Steganeur supports four encoding methods. Each has different trade-offs in cover
quality, bit density, and server requirements.
Rejection (Cachin, 2004) -- recommended
The model's distribution is partitioned into 2^b equal-probability-mass bins.
The message bits select a target bin. The encoder draws a random sample from
the distribution (locally, with a seeded RNG, no extra model calls per
rejection) and emits the token only if it falls in the target bin, redrawing
otherwise.
echo " Meet me at noon " \
| steganeur encode --method rejection --rejection-bits 2 --seed 42 \
--temperature 1.0 --top-k 300 \
--context " The old lighthouse stood at the edge of the cliff, its beam sweeping across the dark waters each night. " \
--llama-url " http://127.0.0.1:11434 " \
--model " Qwen3.6-27B-GGUF " \
--vocab-size 152064 --eos-token 151643
Cover text (real example):
The old lighthouse stood at the edge of the cliff, its beam sweeping across the
dark waters each night. But something changed -- no warning lights blinked on,
and no keeper kept the logbook open. The beam was blind, not from storm, but
because the system that fed it was failing silently, like a machine running on
fang fuel instead. The fog felt thick. Nearshore waters glided under the sky,
untouched.
Pros: exactly zero KL divergence from the model's distribution. The output is
statistically identical to normal generation, which is the strongest
undetectability claim possible. Cover quality is the best of the four methods.
Cons: requires a deterministic server (CPU). The bin boundaries depend on the
cumulative probability distribution, so logprob drift shifts them. Errors do
not cascade (each token's bits are independent), but the bits themselves can
flip under drift.
Arithmetic (Ziegler, Deng, Rush, 2019)
The message is interpreted as a binary fraction. At each step, the model's
distribution partitions [0, 1). The token whose bin contains the message
fraction is emitted, and the interval narrows. This gives variable bit density
(1 to 8 bits per token).
echo " Meet me at noon " \
| steganeur encode --method arithmetic \
--temperature 2.0 --top-k 300 \
--context " She opened the old letter carefully, the paper yellowed with age. " \
--llama-url " http://127.0.0.1:11434 " \
--model " Qwen3.6-27B-GGUF " \
--vocab-size 152064 --eos-token 151643
Cover text (real example):
She opened the old letter carefully, the paper yellowed with age. "To Jennifer
of Carpathia," it read on the upper-right corner. A name so distinct and like
from pure antiquum.
She read the letter,
Pros: best cover quality alongside rejection. Variable bit density means high
capacity per token for flat distributions.
Cons: requires a deterministic server (CPU). A single bin flip from logprob
drift cascades through the interval state and corrupts all subsequent bits, so
arithmetic is the most fragile method under non-determinism. Use the
--arith-block-size flag to reset the interval periodically, which limits
cascade damage to a single block (pair with --ecc to recover).
A length-limited Huffman tree (built with the package-merge algorithm) is
constructed from the model's distribution at each step. Message bits traverse
the tree to select tokens. Higher-probability tokens consume fewer bits.
echo " Meet me at noon " \
| steganeur encode --method huffman \
--temperature 2.0 --top-k 300 \
--context " The autumn leaves drifted down " \
--llama-url " http://127.0.0.1:11434 " \
--model " Qwen3.6-27B-GGUF " \
--vocab-size 152064 --eos-token 151643
Cover text (real example):
The autumn leaves drifted down in crisp glory upon a
bright Indian September. It was clear you could smell Thanksgiving drifting forward
in the
Pros: good cover quality. Variable bit density adapts to the distribution shape.
Cons: requires a deterministic server (CPU). The tree structure can drift after
a mis-assigned token, similar to arithmetic.
The vocabulary is split into 2^b bins via a fixed hash. The message is split
into b-bit chunks. For each chunk, the highest-probability token in the
corresponding bin is emitted. Each token encodes exactly b bits.
echo " Meet me at noon " \
| steganeur encode --method block --block-bits 2 \
--temperature 2.0 --top-k 300 \
--context " She walked through the forest " \
--llama-url " http://127.0.0.1:11434 " \
--model " Qwen3.6-27B-GGUF " \
--vocab-size 152064 --eos-token 151643
Cover text (real example):
She walked through the forest with no one in mind. A simple walk, with simple
goals and simple pleasures: the wind through leaves, sunlight breaking in the
gaps in branches, moss covering tree roots like velvet pillows and flowers of
all shapes adoring every path and clearing in the area around this forest's
many trails.
Her pace remained casual for
Pros: works on any server, including non-deterministic GPU servers (the hash
bins ignore logprob drift). Each token carries a fixed number of bits, so
decode is simple and robust.
Cons: cover quality is the lowest of the four methods. To hit a target bin, the
encoder may be forced to emit an unusual or low-probability token, which can
make the prose look slightly off.
Method
Bits per token
Cover quality
Deterministic server (CPU)
Non-deterministic server (GPU)
Security claim
Rejection
Fixed (e.g. 2)
Best
Reliable
Fragile (bin shift)
Exactly zero KL
Arithmetic
Variable (1-8)
Best
Reliable
Fragile (cascade)
Huffman
Variable
Good
Reliable
Fragile (tree drift)
Block
Fixed (e.g. 2)
Moderate
Reliable
Reliable (hash bins ignore drift)
Recommendation:
On a deterministic (CPU) server: use rejection for the best cover quality
and zero-KL undetectability, or arithmetic for maximum bit density.
On a non-deterministic (GPU) server: use block with --ecc . Its hash
bins ignore logprob drift.
For production use across unknown servers, document the community default as
a deterministic (CPU) server configuration.
A Reed-Solomon layer over GF(2^8) is available to tolerate bit errors from
logprob drift or other noise. It wraps the message bytes before encoding and
unwraps after decoding.
echo " Meet me at noon " \
| steganeur encode --ecc --ecc-parity 10 --method block --block-bits 2 \
--temperature 2.0 --top-k 300 \
--context " She walked through the forest " \
--llama-url " http://127.0.0.1:11434 " \
--model " Qwen3.6-27B-GGUF " \
--vocab-size 152064 --eos-token 151643
cat cover.txt \
| steganeur decode --ecc --ecc-parity 10 --method block --block-bits 2 \
--temperature 2.0 --top-k 300 \
--context " She walked through the forest " \
--llama-url " http://127.0.0.1:11434 " \
--model " Qwen3.6-27B-GGUF " \
--vocab-size 152064 --eos-token 151643
--ecc-parity N adds N parity bytes to the payload. The decoder can correct up
to floor(N/2) byte errors. ECC is most useful on methods where errors are
localized (block, rejection) and less effective on arithmetic, where a single
bin flip cascades into an unbounded burst.
Temperature controls how many bits each token can carry. Higher temperature
spreads the probability mass, giving more room for message bits.
For the arithmetic method, use 2.0 or higher. The block method is less
sensitive to temperature.
echo " Hello " | steganeur encode --context " Some context " --dummy
The dummy LM has no token strings, so text-based decode will not

[truncated]
