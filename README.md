# Service

Service builder on go.

This pkg is mean to make building services as fast a possible and as simple as possible as well.
Trying to keep a code base minimalist but as rich as possible allowing the more commune micro service user case to be at least consider in the functionality of this pkg.
Rather than thinking this as full on framework think of it as struct to service promoter.

Major Dependencies:

- Log with [logrus](https://github.com/sirupsen/logrus)
- http router with [httprouter](https://github.com/julienschmidt/httprouter)

Prove of Concept:

- A simple [Lot Service](https://github.com/posttul/lot-service)