---
source: "https://discoveredmaterials.com/research/accelerating-quantum-espresso/"
hn_url: "https://news.ycombinator.com/item?id=50430969"
title: "Speeding up Quantum Espresso by 36x using Claude Fable 5.2"
article_title: "Accelerating Quantum ESPRESSO by 36x with Claude Fable 5.2 — Discovered Materials"
image: "https://discoveredmaterials.com/assets/og-image.png"
author: "advaith08"
captured_at: "2026-11-09T20:28:14Z"
capture_tool: "hn-digest"
hn_id: 50430969
score: 2
comments: 0
posted_at: "2026-11-09T19:54:44Z"
tags:
  - hacker-news
---

# Speeding up Quantum Espresso by 36x using Claude Fable 5.2

- HN: [50430969](https://news.ycombinator.com/item?id=50430969)
- Source: [discoveredmaterials.com](https://discoveredmaterials.com/research/accelerating-quantum-espresso/)
- Score: 2
- Comments: 0
- Posted: 2026-11-09T19:54:44Z

## Translation

Title: Speeding up Quantum Espresso by 36x using Claude Fable 5.2
Article title: Accelerating Quantum ESPRESSO by 36x with Claude Fable 5.2 — Discovered Materials
Description: We used Claude Fable 5.2 and a set of test suites to speed up DFPT phonon calculations by 36x and hybrid functional DFTs by 9x by patching Quantum ESPRESSO.

Article text:
Accelerating Quantum ESPRESSO by 36x with Claude Fable 5.2 — Discovered Materials Challenges Research News Talk to us Accelerating Quantum ESPRESSO
by 36x with Claude Fable 5.2
We used Claude Fable 5.2 and a set of test suites to speed up DFPT phonon calculations by 36x and hybrid functional DFTs by 9x by patching Quantum ESPRESSO. You can find the open-source patched version on our Github here .
Density Functional Theory (DFT) is the workhorse of computational material science. It is the most widely used method for simulating matter at the atomic scale, and consistently among the largest consumers of time on national supercomputers. DFT is used in nearly every field of materials science: semiconductors, battery cathodes and solid electrolytes, catalysts for fertilizer and hydrogen production, solar absorbers, alloys for aerospace, and in drug discovery, the binding energies, charge distributions, and reaction pathways that classical force fields cannot capture.
Quantum ESPRESSO (QE) is the most widely used open-source repository for running DFT calculations, and is a codebase of roughly 600,000 lines of Fortran code. It provides a suite of modular DFT routines. Conventional DFT is comparatively cheap to run but has well-known flaws: it systematically underestimates band gaps and misplaces defect levels, which matters enormously for semiconductor applications. Hybrid functionals, which more accurately account for the Pauli exclusion principle, offer a promising remedy to these limitations of conventional DFT. However, they are computationally much slower to run, limiting their application to certain smaller toy systems, rather than realistic large scale simulations.
Beyond static properties, QE includes routines based on density functional perturbation theory (DFPT), which allows researchers to compute how atoms in a material vibrate (phonons). Phonons help us predict crystal stability, heat capacity, thermal conductivity, dielectric response, conventional superconductivity and many other phenomena. DFPT is an order of magnitude more expensive than DFT. As a result, phonon data exists for only a subset of known materials in the Materials Project , and dynamic stability checks are often skipped when screening candidates in materials discovery.
Rewrites QE kernels for the GPU
Mixed precision, batching, symmetry
Discovered Materials Test Suite
Representative novel test suite
Reference results from stock QE
Only changes that are faster and match reference results
We run Fable 5.2 in an optimization loop against our test suite to speed-up Quantum Espresso
We gave Claude Fable 5.2 the tasks of speeding up QE’s phonon code and hybrid functional algorithms on GPUs. To verify that its speedups were correct, we used standard benchmark test suites, as well as hidden test sets spanning representative systems to ensure the speedups were not overfit to publicly known test suites.
Table 1: Speedups observed across benchmarks for hybrid DFT calculations
Table 2: Speedups observed across benchmarks for phonon (DFPT) calculations
We also do a correctness check of all speedups, and observe that any value mismatches are within measurement errors of experimental characterization.
Result 1: 9x speedup for hybrid functionals (and multi-GPU support!)
On the hybrid-functional task, Fable 5.2 sped up QE’s code by roughly 9x (geometric mean) and, without being asked, implemented a multi-GPU version of the algorithm. Multi-GPU support is useful as it enables parallelization across devices, which in turn allows much larger simulation cells.
The speedup is driven by the following changes:
GPU-resident orbitals. The orbital buffer stays on the GPU between applications of the exchange operator. Stock QE copied it from host to GPU on every application, resulting in large PCIe traffic latencies.
Batched exchange operator. Stock QE processes band pairs one at a time with many small GPU launches. Here they are grouped into blocks sized to fit the GPU’s L2 cache, and each block is handled by batched FFTs and a few fused kernels.
GPU-resident setup. Building the real-space orbitals at each exchange initialization (FFTs and symmetry rotations) now runs on the GPU instead of the CPU. This removes a serial bottleneck and a large host-to-GPU transfer.
Exchange stress on the GPU. The exchange contribution to the stress uses the same batched machinery.
Latency-aware small eigensolves. The small subspace diagonalizations in the Davidson solver run on the CPU, because for problems that size the GPU solver’s launch and synchronization overhead outweighs the computation.
Result 2: 36x speedup for phonon calculations
Fable’s QE DFPT phonon code is roughly 36x faster than stock QE on the same GPU, (geometric mean across 13 benchmark cases). The total wall-clock time for running our benchmark suite fell from 3.3 hours to under 4 minutes .
The gains come from four main changes, all made while preserving backward compatibility with existing QE inputs and workflows:
Mixed precision. Running the bulk of the linear-response solves in lower precision where accuracy permits, and promoting to double precision only where it matters.
Exploiting symmetry. Using the crystal’s symmetry to avoid recomputing equivalent perturbations.
A better preconditioner. Improving convergence of the iterative solver so fewer iterations are needed per perturbation.
GPU-resident setup. Moving setup work that previously ran on the CPU onto the GPU, removing a serial bottleneck and host-device transfers.
Implementing the above changes requires (at minimum) an understanding of electronic structure physics, GPU kernels and Fortran - a combination of skills that is rare to find amongst experts in the field.
Using these speedups to simulate NV-centers in diamond
To test the two speedups, we put them together on the nitrogen-vacancy (NV) center problem in diamond. These NV defect diamond-based quantum sensors are a leading candidate for room-temperature qubits in quantum computing applications. The usefulness of these NV centers depend on their defect level positions (which need Hybrid functionals to compute) and electron-phonon coupling (which needs DFPT).
In our benchmark, the phonon calculation for a negatively charged NV center in a 63-atom diamond supercell went from 77 s on stock GPU QE to 6 s (13x faster) with the optimized code. With the multi-GPU hybrid-functional implementation, the same system can be treated at hybrid accuracy on four H100s. Together, this turns a workflow that previously took several days into something that runs in a few hours, and makes it feasible to screen hundreds of candidate defect-host combinations for quantum applications.
Towards fast, accurate universal functionals
In DFT, accuracy comes in rungs: Perdew’s “Jacob’s ladder” climbs from (semi-)local functionals (LDA, GGA) through meta-GGAs and hybrids to double hybrids and RPA, each rung closer to the “heaven” of chemical accuracy, but each an order of magnitude or more costlier than the last. For decades the field has had to pick a rung and live with its errors. We view breaking this speed-accuracy tradeoff as a Grand Challenge in material science .
Designing or speeding up functionals is a uniquely well-suited problem for LLMs, as new functionals can be run in simulation and verified using existing slower methods. Looking forward, we are optimistic about the ability of LLMs to rapidly collapse the hierarchy of Jacob’s ladder, and will aim to bring hybrid, double-hybrid, and RPA-level accuracy to the speed of today’s GGA calculations.
To get notified on new research, sign up for our mailing list.
©2026 Matforge Inc. All rights reserved · Terms of Service · Privacy Policy

## Original Extract

We used Claude Fable 5.2 and a set of test suites to speed up DFPT phonon calculations by 36x and hybrid functional DFTs by 9x by patching Quantum ESPRESSO.

Accelerating Quantum ESPRESSO by 36x with Claude Fable 5.2 — Discovered Materials Challenges Research News Talk to us Accelerating Quantum ESPRESSO
by 36x with Claude Fable 5.2
We used Claude Fable 5.2 and a set of test suites to speed up DFPT phonon calculations by 36x and hybrid functional DFTs by 9x by patching Quantum ESPRESSO. You can find the open-source patched version on our Github here .
Density Functional Theory (DFT) is the workhorse of computational material science. It is the most widely used method for simulating matter at the atomic scale, and consistently among the largest consumers of time on national supercomputers. DFT is used in nearly every field of materials science: semiconductors, battery cathodes and solid electrolytes, catalysts for fertilizer and hydrogen production, solar absorbers, alloys for aerospace, and in drug discovery, the binding energies, charge distributions, and reaction pathways that classical force fields cannot capture.
Quantum ESPRESSO (QE) is the most widely used open-source repository for running DFT calculations, and is a codebase of roughly 600,000 lines of Fortran code. It provides a suite of modular DFT routines. Conventional DFT is comparatively cheap to run but has well-known flaws: it systematically underestimates band gaps and misplaces defect levels, which matters enormously for semiconductor applications. Hybrid functionals, which more accurately account for the Pauli exclusion principle, offer a promising remedy to these limitations of conventional DFT. However, they are computationally much slower to run, limiting their application to certain smaller toy systems, rather than realistic large scale simulations.
Beyond static properties, QE includes routines based on density functional perturbation theory (DFPT), which allows researchers to compute how atoms in a material vibrate (phonons). Phonons help us predict crystal stability, heat capacity, thermal conductivity, dielectric response, conventional superconductivity and many other phenomena. DFPT is an order of magnitude more expensive than DFT. As a result, phonon data exists for only a subset of known materials in the Materials Project , and dynamic stability checks are often skipped when screening candidates in materials discovery.
Rewrites QE kernels for the GPU
Mixed precision, batching, symmetry
Discovered Materials Test Suite
Representative novel test suite
Reference results from stock QE
Only changes that are faster and match reference results
We run Fable 5.2 in an optimization loop against our test suite to speed-up Quantum Espresso
We gave Claude Fable 5.2 the tasks of speeding up QE’s phonon code and hybrid functional algorithms on GPUs. To verify that its speedups were correct, we used standard benchmark test suites, as well as hidden test sets spanning representative systems to ensure the speedups were not overfit to publicly known test suites.
Table 1: Speedups observed across benchmarks for hybrid DFT calculations
Table 2: Speedups observed across benchmarks for phonon (DFPT) calculations
We also do a correctness check of all speedups, and observe that any value mismatches are within measurement errors of experimental characterization.
Result 1: 9x speedup for hybrid functionals (and multi-GPU support!)
On the hybrid-functional task, Fable 5.2 sped up QE’s code by roughly 9x (geometric mean) and, without being asked, implemented a multi-GPU version of the algorithm. Multi-GPU support is useful as it enables parallelization across devices, which in turn allows much larger simulation cells.
The speedup is driven by the following changes:
GPU-resident orbitals. The orbital buffer stays on the GPU between applications of the exchange operator. Stock QE copied it from host to GPU on every application, resulting in large PCIe traffic latencies.
Batched exchange operator. Stock QE processes band pairs one at a time with many small GPU launches. Here they are grouped into blocks sized to fit the GPU’s L2 cache, and each block is handled by batched FFTs and a few fused kernels.
GPU-resident setup. Building the real-space orbitals at each exchange initialization (FFTs and symmetry rotations) now runs on the GPU instead of the CPU. This removes a serial bottleneck and a large host-to-GPU transfer.
Exchange stress on the GPU. The exchange contribution to the stress uses the same batched machinery.
Latency-aware small eigensolves. The small subspace diagonalizations in the Davidson solver run on the CPU, because for problems that size the GPU solver’s launch and synchronization overhead outweighs the computation.
Result 2: 36x speedup for phonon calculations
Fable’s QE DFPT phonon code is roughly 36x faster than stock QE on the same GPU, (geometric mean across 13 benchmark cases). The total wall-clock time for running our benchmark suite fell from 3.3 hours to under 4 minutes .
The gains come from four main changes, all made while preserving backward compatibility with existing QE inputs and workflows:
Mixed precision. Running the bulk of the linear-response solves in lower precision where accuracy permits, and promoting to double precision only where it matters.
Exploiting symmetry. Using the crystal’s symmetry to avoid recomputing equivalent perturbations.
A better preconditioner. Improving convergence of the iterative solver so fewer iterations are needed per perturbation.
GPU-resident setup. Moving setup work that previously ran on the CPU onto the GPU, removing a serial bottleneck and host-device transfers.
Implementing the above changes requires (at minimum) an understanding of electronic structure physics, GPU kernels and Fortran - a combination of skills that is rare to find amongst experts in the field.
Using these speedups to simulate NV-centers in diamond
To test the two speedups, we put them together on the nitrogen-vacancy (NV) center problem in diamond. These NV defect diamond-based quantum sensors are a leading candidate for room-temperature qubits in quantum computing applications. The usefulness of these NV centers depend on their defect level positions (which need Hybrid functionals to compute) and electron-phonon coupling (which needs DFPT).
In our benchmark, the phonon calculation for a negatively charged NV center in a 63-atom diamond supercell went from 77 s on stock GPU QE to 6 s (13x faster) with the optimized code. With the multi-GPU hybrid-functional implementation, the same system can be treated at hybrid accuracy on four H100s. Together, this turns a workflow that previously took several days into something that runs in a few hours, and makes it feasible to screen hundreds of candidate defect-host combinations for quantum applications.
Towards fast, accurate universal functionals
In DFT, accuracy comes in rungs: Perdew’s “Jacob’s ladder” climbs from (semi-)local functionals (LDA, GGA) through meta-GGAs and hybrids to double hybrids and RPA, each rung closer to the “heaven” of chemical accuracy, but each an order of magnitude or more costlier than the last. For decades the field has had to pick a rung and live with its errors. We view breaking this speed-accuracy tradeoff as a Grand Challenge in material science .
Designing or speeding up functionals is a uniquely well-suited problem for LLMs, as new functionals can be run in simulation and verified using existing slower methods. Looking forward, we are optimistic about the ability of LLMs to rapidly collapse the hierarchy of Jacob’s ladder, and will aim to bring hybrid, double-hybrid, and RPA-level accuracy to the speed of today’s GGA calculations.
To get notified on new research, sign up for our mailing list.
©2026 Matforge Inc. All rights reserved · Terms of Service · Privacy Policy
