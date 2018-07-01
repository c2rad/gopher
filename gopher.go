// package gopher provides the ability to write gigantic gopher DomePieces for demonstration purposes
package gopher

import (
	"fmt"
	"io"
)

// Write takes an io.Writer and will call the Write() method.
// It will write a gigantic gopher head.
func Write(i io.Writer) {
	num, err := i.Write(DomePiece)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("wrote %d bytes worth of gopherDome!\n", num)
}

var DomePiece = []byte(`
                                          '.--://++++++++///::-.'                                    
                               '.:/+syhdmmmmmmmmmmmmmmmmmmmmmmmmmdhyo+:.                            
                          ':+shddmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmddyo/.                       
                      ':ohdmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmdddddmmmmmmdho-      .--:-.'       
       '.::::-.    '/ydmmmmdhyysssyydmmmmmmmmmmmmmmmmmmmmmmhsssssyyyyssssshmmmmdh+..shdddddddy/.    
    '/ydddmmmddh--sdmmmdsssshdmmmmmdysssymmmmmmmmmmmmmmmhooymNMMMmdmNMMMNmyosmmmmmds/ymmmmmmmmmdo'  
  '+dmmmmmmmmms/ydmmmyosdNMMMMMdyyhNMMMNhosdmmmmmmmmmmdosmMMMMMh-'''.oNMMMMNh/hmmmmmdo/++oymmmmmmy' 
 'ymmmmdsooyh/sdmmmd+yNMMMMMMm-'   .oMMMMMd/hmmmmmmmmh/dMMMMMMd'      oMMMMMMN/ymmmmmmh-  'ommmmmms 
 smmmmy.    :hmmmmd/dMMMMMMMM:       dMMMMMN/hmmmmmmh:NMMMMMMMd   /-  +MMMMMMMN:dmmmmmmd/  :mmmmmmm 
 dmmmm/    :dmmmmm/dMMMMMMMMM/  -o' 'mMMMMMMm:mmmmmm/dMMMMMMMMMs.'y+'/NMMMMMMMMsommmmmmmd+-dmmmmmmh 
 ymmmmd/. :dmmmmmd:MMMMMMMMMMNo-/y-:dMMMMMMMM:dmmmmm.MMMMMMMMMMMNdhdmMMMMMMMMMMd/mmmmmmmmd:ymmmmmd- 
 .hmmmmms-dmmmmmmh/MMMMMMMMMMMMNmmmMMMMMMMMMM/hmmmmm.MMMMMMMMMMMMMMMMMMMMMMMMMMh/mmmmmmmmmh-dmmds.  
  '+hmmd-hmmmmmmmd-MMMMMMMMMMMMMMMMMMMMMMMMMM-dmmmmmosMMMMMMMMMMMMMMMMMMMMMMMMM/ymmmmmmmmmmo+s/.    
    '-+/+mmmmmmmmm+yMMMMMMMMMMMMMMMMMMMMMMMMsommmmmmd/hMMMMMMMMMMMMMMMMMMMMMMNoommmmmmmmmmmd.       
       'dmmmmmmmmmd/hMMMMMMMMMMMMMMMMMMMMMNs+mmmmdddddosmMMMMMMMMMMMMMMMMMMMd+smmmmmmmmmmmmmo       
       +mmmmmmmmmmmdoomMMMMMMMMMMMMMMMMMNh+ymdy+:-.--:ososdNMMMMMMMMMMMMNmhosdmmmmmmmmmmmmmmd'      
       hmmmmmmmmmmmmmhsohmNMMMMMMMMMNmdsoydmh-         'ydyssyhddmmddhyssshmmmmmmmmmmmmmmmmmm/      
      'mmmmmmmmmmmmmmmmdyysssyyyyysssyydmdhs:          ':yhmmdhyyyyyyhdmmmmmmmmmmmmmmmmmmmmmmy      
      :mmmmmmmmmmmmmmmmmmmmmmdddddmmmmmmy/oyh+-.'''..:+yhyo+ydmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd      
      /mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmh-hdddddddhhdddddddddo/dmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm-     
      +mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmo/dddddddddddddddddddd:ymmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm/     
      ommmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd:shhhhyoo+-sooshhddho/dmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmo     
      ommmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmyo:oydNMh-MMMm+:oooymmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmms     
      +mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm-mMMMMs/MMMMh+mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmy     
      +mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm+yMMMMs/MMMMd/mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmh     
      :mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd:dMMd/:dMMMoommmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd     
      -mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmdsoosmdsoooymmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmd
`)
