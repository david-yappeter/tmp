- What steps would you take to decompose the service into smaller, more manageable services?
    - Determine which services can be decoupled first without breaking the others service. In this case we will start with `File Upload` since most of upload service is a independent service that can be decoupled from the monolithic service. and then `authentication` and then .`user_data processing`.
    - If the service rely on others service, we need to ensure there is a way for the service to communicate to other service, such as via `HTTP request` or `RPC request`, depends on the approach we plan to used. There are others several way for example `database duplication`. So for each service will have its own database, but some datas are duplicate within the databases



- How would you ensure that the new system is backward compatible with the old one during the transition?

Using API Versioning, so the for example the old system will used a prefix `/v1/* `and the new system will use new path `/v2/*`
With this approach the previous api can still be used while the new version is still in transition. After a while you can deprecate the old endpoints gradually.
