package models

import "fmt"

type Migration struct {
	ID uint	`json:"id"`
	Name string `json:"name"`
	Aws_Access_Key string `json:"aws_access_key"`
	Aws_Secret_Key string `json:"aws_secret_key"`
	Aws_Bucket_Name string `json:"aws_bucket_name"`
	Path string `json:"path"`
	Max_Object_Size uint `json:"max_object_size"`
// - Max creation date
	Lyve_Access_Key string `json:"lyve_access_key"`
	Lyve_Secret_Key string `json:"lyve_secret_key"`
	Lyve_Bucket_Name string `json:"lyve_bucket_name"`
	Status string `json:"status"`
// - Created date
}