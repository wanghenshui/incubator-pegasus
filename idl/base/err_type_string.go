// Code generated by "enumer -type=ErrType -output=err_type_string.go"; DO NOT EDIT

package base

import (
	"fmt"
)

const _ErrType_name = "ERR_OKERR_UNKNOWNERR_REPLICATION_FAILUREERR_APP_EXISTERR_APP_NOT_EXISTERR_APP_DROPPEDERR_BUSY_CREATINGERR_BUSY_DROPPINGERR_EXPIREDERR_LOCK_ALREADY_EXISTERR_HOLD_BY_OTHERSERR_RECURSIVE_LOCKERR_NO_OWNERERR_NODE_ALREADY_EXISTERR_INCONSISTENT_STATEERR_ARRAY_INDEX_OUT_OF_RANGEERR_SERVICE_NOT_FOUNDERR_SERVICE_ALREADY_RUNNINGERR_IO_PENDINGERR_TIMEOUTERR_SERVICE_NOT_ACTIVEERR_BUSYERR_NETWORK_INIT_FAILEDERR_FORWARD_TO_OTHERSERR_OBJECT_NOT_FOUNDERR_HANDLER_NOT_FOUNDERR_LEARN_FILE_FAILEDERR_GET_LEARN_STATE_FAILEDERR_INVALID_VERSIONERR_INVALID_PARAMETERSERR_CAPACITY_EXCEEDEDERR_INVALID_STATEERR_INACTIVE_STATEERR_NOT_ENOUGH_MEMBERERR_FILE_OPERATION_FAILEDERR_HANDLE_EOFERR_WRONG_CHECKSUMERR_INVALID_DATAERR_INVALID_HANDLEERR_INCOMPLETE_DATAERR_VERSION_OUTDATEDERR_PATH_NOT_FOUNDERR_PATH_ALREADY_EXISTERR_ADDRESS_ALREADY_USEDERR_STATE_FREEZEDERR_LOCAL_APP_FAILUREERR_BIND_IOCP_FAILEDERR_NETWORK_START_FAILEDERR_NOT_IMPLEMENTEDERR_CHECKPOINT_FAILEDERR_WRONG_TIMINGERR_NO_NEED_OPERATEERR_CORRUPTIONERR_TRY_AGAINERR_CLUSTER_NOT_FOUNDERR_CLUSTER_ALREADY_EXISTERR_ZOOKEEPER_OPERATIONERR_K8S_CLUSTER_NOT_FOUNDERR_K8S_KUBECTL_NOT_FOUNDERR_K8S_DEPLOY_FAILEDERR_K8S_UNDEPLOY_FAILEDERR_RESOURCE_NOT_ENOUGHERR_WIN_DEPLOY_FAILEDERR_WIN_UNDEPLOY_FAILEDERR_DOCKER_DAEMON_NOT_FOUNDERR_DOCKER_BINARY_NOT_FOUNDERR_DOCKER_DEPLOY_FAILEDERR_DOCKER_UNDEPLOY_FAILEDERR_FS_INTERNALERR_CLIENT_FAILED"

var _ErrType_index = [...]uint16{0, 6, 17, 40, 53, 70, 85, 102, 119, 130, 152, 170, 188, 200, 222, 244, 272, 293, 320, 334, 345, 367, 375, 398, 419, 439, 460, 481, 507, 526, 548, 569, 586, 604, 625, 650, 664, 682, 698, 716, 735, 755, 773, 795, 819, 836, 857, 877, 901, 920, 941, 957, 976, 990, 1003, 1024, 1049, 1072, 1097, 1122, 1143, 1166, 1189, 1210, 1233, 1260, 1287, 1311, 1337, 1352, 1369}

func (i ErrType) String() string {
	if i < 0 || i >= ErrType(len(_ErrType_index)-1) {
		return fmt.Sprintf("ErrType(%d)", i)
	}
	return _ErrType_name[_ErrType_index[i]:_ErrType_index[i+1]]
}

var _ErrTypeNameToValue_map = map[string]ErrType{
	_ErrType_name[0:6]:       0,
	_ErrType_name[6:17]:      1,
	_ErrType_name[17:40]:     2,
	_ErrType_name[40:53]:     3,
	_ErrType_name[53:70]:     4,
	_ErrType_name[70:85]:     5,
	_ErrType_name[85:102]:    6,
	_ErrType_name[102:119]:   7,
	_ErrType_name[119:130]:   8,
	_ErrType_name[130:152]:   9,
	_ErrType_name[152:170]:   10,
	_ErrType_name[170:188]:   11,
	_ErrType_name[188:200]:   12,
	_ErrType_name[200:222]:   13,
	_ErrType_name[222:244]:   14,
	_ErrType_name[244:272]:   15,
	_ErrType_name[272:293]:   16,
	_ErrType_name[293:320]:   17,
	_ErrType_name[320:334]:   18,
	_ErrType_name[334:345]:   19,
	_ErrType_name[345:367]:   20,
	_ErrType_name[367:375]:   21,
	_ErrType_name[375:398]:   22,
	_ErrType_name[398:419]:   23,
	_ErrType_name[419:439]:   24,
	_ErrType_name[439:460]:   25,
	_ErrType_name[460:481]:   26,
	_ErrType_name[481:507]:   27,
	_ErrType_name[507:526]:   28,
	_ErrType_name[526:548]:   29,
	_ErrType_name[548:569]:   30,
	_ErrType_name[569:586]:   31,
	_ErrType_name[586:604]:   32,
	_ErrType_name[604:625]:   33,
	_ErrType_name[625:650]:   34,
	_ErrType_name[650:664]:   35,
	_ErrType_name[664:682]:   36,
	_ErrType_name[682:698]:   37,
	_ErrType_name[698:716]:   38,
	_ErrType_name[716:735]:   39,
	_ErrType_name[735:755]:   40,
	_ErrType_name[755:773]:   41,
	_ErrType_name[773:795]:   42,
	_ErrType_name[795:819]:   43,
	_ErrType_name[819:836]:   44,
	_ErrType_name[836:857]:   45,
	_ErrType_name[857:877]:   46,
	_ErrType_name[877:901]:   47,
	_ErrType_name[901:920]:   48,
	_ErrType_name[920:941]:   49,
	_ErrType_name[941:957]:   50,
	_ErrType_name[957:976]:   51,
	_ErrType_name[976:990]:   52,
	_ErrType_name[990:1003]:  53,
	_ErrType_name[1003:1024]: 54,
	_ErrType_name[1024:1049]: 55,
	_ErrType_name[1049:1072]: 56,
	_ErrType_name[1072:1097]: 57,
	_ErrType_name[1097:1122]: 58,
	_ErrType_name[1122:1143]: 59,
	_ErrType_name[1143:1166]: 60,
	_ErrType_name[1166:1189]: 61,
	_ErrType_name[1189:1210]: 62,
	_ErrType_name[1210:1233]: 63,
	_ErrType_name[1233:1260]: 64,
	_ErrType_name[1260:1287]: 65,
	_ErrType_name[1287:1311]: 66,
	_ErrType_name[1311:1337]: 67,
	_ErrType_name[1337:1352]: 68,
	_ErrType_name[1352:1369]: 69,
}

func ErrTypeString(s string) (ErrType, error) {
	if val, ok := _ErrTypeNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ErrType values", s)
}
