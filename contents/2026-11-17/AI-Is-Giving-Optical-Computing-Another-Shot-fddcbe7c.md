---
source: "https://spectrum.ieee.org/optical-computing-ai"
hn_url: "https://news.ycombinator.com/item?id=50533886"
title: "AI Is Giving Optical Computing Another Shot"
article_title: "AI Is Giving Optical Computing Another Shot - IEEE Spectrum"
image: "https://spectrum.ieee.org/media-library/close-up-of-a-green-circuit-board-with-a-central-microchip-and-intricate-traces.jpg?id=68692424&width=1200&height=600&coordinates=0%2C531%2C0%2C531"
author: "speckx"
captured_at: "2026-11-17T16:43:35Z"
capture_tool: "hn-digest"
hn_id: 50533886
score: 1
comments: 0
posted_at: "2026-11-17T16:32:10Z"
tags:
  - hacker-news
---

# AI Is Giving Optical Computing Another Shot

- HN: [50533886](https://news.ycombinator.com/item?id=50533886)
- Source: [spectrum.ieee.org](https://spectrum.ieee.org/optical-computing-ai)
- Score: 1
- Comments: 0
- Posted: 2026-11-17T16:32:10Z

## Translation

Title: AI Is Giving Optical Computing Another Shot
Article title: AI Is Giving Optical Computing Another Shot - IEEE Spectrum
Description: The push to cut AI inference’s energy use could pay off in unexpected ways

Article text:
-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
AI Is Giving Optical Computing Another Shot - IEEE Spectrum
IEEE.org IEEE Xplore IEEE Standards IEEE Job Site More Sites Sign In Join IEEE AI Is Giving Optical Computing Another Shot Share FOR THE TECHNOLOGY INSIDER Search: Explore by topic Aerospace AI Biomedical Climate Tech Computing Consumer Electronics Energy History of Technology Robotics Semiconductors Telecommunications Transportation
IEEE Spectrum
FOR THE TECHNOLOGY INSIDER Topics
Enjoy more free content and benefits by creating an account
Saving articles to read later requires an IEEE Spectrum account
The Institute content is only available for members
Downloading full PDF issues is exclusive for IEEE Members
Downloading this e-book is exclusive for IEEE Members
Access to
Spectrum
's Digital Edition is exclusive for IEEE Members
Following topics is a feature exclusive for IEEE Members
Adding your response to an article requires an IEEE Spectrum account
Create an account to access more content and features on
IEEE Spectrum
, including the ability to save articles to read later, download Spectrum Collections, and participate in
conversations with readers and editors. For more exclusive content and features, consider
Joining IEEE
.
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
all of Spectrum’s articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
this e-book plus all of
IEEE Spectrum’s
articles, archives, PDF downloads, and other benefits.
Learn more about IEEE → CREATE AN ACCOUNT SIGN IN JOIN IEEE SIGN IN Close Access Thousands of Articles — Completely Free
Create an account and get exclusive content and features:
Save articles, download collections,
and
post comments
— all free! For full access and benefits,
subscribe
to
Spectrum
.
AI Is Giving Optical Computing Another Shot
The push to cut AI inference’s energy use could pay off in unexpected ways
Data centers already use light in networking. Startup Optalysis aims to use specially engineered waveguides to add computational capacity to the networking path, doing math while data is in transit.
The idea of computing with light has been around for over 70 years, but until recently it’s remained largely in the realm of research. Over the past few years, though, optical computing has emerged as a fast-growing industry: Market valuation of companies in this sector now totals US $2.22 billion, up from $0.4 billion in 2022. One projection suggests that the market will continue to grow exponentially in the coming years.
The reason for this sudden surge will surprise no one: AI. In particular, AI inference—the process of using AI models to respond to user queries—has driven much of the construction of huge, power-hungry data centers , and relatively energy-efficient optical systems are seen as a potential way forward.
“Most of the interest in [ photonics ] is driven by the fact that AI inference has literally become limited by power, and in a rather striking manner,” says Peter McMahon , associate professor of applied and engineering physics at Cornell University , in Ithaca, N.Y., who has researched optical computing for over a decade.
Optics , at least on paper, carries the promise of doing the same math as today’s GPUs at a drastically lower level of power. Switching from electronics to optics would be a pretty radical solution. Nevertheless, like mushrooms after a rainstorm, optical-computing startups are popping up in the United States, China, Germany, and Israel. And while competing against such GPU makers as Nvidia is an ambitious goal, the investment in optical computing has spurred development of complementary approaches, such as designing alternative AI algorithms , taking on scientific computing , and enhancing in-network computing.
Optalysys, based in England, is inserting optical-computing elements into the networking path, where data centers already use light. The company’s basic on-chip optical element, shown in the inset, is a Fourier transform calculator. Optalysys
“I think it would be good for the technology if many of these companies made a good run at it,” says Logan Wright , assistant professor of applied physics at Yale University, in New Haven, Conn., and a former postdoc in McMahon’s lab. “I think that the technology that will be developed through this would be really, really useful, and enable photonic systems that maybe don’t get on this wave of AI but would be capable of doing really quite special things.”
Inference makes up the bulk of AI energy use. The industry is already spending billions of dollars designing advanced GPUs and other chips to run inference as efficiently as possible. At the heart of the inference computation is one operation, done countless times a second: matrix-vector multiplication.
Back in 1978, optical-computing pioneer Joseph W. Goodman and his coworkers proposed a method for doing matrix-vector multiplication with light. The setup was relatively simple. Let’s say a vector of length N needs to be multiplied by a matrix of size N by N . N light-emitting diodes , all in a row, encode the elements of the vector into the intensity of light—the brighter the light from an individual diode , the greater the value that diode encodes. Then, a cylindrical lens spreads out the light vertically, essentially copying the vector N times. The light then passes through an attenuator of N by N pixels. At each pixel, the light is attenuated according to the corresponding value at that location in the matrix. Passing the light through the attenuator performs the multiplication. Finally, a cylindrical lens oriented perpendicular to the first lens re-collects the light from each row, summing it up. The final vector is measured by a row of detectors.
The most computationally expensive part of common AI today is vector-matrix (or matrix-matrix) multiplication. For a vector of length N (here, N = 3), it takes N 2 (here, 9) multiplication steps and three accumulation steps—multiple additions at once.
The argument for why optics might be advantageous over electronics goes something like this: Multiplying a matrix of size N by N by a vector of length N takes N 2 multiplication steps. In electronics, each multiplication takes some energy, so the total energy used scales as N 2 . In optics, however, the multiplication step is passive: As long as the matrix stays the same, passing light through the attenuator costs almost no energy. The only cost is encoding the vector into light and reading out the answer, but the vector is of length N , not N 2 . So the energy used will scale as N , not N 2 , as bigger vectors and matrices are multiplied.
Since Goodman’s 1978 proposal, the quality of the available optical elements has improved dramatically. Indeed, in 2022, a team of researchers at Cornell led by McMahon demonstrated an optical neural network that consumed less than one photon of energy—approximately 2.5 x 10 -19 joules—per multiplication operation. In a standard computer, a single multiply operation can take between 10 -14 and 10 -12 joules, which suggests that optics could possibly reduce energy use by a factor of 100,000.
But McMahon cautions against such extrapolation. First, the cost of encoding data into optics was not included in the calculation. Even if the scaling is favorable with light, in practice it’s not just the scaling that matters, it’s also the coefficients. Let’s say the energy cost of doing one multiplication with electronics is 1 picojoule, and the cost to encode one vector element with optics is 1,000 picojoules. The electronic approach still wins if the length N of the vector is fewer than 1,000 elements.
McMahon says this example is not far off from reality: “One of the key questions is, How big do you have to make the system before you start winning? Lots of people have done this calculation for various kinds of neural networks . The very rough answer is somewhere between N equals 1,000 and N equals 10,000.”
It’s also important to keep in mind a computing principle known as Amdahl’s law .
“Amdahl’s law shouldn’t even be called a law. It’s a simple piece of common sense,” McMahon says. “If you have something that’s going to accelerate a computation, and it can accelerate 90 percent of the code, but there’s 10 percent of the code that it can’t accelerate, then the maximum possible speedup you can get is 10x.”
With AI inferencing, the point is to save energy, not time. But Amdahl’s observation still applies. What fraction of today’s AI inferencing code can be better handled with optics is a complicated question that varies by approach. But it’s not 100 percent. Matrix-vector multiplication, or matrix-matrix multiplication, is only part of the inference workload—there are other, nonlinear operations that need to be performed on the data.
Taking all of these caveats into account, optical-computing companies today are aiming for a 10- to 80-fold improvement in energy efficiency over today’s GPUs. The question is whether they can deliver on that promise at scale before traditional electronics catches up to those gains.
Bridging the gap in free space
The startup Lumai , based in Oxford, England, has taken perhaps the most direct inspiration from Goodman and his team. Lumai’s head of product, Phillip Burr , declares that a free-space approach, where the light travels through air rather than on-chip waveguides, is what will enable the company’s tech to beat GPUs.
Instead of the fixed attenuator used in Goodman’s proposal, Lumai is using a modern programmable device called a spatial light modulator—a pixelated membrane akin to a fancy liquid-crystal display that can modulate both the intensity and phase of light. The rest of the design is more sophisticated in practice but similar in principle to Goodman’s design: A vector is encoded into light, copied N times via a lens, passed through the spatial light modulator, collected with another lens, and the result is converted back to digital electronics. The electronics then handle the nonlinear functions. “A full vector-matrix multiply completes in a single optical cycle,” Burr says.
To accurately compare the energy expenditures of the optical and GPU approaches, you must first specify which part of the workload you’re considering. Modern inference is often split into two phases: prefill , where the input query is processed and propagated through the model, and decode , where the output is generated one word (or token) at a time. Prefill is the more computationally intensive of the two. It basically involves doing vector-matrix multiplications for the same matrix over and over. Decode also involves vector-matrix multiplies, but not as intensively—the time-consuming part of decode is reading the necessary matrices in and out of memory.
Vector-matrix multiplication with free-space optics
The Goodman optical computer is the simplest version of a photonic vector-matrix multiplier. Here, the vector elements are encoded into the intensity of light from three light sources, such as LEDs . Then, a cylindrical lens spreads the light vertically. The matrix elements are encoded into the opacity of an attenuator, such that when light passes through each pixel it is multiplied (attenuated) by that matrix value. Finally, another cylindrical lens collects the light horizontally, performing the accumulation step, and a photodetector measures the output vector.
Lumai, and other companies in the space, are focusing on the compute-intensive prefill phase of inference. Lumai claims that its Iris Nova optical server, made available for evaluation in April 2026, is a proof of concept, and its next-generation Iris Aura server will do prefill using roughly one-tenth of the energy a GPU would use.
The startup Neurophos ,

[truncated]

## Original Extract

The push to cut AI inference’s energy use could pay off in unexpected ways

-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
AI Is Giving Optical Computing Another Shot - IEEE Spectrum
IEEE.org IEEE Xplore IEEE Standards IEEE Job Site More Sites Sign In Join IEEE AI Is Giving Optical Computing Another Shot Share FOR THE TECHNOLOGY INSIDER Search: Explore by topic Aerospace AI Biomedical Climate Tech Computing Consumer Electronics Energy History of Technology Robotics Semiconductors Telecommunications Transportation
IEEE Spectrum
FOR THE TECHNOLOGY INSIDER Topics
Enjoy more free content and benefits by creating an account
Saving articles to read later requires an IEEE Spectrum account
The Institute content is only available for members
Downloading full PDF issues is exclusive for IEEE Members
Downloading this e-book is exclusive for IEEE Members
Access to
Spectrum
's Digital Edition is exclusive for IEEE Members
Following topics is a feature exclusive for IEEE Members
Adding your response to an article requires an IEEE Spectrum account
Create an account to access more content and features on
IEEE Spectrum
, including the ability to save articles to read later, download Spectrum Collections, and participate in
conversations with readers and editors. For more exclusive content and features, consider
Joining IEEE
.
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
all of Spectrum’s articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
this e-book plus all of
IEEE Spectrum’s
articles, archives, PDF downloads, and other benefits.
Learn more about IEEE → CREATE AN ACCOUNT SIGN IN JOIN IEEE SIGN IN Close Access Thousands of Articles — Completely Free
Create an account and get exclusive content and features:
Save articles, download collections,
and
post comments
— all free! For full access and benefits,
subscribe
to
Spectrum
.
AI Is Giving Optical Computing Another Shot
The push to cut AI inference’s energy use could pay off in unexpected ways
Data centers already use light in networking. Startup Optalysis aims to use specially engineered waveguides to add computational capacity to the networking path, doing math while data is in transit.
The idea of computing with light has been around for over 70 years, but until recently it’s remained largely in the realm of research. Over the past few years, though, optical computing has emerged as a fast-growing industry: Market valuation of companies in this sector now totals US $2.22 billion, up from $0.4 billion in 2022. One projection suggests that the market will continue to grow exponentially in the coming years.
The reason for this sudden surge will surprise no one: AI. In particular, AI inference—the process of using AI models to respond to user queries—has driven much of the construction of huge, power-hungry data centers , and relatively energy-efficient optical systems are seen as a potential way forward.
“Most of the interest in [ photonics ] is driven by the fact that AI inference has literally become limited by power, and in a rather striking manner,” says Peter McMahon , associate professor of applied and engineering physics at Cornell University , in Ithaca, N.Y., who has researched optical computing for over a decade.
Optics , at least on paper, carries the promise of doing the same math as today’s GPUs at a drastically lower level of power. Switching from electronics to optics would be a pretty radical solution. Nevertheless, like mushrooms after a rainstorm, optical-computing startups are popping up in the United States, China, Germany, and Israel. And while competing against such GPU makers as Nvidia is an ambitious goal, the investment in optical computing has spurred development of complementary approaches, such as designing alternative AI algorithms , taking on scientific computing , and enhancing in-network computing.
Optalysys, based in England, is inserting optical-computing elements into the networking path, where data centers already use light. The company’s basic on-chip optical element, shown in the inset, is a Fourier transform calculator. Optalysys
“I think it would be good for the technology if many of these companies made a good run at it,” says Logan Wright , assistant professor of applied physics at Yale University, in New Haven, Conn., and a former postdoc in McMahon’s lab. “I think that the technology that will be developed through this would be really, really useful, and enable photonic systems that maybe don’t get on this wave of AI but would be capable of doing really quite special things.”
Inference makes up the bulk of AI energy use. The industry is already spending billions of dollars designing advanced GPUs and other chips to run inference as efficiently as possible. At the heart of the inference computation is one operation, done countless times a second: matrix-vector multiplication.
Back in 1978, optical-computing pioneer Joseph W. Goodman and his coworkers proposed a method for doing matrix-vector multiplication with light. The setup was relatively simple. Let’s say a vector of length N needs to be multiplied by a matrix of size N by N . N light-emitting diodes , all in a row, encode the elements of the vector into the intensity of light—the brighter the light from an individual diode , the greater the value that diode encodes. Then, a cylindrical lens spreads out the light vertically, essentially copying the vector N times. The light then passes through an attenuator of N by N pixels. At each pixel, the light is attenuated according to the corresponding value at that location in the matrix. Passing the light through the attenuator performs the multiplication. Finally, a cylindrical lens oriented perpendicular to the first lens re-collects the light from each row, summing it up. The final vector is measured by a row of detectors.
The most computationally expensive part of common AI today is vector-matrix (or matrix-matrix) multiplication. For a vector of length N (here, N = 3), it takes N 2 (here, 9) multiplication steps and three accumulation steps—multiple additions at once.
The argument for why optics might be advantageous over electronics goes something like this: Multiplying a matrix of size N by N by a vector of length N takes N 2 multiplication steps. In electronics, each multiplication takes some energy, so the total energy used scales as N 2 . In optics, however, the multiplication step is passive: As long as the matrix stays the same, passing light through the attenuator costs almost no energy. The only cost is encoding the vector into light and reading out the answer, but the vector is of length N , not N 2 . So the energy used will scale as N , not N 2 , as bigger vectors and matrices are multiplied.
Since Goodman’s 1978 proposal, the quality of the available optical elements has improved dramatically. Indeed, in 2022, a team of researchers at Cornell led by McMahon demonstrated an optical neural network that consumed less than one photon of energy—approximately 2.5 x 10 -19 joules—per multiplication operation. In a standard computer, a single multiply operation can take between 10 -14 and 10 -12 joules, which suggests that optics could possibly reduce energy use by a factor of 100,000.
But McMahon cautions against such extrapolation. First, the cost of encoding data into optics was not included in the calculation. Even if the scaling is favorable with light, in practice it’s not just the scaling that matters, it’s also the coefficients. Let’s say the energy cost of doing one multiplication with electronics is 1 picojoule, and the cost to encode one vector element with optics is 1,000 picojoules. The electronic approach still wins if the length N of the vector is fewer than 1,000 elements.
McMahon says this example is not far off from reality: “One of the key questions is, How big do you have to make the system before you start winning? Lots of people have done this calculation for various kinds of neural networks . The very rough answer is somewhere between N equals 1,000 and N equals 10,000.”
It’s also important to keep in mind a computing principle known as Amdahl’s law .
“Amdahl’s law shouldn’t even be called a law. It’s a simple piece of common sense,” McMahon says. “If you have something that’s going to accelerate a computation, and it can accelerate 90 percent of the code, but there’s 10 percent of the code that it can’t accelerate, then the maximum possible speedup you can get is 10x.”
With AI inferencing, the point is to save energy, not time. But Amdahl’s observation still applies. What fraction of today’s AI inferencing code can be better handled with optics is a complicated question that varies by approach. But it’s not 100 percent. Matrix-vector multiplication, or matrix-matrix multiplication, is only part of the inference workload—there are other, nonlinear operations that need to be performed on the data.
Taking all of these caveats into account, optical-computing companies today are aiming for a 10- to 80-fold improvement in energy efficiency over today’s GPUs. The question is whether they can deliver on that promise at scale before traditional electronics catches up to those gains.
Bridging the gap in free space
The startup Lumai , based in Oxford, England, has taken perhaps the most direct inspiration from Goodman and his team. Lumai’s head of product, Phillip Burr , declares that a free-space approach, where the light travels through air rather than on-chip waveguides, is what will enable the company’s tech to beat GPUs.
Instead of the fixed attenuator used in Goodman’s proposal, Lumai is using a modern programmable device called a spatial light modulator—a pixelated membrane akin to a fancy liquid-crystal display that can modulate both the intensity and phase of light. The rest of the design is more sophisticated in practice but similar in principle to Goodman’s design: A vector is encoded into light, copied N times via a lens, passed through the spatial light modulator, collected with another lens, and the result is converted back to digital electronics. The electronics then handle the nonlinear functions. “A full vector-matrix multiply completes in a single optical cycle,” Burr says.
To accurately compare the energy expenditures of the optical and GPU approaches, you must first specify which part of the workload you’re considering. Modern inference is often split into two phases: prefill , where the input query is processed and propagated through the model, and decode , where the output is generated one word (or token) at a time. Prefill is the more computationally intensive of the two. It basically involves doing vector-matrix multiplications for the same matrix over and over. Decode also involves vector-matrix multiplies, but not as intensively—the time-consuming part of decode is reading the necessary matrices in and out of memory.
Vector-matrix multiplication with free-space optics
The Goodman optical computer is the simplest version of a photonic vector-matrix multiplier. Here, the vector elements are encoded into the intensity of light from three light sources, such as LEDs . Then, a cylindrical lens spreads the light vertically. The matrix elements are encoded into the opacity of an attenuator, such that when light passes through each pixel it is multiplied (attenuated) by that matrix value. Finally, another cylindrical lens collects the light horizontally, performing the accumulation step, and a photodetector measures the output vector.
Lumai, and other companies in the space, are focusing on the compute-intensive prefill phase of inference. Lumai claims that its Iris Nova optical server, made available for evaluation in April 2026, is a proof of concept, and its next-generation Iris Aura server will do prefill using roughly one-tenth of the energy a GPU would use.
The startup Neurophos ,

[truncated]
