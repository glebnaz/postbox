package entities

//User is entity type
//with this struct you can send email
//user include smtp cred
type User struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	SMTPHost    string `json:"smtp_host" bson:"smtp_host"`
	SMTPAddress string `json:"smtp_address" bson:"smtp_address"`
	SMTPUser    string `json:"smtp_user" bson:"smtp_user"`
	SMTPPass    string `json:"smtp_pass" bson:"smtp_pass"`
}
