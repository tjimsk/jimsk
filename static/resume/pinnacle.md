## PINNACLE HOLDINGS LTD.

<a name="overview"></a>
### Overview
Pinnacle Holdings Ltd. is an affiliate of a Chinese gaming service company established over than 20 years ago.  Its services include a variety of online casino and lottery games across multiple regions in Asia.  The affiliate company began its business in Q1 of 2018.  Its line of business is software development.  

As an established organisation beginning a new operation in Hong Kong, some of the technical challenges were:
- inheritance of a large code base from a system proven to be non scalable
- desire to learn from past development mistakes and implement more stringent coding standards

<a name="role"></a>
### Role
- Position: Senior Software Engineer
- Location: Kowloon, Tsim Sha Tsui
- Period: November 2017 - January 2018

<a name="stack"></a>
### Technology Stack
- Database: MySQL
- Backend: Java / framework core-ng
- Frontend: React/Redux
- Infrastructure: 
	- Kubernetes
	- Docker
	- Jenkins
	- ElasticSearch
	- Kafka

<a name="description"></a>
### Description
In this role, my work was concentrated on developing the CI/CD pipeline, scaffolding of a Chinese lottery game frontend and a general webpack configuration for later games' frontends. 

The pipeline consisted of a Kubernetes cluster with groovy scripting to periodically: 
- pull repositories
- run tests 
- build docker images
- push to gcr
- set cluster pod images

I implemented a single page web app using react-router and redux, which was ideal for keeping the state of multiple game modes/inputs/outcomes of a lottery game.  The game itself had different rules based on the game mode.  Most of its logic lived in the backend.  

The webpack configuration included applying eslint and stylelint standardization rules.

<a name="reasonforleaving"></a>
### Reason For Leaving
The language barrier with the Products team based in the Philippines (Mandarin only) was too big to continue leading the frontend development.
