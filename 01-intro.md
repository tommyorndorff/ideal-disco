---
marp: true
theme: gaia
_class: invert
---
# Howdy
---

<!-- paginate: true -->

## What We'll Cover
In the first couple of weeks, we'll cover the basics of AWS, split into primary categories:
- Ecosystem & Responsibility Model
- Networking
- Compute
- Storage
- Security

---

## Certificate Path Overview

- Best explained by the [AWS diagram](https://aws.amazon.com/certification/).
- Broken into:
  - 'Foundational'
  - 'Associate'
  - 'Professional'
  - 'Specialty'
- Descriptions & costs on [aws.amazon.com](https://aws.amazon.com/certification/exams/?nc2=sb_ce_exm)

---

## Online Resources or Trainings

- AWS provides `AWS Cloud Practitioner Essentials` [online training](https://aws.amazon.com/training/digital/aws-cloud-practitioner-essentials/)
- Pluralsight has great [Associate](https://www.pluralsight.com/courses/aws-certified-solutions-architect-associate) & [Professional](https://www.pluralsight.com/paths/aws-certified-solutions-architect-professional) trainings
- [AWS Well-Architected Framework](https://aws.amazon.com/architecture/well-architected/?wa-lens-whitepapers.sort-by=item.additionalFields.sortDate&wa-lens-whitepapers.sort-order=desc)
- [AWS Whitepapers & Guides](https://aws.amazon.com/whitepapers/?whitepapers-main.sort-by=item.additionalFields.sortDate&whitepapers-main.sort-order=desc&awsf.whitepapers-content-type=*all&awsf.whitepapers-tech-category=*all&awsf.whitepapers-industries=*all&awsf.whitepapers-business-category=*all&awsf.whitepapers-global-methodology=methodology%23well-arch-framework)
- AWS Sample Questions (e.g. [SA Assoc.](https://d1.awsstatic.com/training-and-certification/docs-sa-assoc/AWS-Certified-Solutions-Architect-Associate_Sample-Questions.pdf))

---

<!-- footer: 101 Overview -->

## 101 Overview, Part 1
- **Brief History of Cloud**
- **Shared Responsibility**
- **AWS Ecosystem**

---

<!-- footer: 101 Overview -->

## 101 Overview, Part 2
- **Networking**
- **Compute** 
- **Storage**
- **Database**
- **Security**

---

<!-- footer: 201 Overview -->

## 201 Overview: 
### Networking
- VPC Fundamentals
- Subnets, Route Tables, NACLs
- NAT, Internet, & Egress-Only Gateways
- Hybrid Implementations

---

## 201 Overview: 
### Compute
- "Compute primitives"
- EC2 (Elastic Compute Cloud)
- Types of EC2 ([useful link](https://instances.vantage.sh/))
- Lambda & Serverless

---

## 201 Overview: Storage


---

## 201 Overview: Database

---

## 201 Overview: Security

---

<!-- footer: Bits & Bobs -->

## AWS Account

Need access to an AWS account where you're admin.  EG folks, we'll use the old "GCO Sandbox" account in us-east-1

Otherwise:
- Sign-up for a free account
- Enable multi-factor auth (MFA)
- Stay in the free tier & make a billing alert

---

## 🔨 Must Have Tools: AWS ☁️
- AWS Account: [Offical AWS 'Getting Started'](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-prereqs.html)
- AWS CLI: [docs.aws.amazon.com](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
  - macOS: use `brew install awscli`
  - Windows: use the [.msi](https://awscli.amazonaws.com/AWSCLIV2.msi)

---

## 🔨 Must Have Tools: Python 🐍 
- Python 3.8 or above ([link](https://www.python.org/downloads/))
- 'pip' Module for Python
- [AWS SDK for Python](https://aws.amazon.com/sdk-for-python/) (aka Boto3) 
  - `pip install boto3`

---

## 'Make Life Easier' Tools 🧑‍🌾
- Visual Studio Code (aka VSCode) ([link](https://code.visualstudio.com/download))
- Github account & CLI 
- Docker Desktop
- [Windows Subsystem for Linux](https://docs.microsoft.com/en-us/windows/wsl/install) (WSL)