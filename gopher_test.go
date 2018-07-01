package gopher_test

import (
	"goGetGoingWithGo/03_documentation_testing/01_godoc/gopher"
	"os"
)

func ExampleWrite() {
	gopher.Write(os.Stdout) // write a gopher to stdout
	// Output:
	//                                         '.--://++++++++///::-.'
	//                              '.:/+syhdmmmmmmmmmmmmmmmmmmmmmmmmmdhyo+:.
	//                         ':+shddmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmddyo/.
	//                     ':ohdmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmdddddmmmmmmdho-      .--:-.'
	//      '.::::-.    '/ydmmmmdhyysssyydmmmmmmmmmmmmmmmmmmmmmmhsssssyyyyssssshmmmmdh+..shdddddddy/.
	//   '/ydddmmmddh--sdmmmdsssshdmmmmmdysssymmmmmmmmmmmmmmmhooymNMMMmdmNMMMNmyosmmmmmds/ymmmmmmmmmdo'
	// '+dmmmmmmmmms/ydmmmyosdNMMMMMdyyhNMMMNhosdmmmmmmmmmmdosmMMMMMh-'''.oNMMMMNh/hmmmmmdo/++oymmmmmmy'
	//'ymmmmdsooyh/sdmmmd+yNMMMMMMm-'   .oMMMMMd/hmmmmmmmmh/dMMMMMMd'      oMMMMMMN/ymmmmmmh-  'ommmmmms
	//smmmmy.    :hmmmmd/dMMMMMMMM:       dMMMMMN/hmmmmmmh:NMMMMMMMd   /-  +MMMMMMMN:dmmmmmmd/  :mmmmmmm
	//dmmmm/    :dmmmmm/dMMMMMMMMM/  -o' 'mMMMMMMm:mmmmmm/dMMMMMMMMMs.'y+'/NMMMMMMMMsommmmmmmd+-dmmmmmmh
	//ymmmmd/. :dmmmmmd:MMMMMMMMMMNo-/y-:dMMMMMMMM:dmmmmm.MMMMMMMMMMMNdhdmMMMMMMMMMMd/mmmmmmmmd:ymmmmmd-
	//.hmmmmms-dmmmmmmh/MMMMMMMMMMMMNmmmMMMMMMMMMM/hmmmmm.MMMMMMMMMMMMMMMMMMMMMMMMMMh/mmmmmmmmmh-dmmds.
	// '+hmmd-hmmmmmmmd-MMMMMMMMMMMMMMMMMMMMMMMMMM-dmmmmmosMMMMMMMMMMMMMMMMMMMMMMMMM/ymmmmmmmmmmo+s/.
	//   '-+/+mmmmmmmmm+yMMMMMMMMMMMMMMMMMMMMMMMMsommmmmmd/hMMMMMMMMMMMMMMMMMMMMMMNoommmmmmmmmmmd.
	//      'dmmmmmmmmmd/hMMMMMMMMMMMMMMMMMMMMMNs+mmmmdddddosmMMMMMMMMMMMMMMMMMMMd+smmmmmmmmmmmmmo
	//      +mmmmmmmmmmmdoomMMMMMMMMMMMMMMMMMNh+ymdy+:-.--:ososdNMMMMMMMMMMMMNmhosdmmmmmmmmmmmmmmd'
	//      hmmmmmmmmmmmmmhsohmNMMMMMMMMMNmdsoydmh-         'ydyssyhddmmddhyssshmmmmmmmmmmmmmmmmmm/
	//     'mmmmmmmmmmmmmmmmdyysssyyyyysssyydmdhs:          ':yhmmdhyyyyyyhdmmmmmmmmmmmmmmmmmmmmmmy
	//     :mmmmmmmmmmmmmmmmmmmmmmdddddmmmmmmy/oyh+-.'''..:+yhyo+ydmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd
	//     /mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmh-hdddddddhhdddddddddo/dmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm-
	//     +mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmo/dddddddddddddddddddd:ymmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm/
	//     ommmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd:shhhhyoo+-sooshhddho/dmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmo
	//     ommmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmyo:oydNMh-MMMm+:oooymmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmms
	//     +mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm-mMMMMs/MMMMh+mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmy
	//     +mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm+yMMMMs/MMMMd/mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmh
	//     :mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd:dMMd/:dMMMoommmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd
	//     -mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmdsoosmdsoooymmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd
	//wrote 2724 bytes worth of gopherDome!
}

//func ExampleWrite_second() {
//	// write a gopher to a file
//	file, err := os.Create("./gopher.txt")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	gopher.Write(file)
//}
//
//func ExampleWrite_third() {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		gopher.Write(w)
//	})
//	srv := &http.Server{
//		Addr:    ":8080",
//		Handler: mux,
//	}
//	fmt.Println("Starting HTTP server. RFC 7230 baby!")
//	log.Fatal(srv.ListenAndServe())
//}
