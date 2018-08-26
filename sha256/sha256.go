package main


import (
    "fmt"
    "crypto/sha256"
    "crypto/sha1"
    "encoding/hex"
    "encoding/base64"
    "os"
    "log"
    "io"
)




func main() {
        
    s:="Secret"
    b:="Gaasdasasdasdadsdadasdsd adssdsd"

    l:=Sys_sha(s,b)
    fmt.Println(l)


    s="Secrets"
    b="Gaasdasasdasdadsdadasdsd adssdsd sssss пример написания простого решения"

    l=Sys_sha(s,b)
    fmt.Println(l)
  
    l=Sys_shas("Pass")
    fmt.Println(l)

    l=Sys_shax("Yssss ddd")
fmt.Println(l)    

}



func Sys_shas(Txt string) string{
    sum := sha256.Sum256([]byte(Txt))
    r:=fmt.Sprintf("%x", sum)
    // r:=base64.StdEncoding.EncodeToString([]byte(sum))
    
    return  r
}



// 
// https://play.golang.org/p/dPr9lt1uIS
func Sys_shax(Txt string) string{
    h := sha1.New()
    h.Write([]byte(Txt))
    sha1_hash := hex.EncodeToString(h.Sum(nil))
    // fmt.Println(s, sha1_hash)
    return sha1_hash
}


/********************************************************************************************************************************
 * See 2 (end of page 4) http://www.ietf.org/rfc/rfc2617.txt
 *
 * "To receive authorization, the client sends the userid and password,
 * separated by a single colon (":") character, within a base64
 * encoded string in the credentials."
 *
 * It is not meant to be urlencoded.
 *********************************************************************************************************************************/
func basicAuth(username, password string) string {
    auth := username + ":" + password
    return base64.StdEncoding.EncodeToString([]byte(auth))
}


/******************************************************************************************************
 * TITLE  :  Расчет кодирование SHA1 Hash
 * REMARK :
 ******************************************************************************************************/
func Sys_sha(Salt, Body string) string {

    if Salt == "" || Body == "" {
       return "SHA Error."
    }

    b := Salt + Body
    h := sha1.New()
    h.Write([]byte(b))
    s := hex.EncodeToString(h.Sum(nil))
    return s
}

/******************************************************************************************************
 * TITLE  :  Расчет кодирование SHA1 Hash
 * REMARK :
 ******************************************************************************************************/
func Sys_sha256(Salt, Body string) string {
    b := Salt + Body
    h := sha256.New()
    h.Write([]byte(b))
    s := hex.EncodeToString(h.Sum(nil))
    return s
}



func Sys_Sum256() {
    sum := sha256.Sum256([]byte("hello world\n"))
    fmt.Printf("%x", sum)
    // Output: a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447
}

func Sys_New() {
    h := sha256.New()
    h.Write([]byte("hello world\n"))
    fmt.Printf("%x", h.Sum(nil))
    // Output: a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447
}

func Sys_New_file() {
    f, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    h := sha256.New()
    if _, err := io.Copy(h, f); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%x", h.Sum(nil))
}
