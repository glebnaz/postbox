package errors

//BadCred error for Bad Credential
var BadCred = Exception{Code: SecureCode, Message: "Bad Credentials"}

//Unauthorized error for unauthorized user
var Unauthorized = Exception{Code: SecureCode, Message: "Unauthorized"}

//DataBaseConnection error for fail database connection
var DataBaseConnection = Exception{Code: DataBaseCode, Message: "Bad Connection to Store"}

//DataBaseOperation error for DataBaseOperation
var DataBaseOperation = Exception{Code: DataBaseCode, Message: "Internal Storage Problem"}

//EmptyUsersIDs error if empty users Ids
var EmptyUsersIDs = Exception{Code: EntityCode, Message: "Empty Users IDs"}

//EmptyUsers error if empty users
var EmptyUsers = Exception{Code: EntityCode, Message: "Empty Users"}
