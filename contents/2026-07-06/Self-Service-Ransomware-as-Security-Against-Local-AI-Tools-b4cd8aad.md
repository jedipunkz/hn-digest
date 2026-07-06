---
source: "https://blog.brendankeaton.com/self-service-ransomware-as-security-against-local-ai-tools"
hn_url: "https://news.ycombinator.com/item?id=48805363"
title: "Self-Service Ransomware as Security Against Local AI Tools"
article_title: "Brendan Keaton's Blog"
author: "BrKeaton"
captured_at: "2026-07-06T15:01:11Z"
capture_tool: "hn-digest"
hn_id: 48805363
score: 1
comments: 0
posted_at: "2026-07-06T14:42:59Z"
tags:
  - hacker-news
  - translated
---

# Self-Service Ransomware as Security Against Local AI Tools

- HN: [48805363](https://news.ycombinator.com/item?id=48805363)
- Source: [blog.brendankeaton.com](https://blog.brendankeaton.com/self-service-ransomware-as-security-against-local-ai-tools)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T14:42:59Z

## Translation

タイトル: ローカル AI ツールに対するセキュリティとしてのセルフサービス ランサムウェア
記事タイトル: ブレンダン・キートンのブログ
説明: Brendan Keaton によって書かれた投稿のリスト。

記事本文:
ブレンダン キートンのブログ ブレンダン キートン ポートフォリオ ローカル AI ツールに対するセキュリティとしてのセルフサービス ランサムウェア
読了時間: 7 分 投稿日: 2 週間前 「私たちの世界にあるこれらのものたちが大切にされる唯一の方法は、私たちが住んでいる世界を維持するという責任感を持ち、それについて自ら行動を起こす人がいるかどうかです。」 - これを哲学する [168]、ケアの倫理
このアプリケーションは、ローカル AI エージェント脅威モデルのファイル暗号化実験です。完全に動作するコードは、以下の GitHub リンクから入手できます。これは、自己ホスト型の Supabase (Postgresql、Kong など)、Nextjs、Tauri (Rust)、および FastAPI (Python) を使用して構築されています。コードの一部のセクション (特に、Web アプリのフロントエンドの大部分と fastAPI エンドポイントの一部) に生成 AI が使用されました。このプロジェクトは現在監査されておらず、現状のまま提供されています (MIT ライセンス)。免責事項の全文についてはリポジトリを参照してください
ローカル AI ソフトウェアは、デバイス上のファイル、API キー、画像、その他の機密情報を読み取ることができます。
この影響を最小限に抑えるためにいくつかの試みが行われています ( .llmignore を参照)。ただし、これは最終的にはツールとユーザーの間のハンドシェイク合意です。悪意のあるエージェント、プロンプト インジェクション攻撃、さらには善良なエージェントによる過剰なインデックス作成による秘密や機密データの漏洩を根本的に阻止するものは何もありません。 VM での開発や異なるユーザー スペースでの作業など、より安全なソリューションが提供されています。しかし、Mythos は、Anthropic のレッドチームによると、十分に強力なモデルによってこれを打破できることを示しました。
過去数か月間、私は考えられる解決策に取り組んできました。つまり、このツールは、エージェントの作業中にファイルを暗号化するセルフサービス ランサムウェアです。基本的な流れは以下の通りです。
ユーザーがファイルを選択すると、

.env、テスト ファイル、PII を含むデータ ファイルなどにアクセスしたり編集したりすることはできません。
エージェントを開始する前に、付録の「セッション」を開始します。これにより、ファイル名を含む保存中のすべてのファイルが AES-256-GCM で暗号化されます。ファイルのサイズと数に応じて、通常は数秒から 1 分かかります。
AI ツールを通常どおり使用してください。すべての機密ファイルは完全に暗号化されており、AI ツールで読み取ることはできません。
エージェントの実行が終了したらセッションを終了すると、すべてのファイルが復号化されます。
完全なフローとコードには、プロンプト インジェクション攻撃や不正な AI ツールに対する追加の保護手段が多数含まれています。
プロセスの最初のステップはサインアップ時です。ユーザーは 6 桁の PIN を作成すると、6 つの「回復シード」が提供されます。これらはアカウント回復キーではありません。これについては以下で詳しく説明します。
ユーザーは、この「セッション」で保護したいファイルをすべて選択した後、ボタンをクリックして開始します。これにより、AES-256-GCM 暗号化で使用されるキーのリクエストがクライアントからバックエンドに送信されます。他の暗号化ソフトウェアとは異なり、このキーはランダムではなく、4 つの入力を使用して決定的に作成されます。
サインアップ時に提供される 6 つのリカバリ コードのうちの 1 つ
一度生成されたシークレットは、FastAPI サーバー上にのみ保存され、そこから離れることはありません。
次に、これら 4 つが 1 つの文字列に連結されて、SHA256 のシードになります。
以下の図は、セッション キーの作成の流れを説明しています。
キー生成フロー Tauri Client [Rust] FastAPI サーバー [Python] データベース [Postgresql] [1] ユーザーがセッションを開始し、バックエンドに送信されるキーのリクエスト [2] AES-256-GCM 暗号化用の SHA256 キーを決定的に作成 (上記で説明) [3] セッション キーを Postgresql に保存 [4] セッション キーの保存を確認 [5] セッション キーを Tauri Client に返す

t Postgresql鍵生成フローで鍵の保管が確認されるまで、ローカル暗号化を開始できません
[ 1 ] ユーザーがセッションを開始し、キーのリクエストがバックエンドに送信されます
[ 2 ] AES-256-GCM 暗号化用の SHA256 キーを決定的に作成します (上記で説明)
[ 3 ] Postgresql にセッション キーを保存する
[ 4 ] セッションキーの保存確認
[ 5 ] セッションキーを Tauri クライアントに返す
Postgresql からキーの保存が確認されるまで、ローカル暗号化を開始できません
決定論的に作成された導出可能なキーは、暗号化ツールのアンチパターンのように感じるかもしれません。ただし、この脅威モデルはローカル AI エージェント専用であるため、データベースまたはロジック障害が発生した場合にファイルを回復する機能との安全なトレードオフであると私は考えています。その場合でも、標的を絞った専門的な攻撃の場合、攻撃者は fastAPI サーバーの秘密、ユーザーのリカバリ コード、ユーザーの PIN、およびデバイス自体にアクセスする必要があります。
セッション キーを受信すると、デスクトップ アプリケーションはファイルの暗号化を開始できます。各ファイルは以下のパイプラインを通過します。
ファイル暗号化パイプライン [1] すべてのファイル統計を収集する 暗号化が開始される前に、すべての保護されたファイル パスが収集され、サイズが合計され、推定長が作成されます。これは、「実際の」情報とともに、分析のために追跡され、より良いユーザー エクスペリエンスを提供します。また、ファイルがすべて利用可能であること、開かれていないこと、削除されていないことを確認するためのチェックも行われます。 [2] チェックサムを作成する既存のファイル 変更されていない既存のファイルから開始します。 SHA256 チェックサムを取得し、結果を Postgresql に送信します。これを使用すると、エンドツーエンドの整合性と障害のないことを保証できます。 [3] AES-256-GCM 暗号化 実際の暗号化アルゴリズムを実行します。 ≤ 4MB: 12 バイトのノンス > 4MB: 4MB のチャンク、7 バイトのノンス [5] 元のファイルの破棄 元のファイルをランダムなバイト、パンチ ブロック、r で上書きします。

andom rename (ステップ 4 とは別)、fsync、そして最後にリンクを解除します。このプロセスは TRIM 対応であり、使用しているオペレーティング システムに応じてコード パスが異なります。このレベルの注意は、メモリを直接読み取ることができる強力なエージェントや悪意のあるツールに対する将来の備えとして行われます。 [4] 暗号文の書き込み ファイル名をランダム化し、.annex 拡張子を追加します。これは機密性の高いファイル名を隠すためです。 IE「CLIENT_TAXES_2026.csv」は「aj18xl.annex」になります。 [6] ファイル完了の確認 ファイルが暗号化され、安全に削除されたことの確認をバックエンドに送信します。これのマニフェストを保持すると、クラッシュから回復できます。ファイル暗号化パイプライン
暗号化が開始される前に、保護されたすべてのファイル パスが収集され、サイズが合計され、推定長が作成されます。これは「実際の」情報とともに追跡され、より良いユーザー エクスペリエンスを提供するために分析されます。ファイルがすべて利用可能であること、開かれていないこと、削除されていないことを確認するためのチェックも行われます。
既存の変更されていないファイルから始めます。 SHA256 チェックサムを取得し、結果を Postgresql に送信します。これを使用すると、エンドツーエンドの整合性と障害のないことを保証できます。
実際の暗号化アルゴリズムを実行します。
≤ 4MB: 12 バイトのノンス
> 4MB: 4MB チャンク、7 バイトのノンス
ファイル名をランダム化し、.annex 拡張子を追加します。これは機密性の高いファイル名を隠すためです。 IE「CLIENT_TAXES_2026.csv」は「aj18xl.annex」になります
元のファイルをランダムなバイトで上書きし、ブロックをパンチし、ランダムな名前変更 (手順 4 とは別)、fsync、最後にリンクを解除します。このプロセスは TRIM 対応であり、使用しているオペレーティング システムに応じてコード パスが異なります。
このレベルの注意は、メモリを直接読み取ることができる強力なエージェントや悪意のあるツールに対する将来の備えとして行われます。
ファイルが暗号化され、安全に削除されたことを示す確認をバックエンドに送信します。キーピン

これのマニフェストにより、クラッシュからの回復が可能になります。
すべてのファイルが暗号化され、オリジナルが破棄された後、Rust のゼロ化クレートを使用してセッション キーがメモリから削除されます。ローカル マシン上のどこにも存在しないように最善の努力が払われます。
関連ファイルを暗号化するために、いくつかの追加措置が講じられます。 IE は一時ファイル ( .bak 、 .swp など)、VSCode バックアップ、シェル履歴などをチェックします。これもすべてベストエフォートです。すべてのエディターとサードパーティ ソフトウェアが明示的に処理されるわけではありません。
セッションの終了の大部分は、前のセクションを元に戻すだけです。 1 つ注意したいのは、復号化にはサインアップ時に PIN を入力する必要があるということです。これは、悪意のある AI ツールが、復号化を試みるためのセッション キーを要求する API リクエストをバックエンドに送信しようとすることができないことを意味します。また、ブルート フォース攻撃を避けるために、レート制限も意図的に低く設定されています。
キーを受信すると、ファイル名とコンテンツが復元され、チェックサムが暗号化前に存在していたものと比較され、プロセス中に何も失われていないことが確認されます。
このプロジェクトはさまざまな用途に活用できると思います。抽象的な使用例の 1 つは、AI エージェントからテスト ケースを非表示にして、実際のコードを修正するのではなく、動作するようにテストを編集しないようにすることです。
ローカル AI の安全性のために実装できる追加の安全対策は多数あります。たとえば、機密データを含む API 呼び出しをブロックする出口ファイアウォールや、機密データのみをブロックする部分的なファイル暗号化などです。ただし、このプロジェクトは強力な出発点です。
私が見落としている可能性のある問題、改善点、潜在的な悪用について、お気軽にご連絡ください。
最後の注意: これを「ランサムウェア」と呼ぶのは、その目的ではなく、その仕組みについての冗談です。自分のマシン上の自分のファイルのみを暗号化します。

o 地元のエージェントの手の届かない場所に保管してください。ここには何も攻撃的なものはありません。

## Original Extract

A list of posts written by Brendan Keaton.

Brendan Keaton's Blog Brendan Keaton Portfolio Self-Service Ransomware as Security Against Local AI Tools
Read Time: 7 mins Posted: 2 weeks ago “The only way these things in our world ever get cared for is if there is someone out there that has a sense of responsibility for maintaining the world we live in… and takes it upon themselves to do something about it.” - Philosophize This [168], Ethics of Care
This application is a file-encryption experiment for the local AI agent threat model. The full working code is available at the GitHub link below. It is built using self hosted Supabase (Postgresql, Kong, etc), Nextjs, Tauri (Rust), and FastAPI (Python). Generative AI was used for some sections of code (particularly most of the frontend for the webapp and some of the fastAPI endpoints). This project is currently unaudited and provided as-is (MIT License). See the repo for the full disclaimer
Local AI software can read your files, API keys, images, and any other sensitive information on your device.
Some attempts have been made at minimizing the effects of this (see .llmignore ); however, this is ultimately a handshake agreement between the tools and you. There is nothing fundamentally stopping a bad actor agent, prompt injection attacks, or even a good agent over-indexing from leaking secrets or sensitive data . More secure solutions have been offered, such as developing in VMs or working under different user spaces; however, Mythos has shown this can be broken out of by a powerful enough model according to Anthropic's red team .
Over the course of the last few months, I have worked on a possible solution. In short, the tool is self-service ransomware that encrypts your files while agents work. The basic flow is as follows:
A user selects files that AI should not have access to or be able to edit, such as .env s, testing files, data files with PII, etc.
Before starting an agent, start an Annex “session”. This encrypts all of the files at rest with AES-256-GCM, including the file names. It usually takes a few seconds to a minute depending on the size and count of the files.
Use your AI tools normally. All of your sensitive files are encrypted entirely and cannot be read by AI tools.
At the conclusion of the agent running, end the session, and all of your files are decrypted.
The full flow and code includes a number of additional safeguards against prompt injection attacks and bad actor AI tools:
The first step in the process is during sign up. Users create a 6 digit PIN and are then provided 6 “recovery seeds”. These are not account recovery keys; more on this below.
After a user selects all of the files they want to be protected for this “session”, they click a button to initiate. This sends a request from the client to the backend for a key that will be used in the AES-256-GCM encryption. Unlike other encryption softwares, this key is not random and is deterministically created with 4 inputs:
One of the six recovery codes provided at sign up
A once-generated secret , stored only on the FastAPI server that never leaves
These four are then concatenated into a single string to be a seed for SHA256.
Below is a graphic, explaining the flow of the creation of a session key.
Key Generation Flow Tauri Client [Rust] FastAPI Server [Python] Database [Postgresql] [1] User initiates session, request for key sent to backend [2] Deterministically create a SHA256 key for AES-256-GCM encryption (explained above) [3] Store session key in Postgresql [4] Confirm storage of session key [5] Return session key to Tauri Client Local encryption cannot begin until key storage is confirmed from Postgresql Key Generation Flow
[ 1 ] User initiates session, request for key sent to backend
[ 2 ] Deterministically create a SHA256 key for AES-256-GCM encryption (explained above)
[ 3 ] Store session key in Postgresql
[ 4 ] Confirm storage of session key
[ 5 ] Return session key to Tauri Client
Local encryption cannot begin until key storage is confirmed from Postgresql
Having deterministically created, derivable keys may feel a bit like an anti-pattern for an encryption tool; however, because the threat model is exclusively for local AI agents, I believe it is a safe-trade off for the ability to recover files in the event of a database or logic failure. Even then, for a targeted, professional attack, the bad actors would still need to gain access to the fastAPI server secret, a user's recovery codes, a user's pin, and the device itself.
Upon receiving the session key, the desktop application can begin encrypting files. Each file goes through the pipeline below:
File Encryption Pipeline [1] Collect all file stats Before encryption begins, all protected file paths are gathered, sizes summed, and estimated length is created. This, along with the “actual” is tracked for analytics and to provide a better user experience. There is also a check done to ensure the files are all available, not open, and none have been deleted. [2] Existing File to checksum Start with the existing, unaltered file. SHA256 checksum and send result to Postgresql. This can be used to ensure end-to-end integrity and no failure. [3] AES-256-GCM Encryption Run the actual encryption algorithm. ≤ 4MB: 12-byte nonce > 4MB: 4MB chunks, 7 byte nonce [5] Destroy Original File Overwrite original file with random bytes, punch blocks, random rename (separate from step 4), fsync, and finally unlink. This process is TRIM aware, and has different code paths depending on the operating system in use. This level of care is taken as future proofing against stronger agents and bad actor tools that are capable of reading memory, directly. [4] Write Ciphertext Randomize filename, add .annex extension. This is to hide sensitive file names. IE “CLIENT_TAXES_2026.csv” becomes “aj18xl.annex” [6] Confirm File Completion Send confirmation that the file has been encrypted, and securely deleted to the backend. Keeping a manifest of this allows for crash recovery. File Encryption Pipeline
Before encryption begins, all protected file paths are gathered, sizes summed, and estimated length is created. This, along with the “actual”, is tracked for analytics to provide a better user experience. There is also a check done to ensure the files are all available, not open, and have not been deleted.
Start with the existing, unaltered file. SHA256 checksum and send result to Postgresql. This can be used to ensure end-to-end integrity and no failure.
Run the actual encryption algorithm.
≤ 4MB: 12-byte nonce
> 4MB: 4MB chunks, 7 byte nonce
Randomize filename, add .annex extension. This is to hide sensitive file names. IE “CLIENT_TAXES_2026.csv” becomes “aj18xl.annex”
Overwrite original file with random bytes, punch blocks, random rename (separate from step 4), fsync, and finally unlink. This process is TRIM aware and has different code paths depending on the operating system in use.
This level of care is taken as future proofing against stronger agents and bad actor tools that are capable of reading memory directly.
Send confirmation that the file has been encrypted, and securely deleted to the backend. Keeping a manifest of this allows for crash recovery.
After all files have been encrypted and originals destroyed, the session key is removed from memory using the zeroize crate in Rust. Best efforts are taken to ensure it no longer exists anywhere on the local machine.
Some additional measures are taken to encrypt relevant files. IE checking for temp files ( .bak , .swp , etc.), VSCode backups, shell history, etc. This is all, also, best effort. Not all editors and third party softwares are explicitly handled.
The vast majority of ending a session is just… undoing the previous section. The one callout I would like to make is that decrypting requires inputting the PIN from signup. This means a malicious AI tool cannot attempt to make an API request to our backend asking for the session key to attempt decryption. The rate limit is also intentionally low to avoid brute force attacks.
Once the key has been received, filenames and contents are restored, and checksums are compared with what existed pre-encryption to ensure nothing was lost in the process.
I believe this project can be used for a number of purposes. One abstract usecase would be for hiding test cases from AI agents so that they don't edit the tests to work rather than fix the actual code .
There are a number of additional safety measures that could be implemented for local AI safety, such as an egress firewall that blocks API calls that contain sensitive data or even partial file encryption that only blocks sensitive data. However, this project is a strong starting point.
Feel free to call me out on issues, improvements, or potential exploits that I may have overlooked.
Final note: Calling this "ransomware" is a joke about how it works, not what it's for. It only ever encrypts your own files on your own machine to keep them out of reach of local agents. There's nothing offensive here.
