---
source: "https://shattered.io/rsa-vs-ml-dsa-signatures-2026/"
hn_url: "https://news.ycombinator.com/item?id=50053758"
title: "RSA vs. ML-DSA: 256 vs. 2,420-Byte Signatures (2026)"
article_title: "RSA vs ML-DSA: 256 vs 2,420-Byte Signatures [2026]"
image: "https://shattered.io/wp-content/uploads/2026/10/rsa-vs-ml-dsa-signatures-2026-1.webp"
author: "msgchainhq"
captured_at: "2026-10-12T10:19:42Z"
capture_tool: "hn-digest"
hn_id: 50053758
score: 1
comments: 0
posted_at: "2026-10-12T09:37:39Z"
tags:
  - hacker-news
---

# RSA vs. ML-DSA: 256 vs. 2,420-Byte Signatures (2026)

- HN: [50053758](https://news.ycombinator.com/item?id=50053758)
- Source: [shattered.io](https://shattered.io/rsa-vs-ml-dsa-signatures-2026/)
- Score: 1
- Comments: 0
- Posted: 2026-10-12T09:37:39Z

## Translation

Title: RSA vs. ML-DSA: 256 vs. 2,420-Byte Signatures (2026)
Article title: RSA vs ML-DSA: 256 vs 2,420-Byte Signatures [2026]
Description: RSA vs ML-DSA compared: signature sizes, benchmarks, CNSA 2.0 deadlines, and real adoption data from Cloudflare, AWS, and Google in 2026.

Article text:
RSA vs ML-DSA: 256 vs 2,420-Byte Signatures [2026]
SHA TERED
☰
Home
RSA vs ML-DSA: 256 vs 2,420-Byte Signatures [2026]
RSA has signed the internet’s certificates for three decades. ML-DSA, finalized by NIST in August 2024, is the algorithm built to take over once quantum computers make RSA’s math solvable in practice. The catch is that ML-DSA signatures run 5 to 18 times larger than RSA’s, and that single fact is reshaping how banks, cloud providers, and certificate authorities plan their next five years of infrastructure spending.
Why RSA’s Clock Is Running Out
What Is ML-DSA, and Why Did NIST Pick It?
RSA vs ML-DSA: The Core Specs Table
Signature and Key Sizes: The Byte-for-Byte Gap
Benchmark Results: Signing and Verification Speed
What Bigger Signatures Actually Cost: Bandwidth, TLS, and Storage
NIST Security Levels, Explained in Plain Terms
CNSA 2.0 and Government Migration Deadlines
Real-World Adoption: Who’s Already Shipping ML-DSA
Hybrid Signatures: Running Both Algorithms at Once
Generating Keys: What Changes in Your Code
Migration Guide: Eight Steps From RSA to ML-DSA
Five Use Cases: Picking the Right Algorithm for the Job
Where SLH-DSA and Hybrid Classical Schemes Still Fit
The Verdict: Which One Should You Actually Deploy?
Why is ML-DSA’s signature so much bigger than RSA’s?
Can I use ML-DSA with my current certificate authority?
What’s the difference between ML-DSA-44, ML-DSA-65, and ML-DSA-87?
Does switching to ML-DSA cost more money?
Is ML-DSA the same thing as Dilithium?
Should small businesses worry about this migration yet?
What should I do if my hardware security module doesn’t support ML-DSA?
This comparison walks through the actual numbers: signature and key sizes straight from the FIPS 204 specification, benchmark ranges from liboqs and Cloudflare’s own testing, the CNSA 2.0 deadlines federal contractors are already tracking, and which companies have shipped ML-DSA support into production in 2025 and 2026. By the end you’ll know exactly which algorithm to deploy where, and when.
None of this is hypothetical anymore. OpenSSL shipped native ML-DSA support over a year ago, the major cloud key management services already let you choose it, and a hard regulatory deadline for U.S. national security systems arrives on January 1, 2027. The question for most engineering teams in October 2026 isn’t whether ML-DSA matters, it’s which systems need it first and which can safely wait.
Why RSA’s Clock Is Running Out
RSA signatures rely on the difficulty of factoring the product of two large prime numbers. That problem has held up against classical computers for over 40 years. It does not hold up against Shor’s algorithm running on a sufficiently large, fault-tolerant quantum computer, which can factor an RSA modulus in polynomial time rather than the exponential time classical computers need.
No quantum computer today can factor RSA-2048 at a cryptographically relevant scale. Published estimates for what it would take vary widely, ranging from hundreds to thousands of logical qubits, which translates to potentially millions of physical qubits once error correction overhead is factored in. There is no single agreed “RSA break date,” and serious estimates keep shifting as quantum hardware architecture evolves.
The more pressing issue isn’t a future break, it’s a present one: harvest now, decrypt later. An attacker can record encrypted traffic today and sit on it until a capable quantum computer exists, then decrypt everything retroactively. Signatures carry a related but distinct risk. Once quantum attacks become practical, an adversary could forge RSA signatures on certificates, code-signing packages, and authentication tokens that were never meant to be verified decades into the future. That’s precisely the gap ML-DSA was built to close.
What Is ML-DSA, and Why Did NIST Pick It?
ML-DSA stands for Module-Lattice-Based Digital Signature Algorithm. NIST finalized it as FIPS 204 on August 13, 2024, alongside ML-KEM (FIPS 203, for key exchange) and SLH-DSA (FIPS 205, a hash-based backup signature scheme). Before standardization, ML-DSA competed and won under its original research name, CRYSTALS-Dilithium.
Instead of factoring large numbers, ML-DSA’s security rests on the hardness of certain lattice problems, specifically finding short vectors in high-dimensional structured lattices. That math is believed to resist both classical and quantum attacks, including Shor’s algorithm, which has no known way to solve lattice problems efficiently.
FIPS 204 defines three parameter sets: ML-DSA-44, ML-DSA-65, and ML-DSA-87. The numbers refer to internal lattice dimension parameters, not key or signature sizes, so don’t read them as byte counts. Each set maps to one of NIST’s five security categories, and picking the right one is the first real decision any team migrating off RSA has to make.
RSA vs ML-DSA: The Core Specs Table
Here’s the side-by-side breakdown, pulling exact figures from the FIPS 204 specification and NIST’s published security category mapping.
RSA key sizes above are approximate because the byte count depends heavily on encoding: raw modulus and exponent, PKCS#1 DER, PKCS#8 DER, and PEM (which adds roughly a third more size through Base64 wrapping) all produce different numbers for the “same” key. ML-DSA’s figures, by contrast, are fixed exactly in FIPS 204 Table 2, which is part of why lattice-based schemes are easier to benchmark consistently.
Signature and Key Sizes: The Byte-for-Byte Gap
The headline number in any RSA vs ML-DSA conversation is the signature size gap. An RSA-2048 signature fits in 256 bytes. Swap it for ML-DSA-44 and that same signature balloons to 2,420 bytes, a 9.5x increase. Move up to ML-DSA-87, the parameter set CNSA 2.0 targets for national security systems, and a single signature hits 4,627 bytes, roughly 18 times the size of the RSA-2048 signature it replaces.
Public keys follow the same pattern, though less dramatically. ML-DSA-44’s 1,312-byte public key already exceeds the largest commonly encoded RSA-4096 public key. Private keys show the widest spread: ML-DSA-87’s 4,896-byte private key runs comparable to a PKCS#8-encoded RSA-4096 private key, so the private-key overhead is less painful than the signature overhead.
That asymmetry matters because signatures get transmitted over the wire on every single TLS handshake, every code-signing verification, and every document check. Private keys mostly sit on disk or inside an HSM and never leave. The practical cost of migrating to ML-DSA shows up almost entirely in network and storage overhead tied to signatures and public keys, not in private-key storage.
Benchmark Results: Signing and Verification Speed
Raw speed tells a more forgiving story than size does. Benchmarks from the Open Quantum Safe project’s liboqs library, cross-referenced against Cloudflare’s own post-quantum signature testing and independent measurements published by systemshardening.com, converge on a consistent pattern even though exact throughput numbers swing with CPU model, compiler flags, and whether AVX2 or AVX-512 vectorization is enabled.
The pattern worth remembering: RSA’s verification step is its biggest advantage, often an order of magnitude faster than its own signing step, which is exactly why RSA has dominated TLS server certificates for so long (clients verify far more often than servers sign). ML-DSA’s signing and verification costs sit much closer together, which changes the performance calculus for services that sign constantly, like code-signing pipelines or high-frequency API authentication, where ML-DSA’s more balanced profile can actually come out ahead despite the larger output.
Treat every number in that table as a range, not a constant. liboqs results should always be reported with the library version, compiler, and CPU model attached, because a benchmark showing 20,000 operations per second on one machine and 2,000 on another isn’t a contradiction, it’s two different hardware and implementation configurations.
What Bigger Signatures Actually Cost: Bandwidth, TLS, and Storage
Signature size differences stop being academic the moment you multiply them across millions of daily connections. A TLS 1.3 CertificateVerify message carrying an ML-DSA-87 signature runs roughly 4.6 KB for the signature payload alone, before certificate data and handshake framing are even counted. A full certificate chain signed end-to-end with ML-DSA-65 can add 10 KB or more compared to an equivalent RSA or ECDSA chain.
None of the major cloud key management services charge extra per signature algorithm today. AWS KMS and Google Cloud KMS both bill ML-DSA signing the same way they bill any other asymmetric signing operation, through their standard per-API-call pricing. The real cost of this migration isn’t a line item on a cloud bill, it’s engineering time spent on library upgrades, infrastructure that chokes on oversized ASN.1 fields, and certificate-chain bloat that quietly eats into bandwidth budgets at scale.
NIST Security Levels, Explained in Plain Terms
NIST defines five security categories for post-quantum algorithms, each pegged to the computational effort required to break a reference classical algorithm. ML-DSA-44 sits at Category 2, roughly equivalent to the effort needed to find a SHA-256 collision. ML-DSA-65 sits at Category 3, comparable to breaking AES-192 by brute force. ML-DSA-87 sits at Category 5, comparable to breaking AES-256.
In practice, most organizations land on ML-DSA-65 as the default choice, the same way RSA-2048 became the default rather than RSA-4096 for years: it balances security margin against size and performance cost. ML-DSA-87 gets reserved for cases where regulation demands it, principally CNSA 2.0 compliance for U.S. national security systems, or where an organization wants maximum long-term margin for data or signatures that need to stay trustworthy for decades.
CNSA 2.0 and Government Migration Deadlines
The NSA’s Commercial National Security Algorithm Suite 2.0 is the single clearest deadline driving ML-DSA adoption in 2026, even though it technically applies only to National Security Systems rather than every commercial entity in the country.
January 1, 2027: all new National Security System acquisitions must use CNSA 2.0-compliant algorithms, which for signatures means ML-DSA-87.
Through 2030: organizations are expected to have active migration planning and procurement underway, phasing out classical algorithms in new deployments.
By 2033: CNSA 2.0 calls for exclusive use of the new suite across applicable systems, subject to NSA implementation guidance.
CNSA 2.0 also requires post-quantum key establishment via ML-KEM, not just post-quantum signatures, so contractors touching both encryption and signing infrastructure are dealing with a two-front migration. The deadlines don’t bind ordinary commercial companies or civilian agencies by statute, but defense contractors, their suppliers, and anyone selling into that procurement chain are already building ML-DSA-87 support into 2026 and 2027 product roadmaps regardless of whether a specific contract requires it yet.
Running in parallel, the CA/Browser Forum has been shortening public TLS certificate validity on its own separate timeline: maximum validity dropped to 200 days on March 15, 2026, is scheduled to fall to 100 days on March 15, 2027, and down to 47 days by March 15, 2029. That schedule isn’t an ML-DSA mandate, but shorter-lived certificates push every CA toward heavier automation, which happens to make a future switch to post-quantum signatures operationally easier once public PKI support arrives.
Real-World Adoption: Who’s Already Shipping ML-DSA
ML-DSA moved from paper standard to shipping product faster than most post-quantum algorithms, largely because OpenSSL 3.5.0, released April 8, 2025 as a long-term-support build, brought native ML-KEM, ML-DSA, and SLH-DSA support into the default provider with no external patch req

[truncated]

## Original Extract

RSA vs ML-DSA compared: signature sizes, benchmarks, CNSA 2.0 deadlines, and real adoption data from Cloudflare, AWS, and Google in 2026.

RSA vs ML-DSA: 256 vs 2,420-Byte Signatures [2026]
SHA TERED
☰
Home
RSA vs ML-DSA: 256 vs 2,420-Byte Signatures [2026]
RSA has signed the internet’s certificates for three decades. ML-DSA, finalized by NIST in August 2024, is the algorithm built to take over once quantum computers make RSA’s math solvable in practice. The catch is that ML-DSA signatures run 5 to 18 times larger than RSA’s, and that single fact is reshaping how banks, cloud providers, and certificate authorities plan their next five years of infrastructure spending.
Why RSA’s Clock Is Running Out
What Is ML-DSA, and Why Did NIST Pick It?
RSA vs ML-DSA: The Core Specs Table
Signature and Key Sizes: The Byte-for-Byte Gap
Benchmark Results: Signing and Verification Speed
What Bigger Signatures Actually Cost: Bandwidth, TLS, and Storage
NIST Security Levels, Explained in Plain Terms
CNSA 2.0 and Government Migration Deadlines
Real-World Adoption: Who’s Already Shipping ML-DSA
Hybrid Signatures: Running Both Algorithms at Once
Generating Keys: What Changes in Your Code
Migration Guide: Eight Steps From RSA to ML-DSA
Five Use Cases: Picking the Right Algorithm for the Job
Where SLH-DSA and Hybrid Classical Schemes Still Fit
The Verdict: Which One Should You Actually Deploy?
Why is ML-DSA’s signature so much bigger than RSA’s?
Can I use ML-DSA with my current certificate authority?
What’s the difference between ML-DSA-44, ML-DSA-65, and ML-DSA-87?
Does switching to ML-DSA cost more money?
Is ML-DSA the same thing as Dilithium?
Should small businesses worry about this migration yet?
What should I do if my hardware security module doesn’t support ML-DSA?
This comparison walks through the actual numbers: signature and key sizes straight from the FIPS 204 specification, benchmark ranges from liboqs and Cloudflare’s own testing, the CNSA 2.0 deadlines federal contractors are already tracking, and which companies have shipped ML-DSA support into production in 2025 and 2026. By the end you’ll know exactly which algorithm to deploy where, and when.
None of this is hypothetical anymore. OpenSSL shipped native ML-DSA support over a year ago, the major cloud key management services already let you choose it, and a hard regulatory deadline for U.S. national security systems arrives on January 1, 2027. The question for most engineering teams in October 2026 isn’t whether ML-DSA matters, it’s which systems need it first and which can safely wait.
Why RSA’s Clock Is Running Out
RSA signatures rely on the difficulty of factoring the product of two large prime numbers. That problem has held up against classical computers for over 40 years. It does not hold up against Shor’s algorithm running on a sufficiently large, fault-tolerant quantum computer, which can factor an RSA modulus in polynomial time rather than the exponential time classical computers need.
No quantum computer today can factor RSA-2048 at a cryptographically relevant scale. Published estimates for what it would take vary widely, ranging from hundreds to thousands of logical qubits, which translates to potentially millions of physical qubits once error correction overhead is factored in. There is no single agreed “RSA break date,” and serious estimates keep shifting as quantum hardware architecture evolves.
The more pressing issue isn’t a future break, it’s a present one: harvest now, decrypt later. An attacker can record encrypted traffic today and sit on it until a capable quantum computer exists, then decrypt everything retroactively. Signatures carry a related but distinct risk. Once quantum attacks become practical, an adversary could forge RSA signatures on certificates, code-signing packages, and authentication tokens that were never meant to be verified decades into the future. That’s precisely the gap ML-DSA was built to close.
What Is ML-DSA, and Why Did NIST Pick It?
ML-DSA stands for Module-Lattice-Based Digital Signature Algorithm. NIST finalized it as FIPS 204 on August 13, 2024, alongside ML-KEM (FIPS 203, for key exchange) and SLH-DSA (FIPS 205, a hash-based backup signature scheme). Before standardization, ML-DSA competed and won under its original research name, CRYSTALS-Dilithium.
Instead of factoring large numbers, ML-DSA’s security rests on the hardness of certain lattice problems, specifically finding short vectors in high-dimensional structured lattices. That math is believed to resist both classical and quantum attacks, including Shor’s algorithm, which has no known way to solve lattice problems efficiently.
FIPS 204 defines three parameter sets: ML-DSA-44, ML-DSA-65, and ML-DSA-87. The numbers refer to internal lattice dimension parameters, not key or signature sizes, so don’t read them as byte counts. Each set maps to one of NIST’s five security categories, and picking the right one is the first real decision any team migrating off RSA has to make.
RSA vs ML-DSA: The Core Specs Table
Here’s the side-by-side breakdown, pulling exact figures from the FIPS 204 specification and NIST’s published security category mapping.
RSA key sizes above are approximate because the byte count depends heavily on encoding: raw modulus and exponent, PKCS#1 DER, PKCS#8 DER, and PEM (which adds roughly a third more size through Base64 wrapping) all produce different numbers for the “same” key. ML-DSA’s figures, by contrast, are fixed exactly in FIPS 204 Table 2, which is part of why lattice-based schemes are easier to benchmark consistently.
Signature and Key Sizes: The Byte-for-Byte Gap
The headline number in any RSA vs ML-DSA conversation is the signature size gap. An RSA-2048 signature fits in 256 bytes. Swap it for ML-DSA-44 and that same signature balloons to 2,420 bytes, a 9.5x increase. Move up to ML-DSA-87, the parameter set CNSA 2.0 targets for national security systems, and a single signature hits 4,627 bytes, roughly 18 times the size of the RSA-2048 signature it replaces.
Public keys follow the same pattern, though less dramatically. ML-DSA-44’s 1,312-byte public key already exceeds the largest commonly encoded RSA-4096 public key. Private keys show the widest spread: ML-DSA-87’s 4,896-byte private key runs comparable to a PKCS#8-encoded RSA-4096 private key, so the private-key overhead is less painful than the signature overhead.
That asymmetry matters because signatures get transmitted over the wire on every single TLS handshake, every code-signing verification, and every document check. Private keys mostly sit on disk or inside an HSM and never leave. The practical cost of migrating to ML-DSA shows up almost entirely in network and storage overhead tied to signatures and public keys, not in private-key storage.
Benchmark Results: Signing and Verification Speed
Raw speed tells a more forgiving story than size does. Benchmarks from the Open Quantum Safe project’s liboqs library, cross-referenced against Cloudflare’s own post-quantum signature testing and independent measurements published by systemshardening.com, converge on a consistent pattern even though exact throughput numbers swing with CPU model, compiler flags, and whether AVX2 or AVX-512 vectorization is enabled.
The pattern worth remembering: RSA’s verification step is its biggest advantage, often an order of magnitude faster than its own signing step, which is exactly why RSA has dominated TLS server certificates for so long (clients verify far more often than servers sign). ML-DSA’s signing and verification costs sit much closer together, which changes the performance calculus for services that sign constantly, like code-signing pipelines or high-frequency API authentication, where ML-DSA’s more balanced profile can actually come out ahead despite the larger output.
Treat every number in that table as a range, not a constant. liboqs results should always be reported with the library version, compiler, and CPU model attached, because a benchmark showing 20,000 operations per second on one machine and 2,000 on another isn’t a contradiction, it’s two different hardware and implementation configurations.
What Bigger Signatures Actually Cost: Bandwidth, TLS, and Storage
Signature size differences stop being academic the moment you multiply them across millions of daily connections. A TLS 1.3 CertificateVerify message carrying an ML-DSA-87 signature runs roughly 4.6 KB for the signature payload alone, before certificate data and handshake framing are even counted. A full certificate chain signed end-to-end with ML-DSA-65 can add 10 KB or more compared to an equivalent RSA or ECDSA chain.
None of the major cloud key management services charge extra per signature algorithm today. AWS KMS and Google Cloud KMS both bill ML-DSA signing the same way they bill any other asymmetric signing operation, through their standard per-API-call pricing. The real cost of this migration isn’t a line item on a cloud bill, it’s engineering time spent on library upgrades, infrastructure that chokes on oversized ASN.1 fields, and certificate-chain bloat that quietly eats into bandwidth budgets at scale.
NIST Security Levels, Explained in Plain Terms
NIST defines five security categories for post-quantum algorithms, each pegged to the computational effort required to break a reference classical algorithm. ML-DSA-44 sits at Category 2, roughly equivalent to the effort needed to find a SHA-256 collision. ML-DSA-65 sits at Category 3, comparable to breaking AES-192 by brute force. ML-DSA-87 sits at Category 5, comparable to breaking AES-256.
In practice, most organizations land on ML-DSA-65 as the default choice, the same way RSA-2048 became the default rather than RSA-4096 for years: it balances security margin against size and performance cost. ML-DSA-87 gets reserved for cases where regulation demands it, principally CNSA 2.0 compliance for U.S. national security systems, or where an organization wants maximum long-term margin for data or signatures that need to stay trustworthy for decades.
CNSA 2.0 and Government Migration Deadlines
The NSA’s Commercial National Security Algorithm Suite 2.0 is the single clearest deadline driving ML-DSA adoption in 2026, even though it technically applies only to National Security Systems rather than every commercial entity in the country.
January 1, 2027: all new National Security System acquisitions must use CNSA 2.0-compliant algorithms, which for signatures means ML-DSA-87.
Through 2030: organizations are expected to have active migration planning and procurement underway, phasing out classical algorithms in new deployments.
By 2033: CNSA 2.0 calls for exclusive use of the new suite across applicable systems, subject to NSA implementation guidance.
CNSA 2.0 also requires post-quantum key establishment via ML-KEM, not just post-quantum signatures, so contractors touching both encryption and signing infrastructure are dealing with a two-front migration. The deadlines don’t bind ordinary commercial companies or civilian agencies by statute, but defense contractors, their suppliers, and anyone selling into that procurement chain are already building ML-DSA-87 support into 2026 and 2027 product roadmaps regardless of whether a specific contract requires it yet.
Running in parallel, the CA/Browser Forum has been shortening public TLS certificate validity on its own separate timeline: maximum validity dropped to 200 days on March 15, 2026, is scheduled to fall to 100 days on March 15, 2027, and down to 47 days by March 15, 2029. That schedule isn’t an ML-DSA mandate, but shorter-lived certificates push every CA toward heavier automation, which happens to make a future switch to post-quantum signatures operationally easier once public PKI support arrives.
Real-World Adoption: Who’s Already Shipping ML-DSA
ML-DSA moved from paper standard to shipping product faster than most post-quantum algorithms, largely because OpenSSL 3.5.0, released April 8, 2025 as a long-term-support build, brought native ML-KEM, ML-DSA, and SLH-DSA support into the default provider with no external patch req

[truncated]
