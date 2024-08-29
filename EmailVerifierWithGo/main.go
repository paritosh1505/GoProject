package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	
	scannerentry := bufio.NewScanner(os.Stdin)

	for scannerentry.Scan(){
		validateDomain(scannerentry.Text())
	}
}

func validateDomain(domainEntry string){
	var isMuxPresent,isSpfCheckPresent,isDmarchPresent bool//default value is false here
	mxrecordChecker,err:= net.LookupMX(domainEntry)
	if err!=nil{
		log.Println("Error while finding the mx record")
		return
	}
	if(len(mxrecordChecker)>0){
		isMuxPresent= true
	}
	spfrecordChecker,err :=net.LookupTXT(domainEntry)
	if err!=nil{
		log.Println("Error while fetching the spf record")
	}
	for _,ieterateOver:= range spfrecordChecker{
		if(strings.HasPrefix(ieterateOver, "v=spf")){
			isSpfCheckPresent= true
			break
		}
	}
	dmarchChecker,err := net.LookupTXT("_dmarc."+domainEntry)
	if err!=nil{
		log.Println("Error while finding the dmarc record")

	}
	fmt.Println(dmarchChecker)
	for _,dmarchiterate := range dmarchChecker{
		if(strings.HasPrefix(dmarchiterate,"v=DMARC1")){
			isDmarchPresent=true
		}	
	}


	fmt.Printf("%v %v %v",isMuxPresent,isSpfCheckPresent,isDmarchPresent)

}