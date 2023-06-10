# Intro Module

## Types Of Cloud Computing

- **Infrastructure as Service(IaaS)**
  - Provides Networking, compute, data in it's raw form
  - High level of flexibilty
  - Basic building blocks of cloud IT
  - We can easily migrate from on-premises IT to cloud IT

- **Platform as Service(PaaS)**

  - Removes the need for managing underlying infrastructure
  - We can focus on deployment and management of out applications and data
  - Examples include: Openshift, Windows Azure, Google Cloud etc.

- **Software as Service(Saas)**

  - Completed product that is run and managed by service provider
  - Ex: google workspace, Dropbox, Cisco Webex etc.

Resources:

- IaaS vs SaaS vs PaaS [Link](https://www.bmc.com/blogs/saas-vs-paas-vs-iaas-whats-the-difference-and-how-to-choose)
- What is PaaS [Link](https://www.techtarget.com/searchcloudcomputing/definition/Platform-as-a-Service-PaaS)
  
## AWS Regions

[**OFFICIAL PAGE**](https://aws.amazon.com/about-aws/global-infrastructure/?p=ngi&loc=1)

Read more about **Regions**, **Availablity Zones**, **Services**, **Outposts** etc. [HERE](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/?p=ngi&loc=2&refid=7541ebd3-552d-4f98-9357-b542436aa66c)

## Factors to consider while chosing an AWS region

The decision to choose which region our application is deployed in depends mainly on the following factors:

- **Compliance**: Often times, due to governance and legal requirements, the data never leaves a region without explicit permission.

    Eg: An application's data serving its customers in USA needs to reside within the USA

- **Proximity**: Proximity to customers will help reduce latency.

- **Available services**: We need to make sure the selected AWS region offers all the services required for our application

- **Pricing**: Pricing for each region varies and be learnt about on the service pricing page.
