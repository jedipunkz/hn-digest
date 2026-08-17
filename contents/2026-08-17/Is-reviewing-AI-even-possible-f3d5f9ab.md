---
source: "https://www.colino.net/wordpress/archives/2026/08/16/is-reviewing-ai-even-possible/"
hn_url: "https://news.ycombinator.com/item?id=49336725"
title: "Is reviewing AI even possible?"
article_title: "Is reviewing AI even possible? – colin@colino.net"
image: "https://www.colino.net/wordpress/wp-content/uploads/Screenshot_2026-08-16_18-30-21.png"
author: "ibobev"
captured_at: "2026-08-17T20:15:42Z"
capture_tool: "hn-digest"
hn_id: 49336725
score: 2
comments: 0
posted_at: "2026-08-17T20:00:30Z"
tags:
  - hacker-news
  - translated
---

# Is reviewing AI even possible?

- HN: [49336725](https://news.ycombinator.com/item?id=49336725)
- Source: [www.colino.net](https://www.colino.net/wordpress/archives/2026/08/16/is-reviewing-ai-even-possible/)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T20:00:30Z

## Translation

タイトル: AI のレビューは可能ですか?
記事タイトル：AIの見直しは可能なのか？ –colin@colino.net
説明: それはどうすればいいですか?

記事本文:
AIのレビューは可能でしょうか？ –colin@colino.net
コンテンツにスキップ
コリン@colino.net
AIのレビューは可能でしょうか？
今日の初めに、私は課題を共有しました。これらのサムネイルのデータ エンコーディングがわかりませんでした。画像処理に詳しい人がアイデアをくれることを期待していました。また、誰かがそれを AI に与えて結果を私に投げ捨てるのではないかと心配していました。
まあ、両方とも起こりました！ヘンリーという人がストレージ形式を見つけて私に説明してくれたので、私は彼の説明に基づいてアルゴリズムを更新しました。彼の説明は、「実際には 80*60 ではなく 40*30 ピクセルであり、各ピクセルを 16 ビット、RGGB (4/8/4 ビット) で、40*RGG、次に 40*B からなる 80 バイトのストライドでエンコードします」でした。
別の人がクロードに問題を提出し、結果のコードを私にくれました。
コードは機能します。何らかの理由でいくつかのオプションも実装されていますが、私はそれを使用しませんでした。もしヘンリーが私に適切な説明をしてくれなかったら、私は今でもクロードが吐き出したコードの塊を理解しようとしていたでしょう。
私はクロードのコードを、機能する最低限の部分まで分解して、12 個の付加機能を持たないものにしようとしましたが、理由もなく恐ろしく複雑でした。 AI を使用し、その出力をレビューしていると主張する人々がどのようにそれを行っているのか、本当に不思議に思います。これは自分で行うよりもはるかに手間がかかります。
このアルゴリズムはピクセルあたり 4 ビットを誤って想定していましたが、ストライド (60/20 バイト) は正しく得られました。
void render_thumbnail(FILE *fp, int w, int h, SDL_Surface *screen) {
unsigned char i、x、y;
文字入力バイト[80];
char full_bytes[160];
/* 80 バイトで 2 行をエンコードします。ピクセルは 2x2 のブロックで記述されます。
* ブロックごとに 2 バイト (4 ニブル)、奇妙なレイアウト。 */
for (y = 0; y < 60; y+=2) {
/* 2 行分のデータを読み込みます */
fread(input_bytes, 1, 80, fp);
/* まずニブルを変換します

簡単にするために、s をフルバイトに変更します。
* これらは 0 ～ 15 ですが、使用可能な値になるように << 4 にシフトします。 */
for (i = 0; i < w; i++) {
unsigned char c = input_bytes[i];
unsigned char high_nibble = c & 0xF0;
unsigned char low_nibble = c << 4;
full_bytes[i*2] = high_nibble;
full_bytes[i*2 + 1] = low_nibble;
}
/* ここからは楽しい部分です。 160 バイトがあり、2 行をエンコードしています。私は
* このように、2x2 のブロックでエンコードされていることがわかりました。
* 各ブロックは ab/cd で示されます。
*
* 0 <--- 線幅 --> 79
* アバババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババババ
* cdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcd
*
* これら 4 つの a/b/c/d 値は、次のように入力バッファ (160 バイトの大きさ) から取得されます。
* abcabcabcabc...abcddd...ddd
* ここで、最初の 120 バイトには 40 個の a、b、c のトリプレットが含まれ、最後の 40 バイトには 40 個の d が含まれます。
*
* しかし。出力がきれいではないため、160 の値がピクセル値を直接エンコードしているとは思えません。
※必ず変換が必要です。
*
*/
for (i = 0, x = 0; i < (w*3)/2; i += 3, x += 2) {
unsigned char ピクセル_a、ピクセル_b、ピクセル_c、ピクセル_d;
ピクセル_a = フル_バイト[i];
ピクセル_b = フル_バイト[i+1];
ピクセル_c = フル_バイト[i+2];
ピクセル_d = フル_バイト[120 + i/3];
PIXEL_OUTPUT(x, y, ピクセル_a);
PIXEL_OUTPUT(x+1, y, ピクセル_b);
PIXEL_OUTPUT(x, y+1, ピクセル_c);
PIXEL_OUTPUT(x+1, y+1, ピクセル_d);
}
}
}
正しいアルゴリズム
このバージョンではサムネイルが正しくデコードされます。誰もそれ以上何も要求しなかったので、それ以上何もしません。
void render_thumbnail(FILE *fp, int w, int h, SDL_Surface *screen) {
unsigned char i、j、x、y;
文字入力バイト[80];
/* 80 バイトは 1 つの 40 ピクセル ラインをエンコードします。ピクセルは次のブロックで記述されます。
* 16 ビット、奇妙なレイアウト。 */
for (y = 0; y < 60; y+=2) {
/* 2 行分の内容を読み取ります

データ */
fread(input_bytes, 1, 80, fp);
/* ここからは楽しい部分です。 80 バイトがあり、40 ピクセルをエンコードします。彼らは
* は RGGB (4/8/4 ビット) としてエンコードされ、40*RGG、次に 40*B の順に並べられます。
*/
for (i = 0, x = 0; i < 60;) {
unsigned char r、g、b;
r = (入力バイト[i] & 0xF0);
g = (入力バイト[i] & 0x0F) | (input_bytes[i+1] & 0xF0);
b = (入力バイト[60 + i/3] & 0x0F) << 4;
PIXEL_OUTPUT(x, y, r, g, g);
PIXEL_OUTPUT(x+1, y, r, g, g);
PIXEL_OUTPUT(x, y+1, r, g, g);
PIXEL_OUTPUT(x+1, y+1, r, g, g);
i += 2;
x += 2;
r = (入力バイト[i-1] & 0x0F) << 4;
g = 入力バイト[i];
b = (入力バイト[60 + i/3] & 0xF0);
PIXEL_OUTPUT(x, y, r, g, g);
PIXEL_OUTPUT(x+1, y, r, g, g);
PIXEL_OUTPUT(x, y+1, r, g, g);
PIXEL_OUTPUT(x+1, y+1, r, g, g);
i += 1;
x += 2;
}
}
}
クロードのアルゴリズム
クロードのアルゴリズムも機能します。正しいフラグを指定すると、正しいアルゴリズムと漠然と同じサムネイルが出力されます (色がより飽和しています)。他のサムネイルでテストしていないので、それが正しいかどうかさえわかりません。それはその潜在的な正しさを明らかにしません。さまざまなフラグを使用してさまざまなバリエーションを出力します。 SDL ディスプレイを PNG ライターに置き換えました。申し訳ありませんが、これは 50 行未満の単一の関数には含まれていません。このペーストから配管、main()、および PNG ライターを削除しました。
#define QW 40 /* クワッド全体 */
#define QH 30 /* クワッドダウン */
#define NQUADS (QW * QH)
#define IW (QW * 2) /* 80 個の輝度サンプル */
#define IH (QH * 2) /* 60 ルミナンス サンプル ダウン */
#定義THUMB_BYTES 2400
/* プレーン[0..3] = すべてのクワッドの a、b、c、d サンプル (それぞれ 0..15)。
* a = 左上、b = 右上、c = 左下、d = 右下。 */
静的 unsigned char プレーン[4][NQUADS];
/* 4 つのフィルター位置の相対感度。各平面をその平面で割ると、
* のゲインと減算

fset は 2x2 チェッカーボードを削除するものです。 */
static const double CFA_GAIN[4] = { 1.118, 0.979, 0.992, 0.910 };
static const double CFA_OFFSET[4] = { 0.482, -0.267, -0.480, 0.265 };
/* 色補正行列: (a, b, c, d, 1) からのディスプレイ参照 R、G、B。
* 1200 個のクワッドに対する最小二乗法を当てはめました。 R^2 = 0.937、0.887、0.786。
※反対側の大きな係数は上記のクロマ増幅です。 */
静的 const double CCM[3][5] = {
{ 15.602, -7.763, -0.626, 8.720, 5.059 }, /* R */
{ -8.812, 15.116, -1.019, 8.402, 22.012 }, /* G */
{ -11.359, 9.929, 6.424, 6.018, 4.753 } /* B */
};
/* ------------------------------------------------------------------ */
/* 解凍 -- この権限を持つオリジナルから変更なし */
/* ------------------------------------------------------------------ */
static int unpack(const char *path)
{
unsigned char バンド[80]、nib[160];
ファイル*fp;
長いサイズ。
int b、i、k、q = 0;
fp = fopen(パス, "rb");
if (!fp) {
fprintf(stderr, "%s (%s) を開けません\n", path, strerror(errno));
-1 を返します。
}
if (fseek(fp, 0, SEEK_END) == 0) {
sz = ftell(fp);
if (sz != THUMB_BYTES)
fprintf(stderr, "警告: %s は %ld バイト、期待値 %d\n",
パス、サイズ、THUMB_BYTES);
巻き戻し(fp);
}
for (b = 0; b < QH; b++) {
if (fread(バンド, 1, 80, fp) != 80) {
fprintf(stderr, "バンド %d でショート読み取り\n", b);
fclose(fp);
-1 を返します。
}
/* 80 バイト -> 160 ニブル、上位ニブルが最初 */
for (i = 0; i < 80; i++) {
nib[i * 2] = (バンド[i] >> 4) & 0x0F;
nib[i * 2 + 1] = バンド[i] & 0x0F;
}
/* ニブル 0..119 は
[切り捨てられた]
比較のために、ズームアウトしたエディターのスクリーンショット。左側のペインで私のアルゴリズムが選択され、右側のペインでクロードのアルゴリズムが選択されています。
時期尚早に最適化しないでください。ただし、
Apple での Code 2022 の到来 //c
Apple II 用グライダー開発ログ
前回のサムネイル デコード チャレンジ
あなた

r メールアドレスは公開されません。 * が付いているフィールドは必須です
このサイトはスパムを低減するために Akismet を使っています。コメントデータがどのように処理されるかをご覧ください。
colin@colino.net © 2026.全著作権所有。
Powered by WordPress、テーマは Alx Boxcard から派生、ステータス ページは Updown から提供されます。
一部の人だけのニュース、誰も気にしないこと
「AI」ゼロで作られています。
Apple II プロジェクト Toggle Child Menu
Apple II のマストドン
BurgerDisk – Apple II Smartport ハードドライブ、デイジーチェーン接続可能
SixForty – Apple II の写真共有
surl-server – 8 ビット コンピュータ用のシリアル プロキシ

## Original Extract

What am I supposed to do with that?

Is reviewing AI even possible? – colin@colino.net
Skip to content
colin@colino.net
Is reviewing AI even possible?
Earlier today, I shared a challenge . I couldn’t figure out the data encoding in those thumbnails, and I hoped someone well versed in image processing would have ideas. I also feared someone would feed it to an AI and dump me the results.
Well, both happened! One person, Henry, figured out and explained the storage format to me, and I updated my algorithm from his explanation. His explanation was: “It’s actually 40*30 pixels instead of 80*60; and it encodes each pixel in 16 bits, RGGB (4/8/4 bits), in strides of 80 bytes consisting of 40*RGG then 40*B”.
Another person submitted the problem to Claude, and gave me the resulting code.
The code works. It also has implemented a few options, for some reason… but I didn’t use it. I would still be trying to understand the blob of code that Claude regurgitated if Henry hadn’t given me a good explanation.
I tried dismantling Claude’s code to the bare minimum that would work and not have twelve bells and whistles, but it was horribly complicated for absolutely no reason. That really makes me wonder how people who use AI and claim to review its output do it. This is much more work than doing it oneself.
This algorithm was wrongly assuming 4 bits per pixels, but got the stride (60/20 bytes) right:
void render_thumbnail(FILE *fp, int w, int h, SDL_Surface *screen) {
unsigned char i, x, y;
char input_bytes[80];
char full_bytes[160];
/* 80 bytes encode two lines. Pixels are described in blocks of 2x2, using
* two bytes (four nibbles) per block, with a weird layout. */
for (y = 0; y < 60; y+=2) {
/* Read two lines' worth of data */
fread(input_bytes, 1, 80, fp);
/* First of all, convert nibbles to full bytes, for simplicity's sake.
* they're 0-15, we'll shift them << 4 so that they have a usable value. */
for (i = 0; i < w; i++) {
unsigned char c = input_bytes[i];
unsigned char high_nibble = c & 0xF0;
unsigned char low_nibble = c << 4;
full_bytes[i*2] = high_nibble;
full_bytes[i*2 + 1] = low_nibble;
}
/* Now for the fun part. We have 160 bytes, encoding two lines. I
* figured out that they're encoded in blocks of 2x2, like this,
* where each block is indicated as ab/cd:
*
* 0 <--- line width --> 79
* abababababababababababababababababababababababababababababababababababababababab even line
* cdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcd odd line
*
* Those four a/b/c/d values come from the input buffer (160 bytes large) as follows:
* abcabcabcabc...abcddd...ddd
* where the first 120 bytes contain the 40 a,b,c triplets and the last 40 bytes contain the 40 d.
*
* But. I don't think the 160 values directly encode pixel values, as the output is not clean.
* There must be a transformation.
*
*/
for (i = 0, x = 0; i < (w*3)/2; i += 3, x += 2) {
unsigned char pixel_a, pixel_b, pixel_c, pixel_d;
pixel_a = full_bytes[i];
pixel_b = full_bytes[i+1];
pixel_c = full_bytes[i+2];
pixel_d = full_bytes[120 + i/3];
PIXEL_OUTPUT(x, y, pixel_a);
PIXEL_OUTPUT(x+1, y, pixel_b);
PIXEL_OUTPUT(x, y+1, pixel_c);
PIXEL_OUTPUT(x+1, y+1, pixel_d);
}
}
}
The correct algorithm
This version decodes a thumbnail correctly. It does nothing more as nobody asked it to do anything more.
void render_thumbnail(FILE *fp, int w, int h, SDL_Surface *screen) {
unsigned char i, j, x, y;
char input_bytes[80];
/* 80 bytes encode one 40 pixels line. Pixels are described in blocks of
* 16 bits, with a weird layout. */
for (y = 0; y < 60; y+=2) {
/* Read two lines' worth of data */
fread(input_bytes, 1, 80, fp);
/* Now for the fun part. We have 80 bytes, encoding 40 pixels. They
* are encoded as RGGB (4/8/4 bits), and ordered as 40*RGG then 40*B.
*/
for (i = 0, x = 0; i < 60;) {
unsigned char r, g, b;
r = (input_bytes[i] & 0xF0);
g = (input_bytes[i] & 0x0F) | (input_bytes[i+1] & 0xF0);
b = (input_bytes[60 + i/3] & 0x0F) << 4;
PIXEL_OUTPUT(x, y, r, g, g);
PIXEL_OUTPUT(x+1, y, r, g, g);
PIXEL_OUTPUT(x, y+1, r, g, g);
PIXEL_OUTPUT(x+1, y+1, r, g, g);
i += 2;
x += 2;
r = (input_bytes[i-1] & 0x0F) << 4;
g = input_bytes[i];
b = (input_bytes[60 + i/3] & 0xF0);
PIXEL_OUTPUT(x, y, r, g, g);
PIXEL_OUTPUT(x+1, y, r, g, g);
PIXEL_OUTPUT(x, y+1, r, g, g);
PIXEL_OUTPUT(x+1, y+1, r, g, g);
i += 1;
x += 2;
}
}
}
Claude’s algorithm
Claude’s algorithm also works. Given the correct flags, it does output vaguely the same thumbnail as the correct algorithm (the colors are more saturated). I am not even sure it is correct, as I didn’t test it with other thumbnails. It does not make its potential correctness clear. It outputs different variations with different flags. It replaced the SDL display with a PNG writer. And I am sorry to inform you that it is not contained in a single, < 50 lines function. I did remove the plumbing, main(), and the PNG writer from this paste:
#define QW 40 /* quads across */
#define QH 30 /* quads down */
#define NQUADS (QW * QH)
#define IW (QW * 2) /* 80 luma samples across */
#define IH (QH * 2) /* 60 luma samples down */
#define THUMB_BYTES 2400
/* plane[0..3] = the a, b, c, d sample of every quad, each 0..15.
* a = top-left, b = top-right, c = bottom-left, d = bottom-right. */
static unsigned char plane[4][NQUADS];
/* Relative sensitivity of the four filter positions. Dividing each plane by its
* gain and subtracting its offset is what removes the 2x2 checkerboard. */
static const double CFA_GAIN[4] = { 1.118, 0.979, 0.992, 0.910 };
static const double CFA_OFFSET[4] = { 0.482, -0.267, -0.480, 0.265 };
/* Colour correction matrix: display-referred R, G, B from (a, b, c, d, 1).
* Fitted least-squares over 1200 quads; R^2 = 0.937, 0.887, 0.786.
* The large opposing coefficients are the chroma amplification described above. */
static const double CCM[3][5] = {
{ 15.602, -7.763, -0.626, 8.720, 5.059 }, /* R */
{ -8.812, 15.116, -1.019, 8.402, 22.012 }, /* G */
{ -11.359, 9.929, 6.424, 6.018, 4.753 } /* B */
};
/* ------------------------------------------------------------------ */
/* Unpacking -- unchanged from the original, which had this right */
/* ------------------------------------------------------------------ */
static int unpack(const char *path)
{
unsigned char band[80], nib[160];
FILE *fp;
long sz;
int b, i, k, q = 0;
fp = fopen(path, "rb");
if (!fp) {
fprintf(stderr, "can't open %s (%s)\n", path, strerror(errno));
return -1;
}
if (fseek(fp, 0, SEEK_END) == 0) {
sz = ftell(fp);
if (sz != THUMB_BYTES)
fprintf(stderr, "warning: %s is %ld bytes, expected %d\n",
path, sz, THUMB_BYTES);
rewind(fp);
}
for (b = 0; b < QH; b++) {
if (fread(band, 1, 80, fp) != 80) {
fprintf(stderr, "short read at band %d\n", b);
fclose(fp);
return -1;
}
/* 80 bytes -> 160 nibbles, high nibble first */
for (i = 0; i < 80; i++) {
nib[i * 2] = (band[i] >> 4) & 0x0F;
nib[i * 2 + 1] = band[i] & 0x0F;
}
/* nibbles 0..119 are
[truncated]
For comparison, a screenshot of my editor, zoomed out, with my algorithm selected in the left pane and Claude’s in the right pane:
Don’t optimize prematurely, BUT
Advent of Code 2022 on the Apple //c
Glider for Apple II development log
Previous Thumbnail decoding challenge
Your email address will not be published. Required fields are marked *
This site uses Akismet to reduce spam. Learn how your comment data is processed.
colin@colino.net © 2026. All Rights Reserved.
Powered by WordPress , Theme derived from Alx Boxcard , Status page by Updown .
News for few, stuff no-one cares about
Made with zero "AI".
Apple II projects Toggle Child Menu
Mastodon for Apple II
BurgerDisk – an Apple II Smartport hard drive, daisy-chainable
SixForty – photo sharing for the Apple II
surl-server – a serial proxy for 8bit computers
