---
source: "https://www.backblaze.com/blog/results-from-the-backblaze-generative-ai-media-hackathon/"
hn_url: "https://news.ycombinator.com/item?id=49381062"
title: "Results from the Backblaze Generative AI Media Hackathon"
article_title: "Results from the Backblaze Generative AI Media Hackathon"
image: "https://www.backblaze.com/blog/wp-content/uploads/2026/08/Hackathon-Winners-blog-header.png"
author: "roan-we"
captured_at: "2026-08-20T23:17:03Z"
capture_tool: "hn-digest"
hn_id: 49381062
score: 1
comments: 0
posted_at: "2026-08-20T22:20:53Z"
tags:
  - hacker-news
  - translated
---

# Results from the Backblaze Generative AI Media Hackathon

- HN: [49381062](https://news.ycombinator.com/item?id=49381062)
- Source: [www.backblaze.com](https://www.backblaze.com/blog/results-from-the-backblaze-generative-ai-media-hackathon/)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T22:20:53Z

## Translation

タイトル: Backblaze Generative AI Media Hackathon の結果
説明: Backblaze Generative AI Media Hackathon で 5 つのジェネレーティブ メディア アプリと、それらを本番環境に対応させるオーケストレーション パターン。

記事本文:
クラウド ストレージ クラウド ストレージに関するヒント、ストーリー、洞察、調査、ハウツー。
バックアップ バックアップ戦略と技術、バックアップに関する技術的な詳細と洞察など、データのバックアップに関連するトピック。
ビジネス ラボのブートストラッピング、スタートアップの物語と課題、マーケティングの混乱など。
Backblaze Bits Backblaze 製品とリリースに関連するトピックのほか、Backblaze の新入社員、求人情報、オフィス周りのこと、不遜なユーモアなど。
Tech Lab の記事は、技術的な内容をより深く探求しており、多くの場合、開発者を対象としています。
パートナー ニュース 最新のパートナーシップに関する発表、パートナー コンテンツなどを読むためのスペースです。
Backblaze Generative AI メディア ハッカソンの結果
2026 年 8 月 14 日、ジェロニモ・デ・レオン著 // コメントはありません
オブジェクト ストレージ上でのジェネレーティブ メディア アプリケーションの構築: 最も強力なプロジェクトの共通点
2026 年 6 月 22 日から 8 月 3 日までの間に、1,314 人が数百のプロジェクトを 1 つの概要に入力しました。それは、オープンソースの Genblaze SDK を通じて調整された Backblaze B2 オブジェクト ストレージ上にジェネレーティブ メディア アプリを構築し、賞金総額 10,000 ドルの分け前を獲得するというものでした。
ほぼすべてのエントリで画像、ビデオ、またはサウンドトラックを生成でき、それらを適切に生成できます。最も強力なものを分けたのは、その下にある層、つまり、何が書き留められるか、何が後で証明できるか、そしてファイルをチェック、修正、または削除に対してロックする必要がある場合に何が起こるかでした。
ここでは、最も成果が上がった 5 つの部品と、再利用する価値のある部品を詳しく見ていきます。すべてのエントリはハッカソン プロジェクト ギャラリーで閲覧できます。
概要と意図的に幅を狭くした理由
概要は、ほとんどのハッカソンよりも内容が狭かったです。提出物は、単なる動作デモやリポジトリではありません。 Devpost チームは書面で次のことを説明する必要がありました。

彼らのアプリが実際に Backblaze B2 と Genblaze の両方をどのように使用したか、そしてその背後にあるすべてのプロバイダーとモデルをリストします。審査は、現実世界の実用性、生産準備状況、B2 の有意義な使用、Genblaze の有意義な使用という 4 つの基準に基づいて行われました。
最後のペアは、機能を実行する制約です。 「意味のある使用」では、チームがアセットを生成し、バケットにバイトをドロップし、そのストレージを呼び出すデフォルトのアーキテクチャは除外されます。それは、ストレージ層が実際に何のためにあるのかについての決定を強制します。以下は、テクニカル ディレクターおよびクリエイティブ ディレクターが、まさにそのルーブリックに基づいて 5 つの完成したパイプラインを実行する方法、つまり各アプリが実際に誰向けであるか、デモを経ても持ちこたえられるかどうか、ストレージとオーケストレーションの選択がどの程度深くまで及ぶかを説明する方法です。
Backblaze B2 は、これらのアプリのすべてが構築する必要がある基盤です。S3 互換のオブジェクト ストレージで、生成されたアセット、サムネイル、メタデータ、メディア パイプラインが放出する出所レコードの膨大な量に合わせてサイズ設定されており、10 GB を含めて無料で開始できます。
Backblaze のオープンソース オーケストレーション SDK である Genblaze がこれにフィードを提供します。OpenAI、Google、Runway、Luma、イレブンラボ、Stability Audio などのプロバイダーにまたがる統合パイプライン API に加え、GMI Cloud や NVIDIA NIM などのプラットフォームを通じて提供されるモデルにより、チームはオーケストレーションを書き換えることなくプロバイダーを交換できます。実行ごとに、メディア ファイル自体 (.mp4、.png、.mp3) に直接埋め込み、B2 または任意の S3 互換ストアに保存できる正規の出自マニフェストが生成されます。
このハッカソンでは GMI Cloud とも提携し、チームが画像、ビデオ、オーディオ、チャット、推論、マルチモーダル作業用のオープンソース生成モデルに簡単にアクセスできるようにしました。そのため、以下で GMI Cloud がプロバイダーとして何度も登場します。
2 つのプリミティブ、6 週間、数百のチーム、および 5 つの非常に異なるもの

厳格さの熱烈な例。
https://devpost.com/software/firstframe
github.com/migarci2/firstframe
全体のレンダリングを待たずに済む AI 動画広告のレビュー ルーム。
firstframe は、AI 生成のビデオ広告を依頼するマーケティング チームとクリエイティブ チームのためにレビュー ルームを構築します。誰かが反応する前に完全なマルチシーンのレンダリングを待つのではなく、準備ができた瞬間に最初に完成したシーンをライブ HLS プレイリストとしてストリーミングし、後のシーンが到着するたびにセグメントを追加します。レビュー担当者は、広告の作成後ではなく、作成中にメモを書き始めます。
生成されたすべてのシーンは、人間が見る前に実際のビジョン モデルによってスコア付けされるため、明らかに壊れた出力は、レビュー担当者の受信箱に送信されるのではなく、自動的に検出されて再試行されます。フェイルオーバー ステップは、通常の応答が遅い場合ではなく、本物のプロバイダー エラーが発生した場合にのみバックアップ モデルを交換します。シーンがクリアされると、そのマスター ファイルとマニフェストは 30 日間削除されないようにロックされます。これは、コードがロックされたファイルを削除しようとして拒否をキャッチすることで証明される保証です。自動化された QA と承認内容の改ざん防止記録を組み合わせることで、生成パイプラインがブランドが実際に承認できるものに変わります。
B2 では、バケットのフォルダー構造がワークフローとしても機能します。シーンは、レビューをクリアするにつれて、受信、実行、来歴、承認、および拒否されたプレフィックスを経て移動します。シーンが承認されると、そのマスター ファイルとマニフェストは、ガバナンス モードで B2 のオブジェクト ロックを使用して実際に 30 日間のライトワンス保持を取得します。コードは、ロックされたオブジェクトをバージョン ID で削除しようとし、B2 がスローバックする拒否をキャッチすることで、それが単なる装飾ではないことを証明します。レビュー担当者のアプリケーション キーのスコープは、承認されたフォルダー al に制限する名前プレフィックスが付いた readFiles 機能に限定されます。

1つ。 4 つの個別のライフサイクル ルールがバケットのさまざまなプレフィックスをカバーします。受信で停止したアップロードは 24 時間後にマルチパート パーツがキャンセルされますが、拒否されたオブジェクト、進行中、承認されたオブジェクトはそれぞれ独自のタイマーで期限切れになります。ビデオ セグメント自体は、ffmpeg が各セグメントを終了するときに B2 に配置され、セグメントごとにプレイリストが書き換えられます。これにより、B2 は、事後に埋まるアーカイブではなく、ライブ ブロードキャストのターゲットになります。読み取りはパス スタイルの署名付き URL として送信され、プライベート B2 バケットで仮想ホスト スタイルの署名が失敗するという既知の問題を回避します。 B2 独自のイベント通知 (5 つの署名済み Webhook ルール) は、レビュー ルームをリアルタイムで同期し、アカウントのイベント通知 API が有効になっていない場合は通常のポーリングにフォールバックします。また、アプリは B2 自身のトランザクション上限を監視し、それを超えて通話が拒否された場合でもクラッシュするのではなく、ローカル ディスクにバックオフします。
Genblaze では、ThresholdEvaluator によってスコア付けされた AgentLoop を介して生成が実行されます。ジャッジは実際のビジョン モデルであり、固定の再試行回数ではなく、実際にレンダリングされたキーフレームをグレーディングする NVIDIA NIM llama-3.2-90b-vision-instruct インスタンスです。オーディオとビデオの 2 つのパイプライン ブランチは、ストレート チェーンとして実行されるのではなく、単一のコンポジター ノードにファン状に配置されます。fallback_models フェイルオーバーが組み込まれており、タイムアウトではなく本物のモデル エラーが発生した場合にのみトリガーされることが確認されています。また、すべての実行には 2 つの系統の層が含まれます。つまり、シーンとループ反復間で共有される実行 ID と、その上に重ねられる 2 番目のカスタム チェーン ID です。チームは、配布された MP4 内に直接マニフェストを埋め込んだため、別の検証コマンドを使用して、ファイルが存在すると主張するすべてのアセットを再ダウンロードして再ハッシュすることもできます。その過程で、彼らは 3 つのプル リクエストと Genblaze 自体に対する問題を提出しました。

firstframe は、レビュー担当者の目の前で B2 の機能セットを動作させます。ガバナンス モードのオブジェクト ロック、スコープ指定されたアプリケーション キー、ライフサイクル ルール、イベント通知、および署名付き URL はすべて、レビュー ワークフローの負荷がかかる部分であり、ワークフロー内を移動するすべてのシーンで目に見える作業を実行します。これは、機能するためだけではなく、信頼されるように構築されたアーキテクチャです。
https://devpost.com/software/beavous
キャンペーン ジェネレーターは、自身のストレージを再チェックするように構築されており、読み戻された瞬間にすべてのアセットを再検証します。
beavous は、単一の製品写真から完全な有料ソーシャル キャンペーンを必要とするマーケティング担当者や小規模コマース チーム向けに構築されています。これには、1 つのヒーロー画像ではなく、4 つのクリエイティブ コンセプト、すべての配置で 16 のトリミングされたアスペクト比、ラベル上の広告コピー、およびポートレート ビデオ リールが含まれており、B2 から直接抽出された検証済み ZIP としてパッケージ化されています。
パブリック API はタスク キューの背後でプライベート ワーカーにハンドオフされ、すべてのキャンペーンは組織に名前空間が設定されるため、テナントは互いに分離された状態になります。世代が拒否された場合、アプリは最初からやり直すことはできません。元の試行に修正を連鎖させます。これは、ワンショットの再試行よりも、実際のクリエイティブなレビューの実際の動作に近いものになります。
B2 では、製品が設計上マルチテナントであるため、キーはコンテンツ ハッシュではなく組織およびキャンペーンごとに階層的に編成されます。すべてのアップロードとダウンロードは署名付き URL を介して行われ、データベースには生のリンクは保存されず、オブジェクト キーとハッシュのみが保存されます。アセットが読み戻されるたびに、beavous はマニフェスト チェックだけを信頼するのではなく、アセットを再ダウンロードし、バイトを個別に再ハッシュします。 B2 は、生成されたすべてのアセットの単一の唯一の記録システムとして扱われます。つまり、他に同期を保つ必要のない、クリーンで信頼できる唯一の情報源の設計です。
Genblaze では、3 つのカスタム プロバイダーが Gemini i を処理します

mage 生成、Gemini ビデオ、および Veo 画像からビデオへの変換があり、それぞれモデルに独自の段階的な価格が登録されています。さらに興味深いのは修正チェーンです。世代が拒否されると、そのマニフェストが次の試行の親となり、ほとんどのワンショット生成パイプラインが完全にスキップする明示的な修正系統です。プロンプトはプライベートとしてマークされるため、テキストはパブリック マニフェストには決して配置されず、ハッシュされた参照のみがパブリック マニフェストに配置されます。マニフェストは修正の親として信頼される前に、アプリ自身の書き込みだからといって正しいとはみなされず、独立して再検証されます。
「自分のストレージを信頼しますか?」に対する beavous の答えは単純です。「いいえ、決して信頼できません」です。毎回、もう一度確認してください。これは、自分の書き込みを信頼するよりも遅い設計であり、より正直な設計です。
https://devpost.com/software/takegraph
B2 に対してライブで独自の再利用、回復、リリースの整合性を証明できる生成メディア用のビルド システム。
takegraph は、単一世代ではなく、プロダクションを実行するチーム向けに構築されています。途中でスクリプトを調整するようなプロジェクトでは、すべてを最初から再レンダリングする必要はありません。全体をソフトウェア ビルドのように扱います。つまり、仕様が変更されると、フィンガープリントを再計算し、実際に無効化されたものだけを再構築し、それ以外はすべて再利用するコンテンツ アドレス指定の依存関係グラフです。
また、チームは、再利用、リカバリ、リリースの背後にある実際のバイトを再ダウンロードして再ハッシュし、B2 から直接ライブで、ログを信頼する代わりに整合性を自分たちでチェックすることもできます。この種の自己監査により、数か月間無人で実行されることを意図したパイプラインと、1 回のデモに耐えるために構築されたパイプラインが分離されます。
B2 では、コンテンツ アドレス指定キーは 2 レベルのハッシュ分割を使用するため、ディレクトリ リストは大規模な場合でも高速に維持され、B2 独自のイベント通知 (HMAC-SHA256 署名付き Webhook) が機能します。

バックグラウンド プロセスを実行し、Webhook が見逃された場合に備えて別のリコンサイラーが定期的にすべてを手動で再チェックします。また、一度に 1 つのリコンサイラーのみが実行されるようにデータベース ロックを使用してワーカー間で調整されます。未検証のアップロードは、自動的に期限切れになる実際のライフサイクル ルールに基づく検疫プレフィックスに到達し、不正なキーは黙って書き換えられるのではなく、完全に拒否されます。 2 つの最小特権のアプリケーション キー (日常業務用とリリース用) は、それぞれ 1 つのバケットにスコープされており、作業バケットの CORS ルールは、特に署名付きのブラウザー アップロードをサポートするために存在します。検証 (再ダウンロード、再ハッシュ、証明) は製品の機能であり、内部ツールではありません。
Genblaze では、takegraph は、専用の実行ビルダー、コンテンツ アドレス指定可能なストレージ シンク、マニフェスト、およびすべてのステップに関連付けられた可観測性イベントなどのアイデアに基づいて実際のパイプラインを構築します。チームはメディア生成側を無駄なく保ち、画像とビデオの GMI クラウド コネクタ (ハッカソンのパートナー プラットフォームの 1 つ) を直接呼び出し、すべてのカスタム エンジニアリングをピッチ全体を機能させるレイヤー、つまりその下にあるストレージと一貫性システムに向けました。
takegraph はストレージの検証を製品自体に変えます: 再利用、リカバリ、統合のリリース

[切り捨てられた]

## Original Extract

Five generative media apps and the orchestration patterns that made them production-ready in the Backblaze Generative AI Media Hackathon.

Cloud Storage Tips, stories, insights, investigations, and how-tos about cloud storage.
Backing Up Topics related to backing up data, including backup strategies and techniques, and technical details and insights about backup.
Business Lab Bootstrapping, start-up tales and challenges, marketing mayhem, and more.
Backblaze Bits Topics related to Backblaze products and releases, as well as Backblaze new hires, job postings, things around the office, irreverent humor, and more.
Tech Lab Articles that explore our technical content more deeply, and are often geared towards our developer audience.
Partner News Your space to read up on our latest partnership announcements, partner content, and more.
Results from the Backblaze Generative AI Media Hackathon
August 14, 2026 by Jeronimo De Leon // No Comments
Building generative media applications on object storage: What the strongest projects have in common
Between June 22 and August 3, 2026, 1,314 people entered hundreds of projects into a single brief : build a generative media app on Backblaze B2 object storage, orchestrated through the open-source Genblaze SDK , for a share of a $10,000 prize pool.
Almost every entry could generate an image, a video, or a soundtrack, and generate it well. What separated the strongest was the layer underneath: what gets written down, what can be proved later, and what happens when a file has to be checked, corrected, or locked against deletion.
Here’s a closer look at the five that went furthest, and the parts worth reusing. Every entry is browsable in the hackathon project gallery .
The brief, and why it was narrow on purpose
The brief was narrower than most hackathons get. Submissions couldn’t just be a working demo and a repo. Devpost required teams to explain, in writing, how their app actually used both Backblaze B2 and Genblaze , and to list every provider and model behind it. Judging ran against four criteria: real-world utility, production readiness, meaningful use of B2, and meaningful use of Genblaze.
That last pair is the constraint that did the work. “Meaningful use” rules out the default architecture, where a team generates an asset, drops the bytes in a bucket, and calls that storage. It forces a decision about what the storage layer is actually for. What follows is written the way a technical and creative director would walk five finished pipelines against exactly that rubric: who each app is actually for, whether it holds up past the demo, and how deep the storage and orchestration choices go.
Backblaze B2 is the ground every one of these apps had to build on: S3-compatible object storage, sized for the sheer volume of generated assets, thumbnails, metadata, and provenance records a media pipeline throws off, free to start with 10GB included.
Genblaze, Backblaze’s open-source orchestration SDK, is what feeds it: a unified Pipeline API spanning providers like OpenAI, Google, Runway, Luma, ElevenLabs, and Stability Audio, plus models served through platforms such as GMI Cloud and NVIDIA NIM, so a team can swap providers without rewriting its orchestration. Every run produces a canonical provenance manifest that can be embedded directly into the media file itself (an .mp4, a .png, an .mp3) and persisted to B2 or any S3-compatible store.
The hackathon also partnered with GMI Cloud, giving teams easy access to open-source generative models for image, video, audio, chat, reasoning, and multimodal work, which is why it turns up as a provider more than once below.
Two primitives, six weeks, hundreds of teams, and five very different examples of rigor.
https://devpost.com/software/firstframe
github.com/migarci2/firstframe
A review room for AI video ads that doesn’t make you wait for the whole render.
firstframe builds a review room for the marketing and creative teams who commission AI-generated video ads: instead of waiting on a full multi-scene render before anyone can react, it streams the first finished scene as a live HLS playlist the moment it’s ready, appending segments as later scenes land. A reviewer starts giving notes while the ad is still being made, not after.
Every generated scene is scored by an actual vision model before a human ever sees it, so obviously broken output gets caught and retried automatically rather than shipped to a reviewer’s inbox. A failover step swaps in a backup model only on a genuine provider error, never on an ordinary slow response, and once a scene clears review its master file and manifest are locked against deletion for thirty days: a guarantee the code proves by trying to delete a locked file and catching the rejection. Automated QA paired with a tamper-evident record of what was approved is what turns a generation pipeline into something a brand could actually sign off on.
On B2, the bucket’s folder structure doubles as a workflow: a scene moves through incoming, running, provenance, approved, and rejected prefixes as it clears review. Once a scene is approved, its master file and manifest get a real thirty-day write-once hold using B2’s Object Lock in Governance mode, and the code proves that isn’t just decorative by trying to delete a locked object by its version ID and catching the rejection B2 throws back. A reviewer’s application key is scoped to the readFiles capability with a name prefix restricting it to the approved folder alone. Four separate lifecycle rules cover the bucket’s different prefixes: stalled uploads in incoming have their multipart parts cancelled after 24 hours, while rejected, in-progress, and approved objects each age out on their own separate timers. The video segments themselves land in B2 as ffmpeg finishes each one, with the playlist rewritten after every segment, which makes B2 a live broadcast target rather than an archive that fills up after the fact. Reads go out as path-style presigned URLs, working around a known issue where virtual-host-style presigning fails on a private B2 bucket. B2’s own Event Notifications, five signed webhook rules, keep the review room in sync in real time, with a fallback to plain polling if an account’s Event Notifications API isn’t enabled. The app also watches B2’s own transaction cap and backs off to local disk instead of crashing when a call gets rejected for exceeding it.
On Genblaze, generation runs through an AgentLoop scored by a ThresholdEvaluator. The judge is a real vision model, an NVIDIA NIM llama-3.2-90b-vision-instruct instance grading the actual rendered keyframes, not a fixed retry count. Two pipeline branches, audio and video, fan into a single compositor node instead of running as a straight chain, a fallback_models failover is wired in and confirmed to trigger only on a genuine model error rather than a timeout, and every run carries two layers of lineage: a shared run id across scenes and loop iterations, plus a second, custom chain id layered on top of that. The team even embedded the manifest directly inside the delivered MP4, so a separate verification command can re-download and re-hash every asset the file claims exists. Along the way they filed three pull requests and an issue against Genblaze itself.
firstframe puts B2’s feature set to work in front of the reviewer. Object Lock in Governance mode, scoped application keys, lifecycle rules, Event Notifications, and presigned URLs are all load-bearing parts of the review workflow, doing visible work on every scene that moves through it. That’s architecture built to be trusted, not just to work.
https://devpost.com/software/beavous
A campaign generator built to double-check its own storage, re-verifying every asset the moment it’s read back.
beavous is built for the marketers and small commerce teams who need a full paid-social campaign out of a single product photo: not one hero image, but four creative concepts, sixteen cropped aspect ratios for every placement, on-label ad copy, and a portrait video reel, packaged as a verified ZIP pulled straight from B2.
A public API hands off to a private worker behind a task queue, and every campaign is namespaced to an organization so tenants stay isolated from each other. When a generation gets rejected, the app doesn’t start over. It chains a correction onto the original attempt, which is closer to how a real creative review actually works than a one-shot retry.
On B2, keys are organized hierarchically by organization and campaign rather than by content hash, because the product is multi-tenant by design. Every upload and download goes through a presigned URL, and the database never stores a raw link, only an object key and a hash. Every time an asset is read back, beavous re-downloads it and re-hashes the bytes independently, rather than trusting a manifest check alone. B2 is treated as the single, sole system of record for every generated asset: a clean, one-source-of-truth design with nothing else to keep in sync.
On Genblaze, three custom providers handle Gemini image generation, Gemini video, and Veo image-to-video, each with its own tiered pricing registered on the model. The more interesting move is a correction chain: when a generation gets rejected, its manifest becomes the parent of the next attempt, an explicit correction lineage most one-shot generation pipelines skip entirely. Prompts are marked private so the text never lands in the public manifest, only a hashed reference to it does, and before any manifest is trusted as a correction parent it gets independently re-verified, not assumed correct just because it was the app’s own write.
beavous’s answer to “do you trust your own storage” is simple: no, never. Check it again, every time. That’s a slower design than trusting your own write, and a more honest one.
https://devpost.com/software/takegraph
A build system for generative media that can prove its own reuse, recovery, and release integrity, live, against B2.
takegraph is built for teams running a production, not a single generation: the kind of project where a script tweak halfway through shouldn’t mean re-rendering everything from scratch. It treats the whole thing like a software build: a content-addressed dependency graph that, when a spec changes, recomputes fingerprints, rebuilds only what’s actually invalidated, and reuses everything else.
A team can also re-download and re-hash the actual bytes behind any reuse, recovery, or release straight from B2, live, checking integrity themselves instead of trusting a log. That kind of self-auditing separates a pipeline meant to run unattended for months from one built to survive a single demo.
On B2, content-addressed keys use a two-level hash split so directory listings stay fast at scale, B2’s own Event Notifications (HMAC-SHA256-signed webhooks) feed a background process, and a separate reconciler periodically re-checks everything by hand in case a webhook is ever missed, coordinated across workers with a database lock so only one reconciler runs at a time. Unvalidated uploads land in a quarantine prefix backed by a real lifecycle rule that expires it automatically, and a bad key gets rejected outright rather than silently rewritten. Two least-privilege application keys, one for day-to-day work and one for releases, are each scoped to a single bucket, and CORS rules on the work bucket exist specifically to support presigned browser uploads. Verification (re-download, re-hash, prove it) is a feature of the product, not an internal tool.
On Genblaze, takegraph builds a real pipeline around the idea: a dedicated run builder, a content-addressable storage sink, manifests, and observability events tied to every step. The team kept its media-generation side lean, calling straight through the GMI Cloud connector for image and video (one of the hackathon’s partner platforms), and pointed all of its custom engineering at the layer that makes the whole pitch work: the storage and consistency system underneath.
takegraph turns storage verification into the product itself: reuse, recovery, and release integ

[truncated]
