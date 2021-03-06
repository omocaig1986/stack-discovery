<div align="center">

# P2PFaaS

A Framework for FaaS load balancing  | _`stack-discovery` repository_

![License](https://img.shields.io/badge/license-GPLv3-green?style=flat)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/5471195919a744ab8a9bba3a6be9169b)](https://www.codacy.com/gl/p2p-faas/stack-discovery/dashboard?utm_source=gitlab.com&amp;utm_medium=referral&amp;utm_content=p2p-faas/stack-discovery&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/p2p-faas/stack-discovery)](https://goreportcard.com/badge/gitlab.com/p2p-faas/stack-discovery)

</div>

# Introduction

The P2PFaaS is a framework that allows you to implement a load balancing/scheduling algorithm for FaaS.

For a detailed information about the framework you can read my MSc thesis at [raw.gpm.name/theses/master-thesis.pdf](https://raw.gpm.name/theses/master-thesis.pdf). If you are using P2PFaaS in your work please cite [https://ieeexplore.ieee.org/document/8964273/](https://ieeexplore.ieee.org/document/8964273/):

```bibtex
@article{8964273,
    author={Beraldi, Roberto and Proietti Mattia, Gabriele and Magnani, Giacomo},
    journal={IEEE Transactions on Cloud Computing},
    title={Power of random choices made efficient for fog computing},
    year={2020},
    volume={},
    number={},
    pages={1-1},
    doi={10.1109/TCC.2020.2968443}}
```

# Repository

This is the a simple discovery service of the framework. It's written in Go and it is packaged with Docker. Instructions are the same as `stack-scheduler` repository.