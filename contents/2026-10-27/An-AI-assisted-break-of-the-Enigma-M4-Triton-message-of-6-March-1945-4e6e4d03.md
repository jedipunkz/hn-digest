---
source: "https://cryptocellar.org/bgac/the-m4-triton-break.html"
hn_url: "https://news.ycombinator.com/item?id=50258672"
title: "An AI-assisted break of the Enigma M4 Triton message of 6 March 1945"
article_title: "The Enigma M4 Triton Break"
image: ""
author: "makmanalp"
captured_at: "2026-10-27T14:02:53Z"
capture_tool: "hn-digest"
hn_id: 50258672
score: 1
comments: 0
posted_at: "2026-10-27T13:10:37Z"
tags:
  - hacker-news
---

# An AI-assisted break of the Enigma M4 Triton message of 6 March 1945

- HN: [50258672](https://news.ycombinator.com/item?id=50258672)
- Source: [cryptocellar.org](https://cryptocellar.org/bgac/the-m4-triton-break.html)
- Score: 1
- Comments: 0
- Posted: 2026-10-27T13:10:37Z

## Translation

Title: An AI-assisted break of the Enigma M4 Triton message of 6 March 1945
Article title: The Enigma M4 Triton Break

Article text:
An AI-assisted break of the Enigma M4 Triton message of 6 March 1945
In 2006, Ralph Erskine published three German Naval Enigma messages, M4 messages in the keys Shark, Turtle and Pike , enciphered on the Naval Enigma M4, as another contribution to the then ongoing M4 Message Breaking Project . They were then busy trying to break the four M4 messages intercepted by H.M.S. Hurricane in the Atlantic in November 1942 and published by Ralph Erskine in a Cryptologia article in October 1996. The first two of the 1942 messages were broken by the members of the M4 Message Breaking Project in February and March 2006, while the third message was broken by Dan Girard in January 2013.
If anyone attempted to break the three M4 messages in the keys of Shark, Turtle and Pike, it is not known to me. The file with the three messages from October 1943, January 1944, and March 1945 has been sitting in Crypto Cellar Research since March 2006, and I must admit I had forgotten about them.
I was therefore greatly surprised when Jason Parker-Burlingham contacted me on 25 October 2026 to announce that he had broken the first of the three messages, the Shark (Triton) signal of 6 March 1945.
Here is his report:
I have decrypted the 6 March 1945 Shark signal (“CXSO UNVZ NEWM ESUC IIFP RLEE TPNZ LPYJ WDKL LDUJ”), first on the page at https://cryptocellar.org/bgac/m4-messages.html .
The successful crib was NULVIERDREINULBISNULSECHSDREINUL (“0430 to 0630”), 123rd in a ranked list; a second crib, LAUFENNACHPERIODEJBRUNOJ, gave the same key independently.
Plaintext, with no corrections applied or encoding undone:
EULA UFEN NACH PERI ODEJ BRUN OJXX EINS KKFU NKFE
UERG EUPP EEIN SEIN SAMS IEBE NXMA ERZV ONEI NSBI
SDRE IGEI NSZW OBIS EINS VIER YACH TXNU LVIE RDRE
INUL BISN ULSE CHSD REIN ULYN EUNX EINS NULB ISEI
NSDW OYZW OZWO BISZ WOVG ERXZ WOKK FUNK FEUE RGRU
PPEE CNSZ WVAM SIEB ENRM JERZ VONE INSB SSDR EIYA
CHTX NULV IERD REIN ULBI SNUL SECH SBRE INUL YEIN
SVIE RIRE INUL BISE IOSS ECHS DREI NZLY NEUN XKKF
ORTK KEIN SNUL BISE IESZ WOBV
Plaintext, with words separated, punctuation applied and 14 letters corrected:
ES LAUFEN NACH PERIODE "BRUNO": EINS) FUNKFEUERGRUPPE EINSEINS AM
SIEBEN. MAERZ VON EINS BIS DREI, EINSZWO BIS EINSVIER; ACHT.
NULVIERDREINUL BIS NULSECHSDREINUL; NEUN. EINSNUL BIS EINSZWO,
ZWOZWO BIS ZWOVIER. ZWO) FUNKFEUERGRUPPE EINSZWO AM SIEBEN.
MAERZ VON EINS BIS DREI; ACHT. NULVIERDREINUL BIS NULSECHSDREINUL,
EINSVIERDREINUL BIS EINSSECHSDREINUL; NEUN. (FORT) EINSNUL BIS
EINSZWO
In English:
"The following will operate according to period 'Bruno':
1) Radio-beacon group 11: on 7 March from 0100 to 0300 and 1200 to 1400;
on the 8th, 0430 to 0630; on the 9th, 1000 to 1200 and 2200 to 2400.
2) Radio-beacon group 12: on 7 March from 0100 to 0300;
on the 8th, 0430 to 0630 and 1430 to 1630;
on the 9th, (continued) 1000 to 1200 …"
Machine settings:
M4, thin reflector B, Greek wheel gamma, rotors IV V III,
rings A A P U, start Z K S A, plugs AX BZ CY DG ET IP JN OQ RV SW.
The parts that made up the solution were as follows, given in more or less the order in which I prompted the model to build and run them:
1) collecting unbroken (or not known-broken) Enigma ciphertexts from numerous sources;
2) an Enigma machine, hill-climbing solver and other tools, created via prompting Claude Opus 5.5;
3) an account at The National Archives to retrieve U-boat decrypts on relevant dates;
4) OCR on the resulting PDFs and identification of a translation of this decrypt by matching its time of origin, 1414 on 6 March;
5) Cribs: the content from Bletchley’s translation (ZTPGU/36892 from DEFE 3/741, part 6, pp. 69–70); the phrasing from a German original of the same kind of signal in U–806’s war diary on uboatarchive.net (Bletchley’s translation: ZTPGU/36575);
6) a program that simulated an M4 bombe efficiently and tried every middle ring on 80 threads;
7) an M4 simulator which shares no code with the bombe, that decrypts and re-encrypts the message.
This is a most welcome break and, to my knowledge, the first AI-assisted break of an Enigma M4 message. This break also succeeded by finding a good crib, and it shows that archive research pays off when breaking single Enigma messages like this. Often the ciphertext-only attack using hill-climbing or simulated annealing will fail on the short singletons.
Updated: 26 October 2026 at 11:37 UTC

## Original Extract

An AI-assisted break of the Enigma M4 Triton message of 6 March 1945
In 2006, Ralph Erskine published three German Naval Enigma messages, M4 messages in the keys Shark, Turtle and Pike , enciphered on the Naval Enigma M4, as another contribution to the then ongoing M4 Message Breaking Project . They were then busy trying to break the four M4 messages intercepted by H.M.S. Hurricane in the Atlantic in November 1942 and published by Ralph Erskine in a Cryptologia article in October 1996. The first two of the 1942 messages were broken by the members of the M4 Message Breaking Project in February and March 2006, while the third message was broken by Dan Girard in January 2013.
If anyone attempted to break the three M4 messages in the keys of Shark, Turtle and Pike, it is not known to me. The file with the three messages from October 1943, January 1944, and March 1945 has been sitting in Crypto Cellar Research since March 2006, and I must admit I had forgotten about them.
I was therefore greatly surprised when Jason Parker-Burlingham contacted me on 25 October 2026 to announce that he had broken the first of the three messages, the Shark (Triton) signal of 6 March 1945.
Here is his report:
I have decrypted the 6 March 1945 Shark signal (“CXSO UNVZ NEWM ESUC IIFP RLEE TPNZ LPYJ WDKL LDUJ”), first on the page at https://cryptocellar.org/bgac/m4-messages.html .
The successful crib was NULVIERDREINULBISNULSECHSDREINUL (“0430 to 0630”), 123rd in a ranked list; a second crib, LAUFENNACHPERIODEJBRUNOJ, gave the same key independently.
Plaintext, with no corrections applied or encoding undone:
EULA UFEN NACH PERI ODEJ BRUN OJXX EINS KKFU NKFE
UERG EUPP EEIN SEIN SAMS IEBE NXMA ERZV ONEI NSBI
SDRE IGEI NSZW OBIS EINS VIER YACH TXNU LVIE RDRE
INUL BISN ULSE CHSD REIN ULYN EUNX EINS NULB ISEI
NSDW OYZW OZWO BISZ WOVG ERXZ WOKK FUNK FEUE RGRU
PPEE CNSZ WVAM SIEB ENRM JERZ VONE INSB SSDR EIYA
CHTX NULV IERD REIN ULBI SNUL SECH SBRE INUL YEIN
SVIE RIRE INUL BISE IOSS ECHS DREI NZLY NEUN XKKF
ORTK KEIN SNUL BISE IESZ WOBV
Plaintext, with words separated, punctuation applied and 14 letters corrected:
ES LAUFEN NACH PERIODE "BRUNO": EINS) FUNKFEUERGRUPPE EINSEINS AM
SIEBEN. MAERZ VON EINS BIS DREI, EINSZWO BIS EINSVIER; ACHT.
NULVIERDREINUL BIS NULSECHSDREINUL; NEUN. EINSNUL BIS EINSZWO,
ZWOZWO BIS ZWOVIER. ZWO) FUNKFEUERGRUPPE EINSZWO AM SIEBEN.
MAERZ VON EINS BIS DREI; ACHT. NULVIERDREINUL BIS NULSECHSDREINUL,
EINSVIERDREINUL BIS EINSSECHSDREINUL; NEUN. (FORT) EINSNUL BIS
EINSZWO
In English:
"The following will operate according to period 'Bruno':
1) Radio-beacon group 11: on 7 March from 0100 to 0300 and 1200 to 1400;
on the 8th, 0430 to 0630; on the 9th, 1000 to 1200 and 2200 to 2400.
2) Radio-beacon group 12: on 7 March from 0100 to 0300;
on the 8th, 0430 to 0630 and 1430 to 1630;
on the 9th, (continued) 1000 to 1200 …"
Machine settings:
M4, thin reflector B, Greek wheel gamma, rotors IV V III,
rings A A P U, start Z K S A, plugs AX BZ CY DG ET IP JN OQ RV SW.
The parts that made up the solution were as follows, given in more or less the order in which I prompted the model to build and run them:
1) collecting unbroken (or not known-broken) Enigma ciphertexts from numerous sources;
2) an Enigma machine, hill-climbing solver and other tools, created via prompting Claude Opus 5.5;
3) an account at The National Archives to retrieve U-boat decrypts on relevant dates;
4) OCR on the resulting PDFs and identification of a translation of this decrypt by matching its time of origin, 1414 on 6 March;
5) Cribs: the content from Bletchley’s translation (ZTPGU/36892 from DEFE 3/741, part 6, pp. 69–70); the phrasing from a German original of the same kind of signal in U–806’s war diary on uboatarchive.net (Bletchley’s translation: ZTPGU/36575);
6) a program that simulated an M4 bombe efficiently and tried every middle ring on 80 threads;
7) an M4 simulator which shares no code with the bombe, that decrypts and re-encrypts the message.
This is a most welcome break and, to my knowledge, the first AI-assisted break of an Enigma M4 message. This break also succeeded by finding a good crib, and it shows that archive research pays off when breaking single Enigma messages like this. Often the ciphertext-only attack using hill-climbing or simulated annealing will fail on the short singletons.
Updated: 26 October 2026 at 11:37 UTC
