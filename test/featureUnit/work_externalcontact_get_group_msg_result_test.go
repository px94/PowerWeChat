package featureUnit

import (
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"testing"
)

func Test_ExternalContact_Get_Group_MSG_Result(t *testing.T) {

	msgID := "msg_ViZBwAA72X9XCh4Cx5ku9OVFb2thQ" // walle
	//msgID := "msg_ViZBwAAT_5pzEfUhaT2xbSbFxhTgw"  // matt
	response, _ := Work.ExternalContactGroupChat.Get(msgID, false)

	if response == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error uniformMessage as :", response.ErrMsg)
	}

	fmt.Dump(response)

}
