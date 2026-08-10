---
source: "https://bobdahacker.com/blog/tldv-hack"
hn_url: "https://news.ycombinator.com/item?id=49242739"
title: "Over 181,000 AI meeting recordings left wide open in note taking app"
article_title: "tl;dv (Too Lazy; Didn't Validate): 181,874 Meetings Left Wide Open | bobdahacker"
author: "colesantiago"
captured_at: "2026-08-10T12:44:31Z"
capture_tool: "hn-digest"
hn_id: 49242739
score: 3
comments: 0
posted_at: "2026-08-10T12:26:05Z"
tags:
  - hacker-news
  - translated
---

# Over 181,000 AI meeting recordings left wide open in note taking app

- HN: [49242739](https://news.ycombinator.com/item?id=49242739)
- Source: [bobdahacker.com](https://bobdahacker.com/blog/tldv-hack)
- Score: 3
- Comments: 0
- Posted: 2026-08-10T12:26:05Z

## Translation

タイトル: 181,000 件を超える AI 会議の記録がメモアプリに広く残されている
記事のタイトル: tl;dv (遅すぎる; 検証されなかった): 181,874 の会議が広く放置されたまま |ボブダハッカー
説明: tl;dv 上の Firestore セキュリティ ルールが欠落していたために、35,003 ドメインにわたる 84,312 人のユーザーからの 181,874 件の会議がどのように公開され、その中には招待されずに参加できるライブ通話も含まれていたことと、6 か月間公開されても領収書を見るだけで何も得られなかったことがわかりました。

記事本文:
tl;dv (遅すぎる; 検証されていない): 181,874 件の会議が未公開のまま
2026 年 8 月 4 日
ボブダハッカー
これを報告したのは 2026 年 1 月 28 日で、半年後の現在は 2026 年 7 月です。 Firestore データベースはまだ広く開かれています。 CTOは応答しませんでした。私のメールが長すぎて見てもらえなかったのだと思います。
tl;dv (Too Long; Didn't View) は、AI 会議記録プラットフォームです。 Google Meet、Zoom、Teams の通話にボットを投入し、すべてを記録して文字に起こし、AI で概要を生成します。ユーザー数は 200 万人を超えています。投資家からの支援を受けています。 LinkedIn のセールス インフルエンサー コミュニティの半数によって支持されています。
営業電話、就職面接、業績評価、社内戦略セッションが保存されます。 「この通話は録音されています」と誰かが言い、全員が緊張しながら笑い、その後 45 分間企業秘密を共有するような内容です。
tl;dv にサインアップすると、プラットフォームは JWT で認証し、 gw.tldv.io/v1/users/firebase/token 経由で Firebase トークンと交換します。このトークンを使用すると、projects/lmi-store/databases/(default) にある Firestore データベースにクエリを実行できます。
会議コレクションにはテナントの分離がありません。認証された tl;dv ユーザーは、プラットフォーム上のすべてのアカウントですべての会議をクエリできます。各会議記録には、作成者のメール アドレス、会議 ID (参加可能な Google Meet または Teams ルーム)、プロバイダー、録画ステータス、タイムスタンプが含まれます。
録画ステータスの会議の場合、その会議 ID はライブのアクティブな通話です。コレクションをリアルタイムで監視したり、会議の録画開始を確認したり、ID を取得したり、招待されていないのに誰かの通話に参加したりすることができます。常に、コレクション内に記録中というステータスの会議が約 1,000 件あります。公開された会議 ID を使用した 1,000 件のライブ コール。ボットを使った攻撃者が参加する可能性がある

それらすべてを同時に。
Firestore からカンファレンス ID を取得し、マレーシア教育省が所有するライブ Google Meet に参加しました。 157名以上の参加者を前に、ある女性が講演を行っていました。 tl;dv ボットはすでに参加者リストに含まれていました。私も同じ通話中でした。誰も私を誘ってくれませんでした。 Firestore データベースはそうでした。
私はまた、米国の主要大学の学生がスタートアップ アプリを構築している会議にも参加しました。通話参加者は21名。彼らはプロジェクト全体を画面共有し、プロトタイプについて話し合ったり、冗談ではありませんが、.edu メール アドレスにクライアント側の検証を追加する必要があるかどうかについて話し合ったりしていました。彼らはまた、画面上で Supabase をライブでセットアップしていましたが、ほとんどの人はセットアップしないので、私は「RLS ポリシーをセットアップしてください」としか考えられませんでしたが、結局は tl;dv のようになってしまいました。
すごく言いたかったのです。 「サーバー側の検証も必要かもしれません。」しかし、これはコンセプトの実証であり、相談ではありませんでした。
Firestore 会議コレクションにクエリを実行したところ、35,003 の電子メール ドメインにわたる 84,312 人のユニーク ユーザーに属する 181,874 件の会議記録があることがわかりました。
23か国の政府会合：ブラジル、コロンビア、ペルー、ウクライナ、エルサルバドル、フィリピン、チリ、インドネシア、メキシコ、米国、カタール、マレーシア、ウズベキスタン、スリランカ、ハイチ、南アフリカ、ジャマイカ、ホンジュラス、アルゼンチン、タイ、日本、イスラエル、ベリーズ。すべての .gov ドメイン。無料利用枠のユーザーがすべてを列挙できるプラットフォームで通話を録音する政府職員。
バークレー、東京大学、デ・ラ・サール、コロンビア国立大学からの大学の会合。数十の .edu および .ac ドメイン。
残りの 35,000 ドメインすべてからの企業会議。三井倉庫 (4 つの地方事務所で 484 回の会議)、三井不動産、HubSpot、Confluent、Mekari、AnyMind Group。これまで私たちに関わってきたすべての企業

ed tl;dv は、同じ保護されていないコレクションに会議メタデータを持っていました。
ピーク月は 2025 年 7 月で、会議数は 43,209 件でした。最も混雑する時間帯: 水曜日の午後 2 時 (協定世界時)、7,804 件の会議。こぶの日のスタンドアップアワー。
実際のコンテンツにどれだけアクセスできるのかも知りたかったのですが、デフォルトでは会議は非公開になっています（つまり、ビデオを見たりトランスクリプトを確認したりすることはできません）。そのため、27,334 件の会議 ID を収集し、どれが公開されているかを確認しました。 1,000以上ありました。 715 件の招待者の電子メールが 228 のドメインにわたって公開されました。
ハイライト: WWF、ザ・ネイチャー・コンサーバンシー、コンサベーション・インターナショナル、WRI、サンパウロ州政府が参加するブラジル政府の自然保護会議（PACTO Mata Atlântica）。ウクライナデジタル変革省からの会合。 HubSpot の営業電話。コロンビア国立大学とチリのカマラ・ベルデが参加したセッション。
tl;dv はマイクロサービスにパスタの名前を付けています。サブドメイン スキャンにより、 cappellini 、carbonara 、fusilli 、pasta 、penne 、puttanesca-v0 、および ravioli がすべて tldv.io の下にあることがわかります。イタリア料理レストラン全体に相当する Express サーバー。
サブドメインを探索しているときに https://worldcup.tldv.io を見つけました。 tl;dv 従業員向けに Base44 に基づいて構築された 2026 FIFA ワールド カップのバイブコーディングされた予測ゲーム。これは「World Cup Pick'em」と呼ばれ、内部チームの名前は「Too Long; Didn't Score」です。かわいい。
Player エンティティ API には認証がありません。 GET /api/entities/Player は、セッション Cookie なしですべてのプレーヤー レコードを返します。選手数は43名。 19 人の @tldv.io 従業員 (フルネームと会社の電子メールを含む)。
私の情報開示担当者であるラファエル・アルシュタット氏は、曖昧な安心感を与えた後、沈黙してしまい、298 ポイントで 2 位になりました。 API 応答には彼の個人的な Gmail も含まれていました。グローバルリーダーボードのプレイヤー #5 は「Super Duper CEO」です。それが誰なのかはご想像にお任せします。
予測

フィクスチャ エンティティも幅広く対応しています。何百万もの人々の会議を記録している会社は、社内の楽しいアプリをバイブコーディングして、自社の従業員名簿を漏洩させました。皮肉はアルデンテだ。
1 月 28 日、私は LinkedIn で Raphael Allstadt にメッセージを送り、ユーザー データを漏洩する巨大な脆弱性を発見したことを伝えました。彼は数分以内に「ありがとうございます。弊社の CTO に報告していただけますか。すぐに検討させていただきます。」と返信しました。メールを送りました。彼は「ありがとう！」と言いました。報酬について聞いてみた。 「私の CTO が戻ってきます」と彼は言いました。
CTO は私のところに戻ってきませんでした。
1 月 29 日: 「ところで、あなたの最高技術責任者はまだ連絡を取っておらず、問題は解決していません。」 1 月 30 日、ラファエル: 「チームがすぐにレビューすると確信しています ❤️」 2 月 14 日: 「メールは届いていませんが、脆弱性はまだ機能しています。」ラファエル: 「彼は戻ってきます ☺️」私は彼に、おそらく脆弱性を修正し、顧客を危険にさらしたままにしないように言いました。 2 月 19 日: 「取り組んでいます。少し時間がかかりますが、最後までやり遂げますのでご安心ください。さらに連絡が必要な場合は、CTO に連絡することをお勧めします。」
返答しなかったCTO。あのCTO。
3月6日：「まだ直っていない」。ラファエルが午後5時42分に目撃。返事はありません。
7月22日：「まだ直ってない…」 返事なし。
彼らのセキュリティページはトロフィーケースです。 SOC2準拠。 GDPR準拠。 EU AI 法に準拠。 EUで主催。 AES-256暗号化。創業者のコミットメントビデオ。 6 つのコンプライアンスバッジが横一列に並んでいます。一番下には、「対処すべきプライバシーまたはセキュリティの問題を発見した場合は、いつでも [email protected] までお知らせください。当社のセキュリティ チームが 24 時間以内に対応します。」という 1 行が埋め込まれています。 CTOに直接メールしました。 6か月。応答がありません。 Firestore データベースの稼働時間は受信トレイよりも優れています。
あなたのプラットフォームは人々の最も機密性の高い会話を記録します。就職面接。販売

交渉。政府の説明会。ユーザーは、記録することに明示的に同意したコンテンツについてあなたを信頼していました。
Firestore テナントの分離を修正します。これには Firestore セキュリティ ルールが存在します。他のすべてのコレクションではすでに正しく実行されています (ユーザー、チャット、トランスクリプト、クリップ、録音、ビデオ、メモ、チーム、組織はすべて 403 を返します)。会議を忘れただけです。
ワールドカップアプリを認証するか、削除してください。従業員ディレクトリは GET リクエスト 1 つで取得できます。
セキュリティ研究者に回答します。特に、あなたのプラットフォーム上のすべての会議は、無料アカウントを持つ誰でもクエリ可能であると言われている場合はそうです。
長い間、パスタをありがとう :3

## Original Extract

How a missing Firestore security rule on tl;dv exposed 181,874 meetings from 84,312 users across 35,003 domains, including live calls I could join uninvited, and how six months of disclosure got me nothing but seen receipts.

tl;dv (Too Lazy; Didn't Validate): 181,874 Meetings Left Wide Open
August 4, 2026
BobDaHacker
I reported this on January 28th, 2026. It is now July 2026. Six months later. The Firestore database is still wide open. The CTO never responded. I guess my emails were too long and they didn't view them.
tl;dv (Too Long; Didn't View) is an AI meeting recording platform. It drops a bot into your Google Meet, Zoom, or Teams call, records everything, transcribes it, and generates summaries with AI. Over 2 million users. Backed by investors. Endorsed by half of LinkedIn's sales influencer community.
They store your sales calls, job interviews, performance reviews, internal strategy sessions. The kind of content where someone says "this call is being recorded" and everyone nervously laughs and then shares trade secrets for 45 minutes.
When you sign up for tl;dv, the platform authenticates you with a JWT and exchanges it for a Firebase token via gw.tldv.io/v1/users/firebase/token . That token lets you query their Firestore database at projects/lmi-store/databases/(default) .
The meetings collection has no tenant isolation. Any authenticated tl;dv user can query every meeting across every account on the platform. Each meeting record hands you the creator's email address, the conference ID (which is a joinable Google Meet or Teams room), the provider, the recording status, and timestamps.
For meetings in recording status, that conference ID is a live, active call. You can watch the collection in real time, see a meeting start recording, grab the ID, and walk into someone's call uninvited. At any given time there are roughly 1,000 meetings with status: recording sitting in the collection. A thousand live calls with exposed conference IDs. An attacker with a bot could join all of them simultaneously.
Grabbed a conference ID from Firestore and joined a live Google Meet belonging to the Malaysian Ministry of Education . A lady was presenting to over 157 participants. The tl;dv bot was already in the participant list. I was in the same call. Nobody invited me. The Firestore database did.
I also joined a call where students from a major US university were building a startup app. 21 people in the call. They were screen-sharing their entire project, discussing prototypes, and, I kid you not, talking about how they needed to add client-side validation for .edu email addresses. They were also setting up Supabase live on screen, and all I could think was "please set up RLS policies" because most people don't, and then you end up like tl;dv.
I wanted to say something so badly. "Hey, you might want server-side validation too." But this was a proof of concept, not a consultation.
I queried the Firestore meetings collection and saw there were 181,874 meeting records belonging to 84,312 unique users across 35,003 email domains .
Government meetings from 23 countries : Brazil, Colombia, Peru, Ukraine, El Salvador, the Philippines, Chile, Indonesia, Mexico, the United States, Qatar, Malaysia, Uzbekistan, Sri Lanka, Haiti, South Africa, Jamaica, Honduras, Argentina, Thailand, Japan, Israel, and Belize. All .gov domains. Government employees recording calls on a platform that lets any free-tier user enumerate the whole thing.
University meetings from Berkeley, the University of Tokyo, De La Salle, Universidad Nacional de Colombia. Dozens of .edu and .ac domains.
Corporate meetings from all 35,000 remaining domains. Mitsui-Soko (484 meetings across four regional offices), Mitsui Fudosan, HubSpot, Confluent, Mekari, AnyMind Group. Every company that ever used tl;dv had their meeting metadata in the same unprotected collection.
Peak month was July 2025 with 43,209 meetings . Busiest time slot: Wednesday at 2pm UTC , 7,804 meetings. Hump-day standup hour.
I wanted to know how much actual content was accessible too, by default meetings are private (Meaning you cant watch the video or see the transcript), so I scraped 27,334 meeting IDs and checked which ones were public. Over 1,000 were. 715 invitee emails exposed across 228 domains .
Highlights: a Brazilian government conservation meeting (PACTO Mata Atlântica) with participants from WWF, The Nature Conservancy, Conservation International, WRI, and the São Paulo state government. Meetings from Ukraine's Ministry of Digital Transformation. A HubSpot sales call. Sessions involving Universidad Nacional de Colombia and Chile's Cámara Verde.
tl;dv names their microservices after pasta. A subdomain scan reveals cappellini , carbonara , fusilli , pasta , penne , puttanesca-v0 , and ravioli , all under tldv.io. An entire Italian restaurant worth of Express servers.
While exploring their subdomains I found https://worldcup.tldv.io . A FIFA World Cup 2026 vibecoded prediction game built on Base44 for tl;dv employees. It's called "World Cup Pick'em" and their internal squad is named "Too Long; Didn't Score." Cute.
The Player entity API has zero authentication. GET /api/entities/Player returns every player record without a session cookie. 43 players. 19 @tldv.io employees with full names and corporate emails.
Raphael Allstadt, my disclosure contact who gave vague reassurances and then went quiet, came in 2nd place with 298 points. His personal Gmail was also in the API response. Player #5 on the global leaderboard is "Super Duper CEO." I'll let you guess who that is.
The Prediction and Fixture entities are also wide open. A company that records millions of people's meetings vibecoded an internal fun app that leaks their own employee directory. The irony is al dente.
On January 28th I messaged Raphael Allstadt on LinkedIn and told him I'd found a huge vulnerability that leaks user data. He responded within minutes: "thank you! can you report it to our CTO and we will look at it immediately?" I sent the email. He said "thank you!" I asked about a reward. "My CTO will come back to you," he said.
The CTO never came back to me.
January 29th: "your cto hasnt reached out yet btw and its not fixed." January 30th, Raphael: "I am sure the team is reviewing it very very soon ❤️" February 14th: "havent got an email and the vulnerability stilll works." Raphael: "He'll come back ☺️" I told him to maybe fix the vulnerability and not leave customers exposed. February 19th: "We're on it. It needs some time, but rest assured we're following through. For further communication, i'll recommend reaching out to our CTO."
The CTO who never responded. That CTO.
March 6th: "still not fixed." Seen by Raphael at 5:42 PM. No reply.
July 22nd: "still not fixed..." No reply.
Their security page is a trophy case. SOC2 compliant. GDPR compliant. EU AI Act compliant. Hosted in the EU. AES-256 encryption. A founder commitment video. Six compliance badges lined up in a row. Buried at the bottom, a single line: "If you have discovered a privacy or security issue that we should address, please always let us know at [email protected] . Our security team will respond within 24 hours." I emailed the CTO directly. Six months. No response. Their Firestore database has better uptime than their inbox.
Your platform records people's most sensitive conversations. Job interviews. Sales negotiations. Government briefings. Your users trusted you with content they explicitly consented to record.
Fix the Firestore tenant isolation. Firestore security rules exist for this. You already do it correctly for every other collection (users, chats, transcripts, clips, recordings, videos, notes, teams, organizations all return 403). You just forgot meetings.
Put auth on the World Cup app or take it down. Your employee directory is one GET request away.
Respond to security researchers. Especially when they're telling you that every meeting on your platform is queryable by anyone with a free account.
So long and thanks for all the pasta :3
