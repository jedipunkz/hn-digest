---
source: "https://github.com/livekit/eot-bench"
hn_url: "https://news.ycombinator.com/item?id=48588475"
title: "Eot-bench: Open benchmark suite for end-of-turn detection in voice AI"
article_title: "GitHub - livekit/eot-bench: LiveKit End-of-Turn Benchmark · GitHub"
author: "codestackr"
captured_at: "2026-06-18T18:55:50Z"
capture_tool: "hn-digest"
hn_id: 48588475
score: 1
comments: 0
posted_at: "2026-06-18T17:14:14Z"
tags:
  - hacker-news
  - translated
---

# Eot-bench: Open benchmark suite for end-of-turn detection in voice AI

- HN: [48588475](https://news.ycombinator.com/item?id=48588475)
- Source: [github.com](https://github.com/livekit/eot-bench)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T17:14:14Z

## Translation

タイトル: Eot-bench: 音声 AI におけるターン終了検出のためのオープン ベンチマーク スイート
記事のタイトル: GitHub - livekit/eot-bench: LiveKit ターン終了ベンチマーク · GitHub
説明: LiveKit のターン終了ベンチマーク。 GitHub でアカウントを作成して、livekit/eot-bench の開発に貢献してください。

記事本文:
GitHub - livekit/eot-bench: LiveKit ターン終了ベンチマーク · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ライブキット
/
eotベンチ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 開く

アクションメニュー フォルダーとファイル
5 コミット 5 コミット .github/ workflows .github/ workflows 資産 アセット eot_harness eot_harness 出力/ livekit__eot-bench-data__validation__min_silence_100ms 出力/ livekit__eot-bench-data__validation__min_silence_100ms スクリプト スクリプト サイト サイトテスト テスト .gitignore .gitignore ライセンスライセンス README.md README.md pyproject.toml pyproject.toml要件.txt要件.txt uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ターン終了検出のオープンベンチマーク
すべての音声エージェントは、毎回同じ質問に何度も答えなければなりません。
一時停止: ユーザーは話し終えましたか?応答が早すぎるとエージェントが話を切り上げます
人々。答えるのが遅すぎると、会話が無風になってしまいます。ターン終了時
(EoT) 検出は、エージェントのように感じられるエージェントとの違いです。
会話とトランシーバーのような感覚があり、
最初のエージェントが出荷されて以来、音声 AI における最も困難な未解決の問題。
フィールドとして測定することも困難でした。力強い取り組みがたくさんあります
ターン終了の検出は可能ですが、それを比較するための共有された公開方法はありません。結果は得られます。
さまざまなプライベート データセットやさまざまな方法論から取得したものであるため、
再現したり並べたりするのは困難です。欠けているものは共通です
地面。
eot-bench がその共通点です。これはオープンで再現可能なベンチマークであり、ペアになっています
最初の開いたデータセットを使用して
人間とエージェントの実際の会話を 14 か国語で記録。モデルをスコアリングする代わりに
孤立したクリップでは、ライブ音声エージェントが行うのと同じ方法でそれらを評価します。
実際のレイテンシと中断の予算に基づいて一時停止します。評価するために構築しました
LiveKit ターンディテクター v1 、
EoT モデルを構築する誰もが同じモデルで測定できるように、それをリリースします。
足場。
データセット: 実際の会話、14 言語
ライブキット/eot-ベンチ

-データは
ターン終了検出用のこの種の最初のオープン データセット: real
人間からエージェントへのユーザーのやり取りは、音声とテキストのコンテキストを調整して、全体にわたって行われます。
14言語：アラビア語、中国語、オランダ語、英語、フランス語、ドイツ語、ヒンディー語、
インドネシア語、イタリア語、日本語、韓国語、ポルトガル語、スペイン語、トルコ語。
各行はタスク指向の会話からのユーザーの完全なターンであり、注釈が付けられています
少なくとも 100 ミリ秒の沈黙の休止ごとに。最後の休止が本当の終わりだ
順番。それ以前の一時停止はすべて、エージェントが耳を傾けるべきターン途中の躊躇です
を通して。この構造により、ベンチマークは実際のモデルに基づいてスコアを付けることができます。
個別のクリップではなく、音声エージェントが直面する決定を決定します。それは自由です
このリポジトリと一緒に Apache-2.0 をダウンロードして評価できます。
LiveKit Turn Detector v1 は、当社が提供するすべてのモデルの中で最も強力な全体的な結果を示します。
英語および 14 言語すべてで評価されます。完全に探索してください
インタラクティブなリーダーボード » 。を設定します
レイテンシや誤ったカットオフ バジェットを確認し、パレート上ですべてのモデルが再ランク付けされるのを観察します。
フロンティアと言語ごとのヒートマップ。
最も明確な単一のビューは、各モデルが固定時にどれだけのデッドエアを残すかです。
中断の予算。ユーザーの中断が 5% 以下になるように調整されています。
ユーザーが実際に終了してからエージェントはどれくらいの時間待機しますか
応答しますか？低いほど、会話がより活発になります。 (これはエンドポイント遅延ではなく、
モデル推論時間)
4 つの動作点での英国モデル、300 ミリ秒の誤ったカットオフで順序付け
レイテンシーバジェット (最良のものを優先)。どの指標でも低いほど優れています。
A – ポリシー設定がそのレイテンシ バジェットに達していないことを意味します。沈黙のみ
VAD ベースラインは同一の評価を通じて実行されるため、学習されたすべての評価と
市販の検出器は常にタイミングのみを基準にして測定されます。すべての数字は
コミットされた再現可能なアーティファクトから生成される

出力/ の下にあります。
インタラクティブなリーダーボードには、
パレートフロンティア全体と 14 言語すべての内訳。
優れた方向転換検出器は、互いに対立する 2 つの目標を満たさなければなりません。
1 つ目は、ユーザーを決して切断しないことです。誰かが話す前に中断する
考えを終えたということは、音声エージェントが犯す最も不快な失敗であるため、
主な目的は、誤カットオフ率を最小限に抑えることです: ミッドターンでの発砲
実際にはターンの終わりではなかった一時停止。 2つ目は対応することです
ユーザーの話が終わったらすぐに会話を続けることができるため、
死んだ空気で満たされています。それはレイテンシーです: true の後にエージェントが待機する時間です。
フロアに立つ前にターン終了します。
ここでのレイテンシは、コンピューティングではなく、会話のデッドエアです。政策はいつまでだ
モデルの実行速度ではなく、ターンが終了したと確信する前に保持されます。
推論。 600 ミリ秒待機して確実に 600 ミリ秒の結果が表示されるインスタント モデル
待ち時間。これを最小限に抑えるということは、計算を速くすることではなく、正しく判断することを早くすることを意味します。
どちらかの目標を単独で獲得することは簡単にできます。永遠に待っていれば、決して中断することはありません。
すぐに発射できるので、遅れることはありません。重要なのは、次の間のトレードオフです。
彼ら。 eot-bench は、トレードオフを直接測定します。すべてのモデルについて、
エンドポイント ポリシーを使用して、固定の設定で達成可能な最良のレイテンシーを報告します。
誤ったカットオフ予算（およびその逆）に加えて、完全なパレートフロンティア。左下
より良い: 中断が減り、応答が速くなります。単一の精度スコアでは不可能です
これを捉えるため、代わりにトレードオフを中心にベンチマークが構築されています。
完全な方法論については、「評価モデル」を参照してください。
公開ターンレベル データセット livekit/eot-bench-data :
14 か国語の音声とテキストのコンテキストを使用して、人間からエージェントへの実際のターンを行います。
ローカル モデルおよびプロバイダー API 用のバッチおよびストリーミング アダプター インターフェイス、
参照あり

LiveKit Turn Detector v1 / v1-mini、Deepgram Flux 用のアダプター
AssemblyAI、Soniox、OpenAI GPT Realtime、SmartTurn、および UltraVAD。
再現可能な予測アーティファクト、ポリシースイープメトリクス、パレートフロンティア、
操作点テーブルと、output/ の下にコミットされた多言語ヒートマップ。
1 つの言語またはサポートされているすべての言語に対して新しいアダプターを実行するための CLI コマンド
データセット言語に加えて、スケーリングされたバッチ ジョブ用のモーダル ランナー。
python -m pip install -e " .[dev] "
コミットされた英語の比較アーティファクトを再生成します。
eot-harness 比較モデル \
出力/livekit__eot-bench-data__validation__min_silence_100ms/en
多言語操作点アーティファクトを再生成します。
eot-harness 比較言語 \
Output/livekit__eot-bench-data__validation__min_silence_100ms
評価モデル
評価は、EoT モデルがサポートする必要がある決定に基づいて構築されます。
プロダクション音声エージェント: 沈黙するたびに、アシスタントはすぐに応答するべきか、それとも
聞き続けますか？ EoT 検出をオフライン分類として扱う代わりに、
クリップが分離されている場合、ハーネスは完全なターンを因果的沈黙として評価します。
決定を行ってから、そのスコアがサポートできるポリシーを比較します。
データセットの各行は、人間のユーザーがタスク指向から完全に移行したものです。
会話。行には、少なくとも 100 ミリ秒のすべての無音期間が含まれます。決勝戦
無音の期間はユーザーのターンの真の終了であり、 eot というラベルが付けられます。ごとに
前の無音スパンはターン途中の一時停止であり、hold というラベルが付けられます。
ハーネスは、各モデルにこれらのスパンを因果的にスコア付けするように要求します。での予測については、
時刻 t 、アダプターは音声、トランスクリプト コンテキスト、およびメッセージのみを受信します。
それは t までに利用可能でした。 EoT エラーがあるため、これは重要です
通常、孤立したクリップではなく、実際のターン内のあいまいな一時停止で発生します。
モデルは将来のコンテキストまたはオフラインに暗黙的に依存できる場合

セグメンテーションです。
スパンレベルの評価により、各ユーザーのターンが実際の意思決定ポイントに変わります。
音声システムが見ます。優れたモデルは、高い EoT 確率を最終的なモデルに割り当てます。
通常の躊躇やターン途中で確率を低く保ちながら沈黙する
一時停止します。
ポリシースイープと操作点
生のモデル スコアだけではモデルを比較するのに十分ではありません。実稼働システムでは、
スコアに加えてポリシーが追加され、そのポリシーによってユーザーに表示される
迅速な対応と誤った切断の回避との間のトレードオフ。ハーネス
厳選した 1 つのしきい値やタイムアウトを判断するのではなく、ポリシー スペースをスイープします。
スイープ ポリシーには 3 つのノブがあります。
しきい値 : ユーザーのターンを終了するために必要な EoT 信頼度。
action_delay : システムが許可されるまでの最小沈黙時間。
モデルのスコアに基づいて動作します。
timeout : システムが終了するまでに保持する最大無音期間
モデルが発射されていない場合でもターン。
ハーネスはこれらのノブをまとめて評価します。 action_lay を上げると短縮できる
ターン途中の短い一時停止を無視することで誤って中断しますが、同じことが追加されます。
正しく検出されたすべての真のターン終了までの待ち時間。タイムアウトを増やすと、
システムはターン途中でのより長い一時停止を許容しますが、偽陰性が増加します
本当のターン終了を逃した場合、アシスタントは次の時間まで待機することになるため、費用がかかります。
タイムアウトが切れます。
比較レポートは、明示的な誤カットオフおよび
レイテンシ バジェットに加えて、レイテンシ/カットオフ パレート フロンティア。また、
VAD のみのベースラインは同じポリシー グリッドで評価されるため、学習された EoT モデルは
無音のタイミングのみと比較します。
auc や ap などのスカラー分類メトリックは引き続き利用可能です。
実行ごとの診断は行われますが、モデルをランク付けしたり、比較レポートを作成したりすることはありません。
これらは、展開動作とは異なる質問に答えます。

誤解を招く
サーバー側の内部保留の後にイベントを公開するストリーミング API の場合、または
行動の遅れ。ターン途中の短い一時停止は、挿入せずに正しく拒否されたように見える場合があります
明示的なポリシーノブの対応するレイテンシーコスト。政策全般
そのコストをレイテンシ/カットオフ フロンティアと名前付き運用で可視化します。
ポイント。
ハーネスは、パブリック EoT ベンチマーク スキーマ内の既存のデータセットを想定しています。
データセットの構築、アノテーション、VAD 抽出、およびソース固有の取り込み
このパッケージの外に住んでいます。
オーディオはハグフェイスオーディオです
特徴。ハーネスは、Audio(decode=False) にキャストされた列を使用してそれをロードします。
torchcodec ランタイム デコーダは必要ありません。生のバイト自体をデコードします。
サウンドファイル付き。したがって、行は {"bytes": ...} としてエンコードされて到着する可能性があります。
または、すでに {"array": ..., "sampling_rate": ...} としてデコードされています。ハーネスハンドル
両方。
Silence_spans は、開始秒と終了秒を含むスパンのリストです。予測
世代スキップのスパンは 0.1 秒未満です。生成されたすべてのスパンにはラベルが付けられます
それが、eot というラベルが付いた、 Silence_spans の最後のスパンでない限り、保持されます。
メッセージが存在する場合、それは各因果アダプター入力にコピーされます。もし
単語が存在する場合、ハーネスは現在のターンのユーザー メッセージを追加します

[切り捨てられた]

## Original Extract

LiveKit End-of-Turn Benchmark. Contribute to livekit/eot-bench development by creating an account on GitHub.

GitHub - livekit/eot-bench: LiveKit End-of-Turn Benchmark · GitHub
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
livekit
/
eot-bench
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .github/ workflows .github/ workflows assets assets eot_harness eot_harness output/ livekit__eot-bench-data__validation__min_silence_100ms output/ livekit__eot-bench-data__validation__min_silence_100ms scripts scripts site site tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml requirements.txt requirements.txt uv.lock uv.lock View all files Repository files navigation
The open benchmark for end-of-turn detection
Every voice agent has to answer the same question, over and over, on every
pause: is the user done talking? Answer too early and the agent talks over
people; answer too late and the conversation fills with dead air. End-of-turn
(EoT) detection is the difference between an agent that feels like a
conversation and one that feels like a walkie-talkie, and it has been one of the
hardest open problems in voice AI since the first agents shipped.
It has also been hard to measure as a field. There's a lot of strong work on
end-of-turn detection, but no shared, public way to compare it: results come
from different private datasets and different methodologies, which makes them
difficult to reproduce or line up side by side. What's been missing is common
ground.
eot-bench is that common ground. It's an open, reproducible benchmark, paired
with the first open dataset
of real human-to-agent conversations in 14 languages. Instead of scoring models
on isolated clips, it evaluates them the way a live voice agent does: at real
pauses, under real latency and interruption budgets. We built it to evaluate
LiveKit Turn Detector v1 ,
and we're releasing it so anyone building an EoT model can measure on the same
footing.
The dataset: real conversations, 14 languages
livekit/eot-bench-data is
the first open dataset of its kind for end-of-turn detection: real
human-to-agent user turns , with aligned audio and textual context , across
14 languages : Arabic, Chinese, Dutch, English, French, German, Hindi,
Indonesian, Italian, Japanese, Korean, Portuguese, Spanish, and Turkish.
Each row is a complete user turn from a task-oriented conversation, annotated
with every silence pause of at least 100 ms. The final pause is the true end of
the turn; every earlier pause is a mid-turn hesitation the agent should listen
through. That structure is what lets the benchmark score a model on the actual
decisions a voice agent faces, rather than on isolated clips. It's freely
available for download and evaluation, Apache-2.0 alongside this repo.
LiveKit Turn Detector v1 posts the strongest overall results of any model we
evaluated, in English and across all 14 languages. Explore the full
interactive leaderboard » . Set a
latency or false-cutoff budget and watch every model re-rank on the Pareto
frontier and the per-language heatmap.
The clearest single view is how much dead air each model leaves at a fixed
interruption budget. Tuned to interrupt the user no more than 5% of the time,
how long after the user has actually finished does the agent wait before
responding? Lower means a snappier conversation. (This is endpointing delay, not
model inference time.)
English models at four operating points, ordered by false cutoffs at a 300 ms
latency budget (best first). Lower is better on every metric:
A – means no policy setting reached that latency budget. The silence-only
VAD baseline runs through the identical evaluation, so every learned and
commercial detector is always measured against timing alone. All numbers are
generated from the reproducible artifacts committed under output/ .
The interactive leaderboard adds the
full Pareto frontier and the breakdown across all 14 languages.
A good turn detector has to satisfy two goals that pull against each other.
The first is to never cut the user off . Interrupting before someone has
finished a thought is the most jarring failure a voice agent can make, so the
primary objective is to minimize the false-cutoff rate : firing on a mid-turn
pause that wasn't actually the end of the turn. The second is to respond
quickly once the user is done, so the conversation keeps flowing instead of
filling with dead air. That's latency : the time the agent waits after a true
turn ending before taking the floor.
Latency here is conversational dead air, not compute. It's how long the policy
holds before it's confident the turn is over, not how fast the model runs
inference. An instant model that waits 600 ms to be sure still shows 600 ms of
latency. Minimizing it means deciding correctly sooner , not computing faster.
You can trivially win either goal alone: wait forever and you'll never interrupt;
fire instantly and you'll never lag. What matters is the tradeoff between
them. eot-bench measures that tradeoff directly. For every model it sweeps the
endpointing policy, then reports the best latency achievable at a fixed
false-cutoff budget (and vice versa), plus the full Pareto frontier. Lower-left
is better: fewer interruptions, faster responses. A single accuracy score can't
capture this, which is why the benchmark is built around the tradeoff instead.
See Evaluation Model for the full methodology.
The public turn-level dataset livekit/eot-bench-data :
real human-to-agent turns with audio and text context in 14 languages.
Batch and streaming adapter interfaces for local models and provider APIs,
with reference adapters for LiveKit Turn Detector v1 / v1-mini, Deepgram Flux,
AssemblyAI, Soniox, OpenAI GPT Realtime, SmartTurn, and ultraVAD.
Reproducible prediction artifacts, policy-sweep metrics, Pareto frontiers,
operating-point tables, and multilingual heatmaps committed under output/ .
CLI commands for running a new adapter against one language or every supported
dataset language, plus a Modal runner for scaled batch jobs.
python -m pip install -e " .[dev] "
Regenerate the committed English comparison artifacts:
eot-harness compare-models \
output/livekit__eot-bench-data__validation__min_silence_100ms/en
Regenerate the multilingual operating-point artifacts:
eot-harness compare-languages \
output/livekit__eot-bench-data__validation__min_silence_100ms
Evaluation Model
The evaluation is built around the decision an EoT model has to support in a
production voice agent: at each silence, should the assistant respond now or
keep listening? Instead of treating EoT detection as offline classification over
isolated clips, the harness evaluates complete turns as causal silence
decisions, then compares the policies those scores can support.
Each dataset row is a complete human user turn from a task-oriented
conversation. The row includes every silence span of at least 100 ms. The final
silence span is the true end of the user's turn and is labeled eot ; every
earlier silence span is a mid-turn pause and is labeled hold .
The harness asks each model to score those spans causally. For a prediction at
time t , the adapter receives only the audio, transcript context, and messages
that would have been available by t . This matters because EoT errors
usually happen at ambiguous pauses inside a real turn, not at isolated clips
where the model can implicitly rely on future context or offline segmentation.
Span-level evaluation turns each user turn into the actual decision points a
voice system sees. A good model assigns high EoT probability to the final
silence while keeping probability low through ordinary hesitations and mid-turn
pauses.
Policy Sweep and Operating Points
Raw model scores are not enough to compare models. Production systems apply a
policy on top of the score, and that policy determines the user-visible
tradeoff between responding quickly and avoiding false cutoffs. The harness
sweeps the policy space instead of judging one hand-picked threshold or timeout.
The swept policy has three knobs:
threshold : the EoT confidence needed to end the user turn.
action_delay : the minimum silence duration before the system is allowed to
act on the model score.
timeout : the maximum silence duration the system will hold before ending
the turn even if the model has not fired.
The harness evaluates these knobs together. Raising action_delay can reduce
false interruptions by ignoring short mid-turn pauses, but it adds the same
latency to every correctly detected true turn ending. Raising timeout lets the
system tolerate longer mid-turn pauses, but it makes false negatives more
expensive because a missed true end-of-turn leaves the assistant waiting until
the timeout expires.
Comparison reports focus on operating points under explicit false-cutoff and
latency budgets, plus the latency/cutoff Pareto frontier. They also include a
VAD-only baseline evaluated on the same policy grid, so learned EoT models are
compared against silence timing alone.
Scalar classification metrics such as auc and ap remain available in
per-run diagnostics, but they do not rank models or drive comparison reports.
They answer a different question from deployment behavior and can be misleading
for streaming APIs that expose events after an internal server-side hold or
action delay. Short mid-turn pauses may look correctly rejected without putting
the corresponding latency cost on an explicit policy knob. The policy sweep
keeps that cost visible in the latency/cutoff frontier and named operating
points.
The harness expects an existing dataset in the public EoT benchmark schema.
Dataset construction, annotation, VAD extraction, and source-specific ingestion
live outside this package.
audio is a Hugging Face Audio
feature. The harness loads it with the column cast to Audio(decode=False) — so
no torchcodec runtime decoder is required — and decodes the raw bytes itself
with soundfile . Rows may therefore arrive either encoded as {"bytes": ...}
or already decoded as {"array": ..., "sampling_rate": ...} ; the harness handles
both.
silence_spans is a list of spans with start and end seconds. Prediction
generation skips spans shorter than 0.1s . Every generated span is labeled
hold unless it is the final span in silence_spans , which is labeled eot .
If messages is present, it is copied into each causal adapter input. If
words is present, the harness appends a current-turn user message

[truncated]
