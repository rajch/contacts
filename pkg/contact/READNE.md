# contact

Package contact defines the data model for a "Contact" entity, and an interface that 
lightly follows the repository pattern for storing and retrieving contacts.

It is accompanied by a sub-package, contacttest, which contains a test suite for the 
repository. Implementations of the repository interface should include the sub-package, and call the test function in _their tests only_.
