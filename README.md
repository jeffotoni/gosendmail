# gosendmail - SMTP

This repo is a simple net/smtp abstraction for sending emails using SMTP.
With it we can send emails With copy, with blind copy and attachments.
The body is plain text, we're still going to make a version to accept HTML, at the moment only text/plain in the Body.
The attachment must pass the complete path so that the package can open the file.
There are 4 environment variables to be configured, they are:

	- EMAIL_HOST
	- EMAIL_USERNAME
	- EMAIL_PASSWORD
	- EMAIL_PORT

Below is an example of how you might use pkg.
```go
package main

import gse "github.com/jotonilima/gosendmail"
import "log" 


func main() {
 	email := gse.New()
 	if email == nil {
 		log.Println("Error New() check the required fields: 
		 EMAIL_HOST,EMAiL_USERNAME,EMAIL_PASSWORD,EMAIL_PORT")
 		return
 	}
 	m := gse.NewMessage("Title here", "Body message here.")
 	m.To = []string{"<to-email>@gmail.com"}
 	m.CC = []string{"<copy1-email>@gmail.com", "<copy2-email>@gmail.com"}
 	m.BCC = []string{"<bc-email>@gmail.com"}
 	m.AttachFile("/path/to/file1.pdf")
 	m.AttachFile("/path/to/file2.pdf")
 	if email.Send(m) != nil {
 		log.Println("Error when sending:", err.Error())
 		return
 	}
 	log.Println("Sent with success")
}

```

To run the code above:
```bash
$ export EMAIL_HOST="smtp.<YOUR-SERVER>.com"
$ export EMAIL_USERNAME="<USER>"
$ export EMAIL_PASSWORD="<PASSWORD>"
$ export EMAIL_PORT="587"
$ go mod init <your-pkg>
$ go mod tidy
$ go run main
```