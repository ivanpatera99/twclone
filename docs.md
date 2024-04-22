# TwClone

## project structure
- /cmd: app entry points
- /internal: internal app code
    - adapters: tech stack logic (database queries, interaction with other services)
        - sql
        - documents (not implemented)
        - graph (not implemented)
    - app: adapter injection, use case handling
        - http: inside this package the app will be tailored to work with a regular http server with sql adapters
            - handlers: http handlers
    - domain: business logic
        - entities
        - usecases: business rules
        - ports: interface ports for use cases

This project structure ensures separation of concerns. First, where the tailoring and orchestration of the different external services is defined in the domain layer aswell as the interfaces that this external services should fulfill. Second, the adapter layer has the structures with the methods necessary to implement their desired interfaces, within this methods the specifics of the selected tech stack emerges, if you check the `/internal/adapters/sql` you will find packages with sql specific code intended to interact with the sqlite db that ships along side this app. Lastly the app package defines the entry points of the server and injects the use cases dependencies, aswell as any other needed configuration, it should reside here.

## tech decisions
For development and demonstration purposes this application is built and shipped using an in-memory database. The priority was set to get an mvp and work with sqlite to build the initial data model, then define areas of improvement that can be implemented to achieve the desired performance. Here we will dive deep into some of the conundrums that arrised while thinking the best approach for a production enviroment:

1. sql vs nosql: this first version models the problem in a relational style. With further analysis on the problem when scaling a document-oriented db might be a better approach for two reasons:
    - Sharding: document-oriented dbs are generally designed to scale using sharding.
    - Flexibility: the schema-less document aproach lets us work on new features that require changes in the schema (such as likes, reposts, posts with images, responses, etc) without the need to going throughout a migration.
2. social graph queries: both sql and nosql databases introduce problems when querying social graphs, having a complementary graph db can introduce serious benefits when performing this type of queries, with lower latency than its sql and nosql counterparts, it can reduce the necessity of convoluted caching strategies to not overload the db.
3. caching: while not implemented in the app caching will be necessary for frequent accessed data such as social graphs, frequent/relevant profiles. This was not a huge `conundrum` as described at the beggining since every application that aims to fullfil requests to millons of active users should and have a caching layer.

## human considerations
As relevant as the technology stack to deploy is the team that will have to maintain the different parts of the stack, in the delicate balance between complexity overhead vs saas cost, the product should bare in mind a range of technological solutions to implement the needed services. I'll describe three solutions with different ratios of complexity overhead/solution cost:
1. Serverless: the solution with the minimal complexity overhead but with an expensive solution cost when scaling to millons of users. A set of products from aws to implement are `Lambda` (code execution), `DynamoDB` (document storage), `Neptune` (graph storage) and `ElastiCache` with memcached as a cache motor.
2. Containers: a containeraized solution provides fine tune control over the different properties, with more complexity overhead but with less saas cost would be to deploy the solution with a tool like `docker compose` alongside the caching layer and a db that can support both documents and graphs like `arangoDB`, this forces us to have people maintaining this deployed infrastructure, but you can laverage more flexibility when moving from cloud to cloud as your code is not thightly coupled to any cloud whatsoever.
3. K8s: with the most complexity overhead but with the most flexibility to fine tune your service orchestration a K8s cluster might be a good idea to handle millons of users spending the minimum on saas products. Instead of arangoDB as detailed in `2`, the proposed solution is to deploy separate document and graphs dbs first, to let them scale on separate tracks but also to provide better performance since we can fine tune the document and the graph sides of the application independently. A complete selection of products to integrate to this solutions might be: your preferred flavor to deploy a k8s cluster, `ScyllaDB` for documents would be ideal since it's a db that could help us scale with its design arround multi-core utilization makes it the best desition to handle this high throughput and low latency scenario, `neo4j` for graphs and `memcached` for the caching layer that can also be sharded to work in a high availabilty scenario as presented.




    